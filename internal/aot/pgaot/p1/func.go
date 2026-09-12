package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LookupFuncName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v19 = F_FuncnameGetCandidates(m, l0, l1, v5, v5, v5, v5, l3)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v189
L2:
	;
	v173 = F_NameListToString(m, l0)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L55
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L49
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L43
	}
L5:
	;
	return int32(0)
L6:
	;
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v24 = l1 << (uint(int32(2)) % 32)
	v32 = v5
	v33 = v19
	goto L10
L8:
	;
	goto L9
L9:
	;
	if l3 != 0 {
		v189 = v5
		goto L1
	} else {
		goto L42
	}
L10:
	;
	if base.B2i32(l1 <= int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	if l3 != 0 {
		v189 = v111
		goto L1
	} else {
		goto L40
	}
L12:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v113 != 0 {
		v32 = v111
		v33 = v113
		goto L10
	} else {
		goto L39
	}
L13:
	;
	v40 = v33 + int32(32)
	if base.Ui32(int32(4)) <= base.Ui32(v24) {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	goto L15
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v103 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L16:
	;
	if v102 != 0 {
		v111 = v32
		goto L12
	} else {
		goto L34
	}
L17:
	;
	v102 = int32(0)
	goto L16
L18:
	;
	v76 = v71
	v77 = v72
	v78 = v73
	goto L28
L19:
	;
	if (l2|v40)&int32(3) != 0 {
		v71 = l2
		v72 = v40
		v73 = v24
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v64 = l2
	v65 = v40
	v66 = v24
	goto L21
L21:
	;
	if v66 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v48 = l2
	v49 = v40
	v50 = v24
	goto L23
L23:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v53 != v54 {
		v71 = v48
		v72 = v49
		v73 = v50
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v64 = v59
	v65 = v57
	v66 = v61
	goto L21
L25:
	;
	v56 = int32(4)
	v57 = v49 + v56
	v59 = v48 + v56
	v61 = v50 - v56
	if base.Ui32(int32(3)) < base.Ui32(v61) {
		v48 = v59
		v49 = v57
		v50 = v61
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v71 = v64
	v72 = v65
	v73 = v66
	goto L18
L28:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v81 == v82 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v102 = v81 - v82
	goto L16
L30:
	;
	v84 = int32(1)
	v89 = v78 - v84
	if v89 != 0 {
		v76 = v76 + v84
		v77 = v77 + v84
		v78 = v89
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	goto L15
L35:
	;
	v106 = F_get_func_prokind(m, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v106 == int32(112) {
		v111 = v32
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v32 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v111 = v110
	goto L12
L39:
	;
	goto L11
L40:
	;
	if v111 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v189 = v111
	goto L1
L42:
	;
	goto L4
L43:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	if l1 < int32(0) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v136 = F_func_signature_string(m, l0, l1, int32(0), l2)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v136
	F_errmsg(m, int32(69739), v13+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(499958), int32(2174), int32(382299))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v156 = F_NameListToString(m, l0)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v156
	F_errmsg(m, int32(344432), v13+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(575024), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(499958), int32(2183), int32(382299))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v173
	F_errmsg(m, int32(723116), v13)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(499958), int32(2168), int32(382299))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_func_arg_info(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v19 = F_SysCacheGetAttr(m, int32(47), l0, int32(21), v11+int32(15))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v23 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L66
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L63
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L60
	}
L6:
	;
	v74 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v11+int32(15))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L28
	}
L7:
	;
	v26 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v55 = v13 + v14
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+128))
	v58 = v56 << (uint(int32(2)) % 32)
	v59 = F_palloc(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v28 != int32(1) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	if v31 < int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v34 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v35 != int32(26) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v39 = v31 << (uint(int32(2)) % 32)
	v40 = F_palloc(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v51 = v43
	goto L18
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v51 = (v44<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	if v39 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v68 = v31
	goto L6
L20:
	;
	v53 = F__emscripten_memcpy_bulkmem(m, v40, v51+v26, v39)
	mBase = m.M
	goto L22
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v68 = v56
	goto L6
L25:
	;
	v64 = F__emscripten_memcpy_bulkmem(m, v59, v55+int32(136), v58)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v76 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v134 = F_SysCacheGetAttr(m, int32(47), l0, int32(22), v11+int32(15))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L42
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v82 = F_pg_detoast_datum(m, v74)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_deconstruct_array_builtin(m, v82, int32(25), v11+int32(8), int32(0), v11+int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v92 != v68 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v96 = F_palloc(m, v68<<(uint(int32(2))%32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v96
	if v68 <= int32(0) {
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v105 = int32(0)
	goto L38
L38:
	;
	v110 = v105 << (uint(int32(2)) % 32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110+v111)))
	v114 = F_text_to_cstring(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L29
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v116+v110))) = v114
	v120 = v105 + int32(1)
	if v120 != v68 {
		v105 = v120
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v136 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	m.G0 = v11 + int32(16)
	return v68
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v141 = F_pg_detoast_datum(m, v134)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v143 != int32(1) {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	if v146 != v68 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v148 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	if v149 != int32(18) {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v152 = F_palloc(m, v68)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v152
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v141)+8))
	if v155 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v163 = v155
	goto L55
L54:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v163 = (v156<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L55
L55:
	;
	if v68 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L43
L57:
	;
	v165 = F__emscripten_memcpy_bulkmem(m, v152, v163+v141, v68)
	mBase = m.M
	goto L59
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	F_errmsg_internal(m, int32(152141), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(497683), int32(1411), int32(241706))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(121252), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(497683), int32(1438), int32(241706))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v68
	F_errmsg_internal(m, int32(152326), v11)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(497683), int32(1458), int32(241706))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_func_namespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+68))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_func_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+92))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
