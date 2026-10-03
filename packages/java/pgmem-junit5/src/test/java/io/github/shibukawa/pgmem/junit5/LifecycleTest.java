package io.github.shibukawa.pgmem.junit5;

import static org.junit.jupiter.api.Assertions.*;
import java.lang.reflect.Proxy;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;
import java.util.function.Function;
import io.github.shibukawa.pgmem.Fork;
import io.github.shibukawa.pgmem.Pgmem;
import io.github.shibukawa.pgmem.ProtocolException;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtensionContext;

class LifecycleTest {
    @Test void cleanupRetainsOriginalAndAttemptsEveryAction() {
        AssertionError original = new AssertionError("test failed");
        RuntimeException first = new RuntimeException("first cleanup failed");
        RuntimeException second = new RuntimeException("second cleanup failed");
        int[] attempted = {0};
        AssertionError caught = assertThrows(AssertionError.class, () -> PgmemExtension.cleanup(original,
                () -> { attempted[0]++; throw first; },
                () -> { attempted[0]++; throw second; },
                () -> attempted[0]++));
        assertSame(original, caught);
        assertEquals(3, attempted[0]);
        assertSame(first, caught.getSuppressed()[0]);
        assertSame(second, first.getSuppressed()[0]);
    }

    @Test void afterAllStopsProcessEvenWhenClassForkCloseFails() throws Exception {
        ExtensionContext context = context();
        PgmemExtension extension = PgmemExtension.builder().build();
        extension.beforeAll(context);
        Pgmem process = extension.pgmem();
        try {
            var acquire = PgmemExtension.class.getDeclaredMethod("forkFor", ExtensionContext.class, String.class, ForkScope.class);
            acquire.setAccessible(true);
            Fork fork = (Fork) acquire.invoke(extension, context, "", ForkScope.CLASS);
            // Make this handle's close fail while its owned process and actual fork stay live.
            var id = fork.endpoint().getClass().getDeclaredField("id");
            id.setAccessible(true);
            id.set(fork.endpoint(), "");
            assertThrows(ProtocolException.class, () -> extension.afterAll(context));
            assertFalse(ProcessHandle.of(process.pid()).map(ProcessHandle::isAlive).orElse(false));
        } finally { process.close(); }
    }

    @SuppressWarnings("unchecked")
    private static ExtensionContext context() {
        Map<Object, Object> values = new HashMap<>();
        ExtensionContext.Store store = (ExtensionContext.Store) Proxy.newProxyInstance(
                LifecycleTest.class.getClassLoader(), new Class<?>[] { ExtensionContext.Store.class }, (proxy, method, args) -> {
                    switch (method.getName()) {
                        case "getOrComputeIfAbsent": return values.computeIfAbsent(args[0], (Function<Object, Object>) args[1]);
                        case "remove": return values.remove(args[0]);
                        case "get": return values.get(args[0]);
                        default: throw new UnsupportedOperationException(method.getName());
                    }
                });
        return (ExtensionContext) Proxy.newProxyInstance(LifecycleTest.class.getClassLoader(), new Class<?>[] { ExtensionContext.class },
                (proxy, method, args) -> {
                    switch (method.getName()) {
                        case "getStore": return store;
                        case "getRequiredTestClass": return LifecycleTest.class;
                        case "getUniqueId": return "lifecycle-test";
                        case "getParent": case "getTestMethod": case "getExecutionException": return Optional.empty();
                        default: throw new UnsupportedOperationException(method.getName());
                    }
                });
    }
}
