description = "One-annotation Spring Boot tests backed by pgmem"

dependencies {
    api(project(":pgmem-junit5"))
    compileOnly("org.springframework.boot:spring-boot-starter-test:3.5.7")
    testImplementation("org.springframework.boot:spring-boot-starter-test:3.5.7")
    testImplementation("org.springframework.boot:spring-boot-starter-web:3.5.7")
    testImplementation("org.springframework.boot:spring-boot-starter-jdbc:3.5.7")
    testRuntimeOnly("org.postgresql:postgresql:42.7.7")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
}

configure<PublishingExtension> {
    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
            versionMapping { allVariants { fromResolutionResult() } }
        }
    }
}
