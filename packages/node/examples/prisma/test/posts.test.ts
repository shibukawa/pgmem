import { expect, test } from "vitest";
import { postsBy, publish, signUp } from "../src/blog";

test("publish in a transaction", async () => {
  await signUp("carol@example.com");
  await publish("carol@example.com", "first");
  await publish("carol@example.com", "second");
  expect((await postsBy("carol@example.com")).map((p) => p.title)).toEqual(["first", "second"]);
});

test("a failed transaction leaves nothing behind", async () => {
  await expect(publish("nobody@example.com", "lost")).rejects.toThrow();
  expect(await postsBy("nobody@example.com")).toEqual([]);
});
