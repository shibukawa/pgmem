// ESM entry. The implementation lives in index.cjs so that Jest, which
// loads test files as CommonJS, can use the same code.
import pgmem from "./index.cjs";

export const {
  PROTOCOL,
  PgmemError,
  PgmemServer,
  PgmemClient,
  PgmemSnapshot,
  PgmemFork,
  connect,
  fork,
  withFork,
  useFork,
  currentFork,
  resolveBinary,
} = pgmem;

export default PgmemServer;
