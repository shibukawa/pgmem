package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_finalize_grouping_exprs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v283 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v283)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v286 = F_finalize_grouping_exprs_walker(m, v285, l1)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L26
	} else {
		goto L68
	}
L2:
	;
	return int32(0)
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 - int32(7) {
	case 0, 1:
		goto L2
	case 2:
		goto L7
	case 3:
		goto L6
	default:
		v231 = v14
		goto L5
	}
L4:
	;
	v267 = F_expression_tree_walker_impl(m, l0, int32(480), l1)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L26
	} else {
		goto L67
	}
L5:
	;
	if v231 != int32(67) {
		goto L4
	} else {
		goto L65
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v21 == v22 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v17 == v18 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v17 <= v18 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v218 = v21
	v226 = v22
	goto L12
L12:
	;
	if v226 < v218 {
		goto L2
	} else {
		goto L64
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v207
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v218 = v205
	v226 = v214
	goto L12
L14:
	;
	v205 = v21
	v207 = v3
	goto L13
L15:
	;
	goto L16
L16:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v27 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v205 = v201
	v207 = v195
	goto L13
L18:
	;
	v195 = v3
	goto L17
L19:
	;
	goto L20
L20:
	;
	v35 = v3
	v37 = v3
	goto L21
L21:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v37<<(uint(int32(2))%32))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v46 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L26
	} else {
		goto L59
	}
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v51 = F_flatten_join_alias_vars(m, int32(0), v50, v45)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v55 = v45
	goto L25
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(6) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	return int32(0)
L27:
	;
	v55 = v51
	goto L25
L28:
	;
	goto L22
L29:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	if v150 == int32(0) {
		goto L28
	} else {
		goto L56
	}
L30:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+28))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v59 != v60 {
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v105 != int32(1) {
		goto L28
	} else {
		goto L47
	}
L33:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v62 == int32(0) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v65 <= int32(0) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	v68 = int32(0)
	if v68 < v65 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v71 = v65
	goto L38
L37:
	;
	v71 = v68
	goto L38
L38:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v76 = int32(0)
	goto L39
L39:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72+v76<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 != int32(6) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L28
L41:
	;
	v103 = v76 + int32(1)
	if v103 != v71 {
		v76 = v103
		goto L39
	} else {
		goto L46
	}
L42:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v93 != v94 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)))
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+8)))
	if v96 != v97 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	if v99 == int32(0) {
		v147 = v88
		goto L29
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L40
L47:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v108 != 0 {
		goto L28
	} else {
		goto L48
	}
L48:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v109 == int32(0) {
		goto L28
	} else {
		goto L49
	}
L49:
	;
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v113 <= v112 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v118 = v112
	goto L51
L51:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v118<<(uint(int32(2))%32))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v133 = F_equal(m, v55, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L26
	} else {
		goto L53
	}
L52:
	;
	goto L28
L53:
	;
	if v133 != 0 {
		v147 = v131
		goto L29
	} else {
		goto L54
	}
L54:
	;
	v136 = v118 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v136 < v137 {
		v118 = v136
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v153 = F_lappend_int(m, v35, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L26
	} else {
		goto L57
	}
L57:
	;
	v156 = v37 + int32(1)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v157 <= v156 {
		v195 = v153
		goto L17
	} else {
		goto L58
	}
L58:
	;
	v35 = v153
	v37 = v156
	goto L21
L59:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(306424), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v182 = F_exprLocation(m, v55)
	mBase = m.M
	F_parser_errposition(m, v181, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(498392), int32(1727), int32(221426))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L26
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
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = v228
	goto L5
L65:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v242 + int32(1)
	v248 = F_query_tree_walker_impl(m, l0, int32(480), l1, int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v250 - int32(1)
	return v248
L67:
	;
	return v267
L68:
	;
	v288 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v288)
	return v286
}
func F_gather_grouping_paths(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 float64
	_ = v158
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v162 int32
	_ = v162
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	var v177 float64
	_ = v177
	var v185 float64
	_ = v185
	var v189 float64
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = v16
	goto L3
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v19 < v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = F_list_copy_head(m, v15, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v23 = v15
	goto L6
L6:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v23 = v21
	goto L6
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if int32(0) < v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v39 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v13 + int32(16)
	return
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v39<<(uint(int32(2))%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v50 = v13 + int32(12)
	if v23 == v48 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L12
L15:
	;
	v200 = v39 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v200 < v201 {
		v39 = v200
		goto L13
	} else {
		goto L75
	}
L16:
	;
	if v128 != 0 {
		goto L15
	} else {
		goto L48
	}
L17:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v116
	v128 = int32(1)
	goto L16
L18:
	;
	if v23 != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v23 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
	v128 = int32(1)
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = int32(0)
	v128 = int32(1)
	goto L16
L23:
	;
	goto L24
L24:
	;
	if v48 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v68
	v128 = v68
	goto L16
L26:
	;
	goto L27
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v72 = int32(0)
	if v72 < v71 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v75 = v71
	goto L30
L29:
	;
	v75 = v72
	goto L30
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v80 = int32(0)
	goto L31
L31:
	;
	if v80 < v76 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v92 = v88 + v80<<(uint(int32(2))%32)
	goto L35
L34:
	;
	v92 = int32(0)
	goto L35
L35:
	;
	if v80 == v75 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v75
	v128 = base.B2i32(v92 == int32(0))
	goto L16
L37:
	;
	goto L38
L38:
	;
	v98 = base.B2i32(v92 == int32(0))
	if v92 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v80
	v128 = v98
	goto L16
L40:
	;
	goto L41
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v105 = v102 + v80<<(uint(int32(2))%32)
	if v105 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v80
	v128 = v98
	goto L16
L43:
	;
	goto L44
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v109 != v110 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v80
	v128 = int32(0)
	goto L16
L46:
	;
	v80 = v80 + int32(1)
	goto L31
L48:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v47 != v32 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v158 = *(*float64)(unsafe.Add(mBase, uint32(v154)+32))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+24))
	v160 = base.F64_convert_i32_s(v159)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _consts[381])))
	if v162 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L50:
	;
	v151 = F_create_incremental_sort_path(m, l0, l1, v47, v23, v129, float64(-1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L7
	} else {
		goto L61
	}
L51:
	;
	v147 = F_create_sort_path(m, l1, v47, v23, float64(-1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L60
	}
L52:
	;
	if v129 == int32(0) {
		goto L15
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v129 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L55:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	if v134 == int32(0) {
		goto L15
	} else {
		goto L56
	}
L56:
	;
	if v134 == int32(0) {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L50
L58:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, _consts[394])))
	if v142&int32(1) != 0 {
		goto L50
	} else {
		goto L59
	}
L59:
	;
	goto L51
L60:
	;
	v154 = v147
	goto L49
L61:
	;
	v154 = v151
	goto L49
L62:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v192 = F_create_gather_merge_path(m, l0, l1, v154, v191, v23, v13)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L73
	}
L63:
	;
	v168 = base.F64_add(base.F64_mul(v160, float64(-0.3)), float64(1))
	if base.F64_gt(v168, float64(0)) != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v174 = v160
	goto L65
L65:
	;
	v176 = float64(1e+100)
	v177 = base.F64_mul(v158, v174)
	if base.F64_gt(v177, v176) != 0 {
		v189 = v176
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v172 = v168
	goto L68
L67:
	;
	v172 = math.Float64frombits(uint64(0x8000000000000000))
	goto L68
L68:
	;
	v174 = base.F64_add(v172, v160)
	goto L65
L69:
	;
	goto L62
L70:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v177)&int64(9223372036854775807)) {
		v189 = v176
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v185 = float64(1)
	if base.F64_le(v177, v185) != 0 {
		v189 = v185
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v189 = base.F64_nearest(v177)
	goto L69
L73:
	;
	F_add_path(m, l1, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	goto L15
L75:
	;
	goto L14
}
