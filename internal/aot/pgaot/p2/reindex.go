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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(559)
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
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int64
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v450 int64
	_ = v450
	var v453 int32
	_ = v453
	var v455 int64
	_ = v455
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
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
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
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
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
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(336)
	m.G0 = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)) = uint8(v6)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	F_getrusage(m, v15+int32(168))
	mBase = m.M
	F_gettimeofday(m, v15+int32(152))
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
	v873 = m.ExcPending
	if v873 != 0 {
		goto L10
	} else {
		goto L206
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L10
	} else {
		goto L203
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L10
	} else {
		goto L200
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L10
	} else {
		goto L196
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L10
	} else {
		goto L192
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L10
	} else {
		goto L188
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L10
	} else {
		goto L184
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
	F_errmsg_internal(m, int32(_a_F_reindex_index_0), v15)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3594), int32(_a_F_reindex_index_2))
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
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(332)))) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(328)))) = v75
	goto L28
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+80))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v79 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v78
	goto L29
L29:
	;
	v87 = int32(_a_F_reindex_index_3)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[2]))
	v91 = v89 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[2])) = v91
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
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
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
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v336&int32(4) != 0 {
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
	v110 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v110&int32(1) == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v115 = int32(_a_F_reindex_index_4)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v118 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v117 + v118
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v121 + v118
	v125 = int32(0)
	v127 = int32(_a_F_reindex_index_5)
	v128 = base.AtomicRmwOr32(m, v125, v127, v125)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+224)) = v50
	base.MemoryFill(m, v106+int32(232), v125, int32(160))
	v139 = base.AtomicRmwOr32(m, v125, v127, v125)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v140 + v118
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v146 - v118
	goto L36
L39:
	;
	goto L34
L40:
	;
	goto L39
L41:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v164 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v168&int32(1) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v173 = int32(_a_F_reindex_index_4)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v176 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v175 + v176
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v179 + v176
	v183 = int32(0)
	v186 = base.AtomicRmwOr32(m, v183, int32(_a_F_reindex_index_5), v183)
	goto L45
L44:
	;
	v313 = int32(0)
	v316 = base.AtomicRmwOr32(m, v313, int32(_a_F_reindex_index_5), v313)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v318 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = v317 + v318
	v321 = int32(_a_F_reindex_index_4)
	v323 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v323 - v318
	goto L40
L45:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v278 = int32(0)
	v281 = int32(0)
	goto L53
L53:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)+v281<<(uint(int32(2))%32))))
	v291 = int32(3)
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(128)+v281<<(uint(v291)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v164+int32(232)+v290<<(uint(v291)%32)))) = v297
	v299 = int32(1)
	v302 = v278 + v299
	if v302 != int32(2) {
		v278 = v302
		v281 = v281 + v299
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
	if v381 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v339 = m.G0
	v341 = v339 - int32(16)
	m.G0 = v341
	v344 = F_try_relation_open(m, l1, int32(8))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L10
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v377 = F_index_open(m, l1, int32(8))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L10
	} else {
		goto L68
	}
L60:
	;
	m.G0 = v341 + int32(16)
	v381 = v344
	goto L56
L61:
	;
	if v344 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+48))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+119)))
	if v349|int32(32) == int32(105) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v344)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v341))) = v361 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_22), v341)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_reindex_index_23), int32(204), int32(_a_F_reindex_index_24))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
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
	v381 = v377
	goto L56
L69:
	;
	F_AtEOXact_GUC(m, int32(0), v91)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L10
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v96 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v388
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v387
	goto L73
L73:
	;
	F_relation_close(m, v64, int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	goto L9
L75:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	v398 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v397)+84)))
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v401 == int32(0) {
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
	v405 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v405&int32(1) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v410 = int32(_a_F_reindex_index_4)
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v413 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v412 + v413
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = v416 + v413
	v420 = int32(0)
	v422 = int32(_a_F_reindex_index_5)
	v423 = base.AtomicRmwOr32(m, v420, v422, v420)
	*(*int64)(unsafe.Add(mBase, uint32(v401+int32(64))+232)) = v398
	v431 = base.AtomicRmwOr32(m, v420, v422, v420)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = v432 + v413
	v438 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v438 - v413
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(1259)
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v445
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v445
	v450 = *(*int64)(unsafe.Add(mBase, _c_F_reindex_index[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v15)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v455
	F_EventTriggerCollectSimpleCommand(m, v15+int32(112), v15+int32(96), l0)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L10
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+119)))
	if v464 == int32(73) {
		goto L8
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+118)))
	if v467 == int32(116) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+24)))
	if v470 == int32(0) {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v463)+68))
	if v473 != int32(99) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	if v478 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v476 = F_isTempToastNamespace(m, v473)
	mBase = m.M
	v478 = v476
	goto L94
L93:
	;
	v478 = int32(1)
	goto L94
L94:
	;
	goto L91
L95:
	;
	v479 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v483 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v479 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_TransferPredicateLocksToHeapRelation(m, v381)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L10
	} else {
		goto L123
	}
L101:
	;
	F_CheckTableNotInUse(m, v381, int32(_a_F_reindex_index_10))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L10
	} else {
		goto L122
	}
L102:
	;
	v487 = int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v381)+56))
	if base.Ui32(v488) < base.Ui32(int32(_a_F_reindex_index_20)) {
		v497 = v487
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v497 != 0 {
		goto L5
	} else {
		goto L107
	}
L104:
	;
	goto L103
L105:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+68))
	if v492 == int32(99) {
		v497 = v487
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v495 = F_isTempToastNamespace(m, v492)
	mBase = m.M
	v497 = v495
	goto L104
L107:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v498 == int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v501 = F_CheckRelationTableSpaceMove(m, v381, v498)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	F_CheckTableNotInUse(m, v381, int32(_a_F_reindex_index_10))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v501 == int32(0) {
		goto L100
	} else {
		goto L111
	}
L111:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	F_SetRelationTableSpace(m, v381, v508, int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	F_RelationDropStorage(m, v381)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v515 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v381)+36)) = v515
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v381)+40))
	if v517 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L10
	} else {
		goto L121
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+40)) = v515
	goto L117
L116:
	;
	goto L117
L117:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13]))
	if v522 <= int32(31) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v381)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13])) = v522 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v522<<(uint(int32(2))%32))+uint32(_c_F_reindex_index[14]))) = v525
	goto L114
L119:
	;
	goto L120
L120:
	;
	v536 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[15])) = uint8(v536)
	goto L114
L121:
	;
	goto L100
L122:
	;
	goto L100
L123:
	;
	v548 = F_BuildIndexInfo(m, v381)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
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
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+116)))
	if v550 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8]))
	if v565 != 0 {
		goto L4
	} else {
		goto L133
	}
L128:
	;
	v558 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+100)) = v558
	*(*int64)(unsafe.Add(mBase, uint32(v548)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v548)+116)) = uint8(v558)
	goto L127
L129:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v548)+92))
	if v553 == int32(0) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v556 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)) = uint8(v556)
	goto L128
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v50
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+72))
	if v573 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v576&int32(1) != 0 {
		goto L3
	} else {
		goto L138
	}
L135:
	;
	v576 = int32(1)
	goto L137
L136:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v572)+76)))
	v576 = v575
	goto L137
L137:
	;
	goto L134
L138:
	;
	v579 = int32(_a_F_reindex_index_17)
	v581 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11]))
	v582 = F_list_delete_ptr(m, v581, l1)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11])) = v582
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v587)+28))
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[12])) = v588
	F_RelationSetNewRelfilenumber(m, v381, l3)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v592 = int32(1)
	F_index_build(m, v64, v381, v548, v592, v592)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	v597 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = v597
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v597
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)))
	if v602 == v597 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v607 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v659&int32(1) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L146:
	;
	v611 = F_SearchSysCacheCopy(m, int32(34), l1, int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	if v611 == int32(0) {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v611)+16))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+22)))
	v617 = v615 + v616
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+18)))
	if v618 != int32(1) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	F_relation_close(m, v607, int32(3))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L10
	} else {
		goto L164
	}
L150:
	;
	v640 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v617)+20)) = uint16(v640)
	v642 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+18)) = uint8(v642)
	*(*uint8)(unsafe.Add(mBase, uint32(v617)+19)) = uint8(v639)
	F_CatalogTupleUpdate(m, v607, v611+int32(4), v611)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L10
	} else {
		goto L162
	}
L151:
	;
	v639 = int32(1)
	goto L150
L152:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+122)))
	if v634 == int32(1) {
		goto L151
	} else {
		goto L161
	}
L153:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+20)))
	if v621 != int32(1) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+21)))
	if v624 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+19)))
	if v627 != int32(1) {
		goto L149
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+122)))
	if v632 != 0 {
		goto L151
	} else {
		goto L160
	}
L158:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+122)))
	if v630 != 0 {
		goto L149
	} else {
		goto L159
	}
L159:
	;
	v639 = int32(0)
	goto L150
L160:
	;
	v639 = int32(0)
	goto L150
L161:
	;
	v639 = int32(0)
	goto L150
L162:
	;
	F_CacheInvalidateRelcache(m, v64)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
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
	F_AtEOXact_GUC(m, int32(0), v91)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L10
	} else {
		goto L174
	}
L166:
	;
	v666 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	if v666 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v670 = F_get_rel_name(m, l1)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v670
	F_errmsg(m, int32(_a_F_reindex_index_18), v15+int32(48))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v680 = F_pg_rusage_show(m, v15+int32(152))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v680
	F_errdetail_internal(m, int32(_a_F_reindex_index_19), v15+int32(32))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3896), int32(_a_F_reindex_index_7))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v697
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v696
	goto L175
L175:
	;
	F_relation_close(m, v381, int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_relation_close(m, v64, int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	if v96 == int32(0) {
		goto L9
	} else {
		goto L178
	}
L178:
	;
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v712 == int32(0) {
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
	v716 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v716&int32(1) == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v712)+220))
	if v721 == int32(0) {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v724 = int32(_a_F_reindex_index_4)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v727 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v726 + v727
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v730 + v727
	v734 = int32(0)
	v736 = int32(_a_F_reindex_index_5)
	v737 = base.AtomicRmwOr32(m, v734, v736, v734)
	*(*int32)(unsafe.Add(mBase, uint32(v712)+220)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v712)+224)) = v734
	v745 = base.AtomicRmwOr32(m, v734, v736, v734)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v712)))
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v746 + v727
	v752 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v752 - v727
	goto L180
L184:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+68))
	v774 = F_get_namespace_name(m, v773)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L10
	} else {
		goto L185
	}
L185:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v776 + int32(4)
	F_errmsg_internal(m, int32(_a_F_reindex_index_6), v15+int32(16))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3720), int32(_a_F_reindex_index_7))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
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
	v797 = m.ExcPending
	if v797 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_reindex_index_8), int32(0))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3729), int32(_a_F_reindex_index_7))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
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
	v813 = m.ExcPending
	if v813 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(_a_F_reindex_index_9), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3740), int32(_a_F_reindex_index_7))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
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
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v381)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v830 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_21), v15+int32(80))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3757), int32(_a_F_reindex_index_7))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
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
	v851 = m.ExcPending
	if v851 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_12), int32(_a_F_reindex_index_13))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
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
	v864 = m.ExcPending
	if v864 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_15), int32(_a_F_reindex_index_16))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l1
	F_errmsg_internal(m, int32(_a_F_reindex_index_0), v15-int32(-64))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3859), int32(_a_F_reindex_index_7))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
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
