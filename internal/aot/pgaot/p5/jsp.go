package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_jspGetArg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_jspInitByBuffer(m, l1, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_jspGetRightArg(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_jspInitByBuffer(m, l1, v3, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_jspIsMutableWalker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v16 != 0 {
		v399 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(112)
	return v399
L2:
	;
	v17 = l0
	v21 = v3
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	switch v28 - int32(4) {
	case 0, 1, 10, 11, 12, 13, 14, 37:
		goto L14
	case 2, 3, 15, 16, 26:
		goto L15
	case 4, 5, 6, 7, 8, 9:
		goto L16
	case 17:
		goto L12
	case 18, 21, 27, 28, 29, 30, 31, 32, 34, 35, 36, 39, 40, 42, 43, 44, 45:
		goto L8
	case 19:
		goto L13
	case 20:
		goto L11
	case 22:
		goto L19
	default:
		v375 = v21
		goto L5
	case 24:
		goto L17
	case 25:
		goto L18
	case 33:
		goto L10
	case 38:
		goto L9
	case 41, 46, 47, 48, 49:
		goto L7
	}
L4:
	;
	v399 = v375
	goto L1
L5:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v382 <= int32(0) {
		v399 = v375
		goto L1
	} else {
		goto L125
	}
L6:
	;
	v375 = int32(0)
	goto L5
L7:
	;
	v366 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v366)
	v375 = int32(3)
	goto L5
L8:
	;
	v375 = int32(0)
	goto L5
L9:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v14+int32(56), v357, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L20
	} else {
		goto L123
	}
L10:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v271 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L11:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v269 != 0 {
		goto L84
	} else {
		goto L85
	}
L12:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v266 != 0 {
		goto L81
	} else {
		goto L82
	}
L13:
	;
	v213 = int32(0)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v214 <= v213 {
		goto L12
	} else {
		goto L70
	}
L14:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v14+int32(56), v195, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L66
	}
L15:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v14+int32(56), v185, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L20
	} else {
		goto L64
	}
L16:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v14+int32(56), v154, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L20
	} else {
		goto L54
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v54 = int32(0)
	goto L23
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v21
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v14+int32(56), v36, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v375 = v31
	goto L5
L20:
	;
	return int32(0)
L21:
	;
	v44 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
	v375 = v21
	goto L5
L23:
	;
	v63 = int32(0)
	if v48 == v63 {
		v73 = v63
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v135 = F_exprType(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L20
	} else {
		goto L47
	}
L25:
	;
	if v47 == int32(0) {
		v375 = v21
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v67 <= v54 {
		v73 = int32(0)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v73 = v69 + v54<<(uint(int32(2))%32)
	goto L25
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v76 <= v54 {
		v375 = v21
		goto L5
	} else {
		goto L29
	}
L29:
	;
	if v73 == int32(0) {
		v375 = v21
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v83 = v80 + v54<<(uint(int32(2))%32)
	if v83 == int32(0) {
		v375 = v21
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v50 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v133 != 0 {
		v54 = v54 + int32(1)
		goto L23
	} else {
		goto L46
	}
L33:
	;
	v133 = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v95 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v96 = v89
	v97 = v49
	v98 = v50
	v99 = v95
	goto L40
L37:
	;
	v121 = v49
	v125 = int32(0)
	goto L38
L38:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v133 = v125 - v126
	goto L32
L39:
	;
	v121 = v116
	v125 = v118
	goto L38
L40:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	if v99 != v101 {
		v116 = v97
		v118 = v99
		goto L39
	} else {
		goto L42
	}
L41:
	;
	v116 = v110
	v118 = int32(0)
	goto L39
L42:
	;
	if v101 == int32(0) {
		v116 = v97
		v118 = v99
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v106 = v98 - int32(1)
	if v106 == int32(0) {
		v116 = v97
		v118 = v99
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v109 = int32(1)
	v110 = v97 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
	if v111 != 0 {
		v96 = v96 + v109
		v97 = v110
		v98 = v106
		v99 = v111
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L24
L47:
	;
	if v135 <= int32(1183) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = int32(3)
	if base.Ui32(v135-int32(1082)) < base.Ui32(int32(2)) {
		v375 = v139
		goto L5
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if base.B2i32(v135 != int32(1266))&base.B2i32(v135 != int32(1184)) != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	if v135 == int32(1114) {
		v375 = v139
		goto L5
	} else {
		goto L52
	}
L52:
	;
	goto L6
L53:
	;
	v375 = int32(2)
	goto L5
L54:
	;
	v160 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_jspInitByBuffer(m, v14+int32(56), v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	v170 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L20
	} else {
		goto L57
	}
L57:
	;
	if v160 == int32(0) {
		v375 = v21
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v170 == int32(0) {
		v375 = v21
		goto L5
	} else {
		goto L59
	}
L59:
	;
	if v160 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v181 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v181)
	v375 = v21
	goto L5
L61:
	;
	if v170 == int32(1) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v160 == v170 {
		v375 = v21
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v191 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	v375 = v21
	goto L5
L66:
	;
	v201 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L20
	} else {
		goto L67
	}
L67:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_jspInitByBuffer(m, v14+int32(56), v205, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L68
	}
L68:
	;
	v211 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L20
	} else {
		goto L69
	}
L69:
	;
	v375 = v21
	goto L5
L70:
	;
	v219 = v213
	goto L71
L71:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v232 = v219 << (uint(int32(3)) % 32)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232+v233)))
	F_jspInitByBuffer(m, v14+int32(28), v230, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L20
	} else {
		goto L73
	}
L72:
	;
	goto L12
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v238+v232)+4))
	if v240 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v14, v241, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L20
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v248 = F_jspIsMutableWalker(m, v14+int32(28), l1)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L79
	}
L77:
	;
	v244 = F_jspIsMutableWalker(m, v14, l1)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v251 = v219 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v251 < v252 {
		v219 = v251
		goto L71
	} else {
		goto L80
	}
L80:
	;
	goto L72
L81:
	;
	v267 = v21
	goto L83
L82:
	;
	v267 = int32(0)
	goto L83
L83:
	;
	v375 = v267
	goto L5
L84:
	;
	v270 = int32(0)
	goto L86
L85:
	;
	v270 = v21
	goto L86
L86:
	;
	v375 = v270
	goto L5
L87:
	;
	v375 = int32(1)
	goto L5
L88:
	;
	goto L89
L89:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v14+int32(56), v277, v271)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v280 != int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v375 = int32(0)
	goto L5
L92:
	;
	goto L93
L93:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v285 = F_strlen(m, v284)
	mBase = m.M
	if base.Ui32(v285) <= base.Ui32(int32(155)) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v351 = int32(3)
	if base.Ui32(v351) < base.Ui32(v349) {
		goto L120
	} else {
		goto L121
	}
L95:
	;
	v289 = F_DCH_cache_fetch(m, v284, int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L20
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v312 = int32(12)
	v316 = F_palloc(m, v285*v312+v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L20
	} else {
		goto L108
	}
L98:
	;
	v292 = v289
	v293 = int32(0)
	goto L100
L99:
	;
	v349 = v293
	goto L94
L100:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
	switch v294 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L104
	default:
		v309 = v293
		goto L103
	}
L101:
	;
	goto L99
L102:
	;
	goto L101
L103:
	;
	v292 = v292 + int32(12)
	v293 = v309
	goto L100
L104:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v292)+8))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	switch v298 {
	case 0, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 24, 25, 27, 28, 29, 30, 31, 33, 34, 35, 37, 38, 42, 43, 51, 52, 53, 54, 55, 56, 57, 58, 60, 62, 63, 65, 68, 90, 91, 97:
		goto L105
	case 1, 3, 14, 15, 16, 17, 18, 19, 21, 22, 23, 32, 36, 40, 41, 45, 46, 50, 59, 61, 94, 95:
		goto L107
	default:
		v309 = v293
		goto L103
	case 39, 47, 48, 49, 103:
		goto L106
	}
L105:
	;
	v309 = v293 | int32(1)
	goto L103
L106:
	;
	v292 = v292 + int32(12)
	v293 = v293 | int32(4)
	goto L100
L107:
	;
	v292 = v292 + int32(12)
	v293 = v293 | int32(2)
	goto L100
L108:
	;
	F_parse_format(m, v316, v284, int32(1699168), int32(1698432), int32(1698560), int32(1), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L20
	} else {
		goto L109
	}
L109:
	;
	v326 = v316
	v327 = int32(0)
	goto L111
L110:
	;
	F_pfree(m, v316)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L20
	} else {
		goto L119
	}
L111:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	switch v328 - int32(1) {
	case 0:
		goto L113
	case 1:
		goto L115
	default:
		v343 = v327
		goto L114
	}
L112:
	;
	goto L110
L113:
	;
	goto L112
L114:
	;
	v326 = v326 + int32(12)
	v327 = v343
	goto L111
L115:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v326)+8))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	switch v332 {
	case 0, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 24, 25, 27, 28, 29, 30, 31, 33, 34, 35, 37, 38, 42, 43, 51, 52, 53, 54, 55, 56, 57, 58, 60, 62, 63, 65, 68, 90, 91, 97:
		goto L116
	case 1, 3, 14, 15, 16, 17, 18, 19, 21, 22, 23, 32, 36, 40, 41, 45, 46, 50, 59, 61, 94, 95:
		goto L118
	default:
		v343 = v327
		goto L114
	case 39, 47, 48, 49, 103:
		goto L117
	}
L116:
	;
	v343 = v327 | int32(1)
	goto L114
L117:
	;
	v326 = v326 + int32(12)
	v327 = v327 | int32(4)
	goto L111
L118:
	;
	v326 = v326 + int32(12)
	v327 = v327 | int32(2)
	goto L111
L119:
	;
	v349 = v327
	goto L94
L120:
	;
	v354 = int32(2)
	goto L122
L121:
	;
	v354 = v351
	goto L122
L122:
	;
	v375 = v354
	goto L5
L123:
	;
	v363 = F_jspIsMutableWalker(m, v14+int32(56), l1)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L20
	} else {
		goto L124
	}
L124:
	;
	v375 = v21
	goto L5
L125:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v14+int32(84), v387, v382)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v392 != int32(1) {
		v17 = v14 + int32(84)
		v21 = v375
		goto L3
	} else {
		goto L127
	}
L127:
	;
	goto L4
}
