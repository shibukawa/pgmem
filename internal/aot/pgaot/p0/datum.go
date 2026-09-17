package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datum_to_jsonb_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
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
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
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
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	F_check_stack_depth(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
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
	m.G0 = v14 + int32(160)
	return
L4:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v558 == int32(0) {
		goto L152
	} else {
		goto L153
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(l3-int32(6)) <= base.Ui32(int32(4)) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if base.Ui32(l3-int32(11)) < base.Ui32(int32(-5)) {
		goto L4
	} else {
		goto L150
	}
L9:
	;
	v521 = v14 + int32(32)
	v523 = v14 + int32(72)
	v525 = F_JsonbIteratorNext(m, v521, v523, int32(1))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L148
	}
L10:
	;
	v25 = l5
	goto L12
L11:
	;
	v25 = int32(0)
	goto L12
L12:
	;
	if v25 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	switch l3 - int32(1) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	case 3:
		goto L22
	case 4:
		goto L21
	case 5:
		v379 = l0
		goto L19
	case 6:
		goto L18
	case 7:
		goto L27
	case 8:
		goto L26
	case 9:
		goto L20
	default:
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L144
	}
L16:
	;
	v530 = int32(0)
	goto L8
L17:
	;
	v460 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L135
	}
L18:
	;
	v412 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L122
	}
L19:
	;
	v381 = v14 + int32(72)
	v382 = F_pg_detoast_datum_packed(m, v379)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L118
	}
L20:
	;
	v377 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L117
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v368 = int32(0)
	v371 = F_JsonEncodeDateTime(m, v368, l0, int32(1184), v368)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L116
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v358 = int32(0)
	v361 = F_JsonEncodeDateTime(m, v358, l0, int32(1114), v358)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L115
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v348 = int32(0)
	v351 = F_JsonEncodeDateTime(m, v348, l0, int32(1082), v348)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L114
	}
L24:
	;
	v305 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L96
	}
L25:
	;
	if l5 != 0 {
		goto L87
	} else {
		goto L88
	}
L26:
	;
	v95 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L41
	}
L27:
	;
	v30 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v33
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = v30 + int32(16)
	v39 = F_ArrayGetNItemsSafe(m, v36, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = v39
	if v39 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v46 = F_pushJsonbValue(m, l2, int32(4), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_get_typlenbyvalalign(m, v32, v14+int32(150), v14+int32(149), v14+int32(148))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v46
	v51 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v51
	v530 = v33
	goto L8
L35:
	;
	F_json_categorize_type(m, v32, int32(1), v14+int32(144), v14+int32(140))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+150)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+149)))
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+148)))
	F_deconstruct_array(m, v30, v69, v70, v71, v14+int32(156), v14+int32(152), v14+int32(72))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v14)+144))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v14)+140))
	F_array_dim_to_jsonb(m, l2, int32(0), v36, v38, v81, v82, v14+int32(32), v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	F_pfree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v530 = v33
	goto L8
L41:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v99 = F_lookup_rowtype_tupdesc(m, v97, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(base.Ui32(v101) >> (uint(int32(2)) % 32))
	v108 = F_pushJsonbValue(m, l2, int32(6), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if int32(0) < v111 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v118 = int32(0)
	v121 = v111
	goto L47
L45:
	;
	goto L46
L46:
	;
	v282 = F_pushJsonbValue(m, l2, int32(7), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L84
	}
L47:
	;
	v133 = v99 + v121<<(uint(int32(4))%32) + v118*int32(100)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+111)))
	if v134 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v266 = v118 + int32(1)
	goto L51
L50:
	;
	v139 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v139
	v144 = v133 + int32(24)
	v145 = F_strlen(m, v144)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v145
	v151 = F_pushJsonbValue(m, l2, v139, v14+int32(32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v266 < v267 {
		v118 = v266
		v121 = v267
		goto L47
	} else {
		goto L83
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v151
	v155 = v118 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+18)))
	if base.Ui32(v157&int32(2047)) <= base.Ui32(v118) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	if v232 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L54:
	;
	v163 = F_getmissingattr(m, v99, v155, v14+int32(144))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)) = uint8(v165)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+20)))
	if v167&int32(1) == v165 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v231 = v163
	goto L53
L58:
	;
	v174 = v99 + int32(4) + v155<<(uint(int32(4))%32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if int32(0) <= v175 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(base.Ui32(v118)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v213)>>(uint(v118&int32(7))%32))&int32(1) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+22)))
	v180 = v156 + v178 + v175
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
	if v181 != int32(1) {
		v231 = v180
		goto L53
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v208 = F_nocachegetattr(m, v14+int32(72), v155, v99)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L72
	}
L64:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174)+4)))
	switch v184 - int32(1) {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L65
	case 3:
		goto L66
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L69
	}
L66:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v231 = v189
	goto L53
L67:
	;
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v180))))
	v231 = v188
	goto L53
L68:
	;
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v180))))
	v231 = v187
	goto L53
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = base.I32_extend16_s(v184)
	F_errmsg_internal(m, int32(_a_F_datum_to_jsonb_internal_0), v14+int32(16))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_1), int32(70), int32(_a_F_datum_to_jsonb_internal_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v231 = v208
	goto L53
L73:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)) = uint8(v221)
	v231 = int32(0)
	goto L53
L74:
	;
	goto L75
L75:
	;
	v226 = F_nocachegetattr(m, v14+int32(72), v155, v99)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v231 = v226
	goto L53
L77:
	;
	F_datum_to_jsonb_internal(m, v231, v253&int32(1), l2, v254, v255, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L82
	}
L78:
	;
	v235 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v14)+156)) = v235
	v253 = int32(1)
	v254 = v235
	v255 = v235
	goto L77
L79:
	;
	goto L80
L80:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v133+int32(20))+68))
	F_json_categorize_type(m, v242, int32(1), v14+int32(156), v14+int32(152))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+144)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v14)+156))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	v253 = v250
	v254 = v251
	v255 = v252
	goto L77
L82:
	;
	v266 = v155
	goto L51
L83:
	;
	goto L48
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	if v285 < int32(0) {
		goto L16
	} else {
		goto L85
	}
L85:
	;
	F_DecrTupleDescRefCount(m, v99)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L16
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	if l0 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+76)) = uint8(base.B2i32(l0 != int32(0)))
	goto L4
L90:
	;
	v294 = int32(_a_F_datum_to_jsonb_internal_3)
	goto L92
L91:
	;
	v294 = int32(_a_F_datum_to_jsonb_internal_4)
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v294
	if l0 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v298 = int32(4)
	goto L95
L94:
	;
	v298 = int32(5)
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v298
	goto L4
L96:
	;
	if l5 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v309 = F_strlen(m, v305)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v309
	goto L4
L98:
	;
	goto L99
L99:
	;
	v312 = int32(78)
	v313 = F___strchrnul(m, v305, v312)
	mBase = m.M
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313))))
	if v315 == v312 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v343 = F_strlen(m, v305)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v343
	goto L4
L101:
	;
	if v319 != 0 {
		goto L100
	} else {
		goto L105
	}
L102:
	;
	v319 = v313
	goto L104
L103:
	;
	v319 = int32(0)
	goto L104
L104:
	;
	goto L101
L105:
	;
	v320 = int32(110)
	v321 = F___strchrnul(m, v305, v320)
	mBase = m.M
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v323 == v320 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v327 != 0 {
		goto L100
	} else {
		goto L110
	}
L107:
	;
	v327 = v321
	goto L109
L108:
	;
	v327 = int32(0)
	goto L109
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(2)
	v331 = int32(0)
	v334 = F_DirectFunctionCall3Coll(m, int32(408), v331, v305, v331, int32(-1))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v336 = F_pg_detoast_datum(m, v334)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v336
	F_pfree(m, v305)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L4
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v351
	v354 = F_strlen(m, v351)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v354
	goto L4
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v361
	v364 = F_strlen(m, v361)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v364
	goto L4
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v371
	v374 = F_strlen(m, v371)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v374
	goto L4
L117:
	;
	v379 = v377
	goto L19
L118:
	;
	F_makeJsonLexContext(m, v381, v382, int32(1))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v387 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v387
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = int32(1310)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = int32(1311)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = int32(1312)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = int32(1313)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(1314)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(1315)
	v408 = F_pg_parse_json_or_errsave(m, v381, v14+int32(32), v387)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_freeJsonLexContext(m, v381)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v530 = v387
	goto L8
L122:
	;
	v416 = F_JsonbIteratorInit(m, v412+int32(4))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v416
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412)+7)))
	v423 = int32(base.Ui32(v419&int32(16)) >> (uint(int32(4)) % 32))
	if v423 != 0 {
		goto L9
	} else {
		goto L124
	}
L124:
	;
	v429 = F_JsonbIteratorNext(m, v14+int32(32), v14+int32(72), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v429 == int32(0) {
		v530 = v423
		goto L8
	} else {
		goto L126
	}
L126:
	;
	v434 = v429
	goto L127
L127:
	;
	v445 = v14 + int32(72)
	if v434&int32(-4) != int32(4) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v530 = v423
	goto L8
L129:
	;
	v451 = v445
	goto L131
L130:
	;
	v451 = int32(0)
	goto L131
L131:
	;
	v452 = F_pushJsonbValue(m, l2, v434, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v452
	v458 = F_JsonbIteratorNext(m, v14+int32(32), v445, int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	if v458 != 0 {
		v434 = v458
		goto L127
	} else {
		goto L134
	}
L134:
	;
	goto L128
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = int32(1)
	v464 = F_strlen(m, v460)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+76)) = v464
	if base.Ui32(v464) < base.Ui32(int32(268435456)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v460
	goto L16
L137:
	;
	v469 = F_errsave_start(m, int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	if v469 == int32(0) {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_datum_to_jsonb_internal_5), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(268435455)
	F_errdetail(m, int32(_a_F_datum_to_jsonb_internal_6), v14)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_datum_to_jsonb_internal_7), int32(284), int32(_a_F_datum_to_jsonb_internal_8))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	goto L136
L144:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_datum_to_jsonb_internal_9), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_7), int32(664), int32(_a_F_datum_to_jsonb_internal_10))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v528 = F_JsonbIteratorNext(m, v521, v523, int32(1))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v530 = v423
	goto L8
L150:
	;
	if v530 == int32(0) {
		goto L3
	} else {
		goto L151
	}
L151:
	;
	goto L4
L152:
	;
	v561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v561)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(4294967312)
	v568 = F_pushJsonbValue(m, l2, int32(4), v14+int32(32))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	switch v582 - int32(16) {
	case 0:
		goto L158
	case 1:
		goto L160
	default:
		goto L159
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v568
	v574 = F_pushJsonbValue(m, l2, int32(3), v14+int32(72))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v574
	v579 = F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v579
	goto L3
L158:
	;
	v609 = F_pushJsonbValue(m, l2, int32(3), v14+int32(72))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L168
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L165
	}
L160:
	;
	if l5 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v587 = int32(1)
	goto L163
L162:
	;
	v587 = int32(2)
	goto L163
L163:
	;
	v590 = F_pushJsonbValue(m, l2, v587, v14+int32(72))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v590
	goto L3
L165:
	;
	F_errmsg_internal(m, int32(_a_F_datum_to_jsonb_internal_11), int32(0))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_7), int32(851), int32(_a_F_datum_to_jsonb_internal_10))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v609
	goto L3
}
