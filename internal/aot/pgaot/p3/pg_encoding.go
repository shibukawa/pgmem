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
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = int32(2147483647)
	if l2 == int32(0) {
		v31 = v5
		return v31
	} else {
		if base.B2i32(l0 != int32(39))|base.B2i32(l2 != int32(1)) == int32(0) {
			v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
			if v15 < int32(0) {
				v31 = v5
				return v31
			} else {
				if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
					v31 = int32(1)
					return v31
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen_or_incomplete[0])))
					v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v27
						return v31
					}
				}
			}
		} else {
			if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
				v31 = int32(1)
				return v31
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen_or_incomplete[0])))
				v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = v27
					return v31
				}
			}
		}
	}
}
