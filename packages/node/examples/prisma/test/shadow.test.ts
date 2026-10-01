import { expect, test } from "vitest";
import { shadowPg } from "@pgmem/core";

shadowPg(test, "routes an unchanged Prisma application client without replacing its URL", async () => {
  const configuredUrl = process.env.DATABASE_URL;
  const { prisma } = await import("../src/db");
  const { signUp } = await import("../src/blog");
  try {
    await signUp("shadow@example.com");
    expect(await prisma.user.count()).toBe(1);
    expect(process.env.DATABASE_URL).toBe(configuredUrl);
  } finally {
    await prisma.$disconnect();
  }
});

shadowPg(test, "starts the next Prisma case from a new fork", async () => {
  const { prisma } = await import("../src/db");
  try {
    expect(await prisma.user.count()).toBe(0);
  } finally {
    await prisma.$disconnect();
  }
});
