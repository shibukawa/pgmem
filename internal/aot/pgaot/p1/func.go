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
	var v190 int32
	_ = v190
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
	return v190
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
	v32 = v19
	v33 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	if l3 != 0 {
		v190 = v5
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
		v190 = v111
		goto L1
	} else {
		goto L40
	}
L12:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if v113 != 0 {
		v32 = v113
		v33 = v111
		goto L10
	} else {
		goto L39
	}
L13:
	;
	v40 = v32 + int32(32)
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
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v103 == int32(0) {
		goto L3
	} else {
		goto L35
	}
L16:
	;
	if v102 != 0 {
		v111 = v33
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
		v111 = v33
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v33 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
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
	v190 = v111
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
	F_errmsg(m, int32(_a_F_LookupFuncName_0), v13+int32(16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2174), int32(_a_F_LookupFuncName_2))
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
	F_errmsg(m, int32(_a_F_LookupFuncName_3), v13+int32(32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	F_errhint(m, int32(_a_F_LookupFuncName_4), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2183), int32(_a_F_LookupFuncName_2))
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
	F_errmsg(m, int32(_a_F_LookupFuncName_5), v13)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_LookupFuncName_1), int32(2168), int32(_a_F_LookupFuncName_2))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
	v20 = F_SysCacheGetAttr(m, int32(47), l0, int32(21), v12+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L57
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L51
	}
L6:
	;
	v80 = F_SysCacheGetAttr(m, int32(47), l0, int32(23), v12+int32(15))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L7:
	;
	v27 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v59 = v14 + v15
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+128))
	v62 = v60 << (uint(int32(2)) % 32)
	v63 = F_palloc(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v29 != int32(1) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v32 < int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v35 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v36 != int32(26) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v40 = v32 << (uint(int32(2)) % 32)
	v41 = F_palloc(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v44 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v54 = (v47<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L17:
	;
	v54 = v44
	goto L18
L18:
	;
	if v40 == int32(0) {
		v72 = v32
		goto L6
	} else {
		goto L19
	}
L19:
	;
	base.MemoryCopy(m, v41, v54+v27, v40)
	v72 = v32
	goto L6
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v63
	if v62 == int32(0) {
		v72 = v60
		goto L6
	} else {
		goto L21
	}
L21:
	;
	base.MemoryCopy(m, v63, v59+int32(136), v62)
	v72 = v60
	goto L6
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v82 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v142 = F_SysCacheGetAttr(m, int32(47), l0, int32(22), v12+int32(15))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v88 = F_pg_detoast_datum(m, v80)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_deconstruct_array_builtin(m, v88, int32(25), v12+int32(8), int32(0), v12+int32(4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v98 != v72 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v102 = F_palloc(m, v72<<(uint(int32(2))%32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v102
	if v72 <= int32(0) {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(0)
	goto L32
L32:
	;
	v117 = v108 << (uint(int32(2)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117+v118)))
	v121 = F_text_to_cstring(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L23
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v117))) = v121
	v127 = v108 + int32(1)
	if v127 != v72 {
		v108 = v127
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v144 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	m.G0 = v12 + int32(16)
	return v72
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v149 = F_pg_detoast_datum(m, v142)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v151 != int32(1) {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	if v154 != v72 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	if v156 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	if v157 != int32(18) {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	v160 = F_palloc(m, v72)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v160
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
	if v163 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v173 = (v166<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L49
L48:
	;
	v173 = v163
	goto L49
L49:
	;
	if v72 == int32(0) {
		goto L37
	} else {
		goto L50
	}
L50:
	;
	base.MemoryCopy(m, v160, v149+v173, v72)
	goto L37
L51:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_0), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1411), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_3), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1438), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v72
	F_errmsg_internal(m, int32(_a_F_get_func_arg_info_4), v12)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_get_func_arg_info_1), int32(1458), int32(_a_F_get_func_arg_info_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_func_namespace(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13900(m, l0, int32(47))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_get_func_support(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13898(m, l0, int32(47))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
