package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyFromTextOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v69 int64
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v218 int64
	_ = v218
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int64
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int64
	_ = v437
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	v29 = v27
	goto L3
L2:
	;
	v29 = int32(0)
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	if v34 != int64(0) {
		v238 = v34
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L20
	} else {
		goto L143
	}
L5:
	;
	m.G0 = v23 + int32(128)
	return v584
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L20
	} else {
		goto L139
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v238 + int64(1)
	v242 = int32(0)
	v244 = F_CopyReadLine(m, l0, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L54
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v38 == int32(0) {
		v238 = int64(0)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v38 == int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v89 != int32(-1) {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v45 = int32(1)
	goto L13
L12:
	;
	v45 = v38
	goto L13
L13:
	;
	if v45 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v89 = v38
	v92 = int32(0)
	goto L10
L15:
	;
	goto L16
L16:
	;
	v53 = int32(0)
	goto L17
L17:
	;
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v69 + int64(1)
	v74 = F_CopyReadLine(m, l0, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v89 = v82
	v92 = v74
	goto L10
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	if v74 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v79 = v53 + int32(1)
	if v79 != v45 {
		v53 = v79
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	if v92 != 0 {
		goto L51
	} else {
		goto L52
	}
L25:
	;
	v105 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v107 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v105 == int32(0) {
		goto L24
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v105 != v112 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L4
L31:
	;
	v119 = int32(0)
	goto L32
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v135 <= v119 {
		goto L24
	} else {
		goto L34
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L20
	} else {
		goto L47
	}
L34:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v142 = v119 << (uint(int32(2)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142+v143)))
	v150 = v33 + v137<<(uint(int32(3))%32) + v145*int32(100) - int32(72)
	v152 = v119 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v142+v153)))
	if v155 == int32(0) {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v159 = v150 + int32(4)
	if v159|v155 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v174 == int32(0) {
		v119 = v152
		goto L32
	} else {
		goto L46
	}
L37:
	;
	v165 = int32(-1)
	goto L39
L38:
	;
	v165 = int32(0)
	goto L39
L39:
	;
	if v159 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v166 = int32(1)
	goto L42
L41:
	;
	v166 = v165
	goto L42
L42:
	;
	v167 = int32(0)
	if base.B2i32(v159 == v167)|base.B2i32(v155 == v167) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v174 = v166
	goto L45
L44:
	;
	v173 = F_strncmp(m, v159, v155, int32(64))
	mBase = m.M
	v174 = v173
	goto L45
L45:
	;
	goto L36
L46:
	;
	goto L33
L47:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v152
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_0), v23+int32(112))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(867), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L20
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v584 = int32(0)
	goto L5
L52:
	;
	goto L53
L53:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v238 = v218
	goto L7
L54:
	;
	if v244 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v246 == int32(0) {
		v584 = v242
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v249 = int32(1)
	v250 = F_CopyReadAttributesText(m, l0)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L20
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v254 = int32(0)
	if base.B2i32(v29 < v250)&base.B2i32(v254 < v29) == v254 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v259 == int32(0) {
		v584 = v249
		goto L5
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L20
	} else {
		goto L135
	}
L63:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v262 <= int32(0) {
		v584 = v249
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v265 = int32(0)
	if v265 < v250 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v269 = v250
	goto L67
L66:
	;
	v269 = v265
	goto L67
L67:
	;
	v279 = v265
	v288 = int32(0)
	goto L68
L68:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v295 = v279 << (uint(int32(2)) % 32)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295+v296)))
	v303 = v33 + v290<<(uint(int32(3))%32) + v298*int32(100) - int32(72)
	if v279 != v269 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	v584 = int32(1)
	goto L5
L70:
	;
	v537 = v279 + int32(1)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v259)+4))
	if v537 < v538 {
		v279 = v537
		v288 = v534
		goto L68
	} else {
		goto L134
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = int64(0)
	v534 = v525
	goto L70
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L20
	} else {
		goto L126
	}
L73:
	;
	v306 = v298 - int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v295+v252)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v309 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L20
	} else {
		goto L122
	}
L76:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306+v309))))
	if v311 != int32(1) {
		v534 = v288
		goto L70
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v303 + int32(4)
	if v308 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	v319 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v306))) = uint8(v319)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v323 = l2 + v306<<(uint(int32(3))%32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324+v306))))
	if v326 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v30+v306<<(uint(int32(2))%32))))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332)+24))
	v335 = m.T0[v334].(func(*base.Module, int32, int32, int32) int64)(m, v332, l1, l3+v306)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L20
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v340 = v32 + v306*int32(28)
	v343 = v31 + v306<<(uint(int32(2))%32)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v303)+76))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v347 = F_InputFunctionCallSafe(m, v340, v308, v344, v345, v346, v323)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L20
	} else {
		goto L87
	}
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v335
	v525 = v288
	goto L71
L87:
	;
	if v347 != 0 {
		v525 = v288
		goto L71
	} else {
		goto L88
	}
L88:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v349 - int32(1) {
	case 0:
		v374 = v288
		goto L90
	case 1:
		goto L91
	default:
		v379 = v288
		goto L89
	}
L89:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v380 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L90:
	;
	v375 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v375 + int64(1)
	v379 = v374
	goto L89
L91:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v353 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v352)+4)) = uint8(v353)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355+v306))))
	if v357 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v303)+76))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v364 = F_InputFunctionCallSafe(m, v340, int32(0), v361, v362, v363, v323)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L20
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v369 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v306))) = uint8(v369)
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = int64(0)
	if v288 != 0 {
		v379 = v369
		goto L89
	} else {
		goto L97
	}
L95:
	;
	if v364 == int32(0) {
		goto L72
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v374 = v369
	goto L90
L98:
	;
	v383 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v383)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v385 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v463 = int32(1)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v464 - v463 {
	case 0:
		v584 = v463
		goto L5
	case 1:
		v534 = v379
		goto L70
	default:
		v525 = v379
		goto L71
	}
L101:
	;
	v456 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v456)
	goto L100
L102:
	;
	v386 = F_CopyLimitPrintoutLength(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L20
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v428 != int32(1) {
		goto L101
	} else {
		goto L117
	}
L105:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v388 - int32(1) {
	case 0:
		goto L109
	case 1:
		goto L108
	default:
		goto L106
	}
L106:
	;
	F_pfree(m, v386)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L20
	} else {
		goto L116
	}
L107:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v410
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v409
	F_errmsg(m, v407, v23+int32(32))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L20
	} else {
		goto L114
	}
L108:
	;
	v401 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L20
	} else {
		goto L112
	}
L109:
	;
	v393 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L20
	} else {
		goto L110
	}
L110:
	;
	if v393 == int32(0) {
		goto L106
	} else {
		goto L111
	}
L111:
	;
	v407 = int32(_a_F_CopyFromTextOneRow_3)
	v408 = int32(1151)
	goto L107
L112:
	;
	if v401 == int32(0) {
		goto L106
	} else {
		goto L113
	}
L113:
	;
	v407 = int32(_a_F_CopyFromTextOneRow_4)
	v408 = int32(1157)
	goto L107
L114:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), v408, int32(_a_F_CopyFromTextOneRow_5))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	goto L106
L116:
	;
	goto L101
L117:
	;
	v433 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	if v433 == int32(0) {
		goto L101
	} else {
		goto L119
	}
L119:
	;
	v437 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v437
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_6), v23+int32(16))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1166), int32(_a_F_CopyFromTextOneRow_5))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L20
	} else {
		goto L121
	}
L121:
	;
	goto L101
L122:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L20
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v303 + int32(4)
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_7), v23)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L20
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1019), int32(_a_F_CopyFromTextOneRow_5))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L20
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L20
	} else {
		goto L127
	}
L127:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v493 = F_format_type_be(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L20
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v493
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_8), v23-int32(-64))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L20
	} else {
		goto L129
	}
L129:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v503 = F_format_type_be(m, v502)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L20
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v503
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v501
	v510 = F_errdetail(m, int32(_a_F_CopyFromTextOneRow_9), v23+int32(48))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L20
	} else {
		goto L131
	}
L131:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	F_errdatatype(m, v512)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L20
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1117), int32(_a_F_CopyFromTextOneRow_5))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L20
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	goto L69
L135:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L20
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_10), int32(0))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L20
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(1004), int32(_a_F_CopyFromTextOneRow_5))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L20
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L20
	} else {
		goto L140
	}
L140:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v150 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v564
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v152
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_11), v23+int32(96))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L20
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(860), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L20
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L20
	} else {
		goto L144
	}
L144:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v611 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v611)+4))
	v614 = v612
	goto L147
L146:
	;
	v614 = int32(0)
	goto L147
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v105
	F_errmsg(m, int32(_a_F_CopyFromTextOneRow_12), v23+int32(80))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L20
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_CopyFromTextOneRow_1), int32(844), int32(_a_F_CopyFromTextOneRow_2))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L20
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RunFromStore(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
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
	var v57 int64
	_ = v57
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v11 = F_MakeSingleTupleTableSlot(m, v9, int32(_a_F_RunFromStore_0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	m.T0[v17].(func(*base.Module, int32, int32, int32))(m, l3, int32(1), v16)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 == int32(0) {
		v67 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	m.T0[v68].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	v31 = v8
	goto L6
L6:
	;
	v32 = int32(_a_F_RunFromStore_1)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0])) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v39 = F_tuplestore_gettupleslot(m, v37, base.B2i32(l1 == int32(1)), int32(0), v11)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v67 = l2
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RunFromStore[0])) = v33
	if v39 == int32(0) {
		v67 = v31
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v46 = m.T0[v45].(func(*base.Module, int32, int32) int32)(m, v11, l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v46 == int32(0) {
		v67 = v31
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	m.T0[v51].(func(*base.Module, int32))(m, v11)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v57 = v31 + int64(1)
	if base.B2i32(l2 == int64(0))|base.B2i32(v57 != l2) != 0 {
		v31 = v57
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	F_ExecDropSingleTupleTableSlot(m, v11)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	return v67
}
