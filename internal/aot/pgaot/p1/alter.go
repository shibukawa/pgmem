package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterConstraintNamespaces(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
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
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v16 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = v12 + int32(16)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(0)
	goto L5
L4:
	;
	v24 = l0
	goto L5
L5:
	;
	F_ScanKeyInit(m, v19, int32(9), int32(3), int32(184), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v33 = l0
	goto L9
L8:
	;
	v33 = int32(0)
	goto L9
L9:
	;
	F_ScanKeyInit(m, v12-int32(-64), int32(10), int32(3), int32(184), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = F_systable_beginscan(m, v16, int32(2665), int32(1), int32(0), int32(2), v19)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v42 = F_systable_getnext(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v42
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_systable_endscan(m, v40)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L43
	}
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(2606)
	v57 = v53 + v54
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v58
	v63 = v12 + int32(4)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v70 = v68 - int32(1)
	if v70 < v59 {
		v106 = v59
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	if v106 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	goto L18
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v77 = v70
	goto L21
L21:
	;
	v83 = v74 + v77*int32(12)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v73 != v84 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v106 = v100
	goto L19
L23:
	;
	v100 = int32(0)
	if v100 < v77 {
		v77 = v77 - int32(1)
		goto L21
	} else {
		goto L27
	}
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 != v87 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if base.B2i32(v90 == v91)|base.B2i32(v90 == int32(0)) != 0 {
		v106 = int32(1)
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
	v141 = F_systable_getnext(m, v40)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L41
	}
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_AlterConstraintNamespaces[0]))
	if v127 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v57)+68))
	if v112 != l1 {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v114 = F_heap_copytuple(m, v47)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v116+v117)+68)) = l2
	F_CatalogTupleUpdate(m, v16, v114+int32(4), v114)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v130 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v129, v130, v130, v130)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_add_exact_object_address(m, v12+int32(4), l4)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
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
	if v141 != 0 {
		v47 = v141
		goto L16
	} else {
		goto L42
	}
L42:
	;
	goto L17
L43:
	;
	F_relation_close(m, v16, int32(3))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v12 + int32(112)
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
	var v13 int32
	_ = v13
	var v18 int64
	_ = v18
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
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
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v7 = m.G0
	v9 = v7 - int32(256)
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
	v13 = int32(128)
	base.MemoryFill(m, v9+v13, int32(0), v13)
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+112)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+96)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+64)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+80)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v9)+88)) = v18
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v34 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+87)) = uint8(v37)
	v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4)+7)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+220)) = v39
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+81)) = uint8(v44)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+196)) = v46
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+82)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+200)) = v53
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+83)) = uint8(v58)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+204)) = v60
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+84)) = uint8(v65)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+208)) = v67
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+85)) = uint8(v72)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+212)) = v74
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
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)) = uint8(v79)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v81
	goto L23
L22:
	;
	goto L23
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v90 = F_heap_modify_tuple(m, l2, v83, v9+int32(128), v9+int32(96), v9-int32(-64))
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
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L63
	}
L32:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v153)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+3)) = uint16(v153)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)) = uint8(v153)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v159 != 0 {
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
	v124 = F_SearchSysCache1(m, int32(82), v120)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v124 == int32(0) {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v128
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v128
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+19)) = uint8(v136)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v138)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v142
	F_AlterTypeRecurse(m, v120, int32(1), v124, l3, v9+int32(16))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_ReleaseCatCache(m, v124)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	m.G0 = v9 + int32(256)
	return
L44:
	;
	v165 = v9 + int32(16)
	F_ScanKeyInit(m, v165, int32(26), int32(3), int32(184), l0)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)))
	if v160 != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	if v161 != int32(1) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v171 = int32(0)
	v175 = F_systable_beginscan(m, l3, v171, v171, v171, int32(1), v165)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v177 = F_systable_getnext(m, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v177 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v180 = v177
	goto L54
L52:
	;
	goto L53
L53:
	;
	F_systable_endscan(m, v175)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L62
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+22)))
	v187 = v185 + v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+79)))
	if v188 == int32(100) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	F_AlterTypeRecurse(m, v191, int32(0), v180, l3, l4)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v195 = F_systable_getnext(m, v175)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	if v195 != 0 {
		v180 = v195
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
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_AlterTypeRecurse_1), int32(_a_F_AlterTypeRecurse_2), int32(_a_F_AlterTypeRecurse_3))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
