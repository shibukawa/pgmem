package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fix_join_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 float64
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 float64
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 float64
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v290
L2:
	;
	v290 = v3
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = l0
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v31 != int32(319) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v290 = int32(0)
	goto L1
L7:
	;
	if v275 != 0 {
		v19 = v275
		goto L5
	} else {
		goto L87
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L23
	} else {
		goto L84
	}
L9:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v147 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L10:
	;
	if v31 != int32(6) {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v124 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v76 != 0 {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v72 = F_search_indexed_tlist_for_var(m, v19, v67, int32(-2), v70, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L23
	} else {
		goto L28
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v37 != 0 {
		goto L8
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 == int32(0) {
		goto L14
	} else {
		goto L27
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v41 == int32(0) {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v44 == v41 {
		v67 = v38
		goto L15
	} else {
		goto L22
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v52
	F_errmsg_internal(m, int32(32137), v15)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(494232), int32(3148), int32(208987))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v67 = v64
	goto L15
L28:
	;
	if v72 != 0 {
		v290 = v72
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v80 = F_search_indexed_tlist_for_var(m, v19, v76, int32(-1), v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L23
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v83 == v84 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v80 != 0 {
		v290 = v80
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v87 = F_palloc(m, int32(48))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L23
	} else {
		goto L40
	}
L38:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v19)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+40)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v93
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v97
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
	*(*int64)(unsafe.Add(mBase, uint32(v87))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v101 + v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v87)+36))
	if v105 == int32(0) {
		v290 = v87
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v108 + v105
	v290 = v87
	goto L1
L40:
	;
	F_errmsg_internal(m, int32(117633), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(494232), int32(3186), int32(208987))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v135 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+8)))
	if v127 != int32(1) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v132 = F_search_indexed_tlist_for_phv(m, v19, v124, int32(-2), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L23
	} else {
		goto L46
	}
L46:
	;
	if v132 != 0 {
		v290 = v132
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v275 = v146
	goto L7
L49:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+8)))
	if v138 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v143 = F_search_indexed_tlist_for_phv(m, v19, v135, int32(-1), v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L23
	} else {
		goto L51
	}
L51:
	;
	if v143 != 0 {
		v290 = v143
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v170 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L54:
	;
	if v31 == int32(7) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+9)))
	if v152&int32(1) == int32(0) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v158 = F_tlist_member(m, v19, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L23
	} else {
		goto L57
	}
L57:
	;
	if v158 == int32(0) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v163 = F_makeVarFromTargetEntry(m, int32(-2), v158)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	v165 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v163)+40)) = uint16(v165)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+36)) = v165
	v290 = v163
	goto L1
L60:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_fix_expr_common(m, v254, v19)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L23
	} else {
		goto L82
	}
L61:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v192 != int32(24) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+9)))
	if v173 != int32(1) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v176 == int32(7) {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v180 = F_tlist_member(m, v19, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	if v180 == int32(0) {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v185 = F_makeVarFromTargetEntry(m, int32(-1), v180)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v187 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v185)+40)) = uint16(v187)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+36)) = v187
	v290 = v185
	goto L1
L68:
	;
	if v192 != int32(8) {
		goto L60
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v201 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v203 = int32(0)
	v206 = v203
	v209 = v203
	v214 = float64(0)
	goto L73
L71:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v198 = F_fix_param_node(m, v197, v19)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L23
	} else {
		goto L72
	}
L72:
	;
	v290 = v198
	goto L1
L73:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v209<<(uint(int32(2))%32))))
	v223 = *(*float64)(unsafe.Add(mBase, uint32(v222)+56))
	v224 = *(*float64)(unsafe.Add(mBase, uint32(v222)+64))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v202)+360))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v225+v226-v228))) = uint8(v228)
	v233 = base.F64_add(v223, base.F64_mul(v201, v224))
	v235 = int32(0)
	v239 = base.B2i32(base.F64_ge(v214, v233) == v235) & base.B2i32(v206 != v235)
	if v239 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v202)+364))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v241)+16))
	v249 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v246+v247-v249))) = uint8(v249)
	v275 = v241
	goto L7
L75:
	;
	v240 = v214
	goto L77
L76:
	;
	v240 = v233
	goto L77
L77:
	;
	if v239 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v241 = v206
	goto L80
L79:
	;
	v241 = v222
	goto L80
L80:
	;
	v243 = v209 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v243 < v244 {
		v206 = v241
		v209 = v243
		v214 = v240
		goto L73
	} else {
		goto L81
	}
L81:
	;
	goto L74
L82:
	;
	v258 = F_expression_tree_mutator_impl(m, v19, int32(835), l1)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L23
	} else {
		goto L83
	}
L83:
	;
	v290 = v258
	goto L1
L84:
	;
	F_errmsg_internal(m, int32(75639), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L23
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(494232), int32(3145), int32(208987))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L23
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
	goto L6
}
func F_get_join_index_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	v18 = m.G0
	v20 = v18 - int32(144)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v23 = F_list_member(m, v22, l7)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = F__emscripten_memset_bulkmem(m, v20+int32(12), base.I32_extend8_s(int32(0)), int32(132))
	mBase = m.M
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v20 + int32(144)
	return
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if int32(0) < v33 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = int32(4)
	v43 = v20 + int32(16)
	v55 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_get_index_paths(m, l0, l1, l2, v20+int32(12), l6)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L65
	}
L10:
	;
	v62 = v55 << (uint(int32(2)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l4+v36+v62)))
	if v64 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l5+v36))))
	if v176 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v68 <= v67 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v71 = v62 + v43
	v77 = v67
	goto L15
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v77<<(uint(int32(2))%32))))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+28))
	v96 = int32(0)
	if v95 == v96 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L12
L17:
	;
	if v149 != 0 {
		goto L31
	} else {
		goto L32
	}
L18:
	;
	v149 = int32(1)
	goto L17
L19:
	;
	goto L20
L20:
	;
	if l7 == int32(0) {
		v140 = v96
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v149 = v140
	goto L17
L22:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v106 < v105 {
		v140 = v96
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v108 = int32(1)
	if v105 <= v108 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v111 = v108
	goto L26
L25:
	;
	v111 = v105
	goto L26
L26:
	;
	v112 = int32(8)
	v117 = int32(0)
	goto L27
L27:
	;
	v124 = v117 << (uint(int32(2)) % 32)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v95+v112+v124)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l7+v112))))
	v131 = v126 & (v128 ^ int32(-1))
	v133 = base.B2i32(v131 == int32(0))
	if v131 != 0 {
		v140 = v133
		goto L21
	} else {
		goto L29
	}
L28:
	;
	v140 = v133
	goto L21
L29:
	;
	v135 = v117 + int32(1)
	if v135 != v111 {
		v117 = v135
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v151 = F_lappend(m, v150, v93)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v155 = v77 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v155 < v156 {
		v77 = v155
		goto L15
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v151
	goto L33
L35:
	;
	goto L16
L36:
	;
	v289 = v62 + v43
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v62+(l3+v36))))
	v293 = F_list_concat(m, v290, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L60
	}
L37:
	;
	v179 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v180 <= v179 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v188 = v179
	goto L39
L39:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200+v188<<(uint(int32(2))%32))))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+28))
	v207 = int32(0)
	if v206 == v207 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v267 = v62 + v43
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v269 = F_lappend(m, v268, v204)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L59
	}
L41:
	;
	if v260 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L42:
	;
	v260 = int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if l7 == int32(0) {
		v251 = v207
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v260 = v251
	goto L41
L46:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v217 < v216 {
		v251 = v207
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v219 = int32(1)
	if v216 <= v219 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v222 = v219
	goto L50
L49:
	;
	v222 = v216
	goto L50
L50:
	;
	v223 = int32(8)
	v228 = int32(0)
	goto L51
L51:
	;
	v235 = v228 << (uint(int32(2)) % 32)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v206+v223+v235)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+(l7+v223))))
	v242 = v237 & (v239 ^ int32(-1))
	v244 = base.B2i32(v242 == int32(0))
	if v242 != 0 {
		v251 = v244
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v251 = v244
	goto L45
L53:
	;
	v246 = v228 + int32(1)
	if v246 != v222 {
		v228 = v246
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v264 = v188 + int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v264 < v265 {
		v188 = v264
		goto L39
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L40
L58:
	;
	goto L36
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v269
	goto L36
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v293
	if v293 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+12)) = uint8(v296)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v299 = v55 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v299 < v300 {
		v55 = v299
		goto L10
	} else {
		goto L64
	}
L64:
	;
	goto L11
L65:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	v324 = F_lappend(m, v323, l7)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v324
	goto L5
}
func F_join_clause_is_movable_into(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	v4 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v5 == v4 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v158
L2:
	;
	if v59 == int32(0) {
		v158 = v4
		goto L1
	} else {
		goto L16
	}
L3:
	;
	v59 = int32(1)
	goto L2
L4:
	;
	goto L5
L5:
	;
	if l2 == int32(0) {
		v50 = v4
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v59 = v50
	goto L2
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v16 < v15 {
		v50 = v4
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v18 = int32(1)
	if v15 <= v18 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = v18
	goto L11
L10:
	;
	v21 = v15
	goto L11
L11:
	;
	v22 = int32(8)
	v27 = int32(0)
	goto L12
L12:
	;
	v34 = v27 << (uint(int32(2)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v5+v22+v34)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l2+v22))))
	v41 = v36 & (v38 ^ int32(-1))
	v43 = base.B2i32(v41 == int32(0))
	if v41 != 0 {
		v50 = v43
		goto L6
	} else {
		goto L14
	}
L13:
	;
	v50 = v43
	goto L6
L14:
	;
	v45 = v27 + int32(1)
	if v45 != v21 {
		v27 = v45
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v63 = int32(0)
	if l1 == v63 {
		v104 = v63
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v104 == int32(0) {
		v158 = v4
		goto L1
	} else {
		goto L31
	}
L18:
	;
	goto L17
L19:
	;
	if v62 == int32(0) {
		v104 = v63
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v72 < v73 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v75 = v72
	goto L23
L22:
	;
	v75 = v73
	goto L23
L23:
	;
	if v75 <= int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v78 = int32(1)
	goto L26
L25:
	;
	v78 = v75
	goto L26
L26:
	;
	v79 = int32(8)
	v84 = int32(0)
	goto L27
L27:
	;
	v91 = v84 << (uint(int32(2)) % 32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62+v79+v91)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+(l1+v79))))
	v96 = v93 & v95
	v98 = base.B2i32(v96 != int32(0))
	if v96 != 0 {
		v104 = v98
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v104 = v98
	goto L18
L29:
	;
	v100 = v84 + int32(1)
	if v100 != v78 {
		v84 = v100
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v111 = int32(0)
	if l1 == v111 {
		v152 = v111
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v158 = v152 ^ int32(1)
	goto L1
L33:
	;
	goto L32
L34:
	;
	if v110 == int32(0) {
		v152 = v111
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v120 < v121 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v123 = v120
	goto L38
L37:
	;
	v123 = v121
	goto L38
L38:
	;
	if v123 <= int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v126 = int32(1)
	goto L41
L40:
	;
	v126 = v123
	goto L41
L41:
	;
	v127 = int32(8)
	v132 = int32(0)
	goto L42
L42:
	;
	v139 = v132 << (uint(int32(2)) % 32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v110+v127+v139)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+(l1+v127))))
	v144 = v141 & v143
	v146 = base.B2i32(v144 != int32(0))
	if v144 != 0 {
		v152 = v146
		goto L33
	} else {
		goto L44
	}
L43:
	;
	v152 = v146
	goto L33
L44:
	;
	v148 = v132 + int32(1)
	if v148 != v126 {
		v132 = v148
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
}
func F_join_path_components(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L7
L2:
	;
	goto L3
L3:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v126 != 0 {
		goto L36
	} else {
		goto L37
	}
L4:
	;
	goto L3
L5:
	;
	v123 = F_strlen(m, v112)
	mBase = m.M
	goto L4
L7:
	;
	goto L8
L8:
	;
	v17 = int32(1023)
	if (l0^l1)&int32(3) != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v116)
	goto L5
L10:
	;
	v97 = v92
	v98 = v93
	v99 = v94
	goto L32
L11:
	;
	if v87 == int32(0) {
		v112 = v85
		v113 = v86
		goto L9
	} else {
		goto L31
	}
L12:
	;
	v85 = l1
	v86 = l0
	v87 = v17
	goto L11
L13:
	;
	goto L14
L14:
	;
	if l1&int32(3) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v54 == int32(0) {
		v112 = v51
		v113 = v52
		goto L9
	} else {
		goto L24
	}
L16:
	;
	v51 = l1
	v52 = l0
	v53 = v17
	v54 = int32(1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v30 = l1
	v31 = l0
	v32 = v17
	goto L19
L19:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
	if v34 == int32(0) {
		v92 = v30
		v93 = v31
		v94 = v32
		goto L10
	} else {
		goto L21
	}
L20:
	;
	v51 = v45
	v52 = v39
	v53 = v41
	v54 = v43
	goto L15
L21:
	;
	v38 = int32(1)
	v39 = v31 + v38
	v41 = v32 - v38
	v42 = int32(0)
	v43 = base.B2i32(v41 != v42)
	v45 = v30 + v38
	if v45&int32(3) == v42 {
		v51 = v45
		v52 = v39
		v53 = v41
		v54 = v43
		goto L15
	} else {
		goto L22
	}
L22:
	;
	if v41 != 0 {
		v30 = v45
		v31 = v39
		v32 = v41
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v57 == int32(0) {
		v85 = v51
		v86 = v52
		v87 = v53
		goto L11
	} else {
		goto L25
	}
L25:
	;
	if base.Ui32(v53) < base.Ui32(int32(4)) {
		v85 = v51
		v86 = v52
		v87 = v53
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v63 = v51
	v64 = v52
	v65 = v53
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v71 = int32(-2139062144)
	if (int32(16843008)-v68|v68)&v71 != v71 {
		v92 = v63
		v93 = v64
		v94 = v65
		goto L10
	} else {
		goto L29
	}
L28:
	;
	v85 = v79
	v86 = v77
	v87 = v81
	goto L11
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v68
	v76 = int32(4)
	v77 = v64 + v76
	v79 = v63 + v76
	v81 = v65 - v76
	if base.Ui32(int32(3)) < base.Ui32(v81) {
		v63 = v79
		v64 = v77
		v65 = v81
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v92 = v85
	v93 = v86
	v94 = v87
	goto L10
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v101)
	if v101 == int32(0) {
		v112 = v97
		v113 = v98
		goto L9
	} else {
		goto L34
	}
L33:
	;
	v112 = v108
	v113 = v106
	goto L9
L34:
	;
	v105 = int32(1)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if v110 != 0 {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v127 = F_strlen(m, l0)
	mBase = m.M
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l2
	if v128 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	m.G0 = v8 + int32(16)
	return
L39:
	;
	v132 = int32(570922)
	goto L41
L40:
	;
	v132 = int32(757461)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v132
	v138 = F_pg_snprintf(m, l0+v127, int32(1024)-v127, int32(175835), v8)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	return
L43:
	;
	goto L38
}
