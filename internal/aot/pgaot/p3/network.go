package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_network_to_scalar(m *base.Module, l0 int64, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 float64
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 float64
	_ = v113
	if l1 <= int32(828) {
		if l1 == int32(650) {
			v43 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return float64(0)
			} else {
				v47 = int32(1)
				v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
				if v49&v47 != 0 {
					v52 = v47
				} else {
					v52 = int32(4)
				}
				v53 = v43 + v52
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				v56 = float64(256)
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)))
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)))
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
				v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+5)))
				v75 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v54), v56), base.F64_convert_i32_u(v58)), v56), base.F64_convert_i32_u(v63)), v56), base.F64_convert_i32_u(v68)), v56), base.F64_convert_i32_u(v73))
				if v54 == int32(2) {
					v113 = v75
					return v113
				} else {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+6)))
					return base.F64_add(base.F64_mul(v75, float64(256)), base.F64_convert_i32_u(v80))
				}
			}
		} else {
			if l1 != int32(774) {
				v109 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v109)
				v113 = float64(0)
				return v113
			} else {
				v11 = base.I32_wrap_i64(l0)
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v13 = int32(16711935)
				v15 = int32(8)
				v17 = int32(24)
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				return base.F64_add(base.F64_mul(base.F64_convert_i32_s(base.I32_rotr(v12&v13, v15)|base.I32_rotr(v12, v17)&v13), float64(4.294967296e+09)), base.F64_convert_i32_s(base.I32_rotr(v25&v13, v15)|base.I32_rotr(v25, v17)&v13))
			}
		}
	} else {
		if l1 == int32(829) {
			v84 = base.I32_wrap_i64(l0)
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
			v86 = int32(8)
			v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
			v89 = int32(16)
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+2)))
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+5)))
			v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+4)))
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+3)))
			return base.F64_add(base.F64_mul(base.F64_convert_i32_u(v85<<(uint(v86)%32)|v88<<(uint(v89)%32)|v92), float64(1.6777216e+07)), base.F64_convert_i32_u(v97|(v98<<(uint(v86)%32)|v101<<(uint(v89)%32))))
		} else {
			if l1 != int32(869) {
				v109 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v109)
				v113 = float64(0)
				return v113
			} else {
				v43 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return float64(0)
				} else {
					v47 = int32(1)
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
					if v49&v47 != 0 {
						v52 = v47
					} else {
						v52 = int32(4)
					}
					v53 = v43 + v52
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
					v56 = float64(256)
					v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)))
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)))
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+5)))
					v75 = base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_add(base.F64_mul(base.F64_convert_i32_u(v54), v56), base.F64_convert_i32_u(v58)), v56), base.F64_convert_i32_u(v63)), v56), base.F64_convert_i32_u(v68)), v56), base.F64_convert_i32_u(v73))
					if v54 == int32(2) {
						v113 = v75
						return v113
					} else {
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+6)))
						return base.F64_add(base.F64_mul(v75, float64(256)), base.F64_convert_i32_u(v80))
					}
				}
			}
		}
	}
}
