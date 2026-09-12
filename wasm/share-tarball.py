#!/usr/bin/env python3
"""Pack the installed share tree into internal/assets/share.tar.gz.

usage: share-tarball.py <share/postgresql dir> <out.tar.gz>

Deterministic (sorted entries, zero timestamps and owners, gzip without a
name or time) so a rebuild that changes nothing leaves the archive alone,
and portable: hard links (zic's timezone aliases) are stored as plain
files, and macOS resource-fork files and the postgres.description files
that only pg_dump wants are left out.
"""
import gzip, io, os, sys, tarfile

src, dst = sys.argv[1], sys.argv[2]
skip = {'postgres.description', 'postgres.shdescription'}
buf = io.BytesIO()
with tarfile.open(fileobj=buf, mode='w', format=tarfile.PAX_FORMAT) as tar:
    for dirpath, dirnames, filenames in os.walk(src):
        dirnames.sort()
        rel = os.path.relpath(dirpath, src)
        for name in sorted(filenames):
            if name in skip or name.startswith('._'):
                continue
            full = os.path.join(dirpath, name)
            if os.path.islink(full):
                info = tarfile.TarInfo('./' + os.path.normpath(os.path.join(rel, name)))
                info.type = tarfile.SYMTYPE
                info.linkname = os.readlink(full)
                info.mode = 0o777
                tar.addfile(info)
                continue
            info = tarfile.TarInfo('./' + os.path.normpath(os.path.join(rel, name)))
            info.type = tarfile.REGTYPE
            info.size = os.path.getsize(full)
            info.mode = os.stat(full).st_mode & 0o777
            with open(full, 'rb') as f:
                tar.addfile(info, f)
        for name in dirnames:
            info = tarfile.TarInfo('./' + os.path.normpath(os.path.join(rel, name)))
            info.type = tarfile.DIRTYPE
            info.mode = 0o755
            tar.addfile(info)
with open(dst, 'wb') as out:
    with gzip.GzipFile(fileobj=out, mode='wb', compresslevel=9, mtime=0) as g:
        g.write(buf.getvalue())
print(f'{dst}: {os.path.getsize(dst)} bytes')
