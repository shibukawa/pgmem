package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleCreateEventTrigStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v4 != 0 {
		v5 = F_strlen(m, v4)
		mBase = m.M
		F_AppendJumble(m, l0, v4, v5+int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v14 != 0 {
				v15 = F_strlen(m, v14)
				mBase = m.M
				F_AppendJumble(m, l0, v14, v15+int32(1))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					F__jumbleNode(m, l0, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						F__jumbleNode(m, l0, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v20 + int32(1)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					F__jumbleNode(m, l0, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v10 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		if v14 != 0 {
			v15 = F_strlen(m, v14)
			mBase = m.M
			F_AppendJumble(m, l0, v14, v15+int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F__jumbleNode(m, l0, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					F__jumbleNode(m, l0, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v20 + int32(1)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			F__jumbleNode(m, l0, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				F__jumbleNode(m, l0, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
