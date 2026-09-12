package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_encoding_mblen_or_incomplete(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = int32(2147483647)
	if l2 == int32(0) {
		v28 = v5
		return v28
	} else {
		if l0 == int32(39) {
			if l2 != int32(1) {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[1119])))
				v24 = m.T0[v23].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v24
					return v28
				}
			} else {
				v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v12 {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[1119])))
					v24 = m.T0[v23].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = v24
						return v28
					}
				} else {
					v28 = v5
					return v28
				}
			}
		} else {
			if base.Ui32(int32(41)) < base.Ui32(l0) {
				v28 = int32(1)
				return v28
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_consts[1119])))
				v24 = m.T0[v23].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v24
					return v28
				}
			}
		}
	}
}
