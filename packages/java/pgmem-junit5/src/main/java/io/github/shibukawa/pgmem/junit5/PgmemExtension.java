package io.github.shibukawa.pgmem.junit5;

import java.lang.reflect.Parameter;
import java.nio.file.Path;
import java.time.Duration;
import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Consumer;
import javax.sql.DataSource;
import io.github.shibukawa.pgmem.Fork;
import io.github.shibukawa.pgmem.Pgmem;
import io.github.shibukawa.pgmem.PgmemException;
import io.github.shibukawa.pgmem.Server;
import io.github.shibukawa.pgmem.Snapshot;
import org.junit.jupiter.api.extension.AfterAllCallback;
import org.junit.jupiter.api.extension.AfterEachCallback;
import org.junit.jupiter.api.extension.BeforeAllCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.api.extension.ParameterContext;
import org.junit.jupiter.api.extension.ParameterResolutionException;
import org.junit.jupiter.api.extension.ParameterResolver;

/**
 * Starts one pgmem process per test class, prepares and snapshots each
 * template once, and injects a fresh {@link Fork} (or {@link DataSource},
 * or {@code @PgmemFork String} JDBC URL) into every test method.
 *
 * <pre>{@code
 * @RegisterExtension
 * static PgmemExtension pg = PgmemExtension.builder()
 *         .database("app")
 *         .prepare(template -> migrate(template.jdbcUrl()))
 *         .build();
 *
 * @Test void orders(Fork db) { ... }                       // fresh copy per test
 * @Test void report(@PgmemFork(scope = CLASS) DataSource ds) { ... } // shared by the class
 * }</pre>
 *
 * Register it on a static field (or with {@code @ExtendWith} for the
 * defaults) so the process lives for the whole class. Test methods may run
 * in parallel; each gets its own fork, bounded by {@link Builder#maxForks}.
 */
public final class PgmemExtension
        implements BeforeAllCallback, AfterAllCallback, AfterEachCallback, ParameterResolver {

    private static final ExtensionContext.Namespace NS = ExtensionContext.Namespace.create(PgmemExtension.class);

    /** One template: how to start it and how to prepare it before the snapshot. */
    private static final class Template {
        final String name;
        final String database;
        final String user;
        final List<String> params;
        final Consumer<Server> prepare;
        Server server;
        Snapshot snapshot;

        Template(String name, String database, String user, List<String> params, Consumer<Server> prepare) {
            this.name = name;
            this.database = database;
            this.user = user;
            this.params = params;
            this.prepare = prepare;
        }
    }

    private final Builder cfg;
    private final Map<String, Template> templates = new LinkedHashMap<>();
    private final ForkScope defaultScope;
    private volatile Pgmem pgmem;

    /** Defaults: database {@code postgres}, no preparation, a fresh fork per test method. */
    public PgmemExtension() { this(builder()); }

    private PgmemExtension(Builder b) {
        this.cfg = b;
        this.defaultScope = b.forkScope;
        // The default template (database/prepare on the builder) comes first
        // unless only named templates were configured.
        if (b.templates.isEmpty() || b.prepare != null || b.databaseSet) {
            templates.put("default", new Template("default", b.database, b.user, b.params, b.prepare));
        }
        for (Template t : b.templates) {
            if (templates.put(t.name, t) != null) throw new IllegalArgumentException("duplicate pgmem template " + t.name);
        }
    }

    public static Builder builder() { return new Builder(); }

    // -- accessors for manual use -------------------------------------------

    /** The running process; available after BeforeAll. */
    public Pgmem pgmem() {
        Pgmem p = pgmem;
        if (p == null) throw new IllegalStateException("pgmem not started: register the extension on a static field");
        return p;
    }

    /** The first (or only) template server. */
    public Server template() { return template(firstName()); }

    public Server template(String name) { return get(name).server; }

    /** The snapshot of the first template. */
    public Snapshot snapshot() { return snapshot(firstName()); }

    public Snapshot snapshot(String name) { return get(name).snapshot; }

    /** A fork of the first template that the caller closes. */
    public Fork fork() { return snapshot().fork(); }

    public Fork fork(String name) { return snapshot(name).fork(); }

    private String firstName() { return templates.keySet().iterator().next(); }

    private Template get(String name) {
        Template t = templates.get(name.isEmpty() ? firstName() : name);
        if (t == null) throw new IllegalArgumentException("no pgmem template named " + name + "; have " + templates.keySet());
        if (t.snapshot == null) throw new IllegalStateException("pgmem not started: register the extension on a static field");
        return t;
    }

    // -- lifecycle -----------------------------------------------------------

    @Override public void beforeAll(ExtensionContext context) {
        if (pgmem != null) return; // nested class or re-registration
        Template first = templates.values().iterator().next();
        Pgmem.Builder pb = Pgmem.builder().database(first.database).user(first.user).log(cfg.log).readyTimeout(cfg.readyTimeout);
        for (String p : first.params) {
            int eq = p.indexOf('=');
            pb.param(p.substring(0, eq), p.substring(eq + 1));
        }
        if (cfg.binary != null) pb.binary(cfg.binary);
        Pgmem pg = pb.start();
        try {
            boolean isFirst = true;
            for (Template t : templates.values()) {
                t.server = isFirst ? pg.template() : pg.startServer(t.database, t.user, t.params);
                isFirst = false;
                if (t.prepare != null) t.prepare.accept(t.server);
                t.snapshot = t.server.snapshot(cfg.maxForks);
            }
        } catch (RuntimeException e) {
            pg.close();
            throw e;
        }
        pgmem = pg;
    }

    @Override public void afterAll(ExtensionContext context) {
        // class-scoped forks first, then the process
        ForkSet forks = context.getStore(NS).remove(forkSetKey(context), ForkSet.class);
        if (forks != null) forks.close();
        Pgmem pg = pgmem;
        if (pg != null && context.getRequiredTestClass() == ownerClass(context)) {
            pgmem = null;
            for (Template t : templates.values()) { t.server = null; t.snapshot = null; }
            pg.close();
        }
    }

    /** The outermost class whose BeforeAll started us, so nested classes do not stop the process. */
    private Class<?> ownerClass(ExtensionContext context) {
        ExtensionContext c = context;
        Class<?> owner = context.getRequiredTestClass();
        while (c.getParent().isPresent() && c.getParent().get().getTestClass().isPresent()) {
            c = c.getParent().get();
            owner = c.getRequiredTestClass();
        }
        return owner;
    }

    @Override public void afterEach(ExtensionContext context) {
        ForkSet forks = context.getStore(NS).remove(forkSetKey(context), ForkSet.class);
        if (forks != null) forks.close();
    }

    // -- parameter injection -------------------------------------------------

    @Override public boolean supportsParameter(ParameterContext pc, ExtensionContext ec) {
        Class<?> type = pc.getParameter().getType();
        if (type == Fork.class || type == DataSource.class) return true;
        return type == String.class && pc.isAnnotated(PgmemFork.class);
    }

    @Override public Object resolveParameter(ParameterContext pc, ExtensionContext ec) {
        Parameter p = pc.getParameter();
        PgmemFork ann = pc.findAnnotation(PgmemFork.class).orElse(null);
        String name = ann == null ? "" : ann.value();
        ForkScope scope = ann == null ? defaultScope : ann.scope();
        Fork fork = forkFor(ec, name, scope);
        Class<?> type = p.getType();
        if (type == Fork.class) return fork;
        if (type == DataSource.class) return fork.dataSource();
        if (type == String.class) return fork.jdbcUrl();
        throw new ParameterResolutionException("unsupported parameter " + p);
    }

    /** One fork per (scope context, template); several parameters in one test share it. */
    private Fork forkFor(ExtensionContext ec, String name, ForkScope scope) {
        ExtensionContext scopeCtx = scope == ForkScope.CLASS ? classContext(ec) : ec;
        // Store lookups fall through to parent contexts, so key by this context's id.
        ForkSet set = scopeCtx.getStore(NS).getOrComputeIfAbsent(forkSetKey(scopeCtx), k -> new ForkSet(), ForkSet.class);
        Template t = get(name);
        return set.get(t.name, () -> t.snapshot.fork(cfg.forkTimeout));
    }

    private static String forkSetKey(ExtensionContext ctx) { return "forks:" + ctx.getUniqueId(); }

    private static ExtensionContext classContext(ExtensionContext ec) {
        ExtensionContext c = ec;
        while (c.getTestMethod().isPresent() && c.getParent().isPresent()) c = c.getParent().get();
        return c;
    }

    private static final class ForkSet {
        private final Map<String, Fork> forks = new LinkedHashMap<>();

        synchronized Fork get(String template, java.util.function.Supplier<Fork> create) {
            return forks.computeIfAbsent(template, k -> create.get());
        }

        synchronized void close() {
            PgmemException first = null;
            for (Fork f : forks.values()) {
                try { f.close(); } catch (PgmemException e) { if (first == null) first = e; }
            }
            forks.clear();
            if (first != null) throw first;
        }
    }

    // -- builder -------------------------------------------------------------

    public static final class Builder {
        private String database = "postgres";
        private boolean databaseSet;
        private String user = "postgres";
        private final List<String> params = new ArrayList<>();
        private Consumer<Server> prepare;
        private final List<Template> templates = new ArrayList<>();
        private boolean log;
        private Path binary;
        private int maxForks;
        private ForkScope forkScope = ForkScope.METHOD;
        private Duration forkTimeout;
        private Duration readyTimeout = Duration.ofSeconds(30);

        Builder() {}

        /** Database name of the default template. */
        public Builder database(String database) { this.database = database; this.databaseSet = true; return this; }
        public Builder user(String user) { this.user = user; return this; }
        public Builder param(String name, String value) { params.add(name + "=" + value); return this; }

        /** Runs once against the default template before it is snapshotted (migrations, seed data). */
        public Builder prepare(Consumer<Server> prepare) { this.prepare = prepare; return this; }

        /**
         * Registers an extra template with its own database and preparation. Unnamed
         * parameters use the default template, or the first named one when neither
         * {@link #database} nor {@link #prepare} was set.
         */
        public Builder template(String name, Consumer<Server> prepare) { return template(name, name, prepare); }

        public Builder template(String name, String database, Consumer<Server> prepare) {
            templates.add(new Template(name, database, user, params, prepare));
            return this;
        }

        public Builder log(boolean log) { this.log = log; return this; }
        public Builder binary(Path binary) { this.binary = binary; return this; }

        /** Forks alive at once per template before fork blocks; 0 = available processors. */
        public Builder maxForks(int maxForks) { this.maxForks = maxForks; return this; }

        /** Default scope for parameters without {@link PgmemFork#scope()}. */
        public Builder forkScope(ForkScope scope) { this.forkScope = scope; return this; }

        /** Fail a test instead of waiting forever when every fork slot is taken. */
        public Builder forkTimeout(Duration timeout) { this.forkTimeout = timeout; return this; }

        public Builder readyTimeout(Duration timeout) { this.readyTimeout = timeout; return this; }

        public PgmemExtension build() { return new PgmemExtension(this); }
    }
}
