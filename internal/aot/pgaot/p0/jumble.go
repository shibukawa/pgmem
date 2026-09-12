package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateRoleStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	F_AppendJumble32(m, l0, l1+int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v8 != 0 {
			v9 = F_strlen(m, v8)
			mBase = m.M
			F_AppendJumble(m, l0, v8, v9+int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v14 + int32(1)
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				return
			}
		}
	}
}
