description = "In-memory PostgreSQL for tests: client for the pgmem server process"

dependencies {
    testImplementation(platform("org.junit:junit-bom:5.14.0"))
    testImplementation("org.junit.jupiter:junit-jupiter")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    testImplementation("org.postgresql:postgresql:42.7.7")
}

configure<PublishingExtension> {
    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
            versionMapping { allVariants { fromResolutionResult() } }
        }
    }
}
