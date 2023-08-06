import atexit
import inspect
import os
import pickle
import signal
from functools import wraps
from multiprocessing import Process
from typing import Callable
from uuid import UUID, uuid1

import zmq
from time import sleep

from zproc.util import shutdown, get_ipc_paths, Message, STATE_DICT_DYNAMIC_METHODS, \
    serialize_func, handle_server_response, method_injector, print_crash_report, signal_to_exception, \
    SignalException
from zproc.zproc_server import ZProcServer, zproc_server_proc

signal.signal(signal.SIGTERM, shutdown)  # Add handler for the SIGTERM signal


def ping(uuid: UUID, **kwargs):
    """
    Ping the ZProc Server

    :param uuid: The ``UUID`` object for identifying the Context.

                 You can use the UUID to reconstruct a ZeroState object, from any Process, at any time.
    :param timeout: (optional) Sets the timeout in seconds. By default it is set to -1.

                    If the value is 0, it will return immediately, with a :py:exc:`TimeoutError`.

                    If the value is -1, it will block until the ZProc Server replies.

                    For all other values, it will wait for a reply,
                    for that amount of time before returning with a :py:exc:`TimeoutError`.
    :param ping_data: (optional) Data to ping the server with.

                      By default, it is set to ``os.urandom(56)`` (56 random bytes of data).
    :return: The ZProc Server's ``pid`` if the ping was successful, else :py:class:`False`
    """

    kwargs.setdefault('timeout', -1)
    kwargs.setdefault('ping_data', os.urandom(56))

    ctx = zmq.Context()
    ctx.setsockopt(zmq.LINGER, 0)
    sock = ctx.socket(zmq.DEALER)
    sock.connect(get_ipc_paths(uuid)[0])

    if kwargs['timeout'] != -1:
        sock.setsockopt(zmq.RCVTIMEO, int(kwargs['timeout'] * 1000))

    ping_msg = {
        Message.server_action: ZProcServer.ping.__name__,
        Message.ping_data: kwargs['ping_data']
    }

    sock.send_pyobj(ping_msg, protocol=pickle.HIGHEST_PROTOCOL)  # send msg

    try:
        response = sock.recv_pyobj()  # wait for reply
    except zmq.error.Again:
        raise TimeoutError('Connection to ZProc Server timed out!')
    else:
        if response.get('ping_data') == kwargs['ping_data']:
            return response['pid']
        else:
            return False
    finally:
        sock.close()


class ZeroState:
    def __init__(self, uuid: UUID):
        """
        :param uuid: The ``UUID`` object for identifying the Context.

                     You can use the UUID to reconstruct a ZeroState object, from any Process, at any time.

        :ivar uuid: Passed on from constructor

        | Allows accessing state stored on the ZProc Server, through a dict-like API.
        | Communicates to the ZProc Server using the ROUTER-DEALER pattern.

        | Don't share a ZeroState object between Process/Threads.
        | A ZeroState object is not thread-safe.

        Boasts the following ``dict``-like members:

        - Magic  methods:
            __contains__(),  __delitem__(), __eq__(), __getitem__(), __iter__(), __len__(), __ne__(), __setitem__()
        - Methods:
            clear(), copy(), fromkeys(), get(), items(),  keys(), pop(), popitem(), setdefault(), update(), values()
        """

        self.uuid = uuid
        self.identity = os.urandom(5)

        self.zmq_ctx = zmq.Context()
        self.zmq_ctx.setsockopt(zmq.LINGER, 0)

        self.server_ipc_path, self.bcast_ipc_path = get_ipc_paths(self.uuid)

        self.server_sock = self.zmq_ctx.socket(zmq.DEALER)
        self.server_sock.setsockopt(zmq.IDENTITY, self.identity)
        self.server_sock.connect(self.server_ipc_path)

        self.sub_sock = self._new_sub_sock()

    def _new_sub_sock(self):
        sock = self.zmq_ctx.socket(zmq.SUB)
        sock.connect(self.bcast_ipc_path)

        sock.setsockopt(zmq.SUBSCRIBE, b'')

        return sock

    def _sub_recv(self, sock):
        while True:
            ident, response = sock.recv_multipart()

            if ident != self.identity:
                return handle_server_response(pickle.loads(response))

    def _req_rep(self, request):
        self.server_sock.send_pyobj(request, protocol=pickle.HIGHEST_PROTOCOL)  # send msg
        response = self.server_sock.recv_pyobj()  # wait for reply
        return handle_server_response(response)

    def _get_state_method(self, name):
        def state_method(*args, **kwargs):
            return self._req_rep({
                Message.server_action: ZProcServer.state_method.__name__,
                Message.method_name: name,
                Message.args: args,
                Message.kwargs: kwargs
            })

        return state_method

    def go_live(self):
        self.sub_sock.close()
        self.sub_sock = self._new_sub_sock()

    def task(self, fn, *args, **kwargs):
        """
        | Run a "task" on the state, atomically.
        | In other words, execute the ``fn`` in an *atomic* way.
        |
        | No Process shall access the state in any way whilst ``fn`` is running.
        | This helps avoid race-conditions, almost entirely.

        :param fn: A user-defined function.
        :param args: Passed on to fn
        :param kwargs: Passed on to fn
        :return: Roughly, ``fn(state, *args, **kwargs)``

        .. code:: python

            from zproc import Context

            ctx = Context()
            state = ctx.state

            state['count'] = 0

            def increment(state):
                return state['count'] += 1

            print(state.atomic(increment)) # 1
        """

        return self._req_rep({
            Message.server_action: ZProcServer.state_func.__name__,
            Message.func: serialize_func(fn),
            Message.args: args,
            Message.kwargs: kwargs
        })

    def taskify(self):
        """
        | Hack on a normal looking function to make an atomic "task" out of it.
        | Allows making an arbitrary number of operations on sate, atomic.
        |
        | Decorator version of :py:meth:`~ZeroState.task()`

        :return: An atomic decorator, which itself returns - ``wrapped_fn(state, *args, **kwargs)``
        :rtype: function

        .. code:: python

            from zproc import Context

            ctx = Context()
            state = ctx.state

            state['count'] = 0

            @atomify()
            def increment(state):
                return state['count'] += 1

            print(increment()) # 1
        """

        def atomic_decorator(fn):
            @wraps(fn)
            def atomic_wrapper(*args, **kwargs):
                return self.task(fn, *args, **kwargs)

            return atomic_wrapper

        return atomic_decorator

    def get_when_change(self, *keys, live=True, timeout=-1):
        """
        | Block until a change is observed.

        :param keys: Observe for changes in these ``dict`` key(s).

                     If no key is provided, any change in the ``state`` is respected.
        :param live: Whether to get "live" updates. See :ref:`state_watching`.
        :param timeout: Sets the timeout in seconds. By default it is set to -1.

                        If the value is 0, it will return immediately, with a :py:exc:`TimeoutError`,
                        if no update is available.

                        If the value is -1, it will block until an update is available.

                        For all other values, it will wait for and update,
                        for that amount of time before returning with a :py:exc:`TimeoutError`.

        :return: Roughly,

        .. code:: python

            if len(keys) == 1
                return state.get(key)

            if len(keys) > 1
                return [state.get(key) for key in keys]

            if len(keys) == 0
                return state.copy()
        """

        if live:
            sock = self._new_sub_sock()
        else:
            sock = self.sub_sock

        if timeout != -1:
            sock.setsockopt(zmq.RCVTIMEO, int(timeout * 1000))

        num_keys = len(keys)
        try:
            if num_keys == 1:
                key = keys[0]

                while True:
                    old, new = self._sub_recv(sock)

                    if new.get(key) != old.get(key):
                        return new
            elif num_keys:
                while True:
                    old, new = self._sub_recv(sock)

                    for key in keys:
                        if new.get(key) != old.get(key):
                            return [new.get(key) for key in keys]
            else:
                while True:
                    old, new = self._sub_recv(sock)

                    if new != old:
                        return new
        except zmq.error.Again:
            raise TimeoutError('Timed-out waiting for an update')
        finally:
            if live:
                sock.close()
            elif timeout != -1:
                sock.setsockopt(zmq.RCVTIMEO, -1)

    def get_when(self, test_fn, *, live=True, timeout=-1):
        """
        | Block until ``test_fn(state: ZeroState)`` returns a True-like value
        |
        | Roughly,

        .. code:: python

            if test_fn(state)
                return state.copy()

        :param test_fn: A ``function``, which is called on each state-change.
        :param live: Whether to get "live" updates. See :ref:`state_watching`.
        :param timeout: Sets the timeout in seconds. By default it is set to -1.

                        If the value is 0, it will return immediately, with a :py:exc:`TimeoutError`,
                        if no update is available.

                        If the value is -1, it will block until an update is available.

                        For all other values, it will wait for and update,
                        for that amount of time before returning with a :py:exc:`TimeoutError`.

        :return: state
        :rtype: ``dict``
        """

        if live:
            sock = self._new_sub_sock()
        else:
            sock = self.sub_sock

        if timeout != -1:
            sock.setsockopt(zmq.RCVTIMEO, int(timeout * 1000))

        new = self.copy()
        try:
            while True:
                if test_fn(new):
                    return new
                else:
                    new = self._sub_recv(sock)[-1]
        except zmq.error.Again:
            raise TimeoutError('Timed-out waiting for an update')
        finally:
            if live:
                sock.close()
            elif timeout != -1:
                sock.setsockopt(zmq.RCVTIMEO, -1)

    def get_when_equal(self, key, value, *, live=True, timeout=-1):
        """
        | Block until ``state.get(key) == value``.

        :param key: ``dict`` key
        :param value: ``dict`` value.
        :param live: Whether to get "live" updates. See :ref:`state_watching`.
        :param timeout: Sets the timeout in seconds. By default it is set to -1.

                        If the value is 0, it will return immediately, with a :py:exc:`TimeoutError`,
                        if no update is available.

                        If the value is -1, it will block until an update is available.

                        For all other values, it will wait for and update,
                        for that amount of time before returning with a :py:exc:`TimeoutError`.

        :return: ``state.get(key)``
        """

        if live:
            sock = self._new_sub_sock()
        else:
            sock = self.sub_sock

        if timeout != -1:
            sock.setsockopt(zmq.RCVTIMEO, int(timeout * 1000))

        new = self.get(key)
        try:
            while True:
                if new == value:
                    return new
                else:
                    new = self._sub_recv(sock)[-1].get(key)
        except zmq.error.Again:
            raise TimeoutError('Timed-out waiting for an update')
        finally:
            if live:
                sock.close()
            elif timeout != -1:
                sock.setsockopt(zmq.RCVTIMEO, -1)

    def get_when_not_equal(self, key, value, *, live=True, timeout=-1):
        """
        | Block until ``state.get(key) != value``.

        :param key: ``dict`` key.
        :param value: ``dict`` value.
        :param live: Whether to get "live" updates. See :ref:`state_watching`.
        :param timeout: Sets the timeout in seconds. By default it is set to -1.

                        If the value is 0, it will return immediately, with a :py:exc:`TimeoutError`,
                        if no update is available.

                        If the value is -1, it will block until an update is available.

                        For all other values, it will wait for and update,
                        for that amount of time before returning with a :py:exc:`TimeoutError`.

        :return: ``state.get(key)``
        """

        if live:
            sock = self._new_sub_sock()
        else:
            sock = self.sub_sock

        if timeout != -1:
            sock.setsockopt(zmq.RCVTIMEO, int(timeout * 1000))

        new = self.get(key)
        try:
            while True:
                if new != value:
                    return new
                else:
                    new = self._sub_recv(sock)[-1].get(key)
        except zmq.error.Again:
            raise TimeoutError('Timed-out waiting for an update')
        finally:
            if live:
                sock.close()
            elif timeout != -1:
                sock.setsockopt(zmq.RCVTIMEO, -1)

    def ping(self, **kwargs):
        """
        Ping the ZProc Server corresponding to this State's Context

        :param kwargs: Optional keyword arguments that :py:func:`ping` takes.
        :return: Same as :py:func:`ping`
        """

        return ping(self.uuid, **kwargs)

    def close(self):
        """
        Close this ZeroState and disconnect with the ZProcServer
        """

        self.server_sock.close()
        self.sub_sock.close()
        self.zmq_ctx.destroy()
        self.zmq_ctx.term()

    def copy(self):
        return self._req_rep({Message.server_action: ZProcServer.reply_state.__name__})

    def keys(self):
        return self.copy().keys()

    def items(self):
        return self.copy().items()

    def values(self):
        return self.copy().values()

    def __repr__(self):
        return '<ZeroState state: {} uuid: {}>'.format(self.copy(), self.uuid)


def _get_remote_method(name):
    def remote_method(self, *args, **kwargs):
        return self._get_state_method(name)(*args, **kwargs)

    return remote_method


method_injector(ZeroState, STATE_DICT_DYNAMIC_METHODS, _get_remote_method)


def _child_proc(proc: 'ZeroProcess'):
    state = ZeroState(proc.uuid)

    target_params = inspect.signature(proc.target).parameters.copy()

    if 'kwargs' in target_params:
        target_kwargs = {'state': state, 'props': proc.kwargs['props'], 'proc': proc}
    else:
        target_kwargs = {}
        if 'state' in target_params:
            target_kwargs['state'] = state
        if 'props' in target_params:
            target_kwargs['props'] = proc.kwargs['props']
        if 'proc' in target_params:
            target_kwargs['proc'] = proc

    exceptions = [SignalException]
    for i in proc.kwargs['retry_for']:
        if type(i) == signal.Signals:
            signal_to_exception(i)  # converts signal to exception!
        elif issubclass(i, BaseException):
            exceptions.append(i)
        else:
            raise ValueError(
                '"retry_for" must contain a `signals.Signal` or `Exception`, not `{}`'.format(repr(type(i))))

    exceptions = tuple(exceptions)

    tries = 0
    while True:
        try:
            tries += 1
            proc.target(**target_kwargs)
        except exceptions as e:
            print_crash_report(proc, e, proc.kwargs['retry_delay'], tries, proc.kwargs['max_retries'])

            if proc.kwargs['max_retries'] != -1 and tries > proc.kwargs['max_retries']:
                raise e
            else:
                if 'props' in target_kwargs:
                    target_kwargs['props'] = proc.kwargs['retry_props']

                sleep(proc.kwargs['retry_delay'])
        else:
            break


class ZeroProcess:
    def __init__(self, uuid: UUID, target: Callable, **kwargs):
        """
        Provides a high level wrapper over :py:class:`Process`.

        :param uuid: The :py:class:`UUID` object for identifying the Context.

                     You may use this to reconstruct a :py:class:`ZeroState` object, from any Process, at any time.

        :param target: The Callable to be invoked by the start() method.

                        | It is run inside a :py:class:`Process` with following kwargs:
                        | ``target(state=<ZeroState>, props=<props>, proc=<ZeroProcess>)``

                        | You may omit these, shall you not want them.

        :param props: (optional) Passed on to the target. By default, it is set to :py:class:`None`

        :param start: (optional) Automatically call :py:meth:`.start()` on the process.
                      By default, it is set to ``True``.

        :param retry_for: (optional) Automatically retry  whenever a particular Exception / signal is raised.
                          By default, it is an empty ``tuple``

                          .. code:: python

                                import signal

                                # retry if a ConnectionError, ValueError or signal.SIGTERM is received.
                                ctx.process(
                                    lambda: None,
                                    retry_for=(ConnectionError, ValueError, signal.SIGTERM)
                                )

                          | To retry for *any* **Exception**, just do
                          | ``retry_for=(Exception,)``

                          There is no such feature as retry for *any* **signal**, for now.

        :param retry_delay: (optional) The delay in seconds, before auto-retrying. By default, it is set to 5 secs

        :param retry_props: (optional) Provide different ``props`` to a Process when it's retrying.
                            By default, it is the same as ``props``

                            Used to control how your application behaves, under retry conditions.


        :param max_retries: (optional) Give up after this many attempts. By default, it is set to ``-1``.

                            A value of ``-1`` will result in an *infinite* number of retries.

                            After "max_tries", any Exception will be raised normally.

        :ivar uuid: Passed on from constructor.
        :ivar target: Passed on from constructor.
        :ivar kwargs: Passed on from constructor.
        :ivar proc: The :py:class:`Process` object associated with this ZeroProcess.
        """

        assert callable(target), '"target" must be a `Callable`, not `{}`'.format(type(target))

        self.uuid = uuid
        self.target = target
        self.kwargs = kwargs

        self.kwargs.setdefault('props', None)
        self.kwargs.setdefault('start', True)
        self.kwargs.setdefault('retry_for', ())
        self.kwargs.setdefault('retry_delay', 5)
        self.kwargs.setdefault('retry_props', self.kwargs['props'])
        self.kwargs.setdefault('max_retries', -1)

        self.proc = Process(target=_child_proc, args=(self,))

        if self.kwargs['start']:
            self.start()

    def __repr__(self):
        return '<ZeroProcess target: {} uuid: {}>'.format(self.target, self.uuid)

    def start(self):
        """
        Start this Process

        If the child has already been started once, it will return with an :py:exc:`AssertionError`.

        :return: the process PID
        """

        self.proc.start()
        return self.proc.pid

    def stop(self):
        """
        Stop this process

        :return: :py:attr:`~exitcode`.
        """

        self.proc.terminate()
        return self.proc.exitcode

    @property
    def is_alive(self):
        """
        Whether the child process is alive.

        | Roughly, a process object is alive;
        | from the moment the start() method returns,
        | until the child process is stopped manually (using stop()) or naturally exits
        """

        return self.proc and self.proc.is_alive()

    @property
    def pid(self):
        """
        | The process ID.
        | Before the process is started, this will be None.
        """

        if self.proc is not None:
            return self.proc.pid

    @property
    def exitcode(self):
        """
        | The child’s exit code.
        | This will be None if the process has not yet terminated.
        | A negative value -N indicates that the child was terminated by signal N.
        """

        if self.proc is not None:
            return self.proc.exitcode


class Context:
    def __init__(self, background: bool = False, uuid: UUID = None, timeout=-1, **kwargs):
        """
        A Context holds information about the current process.

        It is responsible for the creation and management of Processes.

        | Don't share a Context object between Process/Threads.
        | A Context object is not thread-safe.

        :param background: Whether to run Processes under this Context as background tasks.

                           Background tasks keep running even when your main (parent) script finishes execution.

                           Avoids manually calling ``signal.pause()``

        :param uuid: The :py:class:`UUID` object for identifying Context. By default, it is set to :py:class:`None`

                     If uuid is :py:class:`None`,
                     then one will be generated and a new ZProc Server :py:class:`Process` will be spawned.

                     If a :py:class:`UUID` object is provided,
                     then it will connect to an existing ZProc Server :py:class:`Process`, corresponding to that uuid.
        :param timeout: Passed on to :py:meth:`~Context.ping`.
        :param \*\*kwargs: Optional keyword arguments that :py:class:`ZeroProcess` takes.

                           If provided, these will be used while creation of any and all Processes under this Context.

        :ivar background: Passed from constructor.
        :ivar kwargs: Passed from constructor.
        :ivar uuid:  A ``UUID`` for identifying the ZProc Server.
        :ivar state: :py:class:`ZeroState` object. The global state.
        :ivar procs: :py:class:`list` ``[`` :py:class:`ZeroProcess` ``]``.
                     The child Process(s) created under this Context.
        :ivar server_proc: A :py:class:`Process` object for the server.
        """

        self.procs = []
        self.background = background
        self.kwargs = kwargs

        if uuid is None:
            self.uuid = uuid1()

            self.server_proc = Process(target=zproc_server_proc, args=(self.uuid,))
            self.server_proc.start()
        else:
            assert isinstance(uuid, UUID), \
                '"uuid" must be `None`, or an instance of `uuid.UUID`, not `{}`'.format((type(uuid)))

            self.uuid = uuid
            self.server_proc = None

        if not background:
            atexit.register(shutdown)  # kill all the child procs when main script exits

        self.state = ZeroState(self.uuid)

    def process(self, target, **kwargs):
        """
        Produce a child process bound to this context.

        :param target: Passed on to :py:class:`ZeroProcess`'s constructor.
        :param \*\*kwargs: Optional keyword arguments that :py:class:`ZeroProcess` takes.
        """

        _kwargs = self.kwargs.copy()
        _kwargs.update(kwargs)

        proc = ZeroProcess(self.uuid, target, **_kwargs)

        self.procs.append(proc)

        return proc

    def processify(self, **kwargs):
        """
        The decorator version of :py:meth:`~Context.process`

        :param \*\*kwargs: Optional keyword arguments that :py:class:`ZeroProcess` takes.
        :return: A processify decorator, which itself returns a :py:class:`ZeroProcess` instance.
        :rtype: function
        """

        def processify_decorator(func):
            return self.process(func, **kwargs)

        return processify_decorator

    def process_factory(self, *targets: Callable, count=1, **kwargs):
        """
        Produce multiple child process(s) bound to this context.

        :param targets: Passed on to :py:class:`ZeroProcess`'s constructor.
        :param count: The number of processes to spawn for each target
        :param \*\*kwargs: Optional keyword arguments that :py:class:`ZeroProcess` takes.
        :return: spawned processes
        :rtype: :py:class:`list` ``[`` :py:class:`ZeroProcess` ``]``
        """

        procs = []
        for target in targets:
            for _ in range(count):
                procs.append(self.process(target, **kwargs))
        self.procs += procs

        return procs

    def _get_watcher_decorator(self, fn_name, *args, **kwargs):
        def watcher_decorator(fn):
            def watcher_proc(state, props, proc):
                target_params = inspect.signature(fn).parameters.copy()

                if 'kwargs' in target_params:
                    target_kwargs = {'state': state, 'props': props, 'proc': proc}
                else:
                    target_kwargs = {}
                    if 'state' in target_params:
                        target_kwargs['state'] = state
                    if 'props' in target_params:
                        target_kwargs['props'] = props
                    if 'proc' in target_params:
                        target_kwargs['proc'] = proc

                while True:
                    getattr(state, fn_name)(*args, **kwargs)
                    fn(**target_kwargs)

            return self.process(watcher_proc)

        return watcher_decorator

    def call_when_change(self, *keys, live=False):
        """Decorator version of :py:meth:`~ZeroState.get_when_change()`"""

        return self._get_watcher_decorator('get_when_change', *keys, live=live)

    def call_when(self, test_fn, *, live=False):
        """Decorator version of :py:meth:`~ZeroState.get_when()`"""

        return self._get_watcher_decorator('get_when', test_fn, live=live)

    def call_when_equal(self, key, value, *, live=False):
        """Decorator version of :py:meth:`~ZeroState.get_when_equal()`"""

        return self._get_watcher_decorator('get_when_equal', key, value, live=live)

    def call_when_not_equal(self, key, value, *, live=False):
        """Decorator version of :py:meth:`~ZeroState.get_when_not_equal()`"""

        return self._get_watcher_decorator('get_when_not_equal', key, value, live=live)

    def start_all(self):
        """
        Call :py:meth:`~ZeroProcess.start()` on all the child processes of this Context

        Ignore if process is already started
        """

        for proc in self.procs:
            try:
                proc.start()
            except AssertionError:
                pass

    def stop_all(self):
        """Call :py:meth:`~ZeroProcess.stop()` on all the child processes of this Context"""

        for proc in self.procs:
            proc.stop()

    def ping(self, **kwargs):
        """
        Ping the ZProc Server corresponding to this Context

        :param kwargs: Optional keyword arguments that :py:func:`ping` takes.
        :return: Same as :py:func:`ping`
        """

        return ping(self.uuid, **kwargs)

    def close(self):
        """
        Close this context and stop all processes associated with it.

        Once closed, you shouldn't use this Context again.
        """

        self.stop_all()

        if self.server_proc is not None:
            self.server_proc.terminate()

        self.state.close()
