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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
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
	var v53 int32
	_ = v53
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v7 != 0 {
		if l2 != 0 {
			v8 = F__emscripten_memcpy_bulkmem(m, l1, v7, l2)
			mBase = m.M
		} else {
		}
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		if v12 != 0 {
			v16 = base.I32_div_s(v11+int32(7), int32(8))
			v24 = (v16 + v10<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		} else {
			v24 = int32(0)
		}
		v27 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), l2)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v10
		v30 = int32(2)
		*(*int32)(unsafe.Add(mBase, uint32(v27))) = l2 << (uint(v30) % 32)
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v33
		v36 = v27 + int32(16)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v39 = v10 << (uint(v30) % 32)
		if v39 != 0 {
			v40 = F__emscripten_memcpy_bulkmem(m, v36, v37, v39)
			mBase = m.M
			v41 = v40
		} else {
			v41 = v36
		}
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v39 != 0 {
			v44 = F__emscripten_memcpy_bulkmem(m, v41+v39, v43, v39)
			mBase = m.M
		} else {
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+44)))
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
		v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+47)))
		F_CopyArrayEls(m, v27, v46, v47, v11, v48, v49, v50, int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
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
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v350 int64
	_ = v350
	var v352 int64
	_ = v352
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v400 int64
	_ = v400
	var v402 int64
	_ = v402
	var v404 int64
	_ = v404
	var v406 int64
	_ = v406
	var v408 int64
	_ = v408
	var v414 int32
	_ = v414
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v491 int32
	_ = v491
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
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v530 int64
	_ = v530
	var v533 int32
	_ = v533
	var v534 int64
	_ = v534
	var v537 int32
	_ = v537
	var v538 int64
	_ = v538
	var v540 int64
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v588 int64
	_ = v588
	var v590 int64
	_ = v590
	var v592 int64
	_ = v592
	var v596 int32
	_ = v596
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int64
	_ = v653
	var v655 int64
	_ = v655
	var v657 int64
	_ = v657
	var v659 int64
	_ = v659
	var v661 int64
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	v9 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(48)
	m.G0 = v25
	v28 = F_palloc0(m, int32(68))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(380)
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = l6
	v45 = int32(0)
	if int64(2) <= base.I64_extend_i32_s(l4)+base.I64_extend_i32_u(base.B2i32(l7 != v45)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v54 = F_palloc(m, int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v56 = v9
	goto L5
L5:
	;
	if int32(0) < l4 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v56 = v54
	goto L5
L7:
	;
	v60 = l4 & int32(3)
	if base.Ui32(l4) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v194 = v45
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = l1
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+28)) = uint8(v211)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = l0
	v218 = v25 + int32(8)
	v222 = m.G0
	v224 = v222 - int32(16)
	m.G0 = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)) = uint8(v211)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v218)+24))
	if v229 != 0 {
		goto L44
	} else {
		goto L45
	}
L10:
	;
	if v60 != 0 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v114 = v45
	v118 = v9
	goto L10
L12:
	;
	goto L13
L13:
	;
	v71 = v45
	v75 = v9
	v78 = v9
	goto L14
L14:
	;
	v87 = base.I32_extend16_s(v75)
	v90 = l5 + v71<<(uint(int32(1))%32)
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90))))
	if v91 < v87 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v114 = v104
	v118 = v102
	goto L10
L16:
	;
	v93 = v87
	goto L18
L17:
	;
	v93 = v91
	goto L18
L18:
	;
	v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+2)))
	if v94 < v93 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v96 = v93
	goto L21
L20:
	;
	v96 = v94
	goto L21
L21:
	;
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+4)))
	if v97 < v96 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v99 = v96
	goto L24
L23:
	;
	v99 = v97
	goto L24
L24:
	;
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+6)))
	if v100 < v99 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v102 = v99
	goto L27
L26:
	;
	v102 = v100
	goto L27
L27:
	;
	v103 = int32(4)
	v104 = v71 + v103
	v106 = v78 + v103
	if v106 != l4&int32(2147483644) {
		v71 = v104
		v75 = v102
		v78 = v106
		goto L14
	} else {
		goto L28
	}
L28:
	;
	goto L15
L29:
	;
	v136 = v114
	v140 = v118
	v141 = v9
	goto L32
L30:
	;
	v174 = v118
	goto L31
L31:
	;
	v194 = v174 & int32(65535)
	goto L9
L32:
	;
	v152 = base.I32_extend16_s(v140)
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5+v136<<(uint(int32(1))%32)))))
	if v156 < v152 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v174 = v158
	goto L31
L34:
	;
	v158 = v152
	goto L36
L35:
	;
	v158 = v156
	goto L36
L36:
	;
	v159 = int32(1)
	v162 = v141 + v159
	if v162 != v60 {
		v136 = v136 + v159
		v140 = v158
		v141 = v162
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	if v311 != 0 {
		goto L66
	} else {
		goto L67
	}
L39:
	;
	m.G0 = v224 + int32(16)
	goto L38
L40:
	;
	v311 = int32(1)
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v218)+28)) = v287
	v300 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+20)) = uint8(v300)
	*(*int32)(unsafe.Add(mBase, uint32(v218)+24)) = v286
	if v287 != int32(1652372) {
		goto L40
	} else {
		goto L65
	}
L42:
	;
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v218)+20)) = uint8(v295)
	*(*int64)(unsafe.Add(mBase, uint32(v218)+24)) = int64(0)
	goto L40
L43:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)))
	if v288 != int32(1) {
		goto L42
	} else {
		goto L62
	}
L44:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)) = uint8(base.B2i32(v230 != int32(0)))
	v286 = v229
	v287 = v230
	goto L43
L45:
	;
	goto L46
L46:
	;
	if v226 == int32(0) {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	switch v236 - int32(2) {
	case 0:
		goto L50
	case 1:
		goto L49
	case 2, 3, 4:
		goto L48
	default:
		goto L42
	}
L48:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v226)+80))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v226)+76))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+100)))
	if v281 != int32(1) {
		v286 = v280
		v287 = v279
		goto L43
	} else {
		goto L61
	}
L49:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v226)+36))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+101)))
	if v260 != int32(1) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v226)+40))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+102)))
	if v240 != int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v239 == int32(0) {
		goto L42
	} else {
		goto L55
	}
L52:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+98)))
	if v243 != int32(1) {
		goto L42
	} else {
		goto L53
	}
L53:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v226)+88))
	if v246 == int32(0) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)) = uint8(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v239)+56))
	v286 = v251
	v287 = v246
	goto L43
L55:
	;
	v257 = F_ExecGetResultSlotOps(m, v239, v224+int32(15))
	mBase = m.M
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v239)+56))
	v286 = v258
	v287 = v257
	goto L43
L56:
	;
	if v259 == int32(0) {
		goto L42
	} else {
		goto L60
	}
L57:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+97)))
	if v263 != int32(1) {
		goto L42
	} else {
		goto L58
	}
L58:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v226)+84))
	if v266 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v269 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)) = uint8(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v259)+56))
	v286 = v271
	v287 = v266
	goto L43
L60:
	;
	v277 = F_ExecGetResultSlotOps(m, v259, v224+int32(15))
	mBase = m.M
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v259)+56))
	v286 = v278
	v287 = v277
	goto L43
L61:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+15)) = uint8(v284)
	v286 = v280
	v287 = v279
	goto L43
L62:
	;
	if v286 == int32(0) {
		goto L42
	} else {
		goto L63
	}
L63:
	;
	if v287 != 0 {
		goto L41
	} else {
		goto L64
	}
L64:
	;
	goto L42
L65:
	;
	v311 = int32(0)
	goto L39
L66:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v315 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	if l7 != 0 {
		goto L79
	} else {
		goto L80
	}
L69:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v337 + int32(1)
	v343 = v336 + v337*int32(40)
	v344 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+32)) = v344
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+24)) = v346
	v348 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+16)) = v348
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v343)+8)) = v350
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v352
	goto L68
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v334
	v336 = v334
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = int32(16)
	v321 = F_palloc(m, int32(640))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v323 != v315 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v334 = v321
	goto L70
L75:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v336 = v325
	goto L69
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v315 << (uint(int32(1)) % 32)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v332 = F_repalloc(m, v329, v315*int32(80))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v334 = v332
	goto L70
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(85)
	v364 = base.B2i32(int32(0) < l4)
	if int32(0) < l4 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v414 = int32(86)
	goto L81
L81:
	;
	if int32(0) < l4 {
		goto L98
	} else {
		goto L99
	}
L82:
	;
	v365 = v56 + int32(4)
	goto L84
L83:
	;
	v365 = v28 + int32(5)
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v365
	if int32(0) < l4 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v369 = v56
	goto L87
L86:
	;
	v369 = v28 + int32(8)
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v371 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v393 + int32(1)
	v399 = v392 + v393*int32(40)
	v400 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+32)) = v400
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+24)) = v402
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+16)) = v404
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v399)+8)) = v406
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v408
	v414 = int32(88)
	goto L81
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v390
	v392 = v390
	goto L88
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = int32(16)
	v377 = F_palloc(m, int32(640))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v379 != v371 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v390 = v377
	goto L89
L94:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v392 = v381
	goto L88
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v371 << (uint(int32(1)) % 32)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v388 = F_repalloc(m, v385, v371*int32(80))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v390 = v388
	goto L89
L98:
	;
	v435 = v414
	v438 = int32(0)
	goto L101
L99:
	;
	goto L100
L100:
	;
	v620 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v620
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(0)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v624 == v620 {
		goto L133
	} else {
		goto L134
	}
L101:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5+v438<<(uint(int32(1))%32)))))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l3+v438<<(uint(int32(2))%32))))
	v459 = F_palloc0(m, int32(28))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	goto L100
L103:
	;
	v461 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v459)+18)) = uint16(v461)
	v463 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v459)+16)) = uint8(v463)
	*(*int32)(unsafe.Add(mBase, uint32(v459)+12)) = v457
	*(*int64)(unsafe.Add(mBase, uint32(v459)+4)) = int64(0)
	v470 = l2 + v438*int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v459))) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v459 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v459 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = int32(7)
	v482 = base.I32_extend16_s(v453 - v461)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(88)+v484<<(uint(int32(4))%32)+v482*int32(100))))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v491
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v495 == v463 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v517 + int32(1)
	v521 = int32(40)
	v523 = v516 + v517*v521
	v525 = v25 + v521
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v525)))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+32)) = v526
	v529 = v25 + int32(32)
	v530 = *(*int64)(unsafe.Add(mBase, uint32(v529)))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+24)) = v530
	v533 = v25 + int32(24)
	v534 = *(*int64)(unsafe.Add(mBase, uint32(v533)))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+16)) = v534
	v537 = v25 + int32(16)
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v537)))
	*(*int64)(unsafe.Add(mBase, uint32(v523)+8)) = v538
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v523))) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v56
	v543 = base.B2i32(v438 == l4-int32(1))
	if v438 == l4-int32(1) {
		goto L114
	} else {
		goto L115
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v514
	v516 = v514
	goto L104
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = int32(16)
	v501 = F_palloc(m, int32(640))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v503 != v495 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v514 = v501
	goto L105
L110:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v516 = v505
	goto L104
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v495 << (uint(int32(1)) % 32)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v512 = F_repalloc(m, v509, v495*int32(80))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v514 = v512
	goto L105
L114:
	;
	v544 = v28 + int32(5)
	goto L116
L115:
	;
	v544 = v56 + int32(4)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v533))) = v470
	if v438 == l4-int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v547 = v28 + int32(8)
	goto L119
L118:
	;
	v547 = v56
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v435
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v459
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	*(*int32)(unsafe.Add(mBase, uint32(v529))) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = int32(-1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	if v555 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v578 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v577 + v578
	v583 = v576 + v577*int32(40)
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v525)))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+32)) = v584
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v529)))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+24)) = v586
	v588 = *(*int64)(unsafe.Add(mBase, uint32(v533)))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+16)) = v588
	v590 = *(*int64)(unsafe.Add(mBase, uint32(v537)))
	*(*int64)(unsafe.Add(mBase, uint32(v583)+8)) = v590
	v592 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v583))) = v592
	v596 = v438 + v578
	if v596 != l4 {
		v435 = int32(88)
		v438 = v596
		goto L101
	} else {
		goto L130
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v574
	v576 = v574
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = int32(16)
	v561 = F_palloc(m, int32(640))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v563 != v555 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v574 = v561
	goto L121
L126:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v576 = v565
	goto L120
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v555 << (uint(int32(1)) % 32)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v572 = F_repalloc(m, v569, v555*int32(80))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v574 = v572
	goto L121
L130:
	;
	goto L102
L131:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v646 + int32(1)
	v652 = v645 + v646*int32(40)
	v653 = *(*int64)(unsafe.Add(mBase, uint32(v25)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v652)+32)) = v653
	v655 = *(*int64)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v652)+24)) = v655
	v657 = *(*int64)(unsafe.Add(mBase, uint32(v25)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v652)+16)) = v657
	v659 = *(*int64)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v652)+8)) = v659
	v661 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v652))) = v661
	v663 = F_jit_compile_expr(m, v28)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L141
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v643
	v645 = v643
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = int32(16)
	v630 = F_palloc(m, int32(640))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v632 != v624 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v643 = v630
	goto L132
L137:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v645 = v634
	goto L131
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v624 << (uint(int32(1)) % 32)
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v641 = F_repalloc(m, v638, v624*int32(80))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v643 = v641
	goto L132
L141:
	;
	if v663 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_ExecReadyInterpretedExpr(m, v28)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	m.G0 = v25 + int32(48)
	return v28
L145:
	;
	goto L144
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
	v31 = base.I32_div_s(v29, int32(32))
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
	if v84&int32(65535) == int32(0) {
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
	F_errmsg_internal(m, int32(466940), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(521193), int32(780), int32(478483))
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v23 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = v23
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
	v28 = v26
	goto L3
L6:
	;
	m.G0 = v18 + int32(16)
	return
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v31 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = l0 + int32(132)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v44 != 0 {
		v47 = v43
		goto L7
	} else {
		goto L16
	}
L11:
	;
	F_ExecInitGenerated(m, l0, l1, int32(2))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v37 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v47 = l0 + int32(136)
	goto L7
L16:
	;
	F_ExecInitGenerated(m, l0, l1, l3)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v47 = v43
	goto L7
L18:
	;
	v52 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	v54 = v49
	goto L20
L20:
	;
	v55 = int32(4562080)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v58
	v61 = v22 << (uint(int32(2)) % 32)
	v62 = F_palloc(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	v54 = v52
	goto L20
L22:
	;
	v64 = F_palloc(m, v22)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v68 < v67 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_slot_getsomeattrs_int(m, l2, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v22 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L26
L28:
	;
	if int32(0) < v22 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, v64, v72, v22)
	mBase = m.M
	v74 = v73
	goto L31
L30:
	;
	v74 = v64
	goto L31
L31:
	;
	goto L28
L32:
	;
	v80 = int32(0)
	goto L35
L33:
	;
	goto L34
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	m.T0[v155].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L49
	}
L35:
	;
	v97 = v21 + int32(20) + v80<<(uint(int32(4))%32)
	v99 = v80 << (uint(int32(2)) % 32)
	v100 = v48 + v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L34
L37:
	;
	v137 = v80 + int32(1)
	if v137 != v22 {
		v80 = v137
		goto L35
	} else {
		goto L48
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l2
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v107 = m.T0[v106].(func(*base.Module, int32, int32, int32) int32)(m, v103, v28, v18+int32(15))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v74))))
	if v124 != 0 {
		goto L37
	} else {
		goto L46
	}
L41:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	if v109 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+4)))
	v114 = F_datumCopy(m, v107, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	v117 = v107
	v118 = v109
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v62))) = v117
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v74))) = uint8(v118)
	goto L37
L45:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	v117 = v114
	v118 = v116
	goto L44
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126+v99)))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+4)))
	v131 = F_datumCopy(m, v128, v129, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v62))) = v131
	goto L37
L48:
	;
	goto L36
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v61 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v22 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v159 = F__emscripten_memcpy_bulkmem(m, v158, v62, v61)
	mBase = m.M
	goto L53
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	v166 = v164 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)) = uint16(v169)
	goto L58
L55:
	;
	v162 = F__emscripten_memcpy_bulkmem(m, v161, v74, v22)
	mBase = m.M
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+28))
	m.T0[v172].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v56
	goto L6
}
func F_ExecCteScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(705), int32(706))
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
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
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
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L18
	} else {
		goto L149
	}
L8:
	;
	m.G0 = v21 + int32(32)
	return v417
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
	v417 = int32(0)
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
		v417 = int32(0)
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
	v203 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v203 < int32(2) {
		v417 = int32(0)
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
		v417 = v139
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
	v99 = m.T0[v98].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v91, l2, v93, v94, v95, int32(1), v72, l5)
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
	v417 = v139
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
	v129 = *(*int32)(unsafe.Add(mBase, _consts[134]))
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
	v417 = int32(0)
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
	F_errmsg(m, int32(448604), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_errhint(m, int32(611271), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(524080), int32(1687), int32(369193))
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
		v417 = v139
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
		v417 = v139
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+4)))
	if v157&int32(2) != 0 {
		v417 = v139
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
	F_errmsg(m, int32(448604), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	F_errhint(m, int32(611271), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(524080), int32(1761), int32(369193))
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
	F_errmsg_internal(m, int32(62964), v21+int32(16))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(524080), int32(1781), int32(369193))
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
	F_errmsg(m, int32(369055), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L18
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(524080), int32(1793), int32(369193))
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
	F_errmsg_internal(m, int32(63039), v21)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L18
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(524080), int32(1799), int32(369193))
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
	F_errmsg(m, int32(374057), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L18
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(524080), int32(1703), int32(369193))
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
	v417 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, uint32(v256)+36)) = v266
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
	if v330 == v332 {
		v417 = v332
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
	v417 = int32(0)
	goto L8
L118:
	;
	goto L119
L119:
	;
	v324 = int32(1)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+8)))
	v330 = v324
	v331 = int32(base.Ui32(v325&int32(2)) >> (uint(v324) % 32))
	goto L113
L120:
	;
	if (l4|v331)&int32(1) == int32(0) {
		v417 = v332
		goto L8
	} else {
		goto L121
	}
L121:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v340 != 0 {
		v374 = v279
		goto L122
	} else {
		goto L123
	}
L122:
	;
	if v331 != 0 {
		goto L137
	} else {
		goto L138
	}
L123:
	;
	v341 = F_ExecGetReturningSlot(m, v24, l1)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L18
	} else {
		goto L124
	}
L124:
	;
	if l3 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_ExecForceStoreHeapTuple(m, l3, v341, int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L18
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v347 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v374 = v341
	goto L122
L129:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v23)+188))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+60))
	v370 = m.T0[v369].(func(*base.Module, int32, int32, int32, int32) int32)(m, v23, l2, int32(4216240), v341)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L18
	} else {
		goto L135
	}
L130:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, _consts[116])))
	if v351&int32(1) != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L18
	} else {
		goto L132
	}
L132:
	;
	F_errmsg_internal(m, int32(353288), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L18
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(343280), int32(1264), int32(284132))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L18
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
	if v370 == int32(0) {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	v374 = v341
	goto L122
L137:
	;
	v375 = F_ExecGetChildToRootMap(m, l1)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v397 = F_ExecProcessReturning(m, l0, l1, int32(4), v374, int32(0), v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L18
	} else {
		goto L146
	}
L140:
	;
	if v375 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v374
	v417 = v332
	goto L8
L142:
	;
	goto L143
L143:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)+120))
	v383 = F_ExecGetReturningSlot(m, v24, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L18
	} else {
		goto L144
	}
L144:
	;
	v385 = F_execute_attr_map_slot(m, v380, v374, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L18
	} else {
		goto L145
	}
L145:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v374)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+36)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v374)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+28)) = v389
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v374)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v385)+32)) = uint16(v391)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v385
	v417 = v332
	goto L8
L146:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+28))
	m.T0[v400].(func(*base.Module, int32))(m, v397)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L18
	} else {
		goto L147
	}
L147:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+12))
	m.T0[v404].(func(*base.Module, int32))(m, v374)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L18
	} else {
		goto L148
	}
L148:
	;
	v417 = v397
	goto L8
L149:
	;
	F_errmsg_internal(m, int32(562727), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L18
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(524080), int32(1856), int32(369193))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
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
	var v32 int32
	_ = v32
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
	var v57 int32
	_ = v57
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
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v205 int32
	_ = v205
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
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v235 int32
	_ = v235
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
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v283 int64
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
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
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int64
	_ = v444
	var v447 int64
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v456 int64
	_ = v456
	var v458 int64
	_ = v458
	var v460 int64
	_ = v460
	var v462 int64
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
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
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int64
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
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
	switch v21 - int32(394) {
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
	v679 = m.ExcPending
	if v679 != 0 {
		goto L3
	} else {
		goto L335
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L3
	} else {
		goto L332
	}
L11:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L3
	} else {
		goto L331
	}
L12:
	;
	F_EvalPlanQualEnd(m, l0+int32(108))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L3
	} else {
		goto L329
	}
L13:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v646 != 0 {
		goto L323
	} else {
		goto L324
	}
L14:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L3
	} else {
		goto L322
	}
L15:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L3
	} else {
		goto L321
	}
L16:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v588 != 0 {
		goto L301
	} else {
		goto L302
	}
L17:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v475 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L18:
	;
	F_ExecEndGroup(m, l0)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L3
	} else {
		goto L251
	}
L19:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v437 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L20:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	F_ExecDropSingleTupleTableSlot(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L3
	} else {
		goto L232
	}
L21:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v410 != 0 {
		goto L227
	} else {
		goto L228
	}
L22:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v402 != 0 {
		goto L222
	} else {
		goto L223
	}
L23:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v391 != 0 {
		goto L216
	} else {
		goto L217
	}
L24:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L214
	}
L25:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L212
	}
L26:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	m.T0[v376].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L211
	}
L27:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+80))
	if v359 == int32(1) {
		goto L202
	} else {
		goto L203
	}
L28:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if l0 == v351 {
		goto L197
	} else {
		goto L198
	}
L29:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v346 != 0 {
		goto L193
	} else {
		goto L194
	}
L30:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v312 < v313 {
		goto L183
	} else {
		goto L184
	}
L31:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	F_ExecEndNode(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L182
	}
L32:
	;
	F_ExecEndSeqScan(m, l0)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L3
	} else {
		goto L181
	}
L33:
	;
	F_ExecEndSeqScan(m, l0)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L180
	}
L34:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v264 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L35:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v241 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L36:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v211 != 0 {
		goto L138
	} else {
		goto L139
	}
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v186 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L38:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L114
	}
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L101
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+32))
	if v143 != 0 {
		goto L93
	} else {
		goto L94
	}
L41:
	;
	F_ExecEndSeqScan(m, l0)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L92
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
	v32 = v2
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
	v39 = v36 + v32*int32(216)
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
	v57 = int32(0)
	goto L63
L61:
	;
	goto L62
L62:
	;
	v86 = v32 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v86 < v87 {
		v32 = v86
		goto L53
	} else {
		goto L68
	}
L63:
	;
	v63 = v57 << (uint(int32(2)) % 32)
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
	v75 = v57 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v39)+100))
	if v75 < v76 {
		v57 = v75
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
	m.T0[v143].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v146 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+188))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	m.T0[v149].(func(*base.Module, int32))(m, v146)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L1
L100:
	;
	goto L99
L101:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v155 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_ExecParallelFinish(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L3
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v158 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	F_pfree(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(0)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v163 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	F_ExecParallelCleanup(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L1
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L112
L114:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v171 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	F_ExecParallelFinish(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v174 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	F_pfree(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v179 != 0 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	F_ExecParallelCleanup(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L3
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L1
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	goto L125
L127:
	;
	if v184 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v190 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v197 = v186 + v190<<(uint(int32(3))%32) + int32(8)
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v197)))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v197))) = v198 + v199
	goto L127
L130:
	;
	F_index_endscan(m, v184)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	if v185 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	F_relation_close(m, v185, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L3
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	goto L1
L137:
	;
	goto L136
L138:
	;
	F_ReleaseBuffer(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L3
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v216 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	goto L140
L142:
	;
	if v209 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v220 < int32(0) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v227 = v216 + v220<<(uint(int32(3))%32) + int32(8)
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v227)))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = v228 + v229
	goto L142
L145:
	;
	F_index_endscan(m, v209)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if v210 != 0 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	goto L147
L149:
	;
	F_relation_close(m, v210, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L3
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	goto L1
L152:
	;
	goto L151
L153:
	;
	if v239 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v245 < int32(0) {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v252 = v241 + v245<<(uint(int32(3))%32) + int32(8)
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v252)))
	v254 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v252))) = v253 + v254
	goto L153
L156:
	;
	F_index_endscan(m, v239)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L3
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	if v240 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L158
L160:
	;
	F_relation_close(m, v240, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	goto L1
L163:
	;
	goto L162
L164:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L3
	} else {
		goto L167
	}
L165:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v268 < int32(0) {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v273 = v264 + v268<<(uint(int32(4))%32)
	v275 = v273 + int32(8)
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v275)))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v275))) = v276 + v277
	v281 = v273 + int32(16)
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v281)))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(l0)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v281))) = v282 + v283
	goto L164
L167:
	;
	if v288 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288)+20))
	if v292 != 0 {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	goto L170
L170:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v302 != 0 {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_tbm_end_iterate(m, v288+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L3
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+188))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	m.T0[v299].(func(*base.Module, int32))(m, v288)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L175
	}
L174:
	;
	goto L173
L175:
	;
	goto L170
L176:
	;
	F_tbm_free(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L3
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	goto L1
L179:
	;
	goto L178
L180:
	;
	goto L1
L181:
	;
	goto L1
L182:
	;
	goto L1
L183:
	;
	v317 = v312
	v320 = v313
	goto L186
L184:
	;
	goto L185
L185:
	;
	goto L1
L186:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v328 = v323 + v317<<(uint(int32(5))%32) + int32(12)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	if v329 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L185
L188:
	;
	F_tuplestore_end(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L3
	} else {
		goto L191
	}
L189:
	;
	v335 = v320
	goto L190
L190:
	;
	v337 = v317 + int32(1)
	if v337 < v335 {
		v317 = v337
		v320 = v335
		goto L186
	} else {
		goto L192
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v335 = v334
	goto L190
L192:
	;
	goto L187
L193:
	;
	F_tuplestore_end(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L3
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = int32(0)
	goto L1
L196:
	;
	goto L195
L197:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_tuplestore_end(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	goto L1
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = int32(0)
	goto L199
L201:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v372 != 0 {
		goto L207
	} else {
		goto L208
	}
L202:
	;
	v366 = int32(28)
	goto L204
L203:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+156))
	if v364 != 0 {
		goto L201
	} else {
		goto L205
	}
L204:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366+v367)))
	m.T0[v369].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L3
	} else {
		goto L206
	}
L205:
	;
	v366 = int32(100)
	goto L204
L206:
	;
	goto L201
L207:
	;
	F_ExecEndNode(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	goto L1
L210:
	;
	goto L209
L211:
	;
	goto L1
L212:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L3
	} else {
		goto L213
	}
L213:
	;
	goto L1
L214:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L215
	}
L215:
	;
	goto L1
L216:
	;
	F_ExecHashTableDestroy(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L220
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	goto L218
L220:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L221
	}
L221:
	;
	goto L1
L222:
	;
	F_tuplestore_end(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = int32(0)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L3
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	goto L1
L227:
	;
	F_tuplesort_end(m, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L3
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L3
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	goto L1
L232:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	F_ExecDropSingleTupleTableSlot(m, v421)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L3
	} else {
		goto L233
	}
L233:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v424 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_tuplesort_end(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L3
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v429 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	goto L236
L238:
	;
	F_tuplesort_end(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L3
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L3
	} else {
		goto L242
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L240
L242:
	;
	goto L1
L243:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	F_MemoryContextDelete(m, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L3
	} else {
		goto L249
	}
L244:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v441 < int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v444 = *(*int64)(unsafe.Add(mBase, uint32(l0)+232))
	if v444 == int64(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v447 = *(*int64)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+232)) = v447
	goto L248
L247:
	;
	goto L248
L248:
	;
	v451 = v437 + v441*int32(40)
	v453 = l0 + int32(200)
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v453)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+40)) = v454
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v453)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+32)) = v456
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v453)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+24)) = v458
	v460 = *(*int64)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+16)) = v460
	v462 = *(*int64)(unsafe.Add(mBase, uint32(v453)))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+8)) = v462
	goto L243
L249:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L3
	} else {
		goto L250
	}
L250:
	;
	goto L1
L251:
	;
	goto L1
L252:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v493 != 0 {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v479 < int32(0) {
		goto L252
	} else {
		goto L254
	}
L254:
	;
	v484 = v475 + v479*int32(24)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+24)) = v485
	v487 = *(*int64)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v484)+16)) = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(v484)+8)) = v489
	goto L252
L255:
	;
	F_tuplesort_end(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L3
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v496 != 0 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	F_tuplesort_end(m, v496)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	F_hashagg_reset_spill_state(m, l0)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L3
	} else {
		goto L263
	}
L262:
	;
	goto L261
L263:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	if v501 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	F_MemoryContextDelete(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v506 != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = int32(0)
	goto L266
L268:
	;
	F_MemoryContextDelete(m, v506)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L3
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v511 <= int32(0) {
		goto L273
	} else {
		goto L274
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = int32(0)
	goto L270
L272:
	;
	v566 = int32(0)
	goto L292
L273:
	;
	v514 = int32(1)
	if v474 <= v514 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	v518 = int32(1)
	if v474 <= v518 {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v517 = v514
	goto L278
L277:
	;
	v517 = v474
	goto L278
L278:
	;
	v560 = v517
	goto L272
L279:
	;
	v521 = v518
	goto L281
L280:
	;
	v521 = v474
	goto L281
L281:
	;
	v526 = v2
	goto L282
L282:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v537 = int32(0)
	goto L284
L283:
	;
	v560 = v521
	goto L272
L284:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v529+v526*int32(224)+int32(208))))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v543+v537<<(uint(int32(2))%32))))
	if v547 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v554 = v526 + int32(1)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v554 < v555 {
		v526 = v554
		goto L282
	} else {
		goto L291
	}
L286:
	;
	F_tuplesort_end(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L3
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	v551 = v537 + int32(1)
	if v551 != v521 {
		v537 = v551
		goto L284
	} else {
		goto L290
	}
L289:
	;
	goto L288
L290:
	;
	goto L285
L291:
	;
	goto L283
L292:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v566<<(uint(int32(2))%32))))
	F_ReScanExprContext(m, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L3
	} else {
		goto L294
	}
L293:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v582 != 0 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v580 = v566 + int32(1)
	if v580 != v560 {
		v566 = v580
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	F_ReScanExprContext(m, v582)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L3
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v585)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L300
	}
L299:
	;
	goto L298
L300:
	;
	goto L1
L301:
	;
	F_tuplestore_end(m, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L3
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	F_release_partition(m, l0)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L305
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L303
L305:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v595 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v600 = int32(0)
	v601 = v595
	goto L309
L307:
	;
	goto L308
L308:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	F_MemoryContextDelete(m, v627)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L3
	} else {
		goto L316
	}
L309:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606+v600*int32(160))+128))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v610 != v611 {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	goto L308
L311:
	;
	F_MemoryContextDelete(m, v610)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L3
	} else {
		goto L314
	}
L312:
	;
	v616 = v601
	goto L313
L313:
	;
	v618 = v600 + int32(1)
	if v618 < v616 {
		v600 = v618
		v601 = v616
		goto L309
	} else {
		goto L315
	}
L314:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v616 = v615
	goto L313
L315:
	;
	goto L310
L316:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	F_MemoryContextDelete(m, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L317
	}
L317:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	F_pfree(m, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L318
	}
L318:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_pfree(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L3
	} else {
		goto L319
	}
L319:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L3
	} else {
		goto L320
	}
L320:
	;
	goto L1
L321:
	;
	goto L1
L322:
	;
	goto L1
L323:
	;
	F_MemoryContextDelete(m, v646)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L3
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v649)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L3
	} else {
		goto L327
	}
L326:
	;
	goto L325
L327:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ExecEndNode(m, v652)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L3
	} else {
		goto L328
	}
L328:
	;
	goto L1
L329:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ExecEndNode(m, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L330
	}
L330:
	;
	goto L1
L331:
	;
	goto L1
L332:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v668
	F_errmsg_internal(m, int32(509011), v10)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L3
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(524470), int32(760), int32(434096))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L3
	} else {
		goto L334
	}
L334:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L335:
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
	var v80 int32
	_ = v80
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v8 != 0 {
		v9 = v8
	} else {
		v9 = l0
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(0) {
		v80 = v3
		return v80
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
				v80 = v3
				return v80
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
				if v28 == int32(0) {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
					v80 = v77
					return v80
				} else {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)))
					if v31 == int32(0) {
						v34 = int32(4562080)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[28]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v41
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
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v35
									v56 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									if v63 == int32(0) {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
										v80 = v77
										return v80
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
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v35
								v56 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v56)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								if v63 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
									v80 = v77
									return v80
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
							v80 = v77
							return v80
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
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(85899346362)
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
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(94489280954)
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
func F_ExecInterpExprStillValid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v4 < v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
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
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v213
	v215 = m.T0[v213].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L50
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
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
	v58 = F_bsearch(m, v18+v52, int32(4459344), int32(120), v52, int32(588))
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
	return int32(0)
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v64 = v62
	goto L6
L12:
	;
	v192 = v36 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
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
	v86 = v78 + v79<<(uint(int32(4))%32) + v70*int32(100)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+10)))
	if v87 == int32(118) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v91 = v86 - int32(80)
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
	F_errmsg_internal(m, int32(491686), v74)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(520199), int32(2404), int32(11819))
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
	F_errmsg_internal(m, int32(436654), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(520199), int32(2410), int32(11819))
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
	F_errmsg(m, int32(473282), v72+int32(-48))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(520199), int32(2416), int32(11819))
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
	F_errmsg(m, int32(387743), v72+int32(-16))
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
	F_errdetail(m, int32(632116), v72+int32(-32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(520199), int32(2425), int32(11819))
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
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int64
	_ = v200
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
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
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L138
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L135
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L132
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
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
	v275 = m.ExcPending
	if v275 != 0 {
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
	v204 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	v205 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v204 <= v205+int64(1) {
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
		v75 = v58
		v76 = v57
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
	v79 = v75 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v79
	if v79 <= v76 {
		goto L27
	} else {
		goto L43
	}
L40:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+128))
	if v62-int64(1) != v58-v57 {
		v75 = v58
		v76 = v57
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
	v75 = v73
	v76 = v72
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
	v200 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v200 + int64(1)
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
	v180 = int32(4562080)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v183
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	v188 = m.T0[v187].(func(*base.Module, int32, int32, int32) int32)(m, v174, v14, v12+int32(15))
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
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v181
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
	if v188 == int32(0) {
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
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v211 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v215 = m.T0[v214].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	if v215 == int32(0) {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+4)))
	if v219&int32(2) != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v225 = *(*int64)(unsafe.Add(mBase, uint32(l0)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+144)) = v225 - int64(1)
	v290 = v215
	goto L10
L106:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v231 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v235 = m.T0[v234].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	if v235 == int32(0) {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+4)))
	if v239&int32(2) != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v235
	v290 = v235
	goto L10
L114:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v247 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v250 != 0 {
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
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v290 = v266
	goto L10
L118:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v254 = m.T0[v253].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v254 == int32(0) {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+4)))
	if v258&int32(2) != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v254
	v290 = v254
	goto L10
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = int32(3)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v290 = v271
	goto L10
L126:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v276
	F_errmsg_internal(m, int32(506946), v12)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(517045), int32(336), int32(108054))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	F_errmsg_internal(m, int32(182588), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(517045), int32(211), int32(108054))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
	F_errmsg_internal(m, int32(182588), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(517045), int32(269), int32(108054))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
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
	F_errmsg_internal(m, int32(182588), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(517045), int32(286), int32(108054))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
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
	F_errmsg_internal(m, int32(182588), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(517045), int32(305), int32(108054))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v11 - int32(394) {
	case 0:
		v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		if v200 != 0 {
			F_ExecMarkPos(m, v200)
			mBase = m.M
			v202 = m.ExcPending
			if v202 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v205 = F_errstart(m, int32(13), int32(0))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return
			} else {
				if v205 != 0 {
					F_errmsg_internal(m, int32(382328), int32(0))
					mBase = m.M
					v210 = m.ExcPending
					if v210 != 0 {
						return
					} else {
						F_errfinish(m, int32(516964), int32(153), int32(145446))
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	default:
		v218 = F_errstart(m, int32(13), int32(0))
		mBase = m.M
		v219 = m.ExcPending
		if v219 != 0 {
			return
		} else {
			if v218 == int32(0) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v222
				F_errmsg_internal(m, int32(509011), v9)
				mBase = m.M
				v226 = m.ExcPending
				if v226 != 0 {
					return
				} else {
					F_errfinish(m, int32(522710), int32(357), int32(145482))
					mBase = m.M
					v231 = m.ExcPending
					if v231 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	case 11:
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
		if v15 == int32(0) {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			F_index_markpos(m, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
			v21 = v19 - int32(1)
			v23 = v21 << (uint(int32(2)) % 32)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v23+v24)))
			if v26 == int32(0) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23)))
				if v31 == int32(0) {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
					F_index_markpos(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v21))))
					if v36 != 0 {
						m.G0 = v9 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(333637), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errfinish(m, int32(521309), int32(860), int32(145429))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
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
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v21))))
				if v36 != 0 {
					m.G0 = v9 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(333637), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(521309), int32(860), int32(145429))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+156))
		if v58 == int32(0) {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
			F_index_markpos(m, v95)
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+72))
			v64 = v62 - int32(1)
			v66 = v64 << (uint(int32(2)) % 32)
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v67)))
			if v69 == int32(0) {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v58)+36))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v72+v66)))
				if v74 == int32(0) {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
					F_index_markpos(m, v95)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v58)+40))
					v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v64))))
					if v79 != 0 {
						m.G0 = v9 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(333585), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_errfinish(m, int32(521267), int32(479), int32(145408))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
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
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v58)+40))
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v64))))
				if v79 != 0 {
					m.G0 = v9 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(333585), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errfinish(m, int32(521267), int32(479), int32(145408))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
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
		v100 = m.G0
		v102 = v100 - int32(16)
		m.G0 = v102
		v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
		v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+20))
		if v105 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
					*(*int32)(unsafe.Add(mBase, uint32(v102))) = v116
					F_errmsg(m, int32(145494), v102)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						F_errfinish(m, int32(522052), int32(145), int32(145464))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
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
			m.T0[v105].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				m.G0 = v102 + int32(16)
				m.G0 = v9 + int32(16)
				return
			}
		}
	case 30:
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
		if v131 != 0 {
			F_tuplestore_copy_read_pointer(m, v131, int32(0), int32(1))
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				F_tuplestore_trim(m, v136)
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	case 32:
		v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
		if v139 == int32(1) {
			v142 = int32(4562080)
			v143 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
			v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+28))
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v146
			v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+64))
			switch v148 - int32(3) {
			case 0:
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v145)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v145)+224)) = v186
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+208)))
				*(*uint8)(unsafe.Add(mBase, uint32(v145)+228)) = uint8(v191)
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v143
				m.G0 = v9 + int32(16)
				return
			case 1:
				v155 = *(*int32)(unsafe.Add(mBase, uint32(v145)+200))
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+40))
				if v156 == int32(0) {
					v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+44))
					v160 = F_palloc(m, v159)
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v155)+40)) = v160
						*(*int64)(unsafe.Add(mBase, uint32(v155)+52)) = int64(0)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v155)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v155)+24)) = v165
						v167 = F_ltsReadFillBuffer(m, v155)
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return
						} else {
							v169 = *(*int64)(unsafe.Add(mBase, uint32(v155)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v145+int32(216)))) = v169
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v155)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v145+int32(224)))) = v171
							v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+208)))
							*(*uint8)(unsafe.Add(mBase, uint32(v145)+228)) = uint8(v191)
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v143
							m.G0 = v9 + int32(16)
							return
						}
					}
				} else {
					v169 = *(*int64)(unsafe.Add(mBase, uint32(v155)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v145+int32(216)))) = v169
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v155)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v145+int32(224)))) = v171
					v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+208)))
					*(*uint8)(unsafe.Add(mBase, uint32(v145)+228)) = uint8(v191)
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v143
					m.G0 = v9 + int32(16)
					return
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(369759), int32(0))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return
					} else {
						F_errfinish(m, int32(516387), int32(2454), int32(145155))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
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
			m.G0 = v9 + int32(16)
			return
		}
	}
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v15 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L6
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_ExecReScan(m, v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v34 = m.T0[v33].(func(*base.Module, int32) int32)(m, v21)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L15
	}
L11:
	;
	goto L10
L12:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_MemoryContextReset(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L29
	}
L13:
	;
	m.G0 = v11 + int32(16)
	return v77
L14:
	;
	if v20 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	if v34 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v36&int32(2) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	m.T0[v42].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v77 = int32(0)
	goto L13
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+32))
	m.T0[v74].(func(*base.Module, int32, int32))(m, v20, v34)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L28
	}
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v48&int32(2) != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v34
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v53 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v56 = int32(4562080)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v59
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
	v64 = m.T0[v63].(func(*base.Module, int32, int32, int32) int32)(m, v53, v13, v11+int32(15))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v57
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_MemoryContextReset(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if v64 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v77 = v20
	goto L13
L29:
	;
	goto L6
}
func F_ExecUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
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
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v425 int64
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = int64(0)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v19 + int32(48)
	return v494
L2:
	;
	if l6 != 0 {
		goto L119
	} else {
		goto L120
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L116
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L8
	} else {
		goto L112
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	m.T0[v30].(func(*base.Module, int32))(m, l5)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L109
	}
L8:
	;
	return int32(0)
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+116)))
	if v36 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v39 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_ExecOpenIndices(m, l1, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v81 != 0 {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+13)))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+188))
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v71 = v43
	goto L18
L18:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+15)))
	if v73 != int32(1) {
		goto L14
	} else {
		goto L26
	}
L19:
	;
	F_ExecPendingInserts(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L22
	}
L20:
	;
	v54 = v49
	goto L21
L21:
	;
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+104))
	v64 = F_ExecBRUpdateTriggers(m, v54, v56, l1, l2, l3, l5, v55, l0+int32(16), base.B2i32(v61 == int32(5)))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v54 = v53
	goto L21
L23:
	;
	if v64 == int32(0) {
		v494 = v55
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v68 == int32(0) {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v71 = v68
	goto L18
L26:
	;
	v76 = F_ExecIRUpdateTriggers(m, v22, l1, l3, l5)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	if v76 != 0 {
		v413 = l4
		v414 = l5
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v494 = int32(0)
	goto L1
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+36)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v86 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v106 = v19 + int32(28)
	v108 = l2 + int32(4)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v109)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v111
	v115 = F_ExecUpdateAct(m, l0, l1, l2, l3, l5, l6, v19+int32(32))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L8
	} else {
		goto L40
	}
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+64))
	v98 = m.T0[v97].(func(*base.Module, int32, int32, int32, int32) int32)(m, v22, l1, l5, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L36
	}
L33:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+17)))
	if v89 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	F_ExecComputeStoredGenerated(m, l1, v22, l5, int32(2))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v98 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v494 = int32(0)
	goto L1
L38:
	;
	goto L39
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+36)) = v103
	v413 = l4
	v414 = v98
	goto L2
L40:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)))
	if v117 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L8
	} else {
		goto L106
	}
L42:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if v335 < int32(2) {
		v494 = int32(0)
		goto L1
	} else {
		goto L101
	}
L43:
	;
	v126 = l4
	v127 = l5
	v129 = v115
	goto L46
L44:
	;
	goto L45
L45:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v494 = v332
	goto L1
L46:
	;
	if v129 != int32(3) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	switch v129 {
	case 0:
		v413 = v126
		v414 = v127
		goto L2
	default:
		goto L41
	case 2:
		goto L51
	case 4:
		goto L42
	}
L49:
	;
	goto L50
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	if int32(2) <= v165 {
		goto L4
	} else {
		goto L58
	}
L51:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	if v141 == v142 {
		v494 = int32(0)
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(448694), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errhint(m, int32(611271), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(524080), int32(2569), int32(374581))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v170 = F_EvalPlanQualSlot(m, v168, v21, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	v172 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v21)+188))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+104))
	v180 = m.T0[v179].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v21, l2, v173, v170, v174, v175, v172, int32(2), l0+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L8
	} else {
		goto L63
	}
L60:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v276)+12)) = v186
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+72))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	m.T0[v282].(func(*base.Module, int32))(m, v280)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L8
	} else {
		goto L97
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L94
	}
L62:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	if v237 == v238 {
		v494 = v172
		goto L1
	} else {
		goto L88
	}
L63:
	;
	if v180 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	switch v180 - int32(2) {
	case 0:
		goto L62
	default:
		goto L61
	case 2:
		v494 = v172
		goto L1
	}
L65:
	;
	goto L66
L66:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v186 = F_EvalPlanQual(m, v184, v21, v185, v170)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L67
	}
L67:
	;
	if v186 == int32(0) {
		v494 = v172
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+4)))
	if v190&int32(2) != 0 {
		v494 = v172
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v193 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ExecInitUpdateProjection(m, v196, l1)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+49)))
	if v199 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	F_UnlockTuple(m, v21, v19+int32(24), int32(7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v212 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v212 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	F_LockTuple(m, v21, l2, int32(7))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _consts[116])))
	if v214&int32(1) == int32(0) {
		goto L3
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v21)+188))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+60))
	v222 = m.T0[v221].(func(*base.Module, int32, int32, int32, int32) int32)(m, v21, l2, int32(4216240), v210)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L8
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	if v222 != 0 {
		goto L60
	} else {
		goto L84
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L85
	}
L85:
	;
	F_errmsg_internal(m, int32(470006), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(524080), int32(2633), int32(374581))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(450))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(448694), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
	;
	F_errhint(m, int32(611271), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(524080), int32(2659), int32(374581))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v180
	F_errmsg_internal(m, int32(62964), v19+int32(16))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(524080), int32(2665), int32(374581))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
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
	v285 = int32(4562080)
	v286 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v279)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v288
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v294 = m.T0[v293].(func(*base.Module, int32, int32, int32) int32)(m, v275+int32(4), v279, int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v286
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+4)))
	v300 = v298 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+4)) = uint16(v300)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+6)) = uint16(v303)
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108))))
	*(*uint16)(unsafe.Add(mBase, uint32(v106))) = uint16(v305)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v307
	v311 = F_ExecUpdateAct(m, l0, l1, l2, l3, v280, l6, v19+int32(32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+32)))
	if v313 == int32(0) {
		v126 = v210
		v127 = v280
		v129 = v311
		goto L46
	} else {
		goto L100
	}
L100:
	;
	goto L47
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(369055), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(524080), int32(2676), int32(374581))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
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
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v129
	F_errmsg_internal(m, int32(63118), v19)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(524080), int32(2682), int32(374581))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(250154), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(524080), int32(2474), int32(374581))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(374057), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(524080), int32(2585), int32(374581))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(353288), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(343280), int32(1264), int32(284132))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v22)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v425 + int64(1)
	goto L121
L120:
	;
	goto L121
L121:
	;
	F_ExecUpdateEpilogue(m, l0, v19+int32(32), l1, l2, l3, v414)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v433 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v494 = int32(0)
	goto L1
L124:
	;
	goto L125
L125:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v433)+72))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+12)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v414
	if v413 != 0 {
		v452 = int32(0)
		v453 = v413
		goto L126
	} else {
		goto L127
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+60)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v438)+56)) = v453
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+8)))
	v459 = v456&int32(231) | v452
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+8)) = uint8(v459)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v433)+72))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v433)+16))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+8))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	m.T0[v464].(func(*base.Module, int32))(m, v462)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L8
	} else {
		goto L130
	}
L127:
	;
	v443 = int32(8)
	v444 = int32(0)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+8)))
	if v445&int32(2) == v444 {
		v452 = v443
		v453 = v444
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v450 = F_ExecGetAllNullSlot(m, v437, l1)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	v452 = v443
	v453 = v450
	goto L126
L130:
	;
	v467 = int32(4562080)
	v468 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v461)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v470
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v433)+24))
	v476 = m.T0[v475].(func(*base.Module, int32, int32, int32) int32)(m, v433+int32(4), v461, int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v468
	v480 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v462)+4)))
	v482 = v480 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v462)+4)) = uint16(v482)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	*(*uint16)(unsafe.Add(mBase, uint32(v462)+6)) = uint16(v485)
	v494 = v462
	goto L1
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
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
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
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
	var v527 int32
	_ = v527
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
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
	return v564
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	if v43 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	v564 = v559
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
	v89 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v89
	*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = int64(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+132))
	if v98 != int32(2) {
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
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v20)+188))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+100))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v20, l2, v28, v75, v76, v77, int32(1), l0+int32(16), l6+int32(8), l6+int32(4))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v564 = v87
	goto L1
L25:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+104))
	if v556 != int32(5) {
		v28 = v551
		goto L2
	} else {
		goto L136
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L7
	} else {
		goto L130
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L7
	} else {
		goto L127
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L124
	}
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+120))
	if v101 == l1 {
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
	v469 = m.ExcPending
	if v469 != 0 {
		goto L7
	} else {
		goto L119
	}
L32:
	;
	F_ExecPartitionCheckEmitError(m, l1, v28, v92)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v91)+200))
	if v105 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v108 = int32(4562080)
	v109 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+120))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v92)+100))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v113
	v115 = F_ExecSetupPartitionTupleRouting(m, v92, v111)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v126 = int32(0)
	v135 = F_ExecDelete(m, l0, l1, l2, l3, v126, int32(1), v126, v18+int32(36), v18+int32(47), v18+int32(40))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+200)) = v115
	v119 = F_table_slot_create(m, v111, int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+196)) = v119
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v109
	goto L38
L41:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+47)))
	if v137 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v226)
	v228 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v229 == v228 {
		v564 = v228
		goto L1
	} else {
		goto L70
	}
L43:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v223 != 0 {
		v551 = int32(0)
		goto L25
	} else {
		goto L69
	}
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v91)+104))
	if v140 == int32(5) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v202 = F_ExecGetChildToRootMap(m, l1)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L62
	}
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v143 == int32(0) {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+4)))
	if v146&int32(2) != 0 {
		goto L42
	} else {
		goto L49
	}
L49:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v149 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	F_ExecInitUpdateProjection(m, v91, l1)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L7
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v156 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v156 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _consts[116])))
	if v158&int32(1) == int32(0) {
		goto L28
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v163)+188))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+60))
	v167 = m.T0[v166].(func(*base.Module, int32, int32, int32, int32) int32)(m, v163, l2, int32(4216240), v154)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	if v167 == int32(0) {
		goto L27
	} else {
		goto L59
	}
L59:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v171
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172)+72))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	m.T0[v179].(func(*base.Module, int32))(m, v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v182 = int32(4562080)
	v183 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v172)+24))
	v191 = m.T0[v190].(func(*base.Module, int32, int32, int32) int32)(m, v172+int32(4), v176, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v183
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+4)))
	v197 = v195 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+4)) = uint16(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+6)) = uint16(v200)
	v551 = v177
	goto L25
L62:
	;
	if v202 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v91)+196))
	v206 = F_execute_attr_map_slot(m, v204, v28, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	v208 = v28
	goto L65
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v91)+120))
	v214 = F_ExecInsert(m, l0, v209, v208, l5, v18+int32(32), v18+int32(28))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L67
	}
L66:
	;
	v208 = v206
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v214
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v91)+204))
	if v217 == int32(0) {
		goto L42
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = int32(0)
	goto L42
L69:
	;
	goto L42
L70:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v232 == int32(0) {
		v564 = v228
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+14)))
	if v235 != int32(1) {
		v564 = v228
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v241 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243)+131)))
	if v244 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v458 = int32(0)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_ExecARUpdateTriggers(m, v459, v238, l1, v229, l2, v458, v239, v458, v458, int32(1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L7
	} else {
		goto L118
	}
L74:
	;
	if v321 == int32(0) {
		goto L73
	} else {
		goto L97
	}
L75:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v245 == int32(0) {
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L94
	}
L78:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+200))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+8))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v249)+56))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v242)+56))
	v252 = F_get_partition_ancestors(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L7
	} else {
		goto L82
	}
L79:
	;
	v321 = v245
	goto L80
L80:
	;
	goto L74
L81:
	;
	v315 = F_lappend(m, v309, v248)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L93
	}
L82:
	;
	if v252 == int32(0) {
		v309 = v241
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v256 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v257 <= v256 {
		v309 = v241
		goto L81
	} else {
		goto L84
	}
L84:
	;
	v263 = v256
	v269 = v241
	goto L85
L85:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v263<<(uint(int32(2))%32))))
	if v279 == v250 {
		v309 = v269
		goto L81
	} else {
		goto L87
	}
L86:
	;
	v309 = v294
	goto L81
L87:
	;
	v282 = F_table_open(m, v279, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v285 = F_palloc0(m, int32(216))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285))) = int32(388)
	v289 = int32(0)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v240)+132))
	F_InitResultRelInfo(m, v285, v282, v289, v289, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v294 = F_lappend(m, v269, v285)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v297 = v263 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v297 < v298 {
		v263 = v297
		v269 = v294
		goto L85
	} else {
		goto L92
	}
L92:
	;
	goto L86
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+212)) = v315
	v321 = v315
	goto L80
L94:
	;
	F_errmsg_internal(m, int32(276310), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(521193), int32(1441), int32(162925))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
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
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	if v348 <= int32(0) {
		goto L73
	} else {
		goto L98
	}
L98:
	;
	v362 = v348
	v364 = int32(0)
	goto L99
L99:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v364<<(uint(int32(2))%32))))
	if v371 == v238 {
		v435 = v362
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L73
L101:
	;
	v441 = v364 + int32(1)
	if v441 < v435 {
		v362 = v435
		v364 = v441
		goto L99
	} else {
		goto L117
	}
L102:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v371)+52))
	if v373 == int32(0) {
		v435 = v362
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+14)))
	if v376 != int32(1) {
		v435 = v362
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v379 = int32(0)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v380 <= v379 {
		v435 = v362
		goto L101
	} else {
		goto L105
	}
L105:
	;
	v387 = v379
	v392 = v380
	goto L106
L106:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v401 = v398 + v387*int32(60)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+16)))
	if v402 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v435 = v424
	goto L101
L108:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401)+8))
	v408 = v405 - int32(1644)
	if base.Ui32(v408) <= base.Ui32(int32(11)) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v420 = v392
	goto L110
L110:
	;
	v422 = v387 + int32(1)
	if v422 < v420 {
		v387 = v422
		v392 = v420
		goto L106
	} else {
		goto L116
	}
L111:
	;
	if v416 == int32(1) {
		goto L26
	} else {
		goto L115
	}
L112:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v408<<(uint(int32(2))%32))+uint32(_consts[507])))
	v416 = v415
	goto L114
L113:
	;
	v416 = int32(0)
	goto L114
L114:
	;
	goto L111
L115:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	v420 = v419
	goto L110
L116:
	;
	goto L107
L117:
	;
	goto L100
L118:
	;
	v564 = v458
	goto L1
L119:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(281265), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	F_errdetail(m, int32(665601), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(524080), int32(1960), int32(374511))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
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
	F_errmsg_internal(m, int32(353288), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(343280), int32(1264), int32(284132))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
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
	F_errmsg_internal(m, int32(470006), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(524080), int32(2049), int32(374511))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
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
	v518 = m.ExcPending
	if v518 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(22521), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+48))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+48))
	v527 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v526 + v527
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v524 + v527
	F_errdetail(m, int32(695684), v18+int32(16))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v539 + int32(4)
	F_errhint(m, int32(697714), v18)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(524080), int32(2422), int32(23044))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
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
	if v11 == v13 {
		v54 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	if v3 == int32(0) {
		v54 = v13
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if v22 < v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v22
	goto L10
L9:
	;
	v25 = v23
	goto L10
L10:
	;
	if v25 <= int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v28 = int32(1)
	goto L13
L12:
	;
	v28 = v25
	goto L13
L13:
	;
	v29 = int32(8)
	v34 = int32(0)
	goto L14
L14:
	;
	v41 = v34 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v3+v29+v41)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+(v11+v29))))
	v46 = v43 & v45
	v48 = base.B2i32(v46 != int32(0))
	if v46 != 0 {
		v54 = v48
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v54 = v48
	goto L5
L16:
	;
	v50 = v34 + int32(1)
	if v50 != v28 {
		v34 = v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v58 = int32(3)
	goto L20
L19:
	;
	v58 = int32(2)
	goto L20
L20:
	;
	return v58
}
func F_ExecValuesScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(773), int32(774))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_ExitParallelMode(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _consts[25]))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
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
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	v10 = m.G0
	v12 = v10 - int32(2400)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	goto L5
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L16
	} else {
		goto L120
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L16
	} else {
		goto L116
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L16
	} else {
		goto L112
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L16
	} else {
		goto L108
	}
L5:
	;
	if base.B2i32(int32(1) < v18) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L16
	} else {
		goto L104
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v38 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v30 = v29
	goto L12
L11:
	;
	v30 = int32(0)
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(2396)))) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v43 = v39 + int32(1)
	goto L15
L14:
	;
	v43 = int32(1)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+312)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v12)+308)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v12)+304)) = v36
	v53 = F_pg_snprintf(m, v12+int32(1344), int32(1024), int32(488791), v12+int32(304))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	v60 = l0 + int32(24)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v63 = l0 + int32(16)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v66 = int32(2)
	v68 = int32(72)
	v73 = v64<<(uint(v66)%32) + v68
	if int32(0) < v61 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = (v61+v64)<<(uint(v66)%32) + v68
	goto L20
L19:
	;
	v76 = v73
	goto L20
L20:
	;
	v77 = F_MemoryContextAlloc(m, v58, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v80 = v77 + int32(48)
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v81
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+40)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+24)) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+56)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+32)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v77))) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v77)+64)) = int64(0)
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v77)+44)) = v99
	v103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v77)+30)) = uint8(v103)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v105 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v118 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L23:
	;
	v107 = v77 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v112 = v110 << (uint(int32(2)) % 32)
	if v112 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(0)
	goto L22
L26:
	;
	goto L22
L27:
	;
	v113 = F__emscripten_memcpy_bulkmem(m, v107, v109, v112)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v139 = int32(4562080)
	v140 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v143 = *(*int32)(unsafe.Add(mBase, _consts[153]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v143
	v146 = F_palloc(m, int32(8))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L16
	} else {
		goto L41
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = int32(0)
	goto L30
L32:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v121 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v124 != int32(1) {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v127 = v77 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v132 = v130 << (uint(int32(2)) % 32)
	if v132 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L35
L37:
	;
	goto L30
L38:
	;
	v133 = F__emscripten_memcpy_bulkmem(m, v127, v129, v132)
	mBase = m.M
	goto L40
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	v150 = F_pstrdup(m, v12+int32(1344))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v150
	v155 = *(*int32)(unsafe.Add(mBase, _consts[1454]))
	v156 = F_lappend(m, v155, v146)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v140
	*(*int32)(unsafe.Add(mBase, _consts[1454])) = v156
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+48)) = v162 + int32(1)
	F_pairingheap_add(m, int32(4216392), v77+int32(52))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	F_initStringInfo(m, v12+int32(2380))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v176)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+288)) = v177
	F_appendStringInfo(m, v12+int32(2380), int32(781058), v12+int32(288))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _consts[712]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+272)) = v187
	F_appendStringInfo(m, v12+int32(2380), int32(785592), v12+int32(272))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v197
	F_appendStringInfo(m, v12+int32(2380), int32(781041), v12+int32(256))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _consts[134]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v207
	F_appendStringInfo(m, v12+int32(2380), int32(785577), v12+int32(240))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _consts[207])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v217
	F_appendStringInfo(m, v12+int32(2380), int32(785585), v12+int32(224))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v226
	F_appendStringInfo(m, v12+int32(2380), int32(781032), v12+int32(208))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v235
	F_appendStringInfo(m, v12+int32(2380), int32(781007), v12+int32(192))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	if v15 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v244))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v15)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v257 = int32(0)
	goto L55
L55:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = v258 + v257
	F_appendStringInfo(m, v12+int32(2380), int32(785528), v12+int32(176))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L16
	} else {
		goto L60
	}
L56:
	;
	v257 = v256
	goto L55
L57:
	;
	v256 = base.B2i32(base.Ui32(v15) < base.Ui32(v244))
	goto L56
L58:
	;
	goto L59
L59:
	;
	v256 = int32(base.Ui32(v15-v244) >> (uint(int32(31)) % 32))
	goto L56
L60:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v268 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v270 = int32(0)
	goto L64
L62:
	;
	goto L63
L63:
	;
	if v257 != 0 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v270<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v283
	F_appendStringInfo(m, v12+int32(2380), int32(781024), v12+int32(160))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L16
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	v293 = v270 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if base.Ui32(v293) < base.Ui32(v294) {
		v270 = v293
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v15
	F_appendStringInfo(m, v12+int32(2380), int32(781024), v12+int32(144))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L16
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+28)))
	if v313 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L70
L72:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+29)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v422
	F_appendStringInfo(m, v12+int32(2380), int32(781050), v12+int32(80))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L16
	} else {
		goto L94
	}
L73:
	;
	F_appendStringInfoString(m, v12+int32(2380), int32(788860))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L16
	} else {
		goto L80
	}
L74:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	v319 = *(*int32)(unsafe.Add(mBase, _consts[288]))
	v321 = *(*int32)(unsafe.Add(mBase, _consts[780]))
	goto L77
L75:
	;
	goto L76
L76:
	;
	F_appendStringInfoString(m, v12+int32(2380), int32(788853))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L16
	} else {
		goto L79
	}
L77:
	;
	if v316+v32 <= (v319+v321)*int32(65) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L72
L80:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v336 + v32
	F_appendStringInfo(m, v12+int32(2380), int32(785527), v12+int32(128))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	if int32(0) < v346 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v350 = int32(0)
	goto L85
L83:
	;
	goto L84
L84:
	;
	v385 = int32(0)
	if v32 <= v385 {
		goto L72
	} else {
		goto L89
	}
L85:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v350<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v363
	F_appendStringInfo(m, v12+int32(2380), int32(781016), v12+int32(112))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L16
	} else {
		goto L87
	}
L86:
	;
	goto L84
L87:
	;
	v373 = v350 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v77)+24))
	if v373 < v374 {
		v350 = v373
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v388 = v385
	goto L90
L90:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2396))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v397+v388<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v401
	F_appendStringInfo(m, v12+int32(2380), int32(781016), v12+int32(96))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L16
	} else {
		goto L92
	}
L91:
	;
	goto L72
L92:
	;
	v411 = v388 + int32(1)
	if v411 != v32 {
		v388 = v411
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v12 + int32(1344)
	v440 = F_pg_snprintf(m, v12+int32(320), int32(1024), int32(247482), v12-int32(-64))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L16
	} else {
		goto L95
	}
L95:
	;
	v445 = F_AllocateFile(m, v12+int32(320), int32(34045))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L16
	} else {
		goto L96
	}
L96:
	;
	if v445 == int32(0) {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2380))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2384))
	v452 = F_fwrite(m, v449, v450, int32(1), v445)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L16
	} else {
		goto L98
	}
L98:
	;
	if v452 != int32(1) {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v456 = F_FreeFile(m, v445)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L16
	} else {
		goto L100
	}
L100:
	;
	if v456 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	v462 = F_rename(m, v12+int32(320), v12+int32(1344))
	mBase = m.M
	if v462 < int32(0) {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v469 = F_pstrdup(m, v12+int32(1344)|int32(13))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L16
	} else {
		goto L103
	}
L103:
	;
	m.G0 = v12 + int32(2400)
	return v469
L104:
	;
	F_errcode(m, int32(16777538))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L16
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(268810), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(519589), int32(1154), int32(93286))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L16
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 + int32(320)
	F_errmsg(m, int32(313698), v12)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L16
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(519589), int32(1252), int32(93286))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L16
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L16
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(320)
	F_errmsg(m, int32(312702), v12+int32(48))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L16
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(519589), int32(1257), int32(93286))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L16
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
	v532 = m.ExcPending
	if v532 != 0 {
		goto L16
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(320)
	F_errmsg(m, int32(312702), v12+int32(32))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L16
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(519589), int32(1264), int32(93286))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L16
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
	v551 = m.ExcPending
	if v551 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(1344)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(320)
	F_errmsg(m, int32(311730), v12+int32(16))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L16
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(519589), int32(1274), int32(93286))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L16
	} else {
		goto L123
	}
L123:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExtendBufferedRel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v13 = v8 + int32(-8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	if v18 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
		if v22 != 0 {
			v50 = v22
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v50
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+118)))
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v53)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v57
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v59
			v69 = F_ExtendBufferedRelCommon(m, v8+int32(-56), l1, l2, l3, int32(1), int32(-1), v8+int32(-20), v8+int32(-24))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
				m.G0 = v10 - int32(-64)
				return v71
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
			*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v26
			v30 = F_smgropen(m, v8+int32(-40), v23)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v30
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
				if v36 != 0 {
					v44 = v36
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
					v44 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v44 + int32(1)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				v50 = v48
				*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v50
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+118)))
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+56)) = uint8(v53)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v57
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v59
				v69 = F_ExtendBufferedRelCommon(m, v8+int32(-56), l1, l2, l3, int32(1), int32(-1), v8+int32(-20), v8+int32(-24))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
					m.G0 = v10 - int32(-64)
					return v71
				}
			}
		}
	} else {
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v57
		v59 = *(*int64)(unsafe.Add(mBase, uint32(v10)+48))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v59
		v69 = F_ExtendBufferedRelCommon(m, v8+int32(-56), l1, l2, l3, int32(1), int32(-1), v8+int32(-20), v8+int32(-24))
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
			m.G0 = v10 - int32(-64)
			return v71
		}
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	return v53
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v40 = F_equal(m, v38, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v53 = v3
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
		v53 = v3
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 == int32(0) {
		goto L2
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v53 = v3
	goto L1
L16:
	;
	goto L2
L17:
	;
	return int32(0)
L18:
	;
	if v40 == int32(0) {
		v53 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v46 != v47 {
		v53 = v3
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v51 = F_equal(m, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v53 = v51
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+6)) = uint16(v4)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v13 == int32(1) {
		v16 = int32(4562080)
		v17 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, _consts[28])) = v19
		v21 = F_cstring_to_text(m, l1)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			if l2 == int32(0) {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
				if v35 == int32(1) {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v39 = F_cstring_to_text(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v39
						v42 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v42)
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v57 = F_heap_form_tuple(m, v52, v8+int32(8), v8+int32(6))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							F_tuplestore_puttuple(m, v59, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v17
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_MemoryContextReset(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(16)
									return int32(0)
								}
							}
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
					v48 = F_cstring_to_text_with_len(m, v44, v46-v44)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v48
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v57 = F_heap_form_tuple(m, v52, v8+int32(8), v8+int32(6))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							F_tuplestore_puttuple(m, v59, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v17
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_MemoryContextReset(m, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(16)
									return int32(0)
								}
							}
						}
					}
				}
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
				if v28 != int32(1) {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
					if v35 == int32(1) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v39 = F_cstring_to_text(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v39
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v42)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v57 = F_heap_form_tuple(m, v52, v8+int32(8), v8+int32(6))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_tuplestore_puttuple(m, v59, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v17
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_MemoryContextReset(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return int32(0)
									}
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
						v48 = F_cstring_to_text_with_len(m, v44, v46-v44)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v48
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v57 = F_heap_form_tuple(m, v52, v8+int32(8), v8+int32(6))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								F_tuplestore_puttuple(m, v59, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v17
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									F_MemoryContextReset(m, v64)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(16)
										return int32(0)
									}
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
					v33 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)) = uint8(v33)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v57 = F_heap_form_tuple(m, v52, v8+int32(8), v8+int32(6))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						F_tuplestore_puttuple(m, v59, v57)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v17
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_MemoryContextReset(m, v64)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return int32(0)
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(16)
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
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
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
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
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
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	v10 = m.G0
	v12 = v10 - int32(128)
	m.G0 = v12
	v15 = int64(base.Ui64(l0) >> (uint(int64(1)) % 64))
	if base.Ui64(l0) <= base.Ui64(int64(19999999999999)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v362 = int32(2)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(v362)%32))+uint32(_consts[1569])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v365
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v280<<(uint(v362)%32))+uint32(_consts[1569])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v12 - int32(-64)
	F_errmsg(m, int32(760470), v12+int32(48))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L81
	} else {
		goto L92
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l0
	m.G0 = v12 + int32(128)
	return
L3:
	;
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+77)) = uint8(v18)
	v28 = v12 - int32(-64) | int32(13)
	v29 = int32(0)
	v31 = v15
	goto L7
L4:
	;
	goto L5
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v15
	v329 = F_pg_snprintf(m, v12-int32(-64), int32(64), int32(39960), v12+int32(16))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L81
	} else {
		goto L87
	}
L6:
	;
	v62 = int32(3)
	v63 = int32(0)
	goto L19
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
	v53 = int32(12)
	v59 = F__emscripten_memset_bulkmem(m, v34+v29-v53, base.I32_extend8_s(int32(48)), v53-v29)
	mBase = m.M
	goto L13
L13:
	;
	goto L6
L14:
	;
	if l2 == v280 {
		goto L2
	} else {
		goto L80
	}
L15:
	;
	if v129 == int32(0) {
		v280 = v62
		v281 = v63
		goto L14
	} else {
		goto L33
	}
L16:
	;
	v129 = int32(0)
	goto L15
L17:
	;
	v103 = int32(583207)
	v104 = v12 - int32(-64)
	v105 = v62
	goto L27
L19:
	;
	goto L20
L20:
	;
	goto L26
L26:
	;
	goto L17
L27:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 == v109 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v129 = v108 - v109
	goto L15
L29:
	;
	v111 = int32(1)
	v116 = v105 - v111
	if v116 != 0 {
		v103 = v103 + v111
		v104 = v104 + v111
		v105 = v116
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L16
L33:
	;
	goto L38
L34:
	;
	if v197 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L35:
	;
	v197 = int32(0)
	goto L34
L36:
	;
	v171 = int32(583484)
	v172 = v12 - int32(-64)
	v173 = int32(3)
	goto L46
L38:
	;
	goto L39
L39:
	;
	goto L45
L45:
	;
	goto L36
L46:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v176 == v177 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v197 = v176 - v177
	goto L34
L48:
	;
	v179 = int32(1)
	v184 = v173 - v179
	if v184 != 0 {
		v171 = v171 + v179
		v172 = v172 + v179
		v173 = v184
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v280 = int32(5)
	v281 = v63
	goto L14
L53:
	;
	goto L54
L54:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	if v201 == int32(809056057) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v280 = int32(4)
	v281 = v63
	goto L14
L56:
	;
	goto L57
L57:
	;
	goto L62
L58:
	;
	if v270 == int32(0) {
		v280 = v62
		v281 = v63
		goto L14
	} else {
		goto L76
	}
L59:
	;
	v270 = int32(0)
	goto L58
L60:
	;
	v244 = int32(579213)
	v245 = v12 - int32(-64)
	v246 = int32(3)
	goto L70
L62:
	;
	goto L63
L63:
	;
	goto L69
L69:
	;
	goto L60
L70:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v249 == v250 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v270 = v249 - v250
	goto L58
L72:
	;
	v252 = int32(1)
	v257 = v246 - v252
	if v257 != 0 {
		v244 = v244 + v252
		v245 = v245 + v252
		v246 = v257
		goto L70
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L71
L75:
	;
	goto L59
L76:
	;
	v278 = base.B2i32(v201&int32(255) != int32(48))
	if v201&int32(255) != int32(48) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v279 = int32(2)
	goto L79
L78:
	;
	v279 = int32(6)
	goto L79
L79:
	;
	v280 = v279
	v281 = v278
	goto L14
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	return
L82:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if v281 == int32(0) {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(587441)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_consts[1569])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v12 - int32(-64)
	F_errmsg(m, int32(760432), v12+int32(32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(520699), int32(419), int32(257544))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(553676)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v12 - int32(-64)
	F_errmsg(m, int32(386444), v12)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L81
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(520699), int32(437), int32(257544))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L81
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errfinish(m, int32(520699), int32(412), int32(257544))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L81
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ec_add_clause_to_derives_hash(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
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
	var v45 int32
	_ = v45
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
	var v67 int32
	_ = v67
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
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
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
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
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
	var v389 float64
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
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
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1012 int32
	_ = v1012
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int64
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1023 int64
	_ = v1023
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1159 int32
	_ = v1159
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1324 int32
	_ = v1324
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1441 int32
	_ = v1441
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1458 int64
	_ = v1458
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1515 int64
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int64
	_ = v1519
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1550 int64
	_ = v1550
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1622 int32
	_ = v1622
	var v1649 int32
	_ = v1649
	var v1658 int32
	_ = v1658
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+12)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v27
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = int32(0)
	goto L3
L2:
	;
	v30 = v25
	goto L3
L3:
	;
	v31 = base.B2i32(base.Ui32(v24) < base.Ui32(v30))
	if base.Ui32(v24) < base.Ui32(v30) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = v30
	goto L6
L5:
	;
	v32 = v24
	goto L6
L6:
	;
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = v32
	goto L9
L8:
	;
	v33 = v24
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v33
	if base.Ui32(v24) < base.Ui32(v30) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = v24
	goto L12
L11:
	;
	v35 = v30
	goto L12
L12:
	;
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = v35
	goto L15
L14:
	;
	v37 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v37
	v39 = int32(12)
	v45 = int32(-1636608420)
	if v21&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v309 = v304
	v318 = v305
	goto L56
L17:
	;
	v277 = int32(14)
	v279 = v273 ^ v274 - base.I32_rotl(v273, v277)
	v283 = v279 ^ v272 - base.I32_rotl(v279, int32(11))
	v287 = v283 ^ v273 - base.I32_rotl(v283, int32(25))
	v291 = v287 ^ v279 - base.I32_rotl(v287, int32(16))
	v295 = v291 ^ v283 - base.I32_rotl(v291, int32(4))
	v299 = v295 ^ v287 - base.I32_rotl(v295, v277)
	goto L16
L18:
	;
	switch v199 - int32(1) {
	case 0:
		v265 = v190
		v266 = v191
		v267 = v195
		goto L45
	case 1:
		v258 = v190
		v259 = v191
		v260 = v195
		goto L46
	case 2:
		v251 = v190
		v252 = v191
		v253 = v195
		goto L47
	case 3:
		v245 = v191
		v246 = v195
		goto L48
	case 4:
		v241 = v191
		v242 = v195
		goto L49
	case 5:
		v235 = v191
		v236 = v195
		goto L50
	case 6:
		v229 = v191
		v230 = v195
		goto L51
	case 7:
		v224 = v195
		goto L52
	case 8:
		v219 = v195
		goto L53
	case 9:
		v214 = v195
		goto L54
	case 10:
		goto L55
	default:
		v272 = v190
		v273 = v191
		v274 = v195
		goto L17
	}
L19:
	;
	v154 = v21
	v155 = v39
	v156 = v45
	v157 = v45
	v158 = v45
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
	switch v97 - int32(1) {
	case 0:
		v151 = v88
		goto L31
	case 1:
		v146 = v88
		goto L32
	case 2:
		goto L33
	case 3:
		v139 = v89
		goto L34
	case 4:
		v136 = v89
		goto L35
	case 5:
		v131 = v89
		goto L36
	case 6:
		goto L37
	case 7:
		v122 = v93
		goto L38
	case 8:
		v117 = v93
		goto L39
	case 9:
		v112 = v93
		goto L40
	case 10:
		goto L41
	default:
		v272 = v88
		v273 = v89
		v274 = v93
		goto L17
	}
L26:
	;
	goto L27
L27:
	;
	v52 = v21
	v53 = v39
	v54 = v45
	v55 = v45
	v56 = v45
	goto L28
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v59 = v58 + v55
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v63 = v62 + v56
	v65 = int32(4)
	v67 = v60 + v54 - v63 ^ base.I32_rotl(v63, v65)
	v71 = v59 - v67 ^ base.I32_rotl(v67, int32(6))
	v72 = v63 + v59
	v73 = v67 + v72
	v74 = v71 + v73
	v78 = v72 - v71 ^ base.I32_rotl(v71, int32(8))
	v82 = v73 - v78 ^ base.I32_rotl(v78, int32(16))
	v86 = v74 - v82 ^ base.I32_rotl(v82, int32(19))
	v87 = v78 + v74
	v88 = v82 + v87
	v89 = v86 + v88
	v93 = v87 - v86 ^ base.I32_rotl(v86, v65)
	v94 = int32(12)
	v95 = v52 + v94
	v97 = v53 - v94
	if base.Ui32(int32(11)) < base.Ui32(v97) {
		v52 = v95
		v53 = v97
		v54 = v88
		v55 = v89
		v56 = v93
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
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v272 = v151 + v152
	v273 = v89
	v274 = v93
	goto L17
L32:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	v151 = v147<<(uint(int32(8))%32) + v146
	goto L31
L33:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+2)))
	v146 = v142<<(uint(int32(16))%32) + v88
	goto L32
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v272 = v140 + v88
	v273 = v139
	v274 = v93
	goto L17
L35:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)))
	v139 = v136 + v137
	goto L34
L36:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+5)))
	v136 = v132<<(uint(int32(8))%32) + v131
	goto L35
L37:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+6)))
	v131 = v127<<(uint(int32(16))%32) + v89
	goto L36
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v272 = v123 + v88
	v273 = v125 + v89
	v274 = v122
	goto L17
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)))
	v122 = v118<<(uint(int32(8))%32) + v117
	goto L38
L40:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+9)))
	v117 = v113<<(uint(int32(16))%32) + v112
	goto L39
L41:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+10)))
	v112 = v108<<(uint(int32(24))%32) + v93
	goto L40
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v161 = v160 + v157
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v154)+8))
	v165 = v164 + v158
	v167 = int32(4)
	v169 = v162 + v156 - v165 ^ base.I32_rotl(v165, v167)
	v173 = v161 - v169 ^ base.I32_rotl(v169, int32(6))
	v174 = v165 + v161
	v175 = v169 + v174
	v176 = v173 + v175
	v180 = v174 - v173 ^ base.I32_rotl(v173, int32(8))
	v184 = v175 - v180 ^ base.I32_rotl(v180, int32(16))
	v188 = v176 - v184 ^ base.I32_rotl(v184, int32(19))
	v189 = v180 + v176
	v190 = v184 + v189
	v191 = v188 + v190
	v195 = v189 - v188 ^ base.I32_rotl(v188, v167)
	v196 = int32(12)
	v197 = v154 + v196
	v199 = v155 - v196
	if base.Ui32(int32(11)) < base.Ui32(v199) {
		v154 = v197
		v155 = v199
		v156 = v190
		v157 = v191
		v158 = v195
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
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v272 = v265 + v268
	v273 = v266
	v274 = v267
	goto L17
L46:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	v265 = v261<<(uint(int32(8))%32) + v258
	v266 = v259
	v267 = v260
	goto L45
L47:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+2)))
	v258 = v254<<(uint(int32(16))%32) + v251
	v259 = v252
	v260 = v253
	goto L46
L48:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+3)))
	v251 = v247<<(uint(int32(24))%32) + v190
	v252 = v245
	v253 = v246
	goto L47
L49:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+4)))
	v245 = v241 + v243
	v246 = v242
	goto L48
L50:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+5)))
	v241 = v237<<(uint(int32(8))%32) + v235
	v242 = v236
	goto L49
L51:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+6)))
	v235 = v231<<(uint(int32(16))%32) + v229
	v236 = v230
	goto L50
L52:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+7)))
	v229 = v225<<(uint(int32(24))%32) + v191
	v230 = v224
	goto L51
L53:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+8)))
	v224 = v220<<(uint(int32(8))%32) + v219
	goto L52
L54:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+9)))
	v219 = v215<<(uint(int32(16))%32) + v214
	goto L53
L55:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+10)))
	v214 = v210<<(uint(int32(24))%32) + v195
	goto L54
L56:
	;
	if base.Ui32(v309) <= base.Ui32(v318) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v1658 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1658
	v309 = v1658
	v318 = v1649
	goto L56
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+16)) = l1
	m.G0 = v21 + int32(16)
	return
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1599)+12)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v1599)+8)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v1599)+4)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v1599))) = int32(1)
	v1622 = v1599
	goto L59
L61:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1591 + int32(1)
	v1599 = v1577
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L78
	} else {
		goto L273
	}
L63:
	;
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	if v328 == int64(4294967296) {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v1117 = int32(0)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1120 = v1119 & (v299 ^ v291 - base.I32_rotl(v299, int32(24)))
	v1123 = v1118 + v1120*int32(20)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)))
	if v1124 == v1117 {
		v1577 = v1123
		goto L61
	} else {
		goto L199
	}
L66:
	;
	v331 = int32(0)
	v333 = int64(2)
	v335 = v328 << (uint(int64(1)) % 64)
	if base.Ui64(v335) <= base.Ui64(v333) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L65
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L78
	} else {
		goto L196
	}
L69:
	;
	v338 = v333
	goto L71
L70:
	;
	v338 = v335
	goto L71
L71:
	;
	v339 = int64(1)
	if v338&(v338-v339) == int64(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v349 = v338
	goto L74
L73:
	;
	v349 = v339 << (uint(int64(64)-base.I64_clz(v338)) % 64)
	goto L74
L74:
	;
	if base.Ui64(v349*int64(20)) < base.Ui64(int64(2147483647)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v361 = F_MemoryContextAllocExtended(m, v356, base.I32_wrap_i64(v349)*int32(20), int32(5))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
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
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L78
	} else {
		goto L193
	}
L78:
	;
	return
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v361
	v364 = int64(1)
	if v349&(v349-v364) == int64(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v374 = v349
	goto L82
L81:
	;
	v374 = v364 << (uint(int64(64)-base.I64_clz(v349)) % 64)
	goto L82
L82:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v374*int64(20)) {
		goto L68
	} else {
		goto L83
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = base.I32_wrap_i64(v374) - int32(1)
	v389 = base.F64_mul(base.F64_convert_i64_u(v374), float64(0.9))
	if base.F64_lt(v389, float64(4.294967296e+09))&base.F64_ge(v389, float64(0)) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v374 == int64(4294967296) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v395 = base.I32_trunc_f64_u(v389)
	v397 = v395
	goto L84
L86:
	;
	goto L87
L87:
	;
	v397 = int32(0)
	goto L84
L88:
	;
	v398 = int32(-85899346)
	goto L90
L89:
	;
	v398 = v397
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v398
	if v355 != int64(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v405 = v331
	goto L95
L92:
	;
	goto L93
L93:
	;
	F_pfree(m, v354)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L78
	} else {
		goto L192
	}
L94:
	;
	v705 = v701
	v706 = v331
	goto L140
L95:
	;
	v422 = v354 + v405*int32(20)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v423 != int32(1) {
		v701 = v405
		goto L94
	} else {
		goto L97
	}
L96:
	;
	v701 = int32(0)
	goto L94
L97:
	;
	v427 = v422 + int32(4)
	v428 = int32(12)
	v434 = int32(-1636608420)
	if v427&int32(3) != 0 {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if (v688^v680-base.I32_rotl(v688, int32(24)))&v693 == v405 {
		v701 = v405
		goto L94
	} else {
		goto L138
	}
L99:
	;
	v666 = int32(14)
	v668 = v662 ^ v663 - base.I32_rotl(v662, v666)
	v672 = v668 ^ v661 - base.I32_rotl(v668, int32(11))
	v676 = v672 ^ v662 - base.I32_rotl(v672, int32(25))
	v680 = v676 ^ v668 - base.I32_rotl(v676, int32(16))
	v684 = v680 ^ v672 - base.I32_rotl(v680, int32(4))
	v688 = v684 ^ v676 - base.I32_rotl(v684, v666)
	goto L98
L100:
	;
	switch v588 - int32(1) {
	case 0:
		v654 = v579
		v655 = v580
		v656 = v584
		goto L127
	case 1:
		v647 = v579
		v648 = v580
		v649 = v584
		goto L128
	case 2:
		v640 = v579
		v641 = v580
		v642 = v584
		goto L129
	case 3:
		v634 = v580
		v635 = v584
		goto L130
	case 4:
		v630 = v580
		v631 = v584
		goto L131
	case 5:
		v624 = v580
		v625 = v584
		goto L132
	case 6:
		v618 = v580
		v619 = v584
		goto L133
	case 7:
		v613 = v584
		goto L134
	case 8:
		v608 = v584
		goto L135
	case 9:
		v603 = v584
		goto L136
	case 10:
		goto L137
	default:
		v661 = v579
		v662 = v580
		v663 = v584
		goto L99
	}
L101:
	;
	v543 = v427
	v544 = v428
	v545 = v434
	v546 = v434
	v547 = v434
	goto L124
L102:
	;
	goto L101
L103:
	;
	goto L104
L104:
	;
	goto L108
L106:
	;
	switch v486 - int32(1) {
	case 0:
		v540 = v477
		goto L113
	case 1:
		v535 = v477
		goto L114
	case 2:
		goto L115
	case 3:
		v528 = v478
		goto L116
	case 4:
		v525 = v478
		goto L117
	case 5:
		v520 = v478
		goto L118
	case 6:
		goto L119
	case 7:
		v511 = v482
		goto L120
	case 8:
		v506 = v482
		goto L121
	case 9:
		v501 = v482
		goto L122
	case 10:
		goto L123
	default:
		v661 = v477
		v662 = v478
		v663 = v482
		goto L99
	}
L108:
	;
	goto L109
L109:
	;
	v441 = v427
	v442 = v428
	v443 = v434
	v444 = v434
	v445 = v434
	goto L110
L110:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v448 = v447 + v444
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	v452 = v451 + v445
	v454 = int32(4)
	v456 = v449 + v443 - v452 ^ base.I32_rotl(v452, v454)
	v460 = v448 - v456 ^ base.I32_rotl(v456, int32(6))
	v461 = v452 + v448
	v462 = v456 + v461
	v463 = v460 + v462
	v467 = v461 - v460 ^ base.I32_rotl(v460, int32(8))
	v471 = v462 - v467 ^ base.I32_rotl(v467, int32(16))
	v475 = v463 - v471 ^ base.I32_rotl(v471, int32(19))
	v476 = v467 + v463
	v477 = v471 + v476
	v478 = v475 + v477
	v482 = v476 - v475 ^ base.I32_rotl(v475, v454)
	v483 = int32(12)
	v484 = v441 + v483
	v486 = v442 - v483
	if base.Ui32(int32(11)) < base.Ui32(v486) {
		v441 = v484
		v442 = v486
		v443 = v477
		v444 = v478
		v445 = v482
		goto L110
	} else {
		goto L112
	}
L111:
	;
	goto L106
L112:
	;
	goto L111
L113:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484))))
	v661 = v540 + v541
	v662 = v478
	v663 = v482
	goto L99
L114:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+1)))
	v540 = v536<<(uint(int32(8))%32) + v535
	goto L113
L115:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+2)))
	v535 = v531<<(uint(int32(16))%32) + v477
	goto L114
L116:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	v661 = v529 + v477
	v662 = v528
	v663 = v482
	goto L99
L117:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+4)))
	v528 = v525 + v526
	goto L116
L118:
	;
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+5)))
	v525 = v521<<(uint(int32(8))%32) + v520
	goto L117
L119:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+6)))
	v520 = v516<<(uint(int32(16))%32) + v478
	goto L118
L120:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	v661 = v512 + v477
	v662 = v514 + v478
	v663 = v511
	goto L99
L121:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+8)))
	v511 = v507<<(uint(int32(8))%32) + v506
	goto L120
L122:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+9)))
	v506 = v502<<(uint(int32(16))%32) + v501
	goto L121
L123:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484)+10)))
	v501 = v497<<(uint(int32(24))%32) + v482
	goto L122
L124:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v550 = v549 + v546
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v543)+8))
	v554 = v553 + v547
	v556 = int32(4)
	v558 = v551 + v545 - v554 ^ base.I32_rotl(v554, v556)
	v562 = v550 - v558 ^ base.I32_rotl(v558, int32(6))
	v563 = v554 + v550
	v564 = v558 + v563
	v565 = v562 + v564
	v569 = v563 - v562 ^ base.I32_rotl(v562, int32(8))
	v573 = v564 - v569 ^ base.I32_rotl(v569, int32(16))
	v577 = v565 - v573 ^ base.I32_rotl(v573, int32(19))
	v578 = v569 + v565
	v579 = v573 + v578
	v580 = v577 + v579
	v584 = v578 - v577 ^ base.I32_rotl(v577, v556)
	v585 = int32(12)
	v586 = v543 + v585
	v588 = v544 - v585
	if base.Ui32(int32(11)) < base.Ui32(v588) {
		v543 = v586
		v544 = v588
		v545 = v579
		v546 = v580
		v547 = v584
		goto L124
	} else {
		goto L126
	}
L125:
	;
	goto L100
L126:
	;
	goto L125
L127:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	v661 = v654 + v657
	v662 = v655
	v663 = v656
	goto L99
L128:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+1)))
	v654 = v650<<(uint(int32(8))%32) + v647
	v655 = v648
	v656 = v649
	goto L127
L129:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+2)))
	v647 = v643<<(uint(int32(16))%32) + v640
	v648 = v641
	v649 = v642
	goto L128
L130:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+3)))
	v640 = v636<<(uint(int32(24))%32) + v579
	v641 = v634
	v642 = v635
	goto L129
L131:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+4)))
	v634 = v630 + v632
	v635 = v631
	goto L130
L132:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+5)))
	v630 = v626<<(uint(int32(8))%32) + v624
	v631 = v625
	goto L131
L133:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+6)))
	v624 = v620<<(uint(int32(16))%32) + v618
	v625 = v619
	goto L132
L134:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+7)))
	v618 = v614<<(uint(int32(24))%32) + v580
	v619 = v613
	goto L133
L135:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+8)))
	v613 = v609<<(uint(int32(8))%32) + v608
	goto L134
L136:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+9)))
	v608 = v604<<(uint(int32(16))%32) + v603
	goto L135
L137:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+10)))
	v603 = v599<<(uint(int32(24))%32) + v584
	goto L136
L138:
	;
	v697 = v405 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v697)) < base.Ui64(v355) {
		v405 = v697
		goto L95
	} else {
		goto L139
	}
L139:
	;
	goto L96
L140:
	;
	v722 = v354 + v705*int32(20)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	if v723 == int32(1) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L93
L142:
	;
	v727 = v722 + int32(4)
	v728 = int32(12)
	v734 = int32(-1636608420)
	if v727&int32(3) != 0 {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	goto L144
L144:
	;
	v1044 = v705 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1044)) < base.Ui64(v355) {
		goto L188
	} else {
		goto L189
	}
L145:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v994 = v988 ^ v980 - base.I32_rotl(v988, int32(24))
	goto L185
L146:
	;
	v966 = int32(14)
	v968 = v962 ^ v963 - base.I32_rotl(v962, v966)
	v972 = v968 ^ v961 - base.I32_rotl(v968, int32(11))
	v976 = v972 ^ v962 - base.I32_rotl(v972, int32(25))
	v980 = v976 ^ v968 - base.I32_rotl(v976, int32(16))
	v984 = v980 ^ v972 - base.I32_rotl(v980, int32(4))
	v988 = v984 ^ v976 - base.I32_rotl(v984, v966)
	goto L145
L147:
	;
	switch v888 - int32(1) {
	case 0:
		v954 = v879
		v955 = v880
		v956 = v884
		goto L174
	case 1:
		v947 = v879
		v948 = v880
		v949 = v884
		goto L175
	case 2:
		v940 = v879
		v941 = v880
		v942 = v884
		goto L176
	case 3:
		v934 = v880
		v935 = v884
		goto L177
	case 4:
		v930 = v880
		v931 = v884
		goto L178
	case 5:
		v924 = v880
		v925 = v884
		goto L179
	case 6:
		v918 = v880
		v919 = v884
		goto L180
	case 7:
		v913 = v884
		goto L181
	case 8:
		v908 = v884
		goto L182
	case 9:
		v903 = v884
		goto L183
	case 10:
		goto L184
	default:
		v961 = v879
		v962 = v880
		v963 = v884
		goto L146
	}
L148:
	;
	v843 = v727
	v844 = v728
	v845 = v734
	v846 = v734
	v847 = v734
	goto L171
L149:
	;
	goto L148
L150:
	;
	goto L151
L151:
	;
	goto L155
L153:
	;
	switch v786 - int32(1) {
	case 0:
		v840 = v777
		goto L160
	case 1:
		v835 = v777
		goto L161
	case 2:
		goto L162
	case 3:
		v828 = v778
		goto L163
	case 4:
		v825 = v778
		goto L164
	case 5:
		v820 = v778
		goto L165
	case 6:
		goto L166
	case 7:
		v811 = v782
		goto L167
	case 8:
		v806 = v782
		goto L168
	case 9:
		v801 = v782
		goto L169
	case 10:
		goto L170
	default:
		v961 = v777
		v962 = v778
		v963 = v782
		goto L146
	}
L155:
	;
	goto L156
L156:
	;
	v741 = v727
	v742 = v728
	v743 = v734
	v744 = v734
	v745 = v734
	goto L157
L157:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	v748 = v747 + v744
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v741)))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v741)+8))
	v752 = v751 + v745
	v754 = int32(4)
	v756 = v749 + v743 - v752 ^ base.I32_rotl(v752, v754)
	v760 = v748 - v756 ^ base.I32_rotl(v756, int32(6))
	v761 = v752 + v748
	v762 = v756 + v761
	v763 = v760 + v762
	v767 = v761 - v760 ^ base.I32_rotl(v760, int32(8))
	v771 = v762 - v767 ^ base.I32_rotl(v767, int32(16))
	v775 = v763 - v771 ^ base.I32_rotl(v771, int32(19))
	v776 = v767 + v763
	v777 = v771 + v776
	v778 = v775 + v777
	v782 = v776 - v775 ^ base.I32_rotl(v775, v754)
	v783 = int32(12)
	v784 = v741 + v783
	v786 = v742 - v783
	if base.Ui32(int32(11)) < base.Ui32(v786) {
		v741 = v784
		v742 = v786
		v743 = v777
		v744 = v778
		v745 = v782
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L153
L159:
	;
	goto L158
L160:
	;
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784))))
	v961 = v840 + v841
	v962 = v778
	v963 = v782
	goto L146
L161:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+1)))
	v840 = v836<<(uint(int32(8))%32) + v835
	goto L160
L162:
	;
	v831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+2)))
	v835 = v831<<(uint(int32(16))%32) + v777
	goto L161
L163:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v961 = v829 + v777
	v962 = v828
	v963 = v782
	goto L146
L164:
	;
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+4)))
	v828 = v825 + v826
	goto L163
L165:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+5)))
	v825 = v821<<(uint(int32(8))%32) + v820
	goto L164
L166:
	;
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+6)))
	v820 = v816<<(uint(int32(16))%32) + v778
	goto L165
L167:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v784)+4))
	v961 = v812 + v777
	v962 = v814 + v778
	v963 = v811
	goto L146
L168:
	;
	v807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+8)))
	v811 = v807<<(uint(int32(8))%32) + v806
	goto L167
L169:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+9)))
	v806 = v802<<(uint(int32(16))%32) + v801
	goto L168
L170:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+10)))
	v801 = v797<<(uint(int32(24))%32) + v782
	goto L169
L171:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
	v850 = v849 + v846
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v843)+8))
	v854 = v853 + v847
	v856 = int32(4)
	v858 = v851 + v845 - v854 ^ base.I32_rotl(v854, v856)
	v862 = v850 - v858 ^ base.I32_rotl(v858, int32(6))
	v863 = v854 + v850
	v864 = v858 + v863
	v865 = v862 + v864
	v869 = v863 - v862 ^ base.I32_rotl(v862, int32(8))
	v873 = v864 - v869 ^ base.I32_rotl(v869, int32(16))
	v877 = v865 - v873 ^ base.I32_rotl(v873, int32(19))
	v878 = v869 + v865
	v879 = v873 + v878
	v880 = v877 + v879
	v884 = v878 - v877 ^ base.I32_rotl(v877, v856)
	v885 = int32(12)
	v886 = v843 + v885
	v888 = v844 - v885
	if base.Ui32(int32(11)) < base.Ui32(v888) {
		v843 = v886
		v844 = v888
		v845 = v879
		v846 = v880
		v847 = v884
		goto L171
	} else {
		goto L173
	}
L172:
	;
	goto L147
L173:
	;
	goto L172
L174:
	;
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886))))
	v961 = v954 + v957
	v962 = v955
	v963 = v956
	goto L146
L175:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+1)))
	v954 = v950<<(uint(int32(8))%32) + v947
	v955 = v948
	v956 = v949
	goto L174
L176:
	;
	v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+2)))
	v947 = v943<<(uint(int32(16))%32) + v940
	v948 = v941
	v949 = v942
	goto L175
L177:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+3)))
	v940 = v936<<(uint(int32(24))%32) + v879
	v941 = v934
	v942 = v935
	goto L176
L178:
	;
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+4)))
	v934 = v930 + v932
	v935 = v931
	goto L177
L179:
	;
	v926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+5)))
	v930 = v926<<(uint(int32(8))%32) + v924
	v931 = v925
	goto L178
L180:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+6)))
	v924 = v920<<(uint(int32(16))%32) + v918
	v925 = v919
	goto L179
L181:
	;
	v914 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+7)))
	v918 = v914<<(uint(int32(24))%32) + v880
	v919 = v913
	goto L180
L182:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+8)))
	v913 = v909<<(uint(int32(8))%32) + v908
	goto L181
L183:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+9)))
	v908 = v904<<(uint(int32(16))%32) + v903
	goto L182
L184:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+10)))
	v903 = v899<<(uint(int32(24))%32) + v884
	goto L183
L185:
	;
	v1012 = v994 & v993
	v1017 = v361 + v1012*int32(20)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)))
	if v1018 != 0 {
		v994 = v1012 + int32(1)
		goto L185
	} else {
		goto L187
	}
L186:
	;
	v1019 = *(*int64)(unsafe.Add(mBase, uint32(v722)))
	*(*int64)(unsafe.Add(mBase, uint32(v1017))) = v1019
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v722)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1017)+16)) = v1021
	v1023 = *(*int64)(unsafe.Add(mBase, uint32(v722)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1017)+8)) = v1023
	goto L144
L187:
	;
	goto L186
L188:
	;
	v1048 = v1044
	goto L190
L189:
	;
	v1048 = int32(0)
	goto L190
L190:
	;
	v1050 = v706 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v1050)) < base.Ui64(v355) {
		v705 = v1048
		v706 = v1050
		goto L140
	} else {
		goto L191
	}
L191:
	;
	goto L141
L192:
	;
	goto L67
L193:
	;
	F_errmsg_internal(m, int32(419840), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L78
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(343467), int32(327), int32(358444))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L78
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errmsg_internal(m, int32(419840), int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L78
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(343467), int32(327), int32(358444))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L78
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	v1130 = v1120
	v1131 = v1123
	v1132 = v1117
	goto L200
L200:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+4))
	if v1145 != v308 {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1577 = v1558
	goto L61
L202:
	;
	v1152 = v1131 + int32(4)
	v1153 = int32(12)
	v1159 = int32(-1636608420)
	if v1152&int32(3) != 0 {
		goto L210
	} else {
		goto L211
	}
L203:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+8))
	if v1147 != v307 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1131)+12))
	if v1149 == v306 {
		v1622 = v1131
		goto L59
	} else {
		goto L205
	}
L205:
	;
	goto L202
L206:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1419 = (v1413 ^ v1405 - base.I32_rotl(v1413, int32(24))) & v1418
	if base.Ui32(v1130) < base.Ui32(v1419) {
		goto L246
	} else {
		goto L247
	}
L207:
	;
	v1391 = int32(14)
	v1393 = v1387 ^ v1388 - base.I32_rotl(v1387, v1391)
	v1397 = v1393 ^ v1386 - base.I32_rotl(v1393, int32(11))
	v1401 = v1397 ^ v1387 - base.I32_rotl(v1397, int32(25))
	v1405 = v1401 ^ v1393 - base.I32_rotl(v1401, int32(16))
	v1409 = v1405 ^ v1397 - base.I32_rotl(v1405, int32(4))
	v1413 = v1409 ^ v1401 - base.I32_rotl(v1409, v1391)
	goto L206
L208:
	;
	switch v1313 - int32(1) {
	case 0:
		v1379 = v1304
		v1380 = v1305
		v1381 = v1309
		goto L235
	case 1:
		v1372 = v1304
		v1373 = v1305
		v1374 = v1309
		goto L236
	case 2:
		v1365 = v1304
		v1366 = v1305
		v1367 = v1309
		goto L237
	case 3:
		v1359 = v1305
		v1360 = v1309
		goto L238
	case 4:
		v1355 = v1305
		v1356 = v1309
		goto L239
	case 5:
		v1349 = v1305
		v1350 = v1309
		goto L240
	case 6:
		v1343 = v1305
		v1344 = v1309
		goto L241
	case 7:
		v1338 = v1309
		goto L242
	case 8:
		v1333 = v1309
		goto L243
	case 9:
		v1328 = v1309
		goto L244
	case 10:
		goto L245
	default:
		v1386 = v1304
		v1387 = v1305
		v1388 = v1309
		goto L207
	}
L209:
	;
	v1268 = v1152
	v1269 = v1153
	v1270 = v1159
	v1271 = v1159
	v1272 = v1159
	goto L232
L210:
	;
	goto L209
L211:
	;
	goto L212
L212:
	;
	goto L216
L214:
	;
	switch v1211 - int32(1) {
	case 0:
		v1265 = v1202
		goto L221
	case 1:
		v1260 = v1202
		goto L222
	case 2:
		goto L223
	case 3:
		v1253 = v1203
		goto L224
	case 4:
		v1250 = v1203
		goto L225
	case 5:
		v1245 = v1203
		goto L226
	case 6:
		goto L227
	case 7:
		v1236 = v1207
		goto L228
	case 8:
		v1231 = v1207
		goto L229
	case 9:
		v1226 = v1207
		goto L230
	case 10:
		goto L231
	default:
		v1386 = v1202
		v1387 = v1203
		v1388 = v1207
		goto L207
	}
L216:
	;
	goto L217
L217:
	;
	v1166 = v1152
	v1167 = v1153
	v1168 = v1159
	v1169 = v1159
	v1170 = v1159
	goto L218
L218:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+4))
	v1173 = v1172 + v1169
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1166)))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+8))
	v1177 = v1176 + v1170
	v1179 = int32(4)
	v1181 = v1174 + v1168 - v1177 ^ base.I32_rotl(v1177, v1179)
	v1185 = v1173 - v1181 ^ base.I32_rotl(v1181, int32(6))
	v1186 = v1177 + v1173
	v1187 = v1181 + v1186
	v1188 = v1185 + v1187
	v1192 = v1186 - v1185 ^ base.I32_rotl(v1185, int32(8))
	v1196 = v1187 - v1192 ^ base.I32_rotl(v1192, int32(16))
	v1200 = v1188 - v1196 ^ base.I32_rotl(v1196, int32(19))
	v1201 = v1192 + v1188
	v1202 = v1196 + v1201
	v1203 = v1200 + v1202
	v1207 = v1201 - v1200 ^ base.I32_rotl(v1200, v1179)
	v1208 = int32(12)
	v1209 = v1166 + v1208
	v1211 = v1167 - v1208
	if base.Ui32(int32(11)) < base.Ui32(v1211) {
		v1166 = v1209
		v1167 = v1211
		v1168 = v1202
		v1169 = v1203
		v1170 = v1207
		goto L218
	} else {
		goto L220
	}
L219:
	;
	goto L214
L220:
	;
	goto L219
L221:
	;
	v1266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209))))
	v1386 = v1265 + v1266
	v1387 = v1203
	v1388 = v1207
	goto L207
L222:
	;
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+1)))
	v1265 = v1261<<(uint(int32(8))%32) + v1260
	goto L221
L223:
	;
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+2)))
	v1260 = v1256<<(uint(int32(16))%32) + v1202
	goto L222
L224:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	v1386 = v1254 + v1202
	v1387 = v1253
	v1388 = v1207
	goto L207
L225:
	;
	v1251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+4)))
	v1253 = v1250 + v1251
	goto L224
L226:
	;
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+5)))
	v1250 = v1246<<(uint(int32(8))%32) + v1245
	goto L225
L227:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+6)))
	v1245 = v1241<<(uint(int32(16))%32) + v1203
	goto L226
L228:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1209)))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+4))
	v1386 = v1237 + v1202
	v1387 = v1239 + v1203
	v1388 = v1236
	goto L207
L229:
	;
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+8)))
	v1236 = v1232<<(uint(int32(8))%32) + v1231
	goto L228
L230:
	;
	v1227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+9)))
	v1231 = v1227<<(uint(int32(16))%32) + v1226
	goto L229
L231:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1209)+10)))
	v1226 = v1222<<(uint(int32(24))%32) + v1207
	goto L230
L232:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+4))
	v1275 = v1274 + v1271
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1268)))
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v1268)+8))
	v1279 = v1278 + v1272
	v1281 = int32(4)
	v1283 = v1276 + v1270 - v1279 ^ base.I32_rotl(v1279, v1281)
	v1287 = v1275 - v1283 ^ base.I32_rotl(v1283, int32(6))
	v1288 = v1279 + v1275
	v1289 = v1283 + v1288
	v1290 = v1287 + v1289
	v1294 = v1288 - v1287 ^ base.I32_rotl(v1287, int32(8))
	v1298 = v1289 - v1294 ^ base.I32_rotl(v1294, int32(16))
	v1302 = v1290 - v1298 ^ base.I32_rotl(v1298, int32(19))
	v1303 = v1294 + v1290
	v1304 = v1298 + v1303
	v1305 = v1302 + v1304
	v1309 = v1303 - v1302 ^ base.I32_rotl(v1302, v1281)
	v1310 = int32(12)
	v1311 = v1268 + v1310
	v1313 = v1269 - v1310
	if base.Ui32(int32(11)) < base.Ui32(v1313) {
		v1268 = v1311
		v1269 = v1313
		v1270 = v1304
		v1271 = v1305
		v1272 = v1309
		goto L232
	} else {
		goto L234
	}
L233:
	;
	goto L208
L234:
	;
	goto L233
L235:
	;
	v1382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311))))
	v1386 = v1379 + v1382
	v1387 = v1380
	v1388 = v1381
	goto L207
L236:
	;
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+1)))
	v1379 = v1375<<(uint(int32(8))%32) + v1372
	v1380 = v1373
	v1381 = v1374
	goto L235
L237:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+2)))
	v1372 = v1368<<(uint(int32(16))%32) + v1365
	v1373 = v1366
	v1374 = v1367
	goto L236
L238:
	;
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+3)))
	v1365 = v1361<<(uint(int32(24))%32) + v1304
	v1366 = v1359
	v1367 = v1360
	goto L237
L239:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+4)))
	v1359 = v1355 + v1357
	v1360 = v1356
	goto L238
L240:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+5)))
	v1355 = v1351<<(uint(int32(8))%32) + v1349
	v1356 = v1350
	goto L239
L241:
	;
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+6)))
	v1349 = v1345<<(uint(int32(16))%32) + v1343
	v1350 = v1344
	goto L240
L242:
	;
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+7)))
	v1343 = v1339<<(uint(int32(24))%32) + v1305
	v1344 = v1338
	goto L241
L243:
	;
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+8)))
	v1338 = v1334<<(uint(int32(8))%32) + v1333
	goto L242
L244:
	;
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+9)))
	v1333 = v1329<<(uint(int32(16))%32) + v1328
	goto L243
L245:
	;
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1311)+10)))
	v1328 = v1324<<(uint(int32(24))%32) + v1309
	goto L244
L246:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v1423 = v1130 + v1421
	goto L248
L247:
	;
	v1423 = v1130
	goto L248
L248:
	;
	v1426 = v1418 & (v1130 + int32(1))
	if base.Ui32(v1423-v1419) < base.Ui32(v1132) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1432 = v1118 + v1426*int32(20)
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1432)))
	if v1433 != 0 {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	v1545 = v1132 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v1545) {
		goto L268
	} else {
		goto L269
	}
L252:
	;
	v1434 = v1426
	v1441 = int32(0)
	goto L255
L253:
	;
	v1471 = v1426
	v1476 = v1432
	goto L254
L254:
	;
	if v1471 != v1130 {
		goto L262
	} else {
		goto L263
	}
L255:
	;
	v1453 = v1441 + int32(1)
	if int32(151) <= v1453 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v1471 = v1466
	v1476 = v1469
	goto L254
L257:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1458 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1456), base.F64_convert_i64_u(v1458)), float64(0.1)) != 0 {
		v1649 = v1456
		goto L58
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	v1466 = (v1434 + int32(1)) & v1418
	v1469 = v1118 + v1466*int32(20)
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1469)))
	if v1470 != 0 {
		v1434 = v1466
		v1441 = v1453
		goto L255
	} else {
		goto L261
	}
L260:
	;
	goto L259
L261:
	;
	goto L256
L262:
	;
	v1490 = v1471
	v1495 = v1476
	goto L265
L263:
	;
	goto L264
L264:
	;
	v1540 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1540 + int32(1)
	v1599 = v1131
	goto L60
L265:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1511 = v1508 & (v1490 - int32(1))
	v1514 = v1118 + v1511*int32(20)
	v1515 = *(*int64)(unsafe.Add(mBase, uint32(v1514)))
	*(*int64)(unsafe.Add(mBase, uint32(v1495))) = v1515
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1514)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1495)+16)) = v1517
	v1519 = *(*int64)(unsafe.Add(mBase, uint32(v1514)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1495)+8)) = v1519
	if v1511 != v1130 {
		v1490 = v1511
		v1495 = v1514
		goto L265
	} else {
		goto L267
	}
L266:
	;
	goto L264
L267:
	;
	goto L266
L268:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1550 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v1548), base.F64_convert_i64_u(v1550)), float64(0.1)) != 0 {
		v1649 = v1548
		goto L58
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1558 = v1118 + v1426*int32(20)
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1558)))
	if v1559 != 0 {
		v1130 = v1426
		v1131 = v1558
		v1132 = v1545
		goto L200
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	goto L201
L273:
	;
	F_errmsg_internal(m, int32(484212), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L78
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(343467), int32(630), int32(326930))
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L78
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_elem_contained_by_multirange(m *base.Module, l0 int32) int32 {
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
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
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v32
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(65536))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(388681), v9)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518389), int32(558), int32(418059))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
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
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v32
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(65536))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+296))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(388681), v9)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518389), int32(558), int32(418059))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
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
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v32
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = F_FunctionCall2Coll(m, l2+int32(104), v6, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
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
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v20 = int32(1)
			v21 = v18 + v20
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v26 = v24 & v20
			if v26 != 0 {
				v27 = v21
			} else {
				v27 = v18 + int32(4)
			}
			if v24 == int32(1) {
				v30 = int32(4)
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				if v32&int32(254) == int32(2) {
					v41 = v30
				} else {
					v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
				}
				if v32 == int32(1) {
					v44 = v30
				} else {
					v44 = v41
				}
				v55 = v44
			} else {
				v45 = int32(1)
				if v26 != 0 {
					v55 = int32(base.Ui32(v24)>>(uint(v45)%32)) - v45
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v57 = *(*int32)(unsafe.Add(mBase, _consts[495]))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
			v59 = F_makeJsonLexContextCstringLen(m, v11+int32(12), v27, v55, v58, v3)
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
							*(*int32)(unsafe.Add(mBase, uint32(v65)+36)) = int32(1370)
							*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(1371)
							*(*int32)(unsafe.Add(mBase, uint32(v65))) = v62
							*(*int32)(unsafe.Add(mBase, uint32(v65)+32)) = int32(1372)
							*(*int32)(unsafe.Add(mBase, uint32(v65)+28)) = int32(1373)
							v84 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+25)) = uint8(v84)
							*(*uint8)(unsafe.Add(mBase, uint32(v62)+24)) = uint8(v3)
							*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v11 + int32(12)
							v92 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							v97 = F_AllocSetContextCreateInternal(m, v92, int32(69967), v84, int32(8192), int32(8388608))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v97
								v102 = F_pg_parse_json(m, v11+int32(12), v65)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return
								} else {
									if v102 != 0 {
										F_json_errsave_error(m, v102, v11+int32(12), int32(0))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return
										} else {
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
											F_MemoryContextDelete(m, v109)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return
											} else {
												F_freeJsonLexContext(m, v11+int32(12))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													v116 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v116)
													m.G0 = v11 + int32(80)
													return
												}
											}
										}
									} else {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
										F_MemoryContextDelete(m, v109)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											F_freeJsonLexContext(m, v11+int32(12))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												v116 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v116)
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v45 int32
	_ = v45
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
	v34 = v28
	v36 = l1
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	if v37 != int32(110) {
		v45 = v36
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return v45
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	if v46 != 0 {
		v34 = v46
		v36 = v45
		goto L12
	} else {
		goto L18
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v41 != 0 {
		v45 = v36
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v42 = F_emptyreachable(m, l0, v40, v36, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v45 = v42
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
	var v39 int32
	_ = v39
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
	*(*int32)(unsafe.Add(mBase, _consts[646])) = v3
	v18 = m.G0
	v20 = v18 - v9
	m.G0 = v20
	F___gettimeofday(m, v20)
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
	v39 = v3
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
	v44 = l0 + v39*int32(24)
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
	v81 = v39 + int32(1)
	if v81 != l1 {
		v39 = v81
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
	F_errmsg_internal(m, int32(498269), v10)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(516127), int32(666), int32(123135))
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
	var v8 int32
	_ = v8
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
			v8 = F___memset(m, l0, int32(0), int32(8196))
			mBase = m.M
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
		v8 = F___memset(m, l0, int32(0), int32(8196))
		mBase = m.M
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
	var v46 int32
	_ = v46
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
				v46 = v20
				m.G0 = v8 + int32(16)
				return v46
			} else {
				v27 = int32(0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
				v35 = F_pgp_cfb_create(m, v8+int32(12), v30, l1+int32(132), v33, v27, int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 < int32(0) {
						v46 = v35
						m.G0 = v8 + int32(16)
						return v46
					} else {
						v40 = F_palloc0(m, int32(8196))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
							v46 = int32(8192)
							m.G0 = v8 + int32(16)
							return v46
						}
					}
				}
			}
		}
	} else {
		v27 = int32(1)
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+164))
		v35 = F_pgp_cfb_create(m, v8+int32(12), v30, l1+int32(132), v33, v27, int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 < int32(0) {
				v46 = v35
				m.G0 = v8 + int32(16)
				return v46
			} else {
				v40 = F_palloc0(m, int32(8196))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v40))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
					v46 = int32(8192)
					m.G0 = v8 + int32(16)
					return v46
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
	var v47 int32
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
	var v302 float64
	_ = v302
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if l1 == int32(0) {
		v27 = v17
		v34 = F_get_restriction_variable(m, v19, v16, v15, v13+int32(16), v13+int32(12), v13+int32(11))
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
				v302 = v40
				m.G0 = v13 + int32(48)
				return v302
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
				if v42 == int32(7) {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
					v50 = F_var_eq_const(m, v13+int32(16), v27, v18, v47, v48, v49, l1)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return float64(0)
					} else {
						v293 = v50
						v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						if v294 == int32(0) {
							v302 = v293
							m.G0 = v13 + int32(48)
							return v302
						} else {
							v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
							m.T0[v297].(func(*base.Module, int32))(m, v294)
							mBase = m.M
							v299 = m.ExcPending
							if v299 != 0 {
								return float64(0)
							} else {
								v302 = v293
								m.G0 = v13 + int32(48)
								return v302
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
												switch v117 - int32(65530) {
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
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
													switch v117 - int32(65530) {
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
									v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
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
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
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
												switch v223 - int32(65530) {
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
										v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
													switch v223 - int32(65530) {
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
									v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
								v302 = v293
								m.G0 = v13 + int32(48)
								return v302
							} else {
								v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
								m.T0[v297].(func(*base.Module, int32))(m, v294)
								mBase = m.M
								v299 = m.ExcPending
								if v299 != 0 {
									return float64(0)
								} else {
									v302 = v293
									m.G0 = v13 + int32(48)
									return v302
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
													switch v117 - int32(65530) {
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
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
														switch v117 - int32(65530) {
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
										v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
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
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
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
													switch v223 - int32(65530) {
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
											v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
														switch v223 - int32(65530) {
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
										v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
									v302 = v293
									m.G0 = v13 + int32(48)
									return v302
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
									}
								}
							}
						} else {
							v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+120))
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
														switch v117 - int32(65530) {
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
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
															switch v117 - int32(65530) {
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
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
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
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v302 = v293
														m.G0 = v13 + int32(48)
														return v302
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
														switch v223 - int32(65530) {
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
												v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
															switch v223 - int32(65530) {
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
											v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
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
									v302 = v293
									m.G0 = v13 + int32(48)
									return v302
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v22 = F_get_negator(m, v17)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return float64(0)
		} else {
			if v22 != 0 {
				v27 = v22
				v34 = F_get_restriction_variable(m, v19, v16, v15, v13+int32(16), v13+int32(12), v13+int32(11))
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
						v302 = v40
						m.G0 = v13 + int32(48)
						return v302
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						if v42 == int32(7) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
							v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+24)))
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
							v50 = F_var_eq_const(m, v13+int32(16), v27, v18, v47, v48, v49, l1)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return float64(0)
							} else {
								v293 = v50
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								if v294 == int32(0) {
									v302 = v293
									m.G0 = v13 + int32(48)
									return v302
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
									m.T0[v297].(func(*base.Module, int32))(m, v294)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return float64(0)
									} else {
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
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
														switch v117 - int32(65530) {
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
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
															switch v117 - int32(65530) {
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
											v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
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
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v302 = v293
														m.G0 = v13 + int32(48)
														return v302
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
														switch v223 - int32(65530) {
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
												v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
															switch v223 - int32(65530) {
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
											v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
										v302 = v293
										m.G0 = v13 + int32(48)
										return v302
									} else {
										v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
										m.T0[v297].(func(*base.Module, int32))(m, v294)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return float64(0)
										} else {
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
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
															switch v117 - int32(65530) {
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
													v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
																switch v117 - int32(65530) {
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
												v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
												} else {
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
													m.T0[v297].(func(*base.Module, int32))(m, v294)
													mBase = m.M
													v299 = m.ExcPending
													if v299 != 0 {
														return float64(0)
													} else {
														v302 = v293
														m.G0 = v13 + int32(48)
														return v302
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
														v302 = v293
														m.G0 = v13 + int32(48)
														return v302
													} else {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
														m.T0[v297].(func(*base.Module, int32))(m, v294)
														mBase = m.M
														v299 = m.ExcPending
														if v299 != 0 {
															return float64(0)
														} else {
															v302 = v293
															m.G0 = v13 + int32(48)
															return v302
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
															switch v223 - int32(65530) {
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
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
																switch v223 - int32(65530) {
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
												v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											}
										}
									}
								} else {
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+120))
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
																switch v117 - int32(65530) {
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
														v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+76))
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
																	switch v117 - int32(65530) {
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
													v140 = *(*float64)(unsafe.Add(mBase, uint32(v134)+120))
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
														v302 = v293
														m.G0 = v13 + int32(48)
														return v302
													} else {
														v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
														m.T0[v297].(func(*base.Module, int32))(m, v294)
														mBase = m.M
														v299 = m.ExcPending
														if v299 != 0 {
															return float64(0)
														} else {
															v302 = v293
															m.G0 = v13 + int32(48)
															return v302
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
															v302 = v293
															m.G0 = v13 + int32(48)
															return v302
														} else {
															v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
															m.T0[v297].(func(*base.Module, int32))(m, v294)
															mBase = m.M
															v299 = m.ExcPending
															if v299 != 0 {
																return float64(0)
															} else {
																v302 = v293
																m.G0 = v13 + int32(48)
																return v302
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
																switch v223 - int32(65530) {
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
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v209)+76))
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
																	switch v223 - int32(65530) {
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
													v246 = *(*float64)(unsafe.Add(mBase, uint32(v240)+120))
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
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
												m.T0[v297].(func(*base.Module, int32))(m, v294)
												mBase = m.M
												v299 = m.ExcPending
												if v299 != 0 {
													return float64(0)
												} else {
													v302 = v293
													m.G0 = v13 + int32(48)
													return v302
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
											v302 = v293
											m.G0 = v13 + int32(48)
											return v302
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
											m.T0[v297].(func(*base.Module, int32))(m, v294)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return float64(0)
											} else {
												v302 = v293
												m.G0 = v13 + int32(48)
												return v302
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v302 = float64(0.995)
				m.G0 = v13 + int32(48)
				return v302
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
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1404])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(475542), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(523098), int32(1457), int32(64945))
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
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_consts[1417]))) = uint8(v26)
		return
	}
}
func F_errmsg_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(4555004)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1403]))
	*(*int32)(unsafe.Add(mBase, _consts[1403])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(4562080)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v25 = v18 * int32(100)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1405])))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1406]))) = l0
	F_initStringInfo(m, v9+int32(16))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1404])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L21
	}
L4:
	;
	return
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1407])))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v41 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v47 = v41
	goto L10
L8:
	;
	goto L9
L9:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1408])))
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	F_enlargeStringInfo(m, v9+int32(16), v47)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1407])))
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v59 = F_appendStringInfoVA(m, v9+int32(16), l0, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	if v59 != 0 {
		v47 = v59
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	F_pfree(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v71 = F_pstrdup(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+uint32(_consts[1408]))) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	F_pfree(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v22
	v79 = int32(4555004)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[1403]))
	*(*int32)(unsafe.Add(mBase, _consts[1403])) = v81 - int32(1)
	m.G0 = v9 + int32(32)
	return
L21:
	;
	F_errmsg_internal(m, int32(475542), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(523098), int32(1164), int32(327416))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_error_multiple_recovery_targets(m *base.Module) {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
			F_errmsg(m, int32(478741), int32(0))
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				F_errdetail(m, int32(610187), int32(0))
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					F_errfinish(m, int32(515869), int32(4826), int32(132951))
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[1404])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(475542), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523098), int32(1473), int32(262184))
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
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_consts[1418]))) = l0
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v5 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
	if v5 <= v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errfinish(m, l1, l2, l3)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L27
	} else {
		goto L36
	}
L2:
	;
	v12 = v8 * int32(100)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1409])))
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
	*(*int32)(unsafe.Add(mBase, _consts[1404])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L33
	}
L5:
	;
	v18 = int32(4555004)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1403]))
	*(*int32)(unsafe.Add(mBase, _consts[1403])) = v20 + int32(1)
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1414]))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1413]))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1412]))) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[1409]))) = int32(21)
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
	goto L30
L29:
	;
	v90 = int32(4164644)
	v92 = *(*int32)(unsafe.Add(mBase, _consts[1404]))
	v93 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1404])) = v92 - v93
	v96 = int32(4555004)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1403]))
	*(*int32)(unsafe.Add(mBase, _consts[1403])) = v98 - v93
	return
L30:
	;
	v88 = F__emscripten_memcpy_bulkmem(m, v84, v12+int32(4555008), int32(100))
	mBase = m.M
	goto L32
L32:
	;
	goto L29
L33:
	;
	F_errmsg_internal(m, int32(475542), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(523098), int32(689), int32(338354))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L27
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
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v57 int64
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v10) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v12 = l0
	v16 = v5
	goto L5
L3:
	;
	v57 = v5
	goto L4
L4:
	;
	m.G0 = v8 + int32(16)
	return v57
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v18 != int32(92) {
		v48 = int32(1)
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v57 = v50
	goto L4
L7:
	;
	v50 = v16 + int64(1)
	v51 = v12 + v48
	if base.Ui32(v51) < base.Ui32(v10) {
		v12 = v51
		v16 = v50
		goto L5
	} else {
		goto L16
	}
L8:
	;
	v22 = v12 + int32(3)
	if base.Ui32(v10) <= base.Ui32(v22) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = v12 + int32(1)
	if base.Ui32(v10) <= base.Ui32(v41) {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v24&int32(252) != int32(48) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v29&int32(248) != int32(48) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v34&int32(248) != int32(48) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v48 = int32(4)
	goto L7
L14:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v43 != int32(92) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(2)
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
	v70 = m.ExcPending
	if v70 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(533208)
	F_errmsg(m, int32(199922), v8)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(524498), int32(578), int32(295803))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v62 int64
	_ = v62
	v4 = int64(0)
	v7 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v7) {
		v9 = l0
		v11 = l2
		v12 = v4
		for {
			v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
			if v15 <= int32(0) {
				v18 = int32(92)
				*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v18)
				v20 = int32(7)
				v22 = int32(48)
				v23 = v15&v20 | v22
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v23)
				v30 = int32(base.Ui32(v15)>>(uint(int32(3))%32))&v20 | v22
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+2)) = uint8(v30)
				v37 = int32(base.Ui32(v15&int32(192))>>(uint(int32(6))%32)) | v22
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)) = uint8(v37)
				v53 = int64(4)
				v54 = v11 + int32(4)
			} else {
				if v15 == int32(92) {
					v44 = int32(23644)
					*(*uint16)(unsafe.Add(mBase, uint32(v11))) = uint16(v44)
					v53 = int64(2)
					v54 = v11 + int32(2)
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v15)
					v53 = int64(1)
					v54 = v11 + int32(1)
				}
			}
			v55 = v12 + v53
			v57 = v9 + int32(1)
			if v57 != v7 {
				v9 = v57
				v11 = v54
				v12 = v55
				continue
			} else {
				break
			}
			break
		}
		v62 = v55
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v103 <= v104+int32(1) {
		goto L40
	} else {
		goto L41
	}
L10:
	;
	v40 = v35 & int32(255)
	switch v40 - int32(8) {
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
	v95 = v34 + int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96 != 0 {
		v34 = v95
		v35 = v96
		goto L10
	} else {
		goto L38
	}
L13:
	;
	F_appendStringInfoString(m, l0, int32(533822))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L37
	}
L14:
	;
	v63 = base.I32_extend8_s(v35)
	if base.Ui32(v40) <= base.Ui32(int32(31)) {
		goto L29
	} else {
		goto L30
	}
L15:
	;
	if v40 == int32(92) {
		goto L13
	} else {
		goto L28
	}
L16:
	;
	F_appendStringInfoString(m, l0, int32(764611))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L27
	}
L17:
	;
	F_appendStringInfoString(m, l0, int32(120333))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L26
	}
L18:
	;
	F_appendStringInfoString(m, l0, int32(242521))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L25
	}
L19:
	;
	F_appendStringInfoString(m, l0, int32(299399))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L24
	}
L20:
	;
	F_appendStringInfoString(m, l0, int32(357057))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	F_appendStringInfoString(m, l0, int32(529620))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v63
	F_appendStringInfo(m, l0, int32(30791), v9)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v70 <= v71+int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L12
L33:
	;
	F_appendStringInfoChar(m, l0, v63)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v71))) = uint8(v35)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = v80 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84+v82))) = uint8(v86)
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
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v113 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v111+v104))) = uint8(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v117 = v115 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119+v117))) = uint8(v121)
	goto L39
L43:
	;
	goto L39
}
func F_estimateHyperLogLog(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 float64
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 float64
	_ = v48
	var v50 int32
	_ = v50
	var v57 float64
	_ = v57
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v188 float64
	_ = v188
	var v192 float64
	_ = v192
	v2 = float64(0)
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 != 0 {
		v12 = int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v11 == v12 {
			v48 = v2
			v50 = v4
		} else {
			v20 = v2
			v22 = v4
			v24 = v4
			for {
				v30 = v22 + v14
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
				v32 = F_scalbn(m, float64(1), v31)
				mBase = m.M
				v33 = float64(1)
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				v36 = F_scalbn(m, v33, v35)
				mBase = m.M
				v41 = base.F64_add(base.F64_add(v20, base.F64_div(v33, v36)), base.F64_div(float64(1), v32))
				v42 = int32(2)
				v43 = v22 + v42
				v45 = v24 + v42
				if v45 != v11&int32(-2) {
					v20 = v41
					v22 = v43
					v24 = v45
					continue
				} else {
					break
				}
				break
			}
			v48 = v41
			v50 = v43
		}
		if v11&v12 != 0 {
			v57 = float64(1)
			v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v14))))
			v61 = F_scalbn(m, v57, v60)
			mBase = m.M
			v64 = base.F64_add(v48, base.F64_div(v57, v61))
		} else {
			v64 = v48
		}
		v65 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v66 = base.F64_div(v65, v64)
		v67 = base.F64_convert_i32_u(v11)
		if base.F64_le(v66, base.F64_mul(v67, float64(2.5))) == int32(0) {
			v171 = v66
			if base.F64_gt(v171, float64(1.4316557653333333e+08)) == int32(0) {
				v192 = v171
			} else {
				v188 = F_log(m, base.F64_add(base.F64_mul(v171, float64(-2.3283064365386963e-10)), float64(1)))
				mBase = m.M
				v192 = base.F64_mul(v188, float64(-4.294967296e+09))
			}
			return v192
		} else {
			v74 = v11 & int32(3)
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v76 = int32(0)
			if base.Ui32(int32(4)) <= base.Ui32(v11) {
				v87 = v76
				v88 = int32(0)
				v89 = v76
				for {
					v94 = v87 + v75
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94))))
					v96 = int32(0)
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+2)))
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+3)))
					v110 = v89 + base.B2i32(v95 == v96) + base.B2i32(v99 == v96) + base.B2i32(v103 == v96) + base.B2i32(v107 == v96)
					v111 = int32(4)
					v112 = v87 + v111
					v114 = v88 + v111
					if v114 != v11&int32(-4) {
						v87 = v112
						v88 = v114
						v89 = v110
						continue
					} else {
						break
					}
					break
				}
				v119 = v112
				v121 = v110
			} else {
				v119 = v76
				v121 = v76
			}
			if v74 != 0 {
				v129 = v119
				v131 = v121
				v132 = v76
				for {
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v75))))
					v140 = v131 + base.B2i32(v137 == int32(0))
					v141 = int32(1)
					v144 = v132 + v141
					if v144 != v74 {
						v129 = v129 + v141
						v131 = v140
						v132 = v144
						continue
					} else {
						break
					}
					break
				}
				v151 = v140
			} else {
				v151 = v121
			}
			if v151 == int32(0) {
				v192 = v66
				return v192
			} else {
				v160 = F_log(m, base.F64_div(v67, base.F64_convert_i32_s(v151)))
				mBase = m.M
				return base.F64_mul(v160, v67)
			}
		}
	} else {
		v163 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		v165 = base.F64_div(v163, float64(0))
		if base.F64_le(v165, base.F64_mul(base.F64_convert_i32_u(v11), float64(2.5))) != 0 {
			v192 = v165
		} else {
			v171 = v165
			if base.F64_gt(v171, float64(1.4316557653333333e+08)) == int32(0) {
				v192 = v171
			} else {
				v188 = F_log(m, base.F64_add(base.F64_mul(v171, float64(-2.3283064365386963e-10)), float64(1)))
				mBase = m.M
				v192 = base.F64_mul(v188, float64(-4.294967296e+09))
			}
		}
		return v192
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
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v8 = v6
	} else {
		v8 = int32(0)
	}
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v11 != 0 {
		v12 = int32(1)
		if v11&(v11-v12) != 0 {
			v20 = v12 << (uint(int32(32)-base.I32_clz(v11)) % 32)
		} else {
			v20 = v11
		}
		v24 = v20 + int32(8)
	} else {
		v24 = int32(0)
	}
	return base.F64_mul(l3, base.F64_convert_i32_u(v24+((v10+int32(23))&int32(-8)+v8<<(uint(int32(3))%32))+int32(12)))
}
func F_estimate_num_groups(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v35 float64
	_ = v35
	var v39 float64
	_ = v39
	var v42 int32
	_ = v42
	var v45 float64
	_ = v45
	var v47 int32
	_ = v47
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 float64
	_ = v120
	var v123 int32
	_ = v123
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v252 float64
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 float64
	_ = v334
	var v338 float64
	_ = v338
	var v343 float64
	_ = v343
	var v346 float64
	_ = v346
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 float64
	_ = v376
	var v379 float64
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 float64
	_ = v397
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 float64
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 float64
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 float64
	_ = v416
	var v418 float64
	_ = v418
	var v420 int32
	_ = v420
	var v421 float64
	_ = v421
	var v422 int32
	_ = v422
	var v425 float64
	_ = v425
	var v428 float64
	_ = v428
	var v436 int32
	_ = v436
	var v441 float64
	_ = v441
	var v449 float64
	_ = v449
	var v457 float64
	_ = v457
	var v461 float64
	_ = v461
	var v465 float64
	_ = v465
	var v474 float64
	_ = v474
	var v479 float64
	_ = v479
	var v487 float64
	_ = v487
	var v491 float64
	_ = v491
	var v493 float64
	_ = v493
	var v501 float64
	_ = v501
	var v505 float64
	_ = v505
	var v514 float64
	_ = v514
	var v527 float64
	_ = v527
	var v529 float64
	_ = v529
	var v532 float64
	_ = v532
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v553 float64
	_ = v553
	var v555 float64
	_ = v555
	var v558 float64
	_ = v558
	var v564 float64
	_ = v564
	v12 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
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
	v25 = float64(1)
	v27 = float64(1e+100)
	if base.F64_gt(l2, v27) != 0 {
		v39 = v27
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if l1 == int32(0) {
		v564 = v25
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L4
L6:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l2)&int64(9223372036854775807)) {
		v39 = v27
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v35 = float64(1)
	if base.F64_le(l2, v35) != 0 {
		v39 = v35
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = base.F64_nearest(l2)
	goto L5
L9:
	;
	m.G0 = v21 + int32(48)
	return v564
L10:
	;
	if l3 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v42 == int32(0) {
		v564 = v25
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v45 = float64(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v47 <= int32(0) {
		v540 = v45
		v541 = v45
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v553 = base.F64_ceil(base.F64_mul(v540, v541))
	if base.F64_gt(v553, v39) != 0 {
		goto L141
	} else {
		goto L142
	}
L16:
	;
	v57 = v45
	v58 = v45
	v63 = v12
	v64 = v12
	v65 = v12
	goto L17
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v64<<(uint(int32(2))%32))))
	if l3 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	if v236 == int32(0) {
		v540 = v228
		v541 = v229
		goto L15
	} else {
		goto L74
	}
L19:
	;
	v240 = v64 + int32(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v240 < v241 {
		v57 = v228
		v58 = v229
		v63 = v234
		v64 = v240
		v65 = v236
		goto L17
	} else {
		goto L73
	}
L20:
	;
	v228 = v210
	v229 = v211
	v234 = v219
	v236 = v218
	goto L19
L21:
	;
	v120 = F_expression_returns_set_rows(m, l0, v72)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L39
	} else {
		goto L40
	}
L22:
	;
	v119 = v63
	goto L21
L23:
	;
	goto L24
L24:
	;
	v76 = v63 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v78 = int32(0)
	if v77 == v78 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v116 == int32(0) {
		v210 = v57
		v211 = v58
		v218 = v65
		v219 = v76
		goto L20
	} else {
		goto L38
	}
L26:
	;
	v116 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v84 <= int32(0) {
		v109 = v78
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v116 = v109
	goto L25
L30:
	;
	v87 = int32(0)
	if v87 < v84 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v90 = v84
	goto L33
L32:
	;
	v90 = v87
	goto L33
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v93 = int32(0)
	goto L34
L34:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91+v93<<(uint(int32(2))%32))))
	v102 = base.B2i32(v101 == v63)
	if v101 == v63 {
		v109 = v102
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v109 = v102
	goto L29
L36:
	;
	v104 = v93 + int32(1)
	if v104 != v90 {
		v93 = v104
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v119 = v76
	goto L21
L39:
	;
	return float64(0)
L40:
	;
	if base.F64_gt(v120, v58) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v125 = v120
	goto L43
L42:
	;
	v125 = v58
	goto L43
L43:
	;
	v126 = F_exprType(m, v72)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	if v126 == int32(16) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v210 = base.F64_add(v57, v57)
	v211 = v125
	v218 = v65
	v219 = v119
	goto L20
L46:
	;
	goto L47
L47:
	;
	F_examine_variable(m, l0, v72, int32(0), v21+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v136 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v153 = F_pull_var_clause(m, v72, int32(42))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L39
	} else {
		goto L57
	}
L50:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+44)))
	if v139 != int32(1) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v144 = F_add_unique_group_var(m, l0, v65, v72, v21+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L39
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v146 == int32(0) {
		v210 = v57
		v211 = v125
		v218 = v144
		v219 = v119
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	m.T0[v149].(func(*base.Module, int32))(m, v146)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L39
	} else {
		goto L56
	}
L56:
	;
	v210 = v57
	v211 = v125
	v218 = v144
	v219 = v119
	goto L20
L57:
	;
	if v153 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v155 <= int32(0) {
		v210 = v57
		v211 = v125
		v218 = v65
		v219 = v119
		goto L20
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v199 = F_contain_volatile_functions(m, v72)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L39
	} else {
		goto L71
	}
L61:
	;
	v171 = int32(0)
	v174 = v65
	goto L62
L62:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v171<<(uint(int32(2))%32))))
	F_examine_variable(m, l0, v181, int32(0), v21+int32(16))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L39
	} else {
		goto L64
	}
L63:
	;
	v210 = v57
	v211 = v125
	v218 = v189
	v219 = v119
	goto L20
L64:
	;
	v189 = F_add_unique_group_var(m, l0, v174, v181, v21+int32(16))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L39
	} else {
		goto L65
	}
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v191 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	m.T0[v192].(func(*base.Module, int32))(m, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L39
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v196 = v171 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v196 < v197 {
		v171 = v196
		v174 = v189
		goto L62
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	goto L63
L71:
	;
	if v199 == int32(0) {
		v228 = v57
		v229 = v125
		v234 = v119
		v236 = v65
		goto L19
	} else {
		goto L72
	}
L72:
	;
	v564 = v39
	goto L9
L73:
	;
	goto L18
L74:
	;
	v252 = v228
	v260 = v236
	goto L75
L75:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v266 = int32(0)
	v268 = F_lappend(m, v266, v264)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L39
	} else {
		goto L77
	}
L76:
	;
	v527 = base.F64_ceil(base.F64_mul(v229, v514))
	if base.F64_gt(v527, v39) != 0 {
		goto L135
	} else {
		goto L136
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v268
	v271 = int32(0)
	v273 = int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v273 < v274 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v278 = v271
	v289 = v273
	goto L81
L79:
	;
	v329 = v268
	v333 = v271
	goto L80
L80:
	;
	v334 = float64(1)
	if v329 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295+v289<<(uint(int32(2))%32))))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	if v300 == v301 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v329 = v314
	v333 = v309
	goto L80
L83:
	;
	v311 = v289 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v311 < v312 {
		v278 = v309
		v289 = v311
		goto L81
	} else {
		goto L89
	}
L84:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v304 = F_lappend(m, v303, v299)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L39
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v307 = F_lappend(m, v278, v299)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L39
	} else {
		goto L88
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v304
	v309 = v278
	goto L83
L88:
	;
	v309 = v307
	goto L83
L89:
	;
	goto L82
L90:
	;
	if v333 != 0 {
		v252 = v514
		v260 = v333
		goto L75
	} else {
		goto L134
	}
L91:
	;
	if base.F64_gt(v457, v465) != 0 {
		goto L124
	} else {
		goto L125
	}
L92:
	;
	v338 = *(*float64)(unsafe.Add(mBase, uint32(v265)+120))
	if base.F64_gt(v338, float64(0)) != 0 {
		v457 = v334
		v461 = v338
		v465 = v338
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v343 = v334
	v346 = v334
	v354 = v266
	goto L96
L95:
	;
	v514 = v252
	goto L90
L96:
	;
	v363 = F_estimate_multivariate_ndistinct(m, l0, v265, v21+int32(12), v21+int32(16))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L39
	} else {
		goto L99
	}
L97:
	;
	v441 = *(*float64)(unsafe.Add(mBase, uint32(v265)+120))
	if base.F64_gt(v441, float64(0)) == int32(0) {
		v514 = v252
		goto L90
	} else {
		goto L118
	}
L98:
	;
	goto L97
L99:
	;
	if v363 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v367 == int32(0) {
		v425 = v343
		v428 = v346
		v436 = v354
		goto L98
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v416 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	if base.F64_lt(v346, v416) != 0 {
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v370 = int32(0)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v371 <= v370 {
		v425 = v343
		v428 = v346
		v436 = v354
		goto L98
	} else {
		goto L104
	}
L104:
	;
	v376 = v343
	v379 = v346
	v386 = v370
	v387 = v354
	goto L105
L105:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392+v386<<(uint(int32(2))%32))))
	v397 = *(*float64)(unsafe.Add(mBase, uint32(v396)+8))
	if l4 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v425 = v411
	v428 = v408
	v436 = v410
	goto L98
L107:
	;
	if base.F64_lt(v379, v397) != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+16)))
	if v401 != int32(1) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v404 | int32(1)
	goto L107
L110:
	;
	v408 = v397
	goto L112
L111:
	;
	v408 = v379
	goto L112
L112:
	;
	v409 = int32(1)
	v410 = v387 + v409
	v411 = base.F64_mul(v376, v397)
	v413 = v386 + v409
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v413 < v414 {
		v376 = v411
		v379 = v408
		v386 = v413
		v387 = v410
		goto L105
	} else {
		goto L113
	}
L113:
	;
	goto L106
L114:
	;
	v418 = v416
	goto L116
L115:
	;
	v418 = v346
	goto L116
L116:
	;
	v420 = v354 + int32(1)
	v421 = base.F64_mul(v343, v416)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v422 != 0 {
		v343 = v421
		v346 = v418
		v354 = v420
		goto L96
	} else {
		goto L117
	}
L117:
	;
	v425 = v421
	v428 = v418
	v436 = v420
	goto L98
L118:
	;
	if v436 < int32(2) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v457 = v425
	v461 = v441
	v465 = v441
	goto L91
L120:
	;
	v449 = base.F64_mul(v441, float64(0.1))
	if base.F64_lt(v449, v428) == int32(0) {
		v457 = v425
		v461 = v441
		v465 = v449
		goto L91
	} else {
		goto L121
	}
L121:
	;
	if base.F64_gt(v428, v441) != 0 {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	v457 = v425
	v461 = v441
	v465 = v428
	goto L91
L123:
	;
	v493 = float64(1e+100)
	if base.F64_gt(v491, v493) != 0 {
		v505 = v493
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v474 = v465
	goto L126
L125:
	;
	v474 = v457
	goto L126
L126:
	;
	if base.F64_gt(v474, float64(0)) == int32(0) {
		v491 = v474
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v479 = *(*float64)(unsafe.Add(mBase, uint32(v265)+16))
	if base.F64_lt(v479, v461) == int32(0) {
		v491 = v474
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v487 = F_pow(m, base.F64_div(base.F64_sub(v461, v479), v461), base.F64_div(v461, v474))
	mBase = m.M
	v491 = base.F64_mul(v474, base.F64_sub(float64(1), v487))
	goto L123
L129:
	;
	v514 = base.F64_mul(v252, v505)
	goto L90
L130:
	;
	goto L129
L131:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v491)&int64(9223372036854775807)) {
		v505 = v493
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v501 = float64(1)
	if base.F64_le(v491, v501) != 0 {
		v505 = v501
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v505 = base.F64_nearest(v491)
	goto L130
L134:
	;
	goto L76
L135:
	;
	v529 = v39
	goto L137
L136:
	;
	v529 = v527
	goto L137
L137:
	;
	if base.F64_lt(v529, float64(1)) != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v532 = float64(1)
	goto L140
L139:
	;
	v532 = v529
	goto L140
L140:
	;
	v564 = v532
	goto L9
L141:
	;
	v555 = v39
	goto L143
L142:
	;
	v555 = v553
	goto L143
L143:
	;
	if base.F64_lt(v555, float64(1)) != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v558 = float64(1)
	goto L146
L145:
	;
	v558 = v555
	goto L146
L146:
	;
	v564 = v558
	goto L9
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
	var v12 int32
	_ = v12
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
	v12 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)) = uint8(v12)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v20 = v13 + v14<<(uint(int32(4))%32) + l1*int32(100)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+11)))
	if v21 != 0 {
		v153 = v4
		m.G0 = v11 + int32(32)
		return v153
	} else {
		v23 = v20 - int32(80)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+90)))
		if v24 == int32(118) {
			v153 = v4
			m.G0 = v11 + int32(32)
			return v153
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v30 = F_SearchSysCache2(m, int32(7), v28, base.I32_extend16_s(l1))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return int32(0)
					} else {
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v165
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
						F_errmsg_internal(m, int32(50216), v11)
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523325), int32(1064), int32(366567))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v40 = F_SysCacheGetAttr(m, int32(7), v30, int32(21), v11+int32(31))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)))
						if v42 == int32(1) {
							F_ReleaseCatCache(m, v30)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v55 = int32(-1)
								v57 = F_palloc0(m, int32(248))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
									if l2 != 0 {
										v60 = F_exprType(m, l2)
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v60
											v63 = F_exprTypmod(m, l2)
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v63
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
												v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+(l1-int32(1))<<(uint(int32(2))%32))))
												if v72 != 0 {
													v80 = v72
													*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														if v85 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v179 = m.ExcPending
															if v179 != 0 {
																return int32(0)
															} else {
																v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
																F_errmsg_internal(m, int32(55200), v11+int32(16))
																mBase = m.M
																v186 = m.ExcPending
																if v186 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(523325), int32(1113), int32(366567))
																	mBase = m.M
																	v191 = m.ExcPending
																	if v191 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
															v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
															v91 = v89 + v90
															*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
															v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
															*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
															v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
															v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
															*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
															v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
															v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
															v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
															*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
															v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
															v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
															*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
															v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
															v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
															v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
															*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
															v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
															v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
															*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
															v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
															*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
															v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
															v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
															*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
															if v133 != 0 {
																v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	if v135 != 0 {
																		v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																		if v139 == int32(0) {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																			if int32(0) < v142 {
																				v153 = v57
																				m.G0 = v11 + int32(32)
																				return v153
																			} else {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			}
																		}
																	} else {
																		F_pfree(m, v85)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v57)
																			mBase = m.M
																			v148 = m.ExcPending
																			if v148 != 0 {
																				return int32(0)
																			} else {
																				v153 = int32(0)
																				m.G0 = v11 + int32(32)
																				return v153
																			}
																		}
																	}
																}
															} else {
																v137 = F_std_typanalyze(m, v57)
																mBase = m.M
																v138 = m.ExcPending
																if v138 != 0 {
																	return int32(0)
																} else {
																	v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																	if v139 == int32(0) {
																		F_pfree(m, v85)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v57)
																			mBase = m.M
																			v148 = m.ExcPending
																			if v148 != 0 {
																				return int32(0)
																			} else {
																				v153 = int32(0)
																				m.G0 = v11 + int32(32)
																				return v153
																			}
																		}
																	} else {
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																		if int32(0) < v142 {
																			v153 = v57
																			m.G0 = v11 + int32(32)
																			return v153
																		} else {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v73 = F_exprCollation(m, l2)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														v80 = v73
														*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
														v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															if v85 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
																	return int32(0)
																} else {
																	v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
																	F_errmsg_internal(m, int32(55200), v11+int32(16))
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(523325), int32(1113), int32(366567))
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
																v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																v91 = v89 + v90
																*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
																v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
																*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
																v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
																v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
																v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
																v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
																v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
																v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
																v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
																v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
																v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
																v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
																v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
																v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
																v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
																v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
																v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
																if v133 != 0 {
																	v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		if v135 != 0 {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																			if v139 == int32(0) {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																				if int32(0) < v142 {
																					v153 = v57
																					m.G0 = v11 + int32(32)
																					return v153
																				} else {
																					F_pfree(m, v85)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v57)
																						mBase = m.M
																						v148 = m.ExcPending
																						if v148 != 0 {
																							return int32(0)
																						} else {
																							v153 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v153
																						}
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		}
																	}
																} else {
																	v137 = F_std_typanalyze(m, v57)
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int32(0)
																	} else {
																		v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																		if v139 == int32(0) {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																			if int32(0) < v142 {
																				v153 = v57
																				m.G0 = v11 + int32(32)
																				return v153
																			} else {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
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
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
										*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v75
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
										*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v77
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
										v80 = v79
										*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											if v85 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int32(0)
												} else {
													v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
													F_errmsg_internal(m, int32(55200), v11+int32(16))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(523325), int32(1113), int32(366567))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
												v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
												v91 = v89 + v90
												*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
												v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
												*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
												v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
												*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
												v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
												v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
												v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
												*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
												v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
												*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
												v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
												v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
												v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
												*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
												v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
												*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
												v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
												*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
												v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
												v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
												*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
												if v133 != 0 {
													v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														if v135 != 0 {
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
															if v139 == int32(0) {
																F_pfree(m, v85)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v57)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v153 = int32(0)
																		m.G0 = v11 + int32(32)
																		return v153
																	}
																}
															} else {
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																if int32(0) < v142 {
																	v153 = v57
																	m.G0 = v11 + int32(32)
																	return v153
																} else {
																	F_pfree(m, v85)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v57)
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return int32(0)
																		} else {
																			v153 = int32(0)
																			m.G0 = v11 + int32(32)
																			return v153
																		}
																	}
																}
															}
														} else {
															F_pfree(m, v85)
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v57)
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	v153 = int32(0)
																	m.G0 = v11 + int32(32)
																	return v153
																}
															}
														}
													}
												} else {
													v137 = F_std_typanalyze(m, v57)
													mBase = m.M
													v138 = m.ExcPending
													if v138 != 0 {
														return int32(0)
													} else {
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
														if v139 == int32(0) {
															F_pfree(m, v85)
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v57)
																mBase = m.M
																v148 = m.ExcPending
																if v148 != 0 {
																	return int32(0)
																} else {
																	v153 = int32(0)
																	m.G0 = v11 + int32(32)
																	return v153
																}
															}
														} else {
															v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
															if int32(0) < v142 {
																v153 = v57
																m.G0 = v11 + int32(32)
																return v153
															} else {
																F_pfree(m, v85)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v57)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v153 = int32(0)
																		m.G0 = v11 + int32(32)
																		return v153
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
							F_ReleaseCatCache(m, v30)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v40<<(uint(int32(16))%32) == int32(0) {
									v153 = v4
									m.G0 = v11 + int32(32)
									return v153
								} else {
									v55 = base.I32_extend16_s(v40)
									v57 = F_palloc0(m, int32(248))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v57))) = v55
										if l2 != 0 {
											v60 = F_exprType(m, l2)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v60
												v63 = F_exprTypmod(m, l2)
												mBase = m.M
												v64 = m.ExcPending
												if v64 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v63
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
													v72 = *(*int32)(unsafe.Add(mBase, uint32(v66+(l1-int32(1))<<(uint(int32(2))%32))))
													if v72 != 0 {
														v80 = v72
														*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
														v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															if v85 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v179 = m.ExcPending
																if v179 != 0 {
																	return int32(0)
																} else {
																	v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
																	F_errmsg_internal(m, int32(55200), v11+int32(16))
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(523325), int32(1113), int32(366567))
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
																v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																v91 = v89 + v90
																*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
																v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
																*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
																v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
																v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
																v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
																v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
																v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
																v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
																v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
																v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
																v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
																v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
																v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
																v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
																v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
																v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
																v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
																v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
																if v133 != 0 {
																	v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		if v135 != 0 {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																			if v139 == int32(0) {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																				if int32(0) < v142 {
																					v153 = v57
																					m.G0 = v11 + int32(32)
																					return v153
																				} else {
																					F_pfree(m, v85)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v57)
																						mBase = m.M
																						v148 = m.ExcPending
																						if v148 != 0 {
																							return int32(0)
																						} else {
																							v153 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v153
																						}
																					}
																				}
																			}
																		} else {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		}
																	}
																} else {
																	v137 = F_std_typanalyze(m, v57)
																	mBase = m.M
																	v138 = m.ExcPending
																	if v138 != 0 {
																		return int32(0)
																	} else {
																		v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																		if v139 == int32(0) {
																			F_pfree(m, v85)
																			mBase = m.M
																			v146 = m.ExcPending
																			if v146 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v57)
																				mBase = m.M
																				v148 = m.ExcPending
																				if v148 != 0 {
																					return int32(0)
																				} else {
																					v153 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v153
																				}
																			}
																		} else {
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																			if int32(0) < v142 {
																				v153 = v57
																				m.G0 = v11 + int32(32)
																				return v153
																			} else {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v73 = F_exprCollation(m, l2)
														mBase = m.M
														v74 = m.ExcPending
														if v74 != 0 {
															return int32(0)
														} else {
															v80 = v73
															*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
															v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
															v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return int32(0)
															} else {
																if v85 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return int32(0)
																	} else {
																		v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
																		F_errmsg_internal(m, int32(55200), v11+int32(16))
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(523325), int32(1113), int32(366567))
																			mBase = m.M
																			v191 = m.ExcPending
																			if v191 != 0 {
																				return int32(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
																	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
																	v91 = v89 + v90
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
																	v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
																	v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
																	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
																	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
																	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
																	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
																	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
																	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
																	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
																	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
																	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
																	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																	*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
																	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
																	*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
																	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
																	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
																	*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
																	if v133 != 0 {
																		v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			if v135 != 0 {
																				v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																				if v139 == int32(0) {
																					F_pfree(m, v85)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v57)
																						mBase = m.M
																						v148 = m.ExcPending
																						if v148 != 0 {
																							return int32(0)
																						} else {
																							v153 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v153
																						}
																					}
																				} else {
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																					if int32(0) < v142 {
																						v153 = v57
																						m.G0 = v11 + int32(32)
																						return v153
																					} else {
																						F_pfree(m, v85)
																						mBase = m.M
																						v146 = m.ExcPending
																						if v146 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v57)
																							mBase = m.M
																							v148 = m.ExcPending
																							if v148 != 0 {
																								return int32(0)
																							} else {
																								v153 = int32(0)
																								m.G0 = v11 + int32(32)
																								return v153
																							}
																						}
																					}
																				}
																			} else {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			}
																		}
																	} else {
																		v137 = F_std_typanalyze(m, v57)
																		mBase = m.M
																		v138 = m.ExcPending
																		if v138 != 0 {
																			return int32(0)
																		} else {
																			v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																			if v139 == int32(0) {
																				F_pfree(m, v85)
																				mBase = m.M
																				v146 = m.ExcPending
																				if v146 != 0 {
																					return int32(0)
																				} else {
																					F_pfree(m, v57)
																					mBase = m.M
																					v148 = m.ExcPending
																					if v148 != 0 {
																						return int32(0)
																					} else {
																						v153 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v153
																					}
																				}
																			} else {
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																				if int32(0) < v142 {
																					v153 = v57
																					m.G0 = v11 + int32(32)
																					return v153
																				} else {
																					F_pfree(m, v85)
																					mBase = m.M
																					v146 = m.ExcPending
																					if v146 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v57)
																						mBase = m.M
																						v148 = m.ExcPending
																						if v148 != 0 {
																							return int32(0)
																						} else {
																							v153 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v153
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
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
											*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = v75
											v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
											*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v77
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
											v80 = v79
											*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v80
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
											v85 = F_SearchSysCacheCopy(m, int32(82), v83, int32(0))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												if v85 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v179 = m.ExcPending
													if v179 != 0 {
														return int32(0)
													} else {
														v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v180
														F_errmsg_internal(m, int32(55200), v11+int32(16))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(523325), int32(1113), int32(366567))
															mBase = m.M
															v191 = m.ExcPending
															if v191 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
													v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
													v91 = v89 + v90
													*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v91
													v94 = *(*int32)(unsafe.Add(mBase, _consts[395]))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+224)) = l1
													*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v94
													v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+184)) = v97
													v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
													*(*uint16)(unsafe.Add(mBase, uint32(v57)+204)) = uint16(v99)
													v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+214)) = uint8(v101)
													v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+188)) = v97
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+219)) = uint8(v103)
													v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
													*(*uint16)(unsafe.Add(mBase, uint32(v57)+206)) = uint16(v106)
													v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+215)) = uint8(v108)
													v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+192)) = v97
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+220)) = uint8(v110)
													v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
													*(*uint16)(unsafe.Add(mBase, uint32(v57)+208)) = uint16(v113)
													v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+216)) = uint8(v115)
													v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+196)) = v97
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+221)) = uint8(v117)
													v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
													*(*uint16)(unsafe.Add(mBase, uint32(v57)+210)) = uint16(v120)
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+217)) = uint8(v122)
													v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
													*(*int32)(unsafe.Add(mBase, uint32(v57)+200)) = v97
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+222)) = uint8(v124)
													v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+76)))
													*(*uint16)(unsafe.Add(mBase, uint32(v57)+212)) = uint16(v127)
													v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+78)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+218)) = uint8(v129)
													v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+128)))
													*(*uint8)(unsafe.Add(mBase, uint32(v57)+223)) = uint8(v131)
													v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+124))
													if v133 != 0 {
														v135 = F_OidFunctionCall1Coll(m, v133, int32(0), v57)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															if v135 != 0 {
																v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
																if v139 == int32(0) {
																	F_pfree(m, v85)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v57)
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return int32(0)
																		} else {
																			v153 = int32(0)
																			m.G0 = v11 + int32(32)
																			return v153
																		}
																	}
																} else {
																	v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																	if int32(0) < v142 {
																		v153 = v57
																		m.G0 = v11 + int32(32)
																		return v153
																	} else {
																		F_pfree(m, v85)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v57)
																			mBase = m.M
																			v148 = m.ExcPending
																			if v148 != 0 {
																				return int32(0)
																			} else {
																				v153 = int32(0)
																				m.G0 = v11 + int32(32)
																				return v153
																			}
																		}
																	}
																}
															} else {
																F_pfree(m, v85)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v57)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v153 = int32(0)
																		m.G0 = v11 + int32(32)
																		return v153
																	}
																}
															}
														}
													} else {
														v137 = F_std_typanalyze(m, v57)
														mBase = m.M
														v138 = m.ExcPending
														if v138 != 0 {
															return int32(0)
														} else {
															v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
															if v139 == int32(0) {
																F_pfree(m, v85)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v57)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v153 = int32(0)
																		m.G0 = v11 + int32(32)
																		return v153
																	}
																}
															} else {
																v142 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
																if int32(0) < v142 {
																	v153 = v57
																	m.G0 = v11 + int32(32)
																	return v153
																} else {
																	F_pfree(m, v85)
																	mBase = m.M
																	v146 = m.ExcPending
																	if v146 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v57)
																		mBase = m.M
																		v148 = m.ExcPending
																		if v148 != 0 {
																			return int32(0)
																		} else {
																			v153 = int32(0)
																			m.G0 = v11 + int32(32)
																			return v153
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
		}
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
	var v28 int32
	_ = v28
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
	v28 = v7
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
	v32 = v28 << (uint(int32(2)) % 32)
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
	v40 = v28 + int32(1)
	if v40 != l1 {
		v28 = v40
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v11 = l1
	for {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v9+v11<<(uint(int32(2))%32))))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		if v18 != int32(3) {
			break
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
			v11 = v48
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
			F_errstart_cold(m, int32(21), int32(583335))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(83886210))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v33
					F_errmsg(m, int32(544432), v7+int32(16))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(525311), int32(8573), int32(414405))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
		F_errstart_cold(m, int32(21), int32(583335))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v55
			F_errmsg_internal(m, int32(507298), v7)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				F_errfinish(m, int32(525311), int32(8584), int32(414405))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
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
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v126
L2:
	;
	v126 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v29 = F_AllocSetContextCreateInternal(m, v24, int32(530191), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v34 = v21
	goto L7
L7:
	;
	v35 = int32(4562080)
	v36 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v34
	v38 = F_makeParamList(m, v20)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v29
	v34 = v29
	goto L7
L10:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 <= int32(0) {
		v126 = v38
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v52 = int32(0)
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v52<<(uint(int32(2))%32))))
	v65 = int32(12)
	v67 = v38 + int32(32) + v52*v65
	v68 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v67)+6)) = uint16(v68)
	v71 = v67 + int32(4)
	v73 = v67 + int32(8)
	v76 = F_exec_eval_expr(m, l0, v64, v71, v73, v16+v65)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v126 = v38
	goto L1
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v76
	v79 = int32(4562080)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v34
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	if v82 == int32(705) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v80
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v107 != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(25)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v87 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v91 != 0 {
		goto L15
	} else {
		goto L21
	}
L19:
	;
	v88 = F_cstring_to_text(m, v76)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v88
	goto L15
L21:
	;
	F_get_typlenbyval(m, v82, v16+int32(10), v16+int32(9))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)))
	if v98 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+10)))
	v102 = F_datumCopy(m, v99, int32(0), v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v102
	goto L15
L25:
	;
	F_SPI_freetuptable(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
	F_MemoryContextReset(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v117 = v52 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v117 < v118 {
		v52 = v117
		goto L12
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L13
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
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
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
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
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
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	v14 = m.G0
	v16 = v14 - int32(320)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v18 != int32(2) {
		v261 = int32(0)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(320)
	return
L2:
	;
	if l2 == int32(0) {
		goto L94
	} else {
		goto L95
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v87 != int32(2249) {
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v42 != v38 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v72 != 0 {
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
	F_DeleteExpandedObject(m, v72+int32(12))
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
	F_DeleteExpandedObject(m, v78+int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v99 == int32(2249) {
		goto L47
	} else {
		goto L48
	}
L39:
	;
	F_revalidate_rectypeid(m, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v96 = F_make_expanded_record_from_tupdesc(m, l3, v86)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L44
	}
L42:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v94 = F_make_expanded_record_from_typeid(m, v92, int32(-1), v86)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v98 = v94
	goto L38
L44:
	;
	v98 = v96
	goto L38
L45:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v220 != v216 {
		goto L74
	} else {
		goto L75
	}
L46:
	;
	v194 = int32(1)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	F_expanded_record_set_tuple(m, v98, l2, v194, (v195^int32(-1))&v194)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L10
	} else {
		goto L72
	}
L47:
	;
	if l2 != 0 {
		goto L46
	} else {
		goto L70
	}
L48:
	;
	if l2 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v99 == v104 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+44))
	if v106 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v109 = F_expanded_record_fetch_tupdesc(m, v98)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L10
	} else {
		goto L54
	}
L52:
	;
	v111 = v106
	goto L53
L53:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v112 != v113 {
		v261 = v98
		goto L2
	} else {
		goto L55
	}
L54:
	;
	v111 = v109
	goto L53
L55:
	;
	if v112 <= int32(0) {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v118 = v112 << (uint(int32(4)) % 32)
	v120 = int32(20)
	v130 = int32(0)
	goto L57
L57:
	;
	v140 = v130 * int32(100)
	v141 = v111 + v118 + v120 + v140
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+91)))
	v143 = v140 + (l3 + v118 + v120)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+91)))
	if v142 != v144 {
		v261 = v98
		goto L2
	} else {
		goto L59
	}
L58:
	;
	goto L47
L59:
	;
	if v142 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v164 = v130 + int32(1)
	if v164 != v112 {
		v130 = v164
		goto L57
	} else {
		goto L69
	}
L61:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+68))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+68))
	if v148 != v149 {
		v261 = v98
		goto L2
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+72)))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+72)))
	if v156 != v157 {
		v261 = v98
		goto L2
	} else {
		goto L67
	}
L64:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
	if v151 < int32(0) {
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	if v151 == v154 {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v261 = v98
	goto L2
L67:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+83)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+83)))
	if v159 != v160 {
		v261 = v98
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L60
L69:
	;
	goto L58
L70:
	;
	F_deconstruct_expanded_record(m, v98)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	goto L45
L72:
	;
	goto L45
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v250 != 0 {
		goto L90
	} else {
		goto L91
	}
L74:
	;
	if v220 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	if v216 != 0 {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	if v225 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v224 == int32(0) {
		goto L77
	} else {
		goto L83
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+28)) = v224
	goto L79
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+20)) = v224
	goto L79
L83:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+24)) = v230
	goto L77
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v216
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v216)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+28)) = v237
	if v237 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v215)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = int32(0)
	goto L76
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+24)) = v215
	goto L89
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216)+20)) = v215
	goto L73
L90:
	;
	F_DeleteExpandedObject(m, v250+int32(12))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v98
	goto L1
L93:
	;
	goto L92
L94:
	;
	v293 = int32(0)
	F_exec_move_row_from_fields(m, l0, l1, v261, v293, v293, v293)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L104
	}
L95:
	;
	if l3 == int32(0) {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui32(v273) < base.Ui32(int32(65)) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	F_heap_deform_tuple(m, l2, l3, v288, v287)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L102
	}
L98:
	;
	v287 = v16
	v288 = v16 - int32(-64)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v282 = F_MemoryContextAlloc(m, v279, v273*int32(5))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v287 = v282 + v273<<(uint(int32(2))%32)
	v288 = v282
	goto L97
L102:
	;
	F_exec_move_row_from_fields(m, l0, l1, v261, v288, v287, l3)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	goto L1
L104:
	;
	goto L1
}
func F_executeUnaryArithmExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
	F_jspGetArg(m, l1, v15+int32(44))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = F_executeItemOptUnwrapResult(m, l0, v15+int32(44), l2, int32(1), v15+int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v15 + int32(80)
	return v195
L4:
	;
	if v31 == int32(2) {
		v195 = int32(2)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = F_jspGetNext(m, l1, v15+int32(44))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = int32(0)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	if v40 != 0 {
		v55 = v39
		v56 = v40
		v57 = int32(0)
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v60 = base.B2i32(l4 != int32(0)) | v37
	v64 = v55
	v70 = v56
	v72 = int32(1)
	goto L15
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = int32(0)
	v55 = v39
	v56 = v44
	v57 = v44
	goto L7
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if int32(1) < v50 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = v46 + int32(4)
	goto L14
L13:
	;
	v53 = int32(0)
	goto L14
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v55 = v53
	v56 = v54
	v57 = v41
	goto L7
L15:
	;
	v76 = v64
	v80 = v70
	goto L19
L16:
	;
	v195 = v186
	goto L3
L17:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v184 = F_executeItemOptUnwrapTarget(m, l0, v15+int32(44), v80, l4, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L56
	}
L18:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v172 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L19:
	;
	if v76 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v148 != int32(1) {
		v195 = int32(2)
		goto L3
	} else {
		goto L45
	}
L21:
	;
	if v80 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	v88 = int32(0)
	v102 = v88
	v103 = v88
	goto L21
L23:
	;
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v92 = v76 + int32(4)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if base.Ui32(v92) < base.Ui32(v94+v95<<(uint(int32(2))%32)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v100 = v92
	goto L27
L26:
	;
	v100 = int32(0)
	goto L27
L27:
	;
	v102 = v90
	v103 = v100
	goto L21
L28:
	;
	v195 = v72
	goto L3
L29:
	;
	goto L30
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if v106 == int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v109 = int32(0)
	if v60&int32(1) == v109 {
		v195 = v109
		goto L3
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v60&int32(1) == int32(0) {
		v76 = v103
		v80 = v102
		goto L19
	} else {
		goto L44
	}
L34:
	;
	if l3 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v116 = F_DirectFunctionCall1Coll(m, l3, int32(0), v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v121 {
		goto L17
	} else {
		goto L40
	}
L38:
	;
	v118 = F_pg_detoast_datum(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v118
	goto L37
L40:
	;
	if l4 == int32(0) {
		v195 = v109
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v126 == int32(0) {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v80
	v137 = F_list_make2_impl(m, v15+int32(12), v15+int32(8))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v139
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v137
	v64 = v103
	v70 = v102
	v72 = v139
	goto L15
L44:
	;
	goto L20
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(302776450))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v159 = F_jspOperationName(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v159
	F_errmsg(m, int32(364204), v15+int32(16))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(525321), int32(2212), int32(218003))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v80
	v64 = v103
	v70 = v102
	v72 = int32(0)
	goto L15
L52:
	;
	goto L53
L53:
	;
	v177 = F_lappend(m, v172, v80)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v177
	v64 = v103
	v70 = v102
	v72 = int32(0)
	goto L15
L55:
	;
	v186 = int32(0)
	if l4 != 0 {
		v64 = v103
		v70 = v102
		v72 = v186
		goto L15
	} else {
		goto L57
	}
L56:
	;
	switch v184 {
	case 0:
		goto L55
	default:
		v64 = v103
		v70 = v102
		goto L15
	case 2:
		v195 = v184
		goto L3
	}
L57:
	;
	goto L16
}
func F_exists(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_exists(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
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
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
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
	v37 = int32(0)
	v43 = int32(0)
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
	return v160
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
	if v48 <= v37 {
		v54 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v54 = v50 + v37<<(uint(int32(2))%32)
	goto L8
L11:
	;
	goto L7
L12:
	;
	v160 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v58 <= v37 {
		v160 = v43
		goto L11
	} else {
		goto L15
	}
L15:
	;
	if v54 == int32(0) {
		v160 = v43
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v65 = v62 + v37<<(uint(int32(2))%32)
	if v65 == int32(0) {
		v160 = v43
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v71 + int32(1)
	v77 = F_makeTargetEntry(m, v68, base.I32_extend16_s(v71), v70, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v79 = F_lappend(m, v43, v77)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	if v81 == int32(0) {
		v143 = l0
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+8)))
	F_markRTEForSelectPriv(m, v143, v153, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v85 = v81 & int32(7)
	if base.Ui32(int32(8)) <= base.Ui32(v81) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v92 = l0
	v94 = int32(0)
	goto L25
L23:
	;
	v114 = l0
	goto L24
L24:
	;
	v124 = int32(0)
	if v85 == v124 {
		v143 = v114
		goto L20
	} else {
		goto L28
	}
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v111 = v94 + int32(8)
	if v111 != v81&int32(-8) {
		v92 = v109
		v94 = v111
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v114 = v109
	goto L24
L27:
	;
	goto L26
L28:
	;
	v128 = v114
	v130 = v124
	goto L29
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v140 = v130 + int32(1)
	if v140 != v85 {
		v128 = v138
		v130 = v140
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v143 = v138
	goto L20
L31:
	;
	goto L30
L32:
	;
	v37 = v37 + int32(1)
	v43 = v79
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
	var v35 int32
	_ = v35
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
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
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
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
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
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
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
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
	v35 = v5
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v35<<(uint(int32(2))%32))))
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
	v66 = v35 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v66 < v67 {
		v35 = v66
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v74 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+119)))
	if v121 == int32(112) {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	if v110 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	goto L18
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= int32(0) {
		v110 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v110 = int32(0)
	goto L19
L23:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = v78
	goto L26
L25:
	;
	v84 = v81
	goto L26
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v87 = int32(0)
	goto L27
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85+v87<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v96 == l3 {
		v110 = v95
		goto L19
	} else {
		goto L29
	}
L28:
	;
	goto L22
L29:
	;
	v99 = v87 + int32(1)
	if v99 != v84 {
		v87 = v99
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v118 = v5
	v119 = v5
	goto L17
L32:
	;
	goto L33
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+32)))
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+32)) = uint8(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v118 = v117
	v119 = v114
	goto L17
L34:
	;
	if v110 != 0 {
		goto L64
	} else {
		goto L65
	}
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+56))
	v126 = F_getRTEPermissionInfo(m, v125, l2)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v132 = F_find_all_inheritors(m, v69, v73, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L11
	} else {
		goto L40
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+36))
	F_expand_partitioned_rtentry(m, l0, l1, l2, l3, v71, v128, v110, v73)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	if v132 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_expand_planner_arrays(m, l0, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	F_expand_planner_arrays(m, l0, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L11
	} else {
		goto L45
	}
L44:
	;
	goto L34
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v142 <= int32(0) {
		goto L34
	} else {
		goto L46
	}
L46:
	;
	v152 = v5
	goto L47
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v159+v152<<(uint(int32(2))%32))))
	if v69 != v163 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L34
L49:
	;
	v198 = v152 + int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v198 < v199 {
		v152 = v198
		goto L47
	} else {
		goto L63
	}
L50:
	;
	v166 = F_table_open(m, v163, int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L11
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, v71, v110, v71, v17+int32(48), v17+int32(92))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L11
	} else {
		goto L61
	}
L53:
	;
	F_expand_single_inheritance_child(m, l0, l2, l3, v71, v110, v166, v17+int32(48), v17+int32(92))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L58
	}
L54:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)+48))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+118)))
	if v169 != int32(116) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+24)))
	if v172 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	F_sequence_close(m, v166, v73)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	goto L49
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v182 = F_build_simple_rel(m, l0, v181, l1)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	F_sequence_close(m, v166, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	goto L49
L61:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	v194 = F_build_simple_rel(m, l0, v193, l1)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	goto L49
L63:
	;
	goto L48
L64:
	;
	v215 = int32(0)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	if v216&int32(-33) == v215 {
		v264 = v215
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	F_sequence_close(m, v71, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L109
	}
L67:
	;
	if v216&int32(32) == int32(0) {
		v325 = v264
		goto L79
	} else {
		goto L80
	}
L68:
	;
	if v118&int32(-33) != 0 {
		v264 = v215
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v224 = int32(-1)
	v227 = int32(0)
	v229 = F_makeVar(m, v223, v224, int32(27), v224, v227, v227)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v231
	v235 = int32(32)
	v239 = F_pg_snprintf(m, v17+int32(48), v235, int32(40082), v17+v235)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v242 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242)+4)))
	v246 = v243 + int32(1)
	goto L74
L73:
	;
	v246 = int32(1)
	goto L74
L74:
	;
	v250 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L75
	}
L75:
	;
	v253 = F_makeTargetEntry(m, v229, base.I32_extend16_s(v246), v250, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v256 = F_lappend(m, v255, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v256
	v260 = F_lappend(m, int32(0), v229)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	v264 = v260
	goto L67
L79:
	;
	v326 = int32(1)
	if v119&v326 != 0 {
		goto L95
	} else {
		goto L96
	}
L80:
	;
	if v118&int32(32) != 0 {
		v325 = v264
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v271 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v288 = int32(0)
	v290 = F_makeWholeRowVar(m, v287, v285, v288, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L86
	}
L83:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v285 = v272
	v286 = v271 + v272<<(uint(int32(2))%32)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+52))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v285 = v279
	v286 = v278 + v279<<(uint(int32(2))%32) - int32(4)
	goto L82
L86:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v292
	v300 = F_pg_snprintf(m, v17+int32(48), int32(32), int32(40071), v17+int32(16))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v303 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303)+4)))
	v307 = v304 + int32(1)
	goto L90
L89:
	;
	v307 = int32(1)
	goto L90
L90:
	;
	v311 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	v314 = F_makeTargetEntry(m, v290, base.I32_extend16_s(v307), v311, int32(1))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v317 = F_lappend(m, v316, v314)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v317
	v320 = F_lappend(m, v264, v290)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	v325 = v320
	goto L79
L95:
	;
	v367 = v325
	goto L97
L96:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v333 = int32(0)
	v335 = F_makeVar(m, v329, int32(-6), int32(26), int32(-1), v333, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L11
	} else {
		goto L98
	}
L97:
	;
	v369 = F_bms_make_singleton(m, int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L11
	} else {
		goto L107
	}
L98:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v337
	v343 = F_pg_snprintf(m, v17+int32(48), int32(32), int32(40089), v17)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	if v345 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345)+4)))
	v349 = v346 + int32(1)
	goto L102
L101:
	;
	v349 = v326
	goto L102
L102:
	;
	v353 = F_pstrdup(m, v17+int32(48))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	v356 = F_makeTargetEntry(m, v335, base.I32_extend16_s(v349), v353, int32(1))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v359 = F_lappend(m, v358, v356)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v359
	v362 = F_lappend(m, v325, v335)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	v367 = v362
	goto L97
L107:
	;
	F_add_vars_to_targetlist(m, l0, v367, v369)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	goto L66
L109:
	;
	goto L1
}
func F_ext_sibling_callback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v5 = *(*int32)(unsafe.Add(mBase, _consts[440]))
	if v5 != 0 {
		v6 = v5
		for {
			if l2 != 0 {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				if v9 != l2 {
				} else {
					v11 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v11)
				}
			} else {
				v11 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v11)
			}
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v13 != 0 {
				v6 = v13
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
