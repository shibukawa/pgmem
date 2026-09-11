package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownSetExpr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
		m.T0[v5].(func(*base.Module, int32))(m, v3)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			if v8 != 0 {
				F_tuplestore_end(m, v8)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					v11 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v11)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v11
					return
				}
			} else {
				v11 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v11)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v11
				return
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v8 != 0 {
			F_tuplestore_end(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v11 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v11)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v11
				return
			}
		} else {
			v11 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)) = uint16(v11)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v11
			return
		}
	}
}
