package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerPrintTupleDesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
	v11 = F_psprintf(m, int32(_a_F_ResOwnerPrintTupleDesc_0), v5)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v11
	}
}
func F_ResOwnerReleaseWaitEventSet(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_pfree(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
