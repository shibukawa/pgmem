package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetMaxSnapshotSubxidCount(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v4 = *(*int32)(unsafe.Add(mBase, _consts[771]))
	return (v2 + v4) * int32(65)
}
