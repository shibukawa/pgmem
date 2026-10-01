package io.github.shibukawa.pgmem.junit5;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/** Selects the default database target for parameters resolved by PgmemExtension. */
@Retention(RetentionPolicy.RUNTIME)
@Target({ElementType.TYPE, ElementType.METHOD})
public @interface PgmemTest {
    /** Fresh fork per method by default; false shares a prepared target for read-only tests. */
    boolean fork() default true;
}
