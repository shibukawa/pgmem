package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResOwnerReleaseCachedPlan(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	v7 = v5 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v7
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(0)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+8)))
		if v11 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
			F_MemoryContextDelete(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ResOwnerReleaseTupleDesc(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	v7 = v5 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v7
	if v7 == int32(0) {
		F_FreeTupleDesc(m, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
