package io.github.shibukawa.pgmem.spring;

import io.github.shibukawa.pgmem.junit5.ShadowPgDatabase;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.AnnotatedElementUtils;
import org.springframework.test.context.TestContext;
import org.springframework.test.context.support.AbstractTestExecutionListener;

/** Reset before Spring's transactional listener opens the test transaction. */
public final class ShadowPgResetListener extends AbstractTestExecutionListener {
    @Override public int getOrder() { return Ordered.HIGHEST_PRECEDENCE; }

    @Override public void beforeTestMethod(TestContext context) {
        Class<?> testClass = context.getTestClass();
        ShadowPg annotation = AnnotatedElementUtils.findMergedAnnotation(testClass, ShadowPg.class);
        if (annotation == null || !annotation.reset()) return;
        ShadowPgDatabase database = ShadowPgCustomizerFactory.database(testClass);
        if (database == null) throw new IllegalStateException("ShadowPg Spring context was not initialized");
        database.reset();
    }
}
