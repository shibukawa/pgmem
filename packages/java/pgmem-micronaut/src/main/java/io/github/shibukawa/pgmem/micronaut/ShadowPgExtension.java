package io.github.shibukawa.pgmem.micronaut;

import io.github.shibukawa.pgmem.junit5.ShadowPgDatabase;
import java.util.LinkedHashMap;
import java.util.Map;
import org.junit.jupiter.api.extension.AfterAllCallback;
import org.junit.jupiter.api.extension.BeforeEachCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.api.extension.TestInstancePostProcessor;

/** Owns the Micronaut context and pgmem fork for one JUnit test class. */
public final class ShadowPgExtension
        implements TestInstancePostProcessor, BeforeEachCallback, AfterAllCallback {
    private static final ExtensionContext.Namespace NS = ExtensionContext.Namespace.create(ShadowPgExtension.class);

    private static final class Runtime implements AutoCloseable {
        final ShadowPgDatabase database;
        final MicronautTestApp app;

        Runtime(Class<?> testClass, ShadowPg annotation) {
            database = ShadowPgDatabase.start(testClass, annotation.database(), annotation.schema(), annotation.seed());
            MicronautTestApp started = null;
            try {
                Map<String, Object> properties = new LinkedHashMap<>();
                if (annotation.schema().isEmpty()) properties.put("flyway.datasources.default.enabled", true);
                for (String entry : annotation.properties()) {
                    int separator = entry.indexOf('=');
                    if (separator < 1) throw new IllegalArgumentException("ShadowPg property must be key=value: " + entry);
                    properties.put(entry.substring(0, separator), entry.substring(separator + 1));
                }
                started = MicronautTestApp.start(database.fork(), properties);
                database.captureApplicationBaseline();
                app = started;
            }
            catch (RuntimeException e) {
                if (started != null) started.close();
                database.close();
                throw e;
            }
        }

        @Override public void close() {
            try { app.close(); }
            finally { database.close(); }
        }
    }

    private static Runtime runtime(ExtensionContext context) {
        Class<?> testClass = context.getRequiredTestClass();
        ShadowPg annotation = testClass.getAnnotation(ShadowPg.class);
        if (annotation == null) throw new IllegalStateException("@ShadowPg is required");
        return context.getStore(NS).getOrComputeIfAbsent("runtime", key -> new Runtime(testClass, annotation), Runtime.class);
    }

    @Override public void postProcessTestInstance(Object testInstance, ExtensionContext context) {
        runtime(context).app.server().getApplicationContext().inject(testInstance);
    }

    @Override public void beforeEach(ExtensionContext context) {
        if (context.getRequiredTestClass().getAnnotation(ShadowPg.class).reset()) runtime(context).database.reset();
    }

    @Override public void afterAll(ExtensionContext context) {
        Runtime current = context.getStore(NS).remove("runtime", Runtime.class);
        if (current != null) current.close();
    }
}
