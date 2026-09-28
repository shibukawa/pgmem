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
	var v22 int64
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
	var v45 int64
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
	var v82 int64
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
	v18 = v11 + int32(24)
	v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
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
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v44))))
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
	v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v84 = F_SearchSysCacheCopy(m, int32(34), v82, int64(0))
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
func F_varcharrecv(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v6+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
		v21 = F_varchar_input(m, v15, v19, v8, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_pfree(m, v15)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return base.I64_extend_i32_u(v21)
			}
		}
	}
}
func F_varchartypmodin(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_anychar_typmodin(m, v3, int32(_a_F_varchartypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_varstr_levenshtein_less_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v214 int32
	_ = v214
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v277 int64
	_ = v277
	var v285 int64
	_ = v285
	var v301 int32
	_ = v301
	var v307 int64
	_ = v307
	var v315 int64
	_ = v315
	var v323 int64
	_ = v323
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v346 int64
	_ = v346
	var v378 int64
	_ = v378
	var v385 int64
	_ = v385
	var v407 int64
	_ = v407
	var v410 int64
	_ = v410
	var v451 int64
	_ = v451
	var v452 int64
	_ = v452
	var v453 int64
	_ = v453
	var v456 int32
	_ = v456
	var __phi456 int32
	_ = __phi456
	var v457 int32
	_ = v457
	var __phi457 int32
	_ = __phi457
	var v458 int32
	_ = v458
	var __phi458 int32
	_ = __phi458
	var v464 int32
	_ = v464
	var __phi464 int32
	_ = __phi464
	var v474 int64
	_ = v474
	var __phi474 int64
	_ = __phi474
	var v475 int32
	_ = v475
	var __phi475 int32
	_ = __phi475
	var v481 int32
	_ = v481
	var __phi481 int32
	_ = __phi481
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int64
	_ = v546
	var v547 int64
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v554 int64
	_ = v554
	var v555 int64
	_ = v555
	var v557 int64
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v648 int64
	_ = v648
	var v683 int64
	_ = v683
	var v716 int64
	_ = v716
	var v718 int64
	_ = v718
	var v721 int32
	_ = v721
	var v729 int64
	_ = v729
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v739 int64
	_ = v739
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v767 int64
	_ = v767
	var v768 int64
	_ = v768
	var v770 int64
	_ = v770
	var v773 int64
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int64
	_ = v778
	var v779 int64
	_ = v779
	var v781 int64
	_ = v781
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v826 int32
	_ = v826
	var v837 int64
	_ = v837
	var v863 int64
	_ = v863
	var v864 int32
	_ = v864
	var v868 int64
	_ = v868
	var v869 int32
	_ = v869
	var v878 int64
	_ = v878
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v897 int64
	_ = v897
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int64
	_ = v924
	var v925 int32
	_ = v925
	var v934 int64
	_ = v934
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int64
	_ = v953
	var v955 int32
	_ = v955
	var v980 int32
	_ = v980
	var v1024 int32
	_ = v1024
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1058 int64
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1095 int64
	_ = v1095
	var v1133 int32
	_ = v1133
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	v10 = int64(0)
	v20 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(16)
	m.G0 = v35
	v37 = base.I64_extend_i32_s(l4)
	v38 = F_pg_mbstrlen_with_len(m, l0, l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v42 = F_pg_mbstrlen_with_len(m, l2, l3)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v38 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L154
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L150
	}
L6:
	;
	m.G0 = v35 + int32(16)
	return v1133
L7:
	;
	v47 = base.I64_extend_i32_s(v42) * v37
	if base.Ui64(v47-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v53 = base.I64_extend_i32_s(l5)
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v1133 = base.I32_wrap_i64(v47)
	goto L6
L11:
	;
	v57 = base.I64_extend_i32_s(v38) * v53
	if base.Ui64(v57-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.B2i32(l8 == int32(0))&(base.B2i32(int32(255) < v38)|base.B2i32(int32(256) <= v42)) != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	v1133 = base.I32_wrap_i64(v57)
	goto L6
L15:
	;
	v71 = base.I64_extend_i32_s(l6)
	v73 = v38 + int32(1)
	if l7 < int32(0) {
		v120 = l7
		v124 = v71
		v125 = v73
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.B2i32(l1 == v38)&base.B2i32(l3 == v42) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L17:
	;
	v76 = base.I64_extend_i32_u(l7)
	v77 = v42 - v38
	if v77 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v111 = base.I64_div_s(v76-v86, v90)
	v114 = base.I64_extend_i32_s((int32(0)-v77)&(v77>>(uint(int32(31))%32))) + v111 + int64(1)
	v115 = base.I64_extend_i32_s(v73)
	if v114 < v115 {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	v1133 = l7 + int32(1)
	goto L6
L20:
	;
	if v76 < v86 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v86 = base.I64_extend_i32_s(int32(0)-v77) * v53
	goto L20
L22:
	;
	goto L23
L23:
	;
	v86 = base.I64_extend_i32_u(v77) * v37
	goto L20
L24:
	;
	if l7 != int32(2147483647) {
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v90 = v53 + v37
	if v90 < v71 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L4
L28:
	;
	v92 = v90
	goto L30
L29:
	;
	v92 = v71
	goto L30
L30:
	;
	if v38 < v42 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v94 = v38
	goto L33
L32:
	;
	v94 = v42
	goto L33
L33:
	;
	if v86+v92*base.I64_extend_i32_s(v94) <= v76 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v120 = int32(-1)
	v124 = v92
	v125 = v73
	goto L16
L35:
	;
	goto L36
L36:
	;
	if int64(0) < v90 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v120 = l7
	v124 = v92
	v125 = v73
	goto L16
L38:
	;
	v117 = v114
	goto L40
L39:
	;
	v117 = v115
	goto L40
L40:
	;
	v120 = l7
	v124 = v92
	v125 = base.I32_wrap_i64(v117)
	goto L16
L41:
	;
	v133 = F_palloc(m, v73<<(uint(int32(2))%32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v243 = v20
	goto L43
L43:
	;
	v254 = F_palloc(m, v73<<(uint(int32(4))%32))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L52
	}
L44:
	;
	v135 = int32(0)
	if v135 < v38 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v140 = l0
	v143 = v135
	goto L48
L46:
	;
	v214 = int32(0)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133+v214<<(uint(int32(2))%32)))) = int32(0)
	v243 = v133
	goto L43
L48:
	;
	v174 = F_pg_mblen_range(m, v140, l0+l1)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v214 = v38
	goto L47
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133+v143<<(uint(int32(2))%32)))) = v174
	v179 = v143 + int32(1)
	if v179 != v38 {
		v140 = v140 + v174
		v143 = v179
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v257 = v42 + int32(1)
	if v125 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if v257 < int32(2) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v260 = base.I64_extend_i32_u(v125)
	v262 = v260 & int64(3)
	v263 = int64(0)
	if base.Ui32(int32(4)) <= base.Ui32(v125) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v277 = v263
	v285 = v10
	goto L58
L56:
	;
	v346 = v263
	goto L57
L57:
	;
	v378 = v346
	v385 = v10
	goto L62
L58:
	;
	v301 = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v254+base.I32_wrap_i64(v277)<<(uint(v301)%32)))) = v277 * v53
	v307 = v277 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v254+base.I32_wrap_i64(v307)<<(uint(v301)%32)))) = v307 * v53
	v315 = v277 | int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v254+base.I32_wrap_i64(v315)<<(uint(v301)%32)))) = v315 * v53
	v323 = v277 | int64(3)
	*(*int64)(unsafe.Add(mBase, uint32(v254+base.I32_wrap_i64(v323)<<(uint(v301)%32)))) = v323 * v53
	v330 = int64(4)
	v331 = v277 + v330
	v333 = v285 + v330
	if v333 != v260&int64(2147483644) {
		v277 = v331
		v285 = v333
		goto L58
	} else {
		goto L60
	}
L59:
	;
	if v262 == int64(0) {
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v346 = v331
	goto L57
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v254+base.I32_wrap_i64(v378)<<(uint(int32(3))%32)))) = v378 * v53
	v407 = int64(1)
	v410 = v385 + v407
	if v410 != v262 {
		v378 = v378 + v407
		v385 = v410
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L53
L64:
	;
	goto L63
L65:
	;
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v1066+v38<<(uint(int32(3))%32))))
	if base.Ui64(v1095-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		goto L4
	} else {
		goto L149
	}
L66:
	;
	v1066 = v254
	goto L65
L67:
	;
	goto L68
L68:
	;
	v451 = base.I64_extend_i32_s(v120)
	v452 = int64(1)
	v453 = v451 + v452
	__phi456 = l0
	__phi457 = v254
	__phi458 = l2
	__phi464 = v254 + v73<<(uint(int32(3))%32)
	__phi474 = v452
	__phi475 = v125
	__phi481 = v20
	v456 = __phi456
	v457 = __phi457
	v458 = __phi458
	v464 = __phi464
	v474 = __phi474
	v475 = __phi475
	v481 = __phi481
	goto L69
L69:
	;
	if l3 != v42 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v1066 = v464
	goto L65
L71:
	;
	v490 = F_pg_mblen_range(m, v458, l2+l3)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v492 = int32(1)
	goto L73
L73:
	;
	if v475 < v73 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v492 = v490
	goto L73
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v457+v475<<(uint(int32(3))%32)))) = v453
	v500 = v475 + int32(1)
	goto L77
L76:
	;
	v500 = v475
	goto L77
L77:
	;
	if v481 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v464))) = v37 * v474
	v506 = int32(1)
	goto L80
L79:
	;
	v506 = v481
	goto L80
L80:
	;
	if v243 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v120 < int32(0) {
		v1024 = v456
		v1043 = v500
		v1049 = v481
		goto L118
	} else {
		goto L119
	}
L82:
	;
	if v500 <= v506 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v500 <= v506 {
		goto L81
	} else {
		goto L105
	}
L85:
	;
	v515 = v506
	v516 = v456
	goto L86
L86:
	;
	v543 = int32(3)
	v544 = v515 << (uint(v543) % 32)
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v457+v544)))
	v547 = v546 + v37
	v550 = v515 - int32(1)
	v552 = v550 << (uint(v543) % 32)
	v554 = *(*int64)(unsafe.Add(mBase, uint32(v464+v552)))
	v555 = v554 + v53
	if v547 < v555 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L81
L88:
	;
	v557 = v547
	goto L90
L89:
	;
	v557 = v555
	goto L90
L90:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v243+v550<<(uint(int32(2))%32))))
	v562 = v516 + v561
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562-int32(1)))))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458+v492-int32(1)))))
	if base.B2i32(v565 != v566)|base.B2i32(v561 != v492) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	if v557 < v716 {
		goto L101
	} else {
		goto L102
	}
L92:
	;
	v683 = *(*int64)(unsafe.Add(mBase, uint32(v457+v552)))
	v716 = v683
	goto L91
L93:
	;
	if v492 == int32(1) {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v648 = *(*int64)(unsafe.Add(mBase, uint32(v457+v552)))
	v716 = v648 + v124
	goto L91
L96:
	;
	v575 = v492
	goto L97
L97:
	;
	if v575 <= int32(0) {
		goto L92
	} else {
		goto L99
	}
L98:
	;
	goto L95
L99:
	;
	v609 = v575 - int32(1)
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516+v609))))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609+v458))))
	if v611 == v613 {
		v575 = v609
		goto L97
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	v718 = v557
	goto L103
L102:
	;
	v718 = v716
	goto L103
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v544+v464))) = v718
	v721 = v515 + int32(1)
	if v721 != v500 {
		v515 = v721
		v516 = v562
		goto L86
	} else {
		goto L104
	}
L104:
	;
	goto L87
L105:
	;
	v729 = *(*int64)(unsafe.Add(mBase, uint32(v464+v506<<(uint(int32(3))%32)-int32(8))))
	v731 = v456
	v734 = v506
	v739 = v729
	goto L106
L106:
	;
	v763 = v734 << (uint(int32(3)) % 32)
	v765 = v763 + v457
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v765)))
	v767 = v766 + v37
	v768 = v739 + v53
	if v767 < v768 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L81
L108:
	;
	v770 = v767
	goto L110
L109:
	;
	v770 = v768
	goto L110
L110:
	;
	v773 = *(*int64)(unsafe.Add(mBase, uint32(v765-int32(8))))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731))))
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v775 != v776 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v778 = v124
	goto L113
L112:
	;
	v778 = int64(0)
	goto L113
L113:
	;
	v779 = v773 + v778
	if v770 < v779 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v781 = v770
	goto L116
L115:
	;
	v781 = v779
	goto L116
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v464+v763))) = v781
	v783 = int32(1)
	v786 = v734 + v783
	if v786 != v500 {
		v731 = v731 + v783
		v734 = v786
		v739 = v781
		goto L106
	} else {
		goto L117
	}
L117:
	;
	goto L107
L118:
	;
	v1058 = v474 + int64(1)
	if v1058 != base.I64_extend_i32_u(v257) {
		__phi456 = v1024
		__phi457 = v464
		__phi458 = v458 + v492
		__phi464 = v457
		__phi474 = v1058
		__phi475 = v1043
		__phi481 = v1049
		v456 = __phi456
		v457 = __phi457
		v458 = __phi458
		v464 = __phi464
		v474 = __phi474
		v475 = __phi475
		v481 = __phi481
		goto L69
	} else {
		goto L148
	}
L119:
	;
	v826 = v38 - v42 + base.I32_wrap_i64(v474)
	v837 = base.I64_extend_i32_s(v500)
	goto L123
L120:
	;
	if v120 != int32(2147483647) {
		v1133 = v120 + int32(1)
		goto L6
	} else {
		goto L147
	}
L121:
	;
	if v980 < v884 {
		v1024 = v955
		v1043 = v884
		v1049 = v980
		goto L118
	} else {
		goto L146
	}
L122:
	;
	if v884 <= v481 {
		v955 = v456
		v980 = v481
		goto L121
	} else {
		goto L131
	}
L123:
	;
	if v837 <= int64(0) {
		v884 = v500 >> (uint(int32(31)) % 32) & v500
		goto L122
	} else {
		goto L125
	}
L124:
	;
	v884 = base.I32_wrap_i64(v837)
	goto L122
L125:
	;
	v863 = v837 - int64(1)
	v864 = base.I32_wrap_i64(v863)
	v868 = *(*int64)(unsafe.Add(mBase, uint32(v464+v864<<(uint(int32(3))%32))))
	v869 = v864 - v826
	if int32(0) < v869 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v451 < v868+v878 {
		v837 = v863
		goto L123
	} else {
		goto L130
	}
L127:
	;
	v878 = base.I64_extend_i32_u(v869) * v37
	goto L126
L128:
	;
	goto L129
L129:
	;
	v878 = base.I64_extend_i32_s(int32(0)-v869) * v53
	goto L126
L130:
	;
	goto L124
L131:
	;
	v888 = v456
	v897 = base.I64_extend_i32_s(v481)
	goto L132
L132:
	;
	v920 = base.I32_wrap_i64(v897)
	v922 = v920 << (uint(int32(3)) % 32)
	v923 = v464 + v922
	v924 = *(*int64)(unsafe.Add(mBase, uint32(v923)))
	v925 = v920 - v826
	if int32(0) < v925 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L120
L134:
	;
	if v924+v934 <= v451 {
		v955 = v888
		v980 = v920
		goto L121
	} else {
		goto L138
	}
L135:
	;
	v934 = base.I64_extend_i32_u(v925) * v37
	goto L134
L136:
	;
	goto L137
L137:
	;
	v934 = base.I64_extend_i32_s(int32(0)-v925) * v53
	goto L134
L138:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v923))) = v453
	*(*int64)(unsafe.Add(mBase, uint32(v922+v457))) = v453
	if v897 != int64(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v243 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v951 = v888
	goto L141
L141:
	;
	v953 = v897 + int64(1)
	if v953 != base.I64_extend_i32_s(v884) {
		v888 = v951
		v897 = v953
		goto L132
	} else {
		goto L145
	}
L142:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v243+v920<<(uint(int32(2))%32)-int32(4))))
	v949 = v947
	goto L144
L143:
	;
	v949 = int32(1)
	goto L144
L144:
	;
	v951 = v949 + v888
	goto L141
L145:
	;
	goto L133
L146:
	;
	goto L120
L147:
	;
	goto L4
L148:
	;
	goto L70
L149:
	;
	v1133 = base.I32_wrap_i64(v1095)
	goto L6
L150:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = int32(255)
	F_errmsg(m, int32(_a_F_varstr_levenshtein_less_equal_0), v35)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_less_equal_1), int32(138), int32(_a_F_varstr_levenshtein_less_equal_2))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	F_errmsg(m, int32(_a_F_varstr_levenshtein_less_equal_3), int32(0))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_varstr_levenshtein_less_equal_4), int32(_a_F_varstr_levenshtein_less_equal_5), int32(_a_F_varstr_levenshtein_less_equal_6))
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
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
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	var v226 int32
	_ = v226
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
	return v226
L11:
	;
	if v51 == int32(0) {
		v226 = v49
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
	v186 = v18
	goto L15
L15:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v193 != 0 {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	if v66 == int32(0) {
		v226 = v49
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
	F_LockBufferInternal(m, v66, int32(3))
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
	v174 = int32(1)
	v176 = int32(_a_F_visibilitymap_prepare_truncate_1)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[2])) = v178 - v174
	F_UnlockReleaseBuffer(m, v66)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
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
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[5])))
	if v161 == int32(0) {
		goto L42
	} else {
		goto L43
	}
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
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_prepare_truncate[6]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+268))
	goto L45
L43:
	;
	goto L44
L44:
	;
	F_log_newpage_buffer(m, v66, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	if base.B2i32(v166 != int32(0)) == int32(0) {
		goto L34
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L34
L48:
	;
	v186 = v18 + v174
	goto L15
L49:
	;
	v217 = v193
	goto L51
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v195
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v197
	v199 = F_smgropen(m, v15, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	v219 = F_smgrnblocks(m, v217, int32(2))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v199
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	if v203 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v217 = v215
	goto L51
L54:
	;
	v211 = v203
	goto L56
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v199)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v205
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v199)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v207
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v199)+72))
	v211 = v209
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+72)) = v211 + int32(1)
	goto L53
L57:
	;
	if base.Ui32(v219) <= base.Ui32(v186) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v222 = int32(-1)
	goto L60
L59:
	;
	v222 = v186
	goto L60
L60:
	;
	v226 = v222
	goto L10
}
func F_visibilitymap_set(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v8 = base.I32_div_u_s(l0, int32(_a_F_visibilitymap_set_0))
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v75 = m.ExcPending
		if v75 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_visibilitymap_set_1), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(288), int32(_a_F_visibilitymap_set_3))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l1 < int32(0) {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[0]))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+(l1^int32(-1))*int32(56))+16))
			v29 = v20
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[1]))
			v23 = int32(56)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+l1*v23-v23)+16))
			v29 = v28
		}
		if v29 != v8 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_set_1), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_set_2), int32(288), int32(_a_F_visibilitymap_set_3))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v34 = l0 << (uint(int32(1)) % 32) & int32(6)
			if l1 < int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[2]))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38+(l1^int32(-1))<<(uint(int32(2))%32))))
				v52 = v44
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_set[3]))
				v52 = v46 + l1<<(uint(int32(13))%32) + int32(-8192)
			}
			v58 = v52 + int32(base.Ui32(l0-v8*int32(_a_F_visibilitymap_set_0))>>(uint(int32(2))%32))
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)))
			v62 = int32(base.Ui32(v59)>>(uint(v34)%32)) & int32(3)
			if l2 != v62 {
				v65 = v59 | l2<<(uint(v34)%32)
				*(*uint8)(unsafe.Add(mBase, uint32(v58)+24)) = uint8(v65)
				F_MarkBufferDirty(m, l1)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					return v62
				}
			} else {
				return v62
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
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
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
	return v215
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
		v215 = v62
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
	F_UnlockBuffer(m, v118)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
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
		v215 = v118
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
		v215 = v118
		goto L10
	} else {
		goto L41
	}
L38:
	;
	F_LockBufferInternal(m, v118, int32(3))
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
	F_LockBufferInternal(m, v118, int32(3))
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
	v215 = v118
	goto L10
}
