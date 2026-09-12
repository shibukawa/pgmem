package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_a_f8_f8(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v80 int32
	_ = v80
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v87 int32
	_ = v87
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v161 float64
	_ = v161
	var v164 int32
	_ = v164
	var v167 float64
	_ = v167
	var v170 int32
	_ = v170
	var v173 float64
	_ = v173
	var v176 int32
	_ = v176
	var v179 float64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v218 float64
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v271 float64
	_ = v271
	var v274 int32
	_ = v274
	var v282 float64
	_ = v282
	var v285 int32
	_ = v285
	var v293 float64
	_ = v293
	var v296 int32
	_ = v296
	var v304 float64
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v346 float64
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v26 = F_array_contains_nulls(m, v19)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L69
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L64
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L60
	}
L7:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v28 = F_array_contains_nulls(m, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v28 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v33 = F_ArrayGetNItems(m, v30, v19+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if int32(101) <= v33 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v40 = F_ArrayGetNItems(m, v37, v24+int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v40 != v33 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = v43
	goto L17
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v51 = (v44<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v52 = v51 + v19
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = v53
	goto L20
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v61 = (v54<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L20
L20:
	;
	v62 = v61 + v24
	v63 = int32(0)
	if v33 <= v63 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v125 = F_palloc0(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L30
	}
L22:
	;
	v122 = v2
	v124 = v33<<(uint(int32(3))%32) + int32(8)
	goto L21
L23:
	;
	v67 = v63
	goto L24
L24:
	;
	v80 = v67 << (uint(int32(3)) % 32)
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v52+v80)))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v80+v62)))
	if base.F64_eq(v82, v84) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v122 = int32(1)
	v124 = v33<<(uint(int32(4))%32) | int32(8)
	goto L21
L26:
	;
	v87 = v67 + int32(1)
	if v33 != v87 {
		v67 = v87
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L22
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v124 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v130&int32(-2147483648) | v33
	if int32(0) < v33 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	m.G0 = v16 + int32(16)
	return v125
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125)+4)) = v33 | int32(-2147483648)
	goto L31
L33:
	;
	v138 = v125 + int32(8)
	v139 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v122 != 0 {
		goto L31
	} else {
		goto L59
	}
L36:
	;
	v145 = v139
	v150 = v2
	goto L39
L37:
	;
	v187 = v139
	goto L38
L38:
	;
	v200 = v33 & int32(3)
	if v200 != 0 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v158 = v145 << (uint(int32(3)) % 32)
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v52+v158)))
	*(*float64)(unsafe.Add(mBase, uint32(v138+v158))) = v161
	v164 = v158 | int32(8)
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v52+v164)))
	*(*float64)(unsafe.Add(mBase, uint32(v138+v164))) = v167
	v170 = v158 | int32(16)
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v52+v170)))
	*(*float64)(unsafe.Add(mBase, uint32(v138+v170))) = v173
	v176 = v158 | int32(24)
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v52+v176)))
	*(*float64)(unsafe.Add(mBase, uint32(v138+v176))) = v179
	v181 = int32(4)
	v182 = v145 + v181
	v184 = v150 + v181
	if v184 != v33&int32(2147483644) {
		v145 = v182
		v150 = v184
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v187 = v182
	goto L38
L41:
	;
	goto L40
L42:
	;
	v202 = v187
	v206 = v2
	goto L45
L43:
	;
	goto L44
L44:
	;
	if v122 == int32(0) {
		goto L32
	} else {
		goto L48
	}
L45:
	;
	v215 = v202 << (uint(int32(3)) % 32)
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v52+v215)))
	*(*float64)(unsafe.Add(mBase, uint32(v138+v215))) = v218
	v220 = int32(1)
	v223 = v206 + v220
	if v223 != v200 {
		v202 = v202 + v220
		v206 = v223
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	goto L46
L48:
	;
	v241 = v33 & int32(3)
	v243 = v125 + int32(8)
	v244 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v252 = v244
	v255 = int32(0)
	goto L52
L50:
	;
	v312 = v244
	goto L51
L51:
	;
	if v241 == int32(0) {
		goto L31
	} else {
		goto L55
	}
L52:
	;
	v265 = int32(3)
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v62+v252<<(uint(v265)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v243+(v252+v33)<<(uint(v265)%32)))) = v271
	v274 = v252 | int32(1)
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v62+v274<<(uint(v265)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v243+(v274+v33)<<(uint(v265)%32)))) = v282
	v285 = v252 | int32(2)
	v293 = *(*float64)(unsafe.Add(mBase, uint32(v62+v285<<(uint(v265)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v243+(v285+v33)<<(uint(v265)%32)))) = v293
	v296 = v252 | v265
	v304 = *(*float64)(unsafe.Add(mBase, uint32(v62+v296<<(uint(v265)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v243+(v296+v33)<<(uint(v265)%32)))) = v304
	v306 = int32(4)
	v307 = v252 + v306
	v309 = v255 + v306
	if v309 != v33&int32(2147483644) {
		v252 = v307
		v255 = v309
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v312 = v307
	goto L51
L54:
	;
	goto L53
L55:
	;
	v327 = v312
	v329 = v244
	goto L56
L56:
	;
	v340 = int32(3)
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v62+v327<<(uint(v340)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v243+(v327+v33)<<(uint(v340)%32)))) = v346
	v348 = int32(1)
	v351 = v329 + v348
	if v351 != v241 {
		v327 = v327 + v348
		v329 = v351
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L31
L58:
	;
	goto L57
L59:
	;
	goto L32
L60:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(174799), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(499283), int32(158), int32(555523))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(419628), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(100)
	F_errdetail(m, int32(591730), v16)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(499283), int32(166), int32(555523))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
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
	F_errcode(m, int32(352845954))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(320763), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(499283), int32(171), int32(555523))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_c_f8_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v92 float64
	_ = v92
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v101 int32
	_ = v101
	var v104 float64
	_ = v104
	var v107 int32
	_ = v107
	var v110 float64
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v159 float64
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 float64
	_ = v246
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 float64
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v277 int32
	_ = v277
	var v278 float64
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 float64
	_ = v309
	var v321 int32
	_ = v321
	var v322 float64
	_ = v322
	var v332 int32
	_ = v332
	var v354 int32
	_ = v354
	var v364 float64
	_ = v364
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
		if base.Ui32(v28&int32(2147483644)) < base.Ui32(int32(100)) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(v35)))
			if int32(0) <= v28 {
				v189 = v28<<(uint(int32(4))%32) + int32(24)
				v190 = F_palloc0(m, v189)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v190))) = v189 << (uint(int32(2)) % 32)
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					v196 = int32(2147483647)
					v199 = v195&v196 + int32(1)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v199 | v200&int32(-2147483648)
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
					v207 = v205 & v196
					if v207 == int32(0) {
						v332 = v199 & int32(2147483647)
					} else {
						v212 = int32(1)
						v213 = v205 & v212
						v215 = v199 & int32(2147483647)
						v216 = int32(8)
						v217 = v190 + v216
						v219 = v24 + v216
						v220 = int32(0)
						if v207 != v212 {
							v225 = v220
							v233 = v2
							for {
								v242 = int32(3)
								v243 = v225 << (uint(v242) % 32)
								v245 = v243 + v219
								v246 = *(*float64)(unsafe.Add(mBase, uint32(v245)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+v243))) = v246
								v257 = base.B2i32(v205 < int32(0))
								if v205 < int32(0) {
									v258 = v245
								} else {
									v258 = v219 + (v225+v205)<<(uint(v242)%32)
								}
								v259 = *(*float64)(unsafe.Add(mBase, uint32(v258)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+(v225+v215)<<(uint(v242)%32)))) = v259
								v262 = v225 | int32(1)
								v263 = int32(3)
								v264 = v262 << (uint(v263) % 32)
								v266 = v219 + v264
								v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+v264))) = v267
								if v205 < int32(0) {
									v277 = v266
								} else {
									v277 = v219 + (v205+v262)<<(uint(v263)%32)
								}
								v278 = *(*float64)(unsafe.Add(mBase, uint32(v277)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+(v262+v215)<<(uint(v263)%32)))) = v278
								v280 = int32(2)
								v281 = v225 + v280
								v283 = v233 + v280
								if v283 != v207-v213 {
									v225 = v281
									v233 = v283
									continue
								} else {
									break
								}
								break
							}
							v286 = v281
						} else {
							v286 = v220
						}
						if v213 == int32(0) {
							v332 = v215
						} else {
							v305 = int32(3)
							v306 = v286 << (uint(v305) % 32)
							v308 = v306 + v219
							v309 = *(*float64)(unsafe.Add(mBase, uint32(v308)))
							*(*float64)(unsafe.Add(mBase, uint32(v217+v306))) = v309
							if v205 < int32(0) {
								v321 = v308
							} else {
								v321 = v219 + (v286+v205)<<(uint(v305)%32)
							}
							v322 = *(*float64)(unsafe.Add(mBase, uint32(v321)))
							*(*float64)(unsafe.Add(mBase, uint32(v217+(v286+v215)<<(uint(v305)%32)))) = v322
							v332 = v215
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v332<<(uint(int32(3))%32)+v190))) = v36
					v354 = v190
					v364 = v34
					v366 = v199 << (uint(int32(1)) % 32)
					*(*float64)(unsafe.Add(mBase, uint32(v366<<(uint(int32(3))%32)+v354))) = v364
					v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v371 != v24 {
						F_pfree(m, v24)
						mBase = m.M
						v374 = m.ExcPending
						if v374 != 0 {
							return int32(0)
						} else {
							m.G0 = v21 + int32(16)
							return v354
						}
					} else {
						m.G0 = v21 + int32(16)
						return v354
					}
				}
			} else {
				if base.F64_ne(v36, v34) != 0 {
					v189 = v28<<(uint(int32(4))%32) + int32(24)
					v190 = F_palloc0(m, v189)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v190))) = v189 << (uint(int32(2)) % 32)
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v196 = int32(2147483647)
						v199 = v195&v196 + int32(1)
						v200 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v199 | v200&int32(-2147483648)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v207 = v205 & v196
						if v207 == int32(0) {
							v332 = v199 & int32(2147483647)
						} else {
							v212 = int32(1)
							v213 = v205 & v212
							v215 = v199 & int32(2147483647)
							v216 = int32(8)
							v217 = v190 + v216
							v219 = v24 + v216
							v220 = int32(0)
							if v207 != v212 {
								v225 = v220
								v233 = v2
								for {
									v242 = int32(3)
									v243 = v225 << (uint(v242) % 32)
									v245 = v243 + v219
									v246 = *(*float64)(unsafe.Add(mBase, uint32(v245)))
									*(*float64)(unsafe.Add(mBase, uint32(v217+v243))) = v246
									v257 = base.B2i32(v205 < int32(0))
									if v205 < int32(0) {
										v258 = v245
									} else {
										v258 = v219 + (v225+v205)<<(uint(v242)%32)
									}
									v259 = *(*float64)(unsafe.Add(mBase, uint32(v258)))
									*(*float64)(unsafe.Add(mBase, uint32(v217+(v225+v215)<<(uint(v242)%32)))) = v259
									v262 = v225 | int32(1)
									v263 = int32(3)
									v264 = v262 << (uint(v263) % 32)
									v266 = v219 + v264
									v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
									*(*float64)(unsafe.Add(mBase, uint32(v217+v264))) = v267
									if v205 < int32(0) {
										v277 = v266
									} else {
										v277 = v219 + (v205+v262)<<(uint(v263)%32)
									}
									v278 = *(*float64)(unsafe.Add(mBase, uint32(v277)))
									*(*float64)(unsafe.Add(mBase, uint32(v217+(v262+v215)<<(uint(v263)%32)))) = v278
									v280 = int32(2)
									v281 = v225 + v280
									v283 = v233 + v280
									if v283 != v207-v213 {
										v225 = v281
										v233 = v283
										continue
									} else {
										break
									}
									break
								}
								v286 = v281
							} else {
								v286 = v220
							}
							if v213 == int32(0) {
								v332 = v215
							} else {
								v305 = int32(3)
								v306 = v286 << (uint(v305) % 32)
								v308 = v306 + v219
								v309 = *(*float64)(unsafe.Add(mBase, uint32(v308)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+v306))) = v309
								if v205 < int32(0) {
									v321 = v308
								} else {
									v321 = v219 + (v286+v205)<<(uint(v305)%32)
								}
								v322 = *(*float64)(unsafe.Add(mBase, uint32(v321)))
								*(*float64)(unsafe.Add(mBase, uint32(v217+(v286+v215)<<(uint(v305)%32)))) = v322
								v332 = v215
							}
						}
						*(*float64)(unsafe.Add(mBase, uint32(v332<<(uint(int32(3))%32)+v190))) = v36
						v354 = v190
						v364 = v34
						v366 = v199 << (uint(int32(1)) % 32)
						*(*float64)(unsafe.Add(mBase, uint32(v366<<(uint(int32(3))%32)+v354))) = v364
						v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v371 != v24 {
							F_pfree(m, v24)
							mBase = m.M
							v374 = m.ExcPending
							if v374 != 0 {
								return int32(0)
							} else {
								m.G0 = v21 + int32(16)
								return v354
							}
						} else {
							m.G0 = v21 + int32(16)
							return v354
						}
					}
				} else {
					v43 = v28<<(uint(int32(3))%32) + int32(16)
					v44 = F_palloc0(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v44))) = v43 << (uint(int32(2)) % 32)
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v51 = v49 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v51 | int32(-2147483648)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
						v57 = v55 & int32(2147483647)
						if v57 == int32(0) {
						} else {
							v61 = v55 & int32(3)
							v62 = int32(8)
							v63 = v44 + v62
							v65 = v24 + v62
							v66 = int32(0)
							if base.Ui32(int32(4)) <= base.Ui32(v57) {
								v71 = v66
								v78 = v2
								for {
									v89 = v71 << (uint(int32(3)) % 32)
									v92 = *(*float64)(unsafe.Add(mBase, uint32(v89+v65)))
									*(*float64)(unsafe.Add(mBase, uint32(v63+v89))) = v92
									v95 = v89 | int32(8)
									v98 = *(*float64)(unsafe.Add(mBase, uint32(v65+v95)))
									*(*float64)(unsafe.Add(mBase, uint32(v63+v95))) = v98
									v101 = v89 | int32(16)
									v104 = *(*float64)(unsafe.Add(mBase, uint32(v65+v101)))
									*(*float64)(unsafe.Add(mBase, uint32(v63+v101))) = v104
									v107 = v89 | int32(24)
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v107+v65)))
									*(*float64)(unsafe.Add(mBase, uint32(v63+v107))) = v110
									v112 = int32(4)
									v113 = v71 + v112
									v115 = v78 + v112
									if v115 != v57-v61 {
										v71 = v113
										v78 = v115
										continue
									} else {
										break
									}
									break
								}
								v118 = v113
							} else {
								v118 = v66
							}
							if v61 == int32(0) {
							} else {
								v138 = v118
								v140 = v2
								for {
									v156 = v138 << (uint(int32(3)) % 32)
									v159 = *(*float64)(unsafe.Add(mBase, uint32(v156+v65)))
									*(*float64)(unsafe.Add(mBase, uint32(v63+v156))) = v159
									v161 = int32(1)
									v164 = v140 + v161
									if v164 != v61 {
										v138 = v138 + v161
										v140 = v164
										continue
									} else {
										break
									}
									break
								}
							}
						}
						v354 = v44
						v364 = v36
						v366 = v51 & int32(2147483647)
						*(*float64)(unsafe.Add(mBase, uint32(v366<<(uint(int32(3))%32)+v354))) = v364
						v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v371 != v24 {
							F_pfree(m, v24)
							mBase = m.M
							v374 = m.ExcPending
							if v374 != 0 {
								return int32(0)
							} else {
								m.G0 = v21 + int32(16)
								return v354
							}
						} else {
							m.G0 = v21 + int32(16)
							return v354
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v382 = m.ExcPending
			if v382 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v385 = m.ExcPending
				if v385 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(419628), int32(0))
					mBase = m.M
					v391 = m.ExcPending
					if v391 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(100)
						F_errdetail(m, int32(591730), v21)
						mBase = m.M
						v397 = m.ExcPending
						if v397 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499283), int32(1883), int32(555501))
							mBase = m.M
							v404 = m.ExcPending
							if v404 != 0 {
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
func F_cube_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v57 float64
	_ = v57
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v111 int32
	_ = v111
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v120 float64
	_ = v120
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v132 float64
	_ = v132
	var v135 float64
	_ = v135
	var v136 int32
	_ = v136
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v150 float64
	_ = v150
	var v152 float64
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(0)
	if v11 == v13 {
		v163 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v188 != v6 {
		goto L42
	} else {
		goto L43
	}
L5:
	;
	v187 = v163
	goto L4
L6:
	;
	if v6 == int32(0) {
		v163 = v13
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v6 + int32(8)
	v42 = v31
	goto L11
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L19
	} else {
		goto L20
	}
L11:
	;
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v37+v42<<(uint(int32(3))%32))))
	if base.F64_ne(v57, float64(0)) != 0 {
		v163 = v13
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v37+(v42+v34)<<(uint(int32(3))%32))))
	if base.F64_ne(v66, float64(0)) != 0 {
		v163 = v13
		goto L5
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v70 = v42 + int32(1)
	if v70 != v34 {
		v42 = v70
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	goto L12
L19:
	;
	v87 = v31
	goto L21
L20:
	;
	v87 = v34
	goto L21
L21:
	;
	if v87 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v187 = int32(1)
	goto L4
L23:
	;
	goto L24
L24:
	;
	v91 = int32(8)
	v92 = v6 + v91
	v94 = v11 + v91
	v98 = int32(0)
	goto L25
L25:
	;
	v111 = v98 << (uint(int32(3)) % 32)
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v94+v111)))
	v115 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v123 = v113
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v163 = v155
	goto L5
L27:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v92+v111)))
	v127 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v135 = v125
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v94+(v98+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v113, v120) != 0 {
		v123 = v113
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v123 = v120
	goto L27
L30:
	;
	v136 = int32(0)
	if base.F64_gt(v123, v135) != 0 {
		v163 = v136
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v92+(v98+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v125, v132) != 0 {
		v135 = v125
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v135 = v132
	goto L30
L33:
	;
	if v29 < int32(0) {
		v144 = v113
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v32 < int32(0) {
		v152 = v125
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v94+(v98+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v113, v142) != 0 {
		v144 = v113
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v144 = v142
	goto L34
L37:
	;
	if base.F64_gt(v152, v144) != 0 {
		v163 = v136
		goto L5
	} else {
		goto L40
	}
L38:
	;
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v92+(v98+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v125, v150) != 0 {
		v152 = v125
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v152 = v150
	goto L37
L40:
	;
	v155 = int32(1)
	v157 = v98 + v155
	if v157 != v87 {
		v98 = v157
		goto L25
	} else {
		goto L41
	}
L41:
	;
	goto L26
L42:
	;
	F_pfree(m, v6)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v192 != v11 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L44
L46:
	;
	F_pfree(m, v11)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	return v187
L49:
	;
	goto L48
}
func F_cube_coord_llur(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 float64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 float64
	_ = v38
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v51 float64
	_ = v51
	var v57 float64
	_ = v57
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 != 0 {
			v16 = v13 >> (uint(int32(31)) % 32)
			v18 = v13 ^ v16 - v16
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if base.Ui32(v19<<(uint(int32(1))%32)) < base.Ui32(v18) {
				v57 = float64(0)
			} else {
				v23 = int32(1)
				v24 = v18 - v23
				v26 = int32(base.Ui32(v24) >> (uint(v23) % 32))
				if v19 < int32(0) {
					v32 = *(*float64)(unsafe.Add(mBase, uint32(v9+v26<<(uint(int32(3))%32))+8))
					v57 = v32
				} else {
					v34 = v9 + int32(8)
					v35 = int32(3)
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v34+v26<<(uint(v35)%32))))
					v45 = *(*float64)(unsafe.Add(mBase, uint32(v34+(v19&int32(2147483647)+v26)<<(uint(v35)%32))))
					if base.F64_gt(v38, v45) != 0 {
						v47 = v38
					} else {
						v47 = v45
					}
					if v24&int32(1) != 0 {
						v57 = v47
					} else {
						if base.F64_lt(v38, v45) != 0 {
							v51 = v38
						} else {
							v51 = v45
						}
						v57 = v51
					}
				}
			}
			if v13 < int32(0) {
				v61 = base.F64_neg(v57)
			} else {
				v61 = v57
			}
			v62 = F_Float8GetDatum(m, v61)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				return v62
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(452581), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499283), int32(1659), int32(206111))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
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
func F_cube_f8_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_eq(v5, v7) != 0 {
		v10 = F_palloc0(m, int32(16))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v5
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(-9223372032559808448)
			return v10
		}
	} else {
		v19 = F_palloc0(m, int32(24))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(96)
			*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v7
			*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v5
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v25&int32(-2147483648) | int32(1)
			return v19
		}
	}
}
func F_cube_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_cube_scanner_init(m, v7, v5+int32(8), v5+int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v21 = F_cube_yyparse(m, v5+int32(12), v18, v19, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			F_cube_scanner_finish(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				m.G0 = v5 + int32(16)
				return v26
			}
		}
	}
}
func F_cube_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 float64
	_ = v286
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 float64
	_ = v350
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 float64
	_ = v406
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v430 int32
	_ = v430
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 != v6 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	v487 = v465
	goto L4
L6:
	;
	v465 = int32(1)
	goto L5
L7:
	;
	v36 = v31
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(8)
	v38 = v11 + v37
	v40 = v6 + v37
	v48 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L47
	} else {
		goto L48
	}
L13:
	;
	v60 = v48 << (uint(int32(3)) % 32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v40+v60)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v72 = v62
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = int32(8)
	v108 = v11 + v107
	v110 = v6 + v107
	v119 = int32(0)
	goto L30
L15:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38+v60)))
	v76 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v83 = v74
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v69) != 0 {
		v72 = v62
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if base.F64_gt(v72, v83) != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v81) != 0 {
		v83 = v74
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v83 = v81
	goto L18
L21:
	;
	if v29 < int32(0) {
		v92 = v62
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 < int32(0) {
		v100 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v90) != 0 {
		v92 = v62
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v90
	goto L22
L25:
	;
	v102 = int32(-1)
	if base.F64_lt(v92, v100) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v98) != 0 {
		v100 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	goto L25
L28:
	;
	v105 = v48 + int32(1)
	if v105 != v36 {
		v48 = v105
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v131 = v119 << (uint(int32(3)) % 32)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v110+v131)))
	v135 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v143 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v108+v131)))
	v147 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v154 = v145
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v140) != 0 {
		v143 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v143 = v140
	goto L32
L35:
	;
	if base.F64_gt(v143, v154) != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v152) != 0 {
		v154 = v145
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = v152
	goto L35
L38:
	;
	if v29 < int32(0) {
		v163 = v133
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v32 < int32(0) {
		v171 = v145
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v161) != 0 {
		v163 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L39
L42:
	;
	if base.F64_lt(v163, v171) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v169) != 0 {
		v171 = v145
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v171 = v169
	goto L42
L45:
	;
	v175 = v119 + int32(1)
	if v175 != v36 {
		v119 = v175
		goto L30
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	v197 = v6 + int32(8)
	v207 = v36
	goto L50
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v220 = v197 + v207<<(uint(int32(3))%32)
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v266 = v36
	goto L68
L52:
	;
	if base.F64_lt(v243, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v225 = int32(3)
	v227 = v197 + (v207+v29)<<(uint(v225)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v207+v31)<<(uint(v225)%32))))
	if base.F64_lt(v221, v232) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_gt(v221, float64(0)) != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v234 = v220
	goto L58
L57:
	;
	v234 = v227
	goto L58
L58:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v234)))
	if base.F64_gt(v235, float64(0)) != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.F64_lt(v221, v238) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v240 = v221
	goto L62
L61:
	;
	v240 = v238
	goto L62
L62:
	;
	v243 = v240
	goto L52
L63:
	;
	v243 = v221
	goto L52
L64:
	;
	v487 = int32(-1)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v250 = v207 + int32(1)
	if v250 != v31 {
		v207 = v250
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	v274 = v197 + v266<<(uint(int32(3))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v487 = int32(-1)
	goto L4
L70:
	;
	if base.F64_lt(v297, float64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v279 = int32(3)
	v281 = v197 + (v29+v266)<<(uint(v279)%32)
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v266+v31)<<(uint(v279)%32))))
	if base.F64_gt(v275, v286) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if base.F64_gt(v275, float64(0)) != 0 {
		goto L6
	} else {
		goto L81
	}
L74:
	;
	v288 = v274
	goto L76
L75:
	;
	v288 = v281
	goto L76
L76:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.F64_gt(v289, float64(0)) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_gt(v275, v292) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v275
	goto L80
L79:
	;
	v294 = v292
	goto L80
L80:
	;
	v297 = v294
	goto L70
L81:
	;
	v297 = v275
	goto L70
L82:
	;
	v304 = int32(1)
	v306 = v266 + v304
	if v306 == v31 {
		v465 = v304
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L69
L85:
	;
	v266 = v306
	goto L68
L86:
	;
	v487 = int32(0)
	goto L4
L87:
	;
	goto L88
L88:
	;
	v312 = v11 + int32(8)
	v328 = v31
	goto L89
L89:
	;
	v335 = v312 + v328<<(uint(int32(3))%32)
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v386 = v36
	goto L106
L91:
	;
	if base.F64_lt(v364, float64(0)) != 0 {
		goto L6
	} else {
		goto L104
	}
L92:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_lt(v336, v361) != 0 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v340 = int32(3)
	v342 = v312 + (v32+v328)<<(uint(v340)%32)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v328+v34)<<(uint(v340)%32))))
	if base.F64_lt(v336, v347) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.F64_gt(v336, float64(0)) == int32(0) {
		v364 = v336
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v349 = v335
	goto L98
L97:
	;
	v349 = v342
	goto L98
L98:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v349)))
	if base.F64_gt(v350, float64(0)) == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v487 = int32(-1)
	goto L4
L100:
	;
	v487 = int32(-1)
	goto L4
L101:
	;
	v363 = v336
	goto L103
L102:
	;
	v363 = v361
	goto L103
L103:
	;
	v364 = v363
	goto L91
L104:
	;
	v370 = v328 + int32(1)
	if v370 != v34 {
		v328 = v370
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v394 = v312 + v386<<(uint(int32(3))%32)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v465 = int32(-1)
	goto L5
L108:
	;
	if base.F64_lt(v423, float64(0)) != 0 {
		goto L6
	} else {
		goto L121
	}
L109:
	;
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v401)))
	if base.F64_gt(v395, v420) != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v399 = int32(3)
	v401 = v312 + (v32+v386)<<(uint(v399)%32)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v386+v34)<<(uint(v399)%32))))
	if base.F64_gt(v395, v406) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.F64_gt(v395, float64(0)) == int32(0) {
		v423 = v395
		goto L108
	} else {
		goto L117
	}
L113:
	;
	v408 = v394
	goto L115
L114:
	;
	v408 = v401
	goto L115
L115:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)))
	if base.F64_gt(v409, float64(0)) == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v487 = int32(-1)
	goto L4
L117:
	;
	v487 = int32(-1)
	goto L4
L118:
	;
	v422 = v395
	goto L120
L119:
	;
	v422 = v420
	goto L120
L120:
	;
	v423 = v422
	goto L108
L121:
	;
	v430 = v386 + int32(1)
	if v34 != v430 {
		v386 = v430
		goto L106
	} else {
		goto L122
	}
L122:
	;
	goto L107
L123:
	;
	F_pfree(m, v6)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v492 != v11 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_pfree(m, v11)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	return base.B2i32(v487 <= int32(0))
L130:
	;
	goto L129
}
func F_cube_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v6 = F_strlen(m, l0)
	mBase = m.M
	if l2 != 0 {
		v9 = F_palloc(m, int32(96))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v9
			if v9 != 0 {
				v33 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(96))
				mBase = m.M
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v35
				v37 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v34)+52)) = v37
				*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v34)+36)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v34)+4)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v34)+12)) = v37
				*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v35
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v50 = F_cube_yy_scan_bytes(m, l0, v6, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
					return
				}
			} else {
				v15 = int32(48)
				*(*int32)(unsafe.Add(mBase, _consts[137])) = v15
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(295468), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_errfinish(m, int32(314367), int32(108), int32(99911))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v15 = int32(28)
		*(*int32)(unsafe.Add(mBase, _consts[137])) = v15
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(295468), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errfinish(m, int32(314367), int32(108), int32(99911))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_cube_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 float64
	_ = v40
	var v42 int32
	_ = v42
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v54 int32
	_ = v54
	var v59 float64
	_ = v59
	var v63 float64
	_ = v63
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v81 float64
	_ = v81
	var v85 int32
	_ = v85
	var v88 float64
	_ = v88
	var v92 float64
	_ = v92
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	v2 = int32(0)
	v10 = float64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v105 = v10
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			if v18 <= int32(0) {
				v105 = v10
			} else {
				v21 = int32(1)
				v24 = v12 + int32(8)
				if v18 == v21 {
					v73 = v2
					v81 = float64(1)
				} else {
					v32 = v2
					v36 = v2
					v40 = float64(1)
					for {
						v42 = int32(3)
						v45 = *(*float64)(unsafe.Add(mBase, uint32(v24+(v32+v18)<<(uint(v42)%32))))
						v49 = *(*float64)(unsafe.Add(mBase, uint32(v24+v32<<(uint(v42)%32))))
						v54 = v32 | int32(1)
						v59 = *(*float64)(unsafe.Add(mBase, uint32(v24+(v54+v18)<<(uint(v42)%32))))
						v63 = *(*float64)(unsafe.Add(mBase, uint32(v24+v54<<(uint(v42)%32))))
						v66 = base.F64_mul(base.F64_mul(v40, base.F64_abs(base.F64_sub(v45, v49))), base.F64_abs(base.F64_sub(v59, v63)))
						v67 = int32(2)
						v68 = v32 + v67
						v70 = v36 + v67
						if v70 != v18&int32(2147483646) {
							v32 = v68
							v36 = v70
							v40 = v66
							continue
						} else {
							break
						}
						break
					}
					v73 = v68
					v81 = v66
				}
				if v18&v21 == int32(0) {
					v105 = v81
				} else {
					v85 = int32(3)
					v88 = *(*float64)(unsafe.Add(mBase, uint32(v24+(v73+v18)<<(uint(v85)%32))))
					v92 = *(*float64)(unsafe.Add(mBase, uint32(v24+v73<<(uint(v85)%32))))
					v105 = base.F64_mul(v81, base.F64_abs(base.F64_sub(v88, v92)))
				}
			}
		}
		v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v106 != v12 {
			F_pfree(m, v12)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return int32(0)
			} else {
				v110 = F_Float8GetDatum(m, v105)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					return v110
				}
			}
		} else {
			v110 = F_Float8GetDatum(m, v105)
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return int32(0)
			} else {
				return v110
			}
		}
	}
}
func F_cube_yyalloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
