package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	if l1 != 0 {
		F_ResourceOwnerForget(m, l1, l0, int32(1743276))
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v8 = v6 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
			if v8 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					F_MemoryContextDelete(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = v6 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v8
		if v8 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				F_MemoryContextDelete(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
