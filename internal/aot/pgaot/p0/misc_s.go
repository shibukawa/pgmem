package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SB_MatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	v6 = int32(0)
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v15 = v11 ^ int32(1)
	goto L3
L2:
	;
	v15 = int32(0)
	goto L3
L3:
	;
	if l3 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v18 != int32(37) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	return int32(1)
L7:
	;
	return int32(0)
L8:
	;
	v29 = int32(0)
	if base.B2i32(l1 <= v29)|base.B2i32(l3 <= v29) != 0 {
		v349 = l2
		v350 = l3
		v352 = base.B2i32(int32(0) < l1)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v388
L10:
	;
	if v352 != 0 {
		v388 = v6
		goto L9
	} else {
		goto L128
	}
L11:
	;
	v34 = l0
	v35 = l1
	v36 = l2
	v37 = l3
	goto L12
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v44 == int32(95) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v349 = v338
	v350 = v336
	v352 = v334
	goto L10
L14:
	;
	v333 = int32(1)
	v334 = base.B2i32(v333 < v35)
	v336 = v332 - v333
	v338 = v331 + v333
	if v35 < int32(2) {
		v349 = v338
		v350 = v336
		v352 = v334
		goto L10
	} else {
		goto L126
	}
L15:
	;
	v331 = v36
	v332 = v37
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v44 == int32(37) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v324 = F_pg_strncoll(m, v185, v219, v34, v35, l4)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L7
	} else {
		goto L121
	}
L19:
	;
	if v44 == int32(92) {
		goto L109
	} else {
		goto L110
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L104
	}
L21:
	;
	if base.Ui32(v37) < base.Ui32(int32(2)) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	if v15&int32(1) == int32(0) {
		goto L19
	} else {
		goto L51
	}
L24:
	;
	return int32(1)
L25:
	;
	goto L26
L26:
	;
	v53 = v34
	v54 = v35
	v55 = v36
	v56 = v37
	goto L28
L27:
	;
	if v54 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	v63 = int32(1)
	v64 = v56 - v63
	v66 = v55 + v63
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	switch v67 - int32(92) {
	case 0:
		goto L30
	case 1, 2:
		v88 = v67
		goto L27
	case 3:
		goto L32
	default:
		goto L33
	}
L29:
	;
	if v64 == int32(1) {
		goto L20
	} else {
		goto L39
	}
L30:
	;
	goto L29
L31:
	;
	if base.Ui32(int32(2)) < base.Ui32(v56) {
		v53 = v80
		v54 = v81
		v55 = v66
		v56 = v64
		goto L28
	} else {
		goto L38
	}
L32:
	;
	if v54 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v67 == int32(37) {
		v80 = v53
		v81 = v54
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v88 = v67
	goto L27
L35:
	;
	return int32(-1)
L36:
	;
	goto L37
L37:
	;
	v76 = int32(1)
	v80 = v53 + v76
	v81 = v54 - v76
	goto L31
L38:
	;
	v388 = int32(1)
	goto L9
L39:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v88 = v87
	goto L27
L40:
	;
	return int32(-1)
L41:
	;
	goto L42
L42:
	;
	v95 = v53
	v96 = v54
	goto L43
L43:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if (base.B2i32(v105 == v88&int32(255))|v15)&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v388 = int32(-1)
	goto L9
L45:
	;
	v110 = F_SB_MatchText(m, v95, v96, v66, v64, l4)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v114 = int32(1)
	if v114 < v96 {
		v95 = v95 + v114
		v96 = v96 - v114
		goto L43
	} else {
		goto L50
	}
L48:
	;
	if v110 != 0 {
		v388 = v110
		goto L9
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L44
L51:
	;
	v128 = v37
	v130 = v44
	v131 = v36
	v134 = int32(0)
	goto L55
L52:
	;
	v270 = F_pg_strncoll(m, v36, v168-v36, v34, v35, l4)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L103
	}
L53:
	;
	v231 = v35
	v235 = v34
	goto L85
L54:
	;
	v185 = F_palloc(m, v182-v36)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L77
	}
L55:
	;
	v136 = v130 & int32(255)
	switch v136 - int32(92) {
	case 0:
		goto L62
	case 1, 2:
		goto L59
	case 3:
		goto L60
	default:
		goto L61
	}
L56:
	;
	if v134 == int32(0) {
		goto L52
	} else {
		goto L76
	}
L57:
	;
	goto L56
L58:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v128 = v173
	v130 = v176
	v131 = v174
	v134 = v175
	goto L55
L59:
	;
	v167 = int32(1)
	v168 = v131 + v167
	v170 = v128 - v167
	if v170 == int32(0) {
		goto L57
	} else {
		goto L75
	}
L60:
	;
	if v134 != 0 {
		goto L72
	} else {
		goto L73
	}
L61:
	;
	if v136 != int32(37) {
		goto L59
	} else {
		goto L71
	}
L62:
	;
	if v128 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v157 = int32(2)
	v158 = v131 + v157
	v159 = int32(1)
	v161 = v128 - v157
	if v161 != 0 {
		v173 = v161
		v174 = v158
		v175 = v159
		goto L58
	} else {
		goto L70
	}
L66:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(236), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v181 = int32(0)
	v182 = v158
	v183 = v159
	goto L54
L71:
	;
	goto L60
L72:
	;
	v181 = v128
	v182 = v131
	v183 = int32(0)
	goto L54
L73:
	;
	goto L74
L74:
	;
	v222 = v36
	v223 = v128
	v226 = v131
	v227 = v131 - v36
	v228 = v6
	goto L53
L75:
	;
	v173 = v170
	v174 = v168
	v175 = v134
	goto L58
L76:
	;
	v181 = int32(0)
	v182 = v168
	v183 = int32(1)
	goto L54
L77:
	;
	if base.Ui32(v36) < base.Ui32(v182) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v190 = v36
	v193 = v185
	goto L81
L79:
	;
	v214 = v185
	goto L80
L80:
	;
	v219 = v214 - v185
	if v183 != 0 {
		goto L18
	} else {
		goto L84
	}
L81:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	v201 = v190 + base.B2i32(v198 == int32(92))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v202)
	v204 = int32(1)
	v205 = v193 + v204
	v207 = v201 + v204
	if base.Ui32(v207) < base.Ui32(v182) {
		v190 = v207
		v193 = v205
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v214 = v205
	goto L80
L83:
	;
	goto L82
L84:
	;
	v222 = v185
	v223 = v181
	v226 = v182
	v227 = v219
	v228 = v185
	goto L53
L85:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_SB_MatchText[0]))
	if v241 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v245 = F_pg_strncoll(m, v222, v227, v34, v235-v34, l4)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	if v231 != 0 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	if v245 != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v247 = F_SB_MatchText(m, v235, v231, v226, v223, l4)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	if v247 != int32(1) {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	if v228 == int32(0) {
		v388 = int32(1)
		goto L9
	} else {
		goto L96
	}
L96:
	;
	F_pfree(m, v228)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	return int32(1)
L98:
	;
	v258 = int32(1)
	v231 = v231 - v258
	v235 = v235 + v258
	goto L85
L99:
	;
	v262 = int32(0)
	if v228 == v262 {
		v388 = v262
		goto L9
	} else {
		goto L101
	}
L101:
	;
	F_pfree(m, v228)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	return int32(0)
L103:
	;
	return base.B2i32(v270 == int32(0))
L104:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(168), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L117
	}
L109:
	;
	if base.Ui32(v37) <= base.Ui32(int32(1)) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v44 == v304 {
		v331 = v36
		v332 = v37
		goto L14
	} else {
		goto L116
	}
L112:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v295 == v296 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v298 = int32(1)
	v331 = v36 + v298
	v332 = v37 - v298
	goto L14
L114:
	;
	goto L115
L115:
	;
	return int32(0)
L116:
	;
	return int32(0)
L117:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(_a_F_SB_MatchText_0), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_SB_MatchText_1), int32(356), int32(_a_F_SB_MatchText_2))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	if v185 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	F_pfree(m, v185)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	return base.B2i32(v324 == int32(0))
L125:
	;
	goto L124
L126:
	;
	v341 = int32(1)
	if v341 < v332 {
		v34 = v34 + v341
		v35 = v35 - v341
		v36 = v338
		v37 = v336
		goto L12
	} else {
		goto L127
	}
L127:
	;
	goto L13
L128:
	;
	v357 = int32(1)
	if v350 <= int32(0) {
		v388 = v357
		goto L9
	} else {
		goto L129
	}
L129:
	;
	v362 = v349
	v363 = v350
	goto L130
L130:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v370 != int32(37) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v388 = v357
	goto L9
L132:
	;
	return int32(-1)
L133:
	;
	goto L134
L134:
	;
	v375 = int32(1)
	if v375 < v363 {
		v362 = v362 + v375
		v363 = v363 - v375
		goto L130
	} else {
		goto L135
	}
L135:
	;
	goto L131
}
func F_ScanKeyEntryInitializeWithInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v23 int64
	_ = v23
	var v25 int64
	_ = v25
	v3 = l2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v11 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v11)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v16 = l0 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_ScanKeyEntryInitializeWithInfo[0]))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l5)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l5)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v11
	return
}
func F_ScanKeyInit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	v2 = l1
	v3 = l2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(4080218931200)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	F_fmgr_info(m, l3, l0+int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		return
	}
}
func F_SerializeSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v14)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)) = uint8(v10)
	if v10&int32(1) != 0 {
		v25 = v9
	} else {
		v25 = v3
	}
	if v14 != 0 {
		v26 = v25
	} else {
		v26 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v28 == int32(0) {
	} else {
		v32 = v28 << (uint(int32(2)) % 32)
		if v32 == int32(0) {
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			base.MemoryCopy(m, l1+int32(24), v37, v32)
		}
	}
	if v26 <= int32(0) {
	} else {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v44 = v42 << (uint(int32(2)) % 32)
		if v44 == int32(0) {
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			base.MemoryCopy(m, l1+v47<<(uint(int32(2))%32)+int32(24), v53, v44)
		}
	}
	return
}
func F_SetMatViewPopulatedState(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
		v17 = F_SearchSysCacheCopy(m, int32(57), v15, int64(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
					F_errmsg_internal(m, int32(_a_F_SetMatViewPopulatedState_0), v8)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetMatViewPopulatedState_1), int32(95), int32(_a_F_SetMatViewPopulatedState_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v35+v36)+129)) = uint8(v2)
				F_CatalogTupleUpdate(m, v12, v17+int32(4), v17)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					F_pfree(m, v17)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_relation_close(m, v12, int32(3))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_CommandCounterIncrement(m)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_SignalHandlerForShutdownRequest(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_SignalHandlerForShutdownRequest_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_SignalHandlerForShutdownRequest_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_SignalHandlerForShutdownRequest[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_SwitchToSharedLatch(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[0]))
	v7 = v5 + int32(316)
	*(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[1])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[2]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(1)
	F_ModifyWaitEvent(m, v10, v11, v11, v7)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v17 = v7
	goto L3
L3:
	;
	v18 = int32(0)
	v21 = base.AtomicRmwOr32(m, v18, int32(_a_F_SwitchToSharedLatch_0), v18)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[1]))
	v17 = v16
	goto L3
L6:
	;
	return
L7:
	;
	goto L6
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1)
	v25 = int32(0)
	v28 = base.AtomicRmwOr32(m, v25, int32(_a_F_SwitchToSharedLatch_0), v25)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v29 == v25 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v32 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[3]))
	if v36 == v32 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = m.G0
	v40 = v38 - int32(16)
	m.G0 = v40
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[4]))
	if v43 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v66 = F_pgmem_kill(m, v32, int32(23))
	mBase = m.M
	goto L7
L14:
	;
	m.G0 = v40 + int32(16)
	goto L6
L15:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+15)) = uint8(v46)
	goto L16
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[5]))
	v54 = F_write(m, v50, v40+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v54 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	goto L14
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToSharedLatch[6]))
	if v58 == int32(27) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
}
func F_SwitchToUntrustedUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1]))
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)))) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_member_can_set_role(m, v18, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		if v19 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v22 = F_member_can_set_role(m, l0, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v22 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1])) = v24
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0])) = l0
					v43 = int32(-1)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[1])) = v24 | int32(2)
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[0])) = l0
					v37 = int32(_a_F_SwitchToUntrustedUser_0)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[2]))
					v41 = v39 + int32(1)
					*(*int32)(unsafe.Add(mBase, _c_F_SwitchToUntrustedUser[2])) = v41
					v43 = v41
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v43
				m.G0 = v8 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v57 = F_GetUserNameFromId(m, v55, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v60 = F_GetUserNameFromId(m, l0, int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v60
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v57
							F_errmsg(m, int32(_a_F_SwitchToUntrustedUser_1), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_SwitchToUntrustedUser_2), int32(45), int32(_a_F_SwitchToUntrustedUser_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
	}
}
func F_sanitize_char_1(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	Fn14374(m, l0, int32(_a_F_sanitize_char_1_0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_scalargesel(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = F_scalarineqsel_wrapper(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_scanNSItemForColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
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
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_scanRTEForColumn(m, l0, v15, v16, l3, l4, v6, v6)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 != 0 {
			v24 = base.B2i32(v19 == int32(-6))
			v25 = int32(0)
			v26 = base.B2i32(v25 <= v19)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			if base.B2i32(v24|v26 == v25)&base.B2i32(v30 == int32(28)) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v274 = m.ExcPending
				if v274 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
					mBase = m.M
					v277 = m.ExcPending
					if v277 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = l3
						F_errmsg(m, int32(_a_F_scanNSItemForColumn_1), v13)
						mBase = m.M
						v281 = m.ExcPending
						if v281 != 0 {
							return int32(0)
						} else {
							F_parser_errposition(m, l0, l4)
							mBase = m.M
							v283 = m.ExcPending
							if v283 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(741), int32(_a_F_scanNSItemForColumn_3))
								mBase = m.M
								v288 = m.ExcPending
								if v288 != 0 {
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
			} else {
				if base.B2i32(v24|v26 == int32(0))&base.B2i32(v30 == int32(43)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v292 = m.ExcPending
					if v292 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l3
							F_errmsg(m, int32(_a_F_scanNSItemForColumn_4), v11+int32(-48))
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								F_parser_errposition(m, l0, l4)
								mBase = m.M
								v303 = m.ExcPending
								if v303 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(754), int32(_a_F_scanNSItemForColumn_3))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
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
				} else {
					if base.B2i32(v24|v26 == int32(0))&base.B2i32(v30 == int32(18)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v312 = m.ExcPending
						if v312 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(_a_F_scanNSItemForColumn_0))
							mBase = m.M
							v315 = m.ExcPending
							if v315 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l3
								F_errmsg(m, int32(_a_F_scanNSItemForColumn_5), v11+int32(-32))
								mBase = m.M
								v321 = m.ExcPending
								if v321 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, l0, l4)
									mBase = m.M
									v323 = m.ExcPending
									if v323 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(765), int32(_a_F_scanNSItemForColumn_3))
										mBase = m.M
										v328 = m.ExcPending
										if v328 != 0 {
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
					} else {
						if int32(0) < v19 {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
							v51 = v48 + v19<<(uint(int32(5))%32)
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(32))))
							if v54 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v332 = m.ExcPending
								if v332 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50360452))
									mBase = m.M
									v335 = m.ExcPending
									if v335 != 0 {
										return int32(0)
									} else {
										v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v337
										*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l3
										F_errmsg(m, int32(_a_F_scanNSItemForColumn_6), v11+int32(-16))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_scanNSItemForColumn_2), int32(779), int32(_a_F_scanNSItemForColumn_3))
											mBase = m.M
											v349 = m.ExcPending
											if v349 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51-int32(28)))))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(24))))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(20))))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(16))))
								v69 = F_makeVar(m, v54, v59, v62, v65, v68, l2)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v51-int32(8))))
									*(*int32)(unsafe.Add(mBase, uint32(v69)+36)) = v73
									v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51-int32(4)))))
									*(*uint16)(unsafe.Add(mBase, uint32(v69)+40)) = uint16(v77)
									v89 = v69
									*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = l4
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v92
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									if v95 == int32(0) {
										v156 = l0
									} else {
										v99 = v95 & int32(7)
										if base.Ui32(int32(8)) <= base.Ui32(v95) {
											v106 = int32(0)
											v108 = l0
											for {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
												v124 = v106 + int32(8)
												if v124 != v95&int32(-8) {
													v106 = v124
													v108 = v122
													continue
												} else {
													break
												}
												break
											}
											if v99 == int32(0) {
												v156 = v122
											} else {
												v131 = v122
												v140 = int32(0)
												v142 = v131
												for {
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													v151 = v140 + int32(1)
													if v151 != v99 {
														v140 = v151
														v142 = v149
														continue
													} else {
														break
													}
													break
												}
												v156 = v149
											}
										} else {
											v131 = l0
											v140 = int32(0)
											v142 = v131
											for {
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												v151 = v140 + int32(1)
												if v151 != v99 {
													v140 = v151
													v142 = v149
													continue
												} else {
													break
												}
												break
											}
											v156 = v149
										}
									}
									if v94 <= int32(0) {
										v185 = v95
										if v185 == int32(0) {
											v243 = l0
										} else {
											v189 = v185 & int32(7)
											if base.Ui32(int32(8)) <= base.Ui32(v185) {
												v195 = l0
												v198 = int32(0)
												for {
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
													v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
													v214 = v198 + int32(8)
													if v214 != v185&int32(-8) {
														v195 = v212
														v198 = v214
														continue
													} else {
														break
													}
													break
												}
												if v189 == int32(0) {
													v243 = v212
												} else {
													v218 = v212
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											} else {
												v218 = l0
												v229 = v218
												v232 = int32(0)
												for {
													v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
													v241 = v232 + int32(1)
													if v241 != v189 {
														v229 = v239
														v232 = v241
														continue
													} else {
														break
													}
													break
												}
												v243 = v239
											}
										}
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
										v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
										F_markRTEForSelectPriv(m, v243, v253, v254)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											v262 = v89
											m.G0 = v13 - int32(-64)
											return v262
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
										if v165 == int32(0) {
											v185 = v95
											if v185 == int32(0) {
												v243 = l0
											} else {
												v189 = v185 & int32(7)
												if base.Ui32(int32(8)) <= base.Ui32(v185) {
													v195 = l0
													v198 = int32(0)
													for {
														v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
														v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
														v214 = v198 + int32(8)
														if v214 != v185&int32(-8) {
															v195 = v212
															v198 = v214
															continue
														} else {
															break
														}
														break
													}
													if v189 == int32(0) {
														v243 = v212
													} else {
														v218 = v212
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												} else {
													v218 = l0
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											}
											v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
											v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
											F_markRTEForSelectPriv(m, v243, v253, v254)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
												return int32(0)
											} else {
												v262 = v89
												m.G0 = v13 - int32(-64)
												return v262
											}
										} else {
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
											if v168 < v94 {
												v185 = v95
												if v185 == int32(0) {
													v243 = l0
												} else {
													v189 = v185 & int32(7)
													if base.Ui32(int32(8)) <= base.Ui32(v185) {
														v195 = l0
														v198 = int32(0)
														for {
															v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
															v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
															v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
															v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
															v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
															v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
															v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
															v214 = v198 + int32(8)
															if v214 != v185&int32(-8) {
																v195 = v212
																v198 = v214
																continue
															} else {
																break
															}
															break
														}
														if v189 == int32(0) {
															v243 = v212
														} else {
															v218 = v212
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													} else {
														v218 = l0
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												}
												v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
												F_markRTEForSelectPriv(m, v243, v253, v254)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													v262 = v89
													m.G0 = v13 - int32(-64)
													return v262
												}
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+v94<<(uint(int32(2))%32)-int32(4))))
												if v176 == int32(0) {
													v185 = v95
													if v185 == int32(0) {
														v243 = l0
													} else {
														v189 = v185 & int32(7)
														if base.Ui32(int32(8)) <= base.Ui32(v185) {
															v195 = l0
															v198 = int32(0)
															for {
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																v214 = v198 + int32(8)
																if v214 != v185&int32(-8) {
																	v195 = v212
																	v198 = v214
																	continue
																} else {
																	break
																}
																break
															}
															if v189 == int32(0) {
																v243 = v212
															} else {
																v218 = v212
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														} else {
															v218 = l0
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													}
													v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
													F_markRTEForSelectPriv(m, v243, v253, v254)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int32(0)
													} else {
														v262 = v89
														m.G0 = v13 - int32(-64)
														return v262
													}
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
													v180 = F_bms_union(m, v179, v176)
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v180
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
														v185 = v183
														if v185 == int32(0) {
															v243 = l0
														} else {
															v189 = v185 & int32(7)
															if base.Ui32(int32(8)) <= base.Ui32(v185) {
																v195 = l0
																v198 = int32(0)
																for {
																	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																	v214 = v198 + int32(8)
																	if v214 != v185&int32(-8) {
																		v195 = v212
																		v198 = v214
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v189 == int32(0) {
																	v243 = v212
																} else {
																	v218 = v212
																	v229 = v218
																	v232 = int32(0)
																	for {
																		v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																		v241 = v232 + int32(1)
																		if v241 != v189 {
																			v229 = v239
																			v232 = v241
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v243 = v239
																}
															} else {
																v218 = l0
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														}
														v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
														F_markRTEForSelectPriv(m, v243, v253, v254)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															v262 = v89
															m.G0 = v13 - int32(-64)
															return v262
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v79 = base.I32_extend16_s(v19)
							v80 = F_SystemAttributeDefinition(m, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+68))
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+76))
								v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+96))
								v86 = F_makeVar(m, v82, v79, v83, v84, v85, l2)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									v89 = v86
									*(*int32)(unsafe.Add(mBase, uint32(v89)+44)) = l4
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v92
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
									if v95 == int32(0) {
										v156 = l0
									} else {
										v99 = v95 & int32(7)
										if base.Ui32(int32(8)) <= base.Ui32(v95) {
											v106 = int32(0)
											v108 = l0
											for {
												v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
												v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
												v124 = v106 + int32(8)
												if v124 != v95&int32(-8) {
													v106 = v124
													v108 = v122
													continue
												} else {
													break
												}
												break
											}
											if v99 == int32(0) {
												v156 = v122
											} else {
												v131 = v122
												v140 = int32(0)
												v142 = v131
												for {
													v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
													v151 = v140 + int32(1)
													if v151 != v99 {
														v140 = v151
														v142 = v149
														continue
													} else {
														break
													}
													break
												}
												v156 = v149
											}
										} else {
											v131 = l0
											v140 = int32(0)
											v142 = v131
											for {
												v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
												v151 = v140 + int32(1)
												if v151 != v99 {
													v140 = v151
													v142 = v149
													continue
												} else {
													break
												}
												break
											}
											v156 = v149
										}
									}
									if v94 <= int32(0) {
										v185 = v95
										if v185 == int32(0) {
											v243 = l0
										} else {
											v189 = v185 & int32(7)
											if base.Ui32(int32(8)) <= base.Ui32(v185) {
												v195 = l0
												v198 = int32(0)
												for {
													v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
													v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
													v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
													v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
													v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
													v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
													v214 = v198 + int32(8)
													if v214 != v185&int32(-8) {
														v195 = v212
														v198 = v214
														continue
													} else {
														break
													}
													break
												}
												if v189 == int32(0) {
													v243 = v212
												} else {
													v218 = v212
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											} else {
												v218 = l0
												v229 = v218
												v232 = int32(0)
												for {
													v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
													v241 = v232 + int32(1)
													if v241 != v189 {
														v229 = v239
														v232 = v241
														continue
													} else {
														break
													}
													break
												}
												v243 = v239
											}
										}
										v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
										v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
										F_markRTEForSelectPriv(m, v243, v253, v254)
										mBase = m.M
										v256 = m.ExcPending
										if v256 != 0 {
											return int32(0)
										} else {
											v262 = v89
											m.G0 = v13 - int32(-64)
											return v262
										}
									} else {
										v165 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
										if v165 == int32(0) {
											v185 = v95
											if v185 == int32(0) {
												v243 = l0
											} else {
												v189 = v185 & int32(7)
												if base.Ui32(int32(8)) <= base.Ui32(v185) {
													v195 = l0
													v198 = int32(0)
													for {
														v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
														v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
														v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
														v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
														v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
														v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
														v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
														v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
														v214 = v198 + int32(8)
														if v214 != v185&int32(-8) {
															v195 = v212
															v198 = v214
															continue
														} else {
															break
														}
														break
													}
													if v189 == int32(0) {
														v243 = v212
													} else {
														v218 = v212
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												} else {
													v218 = l0
													v229 = v218
													v232 = int32(0)
													for {
														v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
														v241 = v232 + int32(1)
														if v241 != v189 {
															v229 = v239
															v232 = v241
															continue
														} else {
															break
														}
														break
													}
													v243 = v239
												}
											}
											v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
											v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
											F_markRTEForSelectPriv(m, v243, v253, v254)
											mBase = m.M
											v256 = m.ExcPending
											if v256 != 0 {
												return int32(0)
											} else {
												v262 = v89
												m.G0 = v13 - int32(-64)
												return v262
											}
										} else {
											v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
											if v168 < v94 {
												v185 = v95
												if v185 == int32(0) {
													v243 = l0
												} else {
													v189 = v185 & int32(7)
													if base.Ui32(int32(8)) <= base.Ui32(v185) {
														v195 = l0
														v198 = int32(0)
														for {
															v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
															v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
															v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
															v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
															v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
															v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
															v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
															v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
															v214 = v198 + int32(8)
															if v214 != v185&int32(-8) {
																v195 = v212
																v198 = v214
																continue
															} else {
																break
															}
															break
														}
														if v189 == int32(0) {
															v243 = v212
														} else {
															v218 = v212
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													} else {
														v218 = l0
														v229 = v218
														v232 = int32(0)
														for {
															v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
															v241 = v232 + int32(1)
															if v241 != v189 {
																v229 = v239
																v232 = v241
																continue
															} else {
																break
															}
															break
														}
														v243 = v239
													}
												}
												v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
												v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
												F_markRTEForSelectPriv(m, v243, v253, v254)
												mBase = m.M
												v256 = m.ExcPending
												if v256 != 0 {
													return int32(0)
												} else {
													v262 = v89
													m.G0 = v13 - int32(-64)
													return v262
												}
											} else {
												v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
												v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+v94<<(uint(int32(2))%32)-int32(4))))
												if v176 == int32(0) {
													v185 = v95
													if v185 == int32(0) {
														v243 = l0
													} else {
														v189 = v185 & int32(7)
														if base.Ui32(int32(8)) <= base.Ui32(v185) {
															v195 = l0
															v198 = int32(0)
															for {
																v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																v214 = v198 + int32(8)
																if v214 != v185&int32(-8) {
																	v195 = v212
																	v198 = v214
																	continue
																} else {
																	break
																}
																break
															}
															if v189 == int32(0) {
																v243 = v212
															} else {
																v218 = v212
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														} else {
															v218 = l0
															v229 = v218
															v232 = int32(0)
															for {
																v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																v241 = v232 + int32(1)
																if v241 != v189 {
																	v229 = v239
																	v232 = v241
																	continue
																} else {
																	break
																}
																break
															}
															v243 = v239
														}
													}
													v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
													v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
													F_markRTEForSelectPriv(m, v243, v253, v254)
													mBase = m.M
													v256 = m.ExcPending
													if v256 != 0 {
														return int32(0)
													} else {
														v262 = v89
														m.G0 = v13 - int32(-64)
														return v262
													}
												} else {
													v179 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
													v180 = F_bms_union(m, v179, v176)
													mBase = m.M
													v181 = m.ExcPending
													if v181 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v180
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
														v185 = v183
														if v185 == int32(0) {
															v243 = l0
														} else {
															v189 = v185 & int32(7)
															if base.Ui32(int32(8)) <= base.Ui32(v185) {
																v195 = l0
																v198 = int32(0)
																for {
																	v205 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
																	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
																	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
																	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
																	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
																	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
																	v214 = v198 + int32(8)
																	if v214 != v185&int32(-8) {
																		v195 = v212
																		v198 = v214
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v189 == int32(0) {
																	v243 = v212
																} else {
																	v218 = v212
																	v229 = v218
																	v232 = int32(0)
																	for {
																		v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																		v241 = v232 + int32(1)
																		if v241 != v189 {
																			v229 = v239
																			v232 = v241
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v243 = v239
																}
															} else {
																v218 = l0
																v229 = v218
																v232 = int32(0)
																for {
																	v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
																	v241 = v232 + int32(1)
																	if v241 != v189 {
																		v229 = v239
																		v232 = v241
																		continue
																	} else {
																		break
																	}
																	break
																}
																v243 = v239
															}
														}
														v253 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
														v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+8)))
														F_markRTEForSelectPriv(m, v243, v253, v254)
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															v262 = v89
															m.G0 = v13 - int32(-64)
															return v262
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
		} else {
			v262 = v6
			m.G0 = v13 - int32(-64)
			return v262
		}
	}
}
func F_scanner_finish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
	if base.Ui32(int32(_a_F_scanner_finish_0)) <= base.Ui32(v4) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		F_pfree(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v11 = v10
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			if int32(_a_F_scanner_finish_0) <= v12 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				F_pfree(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v11 = v3
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
		if int32(_a_F_scanner_finish_0) <= v12 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			F_pfree(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_scram_exchange(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int64
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v712 int32
	_ = v712
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
	var v730 int32
	_ = v730
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1230 int32
	_ = v1230
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1527 int32
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1557 int32
	_ = v1557
	var v1567 int32
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1666 int32
	_ = v1666
	var v1667 int64
	_ = v1667
	var v1669 int64
	_ = v1669
	var v1671 int64
	_ = v1671
	var v1673 int64
	_ = v1673
	var v1675 int64
	_ = v1675
	var v1677 int64
	_ = v1677
	var v1679 int64
	_ = v1679
	var v1681 int64
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1692 int32
	_ = v1692
	var v1708 int32
	_ = v1708
	var v1711 int32
	_ = v1711
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1817 int32
	_ = v1817
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1824 int32
	_ = v1824
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1867 int32
	_ = v1867
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1924 int32
	_ = v1924
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1943 int32
	_ = v1943
	var v1947 int32
	_ = v1947
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2046 int32
	_ = v2046
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2099 int32
	_ = v2099
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2119 int32
	_ = v2119
	var v2125 int32
	_ = v2125
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2138 int32
	_ = v2138
	var v2143 int32
	_ = v2143
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v7
	if l1 == v7 {
		goto L23
	} else {
		goto L24
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L26
	} else {
		goto L540
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L26
	} else {
		goto L524
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L26
	} else {
		goto L521
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L26
	} else {
		goto L505
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L26
	} else {
		goto L500
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L26
	} else {
		goto L495
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L26
	} else {
		goto L490
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L26
	} else {
		goto L486
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L26
	} else {
		goto L483
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L26
	} else {
		goto L479
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L26
	} else {
		goto L475
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L26
	} else {
		goto L471
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L26
	} else {
		goto L467
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L26
	} else {
		goto L463
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L26
	} else {
		goto L458
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L26
	} else {
		goto L452
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1794 = m.ExcPending
	if v1794 != 0 {
		goto L26
	} else {
		goto L447
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L26
	} else {
		goto L441
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L26
	} else {
		goto L436
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1728 = m.ExcPending
	if v1728 != 0 {
		goto L26
	} else {
		goto L431
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L26
	} else {
		goto L426
	}
L22:
	;
	m.G0 = v19 + int32(224)
	return v1692
L23:
	;
	v26 = F_pstrdup(m, int32(_a_F_scram_exchange_0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if l2 == int32(0) {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v26
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(0)
	v1692 = v7
	goto L22
L28:
	;
	v35 = F_strlen(m, l1)
	mBase = m.M
	if v35 != l2 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v37 {
	case 0:
		goto L34
	case 1:
		goto L33
	default:
		goto L32
	}
L30:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1657 != 0 {
		goto L421
	} else {
		goto L422
	}
L31:
	;
	v1633 = int32(0)
	v1634 = int32(2)
	if l5 == v1633 {
		v1643 = v1633
		v1648 = v1634
		goto L30
	} else {
		goto L419
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1623 = m.ExcPending
	if v1623 != 0 {
		goto L26
	} else {
		goto L416
	}
L33:
	;
	v544 = F_pstrdup(m, l1)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L26
	} else {
		goto L148
	}
L34:
	;
	v38 = F_pstrdup(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L26
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v38
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)) = uint8(v41)
	switch v41 - int32(110) {
	case 0:
		goto L42
	default:
		goto L37
	case 2:
		goto L40
	case 11:
		goto L41
	}
L36:
	;
	v229 = v98 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v229
	v231 = F_pstrdup(m, v229)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L26
	} else {
		goto L89
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L26
	} else {
		goto L83
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L26
	} else {
		goto L65
	}
L39:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v99 == int32(44) {
		goto L36
	} else {
		goto L57
	}
L40:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v61 == int32(0) {
		goto L15
	} else {
		goto L47
	}
L41:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v53 == int32(1) {
		goto L17
	} else {
		goto L45
	}
L42:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v45 == int32(1) {
		goto L19
	} else {
		goto L43
	}
L43:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v48 != int32(44) {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v98 = v38 + int32(2)
	goto L39
L45:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v56 != int32(44) {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v98 = v38 + int32(2)
	goto L39
L47:
	;
	v67 = F_read_attr_value(m, v19+int32(192), int32(112))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L26
	} else {
		goto L48
	}
L48:
	;
	v69 = int32(_a_F_scram_exchange_1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[0])))
	if base.B2i32(v72 == int32(0))|base.B2i32(v72 != v75) != 0 {
		v93 = v72
		v94 = v75
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v93-v94 != 0 {
		goto L38
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v78 = v67
	v79 = v69
	goto L52
L52:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	if v83 == int32(0) {
		v93 = v83
		v94 = v82
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v93 = v83
	v94 = v82
	goto L50
L54:
	;
	v86 = int32(1)
	if v83 == v82 {
		v78 = v78 + v86
		v79 = v79 + v86
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v98 = v96
	goto L39
L57:
	;
	if v99 == int32(97) {
		goto L14
	} else {
		goto L58
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L59
	}
L59:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98))))
	F_sanitize_char_2(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(_a_F_scram_exchange_3)
	v123 = F_errdetail(m, int32(_a_F_scram_exchange_4), v19+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1079), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
	;
	v139 = int32(0)
	goto L67
L67:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139+v67))))
	if v155 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187)+uint32(_c_F_scram_exchange[1]))) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = int32(_a_F_scram_exchange_7)
	F_errmsg(m, int32(_a_F_scram_exchange_8), v19+int32(80))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L26
	} else {
		goto L81
	}
L69:
	;
	goto L68
L70:
	;
	v187 = v139
	goto L69
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v155-int32(33))&int32(255)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v165 = int32(63)
	goto L75
L74:
	;
	v165 = v155
	goto L75
L75:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+uint32(_c_F_scram_exchange[1]))) = uint8(v165)
	v168 = v139 | int32(1)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v168))))
	if v170 == int32(0) {
		v187 = v168
		goto L69
	} else {
		goto L76
	}
L76:
	;
	if base.Ui32(int32(94)) <= base.Ui32((v170-int32(33))&int32(255)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v180 = int32(63)
	goto L79
L78:
	;
	v180 = v170
	goto L79
L79:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_scram_exchange[1]))) = uint8(v180)
	v182 = int32(30)
	v184 = v139 + int32(2)
	if v184 != v182 {
		v139 = v184
		goto L67
	} else {
		goto L80
	}
L80:
	;
	v187 = v182
	goto L69
L81:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1057), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L26
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L26
	} else {
		goto L85
	}
L85:
	;
	v215 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38))))
	F_sanitize_char_2(m, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L26
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_scram_exchange_3)
	v221 = F_errdetail(m, int32(_a_F_scram_exchange_9), v19)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L26
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1064), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L26
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v231
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v234 == int32(109) {
		goto L13
	} else {
		goto L90
	}
L90:
	;
	v238 = v19 + int32(192)
	v240 = F_read_attr_value(m, v238, int32(110))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L26
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v240
	v244 = F_read_attr_value(m, v238, int32(114))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L26
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v244
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v248 = base.I32_extend8_s(v247)
	if int32(33) <= v248 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v253 = v248
	v258 = v244
	goto L96
L94:
	;
	v282 = v248
	goto L95
L95:
	;
	if v282 != 0 {
		goto L12
	} else {
		goto L100
	}
L96:
	;
	v268 = v253 & int32(255)
	if base.B2i32(v268 == int32(44))|base.B2i32(v268 == int32(127)) != 0 {
		goto L12
	} else {
		goto L98
	}
L97:
	;
	v282 = v277
	goto L95
L98:
	;
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v258)+1)))
	v277 = base.I32_extend8_s(v274)
	if int32(32) < v277 {
		v253 = v277
		v258 = v258 + int32(1)
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	if v297 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	goto L104
L102:
	;
	goto L103
L103:
	;
	v338 = v19 + int32(192)
	v340 = int32(0)
	v344 = m.G0
	v346 = v344 - int32(16)
	m.G0 = v346
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v340
	v352 = F_open(m, int32(_a_F_scram_exchange_10), v340, v346)
	mBase = m.M
	if v352 != int32(-1) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v317 = F_read_any_attr(m, v19+int32(192), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L26
	} else {
		goto L106
	}
L105:
	;
	goto L103
L106:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v320 != 0 {
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	if v385 == int32(0) {
		goto L11
	} else {
		goto L121
	}
L109:
	;
	goto L113
L110:
	;
	v385 = v340
	goto L111
L111:
	;
	m.G0 = v346 + int32(16)
	goto L108
L112:
	;
	v380 = F_close(m, v352)
	mBase = m.M
	v385 = v378
	goto L111
L113:
	;
	v358 = v338
	v359 = int32(18)
	goto L114
L114:
	;
	v364 = F_read(m, v352, v358, v359)
	mBase = m.M
	if v364 <= int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v378 = int32(1)
	goto L112
L116:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_scram_exchange[2]))
	if v368 == int32(27) {
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v373 = v359 - v364
	if v373 != 0 {
		v358 = v358 + v364
		v359 = v373
		goto L114
	} else {
		goto L120
	}
L119:
	;
	v378 = int32(0)
	goto L112
L120:
	;
	goto L115
L121:
	;
	v396 = base.I32_div_s(int32(20), int32(3))
	v398 = v396 << (uint(int32(2)) % 32)
	goto L122
L122:
	;
	v401 = F_palloc(m, v398+int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L26
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v401
	goto L127
L124:
	;
	if v516 < int32(0) {
		goto L10
	} else {
		goto L145
	}
L125:
	;
	if v398 != 0 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	if v398 < v458-v401+int32(4) {
		goto L125
	} else {
		goto L138
	}
L127:
	;
	v413 = v338
	v414 = int32(0)
	v417 = v401
	v418 = int32(2)
	goto L130
L129:
	;
	v516 = v458 - v401
	goto L124
L130:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	v424 = v420<<(uint(v418<<(uint(int32(3))%32))%32) | v414
	if int32(0) < v418 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v459 != int32(2) {
		goto L126
	} else {
		goto L137
	}
L132:
	;
	v457 = v424
	v458 = v417
	v459 = v418 - int32(1)
	goto L134
L133:
	;
	if v398 < v417-v401+int32(4) {
		goto L125
	} else {
		goto L135
	}
L134:
	;
	v461 = v413 + int32(1)
	if base.Ui32(v461) < base.Ui32(v19+int32(210)) {
		v413 = v461
		v414 = v457
		v417 = v458
		v418 = v459
		goto L130
	} else {
		goto L136
	}
L135:
	;
	v433 = int32(63)
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+3)) = uint8(v435)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417))) = uint8(v439)
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(6))%32))&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+2)) = uint8(v445)
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v424)>>(uint(int32(12))%32))&v433)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)) = uint8(v451)
	v457 = int32(0)
	v458 = v417 + int32(4)
	v459 = int32(2)
	goto L134
L136:
	;
	goto L131
L137:
	;
	goto L129
L138:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458))) = uint8(v479)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+1)) = uint8(v485)
	if v459 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v457)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	v495 = v494
	goto L141
L140:
	;
	v495 = int32(61)
	goto L141
L141:
	;
	v496 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+3)) = uint8(v496)
	*(*uint8)(unsafe.Add(mBase, uint32(v458)+2)) = uint8(v495)
	v516 = v458 + int32(4) - v401
	goto L124
L142:
	;
	base.MemoryFill(m, v401, int32(0), v398)
	goto L144
L143:
	;
	goto L144
L144:
	;
	v516 = int32(-1)
	goto L124
L145:
	;
	v519 = int32(0)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	*(*uint8)(unsafe.Add(mBase, uint32(v520+v516))) = uint8(v519)
	v524 = *(*int64)(unsafe.Add(mBase, uint32(l0)+20))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v526
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v525
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = base.I64_rotl(v524, int64(32))
	v535 = F_psprintf(m, int32(_a_F_scram_exchange_11), v19+int32(32))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L26
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v535
	v538 = F_pstrdup(m, v535)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L26
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v538
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v1643 = int32(0)
	v1648 = v519
	goto L30
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+192)) = v544
	v550 = F_read_attr_value(m, v19+int32(192), int32(99))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L26
	} else {
		goto L149
	}
L149:
	;
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v552 == int32(1) {
		goto L9
	} else {
		goto L150
	}
L150:
	;
	v555 = int32(_a_F_scram_exchange_12)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[4])))
	if base.B2i32(v558 == int32(0))|base.B2i32(v558 != v561) != 0 {
		v579 = v558
		v580 = v561
		goto L153
	} else {
		goto L154
	}
L151:
	;
	v620 = F_read_attr_value(m, v19+int32(192), int32(114))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L26
	} else {
		goto L172
	}
L152:
	;
	if v579-v580 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L153:
	;
	goto L152
L154:
	;
	v564 = v550
	v565 = v555
	goto L155
L155:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+1)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+1)))
	if v569 == int32(0) {
		v579 = v569
		v580 = v568
		goto L153
	} else {
		goto L157
	}
L156:
	;
	v579 = v569
	v580 = v568
	goto L153
L157:
	;
	v572 = int32(1)
	if v569 == v568 {
		v564 = v564 + v572
		v565 = v565 + v572
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	if v584 == int32(110) {
		goto L151
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v587 = int32(_a_F_scram_exchange_13)
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_scram_exchange[5])))
	if base.B2i32(v590 == int32(0))|base.B2i32(v590 != v593) != 0 {
		v611 = v590
		v612 = v593
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L161
L163:
	;
	if v611-v612 != 0 {
		goto L8
	} else {
		goto L170
	}
L164:
	;
	goto L163
L165:
	;
	v596 = v550
	v597 = v587
	goto L166
L166:
	;
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597)+1)))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	if v601 == int32(0) {
		v611 = v601
		v612 = v600
		goto L164
	} else {
		goto L168
	}
L167:
	;
	v611 = v601
	v612 = v600
	goto L164
L168:
	;
	v604 = int32(1)
	if v601 == v600 {
		v596 = v596 + v604
		v597 = v597 + v604
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+124)))
	if v614 != int32(121) {
		goto L8
	} else {
		goto L171
	}
L171:
	;
	goto L151
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v620
	goto L173
L173:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v644 = F_read_any_attr(m, v19+int32(192), v19+int32(160))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L26
	} else {
		goto L175
	}
L174:
	;
	v649 = F_strlen(m, v644)
	mBase = m.M
	v653 = v649 * int32(3) >> (uint(int32(2)) % 32)
	goto L177
L175:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+160)))
	if v646 != int32(112) {
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v654 = F_palloc(m, v653)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L26
	} else {
		goto L178
	}
L178:
	;
	v656 = F_strlen(m, v644)
	mBase = m.M
	v657 = int32(0)
	if v657 < v656 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v842 != v843 {
		goto L7
	} else {
		goto L225
	}
L180:
	;
	if v653 != 0 {
		goto L222
	} else {
		goto L223
	}
L181:
	;
	v666 = v644 + v656
	v667 = v644
	v671 = v657
	v672 = v654
	v674 = v657
	v675 = v657
	goto L184
L182:
	;
	v813 = v654
	goto L183
L183:
	;
	v842 = v813 - v654
	goto L179
L184:
	;
	v679 = v667 + int32(1)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	if v680 != int32(61) {
		goto L192
	} else {
		goto L193
	}
L185:
	;
	if v801 != 0 {
		goto L180
	} else {
		goto L221
	}
L186:
	;
	if base.Ui32(v799) < base.Ui32(v666) {
		v667 = v799
		v671 = v801
		v672 = v802
		v674 = v804
		v675 = v805
		goto L184
	} else {
		goto L220
	}
L187:
	;
	if v653 < v672-v654+int32(1) {
		goto L180
	} else {
		goto L207
	}
L188:
	;
	v757 = v679
	v758 = int32(2)
	v761 = v675 << (uint(int32(6)) % 32)
	goto L187
L189:
	;
	v746 = v743 + v742<<(uint(int32(6))%32)
	v748 = v739 + int32(1)
	if v748 == int32(4) {
		v757 = v740
		v758 = v741
		v761 = v746
		goto L187
	} else {
		goto L206
	}
L190:
	;
	v739 = int32(3)
	v740 = v667 + int32(2)
	v741 = int32(1)
	v742 = v699
	v743 = v694
	goto L189
L191:
	;
	if base.Ui32(int32(125)) < base.Ui32((v718-int32(1))&int32(255)) {
		goto L180
	} else {
		goto L204
	}
L192:
	;
	v684 = v680 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v684))|base.B2i32(int32(1)<<(uint(v684)%32)&int32(_a_F_scram_exchange_14) == int32(0)) != 0 {
		v718 = v680
		v719 = v671
		v720 = v679
		v721 = v674
		v722 = v675
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v694 = int32(0)
	if v674 != 0 {
		v739 = v671
		v740 = v679
		v741 = v674
		v742 = v675
		v743 = v694
		goto L189
	} else {
		goto L196
	}
L195:
	;
	goto L180
L196:
	;
	switch v671 - int32(2) {
	case 0:
		goto L197
	case 1:
		goto L188
	default:
		goto L180
	}
L197:
	;
	if base.Ui32(v666) <= base.Ui32(v679) {
		goto L180
	} else {
		goto L198
	}
L198:
	;
	v699 = v675 << (uint(int32(6)) % 32)
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679))))
	if v700 == int32(61) {
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v704 = v700 - int32(9)
	if int32(1)<<(uint(v704)%32)&int32(_a_F_scram_exchange_14) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v712 = base.B2i32(base.Ui32(v704) <= base.Ui32(int32(23)))
	goto L202
L201:
	;
	v712 = int32(0)
	goto L202
L202:
	;
	if v712 != 0 {
		goto L180
	} else {
		goto L203
	}
L203:
	;
	v718 = v700
	v719 = int32(3)
	v720 = v667 + int32(2)
	v721 = int32(1)
	v722 = v699
	goto L191
L204:
	;
	v730 = int32(*(*int8)(unsafe.Add(mBase, uint32(v718)+uint32(_c_F_scram_exchange[6]))))
	if v730 < int32(0) {
		goto L180
	} else {
		goto L205
	}
L205:
	;
	v739 = v719
	v740 = v720
	v741 = v721
	v742 = v722
	v743 = v730
	goto L189
L206:
	;
	v799 = v740
	v801 = v748
	v802 = v672
	v804 = v741
	v805 = v746
	goto L186
L207:
	;
	v767 = int32(base.Ui32(v761) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v672))) = uint8(v767)
	v770 = v672 + int32(1)
	if base.Ui32(v758) < base.Ui32(int32(2)) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v774 = v758
	goto L210
L209:
	;
	v774 = int32(0)
	goto L210
L210:
	;
	if v774 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	if v653 < v770-v654+int32(1) {
		goto L180
	} else {
		goto L214
	}
L212:
	;
	v786 = v770
	goto L213
L213:
	;
	if v758 != 0 {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	v782 = int32(base.Ui32(v761) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v672)+1)) = uint8(v782)
	v786 = v672 + int32(2)
	goto L213
L215:
	;
	v799 = v757
	v801 = int32(0)
	v802 = v796
	v804 = v797
	v805 = int32(0)
	goto L186
L216:
	;
	v796 = v786
	v797 = v758
	goto L215
L217:
	;
	goto L218
L218:
	;
	if v653 < v786-v654+int32(1) {
		goto L180
	} else {
		goto L219
	}
L219:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v786))) = uint8(v761)
	v796 = v786 + int32(1)
	v797 = int32(0)
	goto L215
L220:
	;
	goto L185
L221:
	;
	v813 = v802
	goto L183
L222:
	;
	base.MemoryFill(m, v654, int32(0), v653)
	goto L224
L223:
	;
	goto L224
L224:
	;
	v842 = int32(-1)
	goto L179
L225:
	;
	v846 = l0 + int32(148)
	if v842 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	base.MemoryCopy(m, v846, v654, v842)
	goto L228
L227:
	;
	goto L228
L228:
	;
	F_pfree(m, v654)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L26
	} else {
		goto L229
	}
L229:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v850))))
	if v851 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	v853 = F_palloc(m, v639-v544)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L26
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v853
	v858 = v544 ^ int32(-1) + v639
	if v858 != 0 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	base.MemoryCopy(m, v853, l1, v858)
	goto L234
L233:
	;
	goto L234
L234:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v862 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v860+v858))) = uint8(v862)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v865 = F_strlen(m, v864)
	mBase = m.M
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v867 = F_strlen(m, v866)
	mBase = m.M
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v869 = F_strlen(m, v868)
	mBase = m.M
	if v869 != v865+v867 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	v872 = int32(0)
	if v865 == v872 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	if v960 != 0 {
		goto L5
	} else {
		goto L252
	}
L237:
	;
	v960 = int32(0)
	goto L236
L238:
	;
	goto L239
L239:
	;
	v880 = v865 & int32(3)
	if base.Ui32(v865) < base.Ui32(int32(4)) {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	v960 = base.B2i32(v946 != int32(0))
	goto L236
L241:
	;
	v926 = v919
	v927 = v920
	v928 = v921
	v932 = v872
	goto L249
L242:
	;
	v919 = v868
	v920 = v864
	v921 = int32(0)
	goto L241
L243:
	;
	goto L244
L244:
	;
	v887 = v868
	v888 = v864
	v889 = int32(0)
	v892 = v872
	goto L245
L245:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	v898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888)+1)))
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)))
	v902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888)+2)))
	v903 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+2)))
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888)+3)))
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+3)))
	v909 = v889 | (v894 ^ v895) | (v898 ^ v899) | (v902 ^ v903) | (v906 ^ v907)
	v910 = int32(4)
	v911 = v888 + v910
	v913 = v887 + v910
	v915 = v892 + v910
	if v915 != v865&int32(-4) {
		v887 = v913
		v888 = v911
		v889 = v909
		v892 = v915
		goto L245
	} else {
		goto L247
	}
L246:
	;
	if v880 == int32(0) {
		v946 = v909
		goto L240
	} else {
		goto L248
	}
L247:
	;
	goto L246
L248:
	;
	v919 = v913
	v920 = v911
	v921 = v909
	goto L241
L249:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v926))))
	v936 = v928 | (v933 ^ v934)
	v937 = int32(1)
	v942 = v932 + v937
	if v942 != v880 {
		v926 = v926 + v937
		v927 = v927 + v937
		v928 = v936
		v932 = v942
		goto L249
	} else {
		goto L251
	}
L250:
	;
	v946 = v936
	goto L240
L251:
	;
	goto L250
L252:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v962 = v961 + v865
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v964 = int32(0)
	if v867 == v964 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	if v1052 != 0 {
		goto L5
	} else {
		goto L269
	}
L254:
	;
	v1052 = int32(0)
	goto L253
L255:
	;
	goto L256
L256:
	;
	v972 = v867 & int32(3)
	if base.Ui32(v867) < base.Ui32(int32(4)) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	v1052 = base.B2i32(v1038 != int32(0))
	goto L253
L258:
	;
	v1018 = v1011
	v1019 = v1012
	v1020 = v1013
	v1024 = v964
	goto L266
L259:
	;
	v1011 = v962
	v1012 = v963
	v1013 = int32(0)
	goto L258
L260:
	;
	goto L261
L261:
	;
	v979 = v962
	v980 = v963
	v981 = int32(0)
	v984 = v964
	goto L262
L262:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980))))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979))))
	v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+1)))
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+1)))
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+2)))
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+2)))
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980)+3)))
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979)+3)))
	v1001 = v981 | (v986 ^ v987) | (v990 ^ v991) | (v994 ^ v995) | (v998 ^ v999)
	v1002 = int32(4)
	v1003 = v980 + v1002
	v1005 = v979 + v1002
	v1007 = v984 + v1002
	if v1007 != v867&int32(-4) {
		v979 = v1005
		v980 = v1003
		v981 = v1001
		v984 = v1007
		goto L262
	} else {
		goto L264
	}
L263:
	;
	if v972 == int32(0) {
		v1038 = v1001
		goto L257
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	v1011 = v1005
	v1012 = v1003
	v1013 = v1001
	goto L258
L266:
	;
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019))))
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1018))))
	v1028 = v1020 | (v1025 ^ v1026)
	v1029 = int32(1)
	v1034 = v1024 + v1029
	if v1034 != v972 {
		v1018 = v1018 + v1029
		v1019 = v1019 + v1029
		v1020 = v1028
		v1024 = v1034
		goto L266
	} else {
		goto L268
	}
L267:
	;
	v1038 = v1028
	goto L257
L268:
	;
	goto L267
L269:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1054 = F_pg_hmac_create(m, v1053)
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L26
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+156)) = int32(0)
	v1059 = l0 + int32(60)
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1061 = F_pg_hmac_init(m, v1054, v1059, v1060)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L26
	} else {
		goto L271
	}
L271:
	;
	if v1061 < int32(0) {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1066 = F_strlen(m, v1065)
	mBase = m.M
	if v1054 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	if v1081 < int32(0) {
		goto L4
	} else {
		goto L280
	}
L274:
	;
	v1081 = int32(-1)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1071 = F_pg_cryptohash_update(m, v1070, v1065, v1066)
	mBase = m.M
	if int32(0) <= v1071 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1081 = int32(0)
	goto L273
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+8)) = int32(2)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1078 = F_pg_cryptohash_error(m, v1077)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+12)) = v1078
	v1081 = int32(-1)
	goto L273
L280:
	;
	if v1054 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	if v1100 < int32(0) {
		goto L4
	} else {
		goto L288
	}
L282:
	;
	v1100 = int32(-1)
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1090 = F_pg_cryptohash_update(m, v1089, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1090 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1100 = int32(0)
	goto L281
L286:
	;
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+8)) = int32(2)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1097 = F_pg_cryptohash_error(m, v1096)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+12)) = v1097
	v1100 = int32(-1)
	goto L281
L288:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v1104 = F_strlen(m, v1103)
	mBase = m.M
	if v1054 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	if v1119 < int32(0) {
		goto L4
	} else {
		goto L296
	}
L290:
	;
	v1119 = int32(-1)
	goto L289
L291:
	;
	goto L292
L292:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1109 = F_pg_cryptohash_update(m, v1108, v1103, v1104)
	mBase = m.M
	if int32(0) <= v1109 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	v1119 = int32(0)
	goto L289
L294:
	;
	goto L295
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+8)) = int32(2)
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1116 = F_pg_cryptohash_error(m, v1115)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+12)) = v1116
	v1119 = int32(-1)
	goto L289
L296:
	;
	if v1054 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L297:
	;
	if v1138 < int32(0) {
		goto L4
	} else {
		goto L304
	}
L298:
	;
	v1138 = int32(-1)
	goto L297
L299:
	;
	goto L300
L300:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1128 = F_pg_cryptohash_update(m, v1127, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1128 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	v1138 = int32(0)
	goto L297
L302:
	;
	goto L303
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+8)) = int32(2)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1135 = F_pg_cryptohash_error(m, v1134)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+12)) = v1135
	v1138 = int32(-1)
	goto L297
L304:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1142 = F_strlen(m, v1141)
	mBase = m.M
	if v1054 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	if v1157 < int32(0) {
		goto L4
	} else {
		goto L312
	}
L306:
	;
	v1157 = int32(-1)
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1147 = F_pg_cryptohash_update(m, v1146, v1141, v1142)
	mBase = m.M
	if int32(0) <= v1147 {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1157 = int32(0)
	goto L305
L310:
	;
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+8)) = int32(2)
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1054)))
	v1154 = F_pg_cryptohash_error(m, v1153)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1054)+12)) = v1154
	v1157 = int32(-1)
	goto L305
L312:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1163 = F_pg_hmac_final(m, v1054, v19+int32(192), v1162)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L26
	} else {
		goto L313
	}
L313:
	;
	if v1163 < int32(0) {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	F_pg_hmac_free(m, v1054)
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L26
	} else {
		goto L315
	}
L315:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v1169 <= int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1268 = v19 + int32(160)
	v1271 = F_scram_H(m, l0+int32(28), v1266, v1169, v1268, v19+int32(156))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L26
	} else {
		goto L325
	}
L317:
	;
	v1173 = l0 + int32(28)
	v1174 = int32(0)
	if v1169 != int32(1) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1189 = v1174
	v1192 = int32(0)
	goto L321
L319:
	;
	v1230 = v1174
	goto L320
L320:
	;
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(192)+v1230))))
	v1245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1230))))
	v1246 = v1243 ^ v1245
	*(*uint8)(unsafe.Add(mBase, uint32(v1230+v1173))) = uint8(v1246)
	goto L316
L321:
	;
	v1200 = v19 + int32(192)
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200+v1189))))
	v1204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1189))))
	v1205 = v1202 ^ v1204
	*(*uint8)(unsafe.Add(mBase, uint32(v1189+v1173))) = uint8(v1205)
	v1208 = v1189 | int32(1)
	v1211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1208+v1200))))
	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v1208))))
	v1214 = v1211 ^ v1213
	*(*uint8)(unsafe.Add(mBase, uint32(v1173+v1208))) = uint8(v1214)
	v1216 = int32(2)
	v1217 = v1189 + v1216
	v1219 = v1192 + v1216
	if v1219 != v1169&int32(2147483646) {
		v1189 = v1217
		v1192 = v1219
		goto L321
	} else {
		goto L323
	}
L322:
	;
	if v1169&int32(1) == int32(0) {
		goto L316
	} else {
		goto L324
	}
L323:
	;
	goto L322
L324:
	;
	v1230 = v1217
	goto L320
L325:
	;
	if v1271 < int32(0) {
		goto L3
	} else {
		goto L326
	}
L326:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1276 = int32(0)
	if v1275 == v1276 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if v1364 != 0 {
		goto L31
	} else {
		goto L343
	}
L328:
	;
	v1364 = int32(0)
	goto L327
L329:
	;
	goto L330
L330:
	;
	v1284 = v1275 & int32(3)
	if base.Ui32(v1275) < base.Ui32(int32(4)) {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	v1364 = base.B2i32(v1350 != int32(0))
	goto L327
L332:
	;
	v1330 = v1323
	v1331 = v1324
	v1332 = v1325
	v1336 = v1276
	goto L340
L333:
	;
	v1323 = v1268
	v1324 = v1059
	v1325 = int32(0)
	goto L332
L334:
	;
	goto L335
L335:
	;
	v1291 = v1268
	v1292 = v1059
	v1293 = int32(0)
	v1296 = v1276
	goto L336
L336:
	;
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292))))
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291))))
	v1302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+1)))
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+1)))
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+2)))
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+2)))
	v1310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292)+3)))
	v1311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+3)))
	v1313 = v1293 | (v1298 ^ v1299) | (v1302 ^ v1303) | (v1306 ^ v1307) | (v1310 ^ v1311)
	v1314 = int32(4)
	v1315 = v1292 + v1314
	v1317 = v1291 + v1314
	v1319 = v1296 + v1314
	if v1319 != v1275&int32(-4) {
		v1291 = v1317
		v1292 = v1315
		v1293 = v1313
		v1296 = v1319
		goto L336
	} else {
		goto L338
	}
L337:
	;
	if v1284 == int32(0) {
		v1350 = v1313
		goto L331
	} else {
		goto L339
	}
L338:
	;
	goto L337
L339:
	;
	v1323 = v1317
	v1324 = v1315
	v1325 = v1313
	goto L332
L340:
	;
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1331))))
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330))))
	v1340 = v1332 | (v1337 ^ v1338)
	v1341 = int32(1)
	v1346 = v1336 + v1341
	if v1346 != v1284 {
		v1330 = v1330 + v1341
		v1331 = v1331 + v1341
		v1332 = v1340
		v1336 = v1346
		goto L340
	} else {
		goto L342
	}
L341:
	;
	v1350 = v1340
	goto L331
L342:
	;
	goto L341
L343:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)))
	if v1365 != 0 {
		goto L31
	} else {
		goto L344
	}
L344:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1367 = F_pg_hmac_create(m, v1366)
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L26
	} else {
		goto L345
	}
L345:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1372 = F_pg_hmac_init(m, v1367, l0+int32(92), v1371)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L26
	} else {
		goto L346
	}
L346:
	;
	if v1372 < int32(0) {
		goto L2
	} else {
		goto L347
	}
L347:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1377 = F_strlen(m, v1376)
	mBase = m.M
	if v1367 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	if v1392 < int32(0) {
		goto L2
	} else {
		goto L355
	}
L349:
	;
	v1392 = int32(-1)
	goto L348
L350:
	;
	goto L351
L351:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1382 = F_pg_cryptohash_update(m, v1381, v1376, v1377)
	mBase = m.M
	if int32(0) <= v1382 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1392 = int32(0)
	goto L348
L353:
	;
	goto L354
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+8)) = int32(2)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1389 = F_pg_cryptohash_error(m, v1388)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+12)) = v1389
	v1392 = int32(-1)
	goto L348
L355:
	;
	if v1367 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v1411 < int32(0) {
		goto L2
	} else {
		goto L363
	}
L357:
	;
	v1411 = int32(-1)
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1401 = F_pg_cryptohash_update(m, v1400, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1401 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1411 = int32(0)
	goto L356
L361:
	;
	goto L362
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+8)) = int32(2)
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1408 = F_pg_cryptohash_error(m, v1407)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+12)) = v1408
	v1411 = int32(-1)
	goto L356
L363:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v1415 = F_strlen(m, v1414)
	mBase = m.M
	if v1367 == int32(0) {
		goto L365
	} else {
		goto L366
	}
L364:
	;
	if v1430 < int32(0) {
		goto L2
	} else {
		goto L371
	}
L365:
	;
	v1430 = int32(-1)
	goto L364
L366:
	;
	goto L367
L367:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1420 = F_pg_cryptohash_update(m, v1419, v1414, v1415)
	mBase = m.M
	if int32(0) <= v1420 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v1430 = int32(0)
	goto L364
L369:
	;
	goto L370
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+8)) = int32(2)
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1427 = F_pg_cryptohash_error(m, v1426)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+12)) = v1427
	v1430 = int32(-1)
	goto L364
L371:
	;
	if v1367 == int32(0) {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	if v1449 < int32(0) {
		goto L2
	} else {
		goto L379
	}
L373:
	;
	v1449 = int32(-1)
	goto L372
L374:
	;
	goto L375
L375:
	;
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1439 = F_pg_cryptohash_update(m, v1438, int32(_a_F_scram_exchange_15), int32(1))
	mBase = m.M
	if int32(0) <= v1439 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1449 = int32(0)
	goto L372
L377:
	;
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+8)) = int32(2)
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1446 = F_pg_cryptohash_error(m, v1445)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+12)) = v1446
	v1449 = int32(-1)
	goto L372
L379:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v1453 = F_strlen(m, v1452)
	mBase = m.M
	if v1367 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	if v1468 < int32(0) {
		goto L2
	} else {
		goto L387
	}
L381:
	;
	v1468 = int32(-1)
	goto L380
L382:
	;
	goto L383
L383:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1458 = F_pg_cryptohash_update(m, v1457, v1452, v1453)
	mBase = m.M
	if int32(0) <= v1458 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1468 = int32(0)
	goto L380
L385:
	;
	goto L386
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+8)) = int32(2)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1367)))
	v1465 = F_pg_cryptohash_error(m, v1464)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1367)+12)) = v1465
	v1468 = int32(-1)
	goto L380
L387:
	;
	v1472 = v19 + int32(192)
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1474 = F_pg_hmac_final(m, v1367, v1472, v1473)
	mBase = m.M
	v1475 = m.ExcPending
	if v1475 != 0 {
		goto L26
	} else {
		goto L388
	}
L388:
	;
	if v1474 < int32(0) {
		goto L2
	} else {
		goto L389
	}
L389:
	;
	F_pg_hmac_free(m, v1367)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L26
	} else {
		goto L390
	}
L390:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1482 = int32(2)
	v1485 = base.I32_div_s(v1481+v1482, int32(3))
	v1487 = v1485 << (uint(v1482) % 32)
	goto L391
L391:
	;
	v1490 = F_palloc(m, v1487+int32(1))
	mBase = m.M
	v1491 = m.ExcPending
	if v1491 != 0 {
		goto L26
	} else {
		goto L392
	}
L392:
	;
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if int32(0) < v1492 {
		goto L396
	} else {
		goto L397
	}
L393:
	;
	if v1604 < int32(0) {
		goto L1
	} else {
		goto L414
	}
L394:
	;
	if v1487 != 0 {
		goto L411
	} else {
		goto L412
	}
L395:
	;
	if v1487 < v1546-v1490+int32(4) {
		goto L394
	} else {
		goto L407
	}
L396:
	;
	v1501 = v1472
	v1502 = int32(0)
	v1505 = v1490
	v1506 = int32(2)
	goto L399
L397:
	;
	v1557 = v1490
	goto L398
L398:
	;
	v1604 = v1557 - v1490
	goto L393
L399:
	;
	v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1501))))
	v1512 = v1508<<(uint(v1506<<(uint(int32(3))%32))%32) | v1502
	if int32(0) < v1506 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	if v1547 != int32(2) {
		goto L395
	} else {
		goto L406
	}
L401:
	;
	v1545 = v1512
	v1546 = v1505
	v1547 = v1506 - int32(1)
	goto L403
L402:
	;
	if v1487 < v1505-v1490+int32(4) {
		goto L394
	} else {
		goto L404
	}
L403:
	;
	v1549 = v1501 + int32(1)
	if base.Ui32(v1549) < base.Ui32(v1472+v1492) {
		v1501 = v1549
		v1502 = v1545
		v1505 = v1546
		v1506 = v1547
		goto L399
	} else {
		goto L405
	}
L404:
	;
	v1521 = int32(63)
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512&v1521)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1505)+3)) = uint8(v1523)
	v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1512)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1505))) = uint8(v1527)
	v1533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1512)>>(uint(int32(6))%32))&v1521)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1505)+2)) = uint8(v1533)
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1512)>>(uint(int32(12))%32))&v1521)+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1505)+1)) = uint8(v1539)
	v1545 = int32(0)
	v1546 = v1505 + int32(4)
	v1547 = int32(2)
	goto L403
L405:
	;
	goto L400
L406:
	;
	v1557 = v1546
	goto L398
L407:
	;
	v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1545)>>(uint(int32(18))%32)))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1546))) = uint8(v1567)
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1545)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1546)+1)) = uint8(v1573)
	if v1547 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1545)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_scram_exchange[3]))))
	v1583 = v1582
	goto L410
L409:
	;
	v1583 = int32(61)
	goto L410
L410:
	;
	v1584 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v1546)+3)) = uint8(v1584)
	*(*uint8)(unsafe.Add(mBase, uint32(v1546)+2)) = uint8(v1583)
	v1604 = v1546 + int32(4) - v1490
	goto L393
L411:
	;
	base.MemoryFill(m, v1490, int32(0), v1487)
	goto L413
L412:
	;
	goto L413
L413:
	;
	v1604 = int32(-1)
	goto L393
L414:
	;
	v1608 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1490+v1604))) = uint8(v1608)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v1490
	v1614 = F_psprintf(m, int32(_a_F_scram_exchange_16), v19+int32(144))
	mBase = m.M
	v1615 = m.ExcPending
	if v1615 != 0 {
		goto L26
	} else {
		goto L415
	}
L415:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1614
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2)
	v1643 = int32(1)
	v1648 = int32(1)
	goto L30
L416:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_17), int32(0))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L26
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(455), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L26
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L419:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v1637 == int32(0) {
		v1643 = v1633
		v1648 = v1634
		goto L30
	} else {
		goto L420
	}
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v1637
	v1643 = v1633
	v1648 = v1634
	goto L30
L421:
	;
	v1658 = F_strlen(m, v1657)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v1658
	goto L423
L422:
	;
	goto L423
L423:
	;
	if v1643 == int32(0) {
		v1692 = v1648
		goto L22
	} else {
		goto L424
	}
L424:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1662 != int32(2) {
		v1692 = v1648
		goto L22
	} else {
		goto L425
	}
L425:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, _c_F_scram_exchange[7]))
	v1667 = *(*int64)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+440)) = v1667
	v1669 = *(*int64)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+432)) = v1669
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+424)) = v1671
	v1673 = *(*int64)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+416)) = v1673
	v1675 = *(*int64)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+464)) = v1675
	v1677 = *(*int64)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+456)) = v1677
	v1679 = *(*int64)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+448)) = v1679
	v1681 = *(*int64)(unsafe.Add(mBase, uint32(l0)+116))
	v1682 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1666)+480)) = uint8(v1682)
	*(*int64)(unsafe.Add(mBase, uint32(v1666)+472)) = v1681
	v1692 = v1648
	goto L22
L426:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L26
	} else {
		goto L427
	}
L427:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L26
	} else {
		goto L428
	}
L428:
	;
	v1718 = F_errdetail(m, int32(_a_F_scram_exchange_19), int32(0))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L26
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(381), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L26
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L26
	} else {
		goto L432
	}
L432:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L26
	} else {
		goto L433
	}
L433:
	;
	v1738 = F_errdetail(m, int32(_a_F_scram_exchange_20), int32(0))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L26
	} else {
		goto L434
	}
L434:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(386), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L26
	} else {
		goto L435
	}
L435:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L436:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L26
	} else {
		goto L437
	}
L437:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L26
	} else {
		goto L438
	}
L438:
	;
	v1758 = F_errdetail(m, int32(_a_F_scram_exchange_21), int32(0))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L26
	} else {
		goto L439
	}
L439:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(994), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L26
	} else {
		goto L440
	}
L440:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L441:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L26
	} else {
		goto L442
	}
L442:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L26
	} else {
		goto L443
	}
L443:
	;
	v1776 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
	F_sanitize_char_2(m, v1776)
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L26
	} else {
		goto L444
	}
L444:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(_a_F_scram_exchange_3)
	v1784 = F_errdetail(m, int32(_a_F_scram_exchange_22), v19+int32(48))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L26
	} else {
		goto L445
	}
L445:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1002), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L26
	} else {
		goto L446
	}
L446:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L447:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L26
	} else {
		goto L448
	}
L448:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1801 = m.ExcPending
	if v1801 != 0 {
		goto L26
	} else {
		goto L449
	}
L449:
	;
	v1804 = F_errdetail(m, int32(_a_F_scram_exchange_21), int32(0))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L26
	} else {
		goto L450
	}
L450:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1016), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L26
	} else {
		goto L451
	}
L451:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L452:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L26
	} else {
		goto L453
	}
L453:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L26
	} else {
		goto L454
	}
L454:
	;
	v1822 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+1)))
	F_sanitize_char_2(m, v1822)
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L26
	} else {
		goto L455
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = int32(_a_F_scram_exchange_3)
	v1830 = F_errdetail(m, int32(_a_F_scram_exchange_22), v19-int32(-64))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L26
	} else {
		goto L456
	}
L456:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1032), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1836 = m.ExcPending
	if v1836 != 0 {
		goto L26
	} else {
		goto L457
	}
L457:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L458:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L26
	} else {
		goto L459
	}
L459:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L26
	} else {
		goto L460
	}
L460:
	;
	v1850 = F_errdetail(m, int32(_a_F_scram_exchange_23), int32(0))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L26
	} else {
		goto L461
	}
L461:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1045), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L26
	} else {
		goto L462
	}
L462:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L463:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L26
	} else {
		goto L464
	}
L464:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_24), int32(0))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L26
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1073), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1872 = m.ExcPending
	if v1872 != 0 {
		goto L26
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L26
	} else {
		goto L468
	}
L468:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_25), int32(0))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L26
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1094), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L26
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L26
	} else {
		goto L472
	}
L472:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_26), int32(0))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L26
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1108), int32(_a_F_scram_exchange_6))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L26
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L26
	} else {
		goto L476
	}
L476:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_27), int32(0))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L26
	} else {
		goto L477
	}
L477:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1238), int32(_a_F_scram_exchange_28))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L26
	} else {
		goto L478
	}
L478:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L479:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L26
	} else {
		goto L480
	}
L480:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_29), int32(0))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L26
	} else {
		goto L481
	}
L481:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1248), int32(_a_F_scram_exchange_28))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L26
	} else {
		goto L482
	}
L482:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L483:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_30), int32(0))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L26
	} else {
		goto L484
	}
L484:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1357), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L26
	} else {
		goto L485
	}
L485:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L486:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L26
	} else {
		goto L487
	}
L487:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_32), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L26
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1372), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L26
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L26
	} else {
		goto L491
	}
L491:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L26
	} else {
		goto L492
	}
L492:
	;
	v1995 = F_errdetail(m, int32(_a_F_scram_exchange_33), int32(0))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L26
	} else {
		goto L493
	}
L493:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1391), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L26
	} else {
		goto L494
	}
L494:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L495:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L26
	} else {
		goto L496
	}
L496:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_2), int32(0))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L26
	} else {
		goto L497
	}
L497:
	;
	v2015 = F_errdetail(m, int32(_a_F_scram_exchange_34), int32(0))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L26
	} else {
		goto L498
	}
L498:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1399), int32(_a_F_scram_exchange_31))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L26
	} else {
		goto L499
	}
L499:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L500:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L26
	} else {
		goto L501
	}
L501:
	;
	F_errmsg(m, int32(_a_F_scram_exchange_35), int32(0))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L26
	} else {
		goto L502
	}
L502:
	;
	v2035 = F_errdetail(m, int32(_a_F_scram_exchange_36), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L26
	} else {
		goto L503
	}
L503:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(419), int32(_a_F_scram_exchange_18))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L26
	} else {
		goto L504
	}
L504:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L505:
	;
	if v1054 == int32(0) {
		goto L507
	} else {
		goto L508
	}
L506:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v2066
	F_errmsg_internal(m, int32(_a_F_scram_exchange_37), v19+int32(96))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L26
	} else {
		goto L519
	}
L507:
	;
	v2066 = int32(_a_F_scram_exchange_38)
	goto L506
L508:
	;
	goto L509
L509:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+12))
	if v2051 != 0 {
		goto L510
	} else {
		goto L511
	}
L510:
	;
	v2063 = v2051
	goto L512
L511:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v1054)+8))
	if v2055 == int32(2) {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v2066 = v2063
	goto L506
L513:
	;
	v2058 = int32(_a_F_scram_exchange_39)
	goto L515
L514:
	;
	v2058 = int32(_a_F_scram_exchange_40)
	goto L515
L515:
	;
	if v2055 == int32(1) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v2061 = int32(_a_F_scram_exchange_38)
	goto L518
L517:
	;
	v2061 = v2058
	goto L518
L518:
	;
	v2063 = v2061
	goto L512
L519:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1175), int32(_a_F_scram_exchange_41))
	mBase = m.M
	v2077 = m.ExcPending
	if v2077 != 0 {
		goto L26
	} else {
		goto L520
	}
L520:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L521:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v19)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v2082
	F_errmsg_internal(m, int32(_a_F_scram_exchange_42), v19+int32(112))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L26
	} else {
		goto L522
	}
L522:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1187), int32(_a_F_scram_exchange_41))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L26
	} else {
		goto L523
	}
L523:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L524:
	;
	if v1367 == int32(0) {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v2119
	F_errmsg_internal(m, int32(_a_F_scram_exchange_43), v19+int32(128))
	mBase = m.M
	v2125 = m.ExcPending
	if v2125 != 0 {
		goto L26
	} else {
		goto L538
	}
L526:
	;
	v2119 = int32(_a_F_scram_exchange_38)
	goto L525
L527:
	;
	goto L528
L528:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+12))
	if v2104 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2116 = v2104
	goto L531
L530:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v1367)+8))
	if v2108 == int32(2) {
		goto L532
	} else {
		goto L533
	}
L531:
	;
	v2119 = v2116
	goto L525
L532:
	;
	v2111 = int32(_a_F_scram_exchange_39)
	goto L534
L533:
	;
	v2111 = int32(_a_F_scram_exchange_40)
	goto L534
L534:
	;
	if v2108 == int32(1) {
		goto L535
	} else {
		goto L536
	}
L535:
	;
	v2114 = int32(_a_F_scram_exchange_38)
	goto L537
L536:
	;
	v2114 = v2111
	goto L537
L537:
	;
	v2116 = v2114
	goto L531
L538:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1433), int32(_a_F_scram_exchange_44))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L26
	} else {
		goto L539
	}
L539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L540:
	;
	F_errmsg_internal(m, int32(_a_F_scram_exchange_45), int32(0))
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L26
	} else {
		goto L541
	}
L541:
	;
	F_errfinish(m, int32(_a_F_scram_exchange_5), int32(1445), int32(_a_F_scram_exchange_44))
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L26
	} else {
		goto L542
	}
L542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_scram_get_mechanisms(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_appendStringInfoString(m, l1, int32(_a_F_scram_get_mechanisms_0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_appendStringInfoChar(m, l1, int32(0))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_script_error_callback(m *base.Module, l0 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v16
	v18 = F_geterrposition(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v282 = int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v287 = F_strlen(m, v283)
	mBase = m.M
	v294 = v287 + v282
	goto L77
L2:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v275 = v270
	v281 = base.B2i32(int32(0) <= v270)
	goto L1
L3:
	;
	return
L4:
	;
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if (base.B2i32(v16 <= int32(0))|base.B2i32(v18 <= v14+v16))&base.B2i32(base.Ui32(v14) <= base.Ui32(v18)) != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v174 = int32(0)
	if v14 < v174 {
		v275 = v14
		v281 = v174
		goto L1
	} else {
		goto L53
	}
L8:
	;
	v82 = v11 + int32(44)
	v84 = v11 + int32(40)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v89 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L9:
	;
	v29 = F_strlen(m, v13)
	mBase = m.M
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v30
	if v29 <= v30 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = int32(0)
	v42 = int32(0)
	goto L11
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v13))))
	if v46 != int32(59) {
		v68 = v40
		v69 = v42
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v71 = v68 + int32(1)
	if v71 < v29 {
		v40 = v71
		v42 = v69
		goto L11
	} else {
		goto L23
	}
L14:
	;
	v50 = v40 + int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v13))))
	if v52 == int32(13) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v55 = v50
	goto L17
L16:
	;
	v55 = v40
	goto L17
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v55)+1)))
	if v57 != int32(10) {
		v68 = v55
		v69 = v42
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v61 = v55 + int32(2)
	if v61 < v18 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v61
	v68 = v55
	v69 = v61
	goto L13
L20:
	;
	goto L21
L21:
	;
	if v61 <= v18 {
		v68 = v55
		v69 = v42
		goto L13
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v61 - v42
	goto L8
L23:
	;
	goto L12
L24:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v158 = F_errposition(m, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L43
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v149
	goto L24
L26:
	;
	v107 = v103
	v110 = v104
	v111 = v105
	goto L33
L27:
	;
	v100 = F_strlen(m, v97)
	mBase = m.M
	if v100 <= int32(0) {
		v146 = v97
		v149 = v100
		v150 = v99
		goto L25
	} else {
		goto L32
	}
L28:
	;
	v97 = v13
	v99 = int32(0)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v93 = v13 + v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if int32(0) < v94 {
		v103 = v93
		v104 = v94
		v105 = v89
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v97 = v93
	v99 = v89
	goto L27
L32:
	;
	v103 = v97
	v104 = v100
	v105 = v99
	goto L26
L33:
	;
	v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107))))
	v115 = F_scanner_isspace(m, v114)
	mBase = m.M
	if v115 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v146 = v140
	v149 = int32(0)
	v150 = v104 + v105
	goto L25
L35:
	;
	v121 = v110
	goto L38
L36:
	;
	goto L37
L37:
	;
	v137 = int32(1)
	v140 = v107 + v137
	if v137 < v110 {
		v107 = v140
		v110 = v110 - v137
		v111 = v111 + v137
		goto L33
	} else {
		goto L42
	}
L38:
	;
	v128 = int32(*(*int8)(unsafe.Add(mBase, uint32(v107+v121-int32(1)))))
	v129 = F_scanner_isspace(m, v128)
	mBase = m.M
	if v129 == int32(0) {
		v146 = v107
		v149 = v121
		v150 = v111
		goto L25
	} else {
		goto L40
	}
L39:
	;
	v146 = v107
	v149 = int32(0)
	v150 = v111
	goto L25
L40:
	;
	v132 = int32(1)
	if v132 < v121 {
		v121 = v121 - v132
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L34
L43:
	;
	v160 = v18 - v155
	if v160 < v156 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v162 = v160
	goto L46
L45:
	;
	v162 = v156
	goto L46
L46:
	;
	v163 = int32(0)
	if v163 <= v160 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v166 = v162
	goto L49
L48:
	;
	v166 = v163
	goto L49
L49:
	;
	F_internalerrposition(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v170 = F_pnstrdup(m, v146, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	v172 = F_internalerrquery(m, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L2
L53:
	;
	v178 = v11 + int32(44)
	v180 = v11 + int32(40)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v185 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L73
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v245
	goto L54
L56:
	;
	v203 = v199
	v206 = v200
	v207 = v201
	goto L63
L57:
	;
	v196 = F_strlen(m, v193)
	mBase = m.M
	if v196 <= int32(0) {
		v242 = v193
		v245 = v196
		v246 = v195
		goto L55
	} else {
		goto L62
	}
L58:
	;
	v193 = v13
	v195 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v189 = v13 + v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if int32(0) < v190 {
		v199 = v189
		v200 = v190
		v201 = v185
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v193 = v189
	v195 = v185
	goto L57
L62:
	;
	v199 = v193
	v200 = v196
	v201 = v195
	goto L56
L63:
	;
	v210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v203))))
	v211 = F_scanner_isspace(m, v210)
	mBase = m.M
	if v211 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v242 = v236
	v245 = int32(0)
	v246 = v200 + v201
	goto L55
L65:
	;
	v217 = v206
	goto L68
L66:
	;
	goto L67
L67:
	;
	v233 = int32(1)
	v236 = v203 + v233
	if v233 < v206 {
		v203 = v236
		v206 = v206 - v233
		v207 = v207 + v233
		goto L63
	} else {
		goto L72
	}
L68:
	;
	v224 = int32(*(*int8)(unsafe.Add(mBase, uint32(v203+v217-int32(1)))))
	v225 = F_scanner_isspace(m, v224)
	mBase = m.M
	if v225 == int32(0) {
		v242 = v203
		v245 = v217
		v246 = v207
		goto L55
	} else {
		goto L70
	}
L69:
	;
	v242 = v203
	v245 = int32(0)
	v246 = v207
	goto L55
L70:
	;
	v228 = int32(1)
	if v228 < v217 {
		v217 = v217 - v228
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L64
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v242
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v255
	F_errcontext_msg(m, int32(_a_F_script_error_callback_0), v11+int32(32))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	goto L2
L75:
	;
	if v306 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v296 = int32(0)
	if v294 == v296 {
		v306 = v296
		goto L76
	} else {
		goto L79
	}
L78:
	;
	v306 = v301
	goto L76
L79:
	;
	v300 = v294 - int32(1)
	v301 = v283 + v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v302 != int32(47) {
		v294 = v300
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v309 = v306 + int32(1)
	goto L83
L82:
	;
	v309 = v283
	goto L83
L83:
	;
	if v281 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	m.G0 = v11 + int32(48)
	return
L85:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v311 == int32(0) {
		v339 = v282
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L96
	}
L88:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L3
	} else {
		goto L94
	}
L89:
	;
	v316 = v275
	v317 = v310
	v319 = v282
	goto L90
L90:
	;
	v323 = v316 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v323
	if v323 < int32(0) {
		v339 = v319
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v339 = v330
	goto L88
L92:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	v330 = v319 + base.B2i32(v327 == int32(10))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	if v331 != 0 {
		v316 = v323
		v317 = v317 + int32(1)
		v319 = v330
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v309
	F_errcontext_msg(m, int32(_a_F_script_error_callback_1), v11)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	goto L84
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v309
	F_errcontext_msg(m, int32(_a_F_script_error_callback_2), v11+int32(16))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	goto L84
}
func F_secure_loaded_verify_locations(m *base.Module) int32 {
	return int32(0)
}
func F_secure_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_ProcessClientReadInterrupt(m, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if int32(0) < v15 {
		v86 = v15
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_ProcessClientReadInterrupt(m, int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L4:
	;
	if base.Ui32(l2) < base.Ui32(v86) {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	goto L6
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = F_pgmem_recv(m, v23, l1, l2)
	mBase = m.M
	if int32(0) <= v24 {
		v103 = v24
		goto L3
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v27 != 0 {
		v103 = v24
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[0]))
	if v29 != int32(6) {
		v103 = v24
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[1]))
	v34 = int32(0)
	F_ModifyWaitEvent(m, v33, v34, int32(2), v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[1]))
	v44 = F_WaitEventSetWait(m, v40, int32(-1), v8, int32(1), int32(100663296))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v46&int32(16) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v46&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	goto L7
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_secure_read[2]))
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55
	v60 = base.AtomicRmwOr32(m, v55, int32(_a_F_secure_read_0), v55)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	if v64 <= int32(0) {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	F_ProcessClientReadInterrupt(m, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v86 = v64
	goto L4
L22:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_secure_read_1), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_secure_read_2), int32(245), int32(_a_F_secure_read_3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v89 = l2
	goto L28
L27:
	;
	v89 = v86
	goto L28
L28:
	;
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+512))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	base.MemoryCopy(m, l1, v90+v91, v89)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+516))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+516)) = v94 + v89
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+520))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+520)) = v97 - v89
	v103 = v89
	goto L3
L32:
	;
	m.G0 = v8 + int32(16)
	return v103
}
func F_serbian_UTF_8_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = F_SN_new_env(m, int32(36))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v7 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v3)+32)) = uint8(v7)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+28)) = v7
		} else {
		}
		return v3
	}
}
func F_serialize_deflist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	F_initStringInfo(m, v10+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l0 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v137 = F_cstring_to_text_with_len(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = int32(0)
	goto L6
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v30 + v27<<(uint(int32(2))%32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = F_defGetString(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v38 = F_quote_identifier(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v38
	v42 = v10 + int32(16)
	F_appendStringInfo(m, v42, int32(_a_F_serialize_deflist_0), v10)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if base.Ui32(v47-int32(473)) < base.Ui32(int32(2)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v33+int32(4)) < base.Ui32(v113+v114<<(uint(int32(2))%32)) {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	F_appendStringInfoString(m, v42, v35)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v54 = int32(92)
	v55 = F___strchrnul(m, v35, v54)
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v57 == v54 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L11
L16:
	;
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v61 = v55
	goto L19
L18:
	;
	v61 = int32(0)
	goto L19
L19:
	;
	goto L16
L20:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(69))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(39))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v75 = v35
	goto L25
L25:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	F_appendStringInfoChar(m, v10+int32(16), int32(39))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	if base.B2i32(v79 != int32(92))&base.B2i32(v79 != int32(39)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v79))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_appendStringInfoChar(m, v10+int32(16), base.I32_extend8_s(v79))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v75 = v75 + int32(1)
	goto L25
L35:
	;
	goto L11
L36:
	;
	F_appendStringInfoString(m, v10+int32(16), int32(_a_F_serialize_deflist_1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v125 = v27 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v125 < v126 {
		v27 = v125
		goto L6
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L7
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	F_pfree(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	m.G0 = v10 + int32(32)
	return v137
}
func F_setRuleCheckAsUser_Query(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v12 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v40 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = v3
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v21<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = l1
	v31 = v21 + int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v31 < v32 {
		v21 = v31
		goto L4
	} else {
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	goto L5
L7:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v74 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v50 = int32(0)
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v50<<(uint(int32(2))%32))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	if v58 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+36))
	F_setRuleCheckAsUser_Query(m, v61, l1)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v65 = v50 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v65 < v66 {
		v50 = v65
		goto L10
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	goto L11
L18:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v105 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v77 = int32(0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v78 <= v77 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v84 = v77
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v84<<(uint(int32(2))%32))))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+16))
	F_setRuleCheckAsUser_Query(m, v92, l1)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v96 = v84 + int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v96 < v97 {
		v84 = v96
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v112 = F_query_tree_walker_impl(m, l0, int32(1117), v9+int32(12), int32(3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	m.G0 = v9 + int32(16)
	return
L28:
	;
	goto L27
}
func F_setRuleCheckAsUser_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v3 == int32(67) {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_setRuleCheckAsUser_Query(m, l0, v6)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		} else {
			v14 = F_expression_tree_walker_impl(m, l0, int32(1117), l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = v14
				return v17
			}
		}
	} else {
		v17 = int32(0)
		return v17
	}
}
func F_set_attnotnull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+118)))
	if v16 == int32(116) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L50
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L46
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	if v19 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_CheckTableNotInUse(m, l1, int32(_a_F_set_attnotnull_0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v32 = v25 + v26<<(uint(int32(3))%32) + l2*int32(100)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+19)))
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	m.G0 = v13 + int32(32)
	return
L10:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32-int32(72))+86)))
	if v36 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_CacheInvalidateRelcache(m, l1)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L45
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v44 = F_SearchSysCacheCopyAttNum(m, v43, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	if v44 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
	v50 = v48 + v49
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v50)+86)) = uint8(v51)
	F_CatalogTupleUpdate(m, v41, v44+int32(4), v44)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v57 = int32(0)
	if base.B2i32(l0 == v57)|base.B2i32(l3 == v57) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L42
	}
L19:
	;
	v63 = F_palloc0(m, int32(20))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(52)
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+74)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v50)+76))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v50)+96))
	v73 = F_makeVar(m, int32(1), v68, v69, v70, v71, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = int32(-1)
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)) = uint8(v77)
	v79 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v63
	v87 = F_list_make1_impl(m, v79, v13+int32(24))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v90 = F_ConstraintImpliedByRelConstraint(m, l1, v87, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	if v90 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v116 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L27:
	;
	if v94 == int32(0) {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v99 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v50 + v99
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v98 + v99
	F_errmsg_internal(m, int32(_a_F_set_attnotnull_1), v13+int32(16))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_3), int32(_a_F_set_attnotnull_4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	goto L18
L31:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+76)) = uint8(v186)
	goto L18
L32:
	;
	v154 = F_palloc0(m, int32(144))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L39
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v119 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v126 = int32(0)
	goto L35
L35:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v122+v126<<(uint(int32(2))%32))))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v138 == v115 {
		v179 = v137
		goto L31
	} else {
		goto L37
	}
L36:
	;
	goto L32
L37:
	;
	v141 = v126 + int32(1)
	if v119 != v141 {
		v126 = v141
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v115
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+4)) = uint8(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v163 = F_CreateTupleDescCopyConstr(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v163
	*(*int64)(unsafe.Add(mBase, uint32(v154)+88)) = int64(0)
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+84)) = uint8(v168)
	v170 = int32(_a_F_set_attnotnull_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v154)+96)) = uint16(v170)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v173 = F_lappend(m, v172, v154)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v173
	v179 = v154
	goto L31
L42:
	;
	F_relation_close(m, v41, int32(3))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_pfree(m, v44)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	goto L9
L45:
	;
	goto L9
L46:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_set_attnotnull_6), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_7), int32(_a_F_set_attnotnull_8))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg_internal(m, int32(_a_F_set_attnotnull_9), v13)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_set_attnotnull_2), int32(_a_F_set_attnotnull_10), int32(_a_F_set_attnotnull_11))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_set_cheapest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 float64
	_ = v260
	var v261 float64
	_ = v261
	var v264 float64
	_ = v264
	var v265 float64
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v13 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v356
	return
L2:
	;
	v32 = v2
	v33 = v2
	v36 = v2
	v37 = v2
	v39 = v2
	goto L11
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v14 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v352 = v2
	v356 = v2
	v357 = v2
	goto L1
L7:
	;
	return
L8:
	;
	F_errmsg_internal(m, int32(_a_F_set_cheapest_0), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_set_cheapest_1), int32(279), int32(_a_F_set_cheapest_2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v47 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v334 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L13:
	;
	v343 = v39 + int32(1)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v343 < v344 {
		v32 = v334
		v33 = v335
		v36 = v338
		v37 = v339
		v39 = v343
		goto L11
	} else {
		goto L146
	}
L14:
	;
	v48 = F_lappend(m, v37, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v32 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L17:
	;
	if v32 != 0 {
		v334 = v32
		v335 = v33
		v338 = v36
		v339 = v48
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v33 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v334 = int32(0)
	v335 = v46
	v338 = v36
	v339 = v48
	goto L13
L20:
	;
	goto L21
L21:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	if v55 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v57 = v56
	goto L24
L23:
	;
	v57 = v53
	goto L24
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v58 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v60 = v59
	goto L27
L26:
	;
	v60 = v53
	goto L27
L27:
	;
	v61 = int32(0)
	if v57 == v61 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v334 = v61
	v335 = v46
	v338 = v36
	v339 = v48
	goto L13
L29:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	if v156 != v157 {
		goto L65
	} else {
		goto L66
	}
L30:
	;
	switch v155 {
	case 0:
		goto L29
	case 1:
		goto L28
	default:
		v334 = v61
		v335 = v33
		v338 = v36
		v339 = v48
		goto L13
	}
L31:
	;
	v155 = base.B2i32(v60 != int32(0))
	goto L30
L32:
	;
	goto L33
L33:
	;
	if v60 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v155 = int32(2)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v78 < v79 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v81 = v78
	goto L39
L38:
	;
	v81 = v79
	goto L39
L39:
	;
	if v81 <= int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v84 = int32(1)
	goto L42
L41:
	;
	v84 = v81
	goto L42
L42:
	;
	v85 = int32(8)
	v89 = int32(0)
	v91 = v89
	v92 = v89
	goto L45
L43:
	;
	v155 = int32(3)
	goto L30
L44:
	;
	v155 = v142
	goto L30
L45:
	;
	v102 = v92 << (uint(int32(2)) % 32)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v57+v85+v102)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+(v60+v85))))
	if v104&(v106^int32(-1)) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	if v79 < v78 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v128 = v92 + int32(1)
	if v128 != v84 {
		v91 = v126
		v92 = v128
		goto L45
	} else {
		goto L54
	}
L48:
	;
	if base.B2i32(v91 == int32(1))|v106&(v104^int32(-1)) != 0 {
		v142 = int32(3)
		goto L44
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v106&(v104^int32(-1)) == int32(0) {
		v126 = v91
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v126 = int32(2)
	goto L47
L52:
	;
	if v91 == int32(2) {
		goto L43
	} else {
		goto L53
	}
L53:
	;
	v126 = int32(1)
	goto L47
L54:
	;
	goto L46
L55:
	;
	if v126 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if v79 <= v78 {
		v142 = v126
		goto L44
	} else {
		goto L61
	}
L58:
	;
	v135 = int32(3)
	goto L60
L59:
	;
	v135 = int32(2)
	goto L60
L60:
	;
	v155 = v135
	goto L30
L61:
	;
	if v126 == int32(2) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v141 = int32(3)
	goto L64
L63:
	;
	v141 = int32(1)
	goto L64
L64:
	;
	v142 = v141
	goto L44
L65:
	;
	if v157 <= v156 {
		v334 = v61
		v335 = v33
		v338 = v36
		v339 = v48
		goto L13
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v160 = *(*float64)(unsafe.Add(mBase, uint32(v46)+56))
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	if base.F64_lt(v160, v161) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v334 = v61
	v335 = v46
	v338 = v36
	v339 = v48
	goto L13
L69:
	;
	v334 = v61
	v335 = v46
	v338 = v36
	v339 = v48
	goto L13
L70:
	;
	goto L71
L71:
	;
	if base.F64_gt(v160, v161) != 0 {
		v334 = v61
		v335 = v33
		v338 = v36
		v339 = v48
		goto L13
	} else {
		goto L72
	}
L72:
	;
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	if base.F64_lt(v164, v165) == int32(0) {
		v334 = v61
		v335 = v33
		v338 = v36
		v339 = v48
		goto L13
	} else {
		goto L73
	}
L73:
	;
	goto L28
L74:
	;
	v334 = v46
	v335 = v33
	v338 = v46
	v339 = v37
	goto L13
L75:
	;
	goto L76
L76:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v36)+40))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v175 != v176 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if v256 != v257 {
		goto L113
	} else {
		goto L114
	}
L78:
	;
	v253 = v46
	goto L77
L79:
	;
	if v176 <= v175 {
		goto L78
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v36)+48))
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	if base.F64_lt(v179, v180) != 0 {
		v253 = v36
		goto L77
	} else {
		goto L83
	}
L82:
	;
	v253 = v36
	goto L77
L83:
	;
	if base.F64_gt(v179, v180) != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v36)+56))
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v46)+56))
	if base.F64_lt(v183, v184) != 0 {
		v253 = v36
		goto L77
	} else {
		goto L85
	}
L85:
	;
	if base.F64_gt(v183, v184) != 0 {
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	if v187 == v188 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v248 != int32(2) {
		v253 = v36
		goto L77
	} else {
		goto L111
	}
L88:
	;
	v248 = int32(0)
	goto L87
L89:
	;
	goto L90
L90:
	;
	v197 = int32(0)
	goto L93
L91:
	;
	if v237 != 0 {
		goto L108
	} else {
		goto L109
	}
L92:
	;
	v232 = int32(0)
	if v219 != 0 {
		goto L105
	} else {
		goto L106
	}
L93:
	;
	v201 = int32(0)
	if v187 == v201 {
		v211 = v201
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v248 = int32(3)
	goto L87
L95:
	;
	if v188 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v205 <= v197 {
		v211 = int32(0)
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v211 = v207 + v197<<(uint(int32(2))%32)
	goto L95
L98:
	;
	v217 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if base.B2i32(v211 == v217)|base.B2i32(v219 == v217) != 0 {
		goto L92
	} else {
		goto L103
	}
L99:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v197 < v212 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v214 = int32(0)
	v237 = base.B2i32(v211 == v214)
	v239 = v214
	goto L91
L102:
	;
	goto L101
L103:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v219+v197<<(uint(int32(2))%32))))
	if v227 == v229 {
		v197 = v197 + int32(1)
		goto L93
	} else {
		goto L104
	}
L104:
	;
	goto L94
L105:
	;
	v236 = int32(2)
	goto L107
L106:
	;
	v236 = v232
	goto L107
L107:
	;
	v237 = base.B2i32(v211 == v232)
	v239 = v236
	goto L91
L108:
	;
	v241 = v239
	goto L110
L109:
	;
	v241 = int32(1)
	goto L110
L110:
	;
	v248 = v241
	goto L87
L111:
	;
	goto L78
L112:
	;
	v334 = v46
	v335 = v33
	v338 = v253
	v339 = v37
	goto L13
L113:
	;
	if v257 <= v256 {
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v32)+56))
	v261 = *(*float64)(unsafe.Add(mBase, uint32(v46)+56))
	if base.F64_lt(v260, v261) != 0 {
		v334 = v32
		v335 = v33
		v338 = v253
		v339 = v37
		goto L13
	} else {
		goto L117
	}
L116:
	;
	v334 = v32
	v335 = v33
	v338 = v253
	v339 = v37
	goto L13
L117:
	;
	if base.F64_gt(v260, v261) != 0 {
		goto L112
	} else {
		goto L118
	}
L118:
	;
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v32)+48))
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	if base.F64_lt(v264, v265) != 0 {
		v334 = v32
		v335 = v33
		v338 = v253
		v339 = v37
		goto L13
	} else {
		goto L119
	}
L119:
	;
	if base.F64_gt(v264, v265) != 0 {
		goto L112
	} else {
		goto L120
	}
L120:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	if v268 == v269 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v329 != int32(2) {
		v334 = v32
		v335 = v33
		v338 = v253
		v339 = v37
		goto L13
	} else {
		goto L145
	}
L122:
	;
	v329 = int32(0)
	goto L121
L123:
	;
	goto L124
L124:
	;
	v278 = int32(0)
	goto L127
L125:
	;
	if v318 != 0 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	v313 = int32(0)
	if v300 != 0 {
		goto L139
	} else {
		goto L140
	}
L127:
	;
	v282 = int32(0)
	if v268 == v282 {
		v292 = v282
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v329 = int32(3)
	goto L121
L129:
	;
	if v269 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v286 <= v278 {
		v292 = int32(0)
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v292 = v288 + v278<<(uint(int32(2))%32)
	goto L129
L132:
	;
	v298 = int32(0)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	if base.B2i32(v292 == v298)|base.B2i32(v300 == v298) != 0 {
		goto L126
	} else {
		goto L137
	}
L133:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v278 < v293 {
		goto L132
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v295 = int32(0)
	v318 = base.B2i32(v292 == v295)
	v320 = v295
	goto L125
L136:
	;
	goto L135
L137:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v300+v278<<(uint(int32(2))%32))))
	if v308 == v310 {
		v278 = v278 + int32(1)
		goto L127
	} else {
		goto L138
	}
L138:
	;
	goto L128
L139:
	;
	v317 = int32(2)
	goto L141
L140:
	;
	v317 = v313
	goto L141
L141:
	;
	v318 = base.B2i32(v292 == v313)
	v320 = v317
	goto L125
L142:
	;
	v322 = v320
	goto L144
L143:
	;
	v322 = int32(1)
	goto L144
L144:
	;
	v329 = v322
	goto L121
L145:
	;
	goto L112
L146:
	;
	goto L12
L147:
	;
	v352 = v335
	v356 = v338
	v357 = v339
	goto L1
L148:
	;
	goto L149
L149:
	;
	v348 = F_lcons(m, v334, v339)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	v352 = v334
	v356 = v338
	v357 = v348
	goto L1
}
func F_set_errcontext_domain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_set_errcontext_domain[0]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_set_errcontext_domain[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_set_errcontext_domain_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_set_errcontext_domain_1), int32(1609), int32(_a_F_set_errcontext_domain_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l0 != 0 {
			v26 = l0
		} else {
			v26 = int32(_a_F_set_errcontext_domain_3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_c_F_set_errcontext_domain[1]))) = v26
		return
	}
}
func F_set_extra_field(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2
	if v5 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v5 == v9 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v5 == v11 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = l0 + int32(56)
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_pfree(m, v5)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	if v5 == v20 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	if v5 != v22 {
		v16 = v19
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L1
L12:
	;
	return
L13:
	;
	goto L1
}
func F_set_joinrel_size_estimates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
	v9 = F_calc_joinrel_size_estimate(m, l0, l1, l2, l3, v7, v8, l4, l5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(l1)+16)) = v9
		return
	}
}
func F_set_limit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 float64
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	F_getTypeOutputInfo(m, int32(700), v7+int32(12), v7+int32(11))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		v21 = F_OidOutputFunctionCall(m, v20, v9)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_SetConfigOption(m, int32(_a_F_set_limit_0), v21, int32(6), int32(13))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v28 = *(*float64)(unsafe.Add(mBase, _c_F_set_limit[0]))
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_s(base.I32_reinterpret_f32(base.F32_demote_f64(v28)))
			}
		}
	}
}
func F_set_spins_per_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_set_spins_per_delay[0])) = l0
	return
}
func F_setlocale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int64
	_ = v20
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
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
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v401 int64
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v438 int32
	_ = v438
	var v442 int64
	_ = v442
	var v445 int64
	_ = v445
	var v448 int64
	_ = v448
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v539 int64
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if base.Ui32(int32(6)) < base.Ui32(l0) {
		v797 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v797
L2:
	;
	if l0 == int32(6) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v590 = int32(0)
	v591 = int32(_a_F_setlocale_0)
	v596 = v3
	goto L197
L4:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if l1 != 0 {
		goto L140
	} else {
		goto L141
	}
L7:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v26
	v29 = int32(0)
	v30 = l1
	goto L9
L8:
	;
	v797 = int32(0)
	goto L1
L9:
	;
	goto L15
L10:
	;
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[3])) = v442
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[4])) = v445
	v448 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, _c_F_setlocale[5])) = v448
	goto L3
L11:
	;
	v133 = v122 - v30
	if v133 <= int32(23) {
		goto L34
	} else {
		goto L35
	}
L12:
	;
	goto L11
L13:
	;
	v112 = v107
	goto L30
L14:
	;
	v107 = v99
	goto L13
L15:
	;
	if v30&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v45 = v30
	goto L21
L19:
	;
	v59 = v30
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v68 = int32(-2139062144)
	if (int32(16843008)-v65|v65)&v68 != v68 {
		v99 = v59
		goto L14
	} else {
		goto L25
	}
L21:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.B2i32(v50 == int32(0))|base.B2i32(int32(59) == v50) != 0 {
		v122 = v45
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v59 = v56
	goto L20
L23:
	;
	v56 = v45 + int32(1)
	if v56&int32(3) != 0 {
		v45 = v56
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v74 = v59
	v76 = v65
	goto L26
L26:
	;
	v80 = v76 ^ int32(993737531)
	v83 = int32(-2139062144)
	if (int32(16843008)-v80|v80)&v83 != v83 {
		v99 = v74
		goto L14
	} else {
		goto L28
	}
L27:
	;
	v107 = v89
	goto L13
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v89 = v74 + int32(4)
	v93 = int32(-2139062144)
	if (v87|(int32(16843008)-v87))&v93 == v93 {
		v74 = v89
		v76 = v87
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v114 == int32(0) {
		v122 = v112
		goto L12
	} else {
		goto L32
	}
L31:
	;
	v122 = v112
	goto L12
L32:
	;
	if v114 != int32(59) {
		v112 = v112 + int32(1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if base.Ui32(int32(512)) <= base.Ui32(v133) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v312 = v30
	goto L36
L36:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v316 != 0 {
		v330 = v11
		goto L88
	} else {
		goto L89
	}
L37:
	;
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11+v133))) = uint8(v306)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v310 != 0 {
		goto L84
	} else {
		goto L85
	}
L38:
	;
	if v133 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v142 = v11 + v133
	if (v11^v30)&int32(3) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	base.MemoryCopy(m, v11, v30, v133)
	goto L43
L42:
	;
	goto L43
L43:
	;
	goto L37
L44:
	;
	if base.Ui32(v274) < base.Ui32(v142) {
		goto L78
	} else {
		goto L79
	}
L45:
	;
	if v11&int32(3) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v142) < base.Ui32(int32(4)) {
		goto L69
	} else {
		goto L70
	}
L48:
	;
	v178 = v142 & int32(-4)
	if base.Ui32(v142) < base.Ui32(int32(64)) {
		v228 = v172
		v229 = v173
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v172 = v30
	v173 = v11
	goto L48
L50:
	;
	goto L51
L51:
	;
	if v133 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v30
	v173 = v11
	goto L48
L53:
	;
	goto L54
L54:
	;
	v155 = v30
	v156 = v11
	goto L55
L55:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v160)
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v156 + v162
	if v165&int32(3) == int32(0) {
		v172 = v163
		v173 = v165
		goto L48
	} else {
		goto L57
	}
L56:
	;
	v172 = v163
	v173 = v165
	goto L48
L57:
	;
	if base.Ui32(v165) < base.Ui32(v142) {
		v155 = v163
		v156 = v165
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if base.Ui32(v178) <= base.Ui32(v229) {
		v273 = v228
		v274 = v229
		goto L44
	} else {
		goto L65
	}
L60:
	;
	v182 = v178 + int32(-64)
	if base.Ui32(v182) < base.Ui32(v173) {
		v228 = v172
		v229 = v173
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v185 = v172
	v186 = v173
	goto L62
L62:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v185)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v196
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+16)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v185)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v185)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+24)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v185)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+28)) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v185)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+32)) = v206
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v185)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+36)) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v185)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+40)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v185)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+44)) = v212
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v185)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+48)) = v214
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+52)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v185)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+56)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v185)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+60)) = v220
	v222 = int32(-64)
	v223 = v185 - v222
	v225 = v186 - v222
	if base.Ui32(v225) <= base.Ui32(v182) {
		v185 = v223
		v186 = v225
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v228 = v223
	v229 = v225
	goto L59
L64:
	;
	goto L63
L65:
	;
	v235 = v228
	v236 = v229
	goto L66
L66:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v240
	v242 = int32(4)
	v243 = v235 + v242
	v245 = v236 + v242
	if base.Ui32(v245) < base.Ui32(v178) {
		v235 = v243
		v236 = v245
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v273 = v243
	v274 = v245
	goto L44
L68:
	;
	goto L67
L69:
	;
	v273 = v30
	v274 = v11
	goto L44
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v133) < base.Ui32(int32(4)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v273 = v30
	v274 = v11
	goto L44
L73:
	;
	goto L74
L74:
	;
	v254 = v30
	v255 = v11
	goto L75
L75:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v255))) = uint8(v259)
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)) = uint8(v261)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+2)) = uint8(v263)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+3)) = uint8(v265)
	v267 = int32(4)
	v268 = v254 + v267
	v270 = v255 + v267
	if base.Ui32(v270) <= base.Ui32(v142-int32(4)) {
		v254 = v268
		v255 = v270
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v273 = v268
	v274 = v270
	goto L44
L77:
	;
	goto L76
L78:
	;
	v280 = v273
	v281 = v274
	goto L81
L79:
	;
	goto L80
L80:
	;
	goto L37
L81:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v285)
	v287 = int32(1)
	v290 = v281 + v287
	if v290 != v142 {
		v280 = v280 + v287
		v281 = v290
		goto L81
	} else {
		goto L83
	}
L82:
	;
	goto L80
L83:
	;
	goto L82
L84:
	;
	v311 = v122 + int32(1)
	goto L86
L85:
	;
	v311 = v30
	goto L86
L86:
	;
	v312 = v311
	goto L36
L87:
	;
	if v428 == int32(-1) {
		goto L8
	} else {
		goto L137
	}
L88:
	;
	v333 = int32(0)
	goto L103
L89:
	;
	v318 = F_getenv(m, int32(_a_F_setlocale_1))
	mBase = m.M
	if v318 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318))))
	if v319 != 0 {
		v330 = v318
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v324 = F_getenv(m, v29*int32(12)+int32(_a_F_setlocale_2))
	mBase = m.M
	if v324 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v325 != 0 {
		v330 = v324
		goto L88
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v327 = F_getenv(m, int32(_a_F_setlocale_3))
	mBase = m.M
	if v327 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	goto L96
L98:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v328 != 0 {
		v330 = v327
		goto L88
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v330 = int32(_a_F_setlocale_4)
	goto L88
L101:
	;
	goto L100
L102:
	;
	v352 = int32(_a_F_setlocale_4)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	if v353 == int32(46) {
		v360 = v352
		goto L113
	} else {
		goto L114
	}
L103:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v333))))
	v338 = int32(0)
	if base.B2i32(v337 == v338)|base.B2i32(v337 == int32(47)) == v338 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v351 = v333
	goto L102
L105:
	;
	v345 = int32(23)
	v347 = v333 + int32(1)
	if v347 != v345 {
		v333 = v347
		goto L103
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	v351 = v345
	goto L102
L109:
	;
	v428 = v420
	goto L87
L110:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	if v381 != 0 {
		goto L124
	} else {
		goto L125
	}
L111:
	;
	if v29 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L112:
	;
	v366 = F_strcmp(m, v364, int32(_a_F_setlocale_4))
	mBase = m.M
	if v366 == int32(0) {
		v371 = v364
		goto L111
	} else {
		goto L118
	}
L113:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+1)))
	if v361 == int32(0) {
		v371 = v360
		goto L111
	} else {
		goto L117
	}
L114:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330+v351))))
	if v357 != 0 {
		v360 = v352
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if v353 != int32(67) {
		v364 = v330
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v360 = v330
	goto L113
L117:
	;
	v364 = v360
	goto L112
L118:
	;
	v370 = F_strcmp(m, v364, int32(_a_F_setlocale_5))
	mBase = m.M
	if v370 != 0 {
		goto L110
	} else {
		goto L119
	}
L119:
	;
	v371 = v364
	goto L111
L120:
	;
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)))
	if v375 == int32(46) {
		v420 = int32(_a_F_setlocale_6)
		goto L109
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v428 = int32(0)
	goto L87
L123:
	;
	goto L122
L124:
	;
	v384 = v381
	goto L127
L125:
	;
	goto L126
L126:
	;
	v399 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v399 != 0 {
		goto L131
	} else {
		goto L132
	}
L127:
	;
	v389 = F_strcmp(m, v364, v384+int32(8))
	mBase = m.M
	if v389 == int32(0) {
		v420 = v384
		goto L109
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v384)+32))
	if v392 != 0 {
		v384 = v392
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v401 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v401
	v404 = v399 + int32(8)
	v405 = F___memcpy(m, v404, v364, v351)
	mBase = m.M
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v404+v351))) = uint8(v407)
	v409 = int32(_a_F_setlocale_7)
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v399)+32)) = v410
	*(*int32)(unsafe.Add(mBase, _c_F_setlocale[6])) = v399
	goto L133
L132:
	;
	goto L133
L133:
	;
	if v29|v399 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v417 = v399
	goto L136
L135:
	;
	v417 = int32(_a_F_setlocale_6)
	goto L136
L136:
	;
	v420 = v417
	goto L109
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11+int32(24)+v29<<(uint(int32(2))%32)))) = v428
	v438 = v29 + int32(1)
	if v438 != int32(6) {
		v29 = v438
		v30 = v312
		goto L9
	} else {
		goto L138
	}
L138:
	;
	goto L10
L139:
	;
	if v575 != 0 {
		goto L194
	} else {
		goto L195
	}
L140:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v454 != 0 {
		v468 = l1
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L142
L142:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5])))
	v575 = v574
	goto L139
L143:
	;
	if v566 == int32(-1) {
		v797 = v3
		goto L1
	} else {
		goto L193
	}
L144:
	;
	v471 = int32(0)
	goto L159
L145:
	;
	v456 = F_getenv(m, int32(_a_F_setlocale_1))
	mBase = m.M
	if v456 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v457 != 0 {
		v468 = v456
		goto L144
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v462 = F_getenv(m, l0*int32(12)+int32(_a_F_setlocale_2))
	mBase = m.M
	if v462 != 0 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L148
L150:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v463 != 0 {
		v468 = v462
		goto L144
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v465 = F_getenv(m, int32(_a_F_setlocale_3))
	mBase = m.M
	if v465 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	goto L152
L154:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v465))))
	if v466 != 0 {
		v468 = v465
		goto L144
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v468 = int32(_a_F_setlocale_4)
	goto L144
L157:
	;
	goto L156
L158:
	;
	v490 = int32(_a_F_setlocale_4)
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v491 == int32(46) {
		v498 = v490
		goto L169
	} else {
		goto L170
	}
L159:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v471))))
	v476 = int32(0)
	if base.B2i32(v475 == v476)|base.B2i32(v475 == int32(47)) == v476 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v489 = v471
	goto L158
L161:
	;
	v483 = int32(23)
	v485 = v471 + int32(1)
	if v485 != v483 {
		v471 = v485
		goto L159
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	goto L160
L164:
	;
	v489 = v483
	goto L158
L165:
	;
	v566 = v558
	goto L143
L166:
	;
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	if v519 != 0 {
		goto L180
	} else {
		goto L181
	}
L167:
	;
	if l0 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L168:
	;
	v504 = F_strcmp(m, v502, int32(_a_F_setlocale_4))
	mBase = m.M
	if v504 == int32(0) {
		v509 = v502
		goto L167
	} else {
		goto L174
	}
L169:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	if v499 == int32(0) {
		v509 = v498
		goto L167
	} else {
		goto L173
	}
L170:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v489))))
	if v495 != 0 {
		v498 = v490
		goto L169
	} else {
		goto L171
	}
L171:
	;
	if v491 != int32(67) {
		v502 = v468
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v498 = v468
	goto L169
L173:
	;
	v502 = v498
	goto L168
L174:
	;
	v508 = F_strcmp(m, v502, int32(_a_F_setlocale_5))
	mBase = m.M
	if v508 != 0 {
		goto L166
	} else {
		goto L175
	}
L175:
	;
	v509 = v502
	goto L167
L176:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v509)+1)))
	if v513 == int32(46) {
		v558 = int32(_a_F_setlocale_6)
		goto L165
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v566 = int32(0)
	goto L143
L179:
	;
	goto L178
L180:
	;
	v522 = v519
	goto L183
L181:
	;
	goto L182
L182:
	;
	v537 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v537 != 0 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v527 = F_strcmp(m, v502, v522+int32(8))
	mBase = m.M
	if v527 == int32(0) {
		v558 = v522
		goto L165
	} else {
		goto L185
	}
L184:
	;
	goto L182
L185:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v522)+32))
	if v530 != 0 {
		v522 = v530
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v539 = *(*int64)(unsafe.Add(mBase, _c_F_setlocale[7]))
	*(*int64)(unsafe.Add(mBase, uint32(v537))) = v539
	v542 = v537 + int32(8)
	v543 = F___memcpy(m, v542, v502, v489)
	mBase = m.M
	v545 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v542+v489))) = uint8(v545)
	v547 = int32(_a_F_setlocale_7)
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v537)+32)) = v548
	*(*int32)(unsafe.Add(mBase, _c_F_setlocale[6])) = v537
	goto L189
L188:
	;
	goto L189
L189:
	;
	if l0|v537 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v555 = v537
	goto L192
L191:
	;
	v555 = int32(_a_F_setlocale_6)
	goto L192
L192:
	;
	v558 = v555
	goto L165
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5]))) = v566
	v575 = v566
	goto L139
L194:
	;
	v579 = v575 + int32(8)
	goto L196
L195:
	;
	v579 = int32(_a_F_setlocale_8)
	goto L196
L196:
	;
	v797 = v579
	goto L1
L197:
	;
	v599 = *(*int32)(unsafe.Add(mBase, _c_F_setlocale[5]))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v590<<(uint(int32(2))%32))+uint32(_c_F_setlocale[5])))
	if v602 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v788 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v788)
	if v783 != int32(6) {
		goto L250
	} else {
		goto L251
	}
L199:
	;
	v606 = v602 + int32(8)
	goto L201
L200:
	;
	v606 = int32(_a_F_setlocale_8)
	goto L201
L201:
	;
	v607 = F_strlen(m, v606)
	mBase = m.M
	if base.Ui32(int32(512)) <= base.Ui32(v607) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v777 = v591 + v607
	v778 = int32(59)
	*(*uint8)(unsafe.Add(mBase, uint32(v777))) = uint8(v778)
	v780 = int32(1)
	v783 = v596 + base.B2i32(v602 == v599)
	v785 = v590 + v780
	if v785 != int32(6) {
		v590 = v785
		v591 = v777 + v780
		v596 = v783
		goto L197
	} else {
		goto L249
	}
L203:
	;
	if v607 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v614 = v591 + v607
	if (v591^v606)&int32(3) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L206:
	;
	base.MemoryCopy(m, v591, v606, v607)
	goto L208
L207:
	;
	goto L208
L208:
	;
	goto L202
L209:
	;
	if base.Ui32(v746) < base.Ui32(v614) {
		goto L243
	} else {
		goto L244
	}
L210:
	;
	if v591&int32(3) == int32(0) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	if base.Ui32(v614) < base.Ui32(int32(4)) {
		goto L234
	} else {
		goto L235
	}
L213:
	;
	v650 = v614 & int32(-4)
	if base.Ui32(v614) < base.Ui32(int32(64)) {
		v700 = v644
		v701 = v645
		goto L224
	} else {
		goto L225
	}
L214:
	;
	v644 = v606
	v645 = v591
	goto L213
L215:
	;
	goto L216
L216:
	;
	if v607 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v644 = v606
	v645 = v591
	goto L213
L218:
	;
	goto L219
L219:
	;
	v627 = v606
	v628 = v591
	goto L220
L220:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	*(*uint8)(unsafe.Add(mBase, uint32(v628))) = uint8(v632)
	v634 = int32(1)
	v635 = v627 + v634
	v637 = v628 + v634
	if v637&int32(3) == int32(0) {
		v644 = v635
		v645 = v637
		goto L213
	} else {
		goto L222
	}
L221:
	;
	v644 = v635
	v645 = v637
	goto L213
L222:
	;
	if base.Ui32(v637) < base.Ui32(v614) {
		v627 = v635
		v628 = v637
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	if base.Ui32(v650) <= base.Ui32(v701) {
		v745 = v700
		v746 = v701
		goto L209
	} else {
		goto L230
	}
L225:
	;
	v654 = v650 + int32(-64)
	if base.Ui32(v654) < base.Ui32(v645) {
		v700 = v644
		v701 = v645
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v657 = v644
	v658 = v645
	goto L227
L227:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	*(*int32)(unsafe.Add(mBase, uint32(v658))) = v662
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+4)) = v664
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v657)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+8)) = v666
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v657)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+12)) = v668
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v657)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+16)) = v670
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v657)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+20)) = v672
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v657)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+24)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v657)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+28)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v657)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+32)) = v678
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v657)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+36)) = v680
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v657)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+40)) = v682
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v657)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+44)) = v684
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v657)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+48)) = v686
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v657)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+52)) = v688
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v657)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+56)) = v690
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v657)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v658)+60)) = v692
	v694 = int32(-64)
	v695 = v657 - v694
	v697 = v658 - v694
	if base.Ui32(v697) <= base.Ui32(v654) {
		v657 = v695
		v658 = v697
		goto L227
	} else {
		goto L229
	}
L228:
	;
	v700 = v695
	v701 = v697
	goto L224
L229:
	;
	goto L228
L230:
	;
	v707 = v700
	v708 = v701
	goto L231
L231:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	*(*int32)(unsafe.Add(mBase, uint32(v708))) = v712
	v714 = int32(4)
	v715 = v707 + v714
	v717 = v708 + v714
	if base.Ui32(v717) < base.Ui32(v650) {
		v707 = v715
		v708 = v717
		goto L231
	} else {
		goto L233
	}
L232:
	;
	v745 = v715
	v746 = v717
	goto L209
L233:
	;
	goto L232
L234:
	;
	v745 = v606
	v746 = v591
	goto L209
L235:
	;
	goto L236
L236:
	;
	if base.Ui32(v607) < base.Ui32(int32(4)) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v745 = v606
	v746 = v591
	goto L209
L238:
	;
	goto L239
L239:
	;
	v726 = v606
	v727 = v591
	goto L240
L240:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726))))
	*(*uint8)(unsafe.Add(mBase, uint32(v727))) = uint8(v731)
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+1)) = uint8(v733)
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+2)) = uint8(v735)
	v737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v726)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+3)) = uint8(v737)
	v739 = int32(4)
	v740 = v726 + v739
	v742 = v727 + v739
	if base.Ui32(v742) <= base.Ui32(v614-int32(4)) {
		v726 = v740
		v727 = v742
		goto L240
	} else {
		goto L242
	}
L241:
	;
	v745 = v740
	v746 = v742
	goto L209
L242:
	;
	goto L241
L243:
	;
	v752 = v745
	v753 = v746
	goto L246
L244:
	;
	goto L245
L245:
	;
	goto L202
L246:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752))))
	*(*uint8)(unsafe.Add(mBase, uint32(v753))) = uint8(v757)
	v759 = int32(1)
	v762 = v753 + v759
	if v762 != v614 {
		v752 = v752 + v759
		v753 = v762
		goto L246
	} else {
		goto L248
	}
L247:
	;
	goto L245
L248:
	;
	goto L247
L249:
	;
	goto L198
L250:
	;
	v793 = int32(_a_F_setlocale_0)
	goto L252
L251:
	;
	v793 = v606
	goto L252
L252:
	;
	v797 = v793
	goto L1
}
func F_sha224_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(2), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_sha512_bytea(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_cryptohash_internal(m, int32(5), v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_shell_archive_configured(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[2])) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_shell_archive_configured_0)
		v20 = F_format_elog_string(m, int32(_a_F_shell_archive_configured_1), v5)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_shell_archive_configured[3])) = v20
			m.G0 = v5 + int32(16)
			return base.B2i32(v9 != int32(0))
		}
	} else {
		m.G0 = v5 + int32(16)
		return base.B2i32(v9 != int32(0))
	}
}
func F_show_instrumentation_count(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 float64
	_ = v13
	var v18 int32
	_ = v18
	var v20 float64
	_ = v20
	var v25 int32
	_ = v25
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v36 int32
	_ = v36
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)))
	if v7 != int32(1) {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
		if v10 == int32(0) {
			return
		} else {
			v13 = *(*float64)(unsafe.Add(mBase, uint32(v10)+416))
			if l1 == int32(2) {
				v18 = int32(432)
			} else {
				v18 = int32(424)
			}
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v10+v18)))
			if base.F64_gt(v20, float64(0)) == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
				if v25 == int32(0) {
					return
				} else {
					v30 = float64(0)
					if base.F64_gt(v13, v30) != 0 {
						v33 = base.F64_div(v20, v13)
					} else {
						v33 = v30
					}
					F_ExplainPropertyFloat(m, l0, int32(0), v33, int32(0), l3)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v30 = float64(0)
				if base.F64_gt(v13, v30) != 0 {
					v33 = base.F64_div(v20, v13)
				} else {
					v33 = v30
				}
				F_ExplainPropertyFloat(m, l0, int32(0), v33, int32(0), l3)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_show_trgm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v145 int32
	_ = v145
	var v164 int32
	_ = v164
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v26 = v24 & v22
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v22
	goto L5
L4:
	;
	v27 = int32(4)
	goto L5
L5:
	;
	if v24 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v56 = F_generate_trgm(m, v18+v27, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v34 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v45 = int32(1)
	if v26 != 0 {
		v55 = int32(base.Ui32(v24)>>(uint(v45)%32)) - v45
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v37 = int32(16)
	goto L12
L11:
	;
	v37 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = int32(4)
	goto L15
L14:
	;
	v44 = v37
	goto L15
L15:
	;
	v55 = v44
	goto L6
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v64 = base.I32_div_u_s(int32(base.Ui32(v58)>>(uint(int32(2))%32))-int32(5), int32(3))
	v67 = F_palloc_mul(m, int32(8), v64+int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v69)>>(uint(int32(2))%32))-int32(5)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v79 = v56 + int32(5)
	v88 = v2
	goto L22
L20:
	;
	v237 = v2
	goto L21
L21:
	;
	v244 = F_construct_array_builtin(m, v67, v237, int32(25))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L46
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91*int32(28))+uint32(_c_F_show_trgm[1])))
	goto L24
L23:
	;
	v237 = v230
	goto L21
L24:
	;
	if int32(4) <= v96 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[0]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101*int32(28))+uint32(_c_F_show_trgm[1])))
	goto L28
L26:
	;
	v112 = int32(16)
	goto L27
L27:
	;
	v113 = F_palloc(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v112 = v106*int32(3) + int32(4)
	goto L27
L29:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_show_trgm[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117*int32(28))+uint32(_c_F_show_trgm[1])))
	goto L32
L30:
	;
	v215 = int32(3)
	*(*int64)(unsafe.Add(mBase, uint32(v67+v88<<(uint(v215)%32)))) = base.I64_extend_i32_u(v113)
	v223 = v88 + int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v230 = base.I32_div_u_s(int32(base.Ui32(v224)>>(uint(int32(2))%32))-int32(5), v215)
	if base.Ui32(v223) < base.Ui32(v230) {
		v79 = v79 + v215
		v88 = v223
		goto L22
	} else {
		goto L45
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(28)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+4)) = uint8(v207)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+5)) = uint8(v209)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+6)) = uint8(v211)
	goto L30
L32:
	;
	if v122 < int32(2) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v126 = base.I32_extend8_s(v125)
	if v126 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+2)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v182 | (v183<<(uint(int32(8))%32) | v125<<(uint(int32(16))%32))
	v192 = v113 + int32(4)
	v195 = F_pg_snprintf(m, v192, int32(12), int32(_a_F_show_trgm_0), v14)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L44
	}
L35:
	;
	goto L36
L36:
	;
	if base.B2i32(base.B2i32(base.Ui32(v125-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v125|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v126 != int32(32)) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v145 < int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L39
L39:
	;
	if base.B2i32(base.B2i32(base.Ui32(v145-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v145|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v145 != int32(32)) != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+2)))
	if v164 < int32(0) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	if base.B2i32(base.Ui32(v164-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v164|int32(32)-int32(97)) < base.Ui32(int32(26)))|base.B2i32(v164 == int32(32)) != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	goto L34
L44:
	;
	v197 = F_strlen(m, v192)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v197<<(uint(int32(2))%32) + int32(16)
	goto L30
L45:
	;
	goto L23
L46:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if base.Ui32(int32(3)) <= base.Ui32(int32(base.Ui32(v246)>>(uint(int32(2))%32))-int32(5)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v262 = v2
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_pfree(m, v67)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L54
	}
L50:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v67+v262<<(uint(int32(3))%32))))
	F_pfree(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	goto L49
L52:
	;
	v271 = v262 + int32(1)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v278 = base.I32_div_u_s(int32(base.Ui32(v272)>>(uint(int32(2))%32))-int32(5), int32(3))
	if base.Ui32(v271) < base.Ui32(v278) {
		v262 = v271
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	F_pfree(m, v56)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v295 != v18 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_pfree(m, v18)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_u(v244)
L59:
	;
	goto L58
}
func F_sigUsr1Handler(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_sigUsr1Handler_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_sigUsr1Handler_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_sigUsr1Handler[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_sigfillset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-15032385537)
	return
}
func F_significant_digits(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	v7 = l0
	goto L1
L1:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v14 = v12 - int32(43)
	v21 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v14))|base.B2i32(int32(1)<<(uint(v14)%32)&int32(37) == v21) == v21 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v28 = v7
	v29 = v12
	v31 = int32(1)
	goto L9
L3:
	;
	v7 = v7 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return v74
L7:
	;
	v74 = v31
	goto L6
L8:
	;
	v44 = v28
	v45 = v29
	v46 = int32(0)
	goto L15
L9:
	;
	switch v29 - int32(46) {
	case 0:
		v38 = v31
		goto L12
	case 1:
		goto L8
	case 2:
		goto L13
	default:
		goto L11
	}
L10:
	;
	if v29 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L11:
	;
	goto L10
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	v28 = v28 + int32(1)
	v29 = v39
	v31 = v38
	goto L9
L13:
	;
	v38 = v31 + int32(1)
	goto L12
L14:
	;
	goto L8
L15:
	;
	v50 = base.B2i32(v45 != int32(46))
	if v50&base.B2i32(base.Ui32(int32(9)) < base.Ui32((v45-int32(48))&int32(255))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v66 != 0 {
		v74 = v66
		goto L6
	} else {
		goto L21
	}
L17:
	;
	v60 = v46 + v50
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v61 != 0 {
		v44 = v44 + int32(1)
		v45 = v61
		v46 = v60
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v66 = v46
	goto L19
L19:
	;
	goto L16
L20:
	;
	v66 = v60
	goto L19
L21:
	;
	goto L7
}
func F_sjis_to_euc_jp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int64
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	F_check_encoding_conversion_args(m, v16, v17, v18, int32(35), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v18 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v328)
	return base.I64_extend_i32_s(v319 - v15)
L4:
	;
	v317 = v14
	v319 = v15
	goto L3
L5:
	;
	goto L6
L6:
	;
	v28 = v14
	v30 = v15
	v31 = v18
	goto L7
L7:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v40 = base.I32_extend8_s(v39)
	if int32(0) <= v40 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v317 = v302
	v319 = v313
	goto L3
L9:
	;
	if int32(0) < v305 {
		v28 = v302
		v30 = v313
		v31 = v305
		goto L7
	} else {
		goto L70
	}
L10:
	;
	if v40 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v58 = F_pg_encoding_verifymbchar(m, int32(35), v30, v31)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v13 != int64(0) {
		v317 = v28
		v319 = v30
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v40)
	v51 = int32(1)
	v302 = v28 + v51
	v305 = v31 - v51
	v313 = v30 + v51
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(35), v30, v31)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v58 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v13 != int64(0) {
		v317 = v28
		v319 = v30
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32((v40+int32(95))&int32(255)) <= base.Ui32(int32(62)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_report_invalid_encoding(m, int32(35), v30, v31)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v302 = v288
	v305 = v31 - v58
	v313 = v30 + v58
	goto L9
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v40)
	v74 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v74)
	v288 = v28 + int32(2)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	v82 = v79 | v39<<(uint(int32(8))%32)
	if base.Ui32(v82-int32(_a_F_sjis_to_euc_jp_0)) <= base.Ui32(int32(767)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v87 = v82
	v89 = v79
	v92 = v39
	v93 = int32(0)
	goto L31
L29:
	;
	v129 = v82
	v131 = v79
	v134 = v39
	goto L30
L30:
	;
	if v129 <= int32(_a_F_sjis_to_euc_jp_1) {
		goto L40
	} else {
		goto L41
	}
L31:
	;
	v100 = v93 << (uint(int32(3)) % 32)
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_sjis_to_euc_jp[0]))))
	if v103 == v87 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v129 = v122
	v131 = v123
	v134 = v124
	goto L30
L33:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_sjis_to_euc_jp[1]))))
	v110 = v105
	v111 = v105 & int32(255)
	v112 = int32(base.Ui32(v105) >> (uint(int32(8)) % 32))
	goto L35
L34:
	;
	v110 = v87
	v111 = v89
	v112 = v92
	goto L35
L35:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_sjis_to_euc_jp[2]))))
	if v115 == v110 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_sjis_to_euc_jp[3]))))
	v122 = v117
	v123 = v117 & int32(255)
	v124 = int32(base.Ui32(v117) >> (uint(int32(8)) % 32))
	goto L38
L37:
	;
	v122 = v110
	v123 = v111
	v124 = v112
	goto L38
L38:
	;
	v126 = v93 + int32(2)
	if v126 != int32(388) {
		v87 = v122
		v89 = v123
		v92 = v124
		v93 = v126
		goto L31
	} else {
		goto L39
	}
L39:
	;
	goto L32
L40:
	;
	v146 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v131))
	if base.Ui32(int32(158)) < base.Ui32(v131) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v168 = int32(0)
	if base.B2i32(base.B2i32(v129 != int32(_a_F_sjis_to_euc_jp_2))&base.B2i32(base.Ui32(v129) < base.Ui32(int32(_a_F_sjis_to_euc_jp_3))) == v168)&base.B2i32(base.Ui32(int32(176)) < base.Ui32(v129-int32(_a_F_sjis_to_euc_jp_4))) == v168 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v147 = int32(2)
	goto L45
L44:
	;
	v147 = int32(96)
	goto L45
L45:
	;
	v151 = v147 + v131 + base.B2i32(base.Ui32(v131) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v151)
	v159 = v134<<(uint(int32(1))%32)&int32(126) | v146 + int32(159)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v159)
	v288 = v28 + int32(2)
	goto L24
L46:
	;
	v177 = int32(_a_F_sjis_to_euc_jp_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v28))) = uint16(v177)
	v288 = v28 + int32(2)
	goto L24
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(v129-int32(_a_F_sjis_to_euc_jp_3)) <= base.Ui32(int32(1279)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v188 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v131))
	if base.Ui32(int32(158)) < base.Ui32(v131) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.Ui32(v129-int32(_a_F_sjis_to_euc_jp_6)) <= base.Ui32(int32(1279)) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v189 = int32(2)
	goto L54
L53:
	;
	v189 = int32(96)
	goto L54
L54:
	;
	v193 = v189 + v131 + base.B2i32(base.Ui32(v131) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v193)
	v203 = (v134<<(uint(int32(1))%32)+int32(34))&int32(126) | v188 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v203)
	v288 = v28 + int32(2)
	goto L24
L55:
	;
	v211 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v211)
	v216 = base.B2i32(base.Ui32(int32(158)) < base.Ui32(v131))
	if base.Ui32(int32(158)) < base.Ui32(v131) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v129) < base.Ui32(int32(_a_F_sjis_to_euc_jp_7)) {
		v288 = v28
		goto L24
	} else {
		goto L61
	}
L58:
	;
	v217 = int32(2)
	goto L60
L59:
	;
	v217 = int32(96)
	goto L60
L60:
	;
	v221 = v217 + v131 + base.B2i32(base.Ui32(v131) < base.Ui32(int32(128)))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+2)) = uint8(v221)
	v231 = (v134<<(uint(int32(1))%32)+int32(24))&int32(126) | v216 + int32(243)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v231)
	v288 = v28 + int32(3)
	goto L24
L61:
	;
	v238 = v129
	v239 = v28
	v244 = int32(0)
	goto L62
L62:
	;
	v251 = v244 << (uint(int32(3)) % 32)
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+uint32(_c_F_sjis_to_euc_jp[1]))))
	if v238 != v254 {
		v281 = v238
		v282 = v239
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v288 = v282
	goto L24
L64:
	;
	v284 = v244 + int32(1)
	if v284 != int32(388) {
		v238 = v281
		v239 = v282
		v244 = v284
		goto L62
	} else {
		goto L69
	}
L65:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)+uint32(_c_F_sjis_to_euc_jp[4])))
	if int32(_a_F_sjis_to_euc_jp_8) <= v256 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v259 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v259)
	v261 = int32(128)
	v262 = v256 | v261
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+2)) = uint8(v262)
	v267 = int32(base.Ui32(v256)>>(uint(int32(8))%32)) | v261
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)) = uint8(v267)
	v281 = v256
	v282 = v239 + int32(3)
	goto L64
L67:
	;
	goto L68
L68:
	;
	v271 = int32(128)
	v272 = v256 | v271
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)) = uint8(v272)
	v277 = int32(base.Ui32(v256)>>(uint(int32(8))%32)) | v271
	*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v277)
	v281 = v256
	v282 = v239 + int32(2)
	goto L64
L69:
	;
	goto L63
L70:
	;
	goto L8
}
func F_slice_del(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	v8 = int32(-1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v9 < int32(0) {
		v51 = v8
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 < v9 {
			v51 = v8
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v14 < v12 {
				v51 = v8
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16-int32(4))))
				if v19 < v14 {
					v51 = v8
				} else {
					if v9 == v12 {
						v44 = v9
					} else {
						v22 = v19 - v12
						if v22 != 0 {
							base.MemoryCopy(m, v9+v16, v12+v16, v22)
						} else {
						}
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v29 = v12 - v9
						*(*int32)(unsafe.Add(mBase, uint32(v26-int32(4)))) = v19 - v29
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v32 - v29
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v36 <= v35 {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v35 - v29
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v44 = v40
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v35 <= v41 {
								v44 = v41
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41
								v44 = v41
							}
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
					v51 = int32(0)
				}
			}
		}
	}
	return v51
}
func F_sort_pending_writebacks_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return l0
L2:
	;
	return l2
L3:
	;
	return v101
L4:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v71) {
		goto L36
	} else {
		goto L37
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if base.Ui32(v18) < base.Ui32(v34) {
		v101 = l1
		goto L3
	} else {
		goto L15
	}
L6:
	;
	if base.Ui32(v18) < base.Ui32(v17) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.Ui32(v15) < base.Ui32(v13) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(v13) < base.Ui32(v15) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if base.Ui32(v16) < base.Ui32(v14) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if base.Ui32(v14) < base.Ui32(v16) {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v25 < v26 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v26 < v25 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v30) <= base.Ui32(v29) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L5
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v34) < base.Ui32(v18) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if base.Ui32(v17) < base.Ui32(v34) {
		goto L2
	} else {
		goto L25
	}
L17:
	;
	if base.Ui32(v13) < base.Ui32(v36) {
		v101 = l1
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v36) < base.Ui32(v13) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v14) < base.Ui32(v37) {
		v101 = l1
		goto L3
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(v37) < base.Ui32(v14) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v43 < v44 {
		v101 = l1
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if v44 < v43 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v47) < base.Ui32(v48) {
		v101 = l1
		goto L3
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	if base.Ui32(v34) < base.Ui32(v17) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(v15) < base.Ui32(v36) {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v36) < base.Ui32(v15) {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v16) < base.Ui32(v37) {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	if base.Ui32(v37) < base.Ui32(v16) {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v58 < v59 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	if v59 < v58 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v62) < base.Ui32(v63) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v65 = l2
	goto L35
L34:
	;
	v65 = l0
	goto L35
L35:
	;
	return v65
L36:
	;
	if base.Ui32(v17) < base.Ui32(v71) {
		goto L1
	} else {
		goto L46
	}
L37:
	;
	if base.Ui32(v71) < base.Ui32(v18) {
		v101 = l1
		goto L3
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v13) < base.Ui32(v69) {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(v69) < base.Ui32(v13) {
		v101 = l1
		goto L3
	} else {
		goto L40
	}
L40:
	;
	if base.Ui32(v14) < base.Ui32(v70) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(v70) < base.Ui32(v14) {
		v101 = l1
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v78 < v79 {
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v79 < v78 {
		v101 = l1
		goto L3
	} else {
		goto L44
	}
L44:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v83) < base.Ui32(v82) {
		v101 = l1
		goto L3
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	if base.Ui32(v71) < base.Ui32(v17) {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32(v15) < base.Ui32(v69) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(v69) < base.Ui32(v15) {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	if base.Ui32(v16) < base.Ui32(v70) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if base.Ui32(v70) < base.Ui32(v16) {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v93 < v94 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v94 < v93 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if base.Ui32(v97) < base.Ui32(v98) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v100 = l0
	goto L56
L55:
	;
	v100 = l2
	goto L56
L56:
	;
	v101 = v100
	goto L3
}
func F_spcache_init(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0]))
	if v4 == int32(0) {
		v17 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
		*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
		if v29 == v20 {
			v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
			v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
				v46 = v39
				v48 = F_MemoryContextAllocZero(m, v46, int32(32))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
					v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
						v63 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
						return
					}
				}
			}
		} else {
			F_MemoryContextReset(m, v29)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
				v46 = v45
				v48 = F_MemoryContextAllocZero(m, v46, int32(32))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
					v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
						*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
						v63 = int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
						return
					}
				}
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])))
		if v8&int32(1) == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
			*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
			*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
			if v29 == v20 {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
				v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
					v46 = v39
					v48 = F_MemoryContextAllocZero(m, v46, int32(32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
						v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
							return
						}
					}
				}
			} else {
				F_MemoryContextReset(m, v29)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
					v46 = v45
					v48 = F_MemoryContextAllocZero(m, v46, int32(32))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
						v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
							*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
							v63 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
							return
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			if base.Ui32(v13) < base.Ui32(int32(256)) {
				return
			} else {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[1])) = uint8(v17)
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v20
				*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[3])) = v20
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
				if v29 == v20 {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[5]))
					v39 = F_AllocSetContextCreateInternal(m, v34, int32(_a_F_spcache_init_0), int32(0), int32(_a_F_spcache_init_1), int32(_a_F_spcache_init_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4])) = v39
						v46 = v39
						v48 = F_MemoryContextAllocZero(m, v46, int32(32))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
							v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
								return
							}
						}
					}
				} else {
					F_MemoryContextReset(m, v29)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_spcache_init[4]))
						v46 = v45
						v48 = F_MemoryContextAllocZero(m, v46, int32(32))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v46
							v55 = F_MemoryContextAllocExtended(m, v46, int32(768), int32(5))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+12)) = int64(120259084319)
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = int64(32)
								*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v55
								v63 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_spcache_init[2])) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, _c_F_spcache_init[0])) = v48
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_spghandler(m *base.Module, l0 int32) int64 {
	return int64(831140)
}
func F_spgrescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v407 int32
	_ = v407
	var v408 int64
	_ = v408
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l1 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if l3 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v14 * int32(56)
	if v18 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v21, l1, v18)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+112)) = v142
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v151 {
		goto L32
	} else {
		goto L33
	}
L6:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+108)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+116)) = v137
	v142 = int32(0)
	v146 = v135
	goto L5
L7:
	;
	if v24 <= int32(0) {
		v134 = v24
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v62 = v24
	goto L9
L9:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+108)) = v62
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+116)) = v73
	v75 = int32(0)
	if v62 <= v75 {
		v142 = v75
		v146 = v71
		goto L5
	} else {
		goto L20
	}
L10:
	;
	v28 = v24 * int32(56)
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v29, l3, v28)
	goto L13
L12:
	;
	goto L13
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v31 <= int32(0) {
		v134 = v31
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(0)
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v38*int32(56))+20))
	v50 = F_get_func_rettype(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v62 = v59
	goto L9
L17:
	;
	return
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v38<<(uint(int32(2))%32)))) = v50
	v58 = v38 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v58 < v59 {
		v38 = v58
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v78 <= int32(0) {
		v142 = v75
		v146 = v71
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v84 = v75
	v85 = int32(0)
	goto L22
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v71)+116))
	v96 = v93 + v85*int32(56)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v97&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v142 = v123
	v146 = v71
	goto L5
L24:
	;
	if v84 != v85 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v123 = v84
	v124 = int32(-1)
	goto L26
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v71)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v85<<(uint(int32(2))%32)))) = v124
	v131 = v85 + int32(1)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v131 < v132 {
		v84 = v123
		v85 = v131
		goto L22
	} else {
		goto L30
	}
L27:
	;
	v105 = v93 + v84*int32(56)
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v96)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+48)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v96)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+40)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v96)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+32)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+24)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+16)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v118
	goto L29
L28:
	;
	goto L29
L29:
	;
	v123 = v84 + int32(1)
	v124 = v84
	goto L26
L30:
	;
	goto L23
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+100)) = v224
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	F_MemoryContextReset(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L17
	} else {
		goto L48
	}
L32:
	;
	v154 = int32(0)
	v158 = v154
	v159 = v154
	v160 = v151
	v164 = v6
	v165 = v6
	goto L36
L33:
	;
	goto L34
L34:
	;
	v219 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+96)) = uint16(v219)
	v224 = int32(0)
	goto L31
L35:
	;
	v216 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v146)+96)) = uint16(v216)
	v224 = v216
	goto L31
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v169 = v166 + v159*int32(56)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v170&int32(64) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	if v203&v204 != 0 {
		goto L35
	} else {
		goto L47
	}
L38:
	;
	v206 = v159 + int32(1)
	if v206 < v202 {
		v158 = v201
		v159 = v206
		v160 = v202
		v164 = v203
		v165 = v204
		goto L36
	} else {
		goto L46
	}
L39:
	;
	v201 = v158
	v202 = v160
	v203 = v164
	v204 = int32(1)
	goto L38
L40:
	;
	goto L41
L41:
	;
	if v170&int32(128) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v201 = v158
	v202 = v160
	v203 = int32(1)
	v204 = v165
	goto L38
L43:
	;
	goto L44
L44:
	;
	if v170&int32(1) != 0 {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v146)+104))
	v182 = v179 + v158*int32(56)
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v169)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+48)) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v169)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+40)) = v185
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v169)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+32)) = v187
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v169)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+24)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v169)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+16)) = v191
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v169)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+8)) = v193
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v169)))
	*(*int64)(unsafe.Add(mBase, uint32(v182))) = v195
	v197 = int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v201 = v158 + v197
	v202 = v200
	v203 = v197
	v204 = v165
	goto L38
L46:
	;
	goto L37
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+97)) = uint8(v203)
	*(*uint8)(unsafe.Add(mBase, uint32(v146)+96)) = uint8(v204)
	v224 = v201
	goto L31
L48:
	;
	v236 = int32(_a_F_spgrescan_0)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v11)+92))
	*(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0])) = v239
	v242 = F_pairingheap_allocate(m, int32(264), v11)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v242
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+96)))
	if v245 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v249 = F_palloc(m, int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+97)))
	if v265 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v251 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v249)+40)) = uint16(v251)
	*(*int32)(unsafe.Add(mBase, uint32(v249)+42)) = v251
	v255 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v249)+16)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v249)+24)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v249)+32)) = int64(562949953421312)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v261, v249)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+188))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	v274 = F_palloc(m, v269<<(uint(int32(3))%32)+int32(48))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_spgrescan[0])) = v237
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	if v309 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v276 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+42)) = uint8(v276)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v11)+112))
	if v278 <= v276 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+45)) = uint8(v289)
	*(*uint16)(unsafe.Add(mBase, uint32(v274)+43)) = uint16(v289)
	v293 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v274)+40)) = uint16(v293)
	v295 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v274)+16)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v274)+24)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v274)+32)) = int64(281474976710656)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v11)+84))
	F_pairingheap_add(m, v301, v274)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L17
	} else {
		goto L62
	}
L60:
	;
	v282 = v278 << (uint(int32(3)) % 32)
	if v282 == int32(0) {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	base.MemoryCopy(m, v274+int32(48), v268, v282)
	goto L59
L62:
	;
	goto L57
L63:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+208)))
	if v349 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v312 <= int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v319 = v312
	v321 = int32(0)
	goto L66
L66:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(_a_F_spgrescan_1)+v321<<(uint(int32(2))%32))))
	if v331 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L63
L68:
	;
	F_pfree(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L17
	} else {
		goto L71
	}
L69:
	;
	v335 = v319
	goto L70
L70:
	;
	v337 = v321 + int32(1)
	if v337 < v335 {
		v319 = v335
		v321 = v337
		goto L66
	} else {
		goto L72
	}
L71:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	v335 = v334
	goto L70
L72:
	;
	goto L67
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+216)) = int64(0)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+272))
	if v391 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L74:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v352 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v361 = int32(0)
	goto L76
L76:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(3488)+v361<<(uint(int32(2))%32))))
	F_pfree(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L17
	} else {
		goto L78
	}
L77:
	;
	goto L73
L78:
	;
	v375 = v361 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	if v375 < v376 {
		v361 = v375
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v407 != 0 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+268)))
	if v394 != int32(1) {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	v401 = v391
	goto L83
L83:
	;
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v401)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v401)+16)) = v402 + int64(1)
	goto L80
L84:
	;
	F_pgstat_assoc_relation(m, v390)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+272))
	v401 = v400
	goto L83
L86:
	;
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v407)))
	*(*int64)(unsafe.Add(mBase, uint32(v407))) = v408 + int64(1)
	goto L88
L87:
	;
	goto L88
L88:
	;
	return
}
func F_split_pathtarget_at_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	var v8 int32
	_ = v8
	F_split_pathtarget_at_srfs_extended(m, l0, l1, l2, l3, l4, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_split_pathtarget_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v128 = F_palloc(m, int32(8))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L43
	}
L2:
	;
	v54 = F_palloc(m, int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L25
	}
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v9 != int32(1) {
		v30 = l0
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v51 = int32(0)
	goto L5
L5:
	;
	return v51
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v32 = F_list_member(m, v31, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L13
	}
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+45)))
	if v14 != int32(1) {
		v30 = l0
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v17 == int32(0) {
		v30 = l0
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+340))
	v21 = F_bms_make_singleton(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v26 = F_remove_nulling_relids(m, l0, v21, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = v26
	goto L6
L13:
	;
	if v32 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	goto L1
L15:
	;
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v34 - int32(6) {
	case 0, 3, 4, 5:
		goto L1
	case 1, 2, 6, 7, 8, 10:
		goto L17
	case 9:
		goto L18
	case 11:
		goto L19
	default:
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	v46 = F_expression_tree_walker_impl(m, l0, int32(946), l1)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L24
	}
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v42 != 0 {
		goto L2
	} else {
		goto L23
	}
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v39 != int32(1) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	if v34 != int32(321) {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L1
L22:
	;
	goto L2
L23:
	;
	goto L17
L24:
	;
	v51 = v46
	goto L5
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = l0
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v62
	v67 = F_expression_tree_walker_impl(m, l0, int32(946), l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = v70
	goto L29
L28:
	;
	v72 = int32(0)
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v75 = v73 + int32(1)
	if v72 <= v75 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v78 = F_lappend(m, v69, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	v92 = v69
	goto L32
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v95 = v75 << (uint(int32(2)) % 32)
	v96 = v93 + v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = F_lappend(m, v97, v54)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L10
	} else {
		goto L36
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v83 = F_lappend(m, v81, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v88 = F_lappend(m, v86, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v92 = v91
	goto L32
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v98
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+12))
	v103 = v102 + v95
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v106 = F_list_concat(m, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v106
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v111 = v110 + v95
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v114 = F_list_concat(m, v112, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v58
	v118 = F_lappend(m, v57, v54)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	if v75 < v56 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v121 = v56
	goto L42
L41:
	;
	v121 = v75
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v118
	return int32(0)
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = l0
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v134 = F_lappend(m, v133, v128)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v134
	return int32(0)
}
func F_sqrt_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v241 int64
	_ = v241
	var v252 int64
	_ = v252
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v297 int64
	_ = v297
	var v308 int64
	_ = v308
	var v313 int64
	_ = v313
	var v321 int32
	_ = v321
	var v334 int64
	_ = v334
	var v346 int64
	_ = v346
	var v348 int64
	_ = v348
	var v373 int64
	_ = v373
	var v384 int64
	_ = v384
	var v387 int64
	_ = v387
	var v391 int64
	_ = v391
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v430 int64
	_ = v430
	var v434 int64
	_ = v434
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v461 int64
	_ = v461
	var v465 int64
	_ = v465
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v486 int64
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int64
	_ = v556
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v585 int64
	_ = v585
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v603 int64
	_ = v603
	var v605 int64
	_ = v605
	var v607 int32
	_ = v607
	var v613 int64
	_ = v613
	var v619 int32
	_ = v619
	var v636 int64
	_ = v636
	var v641 int64
	_ = v641
	var v642 int64
	_ = v642
	var v645 int64
	_ = v645
	var v650 int64
	_ = v650
	var v652 int64
	_ = v652
	var v654 int64
	_ = v654
	var v656 int32
	_ = v656
	var v661 int64
	_ = v661
	var v663 int64
	_ = v663
	var v665 int64
	_ = v665
	var v667 int32
	_ = v667
	var v669 int64
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v683 int64
	_ = v683
	var v689 int32
	_ = v689
	var v690 int64
	_ = v690
	var v696 int32
	_ = v696
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v721 int64
	_ = v721
	var v723 int64
	_ = v723
	var v733 int32
	_ = v733
	var v735 int64
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int64
	_ = v743
	var v744 int64
	_ = v744
	var v745 int64
	_ = v745
	var v748 int64
	_ = v748
	var v764 int64
	_ = v764
	var v765 int64
	_ = v765
	var v766 int64
	_ = v766
	var v767 int64
	_ = v767
	var v780 int32
	_ = v780
	var v781 int64
	_ = v781
	var v782 int64
	_ = v782
	var v783 int64
	_ = v783
	var v784 int64
	_ = v784
	var v789 int64
	_ = v789
	var v792 int64
	_ = v792
	var v795 int64
	_ = v795
	var v798 int64
	_ = v798
	var v799 int64
	_ = v799
	var v803 int64
	_ = v803
	var v810 int64
	_ = v810
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v832 int64
	_ = v832
	var v835 int64
	_ = v835
	var v837 int32
	_ = v837
	var v840 int64
	_ = v840
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v888 int64
	_ = v888
	var v894 int32
	_ = v894
	var v895 int64
	_ = v895
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v929 int64
	_ = v929
	var v932 int64
	_ = v932
	var v937 int32
	_ = v937
	var v939 int64
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int64
	_ = v947
	var v948 int64
	_ = v948
	var v949 int64
	_ = v949
	var v952 int64
	_ = v952
	var v968 int64
	_ = v968
	var v969 int64
	_ = v969
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v984 int32
	_ = v984
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v987 int64
	_ = v987
	var v988 int64
	_ = v988
	var v993 int64
	_ = v993
	var v996 int64
	_ = v996
	var v999 int64
	_ = v999
	var v1002 int64
	_ = v1002
	var v1003 int64
	_ = v1003
	var v1007 int64
	_ = v1007
	var v1014 int64
	_ = v1014
	var v1026 int32
	_ = v1026
	var v1027 int64
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int64
	_ = v1036
	var v1039 int64
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int64
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1081 int32
	_ = v1081
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1234 int32
	_ = v1234
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1294 int32
	_ = v1294
	var v1295 int64
	_ = v1295
	var v1313 int32
	_ = v1313
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1358 int32
	_ = v1358
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1392 int32
	_ = v1392
	var v1399 int32
	_ = v1399
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1445 int32
	_ = v1445
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1505 int32
	_ = v1505
	var v1506 int64
	_ = v1506
	var v1522 int32
	_ = v1522
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1565 int64
	_ = v1565
	var v1567 int64
	_ = v1567
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1588 int64
	_ = v1588
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1606 int32
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1716 int32
	_ = v1716
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1757 int32
	_ = v1757
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1780 int32
	_ = v1780
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1847 int32
	_ = v1847
	var v1854 int32
	_ = v1854
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1897 int32
	_ = v1897
	var v1908 int32
	_ = v1908
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1956 int32
	_ = v1956
	var v1959 int32
	_ = v1959
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1978 int32
	_ = v1978
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2012 int32
	_ = v2012
	var v2016 int32
	_ = v2016
	var v2029 int32
	_ = v2029
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2065 int32
	_ = v2065
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2083 int32
	_ = v2083
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2094 int32
	_ = v2094
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2133 int32
	_ = v2133
	var v2144 int32
	_ = v2144
	var v2154 int32
	_ = v2154
	var v2161 int32
	_ = v2161
	var v2162 int32
	_ = v2162
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int64
	_ = v2209
	var v2211 int64
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2238 int64
	_ = v2238
	var v2240 int64
	_ = v2240
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2260 int32
	_ = v2260
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2269 int32
	_ = v2269
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2301 int32
	_ = v2301
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2312 int32
	_ = v2312
	var v2324 int32
	_ = v2324
	var v2328 int32
	_ = v2328
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2371 int32
	_ = v2371
	var v2379 int32
	_ = v2379
	var v2381 int32
	_ = v2381
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2414 int32
	_ = v2414
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2425 int32
	_ = v2425
	var v2432 int32
	_ = v2432
	var v2437 int32
	_ = v2437
	var v2445 int32
	_ = v2445
	var v2449 int32
	_ = v2449
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2464 int32
	_ = v2464
	var v2475 int32
	_ = v2475
	var v2485 int32
	_ = v2485
	var v2488 int32
	_ = v2488
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2519 int32
	_ = v2519
	var v2523 int32
	_ = v2523
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2545 int32
	_ = v2545
	var v2547 int32
	_ = v2547
	var v2555 int32
	_ = v2555
	var v2557 int32
	_ = v2557
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2572 int32
	_ = v2572
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2590 int32
	_ = v2590
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2608 int32
	_ = v2608
	var v2613 int32
	_ = v2613
	var v2621 int32
	_ = v2621
	var v2625 int32
	_ = v2625
	var v2630 int32
	_ = v2630
	var v2634 int32
	_ = v2634
	var v2640 int32
	_ = v2640
	var v2651 int32
	_ = v2651
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2671 int32
	_ = v2671
	var v2674 int32
	_ = v2674
	var v2683 int32
	_ = v2683
	var v2692 int32
	_ = v2692
	var v2709 int64
	_ = v2709
	var v2715 int64
	_ = v2715
	var v2717 int64
	_ = v2717
	var v2718 int64
	_ = v2718
	var v2720 int64
	_ = v2720
	var v2725 int64
	_ = v2725
	var v2748 int64
	_ = v2748
	var v2754 int64
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2760 int64
	_ = v2760
	var v2766 int32
	_ = v2766
	var v2781 int64
	_ = v2781
	var v2788 int64
	_ = v2788
	var v2792 int64
	_ = v2792
	var v2797 int64
	_ = v2797
	var v2799 int64
	_ = v2799
	var v2801 int64
	_ = v2801
	var v2803 int32
	_ = v2803
	var v2808 int64
	_ = v2808
	var v2810 int64
	_ = v2810
	var v2812 int32
	_ = v2812
	var v2814 int64
	_ = v2814
	var v2824 int32
	_ = v2824
	var v2839 int64
	_ = v2839
	var v2850 int64
	_ = v2850
	var v2855 int64
	_ = v2855
	var v2876 int64
	_ = v2876
	var v2893 int32
	_ = v2893
	var v2906 int64
	_ = v2906
	var v2908 int64
	_ = v2908
	var v2914 int64
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2919 int64
	_ = v2919
	var v2924 int64
	_ = v2924
	var v2925 int64
	_ = v2925
	var v2927 int64
	_ = v2927
	var v2930 int64
	_ = v2930
	var v2931 int64
	_ = v2931
	var v2933 int64
	_ = v2933
	var v2934 int64
	_ = v2934
	var v2938 int64
	_ = v2938
	var v2945 int64
	_ = v2945
	var v2957 int32
	_ = v2957
	var v2958 int64
	_ = v2958
	var v2959 int64
	_ = v2959
	var v2962 int64
	_ = v2962
	var v2963 int64
	_ = v2963
	var v2966 int64
	_ = v2966
	var v2967 int64
	_ = v2967
	var v2968 int64
	_ = v2968
	var v2973 int64
	_ = v2973
	var v2977 int32
	_ = v2977
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v2982 int64
	_ = v2982
	var v2983 int64
	_ = v2983
	var v2986 int64
	_ = v2986
	var v2992 int64
	_ = v2992
	var v2993 int64
	_ = v2993
	var v2996 int64
	_ = v2996
	var v3002 int64
	_ = v3002
	var v3003 int64
	_ = v3003
	var v3004 int64
	_ = v3004
	var v3005 int64
	_ = v3005
	var v3018 int32
	_ = v3018
	var v3019 int64
	_ = v3019
	var v3020 int64
	_ = v3020
	var v3025 int64
	_ = v3025
	var v3026 int64
	_ = v3026
	var v3028 int64
	_ = v3028
	var v3031 int64
	_ = v3031
	var v3032 int64
	_ = v3032
	var v3034 int64
	_ = v3034
	var v3035 int64
	_ = v3035
	var v3039 int64
	_ = v3039
	var v3046 int64
	_ = v3046
	var v3058 int32
	_ = v3058
	var v3063 int64
	_ = v3063
	var v3064 int64
	_ = v3064
	var v3066 int64
	_ = v3066
	var v3069 int64
	_ = v3069
	var v3070 int64
	_ = v3070
	var v3072 int64
	_ = v3072
	var v3073 int64
	_ = v3073
	var v3077 int64
	_ = v3077
	var v3084 int64
	_ = v3084
	var v3096 int32
	_ = v3096
	var v3097 int64
	_ = v3097
	var v3098 int64
	_ = v3098
	var v3099 int64
	_ = v3099
	var v3108 int64
	_ = v3108
	var v3109 int64
	_ = v3109
	var v3111 int64
	_ = v3111
	var v3114 int64
	_ = v3114
	var v3115 int64
	_ = v3115
	var v3117 int64
	_ = v3117
	var v3118 int64
	_ = v3118
	var v3122 int64
	_ = v3122
	var v3129 int64
	_ = v3129
	var v3141 int32
	_ = v3141
	var v3143 int64
	_ = v3143
	var v3146 int64
	_ = v3146
	var v3147 int64
	_ = v3147
	var v3152 int64
	_ = v3152
	var v3153 int64
	_ = v3153
	var v3156 int64
	_ = v3156
	var v3159 int64
	_ = v3159
	var v3160 int64
	_ = v3160
	var v3167 int64
	_ = v3167
	var v3178 int64
	_ = v3178
	var v3179 int64
	_ = v3179
	var v3181 int64
	_ = v3181
	var v3183 int64
	_ = v3183
	var v3188 int64
	_ = v3188
	var v3189 int64
	_ = v3189
	var v3190 int64
	_ = v3190
	var v3193 int64
	_ = v3193
	var v3195 int64
	_ = v3195
	var v3196 int64
	_ = v3196
	var v3197 int64
	_ = v3197
	var v3200 int64
	_ = v3200
	var v3202 int64
	_ = v3202
	var v3203 int64
	_ = v3203
	var v3206 int64
	_ = v3206
	var v3209 int64
	_ = v3209
	var v3210 int64
	_ = v3210
	var v3217 int64
	_ = v3217
	var v3222 int32
	_ = v3222
	var v3223 int64
	_ = v3223
	var v3224 int64
	_ = v3224
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3241 int64
	_ = v3241
	var v3247 int32
	_ = v3247
	var v3248 int64
	_ = v3248
	var v3254 int32
	_ = v3254
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3279 int64
	_ = v3279
	var v3281 int64
	_ = v3281
	var v3290 int32
	_ = v3290
	var v3291 int32
	_ = v3291
	var v3293 int64
	_ = v3293
	var v3297 int32
	_ = v3297
	var v3299 int32
	_ = v3299
	var v3301 int64
	_ = v3301
	var v3302 int64
	_ = v3302
	var v3303 int64
	_ = v3303
	var v3306 int64
	_ = v3306
	var v3322 int64
	_ = v3322
	var v3323 int64
	_ = v3323
	var v3324 int64
	_ = v3324
	var v3325 int64
	_ = v3325
	var v3337 int64
	_ = v3337
	var v3338 int64
	_ = v3338
	var v3339 int64
	_ = v3339
	var v3340 int64
	_ = v3340
	var v3345 int64
	_ = v3345
	var v3348 int64
	_ = v3348
	var v3351 int64
	_ = v3351
	var v3354 int64
	_ = v3354
	var v3355 int64
	_ = v3355
	var v3359 int64
	_ = v3359
	var v3366 int64
	_ = v3366
	var v3378 int32
	_ = v3378
	var v3379 int64
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3388 int64
	_ = v3388
	var v3391 int64
	_ = v3391
	var v3393 int32
	_ = v3393
	var v3396 int64
	_ = v3396
	var v3400 int32
	_ = v3400
	var v3407 int32
	_ = v3407
	var v3409 int32
	_ = v3409
	var v3440 int32
	_ = v3440
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3468 int32
	_ = v3468
	var v3473 int32
	_ = v3473
	var v3482 int32
	_ = v3482
	var v3486 int32
	_ = v3486
	var v3504 int32
	_ = v3504
	var v3512 int32
	_ = v3512
	var v3518 int32
	_ = v3518
	var v3522 int32
	_ = v3522
	var v3526 int32
	_ = v3526
	var v3544 int32
	_ = v3544
	var v3549 int32
	_ = v3549
	var v3551 int32
	_ = v3551
	var v3553 int32
	_ = v3553
	var v3555 int32
	_ = v3555
	var v3560 int32
	_ = v3560
	var v3562 int32
	_ = v3562
	var v3563 int32
	_ = v3563
	var v3564 int32
	_ = v3564
	var v3566 int32
	_ = v3566
	var v3574 int32
	_ = v3574
	var v3578 int32
	_ = v3578
	var v3600 int32
	_ = v3600
	var v3605 int32
	_ = v3605
	var v3615 int32
	_ = v3615
	var v3646 int32
	_ = v3646
	var v3659 int64
	_ = v3659
	var v3667 int64
	_ = v3667
	var v3669 int64
	_ = v3669
	var v3671 int64
	_ = v3671
	var v3672 int64
	_ = v3672
	var v3673 int64
	_ = v3673
	var v3674 int64
	_ = v3674
	var v3680 int64
	_ = v3680
	var v3682 int64
	_ = v3682
	var v3684 int64
	_ = v3684
	var v3692 int64
	_ = v3692
	var v3715 int64
	_ = v3715
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3730 int32
	_ = v3730
	var v3742 int64
	_ = v3742
	var v3748 int64
	_ = v3748
	var v3755 int32
	_ = v3755
	var v3756 int32
	_ = v3756
	var v3769 int64
	_ = v3769
	var v3781 int32
	_ = v3781
	var v3783 int64
	_ = v3783
	var v3786 int64
	_ = v3786
	var v3789 int32
	_ = v3789
	var v3798 int32
	_ = v3798
	var v3801 int32
	_ = v3801
	var v3824 int32
	_ = v3824
	var v3831 int32
	_ = v3831
	var v3835 int32
	_ = v3835
	var v3837 int32
	_ = v3837
	var v3856 int32
	_ = v3856
	var v3859 int32
	_ = v3859
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3872 int32
	_ = v3872
	var v3874 int32
	_ = v3874
	var v3876 int32
	_ = v3876
	var v3877 int64
	_ = v3877
	var v3879 int64
	_ = v3879
	var v3882 int32
	_ = v3882
	var v3885 int32
	_ = v3885
	var v3897 int32
	_ = v3897
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3912 int32
	_ = v3912
	var v3913 int32
	_ = v3913
	var v3924 int32
	_ = v3924
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3931 int32
	_ = v3931
	var v3932 int32
	_ = v3932
	var v3933 int32
	_ = v3933
	var v3935 int32
	_ = v3935
	var v3936 int32
	_ = v3936
	var v3937 int32
	_ = v3937
	var v3940 int32
	_ = v3940
	var v3943 int32
	_ = v3943
	var v3948 int32
	_ = v3948
	var v3952 int32
	_ = v3952
	var v3958 int32
	_ = v3958
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3968 int32
	_ = v3968
	var v3971 int32
	_ = v3971
	var v3973 int32
	_ = v3973
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3978 int32
	_ = v3978
	var v3986 int32
	_ = v3986
	var v3990 int32
	_ = v3990
	var v3991 int32
	_ = v3991
	var v3994 int32
	_ = v3994
	var v4012 int32
	_ = v4012
	var v4013 int32
	_ = v4013
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4048 int32
	_ = v4048
	var v4054 int32
	_ = v4054
	var v4083 int32
	_ = v4083
	var v4084 int32
	_ = v4084
	var v4088 int32
	_ = v4088
	var v4089 int32
	_ = v4089
	var v4102 int32
	_ = v4102
	var v4134 int32
	_ = v4134
	var v4135 int32
	_ = v4135
	var v4161 int32
	_ = v4161
	var v4163 int32
	_ = v4163
	var v4164 int32
	_ = v4164
	var v4166 int32
	_ = v4166
	var v4168 int32
	_ = v4168
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4173 int32
	_ = v4173
	var v4174 int32
	_ = v4174
	var v4178 int32
	_ = v4178
	v4 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(512)
	m.G0 = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v34 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v32 + int32(512)
	return
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v47 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_pfree(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v43
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v43
	goto L1
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	if int32(0) <= v187 {
		goto L59
	} else {
		goto L60
	}
L11:
	;
	v50 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+328)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+312)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+288)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+296)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+304)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+264)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+280)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+240)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+256)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+232)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v50
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v50
	v86 = int32(-1)
	v87 = int32(1)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v90 = v88 >> (uint(v87) % 32)
	v94 = int32(4)
	v97 = base.I32_div_s(l2+v94, v94)
	v101 = base.I32_div_s(l2^v86, int32(-4))
	if int32(0) <= l2+v87 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L54
	}
L14:
	;
	v106 = v97
	goto L16
L15:
	;
	v106 = v101
	goto L16
L16:
	;
	v108 = int32(1)
	v109 = v106 + v90 + v108
	if v109 <= v108 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v112 = v87
	goto L19
L18:
	;
	v112 = v109
	goto L19
L19:
	;
	v114 = int32(1)
	v118 = (v90^v86+v112)<<(uint(v114)%32) + v88 + v114
	if v118 <= v114 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v121 = v87
	goto L22
L21:
	;
	v121 = v118
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+336)) = v121
	if int32(5) <= v118 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v131 = v121
	v132 = v4
	goto L26
L24:
	;
	v183 = v121
	v187 = v86
	goto L25
L25:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v207 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206))))
	if v183 < int32(2) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	v154 = int32(2)
	v155 = int32(base.Ui32(v131) >> (uint(v154) % 32))
	v159 = v132 + int32(1)
	if v131&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v183 = v173
	v187 = v132
	goto L25
L28:
	;
	v170 = v155
	goto L30
L29:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165))))
	v170 = v155 - base.B2i32(v166 < int32(2500))
	goto L30
L30:
	;
	v173 = v131 - v170<<(uint(int32(1))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v159<<(uint(v154)%32)))) = v173
	if int32(4) < v173 {
		v131 = v173
		v132 = v159
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v346 = base.I64_trunc_sat_f64_s(base.F64_sqrt(base.F64_convert_i64_s(v334)))
	v348 = v334 - v346*v346
	if base.B2i32(int64(0) <= v348)&base.B2i32(v348 <= v346<<(uint(int64(1))%64)) != 0 {
		v430 = v346
		v434 = v348
		goto L10
	} else {
		goto L50
	}
L33:
	;
	v321 = int32(1)
	v334 = v207
	goto L32
L34:
	;
	goto L35
L35:
	;
	if v183 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v214 = int32(1)
	v215 = v183 - v214
	v226 = v214
	v229 = int32(0)
	v241 = v207
	goto L39
L37:
	;
	v282 = int32(1)
	v297 = v207
	goto L38
L38:
	;
	v308 = v297 * int64(10000)
	if v34 <= v282 {
		v321 = v183
		v334 = v308
		goto L32
	} else {
		goto L49
	}
L39:
	;
	v252 = v241 * int64(10000)
	if v226 < v34 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v215&v214 == int32(0) {
		v321 = v183
		v334 = v270
		goto L32
	} else {
		goto L48
	}
L41:
	;
	v257 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v226<<(uint(int32(1))%32)))))
	v259 = v252 + v257
	goto L43
L42:
	;
	v259 = v252
	goto L43
L43:
	;
	v261 = v259 * int64(10000)
	v263 = v226 + int32(1)
	if v263 < v34 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v268 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v263<<(uint(int32(1))%32)))))
	v270 = v261 + v268
	goto L46
L45:
	;
	v270 = v261
	goto L46
L46:
	;
	v271 = int32(2)
	v272 = v226 + v271
	v274 = v229 + v271
	if v274 != v215&int32(-2) {
		v226 = v272
		v229 = v274
		v241 = v270
		goto L39
	} else {
		goto L47
	}
L47:
	;
	goto L40
L48:
	;
	v282 = v272
	v297 = v270
	goto L38
L49:
	;
	v313 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v282<<(uint(int32(1))%32)))))
	v321 = v183
	v334 = v308 + v313
	goto L32
L50:
	;
	v373 = v346
	goto L51
L51:
	;
	v384 = base.I64_div_s(v334, v373)
	v387 = base.I64_div_s(v384+v373, int64(2))
	v391 = v334 - v387*v387
	if base.B2i32(v391 < int64(0))|base.B2i32(v387<<(uint(int64(1))%64) < v391) != 0 {
		v373 = v387
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v430 = v387
	v434 = v391
	goto L10
L53:
	;
	goto L52
L54:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	F_errmsg(m, int32(_a_F_sqrt_var_0), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_sqrt_var_1), int32(_a_F_sqrt_var_2), int32(_a_F_sqrt_var_3))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	v3856 = v3831 << (uint(int32(1)) % 32)
	v3859 = F_palloc(m, v3856+int32(2))
	mBase = m.M
	v3860 = m.ExcPending
	if v3860 != 0 {
		goto L8
	} else {
		goto L562
	}
L59:
	;
	v449 = v321
	v453 = v187
	v461 = v430
	v465 = v434
	goto L62
L60:
	;
	v3715 = v430
	goto L61
L61:
	;
	v3727 = F_palloc(m, int32(12))
	mBase = m.M
	v3728 = m.ExcPending
	if v3728 != 0 {
		goto L8
	} else {
		goto L552
	}
L62:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v453<<(uint(int32(2))%32))))
	if v477 <= int32(8) {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v3715 = v3692
	goto L61
L64:
	;
	v3671 = v3659*v465 + v3669
	v3672 = int64(1)
	v3673 = v461 << (uint(v3672) % 64)
	v3674 = base.I64_div_s(v3671, v3673)
	v3680 = v3667 - v3674*v3674 + (v3671-v3673*v3674)*v3659
	v3682 = v3674 + v461*v3659
	v3684 = v3682 - v3672
	if v3680 < int64(0) {
		goto L548
	} else {
		goto L549
	}
L65:
	;
	v3504 = v449 + v482
	if v490 == int32(0) {
		goto L534
	} else {
		goto L535
	}
L66:
	;
	v3465 = int32(_a_F_sqrt_var_4)
	v3466 = v3443 * v3465
	v3468 = v3447 * v3465
	if v34 <= v3440 {
		v3482 = v3466
		v3486 = v3468
		goto L65
	} else {
		goto L531
	}
L67:
	;
	if v482&int32(1) == int32(0) {
		v3482 = v548
		v3486 = v550
		goto L65
	} else {
		goto L530
	}
L68:
	;
	v480 = v477 - v449
	v481 = int32(2)
	v482 = base.I32_div_s(v480, v481)
	if v480 < v481 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v556 = int64(63)
	v566 = v449
	v570 = v453
	v578 = v461
	v580 = v461 >> (uint(v556) % 64)
	v582 = v465
	v585 = v465 >> (uint(v556) % 64)
	goto L86
L71:
	;
	v486 = int64(0)
	v3646 = v449
	v3659 = int64(1)
	v3667 = v486
	v3669 = v486
	goto L64
L72:
	;
	goto L73
L73:
	;
	v488 = int32(1)
	v490 = v482 - v488
	if v490 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v3440 = v449
	v3443 = int32(0)
	v3447 = v488
	goto L66
L75:
	;
	goto L76
L76:
	;
	v498 = int32(0)
	v504 = v449
	v507 = v498
	v508 = v498
	v511 = v488
	goto L77
L77:
	;
	v530 = v507 * int32(_a_F_sqrt_var_4)
	if v504 < v34 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L67
L79:
	;
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v504<<(uint(int32(1))%32)))))
	v537 = v530 + v535
	goto L81
L80:
	;
	v537 = v530
	goto L81
L81:
	;
	v539 = v537 * int32(_a_F_sqrt_var_4)
	v541 = v504 + int32(1)
	if v541 < v34 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v541<<(uint(int32(1))%32)))))
	v548 = v539 + v546
	goto L84
L83:
	;
	v548 = v539
	goto L84
L84:
	;
	v550 = v511 * int32(100000000)
	v551 = int32(2)
	v552 = v504 + v551
	v554 = v508 + v551
	if v482&int32(1073741822) != v554 {
		v504 = v552
		v507 = v548
		v508 = v554
		v511 = v550
		goto L77
	} else {
		goto L85
	}
L85:
	;
	goto L78
L86:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v570<<(uint(int32(2))%32))))
	if v594 <= int32(16) {
		goto L92
	} else {
		goto L93
	}
L87:
	;
	v3230 = F_palloc(m, int32(22))
	mBase = m.M
	v3231 = m.ExcPending
	if v3231 != 0 {
		goto L8
	} else {
		goto L515
	}
L88:
	;
	v2917 = v32 + int32(112)
	v2919 = v2914 >> (uint(int64(63)) % 64)
	v2924 = int64(32)
	v2925 = int64(base.Ui64(v2914) >> (uint(v2924) % 64))
	v2927 = int64(base.Ui64(v582) >> (uint(v2924) % 64))
	v2930 = int64(4294967295)
	v2931 = v2914 & v2930
	v2933 = v582 & v2930
	v2934 = v2931 * v2933
	v2938 = int64(base.Ui64(v2934)>>(uint(v2924)%64)) + v2931*v2927
	v2945 = v2933*v2925 + v2938&v2930
	*(*int64)(unsafe.Add(mBase, uint32(v2917)+8)) = v582*v2919 + v585*v2914 + v2925*v2927 + int64(base.Ui64(v2938)>>(uint(v2924)%64)) + int64(base.Ui64(v2945)>>(uint(v2924)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v2917))) = v2934&v2930 | v2945<<(uint(v2924)%64)
	goto L502
L89:
	;
	v2756 = v566 + v599
	if v605 == int64(1) {
		goto L488
	} else {
		goto L489
	}
L90:
	;
	v2717 = int64(10000)
	v2718 = v2709 * v2717
	v2720 = v2715 * v2717
	if v34 <= v2692 {
		v2748 = v2718
		v2754 = v2720
		goto L89
	} else {
		goto L485
	}
L91:
	;
	if v599&int32(1) == int32(0) {
		v2748 = v663
		v2754 = v665
		goto L89
	} else {
		goto L484
	}
L92:
	;
	v597 = v594 - v566
	v598 = int32(2)
	v599 = base.I32_div_s(v597, v598)
	if v597 < v598 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v672 = F_palloc(m, int32(22))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L8
	} else {
		goto L110
	}
L95:
	;
	v603 = int64(0)
	v2893 = v566
	v2906 = v603
	v2908 = v603
	v2914 = int64(1)
	goto L88
L96:
	;
	goto L97
L97:
	;
	v605 = base.I64_extend_i32_s(v599)
	v607 = base.B2i32(v605 == int64(1))
	if v605 == int64(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v2692 = v566
	v2709 = int64(0)
	v2715 = int64(1)
	goto L90
L99:
	;
	goto L100
L100:
	;
	v613 = int64(0)
	v619 = v566
	v636 = v613
	v641 = v613
	v642 = int64(1)
	goto L101
L101:
	;
	v645 = v636 * int64(10000)
	if v619 < v34 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L91
L103:
	;
	v650 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v619<<(uint(int32(1))%32)))))
	v652 = v645 + v650
	goto L105
L104:
	;
	v652 = v645
	goto L105
L105:
	;
	v654 = v652 * int64(10000)
	v656 = v619 + int32(1)
	if v656 < v34 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v661 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v656<<(uint(int32(1))%32)))))
	v663 = v654 + v661
	goto L108
L107:
	;
	v663 = v654
	goto L108
L108:
	;
	v665 = v642 * int64(100000000)
	v667 = v619 + int32(2)
	v669 = v641 + int64(2)
	if v605&int64(1073741822) != v669 {
		v619 = v667
		v636 = v663
		v641 = v669
		v642 = v665
		goto L101
	} else {
		goto L109
	}
L109:
	;
	goto L102
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v672
	v675 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v672))) = uint16(v675)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+324)) = v675
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v672 + int32(2)
	v683 = int64(0)
	if v580 == v683 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v689 = base.B2i32(v578 != v683)
	goto L113
L112:
	;
	v689 = base.B2i32(v683 < v580)
	goto L113
L113:
	;
	v690 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+320)) = (v689 - base.B2i32(v580 < v690)) & int32(_a_F_sqrt_var_5)
	v696 = int32(0)
	if v578|v580 != v690 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v707 = v696
	v710 = v672 + int32(22)
	v721 = v578
	v723 = v580
	goto L117
L115:
	;
	v850 = v696
	v854 = v696
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v32)+312)) = v850
	v878 = F_palloc(m, int32(22))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L8
	} else {
		goto L125
	}
L117:
	;
	v733 = v32 + int32(176)
	v735 = int64(0)
	v739 = m.G0
	v740 = int32(16)
	v741 = v739 - v740
	m.G0 = v741
	v743 = int64(63)
	v744 = v723 >> (uint(v743) % 64)
	v745 = v721 ^ v744
	v748 = v745 + int64(base.Ui64(v723)>>(uint(v743)%64))
	F___udivmodti4(m, v741, v748, base.I64_extend_i32_u(base.B2i32(base.Ui64(v748) < base.Ui64(v745)))+(v744^v723), int64(10000), base.I64_extend_i32_u(int32(0))+v735)
	mBase = m.M
	v764 = *(*int64)(unsafe.Add(mBase, uint32(v741)+8))
	v765 = v744 ^ v735
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v741)))
	v767 = v765 ^ v766
	*(*int64)(unsafe.Add(mBase, uint32(v733))) = v767 - v765
	*(*int64)(unsafe.Add(mBase, uint32(v733)+8)) = v765 ^ v764 - v765 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v767) < base.Ui64(v765)))
	m.G0 = v741 + v740
	goto L119
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v822
	v850 = v837
	v854 = v707
	goto L116
L119:
	;
	v780 = v32 + int32(160)
	v781 = *(*int64)(unsafe.Add(mBase, uint32(v32)+176))
	v782 = *(*int64)(unsafe.Add(mBase, uint32(v32)+184))
	v783 = int64(4294957296)
	v784 = int64(0)
	v789 = int64(32)
	v792 = int64(base.Ui64(v781) >> (uint(v789) % 64))
	v795 = int64(4294967295)
	v798 = v781 & v795
	v799 = v783 * v798
	v803 = int64(base.Ui64(v799)>>(uint(v789)%64)) + v783*v792
	v810 = v798*v784 + v803&v795
	*(*int64)(unsafe.Add(mBase, uint32(v780)+8)) = v781*v784 + v782*v783 + v784*v792 + int64(base.Ui64(v803)>>(uint(v789)%64)) + int64(base.Ui64(v810)>>(uint(v789)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v780))) = v799&v795 | v810<<(uint(v789)%64)
	goto L120
L120:
	;
	v822 = v710 - int32(2)
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v32)+160))
	v825 = base.I32_wrap_i64(v823 + v721)
	v827 = v825 >> (uint(int32(31)) % 32)
	v829 = v825 ^ v827 - v827
	*(*uint16)(unsafe.Add(mBase, uint32(v822))) = uint16(v829)
	v832 = v721 + int64(9999)
	v835 = v723 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v832) < base.Ui64(v721)))
	v837 = v707 + int32(1)
	v840 = int64(0)
	if v835 == v840 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v844 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v832))
	goto L123
L122:
	;
	v844 = base.B2i32(v835 != v840)
	goto L123
L123:
	;
	if v844 != 0 {
		v707 = v837
		v710 = v822
		v721 = v781
		v723 = v782
		goto L117
	} else {
		goto L124
	}
L124:
	;
	goto L118
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+304)) = v878
	v881 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v878))) = uint16(v881)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+300)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(v32)+308)) = v878 + int32(2)
	v888 = int64(0)
	if v585 == v888 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v894 = base.B2i32(v582 != v888)
	goto L128
L127:
	;
	v894 = base.B2i32(v888 < v585)
	goto L128
L128:
	;
	v895 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+296)) = (v894 - base.B2i32(v585 < v895)) & int32(_a_F_sqrt_var_5)
	if v582|v585 != v895 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v911 = v878 + int32(22)
	v912 = v675
	v929 = v582
	v932 = v585
	goto L132
L130:
	;
	v1055 = v675
	v1057 = int32(0)
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+292)) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(v32)+288)) = v1055
	v1081 = int32(0)
	v1089 = v566
	v1092 = v1081
	v1093 = v570
	v1094 = v1081
	goto L141
L132:
	;
	v937 = v32 + int32(144)
	v939 = int64(0)
	v943 = m.G0
	v944 = int32(16)
	v945 = v943 - v944
	m.G0 = v945
	v947 = int64(63)
	v948 = v932 >> (uint(v947) % 64)
	v949 = v929 ^ v948
	v952 = v949 + int64(base.Ui64(v932)>>(uint(v947)%64))
	F___udivmodti4(m, v945, v952, base.I64_extend_i32_u(base.B2i32(base.Ui64(v952) < base.Ui64(v949)))+(v948^v932), int64(10000), base.I64_extend_i32_u(int32(0))+v939)
	mBase = m.M
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v945)+8))
	v969 = v948 ^ v939
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v945)))
	v971 = v969 ^ v970
	*(*int64)(unsafe.Add(mBase, uint32(v937))) = v971 - v969
	*(*int64)(unsafe.Add(mBase, uint32(v937)+8)) = v969 ^ v968 - v969 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v971) < base.Ui64(v969)))
	m.G0 = v945 + v944
	goto L134
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+308)) = v1026
	v1055 = v1041
	v1057 = v912
	goto L131
L134:
	;
	v984 = v32 + int32(128)
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v32)+144))
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v32)+152))
	v987 = int64(4294957296)
	v988 = int64(0)
	v993 = int64(32)
	v996 = int64(base.Ui64(v985) >> (uint(v993) % 64))
	v999 = int64(4294967295)
	v1002 = v985 & v999
	v1003 = v987 * v1002
	v1007 = int64(base.Ui64(v1003)>>(uint(v993)%64)) + v987*v996
	v1014 = v1002*v988 + v1007&v999
	*(*int64)(unsafe.Add(mBase, uint32(v984)+8)) = v985*v988 + v986*v987 + v988*v996 + int64(base.Ui64(v1007)>>(uint(v993)%64)) + int64(base.Ui64(v1014)>>(uint(v993)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v984))) = v1003&v999 | v1014<<(uint(v993)%64)
	goto L135
L135:
	;
	v1026 = v911 - int32(2)
	v1027 = *(*int64)(unsafe.Add(mBase, uint32(v32)+128))
	v1029 = base.I32_wrap_i64(v1027 + v929)
	v1031 = v1029 >> (uint(int32(31)) % 32)
	v1033 = v1029 ^ v1031 - v1031
	*(*uint16)(unsafe.Add(mBase, uint32(v1026))) = uint16(v1033)
	v1036 = v929 + int64(9999)
	v1039 = v932 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v1036) < base.Ui64(v929)))
	v1041 = v912 + int32(1)
	v1044 = int64(0)
	if v1039 == v1044 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v1048 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v1036))
	goto L138
L137:
	;
	v1048 = base.B2i32(v1039 != v1044)
	goto L138
L138:
	;
	if v1048 != 0 {
		v911 = v1026
		v912 = v1041
		v929 = v985
		v932 = v986
		goto L132
	} else {
		goto L139
	}
L139:
	;
	goto L133
L140:
	;
	v2683 = *(*int32)(unsafe.Add(mBase, uint32(v32)+312))
	v3831 = v2683
	v3835 = v1522
	v3837 = v1313
	goto L58
L141:
	;
	v1114 = int32(2)
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(336)+v1093<<(uint(v1114)%32))))
	v1120 = base.I32_div_s(v1117-v1089, v1114)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v1089 < v1121 {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	if v2294 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L143:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1332 = v1089 + v1120
	if v1332 < v1331 {
		goto L179
	} else {
		goto L180
	}
L144:
	;
	v1123 = v1121 - v1089
	if v1120 < v1123 {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	if v1094 != 0 {
		goto L174
	} else {
		goto L175
	}
L147:
	;
	v1125 = v1120
	goto L149
L148:
	;
	v1125 = v1123
	goto L149
L149:
	;
	if v1094 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	F_pfree(m, v1094)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L8
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v1129 = v1125 << (uint(int32(1)) % 32)
	v1132 = F_palloc(m, v1129+int32(2))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L8
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+256)) = v1132
	v1135 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1132))) = uint16(v1135)
	v1138 = v1132 + int32(2)
	if v1129 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v1138, v1139+v1089<<(uint(int32(1))%32), v1129)
	goto L157
L156:
	;
	goto L157
L157:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = int64(0)
	v1147 = v1120 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+244)) = v1147
	if int32(0) < v1125 {
		goto L160
	} else {
		goto L161
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+240)) = v1267
	*(*int32)(unsafe.Add(mBase, uint32(v32)+260)) = v1266
	v1313 = v1132
	goto L143
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+244)) = int64(0)
	v1266 = v1234
	v1267 = int32(0)
	goto L158
L160:
	;
	v1156 = v1138
	v1157 = v1125
	v1159 = v1147
	goto L163
L161:
	;
	goto L162
L162:
	;
	if v1125 != 0 {
		v1266 = v1138
		v1267 = v1125
		goto L158
	} else {
		goto L173
	}
L163:
	;
	v1181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156))))
	if v1181 != 0 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v1234 = v1138 + v1129
	goto L159
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+244)) = v1159
	v1188 = v1157
	goto L168
L166:
	;
	goto L167
L167:
	;
	v1222 = int32(1)
	if v1222 < v1157 {
		v1156 = v1156 + int32(2)
		v1157 = v1157 - v1222
		v1159 = v1159 - v1222
		goto L163
	} else {
		goto L172
	}
L168:
	;
	v1217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1156+v1188<<(uint(int32(1))%32)-int32(2)))))
	if v1217 != 0 {
		v1266 = v1156
		v1267 = v1188
		goto L158
	} else {
		goto L170
	}
L169:
	;
	v1234 = v1156
	goto L159
L170:
	;
	v1218 = int32(1)
	if v1218 < v1188 {
		v1188 = v1188 - v1218
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	goto L164
L173:
	;
	v1234 = v1138
	goto L159
L174:
	;
	F_pfree(m, v1094)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L8
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v1295 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+256)) = v1295
	*(*int64)(unsafe.Add(mBase, uint32(v32)+248)) = v1295
	*(*int64)(unsafe.Add(mBase, uint32(v32)+240)) = v1295
	v1313 = int32(0)
	goto L143
L177:
	;
	goto L176
L178:
	;
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v32)+288))
	v1544 = v1542 << (uint(int32(1)) % 32)
	v1547 = F_palloc(m, v1544+int32(2))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L8
	} else {
		goto L213
	}
L179:
	;
	v1334 = v1331 - v1332
	if v1120 < v1334 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	if v1092 != 0 {
		goto L209
	} else {
		goto L210
	}
L182:
	;
	v1336 = v1120
	goto L184
L183:
	;
	v1336 = v1334
	goto L184
L184:
	;
	if v1092 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	F_pfree(m, v1092)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L8
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v1340 = v1336 << (uint(int32(1)) % 32)
	v1343 = F_palloc(m, v1340+int32(2))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L8
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+280)) = v1343
	v1346 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1343))) = uint16(v1346)
	v1349 = v1343 + int32(2)
	if v1340 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v1349, v1350+v1332<<(uint(int32(1))%32), v1340)
	goto L192
L191:
	;
	goto L192
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = int64(0)
	v1358 = v1120 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+268)) = v1358
	if int32(0) < v1336 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+264)) = v1478
	*(*int32)(unsafe.Add(mBase, uint32(v32)+284)) = v1477
	v1522 = v1343
	goto L178
L194:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+268)) = int64(0)
	v1477 = v1445
	v1478 = int32(0)
	goto L193
L195:
	;
	v1367 = v1349
	v1368 = v1336
	v1370 = v1358
	goto L198
L196:
	;
	goto L197
L197:
	;
	if v1336 != 0 {
		v1477 = v1349
		v1478 = v1336
		goto L193
	} else {
		goto L208
	}
L198:
	;
	v1392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367))))
	if v1392 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v1445 = v1349 + v1340
	goto L194
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+268)) = v1370
	v1399 = v1368
	goto L203
L201:
	;
	goto L202
L202:
	;
	v1433 = int32(1)
	if v1433 < v1368 {
		v1367 = v1367 + int32(2)
		v1368 = v1368 - v1433
		v1370 = v1370 - v1433
		goto L198
	} else {
		goto L207
	}
L203:
	;
	v1428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1367+v1399<<(uint(int32(1))%32)-int32(2)))))
	if v1428 != 0 {
		v1477 = v1367
		v1478 = v1399
		goto L193
	} else {
		goto L205
	}
L204:
	;
	v1445 = v1367
	goto L194
L205:
	;
	v1429 = int32(1)
	if v1429 < v1399 {
		v1399 = v1399 - v1429
		goto L203
	} else {
		goto L206
	}
L206:
	;
	goto L204
L207:
	;
	goto L199
L208:
	;
	v1445 = v1349
	goto L194
L209:
	;
	F_pfree(m, v1092)
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L8
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v1506 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+280)) = v1506
	*(*int64)(unsafe.Add(mBase, uint32(v32)+272)) = v1506
	*(*int64)(unsafe.Add(mBase, uint32(v32)+264)) = v1506
	v1522 = int32(0)
	goto L178
L212:
	;
	goto L211
L213:
	;
	v1549 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1547))) = uint16(v1549)
	if base.B2i32(v1544 == v1549)|base.B2i32(v1542 <= v1549) == v1549 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v32)+308))
	base.MemoryCopy(m, v1547+int32(2), v1560, v1544)
	goto L216
L215:
	;
	goto L216
L216:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v1562 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	F_pfree(m, v1562)
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L8
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1565 = *(*int64)(unsafe.Add(mBase, uint32(v32)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v1565
	v1567 = *(*int64)(unsafe.Add(mBase, uint32(v32)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v1567
	*(*int32)(unsafe.Add(mBase, uint32(v32)+232)) = v1547
	*(*int32)(unsafe.Add(mBase, uint32(v32)+236)) = v1547 + int32(2)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+220)) = v1573 + v1120
	v1577 = v32 + int32(216)
	F_add_var(m, v1577, v32+int32(240), v1577)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L8
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	v1583 = v32 + int32(312)
	v1585 = v32 + int32(192)
	F_add_var(m, v1583, v1583, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L8
	} else {
		goto L222
	}
L222:
	;
	v1588 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+488)) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v32)+496)) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v32)+504)) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v32)+464)) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v32)+472)) = v1588
	*(*int64)(unsafe.Add(mBase, uint32(v32)+480)) = v1588
	v1601 = v32 + int32(488)
	v1602 = int32(0)
	F_div_var(m, v1577, v1585, v1601, v1602, v1602, v1602)
	mBase = m.M
	v1606 = m.ExcPending
	if v1606 != 0 {
		goto L8
	} else {
		goto L223
	}
L223:
	;
	v1608 = v32 + int32(464)
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v32)+204))
	F_mul_var(m, v1585, v1601, v1608, v1609)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L8
	} else {
		goto L224
	}
L224:
	;
	F_sub_var(m, v1577, v1608, v1608)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L8
	} else {
		goto L225
	}
L225:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	if v1614 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v32)+484))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v32)+468))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v32)+212))
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	v1745 = int32(0)
	if base.B2i32(v1744 < v1741)&base.B2i32(v1745 < v1716) == v1745 {
		v1776 = v1741
		v1780 = v1745
		goto L244
	} else {
		goto L245
	}
L227:
	;
	v1716 = int32(0)
	goto L226
L228:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v32)+472))
	if v1617 == v1618 {
		v1716 = v1614
		goto L226
	} else {
		goto L229
	}
L229:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	goto L230
L230:
	;
	if base.B2i32(v1617 != v1620) == int32(0) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v1716 = v1676
	goto L226
L232:
	;
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	if v1676 == int32(0) {
		goto L227
	} else {
		goto L240
	}
L233:
	;
	v1654 = v32 + int32(488)
	F_sub_var(m, v1654, int32(_a_F_sqrt_var_6), v1654)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L8
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1665 = v32 + int32(488)
	F_add_var(m, v1665, int32(_a_F_sqrt_var_6), v1665)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L8
	} else {
		goto L238
	}
L236:
	;
	v1659 = v32 + int32(464)
	F_add_var(m, v1659, v32+int32(192), v1659)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L8
	} else {
		goto L237
	}
L237:
	;
	goto L232
L238:
	;
	v1670 = v32 + int32(464)
	F_sub_var(m, v1670, v32+int32(192), v1670)
	mBase = m.M
	v1674 = m.ExcPending
	if v1674 != 0 {
		goto L8
	} else {
		goto L239
	}
L239:
	;
	goto L232
L240:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v32)+472))
	if v1617 != v1679 {
		goto L230
	} else {
		goto L241
	}
L241:
	;
	goto L231
L242:
	;
	if int32(0) <= v1918 {
		goto L285
	} else {
		goto L286
	}
L243:
	;
	v1918 = v1908
	goto L242
L244:
	;
	if base.B2i32(v1743 <= int32(0))|base.B2i32(v1744 <= v1776) != 0 {
		v1812 = v1744
		v1814 = v1745
		goto L251
	} else {
		goto L252
	}
L245:
	;
	v1757 = v1741
	v1761 = v1745
	goto L246
L246:
	;
	v1767 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1740+v1761<<(uint(int32(1))%32)))))
	if v1767 != 0 {
		v1908 = int32(1)
		goto L243
	} else {
		goto L248
	}
L247:
	;
	v1776 = v1771
	v1780 = v1769
	goto L244
L248:
	;
	v1768 = int32(1)
	v1769 = v1761 + v1768
	v1771 = v1757 - v1768
	if v1771 <= v1744 {
		v1776 = v1771
		v1780 = v1769
		goto L244
	} else {
		goto L249
	}
L249:
	;
	if v1769 < v1716 {
		v1757 = v1771
		v1761 = v1769
		goto L246
	} else {
		goto L250
	}
L250:
	;
	goto L247
L251:
	;
	if v1776 != v1812 {
		v1854 = v1780
		v1855 = v1814
		goto L258
	} else {
		goto L259
	}
L252:
	;
	v1793 = v1744
	v1795 = v1745
	goto L253
L253:
	;
	v1800 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1742+v1795<<(uint(int32(1))%32)))))
	if v1800 != 0 {
		v1908 = int32(-1)
		goto L243
	} else {
		goto L255
	}
L254:
	;
	v1812 = v1804
	v1814 = v1802
	goto L251
L255:
	;
	v1801 = int32(1)
	v1802 = v1795 + v1801
	v1804 = v1793 - v1801
	if v1804 <= v1776 {
		v1812 = v1804
		v1814 = v1802
		goto L251
	} else {
		goto L256
	}
L256:
	;
	if v1802 < v1743 {
		v1793 = v1804
		v1795 = v1802
		goto L253
	} else {
		goto L257
	}
L257:
	;
	goto L254
L258:
	;
	if v1716 < v1854 {
		goto L267
	} else {
		goto L268
	}
L259:
	;
	v1823 = v1780
	v1824 = v1814
	goto L260
L260:
	;
	if base.B2i32(v1716 <= v1823)|base.B2i32(v1743 <= v1824) != 0 {
		v1854 = v1823
		v1855 = v1824
		goto L258
	} else {
		goto L262
	}
L261:
	;
	if base.I32_extend16_s(v1840) < base.I32_extend16_s(v1838) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1829 = int32(1)
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1740+v1823<<(uint(v1829)%32)))))
	v1840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1824<<(uint(v1829)%32)+v1742))))
	if v1838 == v1840 {
		v1823 = v1823 + v1829
		v1824 = v1824 + v1829
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1847 = int32(1)
	goto L266
L265:
	;
	v1847 = int32(-1)
	goto L266
L266:
	;
	v1918 = v1847
	goto L242
L267:
	;
	v1858 = v1854
	goto L269
L268:
	;
	v1858 = v1716
	goto L269
L269:
	;
	v1865 = v1854
	goto L270
L270:
	;
	if v1858 == v1865 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1908 = v1891
	goto L243
L272:
	;
	if v1743 < v1855 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	v1891 = int32(1)
	v1897 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1740+v1865<<(uint(v1891)%32)))))
	if v1897 == int32(0) {
		v1865 = v1865 + v1891
		goto L270
	} else {
		goto L284
	}
L275:
	;
	v1870 = v1855
	goto L277
L276:
	;
	v1870 = v1743
	goto L277
L277:
	;
	v1878 = v1855
	goto L278
L278:
	;
	if v1870 == v1878 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v1908 = int32(-1)
	goto L243
L280:
	;
	v1918 = int32(0)
	goto L242
L281:
	;
	goto L282
L282:
	;
	v1882 = int32(1)
	v1887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1878<<(uint(v1882)%32)+v1742))))
	if v1887 == int32(0) {
		v1878 = v1878 + v1882
		goto L278
	} else {
		goto L283
	}
L283:
	;
	goto L279
L284:
	;
	goto L271
L285:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v32)+224))
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	goto L288
L286:
	;
	v2161 = v1740
	v2162 = v1716
	goto L287
L287:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v32)+488))
	v2188 = v2186 << (uint(int32(1)) % 32)
	v2191 = F_palloc(m, v2188+int32(2))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L8
	} else {
		goto L342
	}
L288:
	;
	if base.B2i32(v1921 != v1922) == int32(0) {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	v2161 = v1978
	v2162 = v1979
	goto L287
L290:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v32)+484))
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v32)+464))
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v32)+468))
	v1981 = int32(0)
	if base.B2i32(v1744 < v1980)&base.B2i32(v1981 < v1979) == v1981 {
		v2012 = v1980
		v2016 = v1981
		goto L300
	} else {
		goto L301
	}
L291:
	;
	v1956 = v32 + int32(488)
	F_add_var(m, v1956, int32(_a_F_sqrt_var_6), v1956)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L8
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1967 = v32 + int32(488)
	F_sub_var(m, v1967, int32(_a_F_sqrt_var_6), v1967)
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L8
	} else {
		goto L296
	}
L294:
	;
	v1961 = v32 + int32(464)
	F_sub_var(m, v1961, v32+int32(192), v1961)
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
		goto L8
	} else {
		goto L295
	}
L295:
	;
	goto L290
L296:
	;
	v1972 = v32 + int32(464)
	F_add_var(m, v1972, v32+int32(192), v1972)
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L8
	} else {
		goto L297
	}
L297:
	;
	goto L290
L298:
	;
	if int32(0) <= v2154 {
		goto L288
	} else {
		goto L341
	}
L299:
	;
	v2154 = v2144
	goto L298
L300:
	;
	if base.B2i32(v1743 <= int32(0))|base.B2i32(v1744 <= v2012) != 0 {
		v2048 = v1744
		v2050 = v1981
		goto L307
	} else {
		goto L308
	}
L301:
	;
	v1993 = v1980
	v1997 = v1981
	goto L302
L302:
	;
	v2003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1978+v1997<<(uint(int32(1))%32)))))
	if v2003 != 0 {
		v2144 = int32(1)
		goto L299
	} else {
		goto L304
	}
L303:
	;
	v2012 = v2007
	v2016 = v2005
	goto L300
L304:
	;
	v2004 = int32(1)
	v2005 = v1997 + v2004
	v2007 = v1993 - v2004
	if v2007 <= v1744 {
		v2012 = v2007
		v2016 = v2005
		goto L300
	} else {
		goto L305
	}
L305:
	;
	if v2005 < v1979 {
		v1993 = v2007
		v1997 = v2005
		goto L302
	} else {
		goto L306
	}
L306:
	;
	goto L303
L307:
	;
	if v2012 != v2048 {
		v2090 = v2016
		v2091 = v2050
		goto L314
	} else {
		goto L315
	}
L308:
	;
	v2029 = v1744
	v2031 = v1981
	goto L309
L309:
	;
	v2036 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1742+v2031<<(uint(int32(1))%32)))))
	if v2036 != 0 {
		v2144 = int32(-1)
		goto L299
	} else {
		goto L311
	}
L310:
	;
	v2048 = v2040
	v2050 = v2038
	goto L307
L311:
	;
	v2037 = int32(1)
	v2038 = v2031 + v2037
	v2040 = v2029 - v2037
	if v2040 <= v2012 {
		v2048 = v2040
		v2050 = v2038
		goto L307
	} else {
		goto L312
	}
L312:
	;
	if v2038 < v1743 {
		v2029 = v2040
		v2031 = v2038
		goto L309
	} else {
		goto L313
	}
L313:
	;
	goto L310
L314:
	;
	if v1979 < v2090 {
		goto L323
	} else {
		goto L324
	}
L315:
	;
	v2059 = v2016
	v2060 = v2050
	goto L316
L316:
	;
	if base.B2i32(v1979 <= v2059)|base.B2i32(v1743 <= v2060) != 0 {
		v2090 = v2059
		v2091 = v2060
		goto L314
	} else {
		goto L318
	}
L317:
	;
	if base.I32_extend16_s(v2076) < base.I32_extend16_s(v2074) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	v2065 = int32(1)
	v2074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1978+v2059<<(uint(v2065)%32)))))
	v2076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2060<<(uint(v2065)%32)+v1742))))
	if v2074 == v2076 {
		v2059 = v2059 + v2065
		v2060 = v2060 + v2065
		goto L316
	} else {
		goto L319
	}
L319:
	;
	goto L317
L320:
	;
	v2083 = int32(1)
	goto L322
L321:
	;
	v2083 = int32(-1)
	goto L322
L322:
	;
	v2154 = v2083
	goto L298
L323:
	;
	v2094 = v2090
	goto L325
L324:
	;
	v2094 = v1979
	goto L325
L325:
	;
	v2101 = v2090
	goto L326
L326:
	;
	if v2094 == v2101 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v2144 = v2127
	goto L299
L328:
	;
	if v1743 < v2091 {
		goto L331
	} else {
		goto L332
	}
L329:
	;
	goto L330
L330:
	;
	v2127 = int32(1)
	v2133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1978+v2101<<(uint(v2127)%32)))))
	if v2133 == int32(0) {
		v2101 = v2101 + v2127
		goto L326
	} else {
		goto L340
	}
L331:
	;
	v2106 = v2091
	goto L333
L332:
	;
	v2106 = v1743
	goto L333
L333:
	;
	v2114 = v2091
	goto L334
L334:
	;
	if v2106 == v2114 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v2144 = int32(-1)
	goto L299
L336:
	;
	v2154 = int32(0)
	goto L298
L337:
	;
	goto L338
L338:
	;
	v2118 = int32(1)
	v2123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2114<<(uint(v2118)%32)+v1742))))
	if v2123 == int32(0) {
		v2114 = v2114 + v2118
		goto L334
	} else {
		goto L339
	}
L339:
	;
	goto L335
L340:
	;
	goto L327
L341:
	;
	goto L289
L342:
	;
	v2193 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2191))) = uint16(v2193)
	if base.B2i32(v2188 == v2193)|base.B2i32(v2186 <= v2193) == v2193 {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v32)+508))
	base.MemoryCopy(m, v2191+int32(2), v2204, v2188)
	goto L345
L344:
	;
	goto L345
L345:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v2206 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	F_pfree(m, v2206)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L8
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	v2209 = *(*int64)(unsafe.Add(mBase, uint32(v32)+496))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v2209
	v2211 = *(*int64)(unsafe.Add(mBase, uint32(v32)+488))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v2211
	*(*int32)(unsafe.Add(mBase, uint32(v32)+232)) = v2191
	v2214 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+236)) = v2191 + v2214
	v2218 = v2162 << (uint(int32(1)) % 32)
	v2221 = F_palloc(m, v2218+v2214)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L8
	} else {
		goto L350
	}
L349:
	;
	goto L348
L350:
	;
	v2223 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2221))) = uint16(v2223)
	if base.B2i32(v2218 == v2223)|base.B2i32(v2162 <= v2223) == v2223 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	base.MemoryCopy(m, v2221+int32(2), v2161, v2218)
	goto L353
L352:
	;
	goto L353
L353:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	if v2235 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	F_pfree(m, v2235)
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L8
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v2238 = *(*int64)(unsafe.Add(mBase, uint32(v32)+472))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v2238
	v2240 = *(*int64)(unsafe.Add(mBase, uint32(v32)+464))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v2240
	*(*int32)(unsafe.Add(mBase, uint32(v32)+208)) = v2221
	*(*int32)(unsafe.Add(mBase, uint32(v32)+212)) = v2221 + int32(2)
	v2246 = *(*int32)(unsafe.Add(mBase, uint32(v32)+504))
	if v2246 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L356
L358:
	;
	F_pfree(m, v2246)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L8
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, uint32(v32)+480))
	if v2249 != 0 {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L360
L362:
	;
	F_pfree(m, v2249)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L8
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v2252 + v1120
	v2256 = v32 + int32(312)
	v2258 = v32 + int32(216)
	F_add_var(m, v2256, v2258, v2256)
	mBase = m.M
	v2260 = m.ExcPending
	if v2260 != 0 {
		goto L8
	} else {
		goto L366
	}
L365:
	;
	goto L364
L366:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+196)) = v2261 + v1120
	v2265 = v32 + int32(192)
	F_add_var(m, v2265, v32+int32(264), v2265)
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L8
	} else {
		goto L367
	}
L367:
	;
	F_mul_var(m, v2258, v2258, v2258, int32(0))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L8
	} else {
		goto L368
	}
L368:
	;
	if v1093 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v2274 = v32 + int32(288)
	F_sub_var(m, v2265, v2258, v2274)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L8
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	goto L142
L372:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
	if v2277 == int32(_a_F_sqrt_var_5) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	F_add_var(m, v2274, v2256, v2274)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L8
	} else {
		goto L376
	}
L374:
	;
	goto L375
L375:
	;
	if int32(0) < v1093 {
		v1089 = v1120 + v1332
		v1092 = v1522
		v1093 = v1093 - int32(1)
		v1094 = v1313
		goto L141
	} else {
		goto L379
	}
L376:
	;
	F_sub_var(m, v2256, int32(_a_F_sqrt_var_6), v2256)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L8
	} else {
		goto L377
	}
L377:
	;
	F_add_var(m, v2274, v2256, v2274)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L8
	} else {
		goto L378
	}
L378:
	;
	goto L375
L379:
	;
	goto L140
L380:
	;
	v2671 = v32 + int32(312)
	F_sub_var(m, v2671, int32(_a_F_sqrt_var_6), v2671)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L8
	} else {
		goto L483
	}
L381:
	;
	if v2293 == int32(0) {
		goto L140
	} else {
		goto L384
	}
L382:
	;
	goto L383
L383:
	;
	v2301 = *(*int32)(unsafe.Add(mBase, uint32(v32)+200))
	if v2293 == int32(0) {
		goto L386
	} else {
		goto L387
	}
L384:
	;
	if v2292 != int32(_a_F_sqrt_var_5) {
		goto L380
	} else {
		goto L385
	}
L385:
	;
	goto L140
L386:
	;
	if v2301 != 0 {
		goto L380
	} else {
		goto L389
	}
L387:
	;
	goto L388
L388:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v32)+236))
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	v2307 = *(*int32)(unsafe.Add(mBase, uint32(v32)+212))
	if v2301 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L389:
	;
	goto L140
L390:
	;
	if int32(0) <= v2662 {
		goto L140
	} else {
		goto L482
	}
L391:
	;
	if v2292 == int32(_a_F_sqrt_var_5) {
		goto L140
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	if v2292 == int32(0) {
		goto L380
	} else {
		goto L438
	}
L394:
	;
	v2312 = int32(0)
	if base.B2i32(v2304 < v2306)&base.B2i32(v2312 < v2294) == v2312 {
		v2343 = v2306
		v2347 = v2312
		goto L397
	} else {
		goto L398
	}
L395:
	;
	v2662 = v2485
	goto L390
L396:
	;
	v2485 = v2475
	goto L395
L397:
	;
	if base.B2i32(v2293 <= int32(0))|base.B2i32(v2304 <= v2343) != 0 {
		v2379 = v2304
		v2381 = v2312
		goto L404
	} else {
		goto L405
	}
L398:
	;
	v2324 = v2306
	v2328 = v2312
	goto L399
L399:
	;
	v2334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2307+v2328<<(uint(int32(1))%32)))))
	if v2334 != 0 {
		v2475 = int32(1)
		goto L396
	} else {
		goto L401
	}
L400:
	;
	v2343 = v2338
	v2347 = v2336
	goto L397
L401:
	;
	v2335 = int32(1)
	v2336 = v2328 + v2335
	v2338 = v2324 - v2335
	if v2338 <= v2304 {
		v2343 = v2338
		v2347 = v2336
		goto L397
	} else {
		goto L402
	}
L402:
	;
	if v2336 < v2294 {
		v2324 = v2338
		v2328 = v2336
		goto L399
	} else {
		goto L403
	}
L403:
	;
	goto L400
L404:
	;
	if v2343 != v2379 {
		v2421 = v2347
		v2422 = v2381
		goto L411
	} else {
		goto L412
	}
L405:
	;
	v2360 = v2304
	v2362 = v2312
	goto L406
L406:
	;
	v2367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2305+v2362<<(uint(int32(1))%32)))))
	if v2367 != 0 {
		v2475 = int32(-1)
		goto L396
	} else {
		goto L408
	}
L407:
	;
	v2379 = v2371
	v2381 = v2369
	goto L404
L408:
	;
	v2368 = int32(1)
	v2369 = v2362 + v2368
	v2371 = v2360 - v2368
	if v2371 <= v2343 {
		v2379 = v2371
		v2381 = v2369
		goto L404
	} else {
		goto L409
	}
L409:
	;
	if v2369 < v2293 {
		v2360 = v2371
		v2362 = v2369
		goto L406
	} else {
		goto L410
	}
L410:
	;
	goto L407
L411:
	;
	if v2294 < v2421 {
		goto L420
	} else {
		goto L421
	}
L412:
	;
	v2390 = v2347
	v2391 = v2381
	goto L413
L413:
	;
	if base.B2i32(v2294 <= v2390)|base.B2i32(v2293 <= v2391) != 0 {
		v2421 = v2390
		v2422 = v2391
		goto L411
	} else {
		goto L415
	}
L414:
	;
	if base.I32_extend16_s(v2407) < base.I32_extend16_s(v2405) {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	v2396 = int32(1)
	v2405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2307+v2390<<(uint(v2396)%32)))))
	v2407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2391<<(uint(v2396)%32)+v2305))))
	if v2405 == v2407 {
		v2390 = v2390 + v2396
		v2391 = v2391 + v2396
		goto L413
	} else {
		goto L416
	}
L416:
	;
	goto L414
L417:
	;
	v2414 = int32(1)
	goto L419
L418:
	;
	v2414 = int32(-1)
	goto L419
L419:
	;
	v2485 = v2414
	goto L395
L420:
	;
	v2425 = v2421
	goto L422
L421:
	;
	v2425 = v2294
	goto L422
L422:
	;
	v2432 = v2421
	goto L423
L423:
	;
	if v2425 == v2432 {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	v2475 = v2458
	goto L396
L425:
	;
	if v2293 < v2422 {
		goto L428
	} else {
		goto L429
	}
L426:
	;
	goto L427
L427:
	;
	v2458 = int32(1)
	v2464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2307+v2432<<(uint(v2458)%32)))))
	if v2464 == int32(0) {
		v2432 = v2432 + v2458
		goto L423
	} else {
		goto L437
	}
L428:
	;
	v2437 = v2422
	goto L430
L429:
	;
	v2437 = v2293
	goto L430
L430:
	;
	v2445 = v2422
	goto L431
L431:
	;
	if v2437 == v2445 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	v2475 = int32(-1)
	goto L396
L433:
	;
	v2485 = int32(0)
	goto L395
L434:
	;
	goto L435
L435:
	;
	v2449 = int32(1)
	v2454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2445<<(uint(v2449)%32)+v2305))))
	if v2454 == int32(0) {
		v2445 = v2445 + v2449
		goto L431
	} else {
		goto L436
	}
L436:
	;
	goto L432
L437:
	;
	goto L424
L438:
	;
	v2488 = int32(0)
	if base.B2i32(v2306 < v2304)&base.B2i32(v2488 < v2293) == v2488 {
		v2519 = v2304
		v2523 = v2488
		goto L441
	} else {
		goto L442
	}
L439:
	;
	v2662 = v2661
	goto L390
L440:
	;
	v2661 = v2651
	goto L439
L441:
	;
	if base.B2i32(v2294 <= int32(0))|base.B2i32(v2306 <= v2519) != 0 {
		v2555 = v2306
		v2557 = v2488
		goto L448
	} else {
		goto L449
	}
L442:
	;
	v2500 = v2304
	v2504 = v2488
	goto L443
L443:
	;
	v2510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2305+v2504<<(uint(int32(1))%32)))))
	if v2510 != 0 {
		v2651 = int32(1)
		goto L440
	} else {
		goto L445
	}
L444:
	;
	v2519 = v2514
	v2523 = v2512
	goto L441
L445:
	;
	v2511 = int32(1)
	v2512 = v2504 + v2511
	v2514 = v2500 - v2511
	if v2514 <= v2306 {
		v2519 = v2514
		v2523 = v2512
		goto L441
	} else {
		goto L446
	}
L446:
	;
	if v2512 < v2293 {
		v2500 = v2514
		v2504 = v2512
		goto L443
	} else {
		goto L447
	}
L447:
	;
	goto L444
L448:
	;
	if v2519 != v2555 {
		v2597 = v2523
		v2598 = v2557
		goto L455
	} else {
		goto L456
	}
L449:
	;
	v2536 = v2306
	v2538 = v2488
	goto L450
L450:
	;
	v2543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2307+v2538<<(uint(int32(1))%32)))))
	if v2543 != 0 {
		v2651 = int32(-1)
		goto L440
	} else {
		goto L452
	}
L451:
	;
	v2555 = v2547
	v2557 = v2545
	goto L448
L452:
	;
	v2544 = int32(1)
	v2545 = v2538 + v2544
	v2547 = v2536 - v2544
	if v2547 <= v2519 {
		v2555 = v2547
		v2557 = v2545
		goto L448
	} else {
		goto L453
	}
L453:
	;
	if v2545 < v2294 {
		v2536 = v2547
		v2538 = v2545
		goto L450
	} else {
		goto L454
	}
L454:
	;
	goto L451
L455:
	;
	if v2293 < v2597 {
		goto L464
	} else {
		goto L465
	}
L456:
	;
	v2566 = v2523
	v2567 = v2557
	goto L457
L457:
	;
	if base.B2i32(v2293 <= v2566)|base.B2i32(v2294 <= v2567) != 0 {
		v2597 = v2566
		v2598 = v2567
		goto L455
	} else {
		goto L459
	}
L458:
	;
	if base.I32_extend16_s(v2583) < base.I32_extend16_s(v2581) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	v2572 = int32(1)
	v2581 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2305+v2566<<(uint(v2572)%32)))))
	v2583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2567<<(uint(v2572)%32)+v2307))))
	if v2581 == v2583 {
		v2566 = v2566 + v2572
		v2567 = v2567 + v2572
		goto L457
	} else {
		goto L460
	}
L460:
	;
	goto L458
L461:
	;
	v2590 = int32(1)
	goto L463
L462:
	;
	v2590 = int32(-1)
	goto L463
L463:
	;
	v2661 = v2590
	goto L439
L464:
	;
	v2601 = v2597
	goto L466
L465:
	;
	v2601 = v2293
	goto L466
L466:
	;
	v2608 = v2597
	goto L467
L467:
	;
	if v2601 == v2608 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v2651 = v2634
	goto L440
L469:
	;
	if v2294 < v2598 {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	goto L471
L471:
	;
	v2634 = int32(1)
	v2640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2305+v2608<<(uint(v2634)%32)))))
	if v2640 == int32(0) {
		v2608 = v2608 + v2634
		goto L467
	} else {
		goto L481
	}
L472:
	;
	v2613 = v2598
	goto L474
L473:
	;
	v2613 = v2294
	goto L474
L474:
	;
	v2621 = v2598
	goto L475
L475:
	;
	if v2613 == v2621 {
		goto L477
	} else {
		goto L478
	}
L476:
	;
	v2651 = int32(-1)
	goto L440
L477:
	;
	v2661 = int32(0)
	goto L439
L478:
	;
	goto L479
L479:
	;
	v2625 = int32(1)
	v2630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2621<<(uint(v2625)%32)+v2307))))
	if v2630 == int32(0) {
		v2621 = v2621 + v2625
		goto L475
	} else {
		goto L480
	}
L480:
	;
	goto L476
L481:
	;
	goto L468
L482:
	;
	goto L380
L483:
	;
	goto L140
L484:
	;
	v2692 = v667
	v2709 = v663
	v2715 = v665
	goto L90
L485:
	;
	v2725 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2692<<(uint(int32(1))%32)))))
	v2748 = v2718 + v2725
	v2754 = v2720
	goto L89
L486:
	;
	v2893 = v2756 + v599
	v2906 = v2876
	v2908 = v2748
	v2914 = v2754
	goto L88
L487:
	;
	v2850 = v2839 * int64(10000)
	if v34 <= v2824 {
		v2876 = v2850
		goto L486
	} else {
		goto L501
	}
L488:
	;
	v2824 = v2756
	v2839 = int64(0)
	goto L487
L489:
	;
	goto L490
L490:
	;
	v2760 = int64(0)
	v2766 = v2756
	v2781 = v2760
	v2788 = v2760
	goto L491
L491:
	;
	v2792 = v2781 * int64(10000)
	if v2766 < v34 {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	if v599&int32(1) == int32(0) {
		v2876 = v2810
		goto L486
	} else {
		goto L500
	}
L493:
	;
	v2797 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2766<<(uint(int32(1))%32)))))
	v2799 = v2792 + v2797
	goto L495
L494:
	;
	v2799 = v2792
	goto L495
L495:
	;
	v2801 = v2799 * int64(10000)
	v2803 = v2766 + int32(1)
	if v2803 < v34 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2808 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2803<<(uint(int32(1))%32)))))
	v2810 = v2801 + v2808
	goto L498
L497:
	;
	v2810 = v2801
	goto L498
L498:
	;
	v2812 = v2766 + int32(2)
	v2814 = v2788 + int64(2)
	if v2814 != v605&int64(1073741822) {
		v2766 = v2812
		v2781 = v2810
		v2788 = v2814
		goto L491
	} else {
		goto L499
	}
L499:
	;
	goto L492
L500:
	;
	v2824 = v2812
	v2839 = v2810
	goto L487
L501:
	;
	v2855 = int64(*(*int16)(unsafe.Add(mBase, uint32(v206+v2824<<(uint(int32(1))%32)))))
	v2876 = v2850 + v2855
	goto L486
L502:
	;
	v2957 = v32 + int32(96)
	v2958 = *(*int64)(unsafe.Add(mBase, uint32(v32)+112))
	v2959 = v2958 + v2908
	v2962 = *(*int64)(unsafe.Add(mBase, uint32(v32)+120))
	v2963 = int64(63)
	v2966 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v2959) < base.Ui64(v2958))) + (v2962 + v2908>>(uint(v2963)%64))
	v2967 = int64(1)
	v2968 = v578 << (uint(v2967) % 64)
	v2973 = v580<<(uint(v2967)%64) | int64(base.Ui64(v578)>>(uint(v2963)%64))
	v2977 = m.G0
	v2978 = int32(16)
	v2979 = v2977 - v2978
	m.G0 = v2979
	v2982 = v2966 >> (uint(v2963) % 64)
	v2983 = v2959 ^ v2982
	v2986 = v2983 + int64(base.Ui64(v2966)>>(uint(v2963)%64))
	v2992 = v2973 >> (uint(v2963) % 64)
	v2993 = v2992 ^ v2968
	v2996 = v2993 + int64(base.Ui64(v2973)>>(uint(v2963)%64))
	F___udivmodti4(m, v2979, v2986, base.I64_extend_i32_u(base.B2i32(base.Ui64(v2986) < base.Ui64(v2983)))+(v2982^v2966), v2996, base.I64_extend_i32_u(base.B2i32(base.Ui64(v2996) < base.Ui64(v2993)))+(v2992^v2973))
	mBase = m.M
	v3002 = *(*int64)(unsafe.Add(mBase, uint32(v2979)+8))
	v3003 = v2982 ^ v2992
	v3004 = *(*int64)(unsafe.Add(mBase, uint32(v2979)))
	v3005 = v3003 ^ v3004
	*(*int64)(unsafe.Add(mBase, uint32(v2957))) = v3005 - v3003
	*(*int64)(unsafe.Add(mBase, uint32(v2957)+8)) = v3003 ^ v3002 - v3003 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v3005) < base.Ui64(v3003)))
	m.G0 = v2979 + v2978
	goto L503
L503:
	;
	v3018 = v32 + int32(80)
	v3019 = *(*int64)(unsafe.Add(mBase, uint32(v32)+96))
	v3020 = *(*int64)(unsafe.Add(mBase, uint32(v32)+104))
	v3025 = int64(32)
	v3026 = int64(base.Ui64(v2968) >> (uint(v3025) % 64))
	v3028 = int64(base.Ui64(v3019) >> (uint(v3025) % 64))
	v3031 = int64(4294967295)
	v3032 = v2968 & v3031
	v3034 = v3019 & v3031
	v3035 = v3032 * v3034
	v3039 = int64(base.Ui64(v3035)>>(uint(v3025)%64)) + v3032*v3028
	v3046 = v3034*v3026 + v3039&v3031
	*(*int64)(unsafe.Add(mBase, uint32(v3018)+8)) = v3019*v2973 + v3020*v2968 + v3026*v3028 + int64(base.Ui64(v3039)>>(uint(v3025)%64)) + int64(base.Ui64(v3046)>>(uint(v3025)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3018))) = v3035&v3031 | v3046<<(uint(v3025)%64)
	goto L504
L504:
	;
	v3058 = v32 + int32(32)
	v3063 = int64(32)
	v3064 = int64(base.Ui64(v2914) >> (uint(v3063) % 64))
	v3066 = int64(base.Ui64(v578) >> (uint(v3063) % 64))
	v3069 = int64(4294967295)
	v3070 = v2914 & v3069
	v3072 = v578 & v3069
	v3073 = v3070 * v3072
	v3077 = int64(base.Ui64(v3073)>>(uint(v3063)%64)) + v3070*v3066
	v3084 = v3072*v3064 + v3077&v3069
	*(*int64)(unsafe.Add(mBase, uint32(v3058)+8)) = v578*v2919 + v580*v2914 + v3064*v3066 + int64(base.Ui64(v3077)>>(uint(v3063)%64)) + int64(base.Ui64(v3084)>>(uint(v3063)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3058))) = v3073&v3069 | v3084<<(uint(v3063)%64)
	goto L505
L505:
	;
	v3096 = v32 - int32(-64)
	v3097 = *(*int64)(unsafe.Add(mBase, uint32(v32)+80))
	v3098 = v2959 - v3097
	v3099 = *(*int64)(unsafe.Add(mBase, uint32(v32)+88))
	v3108 = int64(32)
	v3109 = int64(base.Ui64(v2914) >> (uint(v3108) % 64))
	v3111 = int64(base.Ui64(v3098) >> (uint(v3108) % 64))
	v3114 = int64(4294967295)
	v3115 = v2914 & v3114
	v3117 = v3098 & v3114
	v3118 = v3115 * v3117
	v3122 = int64(base.Ui64(v3118)>>(uint(v3108)%64)) + v3115*v3111
	v3129 = v3117*v3109 + v3122&v3114
	*(*int64)(unsafe.Add(mBase, uint32(v3096)+8)) = v3098*v2919 + (v2966-v3099-base.I64_extend_i32_u(base.B2i32(base.Ui64(v2959) < base.Ui64(v3097))))*v2914 + v3109*v3111 + int64(base.Ui64(v3122)>>(uint(v3108)%64)) + int64(base.Ui64(v3129)>>(uint(v3108)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3096))) = v3118&v3114 | v3129<<(uint(v3108)%64)
	goto L506
L506:
	;
	v3141 = v32 + int32(48)
	v3143 = v3019 * v3020
	v3146 = int64(32)
	v3147 = int64(base.Ui64(v3019) >> (uint(v3146) % 64))
	v3152 = int64(4294967295)
	v3153 = v3019 & v3152
	v3156 = v3153 * v3153
	v3159 = v3153 * v3147
	v3160 = int64(base.Ui64(v3156)>>(uint(v3146)%64)) + v3159
	v3167 = v3159 + v3160&v3152
	*(*int64)(unsafe.Add(mBase, uint32(v3141)+8)) = v3143 + v3143 + v3147*v3147 + int64(base.Ui64(v3160)>>(uint(v3146)%64)) + int64(base.Ui64(v3167)>>(uint(v3146)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v3141))) = v3156&v3152 | v3167<<(uint(v3146)%64)
	goto L507
L507:
	;
	v3178 = *(*int64)(unsafe.Add(mBase, uint32(v32)+72))
	v3179 = int64(63)
	v3181 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
	v3183 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
	v3188 = v2906 - v3183
	v3189 = *(*int64)(unsafe.Add(mBase, uint32(v32)+64))
	v3190 = v3188 + v3189
	v3193 = v3178 + (v2906>>(uint(v3179)%64) - v3181 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v2906) < base.Ui64(v3183)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v3190) < base.Ui64(v3188)))
	v3195 = v3193 >> (uint(v3179) % 64)
	v3196 = *(*int64)(unsafe.Add(mBase, uint32(v32)+32))
	v3197 = v3019 + v3196
	v3200 = *(*int64)(unsafe.Add(mBase, uint32(v32)+40))
	v3202 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v3197) < base.Ui64(v3019))) + (v3020 + v3200)
	v3203 = int64(0)
	v3206 = v3202 - base.I64_extend_i32_u(base.B2i32(v3197 == v3203))
	v3209 = v3197 - int64(1)
	v3210 = v3209 + v3197
	v3217 = v3190 + v3210&v3195
	v3222 = base.B2i32(v3193 < v3203)
	if v3193 < v3203 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3223 = v3206
	goto L510
L509:
	;
	v3223 = v3202
	goto L510
L510:
	;
	if v3193 < v3203 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v3224 = v3209
	goto L513
L512:
	;
	v3224 = v3197
	goto L513
L513:
	;
	if int32(0) < v570 {
		v566 = v2893
		v570 = v570 - int32(1)
		v578 = v3224
		v580 = v3223
		v582 = v3217
		v585 = v3193 + v3195&(v3206+v3202+base.I64_extend_i32_u(base.B2i32(base.Ui64(v3210) < base.Ui64(v3209)))) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v3217) < base.Ui64(v3190)))
		goto L86
	} else {
		goto L514
	}
L514:
	;
	goto L87
L515:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v3230
	v3233 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3230))) = uint16(v3233)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+324)) = v3233
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3230 + int32(2)
	v3241 = int64(0)
	if v3223 == v3241 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v3247 = base.B2i32(v3224 != v3241)
	goto L518
L517:
	;
	v3247 = base.B2i32(v3241 < v3223)
	goto L518
L518:
	;
	v3248 = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+320)) = (v3247 - base.B2i32(v3223 < v3248)) & int32(_a_F_sqrt_var_5)
	v3254 = int32(0)
	if v3224|v3223 != v3248 {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v3265 = v3230 + int32(22)
	v3266 = v3254
	v3279 = v3224
	v3281 = v3223
	goto L522
L520:
	;
	v3407 = v3254
	v3409 = v3254
	goto L521
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v3409
	*(*int32)(unsafe.Add(mBase, uint32(v32)+312)) = v3407
	v3831 = v3407
	v3835 = int32(0)
	v3837 = v3233
	goto L58
L522:
	;
	v3290 = int32(16)
	v3291 = v32 + v3290
	v3293 = int64(0)
	v3297 = m.G0
	v3299 = v3297 - v3290
	m.G0 = v3299
	v3301 = int64(63)
	v3302 = v3281 >> (uint(v3301) % 64)
	v3303 = v3279 ^ v3302
	v3306 = v3303 + int64(base.Ui64(v3281)>>(uint(v3301)%64))
	F___udivmodti4(m, v3299, v3306, base.I64_extend_i32_u(base.B2i32(base.Ui64(v3306) < base.Ui64(v3303)))+(v3302^v3281), int64(10000), base.I64_extend_i32_u(int32(0))+v3293)
	mBase = m.M
	v3322 = *(*int64)(unsafe.Add(mBase, uint32(v3299)+8))
	v3323 = v3302 ^ v3293
	v3324 = *(*int64)(unsafe.Add(mBase, uint32(v3299)))
	v3325 = v3323 ^ v3324
	*(*int64)(unsafe.Add(mBase, uint32(v3291))) = v3325 - v3323
	*(*int64)(unsafe.Add(mBase, uint32(v3291)+8)) = v3323 ^ v3322 - v3323 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v3325) < base.Ui64(v3323)))
	m.G0 = v3299 + v3290
	goto L524
L523:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3378
	v3407 = v3393
	v3409 = v3266
	goto L521
L524:
	;
	v3337 = *(*int64)(unsafe.Add(mBase, uint32(v32)+16))
	v3338 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
	v3339 = int64(4294957296)
	v3340 = int64(0)
	v3345 = int64(32)
	v3348 = int64(base.Ui64(v3337) >> (uint(v3345) % 64))
	v3351 = int64(4294967295)
	v3354 = v3337 & v3351
	v3355 = v3339 * v3354
	v3359 = int64(base.Ui64(v3355)>>(uint(v3345)%64)) + v3339*v3348
	v3366 = v3354*v3340 + v3359&v3351
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v3337*v3340 + v3338*v3339 + v3340*v3348 + int64(base.Ui64(v3359)>>(uint(v3345)%64)) + int64(base.Ui64(v3366)>>(uint(v3345)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v32))) = v3355&v3351 | v3366<<(uint(v3345)%64)
	goto L525
L525:
	;
	v3378 = v3265 - int32(2)
	v3379 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
	v3381 = base.I32_wrap_i64(v3379 + v3279)
	v3383 = v3381 >> (uint(int32(31)) % 32)
	v3385 = v3381 ^ v3383 - v3383
	*(*uint16)(unsafe.Add(mBase, uint32(v3378))) = uint16(v3385)
	v3388 = v3279 + int64(9999)
	v3391 = v3281 + base.I64_extend_i32_u(base.B2i32(base.Ui64(v3388) < base.Ui64(v3279)))
	v3393 = v3266 + int32(1)
	v3396 = int64(0)
	if v3391 == v3396 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	v3400 = base.B2i32(base.Ui64(int64(19998)) < base.Ui64(v3388))
	goto L528
L527:
	;
	v3400 = base.B2i32(v3391 != v3396)
	goto L528
L528:
	;
	if v3400 != 0 {
		v3265 = v3378
		v3266 = v3393
		v3279 = v3337
		v3281 = v3338
		goto L522
	} else {
		goto L529
	}
L529:
	;
	goto L523
L530:
	;
	v3440 = v552
	v3443 = v548
	v3447 = v550
	goto L66
L531:
	;
	v3473 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3440<<(uint(int32(1))%32)))))
	v3482 = v3466 + v3473
	v3486 = v3468
	goto L65
L532:
	;
	v3646 = v3504 + v482
	v3659 = base.I64_extend_i32_s(v3486)
	v3667 = base.I64_extend_i32_s(v3615)
	v3669 = base.I64_extend_i32_s(v3482)
	goto L64
L533:
	;
	v3600 = v3578 * int32(_a_F_sqrt_var_4)
	if v34 <= v3574 {
		v3615 = v3600
		goto L532
	} else {
		goto L547
	}
L534:
	;
	v3574 = v3504
	v3578 = int32(0)
	goto L533
L535:
	;
	goto L536
L536:
	;
	v3512 = int32(0)
	v3518 = v3504
	v3522 = v3512
	v3526 = v3512
	goto L537
L537:
	;
	v3544 = v3522 * int32(_a_F_sqrt_var_4)
	if v3518 < v34 {
		goto L539
	} else {
		goto L540
	}
L538:
	;
	if v482&int32(1) == int32(0) {
		v3615 = v3562
		goto L532
	} else {
		goto L546
	}
L539:
	;
	v3549 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3518<<(uint(int32(1))%32)))))
	v3551 = v3544 + v3549
	goto L541
L540:
	;
	v3551 = v3544
	goto L541
L541:
	;
	v3553 = v3551 * int32(_a_F_sqrt_var_4)
	v3555 = v3518 + int32(1)
	if v3555 < v34 {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v3560 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3555<<(uint(int32(1))%32)))))
	v3562 = v3553 + v3560
	goto L544
L543:
	;
	v3562 = v3553
	goto L544
L544:
	;
	v3563 = int32(2)
	v3564 = v3518 + v3563
	v3566 = v3526 + v3563
	if v3566 != v482&int32(1073741822) {
		v3518 = v3564
		v3522 = v3562
		v3526 = v3566
		goto L537
	} else {
		goto L545
	}
L545:
	;
	goto L538
L546:
	;
	v3574 = v3564
	v3578 = v3562
	goto L533
L547:
	;
	v3605 = int32(*(*int16)(unsafe.Add(mBase, uint32(v206+v3574<<(uint(int32(1))%32)))))
	v3615 = v3600 + v3605
	goto L532
L548:
	;
	v3692 = v3684
	goto L550
L549:
	;
	v3692 = v3682
	goto L550
L550:
	;
	if int32(0) < v453 {
		v449 = v3646
		v453 = v453 - int32(1)
		v461 = v3692
		v465 = v3680 + (v3684+v3682)&(v3680>>(uint(int64(63))%64))
		goto L62
	} else {
		goto L551
	}
L551:
	;
	goto L63
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+328)) = v3727
	v3730 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3727))) = uint16(v3730)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3727 + int32(2)
	if v3715 < int64(0) {
		goto L555
	} else {
		goto L556
	}
L553:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+316)) = v3801
	*(*int32)(unsafe.Add(mBase, uint32(v32)+312)) = v3798
	v3824 = int32(0)
	v3831 = v3798
	v3835 = v3824
	v3837 = v3824
	goto L58
L554:
	;
	v3755 = v3727 + int32(12)
	v3756 = v3730
	v3769 = v3748
	goto L559
L555:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = int64(16384)
	v3748 = int64(0) - v3715
	goto L554
L556:
	;
	goto L557
L557:
	;
	v3742 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+320)) = v3742
	if v3715 == v3742 {
		v3798 = v3730
		v3801 = int32(0)
		goto L553
	} else {
		goto L558
	}
L558:
	;
	v3748 = v3715
	goto L554
L559:
	;
	v3781 = v3755 - int32(2)
	v3783 = base.I64_div_u_s(v3769, int64(10000))
	v3786 = v3783*int64(55536) + v3769
	*(*uint16)(unsafe.Add(mBase, uint32(v3781))) = uint16(v3786)
	v3789 = v3756 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v3769) {
		v3755 = v3781
		v3756 = v3789
		v3769 = v3783
		goto L559
	} else {
		goto L561
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+332)) = v3781
	v3798 = v3789
	v3801 = v3756
	goto L553
L561:
	;
	goto L560
L562:
	;
	v3861 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v3859))) = uint16(v3861)
	if base.B2i32(v3856 == v3861)|base.B2i32(v3831 <= v3861) == v3861 {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v3872 = *(*int32)(unsafe.Add(mBase, uint32(v32)+332))
	base.MemoryCopy(m, v3859+int32(2), v3872, v3856)
	goto L565
L564:
	;
	goto L565
L565:
	;
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v3874 != 0 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	F_pfree(m, v3874)
	mBase = m.M
	v3876 = m.ExcPending
	if v3876 != 0 {
		goto L8
	} else {
		goto L569
	}
L567:
	;
	goto L568
L568:
	;
	v3877 = *(*int64)(unsafe.Add(mBase, uint32(v32)+320))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v3877
	v3879 = *(*int64)(unsafe.Add(mBase, uint32(v32)+312))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v3879
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v3859
	v3882 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v3859 + v3882
	v3885 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v3885
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	v3897 = l2 + v90<<(uint(v3882)%32)
	if v3897+int32(4) < v3885 {
		goto L571
	} else {
		goto L572
	}
L569:
	;
	goto L568
L570:
	;
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v4013 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v4013 {
		goto L598
	} else {
		goto L599
	}
L571:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L570
L572:
	;
	goto L573
L573:
	;
	v3906 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v3908 = l2 & int32(3)
	v3912 = base.I32_div_s(v3897+int32(7), int32(4))
	v3913 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v3913 <= v3912 {
		goto L578
	} else {
		goto L579
	}
L574:
	;
	goto L570
L575:
	;
	if int32(0) <= v3978 {
		goto L574
	} else {
		goto L595
	}
L576:
	;
	v3958 = v3952
	goto L589
L577:
	;
	v3927 = int32(1)
	v3928 = v3912 - v3927
	v3931 = v3906 + v3928<<(uint(v3927)%32)
	v3932 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3931))))
	v3933 = int32(2)
	v3935 = *(*int32)(unsafe.Add(mBase, uint32(v3908<<(uint(v3933)%32))+uint32(_c_F_sqrt_var[0])))
	v3936 = base.I32_rem_s(v3932, v3935)
	v3937 = v3932 - v3936
	*(*uint16)(unsafe.Add(mBase, uint32(v3931))) = uint16(v3937)
	v3940 = base.I32_div_s(v3935, v3933)
	if v3936 < v3940 {
		v3978 = v3928
		goto L575
	} else {
		goto L584
	}
L578:
	;
	if base.B2i32(v3908 == int32(0))|base.B2i32(v3912 != v3913) != 0 {
		goto L574
	} else {
		goto L581
	}
L579:
	;
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3912
	if v3908 != 0 {
		goto L577
	} else {
		goto L582
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3912
	goto L577
L582:
	;
	v3924 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3906+v3912<<(uint(int32(1))%32)))))
	if v3924 <= int32(_a_F_sqrt_var_7) {
		v3978 = v3912
		goto L575
	} else {
		goto L583
	}
L583:
	;
	v3952 = v3912
	goto L576
L584:
	;
	v3943 = v3935 + base.I32_extend16_s(v3937)
	if int32(_a_F_sqrt_var_8) < v3943 {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v3948 = v3943 + int32(_a_F_sqrt_var_9)
	goto L587
L586:
	;
	v3948 = v3943
	goto L587
L587:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v3931))) = uint16(v3948)
	if v3943 < int32(_a_F_sqrt_var_4) {
		v3978 = v3928
		goto L575
	} else {
		goto L588
	}
L588:
	;
	v3952 = v3928
	goto L576
L589:
	;
	v3964 = int32(1)
	v3965 = v3958 - v3964
	v3968 = v3906 + v3965<<(uint(v3964)%32)
	v3971 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3968))))
	v3973 = base.B2i32(int32(_a_F_sqrt_var_10) < v3971)
	if int32(_a_F_sqrt_var_10) < v3971 {
		goto L591
	} else {
		goto L592
	}
L590:
	;
	v3978 = v3965
	goto L575
L591:
	;
	v3974 = int32(-9999)
	goto L593
L592:
	;
	v3974 = v3964
	goto L593
L593:
	;
	v3975 = v3974 + v3971
	*(*uint16)(unsafe.Add(mBase, uint32(v3968))) = uint16(v3975)
	if int32(_a_F_sqrt_var_10) < v3971 {
		v3958 = v3965
		goto L589
	} else {
		goto L594
	}
L594:
	;
	goto L590
L595:
	;
	v3986 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v3986 - int32(2)
	v3990 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v3991 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v3990 + v3991
	v3994 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3994 + v3991
	goto L574
L596:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4135
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v4134
	v4161 = *(*int32)(unsafe.Add(mBase, uint32(v32)+328))
	if v4161 != 0 {
		goto L612
	} else {
		goto L613
	}
L597:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(0)
	v4134 = v4102
	v4135 = int32(0)
	goto L596
L598:
	;
	v4023 = v4012
	v4024 = v4013
	goto L601
L599:
	;
	goto L600
L600:
	;
	if v4013 != 0 {
		v4134 = v4012
		v4135 = v4013
		goto L596
	} else {
		goto L611
	}
L601:
	;
	v4048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4023))))
	if v4048 != 0 {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	v4102 = v4012 + v4013<<(uint(int32(1))%32)
	goto L597
L603:
	;
	v4054 = v4024
	goto L606
L604:
	;
	goto L605
L605:
	;
	v4088 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v4089 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v4088 - v4089
	if v4089 < v4024 {
		v4023 = v4023 + int32(2)
		v4024 = v4024 - v4089
		goto L601
	} else {
		goto L610
	}
L606:
	;
	v4083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4023+v4054<<(uint(int32(1))%32)-int32(2)))))
	if v4083 != 0 {
		v4134 = v4023
		v4135 = v4054
		goto L596
	} else {
		goto L608
	}
L608:
	;
	v4084 = int32(1)
	if v4084 < v4054 {
		v4054 = v4054 - v4084
		goto L606
	} else {
		goto L609
	}
L609:
	;
	v4102 = v4023
	goto L597
L610:
	;
	goto L602
L611:
	;
	v4102 = v4012
	goto L597
L612:
	;
	F_pfree(m, v4161)
	mBase = m.M
	v4163 = m.ExcPending
	if v4163 != 0 {
		goto L8
	} else {
		goto L615
	}
L613:
	;
	goto L614
L614:
	;
	v4164 = *(*int32)(unsafe.Add(mBase, uint32(v32)+304))
	if v4164 != 0 {
		goto L616
	} else {
		goto L617
	}
L615:
	;
	goto L614
L616:
	;
	F_pfree(m, v4164)
	mBase = m.M
	v4166 = m.ExcPending
	if v4166 != 0 {
		goto L8
	} else {
		goto L619
	}
L617:
	;
	goto L618
L618:
	;
	if v3835 != 0 {
		goto L620
	} else {
		goto L621
	}
L619:
	;
	goto L618
L620:
	;
	F_pfree(m, v3835)
	mBase = m.M
	v4168 = m.ExcPending
	if v4168 != 0 {
		goto L8
	} else {
		goto L623
	}
L621:
	;
	goto L622
L622:
	;
	if v3837 != 0 {
		goto L624
	} else {
		goto L625
	}
L623:
	;
	goto L622
L624:
	;
	F_pfree(m, v3837)
	mBase = m.M
	v4170 = m.ExcPending
	if v4170 != 0 {
		goto L8
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	v4171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	if v4171 != 0 {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L626
L628:
	;
	F_pfree(m, v4171)
	mBase = m.M
	v4173 = m.ExcPending
	if v4173 != 0 {
		goto L8
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	v4174 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	if v4174 == int32(0) {
		goto L1
	} else {
		goto L632
	}
L631:
	;
	goto L630
L632:
	;
	F_pfree(m, v4174)
	mBase = m.M
	v4178 = m.ExcPending
	if v4178 != 0 {
		goto L8
	} else {
		goto L633
	}
L633:
	;
	goto L1
}
func F_stat(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	v5 = F___fstatat(m, int32(-100), l0, l1, int32(0))
	return v5
}
func F_statement_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_statement_timestamp[0]))
	return v3
}
func F_str_initcap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l0 == v4 {
		v92 = v4
		m.G0 = v10 + int32(16)
		return v92
	} else {
		if l2 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_str_initcap_0)
					F_errmsg(m, int32(_a_F_str_initcap_1), v10)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_str_initcap_2), int32(0))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_str_initcap_3), int32(1767), int32(_a_F_str_initcap_4))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
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
		} else {
			v17 = F_pg_newlocale_from_collation(m, l2)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+2)))
				if v21 == int32(1) {
					v24 = F_pnstrdup(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
						if v26 == int32(0) {
							v92 = v24
						} else {
							v29 = v26
							v31 = v24
							v32 = int32(1)
							for {
								if base.Ui32((v29-int32(97))&int32(255)) < base.Ui32(int32(26)) {
									v44 = v29 - int32(32)
								} else {
									v44 = v29
								}
								if base.Ui32((v29-int32(65))&int32(255)) < base.Ui32(int32(26)) {
									v53 = v29 | int32(32)
								} else {
									v53 = v29
								}
								if v32&int32(1) != 0 {
									v56 = v44
								} else {
									v56 = v53
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v56)
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
								if v72 != 0 {
									v29 = v72
									v31 = v31 + int32(1)
									v32 = base.B2i32(base.Ui32((v56&int32(223)-int32(91))&int32(255)) < base.Ui32(int32(230))) & base.B2i32(base.Ui32(base.I32_extend8_s(v56)-int32(58)) < base.Ui32(int32(-10)))
									continue
								} else {
									break
								}
								break
							}
							v92 = v24
						}
						m.G0 = v10 + int32(16)
						return v92
					}
				} else {
					v76 = l1 + int32(1)
					v77 = F_palloc(m, v76)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = F_pg_strtitle(m, v77, v76, l0, l1, v17)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							v82 = v79 + int32(1)
							if base.Ui32(v82) <= base.Ui32(v76) {
								v92 = v77
								m.G0 = v10 + int32(16)
								return v92
							} else {
								v84 = F_repalloc(m, v77, v82)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v86 = F_pg_strtitle(m, v84, v82, l0, l1, v17)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										v92 = v84
										m.G0 = v10 + int32(16)
										return v92
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
func F_strict_word_similarity_commutator_op(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14407(m, l0, int32(_a_F_strict_word_similarity_commutator_op_0), int32(3))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_strlist_to_textarray(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v2
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_strlist_to_textarray_0), v2, int32(_a_F_strlist_to_textarray_1), int32(_a_F_strlist_to_textarray_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = int32(_a_F_strlist_to_textarray_3)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0])) = v23
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_strlist_to_textarray[0])) = v28
	v101 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v101
	v112 = F_construct_md_array(m, v92, v93, v101, v13+int32(12), v13+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L20
	}
L4:
	;
	v35 = F_palloc_mul(m, int32(8), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = F_palloc_mul(m, int32(8), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v39 = F_palloc_mul(m, int32(1), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v92 = v35
	v93 = v39
	goto L3
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = F_palloc_mul(m, int32(1), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v49 <= int32(0) {
		v92 = v43
		v93 = v47
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v57 = v2
	v60 = v2
	goto L12
L12:
	;
	v62 = v47 + v57
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v60<<(uint(int32(2))%32))))
	if v67 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v92 = v43
	v93 = v47
	goto L3
L14:
	;
	v86 = v60 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v86 < v87 {
		v57 = v83
		v60 = v86
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v68)
	v70 = F_cstring_to_text(m, v67)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v80)
	v83 = v57
	goto L14
L18:
	;
	v73 = v57 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v43+v57<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v70)
	v83 = v73
	goto L14
L19:
	;
	goto L13
L20:
	;
	F_MemoryContextDelete(m, v23)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	m.G0 = v13 + int32(16)
	return v112
}
func F_strncoll_libc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	v9 = m.G0
	v11 = v9 - int32(1024)
	m.G0 = v11
	v15 = l1 + l3 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = F_palloc(m, v15)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v22 = v11
	goto L3
L3:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = v18
	goto L3
L6:
	;
	base.MemoryCopy(m, v22, l0, l1)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v24 = l1 + v22
	v25 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v25)
	v28 = v24 + int32(1)
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	base.MemoryCopy(m, v28, l2, l3)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28+l3))) = uint8(v31)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(v36 == v31)|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v22 != v11 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v42 = v22
	v43 = v28
	goto L15
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v57 = v47
	v58 = v46
	goto L13
L17:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_pfree(m, v22)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	m.G0 = v11 + int32(1024)
	return v57 - v58
L22:
	;
	goto L21
}
func F_strrchr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = F_strlen(m, l0)
	mBase = m.M
	v12 = v5 + int32(1)
	goto L2
L1:
	;
	return v24
L2:
	;
	v14 = int32(0)
	if v12 == v14 {
		v24 = v14
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v24 = v19
	goto L1
L4:
	;
	v18 = v12 - int32(1)
	v19 = l0 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v20 != l1&int32(255) {
		v12 = v18
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
}
func F_strspn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	v6 = m.G0
	v8 = v6 - int32(32)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v9
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v17 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = l0
	goto L7
L5:
	;
	goto L6
L6:
	;
	v37 = l1
	v38 = v17
	goto L10
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v32 == v17 {
		v26 = v26 + int32(1)
		goto L7
	} else {
		goto L9
	}
L8:
	;
	return v26 - l0
L9:
	;
	goto L8
L10:
	;
	v45 = v8 + int32(base.Ui32(v38)>>(uint(int32(3))%32))&int32(28)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v46 | v47<<(uint(v38)%32)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v51 != 0 {
		v37 = v37 + v47
		v38 = v51
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v54 == int32(0) {
		v77 = l0
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return v77 - l0
L14:
	;
	v58 = l0
	v59 = v54
	goto L15
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(base.Ui32(v59)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v67)>>(uint(v59)%32))&int32(1) == int32(0) {
		v77 = v58
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v77 = v75
	goto L13
L17:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v75 = v58 + int32(1)
	if v73 != 0 {
		v58 = v75
		v59 = v73
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_strtitle_libc_sb(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	if base.Ui32(l1) < base.Ui32(l3+int32(1)) {
	} else {
		if l3 != 0 {
			base.MemoryCopy(m, l0, l2, l3)
		} else {
		}
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0+l3))) = uint8(v13)
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v15 == v13 {
		} else {
			v18 = l0
			v19 = v15
			v23 = int32(0)
			for {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+3)))
				if v24 == int32(1) {
					if v23 != 0 {
						if base.Ui32((v19-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
							v99 = v19 | int32(32)
							*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v99)
							v102 = v99
						} else {
							if int32(0) <= base.I32_extend8_s(v19) {
								v102 = v19
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								if base.B2i32(base.Ui32(v19&int32(255)-int32(65)) < base.Ui32(int32(26))) == int32(0) {
									v102 = v44
								} else {
									if base.Ui32(v44-int32(65)) < base.Ui32(int32(26)) {
										v53 = v44 | int32(32)
									} else {
										v53 = v44
									}
									v99 = v53
									*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v99)
									v102 = v99
								}
							}
						}
					} else {
						if base.Ui32((v19-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
							v99 = v19 - int32(32)
							*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v99)
							v102 = v99
						} else {
							if int32(0) <= base.I32_extend8_s(v19) {
								v102 = v19
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
								if base.B2i32(base.Ui32(v19&int32(255)-int32(97)) < base.Ui32(int32(26))) == int32(0) {
									v102 = v71
								} else {
									if base.Ui32(v71-int32(97)) < base.Ui32(int32(26)) {
										v80 = v71 & int32(95)
									} else {
										v80 = v71
									}
									v99 = v80
									*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v99)
									v102 = v99
								}
							}
						}
					}
				} else {
					v82 = v19 & int32(255)
					if v23 != 0 {
						if base.Ui32(v82-int32(65)) < base.Ui32(int32(26)) {
							v89 = v82 | int32(32)
						} else {
							v89 = v82
						}
						v99 = v89
					} else {
						if base.Ui32(v82-int32(97)) < base.Ui32(int32(26)) {
							v96 = v82 & int32(95)
						} else {
							v96 = v82
						}
						v99 = v96
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v99)
					v102 = v99
				}
				v106 = v102 & int32(255)
				v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
				if v118 != 0 {
					v18 = v18 + int32(1)
					v19 = v118
					v23 = base.B2i32(base.Ui32(v106-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v106|int32(32)-int32(97)) < base.Ui32(int32(26)))
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return l3
}
func F_subxact_info_read(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(1040)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v13 = v8 + int32(16)
	v16 = F_pg_snprintf(m, v13, int32(1024), int32(_a_F_subxact_info_read_0), v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[0]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
		v23 = F_BufFileOpenFileSet(m, v20, v13, int32(0), int32(1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			if v23 != 0 {
				F_BufFileReadExact(m, v23, int32(_a_F_subxact_info_read_1), int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = int32(_a_F_subxact_info_read_2)
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[1]))
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[2]))
					*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[1])) = v33
					v36 = int32(1)
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[3]))
					v44 = v36 << (uint(int32(0)-base.I32_clz(v39-v36)) % 32)
					*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[4])) = v44
					v47 = F_palloc_mul(m, int32(16), v44)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[1])) = v30
						*(*int32)(unsafe.Add(mBase, _c_F_subxact_info_read[5])) = v47
						v54 = v39 << (uint(int32(4)) % 32)
						if v54 != 0 {
							F_BufFileReadExact(m, v23, v47, v54)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_BufFileClose(m, v23)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									m.G0 = v8 + int32(1040)
									return
								}
							}
						} else {
							F_BufFileClose(m, v23)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								m.G0 = v8 + int32(1040)
								return
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(1040)
				return
			}
		}
	}
}
func F_superuser(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_superuser[0]))
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_superuser[1]))
	if base.B2i32(v5 == v1)|base.B2i32(v5 != v9) == v1 {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])))
		v58 = v15
		return v58 & int32(1)
	} else {
		if v9 == int32(10) {
			v18 = int32(1)
			v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[3])))
			if v20&v18 == int32(0) {
				v58 = v18
				return v58 & int32(1)
			} else {
				v29 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(v9))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v29 != 0 {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v34)+68)))
						F_ReleaseCatCache(m, v29)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = v36
							v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
							if v41 == int32(0) {
								F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									v50 = int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v50)
									*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
									v56 = v39 & int32(1)
									*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
									v58 = v39
									return v58 & int32(1)
								}
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v56 = v39 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
								v58 = v39
								return v58 & int32(1)
							}
						}
					} else {
						v39 = int32(0)
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
						if v41 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v50 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v50)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v56 = v39 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
								v58 = v39
								return v58 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v56 = v39 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
							v58 = v39
							return v58 & int32(1)
						}
					}
				}
			}
		} else {
			v29 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(v9))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 != 0 {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v34)+68)))
					F_ReleaseCatCache(m, v29)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v36
						v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
						if v41 == int32(0) {
							F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								v50 = int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v50)
								*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
								v56 = v39 & int32(1)
								*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
								v58 = v39
								return v58 & int32(1)
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v56 = v39 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
							v58 = v39
							return v58 & int32(1)
						}
					}
				} else {
					v39 = int32(0)
					v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])))
					if v41 == int32(0) {
						F_CacheRegisterSyscacheCallback(m, int32(11), int32(1991), int64(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v50 = int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[4])) = uint8(v50)
							*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
							v56 = v39 & int32(1)
							*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
							v58 = v39
							return v58 & int32(1)
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_superuser[0])) = v9
						v56 = v39 & int32(1)
						*(*uint8)(unsafe.Add(mBase, _c_F_superuser[2])) = uint8(v56)
						v58 = v39
						return v58 & int32(1)
					}
				}
			}
		}
	}
}
func F_swedish_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L4
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v313 < v315 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	if v60 < int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	goto L5
L5:
	;
	goto L6
L6:
	;
	v15 = v8
	v17 = int32(3)
	goto L9
L8:
	;
	v60 = v45
	goto L2
L9:
	;
	if v5 <= v15 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v60 = int32(-1)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v22 = v15 + int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v15))))
	if base.Ui32(v24) < base.Ui32(int32(192)) {
		v45 = v22
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = int32(1)
	if v46 < v17 {
		v15 = v45
		v17 = v17 - v46
		goto L9
	} else {
		goto L21
	}
L15:
	;
	if v5 <= v22 {
		v45 = v22
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v31 = v22
	goto L17
L17:
	;
	v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7+v31))))
	if int32(-65) < v34 {
		v45 = v31
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v45 = v5
	goto L14
L19:
	;
	v38 = v31 + int32(1)
	if v38 != v5 {
		v31 = v38
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	goto L10
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v85 = v8
	goto L25
L23:
	;
	if v180 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L24:
	;
	v180 = v152
	goto L23
L25:
	;
	if v76 <= v85 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v180 = int32(-1)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v92 = int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v77))))
	if base.Ui32(v94) < base.Ui32(int32(192)) {
		v151 = v94
		v152 = v92
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if int32(246) < v151 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v98 = v85 + int32(1)
	if v98 == v76 {
		v151 = v94
		v152 = v92
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+v77))))
	v103 = v101 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v94) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v77))))
	v119 = v117 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v94) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v107 = v85 + int32(2)
	if v107 != v76 {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v151 = v94<<(uint(int32(6))%32)&int32(1984) | v103
	v152 = int32(2)
	goto L30
L37:
	;
	goto L36
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v123))))
	v151 = v136&int32(63) | (v94<<(uint(int32(18))%32)&int32(_a_F_swedish_UTF_8_stem_0) | v103<<(uint(int32(12))%32) | v119<<(uint(int32(6))%32))
	v152 = int32(4)
	goto L30
L39:
	;
	v123 = v85 + int32(3)
	if v123 != v76 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v151 = v94<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v103<<(uint(int32(6))%32) | v119
	v152 = int32(3)
	goto L30
L42:
	;
	goto L41
L43:
	;
	v169 = v152 + v85
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v169
	v85 = v169
	goto L25
L44:
	;
	v156 = v151 - int32(97)
	if v156 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v156)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[0]))))
	if int32(base.Ui32(v162)>>(uint(v156&int32(7))%32))&int32(1) != 0 {
		goto L24
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v184 = v183 + v180
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v207 = v184
	goto L51
L49:
	;
	if v303 < int32(0) {
		goto L1
	} else {
		goto L73
	}
L50:
	;
	v303 = v274
	goto L49
L51:
	;
	if v198 <= v207 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v303 = int32(-1)
	goto L49
L54:
	;
	goto L55
L55:
	;
	v214 = int32(1)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v199))))
	if base.Ui32(v216) < base.Ui32(int32(192)) {
		v273 = v216
		v274 = v214
		goto L56
	} else {
		goto L57
	}
L56:
	;
	if int32(246) < v273 {
		goto L50
	} else {
		goto L69
	}
L57:
	;
	v220 = v207 + int32(1)
	if v220 == v198 {
		v273 = v216
		v274 = v214
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v199))))
	v225 = v223 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v216) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v199))))
	v241 = v239 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v216) {
		goto L65
	} else {
		goto L66
	}
L60:
	;
	v229 = v207 + int32(2)
	if v229 != v198 {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v273 = v216<<(uint(int32(6))%32)&int32(1984) | v225
	v274 = int32(2)
	goto L56
L63:
	;
	goto L62
L64:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199+v245))))
	v273 = v258&int32(63) | (v216<<(uint(int32(18))%32)&int32(_a_F_swedish_UTF_8_stem_0) | v225<<(uint(int32(12))%32) | v241<<(uint(int32(6))%32))
	v274 = int32(4)
	goto L56
L65:
	;
	v245 = v207 + int32(3)
	if v245 != v198 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v273 = v216<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v225<<(uint(int32(6))%32) | v241
	v274 = int32(3)
	goto L56
L68:
	;
	goto L67
L69:
	;
	v278 = v273 - int32(97)
	if v278 < int32(0) {
		goto L50
	} else {
		goto L70
	}
L70:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v278)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[0]))))
	if int32(base.Ui32(v284)>>(uint(v278&int32(7))%32))&int32(1) == int32(0) {
		goto L50
	} else {
		goto L71
	}
L71:
	;
	v292 = v274 + v207
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v292
	v207 = v292
	goto L51
L73:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v307 = v306 + v303
	if v60 < v307 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v309 = v307
	goto L76
L75:
	;
	v309 = v60
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v309
	goto L1
L77:
	;
	return v811
L78:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v525
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v527 <= v525 {
		goto L127
	} else {
		goto L128
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v315
	if v313 <= v315 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	goto L78
L81:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = int32(1)
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v313-v322))))
	if base.B2i32(v324&int32(224) != int32(96))|base.B2i32(v322<<(uint(v324)%32)&int32(_a_F_swedish_UTF_8_stem_2) == int32(0)) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v339 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_3), int32(38), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	return int32(0)
L84:
	;
	if v339 == int32(0) {
		goto L80
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v346
	switch v339 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	case 2:
		goto L86
	default:
		goto L78
	}
L86:
	;
	v514 = F_r_et_condition_2(m, l0)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L83
	} else {
		goto L124
	}
L87:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v354 = int32(2)
	v356 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v358-v359 < v354 {
		v369 = v356
		goto L93
	} else {
		goto L94
	}
L88:
	;
	v350 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v350 {
		goto L78
	} else {
		goto L89
	}
L89:
	;
	v811 = v350
	goto L77
L90:
	;
	v511 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v511 {
		goto L78
	} else {
		goto L123
	}
L91:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v380 = v378 + (v346 - v353)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v380
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L101
L92:
	;
	if v369 == int32(0) {
		goto L91
	} else {
		goto L96
	}
L93:
	;
	goto L92
L94:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v365 = F_memcmp(m, v362+v358-v354, int32(_a_F_swedish_UTF_8_stem_4), v354)
	mBase = m.M
	if v365 != 0 {
		v369 = v356
		goto L93
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v358 - v354
	v369 = int32(1)
	goto L93
L96:
	;
	v372 = F_r_et_condition_2(m, l0)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L83
	} else {
		goto L97
	}
L97:
	;
	if v372 == int32(0) {
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v376
	goto L90
L99:
	;
	if v510 != 0 {
		goto L78
	} else {
		goto L122
	}
L100:
	;
	v510 = v503
	goto L99
L101:
	;
	if v380 <= v395 {
		v503 = int32(-1)
		goto L100
	} else {
		goto L103
	}
L102:
	;
	v503 = int32(0)
	goto L100
L103:
	;
	v412 = int32(1)
	v413 = v380 - v412
	v415 = int32(*(*int8)(unsafe.Add(mBase, uint32(v396+v413))))
	v417 = v415 & int32(255)
	if base.B2i32(v413 == v395)|base.B2i32(int32(0) <= v415) != 0 {
		v475 = v417
		v479 = v412
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if int32(121) < v475 {
		goto L112
	} else {
		goto L113
	}
L105:
	;
	v424 = v417 & int32(63)
	v426 = v380 - int32(2)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v426))))
	v430 = v428 << (uint(int32(6)) % 32)
	if base.B2i32(v426 != v395)&base.B2i32(base.Ui32(v428) < base.Ui32(int32(192))) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v475 = v430&int32(1984) | v424
	v479 = int32(2)
	goto L104
L107:
	;
	goto L108
L108:
	;
	v443 = v430&int32(4032) | v424
	v445 = v380 - int32(3)
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396+v445))))
	if base.B2i32(v445 != v395)&base.B2i32(base.Ui32(v447) < base.Ui32(int32(224))) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v475 = v447<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v443
	v479 = int32(3)
	goto L104
L110:
	;
	goto L111
L111:
	;
	v465 = int32(4)
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380+v396-v465))))
	v475 = v447<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_5) | v467&int32(7)<<(uint(int32(18))%32) | v443
	v479 = v465
	goto L104
L112:
	;
	v510 = v479
	goto L99
L113:
	;
	goto L114
L114:
	;
	v481 = v475 - int32(98)
	if v481 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v510 = v479
	goto L99
L116:
	;
	goto L117
L117:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v481)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[1]))))
	if int32(base.Ui32(v487)>>(uint(v481&int32(7))%32))&int32(1) == int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v510 = v479
	goto L99
L119:
	;
	goto L120
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v380 - v479
	goto L121
L121:
	;
	goto L102
L122:
	;
	goto L90
L123:
	;
	v811 = v511
	goto L77
L124:
	;
	if v514 == int32(0) {
		goto L78
	} else {
		goto L125
	}
L125:
	;
	v518 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v518 {
		goto L78
	} else {
		goto L126
	}
L126:
	;
	v811 = v518
	goto L77
L127:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v527
	v532 = v525 - int32(1)
	if v532 <= v527 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v623 = v525
	v624 = v527
	goto L129
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v623
	if v623 < v624 {
		goto L156
	} else {
		goto L157
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v529
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v623 = v622
	v624 = v621
	goto L129
L131:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534+v532))))
	if base.B2i32(v536&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v536)%32)&int32(_a_F_swedish_UTF_8_stem_6) == int32(0)) != 0 {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v551 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_7), int32(7), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L83
	} else {
		goto L133
	}
L133:
	;
	if v551 == int32(0) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v555
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L137
L135:
	;
	if v611 < int32(0) {
		goto L130
	} else {
		goto L154
	}
L137:
	;
	goto L138
L138:
	;
	goto L139
L139:
	;
	v566 = v555
	v568 = int32(1)
	goto L142
L141:
	;
	v611 = v593
	goto L135
L142:
	;
	if v566 <= v559 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L141
L144:
	;
	v611 = int32(-1)
	goto L135
L145:
	;
	goto L146
L146:
	;
	v573 = v566 - int32(1)
	v575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v558+v573))))
	if base.B2i32(int32(0) <= v575)|base.B2i32(v573 <= v559) != 0 {
		v593 = v573
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v597 = int32(1)
	if v597 < v568 {
		v566 = v593
		v568 = v568 - v597
		goto L142
	} else {
		goto L153
	}
L148:
	;
	v581 = v573
	goto L149
L149:
	;
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558+v581))))
	if base.Ui32(int32(191)) < base.Ui32(v586) {
		v593 = v581
		goto L147
	} else {
		goto L151
	}
L150:
	;
	v593 = v559
	goto L147
L151:
	;
	v590 = v581 - int32(1)
	if v559 < v590 {
		v581 = v590
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	goto L143
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v611
	v616 = F_slice_del(m, l0)
	mBase = m.M
	if v616 < int32(0) {
		v811 = v616
		goto L77
	} else {
		goto L155
	}
L155:
	;
	goto L130
L156:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v808
	v811 = int32(1)
	goto L77
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v623
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v624
	v632 = v623 - int32(1)
	if v632 <= v624 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v629
	goto L156
L159:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634+v632))))
	if base.B2i32(v636&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v636)%32)&int32(_a_F_swedish_UTF_8_stem_8) == int32(0)) != 0 {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v651 = F_find_among_b(m, l0, int32(_a_F_swedish_UTF_8_stem_9), int32(5), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L83
	} else {
		goto L161
	}
L161:
	;
	if v651 == int32(0) {
		goto L158
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v629
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v656
	switch v651 - int32(1) {
	case 0:
		goto L165
	case 1:
		goto L164
	case 2:
		goto L163
	default:
		goto L156
	}
L163:
	;
	v800 = F_slice_from_s(m, l0, int32(4), int32(_a_F_swedish_UTF_8_stem_10))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L83
	} else {
		goto L193
	}
L164:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L169
L165:
	;
	v660 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v660 {
		goto L156
	} else {
		goto L166
	}
L166:
	;
	v811 = v660
	goto L77
L167:
	;
	if v791 != 0 {
		goto L156
	} else {
		goto L190
	}
L168:
	;
	v791 = v784
	goto L167
L169:
	;
	if v675 <= v676 {
		v784 = int32(-1)
		goto L168
	} else {
		goto L171
	}
L170:
	;
	v784 = int32(0)
	goto L168
L171:
	;
	v693 = int32(1)
	v694 = v675 - v693
	v696 = int32(*(*int8)(unsafe.Add(mBase, uint32(v677+v694))))
	v698 = v696 & int32(255)
	if base.B2i32(v694 == v676)|base.B2i32(int32(0) <= v696) != 0 {
		v756 = v698
		v760 = v693
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if int32(118) < v756 {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v705 = v698 & int32(63)
	v707 = v675 - int32(2)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v707))))
	v711 = v709 << (uint(int32(6)) % 32)
	if base.B2i32(v707 != v676)&base.B2i32(base.Ui32(v709) < base.Ui32(int32(192))) == int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v756 = v711&int32(1984) | v705
	v760 = int32(2)
	goto L172
L175:
	;
	goto L176
L176:
	;
	v724 = v711&int32(4032) | v705
	v726 = v675 - int32(3)
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677+v726))))
	if base.B2i32(v726 != v676)&base.B2i32(base.Ui32(v728) < base.Ui32(int32(224))) == int32(0) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v756 = v728<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_1) | v724
	v760 = int32(3)
	goto L172
L178:
	;
	goto L179
L179:
	;
	v746 = int32(4)
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675+v677-v746))))
	v756 = v728<<(uint(int32(12))%32)&int32(_a_F_swedish_UTF_8_stem_5) | v748&int32(7)<<(uint(int32(18))%32) | v724
	v760 = v746
	goto L172
L180:
	;
	v791 = v760
	goto L167
L181:
	;
	goto L182
L182:
	;
	v762 = v756 - int32(105)
	if v762 < int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v791 = v760
	goto L167
L184:
	;
	goto L185
L185:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v762)>>(uint(int32(3))%32)))+uint32(_c_F_swedish_UTF_8_stem[2]))))
	if int32(base.Ui32(v768)>>(uint(v762&int32(7))%32))&int32(1) == int32(0) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v791 = v760
	goto L167
L187:
	;
	goto L188
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675 - v760
	goto L189
L189:
	;
	goto L170
L190:
	;
	v794 = F_slice_from_s(m, l0, int32(3), int32(_a_F_swedish_UTF_8_stem_11))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L83
	} else {
		goto L191
	}
L191:
	;
	if int32(0) <= v794 {
		goto L156
	} else {
		goto L192
	}
L192:
	;
	v811 = v794
	goto L77
L193:
	;
	if int32(0) <= v800 {
		goto L156
	} else {
		goto L194
	}
L194:
	;
	v811 = v800
	goto L77
}
