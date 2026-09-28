package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReleaseCachedPlan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	if l1 != 0 {
		F_ResourceOwnerForget(m, l1, base.I64_extend_i32_u(l0), int32(_a_F_ReleaseCachedPlan_0))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v9 = v7 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v9
			if v9 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
				if v13 != 0 {
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					F_MemoryContextDelete(m, v14)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = v7 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v9
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v13 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				F_MemoryContextDelete(m, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
