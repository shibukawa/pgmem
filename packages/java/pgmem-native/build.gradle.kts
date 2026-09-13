description = "The pgmem server binary, one jar per platform selected by classifier"

val nativeClassifier: String by rootProject.extra
val binaryName: String by rootProject.extra
val nativeDir: Provider<Directory> by rootProject.extra

// Go target -> classifier, as io.github.shibukawa.pgmem.BinaryLocator computes it.
val platforms = mapOf(
    "darwin-amd64" to "darwin-x86_64",
    "darwin-arm64" to "darwin-arm64",
    "linux-amd64" to "linux-x86_64",
    "linux-arm64" to "linux-arm64",
    "windows-amd64" to "windows-x86_64",
    "windows-arm64" to "windows-arm64",
)

// A development build packs the host binary from buildBinary. A release
// build passes -PnativeDist=<dir> with <goos>-<goarch>/pgmem[.exe] for every
// platform (scripts/build-binaries.sh) and gets one classifier jar each.
val nativeDist = findProperty("nativeDist") as String?
val nativeJars: List<TaskProvider<Jar>> = if (nativeDist == null) {
    tasks.jar {
        dependsOn(rootProject.tasks.named("buildBinary"))
        archiveClassifier.set(nativeClassifier)
        from(nativeDir) {
            include(binaryName)
            into("io/github/shibukawa/pgmem/native/$nativeClassifier")
        }
    }
    listOf(tasks.jar)
} else {
    platforms.map { (target, classifier) ->
        val binary = file("$nativeDist/$target/" + if (target.startsWith("windows-")) "pgmem.exe" else "pgmem")
        tasks.register<Jar>("nativeJar-$classifier") {
            archiveClassifier.set(classifier)
            from(binary) { into("io/github/shibukawa/pgmem/native/$classifier") }
            doFirst { check(binary.isFile) { "missing $binary: run scripts/build-binaries.sh" } }
        }
    }
}

// Packaging pom with only the classifier jars: Maven Central wants no main,
// sources or javadoc jar for it.
configure<PublishingExtension> {
    publications {
        create<MavenPublication>("maven") {
            pom { packaging = "pom" }
            nativeJars.forEach { artifact(it) }
        }
    }
}
