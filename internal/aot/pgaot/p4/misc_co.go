package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_CommentObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
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
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_CommentObject[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(9) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return
L2:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v241 == int32(0) {
		goto L1
	} else {
		goto L58
	}
L3:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_CreateComments(m, v227, v228, v229, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L57
	}
L4:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v120 = int32(0)
	v121 = m.G0
	v123 = v121 - int32(160)
	m.G0 = v123
	v125 = int32(1)
	if v119 == v120 {
		v149 = v125
		v150 = v3
		goto L29
	} else {
		goto L30
	}
L5:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+48))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+119)))
	v81 = v79 - int32(99)
	if int32(1)<<(uint(v81)%32)&int32(_a_F_CommentObject_3) != 0 {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v61 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L15
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v26 = F_get_database_oid(m, v24, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v31 = v20
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommentObject[2])))
	F_get_object_address(m, l0, v31, v32, v12+int32(44), int32(4), v37&base.B2i32(v31 == int32(22)))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	if v26 == int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v31 = v30
	goto L9
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_CommentObject[3]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v49
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_check_object_ownership(m, v44, v46, v12+int32(32), v45, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v56 - int32(6) {
	case 0:
		goto L5
	default:
		goto L3
	case 3, 28, 37:
		goto L4
	}
L15:
	;
	if v61 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v24
	F_errmsg(m, int32(_a_F_CommentObject_0), v12)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(62), int32(_a_F_CommentObject_2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v89 = base.B2i32(base.Ui32(v81) <= base.Ui32(int32(19)))
	goto L22
L21:
	;
	v89 = int32(0)
	goto L22
L22:
	;
	if v89 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v98 + int32(4)
	F_errmsg(m, int32(_a_F_CommentObject_4), v12+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+48))
	v109 = int32(*(*int8)(unsafe.Add(mBase, uint32(v108)+119)))
	F_errdetail_relkind_not_supported(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_CommentObject_1), int32(113), int32(_a_F_CommentObject_2))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v152 = v123 + int32(48)
	F_ScanKeyInit(m, v152, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v117))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L33
	}
L30:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	if v128 == int32(0) {
		v149 = v125
		v150 = v3
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+14)) = uint8(v131)
	*(*uint16)(unsafe.Add(mBase, uint32(v123)+12)) = uint16(v131)
	v136 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v123)+8)) = uint16(v136)
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+10)) = uint8(v138)
	*(*int64)(unsafe.Add(mBase, uint32(v123)+24)) = base.I64_extend_i32_u(v118)
	*(*int64)(unsafe.Add(mBase, uint32(v123)+16)) = base.I64_extend_i32_u(v117)
	v145 = F_cstring_to_text(m, v119)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L32
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v123)+32)) = base.I64_extend_i32_u(v145)
	v149 = v131
	v150 = v138
	goto L29
L33:
	;
	F_ScanKeyInit(m, v123+int32(104), int32(2), int32(3), int32(184), base.I64_extend_i32_u(v118))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v169 = F_table_open(m, int32(2396), int32(3))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	F_systable_endscan(m, v175)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L46
	}
L36:
	;
	v175 = F_systable_beginscan(m, v169, int32(2397), int32(1), int32(0), int32(2), v152)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	v177 = F_systable_getnext(m, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	if v177 == int32(0) {
		v198 = v120
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v149 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_simple_heap_delete(m, v169, v177+int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v169)+52))
	v194 = F_heap_modify_tuple(m, v177, v187, v123+int32(16), v123+int32(12), v123+int32(8))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v198 = v120
	goto L35
L44:
	;
	F_CatalogTupleUpdate(m, v169, v177+int32(4), v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v198 = v194
	goto L35
L46:
	;
	v201 = int32(0)
	if base.B2i32(v150 == v201)|base.B2i32(v198 != v201) == v201 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v169)+52))
	v213 = F_heap_form_tuple(m, v208, v123+int32(16), v123+int32(12))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L50
	}
L48:
	;
	v217 = v198
	goto L49
L49:
	;
	if v217 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_CatalogTupleInsert(m, v169, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	v217 = v213
	goto L49
L52:
	;
	F_pfree(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	F_relation_close(m, v169, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	m.G0 = v123 + int32(160)
	goto L2
L57:
	;
	goto L2
L58:
	;
	F_relation_close(m, v241, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	goto L1
}
func F_CompareFurthestCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_gt(v7, v8)
	}
	return v11
}
func F_CompareLists(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_lt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_gt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompareNearestDiscardedCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v9 = *(*float64)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(l1)+20))
	if base.F64_gt(v9, v10) != 0 {
		v12 = int32(-1)
	} else {
		v12 = int32(0)
	}
	if base.F64_lt(v9, v10) != 0 {
		v14 = int32(1)
	} else {
		v14 = v12
	}
	return v14
}
func F_CompleteCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	v9 = l8
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v62
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v65 != 0 {
		goto L28
	} else {
		goto L29
	}
L2:
	;
	v61 = l1
	v62 = v14
	goto L1
L3:
	;
	goto L4
L4:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v20 != v15 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v56 = F_AllocSetContextCreateInternal(m, v15, int32(_a_F_CompleteCachedPlan_0), int32(0), int32(1024), int32(_a_F_CompleteCachedPlan_1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = l2
	v61 = l1
	v62 = l2
	goto L1
L9:
	;
	if v20 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	if v15 != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v24 == int32(0) {
		goto L12
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v24
	goto L14
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v24
	goto L14
L18:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v30
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v15
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v37
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(0)
	goto L11
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = l2
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = l2
	goto L8
L25:
	;
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v56
	v59 = F_copyObjectImpl(m, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v61 = v59
	v62 = v56
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v15
	if int32(0) < l4 {
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_extract_query_dependencies(m, v61, l0-int32(-64), l0+int32(68), l0+int32(85))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L25
	} else {
		goto L44
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v70 - int32(137) {
	case 0, 1, 2, 3, 4, 6, 7, 64, 76, 104, 105:
		v74 = int32(1)
		goto L35
	default:
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v75 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L34:
	;
	if v74 != 0 {
		goto L30
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	v74 = int32(0)
	goto L35
L37:
	;
	goto L28
L38:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v79 != int32(6) {
		v94 = int32(1)
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v94&int32(1) == int32(0) {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v86 = v84 - int32(201)
	if base.Ui32(int32(41)) < base.Ui32(v86) {
		v94 = int32(0)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v94 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v86)) % 64)))
	goto L40
L43:
	;
	goto L30
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v109
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+84)) = uint8(v112)
	v114 = F_GetSearchPathMatcher(m, v62)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v114
	goto L28
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v9)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l4
	v141 = F_ChoosePortalStrategy(m, v61)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L25
	} else {
		goto L56
	}
L47:
	;
	v123 = F_palloc_mul(m, int32(4), l4)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L25
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L46
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v123
	v127 = l4 << (uint(int32(2)) % 32)
	if v127 == int32(0) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	base.MemoryCopy(m, v123, l3, v127)
	goto L46
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v183
	*(*int32)(unsafe.Add(mBase, _c_F_CompleteCachedPlan[0])) = v14
	v194 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v194)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+93)) = uint8(v194)
	return
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+28))
	v177 = F_UtilityTupleDescriptor(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L62
	}
L54:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v151 = int32(0)
	goto L58
L55:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+76))
	v146 = F_ExecCleanTypeFromTL(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L57
	}
L56:
	;
	switch v141 {
	case 0, 2:
		goto L55
	case 1:
		goto L54
	case 3:
		goto L53
	default:
		v183 = int32(0)
		goto L52
	}
L57:
	;
	v183 = v146
	goto L52
L58:
	;
	v164 = int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v148+v151<<(uint(int32(2))%32))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+24)))
	if v168 != v164 {
		v151 = v151 + v164
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167)+96))
	v172 = F_ExecCleanTypeFromTL(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L25
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v183 = v172
	goto L52
L62:
	;
	v183 = v177
	goto L52
}
func F_ConditionVariableSleep(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F_ConditionVariableTimedSleep(m, l0, int32(-1), l1)
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_CopyLoadInputBuf(m *base.Module, l0 int32) {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v10 == v11 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v8
	goto L3
L2:
	;
	goto L3
L3:
	;
	goto L5
L4:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	F_report_invalid_encoding(m, v295, v296+v252, v259-v252)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L19
	} else {
		goto L89
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return
L7:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v9-v8 < v252-v253 {
		goto L76
	} else {
		goto L77
	}
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v25 == v26 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	if v183 == v184 {
		goto L60
	} else {
		goto L61
	}
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v28 != int32(1) {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v34 = v33 + v26
	v35 = v25 - v26
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(base.Ui32(int32(41)) < base.Ui32(v36))|base.B2i32(v36 == int32(7)) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v31 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v31)
	goto L7
L15:
	;
	if v158 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36*int32(28))+uint32(_c_F_CopyLoadInputBuf[0])))
	v47 = m.T0[v46].(func(*base.Module, int32, int32) int32)(m, v34, v35)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v49 = int32(0)
	if base.B2i32(v34&int32(3) == v49)|base.B2i32(v35 == v49) != 0 {
		v80 = v34
		v82 = v35
		v83 = base.B2i32(v35 != v49)
		goto L24
	} else {
		goto L25
	}
L19:
	;
	return
L20:
	;
	v158 = v47
	goto L15
L21:
	;
	if v154 != 0 {
		goto L46
	} else {
		goto L47
	}
L22:
	;
	v154 = int32(0)
	goto L21
L23:
	;
	v132 = v125
	v134 = v127
	goto L40
L24:
	;
	if v83 == int32(0) {
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v63 = v34
	v65 = v35
	goto L26
L26:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v68 == int32(0) {
		v125 = v63
		v127 = v65
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v80 = v75
	v82 = v71
	v83 = v73
	goto L24
L28:
	;
	v70 = int32(1)
	v71 = v65 - v70
	v72 = int32(0)
	v73 = base.B2i32(v71 != v72)
	v75 = v63 + v70
	if v75&int32(3) == v72 {
		v80 = v75
		v82 = v71
		v83 = v73
		goto L24
	} else {
		goto L29
	}
L29:
	;
	if v71 != 0 {
		v63 = v75
		v65 = v71
		goto L26
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v88 = int32(0)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if base.B2i32(v88 == v89)|base.B2i32(base.Ui32(v82) < base.Ui32(int32(4))) == v88 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = v80
	v100 = v82
	goto L35
L33:
	;
	v118 = v80
	v120 = v82
	goto L34
L34:
	;
	if v120 == int32(0) {
		goto L22
	} else {
		goto L39
	}
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v105 = v104 ^ int32(0)
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 != v108 {
		v125 = v98
		v127 = v100
		goto L23
	} else {
		goto L37
	}
L36:
	;
	v118 = v113
	v120 = v115
	goto L34
L37:
	;
	v112 = int32(4)
	v113 = v98 + v112
	v115 = v100 - v112
	if base.Ui32(int32(3)) < base.Ui32(v115) {
		v98 = v113
		v100 = v115
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v125 = v118
	v127 = v120
	goto L23
L40:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if int32(0) == v137 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L22
L42:
	;
	v154 = v132
	goto L21
L43:
	;
	goto L44
L44:
	;
	v139 = int32(1)
	v142 = v134 - v139
	if v142 != 0 {
		v132 = v132 + v139
		v134 = v142
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v156 = v154 - v34
	goto L48
L47:
	;
	v156 = v35
	goto L48
L48:
	;
	v158 = v156
	goto L15
L49:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v161 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v180 + v158
	goto L7
L52:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.B2i32(v164 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v164)) != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+337)) = uint8(v178)
	goto L7
L55:
	;
	if v35 < v176 {
		goto L7
	} else {
		goto L59
	}
L56:
	;
	v176 = int32(1)
	goto L58
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v164*int32(28))+uint32(_c_F_CopyLoadInputBuf[1])))
	v176 = v175
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L54
L60:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if v186 != int32(1) {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v193 = v191 - v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v195 = int32(0)
	if base.B2i32(v192 <= v195)|base.B2i32(v193 <= v195) == v195 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v189)
	goto L7
L64:
	;
	if v193 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v205 = v194
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v193
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v207
	*(*uint8)(unsafe.Add(mBase, uint32(v193+v205))) = uint8(v207)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadInputBuf[2]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	goto L70
L67:
	;
	base.MemoryCopy(m, v194, v194+v192, v193)
	goto L69
L68:
	;
	goto L69
L69:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v205 = v204
	goto L66
L70:
	;
	v223 = v215 - v214
	v224 = v212 + v213
	v228 = F_pg_do_encoding_conversion_buf(m, v217, v218, v221, v214+v216, v223, v224, int32(_a_F_CopyLoadInputBuf_0)-v212, int32(1))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L19
	} else {
		goto L71
	}
L71:
	;
	if v228 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+352)))
	if base.B2i32(v232 == int32(0))&base.B2i32(v223 < int32(16)) != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+344)) = v240 + v228
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v244 = F_strlen(m, v224)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v243 + v244
	goto L7
L75:
	;
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+337)) = uint8(v238)
	goto L7
L76:
	;
	goto L6
L77:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+337)))
	if v256 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+348))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v260 == int32(0) {
		goto L4
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v292 != 0 {
		goto L76
	} else {
		goto L87
	}
L81:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+344))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_CopyLoadInputBuf[2]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	goto L82
L82:
	;
	v277 = F_pg_do_encoding_conversion_buf(m, v266, v267, v270, v264+v265, v259-v264, v252+v263, int32(_a_F_CopyLoadInputBuf_0)-v252, int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L83
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L19
	} else {
		goto L84
	}
L84:
	;
	F_errmsg_internal(m, int32(_a_F_CopyLoadInputBuf_1), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L19
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_CopyLoadInputBuf_2), int32(585), int32(_a_F_CopyLoadInputBuf_3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L19
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L19
	} else {
		goto L88
	}
L88:
	;
	goto L5
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyReadAttributesCSV(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
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
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v21 <= v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L7
	} else {
		goto L99
	}
L2:
	;
	m.G0 = v19 + int32(16)
	return v412
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v24 == int32(0) {
		v412 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v52 = l0 + int32(280)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53))) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v54
	goto L12
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_0), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(2127), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
	if v61 <= v60 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_enlargeStringInfo(m, v52, v60)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	v66 = v60
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+304))
	v68 = v66 + v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v73 = v69
	v75 = v67
	v77 = v2
	goto L18
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	v66 = v65
	goto L15
L17:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v390 - v402
	v412 = v393
	goto L2
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+296))
	if v87 <= v77 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+52))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L7
	} else {
		goto L94
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+296)) = v87 << (uint(int32(1)) % 32)
	v94 = F_repalloc(m, v86, v87<<(uint(int32(3))%32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v97 = v86
	goto L22
L22:
	;
	v99 = v77 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v97+v99))) = v73
	if base.Ui32(v75) < base.Ui32(v68) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+300)) = v94
	v97 = v94
	goto L22
L24:
	;
	goto L19
L25:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v349+v99))) = int32(0)
	v354 = v77 + int32(1)
	if v202 != 0 {
		v73 = v193
		v75 = v194
		v77 = v354
		goto L18
	} else {
		goto L93
	}
L26:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v274 != 0 {
		goto L71
	} else {
		goto L72
	}
L27:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v208 == v201 {
		goto L53
	} else {
		goto L54
	}
L28:
	;
	v105 = v75
	v108 = v73
	v110 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v186)
	v193 = v73 + int32(1)
	v194 = v75
	v196 = v73
	v201 = v186
	v202 = v186
	goto L27
L31:
	;
	v121 = v105 + int32(1)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v123 = base.B2i32(v122 == v50)
	if v123 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v180)
	v260 = v160
	v261 = v145 + int32(1)
	v262 = v145
	v267 = v160 - v75
	v268 = v180
	goto L26
L33:
	;
	v143 = v121
	v145 = v108
	goto L43
L34:
	;
	if v122 == v48 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v132 = v105
	v133 = v108
	goto L36
L36:
	;
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v133))) = uint8(v134)
	v136 = v132 - v75
	v138 = v133 + int32(1)
	if v110 == v134 {
		v193 = v138
		v194 = v121
		v196 = v133
		v201 = v136
		v202 = v123
		goto L27
	} else {
		goto L42
	}
L37:
	;
	if base.Ui32(v121) < base.Ui32(v68) {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v122)
	v130 = v108 + int32(1)
	if base.Ui32(v121) < base.Ui32(v68) {
		v105 = v121
		v108 = v130
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L1
L41:
	;
	v132 = v121
	v133 = v130
	goto L36
L42:
	;
	v260 = v121
	v261 = v138
	v262 = v133
	v267 = v136
	v268 = v123
	goto L26
L43:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v160 = v143 + int32(1)
	if base.B2i32(v46 != v157)|base.B2i32(base.Ui32(v68) <= base.Ui32(v160)) != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v160) < base.Ui32(v68) {
		v105 = v160
		v108 = v145
		v110 = int32(1)
		goto L31
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if base.Ui32(v174) < base.Ui32(v68) {
		v143 = v174
		v145 = v145 + int32(1)
		goto L43
	} else {
		goto L51
	}
L47:
	;
	if v157 == v48 {
		goto L45
	} else {
		goto L50
	}
L48:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if base.B2i32(v46 != v163)&base.B2i32(v163 != v48) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v163)
	v174 = v143 + int32(2)
	goto L46
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v157)
	v174 = v160
	goto L46
L51:
	;
	goto L1
L52:
	;
	goto L32
L53:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v201 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	v260 = v194
	v261 = v193
	v262 = v196
	v267 = v201
	v268 = v202
	goto L26
L56:
	;
	if v255 == int32(0) {
		goto L25
	} else {
		goto L69
	}
L57:
	;
	v255 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v216 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v217 = v75
	v218 = v210
	v219 = v201
	v220 = v216
	goto L64
L61:
	;
	v243 = v210
	v247 = int32(0)
	goto L62
L62:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	v255 = v247 - v248
	goto L56
L63:
	;
	v243 = v238
	v247 = v240
	goto L62
L64:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if base.B2i32(v220 != v222)|base.B2i32(v222 == int32(0)) != 0 {
		v238 = v218
		v240 = v220
		goto L63
	} else {
		goto L66
	}
L65:
	;
	v238 = v232
	v240 = int32(0)
	goto L63
L66:
	;
	v228 = v219 - int32(1)
	if v228 == int32(0) {
		v238 = v218
		v240 = v220
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v231 = int32(1)
	v232 = v218 + v231
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	if v233 != 0 {
		v217 = v217 + v231
		v218 = v232
		v219 = v228
		v220 = v233
		goto L64
	} else {
		goto L68
	}
L68:
	;
	goto L65
L69:
	;
	goto L55
L70:
	;
	v348 = v77 + int32(1)
	if v268 != 0 {
		v73 = v261
		v75 = v260
		v77 = v348
		goto L18
	} else {
		goto L92
	}
L71:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	v277 = v275
	goto L73
L72:
	;
	v277 = int32(0)
	goto L73
L73:
	;
	if v277 <= v77 {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v279 == int32(0) {
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v267 != v282 {
		goto L70
	} else {
		goto L76
	}
L76:
	;
	if v267 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v328 != 0 {
		goto L70
	} else {
		goto L90
	}
L78:
	;
	v328 = int32(0)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v289 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v290 = v75
	v291 = v279
	v292 = v267
	v293 = v289
	goto L85
L82:
	;
	v316 = v279
	v320 = int32(0)
	goto L83
L83:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	v328 = v320 - v321
	goto L77
L84:
	;
	v316 = v311
	v320 = v313
	goto L83
L85:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if base.B2i32(v293 != v295)|base.B2i32(v295 == int32(0)) != 0 {
		v311 = v291
		v313 = v293
		goto L84
	} else {
		goto L87
	}
L86:
	;
	v311 = v305
	v313 = int32(0)
	goto L84
L87:
	;
	v301 = v292 - int32(1)
	if v301 == int32(0) {
		v311 = v291
		v313 = v293
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v304 = int32(1)
	v305 = v291 + v304
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	if v306 != 0 {
		v290 = v290 + v304
		v291 = v305
		v292 = v301
		v293 = v306
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v274)+12))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v330+v99)))
	v334 = v332 - int32(1)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v329+v334<<(uint(int32(2))%32))))
	if v338 == int32(0) {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v341+v334))) = uint8(v343)
	goto L70
L92:
	;
	v390 = v262
	v393 = v348
	goto L17
L93:
	;
	v390 = v196
	v393 = v354
	goto L17
L94:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_3), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v356 + v357<<(uint(int32(3))%32) + v334*int32(100) + int32(32)
	v379 = F_errdetail(m, int32(_a_F_CopyReadAttributesCSV_4), v19)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(2280), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_CopyReadAttributesCSV_5), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_CopyReadAttributesCSV_1), int32(2211), int32(_a_F_CopyReadAttributesCSV_2))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_collect_visibility_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v3
	v18 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
	v28 = v26 - int32(109)
	v35 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v28))|base.B2i32(int32(1)<<(uint(v28)%32)&int32(161) == v35) == v35 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = F_RelationGetNumberOfBlocksInFork(m, v23, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L52
	}
L7:
	;
	v45 = F_palloc0(m, v41+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = int32(0)
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v41
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v51
	v59 = F_read_stream_begin_relation(m, int32(12), v18, v23, v51, int32(3), v13+int32(4), v51)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v61 = v3
	goto L11
L11:
	;
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v61 = v59
	goto L11
L13:
	;
	v63 = v45 + int32(8)
	v65 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	if l1 != 0 {
		goto L43
	} else {
		goto L44
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[0]))
	if v76 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v81 = F_visibilitymap_get_status(m, v23, v65, v13+int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v81&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v85 = v65 + v63
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	v88 = v86 | int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v88)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if v81&int32(2) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = v65 + v63
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v96 = v94 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v96)
	goto L28
L27:
	;
	goto L28
L28:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v100 = F_read_stream_next_buffer(m, v61, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v137 = v65 + int32(1)
	if v137 != v41 {
		v65 = v137
		goto L16
	} else {
		goto L42
	}
L32:
	;
	F_LockBufferInternal(m, v100, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v100 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+10)))
	if v123&int32(4) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[1]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v100^int32(-1))<<(uint(int32(2))%32))))
	v122 = v114
	goto L34
L36:
	;
	goto L37
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_collect_visibility_data[2]))
	v122 = v116 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L34
L38:
	;
	v126 = v65 + v63
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v129 = v127 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
	goto L40
L39:
	;
	goto L40
L40:
	;
	F_UnlockReleaseBuffer(m, v100)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L31
L42:
	;
	goto L17
L43:
	;
	F_read_stream_end(m, v61)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v151 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_ReleaseBuffer(m, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	m.G0 = v13 + int32(16)
	return v45
L52:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v168 + int32(4)
	F_errmsg(m, int32(_a_F_collect_visibility_data_0), v13)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v176 = int32(*(*int8)(unsafe.Add(mBase, uint32(v175)+119)))
	F_errdetail_relkind_not_supported(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_collect_visibility_data_1), int32(932), int32(_a_F_collect_visibility_data_2))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_compare3(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	v6 = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v8) < base.Ui32(v7) {
		v24 = v6
	} else {
		v10 = base.B2i32(v7 != v8)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if base.B2i32(v10 == int32(0))&base.B2i32(base.Ui32(v14) < base.Ui32(v13)) != 0 {
			v24 = v6
		} else {
			v24 = int32(0) - (v10 | base.B2i32(v13 != v14))
		}
	}
	return v24
}
func F_compare_fractional_path_costs(m *base.Module, l0 int32, l1 int32, l2 float64) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v49 int32
	_ = v49
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v7 != v8 {
		if v7 < v8 {
			v13 = int32(-1)
		} else {
			v13 = int32(1)
		}
		return v13
	} else {
		if base.F64_le(l2, float64(0))|base.F64_ge(l2, float64(1)) != 0 {
			v20 = int32(-1)
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			if base.F64_lt(v21, v22) != 0 {
				v49 = v20
				return v49
			} else {
				if base.F64_gt(v21, v22) != 0 {
					return int32(1)
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
					if base.F64_lt(v27, v28) != 0 {
						v49 = v20
						return v49
					} else {
						if base.F64_gt(v27, v28) != 0 {
							v49 = int32(1)
							return v49
						} else {
							return int32(0)
						}
					}
				}
			}
		} else {
			v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
			v36 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
			v39 = base.F64_add(base.F64_mul(l2, base.F64_sub(v35, v36)), v36)
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
			v41 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
			v44 = base.F64_add(base.F64_mul(l2, base.F64_sub(v40, v41)), v41)
			if base.F64_lt(v39, v44) != 0 {
				v49 = int32(-1)
			} else {
				v49 = base.F64_lt(v44, v39)
			}
			return v49
		}
	}
}
func F_compare_subnode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	v15 = l1 + l2
	v17 = l0 + int32(2)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v19 = v17 + v18
	v21 = l1
	goto L3
L3:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32 != int32(95) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L1
L5:
	;
	if v18 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v49 = v21
	goto L14
L7:
	;
	if base.Ui32(v21) < base.Ui32(v15) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v36 = F_pg_mblen_range(m, v21, v15)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v68 = v21
	goto L5
L11:
	;
	return int32(0)
L12:
	;
	v40 = v36 + v21
	if base.Ui32(v40) < base.Ui32(v15) {
		v21 = v40
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v54 = F_pg_mblen_range(m, v49, v15)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v68 = v56
	goto L5
L16:
	;
	v56 = v54 + v49
	if base.Ui32(v15) <= base.Ui32(v56) {
		v68 = v56
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v58 != int32(95) {
		v49 = v56
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v77 = v68 - v21
	v78 = v17
	goto L23
L22:
	;
	if base.Ui32(v68) < base.Ui32(v15) {
		v21 = v21 + v77
		goto L3
	} else {
		goto L43
	}
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 != int32(95) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	return int32(0)
L25:
	;
	v133 = F_ltree_label_match(m, v21, v77, v78, v129, l3, l4)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L40
	}
L26:
	;
	v103 = v78
	goto L33
L27:
	;
	if base.Ui32(v78) < base.Ui32(v19) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v95 = F_pg_mblen_range(m, v78, v19)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L31
	}
L30:
	;
	v129 = int32(0)
	goto L25
L31:
	;
	v97 = v95 + v78
	if base.Ui32(v97) < base.Ui32(v19) {
		v78 = v97
		goto L23
	} else {
		goto L32
	}
L32:
	;
	return int32(0)
L33:
	;
	v113 = F_pg_mblen_range(m, v103, v19)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v129 = v115 - v78
	goto L25
L35:
	;
	v115 = v113 + v103
	if base.Ui32(v115) < base.Ui32(v19) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v117 != int32(95) {
		v103 = v115
		goto L33
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L34
L39:
	;
	goto L38
L40:
	;
	if v133 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	v135 = v78 + v129
	if base.Ui32(v135) < base.Ui32(v19) {
		v78 = v135
		goto L23
	} else {
		goto L42
	}
L42:
	;
	goto L24
L43:
	;
	goto L4
}
func F_comparecost_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_compute_new_xmax_infomask(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v682 int32
	_ = v682
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v938 int32
	_ = v938
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1091 int32
	_ = v1091
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	if l1&int32(2048) != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_1), int32(_a_F_compute_new_xmax_infomask_2))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L13
	} else {
		goto L331
	}
L2:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l7))) = uint16(v1054)
	*(*uint16)(unsafe.Add(mBase, uint32(l8))) = uint16(v1050)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1052
	m.G0 = v19 + int32(80)
	return
L3:
	;
	if l5 != 0 {
		goto L318
	} else {
		goto L319
	}
L4:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_3) != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L13
	} else {
		goto L313
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L13
	} else {
		goto L308
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L13
	} else {
		goto L303
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L13
	} else {
		goto L298
	}
L9:
	;
	if l1&int32(_a_F_compute_new_xmax_infomask_4) == int32(_a_F_compute_new_xmax_infomask_5) {
		v1012 = l4
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v430 = l2 & int32(_a_F_compute_new_xmax_infomask_6)
	if v430 != 0 {
		goto L116
	} else {
		goto L117
	}
L12:
	;
	v30 = l1 & int32(128)
	v33 = F_MultiXactIdIsRunning(m, l0, int32(base.Ui32(v30)>>(uint(int32(7))%32)))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v33 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v30 != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l5 != 0 {
		goto L32
	} else {
		goto L33
	}
L18:
	;
	v37 = int32(0)
	v41 = F_GetMultiXactIdMembers(m, l0, v19+int32(76), v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	if int32(0) < v41 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	v47 = v37
	goto L25
L21:
	;
	v78 = v37
	goto L22
L22:
	;
	v93 = F_TransactionIdDidCommit(m, v78)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L13
	} else {
		goto L30
	}
L23:
	;
	F_pfree(m, v45)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L29
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v74 = v72
	goto L23
L25:
	;
	v64 = v45 + v47<<(uint(int32(3))%32)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v74 = int32(0)
	goto L23
L27:
	;
	v69 = v47 + int32(1)
	if v69 != v41 {
		v47 = v69
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v78 = v74
	goto L22
L30:
	;
	if v93 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	v117 = int32(8)
	goto L34
L33:
	;
	v117 = int32(4)
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v117)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v119 == int32(-1) {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	v122 = int32(0)
	v124 = m.G0
	v126 = v124 - int32(16)
	m.G0 = v126
	v131 = F_GetMultiXactIdMembers(m, l0, v126+int32(12), v122)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L39
	}
L36:
	;
	m.G0 = v126 + int32(16)
	v286 = F_GetMultiXactIdMembers(m, v264, v19+int32(76), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L13
	} else {
		goto L71
	}
L37:
	;
	v253 = v240 + v235<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = l3
	v258 = F_MultiXactIdCreateFromMembers(m, v235+int32(1), v240)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L13
	} else {
		goto L68
	}
L38:
	;
	v233 = F_palloc_mul(m, int32(8), int32(1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L13
	} else {
		goto L67
	}
L39:
	;
	if int32(0) <= v131 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	if v131 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = l3
	v228 = F_MultiXactIdCreateFromMembers(m, int32(1), v126+int32(4))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L66
	}
L43:
	;
	v142 = v122
	goto L44
L44:
	;
	v156 = v135 + v142<<(uint(int32(3))%32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 != l3 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v166 = int32(1)
	if v131 <= v166 {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v164 = v142 + int32(1)
	if v164 != v131 {
		v142 = v164
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v159 != v119 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v135)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	v264 = l0
	goto L36
L50:
	;
	goto L45
L51:
	;
	v169 = v166
	goto L53
L52:
	;
	v169 = v131
	goto L53
L53:
	;
	v174 = F_palloc_mul(m, int32(8), v131+int32(1))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	v177 = int32(0)
	v181 = int32(0)
	goto L55
L55:
	;
	v195 = v135 + v181<<(uint(int32(3))%32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v197 = F_TransactionIdIsInProgress(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L13
	} else {
		goto L58
	}
L56:
	;
	v235 = v218
	v240 = v174
	goto L37
L57:
	;
	v221 = v181 + int32(1)
	if v221 != v169 {
		v177 = v218
		v181 = v221
		goto L55
	} else {
		goto L65
	}
L58:
	;
	if v197 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if base.Ui32(v201) < base.Ui32(int32(4)) {
		v218 = v177
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v211 = v174 + v177<<(uint(int32(3))%32)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v214
	v218 = v177 + int32(1)
	goto L57
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v205 = F_TransactionIdDidCommit(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	if v205 == int32(0) {
		v218 = v177
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L61
L65:
	;
	goto L56
L66:
	;
	v264 = v228
	goto L36
L67:
	;
	v235 = int32(0)
	v240 = v233
	goto L37
L68:
	;
	F_pfree(m, v135)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L13
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v240)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L70
	}
L70:
	;
	v264 = v258
	goto L36
L71:
	;
	if v286 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v1050 = v122
	v1052 = v264
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L73:
	;
	goto L74
L74:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v286 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	F_pfree(m, v291)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L13
	} else {
		goto L104
	}
L76:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v291+v365<<(uint(int32(3))%32))+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v380<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v366) < base.Ui32(v383) {
		goto L98
	} else {
		goto L99
	}
L77:
	;
	v294 = int32(0)
	v362 = v122
	v365 = v294
	v366 = v294
	v373 = v10
	goto L76
L78:
	;
	goto L79
L79:
	;
	v300 = int32(0)
	v303 = v300
	v304 = v122
	v307 = v300
	v308 = v300
	v315 = v10
	goto L80
L80:
	;
	v321 = v291 + v307<<(uint(int32(3))%32)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v322<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v308) < base.Ui32(v325) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v286&int32(1) == int32(0) {
		v395 = v351
		v399 = v353
		v406 = v352
		goto L75
	} else {
		goto L97
	}
L82:
	;
	v327 = v325
	goto L84
L83:
	;
	v327 = v308
	goto L84
L84:
	;
	switch v322 - int32(3) {
	case 0:
		goto L88
	case 1:
		v334 = v304
		goto L86
	case 2:
		goto L87
	default:
		v336 = v304
		v337 = v315
		goto L85
	}
L85:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v338 - int32(3) {
	case 0:
		goto L92
	case 1:
		v349 = v336
		goto L90
	case 2:
		goto L91
	default:
		v351 = v336
		v352 = v337
		goto L89
	}
L86:
	;
	v336 = v334
	v337 = int32(1)
	goto L85
L87:
	;
	v334 = v304 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L86
L88:
	;
	v336 = v304 | int32(_a_F_compute_new_xmax_infomask_6)
	v337 = v315
	goto L85
L89:
	;
	if base.Ui32(v327) < base.Ui32(v341) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v351 = v349
	v352 = int32(1)
	goto L89
L91:
	;
	v349 = v336 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L90
L92:
	;
	v351 = v336 | int32(_a_F_compute_new_xmax_infomask_6)
	v352 = v337
	goto L89
L93:
	;
	v353 = v341
	goto L95
L94:
	;
	v353 = v327
	goto L95
L95:
	;
	v354 = int32(2)
	v355 = v307 + v354
	v357 = v303 + v354
	if v357 != v286&int32(2147483646) {
		v303 = v357
		v304 = v351
		v307 = v355
		v308 = v353
		v315 = v352
		goto L80
	} else {
		goto L96
	}
L96:
	;
	goto L81
L97:
	;
	v362 = v351
	v365 = v355
	v366 = v353
	v373 = v352
	goto L76
L98:
	;
	v385 = v383
	goto L100
L99:
	;
	v385 = v366
	goto L100
L100:
	;
	switch v380 - int32(3) {
	case 0:
		goto L103
	case 1:
		v392 = v362
		goto L101
	case 2:
		goto L102
	default:
		v395 = v362
		v399 = v385
		v406 = v373
		goto L75
	}
L101:
	;
	v395 = v392
	v399 = v385
	v406 = int32(1)
	goto L75
L102:
	;
	v392 = v362 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L101
L103:
	;
	v395 = v362 | int32(_a_F_compute_new_xmax_infomask_6)
	v399 = v385
	v406 = v373
	goto L75
L104:
	;
	if v399&int32(-2) == int32(2) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v406 != 0 {
		v1050 = v395
		v1052 = v264
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if v399 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v1050 = v395
	v1052 = v264
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L109:
	;
	v421 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L111
L110:
	;
	v421 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L111
L111:
	;
	if v399 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v424 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L114
L113:
	;
	v424 = v421
	goto L114
L114:
	;
	if v406 != 0 {
		v1050 = v395
		v1052 = v264
		v1054 = v424
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v1050 = v395
	v1052 = v264
	v1054 = v424 | int32(128)
	goto L2
L116:
	;
	v431 = int32(5)
	goto L118
L117:
	;
	v431 = int32(4)
	goto L118
L118:
	;
	if l1&int32(1024) != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	if l5 != 0 {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	v598 = int32(base.Ui32(l1&int32(128))>>(uint(int32(7))%32)) | base.B2i32(l1&int32(80) == int32(64))
	v599 = F_TransactionIdIsInProgress(m, l0)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L13
	} else {
		goto L172
	}
L122:
	;
	v438 = int32(8)
	goto L124
L123:
	;
	v438 = int32(4)
	goto L124
L124:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v438)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v440 == int32(-1) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	v443 = int32(0)
	v444 = F_MultiXactIdCreate(m, l0, v431, l3, v440)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L13
	} else {
		goto L126
	}
L126:
	;
	v449 = F_GetMultiXactIdMembers(m, v444, v19+int32(76), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L13
	} else {
		goto L127
	}
L127:
	;
	if v449 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1050 = v443
	v1052 = v444
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L129:
	;
	goto L130
L130:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v449 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	F_pfree(m, v454)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L13
	} else {
		goto L160
	}
L132:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v454+v528<<(uint(int32(3))%32))+4))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v543<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v529) < base.Ui32(v546) {
		goto L154
	} else {
		goto L155
	}
L133:
	;
	v457 = int32(0)
	v525 = v443
	v528 = v457
	v529 = v457
	v536 = v10
	goto L132
L134:
	;
	goto L135
L135:
	;
	v463 = int32(0)
	v466 = v463
	v467 = v443
	v470 = v463
	v471 = v463
	v478 = v10
	goto L136
L136:
	;
	v484 = v454 + v470<<(uint(int32(3))%32)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+4))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v471) < base.Ui32(v488) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v449&int32(1) == int32(0) {
		v558 = v514
		v562 = v516
		v569 = v515
		goto L131
	} else {
		goto L153
	}
L138:
	;
	v490 = v488
	goto L140
L139:
	;
	v490 = v471
	goto L140
L140:
	;
	switch v485 - int32(3) {
	case 0:
		goto L144
	case 1:
		v497 = v467
		goto L142
	case 2:
		goto L143
	default:
		v499 = v467
		v500 = v478
		goto L141
	}
L141:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v501<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v501 - int32(3) {
	case 0:
		goto L148
	case 1:
		v512 = v499
		goto L146
	case 2:
		goto L147
	default:
		v514 = v499
		v515 = v500
		goto L145
	}
L142:
	;
	v499 = v497
	v500 = int32(1)
	goto L141
L143:
	;
	v497 = v467 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L142
L144:
	;
	v499 = v467 | int32(_a_F_compute_new_xmax_infomask_6)
	v500 = v478
	goto L141
L145:
	;
	if base.Ui32(v490) < base.Ui32(v504) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v514 = v512
	v515 = int32(1)
	goto L145
L147:
	;
	v512 = v499 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L146
L148:
	;
	v514 = v499 | int32(_a_F_compute_new_xmax_infomask_6)
	v515 = v500
	goto L145
L149:
	;
	v516 = v504
	goto L151
L150:
	;
	v516 = v490
	goto L151
L151:
	;
	v517 = int32(2)
	v518 = v470 + v517
	v520 = v466 + v517
	if v520 != v449&int32(2147483646) {
		v466 = v520
		v467 = v514
		v470 = v518
		v471 = v516
		v478 = v515
		goto L136
	} else {
		goto L152
	}
L152:
	;
	goto L137
L153:
	;
	v525 = v514
	v528 = v518
	v529 = v516
	v536 = v515
	goto L132
L154:
	;
	v548 = v546
	goto L156
L155:
	;
	v548 = v529
	goto L156
L156:
	;
	switch v543 - int32(3) {
	case 0:
		goto L159
	case 1:
		v555 = v525
		goto L157
	case 2:
		goto L158
	default:
		v558 = v525
		v562 = v548
		v569 = v536
		goto L131
	}
L157:
	;
	v558 = v555
	v562 = v548
	v569 = int32(1)
	goto L131
L158:
	;
	v555 = v525 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L157
L159:
	;
	v558 = v525 | int32(_a_F_compute_new_xmax_infomask_6)
	v562 = v548
	v569 = v536
	goto L131
L160:
	;
	if v562&int32(-2) == int32(2) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v569 != 0 {
		v1050 = v558
		v1052 = v444
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	if v562 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1050 = v558
	v1052 = v444
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L165:
	;
	v584 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L167
L166:
	;
	v584 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L167
L167:
	;
	if v562 == int32(1) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v587 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L170
L169:
	;
	v587 = v584
	goto L170
L170:
	;
	if v569 != 0 {
		v1050 = v558
		v1052 = v444
		v1054 = v587
		goto L2
	} else {
		goto L171
	}
L171:
	;
	v1050 = v558
	v1052 = v444
	v1054 = v587 | int32(128)
	goto L2
L172:
	;
	if v599 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v598 == int32(0) {
		v631 = v431
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	if v598 != 0 {
		v1012 = l4
		goto L3
	} else {
		goto L245
	}
L176:
	;
	if l0 == l3 {
		goto L189
	} else {
		goto L190
	}
L177:
	;
	switch int32(base.Ui32(l1)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v631 = int32(0)
		goto L176
	case 1, 2:
		goto L180
	case 3:
		goto L181
	case 4:
		goto L178
	default:
		goto L179
	}
L178:
	;
	v631 = int32(1)
	goto L176
L179:
	;
	v615 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L13
	} else {
		goto L185
	}
L180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L181:
	;
	if v430 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v612 = int32(3)
	goto L184
L183:
	;
	v612 = int32(2)
	goto L184
L184:
	;
	v631 = v612
	goto L176
L185:
	;
	if v615 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_12), v19+int32(16))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_13), int32(_a_F_compute_new_xmax_infomask_14))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L13
	} else {
		goto L188
	}
L188:
	;
	v1012 = l4
	goto L3
L189:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v631<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v635) < base.Ui32(l4) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	if l5 != 0 {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	v637 = l4
	goto L194
L193:
	;
	v637 = v635
	goto L194
L194:
	;
	v1012 = v637
	goto L3
L195:
	;
	v642 = int32(8)
	goto L197
L196:
	;
	v642 = int32(4)
	goto L197
L197:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v642)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v644 == int32(-1) {
		goto L6
	} else {
		goto L198
	}
L198:
	;
	v647 = int32(0)
	v648 = F_MultiXactIdCreate(m, l0, v631, l3, v644)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L13
	} else {
		goto L199
	}
L199:
	;
	v653 = F_GetMultiXactIdMembers(m, v648, v19+int32(76), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L13
	} else {
		goto L200
	}
L200:
	;
	if v653 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1050 = v647
	v1052 = v648
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L202:
	;
	goto L203
L203:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v653 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	F_pfree(m, v658)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L13
	} else {
		goto L233
	}
L205:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v658+v732<<(uint(int32(3))%32))+4))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v747<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v733) < base.Ui32(v750) {
		goto L227
	} else {
		goto L228
	}
L206:
	;
	v661 = int32(0)
	v729 = v647
	v732 = v661
	v733 = v661
	v740 = v10
	goto L205
L207:
	;
	goto L208
L208:
	;
	v667 = int32(0)
	v670 = v667
	v671 = v647
	v674 = v667
	v675 = v667
	v682 = v10
	goto L209
L209:
	;
	v688 = v658 + v674<<(uint(int32(3))%32)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v689<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v675) < base.Ui32(v692) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	if v653&int32(1) == int32(0) {
		v762 = v718
		v766 = v720
		v773 = v719
		goto L204
	} else {
		goto L226
	}
L211:
	;
	v694 = v692
	goto L213
L212:
	;
	v694 = v675
	goto L213
L213:
	;
	switch v689 - int32(3) {
	case 0:
		goto L217
	case 1:
		v701 = v671
		goto L215
	case 2:
		goto L216
	default:
		v703 = v671
		v704 = v682
		goto L214
	}
L214:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v688)+12))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v705<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v705 - int32(3) {
	case 0:
		goto L221
	case 1:
		v716 = v703
		goto L219
	case 2:
		goto L220
	default:
		v718 = v703
		v719 = v704
		goto L218
	}
L215:
	;
	v703 = v701
	v704 = int32(1)
	goto L214
L216:
	;
	v701 = v671 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L215
L217:
	;
	v703 = v671 | int32(_a_F_compute_new_xmax_infomask_6)
	v704 = v682
	goto L214
L218:
	;
	if base.Ui32(v694) < base.Ui32(v708) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v718 = v716
	v719 = int32(1)
	goto L218
L220:
	;
	v716 = v703 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L219
L221:
	;
	v718 = v703 | int32(_a_F_compute_new_xmax_infomask_6)
	v719 = v704
	goto L218
L222:
	;
	v720 = v708
	goto L224
L223:
	;
	v720 = v694
	goto L224
L224:
	;
	v721 = int32(2)
	v722 = v674 + v721
	v724 = v670 + v721
	if v724 != v653&int32(2147483646) {
		v670 = v724
		v671 = v718
		v674 = v722
		v675 = v720
		v682 = v719
		goto L209
	} else {
		goto L225
	}
L225:
	;
	goto L210
L226:
	;
	v729 = v718
	v732 = v722
	v733 = v720
	v740 = v719
	goto L205
L227:
	;
	v752 = v750
	goto L229
L228:
	;
	v752 = v733
	goto L229
L229:
	;
	switch v747 - int32(3) {
	case 0:
		goto L232
	case 1:
		v759 = v729
		goto L230
	case 2:
		goto L231
	default:
		v762 = v729
		v766 = v752
		v773 = v740
		goto L204
	}
L230:
	;
	v762 = v759
	v766 = v752
	v773 = int32(1)
	goto L204
L231:
	;
	v759 = v729 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L230
L232:
	;
	v762 = v729 | int32(_a_F_compute_new_xmax_infomask_6)
	v766 = v752
	v773 = v740
	goto L204
L233:
	;
	if v766&int32(-2) == int32(2) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	if v773 != 0 {
		v1050 = v762
		v1052 = v648
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	if v766 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v1050 = v762
	v1052 = v648
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L238:
	;
	v788 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L240
L239:
	;
	v788 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L240
L240:
	;
	if v766 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v791 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L243
L242:
	;
	v791 = v788
	goto L243
L243:
	;
	if v773 != 0 {
		v1050 = v762
		v1052 = v648
		v1054 = v791
		goto L2
	} else {
		goto L244
	}
L244:
	;
	v1050 = v762
	v1052 = v648
	v1054 = v791 | int32(128)
	goto L2
L245:
	;
	v794 = F_TransactionIdDidCommit(m, l0)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L13
	} else {
		goto L246
	}
L246:
	;
	if v794 == int32(0) {
		v1012 = l4
		goto L3
	} else {
		goto L247
	}
L247:
	;
	if l5 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v802 = int32(8)
	goto L250
L249:
	;
	v802 = int32(4)
	goto L250
L250:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l4*int32(12)+v802)+uint32(_c_F_compute_new_xmax_infomask[0])))
	if v804 == int32(-1) {
		goto L5
	} else {
		goto L251
	}
L251:
	;
	v807 = int32(0)
	v808 = F_MultiXactIdCreate(m, l0, v431, l3, v804)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L13
	} else {
		goto L252
	}
L252:
	;
	v813 = F_GetMultiXactIdMembers(m, v808, v19+int32(76), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L13
	} else {
		goto L253
	}
L253:
	;
	if v813 <= int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1050 = v807
	v1052 = v808
	v1054 = int32(_a_F_compute_new_xmax_infomask_7)
	goto L2
L255:
	;
	goto L256
L256:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v19)+76))
	if v813 == int32(1) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	F_pfree(m, v818)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L13
	} else {
		goto L286
	}
L258:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v818+v892<<(uint(int32(3))%32))+4))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v907<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v893) < base.Ui32(v910) {
		goto L280
	} else {
		goto L281
	}
L259:
	;
	v821 = int32(0)
	v889 = v807
	v892 = v821
	v893 = v821
	v900 = v10
	goto L258
L260:
	;
	goto L261
L261:
	;
	v827 = int32(0)
	v830 = v827
	v831 = v807
	v834 = v827
	v835 = v827
	v842 = v10
	goto L262
L262:
	;
	v848 = v818 + v834<<(uint(int32(3))%32)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v848)+4))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v849<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	if base.Ui32(v835) < base.Ui32(v852) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if v813&int32(1) == int32(0) {
		v922 = v878
		v926 = v880
		v933 = v879
		goto L257
	} else {
		goto L279
	}
L264:
	;
	v854 = v852
	goto L266
L265:
	;
	v854 = v835
	goto L266
L266:
	;
	switch v849 - int32(3) {
	case 0:
		goto L270
	case 1:
		v861 = v831
		goto L268
	case 2:
		goto L269
	default:
		v863 = v831
		v864 = v842
		goto L267
	}
L267:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v848)+12))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865<<(uint(int32(2))%32))+uint32(_c_F_compute_new_xmax_infomask[1])))
	switch v865 - int32(3) {
	case 0:
		goto L274
	case 1:
		v876 = v863
		goto L272
	case 2:
		goto L273
	default:
		v878 = v863
		v879 = v864
		goto L271
	}
L268:
	;
	v863 = v861
	v864 = int32(1)
	goto L267
L269:
	;
	v861 = v831 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L268
L270:
	;
	v863 = v831 | int32(_a_F_compute_new_xmax_infomask_6)
	v864 = v842
	goto L267
L271:
	;
	if base.Ui32(v854) < base.Ui32(v868) {
		goto L275
	} else {
		goto L276
	}
L272:
	;
	v878 = v876
	v879 = int32(1)
	goto L271
L273:
	;
	v876 = v863 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L272
L274:
	;
	v878 = v863 | int32(_a_F_compute_new_xmax_infomask_6)
	v879 = v864
	goto L271
L275:
	;
	v880 = v868
	goto L277
L276:
	;
	v880 = v854
	goto L277
L277:
	;
	v881 = int32(2)
	v882 = v834 + v881
	v884 = v830 + v881
	if v884 != v813&int32(2147483646) {
		v830 = v884
		v831 = v878
		v834 = v882
		v835 = v880
		v842 = v879
		goto L262
	} else {
		goto L278
	}
L278:
	;
	goto L263
L279:
	;
	v889 = v878
	v892 = v882
	v893 = v880
	v900 = v879
	goto L258
L280:
	;
	v912 = v910
	goto L282
L281:
	;
	v912 = v893
	goto L282
L282:
	;
	switch v907 - int32(3) {
	case 0:
		goto L285
	case 1:
		v919 = v889
		goto L283
	case 2:
		goto L284
	default:
		v922 = v889
		v926 = v912
		v933 = v900
		goto L257
	}
L283:
	;
	v922 = v919
	v926 = v912
	v933 = int32(1)
	goto L257
L284:
	;
	v919 = v889 | int32(_a_F_compute_new_xmax_infomask_6)
	goto L283
L285:
	;
	v922 = v889 | int32(_a_F_compute_new_xmax_infomask_6)
	v926 = v912
	v933 = v900
	goto L257
L286:
	;
	if v926&int32(-2) == int32(2) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	if v933 != 0 {
		v1050 = v922
		v1052 = v808
		v1054 = int32(_a_F_compute_new_xmax_infomask_8)
		goto L2
	} else {
		goto L290
	}
L288:
	;
	goto L289
L289:
	;
	if v926 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1050 = v922
	v1052 = v808
	v1054 = int32(_a_F_compute_new_xmax_infomask_9)
	goto L2
L291:
	;
	v948 = int32(_a_F_compute_new_xmax_infomask_3)
	goto L293
L292:
	;
	v948 = int32(_a_F_compute_new_xmax_infomask_10)
	goto L293
L293:
	;
	if v926 == int32(1) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v951 = int32(_a_F_compute_new_xmax_infomask_11)
	goto L296
L295:
	;
	v951 = v948
	goto L296
L296:
	;
	if v933 != 0 {
		v1050 = v922
		v1052 = v808
		v1054 = v951
		goto L2
	} else {
		goto L297
	}
L297:
	;
	v1050 = v922
	v1052 = v808
	v1054 = v951 | int32(128)
	goto L2
L298:
	;
	if l5 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v960 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L301
L300:
	;
	v960 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v960
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19-int32(-64))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L13
	} else {
		goto L302
	}
L302:
	;
	goto L1
L303:
	;
	if l5 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L306
L305:
	;
	v974 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L306
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19+int32(48))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L13
	} else {
		goto L307
	}
L307:
	;
	goto L1
L308:
	;
	if l5 != 0 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v988 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L311
L310:
	;
	v988 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v988
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L13
	} else {
		goto L312
	}
L312:
	;
	goto L1
L313:
	;
	if l5 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1000 = int32(_a_F_compute_new_xmax_infomask_15)
	goto L316
L315:
	;
	v1000 = int32(_a_F_compute_new_xmax_infomask_16)
	goto L316
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1000
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = l4
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_17), v19+int32(32))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L13
	} else {
		goto L317
	}
L317:
	;
	goto L1
L318:
	;
	v1024 = int32(0)
	if v1012 == int32(3) {
		goto L321
	} else {
		goto L322
	}
L319:
	;
	goto L320
L320:
	;
	v1030 = int32(0)
	switch v1012 {
	case 0:
		v1050 = v1030
		v1052 = l3
		v1054 = int32(144)
		goto L2
	case 1:
		goto L327
	case 2:
		goto L326
	case 3:
		goto L325
	default:
		goto L324
	}
L321:
	;
	v1029 = int32(_a_F_compute_new_xmax_infomask_6)
	goto L323
L322:
	;
	v1029 = v1024
	goto L323
L323:
	;
	v1050 = v1029
	v1052 = l3
	v1054 = v1024
	goto L2
L324:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L13
	} else {
		goto L328
	}
L325:
	;
	v1050 = int32(_a_F_compute_new_xmax_infomask_6)
	v1052 = l3
	v1054 = int32(192)
	goto L2
L326:
	;
	v1050 = v1030
	v1052 = l3
	v1054 = int32(192)
	goto L2
L327:
	;
	v1050 = v1030
	v1052 = l3
	v1054 = int32(208)
	goto L2
L328:
	;
	F_errmsg_internal(m, int32(_a_F_compute_new_xmax_infomask_18), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L13
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_compute_new_xmax_infomask_0), int32(_a_F_compute_new_xmax_infomask_19), int32(_a_F_compute_new_xmax_infomask_14))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L13
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_concat_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v15 == v4 {
		v27 = v4
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L12
	} else {
		goto L50
	}
L2:
	;
	m.G0 = v13 + int32(32)
	return v187
L3:
	;
	if v27&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L3
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v19 == int32(0) {
		v27 = v4
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v22 != int32(15) {
		v27 = v4
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+13)))
	v27 = v25
	goto L4
L8:
	;
	v32 = l2 + l1<<(uint(int32(4))%32)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+32)))
	if v33 != 0 {
		v187 = v4
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_initStringInfo(m, v13+int32(8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v35 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v40 = F_array_to_text_internal(m, l2, v35, l0, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v187 = v40
	goto L2
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v47 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	v54 = F_MemoryContextAlloc(m, v50, v51*int32(28))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	v108 = v47
	goto L18
L18:
	;
	v113 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v113 {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if l1 < v56 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = l1
	goto L23
L21:
	;
	goto L22
L22:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+16)) = v54
	v108 = v54
	goto L18
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v69 = F_get_fn_expr_argtype(m, v68, v61)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v69 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_getTypeOutputInfo(m, v69, v13+int32(28), v13+int32(27))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	F_fmgr_info_cxt(m, v79, v54+v61*int32(28), v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L28
	}
L28:
	;
	v88 = v61 + int32(1)
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+18)))
	if v88 < v89 {
		v61 = v88
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	v119 = l1
	v121 = v113
	v124 = int32(1)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v172 = v170 + int32(4)
	v173 = F_palloc(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L12
	} else {
		goto L45
	}
L33:
	;
	v130 = l2 + int32(24) + v119<<(uint(int32(4))%32)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+8)))
	if v131 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	if v124 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v152 = v121
	v153 = v124
	goto L37
L37:
	;
	v156 = v119 + int32(1)
	if v156 < base.I32_extend16_s(v152) {
		v119 = v156
		v121 = v152
		v124 = v153
		goto L33
	} else {
		goto L44
	}
L38:
	;
	F_appendStringInfoString(m, v13+int32(8), l0)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v146 = F_OutputFunctionCall(m, v108+v119*int32(28), v134)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	F_appendStringInfoString(m, v13+int32(8), v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+18)))
	v152 = v151
	v153 = int32(0)
	goto L37
L44:
	;
	goto L34
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v172 << (uint(int32(2)) % 32)
	if v170 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	base.MemoryCopy(m, v173+int32(4), v169, v170)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	F_pfree(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	v187 = v173
	goto L2
L50:
	;
	F_errmsg_internal(m, int32(_a_F_concat_internal_0), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_concat_internal_1), int32(_a_F_concat_internal_2), int32(_a_F_concat_internal_3))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_convert_saop_to_hashed_saop_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l0 == v3 {
		v101 = v3
		m.G0 = v9 + int32(16)
		return v101
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v13 != int32(20) {
			v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int32(0)
			} else {
				v101 = v96
				m.G0 = v9 + int32(16)
				return v101
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			if v18 == int32(0) {
				v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v101 = v96
					m.G0 = v9 + int32(16)
					return v101
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v21 != int32(7) {
					v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v101 = v96
						m.G0 = v9 + int32(16)
						return v101
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+32)))
					if v24 != 0 {
						v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							v101 = v96
							m.G0 = v9 + int32(16)
							return v101
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
						if v27 == int32(1) {
							v30 = F_exprType(m, v25)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v38 = F_get_op_hash_functions_ext(m, v26, v30, v9+int32(12), v9+int32(8))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									if v38 == int32(0) {
										v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											v101 = v96
											m.G0 = v9 + int32(16)
											return v101
										}
									} else {
										v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
										if v42 != v43 {
											v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												v101 = v96
												m.G0 = v9 + int32(16)
												return v101
											}
										} else {
											v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
											v49 = F_ArrayGetNItemsSafe(m, v46, v45+int32(16))
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return int32(0)
											} else {
												if v49 < int32(9) {
													v101 = v3
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
													v86 = int32(12)
													v88 = v54
													*(*int32)(unsafe.Add(mBase, uint32(l0+v86))) = v88
													v101 = v3
												}
												m.G0 = v9 + int32(16)
												return v101
											}
										}
									}
								}
							}
						} else {
							v55 = F_get_negator(m, v26)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								if v55 == int32(0) {
									v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v101 = v96
										m.G0 = v9 + int32(16)
										return v101
									}
								} else {
									v59 = F_exprType(m, v25)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										v65 = F_get_op_hash_functions_ext(m, v55, v59, v9+int32(12), v9+int32(8))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return int32(0)
										} else {
											if v65 == int32(0) {
												v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v101 = v96
													m.G0 = v9 + int32(16)
													return v101
												}
											} else {
												v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
												v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
												if v69 != v70 {
													v96 = F_expression_tree_walker_impl(m, l0, int32(922), int32(0))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														v101 = v96
														m.G0 = v9 + int32(16)
														return v101
													}
												} else {
													v72 = int32(16)
													v73 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
													v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
													v77 = F_ArrayGetNItemsSafe(m, v74, v73+v72)
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														if v77 < int32(9) {
															v101 = v3
															m.G0 = v9 + int32(16)
															return v101
														} else {
															v81 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v81
															v83 = F_get_opcode(m, v55)
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return int32(0)
															} else {
																v86 = v72
																v88 = v83
																*(*int32)(unsafe.Add(mBase, uint32(l0+v86))) = v88
																v101 = v3
																m.G0 = v9 + int32(16)
																return v101
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
			}
		}
	}
}
func F_convert_testexpr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == int32(0) {
		v42 = int32(0)
		m.G0 = v7 + int32(16)
		return v42
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v14 = v12 - int32(8)
		if v14 != 0 {
			if v14 == int32(14) {
				v42 = l0
				m.G0 = v7 + int32(16)
				return v42
			} else {
				v40 = F_expression_tree_mutator_impl(m, l0, int32(893), l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					m.G0 = v7 + int32(16)
					return v42
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 != int32(2) {
				v40 = F_expression_tree_mutator_impl(m, l0, int32(893), l1)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = v40
					m.G0 = v7 + int32(16)
					return v42
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v20 <= int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
						F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(678), int32(_a_F_convert_testexpr_mutator_2))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
							F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(678), int32(_a_F_convert_testexpr_mutator_2))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
						if v26 < v20 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v54
								F_errmsg_internal(m, int32(_a_F_convert_testexpr_mutator_0), v7)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_convert_testexpr_mutator_1), int32(678), int32(_a_F_convert_testexpr_mutator_2))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v20<<(uint(int32(2))%32)-int32(4))))
							v35 = F_copyObjectImpl(m, v34)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v42 = v35
								m.G0 = v7 + int32(16)
								return v42
							}
						}
					}
				}
			}
		}
	}
}
func F_core_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_2(m, int32(_a_F_core_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_cost_append(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 float64
	_ = v19
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 float64
	_ = v227
	var v228 float64
	_ = v228
	var v229 float64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 float64
	_ = v234
	var v236 int32
	_ = v236
	var v238 float64
	_ = v238
	var v240 int32
	_ = v240
	var v243 float64
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v254 float64
	_ = v254
	var v257 float64
	_ = v257
	var v259 float64
	_ = v259
	var v262 float64
	_ = v262
	var v265 int64
	_ = v265
	var v266 float64
	_ = v266
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v280 float64
	_ = v280
	var v282 float64
	_ = v282
	var v283 int32
	_ = v283
	var v286 float64
	_ = v286
	var v290 float64
	_ = v290
	var v292 float64
	_ = v292
	var v293 float64
	_ = v293
	var v295 float64
	_ = v295
	var v296 float64
	_ = v296
	var v300 float64
	_ = v300
	var v303 float64
	_ = v303
	var v307 float64
	_ = v307
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v318 float64
	_ = v318
	var v322 float64
	_ = v322
	var v330 float64
	_ = v330
	var v333 float64
	_ = v333
	var v338 int32
	_ = v338
	var v343 float64
	_ = v343
	var v344 float64
	_ = v344
	var v346 float64
	_ = v346
	var v352 int32
	_ = v352
	var v357 float64
	_ = v357
	var v358 float64
	_ = v358
	var v359 float64
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 float64
	_ = v365
	var v366 float64
	_ = v366
	var v369 float64
	_ = v369
	var v370 float64
	_ = v370
	var v371 float64
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 float64
	_ = v377
	var v379 int32
	_ = v379
	var v385 float64
	_ = v385
	var v389 float64
	_ = v389
	var v392 float64
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 float64
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 float64
	_ = v405
	var v406 float64
	_ = v406
	var v410 float64
	_ = v410
	var v414 float64
	_ = v414
	var v417 float64
	_ = v417
	var v420 float64
	_ = v420
	var v421 float64
	_ = v421
	var v423 float64
	_ = v423
	var v425 float64
	_ = v425
	var v427 float64
	_ = v427
	var v430 float64
	_ = v430
	var v432 float64
	_ = v432
	var v434 float64
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 float64
	_ = v438
	var v447 float64
	_ = v447
	var v451 float64
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v476 float64
	_ = v476
	var v478 float64
	_ = v478
	var v479 float64
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v488 float64
	_ = v488
	var v490 float64
	_ = v490
	var v493 float64
	_ = v493
	var v495 float64
	_ = v495
	var v497 float64
	_ = v497
	var v499 int32
	_ = v499
	var v500 float64
	_ = v500
	var v501 float64
	_ = v501
	var v505 float64
	_ = v505
	var v509 float64
	_ = v509
	var v512 float64
	_ = v512
	var v515 float64
	_ = v515
	var v517 float64
	_ = v517
	var v518 float64
	_ = v518
	var v520 float64
	_ = v520
	var v521 float64
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 float64
	_ = v526
	var v535 float64
	_ = v535
	var v539 float64
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v554 int32
	_ = v554
	var v564 float64
	_ = v564
	var v567 float64
	_ = v567
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v617 int32
	_ = v617
	var v618 float64
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 float64
	_ = v723
	var v724 float64
	_ = v724
	var v730 int32
	_ = v730
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 float64
	_ = v770
	var v774 float64
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 float64
	_ = v780
	var v784 float64
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 float64
	_ = v790
	var v794 float64
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 float64
	_ = v800
	var v804 float64
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	var v867 float64
	_ = v867
	var v871 float64
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1009 float64
	_ = v1009
	var v1013 float64
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 float64
	_ = v1019
	var v1023 float64
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 float64
	_ = v1029
	var v1033 float64
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1039 float64
	_ = v1039
	var v1043 float64
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1103 int32
	_ = v1103
	var v1106 float64
	_ = v1106
	var v1110 float64
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1143 float64
	_ = v1143
	var v1147 float64
	_ = v1147
	var v1148 float64
	_ = v1148
	var v1168 float64
	_ = v1168
	var v1169 float64
	_ = v1169
	var v1176 float64
	_ = v1176
	v3 = int32(0)
	v19 = float64(0)
	v25 = int64(0)
	v26 = m.G0
	v28 = v26 - int32(96)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v25
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v25
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v40 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v41 = int64(-4097)
	goto L3
L2:
	;
	v41 = int64(-266241)
	goto L3
L3:
	;
	v44 = base.B2i32(v31|v41 != int64(-1))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v48 = l0 + int32(48)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v49 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	m.G0 = v28 + int32(96)
	return
L7:
	;
	v1176 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[0]))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(base.F64_mul(base.F64_mul(v1176, float64(0.5)), v1169), v1168)
	goto L6
L8:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v52 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v377 = base.F64_convert_i32_s(v40)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[1])))
	if v379 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L11:
	;
	v109 = v3
	goto L20
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if int32(0) < v53 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v57)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v60 <= int32(0) {
		v1168 = v19
		v1169 = v19
		goto L7
	} else {
		goto L16
	}
L15:
	;
	v1168 = v19
	v1169 = v19
	goto L7
L16:
	;
	v65 = v3
	v67 = v44
	v81 = v19
	v82 = v19
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v65<<(uint(int32(2))%32))))
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v92)+32))
	v94 = base.F64_add(v93, v82)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+40))
	v97 = v67 + v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v97
	v99 = *(*float64)(unsafe.Add(mBase, uint32(v92)+56))
	v100 = base.F64_add(v99, v81)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v100
	v103 = v65 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v103 < v104 {
		v65 = v103
		v67 = v97
		v81 = v100
		v82 = v94
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v1168 = v100
	v1169 = v94
	goto L7
L19:
	;
	goto L18
L20:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v109<<(uint(int32(2))%32))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+64))
	v138 = v28 + int32(76)
	if v52 == v136 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v1168 = v371
	v1169 = v359
	goto L7
L22:
	;
	if v216 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v204
	v216 = int32(1)
	goto L22
L24:
	;
	if v52 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v52 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = int32(0)
	v216 = int32(1)
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = int32(0)
	v216 = int32(1)
	goto L22
L29:
	;
	goto L30
L30:
	;
	if v136 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v156
	v216 = v156
	goto L22
L32:
	;
	goto L33
L33:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	v160 = int32(0)
	if v160 < v159 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v163 = v159
	goto L36
L35:
	;
	v163 = v160
	goto L36
L36:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v169 = int32(0)
	goto L37
L37:
	;
	if v169 < v164 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v180 = v176 + v169<<(uint(int32(2))%32)
	goto L41
L40:
	;
	v180 = int32(0)
	goto L41
L41:
	;
	if v169 == v163 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v163
	v216 = base.B2i32(v180 == int32(0))
	goto L22
L43:
	;
	goto L44
L44:
	;
	v186 = base.B2i32(v180 == int32(0))
	if v180 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v169
	v216 = v186
	goto L22
L46:
	;
	goto L47
L47:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	if v190 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v169
	v216 = v186
	goto L22
L49:
	;
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v190+v169<<(uint(int32(2))%32))))
	if v194 != v198 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v169
	v216 = int32(0)
	goto L22
L52:
	;
	v169 = v169 + int32(1)
	goto L37
L54:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v135)+40))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[2])))
	if v221 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v352 = v135
	goto L56
L56:
	;
	v357 = *(*float64)(unsafe.Add(mBase, uint32(v352)+32))
	v358 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v359 = base.F64_add(v357, v358)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v352)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v361 + v362
	v365 = *(*float64)(unsafe.Add(mBase, uint32(v352)+48))
	v366 = *(*float64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = base.F64_add(v365, v366)
	v369 = *(*float64)(unsafe.Add(mBase, uint32(v352)+56))
	v370 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v371 = base.F64_add(v369, v370)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v371
	v374 = v109 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v374 < v375 {
		v109 = v374
		goto L20
	} else {
		goto L83
	}
L57:
	;
	v352 = v28
	goto L56
L58:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v135)+56))
	v240 = v28 + int32(88)
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v135)+32))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+32))
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_cost_append[3]))
	v249 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	v254 = float64(2)
	if base.F64_lt(v243, v254) != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v28)+76))
	if v224 <= int32(0) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v135)+48))
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v135)+56))
	v229 = *(*float64)(unsafe.Add(mBase, uint32(v135)+32))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+32))
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_cost_append[3]))
	v234 = *(*float64)(unsafe.Add(mBase, uint32(l0)+80))
	F_cost_incremental_sort(m, v28, l1, v52, v224, v219, v227, v228, v229, v231, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	return
L62:
	;
	goto L57
L63:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v28)+32)) = v243
	v338 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_append[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = v219 + (v338 ^ int32(1))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v28)+88))
	v344 = base.F64_add(v238, v343)
	*(*float64)(unsafe.Add(mBase, uint32(v28)+48)) = v344
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v28)+80))
	*(*float64)(unsafe.Add(mBase, uint32(v28)+56)) = base.F64_add(v344, v346)
	goto L57
L64:
	;
	v257 = v254
	goto L66
L65:
	;
	v257 = v243
	goto L66
L66:
	;
	v259 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[5]))
	v262 = base.F64_mul(v257, base.F64_add(base.F64_add(v259, v259), float64(0)))
	v265 = base.I64_extend_i32_s(v248) << (uint(int64(10)) % 64)
	v266 = base.F64_convert_i64_s(v265)
	v273 = base.F64_convert_i32_u((v245+int32(7))&int32(-8) + int32(24))
	v275 = base.F64_mul(v243, v273)
	v279 = base.F64_lt(v249, v257) & base.F64_gt(v249, float64(0))
	if v279 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v240))) = v330
	v333 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[5]))
	*(*float64)(unsafe.Add(mBase, uint32(v28+int32(80)))) = base.F64_mul(v257, v333)
	goto L63
L68:
	;
	v280 = base.F64_mul(v249, v273)
	goto L70
L69:
	;
	v280 = v275
	goto L70
L70:
	;
	if base.F64_lt(v266, v280) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v282 = F_log(m, v257)
	mBase = m.M
	v283 = F_tuplesort_merge_order(m, v265)
	mBase = m.M
	v286 = base.F64_mul(base.F64_div(v282, float64(0.693147180559945)), v262)
	*(*float64)(unsafe.Add(mBase, uint32(v240))) = v286
	v290 = base.F64_ceil(base.F64_mul(v275, float64(0.0001220703125)))
	v292 = base.F64_div(v275, v266)
	v293 = base.F64_convert_i32_s(v283)
	if base.F64_gt(v292, v293) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v279 != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v295 = F_log(m, v292)
	mBase = m.M
	v296 = F_log(m, v293)
	mBase = m.M
	v300 = base.F64_ceil(base.F64_div(v295, v296))
	goto L76
L75:
	;
	v300 = float64(1)
	goto L76
L76:
	;
	v303 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[6]))
	v307 = *(*float64)(unsafe.Add(mBase, _c_F_cost_append[7]))
	v330 = base.F64_add(base.F64_mul(base.F64_mul(base.F64_add(v290, v290), v300), base.F64_add(base.F64_mul(v303, float64(0.75)), base.F64_mul(v307, float64(0.25)))), v286)
	goto L67
L77:
	;
	v313 = v249
	goto L79
L78:
	;
	v313 = v257
	goto L79
L79:
	;
	v314 = base.F64_add(v313, v313)
	if base.F64_gt(v257, v314)|base.F64_gt(v275, v266) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v318 = F_log(m, v314)
	mBase = m.M
	v330 = base.F64_mul(base.F64_div(v318, float64(0.693147180559945)), v262)
	goto L67
L81:
	;
	goto L82
L82:
	;
	v322 = F_log(m, v257)
	mBase = m.M
	v330 = base.F64_mul(base.F64_div(v322, float64(0.693147180559945)), v262)
	goto L67
L83:
	;
	goto L21
L84:
	;
	v385 = base.F64_add(base.F64_mul(v377, float64(-0.3)), float64(1))
	if base.F64_gt(v385, float64(0)) != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v392 = v377
	goto L86
L86:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v393 <= int32(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v389 = v385
	goto L89
L88:
	;
	v389 = math.Float64frombits(uint64(0x8000000000000000))
	goto L89
L89:
	;
	v392 = base.F64_add(v389, v377)
	goto L86
L90:
	;
	if v554 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L91:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v554 = v396
	v564 = v19
	v567 = v19
	goto L90
L92:
	;
	goto L93
L93:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v398)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v399
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v401 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v398)+40))
	v436 = v435 + v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v436
	v438 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v432)&int64(9223372036854775807)))|base.F64_gt(v432, v438) != 0 {
		v451 = v438
		goto L104
	} else {
		goto L105
	}
L95:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v398)+24))
	v405 = base.F64_convert_i32_s(v404)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v398)+32))
	if v379 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v398)+32))
	v430 = base.F64_add(base.F64_div(v427, v392), float64(0))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v430
	v432 = v430
	v434 = v19
	goto L94
L98:
	;
	v410 = base.F64_add(base.F64_mul(v405, float64(-0.3)), float64(1))
	if base.F64_gt(v410, float64(0)) != 0 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v417 = v405
	goto L100
L100:
	;
	v420 = float64(0)
	v421 = base.F64_add(base.F64_mul(v406, base.F64_div(v417, v392)), v420)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v421
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v398)+56))
	v425 = base.F64_add(v423, v420)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v425
	v432 = v421
	v434 = v425
	goto L94
L101:
	;
	v414 = v410
	goto L103
L102:
	;
	v414 = math.Float64frombits(uint64(0x8000000000000000))
	goto L103
L103:
	;
	v417 = base.F64_add(v414, v405)
	goto L100
L104:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v451
	v453 = int32(1)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v454 <= v453 {
		v554 = v401
		v564 = v451
		v567 = v434
		goto L90
	} else {
		goto L107
	}
L105:
	;
	v447 = float64(1)
	if base.F64_le(v432, v447) != 0 {
		v451 = v447
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v451 = base.F64_nearest(v432)
	goto L104
L107:
	;
	v459 = v453
	v461 = v436
	v476 = v451
	v478 = v399
	v479 = v434
	goto L108
L108:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v482+v459<<(uint(int32(2))%32))))
	if v459 < v40 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v554 = v401
	v564 = v539
	v567 = v521
	goto L90
L110:
	;
	v488 = *(*float64)(unsafe.Add(mBase, uint32(v486)+48))
	if base.F64_gt(v488, v478) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v493 = v478
	goto L112
L112:
	;
	if v459 < v401 {
		goto L117
	} else {
		goto L118
	}
L113:
	;
	v490 = v478
	goto L115
L114:
	;
	v490 = v488
	goto L115
L115:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v48))) = v490
	v493 = v490
	goto L112
L116:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v486)+40))
	v524 = v461 + v523
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v524
	v526 = float64(1e+100)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v520)&int64(9223372036854775807)))|base.F64_gt(v520, v526) != 0 {
		v539 = v526
		goto L126
	} else {
		goto L127
	}
L117:
	;
	v495 = *(*float64)(unsafe.Add(mBase, uint32(v486)+32))
	v497 = base.F64_add(v476, base.F64_div(v495, v392))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v497
	v520 = v497
	v521 = v479
	goto L116
L118:
	;
	goto L119
L119:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v486)+24))
	v500 = base.F64_convert_i32_s(v499)
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v486)+32))
	if v379 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v505 = base.F64_add(base.F64_mul(v500, float64(-0.3)), float64(1))
	if base.F64_gt(v505, float64(0)) != 0 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	v512 = v500
	goto L122
L122:
	;
	v515 = base.F64_add(base.F64_mul(v501, base.F64_div(v512, v392)), v476)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v515
	v517 = *(*float64)(unsafe.Add(mBase, uint32(v486)+56))
	v518 = base.F64_add(v517, v479)
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = v518
	v520 = v515
	v521 = v518
	goto L116
L123:
	;
	v509 = v505
	goto L125
L124:
	;
	v509 = math.Float64frombits(uint64(0x8000000000000000))
	goto L125
L125:
	;
	v512 = base.F64_add(v509, v500)
	goto L122
L126:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v539
	v542 = v459 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v542 < v543 {
		v459 = v542
		v461 = v524
		v476 = v539
		v478 = v493
		v479 = v521
		goto L108
	} else {
		goto L129
	}
L127:
	;
	v535 = float64(1)
	if base.F64_le(v520, v535) != 0 {
		v539 = v535
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v539 = base.F64_nearest(v520)
	goto L126
L129:
	;
	goto L109
L130:
	;
	v1168 = base.F64_add(v567, float64(0))
	v1169 = v564
	goto L7
L131:
	;
	goto L132
L132:
	;
	if v40 < v554 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v576 = v40
	goto L135
L134:
	;
	v576 = v554
	goto L135
L135:
	;
	v577 = F_palloc_mul(m, int32(8), v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L61
	} else {
		goto L136
	}
L136:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v579 <= int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v1143 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v1147 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1122<<(uint(int32(3))%32))))
	v1148 = *(*float64)(unsafe.Add(mBase, uint32(l0)+56))
	v1168 = base.F64_add(v1147, v1148)
	v1169 = v1143
	goto L7
L138:
	;
	v963 = int32(3)
	v964 = v576 & v963
	v965 = int32(0)
	if base.Ui32(v963) <= base.Ui32(v576-int32(1)) {
		goto L187
	} else {
		goto L188
	}
L139:
	;
	if v658 != 0 {
		goto L149
	} else {
		goto L150
	}
L140:
	;
	v651 = int32(0)
	if v576 <= v651 {
		v1122 = v651
		goto L137
	} else {
		goto L147
	}
L141:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v576 == int32(0) {
		v655 = v579
		v658 = v582
		goto L139
	} else {
		goto L142
	}
L142:
	;
	v589 = int32(0)
	v590 = v582
	goto L143
L143:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v590+v589<<(uint(int32(2))%32))))
	v618 = *(*float64)(unsafe.Add(mBase, uint32(v617)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v577+v589<<(uint(int32(3))%32)))) = v618
	v621 = v589 + int32(1)
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v622 <= v621 {
		goto L140
	} else {
		goto L145
	}
L144:
	;
	v655 = v622
	v658 = v624
	goto L139
L145:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v576 != v621 {
		v589 = v621
		v590 = v624
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	goto L138
L148:
	;
	if int32(0) < v576 {
		goto L138
	} else {
		goto L186
	}
L149:
	;
	v679 = v576
	goto L151
L150:
	;
	v679 = v655
	goto L151
L151:
	;
	if v655 <= v679 {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v684 = v576 & int32(3)
	v692 = v576 - int32(1)
	v696 = v576
	v699 = v679
	goto L153
L153:
	;
	if v696 == v554 {
		goto L148
	} else {
		goto L155
	}
L154:
	;
	goto L148
L155:
	;
	v717 = v577 + v692<<(uint(int32(3))%32)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v718+v699<<(uint(int32(2))%32))))
	v723 = *(*float64)(unsafe.Add(mBase, uint32(v722)+56))
	v724 = *(*float64)(unsafe.Add(mBase, uint32(v717)))
	*(*float64)(unsafe.Add(mBase, uint32(v717))) = base.F64_add(v723, v724)
	if v576 <= int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v904 = int32(1)
	v907 = v699 + v904
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v907 < v908 {
		v692 = v882
		v696 = v696 + v904
		v699 = v907
		goto L153
	} else {
		goto L185
	}
L157:
	;
	v882 = int32(0)
	goto L156
L158:
	;
	goto L159
L159:
	;
	v730 = int32(0)
	if base.B2i32(base.Ui32(v576) < base.Ui32(int32(4))) == v730 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v737 = v730
	v739 = v730
	v740 = v730
	goto L163
L161:
	;
	v815 = v730
	v817 = v730
	goto L162
L162:
	;
	v840 = v815
	v842 = v817
	v851 = v730
	goto L179
L163:
	;
	v761 = int32(3)
	v762 = v740 | v761
	v764 = v740 | int32(2)
	v766 = v740 | int32(1)
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v577+v740<<(uint(v761)%32))))
	v774 = *(*float64)(unsafe.Add(mBase, uint32(v577+v739<<(uint(v761)%32))))
	if base.F64_lt(v770, v774) != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v684 == int32(0) {
		v882 = v806
		goto L156
	} else {
		goto L178
	}
L165:
	;
	v776 = v740
	goto L167
L166:
	;
	v776 = v739
	goto L167
L167:
	;
	v777 = int32(3)
	v780 = *(*float64)(unsafe.Add(mBase, uint32(v577+v766<<(uint(v777)%32))))
	v784 = *(*float64)(unsafe.Add(mBase, uint32(v577+v776<<(uint(v777)%32))))
	if base.F64_lt(v780, v784) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v786 = v766
	goto L170
L169:
	;
	v786 = v776
	goto L170
L170:
	;
	v787 = int32(3)
	v790 = *(*float64)(unsafe.Add(mBase, uint32(v577+v764<<(uint(v787)%32))))
	v794 = *(*float64)(unsafe.Add(mBase, uint32(v577+v786<<(uint(v787)%32))))
	if base.F64_lt(v790, v794) != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v796 = v764
	goto L173
L172:
	;
	v796 = v786
	goto L173
L173:
	;
	v797 = int32(3)
	v800 = *(*float64)(unsafe.Add(mBase, uint32(v577+v762<<(uint(v797)%32))))
	v804 = *(*float64)(unsafe.Add(mBase, uint32(v577+v796<<(uint(v797)%32))))
	if base.F64_lt(v800, v804) != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v806 = v762
	goto L176
L175:
	;
	v806 = v796
	goto L176
L176:
	;
	v807 = int32(4)
	v808 = v740 + v807
	v810 = v737 + v807
	if v810 != v576&int32(2147483644) {
		v737 = v810
		v739 = v806
		v740 = v808
		goto L163
	} else {
		goto L177
	}
L177:
	;
	goto L164
L178:
	;
	v815 = v808
	v817 = v806
	goto L162
L179:
	;
	v864 = int32(3)
	v867 = *(*float64)(unsafe.Add(mBase, uint32(v577+v840<<(uint(v864)%32))))
	v871 = *(*float64)(unsafe.Add(mBase, uint32(v577+v842<<(uint(v864)%32))))
	if base.F64_lt(v867, v871) != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v882 = v873
	goto L156
L181:
	;
	v873 = v840
	goto L183
L182:
	;
	v873 = v842
	goto L183
L183:
	;
	v874 = int32(1)
	v877 = v851 + v874
	if v877 != v684 {
		v840 = v840 + v874
		v842 = v873
		v851 = v877
		goto L179
	} else {
		goto L184
	}
L184:
	;
	goto L180
L185:
	;
	goto L154
L186:
	;
	v1122 = int32(0)
	goto L137
L187:
	;
	v977 = int32(0)
	v978 = v965
	v979 = v965
	goto L190
L188:
	;
	v1056 = v965
	v1057 = v965
	goto L189
L189:
	;
	v1079 = v965
	v1081 = v1056
	v1082 = v1057
	goto L206
L190:
	;
	v1000 = int32(3)
	v1001 = v978 | v1000
	v1003 = v978 | int32(2)
	v1005 = v978 | int32(1)
	v1009 = *(*float64)(unsafe.Add(mBase, uint32(v577+v978<<(uint(v1000)%32))))
	v1013 = *(*float64)(unsafe.Add(mBase, uint32(v577+v979<<(uint(v1000)%32))))
	if base.F64_gt(v1009, v1013) != 0 {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v964 == int32(0) {
		v1122 = v1045
		goto L137
	} else {
		goto L205
	}
L192:
	;
	v1015 = v978
	goto L194
L193:
	;
	v1015 = v979
	goto L194
L194:
	;
	v1016 = int32(3)
	v1019 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1005<<(uint(v1016)%32))))
	v1023 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1015<<(uint(v1016)%32))))
	if base.F64_gt(v1019, v1023) != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1025 = v1005
	goto L197
L196:
	;
	v1025 = v1015
	goto L197
L197:
	;
	v1026 = int32(3)
	v1029 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1003<<(uint(v1026)%32))))
	v1033 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1025<<(uint(v1026)%32))))
	if base.F64_gt(v1029, v1033) != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1035 = v1003
	goto L200
L199:
	;
	v1035 = v1025
	goto L200
L200:
	;
	v1036 = int32(3)
	v1039 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1001<<(uint(v1036)%32))))
	v1043 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1035<<(uint(v1036)%32))))
	if base.F64_gt(v1039, v1043) != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v1045 = v1001
	goto L203
L202:
	;
	v1045 = v1035
	goto L203
L203:
	;
	v1046 = int32(4)
	v1047 = v978 + v1046
	v1049 = v977 + v1046
	if v1049 != v576&int32(-4) {
		v977 = v1049
		v978 = v1047
		v979 = v1045
		goto L190
	} else {
		goto L204
	}
L204:
	;
	goto L191
L205:
	;
	v1056 = v1047
	v1057 = v1045
	goto L189
L206:
	;
	v1103 = int32(3)
	v1106 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1081<<(uint(v1103)%32))))
	v1110 = *(*float64)(unsafe.Add(mBase, uint32(v577+v1082<<(uint(v1103)%32))))
	if base.F64_gt(v1106, v1110) != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	v1122 = v1112
	goto L137
L208:
	;
	v1112 = v1081
	goto L210
L209:
	;
	v1112 = v1082
	goto L210
L210:
	;
	v1113 = int32(1)
	v1116 = v1079 + v1113
	if v1116 != v964 {
		v1079 = v1116
		v1081 = v1081 + v1113
		v1082 = v1112
		goto L206
	} else {
		goto L211
	}
L211:
	;
	goto L207
}
func F_countVariablesFromJsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	if l0 == int32(0) {
		return base.B2i32(l0 != int32(0))
	} else {
		v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
		if v4&int32(32) != 0 {
			return base.B2i32(l0 != int32(0))
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_countVariablesFromJsonb_0), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v22 = F_errdetail(m, int32(_a_F_countVariablesFromJsonb_1), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_countVariablesFromJsonb_2), int32(3481), int32(_a_F_countVariablesFromJsonb_3))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
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
	}
}
func F_countitem_compare_count(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	return base.B2i32(v7 < v5) - base.B2i32(v5 < v7)
}
