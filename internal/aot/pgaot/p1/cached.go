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
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 float64
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 float64
	_ = v294
	var v296 float64
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 float64
	_ = v302
	var v304 float64
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 float64
	_ = v322
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v349 float64
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v401 int64
	_ = v401
	var v404 float64
	_ = v404
	var v405 float64
	_ = v405
	var v414 int32
	_ = v414
	var v424 float64
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v435 float64
	_ = v435
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v451 float64
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 float64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 float64
	_ = v470
	var v474 float64
	_ = v474
	var v476 int32
	_ = v476
	var v490 float64
	_ = v490
	var v493 float64
	_ = v493
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v514 int64
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v572 int32
	_ = v572
	v13 = float64(0)
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	v513 = l0 + v502
	v514 = *(*int64)(unsafe.Add(mBase, uint32(v513)))
	*(*int64)(unsafe.Add(mBase, uint32(v513))) = v514 + int64(1)
	if l2 != 0 {
		goto L154
	} else {
		goto L155
	}
L2:
	;
	v424 = float64(0)
	v425 = F_BuildCachedPlan(m, l0, v414, l1, l3)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L12
	} else {
		goto L141
	}
L3:
	;
	v156 = F_BuildCachedPlan(m, l0, v19, int32(0), l3)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L57
	}
L4:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v138 == int32(0) {
		goto L3
	} else {
		goto L53
	}
L5:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v502 = int32(136)
	v503 = int32(0)
	v504 = v134
	goto L1
L6:
	;
	v132 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)) = uint8(v132)
	goto L4
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L50
	}
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v16 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v19 = F_RevalidateCachedQuery(m, l0, l3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	return int32(0)
L13:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v23 != 0 {
		v414 = v19
		goto L2
	} else {
		goto L14
	}
L14:
	;
	if l1 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v81 == int32(0) {
		goto L3
	} else {
		goto L36
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	switch v61 - int32(1) {
	case 0:
		goto L15
	case 1:
		v414 = v19
		goto L2
	default:
		goto L31
	}
L18:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	switch v30 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v34 = int32(1)
		goto L22
	default:
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v35 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L21:
	;
	if v34 != 0 {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v34 = int32(0)
	goto L22
L24:
	;
	goto L15
L25:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v39 != int32(6) {
		v54 = int32(1)
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v54&int32(1) == int32(0) {
		goto L15
	} else {
		goto L30
	}
L27:
	;
	goto L26
L28:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = v44 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v46) {
		v54 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v54 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v46)) % 64)))
	goto L27
L30:
	;
	goto L17
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v64&int32(512) != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	if v64&int32(1024) != 0 {
		v414 = v19
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v69 < int64(5) {
		v414 = v19
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v72 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	v73 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.F64_lt(v72, base.F64_div(v73, base.F64_convert_i64_u(v69))) == int32(0) {
		v414 = v19
		goto L2
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v84 != int32(1) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+16)))
	if v87 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v91 != v92 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_AcquireExecutorLocks(m, v99, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v94&int32(1) == int32(0) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)))
	if v103 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	if v106 == int32(0) {
		goto L5
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	F_AcquireExecutorLocks(m, v115, int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L49
	}
L47:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[611]))
	if v106 == v110 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v81)+10)) = uint8(v112)
	goto L46
L49:
	;
	goto L4
L50:
	;
	F_errmsg_internal(m, int32(277251), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(487984), int32(1292), int32(277485))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+28))
	v145 = v143 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v138)+28)) = v145
	if v145 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = int32(0)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+8)))
	if v149 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v138)+32))
	F_MemoryContextDelete(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	goto L3
L57:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v158 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v156
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156)+28))
	v176 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v156)+28)) = v175 + v176
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v156)+32))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v180 == v176 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+28))
	v165 = v163 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+28)) = v165
	if v165 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = int32(0)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+8)))
	if v169 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)+32))
	F_MemoryContextDelete(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v255 == int32(0) {
		v349 = v13
		goto L101
	} else {
		goto L102
	}
L64:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	if v188 != v184 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	if v225 != v221 {
		goto L85
	} else {
		goto L86
	}
L67:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+9)) = uint8(v218)
	goto L63
L68:
	;
	if v188 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	if v184 != 0 {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	if v193 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v192 == int32(0) {
		goto L71
	} else {
		goto L77
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193)+28)) = v192
	goto L73
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = v192
	goto L73
L77:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v198
	goto L71
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = v184
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v184)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+28)) = v205
	if v205 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v179)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = int32(0)
	goto L70
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205)+24)) = v179
	goto L83
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = v179
	goto L67
L84:
	;
	goto L63
L85:
	;
	if v225 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	if v221 != 0 {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	if v230 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v229 == int32(0) {
		goto L88
	} else {
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230)+28)) = v229
	goto L90
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+20)) = v229
	goto L90
L94:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v179)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v229)+24)) = v235
	goto L88
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = v221
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+28)) = v242
	if v242 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v179)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v179)+16)) = int32(0)
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = v179
	goto L100
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = v179
	goto L84
L101:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+112)) = v349
	v353 = int32(0)
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v354 != 0 {
		v414 = v353
		goto L2
	} else {
		goto L119
	}
L102:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v258 <= int32(0) {
		v349 = v13
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v261 = int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	if v258 == v261 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v258&v261 == int32(0) {
		v349 = v322
		goto L101
	} else {
		goto L117
	}
L105:
	;
	v315 = int32(0)
	v322 = v13
	goto L104
L106:
	;
	goto L107
L107:
	;
	v269 = int32(0)
	v275 = v269
	v276 = v269
	v283 = v13
	goto L108
L108:
	;
	v288 = v263 + v276<<(uint(int32(2))%32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v290 != int32(6) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v315 = v306
	v322 = v304
	goto L104
L110:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289)+36))
	v294 = *(*float64)(unsafe.Add(mBase, uint32(v293)+16))
	v296 = base.F64_add(v283, v294)
	goto L112
L111:
	;
	v296 = v283
	goto L112
L112:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v298 != int32(6) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+36))
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v301)+16))
	v304 = base.F64_add(v296, v302)
	goto L115
L114:
	;
	v304 = v296
	goto L115
L115:
	;
	v305 = int32(2)
	v306 = v276 + v305
	v308 = v275 + v305
	if v308 != v258&int32(2147483646) {
		v275 = v308
		v276 = v306
		v283 = v304
		goto L108
	} else {
		goto L116
	}
L116:
	;
	goto L109
L117:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v263+v315<<(uint(int32(2))%32))))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if v331 == int32(6) {
		v349 = v322
		goto L101
	} else {
		goto L118
	}
L118:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330)+36))
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v334)+16))
	v349 = base.F64_add(v322, v335)
	goto L101
L119:
	;
	v355 = int32(136)
	if l1 == int32(0) {
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v358 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	switch v393 - int32(1) {
	case 0:
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	case 1:
		v414 = v353
		goto L2
	default:
		goto L135
	}
L122:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	switch v362 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v366 = int32(1)
		goto L126
	default:
		goto L127
	}
L123:
	;
	goto L124
L124:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v367 == int32(0) {
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	} else {
		goto L129
	}
L125:
	;
	if v366 != 0 {
		goto L121
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	v366 = int32(0)
	goto L126
L128:
	;
	v502 = v355
	v503 = v353
	v504 = v156
	goto L1
L129:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v371 != int32(6) {
		v386 = int32(1)
		goto L131
	} else {
		goto L132
	}
L130:
	;
	if v386&int32(1) == int32(0) {
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	} else {
		goto L134
	}
L131:
	;
	goto L130
L132:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v367)+28))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v378 = v376 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v378) {
		v386 = int32(0)
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v386 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v378)) % 64)))
	goto L131
L134:
	;
	goto L121
L135:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v396&int32(512) != 0 {
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	} else {
		goto L136
	}
L136:
	;
	if v396&int32(1024) != 0 {
		v414 = v353
		goto L2
	} else {
		goto L137
	}
L137:
	;
	v401 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v401 < int64(5) {
		v414 = v353
		goto L2
	} else {
		goto L138
	}
L138:
	;
	v404 = *(*float64)(unsafe.Add(mBase, uint32(l0)+112))
	v405 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.F64_lt(v404, base.F64_div(v405, base.F64_convert_i64_u(v401))) != 0 {
		v502 = v355
		v503 = v353
		v504 = v156
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v414 = v353
	goto L2
L140:
	;
	v493 = *(*float64)(unsafe.Add(mBase, uint32(l0)+120))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+120)) = base.F64_add(v490, v493)
	v502 = int32(128)
	v503 = int32(1)
	v504 = v425
	goto L1
L141:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
	if v427 == int32(0) {
		v490 = v424
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
	if v430 <= int32(0) {
		v490 = v424
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v435 = *(*float64)(unsafe.Add(mBase, _consts[601]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v427)+12))
	v444 = int32(0)
	v451 = v424
	goto L144
L144:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v438+v444<<(uint(int32(2))%32))))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)+4))
	if v458 != int32(6) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	v490 = v474
	goto L140
L146:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457)+36))
	v462 = *(*float64)(unsafe.Add(mBase, uint32(v461)+16))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v457)+44))
	if v464 != 0 {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v474 = v451
	goto L148
L148:
	;
	v476 = v444 + int32(1)
	if v430 != v476 {
		v444 = v476
		v451 = v474
		goto L144
	} else {
		goto L152
	}
L149:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v470 = base.F64_convert_i32_s(v465 + int32(1))
	goto L151
L150:
	;
	v470 = float64(1)
	goto L151
L151:
	;
	v474 = base.F64_add(base.F64_mul(base.F64_mul(v435, float64(1000)), v470), base.F64_add(v451, v462))
	goto L148
L152:
	;
	goto L145
L153:
	;
	if v503 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L154:
	;
	F_ResourceOwnerEnlarge(m, l2)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L12
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v504)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+28)) = v527 + int32(1)
	goto L153
L157:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v504)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+28)) = v520 + int32(1)
	F_ResourceOwnerRemember(m, l2, v504, int32(1710508))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L12
	} else {
		goto L158
	}
L158:
	;
	goto L153
L159:
	;
	return v504
L160:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+94)))
	if v533 != int32(1) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v504)+32))
	v538 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v536)+16))
	if v542 != v538 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v572 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v504)+9)) = uint8(v572)
	goto L159
L163:
	;
	if v542 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	goto L162
L166:
	;
	if v538 != 0 {
		goto L173
	} else {
		goto L174
	}
L167:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v536)+28))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v536)+24))
	if v547 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v546 == int32(0) {
		goto L166
	} else {
		goto L172
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v547)+28)) = v546
	goto L168
L170:
	;
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v542)+20)) = v546
	goto L168
L172:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v536)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v546)+24)) = v552
	goto L166
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v536)+16)) = v538
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v538)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+28)) = v559
	if v559 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v536)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v536)+16)) = int32(0)
	goto L165
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+24)) = v536
	goto L178
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v538)+20)) = v536
	goto L162
}
func F_RevalidateCachedQuery(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int64
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
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
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v3 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v8 != 0 {
		v504 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L20
	} else {
		goto L184
	}
L2:
	;
	return v504
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v43 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch v13 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v17 = int32(1)
		goto L9
	default:
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 == int32(0) {
		v504 = v3
		goto L2
	} else {
		goto L12
	}
L8:
	;
	if v17 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L8
L10:
	;
	v17 = int32(0)
	goto L9
L11:
	;
	v504 = v3
	goto L2
L12:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 != int32(6) {
		v37 = int32(1)
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v37&int32(1) == int32(0) {
		v504 = v3
		goto L2
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = v27 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v29) {
		v37 = int32(0)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v37 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v29)) % 64)))
	goto L14
L17:
	;
	goto L4
L18:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v59 != int32(1) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v47 = F_SearchPathMatchesCurrentEnvironment(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v47 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v53 == v51 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+10)) = uint8(v56)
	goto L18
L24:
	;
	v208 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = v208
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v210)
	v213 = l0 + int32(68)
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v208
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v216 != 0 {
		goto L84
	} else {
		goto L85
	}
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v78 == int32(0) {
		v504 = v3
		goto L2
	} else {
		goto L33
	}
L26:
	;
	if v59 == int32(0) {
		goto L24
	} else {
		goto L32
	}
L27:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v62 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v66 != v67 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1165])))
	if v69 != v71 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v73&int32(1) != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	goto L25
L33:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		v504 = v3
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v87 = int32(0)
	goto L35
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v87<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v97 == int32(6) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)))
	if v138&int32(1) != 0 {
		v504 = v3
		goto L2
	} else {
		goto L58
	}
L37:
	;
	v135 = v87 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v135 < v136 {
		v87 = v135
		goto L35
	} else {
		goto L57
	}
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	v102 = v100
	goto L42
L39:
	;
	v129 = v96
	goto L40
L40:
	;
	F_ScanQueryForLocks(m, v129, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L20
	} else {
		goto L56
	}
L41:
	;
	if v126 == int32(0) {
		goto L37
	} else {
		goto L55
	}
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	switch v104 - int32(241) {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L48
	}
L44:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+28))
	v102 = v124
	goto L42
L45:
	;
	v126 = v122
	goto L41
L46:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v119 == int32(6) {
		v123 = v118
		goto L44
	} else {
		goto L54
	}
L47:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v115 == int32(6) {
		v123 = v114
		goto L44
	} else {
		goto L53
	}
L48:
	;
	if v104 != int32(201) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v126 = int32(0)
	goto L41
L50:
	;
	goto L51
L51:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v111 != int32(6) {
		v122 = v110
		goto L45
	} else {
		goto L52
	}
L52:
	;
	v123 = v110
	goto L44
L53:
	;
	v122 = v114
	goto L45
L54:
	;
	v122 = v118
	goto L45
L55:
	;
	v129 = v126
	goto L40
L56:
	;
	goto L37
L57:
	;
	goto L36
L58:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v141 == int32(0) {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v144 <= int32(0) {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	v150 = int32(0)
	goto L61
L61:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v150<<(uint(int32(2))%32))))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v160 == int32(6) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L24
L63:
	;
	v198 = v150 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v198 < v199 {
		v150 = v198
		goto L61
	} else {
		goto L83
	}
L64:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v165 = v163
	goto L68
L65:
	;
	v192 = v159
	goto L66
L66:
	;
	F_ScanQueryForLocks(m, v192, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L20
	} else {
		goto L82
	}
L67:
	;
	if v189 == int32(0) {
		goto L63
	} else {
		goto L81
	}
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	switch v167 - int32(241) {
	case 0:
		goto L73
	case 1:
		goto L72
	default:
		goto L74
	}
L70:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
	v165 = v187
	goto L68
L71:
	;
	v189 = v185
	goto L67
L72:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v182 == int32(6) {
		v186 = v181
		goto L70
	} else {
		goto L80
	}
L73:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v178 == int32(6) {
		v186 = v177
		goto L70
	} else {
		goto L79
	}
L74:
	;
	if v167 != int32(201) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v189 = int32(0)
	goto L67
L76:
	;
	goto L77
L77:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v174 != int32(6) {
		v185 = v173
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v186 = v173
	goto L70
L79:
	;
	v185 = v177
	goto L71
L80:
	;
	v185 = v181
	goto L71
L81:
	;
	v192 = v189
	goto L66
L82:
	;
	goto L63
L83:
	;
	goto L62
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = int32(0)
	F_MemoryContextDelete(m, v216)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L20
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v221 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v240 = base.B2i32(v238 != int32(0))
	goto L93
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = int32(0)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v221)+28))
	v228 = v226 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v221)+28)) = v228
	if v228 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = int32(0)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+8)))
	if v232 != 0 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v221)+32))
	F_MemoryContextDelete(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L20
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	if v240 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v243 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L20
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v247 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	F_PushActiveSnapshot(m, v243)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L20
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v275 != 0 {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v248 = F_copyObjectImpl(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v259 = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v260 == v259 {
		v274 = v259
		goto L99
	} else {
		goto L109
	}
L103:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v251 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v253 = F_pg_analyze_and_rewrite_withcb(m, v248, v250, v251, v252, l1)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L20
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v257 = F_pg_analyze_and_rewrite_fixedparams(m, v248, v250, v255, v256, l1)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L20
	} else {
		goto L108
	}
L107:
	;
	v274 = v253
	goto L99
L108:
	;
	v274 = v257
	goto L99
L109:
	;
	v263 = F_copyObjectImpl(m, v260)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	F_AcquireRewriteLocks(m, v263, int32(1), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L20
	} else {
		goto L111
	}
L111:
	;
	v269 = F_pg_rewrite_query(m, v263)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L20
	} else {
		goto L112
	}
L112:
	;
	v274 = v269
	goto L99
L113:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	m.T0[v275].(func(*base.Module, int32, int32))(m, v274, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L20
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v240 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	goto L115
L117:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L20
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v283 = F_ChoosePortalStrategy(m, v274)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L20
	} else {
		goto L129
	}
L120:
	;
	goto L119
L121:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v437 = F_AllocSetContextCreateInternal(m, v432, int32(16683), int32(0), int32(1024), int32(8388608))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L20
	} else {
		goto L163
	}
L122:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v418 != 0 {
		goto L159
	} else {
		goto L160
	}
L123:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v341 != 0 {
		goto L139
	} else {
		goto L140
	}
L124:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v331 == int32(0) {
		goto L121
	} else {
		goto L137
	}
L125:
	;
	if v323 != 0 {
		goto L123
	} else {
		goto L136
	}
L126:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+28))
	v314 = F_UtilityTupleDescriptor(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L20
	} else {
		goto L135
	}
L127:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v294 = int32(0)
	goto L131
L128:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+76))
	v288 = F_ExecCleanTypeFromTL(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L20
	} else {
		goto L130
	}
L129:
	;
	switch v283 {
	case 0, 2:
		goto L128
	case 1:
		goto L127
	case 3:
		goto L126
	default:
		goto L124
	}
L130:
	;
	v323 = v288
	goto L125
L131:
	;
	v301 = int32(1)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v294<<(uint(int32(2))%32)+v290)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+24)))
	if v305 != v301 {
		v294 = v294 + v301
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)+96))
	v309 = F_ExecCleanTypeFromTL(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L20
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v323 = v309
	goto L125
L135:
	;
	v323 = v314
	goto L125
L136:
	;
	goto L124
L137:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v334 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v335 = int32(4464496)
	v336 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v338
	v413 = v336
	v417 = int32(0)
	goto L122
L139:
	;
	v342 = int32(0)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	if v346 != v347 {
		v398 = v342
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v402 != 0 {
		goto L1
	} else {
		goto L157
	}
L142:
	;
	if v398 != 0 {
		goto L121
	} else {
		goto L156
	}
L143:
	;
	goto L142
L144:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v349 != v350 {
		v398 = v342
		goto L143
	} else {
		goto L145
	}
L145:
	;
	if v346 <= int32(0) {
		v398 = int32(1)
		goto L143
	} else {
		goto L146
	}
L146:
	;
	v356 = v346 << (uint(int32(4)) % 32)
	v358 = int32(20)
	v364 = int32(0)
	goto L147
L147:
	;
	v371 = v364 * int32(100)
	v372 = v323 + v356 + v358 + v371
	v373 = int32(4)
	v375 = v371 + (v341 + v356 + v358)
	v378 = F_strcmp(m, v372+v373, v375+v373)
	mBase = m.M
	if v378 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v398 = int32(0)
	goto L143
L149:
	;
	goto L148
L150:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v372)+68))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v375)+68))
	if v379 != v380 {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v372)+76))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v375)+76))
	if v382 != v383 {
		goto L149
	} else {
		goto L152
	}
L152:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v372)+96))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v375)+96))
	if v385 != v386 {
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+91)))
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+91)))
	if v388 != v389 {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v391 = int32(1)
	v393 = v364 + v391
	if v346 != v393 {
		v364 = v393
		goto L147
	} else {
		goto L155
	}
L155:
	;
	v398 = v391
	goto L143
L156:
	;
	goto L141
L157:
	;
	v403 = int32(4464496)
	v404 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v406
	v408 = F_CreateTupleDescCopy(m, v323)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L20
	} else {
		goto L158
	}
L158:
	;
	v413 = v404
	v417 = v408
	goto L122
L159:
	;
	F_FreeTupleDesc(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L20
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v417
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v413
	goto L121
L162:
	;
	goto L161
L163:
	;
	v439 = int32(4464496)
	v440 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v437
	v443 = F_copyObjectImpl(m, v274)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	F_extract_query_dependencies(m, v443, l0-int32(-64), v213, l0+int32(85))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[31]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v452
	v455 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1165])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v455)
	v457 = F_GetSearchPathMatcher(m, v437)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L20
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v457
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v440
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	if v466 != v462 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v496 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v496)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v437
	v504 = v274
	goto L2
L168:
	;
	if v466 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	goto L167
L171:
	;
	if v462 != 0 {
		goto L178
	} else {
		goto L179
	}
L172:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v437)+28))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v437)+24))
	if v471 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	if v470 == int32(0) {
		goto L171
	} else {
		goto L177
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+28)) = v470
	goto L173
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v466)+20)) = v470
	goto L173
L177:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v437)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v470)+24)) = v476
	goto L171
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+16)) = v462
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v462)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v437)+28)) = v483
	if v483 != 0 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v437)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+16)) = int32(0)
	goto L170
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v483)+24)) = v437
	goto L183
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462)+20)) = v437
	goto L167
L184:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L20
	} else {
		goto L185
	}
L185:
	;
	F_errmsg(m, int32(359807), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L20
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(487984), int32(860), int32(16759))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L20
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
