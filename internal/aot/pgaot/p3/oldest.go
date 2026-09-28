package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AdvanceOldestClogXid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[0]))
	v9 = F_LWLockAcquire(m, v5+int32(_a_F_AdvanceOldestClogXid_0), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int32(3)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
		if base.B2i32(base.Ui32(l0) < base.Ui32(v11))|base.B2i32(base.Ui32(v15) < base.Ui32(v11)) == int32(0) {
			if v15-l0 < int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l0
			} else {
			}
		} else {
			if base.Ui32(l0) <= base.Ui32(v15) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = l0
			}
		}
		v27 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[0]))
		F_LWLockRelease(m, v27+int32(_a_F_AdvanceOldestClogXid_0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			return
		}
	}
}
