package io.github.shibukawa.pgmem;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.nio.charset.StandardCharsets;
import java.nio.file.Path;
import java.time.Duration;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicLong;

/**
 * One pgmem process: a real PostgreSQL that runs entirely in memory.
 *
 * <pre>{@code
 * try (Pgmem pg = Pgmem.builder().database("app").start()) {
 *     migrate(pg.template().jdbcUrl());
 *     Snapshot snap = pg.template().snapshot();
 *     try (Fork fork = snap.fork()) {            // private copy of the prepared database
 *         DriverManager.getConnection(fork.jdbcUrl()) ...
 *     }
 * }
 * }</pre>
 *
 * Closing the Pgmem, or the JVM exiting, ends the process and everything in it.
 */
public final class Pgmem implements AutoCloseable {
    /** Protocol version this client speaks; the ready line must match. */
    public static final int PROTOCOL = 1;

    private final Process process;
    private final BufferedReader stdout;
    private final BufferedWriter stdin;
    private final Thread reader;
    private final AtomicLong seq = new AtomicLong();
    private final ConcurrentHashMap<Long, CompletableFuture<Map<String, Object>>> pending = new ConcurrentHashMap<>();
    private final Server template;
    private final long pid;
    private final String version;
    private volatile boolean exited;
    private volatile boolean closed;

    private Pgmem(Process process, BufferedReader stdout, Map<String, Object> ready) {
        this.process = process;
        this.stdout = stdout;
        this.stdin = new BufferedWriter(new OutputStreamWriter(process.getOutputStream(), StandardCharsets.UTF_8));
        this.pid = ((Number) ready.get("pid")).longValue();
        this.version = String.valueOf(ready.getOrDefault("version", ""));
        this.template = new Server(this, endpoint(ready));
        this.reader = new Thread(this::readLoop, "pgmem-reader");
        this.reader.setDaemon(true);
        this.reader.start();
    }

    public static Builder builder() { return new Builder(); }

    /** Starts a process with default options (database and user {@code postgres}). */
    public static Pgmem start() { return builder().start(); }

    /** The server started by the process itself. */
    public Server template() { return template; }

    public long pid() { return pid; }

    /** Version string of the binary. */
    public String version() { return version; }

    /** Starts another template server, for suites with several seed sets. */
    public Server startServer(String database) { return startServer(database, "postgres", List.of()); }

    public Server startServer(String database, String user, List<String> params) {
        Map<String, Object> req = new LinkedHashMap<>();
        req.put("database", database);
        req.put("user", user);
        req.put("params", new ArrayList<>(params));
        return new Server(this, endpoint(request("start", req)));
    }

    /** Shuts the process down. Idempotent. */
    @Override public void close() {
        if (closed) return;
        closed = true;
        try {
            if (!exited) request("shutdown", Map.of());
        } catch (PgmemException ignored) {
            // exited meanwhile
        }
        try {
            stdin.close();
        } catch (IOException ignored) {
            // already gone
        }
        try {
            if (!process.waitFor(10, TimeUnit.SECONDS)) {
                process.destroyForcibly();
                process.waitFor();
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            process.destroyForcibly();
        }
    }

    // -- protocol ------------------------------------------------------------

    @SuppressWarnings("unchecked")
    static Endpoint endpoint(Map<String, Object> res) {
        return new Endpoint((Map<String, Object>) res.get("server"));
    }

    Map<String, Object> request(String op, Map<String, Object> fields) {
        if (exited) throw new ServerExitedException("pgmem process has exited");
        long id = seq.incrementAndGet();
        Map<String, Object> msg = new LinkedHashMap<>();
        msg.put("id", id);
        msg.put("op", op);
        msg.putAll(fields);
        CompletableFuture<Map<String, Object>> f = new CompletableFuture<>();
        pending.put(id, f);
        synchronized (stdin) {
            try {
                stdin.write(Json.write(msg));
                stdin.write('\n');
                stdin.flush();
            } catch (IOException e) {
                pending.remove(id);
                throw new ServerExitedException("cannot write to pgmem process", e);
            }
        }
        if (exited) failPending();
        Map<String, Object> res;
        try {
            res = f.get();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            pending.remove(id);
            throw new PgmemException("interrupted while waiting for " + op, e);
        } catch (ExecutionException e) {
            throw e.getCause() instanceof PgmemException ? (PgmemException) e.getCause()
                    : new PgmemException(op + " failed", e.getCause());
        }
        if (!Boolean.TRUE.equals(res.get("ok"))) {
            @SuppressWarnings("unchecked")
            Map<String, Object> err = (Map<String, Object>) res.get("error");
            String code = err == null ? "internal" : String.valueOf(err.get("code"));
            String message = err == null ? "unknown error" : String.valueOf(err.get("message"));
            throw new ProtocolException(code, message);
        }
        return res;
    }

    private void readLoop() {
        try (BufferedReader out = stdout) {
            String line;
            while ((line = out.readLine()) != null) {
                line = line.trim();
                if (line.isEmpty()) continue;
                Map<String, Object> msg;
                try {
                    msg = Json.parseObject(line);
                } catch (IllegalArgumentException e) {
                    continue;
                }
                Object id = msg.get("id");
                if (id == null) {
                    if ("fatal".equals(msg.get("event"))) {
                        System.err.println("pgmem: fatal: " + msg.get("message"));
                    }
                    continue;
                }
                CompletableFuture<Map<String, Object>> f = pending.remove(((Number) id).longValue());
                if (f != null) f.complete(msg);
            }
        } catch (IOException ignored) {
            // process gone
        } finally {
            exited = true;
            failPending();
        }
    }

    private void failPending() {
        for (Long id : new ArrayList<>(pending.keySet())) {
            CompletableFuture<Map<String, Object>> f = pending.remove(id);
            if (f != null) f.completeExceptionally(new ServerExitedException("pgmem process exited before answering"));
        }
    }

    // -- builder -------------------------------------------------------------

    public static final class Builder {
        private String database = "postgres";
        private String user = "postgres";
        private final List<String> params = new ArrayList<>();
        private boolean log;
        private Path binary;
        private Duration readyTimeout = Duration.ofSeconds(30);

        Builder() {}

        public Builder database(String database) { this.database = database; return this; }
        public Builder user(String user) { this.user = user; return this; }

        /** Adds a {@code postgres -c} setting, e.g. {@code param("log_statement", "all")}. */
        public Builder param(String name, String value) { params.add(name + "=" + value); return this; }

        /** Passes the server log through to stderr. */
        public Builder log(boolean log) { this.log = log; return this; }

        /** Overrides the binary lookup (see {@link BinaryLocator}). */
        public Builder binary(Path binary) { this.binary = binary; return this; }

        public Builder readyTimeout(Duration timeout) { this.readyTimeout = timeout; return this; }

        public Pgmem start() {
            List<String> cmd = new ArrayList<>();
            cmd.add(BinaryLocator.locate(binary).toString());
            cmd.add("-database"); cmd.add(database);
            cmd.add("-user"); cmd.add(user);
            if (!params.isEmpty()) { cmd.add("-params"); cmd.add(String.join(",", params)); }
            if (log) cmd.add("-log");
            Process p;
            try {
                p = new ProcessBuilder(cmd).redirectError(ProcessBuilder.Redirect.INHERIT).start();
            } catch (IOException e) {
                throw new PgmemException("cannot start " + cmd.get(0), e);
            }
            BufferedReader stdout = new BufferedReader(new InputStreamReader(p.getInputStream(), StandardCharsets.UTF_8));
            Map<String, Object> ready = readReady(p, stdout, readyTimeout);
            return new Pgmem(p, stdout, ready);
        }

        private static Map<String, Object> readReady(Process p, BufferedReader stdout, Duration timeout) {
            CompletableFuture<String> first = CompletableFuture.supplyAsync(() -> {
                try {
                    return stdout.readLine();
                } catch (IOException e) {
                    return null;
                }
            });
            String line;
            try {
                line = first.get(timeout.toMillis(), TimeUnit.MILLISECONDS);
            } catch (TimeoutException e) {
                p.destroyForcibly();
                throw new PgmemException("pgmem did not become ready within " + timeout);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                p.destroyForcibly();
                throw new PgmemException("interrupted while starting pgmem", e);
            } catch (ExecutionException e) {
                p.destroyForcibly();
                throw new PgmemException("cannot read from pgmem", e.getCause());
            }
            if (line == null) {
                int code;
                try { code = p.waitFor(); } catch (InterruptedException e) { Thread.currentThread().interrupt(); code = -1; }
                throw new PgmemException("pgmem exited with status " + code + " before becoming ready");
            }
            Map<String, Object> ready = Json.parseObject(line);
            if (!"ready".equals(ready.get("event")) || !(ready.get("server") instanceof Map)) {
                p.destroyForcibly();
                throw new PgmemException("unexpected first line from pgmem: " + line);
            }
            Object proto = ready.get("protocol");
            if (!(proto instanceof Number) || ((Number) proto).intValue() != PROTOCOL) {
                p.destroyForcibly();
                throw new PgmemException("pgmem binary speaks protocol " + proto + ", this client needs " + PROTOCOL);
            }
            return ready;
        }
    }
}
