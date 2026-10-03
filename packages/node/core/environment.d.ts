export declare class ServiceEnvConflict extends Error {
  readonly code: "service_env_conflict";
  readonly variable: string;
  readonly owner: string;
  readonly service: string;
}
/** Only registered service claims conflict. Existing unclaimed values are restored on release. */
export declare function installServiceEnv(service: string, values: Record<string, string>, options?: {
  target?: Record<string, string | undefined>;
  scope?: object;
}): () => void;
