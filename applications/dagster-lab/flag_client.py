"""Thread-backed flag client — stand-in for the Amplitude SDK.

The real SDK keeps an internal background poller/connection thread. We model
exactly that: a daemon thread holds the live flag state and answers fetch_v2()
over an in-process queue. The thread is the 'guts' that does NOT survive
os.fork() — so a client built in the Celery parent and reused in a prefork
ForkPoolWorker child returns EMPTY, silently (no exception).
"""

import os
import queue
import threading


class FlagClient:
    def __init__(self) -> None:
        self._owner_pid = os.getpid()
        self._req: queue.Queue = queue.Queue()
        self._resp: queue.Queue = queue.Queue()
        self._ready = threading.Event()
        self._thread = threading.Thread(target=self._serve, daemon=True)
        self._thread.start()
        self._ready.wait(timeout=1.0)

    def _serve(self) -> None:
        # Live state lives in the thread; only runs in the process that started it.
        flags = {"shoppayinstallment_enable_20260130": "on"}
        self._ready.set()
        while True:
            try:
                self._req.get(timeout=0.1)
            except queue.Empty:
                continue
            self._resp.put(dict(flags))

    def fetch_v2(self) -> dict:
        """Ask the bg thread. Post-fork the thread is gone -> nobody answers ->
        empty, no exception. Mirrors all_flag_keys=None in the real bug."""
        self._req.put("fetch")
        try:
            return self._resp.get(timeout=0.5)
        except queue.Empty:
            return {}

    def thread_alive_here(self) -> bool:
        return self._thread.is_alive()
