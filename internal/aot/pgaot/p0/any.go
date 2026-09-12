package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_ANY_sublink_to_join(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v18 = F_pull_varnos_of_level(m, v4, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v239
L2:
	;
	return int32(0)
L3:
	;
	v22 = int32(0)
	if v18 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v75 == int32(0) {
		v239 = v4
		goto L1
	} else {
		goto L18
	}
L5:
	;
	v75 = int32(1)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if l2 == int32(0) {
		v66 = v22
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v75 = v66
	goto L4
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v32 < v31 {
		v66 = v22
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(1)
	if v31 <= v34 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v34
	goto L13
L12:
	;
	v37 = v31
	goto L13
L13:
	;
	v38 = int32(8)
	v43 = int32(0)
	goto L14
L14:
	;
	v50 = v43 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v18+v38+v50)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+(l2+v38))))
	v57 = v52 & (v54 ^ int32(-1))
	v59 = base.B2i32(v57 == int32(0))
	if v57 != 0 {
		v66 = v59
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v66 = v59
	goto L8
L16:
	;
	v61 = v43 + int32(1)
	if v61 != v37 {
		v43 = v61
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v79 = F_pull_varnos(m, l0, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v79 == int32(0) {
		v239 = v4
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v83 = int32(0)
	if v79 == v83 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v136 == int32(0) {
		v239 = v4
		goto L1
	} else {
		goto L35
	}
L22:
	;
	v136 = int32(1)
	goto L21
L23:
	;
	goto L24
L24:
	;
	if l2 == int32(0) {
		v127 = v83
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v136 = v127
	goto L21
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v93 < v92 {
		v127 = v83
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v95 = int32(1)
	if v92 <= v95 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = v95
	goto L30
L29:
	;
	v98 = v92
	goto L30
L30:
	;
	v99 = int32(8)
	v104 = int32(0)
	goto L31
L31:
	;
	v111 = v104 << (uint(int32(2)) % 32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v79+v99+v111)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+(l2+v99))))
	v118 = v113 & (v115 ^ int32(-1))
	v120 = base.B2i32(v118 == int32(0))
	if v118 != 0 {
		v127 = v120
		goto L25
	} else {
		goto L33
	}
L32:
	;
	v127 = v120
	goto L25
L33:
	;
	v122 = v104 + int32(1)
	if v122 != v98 {
		v104 = v122
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v140 = F_contain_volatile_functions(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v140 != 0 {
		v239 = v4
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v142 = int32(0)
	v144 = F_make_parsestate(m, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v148 = F_makeAlias(m, int32(15552), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v150 = int32(0)
	v153 = F_addRangeTableEntryForSubquery(m, v144, v17, v148, base.B2i32(v18 != v150), v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v157 = F_lappend(m, v155, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v157
	if v157 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v162 = v161
	goto L44
L43:
	;
	v162 = int32(0)
	goto L44
L44:
	;
	v164 = F_palloc0(m, int32(8))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v164)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v164))) = int32(63)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	if v169 == int32(0) {
		v208 = v142
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l0
	v219 = F_convert_testexpr_mutator(m, v214, v13+int32(8))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L59
	}
L47:
	;
	v172 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	if v173 <= v172 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v208 = v142
	goto L46
L49:
	;
	goto L50
L50:
	;
	v179 = v172
	v180 = v142
	v182 = v173
	goto L51
L51:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v179<<(uint(int32(2))%32))))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+26)))
	if v191 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v208 = v199
	goto L46
L53:
	;
	v194 = F_makeVarFromTargetEntry(m, v162, v190)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	v199 = v180
	v200 = v182
	goto L55
L55:
	;
	v202 = v179 + int32(1)
	if v202 < v200 {
		v179 = v202
		v180 = v199
		v182 = v200
		goto L51
	} else {
		goto L58
	}
L56:
	;
	v196 = F_lappend(m, v180, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v199 = v196
	v200 = v198
	goto L55
L58:
	;
	goto L52
L59:
	;
	v222 = F_palloc0(m, int32(40))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v224 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v222)+32)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v222)+28)) = v219
	*(*int64)(unsafe.Add(mBase, uint32(v222)+20)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v164
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+12)) = v230
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+8)) = uint8(v230)
	*(*int64)(unsafe.Add(mBase, uint32(v222))) = int64(17179869248)
	v239 = v222
	goto L1
}
func F_executeAnyItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
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
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v147 int32
	_ = v147
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	F_check_stack_depth(m)
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
	if base.Ui32(l6) < base.Ui32(l4) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(48)
	return v147
L4:
	;
	v147 = int32(1)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = F_JsonbIteratorInit(m, l2)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v25
	v28 = int32(1)
	v44 = v28
	goto L8
L8:
	;
	v53 = F_JsonbIteratorNext(m, v17+int32(36), v17+int32(16), int32(1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v147 = int32(0)
	goto L3
L10:
	;
	if v62&int32(-2) != int32(2) {
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v60 = F_JsonbIteratorNext(m, v17+int32(36), v17+int32(16), int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	switch v53 {
	case 0:
		v147 = v44
		goto L3
	case 1:
		goto L11
	default:
		v62 = v53
		goto L10
	}
L13:
	;
	v62 = v60
	goto L10
L14:
	;
	if base.Ui32(l4) < base.Ui32(l5) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L9
L16:
	;
	if base.Ui32(l6) <= base.Ui32(l4) {
		v44 = v121
		goto L8
	} else {
		goto L46
	}
L17:
	;
	if l5&l6 != int32(-1) {
		v121 = v44
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v68 == int32(18) {
		v121 = v44
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if l7 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if l3 == int32(0) {
		goto L15
	} else {
		goto L36
	}
L25:
	;
	v85 = int32(2)
	if v83 == v85 {
		v147 = v85
		goto L3
	} else {
		goto L31
	}
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v72)
	v76 = F_executeItemOptUnwrapTarget(m, l0, l1, v17+int32(16), l3, l8)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v81 = F_executeItemOptUnwrapTarget(m, l0, l1, v17+int32(16), l3, l8)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)) = uint8(v71)
	v83 = v76
	goto L25
L30:
	;
	v83 = v81
	goto L25
L31:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v121 = v83
	goto L16
L33:
	;
	goto L34
L34:
	;
	if v83 != 0 {
		v121 = v83
		goto L16
	} else {
		goto L35
	}
L35:
	;
	goto L15
L36:
	;
	v91 = F_palloc(m, int32(20))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v99 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v91
	v108 = F_list_make2_impl(m, v17+int32(12), v17+int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v113 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v108
	v121 = v44
	goto L16
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v91
	v121 = v44
	goto L16
L43:
	;
	goto L44
L44:
	;
	v117 = F_lappend(m, v113, v91)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v117
	v121 = v44
	goto L16
L46:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v124 != int32(18) {
		v44 = v121
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v129 = F_executeAnyItem(m, l0, l1, v128, l3, l4+v28, l5, l6, l7, l8)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v129 == int32(2) {
		v147 = int32(2)
		goto L3
	} else {
		goto L49
	}
L49:
	;
	if l3 != 0 {
		v44 = v129
		goto L8
	} else {
		goto L50
	}
L50:
	;
	if v129 != 0 {
		v44 = v129
		goto L8
	} else {
		goto L51
	}
L51:
	;
	goto L15
}
func F_has_any_column_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_textToQualifiedNameList(m, v7)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = F_makeRangeVarFromNameList(m, v14)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = int32(0)
					v22 = F_RangeVarGetRelidExtended(m, v16, v18, v18, v18, v18)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = F_convert_any_priv_string(m, v12, int32(1689536))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = F_pg_class_aclcheck(m, v22, v5, v25)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								if v27 == int32(0) {
									return int32(1)
								} else {
									v34 = F_pg_attribute_aclcheck_all(m, v22, v5, v25, int32(1))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return int32(0)
									} else {
										return base.B2i32(v34 == int32(0))
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
