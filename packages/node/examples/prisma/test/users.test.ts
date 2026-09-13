import { expect, test } from "vitest";
import { prisma } from "../src/db";
import { signUp } from "../src/blog";

test("sign up", async () => {
  await signUp("alice@example.com");
  expect(await prisma.user.count()).toBe(1);
});

test("the next test starts from an empty database", async () => {
  expect(await prisma.user.count()).toBe(0);
  await expect(Promise.all([signUp("bob@example.com"), signUp("bob@example.com")])).rejects.toMatchObject({ code: "P2002" });
});
