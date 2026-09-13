/** Control protocol version this package speaks; the binary must match. */
export declare const PROTOCOL: 1;

/** A listening server: its URL and the parts of it. */
export interface Endpoint {
  readonly id: string;
  /** postgres://user@127.0.0.1:port/database?sslmode=disable */
  readonly url: string;
  readonly host: string;
  readonly port: number;
  readonly user: string;
  readonly database: string;
  /**
   * Environment variables pointing at this server: DATABASE_URL by default,
   * or the given names. PGHOST, PGPORT, PGUSER, PGDATABASE and PGSSLMODE get
   * the matching part; any other name gets the URL.
   */
  env(names?: string | readonly string[]): Record<string, string>;
}

export interface StartOptions {
  /** Database to create and serve. Default "postgres". */
  database?: string;
  /** Superuser name. Default "postgres". */
  user?: string;
  /** postgres -c settings, e.g. { log_statement: "all" }. Values must not contain commas. */
  params?: Record<string, string | number | boolean>;
  /**
   * Runs once against the template before it is snapshotted: apply
   * migrations and seed data here. Close or commit every connection it
   * opens; the snapshot waits for open transactions.
   */
  prepare?: (template: Endpoint) => unknown;
  /** Forks alive at once before fork() waits. Default: the number of CPUs. */
  maxForks?: number;
  /** Serve the control socket other processes fork through. Default true. */
  control?: boolean;
  /**
   * End a connection that has waited this long behind another connection's
   * idle transaction, with SQLSTATE 55P03. Default 2000; negative waits forever.
   */
  waitTimeoutMs?: number;
  /** Pass the server log through to stderr. */
  log?: boolean;
  /** The pgmem binary. Default: PGMEM_BINARY, then the @pgmem/<platform> package. */
  binary?: string;
  /** Default 30000. */
  startupTimeoutMs?: number;
  /** How long the snapshot after prepare waits for open transactions. Default 30000. */
  snapshotTimeoutMs?: number;
}

/** An error answered by pgmem. */
export declare class PgmemError extends Error {
  /** busy, pool_timeout, unknown_id, snapshot_closed, unauthorized, forbidden, protocol or internal. */
  readonly code: string;
}

export interface ForkOptions {
  /** Fail with code pool_timeout instead of waiting longer for a free slot. */
  timeoutMs?: number;
}

/** A frozen copy of a server's data that forks start from. */
export declare class PgmemSnapshot {
  readonly id: string;
  /** Start a server on a fresh copy; waits while maxForks forks are alive. */
  fork(options?: ForkOptions): Promise<PgmemFork>;
  /** Run fn with a fresh fork and close the fork afterwards. */
  withFork<T>(fn: (fork: PgmemFork) => T | Promise<T>, options?: ForkOptions): Promise<T>;
  /** Refuse further forks; running forks keep working. */
  close(): Promise<void>;
}

/** A server started from a snapshot. */
export declare class PgmemFork implements Endpoint {
  readonly id: string;
  readonly url: string;
  readonly host: string;
  readonly port: number;
  readonly user: string;
  readonly database: string;
  env(names?: string | readonly string[]): Record<string, string>;
  /**
   * Put the data back to the snapshot the fork came from (or the given one)
   * in place. The URL and open connections stay valid; pooled connections
   * continue in a fresh session with their prepared statements intact.
   * Fails with code busy when a connection keeps a transaction open longer
   * than timeoutMs (default 5000).
   */
  reset(options?: { snapshot?: PgmemSnapshot | string; timeoutMs?: number }): Promise<void>;
  /** Snapshot this fork, for example after beforeAll seeding, to reset to later. */
  snapshot(options?: { maxForks?: number; timeoutMs?: number }): Promise<PgmemSnapshot>;
  /** Stop the fork and free its slot. Idle client connections are left to their pools. */
  close(): Promise<void>;
  [Symbol.asyncDispose](): Promise<void>;
}

/** A pgmem process owned by this process. */
export declare class PgmemServer {
  /** Start pgmem, run prepare against the template and snapshot it. */
  static start(options?: StartOptions): Promise<PgmemServer>;
  readonly pid: number;
  readonly version: string;
  /** The server prepare ran against. */
  readonly template: Endpoint;
  /** template.url */
  readonly url: string;
  /** pgmem-control://token@127.0.0.1:port, unless started with control: false. */
  readonly controlUrl: string | undefined;
  /** The snapshot taken after prepare. */
  readonly snapshot: PgmemSnapshot;
  /** PGMEM_CONTROL and PGMEM_SNAPSHOT, for the processes that fork from this server. */
  env(): { PGMEM_CONTROL: string; PGMEM_SNAPSHOT: string };
  /** Fork the prepared snapshot. */
  fork(options?: ForkOptions): Promise<PgmemFork>;
  /** Run fn with a fork of the prepared snapshot and close it afterwards. */
  withFork<T>(fn: (fork: PgmemFork) => T | Promise<T>, options?: ForkOptions): Promise<T>;
  /** Shut the process down with every fork. */
  close(): Promise<void>;
  [Symbol.asyncDispose](): Promise<void>;
}

export interface ConnectOptions {
  /** Default: PGMEM_CONTROL. */
  controlUrl?: string;
  /** Snapshot id to fork. Default: PGMEM_SNAPSHOT. */
  snapshot?: string;
}

/** A connection to a PgmemServer's control socket from another process. Forks made through it close with it. */
export declare class PgmemClient {
  static connect(options?: ConnectOptions): Promise<PgmemClient>;
  readonly snapshot: PgmemSnapshot | undefined;
  fork(options?: ForkOptions): Promise<PgmemFork>;
  withFork<T>(fn: (fork: PgmemFork) => T | Promise<T>, options?: ForkOptions): Promise<T>;
  close(): void;
}

/** With options, a new client; without, the process-wide client for PGMEM_CONTROL. */
export declare function connect(options?: ConnectOptions): Promise<PgmemClient>;
/** Fork PGMEM_SNAPSHOT through the process-wide client. */
export declare function fork(options?: ForkOptions): Promise<PgmemFork>;
/** Run fn with a fork of PGMEM_SNAPSHOT and close it afterwards. */
export declare function withFork<T>(fn: (fork: PgmemFork) => T | Promise<T>, options?: ForkOptions): Promise<T>;
/**
 * Give this process its fork and write the fork's URL to process.env
 * (DATABASE_URL, or the names in env or PGMEM_ENV). A process that already
 * has one resets it instead. @pgmem/core/register calls this.
 */
export declare function useFork(options?: { env?: readonly string[] }): Promise<PgmemFork>;
/** The fork of this test file, made by @pgmem/core/register or @pgmem/core/jest-environment. */
export declare function currentFork(): PgmemFork;
/** The pgmem binary: the argument, PGMEM_BINARY, then the @pgmem/<platform> package. */
export declare function resolveBinary(binary?: string): string;

export default PgmemServer;
