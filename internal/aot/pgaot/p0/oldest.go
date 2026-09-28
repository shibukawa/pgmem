package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetOldestRestartPoint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[0]))
	v8 = F_LWLockAcquire(m, v4+int32(1152), int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[1]))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestRestartPoint[0]))
		F_LWLockRelease(m, v17+int32(1152))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			return
		}
	}
}
