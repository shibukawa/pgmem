package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datum_to_jsonb_internal(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
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
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int64
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
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int64
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
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
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int64
	_ = v393
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	m.G0 = v15 + int32(160)
	return
L4:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v615 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(0)
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
		goto L179
	}
L9:
	;
	v576 = v15 + int32(32)
	v578 = v15 + int32(72)
	v580 = F_JsonbIteratorNext(m, v576, v578, int32(1))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L177
	}
L10:
	;
	v26 = l5
	goto L12
L11:
	;
	v26 = int32(0)
	goto L12
L12:
	;
	if v26 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	switch l3 - int32(1) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	case 3:
		goto L21
	case 4:
		goto L20
	case 5:
		v395 = l0
		goto L18
	case 6:
		goto L17
	case 7:
		goto L26
	case 8:
		goto L25
	case 9:
		goto L19
	default:
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L173
	}
L16:
	;
	switch l4 - int32(1045) {
	case 0, 2:
		goto L147
	case 1:
		goto L146
	default:
		goto L148
	}
L17:
	;
	v430 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L132
	}
L18:
	;
	v397 = v15 + int32(72)
	v399 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v395))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L128
	}
L19:
	;
	v393 = F_OidFunctionCall1Coll(m, l4, int32(0), l0)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L127
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	v384 = int32(0)
	v387 = F_JsonEncodeDateTime(m, v384, l0, int32(1184), v384)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L126
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	v374 = int32(0)
	v377 = F_JsonEncodeDateTime(m, v374, l0, int32(1114), v374)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L125
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	v364 = int32(0)
	v367 = F_JsonEncodeDateTime(m, v364, l0, int32(1082), v364)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L124
	}
L23:
	;
	if l5 != 0 {
		goto L102
	} else {
		goto L103
	}
L24:
	;
	if l5 != 0 {
		goto L92
	} else {
		goto L93
	}
L25:
	;
	v96 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L40
	}
L26:
	;
	v32 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v40 = v32 + int32(16)
	v41 = F_ArrayGetNItemsSafe(m, v38, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v41
	if v41 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_pushJsonbValue(m, l2, int32(4), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_get_typlenbyvalalign(m, v34, v15+int32(150), v15+int32(149), v15+int32(148))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v589 = v35
	goto L8
L34:
	;
	F_json_categorize_type(m, v34, int32(1), v15+int32(144), v15+int32(140))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+150)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+149)))
	v71 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+148)))
	F_deconstruct_array(m, v32, v69, v70, v71, v15+int32(156), v15+int32(152), v15+int32(72))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v15)+156))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15)+144))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v15)+140))
	F_array_dim_to_jsonb(m, l2, int32(0), v38, v40, v81, v82, v15+int32(32), v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+156))
	F_pfree(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	F_pfree(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v589 = v35
	goto L8
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v100 = F_lookup_rowtype_tupdesc(m, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(base.Ui32(v102) >> (uint(int32(2)) % 32))
	F_pushJsonbValue(m, l2, int32(6), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if int32(0) < v111 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v118 = int32(0)
	v124 = v111
	goto L46
L44:
	;
	goto L45
L45:
	;
	F_pushJsonbValue(m, l2, int32(7), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L87
	}
L46:
	;
	v134 = v100 + v124<<(uint(int32(3))%32) + v118*int32(100)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+119)))
	if v135 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v269 = v118 + int32(1)
	goto L50
L49:
	;
	v140 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v140
	v145 = v134 + int32(32)
	v146 = F_strlen(m, v145)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v146
	F_pushJsonbValue(m, l2, v140, v15+int32(72))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v269 < v270 {
		v118 = v269
		v124 = v270
		goto L46
	} else {
		goto L86
	}
L51:
	;
	v155 = v118 + int32(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+18)))
	if base.Ui32(v157&int32(2047)) <= base.Ui32(v118) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	if v235 == int32(1) {
		goto L81
	} else {
		goto L82
	}
L53:
	;
	v163 = F_getmissingattr(m, v100, v155, v15+int32(144))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)) = uint8(v165)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+20)))
	if v167&int32(1) == v165 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v234 = v163
	goto L52
L57:
	;
	v174 = v100 + int32(20) + v155<<(uint(int32(3))%32)
	v175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v174))))
	if int32(0) <= v175 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+int32(base.Ui32(v118)>>(uint(int32(3))%32)))+23)))
	if int32(base.Ui32(v216)>>(uint(v118&int32(7))%32))&int32(1) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L60:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+22)))
	v180 = v156 + v178 + v175
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
	if v181 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v211 = F_nocachegetattr(m, v15+int32(32), v155, v100)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L75
	}
L63:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v174)+2)))
	if base.I32_popcnt(v184) != int32(1) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v234 = base.I64_extend_i32_u(v180)
	goto L52
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L72
	}
L67:
	;
	switch base.I32_ctz(v184) {
	case 0:
		goto L71
	case 1:
		goto L70
	case 2:
		goto L69
	case 3:
		goto L68
	default:
		goto L66
	}
L68:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	v234 = v192
	goto L52
L69:
	;
	v191 = int64(*(*int32)(unsafe.Add(mBase, uint32(v180))))
	v234 = v191
	goto L52
L70:
	;
	v190 = int64(*(*int16)(unsafe.Add(mBase, uint32(v180))))
	v234 = v190
	goto L52
L71:
	;
	v189 = int64(*(*int8)(unsafe.Add(mBase, uint32(v180))))
	v234 = v189
	goto L52
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v184
	F_errmsg_internal(m, int32(_a_F_datum_to_jsonb_internal_0), v15+int32(16))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_1), int32(123), int32(_a_F_datum_to_jsonb_internal_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
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
	v234 = v211
	goto L52
L76:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)) = uint8(v224)
	v234 = int64(0)
	goto L52
L77:
	;
	goto L78
L78:
	;
	v229 = F_nocachegetattr(m, v15+int32(32), v155, v100)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v234 = v229
	goto L52
L80:
	;
	F_datum_to_jsonb_internal(m, v234, v256&int32(1), l2, v257, v258, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L85
	}
L81:
	;
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v238
	v256 = int32(1)
	v257 = v238
	v258 = v238
	goto L80
L82:
	;
	goto L83
L83:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v134+int32(28))+68))
	F_json_categorize_type(m, v245, int32(1), v15+int32(156), v15+int32(152))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+144)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v15)+156))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v15)+152))
	v256 = v253
	v257 = v254
	v258 = v255
	goto L80
L85:
	;
	v269 = v155
	goto L50
L86:
	;
	goto L47
L87:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	if int32(0) <= v288 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_DecrTupleDescRefCount(m, v100)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v589 = int32(0)
	goto L8
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	v299 = base.B2i32(l0 == int64(0))
	if l0 == int64(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+80)) = uint8(base.B2i32(l0 != int64(0)))
	goto L4
L95:
	;
	v300 = int32(_a_F_datum_to_jsonb_internal_3)
	goto L97
L96:
	;
	v300 = int32(_a_F_datum_to_jsonb_internal_4)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v300
	if l0 == int64(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v304 = int32(5)
	goto L100
L99:
	;
	v304 = int32(4)
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v304
	goto L4
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(2)
	goto L4
L102:
	;
	v352 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L123
	}
L103:
	;
	switch l4 - int32(39) {
	case 0:
		goto L108
	case 1, 2, 3:
		goto L105
	case 4:
		goto L107
	default:
		goto L109
	}
L104:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+4)))
	goto L119
L105:
	;
	v330 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L116
	}
L106:
	;
	v326 = F_int64_to_numeric(m, l0)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L115
	}
L107:
	;
	v324 = F_int64_to_numeric(m, base.I64_extend32_s(l0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L114
	}
L108:
	;
	v321 = F_int64_to_numeric(m, base.I64_extend16_s(l0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L113
	}
L109:
	;
	if l4 == int32(461) {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	if l4 != int32(1702) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v318 = F_pg_detoast_datum(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v340 = v318
	goto L104
L113:
	;
	v340 = v321
	goto L104
L114:
	;
	v340 = v324
	goto L104
L115:
	;
	v340 = v326
	goto L104
L116:
	;
	v335 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v330), int64(0), int64(-1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v338 = F_pg_detoast_datum(m, base.I32_wrap_i64(v335))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v340 = v338
	goto L104
L119:
	;
	if v341&int32(_a_F_datum_to_jsonb_internal_5) == int32(_a_F_datum_to_jsonb_internal_6) {
		goto L102
	} else {
		goto L120
	}
L120:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v340)+4)))
	goto L121
L121:
	;
	if base.B2i32(v346 == int32(_a_F_datum_to_jsonb_internal_7)) == int32(0) {
		goto L101
	} else {
		goto L122
	}
L122:
	;
	goto L102
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	v356 = F_strlen(m, v352)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v356
	goto L4
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v367
	v370 = F_strlen(m, v367)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v370
	goto L4
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v377
	v380 = F_strlen(m, v377)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v380
	goto L4
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v387
	v390 = F_strlen(m, v387)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v390
	goto L4
L127:
	;
	v395 = v393
	goto L18
L128:
	;
	F_makeJsonLexContext(m, v397, v399, int32(1))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v404 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v404
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = int32(1451)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(1452)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = int32(1453)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(1454)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = int32(1455)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = int32(1456)
	v425 = F_pg_parse_json_or_errsave(m, v397, v15+int32(32), v404)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_freeJsonLexContext(m, v397)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v589 = v404
	goto L8
L132:
	;
	v434 = F_JsonbIteratorInit(m, v430+int32(4))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v434
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+7)))
	v441 = int32(base.Ui32(v437&int32(16)) >> (uint(int32(4)) % 32))
	if v441 != 0 {
		goto L9
	} else {
		goto L134
	}
L134:
	;
	v447 = F_JsonbIteratorNext(m, v15+int32(32), v15+int32(72), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	if v447 == int32(0) {
		v589 = v441
		goto L8
	} else {
		goto L136
	}
L136:
	;
	v452 = v447
	goto L137
L137:
	;
	v464 = v15 + int32(72)
	if v452&int32(-4) != int32(4) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v589 = v441
	goto L8
L139:
	;
	v470 = v464
	goto L141
L140:
	;
	v470 = int32(0)
	goto L141
L141:
	;
	F_pushJsonbValue(m, l2, v452, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v476 = F_JsonbIteratorNext(m, v15+int32(32), v464, int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	if v476 != 0 {
		v452 = v476
		goto L137
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v527
	v533 = int32(0)
	if base.Ui32(v529) < base.Ui32(int32(268435456)) {
		v589 = v533
		goto L8
	} else {
		goto L166
	}
L146:
	;
	v523 = F_OidOutputFunctionCall(m, l4, l0)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L165
	}
L147:
	;
	v483 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L151
	}
L148:
	;
	if l4 != int32(47) {
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v514
	v516 = int32(1)
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	if v518&v516 != 0 {
		goto L162
	} else {
		goto L163
	}
L151:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483))))
	if v485 == int32(1) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+1)))
	if v491 == int32(18) {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	goto L154
L154:
	;
	v502 = int32(1)
	if v485&v502 != 0 {
		v514 = int32(base.Ui32(v485)>>(uint(v502)%32)) - v502
		goto L150
	} else {
		goto L161
	}
L155:
	;
	v494 = int32(16)
	goto L157
L156:
	;
	v494 = int32(0)
	goto L157
L157:
	;
	if base.Ui32((v491-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v501 = int32(4)
	goto L160
L159:
	;
	v501 = v494
	goto L160
L160:
	;
	v514 = v501
	goto L150
L161:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v483)))
	v514 = int32(base.Ui32(v508)>>(uint(int32(2))%32)) - int32(4)
	goto L150
L162:
	;
	v521 = v516
	goto L164
L163:
	;
	v521 = int32(4)
	goto L164
L164:
	;
	v527 = v483 + v521
	v529 = v514
	goto L145
L165:
	;
	v525 = F_strlen(m, v523)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v525
	v527 = v523
	v529 = v525
	goto L145
L166:
	;
	v537 = F_errsave_start(m, int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	if v537 == int32(0) {
		v589 = v533
		goto L8
	} else {
		goto L168
	}
L168:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_datum_to_jsonb_internal_8), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(268435455)
	v551 = F_errdetail(m, int32(_a_F_datum_to_jsonb_internal_9), v15)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_datum_to_jsonb_internal_10), int32(276), int32(_a_F_datum_to_jsonb_internal_11))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v589 = v533
	goto L8
L173:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errmsg(m, int32(_a_F_datum_to_jsonb_internal_12), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_10), int32(657), int32(_a_F_datum_to_jsonb_internal_13))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L177:
	;
	v583 = F_JsonbIteratorNext(m, v576, v578, int32(1))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v589 = v441
	goto L8
L179:
	;
	if v589 == int32(0) {
		goto L3
	} else {
		goto L180
	}
L180:
	;
	goto L4
L181:
	;
	v618 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)) = uint8(v618)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v618
	F_pushJsonbValue(m, l2, int32(4), v15+int32(32))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v615)))
	switch v638 - int32(16) {
	case 0:
		goto L187
	case 1:
		goto L189
	default:
		goto L188
	}
L184:
	;
	F_pushJsonbValue(m, l2, int32(3), v15+int32(72))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_pushJsonbValue(m, l2, int32(5), int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L3
L187:
	;
	F_pushJsonbValue(m, l2, int32(3), v15+int32(72))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L197
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L1
	} else {
		goto L194
	}
L189:
	;
	if l5 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v643 = int32(1)
	goto L192
L191:
	;
	v643 = int32(2)
	goto L192
L192:
	;
	F_pushJsonbValue(m, l2, v643, v15+int32(72))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	goto L3
L194:
	;
	F_errmsg_internal(m, int32(_a_F_datum_to_jsonb_internal_14), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_datum_to_jsonb_internal_10), int32(887), int32(_a_F_datum_to_jsonb_internal_13))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	goto L3
}
