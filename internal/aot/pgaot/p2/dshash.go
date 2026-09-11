package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dshash_seq_term(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if int32(0) <= v3 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+32))
		F_LWLockRelease(m, v7+v3*int32(20)+int32(8))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
