"use strict";
function deadline(value, fallback, name) {
  if (value === undefined) return fallback;
  const result = typeof value === "string" ? Number(value) : value;
  if (!Number.isFinite(result) || result <= 0) throw new TypeError(`pgmem: ${name} must be a positive number of milliseconds`);
  return result;
}
function fixtureOptions(options = {}, env = process.env) {
  const readOnly = options.readOnly ?? env.PGMEM_READ_ONLY;
  if (readOnly !== undefined && ![true, false, "true", "false"].includes(readOnly)) throw new TypeError("pgmem: readOnly / PGMEM_READ_ONLY must be true or false");
  return {
    env: options.env ?? env.PGMEM_ENV?.split(",").map((s) => s.trim()).filter(Boolean),
    forkTimeoutMs: deadline(options.forkTimeoutMs ?? env.PGMEM_FORK_TIMEOUT_MS, 30000, "forkTimeoutMs"),
    resetTimeoutMs: deadline(options.resetTimeoutMs ?? options.timeoutMs ?? env.PGMEM_RESET_TIMEOUT_MS, 5000, "resetTimeoutMs"),
    readOnly: readOnly === true || readOnly === "true",
  };
}
module.exports = { fixtureOptions };
