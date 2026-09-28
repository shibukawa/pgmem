package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_VectorItemSize(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = F_mul_size(m, int32(4), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_add_size(m, int32(8), v4)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_VectorUpdateCenter(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 float32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 float32
	_ = v50
	var v53 int32
	_ = v53
	var v56 float32
	_ = v56
	var v59 int32
	_ = v59
	var v62 float32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 float32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v2 = l1
	v4 = int32(0)
	v12 = F_mul_size(m, int32(4), v2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = F_add_size(m, int32(8), v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v14 << (uint(int32(2)) % 32)
			if v2 <= int32(0) {
			} else {
				v23 = v2 & int32(3)
				v25 = l0 + int32(8)
				v26 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v2) {
					v31 = v26
					v38 = v4
					for {
						v41 = v31 << (uint(int32(2)) % 32)
						v44 = *(*float32)(unsafe.Add(mBase, uint32(v41+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v41))) = v44
						v46 = int32(4)
						v47 = v41 | v46
						v50 = *(*float32)(unsafe.Add(mBase, uint32(l2+v47)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v47))) = v50
						v53 = v41 | int32(8)
						v56 = *(*float32)(unsafe.Add(mBase, uint32(l2+v53)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v53))) = v56
						v59 = v41 | int32(12)
						v62 = *(*float32)(unsafe.Add(mBase, uint32(v59+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v59))) = v62
						v65 = v31 + v46
						v67 = v38 + v46
						if v67 != v2&int32(2147483644) {
							v31 = v65
							v38 = v67
							continue
						} else {
							break
						}
						break
					}
					if v23 == int32(0) {
					} else {
						v71 = v65
						v80 = v71
						v88 = v4
						for {
							v90 = v80 << (uint(int32(2)) % 32)
							v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+l2)))
							*(*float32)(unsafe.Add(mBase, uint32(v25+v90))) = v93
							v95 = int32(1)
							v98 = v88 + v95
							if v98 != v23 {
								v80 = v80 + v95
								v88 = v98
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v71 = v26
					v80 = v71
					v88 = v4
					for {
						v90 = v80 << (uint(int32(2)) % 32)
						v93 = *(*float32)(unsafe.Add(mBase, uint32(v90+l2)))
						*(*float32)(unsafe.Add(mBase, uint32(v25+v90))) = v93
						v95 = int32(1)
						v98 = v88 + v95
						if v98 != v23 {
							v80 = v80 + v95
							v88 = v98
							continue
						} else {
							break
						}
						break
					}
				}
			}
			return
		}
	}
}
func F_vector(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14298(m, l0, int32(88), int32(_a_F_vector_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_vector_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v79 float32
	_ = v79
	var v81 float64
	_ = v81
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 float32
	_ = v119
	var v123 int32
	_ = v123
	var v130 float32
	_ = v130
	var v134 int32
	_ = v134
	var v141 float32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 float32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 float32
	_ = v192
	var v196 int32
	_ = v196
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v25 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L46
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L42
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v28 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v31 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v32 != int32(701) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
	v37 = v28 - int32(1)
	if v37 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v42 = int32(8)
	v43 = v23 + v42
	v45 = v18 + int32(24)
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
	v49 = v41 + int32(1)
	v50 = F_palloc_mul(m, v42, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v41 = v35
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v37 != v35 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v41 = v37
	goto L11
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v50))) = base.F64_add(v46, float64(1))
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v214 = F_construct_array(m, v50, v49, int32(701), int32(8), int32(1), int32(100))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L37
	}
L18:
	;
	v55 = int32(0)
	if v41 <= v55 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v41 <= int32(0) {
		goto L17
	} else {
		goto L26
	}
L21:
	;
	v59 = v55
	goto L22
L22:
	;
	v73 = v59 + int32(1)
	v75 = v73 << (uint(int32(3)) % 32)
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v45+v75)))
	v79 = *(*float32)(unsafe.Add(mBase, uint32(v43+v59<<(uint(int32(2))%32))))
	v81 = base.F64_add(v77, base.F64_promote_f32(v79))
	if base.F64_eq(base.F64_abs(v81), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L17
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v50+v75))) = v81
	if v41 != v73 {
		v59 = v73
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v91 = v41 & int32(3)
	v92 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v41) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = int32(0)
	v100 = v92
	goto L30
L28:
	;
	v161 = v92
	goto L29
L29:
	;
	v173 = v161
	v178 = v92
	goto L34
L30:
	;
	v112 = v100 | int32(1)
	v113 = int32(3)
	v116 = int32(2)
	v119 = *(*float32)(unsafe.Add(mBase, uint32(v43+v100<<(uint(v116)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v50+v112<<(uint(v113)%32)))) = base.F64_promote_f32(v119)
	v123 = v100 | v116
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v43+v112<<(uint(v116)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v50+v123<<(uint(v113)%32)))) = base.F64_promote_f32(v130)
	v134 = v100 | v113
	v141 = *(*float32)(unsafe.Add(mBase, uint32(v43+v123<<(uint(v116)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v50+v134<<(uint(v113)%32)))) = base.F64_promote_f32(v141)
	v144 = int32(4)
	v145 = v100 + v144
	v152 = *(*float32)(unsafe.Add(mBase, uint32(v43+v134<<(uint(v116)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v50+v145<<(uint(v113)%32)))) = base.F64_promote_f32(v152)
	v156 = v99 + v144
	if v156 != v41&int32(2147483644) {
		v99 = v156
		v100 = v145
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v91 == int32(0) {
		goto L17
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v161 = v145
	goto L29
L34:
	;
	v184 = int32(1)
	v185 = v173 + v184
	v192 = *(*float32)(unsafe.Add(mBase, uint32(v43+v173<<(uint(int32(2))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v50+v185<<(uint(int32(3))%32)))) = base.F64_promote_f32(v192)
	v196 = v178 + v184
	if v196 != v91 {
		v173 = v185
		v178 = v196
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L17
L36:
	;
	goto L35
L37:
	;
	F_pfree(m, v50)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	m.G0 = v15 + int32(32)
	return base.I64_extend_i32_u(v214)
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(_a_F_vector_accum_0)
	F_errmsg_internal(m, int32(_a_F_vector_accum_1), v15)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_vector_accum_2), int32(169), int32(_a_F_vector_accum_3))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v37
	F_errmsg(m, int32(_a_F_vector_accum_4), v15+int32(16))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_vector_accum_2), int32(88), int32(_a_F_vector_accum_5))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v50 int32
	_ = v50
	var v51 float64
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v99 int64
	_ = v99
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int64
	_ = v179
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int64
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 float64
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v258 float64
	_ = v258
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int64
	_ = v297
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v369 float64
	_ = v369
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v395 float64
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v26 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v383))) = v395
	v401 = F_construct_array(m, v383, v384, int32(701), int32(8), int32(1), int32(100))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L71
	}
L5:
	;
	v383 = v60
	v384 = v39
	v395 = v48
	goto L4
L6:
	;
	v383 = v357
	v384 = v29
	v395 = v369
	goto L4
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L70
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L66
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L63
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L60
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v29 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v32 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v33 != int32(701) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v36 != int32(1) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v39 <= int32(0) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v42 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v43 != int32(701) {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v46 = int32(24)
	v47 = v24 + v46
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v24)+24))
	v50 = v19 + v46
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v50)))
	if base.F64_eq(v51, float64(0)) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v279 = v66
	v282 = v269
	goto L57
L20:
	;
	v55 = v39 - int32(1)
	F_CheckDim_3(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.F64_eq(v48, float64(0)) != 0 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	v60 = F_palloc_mul(m, int32(8), v39)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v39 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v383 = v60
	v384 = int32(1)
	v395 = v48
	goto L4
L26:
	;
	goto L27
L27:
	;
	v64 = int32(3)
	v65 = v55 & v64
	v66 = int32(0)
	if base.Ui32(v39-int32(2)) < base.Ui32(v64) {
		v269 = v66
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v78 = v66
	v82 = v2
	goto L29
L29:
	;
	v87 = int32(3)
	v88 = v78 << (uint(v87) % 32)
	v90 = v88 | int32(8)
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v47+v90)))
	*(*int64)(unsafe.Add(mBase, uint32(v60+v90))) = v93
	v96 = v88 | int32(16)
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v47+v96)))
	*(*int64)(unsafe.Add(mBase, uint32(v60+v96))) = v99
	v102 = v88 | int32(24)
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v47+v102)))
	*(*int64)(unsafe.Add(mBase, uint32(v60+v102))) = v105
	v107 = int32(4)
	v108 = v78 + v107
	v110 = v108 << (uint(v87) % 32)
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v47+v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v60+v110))) = v113
	v116 = v82 + v107
	if v116 != v55&int32(-4) {
		v78 = v108
		v82 = v116
		goto L29
	} else {
		goto L31
	}
L30:
	;
	if v65 != 0 {
		v269 = v108
		goto L19
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L5
L33:
	;
	v121 = v29 - int32(1)
	F_CheckDim_3(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v224 = v29 - int32(1)
	F_CheckDim_3(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L49
	}
L36:
	;
	v126 = F_palloc_mul(m, int32(8), v29)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v29 == int32(1) {
		v383 = v126
		v384 = int32(1)
		v395 = v51
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v130 = int32(3)
	v131 = v121 & v130
	v132 = int32(0)
	if base.Ui32(v130) <= base.Ui32(v29-int32(2)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v145 = v132
	v148 = v2
	goto L42
L40:
	;
	v191 = v132
	goto L41
L41:
	;
	v200 = v132
	v204 = v191
	goto L46
L42:
	;
	v153 = int32(3)
	v154 = v145 << (uint(v153) % 32)
	v156 = v154 | int32(8)
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v50+v156)))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v156))) = v159
	v162 = v154 | int32(16)
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v50+v162)))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v162))) = v165
	v168 = v154 | int32(24)
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v168+v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v168))) = v171
	v173 = int32(4)
	v174 = v145 + v173
	v176 = v174 << (uint(v153) % 32)
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v176+v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v176))) = v179
	v182 = v148 + v173
	if v182 != v121&int32(-4) {
		v145 = v174
		v148 = v182
		goto L42
	} else {
		goto L44
	}
L43:
	;
	if v131 == int32(0) {
		v357 = v126
		v369 = v51
		goto L6
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v191 = v174
	goto L41
L46:
	;
	v212 = int32(1)
	v213 = v204 + v212
	v215 = v213 << (uint(int32(3)) % 32)
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v215+v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v126+v215))) = v218
	v221 = v200 + v212
	if v221 != v131 {
		v200 = v221
		v204 = v213
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v357 = v126
	v369 = v51
	goto L6
L48:
	;
	goto L47
L49:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v29 != v227 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	v229 = base.F64_add(v48, v51)
	v232 = F_palloc_mul(m, int32(8), v29)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v29 == int32(1) {
		v383 = v232
		v384 = int32(1)
		v395 = v229
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v238 = int32(0)
	goto L53
L53:
	;
	v251 = v238 + int32(1)
	v253 = v251 << (uint(int32(3)) % 32)
	v255 = *(*float64)(unsafe.Add(mBase, uint32(v47+v253)))
	v257 = *(*float64)(unsafe.Add(mBase, uint32(v50+v253)))
	v258 = base.F64_add(v255, v257)
	if base.F64_eq(base.F64_abs(v258), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L55
	}
L54:
	;
	v357 = v232
	v369 = v229
	goto L6
L55:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v232+v253))) = v258
	if v251 != v224 {
		v238 = v251
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v291 = int32(1)
	v292 = v282 + v291
	v294 = v292 << (uint(int32(3)) % 32)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v47+v294)))
	*(*int64)(unsafe.Add(mBase, uint32(v60+v294))) = v297
	v300 = v279 + v291
	if v300 != v65 {
		v279 = v300
		v282 = v292
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L5
L59:
	;
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_vector_combine_0)
	F_errmsg_internal(m, int32(_a_F_vector_combine_1), v16)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_vector_combine_2), int32(169), int32(_a_F_vector_combine_3))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(_a_F_vector_combine_0)
	F_errmsg_internal(m, int32(_a_F_vector_combine_1), v16+int32(16))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_vector_combine_2), int32(169), int32(_a_F_vector_combine_3))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v227 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v224
	F_errmsg(m, int32(_a_F_vector_combine_4), v16+int32(32))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_vector_combine_2), int32(88), int32(_a_F_vector_combine_5))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_pfree(m, v383)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	m.G0 = v16 + int32(48)
	return base.I64_extend_i32_u(v401)
}
