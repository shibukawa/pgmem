package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_map_sql_value_to_xml_value(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int64
	_ = v387
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	v6 = m.G0
	v8 = v6 - int32(208)
	m.G0 = v8
	v10 = F_get_base_element_type(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(208)
	return v497
L2:
	;
	return int32(0)
L3:
	;
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v14 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v92 = F_getBaseType(m, l1)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L33
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_get_typlenbyvalalign(m, v16, v8+int32(12), v8+int32(207), v8+int32(206))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+12)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+207)))
	v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+206)))
	F_deconstruct_array(m, v14, v25, v26, v27, v8+int32(200), v8+int32(196), v8+int32(152))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_initStringInfo(m, v8+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+152))
	if int32(0) < v40 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = v40
	v45 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	F_pfree(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L24
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+v45))))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_appendStringInfoString(m, v8+int32(16), int32(546223))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v76 = v44
	goto L18
L18:
	;
	v78 = v45 + int32(1)
	if v78 < v76 {
		v44 = v76
		v45 = v78
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v45<<(uint(int32(2))%32))))
	v66 = F_map_sql_value_to_xml_value(m, v65, v16)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoString(m, v8+int32(16), v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v8+int32(16), int32(546233))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v8)+152))
	v76 = v75
	goto L18
L23:
	;
	goto L15
L24:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	F_pfree(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v497 = v91
	goto L1
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L120
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L115
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L2
	} else {
		goto L110
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v387 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v387-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L26
	} else {
		goto L99
	}
L30:
	;
	v338 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v338-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L27
	} else {
		goto L88
	}
L31:
	;
	if l0 != 0 {
		goto L85
	} else {
		goto L86
	}
L32:
	;
	F_getTypeOutputInfo(m, v92, v8+int32(16), v8+int32(152))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L81
	}
L33:
	;
	if v92 <= int32(1113) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v92 == int32(16) {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v92 == int32(1114) {
		goto L30
	} else {
		goto L79
	}
L37:
	;
	if v92 != int32(1082) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(l0-int32(2147483647)) <= base.Ui32(int32(1)) {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v115 = l0 + int32(2483589)
	v116 = int32(146097)
	v117 = base.I32_div_u_s(v115, v116)
	v118 = int32(3)
	v124 = int32(2)
	v129 = base.I32_div_u_s((v117*int32(1073595727)+v115)<<(uint(v124)%32)|v118, v116)
	v132 = l0 + int32(2451545) + v117*v118 + v129 + int32(32104)
	v133 = int32(1461)
	v134 = base.I32_div_u_s(v132, v133)
	v137 = v134*int32(-1461) + v132
	v139 = v137 << (uint(v124) % 32)
	if base.Ui32(v133) <= base.Ui32(v139) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v179 = v8 + int32(152)
	v182 = v8 + int32(16)
	switch int32(3) {
	case 0, 3:
		goto L50
	case 1:
		goto L49
	case 2:
		goto L48
	default:
		goto L47
	}
L41:
	;
	v152 = base.I32_div_u_s(v139, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(172)))) = v152 + v134<<(uint(int32(2))%32) - int32(4800)
	v160 = v150 + int32(123)
	v164 = int32(base.Ui32(v160*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(164)))) = v160 - int32(base.Ui32(v164*int32(7834))>>(uint(int32(8))%32))
	v174 = base.I32_rem_u_s(v164+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(168)))) = v174 + int32(1)
	goto L40
L42:
	;
	v145 = base.I32_rem_u_s(v137+int32(305), int32(365))
	v150 = v145
	goto L41
L43:
	;
	goto L44
L44:
	;
	v149 = base.I32_rem_u_s(v137+int32(306), int32(366))
	v150 = v149
	goto L41
L45:
	;
	v316 = F_pstrdup(m, v8+int32(16))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L78
	}
L46:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if v300 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L47:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _consts[1062]))
	v268 = base.B2i32(v266 == int32(1))
	if v266 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L48:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v242 = int32(2)
	v243 = F_pg_ultostr_zeropad(m, v182, v241, v242)
	mBase = m.M
	v244 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v243))) = uint8(v244)
	v246 = int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v250 = F_pg_ultostr_zeropad(m, v243+v246, v248, v242)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v244)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if int32(0) < v255 {
		goto L63
	} else {
		goto L64
	}
L49:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[1062]))
	v212 = base.B2i32(v210 == int32(1))
	if v210 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if int32(0) < v185 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v190 = v185
	goto L53
L52:
	;
	v190 = int32(1) - v185
	goto L53
L53:
	;
	v192 = F_pg_ultostr_zeropad(m, v182, v190, int32(4))
	mBase = m.M
	v193 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v193)
	v195 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v198 = int32(2)
	v199 = F_pg_ultostr_zeropad(m, v192+v195, v197, v198)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v193)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v206 = F_pg_ultostr_zeropad(m, v199+v195, v204, v198)
	mBase = m.M
	v299 = v206
	goto L46
L54:
	;
	v213 = int32(12)
	goto L56
L55:
	;
	v213 = int32(16)
	goto L56
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v179+v213)))
	v217 = F_pg_ultostr_zeropad(m, v182, v215, int32(2))
	mBase = m.M
	v218 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v218)
	if v210 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v224 = int32(16)
	goto L59
L58:
	;
	v224 = int32(12)
	goto L59
L59:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v179+v224)))
	v228 = F_pg_ultostr_zeropad(m, v217+int32(1), v226, int32(2))
	mBase = m.M
	v229 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v229)
	v231 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if int32(0) < v233 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = v233
	goto L62
L61:
	;
	v238 = v231 - v233
	goto L62
L62:
	;
	v240 = F_pg_ultostr_zeropad(m, v228+v231, v238, int32(4))
	mBase = m.M
	v299 = v240
	goto L46
L63:
	;
	v260 = v255
	goto L65
L64:
	;
	v260 = v246 - v255
	goto L65
L65:
	;
	v262 = F_pg_ultostr_zeropad(m, v250+v246, v260, int32(4))
	mBase = m.M
	v299 = v262
	goto L46
L66:
	;
	v269 = int32(12)
	goto L68
L67:
	;
	v269 = int32(16)
	goto L68
L68:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v179+v269)))
	v273 = F_pg_ultostr_zeropad(m, v182, v271, int32(2))
	mBase = m.M
	v274 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v273))) = uint8(v274)
	if v266 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v280 = int32(16)
	goto L71
L70:
	;
	v280 = int32(12)
	goto L71
L71:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v179+v280)))
	v284 = F_pg_ultostr_zeropad(m, v273+int32(1), v282, int32(2))
	mBase = m.M
	v285 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v284))) = uint8(v285)
	v287 = int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	if int32(0) < v289 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v294 = v289
	goto L74
L73:
	;
	v294 = v287 - v289
	goto L74
L74:
	;
	v296 = F_pg_ultostr_zeropad(m, v284+v287, v294, int32(4))
	mBase = m.M
	v299 = v296
	goto L46
L75:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1063])))
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+2)) = uint8(v304)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1064])))
	*(*uint16)(unsafe.Add(mBase, uint32(v299))) = uint16(v307)
	v311 = v299 + int32(3)
	goto L77
L76:
	;
	v311 = v299
	goto L77
L77:
	;
	v312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v311))) = uint8(v312)
	goto L45
L78:
	;
	v497 = v316
	goto L1
L79:
	;
	if v92 == int32(1184) {
		goto L29
	} else {
		goto L80
	}
L80:
	;
	goto L32
L81:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v329 = F_OidOutputFunctionCall(m, v328, l0)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	if v92 == int32(142) {
		v497 = v329
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v333 = F_escape_xml(m, v329)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v497 = v333
	goto L1
L85:
	;
	v337 = int32(344470)
	goto L87
L86:
	;
	v337 = int32(361486)
	goto L87
L87:
	;
	v497 = v337
	goto L1
L88:
	;
	v343 = int32(0)
	v350 = F_timestamp2tm(m, v338, v343, v8+int32(152), v8+int32(200), v343, v343)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v350 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v357 = int32(0)
	F_EncodeDateTime(m, v8+int32(152), v356, v357, v357, v357, int32(4), v8+int32(16))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L95
	}
L93:
	;
	v367 = F_pstrdup(m, v8+int32(16))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v497 = v367
	goto L1
L95:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(402548), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(497496), int32(2582), int32(345254))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
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
	v401 = F_timestamp2tm(m, v387, v8+int32(200), v8+int32(152), v8+int32(196), v8+int32(12), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v401 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_EncodeDateTime(m, v8+int32(152), v407, int32(1), v409, v410, int32(4), v8+int32(16))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L2
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L106
	}
L104:
	;
	v418 = F_pstrdup(m, v8+int32(16))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v497 = v418
	goto L1
L106:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(402548), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(497496), int32(2609), int32(345254))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(402633), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	F_errdetail(m, int32(596233), int32(0))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(497496), int32(2554), int32(345254))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(402548), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	F_errdetail(m, int32(596101), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(497496), int32(2576), int32(345254))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(402548), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_errdetail(m, int32(596101), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(497496), int32(2603), int32(345254))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sql_compile_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v8 != 0 {
		v9 = F_geterrposition(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			if v9 <= int32(0) {
				F_set_errcontext_domain(m, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27
					F_errcontext_msg(m, int32(232132), v6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				}
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v13 == int32(0) {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27
						F_errcontext_msg(m, int32(232132), v6)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					}
				} else {
					v17 = F_errposition(m, int32(0))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_internalerrposition(m, v9)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v22 = F_internalerrquery(m, v21)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return
							} else {
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v26 = m.ExcPending
								if v26 != 0 {
									return
								} else {
									v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v27
									F_errcontext_msg(m, int32(232132), v6)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_sql_fn_post_column_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 != 0 {
		v144 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v144
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = int32(2)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(v29)%32)-int32(4))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v38 = v25 - base.B2i32(v35 == int32(77))
	if v29 <= v38 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, 4))
	v24 = v20
	v25 = v4
	goto L3
L5:
	;
	goto L6
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(3) < v21 {
		v144 = v4
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v24 = v21
	v25 = v21
	goto L3
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v43 = v41
	v44 = v42
	goto L10
L9:
	;
	v43 = int32(0)
	v44 = v4
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	switch v38 - int32(2) {
	case 0:
		goto L13
	case 1:
		goto L14
	default:
		goto L12
	}
L11:
	;
	if v119 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L12:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v116 = F_sql_fn_resolve_param_name(m, v15, v45, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L26
	} else {
		goto L40
	}
L13:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v87 == int32(0) {
		v106 = v86
		v107 = v87
		goto L29
	} else {
		goto L30
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v52 == int32(0) {
		v71 = v51
		v72 = v52
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v72-v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	goto L15
L17:
	;
	if v51 != v52 {
		v71 = v51
		v72 = v52
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v56 = v45
	v57 = v48
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v61 == int32(0) {
		v71 = v60
		v72 = v61
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v71 = v60
	v72 = v61
	goto L16
L21:
	;
	v64 = int32(1)
	if v60 == v61 {
		v56 = v56 + v64
		v57 = v57 + v64
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v144 = int32(0)
	goto L1
L24:
	;
	goto L25
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v76 = F_sql_fn_resolve_param_name(m, v15, v44, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int32(0)
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v118 = v82
	v119 = v76
	goto L11
L28:
	;
	if v107-v106 != 0 {
		goto L12
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	if v86 != v87 {
		v106 = v86
		v107 = v87
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v91 = v45
	v92 = v83
	goto L32
L32:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v95
		v107 = v96
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v106 = v95
	v107 = v96
	goto L29
L34:
	;
	v99 = int32(1)
	if v95 == v96 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v110 = F_sql_fn_resolve_param_name(m, v15, v44, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	if v110 != 0 {
		v144 = v110
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v113 = F_sql_fn_resolve_param_name(m, v15, v45, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	v118 = v43
	v119 = v113
	goto L11
L40:
	;
	v118 = v43
	v119 = v116
	goto L11
L41:
	;
	v144 = int32(0)
	goto L1
L42:
	;
	goto L43
L43:
	;
	if v118 == int32(0) {
		v144 = v119
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v118
	v130 = F_list_make1_impl(m, int32(1), v13+int32(4))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v119
	v135 = F_list_make1_impl(m, int32(1), v13)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v138 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v141 = F_ParseFuncOrColumn(m, l0, v130, v135, v137, v138, v138, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L26
	} else {
		goto L47
	}
L47:
	;
	v144 = v141
	goto L1
}
