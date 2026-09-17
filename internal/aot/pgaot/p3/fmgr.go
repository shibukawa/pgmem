package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fmgr_internal_validator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_CheckFunctionValidatorAccess(m, v13, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L3
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L29
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = F_SearchSysCache1(m, int32(47), v14)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	m.G0 = v10 + int32(32)
	return int32(0)
L8:
	;
	if v20 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v26 = F_SysCacheGetAttrNotNull(m, int32(47), v20, int32(26))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v28 = F_text_to_cstring(m, v26)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_internal_validator[0]))
	if v31 <= int32(0) {
		v82 = v2
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v82 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L13:
	;
	v37 = v2
	goto L14
L14:
	;
	v42 = v37 << (uint(int32(4)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_fmgr_internal_validator[1])))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.B2i32(v48 == int32(0))|base.B2i32(v48 != v51) != 0 {
		v69 = v48
		v70 = v51
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_fmgr_internal_validator[2])))
	v82 = v75
	goto L12
L16:
	;
	if v69-v70 != 0 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	v54 = v28
	v55 = v45
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v59
		v70 = v58
		goto L17
	} else {
		goto L21
	}
L20:
	;
	v69 = v59
	v70 = v58
	goto L17
L21:
	;
	v62 = int32(1)
	if v59 == v58 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v73 = v37 + int32(1)
	if v31 != v73 {
		v37 = v73
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
	v82 = v2
	goto L12
L27:
	;
	F_ReleaseCatCache(m, v20)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
	F_errmsg_internal(m, int32(_a_F_fmgr_internal_validator_0), v10)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_fmgr_internal_validator_1), int32(763), int32(_a_F_fmgr_internal_validator_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v28
	F_errmsg(m, int32(_a_F_fmgr_internal_validator_3), v10+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_fmgr_internal_validator_1), int32(772), int32(_a_F_fmgr_internal_validator_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
func F_fmgr_sql(m *base.Module, l0 int32) int32 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v106 int64
	_ = v106
	var v112 int32
	_ = v112
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
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
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
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
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
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
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v802 int32
	_ = v802
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
	if v23 == v21 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	v814 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v814)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0])) = v817
	m.G0 = v19 + int32(16)
	return v802
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = int32(0)
	v802 = v784
	goto L1
L3:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v771 != int32(1) {
		v784 = v770
		goto L2
	} else {
		goto L212
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L14
	} else {
		goto L208
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v26 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v48 = v21
	v50 = v2
	v51 = v2
	goto L7
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v52 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(383) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v33 = int32(3)
	if v32&v33 != v33 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v48 = base.B2i32(v32&int32(8) == int32(0))
	v50 = v38
	v51 = int32(base.Ui32(v32&int32(4)) >> (uint(int32(2)) % 32))
	goto L7
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v57 = F_MemoryContextAllocZero(m, v55, int32(84))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v75 = v52
	goto L13
L13:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	if v77 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	return int32(0)
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+72)) = int32(681)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+60)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v57)+76)) = v57
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v68 = v57 + int32(72)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v69
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+4)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+40)) = v68
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v57
	v75 = v57
	goto L13
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v50
	v237 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v237)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(684)
	v243 = int32(_a_F_fmgr_sql_0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[0])) = v19 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v244
	v251 = v75 + int32(44)
	v254 = v251
	goto L54
L18:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v97 = F_cached_function_compile(m, l0, v91, int32(682), int32(683), int32(88), int32(1), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v80)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+44))
	if v90 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v97 != v99 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v99 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v112 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = v101 - int64(1)
	goto L29
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v97
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+24)) = v106 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+24)) = int32(0)
	goto L26
L30:
	;
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+7)) = uint8(v213)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+5)) = uint8(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+56)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v75)+40)) = int64(0)
	goto L17
L31:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+44))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+40))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v119 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = int32(0)
	goto L30
L34:
	;
	v122 = int32(_a_F_fmgr_sql_1)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v125
	v127 = F_makeParamList(m, v112)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L37
	}
L35:
	;
	v132 = v119
	goto L36
L36:
	;
	v141 = int32(0)
	goto L38
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v127
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v123
	v132 = v127
	goto L36
L38:
	;
	v157 = v132 + int32(32) + v141*int32(12)
	v160 = l0 + int32(20) + v141<<(uint(int32(3))%32)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)) = uint8(v161)
	if v161 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L30
L40:
	;
	v184 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+6)) = uint16(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v183
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v118+v141<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v190
	v193 = v141 + v184
	if v193 != v112 {
		v141 = v193
		goto L38
	} else {
		goto L50
	}
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v173 != int32(1) {
		v182 = v172
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116+v141<<(uint(int32(1))%32)))))
	if v168 == int32(_a_F_fmgr_sql_2) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v183 = v171
	goto L40
L45:
	;
	goto L44
L46:
	;
	v183 = v182
	goto L40
L47:
	;
	goto L46
L48:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	if v176 != int32(3) {
		v182 = v172
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)+2))
	v182 = v179 + int32(18)
	goto L47
L50:
	;
	goto L39
L51:
	;
	if v738 != 0 {
		v802 = v740
		goto L1
	} else {
		goto L207
	}
L52:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L14
	} else {
		goto L206
	}
L53:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674)+56)))
	if v686 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L54:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v268 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628)+55)))
	if v629 != 0 {
		goto L179
	} else {
		goto L180
	}
L56:
	;
	goto L55
L57:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v269 == int32(2) {
		v254 = v268
		goto L54
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v610 = F_init_execution_state(m, v75)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L14
	} else {
		goto L177
	}
L60:
	;
	v274 = v268
	goto L61
L61:
	;
	v291 = v274
	v295 = int32(0)
	goto L64
L62:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+55)))
	if v548 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L63:
	;
	goto L62
L64:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+57)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v307 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	if v441 != 0 {
		goto L139
	} else {
		goto L140
	}
L66:
	;
	v443 = int32(_a_F_fmgr_sql_1)
	v444 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v449 == int32(6) {
		goto L116
	} else {
		goto L117
	}
L67:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)+36))
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[2]))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v413)))
	goto L105
L68:
	;
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[3]))
	v407 = v406
	goto L67
L69:
	;
	if v306&int32(1) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if (v306|v295)&int32(1) != 0 {
		v441 = v295
		goto L66
	} else {
		goto L103
	}
L72:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L14
	} else {
		goto L75
	}
L73:
	;
	v325 = v295
	goto L74
L74:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1]))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v328 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L75:
	;
	if v295 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v325 = int32(1)
	goto L74
L77:
	;
	v318 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L14
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L14
	} else {
		goto L82
	}
L80:
	;
	F_PushActiveSnapshot(m, v318)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L14
	} else {
		goto L81
	}
L81:
	;
	goto L76
L82:
	;
	goto L76
L83:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)) = uint8(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+68)) = v352
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	if v355 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	v348 = F_AllocSetContextCreateInternal(m, v341, int32(_a_F_fmgr_sql_3), int32(0), int32(_a_F_fmgr_sql_4), int32(_a_F_fmgr_sql_5))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L14
	} else {
		goto L89
	}
L85:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v338 != int32(6) {
		v351 = int32(0)
		v352 = v327
		goto L83
	} else {
		goto L88
	}
L86:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+55)))
	if v332 != int32(1) {
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v341 = v335
	goto L84
L88:
	;
	v341 = v327
	goto L84
L89:
	;
	v351 = int32(1)
	v352 = v348
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v352
	goto L68
L91:
	;
	goto L92
L92:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+55)))
	if v361 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v386 = F_CreateDestReceiver(m, int32(9))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L101
	}
L94:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v370
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)))
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[4]))
	v376 = F_tuplestore_begin_heap(m, v372, int32(0), v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L14
	} else {
		goto L99
	}
L95:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v364 != int32(1) {
		goto L94
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v352
	goto L93
L98:
	;
	goto L97
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v376
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v381
	if v379 != int32(1) {
		goto L68
	} else {
		goto L100
	}
L100:
	;
	goto L93
L101:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+20)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v390)+16))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+8))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	m.T0[v394].(func(*base.Module, int32))(m, v392)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	v407 = v386
	goto L67
L103:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+12))
	F_PushActiveSnapshot(m, v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	v441 = int32(1)
	goto L66
L105:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	if v417 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)+28))
	v420 = v418
	goto L108
L107:
	;
	v420 = int32(0)
	goto L108
L108:
	;
	v422 = F_CreateQueryDesc(m, v409, v411, v414, int32(0), v407, v416, v420, int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L14
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v422
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v425 != int32(6) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	F_ExecutorStart(m, v422, v428<<(uint(int32(5))%32)&int32(32))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L14
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v327
	v441 = v325
	goto L66
L113:
	;
	goto L112
L114:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v518 != int32(2) {
		goto L63
	} else {
		goto L137
	}
L115:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v483
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(2)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	if v487 != int32(6) {
		goto L126
	} else {
		goto L127
	}
L116:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+36))
	v455 = int32(1)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v448)+28))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	F_ProcessUtility(m, v452, v454, v455, v455, v457, v458, v459, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L14
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v465 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	F_ExecutorRun(m, v448, int32(1), v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L14
	} else {
		goto L120
	}
L119:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v480 = v463
	goto L115
L120:
	;
	if v465 == int64(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v480 = v470
	goto L115
L122:
	;
	goto L123
L123:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)+40))
	v473 = *(*int64)(unsafe.Add(mBase, uint32(v472)+112))
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v444
	if v473 == int64(0) {
		v480 = v471
		goto L115
	} else {
		goto L124
	}
L124:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478)+55)))
	if v479 != 0 {
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v480 = v471
	goto L115
L126:
	;
	F_ExecutorFinish(m, v480)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L14
	} else {
		goto L129
	}
L127:
	;
	v496 = v480
	goto L128
L128:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+20))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	m.T0[v498].(func(*base.Module, int32))(m, v497)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L14
	} else {
		goto L131
	}
L129:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_ExecutorEnd(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v496 = v495
	goto L128
L131:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_FreeQueryDesc(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_fmgr_sql[1])) = v444
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	if v508 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	F_MemoryContextDelete(m, v511)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L14
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+68)) = int32(0)
	goto L114
L136:
	;
	goto L135
L137:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v521 != 0 {
		v291 = v521
		v295 = v441
		goto L64
	} else {
		goto L138
	}
L138:
	;
	goto L65
L139:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L14
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v524 = F_init_execution_state(m, v75)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L14
	} else {
		goto L143
	}
L142:
	;
	goto L141
L143:
	;
	if v524 == int32(0) {
		goto L56
	} else {
		goto L144
	}
L144:
	;
	goto L145
L145:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v544 != 0 {
		v274 = v544
		goto L61
	} else {
		goto L147
	}
L146:
	;
	goto L56
L147:
	;
	v545 = F_init_execution_state(m, v75)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L14
	} else {
		goto L148
	}
L148:
	;
	if v545 != 0 {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v551 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+56)))
	if v567 == int32(1) {
		goto L160
	} else {
		goto L161
	}
L153:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)+16))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+4)))
	if v553&int32(2) == int32(0) {
		v672 = v291
		v673 = v552
		v674 = v547
		v676 = v441
		goto L53
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v561)
	v563 = int32(0)
	if v441 != 0 {
		v720 = v291
		v722 = v563
		goto L52
	} else {
		goto L158
	}
L156:
	;
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v558)
	v560 = int32(0)
	if v441 != 0 {
		v720 = v291
		v722 = v560
		goto L52
	} else {
		goto L157
	}
L157:
	;
	v802 = v560
	goto L1
L158:
	;
	v802 = v563
	goto L1
L159:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v566)+8))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v593)+12))
	m.T0[v594].(func(*base.Module, int32))(m, v566)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L14
	} else {
		goto L170
	}
L160:
	;
	v570 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v570)
	v572 = F_ExecFetchSlotHeapTupleDatum(m, v566)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L14
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v574 = int32(*(*int16)(unsafe.Add(mBase, uint32(v566)+6)))
	if v574 <= int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v592 = v572
	goto L159
L164:
	;
	F_slot_getsomeattrs_int(m, v566, int32(1))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L14
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v566)+20))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v581)
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v566)+16))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	if v581 != 0 {
		v592 = v584
		goto L159
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+54)))
	v587 = int32(*(*int16)(unsafe.Add(mBase, uint32(v585)+52)))
	v588 = F_datumCopy(m, v584, v586, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L14
	} else {
		goto L169
	}
L169:
	;
	v592 = v588
	goto L159
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v564)+20)) = int32(1)
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v599 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v564)+4))
	F_RegisterExprContextCallback(m, v602, int32(685), v75)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L14
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	if v441 == int32(0) {
		v802 = v592
		goto L1
	} else {
		goto L176
	}
L174:
	;
	v606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v606)
	if v441 != 0 {
		v720 = v291
		v722 = v592
		goto L52
	} else {
		goto L175
	}
L175:
	;
	v802 = v592
	goto L1
L176:
	;
	v720 = v291
	v722 = v592
	goto L52
L177:
	;
	if v610 != 0 {
		v254 = v251
		goto L54
	} else {
		goto L178
	}
L178:
	;
	goto L56
L179:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+7)))
	if v631 == int32(1) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v653 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+20)) = int32(2)
	v636 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v636)
	v770 = int32(0)
	goto L3
L183:
	;
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+16)) = int32(2)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v630)+24)) = v641
	v643 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v643
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v646 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+8))
	v648 = F_CreateTupleDescCopy(m, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L14
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v651 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v651)
	v770 = v643
	goto L3
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+28)) = v648
	goto L187
L189:
	;
	v656 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v656)
	v784 = int32(0)
	goto L2
L190:
	;
	goto L191
L191:
	;
	v659 = int32(0)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v653)+16))
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v661)+4)))
	if v662&int32(2) == v659 {
		v672 = v659
		v673 = v661
		v674 = v628
		v676 = v659
		goto L53
	} else {
		goto L192
	}
L192:
	;
	v667 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v667)
	v784 = int32(0)
	goto L2
L193:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v673)+8))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	m.T0[v713].(func(*base.Module, int32))(m, v673)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L14
	} else {
		goto L204
	}
L194:
	;
	v689 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v689)
	v691 = F_ExecFetchSlotHeapTupleDatum(m, v673)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L14
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v693 = int32(*(*int16)(unsafe.Add(mBase, uint32(v673)+6)))
	if v693 <= int32(0) {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v711 = v691
	goto L193
L198:
	;
	F_slot_getsomeattrs_int(m, v673, int32(1))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L14
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v673)+20))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v700)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v673)+16))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v702)))
	if v700 != 0 {
		v711 = v703
		goto L193
	} else {
		goto L202
	}
L201:
	;
	goto L200
L202:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v704)+54)))
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(v704)+52)))
	v707 = F_datumCopy(m, v703, v705, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L14
	} else {
		goto L203
	}
L203:
	;
	v711 = v707
	goto L193
L204:
	;
	if v676 == int32(0) {
		v738 = v672
		v740 = v711
		goto L51
	} else {
		goto L205
	}
L205:
	;
	v720 = v672
	v722 = v711
	goto L52
L206:
	;
	v738 = v720
	v740 = v722
	goto L51
L207:
	;
	v784 = v740
	goto L2
L208:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L14
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_fmgr_sql_6), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L14
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_fmgr_sql_7), int32(1604), int32(_a_F_fmgr_sql_8))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L14
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	F_UnregisterExprContextCallback(m, v774, int32(685), v75)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L14
	} else {
		goto L213
	}
L213:
	;
	v778 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v778)
	v784 = v770
	goto L2
}
