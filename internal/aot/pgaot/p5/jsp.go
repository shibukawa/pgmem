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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v16 != 0 {
		v388 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(112)
	return v388
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
		v366 = v21
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
	v388 = v366
	goto L1
L5:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v373 <= int32(0) {
		v388 = v366
		goto L1
	} else {
		goto L119
	}
L6:
	;
	v366 = int32(0)
	goto L5
L7:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v357)
	v366 = int32(3)
	goto L5
L8:
	;
	v366 = int32(0)
	goto L5
L9:
	;
	v349 = v14 + int32(56)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v349, v350, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L20
	} else {
		goto L117
	}
L10:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v264 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L11:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v262 != 0 {
		goto L78
	} else {
		goto L79
	}
L12:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v259 != 0 {
		goto L75
	} else {
		goto L76
	}
L13:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v207 <= v206 {
		goto L12
	} else {
		goto L64
	}
L14:
	;
	v193 = v14 + int32(56)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v193, v194, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L20
	} else {
		goto L60
	}
L15:
	;
	v185 = v14 + int32(56)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v185, v186, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L58
	}
L16:
	;
	v154 = v14 + int32(56)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v154, v155, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L20
	} else {
		goto L53
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v52 = int32(0)
	goto L23
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v21
	v35 = v14 + int32(56)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_jspInitByBuffer(m, v35, v36, v37)
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
	v366 = v31
	goto L5
L20:
	;
	return int32(0)
L21:
	;
	v42 = F_jspIsMutableWalker(m, v35, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v32
	v366 = v21
	goto L5
L23:
	;
	v61 = int32(0)
	if v46 == v61 {
		v71 = v61
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v45 == int32(0) {
		v366 = v21
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v65 <= v52 {
		v71 = int32(0)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v71 = v67 + v52<<(uint(int32(2))%32)
	goto L25
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.B2i32(v71 == int32(0))|base.B2i32(v76 <= v52) != 0 {
		v366 = v21
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v79 == int32(0) {
		v366 = v21
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v48 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v128 != 0 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v128 = int32(0)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v89 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v90 = v83
	v91 = v47
	v92 = v48
	v93 = v89
	goto L39
L36:
	;
	v116 = v47
	v120 = int32(0)
	goto L37
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v128 = v120 - v121
	goto L31
L38:
	;
	v116 = v111
	v120 = v113
	goto L37
L39:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if base.B2i32(v93 != v95)|base.B2i32(v95 == int32(0)) != 0 {
		v111 = v91
		v113 = v93
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v111 = v105
	v113 = int32(0)
	goto L38
L41:
	;
	v101 = v92 - int32(1)
	if v101 == int32(0) {
		v111 = v91
		v113 = v93
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v104 = int32(1)
	v105 = v91 + v104
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v106 != 0 {
		v90 = v90 + v104
		v91 = v105
		v92 = v101
		v93 = v106
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v52 = v52 + int32(1)
	goto L23
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v79+v52<<(uint(int32(2))%32))))
	v135 = F_exprType(m, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	if v135 <= int32(1183) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if base.B2i32(v135 == int32(1114))|base.B2i32(base.Ui32(v135-int32(1082)) < base.Ui32(int32(2))) != 0 {
		v366 = int32(3)
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
		goto L52
	}
L51:
	;
	goto L6
L52:
	;
	v366 = int32(2)
	goto L5
L53:
	;
	v159 = F_jspIsMutableWalker(m, v154, l1)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_jspInitByBuffer(m, v154, v161, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L20
	} else {
		goto L55
	}
L55:
	;
	v167 = F_jspIsMutableWalker(m, v154, l1)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	v169 = int32(0)
	v172 = int32(1)
	if base.B2i32(v159 == int32(0))|base.B2i32(v167 == v169)|base.B2i32(base.B2i32(v159 == v172)|base.B2i32(v167 == v172) == v169)&base.B2i32(v159 == v167) != 0 {
		v366 = v21
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)) = uint8(v182)
	v366 = v21
	goto L5
L58:
	;
	v190 = F_jspIsMutableWalker(m, v185, l1)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	v366 = v21
	goto L5
L60:
	;
	v198 = F_jspIsMutableWalker(m, v193, l1)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	F_jspInitByBuffer(m, v193, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	v204 = F_jspIsMutableWalker(m, v193, l1)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	v366 = v21
	goto L5
L64:
	;
	v212 = v206
	goto L65
L65:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v225 = v212 << (uint(int32(3)) % 32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225+v226)))
	F_jspInitByBuffer(m, v14+int32(28), v223, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L67
	}
L66:
	;
	goto L12
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231+v225)+4))
	if v233 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v14, v234, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L20
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v241 = F_jspIsMutableWalker(m, v14+int32(28), l1)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L20
	} else {
		goto L73
	}
L71:
	;
	v237 = F_jspIsMutableWalker(m, v14, l1)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v244 = v212 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v244 < v245 {
		v212 = v244
		goto L65
	} else {
		goto L74
	}
L74:
	;
	goto L66
L75:
	;
	v260 = v21
	goto L77
L76:
	;
	v260 = int32(0)
	goto L77
L77:
	;
	v366 = v260
	goto L5
L78:
	;
	v263 = int32(0)
	goto L80
L79:
	;
	v263 = v21
	goto L80
L80:
	;
	v366 = v263
	goto L5
L81:
	;
	v366 = int32(1)
	goto L5
L82:
	;
	goto L83
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v14+int32(56), v270, v264)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L20
	} else {
		goto L84
	}
L84:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	if v273 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v366 = int32(0)
	goto L5
L86:
	;
	goto L87
L87:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v278 = F_strlen(m, v277)
	mBase = m.M
	if base.Ui32(v278) <= base.Ui32(int32(155)) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v344 = int32(3)
	if base.Ui32(v344) < base.Ui32(v342) {
		goto L114
	} else {
		goto L115
	}
L89:
	;
	v282 = F_DCH_cache_fetch(m, v277, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L20
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v305 = int32(12)
	v309 = F_palloc(m, v278*v305+v305)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L20
	} else {
		goto L102
	}
L92:
	;
	v285 = v282
	v286 = int32(0)
	goto L94
L93:
	;
	v342 = v286
	goto L88
L94:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	switch v287 - int32(1) {
	case 0:
		goto L96
	case 1:
		goto L98
	default:
		v302 = v286
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	goto L95
L97:
	;
	v285 = v285 + int32(12)
	v286 = v302
	goto L94
L98:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v285)+8))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+8))
	switch v291 {
	case 0, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 24, 25, 27, 28, 29, 30, 31, 33, 34, 35, 37, 38, 42, 43, 51, 52, 53, 54, 55, 56, 57, 58, 60, 62, 63, 65, 68, 90, 91, 97:
		goto L99
	case 1, 3, 14, 15, 16, 17, 18, 19, 21, 22, 23, 32, 36, 40, 41, 45, 46, 50, 59, 61, 94, 95:
		goto L101
	default:
		v302 = v286
		goto L97
	case 39, 47, 48, 49, 103:
		goto L100
	}
L99:
	;
	v302 = v286 | int32(1)
	goto L97
L100:
	;
	v285 = v285 + int32(12)
	v286 = v286 | int32(4)
	goto L94
L101:
	;
	v285 = v285 + int32(12)
	v286 = v286 | int32(2)
	goto L94
L102:
	;
	F_parse_format(m, v309, v277, int32(_a_F_jspIsMutableWalker_0), int32(_a_F_jspIsMutableWalker_1), int32(_a_F_jspIsMutableWalker_2), int32(1), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L20
	} else {
		goto L103
	}
L103:
	;
	v319 = v309
	v320 = int32(0)
	goto L105
L104:
	;
	F_pfree(m, v309)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L20
	} else {
		goto L113
	}
L105:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	switch v321 - int32(1) {
	case 0:
		goto L107
	case 1:
		goto L109
	default:
		v336 = v320
		goto L108
	}
L106:
	;
	goto L104
L107:
	;
	goto L106
L108:
	;
	v319 = v319 + int32(12)
	v320 = v336
	goto L105
L109:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	switch v325 {
	case 0, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 24, 25, 27, 28, 29, 30, 31, 33, 34, 35, 37, 38, 42, 43, 51, 52, 53, 54, 55, 56, 57, 58, 60, 62, 63, 65, 68, 90, 91, 97:
		goto L110
	case 1, 3, 14, 15, 16, 17, 18, 19, 21, 22, 23, 32, 36, 40, 41, 45, 46, 50, 59, 61, 94, 95:
		goto L112
	default:
		v336 = v320
		goto L108
	case 39, 47, 48, 49, 103:
		goto L111
	}
L110:
	;
	v336 = v320 | int32(1)
	goto L108
L111:
	;
	v319 = v319 + int32(12)
	v320 = v320 | int32(4)
	goto L105
L112:
	;
	v319 = v319 + int32(12)
	v320 = v320 | int32(2)
	goto L105
L113:
	;
	v342 = v320
	goto L88
L114:
	;
	v347 = int32(2)
	goto L116
L115:
	;
	v347 = v344
	goto L116
L116:
	;
	v366 = v347
	goto L5
L117:
	;
	v354 = F_jspIsMutableWalker(m, v349, l1)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L20
	} else {
		goto L118
	}
L118:
	;
	v366 = v21
	goto L5
L119:
	;
	v377 = v14 + int32(84)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_jspInitByBuffer(m, v377, v378, v373)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v381 != int32(1) {
		v17 = v377
		v21 = v366
		goto L3
	} else {
		goto L121
	}
L121:
	;
	goto L4
}
