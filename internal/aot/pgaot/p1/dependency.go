package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_makeDependencyGraphWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	v3 = int32(0)
	if l0 == v3 {
		v221 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+8)))
	if v229 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L2:
	;
	return v221
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 - int32(110) {
	case 0:
		v221 = v3
		goto L2
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26:
		goto L4
	case 27:
		goto L8
	case 28:
		goto L7
	case 29:
		goto L6
	case 30:
		goto L5
	case 31:
		goto L9
	default:
		goto L10
	}
L4:
	;
	v215 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l1)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L57
	} else {
		goto L65
	}
L5:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v210 == int32(0) {
		goto L4
	} else {
		goto L64
	}
L6:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v207 == int32(0) {
		goto L4
	} else {
		goto L63
	}
L7:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v204 == int32(0) {
		goto L4
	} else {
		goto L62
	}
L8:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v201 == int32(0) {
		goto L4
	} else {
		goto L61
	}
L9:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v198 == int32(0) {
		goto L4
	} else {
		goto L60
	}
L10:
	;
	if v13 != int32(3) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != 0 {
		v221 = v3
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v19 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v125 <= int32(0) {
		v221 = v3
		goto L2
	} else {
		goto L41
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v25 = int32(0)
	if v25 < v22 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v28 = v22
	goto L18
L17:
	;
	v28 = v25
	goto L18
L18:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v35 = v3
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29+v35<<(uint(int32(2))%32))))
	if v43 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L13
L21:
	;
	v113 = v35 + int32(1)
	if v113 != v28 {
		v35 = v113
		goto L19
	} else {
		goto L40
	}
L22:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v46 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v49 = int32(0)
	if v49 < v46 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v52 = v46
	goto L26
L25:
	;
	v52 = v49
	goto L26
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v58 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54+v58<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.B2i32(v73 == int32(0))|base.B2i32(v73 != v76) != 0 {
		v94 = v73
		v95 = v76
		goto L30
	} else {
		goto L31
	}
L28:
	;
	return int32(0)
L29:
	;
	if v94-v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v79 = v53
	v80 = v70
	goto L32
L32:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v84 == int32(0) {
		v94 = v84
		v95 = v83
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v94 = v84
	v95 = v83
	goto L30
L34:
	;
	v87 = int32(1)
	if v84 == v83 {
		v79 = v79 + v87
		v80 = v80 + v87
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v98 = v58 + int32(1)
	if v52 != v98 {
		v58 = v98
		goto L27
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L28
L39:
	;
	goto L21
L40:
	;
	goto L20
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v131 = int32(0)
	goto L42
L42:
	;
	v143 = v129 + v131*int32(12)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if base.B2i32(v148 == int32(0))|base.B2i32(v148 != v151) != 0 {
		v169 = v148
		v170 = v151
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v221 = int32(0)
	goto L2
L44:
	;
	if v169-v170 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	v154 = v128
	v155 = v145
	goto L47
L47:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+1)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	if v159 == int32(0) {
		v169 = v159
		v170 = v158
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v169 = v159
	v170 = v158
	goto L45
L49:
	;
	v162 = int32(1)
	if v159 == v158 {
		v154 = v154 + v162
		v155 = v155 + v162
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v131 != v174 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v196 = v131 + int32(1)
	if v196 != v125 {
		v131 = v196
		goto L42
	} else {
		goto L59
	}
L54:
	;
	v177 = v174 * int32(12)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v129+v177)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v181 = F_bms_add_member(m, v179, v180)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+32)) = uint8(v190)
	return int32(0)
L57:
	;
	return int32(0)
L58:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v185+v177)+8)) = v181
	return int32(0)
L59:
	;
	goto L43
L60:
	;
	v228 = v198
	goto L1
L61:
	;
	v228 = v201
	goto L1
L62:
	;
	v228 = v204
	goto L1
L63:
	;
	v228 = v207
	goto L1
L64:
	;
	v228 = v210
	goto L1
L65:
	;
	v221 = v215
	goto L2
L66:
	;
	return int32(0)
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v234 = F_lcons(m, v232, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L57
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v284 = F_lcons(m, int32(0), v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L57
	} else {
		goto L80
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v234
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v237 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v276 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l1)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L57
	} else {
		goto L78
	}
L72:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v240 <= int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v246 = v3
	goto L74
L74:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253+v246<<(uint(int32(2))%32))))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+16))
	v259 = F_makeDependencyGraphWalker(m, v258, l1)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L57
	} else {
		goto L76
	}
L75:
	;
	goto L71
L76:
	;
	v262 = v246 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v262 < v263 {
		v246 = v262
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v279 = F_list_delete_first(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L57
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v279
	goto L66
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v284
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v287 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v334 = F_raw_expression_tree_walker_impl(m, l0, int32(484), l1)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L57
	} else {
		goto L92
	}
L82:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v290 <= int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v296 = v3
	goto L84
L84:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v287)+12))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v296<<(uint(int32(2))%32))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+16))
	v309 = F_makeDependencyGraphWalker(m, v308, l1)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L57
	} else {
		goto L86
	}
L85:
	;
	goto L81
L86:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v311 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	v314 = v312
	goto L89
L88:
	;
	v314 = int32(0)
	goto L89
L89:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v316 = F_lappend(m, v315, v307)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L57
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314))) = v316
	v320 = v296 + int32(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v320 < v321 {
		v296 = v320
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v337 = F_list_delete_first(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L57
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v337
	goto L66
}
