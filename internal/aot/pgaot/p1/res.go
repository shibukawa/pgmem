package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintTupleDesc(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = base.I32_wrap_i64(l0)
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v13 = F_psprintf(m, int32(_a_F_ResOwnerPrintTupleDesc_0), v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		m.G0 = v6 + int32(16)
		return v13
	}
}
func F_ResOwnerReleaseWaitEventSet(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = base.I32_wrap_i64(l0)
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(0)
	F_pfree(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
