package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_InitVector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v1 = l0
	v6 = F_mul_size(m, int32(4), v1)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_add_size(m, int32(8), v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_palloc0(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v10 << (uint(int32(2)) % 32)
				return v12
			}
		}
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 float32
	_ = v36
	var v38 float32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 float32
	_ = v44
	var v46 float32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 float32
	_ = v52
	var v54 float32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float32
	_ = v60
	var v62 float32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 float32
	_ = v95
	var v97 float32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v3 = int32(0)
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11 <= v3 {
	} else {
		v15 = v11 & int32(3)
		v17 = l0 + int32(8)
		v18 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(v11) {
			v23 = v18
			v28 = v3
			for {
				v34 = v23 << (uint(int32(2)) % 32)
				v35 = l1 + v34
				v36 = *(*float32)(unsafe.Add(mBase, uint32(v35)))
				v38 = *(*float32)(unsafe.Add(mBase, uint32(v34+v17)))
				*(*float32)(unsafe.Add(mBase, uint32(v35))) = base.F32_add(v36, v38)
				v41 = int32(4)
				v42 = v34 | v41
				v43 = l1 + v42
				v44 = *(*float32)(unsafe.Add(mBase, uint32(v43)))
				v46 = *(*float32)(unsafe.Add(mBase, uint32(v42+v17)))
				*(*float32)(unsafe.Add(mBase, uint32(v43))) = base.F32_add(v44, v46)
				v50 = v34 | int32(8)
				v51 = l1 + v50
				v52 = *(*float32)(unsafe.Add(mBase, uint32(v51)))
				v54 = *(*float32)(unsafe.Add(mBase, uint32(v50+v17)))
				*(*float32)(unsafe.Add(mBase, uint32(v51))) = base.F32_add(v52, v54)
				v58 = v34 | int32(12)
				v59 = l1 + v58
				v60 = *(*float32)(unsafe.Add(mBase, uint32(v59)))
				v62 = *(*float32)(unsafe.Add(mBase, uint32(v58+v17)))
				*(*float32)(unsafe.Add(mBase, uint32(v59))) = base.F32_add(v60, v62)
				v66 = v23 + v41
				v68 = v28 + v41
				if v68 != v11&int32(32764) {
					v23 = v66
					v28 = v68
					continue
				} else {
					break
				}
				break
			}
			v70 = v66
		} else {
			v70 = v18
		}
		if v15 == int32(0) {
		} else {
			v82 = v70
			v90 = v3
			for {
				v93 = v82 << (uint(int32(2)) % 32)
				v94 = l1 + v93
				v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)))
				v97 = *(*float32)(unsafe.Add(mBase, uint32(v93+v17)))
				*(*float32)(unsafe.Add(mBase, uint32(v94))) = base.F32_add(v95, v97)
				v100 = int32(1)
				v103 = v90 + v100
				if v103 != v15 {
					v82 = v82 + v100
					v90 = v103
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
	var v75 int32
	_ = v75
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
	return v75
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
	v47 = *(*float32)(unsafe.Add(mBase, uint32(v43+(v18+v26))))
	if base.F32_lt(v45, v47) != 0 {
		v75 = int32(0)
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
	v75 = base.B2i32(v21 < v20)
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
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	v9 = m.G0
	v11 = v9 - int32(64128)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = v14
	goto L1
L1:
	;
	v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15))))
	goto L3
L2:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v35 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v25 == int32(32))|base.B2i32(base.Ui32((v25-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v15 = v15 + int32(1)
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
	v348 = m.ExcPending
	if v348 != 0 {
		goto L25
	} else {
		goto L88
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L25
	} else {
		goto L83
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L25
	} else {
		goto L78
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L25
	} else {
		goto L74
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L25
	} else {
		goto L70
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L25
	} else {
		goto L66
	}
L11:
	;
	v40 = v15
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L25
	} else {
		goto L61
	}
L14:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+1)))
	v48 = v40 + int32(1)
	goto L16
L15:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v58 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v46 == int32(32))|base.B2i32(base.Ui32((v46-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v40 = v48
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v61 = v48
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
	*(*int32)(unsafe.Add(mBase, _consts[137])) = int32(0)
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
	v96 = *(*int32)(unsafe.Add(mBase, _consts[137]))
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
	if v132 != int32(16000) {
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
	F_errmsg(m, int32(753361), v11+int32(48))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(514923), int32(260), int32(290368))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(16000)
	F_errmsg(m, int32(154976), v11-int32(-64))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(514923), int32(214), int32(290368))
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
	v223 = v67<<(uint(v213)%32) + int32(4)
	if v223 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	m.G0 = v11 + int32(64128)
	return v210
L58:
	;
	v224 = F__emscripten_memcpy_bulkmem(m, v210+int32(8), v11+int32(128), v223)
	mBase = m.M
	goto L60
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v14
	F_errmsg(m, int32(753361), v11+int32(112))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(694148), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(514923), int32(194), int32(290368))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
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
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(282642), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(514923), int32(204), int32(290368))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
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
	F_errcode(m, int32(33685634))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	F_errmsg(m, int32(753361), v11)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(514923), int32(223), int32(290368))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L25
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L25
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
	F_errmsg(m, int32(753361), v11+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(514923), int32(234), int32(290368))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L25
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	v310 = F_pnstrdup(m, v61, v93-v61)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v310
	F_errmsg(m, int32(217703), v11+int32(32))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(514923), int32(240), int32(290368))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L25
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v14
	F_errmsg(m, int32(753361), v11+int32(96))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	F_errdetail(m, int32(667615), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(514923), int32(271), int32(290368))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L25
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L25
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v13
	F_errmsg(m, int32(485731), v11+int32(80))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(514923), int32(88), int32(300297))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_typmod_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 == int32(1) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v19 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(584384), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(514923), int32(361), int32(291007))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(16001)) <= base.Ui32(v19) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(16000)
								F_errmsg(m, int32(497779), v5)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(514923), int32(366), int32(291007))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v5 + int32(16)
						return v19
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(231932), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(514923), int32(356), int32(291007))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
