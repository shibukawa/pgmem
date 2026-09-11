package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_ls_tmpdir_1arg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_ls_tmpdir(m, l0, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
