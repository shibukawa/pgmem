package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexPartitions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = F_get_rel_name(m, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = F_get_namespace_name(m, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = F_pstrdup(m, v16)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v22
	v25 = F_pstrdup(m, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)) = uint8(v14)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(558)
	v31 = int32(0)
	v32 = int32(4418184)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v12 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v12 + int32(8)
	if v14 == int32(112) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = int32(508716)
	goto L10
L9:
	;
	v46 = int32(478318)
	goto L10
L10:
	;
	F_PreventInTransactionBlock(m, l3, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int32)(unsafe.Add(mBase, _consts[84])) = v50
	v53 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	v58 = F_AllocSetContextCreateInternal(m, v53, int32(25480), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v62 = F_find_all_inheritors(m, l1, int32(5), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	F_ReindexMultipleInternal(m, l0, v109, l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L26
	}
L14:
	;
	if v62 == int32(0) {
		v109 = v31
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v67 <= v66 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v109 = v31
	goto L13
L17:
	;
	goto L18
L18:
	;
	v71 = v66
	v75 = v31
	goto L19
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v71<<(uint(int32(2))%32))))
	v84 = F_get_rel_relkind(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	v109 = v98
	goto L13
L21:
	;
	v101 = v71 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v101 < v102 {
		v71 = v101
		v75 = v98
		goto L19
	} else {
		goto L25
	}
L22:
	;
	v90 = int32(4425280)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v58
	v94 = F_lappend_oid(m, v75, v83)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	switch v84&int32(255) - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L22
	default:
		v98 = v75
		goto L21
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v91
	v98 = v94
	goto L21
L25:
	;
	goto L20
L26:
	;
	F_MemoryContextDelete(m, v58)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	m.G0 = v12 + int32(32)
	return
}
func F_reindex_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v744 int32
	_ = v744
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(336)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)) = uint8(v6)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_getrusage(m, v15+int32(168))
	mBase = m.M
	F___gettimeofday(m, v15+int32(152))
	mBase = m.M
	goto L1
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v28 = F_SearchSysCache1(m, int32(34), l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L10
	} else {
		goto L216
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L10
	} else {
		goto L213
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L10
	} else {
		goto L210
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L10
	} else {
		goto L206
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L10
	} else {
		goto L202
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L10
	} else {
		goto L198
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L10
	} else {
		goto L194
	}
L9:
	;
	m.G0 = v15 + int32(336)
	return
L10:
	;
	return
L11:
	;
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v26&int32(4) != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+22)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48)+4))
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	F_errmsg_internal(m, int32(37172), v15)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(462298), int32(3594), int32(248392))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if v50 == int32(0) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v55&int32(4) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v64 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L22:
	;
	v59 = F_try_table_open(m, v50, int32(5))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v62 = F_table_open(m, v50, int32(5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	v64 = v59
	goto L21
L26:
	;
	v64 = v62
	goto L21
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(332)))) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _consts[277]))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(328)))) = v75
	goto L28
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _consts[277])) = v79 | int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v78
	goto L29
L29:
	;
	v87 = int32(4423320)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	v91 = v89 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[278])) = v91
	goto L30
L30:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v96 = v19 & int32(2)
	if v96 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = int64(25769803776)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+136)) = base.I64_extend_i32_u(l1)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = int64(3)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v106 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v342&int32(4) != 0 {
		goto L65
	} else {
		goto L66
	}
L35:
	;
	v169 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v176 == v169 {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	goto L35
L37:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	if v110 != int32(1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v113 = int32(4419940)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v116 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v115 + v116
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v119 + v116
	*(*int32)(unsafe.Add(mBase, uint32(v106)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+224)) = v50
	v126 = v106 + int32(232)
	if v126&int32(3) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v153 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v152 + v153
	v156 = int32(4419940)
	v158 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v158 - v153
	goto L36
L40:
	;
	v132 = v106 + int32(392)
	if base.Ui32(v132) <= base.Ui32(v126) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v149 = F___memset(m, v126, int32(0), int32(160))
	mBase = m.M
	goto L39
L43:
	;
	v136 = v106 + int32(236)
	if base.Ui32(v136) < base.Ui32(v132) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v138 = v132
	goto L46
L45:
	;
	v138 = v136
	goto L46
L46:
	;
	v146 = F___memset(m, v126, int32(0), (v138-v106-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L39
L47:
	;
	goto L34
L48:
	;
	goto L47
L49:
	;
	goto L50
L50:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	if v182&int32(1) == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v187 = int32(4419940)
	v189 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v190 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v189 + v190
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v193 + v190
	goto L53
L52:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v324 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v323 + v324
	v327 = int32(4419940)
	v329 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v329 - v324
	goto L48
L53:
	;
	goto L55
L55:
	;
	goto L56
L56:
	;
	goto L60
L60:
	;
	v288 = int32(0)
	v291 = v169
	goto L61
L61:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)+v291<<(uint(int32(2))%32))))
	v301 = int32(3)
	v307 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(128)+v291<<(uint(v301)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v176+int32(232)+v300<<(uint(v301)%32)))) = v307
	v309 = int32(1)
	v312 = v288 + v309
	if v312 != int32(2) {
		v288 = v312
		v291 = v291 + v309
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L52
L63:
	;
	goto L62
L64:
	;
	if v387 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L65:
	;
	v345 = m.G0
	v347 = v345 - int32(16)
	m.G0 = v347
	v350 = F_try_relation_open(m, l1, int32(8))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v383 = F_index_open(m, l1, int32(8))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L10
	} else {
		goto L76
	}
L68:
	;
	m.G0 = v347 + int32(16)
	v387 = v350
	goto L64
L69:
	;
	if v350 == int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v350)+48))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+119)))
	if v355|int32(32) == int32(105) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v350)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v367 + int32(4)
	F_errmsg(m, int32(25895), v347)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(466631), int32(204), int32(399594))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v387 = v383
	goto L64
L77:
	;
	F_AtEOXact_GUC(m, int32(0), v91)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L10
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v96 != 0 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _consts[277])) = v394
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v393
	goto L81
L81:
	;
	F_sequence_close(m, v64, int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L10
	} else {
		goto L82
	}
L82:
	;
	goto L9
L83:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	v404 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v403)+84)))
	v407 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v407 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	goto L85
L85:
	;
	if l0 != 0 {
		goto L90
	} else {
		goto L91
	}
L86:
	;
	goto L85
L87:
	;
	goto L86
L88:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	if v411 != int32(1) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v414 = int32(4419940)
	v416 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v417 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v416 + v417
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v420 + v417
	*(*int64)(unsafe.Add(mBase, uint32(v407+int32(64))+232)) = v404
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v428 + v417
	v434 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v434 - v417
	goto L87
L90:
	;
	v438 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v438
	v443 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(1259)
	v449 = *(*int64)(unsafe.Add(mBase, _consts[280]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v15)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v451
	F_EventTriggerCollectSimpleCommand(m, v15+int32(112), v15+int32(96), l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L10
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+119)))
	if v460 == int32(73) {
		goto L8
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+118)))
	if v463 == int32(116) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+24)))
	if v466 == int32(0) {
		goto L7
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v459)+68))
	if v469 != int32(99) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L97
L99:
	;
	if v475 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v474 = F_isTempToastNamespace(m, v469)
	mBase = m.M
	v475 = v474
	goto L102
L101:
	;
	v475 = int32(1)
	goto L102
L102:
	;
	goto L99
L103:
	;
	v476 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L10
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v480 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	if v476 == int32(0) {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	F_TransferPredicateLocksToHeapRelation(m, v387)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L10
	} else {
		goto L131
	}
L109:
	;
	F_CheckTableNotInUse(m, v387, int32(478318))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L130
	}
L110:
	;
	v484 = int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v387)+56))
	if base.Ui32(v485) < base.Ui32(int32(12000)) {
		v494 = v484
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v494 != 0 {
		goto L5
	} else {
		goto L115
	}
L112:
	;
	goto L111
L113:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v488)+68))
	if v489 == int32(99) {
		v494 = v484
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v492 = F_isTempToastNamespace(m, v489)
	mBase = m.M
	v494 = v492
	goto L112
L115:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v495 == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v498 = F_CheckRelationTableSpaceMove(m, v387, v495)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L10
	} else {
		goto L117
	}
L117:
	;
	F_CheckTableNotInUse(m, v387, int32(478318))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	if v498 == int32(0) {
		goto L108
	} else {
		goto L119
	}
L119:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	F_SetRelationTableSpace(m, v387, v505, int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	F_RelationDropStorage(m, v387)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L10
	} else {
		goto L121
	}
L121:
	;
	v512 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v387)+36)) = v512
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v387)+40))
	if v514 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L10
	} else {
		goto L129
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+40)) = v512
	goto L125
L124:
	;
	goto L125
L125:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	if v519 <= int32(31) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v387)+56))
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v519 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v519<<(uint(int32(2))%32))+uint32(_consts[284]))) = v522
	goto L122
L127:
	;
	goto L128
L128:
	;
	v533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[285])) = uint8(v533)
	goto L122
L129:
	;
	goto L108
L130:
	;
	goto L108
L131:
	;
	v545 = F_BuildIndexInfo(m, v387)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L10
	} else {
		goto L132
	}
L132:
	;
	if l2 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+116)))
	if v547 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	goto L135
L135:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _consts[281]))
	if v562 != 0 {
		goto L4
	} else {
		goto L141
	}
L136:
	;
	v555 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v545)+100)) = v555
	*(*int64)(unsafe.Add(mBase, uint32(v545)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v545)+116)) = uint8(v555)
	goto L135
L137:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v545)+92))
	if v550 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v553 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)) = uint8(v553)
	goto L136
L140:
	;
	goto L139
L141:
	;
	*(*int32)(unsafe.Add(mBase, _consts[118])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v50
	v571 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+72))
	if v572 != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if v574&int32(1) != 0 {
		goto L3
	} else {
		goto L146
	}
L143:
	;
	v574 = int32(1)
	goto L145
L144:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+76)))
	v574 = v573
	goto L145
L145:
	;
	goto L142
L146:
	;
	v577 = int32(4322012)
	v579 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v580 = F_list_delete_ptr(m, v579, l1)
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, _consts[119])) = v580
	v585 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+28))
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, _consts[282])) = v586
	F_RelationSetNewRelfilenumber(m, v387, l3)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	v590 = int32(1)
	F_index_build(m, v64, v387, v545, v590, v590)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	v595 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[118])) = v595
	*(*int32)(unsafe.Add(mBase, _consts[281])) = v595
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)))
	if v600 == v595 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v605 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L10
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v661&int32(1) == int32(0) {
		goto L175
	} else {
		goto L176
	}
L154:
	;
	v609 = F_SearchSysCacheCopy(m, int32(34), l1, int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L10
	} else {
		goto L155
	}
L155:
	;
	if v609 == int32(0) {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v609)+16))
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613)+22)))
	v615 = v613 + v614
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+18)))
	if v616 != int32(1) {
		goto L161
	} else {
		goto L162
	}
L157:
	;
	F_sequence_close(m, v605, int32(3))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L10
	} else {
		goto L174
	}
L158:
	;
	v643 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v615)+20)) = uint16(v643)
	v645 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+18)) = uint8(v645)
	F_CatalogTupleUpdate(m, v605, v609+int32(4), v609)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L10
	} else {
		goto L172
	}
L159:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v615)+19)) = uint8(v640)
	goto L158
L160:
	;
	v640 = int32(1)
	goto L159
L161:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+122)))
	if v634 == int32(1) {
		goto L160
	} else {
		goto L171
	}
L162:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+20)))
	if v619 != int32(1) {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+21)))
	if v622 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+122)))
	if v625&int32(1) != 0 {
		goto L160
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+19)))
	if v629 != int32(1) {
		goto L157
	} else {
		goto L168
	}
L167:
	;
	v640 = int32(0)
	goto L159
L168:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+122)))
	if v632 != 0 {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	if v632 != 0 {
		goto L158
	} else {
		goto L170
	}
L170:
	;
	v640 = int32(0)
	goto L159
L171:
	;
	v640 = int32(0)
	goto L159
L172:
	;
	F_CacheInvalidateRelcache(m, v64)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	goto L157
L174:
	;
	goto L153
L175:
	;
	F_AtEOXact_GUC(m, int32(0), v91)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L184
	}
L176:
	;
	v668 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	if v668 == int32(0) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	v672 = F_get_rel_name(m, l1)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L10
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v672
	F_errmsg(m, int32(412104), v15+int32(48))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L10
	} else {
		goto L180
	}
L180:
	;
	v682 = F_pg_rusage_show(m, v15+int32(152))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L10
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v682
	F_errdetail_internal(m, int32(193943), v15+int32(32))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(462298), int32(3896), int32(25539))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	goto L175
L184:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _consts[277])) = v699
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v698
	goto L185
L185:
	;
	F_relation_close(m, v387, int32(0))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_sequence_close(m, v64, int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	if v96 == int32(0) {
		goto L9
	} else {
		goto L188
	}
L188:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _consts[22]))
	if v714 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L9
L190:
	;
	goto L189
L191:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, _consts[24])))
	if v718 != int32(1) {
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v714)+220))
	if v721 == int32(0) {
		goto L190
	} else {
		goto L193
	}
L193:
	;
	v724 = int32(4419940)
	v726 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v727 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v726 + v727
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = v730 + v727
	v734 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v714)+220)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v714)+224)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = v730 + int32(2)
	v744 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v744 - v727
	goto L190
L194:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+68))
	v766 = F_get_namespace_name(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L10
	} else {
		goto L195
	}
L195:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v768 + int32(4)
	F_errmsg_internal(m, int32(640530), v15+int32(16))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(462298), int32(3720), int32(25539))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	F_errmsg(m, int32(132795), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(462298), int32(3729), int32(25539))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	F_errmsg(m, int32(371414), int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(462298), int32(3740), int32(25539))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L10
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v387)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v822 + int32(4)
	F_errmsg(m, int32(656468), v15+int32(80))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(462298), int32(3757), int32(25539))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errmsg_internal(m, int32(307955), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L10
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(462298), int32(4156), int32(308987))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	F_errmsg_internal(m, int32(245085), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L10
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(462298), int32(4203), int32(315825))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l1
	F_errmsg_internal(m, int32(37172), v15-int32(-64))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L10
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(462298), int32(3859), int32(25539))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
