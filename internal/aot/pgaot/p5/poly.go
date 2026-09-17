package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_overright(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_le(v14, v15)
						}
					} else {
						return base.F64_le(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_le(v14, v15)
					}
				} else {
					return base.F64_le(v14, v15)
				}
			}
		}
	}
}
