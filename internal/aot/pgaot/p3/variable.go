package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expandRecordVariable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
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
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v15 = l1
	v19 = int32(0)
	goto L4
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L23
	} else {
		goto L137
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L23
	} else {
		goto L134
	}
L3:
	;
	m.G0 = v12 + int32(160)
	return v548
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v25 = v24 + v19
	v26 = int32(0)
	if v25 <= v26 {
		v73 = l0
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v546 = F_get_expr_result_tupdesc(m, v542, int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L23
	} else {
		goto L133
	}
L6:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+8)))
	if v87 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+v23<<(uint(int32(2))%32)-int32(4))))
	goto L6
L8:
	;
	v32 = v25 & int32(7)
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v25) < base.Ui32(int32(8)) {
		v73 = v47
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v47 = l0
	v50 = v25
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = l0
	v38 = v25
	v40 = v26
	goto L13
L13:
	;
	v41 = int32(1)
	v42 = v38 - v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v45 = v40 + v41
	if v45 != v32 {
		v35 = v43
		v38 = v42
		v40 = v45
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v47 = v43
	v50 = v42
	goto L9
L15:
	;
	goto L14
L16:
	;
	v55 = v47
	v58 = v50
	goto L17
L17:
	;
	v61 = int32(8)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v61 < v58 {
		v55 = v70
		v58 = v58 - v61
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v73 = v70
	goto L7
L19:
	;
	goto L18
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_expandRTE(m, v86, v90, v91, v92, v93, v91, v12+int32(36), v12+int32(156))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	if v256 != int32(2) {
		goto L64
	} else {
		goto L65
	}
L23:
	;
	return int32(0)
L24:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if v104 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v107 = v105
	goto L27
L26:
	;
	v107 = int32(0)
	goto L27
L27:
	;
	v108 = F_CreateTemplateTupleDesc(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v114 = int32(0)
	v118 = int32(1)
	goto L29
L29:
	;
	v122 = int32(0)
	if v111 == v122 {
		v131 = v122
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v110 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v125 <= v114 {
		v131 = v122
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v131 = v127 + v114<<(uint(int32(2))%32)
	goto L31
L34:
	;
	v228 = base.I32_extend16_s(v118)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v139+v114<<(uint(int32(2))%32))))
	v235 = F_exprType(m, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L23
	} else {
		goto L58
	}
L35:
	;
	v141 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v141 < v150 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if base.B2i32(v131 == int32(0))|base.B2i32(v136 <= v114) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	if v139 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v548 = v108
	goto L3
L40:
	;
	v154 = v108 + int32(28)
	v161 = v141
	v162 = v150
	v164 = v141
	goto L44
L41:
	;
	v218 = v141
	v225 = v150
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v108)+20)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v218
	goto L39
L43:
	;
	v218 = v212
	v225 = v191
	goto L42
L44:
	;
	v170 = v154 + v150<<(uint(int32(3))%32) + v161*int32(100)
	v173 = v154 + v161<<(uint(int32(3))%32)
	if v150 != v162 {
		v191 = v162
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v212 = v150
	goto L43
L46:
	;
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173)+2)))
	if v192 <= int32(0) {
		v212 = v161
		goto L43
	} else {
		goto L54
	}
L47:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+7)))
	if v175 != int32(118) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v191 = v161
	goto L46
L49:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+4)))
	if v178 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+6)))
	if v181&int32(6) != 0 {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v173)+2)))
	if v184 <= int32(0) {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+90)))
	if v187 != int32(118) {
		v191 = v150
		goto L46
	} else {
		goto L53
	}
L53:
	;
	goto L48
L54:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+90)))
	if v195 == int32(118) {
		v212 = v161
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+5)))
	v204 = (v164 + v198 - int32(1)) & (int32(0) - v198)
	if int32(_a_F_expandRecordVariable_0) < v204 {
		v212 = v161
		goto L43
	} else {
		goto L56
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v173))) = uint16(v204)
	v210 = v161 + int32(1)
	if v210 != v150 {
		v161 = v210
		v162 = v191
		v164 = v204 + v192
		goto L44
	} else {
		goto L57
	}
L57:
	;
	goto L45
L58:
	;
	v237 = F_exprTypmod(m, v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_TupleDescInitEntry(m, v108, v228, v230, v235, v237, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	v242 = F_exprCollation(m, v234)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L23
	} else {
		goto L61
	}
L61:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v108+v244<<(uint(int32(3))%32)+v228*int32(100))+24)) = v242
	goto L62
L62:
	;
	v252 = int32(1)
	v114 = v114 + v252
	v118 = v118 + v252
	goto L29
L63:
	;
	goto L5
L64:
	;
	switch v256 - int32(1) {
	case 0:
		goto L68
	default:
		v542 = v15
		goto L63
	case 5:
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v531)+12))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v532+v87<<(uint(int32(2))%32)-int32(4))))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v539 == int32(6) {
		v15 = v538
		v19 = v25
		goto L4
	} else {
		goto L132
	}
L67:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+92)))
	if v389 != 0 {
		v542 = v15
		goto L63
	} else {
		goto L98
	}
L68:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+76))
	if v262 != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	if v300 == int32(0) {
		goto L2
	} else {
		goto L82
	}
L70:
	;
	goto L69
L71:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v266 <= int32(0) {
		v300 = int32(0)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v300 = int32(0)
	goto L70
L74:
	;
	v269 = int32(0)
	if v269 < v266 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v272 = v266
	goto L77
L76:
	;
	v272 = v269
	goto L77
L77:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v277 = int32(0)
	goto L78
L78:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v273+v277<<(uint(int32(2))%32))))
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+8)))
	if v286 == v87&int32(_a_F_expandRecordVariable_1) {
		v300 = v285
		goto L70
	} else {
		goto L80
	}
L79:
	;
	goto L73
L80:
	;
	v289 = v277 + int32(1)
	if v289 != v272 {
		v277 = v289
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+26)))
	if v304 == int32(1) {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	if v308 != int32(6) {
		v542 = v307
		goto L63
	} else {
		goto L84
	}
L84:
	;
	v313 = int32(0)
	base.MemoryFill(m, v12+int32(40), v313, int32(116))
	if v25 == v313 {
		v372 = l0
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v372
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v383
	v387 = F_expandRecordVariable(m, v12+int32(36), v307)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L23
	} else {
		goto L97
	}
L86:
	;
	v318 = int32(7)
	v319 = v25 & v318
	if base.Ui32(v318) <= base.Ui32(v25-int32(1)) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v327 = l0
	v332 = int32(0)
	goto L90
L88:
	;
	v349 = l0
	goto L89
L89:
	;
	v359 = v349
	v364 = int32(0)
	goto L94
L90:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v345 = v332 + int32(8)
	if v345 != v25&int32(-8) {
		v327 = v343
		v332 = v345
		goto L90
	} else {
		goto L92
	}
L91:
	;
	if v319 == int32(0) {
		v372 = v343
		goto L85
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v349 = v343
	goto L89
L94:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v370 = v364 + int32(1)
	if v370 != v319 {
		v359 = v368
		v364 = v370
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v372 = v368
	goto L85
L96:
	;
	goto L95
L97:
	;
	v548 = v387
	goto L3
L98:
	;
	v390 = F_GetCTEForRTE(m, l0, v86, v25)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L23
	} else {
		goto L99
	}
L99:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v395 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v398 = int32(76)
	goto L102
L101:
	;
	v398 = int32(96)
	goto L102
L102:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v392+v398)))
	if v400 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	if v438 == int32(0) {
		goto L1
	} else {
		goto L116
	}
L104:
	;
	goto L103
L105:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v400)+4))
	if v404 <= int32(0) {
		v438 = int32(0)
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v438 = int32(0)
	goto L104
L108:
	;
	v407 = int32(0)
	if v407 < v404 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v410 = v404
	goto L111
L110:
	;
	v410 = v407
	goto L111
L111:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	v415 = int32(0)
	goto L112
L112:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v411+v415<<(uint(int32(2))%32))))
	v424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v423)+8)))
	if v424 == v87&int32(_a_F_expandRecordVariable_1) {
		v438 = v423
		goto L104
	} else {
		goto L114
	}
L113:
	;
	goto L107
L114:
	;
	v427 = v415 + int32(1)
	if v427 != v410 {
		v415 = v427
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438)+26)))
	if v442 == int32(1) {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)))
	if v446 != int32(6) {
		v542 = v445
		goto L63
	} else {
		goto L118
	}
L118:
	;
	v451 = int32(0)
	base.MemoryFill(m, v12+int32(40), v451, int32(116))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	v455 = v454 + v25
	if v455 == v451 {
		v514 = l0
		goto L119
	} else {
		goto L120
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v514
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v525
	v529 = F_expandRecordVariable(m, v12+int32(36), v445)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L23
	} else {
		goto L131
	}
L120:
	;
	v458 = int32(7)
	v459 = v455 & v458
	if base.Ui32(v458) <= base.Ui32(v19+v454+v24-int32(1)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v469 = l0
	v471 = int32(0)
	goto L124
L122:
	;
	v491 = l0
	goto L123
L123:
	;
	v501 = v491
	v503 = int32(0)
	goto L128
L124:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	v487 = v471 + int32(8)
	if v487 != v455&int32(-8) {
		v469 = v485
		v471 = v487
		goto L124
	} else {
		goto L126
	}
L125:
	;
	if v459 == int32(0) {
		v514 = v485
		goto L119
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	v491 = v485
	goto L123
L128:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	v512 = v503 + int32(1)
	if v512 != v459 {
		v501 = v510
		v503 = v512
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v514 = v510
	goto L119
L130:
	;
	goto L129
L131:
	;
	v548 = v529
	goto L3
L132:
	;
	v542 = v538
	goto L63
L133:
	;
	v548 = v546
	goto L3
L134:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v566
	F_errmsg_internal(m, int32(_a_F_expandRecordVariable_2), v12+int32(16))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L23
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_expandRecordVariable_3), int32(1603), int32(_a_F_expandRecordVariable_4))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L23
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v584
	F_errmsg_internal(m, int32(_a_F_expandRecordVariable_5), v12)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L23
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_expandRecordVariable_3), int32(1662), int32(_a_F_expandRecordVariable_4))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L23
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 float64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 float64
	_ = v191
	var v195 int32
	_ = v195
	var v196 float32
	_ = v196
	var v199 float32
	_ = v199
	var v202 float32
	_ = v202
	var v205 float32
	_ = v205
	var v207 float64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v228 float64
	_ = v228
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 float64
	_ = v243
	var v248 float32
	_ = v248
	var v250 float64
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v269 float64
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 float32
	_ = v276
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v358 int32
	_ = v358
	v6 = int32(0)
	v14 = float64(0)
	v15 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)) = uint8(v6)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v26 == v6 {
		v358 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v18 + int32(96)
	return v358 & int32(1)
L2:
	;
	v29 = F_get_opcode(m, l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_typlenbyval(m, v57, v18+int32(76), v18+int32(75))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L15
	}
L6:
	;
	if v29 == int32(0) {
		v358 = v6
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v36 = F_get_func_leakproof(m, v29)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v40 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v40 == int32(0) {
		v358 = v6
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = F_get_func_name(m, v29)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v44
	F_errmsg_internal(m, int32(_a_F_get_variable_range_0), v18)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_get_variable_range_1), int32(_a_F_get_variable_range_2), int32(_a_F_get_variable_range_3))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v358 = v6
	goto L1
L15:
	;
	v65 = v18 + int32(8)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = F_get_attstatsslot(m, v65, v66, int32(2), l1, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L22
	}
L16:
	;
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v18)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v348
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v18)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l4))) = v350
	v358 = v339
	goto L1
L17:
	;
	F_free_attstatsslot(m, v18+int32(8))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L3
	} else {
		goto L59
	}
L18:
	;
	v302 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+76)))
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)))
	F_get_stats_slot_range(m, v18+int32(8), v29, v18+int32(44), l2, v302, v303, v18+int32(88), v18+int32(80), v18+int32(79))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L58
	}
L19:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	if v167 <= int32(0) {
		v269 = v14
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v153 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v160 = F_get_attstatsslot(m, v18+int32(8), v156, int32(1), v153, int32(3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L43
	}
L21:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v147 = F_get_attstatsslot(m, v18+int32(8), v144, int32(1), int32(0), v141)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L40
	}
L22:
	;
	if v69 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v71 != l2 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v109 = v18 + int32(8)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v114 = F_get_attstatsslot(m, v109, v110, int32(2), int32(0), int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L33
	}
L26:
	;
	F_free_attstatsslot(m, v18+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L32
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v74 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)))
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+76)))
	v81 = F_datumCopy(m, v78, v79, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v84+v85<<(uint(int32(3))%32)-int32(8))))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+76)))
	v94 = F_datumCopy(m, v91, v92, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)) = uint8(v96)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v94
	F_free_attstatsslot(m, v65)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v139 = int32(1)
	v141 = int32(1)
	goto L21
L32:
	;
	goto L25
L33:
	;
	if v114 == int32(0) {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+76)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+75)))
	F_get_stats_slot_range(m, v109, v29, v18+int32(44), l2, v120, v121, v18+int32(88), v18+int32(80), v18+int32(79))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_free_attstatsslot(m, v109)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v132 = int32(1)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)))
	if v134&v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = v132
	goto L39
L38:
	;
	v137 = int32(3)
	goto L39
L39:
	;
	v139 = v134
	v141 = v137
	goto L21
L40:
	;
	if v147 == int32(0) {
		v339 = v139
		goto L16
	} else {
		goto L41
	}
L41:
	;
	if v139&int32(1) != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	goto L19
L43:
	;
	if v160 == int32(0) {
		v339 = v153
		goto L16
	} else {
		goto L44
	}
L44:
	;
	goto L19
L45:
	;
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+22)))
	v276 = *(*float32)(unsafe.Add(mBase, uint32(v273+v274)+8))
	if base.F64_gt(base.F64_add(v269, base.F64_promote_f32(v276)), float64(0.99999)) == v271 {
		v328 = v271
		goto L17
	} else {
		goto L57
	}
L46:
	;
	v171 = v167 & int32(3)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v173 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v167) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v184 = v173
	v189 = v6
	v191 = v14
	goto L50
L48:
	;
	v221 = v173
	v228 = v14
	goto L49
L49:
	;
	v236 = v221
	v242 = v6
	v243 = v228
	goto L54
L50:
	;
	v195 = v172 + v184<<(uint(int32(2))%32)
	v196 = *(*float32)(unsafe.Add(mBase, uint32(v195)))
	v199 = *(*float32)(unsafe.Add(mBase, uint32(v195)+4))
	v202 = *(*float32)(unsafe.Add(mBase, uint32(v195)+8))
	v205 = *(*float32)(unsafe.Add(mBase, uint32(v195)+12))
	v207 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v191, base.F64_promote_f32(v196)), base.F64_promote_f32(v199)), base.F64_promote_f32(v202)), base.F64_promote_f32(v205))
	v208 = int32(4)
	v209 = v184 + v208
	v211 = v189 + v208
	if v211 != v167&int32(2147483644) {
		v184 = v209
		v189 = v211
		v191 = v207
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if v171 == int32(0) {
		v269 = v207
		goto L45
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v221 = v209
	v228 = v207
	goto L49
L54:
	;
	v248 = *(*float32)(unsafe.Add(mBase, uint32(v172+v236<<(uint(int32(2))%32))))
	v250 = base.F64_add(v243, base.F64_promote_f32(v248))
	v251 = int32(1)
	v254 = v242 + v251
	if v254 != v171 {
		v236 = v236 + v251
		v242 = v254
		v243 = v250
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v269 = v250
	goto L45
L56:
	;
	goto L55
L57:
	;
	goto L18
L58:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+79)))
	v328 = v312
	goto L17
L59:
	;
	v339 = v328
	goto L16
}
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v6)
	v22 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1133), v9+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(32)
		return v22
	}
}
