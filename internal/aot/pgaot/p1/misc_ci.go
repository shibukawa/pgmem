package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cidr_abbrev(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v529 int32
	_ = v529
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26&v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = v24
	goto L5
L4:
	;
	v29 = int32(4)
	goto L5
L5:
	;
	v30 = v20 + v29
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v32 = int32(2)
	v33 = v30 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v35 = int32(50)
	v36 = m.G0
	v38 = v36 - int32(192)
	m.G0 = v38
	switch v31 - v32 {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L6:
	;
	m.G0 = v38 + int32(192)
	if v529 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L7:
	;
	v529 = int32(0)
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(5)
	goto L7
L9:
	;
	if base.Ui32(int32(129)) <= base.Ui32(v34) {
		goto L39
	} else {
		goto L40
	}
L10:
	;
	if base.Ui32(int32(33)) <= base.Ui32(v34) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L12:
	;
	goto L13
L13:
	;
	if v34 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L15:
	;
	if base.Ui32(v149) < base.Ui32(int32(5)) {
		goto L14
	} else {
		goto L37
	}
L16:
	;
	v49 = int32(48)
	*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v49)
	v146 = v15 + int32(-63)
	v149 = int32(49)
	goto L15
L17:
	;
	goto L18
L18:
	;
	if base.Ui32(v34) < base.Ui32(int32(8)) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v113 = v34 & int32(7)
	if v113 == int32(0) {
		v146 = v102
		v149 = v105
		goto L15
	} else {
		goto L31
	}
L20:
	;
	v99 = v33
	v102 = v17
	v105 = v35
	goto L19
L21:
	;
	goto L22
L22:
	;
	v59 = v33
	v61 = v17
	v63 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v65 = v35
	goto L23
L23:
	;
	if base.Ui32(v65) < base.Ui32(int32(6)) {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	v99 = v77
	v102 = v92
	v105 = v93
	goto L19
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v74
	v77 = v59 + int32(1)
	v81 = F_pg_sprintf(m, v61, int32(_a_F_cidr_abbrev_0), v38+int32(32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v83 = v81 + v61
	if v63 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = v77
	v102 = v83
	v105 = v61 + v65 - v83
	goto L19
L28:
	;
	goto L29
L29:
	;
	v88 = int32(46)
	*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v88)
	v91 = int32(1)
	v92 = v83 + v91
	v93 = v61 + v65 - v92
	if v91 < v63 {
		v59 = v77
		v61 = v92
		v63 = v63 - v91
		v65 = v93
		goto L23
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	if base.Ui32(v105) < base.Ui32(int32(6)) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	if v17 != v102 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v119)
	v123 = v102 + int32(1)
	goto L35
L34:
	;
	v123 = v17
	goto L35
L35:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v125 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v124 & ((v125<<(uint(v113)%32) ^ v125) << (uint(int32(8)-v113) % 32))
	v138 = F_pg_sprintf(m, v123, int32(_a_F_cidr_abbrev_0), v38+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v140 = v138 + v123
	v146 = v140
	v149 = v102 + v105 - v140
	goto L15
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v34
	v160 = F_pg_sprintf(m, v146, int32(_a_F_cidr_abbrev_1), v38)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v529 = v17
	goto L6
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(28)
	goto L7
L40:
	;
	goto L41
L41:
	;
	if v34 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v34
	v423 = F_pg_sprintf(m, v406, int32(_a_F_cidr_abbrev_1), v38+int32(48))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L116
	}
L43:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+114)) = uint8(v186)
	v188 = int32(_a_F_cidr_abbrev_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+112)) = uint16(v188)
	v406 = v38 + int32(112) | int32(2)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v199 = int32(base.Ui32(v34+int32(7)) >> (uint(int32(3)) % 32))
	if v199 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v202 = int32(0)
	v210 = F__emscripten_memset_bulkmem(m, v38+int32(176)+v199, base.I32_extend8_s(v202), int32(16)-v199)
	mBase = m.M
	goto L50
L47:
	;
	v200 = F__emscripten_memcpy_bulkmem(m, v38+int32(176), v33, v199)
	mBase = m.M
	goto L49
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v212 = v34 & int32(7)
	if v212 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v215 = v38 + v199 + int32(175)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v221 = v216 & (int32(-1) << (uint(int32(8)-v212) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v221)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v228 = int32(base.Ui32(v34+int32(15)) >> (uint(int32(4)) % 32))
	if v228 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v231 = int32(2)
	goto L56
L55:
	;
	v231 = v228
	goto L56
L56:
	;
	v236 = int32(0)
	v239 = v202
	v243 = v2
	v244 = v2
	v245 = v2
	goto L57
L57:
	;
	v251 = v38 + int32(176) + v239
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+1)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v252|v253 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v273 = int32(0)
	v277 = base.B2i32(v266 != v273) & base.B2i32(v268 < v266)
	if v277 != 0 {
		goto L71
	} else {
		goto L72
	}
L59:
	;
	v271 = v239 + int32(2)
	if base.Ui32(v271) < base.Ui32(v231<<(uint(int32(1))%32)) {
		v236 = v266
		v239 = v271
		v243 = v267
		v244 = v268
		v245 = v269
		goto L57
	} else {
		goto L70
	}
L60:
	;
	if v236 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v236 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v259 = v243
	goto L65
L64:
	;
	v259 = int32(base.Ui32(v239) >> (uint(int32(1)) % 32))
	goto L65
L65:
	;
	v266 = v236 + int32(1)
	v267 = v259
	v268 = v244
	v269 = v245
	goto L59
L66:
	;
	if v236 <= v244 {
		v266 = v236
		v267 = v243
		v268 = v244
		v269 = v245
		goto L59
	} else {
		goto L69
	}
L67:
	;
	v263 = v244
	v264 = v245
	goto L68
L68:
	;
	v266 = int32(0)
	v267 = v243
	v268 = v263
	v269 = v264
	goto L59
L69:
	;
	v263 = v236
	v264 = v243
	goto L68
L70:
	;
	goto L58
L71:
	;
	v278 = v267
	goto L73
L72:
	;
	v278 = v269
	goto L73
L73:
	;
	v279 = int32(0)
	if v277 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v308 = v38 + int32(112)
	v310 = v38 + int32(176)
	v311 = v273
	goto L87
L75:
	;
	v280 = v266
	goto L77
L76:
	;
	v280 = v268
	goto L77
L77:
	;
	if v280 == v231 {
		v297 = v279
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v278 != 0 {
		v297 = v279
		goto L74
	} else {
		goto L79
	}
L79:
	;
	switch v280 - int32(5) {
	case 0:
		goto L82
	case 1:
		goto L80
	case 2:
		goto L81
	default:
		v297 = v279
		goto L74
	}
L80:
	;
	v297 = int32(1)
	goto L74
L81:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+190)))
	if v290 == int32(0) {
		v297 = v279
		goto L74
	} else {
		goto L85
	}
L82:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+186)))
	if v284 != int32(255) {
		v297 = v279
		goto L74
	} else {
		goto L83
	}
L83:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+187)))
	if v287 == int32(255) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v297 = v279
	goto L74
L85:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+191)))
	if v293 == int32(1) {
		v297 = v279
		goto L74
	} else {
		goto L86
	}
L86:
	;
	goto L80
L87:
	;
	if v280 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v406 = v400
	goto L42
L89:
	;
	v403 = v311 + int32(1)
	if v403 != v231 {
		v308 = v400
		v310 = v401
		v311 = v403
		goto L87
	} else {
		goto L115
	}
L90:
	;
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v311))&v297 != 0 {
		goto L100
	} else {
		goto L101
	}
L91:
	;
	if v311 < v278 {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	if v278+v280 <= v311 {
		goto L90
	} else {
		goto L93
	}
L93:
	;
	if v311 == v278 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v326 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v326)
	v330 = v308 + int32(1)
	goto L96
L95:
	;
	v330 = v308
	goto L96
L96:
	;
	if v311 == v231-int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v332 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v332)
	v336 = v330 + int32(1)
	goto L99
L98:
	;
	v336 = v330
	goto L99
L99:
	;
	v400 = v336
	v401 = v310 + int32(2)
	goto L89
L100:
	;
	if v311 == int32(6) {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	if v38+int32(112) == v308 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	v346 = int32(58)
	goto L105
L104:
	;
	v346 = int32(46)
	goto L105
L105:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v346)
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v348
	v351 = v308 + int32(1)
	v355 = F_pg_sprintf(m, v351, int32(_a_F_cidr_abbrev_0), v38+int32(80))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v357 = v355 + v351
	if base.Ui32(int32(120)) < base.Ui32(v34) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v362 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v357))) = uint8(v362)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v364
	v367 = v357 + int32(1)
	v371 = F_pg_sprintf(m, v367, int32(_a_F_cidr_abbrev_0), v38-int32(-64))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	if v311 != int32(7) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v400 = v357
	v401 = v310 + int32(1)
	goto L89
L110:
	;
	v400 = v371 + v367
	v401 = v310 + int32(2)
	goto L89
L111:
	;
	v385 = v38 + int32(112)
	goto L113
L112:
	;
	v381 = int32(58)
	*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v381)
	v385 = v308 + int32(1)
	goto L113
L113:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+1)))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v386 | v387<<(uint(int32(8))%32)
	v397 = F_pg_sprintf(m, v385, int32(_a_F_cidr_abbrev_3), v38+int32(96))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v400 = v397 + v385
	v401 = v310 + int32(2)
	goto L89
L115:
	;
	goto L88
L116:
	;
	v427 = F_strlen(m, v38+int32(112))
	mBase = m.M
	if base.Ui32(v427+int32(1)) <= base.Ui32(int32(50)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v433 = v38 + int32(112)
	if (v433^v17)&int32(3) != 0 {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cidr_abbrev[0])) = int32(35)
	goto L7
L120:
	;
	v529 = v17
	goto L6
L121:
	;
	goto L120
L122:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v488))) = uint8(v487)
	if v487&int32(255) == int32(0) {
		goto L121
	} else {
		goto L137
	}
L123:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	v486 = v433
	v487 = v439
	v488 = v17
	goto L122
L124:
	;
	goto L125
L125:
	;
	if v433&int32(3) != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v443 = v433
	v445 = v17
	goto L129
L127:
	;
	v457 = v433
	v459 = v17
	goto L128
L128:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	v464 = int32(-2139062144)
	if (int32(16843008)-v461|v461)&v464 != v464 {
		v486 = v457
		v487 = v461
		v488 = v459
		goto L122
	} else {
		goto L133
	}
L129:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	*(*uint8)(unsafe.Add(mBase, uint32(v445))) = uint8(v446)
	if v446 == int32(0) {
		goto L121
	} else {
		goto L131
	}
L130:
	;
	v457 = v453
	v459 = v451
	goto L128
L131:
	;
	v450 = int32(1)
	v451 = v445 + v450
	v453 = v443 + v450
	if v453&int32(3) != 0 {
		v443 = v453
		v445 = v451
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v469 = v457
	v470 = v461
	v471 = v459
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v470
	v473 = int32(4)
	v474 = v471 + v473
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	v477 = v469 + v473
	v481 = int32(-2139062144)
	if (v475|(int32(16843008)-v475))&v481 == v481 {
		v469 = v477
		v470 = v475
		v471 = v474
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v486 = v477
	v487 = v475
	v488 = v474
	goto L122
L136:
	;
	goto L135
L137:
	;
	v495 = v486
	v497 = v488
	goto L138
L138:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v497)+1)) = uint8(v498)
	v500 = int32(1)
	if v498 != 0 {
		v495 = v495 + v500
		v497 = v497 + v500
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L121
L140:
	;
	goto L139
L141:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v564 = F_cstring_to_text(m, v17)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L148
	}
L144:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_cidr_abbrev_4), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_cidr_abbrev_5), int32(1217), int32(_a_F_cidr_abbrev_6))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	m.G0 = v17 - int32(-64)
	return v564
}
func F_cidr_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_out(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_cidr_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_network_send(m, v3, int32(1))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_cidr_set_masklen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v17 == int32(-1) {
			v22 = int32(1)
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v26 = v24 & v22
			if v26 != 0 {
				v27 = v22
			} else {
				v27 = int32(4)
			}
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v27))))
			if v29 == int32(2) {
				v32 = int32(32)
			} else {
				v32 = int32(128)
			}
			v38 = v26
			v39 = v32
			if v38 != 0 {
				v44 = int32(1)
			} else {
				v44 = int32(4)
			}
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
			if v46 == int32(2) {
				v49 = int32(32)
			} else {
				v49 = int32(128)
			}
			if base.Ui32(v49) < base.Ui32(v39) {
				v139 = v39
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
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
				v52 = F_palloc0(m, int32(22))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					v58 = v56 & v54
					if v58 != 0 {
						v59 = v54
					} else {
						v59 = int32(4)
					}
					v61 = int32(1)
					v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
					if v63&v61 != 0 {
						v66 = v61
					} else {
						v66 = int32(4)
					}
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v66))))
					*(*uint8)(unsafe.Add(mBase, uint32(v52+v59))) = uint8(v68)
					v71 = v52 + int32(1)
					v73 = v52 + int32(4)
					if v58 != 0 {
						v74 = v71
					} else {
						v74 = v73
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)) = uint8(v39)
					if v39 == int32(0) {
					} else {
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v78&int32(1) != 0 {
							v81 = v71
						} else {
							v81 = v73
						}
						v84 = int32(1)
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v88&v84 != 0 {
							v91 = v13 + v84
						} else {
							v91 = v13 + int32(4)
						}
						v97 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
						if v97 != 0 {
							v98 = F__emscripten_memcpy_bulkmem(m, v81+int32(2), v91+int32(2), v97)
							mBase = m.M
						} else {
						}
						v101 = v39 & int32(7)
						if v101 == int32(0) {
						} else {
							v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v106&int32(1) != 0 {
								v109 = v71
							} else {
								v109 = v73
							}
							v112 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v109 + int32(2)
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
							v116 = v113 & (int32(-256) >> (uint(v101) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v116)
						}
					}
					v122 = int32(1)
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
					if v124&v122 != 0 {
						v127 = v122
					} else {
						v127 = int32(4)
					}
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v127))))
					if v129 == int32(2) {
						v132 = int32(40)
					} else {
						v132 = int32(88)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v132
					m.G0 = v10 + int32(16)
					return v52
				}
			}
		} else {
			if v17 < int32(0) {
				v139 = v17
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v146 = m.ExcPending
					if v146 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
						F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
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
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v38 = v35 & int32(1)
				v39 = v17
				if v38 != 0 {
					v44 = int32(1)
				} else {
					v44 = int32(4)
				}
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v44))))
				if v46 == int32(2) {
					v49 = int32(32)
				} else {
					v49 = int32(128)
				}
				if base.Ui32(v49) < base.Ui32(v39) {
					v139 = v39
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v139
							F_errmsg(m, int32(_a_F_cidr_set_masklen_0), v10)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_cidr_set_masklen_1), int32(357), int32(_a_F_cidr_set_masklen_2))
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
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
					v52 = F_palloc0(m, int32(22))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = int32(1)
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						v58 = v56 & v54
						if v58 != 0 {
							v59 = v54
						} else {
							v59 = int32(4)
						}
						v61 = int32(1)
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						if v63&v61 != 0 {
							v66 = v61
						} else {
							v66 = int32(4)
						}
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v66))))
						*(*uint8)(unsafe.Add(mBase, uint32(v52+v59))) = uint8(v68)
						v71 = v52 + int32(1)
						v73 = v52 + int32(4)
						if v58 != 0 {
							v74 = v71
						} else {
							v74 = v73
						}
						*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)) = uint8(v39)
						if v39 == int32(0) {
						} else {
							v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
							if v78&int32(1) != 0 {
								v81 = v71
							} else {
								v81 = v73
							}
							v84 = int32(1)
							v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
							if v88&v84 != 0 {
								v91 = v13 + v84
							} else {
								v91 = v13 + int32(4)
							}
							v97 = int32(base.Ui32(v39+int32(7)) >> (uint(int32(3)) % 32))
							if v97 != 0 {
								v98 = F__emscripten_memcpy_bulkmem(m, v81+int32(2), v91+int32(2), v97)
								mBase = m.M
							} else {
							}
							v101 = v39 & int32(7)
							if v101 == int32(0) {
							} else {
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
								if v106&int32(1) != 0 {
									v109 = v71
								} else {
									v109 = v73
								}
								v112 = int32(base.Ui32(v39)>>(uint(int32(3))%32)) + v109 + int32(2)
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
								v116 = v113 & (int32(-256) >> (uint(v101) % 32))
								*(*uint8)(unsafe.Add(mBase, uint32(v112))) = uint8(v116)
							}
						}
						v122 = int32(1)
						v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
						if v124&v122 != 0 {
							v127 = v122
						} else {
							v127 = int32(4)
						}
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v127))))
						if v129 == int32(2) {
							v132 = int32(40)
						} else {
							v132 = int32(88)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v132
						m.G0 = v10 + int32(16)
						return v52
					}
				}
			}
		}
	}
}
func F_citextcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v8 = int32(1)
	v9 = l0 + v8
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v14 = v12 & v8
	if v14 != 0 {
		v15 = v9
	} else {
		v15 = l0 + int32(4)
	}
	if v12 == int32(1) {
		v18 = int32(4)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
		if v20&int32(254) == int32(2) {
			v29 = v18
		} else {
			v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
		}
		if v20 == int32(1) {
			v32 = v18
		} else {
			v32 = v29
		}
		v43 = v32
	} else {
		v33 = int32(1)
		if v14 != 0 {
			v43 = int32(base.Ui32(v12)>>(uint(v33)%32)) - v33
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v45 = l1 + int32(1)
	v47 = F_str_tolower(m, v15, v43, int32(100))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return int32(0)
	} else {
		v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v55 = v53 & int32(1)
		if v55 != 0 {
			v56 = v45
		} else {
			v56 = l1 + int32(4)
		}
		if v53 == int32(1) {
			v59 = int32(4)
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
			if v61&int32(254) == int32(2) {
				v70 = v59
			} else {
				v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
			}
			if v61 == int32(1) {
				v73 = v59
			} else {
				v73 = v70
			}
			v84 = v73
		} else {
			v74 = int32(1)
			if v55 != 0 {
				v84 = int32(base.Ui32(v53)>>(uint(v74)%32)) - v74
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v86 = F_str_tolower(m, v56, v84, int32(100))
		mBase = m.M
		v87 = m.ExcPending
		if v87 != 0 {
			return int32(0)
		} else {
			v88 = F_strlen(m, v47)
			mBase = m.M
			v89 = F_strlen(m, v86)
			mBase = m.M
			v90 = F_varstr_cmp(m, v47, v88, v86, v89, l2)
			mBase = m.M
			v91 = m.ExcPending
			if v91 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v47)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v86)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						return v90
					}
				}
			}
		}
	}
}
