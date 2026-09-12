import java.util.Locale

// Target platform for the Go binary: -Pgoos=linux -Pgoarch=arm64 cross-builds;
// the default is the build host.
fun hostGoos(): String {
    val os = System.getProperty("os.name").lowercase(Locale.ROOT)
    return when {
        os.contains("mac") || os.contains("darwin") -> "darwin"
        os.contains("win") -> "windows"
        else -> "linux"
    }
}
fun hostGoarch(): String = when (System.getProperty("os.arch").lowercase(Locale.ROOT)) {
    "aarch64", "arm64" -> "arm64"
    else -> "amd64"
}
val goos: String = (findProperty("goos") as String?) ?: hostGoos()
val goarch: String = (findProperty("goarch") as String?) ?: hostGoarch()
// Classifier as jp.shibu.pgmem.BinaryLocator computes it from os.name/os.arch.
val nativeClassifier: String = "$goos-" + (if (goarch == "amd64") "x86_64" else "arm64")
val binaryName: String = if (goos == "windows") "pgmem.exe" else "pgmem"
val nativeDir: Provider<Directory> = layout.buildDirectory.dir("native/$nativeClassifier")

extra["nativeClassifier"] = nativeClassifier
extra["binaryName"] = binaryName
extra["nativeDir"] = nativeDir

// One go build shared by the native jar and the test tasks. PGMEM_BINARY
// skips the build and copies a prebuilt binary instead.
val buildBinary by tasks.registering(Exec::class) {
    val repoRoot = rootDir.parentFile.parentFile
    val prebuilt = System.getenv("PGMEM_BINARY")
    inputs.dir(File(repoRoot, "cmd/pgmem"))
    outputs.file(nativeDir.map { it.file(binaryName) })
    workingDir = repoRoot
    environment("GOOS", goos)
    environment("GOARCH", goarch)
    environment("CGO_ENABLED", "0")
    if (prebuilt != null) {
        commandLine("cp", prebuilt, nativeDir.get().file(binaryName).asFile.absolutePath)
    } else {
        commandLine(
            "go", "build", "-trimpath", "-ldflags=-s -w",
            "-o", nativeDir.get().file(binaryName).asFile.absolutePath, "./cmd/pgmem",
        )
    }
    doFirst { nativeDir.get().asFile.mkdirs() }
}

allprojects {
    group = "jp.shibu"
    version = "0.1.0"
    repositories { mavenCentral() }
}

subprojects {
    apply(plugin = "java-library")
    apply(plugin = "maven-publish")

    extensions.configure<JavaPluginExtension> {
        withSourcesJar()
        withJavadocJar()
    }
    tasks.withType<JavaCompile>().configureEach {
        options.release.set(17)
        options.encoding = "UTF-8"
    }
    tasks.withType<Jar>().configureEach {
        manifest { attributes("Implementation-Version" to project.version, "Implementation-Title" to project.name) }
    }
    tasks.withType<Javadoc>().configureEach {
        (options as StandardJavadocDocletOptions).addBooleanOption("Xdoclint:none", true)
    }
    tasks.withType<Test>().configureEach {
        useJUnitPlatform()
        dependsOn(buildBinary)
        systemProperty("pgmem.binary", nativeDir.get().file(binaryName).asFile.absolutePath)
        testLogging { events("passed", "failed", "skipped"); showStandardStreams = false }
    }
    extensions.configure<PublishingExtension> {
        publications { create<MavenPublication>("maven") { from(components["java"]) } }
    }
}
