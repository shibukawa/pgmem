package jp.shibu.pgmem.junit5;

/** How long an injected fork lives. */
public enum ForkScope {
    /** A fresh fork for every test method (default). */
    METHOD,
    /** One fork shared by all test methods of the class; for read-only tests. */
    CLASS
}
