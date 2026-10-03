"use strict";
const pgmem = require("./index.cjs");
const { installServiceEnv } = require("./environment.cjs");
const { fixtureOptions } = require("./fixture-config.cjs");
let workerFork;

/** Add pgmem to a user-owned Jest environment, including another service's wrapper. */
function withPgmemEnvironment(BaseEnvironment, options = {}) {
  return class extends BaseEnvironment {
    async setup() {
      try {
        await super.setup();
        const cfg = fixtureOptions(options);
        const reuse = workerFork !== undefined;
        workerFork ??= pgmem.fork({ timeoutMs: cfg.forkTimeoutMs });
        const fork = await workerFork;
        if (reuse) await fork.reset({ timeoutMs: cfg.resetTimeoutMs });
        this.releaseEnv = installServiceEnv("pgmem", fork.env(cfg.env?.length ? cfg.env : undefined, { readOnly: cfg.readOnly }), { target: this.global.process.env, scope: this.global });
        this.global[Symbol.for("pgmem.currentFork")] = fork;
      } catch (err) {
        const failedFork = workerFork;
        workerFork = undefined;
        const failures = [err];
        try { if (failedFork) await (await failedFork).close(); }
        catch (cleanup) { if (cleanup !== err) failures.push(cleanup); }
        try { this.releaseEnv?.(); } catch (cleanup) { failures.push(cleanup); }
        try { await super.teardown(); }
        catch (cleanup) { failures.push(cleanup); }
        if (failures.length > 1) throw new AggregateError(failures, "pgmem: environment setup and cleanup both failed");
        throw err;
      }
    }
    async teardown() {
      let failure;
      try { this.releaseEnv?.(); } catch (err) { failure = err; }
      try { await super.teardown(); } catch (err) {
        if (failure) throw new AggregateError([failure, err], "pgmem: environment cleanup failures");
        throw err;
      }
      if (failure) throw failure;
    }
  };
}
module.exports.withPgmemEnvironment = withPgmemEnvironment;
