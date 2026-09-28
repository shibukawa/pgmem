package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AllocSetAllocFromNewBlock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = v13 - v14
	if base.Ui32(int32(16)) <= base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v15
	goto L4
L2:
	;
	goto L3
L3:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v87 = v85 << (uint(int32(1)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if base.Ui32(v87) < base.Ui32(v88) {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v39 = int32(8)
	v40 = v24 - v39
	if base.Ui32(v39) < base.Ui32(v40) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v43 = int32(29) - base.I32_clz(v24-int32(9))
	goto L8
L7:
	;
	v43 = int32(0)
	goto L8
L8:
	;
	v45 = int32(8)
	v49 = base.B2i32(v40 != v45<<(uint(v43)%32))
	if v40 != v45<<(uint(v43)%32) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v50 = int32(4)<<(uint(v43)%32) + v45
	goto L11
L10:
	;
	v50 = v24
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v31 + v50
	v53 = v43 - v49
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = base.I64_extend_i32_u(v53)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v31-v12)<<(uint(int64(34))%64) | int64(3)
	v67 = l0 + int32(48) + v53<<(uint(int32(2))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v31
	v71 = v24 - v50
	if base.Ui32(int32(15)) < base.Ui32(v71) {
		v24 = v71
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	v90 = v87
	goto L15
L14:
	;
	v90 = v88
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v90
	v93 = int32(8) << (uint(l3) % 32)
	v95 = v93 + int32(32)
	v102 = v85
	goto L16
L16:
	;
	if base.Ui32(v102) < base.Ui32(v95) {
		v102 = v102 << (uint(int32(1)) % 32)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v110 = F_emscripten_builtin_malloc(m, v102)
	mBase = m.M
	v112 = base.B2i32(v110 == int32(0))
	if base.Ui32(v102) < base.Ui32(int32(_a_F_AllocSetAllocFromNewBlock_0)) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	goto L17
L19:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v165 + v142
	*(*int32)(unsafe.Add(mBase, uint32(v141)+16)) = v141 + v142
	v171 = v141 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v176
	if v176 != 0 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	v160 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	if v146 == int32(0) {
		goto L19
	} else {
		goto L33
	}
L22:
	;
	v141 = v110
	v142 = v102
	v146 = v112
	goto L21
L23:
	;
	goto L24
L24:
	;
	if v110 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v141 = v110
	v142 = v102
	v146 = v112
	goto L21
L26:
	;
	goto L27
L27:
	;
	v119 = v102
	goto L28
L28:
	;
	v127 = int32(base.Ui32(v119) >> (uint(int32(1)) % 32))
	if base.Ui32(v127) < base.Ui32(v95) {
		goto L20
	} else {
		goto L30
	}
L29:
	;
	v141 = v129
	v142 = v127
	v146 = v131
	goto L21
L30:
	;
	v129 = F_emscripten_builtin_malloc(m, v127)
	mBase = m.M
	v131 = base.B2i32(v129 == int32(0))
	if base.Ui32(v119) < base.Ui32(int32(_a_F_AllocSetAllocFromNewBlock_1)) {
		v141 = v129
		v142 = v127
		v146 = v131
		goto L21
	} else {
		goto L31
	}
L31:
	;
	if v129 == int32(0) {
		v119 = v127
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L20
L34:
	;
	return int32(0)
L35:
	;
	return v160
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v141
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v180 = v179
	goto L38
L37:
	;
	v180 = v171
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v141
	v183 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v180 + v93 + v183
	*(*int64)(unsafe.Add(mBase, uint32(v180))) = base.I64_extend_i32_u(l3<<(uint(int32(5))%32)) | base.I64_extend_i32_u(v180-v141)<<(uint(int64(34))%64) | int64(3)
	return v180 + v183
}
func F_CopyFromBinaryInFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_getTypeBinaryInputInfo(m, l1, v7+int32(12), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		F_fmgr_info(m, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_CopyFromCSVOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v52 int32
	_ = v52
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v216 int64
	_ = v216
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int64
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int64
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int64
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	v5 = int32(0)
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
		v236 = v34
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L22
	} else {
		goto L158
	}
L5:
	;
	m.G0 = v23 + int32(128)
	return v673
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L22
	} else {
		goto L154
	}
L7:
	;
	v267 = F_CopyReadAttributesCSV(m, l0)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L22
	} else {
		goto L57
	}
L8:
	;
	v673 = int32(0)
	goto L5
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v236 + int64(1)
	v241 = F_CopyReadLine(m, l0, int32(1))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L22
	} else {
		goto L54
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v38 == int32(0) {
		v236 = int64(0)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v38 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v87 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L13:
	;
	v45 = int32(1)
	goto L15
L14:
	;
	v45 = v38
	goto L15
L15:
	;
	if v45 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v87 = v38
	v92 = v5
	goto L12
L17:
	;
	goto L18
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+192)) = v68 + int64(1)
	v73 = F_CopyReadLine(m, l0, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v87 = v81
	v92 = v73
	goto L12
L21:
	;
	goto L20
L22:
	;
	return int32(0)
L23:
	;
	if v73 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v78 = v52 + int32(1)
	if v78 != v45 {
		v52 = v78
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	if v92 != 0 {
		goto L8
	} else {
		goto L53
	}
L27:
	;
	v104 = F_CopyReadAttributesCSV(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v106 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v104 == int32(0) {
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v104 != v111 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L4
L33:
	;
	v118 = int32(0)
	goto L34
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v134 <= v118 {
		goto L26
	} else {
		goto L36
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L49
	}
L36:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v141 = v118 << (uint(int32(2)) % 32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141+v142)))
	v149 = v33 + v136<<(uint(int32(3))%32) + v144*int32(100) - int32(72)
	v151 = v118 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v141+v152)))
	if v154 == int32(0) {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v158 = v149 + int32(4)
	if v158|v154 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v173 == int32(0) {
		v118 = v151
		goto L34
	} else {
		goto L48
	}
L39:
	;
	v164 = int32(-1)
	goto L41
L40:
	;
	v164 = int32(0)
	goto L41
L41:
	;
	if v158 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v165 = int32(1)
	goto L44
L43:
	;
	v165 = v164
	goto L44
L44:
	;
	v166 = int32(0)
	if base.B2i32(v158 == v166)|base.B2i32(v154 == v166) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = v165
	goto L47
L46:
	;
	v172 = F_strncmp(m, v158, v154, int32(64))
	mBase = m.M
	v173 = v172
	goto L47
L47:
	;
	goto L38
L48:
	;
	goto L35
L49:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+120)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v151
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_0), v23+int32(112))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(867), int32(_a_F_CopyFromCSVOneRow_2))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L22
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
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v236 = v216
	goto L9
L54:
	;
	if v241 == int32(0) {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v245 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	goto L8
L57:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v271 = int32(0)
	if base.B2i32(v29 < v267)&base.B2i32(v271 < v29) == v271 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v276 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L22
	} else {
		goto L150
	}
L61:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v278 <= int32(0) {
		v673 = int32(1)
		goto L5
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v673 = int32(1)
	goto L5
L64:
	;
	v281 = int32(0)
	if v281 < v267 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v285 = v267
	goto L67
L66:
	;
	v285 = v281
	goto L67
L67:
	;
	v296 = v281
	v304 = v5
	goto L68
L68:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v311 = v296 << (uint(int32(2)) % 32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311+v312)))
	v319 = v33 + v306<<(uint(int32(3))%32) + v314*int32(100) - int32(72)
	if v296 != v285 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	goto L63
L70:
	;
	v606 = v296 + int32(1)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v606 < v607 {
		v296 = v606
		v304 = v603
		goto L68
	} else {
		goto L149
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+200)) = int64(0)
	v603 = v594
	goto L70
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L22
	} else {
		goto L141
	}
L73:
	;
	v322 = v314 - int32(1)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v311+v269)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v325 != 0 {
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
	v539 = m.ExcPending
	if v539 != 0 {
		goto L22
	} else {
		goto L137
	}
L76:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322+v325))))
	if v327 != int32(1) {
		v603 = v304
		goto L70
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v324 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v392 = l2 + v322<<(uint(int32(3))%32)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v322))))
	if v395 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L81:
	;
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v322))) = uint8(v387)
	v389 = v385
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v319 + int32(4)
	v385 = v324
	goto L81
L83:
	;
	v375 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v319 + int32(4)
	v389 = v375
	goto L80
L84:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+v322))))
	if v334 == int32(0) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343+v322))))
	if v345 != int32(1) {
		goto L82
	} else {
		goto L89
	}
L87:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v319 + int32(4)
	if v337 != 0 {
		v385 = v337
		goto L81
	} else {
		goto L88
	}
L88:
	;
	v389 = int32(0)
	goto L80
L89:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	if base.B2i32(v351 == int32(0))|base.B2i32(v351 != v354) != 0 {
		v372 = v351
		v373 = v354
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v372-v373 != 0 {
		goto L82
	} else {
		goto L97
	}
L91:
	;
	goto L90
L92:
	;
	v357 = v324
	v358 = v348
	goto L93
L93:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v357)+1)))
	if v362 == int32(0) {
		v372 = v362
		v373 = v361
		goto L91
	} else {
		goto L95
	}
L94:
	;
	v372 = v362
	v373 = v361
	goto L91
L95:
	;
	v365 = int32(1)
	if v362 == v361 {
		v357 = v357 + v365
		v358 = v358 + v365
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	goto L83
L98:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v30+v322<<(uint(int32(2))%32))))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v401)+24))
	v404 = m.T0[v403].(func(*base.Module, int32, int32, int32) int64)(m, v401, l1, l3+v322)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L22
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v409 = v32 + v322*int32(28)
	v412 = v31 + v322<<(uint(int32(2))%32)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v319)+76))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v416 = F_InputFunctionCallSafe(m, v409, v389, v413, v414, v415, v392)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L22
	} else {
		goto L102
	}
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = v404
	v594 = v304
	goto L71
L102:
	;
	if v416 != 0 {
		v594 = v304
		goto L71
	} else {
		goto L103
	}
L103:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v418 - int32(1) {
	case 0:
		v443 = v304
		goto L105
	case 1:
		goto L106
	default:
		v448 = v304
		goto L104
	}
L104:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v449 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v444 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v444 + int64(1)
	v448 = v443
	goto L104
L106:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v422 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v421)+4)) = uint8(v422)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424+v322))))
	if v426 == int32(1) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v319)+76))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	v433 = F_InputFunctionCallSafe(m, v409, int32(0), v430, v431, v432, v392)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L22
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v438 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v322))) = uint8(v438)
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = int64(0)
	if v304 != 0 {
		v448 = v438
		goto L104
	} else {
		goto L112
	}
L110:
	;
	if v433 == int32(0) {
		goto L72
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v443 = v438
	goto L105
L113:
	;
	v452 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v452)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v454 != 0 {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	v532 = int32(1)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v533 - v532 {
	case 0:
		v673 = v532
		goto L5
	case 1:
		v603 = v448
		goto L70
	default:
		v594 = v448
		goto L71
	}
L116:
	;
	v525 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v525)
	goto L115
L117:
	;
	v455 = F_CopyLimitPrintoutLength(m, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L22
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v497 != int32(1) {
		goto L116
	} else {
		goto L132
	}
L120:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	switch v457 - int32(1) {
	case 0:
		goto L124
	case 1:
		goto L123
	default:
		goto L121
	}
L121:
	;
	F_pfree(m, v455)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L22
	} else {
		goto L131
	}
L122:
	;
	v478 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+44)) = v455
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v479
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v478
	F_errmsg(m, v476, v23+int32(32))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L22
	} else {
		goto L129
	}
L123:
	;
	v470 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L22
	} else {
		goto L127
	}
L124:
	;
	v462 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	if v462 == int32(0) {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v476 = int32(_a_F_CopyFromCSVOneRow_3)
	v477 = int32(1151)
	goto L122
L127:
	;
	if v470 == int32(0) {
		goto L121
	} else {
		goto L128
	}
L128:
	;
	v476 = int32(_a_F_CopyFromCSVOneRow_4)
	v477 = int32(1157)
	goto L122
L129:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), v477, int32(_a_F_CopyFromCSVOneRow_5))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L22
	} else {
		goto L130
	}
L130:
	;
	goto L121
L131:
	;
	goto L116
L132:
	;
	v502 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L22
	} else {
		goto L133
	}
L133:
	;
	if v502 == int32(0) {
		goto L116
	} else {
		goto L134
	}
L134:
	;
	v506 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v506
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_6), v23+int32(16))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L22
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(1166), int32(_a_F_CopyFromCSVOneRow_5))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L22
	} else {
		goto L136
	}
L136:
	;
	goto L116
L137:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L22
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v319 + int32(4)
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_7), v23)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L22
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(1019), int32(_a_F_CopyFromCSVOneRow_5))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L22
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L22
	} else {
		goto L142
	}
L142:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v562 = F_format_type_be(m, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L22
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v562
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_8), v23-int32(-64))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L22
	} else {
		goto L144
	}
L144:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	v572 = F_format_type_be(m, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L22
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v572
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v570
	v579 = F_errdetail(m, int32(_a_F_CopyFromCSVOneRow_9), v23+int32(48))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L22
	} else {
		goto L146
	}
L146:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	F_errdatatype(m, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L22
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(1117), int32(_a_F_CopyFromCSVOneRow_5))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L22
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	goto L69
L150:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L22
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_10), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L22
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(1004), int32(_a_F_CopyFromCSVOneRow_5))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L22
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
	F_errcode(m, int32(67240066))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L22
	} else {
		goto L155
	}
L155:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v149 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+100)) = v653
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v151
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_11), v23+int32(96))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L22
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(860), int32(_a_F_CopyFromCSVOneRow_2))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L22
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L22
	} else {
		goto L159
	}
L159:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v700 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v703 = v701
	goto L162
L161:
	;
	v703 = int32(0)
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v104
	F_errmsg(m, int32(_a_F_CopyFromCSVOneRow_12), v23+int32(80))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L22
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_CopyFromCSVOneRow_1), int32(844), int32(_a_F_CopyFromCSVOneRow_2))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L22
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromTextLikeStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v8 = F_palloc(m, int32(_a_F_CopyFromTextLikeStart_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+328)) = int64(0)
			v13 = v8
			v14 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v14)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v13
			F_initStringInfo(m, l0+int32(304))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v21 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
					v25 = v22 << (uint(int32(16)) % 32)
				} else {
					v25 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v25 >> (uint(int32(16)) % 32)
				v31 = F_palloc(m, v25>>(uint(int32(14))%32))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v31
					return
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
		v13 = v12
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v14)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+324)) = v13
		F_initStringInfo(m, l0+int32(304))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				v25 = v22 << (uint(int32(16)) % 32)
			} else {
				v25 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v25 >> (uint(int32(16)) % 32)
			v31 = F_palloc(m, v25>>(uint(int32(14))%32))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v31
				return
			}
		}
	}
}
func F_get_from_clause_item(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
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
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(112)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v21 - int32(63) {
	case 0:
		goto L4
	case 1:
		goto L3
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L8
	} else {
		goto L190
	}
L2:
	;
	m.G0 = v15 + int32(112)
	return
L3:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v438+v439<<(uint(int32(2))%32)-int32(4))))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v446&int32(1) != 0 {
		goto L118
	} else {
		goto L119
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v24<<(uint(int32(2))%32) - int32(4)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v28)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+124)))
	if v37 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	switch v43 {
	case 0:
		goto L13
	case 1:
		goto L19
	default:
		goto L14
	case 3:
		goto L18
	case 4:
		goto L17
	case 5:
		goto L16
	case 6:
		goto L15
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v344 != 0 {
		goto L2
	} else {
		goto L98
	}
L11:
	;
	F_get_column_alias_list(m, v32, l2)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L8
	} else {
		goto L97
	}
L12:
	;
	F_get_rte_alias(m, v36, v24, int32(0), l2)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L8
	} else {
		goto L96
	}
L13:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+20)))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v302 = F_generate_relation_name(m, v300, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L8
	} else {
		goto L91
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L88
	}
L15:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
	v279 = F_quote_identifier(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L86
	}
L16:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L83
	}
L17:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	switch v262 {
	case 0:
		goto L80
	case 1:
		goto L79
	default:
		goto L12
	}
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v61 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	F_get_query_def(m, v47, v17, v48, int32(0), int32(1), v51, v52, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)))
	if v245 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L24:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L70
	}
L25:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L8
	} else {
		goto L49
	}
L26:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L8
	} else {
		goto L47
	}
L27:
	;
	v75 = int32(0)
	goto L37
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	if v65 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if int32(0) < v61 {
		goto L27
	} else {
		goto L36
	}
L31:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+72)))
	if v66 != 0 {
		goto L27
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	F_get_rule_expr_funccall(m, v67, l2)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v234 = v64
	goto L23
L36:
	;
	v134 = v4
	goto L26
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60+v75<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 != int32(15) {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	v101 = int32(0)
	v104 = v101
	v109 = v101
	goto L43
L39:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v94 != int32(2331) {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v97 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v99 = v75 + int32(1)
	if v99 != v61 {
		v75 = v99
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v104<<(uint(int32(2))%32))))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+28))
	v122 = F_list_concat(m, v109, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L45
	}
L44:
	;
	v134 = v122
	goto L26
L45:
	;
	v125 = v104 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v125 < v126 {
		v104 = v125
		v109 = v122
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	F_get_rule_expr(m, v134, l2, int32(1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	goto L24
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	if v149 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	goto L24
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v152 <= int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	F_get_rule_expr_funccall(m, v157, l2)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	if v160 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_3))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v167 = int32(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v168 <= v167 {
		goto L50
	} else {
		goto L59
	}
L57:
	;
	F_get_from_clause_coldeflist(m, v156, int32(0), l2)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v175 = v167
	goto L60
L60:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183+v175<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_4))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L62
	}
L61:
	;
	goto L50
L62:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	F_get_rule_expr_funccall(m, v191, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	if v194 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_3))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L8
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v202 = v175 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v202 < v203 {
		v175 = v202
		goto L60
	} else {
		goto L69
	}
L67:
	;
	F_get_from_clause_coldeflist(m, v187, int32(0), l2)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	goto L61
L70:
	;
	v234 = int32(0)
	goto L23
L71:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_5))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_get_rte_alias(m, v36, v24, int32(0), l2)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L8
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	if v234 == int32(0) {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	if v256 == int32(0) {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	F_get_from_clause_coldeflist(m, v234, v32, l2)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	goto L10
L79:
	;
	F_get_json_table(m, v261, l2, int32(1))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	F_get_xmltable(m, v261, l2, int32(1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	goto L12
L82:
	;
	goto L12
L83:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	F_get_values_def(m, v272, l2)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	goto L12
L86:
	;
	F_appendStringInfoString(m, v17, v279)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	goto L12
L88:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v287
	F_errmsg_internal(m, int32(_a_F_get_from_clause_item_6), v15+int32(16))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_get_from_clause_item_7), int32(_a_F_get_from_clause_item_8), int32(_a_F_get_from_clause_item_9))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L8
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v302
	if v299 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v307 = int32(_a_F_get_from_clause_item_10)
	goto L94
L93:
	;
	v307 = int32(_a_F_get_from_clause_item_11)
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v307
	F_appendStringInfo(m, v17, int32(_a_F_get_from_clause_item_12), v15+int32(48))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	goto L12
L96:
	;
	goto L11
L97:
	;
	goto L10
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	if v345 == int32(0) {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+108)) = int32(2281)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v353 = int32(0)
	v359 = F_generate_function_name(m, v351, int32(1), v353, v15+int32(108), v353, v353, v353)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v359
	F_appendStringInfo(m, v348, int32(_a_F_get_from_clause_item_13), v15+int32(32))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	if v367 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_appendStringInfoChar(m, v348, int32(41))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L8
	} else {
		goto L112
	}
L103:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v370 <= int32(0) {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	F_get_rule_expr(m, v374, l2, int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v378 = int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v379 <= v378 {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v386 = v378
	goto L107
L107:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	F_appendStringInfoString(m, v348, int32(_a_F_get_from_clause_item_4))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L8
	} else {
		goto L109
	}
L108:
	;
	goto L102
L109:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v394+v386<<(uint(int32(2))%32))))
	F_get_rule_expr(m, v401, l2, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	v406 = v386 + int32(1)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v406 < v407 {
		v386 = v406
		goto L107
	} else {
		goto L111
	}
L111:
	;
	goto L108
L112:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	if v424 == int32(0) {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_appendStringInfoString(m, v348, int32(_a_F_get_from_clause_item_14))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	F_get_rule_expr(m, v430, l2, int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	F_appendStringInfoChar(m, v348, int32(41))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	goto L2
L117:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_get_from_clause_item(m, v468, l1, l2)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L8
	} else {
		goto L126
	}
L118:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	switch v450 - int32(63) {
	case 0:
		v457 = v4
		goto L121
	case 1:
		goto L123
	default:
		goto L122
	}
L119:
	;
	v462 = v4
	goto L120
L120:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L8
	} else {
		goto L125
	}
L121:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v458 == int32(0) {
		v467 = v457
		goto L117
	} else {
		goto L124
	}
L122:
	;
	v457 = int32(1)
	goto L121
L123:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v449)+32))
	v457 = base.B2i32(v453 == int32(0))
	goto L121
L124:
	;
	v462 = v457
	goto L120
L125:
	;
	v467 = v462
	goto L117
L126:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v472 {
	case 0:
		goto L128
	case 1:
		v495 = int32(_a_F_get_from_clause_item_15)
		goto L127
	case 2:
		goto L131
	case 3:
		goto L130
	default:
		goto L129
	}
L127:
	;
	F_appendContextKeyword(m, l2, v495, int32(-8), int32(8), int32(4))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L8
	} else {
		goto L138
	}
L128:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v493 != 0 {
		goto L135
	} else {
		goto L136
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L8
	} else {
		goto L132
	}
L130:
	;
	v495 = int32(_a_F_get_from_clause_item_16)
	goto L127
L131:
	;
	v495 = int32(_a_F_get_from_clause_item_17)
	goto L127
L132:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v479
	F_errmsg_internal(m, int32(_a_F_get_from_clause_item_18), v15-int32(-64))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_get_from_clause_item_7), int32(_a_F_get_from_clause_item_19), int32(_a_F_get_from_clause_item_9))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v494 = int32(_a_F_get_from_clause_item_20)
	goto L137
L136:
	;
	v494 = int32(_a_F_get_from_clause_item_21)
	goto L137
L137:
	;
	v495 = v494
	goto L127
L138:
	;
	if v467 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v513 != 0 {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_from_clause_item(m, v510, l1, l2)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L8
	} else {
		goto L146
	}
L143:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_get_from_clause_item(m, v504, l1, l2)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L8
	} else {
		goto L144
	}
L144:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	goto L139
L146:
	;
	goto L139
L147:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v628&int32(1) != 0 {
		goto L181
	} else {
		goto L182
	}
L148:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_22))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L8
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v588 != 0 {
		goto L168
	} else {
		goto L169
	}
L151:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v445)+44))
	if v517 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L8
	} else {
		goto L164
	}
L153:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	if v521 <= int32(0) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v517)+12))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)))
	v526 = F_quote_identifier(m, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L8
	} else {
		goto L155
	}
L155:
	;
	F_appendStringInfoString(m, v17, v526)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	if v530 <= int32(1) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	v534 = int32(1)
	goto L158
L158:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v517)+12))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v534<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_4))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L8
	} else {
		goto L160
	}
L159:
	;
	goto L152
L160:
	;
	v553 = F_quote_identifier(m, v549)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L8
	} else {
		goto L161
	}
L161:
	;
	F_appendStringInfoString(m, v17, v553)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L8
	} else {
		goto L162
	}
L162:
	;
	v558 = v534 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
	if v558 < v559 {
		v534 = v558
		goto L158
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v576 == int32(0) {
		goto L147
	} else {
		goto L165
	}
L165:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v576)+4))
	v580 = F_quote_identifier(m, v579)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L8
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v580
	F_appendStringInfo(m, v17, int32(_a_F_get_from_clause_item_23), v15+int32(96))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L8
	} else {
		goto L167
	}
L167:
	;
	goto L147
L168:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_24))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L8
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v610 == int32(0) {
		goto L147
	} else {
		goto L179
	}
L171:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v592&int32(1) == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	F_appendStringInfoChar(m, v17, int32(40))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L8
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_get_rule_expr(m, v600, l2, int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L8
	} else {
		goto L176
	}
L175:
	;
	goto L174
L176:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v604&int32(1) != 0 {
		goto L147
	} else {
		goto L177
	}
L177:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L178
	}
L178:
	;
	goto L147
L179:
	;
	F_appendStringInfoString(m, v17, int32(_a_F_get_from_clause_item_25))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L8
	} else {
		goto L180
	}
L180:
	;
	goto L147
L181:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v631 == int32(0) {
		goto L2
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	F_appendStringInfoChar(m, v17, int32(41))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L8
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v637 == int32(0) {
		goto L2
	} else {
		goto L186
	}
L186:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+12))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)+12))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v644+v645<<(uint(int32(2))%32)-int32(4))))
	v652 = F_quote_identifier(m, v651)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L8
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v652
	F_appendStringInfo(m, v17, int32(_a_F_get_from_clause_item_26), v15+int32(80))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L8
	} else {
		goto L188
	}
L188:
	;
	F_get_column_alias_list(m, v445, l2)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L8
	} else {
		goto L189
	}
L189:
	;
	goto L2
L190:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v681
	F_errmsg_internal(m, int32(_a_F_get_from_clause_item_27), v15)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L8
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_get_from_clause_item_7), int32(_a_F_get_from_clause_item_28), int32(_a_F_get_from_clause_item_9))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L8
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_prepare_sort_from_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v255 int32
	_ = v255
	var v257 float64
	_ = v257
	var v259 float64
	_ = v259
	var v261 float64
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v11 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = v29
	goto L3
L2:
	;
	v30 = v11
	goto L3
L3:
	;
	v33 = F_palloc(m, v30<<(uint(int32(1))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v38 = v30 << (uint(int32(2)) % 32)
	v39 = F_palloc(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v41 = F_palloc(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v43 = F_palloc(m, v30)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	if l1 == int32(0) {
		v335 = l0
		v350 = v11
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L4
	} else {
		goto L78
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L75
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v350
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v43
	m.G0 = v26 + int32(16)
	return v335
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v47 <= int32(0) {
		v335 = l0
		v350 = v11
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v50 = l0
	v54 = l4
	v61 = v28
	v65 = v11
	goto L14
L14:
	;
	v74 = v65 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v74+v75)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+41)))
	if v79 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v335 = v289
	v350 = v332
	goto L11
L16:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v314 = F_get_opfamily_member_for_cmptype(m, v312, v299, v299, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L4
	} else {
		goto L72
	}
L17:
	;
	v222 = int32(0)
	v224 = F_find_computable_ec_member(m, v222, v78, v61, l2, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L57
	}
L18:
	;
	v191 = F_get_sortgroupref_tle(m, v80, v61)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L55
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	if v80 != 0 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errmsg_internal(m, int32(_a_F_prepare_sort_from_pathkeys_0), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_prepare_sort_from_pathkeys_1), int32(_a_F_prepare_sort_from_pathkeys_2), int32(_a_F_prepare_sort_from_pathkeys_3))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v163 = v98
	goto L48
L27:
	;
	if v61 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3+v65<<(uint(int32(1))%32)))))
	if v61 != 0 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v98 = int32(0)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v98 < v99 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	if v143 == int32(0) {
		goto L17
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v109 <= int32(0) {
		v143 = int32(0)
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v143 = int32(0)
	goto L33
L37:
	;
	v112 = int32(0)
	if v112 < v109 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = v109
	goto L40
L39:
	;
	v115 = v112
	goto L40
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v120 = int32(0)
	goto L41
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v116+v120<<(uint(int32(2))%32))))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+8)))
	if v129 == v105&int32(_a_F_prepare_sort_from_pathkeys_4) {
		v143 = v128
		goto L33
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v132 = v120 + int32(1)
	if v132 != v115 {
		v120 = v132
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v148 = F_find_ec_member_matching_expr(m, v78, v147, l2)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	if v148 == int32(0) {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v289 = v50
	v293 = v54
	v299 = v152
	v300 = v61
	v302 = v143
	goto L16
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v163<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	v182 = F_find_ec_member_matching_expr(m, v78, v181, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L50
	}
L49:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v289 = v50
	v293 = v54
	v299 = v190
	v300 = v61
	v302 = v180
	goto L16
L50:
	;
	if v182 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v187 = v163 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v187 < v188 {
		v163 = v187
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L49
L54:
	;
	goto L17
L55:
	;
	if v191 == int32(0) {
		goto L17
	} else {
		goto L56
	}
L56:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v289 = v50
	v293 = v54
	v299 = v198
	v300 = v61
	v302 = v191
	goto L16
L57:
	;
	if v224 == int32(0) {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	if v54&int32(1) != 0 {
		v268 = v50
		v270 = v61
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v274 = F_copyObjectImpl(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L66
	}
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	switch v231 - int32(336) {
	case 0, 1, 2, 3, 4, 28, 29, 30, 35, 38, 39, 40, 41:
		goto L61
	default:
		v268 = v50
		v270 = v61
		goto L59
	case 23:
		goto L62
	}
L61:
	;
	v237 = F_copyObjectImpl(m, v61)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L64
	}
L62:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+80)))
	if v234&int32(4) != 0 {
		v268 = v50
		v270 = v61
		goto L59
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+37)))
	v241 = F_palloc0(m, int32(88))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+80)) = v243
	*(*int64)(unsafe.Add(mBase, uint32(v241)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+56)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v241)+52)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v241)+48)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v241)+44)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v241))) = int32(335)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+4)) = v255
	v257 = *(*float64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v241)+8)) = v257
	v259 = *(*float64)(unsafe.Add(mBase, uint32(v50)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v241)+16)) = v259
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v50)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v241)+24)) = v261
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+37)) = uint8(v239)
	*(*uint8)(unsafe.Add(mBase, uint32(v241)+36)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+32)) = v263
	v268 = v241
	v270 = v237
	goto L59
L66:
	;
	if v270 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270)+4)))
	v280 = v276 + int32(1)
	goto L69
L68:
	;
	v280 = int32(1)
	goto L69
L69:
	;
	v284 = F_makeTargetEntry(m, v274, base.I32_extend16_s(v280), int32(0), int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	v286 = F_lappend(m, v270, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v268)+44)) = v286
	v289 = v268
	v293 = int32(1)
	v299 = v228
	v300 = v286
	v302 = v284
	goto L16
L72:
	;
	if v314 == int32(0) {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v318 = int32(1)
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v33+v65<<(uint(v318)%32)))) = uint16(v321)
	*(*int32)(unsafe.Add(mBase, uint32(v39+v74))) = v314
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v74))) = v326
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v43))) = uint8(v329)
	v332 = v65 + v318
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v332 < v333 {
		v50 = v289
		v54 = v293
		v61 = v300
		v65 = v332
		goto L14
	} else {
		goto L74
	}
L74:
	;
	goto L15
L75:
	;
	F_errmsg_internal(m, int32(_a_F_prepare_sort_from_pathkeys_5), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_prepare_sort_from_pathkeys_1), int32(_a_F_prepare_sort_from_pathkeys_6), int32(_a_F_prepare_sort_from_pathkeys_3))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v384
	F_errmsg_internal(m, int32(_a_F_prepare_sort_from_pathkeys_7), v26)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_prepare_sort_from_pathkeys_1), int32(_a_F_prepare_sort_from_pathkeys_8), int32(_a_F_prepare_sort_from_pathkeys_3))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
