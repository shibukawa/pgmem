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
	var v110 int64
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
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
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
	return v191
L2:
	;
	return int32(0)
L3:
	;
	if v22 == int32(2) {
		v191 = v4
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
		v191 = v4
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
		v191 = v4
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
	v179 = m.ExcPending
	if v179 != 0 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
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
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v123<<(uint(int32(2))%32))))
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
	v153 = int32(_a_F_BuildIndexValueDescription_1)
	goto L31
L31:
	;
	v155 = v14 + int32(16)
	F_appendStringInfoString(m, v155, int32(_a_F_BuildIndexValueDescription_2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(l1+v123<<(uint(int32(3))%32))))
	v151 = F_OidOutputFunctionCall(m, v146, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v153 = v151
	goto L31
L34:
	;
	F_appendStringInfoString(m, v155, v153)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v162 = v123 + int32(1)
	if v162 != v18 {
		v123 = v162
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v191 = v180
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+116)))
	if v22 != v23 {
		v236 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L29
	} else {
		goto L51
	}
L2:
	;
	m.G0 = v20 + int32(16)
	return v236
L3:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+117)))
	if v25 != v26 {
		v236 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v28 != v29 {
		v236 = v8
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v31 != v32 {
		v236 = v8
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v34 != v35 {
		v236 = v8
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
		v236 = v122
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
		v236 = v71
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
		v236 = v71
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
		v236 = v87
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
		v236 = v87
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
		v236 = v122
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
		v236 = v122
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
		v236 = v122
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
	v166 = int32(0)
	v167 = base.B2i32(v165 == v166)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v170 = base.B2i32(v168 == v166)
	if v167|v170 != 0 {
		v236 = base.B2i32(v167^v170 == v166)
		goto L2
	} else {
		goto L42
	}
L38:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+14)))
	if v158 != 0 {
		v236 = v122
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
		v236 = v122
		goto L2
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v179 = int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v180 <= int32(0) {
		v236 = v179
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v190 = int32(0)
	goto L44
L44:
	;
	v202 = v190 << (uint(int32(2)) % 32)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v165+v202)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v168)))
	if v204 != v206 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v236 = int32(0)
	goto L2
L46:
	;
	goto L45
L47:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208+v202)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211+v202)))
	if v210 != v213 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v216 = v190 << (uint(int32(1)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216+v217))))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220+v216))))
	if v219 != v222 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v225 = v190 + int32(1)
	if v180 != v225 {
		v190 = v225
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v236 = v179
	goto L2
L51:
	;
	F_errmsg_internal(m, int32(_a_F_CompareIndexInfo_1), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L29
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_CompareIndexInfo_2), int32(2604), int32(_a_F_CompareIndexInfo_3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L29
	} else {
		goto L53
	}
L53:
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
func F_ExecInsertIndexTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v543 int32
	_ = v543
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
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v648 int32
	_ = v648
	v7 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(304)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v40 == v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = F_MakePerTupleExprContext(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v47 = v40
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l3
	if int32(0) < v39 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v47 = v43
	goto L3
L6:
	;
	v52 = l3 + int32(32)
	v62 = int32(0)
	v84 = v7
	goto L9
L7:
	;
	v648 = v7
	goto L8
L8:
	;
	m.G0 = v34 + int32(304)
	return v648
L9:
	;
	v92 = v62 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v38+v92)))
	if v94 == int32(0) {
		v614 = v84
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v648 = v614
	goto L8
L11:
	;
	v622 = v62 + int32(1)
	if v622 != v39 {
		v62 = v622
		v84 = v614
		goto L9
	} else {
		goto L124
	}
L12:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92+v37)))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+118)))
	if v99 != int32(1) {
		v614 = v84
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l2&int32(4) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+123)))
	if v102 != int32(1) {
		v614 = v84
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+84))
	if v105 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_FormIndexDatum(m, v98, l3, l1, v34+int32(32), v34)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L27
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+88))
	if v108 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v111 = F_ExecPrepareQual(m, v105, l1)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	v116 = v108
	goto L22
L22:
	;
	v117 = int32(_a_F_ExecInsertIndexTuples_0)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0])) = v120
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+24))
	v125 = m.T0[v124].(func(*base.Module, int32, int32, int32) int64)(m, v116, v47, v34+int32(303))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+88)) = v111
	if v111 == int32(0) {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v116 = v111
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecInsertIndexTuples[0])) = v118
	if v125 == int64(0) {
		v614 = v84
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if l2&int32(2) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v212 = int32(0)
	if l2&int32(1) == v212 {
		v517 = v212
		goto L58
	} else {
		goto L59
	}
L29:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v94)+192))
	if l4 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v189 = int32(0)
	goto L31
L31:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v94)+192))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+12)))
	if v193 != 0 {
		goto L51
	} else {
		goto L52
	}
L32:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+12)))
	if v145 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v149 = int32(0)
	if l4 == v149 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v146 = int32(2)
	goto L37
L36:
	;
	v146 = int32(0)
	goto L37
L37:
	;
	v208 = int32(1)
	v209 = v145
	v210 = v146
	v211 = v94 + int32(192)
	goto L28
L38:
	;
	v189 = v187
	goto L31
L39:
	;
	v187 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v155 <= int32(0) {
		v181 = v149
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v187 = v181
	goto L38
L43:
	;
	v158 = int32(0)
	if v158 < v155 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v161 = v155
	goto L46
L45:
	;
	v161 = v158
	goto L46
L46:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v164 = int32(0)
	goto L47
L47:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v162+v164<<(uint(int32(2))%32))))
	v173 = base.B2i32(v172 == v148)
	if v172 == v148 {
		v181 = v173
		goto L42
	} else {
		goto L49
	}
L48:
	;
	v181 = v173
	goto L42
L49:
	;
	v175 = v164 + int32(1)
	if v175 != v161 {
		v164 = v175
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v194 = int32(2)
	goto L53
L52:
	;
	v194 = int32(0)
	goto L53
L53:
	;
	v196 = v94 + int32(192)
	if base.B2i32(v193 == int32(0))|v189 != 0 {
		v208 = v189
		v209 = v193
		v210 = v194
		v211 = v196
		goto L28
	} else {
		goto L54
	}
L54:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+16)))
	if v202 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v203 = int32(1)
	goto L57
L56:
	;
	v203 = int32(2)
	goto L57
L57:
	;
	v208 = int32(0)
	v209 = v202 ^ int32(1)
	v210 = v203
	v211 = v196
	goto L28
L58:
	;
	v543 = v34 + int32(32)
	v546 = F_index_insert(m, v94, v543, v34, v52, v36, v210, v517&int32(1), v98)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L106
	}
L59:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+119)))
	if v215 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+120)))
	v517 = v218
	goto L58
L61:
	;
	goto L62
L62:
	;
	v219 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+119)) = uint8(v219)
	v221 = F_ExecGetUpdatedCols(m, l0, l1)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v223 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_ExecInitGenerated(m, l0, l1, int32(2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v230 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+120)) = uint8(v509)
	v517 = v509
	goto L58
L69:
	;
	v509 = int32(0)
	goto L68
L70:
	;
	v509 = int32(1)
	goto L68
L71:
	;
	v234 = v98 + int32(12)
	v242 = v212
	v257 = v230
	v259 = int32(0)
	goto L73
L72:
	;
	if v229 != 0 {
		goto L88
	} else {
		goto L89
	}
L73:
	;
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234+v242<<(uint(int32(1))%32)))))
	if v270 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v337 == int32(0) {
		goto L70
	} else {
		goto L87
	}
L75:
	;
	v279 = v242
	goto L78
L76:
	;
	v320 = v242
	v324 = v270
	v337 = v259
	goto L77
L77:
	;
	v348 = (v324 + int32(7)) & int32(_a_F_ExecInsertIndexTuples_1)
	v349 = F_bms_is_member(m, v348, v221)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L82
	}
L78:
	;
	v305 = v279 + int32(1)
	if v257 <= v305 {
		goto L72
	} else {
		goto L80
	}
L79:
	;
	v320 = v305
	v324 = v311
	v337 = v307
	goto L77
L80:
	;
	v307 = int32(1)
	v311 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234+v305<<(uint(v307)%32)))))
	if v311 <= int32(0) {
		v279 = v305
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	if v349 != 0 {
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v351 = F_bms_is_member(m, v348, v229)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	if v351 != 0 {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v354 = v320 + int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	if v354 < v355 {
		v242 = v354
		v257 = v355
		v259 = v337
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
	v390 = F_bms_union(m, v221, v229)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	v392 = v221
	goto L90
L90:
	;
	v394 = F_RelationGetIndexExpressions(m, v94)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	v392 = v390
	goto L90
L92:
	;
	F_list_free(m, v394)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L100
	}
L93:
	;
	if v394 == int32(0) {
		v409 = int32(0)
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	if v398 == int32(6) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v401 = int32(*(*int16)(unsafe.Add(mBase, uint32(v394)+8)))
	v404 = F_bms_is_member(m, v401+int32(7), v392)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v407 = F_expression_tree_walker_impl(m, v394, int32(671), v392)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	v409 = v404
	goto L92
L99:
	;
	v409 = v407
	goto L92
L100:
	;
	if v229 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	F_bms_free(m, v392)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v409 != 0 {
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
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
	if v548 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	if v208 != 0 {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v565 = v546
	goto L109
L109:
	;
	if v209 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	v563 = F_check_exclusion_or_unique_constraint(m, v36, v94, v98, v52, v543, v34, l1, int32(0), v559, v558&int32(1), int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L114
	}
L111:
	;
	v558 = int32(1)
	v559 = int32(2)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+16)))
	v554 = int32(1)
	v555 = v553 ^ v554
	v558 = v555
	v559 = v555 & v554
	goto L110
L114:
	;
	v565 = v563
	goto L109
L115:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v94)+56))
	v578 = F_lappend_oid(m, v84, v577)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L4
	} else {
		goto L121
	}
L116:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v98)+92))
	v570 = int32(0)
	if (base.B2i32(v569 == v570)|v565)&int32(1) == v570 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if v565 != 0 {
		v614 = v84
		goto L11
	} else {
		goto L120
	}
L119:
	;
	v614 = v84
	goto L11
L120:
	;
	goto L115
L121:
	;
	if l5 == int32(0) {
		v614 = v578
		goto L11
	} else {
		goto L122
	}
L122:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v94)+192))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v582)+16)))
	if v583&int32(1) == int32(0) {
		v614 = v578
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v588 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v588)
	v614 = v578
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 float64
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v18 = v15 * v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v21 != 0 {
		v48 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v49 = F_index_getnext_slot(m, v48, v18, v19)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L15
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v30 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(1024)
	goto L7
L6:
	;
	v34 = int32(0)
	goto L7
L7:
	;
	v35 = F_index_beginscan(m, v22, v23, v24, v25, v26, v27, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v38 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+148)))
	if v39 != int32(1) {
		v48 = v35
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	F_index_rescan(m, v35, v42, v43, v44, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v48 = v35
	goto L1
L14:
	;
	m.G0 = v12 + int32(16)
	return v19
L15:
	;
	if v49 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	goto L19
L17:
	;
	goto L18
L18:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+176)) = uint8(v107)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	m.T0[v110].(func(*base.Module, int32))(m, v19)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L38
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNext[0]))
	if v61 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+72)))
	if v64 != int32(1) {
		goto L14
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v19
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v68 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	F_MemoryContextReset(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v74 = int32(_a_F_IndexNext_0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1]))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1])) = v77
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32) int64)(m, v68, v20, v12+int32(15))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L30
	}
L29:
	;
	goto L14
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexNext[1])) = v75
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	F_MemoryContextReset(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v82 != int64(0) {
		goto L14
	} else {
		goto L32
	}
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v91 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v91)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v91)+432)) = base.F64_add(v92, float64(1))
	goto L35
L34:
	;
	goto L35
L35:
	;
	v96 = F_index_getnext_slot(m, v48, v18, v19)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if v96 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	goto L14
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
	v8 = v6 - int32(288)
	m.G0 = v8
	if l2 == int32(0) {
		v48 = int32(0)
		m.G0 = v8 + int32(288)
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
									m.G0 = v8 + int32(288)
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
										m.G0 = v8 + int32(288)
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
											m.G0 = v8 + int32(288)
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
												m.G0 = v8 + int32(288)
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
						F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(782), int32(_a_F_check_index_is_clusterable_2))
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
							F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(782), int32(_a_F_check_index_is_clusterable_2))
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
								F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(789), int32(_a_F_check_index_is_clusterable_2))
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
										F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(801), int32(_a_F_check_index_is_clusterable_2))
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
											F_errfinish(m, int32(_a_F_check_index_is_clusterable_1), int32(815), int32(_a_F_check_index_is_clusterable_2))
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
	var v46 int32
	_ = v46
	v15 = F_palloc0(m, int32(112))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(283)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v21
		if l7 != 0 {
			v25 = int32(346)
		} else {
			v25 = int32(345)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v25
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
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
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+104)))
				if v46 == int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(1)
				} else {
				}
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v5 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+22)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v10)+22)))
			F_ReleaseCatCache(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = v12
				return v15 & int32(1)
			}
		} else {
			v15 = int32(0)
			return v15 & int32(1)
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
					F_errfinish(m, int32(_a_F_index_beginscan_internal_2), int32(333), int32(_a_F_index_beginscan_internal_3))
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
					F_errfinish(m, int32(_a_F_index_beginscan_internal_2), int32(332), int32(_a_F_index_beginscan_internal_3))
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
			v19 = F_expression_tree_walker_impl(m, l0, int32(671), l1)
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
func F_index_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v5 <= v14 {
		v21 = l2 + l1<<(uint(int32(3))%32) + int32(20)
		v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21))))
		if v22 < int32(0) {
			v72 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int64(0)
			} else {
				v78 = v72
				m.G0 = v10 + int32(16)
				return v78
			}
		} else {
			v27 = l0 + v22 + int32(8)
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
			if v28 == int32(1) {
				v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+2)))
				if base.I32_popcnt(v31) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
						F_errmsg_internal(m, int32(_a_F_index_getattr_2_0), v10)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_index_getattr_2_1), int32(123), int32(_a_F_index_getattr_2_2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v31) {
					case 0:
						v36 = int64(*(*int8)(unsafe.Add(mBase, uint32(v27))))
						v78 = v36
						m.G0 = v10 + int32(16)
						return v78
					case 1:
						v37 = int64(*(*int16)(unsafe.Add(mBase, uint32(v27))))
						v78 = v37
						m.G0 = v10 + int32(16)
						return v78
					case 2:
						v38 = int64(*(*int32)(unsafe.Add(mBase, uint32(v27))))
						v78 = v38
						m.G0 = v10 + int32(16)
						return v78
					case 3:
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
						v78 = v39
						m.G0 = v10 + int32(16)
						return v78
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v31
							F_errmsg_internal(m, int32(_a_F_index_getattr_2_0), v10)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_index_getattr_2_1), int32(123), int32(_a_F_index_getattr_2_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v78 = base.I64_extend_i32_u(v27)
				m.G0 = v10 + int32(16)
				return v78
			}
		}
	} else {
		v56 = int32(1)
		v57 = l1 - v56
		v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v57>>(uint(int32(3))%32))+8)))
		if int32(base.Ui32(v61)>>(uint(v57&int32(7))%32))&v56 != 0 {
			v72 = F_nocache_index_getattr(m, l0, l1, l2)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int64(0)
			} else {
				v78 = v72
				m.G0 = v10 + int32(16)
				return v78
			}
		} else {
			v67 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v67)
			v78 = int64(0)
			m.G0 = v10 + int32(16)
			return v78
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
				F_errfinish(m, int32(_a_F_index_rescan_2), int32(373), int32(_a_F_index_rescan_3))
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v17 = F_SearchSysCacheCopy(m, int32(34), base.I64_extend_i32_u(l0), int64(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
				v21 = v19 + v20
				switch l1 {
				case 0:
					v22 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)) = uint8(v22)
				case 1:
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v21)+18)) = uint8(v24)
				case 2:
					v26 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)) = uint8(v26)
					*(*uint16)(unsafe.Add(mBase, uint32(v21)+17)) = uint16(v26)
				case 3:
					v30 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v30)
				default:
				}
				F_CatalogTupleUpdate(m, v12, v17+int32(4), v17)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_relation_close(m, v12, int32(3))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_index_set_state_flags_0), v8)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_index_set_state_flags_1), int32(3637), int32(_a_F_index_set_state_flags_2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
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
