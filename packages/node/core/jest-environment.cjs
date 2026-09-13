"use strict";
// Jest test environment: every test file runs against a fork of
// PGMEM_SNAPSHOT, reached through PGMEM_CONTROL. The fork belongs to the
// Jest worker and is reset between the test files the worker runs; its URL
// is written to the test file's process.env (DATABASE_URL, or the names in
// PGMEM_ENV) before the file is loaded.
const { TestEnvironment } = require("jest-environment-node");
const pgmem = require("./index.cjs");

let workerFork;

class PgmemEnvironment extends TestEnvironment {
  async setup() {
    await super.setup();
    const reuse = workerFork !== undefined;
    workerFork ??= pgmem.fork();
    const fork = await workerFork;
    if (reuse) await fork.reset();
    const names = process.env.PGMEM_ENV?.split(",").map((s) => s.trim()).filter(Boolean);
    Object.assign(this.global.process.env, fork.env(names?.length ? names : undefined));
    this.global[Symbol.for("pgmem.currentFork")] = fork;
  }
}

module.exports = PgmemEnvironment;
module.exports.TestEnvironment = PgmemEnvironment;
