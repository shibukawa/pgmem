description = "Micronaut HTTP test server connected to a pgmem fork"

dependencies {
    api(project(":pgmem"))
    api(project(":pgmem-junit5"))
    // The application supplies its Micronaut version and server implementation.
    compileOnly("io.micronaut:micronaut-runtime:4.10.28")
    testImplementation(platform("org.junit:junit-bom:5.14.0"))
    testImplementation("org.junit.jupiter:junit-jupiter")
    testImplementation("io.micronaut:micronaut-runtime:4.10.28")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    testRuntimeOnly("io.micronaut:micronaut-http-server-netty:4.10.28")
    testRuntimeOnly("io.micronaut:micronaut-jackson-databind:4.10.28")
    testImplementation("io.micronaut.sql:micronaut-jdbc-hikari:6.3.4")
    testAnnotationProcessor("io.micronaut:micronaut-inject-java:4.10.28")
    testRuntimeOnly("org.postgresql:postgresql:42.7.7")
}

configure<PublishingExtension> {
    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
            versionMapping { allVariants { fromResolutionResult() } }
        }
    }
}
