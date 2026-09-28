package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterConstraintNamespaces(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v17 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = v13 + int32(16)
	v25 = base.I64_extend_i32_u(l0)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int64(0)
	goto L5
L4:
	;
	v26 = v25
	goto L5
L5:
	;
	F_ScanKeyInit(m, v20, int32(9), int32(3), int32(184), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v35 = v25
	goto L9
L8:
	;
	v35 = int64(0)
	goto L9
L9:
	;
	F_ScanKeyInit(m, v13+int32(72), int32(10), int32(3), int32(184), v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v42 = F_systable_beginscan(m, v17, int32(2665), int32(1), int32(0), int32(2), v20)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v44
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(2606)
	v60 = v56 + v57
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v61
	v66 = v13 + int32(4)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v73 = v71 - int32(1)
	if v73 < v62 {
		v109 = v62
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	if v109 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	goto L18
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v80 = v73
	goto L21
L21:
	;
	v86 = v77 + v80*int32(12)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v76 != v87 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v109 = v103
	goto L19
L23:
	;
	v103 = int32(0)
	if v103 < v80 {
		v80 = v80 - int32(1)
		goto L21
	} else {
		goto L27
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v89 != v90 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	if base.B2i32(v93 == v94)|base.B2i32(v93 == int32(0)) != 0 {
		v109 = int32(1)
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L22
L28:
	;
	if l1 == l2 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v144 = F_systable_getnext(m, v42)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L41
	}
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_AlterConstraintNamespaces[0]))
	if v130 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v60)+68))
	if v115 != l1 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v117 = F_heap_copytuple(m, v49)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v119+v120)+68)) = l2
	F_CatalogTupleUpdate(m, v17, v117+int32(4), v117)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v133 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v132, v133, v133, v133)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_add_exact_object_address(m, v13+int32(4), l4)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L30
L41:
	;
	if v144 != 0 {
		v49 = v144
		goto L16
	} else {
		goto L42
	}
L42:
	;
	goto L17
L43:
	;
	F_relation_close(m, v17, int32(3))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v13 + int32(128)
	return
}
func F_AlterFKConstrEnforceabilityRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v20 = v15 + int32(-56)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v27 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v24+v25))))
	F_ScanKeyInit(m, v20, int32(12), int32(3), int32(184), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = int32(1)
	v34 = F_systable_beginscan(m, l2, int32(2579), v31, int32(0), v31, v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	goto L4
L4:
	;
	v50 = F_systable_getnext(m, v34)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	if v50 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v52 = F_ATExecAlterFKConstrEnforceability(m, l0, l1, l2, l3, l4, l5, v50, l7, l8, l9, l10, l11)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L5
L10:
	;
	goto L4
L11:
	;
	m.G0 = v17 - int32(-64)
	return
}
func F_AlterTableGetLockLevel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l0 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v256
L2:
	;
	v256 = int32(4)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = int32(4)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= int32(0) {
		v256 = v21
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v28 = v21
	v32 = v2
	goto L6
L6:
	;
	v38 = int32(8)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v32<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	switch v44 {
	case 0, 1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 19, 21, 22, 24, 25, 26, 29, 30, 31, 32, 33, 36, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 62, 63, 64:
		v235 = v38
		goto L8
	case 8, 9, 10, 20, 27, 28, 59, 61:
		goto L13
	default:
		goto L10
	case 16, 17, 18:
		goto L14
	case 34, 35:
		goto L12
	case 37, 38, 39, 40, 41, 42, 43, 44:
		goto L9
	case 60:
		goto L11
	}
L7:
	;
	v256 = v248
	goto L1
L8:
	;
	if v28 < v235 {
		goto L63
	} else {
		goto L64
	}
L9:
	;
	v235 = int32(6)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L24
	} else {
		goto L60
	}
L11:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+12)))
	if v217 != 0 {
		goto L57
	} else {
		goto L58
	}
L12:
	;
	v56 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v59 == v56 {
		v213 = int32(8)
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v235 = int32(4)
	goto L8
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(161) {
		v235 = v38
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v51 == int32(9) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = int32(6)
	goto L18
L17:
	;
	v54 = int32(8)
	goto L18
L18:
	;
	v235 = v54
	goto L8
L19:
	;
	v235 = v213
	goto L8
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AlterTableGetLockLevel[0])))
	if v63 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_initialize_reloptions(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if int32(0) < v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	return int32(0)
L25:
	;
	goto L23
L26:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTableGetLockLevel[1]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v81 = v56
	v83 = v56
	goto L29
L27:
	;
	v191 = v56
	goto L28
L28:
	;
	v213 = v191
	goto L19
L29:
	;
	if v76 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v191 = v175
	goto L28
L31:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v73+v83<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v97 = v76
	v100 = v81
	v104 = int32(0)
	goto L34
L32:
	;
	v175 = v81
	goto L33
L33:
	;
	v185 = v83 + int32(1)
	if v185 != v70 {
		v81 = v175
		v83 = v185
		goto L29
	} else {
		goto L56
	}
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	v112 = v110 + int32(1)
	if v112 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v175 = v164
	goto L33
L36:
	;
	if v157 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	v157 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v118 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v119 = v109
	v120 = v94
	v121 = v112
	v122 = v118
	goto L44
L41:
	;
	v145 = v94
	v149 = int32(0)
	goto L42
L42:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v157 = v149 - v150
	goto L36
L43:
	;
	v145 = v140
	v149 = v142
	goto L42
L44:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if base.B2i32(v122 != v124)|base.B2i32(v124 == int32(0)) != 0 {
		v140 = v120
		v142 = v122
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v140 = v134
	v142 = int32(0)
	goto L43
L46:
	;
	v130 = v121 - int32(1)
	if v130 == int32(0) {
		v140 = v120
		v142 = v122
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v133 = int32(1)
	v134 = v120 + v133
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v135 != 0 {
		v119 = v119 + v133
		v120 = v134
		v121 = v130
		v122 = v135
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	if v160 < v100 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v164 = v100
	goto L51
L51:
	;
	v166 = v104 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v75+v166<<(uint(int32(2))%32))))
	if v170 != 0 {
		v97 = v170
		v100 = v164
		v104 = v166
		goto L34
	} else {
		goto L55
	}
L52:
	;
	v162 = v100
	goto L54
L53:
	;
	v162 = v160
	goto L54
L54:
	;
	v164 = v162
	goto L51
L55:
	;
	goto L35
L56:
	;
	goto L30
L57:
	;
	v218 = int32(4)
	goto L59
L58:
	;
	v218 = int32(8)
	goto L59
L59:
	;
	v235 = v218
	goto L8
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v223
	F_errmsg_internal(m, int32(_a_F_AlterTableGetLockLevel_0), v16)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L24
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_AlterTableGetLockLevel_1), int32(_a_F_AlterTableGetLockLevel_2), int32(_a_F_AlterTableGetLockLevel_3))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v248 = v235
	goto L65
L64:
	;
	v248 = v28
	goto L65
L65:
	;
	v250 = v32 + int32(1)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v250 < v251 {
		v28 = v248
		v32 = v250
		goto L6
	} else {
		goto L66
	}
L66:
	;
	goto L7
}
func F_AlterTypeRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int64
	_ = v18
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int64
	_ = v129
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v7 = m.G0
	v9 = v7 - int32(400)
	m.G0 = v9
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	base.MemoryFill(m, v9+int32(144), int32(0), int32(256))
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+136)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+128)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v18
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v34 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+103)) = uint8(v37)
	v39 = int64(*(*int8)(unsafe.Add(mBase, uint32(l4)+7)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+328)) = v39
	goto L5
L4:
	;
	goto L5
L5:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v41 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+97)) = uint8(v44)
	v46 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+280)) = v46
	goto L8
L7:
	;
	goto L8
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)))
	if v48 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+98)) = uint8(v51)
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+288)) = v53
	goto L11
L10:
	;
	goto L11
L11:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v55 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+99)) = uint8(v58)
	v60 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+16)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+296)) = v60
	goto L14
L13:
	;
	goto L14
L14:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v62 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+100)) = uint8(v65)
	v67 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+20)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+304)) = v67
	goto L17
L16:
	;
	goto L17
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	if v69 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+101)) = uint8(v72)
	v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+24)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+312)) = v74
	goto L20
L19:
	;
	goto L20
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)))
	if v76 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v79 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+92)) = uint8(v79)
	v81 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l4)+28)))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+240)) = v81
	goto L23
L22:
	;
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v90 = F_heap_modify_tuple(m, l2, v83, v9+int32(144), v9+int32(112), v9+int32(80))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_CatalogTupleUpdate(m, l3, v90+int32(4), v90)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v96 = int32(0)
	F_GenerateTypeDependencies(m, v90, l3, v96, v96, v96, l1, l1, v96, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTypeRecurse[0]))
	if v104 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), l0, v106, v106, v106)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if l1 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L63
	}
L32:
	;
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v154)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+3)) = uint16(v154)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)) = uint8(v154)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v160 != 0 {
		goto L44
	} else {
		goto L45
	}
L33:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v111 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v114 != int32(1) {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)+96))
	if v120 == int32(0) {
		goto L32
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v125 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v120))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v125 == int32(0) {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v129 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v129
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+19)) = uint8(v137)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v139)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v143
	F_AlterTypeRecurse(m, v120, int32(1), v125, l3, v9+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_ReleaseCatCache(m, v125)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	m.G0 = v9 + int32(400)
	return
L44:
	;
	v166 = v9 + int32(16)
	F_ScanKeyInit(m, v166, int32(26), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)))
	if v161 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	if v162 != int32(1) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v173 = int32(0)
	v177 = F_systable_beginscan(m, l3, v173, v173, v173, int32(1), v166)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v179 = F_systable_getnext(m, v177)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v179 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v182 = v179
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_systable_endscan(m, v177)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L62
	}
L54:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+79)))
	if v190 == int32(100) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	F_AlterTypeRecurse(m, v193, int32(0), v182, l3, l4)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v197 = F_systable_getnext(m, v177)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v197 != 0 {
		v182 = v197
		goto L54
	} else {
		goto L61
	}
L61:
	;
	goto L55
L62:
	;
	goto L43
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v120
	F_errmsg_internal(m, int32(_a_F_AlterTypeRecurse_0), v9)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_AlterTypeRecurse_1), int32(_a_F_AlterTypeRecurse_2), int32(_a_F_AlterTypeRecurse_3))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
