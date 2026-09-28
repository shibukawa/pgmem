package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_try_partial_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v42 float64
	_ = v42
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v63 int32
	_ = v63
	var v69 float64
	_ = v69
	var v73 float64
	_ = v73
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v139 int64
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v15 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		if v16 != 0 {
			m.G0 = v13 + int32(96)
			return
		} else {
			v27 = m.G0
			v29 = v27 - int32(16)
			m.G0 = v29
			v31 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
			v34 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[0]))
			if l4 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
				v38 = base.F64_convert_i32_s(v35)
			} else {
				v38 = float64(0)
			}
			v39 = base.F64_mul(v34, v38)
			v42 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[1]))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v46 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
			v47 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
			v49 = float64(0)
			v51 = base.F64_add(base.F64_mul(v39, v32), base.F64_add(base.F64_sub(v46, v47), v49))
			v54 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
			v56 = base.F64_add(base.F64_mul(base.F64_add(v39, v42), v31), base.F64_add(base.F64_add(v47, v49), v54))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
			v59 = *(*int64)(unsafe.Add(mBase, uint32(l6)+40))
			if l7 != 0 {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
				v61 = base.F64_convert_i32_s(v60)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[2])))
				if v63 == int32(1) {
					v69 = base.F64_add(base.F64_mul(v61, float64(-0.3)), float64(1))
					if base.F64_gt(v69, float64(0)) != 0 {
						v73 = v69
					} else {
						v73 = math.Float64frombits(uint64(0x8000000000000000))
					}
					v76 = base.F64_add(v73, v61)
				} else {
					v76 = v61
				}
				v78 = base.F64_mul(v31, v76)
			} else {
				v78 = v31
			}
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
			F_ExecChooseHashTableSize(m, v78, v81, int32(1), l7, v45, v29, v29+int32(12), v29+int32(8), v29+int32(4))
			mBase = m.M
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
			if int32(2) <= v90 {
				v94 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[3]))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
				v97 = int32(7)
				v99 = int32(-8)
				v101 = int32(24)
				v105 = float64(0.0001220703125)
				v107 = base.F64_ceil(base.F64_mul(base.F64_mul(v32, base.F64_convert_i32_u((v96+v97)&v99+v101)), v105))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
				v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
				v121 = base.F64_ceil(base.F64_mul(base.F64_mul(v31, base.F64_convert_i32_u((v110+v97)&v99+v101)), v105))
				v127 = base.F64_add(base.F64_mul(v94, v121), v56)
				v129 = base.F64_add(base.F64_mul(v94, base.F64_add(base.F64_add(v107, v107), v121)), v51)
			} else {
				v127 = v56
				v129 = v51
			}
			*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v129
			*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v127
			*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_add(v129, v127)
			if v45 != 0 {
				v139 = int64(-2049)
			} else {
				v139 = int64(-264193)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57 + v58 + base.B2i32(v139|v59 != int64(-1))
			v145 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+88)) = v78
			*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v90
			*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v145
			m.G0 = v29 + int32(16)
			v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v153 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
			v154 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
			v155 = int32(0)
			v156 = F_add_partial_path_precheck(m, l1, v152, v153, v154, v155)
			mBase = m.M
			if v156 == v155 {
				m.G0 = v13 + int32(96)
				return
			} else {
				v159 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
				v161 = F_create_hashjoin_path(m, l0, l1, l5, v13, l6, l2, l3, l7, v159, int32(0), l4)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return
				} else {
					F_add_partial_path(m, l1, v161)
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						m.G0 = v13 + int32(96)
						return
					}
				}
			}
		}
	} else {
		v27 = m.G0
		v29 = v27 - int32(16)
		m.G0 = v29
		v31 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v34 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[0]))
		if l4 != 0 {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
			v38 = base.F64_convert_i32_s(v35)
		} else {
			v38 = float64(0)
		}
		v39 = base.F64_mul(v34, v38)
		v42 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[1]))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v46 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v47 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
		v49 = float64(0)
		v51 = base.F64_add(base.F64_mul(v39, v32), base.F64_add(base.F64_sub(v46, v47), v49))
		v54 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
		v56 = base.F64_add(base.F64_mul(base.F64_add(v39, v42), v31), base.F64_add(base.F64_add(v47, v49), v54))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
		v59 = *(*int64)(unsafe.Add(mBase, uint32(l6)+40))
		if l7 != 0 {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
			v61 = base.F64_convert_i32_s(v60)
			v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[2])))
			if v63 == int32(1) {
				v69 = base.F64_add(base.F64_mul(v61, float64(-0.3)), float64(1))
				if base.F64_gt(v69, float64(0)) != 0 {
					v73 = v69
				} else {
					v73 = math.Float64frombits(uint64(0x8000000000000000))
				}
				v76 = base.F64_add(v73, v61)
			} else {
				v76 = v61
			}
			v78 = base.F64_mul(v31, v76)
		} else {
			v78 = v31
		}
		v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
		v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
		F_ExecChooseHashTableSize(m, v78, v81, int32(1), l7, v45, v29, v29+int32(12), v29+int32(8), v29+int32(4))
		mBase = m.M
		v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
		if int32(2) <= v90 {
			v94 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[3]))
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
			v97 = int32(7)
			v99 = int32(-8)
			v101 = int32(24)
			v105 = float64(0.0001220703125)
			v107 = base.F64_ceil(base.F64_mul(base.F64_mul(v32, base.F64_convert_i32_u((v96+v97)&v99+v101)), v105))
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
			v121 = base.F64_ceil(base.F64_mul(base.F64_mul(v31, base.F64_convert_i32_u((v110+v97)&v99+v101)), v105))
			v127 = base.F64_add(base.F64_mul(v94, v121), v56)
			v129 = base.F64_add(base.F64_mul(v94, base.F64_add(base.F64_add(v107, v107), v121)), v51)
		} else {
			v127 = v56
			v129 = v51
		}
		*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v129
		*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v127
		*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_add(v129, v127)
		if v45 != 0 {
			v139 = int64(-2049)
		} else {
			v139 = int64(-264193)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57 + v58 + base.B2i32(v139|v59 != int64(-1))
		v145 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
		*(*float64)(unsafe.Add(mBase, uint32(v13)+88)) = v78
		*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v90
		*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v145
		m.G0 = v29 + int32(16)
		v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v153 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
		v154 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
		v155 = int32(0)
		v156 = F_add_partial_path_precheck(m, l1, v152, v153, v154, v155)
		mBase = m.M
		if v156 == v155 {
			m.G0 = v13 + int32(96)
			return
		} else {
			v159 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
			v161 = F_create_hashjoin_path(m, l0, l1, l5, v13, l6, l2, l3, l7, v159, int32(0), l4)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				F_add_partial_path(m, l1, v161)
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return
				} else {
					m.G0 = v13 + int32(96)
					return
				}
			}
		}
	}
}
