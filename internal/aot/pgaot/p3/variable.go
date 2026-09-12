package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_expandRecordVariable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
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
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v15 = l1
	v19 = int32(0)
	goto L4
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L23
	} else {
		goto L119
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L23
	} else {
		goto L116
	}
L3:
	;
	m.G0 = v12 + int32(160)
	return v469
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v25 = v24 + v19
	v26 = int32(0)
	if v25 <= v26 {
		v75 = l0
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v463 = F_get_expr_result_tupdesc(m, v459, int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L23
	} else {
		goto L115
	}
L6:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+8)))
	if v89 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L7:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82+v23<<(uint(int32(2))%32)-int32(4))))
	goto L6
L8:
	;
	v32 = v25 & int32(7)
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v25) < base.Ui32(int32(8)) {
		v75 = v47
		goto L7
	} else {
		goto L16
	}
L10:
	;
	v47 = l0
	v50 = v25
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = l0
	v38 = v25
	v39 = v26
	goto L13
L13:
	;
	v41 = int32(1)
	v42 = v38 - v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v45 = v39 + v41
	if v45 != v32 {
		v35 = v43
		v38 = v42
		v39 = v45
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v47 = v43
	v50 = v42
	goto L9
L15:
	;
	goto L14
L16:
	;
	v55 = v47
	v58 = v50
	goto L17
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if base.Ui32(v58-int32(9)) < base.Ui32(int32(-2)) {
		v55 = v70
		v58 = v58 - int32(8)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v75 = v70
	goto L7
L19:
	;
	goto L18
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v93 = int32(0)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	F_expandRTE(m, v88, v92, v93, v94, v95, v93, v12+int32(32), v12+int32(156))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v171 != int32(2) {
		goto L44
	} else {
		goto L45
	}
L23:
	;
	return int32(0)
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if v106 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v109 = v107
	goto L27
L26:
	;
	v109 = int32(0)
	goto L27
L27:
	;
	v110 = F_CreateTemplateTupleDesc(m, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v116 = int32(0)
	v120 = int32(1)
	goto L29
L29:
	;
	v124 = int32(0)
	if v113 == v124 {
		v133 = v124
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v112 == int32(0) {
		v469 = v110
		goto L3
	} else {
		goto L34
	}
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v127 <= v116 {
		v133 = v124
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v133 = v129 + v116<<(uint(int32(2))%32)
	goto L31
L34:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v136 <= v116 {
		v469 = v110
		goto L3
	} else {
		goto L35
	}
L35:
	;
	if v133 == int32(0) {
		v469 = v110
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v143 = v140 + v116<<(uint(int32(2))%32)
	if v143 == int32(0) {
		v469 = v110
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v146 = base.I32_extend16_s(v120)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v150 = F_exprType(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L23
	} else {
		goto L38
	}
L38:
	;
	v152 = F_exprTypmod(m, v149)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	F_TupleDescInitEntry(m, v110, v146, v148, v150, v152, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v157 = F_exprCollation(m, v149)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L23
	} else {
		goto L41
	}
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v110+v159<<(uint(int32(4))%32)+v146*int32(100))+16)) = v157
	goto L42
L42:
	;
	v167 = int32(1)
	v116 = v116 + v167
	v120 = v120 + v167
	goto L29
L43:
	;
	goto L5
L44:
	;
	switch v171 - int32(1) {
	case 0:
		goto L48
	default:
		v459 = v15
		goto L43
	case 5:
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v88)+52))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v449+v89<<(uint(int32(2))%32)-int32(4))))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	if v456 == int32(6) {
		v15 = v455
		v19 = v25
		goto L4
	} else {
		goto L114
	}
L47:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+92)))
	if v305 != 0 {
		v459 = v15
		goto L43
	} else {
		goto L79
	}
L48:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+76))
	if v177 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	if v215 == int32(0) {
		goto L2
	} else {
		goto L62
	}
L50:
	;
	goto L49
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	if v181 <= int32(0) {
		v215 = int32(0)
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v215 = int32(0)
	goto L50
L54:
	;
	v184 = int32(0)
	if v184 < v181 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v187 = v181
	goto L57
L56:
	;
	v187 = v184
	goto L57
L57:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v192 = int32(0)
	goto L58
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v188+v192<<(uint(int32(2))%32))))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+8)))
	if v201 == v89&int32(65535) {
		v215 = v200
		goto L50
	} else {
		goto L60
	}
L59:
	;
	goto L53
L60:
	;
	v204 = v192 + int32(1)
	if v204 != v187 {
		v192 = v204
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+26)))
	if v219 == int32(1) {
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v223 != int32(6) {
		v459 = v222
		goto L43
	} else {
		goto L64
	}
L64:
	;
	v231 = F__emscripten_memset_bulkmem(m, v12+int32(36), base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	goto L65
L65:
	;
	if v25 == int32(0) {
		v288 = l0
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v288
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v88)+36))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v299
	v303 = F_expandRecordVariable(m, v12+int32(32), v222)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L23
	} else {
		goto L78
	}
L67:
	;
	v234 = int32(7)
	v235 = v25 & v234
	if base.Ui32(v234) <= base.Ui32(v25-int32(1)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v243 = l0
	v248 = int32(0)
	goto L71
L69:
	;
	v263 = l0
	goto L70
L70:
	;
	if v235 == int32(0) {
		v288 = v263
		goto L66
	} else {
		goto L74
	}
L71:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v261 = v248 + int32(8)
	if v261 != v25&int32(-8) {
		v243 = v259
		v248 = v261
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v263 = v259
	goto L70
L73:
	;
	goto L72
L74:
	;
	v275 = v263
	v280 = int32(0)
	goto L75
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	v286 = v280 + int32(1)
	if v286 != v235 {
		v275 = v284
		v280 = v286
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v288 = v284
	goto L66
L77:
	;
	goto L76
L78:
	;
	v469 = v303
	goto L3
L79:
	;
	v306 = F_GetCTEForRTE(m, l0, v88, v25)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L23
	} else {
		goto L80
	}
L80:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v306)+16))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v311 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v314 = int32(76)
	goto L83
L82:
	;
	v314 = int32(96)
	goto L83
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v308+v314)))
	if v316 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	if v354 == int32(0) {
		goto L1
	} else {
		goto L97
	}
L85:
	;
	goto L84
L86:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v320 <= int32(0) {
		v354 = int32(0)
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v354 = int32(0)
	goto L85
L89:
	;
	v323 = int32(0)
	if v323 < v320 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v326 = v320
	goto L92
L91:
	;
	v326 = v323
	goto L92
L92:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v331 = int32(0)
	goto L93
L93:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v327+v331<<(uint(int32(2))%32))))
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339)+8)))
	if v340 == v89&int32(65535) {
		v354 = v339
		goto L85
	} else {
		goto L95
	}
L94:
	;
	goto L88
L95:
	;
	v343 = v331 + int32(1)
	if v343 != v326 {
		v331 = v343
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+26)))
	if v358 == int32(1) {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v362 != int32(6) {
		v459 = v361
		goto L43
	} else {
		goto L99
	}
L99:
	;
	v370 = F__emscripten_memset_bulkmem(m, v12+int32(36), base.I32_extend8_s(int32(0)), int32(120))
	mBase = m.M
	goto L100
L100:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v88)+88))
	v372 = v371 + v25
	if v372 == int32(0) {
		v431 = l0
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v431
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v306)+16))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v442
	v446 = F_expandRecordVariable(m, v12+int32(32), v361)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L23
	} else {
		goto L113
	}
L102:
	;
	v375 = int32(7)
	v376 = v372 & v375
	if base.Ui32(v375) <= base.Ui32(v19+v371+v24-int32(1)) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v386 = l0
	v388 = int32(0)
	goto L106
L104:
	;
	v406 = l0
	goto L105
L105:
	;
	if v376 == int32(0) {
		v431 = v406
		goto L101
	} else {
		goto L109
	}
L106:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v404 = v388 + int32(8)
	if v404 != v372&int32(-8) {
		v386 = v402
		v388 = v404
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v406 = v402
	goto L105
L108:
	;
	goto L107
L109:
	;
	v418 = v406
	v420 = int32(0)
	goto L110
L110:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	v429 = v420 + int32(1)
	if v429 != v376 {
		v418 = v427
		v420 = v429
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v431 = v427
	goto L101
L112:
	;
	goto L111
L113:
	;
	v469 = v446
	goto L3
L114:
	;
	v459 = v455
	goto L43
L115:
	;
	v469 = v463
	goto L3
L116:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v483
	F_errmsg_internal(m, int32(493484), v12+int32(16))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L23
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(512790), int32(1600), int32(412824))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L23
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v88)+8))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v501
	F_errmsg_internal(m, int32(493523), v12)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L23
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(512790), int32(1660), int32(412824))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L23
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_variable_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 float64
	_ = v195
	var v198 int32
	_ = v198
	var v199 float32
	_ = v199
	var v202 float32
	_ = v202
	var v205 float32
	_ = v205
	var v208 float32
	_ = v208
	var v210 float64
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v229 float64
	_ = v229
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 float64
	_ = v245
	var v249 float32
	_ = v249
	var v251 float64
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v270 float64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 float32
	_ = v276
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	v6 = int32(0)
	v14 = float64(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)) = uint8(v6)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v25 == v6 {
		v355 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(96)
	return v355 & int32(1)
L2:
	;
	v28 = F_get_opcode(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_get_typlenbyval(m, v56, v17+int32(84), v17+int32(83))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L3
	} else {
		goto L15
	}
L6:
	;
	if v28 == int32(0) {
		v355 = v6
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v35 = F_get_func_leakproof(m, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if v35 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v39 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v39 == int32(0) {
		v355 = v6
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = F_get_func_name(m, v28)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v43
	F_errmsg_internal(m, int32(352530), v17)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(514641), int32(6242), int32(330924))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v355 = v6
	goto L1
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v68 = F_get_attstatsslot(m, v17+int32(16), v65, int32(2), l1, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L22
	}
L16:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v17)+92))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v345
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v347
	v355 = v337
	goto L1
L17:
	;
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L62
	}
L18:
	;
	v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	F_get_stats_slot_range(m, v17+int32(16), v28, v17+int32(52), l2, v301, v302, v17+int32(92), v17+int32(88), v17+int32(87))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L3
	} else {
		goto L61
	}
L19:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	if v170 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	v157 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v164 = F_get_attstatsslot(m, v17+int32(16), v160, int32(1), v157, int32(3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L43
	}
L21:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v151 = F_get_attstatsslot(m, v17+int32(16), v148, int32(1), int32(0), v145)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L40
	}
L22:
	;
	if v68 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v70 != l2 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v115 = F_get_attstatsslot(m, v17+int32(16), v111, int32(2), int32(0), int32(1))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L33
	}
L26:
	;
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L32
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v73 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v80 = F_datumCopy(m, v77, v78, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v83+v84<<(uint(int32(2))%32)-int32(4))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v93 = F_datumCopy(m, v90, v91, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)) = uint8(v95)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v93
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v144 = int32(1)
	v145 = int32(1)
	goto L21
L32:
	;
	goto L25
L33:
	;
	if v115 == int32(0) {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+84)))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+83)))
	F_get_stats_slot_range(m, v17+int32(16), v28, v17+int32(52), l2, v123, v124, v17+int32(92), v17+int32(88), v17+int32(87))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L35
	}
L35:
	;
	F_free_attstatsslot(m, v17+int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v137 = int32(1)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)))
	if v139&v137 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v142 = v137
	goto L39
L38:
	;
	v142 = int32(3)
	goto L39
L39:
	;
	v144 = v139
	v145 = v142
	goto L21
L40:
	;
	if v151 == int32(0) {
		v337 = v144
		goto L16
	} else {
		goto L41
	}
L41:
	;
	if v144&int32(1) != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	goto L19
L43:
	;
	if v164 == int32(0) {
		v337 = v157
		goto L16
	} else {
		goto L44
	}
L44:
	;
	goto L19
L45:
	;
	v271 = int32(0)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+22)))
	v276 = *(*float32)(unsafe.Add(mBase, uint32(v273+v274)+8))
	if base.F64_gt(base.F64_add(v270, base.F64_promote_f32(v276)), float64(0.99999)) == v271 {
		v326 = v271
		goto L17
	} else {
		goto L60
	}
L46:
	;
	v270 = v14
	goto L45
L47:
	;
	goto L48
L48:
	;
	v174 = v170 & int32(3)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	if base.Ui32(v170) < base.Ui32(int32(4)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v174 == int32(0) {
		v270 = v229
		goto L45
	} else {
		goto L56
	}
L50:
	;
	v222 = int32(0)
	v229 = v14
	goto L49
L51:
	;
	goto L52
L52:
	;
	v188 = int32(0)
	v193 = v6
	v195 = v14
	goto L53
L53:
	;
	v198 = v175 + v188<<(uint(int32(2))%32)
	v199 = *(*float32)(unsafe.Add(mBase, uint32(v198)))
	v202 = *(*float32)(unsafe.Add(mBase, uint32(v198)+4))
	v205 = *(*float32)(unsafe.Add(mBase, uint32(v198)+8))
	v208 = *(*float32)(unsafe.Add(mBase, uint32(v198)+12))
	v210 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v195, base.F64_promote_f32(v199)), base.F64_promote_f32(v202)), base.F64_promote_f32(v205)), base.F64_promote_f32(v208))
	v211 = int32(4)
	v212 = v188 + v211
	v214 = v193 + v211
	if v214 != v170&int32(2147483644) {
		v188 = v212
		v193 = v214
		v195 = v210
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v222 = v212
	v229 = v210
	goto L49
L55:
	;
	goto L54
L56:
	;
	v238 = v222
	v241 = v6
	v245 = v229
	goto L57
L57:
	;
	v249 = *(*float32)(unsafe.Add(mBase, uint32(v175+v238<<(uint(int32(2))%32))))
	v251 = base.F64_add(v245, base.F64_promote_f32(v249))
	v252 = int32(1)
	v255 = v241 + v252
	if v255 != v174 {
		v238 = v238 + v252
		v241 = v255
		v245 = v251
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v270 = v251
	goto L45
L59:
	;
	goto L58
L60:
	;
	goto L18
L61:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+87)))
	v326 = v311
	goto L17
L62:
	;
	v337 = v326
	goto L16
}
func F_map_variable_attnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v6)
	v22 = F_query_or_expression_tree_mutator_impl(m, l0, int32(1055), v9+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(32)
		return v22
	}
}
