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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int64
	_ = v377
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
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
	return v483
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
	v88 = F_getBaseType(m, l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	F_pfree(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
	v55 = v8 + int32(16)
	F_appendStringInfoString(m, v55, int32(_a_F_map_sql_value_to_xml_value_0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	v72 = v44
	goto L18
L18:
	;
	v74 = v45 + int32(1)
	if v74 < v72 {
		v44 = v72
		v45 = v74
		goto L14
	} else {
		goto L23
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v45<<(uint(int32(2))%32))))
	v64 = F_map_sql_value_to_xml_value(m, v63, v16)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	F_appendStringInfoString(m, v55, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v55, int32(_a_F_map_sql_value_to_xml_value_1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v8)+152))
	v72 = v71
	goto L18
L23:
	;
	goto L15
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	F_pfree(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v483 = v87
	goto L1
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L120
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L2
	} else {
		goto L115
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L2
	} else {
		goto L110
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	v377 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v377-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		goto L26
	} else {
		goto L99
	}
L30:
	;
	v332 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui64(v332-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
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
	F_getTypeOutputInfo(m, v88, v8+int32(16), v8+int32(152))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L81
	}
L33:
	;
	if v88 <= int32(1113) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v88 == int32(16) {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v88 == int32(1114) {
		goto L30
	} else {
		goto L79
	}
L37:
	;
	if v88 != int32(1082) {
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
	v111 = l0 + int32(_a_F_map_sql_value_to_xml_value_2)
	v112 = int32(_a_F_map_sql_value_to_xml_value_3)
	v113 = base.I32_div_u_s(v111, v112)
	v114 = int32(3)
	v120 = int32(2)
	v125 = base.I32_div_u_s((v113*int32(1073595727)+v111)<<(uint(v120)%32)|v114, v112)
	v128 = l0 + int32(_a_F_map_sql_value_to_xml_value_4) + v113*v114 + v125 + int32(_a_F_map_sql_value_to_xml_value_5)
	v129 = int32(1461)
	v130 = base.I32_div_u_s(v128, v129)
	v133 = v130*int32(-1461) + v128
	v135 = v133 << (uint(v120) % 32)
	if base.Ui32(v129) <= base.Ui32(v135) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v175 = v8 + int32(152)
	v178 = v8 + int32(16)
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
	v148 = base.I32_div_u_s(v135, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(172)))) = v148 + v130<<(uint(int32(2))%32) - int32(_a_F_map_sql_value_to_xml_value_6)
	v156 = v146 + int32(123)
	v160 = int32(base.Ui32(v156*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(164)))) = v156 - int32(base.Ui32(v160*int32(_a_F_map_sql_value_to_xml_value_7))>>(uint(int32(8))%32))
	v170 = base.I32_rem_u_s(v160+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(168)))) = v170 + int32(1)
	goto L40
L42:
	;
	v141 = base.I32_rem_u_s(v133+int32(305), int32(365))
	v146 = v141
	goto L41
L43:
	;
	goto L44
L44:
	;
	v145 = base.I32_rem_u_s(v133+int32(306), int32(366))
	v146 = v145
	goto L41
L45:
	;
	v310 = F_pstrdup(m, v178)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L2
	} else {
		goto L78
	}
L46:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if v296 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L47:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[0]))
	v264 = base.B2i32(v262 == int32(1))
	if v262 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L48:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v238 = int32(2)
	v239 = F_pg_ultostr_zeropad(m, v178, v237, v238)
	mBase = m.M
	v240 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v239))) = uint8(v240)
	v242 = int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v246 = F_pg_ultostr_zeropad(m, v239+v242, v244, v238)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v246))) = uint8(v240)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if int32(0) < v251 {
		goto L63
	} else {
		goto L64
	}
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[0]))
	v208 = base.B2i32(v206 == int32(1))
	if v206 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if int32(0) < v181 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v186 = v181
	goto L53
L52:
	;
	v186 = int32(1) - v181
	goto L53
L53:
	;
	v188 = F_pg_ultostr_zeropad(m, v178, v186, int32(4))
	mBase = m.M
	v189 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v189)
	v191 = int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v175)+16))
	v194 = int32(2)
	v195 = F_pg_ultostr_zeropad(m, v188+v191, v193, v194)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v189)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v202 = F_pg_ultostr_zeropad(m, v195+v191, v200, v194)
	mBase = m.M
	v295 = v202
	goto L46
L54:
	;
	v209 = int32(12)
	goto L56
L55:
	;
	v209 = int32(16)
	goto L56
L56:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v175+v209)))
	v213 = F_pg_ultostr_zeropad(m, v178, v211, int32(2))
	mBase = m.M
	v214 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v214)
	if v206 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v220 = int32(16)
	goto L59
L58:
	;
	v220 = int32(12)
	goto L59
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v175+v220)))
	v224 = F_pg_ultostr_zeropad(m, v213+int32(1), v222, int32(2))
	mBase = m.M
	v225 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v224))) = uint8(v225)
	v227 = int32(1)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if int32(0) < v229 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v234 = v229
	goto L62
L61:
	;
	v234 = v227 - v229
	goto L62
L62:
	;
	v236 = F_pg_ultostr_zeropad(m, v224+v227, v234, int32(4))
	mBase = m.M
	v295 = v236
	goto L46
L63:
	;
	v256 = v251
	goto L65
L64:
	;
	v256 = v242 - v251
	goto L65
L65:
	;
	v258 = F_pg_ultostr_zeropad(m, v246+v242, v256, int32(4))
	mBase = m.M
	v295 = v258
	goto L46
L66:
	;
	v265 = int32(12)
	goto L68
L67:
	;
	v265 = int32(16)
	goto L68
L68:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v175+v265)))
	v269 = F_pg_ultostr_zeropad(m, v178, v267, int32(2))
	mBase = m.M
	v270 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v270)
	if v262 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v276 = int32(16)
	goto L71
L70:
	;
	v276 = int32(12)
	goto L71
L71:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v175+v276)))
	v280 = F_pg_ultostr_zeropad(m, v269+int32(1), v278, int32(2))
	mBase = m.M
	v281 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v281)
	v283 = int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v175)+20))
	if int32(0) < v285 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v290 = v285
	goto L74
L73:
	;
	v290 = v283 - v285
	goto L74
L74:
	;
	v292 = F_pg_ultostr_zeropad(m, v280+v283, v290, int32(4))
	mBase = m.M
	v295 = v292
	goto L46
L75:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)) = uint8(v300)
	v303 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_map_sql_value_to_xml_value[2])))
	*(*uint16)(unsafe.Add(mBase, uint32(v295))) = uint16(v303)
	v307 = v295 + int32(3)
	goto L77
L76:
	;
	v307 = v295
	goto L77
L77:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v308)
	goto L45
L78:
	;
	v483 = v310
	goto L1
L79:
	;
	if v88 == int32(1184) {
		goto L29
	} else {
		goto L80
	}
L80:
	;
	goto L32
L81:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v323 = F_OidOutputFunctionCall(m, v322, l0)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L2
	} else {
		goto L82
	}
L82:
	;
	if v88 == int32(142) {
		v483 = v323
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v327 = F_escape_xml(m, v323)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v483 = v327
	goto L1
L85:
	;
	v331 = int32(_a_F_map_sql_value_to_xml_value_8)
	goto L87
L86:
	;
	v331 = int32(_a_F_map_sql_value_to_xml_value_9)
	goto L87
L87:
	;
	v483 = v331
	goto L1
L88:
	;
	v337 = int32(0)
	v339 = v8 + int32(152)
	v344 = F_timestamp2tm(m, v332, v337, v339, v8+int32(200), v337, v337)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	if v344 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v349 = int32(0)
	v354 = v8 + int32(16)
	F_EncodeDateTime(m, v339, v348, v349, v349, v349, int32(4), v354)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
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
	v362 = m.ExcPending
	if v362 != 0 {
		goto L2
	} else {
		goto L95
	}
L93:
	;
	v357 = F_pstrdup(m, v354)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	v483 = v357
	goto L1
L95:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2582), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
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
	v385 = v8 + int32(152)
	v391 = F_timestamp2tm(m, v377, v8+int32(200), v385, v8+int32(196), v8+int32(12), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	if v391 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v8)+196))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v8)+200))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v401 = v8 + int32(16)
	F_EncodeDateTime(m, v385, v395, int32(1), v397, v398, int32(4), v401)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
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
	v409 = m.ExcPending
	if v409 != 0 {
		goto L2
	} else {
		goto L106
	}
L104:
	;
	v404 = F_pstrdup(m, v401)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L105
	}
L105:
	;
	v483 = v404
	goto L1
L106:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2609), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
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
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_13), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_14), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2554), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_15), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2576), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
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
	v468 = m.ExcPending
	if v468 != 0 {
		goto L2
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_map_sql_value_to_xml_value_10), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	F_errdetail(m, int32(_a_F_map_sql_value_to_xml_value_15), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_map_sql_value_to_xml_value_11), int32(2603), int32(_a_F_map_sql_value_to_xml_value_12))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
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
