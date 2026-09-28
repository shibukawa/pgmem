package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownSetExpr(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+48))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		m.T0[v7].(func(*base.Module, int32))(m, v5)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+44))
			if v10 != 0 {
				F_tuplestore_end(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v4)+58)) = uint16(v13)
					*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v13
					return
				}
			} else {
				v13 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v4)+58)) = uint16(v13)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v13
				return
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)+44))
		if v10 != 0 {
			F_tuplestore_end(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v4)+58)) = uint16(v13)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v13
				return
			}
		} else {
			v13 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v4)+58)) = uint16(v13)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = v13
			return
		}
	}
}
