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
	F_errmsg_internal(m, int32(40143), v11+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(494763), int32(21849), int32(29413))
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
	F_errmsg_internal(m, int32(40143), v11)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(494763), int32(21875), int32(29413))
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
		v8 = F_anychar_typmodin(m, v3, int32(229998))
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v181 int32
	_ = v181
	var v206 int32
	_ = v206
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var __phi414 int32
	_ = __phi414
	var v416 int32
	_ = v416
	var __phi416 int32
	_ = __phi416
	var v422 int32
	_ = v422
	var __phi422 int32
	_ = __phi422
	var v423 int32
	_ = v423
	var __phi423 int32
	_ = __phi423
	var v427 int32
	_ = v427
	var __phi427 int32
	_ = __phi427
	var v432 int32
	_ = v432
	var __phi432 int32
	_ = __phi432
	var v437 int32
	_ = v437
	var __phi437 int32
	_ = __phi437
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v594 int32
	_ = v594
	var v628 int32
	_ = v628
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v768 int32
	_ = v768
	var v782 int32
	_ = v782
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v843 int32
	_ = v843
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v900 int32
	_ = v900
	var v914 int32
	_ = v914
	var v922 int32
	_ = v922
	var v932 int32
	_ = v932
	var v961 int32
	_ = v961
	var v983 int32
	_ = v983
	var v1000 int32
	_ = v1000
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	v10 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(16)
	m.G0 = v34
	v36 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v40 = F_pg_mbstrlen_with_len(m, l2, l3)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v36 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L147
	}
L5:
	;
	m.G0 = v34 + int32(16)
	return v1000
L6:
	;
	v1000 = l4 * v40
	goto L5
L7:
	;
	goto L8
L8:
	;
	if v40 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v1000 = l5 * v36
	goto L5
L10:
	;
	goto L11
L11:
	;
	if l8 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if int32(255) < v36 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = v36 + int32(1)
	if l7 < int32(0) {
		v90 = l6
		v91 = l7
		v95 = v55
		goto L17
	} else {
		goto L18
	}
L15:
	;
	if int32(256) <= v40 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if base.B2i32(l1 == v36)&base.B2i32(l3 == v40) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L18:
	;
	v58 = int32(0)
	v59 = v40 - v36
	if v59 < v58 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v81 = base.I32_div_s(l7-v65, v67)
	v87 = v81 - v59>>(uint(int32(31))%32)&v59 + int32(1)
	if v36 < v87 {
		goto L36
	} else {
		goto L37
	}
L20:
	;
	v65 = v58 - l5*v59
	goto L22
L21:
	;
	v65 = l4 * v59
	goto L22
L22:
	;
	if v65 <= l7 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v67 = l4 + l5
	if v67 < l6 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v1000 = l7 + int32(1)
	goto L5
L26:
	;
	v69 = v67
	goto L28
L27:
	;
	v69 = l6
	goto L28
L28:
	;
	if v36 < v40 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v71 = v36
	goto L31
L30:
	;
	v71 = v40
	goto L31
L31:
	;
	if v65+v69*v71 <= l7 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v90 = v69
	v91 = int32(-1)
	v95 = v55
	goto L17
L33:
	;
	goto L34
L34:
	;
	if int32(0) < v67 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	v90 = v69
	v91 = l7
	v95 = v55
	goto L17
L36:
	;
	v89 = v55
	goto L38
L37:
	;
	v89 = v87
	goto L38
L38:
	;
	v90 = v69
	v91 = l7
	v95 = v89
	goto L17
L39:
	;
	v103 = F_palloc(m, v55<<(uint(int32(2))%32))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v206 = v10
	goto L41
L41:
	;
	v220 = F_palloc(m, v55<<(uint(int32(3))%32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L50
	}
L42:
	;
	v105 = int32(0)
	if v105 < v36 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v110 = l0
	v118 = v105
	goto L46
L44:
	;
	v181 = v105
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103+v181<<(uint(int32(2))%32)))) = int32(0)
	v206 = v103
	goto L41
L46:
	;
	v143 = F_pg_mblen_range(m, v110, l0+l1)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v181 = v36
	goto L45
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103+v118<<(uint(int32(2))%32)))) = v143
	v148 = v118 + int32(1)
	if v148 != v36 {
		v110 = v110 + v143
		v118 = v148
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	if v95 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v40+int32(1) < int32(2) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v227 = v95 & int32(3)
	v228 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v95) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v236 = v228
	v247 = int32(0)
	goto L56
L54:
	;
	v298 = v228
	goto L55
L55:
	;
	if v227 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v266 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v220+v236<<(uint(v266)%32)))) = v236 * l5
	v272 = v236 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v220+v272<<(uint(v266)%32)))) = l5 * v272
	v279 = v236 | v266
	*(*int32)(unsafe.Add(mBase, uint32(v220+v279<<(uint(v266)%32)))) = l5 * v279
	v286 = v236 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v220+v286<<(uint(v266)%32)))) = l5 * v286
	v292 = int32(4)
	v293 = v236 + v292
	v295 = v247 + v292
	if v295 != v95&int32(2147483644) {
		v236 = v293
		v247 = v295
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v298 = v293
	goto L55
L58:
	;
	goto L57
L59:
	;
	v331 = v298
	v340 = v228
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220+v331<<(uint(int32(2))%32)))) = v331 * l5
	v366 = int32(1)
	v369 = v340 + v366
	if v369 != v227 {
		v331 = v331 + v366
		v340 = v369
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v961+v36<<(uint(int32(2))%32))))
	v1000 = v983
	goto L5
L64:
	;
	v961 = v220
	goto L63
L65:
	;
	goto L66
L66:
	;
	v411 = int32(1)
	v412 = v91 + v411
	__phi414 = l0
	__phi416 = l2
	__phi422 = v95
	__phi423 = v220
	__phi427 = v220 + v55<<(uint(int32(2))%32)
	__phi432 = v10
	__phi437 = v411
	v414 = __phi414
	v416 = __phi416
	v422 = __phi422
	v423 = __phi423
	v427 = __phi427
	v432 = __phi432
	v437 = __phi437
	goto L67
L67:
	;
	if l3 != v40 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v961 = v427
	goto L63
L69:
	;
	v447 = F_pg_mblen_range(m, v416, l2+l3)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	v449 = int32(1)
	goto L71
L71:
	;
	if v55 <= v422 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v449 = v447
	goto L71
L73:
	;
	v457 = v422
	goto L75
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v423+v422<<(uint(int32(2))%32)))) = v412
	v457 = v422 + int32(1)
	goto L75
L75:
	;
	if v432 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = l4 * v437
	v463 = int32(1)
	goto L78
L77:
	;
	v463 = v432
	goto L78
L78:
	;
	if v206 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v91 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L80:
	;
	if v457 <= v463 {
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v457 <= v463 {
		goto L79
	} else {
		goto L103
	}
L83:
	;
	v476 = v463
	v477 = v414
	goto L84
L84:
	;
	v499 = int32(2)
	v500 = v476 << (uint(v499) % 32)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v423+v500)))
	v503 = v502 + l4
	v504 = int32(1)
	v507 = (v476 - v504) << (uint(v499) % 32)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v427+v507)))
	v510 = v509 + l5
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v206+v507)))
	v513 = v477 + v512
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513-v504))))
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416+v449-int32(1)))))
	if v516 != v517 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	goto L79
L86:
	;
	if v503 < v510 {
		goto L96
	} else {
		goto L97
	}
L87:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v423+v507)))
	v660 = v628
	goto L86
L88:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v423+v507)))
	v660 = v594 + v90
	goto L86
L89:
	;
	if v512 != v449 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	if v449 == int32(1) {
		goto L87
	} else {
		goto L91
	}
L91:
	;
	v523 = v449
	goto L92
L92:
	;
	if v523 <= int32(0) {
		goto L87
	} else {
		goto L94
	}
L93:
	;
	goto L88
L94:
	;
	v556 = v523 - int32(1)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v556))))
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556+v416))))
	if v558 == v560 {
		v523 = v556
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v663 = v503
	goto L98
L97:
	;
	v663 = v510
	goto L98
L98:
	;
	if v663 < v660 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v665 = v663
	goto L101
L100:
	;
	v665 = v660
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427+v500))) = v665
	v668 = v476 + int32(1)
	if v668 != v457 {
		v476 = v668
		v477 = v513
		goto L84
	} else {
		goto L102
	}
L102:
	;
	goto L85
L103:
	;
	v671 = int32(4)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v427+v463<<(uint(int32(2))%32)-v671)))
	v680 = v414
	v687 = v463
	v690 = v678
	goto L104
L104:
	;
	v711 = v687 << (uint(int32(2)) % 32)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v711+v423)))
	v715 = v714 + l4
	v716 = l5 + v690
	if v715 < v716 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L79
L106:
	;
	v718 = v715
	goto L108
L107:
	;
	v718 = v716
	goto L108
L108:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v711+(v423-v671))))
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v722 != v723 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v725 = v90
	goto L111
L110:
	;
	v725 = int32(0)
	goto L111
L111:
	;
	v726 = v720 + v725
	if v718 < v726 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v728 = v718
	goto L114
L113:
	;
	v728 = v726
	goto L114
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427+v711))) = v728
	v730 = int32(1)
	v733 = v687 + v730
	if v733 != v457 {
		v680 = v680 + v730
		v687 = v733
		v690 = v728
		goto L104
	} else {
		goto L115
	}
L115:
	;
	goto L105
L116:
	;
	if v40 != v437 {
		__phi414 = v914
		__phi416 = v416 + v449
		__phi422 = v922
		__phi423 = v427
		__phi427 = v423
		__phi432 = v932
		__phi437 = v437 + int32(1)
		v414 = __phi414
		v416 = __phi416
		v422 = __phi422
		v423 = __phi423
		v427 = __phi427
		v432 = __phi432
		v437 = __phi437
		goto L67
	} else {
		goto L146
	}
L117:
	;
	v914 = v414
	v922 = v457
	v932 = v432
	goto L116
L118:
	;
	goto L119
L119:
	;
	v768 = v437 + (v36 - v40)
	v782 = v457
	goto L120
L120:
	;
	if v782 <= int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	if v822 <= v432 {
		v882 = v414
		v900 = v432
		goto L130
	} else {
		goto L131
	}
L122:
	;
	goto L121
L123:
	;
	v822 = v457 >> (uint(int32(31)) % 32) & v457
	goto L122
L124:
	;
	goto L125
L125:
	;
	v806 = v782 - int32(1)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v427+v806<<(uint(int32(2))%32))))
	v811 = v806 - v768
	v813 = int32(0)
	if v813 < v811 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v818 = v811 * l4
	goto L128
L127:
	;
	v818 = v813 - v811*l5
	goto L128
L128:
	;
	if v91 < v810+v818 {
		v782 = v806
		goto L120
	} else {
		goto L129
	}
L129:
	;
	v822 = v782
	goto L122
L130:
	;
	if v822 <= v900 {
		v1000 = v412
		goto L5
	} else {
		goto L145
	}
L131:
	;
	v825 = v414
	v843 = v432
	goto L132
L132:
	;
	v857 = v843 << (uint(int32(2)) % 32)
	v858 = v427 + v857
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)))
	v860 = v843 - v768
	v862 = int32(0)
	if v862 < v860 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v1000 = v412
	goto L5
L134:
	;
	v867 = v860 * l4
	goto L136
L135:
	;
	v867 = v862 - v860*l5
	goto L136
L136:
	;
	if v859+v867 <= v91 {
		v882 = v825
		v900 = v843
		goto L130
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v858))) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v857+v423))) = v412
	if v843 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if v206 != 0 {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v878 = v825
	goto L140
L140:
	;
	v880 = v843 + int32(1)
	if v880 != v822 {
		v825 = v878
		v843 = v880
		goto L132
	} else {
		goto L144
	}
L141:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v857+(v206-int32(4)))))
	v876 = v874
	goto L143
L142:
	;
	v876 = int32(1)
	goto L143
L143:
	;
	v878 = v876 + v825
	goto L140
L144:
	;
	goto L133
L145:
	;
	v914 = v882
	v922 = v822
	v932 = v900
	goto L116
L146:
	;
	goto L68
L147:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(255)
	F_errmsg(m, int32(133454), v34)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(496902), int32(135), int32(310004))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	v17 = base.I32_div_u_s(l1, int32(32672))
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
	v63 = l1 - v17*int32(32672)
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
	v79 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v70^int32(-1))<<(uint(int32(2))%32))))
	v93 = v85
	goto L18
L20:
	;
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v93 = v87 + v70<<(uint(int32(13))%32) + int32(-8192)
	goto L18
L22:
	;
	v97 = int32(4510260)
	v99 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v99 + int32(1)
	v104 = int32(8167) - v75
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
	if base.Ui32(v63) < base.Ui32(int32(28572)) {
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
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[113])))
	if v149 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v174 = int32(1)
	v176 = int32(4510260)
	v178 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	*(*int32)(unsafe.Add(mBase, _consts[11])) = v178 - v174
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
	v155 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v161 = *(*int32)(unsafe.Add(mBase, _consts[109]))
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
	v168 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
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
	v12 = base.I32_div_u_s(l1, int32(32672))
	if l2 != 0 {
		if l2 < int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[8]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l2^int32(-1))<<(uint(int32(6))%32))+16))
			v31 = v22
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
				F_errmsg_internal(m, int32(106189), int32(0))
				mBase = m.M
				v252 = m.ExcPending
				if v252 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496214), int32(270), int32(106261))
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
					F_errmsg_internal(m, int32(106235), int32(0))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496214), int32(274), int32(106261))
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
					v38 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(l4^int32(-1))<<(uint(int32(6))%32))+16))
					v53 = v44
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
						F_errmsg_internal(m, int32(106235), int32(0))
						mBase = m.M
						v265 = m.ExcPending
						if v265 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(496214), int32(274), int32(106261))
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
						v67 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(l4^int32(-1))<<(uint(int32(2))%32))))
						v81 = v73
					} else {
						v75 = *(*int32)(unsafe.Add(mBase, _consts[1]))
						v81 = v75 + l4<<(uint(int32(13))%32) + int32(-8192)
					}
					F_LockBuffer(m, l4, int32(2))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v89 = v81 + int32(base.Ui32(l1-v12*int32(32672))>>(uint(int32(2))%32)) + int32(24)
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v93 = int32(base.Ui32(v90)>>(uint(v58)%32)) & int32(3)
						if v7 != v93 {
							v95 = int32(4510260)
							v97 = *(*int32)(unsafe.Add(mBase, _consts[11]))
							*(*int32)(unsafe.Add(mBase, _consts[11])) = v97 + int32(1)
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
									v231 = int32(4510260)
									v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
									*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
									F_LockBuffer(m, l4, int32(0))
									mBase = m.M
									v243 = m.ExcPending
									if v243 != 0 {
										return int32(0)
									} else {
										return v93
									}
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
									if v112 <= int32(0) {
										v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										if v115 != 0 {
											v231 = int32(4510260)
											v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
											*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
												v231 = int32(4510260)
												v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
												*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
													v231 = int32(4510260)
													v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
													*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
													v126 = *(*int32)(unsafe.Add(mBase, _consts[2]))
													if v126 < int32(2) {
													} else {
														v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
														v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
														if v130 != int32(112) {
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															if base.B2i32(base.Ui32(v133) < base.Ui32(int32(12000))) == int32(0) {
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
																v165 = *(*int32)(unsafe.Add(mBase, _consts[109]))
																v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
																v169 = int32(1)
																v170 = int32(8)
																v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
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
																		v189 = *(*int32)(unsafe.Add(mBase, _consts[109]))
																		v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																		if base.B2i32(v190 != int32(0)) == int32(0) {
																			v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
																			if v196 != int32(1) {
																				v220 = v183
																			} else {
																				if l2 < int32(0) {
																					v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																					v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																					v216 = v208
																				} else {
																					v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																					v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																				}
																				*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																				v220 = v183
																			}
																		} else {
																			if l2 < int32(0) {
																				v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v216 = v208
																			} else {
																				v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																				v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																			v220 = v183
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																		v231 = int32(4510260)
																		v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
																		*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
											v231 = int32(4510260)
											v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
											*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
											v126 = *(*int32)(unsafe.Add(mBase, _consts[2]))
											if v126 < int32(2) {
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
												v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
												if v130 != int32(112) {
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if base.B2i32(base.Ui32(v133) < base.Ui32(int32(12000))) == int32(0) {
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
														v165 = *(*int32)(unsafe.Add(mBase, _consts[109]))
														v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
														v169 = int32(1)
														v170 = int32(8)
														v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
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
																v189 = *(*int32)(unsafe.Add(mBase, _consts[109]))
																v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																if base.B2i32(v190 != int32(0)) == int32(0) {
																	v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
																	if v196 != int32(1) {
																		v220 = v183
																	} else {
																		if l2 < int32(0) {
																			v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v216 = v208
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																			v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																		v220 = v183
																	}
																} else {
																	if l2 < int32(0) {
																		v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																		v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v216 = v208
																	} else {
																		v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																		v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																	v220 = v183
																}
																*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																v231 = int32(4510260)
																v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
																*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
				F_errmsg_internal(m, int32(106235), int32(0))
				mBase = m.M
				v265 = m.ExcPending
				if v265 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496214), int32(274), int32(106261))
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
				v38 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(l4^int32(-1))<<(uint(int32(6))%32))+16))
				v53 = v44
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
					F_errmsg_internal(m, int32(106235), int32(0))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(496214), int32(274), int32(106261))
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
					v67 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(l4^int32(-1))<<(uint(int32(2))%32))))
					v81 = v73
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, _consts[1]))
					v81 = v75 + l4<<(uint(int32(13))%32) + int32(-8192)
				}
				F_LockBuffer(m, l4, int32(2))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					v89 = v81 + int32(base.Ui32(l1-v12*int32(32672))>>(uint(int32(2))%32)) + int32(24)
					v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
					v93 = int32(base.Ui32(v90)>>(uint(v58)%32)) & int32(3)
					if v7 != v93 {
						v95 = int32(4510260)
						v97 = *(*int32)(unsafe.Add(mBase, _consts[11]))
						*(*int32)(unsafe.Add(mBase, _consts[11])) = v97 + int32(1)
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
								v231 = int32(4510260)
								v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
								*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
								F_LockBuffer(m, l4, int32(0))
								mBase = m.M
								v243 = m.ExcPending
								if v243 != 0 {
									return int32(0)
								} else {
									return v93
								}
							} else {
								v112 = *(*int32)(unsafe.Add(mBase, _consts[2]))
								if v112 <= int32(0) {
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									if v115 != 0 {
										v231 = int32(4510260)
										v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
										*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
											v231 = int32(4510260)
											v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
											*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
												v231 = int32(4510260)
												v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
												*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
												v126 = *(*int32)(unsafe.Add(mBase, _consts[2]))
												if v126 < int32(2) {
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
													v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
													if v130 != int32(112) {
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if base.B2i32(base.Ui32(v133) < base.Ui32(int32(12000))) == int32(0) {
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
															v165 = *(*int32)(unsafe.Add(mBase, _consts[109]))
															v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
															v169 = int32(1)
															v170 = int32(8)
															v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
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
																	v189 = *(*int32)(unsafe.Add(mBase, _consts[109]))
																	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
																	if base.B2i32(v190 != int32(0)) == int32(0) {
																		v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
																		if v196 != int32(1) {
																			v220 = v183
																		} else {
																			if l2 < int32(0) {
																				v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																				v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																				v216 = v208
																			} else {
																				v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																				v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																			v220 = v183
																		}
																	} else {
																		if l2 < int32(0) {
																			v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																			v216 = v208
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																			v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																		v220 = v183
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
																	v231 = int32(4510260)
																	v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
																	*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
										v231 = int32(4510260)
										v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
										*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
										v126 = *(*int32)(unsafe.Add(mBase, _consts[2]))
										if v126 < int32(2) {
										} else {
											v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+118)))
											if v130 != int32(112) {
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if base.B2i32(base.Ui32(v133) < base.Ui32(int32(12000))) == int32(0) {
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
													v165 = *(*int32)(unsafe.Add(mBase, _consts[109]))
													v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+252))
													v169 = int32(1)
													v170 = int32(8)
													v174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
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
															v189 = *(*int32)(unsafe.Add(mBase, _consts[109]))
															v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+252))
															if base.B2i32(v190 != int32(0)) == int32(0) {
																v196 = int32(*(*uint8)(unsafe.Add(mBase, _consts[110])))
																if v196 != int32(1) {
																	v220 = v183
																} else {
																	if l2 < int32(0) {
																		v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																		v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																		v216 = v208
																	} else {
																		v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																		v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																	v220 = v183
																}
															} else {
																if l2 < int32(0) {
																	v202 = *(*int32)(unsafe.Add(mBase, _consts[0]))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v202+(l2^int32(-1))<<(uint(int32(2))%32))))
																	v216 = v208
																} else {
																	v210 = *(*int32)(unsafe.Add(mBase, _consts[1]))
																	v216 = v210 + l2<<(uint(int32(13))%32) + int32(-8192)
																}
																*(*int64)(unsafe.Add(mBase, uint32(v216))) = base.I64_rotr(v183, int64(32))
																v220 = v183
															}
															*(*int64)(unsafe.Add(mBase, uint32(v81))) = base.I64_rotr(v220, int64(32))
															v231 = int32(4510260)
															v233 = *(*int32)(unsafe.Add(mBase, _consts[11]))
															*(*int32)(unsafe.Add(mBase, _consts[11])) = v233 - int32(1)
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
	v128 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v145 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143+v145-int32(8178)))))
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
	v136 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v155 = v154 + v143
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155-int32(8178)))))
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
	*(*int32)(unsafe.Add(mBase, uint32(v161)+10)) = int32(1572864)
	v194 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+18)) = uint16(v194)
	v200 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+16)) = uint16(v200)
	*(*uint16)(unsafe.Add(mBase, uint32(v161)+14)) = uint16(v200)
	goto L44
L46:
	;
	v188 = F___memset(m, v161, int32(0), int32(8192))
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
