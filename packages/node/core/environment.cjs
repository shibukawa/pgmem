"use strict";
const OWNERS = Symbol.for("test.serviceEnvOwners");

class ServiceEnvConflict extends Error {
  constructor(variable, owner, service) {
    super(`test environment: ${service} cannot claim ${variable}; already owned by ${owner}`);
    this.name = "ServiceEnvConflict";
    this.code = "service_env_conflict";
    this.variable = variable;
    this.owner = owner;
    this.service = service;
  }
}

/** Register environment ownership; returns an idempotent release function. */
function installServiceEnv(service, values, { target = process.env, scope = globalThis } = {}) {
  if (typeof service !== "string" || !service) throw new TypeError("service must be a nonempty name");
  const registry = scope[OWNERS] ??= new WeakMap();
  let owners = registry.get(target);
  if (!owners) registry.set(target, owners = new Map());
  const entries = Object.entries(values);
  for (const [name, value] of entries) {
    if (typeof value !== "string") throw new TypeError(`${service}: ${name} must be a string`);
    const prior = owners.get(name);
    if (prior && prior.service !== service) throw new ServiceEnvConflict(name, prior.service, service);
  }
  const lease = {};
  const before = entries.map(([name]) => [name, owners.get(name), target[name]]);
  try {
    for (const [name, value] of entries) {
      const prior = owners.get(name);
      owners.set(name, { service, lease, original: prior ? prior.original : target[name], value });
      target[name] = value;
    }
  } catch (error) {
    const failures = [error];
    for (const [name, claim, value] of before) {
      if (claim) owners.set(name, claim); else owners.delete(name);
      try {
        if (value === undefined) delete target[name]; else target[name] = value;
      } catch (cleanup) { failures.push(cleanup); }
    }
    if (failures.length > 1) throw new AggregateError(failures, `${service}: environment install and rollback failed`);
    throw error;
  }
  return () => {
    const failures = [];
    for (const [name] of entries) {
      const claim = owners.get(name);
      if (claim?.lease !== lease) continue;
      try {
        if (target[name] === claim.value) {
          if (claim.original === undefined) delete target[name];
          else target[name] = claim.original;
        }
      } catch (error) { failures.push(error); }
      finally { owners.delete(name); }
    }
    if (failures.length) throw new AggregateError(failures, `${service}: environment release failed`);
  };
}
module.exports = { installServiceEnv, ServiceEnvConflict };
