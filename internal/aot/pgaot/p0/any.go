package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_ANY_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
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
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
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
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v647 int32
	_ = v647
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l2 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v647
L2:
	;
	v413 = int32(0)
	v415 = F_pull_varnos_of_level(m, v413, v20)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L12
	} else {
		goto L94
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 != int32(2) {
		v647 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v27 == int32(0) {
		v647 = v5
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	switch v30 - int32(17) {
	case 0:
		goto L9
	default:
		v647 = v5
		goto L1
	case 4:
		goto L8
	case 20:
		goto L7
	}
L6:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = F_flatten_join_alias_vars(m, l0, v171, v160)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L40
	}
L7:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v104 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if base.Ui32(int32(1)) < base.Ui32(v52) {
		v647 = v5
		goto L1
	} else {
		goto L16
	}
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v33 == int32(0) {
		v647 = v5
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 != int32(2) {
		v647 = v5
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v40 = F_op_is_safe_index_member(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	if v40 == int32(0) {
		v647 = v5
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = F_lappend(m, int32(0), v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v160 = v50
	goto L6
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v55 == int32(0) {
		v160 = v5
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v58 <= int32(0) {
		v160 = v5
		goto L6
	} else {
		goto L18
	}
L18:
	;
	v65 = v5
	v70 = v5
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76+v70<<(uint(int32(2))%32))))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v81 != int32(17) {
		v647 = v5
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v160 = v98
	goto L6
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	if v84 == int32(0) {
		v647 = v5
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 != int32(2) {
		v647 = v5
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v91 = F_op_is_safe_index_member(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	if v91 == int32(0) {
		v647 = v5
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = F_lappend(m, v65, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v101 = v70 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v101 < v102 {
		v65 = v98
		v70 = v101
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v154 = F_list_concat(m, int32(0), v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L38
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v107 <= int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v114 = v5
	goto L31
L31:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125+v114<<(uint(int32(2))%32))))
	v130 = F_op_is_safe_index_member(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L33
	}
L32:
	;
	v647 = int32(0)
	goto L1
L33:
	;
	if v130 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v133 = v114 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v133 < v134 {
		v114 = v133
		goto L31
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L32
L37:
	;
	goto L28
L38:
	;
	v160 = v154
	goto L6
L39:
	;
	v220 = int32(0)
	v223 = m.G0
	v225 = v223 - int32(416)
	m.G0 = v225
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v220
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
	if v229 != 0 {
		v385 = v220
		goto L47
	} else {
		goto L48
	}
L40:
	;
	if v172 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v181 = int32(0)
	goto L42
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v192 <= v181 {
		goto L39
	} else {
		goto L44
	}
L43:
	;
	v647 = int32(0)
	goto L1
L44:
	;
	v197 = int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v181<<(uint(int32(2))%32)+v199)))
	v203 = F_expr_is_nonnullable(m, l0, v201, v197)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	if v203 != 0 {
		v181 = v181 + v197
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	m.G0 = v225 + int32(416)
	if v385 != 0 {
		goto L2
	} else {
		goto L93
	}
L48:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v230 != 0 {
		v385 = v220
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v233 = int32(0)
	base.MemoryFill(m, v225+int32(16), v233, int32(400))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+20)) = v20
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v237 == v233 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v385 = int32(1)
	goto L47
L51:
	;
	goto L52
L52:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v242 <= int32(0) {
		v385 = int32(1)
		goto L47
	} else {
		goto L53
	}
L53:
	;
	v253 = v220
	v254 = v220
	v259 = v5
	goto L54
L54:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v254<<(uint(int32(2))%32))))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+26)))
	if v265 != 0 {
		v367 = v253
		v373 = v259
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v385 = v374
	goto L47
L56:
	;
	v374 = int32(1)
	v376 = v254 + v374
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v376 < v377 {
		v253 = v367
		v254 = v376
		v259 = v373
		goto L54
	} else {
		goto L92
	}
L57:
	;
	v270 = v264
	goto L58
L58:
	;
	v281 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v282 == v281 {
		v385 = v281
		goto L47
	} else {
		goto L60
	}
L59:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+45)))
	if v291 == int32(1) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v285 == int32(27) {
		v270 = v282
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v295 = F_flatten_group_exprs(m, int32(0), v20, v282)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L65
	}
L63:
	;
	v297 = v282
	goto L64
L64:
	;
	v298 = F_flatten_join_alias_vars(m, int32(0), v20, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L66
	}
L65:
	;
	v297 = v295
	goto L64
L66:
	;
	v301 = F_expr_is_nonnullable(m, v225+int32(16), v298, int32(2))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	if v301 != 0 {
		v367 = v253
		v373 = v259
		goto L56
	} else {
		goto L68
	}
L68:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	if v303 != int32(6) {
		v385 = v281
		goto L47
	} else {
		goto L69
	}
L69:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v298)+28))
	if v306 != 0 {
		v385 = v281
		goto L47
	} else {
		goto L70
	}
L70:
	;
	if v259 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	F_find_subquery_safe_quals(m, v309, v225+int32(12))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L74
	}
L72:
	;
	v322 = v253
	goto L73
L73:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v298)+8)))
	v327 = v325 + int32(7)
	if int32(0) <= v324|v327 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v316 = F_flatten_join_alias_vars(m, int32(0), v20, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = v316
	v320 = F_find_nonnullable_vars_walker(m, v316, int32(1))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v322 = v320
	goto L73
L77:
	;
	if v343 == int32(0) {
		v385 = v281
		goto L47
	} else {
		goto L91
	}
L78:
	;
	if v322 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L88
	}
L81:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v333 = v331
	goto L83
L82:
	;
	v333 = int32(0)
	goto L83
L83:
	;
	if v324 < v333 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335+v324<<(uint(int32(2))%32))))
	v340 = F_bms_is_member(m, v327, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L12
	} else {
		goto L87
	}
L85:
	;
	v343 = int32(0)
	goto L86
L86:
	;
	goto L77
L87:
	;
	v343 = v340
	goto L86
L88:
	;
	F_errmsg_internal(m, int32(_a_F_convert_ANY_sublink_to_join_0), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_convert_ANY_sublink_to_join_1), int32(132), int32(_a_F_convert_ANY_sublink_to_join_2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v367 = v322
	v373 = int32(1)
	goto L56
L92:
	;
	goto L55
L93:
	;
	v647 = int32(0)
	goto L1
L94:
	;
	v417 = int32(0)
	if v415 == v417 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v470 == int32(0) {
		v647 = v413
		goto L1
	} else {
		goto L109
	}
L96:
	;
	v470 = int32(1)
	goto L95
L97:
	;
	goto L98
L98:
	;
	if l3 == int32(0) {
		v463 = v417
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v470 = v463
	goto L95
L100:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v427 < v426 {
		v463 = v417
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v429 = int32(1)
	if v426 <= v429 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v432 = v429
	goto L104
L103:
	;
	v432 = v426
	goto L104
L104:
	;
	v433 = int32(8)
	v438 = int32(0)
	goto L105
L105:
	;
	v445 = v438 << (uint(int32(2)) % 32)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v415+v433+v445)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l3+v433+v445)))
	v452 = v447 & (v449 ^ int32(-1))
	v454 = base.B2i32(v452 == int32(0))
	if v452 != 0 {
		v463 = v454
		goto L99
	} else {
		goto L107
	}
L106:
	;
	v463 = v454
	goto L99
L107:
	;
	v456 = v438 + int32(1)
	if v456 != v432 {
		v438 = v456
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v474 = F_pull_varnos(m, l0, v473)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L12
	} else {
		goto L110
	}
L110:
	;
	if v474 == int32(0) {
		v647 = v413
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v478 = int32(0)
	if v474 == v478 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v531 == int32(0) {
		v647 = v413
		goto L1
	} else {
		goto L126
	}
L113:
	;
	v531 = int32(1)
	goto L112
L114:
	;
	goto L115
L115:
	;
	if l3 == int32(0) {
		v524 = v478
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v531 = v524
	goto L112
L117:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v474)+4))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v488 < v487 {
		v524 = v478
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v490 = int32(1)
	if v487 <= v490 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v493 = v490
	goto L121
L120:
	;
	v493 = v487
	goto L121
L121:
	;
	v494 = int32(8)
	v499 = int32(0)
	goto L122
L122:
	;
	v506 = v499 << (uint(int32(2)) % 32)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v474+v494+v506)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l3+v494+v506)))
	v513 = v508 & (v510 ^ int32(-1))
	v515 = base.B2i32(v513 == int32(0))
	if v513 != 0 {
		v524 = v515
		goto L116
	} else {
		goto L124
	}
L123:
	;
	v524 = v515
	goto L116
L124:
	;
	v517 = v499 + int32(1)
	if v517 != v493 {
		v499 = v517
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v535 = F_contain_volatile_functions(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L12
	} else {
		goto L127
	}
L127:
	;
	if v535 != 0 {
		v647 = v413
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v537 = int32(0)
	v539 = F_make_parsestate(m, v537)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L12
	} else {
		goto L129
	}
L129:
	;
	v541 = int32(0)
	v545 = F_addRangeTableEntryForSubquery(m, v539, v20, v541, base.B2i32(v415 != v541), v541)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L12
	} else {
		goto L130
	}
L130:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	v549 = F_lappend(m, v547, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L12
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v549
	if v549 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v553 = v552
	goto L134
L133:
	;
	v553 = v413
	goto L134
L134:
	;
	v555 = F_palloc0(m, int32(8))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(63)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
	if v560 == int32(0) {
		v603 = v537
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l0
	v620 = F_convert_testexpr_mutator(m, v615, v18+int32(8))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L12
	} else {
		goto L147
	}
L137:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	if v563 <= int32(0) {
		v603 = v537
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v570 = v537
	v571 = int32(0)
	v575 = v563
	goto L139
L139:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v582+v571<<(uint(int32(2))%32))))
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+26)))
	if v587 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v603 = v595
	goto L136
L141:
	;
	v590 = F_makeVarFromTargetEntry(m, v553, v586)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L12
	} else {
		goto L144
	}
L142:
	;
	v595 = v570
	v596 = v575
	goto L143
L143:
	;
	v598 = v571 + int32(1)
	if v598 < v596 {
		v570 = v595
		v571 = v598
		v575 = v596
		goto L139
	} else {
		goto L146
	}
L144:
	;
	v592 = F_lappend(m, v570, v590)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L12
	} else {
		goto L145
	}
L145:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v595 = v592
	v596 = v594
	goto L143
L146:
	;
	goto L140
L147:
	;
	v623 = F_palloc0(m, int32(40))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L12
	} else {
		goto L148
	}
L148:
	;
	v625 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v623)+32)) = v625
	*(*int32)(unsafe.Add(mBase, uint32(v623)+28)) = v620
	*(*int64)(unsafe.Add(mBase, uint32(v623)+20)) = v625
	*(*int32)(unsafe.Add(mBase, uint32(v623)+16)) = v555
	v631 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v623)+12)) = v631
	*(*uint8)(unsafe.Add(mBase, uint32(v623)+8)) = uint8(v631)
	if l2 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v637 = int32(5)
	goto L151
L150:
	;
	v637 = int32(4)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+4)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v623))) = int32(64)
	v647 = v623
	goto L1
}
func F_executeAnyItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v168 int32
	_ = v168
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	F_check_stack_depth(m)
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
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v18 + int32(48)
	return v168
L4:
	;
	v168 = int32(1)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v26 = F_JsonbIteratorInit(m, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v26
	v29 = int32(1)
	v45 = v29
	goto L8
L8:
	;
	v55 = F_JsonbIteratorNext(m, v18+int32(44), v18+int32(8), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v168 = int32(0)
	goto L3
L10:
	;
	if v64&int32(-2) != int32(2) {
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v62 = F_JsonbIteratorNext(m, v18+int32(44), v18+int32(8), int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	switch v55 {
	case 0:
		v168 = v45
		goto L3
	case 1:
		goto L11
	default:
		v64 = v55
		goto L10
	}
L13:
	;
	v64 = v62
	goto L10
L14:
	;
	if base.Ui32(l4) < base.Ui32(l5) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L9
L16:
	;
	if base.Ui32(l6) <= base.Ui32(l4) {
		v45 = v139
		goto L8
	} else {
		goto L44
	}
L17:
	;
	if l5&l6 != int32(-1) {
		v139 = v45
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v70 == int32(18) {
		v139 = v45
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if l7 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L36
	}
L25:
	;
	v87 = int32(2)
	if v85 == v87 {
		v168 = v87
		goto L3
	} else {
		goto L31
	}
L26:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v74)
	v78 = F_executeItemOptUnwrapTarget(m, l0, l1, v18+int32(8), l3, l8)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v83 = F_executeItemOptUnwrapTarget(m, l0, l1, v18+int32(8), l3, l8)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v73)
	v85 = v78
	goto L25
L30:
	;
	v85 = v83
	goto L25
L31:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v139 = v85
	goto L16
L33:
	;
	goto L34
L34:
	;
	if v85 != 0 {
		v139 = v85
		goto L16
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v93 < v94 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = v92 + v93<<(uint(int32(5))%32)
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+40)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+32)) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+16)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v107 + int32(1)
	v139 = v45
	goto L16
L38:
	;
	goto L39
L39:
	;
	v111 = int32(16)
	v113 = v94 << (uint(int32(1)) % 32)
	if v113 <= v111 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v116 = v111
	goto L42
L41:
	;
	v116 = v113
	goto L42
L42:
	;
	v121 = F_palloc(m, v116<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = int32(1)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+16)) = v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+24)) = v130
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+32)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+40)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v121
	v139 = v45
	goto L16
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v143 != int32(18) {
		v45 = v139
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v148 = F_executeAnyItem(m, l0, l1, v147, l3, l4+v29, l5, l6, l7, l8)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v148 == int32(2) {
		v168 = int32(2)
		goto L3
	} else {
		goto L47
	}
L47:
	;
	if v148|l3 != 0 {
		v45 = v148
		goto L8
	} else {
		goto L48
	}
L48:
	;
	goto L15
}
func F_has_any_column_privilege_id_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = F_textToQualifiedNameList(m, v7)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = F_makeRangeVarFromNameList(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					v18 = int32(0)
					v22 = F_RangeVarGetRelidExtended(m, v16, v18, v18, v18, v18)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v25 = F_convert_any_priv_string(m, v12, int32(_a_F_has_any_column_privilege_id_name_0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							v27 = F_pg_class_aclcheck(m, v22, v5, v25)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								if v27 == int32(0) {
									return int64(1)
								} else {
									v34 = F_pg_attribute_aclcheck_all(m, v22, v5, v25, int32(1))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int64(0)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v34 == int32(0)))
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
