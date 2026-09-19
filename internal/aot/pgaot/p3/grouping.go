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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
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
		v34 = int32(1)
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v34
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
	v21 = v2
	goto L9
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17+v21<<(uint(int32(2))%32))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v26 != int32(1) {
		v34 = v26
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v34 = v26
	goto L4
L11:
	;
	v30 = v21 + int32(1)
	if v30 != v16 {
		v21 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_makeGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13965(m, l0, l1, l2, int32(107))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	var v157 int32
	_ = v157
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
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	F_ExplainOpenGroup(m, int32(_a_F_show_grouping_set_keys_0), v8, int32(1), l6)
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
	v36 = int32(_a_F_show_grouping_set_keys_1)
	goto L5
L4:
	;
	v36 = int32(_a_F_show_grouping_set_keys_2)
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
	F_show_sort_group_keys(m, l0, int32(_a_F_show_grouping_set_keys_3), v40, int32(0), v42, v43, v44, v45, l5, l6)
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
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L98
	}
L12:
	;
	F_ExplainCloseGroup(m, v36, int32(0), l6)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L93
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v58 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v24&int32(-2) == int32(2) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = int32(_a_F_show_grouping_set_keys_4)
	goto L17
L16:
	;
	v63 = int32(_a_F_show_grouping_set_keys_5)
	goto L17
L17:
	;
	v75 = v8
	goto L18
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v75<<(uint(int32(2))%32))))
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L12
L20:
	;
	v432 = v75 + int32(1)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v432 < v433 {
		v75 = v432
		goto L18
	} else {
		goto L92
	}
L21:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	switch v220 {
	case 0, 1:
		goto L51
	case 2:
		goto L53
	case 3:
		goto L52
	default:
		goto L50
	}
L22:
	;
	v85 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v85 < v87 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v200 != 0 {
		v204 = int32(0)
		goto L21
	} else {
		goto L48
	}
L25:
	;
	v90 = v85
	v91 = v85
	goto L28
L26:
	;
	v167 = v85
	goto L27
L27:
	;
	if v167 != 0 {
		v204 = v167
		goto L21
	} else {
		goto L47
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v91<<(uint(int32(2))%32))))
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21+v111<<(uint(int32(1))%32)))))
	if v106 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v167 = v161
	goto L27
L30:
	;
	if v153 == int32(0) {
		goto L11
	} else {
		goto L43
	}
L31:
	;
	goto L30
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v119 <= int32(0) {
		v153 = int32(0)
		goto L31
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v153 = int32(0)
	goto L31
L35:
	;
	v122 = int32(0)
	if v122 < v119 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v125 = v119
	goto L38
L37:
	;
	v125 = v122
	goto L38
L38:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v106)+12))
	v130 = int32(0)
	goto L39
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v126+v130<<(uint(int32(2))%32))))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+8)))
	if v139 == v115&int32(_a_F_show_grouping_set_keys_6) {
		v153 = v138
		goto L31
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v142 = v130 + int32(1)
	if v142 != v125 {
		v130 = v142
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v159 = F_deparse_expression(m, v157, l3, l4, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v161 = F_lappend(m, v90, v159)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v164 = v91 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if v164 < v165 {
		v90 = v161
		v91 = v164
		goto L28
	} else {
		goto L46
	}
L46:
	;
	goto L29
L47:
	;
	goto L24
L48:
	;
	F_ExplainPropertyText(m, v63, int32(_a_F_show_grouping_set_keys_7), l6)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L20
L50:
	;
	goto L20
L51:
	;
	F_ExplainPropertyList(m, v63, v204, l6)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L91
	}
L52:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	if v312 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l6)+28))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	if v223 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v230, int32(10))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L59
	}
L55:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v224, int32(44))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = int32(1)
	goto L54
L58:
	;
	goto L54
L59:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v234, v235<<(uint(int32(1))%32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v240, int32(91))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v204 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v306, int32(93))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L72
	}
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v246 <= int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	F_escape_json(m, v249, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v254 = int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v255 <= v254 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v259 = v254
	goto L67
L67:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v275, int32(_a_F_show_grouping_set_keys_8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v274+v259<<(uint(int32(2))%32))))
	F_escape_json(m, v279, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v287 = v259 + int32(1)
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v287 < v288 {
		v259 = v287
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	goto L50
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v327, int32(_a_F_show_grouping_set_keys_9))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L79
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = int32(1)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v317, int32(10))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	F_appendStringInfoSpaces(m, v321, v322<<(uint(int32(1))%32))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	if v204 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoChar(m, v393, int32(93))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L90
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v334 <= int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	F_escape_json(m, v337, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v342 <= int32(1) {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v346 = int32(1)
	goto L85
L85:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	F_appendStringInfoString(m, v362, int32(_a_F_show_grouping_set_keys_8))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L80
L87:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v361+v346<<(uint(int32(2))%32))))
	F_escape_json(m, v366, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v374 = v346 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v374 < v375 {
		v346 = v374
		goto L85
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	goto L50
L91:
	;
	goto L50
L92:
	;
	goto L19
L93:
	;
	if l2 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_ExplainCloseGroup(m, int32(_a_F_show_grouping_set_keys_0), int32(1), l6)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L97
	}
L95:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	if v456 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l6)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v457 - int32(1)
	goto L94
L97:
	;
	m.G0 = v19 + int32(16)
	return
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v115
	F_errmsg_internal(m, int32(_a_F_show_grouping_set_keys_10), v19)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_show_grouping_set_keys_11), int32(2722), int32(_a_F_show_grouping_set_keys_12))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
