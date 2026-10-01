import { expect, test } from "vitest";
import { shadowPg } from "@pgmem/core";

shadowPg(test, "routes an unchanged Prisma application client without replacing its URL", async () => {
  const previous = process.env.DATABASE_URL;
  process.env.DATABASE_URL = "postgres://bad@127.0.0.1:1/production";
  try {
    const { prisma } = await import("../src/db");
    const { signUp } = await import("../src/blog");
    try {
      await signUp("shadow@example.com");
      expect(await prisma.user.count()).toBe(1);
    } finally {
      await prisma.$disconnect();
    }
  } finally {
    if (previous === undefined) delete process.env.DATABASE_URL;
    else process.env.DATABASE_URL = previous;
  }
});
