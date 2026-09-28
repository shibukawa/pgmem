package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compare_text_lexemes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v10 = int32(1)
	v11 = v9 & v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v9 == v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v40 = base.I32_wrap_i64(v12)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v42 = int32(1)
	v43 = v41 & v42
	if v41 == v42 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v18 == int32(18) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v29 = int32(1)
	if v11 != 0 {
		v39 = int32(base.Ui32(v9)>>(uint(v29)%32)) - v29
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v21 = int32(16)
	goto L7
L6:
	;
	v21 = int32(0)
	goto L7
L7:
	;
	if base.Ui32((v18-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = int32(4)
	goto L10
L9:
	;
	v28 = v21
	goto L10
L10:
	;
	v39 = v28
	goto L1
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	if v39 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v49 == int32(18) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v60 = int32(1)
	if v43 != 0 {
		v70 = int32(base.Ui32(v41)>>(uint(v60)%32)) - v60
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v52 = int32(16)
	goto L18
L17:
	;
	v52 = int32(0)
	goto L18
L18:
	;
	if base.Ui32((v49-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v59 = int32(4)
	goto L21
L20:
	;
	v59 = v52
	goto L21
L21:
	;
	v70 = v59
	goto L12
L22:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
	goto L12
L23:
	;
	v74 = int32(0)
	if v74 < v70 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v70 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v77 = int32(-1)
	goto L28
L27:
	;
	v77 = v74
	goto L28
L28:
	;
	return v77
L29:
	;
	return base.B2i32(int32(0) < v39)
L30:
	;
	goto L31
L31:
	;
	if v11 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	return v162
L33:
	;
	v86 = int32(1)
	goto L35
L34:
	;
	v86 = int32(4)
	goto L35
L35:
	;
	v87 = v8 + v86
	if v43 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v90 = int32(1)
	goto L38
L37:
	;
	v90 = int32(4)
	goto L38
L38:
	;
	v91 = v40 + v90
	if base.Ui32(v39) < base.Ui32(v70) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v93 = v39
	goto L41
L40:
	;
	v93 = v70
	goto L41
L41:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v93) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if v155 != 0 {
		v162 = v155
		goto L32
	} else {
		goto L60
	}
L43:
	;
	v155 = int32(0)
	goto L42
L44:
	;
	v129 = v124
	v130 = v125
	v131 = v126
	goto L54
L45:
	;
	if (v87|v91)&int32(3) != 0 {
		v124 = v87
		v125 = v91
		v126 = v93
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v117 = v87
	v118 = v91
	v119 = v93
	goto L47
L47:
	;
	if v119 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v101 = v87
	v102 = v91
	v103 = v93
	goto L49
L49:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v106 != v107 {
		v124 = v101
		v125 = v102
		v126 = v103
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v117 = v112
	v118 = v110
	v119 = v114
	goto L47
L51:
	;
	v109 = int32(4)
	v110 = v102 + v109
	v112 = v101 + v109
	v114 = v103 - v109
	if base.Ui32(int32(3)) < base.Ui32(v114) {
		v101 = v112
		v102 = v110
		v103 = v114
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L44
L54:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if v134 == v135 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v155 = v134 - v135
	goto L42
L56:
	;
	v137 = int32(1)
	v142 = v131 - v137
	if v142 != 0 {
		v129 = v129 + v137
		v130 = v130 + v137
		v131 = v142
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L43
L60:
	;
	if v39 == v70 {
		v162 = int32(0)
		goto L32
	} else {
		goto L61
	}
L61:
	;
	if v39 < v70 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v161 = int32(-1)
	goto L64
L63:
	;
	v161 = int32(1)
	goto L64
L64:
	;
	v162 = v161
	goto L32
}
func F_textToQualifiedNameList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v46 = F_palloc(m, v43+int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L14
	}
L2:
	;
	return int32(0)
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v14 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v20 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	if v14&v31 != 0 {
		v43 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v23 = int32(16)
	goto L9
L8:
	;
	v23 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = int32(4)
	goto L12
L11:
	;
	v30 = v23
	goto L12
L12:
	;
	v43 = v30
	goto L1
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = int32(1)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v50&v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v43))) = uint8(v57)
	if l0 != v10 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v53 = v48
	goto L20
L19:
	;
	v53 = int32(4)
	goto L20
L20:
	;
	base.MemoryCopy(m, v46, v10+v53, v43)
	goto L17
L21:
	;
	F_pfree(m, v10)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = F_SplitIdentifierString(m, v46, int32(46), v8+int32(12))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L46
	}
L26:
	;
	if v65 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v67 == int32(0) {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L42
	}
L30:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v70 < v71 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v75 = v70
	v78 = int32(0)
	goto L34
L32:
	;
	v95 = v70
	goto L33
L33:
	;
	F_pfree(m, v46)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v78<<(uint(int32(2))%32))))
	v85 = F_pstrdup(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L36
	}
L35:
	;
	v95 = v89
	goto L33
L36:
	;
	v87 = F_makeString(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v89 = F_lappend(m, v75, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v92 = v78 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v92 < v93 {
		v75 = v89
		v78 = v92
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	F_list_free(m, v67)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	m.G0 = v8 + int32(16)
	return v95
L42:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_textToQualifiedNameList_0), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_textToQualifiedNameList_1), int32(2736), int32(_a_F_textToQualifiedNameList_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_textToQualifiedNameList_0), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_textToQualifiedNameList_1), int32(2741), int32(_a_F_textToQualifiedNameList_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_pattern_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v94))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) <= v94))
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) <= v94))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(int32(0) <= v94))
				}
			}
		}
	}
}
func F_text_pattern_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v94) >> (uint(int32(31)) % 32)))
						}
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v94) >> (uint(int32(31)) % 32)))
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v94) >> (uint(int32(31)) % 32)))
					}
				} else {
					return base.I64_extend_i32_u(int32(base.Ui32(v94) >> (uint(int32(31)) % 32)))
				}
			}
		}
	}
}
func F_text_position_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
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
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 <= v2 {
		v287 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v287
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1052))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v25 = v21
	v26 = v14
	goto L7
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1056))
	v21 = v17 + v18
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v21 = v20
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1060)) = v26
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v40 = v38 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v222 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L10:
	;
	if v206 == int32(0) {
		v287 = v2
		goto L1
	} else {
		goto L56
	}
L11:
	;
	if base.Ui32(v40) <= base.Ui32(v25) {
		v287 = v2
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v26 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	v48 = v25
	goto L15
L15:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v60|base.B2i32(v40-v48 < v26) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	v73 = int32(0)
	v74 = v48
	goto L22
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = F_pg_strncoll(m, v48, v26, v41, v26, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v65 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v206 = v48
	goto L10
L22:
	;
	v83 = F_pg_mblen_range(m, v74, v40)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L19
	} else {
		goto L26
	}
L23:
	;
	if v96 != 0 {
		v212 = v96
		goto L9
	} else {
		goto L31
	}
L24:
	;
	goto L23
L25:
	;
	if base.Ui32(v85) < base.Ui32(v40) {
		v73 = v94
		v74 = v85
		goto L22
	} else {
		goto L30
	}
L26:
	;
	v85 = v83 + v74
	v86 = v85 - v48
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = F_pg_strncoll(m, v48, v86, v41, v26, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	if v88 != 0 {
		v94 = v73
		goto L25
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1060)) = v86
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v91 == int32(1) {
		v94 = v48
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v96 = v48
	goto L24
L30:
	;
	v96 = v94
	goto L24
L31:
	;
	v97 = F_pg_mblen_range(m, v48, v40)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v99 = v97 + v48
	if base.Ui32(v99) < base.Ui32(v40) {
		v48 = v99
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	if base.Ui32(v40) <= base.Ui32(v25) {
		v287 = v2
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v127 = v25 + v26 - int32(1)
	if base.Ui32(v40) <= base.Ui32(v127) {
		v287 = v2
		goto L1
	} else {
		goto L44
	}
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v108 = v25
	goto L38
L38:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v106 == v120 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v287 = v2
	goto L1
L40:
	;
	v212 = v108
	goto L9
L41:
	;
	goto L42
L42:
	;
	v123 = v108 + int32(1)
	if v123 != v40 {
		v108 = v123
		goto L38
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v130 = int32(1)
	v134 = v26 + v41 - v130
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	v139 = v127
	goto L45
L45:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v149 == v135 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v206 = v139 + (v130 - v26)
	goto L10
L47:
	;
	goto L46
L48:
	;
	v152 = v139
	v155 = v134
	goto L51
L49:
	;
	goto L50
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28)+v129&v149<<(uint(int32(2))%32))))
	v190 = v139 + v189
	if base.Ui32(v190) < base.Ui32(v40) {
		v139 = v190
		goto L45
	} else {
		goto L55
	}
L51:
	;
	if v155 == v41 {
		goto L47
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v165 = int32(1)
	v166 = v155 - v165
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v169 = v152 - v165
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v167 == v170 {
		v152 = v169
		v155 = v166
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v287 = v2
	goto L1
L56:
	;
	v212 = v206
	goto L9
L57:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = v250
	v26 = v274
	goto L7
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1052)) = v212
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1060))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1056)) = v271
	v287 = int32(1)
	goto L1
L59:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	if v226 != int32(1) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1064))
	v234 = v232
	goto L61
L61:
	;
	if base.Ui32(v212) <= base.Ui32(v234) {
		goto L58
	} else {
		goto L63
	}
L62:
	;
	goto L57
L63:
	;
	v247 = F_pg_mblen_range(m, v234, v229+v230)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L19
	} else {
		goto L64
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1064))
	v250 = v247 + v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1064)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1068)) = v252 + int32(1)
	if base.Ui32(v250) <= base.Ui32(v212) {
		v234 = v250
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
}
func F_text_position_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
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
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	v5 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v17 == int32(18) {
			v20 = int32(16)
		} else {
			v20 = int32(0)
		}
		if base.Ui32((v17-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v27 = int32(4)
		} else {
			v27 = v20
		}
		v40 = v27
	} else {
		v28 = int32(1)
		if v11&v28 != 0 {
			v40 = int32(base.Ui32(v11)>>(uint(v28)%32)) - v28
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v40 = int32(base.Ui32(v34)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v41 == int32(1) {
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v47 == int32(18) {
			v50 = int32(16)
		} else {
			v50 = int32(0)
		}
		if base.Ui32((v47-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v57 = int32(4)
		} else {
			v57 = v50
		}
		v70 = v57
	} else {
		v58 = int32(1)
		if v41&v58 != 0 {
			v70 = int32(base.Ui32(v41)>>(uint(v58)%32)) - v58
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v70 = int32(base.Ui32(v64)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	if l2 != 0 {
		v71 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return
		} else {
			v73 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v73)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v71
			v77 = *(*int32)(unsafe.Add(mBase, _c_F_text_position_setup[0]))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
			v83 = *(*int32)(unsafe.Add(mBase, uint32(v78*int32(28))+uint32(_c_F_text_position_setup[1])))
			if v83 == int32(1) {
				v86 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v86)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, _c_F_text_position_setup[0]))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
				if v90 == int32(6) {
					v93 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v93)
				} else {
					v95 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v95)
				}
			}
			v97 = int32(1)
			v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v99&v97 != 0 {
				v102 = v97
			} else {
				v102 = int32(4)
			}
			v103 = l0 + v102
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v103
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v106 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1068)) = v106
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1064)) = v103
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1052)) = v106
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v70
			*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v40
			v113 = int32(1)
			if v105&v113 != 0 {
				v117 = v113
			} else {
				v117 = int32(4)
			}
			v118 = l1 + v117
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v118
			if base.B2i32(v70 < int32(2))|base.B2i32(v40 < v70) != 0 {
			} else {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
				if v125 != int32(1) {
				} else {
					v129 = v40 - v70
					if v129 < int32(16) {
						v149 = int32(3)
					} else {
						if base.Ui32(v129) < base.Ui32(int32(64)) {
							v149 = int32(7)
						} else {
							if base.Ui32(v129) < base.Ui32(int32(128)) {
								v149 = int32(15)
							} else {
								if base.Ui32(v129) < base.Ui32(int32(512)) {
									v149 = int32(31)
								} else {
									if base.Ui32(v129) < base.Ui32(int32(2048)) {
										v149 = int32(63)
									} else {
										if base.Ui32(v129) < base.Ui32(int32(_a_F_text_position_setup_0)) {
											v148 = int32(127)
										} else {
											v148 = int32(255)
										}
										v149 = v148
									}
								}
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v149
					v152 = l3 + int32(28)
					v153 = int32(0)
					if base.Ui32(int32(7)) <= base.Ui32(v149) {
						v157 = v149 + int32(1)
						v163 = v153
						v168 = v5
						for {
							v174 = v152 + v163<<(uint(int32(2))%32)
							*(*int32)(unsafe.Add(mBase, uint32(v174)+28)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+24)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+20)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+16)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+12)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+8)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v70
							*(*int32)(unsafe.Add(mBase, uint32(v174))) = v70
							v183 = int32(8)
							v184 = v163 + v183
							v186 = v168 + v183
							if v186 != v157&int32(504) {
								v163 = v184
								v168 = v186
								continue
							} else {
								break
							}
							break
						}
						if v157&int32(4) == int32(0) {
						} else {
							v191 = v184
							v201 = v191
							v208 = v5
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v152+v201<<(uint(int32(2))%32)))) = v70
								v214 = int32(1)
								v217 = v208 + v214
								if v217 != int32(4) {
									v201 = v201 + v214
									v208 = v217
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v191 = v153
						v201 = v191
						v208 = v5
						for {
							*(*int32)(unsafe.Add(mBase, uint32(v152+v201<<(uint(int32(2))%32)))) = v70
							v214 = int32(1)
							v217 = v208 + v214
							if v217 != int32(4) {
								v201 = v201 + v214
								v208 = v217
								continue
							} else {
								break
							}
							break
						}
					}
					v231 = v70 - int32(1)
					v232 = int32(3)
					v233 = v231 & v232
					v234 = int32(0)
					if base.Ui32(v232) <= base.Ui32(v70-int32(2)) {
						v246 = v234
						v247 = int32(0)
						for {
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v118))))
							v256 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v254<<(uint(v256)%32)))) = v231 - v246
							v262 = v246 | int32(1)
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v262))))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v264<<(uint(v256)%32)))) = v231 - v262
							v272 = v246 | v256
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v272))))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v274<<(uint(v256)%32)))) = v231 - v272
							v282 = v246 | int32(3)
							v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118+v282))))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v284<<(uint(v256)%32)))) = v231 - v282
							v291 = int32(4)
							v292 = v246 + v291
							v294 = v247 + v291
							if v294 != v231&int32(-4) {
								v246 = v292
								v247 = v294
								continue
							} else {
								break
							}
							break
						}
						if v233 == int32(0) {
						} else {
							v301 = v292
							v311 = v301
							v316 = v234
							for {
								v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311+v118))))
								*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v319<<(uint(int32(2))%32)))) = v231 - v311
								v326 = int32(1)
								v329 = v316 + v326
								if v329 != v233 {
									v311 = v311 + v326
									v316 = v329
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v301 = v234
						v311 = v301
						v316 = v234
						for {
							v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311+v118))))
							*(*int32)(unsafe.Add(mBase, uint32(v152+v149&v319<<(uint(int32(2))%32)))) = v231 - v311
							v326 = int32(1)
							v329 = v316 + v326
							if v329 != v233 {
								v311 = v311 + v326
								v316 = v329
								continue
							} else {
								break
							}
							break
						}
					}
				}
			}
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v344 = m.ExcPending
		if v344 != 0 {
			return
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v347 = m.ExcPending
			if v347 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_text_position_setup_1), int32(0))
				mBase = m.M
				v351 = m.ExcPending
				if v351 != 0 {
					return
				} else {
					F_errhint(m, int32(_a_F_text_position_setup_2), int32(0))
					mBase = m.M
					v355 = m.ExcPending
					if v355 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_text_position_setup_3), int32(1337), int32(_a_F_text_position_setup_4))
						mBase = m.M
						v360 = m.ExcPending
						if v360 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_text_substr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v6 = F_text_substring(m, v2, v3, v4, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v6)
	}
}
func F_text_substring(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v356 int32
	_ = v356
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	v13 = int32(1)
	if l1 <= v13 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = v13
	goto L3
L2:
	;
	v16 = l1
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_text_substring[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_c_F_text_substring[1])))
	goto L11
L4:
	;
	v370 = F_palloc(m, int32(4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L22
	} else {
		goto L122
	}
L5:
	;
	if v346 != v81 {
		goto L118
	} else {
		goto L119
	}
L6:
	;
	if v16 <= v232 {
		goto L84
	} else {
		goto L85
	}
L7:
	;
	v194 = int32(1)
	v195 = v78 - v194
	if v191&v194 != 0 {
		goto L73
	} else {
		goto L74
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L22
	} else {
		goto L70
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L66
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L22
	} else {
		goto L62
	}
L11:
	;
	if v24 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v27 = int32(-1)
	if l3 != 0 {
		v38 = v27
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v24 < int32(2) {
		goto L8
	} else {
		goto L24
	}
L15:
	;
	v42 = F_detoast_attr_slice(m, base.I32_wrap_i64(l0), v16-int32(1), v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v29 = base.B2i32(l2 < int32(0))
	if l2 < int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v30 = l1 + l2
	if base.B2i32(v30 < l1)^v29 != 0 {
		v38 = v27
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v30 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	goto L4
L20:
	;
	goto L21
L21:
	;
	v38 = v30 - v16
	goto L15
L22:
	;
	return int32(0)
L23:
	;
	return v42
L24:
	;
	v49 = int32(-1)
	if l3 != 0 {
		v77 = v49
		v78 = v49
		v80 = v49
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = base.I32_wrap_i64(l0)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v82 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v53 = base.B2i32(l2 < int32(0))
	if l2 < int32(0) {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v55 = l1 + l2
	if base.B2i32(v55 < l1)^v53 != 0 {
		v77 = v49
		v78 = v55
		v80 = int32(-1)
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v55 <= int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	goto L4
L30:
	;
	goto L31
L31:
	;
	v65 = base.I64_extend_i32_s(v55-int32(1)) * base.I64_extend_i32_s(v24)
	v66 = base.I32_wrap_i64(v65)
	if base.I32_wrap_i64(int64(base.Ui64(v65)>>(uint(int64(32))%64))) != v66>>(uint(int32(31))%32) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v73 = int32(-1)
	goto L34
L33:
	;
	v73 = v66
	goto L34
L34:
	;
	v77 = v73
	v78 = v55
	v80 = v55 - v16
	goto L25
L35:
	;
	v138 = int32(1)
	if v135&v138 != 0 {
		goto L58
	} else {
		goto L59
	}
L36:
	;
	v131 = int32(4)
	v132 = int32(1)
	if v77 != int32(-1) {
		v191 = v132
		v192 = v131
		v193 = v91
		goto L7
	} else {
		goto L57
	}
L37:
	;
	if v123 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L38:
	;
	if v109&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	if v82&int32(3) != int32(2) {
		v109 = v82
		v110 = v81
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v91 = F_detoast_attr_slice(m, v81, int32(0), v77)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L22
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v93 != int32(1) {
		v109 = v93
		v110 = v91
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if base.Ui32((v96-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	if v96 == int32(18) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v107 = int32(16)
	goto L48
L47:
	;
	v107 = int32(0)
	goto L48
L48:
	;
	v122 = int32(1)
	v123 = v107
	v124 = v91
	goto L37
L49:
	;
	v113 = int32(1)
	v122 = v109
	v123 = int32(base.Ui32(v109)>>(uint(v113)%32)) - v113
	v124 = v110
	goto L37
L50:
	;
	goto L51
L51:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v122 = v109
	v123 = int32(base.Ui32(v117)>>(uint(int32(2))%32)) - int32(4)
	v124 = v110
	goto L37
L52:
	;
	v346 = v124
	goto L5
L53:
	;
	goto L54
L54:
	;
	if v77 == int32(-1) {
		v135 = v122
		v136 = v123
		v137 = v124
		goto L35
	} else {
		goto L55
	}
L55:
	;
	if v123 <= int32(0) {
		v346 = v124
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v191 = v122
	v192 = v123
	v193 = v124
	goto L7
L57:
	;
	v135 = v132
	v136 = v131
	v137 = v91
	goto L35
L58:
	;
	v142 = v138
	goto L60
L59:
	;
	v142 = int32(4)
	goto L60
L60:
	;
	v144 = F_pg_mbstrlen_with_len(m, v137+v142, v136)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	v231 = v137
	v232 = v144
	goto L6
L62:
	;
	F_errcode(m, int32(17039490))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L22
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_text_substring_0), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_text_substring_1), int32(612), int32(_a_F_text_substring_2))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	F_errcode(m, int32(17039490))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_text_substring_0), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_text_substring_1), int32(676), int32(_a_F_text_substring_2))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L22
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errmsg_internal(m, int32(_a_F_text_substring_3), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_text_substring_1), int32(790), int32(_a_F_text_substring_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L22
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v200 = v194
	goto L75
L74:
	;
	v200 = int32(4)
	goto L75
L75:
	;
	v205 = v193 + v200
	v206 = v192
	v208 = int32(0)
	goto L76
L76:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v215 == int32(0) {
		v231 = v193
		v232 = v208
		goto L6
	} else {
		goto L78
	}
L77:
	;
	v231 = v193
	v232 = v221
	goto L6
L78:
	;
	v218 = F_pg_mblen_with_len(m, v205, v206)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L22
	} else {
		goto L79
	}
L79:
	;
	v221 = v208 + int32(1)
	if v195 == v221 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v231 = v193
	v232 = v195
	goto L6
L81:
	;
	goto L82
L82:
	;
	v224 = v206 - v218
	if int32(0) < v224 {
		v205 = v205 + v218
		v206 = v224
		v208 = v221
		goto L76
	} else {
		goto L83
	}
L83:
	;
	goto L77
L84:
	;
	if int32(0) <= v80 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v346 = v231
	goto L5
L86:
	;
	v252 = int32(1)
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v254&v252 != 0 {
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v242 = v16 + v80
	v244 = v232 + int32(1)
	if v242 < v244 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	v251 = v232 + int32(1)
	goto L86
L90:
	;
	v246 = v242
	goto L92
L91:
	;
	v246 = v244
	goto L92
L92:
	;
	v251 = v246
	goto L86
L93:
	;
	v257 = v252
	goto L95
L94:
	;
	v257 = int32(4)
	goto L95
L95:
	;
	v258 = v231 + v257
	if int32(2) <= l1 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v266 = v258
	v267 = int32(0)
	goto L99
L97:
	;
	v286 = v258
	goto L98
L98:
	;
	if v16 < v251 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v276 = F_pg_mblen_unbounded(m, v266)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L22
	} else {
		goto L101
	}
L100:
	;
	v286 = v278
	goto L98
L101:
	;
	v278 = v276 + v266
	if base.B2i32(v16-int32(2) == v267) == int32(0) {
		v266 = v278
		v267 = v267 + int32(1)
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v300 = v286
	v304 = v16
	goto L106
L104:
	;
	v318 = v286
	goto L105
L105:
	;
	v327 = v318 - v286
	v329 = v327 + int32(4)
	v330 = F_palloc(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L22
	} else {
		goto L110
	}
L106:
	;
	v309 = F_pg_mblen_unbounded(m, v300)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L22
	} else {
		goto L108
	}
L107:
	;
	v318 = v311
	goto L105
L108:
	;
	v311 = v309 + v300
	v313 = v304 + int32(1)
	if v313 != v251 {
		v300 = v311
		v304 = v313
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v329 << (uint(int32(2)) % 32)
	if v327 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	base.MemoryCopy(m, v330+int32(4), v286, v327)
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v231 != v81 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_pfree(m, v231)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L22
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	return v330
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, v346)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L22
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	goto L4
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v370))) = int32(16)
	return v370
}
func F_text_to_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	v3 = int64(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v3
	v12 = F_split_text(m, l0, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v32 = v3
			m.G0 = v6 + int32(16)
			return v32
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v20 == int32(0) {
				v24 = F_construct_empty_array(m, int32(25))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v32 = base.I64_extend_i32_u(v24)
					m.G0 = v6 + int32(16)
					return v32
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_text_to_array[0]))
				v29 = F_makeArrayResult(m, v20, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v32 = v29
					m.G0 = v6 + int32(16)
					return v32
				}
			}
		}
	}
}
