package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgss_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int64
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v189 int32
	_ = v189
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v236 int32
	_ = v236
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
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v310 int64
	_ = v310
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v321 int64
	_ = v321
	var v323 int64
	_ = v323
	var v324 int64
	_ = v324
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v342 int64
	_ = v342
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v356 int64
	_ = v356
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v398 int64
	_ = v398
	var v400 int64
	_ = v400
	var v401 int64
	_ = v401
	var v405 int64
	_ = v405
	var v407 int64
	_ = v407
	var v408 int64
	_ = v408
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v415 int64
	_ = v415
	var v419 int64
	_ = v419
	var v443 int64
	_ = v443
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v450 int64
	_ = v450
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v457 int64
	_ = v457
	var v459 int64
	_ = v459
	var v460 int64
	_ = v460
	var v464 int64
	_ = v464
	var v466 int64
	_ = v466
	var v467 int64
	_ = v467
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v490 int32
	_ = v490
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
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
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v590 int32
	_ = v590
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int64
	_ = v656
	var v657 int64
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	v5 = int32(0)
	v25 = int64(0)
	v30 = m.G0
	v32 = v30 - int32(80)
	m.G0 = v32
	v35 = l0 + int32(16)
	v42 = v5
	v43 = v5
	v44 = v5
	v45 = v5
	v46 = v5
	v47 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	v51 = v5
	v52 = v5
	v53 = v5
	v54 = v5
	v55 = int32(-1)
	v56 = v5
	v57 = v5
	v58 = v32
	v61 = v25
	v62 = v25
	goto L3
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	m.G0 = v32 + int32(80)
	return v680
L3:
	;
	goto L6
L4:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v533
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v530
	v674 = int32(4735024)
	v675 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v675 - int32(1)
	v680 = v669
	goto L2
L5:
	;
	goto L4
L6:
	;
	switch v55 - int32(1) {
	case 0:
		v524 = v42
		v525 = v43
		v526 = v44
		v527 = v45
		v528 = v46
		v529 = v47
		v530 = v48
		v531 = v49
		v532 = v50
		v533 = v51
		v534 = v54
		v536 = v58
		goto L9
	case 1:
		v168 = v42
		v169 = v43
		v170 = v44
		v171 = v45
		v172 = v46
		v173 = v47
		v174 = v49
		v175 = v50
		v176 = v52
		v177 = v53
		v178 = v54
		v180 = v58
		v182 = v61
		v183 = v62
		goto L11
	default:
		goto L12
	}
L7:
	;
	v669 = v607
	goto L5
L8:
	;
	v632 = int32(m.ExcTag)
	v633 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v632 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L9:
	;
	if v525 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L10:
	;
	v509 = int32(4735024)
	v510 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v510 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	v517 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	goto L43
L11:
	;
	if v169 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L12:
	;
	v68 = int32(16)
	v69 = v58 - v68
	m.G0 = v69
	v72 = v69 - v68
	m.G0 = v72
	v74 = int32(128)
	v75 = v72 - v74
	m.G0 = v75
	v78 = v75 - v74
	m.G0 = v78
	v80 = int32(32)
	v81 = v78 - v80
	m.G0 = v81
	v84 = v81 - v80
	m.G0 = v84
	v86 = int32(160)
	v87 = v84 - v86
	m.G0 = v87
	v90 = v87 - v86
	m.G0 = v90
	v93 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	if int32(0) <= v93 {
		v506 = v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[1465]))
	if v97 != int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v97 != int32(1) {
		v506 = v44
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if l1 == int32(0) {
		v506 = v44
		goto L10
	} else {
		goto L19
	}
L17:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	if v105 != 0 {
		v506 = v44
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1466])))
	if v112&int32(1) == int32(0) {
		v506 = v44
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v117 == int64(0) {
		v506 = v35
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L23
L22:
	;
	v124 = int32(4460456)
	v125 = *(*int64)(unsafe.Add(mBase, _consts[41]))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, _consts[43]))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+16)) = v127
	v129 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v129
	v131 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v122
	v149 = int32(1)
	F___clock_gettime(m, v149, v72)
	mBase = m.M
	v151 = int64(*(*int32)(unsafe.Add(mBase, uint32(v72)+8)))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	v153 = int32(4735024)
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v154 + v149
	v159 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	v161 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	goto L26
L23:
	;
	v122 = F__emscripten_memcpy_bulkmem(m, v75, int32(4460328), int32(128))
	mBase = m.M
	goto L25
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v32 + int32(4)
	goto L29
L27:
	;
	v168 = v84
	v169 = int32(0)
	v170 = v35
	v171 = v90
	v172 = v81
	v173 = v78
	v174 = v69
	v175 = v87
	v176 = v161
	v177 = v159
	v178 = v75
	v180 = v90
	v182 = v151
	v183 = v152
	goto L11
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v177
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v176
	v262 = int32(4735024)
	v263 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	v264 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v263 - v264
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	F___clock_gettime(m, v264, v174)
	mBase = m.M
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	v286 = int64(*(*int32)(unsafe.Add(mBase, uint32(v174)+8)))
	v290 = F__emscripten_memset_bulkmem(m, v173, base.I32_extend8_s(int32(0)), int32(128))
	mBase = m.M
	goto L39
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	v253 = F_standard_planner(m, l0, l2, l3)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		v625 = v180
		goto L8
	} else {
		goto L38
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v175
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1467]))
	if v189 == int32(0) {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v176
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v177
	v214 = int32(4735024)
	v215 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v215 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	F_pg_re_throw(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		v625 = v180
		goto L8
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	v208 = m.T0[v189].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		v625 = v180
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v255 = v56
	v256 = v208
	v257 = v208
	goto L30
L37:
	;
	goto L1
L38:
	;
	v255 = v253
	v256 = v57
	v257 = v253
	goto L30
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v290)))
	v309 = *(*int64)(unsafe.Add(mBase, _consts[103]))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
	*(*int64)(unsafe.Add(mBase, uint32(v290))) = v307 + (v309 - v310)
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v290)+8))
	v316 = *(*int64)(unsafe.Add(mBase, _consts[104]))
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v178)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+8)) = v314 + (v316 - v317)
	v321 = *(*int64)(unsafe.Add(mBase, uint32(v290)+16))
	v323 = *(*int64)(unsafe.Add(mBase, _consts[105]))
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v178)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+16)) = v321 + (v323 - v324)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v290)+24))
	v330 = *(*int64)(unsafe.Add(mBase, _consts[106]))
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v178)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+24)) = v328 + (v330 - v331)
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v290)+32))
	v337 = *(*int64)(unsafe.Add(mBase, _consts[107]))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v178)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+32)) = v335 + (v337 - v338)
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v290)+40))
	v344 = *(*int64)(unsafe.Add(mBase, _consts[108]))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v178)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+40)) = v342 + (v344 - v345)
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v290)+48))
	v351 = *(*int64)(unsafe.Add(mBase, _consts[109]))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v178)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+48)) = v349 + (v351 - v352)
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v290)+56))
	v358 = *(*int64)(unsafe.Add(mBase, _consts[110]))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v178)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+56)) = v356 + (v358 - v359)
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v290)+64))
	v365 = *(*int64)(unsafe.Add(mBase, _consts[111]))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v178)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+64)) = v363 + (v365 - v366)
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v290)+72))
	v372 = *(*int64)(unsafe.Add(mBase, _consts[112]))
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v178)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+72)) = v370 + (v372 - v373)
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v290)+80))
	v379 = *(*int64)(unsafe.Add(mBase, _consts[113]))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v178)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+80)) = v377 + (v379 - v380)
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v290)+88))
	v386 = *(*int64)(unsafe.Add(mBase, _consts[114]))
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v178)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+88)) = v384 + (v386 - v387)
	v391 = *(*int64)(unsafe.Add(mBase, uint32(v290)+96))
	v393 = *(*int64)(unsafe.Add(mBase, _consts[115]))
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v178)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+96)) = v391 + (v393 - v394)
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v290)+104))
	v400 = *(*int64)(unsafe.Add(mBase, _consts[116]))
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v178)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+104)) = v398 + (v400 - v401)
	v405 = *(*int64)(unsafe.Add(mBase, uint32(v290)+112))
	v407 = *(*int64)(unsafe.Add(mBase, _consts[117]))
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v178)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+112)) = v405 + (v407 - v408)
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v290)+120))
	v414 = *(*int64)(unsafe.Add(mBase, _consts[118]))
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v178)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v290)+120)) = v412 + (v414 - v415)
	goto L40
L40:
	;
	v419 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v168)+24)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v419
	*(*int64)(unsafe.Add(mBase, uint32(v168)+8)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v168)+16))
	v445 = *(*int64)(unsafe.Add(mBase, _consts[43]))
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v172)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v443 + (v445 - v446)
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	v452 = *(*int64)(unsafe.Add(mBase, _consts[47]))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v450 + (v452 - v453)
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v168)+8))
	v459 = *(*int64)(unsafe.Add(mBase, _consts[45]))
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v172)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+8)) = v457 + (v459 - v460)
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v168)+24))
	v466 = *(*int64)(unsafe.Add(mBase, _consts[41]))
	v467 = *(*int64)(unsafe.Add(mBase, uint32(v172)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+24)) = v464 + (v466 - v467)
	goto L41
L41:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v177
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v182
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v174
	v490 = int32(0)
	F_pgss_store(m, l1, v473, v472, v471, v490, base.F64_div(base.F64_convert_i64_s(v286-v182+(v285-v183)*int64(1000000000)), float64(1e+06)), int64(0), v290, v168, v490, v490, v490, v490)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		v625 = v180
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v680 = v257
	goto L2
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v32 + int32(4)
	goto L46
L44:
	;
	v524 = v84
	v525 = int32(0)
	v526 = v506
	v527 = v90
	v528 = v81
	v529 = v78
	v530 = v517
	v531 = v69
	v532 = v87
	v533 = v515
	v534 = v75
	v536 = v90
	goto L9
L46:
	;
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v531
	v607 = F_standard_planner(m, l0, l2, l3)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		v625 = v536
		goto L8
	} else {
		goto L54
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v527
	v543 = *(*int32)(unsafe.Add(mBase, _consts[1467]))
	if v543 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[88])) = v530
	*(*int32)(unsafe.Add(mBase, _consts[264])) = v533
	v568 = int32(4735024)
	v569 = *(*int32)(unsafe.Add(mBase, _consts[1464]))
	*(*int32)(unsafe.Add(mBase, _consts[1464])) = v569 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v531
	F_pg_re_throw(m)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		v625 = v536
		goto L8
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v530
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+28)) = v53
	*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v32)+56)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v32)+60)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v32)+68)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v32)+76)) = v531
	v562 = m.T0[v543].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		v625 = v536
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v669 = v562
	goto L5
L53:
	;
	goto L1
L54:
	;
	goto L7
L55:
	;
	v637 = int32(v633)
	m.G0 = v625
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v637)))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	if v32+int32(4) == v644 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	m.ExcPending = 1
	goto L64
L57:
	;
	if v647 != 0 {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	v647 = v646
	goto L60
L59:
	;
	v647 = int32(0)
	goto L60
L60:
	;
	goto L57
L61:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v32)+60))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	v656 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v42 = v652
	v43 = v639
	v44 = v655
	v45 = v654
	v46 = v651
	v47 = v650
	v48 = v663
	v49 = v648
	v50 = v653
	v51 = v662
	v52 = v659
	v53 = v658
	v54 = v649
	v55 = v647
	v56 = v661
	v57 = v660
	v58 = v625
	v61 = v657
	v62 = v656
	goto L3
L62:
	;
	goto L63
L63:
	;
	F___wasm_longjmp(m, v640, v639)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgss_shmem_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v162 int32
	_ = v162
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v320 int64
	_ = v320
	var v322 int64
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int64
	_ = v354
	var v356 int64
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	v1 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(544)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1456]))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.T0[v18].(func(*base.Module))(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1457])) = v22
	*(*int32)(unsafe.Add(mBase, _consts[1458])) = v22
	v28 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v32 = F_LWLockAcquire(m, v28+int32(2688), v22)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v39 = F_ShmemInitStruct(m, int32(131133), int32(56), v15+int32(543))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1458])) = v39
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+543)))
	if v42 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[1459]))
	if int32(0) < v46 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+508)) = int64(1855425871896)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[1460]))
	v185 = F_ShmemInitHash(m, int32(339876), v181, v181, v15+int32(492), int32(40))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L31
	}
L11:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v131 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v132 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v131)+40)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v131)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v131)+24)) = v132
	*(*int64)(unsafe.Add(mBase, uint32(v131)+16)) = int64(1024)
	*(*int64)(unsafe.Add(mBase, uint32(v131)+8)) = int64(4621819117588971520)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v126 + v53<<(uint(int32(7))%32)
	v146 = m.G0
	v147 = int32(16)
	v148 = v146 - v147
	m.G0 = v148
	F___gettimeofday(m, v148)
	mBase = m.M
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v148)+8)))
	m.G0 = v148 + v147
	goto L30
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[1461]))
	v53 = int32(214)
	v54 = v1
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L27
	}
L15:
	;
	v66 = v50 + v54*int32(68)
	v67 = int32(131133)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1462])))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 == int32(0) {
		v90 = v70
		v91 = v71
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	if v91-v90 == int32(0) {
		goto L11
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	if v70 != v71 {
		v90 = v70
		v91 = v71
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v75 = v66
	v76 = v67
	goto L21
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v79
		v91 = v80
		goto L18
	} else {
		goto L23
	}
L22:
	;
	v90 = v79
	v91 = v80
	goto L18
L23:
	;
	v83 = int32(1)
	if v79 == v80 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v66)+64))
	v98 = v54 + int32(1)
	if v98 != v46 {
		v53 = v95 + v53
		v54 = v98
		goto L15
	} else {
		goto L26
	}
L26:
	;
	goto L16
L27:
	;
	F_errmsg_internal(m, int32(473189), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(523489), int32(605), int32(418773))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+48)) = v152 + v151*int64(1000000) - int64(946684800000000)
	goto L10
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1457])) = v185
	v189 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v189+int32(2688))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _consts[96])))
	if v195 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_on_shmem_exit(m, int32(7699), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+543)))
	if v202 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	m.G0 = v15 + int32(544)
	return
L38:
	;
	v203 = int32(119666)
	v204 = F_unlink(m, v203)
	mBase = m.M
	v205 = int32(0)
	v210 = F_AllocateFile(m, v203, int32(34101))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L42
	}
L39:
	;
	if v498 != 0 {
		goto L111
	} else {
		goto L112
	}
L40:
	;
	F_errfinish(m, int32(518309), v491, int32(243939))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L4
	} else {
		goto L110
	}
L41:
	;
	v467 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L106
	}
L42:
	;
	if v210 == int32(0) {
		v454 = v205
		v456 = v205
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1463])))
	if v215 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v218 = F_FreeFile(m, v210)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v222 = F_AllocateFile(m, int32(119700), int32(242988))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	goto L37
L48:
	;
	v438 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L102
	}
L49:
	;
	v410 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L98
	}
L50:
	;
	if v222 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v227 != int32(44) {
		v399 = v205
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v233 = F_palloc(m, int32(2048))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	v230 = F_FreeFile(m, v210)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L37
L56:
	;
	v239 = F_fread(m, v15+int32(488), int32(4), int32(1), v222)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v239 != int32(1) {
		v399 = v233
		goto L49
	} else {
		goto L58
	}
L58:
	;
	v247 = F_fread(m, v15+int32(480), int32(4), int32(1), v222)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	if v247 != int32(1) {
		v399 = v233
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v255 = F_fread(m, v15+int32(484), int32(4), int32(1), v222)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if v255 != int32(1) {
		v399 = v233
		goto L49
	} else {
		goto L62
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v15)+488))
	if v259 != int32(539100168) {
		v427 = v233
		goto L48
	} else {
		goto L63
	}
L63:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v15)+480))
	if v262 != int32(1800) {
		v427 = v233
		goto L48
	} else {
		goto L64
	}
L64:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v15)+484))
	if int32(0) < v265 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v274 = v233
	v276 = int32(2048)
	v281 = v1
	goto L68
L66:
	;
	v369 = v233
	goto L67
L67:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v384 = F_fread(m, v379+int32(40), int32(16), int32(1), v222)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L4
	} else {
		goto L93
	}
L68:
	;
	v287 = F_fread(m, v15+int32(48), int32(432), int32(1), v222)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	v369 = v306
	goto L67
L70:
	;
	if v287 != int32(1) {
		v399 = v274
		goto L49
	} else {
		goto L71
	}
L71:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	if base.Ui32(int32(35)) <= base.Ui32(v291) {
		v427 = v274
		goto L48
	} else {
		goto L72
	}
L72:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v276 <= v294 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v296 = int32(1)
	v297 = v276 << (uint(v296) % 32)
	v299 = v294 + v296
	if v299 < v297 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v305 = v294
	v306 = v274
	v307 = v276
	goto L75
L75:
	;
	v308 = int32(1)
	v311 = F_fread(m, v306, v308, v305+v308, v222)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L80
	}
L76:
	;
	v301 = v297
	goto L78
L77:
	;
	v301 = v299
	goto L78
L78:
	;
	v302 = F_repalloc(m, v274, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	v305 = v304
	v306 = v302
	v307 = v301
	goto L75
L80:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v311 != v313+int32(1) {
		v399 = v306
		goto L49
	} else {
		goto L81
	}
L81:
	;
	v318 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v313+v306))) = uint8(v318)
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v15)+72))
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
	if v320 != int64(0)-v322 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+24))
	v328 = int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	v332 = F_fwrite(m, v306, v328, v329+v328, v210)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v363 = v281 + int32(1)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v15)+484))
	if v363 < v364 {
		v274 = v306
		v276 = v307
		v281 = v363
		goto L68
	} else {
		goto L92
	}
L85:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v15)+444))
	if v332 != v334+int32(1) {
		v454 = v222
		v456 = v306
		goto L41
	} else {
		goto L86
	}
L86:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[1458]))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v339)+24)) = v340 + v332
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v15)+448))
	v347 = F_entry_alloc(m, v15+int32(48), v327, v334, v345, int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	goto L89
L88:
	;
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v15)+456))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+408)) = v354
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v15)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v347)+416)) = v356
	goto L84
L89:
	;
	v352 = F__emscripten_memcpy_bulkmem(m, v347+int32(24), v15+int32(72), int32(368))
	mBase = m.M
	goto L91
L91:
	;
	goto L88
L92:
	;
	goto L69
L93:
	;
	if v384 != int32(1) {
		v399 = v369
		goto L49
	} else {
		goto L94
	}
L94:
	;
	F_pfree(m, v369)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v390 = F_FreeFile(m, v222)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	v392 = F_FreeFile(m, v210)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v395 = F_unlink(m, int32(119700))
	mBase = m.M
	goto L37
L98:
	;
	if v410 == int32(0) {
		v496 = v222
		v498 = v399
		goto L39
	} else {
		goto L99
	}
L99:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(119700)
	F_errmsg(m, int32(314376), v15+int32(16))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v480 = v222
	v482 = v399
	v491 = int32(701)
	goto L40
L102:
	;
	if v438 == int32(0) {
		v496 = v222
		v498 = v427
		goto L39
	} else {
		goto L103
	}
L103:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(119700)
	F_errmsg(m, int32(752409), v15+int32(32))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	v480 = v222
	v482 = v427
	v491 = int32(707)
	goto L40
L106:
	;
	if v467 == int32(0) {
		v496 = v454
		v498 = v456
		goto L39
	} else {
		goto L107
	}
L107:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(119666)
	F_errmsg(m, int32(314042), v15)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v480 = v454
	v482 = v456
	v491 = int32(713)
	goto L40
L110:
	;
	v496 = v480
	v498 = v482
	goto L39
L111:
	;
	F_pfree(m, v498)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if v496 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L113
L115:
	;
	v509 = F_FreeFile(m, v496)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v210 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v511 = F_FreeFile(m, v210)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v514 = F_unlink(m, int32(119700))
	mBase = m.M
	goto L37
L122:
	;
	goto L121
}
