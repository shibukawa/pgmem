import os
import shutil
import sys
from pathlib import Path


def find_binary(explicit=None):
    """Locate the pgmem server binary.

    Order: the ``explicit`` argument, the ``PGMEM_BINARY`` environment
    variable, the binary bundled in this wheel, then ``pgmem`` on PATH.
    """
    name = "pgmem.exe" if sys.platform == "win32" else "pgmem"
    candidates = [explicit, os.environ.get("PGMEM_BINARY"), str(Path(__file__).parent / "_bin" / name)]
    for c in candidates:
        if c and os.path.isfile(c):
            return c
    on_path = shutil.which("pgmem")
    if on_path:
        return on_path
    raise FileNotFoundError(
        "pgmem binary not found: this wheel has no bundled binary for your platform; "
        "set PGMEM_BINARY to a build of github.com/shibukawa/pgmem/cmd/pgmem"
    )
