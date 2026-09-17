package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___get_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
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
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v6 != 0 {
		v161 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v164 = int32(0)
	goto L64
L2:
	;
	v7 = int32(_a_F___get_locale_0)
	v8 = int32(0)
	v13 = F___strchrnul(m, v7, int32(61))
	mBase = m.M
	if v7 == v13 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v55 != 0 {
		goto L19
	} else {
		goto L20
	}
L4:
	;
	v55 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = v13 - v7
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F___get_locale[0]))))
	if v18 != 0 {
		v49 = v8
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v55 = v49
	goto L3
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F___get_locale[1]))
	if v20 == int32(0) {
		v49 = v8
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v23 == int32(0) {
		v49 = v8
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v27 = v20
	v28 = v23
	goto L11
L11:
	;
	v31 = F_strncmp(m, v7, v28, v16)
	mBase = m.M
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v49 = v35 + int32(1)
	goto L7
L13:
	;
	goto L12
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v35 = v34 + v16
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 == int32(61) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v40 != 0 {
		v27 = v27 + int32(4)
		v28 = v40
		goto L11
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v49 = v8
	goto L7
L19:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != 0 {
		v161 = v55
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v60 = l0*int32(12) + int32(_a_F___get_locale_1)
	v61 = int32(0)
	v66 = F___strchrnul(m, v60, int32(61))
	mBase = m.M
	if v60 == v66 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	if v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L24:
	;
	v108 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v69 = v66 - v60
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v69))))
	if v71 != 0 {
		v102 = v61
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v108 = v102
	goto L23
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F___get_locale[1]))
	if v73 == int32(0) {
		v102 = v61
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v76 == int32(0) {
		v102 = v61
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v80 = v73
	v81 = v76
	goto L31
L31:
	;
	v84 = F_strncmp(m, v60, v81, v69)
	mBase = m.M
	if v84 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v102 = v88 + int32(1)
	goto L27
L33:
	;
	goto L32
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v88 = v87 + v69
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v89 == int32(61) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v93 != 0 {
		v80 = v80 + int32(4)
		v81 = v93
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v102 = v61
	goto L27
L39:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 != 0 {
		v161 = v108
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v110 = int32(_a_F___get_locale_2)
	v111 = int32(0)
	v116 = F___strchrnul(m, v110, int32(61))
	mBase = m.M
	if v110 == v116 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L41
L43:
	;
	if v158 != 0 {
		goto L59
	} else {
		goto L60
	}
L44:
	;
	v158 = int32(0)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v119 = v116 - v110
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+uint32(_c_F___get_locale[2]))))
	if v121 != 0 {
		v152 = v111
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v158 = v152
	goto L43
L48:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F___get_locale[1]))
	if v123 == int32(0) {
		v152 = v111
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v126 == int32(0) {
		v152 = v111
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v130 = v123
	v131 = v126
	goto L51
L51:
	;
	v134 = F_strncmp(m, v110, v131, v119)
	mBase = m.M
	if v134 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v152 = v138 + int32(1)
	goto L47
L53:
	;
	goto L52
L54:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v138 = v137 + v119
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v139 == int32(61) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v143 != 0 {
		v130 = v130 + int32(4)
		v131 = v143
		goto L51
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v152 = v111
	goto L47
L59:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 != 0 {
		v161 = v158
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v161 = int32(_a_F___get_locale_3)
	goto L1
L62:
	;
	goto L61
L63:
	;
	v183 = int32(_a_F___get_locale_3)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v184 == int32(46) {
		v191 = v183
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v164))))
	v169 = int32(0)
	if base.B2i32(v168 == v169)|base.B2i32(v168 == int32(47)) == v169 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v182 = v164
	goto L63
L66:
	;
	v176 = int32(23)
	v178 = v164 + int32(1)
	if v178 != v176 {
		v164 = v178
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	v182 = v176
	goto L63
L70:
	;
	return v495
L71:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F___get_locale[3]))
	if v263 != 0 {
		goto L99
	} else {
		goto L100
	}
L72:
	;
	if l0 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L73:
	;
	v196 = int32(_a_F___get_locale_3)
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___get_locale[4])))
	if base.B2i32(v199 == int32(0))|base.B2i32(v199 != v202) != 0 {
		v220 = v199
		v221 = v202
		goto L80
	} else {
		goto L81
	}
L74:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v192 == int32(0) {
		v252 = v191
		goto L72
	} else {
		goto L78
	}
L75:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161+v182))))
	if v188 != 0 {
		v191 = v183
		goto L74
	} else {
		goto L76
	}
L76:
	;
	if v184 != int32(67) {
		v195 = v161
		goto L73
	} else {
		goto L77
	}
L77:
	;
	v191 = v161
	goto L74
L78:
	;
	v195 = v191
	goto L73
L79:
	;
	if v220-v221 == int32(0) {
		v252 = v195
		goto L72
	} else {
		goto L86
	}
L80:
	;
	goto L79
L81:
	;
	v205 = v195
	v206 = v196
	goto L82
L82:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if v210 == int32(0) {
		v220 = v210
		v221 = v209
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v220 = v210
	v221 = v209
	goto L80
L84:
	;
	v213 = int32(1)
	if v210 == v209 {
		v205 = v205 + v213
		v206 = v206 + v213
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v225 = int32(_a_F___get_locale_4)
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, _c_F___get_locale[5])))
	if base.B2i32(v228 == int32(0))|base.B2i32(v228 != v231) != 0 {
		v249 = v228
		v250 = v231
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v249-v250 != 0 {
		goto L71
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v234 = v195
	v235 = v225
	goto L90
L90:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	if v239 == int32(0) {
		v249 = v239
		v250 = v238
		goto L88
	} else {
		goto L92
	}
L91:
	;
	v249 = v239
	v250 = v238
	goto L88
L92:
	;
	v242 = int32(1)
	if v239 == v238 {
		v234 = v234 + v242
		v235 = v235 + v242
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v252 = v195
	goto L72
L95:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	if v256 == int32(46) {
		v495 = int32(_a_F___get_locale_5)
		goto L70
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	return int32(0)
L98:
	;
	goto L97
L99:
	;
	v266 = v263
	goto L102
L100:
	;
	goto L101
L101:
	;
	v306 = F_emscripten_builtin_malloc(m, int32(36))
	mBase = m.M
	if v306 != 0 {
		goto L113
	} else {
		goto L114
	}
L102:
	;
	v270 = v266 + int32(8)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if base.B2i32(v273 == int32(0))|base.B2i32(v273 != v276) != 0 {
		v294 = v273
		v295 = v276
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L101
L104:
	;
	if v294-v295 == int32(0) {
		v495 = v266
		goto L70
	} else {
		goto L111
	}
L105:
	;
	goto L104
L106:
	;
	v279 = v195
	v280 = v270
	goto L107
L107:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+1)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	if v284 == int32(0) {
		v294 = v284
		v295 = v283
		goto L105
	} else {
		goto L109
	}
L108:
	;
	v294 = v284
	v295 = v283
	goto L105
L109:
	;
	v287 = int32(1)
	if v284 == v283 {
		v279 = v279 + v287
		v280 = v280 + v287
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v266)+32))
	if v299 != 0 {
		v266 = v299
		goto L102
	} else {
		goto L112
	}
L112:
	;
	goto L103
L113:
	;
	v308 = *(*int64)(unsafe.Add(mBase, _c_F___get_locale[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v306))) = v308
	v311 = v306 + int32(8)
	if base.Ui32(int32(512)) <= base.Ui32(v182) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	if l0|v306 != 0 {
		goto L163
	} else {
		goto L164
	}
L116:
	;
	v482 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v311+v182))) = uint8(v482)
	v484 = int32(_a_F___get_locale_6)
	v485 = *(*int32)(unsafe.Add(mBase, _c_F___get_locale[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+32)) = v485
	*(*int32)(unsafe.Add(mBase, _c_F___get_locale[3])) = v306
	goto L115
L117:
	;
	if v182 != 0 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v318 = v311 + v182
	if (v311^v195)&int32(3) == int32(0) {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	base.MemoryCopy(m, v311, v195, v182)
	goto L122
L121:
	;
	goto L122
L122:
	;
	goto L116
L123:
	;
	if base.Ui32(v450) < base.Ui32(v318) {
		goto L157
	} else {
		goto L158
	}
L124:
	;
	if v311&int32(3) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	goto L126
L126:
	;
	if base.Ui32(v318) < base.Ui32(int32(4)) {
		goto L148
	} else {
		goto L149
	}
L127:
	;
	v354 = v318 & int32(-4)
	if base.Ui32(v318) < base.Ui32(int32(64)) {
		v404 = v348
		v405 = v349
		goto L138
	} else {
		goto L139
	}
L128:
	;
	v348 = v195
	v349 = v311
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v182 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v348 = v195
	v349 = v311
	goto L127
L132:
	;
	goto L133
L133:
	;
	v331 = v195
	v332 = v311
	goto L134
L134:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v336)
	v338 = int32(1)
	v339 = v331 + v338
	v341 = v332 + v338
	if v341&int32(3) == int32(0) {
		v348 = v339
		v349 = v341
		goto L127
	} else {
		goto L136
	}
L135:
	;
	v348 = v339
	v349 = v341
	goto L127
L136:
	;
	if base.Ui32(v341) < base.Ui32(v318) {
		v331 = v339
		v332 = v341
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	if base.Ui32(v354) <= base.Ui32(v405) {
		v449 = v404
		v450 = v405
		goto L123
	} else {
		goto L144
	}
L139:
	;
	v358 = v354 + int32(-64)
	if base.Ui32(v358) < base.Ui32(v349) {
		v404 = v348
		v405 = v349
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v361 = v348
	v362 = v349
	goto L141
L141:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = v366
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = v368
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+8)) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+12)) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v361)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+16)) = v374
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v361)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+20)) = v376
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+24)) = v378
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+28)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v361)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+32)) = v382
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v361)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+36)) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v361)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+40)) = v386
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v361)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+44)) = v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v361)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+48)) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v361)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+52)) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v361)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+56)) = v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v361)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+60)) = v396
	v398 = int32(-64)
	v399 = v361 - v398
	v401 = v362 - v398
	if base.Ui32(v401) <= base.Ui32(v358) {
		v361 = v399
		v362 = v401
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v404 = v399
	v405 = v401
	goto L138
L143:
	;
	goto L142
L144:
	;
	v411 = v404
	v412 = v405
	goto L145
L145:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	*(*int32)(unsafe.Add(mBase, uint32(v412))) = v416
	v418 = int32(4)
	v419 = v411 + v418
	v421 = v412 + v418
	if base.Ui32(v421) < base.Ui32(v354) {
		v411 = v419
		v412 = v421
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v449 = v419
	v450 = v421
	goto L123
L147:
	;
	goto L146
L148:
	;
	v449 = v195
	v450 = v311
	goto L123
L149:
	;
	goto L150
L150:
	;
	if base.Ui32(v182) < base.Ui32(int32(4)) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v449 = v195
	v450 = v311
	goto L123
L152:
	;
	goto L153
L153:
	;
	v430 = v195
	v431 = v311
	goto L154
L154:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	*(*uint8)(unsafe.Add(mBase, uint32(v431))) = uint8(v435)
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+1)) = uint8(v437)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+2)) = uint8(v439)
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+3)) = uint8(v441)
	v443 = int32(4)
	v444 = v430 + v443
	v446 = v431 + v443
	if base.Ui32(v446) <= base.Ui32(v318-int32(4)) {
		v430 = v444
		v431 = v446
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v449 = v444
	v450 = v446
	goto L123
L156:
	;
	goto L155
L157:
	;
	v456 = v449
	v457 = v450
	goto L160
L158:
	;
	goto L159
L159:
	;
	goto L116
L160:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v461)
	v463 = int32(1)
	v466 = v457 + v463
	if v466 != v318 {
		v456 = v456 + v463
		v457 = v466
		goto L160
	} else {
		goto L162
	}
L161:
	;
	goto L159
L162:
	;
	goto L161
L163:
	;
	v492 = v306
	goto L165
L164:
	;
	v492 = int32(_a_F___get_locale_5)
	goto L165
L165:
	;
	v495 = v492
	goto L70
}
func F_check_locale_numeric(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = F_check_locale(m, int32(1), v5, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
