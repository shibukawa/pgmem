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
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int64
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v426 int64
	_ = v426
	var v429 int32
	_ = v429
	var v431 int64
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
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
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
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
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
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
	v841 = m.ExcPending
	if v841 != 0 {
		goto L10
	} else {
		goto L206
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L10
	} else {
		goto L203
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L10
	} else {
		goto L200
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L10
	} else {
		goto L196
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L192
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L10
	} else {
		goto L188
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
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
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v320&int32(4) != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v106)+220)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v106)+224)) = v50
	base.MemoryFill(m, v106+int32(232), int32(0), int32(160))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v132 + v118
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v138 - v118
	goto L36
L39:
	;
	goto L34
L40:
	;
	goto L39
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v156 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v160&int32(1) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v165 = int32(_a_F_reindex_index_4)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v168 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v167 + v168
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v171 + v168
	goto L45
L44:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v302 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v301 + v302
	v305 = int32(_a_F_reindex_index_4)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v307 - v302
	goto L40
L45:
	;
	goto L47
L47:
	;
	goto L48
L48:
	;
	v266 = int32(0)
	v269 = int32(0)
	goto L53
L53:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v15+int32(144)+v269<<(uint(int32(2))%32))))
	v279 = int32(3)
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v15+int32(128)+v269<<(uint(v279)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v156+int32(232)+v278<<(uint(v279)%32)))) = v285
	v287 = int32(1)
	v290 = v266 + v287
	if v290 != int32(2) {
		v266 = v290
		v269 = v269 + v287
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
	if v365 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v323 = m.G0
	v325 = v323 - int32(16)
	m.G0 = v325
	v328 = F_try_relation_open(m, l1, int32(8))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v361 = F_index_open(m, l1, int32(8))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L68
	}
L60:
	;
	m.G0 = v325 + int32(16)
	v365 = v328
	goto L56
L61:
	;
	if v328 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328)+48))
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+119)))
	if v333|int32(32) == int32(105) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v328)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v345 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_21), v325)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_reindex_index_22), int32(204), int32(_a_F_reindex_index_23))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
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
	v365 = v361
	goto L56
L69:
	;
	F_AtEOXact_GUC(m, int32(0), v91)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v372
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v371
	goto L73
L73:
	;
	F_relation_close(m, v64, int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	goto L9
L75:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	v382 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v381)+84)))
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v385 == int32(0) {
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
	v389 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v389&int32(1) == int32(0) {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v394 = int32(_a_F_reindex_index_4)
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v397 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v396 + v397
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v400 + v397
	*(*int64)(unsafe.Add(mBase, uint32(v385+int32(64))+232)) = v382
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v408 + v397
	v414 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v414 - v397
	goto L79
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(1259)
	v421 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v421
	v426 = *(*int64)(unsafe.Add(mBase, _c_F_reindex_index[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+96)) = v426
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = v429
	v431 = *(*int64)(unsafe.Add(mBase, uint32(v15)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v431
	F_EventTriggerCollectSimpleCommand(m, v15+int32(112), v15+int32(96), l0)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L10
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+119)))
	if v440 == int32(73) {
		goto L8
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+118)))
	if v443 == int32(116) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365)+24)))
	if v446 == int32(0) {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v439)+68))
	if v449 != int32(99) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L89
L91:
	;
	if v454 != 0 {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	v452 = F_isTempToastNamespace(m, v449)
	mBase = m.M
	v454 = v452
	goto L94
L93:
	;
	v454 = int32(1)
	goto L94
L94:
	;
	goto L91
L95:
	;
	v455 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L10
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v459 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	if v455 == int32(0) {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_TransferPredicateLocksToHeapRelation(m, v365)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L10
	} else {
		goto L123
	}
L101:
	;
	F_CheckTableNotInUse(m, v365, int32(_a_F_reindex_index_9))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L10
	} else {
		goto L122
	}
L102:
	;
	v463 = int32(1)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v365)+56))
	if base.Ui32(v464) < base.Ui32(int32(_a_F_reindex_index_19)) {
		v473 = v463
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v473 != 0 {
		goto L5
	} else {
		goto L107
	}
L104:
	;
	goto L103
L105:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+68))
	if v468 == int32(99) {
		v473 = v463
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v471 = F_isTempToastNamespace(m, v468)
	mBase = m.M
	v473 = v471
	goto L104
L107:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v474 == int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	v477 = F_CheckRelationTableSpaceMove(m, v365, v474)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	F_CheckTableNotInUse(m, v365, int32(_a_F_reindex_index_9))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v477 == int32(0) {
		goto L100
	} else {
		goto L111
	}
L111:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	F_SetRelationTableSpace(m, v365, v484, int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	F_RelationDropStorage(m, v365)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v491 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v365)+36)) = v491
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v365)+40))
	if v493 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L10
	} else {
		goto L121
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365)+40)) = v491
	goto L117
L116:
	;
	goto L117
L117:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13]))
	if v498 <= int32(31) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v365)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[13])) = v498 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v498<<(uint(int32(2))%32))+uint32(_c_F_reindex_index[14]))) = v501
	goto L114
L119:
	;
	goto L120
L120:
	;
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[15])) = uint8(v512)
	goto L114
L121:
	;
	goto L100
L122:
	;
	goto L100
L123:
	;
	v524 = F_BuildIndexInfo(m, v365)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
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
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+116)))
	if v526 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8]))
	if v541 != 0 {
		goto L4
	} else {
		goto L133
	}
L128:
	;
	v534 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v524)+100)) = v534
	*(*int64)(unsafe.Add(mBase, uint32(v524)+92)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v524)+116)) = uint8(v534)
	goto L127
L129:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v524)+92))
	if v529 == int32(0) {
		goto L128
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)) = uint8(v532)
	goto L128
L132:
	;
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v50
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+72))
	if v549 != 0 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v552&int32(1) != 0 {
		goto L3
	} else {
		goto L138
	}
L135:
	;
	v552 = int32(1)
	goto L137
L136:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+76)))
	v552 = v551
	goto L137
L137:
	;
	goto L134
L138:
	;
	v555 = int32(_a_F_reindex_index_16)
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11]))
	v558 = F_list_delete_ptr(m, v557, l1)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[11])) = v558
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[10]))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+28))
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[12])) = v564
	F_RelationSetNewRelfilenumber(m, v365, l3)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v568 = int32(1)
	F_index_build(m, v64, v365, v524, v568, v568)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	v573 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[9])) = v573
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[8])) = v573
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+327)))
	if v578 == v573 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v583 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L10
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v635&int32(1) == int32(0) {
		goto L165
	} else {
		goto L166
	}
L146:
	;
	v587 = F_SearchSysCacheCopy(m, int32(34), l1, int32(0))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	if v587 == int32(0) {
		goto L2
	} else {
		goto L148
	}
L148:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+22)))
	v593 = v591 + v592
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+18)))
	if v594 != int32(1) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	F_relation_close(m, v583, int32(3))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L10
	} else {
		goto L164
	}
L150:
	;
	v616 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v593)+20)) = uint16(v616)
	v618 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v593)+18)) = uint8(v618)
	*(*uint8)(unsafe.Add(mBase, uint32(v593)+19)) = uint8(v615)
	F_CatalogTupleUpdate(m, v583, v587+int32(4), v587)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L10
	} else {
		goto L162
	}
L151:
	;
	v615 = int32(1)
	goto L150
L152:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+122)))
	if v610 == int32(1) {
		goto L151
	} else {
		goto L161
	}
L153:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+20)))
	if v597 != int32(1) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+21)))
	if v600 == int32(1) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+19)))
	if v603 != int32(1) {
		goto L149
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+122)))
	if v608 != 0 {
		goto L151
	} else {
		goto L160
	}
L158:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+122)))
	if v606 != 0 {
		goto L149
	} else {
		goto L159
	}
L159:
	;
	v615 = int32(0)
	goto L150
L160:
	;
	v615 = int32(0)
	goto L150
L161:
	;
	v615 = int32(0)
	goto L150
L162:
	;
	F_CacheInvalidateRelcache(m, v64)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
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
	v671 = m.ExcPending
	if v671 != 0 {
		goto L10
	} else {
		goto L174
	}
L166:
	;
	v642 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	if v642 == int32(0) {
		goto L165
	} else {
		goto L168
	}
L168:
	;
	v646 = F_get_rel_name(m, l1)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v646
	F_errmsg(m, int32(_a_F_reindex_index_17), v15+int32(48))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v656 = F_pg_rusage_show(m, v15+int32(152))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v656
	F_errdetail_internal(m, int32(_a_F_reindex_index_18), v15+int32(32))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3896), int32(_a_F_reindex_index_6))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	goto L165
L174:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v15)+332))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v15)+328))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[1])) = v673
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[0])) = v672
	goto L175
L175:
	;
	F_relation_close(m, v365, int32(0))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_relation_close(m, v64, int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
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
	v688 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[3]))
	if v688 == int32(0) {
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
	v692 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_index[4])))
	if v692&int32(1) == int32(0) {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v688)+220))
	if v697 == int32(0) {
		goto L180
	} else {
		goto L183
	}
L183:
	;
	v700 = int32(_a_F_reindex_index_4)
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	v703 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v702 + v703
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v688)))
	*(*int32)(unsafe.Add(mBase, uint32(v688))) = v706 + v703
	v710 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v688)+220)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v688)+224)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v688))) = v706 + int32(2)
	v720 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_index[5])) = v720 - v703
	goto L180
L184:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)+68))
	v742 = F_get_namespace_name(m, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L10
	} else {
		goto L185
	}
L185:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v744 + int32(4)
	F_errmsg_internal(m, int32(_a_F_reindex_index_5), v15+int32(16))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3720), int32(_a_F_reindex_index_6))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
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
	v765 = m.ExcPending
	if v765 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	F_errmsg(m, int32(_a_F_reindex_index_7), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3729), int32(_a_F_reindex_index_6))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
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
	v781 = m.ExcPending
	if v781 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	F_errmsg(m, int32(_a_F_reindex_index_8), int32(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3740), int32(_a_F_reindex_index_6))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
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
	v797 = m.ExcPending
	if v797 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v365)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v798 + int32(4)
	F_errmsg(m, int32(_a_F_reindex_index_20), v15+int32(80))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3757), int32(_a_F_reindex_index_6))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_reindex_index_10), int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_11), int32(_a_F_reindex_index_12))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_reindex_index_13), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(_a_F_reindex_index_14), int32(_a_F_reindex_index_15))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
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
	v847 = m.ExcPending
	if v847 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_reindex_index_1), int32(3859), int32(_a_F_reindex_index_6))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
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
