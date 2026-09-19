package pgmem

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func TestDefaultMaxForks(t *testing.T) {
	const gb = 1 << 30
	def := []string{"-c", "shared_buffers=32MB", "-c", "io_method=sync"}
	cases := []struct {
		name   string
		params []string
		limit  uint64
		cpus   int
		want   int
	}{
		{"7 GB CI runner, default buffers", def, 7 * gb, 2, 28},
		{"16 GB workstation", def, 16 * gb, 8, 64},
		{"2 GB container", def, 2 * gb, 8, 8},
		{"no shared_buffers argument uses the default", nil, 7 * gb, 2, 28},
		{"larger buffers cost more", []string{"--shared_buffers=128MB"}, 7 * gb, 2, 11},
		{"buffers in blocks", []string{"-c", "shared_buffers=4096"}, 7 * gb, 2, 28},
		{"quoted value", []string{"-c", "shared_buffers='32MB'"}, 7 * gb, 2, 28},
		{"unknown memory falls back to the CPU count", def, 0, 2, 2},
		{"unknown memory, cpus 0", def, 0, 0, 1},
		{"too little memory still allows one fork", def, 64 << 20, 8, 1},
	}
	for _, c := range cases {
		if got := defaultMaxForks(c.params, c.limit, c.cpus); got != c.want {
			t.Errorf("%s: defaultMaxForks = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestParseMemorySetting(t *testing.T) {
	cases := []struct {
		in   string
		want uint64
		ok   bool
	}{
		{"32MB", 32 << 20, true},
		{"1.5GB", 3 << 29, true},
		{"1 GB", 1 << 30, true},
		{"4096", 4096 * 8192, true},
		{"512kB", 512 << 10, true},
		{"1TB", 1 << 40, true},
		{"64B", 64, true},
		{"32mb", 0, false},
		{"MB", 0, false},
		{"", 0, false},
		{"-1MB", 0, false},
	}
	for _, c := range cases {
		got, ok := parseMemorySetting(c.in, 8192)
		if ok != c.ok || got != c.want {
			t.Errorf("parseMemorySetting(%q) = %d, %v; want %d, %v", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestMemoryLimit(t *testing.T) {
	switch runtime.GOOS {
	case "linux", "darwin", "windows":
	default:
		t.Skip("physical memory is not read on", runtime.GOOS)
	}
	t.Logf("memoryLimit = %d MB, default MaxForks here = %d", memoryLimit()>>20, defaultMaxForks(nil, memoryLimit(), runtime.GOMAXPROCS(0)))
	if got := memoryLimit(); got < 256<<20 {
		t.Fatalf("memoryLimit = %d, want at least 256 MB on a machine running the tests", got)
	}
	// A GOMEMLIMIT below the machine's (or container's) memory wins.
	before := memoryLimit()
	old := debug.SetMemoryLimit(int64(before / 2))
	defer debug.SetMemoryLimit(old)
	if got := memoryLimit(); got != before/2 {
		t.Fatalf("memoryLimit with GOMEMLIMIT=%d = %d, want %d", before/2, got, before/2)
	}
}
