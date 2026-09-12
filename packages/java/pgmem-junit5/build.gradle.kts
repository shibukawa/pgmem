description = "JUnit 5 extension that gives every test its own pgmem fork"

dependencies {
    api(project(":pgmem"))
    api(platform("org.junit:junit-bom:5.14.0"))
    api("org.junit.jupiter:junit-jupiter-api")
    testImplementation("org.junit.jupiter:junit-jupiter")
    testRuntimeOnly("org.junit.platform:junit-platform-launcher")
    testImplementation("org.postgresql:postgresql:42.7.7")
}
