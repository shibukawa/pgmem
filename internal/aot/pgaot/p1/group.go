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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
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
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	v11 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(96)
	m.G0 = v25
	if l2 <= v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v25 + int32(96)
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_initStringInfo(m, v25+int32(76))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l9)+44))
	v35 = F_set_deparse_context_plan(m, v34, v29, l8)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l9)+56))
	if v37 <= int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l9)+4)))
	v42 = v40
	goto L8
L7:
	;
	v42 = int32(1)
	goto L8
L8:
	;
	v54 = int32(0)
	v60 = v11
	v65 = v11
	goto L9
L9:
	;
	v69 = v25 + int32(76)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(l4+v54<<(uint(int32(1))%32)))))
	if v70 != 0 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	F_ExplainPropertyList(m, l1, v272, l9)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L3
	} else {
		goto L76
	}
L11:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v270 = F_pstrdup(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L69
	}
L12:
	;
	F_appendStringInfoString(m, v69, v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L68
	}
L13:
	;
	v249 = int32(1)
	v251 = int32(0)
	if base.B2i32(v248&v249 == v251)|base.B2i32(v247&v249 == v251) != 0 {
		goto L11
	} else {
		goto L67
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L64
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L3
	} else {
		goto L61
	}
L16:
	;
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	goto L16
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v78 <= int32(0) {
		v112 = int32(0)
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v112 = int32(0)
	goto L17
L21:
	;
	v81 = int32(0)
	if v81 < v78 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L24
L23:
	;
	v84 = v81
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v89 = int32(0)
	goto L25
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+v89<<(uint(int32(2))%32))))
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+8)))
	if v98 == v74&int32(_a_F_show_sort_group_keys_0) {
		v112 = v97
		goto L17
	} else {
		goto L27
	}
L26:
	;
	goto L20
L27:
	;
	v101 = v89 + int32(1)
	if v101 != v84 {
		v89 = v101
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v116 = F_deparse_expression(m, v114, v35, v42&int32(1), int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L3
	} else {
		goto L58
	}
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v69)+4)) = v119
	goto L33
L33:
	;
	F_appendStringInfoString(m, v69, v116)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L34
	}
L34:
	;
	if l5 == int32(0) {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7+v54))))
	v132 = v54 << (uint(int32(2)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l6+v132)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l5+v132)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v138 = F_exprType(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+95)) = uint8(v140)
	v143 = F_lookup_type_cache(m, v138, int32(6))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	if v134 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v143)+60))
	if v163 == v136 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v147 = F_get_typcollation(m, v138)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if v147 == v134 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v150 = F_get_collation_name(m, v134)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	if v150 == int32(0) {
		goto L15
	} else {
		goto L43
	}
L43:
	;
	v154 = F_quote_identifier(m, v150)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v154
	F_appendStringInfo(m, v69, int32(_a_F_show_sort_group_keys_1), v25-int32(-64))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	goto L38
L46:
	;
	F_appendStringInfoString(m, v25+int32(76), int32(_a_F_show_sort_group_keys_2))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v143)+56))
	if v176 != v136 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+95)) = uint8(v170)
	v247 = v170
	v248 = v130 ^ v170
	goto L13
L50:
	;
	v178 = F_get_opname(m, v136)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	v195 = int32(0)
	goto L52
L52:
	;
	v196 = int32(1)
	v197 = v130 ^ v196
	if v197&v196|v195&v196 != 0 {
		v247 = v195
		v248 = v197
		goto L13
	} else {
		goto L57
	}
L53:
	;
	if v178 == int32(0) {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v178
	F_appendStringInfo(m, v25+int32(76), int32(_a_F_show_sort_group_keys_3), v25+int32(32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v192 = F_get_equality_op_for_ordering_op(m, v136, v25+int32(95))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+95)))
	v195 = v194
	goto L52
L57:
	;
	v261 = int32(_a_F_show_sort_group_keys_4)
	goto L12
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v74
	F_errmsg_internal(m, int32(_a_F_show_sort_group_keys_5), v25)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_show_sort_group_keys_6), int32(2801), int32(_a_F_show_sort_group_keys_7))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v134
	F_errmsg_internal(m, int32(_a_F_show_sort_group_keys_8), v25+int32(48))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_show_sort_group_keys_6), int32(2852), int32(_a_F_show_sort_group_keys_9))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L3
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
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v136
	F_errmsg_internal(m, int32(_a_F_show_sort_group_keys_10), v25+int32(16))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_show_sort_group_keys_6), int32(2867), int32(_a_F_show_sort_group_keys_9))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v261 = int32(_a_F_show_sort_group_keys_11)
	goto L12
L68:
	;
	goto L11
L69:
	;
	v272 = F_lappend(m, v65, v270)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	if v54 < l3 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v275 = F_lappend(m, v60, v116)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L3
	} else {
		goto L74
	}
L72:
	;
	v277 = v60
	goto L73
L73:
	;
	v279 = v54 + int32(1)
	if v279 != l2 {
		v54 = v279
		v60 = v277
		v65 = v272
		goto L9
	} else {
		goto L75
	}
L74:
	;
	v277 = v275
	goto L73
L75:
	;
	goto L10
L76:
	;
	if l3 <= int32(0) {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_ExplainPropertyList(m, int32(_a_F_show_sort_group_keys_12), v277, l9)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	goto L1
}
func F_transformGroupClauseExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	if l7 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v150
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v150 = v143
	goto L1
L3:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v126 = F_exprLocation(m, l3)
	mBase = m.M
	v127 = F_addTargetToGroupList(m, l2, v19, v114, v125, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L39
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v13 = F_findTargetlistEntrySQL99(m, l2, l3, l4, l6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v17 = F_findTargetlistEntrySQL92(m, l2, l3, l4, l6)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	v19 = v13
	goto L4
L10:
	;
	v19 = v17
	goto L4
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = v23
	goto L3
L12:
	;
	goto L13
L13:
	;
	v25 = F_bms_is_member(m, v20, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if v25 != 0 {
		v150 = int32(0)
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v28 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.B2i32(v27 == v28)|base.B2i32(v30 == v28) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l5 == int32(0) {
		v114 = v30
		goto L3
	} else {
		goto L23
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v34 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v46 = int32(0)
	goto L19
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v37+v46<<(uint(int32(2))%32))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v55 == v27 {
		v150 = v27
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L16
L21:
	;
	v58 = v46 + int32(1)
	if v58 != v34 {
		v46 = v58
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v74 <= int32(0) {
		v114 = v30
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v77 = int32(0)
	if v77 < v74 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = v74
	goto L27
L26:
	;
	v81 = v77
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v90 = v77
	goto L28
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v82+v90<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v27 != v99 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v104 = F_copyObjectImpl(m, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L8
	} else {
		goto L34
	}
L30:
	;
	v102 = v90 + int32(1)
	if v81 != v102 {
		v90 = v102
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
	v114 = v30
	goto L3
L34:
	;
	if l8 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v104)+17)) = uint8(v108)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = F_lappend(m, v110, v104)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v141 = v111
	goto L2
L39:
	;
	v141 = v127
	goto L2
}
