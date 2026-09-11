package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_network_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 float64
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 float64
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 float64
	_ = v122
	if l1 <= int32(828) {
		if l1 == int32(650) {
			v53 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return float64(0)
			} else {
				v57 = int32(1)
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v59&v57 != 0 {
					v62 = v57
				} else {
					v62 = int32(4)
				}
				v63 = v53 + v62
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
				v66 = float64(256)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
				v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
				v85 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v64), v66), base.F64_convert_i32_u(v68)), v66), base.F64_convert_i32_u(v73)), v66), base.F64_convert_i32_u(v78)), v66), base.F64_convert_i32_u(v83))
				if v64 == int32(2) {
					v122 = v85
					return v122
				} else {
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
					return base.F64_add(base.F64_mul(v85, float64(256)), base.F64_convert_i32_u(v90))
				}
			}
		} else {
			if l1 != int32(774) {
				v118 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v118)
				v122 = float64(0)
				return v122
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v12 = int32(24)
				v14 = int32(_a_F_convert_network_to_scalar_0)
				v16 = int32(8)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				return base.F64_add(base.F64_mul(base.F64_convert_i32_s(v11<<(uint(v12)%32)|v11&v14<<(uint(v16)%32)|(int32(base.Ui32(v11)>>(uint(v16)%32))&v14|int32(base.Ui32(v11)>>(uint(v12)%32)))), float64(4.294967296e+09)), base.F64_convert_i32_s(v30<<(uint(v12)%32)|v30&v14<<(uint(v16)%32)|(int32(base.Ui32(v30)>>(uint(v16)%32))&v14|int32(base.Ui32(v30)>>(uint(v12)%32)))))
			}
		}
	} else {
		if l1 == int32(829) {
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v96 = int32(8)
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v99 = int32(16)
			v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
			return base.F64_add(base.F64_mul(base.F64_convert_i32_u(v94|(v95<<(uint(v96)%32)|v98<<(uint(v99)%32))), float64(1.6777216e+07)), base.F64_convert_i32_u(v106|(v107<<(uint(v96)%32)|v110<<(uint(v99)%32))))
		} else {
			if l1 != int32(869) {
				v118 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v118)
				v122 = float64(0)
				return v122
			} else {
				v53 = F_pg_detoast_datum_packed(m, l0)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return float64(0)
				} else {
					v57 = int32(1)
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
					if v59&v57 != 0 {
						v62 = v57
					} else {
						v62 = int32(4)
					}
					v63 = v53 + v62
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
					v66 = float64(256)
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)))
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
					v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
					v85 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v64), v66), base.F64_convert_i32_u(v68)), v66), base.F64_convert_i32_u(v73)), v66), base.F64_convert_i32_u(v78)), v66), base.F64_convert_i32_u(v83))
					if v64 == int32(2) {
						v122 = v85
						return v122
					} else {
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
						return base.F64_add(base.F64_mul(v85, float64(256)), base.F64_convert_i32_u(v90))
					}
				}
			}
		}
	}
}
