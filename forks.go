package pgmem

import (
	"math"
	"runtime/debug"
	"strconv"
	"strings"
)

// A fork's memory is its backend's linear memory: the buffer cache
// (shared_buffers) plus PostgreSQL's code data, heap and the other
// shared-memory structures. That rest measured 25 MB resident right after
// boot with the default settings (2026-09-19, Apple M3), before the buffer
// cache fills, so a fork is budgeted at shared_buffers + forkOverheadBytes.
const forkOverheadBytes = 32 << 20

// forkBudgetDivisor is the share of the memory limit that one snapshot's
// live forks may occupy: a quarter, leaving room for the template server,
// the test process and whatever else runs alongside.
const forkBudgetDivisor = 4

// defaultSharedBuffers is the shared_buffers Options.withDefaults sets.
const defaultSharedBuffers = 32 << 20

// defaultMaxForks is the fork cap when SnapshotOptions.MaxForks is 0: a
// quarter of the memory limit divided by the cost of one fork under
// params. limit is the process's memory limit in bytes; 0 (unknown) falls
// back to cpus, the number of forks that can execute at once.
func defaultMaxForks(params []string, limit uint64, cpus int) int {
	if limit == 0 {
		return max(cpus, 1)
	}
	perFork := sharedBuffersBytes(params) + forkOverheadBytes
	n := limit / forkBudgetDivisor / perFork
	if n < 1 {
		return 1
	}
	if n > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(n)
}

// memoryLimit is the memory this process may use: the smallest of
// GOMEMLIMIT (when set), the cgroup limit (Linux, as containers set it)
// and the physical memory. 0 means it could not be determined.
func memoryLimit() uint64 {
	limit := physicalMemory()
	if c := cgroupMemoryLimit(); c > 0 && (limit == 0 || c < limit) {
		limit = c
	}
	// A negative argument reports the current limit without changing it;
	// math.MaxInt64 is the runtime's "no limit".
	if l := debug.SetMemoryLimit(-1); l > 0 && l < math.MaxInt64 && (limit == 0 || uint64(l) < limit) {
		limit = uint64(l)
	}
	return limit
}

// sharedBuffersBytes is the shared_buffers setting in params, or the
// default when it is absent or unparsable (postgres rejects it then).
func sharedBuffersBytes(params []string) uint64 {
	if v, ok := settingValue(params, "shared_buffers"); ok {
		if n, ok := parseMemorySetting(v, 8192); ok && n > 0 {
			return n
		}
	}
	return defaultSharedBuffers
}

// settingValue is the value of the "name=value" argument in params, in
// either postgres command-line form: after a -c, or as --name=value.
func settingValue(params []string, name string) (string, bool) {
	for _, p := range params {
		p = strings.TrimPrefix(strings.TrimSpace(p), "--")
		if v, ok := strings.CutPrefix(p, name+"="); ok {
			return strings.Trim(strings.TrimSpace(v), `'"`), true
		}
	}
	return "", false
}

// parseMemorySetting parses a PostgreSQL memory setting such as "32MB",
// "1.5 GB" or "4096". Units are case-sensitive as in postgresql.conf; a
// bare number counts in blockBytes, the setting's own unit.
func parseMemorySetting(s string, blockBytes uint64) (uint64, bool) {
	s = strings.TrimSpace(s)
	i := 0
	for i < len(s) && (s[i] >= '0' && s[i] <= '9' || s[i] == '.') {
		i++
	}
	if i == 0 {
		return 0, false
	}
	num, err := strconv.ParseFloat(s[:i], 64)
	if err != nil || num < 0 {
		return 0, false
	}
	unit := blockBytes
	switch strings.TrimSpace(s[i:]) {
	case "":
	case "B":
		unit = 1
	case "kB":
		unit = 1 << 10
	case "MB":
		unit = 1 << 20
	case "GB":
		unit = 1 << 30
	case "TB":
		unit = 1 << 40
	default:
		return 0, false
	}
	return uint64(num * float64(unit)), true
}
