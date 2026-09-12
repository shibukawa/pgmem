package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_grouping_is_hashable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= int32(0) {
		v35 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v35
L5:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v16 = v10
	goto L8
L7:
	;
	v16 = v13
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = v2
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17+v20<<(uint(int32(2))%32))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v26 != int32(1) {
		v35 = v26
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v35 = v26
	goto L4
L11:
	;
	v30 = v20 + int32(1)
	if v30 != v16 {
		v20 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(107)
		return v6
	}
}
func F_show_grouping_set_keys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
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
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainOpenGroup(m, int32(109100), v8, int32(1), l6)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = base.B2i32(v24&int32(-2) == int32(2))
	if v24&int32(-2) == int32(2) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v36 = int32(113510)
	goto L5
L4:
	;
	v36 = int32(113499)
	goto L5
L5:
	;
	if l2 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_ExplainOpenGroup(m, v36, v36, int32(0), l6)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	F_show_sort_group_keys(m, l0, int32(22737), v40, int32(0), v42, v43, v44, v45, l5, l6)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v48 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v49 + int32(1)
	goto L6
L10:
	;
	if v22 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_ExplainCloseGroup(m, v36, int32(0), l6)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L96
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v58 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if v24&int32(-2) == int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = int32(22778)
	goto L16
L15:
	;
	v63 = int32(22768)
	goto L16
L16:
	;
	v75 = v8
	goto L17
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v75<<(uint(int32(2))%32))))
	if v84 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L11
L19:
	;
	v443 = v75 + int32(1)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v443 < v444 {
		v75 = v443
		goto L17
	} else {
		goto L95
	}
L20:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	switch v231 {
	case 0, 1:
		goto L54
	case 2:
		goto L56
	case 3:
		goto L55
	default:
		goto L53
	}
L21:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v87 <= v85 {
		v178 = v85
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v211 != 0 {
		v215 = int32(0)
		goto L20
	} else {
		goto L51
	}
L24:
	;
	if v178 != 0 {
		v215 = v178
		goto L20
	} else {
		goto L50
	}
L25:
	;
	v90 = v85
	v91 = v85
	goto L26
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v91<<(uint(int32(2))%32))))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v111<<(uint(int32(1))%32)))))
	if v106 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L28:
	;
	if v153 != 0 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	goto L28
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v119 <= int32(0) {
		v153 = int32(0)
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v153 = int32(0)
	goto L29
L33:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v125 = v119
	goto L36
L35:
	;
	v125 = v122
	goto L36
L36:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v130 = int32(0)
	goto L37
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126+v130<<(uint(int32(2))%32))))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+8)))
	if v139 == v115&int32(65535) {
		v153 = v138
		goto L29
	} else {
		goto L39
	}
L38:
	;
	goto L32
L39:
	;
	v142 = v130 + int32(1)
	if v142 != v125 {
		v130 = v142
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v157 = F_deparse_expression(m, v155, l3, l4, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L27
L44:
	;
	v159 = F_lappend(m, v90, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v162 = v91 + int32(1)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v162 < v163 {
		v90 = v159
		v91 = v162
		goto L26
	} else {
		goto L46
	}
L46:
	;
	v178 = v159
	goto L24
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v115
	F_errmsg_internal(m, int32(468733), v19)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(498126), int32(2722), int32(113262))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
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
	goto L23
L51:
	;
	F_ExplainPropertyText(m, v63, int32(686651), l6)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L19
L53:
	;
	goto L19
L54:
	;
	F_ExplainPropertyList(m, v63, v215, l6)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L94
	}
L55:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	if v323 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v234 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v241, int32(10))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v235, int32(44))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = int32(1)
	goto L57
L61:
	;
	goto L57
L62:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v245, v246<<(uint(int32(1))%32))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v251, int32(91))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v215 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v317, int32(93))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L75
	}
L66:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v257 <= int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	F_escape_json(m, v260, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v265 = int32(1)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v266 <= v265 {
		goto L65
	} else {
		goto L69
	}
L69:
	;
	v270 = v265
	goto L70
L70:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v286, int32(748219))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L72
	}
L71:
	;
	goto L65
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v285+v270<<(uint(int32(2))%32))))
	F_escape_json(m, v290, v294)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v298 = v270 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v298 < v299 {
		v270 = v298
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	goto L53
L76:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v338, int32(509456))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L82
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = int32(1)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v328, int32(10))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v332, v333<<(uint(int32(1))%32))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L76
L82:
	;
	if v215 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v404, int32(93))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L93
	}
L84:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v345 <= int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	F_escape_json(m, v348, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v353 <= int32(1) {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v357 = int32(1)
	goto L88
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v373, int32(748219))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	goto L83
L90:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v372+v357<<(uint(int32(2))%32))))
	F_escape_json(m, v377, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v385 = v357 + int32(1)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v385 < v386 {
		v357 = v385
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	goto L53
L94:
	;
	goto L53
L95:
	;
	goto L18
L96:
	;
	if l2 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_ExplainCloseGroup(m, int32(109100), int32(1), l6)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v467 != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v468 - int32(1)
	goto L97
L100:
	;
	m.G0 = v19 + int32(16)
	return
}
