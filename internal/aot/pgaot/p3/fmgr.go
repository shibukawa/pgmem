package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_internal_validator(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = base.I32_wrap_i64(v15)
	v17 = F_CheckFunctionValidatorAccess(m, v14, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L3
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L29
	}
L3:
	;
	return int64(0)
L4:
	;
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = F_SearchSysCache1(m, int32(47), v15&int64(4294967295))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	m.G0 = v11 + int32(32)
	return int64(0)
L8:
	;
	if v24 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v30 = F_SysCacheGetAttrNotNull(m, int32(47), v24, int32(26))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = F_text_to_cstring(m, base.I32_wrap_i64(v30))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_internal_validator[0]))
	if v36 <= int32(0) {
		v88 = v2
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v88 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L13:
	;
	v42 = v2
	goto L14
L14:
	;
	v48 = v42 << (uint(int32(4)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_fmgr_internal_validator[1])))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.B2i32(v54 == int32(0))|base.B2i32(v54 != v57) != 0 {
		v75 = v54
		v76 = v57
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_fmgr_internal_validator[2])))
	v88 = v81
	goto L12
L16:
	;
	if v75-v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v60 = v33
	v61 = v51
	goto L19
L19:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v65 == int32(0) {
		v75 = v65
		v76 = v64
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v75 = v65
	v76 = v64
	goto L17
L21:
	;
	v68 = int32(1)
	if v65 == v64 {
		v60 = v60 + v68
		v61 = v61 + v68
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v79 = v42 + int32(1)
	if v36 != v79 {
		v42 = v79
		goto L14
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L15
L26:
	;
	v88 = v2
	goto L12
L27:
	;
	F_ReleaseCatCache(m, v24)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
	F_errmsg_internal(m, int32(_a_F_fmgr_internal_validator_0), v11)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_fmgr_internal_validator_1), int32(794), int32(_a_F_fmgr_internal_validator_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
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
	F_errcode(m, int32(52461700))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v33
	F_errmsg(m, int32(_a_F_fmgr_internal_validator_3), v11+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_fmgr_internal_validator_1), int32(803), int32(_a_F_fmgr_internal_validator_2))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fmgr_sql(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v111 int32
	_ = v111
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v213 int32
	_ = v213
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
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
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
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
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int64
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int64
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
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
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v560 int64
	_ = v560
	var v561 int32
	_ = v561
	var v563 int64
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int64
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int64
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v591 int32
	_ = v591
	var v594 int64
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
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
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v641 int64
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v665 int64
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int64
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int64
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int64
	_ = v732
	var v733 int32
	_ = v733
	var v736 int64
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v758 int64
	_ = v758
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v776 int64
	_ = v776
	var v792 int64
	_ = v792
	var v810 int64
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
	if v23 == v21 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L11
	} else {
		goto L211
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v26 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v47 = v21
	v49 = v2
	v50 = v2
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v51 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(389) {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v33 = int32(3)
	if v32&v33 != v33 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v47 = base.B2i32(v32&int32(8) == int32(0))
	v49 = int32(base.Ui32(v32&int32(4)) >> (uint(int32(2)) % 32))
	v50 = v42
	goto L4
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v56 = F_MemoryContextAllocZero(m, v54, int32(84))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v74 = v51
	goto L10
L10:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
	if v76 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return int64(0)
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+72)) = int32(727)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+60)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v56)+76)) = v56
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v67 = v56 + int32(72)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+8)) = v68
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+40)) = v67
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v56
	v74 = v56
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = v50
	v237 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)) = uint8(v237)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+8)) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(730)
	v243 = int32(_a_F_fmgr_sql_0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0])) = v19 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v244
	v251 = v74 + int32(44)
	v254 = v251
	goto L53
L15:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v96 = F_cached_function_compile(m, l0, v90, int32(728), int32(729), int32(88), int32(1), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L11
	} else {
		goto L20
	}
L16:
	;
	v79 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+44)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v74)+32)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v79
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)) = uint8(v79)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)) = uint8(v79)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v74)+44))
	if v89 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v96 != v98 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if v98 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v111 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v98)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v98)+24)) = v100 - int64(1)
	goto L26
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v96
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v96)+24)) = v105 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+24)) = int32(0)
	goto L23
L27:
	;
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+7)) = uint8(v213)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+5)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+56)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v74)+40)) = int64(0)
	goto L14
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+44))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+40))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	if v118 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = int32(0)
	goto L27
L31:
	;
	v121 = int32(_a_F_fmgr_sql_1)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v74)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v124
	v126 = F_makeParamList(m, v111)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L34
	}
L32:
	;
	v131 = v118
	goto L33
L33:
	;
	v140 = int32(0)
	goto L35
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v126
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v122
	v131 = v126
	goto L33
L35:
	;
	v155 = v140 << (uint(int32(4)) % 32)
	v156 = v131 + int32(32) + v155
	v157 = v155 + (l0 + int32(24))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v156)+8)) = uint8(v158)
	if v158 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	goto L27
L37:
	;
	v184 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v156)+10)) = uint16(v184)
	*(*int64)(unsafe.Add(mBase, uint32(v156))) = v183
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v117+v140<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v156)+12)) = v190
	v193 = v140 + v184
	if v193 != v111 {
		v140 = v193
		goto L35
	} else {
		goto L47
	}
L38:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	v171 = base.I32_wrap_i64(v169)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v172 != int32(1) {
		v182 = v169
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115+v140<<(uint(int32(1))%32)))))
	if v165 == int32(_a_F_fmgr_sql_2) {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	v183 = v168
	goto L37
L42:
	;
	goto L41
L43:
	;
	v183 = v182
	goto L37
L44:
	;
	goto L43
L45:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
	if v175 != int32(3) {
		v182 = v169
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v171)+2))
	v182 = base.I64_extend_i32_u(v178 + int32(18))
	goto L44
L47:
	;
	goto L36
L48:
	;
	v811 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)) = uint8(v811)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0])) = v814
	m.G0 = v19 + int32(16)
	return v810
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = int32(0)
	v810 = v792
	goto L48
L50:
	;
	if v763 != 0 {
		v810 = v776
		goto L48
	} else {
		goto L210
	}
L51:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L11
	} else {
		goto L209
	}
L52:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+56)))
	if v709 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L53:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v268 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v631)+55)))
	if v632 != 0 {
		goto L178
	} else {
		goto L179
	}
L55:
	;
	goto L54
L56:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v269 == int32(2) {
		v254 = v268
		goto L53
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v613 = F_init_execution_state(m, v74)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L11
	} else {
		goto L176
	}
L59:
	;
	v274 = v268
	goto L60
L60:
	;
	v291 = v274
	v295 = int32(0)
	goto L63
L61:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+55)))
	if v548 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L62:
	;
	goto L61
L63:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+57)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v307 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v441 != 0 {
		goto L138
	} else {
		goto L139
	}
L65:
	;
	v443 = int32(_a_F_fmgr_sql_1)
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v74)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v449 == int32(6) {
		goto L115
	} else {
		goto L116
	}
L66:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+36))
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[2]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	goto L104
L67:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[3]))
	v407 = v406
	goto L66
L68:
	;
	if v306&int32(1) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if (v306|v295)&int32(1) != 0 {
		v441 = v295
		goto L65
	} else {
		goto L102
	}
L71:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L11
	} else {
		goto L74
	}
L72:
	;
	v325 = v295
	goto L73
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v328 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L74:
	;
	if v295 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v325 = int32(1)
	goto L73
L76:
	;
	v318 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L11
	} else {
		goto L81
	}
L79:
	;
	F_PushActiveSnapshot(m, v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	goto L75
L82:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+9)) = uint8(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v74)+68)) = v352
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	if v355 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v348 = F_AllocSetContextCreateInternal(m, v341, int32(_a_F_fmgr_sql_3), int32(0), int32(_a_F_fmgr_sql_4), int32(_a_F_fmgr_sql_5))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L11
	} else {
		goto L88
	}
L84:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v338 != int32(6) {
		v351 = int32(0)
		v352 = v327
		goto L82
	} else {
		goto L87
	}
L85:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+55)))
	if v332 != int32(1) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v74)+60))
	v341 = v335
	goto L83
L87:
	;
	v341 = v327
	goto L83
L88:
	;
	v351 = int32(1)
	v352 = v348
	goto L82
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v352
	goto L67
L90:
	;
	goto L91
L91:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+55)))
	if v361 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v386 = F_CreateDestReceiver(m, int32(9))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L11
	} else {
		goto L100
	}
L93:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v370
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+8)))
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[4]))
	v376 = F_tuplestore_begin_heap(m, v372, int32(0), v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L11
	} else {
		goto L98
	}
L94:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v364 != int32(1) {
		goto L93
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v352
	goto L92
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = v376
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v74)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v381
	if v379 != int32(1) {
		goto L67
	} else {
		goto L99
	}
L99:
	;
	goto L92
L100:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+20)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	m.T0[v394].(func(*base.Module, int32))(m, v392)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L101
	}
L101:
	;
	v407 = v386
	goto L66
L102:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	F_PushActiveSnapshot(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	v441 = int32(1)
	goto L65
L104:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	if v417 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	v420 = v418
	goto L107
L106:
	;
	v420 = int32(0)
	goto L107
L107:
	;
	v422 = F_CreateQueryDesc(m, v409, v411, v414, int32(0), v407, v416, v420, int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v422
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v425 != int32(6) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	F_ExecutorStart(m, v422, v428<<(uint(int32(5))%32)&int32(32))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L11
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v327
	v441 = v325
	goto L65
L112:
	;
	goto L111
L113:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v518 != int32(2) {
		goto L62
	} else {
		goto L136
	}
L114:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v74)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(2)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	if v487 != int32(6) {
		goto L125
	} else {
		goto L126
	}
L115:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+36))
	v455 = int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v448)+28))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	F_ProcessUtility(m, v452, v454, v455, v455, v457, v458, v459, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L11
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v465 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	F_ExecutorRun(m, v448, int32(1), v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L11
	} else {
		goto L119
	}
L118:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v480 = v463
	goto L114
L119:
	;
	if v465 == int64(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v480 = v470
	goto L114
L121:
	;
	goto L122
L122:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+44))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v472)+112))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v444
	if v473 == int64(0) {
		v480 = v471
		goto L114
	} else {
		goto L123
	}
L123:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+55)))
	if v479 != 0 {
		goto L113
	} else {
		goto L124
	}
L124:
	;
	v480 = v471
	goto L114
L125:
	;
	F_ExecutorFinish(m, v480)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L11
	} else {
		goto L128
	}
L126:
	;
	v496 = v480
	goto L127
L127:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+20))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	m.T0[v498].(func(*base.Module, int32))(m, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L130
	}
L128:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_ExecutorEnd(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v496 = v495
	goto L127
L130:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_FreeQueryDesc(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v444
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+9)))
	if v508 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v74)+68))
	F_MemoryContextDelete(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L11
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+68)) = int32(0)
	goto L113
L135:
	;
	goto L134
L136:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v521 != 0 {
		v291 = v521
		v295 = v441
		goto L63
	} else {
		goto L137
	}
L137:
	;
	goto L64
L138:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	v524 = F_init_execution_state(m, v74)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	if v524 == int32(0) {
		goto L55
	} else {
		goto L143
	}
L143:
	;
	goto L144
L144:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v544 != 0 {
		v274 = v544
		goto L60
	} else {
		goto L146
	}
L145:
	;
	goto L55
L146:
	;
	v545 = F_init_execution_state(m, v74)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L147
	}
L147:
	;
	if v545 != 0 {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	goto L145
L149:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v551 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L151
L151:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+56)))
	if v567 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L152:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+16))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+4)))
	if v553&int32(2) == int32(0) {
		v695 = v291
		v696 = v552
		v697 = v547
		v699 = v441
		goto L52
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v561)
	v563 = int64(0)
	if v441 != 0 {
		v745 = v291
		v758 = v563
		goto L51
	} else {
		goto L157
	}
L155:
	;
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v558)
	v560 = int64(0)
	if v441 != 0 {
		v745 = v291
		v758 = v560
		goto L51
	} else {
		goto L156
	}
L156:
	;
	v810 = v560
	goto L48
L157:
	;
	v810 = v563
	goto L48
L158:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	m.T0[v596].(func(*base.Module, int32))(m, v566)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L11
	} else {
		goto L169
	}
L159:
	;
	v570 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v570)
	v572 = F_ExecFetchSlotHeapTupleDatum(m, v566)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L11
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v574 = int32(*(*int16)(unsafe.Add(mBase, uint32(v566)+6)))
	if v574 <= int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v594 = v572
	goto L158
L163:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+16))
	m.T0[v579].(func(*base.Module, int32, int32))(m, v566, int32(1))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L11
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v566)+20))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v583)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v566)+16))
	v586 = *(*int64)(unsafe.Add(mBase, uint32(v585)))
	if v583 != 0 {
		v594 = v586
		goto L158
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+54)))
	v589 = int32(*(*int16)(unsafe.Add(mBase, uint32(v587)+52)))
	v590 = F_datumCopy(m, v586, v588, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L11
	} else {
		goto L168
	}
L168:
	;
	v594 = v590
	goto L158
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564)+20)) = int32(1)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	if v601 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	F_RegisterExprContextCallback(m, v604, int32(731), base.I64_extend_i32_u(v74))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L11
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v441 == int32(0) {
		v810 = v594
		goto L48
	} else {
		goto L175
	}
L173:
	;
	v609 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)) = uint8(v609)
	if v441 != 0 {
		v745 = v291
		v758 = v594
		goto L51
	} else {
		goto L174
	}
L174:
	;
	v810 = v594
	goto L48
L175:
	;
	v745 = v291
	v758 = v594
	goto L51
L176:
	;
	if v613 != 0 {
		v254 = v251
		goto L53
	} else {
		goto L177
	}
L177:
	;
	goto L55
L178:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+7)))
	if v634 == int32(1) {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v676 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633)+20)) = int32(2)
	v639 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v639)
	v641 = int64(0)
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	if v642 != v639 {
		v792 = v641
		goto L49
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633)+16)) = int32(2)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v633)+24)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v74)+16)) = int32(0)
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v658 != 0 {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	F_UnregisterExprContextCallback(m, v645, int32(731), base.I64_extend_i32_u(v74))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	v650 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)) = uint8(v650)
	v792 = v641
	goto L49
L186:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+8))
	v660 = F_CreateTupleDescCopy(m, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L11
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v663 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v663)
	v665 = int64(0)
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	if v666 != v663 {
		v792 = v665
		goto L49
	} else {
		goto L190
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v633)+28)) = v660
	goto L188
L190:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	F_UnregisterExprContextCallback(m, v669, int32(731), base.I64_extend_i32_u(v74))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L11
	} else {
		goto L191
	}
L191:
	;
	v674 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)) = uint8(v674)
	v792 = v665
	goto L49
L192:
	;
	v679 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v679)
	v792 = int64(0)
	goto L49
L193:
	;
	goto L194
L194:
	;
	v682 = int32(0)
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v676)+16))
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v684)+4)))
	if v685&int32(2) == v682 {
		v695 = v682
		v696 = v684
		v697 = v631
		v699 = v682
		goto L52
	} else {
		goto L195
	}
L195:
	;
	v690 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v690)
	v792 = int64(0)
	goto L49
L196:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+12))
	m.T0[v738].(func(*base.Module, int32))(m, v696)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L11
	} else {
		goto L207
	}
L197:
	;
	v712 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v712)
	v714 = F_ExecFetchSlotHeapTupleDatum(m, v696)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L11
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v716 = int32(*(*int16)(unsafe.Add(mBase, uint32(v696)+6)))
	if v716 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v736 = v714
	goto L196
L201:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+16))
	m.T0[v721].(func(*base.Module, int32, int32))(m, v696, int32(1))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L11
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v696)+20))
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v725)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v696)+16))
	v728 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
	if v725 != 0 {
		v736 = v728
		goto L196
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+54)))
	v731 = int32(*(*int16)(unsafe.Add(mBase, uint32(v729)+52)))
	v732 = F_datumCopy(m, v728, v730, v731)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L11
	} else {
		goto L206
	}
L206:
	;
	v736 = v732
	goto L196
L207:
	;
	if v699 == int32(0) {
		v763 = v695
		v776 = v736
		goto L50
	} else {
		goto L208
	}
L208:
	;
	v745 = v695
	v758 = v736
	goto L51
L209:
	;
	v763 = v745
	v776 = v758
	goto L50
L210:
	;
	v792 = v776
	goto L49
L211:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L11
	} else {
		goto L212
	}
L212:
	;
	F_errmsg(m, int32(_a_F_fmgr_sql_6), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L11
	} else {
		goto L213
	}
L213:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_7), int32(1605), int32(_a_F_fmgr_sql_8))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L11
	} else {
		goto L214
	}
L214:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
