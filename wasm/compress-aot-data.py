#!/usr/bin/env python3
"""Create a deterministic gzip copy of wasm2go's linear-memory data."""

import gzip
import shutil
import sys


def main() -> None:
    if len(sys.argv) != 3:
        raise SystemExit(f"usage: {sys.argv[0]} INPUT OUTPUT")

    source, destination = sys.argv[1:]
    with open(source, "rb") as src, open(destination, "wb") as dst:
        with gzip.GzipFile(
            filename="", mode="wb", fileobj=dst, compresslevel=9, mtime=0
        ) as compressed:
            shutil.copyfileobj(src, compressed)


if __name__ == "__main__":
    main()
