"""Aggregate raw.jsonl samples into medians and attach sizes and machine info."""
import json, platform, statistics, subprocess, sys
from collections import defaultdict

raw, sizes = sys.argv[1], sys.argv[2]
# Memory fields. For docker and testcontainers they come only from runs on
# a freshly restarted OrbStack VM (fresh_vm), where the VM's host-side
# growth is measurable; timing fields come only from the other runs, where
# the VM was already warm as on a developer's machine.
MEMORY = {
    "mem_idle_mb", "mem_after_mb", "service_idle_mb", "service_after_mb",
    "rss_idle_mb", "container_mb", "container_after_mb", "vm_total_mb",
    "process_mb", "vm_growth_mb", "process_after_mb", "vm_growth_after_mb",
    "mem_idle_min_mb", "mem_idle_max_mb", "mem_after_min_mb", "mem_after_max_mb",
    "service_idle_min_mb", "service_idle_max_mb", "service_after_min_mb", "service_after_max_mb",
}
CONTAINER_TARGETS = {"docker", "testcontainers"}

samples = defaultdict(lambda: defaultdict(list))
versions = {}
for line in open(raw):
    r = json.loads(line)
    t = r["target"]
    fresh = bool(r.get("fresh_vm"))
    if t in CONTAINER_TARGETS and fresh:
        # Older raw rows contain the same idle components but not the service
        # total; derive it without mixing in the OrbStack VM's net change.
        r.setdefault("service_idle_mb", r.get("container_mb", 0) + r.get("process_mb", 0))
        if "container_after_mb" in r and "process_after_mb" in r:
            r.setdefault("service_after_mb", r["container_after_mb"] + r["process_after_mb"])
    elif t == "binary":
        r.setdefault("service_idle_mb", r.get("mem_idle_mb"))
    else:
        r.setdefault("service_idle_mb", r.get("mem_idle_mb"))
        if "mem_after_mb" in r:
            r.setdefault("service_after_mb", r["mem_after_mb"])
    for k, v in r.items():
        if not isinstance(v, (int, float)) or isinstance(v, bool) or not v:
            continue
        if t in CONTAINER_TARGETS and (k in MEMORY) != fresh:
            continue
        samples[t][k].append(v)
    if r.get("version"):
        versions[t] = r["version"]

def cmd(*a):
    try:
        return subprocess.check_output(a, text=True, stderr=subprocess.DEVNULL).strip()
    except Exception:
        return ""

def orbstack_version():
    for line in cmd("orb", "version").splitlines():
        if line.lower().startswith("version:"):
            return "OrbStack " + line.split(":", 1)[1].split()[0]
    return ""

targets = {}
for t, fields in samples.items():
    targets[t] = {k: round(statistics.median(v), 2) for k, v in fields.items()}
    for field in ("mem_idle_mb", "mem_after_mb", "service_idle_mb", "service_after_mb"):
        if fields.get(field):
            stem = field.removesuffix("_mb")
            targets[t][f"{stem}_min_mb"] = round(min(fields[field]), 2)
            targets[t][f"{stem}_max_mb"] = round(max(fields[field]), 2)
    if t == "pgmem":
        # pgmem runs the server in the measured process, so all host growth
        # belongs to that process even in older raw samples without components.
        targets[t]["process_mb"] = targets[t]["mem_idle_mb"]
        targets[t]["process_after_mb"] = targets[t]["mem_after_mb"]
    targets[t]["runs"] = max(len(v) for k, v in fields.items() if k not in MEMORY)
    if t in CONTAINER_TARGETS and fields.get("mem_idle_mb"):
        targets[t]["mem_runs"] = len(fields["mem_idle_mb"])
    if t in versions:
        targets[t]["postgres"] = versions[t]

print(json.dumps({
    "measured_at": cmd("date", "-u", "+%Y-%m-%d"),
    "machine": {
        "cpu": cmd("sysctl", "-n", "machdep.cpu.brand_string") or platform.processor(),
        "cores": cmd("sysctl", "-n", "hw.ncpu"),
        "memory_gb": round(int(cmd("sysctl", "-n", "hw.memsize") or 0) / 2**30),
        "os": f"{platform.system()} {platform.release()}",
        "go": cmd("go", "env", "GOVERSION"),
        "docker": cmd("docker", "version", "--format", "{{.Server.Version}}"),
        "docker_runtime": orbstack_version(),
        "devbox": cmd("devbox", "version").splitlines()[-1] if cmd("devbox", "version") else "",
    },
    "targets": targets,
    "sizes": json.load(open(sizes)),
}, indent=2))
