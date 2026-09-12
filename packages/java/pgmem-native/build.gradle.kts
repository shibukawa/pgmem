description = "The pgmem server binary for one platform, selected by classifier"

val nativeClassifier: String by rootProject.extra
val binaryName: String by rootProject.extra
val nativeDir: Provider<Directory> by rootProject.extra

tasks.jar {
    dependsOn(rootProject.tasks.named("buildBinary"))
    archiveClassifier.set(nativeClassifier)
    from(nativeDir) {
        include(binaryName)
        into("jp/shibu/pgmem/native/$nativeClassifier")
    }
}
// No Java sources: the sources/javadoc jars stay empty but keep Maven Central happy.
