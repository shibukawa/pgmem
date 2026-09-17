package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildIndexValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v189 int32
	_ = v189
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v22 = F_check_enable_rls(m, v19, v4, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return v189
L2:
	;
	return int32(0)
L3:
	;
	if v22 == int32(2) {
		v189 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndexValueDescription[0]))
	v31 = F_pg_class_aclcheck(m, v19, v29, int64(2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v33 = int32(0)
	if base.B2i32(v31 == v33)|base.B2i32(v18 <= v33) == v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v47 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v80 = v14 + int32(16)
	F_initStringInfo(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L15
	}
L9:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17+int32(48)+v47<<(uint(int32(1))%32)))))
	if v57 == int32(0) {
		v189 = v4
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndexValueDescription[0]))
	v63 = F_pg_attribute_aclcheck(m, v19, v57, v61, int64(2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v63 != 0 {
		v189 = v4
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v66 = v47 + int32(1)
	if v66 != v18 {
		v47 = v66
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v83 = int32(0)
	v85 = int32(1)
	v91 = F_pg_get_indexdef_worker(m, v16, v83, v83, v85, v85, v83, v83, int32(7), v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v91
	F_appendStringInfo(m, v80, int32(_a_F_BuildIndexValueDescription_0), v14)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	if v18 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_appendStringInfoChar(m, v14+int32(16), int32(41))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L37
	}
L19:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v99 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v113 = int32(_a_F_BuildIndexValueDescription_1)
	goto L22
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	F_getTypeOutputInfo(m, v102, v14+int32(12), v14+int32(11))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L23
	}
L22:
	;
	F_appendStringInfoString(m, v80, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v111 = F_OidOutputFunctionCall(m, v109, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v113 = v111
	goto L22
L25:
	;
	v116 = int32(1)
	if v18 == v116 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v123 = v116
	goto L27
L27:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v123))))
	if v132 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L18
L29:
	;
	v136 = v123 << (uint(int32(2)) % 32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136+v137)))
	F_getTypeOutputInfo(m, v139, v14+int32(12), v14+int32(11))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L32
	}
L30:
	;
	v151 = int32(_a_F_BuildIndexValueDescription_1)
	goto L31
L31:
	;
	v153 = v14 + int32(16)
	F_appendStringInfoString(m, v153, int32(_a_F_BuildIndexValueDescription_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1+v136)))
	v149 = F_OidOutputFunctionCall(m, v146, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v151 = v149
	goto L31
L34:
	;
	F_appendStringInfoString(m, v153, v151)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v160 = v123 + int32(1)
	if v160 != v18 {
		v123 = v160
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v189 = v178
	goto L1
}
func F_CompareIndexInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+116)))
	if v22 != v23 {
		v177 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L29
	} else {
		goto L43
	}
L2:
	;
	m.G0 = v20 + int32(16)
	return v177
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v25 != v26 {
		v177 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v28 != v29 {
		v177 = v8
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 != v32 {
		v177 = v8
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != v35 {
		v177 = v8
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v31 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = int32(12)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v54 = v8
	goto L11
L9:
	;
	goto L10
L10:
	;
	v122 = int32(0)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if base.B2i32(v123 == v122) == base.B2i32(v126 != v122) {
		v177 = v122
		goto L2
	} else {
		goto L25
	}
L11:
	;
	v62 = v54 << (uint(int32(1)) % 32)
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1+v39+v62))))
	if v43 < v64 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+(l0+v39)))))
	if (v67|v64)&int32(_a_F_CompareIndexInfo_0) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = int32(0)
	if base.B2i32(v64 == v71)|base.B2i32(v67 == v71) != 0 {
		v177 = v71
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v54 < v34 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77+v64<<(uint(int32(1))%32)-int32(2)))))
	if v83 != v67 {
		v177 = v71
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v87 = int32(0)
	v89 = v54 << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l2+v89)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3+v89)))
	if v91 != v93 {
		v177 = v87
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v103 = v54 + int32(1)
	if v103 != v31 {
		v54 = v103
		goto L11
	} else {
		goto L24
	}
L22:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l4+v89)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l5+v89)))
	if v96 != v98 {
		v177 = v87
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L12
L25:
	;
	if v126 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v134 = F_map_variable_attnos(m, v123, int32(1), l6, int32(0), v20+int32(15))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v146 = int32(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if base.B2i32(v145 == v146) == base.B2i32(v148 != v146) {
		v177 = v122
		goto L2
	} else {
		goto L34
	}
L29:
	;
	return int32(0)
L30:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+15)))
	if v138 != 0 {
		v177 = v122
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v140 = F_equal(m, v139, v134)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if v140 == int32(0) {
		v177 = v122
		goto L2
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	if v145 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v156 = F_map_variable_attnos(m, v148, int32(1), l6, int32(0), v20+int32(14))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L29
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v165 != 0 {
		v177 = v122
		goto L2
	} else {
		goto L42
	}
L38:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v158 != 0 {
		v177 = v122
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v160 = F_equal(m, v159, v156)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L40
	}
L40:
	;
	if v160 == int32(0) {
		v177 = v122
		goto L2
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v177 = base.B2i32(v166 == int32(0))
	goto L2
L43:
	;
	F_errmsg_internal(m, int32(_a_F_CompareIndexInfo_1), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L29
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_CompareIndexInfo_2), int32(2571), int32(_a_F_CompareIndexInfo_3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyIndexTuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v6 = v4 & int32(_a_F_CopyIndexTuple_0)
	v7 = F_palloc(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v6 != 0 {
			base.MemoryCopy(m, v7, l0, v6)
		} else {
		}
		return v7
	}
}
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
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
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v493 int32
	_ = v493
	var v503 int32
	_ = v503
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v633 int32
	_ = v633
	v9 = int32(0)
	v31 = m.G0
	v33 = v31 - int32(176)
	m.G0 = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+152))
	if v39 == v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v42 = F_MakePerTupleExprContext(m, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = v39
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = l1
	if int32(0) < v38 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v46 = v42
	goto L3
L6:
	;
	v51 = l1 + int32(28)
	v74 = v9
	v79 = v9
	goto L9
L7:
	;
	v633 = v9
	goto L8
L8:
	;
	m.G0 = v33 + int32(176)
	return v633
L9:
	;
	v83 = v74 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v37+v83)))
	if v85 == int32(0) {
		v600 = v79
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v633 = v600
	goto L8
L11:
	;
	v604 = v74 + int32(1)
	if v604 != v38 {
		v74 = v604
		v79 = v600
		goto L9
	} else {
		goto L124
	}
L12:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+v36)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+118)))
	if v90 != int32(1) {
		v600 = v79
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l7 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+123)))
	if v93 != int32(1) {
		v600 = v79
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+84))
	if v96 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_FormIndexDatum(m, v89, l1, l2, v33+int32(32), v33)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L27
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	if v99 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v102 = F_ExecPrepareQual(m, v96, l2)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v107 = v99
	goto L22
L22:
	;
	v108 = int32(_a_F_ExecInsertIndexTuples_0)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0])) = v111
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v116 = m.T0[v115].(func(*base.Module, int32, int32, int32) int32)(m, v107, v46, v33+int32(175))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+88)) = v102
	if v102 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v107 = v102
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0])) = v109
	if v116 == int32(0) {
		v600 = v79
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if l4 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v203 = int32(0)
	if l3 == v203 {
		v503 = v203
		goto L58
	} else {
		goto L59
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	if l6 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v180 = int32(0)
	goto L31
L31:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+12)))
	if v184 != 0 {
		goto L51
	} else {
		goto L52
	}
L32:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+12)))
	if v136 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v140 = int32(0)
	if l6 == v140 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v137 = int32(2)
	goto L37
L36:
	;
	v137 = int32(0)
	goto L37
L37:
	;
	v199 = int32(1)
	v200 = v136
	v201 = v137
	v202 = v85 + int32(192)
	goto L28
L38:
	;
	v180 = v178
	goto L31
L39:
	;
	v178 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v146 <= int32(0) {
		v172 = v140
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v178 = v172
	goto L38
L43:
	;
	v149 = int32(0)
	if v149 < v146 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v152 = v146
	goto L46
L45:
	;
	v152 = v149
	goto L46
L46:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v155 = int32(0)
	goto L47
L47:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153+v155<<(uint(int32(2))%32))))
	v164 = base.B2i32(v163 == v139)
	if v163 == v139 {
		v172 = v164
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v172 = v164
	goto L42
L49:
	;
	v166 = v155 + int32(1)
	if v166 != v152 {
		v155 = v166
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v185 = int32(2)
	goto L53
L52:
	;
	v185 = int32(0)
	goto L53
L53:
	;
	v187 = v85 + int32(192)
	if base.B2i32(v184 == int32(0))|v180 != 0 {
		v199 = v180
		v200 = v184
		v201 = v185
		v202 = v187
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+16)))
	if v193 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v194 = int32(1)
	goto L57
L56:
	;
	v194 = int32(2)
	goto L57
L57:
	;
	v199 = int32(0)
	v200 = v193 ^ int32(1)
	v201 = v194
	v202 = v187
	goto L28
L58:
	;
	v526 = v33 + int32(32)
	v529 = F_index_insert(m, v85, v526, v33, v51, v35, v201, v503&int32(1), v89)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L106
	}
L59:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+119)))
	if v206 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+120)))
	v503 = v209
	goto L58
L61:
	;
	goto L62
L62:
	;
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+119)) = uint8(v210)
	v212 = F_ExecGetUpdatedCols(m, l0, l2)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v214 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_ExecInitGenerated(m, l0, l2, int32(2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v221 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v89)+120)) = uint8(v493)
	v503 = v493
	goto L58
L69:
	;
	v493 = int32(0)
	goto L68
L70:
	;
	v493 = int32(1)
	goto L68
L71:
	;
	v225 = v89 + int32(12)
	v235 = v203
	v251 = v221
	v253 = int32(0)
	goto L73
L72:
	;
	if v220 != 0 {
		goto L88
	} else {
		goto L89
	}
L73:
	;
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225+v235<<(uint(int32(1))%32)))))
	if v260 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v329 == int32(0) {
		goto L70
	} else {
		goto L87
	}
L75:
	;
	v271 = v235
	goto L78
L76:
	;
	v311 = v235
	v315 = v260
	v329 = v253
	goto L77
L77:
	;
	v336 = (v315 + int32(7)) & int32(_a_F_ExecInsertIndexTuples_1)
	v337 = F_bms_is_member(m, v336, v212)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L82
	}
L78:
	;
	v294 = v271 + int32(1)
	if v251 <= v294 {
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v311 = v294
	v315 = v300
	v329 = v296
	goto L77
L80:
	;
	v296 = int32(1)
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225+v294<<(uint(v296)%32)))))
	if v300 <= int32(0) {
		v271 = v294
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if v337 != 0 {
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v339 = F_bms_is_member(m, v336, v220)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	if v339 != 0 {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v342 = v311 + int32(1)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v342 < v343 {
		v235 = v342
		v251 = v343
		v253 = v329
		goto L73
	} else {
		goto L86
	}
L86:
	;
	goto L74
L87:
	;
	goto L72
L88:
	;
	v377 = F_bms_union(m, v212, v220)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v379 = v212
	goto L90
L90:
	;
	v381 = F_RelationGetIndexExpressions(m, v85)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	v379 = v377
	goto L90
L92:
	;
	F_list_free(m, v381)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L100
	}
L93:
	;
	if v381 == int32(0) {
		v396 = int32(0)
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v385 == int32(6) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v381)+8)))
	v391 = F_bms_is_member(m, v388+int32(7), v379)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v394 = F_expression_tree_walker_impl(m, v381, int32(625), v379)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	v396 = v391
	goto L92
L99:
	;
	v396 = v394
	goto L92
L100:
	;
	if v220 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_bms_free(m, v379)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L4
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v396 != 0 {
		goto L69
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	goto L70
L106:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v89)+92))
	if v531 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v199 != 0 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v548 = v529
	goto L109
L109:
	;
	if v200 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	v546 = F_check_exclusion_or_unique_constraint(m, v35, v85, v89, v51, v526, v33, l2, int32(0), v542, v541&int32(1), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L114
	}
L111:
	;
	v541 = int32(1)
	v542 = int32(2)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535)+16)))
	v537 = int32(1)
	v538 = v536 ^ v537
	v541 = v538
	v542 = v538 & v537
	goto L110
L114:
	;
	v548 = v546
	goto L109
L115:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v85)+56))
	v561 = F_lappend_oid(m, v79, v560)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L4
	} else {
		goto L121
	}
L116:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v89)+92))
	v553 = int32(0)
	if (base.B2i32(v552 == v553)|v548)&int32(1) == v553 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v548 != 0 {
		v600 = v79
		goto L11
	} else {
		goto L120
	}
L119:
	;
	v600 = v79
	goto L11
L120:
	;
	goto L115
L121:
	;
	if l5 == int32(0) {
		v600 = v561
		goto L11
	} else {
		goto L122
	}
L122:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v85)+192))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+16)))
	if v566&int32(1) == int32(0) {
		v600 = v561
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v571)
	v600 = v561
	goto L11
L124:
	;
	goto L10
}
func F_IndexAmTranslateStrategy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if base.B2i32(l1 != int32(403))|base.B2i32(base.Ui32(int32(5)) <= base.Ui32((l0-int32(1))&int32(_a_F_IndexAmTranslateStrategy_0))) != 0 {
		v19 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
			if v23 != 0 {
				v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l2)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = v24
					v28 = v27
					m.G0 = v7 + int32(16)
					return v28
				}
			} else {
				v27 = int32(0)
				v28 = v27
				m.G0 = v7 + int32(16)
				return v28
			}
		}
	} else {
		v28 = l0
		m.G0 = v7 + int32(16)
		return v28
	}
}
func F_IndexNext(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v17 = v14 * v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v20 != 0 {
		v43 = v20
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v44 = F_index_getnext_slot(m, v43, v17, v18)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L11
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v28 = F_index_beginscan(m, v21, v22, v23, l0+int32(168), v26, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v28
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v33 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v34 != int32(1) {
		v43 = v28
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_index_rescan(m, v28, v37, v38, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	v43 = v28
	goto L1
L10:
	;
	m.G0 = v11 + int32(16)
	return v18
L11:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	goto L15
L13:
	;
	goto L14
L14:
	;
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+184)) = uint8(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	m.T0[v101].(func(*base.Module, int32))(m, v18)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L34
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNext[0]))
	if v55 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+72)))
	if v58 != int32(1) {
		goto L10
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v18
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L3
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v68 = int32(_a_F_IndexNext_0)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1])) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+20))
	v76 = m.T0[v75].(func(*base.Module, int32, int32, int32) int32)(m, v62, v19, v11+int32(15))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L3
	} else {
		goto L26
	}
L25:
	;
	goto L10
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1])) = v69
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_MemoryContextReset(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if v76 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v83)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v83)+248)) = base.F64_add(v84, float64(1))
	goto L31
L30:
	;
	goto L31
L31:
	;
	v88 = F_index_getnext_slot(m, v43, v17, v18)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if v88 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	goto L10
}
func F_build_index_value_desc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	if l2 == int32(0) {
		v48 = int32(0)
		m.G0 = v8 + int32(160)
		return v48
	} else {
		v14 = F_index_open(m, l3, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			if v18 != int32(_a_F_build_index_value_desc_0) {
				v29 = l2
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				if v30 != 0 {
					v33 = v30
					*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
					v35 = F_BuildIndexInfo(m, v14)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v38 = v8 + int32(32)
						F_FormIndexDatum(m, v35, v29, l0, v38, v8)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = F_BuildIndexValueDescription(m, v14, v38, v8)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								F_relation_close(m, v14, int32(0))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v48 = v41
									m.G0 = v8 + int32(160)
									return v48
								}
							}
						}
					}
				} else {
					v31 = F_MakePerTupleExprContext(m, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = v31
						*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
						v35 = F_BuildIndexInfo(m, v14)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v38 = v8 + int32(32)
							F_FormIndexDatum(m, v35, v29, l0, v38, v8)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = F_BuildIndexValueDescription(m, v14, v38, v8)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_relation_close(m, v14, int32(0))
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										v48 = v41
										m.G0 = v8 + int32(160)
										return v48
									}
								}
							}
						}
					}
				}
			} else {
				v23 = F_table_slot_create(m, l1, l0+int32(104))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
					m.T0[v26].(func(*base.Module, int32, int32))(m, v23, l2)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v23
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
						if v30 != 0 {
							v33 = v30
							*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
							v35 = F_BuildIndexInfo(m, v14)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v38 = v8 + int32(32)
								F_FormIndexDatum(m, v35, v29, l0, v38, v8)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = F_BuildIndexValueDescription(m, v14, v38, v8)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										F_relation_close(m, v14, int32(0))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											v48 = v41
											m.G0 = v8 + int32(160)
											return v48
										}
									}
								}
							}
						} else {
							v31 = F_MakePerTupleExprContext(m, l0)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								v33 = v31
								*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v29
								v35 = F_BuildIndexInfo(m, v14)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v38 = v8 + int32(32)
									F_FormIndexDatum(m, v35, v29, l0, v38, v8)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v41 = F_BuildIndexValueDescription(m, v14, v38, v8)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v14, int32(0))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v48 = v41
												m.G0 = v8 + int32(160)
												return v48
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
	}
}
func F_check_index_is_clusterable(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = F_index_open(m, l1, l2)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+192))
		if v11 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v47 = int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v46 + v47
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v45 + v47
					F_errmsg(m, int32(_a_F_check_index_is_clusterable_0), v7)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(509), int32(_a_F_check_index_is_clusterable_2))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v14 != v15 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v47 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v46 + v47
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v45 + v47
						F_errmsg(m, int32(_a_F_check_index_is_clusterable_0), v7)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(509), int32(_a_F_check_index_is_clusterable_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				if v18 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v68 + int32(4)
							F_errmsg(m, int32(_a_F_check_index_is_clusterable_3), v5+int32(-16))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(516), int32(_a_F_check_index_is_clusterable_2))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+196))
					v24 = F_heap_attisnull(m, v21, int32(21), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						if v24 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v89 + int32(4)
									F_errmsg(m, int32(_a_F_check_index_is_clusterable_4), v5+int32(-32))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(528), int32(_a_F_check_index_is_clusterable_2))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+192))
							v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
							if v29 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v110 + int32(4)
										F_errmsg(m, int32(_a_F_check_index_is_clusterable_5), v5+int32(-48))
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(542), int32(_a_F_check_index_is_clusterable_2))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_relation_close(m, v9, int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									m.G0 = v7 - int32(-64)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_create_index_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 float64, l10 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	v15 = F_palloc0(m, int32(112))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(280)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v21
		if l7 != 0 {
			v25 = int32(342)
		} else {
			v25 = int32(341)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v27
		v29 = F_get_baserel_parampathinfo(m, l0, v21, l8)
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			v31 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)) = uint8(v31)
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v29
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+26)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = l6
			*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = l4
			*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = l5
			*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+21)) = uint8(v34)
			F_cost_index(m, v15, l0, l9, l10)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_get_index_am_oid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_get_am_type_oid(m, l0, int32(105), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_get_index_isreplident(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v4 = F_SearchSysCache1(m, int32(34), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9)+22)))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = v11
				return v14 & int32(1)
			}
		} else {
			v14 = int32(0)
			return v14 & int32(1)
		}
	}
}
func F_index_beginscan_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v6 = l5
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_index_beginscan_internal[0]))
	if v15 != v13 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_beginscan_internal[1]))
		v19 = F_list_member_ptr(m, v18, v13)
		mBase = m.M
		v21 = v19
	} else {
		v21 = int32(1)
	}
	if v21 == int32(0) {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+92))
		if v25 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_index_beginscan_internal_0)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v70 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_beginscan_internal_1), v11+int32(16))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_beginscan_internal_2), int32(321), int32(_a_F_index_beginscan_internal_3))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)))
			if v28 == int32(0) {
				F_PredicateLockRelation(m, l0, l3)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_RelationIncrementReferenceCount(m, l0)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+92))
						v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v39)+29)) = uint8(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = l4
							m.G0 = v11 + int32(32)
							return v39
						}
					}
				}
			} else {
				F_RelationIncrementReferenceCount(m, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+92))
					v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v39)+29)) = uint8(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = l4
						m.G0 = v11 + int32(32)
						return v39
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54 + int32(4)
				F_errmsg(m, int32(_a_F_index_beginscan_internal_4), v11)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_beginscan_internal_2), int32(320), int32(_a_F_index_beginscan_internal_3))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
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
func F_index_expression_changed_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 == int32(6) {
			v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v13 = F_bms_is_member(m, v10+int32(7), l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v13
			}
		} else {
			v19 = F_expression_tree_walker_impl(m, l0, int32(625), l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v19
			}
		}
	}
}
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v13 {
		v16 = int32(4)
		v20 = l2 + l1<<(uint(v16)%32) + v16
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		if v21 < int32(0) {
			v69 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v75 = v69
				m.G0 = v9 + int32(16)
				return v75
			}
		} else {
			v26 = l0 + v21 + int32(8)
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+6)))
			if v27 != int32(1) {
				v75 = v26
				m.G0 = v9 + int32(16)
				return v75
			} else {
				v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
				switch v30&int32(_a_F_index_getattr_2_0) - int32(1) {
				case 0:
					v35 = int32(*(*int8)(unsafe.Add(mBase, uint32(v26))))
					v75 = v35
					m.G0 = v9 + int32(16)
					return v75
				case 1:
					v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26))))
					v75 = v36
					m.G0 = v9 + int32(16)
					return v75
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v30
						F_errmsg_internal(m, int32(_a_F_index_getattr_2_1), v9)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_index_getattr_2_2), int32(70), int32(_a_F_index_getattr_2_3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					v75 = v37
					m.G0 = v9 + int32(16)
					return v75
				}
			}
		}
	} else {
		v53 = int32(1)
		v54 = l1 - v53
		v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54>>(uint(int32(3))%32))+8)))
		if int32(base.Ui32(v58)>>(uint(v54&int32(7))%32))&v53 != 0 {
			v69 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v75 = v69
				m.G0 = v9 + int32(16)
				return v75
			}
		} else {
			v64 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v64)
			v75 = int32(0)
			m.G0 = v9 + int32(16)
			return v75
		}
	}
}
func F_index_rescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	if v15 != 0 {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v16 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
			m.T0[v19].(func(*base.Module, int32))(m, v16)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v23 = v22
				v24 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v24)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v24)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
				m.T0[v29].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			}
		} else {
			v23 = v13
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v24)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v24)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
			m.T0[v29].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, l1, l2, l3, l4)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_index_rescan_0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v40 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_rescan_1), v11)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_rescan_2), int32(361), int32(_a_F_index_rescan_3))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
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
func F_index_set_state_flags(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_SearchSysCacheCopy(m, int32(34), l0, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v20 = v18 + v19
				switch l1 {
				case 0:
					v21 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v21)
				case 1:
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+18)) = uint8(v23)
				case 2:
					v25 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)) = uint8(v25)
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+17)) = uint16(v25)
				case 3:
					v29 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v20)+20)) = uint16(v29)
				default:
				}
				F_CatalogTupleUpdate(m, v12, v16+int32(4), v16)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_relation_close(m, v12, int32(3))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_index_set_state_flags_0), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_index_set_state_flags_1), int32(3515), int32(_a_F_index_set_state_flags_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
