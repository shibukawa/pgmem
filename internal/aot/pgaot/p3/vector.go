package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_vector_add(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 float32
	_ = v67
	var v69 float32
	_ = v69
	var v73 int32
	_ = v73
	var v76 float32
	_ = v76
	var v78 float32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 float32
	_ = v104
	var v106 float32
	_ = v106
	var v121 int32
	_ = v121
	var v133 float32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v23 == v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v13 + int32(16)
	return v33
L5:
	;
	v29 = F_mul_size(m, int32(4), base.I32_extend16_s(v23))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L27
	}
L8:
	;
	v31 = F_add_size(m, int32(8), v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v33 = F_palloc0(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v31 << (uint(int32(2)) % 32)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	if v39 <= int32(0) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v42 = int32(8)
	v43 = v16 + v42
	v45 = v21 + v42
	v47 = v33 + v42
	v48 = int32(0)
	if v39 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v121 = int32(0)
	goto L20
L13:
	;
	v54 = v48
	v62 = int32(0)
	goto L16
L14:
	;
	v91 = v48
	goto L15
L15:
	;
	v101 = v91 << (uint(int32(2)) % 32)
	v104 = *(*float32)(unsafe.Add(mBase, uint32(v101+v45)))
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v101+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v101))) = base.F32_add(v104, v106)
	goto L12
L16:
	;
	v63 = int32(2)
	v64 = v54 << (uint(v63) % 32)
	v67 = *(*float32)(unsafe.Add(mBase, uint32(v64+v45)))
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v64+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v64))) = base.F32_add(v67, v69)
	v73 = v64 | int32(4)
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v73+v45)))
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v73+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v73))) = base.F32_add(v76, v78)
	v82 = v54 + v63
	v84 = v62 + v63
	if v84 != v39&int32(_a_F_vector_add_0) {
		v54 = v82
		v62 = v84
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v39&int32(1) == int32(0) {
		goto L12
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v91 = v82
	goto L15
L20:
	;
	v133 = *(*float32)(unsafe.Add(mBase, uint32(v47+v121<<(uint(int32(2))%32))))
	if base.F32_ne(base.F32_abs(v133), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	v138 = v121 + int32(1)
	if v39 != v138 {
		v121 = v138
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L4
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v149
	F_errmsg(m, int32(_a_F_vector_add_1), v13)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_vector_add_2), int32(76), int32(_a_F_vector_add_3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 float32
	_ = v43
	var v45 float32
	_ = v45
	var v53 int32
	_ = v53
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v21 | base.B2i32(v19 <= v20)
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(2)) % 32)
	v43 = *(*float32)(unsafe.Add(mBase, uint32(v12+v25+v41)))
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v17+v25+v41)))
	if base.F32_lt(v43, v45) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	return int32(1)
L12:
	;
	goto L13
L13:
	;
	if base.F32_gt(v43, v45) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v30 + int32(1)
	if v53 == v22 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L10
L17:
	;
	v30 = v53
	goto L9
}
func F_vector_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 float32
	_ = v43
	var v45 float32
	_ = v45
	var v52 int32
	_ = v52
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v21 | base.B2i32(v20 < v19)
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(2)) % 32)
	v43 = *(*float32)(unsafe.Add(mBase, uint32(v12+v25+v41)))
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v17+v25+v41)))
	if base.F32_gt(v43, v45)|base.F32_lt(v43, v45) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int32(1)
L11:
	;
	v52 = v30 + int32(1)
	if v22 != v52 {
		v30 = v52
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L4
}
func F_vector_spherical_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 float32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 float32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float32
	_ = v58
	var v60 float32
	_ = v60
	var v63 int32
	_ = v63
	var v65 float32
	_ = v65
	var v67 float32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 float32
	_ = v72
	var v74 float32
	_ = v74
	var v77 float32
	_ = v77
	var v79 float32
	_ = v79
	var v84 float32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v105 float32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 float32
	_ = v119
	var v121 int32
	_ = v121
	var v123 float32
	_ = v123
	var v125 float32
	_ = v125
	var v127 float32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v144 float32
	_ = v144
	var v166 float64
	_ = v166
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v188 float64
	_ = v188
	var v199 float64
	_ = v199
	var v211 float64
	_ = v211
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v218 float64
	_ = v218
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v230 float64
	_ = v230
	var v236 float64
	_ = v236
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	v2 = int32(0)
	v12 = float32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
			if v25 == v26 {
				v29 = base.I32_extend16_s(v25)
				if v29 <= int32(0) {
					v166 = float64(0)
				} else {
					v32 = int32(8)
					v33 = v23 + v32
					v35 = v18 + v32
					v36 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v25) {
						v41 = v36
						v50 = v2
						v52 = v12
						for {
							v54 = v41 << (uint(int32(2)) % 32)
							v56 = v54 | int32(12)
							v58 = *(*float32)(unsafe.Add(mBase, uint32(v35+v56)))
							v60 = *(*float32)(unsafe.Add(mBase, uint32(v33+v56)))
							v63 = v54 | int32(8)
							v65 = *(*float32)(unsafe.Add(mBase, uint32(v35+v63)))
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v33+v63)))
							v69 = int32(4)
							v70 = v54 | v69
							v72 = *(*float32)(unsafe.Add(mBase, uint32(v35+v70)))
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v33+v70)))
							v77 = *(*float32)(unsafe.Add(mBase, uint32(v35+v54)))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v33+v54)))
							v84 = base.F32_add(base.F32_mul(v58, v60), base.F32_add(base.F32_mul(v65, v67), base.F32_add(base.F32_mul(v72, v74), base.F32_add(base.F32_mul(v77, v79), v52))))
							v86 = v41 + v69
							v88 = v50 + v69
							if v88 != v29&int32(_a_F_vector_spherical_distance_0) {
								v41 = v86
								v50 = v88
								v52 = v84
								continue
							} else {
								break
							}
							break
						}
						if v25&int32(3) == int32(0) {
							v144 = v84
						} else {
							v94 = v86
							v105 = v84
							v108 = v94
							v118 = v2
							v119 = v105
							for {
								v121 = v108 << (uint(int32(2)) % 32)
								v123 = *(*float32)(unsafe.Add(mBase, uint32(v35+v121)))
								v125 = *(*float32)(unsafe.Add(mBase, uint32(v33+v121)))
								v127 = base.F32_add(base.F32_mul(v123, v125), v119)
								v128 = int32(1)
								v131 = v118 + v128
								if v131 != v29&int32(3) {
									v108 = v108 + v128
									v118 = v131
									v119 = v127
									continue
								} else {
									break
								}
								break
							}
							v144 = v127
						}
					} else {
						v94 = v36
						v105 = v12
						v108 = v94
						v118 = v2
						v119 = v105
						for {
							v121 = v108 << (uint(int32(2)) % 32)
							v123 = *(*float32)(unsafe.Add(mBase, uint32(v35+v121)))
							v125 = *(*float32)(unsafe.Add(mBase, uint32(v33+v121)))
							v127 = base.F32_add(base.F32_mul(v123, v125), v119)
							v128 = int32(1)
							v131 = v118 + v128
							if v131 != v29&int32(3) {
								v108 = v108 + v128
								v118 = v131
								v119 = v127
								continue
							} else {
								break
							}
							break
						}
						v144 = v127
					}
					if base.F32_gt(v144, float32(1)) != 0 {
						v166 = float64(1)
					} else {
						if base.F32_lt(v144, float32(-1)) == int32(0) {
							v166 = base.F64_promote_f32(v144)
						} else {
							v166 = float64(-1)
						}
					}
				}
				v170 = base.I64_reinterpret_f64(v166)
				v175 = base.I32_wrap_i64(int64(base.Ui64(v170)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072693248)) <= base.Ui32(v175) {
					if base.I32_wrap_i64(v170)|(v175-int32(1072693248)) == int32(0) {
						if int64(0) <= v170 {
							v188 = float64(0)
						} else {
							v188 = float64(3.141592653589793)
						}
						v243 = v188
					} else {
						v243 = base.F64_div(float64(0), base.F64_sub(v166, v166))
					}
				} else {
					if base.Ui32(v175) <= base.Ui32(int32(1071644671)) {
						if base.Ui32(v175) < base.Ui32(int32(1012924417)) {
							v240 = float64(1.5707963267948966)
							v243 = v240
						} else {
							v199 = F_R(m, base.F64_mul(v166, v166))
							mBase = m.M
							v243 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v166, v199)), v166), float64(1.5707963267948966))
						}
					} else {
						if v170 < int64(0) {
							v211 = base.F64_mul(base.F64_add(v166, float64(1)), float64(0.5))
							v212 = base.F64_sqrt(v211)
							v213 = F_R(m, v211)
							mBase = m.M
							v218 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v212, base.F64_add(base.F64_mul(v212, v213), float64(-6.123233995736766e-17))))
							v243 = base.F64_add(v218, v218)
						} else {
							v223 = base.F64_mul(base.F64_sub(float64(1), v166), float64(0.5))
							v224 = base.F64_sqrt(v223)
							v225 = F_R(m, v223)
							mBase = m.M
							v230 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v224) & int64(-4294967296))
							v236 = base.F64_add(base.F64_add(base.F64_mul(v224, v225), base.F64_div(base.F64_sub(v223, base.F64_mul(v230, v230)), base.F64_add(v224, v230))), v230)
							v240 = base.F64_add(v236, v236)
							v243 = v240
						}
					}
				}
				v246 = F_Float8GetDatum(m, base.F64_div(v243, float64(3.141592653589793)))
				mBase = m.M
				v247 = m.ExcPending
				if v247 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v246
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v258 = m.ExcPending
					if v258 != 0 {
						return int32(0)
					} else {
						v259 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v260
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v259
						F_errmsg(m, int32(_a_F_vector_spherical_distance_1), v15)
						mBase = m.M
						v265 = m.ExcPending
						if v265 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_vector_spherical_distance_2), int32(76), int32(_a_F_vector_spherical_distance_3))
							mBase = m.M
							v270 = m.ExcPending
							if v270 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_vector_to_halfvec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v59 float32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	F_CheckDim_1(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if base.B2i32(v16 != int32(-1))&base.B2i32(v16 != v22) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = F_mul_size(m, int32(2), v22)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v31 = F_add_size(m, int32(8), v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = F_palloc0(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v22)
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v31 << (uint(int32(2)) % 32)
	v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if int32(0) < v39 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = int32(8)
	v47 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v9 + int32(16)
	return v33
L13:
	;
	v59 = *(*float32)(unsafe.Add(mBase, uint32(v12+v42+v47<<(uint(int32(2))%32))))
	v60 = F_Float4ToHalf(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v33+v42+v47<<(uint(int32(1))%32)))) = uint16(v60)
	v64 = v47 + int32(1)
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	if v64 < v65 {
		v47 = v64
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
	F_errmsg(m, int32(_a_F_vector_to_halfvec_0), v9)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_vector_to_halfvec_1), int32(92), int32(_a_F_vector_to_halfvec_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
