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
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
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
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 float32
	_ = v102
	var v104 float32
	_ = v104
	var v110 int32
	_ = v110
	var v122 float32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
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
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L28
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
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v55 = v48
	v61 = int32(0)
	goto L15
L13:
	;
	v88 = v48
	goto L14
L14:
	;
	if v39&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v63 = int32(2)
	v64 = v55 << (uint(v63) % 32)
	v67 = *(*float32)(unsafe.Add(mBase, uint32(v64+v45)))
	v69 = *(*float32)(unsafe.Add(mBase, uint32(v64+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v64))) = base.F32_add(v67, v69)
	v73 = v64 | int32(4)
	v76 = *(*float32)(unsafe.Add(mBase, uint32(v73+v45)))
	v78 = *(*float32)(unsafe.Add(mBase, uint32(v73+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v73))) = base.F32_add(v76, v78)
	v82 = v55 + v63
	v84 = v61 + v63
	if v84 != v39&int32(32766) {
		v55 = v82
		v61 = v84
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v88 = v82
	goto L14
L17:
	;
	goto L16
L18:
	;
	v99 = v88 << (uint(int32(2)) % 32)
	v102 = *(*float32)(unsafe.Add(mBase, uint32(v99+v45)))
	v104 = *(*float32)(unsafe.Add(mBase, uint32(v99+v43)))
	*(*float32)(unsafe.Add(mBase, uint32(v47+v99))) = base.F32_add(v102, v104)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v110 = int32(0)
	goto L21
L21:
	;
	v122 = *(*float32)(unsafe.Add(mBase, uint32(v47+v110<<(uint(int32(2))%32))))
	if base.F32_ne(base.F32_abs(v122), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L27
	}
L23:
	;
	v127 = v110 + int32(1)
	if v39 != v127 {
		v110 = v127
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L4
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v138
	F_errmsg(m, int32(501556), v13)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(520291), int32(76), int32(160460))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 float32
	_ = v46
	var v48 float32
	_ = v48
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = base.B2i32(v21 < v22)
	if v21 < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v21 < v22 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L9
L9:
	;
	v44 = v32 << (uint(int32(2)) % 32)
	v46 = *(*float32)(unsafe.Add(mBase, uint32(v13+v27+v44)))
	v48 = *(*float32)(unsafe.Add(mBase, uint32(v44+(v18+v27))))
	if base.F32_lt(v46, v48) != 0 {
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
	if base.F32_gt(v46, v48) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = v32 + int32(1)
	if v56 == v24 {
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
	v32 = v56
	goto L9
L18:
	;
	v72 = int32(1)
	goto L20
L19:
	;
	v72 = base.B2i32(v21 <= v22)
	goto L20
L20:
	;
	return v72
}
func F_vector_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 float32
	_ = v46
	var v48 float32
	_ = v48
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = base.B2i32(v21 < v22)
	if v21 < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v21 < v22 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L9
L9:
	;
	v44 = v32 << (uint(int32(2)) % 32)
	v46 = *(*float32)(unsafe.Add(mBase, uint32(v13+v27+v44)))
	v48 = *(*float32)(unsafe.Add(mBase, uint32(v44+(v18+v27))))
	if base.F32_gt(v46, v48)|base.F32_lt(v46, v48) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int32(1)
L11:
	;
	v55 = v32 + int32(1)
	if v24 != v55 {
		v32 = v55
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
	v71 = int32(1)
	goto L17
L16:
	;
	v71 = base.B2i32(v22 < v21)
	goto L17
L17:
	;
	return v71
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
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 float32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 float32
	_ = v59
	var v61 float32
	_ = v61
	var v64 int32
	_ = v64
	var v66 float32
	_ = v66
	var v68 float32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 float32
	_ = v73
	var v75 float32
	_ = v75
	var v78 float32
	_ = v78
	var v80 float32
	_ = v80
	var v85 float32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v102 float32
	_ = v102
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 float32
	_ = v118
	var v120 int32
	_ = v120
	var v122 float32
	_ = v122
	var v124 float32
	_ = v124
	var v126 float32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v143 float32
	_ = v143
	var v165 float64
	_ = v165
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
	var v241 float64
	_ = v241
	var v245 float64
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
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
					v165 = float64(0)
				} else {
					v32 = int32(8)
					v33 = v23 + v32
					v35 = v18 + v32
					if base.Ui32(v25) < base.Ui32(int32(4)) {
						v91 = int32(0)
						v102 = v12
					} else {
						v42 = int32(0)
						v48 = v2
						v53 = v12
						for {
							v55 = v42 << (uint(int32(2)) % 32)
							v57 = v55 | int32(12)
							v59 = *(*float32)(unsafe.Add(mBase, uint32(v35+v57)))
							v61 = *(*float32)(unsafe.Add(mBase, uint32(v33+v57)))
							v64 = v55 | int32(8)
							v66 = *(*float32)(unsafe.Add(mBase, uint32(v35+v64)))
							v68 = *(*float32)(unsafe.Add(mBase, uint32(v33+v64)))
							v70 = int32(4)
							v71 = v55 | v70
							v73 = *(*float32)(unsafe.Add(mBase, uint32(v35+v71)))
							v75 = *(*float32)(unsafe.Add(mBase, uint32(v33+v71)))
							v78 = *(*float32)(unsafe.Add(mBase, uint32(v35+v55)))
							v80 = *(*float32)(unsafe.Add(mBase, uint32(v33+v55)))
							v85 = base.F32_add(base.F32_mul(v59, v61), base.F32_add(base.F32_mul(v66, v68), base.F32_add(base.F32_mul(v73, v75), base.F32_add(base.F32_mul(v78, v80), v53))))
							v87 = v42 + v70
							v89 = v48 + v70
							if v89 != v29&int32(32764) {
								v42 = v87
								v48 = v89
								v53 = v85
								continue
							} else {
								break
							}
							break
						}
						v91 = v87
						v102 = v85
					}
					if v25&int32(3) != 0 {
						v107 = v91
						v116 = v2
						v118 = v102
						for {
							v120 = v107 << (uint(int32(2)) % 32)
							v122 = *(*float32)(unsafe.Add(mBase, uint32(v35+v120)))
							v124 = *(*float32)(unsafe.Add(mBase, uint32(v33+v120)))
							v126 = base.F32_add(base.F32_mul(v122, v124), v118)
							v127 = int32(1)
							v130 = v116 + v127
							if v130 != v29&int32(3) {
								v107 = v107 + v127
								v116 = v130
								v118 = v126
								continue
							} else {
								break
							}
							break
						}
						v143 = v126
					} else {
						v143 = v102
					}
					if base.F32_gt(v143, float32(1)) != 0 {
						v165 = float64(1)
					} else {
						if base.F32_lt(v143, float32(-1)) == int32(0) {
							v165 = base.F64_promote_f32(v143)
						} else {
							v165 = float64(-1)
						}
					}
				}
				v170 = base.I64_reinterpret_f64(v165)
				v175 = base.I32_wrap_i64(int64(base.Ui64(v170)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1072693248)) <= base.Ui32(v175) {
					if base.I32_wrap_i64(v170)|(v175-int32(1072693248)) == int32(0) {
						if int64(0) <= v170 {
							v188 = float64(0)
						} else {
							v188 = float64(3.141592653589793)
						}
						v245 = v188
					} else {
						v245 = base.F64_div(float64(0), base.F64_sub(v165, v165))
					}
				} else {
					if base.Ui32(v175) <= base.Ui32(int32(1071644671)) {
						if base.Ui32(v175) < base.Ui32(int32(1012924417)) {
							v241 = float64(1.5707963267948966)
							v245 = v241
						} else {
							v199 = F_R(m, base.F64_mul(v165, v165))
							mBase = m.M
							v245 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v165, v199)), v165), float64(1.5707963267948966))
						}
					} else {
						if v170 < int64(0) {
							v211 = base.F64_mul(base.F64_add(v165, float64(1)), float64(0.5))
							v212 = base.F64_sqrt(v211)
							v213 = F_R(m, v211)
							mBase = m.M
							v218 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v212, base.F64_add(base.F64_mul(v212, v213), float64(-6.123233995736766e-17))))
							v245 = base.F64_add(v218, v218)
						} else {
							v223 = base.F64_mul(base.F64_sub(float64(1), v165), float64(0.5))
							v224 = base.F64_sqrt(v223)
							v225 = F_R(m, v223)
							mBase = m.M
							v230 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v224) & int64(-4294967296))
							v236 = base.F64_add(base.F64_add(base.F64_mul(v224, v225), base.F64_div(base.F64_sub(v223, base.F64_mul(v230, v230)), base.F64_add(v224, v230))), v230)
							v241 = base.F64_add(v236, v236)
							v245 = v241
						}
					}
				}
				v248 = F_Float8GetDatum(m, base.F64_div(v245, float64(3.141592653589793)))
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return int32(0)
				} else {
					m.G0 = v15 + int32(16)
					return v248
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v257 = m.ExcPending
				if v257 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v260 = m.ExcPending
					if v260 != 0 {
						return int32(0)
					} else {
						v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
						v262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v262
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v261
						F_errmsg(m, int32(501556), v15)
						mBase = m.M
						v267 = m.ExcPending
						if v267 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520291), int32(76), int32(160460))
							mBase = m.M
							v272 = m.ExcPending
							if v272 != 0 {
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
	F_errmsg(m, int32(490583), v9)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(526349), int32(92), int32(303246))
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
