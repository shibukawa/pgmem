package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_C_IMatchText(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
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
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
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
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
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
		v411 = l2
		v412 = l3
		v415 = base.B2i32(int32(0) < l1)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return v450
L10:
	;
	if v415 != 0 {
		v450 = v6
		goto L9
	} else {
		goto L146
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
	v411 = v400
	v412 = v398
	v415 = v396
	goto L10
L14:
	;
	v395 = int32(1)
	v396 = base.B2i32(v395 < v35)
	v398 = v394 - v395
	v400 = v392 + v395
	if v35 < int32(2) {
		v411 = v400
		v412 = v398
		v415 = v396
		goto L10
	} else {
		goto L144
	}
L15:
	;
	v392 = v36
	v394 = v37
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
	v385 = F_pg_strncoll(m, v204, v238, v34, v35, l4)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L139
	}
L19:
	;
	if v44 == int32(92) {
		goto L115
	} else {
		goto L116
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L7
	} else {
		goto L110
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
		goto L57
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
	v450 = int32(1)
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
	if base.Ui32((v88-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v101 = v88 | int32(32)
	goto L45
L44:
	;
	v101 = v88
	goto L45
L45:
	;
	v103 = v53
	v104 = v54
	goto L46
L46:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if base.Ui32((v113-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v450 = int32(-1)
	goto L9
L48:
	;
	v122 = v113 | int32(32)
	goto L50
L49:
	;
	v122 = v113
	goto L50
L50:
	;
	if (base.B2i32(v122&int32(255) == base.I32_extend8_s(v101))|v15)&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v129 = F_C_IMatchText(m, v103, v104, v66, v64, l4)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v133 = int32(1)
	if v133 < v104 {
		v103 = v103 + v133
		v104 = v104 - v133
		goto L46
	} else {
		goto L56
	}
L54:
	;
	if v129 != 0 {
		v450 = v129
		goto L9
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L47
L57:
	;
	v147 = v37
	v149 = v36
	v150 = v44
	v153 = int32(0)
	goto L61
L58:
	;
	v289 = F_pg_strncoll(m, v36, v187-v36, v34, v35, l4)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L109
	}
L59:
	;
	v250 = v35
	v255 = v34
	goto L91
L60:
	;
	v204 = F_palloc(m, v201-v36)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L83
	}
L61:
	;
	v155 = v150 & int32(255)
	switch v155 - int32(92) {
	case 0:
		goto L68
	case 1, 2:
		goto L65
	case 3:
		goto L66
	default:
		goto L67
	}
L62:
	;
	if v153 == int32(0) {
		goto L58
	} else {
		goto L82
	}
L63:
	;
	goto L62
L64:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v147 = v192
	v149 = v193
	v150 = v195
	v153 = v194
	goto L61
L65:
	;
	v186 = int32(1)
	v187 = v149 + v186
	v189 = v147 - v186
	if v189 == int32(0) {
		goto L63
	} else {
		goto L81
	}
L66:
	;
	if v153 != 0 {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	if v155 != int32(37) {
		goto L65
	} else {
		goto L77
	}
L68:
	;
	if v147 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v176 = int32(2)
	v177 = v149 + v176
	v178 = int32(1)
	v180 = v147 - v176
	if v180 != 0 {
		v192 = v180
		v193 = v177
		v194 = v178
		goto L64
	} else {
		goto L76
	}
L72:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_C_IMatchText_0), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_C_IMatchText_1), int32(236), int32(_a_F_C_IMatchText_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
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
	v200 = int32(0)
	v201 = v177
	v202 = v178
	goto L60
L77:
	;
	goto L66
L78:
	;
	v200 = v147
	v201 = v149
	v202 = int32(0)
	goto L60
L79:
	;
	goto L80
L80:
	;
	v241 = v36
	v242 = v147
	v244 = v149
	v246 = v149 - v36
	v247 = v6
	goto L59
L81:
	;
	v192 = v189
	v193 = v187
	v194 = v153
	goto L64
L82:
	;
	v200 = int32(0)
	v201 = v187
	v202 = int32(1)
	goto L60
L83:
	;
	if base.Ui32(v36) < base.Ui32(v201) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v209 = v36
	v213 = v204
	goto L87
L85:
	;
	v234 = v204
	goto L86
L86:
	;
	v238 = v234 - v204
	if v202 != 0 {
		goto L18
	} else {
		goto L90
	}
L87:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	v220 = v209 + base.B2i32(v217 == int32(92))
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v221)
	v223 = int32(1)
	v224 = v213 + v223
	v226 = v220 + v223
	if base.Ui32(v226) < base.Ui32(v201) {
		v209 = v226
		v213 = v224
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v234 = v224
	goto L86
L89:
	;
	goto L88
L90:
	;
	v241 = v204
	v242 = v200
	v244 = v201
	v246 = v238
	v247 = v204
	goto L59
L91:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_C_IMatchText[0]))
	if v260 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v264 = F_pg_strncoll(m, v241, v246, v34, v255-v34, l4)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	if v250 != 0 {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	if v264 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v266 = F_C_IMatchText(m, v255, v250, v244, v242, l4)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	if v266 != int32(1) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	if v247 == int32(0) {
		v450 = int32(1)
		goto L9
	} else {
		goto L102
	}
L102:
	;
	F_pfree(m, v247)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	return int32(1)
L104:
	;
	v277 = int32(1)
	v250 = v250 - v277
	v255 = v255 + v277
	goto L91
L105:
	;
	v281 = int32(0)
	if v247 == v281 {
		v450 = v281
		goto L9
	} else {
		goto L107
	}
L107:
	;
	F_pfree(m, v247)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	return int32(0)
L109:
	;
	return base.B2i32(v289 == int32(0))
L110:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_C_IMatchText_0), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_C_IMatchText_1), int32(168), int32(_a_F_C_IMatchText_2))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L135
	}
L115:
	;
	if base.Ui32(v37) <= base.Ui32(int32(1)) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.Ui32((v345-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L128
	} else {
		goto L129
	}
L118:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if base.Ui32((v314-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v323 = v314 | int32(32)
	goto L121
L120:
	;
	v323 = v314
	goto L121
L121:
	;
	v324 = int32(255)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if base.Ui32((v326-int32(65))&v324) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v335 = v326 | int32(32)
	goto L124
L123:
	;
	v335 = v326
	goto L124
L124:
	;
	if v323&v324 == v335&int32(255) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v339 = int32(1)
	v392 = v36 + v339
	v394 = v37 - v339
	goto L14
L126:
	;
	goto L127
L127:
	;
	return int32(0)
L128:
	;
	v354 = v345 | int32(32)
	goto L130
L129:
	;
	v354 = v345
	goto L130
L130:
	;
	v355 = int32(255)
	if base.Ui32((v44-int32(65))&v355) < base.Ui32(int32(26)) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v365 = v44 | int32(32)
	goto L133
L132:
	;
	v365 = v44
	goto L133
L133:
	;
	if v354&v355 == v365 {
		v392 = v36
		v394 = v37
		goto L14
	} else {
		goto L134
	}
L134:
	;
	return int32(0)
L135:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(_a_F_C_IMatchText_0), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_C_IMatchText_1), int32(356), int32(_a_F_C_IMatchText_2))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	if v204 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_pfree(m, v204)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	return base.B2i32(v385 == int32(0))
L143:
	;
	goto L142
L144:
	;
	v403 = int32(1)
	if v403 < v394 {
		v34 = v34 + v403
		v35 = v35 - v403
		v36 = v400
		v37 = v398
		goto L12
	} else {
		goto L145
	}
L145:
	;
	goto L13
L146:
	;
	v419 = int32(1)
	if v412 <= int32(0) {
		v450 = v419
		goto L9
	} else {
		goto L147
	}
L147:
	;
	v424 = v411
	v425 = v412
	goto L148
L148:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424))))
	if v432 != int32(37) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v450 = v419
	goto L9
L150:
	;
	return int32(-1)
L151:
	;
	goto L152
L152:
	;
	v437 = int32(1)
	if v437 < v425 {
		v424 = v424 + v437
		v425 = v425 - v437
		goto L148
	} else {
		goto L153
	}
L153:
	;
	goto L149
}
