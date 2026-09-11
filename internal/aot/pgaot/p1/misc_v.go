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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v38 int32
	_ = v38
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v11+int32(32), int32(2), int32(3), int32(184), v22)
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
	v31 = F_systable_beginscan(m, v15, int32(2187), v26, int32(0), v26, v11+int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v33 = F_systable_getnext(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v33 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v37 = v33
	v38 = v3
	goto L11
L9:
	;
	v64 = v3
	goto L10
L10:
	;
	F_systable_endscan(m, v31)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v48 = F_SearchSysCache1(m, int32(34), v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v64 = v58
	goto L10
L13:
	;
	if v48 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+v53)+18)))
	F_ReleaseCatCache(m, v48)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v58 = v38 + v55
	v59 = F_systable_getnext(m, v31)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v59 != 0 {
		v37 = v59
		v38 = v58
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
L18:
	;
	F_sequence_close(m, v15, int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v75 = F_RelationGetPartitionDesc(m, l1, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v64 != v77 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v81 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v86 = F_SearchSysCacheCopy(m, int32(34), v84, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v86 == int32(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+22)))
	v93 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v90+v91)+18)) = uint8(v93)
	F_CatalogTupleUpdate(m, v81, v86+int32(4), v86)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_sequence_close(m, v81, int32(3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v86)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+131)))
	if v105 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v112 = F_get_partition_parent(m, v110, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v116 = F_get_partition_parent(m, v114, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v119 = F_relation_open(m, v112, int32(8))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v122 = F_relation_open(m, v116, int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_validatePartitionedIndex(m, v119, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_relation_close(m, v119, int32(8))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_relation_close(m, v122, int32(8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
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
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v156 int32
	_ = v156
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var __phi369 int32
	_ = __phi369
	var v371 int32
	_ = v371
	var __phi371 int32
	_ = __phi371
	var v374 int32
	_ = v374
	var __phi374 int32
	_ = __phi374
	var v375 int32
	_ = v375
	var __phi375 int32
	_ = __phi375
	var v379 int32
	_ = v379
	var __phi379 int32
	_ = __phi379
	var v384 int32
	_ = v384
	var __phi384 int32
	_ = __phi384
	var v389 int32
	_ = v389
	var __phi389 int32
	_ = __phi389
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v538 int32
	_ = v538
	var v570 int32
	_ = v570
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v701 int32
	_ = v701
	var v712 int32
	_ = v712
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v820 int32
	_ = v820
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v849 int32
	_ = v849
	var v875 int32
	_ = v875
	var v897 int32
	_ = v897
	var v911 int32
	_ = v911
	v6 = int32(0)
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
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v31 + int32(16)
	return v911
L5:
	;
	v911 = v37
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v911 = v33
	goto L4
L9:
	;
	goto L10
L10:
	;
	v44 = v33 + int32(1)
	if l4 < int32(0) {
		v72 = l4
		v75 = v44
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if base.B2i32(l1 == v33)&base.B2i32(l3 == v37) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	v47 = int32(0)
	v48 = v37 - v33
	if v48 < v47 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v63 = base.I32_div_s(l4-v52, int32(2))
	v69 = v63 - v48>>(uint(int32(31))%32)&v48 + int32(1)
	if v33 < v69 {
		goto L26
	} else {
		goto L27
	}
L14:
	;
	v52 = v47 - v48
	goto L16
L15:
	;
	v52 = v48
	goto L16
L16:
	;
	if v52 <= l4 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v33 < v37 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v911 = l4 + int32(1)
	goto L4
L20:
	;
	v55 = v33
	goto L22
L21:
	;
	v55 = v37
	goto L22
L22:
	;
	if v52+v55 <= l4 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = int32(-1)
	v75 = v44
	goto L11
L24:
	;
	goto L25
L25:
	;
	goto L13
L26:
	;
	v71 = v44
	goto L28
L27:
	;
	v71 = v69
	goto L28
L28:
	;
	v72 = l4
	v75 = v71
	goto L11
L29:
	;
	v83 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v179 = v6
	goto L31
L31:
	;
	v192 = F_palloc(m, v44<<(uint(int32(3))%32))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L40
	}
L32:
	;
	v85 = int32(0)
	if v85 < v33 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = l0
	v95 = v85
	goto L36
L34:
	;
	v156 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v156<<(uint(int32(2))%32)))) = int32(0)
	v179 = v83
	goto L31
L36:
	;
	v120 = F_pg_mblen_range(m, v90, l0+l1)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v156 = v33
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83+v95<<(uint(int32(2))%32)))) = v120
	v125 = v95 + int32(1)
	if v125 != v33 {
		v90 = v90 + v120
		v95 = v125
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v75 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	if v37+int32(1) < int32(2) {
		goto L54
	} else {
		goto L55
	}
L42:
	;
	v199 = v75 & int32(3)
	v200 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v75) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v208 = v200
	v216 = int32(0)
	goto L46
L44:
	;
	v263 = v200
	goto L45
L45:
	;
	if v199 == int32(0) {
		goto L41
	} else {
		goto L49
	}
L46:
	;
	v235 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v208<<(uint(v235)%32)))) = v208
	v240 = v208 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v240<<(uint(v235)%32)))) = v240
	v246 = v208 | v235
	*(*int32)(unsafe.Add(mBase, uint32(v192+v246<<(uint(v235)%32)))) = v246
	v252 = v208 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v192+v252<<(uint(v235)%32)))) = v252
	v257 = int32(4)
	v258 = v208 + v257
	v260 = v216 + v257
	if v260 != v75&int32(2147483644) {
		v208 = v258
		v216 = v260
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v263 = v258
	goto L45
L48:
	;
	goto L47
L49:
	;
	v293 = v263
	v299 = v200
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192+v293<<(uint(int32(2))%32)))) = v293
	v324 = int32(1)
	v327 = v299 + v324
	if v327 != v199 {
		v293 = v293 + v324
		v299 = v327
		goto L50
	} else {
		goto L52
	}
L51:
	;
	goto L41
L52:
	;
	goto L51
L53:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v875+v33<<(uint(int32(2))%32))))
	v911 = v897
	goto L4
L54:
	;
	v875 = v192
	goto L53
L55:
	;
	goto L56
L56:
	;
	v366 = int32(1)
	v367 = v72 + v366
	__phi369 = l0
	__phi371 = l2
	__phi374 = v75
	__phi375 = v192
	__phi379 = v192 + v44<<(uint(int32(2))%32)
	__phi384 = v6
	__phi389 = v366
	v369 = __phi369
	v371 = __phi371
	v374 = __phi374
	v375 = __phi375
	v379 = __phi379
	v384 = __phi384
	v389 = __phi389
	goto L57
L57:
	;
	if l3 != v37 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v875 = v379
	goto L53
L59:
	;
	v399 = F_pg_mblen_range(m, v371, l2+l3)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	v401 = int32(1)
	goto L61
L61:
	;
	if v44 <= v374 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v401 = v399
	goto L61
L63:
	;
	v409 = v374
	goto L65
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375+v374<<(uint(int32(2))%32)))) = v367
	v409 = v374 + int32(1)
	goto L65
L65:
	;
	if v384 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = v389
	v414 = int32(1)
	goto L68
L67:
	;
	v414 = v384
	goto L68
L68:
	;
	if v179 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v72 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L70:
	;
	if v409 <= v414 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v409 <= v414 {
		goto L69
	} else {
		goto L93
	}
L73:
	;
	v424 = v414
	v425 = v369
	goto L74
L74:
	;
	v447 = int32(2)
	v448 = v424 << (uint(v447) % 32)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v375+v448)))
	v451 = int32(1)
	v452 = v450 + v451
	v456 = (v424 - v451) << (uint(v447) % 32)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v379+v456)))
	v460 = v458 + v451
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v456+v179)))
	v463 = v425 + v462
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463-v451))))
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371+v401-int32(1)))))
	if v466 != v467 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L69
L76:
	;
	if v452 < v460 {
		goto L86
	} else {
		goto L87
	}
L77:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v375+v456)))
	v599 = v570
	goto L76
L78:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v375+v456)))
	v599 = v538 + int32(1)
	goto L76
L79:
	;
	if v462 != v401 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	if v401 == int32(1) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v473 = v401
	goto L82
L82:
	;
	if v473 <= int32(0) {
		goto L77
	} else {
		goto L84
	}
L83:
	;
	goto L78
L84:
	;
	v503 = v473 - int32(1)
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v425+v503))))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503+v371))))
	if v505 == v507 {
		v473 = v503
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v602 = v452
	goto L88
L87:
	;
	v602 = v460
	goto L88
L88:
	;
	if v602 < v599 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v604 = v602
	goto L91
L90:
	;
	v604 = v599
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379+v448))) = v604
	v607 = v424 + int32(1)
	if v607 != v409 {
		v424 = v607
		v425 = v463
		goto L74
	} else {
		goto L92
	}
L92:
	;
	goto L75
L93:
	;
	v610 = int32(4)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v379+v414<<(uint(int32(2))%32)-v610)))
	v619 = v369
	v623 = v414
	v626 = v617
	goto L94
L94:
	;
	v647 = v623 << (uint(int32(2)) % 32)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v647+v375)))
	v651 = int32(1)
	v652 = v650 + v651
	v654 = v626 + v651
	if v652 < v654 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L69
L96:
	;
	v656 = v652
	goto L98
L97:
	;
	v656 = v654
	goto L98
L98:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v647+(v375-v610))))
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371))))
	v662 = v658 + base.B2i32(v659 != v660)
	if v656 < v662 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v664 = v656
	goto L101
L100:
	;
	v664 = v662
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379+v647))) = v664
	v666 = int32(1)
	v669 = v623 + v666
	if v669 != v409 {
		v619 = v619 + v666
		v623 = v669
		v626 = v664
		goto L94
	} else {
		goto L102
	}
L102:
	;
	goto L95
L103:
	;
	if v37 != v389 {
		__phi369 = v834
		__phi371 = v371 + v401
		__phi374 = v839
		__phi375 = v379
		__phi379 = v375
		__phi384 = v849
		__phi389 = v389 + int32(1)
		v369 = __phi369
		v371 = __phi371
		v374 = __phi374
		v375 = __phi375
		v379 = __phi379
		v384 = __phi384
		v389 = __phi389
		goto L57
	} else {
		goto L133
	}
L104:
	;
	v834 = v369
	v839 = v409
	v849 = v384
	goto L103
L105:
	;
	goto L106
L106:
	;
	v701 = v389 + (v33 - v37)
	v712 = v409
	goto L107
L107:
	;
	if v712 <= int32(0) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	if v750 <= v384 {
		v805 = v369
		v820 = v384
		goto L117
	} else {
		goto L118
	}
L109:
	;
	goto L108
L110:
	;
	v750 = v409 >> (uint(int32(31)) % 32) & v409
	goto L109
L111:
	;
	goto L112
L112:
	;
	v736 = v712 - int32(1)
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v379+v736<<(uint(int32(2))%32))))
	v741 = v736 - v701
	v742 = int32(0)
	if v742 < v741 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v746 = v741
	goto L115
L114:
	;
	v746 = v742 - v741
	goto L115
L115:
	;
	if v72 < v740+v746 {
		v712 = v736
		goto L107
	} else {
		goto L116
	}
L116:
	;
	v750 = v712
	goto L109
L117:
	;
	if v750 <= v820 {
		v911 = v367
		goto L4
	} else {
		goto L132
	}
L118:
	;
	v753 = v369
	v768 = v384
	goto L119
L119:
	;
	v782 = v768 << (uint(int32(2)) % 32)
	v783 = v379 + v782
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v783)))
	v785 = v768 - v701
	v786 = int32(0)
	if v786 < v785 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v911 = v367
	goto L4
L121:
	;
	v790 = v785
	goto L123
L122:
	;
	v790 = v786 - v785
	goto L123
L123:
	;
	if v784+v790 <= v72 {
		v805 = v753
		v820 = v768
		goto L117
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v783))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v782+v375))) = v367
	if v768 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	if v179 != 0 {
		goto L128
	} else {
		goto L129
	}
L126:
	;
	v801 = v753
	goto L127
L127:
	;
	v803 = v768 + int32(1)
	if v803 != v750 {
		v753 = v801
		v768 = v803
		goto L119
	} else {
		goto L131
	}
L128:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v782+(v179-int32(4)))))
	v799 = v797
	goto L130
L129:
	;
	v799 = int32(1)
	goto L130
L130:
	;
	v801 = v799 + v753
	goto L127
L131:
	;
	goto L120
L132:
	;
	v834 = v805
	v839 = v750
	v849 = v820
	goto L103
L133:
	;
	goto L58
}
func F_visibilitymap_prepare_truncate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = base.I32_div_u_s(l1, int32(_a_F_visibilitymap_prepare_truncate_0))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v24
	v28 = F_smgropen(m, v14+int32(16), v21)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v18
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v28
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+72))
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = v46
	goto L3
L7:
	;
	v42 = v34
	goto L9
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+72))
	v42 = v40
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v42 + int32(1)
	goto L6
L10:
	;
	m.G0 = v14 + int32(32)
	return v224
L11:
	;
	if v51 == int32(0) {
		v224 = v49
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = l1 << (uint(int32(1)) % 32) & int32(6)
	v59 = int32(0)
	v63 = l1 - v17*int32(_a_F_visibilitymap_prepare_truncate_0)
	if base.B2i32(v58 == v59)&base.B2i32(base.Ui32(v63) <= base.Ui32(int32(3))) == v59 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v70 = F_vm_readbuf(m, l0, v17, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v186 = v17
	goto L15
L15:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v193 != 0 {
		goto L50
	} else {
		goto L51
	}
L16:
	;
	if v70 == int32(0) {
		v224 = v49
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v75 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
	if v70 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	F_LockBuffer(m, v70, int32(2))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L22
	}
L19:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[0]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v70^int32(-1))<<(uint(int32(2))%32))))
	v93 = v85
	goto L18
L20:
	;
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[1]))
	v93 = v87 + v70<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	v97 = int32(_a_F_visibilitymap_prepare_truncate_1)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2])) = v99 + int32(1)
	v104 = int32(_a_F_visibilitymap_prepare_truncate_2) - v75
	v105 = v75 + v93
	v107 = v105 + int32(24)
	v109 = v105 + int32(25)
	if v109&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	v140 = int32(-1)
	v144 = v139 & (v140<<(uint(v58)%32) ^ v140)
	*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v144)
	F_MarkBufferDirty(m, v70)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L34
	}
L24:
	;
	v136 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), v104)
	mBase = m.M
	goto L33
L25:
	;
	if v104&int32(3) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v63) < base.Ui32(int32(_a_F_visibilitymap_prepare_truncate_3)) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v104+v109) <= base.Ui32(v109) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v120 = v105 + int32(29)
	v122 = v93 - int32(-8192)
	if base.Ui32(v122) < base.Ui32(v120) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v124 = v120
	goto L31
L30:
	;
	v124 = v122
	goto L31
L31:
	;
	v133 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), (v124-v105-int32(26))&int32(-4)+int32(4))
	mBase = m.M
	goto L32
L32:
	;
	goto L23
L33:
	;
	goto L23
L34:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[3])))
	if v149 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v174 = int32(1)
	v176 = int32(_a_F_visibilitymap_prepare_truncate_1)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2])) = v178 - v174
	F_UnlockReleaseBuffer(m, v70)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L49
	}
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+118)))
	if v151 != int32(112) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[4]))
	if v155 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v158 != 0 {
		goto L35
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[5]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+252))
	goto L43
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v159 != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if base.B2i32(v162 != int32(0)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[6])))
	if v168 != int32(1) {
		goto L35
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_log_newpage_buffer(m, v70, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	goto L35
L49:
	;
	v186 = v17 + v174
	goto L15
L50:
	;
	v217 = v193
	goto L52
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v195
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v197
	v199 = F_smgropen(m, v14, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v219 = F_smgrnblocks(m, v217, int32(2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	if v203 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v217 = v215
	goto L52
L55:
	;
	v211 = v203
	goto L57
L56:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	v211 = v209
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+72)) = v211 + int32(1)
	goto L54
L58:
	;
	if base.Ui32(v219) <= base.Ui32(v186) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v222 = int32(-1)
	goto L61
L60:
	;
	v222 = v186
	goto L61
L61:
	;
	v224 = v222
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
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
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
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
			v248 = m.ExcPending
			if v248 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_set_1), int32(0))
				mBase = m.M
				v252 = m.ExcPending
				if v252 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(270), int32(_a_F_visibilitymap_set_3))
					mBase = m.M
					v257 = m.ExcPending
					if v257 != 0 {
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
				v261 = m.ExcPending
				if v261 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
						mBase = m.M
						v270 = m.ExcPending
						if v270 != 0 {
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
					v261 = m.ExcPending
					if v261 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
						mBase = m.M
						v265 = m.ExcPending
						if v265 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
							mBase = m.M
							v270 = m.ExcPending
							if v270 != 0 {
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
						v89 = v81 + int32(base.Ui32(l1-v12*int32(_a_F_visibilitymap_set_0))>>(uint(int32(2))%32)) + int32(24)
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v93 = int32(base.Ui32(v90)>>(uint(v58)%32)) & int32(3)
						if v7 != v93 {
							v95 = int32(_a_F_visibilitymap_set_5)
							v97 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
							*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v97 + int32(1)
							v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
							v103 = v101 | v7<<(uint(v58)%32)
							*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v103)
							F_MarkBufferDirty(m, l4)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+118)))
								if v108 != int32(112) {
									v231 = int32(_a_F_visibilitymap_set_5)
									v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
									*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
									F_LockBuffer(m, l4, int32(0))
									mBase = m.M
									v243 = m.ExcPending
									if v243 != 0 {
										return int32(0)
									} else {
										return v93
									}
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
									if v112 <= int32(0) {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v115 != 0 {
											v231 = int32(_a_F_visibilitymap_set_5)
											v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v243 = m.ExcPending
											if v243 != 0 {
												return int32(0)
											} else {
												return v93
											}
										} else {
											v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v116 != 0 {
												v231 = int32(_a_F_visibilitymap_set_5)
												v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
												*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
												F_LockBuffer(m, l4, int32(0))
												mBase = m.M
												v243 = m.ExcPending
												if v243 != 0 {
													return int32(0)
												} else {
													return v93
												}
											} else {
												if l3 != int64(0) {
													v220 = l3
													*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
													v231 = int32(_a_F_visibilitymap_set_5)
													v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
													*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
													F_LockBuffer(m, l4, int32(0))
													mBase = m.M
													v243 = m.ExcPending
													if v243 != 0 {
														return int32(0)
													} else {
														return v93
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
																			if v196 != int32(1) {
																				v220 = v183
																			} else {
																				if l2 < int32(0) {
																					v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																					v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																					v216 = v208
																				} else {
																					v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																					v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																				}
																				*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																				v220 = v183
																			}
																		} else {
																			if l2 < int32(0) {
																				v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v216 = v208
																			} else {
																				v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																				v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																			v220 = v183
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																		v231 = int32(_a_F_visibilitymap_set_5)
																		v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																		*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
																		F_LockBuffer(m, l4, int32(0))
																		mBase = m.M
																		v243 = m.ExcPending
																		if v243 != 0 {
																			return int32(0)
																		} else {
																			return v93
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
											v220 = l3
											*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
											v231 = int32(_a_F_visibilitymap_set_5)
											v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v243 = m.ExcPending
											if v243 != 0 {
												return int32(0)
											} else {
												return v93
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
																	if v196 != int32(1) {
																		v220 = v183
																	} else {
																		if l2 < int32(0) {
																			v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v216 = v208
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																			v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																		v220 = v183
																	}
																} else {
																	if l2 < int32(0) {
																		v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																		v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v216 = v208
																	} else {
																		v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																		v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																	v220 = v183
																}
																*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																v231 = int32(_a_F_visibilitymap_set_5)
																v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
																F_LockBuffer(m, l4, int32(0))
																mBase = m.M
																v243 = m.ExcPending
																if v243 != 0 {
																	return int32(0)
																} else {
																	return v93
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
							v243 = m.ExcPending
							if v243 != 0 {
								return int32(0)
							} else {
								return v93
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
			v261 = m.ExcPending
			if v261 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
					mBase = m.M
					v270 = m.ExcPending
					if v270 != 0 {
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
				v261 = m.ExcPending
				if v261 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_visibilitymap_set_4), int32(0))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(274), int32(_a_F_visibilitymap_set_3))
						mBase = m.M
						v270 = m.ExcPending
						if v270 != 0 {
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
					v89 = v81 + int32(base.Ui32(l1-v12*int32(_a_F_visibilitymap_set_0))>>(uint(int32(2))%32)) + int32(24)
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
					v93 = int32(base.Ui32(v90)>>(uint(v58)%32)) & int32(3)
					if v7 != v93 {
						v95 = int32(_a_F_visibilitymap_set_5)
						v97 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
						*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v97 + int32(1)
						v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v103 = v101 | v7<<(uint(v58)%32)
						*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v103)
						F_MarkBufferDirty(m, l4)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+118)))
							if v108 != int32(112) {
								v231 = int32(_a_F_visibilitymap_set_5)
								v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
								*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
								F_LockBuffer(m, l4, int32(0))
								mBase = m.M
								v243 = m.ExcPending
								if v243 != 0 {
									return int32(0)
								} else {
									return v93
								}
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[5]))
								if v112 <= int32(0) {
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v115 != 0 {
										v231 = int32(_a_F_visibilitymap_set_5)
										v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
										*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
										F_LockBuffer(m, l4, int32(0))
										mBase = m.M
										v243 = m.ExcPending
										if v243 != 0 {
											return int32(0)
										} else {
											return v93
										}
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v116 != 0 {
											v231 = int32(_a_F_visibilitymap_set_5)
											v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
											*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
											F_LockBuffer(m, l4, int32(0))
											mBase = m.M
											v243 = m.ExcPending
											if v243 != 0 {
												return int32(0)
											} else {
												return v93
											}
										} else {
											if l3 != int64(0) {
												v220 = l3
												*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
												v231 = int32(_a_F_visibilitymap_set_5)
												v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
												*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
												F_LockBuffer(m, l4, int32(0))
												mBase = m.M
												v243 = m.ExcPending
												if v243 != 0 {
													return int32(0)
												} else {
													return v93
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
																		if v196 != int32(1) {
																			v220 = v183
																		} else {
																			if l2 < int32(0) {
																				v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v216 = v208
																			} else {
																				v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																				v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																			v220 = v183
																		}
																	} else {
																		if l2 < int32(0) {
																			v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v216 = v208
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																			v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																		v220 = v183
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																	v231 = int32(_a_F_visibilitymap_set_5)
																	v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
																	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
																	F_LockBuffer(m, l4, int32(0))
																	mBase = m.M
																	v243 = m.ExcPending
																	if v243 != 0 {
																		return int32(0)
																	} else {
																		return v93
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
										v220 = l3
										*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
										v231 = int32(_a_F_visibilitymap_set_5)
										v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
										*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
										F_LockBuffer(m, l4, int32(0))
										mBase = m.M
										v243 = m.ExcPending
										if v243 != 0 {
											return int32(0)
										} else {
											return v93
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
																if v196 != int32(1) {
																	v220 = v183
																} else {
																	if l2 < int32(0) {
																		v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																		v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v216 = v208
																	} else {
																		v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																		v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																	v220 = v183
																}
															} else {
																if l2 < int32(0) {
																	v202 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																	v216 = v208
																} else {
																	v210 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
																	v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																}
																*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																v220 = v183
															}
															*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
															v231 = int32(_a_F_visibilitymap_set_5)
															v233 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4]))
															*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[4])) = v233 - int32(1)
															F_LockBuffer(m, l4, int32(0))
															mBase = m.M
															v243 = m.ExcPending
															if v243 != 0 {
																return int32(0)
															} else {
																return v93
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
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							return v93
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
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int64
	_ = v111
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9-int32(-64)))) = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = v19
	v23 = F_smgropen(m, v9+int32(56), v14)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v11
	goto L3
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	if v43 == int32(-1) {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = v41
	goto L3
L7:
	;
	v37 = v29
	goto L9
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	v37 = v35
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v37 + int32(1)
	goto L6
L10:
	;
	m.G0 = v9 + int32(80)
	return v209
L11:
	;
	if v120 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v64 = int32(0)
	if l2 == v64 {
		v209 = v64
		goto L10
	} else {
		goto L22
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+28)) = int32(0)
	goto L12
L14:
	;
	v47 = F_smgrexists(m, v42, int32(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v55 = v43
	goto L16
L16:
	;
	if base.Ui32(v55) <= base.Ui32(l1) {
		goto L12
	} else {
		goto L20
	}
L17:
	;
	if v47 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v52 = F_smgrnblocks(m, v42, int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	v55 = v54
	goto L16
L20:
	;
	v60 = F_ReadBufferExtended(m, l0, int32(2), l1, int32(3), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v120 = v60
	goto L11
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = int64(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v9)+68)) = l0
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v9)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v72
	v81 = F_ExtendBufferedRelTo(m, v9+int32(40), int32(2), int32(20), l1+int32(1), int32(3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v83 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v89
	v93 = F_smgropen(m, v9+int32(24), v86)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	v110 = v83
	goto L26
L26:
	;
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v111
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v113
	F_CacheInvalidateSmgr(m, v9+int32(8))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L32
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v93
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+72))
	if v97 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v110 = v109
	goto L26
L29:
	;
	v105 = v97
	goto L31
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+76))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v93)+72))
	v105 = v103
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+72)) = v105 + int32(1)
	goto L28
L32:
	;
	v120 = v81
	goto L11
L33:
	;
	F_LockBuffer(m, v120, int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L54
	}
L34:
	;
	if v161&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L35:
	;
	v126 = (v120 ^ int32(-1)) << (uint(int32(2)) % 32)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v128)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+14)))
	if v131 != 0 {
		v209 = v120
		goto L10
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v143 = v120 << (uint(int32(13)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[1]))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+v145-int32(_a_F_vm_readbuf_3)))))
	if v149 != 0 {
		v209 = v120
		goto L10
	} else {
		goto L41
	}
L38:
	;
	F_LockBuffer(m, v120, int32(2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[0]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136+v126)))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+14)))
	if v139 == int32(0) {
		v161 = v138
		goto L34
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	F_LockBuffer(m, v120, int32(2))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_vm_readbuf[1]))
	v155 = v154 + v143
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155-int32(_a_F_vm_readbuf_3)))))
	if v158 != 0 {
		goto L33
	} else {
		goto L43
	}
L43:
	;
	v161 = v155 + int32(-8192)
	goto L34
L44:
	;
	goto L33
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+10)) = int32(_a_F_vm_readbuf_1)
	v194 = int32(_a_F_vm_readbuf_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)) = uint16(v194)
	v200 = int32(_a_F_vm_readbuf_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+16)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+14)) = uint16(v200)
	goto L44
L46:
	;
	v188 = F___memset(m, v161, int32(0), int32(_a_F_vm_readbuf_0))
	mBase = m.M
	goto L45
L47:
	;
	goto L46
L54:
	;
	v209 = v120
	goto L10
}
