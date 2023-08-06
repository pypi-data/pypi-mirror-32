#
#    Copyright (c) 2012+ Anton Tyurin <noxiouz@yandex.ru>
#    Copyright (c) 2013+ Evgeny Safronov <division494@gmail.com>
#    Copyright (c) 2011-2014 Other contributors as noted in the AUTHORS file.
#
#    This file is part of Cocaine.
#
#    Cocaine is free software; you can redistribute it and/or modify
#    it under the terms of the GNU Lesser General Public License as published by
#    the Free Software Foundation; either version 3 of the License, or
#    (at your option) any later version.
#
#    Cocaine is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
#    GNU Lesser General Public License for more details.
#
#    You should have received a copy of the GNU Lesser General Public License
#    along with this program. If not, see <http://www.gnu.org/licenses/>.
#

import functools
import itertools
import json
import logging
import threading
import warnings

import six
from six.moves import cStringIO as BytesIO

from tornado import gen
from tornado import queues
from tornado.gen import coroutine
from tornado.ioloop import IOLoop
from tornado.locks import Lock
from tornado.tcpclient import TCPClient

from .api import API
from .defaults import Defaults
from .defaults import GetOptError
from .util import msgpack_pack, msgpack_packb, msgpack_unpacker


__all__ = ["Logger", "CocaineHandler"]

LOCATOR_DEFAULT_ENDPOINTS = Defaults.locators

(DEBUG_LEVEL, INFO_LEVEL, WARNING_LEVEL, ERROR_LEVEL) = range(4)

# look at Locator and LoggerAPI
EMIT = 0
VERBOSITY = 1

RESOLVE = 0
VALUE_CODE = 0
ERROR_CODE = 1
assert API.Logger[EMIT][0] == b"emit"
assert API.Locator[RESOLVE][0] == b"resolve"

ATTRS_TYPES = six.string_types + six.integer_types + (float, bool)


def thread_once(class_init):
    @functools.wraps(class_init)
    def wrapper(self, *args, **kwargs):
        if getattr(self._current, "initialized", False):
            return

        class_init(self, *args, **kwargs)
        self._current.initialized = True
    return wrapper


fallback_logger = logging.getLogger("fallback")
fallback_logger.propagate = False
fallback_logger.setLevel(logging.DEBUG)


class Logger(object):
    _name = "logging"
    _current = threading.local()

    def __new__(cls, *args, **kwargs):
        if not getattr(cls._current, "instance", None):
            cls._current.instance = object.__new__(cls, *args, **kwargs)
        return cls._current.instance

    @thread_once
    def __init__(self, endpoints=LOCATOR_DEFAULT_ENDPOINTS, io_loop=None):
        if io_loop:
            warnings.warn('io_loop argument is deprecated.', DeprecationWarning)
        self.io_loop = io_loop or IOLoop.current()
        self.endpoints = endpoints
        self._lock = Lock()

        self.counter = itertools.count(1)

        self.pipe = None
        self.target = Defaults.app
        self.verbosity = DEBUG_LEVEL
        self.queue = queues.Queue(10000)

        # level could be reset from update_verbosity in the future
        if not fallback_logger.handlers:
            sh = logging.StreamHandler()
            sh.setFormatter(logging.Formatter(fmt="[%(asctime)s.%(msecs)d] %(levelname)s fallback %(message)s", datefmt="%z %d/%b/%Y:%H:%M:%S"))
            sh.setLevel(logging.DEBUG)
            fallback_logger.addHandler(sh)

        self._send()
        try:
            uuid = Defaults.uuid
            self._defaultattrs = [("uuid", uuid)]
        except GetOptError:
            self._defaultattrs = []

    def prepare_message_args(self, level, message, *args, **kwargs):
        if args:
            try:
                message %= args
            except Exception:
                message = "unformatted: %s %s" % (message, args)
                level = ERROR_LEVEL

        if "extra" not in kwargs:
            if self._defaultattrs:
                msg = [level, self.target, message, self._defaultattrs]
            else:
                msg = [level, self.target, message]
        else:
            attrs = [(str(k), (v if isinstance(v, ATTRS_TYPES) else str(v))) for k, v in six.iteritems(kwargs["extra"])]
            msg = [level, self.target, message, attrs + self._defaultattrs]

        return msg

    def emit(self, level, message, *args, **kwargs):
        msg = self.prepare_message_args(level, message, *args, **kwargs)
        # if the queue is full log new messages to the fallback Logger
        # to make most recent errors be printed at least to stderr
        try:
            self.queue.put_nowait(msg)
        except queues.QueueFull:
            self._log_to_fallback(msg)

    @coroutine
    def _send(self):
        """ Send a message lazy formatted with args.
        External log attributes can be passed via named attribute `extra`,
        like in logging from the standart library.

        Note:
            * Attrs must be dict, otherwise the whole message would be skipped.
            * The key field in an attr is converted to string.
            * The value is sent as is if isinstance of (str, unicode, int, float, long, bool),
              otherwise we convert the value to string.
        """
        buff = BytesIO()
        while True:
            msgs = list()
            try:
                msg = yield self.queue.get()

                # we need to connect first, as we issue verbosity request just after connection
                # and channels should strictly go in ascending order
                if not self._connected:
                    yield self.connect()

                try:
                    while True:
                        msgs.append(msg)
                        counter = next(self.counter)
                        msgpack_pack([counter, EMIT, msg], buff)
                        msg = self.queue.get_nowait()
                except queues.QueueEmpty:
                    pass

                try:
                    yield self.pipe.write(buff.getvalue())
                except Exception:
                    pass
                # clean the buffer or we will end up without memory
                buff.truncate(0)
            except Exception:
                for message in msgs:
                    self._log_to_fallback(message)

    def _log_to_fallback(self, message):
        level, target, text, attrs = message
        if level >= ERROR_LEVEL:
            actual_level = logging.ERROR
        elif level >= WARNING_LEVEL:
            actual_level = logging.WARNING
        elif level >= INFO_LEVEL:
            actual_level = logging.INFO
        else:
            actual_level = logging.DEBUG
        fallback_logger.log(actual_level, "%s %s %s", target, text, json.dumps(attrs))

    def debug(self, message, *args, **kwargs):
        if self.enable_for(DEBUG_LEVEL):
            self.emit(DEBUG_LEVEL, message, *args, **kwargs)

    def warn(self, message, *args, **kwargs):
        self.warning(message, *args, **kwargs)

    def warning(self, message, *args, **kwargs):
        if self.enable_for(WARNING_LEVEL):
            self.emit(WARNING_LEVEL, message, *args, **kwargs)

    def info(self, message, *args, **kwargs):
        if self.enable_for(INFO_LEVEL):
            self.emit(INFO_LEVEL, message, *args, **kwargs)

    def error(self, message, *args, **kwargs):
        if self.enable_for(ERROR_LEVEL):
            self.emit(ERROR_LEVEL, message, *args, **kwargs)

    def enable_for(self, level):
        return self.verbosity <= level

    @coroutine
    def update_verbosity(self):
        counter = next(self.counter)
        verbosity_request = msgpack_packb([counter, VERBOSITY, []])
        self.pipe.write(verbosity_request)
        buff = msgpack_unpacker()
        while True:
            data = yield self.pipe.read_bytes(1024, partial=True)
            buff.feed(data)
            for msg in buff:
                _, code, payload = msg[:3]
                if code == VALUE_CODE:
                    self.verbosity = payload[0]
                else:
                    self.verbosity = DEBUG_LEVEL
                return

    @coroutine
    def connect(self):
        with (yield self._lock.acquire()):
            if self._connected:
                return

            for host, port in (yield resolve_logging(self.endpoints, self._name,
                                                     self.io_loop)):
                try:
                    self.pipe = yield TCPClient(io_loop=self.io_loop).connect(host, port)
                    self.pipe.set_nodelay(True)
                    yield self.update_verbosity()
                    return
                except IOError:
                    pass

    @property
    def _connected(self):
        return self.pipe is not None and not self.pipe.closed()

    def disconnect(self):
        if self.pipe is None:
            return

        self.pipe.close()
        self.pipe = None

    def __del__(self):
        # we have to close owned connection
        # otherwise it would be a fd-leak
        self.disconnect()


@coroutine
def resolve_logging(endpoints, name="logging", io_loop=None):
    if io_loop:
        warnings.warn('io_loop argument is deprecated.', DeprecationWarning)

    for host, port in endpoints:
        buff = msgpack_unpacker()
        locator_pipe = None
        try:
            locator_pipe = yield TCPClient(io_loop=io_loop).connect(host, port)
            locator_pipe.set_nodelay(True)
            request = msgpack_packb([999999, RESOLVE, [name]])
            locator_pipe.write(request)
            while True:
                data = yield locator_pipe.read_bytes(1024, partial=True)
                buff.feed(data)
                for msg in buff:
                    _, code, payload = msg[:3]
                    if code == VALUE_CODE:
                        raise gen.Return(payload[0])
        except (IOError, ValueError):
            pass
        finally:
            if locator_pipe:
                locator_pipe.close()

    raise Exception("unable to resolve logging")


class CocaineHandler(logging.Handler):
    def __init__(self, *args, **kwargs):
        logging.Handler.__init__(self)
        self._logger = Logger(*args, **kwargs)

    def emit(self, record):
        lvl = record.levelno
        extra = getattr(record, "extra", {})
        if lvl >= logging.ERROR:
            # to avoid message formatting
            if self._logger.enable_for(ERROR_LEVEL):
                self._logger.error(self.format(record), extra=extra)
        elif lvl >= logging.WARNING:
            if self._logger.enable_for(WARNING_LEVEL):
                self._logger.warning(self.format(record), extra=extra)
        elif lvl >= logging.INFO:
            if self._logger.enable_for(INFO_LEVEL):
                self._logger.info(self.format(record), extra=extra)
        elif lvl >= logging.DEBUG:
            if self._logger.enable_for(DEBUG_LEVEL):
                self._logger.debug(self.format(record), extra=extra)


class LoggerWithExtraInRecord(logging.getLoggerClass()):
    def makeRecord(self, name, level, fn, lno, msg, args, exc_info, func=None, extra=None):  # noqa
        rv = super(LoggerWithExtraInRecord, self).makeRecord(name, level, fn, lno, msg, args, exc_info, func, extra)
        if extra is not None:
            rv.__dict__["extra"] = extra
        return rv
