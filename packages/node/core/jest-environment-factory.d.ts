import type { TestFixtureOptions } from "./index";
export declare function withPgmemEnvironment<T extends new (...args: any[]) => {
  setup(): Promise<void>;
  teardown(): Promise<void>;
  global: any;
}>(BaseEnvironment: T, options?: TestFixtureOptions & { /** @deprecated Use resetTimeoutMs. */ timeoutMs?: number }): T;
