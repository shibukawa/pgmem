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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	F_ScanKeyInit(m, v12+int32(16), int32(9), int32(3), int32(184), v24)
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
	v42 = F_systable_beginscan(m, v16, int32(2665), int32(1), int32(0), int32(2), v12+int32(16))
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
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L44
	}
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(2606)
	v59 = v55 + v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v60
	v65 = v12 + int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v72 = v70 - int32(1)
	if v72 < v61 {
		v107 = v61
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L15
L18:
	;
	if v107 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	goto L18
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v79 = v72
	goto L21
L21:
	;
	v85 = v76 + v79*int32(12)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v75 != v86 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v107 = v101
	goto L19
L23:
	;
	v101 = int32(0)
	if v101 < v79 {
		v79 = v79 - int32(1)
		goto L21
	} else {
		goto L28
	}
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v88 != v89 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v91 = int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if v92 == v93 {
		v107 = v91
		goto L19
	} else {
		goto L26
	}
L26:
	;
	if v92 == int32(0) {
		v107 = v91
		goto L19
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	goto L22
L29:
	;
	if l1 == l2 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v142 = F_systable_getnext(m, v42)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L42
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v128 != 0 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v59)+68))
	if v113 != l1 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v115 = F_heap_copytuple(m, v49)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v118)+68)) = l2
	F_CatalogTupleUpdate(m, v16, v115+int32(4), v115)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v131 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v130, v131, v131, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_add_exact_object_address(m, v12+int32(4), l4)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L31
L42:
	;
	if v142 != 0 {
		v49 = v142
		goto L16
	} else {
		goto L43
	}
L43:
	;
	goto L17
L44:
	;
	F_sequence_close(m, v16, int32(3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
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
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
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
	return v255
L2:
	;
	v255 = int32(4)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = int32(4)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= int32(0) {
		v255 = v21
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
		v234 = v38
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
	v255 = v247
	goto L1
L8:
	;
	if v28 < v234 {
		goto L64
	} else {
		goto L65
	}
L9:
	;
	v234 = int32(6)
	goto L8
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L24
	} else {
		goto L61
	}
L11:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+12)))
	if v216 != 0 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v56 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	if v59 == v56 {
		v212 = int32(8)
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v234 = int32(4)
	goto L8
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(161) {
		v234 = v38
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
	v234 = v54
	goto L8
L19:
	;
	v234 = v212
	goto L8
L20:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[404])))
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
	v75 = *(*int32)(unsafe.Add(mBase, _consts[405]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v81 = v56
	v83 = v56
	goto L29
L27:
	;
	v190 = v56
	goto L28
L28:
	;
	v212 = v190
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
	v190 = v174
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
	v174 = v81
	goto L33
L33:
	;
	v184 = v83 + int32(1)
	if v184 != v70 {
		v81 = v174
		v83 = v184
		goto L29
	} else {
		goto L57
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
	v174 = v163
	goto L33
L36:
	;
	if v156 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L37:
	;
	v156 = int32(0)
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
	v144 = v94
	v148 = int32(0)
	goto L42
L42:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v156 = v148 - v149
	goto L36
L43:
	;
	v144 = v139
	v148 = v141
	goto L42
L44:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v122 != v124 {
		v139 = v120
		v141 = v122
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v139 = v133
	v141 = int32(0)
	goto L43
L46:
	;
	if v124 == int32(0) {
		v139 = v120
		v141 = v122
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v129 = v121 - int32(1)
	if v129 == int32(0) {
		v139 = v120
		v141 = v122
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v132 = int32(1)
	v133 = v120 + v132
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+1)))
	if v134 != 0 {
		v119 = v119 + v132
		v120 = v133
		v121 = v129
		v122 = v134
		goto L44
	} else {
		goto L49
	}
L49:
	;
	goto L45
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	if v159 < v100 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v163 = v100
	goto L52
L52:
	;
	v165 = v104 + int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v75+v165<<(uint(int32(2))%32))))
	if v169 != 0 {
		v97 = v169
		v100 = v163
		v104 = v165
		goto L34
	} else {
		goto L56
	}
L53:
	;
	v161 = v100
	goto L55
L54:
	;
	v161 = v159
	goto L55
L55:
	;
	v163 = v161
	goto L52
L56:
	;
	goto L35
L57:
	;
	goto L30
L58:
	;
	v217 = int32(4)
	goto L60
L59:
	;
	v217 = int32(8)
	goto L60
L60:
	;
	v234 = v217
	goto L8
L61:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v222
	F_errmsg_internal(m, int32(485786), v16)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L24
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(494621), int32(4849), int32(306704))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L24
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v247 = v234
	goto L66
L65:
	;
	v247 = v28
	goto L66
L66:
	;
	v249 = v32 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v249 < v250 {
		v28 = v247
		v32 = v249
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L7
}
func F_AlterTypeRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v8 = m.G0
	v10 = v8 - int32(256)
	m.G0 = v10
	F_check_stack_depth(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = int32(128)
	v19 = F__emscripten_memset_bulkmem(m, v10+v14, base.I32_extend8_s(int32(0)), v14)
	mBase = m.M
	goto L3
L3:
	;
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v20
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v36 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v39 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+87)) = uint8(v39)
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4)+7)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+220)) = v41
	goto L6
L5:
	;
	goto L6
L6:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v43 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+81)) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+196)) = v48
	goto L9
L8:
	;
	goto L9
L9:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)))
	if v50 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+82)) = uint8(v53)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+200)) = v55
	goto L12
L11:
	;
	goto L12
L12:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v57 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+83)) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+204)) = v62
	goto L15
L14:
	;
	goto L15
L15:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v64 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+84)) = uint8(v67)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+208)) = v69
	goto L18
L17:
	;
	goto L18
L18:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	if v71 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+85)) = uint8(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+212)) = v76
	goto L21
L20:
	;
	goto L21
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)))
	if v78 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+76)) = uint8(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+176)) = v83
	goto L24
L23:
	;
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	v92 = F_heap_modify_tuple(m, l2, v85, v10+int32(128), v10+int32(96), v10-int32(-64))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_CatalogTupleUpdate(m, l3, v92+int32(4), v92)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v98 = int32(0)
	F_GenerateTypeDependencies(m, v92, l3, v98, v98, v98, l1, l1, v98, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[297]))
	if v106 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1247), l0, v108, v108, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if l1 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L64
	}
L33:
	;
	v158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+6)) = uint8(v158)
	*(*uint16)(unsafe.Add(mBase, uint32(l4)+3)) = uint16(v158)
	*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)) = uint8(v158)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v164 != 0 {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	if v113 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v116 != int32(1) {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v92)+16))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+22)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119+v120)+96))
	if v122 == int32(0) {
		goto L33
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v126 = F_SearchSysCache1(m, int32(82), v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v126 == int32(0) {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v130 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v130
	v133 = v10 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v130
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v130
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+19)) = uint8(v140)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v142)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v146
	F_AlterTypeRecurse(m, v122, int32(1), v126, l3, v10+int32(16))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_ReleaseCatCache(m, v126)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	m.G0 = v10 + int32(256)
	return
L45:
	;
	F_ScanKeyInit(m, v10+int32(16), int32(26), int32(3), int32(184), l0)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+2)))
	if v165 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+5)))
	if v166 != int32(1) {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v176 = int32(0)
	v182 = F_systable_beginscan(m, l3, v176, v176, v176, int32(1), v10+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v184 = F_systable_getnext(m, v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v184 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v187 = v184
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_systable_endscan(m, v182)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L63
	}
L55:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+16))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+22)))
	v195 = v193 + v194
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+79)))
	if v196 == int32(100) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	F_AlterTypeRecurse(m, v199, int32(0), v187, l3, l4)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v203 = F_systable_getnext(m, v182)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	if v203 != 0 {
		v187 = v203
		goto L55
	} else {
		goto L62
	}
L62:
	;
	goto L56
L63:
	;
	goto L44
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v122
	F_errmsg_internal(m, int32(50356), v10)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(494610), int32(4653), int32(360799))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
