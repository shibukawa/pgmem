import { defineConfig } from "vitest/config";

export default defineConfig({
  test: {
    globalSetup: ["./test/global-setup.ts"],
    setupFiles: ["@pgmem/core/register", "./test/reset.ts"],
  },
});
