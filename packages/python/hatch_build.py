"""Hatch build hook: put the platform `pgmem` binary into the wheel.

The binary comes from PGMEM_BINARY when set, otherwise it is cross-compiled
from ../../cmd/pgmem with GOOS/GOARCH (default: the build host). The wheel
is tagged py3-none-<platform>; the binary is static, so one wheel per
platform serves every CPython and PyPy version.
"""

import os
import platform
import shutil
import subprocess
import sys

from hatchling.builders.hooks.plugin.interface import BuildHookInterface

# GOOS/GOARCH -> wheel platform tag. The binary has no libc dependency, so
# the oldest tag each platform accepts is correct.
PLATFORM_TAGS = {
    ("linux", "amd64"): "manylinux_2_17_x86_64.manylinux2014_x86_64",
    ("linux", "arm64"): "manylinux_2_17_aarch64.manylinux2014_aarch64",
    ("darwin", "amd64"): "macosx_10_13_x86_64",
    ("darwin", "arm64"): "macosx_11_0_arm64",
    ("windows", "amd64"): "win_amd64",
    ("windows", "arm64"): "win_arm64",
}


def host_goos_goarch():
    goos = {"Linux": "linux", "Darwin": "darwin", "Windows": "windows"}[platform.system()]
    machine = platform.machine().lower()
    goarch = {"x86_64": "amd64", "amd64": "amd64", "arm64": "arm64", "aarch64": "arm64"}[machine]
    return goos, goarch


class CustomBuildHook(BuildHookInterface):
    def initialize(self, version, build_data):
        if self.target_name != "wheel":
            return
        goos = os.environ.get("GOOS") or host_goos_goarch()[0]
        goarch = os.environ.get("GOARCH") or host_goos_goarch()[1]
        name = "pgmem.exe" if goos == "windows" else "pgmem"
        dst = os.path.join(self.root, "src", "pgmem", "_bin", name)
        if version == "editable":
            # dev install: tests supply PGMEM_BINARY or build on demand
            return
        src = os.environ.get("PGMEM_BINARY")
        if src:
            shutil.copyfile(src, dst)
        else:
            repo = os.path.abspath(os.path.join(self.root, "..", ".."))
            env = dict(os.environ, GOOS=goos, GOARCH=goarch, CGO_ENABLED="0")
            subprocess.run(
                ["go", "build", "-trimpath", "-ldflags=-s -w", "-o", dst, "./cmd/pgmem"],
                cwd=repo, env=env, check=True,
            )
        if goos != "windows":
            os.chmod(dst, 0o755)
        tag = os.environ.get("PGMEM_PLATFORM_TAG") or PLATFORM_TAGS[(goos, goarch)]
        build_data["pure_python"] = False
        build_data["infer_tag"] = False
        build_data["tag"] = f"py3-none-{tag}"
        # _bin/ is gitignored, so include the binary explicitly
        build_data["force_include"][dst] = f"pgmem/_bin/{name}"
        print(f"pgmem: bundled {dst} as {build_data['tag']}", file=sys.stderr)
