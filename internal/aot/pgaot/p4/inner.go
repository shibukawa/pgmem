package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inner_int_union(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v601)+4))
	v615 = F_ArrayGetNItems(m, v612, v601+int32(16))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L8
	} else {
		goto L122
	}
L2:
	;
	if v56 <= v220 {
		v322 = v219
		goto L71
	} else {
		goto L72
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L67
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L8
	} else {
		goto L63
	}
L5:
	;
	v16 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v20 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return int32(0)
L9:
	;
	if v16 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v21 = F_array_contains_nulls(m, l1)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = l0 + int32(16)
	v26 = F_ArrayGetNItems(m, v23, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L17
	}
L14:
	;
	if v21 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = F_ArrayGetNItems(m, v37, v25)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L22
	}
L17:
	;
	if v26 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = F_ArrayGetNItems(m, v28, l1+int32(16))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	if v31 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v34 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	return v34
L22:
	;
	if v38 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v42 = F_copy_intArrayType(m, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	v44 = int32(0)
	goto L25
L25:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v47 = l1 + int32(16)
	v48 = F_ArrayGetNItems(m, v45, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v44 = v42
	goto L25
L27:
	;
	if v48 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v52 = F_copy_intArrayType(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	v54 = v44
	goto L30
L30:
	;
	if v54 != 0 {
		v601 = v54
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v54 = v52
	goto L30
L32:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v56 = F_ArrayGetNItems(m, v55, v25)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v59 = F_ArrayGetNItems(m, v58, v47)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v61 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = (v64<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L36:
	;
	v71 = v61
	goto L37
L37:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v72 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v82 = (v75<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L40
L39:
	;
	v82 = v72
	goto L40
L40:
	;
	v83 = v56 + v59
	if int32(0) < v83 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v125 = l0 + v71
	v126 = l1 + v82
	v127 = v121 + v122
	v128 = int32(0)
	if base.B2i32(v128 < v56)&base.B2i32(v128 < v59) == v128 {
		goto L49
	} else {
		goto L50
	}
L42:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v121 = v110
	v122 = (v114<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v124 = v113
	goto L41
L43:
	;
	v89 = v83<<(uint(int32(2))%32) + int32(24)
	v90 = F_palloc0(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v105 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v90)+12)) = int32(23)
	*(*int64)(unsafe.Add(mBase, uint32(v90)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v89 << (uint(int32(2)) % 32)
	v110 = v90
	v113 = v90 + int32(8)
	goto L42
L47:
	;
	v108 = v105 + int32(8)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	if v109 != 0 {
		v121 = v105
		v122 = v109
		v124 = v108
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v110 = v105
	v113 = v108
	goto L42
L49:
	;
	v219 = v127
	v220 = int32(0)
	v221 = v128
	goto L2
L50:
	;
	goto L51
L51:
	;
	v138 = v127
	v139 = int32(0)
	v140 = v128
	goto L52
L52:
	;
	v152 = int32(2)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v125+v139<<(uint(v152)%32))))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v126+v140<<(uint(v152)%32))))
	if v155 == v159 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v219 = v176
	v220 = v173
	v221 = v174
	goto L2
L54:
	;
	v176 = v138 + int32(4)
	if v56 <= v173 {
		v219 = v176
		v220 = v173
		v221 = v174
		goto L2
	} else {
		goto L61
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v155
	v162 = int32(1)
	v173 = v139 + v162
	v174 = v140 + v162
	goto L54
L56:
	;
	goto L57
L57:
	;
	if v155 < v159 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v155
	v173 = v139 + int32(1)
	v174 = v140
	goto L54
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v159
	v173 = v139
	v174 = v140 + int32(1)
	goto L54
L61:
	;
	if v174 < v59 {
		v138 = v176
		v139 = v173
		v140 = v174
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
L63:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(156210), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(509295), int32(83), int32(277552))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(156210), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(509295), int32(84), int32(277552))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	if v59 <= v221 {
		v425 = v322
		goto L84
	} else {
		goto L85
	}
L72:
	;
	v236 = (v56 - v220) & int32(3)
	if v236 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v220-v56) {
		v322 = v266
		goto L71
	} else {
		goto L80
	}
L74:
	;
	v266 = v219
	v270 = v220
	goto L73
L75:
	;
	goto L76
L76:
	;
	v240 = v219
	v244 = v220
	v246 = int32(0)
	goto L77
L77:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v125+v244<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v257
	v260 = v240 + int32(4)
	v261 = int32(1)
	v262 = v244 + v261
	v264 = v246 + v261
	if v264 != v236 {
		v240 = v260
		v244 = v262
		v246 = v264
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v266 = v260
	v270 = v262
	goto L73
L79:
	;
	goto L78
L80:
	;
	v289 = v266
	v293 = v270
	goto L81
L81:
	;
	v304 = v293 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v125+v304)))
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v306
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v304+(v125+int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v309
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v304+(v125+int32(8)))))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+8)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v304+(v125+int32(12)))))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+12)) = v315
	v318 = v289 + int32(16)
	v320 = v293 + int32(4)
	if v320 != v56 {
		v289 = v318
		v293 = v320
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v322 = v318
	goto L71
L83:
	;
	goto L82
L84:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v439 != 0 {
		goto L97
	} else {
		goto L98
	}
L85:
	;
	v339 = (v59 - v221) & int32(3)
	if v339 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v221-v59) {
		v425 = v369
		goto L84
	} else {
		goto L93
	}
L87:
	;
	v369 = v322
	v370 = v221
	goto L86
L88:
	;
	goto L89
L89:
	;
	v343 = v322
	v344 = v221
	v347 = int32(0)
	goto L90
L90:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v126+v344<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v343))) = v360
	v363 = v343 + int32(4)
	v364 = int32(1)
	v365 = v344 + v364
	v367 = v347 + v364
	if v367 != v339 {
		v343 = v363
		v344 = v365
		v347 = v367
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v369 = v363
	v370 = v365
	goto L86
L92:
	;
	goto L91
L93:
	;
	v392 = v369
	v393 = v370
	goto L94
L94:
	;
	v407 = v393 << (uint(int32(2)) % 32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v126+v407)))
	*(*int32)(unsafe.Add(mBase, uint32(v392))) = v409
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v407+(v126+int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+4)) = v412
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v407+(v126+int32(8)))))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+8)) = v415
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v407+(v126+int32(12)))))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = v418
	v421 = v392 + int32(16)
	v423 = v393 + int32(4)
	if v423 != v59 {
		v392 = v421
		v393 = v423
		goto L94
	} else {
		goto L96
	}
L95:
	;
	v425 = v421
	goto L84
L96:
	;
	goto L95
L97:
	;
	v447 = v439
	goto L99
L98:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v447 = (v440<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L99
L99:
	;
	v449 = v425 - (v447 + v121)
	v451 = v449 >> (uint(int32(2)) % 32)
	if v451 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v455 = F_construct_empty_array(m, int32(23))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v460 = F_ArrayGetNItems(m, v457, v121+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L104
	}
L103:
	;
	v601 = v455
	goto L1
L104:
	;
	if v460 == v451 {
		v601 = v121
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v463 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v471 = v463
	goto L108
L107:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v471 = (v464<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L108
L108:
	;
	v472 = v471 + v449
	v473 = F_repalloc(m, v121, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v472 << (uint(int32(2)) % 32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v478 <= int32(0) {
		v601 = v473
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473)+16)) = v451
	if v478 == int32(1) {
		v601 = v473
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v485 = v473 + int32(16)
	v486 = int32(1)
	v487 = v478 - v486
	v488 = int32(7)
	v489 = v487 & v488
	if base.Ui32(v488) <= base.Ui32(v478-int32(2)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v513 = v486
	v516 = int32(0)
	goto L115
L113:
	;
	v558 = v486
	goto L114
L114:
	;
	if v489 == int32(0) {
		v601 = v473
		goto L1
	} else {
		goto L118
	}
L115:
	;
	v527 = v513 << (uint(int32(2)) % 32)
	v529 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v485+v527))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(20))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(24))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(28))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(32))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(36))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(40))))) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v527+(v473+int32(44))))) = v529
	v552 = int32(8)
	v553 = v513 + v552
	v555 = v516 + v552
	if v555 != v487&int32(-8) {
		v513 = v553
		v516 = v555
		goto L115
	} else {
		goto L117
	}
L116:
	;
	v558 = v553
	goto L114
L117:
	;
	goto L116
L118:
	;
	v574 = int32(0)
	v575 = v558
	goto L119
L119:
	;
	v591 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v485+v575<<(uint(int32(2))%32)))) = v591
	v596 = v574 + v591
	if v596 != v489 {
		v574 = v596
		v575 = v575 + v591
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v601 = v473
	goto L1
L121:
	;
	goto L120
L122:
	;
	if int32(2) <= v615 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v619 = F__int_unique(m, v601)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L8
	} else {
		goto L126
	}
L124:
	;
	v621 = v601
	goto L125
L125:
	;
	return v621
L126:
	;
	v621 = v619
	goto L125
}
