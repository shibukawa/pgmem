package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datum_to_jsonb_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
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
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v17 + int32(160)
	return
L4:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v584 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if l5 != 0 {
		goto L23
	} else {
		goto L24
	}
L8:
	;
	if base.Ui32(int32(4)) < base.Ui32(l3-int32(6)) {
		goto L4
	} else {
		goto L147
	}
L9:
	;
	v520 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L1
	} else {
		goto L138
	}
L10:
	;
	v443 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L122
	}
L11:
	;
	v408 = F_pg_detoast_datum_packed(m, v404)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L118
	}
L12:
	;
	v402 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L117
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v393 = int32(0)
	v396 = F_JsonEncodeDateTime(m, v393, l0, int32(1184), v393)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L116
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v383 = int32(0)
	v386 = F_JsonEncodeDateTime(m, v383, l0, int32(1114), v383)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L115
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v373 = int32(0)
	v376 = F_JsonEncodeDateTime(m, v373, l0, int32(1082), v373)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L114
	}
L16:
	;
	v335 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L100
	}
L17:
	;
	v328 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L98
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+76)) = uint8(base.B2i32(l0 != int32(0)))
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	if l0 != 0 {
		goto L92
	} else {
		goto L93
	}
L20:
	;
	v113 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L44
	}
L21:
	;
	v47 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	switch l3 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L17
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	default:
		goto L9
	}
L23:
	;
	if base.Ui32(int32(4)) < base.Ui32(l3-int32(6)) {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	switch l3 - int32(1) {
	case 0:
		goto L18
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		v404 = l0
		goto L11
	case 6:
		goto L10
	case 7:
		goto L21
	case 8:
		goto L20
	case 9:
		goto L12
	default:
		goto L9
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(245218), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(499796), int32(664), int32(312276))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v54 = v47 + int32(16)
	v55 = F_ArrayGetNItems(m, v52, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v55
	if v55 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v62 = F_pushJsonbValue(m, l2, int32(4), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_get_typlenbyvalalign(m, v49, v17+int32(150), v17+int32(149), v17+int32(148))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v62
	v67 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v67
	v553 = int32(1)
	goto L8
L38:
	;
	v79 = int32(1)
	F_json_categorize_type(m, v49, v79, v17+int32(144), v17+int32(140))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+150)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+149)))
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+148)))
	F_deconstruct_array(m, v47, v87, v88, v89, v17+int32(156), v17+int32(152), v17+int32(72))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+144))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)+140))
	F_array_dim_to_jsonb(m, l2, int32(0), v52, v54, v99, v100, v17+int32(32), v103, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	F_pfree(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	F_pfree(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v553 = v79
	goto L8
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v117 = F_lookup_rowtype_tupdesc(m, v115, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(base.Ui32(v119) >> (uint(int32(2)) % 32))
	v126 = F_pushJsonbValue(m, l2, int32(6), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if int32(0) < v129 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v133 = v117 + int32(20)
	v135 = v129
	v136 = int32(0)
	goto L50
L48:
	;
	goto L49
L49:
	;
	v304 = F_pushJsonbValue(m, l2, int32(7), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L87
	}
L50:
	;
	v154 = v133 + v135<<(uint(int32(4))%32) + v136*int32(100)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+91)))
	if v155 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v285 = v136 + int32(1)
	goto L54
L53:
	;
	v160 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v160
	v163 = v154 + int32(4)
	v164 = F_strlen(m, v163)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v164
	v170 = F_pushJsonbValue(m, l2, v160, v17+int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v285 < v286 {
		v135 = v286
		v136 = v285
		goto L50
	} else {
		goto L86
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v170
	v174 = v136 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v17)+88))
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v175)+18)))
	if base.Ui32(v176&int32(2047)) <= base.Ui32(v136) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)))
	if v251 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L57:
	;
	v182 = F_getmissingattr(m, v117, v174, v17+int32(144))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)) = uint8(v184)
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+20)))
	if v186&int32(1) == v184 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v250 = v182
	goto L56
L61:
	;
	v193 = v133 + v136<<(uint(int32(4))%32)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	if int32(0) <= v194 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+int32(base.Ui32(v136)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v232)>>(uint(v136&int32(7))%32))&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+22)))
	v199 = v175 + v197 + v194
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
	if v200 != int32(1) {
		v250 = v199
		goto L56
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v227 = F_nocachegetattr(m, v17+int32(72), v174, v117)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L75
	}
L67:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+4)))
	switch v203 - int32(1) {
	case 0:
		goto L71
	case 1:
		goto L70
	default:
		goto L68
	case 3:
		goto L69
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L72
	}
L69:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v250 = v208
	goto L56
L70:
	;
	v207 = int32(*(*int16)(unsafe.Add(mBase, uint32(v199))))
	v250 = v207
	goto L56
L71:
	;
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v199))))
	v250 = v206
	goto L56
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = base.I32_extend16_s(v203)
	F_errmsg_internal(m, int32(483040), v17+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(326391), int32(70), int32(67779))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v250 = v227
	goto L56
L76:
	;
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)) = uint8(v240)
	v250 = int32(0)
	goto L56
L77:
	;
	goto L78
L78:
	;
	v245 = F_nocachegetattr(m, v17+int32(72), v174, v117)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v250 = v245
	goto L56
L80:
	;
	F_datum_to_jsonb_internal(m, v250, v272&int32(1), l2, v273, v274, int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	v254 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v254
	v272 = int32(1)
	v273 = v254
	v274 = v254
	goto L80
L82:
	;
	goto L83
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v154)+68))
	F_json_categorize_type(m, v261, int32(1), v17+int32(156), v17+int32(152))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+144)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v17)+156))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v272 = v269
	v273 = v270
	v274 = v271
	goto L80
L85:
	;
	v285 = v174
	goto L54
L86:
	;
	goto L51
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v304
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if int32(0) <= v307 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_DecrTupleDescRefCount(m, v117)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v553 = int32(1)
	goto L8
L91:
	;
	goto L90
L92:
	;
	v317 = int32(344091)
	goto L94
L93:
	;
	v317 = int32(361107)
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v317
	if l0 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v321 = int32(4)
	goto L97
L96:
	;
	v321 = int32(5)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v321
	goto L4
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v332 = F_strlen(m, v328)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v332
	goto L4
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v368 = F_strlen(m, v335)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v368
	goto L4
L100:
	;
	v337 = int32(78)
	v338 = F___strchrnul(m, v335, v337)
	mBase = m.M
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v340 == v337 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v344 != 0 {
		goto L99
	} else {
		goto L105
	}
L102:
	;
	v344 = v338
	goto L104
L103:
	;
	v344 = int32(0)
	goto L104
L104:
	;
	goto L101
L105:
	;
	v345 = int32(110)
	v346 = F___strchrnul(m, v335, v345)
	mBase = m.M
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v348 == v345 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v352 != 0 {
		goto L99
	} else {
		goto L110
	}
L107:
	;
	v352 = v346
	goto L109
L108:
	;
	v352 = int32(0)
	goto L109
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(2)
	v356 = int32(0)
	v359 = F_DirectFunctionCall3Coll(m, int32(408), v356, v335, v356, int32(-1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v361 = F_pg_detoast_datum(m, v359)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v361
	F_pfree(m, v335)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L4
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v376
	v379 = F_strlen(m, v376)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v379
	goto L4
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v386
	v389 = F_strlen(m, v386)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v389
	goto L4
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v396
	v399 = F_strlen(m, v396)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v399
	goto L4
L117:
	;
	v404 = v402
	goto L11
L118:
	;
	F_makeJsonLexContext(m, v17+int32(72), v408, int32(1))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v415 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17-int32(-64)))) = v415
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = int32(1326)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(1327)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(1328)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = int32(1329)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = int32(1330)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = int32(1331)
	v437 = F_pg_parse_json_or_errsave(m, v17+int32(72), v17+int32(32), v415)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_freeJsonLexContext(m, v17+int32(72))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v553 = int32(1)
	goto L8
L122:
	;
	v447 = F_JsonbIteratorInit(m, v443+int32(4))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v447
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v452 = v450 & int32(268435456)
	if v452 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v553 = base.B2i32(v452 == int32(0))
	goto L8
L125:
	;
	goto L128
L126:
	;
	goto L127
L127:
	;
	v494 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), int32(1))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L136
	}
L128:
	;
	v469 = int32(0)
	v475 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), v469)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v477 = int32(4)
	if base.Ui32(v477) <= base.Ui32(v475-v477) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	if v475 == int32(0) {
		goto L124
	} else {
		goto L134
	}
L132:
	;
	v485 = v469
	goto L133
L133:
	;
	v486 = F_pushJsonbValue(m, l2, v475, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L135
	}
L134:
	;
	v485 = v17 + int32(72)
	goto L133
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v486
	goto L128
L136:
	;
	v501 = F_JsonbIteratorNext(m, v17+int32(32), v17+int32(72), int32(1))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	goto L124
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = int32(1)
	v524 = F_strlen(m, v520)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v524
	if base.Ui32(v524) < base.Ui32(int32(268435456)) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v520
	v553 = int32(1)
	goto L8
L140:
	;
	v529 = F_errsave_start(m, int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if v529 == int32(0) {
		goto L139
	} else {
		goto L142
	}
L142:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errmsg(m, int32(330327), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(268435455)
	F_errdetail(m, int32(596915), v17)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errsave_finish(m, int32(0), int32(499796), int32(284), int32(282385))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	goto L139
L147:
	;
	if v553 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	goto L4
L149:
	;
	v587 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)) = uint8(v587)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+32)) = int64(4294967312)
	v594 = F_pushJsonbValue(m, l2, int32(4), v17+int32(32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	switch v608 - int32(16) {
	case 0:
		goto L155
	case 1:
		goto L157
	default:
		goto L156
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v594
	v600 = F_pushJsonbValue(m, l2, int32(3), v17+int32(72))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v600
	v605 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v605
	goto L3
L155:
	;
	v635 = F_pushJsonbValue(m, l2, int32(3), v17+int32(72))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L165
	}
L156:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L162
	}
L157:
	;
	if l5 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v613 = int32(1)
	goto L160
L159:
	;
	v613 = int32(2)
	goto L160
L160:
	;
	v616 = F_pushJsonbValue(m, l2, v613, v17+int32(72))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v616
	goto L3
L162:
	;
	F_errmsg_internal(m, int32(362490), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(499796), int32(851), int32(312276))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v635
	goto L3
}
