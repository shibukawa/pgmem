package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ArrayGetIntegerTypmods(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 == int32(2275) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L28
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L24
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 != int32(1) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L7
	} else {
		goto L20
	}
L6:
	;
	v16 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if v16 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_deconstruct_array_builtin(m, l0, int32(2275), v8+int32(12), int32(0), l1)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = F_palloc(m, v26<<(uint(int32(2))%32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v31 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v37 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_pfree(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L19
	}
L15:
	;
	v40 = v37 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40)))
	v45 = F_pg_strtoint32(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29+v40))) = v45
	v49 = v37 + int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v49 < v50 {
		v37 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	m.G0 = v8 + int32(16)
	return v29
L20:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(533704), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(517844), int32(242), int32(182626))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(328409), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(517844), int32(247), int32(182626))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(162277), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(517844), int32(252), int32(182626))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_array_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
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
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
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
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = v15 + int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if v48 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v48 = v45
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v40
	v45 = v41
	goto L2
L4:
	;
	v37 = int32(0)
	if v18 == v37 {
		v45 = v37
		goto L2
	} else {
		goto L14
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	switch v23 - int32(429) {
	case 0:
		goto L7
	case 1:
		goto L6
	default:
		goto L4
	}
L6:
	;
	if v18 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	if v18 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v48 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+168))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v40 = v30
	v41 = int32(1)
	goto L3
L11:
	;
	v48 = int32(2)
	goto L1
L12:
	;
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+368))
	v40 = v35
	v41 = int32(2)
	goto L3
L14:
	;
	v40 = v37
	v41 = v2
	goto L3
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L31
	} else {
		goto L159
	}
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v53 = v52
	goto L20
L19:
	;
	v53 = v2
	goto L20
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v54 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	m.G0 = v15 + int32(16)
	return v584
L22:
	;
	if v53 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v57 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v53 != 0 {
		v584 = v53
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
	v584 = int32(0)
	goto L21
L28:
	;
	v64 = int32(4562080)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	v70 = int32(0)
	v72 = F_initArrayResultArr(m, v69, v70, v67, v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v131 <= int32(0) {
		v584 = v53
		goto L21
	} else {
		goto L46
	}
L31:
	;
	return int32(0)
L32:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v76
	v78 = F_palloc(m, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v81 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v86 = base.I32_div_s(v82+int32(7), int32(8))
	v87 = F_palloc(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	v96 = v78
	goto L36
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v98 != 0 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v86 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v96 = v93
	goto L36
L39:
	;
	v91 = F__emscripten_memcpy_bulkmem(m, v87, v90, v86)
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+16)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+28)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v57)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+48)) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v57)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+40)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+32)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v57)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+56)) = v115
	v117 = int32(-64)
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v57-v117)))
	*(*int64)(unsafe.Add(mBase, uint32(v72-v117))) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v57)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+72)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v57)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+80)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v57)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+84)) = v127
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v65
	v584 = v72
	goto L21
L43:
	;
	v99 = F__emscripten_memcpy_bulkmem(m, v96, v97, v98)
	mBase = m.M
	goto L45
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v53)+28))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
	if v134 == v135 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v221 = int32(4562080)
	v222 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v224
	v226 = v137 + v138
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v227 < v226 {
		goto L67
	} else {
		goto L68
	}
L48:
	;
	v170 = int32(1)
	goto L57
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v134 < int32(2) {
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L31
	} else {
		goto L53
	}
L52:
	;
	v141 = int32(56)
	v145 = int32(32)
	goto L48
L53:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(11928), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(518919), int32(1051), int32(392184))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	v179 = v170 << (uint(int32(2)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v53+v145+v179)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+(v57+v145))))
	if v181 != v183 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L31
	} else {
		goto L63
	}
L59:
	;
	goto L58
L60:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179+(v53+v141))))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v179+(v57+v141))))
	if v186 != v188 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v191 = v170 + int32(1)
	if v134 != v191 {
		v170 = v191
		goto L57
	} else {
		goto L62
	}
L62:
	;
	goto L47
L63:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L31
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(11928), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(518919), int32(1059), int32(392184))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L31
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
	v229 = int32(1)
	if v226&(v226-v229) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v244 != 0 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v237 = v229 << (uint(int32(32)-base.I32_clz(v226)) % 32)
	goto L72
L71:
	;
	v237 = v226
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v240 = F_repalloc(m, v239, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L31
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v240
	goto L69
L74:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v247 = v245 + v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v248 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v566 != 0 {
		goto L156
	} else {
		goto L157
	}
L77:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v431 = int32(0)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v432 <= v431 {
		goto L125
	} else {
		goto L126
	}
L78:
	;
	v251 = int32(1)
	v253 = int32(256)
	v255 = v247 + v251
	if v255 <= v253 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	if v247 <= v405 {
		v426 = v248
		goto L77
	} else {
		goto L119
	}
L81:
	;
	v258 = v253
	goto L83
L82:
	;
	v258 = v255
	goto L83
L83:
	;
	if v258&(v258-int32(1)) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v265 = v251 << (uint(int32(32)-base.I32_clz(v258)) % 32)
	goto L86
L85:
	;
	v265 = v258
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v265
	v270 = base.I32_div_s(v265+int32(7), int32(8))
	v271 = F_palloc(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L31
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v271
	v274 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	if v277 <= v274 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v426 = v404
	goto L77
L89:
	;
	goto L88
L90:
	;
	v289 = base.I32_div_s(v274, int32(8))
	v290 = v271 + v289
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	goto L92
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v319))) = uint8(v317)
	goto L89
L92:
	;
	v295 = v291
	v298 = v277
	v299 = int32(1)
	v300 = v290
	goto L95
L95:
	;
	v303 = v295 | v299
	v304 = int32(1)
	v305 = v298 - v304
	v307 = v299 << (uint(v304) % 32)
	if v307 == int32(256) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v318 != int32(1) {
		goto L91
	} else {
		goto L102
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v303)
	if v305 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L98:
	;
	v317 = v303
	v318 = v307
	v319 = v300
	goto L99
L99:
	;
	if base.Ui32(int32(1)) < base.Ui32(v298) {
		v295 = v317
		v298 = v305
		v299 = v318
		v300 = v319
		goto L95
	} else {
		goto L101
	}
L100:
	;
	v313 = int32(1)
	v315 = v300 + v313
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	v317 = v316
	v318 = v313
	v319 = v315
	goto L99
L101:
	;
	goto L96
L102:
	;
	goto L89
L119:
	;
	v407 = int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v410 = v409 + v405
	if v410&(v410-v407) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v417 = v407 << (uint(int32(32)-base.I32_clz(v410)) % 32)
	goto L122
L121:
	;
	v417 = v410
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v417
	v422 = base.I32_div_s(v417+int32(7), int32(8))
	v423 = F_repalloc(m, v248, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L31
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v423
	v426 = v423
	goto L77
L124:
	;
	goto L76
L125:
	;
	goto L124
L126:
	;
	v442 = int32(1) << (uint(v429&int32(7)) % 32)
	v444 = base.I32_div_s(v429, int32(8))
	v445 = v426 + v444
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445))))
	if v430 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v546))) = uint8(v541)
	goto L125
L128:
	;
	v450 = v446
	v453 = v432
	v454 = v442
	v455 = v445
	goto L131
L129:
	;
	goto L130
L130:
	;
	v484 = base.I32_div_s(v431, int32(8))
	v485 = v430 + v484
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	v487 = int32(1)
	v488 = v446
	v491 = v432
	v492 = v442
	v493 = v445
	v494 = v485
	v495 = v486
	goto L139
L131:
	;
	v458 = v450 | v454
	v459 = int32(1)
	v460 = v453 - v459
	v462 = v454 << (uint(v459) % 32)
	if v462 == int32(256) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v473 != int32(1) {
		v541 = v472
		v546 = v474
		goto L127
	} else {
		goto L138
	}
L133:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v455))) = uint8(v458)
	if v460 == int32(0) {
		goto L125
	} else {
		goto L136
	}
L134:
	;
	v472 = v458
	v473 = v462
	v474 = v455
	goto L135
L135:
	;
	if base.Ui32(int32(1)) < base.Ui32(v453) {
		v450 = v472
		v453 = v460
		v454 = v473
		v455 = v474
		goto L131
	} else {
		goto L137
	}
L136:
	;
	v468 = int32(1)
	v470 = v455 + v468
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	v472 = v471
	v473 = v468
	v474 = v470
	goto L135
L137:
	;
	goto L132
L138:
	;
	goto L125
L139:
	;
	if v487&v495 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	if v516 == int32(1) {
		goto L125
	} else {
		goto L154
	}
L141:
	;
	v501 = v488 | v492
	goto L143
L142:
	;
	v501 = v488 & (v492 ^ int32(-1))
	goto L143
L143:
	;
	v502 = int32(1)
	v503 = v491 - v502
	v505 = v492 << (uint(v502) % 32)
	if v505 == int32(256) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v493))) = uint8(v501)
	if v503 == int32(0) {
		goto L125
	} else {
		goto L147
	}
L145:
	;
	v515 = v501
	v516 = v505
	v517 = v493
	goto L146
L146:
	;
	v519 = v487 << (uint(int32(1)) % 32)
	if v519 == int32(256) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v511 = int32(1)
	v513 = v493 + v511
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	v515 = v514
	v516 = v511
	v517 = v513
	goto L146
L148:
	;
	goto L140
L149:
	;
	if v503 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L150:
	;
	v528 = v519
	v529 = v494
	v530 = v495
	goto L151
L151:
	;
	if base.Ui32(int32(1)) < base.Ui32(v491) {
		v487 = v528
		v488 = v515
		v491 = v503
		v492 = v516
		v493 = v517
		v494 = v529
		v495 = v530
		goto L139
	} else {
		goto L153
	}
L152:
	;
	v524 = int32(1)
	v525 = v494 + v524
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v528 = v524
	v529 = v525
	v530 = v526
	goto L151
L153:
	;
	goto L148
L154:
	;
	v541 = v515
	v546 = v517
	goto L127
L155:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v569 + v570
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = v573 + v574
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+32)) = v577 + v578
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v222
	v584 = v53
	goto L21
L156:
	;
	v567 = F__emscripten_memcpy_bulkmem(m, v562+v563, v565, v566)
	mBase = m.M
	goto L158
L157:
	;
	goto L158
L158:
	;
	goto L155
L159:
	;
	F_errmsg_internal(m, int32(66939), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L31
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(518919), int32(985), int32(392184))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L31
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_agg_array_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(1)
		v17 = v12 + v16
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		v22 = v20 & v16
		if v22 != 0 {
			v23 = v17
		} else {
			v23 = v12 + int32(4)
		}
		if v20 == int32(1) {
			v26 = int32(4)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			if v28&int32(254) == int32(2) {
				v37 = v26
			} else {
				v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
			}
			if v28 == int32(1) {
				v40 = v26
			} else {
				v40 = v37
			}
			v51 = v40
		} else {
			v41 = int32(1)
			if v22 != 0 {
				v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v51
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v23
		v57 = F_pq_getmsgint(m, v9, int32(4))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v60 = F_pq_getmsgint(m, v9, int32(4))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v63 = F_pq_getmsgint(m, v9, int32(4))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, _consts[10]))
					v69 = F_initArrayResultArr(m, v60, v57, v67, int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v74 = int32(1024)
						for {
							if v74 < v63 {
								v74 = v74 << (uint(int32(1)) % 32)
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v74
						v81 = F_palloc(m, v74)
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v81
							v84 = F_pq_getmsgbytes(m, v9, v63)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
								if v63 != 0 {
									v87 = F__emscripten_memcpy_bulkmem(m, v86, v84, v63)
									mBase = m.M
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v63
								v91 = F_pq_getmsgint(m, v9, int32(4))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v91
									v95 = F_pq_getmsgint(m, v9, int32(4))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = v95
										if int32(0) < v95 {
											v103 = base.I32_div_s(v95+int32(7), int32(8))
											v104 = F_palloc(m, v103)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = v104
												v107 = F_pq_getmsgbytes(m, v9, v103)
												mBase = m.M
												v108 = m.ExcPending
												if v108 != 0 {
													return int32(0)
												} else {
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
													if v103 != 0 {
														v110 = F__emscripten_memcpy_bulkmem(m, v109, v107, v103)
														mBase = m.M
													} else {
													}
													v117 = F_pq_getmsgint(m, v9, int32(4))
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v117
														v121 = F_pq_getmsgint(m, v9, int32(4))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v121
															v125 = F_pq_getmsgbytes(m, v9, int32(24))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v127 = *(*int64)(unsafe.Add(mBase, uint32(v125)))
																*(*int64)(unsafe.Add(mBase, uint32(v69)+32)) = v127
																v129 = *(*int64)(unsafe.Add(mBase, uint32(v125)+16))
																*(*int64)(unsafe.Add(mBase, uint32(v69)+48)) = v129
																v131 = *(*int64)(unsafe.Add(mBase, uint32(v125)+8))
																*(*int64)(unsafe.Add(mBase, uint32(v69)+40)) = v131
																v134 = F_pq_getmsgbytes(m, v9, int32(24))
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	v136 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
																	*(*int64)(unsafe.Add(mBase, uint32(v69)+56)) = v136
																	v138 = *(*int64)(unsafe.Add(mBase, uint32(v134)+16))
																	*(*int64)(unsafe.Add(mBase, uint32(v69)+72)) = v138
																	v142 = *(*int64)(unsafe.Add(mBase, uint32(v134)+8))
																	*(*int64)(unsafe.Add(mBase, uint32(v69-int32(-64)))) = v142
																	F_pq_getmsgend(m, v9)
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v9 + int32(16)
																		return v69
																	}
																}
															}
														}
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = int32(0)
											v117 = F_pq_getmsgint(m, v9, int32(4))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v117
												v121 = F_pq_getmsgint(m, v9, int32(4))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v69)+28)) = v121
													v125 = F_pq_getmsgbytes(m, v9, int32(24))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v127 = *(*int64)(unsafe.Add(mBase, uint32(v125)))
														*(*int64)(unsafe.Add(mBase, uint32(v69)+32)) = v127
														v129 = *(*int64)(unsafe.Add(mBase, uint32(v125)+16))
														*(*int64)(unsafe.Add(mBase, uint32(v69)+48)) = v129
														v131 = *(*int64)(unsafe.Add(mBase, uint32(v125)+8))
														*(*int64)(unsafe.Add(mBase, uint32(v69)+40)) = v131
														v134 = F_pq_getmsgbytes(m, v9, int32(24))
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															v136 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
															*(*int64)(unsafe.Add(mBase, uint32(v69)+56)) = v136
															v138 = *(*int64)(unsafe.Add(mBase, uint32(v134)+16))
															*(*int64)(unsafe.Add(mBase, uint32(v69)+72)) = v138
															v142 = *(*int64)(unsafe.Add(mBase, uint32(v134)+8))
															*(*int64)(unsafe.Add(mBase, uint32(v69-int32(-64)))) = v142
															F_pq_getmsgend(m, v9)
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
																return int32(0)
															} else {
																m.G0 = v9 + int32(16)
																return v69
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_array_create_iterator(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v9 = F_palloc0(m, int32(48))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l1 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(220191), int32(0))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518800), int32(4612), int32(220212))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v15 < l1 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(220191), int32(0))
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518800), int32(4612), int32(220212))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v18 != 0 {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v26 = l0 + v19<<(uint(int32(3))%32) + int32(16)
				} else {
					v26 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = l0 + int32(16)
				v31 = F_ArrayGetNItems(m, v28, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31
					if l2 != 0 {
						v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)) = uint16(v34)
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v36)
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v38)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
						if l1 != 0 {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v51 = int32(2)
							v55 = l1 << (uint(v51) % 32)
							v56 = v30 + v50<<(uint(v51)%32) - v55
							*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v56
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v60 = v58 << (uint(v51) % 32)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v30 + v60 + v60 - v55
							v65 = F_ArrayGetNItems(m, l1, v56)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v65
								v70 = F_palloc(m, v65<<(uint(int32(2))%32))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v70
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v74 = F_palloc(m, v73)
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v74
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										if v81 == int32(0) {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										} else {
											v91 = v81
										}
										*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
										return v9
									}
								}
							}
						} else {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							if v81 == int32(0) {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							} else {
								v91 = v81
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
							return v9
						}
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						F_get_typlenbyvalalign(m, v40, v9+int32(12), v9+int32(14), v9+int32(15))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l1
							if l1 != 0 {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v51 = int32(2)
								v55 = l1 << (uint(v51) % 32)
								v56 = v30 + v50<<(uint(v51)%32) - v55
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v56
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v60 = v58 << (uint(v51) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v30 + v60 + v60 - v55
								v65 = F_ArrayGetNItems(m, l1, v56)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v65
									v70 = F_palloc(m, v65<<(uint(int32(2))%32))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v70
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
										v74 = F_palloc(m, v73)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v74
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											if v81 == int32(0) {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											} else {
												v91 = v81
											}
											*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
											return v9
										}
									}
								}
							} else {
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v81 == int32(0) {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v91 = v81
								}
								*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l0 + v91
								return v9
							}
						}
					}
				}
			}
		}
	}
}
func F_array_dim_to_json(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v84 int32
	_ = v84
	F_appendStringInfoChar(m, l0, int32(91))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = l3 + l1<<(uint(int32(2))%32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if int32(0) < v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l9 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v25 = int32(780778)
	goto L8
L7:
	;
	v25 = int32(700613)
	goto L8
L8:
	;
	v26 = int32(1)
	v27 = l1 + v26
	v30 = v26
	goto L9
L9:
	;
	if int32(2) <= v30 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	F_appendStringInfoString(m, l0, v25)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l2 == v27 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v66 = v30 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v66 <= v67 {
		v30 = v66
		goto L9
	} else {
		goto L21
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4+v47<<(uint(int32(2))%32))))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v47))))
	F_datum_to_json_internal(m, v51, v53, l0, l7, l8, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_array_dim_to_json(m, l0, v27, l2, l3, l4, l5, l6, l7, l8, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v57 + int32(1)
	goto L15
L20:
	;
	goto L15
L21:
	;
	goto L10
L22:
	;
	return
}
func F_array_dim_to_jsonb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v15 = F_pushJsonbValue(m, l0, int32(4), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
	v20 = l3 + l1<<(uint(int32(2))%32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if int32(0) < v21 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(1)
	v25 = l1 + v24
	v28 = v24
	goto L6
L4:
	;
	goto L5
L5:
	;
	v75 = F_pushJsonbValue(m, l0, int32(5), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	if l2 == v25 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v58 = v28 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v58 <= v59 {
		v28 = v58
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l4+v40<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v40))))
	F_datum_to_jsonb_internal(m, v44, v46, l0, l7, l8, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_array_dim_to_jsonb(m, l0, v25, l2, l3, l4, l5, l6, l7, l8)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v50 + int32(1)
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L7
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75
	return
}
func F_array_exec_setup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v11 < int32(7) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v11 != v14 {
			v17 = v14
		} else {
			v17 = int32(0)
		}
		if v17 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(336852), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					F_errfinish(m, int32(519325), int32(495), int32(244270))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v19 = F_palloc(m, int32(60))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v19
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v25 = F_get_typlen(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)) = uint16(v25)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_get_typlenbyvalalign(m, v28, v19+int32(6), v19+int32(8), v19+int32(9))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v14 != 0 {
							v39 = int32(1269)
						} else {
							v39 = int32(1270)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v39
						if v14 != 0 {
							v43 = int32(1271)
						} else {
							v43 = int32(1272)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v43
						if v14 != 0 {
							v47 = int32(1273)
						} else {
							v47 = int32(1274)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(1275)
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v61
				F_errmsg(m, int32(711130), v9)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(519325), int32(490), int32(244270))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
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
}
func F_array_fill_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int64
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v444 int32
	_ = v444
	var v447 int64
	_ = v447
	var v451 int32
	_ = v451
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 < int32(2) {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L15
	} else {
		goto L166
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L15
	} else {
		goto L161
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L15
	} else {
		goto L157
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L15
	} else {
		goto L152
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L15
	} else {
		goto L148
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L15
	} else {
		goto L144
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L15
	} else {
		goto L140
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L15
	} else {
		goto L135
	}
L11:
	;
	if v168 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	v27 = l0 + int32(16)
	v28 = F_ArrayGetNItems(m, v22, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v145 = v22
	goto L14
L14:
	;
	v168 = v145
	v175 = (v145<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v32 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = v27 + v33<<(uint(int32(3))%32)
	goto L19
L18:
	;
	v38 = int32(0)
	goto L19
L19:
	;
	if int32(8) <= v28 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v54 = v28
	v55 = v38
	goto L23
L21:
	;
	v80 = v28
	v81 = v38
	goto L22
L22:
	;
	if int32(0) < v80 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v65 != int32(255) {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v80 = v73
	v81 = v38 + int32(base.Ui32(v28-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L22
L25:
	;
	v73 = v54 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v54) {
		v54 = v73
		v55 = v55 + int32(1)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v101 = v80
	v102 = int32(1)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v32 != 0 {
		v168 = v138
		v175 = v32
		goto L11
	} else {
		goto L34
	}
L30:
	;
	if v102&v93 == int32(0) {
		goto L7
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	v115 = int32(1)
	if v115 < v101 {
		v101 = v101 - v115
		v102 = v102 << (uint(v115) % 32)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v145 = v138
	goto L14
L35:
	;
	if l1 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	v187 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v182 < int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v182) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v187 = v182
	goto L35
L41:
	;
	v346 = l0 + v175
	v347 = F_ArrayGetNItems(m, v187, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L15
	} else {
		goto L71
	}
L42:
	;
	v190 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+64)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v20)+56)) = v190
	*(*int64)(unsafe.Add(mBase, uint32(v20)+48)) = v190
	v345 = v20 + int32(48)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(2) <= v198 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v201 == int32(0) {
		v309 = v7
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v314 {
		goto L64
	} else {
		goto L65
	}
L47:
	;
	v205 = l1 + int32(16)
	v206 = F_ArrayGetNItems(m, v198, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L15
	} else {
		goto L48
	}
L48:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v208 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v214 = v205 + v209<<(uint(int32(3))%32)
	goto L51
L50:
	;
	v214 = int32(0)
	goto L51
L51:
	;
	if int32(8) <= v206 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v230 = v206
	v231 = v214
	goto L55
L53:
	;
	v256 = v206
	v257 = v214
	goto L54
L54:
	;
	if v256 <= int32(0) {
		v309 = v208
		goto L46
	} else {
		goto L59
	}
L55:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v241 != int32(255) {
		goto L3
	} else {
		goto L57
	}
L56:
	;
	v256 = v249
	v257 = v214 + int32(base.Ui32(v206-int32(8))>>(uint(int32(3))%32)) + int32(1)
	goto L54
L57:
	;
	v249 = v230 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v230) {
		v230 = v249
		v231 = v231 + int32(1)
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	v277 = v256
	v278 = int32(1)
	goto L60
L60:
	;
	if v278&v269 == int32(0) {
		goto L3
	} else {
		goto L62
	}
L61:
	;
	v309 = v208
	goto L46
L62:
	;
	v291 = int32(1)
	if v291 < v277 {
		v277 = v277 - v291
		v278 = v278 << (uint(v291) % 32)
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v318 = v317
	goto L66
L65:
	;
	v318 = v7
	goto L66
L66:
	;
	if v318 != v187 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	if v309 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v326 = v309
	goto L70
L69:
	;
	v326 = (v314<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L70
L70:
	;
	v345 = l1 + v326
	goto L41
L71:
	;
	F_ArrayCheckBounds(m, v187, v346, v345)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L15
	} else {
		goto L72
	}
L72:
	;
	if v347 <= int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	m.G0 = v20 + int32(80)
	return v548
L74:
	;
	v354 = F_palloc0(m, int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L15
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	if v362 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v354)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = int64(64)
	v548 = v354
	goto L73
L78:
	;
	if l3 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	F_get_typlenbyvalalign(m, l4, v377+int32(4), v377+int32(6), v377+int32(7))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L15
	} else {
		goto L85
	}
L80:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	v367 = F_MemoryContextAlloc(m, v365, int32(48))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L15
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	if v375 == l4 {
		v387 = v362
		goto L78
	} else {
		goto L84
	}
L83:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v369)+16)) = v367
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = int32(0)
	v377 = v372
	goto L79
L84:
	;
	v377 = v362
	goto L79
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v377))) = l4
	v387 = v377
	goto L78
L86:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+6)))
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+7)))
	v392 = int32(*(*int16)(unsafe.Add(mBase, uint32(v387)+4)))
	if v392 != int32(-1) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	goto L88
L88:
	;
	v522 = base.I32_div_s(v347+int32(7), int32(8))
	v529 = (v522 + v187<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v530 = F_palloc0(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L15
	} else {
		goto L126
	}
L89:
	;
	switch v391 - int32(99) {
	case 0:
		v444 = v429
		goto L104
	case 1:
		goto L106
	default:
		goto L105
	case 6:
		goto L107
	}
L90:
	;
	v424 = F_strlen(m, l2)
	mBase = m.M
	v427 = l2
	v429 = v424 + int32(1)
	goto L89
L91:
	;
	if v392 <= int32(0) {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v397 = F_pg_detoast_datum(m, l2)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L15
	} else {
		goto L95
	}
L94:
	;
	v427 = l2
	v429 = v392
	goto L89
L95:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	if v399 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+1)))
	if base.Ui32((v403-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v427 = v397
		v429 = int32(6)
		goto L89
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v417 = int32(1)
	if v399&v417 != 0 {
		v427 = v397
		v429 = int32(base.Ui32(v399) >> (uint(v417) % 32))
		goto L89
	} else {
		goto L103
	}
L99:
	;
	v410 = int32(18)
	if v403&int32(255) == v410 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v416 = v410
	goto L102
L101:
	;
	v416 = int32(2)
	goto L102
L102:
	;
	v427 = v397
	v429 = v416
	goto L89
L103:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v427 = v397
	v429 = int32(base.Ui32(v421) >> (uint(int32(2)) % 32))
	goto L89
L104:
	;
	v447 = base.I64_extend_i32_s(v444) * base.I64_extend_i32_s(v347)
	v451 = base.I32_wrap_i64(v447)
	if base.I32_wrap_i64(int64(base.Ui64(v447)>>(uint(int64(32))%64))) != v451>>(uint(int32(31))%32) {
		goto L1
	} else {
		goto L108
	}
L105:
	;
	v444 = (v429 + int32(1)) & int32(-2)
	goto L104
L106:
	;
	v444 = (v429 + int32(7)) & int32(-8)
	goto L104
L107:
	;
	v444 = (v429 + int32(3)) & int32(-4)
	goto L104
L108:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v451) {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v463 = v451 + (v187<<(uint(int32(3))%32)+int32(23))&int32(120)
	v464 = F_palloc0(m, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L15
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+12)) = l4
	v467 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v464)+8)) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v464)+4)) = v187
	v471 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v463 << (uint(v471) % 32)
	v475 = v464 + int32(16)
	v477 = v187 << (uint(v471) % 32)
	if v477 != 0 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v477 != 0 {
		goto L116
	} else {
		goto L117
	}
L112:
	;
	v478 = F__emscripten_memcpy_bulkmem(m, v475, v346, v477)
	mBase = m.M
	v479 = v478
	goto L114
L113:
	;
	v479 = v475
	goto L114
L114:
	;
	goto L111
L115:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v464)+8))
	if v484 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v481 = F__emscripten_memcpy_bulkmem(m, v479+v477, v345, v477)
	mBase = m.M
	goto L118
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	v492 = v484
	goto L121
L120:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v492 = (v485<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L121
L121:
	;
	v502 = v492 + v464
	v503 = v467
	goto L122
L122:
	;
	v513 = F_ArrayCastAndSet(m, v427, v392, v390&int32(1), base.I32_extend8_s(v391), v502)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L15
	} else {
		goto L124
	}
L123:
	;
	v548 = v464
	goto L73
L124:
	;
	v517 = v503 + int32(1)
	if v517 != v347 {
		v502 = v513 + v502
		v503 = v517
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v530)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v530)+8)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v530)+4)) = v187
	v535 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v529 << (uint(v535) % 32)
	v539 = v530 + int32(16)
	v541 = v187 << (uint(v535) % 32)
	if v541 != 0 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v541 != 0 {
		goto L132
	} else {
		goto L133
	}
L128:
	;
	v542 = F__emscripten_memcpy_bulkmem(m, v539, v346, v541)
	mBase = m.M
	v543 = v542
	goto L130
L129:
	;
	v543 = v539
	goto L130
L130:
	;
	goto L127
L131:
	;
	v548 = v530
	goto L73
L132:
	;
	v545 = F__emscripten_memcpy_bulkmem(m, v543+v541, v345, v541)
	mBase = m.M
	goto L134
L133:
	;
	goto L134
L134:
	;
	goto L131
L135:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L15
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(126195), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L15
	} else {
		goto L137
	}
L137:
	;
	F_errdetail(m, int32(651424), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L15
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(518800), int32(6113), int32(326784))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L15
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L15
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(317044), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L15
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(518800), int32(6118), int32(326784))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L15
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L15
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v182
	F_errmsg(m, int32(504831), v20)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L15
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(518800), int32(6126), int32(326784))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L15
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
	F_errcode(m, int32(261))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L15
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v182
	F_errmsg(m, int32(711130), v20+int32(16))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L15
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(518800), int32(6131), int32(326784))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L15
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L15
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(126195), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	F_errdetail(m, int32(651424), int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L15
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(518800), int32(6139), int32(326784))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L15
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L15
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(317044), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L15
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(518800), int32(6144), int32(326784))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L15
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L15
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(126195), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L15
	} else {
		goto L163
	}
L163:
	;
	F_errdetail(m, int32(603855), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L15
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(518800), int32(6150), int32(326784))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L15
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L15
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(1073741823)
	F_errmsg(m, int32(711032), v20+int32(32))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L15
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(518800), int32(6223), int32(326784))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L15
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_free_iterator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v2 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			F_pfree(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_array_get_element(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v461 int32
	_ = v461
	v9 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	if v9 < l3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v461
L2:
	;
	v461 = int32(0)
	goto L1
L3:
	;
	if l1 != v236 {
		goto L47
	} else {
		goto L48
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	v25 = base.I32_div_s(l3, l4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v25
	v234 = l0
	v236 = int32(1)
	v238 = v18 + int32(12)
	v240 = v9
	v241 = v18 + int32(8)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v31 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v214 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L33
	} else {
		goto L38
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v34&int32(254) != int32(2) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if v40 != l1 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v99 = int32(0)
	v108 = l1 - int32(1)
	if v108 < v99 {
		v186 = v99
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v61 = v48
	goto L16
L12:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v51)
	goto L2
L13:
	;
	if base.Ui32(v40-int32(7)) < base.Ui32(int32(-6)) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v48 = int32(0)
	if l1 <= v48 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v69 = v61 << (uint(int32(2)) % 32)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2+v69)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v46)))
	if v71 < v73 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v82)
	goto L2
L18:
	;
	goto L17
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69+v47)))
	if v76+v73 <= v71 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v80 = v61 + int32(1)
	if l1 != v80 {
		v61 = v80
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L10
L22:
	;
	F_deconstruct_expanded_array(m, v39)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L33
	} else {
		goto L34
	}
L23:
	;
	goto L22
L24:
	;
	v111 = int32(1)
	if v108 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if l1&v111 == int32(0) {
		v186 = v163
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v163 = v99
	v164 = v108
	v165 = v111
	goto L25
L27:
	;
	goto L28
L28:
	;
	v122 = v99
	v123 = v108
	v124 = v111
	v125 = v99
	goto L29
L29:
	;
	v130 = int32(2)
	v131 = v123 << (uint(v130) % 32)
	v133 = v131 - int32(4)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2+v133)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v46+v133)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v131+v47)))
	v141 = v140 * v124
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v131+l2)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v131+v46)))
	v150 = (v135-v137)*v141 + ((v144-v146)*v124 + v122)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v47+v133)))
	v153 = v152 * v141
	v155 = v123 - v130
	v157 = v125 + v130
	if v157 != l1&int32(-2) {
		v122 = v150
		v123 = v155
		v124 = v153
		v125 = v157
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v163 = v150
	v164 = v155
	v165 = v153
	goto L25
L31:
	;
	goto L30
L32:
	;
	v174 = v164 << (uint(int32(2)) % 32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l2+v174)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v46)))
	v186 = (v176-v178)*v165 + v163
	goto L23
L33:
	;
	return int32(0)
L34:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	if v199 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v208)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v198+v186<<(uint(int32(2))%32))))
	v461 = v213
	goto L1
L36:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v199))))
	if v203 != int32(1) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v206)
	goto L2
L38:
	;
	v217 = v214 + int32(16)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v220 = v218 << (uint(int32(3)) % 32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	if v223 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v224 = v217 + v220
	goto L41
L40:
	;
	v224 = int32(0)
	goto L41
L41:
	;
	if v223 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v229 = v223
	goto L44
L43:
	;
	v229 = (v220 + int32(23)) & int32(-8)
	goto L44
L44:
	;
	v234 = v214 + v229
	v236 = v218
	v238 = v217
	v240 = v224
	v241 = v217 + v218<<(uint(int32(2))%32)
	goto L3
L45:
	;
	v298 = int32(0)
	v307 = l1 - int32(1)
	if v307 < v298 {
		v385 = v298
		goto L58
	} else {
		goto L59
	}
L46:
	;
	v260 = v247
	goto L51
L47:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v250)
	goto L2
L48:
	;
	if base.Ui32(v236-int32(7)) < base.Ui32(int32(-6)) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v247 = int32(0)
	if l1 <= v247 {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v268 = v260 << (uint(int32(2)) % 32)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l2+v268)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v268+v241)))
	if v270 < v272 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v281)
	goto L2
L53:
	;
	goto L52
L54:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v268+v238)))
	if v275+v272 <= v270 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v279 = v260 + int32(1)
	if l1 != v279 {
		v260 = v279
		goto L51
	} else {
		goto L56
	}
L56:
	;
	goto L45
L57:
	;
	if v240 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	goto L57
L59:
	;
	v310 = int32(1)
	if v307 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if l1&v310 == int32(0) {
		v385 = v362
		goto L58
	} else {
		goto L67
	}
L61:
	;
	v362 = v298
	v363 = v307
	v364 = v310
	goto L60
L62:
	;
	goto L63
L63:
	;
	v321 = v298
	v322 = v307
	v323 = v310
	v324 = v298
	goto L64
L64:
	;
	v329 = int32(2)
	v330 = v322 << (uint(v329) % 32)
	v332 = v330 - int32(4)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l2+v332)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v241+v332)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v330+v238)))
	v340 = v339 * v323
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v330+l2)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v330+v241)))
	v349 = (v334-v336)*v340 + ((v343-v345)*v323 + v321)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v238+v332)))
	v352 = v351 * v340
	v354 = v322 - v329
	v356 = v324 + v329
	if v356 != l1&int32(-2) {
		v321 = v349
		v322 = v354
		v323 = v352
		v324 = v356
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v362 = v349
	v363 = v354
	v364 = v352
	goto L60
L66:
	;
	goto L65
L67:
	;
	v373 = v363 << (uint(int32(2)) % 32)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l2+v373)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v241)))
	v385 = (v375-v377)*v364 + v362
	goto L58
L68:
	;
	v406 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v406)
	v409 = F_array_seek(m, v234, v406, v240, v385, l4, l6)
	mBase = m.M
	if l5 == v406 {
		v461 = v409
		goto L1
	} else {
		goto L71
	}
L69:
	;
	v396 = base.I32_div_s(v385, int32(8))
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v396))))
	if int32(base.Ui32(v398)>>(uint(v385&int32(7))%32))&int32(1) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v404 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v404)
	goto L2
L71:
	;
	switch l4 - int32(1) {
	case 0:
		goto L75
	case 1:
		goto L74
	default:
		goto L72
	case 3:
		goto L73
	}
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L33
	} else {
		goto L76
	}
L73:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v461 = v416
	goto L1
L74:
	;
	v415 = int32(*(*int16)(unsafe.Add(mBase, uint32(v409))))
	v461 = v415
	goto L1
L75:
	;
	v414 = int32(*(*int8)(unsafe.Add(mBase, uint32(v409))))
	v461 = v414
	goto L1
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l4
	F_errmsg_internal(m, int32(506326), v18)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L33
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(343236), int32(70), int32(73749))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L33
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_le(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_array_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v5 = int32(-1)
	v9 = F_array_get_element(m, l0, int32(1), l1, v5, v5, int32(0), int32(105), l2)
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_array_subscript_fetch_old_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != 0 {
		v37 = int32(1)
		v48 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v37)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v48
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+4)))
		v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+6)))
		v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+9)))
		v34 = F_array_get_slice(m, v21, v22, v23+int32(12), v23+int32(36), v28, v29, v30, v31, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			v37 = int32(0)
			v48 = v34
			*(*uint8)(unsafe.Add(mBase, uint32(v14)+52)) = uint8(v37)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v48
			return
		}
	}
}
func F_array_to_halfvec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v130 float32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L69
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L65
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L61
	}
L4:
	;
	return int32(0)
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v19 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L57
	}
L9:
	;
	v24 = F_array_contains_nulls(m, v15)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_get_typlenbyvalalign(m, v26, v12+int32(30), v12+int32(29), v12+int32(28))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	if v24 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+30)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+29)))
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+28)))
	F_deconstruct_array(m, v15, v36, v37, v38, v12+int32(24), int32(0), v12+int32(20))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	F_CheckDim_1(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if base.B2i32(v22 != int32(-1))&base.B2i32(v51 != v22) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v56 = F_mul_size(m, int32(2), v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v58 = F_add_size(m, int32(8), v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v60 = F_palloc0(m, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v58 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	switch v66 - int32(700) {
	case 0:
		goto L24
	case 1:
		goto L23
	default:
		goto L25
	}
L21:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	F_pfree(m, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L49
	}
L22:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v170 <= int32(0) {
		goto L21
	} else {
		goto L44
	}
L23:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v138 <= int32(0) {
		goto L21
	} else {
		goto L39
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v108 <= int32(0) {
		goto L21
	} else {
		goto L34
	}
L25:
	;
	if v66 == int32(23) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	if v66 != int32(1700) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v73 <= int32(0) {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	v80 = int32(0)
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v80<<(uint(int32(2))%32))))
	v98 = F_DirectFunctionCall1Coll(m, int32(1338), int32(0), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L21
L31:
	;
	v101 = F_Float4ToHalf(m, base.F32_reinterpret_i32(v98))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60+int32(8)+v80<<(uint(int32(1))%32)))) = uint16(v101)
	v105 = v80 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v105 < v106 {
		v80 = v105
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v115 = int32(0)
	goto L35
L35:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v130 = *(*float32)(unsafe.Add(mBase, uint32(v126+v115<<(uint(int32(2))%32))))
	v131 = F_Float4ToHalf(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L37
	}
L36:
	;
	goto L21
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60+int32(8)+v115<<(uint(int32(1))%32)))) = uint16(v131)
	v135 = v115 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v135 < v136 {
		v115 = v135
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v145 = int32(0)
	goto L40
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v145<<(uint(int32(2))%32))))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v160)))
	v163 = F_Float4ToHalf(m, base.F32_demote_f64(v161))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L21
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60+int32(8)+v145<<(uint(int32(1))%32)))) = uint16(v163)
	v167 = v145 + int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v167 < v168 {
		v145 = v167
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v177 = int32(0)
	goto L45
L45:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+v177<<(uint(int32(2))%32))))
	v194 = F_Float4ToHalf(m, base.F32_convert_i32_s(v192))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L21
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60+int32(8)+v177<<(uint(int32(1))%32)))) = uint16(v194)
	v198 = v177 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v198 < v199 {
		v177 = v198
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v213 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+4)))
	if int32(0) < v213 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v220 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	m.G0 = v12 + int32(32)
	return v60
L53:
	;
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60+int32(8)+v220<<(uint(int32(1))%32)))))
	F_CheckElement_1(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	v235 = v220 + int32(1)
	v236 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+4)))
	if v235 < v236 {
		v220 = v235
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errmsg(m, int32(570932), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(525467), int32(456), int32(514707))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(162284), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(525467), int32(461), int32(514707))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
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
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v22
	F_errmsg(m, int32(489736), v12)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(525467), int32(92), int32(302808))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
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
	F_errcode(m, int32(130))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(385435), int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(525467), int32(495), int32(514707))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_to_json_pretty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_makeStringInfo(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		F_array_to_json_internal(m, v4, v5, base.B2i32(v3 != int32(0)))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			v15 = F_cstring_to_text_with_len(m, v13, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_array_to_text_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v22 = l1 + int32(16)
	v23 = F_ArrayGetNItems(m, v20, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 - int32(-64)
	return v283
L2:
	;
	return int32(0)
L3:
	;
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = F_palloc(m, int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_initStringInfo(m, v16+int32(-16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(16)
	v283 = v30
	goto L1
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if int32(0) < v23 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	F_get_type_io_data(m, v34, int32(1), v56+int32(4), v56+int32(6), v56+int32(7), v56+int32(8), v56+int32(12), v56+int32(16))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L16
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v45 = F_MemoryContextAlloc(m, v43, int32(48))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v54 == v34 {
		v80 = v40
		goto L9
	} else {
		goto L15
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v34 ^ int32(-1)
	v56 = v50
	goto L10
L15:
	;
	v56 = v40
	goto L10
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	F_fmgr_info_cxt(m, v72, v56+int32(20), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v34
	v80 = v56
	goto L9
L18:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	v84 = base.I32_extend16_s(v83)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v87 = v85 << (uint(int32(3)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v90 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v270 = v268 + int32(4)
	v271 = F_palloc(m, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
	} else {
		goto L89
	}
L21:
	;
	v91 = v22 + v87
	goto L23
L22:
	;
	v91 = int32(0)
	goto L23
L23:
	;
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v98 = v90
	goto L26
L25:
	;
	v98 = (v87 + int32(23)) & int32(-8)
	goto L26
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)))
	v101 = int32(1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+7)))
	v108 = int32(0)
	v111 = v91
	v112 = l1 + v98
	v115 = v101
	v118 = v108
	v119 = v108
	goto L27
L27:
	;
	if v111 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L20
L29:
	;
	v239 = int32(1)
	v241 = v115 << (uint(v239) % 32)
	v243 = base.B2i32(v241 == int32(256))
	if v241 == int32(256) {
		goto L79
	} else {
		goto L80
	}
L30:
	;
	v236 = v233
	v238 = int32(1)
	goto L29
L31:
	;
	if v100&v101 != 0 {
		goto L41
	} else {
		goto L42
	}
L32:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	if v115&v128 != 0 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if l3 == int32(0) {
		v236 = v112
		v238 = v119
		goto L29
	} else {
		goto L34
	}
L34:
	;
	if v119&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l2
	F_appendStringInfo(m, v16+int32(-16), int32(186498), v16+int32(-32))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_appendStringInfoString(m, v16+int32(-16), l3)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L39
	}
L38:
	;
	v233 = v112
	goto L30
L39:
	;
	v233 = v112
	goto L30
L40:
	;
	v164 = F_OutputFunctionCall(m, v80+int32(20), v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L51
	}
L41:
	;
	switch v83 - v101 {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L44
	case 3:
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v163 = v112
	goto L40
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v163 = v149
	goto L40
L46:
	;
	v148 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112))))
	v163 = v148
	goto L40
L47:
	;
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v112))))
	v163 = v147
	goto L40
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v84
	F_errmsg_internal(m, int32(506326), v18)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(343236), int32(70), int32(73749))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
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
	if v119&int32(1) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if int32(0) < v84 {
		v220 = v112 + v84
		goto L58
	} else {
		goto L59
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = l2
	F_appendStringInfo(m, v16+int32(-16), int32(186498), v16+int32(-48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_appendStringInfoString(m, v16+int32(-16), v164)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L57
	}
L56:
	;
	goto L52
L57:
	;
	goto L52
L58:
	;
	switch v105 - int32(99) {
	case 0:
		v233 = v220
		goto L30
	case 1:
		goto L77
	default:
		goto L76
	case 6:
		goto L78
	}
L59:
	;
	if v84 == int32(-1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v186 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v215 = F_strlen(m, v112)
	mBase = m.M
	v220 = v215 + v112 + int32(1)
	goto L58
L63:
	;
	v189 = int32(6)
	v191 = int32(18)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	if v193 == v191 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v206 = int32(1)
	if v186&v206 != 0 {
		v220 = v112 + int32(base.Ui32(v186)>>(uint(v206)%32))
		goto L58
	} else {
		goto L75
	}
L66:
	;
	v196 = v191
	goto L68
L67:
	;
	v196 = int32(2)
	goto L68
L68:
	;
	if v193&int32(254) == int32(2) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v201 = v189
	goto L71
L70:
	;
	v201 = v196
	goto L71
L71:
	;
	if v193 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v204 = v189
	goto L74
L73:
	;
	v204 = v201
	goto L74
L74:
	;
	v220 = v112 + v204
	goto L58
L75:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v220 = v112 + int32(base.Ui32(v211)>>(uint(int32(2))%32))
	goto L58
L76:
	;
	v233 = (v220 + int32(1)) & int32(-2)
	goto L30
L77:
	;
	v233 = (v220 + int32(7)) & int32(-8)
	goto L30
L78:
	;
	v233 = (v220 + int32(3)) & int32(-4)
	goto L30
L79:
	;
	v244 = v239
	goto L81
L80:
	;
	v244 = v241
	goto L81
L81:
	;
	if v111 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v245 = v244
	goto L84
L83:
	;
	v245 = v115
	goto L84
L84:
	;
	if v111 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v248 = v111 + v243
	goto L87
L86:
	;
	v248 = int32(0)
	goto L87
L87:
	;
	v250 = v118 + int32(1)
	if v250 != v23 {
		v111 = v248
		v112 = v236
		v115 = v245
		v118 = v250
		v119 = v238
		goto L27
	} else {
		goto L88
	}
L88:
	;
	goto L28
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v270 << (uint(int32(2)) % 32)
	if v268 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_pfree(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L94
	}
L91:
	;
	v278 = F__emscripten_memcpy_bulkmem(m, v271+int32(4), v267, v268)
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	v283 = v271
	goto L1
}
func F_array_to_text_null(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v2 = int32(0)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == v2 {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v11 != int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v24 = F_pg_detoast_datum_packed(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = F_pg_detoast_datum_packed(m, v24)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
						if v28 == int32(1) {
							v31 = int32(4)
							v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
							if v33&int32(254) == int32(2) {
								v42 = v31
							} else {
								v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
							}
							if v33 == int32(1) {
								v45 = v31
							} else {
								v45 = v42
							}
							v58 = v45
						} else {
							v46 = int32(1)
							if v28&v46 != 0 {
								v58 = int32(base.Ui32(v28)>>(uint(v46)%32)) - v46
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
								v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v61 = F_palloc(m, v58+int32(1))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v63 = int32(1)
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
							if v65&v63 != 0 {
								v68 = v63
							} else {
								v68 = int32(4)
							}
							if v58 != 0 {
								v70 = F__emscripten_memcpy_bulkmem(m, v61, v26+v68, v58)
								mBase = m.M
								v71 = v70
							} else {
								v71 = v61
							}
							v73 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v71+v58))) = uint8(v73)
							if v26 != v24 {
								F_pfree(m, v26)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
									if v78 != 0 {
										v137 = v2
										v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
										mBase = m.M
										v139 = m.ExcPending
										if v139 != 0 {
											return int32(0)
										} else {
											return v138
										}
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v80 = F_pg_detoast_datum_packed(m, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											v82 = F_pg_detoast_datum_packed(m, v80)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
												if v84 == int32(1) {
													v87 = int32(4)
													v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
													if v89&int32(254) == int32(2) {
														v98 = v87
													} else {
														v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
													}
													if v89 == int32(1) {
														v101 = v87
													} else {
														v101 = v98
													}
													v114 = v101
												} else {
													v102 = int32(1)
													if v84&v102 != 0 {
														v114 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
													} else {
														v108 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
														v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
													}
												}
												v117 = F_palloc(m, v114+int32(1))
												mBase = m.M
												v118 = m.ExcPending
												if v118 != 0 {
													return int32(0)
												} else {
													v119 = int32(1)
													v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
													if v121&v119 != 0 {
														v124 = v119
													} else {
														v124 = int32(4)
													}
													if v114 != 0 {
														v126 = F__emscripten_memcpy_bulkmem(m, v117, v82+v124, v114)
														mBase = m.M
														v127 = v126
													} else {
														v127 = v117
													}
													v129 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v127+v114))) = uint8(v129)
													if v82 == v80 {
														v137 = v117
														v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															return v138
														}
													} else {
														F_pfree(m, v82)
														mBase = m.M
														v133 = m.ExcPending
														if v133 != 0 {
															return int32(0)
														} else {
															v137 = v117
															v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int32(0)
															} else {
																return v138
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
								if v78 != 0 {
									v137 = v2
									v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										return v138
									}
								} else {
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v80 = F_pg_detoast_datum_packed(m, v79)
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return int32(0)
									} else {
										v82 = F_pg_detoast_datum_packed(m, v80)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
											if v84 == int32(1) {
												v87 = int32(4)
												v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
												if v89&int32(254) == int32(2) {
													v98 = v87
												} else {
													v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
												}
												if v89 == int32(1) {
													v101 = v87
												} else {
													v101 = v98
												}
												v114 = v101
											} else {
												v102 = int32(1)
												if v84&v102 != 0 {
													v114 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
												} else {
													v108 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
													v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
												}
											}
											v117 = F_palloc(m, v114+int32(1))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return int32(0)
											} else {
												v119 = int32(1)
												v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
												if v121&v119 != 0 {
													v124 = v119
												} else {
													v124 = int32(4)
												}
												if v114 != 0 {
													v126 = F__emscripten_memcpy_bulkmem(m, v117, v82+v124, v114)
													mBase = m.M
													v127 = v126
												} else {
													v127 = v117
												}
												v129 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v127+v114))) = uint8(v129)
												if v82 == v80 {
													v137 = v117
													v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return int32(0)
													} else {
														return v138
													}
												} else {
													F_pfree(m, v82)
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int32(0)
													} else {
														v137 = v117
														v138 = F_array_to_text_internal(m, l0, v19, v71, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															return v138
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
			return int32(0)
		}
	} else {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		return int32(0)
	}
}
func F_compute_array_stats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v114 int64
	_ = v114
	var v121 int32
	_ = v121
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v176 int64
	_ = v176
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int64
	_ = v224
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v351 int64
	_ = v351
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v390 int64
	_ = v390
	var v397 int64
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v455 int64
	_ = v455
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v492 int32
	_ = v492
	var v505 int64
	_ = v505
	var v510 float64
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v607 int64
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v644 int64
	_ = v644
	var v646 int64
	_ = v646
	var v650 int64
	_ = v650
	var v656 int64
	_ = v656
	var v658 int64
	_ = v658
	var v662 int64
	_ = v662
	var v663 int32
	_ = v663
	var v665 int64
	_ = v665
	var v666 int64
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v678 int32
	_ = v678
	var v699 int64
	_ = v699
	var v705 float64
	_ = v705
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v735 int64
	_ = v735
	var v736 int32
	_ = v736
	var v737 int64
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 float64
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v874 int32
	_ = v874
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1081 int32
	_ = v1081
	var v1083 int64
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int64
	_ = v1085
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1117 int64
	_ = v1117
	var v1129 int32
	_ = v1129
	var v1134 int32
	_ = v1134
	var v1156 int64
	_ = v1156
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int64
	_ = v1169
	var v1171 int64
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1200 int64
	_ = v1200
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	v5 = int32(0)
	v27 = int64(0)
	v34 = m.G0
	v36 = v34 - int32(176)
	m.G0 = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	m.T0[v41].(func(*base.Module, int32, int32, int32, float64))(m, l0, l1, l2, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v38
	*(*int32)(unsafe.Add(mBase, _consts[1027])) = v38
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+136)) = int32(1264)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+132)) = int32(1265)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+124)) = int64(68719476740)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+148)) = v55
	v59 = v47 * int32(10)
	v63 = F_hash_create(m, int32(411689), v59, v36+int32(108), int32(1224))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+76)) = int64(34359738372)
	v68 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+100)) = v68
	v73 = base.I32_div_s(v47*int32(10000), int32(7))
	v79 = F_hash_create(m, int32(411349), int32(64), v36+int32(60), int32(1064))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l2 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	if v511 == int32(0) {
		v524 = v5
		goto L69
	} else {
		goto L70
	}
L6:
	;
	v492 = v5
	v505 = v27
	v510 = float64(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v94 = int32(1)
	v97 = v5
	v101 = v5
	v105 = v5
	v114 = v27
	goto L10
L9:
	;
	v492 = v442
	v505 = v455
	v510 = base.F64_convert_i32_s(v446)
	goto L5
L10:
	;
	F_vacuum_delay_point(m, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L65
	}
L12:
	;
	goto L11
L13:
	;
	v124 = m.T0[l1].(func(*base.Module, int32, int32, int32) int32)(m, l0, v97, v36+int32(59))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+59)))
	if v126 != 0 {
		v435 = v94
		v442 = v101
		v446 = v105
		v455 = v114
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v461 = v97 + int32(1)
	if l2 != v461 {
		v94 = v435
		v97 = v461
		v101 = v442
		v105 = v446
		v114 = v455
		goto L10
	} else {
		goto L64
	}
L16:
	;
	v127 = F_toast_raw_datum_size(m, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if base.Ui32(int32(65536)) < base.Ui32(v127) {
		v435 = v94
		v442 = v101
		v446 = v105
		v455 = v114
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v131 = F_pg_detoast_datum(m, v124)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+14)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+16)))
	F_deconstruct_array(m, v131, v134, v135, v136, v36+int32(48), v36+int32(44), v36+int32(52))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v145 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	if v145 < v147 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v156 = v145
	v158 = v94
	v164 = v145
	v176 = v114
	goto L24
L22:
	;
	v372 = v94
	v378 = v145
	v390 = v114
	goto L23
L23:
	;
	v397 = v390 - v114
	*(*uint32)(unsafe.Add(mBase, uint32(v36)+156)) = uint32(v397)
	v399 = int32(1)
	v405 = F_hash_search(m, v79, v36+int32(156), v399, v36+int32(40))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L54
	}
L24:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v36)+44))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v156))))
	if v185 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v372 = v333
	v378 = v339 & int32(1)
	v390 = v351
	goto L23
L26:
	;
	v359 = v156 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	if v359 < v360 {
		v156 = v359
		v158 = v333
		v164 = v339
		v176 = v351
		goto L24
	} else {
		goto L53
	}
L27:
	;
	v333 = v158
	v339 = int32(1)
	v351 = v176
	goto L26
L28:
	;
	goto L29
L29:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v156<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+40)) = v191
	v198 = F_hash_search(m, v63, v36+int32(40), int32(1), v36+int32(39))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+39)))
	if v200 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v224 = v176 + int64(1)
	v225 = base.I64_rem_s(v224, base.I64_extend_i32_s(v73))
	if v225 != int64(0) {
		v333 = v158
		v339 = v164
		v351 = v224
		goto L26
	} else {
		goto L37
	}
L32:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if v203 == v97 {
		v333 = v158
		v339 = v164
		v351 = v176
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	v212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+14)))
	v213 = F_datumCopy(m, v210, v211, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v97
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v198)+4)) = v206 + int32(1)
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v97
	v216 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v198)+8)) = v158 - v216
	*(*int32)(unsafe.Add(mBase, uint32(v198)+4)) = v216
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v213
	goto L31
L37:
	;
	F_hash_seq_init(m, v36+int32(156), v63)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v234 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v234 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v240 = v234
	goto L43
L41:
	;
	goto L42
L42:
	;
	v333 = v158 + int32(1)
	v339 = v164
	v351 = v224
	goto L26
L43:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v240)+4))
	if v158 < v269+v270 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v288 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L51
	}
L46:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v276 = F_hash_search(m, v63, v240, int32(2), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v276 == int32(0) {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[1027]))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+12)))
	if v282 != 0 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	F_pfree(m, v273)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L45
L51:
	;
	if v288 != 0 {
		v240 = v288
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	goto L25
L54:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+40)))
	if v407 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v413 = v410 + int32(1)
	goto L57
L56:
	;
	v413 = v399
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v405)+4)) = v413
	if v124 != v131 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_pfree(m, v131)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	F_pfree(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v36)+44))
	F_pfree(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v435 = v372
	v442 = v101 + int32(1)
	v446 = v378 + v105
	v455 = v390
	goto L15
L64:
	;
	goto L9
L65:
	;
	F_errmsg_internal(m, int32(466319), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(523300), int32(695), int32(410809))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L155
	}
L69:
	;
	if v492 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L70:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	if v514 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v524 = int32(1)
	goto L69
L72:
	;
	goto L73
L73:
	;
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	if v518 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v524 = int32(2)
	goto L69
L75:
	;
	goto L76
L76:
	;
	v522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+58)))
	if v522 != 0 {
		goto L68
	} else {
		goto L77
	}
L77:
	;
	v524 = int32(3)
	goto L69
L78:
	;
	m.G0 = v36 + int32(176)
	return
L79:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v528)+412))
	if v530 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v598 = F_palloc(m, v595<<(uint(int32(2))%32))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L84
	}
L81:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v528)+376))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528)+364))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)+352))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528)+340))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v528)+328))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v528)+316))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v528)+304))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v528)+292))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v528)+280))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v528)+268))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v528)+256))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v528)+244))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v528)+232))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v528)+220))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v528)+208))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v528)+196))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v528)+184))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v528)+172))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v528)+160))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v528)+148))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v528)+136))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v528)+124))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v528)+112))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v528)+100))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v528)+88))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v528)+76))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v528-int32(-64))))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v528)+52))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v528)+40))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v528)+28))
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v528)+16))
	v595 = v531 + (v532 + (v533 + (v534 + (v535 + (v536 + (v537 + (v538 + (v539 + (v540 + (v541 + (v542 + (v543 + (v544 + (v545 + (v546 + (v547 + (v548 + (v549 + (v550 + (v551 + (v552 + (v553 + (v554 + (v555 + (v556 + (v559 + (v560 + (v561 + (v562 + (v563 + v529))))))))))))))))))))))))))))))
	goto L83
L82:
	;
	v595 = v529
	goto L83
L83:
	;
	goto L80
L84:
	;
	F_hash_seq_init(m, v36+int32(156), v63)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v607 = base.I64_div_s(v505*int64(9), base.I64_extend_i32_s(v73))
	v610 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v708 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L104
	}
L87:
	;
	if v610 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v678 = int32(0)
	v699 = v505
	v705 = float64(0)
	goto L86
L89:
	;
	goto L90
L90:
	;
	v621 = v610
	v623 = int32(0)
	v644 = v505
	v646 = v27
	goto L91
L91:
	;
	v650 = int64(*(*int32)(unsafe.Add(mBase, uint32(v621)+4)))
	if v607 < v650 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v678 = v663
	v699 = v665
	v705 = base.F64_convert_i64_u(v666)
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v598+v623<<(uint(int32(2))%32)))) = v621
	v656 = int64(*(*int32)(unsafe.Add(mBase, uint32(v621)+4)))
	if v656 < v646 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v663 = v623
	v665 = v644
	v666 = v646
	goto L95
L95:
	;
	v669 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L102
	}
L96:
	;
	v658 = v646
	goto L98
L97:
	;
	v658 = v656
	goto L98
L98:
	;
	if v644 < v656 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v662 = v644
	goto L101
L100:
	;
	v662 = v656
	goto L101
L101:
	;
	v663 = v623 + int32(1)
	v665 = v662
	v666 = v658
	goto L95
L102:
	;
	if v669 != 0 {
		v621 = v669
		v623 = v663
		v644 = v665
		v646 = v666
		goto L91
	} else {
		goto L103
	}
L103:
	;
	goto L92
L104:
	;
	if v708 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v595
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v59
	F_errmsg_internal(m, int32(503055), v36)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v678 <= v59 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	F_errfinish(m, int32(523300), int32(494), int32(133798))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v739 = l0 + int32(52)
	if int32(0) < v736 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v736 = v678
	v737 = v699
	goto L110
L112:
	;
	goto L113
L113:
	;
	F_qsort_interruptible(m, v598, v678, int32(4), int32(1266), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v598+v59<<(uint(int32(2))%32)-int32(4))))
	v735 = int64(*(*int32)(unsafe.Add(mBase, uint32(v734)+4)))
	v736 = v59
	v737 = v735
	goto L110
L115:
	;
	v742 = int32(0)
	F_qsort_interruptible(m, v598, v736, int32(4), int32(1267), v742)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	v874 = v524
	goto L117
L117:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v896)+412))
	if v898 != 0 {
		goto L126
	} else {
		goto L127
	}
L118:
	;
	v748 = int32(4562080)
	v749 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v751
	v753 = base.F64_convert_i32_u(v492)
	v756 = F_palloc(m, v736<<(uint(int32(2))%32))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v759 = v736 + int32(3)
	v762 = F_palloc(m, v759<<(uint(int32(2))%32))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v768 = v742
	goto L121
L121:
	;
	v798 = v768 << (uint(int32(2)) % 32)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798+v598)))
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v801)))
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	v804 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+14)))
	v805 = F_datumCopy(m, v802, v803, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	v817 = int32(2)
	v819 = v762 + v736<<(uint(v817)%32)
	*(*float32)(unsafe.Add(mBase, uint32(v819)+8)) = base.F32_demote_f64(base.F64_div(v510, v753))
	*(*float32)(unsafe.Add(mBase, uint32(v819)+4)) = base.F32_demote_f64(base.F64_div(v705, v753))
	*(*float32)(unsafe.Add(mBase, uint32(v819))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v737), v753))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v749
	v832 = int32(1)
	v833 = v524 << (uint(v832) % 32)
	v835 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v739+v833))) = uint16(v835)
	v839 = l0 + v524<<(uint(v817)%32)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v839-int32(-64)))) = v842
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v839)+124)) = v762
	*(*int32)(unsafe.Add(mBase, uint32(v839)+84)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v839)+164)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v839)+104)) = v759
	*(*int32)(unsafe.Add(mBase, uint32(v839)+144)) = v736
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	*(*int32)(unsafe.Add(mBase, uint32(v839)+184)) = v850
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v833)+204)) = uint16(v853)
	v855 = l0 + v524
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v855)+214)) = uint8(v856)
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v855)+219)) = uint8(v858)
	v874 = v524 + v832
	goto L117
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756+v798))) = v805
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v801)+4))
	*(*float32)(unsafe.Add(mBase, uint32(v798+v762))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i32_s(v809), v753))
	v815 = v768 + int32(1)
	if v815 != v736 {
		v768 = v815
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	if v963 <= int32(0) {
		goto L78
	} else {
		goto L129
	}
L126:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v896)+376))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v896)+364))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v896)+352))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v896)+340))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v896)+328))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v896)+316))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v896)+304))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v896)+292))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v896)+280))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v896)+268))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v896)+256))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v896)+244))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v896)+232))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v896)+220))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v896)+208))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v896)+196))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v896)+184))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v896)+172))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v896)+160))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v896)+148))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v896)+136))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v896)+124))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v896)+112))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v896)+100))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v896)+88))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v896)+76))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v896-int32(-64))))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v896)+52))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v896)+40))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v896)+28))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v896)+16))
	v963 = v899 + (v900 + (v901 + (v902 + (v903 + (v904 + (v905 + (v906 + (v907 + (v908 + (v909 + (v910 + (v911 + (v912 + (v913 + (v914 + (v915 + (v916 + (v917 + (v918 + (v919 + (v920 + (v921 + (v922 + (v923 + (v924 + (v927 + (v928 + (v929 + (v930 + (v931 + v897))))))))))))))))))))))))))))))
	goto L128
L127:
	;
	v963 = v897
	goto L128
L128:
	;
	goto L125
L129:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v969 = F_palloc(m, v963<<(uint(int32(2))%32))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_hash_seq_init(m, v36+int32(156), v79)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v975 = int32(2)
	if v966 <= v975 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v978 = v975
	goto L134
L133:
	;
	v978 = v966
	goto L134
L134:
	;
	v981 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v981 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v988 = int32(0)
	v990 = v981
	goto L139
L137:
	;
	goto L138
L138:
	;
	v1060 = int32(0)
	F_qsort_interruptible(m, v969, v963, int32(4), int32(1268), v1060)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L143
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v969+v988<<(uint(int32(2))%32)))) = v990
	v1025 = F_hash_seq_search(m, v36+int32(156))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L138
L141:
	;
	if v1025 != 0 {
		v988 = v988 + int32(1)
		v990 = v1025
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1068 = v978 + int32(1)
	v1071 = F_MemoryContextAlloc(m, v1066, v1068<<(uint(int32(2))%32))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v1071+v978<<(uint(int32(2))%32)))) = base.F32_demote_f64(base.F64_div(base.F64_convert_i64_s(v505), base.F64_convert_i32_u(v492)))
	v1081 = int32(1)
	v1083 = base.I64_extend_i32_u(v978 - v1081)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v969)))
	v1085 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1084)+4)))
	v1095 = v1060
	v1100 = int32(0)
	v1117 = v1083 * v1085
	goto L145
L145:
	;
	if int64(0) < v1117 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	v1220 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v739+v874<<(uint(int32(1))%32)))) = uint16(v1220)
	v1224 = l0 + v874<<(uint(int32(2))%32)
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1224-int32(-64)))) = v1227
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+124)) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+84)) = v1229
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+104)) = v1068
	goto L78
L147:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1180)))
	*(*float32)(unsafe.Add(mBase, uint32(v1071+v1100<<(uint(int32(2))%32)))) = base.F32_convert_i32_s(v1210)
	v1215 = v1100 + int32(1)
	if v1215 != v978 {
		v1095 = v1178
		v1100 = v1215
		v1117 = v1200 - base.I64_extend_i32_u(v492-v1081)
		goto L145
	} else {
		goto L154
	}
L148:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v969+v1095<<(uint(int32(2))%32))))
	v1178 = v1095
	v1180 = v1129
	v1200 = v1117
	goto L147
L149:
	;
	goto L150
L150:
	;
	v1134 = v1095
	v1156 = v1117
	goto L151
L151:
	;
	v1164 = v1134 + int32(1)
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v969+v1164<<(uint(int32(2))%32))))
	v1169 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)))
	v1171 = v1169*v1083 + v1156
	if v1171 <= int64(0) {
		v1134 = v1164
		v1156 = v1171
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v1178 = v1164
	v1180 = v1168
	v1200 = v1171
	goto L147
L153:
	;
	goto L152
L154:
	;
	goto L146
L155:
	;
	F_errmsg_internal(m, int32(134306), int32(0))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(523300), int32(440), int32(133798))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_construct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(99)
	v13 = int32(1)
	if l2 <= int32(699) {
		switch l2 - int32(18) {
		case 0:
			v60 = int32(1)
			v61 = v12
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 1:
			v60 = int32(64)
			v61 = v12
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 2:
			v60 = int32(8)
			v61 = int32(100)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 3:
			v60 = int32(2)
			v61 = int32(115)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
				F_errmsg_internal(m, int32(715626), v10)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518800), int32(3463), int32(288564))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 5, 8, 10:
			v60 = int32(4)
			v61 = int32(105)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 7:
			v60 = int32(-1)
			v61 = int32(105)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 9:
			v60 = int32(6)
			v61 = int32(115)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		}
	} else {
		switch l2 - int32(700) {
		case 0:
			v60 = int32(4)
			v61 = int32(105)
			v62 = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		case 1:
			v60 = int32(8)
			v61 = int32(100)
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v64 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
			v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				m.G0 = v10 + int32(16)
				return v72
			}
		default:
			if l2 == int32(2206) {
				v60 = int32(4)
				v61 = int32(105)
				v62 = v13
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
				v64 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
				v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					m.G0 = v10 + int32(16)
					return v72
				}
			} else {
				if l2 != int32(2275) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
						F_errmsg_internal(m, int32(715626), v10)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518800), int32(3463), int32(288564))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v60 = int32(-2)
					v61 = v12
					v62 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
					v64 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v64
					v72 = F_construct_md_array(m, l0, int32(0), v64, v10+int32(12), v10+int32(8), l2, v60, v62, v61)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v72
					}
				}
			}
		}
	}
}
func F_get_array_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+96))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_parse_array(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v6 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v79
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = m.T0[v6].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v15 + int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v19 != int32(5) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if v12 != 0 {
		v79 = v12
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v22 = int32(11)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v31 = F_json_lex(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	v26 = int32(6)
	goto L14
L13:
	;
	v26 = v22
	goto L14
L14:
	;
	if v19 == int32(12) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v29 = v22
	goto L17
L16:
	;
	v29 = v26
	goto L17
L17:
	;
	return v29
L18:
	;
	if v31 != 0 {
		v79 = v31
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v33 == int32(6) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v66 = F_json_lex(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v36 = F_parse_array_element(m, l0, l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v36 != 0 {
		v79 = v36
		goto L3
	} else {
		goto L23
	}
L23:
	;
	goto L24
L24:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v42 != int32(7) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v79 = v58
	goto L3
L26:
	;
	if v42 == int32(6) {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v56 = F_json_lex(m, l0)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L36
	}
L29:
	;
	v47 = int32(11)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v51 = int32(7)
	goto L32
L31:
	;
	v51 = v47
	goto L32
L32:
	;
	if v42 == int32(12) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v54 = v47
	goto L35
L34:
	;
	v54 = v51
	goto L35
L35:
	;
	return v54
L36:
	;
	if v56 != 0 {
		v79 = v56
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v58 = F_parse_array_element(m, l0, l1)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v58 == int32(0) {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	goto L25
L40:
	;
	if v66 != 0 {
		v79 = v66
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v68 - int32(1)
	if v5 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = m.T0[v5].(func(*base.Module, int32) int32)(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v79 = int32(0)
	goto L3
L45:
	;
	if v73 != 0 {
		v79 = v73
		goto L3
	} else {
		goto L46
	}
L46:
	;
	goto L44
}
