package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_InitVector(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13853(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VectorSumCenter(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 float32
	_ = v34
	var v36 float32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v44 float32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 float32
	_ = v50
	var v52 float32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float32
	_ = v58
	var v60 float32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 float32
	_ = v97
	var v99 float32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v3 = int32(0)
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11 <= v3 {
	} else {
		v15 = l0 + int32(8)
		v16 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(v11) {
			v21 = v16
			v29 = v3
			for {
				v32 = v21 << (uint(int32(2)) % 32)
				v33 = l1 + v32
				v34 = *(*float32)(unsafe.Add(mBase, uint32(v33)))
				v36 = *(*float32)(unsafe.Add(mBase, uint32(v32+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v33))) = base.F32_add(v34, v36)
				v39 = int32(4)
				v40 = v32 | v39
				v41 = l1 + v40
				v42 = *(*float32)(unsafe.Add(mBase, uint32(v41)))
				v44 = *(*float32)(unsafe.Add(mBase, uint32(v40+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v41))) = base.F32_add(v42, v44)
				v48 = v32 | int32(8)
				v49 = l1 + v48
				v50 = *(*float32)(unsafe.Add(mBase, uint32(v49)))
				v52 = *(*float32)(unsafe.Add(mBase, uint32(v48+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v49))) = base.F32_add(v50, v52)
				v56 = v32 | int32(12)
				v57 = l1 + v56
				v58 = *(*float32)(unsafe.Add(mBase, uint32(v57)))
				v60 = *(*float32)(unsafe.Add(mBase, uint32(v56+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v57))) = base.F32_add(v58, v60)
				v64 = v21 + v39
				v66 = v29 + v39
				if v66 != v11&int32(_a_F_VectorSumCenter_0) {
					v21 = v64
					v29 = v66
					continue
				} else {
					break
				}
				break
			}
			if v11&int32(3) == int32(0) {
			} else {
				v72 = v64
				v84 = v72
				v93 = v3
				for {
					v95 = v84 << (uint(int32(2)) % 32)
					v96 = l1 + v95
					v97 = *(*float32)(unsafe.Add(mBase, uint32(v96)))
					v99 = *(*float32)(unsafe.Add(mBase, uint32(v95+v15)))
					*(*float32)(unsafe.Add(mBase, uint32(v96))) = base.F32_add(v97, v99)
					v102 = int32(1)
					v105 = v93 + v102
					if v105 != v11&int32(3) {
						v84 = v84 + v102
						v93 = v105
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v72 = v16
			v84 = v72
			v93 = v3
			for {
				v95 = v84 << (uint(int32(2)) % 32)
				v96 = l1 + v95
				v97 = *(*float32)(unsafe.Add(mBase, uint32(v96)))
				v99 = *(*float32)(unsafe.Add(mBase, uint32(v95+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v96))) = base.F32_add(v97, v99)
				v102 = int32(1)
				v105 = v93 + v102
				if v105 != v11&int32(3) {
					v84 = v84 + v102
					v93 = v105
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
func F_vector_gt(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 float32
	_ = v45
	var v47 float32
	_ = v47
	var v53 int32
	_ = v53
	var v79 int32
	_ = v79
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
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v79
L5:
	;
	if v20 < v21 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v23 = v20
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	if v23 <= int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L10
L10:
	;
	v43 = v31 << (uint(int32(2)) % 32)
	v45 = *(*float32)(unsafe.Add(mBase, uint32(v13+v26+v43)))
	v47 = *(*float32)(unsafe.Add(mBase, uint32(v18+v26+v43)))
	if base.F32_lt(v45, v47) != 0 {
		v79 = int32(0)
		goto L4
	} else {
		goto L12
	}
L11:
	;
	return int32(1)
L12:
	;
	if base.F32_gt(v45, v47) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v53 = v31 + int32(1)
	if v53 == v23 {
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L11
L16:
	;
	v31 = v53
	goto L10
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v79 = base.B2i32(v21 < v20)
	goto L4
}
func F_vector_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v89 float32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	v9 = m.G0
	v11 = v9 - int32(_a_F_vector_in_0)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = v14
	goto L1
L1:
	;
	v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17))))
	goto L3
L2:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v35 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v25 == int32(32))|base.B2i32(base.Ui32((v25-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v17 = v17 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L25
	} else {
		goto L87
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L25
	} else {
		goto L82
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L25
	} else {
		goto L77
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L25
	} else {
		goto L73
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L25
	} else {
		goto L69
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L25
	} else {
		goto L65
	}
L11:
	;
	v38 = v17
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L25
	} else {
		goto L60
	}
L14:
	;
	v47 = v38 + int32(1)
	v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
	goto L16
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v58 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v48 == int32(32))|base.B2i32(base.Ui32((v48-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v38 = v47
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v61 = v47
	v67 = int32(0)
	goto L20
L19:
	;
	v177 = v120
	goto L47
L20:
	;
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L25
	} else {
		goto L43
	}
L22:
	;
	if base.B2i32(v71 == int32(32))|base.B2i32(base.Ui32((v71-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v61 = v61 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v81 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vector_in[0])) = int32(0)
	v89 = F_strtof(m, v61, v11+int32(124))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	if v93 == v61 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_vector_in[0]))
	if base.B2i32(v96 == int32(68))&base.F32_eq(base.F32_abs(v89), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	F_CheckElement_3(m, v89)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v11+int32(128)+v67<<(uint(int32(2))%32)))) = v89
	v111 = v93
	goto L30
L30:
	;
	v120 = v111 + int32(1)
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v111))))
	goto L32
L31:
	;
	v132 = v67 + int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v133 != int32(44) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if base.B2i32(v121 == int32(32))|base.B2i32(base.Ui32((v121-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v111 = v120
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v133 == int32(93) {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v132 != int32(_a_F_vector_in_1) {
		v61 = v120
		v67 = v132
		goto L20
	} else {
		goto L42
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v14
	F_errmsg(m, int32(_a_F_vector_in_2), v11+int32(48))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(260), int32(_a_F_vector_in_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L25
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
	goto L21
L43:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_vector_in_1)
	F_errmsg(m, int32(_a_F_vector_in_5), v11-int32(-64))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(214), int32(_a_F_vector_in_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v177))))
	goto L49
L48:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v197 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	if base.B2i32(v187 == int32(32))|base.B2i32(base.Ui32((v187-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v177 = v177 + int32(1)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_CheckDim_3(m, v132)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	if base.B2i32(v13 != int32(-1))&base.B2i32(v132 != v13) != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v206 = F_mul_size(m, int32(4), v132)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v208 = F_add_size(m, int32(8), v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	v210 = F_palloc0(m, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v210)+4)) = uint16(v132)
	v213 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v208 << (uint(v213) % 32)
	v219 = v67<<(uint(v213)%32) + int32(4)
	if v219 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v210+int32(8), v11+int32(128), v219)
	goto L59
L58:
	;
	goto L59
L59:
	;
	m.G0 = v11 + int32(_a_F_vector_in_0)
	return v210
L60:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L25
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v14
	F_errmsg(m, int32(_a_F_vector_in_2), v11+int32(112))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	F_errdetail(m, int32(_a_F_vector_in_6), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(194), int32(_a_F_vector_in_4))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_vector_in_7), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(204), int32(_a_F_vector_in_4))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	F_errmsg(m, int32(_a_F_vector_in_2), v11)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(223), int32(_a_F_vector_in_4))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L25
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L25
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
	F_errmsg(m, int32(_a_F_vector_in_2), v11+int32(16))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L25
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(234), int32(_a_F_vector_in_4))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v309 = F_pnstrdup(m, v61, v93-v61)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v309
	F_errmsg(m, int32(_a_F_vector_in_8), v11+int32(32))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(240), int32(_a_F_vector_in_4))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v14
	F_errmsg(m, int32(_a_F_vector_in_2), v11+int32(96))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	F_errdetail(m, int32(_a_F_vector_in_9), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(271), int32(_a_F_vector_in_4))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L25
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v13
	F_errmsg(m, int32(_a_F_vector_in_10), v11+int32(80))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L25
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(88), int32(_a_F_vector_in_11))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_typmod_in(m *base.Module, l0 int32) int32 {
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn13934(m, l0, int32(_a_F_vector_typmod_in_0), int32(366), int32(_a_F_vector_typmod_in_1), int32(_a_F_vector_typmod_in_2), int32(_a_F_vector_typmod_in_3), int32(361), int32(_a_F_vector_typmod_in_4), int32(356), int32(_a_F_vector_typmod_in_5))
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
