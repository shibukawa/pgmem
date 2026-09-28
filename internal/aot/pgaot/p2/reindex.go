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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(601)
	v31 = int32(0)
	v32 = int32(_a_F_ReindexPartitions_0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[0])) = v12 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v12 + int32(8)
	if v14 == int32(112) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = int32(_a_F_ReindexPartitions_1)
	goto L10
L9:
	;
	v46 = int32(_a_F_ReindexPartitions_2)
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
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[0])) = v50
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[1]))
	v58 = F_AllocSetContextCreateInternal(m, v53, int32(_a_F_ReindexPartitions_3), int32(0), int32(_a_F_ReindexPartitions_4), int32(_a_F_ReindexPartitions_5))
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
		goto L24
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
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v66 <= int32(0) {
		v109 = v31
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v71 = int32(0)
	v75 = v31
	goto L17
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+v71<<(uint(int32(2))%32))))
	v84 = F_get_rel_relkind(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L18:
	;
	v109 = v98
	goto L13
L19:
	;
	v101 = v71 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v101 < v102 {
		v71 = v101
		v75 = v98
		goto L17
	} else {
		goto L23
	}
L20:
	;
	v90 = int32(_a_F_ReindexPartitions_6)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[2])) = v58
	v94 = F_lappend_oid(m, v75, v83)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	switch v84&int32(255) - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L20
	default:
		v98 = v75
		goto L19
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexPartitions[2])) = v91
	v98 = v94
	goto L19
L23:
	;
	goto L18
L24:
	;
	F_MemoryContextDelete(m, v58)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v12 + int32(32)
	return
}
func F_reindex_index(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v451 int64
	_ = v451
	var v454 int32
	_ = v454
	var v456 int64
	_ = v456
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(336)
	m.G0 = v16
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+327)) = uint8(v6)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_getrusage(m, v16+int32(168))
	mBase = m.M
	F_gettimeofday(m, v16+int32(152))
	mBase = m.M
	goto L1
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v29 = base.I64_extend_i32_u(l1)
	v30 = F_SearchSysCache1(m, int32(34), v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L10
	} else {
		goto L206
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L10
	} else {
		goto L203
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L10
	} else {
		goto L200
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L196
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L10
	} else {
		goto L192
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L10
	} else {
		goto L188
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L10
	} else {
		goto L184
	}
L9:
	;
	m.G0 = v16 + int32(336)
	return
L10:
	;
	return
L11:
	;
	if v30 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v27&int32(4) != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)+4))
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L19
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg_internal(m, int32(_a_F_reindex_index_0), v16)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3716), int32(_a_F_reindex_index_2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	if v52 == int32(0) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v57&int32(4) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v66 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L22:
	;
	v61 = F_try_table_open(m, v52, int32(5))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v64 = F_table_open(m, v52, int32(5))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	v66 = v61
	goto L21
L26:
	;
	v66 = v64
	goto L21
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(332)))) = v74
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(328)))) = v77
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+48))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+80))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v16)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v81 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v80
	goto L29
L29:
	;
	v89 = int32(_a_F_reindex_index_3)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[2]))
	v93 = v91 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[2])) = v93
	goto L30
L30:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	v98 = v20 & int32(2)
	if v98 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = int64(25769803776)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = int64(3)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v107 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v337&int32(4) != 0 {
		goto L57
	} else {
		goto L58
	}
L35:
	;
	goto L41
L36:
	;
	goto L35
L37:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v111&int32(1) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v116 = int32(_a_F_reindex_index_4)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v119 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v118 + v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v122 + v119
	v126 = int32(0)
	v128 = int32(_a_F_reindex_index_5)
	v129 = base.AtomicRmwOr32(m, v126, v128, v126)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+224)) = v52
	base.MemoryFill(m, v107+int32(232), v126, int32(160))
	v140 = base.AtomicRmwOr32(m, v126, v128, v126)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v141 + v119
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v147 - v119
	goto L36
L39:
	;
	goto L34
L40:
	;
	goto L39
L41:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v165 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v169&int32(1) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v174 = int32(_a_F_reindex_index_4)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v177 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v176 + v177
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v180 + v177
	v184 = int32(0)
	v187 = base.AtomicRmwOr32(m, v184, int32(_a_F_reindex_index_5), v184)
	goto L45
L44:
	;
	v314 = int32(0)
	v317 = base.AtomicRmwOr32(m, v314, int32(_a_F_reindex_index_5), v314)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v319 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v318 + v319
	v322 = int32(_a_F_reindex_index_4)
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v324 - v319
	goto L40
L45:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v279 = int32(0)
	v282 = int32(0)
	goto L53
L53:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(144)+v282<<(uint(int32(2))%32))))
	v292 = int32(3)
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v16+int32(128)+v282<<(uint(v292)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v165+int32(232)+v291<<(uint(v292)%32)))) = v298
	v300 = int32(1)
	v303 = v279 + v300
	if v303 != int32(2) {
		v279 = v303
		v282 = v282 + v300
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L44
L55:
	;
	goto L54
L56:
	;
	if v382 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v340 = m.G0
	v342 = v340 - int32(16)
	m.G0 = v342
	v345 = F_try_relation_open(m, l1, int32(8))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L10
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v378 = F_index_open(m, l1, int32(8))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L10
	} else {
		goto L68
	}
L60:
	;
	m.G0 = v342 + int32(16)
	v382 = v345
	goto L56
L61:
	;
	if v345 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+48))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+119)))
	if v350|int32(32) == int32(105) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v345)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v362 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_22), v342)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_reindex_index_23), int32(205), int32(_a_F_reindex_index_24))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
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
	v382 = v378
	goto L56
L69:
	;
	F_AtEOXact_GUC(m, int32(0), v93)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v98 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v16)+332))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v16)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v389
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v388
	goto L73
L73:
	;
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	goto L9
L75:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v399 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v398)+84)))
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v402 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	goto L77
L77:
	;
	if l0 != 0 {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	goto L77
L79:
	;
	goto L78
L80:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v406&int32(1) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v411 = int32(_a_F_reindex_index_4)
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v414 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v413 + v414
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v417 + v414
	v421 = int32(0)
	v423 = int32(_a_F_reindex_index_5)
	v424 = base.AtomicRmwOr32(m, v421, v423, v421)
	*(*int64)(unsafe.Add(mBase, uint32(v402+int32(64))+232)) = v399
	v432 = base.AtomicRmwOr32(m, v421, v423, v421)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v433 + v414
	v439 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v439 - v414
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = int32(1259)
	v446 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+136)) = v446
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v446
	v451 = *(*int64)(unsafe.Add(mBase, _c_F_reindex_index[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v451
	v454 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v454
	v456 = *(*int64)(unsafe.Add(mBase, uint32(v16)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+112)) = v456
	F_EventTriggerCollectSimpleCommand(m, v16+int32(112), v16+int32(96), l0)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L10
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+119)))
	if v465 == int32(73) {
		goto L8
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+118)))
	if v468 == int32(116) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+24)))
	if v471 == int32(0) {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v464)+68))
	if v474 != int32(99) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	if v479 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v477 = F_isTempToastNamespace(m, v474)
	mBase = m.M
	v479 = v477
	goto L94
L93:
	;
	v479 = int32(1)
	goto L94
L94:
	;
	goto L91
L95:
	;
	v480 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v484 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v480 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_TransferPredicateLocksToHeapRelation(m, v382)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L10
	} else {
		goto L123
	}
L101:
	;
	F_CheckTableNotInUse(m, v382, int32(_a_F_reindex_index_10))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L122
	}
L102:
	;
	v488 = int32(1)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v382)+56))
	if base.Ui32(v489) < base.Ui32(int32(_a_F_reindex_index_20)) {
		v498 = v488
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v498 != 0 {
		goto L5
	} else {
		goto L107
	}
L104:
	;
	goto L103
L105:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+68))
	if v493 == int32(99) {
		v498 = v488
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v496 = F_isTempToastNamespace(m, v493)
	mBase = m.M
	v498 = v496
	goto L104
L107:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v499 == int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v502 = F_CheckRelationTableSpaceMove(m, v382, v499)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	F_CheckTableNotInUse(m, v382, int32(_a_F_reindex_index_10))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v502 == int32(0) {
		goto L100
	} else {
		goto L111
	}
L111:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	F_SetRelationTableSpace(m, v382, v509, int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	F_RelationDropStorage(m, v382)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v516 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v382)+36)) = v516
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v382)+40))
	if v518 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L121
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v382)+40)) = v516
	goto L117
L116:
	;
	goto L117
L117:
	;
	v523 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13]))
	if v523 <= int32(31) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v382)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13])) = v523 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v523<<(uint(int32(2))%32))+uint32(_c_F_reindex_index[14]))) = v526
	goto L114
L119:
	;
	goto L120
L120:
	;
	v537 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[15])) = uint8(v537)
	goto L114
L121:
	;
	goto L100
L122:
	;
	goto L100
L123:
	;
	v549 = F_BuildIndexInfo(m, v382)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	if l2 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+116)))
	if v551 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8]))
	if v566 != 0 {
		goto L4
	} else {
		goto L133
	}
L128:
	;
	v559 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v549)+100)) = v559
	*(*int64)(unsafe.Add(mBase, uint32(v549)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v549)+116)) = uint8(v559)
	goto L127
L129:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v549)+92))
	if v554 == int32(0) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v557 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+327)) = uint8(v557)
	goto L128
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v52
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)+72))
	if v574 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v577&int32(1) != 0 {
		goto L3
	} else {
		goto L138
	}
L135:
	;
	v577 = int32(1)
	goto L137
L136:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573)+76)))
	v577 = v576
	goto L137
L137:
	;
	goto L134
L138:
	;
	v580 = int32(_a_F_reindex_index_17)
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11]))
	v583 = F_list_delete_ptr(m, v582, l1)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11])) = v583
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+28))
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[12])) = v589
	F_RelationSetNewRelfilenumber(m, v382, l3)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v593 = int32(1)
	F_index_build(m, v66, v382, v549, v593, v593, base.B2i32(v98 != int32(0)))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	v600 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = v600
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v600
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+327)))
	if v605 == v600 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v610 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L10
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v662&int32(1) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L146:
	;
	v614 = F_SearchSysCacheCopy(m, int32(34), v29, int64(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	if v614 == int32(0) {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614)+16))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618)+22)))
	v620 = v618 + v619
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+18)))
	if v621 != int32(1) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	F_relation_close(m, v610, int32(3))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L10
	} else {
		goto L164
	}
L150:
	;
	v643 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v620)+20)) = uint16(v643)
	v645 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v620)+18)) = uint8(v645)
	*(*uint8)(unsafe.Add(mBase, uint32(v620)+19)) = uint8(v642)
	F_CatalogTupleUpdate(m, v610, v614+int32(4), v614)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L10
	} else {
		goto L162
	}
L151:
	;
	v642 = int32(1)
	goto L150
L152:
	;
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+122)))
	if v637 == int32(1) {
		goto L151
	} else {
		goto L161
	}
L153:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+20)))
	if v624 != int32(1) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+21)))
	if v627 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+19)))
	if v630 != int32(1) {
		goto L149
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+122)))
	if v635 != 0 {
		goto L151
	} else {
		goto L160
	}
L158:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+122)))
	if v633 != 0 {
		goto L149
	} else {
		goto L159
	}
L159:
	;
	v642 = int32(0)
	goto L150
L160:
	;
	v642 = int32(0)
	goto L150
L161:
	;
	v642 = int32(0)
	goto L150
L162:
	;
	F_CacheInvalidateRelcache(m, v66)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	goto L149
L164:
	;
	goto L145
L165:
	;
	F_AtEOXact_GUC(m, int32(0), v93)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L10
	} else {
		goto L174
	}
L166:
	;
	v669 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	if v669 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v673 = F_get_rel_name(m, l1)
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v673
	F_errmsg(m, int32(_a_F_reindex_index_18), v16+int32(48))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v683 = F_pg_rusage_show(m, v16+int32(152))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v683
	F_errdetail_internal(m, int32(_a_F_reindex_index_19), v16+int32(32))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(4018), int32(_a_F_reindex_index_7))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v16)+332))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v16)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v700
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v699
	goto L175
L175:
	;
	F_relation_close(m, v382, int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_relation_close(m, v66, int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	if v98 == int32(0) {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v715 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L9
L180:
	;
	goto L179
L181:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v719&int32(1) == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v715)+220))
	if v724 == int32(0) {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v727 = int32(_a_F_reindex_index_4)
	v729 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v730 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v729 + v730
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	*(*int32)(unsafe.Add(mBase, uint32(v715))) = v733 + v730
	v737 = int32(0)
	v739 = int32(_a_F_reindex_index_5)
	v740 = base.AtomicRmwOr32(m, v737, v739, v737)
	*(*int32)(unsafe.Add(mBase, uint32(v715)+220)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v715)+224)) = v737
	v748 = base.AtomicRmwOr32(m, v737, v739, v737)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v715)))
	*(*int32)(unsafe.Add(mBase, uint32(v715))) = v749 + v730
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v755 - v730
	goto L180
L184:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)+68))
	v777 = F_get_namespace_name(m, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L185
	}
L185:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v777
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v779 + int32(4)
	F_errmsg_internal(m, int32(_a_F_reindex_index_6), v16+int32(16))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3842), int32(_a_F_reindex_index_7))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_reindex_index_8), int32(0))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3851), int32(_a_F_reindex_index_7))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(_a_F_reindex_index_9), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3862), int32(_a_F_reindex_index_7))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L10
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v382)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v833 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_21), v16+int32(80))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3879), int32(_a_F_reindex_index_7))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(_a_F_reindex_index_11), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_12), int32(_a_F_reindex_index_13))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L10
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	F_errmsg_internal(m, int32(_a_F_reindex_index_14), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_15), int32(_a_F_reindex_index_16))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l1
	F_errmsg_internal(m, int32(_a_F_reindex_index_0), v16-int32(-64))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3981), int32(_a_F_reindex_index_7))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
