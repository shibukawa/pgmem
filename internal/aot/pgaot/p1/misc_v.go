package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_validatePartitionedIndex(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
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
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v15 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = v11 + int32(32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v18, int32(2), int32(3), int32(184), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = int32(1)
	v29 = F_systable_beginscan(m, v15, int32(2187), v26, int32(0), v26, v18)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L38
	}
L6:
	;
	v31 = F_systable_getnext(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = v31
	v40 = v3
	goto L11
L9:
	;
	v66 = v3
	goto L10
L10:
	;
	F_systable_endscan(m, v29)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = F_SearchSysCache1(m, int32(34), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v66 = v56
	goto L10
L13:
	;
	if v46 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v51)+18)))
	F_ReleaseCatCache(m, v46)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v56 = v53 + v40
	v57 = F_systable_getnext(m, v29)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v57 != 0 {
		v35 = v57
		v40 = v56
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	F_relation_close(m, v15, int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = F_RelationGetPartitionDesc(m, l1, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	m.G0 = v11 + int32(80)
	return
L21:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v66 != v75 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v79 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v84 = F_SearchSysCacheCopy(m, int32(34), v82, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v84 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+22)))
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88+v89)+18)) = uint8(v91)
	F_CatalogTupleUpdate(m, v79, v84+int32(4), v84)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_relation_close(m, v79, int32(3))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v84)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+131)))
	if v103 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v110 = F_get_partition_parent(m, v108, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v114 = F_get_partition_parent(m, v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v117 = F_relation_open(m, v110, int32(8))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v120 = F_relation_open(m, v114, int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_validatePartitionedIndex(m, v117, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_relation_close(m, v117, int32(8))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_relation_close(m, v120, int32(8))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v142
	F_errmsg_internal(m, int32(_a_F_validatePartitionedIndex_0), v11+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_validatePartitionedIndex_1), int32(_a_F_validatePartitionedIndex_2), int32(_a_F_validatePartitionedIndex_3))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v158
	F_errmsg_internal(m, int32(_a_F_validatePartitionedIndex_0), v11)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_validatePartitionedIndex_1), int32(_a_F_validatePartitionedIndex_4), int32(_a_F_validatePartitionedIndex_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varcharrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v6+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v21 = F_varchar_input(m, v15, v19, v8, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v21
			}
		}
	}
}
func F_varchartypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anychar_typmodin(m, v3, int32(_a_F_varchartypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v175 int32
	_ = v175
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var __phi391 int32
	_ = __phi391
	var v393 int32
	_ = v393
	var __phi393 int32
	_ = __phi393
	var v399 int32
	_ = v399
	var __phi399 int32
	_ = __phi399
	var v400 int32
	_ = v400
	var __phi400 int32
	_ = __phi400
	var v407 int32
	_ = v407
	var __phi407 int32
	_ = __phi407
	var v410 int32
	_ = v410
	var __phi410 int32
	_ = __phi410
	var v412 int32
	_ = v412
	var __phi412 int32
	_ = __phi412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v500 int32
	_ = v500
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v565 int32
	_ = v565
	var v596 int32
	_ = v596
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v723 int32
	_ = v723
	var v737 int32
	_ = v737
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v796 int32
	_ = v796
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v852 int32
	_ = v852
	var v862 int32
	_ = v862
	var v870 int32
	_ = v870
	var v881 int32
	_ = v881
	var v906 int32
	_ = v906
	var v925 int32
	_ = v925
	var v940 int32
	_ = v940
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	v10 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(16)
	m.G0 = v31
	v33 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = F_pg_mbstrlen_with_len(m, l2, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L143
	}
L5:
	;
	m.G0 = v31 + int32(16)
	return v940
L6:
	;
	v940 = l4 * v37
	goto L5
L7:
	;
	goto L8
L8:
	;
	if v37 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v940 = l5 * v33
	goto L5
L10:
	;
	goto L11
L11:
	;
	if base.B2i32(l8 == int32(0))&(base.B2i32(int32(255) < v33)|base.B2i32(int32(256) <= v37)) != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v54 = v33 + int32(1)
	if l7 < int32(0) {
		v89 = l6
		v90 = l7
		v94 = v54
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if base.B2i32(l1 == v33)&base.B2i32(l3 == v37) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L14:
	;
	v57 = int32(0)
	v58 = v37 - v33
	if v58 < v57 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v80 = base.I32_div_s(l7-v64, v66)
	v86 = v80 - v58>>(uint(int32(31))%32)&v58 + int32(1)
	if v33 < v86 {
		goto L32
	} else {
		goto L33
	}
L16:
	;
	v64 = v57 - l5*v58
	goto L18
L17:
	;
	v64 = l4 * v58
	goto L18
L18:
	;
	if v64 <= l7 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = l4 + l5
	if v66 < l6 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v940 = l7 + int32(1)
	goto L5
L22:
	;
	v68 = v66
	goto L24
L23:
	;
	v68 = l6
	goto L24
L24:
	;
	if v33 < v37 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v70 = v33
	goto L27
L26:
	;
	v70 = v37
	goto L27
L27:
	;
	if v64+v68*v70 <= l7 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v89 = v68
	v90 = int32(-1)
	v94 = v54
	goto L13
L29:
	;
	goto L30
L30:
	;
	if int32(0) < v66 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v89 = v68
	v90 = l7
	v94 = v54
	goto L13
L32:
	;
	v88 = v54
	goto L34
L33:
	;
	v88 = v86
	goto L34
L34:
	;
	v89 = v68
	v90 = l7
	v94 = v88
	goto L13
L35:
	;
	v102 = F_palloc(m, v54<<(uint(int32(2))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v201 = v10
	goto L37
L37:
	;
	v211 = F_palloc(m, v54<<(uint(int32(3))%32))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	v104 = int32(0)
	if v104 < v33 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v109 = l0
	v117 = v104
	goto L42
L40:
	;
	v175 = int32(0)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102+v175<<(uint(int32(2))%32)))) = int32(0)
	v201 = v102
	goto L37
L42:
	;
	v139 = F_pg_mblen_range(m, v109, l0+l1)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v175 = v33
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v102+v117<<(uint(int32(2))%32)))) = v139
	v144 = v117 + int32(1)
	if v144 != v33 {
		v109 = v109 + v139
		v117 = v144
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	if v94 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	if v37+int32(1) < int32(2) {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	v218 = v94 & int32(3)
	v219 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v94) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v227 = v219
	v238 = int32(0)
	goto L52
L50:
	;
	v288 = v219
	goto L51
L51:
	;
	v316 = v288
	v325 = v219
	goto L56
L52:
	;
	v254 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v211+v227<<(uint(v254)%32)))) = v227 * l5
	v260 = v227 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211+v260<<(uint(v254)%32)))) = l5 * v260
	v267 = v227 | v254
	*(*int32)(unsafe.Add(mBase, uint32(v211+v267<<(uint(v254)%32)))) = l5 * v267
	v274 = v227 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v211+v274<<(uint(v254)%32)))) = l5 * v274
	v280 = int32(4)
	v281 = v227 + v280
	v283 = v238 + v280
	if v283 != v94&int32(2147483644) {
		v227 = v281
		v238 = v283
		goto L52
	} else {
		goto L54
	}
L53:
	;
	if v218 == int32(0) {
		goto L47
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	v288 = v281
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211+v316<<(uint(int32(2))%32)))) = v316 * l5
	v348 = int32(1)
	v351 = v325 + v348
	if v351 != v218 {
		v316 = v316 + v348
		v325 = v351
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L47
L58:
	;
	goto L57
L59:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v906+v33<<(uint(int32(2))%32))))
	v940 = v925
	goto L5
L60:
	;
	v906 = v211
	goto L59
L61:
	;
	goto L62
L62:
	;
	v388 = int32(1)
	v389 = v90 + v388
	__phi391 = l0
	__phi393 = l2
	__phi399 = v94
	__phi400 = v211
	__phi407 = v211 + v54<<(uint(int32(2))%32)
	__phi410 = v10
	__phi412 = v388
	v391 = __phi391
	v393 = __phi393
	v399 = __phi399
	v400 = __phi400
	v407 = __phi407
	v410 = __phi410
	v412 = __phi412
	goto L63
L63:
	;
	if l3 != v37 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v906 = v407
	goto L59
L65:
	;
	v421 = F_pg_mblen_range(m, v393, l2+l3)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	v423 = int32(1)
	goto L67
L67:
	;
	if v54 <= v399 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v423 = v421
	goto L67
L69:
	;
	v431 = v399
	goto L71
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400+v399<<(uint(int32(2))%32)))) = v389
	v431 = v399 + int32(1)
	goto L71
L71:
	;
	if v410 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = l4 * v412
	v437 = int32(1)
	goto L74
L73:
	;
	v437 = v410
	goto L74
L74:
	;
	if v201 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v90 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L76:
	;
	if v431 <= v437 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v431 <= v437 {
		goto L75
	} else {
		goto L99
	}
L79:
	;
	v450 = v437
	v451 = v391
	goto L80
L80:
	;
	v470 = int32(2)
	v471 = v450 << (uint(v470) % 32)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v400+v471)))
	v474 = v473 + l4
	v479 = (v450 - int32(1)) << (uint(v470) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v407+v479)))
	v482 = v481 + l5
	if v474 < v482 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L75
L82:
	;
	v484 = v474
	goto L84
L83:
	;
	v484 = v482
	goto L84
L84:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v201+v479)))
	v487 = v451 + v486
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487-int32(1)))))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v423-int32(1)))))
	if base.B2i32(v490 != v491)|base.B2i32(v486 != v423) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	if v484 < v625 {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v400+v479)))
	v625 = v596
	goto L85
L87:
	;
	if v423 == int32(1) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v400+v479)))
	v625 = v565 + v89
	goto L85
L90:
	;
	v500 = v423
	goto L91
L91:
	;
	if v500 <= int32(0) {
		goto L86
	} else {
		goto L93
	}
L92:
	;
	goto L89
L93:
	;
	v530 = v500 - int32(1)
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451+v530))))
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530+v393))))
	if v532 == v534 {
		v500 = v530
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v627 = v484
	goto L97
L96:
	;
	v627 = v625
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471+v407))) = v627
	v630 = v450 + int32(1)
	if v630 != v431 {
		v450 = v630
		v451 = v487
		goto L80
	} else {
		goto L98
	}
L98:
	;
	goto L81
L99:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v407+v437<<(uint(int32(2))%32)-int32(4))))
	v640 = v391
	v647 = v437
	v648 = v638
	goto L100
L100:
	;
	v668 = v647 << (uint(int32(2)) % 32)
	v670 = v668 + v400
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	v672 = v671 + l4
	v673 = l5 + v648
	if v672 < v673 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L75
L102:
	;
	v675 = v672
	goto L104
L103:
	;
	v675 = v673
	goto L104
L104:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v670-int32(4))))
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640))))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	if v680 != v681 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v683 = v89
	goto L107
L106:
	;
	v683 = int32(0)
	goto L107
L107:
	;
	v684 = v678 + v683
	if v675 < v684 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v686 = v675
	goto L110
L109:
	;
	v686 = v684
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v407+v668))) = v686
	v688 = int32(1)
	v691 = v647 + v688
	if v691 != v431 {
		v640 = v640 + v688
		v647 = v691
		v648 = v686
		goto L100
	} else {
		goto L111
	}
L111:
	;
	goto L101
L112:
	;
	if v37 != v412 {
		__phi391 = v862
		__phi393 = v393 + v423
		__phi399 = v870
		__phi400 = v407
		__phi407 = v400
		__phi410 = v881
		__phi412 = v412 + int32(1)
		v391 = __phi391
		v393 = __phi393
		v399 = __phi399
		v400 = __phi400
		v407 = __phi407
		v410 = __phi410
		v412 = __phi412
		goto L63
	} else {
		goto L142
	}
L113:
	;
	v862 = v391
	v870 = v431
	v881 = v410
	goto L112
L114:
	;
	goto L115
L115:
	;
	v723 = v412 + (v33 - v37)
	v737 = v431
	goto L116
L116:
	;
	if v737 <= int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	if v773 <= v410 {
		v833 = v391
		v852 = v410
		goto L126
	} else {
		goto L127
	}
L118:
	;
	goto L117
L119:
	;
	v773 = v431 >> (uint(int32(31)) % 32) & v431
	goto L118
L120:
	;
	goto L121
L121:
	;
	v758 = v737 - int32(1)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v407+v758<<(uint(int32(2))%32))))
	v763 = v758 - v723
	v765 = int32(0)
	if v765 < v763 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v770 = v763 * l4
	goto L124
L123:
	;
	v770 = v765 - l5*v763
	goto L124
L124:
	;
	if v90 < v762+v770 {
		v737 = v758
		goto L116
	} else {
		goto L125
	}
L125:
	;
	v773 = v737
	goto L118
L126:
	;
	if v773 <= v852 {
		v940 = v389
		goto L5
	} else {
		goto L141
	}
L127:
	;
	v777 = v391
	v796 = v410
	goto L128
L128:
	;
	v806 = v796 << (uint(int32(2)) % 32)
	v807 = v407 + v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)))
	v809 = v796 - v723
	v811 = int32(0)
	if v811 < v809 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v940 = v389
	goto L5
L130:
	;
	v816 = v809 * l4
	goto L132
L131:
	;
	v816 = v811 - l5*v809
	goto L132
L132:
	;
	if v808+v816 <= v90 {
		v833 = v777
		v852 = v796
		goto L126
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v807))) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v806+v400))) = v389
	if v796 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if v201 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v829 = v777
	goto L136
L136:
	;
	v831 = v796 + int32(1)
	if v831 != v773 {
		v777 = v829
		v796 = v831
		goto L128
	} else {
		goto L140
	}
L137:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v806+v201-int32(4))))
	v827 = v825
	goto L139
L138:
	;
	v827 = int32(1)
	goto L139
L139:
	;
	v829 = v827 + v777
	goto L136
L140:
	;
	goto L129
L141:
	;
	v862 = v833
	v870 = v773
	v881 = v852
	goto L112
L142:
	;
	goto L64
L143:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(255)
	F_errmsg(m, int32(_a_F_varstr_levenshtein_less_equal_0), v31)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_less_equal_1), int32(135), int32(_a_F_varstr_levenshtein_less_equal_2))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_visibilitymap_prepare_truncate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
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
	var v25 int64
	_ = v25
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = base.I32_div_u_s(l1, int32(_a_F_visibilitymap_prepare_truncate_0))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v25
	v29 = F_smgropen(m, v15+int32(16), v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v19
	goto L3
L3:
	;
	v49 = int32(-1)
	v51 = F_smgrexists(m, v48, int32(2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v47
	goto L3
L7:
	;
	v43 = v35
	goto L9
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	v43 = v41
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v43 + int32(1)
	goto L6
L10:
	;
	m.G0 = v15 + int32(32)
	return v228
L11:
	;
	if v51 == int32(0) {
		v228 = v49
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v57 = l1 - v18*int32(_a_F_visibilitymap_prepare_truncate_0)
	v59 = int32(base.Ui32(v57) >> (uint(int32(2)) % 32))
	v63 = l1 << (uint(int32(1)) % 32) & int32(6)
	if v59|v63 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = F_vm_readbuf(m, l0, v18, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v188 = v18
	goto L15
L15:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v195 != 0 {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	if v66 == int32(0) {
		v228 = v49
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v66 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_LockBuffer(m, v66, int32(2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[0]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73+(v66^int32(-1))<<(uint(int32(2))%32))))
	v87 = v79
	goto L18
L20:
	;
	goto L21
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[1]))
	v87 = v81 + v66<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	v91 = int32(_a_F_visibilitymap_prepare_truncate_1)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2])) = v93 + int32(1)
	v97 = v59 + v87
	v99 = v97 + int32(24)
	v101 = v97 + int32(25)
	v102 = int32(3)
	v105 = int32(_a_F_visibilitymap_prepare_truncate_2) - v59
	if v101&v102|v105&v102|base.B2i32(base.Ui32(v57) < base.Ui32(int32(_a_F_visibilitymap_prepare_truncate_3))) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v140 = int32(-1)
	v144 = v139 & (v140<<(uint(v63)%32) ^ v140)
	*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v144)
	F_MarkBufferDirty(m, v66)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L33
	}
L24:
	;
	if v59 == int32(_a_F_visibilitymap_prepare_truncate_2) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v105 == int32(0) {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v117 = v97 + int32(29)
	v119 = v87 - int32(-8192)
	if base.Ui32(v119) < base.Ui32(v117) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v121 = v117
	goto L30
L29:
	;
	v121 = v119
	goto L30
L30:
	;
	v128 = (v121-v97-int32(26))&int32(-4) + int32(4)
	if v128 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.MemoryFill(m, v101, int32(0), v128)
	goto L23
L32:
	;
	base.MemoryFill(m, v101, int32(0), v105)
	goto L23
L33:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[3])))
	if v149 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v176 = int32(1)
	v178 = int32(_a_F_visibilitymap_prepare_truncate_1)
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2])) = v180 - v176
	F_UnlockReleaseBuffer(m, v66)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L48
	}
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+118)))
	if v151 != int32(112) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[4]))
	if v155 <= int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v158 != 0 {
		goto L34
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[5]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+252))
	goto L42
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v159 != 0 {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if base.B2i32(v162 != int32(0)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[6])))
	if v168&int32(1) == int32(0) {
		goto L34
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_log_newpage_buffer(m, v66, int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	goto L34
L48:
	;
	v188 = v18 + v176
	goto L15
L49:
	;
	v219 = v195
	goto L51
L50:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v197
	v199 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v199
	v201 = F_smgropen(m, v15, v196)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	v221 = F_smgrnblocks(m, v219, int32(2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v201
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+72))
	if v205 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v219 = v217
	goto L51
L54:
	;
	v213 = v205
	goto L56
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201)+76))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v201)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v201)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v201)+72))
	v213 = v211
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+72)) = v213 + int32(1)
	goto L53
L57:
	;
	if base.Ui32(v221) <= base.Ui32(v188) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v224 = int32(-1)
	goto L60
L59:
	;
	v224 = v188
	goto L60
L60:
	;
	v228 = v224
	goto L10
}
func F_visibilitymap_set(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v224 int64
	_ = v224
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	v7 = l6
	v12 = base.I32_div_u_s(l1, int32(_a_F_visibilitymap_set_0))
	if l2 != 0 {
		if l2 < int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[0]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v31 = v22
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[1]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+l2<<(uint(int32(6))%32)+int32(-64))+16))
			v31 = v30
		}
		if v31 != l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v250 = m.ExcPending
			if v250 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_set_1), int32(0))
				mBase = m.M
				v254 = m.ExcPending
				if v254 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(270), int32(_a_F_visibilitymap_set_3))
					mBase = m.M
					v259 = m.ExcPending
					if v259 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if l4 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
					mBase = m.M
					v267 = m.ExcPending
					if v267 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l4 < int32(0) {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[0]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(l4^int32(-1))<<(uint(int32(6))%32))+16))
					v53 = v44
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[1]))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+l4<<(uint(int32(6))%32)+int32(-64))+16))
					v53 = v52
				}
				if v53 != v12 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
						mBase = m.M
						v267 = m.ExcPending
						if v267 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
							mBase = m.M
							v272 = m.ExcPending
							if v272 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v58 = l1 << (uint(int32(1)) % 32) & int32(6)
					if l4 < int32(0) {
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(l4^int32(-1))<<(uint(int32(2))%32))))
						v81 = v73
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
						v81 = v75 + l4<<(uint(int32(13))%32) + int32(-8192)
					}
					F_LockBuffer(m, l4, int32(2))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v87 = int32(base.Ui32(l1-v12*int32(_a_F_visibilitymap_set_0))>>(uint(int32(2))%32)) + v81
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+24)))
						v91 = int32(base.Ui32(v88)>>(uint(v58)%32)) & int32(3)
						if v7 != v91 {
							v93 = int32(_a_F_visibilitymap_set_5)
							v95 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v95 + int32(1)
							v100 = v87 + int32(24)
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
							v103 = v101 | v7<<(uint(v58)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
							F_MarkBufferDirty(m, l4)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+118)))
								if v108 != int32(112) {
									v233 = int32(_a_F_visibilitymap_set_5)
									v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
									F_LockBuffer(m, l4, int32(0))
									mBase = m.M
									v245 = m.ExcPending
									if v245 != 0 {
										return int32(0)
									} else {
										return v91
									}
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
									if v112 <= int32(0) {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v115 != 0 {
											v233 = int32(_a_F_visibilitymap_set_5)
											v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int32(0)
											} else {
												return v91
											}
										} else {
											v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v116 != 0 {
												v233 = int32(_a_F_visibilitymap_set_5)
												v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
												*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
												F_LockBuffer(m, l4, int32(0))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
													return int32(0)
												} else {
													return v91
												}
											} else {
												if l3 != int64(0) {
													v224 = l3
													*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
													v233 = int32(_a_F_visibilitymap_set_5)
													v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
													*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
													F_LockBuffer(m, l4, int32(0))
													mBase = m.M
													v245 = m.ExcPending
													if v245 != 0 {
														return int32(0)
													} else {
														return v91
													}
												} else {
													v119 = m.G0
													v121 = v119 - int32(16)
													m.G0 = v121
													*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v7)
													*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = l5
													v126 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
													if v126 < int32(2) {
													} else {
														v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
														if v130 != int32(112) {
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															if base.B2i32(base.Ui32(v133) < base.Ui32(int32(_a_F_visibilitymap_set_6))) == int32(0) {
																v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
																if v138 == int32(0) {
																} else {
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
																	switch v142 - int32(109) {
																	case 0, 5:
																		v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+104)))
																		if v145 != int32(1) {
																		} else {
																			v150 = v7 | int32(4)
																			*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
																		}
																	default:
																	}
																}
															} else {
																v150 = v7 | int32(4)
																*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
															}
														}
													}
													F_XLogBeginInsert(m)
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return int32(0)
													} else {
														F_XLogRegisterData(m, v121+int32(8), int32(5))
														mBase = m.M
														v159 = m.ExcPending
														if v159 != 0 {
															return int32(0)
														} else {
															v160 = int32(0)
															F_XLogRegisterBuffer(m, v160, l4, v160)
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return int32(0)
															} else {
																v165 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
																v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
																v169 = int32(1)
																v170 = int32(8)
																v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
																if v174&v169 != 0 {
																	v177 = v170
																} else {
																	v177 = int32(10)
																}
																if v166 != int32(0) {
																	v178 = v170
																} else {
																	v178 = v177
																}
																F_XLogRegisterBuffer(m, v169, l2, v178)
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return int32(0)
																} else {
																	v183 = F_XLogInsert(m, int32(9), int32(64))
																	mBase = m.M
																	v184 = m.ExcPending
																	if v184 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v121 + int32(16)
																		v189 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
																		v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																		if base.B2i32(v190 != int32(0)) == int32(0) {
																			v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
																			if v196&int32(1) == int32(0) {
																				v224 = v183
																			} else {
																				if l2 < int32(0) {
																					v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																					v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																					v218 = v210
																				} else {
																					v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																					v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																				}
																				*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																				v224 = v183
																			}
																		} else {
																			if l2 < int32(0) {
																				v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																				v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v218 = v210
																			} else {
																				v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																				v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																			v224 = v183
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
																		v233 = int32(_a_F_visibilitymap_set_5)
																		v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																		*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
																		F_LockBuffer(m, l4, int32(0))
																		mBase = m.M
																		v245 = m.ExcPending
																		if v245 != 0 {
																			return int32(0)
																		} else {
																			return v91
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
										if l3 != int64(0) {
											v224 = l3
											*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
											v233 = int32(_a_F_visibilitymap_set_5)
											v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int32(0)
											} else {
												return v91
											}
										} else {
											v119 = m.G0
											v121 = v119 - int32(16)
											m.G0 = v121
											*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v7)
											*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = l5
											v126 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
											if v126 < int32(2) {
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
												if v130 != int32(112) {
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if base.B2i32(base.Ui32(v133) < base.Ui32(int32(_a_F_visibilitymap_set_6))) == int32(0) {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
														if v138 == int32(0) {
														} else {
															v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
															v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
															switch v142 - int32(109) {
															case 0, 5:
																v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+104)))
																if v145 != int32(1) {
																} else {
																	v150 = v7 | int32(4)
																	*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
																}
															default:
															}
														}
													} else {
														v150 = v7 | int32(4)
														*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
													}
												}
											}
											F_XLogBeginInsert(m)
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												F_XLogRegisterData(m, v121+int32(8), int32(5))
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													v160 = int32(0)
													F_XLogRegisterBuffer(m, v160, l4, v160)
													mBase = m.M
													v163 = m.ExcPending
													if v163 != 0 {
														return int32(0)
													} else {
														v165 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
														v169 = int32(1)
														v170 = int32(8)
														v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
														if v174&v169 != 0 {
															v177 = v170
														} else {
															v177 = int32(10)
														}
														if v166 != int32(0) {
															v178 = v170
														} else {
															v178 = v177
														}
														F_XLogRegisterBuffer(m, v169, l2, v178)
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return int32(0)
														} else {
															v183 = F_XLogInsert(m, int32(9), int32(64))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int32(0)
															} else {
																m.G0 = v121 + int32(16)
																v189 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
																v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																if base.B2i32(v190 != int32(0)) == int32(0) {
																	v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
																	if v196&int32(1) == int32(0) {
																		v224 = v183
																	} else {
																		if l2 < int32(0) {
																			v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v218 = v210
																		} else {
																			v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																			v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																		v224 = v183
																	}
																} else {
																	if l2 < int32(0) {
																		v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																		v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v218 = v210
																	} else {
																		v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																		v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																	v224 = v183
																}
																*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
																v233 = int32(_a_F_visibilitymap_set_5)
																v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
																F_LockBuffer(m, l4, int32(0))
																mBase = m.M
																v245 = m.ExcPending
																if v245 != 0 {
																	return int32(0)
																} else {
																	return v91
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
							F_LockBuffer(m, l4, int32(0))
							mBase = m.M
							v245 = m.ExcPending
							if v245 != 0 {
								return int32(0)
							} else {
								return v91
							}
						}
					}
				}
			}
		}
	} else {
		if l4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v263 = m.ExcPending
			if v263 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
				mBase = m.M
				v267 = m.ExcPending
				if v267 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if l4 < int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[0]))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(l4^int32(-1))<<(uint(int32(6))%32))+16))
				v53 = v44
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[1]))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+l4<<(uint(int32(6))%32)+int32(-64))+16))
				v53 = v52
			}
			if v53 != v12 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
					mBase = m.M
					v267 = m.ExcPending
					if v267 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
						mBase = m.M
						v272 = m.ExcPending
						if v272 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v58 = l1 << (uint(int32(1)) % 32) & int32(6)
				if l4 < int32(0) {
					v67 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(l4^int32(-1))<<(uint(int32(2))%32))))
					v81 = v73
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
					v81 = v75 + l4<<(uint(int32(13))%32) + int32(-8192)
				}
				F_LockBuffer(m, l4, int32(2))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v87 = int32(base.Ui32(l1-v12*int32(_a_F_visibilitymap_set_0))>>(uint(int32(2))%32)) + v81
					v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+24)))
					v91 = int32(base.Ui32(v88)>>(uint(v58)%32)) & int32(3)
					if v7 != v91 {
						v93 = int32(_a_F_visibilitymap_set_5)
						v95 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v95 + int32(1)
						v100 = v87 + int32(24)
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
						v103 = v101 | v7<<(uint(v58)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v103)
						F_MarkBufferDirty(m, l4)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+118)))
							if v108 != int32(112) {
								v233 = int32(_a_F_visibilitymap_set_5)
								v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
								F_LockBuffer(m, l4, int32(0))
								mBase = m.M
								v245 = m.ExcPending
								if v245 != 0 {
									return int32(0)
								} else {
									return v91
								}
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
								if v112 <= int32(0) {
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v115 != 0 {
										v233 = int32(_a_F_visibilitymap_set_5)
										v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
										*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
										F_LockBuffer(m, l4, int32(0))
										mBase = m.M
										v245 = m.ExcPending
										if v245 != 0 {
											return int32(0)
										} else {
											return v91
										}
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v116 != 0 {
											v233 = int32(_a_F_visibilitymap_set_5)
											v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v245 = m.ExcPending
											if v245 != 0 {
												return int32(0)
											} else {
												return v91
											}
										} else {
											if l3 != int64(0) {
												v224 = l3
												*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
												v233 = int32(_a_F_visibilitymap_set_5)
												v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
												*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
												F_LockBuffer(m, l4, int32(0))
												mBase = m.M
												v245 = m.ExcPending
												if v245 != 0 {
													return int32(0)
												} else {
													return v91
												}
											} else {
												v119 = m.G0
												v121 = v119 - int32(16)
												m.G0 = v121
												*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v7)
												*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = l5
												v126 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
												if v126 < int32(2) {
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
													if v130 != int32(112) {
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if base.B2i32(base.Ui32(v133) < base.Ui32(int32(_a_F_visibilitymap_set_6))) == int32(0) {
															v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
															if v138 == int32(0) {
															} else {
																v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
																v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
																switch v142 - int32(109) {
																case 0, 5:
																	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+104)))
																	if v145 != int32(1) {
																	} else {
																		v150 = v7 | int32(4)
																		*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
																	}
																default:
																}
															}
														} else {
															v150 = v7 | int32(4)
															*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
														}
													}
												}
												F_XLogBeginInsert(m)
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													F_XLogRegisterData(m, v121+int32(8), int32(5))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v160 = int32(0)
														F_XLogRegisterBuffer(m, v160, l4, v160)
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return int32(0)
														} else {
															v165 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
															v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
															v169 = int32(1)
															v170 = int32(8)
															v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
															if v174&v169 != 0 {
																v177 = v170
															} else {
																v177 = int32(10)
															}
															if v166 != int32(0) {
																v178 = v170
															} else {
																v178 = v177
															}
															F_XLogRegisterBuffer(m, v169, l2, v178)
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return int32(0)
															} else {
																v183 = F_XLogInsert(m, int32(9), int32(64))
																mBase = m.M
																v184 = m.ExcPending
																if v184 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v121 + int32(16)
																	v189 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
																	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																	if base.B2i32(v190 != int32(0)) == int32(0) {
																		v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
																		if v196&int32(1) == int32(0) {
																			v224 = v183
																		} else {
																			if l2 < int32(0) {
																				v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																				v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v218 = v210
																			} else {
																				v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																				v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																			v224 = v183
																		}
																	} else {
																		if l2 < int32(0) {
																			v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v218 = v210
																		} else {
																			v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																			v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																		v224 = v183
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
																	v233 = int32(_a_F_visibilitymap_set_5)
																	v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
																	F_LockBuffer(m, l4, int32(0))
																	mBase = m.M
																	v245 = m.ExcPending
																	if v245 != 0 {
																		return int32(0)
																	} else {
																		return v91
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
									if l3 != int64(0) {
										v224 = l3
										*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
										v233 = int32(_a_F_visibilitymap_set_5)
										v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
										*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
										F_LockBuffer(m, l4, int32(0))
										mBase = m.M
										v245 = m.ExcPending
										if v245 != 0 {
											return int32(0)
										} else {
											return v91
										}
									} else {
										v119 = m.G0
										v121 = v119 - int32(16)
										m.G0 = v121
										*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v7)
										*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = l5
										v126 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
										if v126 < int32(2) {
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
											if v130 != int32(112) {
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if base.B2i32(base.Ui32(v133) < base.Ui32(int32(_a_F_visibilitymap_set_6))) == int32(0) {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
													if v138 == int32(0) {
													} else {
														v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+119)))
														switch v142 - int32(109) {
														case 0, 5:
															v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+104)))
															if v145 != int32(1) {
															} else {
																v150 = v7 | int32(4)
																*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
															}
														default:
														}
													}
												} else {
													v150 = v7 | int32(4)
													*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v150)
												}
											}
										}
										F_XLogBeginInsert(m)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											F_XLogRegisterData(m, v121+int32(8), int32(5))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												v160 = int32(0)
												F_XLogRegisterBuffer(m, v160, l4, v160)
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int32(0)
												} else {
													v165 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
													v169 = int32(1)
													v170 = int32(8)
													v174 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
													if v174&v169 != 0 {
														v177 = v170
													} else {
														v177 = int32(10)
													}
													if v166 != int32(0) {
														v178 = v170
													} else {
														v178 = v177
													}
													F_XLogRegisterBuffer(m, v169, l2, v178)
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return int32(0)
													} else {
														v183 = F_XLogInsert(m, int32(9), int32(64))
														mBase = m.M
														v184 = m.ExcPending
														if v184 != 0 {
															return int32(0)
														} else {
															m.G0 = v121 + int32(16)
															v189 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[6]))
															v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
															if base.B2i32(v190 != int32(0)) == int32(0) {
																v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_set[7])))
																if v196&int32(1) == int32(0) {
																	v224 = v183
																} else {
																	if l2 < int32(0) {
																		v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																		v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v218 = v210
																	} else {
																		v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																		v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																	v224 = v183
																}
															} else {
																if l2 < int32(0) {
																	v204 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v204+(l2^int32(-1))<<(uint(int32(2))%32))))
																	v218 = v210
																} else {
																	v212 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																	v218 = v212 + l2<<(uint(int32(13))%32) + int32(-8192)
																}
																*(*int64)(unsafe.Add(mBase, uint32(v218))) = base.I64_rotr(v183, int64(32))
																v224 = v183
															}
															*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v224, int64(32))
															v233 = int32(_a_F_visibilitymap_set_5)
															v235 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
															*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v235 - int32(1)
															F_LockBuffer(m, l4, int32(0))
															mBase = m.M
															v245 = m.ExcPending
															if v245 != 0 {
																return int32(0)
															} else {
																return v91
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
						F_LockBuffer(m, l4, int32(0))
						mBase = m.M
						v245 = m.ExcPending
						if v245 != 0 {
							return int32(0)
						} else {
							return v91
						}
					}
				}
			}
		}
	}
}
func F_vm_readbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
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
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v17
	v21 = F_smgropen(m, v9+int32(56), v14)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v40 = v11
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	if v41 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v21
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = v39
	goto L3
L7:
	;
	v35 = v27
	goto L9
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	v35 = v33
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v35 + int32(1)
	goto L6
L10:
	;
	m.G0 = v9 + int32(80)
	return v216
L11:
	;
	if v118 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v62 = int32(0)
	if l2 == v62 {
		v216 = v62
		goto L10
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = int32(0)
	goto L12
L14:
	;
	v45 = F_smgrexists(m, v40, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v53 = v41
	goto L16
L16:
	;
	if base.Ui32(v53) <= base.Ui32(l1) {
		goto L12
	} else {
		goto L20
	}
L17:
	;
	if v45 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v50 = F_smgrnblocks(m, v40, int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	v53 = v52
	goto L16
L20:
	;
	v58 = F_ReadBufferExtended(m, l0, int32(2), l1, int32(3), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v118 = v58
	goto L11
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = l0
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v9)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v70
	v79 = F_ExtendBufferedRelTo(m, v9+int32(40), int32(2), int32(20), l1+int32(1), int32(3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v81 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v87
	v91 = F_smgropen(m, v9+int32(24), v84)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v108 = v81
	goto L26
L26:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v109
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v111
	F_CacheInvalidateSmgr(m, v9+int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v91
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)+72))
	if v95 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v108 = v107
	goto L26
L29:
	;
	v103 = v95
	goto L31
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91)+72))
	v103 = v101
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+72)) = v103 + int32(1)
	goto L28
L32:
	;
	v118 = v79
	goto L11
L33:
	;
	F_LockBuffer(m, v118, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L55
	}
L34:
	;
	v160 = int32(_a_F_vm_readbuf_0)
	v161 = int32(0)
	if v161|(v159&int32(3)|int32(1)) == v161 {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v124 = (v118 ^ int32(-1)) << (uint(int32(2)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v126)))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+14)))
	if v129 != 0 {
		v216 = v118
		goto L10
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v141 = v118 << (uint(int32(13)) % 32)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[1]))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+v143-int32(_a_F_vm_readbuf_3)))))
	if v147 != 0 {
		v216 = v118
		goto L10
	} else {
		goto L41
	}
L38:
	;
	F_LockBuffer(m, v118, int32(2))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134+v124)))
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+14)))
	if v137 == int32(0) {
		v159 = v136
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	F_LockBuffer(m, v118, int32(2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[1]))
	v153 = v152 + v141
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153-int32(_a_F_vm_readbuf_3)))))
	if v156 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v159 = v153 + int32(-8192)
	goto L34
L44:
	;
	goto L33
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+10)) = int32(_a_F_vm_readbuf_1)
	v201 = int32(_a_F_vm_readbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v159)+18)) = uint16(v201)
	v207 = int32(_a_F_vm_readbuf_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v159)+16)) = uint16(v207)
	*(*uint16)(unsafe.Add(mBase, uint32(v159)+14)) = uint16(v207)
	goto L44
L46:
	;
	goto L49
L47:
	;
	goto L48
L48:
	;
	goto L54
L49:
	;
	v178 = v159 + v160
	v180 = v159 + int32(4)
	if base.Ui32(v180) < base.Ui32(v178) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v182 = v178
	goto L52
L51:
	;
	v182 = v180
	goto L52
L52:
	;
	v187 = (v159^int32(-1)+v182)&int32(-4) + int32(4)
	if v187 == int32(0) {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	base.MemoryFill(m, v159, int32(0), v187)
	goto L45
L54:
	;
	base.MemoryFill(m, v159, int32(0), v160)
	goto L45
L55:
	;
	v216 = v118
	goto L10
}
