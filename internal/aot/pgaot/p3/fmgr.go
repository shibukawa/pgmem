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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_CheckFunctionValidatorAccess(m, v12, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L34
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L31
	}
L3:
	;
	return int32(0)
L4:
	;
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = F_SearchSysCache1(m, int32(47), v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	m.G0 = v9 + int32(32)
	return int32(0)
L8:
	;
	if v19 == int32(0) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v25 = F_SysCacheGetAttrNotNull(m, int32(47), v19, int32(26))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v27 = F_text_to_cstring(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[450]))
	if v30 <= int32(0) {
		v83 = v2
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v83 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L13:
	;
	v34 = v2
	goto L14
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34<<(uint(int32(4))%32))+uint32(_consts[451])))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v47 == int32(0) {
		v66 = v46
		v67 = v47
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v73 = v34 << (uint(int32(4)) % 32)
	if v73+int32(1806176) == int32(0) {
		v83 = v2
		goto L12
	} else {
		goto L28
	}
L16:
	;
	if v67-v66 != 0 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	goto L16
L18:
	;
	if v46 != v47 {
		v66 = v46
		v67 = v47
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = v27
	v52 = v43
	goto L20
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v55
		v67 = v56
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v66 = v55
	v67 = v56
	goto L17
L22:
	;
	v59 = int32(1)
	if v55 == v56 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v70 = v34 + int32(1)
	if v70 != v30 {
		v34 = v70
		goto L14
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L15
L27:
	;
	v83 = v2
	goto L12
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)+uint32(_consts[452])))
	v83 = v78
	goto L12
L29:
	;
	F_ReleaseCatCache(m, v19)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	goto L7
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13
	F_errmsg_internal(m, int32(48455), v9)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(525184), int32(763), int32(220776))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v27
	F_errmsg(m, int32(757451), v9+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(525184), int32(772), int32(220776))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
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
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
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
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
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
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int64
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int64
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v827 int32
	_ = v827
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+11)))
	if v23 == v21 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	v840 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v840)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v843
	m.G0 = v19 + int32(16)
	return v827
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = int32(0)
	v827 = v809
	goto L1
L3:
	;
	v804 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v804)
	v809 = v802
	goto L2
L4:
	;
	if v788 != 0 {
		v827 = v789
		goto L1
	} else {
		goto L215
	}
L5:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L17
	} else {
		goto L214
	}
L6:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+16))
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+56)))
	if v721 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L17
	} else {
		goto L192
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v26 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v49 = v2
	v50 = v21
	v51 = v2
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v52 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v29 != int32(383) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v33 = int32(3)
	if v32&v33 != v33 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v49 = int32(base.Ui32(v32&int32(4)) >> (uint(int32(2)) % 32))
	v50 = base.B2i32(v32&int32(8) == int32(0))
	v51 = v42
	goto L10
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	v57 = F_MemoryContextAllocZero(m, v55, int32(84))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v75 = v52
	goto L16
L16:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	if v77 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	return int32(0)
L18:
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
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v57
	v75 = v57
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v51
	v237 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v237)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)) = uint8(v49)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(684)
	v243 = int32(4554984)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v19 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v244
	v251 = v75 + int32(44)
	v254 = v251
	goto L58
L21:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v97 = F_cached_function_compile(m, l0, v91, int32(682), int32(683), int32(88), int32(1), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L17
	} else {
		goto L26
	}
L22:
	;
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+44)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v80)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v75)+44))
	if v90 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v97 != v99 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	if v99 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(0) < v112 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = v101 - int64(1)
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v97
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v97)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+24)) = v106 + int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+24)) = int32(0)
	goto L29
L33:
	;
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+7)) = uint8(v213)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+5)) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+56)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v75)+40)) = int64(0)
	goto L20
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+44))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+40))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	if v119 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = int32(0)
	goto L33
L37:
	;
	v122 = int32(4562080)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v125
	v127 = F_makeParamList(m, v112)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L40
	}
L38:
	;
	v132 = v119
	goto L39
L39:
	;
	v141 = int32(0)
	goto L41
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v127
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v123
	v132 = v127
	goto L39
L41:
	;
	v157 = v132 + int32(32) + v141*int32(12)
	v160 = l0 + int32(20) + v141<<(uint(int32(3))%32)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)) = uint8(v161)
	if v161 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L33
L43:
	;
	v184 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+6)) = uint16(v184)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v183
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v118+v141<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v190
	v193 = v141 + v184
	if v193 != v112 {
		v141 = v193
		goto L41
	} else {
		goto L53
	}
L44:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v173 != int32(1) {
		v182 = v172
		goto L50
	} else {
		goto L51
	}
L45:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116+v141<<(uint(int32(1))%32)))))
	if v168 == int32(65535) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v183 = v171
	goto L43
L48:
	;
	goto L47
L49:
	;
	v183 = v182
	goto L43
L50:
	;
	goto L49
L51:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	if v176 != int32(3) {
		v182 = v172
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)+2))
	v182 = v179 + int32(18)
	goto L50
L53:
	;
	goto L42
L54:
	;
	v696 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v696)
	v698 = int32(0)
	if v445&v696 != 0 {
		v770 = v291
		v771 = v698
		goto L5
	} else {
		goto L191
	}
L55:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651)+56)))
	if v664 == int32(1) {
		goto L179
	} else {
		goto L180
	}
L56:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v631 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L57:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+55)))
	if v616 != 0 {
		goto L6
	} else {
		goto L170
	}
L58:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v268 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574)+55)))
	if v575 == int32(0) {
		goto L56
	} else {
		goto L158
	}
L60:
	;
	goto L59
L61:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v269 == int32(2) {
		v254 = v268
		goto L58
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v556 = F_init_execution_state(m, v75)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L17
	} else {
		goto L156
	}
L64:
	;
	v274 = v268
	goto L65
L65:
	;
	v291 = v274
	v295 = int32(0)
	goto L67
L66:
	;
	goto L60
L67:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+57)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v307 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v445&int32(1) != 0 {
		goto L145
	} else {
		goto L146
	}
L69:
	;
	v449 = int32(4562080)
	v450 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	if v455 == int32(6) {
		goto L122
	} else {
		goto L123
	}
L70:
	;
	if v306&int32(1) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if (v306|v295)&int32(1) != 0 {
		goto L116
	} else {
		goto L117
	}
L73:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L17
	} else {
		goto L76
	}
L74:
	;
	v327 = v295
	goto L75
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v330 != int32(1) {
		goto L86
	} else {
		goto L87
	}
L76:
	;
	if v295&int32(1) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v327 = int32(1)
	goto L75
L78:
	;
	v320 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L17
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L17
	} else {
		goto L83
	}
L81:
	;
	F_PushActiveSnapshot(m, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L17
	} else {
		goto L82
	}
L82:
	;
	goto L77
L83:
	;
	goto L77
L84:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)) = uint8(v352)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+68)) = v354
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	if v357 == int32(1) {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v350 = F_AllocSetContextCreateInternal(m, v344, int32(258817), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L17
	} else {
		goto L90
	}
L86:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v340 != int32(6) {
		v352 = int32(0)
		v354 = v329
		goto L84
	} else {
		goto L89
	}
L87:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+55)))
	if v334 != int32(1) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v75)+60))
	v344 = v337
	goto L85
L89:
	;
	v344 = v329
	goto L85
L90:
	;
	v352 = int32(1)
	v354 = v350
	goto L84
L91:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+36))
	v407 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	goto L104
L92:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _consts[537]))
	v401 = v400
	goto L91
L93:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+55)))
	if v361 != int32(1) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v354
	goto L92
L96:
	;
	v385 = F_CreateDestReceiver(m, int32(9))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L17
	} else {
		goto L102
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v354
	goto L96
L98:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v364 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v366
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+8)))
	v371 = *(*int32)(unsafe.Add(mBase, _consts[538]))
	v372 = F_tuplestore_begin_heap(m, v368, int32(0), v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L17
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v372
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+8)))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v377
	if v375&int32(1) != 0 {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	goto L92
L102:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v385)+24)) = v389
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+8))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	m.T0[v393].(func(*base.Module, int32))(m, v391)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L17
	} else {
		goto L103
	}
L103:
	;
	v401 = v385
	goto L91
L104:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	if v411 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+28))
	v414 = v412
	goto L107
L106:
	;
	v414 = int32(0)
	goto L107
L107:
	;
	v416 = F_CreateQueryDesc(m, v403, v405, v408, int32(0), v401, v410, v414, int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L17
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v416
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	if v419 != int32(6) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	if v424 != 0 {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v329
	v445 = v327
	goto L69
L112:
	;
	v425 = int32(32)
	goto L114
L113:
	;
	v425 = int32(0)
	goto L114
L114:
	;
	F_ExecutorStart(m, v416, v425)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L17
	} else {
		goto L115
	}
L115:
	;
	goto L111
L116:
	;
	v445 = v306 ^ int32(1) | v295
	goto L69
L117:
	;
	goto L118
L118:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	F_PushActiveSnapshot(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	v445 = int32(1)
	goto L69
L120:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v525 != int32(2) {
		goto L57
	} else {
		goto L143
	}
L121:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = int32(2)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v487)))
	if v494 != int32(6) {
		goto L132
	} else {
		goto L133
	}
L122:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454)+4))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+36))
	v461 = int32(1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v454)+24))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v454)+28))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v454)+20))
	F_ProcessUtility(m, v458, v460, v461, v461, v463, v464, v465, int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L17
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v471 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v291)+9)))
	F_ExecutorRun(m, v454, int32(1), v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L17
	} else {
		goto L126
	}
L125:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v487 = v469
	goto L121
L126:
	;
	if base.I32_wrap_i64(v471) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v487 = v477
	goto L121
L128:
	;
	goto L129
L129:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+40))
	v480 = *(*int64)(unsafe.Add(mBase, uint32(v479)+112))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v450
	if v480 == int64(0) {
		v487 = v478
		goto L121
	} else {
		goto L130
	}
L130:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+55)))
	if v486 != 0 {
		goto L120
	} else {
		goto L131
	}
L131:
	;
	v487 = v478
	goto L121
L132:
	;
	F_ExecutorFinish(m, v487)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L17
	} else {
		goto L135
	}
L133:
	;
	v503 = v487
	goto L134
L134:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v503)+20))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)+12))
	m.T0[v505].(func(*base.Module, int32))(m, v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L17
	} else {
		goto L137
	}
L135:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_ExecutorEnd(m, v499)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	v503 = v502
	goto L134
L137:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v291)+16))
	F_FreeQueryDesc(m, v508)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L17
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v450
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+9)))
	if v515 == int32(1) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v75)+68))
	F_MemoryContextDelete(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L17
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+68)) = int32(0)
	goto L120
L142:
	;
	goto L141
L143:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v528 != 0 {
		v291 = v528
		v295 = v445
		goto L67
	} else {
		goto L144
	}
L144:
	;
	goto L68
L145:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L17
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v533 = F_init_execution_state(m, v75)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L17
	} else {
		goto L149
	}
L148:
	;
	goto L147
L149:
	;
	if v533 == int32(0) {
		goto L60
	} else {
		goto L150
	}
L150:
	;
	goto L151
L151:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v553 != 0 {
		v274 = v553
		goto L65
	} else {
		goto L153
	}
L152:
	;
	goto L66
L153:
	;
	v554 = F_init_execution_state(m, v75)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	if v554 != 0 {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	goto L152
L156:
	;
	if v556 != 0 {
		v254 = v251
		goto L58
	} else {
		goto L157
	}
L157:
	;
	goto L60
L158:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+7)))
	if v579 == int32(1) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+20)) = int32(2)
	v584 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v584)
	v586 = int32(0)
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v587 != v584 {
		v809 = v586
		goto L2
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+16)) = int32(2)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+24)) = v596
	v598 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v598
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v601 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	F_UnregisterExprContextCallback(m, v590, int32(685), v75)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L17
	} else {
		goto L163
	}
L163:
	;
	v802 = v586
	goto L3
L164:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v601)+8))
	v603 = F_CreateTupleDescCopy(m, v602)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L17
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v606)
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v608 != v606 {
		v809 = v598
		goto L2
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v578)+28)) = v603
	goto L166
L168:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	F_UnregisterExprContextCallback(m, v611, int32(685), v75)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L17
	} else {
		goto L169
	}
L169:
	;
	v802 = v598
	goto L3
L170:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v75)+24))
	if v617 == int32(0) {
		goto L54
	} else {
		goto L171
	}
L171:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v617)+16))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+4)))
	if v621&int32(2) == int32(0) {
		v650 = v291
		v651 = v615
		v652 = v620
		v654 = v445
		goto L55
	} else {
		goto L172
	}
L172:
	;
	v626 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v626)
	v628 = int32(0)
	if v445&v626 != 0 {
		v770 = v291
		v771 = v628
		goto L5
	} else {
		goto L173
	}
L173:
	;
	v827 = v628
	goto L1
L174:
	;
	v634 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v634)
	v809 = int32(0)
	goto L2
L175:
	;
	goto L176
L176:
	;
	v637 = int32(0)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v631)+16))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+4)))
	if v640&int32(2) == v637 {
		v650 = v637
		v651 = v574
		v652 = v639
		v654 = v637
		goto L55
	} else {
		goto L177
	}
L177:
	;
	v645 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v645)
	v809 = int32(0)
	goto L2
L178:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v652)+8))
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v690)+12))
	m.T0[v691].(func(*base.Module, int32))(m, v652)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L17
	} else {
		goto L189
	}
L179:
	;
	v667 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v667)
	v669 = F_ExecFetchSlotHeapTupleDatum(m, v652)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L17
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v671 = int32(*(*int16)(unsafe.Add(mBase, uint32(v652)+6)))
	if v671 <= int32(0) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v689 = v669
	goto L178
L183:
	;
	F_slot_getsomeattrs_int(m, v652, int32(1))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L17
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v652)+20))
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v678)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v652)+16))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	if v678 != 0 {
		v689 = v681
		goto L178
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682)+54)))
	v684 = int32(*(*int16)(unsafe.Add(mBase, uint32(v682)+52)))
	v685 = F_datumCopy(m, v681, v683, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L17
	} else {
		goto L188
	}
L188:
	;
	v689 = v685
	goto L178
L189:
	;
	if v654&int32(1) != 0 {
		v770 = v650
		v771 = v689
		goto L5
	} else {
		goto L190
	}
L190:
	;
	v788 = v650
	v789 = v689
	goto L4
L191:
	;
	v827 = v698
	goto L1
L192:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L17
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(113957), int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L17
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(517657), int32(1604), int32(315498))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L17
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
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)+12))
	m.T0[v748].(func(*base.Module, int32))(m, v720)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L17
	} else {
		goto L207
	}
L197:
	;
	v724 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v724)
	v726 = F_ExecFetchSlotHeapTupleDatum(m, v720)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L17
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v728 = int32(*(*int16)(unsafe.Add(mBase, uint32(v720)+6)))
	if v728 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v746 = v726
	goto L196
L201:
	;
	F_slot_getsomeattrs_int(m, v720, int32(1))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L17
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v720)+20))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v735)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v720)+16))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)))
	if v735 != 0 {
		v746 = v738
		goto L196
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739)+54)))
	v741 = int32(*(*int16)(unsafe.Add(mBase, uint32(v739)+52)))
	v742 = F_datumCopy(m, v738, v740, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L17
	} else {
		goto L206
	}
L206:
	;
	v746 = v742
	goto L196
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+20)) = int32(1)
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)))
	if v753 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	F_RegisterExprContextCallback(m, v756, int32(685), v75)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L17
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if v445&int32(1) == int32(0) {
		v827 = v746
		goto L1
	} else {
		goto L213
	}
L211:
	;
	v760 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+6)) = uint8(v760)
	if v445&v760 != 0 {
		v770 = v291
		v771 = v746
		goto L5
	} else {
		goto L212
	}
L212:
	;
	v827 = v746
	goto L1
L213:
	;
	v770 = v291
	v771 = v746
	goto L5
L214:
	;
	v788 = v770
	v789 = v771
	goto L4
L215:
	;
	v809 = v789
	goto L2
}
