package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_sql_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_sql_expr[0])))
	if v12 == int32(1) {
		v15 = int32(_a_F_check_sql_expr_0)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[1]))
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[2]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[1])) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_check_sql_expr_1)
		v25 = int32(_a_F_check_sql_expr_2)
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[3]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[3])) = v9 + int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v9 + int32(24)
		v35 = F_raw_parser(m, l0, l1)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[1])) = v16
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_check_sql_expr[3])) = v40
			m.G0 = v9 + int32(32)
			return
		}
	} else {
		m.G0 = v9 + int32(32)
		return
	}
}
func F_map_sql_value_to_xml_value(m *base.Module, l0 int64, l1 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
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
	return v485
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
	v15 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v89 = F_getBaseType(m, l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L33
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	F_get_typlenbyvalalign(m, v17, v8+int32(12), v8+int32(207), v8+int32(206))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+12)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+207)))
	v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v8)+206)))
	F_deconstruct_array(m, v15, v26, v27, v28, v8+int32(200), v8+int32(196), v8+int32(152))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	F_initStringInfo(m, v8+int32(16))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+152))
	if int32(0) < v41 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = int32(0)
	v48 = v41
	goto L14
L12:
	;
	goto L13
L13:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	F_pfree(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L24
	}
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v46))))
	if v52 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v56 = v8 + int32(16)
	F_appendStringInfoString(m, v56, int32(_a_F_map_sql_value_to_xml_value_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v73 = v48
	goto L18
L18:
	;
	v75 = v46 + int32(1)
	if v75 < v73 {
		v46 = v75
		v48 = v73
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v60+v46<<(uint(int32(3))%32))))
	v65 = F_map_sql_value_to_xml_value(m, v64, v17)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoString(m, v56, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v56, int32(_a_F_map_sql_value_to_xml_value_1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v8)+152))
	v73 = v72
	goto L18
L23:
	;
	goto L15
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	F_pfree(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v485 = v88
	goto L1
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L2
	} else {
		goto L120
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L2
	} else {
		goto L115
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L2
	} else {
		goto L110
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	if base.Ui64(l0-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L26
	} else {
		goto L99
	}
L30:
	;
	if base.Ui64(l0-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L27
	} else {
		goto L88
	}
L31:
	;
	if l0 == int64(0) {
		goto L85
	} else {
		goto L86
	}
L32:
	;
	F_getTypeOutputInfo(m, v89, v8+int32(16), v8+int32(152))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L2
	} else {
		goto L81
	}
L33:
	;
	if v89 <= int32(1113) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v89 == int32(16) {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v89 == int32(1114) {
		goto L30
	} else {
		goto L79
	}
L37:
	;
	if v89 != int32(1082) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v97 = base.I32_wrap_i64(l0)
	if base.Ui32(v97-int32(2147483647)) <= base.Ui32(int32(1)) {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	v113 = v97 + int32(_a_F_map_sql_value_to_xml_value_2)
	v114 = int32(_a_F_map_sql_value_to_xml_value_3)
	v115 = base.I32_div_u_s(v113, v114)
	v116 = int32(3)
	v122 = int32(2)
	v127 = base.I32_div_u_s((v115*int32(1073595727)+v113)<<(uint(v122)%32)|v116, v114)
	v130 = v97 + int32(_a_F_map_sql_value_to_xml_value_4) + v115*v116 + v127 + int32(_a_F_map_sql_value_to_xml_value_5)
	v131 = int32(1461)
	v132 = base.I32_div_u_s(v130, v131)
	v135 = v132*int32(-1461) + v130
	v137 = v135 << (uint(v122) % 32)
	if base.Ui32(v131) <= base.Ui32(v137) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v177 = v8 + int32(152)
	v180 = v8 + int32(16)
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
	v150 = base.I32_div_u_s(v137, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(172)))) = v150 + v132<<(uint(int32(2))%32) - int32(_a_F_map_sql_value_to_xml_value_6)
	v158 = v148 + int32(123)
	v162 = int32(base.Ui32(v158*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(164)))) = v158 - int32(base.Ui32(v162*int32(_a_F_map_sql_value_to_xml_value_7))>>(uint(int32(8))%32))
	v172 = base.I32_rem_u_s(v162+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(168)))) = v172 + int32(1)
	goto L40
L42:
	;
	v143 = base.I32_rem_u_s(v135+int32(305), int32(365))
	v148 = v143
	goto L41
L43:
	;
	goto L44
L44:
	;
	v147 = base.I32_rem_u_s(v135+int32(306), int32(366))
	v148 = v147
	goto L41
L45:
	;
	v312 = F_pstrdup(m, v180)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L78
	}
L46:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if v298 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L47:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[0]))
	v266 = base.B2i32(v264 == int32(1))
	if v264 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L48:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v240 = int32(2)
	v241 = F_pg_ultostr_zeropad(m, v180, v239, v240)
	mBase = m.M
	v242 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v242)
	v244 = int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v248 = F_pg_ultostr_zeropad(m, v241+v244, v246, v240)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v242)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if int32(0) < v253 {
		goto L63
	} else {
		goto L64
	}
L49:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[0]))
	v210 = base.B2i32(v208 == int32(1))
	if v208 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if int32(0) < v183 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v188 = v183
	goto L53
L52:
	;
	v188 = int32(1) - v183
	goto L53
L53:
	;
	v190 = F_pg_ultostr_zeropad(m, v180, v188, int32(4))
	mBase = m.M
	v191 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v191)
	v193 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v196 = int32(2)
	v197 = F_pg_ultostr_zeropad(m, v190+v193, v195, v196)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v191)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v204 = F_pg_ultostr_zeropad(m, v197+v193, v202, v196)
	mBase = m.M
	v297 = v204
	goto L46
L54:
	;
	v211 = int32(12)
	goto L56
L55:
	;
	v211 = int32(16)
	goto L56
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v177+v211)))
	v215 = F_pg_ultostr_zeropad(m, v180, v213, int32(2))
	mBase = m.M
	v216 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v216)
	if v208 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v222 = int32(16)
	goto L59
L58:
	;
	v222 = int32(12)
	goto L59
L59:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v177+v222)))
	v226 = F_pg_ultostr_zeropad(m, v215+int32(1), v224, int32(2))
	mBase = m.M
	v227 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v227)
	v229 = int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if int32(0) < v231 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v236 = v231
	goto L62
L61:
	;
	v236 = v229 - v231
	goto L62
L62:
	;
	v238 = F_pg_ultostr_zeropad(m, v226+v229, v236, int32(4))
	mBase = m.M
	v297 = v238
	goto L46
L63:
	;
	v258 = v253
	goto L65
L64:
	;
	v258 = v244 - v253
	goto L65
L65:
	;
	v260 = F_pg_ultostr_zeropad(m, v248+v244, v258, int32(4))
	mBase = m.M
	v297 = v260
	goto L46
L66:
	;
	v267 = int32(12)
	goto L68
L67:
	;
	v267 = int32(16)
	goto L68
L68:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v177+v267)))
	v271 = F_pg_ultostr_zeropad(m, v180, v269, int32(2))
	mBase = m.M
	v272 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v272)
	if v264 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v278 = int32(16)
	goto L71
L70:
	;
	v278 = int32(12)
	goto L71
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v177+v278)))
	v282 = F_pg_ultostr_zeropad(m, v271+int32(1), v280, int32(2))
	mBase = m.M
	v283 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v283)
	v285 = int32(1)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v177)+20))
	if int32(0) < v287 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v292 = v287
	goto L74
L73:
	;
	v292 = v285 - v287
	goto L74
L74:
	;
	v294 = F_pg_ultostr_zeropad(m, v282+v285, v292, int32(4))
	mBase = m.M
	v297 = v294
	goto L46
L75:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(v297)+2)) = uint8(v302)
	v305 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[2])))
	*(*uint16)(unsafe.Add(mBase, uint32(v297))) = uint16(v305)
	v309 = v297 + int32(3)
	goto L77
L76:
	;
	v309 = v297
	goto L77
L77:
	;
	v310 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v310)
	goto L45
L78:
	;
	v485 = v312
	goto L1
L79:
	;
	if v89 == int32(1184) {
		goto L29
	} else {
		goto L80
	}
L80:
	;
	goto L32
L81:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v325 = F_OidOutputFunctionCall(m, v324, l0)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	if v89 == int32(142) {
		v485 = v325
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v329 = F_escape_xml(m, v325)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v485 = v329
	goto L1
L85:
	;
	v335 = int32(_a_F_map_sql_value_to_xml_value_8)
	goto L87
L86:
	;
	v335 = int32(_a_F_map_sql_value_to_xml_value_9)
	goto L87
L87:
	;
	v485 = v335
	goto L1
L88:
	;
	v340 = int32(0)
	v342 = v8 + int32(152)
	v347 = F_timestamp2tm(m, l0, v340, v342, v8+int32(200), v340, v340)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v347 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v352 = int32(0)
	v357 = v8 + int32(16)
	F_EncodeDateTime(m, v342, v351, v352, v352, v352, int32(4), v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
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
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L95
	}
L93:
	;
	v360 = F_pstrdup(m, v357)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v485 = v360
	goto L1
L95:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2625), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
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
	v387 = v8 + int32(152)
	v393 = F_timestamp2tm(m, l0, v8+int32(200), v387, v8+int32(196), v8+int32(12), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v393 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v403 = v8 + int32(16)
	F_EncodeDateTime(m, v387, v397, int32(1), v399, v400, int32(4), v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
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
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L106
	}
L104:
	;
	v406 = F_pstrdup(m, v403)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v485 = v406
	goto L1
L106:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2652), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
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
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_13), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v437 = F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_14), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2597), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
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
	v450 = m.ExcPending
	if v450 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	v457 = F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_15), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2619), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v477 = F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_15), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2646), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
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
					F_errcontext_msg(m, int32(_a_F_sql_compile_error_callback_0), v6)
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
						F_errcontext_msg(m, int32(_a_F_sql_compile_error_callback_0), v6)
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
									F_errcontext_msg(m, int32(_a_F_sql_compile_error_callback_0), v6)
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
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l2 != 0 {
		v146 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v146
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
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
		v146 = v4
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
	if v121 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L12:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v118 = F_sql_fn_resolve_param_name(m, v15, v45, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L25
	} else {
		goto L38
	}
L13:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if base.B2i32(v87 == int32(0))|base.B2i32(v87 != v90) != 0 {
		v108 = v87
		v109 = v90
		goto L28
	} else {
		goto L29
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(v51 == int32(0))|base.B2i32(v51 != v54) != 0 {
		v72 = v51
		v73 = v54
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v72-v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	goto L15
L17:
	;
	v57 = v45
	v58 = v48
	goto L18
L18:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v62 == int32(0) {
		v72 = v62
		v73 = v61
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v72 = v62
	v73 = v61
	goto L16
L20:
	;
	v65 = int32(1)
	if v62 == v61 {
		v57 = v57 + v65
		v58 = v58 + v65
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v146 = int32(0)
	goto L1
L23:
	;
	goto L24
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v77 = F_sql_fn_resolve_param_name(m, v15, v44, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int32(0)
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v120 = v83
	v121 = v77
	goto L11
L27:
	;
	if v108-v109 != 0 {
		goto L12
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v93 = v45
	v94 = v84
	goto L30
L30:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+1)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v98 == int32(0) {
		v108 = v98
		v109 = v97
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v108 = v98
	v109 = v97
	goto L28
L32:
	;
	v101 = int32(1)
	if v98 == v97 {
		v93 = v93 + v101
		v94 = v94 + v101
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v112 = F_sql_fn_resolve_param_name(m, v15, v44, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L25
	} else {
		goto L35
	}
L35:
	;
	if v112 != 0 {
		v146 = v112
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v115 = F_sql_fn_resolve_param_name(m, v15, v45, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v120 = v43
	v121 = v115
	goto L11
L38:
	;
	v120 = v43
	v121 = v118
	goto L11
L39:
	;
	v146 = int32(0)
	goto L1
L40:
	;
	goto L41
L41:
	;
	if v120 == int32(0) {
		v146 = v121
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v120
	v132 = F_list_make1_impl(m, int32(1), v13+int32(4))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v121
	v137 = F_list_make1_impl(m, int32(1), v13)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v140 = int32(0)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v143 = F_ParseFuncOrColumn(m, l0, v132, v137, v139, v140, v140, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v146 = v143
	goto L1
}
