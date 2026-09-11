# Using pgmem from other languages

pgmem speaks the ordinary PostgreSQL wire protocol, so Java, Node.js,
Python or anything else with a PostgreSQL driver can use it. No client
library is needed: a wrapper only has to start the `pgmem` binary as a
subprocess, read the ready line, hand the DSN to the driver, and close
the process's stdin when the test run ends.

## The `pgmem` command

```
pgmem [-port N] [-database NAME] [-user NAME] [-params k=v,k=v] [-log] [-no-stdin]
```

Build it with `go build -ldflags="-s -w" ./cmd/pgmem` (about 37 MB, pure
Go, cross-compiles with `GOOS`/`GOARCH`).

Behaviour that a wrapper can rely on:

- Exactly one JSON line on stdout when the server accepts connections:

  ```json
  {"host":"127.0.0.1","port":54321,"user":"postgres","database":"app","dsn":"postgres://postgres@127.0.0.1:54321/app?sslmode=disable","pid":1234}
  ```

- It listens on `127.0.0.1` only (TCP; a free port unless `-port` is
  given). Loopback TCP is the portable choice: JDBC has no built-in Unix
  domain socket support and Windows drivers differ.
- It exits when its **stdin is closed** or on SIGINT/SIGTERM. Closing
  stdin is the mechanism that works on every platform and also ends the
  server when the parent process dies without cleanup. Pass `-no-stdin`
  to disable that (for example when started from a shell by hand).
- The server log goes to stderr only with `-log`.
- `-params` sets `postgres -c` options (`shared_buffers` defaults to
  32 MB; initdb's 128 MB would add about 80 MB of resident memory).

One process is one single-session PostgreSQL, about 150 MB resident and
ready in about 0.1 s. Tests that need isolated databases can start
several processes.

## Wrappers

The three snippets below do the same thing: spawn, wait for the line,
expose the DSN, stop by closing stdin.

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

// const pg = await startPgmem(["-database", "app"]);
// const client = new Client({ connectionString: pg.dsn }); ...
// pg.stop();
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

# with contextlib.closing(Pgmem("-database", "app")) as pg:
#     conn = psycopg.connect(pg.dsn)  # or asyncpg / SQLAlchemy
```

### Java

```java
public final class Pgmem implements AutoCloseable {
    private final Process proc;
    public final String jdbcUrl;

    public Pgmem(String... args) throws IOException {
        var cmd = new ArrayList<>(List.of(pgmemBinaryPath()));
        cmd.addAll(List.of(args));
        proc = new ProcessBuilder(cmd).redirectError(ProcessBuilder.Redirect.INHERIT).start();
        var line = new BufferedReader(new InputStreamReader(proc.getInputStream())).readLine();
        var info = new ObjectMapper().readTree(line);
        jdbcUrl = "jdbc:postgresql://127.0.0.1:" + info.get("port").asInt()
                + "/" + info.get("database").asText() + "?user=" + info.get("user").asText();
    }

    @Override public void close() throws Exception {
        proc.getOutputStream().close();   // closes the child's stdin
        proc.waitFor(10, TimeUnit.SECONDS);
    }
}
```

Use one connection at a time (the server is a single session): set the
pool size to 1, or keep a single connection per test.

## Shipping the binary

The binary is self-contained, so each ecosystem can use its usual
"platform-specific binary in a package" pattern:

- **npm**: one package per platform (`pgmem-darwin-arm64`, ...) listed as
  `optionalDependencies` of a main package that resolves the right one at
  runtime. This is the esbuild / turbo / biome layout.
- **PyPI**: platform wheels that contain the binary (the `ruff` / `uv`
  layout), or a pure-Python package that downloads the binary from a
  GitHub release on first use into a cache directory.
- **Maven**: one artifact per platform with a classifier
  (`pgmem-<version>-linux-x86_64.jar`), extracted to a temp file on first
  use. `zonky embedded-postgres` uses this layout for real PostgreSQL
  binaries.

Cross-compiling is a plain `GOOS=linux GOARCH=amd64 go build` and the
same for `windows/amd64`, `linux/arm64`, `darwin/arm64`. The default
backend is generated Go with no assembly, so every `GOOS`/`GOARCH` pair
Go supports builds; only darwin/arm64 has been exercised so far.
