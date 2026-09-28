package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EA_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v58 int32
	_ = v58
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v7 != 0 {
		if l2 == int32(0) {
			return
		} else {
			base.MemoryCopy(m, l1, v7, l2)
			return
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v16 = base.I32_div_s(v12+int32(7), int32(8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v25 != 0 {
			v26 = (v16 + v11<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		} else {
			v26 = int32(0)
		}
		if l2 != 0 {
			base.MemoryFill(m, l1, int32(0), l2)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v11
		v31 = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2 << (uint(v31) % 32)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v34
		v37 = l1 + int32(16)
		v39 = v11 << (uint(v31) % 32)
		v40 = int32(0)
		v41 = base.B2i32(v39 == v40)
		if v41 == v40 {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			base.MemoryCopy(m, v37, v44, v39)
		} else {
		}
		if v41 == int32(0) {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			base.MemoryCopy(m, v39+v37, v49, v39)
		} else {
		}
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
		v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
		v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+47)))
		F_CopyArrayEls(m, l1, v51, v52, v12, v53, v54, v55, int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			return
		}
	}
}
func F_ExecASTruncateTriggers(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v4 == int32(0) {
		return
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+24)))
		if v7 != int32(1) {
			return
		} else {
			v10 = int32(0)
			F_AfterTriggerSaveEvent(m, l0, l1, v10, v10, int32(3), v10, v10, v10, v10, v10, v10, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ExecBuildHash32FromAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v368 int64
	_ = v368
	var v370 int64
	_ = v370
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v487 int64
	_ = v487
	var v489 int64
	_ = v489
	var v491 int64
	_ = v491
	var v493 int64
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int64
	_ = v538
	var v540 int64
	_ = v540
	var v542 int64
	_ = v542
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v550 int32
	_ = v550
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int64
	_ = v604
	var v606 int64
	_ = v606
	var v608 int64
	_ = v608
	var v610 int64
	_ = v610
	var v612 int64
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	v9 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v24 = F_palloc0(m, int32(72))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(386)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = l6
	if int64(2) <= base.I64_extend_i32_s(l4)+base.I64_extend_i32_u(base.B2i32(l7 != int32(0))) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v49 = F_palloc(m, int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v51 = v9
	goto L5
L5:
	;
	if l4 <= int32(0) {
		v160 = v9
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v51 = v49
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = l1
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+28)) = uint8(v169)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l0
	v176 = v21 + int32(8)
	v180 = m.G0
	v182 = v180 - int32(16)
	m.G0 = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v169)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+24))
	if v187 != 0 {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	v55 = l4 & int32(3)
	v56 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(l4) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v67 = v56
	v71 = v9
	v78 = v9
	goto L12
L10:
	;
	v108 = v56
	v112 = v9
	goto L11
L11:
	;
	v126 = v108
	v130 = v112
	v134 = v9
	goto L28
L12:
	;
	v79 = base.I32_extend16_s(v71)
	v82 = l5 + v67<<(uint(int32(1))%32)
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
	if v83 < v79 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v55 == int32(0) {
		v160 = v94
		goto L7
	} else {
		goto L27
	}
L14:
	;
	v85 = v79
	goto L16
L15:
	;
	v85 = v83
	goto L16
L16:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+2)))
	if v86 < v85 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v88 = v85
	goto L19
L18:
	;
	v88 = v86
	goto L19
L19:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+4)))
	if v89 < v88 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = v88
	goto L22
L21:
	;
	v91 = v89
	goto L22
L22:
	;
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82)+6)))
	if v92 < v91 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v94 = v91
	goto L25
L24:
	;
	v94 = v92
	goto L25
L25:
	;
	v95 = int32(4)
	v96 = v67 + v95
	v98 = v78 + v95
	if v98 != l4&int32(2147483644) {
		v67 = v96
		v71 = v94
		v78 = v98
		goto L12
	} else {
		goto L26
	}
L26:
	;
	goto L13
L27:
	;
	v108 = v96
	v112 = v94
	goto L11
L28:
	;
	v138 = base.I32_extend16_s(v130)
	v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v126<<(uint(int32(1))%32)))))
	if v142 < v138 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v160 = v144
	goto L7
L30:
	;
	v144 = v138
	goto L32
L31:
	;
	v144 = v142
	goto L32
L32:
	;
	v145 = int32(1)
	v148 = v134 + v145
	if v148 != v55 {
		v126 = v126 + v145
		v130 = v144
		v134 = v148
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	if v274 != 0 {
		goto L62
	} else {
		goto L63
	}
L35:
	;
	m.G0 = v182 + int32(16)
	goto L34
L36:
	;
	v274 = int32(1)
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+28)) = v249
	v263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+20)) = uint8(v263)
	*(*int32)(unsafe.Add(mBase, uint32(v176)+24)) = v248
	if v249 != int32(_a_F_ExecBuildHash32FromAttrs_0) {
		goto L36
	} else {
		goto L61
	}
L38:
	;
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+20)) = uint8(v258)
	*(*int64)(unsafe.Add(mBase, uint32(v176)+24)) = int64(0)
	goto L36
L39:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)))
	if base.B2i32(v248 == int32(0))|base.B2i32(v252 != int32(1)) != 0 {
		goto L38
	} else {
		goto L59
	}
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v176)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(base.B2i32(v188 != int32(0)))
	v248 = v187
	v249 = v188
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v184 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	switch v194 - int32(2) {
	case 0:
		goto L46
	case 1:
		goto L45
	default:
		goto L44
	}
L44:
	;
	if base.Ui32(int32(2)) < base.Ui32(v194-int32(4)) {
		goto L38
	} else {
		goto L57
	}
L45:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v184)+36))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+101)))
	if v218 != int32(1) {
		goto L52
	} else {
		goto L53
	}
L46:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v184)+40))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+102)))
	if v198 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v197 == int32(0) {
		goto L38
	} else {
		goto L51
	}
L48:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+98)))
	if v201 != int32(1) {
		goto L38
	} else {
		goto L49
	}
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v184)+88))
	if v204 == int32(0) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v207)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v248 = v209
	v249 = v204
	goto L39
L51:
	;
	v215 = F_ExecGetResultSlotOps(m, v197, v182+int32(15))
	mBase = m.M
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v248 = v216
	v249 = v215
	goto L39
L52:
	;
	if v217 == int32(0) {
		goto L38
	} else {
		goto L56
	}
L53:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+97)))
	if v221 != int32(1) {
		goto L38
	} else {
		goto L54
	}
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v184)+84))
	if v224 == int32(0) {
		goto L52
	} else {
		goto L55
	}
L55:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v217)+56))
	v248 = v229
	v249 = v224
	goto L39
L56:
	;
	v235 = F_ExecGetResultSlotOps(m, v217, v182+int32(15))
	mBase = m.M
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v217)+56))
	v248 = v236
	v249 = v235
	goto L39
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v184)+80))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v184)+76))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+100)))
	if v243 != int32(1) {
		v248 = v242
		v249 = v241
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v182)+15)) = uint8(v246)
	v248 = v242
	v249 = v241
	goto L39
L59:
	;
	if v249 != 0 {
		goto L37
	} else {
		goto L60
	}
L60:
	;
	goto L38
L61:
	;
	v274 = int32(0)
	goto L35
L62:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v278 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	goto L64
L64:
	;
	if l7 != 0 {
		goto L75
	} else {
		goto L76
	}
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v301 + int32(1)
	v307 = v300 + v301*int32(40)
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+32)) = v308
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+24)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+16)) = v312
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+8)) = v314
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v316
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v298
	v300 = v298
	goto L65
L67:
	;
	v281 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v281
	v285 = F_palloc_mul(m, int32(40), v281)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v287 != v278 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v298 = v285
	goto L66
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v300 = v289
	goto L65
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v278 << (uint(int32(1)) % 32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v296 = F_repalloc(m, v293, v278*int32(80))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v298 = v296
	goto L66
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(85)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = base.I64_extend_i32_u(l7)
	v329 = base.B2i32(int32(0) < l4)
	if int32(0) < l4 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v380 = int32(86)
	goto L77
L77:
	;
	if int32(0) < l4 {
		goto L94
	} else {
		goto L95
	}
L78:
	;
	v330 = v51 + int32(8)
	goto L80
L79:
	;
	v330 = v24 + int32(5)
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v330
	if int32(0) < l4 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v334 = v51
	goto L83
L82:
	;
	v334 = v24 + int32(8)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v334
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v336 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v359 + int32(1)
	v365 = v358 + v359*int32(40)
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+32)) = v366
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+24)) = v368
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+16)) = v370
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v365)+8)) = v372
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v365))) = v374
	v380 = int32(88)
	goto L77
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v356
	v358 = v356
	goto L84
L86:
	;
	v339 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v339
	v343 = F_palloc_mul(m, int32(40), v339)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v345 != v336 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v356 = v343
	goto L85
L90:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v358 = v347
	goto L84
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v336 << (uint(int32(1)) % 32)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v354 = F_repalloc(m, v351, v336*int32(80))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v356 = v354
	goto L85
L94:
	;
	v385 = int32(8)
	v399 = v380
	v402 = int32(0)
	goto L97
L95:
	;
	goto L96
L96:
	;
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = int64(0)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v574 == v570 {
		goto L129
	} else {
		goto L130
	}
L97:
	;
	v413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5+v402<<(uint(int32(1))%32)))))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l3+v402<<(uint(int32(2))%32))))
	v419 = F_palloc0(m, int32(40))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L99
	}
L98:
	;
	goto L96
L99:
	;
	v421 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v419)+18)) = uint16(v421)
	v423 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v419)+16)) = uint8(v423)
	*(*int32)(unsafe.Add(mBase, uint32(v419)+12)) = v417
	*(*int64)(unsafe.Add(mBase, uint32(v419)+4)) = int64(0)
	v430 = l2 + v402*int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v419 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v419 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(7)
	v442 = base.I32_extend16_s(v413 - v421)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v442
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0+v444<<(uint(int32(3))%32)+v442*int32(100))+96))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v451
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v455 == v423 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v478 + int32(1)
	v484 = v477 + v478*int32(40)
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+32)) = v485
	v487 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+24)) = v487
	v489 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+16)) = v489
	v491 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+8)) = v491
	v493 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v484))) = v493
	v495 = base.B2i32(v402 == l4-int32(1))
	if v402 == l4-int32(1) {
		goto L110
	} else {
		goto L111
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v475
	v477 = v475
	goto L100
L102:
	;
	v458 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v458
	v462 = F_palloc_mul(m, int32(40), v458)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v464 != v455 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v475 = v462
	goto L101
L106:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v477 = v466
	goto L100
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v455 << (uint(int32(1)) % 32)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v473 = F_repalloc(m, v470, v455*int32(80))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v475 = v473
	goto L101
L110:
	;
	v496 = v24 + int32(5)
	goto L112
L111:
	;
	v496 = v51 + v385
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v496
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v419
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v51
	if v402 == l4-int32(1) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v501 = v24 + v385
	goto L115
L114:
	;
	v501 = v51
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v430
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v504
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v508 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v532 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v531 + v532
	v537 = v530 + v531*int32(40)
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+32)) = v538
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+24)) = v540
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+16)) = v542
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v537)+8)) = v544
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v537))) = v546
	v550 = v402 + v532
	if v550 != l4 {
		v399 = int32(88)
		v402 = v550
		goto L97
	} else {
		goto L126
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v528
	v530 = v528
	goto L116
L118:
	;
	v511 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v511
	v515 = F_palloc_mul(m, int32(40), v511)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v517 != v508 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v528 = v515
	goto L117
L122:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v530 = v519
	goto L116
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v508 << (uint(int32(1)) % 32)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v526 = F_repalloc(m, v523, v508*int32(80))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v528 = v526
	goto L117
L126:
	;
	goto L98
L127:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v597 + int32(1)
	v603 = v596 + v597*int32(40)
	v604 = *(*int64)(unsafe.Add(mBase, uint32(v21)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+32)) = v604
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+24)) = v606
	v608 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+16)) = v608
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v603)+8)) = v610
	v612 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v603))) = v612
	v614 = F_jit_compile_expr(m, v24)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L137
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v594
	v596 = v594
	goto L127
L129:
	;
	v577 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v577
	v581 = F_palloc_mul(m, int32(40), v577)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v583 != v574 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	v594 = v581
	goto L128
L133:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v596 = v585
	goto L127
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+40)) = v574 << (uint(int32(1)) % 32)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v592 = F_repalloc(m, v589, v574*int32(80))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v594 = v592
	goto L128
L137:
	;
	if v614 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	F_ExecReadyInterpretedExpr(m, v24)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	m.G0 = v21 + int32(48)
	return v24
L141:
	;
	goto L140
}
func F_ExecCheckPermissionsModified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v5 = int32(0)
	if l2 == v5 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L6
	} else {
		goto L26
	}
L2:
	;
	return v99
L3:
	;
	v11 = F_pg_attribute_aclcheck_all(m, l0, l1, l3, int32(1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v20 = int32(-1)
	goto L9
L6:
	;
	return int32(0)
L7:
	;
	if v11 != 0 {
		v99 = v5
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if l2 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v99 = v80
	goto L2
L11:
	;
	v80 = int32(base.Ui32(v78) >> (uint(int32(31)) % 32))
	if v78 < int32(0) {
		v99 = v80
		goto L2
	} else {
		goto L22
	}
L12:
	;
	v78 = base.I32_ctz(v64) | v65<<(uint(int32(5))%32)
	goto L11
L13:
	;
	v78 = int32(-2)
	goto L11
L14:
	;
	v29 = v20 + int32(1)
	v31 = int32(base.Ui32(v29) >> (uint(int32(5)) % 32))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 <= v31 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v35 = l2 + int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(int32(2))%32))))
	v42 = v39 & (int32(-1) << (uint(v29) % 32))
	if v42 != 0 {
		v64 = v42
		v65 = v31
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v44 = v31 + int32(1)
	if v44 == v32 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v47 = v44
	goto L18
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v35+v47<<(uint(int32(2))%32))))
	if v54 != 0 {
		v64 = v54
		v65 = v47
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L13
L20:
	;
	v56 = v47 + int32(1)
	if v56 != v32 {
		v47 = v56
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v84 = v78 - int32(7)
	if v84&int32(_a_F_ExecCheckPermissionsModified_0) == int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v90 = F_pg_attribute_aclcheck(m, l0, base.I32_extend16_s(v84), l1, l3)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	if v90 == int32(0) {
		v20 = v78
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	F_errmsg_internal(m, int32(_a_F_ExecCheckPermissionsModified_1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ExecCheckPermissionsModified_2), int32(791), int32(_a_F_ExecCheckPermissionsModified_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecComputeStoredGenerated(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v24 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v29 = v24
	goto L3
L3:
	;
	if l3 == int32(2) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	v29 = v27
	goto L3
L6:
	;
	m.G0 = v19 + int32(16)
	return
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v50 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v32 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v44 = l0 + int32(132)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v45 != 0 {
		v48 = v44
		goto L7
	} else {
		goto L16
	}
L11:
	;
	F_ExecInitGenerated(m, l0, l1, int32(2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v48 = l0 + int32(136)
	goto L7
L16:
	;
	F_ExecInitGenerated(m, l0, l1, l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v48 = v44
	goto L7
L18:
	;
	v53 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	v55 = v50
	goto L20
L20:
	;
	v56 = int32(_a_F_ExecComputeStoredGenerated_0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_ExecComputeStoredGenerated[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecComputeStoredGenerated[0])) = v59
	v62 = F_palloc_mul(m, int32(8), v23)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	v55 = v53
	goto L20
L22:
	;
	v65 = F_palloc_mul(m, int32(1), v23)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v69 < v68 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	m.T0[v72].(func(*base.Module, int32, int32))(m, l2, v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v75 = int32(0)
	v76 = base.B2i32(v23 == v75)
	if v76 == v75 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	base.MemoryCopy(m, v65, v79, v23)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if int32(0) < v23 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v86 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	m.T0[v165].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L48
	}
L34:
	;
	v103 = v86 << (uint(int32(3)) % 32)
	v104 = v22 + int32(28) + v103
	v107 = v49 + v86<<(uint(int32(2))%32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L33
L36:
	;
	v146 = v86 + int32(1)
	if v146 != v23 {
		v86 = v146
		goto L34
	} else {
		goto L47
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = l2
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)+24))
	v115 = m.T0[v114].(func(*base.Module, int32, int32, int32) int64)(m, v111, v29, v19+int32(15))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v65))))
	if v132 != 0 {
		goto L36
	} else {
		goto L45
	}
L40:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v117 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+2)))
	v122 = F_datumCopy(m, v115, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	v125 = int32(1)
	v126 = v115
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62+v103))) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v86+v65))) = uint8(v125)
	goto L36
L44:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	v125 = v124
	v126 = v122
	goto L43
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v134+v103)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
	v138 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+2)))
	v139 = F_datumCopy(m, v136, v137, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v62+v103))) = v139
	goto L36
L47:
	;
	goto L35
L48:
	;
	v169 = v23 << (uint(int32(3)) % 32)
	if v169 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	base.MemoryCopy(m, v170, v62, v169)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if v76 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	base.MemoryCopy(m, v174, v65, v23)
	goto L54
L53:
	;
	goto L54
L54:
	;
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v178 = v176 & int32(_a_F_ExecComputeStoredGenerated_1)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v181)
	goto L55
L55:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+28))
	m.T0[v184].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecComputeStoredGenerated[0])) = v57
	goto L6
}
func F_ExecCteScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(751), int32(752))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExecDelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
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
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v279 int32
	_ = v279
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	v11 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v25)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l7 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	goto L6
L5:
	;
	goto L6
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v29 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L18
	} else {
		goto L148
	}
L8:
	;
	m.G0 = v21 + int32(32)
	return v418
L9:
	;
	if l6 != 0 {
		goto L95
	} else {
		goto L96
	}
L10:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v68 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+18)))
	if v32 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+188))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v58 = v29
	goto L14
L14:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+20)))
	if v59 != int32(1) {
		goto L10
	} else {
		goto L25
	}
L15:
	;
	F_ExecPendingInserts(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v42 = v35
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	v50 = F_ExecBRDeleteTriggers(m, v42, v43, l1, l2, l3, l9, l7, l0+int32(16), base.B2i32(v47 == int32(5)))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L18
	} else {
		goto L20
	}
L18:
	;
	return int32(0)
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v42 = v41
	goto L17
L20:
	;
	if v50 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v418 = int32(0)
	goto L8
L22:
	;
	goto L23
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v55 == int32(0) {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v58 = v55
	goto L14
L25:
	;
	v63 = F_ExecIRDeleteTriggers(m, v24, l1, l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	if v63 == int32(0) {
		v418 = int32(0)
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v279 = v11
	goto L9
L28:
	;
	v72 = l0 + int32(16)
	goto L36
L29:
	;
	goto L30
L30:
	;
	v251 = F_ExecGetReturningSlot(m, v24, l1)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L18
	} else {
		goto L86
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L18
	} else {
		goto L82
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L18
	} else {
		goto L79
	}
L33:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDelete[0]))
	if v203 < int32(2) {
		v418 = int32(0)
		goto L8
	} else {
		goto L74
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L18
	} else {
		goto L71
	}
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	if v163 == v164 {
		v418 = v139
		goto L8
	} else {
		goto L65
	}
L36:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+64))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+188))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+96))
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v91, l2, v93, l5, v94, v95, int32(1), v72)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v153
	v418 = v139
	goto L8
L38:
	;
	if l7 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v99
	goto L41
L40:
	;
	goto L41
L41:
	;
	if v99 != int32(3) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	switch v99 {
	case 0:
		v279 = v11
		goto L9
	default:
		goto L32
	case 2:
		goto L45
	case 4:
		goto L33
	}
L43:
	;
	goto L44
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDelete[0]))
	if int32(2) <= v129 {
		goto L31
	} else {
		goto L54
	}
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	if v104 == v105 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v418 = int32(0)
	goto L8
L47:
	;
	goto L48
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_ExecDelete_0), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_errhint(m, int32(_a_F_ExecDelete_1), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1722), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_EvalPlanQualBegin(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v137 = F_EvalPlanQualSlot(m, v135, v23, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v24)+64))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+104))
	v147 = m.T0[v146].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v23, l2, v140, v137, v141, int32(3), v139, int32(2), v72)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	if v147 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	switch v147 - int32(2) {
	case 0:
		goto L35
	default:
		goto L34
	case 2:
		v418 = v139
		goto L8
	}
L59:
	;
	goto L60
L60:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v153 = F_EvalPlanQual(m, v151, v23, v152, v137)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	if v153 == int32(0) {
		v418 = v139
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+4)))
	if v157&int32(2) != 0 {
		v418 = v139
		goto L8
	} else {
		goto L63
	}
L63:
	;
	if l9 == int32(0) {
		goto L36
	} else {
		goto L64
	}
L64:
	;
	goto L37
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_ExecDelete_0), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	F_errhint(m, int32(_a_F_ExecDelete_1), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1796), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L18
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v147
	F_errmsg_internal(m, int32(_a_F_ExecDelete_4), v21+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1816), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L18
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L18
	} else {
		goto L75
	}
L75:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L18
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_ExecDelete_5), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1828), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L18
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v99
	F_errmsg_internal(m, int32(_a_F_ExecDelete_6), v21)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1834), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_ExecDelete_7), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1738), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L18
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+68))
	v256 = m.T0[v255].(func(*base.Module, int32, int32, int32, int32) int32)(m, v24, l1, v251, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L18
	} else {
		goto L87
	}
L87:
	;
	if v256 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v418 = int32(0)
	goto L8
L89:
	;
	goto L90
L90:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+4)))
	if v261&int32(2) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	F_ExecStoreAllNullTuple(m, v256)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L18
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v256)+40)) = v266
	v279 = v256
	goto L9
L94:
	;
	goto L93
L95:
	;
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v24)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+112)) = v286 + int64(1)
	goto L97
L96:
	;
	goto L97
L97:
	;
	if l8 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v290)
	goto L100
L99:
	;
	goto L100
L100:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+204))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)+104))
	if v295 != int32(2) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	F_ExecARDeleteTriggers(m, v292, l1, l2, l3, v312, l5)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L18
	} else {
		goto L112
	}
L102:
	;
	v312 = v294
	goto L101
L103:
	;
	goto L104
L104:
	;
	if v294 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v312 = int32(0)
	goto L101
L106:
	;
	goto L107
L107:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+1)))
	if v301 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v312 = v294
	goto L101
L109:
	;
	goto L110
L110:
	;
	v304 = int32(0)
	F_ExecARUpdateTriggers(m, v292, l1, v304, v304, l2, l3, v304, v304, v294, v304)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L18
	} else {
		goto L111
	}
L111:
	;
	v312 = v304
	goto L101
L112:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if l5 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v332 = int32(0)
	if base.B2i32(v330 == v332)|base.B2i32((l4|v331)&int32(1) == v332) != 0 {
		v418 = v332
		goto L8
	} else {
		goto L120
	}
L114:
	;
	v318 = int32(0)
	v330 = base.B2i32(v315 != v318)
	v331 = v318
	goto L113
L115:
	;
	goto L116
L116:
	;
	if v315 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v418 = int32(0)
	goto L8
L118:
	;
	goto L119
L119:
	;
	v324 = int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+12)))
	v330 = v324
	v331 = int32(base.Ui32(v325&int32(2)) >> (uint(v324) % 32))
	goto L113
L120:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v341 != 0 {
		v375 = v279
		goto L121
	} else {
		goto L122
	}
L121:
	;
	if v331 != 0 {
		goto L136
	} else {
		goto L137
	}
L122:
	;
	v342 = F_ExecGetReturningSlot(m, v24, l1)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L18
	} else {
		goto L123
	}
L123:
	;
	if l3 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	F_ExecForceStoreHeapTuple(m, l3, v342, int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L18
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ExecDelete[1]))
	if v348 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v375 = v342
	goto L121
L128:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+60))
	v371 = m.T0[v370].(func(*base.Module, int32, int32, int32, int32) int32)(m, v23, l2, int32(_a_F_ExecDelete_8), v342)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L18
	} else {
		goto L134
	}
L129:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecDelete[2])))
	if v352&int32(1) != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L18
	} else {
		goto L131
	}
L131:
	;
	F_errmsg_internal(m, int32(_a_F_ExecDelete_9), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_10), int32(1355), int32(_a_F_ExecDelete_11))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L18
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
	if v371 == int32(0) {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	v375 = v342
	goto L121
L136:
	;
	v376 = F_ExecGetChildToRootMap(m, l1)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L18
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v398 = F_ExecProcessReturning(m, l0, l1, int32(1), v375, int32(0), v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L18
	} else {
		goto L145
	}
L139:
	;
	if v376 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v375
	v418 = v332
	goto L8
L141:
	;
	goto L142
L142:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v376)+8))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+120))
	v384 = F_ExecGetReturningSlot(m, v24, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L18
	} else {
		goto L143
	}
L143:
	;
	v386 = F_execute_attr_map_slot(m, v381, v375, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v375)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+40)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v375)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+32)) = v390
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v386)+36)) = uint16(v392)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v386
	v418 = v332
	goto L8
L145:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+28))
	m.T0[v401].(func(*base.Module, int32))(m, v398)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L18
	} else {
		goto L146
	}
L146:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+12))
	m.T0[v405].(func(*base.Module, int32))(m, v375)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	v418 = v398
	goto L8
L148:
	;
	F_errmsg_internal(m, int32(_a_F_ExecDelete_12), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(_a_F_ExecDelete_2), int32(1891), int32(_a_F_ExecDelete_3))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecEndAppend(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if int32(0) < v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v9 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9<<(uint(int32(2))%32))))
	F_ExecEndNode(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v19 = v9 + int32(1)
	if v19 != v4 {
		v9 = v19
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
func F_ExecEndNode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
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
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v289 int64
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v337 int64
	_ = v337
	var v338 int64
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v429 int64
	_ = v429
	var v430 int64
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int64
	_ = v441
	var v442 int64
	_ = v442
	var v445 int64
	_ = v445
	var v446 int64
	_ = v446
	var v449 int64
	_ = v449
	var v450 int64
	_ = v450
	var v453 int64
	_ = v453
	var v454 int64
	_ = v454
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v615 int64
	_ = v615
	var v618 int64
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int64
	_ = v625
	var v627 int64
	_ = v627
	var v629 int64
	_ = v629
	var v631 int64
	_ = v631
	var v633 int64
	_ = v633
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_bms_free(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v21 - int32(400) {
	case 0:
		goto L9
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	case 5:
		goto L44
	case 6:
		goto L43
	case 7:
		goto L42
	default:
		goto L10
	case 9:
		goto L41
	case 10:
		goto L40
	case 11:
		goto L37
	case 12:
		goto L36
	case 13:
		goto L35
	case 14:
		goto L34
	case 15:
		goto L33
	case 16:
		goto L32
	case 17:
		goto L31
	case 18:
		goto L30
	case 19, 22, 23:
		goto L1
	case 20:
		goto L29
	case 21:
		goto L28
	case 24:
		goto L27
	case 25:
		goto L26
	case 27:
		goto L25
	case 28:
		goto L24
	case 29:
		goto L23
	case 30:
		goto L22
	case 31:
		goto L19
	case 32:
		goto L21
	case 33:
		goto L20
	case 34:
		goto L18
	case 35:
		goto L17
	case 36:
		goto L16
	case 37:
		goto L15
	case 38:
		goto L39
	case 39:
		goto L38
	case 40:
		goto L14
	case 41:
		goto L13
	case 42:
		goto L12
	case 43:
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = int32(0)
	goto L7
L9:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L3
	} else {
		goto L382
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L3
	} else {
		goto L379
	}
L11:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L3
	} else {
		goto L378
	}
L12:
	;
	F_EvalPlanQualEnd(m, l0+int32(108))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L3
	} else {
		goto L376
	}
L13:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v815 != 0 {
		goto L370
	} else {
		goto L371
	}
L14:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L3
	} else {
		goto L369
	}
L15:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L3
	} else {
		goto L368
	}
L16:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v757 != 0 {
		goto L348
	} else {
		goto L349
	}
L17:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v646 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L18:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L3
	} else {
		goto L298
	}
L19:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v608 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L20:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	F_ExecDropSingleTupleTableSlot(m, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L3
	} else {
		goto L279
	}
L21:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v581 != 0 {
		goto L274
	} else {
		goto L275
	}
L22:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v573 != 0 {
		goto L269
	} else {
		goto L270
	}
L23:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v552 != 0 {
		goto L255
	} else {
		goto L256
	}
L24:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L3
	} else {
		goto L253
	}
L25:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L3
	} else {
		goto L251
	}
L26:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v535)+12))
	m.T0[v536].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L250
	}
L27:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+80))
	if v519 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L28:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if l0 == v511 {
		goto L236
	} else {
		goto L237
	}
L29:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v506 != 0 {
		goto L232
	} else {
		goto L233
	}
L30:
	;
	v474 = int32(0)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v474 < v475 {
		goto L222
	} else {
		goto L223
	}
L31:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_ExecEndNode(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L3
	} else {
		goto L221
	}
L32:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v408 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L33:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v401 != 0 {
		goto L202
	} else {
		goto L203
	}
L34:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	if v323 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L35:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v301 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L36:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v272 != 0 {
		goto L152
	} else {
		goto L153
	}
L37:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v248 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L38:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L3
	} else {
		goto L128
	}
L39:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L115
	}
L40:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+32))
	if v205 != 0 {
		goto L107
	} else {
		goto L108
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v141 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L42:
	;
	F_ExecEndBitmapAnd(m, l0)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L91
	}
L43:
	;
	F_ExecEndBitmapAnd(m, l0)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L90
	}
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	F_tuplestore_end(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L3
	} else {
		goto L78
	}
L45:
	;
	F_ExecEndAppend(m, l0)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L77
	}
L46:
	;
	F_ExecEndAppend(m, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L76
	}
L47:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if int32(0) < v26 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	goto L1
L50:
	;
	v31 = v2
	goto L53
L51:
	;
	goto L52
L52:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v96 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L53:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v39 = v36 + v31*int32(216)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+92)))
	if v40 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if int32(0) < v51 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+84))
	if v41 == int32(0) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+72))
	if v44 == int32(0) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	m.T0[v44].(func(*base.Module, int32, int32))(m, v47, v39)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v58 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v86 = v31 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v86 < v87 {
		v31 = v86
		goto L53
	} else {
		goto L68
	}
L63:
	;
	v63 = v58 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v39)+108))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63+v64)))
	F_ExecDropSingleTupleTableSlot(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L65
	}
L64:
	;
	goto L62
L65:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v39)+112))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v63)))
	F_ExecDropSingleTupleTableSlot(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	v75 = v58 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if v75 < v76 {
		v58 = v75
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L54
L69:
	;
	F_EvalPlanQualEnd(m, l0+int32(124))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L74
	}
L70:
	;
	F_ExecCleanupTupleRouting(m, l0, v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v101 == int32(0) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	F_ExecDropSingleTupleTableSlot(m, v101)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	goto L1
L76:
	;
	goto L1
L77:
	;
	goto L1
L78:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_tuplestore_end(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v124 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_MemoryContextDelete(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v127 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	F_MemoryContextDelete(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L3
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	goto L1
L90:
	;
	goto L1
L91:
	;
	goto L1
L92:
	;
	goto L1
L93:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+188))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	m.T0[v198].(func(*base.Module, int32))(m, v140)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L106
	}
L94:
	;
	if v140 == int32(0) {
		goto L92
	} else {
		goto L105
	}
L95:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v145 < int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if v140 == int32(0) {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v140)+36))
	if v150 == int32(0) {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v157 = v141 + v145*int32(56) + int32(8)
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v150)))
	*(*int64)(unsafe.Add(mBase, uint32(v157))) = v158 + v159
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v157)+8))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+8)) = v162 + v163
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+16)))
	v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+16)))
	if v167 < v166 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+16)) = uint16(v166)
	goto L101
L100:
	;
	goto L101
L101:
	;
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v150)+18)))
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+18)))
	if v171 < v170 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+18)) = uint16(v170)
	goto L104
L103:
	;
	goto L104
L104:
	;
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v150)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+24)) = v174 + v175
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v157)+32))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v150)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+32)) = v178 + v179
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v157)+40))
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v150)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+40)) = v182 + v183
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v157)+48))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v150)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v157)+48)) = v186 + v187
	goto L93
L105:
	;
	goto L93
L106:
	;
	goto L92
L107:
	;
	m.T0[v205].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v208 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+188))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	m.T0[v211].(func(*base.Module, int32))(m, v208)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L1
L114:
	;
	goto L113
L115:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v217 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	F_ExecParallelFinish(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v220 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	goto L118
L120:
	;
	F_pfree(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L3
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v225 != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L122
L124:
	;
	F_ExecParallelCleanup(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	goto L1
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L126
L128:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v233 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_ExecParallelFinish(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v236 != 0 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L131
L133:
	;
	F_pfree(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L3
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(0)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v241 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	goto L135
L137:
	;
	F_ExecParallelCleanup(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	goto L1
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	goto L139
L141:
	;
	if v246 != 0 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v252 < int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v257 = v248 + v252<<(uint(int32(3))%32)
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v257)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v259)))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+8)) = v258 + v260
	goto L141
L144:
	;
	F_index_endscan(m, v246)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L3
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	if v247 != 0 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L146
L148:
	;
	F_relation_close(m, v247, int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	goto L1
L151:
	;
	goto L150
L152:
	;
	F_ReleaseBuffer(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L3
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v277 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	goto L154
L156:
	;
	if v270 != 0 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v281 < int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v286 = v277 + v281<<(uint(int32(3))%32)
	v287 = *(*int64)(unsafe.Add(mBase, uint32(v286)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v289 = *(*int64)(unsafe.Add(mBase, uint32(v288)))
	*(*int64)(unsafe.Add(mBase, uint32(v286)+8)) = v287 + v289
	goto L156
L159:
	;
	F_index_endscan(m, v270)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L3
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	if v271 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	goto L161
L163:
	;
	F_relation_close(m, v271, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	goto L1
L166:
	;
	goto L165
L167:
	;
	if v299 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v305 < int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v310 = v301 + v305<<(uint(int32(3))%32)
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v310)+8))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v311 + v313
	goto L167
L170:
	;
	F_index_endscan(m, v299)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L3
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v300 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	goto L172
L174:
	;
	F_relation_close(m, v300, int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L3
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	goto L1
L177:
	;
	goto L176
L178:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L189
	}
L179:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v327 < int32(0) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	v332 = v323 + v327*int32(72)
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v332)+8))
	v334 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v332)+8)) = v333 + v334
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v332)+16))
	v338 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v332)+16)) = v337 + v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v341 == int32(0) {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v341)+36))
	if v344 == int32(0) {
		goto L178
	} else {
		goto L182
	}
L182:
	;
	v348 = v332 + int32(8)
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v348)+16))
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v344)))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+16)) = v349 + v350
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v348)+24))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v344)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+24)) = v353 + v354
	v357 = int32(*(*int16)(unsafe.Add(mBase, uint32(v344)+16)))
	v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v348)+32)))
	if v358 < v357 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v348)+32)) = uint16(v357)
	goto L185
L184:
	;
	goto L185
L185:
	;
	v361 = int32(*(*int16)(unsafe.Add(mBase, uint32(v344)+18)))
	v362 = int32(*(*int16)(unsafe.Add(mBase, uint32(v348)+34)))
	if v362 < v361 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v348)+34)) = uint16(v361)
	goto L188
L187:
	;
	goto L188
L188:
	;
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v348)+40))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v344)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+40)) = v365 + v366
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v348)+48))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v344)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+48)) = v369 + v370
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v348)+56))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v344)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+56)) = v373 + v374
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v348)+64))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v344)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+64)) = v377 + v378
	goto L178
L189:
	;
	if v384 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384)+20))
	if v388 != 0 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	goto L192
L192:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v398 != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	F_tbm_end_iterate(m, v384+int32(16))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L3
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+188))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	m.T0[v395].(func(*base.Module, int32))(m, v384)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L3
	} else {
		goto L197
	}
L196:
	;
	goto L195
L197:
	;
	goto L192
L198:
	;
	F_tbm_free(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L3
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	goto L1
L201:
	;
	goto L200
L202:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)+188))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	m.T0[v404].(func(*base.Module, int32))(m, v401)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L3
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	goto L1
L205:
	;
	goto L204
L206:
	;
	goto L1
L207:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+188))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+12))
	m.T0[v465].(func(*base.Module, int32))(m, v407)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L220
	}
L208:
	;
	if v407 == int32(0) {
		goto L206
	} else {
		goto L219
	}
L209:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v412 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	if v407 == int32(0) {
		goto L206
	} else {
		goto L211
	}
L211:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v407)+36))
	if v417 == int32(0) {
		goto L207
	} else {
		goto L212
	}
L212:
	;
	v424 = v408 + v412*int32(56) + int32(8)
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v424)))
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v417)))
	*(*int64)(unsafe.Add(mBase, uint32(v424))) = v425 + v426
	v429 = *(*int64)(unsafe.Add(mBase, uint32(v424)+8))
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v417)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+8)) = v429 + v430
	v433 = int32(*(*int16)(unsafe.Add(mBase, uint32(v417)+16)))
	v434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v424)+16)))
	if v434 < v433 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v424)+16)) = uint16(v433)
	goto L215
L214:
	;
	goto L215
L215:
	;
	v437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v417)+18)))
	v438 = int32(*(*int16)(unsafe.Add(mBase, uint32(v424)+18)))
	if v438 < v437 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v424)+18)) = uint16(v437)
	goto L218
L217:
	;
	goto L218
L218:
	;
	v441 = *(*int64)(unsafe.Add(mBase, uint32(v424)+24))
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v417)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+24)) = v441 + v442
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v424)+32))
	v446 = *(*int64)(unsafe.Add(mBase, uint32(v417)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+32)) = v445 + v446
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v424)+40))
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v417)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+40)) = v449 + v450
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v424)+48))
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v417)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v424)+48)) = v453 + v454
	goto L207
L219:
	;
	goto L207
L220:
	;
	goto L206
L221:
	;
	goto L1
L222:
	;
	v479 = v474
	v481 = v475
	goto L225
L223:
	;
	goto L224
L224:
	;
	goto L1
L225:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v488 = v485 + v479<<(uint(int32(5))%32)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+12))
	if v489 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L224
L227:
	;
	F_tuplestore_end(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L230
	}
L228:
	;
	v495 = v481
	goto L229
L229:
	;
	v497 = v479 + int32(1)
	if v497 < v495 {
		v479 = v497
		v481 = v495
		goto L225
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488)+12)) = int32(0)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v495 = v494
	goto L229
L231:
	;
	goto L226
L232:
	;
	F_tuplestore_end(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	goto L1
L235:
	;
	goto L234
L236:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_tuplestore_end(m, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L3
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	goto L1
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	goto L238
L240:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v532 != 0 {
		goto L246
	} else {
		goto L247
	}
L241:
	;
	v526 = int32(28)
	goto L243
L242:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+156))
	if v524 != 0 {
		goto L240
	} else {
		goto L244
	}
L243:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526+v527)))
	m.T0[v529].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L245
	}
L244:
	;
	v526 = int32(100)
	goto L243
L245:
	;
	goto L240
L246:
	;
	F_ExecEndNode(m, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L3
	} else {
		goto L249
	}
L247:
	;
	goto L248
L248:
	;
	goto L1
L249:
	;
	goto L248
L250:
	;
	goto L1
L251:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v542)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L3
	} else {
		goto L252
	}
L252:
	;
	goto L1
L253:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L3
	} else {
		goto L254
	}
L254:
	;
	goto L1
L255:
	;
	F_tuplestore_end(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L3
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v551)+120))
	if v557 != 0 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	goto L257
L259:
	;
	F_tuplestore_end(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L3
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v562 != 0 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v551)+120)) = int32(0)
	goto L261
L263:
	;
	F_ExecHashTableDestroy(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L3
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L3
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L265
L267:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v570)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L3
	} else {
		goto L268
	}
L268:
	;
	goto L1
L269:
	;
	F_tuplestore_end(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L3
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v578)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L3
	} else {
		goto L273
	}
L272:
	;
	goto L271
L273:
	;
	goto L1
L274:
	;
	F_tuplesort_end(m, v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L3
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L3
	} else {
		goto L278
	}
L277:
	;
	goto L276
L278:
	;
	goto L1
L279:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_ExecDropSingleTupleTableSlot(m, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L280
	}
L280:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v595 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	F_tuplesort_end(m, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L3
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v600 != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	goto L283
L285:
	;
	F_tuplesort_end(m, v600)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L3
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L3
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L287
L289:
	;
	goto L1
L290:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	F_MemoryContextDelete(m, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L3
	} else {
		goto L296
	}
L291:
	;
	v612 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v612 < int32(0) {
		goto L290
	} else {
		goto L292
	}
L292:
	;
	v615 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v615 == int64(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v618 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v618
	goto L295
L294:
	;
	goto L295
L295:
	;
	v622 = v608 + v612*int32(40)
	v624 = l0 + int32(200)
	v625 = *(*int64)(unsafe.Add(mBase, uint32(v624)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+40)) = v625
	v627 = *(*int64)(unsafe.Add(mBase, uint32(v624)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+32)) = v627
	v629 = *(*int64)(unsafe.Add(mBase, uint32(v624)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+24)) = v629
	v631 = *(*int64)(unsafe.Add(mBase, uint32(v624)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+16)) = v631
	v633 = *(*int64)(unsafe.Add(mBase, uint32(v624)))
	*(*int64)(unsafe.Add(mBase, uint32(v622)+8)) = v633
	goto L290
L296:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v640)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L3
	} else {
		goto L297
	}
L297:
	;
	goto L1
L298:
	;
	goto L1
L299:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v664 != 0 {
		goto L302
	} else {
		goto L303
	}
L300:
	;
	v650 = *(*int32)(unsafe.Add(mBase, _c_F_ExecEndNode[0]))
	if v650 < int32(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v655 = v646 + v650*int32(24)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+24)) = v656
	v658 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v655)+16)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v655)+8)) = v660
	goto L299
L302:
	;
	F_tuplesort_end(m, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L3
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v667 != 0 {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	goto L304
L306:
	;
	F_tuplesort_end(m, v667)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L3
	} else {
		goto L309
	}
L307:
	;
	goto L308
L308:
	;
	F_hashagg_reset_spill_state(m, l0)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L3
	} else {
		goto L310
	}
L309:
	;
	goto L308
L310:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	if v672 != 0 {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	F_MemoryContextDelete(m, v672)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L3
	} else {
		goto L314
	}
L312:
	;
	goto L313
L313:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v677 != 0 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = int32(0)
	goto L313
L315:
	;
	F_MemoryContextDelete(m, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L3
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v682 <= int32(0) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = int32(0)
	goto L317
L319:
	;
	v735 = int32(0)
	goto L339
L320:
	;
	v685 = int32(1)
	if v645 <= v685 {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	goto L322
L322:
	;
	v689 = int32(1)
	if v645 <= v689 {
		goto L326
	} else {
		goto L327
	}
L323:
	;
	v688 = v685
	goto L325
L324:
	;
	v688 = v645
	goto L325
L325:
	;
	v728 = v688
	goto L319
L326:
	;
	v692 = v689
	goto L328
L327:
	;
	v692 = v645
	goto L328
L328:
	;
	v696 = v2
	goto L329
L329:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v706 = int32(0)
	goto L331
L330:
	;
	v728 = v692
	goto L319
L331:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v700+v696*int32(240))+220))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v712+v706<<(uint(int32(2))%32))))
	if v716 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v723 = v696 + int32(1)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v723 < v724 {
		v696 = v723
		goto L329
	} else {
		goto L338
	}
L333:
	;
	F_tuplesort_end(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L3
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v720 = v706 + int32(1)
	if v720 != v692 {
		v706 = v720
		goto L331
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	goto L332
L338:
	;
	goto L330
L339:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v741+v735<<(uint(int32(2))%32))))
	F_ReScanExprContext(m, v745)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L3
	} else {
		goto L341
	}
L340:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v751 != 0 {
		goto L343
	} else {
		goto L344
	}
L341:
	;
	v749 = v735 + int32(1)
	if v749 != v728 {
		v735 = v749
		goto L339
	} else {
		goto L342
	}
L342:
	;
	goto L340
L343:
	;
	F_ReScanExprContext(m, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L3
	} else {
		goto L347
	}
L346:
	;
	goto L345
L347:
	;
	goto L1
L348:
	;
	F_tuplestore_end(m, v757)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L3
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	F_release_partition(m, l0)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L3
	} else {
		goto L352
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L350
L352:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v764 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v769 = int32(0)
	v771 = v764
	goto L356
L354:
	;
	goto L355
L355:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	F_MemoryContextDelete(m, v796)
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L3
	} else {
		goto L363
	}
L356:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v775+v769*int32(184))+144))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	if v779 != v780 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L355
L358:
	;
	F_MemoryContextDelete(m, v779)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L3
	} else {
		goto L361
	}
L359:
	;
	v785 = v771
	goto L360
L360:
	;
	v787 = v769 + int32(1)
	if v787 < v785 {
		v769 = v787
		v771 = v785
		goto L356
	} else {
		goto L362
	}
L361:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v785 = v784
	goto L360
L362:
	;
	goto L357
L363:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	F_MemoryContextDelete(m, v799)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L3
	} else {
		goto L364
	}
L364:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	F_pfree(m, v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L3
	} else {
		goto L365
	}
L365:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v805)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L3
	} else {
		goto L366
	}
L366:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v808)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L3
	} else {
		goto L367
	}
L367:
	;
	goto L1
L368:
	;
	goto L1
L369:
	;
	goto L1
L370:
	;
	F_MemoryContextDelete(m, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L3
	} else {
		goto L373
	}
L371:
	;
	goto L372
L372:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v818)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L3
	} else {
		goto L374
	}
L373:
	;
	goto L372
L374:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L3
	} else {
		goto L375
	}
L375:
	;
	goto L1
L376:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v828)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L3
	} else {
		goto L377
	}
L377:
	;
	goto L1
L378:
	;
	goto L1
L379:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v837
	F_errmsg_internal(m, int32(_a_F_ExecEndNode_0), v10)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L3
	} else {
		goto L380
	}
L380:
	;
	F_errfinish(m, int32(_a_F_ExecEndNode_1), int32(741), int32(_a_F_ExecEndNode_2))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L3
	} else {
		goto L381
	}
L381:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L382:
	;
	goto L1
}
func F_ExecGetChildToRootMap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
		return v6
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
		if v8 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
			v13 = F_convert_tuples_by_name(m, v10, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = v13
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v19)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v18
				return v18
			}
		} else {
			v18 = int32(0)
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v19)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v18
			return v18
		}
	}
}
func F_ExecGetInsertedCols(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v8 != 0 {
		v9 = v8
	} else {
		v9 = l0
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(0) {
		v82 = v3
		return v82
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+v10<<(uint(int32(2))%32)-int32(4))))
		v22 = F_getRTEPermissionInfo(m, v13, v21)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
				v82 = v3
				return v82
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v28 == int32(0) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
					v82 = v77
					return v82
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
					if v31 == int32(0) {
						v34 = int32(_a_F_ExecGetInsertedCols_0)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_ExecGetInsertedCols[0]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecGetInsertedCols[0])) = v41
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+131)))
						v49 = F_build_attrmap_by_name_if_req(m, v39, v37, (v44^int32(-1))&int32(1))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = F_convert_tuples_by_name_attrmap(m, v39, v37, v49)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v51
									*(*int32)(unsafe.Add(mBase, _c_F_ExecGetInsertedCols[0])) = v35
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									if v63 == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
										v82 = v77
										return v82
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
										v68 = F_execute_attr_map_cols(m, v66, v67)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											return v68
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ExecGetInsertedCols[0])) = v35
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v63 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
									v82 = v77
									return v82
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
									v68 = F_execute_attr_map_cols(m, v66, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										return v68
									}
								}
							}
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						if v63 == int32(0) {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
							v82 = v77
							return v82
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
							v68 = F_execute_attr_map_cols(m, v66, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v68
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecIRInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+36)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11)+28)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11)+20)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(85899346368)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v27 <= v4 {
		v105 = int32(1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v105
L2:
	;
	v35 = v4
	v37 = v4
	goto L3
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v41 = v38 + v37*int32(60)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
	if v42&int32(71) != int32(69) {
		v95 = v35
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v105 = v96
	goto L1
L5:
	;
	v96 = int32(1)
	v98 = v37 + v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v98 < v99 {
		v35 = v95
		v37 = v98
		goto L3
	} else {
		goto L32
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v48 = int32(0)
	v50 = F_TriggerEnabled(m, l0, l1, v41, v47, v48, v48, l2)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	if v50 == int32(0) {
		v95 = v35
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v61 = F_ExecFetchSlotHeapTuple(m, l2, int32(1), v11+int32(47))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	v63 = v35
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v69 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v63 = v61
	goto L12
L14:
	;
	v72 = v69
	goto L16
L15:
	;
	v70 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L17
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v74 = F_ExecCallTriggerFunc(m, v11, v37, v67, v68, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v72 = v70
	goto L16
L18:
	;
	if v74 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = int32(0)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
	if v79 != int32(1) {
		v105 = v78
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v74 == v63 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	F_pfree(m, v63)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v105 = v78
	goto L1
L24:
	;
	v95 = v74
	goto L5
L25:
	;
	goto L26
L26:
	;
	F_ExecForceStoreHeapTuple(m, v74, l2, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+47)))
	if v88 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_pfree(m, v63)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v95 = int32(0)
	goto L5
L31:
	;
	goto L30
L32:
	;
	goto L4
}
func F_ExecIRUpdateTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v16 = F_ExecGetTriggerOldSlot(m, l0, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v13)+20)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(94489280960)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v30
	F_ExecForceStoreHeapTuple(m, l2, v16, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v36 <= int32(0) {
		v111 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v13 + int32(48)
	return v111
L5:
	;
	v45 = v5
	v48 = v5
	goto L6
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v52 = v49 + v48*int32(60)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+12)))
	if v53&int32(83) != int32(81) {
		v100 = v45
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v111 = v101
	goto L4
L8:
	;
	v101 = int32(1)
	v103 = v48 + v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v103 < v104 {
		v45 = v100
		v48 = v103
		goto L6
	} else {
		goto L32
	}
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v60 = F_TriggerEnabled(m, l0, l1, v52, v58, int32(0), v16, l3)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v60 == int32(0) {
		v100 = v45
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v45 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v69 = F_ExecFetchSlotHeapTuple(m, l3, int32(1), v13+int32(47))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v71 = v45
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v71
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v79 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v71 = v69
	goto L14
L16:
	;
	v82 = v79
	goto L18
L17:
	;
	v80 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v84 = F_ExecCallTriggerFunc(m, v13, v48, v77, v78, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v82 = v80
	goto L18
L20:
	;
	if v84 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v111 = int32(0)
	goto L4
L22:
	;
	goto L23
L23:
	;
	if v84 == v71 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v100 = v84
	goto L8
L25:
	;
	goto L26
L26:
	;
	F_ExecForceStoreHeapTuple(m, v84, l3, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)))
	if v93 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	F_pfree(m, v71)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v100 = int32(0)
	goto L8
L31:
	;
	goto L30
L32:
	;
	goto L7
}
func F_ExecInterpExprStillValid(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v4 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v36 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v18 + int32(16)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v213
	v215 = m.T0[v213].(func(*base.Module, int32, int32, int32) int64)(m, l0, l1, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L50
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v46 = v43 + v36*int32(40)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v47&int32(64) != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	switch v64 - int32(7) {
	case 0:
		v67 = v27
		goto L13
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		goto L12
	}
L7:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v50
	v52 = int32(8)
	v58 = F_bsearch(m, v18+v52, int32(_a_F_ExecInterpExprStillValid_0), int32(120), v52, int32(634))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v64 = v63
	goto L6
L10:
	;
	return int64(0)
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v64 = v62
	goto L6
L12:
	;
	v192 = v36 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v192 < v193 {
		v36 = v192
		goto L4
	} else {
		goto L49
	}
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v70 = v68 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v72 = m.G0
	v74 = v72 + int32(-64)
	m.G0 = v74
	if int32(0) < v70 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	v67 = v23
	goto L13
L15:
	;
	v67 = v24
	goto L13
L16:
	;
	v67 = v25
	goto L13
L17:
	;
	v67 = v26
	goto L13
L18:
	;
	goto L12
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L41
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L10
	} else {
		goto L36
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L33
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L30
	}
L23:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 < v70 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	m.G0 = v74 - int32(-64)
	goto L18
L26:
	;
	v86 = v78 + v79<<(uint(int32(3))%32) + v70*int32(100)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+18)))
	if v87 == int32(118) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v91 = v86 - int32(72)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+91)))
	if v92 == int32(1) {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	if v71 != v95 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v70
	F_errmsg_internal(m, int32(_a_F_ExecInterpExprStillValid_1), v74)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_ExecInterpExprStillValid_2), int32(2414), int32(_a_F_ExecInterpExprStillValid_3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errmsg_internal(m, int32(_a_F_ExecInterpExprStillValid_4), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_ExecInterpExprStillValid_2), int32(2420), int32(_a_F_ExecInterpExprStillValid_3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v138 = F_format_type_be(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v70
	F_errmsg(m, int32(_a_F_ExecInterpExprStillValid_5), v72+int32(-48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_ExecInterpExprStillValid_2), int32(2426), int32(_a_F_ExecInterpExprStillValid_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v160 = F_format_type_be(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+52)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v74)+48)) = v70
	F_errmsg(m, int32(_a_F_ExecInterpExprStillValid_6), v72+int32(-16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	v170 = F_format_type_be(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v172 = F_format_type_be(m, v71)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+36)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v74)+32)) = v170
	v179 = F_errdetail(m, int32(_a_F_ExecInterpExprStillValid_7), v72+int32(-32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_ExecInterpExprStillValid_2), int32(2435), int32(_a_F_ExecInterpExprStillValid_3))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	goto L5
L50:
	;
	return v215
}
func F_ExecLimit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
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
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int64
	_ = v79
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int64
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int64
	_ = v199
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v224 int64
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLimit[0]))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	switch v24 {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		v290 = v2
		goto L10
	case 3:
		goto L19
	case 4:
		goto L18
	case 5:
		goto L15
	case 6:
		goto L14
	case 7:
		goto L13
	default:
		goto L12
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L138
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L135
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L132
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L129
	}
L10:
	;
	m.G0 = v12 + int32(16)
	return v290
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(6)
	v290 = v2
	goto L10
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L126
	}
L13:
	;
	if v23 != int32(1) {
		v290 = v2
		goto L10
	} else {
		goto L125
	}
L14:
	;
	if v23 == int32(1) {
		v290 = v2
		goto L10
	} else {
		goto L114
	}
L15:
	;
	if v23 == int32(1) {
		v290 = v2
		goto L10
	} else {
		goto L106
	}
L16:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v203 <= v204+int64(1) {
		goto L96
	} else {
		goto L97
	}
L17:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v158 != 0 {
		goto L78
	} else {
		goto L79
	}
L18:
	;
	if v23 != int32(1) {
		goto L16
	} else {
		goto L77
	}
L19:
	;
	if v23 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L20:
	;
	if v23 != int32(1) {
		v290 = v2
		goto L10
	} else {
		goto L23
	}
L21:
	;
	F_recompute_limits(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if int64(0) < v29 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	goto L27
L25:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)))
	if v32 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(2)
	v290 = v2
	goto L10
L27:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v44 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v290 = v48
	goto L10
L29:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v59 != int32(1) {
		v75 = v57
		v76 = v58
		goto L39
	} else {
		goto L40
	}
L34:
	;
	if v48 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
	if v50&int32(2) == int32(0) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(2)
	v290 = v2
	goto L10
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v48
	v79 = v76 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v79
	if v79 <= v75 {
		goto L27
	} else {
		goto L43
	}
L40:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v62-int64(1) != v58-v57 {
		v75 = v57
		v76 = v58
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	m.T0[v69].(func(*base.Module, int32, int32))(m, v67, v48)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v75 = v72
	v76 = v73
	goto L39
L43:
	;
	goto L28
L44:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+136)))
	if v86 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v133 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v133 <= v134+int64(1) {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v99 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v88-v89 < v87 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v92 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(6)
	v290 = v2
	goto L10
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(4)
	goto L17
L53:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v103 = m.T0[v102].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v112 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	if v103 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
	if v105&int32(2) == int32(0) {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(5)
	v290 = v2
	goto L10
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v103
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v129 + int64(1)
	v290 = v103
	goto L10
L64:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v115-int64(1) != v118-v119 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+32))
	m.T0[v124].(func(*base.Module, int32, int32))(m, v122, v103)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(7)
	v290 = v2
	goto L10
L68:
	;
	goto L69
L69:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v140 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v144 = m.T0[v143].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	if v144 == int32(0) {
		goto L9
	} else {
		goto L75
	}
L75:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+4)))
	if v148&int32(2) != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v144
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v152 - int64(1)
	v290 = v144
	goto L10
L77:
	;
	goto L17
L78:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v162 = m.T0[v161].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v162
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v174 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	if v162 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
	if v164&int32(2) == int32(0) {
		goto L82
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(5)
	v290 = v2
	goto L10
L87:
	;
	goto L86
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v162
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v199 + int64(1)
	v290 = v162
	goto L10
L89:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_MemoryContextReset(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v180 = int32(_a_F_ExecLimit_0)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_ExecLimit[1]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLimit[1])) = v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	v188 = m.T0[v187].(func(*base.Module, int32, int32, int32) int64)(m, v174, v14, v12+int32(15))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecLimit[1])) = v181
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_MemoryContextReset(m, v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	if v188 == int64(0) {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	goto L88
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(7)
	v290 = v2
	goto L10
L97:
	;
	goto L98
L98:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v210 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v214 = m.T0[v213].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	if v214 == int32(0) {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+4)))
	if v218&int32(2) != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v224 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v224 - int64(1)
	v290 = v214
	goto L10
L106:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v230 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v234 = m.T0[v233].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	if v234 == int32(0) {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+4)))
	if v238&int32(2) != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v234
	v290 = v234
	goto L10
L114:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v246 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v249 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v290 = v265
	goto L10
L118:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v253 = m.T0[v252].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v253 == int32(0) {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+4)))
	if v257&int32(2) != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v253
	v290 = v253
	goto L10
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v290 = v270
	goto L10
L126:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v275
	F_errmsg_internal(m, int32(_a_F_ExecLimit_1), v12)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_ExecLimit_2), int32(336), int32(_a_F_ExecLimit_3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLimit_4), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ExecLimit_2), int32(211), int32(_a_F_ExecLimit_3))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLimit_4), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_ExecLimit_2), int32(269), int32(_a_F_ExecLimit_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
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
	F_errmsg_internal(m, int32(_a_F_ExecLimit_4), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ExecLimit_2), int32(286), int32(_a_F_ExecLimit_3))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	F_errmsg_internal(m, int32(_a_F_ExecLimit_4), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_ExecLimit_2), int32(305), int32(_a_F_ExecLimit_3))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecMarkPos(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
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
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 - int32(400) {
	case 0:
		v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v195 != 0 {
			F_ExecMarkPos(m, v195)
			mBase = m.M
			v197 = m.ExcPending
			if v197 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			v200 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v201 = m.ExcPending
			if v201 != 0 {
				return
			} else {
				if v200 != 0 {
					F_errmsg_internal(m, int32(_a_F_ExecMarkPos_0), int32(0))
					mBase = m.M
					v205 = m.ExcPending
					if v205 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecMarkPos_1), int32(153), int32(_a_F_ExecMarkPos_2))
						mBase = m.M
						v210 = m.ExcPending
						if v210 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	default:
		v213 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v214 = m.ExcPending
		if v214 != 0 {
			return
		} else {
			if v213 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v217
				F_errmsg_internal(m, int32(_a_F_ExecMarkPos_3), v8)
				mBase = m.M
				v221 = m.ExcPending
				if v221 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ExecMarkPos_4), int32(358), int32(_a_F_ExecMarkPos_5))
					mBase = m.M
					v226 = m.ExcPending
					if v226 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	case 11:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
		if v14 == int32(0) {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			F_index_markpos(m, v51)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
			v20 = v18 - int32(1)
			v22 = v20 << (uint(int32(2)) % 32)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22+v23)))
			if v25 == int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22)))
				if v30 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					F_index_markpos(m, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v20))))
					if v35 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ExecMarkPos_6), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecMarkPos_7), int32(864), int32(_a_F_ExecMarkPos_8))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
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
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v20))))
				if v35 != 0 {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecMarkPos_6), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecMarkPos_7), int32(864), int32(_a_F_ExecMarkPos_8))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
	case 12:
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+156))
		if v57 == int32(0) {
			v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
			F_index_markpos(m, v94)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
			v63 = v61 - int32(1)
			v65 = v63 << (uint(int32(2)) % 32)
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v66)))
			if v68 == int32(0) {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v71+v65)))
				if v73 == int32(0) {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					F_index_markpos(m, v94)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v63))))
					if v78 != 0 {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F_ExecMarkPos_9), int32(0))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecMarkPos_10), int32(487), int32(_a_F_ExecMarkPos_11))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
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
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v63))))
				if v78 != 0 {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_ExecMarkPos_9), int32(0))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecMarkPos_10), int32(487), int32(_a_F_ExecMarkPos_11))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
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
	case 25:
		v99 = m.G0
		v101 = v99 - int32(16)
		m.G0 = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
		if v104 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return
				} else {
					v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
					*(*int32)(unsafe.Add(mBase, uint32(v101))) = v115
					F_errmsg(m, int32(_a_F_ExecMarkPos_12), v101)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecMarkPos_13), int32(145), int32(_a_F_ExecMarkPos_14))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.T0[v104].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return
			} else {
				m.G0 = v101 + int32(16)
				m.G0 = v8 + int32(16)
				return
			}
		}
	case 30:
		v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		if v130 != 0 {
			F_tuplestore_copy_read_pointer(m, v130, int32(0), int32(1))
			mBase = m.M
			v134 = m.ExcPending
			if v134 != 0 {
				return
			} else {
				v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				F_tuplestore_trim(m, v135)
				mBase = m.M
				v137 = m.ExcPending
				if v137 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	case 32:
		v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
		if v138 == int32(1) {
			v141 = int32(_a_F_ExecMarkPos_15)
			v142 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMarkPos[0]))
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+28))
			*(*int32)(unsafe.Add(mBase, _c_F_ExecMarkPos[0])) = v145
			v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)+64))
			switch v147 - int32(3) {
			case 0:
				v185 = *(*int32)(unsafe.Add(mBase, uint32(v144)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v144)+224)) = v185
				v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+208)))
				*(*uint8)(unsafe.Add(mBase, uint32(v144)+228)) = uint8(v188)
				*(*int32)(unsafe.Add(mBase, _c_F_ExecMarkPos[0])) = v142
				m.G0 = v8 + int32(16)
				return
			case 1:
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+200))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+40))
				if v151 == int32(0) {
					v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+44))
					v155 = F_palloc(m, v154)
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v150)+40)) = v155
						*(*int64)(unsafe.Add(mBase, uint32(v150)+52)) = int64(0)
						v160 = *(*int64)(unsafe.Add(mBase, uint32(v150)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v150)+24)) = v160
						v162 = F_ltsReadFillBuffer(m, v150)
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return
						} else {
							v166 = *(*int64)(unsafe.Add(mBase, uint32(v150)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v144+int32(216)))) = v166
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v150)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v144+int32(224)))) = v170
							v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+208)))
							*(*uint8)(unsafe.Add(mBase, uint32(v144)+228)) = uint8(v188)
							*(*int32)(unsafe.Add(mBase, _c_F_ExecMarkPos[0])) = v142
							m.G0 = v8 + int32(16)
							return
						}
					}
				} else {
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v150)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v144+int32(216)))) = v166
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v150)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v144+int32(224)))) = v170
					v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+208)))
					*(*uint8)(unsafe.Add(mBase, uint32(v144)+228)) = uint8(v188)
					*(*int32)(unsafe.Add(mBase, _c_F_ExecMarkPos[0])) = v142
					m.G0 = v8 + int32(16)
					return
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v175 = m.ExcPending
				if v175 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ExecMarkPos_16), int32(0))
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecMarkPos_17), int32(2351), int32(_a_F_ExecMarkPos_18))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_ExecMergeNotMatched(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v130 int32
	_ = v130
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+172))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v4
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v23
	if v17 == v4 {
		v130 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v130
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v29 <= int32(0) {
		v130 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = v4
	goto L5
L4:
	;
	switch v50 - int32(3) {
	case 0:
		goto L12
	default:
		goto L13
	case 4:
		v130 = int32(0)
		goto L1
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v33<<(uint(int32(2))%32))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v51 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	v130 = int32(0)
	goto L1
L7:
	;
	v54 = int32(_a_F_ExecMergeNotMatched_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
	v62 = m.T0[v61].(func(*base.Module, int32, int32, int32) int64)(m, v51, v20, v15+int32(15))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0])) = v55
	if v62 != int64(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v71 = v33 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v71 < v72 {
		v33 = v71
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+80))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+24))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	m.T0[v96].(func(*base.Module, int32))(m, v94)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L17
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errmsg_internal(m, int32(_a_F_ExecMergeNotMatched_1), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_ExecMergeNotMatched_2), int32(3889), int32(_a_F_ExecMergeNotMatched_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v99 = int32(_a_F_ExecMergeNotMatched_0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0])) = v102
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)+32))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32) int64)(m, v92+int32(8), v93, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecMergeNotMatched[0])) = v100
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v114 = v112 & int32(_a_F_ExecMergeNotMatched_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)) = uint16(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)) = uint16(v117)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+216)) = v48
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v19)+120))
	v121 = int32(0)
	v123 = F_ExecInsert(m, l0, v120, v94, l2, v121, v121)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v19)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+224)) = base.F64_add(v125, float64(1))
	v130 = v123
	goto L1
}
func F_ExecProjectSet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ExecProjectSet[0]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	F_MemoryContextReset(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v15 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v51
L8:
	;
	v19 = F_ExecProjectSRF(m, l0, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L13
L11:
	;
	if v19 != 0 {
		v51 = v19
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	F_MemoryContextReset(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_ExecReScan(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v33 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v35 = m.T0[v34].(func(*base.Module, int32) int32)(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	if v35 == int32(0) {
		v51 = v33
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)))
	if v39&int32(2) != 0 {
		v51 = v33
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v35
	v44 = F_ExecProjectSRF(m, l0, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v44 != 0 {
		v51 = v44
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	F_MemoryContextReset(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	goto L13
}
func F_ExecRecursiveUnion(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRecursiveUnion[0]))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	if v20 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v9 + int32(16)
	return v120
L7:
	;
	goto L11
L8:
	;
	goto L9
L9:
	;
	goto L27
L10:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)) = uint8(v58)
	goto L9
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	F_tuplestore_puttupleslot(m, v55, v33)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L26
	}
L13:
	;
	F_ExecReScan(m, v13)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v33 = m.T0[v32].(func(*base.Module, int32) int32)(m, v13)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if v33 == int32(0) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	if v37&int32(2) != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if int32(0) < v40 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v47 = F_LookupTupleHashEntry(m, v43, v33, v9+int32(15), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	goto L12
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v52 != int32(1) {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v120 = v33
	goto L6
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	if v72 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	F_tuplestore_puttupleslot(m, v116, v76)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L50
	}
L29:
	;
	F_ExecReScan(m, v12)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v76 = m.T0[v75].(func(*base.Module, int32) int32)(m, v12)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	if int32(0) < v99 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if v76 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	if v78&int32(2) == int32(0) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)))
	if v83 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	v120 = int32(0)
	goto L6
L40:
	;
	goto L41
L41:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	F_tuplestore_clear(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+105)) = uint8(v88)
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+108)) = base.I64_rotl(v90, int64(32))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v12)+52))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	v96 = F_bms_add_member(m, v94, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v96
	goto L27
L44:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v106 = F_LookupTupleHashEntry(m, v102, v76, v9+int32(15), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L28
L47:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_MemoryContextReset(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v111 != int32(1) {
		goto L27
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v120 = v76
	goto L6
}
func F_ExecUnique(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUnique[0]))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ExecReScan(m, v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v36 = m.T0[v35].(func(*base.Module, int32) int32)(m, v22)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_MemoryContextReset(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L29
	}
L13:
	;
	m.G0 = v12 + int32(16)
	return v81
L14:
	;
	if v21 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v38&int32(2) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	m.T0[v44].(func(*base.Module, int32))(m, v21)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v81 = int32(0)
	goto L13
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+32))
	m.T0[v78].(func(*base.Module, int32, int32))(m, v21, v36)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L28
	}
L22:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v50&int32(2) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v36
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v55 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v58 = int32(_a_F_ExecUnique_0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUnique[1]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUnique[1])) = v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+24))
	v66 = m.T0[v65].(func(*base.Module, int32, int32, int32) int64)(m, v55, v14, v12+int32(15))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUnique[1])) = v59
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	F_MemoryContextReset(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v66 != int64(0) {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v81 = v21
	goto L13
L29:
	;
	goto L6
}
func F_ExecUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
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
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int64
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v436 int64
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
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
	var v486 int32
	_ = v486
	var v487 int64
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[0]))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L9
	} else {
		goto L141
	}
L2:
	;
	m.G0 = v17 + int32(48)
	return v505
L3:
	;
	if l6 != 0 {
		goto L128
	} else {
		goto L129
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L9
	} else {
		goto L124
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L121
	}
L6:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+28))
	m.T0[v31].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
	} else {
		goto L118
	}
L9:
	;
	return int32(0)
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+116)))
	if v37 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v45 = l0 + int32(16)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v40 != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	F_ExecOpenIndices(m, l1, int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v71 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+13)))
	if v49 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+188))
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_ExecPendingInserts(m, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L21
	}
L19:
	;
	v57 = v52
	goto L20
L20:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+104))
	v65 = F_ExecBRUpdateTriggers(m, v57, v59, l1, l2, l3, l5, v58, v45, base.B2i32(v62 == int32(5)))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L22
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v57 = v56
	goto L20
L22:
	;
	if v65 == int32(0) {
		v505 = v58
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[1]))
	if v75 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v91 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecUpdate[2])))
	if v77&int32(1) == int32(0) {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v82)+188))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+60))
	v86 = m.T0[v85].(func(*base.Module, int32, int32, int32, int32) int32)(m, v82, l2, int32(_a_F_ExecUpdate_0), l4)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v86 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v100 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+15)))
	if v94 != int32(1) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v97 = F_ExecIRUpdateTriggers(m, v20, l1, l3, l5)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v97 != 0 {
		v426 = l4
		v427 = l5
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v505 = int32(0)
	goto L2
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+40)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+52))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+24))
	if v105 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+28)) = uint16(v124)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v126
	v130 = F_ExecUpdateAct(m, l0, l1, l2, l3, l5, l6, v17+int32(32))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L49
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+64))
	v117 = m.T0[v116].(func(*base.Module, int32, int32, int32, int32) int32)(m, v20, l1, l5, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+17)))
	if v108 != int32(1) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	F_ExecComputeStoredGenerated(m, l1, v20, l5, int32(2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	if v117 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v505 = int32(0)
	goto L2
L47:
	;
	goto L48
L48:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+40)) = v122
	v426 = l4
	v427 = v117
	goto L3
L49:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v132&int32(1) == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L115
	}
L51:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[3]))
	if v348 < int32(2) {
		v505 = int32(0)
		goto L2
	} else {
		goto L110
	}
L52:
	;
	v141 = l4
	v142 = l5
	v144 = v130
	goto L55
L53:
	;
	goto L54
L54:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v505 = v345
	goto L2
L55:
	;
	if v144 != int32(3) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	switch v144 {
	case 0:
		v426 = v141
		v427 = v142
		goto L3
	default:
		goto L50
	case 2:
		goto L60
	case 4:
		goto L51
	}
L58:
	;
	goto L59
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[3]))
	if int32(2) <= v178 {
		goto L4
	} else {
		goto L67
	}
L60:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	if v154 == v155 {
		v505 = int32(0)
		goto L2
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_ExecUpdate_1), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	F_errhint(m, int32(_a_F_ExecUpdate_2), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2628), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
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
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v183 = F_EvalPlanQualSlot(m, v181, v19, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v185 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+104))
	v193 = m.T0[v192].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v19, l2, v186, v183, v187, v188, v185, int32(2), v45)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L72
	}
L69:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v289)+12)) = v199
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+80))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v288)+24))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	m.T0[v295].(func(*base.Module, int32))(m, v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L9
	} else {
		goto L106
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L103
	}
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v20)+64))
	if v250 == v251 {
		v505 = v185
		goto L2
	} else {
		goto L97
	}
L72:
	;
	if v193 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	switch v193 - int32(2) {
	case 0:
		goto L71
	default:
		goto L70
	case 2:
		v505 = v185
		goto L2
	}
L74:
	;
	goto L75
L75:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v199 = F_EvalPlanQual(m, v197, v19, v198, v183)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	if v199 == int32(0) {
		v505 = v185
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+4)))
	if v203&int32(2) != 0 {
		v505 = v185
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v206 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ExecInitUpdateProjection(m, v209, l1)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+49)))
	if v212 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_UnlockTuple(m, v19, v17+int32(24), int32(7))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[1]))
	if v225 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	F_LockTuple(m, v19, l2, int32(7))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecUpdate[2])))
	if v227&int32(1) == int32(0) {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+60))
	v235 = m.T0[v234].(func(*base.Module, int32, int32, int32, int32) int32)(m, v19, l2, int32(_a_F_ExecUpdate_0), v223)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	if v235 != 0 {
		goto L69
	} else {
		goto L93
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_5), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2692), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_ExecUpdate_1), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errhint(m, int32(_a_F_ExecUpdate_2), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2718), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v193
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_6), v17+int32(16))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2724), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v298 = int32(_a_F_ExecUpdate_7)
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4]))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v292)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4])) = v301
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v288)+32))
	v307 = m.T0[v306].(func(*base.Module, int32, int32, int32) int64)(m, v288+int32(8), v292, int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4])) = v299
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293)+4)))
	v313 = v311 & int32(_a_F_ExecUpdate_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v293)+4)) = uint16(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	*(*uint16)(unsafe.Add(mBase, uint32(v293)+6)) = uint16(v316)
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v17)+28)) = uint16(v318)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v320
	v324 = F_ExecUpdateAct(m, l0, l1, l2, l3, v293, l6, v17+int32(32))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+32)))
	if v326&int32(1) == int32(0) {
		v141 = v223
		v142 = v293
		v144 = v324
		goto L55
	} else {
		goto L109
	}
L109:
	;
	goto L56
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_ExecUpdate_9), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2735), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v144
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_10), v17)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2741), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_11), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2514), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L9
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_12), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2539), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L9
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L125
	}
L125:
	;
	F_errmsg(m, int32(_a_F_ExecUpdate_13), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_3), int32(2644), int32(_a_F_ExecUpdate_4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L9
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	v436 = *(*int64)(unsafe.Add(mBase, uint32(v20)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v436 + int64(1)
	goto L130
L129:
	;
	goto L130
L130:
	;
	F_ExecUpdateEpilogue(m, l0, v17+int32(32), l1, l2, l3, v427)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L9
	} else {
		goto L131
	}
L131:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v444 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v505 = int32(0)
	goto L2
L133:
	;
	goto L134
L134:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v444)+80))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v449)+12)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v449)+4)) = v427
	if v426 != 0 {
		v463 = int32(0)
		v464 = v426
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449)+72)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v449)+68)) = v464
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+12)))
	v470 = v467&int32(231) | v463
	*(*uint8)(unsafe.Add(mBase, uint32(v444)+12)) = uint8(v470)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v444)+80))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v444)+24))
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v473)+8))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v474)+12))
	m.T0[v475].(func(*base.Module, int32))(m, v473)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L9
	} else {
		goto L139
	}
L136:
	;
	v454 = int32(8)
	v455 = int32(0)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+12)))
	if v456&int32(2) == v455 {
		v463 = v454
		v464 = v455
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v461 = F_ExecGetAllNullSlot(m, v448, l1)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L9
	} else {
		goto L138
	}
L138:
	;
	v463 = v454
	v464 = v461
	goto L135
L139:
	;
	v478 = int32(_a_F_ExecUpdate_7)
	v479 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4]))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v472)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4])) = v481
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v444)+32))
	v487 = m.T0[v486].(func(*base.Module, int32, int32, int32) int64)(m, v444+int32(8), v472, int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdate[4])) = v479
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+4)))
	v493 = v491 & int32(_a_F_ExecUpdate_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+4)) = uint16(v493)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	*(*uint16)(unsafe.Add(mBase, uint32(v473)+6)) = uint16(v496)
	v505 = v473
	goto L2
L141:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdate_14), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L9
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_ExecUpdate_15), int32(1355), int32(_a_F_ExecUpdate_16))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L9
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecUpdateAct(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
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
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v8)
	v28 = l4
	goto L2
L1:
	;
	m.G0 = v18 + int32(48)
	return v563
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
	if v43 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v563 = v558
	goto L1
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	m.T0[v55].(func(*base.Module, int32))(m, v28)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L9
	}
L5:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+17)))
	if v46 != int32(1) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_ExecComputeStoredGenerated(m, l1, v21, v28, int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L4
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+131)))
	if v59 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+128))
	if v99 != int32(2) {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v63 = F_ExecPartitionCheck(m, l1, v28, v21, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v67 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v63 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	F_ExecWithCheckOptions(m, int32(2), l1, v28, v21)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+24))
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	F_ExecConstraints(m, l1, v28, v21)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v20)+188))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+100))
	v88 = m.T0[v87].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v20, l2, v28, v75, int32(0), v77, v78, int32(1), l0+int32(16), l6+int32(8), l6+int32(4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v563 = v88
	goto L1
L25:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v554)+104))
	if v555 != int32(5) {
		v28 = v550
		goto L2
	} else {
		goto L136
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L7
	} else {
		goto L130
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L7
	} else {
		goto L127
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L7
	} else {
		goto L124
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+120))
	if v102 == l1 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L7
	} else {
		goto L119
	}
L32:
	;
	F_ExecPartitionCheckEmitError(m, l1, v28, v93)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L7
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v92)+200))
	if v106 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v109 = int32(_a_F_ExecUpdateAct_0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v92)+120))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v93)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0])) = v114
	v116 = F_ExecSetupPartitionTupleRouting(m, v93, v112)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v127 = int32(0)
	v136 = F_ExecDelete(m, l0, l1, l2, l3, v127, int32(1), v127, v18+int32(36), v18+int32(47), v18+int32(40))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+200)) = v116
	v120 = F_table_slot_create(m, v112, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+196)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0])) = v110
	goto L38
L41:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	if v138 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v227)
	v229 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v230 == v229 {
		v563 = v229
		goto L1
	} else {
		goto L70
	}
L43:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v224 != 0 {
		v550 = int32(0)
		goto L25
	} else {
		goto L69
	}
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v92)+104))
	if v141 == int32(5) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v203 = F_ExecGetChildToRootMap(m, l1)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L62
	}
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v144 == int32(0) {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+4)))
	if v147&int32(2) != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v150 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_ExecInitUpdateProjection(m, v92, l1)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[1]))
	if v157 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecUpdateAct[2])))
	if v159&int32(1) == int32(0) {
		goto L28
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v164)+188))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+60))
	v168 = m.T0[v167].(func(*base.Module, int32, int32, int32, int32) int32)(m, v164, l2, int32(_a_F_ExecUpdateAct_1), v155)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v168 == int32(0) {
		goto L27
	} else {
		goto L59
	}
L59:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v172
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+80))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)+24))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	m.T0[v180].(func(*base.Module, int32))(m, v178)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v183 = int32(_a_F_ExecUpdateAct_0)
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0])) = v186
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	v192 = m.T0[v191].(func(*base.Module, int32, int32, int32) int64)(m, v173+int32(8), v177, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecUpdateAct[0])) = v184
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)))
	v198 = v196 & int32(_a_F_ExecUpdateAct_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*uint16)(unsafe.Add(mBase, uint32(v178)+6)) = uint16(v201)
	v550 = v178
	goto L25
L62:
	;
	if v203 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v92)+196))
	v207 = F_execute_attr_map_slot(m, v205, v28, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	v209 = v28
	goto L65
L65:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v92)+120))
	v215 = F_ExecInsert(m, l0, v210, v209, l5, v18+int32(32), v18+int32(28))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	v209 = v207
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v215
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v92)+204))
	if v218 == int32(0) {
		goto L42
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+4)) = int32(0)
	goto L42
L69:
	;
	goto L42
L70:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v233 == int32(0) {
		v563 = v229
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+14)))
	if v236 != int32(1) {
		v563 = v229
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v242 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+48))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+131)))
	if v245 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v457 = int32(0)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecARUpdateTriggers(m, v458, v239, l1, v230, l2, v457, v240, v457, v457, int32(1))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L118
	}
L74:
	;
	if v322 == int32(0) {
		goto L73
	} else {
		goto L97
	}
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v246 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L94
	}
L78:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v243)+56))
	v253 = F_get_partition_ancestors(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L82
	}
L79:
	;
	v322 = v246
	goto L80
L80:
	;
	goto L74
L81:
	;
	v316 = F_lappend(m, v306, v249)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L93
	}
L82:
	;
	if v253 == int32(0) {
		v306 = v242
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v257 <= int32(0) {
		v306 = v242
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v264 = int32(0)
	v266 = v242
	goto L85
L85:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v264<<(uint(int32(2))%32))))
	if v280 == v251 {
		v306 = v266
		goto L81
	} else {
		goto L87
	}
L86:
	;
	v306 = v295
	goto L81
L87:
	;
	v283 = F_table_open(m, v280, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v286 = F_palloc0(m, int32(216))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286))) = int32(394)
	v290 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v241)+132))
	F_InitResultRelInfo(m, v286, v283, v290, v290, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v295 = F_lappend(m, v266, v286)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v298 = v264 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v298 < v299 {
		v264 = v298
		v266 = v295
		goto L85
	} else {
		goto L92
	}
L92:
	;
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+212)) = v316
	v322 = v316
	goto L80
L94:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdateAct_3), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ExecUpdateAct_4), int32(1452), int32(_a_F_ExecUpdateAct_5))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v349 <= int32(0) {
		goto L73
	} else {
		goto L98
	}
L98:
	;
	v362 = int32(0)
	v364 = v349
	goto L99
L99:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v368+v362<<(uint(int32(2))%32))))
	if v372 == v239 {
		v435 = v364
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L73
L101:
	;
	v440 = v362 + int32(1)
	if v440 < v435 {
		v362 = v440
		v364 = v435
		goto L99
	} else {
		goto L117
	}
L102:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v372)+52))
	if v374 == int32(0) {
		v435 = v364
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+14)))
	if v377 != int32(1) {
		v435 = v364
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v380 = int32(0)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	if v381 <= v380 {
		v435 = v364
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v388 = v380
	v389 = v381
	goto L106
L106:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v402 = v399 + v388*int32(60)
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+16)))
	if v403 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	v435 = v423
	goto L101
L108:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402)+8))
	v408 = v406 - int32(1644)
	if base.Ui32(v408) <= base.Ui32(int32(11)) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v419 = v389
	goto L110
L110:
	;
	v421 = v388 + int32(1)
	if v421 < v419 {
		v388 = v421
		v389 = v419
		goto L106
	} else {
		goto L116
	}
L111:
	;
	if v415 == int32(1) {
		goto L26
	} else {
		goto L115
	}
L112:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v408<<(uint(int32(2))%32))+uint32(_c_F_ExecUpdateAct[3])))
	v415 = v413
	goto L114
L113:
	;
	v415 = int32(0)
	goto L114
L114:
	;
	goto L111
L115:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v419 = v418
	goto L110
L116:
	;
	goto L107
L117:
	;
	goto L100
L118:
	;
	v563 = v457
	goto L1
L119:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_ExecUpdateAct_6), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	v478 = F_errdetail(m, int32(_a_F_ExecUpdateAct_7), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_ExecUpdateAct_8), int32(1995), int32(_a_F_ExecUpdateAct_9))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdateAct_10), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ExecUpdateAct_11), int32(1355), int32(_a_F_ExecUpdateAct_12))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errmsg_internal(m, int32(_a_F_ExecUpdateAct_13), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ExecUpdateAct_8), int32(2084), int32(_a_F_ExecUpdateAct_9))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L7
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_ExecUpdateAct_14), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+48))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+48))
	v526 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v525 + v526
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v523 + v526
	v535 = F_errdetail(m, int32(_a_F_ExecUpdateAct_15), v18+int32(16))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v538 + int32(4)
	F_errhint(m, int32(_a_F_ExecUpdateAct_16), v18)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ExecUpdateAct_8), int32(2462), int32(_a_F_ExecUpdateAct_17))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	goto L3
}
func F_ExecUpdateLockMode(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v3 = F_ExecGetAllUpdatedCols(m, l1, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v11 = F_RelationGetIndexAttrBitmap(m, v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(0)
	if base.B2i32(v11 == v13)|base.B2i32(v3 == v13) != 0 {
		v58 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v23 < v24 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = v23
	goto L9
L8:
	;
	v26 = v24
	goto L9
L9:
	;
	if v26 <= int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v29 = int32(1)
	goto L12
L11:
	;
	v29 = v26
	goto L12
L12:
	;
	v30 = int32(8)
	v35 = int32(0)
	goto L13
L13:
	;
	v42 = v35 << (uint(int32(2)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v3+v30+v42)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11+v30+v42)))
	v47 = v44 & v46
	v49 = base.B2i32(v47 != int32(0))
	if v47 != 0 {
		v58 = v49
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v58 = v49
	goto L5
L15:
	;
	v51 = v35 + int32(1)
	if v51 != v29 {
		v35 = v51
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v59 = int32(3)
	goto L19
L18:
	;
	v59 = int32(2)
	goto L19
L19:
	;
	return v59
}
func F_ExecValuesScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(819), int32(820))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExecuteTruncateGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v251 int32
	_ = v251
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v308 int32
	_ = v308
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v398 int32
	_ = v398
	var v416 int32
	_ = v416
	var v435 int32
	_ = v435
	var v455 int32
	_ = v455
	var v473 int32
	_ = v473
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v544 int32
	_ = v544
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v671 int32
	_ = v671
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v849 int32
	_ = v849
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v891 int32
	_ = v891
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v952 int32
	_ = v952
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v989 int32
	_ = v989
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1297 int32
	_ = v1297
	var v1314 int32
	_ = v1314
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1458 int64
	_ = v1458
	var v1480 int32
	_ = v1480
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1668 int32
	_ = v1668
	var v1686 int32
	_ = v1686
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1732 int32
	_ = v1732
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1754 int32
	_ = v1754
	var v1764 int32
	_ = v1764
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1833 int32
	_ = v1833
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1898 int32
	_ = v1898
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1934 int32
	_ = v1934
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1967 int32
	_ = v1967
	var v2007 int32
	_ = v2007
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2086 int32
	_ = v2086
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2115 int32
	_ = v2115
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2144 int32
	_ = v2144
	var v2159 int32
	_ = v2159
	var v2163 int32
	_ = v2163
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int64
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int64
	_ = v2247
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2292 int32
	_ = v2292
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2310 int32
	_ = v2310
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2338 int32
	_ = v2338
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2381 int32
	_ = v2381
	var v2412 int32
	_ = v2412
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2461 int32
	_ = v2461
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2481 int32
	_ = v2481
	var v2500 int64
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2592 int32
	_ = v2592
	var v2595 int32
	_ = v2595
	var v2597 int32
	_ = v2597
	var v2614 int32
	_ = v2614
	var v2631 int32
	_ = v2631
	var v2647 int32
	_ = v2647
	var v2650 int32
	_ = v2650
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2727 int32
	_ = v2727
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2773 int32
	_ = v2773
	var v2784 int32
	_ = v2784
	var v2788 int32
	_ = v2788
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2880 int32
	_ = v2880
	var v2881 int64
	_ = v2881
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2897 int32
	_ = v2897
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2901 int32
	_ = v2901
	var v2902 int32
	_ = v2902
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2912 int32
	_ = v2912
	v7 = int32(0)
	v34 = m.G0
	v36 = v34 - int32(352)
	m.G0 = v36
	v39 = l0
	v40 = l1
	v41 = l2
	v42 = l3
	v43 = l4
	v44 = l5
	v45 = v36
	v46 = v7
	v47 = v7
	v48 = v7
	v49 = v7
	v50 = v7
	v51 = v7
	v52 = v7
	v53 = v7
	v54 = v7
	v55 = v7
	v56 = v7
	v57 = v7
	v59 = int32(-1)
	v60 = v7
	v61 = v7
	v62 = v7
	v67 = v7
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v59 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v2880 = int32(m.ExcTag)
	v2881 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2880 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L7:
	;
	if v2107 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[0])) = v1815
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[1])) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	v2062 = int32(1)
	v2063 = v1828 & v2062
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2063)
	v2066 = v1827 & v2062
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2066)
	F_hash_destroy(m, v1830)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L193
	}
L9:
	;
	v2094 = v2019
	v2095 = v2020
	v2096 = v2021
	v2097 = v2022
	v2098 = v2023
	v2099 = v2037
	v2100 = v2025
	v2101 = v2026
	v2102 = v2027
	v2103 = v2028
	v2106 = v2031
	v2107 = v2032
	v2108 = v2019
	v2109 = v2034
	v2110 = v2035
	v2111 = v2036
	v2115 = v2021
	goto L7
L10:
	;
	if v1833 != 0 {
		goto L8
	} else {
		goto L181
	}
L11:
	;
	v1812 = v46
	v1813 = v47
	v1814 = v48
	v1815 = v49
	v1816 = v50
	v1819 = v53
	v1820 = v54
	v1821 = v55
	v1824 = v56
	v1825 = v57
	v1827 = v61
	v1828 = v62
	v1829 = v60
	v1830 = v51
	v1833 = v67
	goto L10
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v53
	v87 = int32(1)
	v88 = v62 & v87
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v91 = v61 & v87
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v91)
	v93 = F_list_copy(m, v39)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v95 = int32(1)
	v96 = base.B2i32(v42 == v95)
	if v42 != v95 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v705 = int32(0)
	if base.B2i32(v43 == v705)|base.B2i32(v691 == v705) != 0 {
		v989 = v705
		goto L79
	} else {
		goto L80
	}
L16:
	;
	if v42 != 0 {
		v686 = v635
		v688 = v637
		v691 = v640
		v696 = v645
		goto L15
	} else {
		goto L77
	}
L17:
	;
	v635 = v53
	v637 = v55
	v640 = v93
	v645 = v41
	goto L16
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v53
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v114 = F_heap_truncate_find_FKs(m, v40)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if v114 == int32(0) {
		v686 = v53
		v688 = v55
		v691 = v93
		v696 = v41
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v132 = v53
	v134 = v55
	v137 = v93
	v142 = v41
	v143 = v40
	v144 = v114
	goto L22
L22:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if int32(0) < v151 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v635 = v619
	v637 = v587
	v640 = v590
	v645 = v595
	goto L16
L24:
	;
	v171 = v134
	v174 = v137
	v179 = v142
	v180 = v143
	v183 = int32(0)
	goto L27
L25:
	;
	v587 = v134
	v590 = v137
	v595 = v142
	v596 = v143
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v587
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v619 = F_heap_truncate_find_FKs(m, v596)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L75
	}
L27:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+v183<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v209 = F_table_open(m, v192, int32(8))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L29
	}
L28:
	;
	v587 = v564
	v590 = v489
	v595 = v565
	v596 = v506
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v228 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if v228 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v230 + int32(4)
	F_errmsg(m, int32(_a_F_ExecuteTruncateGuts_0), v45)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_truncate_check_rel(m, v192, v273)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L36
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_errfinish(m, int32(_a_F_ExecuteTruncateGuts_1), int32(2061), int32(_a_F_ExecuteTruncateGuts_2))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v325 = F_pg_class_aclcheck(m, v192, v308, int64(16))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v325 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v327 = int32(*(*int8)(unsafe.Add(mBase, uint32(v291)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	switch v327 - int32(73) {
	case 0, 32:
		v352 = int32(20)
		goto L42
	default:
		goto L43
	case 10:
		goto L47
	case 29:
		goto L44
	case 36:
		goto L45
	case 45:
		goto L46
	}
L39:
	;
	goto L40
L40:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+118)))
	if v376 != int32(116) {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_aclcheck_error(m, v325, v354, v291+int32(4))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L48
	}
L42:
	;
	v354 = v352
	goto L41
L43:
	;
	v352 = int32(42)
	goto L42
L44:
	;
	v354 = int32(18)
	goto L41
L45:
	;
	v354 = int32(23)
	goto L41
L46:
	;
	v354 = int32(52)
	goto L41
L47:
	;
	v354 = int32(38)
	goto L41
L48:
	;
	goto L40
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_CheckTableNotInUse(m, v209, int32(_a_F_ExecuteTruncateGuts_3))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L56
	}
L50:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+24)))
	if v379 != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_errcode(m, int32(1088))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_errmsg(m, int32(_a_F_ExecuteTruncateGuts_4), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_errfinish(m, int32(_a_F_ExecuteTruncateGuts_1), int32(2481), int32(_a_F_ExecuteTruncateGuts_5))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L55
	}
L55:
	;
	goto L3
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v489 = F_lappend(m, v174, v209)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v506 = F_lappend_oid(m, v180, v192)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L58
	}
L58:
	;
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[3]))
	if v509 <= int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v568 = v183 + int32(1)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v568 < v569 {
		v171 = v564
		v174 = v489
		v179 = v565
		v180 = v506
		v183 = v568
		goto L27
	} else {
		goto L74
	}
L60:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[4])))
	if v513&int32(1) == int32(0) {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v209)+48))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+118)))
	if v519 != int32(112) {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	if v509 <= int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v209)+32))
	if v524 != 0 {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+119)))
	if v526 == int32(102) {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L70
	}
L68:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
	if v525 != 0 {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v209)+56))
	goto L71
L71:
	;
	if base.Ui32(v544) < base.Ui32(int32(_a_F_ExecuteTruncateGuts_6)) {
		v564 = v171
		v565 = v179
		goto L59
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v132
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v562 = F_lappend_oid(m, v179, v192)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L73
	}
L73:
	;
	v564 = v562
	v565 = v562
	goto L59
L74:
	;
	goto L28
L75:
	;
	if v619 != 0 {
		v132 = v619
		v134 = v587
		v137 = v590
		v142 = v595
		v143 = v596
		v144 = v619
		goto L22
	} else {
		goto L76
	}
L76:
	;
	goto L23
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v637
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v635
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_heap_truncate_check_FKs(m, v640, int32(0))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v686 = v635
	v688 = v637
	v691 = v640
	v696 = v645
	goto L15
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v1017 = int32(_a_F_ExecuteTruncateGuts_7)
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[5])) = v1019 + int32(1)
	goto L100
L80:
	;
	v711 = int32(0)
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v712 <= v711 {
		v989 = v705
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v735 = v705
	v740 = v711
	goto L82
L82:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v748+v740<<(uint(int32(2))%32))))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v752)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v769 = F_getOwnedSequences(m, v753)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L85
	}
L83:
	;
	v989 = v952
	goto L79
L84:
	;
	v966 = v740 + int32(1)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v966 < v967 {
		v735 = v952
		v740 = v966
		goto L82
	} else {
		goto L99
	}
L85:
	;
	if v769 == int32(0) {
		v952 = v735
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v773 = int32(0)
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v774 <= v773 {
		v952 = v735
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v797 = v735
	v798 = v773
	goto L88
L88:
	;
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v769)+12))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v810+v798<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v831 = F_relation_open(m, v814, int32(8))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L90
	}
L89:
	;
	v952 = v908
	goto L84
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v866 = F_object_ownercheck(m, int32(1259), v814, v849)
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L91
	}
L91:
	;
	if v866 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v831)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_aclcheck_error(m, int32(2), int32(38), v870+int32(4))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v908 = F_lappend_oid(m, v797, v814)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	F_relation_close(m, v831, int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v929 = v798 + int32(1)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	if v929 < v930 {
		v797 = v908
		v798 = v929
		goto L88
	} else {
		goto L98
	}
L98:
	;
	goto L89
L99:
	;
	goto L83
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v88)
	v1038 = F_CreateExecutorState(m)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L101
	}
L101:
	;
	v1041 = base.B2i32(v691 == int32(0))
	if v691 != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v1222 = int32(0)
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	if v1222 < v1223 {
		goto L115
	} else {
		goto L116
	}
L103:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1060 = F_palloc(m, v1042*int32(216))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1167 = F_palloc(m, int32(0))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L113
	}
L106:
	;
	v1063 = v691 + int32(4)
	v1064 = int32(0)
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v1065 <= v1064 {
		goto L102
	} else {
		goto L107
	}
L107:
	;
	v1085 = v1064
	v1086 = v1060
	goto L108
L108:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1101+v1085<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1121 = int32(0)
	F_InitResultRelInfo(m, v1086, v1105, v1121, v1121, v1121)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L110
	}
L109:
	;
	goto L102
L110:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1142 = F_lappend(m, v1126, v1086)
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1038)+72)) = v1142
	v1148 = v1085 + int32(1)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	if v1148 < v1149 {
		v1085 = v1148
		v1086 = v1086 + int32(216)
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	goto L114
L114:
	;
	v2094 = v46
	v2095 = v1038
	v2096 = v48
	v2097 = v49
	v2098 = v50
	v2099 = v51
	v2100 = v52
	v2101 = v686
	v2102 = v54
	v2103 = v688
	v2106 = v691
	v2107 = v989
	v2108 = v1167
	v2109 = v96
	v2110 = v1041
	v2111 = v696
	v2115 = v691 + int32(4)
	goto L7
L115:
	;
	v1243 = v1222
	v1244 = v1060
	goto L118
L116:
	;
	goto L117
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[6]))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+8))
	goto L129
L118:
	;
	if v44 != 0 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L117
L120:
	;
	v1336 = v1243 + int32(1)
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	if v1336 < v1337 {
		v1243 = v1336
		v1244 = v1244 + int32(216)
		goto L118
	} else {
		goto L128
	}
L121:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+8))
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+48))
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1260)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1278 = v45 + int32(280)
	F_SwitchToUntrustedUser(m, v1261, v1278)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_ExecBSTruncateTriggers(m, v1038, v1244)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L127
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_ExecBSTruncateTriggers(m, v1038, v1244)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_RestoreUserContext(m, v1278)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L126
	}
L126:
	;
	goto L120
L127:
	;
	goto L120
L128:
	;
	goto L119
L129:
	;
	v1390 = int32(0)
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v1390 < v1392 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v1410 = v54
	v1413 = v1390
	v1420 = v1390
	goto L133
L131:
	;
	v1754 = v54
	v1764 = v1390
	goto L132
L132:
	;
	if v1764 == int32(0) {
		v2019 = v1060
		v2020 = v1038
		v2021 = v1063
		v2022 = v49
		v2023 = v50
		v2025 = v52
		v2026 = v686
		v2027 = v1754
		v2028 = v688
		v2031 = v691
		v2032 = v989
		v2034 = v96
		v2035 = v1041
		v2036 = v696
		v2037 = v1764
		goto L9
	} else {
		goto L175
	}
L133:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1428+v1413<<(uint(int32(2))%32))))
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+48))
	v1434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1433)+119)))
	switch v1434 - int32(102) {
	case 0:
		goto L137
	default:
		goto L136
	case 10:
		v1730 = v1410
		v1732 = v1420
		goto L135
	}
L134:
	;
	v1754 = v1730
	v1764 = v1732
	goto L132
L135:
	;
	v1736 = v1413 + int32(1)
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1063)))
	if v1736 < v1737 {
		v1410 = v1730
		v1413 = v1736
		v1420 = v1732
		goto L133
	} else {
		goto L174
	}
L136:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+32))
	if v1389 != v1545 {
		goto L151
	} else {
		goto L152
	}
L137:
	;
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1453 = F_GetForeignServerIdByRelId(m, v1437)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+276)) = v1453
	if v1420 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v1458 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v45)+256)) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v45)+264)) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v45)+248)) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v45)+240)) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v45)+224)) = v1458
	*(*int64)(unsafe.Add(mBase, uint32(v45)+232)) = int64(34359738372)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	v1480 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+260)) = v1480
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1493 = F_hash_create(m, int32(_a_F_ExecuteTruncateGuts_8), int64(32), v45+int32(224), int32(1064))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L142
	}
L140:
	;
	v1495 = v1410
	v1496 = v1420
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1517 = F_hash_search(m, v1496, v45+int32(276), int32(1), v45+int32(275))
	mBase = m.M
	v1518 = m.ExcPending
	if v1518 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L143
	}
L142:
	;
	v1495 = v1493
	v1496 = v1493
	goto L141
L143:
	;
	v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+275)))
	if v1519 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1495
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1542 = F_lappend(m, v1526, v1432)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L148
	}
L145:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1517)+4))
	v1526 = v1522
	goto L144
L146:
	;
	goto L147
L147:
	;
	v1523 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+4)) = v1523
	v1526 = v1523
	goto L144
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+4)) = v1542
	v1730 = v1495
	v1732 = v1496
	goto L135
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_pgstat_count_truncate(m, v1432)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L173
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v45)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_CheckTableForSerializableConflictIn(m, v1432)
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L164
	}
L151:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+36))
	if v1547 != v1389 {
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+48))
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564)+119)))
	if v1565 == int32(112) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	goto L153
L155:
	;
	goto L149
L156:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+188))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+116))
	m.T0[v1569].(func(*base.Module, int32))(m, v1432)
	mBase = m.M
	v1571 = m.ExcPending
	if v1571 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L157
	}
L157:
	;
	F_RelationTruncateIndexes(m, v1432)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L158
	}
L158:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+48))
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+112))
	if v1575 == int32(0) {
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v1579 = F_table_open(m, v1575, int32(8))
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L160
	}
L160:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1579)+188))
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+116))
	m.T0[v1582].(func(*base.Module, int32))(m, v1579)
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L161
	}
L161:
	;
	F_RelationTruncateIndexes(m, v1579)
	mBase = m.M
	v1586 = m.ExcPending
	if v1586 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L162
	}
L162:
	;
	F_relation_close(m, v1579, int32(0))
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L163
	}
L163:
	;
	goto L155
L164:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+48))
	v1611 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1610)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_RelationSetNewRelfilenumber(m, v1432, v1611)
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L165
	}
L165:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+56))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+48))
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+112))
	if v1631 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1648 = F_relation_open(m, v1631, int32(8))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	v1708 = F_reindex_relation(m, int32(0), v1629, int32(1), v45+int32(216))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L172
	}
L169:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1648)+48))
	v1651 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1650)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_RelationSetNewRelfilenumber(m, v1648, v1651)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_relation_close(m, v1648, int32(0))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	goto L149
L173:
	;
	v1730 = v1410
	v1732 = v1420
	goto L135
L174:
	;
	goto L134
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1754
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1764
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1063
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1060
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1041)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1038
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v686
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v96)
	F_hash_seq_init(m, v45+int32(196), v1764)
	mBase = m.M
	v1792 = m.ExcPending
	if v1792 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[0]))
	v1797 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[1]))
	goto L177
L177:
	;
	v1799 = v45 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v1799)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1799))) = v45 + int32(4)
	goto L180
L178:
	;
	v1812 = v1060
	v1813 = v1038
	v1814 = v1063
	v1815 = v1795
	v1816 = v1797
	v1819 = v686
	v1820 = v1754
	v1821 = v688
	v1824 = v691
	v1825 = v989
	v1827 = v96
	v1828 = v1041
	v1829 = v696
	v1830 = v1764
	v1833 = int32(0)
	goto L10
L180:
	;
	goto L178
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[1])) = v45 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	v1855 = int32(1)
	v1856 = v1828 & v1855
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1856)
	v1859 = v1827 & v1855
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v1859)
	v1863 = F_hash_seq_search(m, v45+int32(196))
	mBase = m.M
	v1864 = m.ExcPending
	if v1864 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L182
	}
L182:
	;
	if v1863 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v1877 = v1863
	v1878 = v52
	goto L186
L184:
	;
	v1967 = v52
	goto L185
L185:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[0])) = v1815
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[1])) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1856)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v1859)
	F_hash_destroy(m, v1830)
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L192
	}
L186:
	;
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1877)))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v1878
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1856)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v1859)
	v1914 = F_GetFdwRoutineByServerId(m, v1898)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L188
	}
L187:
	;
	v1967 = v1952
	goto L185
L188:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1877)+4))
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1914)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v1878
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1856)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v1859)
	m.T0[v1917].(func(*base.Module, int32, int32, int32))(m, v1916, v42, v43)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v1878
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v1856)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v1859)
	v1952 = F_hash_seq_search(m, v45+int32(196))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L190
	}
L190:
	;
	if v1952 != 0 {
		v1877 = v1952
		v1878 = v1952
		goto L186
	} else {
		goto L191
	}
L191:
	;
	goto L187
L192:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[0])) = v1815
	*(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[1])) = v1816
	v2019 = v1812
	v2020 = v1813
	v2021 = v1814
	v2022 = v1815
	v2023 = v1816
	v2025 = v1967
	v2026 = v1819
	v2027 = v1820
	v2028 = v1821
	v2031 = v1824
	v2032 = v1825
	v2034 = v1827
	v2035 = v1828
	v2036 = v1829
	v2037 = v1830
	goto L9
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v1815
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v1816
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v1830
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v1814
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v1824
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v1821
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v1819
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2063)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2066)
	F_pg_re_throw(m)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L194
	}
L194:
	;
	goto L3
L195:
	;
	if v2111 != 0 {
		goto L216
	} else {
		goto L217
	}
L196:
	;
	v2122 = int32(0)
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+4))
	if v2123 <= v2122 {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v2144 = v2122
	goto L198
L198:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+12))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2159+v2144<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	v2171 = int32(1)
	v2172 = v2110 & v2171
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2172)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	v2181 = v2109 & v2171
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2181)
	v2183 = m.G0
	v2185 = v2183 - int32(48)
	m.G0 = v2185
	F_init_sequence(m, v2163, v2185+int32(40), v2185+int32(44))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L200
	}
L199:
	;
	goto L195
L200:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+44))
	v2198 = F_read_seq_tuple(m, v2193, v2185+int32(36), v2185+int32(16))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L201
	}
L201:
	;
	v2202 = F_SearchSysCache1(m, int32(61), base.I64_extend_i32_u(v2163))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L202
	}
L202:
	;
	if v2202 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2202)+16))
	v2220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2219)+22)))
	v2222 = *(*int64)(unsafe.Add(mBase, uint32(v2219+v2220)+8))
	F_ReleaseCatCache(m, v2202)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L209
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2185))) = v2163
	F_errmsg_internal(m, int32(_a_F_ExecuteTruncateGuts_9), v2185)
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_ExecuteTruncateGuts_10), int32(277), int32(_a_F_ExecuteTruncateGuts_11))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v2227 = F_heap_copytuple(m, v2185+int32(16))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+36))
	F_UnlockReleaseBuffer(m, v2229)
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L211
	}
L211:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2227)+16))
	v2233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2232)+22)))
	v2234 = v2232 + v2233
	v2235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2234)+16)) = uint8(v2235)
	*(*int64)(unsafe.Add(mBase, uint32(v2234))) = v2222
	*(*int64)(unsafe.Add(mBase, uint32(v2234)+8)) = int64(0)
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2193)+48))
	v2241 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2240)+118)))
	F_RelationSetNewRelfilenumber(m, v2193, v2241)
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L212
	}
L212:
	;
	F_fill_seq_with_data(m, v2193, v2227)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L213
	}
L213:
	;
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v2185)+40))
	v2247 = *(*int64)(unsafe.Add(mBase, uint32(v2246)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2246)+24)) = v2247
	F_relation_close(m, v2193, int32(0))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L214
	}
L214:
	;
	m.G0 = v2185 + int32(48)
	v2256 = v2144 + int32(1)
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v2107)+4))
	if v2256 < v2257 {
		v2144 = v2256
		goto L198
	} else {
		goto L215
	}
L215:
	;
	goto L199
L216:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	v2300 = int32(1)
	v2301 = v2110 & v2300
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	v2310 = v2109 & v2300
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	v2314 = F_palloc(m, v2292<<(uint(int32(2))%32))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v2536 = v2110 & int32(1)
	if v2536 != 0 {
		goto L234
	} else {
		goto L235
	}
L219:
	;
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	if int32(0) < v2316 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v2338 = int32(0)
	goto L223
L221:
	;
	v2381 = v2316
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+24)) = v2381
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	v2412 = *(*int32)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v2412
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	if v43 != 0 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v2354 = v2338 << (uint(int32(2)) % 32)
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+12))
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2356+v2354)))
	*(*int32)(unsafe.Add(mBase, uint32(v2314+v2354))) = v2358
	v2361 = v2338 + int32(1)
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	if v2361 < v2362 {
		v2338 = v2361
		goto L223
	} else {
		goto L225
	}
L224:
	;
	v2381 = v2362
	goto L222
L225:
	;
	goto L224
L226:
	;
	v2418 = v2310 | int32(2)
	goto L228
L227:
	;
	v2418 = v2310
	goto L228
L228:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)) = uint8(v2418)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	F_XLogRegisterData(m, v45+int32(20), int32(12))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	F_XLogRegisterData(m, v2314, v2442<<(uint(int32(2))%32))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	v2478 = int32(_a_F_ExecuteTruncateGuts_12)
	v2480 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[9])))
	v2481 = v2480 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExecuteTruncateGuts[9])) = uint8(v2481)
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2301)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2310)
	v2500 = F_XLogInsert(m, int32(10), int32(48))
	mBase = m.M
	v2501 = m.ExcPending
	if v2501 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L233
	}
L233:
	;
	goto L218
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	v2707 = v2109 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2707)
	F_AfterTriggerEndQuery(m, v2095)
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L248
	}
L235:
	;
	v2537 = int32(0)
	v2538 = *(*int32)(unsafe.Add(mBase, uint32(v2115)))
	if v2538 <= v2537 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v2559 = v2537
	v2562 = v2108
	goto L237
L237:
	;
	if v44 != 0 {
		goto L240
	} else {
		goto L241
	}
L238:
	;
	goto L234
L239:
	;
	v2656 = v2559 + int32(1)
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v2115)))
	if v2656 < v2657 {
		v2559 = v2656
		v2562 = v2562 + int32(216)
		goto L237
	} else {
		goto L247
	}
L240:
	;
	v2574 = *(*int32)(unsafe.Add(mBase, uint32(v2562)+8))
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v2574)+48))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v2575)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	v2592 = v2109 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2592)
	v2595 = v45 + int32(8)
	F_SwitchToUntrustedUser(m, v2576, v2595)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	v2647 = v2109 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2647)
	F_ExecASTruncateTriggers(m, v2095, v2562)
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L246
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2592)
	F_ExecASTruncateTriggers(m, v2095, v2562)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2592)
	F_RestoreUserContext(m, v2595)
	mBase = m.M
	v2631 = m.ExcPending
	if v2631 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L245
	}
L245:
	;
	goto L239
L246:
	;
	goto L239
L247:
	;
	goto L238
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2707)
	F_FreeExecutorState(m, v2095)
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2707)
	v2743 = F_list_difference_ptr(m, v2106, v39)
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L251
	}
L250:
	;
	m.G0 = v45 + int32(352)
	return
L251:
	;
	if v2743 == int32(0) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v2747 = int32(0)
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+4))
	if v2748 <= v2747 {
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v2773 = v2747
	goto L254
L254:
	;
	v2784 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+12))
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2784+v2773<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+296)) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v45)+292)) = v2100
	*(*int32)(unsafe.Add(mBase, uint32(v45)+300)) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v45)+304)) = v2102
	*(*int32)(unsafe.Add(mBase, uint32(v45)+308)) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v45)+312)) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v45)+316)) = v2094
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+323)) = uint8(v2536)
	*(*int32)(unsafe.Add(mBase, uint32(v45)+324)) = v2095
	*(*int32)(unsafe.Add(mBase, uint32(v45)+328)) = v2107
	*(*int32)(unsafe.Add(mBase, uint32(v45)+332)) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v45)+336)) = v2111
	*(*int32)(unsafe.Add(mBase, uint32(v45)+340)) = v2103
	*(*int32)(unsafe.Add(mBase, uint32(v45)+344)) = v2101
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+351)) = uint8(v2707)
	F_relation_close(m, v2788, int32(0))
	mBase = m.M
	v2806 = m.ExcPending
	if v2806 != 0 {
		v2847 = v39
		v2848 = v40
		v2849 = v41
		v2850 = v42
		v2851 = v43
		v2852 = v44
		v2853 = v45
		goto L6
	} else {
		goto L256
	}
L255:
	;
	goto L250
L256:
	;
	v2808 = v2773 + int32(1)
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v2743)+4))
	if v2808 < v2809 {
		v2773 = v2808
		goto L254
	} else {
		goto L257
	}
L257:
	;
	goto L255
L258:
	;
	v2885 = int32(v2881)
	m.G0 = v2853
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v2885)+4))
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2885)))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2888)))
	if v2853+int32(4) == v2891 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	m.ExcPending = 1
	goto L267
L260:
	;
	if v2895 != 0 {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v2888)+4))
	v2895 = v2893
	goto L263
L262:
	;
	v2895 = int32(0)
	goto L263
L263:
	;
	goto L260
L264:
	;
	v2896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2853)+351)))
	v2897 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+344))
	v2898 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+340))
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+336))
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+332))
	v2901 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+328))
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+324))
	v2903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2853)+323)))
	v2904 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+316))
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+312))
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+308))
	v2907 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+304))
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+300))
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+296))
	v2910 = *(*int32)(unsafe.Add(mBase, uint32(v2853)+292))
	v39 = v2847
	v40 = v2848
	v41 = v2849
	v42 = v2850
	v43 = v2851
	v44 = v2852
	v45 = v2853
	v46 = v2904
	v47 = v2902
	v48 = v2905
	v49 = v2909
	v50 = v2908
	v51 = v2906
	v52 = v2910
	v53 = v2897
	v54 = v2907
	v55 = v2898
	v56 = v2900
	v57 = v2901
	v59 = v2895
	v60 = v2899
	v61 = v2896
	v62 = v2903
	v67 = v2887
	goto L1
L265:
	;
	goto L266
L266:
	;
	F___wasm_longjmp(m, v2888, v2887)
	mBase = m.M
	v2912 = m.ExcPending
	if v2912 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	return
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExitParallelMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ExitParallelMode[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+72)) = v4 - int32(1)
	return
}
func F_ExitPostmaster(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_proc_exit(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_ExportSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	v9 = m.G0
	v11 = v9 - int32(2448)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[0]))
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L6
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L22
	} else {
		goto L128
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L22
	} else {
		goto L124
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L22
	} else {
		goto L120
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L22
	} else {
		goto L116
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L22
	} else {
		goto L109
	}
L6:
	;
	if base.B2i32(int32(1) < v17) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[1]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	if v26 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L22
	} else {
		goto L105
	}
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v32 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v29 = v27
	goto L13
L12:
	;
	v29 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(2444)))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	goto L10
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[2]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[3]))
	goto L17
L15:
	;
	goto L16
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[4]))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[5]))
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if (v38+v40)*int32(65) < v35+v31 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v55 = v51 + int32(1)
	goto L21
L20:
	;
	v55 = int32(1)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+360)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v11)+356)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v11)+352)) = v48
	v65 = F_pg_snprintf(m, v11+int32(1392), int32(1024), int32(_a_F_ExportSnapshot_0), v11+int32(352))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[6]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v74 = int32(2)
	v76 = int32(72)
	v81 = v72<<(uint(v74)%32) + v76
	if int32(0) < v71 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = (v71+v72)<<(uint(v74)%32) + v76
	goto L26
L25:
	;
	v84 = v81
	goto L26
L26:
	;
	v85 = F_MemoryContextAlloc(m, v70, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v101
	*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = int64(0)
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+48)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v85)+44)) = v105
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+30)) = uint8(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v111 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v126 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v113 = v85 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v117 = v115 << (uint(int32(2)) % 32)
	if v117 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = int32(0)
	goto L28
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v113, v120, v117)
	goto L28
L33:
	;
	v149 = int32(_a_F_ExportSnapshot_1)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[7]))
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[7])) = v153
	v156 = F_palloc(m, int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L22
	} else {
		goto L41
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = int32(0)
	goto L33
L35:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v129 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v132 != int32(1) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v135 = v85 + v81
	*(*int32)(unsafe.Add(mBase, uint32(v85)+20)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v139 = v137 << (uint(int32(2)) % 32)
	if v139 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v135, v142, v139)
	goto L33
L41:
	;
	v160 = F_pstrdup(m, v11+int32(1392))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+4)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v160
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[5]))
	v166 = F_lappend(m, v165, v156)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[7])) = v150
	*(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[5])) = v166
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v85)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+48)) = v172 + int32(1)
	F_pairingheap_add(m, int32(_a_F_ExportSnapshot_2), v85+int32(52))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	v182 = v11 + int32(2428)
	F_initStringInfo(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[4]))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+336)) = v187
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_3), v11+int32(336))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+320)) = v195
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_4), v11+int32(320))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+304)) = v203
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_5), v11+int32(304))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L22
	} else {
		goto L48
	}
L48:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+288)) = v211
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_6), v11+int32(288))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExportSnapshot[11])))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+272)) = v219
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_7), v11+int32(272))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+256)) = v226
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_8), v11+int32(256))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+240)) = v233
	F_appendStringInfo(m, v182, int32(_a_F_ExportSnapshot_9), v11+int32(240))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L22
	} else {
		goto L52
	}
L52:
	;
	if v14 == int32(0) {
		v255 = int32(0)
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+224)) = v256 + v255
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_10), v11+int32(224))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L22
	} else {
		goto L58
	}
L54:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v242))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v14)) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v255 = base.B2i32(base.Ui32(v14) < base.Ui32(v242))
	goto L53
L56:
	;
	goto L57
L57:
	;
	v255 = int32(base.Ui32(v14-v242) >> (uint(int32(31)) % 32))
	goto L53
L58:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v266 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v268 = int32(0)
	goto L62
L60:
	;
	goto L61
L61:
	;
	if v255 != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v268<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v280
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_11), v11+int32(208))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L22
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v290 = v268 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if base.Ui32(v290) < base.Ui32(v291) {
		v268 = v290
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = v14
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_11), v11+int32(192))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L22
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	v310 = v309 + v31
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+28)))
	if v311 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v329
	v332 = v11 + int32(2428)
	F_appendStringInfo(m, v332, int32(_a_F_ExportSnapshot_12), v11+int32(176))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L22
	} else {
		goto L79
	}
L71:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[2]))
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[3]))
	goto L74
L72:
	;
	goto L73
L73:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+29)))
	if v325 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	if v310 <= (v316+v318)*int32(65) {
		v328 = v310
		v329 = int32(0)
		goto L70
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v326 = v310
	goto L78
L77:
	;
	v326 = int32(0)
	goto L78
L78:
	;
	v328 = v326
	v329 = int32(1)
	goto L70
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = v328
	F_appendStringInfo(m, v332, int32(_a_F_ExportSnapshot_13), v11+int32(160))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L22
	} else {
		goto L80
	}
L80:
	;
	if v328 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+29)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v418
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_14), v11+int32(112))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L22
	} else {
		goto L95
	}
L82:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	if int32(0) < v346 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v350 = int32(0)
	goto L86
L84:
	;
	goto L85
L85:
	;
	v383 = int32(0)
	if v31 <= v383 {
		goto L81
	} else {
		goto L90
	}
L86:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v350<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v362
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_15), v11+int32(144))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L22
	} else {
		goto L88
	}
L87:
	;
	goto L85
L88:
	;
	v372 = v350 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v85)+24))
	if v372 < v373 {
		v350 = v372
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v386 = v383
	goto L91
L91:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2444))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394+v386<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = v398
	F_appendStringInfo(m, v11+int32(2428), int32(_a_F_ExportSnapshot_15), v11+int32(128))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L22
	} else {
		goto L93
	}
L92:
	;
	goto L81
L93:
	;
	v408 = v386 + int32(1)
	if v408 != v31 {
		v386 = v408
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v428 = v11 + int32(1392)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v428
	v431 = v11 + int32(368)
	v436 = F_pg_snprintf(m, v431, int32(1024), int32(_a_F_ExportSnapshot_16), v11+int32(96))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	v439 = F_AllocateFile(m, v431, int32(_a_F_ExportSnapshot_17))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L22
	} else {
		goto L97
	}
L97:
	;
	if v439 == int32(0) {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2428))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v11)+2432))
	v446 = F_fwrite(m, v443, v444, int32(1), v439)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L22
	} else {
		goto L99
	}
L99:
	;
	if v446 != int32(1) {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	v450 = F_FreeFile(m, v439)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L22
	} else {
		goto L101
	}
L101:
	;
	if v450 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	v452 = F_rename(m, v431, v428)
	mBase = m.M
	if v452 < int32(0) {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v457 = F_pstrdup(m, v428|int32(13))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L22
	} else {
		goto L104
	}
L104:
	;
	m.G0 = v11 + int32(2448)
	return v457
L105:
	;
	F_errcode(m, int32(16777538))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_ExportSnapshot_18), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L22
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1158), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L22
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L22
	} else {
		goto L110
	}
L110:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v487 = v486 + v31
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v487
	F_errmsg_plural(m, int32(_a_F_ExportSnapshot_21), int32(_a_F_ExportSnapshot_22), v487, v11+int32(16))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[2]))
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[3]))
	goto L112
L112:
	;
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[2]))
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_ExportSnapshot[3]))
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = (v503 + v505) * int32(65)
	v510 = int32(_a_F_ExportSnapshot_23)
	F_errdetail_plural(m, v510, v510, (v496+v498)*int32(65), v11)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L22
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1190), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L22
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L22
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(368)
	F_errmsg(m, int32(_a_F_ExportSnapshot_24), v11+int32(32))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L22
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1294), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L22
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L22
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = v11 + int32(368)
	F_errmsg(m, int32(_a_F_ExportSnapshot_25), v11+int32(80))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L22
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1299), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L124:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v11 + int32(368)
	F_errmsg(m, int32(_a_F_ExportSnapshot_25), v11-int32(-64))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L22
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1306), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L22
	} else {
		goto L127
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L128:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v11 + int32(1392)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v11 + int32(368)
	F_errmsg(m, int32(_a_F_ExportSnapshot_26), v11+int32(48))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L22
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ExportSnapshot_19), int32(1316), int32(_a_F_ExportSnapshot_20))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L22
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExtendBufferedRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v13
	if v11&int32(255) == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+118)))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)) = uint8(v21)
	} else {
	}
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v25
	v35 = F_ExtendBufferedRelCommon(m, v9+int32(8), l1, l2, l3, int32(1), int32(-1), v9+int32(28), v9+int32(24))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		m.G0 = v9 + int32(48)
		return v39
	}
}
func F__equalPLAssignStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v54
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = F_equal(m, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L3:
	;
	if v6 == int32(0) {
		v54 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 != v7 {
		v54 = v3
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v34 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v18 = v7
	v19 = v6
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v33 = v23
	v34 = v22
	goto L8
L12:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v54 = v3
	goto L1
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	if v41 == int32(0) {
		v54 = v3
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v47 != v48 {
		v54 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v52 = F_equal(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v54 = v52
	goto L1
}
func F_each_object_field_end(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+14)) = uint16(v4)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v13 == int32(1) {
		v16 = int32(_a_F_each_object_field_end_0)
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v19
		v21 = F_cstring_to_text(m, l1)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = base.I64_extend_i32_u(v21)
			if l2 == int32(0) {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
				if v36 == int32(1) {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v40 = F_cstring_to_text(m, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v40)
						v44 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v44)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v60 = F_heap_form_tuple(m, v55, v8+int32(16), v8+int32(14))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							F_tuplestore_puttuple(m, v62, v60)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v17
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_MemoryContextReset(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return int32(0)
								}
							}
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
					v50 = F_cstring_to_text_with_len(m, v46, v48-v46)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v50)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v60 = F_heap_form_tuple(m, v55, v8+int32(16), v8+int32(14))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							F_tuplestore_puttuple(m, v62, v60)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v17
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_MemoryContextReset(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(32)
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v29 != int32(1) {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
					if v36 == int32(1) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v40 = F_cstring_to_text(m, v39)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v40)
							v44 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v44)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v60 = F_heap_form_tuple(m, v55, v8+int32(16), v8+int32(14))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_tuplestore_puttuple(m, v62, v60)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v17
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_MemoryContextReset(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(32)
										return int32(0)
									}
								}
							}
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
						v50 = F_cstring_to_text_with_len(m, v46, v48-v46)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = base.I64_extend_i32_u(v50)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v60 = F_heap_form_tuple(m, v55, v8+int32(16), v8+int32(14))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_tuplestore_puttuple(m, v62, v60)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v17
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_MemoryContextReset(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(32)
										return int32(0)
									}
								}
							}
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
					v34 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v34)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v60 = F_heap_form_tuple(m, v55, v8+int32(16), v8+int32(14))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_tuplestore_puttuple(m, v62, v60)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_each_object_field_end[0])) = v17
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_MemoryContextReset(m, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(32)
		return int32(0)
	}
}
func F_ean2isn(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = int64(base.Ui64(l0) >> (uint(int64(1)) % 64))
	if base.Ui64(l0) <= base.Ui64(int64(19999999999999)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v158 = int32(2)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(v158)%32))+uint32(_c_F_ean2isn[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v162
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v102<<(uint(v158)%32))+uint32(_c_F_ean2isn[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v12 + int32(48)
	F_errmsg(m, int32(_a_F_ean2isn_0), v12+int32(32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L27
	} else {
		goto L37
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l0
	m.G0 = v12 + int32(80)
	return
L3:
	;
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)) = uint8(v18)
	v28 = v12 + int32(48) | int32(13)
	v29 = int32(0)
	v31 = v15
	goto L7
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L27
	} else {
		goto L33
	}
L6:
	;
	v64 = int32(3)
	v65 = int32(0)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+48)))
	v68 = v66 ^ int32(_a_F_ean2isn_1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+50)))
	if v68|(v69^int32(56)) == v65 {
		v102 = v64
		v103 = v65
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v34 = v28 - int32(1)
	v35 = int64(10)
	v36 = base.I64_div_u_s(v31, v35)
	v42 = base.I32_wrap_i64(v31-v36*v35) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v42)
	if base.Ui64(v31) < base.Ui64(v35) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(int32(11)) < base.Ui32(v29) {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	goto L8
L10:
	;
	v47 = v29 + int32(1)
	if v47 != int32(13) {
		v28 = v34
		v29 = v47
		v31 = v36
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L6
L12:
	;
	v53 = int32(12) - v29
	if v53 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	base.MemoryFill(m, v34+v29-int32(12), int32(48), v53)
	goto L6
L14:
	;
	if l2 == v102 {
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+50)))
	if v75^int32(55)|v68 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v102 = int32(5)
	v103 = v65
	goto L14
L17:
	;
	goto L18
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v82 == int32(809056057) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v102 = int32(4)
	v103 = v65
	goto L14
L20:
	;
	goto L21
L21:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+48)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+50)))
	if v86^int32(_a_F_ean2isn_1)|(v89^int32(57)) == int32(0) {
		v102 = v64
		v103 = v65
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v100 = base.B2i32(v82&int32(255) != int32(48))
	if v82&int32(255) != int32(48) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v101 = int32(2)
	goto L25
L24:
	;
	v101 = int32(6)
	goto L25
L25:
	;
	v102 = v101
	v103 = v100
	goto L14
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return
L28:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v103 == int32(0) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_ean2isn_2)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_c_F_ean2isn[0])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v12 + int32(48)
	F_errmsg(m, int32(_a_F_ean2isn_3), v12+int32(16))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_ean2isn_4), int32(419), int32(_a_F_ean2isn_5))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(_a_F_ean2isn_6)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg(m, int32(_a_F_ean2isn_7), v12)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_ean2isn_4), int32(429), int32(_a_F_ean2isn_5))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errfinish(m, int32(_a_F_ean2isn_4), int32(412), int32(_a_F_ean2isn_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L27
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ec_add_clause_to_derives_hash(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v29 int32
	_ = v29
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
	var v44 int32
	_ = v44
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
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
	var v309 int32
	_ = v309
	var v328 int64
	_ = v328
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v349 int64
	_ = v349
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int64
	_ = v364
	var v374 int64
	_ = v374
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
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
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v1002 int32
	_ = v1002
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int64
	_ = v1011
	var v1013 int64
	_ = v1013
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1382 int32
	_ = v1382
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int64
	_ = v1399
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int64
	_ = v1455
	var v1457 int64
	_ = v1457
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1483 int64
	_ = v1483
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1497 int32
	_ = v1497
	var v1514 int32
	_ = v1514
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1540 int32
	_ = v1540
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	var v1586 int32
	_ = v1586
	var v1596 int32
	_ = v1596
	var v1600 int32
	_ = v1600
	var v1605 int32
	_ = v1605
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+12)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v26
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = int32(0)
	goto L3
L2:
	;
	v29 = v24
	goto L3
L3:
	;
	v30 = base.B2i32(base.Ui32(v23) < base.Ui32(v29))
	if base.Ui32(v23) < base.Ui32(v29) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = v29
	goto L6
L5:
	;
	v31 = v23
	goto L6
L6:
	;
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v32 = v31
	goto L9
L8:
	;
	v32 = v23
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v32
	if base.Ui32(v23) < base.Ui32(v29) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v34 = v23
	goto L12
L11:
	;
	v34 = v29
	goto L12
L12:
	;
	if v29 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = v34
	goto L15
L14:
	;
	v36 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v36
	v38 = int32(12)
	v44 = int32(-1636608420)
	if v20&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v309 = base.B2i32(base.Ui32(v306) < base.Ui32(v307))
	goto L56
L17:
	;
	v276 = int32(14)
	v278 = v272 ^ v273 - base.I32_rotl(v272, v276)
	v282 = v278 ^ v271 - base.I32_rotl(v278, int32(11))
	v286 = v282 ^ v272 - base.I32_rotl(v282, int32(25))
	v290 = v286 ^ v278 - base.I32_rotl(v286, int32(16))
	v294 = v290 ^ v282 - base.I32_rotl(v290, int32(4))
	v298 = v294 ^ v286 - base.I32_rotl(v294, v276)
	goto L16
L18:
	;
	switch v198 - int32(1) {
	case 0:
		v264 = v189
		v265 = v190
		v266 = v194
		goto L45
	case 1:
		v257 = v189
		v258 = v190
		v259 = v194
		goto L46
	case 2:
		v250 = v189
		v251 = v190
		v252 = v194
		goto L47
	case 3:
		v244 = v190
		v245 = v194
		goto L48
	case 4:
		v240 = v190
		v241 = v194
		goto L49
	case 5:
		v234 = v190
		v235 = v194
		goto L50
	case 6:
		v228 = v190
		v229 = v194
		goto L51
	case 7:
		v223 = v194
		goto L52
	case 8:
		v218 = v194
		goto L53
	case 9:
		v213 = v194
		goto L54
	case 10:
		goto L55
	default:
		v271 = v189
		v272 = v190
		v273 = v194
		goto L17
	}
L19:
	;
	v153 = v20
	v154 = v38
	v155 = v44
	v156 = v44
	v157 = v44
	goto L42
L20:
	;
	goto L19
L21:
	;
	goto L22
L22:
	;
	goto L26
L24:
	;
	switch v96 - int32(1) {
	case 0:
		v150 = v87
		goto L31
	case 1:
		v145 = v87
		goto L32
	case 2:
		goto L33
	case 3:
		v138 = v88
		goto L34
	case 4:
		v135 = v88
		goto L35
	case 5:
		v130 = v88
		goto L36
	case 6:
		goto L37
	case 7:
		v121 = v92
		goto L38
	case 8:
		v116 = v92
		goto L39
	case 9:
		v111 = v92
		goto L40
	case 10:
		goto L41
	default:
		v271 = v87
		v272 = v88
		v273 = v92
		goto L17
	}
L26:
	;
	goto L27
L27:
	;
	v51 = v20
	v52 = v38
	v53 = v44
	v54 = v44
	v55 = v44
	goto L28
L28:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v58 = v57 + v54
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v62 = v61 + v55
	v64 = int32(4)
	v66 = v59 + v53 - v62 ^ base.I32_rotl(v62, v64)
	v70 = v58 - v66 ^ base.I32_rotl(v66, int32(6))
	v71 = v62 + v58
	v72 = v66 + v71
	v73 = v70 + v72
	v77 = v71 - v70 ^ base.I32_rotl(v70, int32(8))
	v81 = v72 - v77 ^ base.I32_rotl(v77, int32(16))
	v85 = v73 - v81 ^ base.I32_rotl(v81, int32(19))
	v86 = v77 + v73
	v87 = v81 + v86
	v88 = v85 + v87
	v92 = v86 - v85 ^ base.I32_rotl(v85, v64)
	v93 = int32(12)
	v94 = v51 + v93
	v96 = v52 - v93
	if base.Ui32(int32(11)) < base.Ui32(v96) {
		v51 = v94
		v52 = v96
		v53 = v87
		v54 = v88
		v55 = v92
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L24
L30:
	;
	goto L29
L31:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
	v271 = v150 + v151
	v272 = v88
	v273 = v92
	goto L17
L32:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v150 = v146<<(uint(int32(8))%32) + v145
	goto L31
L33:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
	v145 = v141<<(uint(int32(16))%32) + v87
	goto L32
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v271 = v139 + v87
	v272 = v138
	v273 = v92
	goto L17
L35:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	v138 = v135 + v136
	goto L34
L36:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+5)))
	v135 = v131<<(uint(int32(8))%32) + v130
	goto L35
L37:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+6)))
	v130 = v126<<(uint(int32(16))%32) + v88
	goto L36
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v271 = v122 + v87
	v272 = v124 + v88
	v273 = v121
	goto L17
L39:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)))
	v121 = v117<<(uint(int32(8))%32) + v116
	goto L38
L40:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+9)))
	v116 = v112<<(uint(int32(16))%32) + v111
	goto L39
L41:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+10)))
	v111 = v107<<(uint(int32(24))%32) + v92
	goto L40
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v160 = v159 + v156
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v164 = v163 + v157
	v166 = int32(4)
	v168 = v161 + v155 - v164 ^ base.I32_rotl(v164, v166)
	v172 = v160 - v168 ^ base.I32_rotl(v168, int32(6))
	v173 = v164 + v160
	v174 = v168 + v173
	v175 = v172 + v174
	v179 = v173 - v172 ^ base.I32_rotl(v172, int32(8))
	v183 = v174 - v179 ^ base.I32_rotl(v179, int32(16))
	v187 = v175 - v183 ^ base.I32_rotl(v183, int32(19))
	v188 = v179 + v175
	v189 = v183 + v188
	v190 = v187 + v189
	v194 = v188 - v187 ^ base.I32_rotl(v187, v166)
	v195 = int32(12)
	v196 = v153 + v195
	v198 = v154 - v195
	if base.Ui32(int32(11)) < base.Ui32(v198) {
		v153 = v196
		v154 = v198
		v155 = v189
		v156 = v190
		v157 = v194
		goto L42
	} else {
		goto L44
	}
L43:
	;
	goto L18
L44:
	;
	goto L43
L45:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v271 = v264 + v267
	v272 = v265
	v273 = v266
	goto L17
L46:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	v264 = v260<<(uint(int32(8))%32) + v257
	v265 = v258
	v266 = v259
	goto L45
L47:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+2)))
	v257 = v253<<(uint(int32(16))%32) + v250
	v258 = v251
	v259 = v252
	goto L46
L48:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+3)))
	v250 = v246<<(uint(int32(24))%32) + v189
	v251 = v244
	v252 = v245
	goto L47
L49:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+4)))
	v244 = v240 + v242
	v245 = v241
	goto L48
L50:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+5)))
	v240 = v236<<(uint(int32(8))%32) + v234
	v241 = v235
	goto L49
L51:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+6)))
	v234 = v230<<(uint(int32(16))%32) + v228
	v235 = v229
	goto L50
L52:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+7)))
	v228 = v224<<(uint(int32(24))%32) + v190
	v229 = v223
	goto L51
L53:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+8)))
	v223 = v219<<(uint(int32(8))%32) + v218
	goto L52
L54:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+9)))
	v218 = v214<<(uint(int32(16))%32) + v213
	goto L53
L55:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+10)))
	v213 = v209<<(uint(int32(24))%32) + v194
	goto L54
L56:
	;
	if v309 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L75
	} else {
		goto L267
	}
L58:
	;
	goto L57
L59:
	;
	v1586 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1586
	v309 = v1586
	goto L56
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1559 = m.ExcPending
	if v1559 != 0 {
		goto L75
	} else {
		goto L264
	}
L61:
	;
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if v328 == int64(4294967296) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1064 = v1063 & (v298 ^ v290 - base.I32_rotl(v298, int32(24)))
	v1067 = v1062 + v1064*int32(20)
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1067)))
	if v1068 != 0 {
		goto L187
	} else {
		goto L188
	}
L64:
	;
	v331 = int32(0)
	v333 = int64(2)
	v335 = v328 << (uint(int64(1)) % 64)
	if base.Ui64(v335) <= base.Ui64(v333) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v309 = int32(1)
	goto L56
L66:
	;
	v338 = v333
	goto L68
L67:
	;
	v338 = v335
	goto L68
L68:
	;
	v339 = int64(1)
	if v338&(v338-v339) == int64(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v349 = v338
	goto L71
L70:
	;
	v349 = v339 << (uint(int64(64)-base.I64_clz(v338)) % 64)
	goto L71
L71:
	;
	if base.Ui64(v349*int64(20)) < base.Ui64(int64(2147483647)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v361 = F_MemoryContextAllocExtended(m, v356, base.I32_wrap_i64(v349)*int32(20), int32(5))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	goto L58
L75:
	;
	return
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v361
	v364 = int64(1)
	if v349&(v349-v364) == int64(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v374 = v349
	goto L79
L78:
	;
	v374 = v364 << (uint(int64(64)-base.I64_clz(v349)) % 64)
	goto L79
L79:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v374*int64(20)) {
		goto L58
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = base.I32_wrap_i64(v374) - int32(1)
	if v374 == int64(4294967296) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v391 = int32(-85899346)
	goto L83
L82:
	;
	v391 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v374), float64(0.9)))
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v391
	if v355 != int64(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v395 = v331
	goto L88
L85:
	;
	goto L86
L86:
	;
	F_pfree(m, v354)
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L75
	} else {
		goto L185
	}
L87:
	;
	v694 = v693
	v702 = v331
	goto L133
L88:
	;
	v414 = v354 + v395*int32(20)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	if v415 != int32(1) {
		v693 = v395
		goto L87
	} else {
		goto L90
	}
L89:
	;
	v693 = int32(0)
	goto L87
L90:
	;
	v419 = v414 + int32(4)
	v420 = int32(12)
	v426 = int32(-1636608420)
	if v419&int32(3) != 0 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if (v680^v672-base.I32_rotl(v680, int32(24)))&v685 == v395 {
		v693 = v395
		goto L87
	} else {
		goto L131
	}
L92:
	;
	v658 = int32(14)
	v660 = v654 ^ v655 - base.I32_rotl(v654, v658)
	v664 = v660 ^ v653 - base.I32_rotl(v660, int32(11))
	v668 = v664 ^ v654 - base.I32_rotl(v664, int32(25))
	v672 = v668 ^ v660 - base.I32_rotl(v668, int32(16))
	v676 = v672 ^ v664 - base.I32_rotl(v672, int32(4))
	v680 = v676 ^ v668 - base.I32_rotl(v676, v658)
	goto L91
L93:
	;
	switch v580 - int32(1) {
	case 0:
		v646 = v571
		v647 = v572
		v648 = v576
		goto L120
	case 1:
		v639 = v571
		v640 = v572
		v641 = v576
		goto L121
	case 2:
		v632 = v571
		v633 = v572
		v634 = v576
		goto L122
	case 3:
		v626 = v572
		v627 = v576
		goto L123
	case 4:
		v622 = v572
		v623 = v576
		goto L124
	case 5:
		v616 = v572
		v617 = v576
		goto L125
	case 6:
		v610 = v572
		v611 = v576
		goto L126
	case 7:
		v605 = v576
		goto L127
	case 8:
		v600 = v576
		goto L128
	case 9:
		v595 = v576
		goto L129
	case 10:
		goto L130
	default:
		v653 = v571
		v654 = v572
		v655 = v576
		goto L92
	}
L94:
	;
	v535 = v419
	v536 = v420
	v537 = v426
	v538 = v426
	v539 = v426
	goto L117
L95:
	;
	goto L94
L96:
	;
	goto L97
L97:
	;
	goto L101
L99:
	;
	switch v478 - int32(1) {
	case 0:
		v532 = v469
		goto L106
	case 1:
		v527 = v469
		goto L107
	case 2:
		goto L108
	case 3:
		v520 = v470
		goto L109
	case 4:
		v517 = v470
		goto L110
	case 5:
		v512 = v470
		goto L111
	case 6:
		goto L112
	case 7:
		v503 = v474
		goto L113
	case 8:
		v498 = v474
		goto L114
	case 9:
		v493 = v474
		goto L115
	case 10:
		goto L116
	default:
		v653 = v469
		v654 = v470
		v655 = v474
		goto L92
	}
L101:
	;
	goto L102
L102:
	;
	v433 = v419
	v434 = v420
	v435 = v426
	v436 = v426
	v437 = v426
	goto L103
L103:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v440 = v439 + v436
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v433)+8))
	v444 = v443 + v437
	v446 = int32(4)
	v448 = v441 + v435 - v444 ^ base.I32_rotl(v444, v446)
	v452 = v440 - v448 ^ base.I32_rotl(v448, int32(6))
	v453 = v444 + v440
	v454 = v448 + v453
	v455 = v452 + v454
	v459 = v453 - v452 ^ base.I32_rotl(v452, int32(8))
	v463 = v454 - v459 ^ base.I32_rotl(v459, int32(16))
	v467 = v455 - v463 ^ base.I32_rotl(v463, int32(19))
	v468 = v459 + v455
	v469 = v463 + v468
	v470 = v467 + v469
	v474 = v468 - v467 ^ base.I32_rotl(v467, v446)
	v475 = int32(12)
	v476 = v433 + v475
	v478 = v434 - v475
	if base.Ui32(int32(11)) < base.Ui32(v478) {
		v433 = v476
		v434 = v478
		v435 = v469
		v436 = v470
		v437 = v474
		goto L103
	} else {
		goto L105
	}
L104:
	;
	goto L99
L105:
	;
	goto L104
L106:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	v653 = v532 + v533
	v654 = v470
	v655 = v474
	goto L92
L107:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+1)))
	v532 = v528<<(uint(int32(8))%32) + v527
	goto L106
L108:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+2)))
	v527 = v523<<(uint(int32(16))%32) + v469
	goto L107
L109:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	v653 = v521 + v469
	v654 = v520
	v655 = v474
	goto L92
L110:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+4)))
	v520 = v517 + v518
	goto L109
L111:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+5)))
	v517 = v513<<(uint(int32(8))%32) + v512
	goto L110
L112:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+6)))
	v512 = v508<<(uint(int32(16))%32) + v470
	goto L111
L113:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	v653 = v504 + v469
	v654 = v506 + v470
	v655 = v503
	goto L92
L114:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+8)))
	v503 = v499<<(uint(int32(8))%32) + v498
	goto L113
L115:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+9)))
	v498 = v494<<(uint(int32(16))%32) + v493
	goto L114
L116:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476)+10)))
	v493 = v489<<(uint(int32(24))%32) + v474
	goto L115
L117:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v535)+4))
	v542 = v541 + v538
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v535)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v535)+8))
	v546 = v545 + v539
	v548 = int32(4)
	v550 = v543 + v537 - v546 ^ base.I32_rotl(v546, v548)
	v554 = v542 - v550 ^ base.I32_rotl(v550, int32(6))
	v555 = v546 + v542
	v556 = v550 + v555
	v557 = v554 + v556
	v561 = v555 - v554 ^ base.I32_rotl(v554, int32(8))
	v565 = v556 - v561 ^ base.I32_rotl(v561, int32(16))
	v569 = v557 - v565 ^ base.I32_rotl(v565, int32(19))
	v570 = v561 + v557
	v571 = v565 + v570
	v572 = v569 + v571
	v576 = v570 - v569 ^ base.I32_rotl(v569, v548)
	v577 = int32(12)
	v578 = v535 + v577
	v580 = v536 - v577
	if base.Ui32(int32(11)) < base.Ui32(v580) {
		v535 = v578
		v536 = v580
		v537 = v571
		v538 = v572
		v539 = v576
		goto L117
	} else {
		goto L119
	}
L118:
	;
	goto L93
L119:
	;
	goto L118
L120:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v653 = v646 + v649
	v654 = v647
	v655 = v648
	goto L92
L121:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+1)))
	v646 = v642<<(uint(int32(8))%32) + v639
	v647 = v640
	v648 = v641
	goto L120
L122:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+2)))
	v639 = v635<<(uint(int32(16))%32) + v632
	v640 = v633
	v641 = v634
	goto L121
L123:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+3)))
	v632 = v628<<(uint(int32(24))%32) + v571
	v633 = v626
	v634 = v627
	goto L122
L124:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+4)))
	v626 = v622 + v624
	v627 = v623
	goto L123
L125:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+5)))
	v622 = v618<<(uint(int32(8))%32) + v616
	v623 = v617
	goto L124
L126:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+6)))
	v616 = v612<<(uint(int32(16))%32) + v610
	v617 = v611
	goto L125
L127:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+7)))
	v610 = v606<<(uint(int32(24))%32) + v572
	v611 = v605
	goto L126
L128:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+8)))
	v605 = v601<<(uint(int32(8))%32) + v600
	goto L127
L129:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+9)))
	v600 = v596<<(uint(int32(16))%32) + v595
	goto L128
L130:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+10)))
	v595 = v591<<(uint(int32(24))%32) + v576
	goto L129
L131:
	;
	v689 = v395 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v689)) < base.Ui64(v355) {
		v395 = v689
		goto L88
	} else {
		goto L132
	}
L132:
	;
	goto L89
L133:
	;
	v713 = v354 + v694*int32(20)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)))
	if v714 == int32(1) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L86
L135:
	;
	v718 = v713 + int32(4)
	v719 = int32(12)
	v725 = int32(-1636608420)
	if v718&int32(3) != 0 {
		goto L142
	} else {
		goto L143
	}
L136:
	;
	goto L137
L137:
	;
	v1033 = v694 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1033)) < base.Ui64(v355) {
		goto L181
	} else {
		goto L182
	}
L138:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v988 = v979 ^ v971 - base.I32_rotl(v979, int32(24))
	goto L178
L139:
	;
	v957 = int32(14)
	v959 = v953 ^ v954 - base.I32_rotl(v953, v957)
	v963 = v959 ^ v952 - base.I32_rotl(v959, int32(11))
	v967 = v963 ^ v953 - base.I32_rotl(v963, int32(25))
	v971 = v967 ^ v959 - base.I32_rotl(v967, int32(16))
	v975 = v971 ^ v963 - base.I32_rotl(v971, int32(4))
	v979 = v975 ^ v967 - base.I32_rotl(v975, v957)
	goto L138
L140:
	;
	switch v879 - int32(1) {
	case 0:
		v945 = v870
		v946 = v871
		v947 = v875
		goto L167
	case 1:
		v938 = v870
		v939 = v871
		v940 = v875
		goto L168
	case 2:
		v931 = v870
		v932 = v871
		v933 = v875
		goto L169
	case 3:
		v925 = v871
		v926 = v875
		goto L170
	case 4:
		v921 = v871
		v922 = v875
		goto L171
	case 5:
		v915 = v871
		v916 = v875
		goto L172
	case 6:
		v909 = v871
		v910 = v875
		goto L173
	case 7:
		v904 = v875
		goto L174
	case 8:
		v899 = v875
		goto L175
	case 9:
		v894 = v875
		goto L176
	case 10:
		goto L177
	default:
		v952 = v870
		v953 = v871
		v954 = v875
		goto L139
	}
L141:
	;
	v834 = v718
	v835 = v719
	v836 = v725
	v837 = v725
	v838 = v725
	goto L164
L142:
	;
	goto L141
L143:
	;
	goto L144
L144:
	;
	goto L148
L146:
	;
	switch v777 - int32(1) {
	case 0:
		v831 = v768
		goto L153
	case 1:
		v826 = v768
		goto L154
	case 2:
		goto L155
	case 3:
		v819 = v769
		goto L156
	case 4:
		v816 = v769
		goto L157
	case 5:
		v811 = v769
		goto L158
	case 6:
		goto L159
	case 7:
		v802 = v773
		goto L160
	case 8:
		v797 = v773
		goto L161
	case 9:
		v792 = v773
		goto L162
	case 10:
		goto L163
	default:
		v952 = v768
		v953 = v769
		v954 = v773
		goto L139
	}
L148:
	;
	goto L149
L149:
	;
	v732 = v718
	v733 = v719
	v734 = v725
	v735 = v725
	v736 = v725
	goto L150
L150:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v732)+4))
	v739 = v738 + v735
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v732)+8))
	v743 = v742 + v736
	v745 = int32(4)
	v747 = v740 + v734 - v743 ^ base.I32_rotl(v743, v745)
	v751 = v739 - v747 ^ base.I32_rotl(v747, int32(6))
	v752 = v743 + v739
	v753 = v747 + v752
	v754 = v751 + v753
	v758 = v752 - v751 ^ base.I32_rotl(v751, int32(8))
	v762 = v753 - v758 ^ base.I32_rotl(v758, int32(16))
	v766 = v754 - v762 ^ base.I32_rotl(v762, int32(19))
	v767 = v758 + v754
	v768 = v762 + v767
	v769 = v766 + v768
	v773 = v767 - v766 ^ base.I32_rotl(v766, v745)
	v774 = int32(12)
	v775 = v732 + v774
	v777 = v733 - v774
	if base.Ui32(int32(11)) < base.Ui32(v777) {
		v732 = v775
		v733 = v777
		v734 = v768
		v735 = v769
		v736 = v773
		goto L150
	} else {
		goto L152
	}
L151:
	;
	goto L146
L152:
	;
	goto L151
L153:
	;
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775))))
	v952 = v831 + v832
	v953 = v769
	v954 = v773
	goto L139
L154:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+1)))
	v831 = v827<<(uint(int32(8))%32) + v826
	goto L153
L155:
	;
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+2)))
	v826 = v822<<(uint(int32(16))%32) + v768
	goto L154
L156:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v952 = v820 + v768
	v953 = v819
	v954 = v773
	goto L139
L157:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+4)))
	v819 = v816 + v817
	goto L156
L158:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+5)))
	v816 = v812<<(uint(int32(8))%32) + v811
	goto L157
L159:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+6)))
	v811 = v807<<(uint(int32(16))%32) + v769
	goto L158
L160:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v775)+4))
	v952 = v803 + v768
	v953 = v805 + v769
	v954 = v802
	goto L139
L161:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+8)))
	v802 = v798<<(uint(int32(8))%32) + v797
	goto L160
L162:
	;
	v793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+9)))
	v797 = v793<<(uint(int32(16))%32) + v792
	goto L161
L163:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v775)+10)))
	v792 = v788<<(uint(int32(24))%32) + v773
	goto L162
L164:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v841 = v840 + v837
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v834)))
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v834)+8))
	v845 = v844 + v838
	v847 = int32(4)
	v849 = v842 + v836 - v845 ^ base.I32_rotl(v845, v847)
	v853 = v841 - v849 ^ base.I32_rotl(v849, int32(6))
	v854 = v845 + v841
	v855 = v849 + v854
	v856 = v853 + v855
	v860 = v854 - v853 ^ base.I32_rotl(v853, int32(8))
	v864 = v855 - v860 ^ base.I32_rotl(v860, int32(16))
	v868 = v856 - v864 ^ base.I32_rotl(v864, int32(19))
	v869 = v860 + v856
	v870 = v864 + v869
	v871 = v868 + v870
	v875 = v869 - v868 ^ base.I32_rotl(v868, v847)
	v876 = int32(12)
	v877 = v834 + v876
	v879 = v835 - v876
	if base.Ui32(int32(11)) < base.Ui32(v879) {
		v834 = v877
		v835 = v879
		v836 = v870
		v837 = v871
		v838 = v875
		goto L164
	} else {
		goto L166
	}
L165:
	;
	goto L140
L166:
	;
	goto L165
L167:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877))))
	v952 = v945 + v948
	v953 = v946
	v954 = v947
	goto L139
L168:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)))
	v945 = v941<<(uint(int32(8))%32) + v938
	v946 = v939
	v947 = v940
	goto L167
L169:
	;
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+2)))
	v938 = v934<<(uint(int32(16))%32) + v931
	v939 = v932
	v940 = v933
	goto L168
L170:
	;
	v927 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+3)))
	v931 = v927<<(uint(int32(24))%32) + v870
	v932 = v925
	v933 = v926
	goto L169
L171:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+4)))
	v925 = v921 + v923
	v926 = v922
	goto L170
L172:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+5)))
	v921 = v917<<(uint(int32(8))%32) + v915
	v922 = v916
	goto L171
L173:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+6)))
	v915 = v911<<(uint(int32(16))%32) + v909
	v916 = v910
	goto L172
L174:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+7)))
	v909 = v905<<(uint(int32(24))%32) + v871
	v910 = v904
	goto L173
L175:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+8)))
	v904 = v900<<(uint(int32(8))%32) + v899
	goto L174
L176:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+9)))
	v899 = v895<<(uint(int32(16))%32) + v894
	goto L175
L177:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+10)))
	v894 = v890<<(uint(int32(24))%32) + v875
	goto L176
L178:
	;
	v1002 = v988 & v984
	v1007 = v361 + v1002*int32(20)
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)))
	if v1008 != 0 {
		v988 = v1002 + int32(1)
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v713)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+16)) = v1009
	v1011 = *(*int64)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1007)+8)) = v1011
	v1013 = *(*int64)(unsafe.Add(mBase, uint32(v713)))
	*(*int64)(unsafe.Add(mBase, uint32(v1007))) = v1013
	goto L137
L180:
	;
	goto L179
L181:
	;
	v1037 = v1033
	goto L183
L182:
	;
	v1037 = int32(0)
	goto L183
L183:
	;
	v1039 = v702 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1039)) < base.Ui64(v355) {
		v694 = v1037
		v702 = v1039
		goto L133
	} else {
		goto L184
	}
L184:
	;
	goto L134
L185:
	;
	goto L65
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1540)+16)) = l1
	m.G0 = v20 + int32(16)
	return
L187:
	;
	v1073 = v1064
	v1074 = int32(0)
	v1075 = v1067
	goto L191
L188:
	;
	v1514 = v1067
	goto L189
L189:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1527 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v1526 + v1527
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+12)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+8)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v1514)+4)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v1514))) = v1527
	v1540 = v1514
	goto L186
L190:
	;
	v1514 = v1497
	goto L189
L191:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+4))
	if v1087 != v305 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v1497 = v1490
	goto L190
L193:
	;
	v1094 = v1075 + int32(4)
	v1095 = int32(12)
	v1101 = int32(-1636608420)
	if v1094&int32(3) != 0 {
		goto L201
	} else {
		goto L202
	}
L194:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+8))
	if v1089 != v304 {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1075)+12))
	if v1091 == v303 {
		v1540 = v1075
		goto L186
	} else {
		goto L196
	}
L196:
	;
	goto L193
L197:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1361 = (v1355 ^ v1347 - base.I32_rotl(v1355, int32(24))) & v1360
	if base.Ui32(v1073) < base.Ui32(v1361) {
		goto L237
	} else {
		goto L238
	}
L198:
	;
	v1333 = int32(14)
	v1335 = v1329 ^ v1330 - base.I32_rotl(v1329, v1333)
	v1339 = v1335 ^ v1328 - base.I32_rotl(v1335, int32(11))
	v1343 = v1339 ^ v1329 - base.I32_rotl(v1339, int32(25))
	v1347 = v1343 ^ v1335 - base.I32_rotl(v1343, int32(16))
	v1351 = v1347 ^ v1339 - base.I32_rotl(v1347, int32(4))
	v1355 = v1351 ^ v1343 - base.I32_rotl(v1351, v1333)
	goto L197
L199:
	;
	switch v1255 - int32(1) {
	case 0:
		v1321 = v1246
		v1322 = v1247
		v1323 = v1251
		goto L226
	case 1:
		v1314 = v1246
		v1315 = v1247
		v1316 = v1251
		goto L227
	case 2:
		v1307 = v1246
		v1308 = v1247
		v1309 = v1251
		goto L228
	case 3:
		v1301 = v1247
		v1302 = v1251
		goto L229
	case 4:
		v1297 = v1247
		v1298 = v1251
		goto L230
	case 5:
		v1291 = v1247
		v1292 = v1251
		goto L231
	case 6:
		v1285 = v1247
		v1286 = v1251
		goto L232
	case 7:
		v1280 = v1251
		goto L233
	case 8:
		v1275 = v1251
		goto L234
	case 9:
		v1270 = v1251
		goto L235
	case 10:
		goto L236
	default:
		v1328 = v1246
		v1329 = v1247
		v1330 = v1251
		goto L198
	}
L200:
	;
	v1210 = v1094
	v1211 = v1095
	v1212 = v1101
	v1213 = v1101
	v1214 = v1101
	goto L223
L201:
	;
	goto L200
L202:
	;
	goto L203
L203:
	;
	goto L207
L205:
	;
	switch v1153 - int32(1) {
	case 0:
		v1207 = v1144
		goto L212
	case 1:
		v1202 = v1144
		goto L213
	case 2:
		goto L214
	case 3:
		v1195 = v1145
		goto L215
	case 4:
		v1192 = v1145
		goto L216
	case 5:
		v1187 = v1145
		goto L217
	case 6:
		goto L218
	case 7:
		v1178 = v1149
		goto L219
	case 8:
		v1173 = v1149
		goto L220
	case 9:
		v1168 = v1149
		goto L221
	case 10:
		goto L222
	default:
		v1328 = v1144
		v1329 = v1145
		v1330 = v1149
		goto L198
	}
L207:
	;
	goto L208
L208:
	;
	v1108 = v1094
	v1109 = v1095
	v1110 = v1101
	v1111 = v1101
	v1112 = v1101
	goto L209
L209:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+4))
	v1115 = v1114 + v1111
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1108)))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1108)+8))
	v1119 = v1118 + v1112
	v1121 = int32(4)
	v1123 = v1116 + v1110 - v1119 ^ base.I32_rotl(v1119, v1121)
	v1127 = v1115 - v1123 ^ base.I32_rotl(v1123, int32(6))
	v1128 = v1119 + v1115
	v1129 = v1123 + v1128
	v1130 = v1127 + v1129
	v1134 = v1128 - v1127 ^ base.I32_rotl(v1127, int32(8))
	v1138 = v1129 - v1134 ^ base.I32_rotl(v1134, int32(16))
	v1142 = v1130 - v1138 ^ base.I32_rotl(v1138, int32(19))
	v1143 = v1134 + v1130
	v1144 = v1138 + v1143
	v1145 = v1142 + v1144
	v1149 = v1143 - v1142 ^ base.I32_rotl(v1142, v1121)
	v1150 = int32(12)
	v1151 = v1108 + v1150
	v1153 = v1109 - v1150
	if base.Ui32(int32(11)) < base.Ui32(v1153) {
		v1108 = v1151
		v1109 = v1153
		v1110 = v1144
		v1111 = v1145
		v1112 = v1149
		goto L209
	} else {
		goto L211
	}
L210:
	;
	goto L205
L211:
	;
	goto L210
L212:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151))))
	v1328 = v1207 + v1208
	v1329 = v1145
	v1330 = v1149
	goto L198
L213:
	;
	v1203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+1)))
	v1207 = v1203<<(uint(int32(8))%32) + v1202
	goto L212
L214:
	;
	v1198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+2)))
	v1202 = v1198<<(uint(int32(16))%32) + v1144
	goto L213
L215:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1151)))
	v1328 = v1196 + v1144
	v1329 = v1195
	v1330 = v1149
	goto L198
L216:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+4)))
	v1195 = v1192 + v1193
	goto L215
L217:
	;
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+5)))
	v1192 = v1188<<(uint(int32(8))%32) + v1187
	goto L216
L218:
	;
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+6)))
	v1187 = v1183<<(uint(int32(16))%32) + v1145
	goto L217
L219:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1151)))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	v1328 = v1179 + v1144
	v1329 = v1181 + v1145
	v1330 = v1178
	goto L198
L220:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+8)))
	v1178 = v1174<<(uint(int32(8))%32) + v1173
	goto L219
L221:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+9)))
	v1173 = v1169<<(uint(int32(16))%32) + v1168
	goto L220
L222:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151)+10)))
	v1168 = v1164<<(uint(int32(24))%32) + v1149
	goto L221
L223:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	v1217 = v1216 + v1213
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1210)))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+8))
	v1221 = v1220 + v1214
	v1223 = int32(4)
	v1225 = v1218 + v1212 - v1221 ^ base.I32_rotl(v1221, v1223)
	v1229 = v1217 - v1225 ^ base.I32_rotl(v1225, int32(6))
	v1230 = v1221 + v1217
	v1231 = v1225 + v1230
	v1232 = v1229 + v1231
	v1236 = v1230 - v1229 ^ base.I32_rotl(v1229, int32(8))
	v1240 = v1231 - v1236 ^ base.I32_rotl(v1236, int32(16))
	v1244 = v1232 - v1240 ^ base.I32_rotl(v1240, int32(19))
	v1245 = v1236 + v1232
	v1246 = v1240 + v1245
	v1247 = v1244 + v1246
	v1251 = v1245 - v1244 ^ base.I32_rotl(v1244, v1223)
	v1252 = int32(12)
	v1253 = v1210 + v1252
	v1255 = v1211 - v1252
	if base.Ui32(int32(11)) < base.Ui32(v1255) {
		v1210 = v1253
		v1211 = v1255
		v1212 = v1246
		v1213 = v1247
		v1214 = v1251
		goto L223
	} else {
		goto L225
	}
L224:
	;
	goto L199
L225:
	;
	goto L224
L226:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	v1328 = v1321 + v1324
	v1329 = v1322
	v1330 = v1323
	goto L198
L227:
	;
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+1)))
	v1321 = v1317<<(uint(int32(8))%32) + v1314
	v1322 = v1315
	v1323 = v1316
	goto L226
L228:
	;
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+2)))
	v1314 = v1310<<(uint(int32(16))%32) + v1307
	v1315 = v1308
	v1316 = v1309
	goto L227
L229:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+3)))
	v1307 = v1303<<(uint(int32(24))%32) + v1246
	v1308 = v1301
	v1309 = v1302
	goto L228
L230:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+4)))
	v1301 = v1297 + v1299
	v1302 = v1298
	goto L229
L231:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+5)))
	v1297 = v1293<<(uint(int32(8))%32) + v1291
	v1298 = v1292
	goto L230
L232:
	;
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+6)))
	v1291 = v1287<<(uint(int32(16))%32) + v1285
	v1292 = v1286
	goto L231
L233:
	;
	v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+7)))
	v1285 = v1281<<(uint(int32(24))%32) + v1247
	v1286 = v1280
	goto L232
L234:
	;
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+8)))
	v1280 = v1276<<(uint(int32(8))%32) + v1275
	goto L233
L235:
	;
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+9)))
	v1275 = v1271<<(uint(int32(16))%32) + v1270
	goto L234
L236:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253)+10)))
	v1270 = v1266<<(uint(int32(24))%32) + v1251
	goto L235
L237:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v1365 = v1073 + v1363
	goto L239
L238:
	;
	v1365 = v1073
	goto L239
L239:
	;
	v1368 = v1360 & (v1073 + int32(1))
	if base.Ui32(v1365-v1361) < base.Ui32(v1074) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1373 = v1062 + v1368*int32(20)
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1373)))
	if v1374 != 0 {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L242
L242:
	;
	v1478 = v1074 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1478) {
		goto L259
	} else {
		goto L260
	}
L243:
	;
	v1376 = v1368
	v1382 = int32(0)
	goto L246
L244:
	;
	v1411 = v1368
	v1415 = v1373
	goto L245
L245:
	;
	if v1411 != v1073 {
		goto L253
	} else {
		goto L254
	}
L246:
	;
	v1394 = v1382 + int32(1)
	if int32(151) <= v1394 {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	v1411 = v1406
	v1415 = v1409
	goto L245
L248:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1399 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1397), base.F64_convert_i64_u(v1399)), float64(0.1)) != 0 {
		goto L59
	} else {
		goto L251
	}
L249:
	;
	goto L250
L250:
	;
	v1406 = (v1376 + int32(1)) & v1360
	v1409 = v1062 + v1406*int32(20)
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1409)))
	if v1410 != 0 {
		v1376 = v1406
		v1382 = v1394
		goto L246
	} else {
		goto L252
	}
L251:
	;
	goto L250
L252:
	;
	goto L247
L253:
	;
	v1429 = v1411
	v1433 = v1415
	goto L256
L254:
	;
	goto L255
L255:
	;
	v1497 = v1075
	goto L190
L256:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v1449 = v1446 & (v1429 - int32(1))
	v1452 = v1062 + v1449*int32(20)
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(v1452)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1433)+16)) = v1453
	v1455 = *(*int64)(unsafe.Add(mBase, uint32(v1452)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1433)+8)) = v1455
	v1457 = *(*int64)(unsafe.Add(mBase, uint32(v1452)))
	*(*int64)(unsafe.Add(mBase, uint32(v1433))) = v1457
	if v1449 != v1073 {
		v1429 = v1449
		v1433 = v1452
		goto L256
	} else {
		goto L258
	}
L257:
	;
	goto L255
L258:
	;
	goto L257
L259:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v1483 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1481), base.F64_convert_i64_u(v1483)), float64(0.1)) != 0 {
		goto L59
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1490 = v1062 + v1368*int32(20)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1490)))
	if v1491 != 0 {
		v1073 = v1368
		v1074 = v1478
		v1075 = v1490
		goto L191
	} else {
		goto L263
	}
L262:
	;
	goto L261
L263:
	;
	goto L192
L264:
	;
	F_errmsg_internal(m, int32(_a_F_ec_add_clause_to_derives_hash_0), int32(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L75
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_ec_add_clause_to_derives_hash_1), int32(635), int32(_a_F_ec_add_clause_to_derives_hash_2))
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L75
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errmsg_internal(m, int32(_a_F_ec_add_clause_to_derives_hash_3), int32(0))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L75
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_ec_add_clause_to_derives_hash_1), int32(332), int32(_a_F_ec_add_clause_to_derives_hash_4))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L75
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_elem_contained_by_multirange(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
				v32 = F_multirange_contains_elem_internal(m, v31, v13, v11)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v32)
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(_a_F_elem_contained_by_multirange_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(_a_F_elem_contained_by_multirange_1), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_elem_contained_by_multirange_2), int32(561), int32(_a_F_elem_contained_by_multirange_3))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
						v32 = F_multirange_contains_elem_internal(m, v31, v13, v11)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v32)
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(_a_F_elem_contained_by_multirange_0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(_a_F_elem_contained_by_multirange_1), v9)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_elem_contained_by_multirange_2), int32(561), int32(_a_F_elem_contained_by_multirange_3))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+296))
					v32 = F_multirange_contains_elem_internal(m, v31, v13, v11)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v32)
					}
				}
			}
		}
	}
}
func F_element_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v9 = F_FunctionCall2Coll(m, l2+int32(104), v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return base.I32_wrap_i64(v9)
	}
}
func F_elements_worker(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			if v18 == int32(1) {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
				if v24 == int32(18) {
					v27 = int32(16)
				} else {
					v27 = int32(0)
				}
				if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v34 = int32(4)
				} else {
					v34 = v27
				}
				v47 = v34
			} else {
				v35 = int32(1)
				if v18&v35 != 0 {
					v47 = int32(base.Ui32(v18)>>(uint(v35)%32)) - v35
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = v11 + int32(12)
			v50 = int32(1)
			if v18&v50 != 0 {
				v54 = v50
			} else {
				v54 = int32(4)
			}
			v57 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker[0]))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
			v59 = F_makeJsonLexContextCstringLen(m, v49, v16+v54, v47, v58, v3)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				v62 = F_palloc0(m, int32(32))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v65 = F_palloc0(m, int32(40))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_InitMaterializedSRF(m, l0, int32(3))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v71
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v65)+36)) = int32(1491)
							*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(1492)
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
							*(*int32)(unsafe.Add(mBase, uint32(v65)+32)) = int32(1493)
							*(*int32)(unsafe.Add(mBase, uint32(v65)+28)) = int32(1494)
							v84 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+25)) = uint8(v84)
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+24)) = uint8(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v49
							v90 = *(*int32)(unsafe.Add(mBase, _c_F_elements_worker[1]))
							v95 = F_AllocSetContextCreateInternal(m, v90, int32(_a_F_elements_worker_0), v84, int32(_a_F_elements_worker_1), int32(_a_F_elements_worker_2))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v95
								v98 = F_pg_parse_json(m, v49, v65)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									if v98 != 0 {
										F_json_errsave_error(m, v98, v49, int32(0))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
											F_MemoryContextDelete(m, v103)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												F_freeJsonLexContext(m, v11+int32(12))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
										F_MemoryContextDelete(m, v103)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											F_freeJsonLexContext(m, v11+int32(12))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
												m.G0 = v11 + int32(80)
												return
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
func F_emptyreachable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v46 int32
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = m.T0[v9].(func(*base.Module) int32)(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(101)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l3+v24<<(uint(int32(2))%32))))
	if v28 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v20 = v18
	goto L8
L7:
	;
	v20 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v20
	return l2
L9:
	;
	return l1
L10:
	;
	goto L11
L11:
	;
	v33 = l1
	v34 = v28
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v37 != int32(110) {
		v44 = v33
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return v44
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v46 != 0 {
		v33 = v44
		v34 = v46
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v41 != 0 {
		v44 = v33
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v42 = F_emptyreachable(m, l0, v40, v33, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v44 = v42
	goto L14
L18:
	;
	goto L13
}
func F_enable_timeouts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v32 int64
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	v3 = int32(0)
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, _c_F_enable_timeouts[0])) = v3
	v18 = m.G0
	v20 = v18 - v9
	m.G0 = v20
	F_gettimeofday(m, v20)
	mBase = m.M
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	m.G0 = v20 + v9
	v32 = v24 + v23*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	if int32(0) < l1 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v40 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	F_schedule_alarm(m, v32)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L12
	} else {
		goto L20
	}
L5:
	;
	v44 = l0 + v40*int32(24)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	switch v46 {
	case 0:
		goto L8
	case 1:
		goto L11
	case 2:
		goto L10
	default:
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	v81 = v40 + int32(1)
	if v81 != l1 {
		v40 = v81
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v72 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+8)))
	F_enable_timeout(m, v45, v32, v72*int64(1000)+v32, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L18
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L12
	} else {
		goto L15
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	F_enable_timeout(m, v45, v32, base.I64_extend_i32_s(v51)*int64(1000)+v32, v51)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L14
	}
L11:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v44)+16))
	F_enable_timeout(m, v45, v32, v47, int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	goto L7
L14:
	;
	goto L7
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v62
	F_errmsg_internal(m, int32(_a_F_enable_timeouts_0), v10)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_enable_timeouts_1), int32(666), int32(_a_F_enable_timeouts_2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	goto L7
L19:
	;
	goto L6
L20:
	;
	m.G0 = v10 + int32(16)
	return
}
func F_encrypt_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 != 0 {
		F_pgp_cfb_free(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			base.MemoryFill(m, l0, int32(0), int32(_a_F_encrypt_free_0))
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		base.MemoryFill(m, l0, int32(0), int32(_a_F_encrypt_free_0))
		F_pfree(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	}
}
func F_encrypt_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v11 == int32(0) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)) = uint8(v14)
		v20 = F_pushf_write(m, l0, v8+int32(11), v14)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v20 < int32(0) {
				v47 = v20
				m.G0 = v8 + int32(16)
				return v47
			} else {
				v27 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
				v34 = int32(0)
				v36 = F_pgp_cfb_create(m, v8+int32(12), v30, l1+int32(136), v33, v27, v34, v34)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 < int32(0) {
						v47 = v36
						m.G0 = v8 + int32(16)
						return v47
					} else {
						v41 = F_palloc0(m, int32(_a_F_encrypt_init_0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
							v47 = int32(_a_F_encrypt_init_1)
							m.G0 = v8 + int32(16)
							return v47
						}
					}
				}
			}
		}
	} else {
		v27 = int32(1)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+168))
		v34 = int32(0)
		v36 = F_pgp_cfb_create(m, v8+int32(12), v30, l1+int32(136), v33, v27, v34, v34)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if v36 < int32(0) {
				v47 = v36
				m.G0 = v8 + int32(16)
				return v47
			} else {
				v41 = F_palloc0(m, int32(_a_F_encrypt_init_0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v43
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
					v47 = int32(_a_F_encrypt_init_1)
					m.G0 = v8 + int32(16)
					return v47
				}
			}
		}
	}
}
func F_eqsel_internal(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 float32
	_ = v62
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 float64
	_ = v73
	var v83 float64
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 float32
	_ = v95
	var v97 float32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 float64
	_ = v140
	var v143 int32
	_ = v143
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v159 float64
	_ = v159
	var v163 float64
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 float32
	_ = v178
	var v179 float64
	_ = v179
	var v183 float64
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 float32
	_ = v201
	var v203 float32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 float64
	_ = v246
	var v249 int32
	_ = v249
	var v256 float64
	_ = v256
	var v259 float64
	_ = v259
	var v260 int32
	_ = v260
	var v265 float64
	_ = v265
	var v267 float64
	_ = v267
	var v277 float64
	_ = v277
	var v285 float64
	_ = v285
	var v293 float64
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 float64
	_ = v301
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l1 == int32(0) {
		v27 = v18
		v34 = F_get_restriction_variable(m, v19, v17, v16, v13+int32(16), v13+int32(12), v13+int32(11))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return float64(0)
		} else {
			if v34 == int32(0) {
				if l1 != 0 {
					v40 = float64(0.995)
				} else {
					v40 = float64(0.005)
				}
				v301 = v40
				m.G0 = v13 + int32(48)
				return v301
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v42 == int32(7) {
					v47 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+32)))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
					v50 = F_var_eq_const(m, v13+int32(16), v27, v15, v47, v48, v49, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return float64(0)
					} else {
						v293 = v50
						v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						if v294 == int32(0) {
							v301 = v293
							m.G0 = v13 + int32(48)
							return v301
						} else {
							v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
							m.T0[v297].(func(*base.Module, int32))(m, v294)
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return float64(0)
							} else {
								v301 = v293
								m.G0 = v13 + int32(48)
								return v301
							}
						}
					}
				} else {
					v52 = m.G0
					v54 = v52 - int32(48)
					m.G0 = v54
					v57 = v13 + int32(16)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
					if v58 != 0 {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
						v62 = *(*float32)(unsafe.Add(mBase, uint32(v59+v60)+8))
						v66 = base.F64_promote_f32(v62)
					} else {
						v66 = float64(0)
					}
					v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
					if v67 != int32(1) {
						if v58 != 0 {
							v83 = base.F64_sub(float64(1), v66)
							v85 = v54 + int32(47)
							v86 = int32(0)
							v87 = float64(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							if v91 != 0 {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
								v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
								v94 = v92 + v93
								v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
								v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
								v124 = base.F64_promote_f32(v97)
								v125 = base.F64_promote_f32(v95)
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
								if v99 == int32(16) {
									v124 = float64(2)
									v125 = v87
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									if v103 == int32(0) {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										if v110 == int32(0) {
											v124 = float64(0)
											v125 = v87
										} else {
											v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
											if v113 != int32(6) {
												v124 = float64(0)
												v125 = v87
											} else {
												v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
												switch v117 - int32(_a_F_eqsel_internal_0) {
												case 0:
													v124 = float64(1)
													v125 = v87
												default:
													v124 = float64(0)
													v125 = v87
												case 5:
													v124 = float64(-1)
													v125 = v87
												}
											}
										}
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
										if v106 != int32(5) {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											if v110 == int32(0) {
												v124 = float64(0)
												v125 = v87
											} else {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
												if v113 != int32(6) {
													v124 = float64(0)
													v125 = v87
												} else {
													v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
													switch v117 - int32(_a_F_eqsel_internal_0) {
													case 0:
														v124 = float64(1)
														v125 = v87
													default:
														v124 = float64(0)
														v125 = v87
													case 5:
														v124 = float64(-1)
														v125 = v87
													}
												}
											}
										} else {
											v124 = float64(-1)
											v125 = v87
										}
									}
								}
							}
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
							if v129 != 0 {
								v130 = base.F64_neg(base.F64_sub(float64(1), v125))
							} else {
								v130 = v124
							}
							if base.F64_gt(v130, float64(0)) != 0 {
								v133 = F_clamp_row_est(m, v130)
								mBase = m.M
								v159 = v133
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
								if v134 == int32(0) {
									v137 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
									v159 = float64(200)
								} else {
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
									if base.F64_le(v140, float64(0)) != 0 {
										v143 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
										v159 = float64(200)
									} else {
										if base.F64_lt(v130, float64(0)) != 0 {
											v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
											mBase = m.M
											v159 = v150
										} else {
											if base.F64_lt(v140, float64(200)) != 0 {
												v153 = F_clamp_row_est(m, v140)
												mBase = m.M
												v159 = v153
											} else {
												v154 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
												v159 = float64(200)
											}
										}
									}
								}
							}
							if base.F64_gt(v159, float64(1)) != 0 {
								v163 = base.F64_div(v83, v159)
							} else {
								v163 = v83
							}
							v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return float64(0)
							} else {
								if v170 == int32(0) {
									v267 = v163
									m.G0 = v54 + int32(48)
									if l1 != 0 {
										v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
									} else {
										v277 = v267
									}
									if base.F64_lt(v277, float64(0)) != 0 {
										v285 = float64(0)
									} else {
										if base.F64_gt(v277, float64(1)) == int32(0) {
											v285 = v277
										} else {
											v285 = float64(1)
										}
									}
									v293 = v285
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									if v294 == int32(0) {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										}
									}
								} else {
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
									if v174 <= int32(0) {
										v183 = v163
									} else {
										v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
										v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
										v179 = base.F64_promote_f32(v178)
										if base.F64_gt(v163, v179) == int32(0) {
											v183 = v163
										} else {
											v183 = v179
										}
									}
									F_free_attstatsslot(m, v54+int32(8))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return float64(0)
									} else {
										v267 = v183
										m.G0 = v54 + int32(48)
										if l1 != 0 {
											v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
										} else {
											v277 = v267
										}
										if base.F64_lt(v277, float64(0)) != 0 {
											v285 = float64(0)
										} else {
											if base.F64_gt(v277, float64(1)) == int32(0) {
												v285 = v277
											} else {
												v285 = float64(1)
											}
										}
										v293 = v285
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
										if v294 == int32(0) {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											}
										}
									}
								}
							}
						} else {
							v191 = v54 + int32(47)
							v192 = int32(0)
							v193 = float64(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
							v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							if v197 != 0 {
								v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
								v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
								v200 = v198 + v199
								v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
								v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
								v230 = base.F64_promote_f32(v203)
								v231 = base.F64_promote_f32(v201)
							} else {
								v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
								if v205 == int32(16) {
									v230 = float64(2)
									v231 = v193
								} else {
									v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									if v209 == int32(0) {
										v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										if v216 == int32(0) {
											v230 = float64(0)
											v231 = v193
										} else {
											v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
											if v219 != int32(6) {
												v230 = float64(0)
												v231 = v193
											} else {
												v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
												switch v223 - int32(_a_F_eqsel_internal_0) {
												case 0:
													v230 = float64(1)
													v231 = v193
												default:
													v230 = float64(0)
													v231 = v193
												case 5:
													v230 = float64(-1)
													v231 = v193
												}
											}
										}
									} else {
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
										if v212 != int32(5) {
											v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											if v216 == int32(0) {
												v230 = float64(0)
												v231 = v193
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
												if v219 != int32(6) {
													v230 = float64(0)
													v231 = v193
												} else {
													v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
													switch v223 - int32(_a_F_eqsel_internal_0) {
													case 0:
														v230 = float64(1)
														v231 = v193
													default:
														v230 = float64(0)
														v231 = v193
													case 5:
														v230 = float64(-1)
														v231 = v193
													}
												}
											}
										} else {
											v230 = float64(-1)
											v231 = v193
										}
									}
								}
							}
							v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
							if v235 != 0 {
								v236 = base.F64_neg(base.F64_sub(float64(1), v231))
							} else {
								v236 = v230
							}
							if base.F64_gt(v236, float64(0)) != 0 {
								v239 = F_clamp_row_est(m, v236)
								mBase = m.M
								v265 = v239
							} else {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
								if v240 == int32(0) {
									v243 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
									v265 = float64(200)
								} else {
									v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
									if base.F64_le(v246, float64(0)) != 0 {
										v249 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
										v265 = float64(200)
									} else {
										if base.F64_lt(v236, float64(0)) != 0 {
											v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
											mBase = m.M
											v265 = v256
										} else {
											if base.F64_lt(v246, float64(200)) != 0 {
												v259 = F_clamp_row_est(m, v246)
												mBase = m.M
												v265 = v259
											} else {
												v260 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
												v265 = float64(200)
											}
										}
									}
								}
							}
							v267 = base.F64_div(float64(1), v265)
							m.G0 = v54 + int32(48)
							if l1 != 0 {
								v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
							} else {
								v277 = v267
							}
							if base.F64_lt(v277, float64(0)) != 0 {
								v285 = float64(0)
							} else {
								if base.F64_gt(v277, float64(1)) == int32(0) {
									v285 = v277
								} else {
									v285 = float64(1)
								}
							}
							v293 = v285
							v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							if v294 == int32(0) {
								v301 = v293
								m.G0 = v13 + int32(48)
								return v301
							} else {
								v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
								m.T0[v297].(func(*base.Module, int32))(m, v294)
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return float64(0)
								} else {
									v301 = v293
									m.G0 = v13 + int32(48)
									return v301
								}
							}
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						if v70 == int32(0) {
							if v58 != 0 {
								v83 = base.F64_sub(float64(1), v66)
								v85 = v54 + int32(47)
								v86 = int32(0)
								v87 = float64(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
								v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
								if v91 != 0 {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
									v94 = v92 + v93
									v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
									v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
									v124 = base.F64_promote_f32(v97)
									v125 = base.F64_promote_f32(v95)
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
									if v99 == int32(16) {
										v124 = float64(2)
										v125 = v87
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v103 == int32(0) {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											if v110 == int32(0) {
												v124 = float64(0)
												v125 = v87
											} else {
												v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
												if v113 != int32(6) {
													v124 = float64(0)
													v125 = v87
												} else {
													v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
													switch v117 - int32(_a_F_eqsel_internal_0) {
													case 0:
														v124 = float64(1)
														v125 = v87
													default:
														v124 = float64(0)
														v125 = v87
													case 5:
														v124 = float64(-1)
														v125 = v87
													}
												}
											}
										} else {
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
											if v106 != int32(5) {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v110 == int32(0) {
													v124 = float64(0)
													v125 = v87
												} else {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
													if v113 != int32(6) {
														v124 = float64(0)
														v125 = v87
													} else {
														v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
														switch v117 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v124 = float64(1)
															v125 = v87
														default:
															v124 = float64(0)
															v125 = v87
														case 5:
															v124 = float64(-1)
															v125 = v87
														}
													}
												}
											} else {
												v124 = float64(-1)
												v125 = v87
											}
										}
									}
								}
								v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
								if v129 != 0 {
									v130 = base.F64_neg(base.F64_sub(float64(1), v125))
								} else {
									v130 = v124
								}
								if base.F64_gt(v130, float64(0)) != 0 {
									v133 = F_clamp_row_est(m, v130)
									mBase = m.M
									v159 = v133
								} else {
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									if v134 == int32(0) {
										v137 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
										v159 = float64(200)
									} else {
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
										if base.F64_le(v140, float64(0)) != 0 {
											v143 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
											v159 = float64(200)
										} else {
											if base.F64_lt(v130, float64(0)) != 0 {
												v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
												mBase = m.M
												v159 = v150
											} else {
												if base.F64_lt(v140, float64(200)) != 0 {
													v153 = F_clamp_row_est(m, v140)
													mBase = m.M
													v159 = v153
												} else {
													v154 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
													v159 = float64(200)
												}
											}
										}
									}
								}
								if base.F64_gt(v159, float64(1)) != 0 {
									v163 = base.F64_div(v83, v159)
								} else {
									v163 = v83
								}
								v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
								v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
								mBase = m.M
								v171 = m.ExcPending
								if v171 != 0 {
									return float64(0)
								} else {
									if v170 == int32(0) {
										v267 = v163
										m.G0 = v54 + int32(48)
										if l1 != 0 {
											v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
										} else {
											v277 = v267
										}
										if base.F64_lt(v277, float64(0)) != 0 {
											v285 = float64(0)
										} else {
											if base.F64_gt(v277, float64(1)) == int32(0) {
												v285 = v277
											} else {
												v285 = float64(1)
											}
										}
										v293 = v285
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
										if v294 == int32(0) {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											}
										}
									} else {
										v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
										if v174 <= int32(0) {
											v183 = v163
										} else {
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
											v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
											v179 = base.F64_promote_f32(v178)
											if base.F64_gt(v163, v179) == int32(0) {
												v183 = v163
											} else {
												v183 = v179
											}
										}
										F_free_attstatsslot(m, v54+int32(8))
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return float64(0)
										} else {
											v267 = v183
											m.G0 = v54 + int32(48)
											if l1 != 0 {
												v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
											} else {
												v277 = v267
											}
											if base.F64_lt(v277, float64(0)) != 0 {
												v285 = float64(0)
											} else {
												if base.F64_gt(v277, float64(1)) == int32(0) {
													v285 = v277
												} else {
													v285 = float64(1)
												}
											}
											v293 = v285
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
											if v294 == int32(0) {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												}
											}
										}
									}
								}
							} else {
								v191 = v54 + int32(47)
								v192 = int32(0)
								v193 = float64(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
								v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
								if v197 != 0 {
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
									v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
									v200 = v198 + v199
									v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
									v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
									v230 = base.F64_promote_f32(v203)
									v231 = base.F64_promote_f32(v201)
								} else {
									v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
									if v205 == int32(16) {
										v230 = float64(2)
										v231 = v193
									} else {
										v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v209 == int32(0) {
											v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											if v216 == int32(0) {
												v230 = float64(0)
												v231 = v193
											} else {
												v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
												if v219 != int32(6) {
													v230 = float64(0)
													v231 = v193
												} else {
													v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
													switch v223 - int32(_a_F_eqsel_internal_0) {
													case 0:
														v230 = float64(1)
														v231 = v193
													default:
														v230 = float64(0)
														v231 = v193
													case 5:
														v230 = float64(-1)
														v231 = v193
													}
												}
											}
										} else {
											v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
											if v212 != int32(5) {
												v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v216 == int32(0) {
													v230 = float64(0)
													v231 = v193
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
													if v219 != int32(6) {
														v230 = float64(0)
														v231 = v193
													} else {
														v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
														switch v223 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v230 = float64(1)
															v231 = v193
														default:
															v230 = float64(0)
															v231 = v193
														case 5:
															v230 = float64(-1)
															v231 = v193
														}
													}
												}
											} else {
												v230 = float64(-1)
												v231 = v193
											}
										}
									}
								}
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
								if v235 != 0 {
									v236 = base.F64_neg(base.F64_sub(float64(1), v231))
								} else {
									v236 = v230
								}
								if base.F64_gt(v236, float64(0)) != 0 {
									v239 = F_clamp_row_est(m, v236)
									mBase = m.M
									v265 = v239
								} else {
									v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									if v240 == int32(0) {
										v243 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
										v265 = float64(200)
									} else {
										v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
										if base.F64_le(v246, float64(0)) != 0 {
											v249 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
											v265 = float64(200)
										} else {
											if base.F64_lt(v236, float64(0)) != 0 {
												v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
												mBase = m.M
												v265 = v256
											} else {
												if base.F64_lt(v246, float64(200)) != 0 {
													v259 = F_clamp_row_est(m, v246)
													mBase = m.M
													v265 = v259
												} else {
													v260 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
													v265 = float64(200)
												}
											}
										}
									}
								}
								v267 = base.F64_div(float64(1), v265)
								m.G0 = v54 + int32(48)
								if l1 != 0 {
									v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
								} else {
									v277 = v267
								}
								if base.F64_lt(v277, float64(0)) != 0 {
									v285 = float64(0)
								} else {
									if base.F64_gt(v277, float64(1)) == int32(0) {
										v285 = v277
									} else {
										v285 = float64(1)
									}
								}
								v293 = v285
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								if v294 == int32(0) {
									v301 = v293
									m.G0 = v13 + int32(48)
									return v301
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									}
								}
							}
						} else {
							v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+128))
							if base.F64_ge(v73, float64(1)) == int32(0) {
								if v58 != 0 {
									v83 = base.F64_sub(float64(1), v66)
									v85 = v54 + int32(47)
									v86 = int32(0)
									v87 = float64(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									if v91 != 0 {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
										v94 = v92 + v93
										v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
										v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
										v124 = base.F64_promote_f32(v97)
										v125 = base.F64_promote_f32(v95)
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
										if v99 == int32(16) {
											v124 = float64(2)
											v125 = v87
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v103 == int32(0) {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v110 == int32(0) {
													v124 = float64(0)
													v125 = v87
												} else {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
													if v113 != int32(6) {
														v124 = float64(0)
														v125 = v87
													} else {
														v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
														switch v117 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v124 = float64(1)
															v125 = v87
														default:
															v124 = float64(0)
															v125 = v87
														case 5:
															v124 = float64(-1)
															v125 = v87
														}
													}
												}
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
												if v106 != int32(5) {
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v110 == int32(0) {
														v124 = float64(0)
														v125 = v87
													} else {
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
														if v113 != int32(6) {
															v124 = float64(0)
															v125 = v87
														} else {
															v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
															switch v117 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v124 = float64(1)
																v125 = v87
															default:
																v124 = float64(0)
																v125 = v87
															case 5:
																v124 = float64(-1)
																v125 = v87
															}
														}
													}
												} else {
													v124 = float64(-1)
													v125 = v87
												}
											}
										}
									}
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
									if v129 != 0 {
										v130 = base.F64_neg(base.F64_sub(float64(1), v125))
									} else {
										v130 = v124
									}
									if base.F64_gt(v130, float64(0)) != 0 {
										v133 = F_clamp_row_est(m, v130)
										mBase = m.M
										v159 = v133
									} else {
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v134 == int32(0) {
											v137 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
											v159 = float64(200)
										} else {
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
											if base.F64_le(v140, float64(0)) != 0 {
												v143 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
												v159 = float64(200)
											} else {
												if base.F64_lt(v130, float64(0)) != 0 {
													v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
													mBase = m.M
													v159 = v150
												} else {
													if base.F64_lt(v140, float64(200)) != 0 {
														v153 = F_clamp_row_est(m, v140)
														mBase = m.M
														v159 = v153
													} else {
														v154 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
														v159 = float64(200)
													}
												}
											}
										}
									}
									if base.F64_gt(v159, float64(1)) != 0 {
										v163 = base.F64_div(v83, v159)
									} else {
										v163 = v83
									}
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return float64(0)
									} else {
										if v170 == int32(0) {
											v267 = v163
											m.G0 = v54 + int32(48)
											if l1 != 0 {
												v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
											} else {
												v277 = v267
											}
											if base.F64_lt(v277, float64(0)) != 0 {
												v285 = float64(0)
											} else {
												if base.F64_gt(v277, float64(1)) == int32(0) {
													v285 = v277
												} else {
													v285 = float64(1)
												}
											}
											v293 = v285
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
											if v294 == int32(0) {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												}
											}
										} else {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
											if v174 <= int32(0) {
												v183 = v163
											} else {
												v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
												v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
												v179 = base.F64_promote_f32(v178)
												if base.F64_gt(v163, v179) == int32(0) {
													v183 = v163
												} else {
													v183 = v179
												}
											}
											F_free_attstatsslot(m, v54+int32(8))
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return float64(0)
											} else {
												v267 = v183
												m.G0 = v54 + int32(48)
												if l1 != 0 {
													v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
												} else {
													v277 = v267
												}
												if base.F64_lt(v277, float64(0)) != 0 {
													v285 = float64(0)
												} else {
													if base.F64_gt(v277, float64(1)) == int32(0) {
														v285 = v277
													} else {
														v285 = float64(1)
													}
												}
												v293 = v285
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
												if v294 == int32(0) {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v301 = v293
														m.G0 = v13 + int32(48)
														return v301
													}
												}
											}
										}
									}
								} else {
									v191 = v54 + int32(47)
									v192 = int32(0)
									v193 = float64(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									if v197 != 0 {
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
										v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
										v200 = v198 + v199
										v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
										v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
										v230 = base.F64_promote_f32(v203)
										v231 = base.F64_promote_f32(v201)
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
										if v205 == int32(16) {
											v230 = float64(2)
											v231 = v193
										} else {
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v209 == int32(0) {
												v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v216 == int32(0) {
													v230 = float64(0)
													v231 = v193
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
													if v219 != int32(6) {
														v230 = float64(0)
														v231 = v193
													} else {
														v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
														switch v223 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v230 = float64(1)
															v231 = v193
														default:
															v230 = float64(0)
															v231 = v193
														case 5:
															v230 = float64(-1)
															v231 = v193
														}
													}
												}
											} else {
												v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
												if v212 != int32(5) {
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v216 == int32(0) {
														v230 = float64(0)
														v231 = v193
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
														if v219 != int32(6) {
															v230 = float64(0)
															v231 = v193
														} else {
															v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
															switch v223 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v230 = float64(1)
																v231 = v193
															default:
																v230 = float64(0)
																v231 = v193
															case 5:
																v230 = float64(-1)
																v231 = v193
															}
														}
													}
												} else {
													v230 = float64(-1)
													v231 = v193
												}
											}
										}
									}
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
									if v235 != 0 {
										v236 = base.F64_neg(base.F64_sub(float64(1), v231))
									} else {
										v236 = v230
									}
									if base.F64_gt(v236, float64(0)) != 0 {
										v239 = F_clamp_row_est(m, v236)
										mBase = m.M
										v265 = v239
									} else {
										v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v240 == int32(0) {
											v243 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
											v265 = float64(200)
										} else {
											v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
											if base.F64_le(v246, float64(0)) != 0 {
												v249 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
												v265 = float64(200)
											} else {
												if base.F64_lt(v236, float64(0)) != 0 {
													v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
													mBase = m.M
													v265 = v256
												} else {
													if base.F64_lt(v246, float64(200)) != 0 {
														v259 = F_clamp_row_est(m, v246)
														mBase = m.M
														v265 = v259
													} else {
														v260 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
														v265 = float64(200)
													}
												}
											}
										}
									}
									v267 = base.F64_div(float64(1), v265)
									m.G0 = v54 + int32(48)
									if l1 != 0 {
										v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
									} else {
										v277 = v267
									}
									if base.F64_lt(v277, float64(0)) != 0 {
										v285 = float64(0)
									} else {
										if base.F64_gt(v277, float64(1)) == int32(0) {
											v285 = v277
										} else {
											v285 = float64(1)
										}
									}
									v293 = v285
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									if v294 == int32(0) {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										}
									}
								}
							} else {
								v267 = base.F64_div(float64(1), v73)
								m.G0 = v54 + int32(48)
								if l1 != 0 {
									v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
								} else {
									v277 = v267
								}
								if base.F64_lt(v277, float64(0)) != 0 {
									v285 = float64(0)
								} else {
									if base.F64_gt(v277, float64(1)) == int32(0) {
										v285 = v277
									} else {
										v285 = float64(1)
									}
								}
								v293 = v285
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								if v294 == int32(0) {
									v301 = v293
									m.G0 = v13 + int32(48)
									return v301
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v22 = F_get_negator(m, v18)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return float64(0)
		} else {
			if v22 != 0 {
				v27 = v22
				v34 = F_get_restriction_variable(m, v19, v17, v16, v13+int32(16), v13+int32(12), v13+int32(11))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return float64(0)
				} else {
					if v34 == int32(0) {
						if l1 != 0 {
							v40 = float64(0.995)
						} else {
							v40 = float64(0.005)
						}
						v301 = v40
						m.G0 = v13 + int32(48)
						return v301
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						if v42 == int32(7) {
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v41)+24))
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+32)))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
							v50 = F_var_eq_const(m, v13+int32(16), v27, v15, v47, v48, v49, l1)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return float64(0)
							} else {
								v293 = v50
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								if v294 == int32(0) {
									v301 = v293
									m.G0 = v13 + int32(48)
									return v301
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									}
								}
							}
						} else {
							v52 = m.G0
							v54 = v52 - int32(48)
							m.G0 = v54
							v57 = v13 + int32(16)
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
							if v58 != 0 {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
								v62 = *(*float32)(unsafe.Add(mBase, uint32(v59+v60)+8))
								v66 = base.F64_promote_f32(v62)
							} else {
								v66 = float64(0)
							}
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
							if v67 != int32(1) {
								if v58 != 0 {
									v83 = base.F64_sub(float64(1), v66)
									v85 = v54 + int32(47)
									v86 = int32(0)
									v87 = float64(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									if v91 != 0 {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
										v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
										v94 = v92 + v93
										v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
										v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
										v124 = base.F64_promote_f32(v97)
										v125 = base.F64_promote_f32(v95)
									} else {
										v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
										if v99 == int32(16) {
											v124 = float64(2)
											v125 = v87
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v103 == int32(0) {
												v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v110 == int32(0) {
													v124 = float64(0)
													v125 = v87
												} else {
													v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
													if v113 != int32(6) {
														v124 = float64(0)
														v125 = v87
													} else {
														v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
														switch v117 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v124 = float64(1)
															v125 = v87
														default:
															v124 = float64(0)
															v125 = v87
														case 5:
															v124 = float64(-1)
															v125 = v87
														}
													}
												}
											} else {
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
												if v106 != int32(5) {
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v110 == int32(0) {
														v124 = float64(0)
														v125 = v87
													} else {
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
														if v113 != int32(6) {
															v124 = float64(0)
															v125 = v87
														} else {
															v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
															switch v117 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v124 = float64(1)
																v125 = v87
															default:
																v124 = float64(0)
																v125 = v87
															case 5:
																v124 = float64(-1)
																v125 = v87
															}
														}
													}
												} else {
													v124 = float64(-1)
													v125 = v87
												}
											}
										}
									}
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
									if v129 != 0 {
										v130 = base.F64_neg(base.F64_sub(float64(1), v125))
									} else {
										v130 = v124
									}
									if base.F64_gt(v130, float64(0)) != 0 {
										v133 = F_clamp_row_est(m, v130)
										mBase = m.M
										v159 = v133
									} else {
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v134 == int32(0) {
											v137 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
											v159 = float64(200)
										} else {
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
											if base.F64_le(v140, float64(0)) != 0 {
												v143 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
												v159 = float64(200)
											} else {
												if base.F64_lt(v130, float64(0)) != 0 {
													v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
													mBase = m.M
													v159 = v150
												} else {
													if base.F64_lt(v140, float64(200)) != 0 {
														v153 = F_clamp_row_est(m, v140)
														mBase = m.M
														v159 = v153
													} else {
														v154 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
														v159 = float64(200)
													}
												}
											}
										}
									}
									if base.F64_gt(v159, float64(1)) != 0 {
										v163 = base.F64_div(v83, v159)
									} else {
										v163 = v83
									}
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return float64(0)
									} else {
										if v170 == int32(0) {
											v267 = v163
											m.G0 = v54 + int32(48)
											if l1 != 0 {
												v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
											} else {
												v277 = v267
											}
											if base.F64_lt(v277, float64(0)) != 0 {
												v285 = float64(0)
											} else {
												if base.F64_gt(v277, float64(1)) == int32(0) {
													v285 = v277
												} else {
													v285 = float64(1)
												}
											}
											v293 = v285
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
											if v294 == int32(0) {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												}
											}
										} else {
											v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
											if v174 <= int32(0) {
												v183 = v163
											} else {
												v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
												v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
												v179 = base.F64_promote_f32(v178)
												if base.F64_gt(v163, v179) == int32(0) {
													v183 = v163
												} else {
													v183 = v179
												}
											}
											F_free_attstatsslot(m, v54+int32(8))
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return float64(0)
											} else {
												v267 = v183
												m.G0 = v54 + int32(48)
												if l1 != 0 {
													v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
												} else {
													v277 = v267
												}
												if base.F64_lt(v277, float64(0)) != 0 {
													v285 = float64(0)
												} else {
													if base.F64_gt(v277, float64(1)) == int32(0) {
														v285 = v277
													} else {
														v285 = float64(1)
													}
												}
												v293 = v285
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
												if v294 == int32(0) {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v301 = v293
														m.G0 = v13 + int32(48)
														return v301
													}
												}
											}
										}
									}
								} else {
									v191 = v54 + int32(47)
									v192 = int32(0)
									v193 = float64(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
									v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									if v197 != 0 {
										v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
										v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
										v200 = v198 + v199
										v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
										v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
										v230 = base.F64_promote_f32(v203)
										v231 = base.F64_promote_f32(v201)
									} else {
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
										if v205 == int32(16) {
											v230 = float64(2)
											v231 = v193
										} else {
											v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v209 == int32(0) {
												v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
												if v216 == int32(0) {
													v230 = float64(0)
													v231 = v193
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
													if v219 != int32(6) {
														v230 = float64(0)
														v231 = v193
													} else {
														v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
														switch v223 - int32(_a_F_eqsel_internal_0) {
														case 0:
															v230 = float64(1)
															v231 = v193
														default:
															v230 = float64(0)
															v231 = v193
														case 5:
															v230 = float64(-1)
															v231 = v193
														}
													}
												}
											} else {
												v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
												if v212 != int32(5) {
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v216 == int32(0) {
														v230 = float64(0)
														v231 = v193
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
														if v219 != int32(6) {
															v230 = float64(0)
															v231 = v193
														} else {
															v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
															switch v223 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v230 = float64(1)
																v231 = v193
															default:
																v230 = float64(0)
																v231 = v193
															case 5:
																v230 = float64(-1)
																v231 = v193
															}
														}
													}
												} else {
													v230 = float64(-1)
													v231 = v193
												}
											}
										}
									}
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
									if v235 != 0 {
										v236 = base.F64_neg(base.F64_sub(float64(1), v231))
									} else {
										v236 = v230
									}
									if base.F64_gt(v236, float64(0)) != 0 {
										v239 = F_clamp_row_est(m, v236)
										mBase = m.M
										v265 = v239
									} else {
										v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										if v240 == int32(0) {
											v243 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
											v265 = float64(200)
										} else {
											v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
											if base.F64_le(v246, float64(0)) != 0 {
												v249 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
												v265 = float64(200)
											} else {
												if base.F64_lt(v236, float64(0)) != 0 {
													v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
													mBase = m.M
													v265 = v256
												} else {
													if base.F64_lt(v246, float64(200)) != 0 {
														v259 = F_clamp_row_est(m, v246)
														mBase = m.M
														v265 = v259
													} else {
														v260 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
														v265 = float64(200)
													}
												}
											}
										}
									}
									v267 = base.F64_div(float64(1), v265)
									m.G0 = v54 + int32(48)
									if l1 != 0 {
										v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
									} else {
										v277 = v267
									}
									if base.F64_lt(v277, float64(0)) != 0 {
										v285 = float64(0)
									} else {
										if base.F64_gt(v277, float64(1)) == int32(0) {
											v285 = v277
										} else {
											v285 = float64(1)
										}
									}
									v293 = v285
									v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									if v294 == int32(0) {
										v301 = v293
										m.G0 = v13 + int32(48)
										return v301
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										}
									}
								}
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
								if v70 == int32(0) {
									if v58 != 0 {
										v83 = base.F64_sub(float64(1), v66)
										v85 = v54 + int32(47)
										v86 = int32(0)
										v87 = float64(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
										if v91 != 0 {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
											v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
											v94 = v92 + v93
											v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
											v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
											v124 = base.F64_promote_f32(v97)
											v125 = base.F64_promote_f32(v95)
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
											if v99 == int32(16) {
												v124 = float64(2)
												v125 = v87
											} else {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												if v103 == int32(0) {
													v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v110 == int32(0) {
														v124 = float64(0)
														v125 = v87
													} else {
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
														if v113 != int32(6) {
															v124 = float64(0)
															v125 = v87
														} else {
															v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
															switch v117 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v124 = float64(1)
																v125 = v87
															default:
																v124 = float64(0)
																v125 = v87
															case 5:
																v124 = float64(-1)
																v125 = v87
															}
														}
													}
												} else {
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
													if v106 != int32(5) {
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														if v110 == int32(0) {
															v124 = float64(0)
															v125 = v87
														} else {
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
															if v113 != int32(6) {
																v124 = float64(0)
																v125 = v87
															} else {
																v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
																switch v117 - int32(_a_F_eqsel_internal_0) {
																case 0:
																	v124 = float64(1)
																	v125 = v87
																default:
																	v124 = float64(0)
																	v125 = v87
																case 5:
																	v124 = float64(-1)
																	v125 = v87
																}
															}
														}
													} else {
														v124 = float64(-1)
														v125 = v87
													}
												}
											}
										}
										v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
										if v129 != 0 {
											v130 = base.F64_neg(base.F64_sub(float64(1), v125))
										} else {
											v130 = v124
										}
										if base.F64_gt(v130, float64(0)) != 0 {
											v133 = F_clamp_row_est(m, v130)
											mBase = m.M
											v159 = v133
										} else {
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v134 == int32(0) {
												v137 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
												v159 = float64(200)
											} else {
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
												if base.F64_le(v140, float64(0)) != 0 {
													v143 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
													v159 = float64(200)
												} else {
													if base.F64_lt(v130, float64(0)) != 0 {
														v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
														mBase = m.M
														v159 = v150
													} else {
														if base.F64_lt(v140, float64(200)) != 0 {
															v153 = F_clamp_row_est(m, v140)
															mBase = m.M
															v159 = v153
														} else {
															v154 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
															v159 = float64(200)
														}
													}
												}
											}
										}
										if base.F64_gt(v159, float64(1)) != 0 {
											v163 = base.F64_div(v83, v159)
										} else {
											v163 = v83
										}
										v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
										v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return float64(0)
										} else {
											if v170 == int32(0) {
												v267 = v163
												m.G0 = v54 + int32(48)
												if l1 != 0 {
													v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
												} else {
													v277 = v267
												}
												if base.F64_lt(v277, float64(0)) != 0 {
													v285 = float64(0)
												} else {
													if base.F64_gt(v277, float64(1)) == int32(0) {
														v285 = v277
													} else {
														v285 = float64(1)
													}
												}
												v293 = v285
												v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
												if v294 == int32(0) {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v301 = v293
														m.G0 = v13 + int32(48)
														return v301
													}
												}
											} else {
												v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
												if v174 <= int32(0) {
													v183 = v163
												} else {
													v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
													v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
													v179 = base.F64_promote_f32(v178)
													if base.F64_gt(v163, v179) == int32(0) {
														v183 = v163
													} else {
														v183 = v179
													}
												}
												F_free_attstatsslot(m, v54+int32(8))
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return float64(0)
												} else {
													v267 = v183
													m.G0 = v54 + int32(48)
													if l1 != 0 {
														v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
													} else {
														v277 = v267
													}
													if base.F64_lt(v277, float64(0)) != 0 {
														v285 = float64(0)
													} else {
														if base.F64_gt(v277, float64(1)) == int32(0) {
															v285 = v277
														} else {
															v285 = float64(1)
														}
													}
													v293 = v285
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
													if v294 == int32(0) {
														v301 = v293
														m.G0 = v13 + int32(48)
														return v301
													} else {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
														m.T0[v297].(func(*base.Module, int32))(m, v294)
														mBase = m.M
														v299 = m.ExcPending
														if v299 != 0 {
															return float64(0)
														} else {
															v301 = v293
															m.G0 = v13 + int32(48)
															return v301
														}
													}
												}
											}
										}
									} else {
										v191 = v54 + int32(47)
										v192 = int32(0)
										v193 = float64(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
										v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
										if v197 != 0 {
											v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
											v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
											v200 = v198 + v199
											v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
											v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
											v230 = base.F64_promote_f32(v203)
											v231 = base.F64_promote_f32(v201)
										} else {
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
											if v205 == int32(16) {
												v230 = float64(2)
												v231 = v193
											} else {
												v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												if v209 == int32(0) {
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
													if v216 == int32(0) {
														v230 = float64(0)
														v231 = v193
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
														if v219 != int32(6) {
															v230 = float64(0)
															v231 = v193
														} else {
															v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
															switch v223 - int32(_a_F_eqsel_internal_0) {
															case 0:
																v230 = float64(1)
																v231 = v193
															default:
																v230 = float64(0)
																v231 = v193
															case 5:
																v230 = float64(-1)
																v231 = v193
															}
														}
													}
												} else {
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
													if v212 != int32(5) {
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														if v216 == int32(0) {
															v230 = float64(0)
															v231 = v193
														} else {
															v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
															if v219 != int32(6) {
																v230 = float64(0)
																v231 = v193
															} else {
																v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
																switch v223 - int32(_a_F_eqsel_internal_0) {
																case 0:
																	v230 = float64(1)
																	v231 = v193
																default:
																	v230 = float64(0)
																	v231 = v193
																case 5:
																	v230 = float64(-1)
																	v231 = v193
																}
															}
														}
													} else {
														v230 = float64(-1)
														v231 = v193
													}
												}
											}
										}
										v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
										if v235 != 0 {
											v236 = base.F64_neg(base.F64_sub(float64(1), v231))
										} else {
											v236 = v230
										}
										if base.F64_gt(v236, float64(0)) != 0 {
											v239 = F_clamp_row_est(m, v236)
											mBase = m.M
											v265 = v239
										} else {
											v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											if v240 == int32(0) {
												v243 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
												v265 = float64(200)
											} else {
												v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
												if base.F64_le(v246, float64(0)) != 0 {
													v249 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
													v265 = float64(200)
												} else {
													if base.F64_lt(v236, float64(0)) != 0 {
														v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
														mBase = m.M
														v265 = v256
													} else {
														if base.F64_lt(v246, float64(200)) != 0 {
															v259 = F_clamp_row_est(m, v246)
															mBase = m.M
															v265 = v259
														} else {
															v260 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
															v265 = float64(200)
														}
													}
												}
											}
										}
										v267 = base.F64_div(float64(1), v265)
										m.G0 = v54 + int32(48)
										if l1 != 0 {
											v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
										} else {
											v277 = v267
										}
										if base.F64_lt(v277, float64(0)) != 0 {
											v285 = float64(0)
										} else {
											if base.F64_gt(v277, float64(1)) == int32(0) {
												v285 = v277
											} else {
												v285 = float64(1)
											}
										}
										v293 = v285
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
										if v294 == int32(0) {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											}
										}
									}
								} else {
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+128))
									if base.F64_ge(v73, float64(1)) == int32(0) {
										if v58 != 0 {
											v83 = base.F64_sub(float64(1), v66)
											v85 = v54 + int32(47)
											v86 = int32(0)
											v87 = float64(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
											if v91 != 0 {
												v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
												v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
												v94 = v92 + v93
												v95 = *(*float32)(unsafe.Add(mBase, uint32(v94)+8))
												v97 = *(*float32)(unsafe.Add(mBase, uint32(v94)+16))
												v124 = base.F64_promote_f32(v97)
												v125 = base.F64_promote_f32(v95)
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
												if v99 == int32(16) {
													v124 = float64(2)
													v125 = v87
												} else {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													if v103 == int32(0) {
														v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														if v110 == int32(0) {
															v124 = float64(0)
															v125 = v87
														} else {
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
															if v113 != int32(6) {
																v124 = float64(0)
																v125 = v87
															} else {
																v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
																switch v117 - int32(_a_F_eqsel_internal_0) {
																case 0:
																	v124 = float64(1)
																	v125 = v87
																default:
																	v124 = float64(0)
																	v125 = v87
																case 5:
																	v124 = float64(-1)
																	v125 = v87
																}
															}
														}
													} else {
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+84))
														if v106 != int32(5) {
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
															if v110 == int32(0) {
																v124 = float64(0)
																v125 = v87
															} else {
																v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
																if v113 != int32(6) {
																	v124 = float64(0)
																	v125 = v87
																} else {
																	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+8)))
																	switch v117 - int32(_a_F_eqsel_internal_0) {
																	case 0:
																		v124 = float64(1)
																		v125 = v87
																	default:
																		v124 = float64(0)
																		v125 = v87
																	case 5:
																		v124 = float64(-1)
																		v125 = v87
																	}
																}
															}
														} else {
															v124 = float64(-1)
															v125 = v87
														}
													}
												}
											}
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
											if v129 != 0 {
												v130 = base.F64_neg(base.F64_sub(float64(1), v125))
											} else {
												v130 = v124
											}
											if base.F64_gt(v130, float64(0)) != 0 {
												v133 = F_clamp_row_est(m, v130)
												mBase = m.M
												v159 = v133
											} else {
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												if v134 == int32(0) {
													v137 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v137)
													v159 = float64(200)
												} else {
													v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+128))
													if base.F64_le(v140, float64(0)) != 0 {
														v143 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v143)
														v159 = float64(200)
													} else {
														if base.F64_lt(v130, float64(0)) != 0 {
															v150 = F_clamp_row_est(m, base.F64_mul(v140, base.F64_neg(v130)))
															mBase = m.M
															v159 = v150
														} else {
															if base.F64_lt(v140, float64(200)) != 0 {
																v153 = F_clamp_row_est(m, v140)
																mBase = m.M
																v159 = v153
															} else {
																v154 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v154)
																v159 = float64(200)
															}
														}
													}
												}
											}
											if base.F64_gt(v159, float64(1)) != 0 {
												v163 = base.F64_div(v83, v159)
											} else {
												v163 = v83
											}
											v166 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
											v170 = F_get_attstatsslot(m, v54+int32(8), v166, int32(1), int32(0), int32(2))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return float64(0)
											} else {
												if v170 == int32(0) {
													v267 = v163
													m.G0 = v54 + int32(48)
													if l1 != 0 {
														v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
													} else {
														v277 = v267
													}
													if base.F64_lt(v277, float64(0)) != 0 {
														v285 = float64(0)
													} else {
														if base.F64_gt(v277, float64(1)) == int32(0) {
															v285 = v277
														} else {
															v285 = float64(1)
														}
													}
													v293 = v285
													v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
													if v294 == int32(0) {
														v301 = v293
														m.G0 = v13 + int32(48)
														return v301
													} else {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
														m.T0[v297].(func(*base.Module, int32))(m, v294)
														mBase = m.M
														v299 = m.ExcPending
														if v299 != 0 {
															return float64(0)
														} else {
															v301 = v293
															m.G0 = v13 + int32(48)
															return v301
														}
													}
												} else {
													v174 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
													if v174 <= int32(0) {
														v183 = v163
													} else {
														v177 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
														v178 = *(*float32)(unsafe.Add(mBase, uint32(v177)))
														v179 = base.F64_promote_f32(v178)
														if base.F64_gt(v163, v179) == int32(0) {
															v183 = v163
														} else {
															v183 = v179
														}
													}
													F_free_attstatsslot(m, v54+int32(8))
													mBase = m.M
													v188 = m.ExcPending
													if v188 != 0 {
														return float64(0)
													} else {
														v267 = v183
														m.G0 = v54 + int32(48)
														if l1 != 0 {
															v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
														} else {
															v277 = v267
														}
														if base.F64_lt(v277, float64(0)) != 0 {
															v285 = float64(0)
														} else {
															if base.F64_gt(v277, float64(1)) == int32(0) {
																v285 = v277
															} else {
																v285 = float64(1)
															}
														}
														v293 = v285
														v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
														if v294 == int32(0) {
															v301 = v293
															m.G0 = v13 + int32(48)
															return v301
														} else {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
															m.T0[v297].(func(*base.Module, int32))(m, v294)
															mBase = m.M
															v299 = m.ExcPending
															if v299 != 0 {
																return float64(0)
															} else {
																v301 = v293
																m.G0 = v13 + int32(48)
																return v301
															}
														}
													}
												}
											}
										} else {
											v191 = v54 + int32(47)
											v192 = int32(0)
											v193 = float64(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
											v197 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
											if v197 != 0 {
												v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+22)))
												v200 = v198 + v199
												v201 = *(*float32)(unsafe.Add(mBase, uint32(v200)+8))
												v203 = *(*float32)(unsafe.Add(mBase, uint32(v200)+16))
												v230 = base.F64_promote_f32(v203)
												v231 = base.F64_promote_f32(v201)
											} else {
												v205 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
												if v205 == int32(16) {
													v230 = float64(2)
													v231 = v193
												} else {
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													if v209 == int32(0) {
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
														if v216 == int32(0) {
															v230 = float64(0)
															v231 = v193
														} else {
															v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
															if v219 != int32(6) {
																v230 = float64(0)
																v231 = v193
															} else {
																v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
																switch v223 - int32(_a_F_eqsel_internal_0) {
																case 0:
																	v230 = float64(1)
																	v231 = v193
																default:
																	v230 = float64(0)
																	v231 = v193
																case 5:
																	v230 = float64(-1)
																	v231 = v193
																}
															}
														}
													} else {
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+84))
														if v212 != int32(5) {
															v216 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
															if v216 == int32(0) {
																v230 = float64(0)
																v231 = v193
															} else {
																v219 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
																if v219 != int32(6) {
																	v230 = float64(0)
																	v231 = v193
																} else {
																	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+8)))
																	switch v223 - int32(_a_F_eqsel_internal_0) {
																	case 0:
																		v230 = float64(1)
																		v231 = v193
																	default:
																		v230 = float64(0)
																		v231 = v193
																	case 5:
																		v230 = float64(-1)
																		v231 = v193
																	}
																}
															}
														} else {
															v230 = float64(-1)
															v231 = v193
														}
													}
												}
											}
											v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+28)))
											if v235 != 0 {
												v236 = base.F64_neg(base.F64_sub(float64(1), v231))
											} else {
												v236 = v230
											}
											if base.F64_gt(v236, float64(0)) != 0 {
												v239 = F_clamp_row_est(m, v236)
												mBase = m.M
												v265 = v239
											} else {
												v240 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												if v240 == int32(0) {
													v243 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v243)
													v265 = float64(200)
												} else {
													v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+128))
													if base.F64_le(v246, float64(0)) != 0 {
														v249 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v249)
														v265 = float64(200)
													} else {
														if base.F64_lt(v236, float64(0)) != 0 {
															v256 = F_clamp_row_est(m, base.F64_mul(v246, base.F64_neg(v236)))
															mBase = m.M
															v265 = v256
														} else {
															if base.F64_lt(v246, float64(200)) != 0 {
																v259 = F_clamp_row_est(m, v246)
																mBase = m.M
																v265 = v259
															} else {
																v260 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v260)
																v265 = float64(200)
															}
														}
													}
												}
											}
											v267 = base.F64_div(float64(1), v265)
											m.G0 = v54 + int32(48)
											if l1 != 0 {
												v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
											} else {
												v277 = v267
											}
											if base.F64_lt(v277, float64(0)) != 0 {
												v285 = float64(0)
											} else {
												if base.F64_gt(v277, float64(1)) == int32(0) {
													v285 = v277
												} else {
													v285 = float64(1)
												}
											}
											v293 = v285
											v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
											if v294 == int32(0) {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v301 = v293
													m.G0 = v13 + int32(48)
													return v301
												}
											}
										}
									} else {
										v267 = base.F64_div(float64(1), v73)
										m.G0 = v54 + int32(48)
										if l1 != 0 {
											v277 = base.F64_sub(base.F64_sub(float64(1), v267), v66)
										} else {
											v277 = v267
										}
										if base.F64_lt(v277, float64(0)) != 0 {
											v285 = float64(0)
										} else {
											if base.F64_gt(v277, float64(1)) == int32(0) {
												v285 = v277
											} else {
												v285 = float64(1)
											}
										}
										v293 = v285
										v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
										if v294 == int32(0) {
											v301 = v293
											m.G0 = v13 + int32(48)
											return v301
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v301 = v293
												m.G0 = v13 + int32(48)
												return v301
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v301 = float64(0.995)
				m.G0 = v13 + int32(48)
				return v301
			}
		}
	}
}
func F_errhidecontext(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_errhidecontext[0]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_errhidecontext[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errhidecontext_0), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errhidecontext_1), int32(1648), int32(_a_F_errhidecontext_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v24 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_errhidecontext[1]))) = uint8(v24)
		return
	}
}
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn14261(m, l0, l1, int32(_a_F_errmsg_internal_0), int32(1355))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_error_multiple_recovery_targets(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	F_errstart_cold(m, int32(21), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_errcode(m, int32(50856066))
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_errmsg(m, int32(_a_F_error_multiple_recovery_targets_0), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v14 = F_errdetail(m, int32(_a_F_error_multiple_recovery_targets_1), int32(0))
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_error_multiple_recovery_targets_2), int32(_a_F_error_multiple_recovery_targets_3), int32(_a_F_error_multiple_recovery_targets_4))
					v20 = m.ExcPending
					if v20 != 0 {
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
func F_errposition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_errposition[0]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_errposition[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_errposition_0), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_errposition_1), int32(1664), int32(_a_F_errposition_2))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_c_F_errposition[1]))) = l0
		return int32(0)
	}
}
func F_errsave_finish(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[0]))
	if v5 <= v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errfinish(m, l1, l2, l3)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L27
	} else {
		goto L32
	}
L2:
	;
	v12 = v8 * int32(100)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_errsave_finish[1])))
	if int32(21) <= v15 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[0])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L27
	} else {
		goto L29
	}
L5:
	;
	v18 = int32(_a_F_errsave_finish_0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[2])) = v20 + int32(1)
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = F_strlen(m, l1)
	mBase = m.M
	v34 = v27 + int32(1)
	goto L11
L7:
	;
	v77 = v5
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_errsave_finish[3]))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_errsave_finish[4]))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_errsave_finish[5]))) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_errsave_finish[1]))) = int32(21)
	v84 = F_palloc(m, int32(100))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L27
	} else {
		goto L28
	}
L9:
	;
	if v46 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	v36 = int32(0)
	if v34 == v36 {
		v46 = v36
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v46 = v41
	goto L10
L13:
	;
	v40 = v34 - int32(1)
	v41 = l1 + v40
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 != int32(47) {
		v34 = v40
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v49 = v46 + int32(1)
	goto L17
L16:
	;
	v49 = l1
	goto L17
L17:
	;
	v53 = F_strlen(m, v49)
	mBase = m.M
	v60 = v53 + int32(1)
	goto L20
L18:
	;
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v62 = int32(0)
	if v60 == v62 {
		v72 = v62
		goto L19
	} else {
		goto L22
	}
L21:
	;
	v72 = v67
	goto L19
L22:
	;
	v66 = v60 - int32(1)
	v67 = v49 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v68 != int32(92) {
		v60 = v66
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v75 = v72 + int32(1)
	goto L26
L25:
	;
	v75 = v49
	goto L26
L26:
	;
	v77 = v75
	goto L8
L27:
	;
	return
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v84
	base.MemoryCopy(m, v84, v12+int32(_a_F_errsave_finish_1), int32(100))
	v89 = int32(_a_F_errsave_finish_2)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[0]))
	v92 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[0])) = v91 - v92
	v95 = int32(_a_F_errsave_finish_0)
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_errsave_finish[2])) = v97 - v92
	return
L29:
	;
	F_errmsg_internal(m, int32(_a_F_errsave_finish_3), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_errsave_finish_4), int32(710), int32(_a_F_errsave_finish_5))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_esc_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v10 = l0 + l1
	v11 = l0
	v15 = v5
	goto L5
L3:
	;
	v56 = v5
	goto L4
L4:
	;
	m.G0 = v8 + int32(16)
	return v56
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v17 != int32(92) {
		v47 = int32(1)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v56 = v49
	goto L4
L7:
	;
	v49 = v15 + int64(1)
	v50 = v11 + v47
	if base.Ui32(v50) < base.Ui32(v10) {
		v11 = v50
		v15 = v49
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v21 = v11 + int32(3)
	if base.Ui32(v10) <= base.Ui32(v21) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = v11 + int32(1)
	if base.Ui32(v10) <= base.Ui32(v40) {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v23&int32(252) != int32(48) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)))
	if v28&int32(248) != int32(48) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v33&int32(248) != int32(48) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(4)
	goto L7
L14:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 != int32(92) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v47 = int32(2)
	goto L7
L16:
	;
	goto L6
L17:
	;
	return int64(0)
L18:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(_a_F_esc_dec_len_0)
	F_errmsg(m, int32(_a_F_esc_dec_len_1), v8)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_esc_dec_len_2), int32(820), int32(_a_F_esc_dec_len_3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_esc_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	v4 = int64(0)
	if l1 != 0 {
		v8 = l0
		v10 = l2
		v12 = v4
		for {
			v14 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8))))
			if v14 <= int32(0) {
				v17 = int32(92)
				*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v17)
				v19 = int32(7)
				v21 = int32(48)
				v22 = v14&v19 | v21
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)) = uint8(v22)
				v29 = int32(base.Ui32(v14)>>(uint(int32(3))%32))&v19 | v21
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)) = uint8(v29)
				v36 = int32(base.Ui32(v14&int32(192))>>(uint(int32(6))%32)) | v21
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)) = uint8(v36)
				v52 = int64(4)
				v53 = v10 + int32(4)
			} else {
				if v14 == int32(92) {
					v43 = int32(_a_F_esc_encode_0)
					*(*uint16)(unsafe.Add(mBase, uint32(v10))) = uint16(v43)
					v52 = int64(2)
					v53 = v10 + int32(2)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v14)
					v52 = int64(1)
					v53 = v10 + int32(1)
				}
			}
			v54 = v52 + v12
			v56 = v8 + int32(1)
			if base.Ui32(v56) < base.Ui32(l0+l1) {
				v8 = v56
				v10 = v53
				v12 = v54
				continue
			} else {
				break
			}
			break
		}
		v62 = v54
	} else {
		v62 = v4
	}
	return v62
}
func F_escape_json(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 <= v12+int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19+v12))) = uint8(v21)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = v23 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27+v25))) = uint8(v29)
	goto L1
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v34 = l1
	v35 = v32
	goto L10
L8:
	;
	goto L9
L9:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v100 <= v101+int32(1) {
		goto L40
	} else {
		goto L41
	}
L10:
	;
	switch v35 - int32(8) {
	case 0:
		goto L21
	case 1:
		goto L17
	case 2:
		goto L19
	case 3, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25:
		goto L14
	case 4:
		goto L20
	case 5:
		goto L18
	case 26:
		goto L16
	default:
		goto L15
	}
L11:
	;
	goto L9
L12:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v93 != 0 {
		v34 = v34 + int32(1)
		v35 = v93
		goto L10
	} else {
		goto L38
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L37
	}
L14:
	;
	v61 = base.I32_extend8_s(v35)
	if base.Ui32(v35) <= base.Ui32(int32(31)) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	if v35 == int32(92) {
		goto L13
	} else {
		goto L28
	}
L16:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L27
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L26
	}
L18:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_3))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L25
	}
L19:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L24
	}
L20:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_5))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_escape_json_6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L12
L23:
	;
	goto L12
L24:
	;
	goto L12
L25:
	;
	goto L12
L26:
	;
	goto L12
L27:
	;
	goto L12
L28:
	;
	goto L14
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v61
	F_appendStringInfo(m, l0, int32(_a_F_escape_json_7), v9)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 <= v69+int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L12
L33:
	;
	F_appendStringInfoChar(m, l0, v61)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v69))) = uint8(v35)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v80 = v78 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82+v80))) = uint8(v84)
	goto L12
L36:
	;
	goto L12
L37:
	;
	goto L12
L38:
	;
	goto L11
L39:
	;
	m.G0 = v9 + int32(16)
	return
L40:
	;
	F_appendStringInfoChar(m, l0, int32(34))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v108+v101))) = uint8(v110)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v114 = v112 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116+v114))) = uint8(v118)
	goto L39
L43:
	;
	goto L39
}
func F_estimateHyperLogLog(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 float64
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 float64
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 float64
	_ = v95
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v121 float64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v162 float64
	_ = v162
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 float64
	_ = v175
	var v176 int32
	_ = v176
	var v194 float64
	_ = v194
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v293 float64
	_ = v293
	var v296 float64
	_ = v296
	var v298 float64
	_ = v298
	var v311 float64
	_ = v311
	var v321 float64
	_ = v321
	var v332 float64
	_ = v332
	v2 = int32(0)
	v9 = float64(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v11 != int32(1) {
			v21 = v2
			v22 = v2
			v27 = v9
			for {
				v29 = float64(1)
				v30 = v21 + v12
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
				if int32(1024) <= v31 {
					v35 = base.F64_mul(v29, float64(8.98846567431158e+307))
					if base.Ui32(v31) < base.Ui32(int32(2047)) {
						v64 = v35
						v65 = v31 - int32(1023)
					} else {
						v42 = int32(3069)
						if base.Ui32(v42) <= base.Ui32(v31) {
							v45 = v42
						} else {
							v45 = v31
						}
						v64 = base.F64_mul(v35, float64(8.98846567431158e+307))
						v65 = v45 - int32(2046)
					}
				} else {
					if int32(-1023) < v31 {
						v64 = v29
						v65 = v31
					} else {
						v51 = base.F64_mul(v29, float64(2.004168360008973e-292))
						if base.Ui32(int32(-1992)) < base.Ui32(v31) {
							v64 = v51
							v65 = v31 + int32(969)
						} else {
							v58 = int32(-2960)
							if base.Ui32(v31) <= base.Ui32(v58) {
								v61 = v58
							} else {
								v61 = v31
							}
							v64 = base.F64_mul(v51, float64(2.004168360008973e-292))
							v65 = v61 + int32(1938)
						}
					}
				}
				v73 = float64(1)
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if int32(1024) <= v75 {
					v79 = base.F64_mul(v73, float64(8.98846567431158e+307))
					if base.Ui32(v75) < base.Ui32(int32(2047)) {
						v108 = v79
						v109 = v75 - int32(1023)
					} else {
						v86 = int32(3069)
						if base.Ui32(v86) <= base.Ui32(v75) {
							v89 = v86
						} else {
							v89 = v75
						}
						v108 = base.F64_mul(v79, float64(8.98846567431158e+307))
						v109 = v89 - int32(2046)
					}
				} else {
					if int32(-1023) < v75 {
						v108 = v73
						v109 = v75
					} else {
						v95 = base.F64_mul(v73, float64(2.004168360008973e-292))
						if base.Ui32(int32(-1992)) < base.Ui32(v75) {
							v108 = v95
							v109 = v75 + int32(969)
						} else {
							v102 = int32(-2960)
							if base.Ui32(v75) <= base.Ui32(v102) {
								v105 = v102
							} else {
								v105 = v75
							}
							v108 = base.F64_mul(v95, float64(2.004168360008973e-292))
							v109 = v105 + int32(1938)
						}
					}
				}
				v121 = base.F64_add(base.F64_add(v27, base.F64_div(v73, base.F64_mul(v108, base.F64_reinterpret_i64(base.I64_extend_i32_u(v109+int32(1023))<<(uint(int64(52))%64))))), base.F64_div(float64(1), base.F64_mul(v64, base.F64_reinterpret_i64(base.I64_extend_i32_u(v65+int32(1023))<<(uint(int64(52))%64)))))
				v122 = int32(2)
				v123 = v21 + v122
				v125 = v22 + v122
				if v125 != v11&int32(-2) {
					v21 = v123
					v22 = v125
					v27 = v121
					continue
				} else {
					break
				}
				break
			}
			if v11&int32(1) == int32(0) {
				v194 = v121
			} else {
				v131 = v123
				v137 = v121
				v139 = float64(1)
				v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v12))))
				if int32(1024) <= v142 {
					v146 = base.F64_mul(v139, float64(8.98846567431158e+307))
					if base.Ui32(v142) < base.Ui32(int32(2047)) {
						v175 = v146
						v176 = v142 - int32(1023)
					} else {
						v153 = int32(3069)
						if base.Ui32(v153) <= base.Ui32(v142) {
							v156 = v153
						} else {
							v156 = v142
						}
						v175 = base.F64_mul(v146, float64(8.98846567431158e+307))
						v176 = v156 - int32(2046)
					}
				} else {
					if int32(-1023) < v142 {
						v175 = v139
						v176 = v142
					} else {
						v162 = base.F64_mul(v139, float64(2.004168360008973e-292))
						if base.Ui32(int32(-1992)) < base.Ui32(v142) {
							v175 = v162
							v176 = v142 + int32(969)
						} else {
							v169 = int32(-2960)
							if base.Ui32(v142) <= base.Ui32(v169) {
								v172 = v169
							} else {
								v172 = v142
							}
							v175 = base.F64_mul(v162, float64(2.004168360008973e-292))
							v176 = v172 + int32(1938)
						}
					}
				}
				v194 = base.F64_add(v137, base.F64_div(v139, base.F64_mul(v175, base.F64_reinterpret_i64(base.I64_extend_i32_u(v176+int32(1023))<<(uint(int64(52))%64)))))
			}
		} else {
			v131 = v2
			v137 = v9
			v139 = float64(1)
			v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v12))))
			if int32(1024) <= v142 {
				v146 = base.F64_mul(v139, float64(8.98846567431158e+307))
				if base.Ui32(v142) < base.Ui32(int32(2047)) {
					v175 = v146
					v176 = v142 - int32(1023)
				} else {
					v153 = int32(3069)
					if base.Ui32(v153) <= base.Ui32(v142) {
						v156 = v153
					} else {
						v156 = v142
					}
					v175 = base.F64_mul(v146, float64(8.98846567431158e+307))
					v176 = v156 - int32(2046)
				}
			} else {
				if int32(-1023) < v142 {
					v175 = v139
					v176 = v142
				} else {
					v162 = base.F64_mul(v139, float64(2.004168360008973e-292))
					if base.Ui32(int32(-1992)) < base.Ui32(v142) {
						v175 = v162
						v176 = v142 + int32(969)
					} else {
						v169 = int32(-2960)
						if base.Ui32(v142) <= base.Ui32(v169) {
							v172 = v169
						} else {
							v172 = v142
						}
						v175 = base.F64_mul(v162, float64(2.004168360008973e-292))
						v176 = v172 + int32(1938)
					}
				}
			}
			v194 = base.F64_add(v137, base.F64_div(v139, base.F64_mul(v175, base.F64_reinterpret_i64(base.I64_extend_i32_u(v176+int32(1023))<<(uint(int64(52))%64)))))
		}
		v196 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v197 = base.F64_div(v196, v194)
		v198 = base.F64_convert_i32_u(v11)
		if base.F64_le(v197, base.F64_mul(v198, float64(2.5))) == int32(0) {
			v311 = v197
			if base.F64_gt(v311, float64(1.4316557653333333e+08)) == int32(0) {
				v332 = v311
			} else {
				v321 = F_log(m, base.F64_add(base.F64_mul(v311, float64(-2.3283064365386963e-10)), float64(1)))
				mBase = m.M
				v332 = base.F64_mul(v321, float64(-4.294967296e+09))
			}
			return v332
		} else {
			v205 = v11 & int32(3)
			v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v207 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v11) {
				v216 = int32(0)
				v217 = v207
				v218 = v207
				for {
					v225 = v217 + v206
					v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
					v227 = int32(0)
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
					v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+2)))
					v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+3)))
					v241 = v218 + base.B2i32(v226 == v227) + base.B2i32(v230 == v227) + base.B2i32(v234 == v227) + base.B2i32(v238 == v227)
					v242 = int32(4)
					v243 = v217 + v242
					v245 = v216 + v242
					if v245 != v11&int32(-4) {
						v216 = v245
						v217 = v243
						v218 = v241
						continue
					} else {
						break
					}
					break
				}
				if v205 == int32(0) {
					v282 = v241
				} else {
					v251 = v243
					v252 = v241
					v261 = v251
					v262 = v252
					v265 = v207
					for {
						v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v206))))
						v273 = v262 + base.B2i32(v270 == int32(0))
						v274 = int32(1)
						v277 = v265 + v274
						if v277 != v205 {
							v261 = v261 + v274
							v262 = v273
							v265 = v277
							continue
						} else {
							break
						}
						break
					}
					v282 = v273
				}
			} else {
				v251 = v207
				v252 = v207
				v261 = v251
				v262 = v252
				v265 = v207
				for {
					v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261+v206))))
					v273 = v262 + base.B2i32(v270 == int32(0))
					v274 = int32(1)
					v277 = v265 + v274
					if v277 != v205 {
						v261 = v261 + v274
						v262 = v273
						v265 = v277
						continue
					} else {
						break
					}
					break
				}
				v282 = v273
			}
			if v282 == int32(0) {
				v332 = v197
				return v332
			} else {
				v293 = F_log(m, base.F64_div(v198, base.F64_convert_i32_s(v282)))
				mBase = m.M
				return base.F64_mul(v293, v198)
			}
		}
	} else {
		v296 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v298 = base.F64_div(v296, float64(0))
		if base.F64_le(v298, base.F64_mul(base.F64_convert_i32_u(v11), float64(2.5))) != 0 {
			v332 = v298
		} else {
			v311 = v298
			if base.F64_gt(v311, float64(1.4316557653333333e+08)) == int32(0) {
				v332 = v311
			} else {
				v321 = F_log(m, base.F64_add(base.F64_mul(v311, float64(-2.3283064365386963e-10)), float64(1)))
				mBase = m.M
				v332 = base.F64_mul(v321, float64(-4.294967296e+09))
			}
		}
		return v332
	}
}
func F_estimate_hashagg_tablesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v8 = v6
	} else {
		v8 = int32(0)
	}
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v19 = int32(1)
	if v11&(v11-v19) != 0 {
		v27 = v19 << (uint(int32(32)-base.I32_clz(v11)) % 32)
	} else {
		v27 = v11
	}
	if v11 != 0 {
		v31 = v27 + int32(8)
	} else {
		v31 = int32(0)
	}
	return base.F64_mul(l3, base.F64_convert_i32_u((v10+int32(23))&int32(-8)+v8<<(uint(int32(4))%32)+v31+int32(12)))
}
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v36 float64
	_ = v36
	var v40 float64
	_ = v40
	var v43 int32
	_ = v43
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 float64
	_ = v122
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 float64
	_ = v241
	var v243 float64
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v266 float64
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v341 int32
	_ = v341
	var v342 float64
	_ = v342
	var v346 float64
	_ = v346
	var v350 int32
	_ = v350
	var v351 float64
	_ = v351
	var v362 float64
	_ = v362
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 float64
	_ = v385
	var v390 int32
	_ = v390
	var v396 float64
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 float64
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 float64
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 float64
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 float64
	_ = v426
	var v428 float64
	_ = v428
	var v430 int32
	_ = v430
	var v431 float64
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 float64
	_ = v435
	var v446 float64
	_ = v446
	var v452 float64
	_ = v452
	var v460 float64
	_ = v460
	var v468 float64
	_ = v468
	var v480 float64
	_ = v480
	var v481 float64
	_ = v481
	var v486 float64
	_ = v486
	var v491 float64
	_ = v491
	var v499 float64
	_ = v499
	var v503 float64
	_ = v503
	var v504 float64
	_ = v504
	var v513 float64
	_ = v513
	var v517 float64
	_ = v517
	var v535 float64
	_ = v535
	var v540 float64
	_ = v540
	var v542 float64
	_ = v542
	var v545 float64
	_ = v545
	var v562 float64
	_ = v562
	var v564 float64
	_ = v564
	var v567 float64
	_ = v567
	var v569 float64
	_ = v569
	var v572 float64
	_ = v572
	var v586 float64
	_ = v586
	v6 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v26 = float64(1)
	v27 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l2)&int64(9223372036854775807)))|base.F64_gt(l2, v27) != 0 {
		v40 = v27
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if l1 == int32(0) {
		v586 = v26
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L4
L6:
	;
	v36 = float64(1)
	if base.F64_le(l2, v36) != 0 {
		v40 = v36
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v40 = base.F64_nearest(l2)
	goto L5
L8:
	;
	m.G0 = v22 + int32(48)
	return v586
L9:
	;
	if l3 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v43 == int32(0) {
		v586 = v26
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v46 = float64(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v48 <= int32(0) {
		v562 = v46
		v564 = v46
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	v567 = base.F64_ceil(base.F64_mul(v562, v564))
	if base.F64_gt(v567, v40) != 0 {
		goto L139
	} else {
		goto L140
	}
L15:
	;
	v56 = v6
	v59 = v6
	v60 = v6
	v67 = v46
	v69 = v46
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v60<<(uint(int32(2))%32))))
	if l3 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	if v233 == int32(0) {
		v562 = v241
		v564 = v243
		goto L14
	} else {
		goto L73
	}
L18:
	;
	v245 = v60 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v245 < v246 {
		v56 = v230
		v59 = v233
		v60 = v245
		v67 = v241
		v69 = v243
		goto L16
	} else {
		goto L72
	}
L19:
	;
	v230 = v216
	v233 = v214
	v241 = v222
	v243 = v224
	goto L18
L20:
	;
	v122 = F_expression_returns_set_rows(m, l0, v74)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v121 = v56
	goto L20
L22:
	;
	goto L23
L23:
	;
	v78 = v56 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v80 = int32(0)
	if v79 == v80 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v118 == int32(0) {
		v214 = v59
		v216 = v78
		v222 = v67
		v224 = v69
		goto L19
	} else {
		goto L37
	}
L25:
	;
	v118 = int32(0)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v86 <= int32(0) {
		v112 = v80
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = v112
	goto L24
L29:
	;
	v89 = int32(0)
	if v89 < v86 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v92 = v86
	goto L32
L31:
	;
	v92 = v89
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v95 = int32(0)
	goto L33
L33:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93+v95<<(uint(int32(2))%32))))
	v104 = base.B2i32(v103 == v56)
	if v103 == v56 {
		v112 = v104
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v112 = v104
	goto L28
L35:
	;
	v106 = v95 + int32(1)
	if v106 != v92 {
		v95 = v106
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v121 = v78
	goto L20
L38:
	;
	return float64(0)
L39:
	;
	if base.F64_gt(v122, v69) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v127 = v122
	goto L42
L41:
	;
	v127 = v69
	goto L42
L42:
	;
	v128 = F_exprType(m, v74)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	if v128 == int32(16) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v214 = v59
	v216 = v121
	v222 = base.F64_add(v67, v67)
	v224 = v127
	goto L19
L45:
	;
	goto L46
L46:
	;
	F_examine_variable(m, l0, v74, int32(0), v22+int32(16))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v138 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v157 = F_pull_var_clause(m, v74, int32(42))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L38
	} else {
		goto L56
	}
L49:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+44)))
	if v141&int32(1) == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v148 = F_add_unique_group_var(m, l0, v59, v74, v22+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L38
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v150 == int32(0) {
		v214 = v148
		v216 = v121
		v222 = v67
		v224 = v127
		goto L19
	} else {
		goto L54
	}
L54:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	m.T0[v153].(func(*base.Module, int32))(m, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L38
	} else {
		goto L55
	}
L55:
	;
	v214 = v148
	v216 = v121
	v222 = v67
	v224 = v127
	goto L19
L56:
	;
	if v157 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v159 <= int32(0) {
		v214 = v59
		v216 = v121
		v222 = v67
		v224 = v127
		goto L19
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v202 = F_contain_volatile_functions(m, v74)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L38
	} else {
		goto L70
	}
L60:
	;
	v170 = int32(0)
	v171 = v59
	goto L61
L61:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182+v170<<(uint(int32(2))%32))))
	v189 = v22 + int32(16)
	F_examine_variable(m, l0, v186, int32(0), v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L38
	} else {
		goto L63
	}
L62:
	;
	v214 = v192
	v216 = v121
	v222 = v67
	v224 = v127
	goto L19
L63:
	;
	v192 = F_add_unique_group_var(m, l0, v171, v186, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v194 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	m.T0[v195].(func(*base.Module, int32))(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L38
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v199 = v170 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v199 < v200 {
		v170 = v199
		v171 = v192
		goto L61
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	goto L62
L70:
	;
	if v202 == int32(0) {
		v230 = v121
		v233 = v59
		v241 = v67
		v243 = v127
		goto L18
	} else {
		goto L71
	}
L71:
	;
	v586 = v40
	goto L8
L72:
	;
	goto L17
L73:
	;
	v258 = v233
	v266 = v241
	goto L74
L74:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	v272 = int32(0)
	v274 = F_lappend(m, v272, v270)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L38
	} else {
		goto L76
	}
L75:
	;
	v540 = base.F64_ceil(base.F64_mul(v243, v535))
	if base.F64_gt(v540, v40) != 0 {
		goto L133
	} else {
		goto L134
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v274
	v278 = int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v278 < v279 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v289 = v278
	v291 = int32(0)
	goto L80
L78:
	;
	v327 = v274
	v341 = int32(0)
	goto L79
L79:
	;
	v342 = float64(1)
	if v327 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L80:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301+v289<<(uint(int32(2))%32))))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v306 == v307 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v327 = v320
	v341 = v315
	goto L79
L82:
	;
	v317 = v289 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v317 < v318 {
		v289 = v317
		v291 = v315
		goto L80
	} else {
		goto L88
	}
L83:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v310 = F_lappend(m, v309, v305)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L38
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v313 = F_lappend(m, v291, v305)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L38
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v310
	v315 = v291
	goto L82
L87:
	;
	v315 = v313
	goto L82
L88:
	;
	goto L81
L89:
	;
	if v341 != 0 {
		v258 = v341
		v266 = v535
		goto L74
	} else {
		goto L132
	}
L90:
	;
	if base.F64_gt(v468, v480) != 0 {
		goto L123
	} else {
		goto L124
	}
L91:
	;
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v271)+128))
	if base.F64_gt(v346, float64(0)) != 0 {
		v468 = v342
		v480 = v346
		v481 = v346
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v350 = v272
	v351 = v342
	v362 = v342
	goto L95
L94:
	;
	v535 = v266
	goto L89
L95:
	;
	v372 = F_estimate_multivariate_ndistinct(m, l0, v271, v22+int32(12), v22+int32(16))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L38
	} else {
		goto L98
	}
L96:
	;
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v271)+128))
	if base.F64_gt(v452, float64(0)) == int32(0) {
		v535 = v266
		goto L89
	} else {
		goto L117
	}
L97:
	;
	goto L96
L98:
	;
	if v372 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v376 == int32(0) {
		v434 = v350
		v435 = v351
		v446 = v362
		goto L97
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v426 = *(*float64)(unsafe.Add(mBase, uint32(v22)+16))
	if base.F64_lt(v362, v426) != 0 {
		goto L113
	} else {
		goto L114
	}
L102:
	;
	v379 = int32(0)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v380 <= v379 {
		v434 = v350
		v435 = v351
		v446 = v362
		goto L97
	} else {
		goto L103
	}
L103:
	;
	v384 = v350
	v385 = v351
	v390 = v379
	v396 = v362
	goto L104
L104:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v376)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v402+v390<<(uint(int32(2))%32))))
	v407 = *(*float64)(unsafe.Add(mBase, uint32(v406)+8))
	if l4 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v434 = v420
	v435 = v421
	v446 = v418
	goto L97
L106:
	;
	if base.F64_lt(v396, v407) != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+16)))
	if v411 != int32(1) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v414 | int32(1)
	goto L106
L109:
	;
	v418 = v407
	goto L111
L110:
	;
	v418 = v396
	goto L111
L111:
	;
	v419 = int32(1)
	v420 = v384 + v419
	v421 = base.F64_mul(v385, v407)
	v423 = v390 + v419
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	if v423 < v424 {
		v384 = v420
		v385 = v421
		v390 = v423
		v396 = v418
		goto L104
	} else {
		goto L112
	}
L112:
	;
	goto L105
L113:
	;
	v428 = v426
	goto L115
L114:
	;
	v428 = v362
	goto L115
L115:
	;
	v430 = v350 + int32(1)
	v431 = base.F64_mul(v351, v426)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v432 != 0 {
		v350 = v430
		v351 = v431
		v362 = v428
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v434 = v430
	v435 = v431
	v446 = v428
	goto L97
L117:
	;
	if v434 < int32(2) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v468 = v435
	v480 = v452
	v481 = v452
	goto L90
L119:
	;
	v460 = base.F64_mul(v452, float64(0.1))
	if base.F64_lt(v460, v446) == int32(0) {
		v468 = v435
		v480 = v460
		v481 = v452
		goto L90
	} else {
		goto L120
	}
L120:
	;
	if base.F64_gt(v446, v452) != 0 {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	v468 = v435
	v480 = v446
	v481 = v452
	goto L90
L122:
	;
	v504 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v503)&int64(9223372036854775807)))|base.F64_gt(v503, v504) != 0 {
		v517 = v504
		goto L129
	} else {
		goto L130
	}
L123:
	;
	v486 = v480
	goto L125
L124:
	;
	v486 = v468
	goto L125
L125:
	;
	if base.F64_gt(v486, float64(0)) == int32(0) {
		v503 = v486
		goto L122
	} else {
		goto L126
	}
L126:
	;
	v491 = *(*float64)(unsafe.Add(mBase, uint32(v271)+16))
	if base.F64_lt(v491, v481) == int32(0) {
		v503 = v486
		goto L122
	} else {
		goto L127
	}
L127:
	;
	v499 = F_pow(m, base.F64_div(base.F64_sub(v481, v491), v481), base.F64_div(v481, v486))
	mBase = m.M
	v503 = base.F64_mul(v486, base.F64_sub(float64(1), v499))
	goto L122
L128:
	;
	v535 = base.F64_mul(v266, v517)
	goto L89
L129:
	;
	goto L128
L130:
	;
	v513 = float64(1)
	if base.F64_le(v503, v513) != 0 {
		v517 = v513
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v517 = base.F64_nearest(v503)
	goto L129
L132:
	;
	goto L75
L133:
	;
	v542 = v40
	goto L135
L134:
	;
	v542 = v540
	goto L135
L135:
	;
	if base.F64_lt(v542, float64(1)) != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v545 = float64(1)
	goto L138
L137:
	;
	v545 = v542
	goto L138
L138:
	;
	v586 = v545
	goto L8
L139:
	;
	v569 = v40
	goto L141
L140:
	;
	v569 = v567
	goto L141
L141:
	;
	if base.F64_lt(v569, float64(1)) != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v572 = float64(1)
	goto L144
L143:
	;
	v572 = v569
	goto L144
L144:
	;
	v586 = v572
	goto L8
}
func F_eval_const_expressions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = v10
	} else {
		v11 = int32(0)
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v16)
	v20 = F_eval_const_expressions_mutator(m, l1, v7+int32(12))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(32)
		return v20
	}
}
func F_examine_attribute(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int64
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v4 = int32(0)
	v12 = m.G0
	v13 = int32(16)
	v14 = v12 - v13
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v25 = v16 + v17<<(uint(int32(3))%32) + l1*int32(100) - int32(72)
	v27 = v14 + int32(12)
	v28 = m.G0
	v30 = v28 - v13
	m.G0 = v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+91)))
	if v32 != 0 {
		v73 = v4
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v196
L3:
	;
	if v73 == int32(0) {
		v196 = v4
		goto L2
	} else {
		goto L23
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L20
	}
L5:
	;
	m.G0 = v30 + int32(16)
	goto L3
L6:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+90)))
	if v33 == int32(118) {
		v73 = v4
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v40 = F_SearchSysCache2(m, int32(7), v37, base.I64_extend16_s(base.I64_extend_i32_u(l1)))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v50 = F_SysCacheGetAttr(m, int32(7), v40, int32(21), v30+int32(15))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v68 = int32(1)
	if v27 == int32(0) {
		v73 = v68
		goto L5
	} else {
		goto L19
	}
L13:
	;
	F_ReleaseCatCache(m, v40)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_ReleaseCatCache(m, v40)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L17
	}
L16:
	;
	v67 = int32(-1)
	goto L12
L17:
	;
	v60 = base.I32_wrap_i64(v50)
	if v60&int32(_a_F_examine_attribute_5) == int32(0) {
		v73 = v4
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v67 = base.I32_extend16_s(v60)
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v67
	v73 = v68
	goto L5
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = l1
	F_errmsg_internal(m, int32(_a_F_examine_attribute_3), v30)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_examine_attribute_1), int32(1205), int32(_a_F_examine_attribute_4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v96 = F_palloc0(m, int32(248))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v98
	if l2 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+16)) = v120
	v123 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v96)+4)))
	v125 = F_SearchSysCacheCopy(m, int32(82), v123, int64(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L33
	}
L26:
	;
	v100 = F_exprType(m, l2)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v25)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+96))
	v120 = v119
	goto L25
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v100
	v103 = F_exprTypmod(m, l2)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v103
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106+(l1-int32(1))<<(uint(int32(2))%32))))
	if v112 != 0 {
		v120 = v112
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v113 = F_exprCollation(m, l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v120 = v113
	goto L25
L33:
	;
	if v125 == int32(0) {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+22)))
	v131 = v129 + v130
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v131
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_examine_attribute[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+224)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v96)+20)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+184)) = v137
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+204)) = uint16(v139)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+214)) = uint8(v141)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+188)) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+219)) = uint8(v143)
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+206)) = uint16(v146)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+215)) = uint8(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+192)) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+220)) = uint8(v150)
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+208)) = uint16(v153)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+216)) = uint8(v155)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+196)) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+221)) = uint8(v157)
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+210)) = uint16(v160)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+217)) = uint8(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+128)))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+200)) = v137
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+222)) = uint8(v164)
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v96)+212)) = uint16(v167)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+218)) = uint8(v169)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96)+223)) = uint8(v171)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v131)+124))
	if v173 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_pfree(m, v125)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L45
	}
L36:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	if v182 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L37:
	;
	v176 = F_OidFunctionCall1Coll(m, v173, int32(0), base.I64_extend_i32_u(v96))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v180 = F_std_typanalyze(m, v96)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	if v176 != int64(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L35
L42:
	;
	goto L36
L43:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	if int32(0) < v185 {
		v196 = v96
		goto L2
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	F_pfree(m, v96)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v196 = int32(0)
	goto L2
L47:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v206
	F_errmsg_internal(m, int32(_a_F_examine_attribute_0), v14)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_examine_attribute_1), int32(1137), int32(_a_F_examine_attribute_2))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_execTuplesMatchPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v7 = int32(0)
	if l1 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v16 = F_palloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if int32(0) < l1 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = v7
	goto L9
L7:
	;
	goto L8
L8:
	;
	v51 = int32(0)
	v53 = F_ExecBuildGroupingEqual(m, l0, l0, v51, v51, l1, l2, v16, l4, l5)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L13
	}
L9:
	;
	v32 = v29 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3+v32)))
	v36 = F_get_opcode(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v32))) = v36
	v40 = v29 + int32(1)
	if v40 != l1 {
		v29 = v40
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	return v53
}
func F_exec_check_assignable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v11 = l1
	for {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v9+v11<<(uint(int32(2))%32))))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v18 != int32(3) {
			break
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v11 = v43
			continue
		}
		break
	}
	switch v18 {
	case 0, 2, 4:
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)))
		if v21 != int32(1) {
			m.G0 = v7 + int32(32)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(_a_F_exec_check_assignable_0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				F_errcode(m, int32(83886210))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v31
					F_errmsg(m, int32(_a_F_exec_check_assignable_1), v7+int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_exec_check_assignable_2), int32(_a_F_exec_check_assignable_3), int32(_a_F_exec_check_assignable_4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
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
	case 1:
		m.G0 = v7 + int32(32)
		return
	default:
		F_errstart_cold(m, int32(21), int32(_a_F_exec_check_assignable_0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v48
			F_errmsg_internal(m, int32(_a_F_exec_check_assignable_5), v7)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_exec_check_assignable_2), int32(_a_F_exec_check_assignable_6), int32(_a_F_exec_check_assignable_4))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
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
func F_exec_eval_using_params(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l1 == v3 {
		v132 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v132
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v21 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(_a_F_exec_eval_using_params_0), int32(0), int32(_a_F_exec_eval_using_params_1), int32(_a_F_exec_eval_using_params_2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v34 = v21
	goto L5
L5:
	;
	v35 = int32(_a_F_exec_eval_using_params_3)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0])) = v34
	v39 = F_makeParamList(m, v20)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return int32(0)
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v29
	v34 = v29
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0])) = v36
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v43 <= int32(0) {
		v132 = v39
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v53 = int32(0)
	goto L10
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v53<<(uint(int32(2))%32))))
	v69 = v39 + int32(32) + v53<<(uint(int32(4))%32)
	v70 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v69)+10)) = uint16(v70)
	v73 = v69 + int32(8)
	v74 = int32(12)
	v75 = v69 + v74
	v78 = F_exec_eval_expr(m, l0, v66, v73, v75, v16+v74)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v132 = v39
	goto L1
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v78
	v81 = int32(_a_F_exec_eval_using_params_3)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0])) = v34
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	if v85 == int32(705) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_using_params[0])) = v82
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v112 != 0 {
		goto L23
	} else {
		goto L24
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(25)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v90 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v96 != 0 {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v92 = F_cstring_to_text(m, base.I32_wrap_i64(v78))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = base.I64_extend_i32_u(v92)
	goto L13
L19:
	;
	F_get_typlenbyval(m, v85, v16+int32(10), v16+int32(9))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)))
	if v103 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	v106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+10)))
	v107 = F_datumCopy(m, v104, int32(0), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v107
	goto L13
L23:
	;
	F_SPI_freetuptable(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = int32(0)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v117 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+20))
	F_MemoryContextReset(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v122 = v53 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v122 < v123 {
		v53 = v122
		goto L10
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	goto L11
}
func F_exec_move_row(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	v14 = m.G0
	v16 = v14 - int32(576)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18 != int32(2) {
		v264 = int32(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(576)
	return
L2:
	;
	v271 = int32(0)
	if base.B2i32(l2 == v271)|base.B2i32(l3 == v271) == v271 {
		goto L93
	} else {
		goto L94
	}
L3:
	;
	if l3 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v88 != int32(2249) {
		goto L39
	} else {
		goto L40
	}
L7:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v78 != 0 {
		goto L34
	} else {
		goto L35
	}
L8:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+15)))
	if v26 != int32(100) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v30 = F_make_expanded_record_for_rec(m, l0, l1, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v32 = int32(0)
	F_expanded_record_set_tuple(m, v30, v32, v32, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v42 != v38 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v71 != 0 {
		goto L30
	} else {
		goto L31
	}
L14:
	;
	if v42 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	if v38 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v46 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v46
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+20)) = v46
	goto L19
L23:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v52
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v38
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = v59
	if v59 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = int32(0)
	goto L16
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v37
	goto L29
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v37
	goto L13
L30:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v71+int32(12)))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v30
	goto L1
L33:
	;
	goto L32
L34:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v78+int32(12)))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	goto L1
L37:
	;
	goto L36
L38:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if base.B2i32(l2 == int32(0))|base.B2i32(v102 == int32(2249)) != 0 {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	F_revalidate_rectypeid(m, l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v97 = F_make_expanded_record_from_tupdesc(m, l3, v87)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L44
	}
L42:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v95 = F_make_expanded_record_from_typeid(m, v93, int32(-1), v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v99 = v95
	goto L38
L44:
	;
	v99 = v97
	goto L38
L45:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	if v222 != v218 {
		goto L73
	} else {
		goto L74
	}
L46:
	;
	v196 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	F_expanded_record_set_tuple(m, v99, l2, v196, (v197^int32(-1))&v196)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L71
	}
L47:
	;
	if l2 != 0 {
		goto L46
	} else {
		goto L69
	}
L48:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v102 == v106 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)+44))
	if v108 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v111 = F_expanded_record_fetch_tupdesc(m, v99)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L10
	} else {
		goto L53
	}
L51:
	;
	v113 = v108
	goto L52
L52:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v114 != v115 {
		v264 = v99
		goto L2
	} else {
		goto L54
	}
L53:
	;
	v113 = v111
	goto L52
L54:
	;
	if v114 <= int32(0) {
		goto L46
	} else {
		goto L55
	}
L55:
	;
	v120 = v114 << (uint(int32(3)) % 32)
	v122 = int32(28)
	v132 = int32(0)
	goto L56
L56:
	;
	v142 = v132 * int32(100)
	v143 = v113 + v120 + v122 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+91)))
	v145 = v142 + (l3 + v120 + v122)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+91)))
	if v144 != v146 {
		v264 = v99
		goto L2
	} else {
		goto L58
	}
L57:
	;
	goto L47
L58:
	;
	if v144 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v166 = v132 + int32(1)
	if v166 != v114 {
		v132 = v166
		goto L56
	} else {
		goto L68
	}
L60:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v143)+68))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145)+68))
	if v150 != v151 {
		v264 = v99
		goto L2
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+72)))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+72)))
	if v158 != v159 {
		v264 = v99
		goto L2
	} else {
		goto L66
	}
L63:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	if v153 < int32(0) {
		goto L59
	} else {
		goto L64
	}
L64:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v145)+76))
	if v153 == v156 {
		goto L59
	} else {
		goto L65
	}
L65:
	;
	v264 = v99
	goto L2
L66:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+83)))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+83)))
	if v161 != v162 {
		v264 = v99
		goto L2
	} else {
		goto L67
	}
L67:
	;
	goto L59
L68:
	;
	goto L57
L69:
	;
	F_deconstruct_expanded_record(m, v99)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	goto L45
L71:
	;
	goto L45
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v251 != 0 {
		goto L89
	} else {
		goto L90
	}
L73:
	;
	if v222 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	goto L72
L76:
	;
	if v218 != 0 {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	if v227 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v226 == int32(0) {
		goto L76
	} else {
		goto L82
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+28)) = v226
	goto L78
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+20)) = v226
	goto L78
L82:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+24)) = v232
	goto L76
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v218
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+28)) = v239
	if v239 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v217)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = int32(0)
	goto L75
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239)+24)) = v217
	goto L88
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+20)) = v217
	goto L72
L89:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v251+int32(12)))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v99
	goto L1
L92:
	;
	goto L91
L93:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui32(v278) < base.Ui32(int32(65)) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v298 = int32(0)
	F_exec_move_row_from_fields(m, l0, l1, v264, v298, v298, v298)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L10
	} else {
		goto L103
	}
L96:
	;
	F_heap_deform_tuple(m, l2, l3, v293, v292)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L10
	} else {
		goto L101
	}
L97:
	;
	v292 = v16
	v293 = v16 - int32(-64)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	v287 = F_MemoryContextAlloc(m, v284, v278*int32(9))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	v292 = v287 + v278<<(uint(int32(3))%32)
	v293 = v287
	goto L96
L101:
	;
	F_exec_move_row_from_fields(m, l0, l1, v264, v293, v292, l3)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	goto L1
L103:
	;
	goto L1
}
func F_executeUnaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int64
	_ = v192
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = int64(8589934592)
	v23 = v16 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v23
	v26 = v16 + int32(100)
	F_jspGetArg(m, l1, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = F_executeItemOptUnwrapResult(m, l0, v26, l2, int32(1), v23)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v16 + int32(128)
	return v246
L4:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v222 == int32(0) {
		v246 = v214
		goto L3
	} else {
		goto L50
	}
L5:
	;
	if v33 == int32(2) {
		v214 = int32(2)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v37 = F_jspGetNext(m, l1, v26)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v41 = v37 | base.B2i32(l4 != int32(0))
	v45 = v23
	v51 = int32(1)
	v55 = v6
	goto L8
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v55 < v56 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v214 = v207
	goto L4
L10:
	;
	v63 = v61 + int32(1)
	v68 = v60 + v61<<(uint(int32(5))%32) + int32(16)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v60 = v45
	v61 = v55
	goto L10
L12:
	;
	goto L13
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v59 != 0 {
		v60 = v59
		v61 = int32(0)
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v214 = v51
	goto L4
L15:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v205 = F_executeItemOptUnwrapTarget(m, l0, v16+int32(100), v68, l4, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L48
	}
L16:
	;
	v175 = int32(16)
	v177 = v92 << (uint(int32(1)) % 32)
	if v177 <= v175 {
		goto L43
	} else {
		goto L44
	}
L17:
	;
	v72 = int32(0)
	if v41&int32(1) == v72 {
		v214 = v72
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v41&int32(1) == int32(0) {
		v45 = v60
		v55 = v63
		goto L8
	} else {
		goto L29
	}
L20:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v78 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v68)+8)))
	v79 = F_DirectFunctionCall1Coll(m, l3, int32(0), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v85 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	v82 = F_pg_detoast_datum(m, base.I32_wrap_i64(v79))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v82
	goto L23
L26:
	;
	if l4 == int32(0) {
		v214 = v72
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v92 <= v91 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v96 = v90 + v91<<(uint(int32(5))%32)
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v68)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+40)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v68)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+32)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+24)) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+16)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v105 + int32(1)
	v45 = v60
	v51 = int32(0)
	v55 = v63
	goto L8
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v114 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v117 = v114
	goto L33
L31:
	;
	goto L32
L32:
	;
	v144 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v16 + int32(16)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v152 != int32(1) {
		v246 = int32(2)
		goto L3
	} else {
		goto L37
	}
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	F_pfree(m, v117)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	if v128 != 0 {
		v117 = v128
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(302776450))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v163 = F_jspOperationName(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v163
	F_errmsg(m, int32(_a_F_executeUnaryArithmExpr_0), v16)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_executeUnaryArithmExpr_1), int32(2301), int32(_a_F_executeUnaryArithmExpr_2))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v180 = v175
	goto L45
L44:
	;
	v180 = v177
	goto L45
L45:
	;
	v185 = F_palloc(m, v180<<(uint(int32(5))%32)|int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = int32(1)
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v185)+16)) = v192
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v185)+24)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v68)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v185)+32)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v68)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v185)+40)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v185
	v45 = v60
	v51 = int32(0)
	v55 = v63
	goto L8
L47:
	;
	v207 = int32(0)
	if l4 != 0 {
		v45 = v60
		v51 = v207
		v55 = v63
		goto L8
	} else {
		goto L49
	}
L48:
	;
	switch v205 {
	case 0:
		goto L47
	default:
		v45 = v60
		v55 = v63
		goto L8
	case 2:
		v214 = v205
		goto L4
	}
L49:
	;
	goto L9
L50:
	;
	v227 = v222
	goto L51
L51:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	F_pfree(m, v227)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v246 = v214
	goto L3
L53:
	;
	if v238 != 0 {
		v227 = v238
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
}
func F_exists(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_exists(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_expandNSItemAttrs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_expandNSItemVars(m, l0, l1, l2, l3, v14+int32(12))
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v24 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v27 | int64(2)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v35 = int32(0)
	v41 = int32(0)
	goto L6
L6:
	;
	v44 = int32(0)
	if v31 == v44 {
		v54 = v44
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v14 + int32(16)
	return v161
L8:
	;
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v48 <= v35 {
		v54 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v54 = v50 + v35<<(uint(int32(2))%32)
	goto L8
L11:
	;
	goto L7
L12:
	;
	v161 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if base.B2i32(v54 == int32(0))|base.B2i32(v60 <= v35) != 0 {
		v161 = v41
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v63 == int32(0) {
		v161 = v41
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v35<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v72 + int32(1)
	v78 = F_makeTargetEntry(m, v69, base.I32_extend16_s(v72), v71, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v80 = F_lappend(m, v41, v78)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+28))
	if v82 == int32(0) {
		v144 = l0
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+8)))
	F_markRTEForSelectPriv(m, v144, v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L31
	}
L20:
	;
	v86 = v82 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v82) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v93 = l0
	v95 = int32(0)
	goto L24
L22:
	;
	v117 = l0
	goto L23
L23:
	;
	v129 = v117
	v131 = int32(0)
	goto L28
L24:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v112 = v95 + int32(8)
	if v112 != v82&int32(-8) {
		v93 = v110
		v95 = v112
		goto L24
	} else {
		goto L26
	}
L25:
	;
	if v86 == int32(0) {
		v144 = v110
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v117 = v110
	goto L23
L28:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v141 = v131 + int32(1)
	if v141 != v86 {
		v129 = v139
		v131 = v141
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v144 = v139
	goto L19
L30:
	;
	goto L29
L31:
	;
	v35 = v35 + int32(1)
	v41 = v80
	goto L6
}
func F_expand_inherited_rtentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v19 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(96)
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v22 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v71 = F_table_open(m, v69, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L16
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v34 = v5
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v34<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v47 != l3 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	v66 = v34 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v66 < v67 {
		v34 = v66
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50<<(uint(int32(2))%32))))
	v55 = F_build_simple_rel(m, l0, v50, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)))
	if v57 != int32(1) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	F_expand_inherited_rtentry(m, l0, v55, v54, v50)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	goto L8
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v74 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	goto L17
L19:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= int32(0) {
		v110 = int32(0)
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v110 = int32(0)
	goto L18
L22:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v84 = v78
	goto L25
L24:
	;
	v84 = v81
	goto L25
L25:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v87 = int32(0)
	goto L26
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85+v87<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v96 == l3 {
		v110 = v95
		goto L18
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v99 = v87 + int32(1)
	if v99 != v84 {
		v87 = v99
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+32)))
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+32)) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v116 = v115
	v117 = v112
	goto L32
L31:
	;
	v116 = v5
	v117 = v5
	goto L32
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+119)))
	if v119 == int32(112) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v110 != 0 {
		goto L63
	} else {
		goto L64
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+56))
	v124 = F_getRTEPermissionInfo(m, v123, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L11
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v130 = F_find_all_inheritors(m, v69, v73, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+36))
	F_expand_partitioned_rtentry(m, l0, l1, l2, l3, v71, v126, v110, v73)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	if v130 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	F_expand_planner_arrays(m, l0, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_expand_planner_arrays(m, l0, int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L11
	} else {
		goto L62
	}
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v135 <= int32(0) {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	v144 = v5
	goto L45
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v144<<(uint(int32(2))%32))))
	if v69 != v156 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L33
L47:
	;
	v191 = v144 + int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v191 < v192 {
		v144 = v191
		goto L45
	} else {
		goto L61
	}
L48:
	;
	v159 = F_table_open(m, v156, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L11
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, v71, v110, v71, v17+int32(48), v17+int32(92))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L59
	}
L51:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, v71, v110, v159, v17+int32(48), v17+int32(92))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L11
	} else {
		goto L56
	}
L52:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v159)+48))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+118)))
	if v162 != int32(116) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+24)))
	if v165 != 0 {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	F_relation_close(m, v159, v73)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	goto L47
L56:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v175 = F_build_simple_rel(m, l0, v174, l1)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	F_relation_close(m, v159, int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v187 = F_build_simple_rel(m, l0, v186, l1)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	goto L47
L61:
	;
	goto L46
L62:
	;
	goto L33
L63:
	;
	v211 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v213 = int32(-33)
	if base.B2i32(v212&v213 == v211)|v116&v213 == v211 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	F_relation_close(m, v71, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L108
	}
L66:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v223 = int32(-1)
	v226 = int32(0)
	v228 = F_makeVar(m, v222, v223, int32(27), v223, v226, v226)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L69
	}
L67:
	;
	v263 = v211
	goto L68
L68:
	;
	v264 = int32(32)
	v266 = int32(0)
	if base.B2i32(v212&v264 == v266)|v116&v264 == v266 {
		goto L78
	} else {
		goto L79
	}
L69:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v230
	v234 = int32(32)
	v238 = F_pg_snprintf(m, v17+int32(48), v234, int32(_a_F_expand_inherited_rtentry_0), v17+v234)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v240 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240)+4)))
	v245 = v241 + int32(1)
	goto L73
L72:
	;
	v245 = int32(1)
	goto L73
L73:
	;
	v249 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	v252 = F_makeTargetEntry(m, v228, base.I32_extend16_s(v245), v249, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v255 = F_lappend(m, v254, v252)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v255
	v259 = F_lappend(m, int32(0), v228)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	v263 = v259
	goto L68
L78:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v273 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v326 = v263
	goto L80
L80:
	;
	if v117&int32(1) != 0 {
		goto L94
	} else {
		goto L95
	}
L81:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v290 = int32(0)
	v292 = F_makeWholeRowVar(m, v289, v287, v290, v290)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L11
	} else {
		goto L85
	}
L82:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v287 = v274
	v288 = v273 + v274<<(uint(int32(2))%32)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v287 = v281
	v288 = v280 + v281<<(uint(int32(2))%32) - int32(4)
	goto L81
L85:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v294
	v302 = F_pg_snprintf(m, v17+int32(48), int32(32), int32(_a_F_expand_inherited_rtentry_1), v17+int32(16))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v304 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+4)))
	v309 = v305 + int32(1)
	goto L89
L88:
	;
	v309 = int32(1)
	goto L89
L89:
	;
	v313 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v316 = F_makeTargetEntry(m, v292, base.I32_extend16_s(v309), v313, int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v319 = F_lappend(m, v318, v316)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v319
	v322 = F_lappend(m, v263, v292)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v326 = v322
	goto L80
L94:
	;
	v367 = v326
	goto L96
L95:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v333 = int32(0)
	v335 = F_makeVar(m, v329, int32(-6), int32(26), int32(-1), v333, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L97
	}
L96:
	;
	v369 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L11
	} else {
		goto L106
	}
L97:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v337
	v343 = F_pg_snprintf(m, v17+int32(48), int32(32), int32(_a_F_expand_inherited_rtentry_2), v17)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v345 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+4)))
	v350 = v346 + int32(1)
	goto L101
L100:
	;
	v350 = int32(1)
	goto L101
L101:
	;
	v354 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v357 = F_makeTargetEntry(m, v335, base.I32_extend16_s(v350), v354, int32(1))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v360 = F_lappend(m, v359, v357)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v360
	v363 = F_lappend(m, v326, v335)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	v367 = v363
	goto L96
L106:
	;
	F_add_vars_to_targetlist(m, l0, v367, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	goto L65
L108:
	;
	goto L1
}
func F_ext_sibling_callback(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ext_sibling_callback[0]))
	if v5 != 0 {
		v7 = v5
		for {
			if l2 != 0 {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				if v9 != l2 {
				} else {
					v11 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v11)
				}
			} else {
				v11 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v11)
			}
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if v13 != 0 {
				v7 = v13
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
