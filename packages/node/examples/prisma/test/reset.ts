import { beforeEach } from "vitest";
import { currentFork } from "@pgmem/core";

// Every test starts from the migrated, empty database. The URL and the
// Prisma client's pooled connections stay the same.
beforeEach(() => currentFork().reset());
