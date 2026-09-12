package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_compare_text_lexemes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(1)
	v11 = v9 + v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v14 = v12 & v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v12 == v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = int32(1)
	v45 = v15 + v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v48 = v46 & v44
	if v46 == v44 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v18 = int32(4)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v20&int32(254) == int32(2) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v33 = int32(1)
	if v14 != 0 {
		v43 = int32(base.Ui32(v12)>>(uint(v33)%32)) - v33
		goto L1
	} else {
		goto L11
	}
L5:
	;
	v29 = v18
	goto L7
L6:
	;
	v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
	goto L7
L7:
	;
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = v18
	goto L10
L9:
	;
	v32 = v29
	goto L10
L10:
	;
	v43 = v32
	goto L1
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L12:
	;
	if v43 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v51 = int32(4)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v53&int32(254) == int32(2) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v66 = int32(1)
	if v48 != 0 {
		v76 = int32(base.Ui32(v46)>>(uint(v66)%32)) - v66
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v62 = v51
	goto L18
L17:
	;
	v62 = base.B2i32(v53 == int32(18)) << (uint(v51) % 32)
	goto L18
L18:
	;
	if v53 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v65 = v51
	goto L21
L20:
	;
	v65 = v62
	goto L21
L21:
	;
	v76 = v65
	goto L12
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L12
L23:
	;
	v80 = int32(0)
	if v80 < v76 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v76 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v83 = int32(-1)
	goto L28
L27:
	;
	v83 = v80
	goto L28
L28:
	;
	return v83
L29:
	;
	return base.B2i32(int32(0) < v43)
L30:
	;
	goto L31
L31:
	;
	if v14 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	return v166
L33:
	;
	v92 = v11
	goto L35
L34:
	;
	v92 = v9 + int32(4)
	goto L35
L35:
	;
	if v48 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v95 = v45
	goto L38
L37:
	;
	v95 = v15 + int32(4)
	goto L38
L38:
	;
	if base.Ui32(v43) < base.Ui32(v76) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v97 = v43
	goto L41
L40:
	;
	v97 = v76
	goto L41
L41:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v97) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if v159 != 0 {
		v166 = v159
		goto L32
	} else {
		goto L60
	}
L43:
	;
	v159 = int32(0)
	goto L42
L44:
	;
	v133 = v128
	v134 = v129
	v135 = v130
	goto L54
L45:
	;
	if (v92|v95)&int32(3) != 0 {
		v128 = v92
		v129 = v95
		v130 = v97
		goto L44
	} else {
		goto L48
	}
L46:
	;
	v121 = v92
	v122 = v95
	v123 = v97
	goto L47
L47:
	;
	if v123 == int32(0) {
		goto L43
	} else {
		goto L53
	}
L48:
	;
	v105 = v92
	v106 = v95
	v107 = v97
	goto L49
L49:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v110 != v111 {
		v128 = v105
		v129 = v106
		v130 = v107
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v121 = v116
	v122 = v114
	v123 = v118
	goto L47
L51:
	;
	v113 = int32(4)
	v114 = v106 + v113
	v116 = v105 + v113
	v118 = v107 - v113
	if base.Ui32(int32(3)) < base.Ui32(v118) {
		v105 = v116
		v106 = v114
		v107 = v118
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v128 = v121
	v129 = v122
	v130 = v123
	goto L44
L54:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v138 == v139 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v159 = v138 - v139
	goto L42
L56:
	;
	v141 = int32(1)
	v146 = v135 - v141
	if v146 != 0 {
		v133 = v133 + v141
		v134 = v134 + v141
		v135 = v146
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
	if v76 == v43 {
		v166 = int32(0)
		goto L32
	} else {
		goto L61
	}
L61:
	;
	if v43 < v76 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v165 = int32(-1)
	goto L64
L63:
	;
	v165 = int32(1)
	goto L64
L64:
	;
	v166 = v165
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
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
	v47 = F_palloc(m, v44+int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
	v17 = int32(4)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v19&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(1)
	if v14&v32 != 0 {
		v44 = int32(base.Ui32(v14)>>(uint(v32)%32)) - v32
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v28 = v17
	goto L9
L8:
	;
	v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
	goto L9
L9:
	;
	if v19 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v17
	goto L12
L11:
	;
	v31 = v28
	goto L12
L12:
	;
	v44 = v31
	goto L1
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	v49 = int32(1)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v51&v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = v49
	goto L17
L16:
	;
	v54 = int32(4)
	goto L17
L17:
	;
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44+v57))) = uint8(v59)
	if l0 != v10 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v56 = F__emscripten_memcpy_bulkmem(m, v47, v10+v54, v44)
	mBase = m.M
	v57 = v56
	goto L21
L20:
	;
	v57 = v47
	goto L21
L21:
	;
	goto L18
L22:
	;
	F_pfree(m, v10)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v67 = F_SplitIdentifierString(m, v57, int32(46), v8+int32(12))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L47
	}
L27:
	;
	if v67 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v69 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L43
	}
L31:
	;
	v72 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v72 < v73 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v77 = v72
	v79 = int32(0)
	goto L35
L33:
	;
	v97 = v72
	goto L34
L34:
	;
	F_pfree(m, v57)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L41
	}
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v79<<(uint(int32(2))%32))))
	v87 = F_pstrdup(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	v97 = v91
	goto L34
L37:
	;
	v89 = F_makeString(m, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v91 = F_lappend(m, v77, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v94 = v79 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v94 < v95 {
		v77 = v91
		v79 = v94
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	F_list_free(m, v69)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v8 + int32(16)
	return v97
L43:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(28076), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(483184), int32(3537), int32(73926))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(28076), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(483184), int32(3542), int32(73926))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_text_catenate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v9 == int32(1) {
		v12 = int32(4)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v14&int32(254) == int32(2) {
			v23 = v12
		} else {
			v23 = base.B2i32(v14 == int32(18)) << (uint(v12) % 32)
		}
		if v14 == int32(1) {
			v26 = v12
		} else {
			v26 = v23
		}
		v39 = v26
	} else {
		v27 = int32(1)
		if v9&v27 != 0 {
			v39 = int32(base.Ui32(v9)>>(uint(v27)%32)) - v27
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = int32(base.Ui32(v33)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v40 = int32(0)
	if v40 < v39 {
		v43 = v39
	} else {
		v43 = v40
	}
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v44 == int32(1) {
		v47 = int32(4)
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v49&int32(254) == int32(2) {
			v58 = v47
		} else {
			v58 = base.B2i32(v49 == int32(18)) << (uint(v47) % 32)
		}
		if v49 == int32(1) {
			v61 = v47
		} else {
			v61 = v58
		}
		v74 = v61
	} else {
		v62 = int32(1)
		if v44&v62 != 0 {
			v74 = int32(base.Ui32(v44)>>(uint(v62)%32)) - v62
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v75 = int32(0)
	if v75 < v74 {
		v78 = v74
	} else {
		v78 = v75
	}
	v81 = v43 + v78 + int32(4)
	v82 = F_palloc(m, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v82))) = v81 << (uint(int32(2)) % 32)
		v90 = v82 + int32(4)
		if int32(0) < v39 {
			v93 = int32(1)
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v95&v93 != 0 {
				v98 = v93
			} else {
				v98 = int32(4)
			}
			if v43 != 0 {
				v100 = F__emscripten_memcpy_bulkmem(m, v90, l0+v98, v43)
				mBase = m.M
			} else {
			}
		} else {
		}
		if int32(0) < v74 {
			v105 = int32(1)
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			if v107&v105 != 0 {
				v110 = v105
			} else {
				v110 = int32(4)
			}
			if v78 != 0 {
				v112 = F__emscripten_memcpy_bulkmem(m, v90+v43, l1+v110, v78)
				mBase = m.M
			} else {
			}
		} else {
		}
		return v82
	}
}
func F_text_pattern_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v51 == int32(1) {
				v54 = int32(4)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v56&int32(254) == int32(2) {
					v65 = v54
				} else {
					v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
				}
				if v56 == int32(1) {
					v68 = v54
				} else {
					v68 = v65
				}
				v81 = v68
			} else {
				v69 = int32(1)
				if v51&v69 != 0 {
					v81 = int32(base.Ui32(v51)>>(uint(v69)%32)) - v69
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v82 = int32(1)
			if v20&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = int32(1)
			if v51&v88 != 0 {
				v92 = v88
			} else {
				v92 = int32(4)
			}
			v94 = base.B2i32(v50 < v81)
			if v50 < v81 {
				v95 = v50
			} else {
				v95 = v81
			}
			v96 = F_memcmp(m, v7+v86, v14+v92, v95)
			mBase = m.M
			if v96 != 0 {
				v99 = v96
			} else {
				if v50 < v81 {
					v99 = int32(-1)
				} else {
					v99 = base.B2i32(v81 < v50)
				}
			}
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v100 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v104 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v99^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v99^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v104 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v99^int32(-1)) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v99^int32(-1)) >> (uint(int32(31)) % 32))
				}
			}
		}
	}
}
func F_text_pattern_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v51 == int32(1) {
				v54 = int32(4)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v56&int32(254) == int32(2) {
					v65 = v54
				} else {
					v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
				}
				if v56 == int32(1) {
					v68 = v54
				} else {
					v68 = v65
				}
				v81 = v68
			} else {
				v69 = int32(1)
				if v51&v69 != 0 {
					v81 = int32(base.Ui32(v51)>>(uint(v69)%32)) - v69
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v82 = int32(1)
			if v20&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = int32(1)
			if v51&v88 != 0 {
				v92 = v88
			} else {
				v92 = int32(4)
			}
			v94 = base.B2i32(v50 < v81)
			if v50 < v81 {
				v95 = v50
			} else {
				v95 = v81
			}
			v96 = F_memcmp(m, v7+v86, v14+v92, v95)
			mBase = m.M
			if v96 != 0 {
				v99 = v96
			} else {
				if v50 < v81 {
					v99 = int32(-1)
				} else {
					v99 = base.B2i32(v81 < v50)
				}
			}
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v100 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v104 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v104 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v99) >> (uint(int32(31)) % 32))
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
	var v27 int32
	_ = v27
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v282 int32
	_ = v282
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 <= v2 {
		v282 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v282
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
	v27 = v14
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1060)) = v27
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v40 = v38 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v43 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v221 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L10:
	;
	if v205 == int32(0) {
		v282 = v2
		goto L1
	} else {
		goto L57
	}
L11:
	;
	if base.Ui32(v40) <= base.Ui32(v25) {
		v282 = v2
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v27 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L14:
	;
	v48 = v25
	goto L15
L15:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	v71 = v48
	v74 = int32(0)
	goto L23
L18:
	;
	if v40-v48 < v27 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = F_pg_strncoll(m, v48, v27, v41, v27, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v205 = v48
	goto L10
L23:
	;
	v82 = F_pg_mblen_range(m, v71, v40)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L20
	} else {
		goto L27
	}
L24:
	;
	if v95 != 0 {
		v213 = v95
		goto L9
	} else {
		goto L32
	}
L25:
	;
	goto L24
L26:
	;
	if base.Ui32(v84) < base.Ui32(v40) {
		v71 = v84
		v74 = v93
		goto L23
	} else {
		goto L31
	}
L27:
	;
	v84 = v82 + v71
	v85 = v84 - v48
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = F_pg_strncoll(m, v48, v85, v41, v27, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	if v87 != 0 {
		v93 = v74
		goto L26
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1060)) = v85
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v90 == int32(1) {
		v93 = v48
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v95 = v48
	goto L25
L31:
	;
	v95 = v93
	goto L25
L32:
	;
	v96 = F_pg_mblen_range(m, v48, v40)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v98 = v96 + v48
	if base.Ui32(v98) < base.Ui32(v40) {
		v48 = v98
		goto L15
	} else {
		goto L34
	}
L34:
	;
	goto L16
L35:
	;
	if base.Ui32(v40) <= base.Ui32(v25) {
		v282 = v2
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v126 = v25 + v27 - int32(1)
	if base.Ui32(v40) <= base.Ui32(v126) {
		v282 = v2
		goto L1
	} else {
		goto L45
	}
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v107 = v25
	goto L39
L39:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v105 == v119 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v282 = v2
	goto L1
L41:
	;
	v213 = v107
	goto L9
L42:
	;
	goto L43
L43:
	;
	v122 = v107 + int32(1)
	if v122 != v40 {
		v107 = v122
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v129 = int32(1)
	v133 = v27 + v41 - v129
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v141 = v126
	goto L46
L46:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v148 == v134 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v205 = v141 + (v129 - v27)
	goto L10
L48:
	;
	goto L47
L49:
	;
	v151 = v141
	v152 = v133
	goto L52
L50:
	;
	goto L51
L51:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28)+v148&v128<<(uint(int32(2))%32))))
	v189 = v141 + v188
	if base.Ui32(v189) < base.Ui32(v40) {
		v141 = v189
		goto L46
	} else {
		goto L56
	}
L52:
	;
	if v152 == v41 {
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v164 = int32(1)
	v165 = v152 - v164
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v168 = v151 - v164
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v166 == v169 {
		v151 = v168
		v152 = v165
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v282 = v2
	goto L1
L57:
	;
	v213 = v205
	goto L9
L58:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = v249
	v27 = v273
	goto L7
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1052)) = v213
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1060))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1056)) = v270
	v282 = int32(1)
	goto L1
L60:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v225 != int32(1) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1064))
	v233 = v231
	goto L62
L62:
	;
	if base.Ui32(v213) <= base.Ui32(v233) {
		goto L59
	} else {
		goto L64
	}
L63:
	;
	goto L58
L64:
	;
	v246 = F_pg_mblen_range(m, v233, v228+v229)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1064))
	v249 = v246 + v248
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1064)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1068))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1068)) = v251 + int32(1)
	if base.Ui32(v249) <= base.Ui32(v213) {
		v233 = v249
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
}
func F_text_position_setup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	v5 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		v14 = int32(4)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v16&int32(254) == int32(2) {
			v25 = v14
		} else {
			v25 = base.B2i32(v16 == int32(18)) << (uint(v14) % 32)
		}
		if v16 == int32(1) {
			v28 = v14
		} else {
			v28 = v25
		}
		v41 = v28
	} else {
		v29 = int32(1)
		if v11&v29 != 0 {
			v41 = int32(base.Ui32(v11)>>(uint(v29)%32)) - v29
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v42 == int32(1) {
		v45 = int32(4)
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v47&int32(254) == int32(2) {
			v56 = v45
		} else {
			v56 = base.B2i32(v47 == int32(18)) << (uint(v45) % 32)
		}
		if v47 == int32(1) {
			v59 = v45
		} else {
			v59 = v56
		}
		v72 = v59
	} else {
		v60 = int32(1)
		if v42&v60 != 0 {
			v72 = int32(base.Ui32(v42)>>(uint(v60)%32)) - v60
		} else {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v72 = int32(base.Ui32(v66)>>(uint(int32(2))%32)) - int32(4)
		}
	}
	if l2 != 0 {
		v73 = F_pg_newlocale_from_collation(m, l2)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			v75 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v75)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v73
			v79 = *(*int32)(unsafe.Add(mBase, _consts[451]))
			v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
			v85 = *(*int32)(unsafe.Add(mBase, uint32(v80*int32(28))+uint32(_consts[970])))
			if v85 == int32(1) {
				v88 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v88)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, _consts[451]))
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				if v92 == int32(6) {
					v95 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v95)
				} else {
					v97 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v97)
				}
			}
			v99 = int32(1)
			v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v101&v99 != 0 {
				v104 = v99
			} else {
				v104 = int32(4)
			}
			v105 = l0 + v104
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v105
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
			v108 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1068)) = v108
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1064)) = v105
			*(*int32)(unsafe.Add(mBase, uint32(l3)+1052)) = v108
			*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v72
			*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v41
			v115 = int32(1)
			if v107&v115 != 0 {
				v119 = v115
			} else {
				v119 = int32(4)
			}
			v120 = l1 + v119
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v120
			if v41 < v72 {
			} else {
				if v72 < int32(2) {
				} else {
					v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+1)))
					if v126 != int32(1) {
					} else {
						v130 = v41 - v72
						if v130 < int32(16) {
							v150 = int32(3)
						} else {
							if base.Ui32(v130) < base.Ui32(int32(64)) {
								v150 = int32(7)
							} else {
								if base.Ui32(v130) < base.Ui32(int32(128)) {
									v150 = int32(15)
								} else {
									if base.Ui32(v130) < base.Ui32(int32(512)) {
										v150 = int32(31)
									} else {
										if base.Ui32(v130) < base.Ui32(int32(2048)) {
											v150 = int32(63)
										} else {
											if base.Ui32(v130) < base.Ui32(int32(4096)) {
												v149 = int32(127)
											} else {
												v149 = int32(255)
											}
											v150 = v149
										}
									}
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v150
						v153 = l3 + int32(28)
						v155 = v150 + int32(1)
						v157 = v155 & int32(4)
						v158 = int32(0)
						if base.Ui32(int32(7)) <= base.Ui32(v150) {
							v164 = v158
							v169 = v5
							for {
								v175 = v153 + v164<<(uint(int32(2))%32)
								*(*int32)(unsafe.Add(mBase, uint32(v175))) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+20)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+24)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v175)+28)) = v72
								v184 = int32(8)
								v185 = v164 + v184
								v187 = v169 + v184
								if v187 != v155&int32(504) {
									v164 = v185
									v169 = v187
									continue
								} else {
									break
								}
								break
							}
							v190 = v185
						} else {
							v190 = v158
						}
						if v157 != 0 {
							v200 = v190
							v207 = v5
							for {
								*(*int32)(unsafe.Add(mBase, uint32(v153+v200<<(uint(int32(2))%32)))) = v72
								v213 = int32(1)
								v216 = v207 + v213
								if v216 != v157 {
									v200 = v200 + v213
									v207 = v216
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v229 = v72 - int32(1)
						v230 = int32(3)
						v231 = v229 & v230
						v232 = int32(0)
						if base.Ui32(v230) <= base.Ui32(v72-int32(2)) {
							v244 = v232
							v245 = int32(0)
							for {
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v244))))
								v254 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v153+v150&v252<<(uint(v254)%32)))) = v229 - v244
								v260 = v244 | int32(1)
								v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v260))))
								*(*int32)(unsafe.Add(mBase, uint32(v153+v150&v262<<(uint(v254)%32)))) = v229 - v260
								v270 = v244 | v254
								v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v270))))
								*(*int32)(unsafe.Add(mBase, uint32(v153+v150&v272<<(uint(v254)%32)))) = v229 - v270
								v280 = v244 | int32(3)
								v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v280))))
								*(*int32)(unsafe.Add(mBase, uint32(v153+v150&v282<<(uint(v254)%32)))) = v229 - v280
								v289 = int32(4)
								v290 = v244 + v289
								v292 = v245 + v289
								if v292 != v229&int32(-4) {
									v244 = v290
									v245 = v292
									continue
								} else {
									break
								}
								break
							}
							v297 = v290
						} else {
							v297 = v232
						}
						if v231 == int32(0) {
						} else {
							v309 = v297
							v314 = v232
							for {
								v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120+v309))))
								*(*int32)(unsafe.Add(mBase, uint32(v153+v150&v317<<(uint(int32(2))%32)))) = v229 - v309
								v324 = int32(1)
								v327 = v314 + v324
								if v327 != v231 {
									v309 = v309 + v324
									v314 = v327
									continue
								} else {
									break
								}
								break
							}
						}
					}
				}
			}
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v342 = m.ExcPending
		if v342 != 0 {
			return
		} else {
			F_errcode(m, int32(34209924))
			mBase = m.M
			v345 = m.ExcPending
			if v345 != 0 {
				return
			} else {
				F_errmsg(m, int32(236373), int32(0))
				mBase = m.M
				v349 = m.ExcPending
				if v349 != 0 {
					return
				} else {
					F_errhint(m, int32(541241), int32(0))
					mBase = m.M
					v353 = m.ExcPending
					if v353 != 0 {
						return
					} else {
						F_errfinish(m, int32(483184), int32(1648), int32(100695))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
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
func F_text_substr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = F_text_substring(m, v2, v3, v4, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_text_substring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
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
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	v12 = int32(1)
	if l1 <= v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v12
	goto L3
L2:
	;
	v15 = l1
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[451]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18*int32(28))+uint32(_consts[970])))
	goto L11
L4:
	;
	if int32(0) <= v92 {
		goto L95
	} else {
		goto L96
	}
L5:
	;
	if l0 != v262 {
		goto L89
	} else {
		goto L90
	}
L6:
	;
	if v15 <= v249 {
		goto L4
	} else {
		goto L88
	}
L7:
	;
	v213 = int32(1)
	v214 = v90 - v213
	if v210&v213 != 0 {
		goto L77
	} else {
		goto L78
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L22
	} else {
		goto L74
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L22
	} else {
		goto L70
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L22
	} else {
		goto L66
	}
L11:
	;
	if v23 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v26 = int32(-1)
	if l3 != 0 {
		v46 = v26
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v23 < int32(2) {
		goto L8
	} else {
		goto L25
	}
L15:
	;
	v49 = F_pg_detoast_datum_slice(m, l0, v15-int32(1), v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L22
	} else {
		goto L24
	}
L16:
	;
	if l2 < int32(0) {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v31 = l1 + l2
	if base.B2i32(l2 < int32(0))^base.B2i32(v31 < l1) != 0 {
		v46 = v26
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v31 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v37 = F_palloc(m, int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v46 = v31 - v15
	goto L15
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = int32(16)
	return v37
L24:
	;
	return v49
L25:
	;
	v54 = int32(-1)
	if l3 != 0 {
		v89 = v54
		v90 = v54
		v92 = v54
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v93 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L27:
	;
	if l2 < int32(0) {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v62 = l1 + l2
	if base.B2i32(l2 < int32(0))^base.B2i32(v62 < l1) != 0 {
		v89 = v54
		v90 = v62
		v92 = int32(-1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	if v62 <= int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v68 = F_palloc(m, int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L22
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v78 = base.I64_extend_i32_s(v62-int32(1)) * base.I64_extend_i32_s(v23)
	v79 = base.I32_wrap_i64(v78)
	if base.I32_wrap_i64(int64(base.Ui64(v78)>>(uint(int64(32))%64))) != v79>>(uint(int32(31))%32) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = int32(16)
	return v68
L34:
	;
	v86 = int32(-1)
	goto L36
L35:
	;
	v86 = v79
	goto L36
L36:
	;
	v89 = v86
	v90 = v62
	v92 = v62 - v15
	goto L26
L37:
	;
	v157 = int32(1)
	if v154&v157 != 0 {
		goto L62
	} else {
		goto L63
	}
L38:
	;
	v151 = int32(1)
	if v89 != int32(-1) {
		v210 = v151
		v211 = v107
		v212 = v102
		goto L7
	} else {
		goto L61
	}
L39:
	;
	if v134 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	if v120&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	if v93&int32(3) != int32(2) {
		v120 = v93
		v121 = l0
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v102 = F_pg_detoast_datum_slice(m, l0, int32(0), v89)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L22
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v104 != int32(1) {
		v120 = v104
		v121 = v102
		goto L40
	} else {
		goto L46
	}
L46:
	;
	v107 = int32(4)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if base.Ui32((v108-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v133 = int32(1)
	v134 = base.B2i32(v108 == int32(18)) << (uint(int32(4)) % 32)
	v135 = v102
	goto L39
L48:
	;
	v124 = int32(1)
	v133 = v120
	v134 = int32(base.Ui32(v120)>>(uint(v124)%32)) - v124
	v135 = v121
	goto L39
L49:
	;
	goto L50
L50:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v133 = v120
	v134 = int32(base.Ui32(v128)>>(uint(int32(2))%32)) - int32(4)
	v135 = v121
	goto L39
L51:
	;
	if l0 != v135 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v89 == int32(-1) {
		v154 = v133
		v155 = v134
		v156 = v135
		goto L37
	} else {
		goto L59
	}
L54:
	;
	F_pfree(m, v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L22
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v142 = F_palloc(m, int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L22
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = int32(16)
	return v142
L59:
	;
	if v134 <= int32(0) {
		v262 = v135
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v210 = v133
	v211 = v134
	v212 = v135
	goto L7
L61:
	;
	v154 = v151
	v155 = v107
	v156 = v102
	goto L37
L62:
	;
	v161 = v157
	goto L64
L63:
	;
	v161 = int32(4)
	goto L64
L64:
	;
	v163 = F_pg_mbstrlen_with_len(m, v156+v161, v155)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	v249 = v163
	v250 = v156
	goto L6
L66:
	;
	F_errcode(m, int32(17039490))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(424391), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(483184), int32(923), int32(317683))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
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
	F_errcode(m, int32(17039490))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(424391), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L22
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(483184), int32(987), int32(317683))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(536094), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L22
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(483184), int32(1101), int32(317683))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L22
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v219 = v213
	goto L79
L78:
	;
	v219 = int32(4)
	goto L79
L79:
	;
	v224 = v212 + v219
	v225 = v211
	v226 = int32(0)
	goto L80
L80:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v233 == int32(0) {
		v249 = v226
		v250 = v212
		goto L6
	} else {
		goto L82
	}
L81:
	;
	v249 = v239
	v250 = v212
	goto L6
L82:
	;
	v236 = F_pg_mblen_with_len(m, v224, v225)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L22
	} else {
		goto L83
	}
L83:
	;
	v239 = v226 + int32(1)
	if v214 == v239 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v249 = v214
	v250 = v212
	goto L6
L85:
	;
	goto L86
L86:
	;
	v242 = v225 - v236
	if int32(0) < v242 {
		v224 = v224 + v236
		v225 = v242
		v226 = v239
		goto L80
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	v262 = v250
	goto L5
L89:
	;
	F_pfree(m, v262)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L22
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v272 = F_palloc(m, int32(4))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L22
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v272))) = int32(16)
	return v272
L94:
	;
	v289 = int32(1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v291&v289 != 0 {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v279 = v15 + v92
	v281 = v249 + int32(1)
	if v279 < v281 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v288 = v249 + int32(1)
	goto L94
L98:
	;
	v283 = v279
	goto L100
L99:
	;
	v283 = v281
	goto L100
L100:
	;
	v288 = v283
	goto L94
L101:
	;
	v294 = v289
	goto L103
L102:
	;
	v294 = int32(4)
	goto L103
L103:
	;
	v295 = v250 + v294
	if int32(2) <= l1 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v303 = v295
	v304 = int32(0)
	goto L107
L105:
	;
	v322 = v295
	goto L106
L106:
	;
	if v15 < v288 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v312 = F_pg_mblen_unbounded(m, v303)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L22
	} else {
		goto L109
	}
L108:
	;
	v322 = v314
	goto L106
L109:
	;
	v314 = v312 + v303
	if base.B2i32(v304 == v15-int32(2)) == int32(0) {
		v303 = v314
		v304 = v304 + int32(1)
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v335 = v322
	v339 = v15
	goto L114
L112:
	;
	v352 = v322
	goto L113
L113:
	;
	v360 = v352 - v322
	v362 = v360 + int32(4)
	v363 = F_palloc(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L22
	} else {
		goto L118
	}
L114:
	;
	v343 = F_pg_mblen_unbounded(m, v335)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L22
	} else {
		goto L116
	}
L115:
	;
	v352 = v345
	goto L113
L116:
	;
	v345 = v343 + v335
	v347 = v339 + int32(1)
	if v347 != v288 {
		v335 = v345
		v339 = v347
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = v362 << (uint(int32(2)) % 32)
	if v360 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if l0 != v250 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	v370 = F__emscripten_memcpy_bulkmem(m, v363+int32(4), v322, v360)
	mBase = m.M
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	F_pfree(m, v250)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L22
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	return v363
L126:
	;
	goto L125
}
func F_text_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(0)
	v12 = F_split_text(m, l0, v6)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v31 = v2
			m.G0 = v6 + int32(16)
			return v31
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v20 == int32(0) {
				v24 = F_construct_empty_array(m, int32(25))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v31 = v24
					m.G0 = v6 + int32(16)
					return v31
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v28 = F_makeArrayResult(m, v20, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = v28
					m.G0 = v6 + int32(16)
					return v31
				}
			}
		}
	}
}
