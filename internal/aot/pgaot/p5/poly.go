package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_overright(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+24))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_le(v17, v18)
						}
					} else {
						return base.F64_le(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_le(v17, v18)
					}
				} else {
					return base.F64_le(v17, v18)
				}
			}
		}
	}
}
