package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_flatten_group_exprs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+39)))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(v13)
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)) = uint8(v13)
	v16 = F_flatten_group_exprs_mutator(m, l2, v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return v16
	}
}
func F_show_sort_group_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
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
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v11 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(96)
	m.G0 = v24
	if l2 <= v11 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L73
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L70
	}
L4:
	;
	m.G0 = v24 + int32(96)
	return
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_initStringInfo(m, v24+int32(76))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l9)+44))
	v35 = F_set_deparse_context_plan(m, v34, v29, l8)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l9)+56))
	if v37 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l9)+4)))
	v41 = v40
	goto L11
L10:
	;
	v41 = int32(1)
	goto L11
L11:
	;
	v56 = int32(0)
	v58 = v11
	v60 = v11
	goto L12
L12:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4+v56<<(uint(int32(1))%32)))))
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	F_ExplainPropertyList(m, l1, v219, l9)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L67
	}
L14:
	;
	if v108 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L15:
	;
	goto L14
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v74 <= int32(0) {
		v108 = int32(0)
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v108 = int32(0)
	goto L15
L19:
	;
	v77 = int32(0)
	if v77 < v74 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = v74
	goto L22
L21:
	;
	v80 = v77
	goto L22
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v85 = int32(0)
	goto L23
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81+v85<<(uint(int32(2))%32))))
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+8)))
	if v94 == v70&int32(65535) {
		v108 = v93
		goto L15
	} else {
		goto L25
	}
L24:
	;
	goto L18
L25:
	;
	v97 = v85 + int32(1)
	if v97 != v80 {
		v85 = v97
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v114 = F_deparse_expression(m, v112, v35, v41&int32(1), int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v117 = v24 + int32(76)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v119
	goto L29
L29:
	;
	F_appendStringInfoString(m, v24+int32(76), v114)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	if l5 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	v217 = F_pstrdup(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L60
	}
L32:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7+v56))))
	v134 = v56 << (uint(int32(2)) % 32)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l6+v134)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l5+v134)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v140 = F_exprType(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+95)) = uint8(v142)
	v145 = F_lookup_type_cache(m, v140, int32(6))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	if v136 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v145)+60))
	if v167 == v138 {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v149 = F_get_typcollation(m, v140)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v149 == v136 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v152 = F_get_collation_name(m, v136)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if v152 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v156 = F_quote_identifier(m, v152)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v156
	F_appendStringInfo(m, v24+int32(76), int32(197832), v24-int32(-64))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+95)))
	if v132&int32(1) != 0 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	F_appendStringInfoString(m, v24+int32(76), int32(542294))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v145)+56))
	if v138 == v176 {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+95)) = uint8(v174)
	goto L43
L48:
	;
	v178 = F_get_opname(m, v138)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	if v178 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v178
	F_appendStringInfo(m, v24+int32(76), int32(197718), v24+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	v192 = F_get_equality_op_for_ordering_op(m, v138, v24+int32(95))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	goto L43
L53:
	;
	F_appendStringInfoString(m, v24+int32(76), v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L59
	}
L54:
	;
	if v195&int32(1) != 0 {
		goto L31
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v195&int32(1) == int32(0) {
		goto L31
	} else {
		goto L58
	}
L57:
	;
	v208 = int32(515169)
	goto L53
L58:
	;
	v208 = int32(515306)
	goto L53
L59:
	;
	goto L31
L60:
	;
	v219 = F_lappend(m, v60, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v56 < l3 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v222 = F_lappend(m, v58, v114)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	v224 = v58
	goto L64
L64:
	;
	v226 = v56 + int32(1)
	if v226 != l2 {
		v56 = v226
		v58 = v224
		v60 = v219
		goto L12
	} else {
		goto L66
	}
L65:
	;
	v224 = v222
	goto L64
L66:
	;
	goto L13
L67:
	;
	if l3 <= int32(0) {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_ExplainPropertyList(m, int32(22768), v224, l9)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L4
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v70
	F_errmsg_internal(m, int32(465878), v24)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(494866), int32(2801), int32(112676))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
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
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v136
	F_errmsg_internal(m, int32(45878), v24+int32(48))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(494866), int32(2852), int32(136888))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v138
	F_errmsg_internal(m, int32(43030), v24+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(494866), int32(2867), int32(136888))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformGroupClauseExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	if l7 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v155
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v155 = v148
	goto L1
L3:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v130 = F_exprLocation(m, l3)
	mBase = m.M
	v131 = F_addTargetToGroupList(m, l2, v20, v125, v129, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L40
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v21 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v14 = F_findTargetlistEntrySQL99(m, l2, l3, l4, l6)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v18 = F_findTargetlistEntrySQL92(m, l2, l3, l4, l6)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	v20 = v14
	goto L4
L10:
	;
	v20 = v18
	goto L4
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = v24
	goto L3
L12:
	;
	goto L13
L13:
	;
	v26 = F_bms_is_member(m, v21, l1)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v26 != 0 {
		v155 = int32(0)
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l5 == int32(0) {
		v125 = v28
		goto L3
	} else {
		goto L24
	}
L17:
	;
	if v28 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v34 <= int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v46 = int32(0)
	goto L20
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v37+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v56 == v29 {
		v155 = v29
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L16
L22:
	;
	v59 = v46 + int32(1)
	if v59 != v34 {
		v46 = v59
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v76 <= int32(0) {
		v125 = v28
		goto L3
	} else {
		goto L25
	}
L25:
	;
	v79 = int32(0)
	if v79 < v76 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v83 = v76
	goto L28
L27:
	;
	v83 = v79
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v92 = v79
	goto L29
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v84+v92<<(uint(int32(2))%32))))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	if v29 != v102 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v107 = F_copyObjectImpl(m, v101)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L35
	}
L31:
	;
	v105 = v92 + int32(1)
	if v83 != v105 {
		v92 = v105
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	v125 = v28
	goto L3
L35:
	;
	if l8 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+17)) = uint8(v111)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = F_lappend(m, v113, v107)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v146 = v114
	goto L2
L40:
	;
	v146 = v131
	goto L2
}
