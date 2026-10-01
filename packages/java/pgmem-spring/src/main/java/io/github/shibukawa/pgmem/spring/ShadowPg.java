package io.github.shibukawa.pgmem.spring;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.core.annotation.AliasFor;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.context.ContextCustomizerFactories;
import org.springframework.test.context.TestExecutionListeners;

/** Start a Spring Boot test with its normal DataSource connected to a prepared pgmem fork. */
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@DirtiesContext(classMode = DirtiesContext.ClassMode.AFTER_CLASS)
@ContextCustomizerFactories(ShadowPgCustomizerFactory.class)
@TestExecutionListeners(listeners = ShadowPgResetListener.class,
        mergeMode = TestExecutionListeners.MergeMode.MERGE_WITH_DEFAULTS)
public @interface ShadowPg {
    /** Spring Boot application configuration, when it cannot be discovered automatically. */
    @AliasFor(annotation = SpringBootTest.class, attribute = "classes")
    Class<?>[] classes() default {};

    /** Extra Spring Boot properties for this test application. */
    @AliasFor(annotation = SpringBootTest.class, attribute = "properties")
    String[] properties() default {};

    String database() default "postgres";

    /** Classpath SQL resource run before snapshotting the template. */
    String schema() default "";

    /** Classpath SQL resource run after schema and before snapshotting. */
    String seed() default "";

    /** Restore the prepared dataset before each test method. */
    boolean reset() default true;
}
