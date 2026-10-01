package io.github.shibukawa.pgmem.spring;

import io.github.shibukawa.pgmem.junit5.ShadowPgDatabase;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import org.springframework.boot.context.event.ApplicationReadyEvent;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.event.ContextClosedEvent;
import org.springframework.core.annotation.AnnotatedElementUtils;
import org.springframework.core.env.MapPropertySource;
import org.springframework.test.context.ContextConfigurationAttributes;
import org.springframework.test.context.ContextCustomizer;
import org.springframework.test.context.ContextCustomizerFactory;
import org.springframework.test.context.MergedContextConfiguration;

/** Installs pgmem properties before Spring builds the application's connection pool. */
public final class ShadowPgCustomizerFactory implements ContextCustomizerFactory {
    private static final Map<Class<?>, ShadowPgDatabase> DATABASES = new ConcurrentHashMap<>();

    static ShadowPgDatabase database(Class<?> testClass) { return DATABASES.get(testClass); }

    @Override public ContextCustomizer createContextCustomizer(
            Class<?> testClass, List<ContextConfigurationAttributes> configAttributes) {
        ShadowPg annotation = AnnotatedElementUtils.findMergedAnnotation(testClass, ShadowPg.class);
        return annotation == null ? null : new Customizer(testClass, annotation);
    }

    private static final class Customizer implements ContextCustomizer {
        private final Class<?> testClass;
        private final ShadowPg annotation;

        Customizer(Class<?> testClass, ShadowPg annotation) {
            this.testClass = testClass;
            this.annotation = annotation;
        }

        @Override public void customizeContext(ConfigurableApplicationContext context,
                                               MergedContextConfiguration mergedConfig) {
            ShadowPgDatabase database = ShadowPgDatabase.start(testClass, annotation.database(),
                    annotation.schema(), annotation.seed());
            if (DATABASES.putIfAbsent(testClass, database) != null) {
                database.close();
                throw new IllegalStateException("ShadowPg context already active for " + testClass.getName());
            }
            context.addApplicationListener(event -> {
                if (event instanceof ApplicationReadyEvent ready && ready.getApplicationContext() == context) {
                    database.captureApplicationBaseline();
                } else if (event instanceof ContextClosedEvent && event.getSource() == context) {
                    if (DATABASES.remove(testClass, database)) database.close();
                }
            });
            Map<String, Object> properties = new LinkedHashMap<>();
            properties.put("spring.datasource.url", database.fork().jdbcUrl());
            properties.put("spring.datasource.username", database.fork().user());
            properties.put("spring.datasource.password", "");
            properties.put("spring.datasource.driver-class-name", "org.postgresql.Driver");
            if (!annotation.schema().isEmpty()) {
                properties.put("spring.flyway.enabled", false);
                properties.put("spring.liquibase.enabled", false);
            }
            context.getEnvironment().getPropertySources().addFirst(new MapPropertySource("shadowPg", properties));
        }

        @Override public boolean equals(Object other) {
            return other instanceof Customizer that && testClass.equals(that.testClass);
        }

        @Override public int hashCode() { return testClass.hashCode(); }
    }
}
