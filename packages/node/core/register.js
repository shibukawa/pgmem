// Setup entry: before the test file is imported, give this process its own
// fork of PGMEM_SNAPSHOT and write the fork's URL to DATABASE_URL (or to the
// comma-separated names in PGMEM_ENV). Use it as a Vitest setup file, with
// node --test --import, or with bun test --isolate --preload.
import { useFork } from "./index.js";

await useFork();
