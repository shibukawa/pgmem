"""Client for the pgmem control protocol (JSON lines over stdin/stdout)."""

import json
import os
import subprocess
import sys
import threading
from typing import Dict, List, Optional

from ._binary import find_binary

PROTOCOL = 1


class PgmemError(Exception):
    """Base class for pgmem errors."""


class ProtocolError(PgmemError):
    """The server answered a request with an error."""

    def __init__(self, code, message):
        super().__init__(f"{code}: {message}")
        self.code = code
        self.message = message


class ServerExited(PgmemError):
    """The pgmem process ended while a request was pending."""


class _Waiter:
    __slots__ = ("event", "response")

    def __init__(self):
        self.event = threading.Event()
        self.response = None


class Pgmem:
    """One pgmem process: the default template, extra templates, snapshots, forks.

    Use :func:`start` to create it. Closing it (or the process exiting)
    releases everything the process owns.
    """

    def __init__(self, proc: subprocess.Popen, ready: dict):
        self._proc = proc
        self._lock = threading.Lock()  # guards stdin writes and the waiter table
        self._seq = 0
        self._waiters: Dict[int, _Waiter] = {}
        self._exited = False
        self._closed = False
        self.pid = ready["pid"]
        self.version = ready.get("version", "")
        self.template = Server(self, ready["server"])
        self._reader = threading.Thread(target=self._read_loop, name="pgmem-reader", daemon=True)
        self._reader.start()

    # -- lifecycle ---------------------------------------------------------

    def start_server(self, database="postgres", user="postgres", params: Optional[List[str]] = None) -> "Server":
        """Start another template server in the same process (op ``start``)."""
        res = self._request("start", database=database, user=user, params=list(params or []))
        return Server(self, res["server"])

    def close(self):
        """Shut the process down. Idempotent."""
        if self._closed:
            return
        self._closed = True
        try:
            if not self._exited:
                self._request("shutdown")
        except PgmemError:
            pass
        finally:
            try:
                self._proc.stdin.close()
            except OSError:
                pass
            try:
                self._proc.wait(timeout=10)
            except subprocess.TimeoutExpired:
                self._proc.kill()
                self._proc.wait()

    def __enter__(self):
        return self

    def __exit__(self, *exc):
        self.close()

    # -- protocol ----------------------------------------------------------

    def _request(self, op, **fields):
        with self._lock:
            if self._exited:
                raise ServerExited("pgmem process has exited")
            self._seq += 1
            rid = self._seq
            w = _Waiter()
            self._waiters[rid] = w
            line = json.dumps({"id": rid, "op": op, **fields}) + "\n"
            try:
                self._proc.stdin.write(line)
                self._proc.stdin.flush()
            except (OSError, ValueError) as e:
                del self._waiters[rid]
                raise ServerExited(f"cannot write to pgmem process: {e}") from e
        w.event.wait()
        res = w.response
        if res is None:
            raise ServerExited("pgmem process exited before answering")
        if not res.get("ok"):
            err = res.get("error") or {}
            raise ProtocolError(err.get("code", "internal"), err.get("message", "unknown error"))
        return res

    def _read_loop(self):
        out = self._proc.stdout
        try:
            for line in out:
                line = line.strip()
                if not line:
                    continue
                try:
                    msg = json.loads(line)
                except ValueError:
                    continue
                if "event" in msg and msg.get("id") is None:
                    if msg["event"] == "fatal":
                        sys.stderr.write(f"pgmem: fatal: {msg.get('message')}\n")
                    continue
                rid = msg.get("id")
                with self._lock:
                    w = self._waiters.pop(rid, None)
                if w is not None:
                    w.response = msg
                    w.event.set()
        finally:
            with self._lock:
                self._exited = True
                pending = list(self._waiters.values())
                self._waiters.clear()
            for w in pending:
                w.event.set()


class Server:
    """A listening PostgreSQL server inside the pgmem process."""

    def __init__(self, pg: Pgmem, endpoint: dict):
        self._pg = pg
        self.id = endpoint["id"]
        self.host = endpoint["host"]
        self.port = endpoint["port"]
        self.user = endpoint["user"]
        self.database = endpoint["database"]
        self.dsn = endpoint["dsn"]
        self._closed = False

    def snapshot(self, max_forks: Optional[int] = None, timeout: Optional[float] = 30.0) -> "Snapshot":
        """Checkpoint and copy this server's state; forks start from the copy.

        The snapshot waits for open transactions to end, so commit or close
        every connection first. After ``timeout`` seconds (None = forever)
        it fails with a ``ProtocolError`` whose code is ``busy``.

        ``max_forks`` caps the forks alive at once; None derives it from the
        server's memory (a quarter of its limit divided by a fork's cost).
        """
        fields = {"server": self.id}
        if max_forks:
            fields["max_forks"] = max_forks
        if timeout is not None:
            fields["timeout_ms"] = int(timeout * 1000)
        res = self._pg._request("snapshot", **fields)
        return Snapshot(self._pg, res["snapshot"], self)

    def close(self):
        if self._closed:
            return
        self._closed = True
        self._pg._request("close", server=self.id)

    def __enter__(self):
        return self

    def __exit__(self, *exc):
        self.close()

    def __repr__(self):
        return f"<pgmem.{type(self).__name__} {self.id} {self.dsn}>"


class Fork(Server):
    """A server started from a snapshot; close it to free its pool slot."""


class Snapshot:
    """A frozen copy of a server's state."""

    def __init__(self, pg: Pgmem, sid: str, origin: Server):
        self._pg = pg
        self.id = sid
        self.origin = origin
        self._closed = False

    def fork(self, timeout: Optional[float] = None) -> Fork:
        """Start a fresh server on a copy of the snapshot.

        Blocks while ``max_forks`` forks are alive; ``timeout`` (seconds)
        turns that wait into a ``ProtocolError`` with code ``pool_timeout``.
        """
        fields = {"snapshot": self.id}
        if timeout is not None:
            fields["timeout_ms"] = int(timeout * 1000)
        res = self._pg._request("fork", **fields)
        return Fork(self._pg, res["server"])

    def close(self):
        if self._closed:
            return
        self._closed = True
        self._pg._request("close", snapshot=self.id)

    def __enter__(self):
        return self

    def __exit__(self, *exc):
        self.close()


def start(database="postgres", user="postgres", params: Optional[List[str]] = None,
          log=False, binary: Optional[str] = None, timeout: float = 30.0) -> Pgmem:
    """Spawn a pgmem process and wait until its template server is ready.

    ``params`` are ``postgres -c`` settings such as ``["log_statement=all"]``.
    ``log`` passes the server log through to stderr. ``binary`` overrides
    the binary lookup (see :func:`find_binary`).
    """
    args = [find_binary(binary), "-database", database, "-user", user]
    if params:
        args += ["-params", ",".join(params)]
    if log:
        args.append("-log")
    proc = subprocess.Popen(
        args, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=None,
        text=True, encoding="utf-8", bufsize=1,
    )
    ready = _read_ready(proc, timeout)
    return Pgmem(proc, ready)


def _read_ready(proc, timeout):
    box = {}

    def reader():
        box["line"] = proc.stdout.readline()

    t = threading.Thread(target=reader, daemon=True)
    t.start()
    t.join(timeout)
    if t.is_alive():
        proc.kill()
        raise PgmemError(f"pgmem did not become ready within {timeout}s")
    line = box.get("line") or ""
    if not line:
        code = proc.wait()
        raise PgmemError(f"pgmem exited with status {code} before becoming ready")
    ready = json.loads(line)
    if ready.get("event") != "ready" or "server" not in ready:
        proc.kill()
        raise PgmemError(f"unexpected first line from pgmem: {line.strip()}")
    if ready.get("protocol") != PROTOCOL:
        proc.kill()
        raise PgmemError(f"pgmem binary speaks protocol {ready.get('protocol')}, this package needs {PROTOCOL}")
    return ready
