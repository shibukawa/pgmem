package jp.shibu.pgmem;

import java.io.IOException;
import java.io.InputStream;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.util.Locale;

/**
 * Finds the pgmem binary: the {@code pgmem.binary} system property, the
 * {@code PGMEM_BINARY} environment variable, the classpath resource shipped by
 * the {@code pgmem-native} artifact for this platform (extracted once into
 * the user's cache directory), then {@code pgmem} on PATH.
 */
public final class BinaryLocator {
    private BinaryLocator() {}

    /** Classifier of the pgmem-native artifact for this JVM's platform, e.g. darwin-arm64. */
    public static String classifier() {
        String os = System.getProperty("os.name").toLowerCase(Locale.ROOT);
        String osName = os.contains("mac") || os.contains("darwin") ? "darwin"
                : os.contains("win") ? "windows" : "linux";
        String arch = System.getProperty("os.arch").toLowerCase(Locale.ROOT);
        String archName = arch.equals("aarch64") || arch.equals("arm64") ? "arm64" : "x86_64";
        return osName + "-" + archName;
    }

    static boolean isWindows() {
        return System.getProperty("os.name").toLowerCase(Locale.ROOT).contains("win");
    }

    public static Path locate(Path explicit) {
        String name = isWindows() ? "pgmem.exe" : "pgmem";
        if (explicit != null) {
            if (!Files.isRegularFile(explicit)) throw new PgmemException("pgmem binary not found: " + explicit);
            return explicit;
        }
        for (String s : new String[] {System.getProperty("pgmem.binary"), System.getenv("PGMEM_BINARY")}) {
            if (s != null && !s.isEmpty()) {
                Path p = Paths.get(s);
                if (Files.isRegularFile(p)) return p;
            }
        }
        Path extracted = extractResource(name);
        if (extracted != null) return extracted;
        String pathEnv = System.getenv("PATH");
        if (pathEnv != null) {
            for (String dir : pathEnv.split(java.io.File.pathSeparator)) {
                Path p = Paths.get(dir, name);
                if (Files.isRegularFile(p) && Files.isExecutable(p)) return p;
            }
        }
        throw new PgmemException("pgmem binary not found: add jp.shibu:pgmem-native with classifier "
                + classifier() + ", or set PGMEM_BINARY to a build of github.com/shibukawa/pgmem/cmd/pgmem");
    }

    private static Path extractResource(String name) {
        String resource = "/jp/shibu/pgmem/native/" + classifier() + "/" + name;
        try (InputStream in = BinaryLocator.class.getResourceAsStream(resource)) {
            if (in == null) return null;
            String version = Pgmem.class.getPackage().getImplementationVersion();
            Path dir = cacheDir().resolve(version == null ? "dev" : version);
            Files.createDirectories(dir);
            Path target = dir.resolve(name);
            if (Files.isRegularFile(target)) return target;
            Path tmp = Files.createTempFile(dir, name, ".part");
            Files.copy(in, tmp, StandardCopyOption.REPLACE_EXISTING);
            if (!isWindows()) tmp.toFile().setExecutable(true, false);
            try {
                Files.move(tmp, target, StandardCopyOption.ATOMIC_MOVE);
            } catch (IOException raced) {
                Files.deleteIfExists(tmp); // another JVM won the race
            }
            return target;
        } catch (IOException e) {
            throw new PgmemException("cannot extract pgmem binary", e);
        }
    }

    private static Path cacheDir() {
        String override = System.getenv("PGMEM_CACHE_DIR");
        if (override != null && !override.isEmpty()) return Paths.get(override);
        String home = System.getProperty("user.home");
        if (home != null && Files.isDirectory(Paths.get(home))) return Paths.get(home, ".cache", "pgmem");
        return Paths.get(System.getProperty("java.io.tmpdir"), "pgmem");
    }
}
