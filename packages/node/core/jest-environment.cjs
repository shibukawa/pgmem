"use strict";
const { TestEnvironment } = require("jest-environment-node");
const { withPgmemEnvironment } = require("./jest-environment-factory.cjs");
const PgmemEnvironment = withPgmemEnvironment(TestEnvironment);
module.exports = PgmemEnvironment;
module.exports.TestEnvironment = PgmemEnvironment;
module.exports.withPgmemEnvironment = withPgmemEnvironment;
