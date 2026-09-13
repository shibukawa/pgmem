package io.github.shibukawa.pgmem.junit5;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * Selects which template a parameter's fork comes from and how long it lives.
 * Optional on {@code Fork} and {@code DataSource} parameters; required on
 * {@code String} parameters so ordinary strings are not resolved by mistake.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.PARAMETER)
public @interface PgmemFork {
    /** Template name as registered with {@code PgmemExtension.Builder#template}; empty = the first. */
    String value() default "";

    ForkScope scope() default ForkScope.METHOD;
}
