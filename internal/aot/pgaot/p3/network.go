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
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 float64
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 float64
	_ = v110
	if l1 <= int32(828) {
		if l1 == int32(650) {
			v41 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return float64(0)
			} else {
				v45 = int32(1)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
				if v47&v45 != 0 {
					v50 = v45
				} else {
					v50 = int32(4)
				}
				v51 = v41 + v50
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				v54 = float64(256)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+2)))
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+3)))
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+5)))
				v73 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v52), v54), base.F64_convert_i32_u(v56)), v54), base.F64_convert_i32_u(v61)), v54), base.F64_convert_i32_u(v66)), v54), base.F64_convert_i32_u(v71))
				if v52 == int32(2) {
					v110 = v73
					return v110
				} else {
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+6)))
					return base.F64_add(base.F64_mul(v73, float64(256)), base.F64_convert_i32_u(v78))
				}
			}
		} else {
			if l1 != int32(774) {
				v106 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v106)
				v110 = float64(0)
				return v110
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v12 = int32(16711935)
				v14 = int32(8)
				v16 = int32(24)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				return base.F64_add(base.F64_mul(base.F64_convert_i32_s(base.I32_rotr(v11&v12, v14)|base.I32_rotr(v11, v16)&v12), float64(4.294967296e+09)), base.F64_convert_i32_s(base.I32_rotr(v24&v12, v14)|base.I32_rotr(v24, v16)&v12))
			}
		}
	} else {
		if l1 == int32(829) {
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v84 = int32(8)
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			v87 = int32(16)
			v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
			return base.F64_add(base.F64_mul(base.F64_convert_i32_u(v82|(v83<<(uint(v84)%32)|v86<<(uint(v87)%32))), float64(1.6777216e+07)), base.F64_convert_i32_u(v94|(v95<<(uint(v84)%32)|v98<<(uint(v87)%32))))
		} else {
			if l1 != int32(869) {
				v106 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v106)
				v110 = float64(0)
				return v110
			} else {
				v41 = F_pg_detoast_datum_packed(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return float64(0)
				} else {
					v45 = int32(1)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
					if v47&v45 != 0 {
						v50 = v45
					} else {
						v50 = int32(4)
					}
					v51 = v41 + v50
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					v54 = float64(256)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+2)))
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+3)))
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+5)))
					v73 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v52), v54), base.F64_convert_i32_u(v56)), v54), base.F64_convert_i32_u(v61)), v54), base.F64_convert_i32_u(v66)), v54), base.F64_convert_i32_u(v71))
					if v52 == int32(2) {
						v110 = v73
						return v110
					} else {
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+6)))
						return base.F64_add(base.F64_mul(v73, float64(256)), base.F64_convert_i32_u(v78))
					}
				}
			}
		}
	}
}
