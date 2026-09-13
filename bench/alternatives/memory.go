package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// footprintMB returns the summed phys_footprint of pids as reported by
// macOS footprint(1). Unlike RSS it does not count clean, file-backed
// pages the kernel can drop, and it counts memory a process maps from the
// hypervisor. Processes that exited meanwhile are skipped.
func footprintMB(ctx context.Context, pids ...int) (float64, error) {
	dir, err := os.MkdirTemp("", "footprint-")
	if err != nil {
		return 0, err
	}
	defer os.RemoveAll(dir)
	var total float64
	for _, pid := range pids {
		out := filepath.Join(dir, strconv.Itoa(pid)+".json")
		cmd := exec.CommandContext(ctx, "/usr/bin/footprint", "-f", "bytes", "--noCategories", "-j", out, strconv.Itoa(pid))
		if err := cmd.Run(); err != nil {
			continue
		}
		b, err := os.ReadFile(out)
		if err != nil {
			continue
		}
		var r struct {
			Total     float64 `json:"total footprint"`
			Processes []struct {
				Auxiliary struct {
					PhysFootprint float64 `json:"phys_footprint"`
				} `json:"auxiliary"`
			} `json:"processes"`
		}
		if err := json.Unmarshal(b, &r); err != nil {
			return 0, fmt.Errorf("footprint %d: %w", pid, err)
		}
		// phys_footprint is the kernel's ledger (what Activity Monitor
		// shows). For the OrbStack VM it includes guest memory mapped
		// through the hypervisor, which the region-based total misses.
		if len(r.Processes) > 0 && r.Processes[0].Auxiliary.PhysFootprint > 0 {
			total += r.Processes[0].Auxiliary.PhysFootprint
		} else {
			total += r.Total
		}
	}
	return total / (1 << 20), nil
}

// orbVM tracks the OrbStack virtual machine process that runs every
// container on macOS. Its host-side footprint can rise or fall as the VM
// reclaims memory, so VM net change is recorded separately from container
// usage and benchmark-process growth.
type orbVM struct {
	baseMB float64
}

func vmPID(ctx context.Context) (int, error) {
	out, err := run(ctx, "pgrep", "-f", "OrbStack Helper vmgr")
	if err != nil {
		return 0, fmt.Errorf("OrbStack VM process not found: %w", err)
	}
	return strconv.Atoi(strings.Fields(out)[0])
}

// restart stops and starts OrbStack, waits for the Docker API and records
// the idle VM footprint. Testcontainers' Ryuk reaper outlives the previous
// run by a few seconds, so containers labeled by testcontainers get up to
// 30 s to exit; any other running container makes it refuse, since a
// restart would stop it.
func (v *orbVM) restart(ctx context.Context) error {
	deadline := time.Now().Add(30 * time.Second)
	for {
		all, err := run(ctx, "docker", "ps", "-q")
		if err != nil {
			return err
		}
		if all == "" {
			break
		}
		ours, err := run(ctx, "docker", "ps", "-q", "--filter", "label=org.testcontainers=true")
		if err != nil {
			return err
		}
		if len(strings.Fields(ours)) < len(strings.Fields(all)) || time.Now().After(deadline) {
			return fmt.Errorf("refusing to restart OrbStack: containers are running (%s)", strings.Join(strings.Fields(all), " "))
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(500 * time.Millisecond):
		}
	}
	if _, err := run(ctx, "orbctl", "stop"); err != nil {
		return err
	}
	if _, err := run(ctx, "orbctl", "start"); err != nil {
		return err
	}
	for {
		if _, err := run(ctx, "docker", "info"); err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(200 * time.Millisecond):
		}
	}
	time.Sleep(3 * time.Second) // let the VM settle; the idle footprint is flat after this
	pid, err := vmPID(ctx)
	if err != nil {
		return err
	}
	v.baseMB, err = footprintMB(ctx, pid)
	return err
}

// usage returns the VM's net footprint change since restart and its current
// absolute footprint. The net change may shrink after a workload.
func (v *orbVM) usage(ctx context.Context) (growth, total float64, err error) {
	pid, err := vmPID(ctx)
	if err != nil {
		return 0, 0, err
	}
	total, err = footprintMB(ctx, pid)
	return total - v.baseMB, total, err
}

// addBenchmarkProcessGrowth records the harness process's increase since the
// fresh-VM baseline. The comparable service reading sums that increase with
// the container stats; the VM net change stays a separate host-level reading.
func addBenchmarkProcessGrowth(ctx context.Context, m memory, baseline float64) (memory, error) {
	current, err := footprintMB(ctx, os.Getpid())
	if err != nil {
		return memory{}, err
	}
	m.Process = current - baseline
	m.VMGrowth = m.Host
	m.Host += m.Process
	m.Service = m.Container + m.Process
	return m, nil
}
