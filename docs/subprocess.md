# Using pgmem from other languages

pgmem speaks the ordinary PostgreSQL wire protocol, so Java, Node.js,
Python or anything else with a PostgreSQL driver can use it. A wrapper
starts the `pgmem` binary as a subprocess, reads the ready line, hands the
DSN to the driver, and talks a small JSON protocol on the process's stdin
and stdout to snapshot the prepared database and fork a copy per test.

Ready-made wrappers live in this repository:

- [`packages/python`](../packages/python/README.md): PyPI package `pgmem`
  with a pytest plugin (`pgmem_dsn` fixture).
- [`packages/java`](../packages/java/README.md): Maven artifacts
  `jp.shibu:pgmem` (client), `jp.shibu:pgmem-junit5` (extension) and
  `jp.shibu:pgmem-native` (binary per platform classifier).
- [`packages/node`](../packages/node/core/README.md): npm package
  `@pgmem/core` with a Vitest/node:test setup entry
  (`@pgmem/core/register`) and a Jest environment.

## The `pgmem` command

```
pgmem [-port N] [-database NAME] [-user NAME] [-params k=v,k=v] [-log] [-no-stdin]
      [-control ADDR] [-wait-timeout D]
```

Build it with `go build -ldflags="-s -w" ./cmd/pgmem` (about 37 MB, pure
Go, cross-compiles with `GOOS`/`GOARCH`).

Behaviour that a wrapper can rely on:

- Exactly one JSON line on stdout when the template server accepts
  connections. The flat fields are what pre-protocol wrappers read; the
  rest negotiates the control protocol:

  ```json
  {"event":"ready","protocol":1,"version":"v0.1.0","pid":1234,
   "server":{"id":"template","host":"127.0.0.1","port":54321,"user":"postgres","database":"app","dsn":"postgres://postgres@127.0.0.1:54321/app?sslmode=disable"},
   "host":"127.0.0.1","port":54321,"user":"postgres","database":"app","dsn":"postgres://postgres@127.0.0.1:54321/app?sslmode=disable"}
  ```

- Every server listens on `127.0.0.1` only (TCP; a free port unless
  `-port` is given for the template). Loopback TCP is the portable choice:
  JDBC has no built-in Unix domain socket support and Windows drivers
  differ.
- It exits when its **stdin is closed** or on SIGINT/SIGTERM, closing
  every fork with it. Closing stdin is the mechanism that works on every
  platform and also ends the server when the parent process dies without
  cleanup. Pass `-no-stdin` to disable that (for example when started from
  a shell by hand); the control protocol is then unavailable.
- The server log goes to stderr only with `-log`. Nothing but protocol
  messages is ever written to stdout.
- `-params` sets `postgres -c` options (`shared_buffers` defaults to
  32 MB; initdb's 128 MB would add about 80 MB of resident memory).
- `-control 127.0.0.1:0` also serves the control protocol on a loopback
  socket (see [Control socket](#control-socket)); the ready line then
  carries `"control":{"addr":"127.0.0.1:54400","token":"<hex>","url":"pgmem-control://<hex>@127.0.0.1:54400"}`.
- `-wait-timeout 2s` is how long a connection may wait behind another
  connection's idle transaction before it is ended with SQLSTATE 55P03
  (negative waits forever).

One process hosts the template, its snapshots and every fork; a fork is a
full single-session PostgreSQL started in about 20 ms because the loaded
engine is shared. Each test process (pytest-xdist worker, forked test
JVM) starts its own `pgmem` process. Node.js test runners, whose workers
are short-lived (Vitest runs every test file in a new process), share one
process started in the global setup and fork through the control socket.

## Control protocol

Newline-delimited JSON, UTF-8. Requests go to stdin, responses come back
on stdout matched by the request's `id` (any JSON value the client
chooses). Requests are handled concurrently and responses may arrive out
of order: `fork` blocks while the pool is full, and the `close` that
frees a slot must not queue behind it. A wrapper therefore needs one
reader thread that dispatches responses by id.

```
→ {"id":1,"op":"snapshot","server":"template","max_forks":4,"timeout_ms":30000}
← {"id":1,"ok":true,"snapshot":"s1"}
→ {"id":2,"op":"fork","snapshot":"s1","timeout_ms":60000}
← {"id":2,"ok":true,"server":{"id":"f2","host":"127.0.0.1","port":54322,"user":"postgres","database":"app","dsn":"postgres://..."}}
→ {"id":3,"op":"close","server":"f2"}
← {"id":3,"ok":true}
→ {"id":4,"op":"start","database":"audit","user":"postgres","params":["log_statement=all"]}
← {"id":4,"ok":true,"server":{"id":"t3",...}}
→ {"id":5,"op":"reset","server":"f2"}
← {"id":5,"ok":true}
→ {"id":6,"op":"shutdown"}
← {"id":6,"ok":true}            (then the process exits 0)
← {"id":9,"ok":false,"error":{"code":"pool_timeout","message":"..."}}
```

| op | fields | result |
|---|---|---|
| `snapshot` | `server` (default `template`), `max_forks` (0 = CPUs), `timeout_ms` | `snapshot` id. Waits for open transactions to end; commit or close every connection first. `busy` after the timeout. |
| `fork` | `snapshot`, `timeout_ms` | `server` endpoint of a new server on a copy of the snapshot. Blocks while `max_forks` forks are alive; `pool_timeout` after the timeout. |
| `close` | `server` or `snapshot` | Stops a fork or template, or rejects further forks from a snapshot. Idempotent. |
| `reset` | `server`, `snapshot` (optional), `timeout_ms` | Puts a server back to a snapshot in place: a fork's own snapshot by default, the named one otherwise (a template needs it). The port and client connections survive; pooled connections keep their named prepared statements and LISTEN registrations. `busy` after the timeout. |
| `start` | `database`, `user`, `params` | `server` endpoint of another template, for suites with several seed sets. |
| `hello` | `token` (control socket) | `protocol` and `version`. The first request on a control socket. |
| `shutdown` | | Closes everything and exits. Closing stdin does the same without a response. Refused on a control socket (`forbidden`). |

Error codes: `unknown_op`, `unknown_id`, `snapshot_closed`, `pool_timeout`,
`busy`, `protocol` (malformed request; `id` is `null` when it could not be
parsed), `unauthorized` and `forbidden` (control socket), `internal`. A
`{"event":"fatal","message":...}` line without an id may precede an
abnormal exit.

### Control socket

With `-control`, the same requests are accepted on a loopback TCP socket,
so processes other than the one that started pgmem (test runner workers)
can fork and reset. A connection starts with
`{"id":1,"op":"hello","token":"<token from the ready line>"}`; any other
first request, or a wrong token, is answered with `unauthorized` and the
connection is closed. Servers and snapshots created on a socket
connection are closed when that connection ends, so a worker that exits
or crashes never keeps fork slots. The process itself still belongs to
whoever holds its stdin.

## Minimal wrappers

The snippets below only start a server and expose its DSN; the packages
above add the protocol client and test-framework integration.

### Node.js

```js
import { spawn } from "node:child_process";
import readline from "node:readline";

export async function startPgmem(args = []) {
  const child = spawn(pgmemBinaryPath(), args, { stdio: ["pipe", "pipe", "inherit"] });
  const rl = readline.createInterface({ input: child.stdout });
  const line = await new Promise((resolve, reject) => {
    rl.once("line", resolve);
    child.once("exit", (code) => reject(new Error(`pgmem exited with ${code}`)));
  });
  const info = JSON.parse(line);
  return { ...info, stop: () => child.stdin.end() };
}
```

### Python

```python
import json, subprocess

class Pgmem:
    def __init__(self, *args):
        self.proc = subprocess.Popen([pgmem_binary_path(), *args],
                                     stdin=subprocess.PIPE, stdout=subprocess.PIPE, text=True)
        self.info = json.loads(self.proc.stdout.readline())
        self.dsn = self.info["dsn"]

    def close(self):
        self.proc.stdin.close()
        self.proc.wait(timeout=10)
```

### Java

```java
var proc = new ProcessBuilder(pgmemBinaryPath(), "-database", "app")
        .redirectError(ProcessBuilder.Redirect.INHERIT).start();
var line = new BufferedReader(new InputStreamReader(proc.getInputStream())).readLine();
// parse the JSON line, build jdbc:postgresql://127.0.0.1:<port>/<database>?user=<user>
proc.getOutputStream().close();   // closes the child's stdin: server exits
```

Connection pools of any size work: each server is a single session, so
connections are serialized at transaction boundaries the way a
transaction-mode pooler does (see "Limits" in the README for what that
means for `SET` and temp tables). `LISTEN`/`NOTIFY` works across
connections.

## Shipping the binary

The binary is self-contained, so each ecosystem uses its usual
"platform-specific binary in a package" pattern, and nothing is downloaded
at run time:

- **PyPI**: one wheel per platform tagged `py3-none-<platform>` with the
  binary at `pgmem/_bin/pgmem` (the `ruff` / `uv` layout).
  `packages/python/hatch_build.py` builds or copies it (`PGMEM_BINARY`,
  `GOOS`/`GOARCH`).
- **Maven**: `jp.shibu:pgmem-native` with one classifier per platform
  (`linux-x86_64`, `linux-arm64`, `darwin-x86_64`, `darwin-arm64`,
  `windows-x86_64`), extracted to `~/.cache/pgmem/<version>/` on first
  use. `zonky embedded-postgres` uses this layout for real PostgreSQL
  binaries. `./gradlew -Pgoos=linux -Pgoarch=amd64 build` cross-compiles.
- **npm**: `@pgmem/core` lists one `@pgmem/<platform>` package per
  platform as `optionalDependencies` (the esbuild layout, with no install
  script, since pnpm, Yarn, Bun and npm 11 hold install scripts by
  default). `scripts/build-npm.sh` cross-compiles into
  `packages/node/platforms/<platform>/bin/`.

Cross-compiling is a plain `GOOS=linux GOARCH=amd64 go build` and the
same for `windows/amd64`, `linux/arm64`, `darwin/arm64`. The default
backend is generated Go with no assembly, so every `GOOS`/`GOARCH` pair
Go supports builds; only darwin/arm64 has been exercised so far.
