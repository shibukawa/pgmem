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
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
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
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v279 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v279)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v282 = F_finalize_grouping_exprs_walker(m, v281, l1)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L23
	} else {
		goto L65
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
		v227 = v14
		goto L5
	}
L4:
	;
	v263 = F_expression_tree_walker_impl(m, l0, int32(480), l1)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L23
	} else {
		goto L64
	}
L5:
	;
	if v227 != int32(67) {
		goto L4
	} else {
		goto L62
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
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v213 = v21
	v222 = v22
	goto L12
L12:
	;
	if v222 < v213 {
		goto L2
	} else {
		goto L61
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 <= int32(0) {
		v194 = v3
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v200 = v21
	v206 = v3
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v206
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v213 = v200
	v222 = v210
	goto L12
L16:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v200 = v197
	v206 = v194
	goto L15
L17:
	;
	v34 = v3
	v36 = v3
	goto L18
L18:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v34<<(uint(int32(2))%32))))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v44 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L56
	}
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v47 = F_flatten_join_alias_vars(m, int32(0), v46, v43)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v51 = v43
	goto L22
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 == int32(6) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	return int32(0)
L24:
	;
	v51 = v47
	goto L22
L25:
	;
	goto L19
L26:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v146 == int32(0) {
		goto L25
	} else {
		goto L53
	}
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v55 != v56 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v101 != int32(1) {
		goto L25
	} else {
		goto L44
	}
L30:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v58 == int32(0) {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v64 = int32(0)
	if v64 < v61 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v67 = v61
	goto L35
L34:
	;
	v67 = v64
	goto L35
L35:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v72 = int32(0)
	goto L36
L36:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v68+v72<<(uint(int32(2))%32))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 != int32(6) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L25
L38:
	;
	v99 = v72 + int32(1)
	if v99 != v67 {
		v72 = v99
		goto L36
	} else {
		goto L43
	}
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v89 != v90 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+8)))
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+8)))
	if v92 != v93 {
		goto L38
	} else {
		goto L41
	}
L41:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+28))
	if v95 == int32(0) {
		v142 = v84
		goto L26
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	goto L37
L44:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v104 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v105 == int32(0) {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	v108 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v109 <= v108 {
		goto L25
	} else {
		goto L47
	}
L47:
	;
	v114 = v108
	goto L48
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v114<<(uint(int32(2))%32))))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	v129 = F_equal(m, v51, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L23
	} else {
		goto L50
	}
L49:
	;
	goto L25
L50:
	;
	if v129 != 0 {
		v142 = v127
		goto L26
	} else {
		goto L51
	}
L51:
	;
	v132 = v114 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v132 < v133 {
		v114 = v132
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v149 = F_lappend_int(m, v36, v146)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L23
	} else {
		goto L54
	}
L54:
	;
	v152 = v34 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v152 < v153 {
		v34 = v152
		v36 = v149
		goto L18
	} else {
		goto L55
	}
L55:
	;
	v194 = v149
	goto L16
L56:
	;
	F_errcode(m, int32(50364548))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_finalize_grouping_exprs_walker_0), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v178 = F_exprLocation(m, v51)
	mBase = m.M
	F_parser_errposition(m, v177, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_finalize_grouping_exprs_walker_1), int32(1727), int32(_a_F_finalize_grouping_exprs_walker_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L23
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
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v227 = v224
	goto L5
L62:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v238 + int32(1)
	v244 = F_query_tree_walker_impl(m, l0, int32(480), l1, int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L23
	} else {
		goto L63
	}
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v246 - int32(1)
	return v244
L64:
	;
	return v263
L65:
	;
	v284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v284)
	return v282
}
func F_gather_grouping_paths(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 float64
	_ = v159
	var v160 int32
	_ = v160
	var v161 float64
	_ = v161
	var v163 int32
	_ = v163
	var v169 float64
	_ = v169
	var v173 float64
	_ = v173
	var v175 float64
	_ = v175
	var v177 float64
	_ = v177
	var v178 float64
	_ = v178
	var v187 float64
	_ = v187
	var v191 float64
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L3
L2:
	;
	v19 = int32(0)
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v20 < v19 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = F_list_copy_head(m, v16, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v24 = v16
	goto L6
L6:
	;
	F_generate_useful_gather_paths(m, l0, l1, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v24 = v22
	goto L6
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if int32(0) < v29 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v42 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v14 + int32(16)
	return
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v42<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v52 = v14 + int32(12)
	if v24 == v50 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	goto L12
L15:
	;
	v202 = v42 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v202 < v203 {
		v42 = v202
		goto L13
	} else {
		goto L76
	}
L16:
	;
	if v130 != 0 {
		goto L15
	} else {
		goto L48
	}
L17:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v118
	v130 = int32(1)
	goto L16
L18:
	;
	if v24 != 0 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v24 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	v130 = int32(1)
	goto L16
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = int32(0)
	v130 = int32(1)
	goto L16
L23:
	;
	goto L24
L24:
	;
	if v50 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v70
	v130 = v70
	goto L16
L26:
	;
	goto L27
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v74 = int32(0)
	if v74 < v73 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v77 = v73
	goto L30
L29:
	;
	v77 = v74
	goto L30
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v83 = int32(0)
	goto L31
L31:
	;
	if v83 < v78 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v94 = v90 + v83<<(uint(int32(2))%32)
	goto L35
L34:
	;
	v94 = int32(0)
	goto L35
L35:
	;
	if v83 == v77 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v77
	v130 = base.B2i32(v94 == int32(0))
	goto L16
L37:
	;
	goto L38
L38:
	;
	v100 = base.B2i32(v94 == int32(0))
	if v94 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = v100
	goto L16
L40:
	;
	goto L41
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v104 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = v100
	goto L16
L43:
	;
	goto L44
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104+v83<<(uint(int32(2))%32))))
	if v108 != v112 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v83
	v130 = int32(0)
	goto L16
L46:
	;
	v83 = v83 + int32(1)
	goto L31
L48:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gather_grouping_paths[0])))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v49 == v33 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v142&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v142 = v132
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v133 == int32(0) {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	v137 = int32(1)
	if v132&v137 == int32(0) {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	v142 = v137
	goto L49
L55:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v155)+32))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+24))
	v161 = base.F64_convert_i32_s(v160)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gather_grouping_paths[1])))
	if v163 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L56:
	;
	v146 = v133
	goto L58
L57:
	;
	v146 = int32(0)
	goto L58
L58:
	;
	if v146 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v150 = F_create_sort_path(m, l1, v49, v24, float64(-1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v153 = F_create_incremental_sort_path(m, l0, l1, v49, v24, v133, float64(-1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L7
	} else {
		goto L63
	}
L62:
	;
	v155 = v150
	goto L55
L63:
	;
	v155 = v153
	goto L55
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v194 = F_create_gather_merge_path(m, l0, l1, v155, v193, v24, v14)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L74
	}
L65:
	;
	v169 = base.F64_add(base.F64_mul(v161, float64(-0.3)), float64(1))
	if base.F64_gt(v169, float64(0)) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v175 = v161
	goto L67
L67:
	;
	v177 = float64(1e+100)
	v178 = base.F64_mul(v159, v175)
	if base.F64_gt(v178, v177)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v178)&int64(9223372036854775807))) != 0 {
		v191 = v177
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v173 = v169
	goto L70
L69:
	;
	v173 = math.Float64frombits(uint64(0x8000000000000000))
	goto L70
L70:
	;
	v175 = base.F64_add(v173, v161)
	goto L67
L71:
	;
	goto L64
L72:
	;
	v187 = float64(1)
	if base.F64_le(v178, v187) != 0 {
		v191 = v187
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v191 = base.F64_nearest(v178)
	goto L71
L74:
	;
	F_add_path(m, l1, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L15
L76:
	;
	goto L14
}
