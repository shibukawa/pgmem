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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = int32(2147483647)
	if l2 == int32(0) {
		v26 = v5
		return v26
	} else {
		if l0 == int32(39) {
			if l2 != int32(1) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen_or_incomplete[0])))
				v22 = m.T0[v21].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v22
					return v26
				}
			} else {
				v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
				if int32(0) <= v12 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen_or_incomplete[0])))
					v22 = m.T0[v21].(func(*base.Module, int32) int32)(m, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v22
						return v26
					}
				} else {
					v26 = v5
					return v26
				}
			}
		} else {
			if base.Ui32(int32(41)) < base.Ui32(l0) {
				v26 = int32(1)
				return v26
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(28))+uint32(_c_F_pg_encoding_mblen_or_incomplete[0])))
				v22 = m.T0[v21].(func(*base.Module, int32) int32)(m, l1)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v22
					return v26
				}
			}
		}
	}
}
