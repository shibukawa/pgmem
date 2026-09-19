package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_InitShmemAccess(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAccess[0])) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAccess[1])) = l0
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_InitShmemAccess[2])) = l0 + v7
	return
}
func F_ShmemAllocNoError(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
	v14 = base.AtomicRmwXchg32(m, v11, v2, int32(1))
	if v14 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
		F_s_lock(m, v16, int32(_a_F_ShmemAllocNoError_0), int32(208), int32(_a_F_ShmemAllocNoError_1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[1]))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
			v27 = (l0+int32(127))&int32(-128) + v26
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			if base.Ui32(v27) <= base.Ui32(v28) {
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v27
				v34 = v31 + v26
			} else {
				v34 = v2
			}
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
			v37 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36))), uint32(v37))
			return v34
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[1]))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
		v27 = (l0+int32(127))&int32(-128) + v26
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
		if base.Ui32(v27) <= base.Ui32(v28) {
			v31 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[2]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v27
			v34 = v31 + v26
		} else {
			v34 = v2
		}
		v36 = *(*int32)(unsafe.Add(mBase, _c_F_ShmemAllocNoError[0]))
		v37 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v36))), uint32(v37))
		return v34
	}
}
