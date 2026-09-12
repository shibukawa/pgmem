package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyFromBinaryStart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
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
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
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
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v303 int32
	_ = v303
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v14-v15 <= int32(10) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	return
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L14
	} else {
		goto L139
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L14
	} else {
		goto L135
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L14
	} else {
		goto L131
	}
L5:
	;
	v71 = int32(1590744)
	v72 = int32(11)
	goto L29
L6:
	;
	v20 = v12
	v21 = v15
	v23 = int32(0)
	v24 = v14
	goto L9
L7:
	;
	goto L8
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v54 = v53 + v15
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v54)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+7))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+7)) = v57
	v60 = v15 + int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v60
	v64 = v60
	v67 = v14
	v69 = v53
	goto L5
L9:
	;
	if v21 == v24 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v48 != int32(11) {
		goto L4
	} else {
		goto L25
	}
L11:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v34 = v21
	v35 = v24
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v39 = int32(11) - v23
	v40 = v35 - v34
	if v39 < v40 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	return
L15:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v31 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v34 = v33
	v35 = v32
	goto L13
L17:
	;
	v42 = v39
	goto L19
L18:
	;
	v42 = v40
	goto L19
L19:
	;
	if v42 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v45 = v42 + v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v45
	v48 = v42 + v23
	if v48 < int32(11) {
		v20 = v42 + v44
		v21 = v45
		v23 = v48
		v24 = v35
		goto L9
	} else {
		goto L24
	}
L21:
	;
	v43 = F__emscripten_memcpy_bulkmem(m, v20, v36+v34, v42)
	mBase = m.M
	v44 = v43
	goto L23
L22:
	;
	v44 = v20
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L10
L25:
	;
	v64 = v45
	v67 = v35
	v69 = v36
	goto L5
L26:
	;
	if v134 != 0 {
		goto L4
	} else {
		goto L44
	}
L27:
	;
	v134 = int32(0)
	goto L26
L28:
	;
	v108 = v103
	v109 = v104
	v110 = v105
	goto L38
L29:
	;
	if (v12|v71)&int32(3) != 0 {
		v103 = v12
		v104 = v71
		v105 = v72
		goto L28
	} else {
		goto L32
	}
L31:
	;
	if v93 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v80 = v12
	v81 = v71
	v82 = v72
	goto L33
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v85 != v86 {
		v103 = v80
		v104 = v81
		v105 = v82
		goto L28
	} else {
		goto L35
	}
L34:
	;
	goto L31
L35:
	;
	v88 = int32(4)
	v89 = v81 + v88
	v91 = v80 + v88
	v93 = v82 - v88
	if base.Ui32(int32(3)) < base.Ui32(v93) {
		v80 = v91
		v81 = v89
		v82 = v93
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v103 = v91
	v104 = v89
	v105 = v93
	goto L28
L38:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if v113 == v114 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v134 = v113 - v114
	goto L26
L40:
	;
	v116 = int32(1)
	v121 = v110 - v116
	if v121 != 0 {
		v108 = v108 + v116
		v109 = v109 + v116
		v110 = v121
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	if v67-v64 <= int32(3) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v214 = int32(24)
	v216 = int32(65280)
	v218 = int32(8)
	v228 = v205<<(uint(v214)%32) | v205&v216<<(uint(v218)%32) | (int32(base.Ui32(v205)>>(uint(v218)%32))&v216 | int32(base.Ui32(v205)>>(uint(v214)%32)))
	if v228&int32(65536) != 0 {
		goto L3
	} else {
		goto L70
	}
L46:
	;
	v142 = v12 + int32(12)
	v143 = v64
	v145 = int32(0)
	v146 = v67
	v148 = v69
	goto L50
L47:
	;
	goto L48
L48:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v64+v69)))
	v203 = v64 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v203
	v205 = v201
	v207 = v203
	v210 = v67
	v212 = v69
	goto L45
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L14
	} else {
		goto L66
	}
L50:
	;
	if v143 == v146 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v171 != int32(4) {
		goto L49
	} else {
		goto L65
	}
L52:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L55
	}
L53:
	;
	v157 = v143
	v158 = v146
	v159 = v148
	goto L54
L54:
	;
	v162 = int32(4) - v145
	v163 = v158 - v157
	if v162 < v163 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v153 != 0 {
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v157 = v156
	v158 = v154
	v159 = v155
	goto L54
L57:
	;
	v165 = v162
	goto L59
L58:
	;
	v165 = v163
	goto L59
L59:
	;
	if v165 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v168 = v165 + v157
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v168
	v171 = v165 + v145
	if v171 < int32(4) {
		v142 = v165 + v167
		v143 = v168
		v145 = v171
		v146 = v158
		v148 = v159
		goto L50
	} else {
		goto L64
	}
L61:
	;
	v166 = F__emscripten_memcpy_bulkmem(m, v142, v157+v159, v165)
	mBase = m.M
	v167 = v166
	goto L63
L62:
	;
	v167 = v142
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L51
L65:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v205 = v176
	v207 = v168
	v210 = v158
	v212 = v159
	goto L45
L66:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(648255), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(490126), int32(205), int32(224334))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L14
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
	if base.Ui32(int32(65536)) <= base.Ui32(v228) {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	if v210-v207 <= int32(3) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v331 = v303
	v333 = v282
	v336 = v285
	v338 = v287
	goto L99
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L14
	} else {
		goto L95
	}
L74:
	;
	v289 = int32(24)
	v291 = int32(65280)
	v293 = int32(8)
	v303 = v280<<(uint(v289)%32) | v280&v291<<(uint(v293)%32) | (int32(base.Ui32(v280)>>(uint(v293)%32))&v291 | int32(base.Ui32(v280)>>(uint(v289)%32)))
	if v303 < int32(0) {
		goto L73
	} else {
		goto L94
	}
L75:
	;
	v240 = v12 + int32(12)
	v241 = v207
	v243 = int32(0)
	v244 = v210
	v246 = v212
	goto L78
L76:
	;
	goto L77
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v207+v212)))
	v278 = v207 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v278
	v280 = v276
	v282 = v278
	v285 = v210
	v287 = v212
	goto L74
L78:
	;
	if v241 == v244 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v269 != int32(4) {
		goto L73
	} else {
		goto L93
	}
L80:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L14
	} else {
		goto L83
	}
L81:
	;
	v255 = v241
	v256 = v244
	v257 = v246
	goto L82
L82:
	;
	v260 = int32(4) - v243
	v261 = v256 - v255
	if v260 < v261 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v251 != 0 {
		goto L73
	} else {
		goto L84
	}
L84:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v255 = v254
	v256 = v252
	v257 = v253
	goto L82
L85:
	;
	v263 = v260
	goto L87
L86:
	;
	v263 = v261
	goto L87
L87:
	;
	if v263 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v266 = v263 + v255
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v266
	v269 = v263 + v243
	if v269 < int32(4) {
		v240 = v263 + v265
		v241 = v266
		v243 = v269
		v244 = v256
		v246 = v257
		goto L78
	} else {
		goto L92
	}
L89:
	;
	v264 = F__emscripten_memcpy_bulkmem(m, v240, v255+v257, v263)
	mBase = m.M
	v265 = v264
	goto L91
L90:
	;
	v265 = v240
	goto L91
L91:
	;
	goto L88
L92:
	;
	goto L79
L93:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v280 = v274
	v282 = v266
	v285 = v256
	v287 = v257
	goto L74
L94:
	;
	goto L72
L95:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(650792), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(490126), int32(220), int32(224334))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	if int32(0) < v331 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	m.G0 = v12 + int32(16)
	goto L1
L101:
	;
	v342 = v331
	v344 = v333
	goto L104
L102:
	;
	goto L103
L103:
	;
	goto L100
L104:
	;
	v352 = v342 - int32(1)
	if v336-v344 <= int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L103
L106:
	;
	v358 = v12
	v359 = v344
	v361 = int32(0)
	v362 = v336
	v364 = v338
	goto L110
L107:
	;
	goto L108
L108:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344+v338))))
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v415)
	v417 = int32(1)
	v418 = v344 + v417
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v418
	if v417 < v342 {
		v342 = v352
		v344 = v418
		goto L104
	} else {
		goto L130
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L14
	} else {
		goto L126
	}
L110:
	;
	if v359 == v362 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v387 == int32(1) {
		v331 = v352
		v333 = v384
		v336 = v374
		v338 = v375
		goto L99
	} else {
		goto L125
	}
L112:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L14
	} else {
		goto L115
	}
L113:
	;
	v373 = v359
	v374 = v362
	v375 = v364
	goto L114
L114:
	;
	v378 = int32(1) - v361
	v379 = v374 - v373
	if v378 < v379 {
		goto L117
	} else {
		goto L118
	}
L115:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v369 != 0 {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v373 = v372
	v374 = v370
	v375 = v371
	goto L114
L117:
	;
	v381 = v378
	goto L119
L118:
	;
	v381 = v379
	goto L119
L119:
	;
	if v381 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v384 = v381 + v373
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v384
	v387 = v381 + v361
	if v387 <= int32(0) {
		v358 = v381 + v383
		v359 = v384
		v361 = v387
		v362 = v374
		v364 = v375
		goto L110
	} else {
		goto L124
	}
L121:
	;
	v382 = F__emscripten_memcpy_bulkmem(m, v358, v373+v375, v381)
	mBase = m.M
	v383 = v382
	goto L123
L122:
	;
	v383 = v358
	goto L123
L123:
	;
	goto L120
L124:
	;
	goto L111
L125:
	;
	goto L109
L126:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	F_errmsg(m, int32(650752), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L14
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(490126), int32(227), int32(224334))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L14
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	goto L105
L131:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(430491), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(490126), int32(200), int32(224334))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L14
	} else {
		goto L136
	}
L136:
	;
	F_errmsg(m, int32(656591), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(490126), int32(209), int32(224334))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L14
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
	F_errcode(m, int32(67240066))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L14
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(224123), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L14
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(490126), int32(214), int32(224334))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L14
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_from_char_seq_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v639 int32
	_ = v639
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v770 int32
	_ = v770
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if l3 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v770
L2:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v759 + v667
	v770 = int32(1)
	goto L1
L3:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v689 = F_pstrdup(m, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L96
	} else {
		goto L199
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = (v664 - v658) >> (uint(int32(2)) % 32)
	if int32(0) < v667 {
		goto L2
	} else {
		goto L197
	}
L5:
	;
	v658 = l2
	v664 = v46
	v667 = v77 - v18
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(-1)
	goto L3
L7:
	;
	v23 = v19 & int32(255)
	if v23 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v19&int32(255) == int32(0) {
		goto L6
	} else {
		goto L38
	}
L10:
	;
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v35 == int32(0) {
		goto L6
	} else {
		goto L15
	}
L12:
	;
	v34 = v23 | int32(32)
	goto L14
L13:
	;
	v34 = v23
	goto L14
L14:
	;
	goto L11
L15:
	;
	v45 = v35
	v46 = l2
	goto L16
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if base.Ui32((v51-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L6
L18:
	;
	v120 = v46 + int32(4)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 != 0 {
		v45 = v121
		v46 = v120
		goto L16
	} else {
		goto L37
	}
L19:
	;
	if v60 != v34 {
		goto L18
	} else {
		goto L23
	}
L20:
	;
	v60 = v51 | int32(32)
	goto L22
L21:
	;
	v60 = v51
	goto L22
L22:
	;
	goto L19
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v70 = v18
	v73 = v62
	goto L24
L24:
	;
	v76 = int32(1)
	v77 = v70 + v76
	v79 = v73 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 == int32(0) {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	goto L18
L26:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v80-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if base.Ui32((v95-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v94 = v80 | int32(32)
	goto L31
L30:
	;
	v94 = v80
	goto L31
L31:
	;
	goto L28
L32:
	;
	if v94 == v104 {
		v70 = v77
		v73 = v79
		goto L24
	} else {
		goto L36
	}
L33:
	;
	v104 = v95 | int32(32)
	goto L35
L34:
	;
	v104 = v95
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L25
L37:
	;
	goto L17
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v126 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v134 = v126
	v135 = l3
	goto L42
L40:
	;
	goto L41
L41:
	;
	if v18&int32(3) == int32(0) {
		v282 = v18
		goto L81
	} else {
		goto L82
	}
L42:
	;
	if v134&int32(3) == int32(0) {
		v163 = v134
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L41
L44:
	;
	if v196 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	v196 = v188 - v134
	goto L44
L46:
	;
	v167 = v163
	goto L55
L47:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v147 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v196 = int32(0)
	goto L44
L49:
	;
	goto L50
L50:
	;
	v152 = v134
	goto L51
L51:
	;
	v156 = v152 + int32(1)
	if v156&int32(3) == int32(0) {
		v163 = v156
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v188 = v156
	goto L45
L53:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if v161 != 0 {
		v152 = v156
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v176 = int32(-2139062144)
	if (int32(16843008)-v173|v173)&v176 == v176 {
		v167 = v167 + int32(4)
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v182 = v167
	goto L58
L57:
	;
	goto L56
L58:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if v186 != 0 {
		v182 = v182 + int32(1)
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v188 = v182
	goto L45
L60:
	;
	goto L59
L61:
	;
	if v240 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L62:
	;
	v240 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v202 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v203 = v18
	v204 = v134
	v205 = v196
	v206 = v202
	goto L69
L66:
	;
	v228 = v134
	v232 = int32(0)
	goto L67
L67:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	v240 = v232 - v233
	goto L61
L68:
	;
	v228 = v223
	v232 = v225
	goto L67
L69:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v206 != v208 {
		v223 = v204
		v225 = v206
		goto L68
	} else {
		goto L71
	}
L70:
	;
	v223 = v217
	v225 = int32(0)
	goto L68
L71:
	;
	if v208 == int32(0) {
		v223 = v204
		v225 = v206
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v213 = v205 - int32(1)
	if v213 == int32(0) {
		v223 = v204
		v225 = v206
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v216 = int32(1)
	v217 = v204 + v216
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	if v218 != 0 {
		v203 = v203 + v216
		v204 = v217
		v205 = v213
		v206 = v218
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L70
L75:
	;
	v658 = l3
	v664 = v135
	v667 = v196
	goto L4
L76:
	;
	goto L77
L77:
	;
	v244 = v135 + int32(4)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	if v245 != 0 {
		v134 = v245
		v135 = v244
		goto L42
	} else {
		goto L78
	}
L78:
	;
	goto L43
L79:
	;
	v316 = F_str_toupper(m, v18, v315, l4)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L96
	} else {
		goto L97
	}
L80:
	;
	v315 = v307 - v18
	goto L79
L81:
	;
	v286 = v282
	goto L90
L82:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v266 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v315 = int32(0)
	goto L79
L84:
	;
	goto L85
L85:
	;
	v271 = v18
	goto L86
L86:
	;
	v275 = v271 + int32(1)
	if v275&int32(3) == int32(0) {
		v282 = v275
		goto L81
	} else {
		goto L88
	}
L87:
	;
	v307 = v275
	goto L80
L88:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v280 != 0 {
		v271 = v275
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	v295 = int32(-2139062144)
	if (int32(16843008)-v292|v292)&v295 == v295 {
		v286 = v286 + int32(4)
		goto L90
	} else {
		goto L92
	}
L91:
	;
	v301 = v286
	goto L93
L92:
	;
	goto L91
L93:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	if v305 != 0 {
		v301 = v301 + int32(1)
		goto L93
	} else {
		goto L95
	}
L94:
	;
	v307 = v301
	goto L80
L95:
	;
	goto L94
L96:
	;
	return int32(0)
L97:
	;
	if v316&int32(3) == int32(0) {
		v343 = v316
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v377 = F_str_tolower(m, v316, v376, l4)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L96
	} else {
		goto L115
	}
L99:
	;
	v376 = v368 - v316
	goto L98
L100:
	;
	v347 = v343
	goto L109
L101:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	if v327 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v376 = int32(0)
	goto L98
L103:
	;
	goto L104
L104:
	;
	v332 = v316
	goto L105
L105:
	;
	v336 = v332 + int32(1)
	if v336&int32(3) == int32(0) {
		v343 = v336
		goto L100
	} else {
		goto L107
	}
L106:
	;
	v368 = v336
	goto L99
L107:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	if v341 != 0 {
		v332 = v336
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v356 = int32(-2139062144)
	if (int32(16843008)-v353|v353)&v356 == v356 {
		v347 = v347 + int32(4)
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v362 = v347
	goto L112
L111:
	;
	goto L110
L112:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v366 != 0 {
		v362 = v362 + int32(1)
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v368 = v362
	goto L99
L114:
	;
	goto L113
L115:
	;
	F_pfree(m, v316)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L96
	} else {
		goto L116
	}
L116:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v381 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v389 = v381
	v390 = l3
	goto L120
L118:
	;
	goto L119
L119:
	;
	F_pfree(m, v377)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L96
	} else {
		goto L196
	}
L120:
	;
	if v389&int32(3) == int32(0) {
		v418 = v389
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L119
L122:
	;
	v452 = F_str_toupper(m, v389, v451, l4)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L96
	} else {
		goto L139
	}
L123:
	;
	v451 = v443 - v389
	goto L122
L124:
	;
	v422 = v418
	goto L133
L125:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v402 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v451 = int32(0)
	goto L122
L127:
	;
	goto L128
L128:
	;
	v407 = v389
	goto L129
L129:
	;
	v411 = v407 + int32(1)
	if v411&int32(3) == int32(0) {
		v418 = v411
		goto L124
	} else {
		goto L131
	}
L130:
	;
	v443 = v411
	goto L123
L131:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v416 != 0 {
		v407 = v411
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	v431 = int32(-2139062144)
	if (int32(16843008)-v428|v428)&v431 == v431 {
		v422 = v422 + int32(4)
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v437 = v422
	goto L136
L135:
	;
	goto L134
L136:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437))))
	if v441 != 0 {
		v437 = v437 + int32(1)
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v443 = v437
	goto L123
L138:
	;
	goto L137
L139:
	;
	if v452&int32(3) == int32(0) {
		v477 = v452
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v511 = F_str_tolower(m, v452, v510, l4)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L96
	} else {
		goto L157
	}
L141:
	;
	v510 = v502 - v452
	goto L140
L142:
	;
	v481 = v477
	goto L151
L143:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v461 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v510 = int32(0)
	goto L140
L145:
	;
	goto L146
L146:
	;
	v466 = v452
	goto L147
L147:
	;
	v470 = v466 + int32(1)
	if v470&int32(3) == int32(0) {
		v477 = v470
		goto L142
	} else {
		goto L149
	}
L148:
	;
	v502 = v470
	goto L141
L149:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	if v475 != 0 {
		v466 = v470
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v481)))
	v490 = int32(-2139062144)
	if (int32(16843008)-v487|v487)&v490 == v490 {
		v481 = v481 + int32(4)
		goto L151
	} else {
		goto L153
	}
L152:
	;
	v496 = v481
	goto L154
L153:
	;
	goto L152
L154:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496))))
	if v500 != 0 {
		v496 = v496 + int32(1)
		goto L154
	} else {
		goto L156
	}
L155:
	;
	v502 = v496
	goto L141
L156:
	;
	goto L155
L157:
	;
	F_pfree(m, v452)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L96
	} else {
		goto L158
	}
L158:
	;
	if v511&int32(3) == int32(0) {
		v538 = v511
		goto L161
	} else {
		goto L162
	}
L159:
	;
	if v571 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L160:
	;
	v571 = v563 - v511
	goto L159
L161:
	;
	v542 = v538
	goto L170
L162:
	;
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511))))
	if v522 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v571 = int32(0)
	goto L159
L164:
	;
	goto L165
L165:
	;
	v527 = v511
	goto L166
L166:
	;
	v531 = v527 + int32(1)
	if v531&int32(3) == int32(0) {
		v538 = v531
		goto L161
	} else {
		goto L168
	}
L167:
	;
	v563 = v531
	goto L160
L168:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	if v536 != 0 {
		v527 = v531
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	v551 = int32(-2139062144)
	if (int32(16843008)-v548|v548)&v551 == v551 {
		v542 = v542 + int32(4)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v557 = v542
	goto L173
L172:
	;
	goto L171
L173:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	if v561 != 0 {
		v557 = v557 + int32(1)
		goto L173
	} else {
		goto L175
	}
L174:
	;
	v563 = v557
	goto L160
L175:
	;
	goto L174
L176:
	;
	F_pfree(m, v511)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L96
	} else {
		goto L190
	}
L177:
	;
	v615 = int32(0)
	goto L176
L178:
	;
	goto L179
L179:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377))))
	if v577 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v578 = v377
	v579 = v511
	v580 = v571
	v581 = v577
	goto L184
L181:
	;
	v603 = v511
	v607 = int32(0)
	goto L182
L182:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	v615 = v607 - v608
	goto L176
L183:
	;
	v603 = v598
	v607 = v600
	goto L182
L184:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579))))
	if v581 != v583 {
		v598 = v579
		v600 = v581
		goto L183
	} else {
		goto L186
	}
L185:
	;
	v598 = v592
	v600 = int32(0)
	goto L183
L186:
	;
	if v583 == int32(0) {
		v598 = v579
		v600 = v581
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v588 = v580 - int32(1)
	if v588 == int32(0) {
		v598 = v579
		v600 = v581
		goto L183
	} else {
		goto L188
	}
L188:
	;
	v591 = int32(1)
	v592 = v579 + v591
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578)+1)))
	if v593 != 0 {
		v578 = v578 + v591
		v579 = v592
		v580 = v588
		v581 = v593
		goto L184
	} else {
		goto L189
	}
L189:
	;
	goto L185
L190:
	;
	if v615 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_pfree(m, v377)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L96
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v623 = v390 + int32(4)
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	if v624 != 0 {
		v389 = v624
		v390 = v623
		goto L120
	} else {
		goto L195
	}
L194:
	;
	v658 = l3
	v664 = v390
	v667 = v571
	goto L4
L195:
	;
	goto L121
L196:
	;
	goto L6
L197:
	;
	goto L3
L198:
	;
	v735 = int32(0)
	v736 = F_errsave_start(m, l6)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L96
	} else {
		goto L208
	}
L199:
	;
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v689))))
	if v691 == int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v701 = v691
	v702 = v689
	goto L201
L201:
	;
	v707 = base.I32_extend8_s(v701)
	goto L203
L202:
	;
	goto L198
L203:
	;
	if base.B2i32(v707 == int32(32))|base.B2i32(base.Ui32((v707-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v717 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v702))) = uint8(v717)
	goto L198
L205:
	;
	goto L206
L206:
	;
	v720 = v702 + int32(1)
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	if v721 != 0 {
		v701 = v721
		v702 = v720
		goto L201
	} else {
		goto L207
	}
L207:
	;
	goto L202
L208:
	;
	if v736 == int32(0) {
		v770 = v735
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L96
	} else {
		goto L210
	}
L210:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v689
	F_errmsg(m, int32(675001), v16)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L96
	} else {
		goto L211
	}
L211:
	;
	F_errdetail(m, int32(620919), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L96
	} else {
		goto L212
	}
L212:
	;
	F_errsave_finish(m, l6, int32(489627), int32(2493), int32(320483))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L96
	} else {
		goto L213
	}
L213:
	;
	v770 = v735
	goto L1
}
