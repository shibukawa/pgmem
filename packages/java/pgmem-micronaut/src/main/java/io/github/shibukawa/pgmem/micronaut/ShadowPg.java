package io.github.shibukawa.pgmem.micronaut;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import org.junit.jupiter.api.extension.ExtendWith;

/** Start and inject a Micronaut HTTP app connected to a prepared pgmem fork. */
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@ExtendWith(ShadowPgExtension.class)
public @interface ShadowPg {
    String database() default "postgres";
    String schema() default "";
    String seed() default "";
    boolean reset() default true;
    /** Extra Micronaut properties in {@code key=value} form. */
    String[] properties() default {};
}
