package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_vector_add(m *base.Module, l0 int32) int64 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(v33)
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
func F_vector_le(m *base.Module, l0 int32) int64 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	if v19 < v20 {
		goto L18
	} else {
		goto L19
	}
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
	return int64(0)
L11:
	;
	return int64(1)
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
L18:
	;
	return int64(1)
L19:
	;
	goto L20
L20:
	;
	return base.I64_extend_i32_u(base.B2i32(v19 <= v20))
}
func F_vector_ne(m *base.Module, l0 int32) int64 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	if v19 < v20 {
		goto L15
	} else {
		goto L16
	}
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
	return int64(1)
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
L15:
	;
	return int64(1)
L16:
	;
	goto L17
L17:
	;
	return base.I64_extend_i32_u(base.B2i32(v20 < v19))
}
func F_vector_spherical_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 float32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v53 float32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 float32
	_ = v60
	var v62 float32
	_ = v62
	var v65 int32
	_ = v65
	var v67 float32
	_ = v67
	var v69 float32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 float32
	_ = v74
	var v76 float32
	_ = v76
	var v79 float32
	_ = v79
	var v81 float32
	_ = v81
	var v86 float32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v107 float32
	_ = v107
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v122 float32
	_ = v122
	var v125 int32
	_ = v125
	var v127 float32
	_ = v127
	var v129 float32
	_ = v129
	var v131 float32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v148 float32
	_ = v148
	var v172 float64
	_ = v172
	var v176 int64
	_ = v176
	var v181 int32
	_ = v181
	var v194 float64
	_ = v194
	var v205 float64
	_ = v205
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v224 float64
	_ = v224
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v236 float64
	_ = v236
	var v242 float64
	_ = v242
	var v246 float64
	_ = v246
	var v249 float64
	_ = v249
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v2 = int32(0)
	v12 = float32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v24 = F_pg_detoast_datum(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
			if v26 == v27 {
				v30 = base.I32_extend16_s(v26)
				if v30 <= int32(0) {
					v172 = float64(0)
				} else {
					v33 = int32(8)
					v34 = v24 + v33
					v36 = v19 + v33
					v37 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v26) {
						v42 = v37
						v51 = v2
						v53 = v12
						for {
							v56 = v42 << (uint(int32(2)) % 32)
							v58 = v56 | int32(12)
							v60 = *(*float32)(unsafe.Add(mBase, uint32(v36+v58)))
							v62 = *(*float32)(unsafe.Add(mBase, uint32(v34+v58)))
							v65 = v56 | int32(8)
							v67 = *(*float32)(unsafe.Add(mBase, uint32(v36+v65)))
							v69 = *(*float32)(unsafe.Add(mBase, uint32(v34+v65)))
							v71 = int32(4)
							v72 = v56 | v71
							v74 = *(*float32)(unsafe.Add(mBase, uint32(v36+v72)))
							v76 = *(*float32)(unsafe.Add(mBase, uint32(v34+v72)))
							v79 = *(*float32)(unsafe.Add(mBase, uint32(v36+v56)))
							v81 = *(*float32)(unsafe.Add(mBase, uint32(v34+v56)))
							v86 = base.F32_add(base.F32_mul(v60, v62), base.F32_add(base.F32_mul(v67, v69), base.F32_add(base.F32_mul(v74, v76), base.F32_add(base.F32_mul(v79, v81), v53))))
							v88 = v42 + v71
							v90 = v51 + v71
							if v90 != v30&int32(_a_F_vector_spherical_distance_0) {
								v42 = v88
								v51 = v90
								v53 = v86
								continue
							} else {
								break
							}
							break
						}
						if v26&int32(3) == int32(0) {
							v148 = v86
						} else {
							v96 = v88
							v107 = v86
							v111 = v96
							v121 = v2
							v122 = v107
							for {
								v125 = v111 << (uint(int32(2)) % 32)
								v127 = *(*float32)(unsafe.Add(mBase, uint32(v36+v125)))
								v129 = *(*float32)(unsafe.Add(mBase, uint32(v34+v125)))
								v131 = base.F32_add(base.F32_mul(v127, v129), v122)
								v132 = int32(1)
								v135 = v121 + v132
								if v135 != v30&int32(3) {
									v111 = v111 + v132
									v121 = v135
									v122 = v131
									continue
								} else {
									break
								}
								break
							}
							v148 = v131
						}
					} else {
						v96 = v37
						v107 = v12
						v111 = v96
						v121 = v2
						v122 = v107
						for {
							v125 = v111 << (uint(int32(2)) % 32)
							v127 = *(*float32)(unsafe.Add(mBase, uint32(v36+v125)))
							v129 = *(*float32)(unsafe.Add(mBase, uint32(v34+v125)))
							v131 = base.F32_add(base.F32_mul(v127, v129), v122)
							v132 = int32(1)
							v135 = v121 + v132
							if v135 != v30&int32(3) {
								v111 = v111 + v132
								v121 = v135
								v122 = v131
								continue
							} else {
								break
							}
							break
						}
						v148 = v131
					}
					if base.F32_gt(v148, float32(1)) != 0 {
						v172 = float64(1)
					} else {
						if base.F32_lt(v148, float32(-1)) == int32(0) {
							v172 = base.F64_promote_f32(v148)
						} else {
							v172 = float64(-1)
						}
					}
				}
				v176 = base.I64_reinterpret_f64(v172)
				v181 = base.I32_wrap_i64(int64(base.Ui64(v176)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072693248)) <= base.Ui32(v181) {
					if base.I32_wrap_i64(v176)|(v181-int32(1072693248)) == int32(0) {
						if int64(0) <= v176 {
							v194 = float64(0)
						} else {
							v194 = float64(3.141592653589793)
						}
						v249 = v194
					} else {
						v249 = base.F64_div(float64(0), base.F64_sub(v172, v172))
					}
				} else {
					if base.Ui32(v181) <= base.Ui32(int32(1071644671)) {
						if base.Ui32(v181) < base.Ui32(int32(1012924417)) {
							v246 = float64(1.5707963267948966)
							v249 = v246
						} else {
							v205 = F_R(m, base.F64_mul(v172, v172))
							mBase = m.M
							v249 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v172, v205)), v172), float64(1.5707963267948966))
						}
					} else {
						if v176 < int64(0) {
							v217 = base.F64_mul(base.F64_add(v172, float64(1)), float64(0.5))
							v218 = base.F64_sqrt(v217)
							v219 = F_R(m, v217)
							mBase = m.M
							v224 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v218, base.F64_add(base.F64_mul(v218, v219), float64(-6.123233995736766e-17))))
							v249 = base.F64_add(v224, v224)
						} else {
							v229 = base.F64_mul(base.F64_sub(float64(1), v172), float64(0.5))
							v230 = base.F64_sqrt(v229)
							v231 = F_R(m, v229)
							mBase = m.M
							v236 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v230) & int64(-4294967296))
							v242 = base.F64_add(base.F64_add(base.F64_mul(v230, v231), base.F64_div(base.F64_sub(v229, base.F64_mul(v236, v236)), base.F64_add(v230, v236))), v236)
							v246 = base.F64_add(v242, v242)
							v249 = v246
						}
					}
				}
				m.G0 = v16 + int32(16)
				return base.I64_reinterpret_f64(base.F64_div(v249, float64(3.141592653589793)))
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v260 = m.ExcPending
				if v260 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int64(0)
					} else {
						v264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+4)))
						v265 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v265
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v264
						F_errmsg(m, int32(_a_F_vector_spherical_distance_1), v16)
						mBase = m.M
						v270 = m.ExcPending
						if v270 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_vector_spherical_distance_2), int32(76), int32(_a_F_vector_spherical_distance_3))
							mBase = m.M
							v275 = m.ExcPending
							if v275 != 0 {
								return int64(0)
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
func F_vector_to_halfvec(m *base.Module, l0 int32) int64 {
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v81 = m.ExcPending
	if v81 != 0 {
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
	return base.I64_extend_i32_u(v33)
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
	v84 = m.ExcPending
	if v84 != 0 {
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
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_vector_to_halfvec_1), int32(92), int32(_a_F_vector_to_halfvec_2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
