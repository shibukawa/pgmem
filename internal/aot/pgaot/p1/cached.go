package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int64
	_ = v69
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 float64
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 float64
	_ = v274
	var v276 float64
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v284 float64
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v304 float64
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 float64
	_ = v315
	var v329 float64
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int64
	_ = v381
	var v384 float64
	_ = v384
	var v385 float64
	_ = v385
	var v394 int32
	_ = v394
	var v404 float64
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 float64
	_ = v415
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v431 float64
	_ = v431
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 float64
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 float64
	_ = v450
	var v454 float64
	_ = v454
	var v456 int32
	_ = v456
	var v470 float64
	_ = v470
	var v473 float64
	_ = v473
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v510 int64
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	v13 = float64(0)
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v509 = l0 + v500
	v510 = *(*int64)(unsafe.Add(mBase, uint32(v509)))
	*(*int64)(unsafe.Add(mBase, uint32(v509))) = v510 + int64(1)
	if l2 != 0 {
		goto L155
	} else {
		goto L156
	}
L2:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v498 = v491
	v499 = int32(0)
	v500 = int32(136)
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L8
	} else {
		goto L151
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v16 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v19 = F_RevalidateCachedQuery(m, l0, l3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L6
L8:
	;
	return int32(0)
L9:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v23 != 0 {
		v394 = v19
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v404 = float64(0)
	v405 = F_BuildCachedPlan(m, l0, v394, l1, l3)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L139
	}
L11:
	;
	if l1 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v81 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[2]))
	switch v61 - int32(1) {
	case 0:
		goto L12
	case 1:
		v394 = v19
		goto L10
	default:
		goto L28
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	switch v30 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v34 = int32(1)
		goto L19
	default:
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v35 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L18:
	;
	if v34 != 0 {
		goto L14
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v34 = int32(0)
	goto L19
L21:
	;
	goto L12
L22:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 != int32(6) {
		v54 = int32(1)
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v54&int32(1) == int32(0) {
		goto L12
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = v44 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v46) {
		v54 = int32(0)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v54 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v46)) % 64)))
	goto L24
L27:
	;
	goto L14
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v64&int32(512) != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	if v64&int32(1024) != 0 {
		v394 = v19
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v69 < int64(5) {
		v394 = v19
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v72 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	v73 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.F64_lt(v72, base.F64_div(v73, base.F64_convert_i64_u(v69))) == int32(0) {
		v394 = v19
		goto L10
	} else {
		goto L32
	}
L32:
	;
	goto L12
L33:
	;
	v138 = F_BuildCachedPlan(m, l0, v19, int32(0), l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L55
	}
L34:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v84 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v120 == int32(0) {
		goto L33
	} else {
		goto L51
	}
L36:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
	if v87 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[3]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v91 != v92 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_AcquireExecutorLocks(m, v99, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L8
	} else {
		goto L44
	}
L40:
	;
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)) = uint8(v94)
	goto L35
L41:
	;
	goto L42
L42:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v96 != int32(1) {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v103 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v106 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_AcquireExecutorLocks(m, v115, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[4]))
	if v106 == v110 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)) = uint8(v112)
	goto L47
L50:
	;
	goto L35
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	v127 = v125 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = v127
	if v127 != 0 {
		goto L33
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(0)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	if v131 != 0 {
		goto L33
	} else {
		goto L53
	}
L53:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v120)+32))
	F_MemoryContextDelete(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	goto L33
L55:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v140 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v138
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v158 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+28)) = v157 + v158
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v138)+32))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v162 == v158 {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)+28))
	v147 = v145 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v140)+28)) = v147
	if v147 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = int32(0)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+8)))
	if v151 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v140)+32))
	F_MemoryContextDelete(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	goto L56
L61:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v235 == int32(0) {
		v329 = v13
		goto L99
	} else {
		goto L100
	}
L62:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[1]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	if v170 != v166 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	if v206 != v202 {
		goto L83
	} else {
		goto L84
	}
L65:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+9)) = uint8(v199)
	goto L61
L66:
	;
	if v170 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	if v166 != 0 {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v161)+28))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v175 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v174 == int32(0) {
		goto L69
	} else {
		goto L75
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+28)) = v174
	goto L71
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v174
	goto L71
L75:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = v180
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v166
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v166)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+28)) = v187
	if v187 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v161)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = int32(0)
	goto L68
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+24)) = v161
	goto L81
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+20)) = v161
	goto L65
L82:
	;
	goto L61
L83:
	;
	if v206 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	goto L82
L86:
	;
	if v202 != 0 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v161)+28))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	if v211 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v210 == int32(0) {
		goto L86
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+28)) = v210
	goto L88
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+20)) = v210
	goto L88
L92:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v161)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+24)) = v216
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v202
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v202)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+28)) = v223
	if v223 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v161)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = int32(0)
	goto L85
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+24)) = v161
	goto L98
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+20)) = v161
	goto L82
L99:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+112)) = v329
	v333 = int32(0)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v334 != 0 {
		v394 = v333
		goto L10
	} else {
		goto L117
	}
L100:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v238 <= int32(0) {
		v329 = v13
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v235)+12))
	if v238 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v241+v297<<(uint(int32(2))%32))))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	if v311 == int32(6) {
		v329 = v304
		goto L99
	} else {
		goto L116
	}
L103:
	;
	v297 = int32(0)
	v304 = v13
	goto L102
L104:
	;
	goto L105
L105:
	;
	v249 = int32(0)
	v256 = v249
	v257 = v249
	v263 = v13
	goto L106
L106:
	;
	v268 = v241 + v256<<(uint(int32(2))%32)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v270 != int32(6) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v238&int32(1) == int32(0) {
		v329 = v284
		goto L99
	} else {
		goto L115
	}
L108:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269)+36))
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v273)+16))
	v276 = base.F64_add(v263, v274)
	goto L110
L109:
	;
	v276 = v263
	goto L110
L110:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	if v278 != int32(6) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)+36))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v281)+16))
	v284 = base.F64_add(v276, v282)
	goto L113
L112:
	;
	v284 = v276
	goto L113
L113:
	;
	v285 = int32(2)
	v286 = v256 + v285
	v288 = v257 + v285
	if v288 != v238&int32(2147483646) {
		v256 = v286
		v257 = v288
		v263 = v284
		goto L106
	} else {
		goto L114
	}
L114:
	;
	goto L107
L115:
	;
	v297 = v286
	v304 = v284
	goto L102
L116:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v310)+36))
	v315 = *(*float64)(unsafe.Add(mBase, uint32(v314)+16))
	v329 = base.F64_add(v304, v315)
	goto L99
L117:
	;
	v335 = int32(136)
	if l1 == int32(0) {
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v338 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[2]))
	switch v373 - int32(1) {
	case 0:
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	case 1:
		v394 = v333
		goto L10
	default:
		goto L133
	}
L120:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	switch v342 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v346 = int32(1)
		goto L124
	default:
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v347 == int32(0) {
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	} else {
		goto L127
	}
L123:
	;
	if v346 != 0 {
		goto L119
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v346 = int32(0)
	goto L124
L126:
	;
	v498 = v138
	v499 = v333
	v500 = v335
	goto L1
L127:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v347)+4))
	if v351 != int32(6) {
		v366 = int32(1)
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if v366&int32(1) == int32(0) {
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	} else {
		goto L132
	}
L129:
	;
	goto L128
L130:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v347)+28))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v358 = v356 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v358) {
		v366 = int32(0)
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v366 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v358)) % 64)))
	goto L129
L132:
	;
	goto L119
L133:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v376&int32(512) != 0 {
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v376&int32(1024) != 0 {
		v394 = v333
		goto L10
	} else {
		goto L135
	}
L135:
	;
	v381 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v381 < int64(5) {
		v394 = v333
		goto L10
	} else {
		goto L136
	}
L136:
	;
	v384 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	v385 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.F64_lt(v384, base.F64_div(v385, base.F64_convert_i64_u(v381))) != 0 {
		v498 = v138
		v499 = v333
		v500 = v335
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v394 = v333
	goto L10
L138:
	;
	v473 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+120)) = base.F64_add(v470, v473)
	v498 = v405
	v499 = int32(1)
	v500 = int32(128)
	goto L1
L139:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v407 == int32(0) {
		v470 = v404
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v410 <= int32(0) {
		v470 = v404
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v415 = *(*float64)(unsafe.Add(mBase, _c_F_GetCachedPlan[0]))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	v424 = int32(0)
	v431 = v404
	goto L142
L142:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v418+v424<<(uint(int32(2))%32))))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v438 != int32(6) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v470 = v454
	goto L138
L144:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)+36))
	v442 = *(*float64)(unsafe.Add(mBase, uint32(v441)+16))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v437)+44))
	if v444 != 0 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v454 = v431
	goto L146
L146:
	;
	v456 = v424 + int32(1)
	if v410 != v456 {
		v424 = v456
		v431 = v454
		goto L142
	} else {
		goto L150
	}
L147:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v450 = base.F64_convert_i32_s(v445 + int32(1))
	goto L149
L148:
	;
	v450 = float64(1)
	goto L149
L149:
	;
	v454 = base.F64_add(base.F64_mul(base.F64_mul(v415, float64(1000)), v450), base.F64_add(v431, v442))
	goto L146
L150:
	;
	goto L143
L151:
	;
	F_errmsg_internal(m, int32(_a_F_GetCachedPlan_0), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L8
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_GetCachedPlan_1), int32(1292), int32(_a_F_GetCachedPlan_2))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	if v499 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L155:
	;
	F_ResourceOwnerEnlarge(m, l2)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L8
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v498)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+28)) = v523 + int32(1)
	goto L154
L158:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v498)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v498)+28)) = v516 + int32(1)
	F_ResourceOwnerRemember(m, l2, v498, int32(_a_F_GetCachedPlan_3))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L8
	} else {
		goto L159
	}
L159:
	;
	goto L154
L160:
	;
	return v498
L161:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v529 != int32(1) {
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v498)+32))
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_GetCachedPlan[1]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	if v538 != v534 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v567 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v498)+9)) = uint8(v567)
	goto L160
L164:
	;
	if v538 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	goto L163
L167:
	;
	if v534 != 0 {
		goto L174
	} else {
		goto L175
	}
L168:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v532)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v532)+24))
	if v543 != 0 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v542 == int32(0) {
		goto L167
	} else {
		goto L173
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v543)+28)) = v542
	goto L169
L171:
	;
	goto L172
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+20)) = v542
	goto L169
L173:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v532)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v542)+24)) = v548
	goto L167
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v532)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v532)+16)) = v534
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v532)+28)) = v555
	if v555 != 0 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v532)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v532)+16)) = int32(0)
	goto L166
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+24)) = v532
	goto L179
L178:
	;
	goto L179
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v534)+20)) = v532
	goto L163
}
func F_RevalidateCachedQuery(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	v3 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v7 != 0 {
		v489 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L20
	} else {
		goto L182
	}
L2:
	;
	return v489
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v42 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	switch v12 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v16 = int32(1)
		goto L9
	default:
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 == int32(0) {
		v489 = v3
		goto L2
	} else {
		goto L12
	}
L8:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v16 = int32(0)
	goto L9
L11:
	;
	v489 = v3
	goto L2
L12:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 != int32(6) {
		v36 = int32(1)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v36&int32(1) == int32(0) {
		v489 = v3
		goto L2
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = v26 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v28) {
		v36 = int32(0)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v36 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v28)) % 64)))
	goto L14
L17:
	;
	goto L4
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	v60 = base.B2i32(v58 != int32(1))
	if v58 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v46 = F_SearchPathMatchesCurrentEnvironment(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v46 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v50)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v52 == v50 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+10)) = uint8(v55)
	goto L18
L24:
	;
	v198 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v198
	v200 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v200)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v198
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v204 != 0 {
		goto L82
	} else {
		goto L83
	}
L25:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v61 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v65 != v66 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v73 == int32(0) {
		v489 = v3
		goto L2
	} else {
		goto L31
	}
L29:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[1])))
	if v60|base.B2i32(v68 != v70) != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v76 <= int32(0) {
		v489 = v3
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v82 = int32(0)
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v82<<(uint(int32(2))%32))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v91 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v132 != 0 {
		v489 = v3
		goto L2
	} else {
		goto L56
	}
L35:
	;
	v129 = v82 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v129 < v130 {
		v82 = v129
		goto L33
	} else {
		goto L55
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+28))
	v96 = v94
	goto L40
L37:
	;
	v123 = v90
	goto L38
L38:
	;
	F_ScanQueryForLocks(m, v123, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L54
	}
L39:
	;
	if v120 == int32(0) {
		goto L35
	} else {
		goto L53
	}
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	switch v98 - int32(241) {
	case 0:
		goto L45
	case 1:
		goto L44
	default:
		goto L46
	}
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+28))
	v96 = v118
	goto L40
L43:
	;
	v120 = v116
	goto L39
L44:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v113 == int32(6) {
		v117 = v112
		goto L42
	} else {
		goto L52
	}
L45:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v109 == int32(6) {
		v117 = v108
		goto L42
	} else {
		goto L51
	}
L46:
	;
	if v98 != int32(201) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v120 = int32(0)
	goto L39
L48:
	;
	goto L49
L49:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v105 != int32(6) {
		v116 = v104
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v117 = v104
	goto L42
L51:
	;
	v116 = v108
	goto L43
L52:
	;
	v116 = v112
	goto L43
L53:
	;
	v123 = v120
	goto L38
L54:
	;
	goto L35
L55:
	;
	goto L34
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v133 == int32(0) {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v136 <= int32(0) {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	v142 = int32(0)
	goto L59
L59:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v142<<(uint(int32(2))%32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	if v151 == int32(6) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L24
L61:
	;
	v189 = v142 + int32(1)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v189 < v190 {
		v142 = v189
		goto L59
	} else {
		goto L81
	}
L62:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+28))
	v156 = v154
	goto L66
L63:
	;
	v183 = v150
	goto L64
L64:
	;
	F_ScanQueryForLocks(m, v183, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L20
	} else {
		goto L80
	}
L65:
	;
	if v180 == int32(0) {
		goto L61
	} else {
		goto L79
	}
L66:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	switch v158 - int32(241) {
	case 0:
		goto L71
	case 1:
		goto L70
	default:
		goto L72
	}
L68:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+28))
	v156 = v178
	goto L66
L69:
	;
	v180 = v176
	goto L65
L70:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v173 == int32(6) {
		v177 = v172
		goto L68
	} else {
		goto L78
	}
L71:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v169 == int32(6) {
		v177 = v168
		goto L68
	} else {
		goto L77
	}
L72:
	;
	if v158 != int32(201) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v180 = int32(0)
	goto L65
L74:
	;
	goto L75
L75:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v165 != int32(6) {
		v176 = v164
		goto L69
	} else {
		goto L76
	}
L76:
	;
	v177 = v164
	goto L68
L77:
	;
	v176 = v168
	goto L69
L78:
	;
	v176 = v172
	goto L69
L79:
	;
	v183 = v180
	goto L64
L80:
	;
	goto L61
L81:
	;
	goto L60
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	F_MemoryContextDelete(m, v204)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v209 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L84
L86:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[2]))
	v228 = base.B2i32(v226 != int32(0))
	goto L91
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+28))
	v216 = v214 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+28)) = v216
	if v216 != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = int32(0)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
	if v220 != 0 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
	F_MemoryContextDelete(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	if v228 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v231 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L20
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v235 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	F_PushActiveSnapshot(m, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v264 != 0 {
		goto L111
	} else {
		goto L112
	}
L98:
	;
	v236 = F_copyObjectImpl(m, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L20
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v247 = int32(0)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v248 == v247 {
		v263 = v247
		goto L97
	} else {
		goto L107
	}
L101:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v239 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v241 = F_pg_analyze_and_rewrite_withcb(m, v236, v238, v239, v240, l1)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L20
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v245 = F_pg_analyze_and_rewrite_fixedparams(m, v236, v238, v243, v244, l1)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L20
	} else {
		goto L106
	}
L105:
	;
	v263 = v241
	goto L97
L106:
	;
	v263 = v245
	goto L97
L107:
	;
	v251 = F_copyObjectImpl(m, v248)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L20
	} else {
		goto L108
	}
L108:
	;
	F_AcquireRewriteLocks(m, v251, int32(1), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L20
	} else {
		goto L109
	}
L109:
	;
	v257 = F_pg_rewrite_query(m, v251)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	v263 = v257
	goto L97
L111:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v264].(func(*base.Module, int32, int32))(m, v263, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L20
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v228 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L113
L115:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L20
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v272 = F_ChoosePortalStrategy(m, v263)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L20
	} else {
		goto L127
	}
L118:
	;
	goto L117
L119:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3]))
	v421 = F_AllocSetContextCreateInternal(m, v416, int32(_a_F_RevalidateCachedQuery_0), int32(0), int32(1024), int32(_a_F_RevalidateCachedQuery_1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L20
	} else {
		goto L161
	}
L120:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v403 != 0 {
		goto L157
	} else {
		goto L158
	}
L121:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v327 != 0 {
		goto L137
	} else {
		goto L138
	}
L122:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v317 == int32(0) {
		goto L119
	} else {
		goto L135
	}
L123:
	;
	if v310 != 0 {
		goto L121
	} else {
		goto L134
	}
L124:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+28))
	v302 = F_UtilityTupleDescriptor(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L20
	} else {
		goto L133
	}
L125:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v283 = int32(0)
	goto L129
L126:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+76))
	v277 = F_ExecCleanTypeFromTL(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L20
	} else {
		goto L128
	}
L127:
	;
	switch v272 {
	case 0, 2:
		goto L126
	case 1:
		goto L125
	case 3:
		goto L124
	default:
		goto L122
	}
L128:
	;
	v310 = v277
	goto L123
L129:
	;
	v289 = int32(1)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v279+v283<<(uint(int32(2))%32))))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+24)))
	if v293 != v289 {
		v283 = v283 + v289
		goto L129
	} else {
		goto L131
	}
L130:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292)+96))
	v297 = F_ExecCleanTypeFromTL(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L20
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	v310 = v297
	goto L123
L133:
	;
	v310 = v302
	goto L123
L134:
	;
	goto L122
L135:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v320 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v321 = int32(_a_F_RevalidateCachedQuery_2)
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3])) = v324
	v399 = v322
	v402 = int32(0)
	goto L120
L137:
	;
	v328 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v332 != v333 {
		v384 = v328
		goto L141
	} else {
		goto L142
	}
L138:
	;
	goto L139
L139:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v388 != 0 {
		goto L1
	} else {
		goto L155
	}
L140:
	;
	if v384 != 0 {
		goto L119
	} else {
		goto L154
	}
L141:
	;
	goto L140
L142:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v335 != v336 {
		v384 = v328
		goto L141
	} else {
		goto L143
	}
L143:
	;
	if v332 <= int32(0) {
		v384 = int32(1)
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v342 = v332 << (uint(int32(4)) % 32)
	v344 = int32(20)
	v350 = int32(0)
	goto L145
L145:
	;
	v357 = v350 * int32(100)
	v358 = v310 + v342 + v344 + v357
	v359 = int32(4)
	v361 = v357 + (v327 + v342 + v344)
	v364 = F_strcmp(m, v358+v359, v361+v359)
	mBase = m.M
	if v364 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v384 = int32(0)
	goto L141
L147:
	;
	goto L146
L148:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358)+68))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)+68))
	if v365 != v366 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v358)+76))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v361)+76))
	if v368 != v369 {
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v358)+96))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v361)+96))
	if v371 != v372 {
		goto L147
	} else {
		goto L151
	}
L151:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+91)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+91)))
	if v374 != v375 {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v377 = int32(1)
	v379 = v350 + v377
	if v332 != v379 {
		v350 = v379
		goto L145
	} else {
		goto L153
	}
L153:
	;
	v384 = v377
	goto L141
L154:
	;
	goto L139
L155:
	;
	v389 = int32(_a_F_RevalidateCachedQuery_2)
	v390 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3])) = v392
	v394 = F_CreateTupleDescCopy(m, v310)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	v399 = v390
	v402 = v394
	goto L120
L157:
	;
	F_FreeTupleDesc(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v402
	*(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3])) = v399
	goto L119
L160:
	;
	goto L159
L161:
	;
	v423 = int32(_a_F_RevalidateCachedQuery_2)
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3])) = v421
	v427 = F_copyObjectImpl(m, v263)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L20
	} else {
		goto L162
	}
L162:
	;
	F_extract_query_dependencies(m, v427, l0-int32(-64), l0+int32(68), l0+int32(85))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L20
	} else {
		goto L163
	}
L163:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v438
	v441 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v441)
	v443 = F_GetSearchPathMatcher(m, v421)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v443
	*(*int32)(unsafe.Add(mBase, _c_F_RevalidateCachedQuery[3])) = v424
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v421)+16))
	if v452 != v448 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v481 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v481)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v421
	v489 = v263
	goto L2
L166:
	;
	if v452 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L168
L168:
	;
	goto L165
L169:
	;
	if v448 != 0 {
		goto L176
	} else {
		goto L177
	}
L170:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v421)+28))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v421)+24))
	if v457 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if v456 == int32(0) {
		goto L169
	} else {
		goto L175
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v457)+28)) = v456
	goto L171
L173:
	;
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452)+20)) = v456
	goto L171
L175:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v421)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v456)+24)) = v462
	goto L169
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+16)) = v448
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v421)+28)) = v469
	if v469 != 0 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	goto L178
L178:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v421)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+16)) = int32(0)
	goto L168
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v469)+24)) = v421
	goto L181
L180:
	;
	goto L181
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448)+20)) = v421
	goto L165
L182:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L20
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_RevalidateCachedQuery_3), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L20
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_RevalidateCachedQuery_4), int32(860), int32(_a_F_RevalidateCachedQuery_5))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
