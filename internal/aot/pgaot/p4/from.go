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
	v71 = int32(1616392)
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
	F_errmsg(m, int32(673555), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L14
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(498648), int32(205), int32(227870))
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
	F_errmsg(m, int32(676092), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(498648), int32(220), int32(227870))
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
	F_errmsg(m, int32(676052), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L14
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(498648), int32(227), int32(227870))
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
	F_errmsg(m, int32(437595), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(498648), int32(200), int32(227870))
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
	F_errmsg(m, int32(681951), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(498648), int32(209), int32(227870))
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
	F_errmsg(m, int32(227659), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L14
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(498648), int32(214), int32(227870))
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
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v303 int32
	_ = v303
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
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
	return v434
L2:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v423 + v331
	v434 = int32(1)
	goto L1
L3:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v353 = F_pstrdup(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L62
	} else {
		goto L97
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = (v328 - v322) >> (uint(int32(2)) % 32)
	if int32(0) < v331 {
		goto L2
	} else {
		goto L95
	}
L5:
	;
	v322 = l2
	v328 = v46
	v331 = v77 - v18
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
	v203 = F_strlen(m, v18)
	mBase = m.M
	v204 = F_str_toupper(m, v18, v203, l4)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	v140 = F_strlen(m, v134)
	mBase = m.M
	if v140 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	if v184 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L45:
	;
	v184 = int32(0)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v146 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v147 = v18
	v148 = v134
	v149 = v140
	v150 = v146
	goto L52
L49:
	;
	v172 = v134
	v176 = int32(0)
	goto L50
L50:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	v184 = v176 - v177
	goto L44
L51:
	;
	v172 = v167
	v176 = v169
	goto L50
L52:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v150 != v152 {
		v167 = v148
		v169 = v150
		goto L51
	} else {
		goto L54
	}
L53:
	;
	v167 = v161
	v169 = int32(0)
	goto L51
L54:
	;
	if v152 == int32(0) {
		v167 = v148
		v169 = v150
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v157 = v149 - int32(1)
	if v157 == int32(0) {
		v167 = v148
		v169 = v150
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v160 = int32(1)
	v161 = v148 + v160
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	if v162 != 0 {
		v147 = v147 + v160
		v148 = v161
		v149 = v157
		v150 = v162
		goto L52
	} else {
		goto L57
	}
L57:
	;
	goto L53
L58:
	;
	v322 = l3
	v328 = v135
	v331 = v140
	goto L4
L59:
	;
	goto L60
L60:
	;
	v188 = v135 + int32(4)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v189 != 0 {
		v134 = v189
		v135 = v188
		goto L42
	} else {
		goto L61
	}
L61:
	;
	goto L43
L62:
	;
	return int32(0)
L63:
	;
	v208 = F_strlen(m, v204)
	mBase = m.M
	v209 = F_str_tolower(m, v204, v208, l4)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	F_pfree(m, v204)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v213 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v221 = v213
	v222 = l3
	goto L69
L67:
	;
	goto L68
L68:
	;
	F_pfree(m, v209)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L62
	} else {
		goto L94
	}
L69:
	;
	v227 = F_strlen(m, v221)
	mBase = m.M
	v228 = F_str_toupper(m, v221, v227, l4)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L62
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	v230 = F_strlen(m, v228)
	mBase = m.M
	v231 = F_str_tolower(m, v228, v230, l4)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L62
	} else {
		goto L72
	}
L72:
	;
	F_pfree(m, v228)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L62
	} else {
		goto L73
	}
L73:
	;
	v235 = F_strlen(m, v231)
	mBase = m.M
	if v235 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	F_pfree(m, v231)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L62
	} else {
		goto L88
	}
L75:
	;
	v279 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v241 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v242 = v209
	v243 = v231
	v244 = v235
	v245 = v241
	goto L82
L79:
	;
	v267 = v231
	v271 = int32(0)
	goto L80
L80:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	v279 = v271 - v272
	goto L74
L81:
	;
	v267 = v262
	v271 = v264
	goto L80
L82:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v245 != v247 {
		v262 = v243
		v264 = v245
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v262 = v256
	v264 = int32(0)
	goto L81
L84:
	;
	if v247 == int32(0) {
		v262 = v243
		v264 = v245
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v252 = v244 - int32(1)
	if v252 == int32(0) {
		v262 = v243
		v264 = v245
		goto L81
	} else {
		goto L86
	}
L86:
	;
	v255 = int32(1)
	v256 = v243 + v255
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	if v257 != 0 {
		v242 = v242 + v255
		v243 = v256
		v244 = v252
		v245 = v257
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	if v279 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	F_pfree(m, v209)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L62
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v287 = v222 + int32(4)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	if v288 != 0 {
		v221 = v288
		v222 = v287
		goto L69
	} else {
		goto L93
	}
L92:
	;
	v322 = l3
	v328 = v222
	v331 = v235
	goto L4
L93:
	;
	goto L70
L94:
	;
	goto L6
L95:
	;
	goto L3
L96:
	;
	v399 = int32(0)
	v400 = F_errsave_start(m, l6)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L62
	} else {
		goto L106
	}
L97:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	if v355 == int32(0) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v365 = v355
	v366 = v353
	goto L99
L99:
	;
	v371 = base.I32_extend8_s(v365)
	goto L101
L100:
	;
	goto L96
L101:
	;
	if base.B2i32(v371 == int32(32))|base.B2i32(base.Ui32((v371-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v381 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v381)
	goto L96
L103:
	;
	goto L104
L104:
	;
	v384 = v366 + int32(1)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	if v385 != 0 {
		v365 = v385
		v366 = v384
		goto L99
	} else {
		goto L105
	}
L105:
	;
	goto L100
L106:
	;
	if v400 == int32(0) {
		v434 = v399
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L62
	} else {
		goto L108
	}
L108:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v353
	F_errmsg(m, int32(700364), v16)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L62
	} else {
		goto L109
	}
L109:
	;
	F_errdetail(m, int32(646039), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L62
	} else {
		goto L110
	}
L110:
	;
	F_errsave_finish(m, l6, int32(498130), int32(2493), int32(325867))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L62
	} else {
		goto L111
	}
L111:
	;
	v434 = v399
	goto L1
}
