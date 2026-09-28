package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_NUM_cache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v584 int32
	_ = v584
	var v586 int64
	_ = v586
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	v5 = int32(0)
	v13 = F_text_to_cstring(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui32(int32(57)) <= base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = F_palloc_mul(m, int32(16), l0+int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	if int32(2147483646) <= v49 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v24)
	v26 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v26
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v28
	F_parse_format(m, v22, v13, int32(_a_F_NUM_cache_0), v26, int32(_a_F_NUM_cache_1), int32(2), l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_pfree(m, v13)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	return v22
L9:
	;
	if v47 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v159 = v49
	goto L11
L11:
	;
	if v47 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v151 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v151
	v159 = v151
	goto L11
L13:
	;
	v55 = v47 & int32(3)
	v56 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v47) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v61 = v56
	v71 = v5
	goto L17
L15:
	;
	v102 = v56
	goto L16
L16:
	;
	v114 = v102
	v125 = v5
	goto L21
L17:
	;
	v74 = v61 << (uint(int32(2)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_NUM_cache[2])))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+972))
	v77 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75)+972)) = v76 >> (uint(v77) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_NUM_cache[3])))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+972))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+972)) = v81 >> (uint(v77) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_NUM_cache[4])))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+972))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+972)) = v86 >> (uint(v77) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_NUM_cache[5])))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+972))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+972)) = v91 >> (uint(v77) % 32)
	v95 = int32(4)
	v96 = v61 + v95
	v98 = v71 + v95
	if v98 != v47&int32(2147483644) {
		v61 = v96
		v71 = v98
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v55 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v102 = v96
	goto L16
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v114<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+972))
	v130 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+972)) = v129 >> (uint(v130) % 32)
	v136 = v125 + v130
	if v136 != v55 {
		v114 = v114 + v130
		v125 = v136
		goto L21
	} else {
		goto L23
	}
L22:
	;
	goto L12
L23:
	;
	goto L22
L24:
	;
	v616 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v616)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v604)+988))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v618
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v604)+984))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v620
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v604)+976))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v622
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v604)+980))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v624
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v604)+992))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v626
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+1008)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)) = uint8(v628)
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v604)+996))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v604)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v604)+1004))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v634
	F_pfree(m, v13)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L125
	}
L25:
	;
	v584 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v572)+1008)) = v584
	v586 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v572)+1000)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v572)+992)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v572)+984)) = v586
	*(*int64)(unsafe.Add(mBase, uint32(v572)+976)) = v586
	F_parse_format(m, v572, v13, int32(_a_F_NUM_cache_0), v584, int32(_a_F_NUM_cache_1), int32(2), v572+int32(976))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L124
	}
L26:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+969)) = uint8(v442)
	v445 = v430 + int32(912)
	goto L96
L27:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[6]))
	v285 = F_MemoryContextAllocZero(m, v283, int32(1012))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L61
	}
L28:
	;
	v170 = int32(0)
	goto L30
L29:
	;
	v267 = v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v182)+972)) = v267
	v604 = v182
	goto L24
L30:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v170<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+969)))
	if v183 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v47 < int32(20) {
		goto L27
	} else {
		goto L44
	}
L32:
	;
	v187 = v182 + int32(912)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v190 == int32(0))|base.B2i32(v190 != v193) != 0 {
		v211 = v190
		v212 = v193
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v217 = v170 + int32(1)
	if v217 != v47 {
		v170 = v217
		goto L30
	} else {
		goto L43
	}
L35:
	;
	if v211-v212 == int32(0) {
		goto L29
	} else {
		goto L42
	}
L36:
	;
	goto L35
L37:
	;
	v196 = v187
	v197 = v13
	goto L38
L38:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v201 == int32(0) {
		v211 = v201
		v212 = v200
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v211 = v201
	v212 = v200
	goto L36
L40:
	;
	v204 = int32(1)
	if v201 == v200 {
		v196 = v196 + v204
		v197 = v197 + v204
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L34
L43:
	;
	goto L31
L44:
	;
	v221 = int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[2]))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+969)))
	if v224 != v221 {
		v430 = v223
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v227 = v223
	v229 = v221
	goto L46
L46:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v229<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+969)))
	if v242 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v430 = v241
	goto L26
L49:
	;
	goto L50
L50:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v241)+972))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227)+972))
	if v245 < v246 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v248 = v241
	goto L53
L52:
	;
	v248 = v227
	goto L53
L53:
	;
	v250 = v229 + int32(1)
	if v250 == int32(20) {
		v430 = v248
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+969)))
	if v256 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v430 = v255
	goto L26
L56:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+972))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v248)+972))
	if v259 < v260 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v262 = v255
	goto L60
L59:
	;
	v262 = v248
	goto L60
L60:
	;
	v227 = v262
	v229 = v229 + int32(2)
	goto L46
L61:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v288<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2]))) = v285
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v285)+969)) = uint8(v294)
	v297 = v285 + int32(912)
	goto L65
L62:
	;
	v417 = int32(_a_F_NUM_cache_2)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	v420 = int32(1)
	v421 = v419 + v420
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v421
	*(*int32)(unsafe.Add(mBase, uint32(v285)+972)) = v421
	v424 = int32(_a_F_NUM_cache_3)
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0])) = v426 + v420
	v572 = v285
	goto L25
L63:
	;
	v414 = F_strlen(m, v403)
	mBase = m.M
	goto L62
L65:
	;
	goto L66
L66:
	;
	v304 = int32(56)
	if (v297^v13)&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v404))) = uint8(v407)
	goto L63
L68:
	;
	v388 = v383
	v389 = v384
	v390 = v385
	goto L89
L69:
	;
	if v378 == int32(0) {
		v403 = v376
		v404 = v377
		goto L67
	} else {
		goto L88
	}
L70:
	;
	v376 = v13
	v377 = v297
	v378 = v304
	goto L69
L71:
	;
	goto L72
L72:
	;
	v308 = int32(0)
	if base.B2i32(v13&int32(3) == v308)|int32(0) == v308 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v344 == int32(0) {
		v403 = v341
		v404 = v342
		goto L67
	} else {
		goto L82
	}
L74:
	;
	v320 = v13
	v321 = v297
	v322 = v304
	goto L77
L75:
	;
	goto L76
L76:
	;
	v341 = v13
	v342 = v297
	v343 = v304
	v344 = int32(1)
	goto L73
L77:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	*(*uint8)(unsafe.Add(mBase, uint32(v321))) = uint8(v324)
	if v324 == int32(0) {
		v383 = v320
		v384 = v321
		v385 = v322
		goto L68
	} else {
		goto L79
	}
L78:
	;
	v341 = v335
	v342 = v329
	v343 = v331
	v344 = v333
	goto L73
L79:
	;
	v328 = int32(1)
	v329 = v321 + v328
	v331 = v322 - v328
	v332 = int32(0)
	v333 = base.B2i32(v331 != v332)
	v335 = v320 + v328
	if v335&int32(3) == v332 {
		v341 = v335
		v342 = v329
		v343 = v331
		v344 = v333
		goto L73
	} else {
		goto L80
	}
L80:
	;
	if v331 != 0 {
		v320 = v335
		v321 = v329
		v322 = v331
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
	if base.B2i32(v347 == int32(0))|base.B2i32(base.Ui32(v343) < base.Ui32(int32(4))) != 0 {
		v376 = v341
		v377 = v342
		v378 = v343
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v354 = v341
	v355 = v342
	v356 = v343
	goto L84
L84:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v354)))
	v362 = int32(-2139062144)
	if (int32(16843008)-v359|v359)&v362 != v362 {
		v383 = v354
		v384 = v355
		v385 = v356
		goto L68
	} else {
		goto L86
	}
L85:
	;
	v376 = v370
	v377 = v368
	v378 = v372
	goto L69
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v359
	v367 = int32(4)
	v368 = v355 + v367
	v370 = v354 + v367
	v372 = v356 - v367
	if base.Ui32(int32(3)) < base.Ui32(v372) {
		v354 = v370
		v355 = v368
		v356 = v372
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v383 = v376
	v384 = v377
	v385 = v378
	goto L68
L89:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v388))))
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v392)
	if v392 == int32(0) {
		v403 = v388
		v404 = v389
		goto L67
	} else {
		goto L91
	}
L90:
	;
	v403 = v399
	v404 = v397
	goto L67
L91:
	;
	v396 = int32(1)
	v397 = v389 + v396
	v399 = v388 + v396
	v401 = v390 - v396
	if v401 != 0 {
		v388 = v399
		v389 = v397
		v390 = v401
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v565 = int32(_a_F_NUM_cache_2)
	v567 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	v569 = v567 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v430)+972)) = v569
	v572 = v430
	goto L25
L94:
	;
	v562 = F_strlen(m, v551)
	mBase = m.M
	goto L93
L96:
	;
	goto L97
L97:
	;
	v452 = int32(56)
	if (v445^v13)&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v555)
	goto L94
L99:
	;
	v536 = v531
	v537 = v532
	v538 = v533
	goto L120
L100:
	;
	if v526 == int32(0) {
		v551 = v524
		v552 = v525
		goto L98
	} else {
		goto L119
	}
L101:
	;
	v524 = v13
	v525 = v445
	v526 = v452
	goto L100
L102:
	;
	goto L103
L103:
	;
	v456 = int32(0)
	if base.B2i32(v13&int32(3) == v456)|int32(0) == v456 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v492 == int32(0) {
		v551 = v489
		v552 = v490
		goto L98
	} else {
		goto L113
	}
L105:
	;
	v468 = v13
	v469 = v445
	v470 = v452
	goto L108
L106:
	;
	goto L107
L107:
	;
	v489 = v13
	v490 = v445
	v491 = v452
	v492 = int32(1)
	goto L104
L108:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	*(*uint8)(unsafe.Add(mBase, uint32(v469))) = uint8(v472)
	if v472 == int32(0) {
		v531 = v468
		v532 = v469
		v533 = v470
		goto L99
	} else {
		goto L110
	}
L109:
	;
	v489 = v483
	v490 = v477
	v491 = v479
	v492 = v481
	goto L104
L110:
	;
	v476 = int32(1)
	v477 = v469 + v476
	v479 = v470 - v476
	v480 = int32(0)
	v481 = base.B2i32(v479 != v480)
	v483 = v468 + v476
	if v483&int32(3) == v480 {
		v489 = v483
		v490 = v477
		v491 = v479
		v492 = v481
		goto L104
	} else {
		goto L111
	}
L111:
	;
	if v479 != 0 {
		v468 = v483
		v469 = v477
		v470 = v479
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v489))))
	if base.B2i32(v495 == int32(0))|base.B2i32(base.Ui32(v491) < base.Ui32(int32(4))) != 0 {
		v524 = v489
		v525 = v490
		v526 = v491
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v502 = v489
	v503 = v490
	v504 = v491
	goto L115
L115:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v510 = int32(-2139062144)
	if (int32(16843008)-v507|v507)&v510 != v510 {
		v531 = v502
		v532 = v503
		v533 = v504
		goto L99
	} else {
		goto L117
	}
L116:
	;
	v524 = v518
	v525 = v516
	v526 = v520
	goto L100
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v503))) = v507
	v515 = int32(4)
	v516 = v503 + v515
	v518 = v502 + v515
	v520 = v504 - v515
	if base.Ui32(int32(3)) < base.Ui32(v520) {
		v502 = v518
		v503 = v516
		v504 = v520
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v531 = v524
	v532 = v525
	v533 = v526
	goto L99
L120:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536))))
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v540)
	if v540 == int32(0) {
		v551 = v536
		v552 = v537
		goto L98
	} else {
		goto L122
	}
L121:
	;
	v551 = v547
	v552 = v545
	goto L98
L122:
	;
	v544 = int32(1)
	v545 = v537 + v544
	v547 = v536 + v544
	v549 = v538 - v544
	if v549 != 0 {
		v536 = v547
		v537 = v545
		v538 = v549
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v602 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v572)+969)) = uint8(v602)
	v604 = v572
	goto L24
L125:
	;
	return v604
}
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v244 int32
	_ = v244
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v817 int32
	_ = v817
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1397 int32
	_ = v1397
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1443 int64
	_ = v1443
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1508 int32
	_ = v1508
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1632 int32
	_ = v1632
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1685 int32
	_ = v1685
	var v1694 int32
	_ = v1694
	var v1705 int32
	_ = v1705
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	var v1770 int32
	_ = v1770
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1785 int32
	_ = v1785
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1801 int32
	_ = v1801
	var v1805 int32
	_ = v1805
	var v1811 int32
	_ = v1811
	var v1826 int32
	_ = v1826
	var v1831 int32
	_ = v1831
	var v1835 int32
	_ = v1835
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1995 int64
	_ = v1995
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2010 int64
	_ = v2010
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2067 int32
	_ = v2067
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2108 int32
	_ = v2108
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2121 int32
	_ = v2121
	var v2123 int64
	_ = v2123
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2138 int64
	_ = v2138
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2279 int32
	_ = v2279
	var v2304 int32
	_ = v2304
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2396 int32
	_ = v2396
	var v2401 int32
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2417 int32
	_ = v2417
	v8 = l7
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(192)
	m.G0 = v22
	base.MemoryFill(m, v22+int32(28), v9, int32(84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v22)+60)) = int64(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v35 - int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v39&int32(_a_F_NUM_processor_0) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2405 = m.ExcPending
	if v2405 != 0 {
		goto L76
	} else {
		goto L681
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L76
	} else {
		goto L677
	}
L6:
	;
	m.G0 = v22 + int32(192)
	return
L7:
	;
	if v8 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v8 != 0 {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	if (l3^l2)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L6
L12:
	;
	goto L11
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v97)
	if v97&int32(255) == int32(0) {
		goto L12
	} else {
		goto L28
	}
L14:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v96 = l3
	v97 = v49
	v98 = l2
	goto L13
L15:
	;
	goto L16
L16:
	;
	if l3&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = l3
	v55 = l2
	goto L20
L18:
	;
	v67 = l3
	v69 = l2
	goto L19
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v74 = int32(-2139062144)
	if (int32(16843008)-v71|v71)&v74 != v74 {
		v96 = v67
		v97 = v71
		v98 = v69
		goto L13
	} else {
		goto L24
	}
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v56)
	if v56 == int32(0) {
		goto L12
	} else {
		goto L22
	}
L21:
	;
	v67 = v63
	v69 = v61
	goto L19
L22:
	;
	v60 = int32(1)
	v61 = v55 + v60
	v63 = v53 + v60
	if v63&int32(3) != 0 {
		v53 = v63
		v55 = v61
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v79 = v67
	v80 = v71
	v81 = v69
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v80
	v83 = int32(4)
	v84 = v81 + v83
	v86 = v79 + v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v91 = int32(-2139062144)
	if (int32(16843008)-v88|v88)&v91 == v91 {
		v79 = v86
		v80 = v88
		v81 = v84
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v96 = v86
	v97 = v88
	v98 = v84
	goto L13
L27:
	;
	goto L26
L28:
	;
	v105 = v96
	v107 = v98
	goto L29
L29:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)) = uint8(v108)
	v110 = int32(1)
	if v108 != 0 {
		v105 = v105 + v110
		v107 = v107 + v110
		goto L29
	} else {
		goto L31
	}
L30:
	;
	goto L12
L31:
	;
	goto L30
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = int64(0)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v309 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l6
	v120 = v39 & int32(768)
	if v120 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v276 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v276
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v279 + v278 - int32(1)
	v286 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v286)
	goto L32
L36:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = l5
	v162 = v160 + v159
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v162 - int32(1)
	v166 = int32(34)
	if v157&v166 != v166 {
		goto L48
	} else {
		goto L49
	}
L37:
	;
	if v120 == int32(512) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v130 = int32(0)
	if base.B2i32(v39&int32(32) == v130)|base.B2i32(l6 == int32(45)) == v130 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(0)
	v157 = v39
	v158 = v9
	goto L36
L41:
	;
	goto L42
L42:
	;
	v125 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v125
	v157 = v39
	v158 = v125
	goto L36
L43:
	;
	v138 = v39 & int32(-17281)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v138
	v140 = v138
	goto L45
L44:
	;
	v140 = v39
	goto L45
L45:
	;
	v147 = base.B2i32(l6 == int32(43)) & base.B2i32(v140&int32(96) == int32(32))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v149 != int32(-1) {
		v157 = v140
		v158 = v147
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v152 != v153 {
		v157 = v140
		v158 = v147
		goto L36
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	v157 = v140
	v158 = v147
	goto L36
L48:
	;
	if l5|v158 != 0 {
		goto L32
	} else {
		goto L69
	}
L49:
	;
	v170 = int32(46)
	v171 = F___strchrnul(m, l3, v170)
	mBase = m.M
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v173 == v170 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v244
	goto L48
L51:
	;
	if v177 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v177 = v171
	goto L54
L53:
	;
	v177 = int32(0)
	goto L54
L54:
	;
	goto L51
L55:
	;
	v244 = int32(0)
	goto L50
L56:
	;
	goto L57
L57:
	;
	v190 = v177
	goto L58
L58:
	;
	v209 = v190
	goto L60
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+88)) = v190
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v225 <= l5 {
		goto L48
	} else {
		goto L64
	}
L60:
	;
	v220 = v209 + int32(1)
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v221 == int32(48) {
		v209 = v220
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if v221 != 0 {
		v190 = v220
		goto L58
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	goto L59
L64:
	;
	v227 = F_strlen(m, l3)
	mBase = m.M
	v229 = v227 - int32(1)
	v230 = v225 - l5
	if base.Ui32(v229) < base.Ui32(v230) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v232 = v229
	goto L67
L66:
	;
	v232 = v230
	goto L67
L67:
	;
	v233 = l3 + v232
	if base.Ui32(v233) <= base.Ui32(v190) {
		goto L48
	} else {
		goto L68
	}
L68:
	;
	v244 = v233
	goto L50
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v162
	goto L32
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+108)) = v365
	v371 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = l3 + (v8 ^ v371)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v375 == v371 {
		goto L106
	} else {
		goto L107
	}
L71:
	;
	v365 = int32(_a_F_NUM_processor_1)
	goto L70
L72:
	;
	v312 = F_PGLC_localeconv(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = int32(_a_F_NUM_processor_2)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = int32(_a_F_NUM_processor_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(_a_F_NUM_processor_4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+92)) = int32(_a_F_NUM_processor_5)
	goto L71
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+92)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v312)+32))
	if v319 != 0 {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	return
L77:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v312)+36))
	if v314 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v315 != 0 {
		v317 = v314
		goto L75
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v317 = int32(_a_F_NUM_processor_5)
	goto L75
L81:
	;
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v324 != 0 {
		goto L88
	} else {
		goto L89
	}
L83:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319))))
	if v320 != 0 {
		v322 = v319
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v322 = int32(_a_F_NUM_processor_4)
	goto L82
L86:
	;
	goto L85
L87:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v329&int32(4) != 0 {
		goto L92
	} else {
		goto L93
	}
L88:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324))))
	if v325 != 0 {
		v327 = v324
		goto L87
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v327 = int32(_a_F_NUM_processor_3)
	goto L87
L91:
	;
	goto L90
L92:
	;
	v332 = v327
	goto L94
L93:
	;
	v332 = int32(_a_F_NUM_processor_3)
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	if v334 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+104)) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	if v346 == int32(0) {
		goto L71
	} else {
		goto L104
	}
L96:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v335 != 0 {
		v344 = v334
		goto L95
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
	if v337 != int32(44) {
		v344 = int32(_a_F_NUM_processor_2)
		goto L95
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
	if v342 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v343 = int32(_a_F_NUM_processor_2)
	goto L103
L102:
	;
	v343 = int32(_a_F_NUM_processor_3)
	goto L103
L103:
	;
	v344 = v343
	goto L95
L104:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v349 == int32(0) {
		goto L71
	} else {
		goto L105
	}
L105:
	;
	v365 = v346
	goto L70
L106:
	;
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
	if v2326 == int32(1) {
		goto L670
	} else {
		goto L671
	}
L107:
	;
	v378 = l0
	v379 = v375
	goto L108
L108:
	;
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
	if v397 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L109:
	;
	goto L106
L110:
	;
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v378)+16))
	if v2304 != int32(1) {
		v378 = v378 + int32(16)
		v379 = v2304
		goto L108
	} else {
		goto L669
	}
L111:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+8))
	switch v495 {
	case 0:
		goto L151
	case 1, 2, 3, 6:
		goto L152
	default:
		goto L110
	case 9:
		goto L150
	case 10:
		goto L149
	case 11:
		goto L145
	case 12:
		goto L144
	case 14, 30:
		goto L148
	case 15:
		goto L143
	case 18:
		goto L146
	case 34:
		goto L147
	}
L112:
	;
	v488 = F_pg_mblen_range(m, v400, v402)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L76
	} else {
		goto L140
	}
L113:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v402 = v401 + l4
	if base.Ui32(v402) <= base.Ui32(v400) {
		goto L106
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v379 == int32(2) {
		goto L111
	} else {
		goto L118
	}
L116:
	;
	if v379 != int32(2) {
		goto L112
	} else {
		goto L117
	}
L117:
	;
	goto L111
L118:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v410 = v378 + int32(4)
	if (v410^v408)&int32(3) != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v485 = F_strlen(m, v408)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v485 + v408
	goto L110
L120:
	;
	goto L119
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v464)
	if v464&int32(255) == int32(0) {
		goto L120
	} else {
		goto L136
	}
L122:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v463 = v410
	v464 = v416
	v465 = v408
	goto L121
L123:
	;
	goto L124
L124:
	;
	if v410&int32(3) != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v420 = v410
	v422 = v408
	goto L128
L126:
	;
	v434 = v410
	v436 = v408
	goto L127
L127:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434)))
	v441 = int32(-2139062144)
	if (int32(16843008)-v438|v438)&v441 != v441 {
		v463 = v434
		v464 = v438
		v465 = v436
		goto L121
	} else {
		goto L132
	}
L128:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420))))
	*(*uint8)(unsafe.Add(mBase, uint32(v422))) = uint8(v423)
	if v423 == int32(0) {
		goto L120
	} else {
		goto L130
	}
L129:
	;
	v434 = v430
	v436 = v428
	goto L127
L130:
	;
	v427 = int32(1)
	v428 = v422 + v427
	v430 = v420 + v427
	if v430&int32(3) != 0 {
		v420 = v430
		v422 = v428
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v446 = v434
	v447 = v438
	v448 = v436
	goto L133
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v448))) = v447
	v450 = int32(4)
	v451 = v448 + v450
	v453 = v446 + v450
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	v458 = int32(-2139062144)
	if (int32(16843008)-v455|v455)&v458 == v458 {
		v446 = v453
		v447 = v455
		v448 = v451
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v463 = v453
	v464 = v455
	v465 = v451
	goto L121
L135:
	;
	goto L134
L136:
	;
	v472 = v463
	v474 = v465
	goto L137
L137:
	;
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v472)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)) = uint8(v475)
	v477 = int32(1)
	if v475 != 0 {
		v472 = v472 + v477
		v474 = v474 + v477
		goto L137
	} else {
		goto L139
	}
L138:
	;
	goto L120
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v488 + v400
	goto L110
L141:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2279 + int32(1)
	goto L110
L142:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1123))) = uint8(v1285)
	goto L141
L143:
	;
	if v397 != 0 {
		goto L661
	} else {
		goto L662
	}
L144:
	;
	if v397 != 0 {
		goto L649
	} else {
		goto L650
	}
L145:
	;
	if v397 != 0 {
		goto L637
	} else {
		goto L638
	}
L146:
	;
	v2024 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2024)+12))
	if v2025&int32(1024)|v2025&int32(2) != 0 {
		goto L110
	} else {
		goto L603
	}
L147:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1896)+12))
	if v1897&int32(1024)|v1897&int32(2) != 0 {
		goto L110
	} else {
		goto L569
	}
L148:
	;
	if v397 != 0 {
		goto L458
	} else {
		goto L459
	}
L149:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v22)+108))
	if v397 != 0 {
		goto L440
	} else {
		goto L441
	}
L150:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v22)+104))
	v1312 = F_strlen(m, v1311)
	mBase = m.M
	if v397 != 0 {
		goto L403
	} else {
		goto L404
	}
L151:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v397 != 0 {
		goto L391
	} else {
		goto L392
	}
L152:
	;
	if v397 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496)+13)))
	if v497&int32(4) != 0 {
		goto L110
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v831 = v830 + l4
	if base.Ui32(v831) <= base.Ui32(v829) {
		goto L141
	} else {
		goto L272
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(0)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if v502 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	if base.B2i32(int32(1)<<(uint(v495)%32)&int32(78) == int32(0))|base.B2i32(base.Ui32(int32(6)) < base.Ui32(v495)) != 0 {
		goto L196
	} else {
		goto L197
	}
L158:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v496)+12))
	v505 = v503 & int32(8)
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	if v506 < v507 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	if v503&int32(64) != 0 {
		goto L173
	} else {
		goto L174
	}
L160:
	;
	if v505 == int32(0) {
		goto L157
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if v505 != 0 {
		goto L159
	} else {
		goto L165
	}
L163:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v496)+24))
	if v511 == v506 {
		goto L159
	} else {
		goto L164
	}
L164:
	;
	goto L157
L165:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	if v513 != v514 {
		goto L159
	} else {
		goto L166
	}
L166:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513))))
	if v516 != int32(48) {
		goto L159
	} else {
		goto L167
	}
L167:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v496)+4))
	if v519 == int32(0) {
		goto L159
	} else {
		goto L168
	}
L168:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	if v522 == int32(0) {
		goto L157
	} else {
		goto L169
	}
L169:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522))))
	if v525 != int32(46) {
		goto L157
	} else {
		goto L170
	}
L170:
	;
	goto L159
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(1)
	goto L157
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v586
	goto L171
L173:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v496)+8))
	if v531 != int32(-1) {
		goto L157
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v503&int32(128) != 0 {
		goto L187
	} else {
		goto L188
	}
L176:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v538 == int32(45) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v541 = int32(64)
	goto L179
L178:
	;
	v541 = int32(68)
	goto L179
L179:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(28)+v541)))
	v544 = F_strlen(m, v543)
	mBase = m.M
	if base.Ui32(int32(9)) <= base.Ui32(v544) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v548 = F_pg_mbcliplen(m, v543, v544, int32(8))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L76
	} else {
		goto L183
	}
L181:
	;
	v550 = v544
	goto L182
L182:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v550 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v550 = v548
	goto L182
L184:
	;
	base.MemoryCopy(m, v551, v543, v550)
	goto L186
L185:
	;
	goto L186
L186:
	;
	v586 = v550 + v551
	goto L172
L187:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v554 == int32(43) {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	goto L189
L189:
	;
	switch v554 - int32(43) {
	case 0:
		goto L194
	default:
		goto L157
	case 2:
		goto L193
	}
L190:
	;
	v562 = int32(32)
	goto L192
L191:
	;
	v562 = int32(60)
	goto L192
L192:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v557))) = uint8(v562)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v586 = v564 + int32(1)
	goto L172
L193:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v578 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v577))) = uint8(v578)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v586 = v580 + int32(1)
	goto L172
L194:
	;
	if v503&int32(32) != 0 {
		goto L171
	} else {
		goto L195
	}
L195:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v572 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v571))) = uint8(v572)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v586 = v574 + int32(1)
	goto L172
L196:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v825 + int32(1)
	goto L110
L197:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	if v606 < v607 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v743 = int32(0)
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+12))
	v748 = int32(1)
	v752 = v741 + base.B2i32(v742 != v743) + int32(base.Ui32(v747)>>(uint(v748)%32))&v748
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	if v753 == v743 {
		goto L246
	} else {
		goto L247
	}
L199:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v605)+24))
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	if v612&int32(8) != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v605)+12))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638))))
	if v639 == int32(46) {
		goto L210
	} else {
		goto L211
	}
L202:
	;
	v615 = base.B2i32(v609 <= v606)
	goto L204
L203:
	;
	v615 = int32(0)
	goto L204
L204:
	;
	if v615 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	if v612&int32(32) != 0 {
		goto L198
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v628 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v628)
	v630 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v632 + v630
	goto L198
L208:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v621 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v620))) = uint8(v621)
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v623 + int32(1)
	goto L198
L209:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v730))))
	if v731 == int32(0) {
		goto L198
	} else {
		goto L244
	}
L210:
	;
	if v637 != 0 {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L212
L212:
	;
	v676 = int32(0)
	if base.B2i32(base.B2i32(v637 == v676)|base.B2i32(v495 == int32(2)) == v676)&base.B2i32(base.Ui32(v637) < base.Ui32(v638)) != 0 {
		goto L209
	} else {
		goto L233
	}
L213:
	;
	if v636&int32(32) == int32(0) {
		goto L209
	} else {
		goto L225
	}
L214:
	;
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v642 == int32(46) {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v646 = F_strlen(m, v645)
	mBase = m.M
	if base.Ui32(int32(9)) <= base.Ui32(v646) {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	goto L216
L218:
	;
	v650 = F_pg_mbcliplen(m, v645, v646, int32(8))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L76
	} else {
		goto L221
	}
L219:
	;
	v652 = v646
	goto L220
L220:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v652 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v652 = v650
	goto L220
L222:
	;
	base.MemoryCopy(m, v653, v645, v652)
	goto L224
L223:
	;
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v652 + v653
	goto L209
L225:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v662 = F_strlen(m, v661)
	mBase = m.M
	if base.Ui32(int32(9)) <= base.Ui32(v662) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v666 = F_pg_mbcliplen(m, v661, v662, int32(8))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L76
	} else {
		goto L229
	}
L227:
	;
	v668 = v662
	goto L228
L228:
	;
	v670 = v22 + int32(28)
	if v668 != 0 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v668 = v666
	goto L228
L230:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)+56))
	base.MemoryCopy(m, v671, v661, v668)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v670)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+56)) = v673 + v668
	goto L209
L233:
	;
	if v636&int32(8) != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*uint8)(unsafe.Add(mBase, uint32(v719))) = uint8(v639)
	v721 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v721
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v723 + v721
	goto L209
L235:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	if v638 != v687 {
		goto L234
	} else {
		goto L236
	}
L236:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687))))
	if v689 != int32(48) {
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v605)+4))
	if v692 == int32(0) {
		goto L234
	} else {
		goto L238
	}
L238:
	;
	if v636&int32(32) == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v700 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v699))) = uint8(v700)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v702 + int32(1)
	goto L209
L240:
	;
	goto L241
L241:
	;
	if v637 == int32(0) {
		goto L209
	} else {
		goto L242
	}
L242:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637))))
	if v708 != int32(46) {
		goto L209
	} else {
		goto L243
	}
L243:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v712 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v712)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v714 + int32(1)
	goto L209
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v730 + int32(1)
	goto L198
L245:
	;
	if v762+int32(1) != v761 {
		goto L196
	} else {
		goto L252
	}
L246:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v761 = v752
	v762 = v756
	goto L245
L247:
	;
	goto L248
L248:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	if v753 == v758 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v760 = v757
	goto L251
L250:
	;
	v760 = v752
	goto L251
L251:
	;
	v761 = v760
	v762 = v757
	goto L245
L252:
	;
	v768 = int32(0)
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	if base.B2i32(v747&int32(128) == v768)|base.B2i32(v770 != int32(1)) == v768 {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v817
	goto L196
L254:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v779 == int32(43) {
		goto L257
	} else {
		goto L258
	}
L255:
	;
	goto L256
L256:
	;
	if v747&int32(64) == int32(0) {
		goto L196
	} else {
		goto L260
	}
L257:
	;
	v782 = int32(32)
	goto L259
L258:
	;
	v782 = int32(62)
	goto L259
L259:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v776))) = uint8(v782)
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v817 = v784 + int32(1)
	goto L253
L260:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v746)+8))
	if v791 != int32(1) {
		goto L196
	} else {
		goto L261
	}
L261:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v798 == int32(45) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v801 = int32(64)
	goto L264
L263:
	;
	v801 = int32(68)
	goto L264
L264:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(28)+v801)))
	v804 = F_strlen(m, v803)
	mBase = m.M
	if base.Ui32(int32(9)) <= base.Ui32(v804) {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v808 = F_pg_mbcliplen(m, v803, v804, int32(8))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L76
	} else {
		goto L268
	}
L266:
	;
	v810 = v804
	goto L267
L267:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v810 != 0 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v810 = v808
	goto L267
L269:
	;
	base.MemoryCopy(m, v811, v803, v810)
	goto L271
L270:
	;
	goto L271
L271:
	;
	v817 = v810 + v811
	goto L253
L272:
	;
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829))))
	if v833 == int32(32) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v837 = v829 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v837
	v839 = v837
	goto L275
L274:
	;
	v839 = v829
	goto L275
L275:
	;
	if base.Ui32(v831) <= base.Ui32(v839) {
		goto L141
	} else {
		goto L276
	}
L276:
	;
	if v495&int32(-2) != int32(2) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	if base.Ui32(v1009+l4) <= base.Ui32(v1008) {
		goto L141
	} else {
		goto L319
	}
L278:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	if v846 != int32(32) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	if v849 != int32(0)-v851 {
		goto L277
	} else {
		goto L280
	}
L280:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v854)+12))
	if v855&int32(64) == int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	v982 = int32(0)
	if base.B2i32(v977 != int32(45))&(base.B2i32(v855&int32(128) == v982)|base.B2i32(v977 != int32(60))) == v982 {
		goto L315
	} else {
		goto L316
	}
L282:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v854)+8))
	if v860 != int32(-1) {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	v864 = F_strlen(m, v863)
	mBase = m.M
	if base.B2i32(v864 == int32(0))|base.B2i32(base.Ui32(v830+(l4-v864)) < base.Ui32(v839)) != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v921 = F_strlen(m, v920)
	mBase = m.M
	if base.B2i32(v921 == int32(0))|base.B2i32(base.Ui32(v830+(l4-v921)) < base.Ui32(v839)) != 0 {
		goto L277
	} else {
		goto L300
	}
L285:
	;
	if v864 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	if v915 != 0 {
		goto L284
	} else {
		goto L299
	}
L287:
	;
	v915 = int32(0)
	goto L286
L288:
	;
	goto L289
L289:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if v876 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v877 = v839
	v878 = v863
	v879 = v864
	v880 = v876
	goto L294
L291:
	;
	v903 = v863
	v907 = int32(0)
	goto L292
L292:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	v915 = v907 - v908
	goto L286
L293:
	;
	v903 = v898
	v907 = v900
	goto L292
L294:
	;
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	if base.B2i32(v880 != v882)|base.B2i32(v882 == int32(0)) != 0 {
		v898 = v878
		v900 = v880
		goto L293
	} else {
		goto L296
	}
L295:
	;
	v898 = v892
	v900 = int32(0)
	goto L293
L296:
	;
	v888 = v879 - int32(1)
	if v888 == int32(0) {
		v898 = v878
		v900 = v880
		goto L293
	} else {
		goto L297
	}
L297:
	;
	v891 = int32(1)
	v892 = v878 + v891
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877)+1)))
	if v893 != 0 {
		v877 = v877 + v891
		v878 = v892
		v879 = v888
		v880 = v893
		goto L294
	} else {
		goto L298
	}
L298:
	;
	goto L295
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v839 + v864
	v918 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v845))) = uint8(v918)
	goto L277
L300:
	;
	if v921 == int32(0) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	if v972 != 0 {
		goto L277
	} else {
		goto L314
	}
L302:
	;
	v972 = int32(0)
	goto L301
L303:
	;
	goto L304
L304:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if v933 != 0 {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v934 = v839
	v935 = v920
	v936 = v921
	v937 = v933
	goto L309
L306:
	;
	v960 = v920
	v964 = int32(0)
	goto L307
L307:
	;
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960))))
	v972 = v964 - v965
	goto L301
L308:
	;
	v960 = v955
	v964 = v957
	goto L307
L309:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	if base.B2i32(v937 != v939)|base.B2i32(v939 == int32(0)) != 0 {
		v955 = v935
		v957 = v937
		goto L308
	} else {
		goto L311
	}
L310:
	;
	v955 = v949
	v957 = int32(0)
	goto L308
L311:
	;
	v945 = v936 - int32(1)
	if v945 == int32(0) {
		v955 = v935
		v957 = v937
		goto L308
	} else {
		goto L312
	}
L312:
	;
	v948 = int32(1)
	v949 = v935 + v948
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+1)))
	if v950 != 0 {
		v934 = v934 + v948
		v935 = v949
		v936 = v945
		v937 = v950
		goto L309
	} else {
		goto L313
	}
L313:
	;
	goto L310
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v839 + v921
	v975 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v845))) = uint8(v975)
	goto L277
L315:
	;
	v990 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v845))) = uint8(v990)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v992 + int32(1)
	goto L277
L316:
	;
	goto L317
L317:
	;
	if v977 != int32(43) {
		goto L277
	} else {
		goto L318
	}
L318:
	;
	v998 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v845))) = uint8(v998)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1000 + int32(1)
	goto L277
L319:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	if base.Ui32((v1012-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v1121 = v1120 + l4
	if base.Ui32(v1121) <= base.Ui32(v1119) {
		goto L141
	} else {
		goto L348
	}
L321:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	if v1019 != 0 {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	goto L323
L323:
	;
	v1041 = int32(0)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1042)+12)))
	if v1043&int32(2) == v1041 {
		v1118 = v1041
		goto L320
	} else {
		goto L331
	}
L324:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	if v1020 == v1022 {
		goto L141
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	*(*uint8)(unsafe.Add(mBase, uint32(v1024))) = uint8(v1012)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v1026 + int32(1)
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	if v1030 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	goto L326
L328:
	;
	v1031 = int32(1)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1032 + v1031
	v1118 = v1031
	goto L320
L329:
	;
	goto L330
L330:
	;
	v1036 = int32(1)
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v1037 + v1036
	v1118 = v1036
	goto L320
L331:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	if v1048 != 0 {
		v1118 = v1041
		goto L320
	} else {
		goto L332
	}
L332:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v1050 = F_strlen(m, v1049)
	mBase = m.M
	if base.B2i32(v1050 == int32(0))|base.B2i32(base.Ui32(v1009+(l4-v1050)) < base.Ui32(v1008)) != 0 {
		v1118 = v1041
		goto L320
	} else {
		goto L333
	}
L333:
	;
	if v1050 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	if v1101 != 0 {
		v1118 = v1041
		goto L320
	} else {
		goto L347
	}
L335:
	;
	v1101 = int32(0)
	goto L334
L336:
	;
	goto L337
L337:
	;
	v1062 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	if v1062 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1063 = v1008
	v1064 = v1049
	v1065 = v1050
	v1066 = v1062
	goto L342
L339:
	;
	v1089 = v1049
	v1093 = int32(0)
	goto L340
L340:
	;
	v1094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089))))
	v1101 = v1093 - v1094
	goto L334
L341:
	;
	v1089 = v1084
	v1093 = v1086
	goto L340
L342:
	;
	v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1064))))
	if base.B2i32(v1066 != v1068)|base.B2i32(v1068 == int32(0)) != 0 {
		v1084 = v1064
		v1086 = v1066
		goto L341
	} else {
		goto L344
	}
L343:
	;
	v1084 = v1078
	v1086 = int32(0)
	goto L341
L344:
	;
	v1074 = v1065 - int32(1)
	if v1074 == int32(0) {
		v1084 = v1064
		v1086 = v1066
		goto L341
	} else {
		goto L345
	}
L345:
	;
	v1077 = int32(1)
	v1078 = v1064 + v1077
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063)+1)))
	if v1079 != 0 {
		v1063 = v1063 + v1077
		v1064 = v1078
		v1065 = v1074
		v1066 = v1079
		goto L342
	} else {
		goto L346
	}
L346:
	;
	goto L343
L347:
	;
	v1103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1050 + v1008 - v1103
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v1107 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v1106))) = uint8(v1107)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+60)) = v1103
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v1112 + v1103
	v1118 = v1103
	goto L320
L348:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123))))
	if v1124 != int32(32) {
		goto L141
	} else {
		goto L349
	}
L349:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	if v1127+v1128 <= int32(0) {
		goto L141
	} else {
		goto L350
	}
L350:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+12))
	v1135 = v1133 & int32(64)
	v1136 = int32(0)
	if base.B2i32(v1135 == v1136)|(v1118^int32(1)) == v1136 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1144 = v1119 + int32(1)
	if base.Ui32(v1121) <= base.Ui32(v1144) {
		goto L141
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	if v1118|(v1135|base.B2i32(v1133&int32(768) == int32(0))) != 0 {
		goto L141
	} else {
		goto L390
	}
L354:
	;
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144))))
	if base.Ui32((v1146-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L141
	} else {
		goto L355
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1144
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	v1155 = F_strlen(m, v1154)
	mBase = m.M
	if base.B2i32(v1155 == int32(0))|base.B2i32(base.Ui32(v1120+(l4-v1155)) < base.Ui32(v1144)) != 0 {
		goto L358
	} else {
		goto L359
	}
L356:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1274))))
	if v1275 != int32(32) {
		goto L141
	} else {
		goto L389
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1267
	*(*uint8)(unsafe.Add(mBase, uint32(v1123))) = uint8(v1268)
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v1274 = v1271
	goto L356
L358:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v1210 = F_strlen(m, v1209)
	mBase = m.M
	if base.B2i32(v1210 == int32(0))|base.B2i32(base.Ui32(v1120+(l4-v1210)) < base.Ui32(v1144)) != 0 {
		v1274 = v1123
		goto L356
	} else {
		goto L374
	}
L359:
	;
	if v1155 == int32(0) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	if v1206 != 0 {
		goto L358
	} else {
		goto L373
	}
L361:
	;
	v1206 = int32(0)
	goto L360
L362:
	;
	goto L363
L363:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144))))
	if v1167 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1168 = v1144
	v1169 = v1154
	v1170 = v1155
	v1171 = v1167
	goto L368
L365:
	;
	v1194 = v1154
	v1198 = int32(0)
	goto L366
L366:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1194))))
	v1206 = v1198 - v1199
	goto L360
L367:
	;
	v1194 = v1189
	v1198 = v1191
	goto L366
L368:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169))))
	if base.B2i32(v1171 != v1173)|base.B2i32(v1173 == int32(0)) != 0 {
		v1189 = v1169
		v1191 = v1171
		goto L367
	} else {
		goto L370
	}
L369:
	;
	v1189 = v1183
	v1191 = int32(0)
	goto L367
L370:
	;
	v1179 = v1170 - int32(1)
	if v1179 == int32(0) {
		v1189 = v1169
		v1191 = v1171
		goto L367
	} else {
		goto L371
	}
L371:
	;
	v1182 = int32(1)
	v1183 = v1169 + v1182
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168)+1)))
	if v1184 != 0 {
		v1168 = v1168 + v1182
		v1169 = v1183
		v1170 = v1179
		v1171 = v1184
		goto L368
	} else {
		goto L372
	}
L372:
	;
	goto L369
L373:
	;
	v1267 = v1155 + v1119
	v1268 = int32(45)
	goto L357
L374:
	;
	if v1210 == int32(0) {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	if v1261 != 0 {
		v1274 = v1123
		goto L356
	} else {
		goto L388
	}
L376:
	;
	v1261 = int32(0)
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144))))
	if v1222 != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1223 = v1144
	v1224 = v1209
	v1225 = v1210
	v1226 = v1222
	goto L383
L380:
	;
	v1249 = v1209
	v1253 = int32(0)
	goto L381
L381:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1249))))
	v1261 = v1253 - v1254
	goto L375
L382:
	;
	v1249 = v1244
	v1253 = v1246
	goto L381
L383:
	;
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1224))))
	if base.B2i32(v1226 != v1228)|base.B2i32(v1228 == int32(0)) != 0 {
		v1244 = v1224
		v1246 = v1226
		goto L382
	} else {
		goto L385
	}
L384:
	;
	v1244 = v1238
	v1246 = int32(0)
	goto L382
L385:
	;
	v1234 = v1225 - int32(1)
	if v1234 == int32(0) {
		v1244 = v1224
		v1246 = v1226
		goto L382
	} else {
		goto L386
	}
L386:
	;
	v1237 = int32(1)
	v1238 = v1224 + v1237
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223)+1)))
	if v1239 != 0 {
		v1223 = v1223 + v1237
		v1224 = v1238
		v1225 = v1234
		v1226 = v1239
		goto L383
	} else {
		goto L387
	}
L387:
	;
	goto L384
L388:
	;
	v1267 = v1210 + v1144 - int32(1)
	v1268 = int32(43)
	goto L357
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1119
	goto L141
L390:
	;
	v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119))))
	switch v1285 - int32(43) {
	case 0, 2:
		goto L142
	default:
		goto L141
	}
L391:
	;
	if v1288 == int32(0) {
		goto L394
	} else {
		goto L395
	}
L392:
	;
	goto L393
L393:
	;
	if v1288 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L394:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291)+12)))
	if v1292&int32(32) != 0 {
		goto L110
	} else {
		goto L397
	}
L395:
	;
	goto L396
L396:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1299 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v1298))) = uint8(v1299)
	goto L141
L397:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1296 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1295))) = uint8(v1296)
	goto L141
L398:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303)+12)))
	if v1304&int32(32) != 0 {
		goto L110
	} else {
		goto L401
	}
L399:
	;
	goto L400
L400:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307))))
	if v1308 == int32(44) {
		goto L141
	} else {
		goto L402
	}
L401:
	;
	goto L400
L402:
	;
	goto L110
L403:
	;
	if base.Ui32(int32(9)) <= base.Ui32(v1312) {
		goto L406
	} else {
		goto L407
	}
L404:
	;
	goto L405
L405:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v1341 == int32(0) {
		goto L421
	} else {
		goto L422
	}
L406:
	;
	v1316 = F_pg_mbcliplen(m, v1311, v1312, int32(8))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L76
	} else {
		goto L409
	}
L407:
	;
	v1318 = v1312
	goto L408
L408:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	if v1319 == int32(0) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	v1318 = v1316
	goto L408
L410:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322)+12)))
	if v1323&int32(32) != 0 {
		goto L110
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v1318 != 0 {
		goto L418
	} else {
		goto L419
	}
L413:
	;
	v1326 = F_pg_mbstrlen_with_len(m, v1311, v1318)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L76
	} else {
		goto L414
	}
L414:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v1326 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	base.MemoryFill(m, v1328, int32(32), v1326)
	goto L417
L416:
	;
	goto L417
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1328 + v1326 - int32(1)
	goto L141
L418:
	;
	base.MemoryCopy(m, v1335, v1311, v1318)
	goto L420
L419:
	;
	goto L420
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1335 + v1318 - int32(1)
	goto L141
L421:
	;
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344)+12)))
	if v1345&int32(32) != 0 {
		goto L110
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	if base.Ui32(v1349+(l4-v1312)) < base.Ui32(v1348) {
		goto L110
	} else {
		goto L425
	}
L424:
	;
	goto L423
L425:
	;
	if v1312 == int32(0) {
		goto L427
	} else {
		goto L428
	}
L426:
	;
	if v1397 != 0 {
		goto L110
	} else {
		goto L439
	}
L427:
	;
	v1397 = int32(0)
	goto L426
L428:
	;
	goto L429
L429:
	;
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1348))))
	if v1358 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1359 = v1348
	v1360 = v1311
	v1361 = v1312
	v1362 = v1358
	goto L434
L431:
	;
	v1385 = v1311
	v1389 = int32(0)
	goto L432
L432:
	;
	v1390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1385))))
	v1397 = v1389 - v1390
	goto L426
L433:
	;
	v1385 = v1380
	v1389 = v1382
	goto L432
L434:
	;
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360))))
	if base.B2i32(v1362 != v1364)|base.B2i32(v1364 == int32(0)) != 0 {
		v1380 = v1360
		v1382 = v1362
		goto L433
	} else {
		goto L436
	}
L435:
	;
	v1380 = v1374
	v1382 = int32(0)
	goto L433
L436:
	;
	v1370 = v1361 - int32(1)
	if v1370 == int32(0) {
		v1380 = v1360
		v1382 = v1362
		goto L433
	} else {
		goto L437
	}
L437:
	;
	v1373 = int32(1)
	v1374 = v1360 + v1373
	v1375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+1)))
	if v1375 != 0 {
		v1359 = v1359 + v1373
		v1360 = v1374
		v1361 = v1370
		v1362 = v1375
		goto L434
	} else {
		goto L438
	}
L438:
	;
	goto L435
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1348 + v1312 - int32(1)
	goto L141
L440:
	;
	v1403 = F_strlen(m, v1402)
	mBase = m.M
	if base.Ui32(int32(9)) <= base.Ui32(v1403) {
		goto L443
	} else {
		goto L444
	}
L441:
	;
	goto L442
L442:
	;
	v1416 = F_pg_mbstrlen(m, v1402)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L76
	} else {
		goto L450
	}
L443:
	;
	v1407 = F_pg_mbcliplen(m, v1402, v1403, int32(8))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L76
	} else {
		goto L446
	}
L444:
	;
	v1409 = v1403
	goto L445
L445:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	if v1409 != 0 {
		goto L447
	} else {
		goto L448
	}
L446:
	;
	v1409 = v1407
	goto L445
L447:
	;
	base.MemoryCopy(m, v1410, v1402, v1409)
	goto L449
L448:
	;
	goto L449
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1409 + v1410 - int32(1)
	goto L141
L450:
	;
	if v1416 <= int32(0) {
		goto L110
	} else {
		goto L451
	}
L451:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v1421 = v1420 + l4
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1423 = v1422
	v1424 = v1416
	goto L452
L452:
	;
	if base.Ui32(v1421) <= base.Ui32(v1423) {
		goto L110
	} else {
		goto L454
	}
L453:
	;
	goto L110
L454:
	;
	v1443 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1423))))
	if base.B2i32(base.Ui64(v1443) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v1443)%64)&int64(288080842570334209) != int64(0)) != 0 {
		goto L110
	} else {
		goto L455
	}
L455:
	;
	v1453 = F_pg_mblen_range(m, v1423, v1421)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L76
	} else {
		goto L456
	}
L456:
	;
	v1455 = v1453 + v1423
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1455
	v1457 = int32(1)
	if base.Ui32(v1457) < base.Ui32(v1424) {
		v1423 = v1455
		v1424 = v1424 - v1457
		goto L452
	} else {
		goto L457
	}
L457:
	;
	goto L453
L458:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	if v495 != int32(30) {
		v1508 = v1461
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L460
L460:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v1611 = v1610 + l4
	if base.Ui32(v1611) <= base.Ui32(v1609) {
		v1644 = v1609
		goto L497
	} else {
		goto L498
	}
L461:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v1523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1522)+12)))
	if v1523&int32(32) != 0 {
		goto L472
	} else {
		goto L473
	}
L462:
	;
	v1464 = F_strlen(m, v1461)
	mBase = m.M
	v1465 = F_pnstrdup(m, v1461, v1464)
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L76
	} else {
		goto L463
	}
L463:
	;
	v1467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465))))
	if v1467 == int32(0) {
		v1508 = v1465
		goto L461
	} else {
		goto L464
	}
L464:
	;
	v1470 = v1467
	v1471 = v1465
	goto L465
L465:
	;
	if base.Ui32((v1470-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	v1508 = v1465
	goto L461
L467:
	;
	v1497 = v1470 | int32(32)
	goto L469
L468:
	;
	v1497 = v1470
	goto L469
L469:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1471))) = uint8(v1497)
	v1499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471)+1)))
	if v1499 != 0 {
		v1470 = v1499
		v1471 = v1471 + int32(1)
		goto L465
	} else {
		goto L470
	}
L470:
	;
	goto L466
L471:
	;
	v1604 = F_strlen(m, v1521)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1604 + v1521 - int32(1)
	goto L141
L472:
	;
	if (v1508^v1521)&int32(3) != 0 {
		goto L478
	} else {
		goto L479
	}
L473:
	;
	goto L474
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1508
	v1602 = F_pg_sprintf(m, v1521, int32(_a_F_NUM_processor_6), v22)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L76
	} else {
		goto L496
	}
L475:
	;
	goto L471
L476:
	;
	goto L475
L477:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1580))) = uint8(v1579)
	if v1579&int32(255) == int32(0) {
		goto L476
	} else {
		goto L492
	}
L478:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1508))))
	v1578 = v1508
	v1579 = v1531
	v1580 = v1521
	goto L477
L479:
	;
	goto L480
L480:
	;
	if v1508&int32(3) != 0 {
		goto L481
	} else {
		goto L482
	}
L481:
	;
	v1535 = v1508
	v1537 = v1521
	goto L484
L482:
	;
	v1549 = v1508
	v1551 = v1521
	goto L483
L483:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1549)))
	v1556 = int32(-2139062144)
	if (int32(16843008)-v1553|v1553)&v1556 != v1556 {
		v1578 = v1549
		v1579 = v1553
		v1580 = v1551
		goto L477
	} else {
		goto L488
	}
L484:
	;
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1537))) = uint8(v1538)
	if v1538 == int32(0) {
		goto L476
	} else {
		goto L486
	}
L485:
	;
	v1549 = v1545
	v1551 = v1543
	goto L483
L486:
	;
	v1542 = int32(1)
	v1543 = v1537 + v1542
	v1545 = v1535 + v1542
	if v1545&int32(3) != 0 {
		v1535 = v1545
		v1537 = v1543
		goto L484
	} else {
		goto L487
	}
L487:
	;
	goto L485
L488:
	;
	v1561 = v1549
	v1562 = v1553
	v1563 = v1551
	goto L489
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1563))) = v1562
	v1565 = int32(4)
	v1566 = v1563 + v1565
	v1568 = v1561 + v1565
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1561)+4))
	v1573 = int32(-2139062144)
	if (int32(16843008)-v1570|v1570)&v1573 == v1573 {
		v1561 = v1568
		v1562 = v1570
		v1563 = v1566
		goto L489
	} else {
		goto L491
	}
L490:
	;
	v1578 = v1568
	v1579 = v1570
	v1580 = v1566
	goto L477
L491:
	;
	goto L490
L492:
	;
	v1587 = v1578
	v1589 = v1580
	goto L493
L493:
	;
	v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1589)+1)) = uint8(v1590)
	v1592 = int32(1)
	if v1590 != 0 {
		v1587 = v1587 + v1592
		v1589 = v1589 + v1592
		goto L493
	} else {
		goto L495
	}
L494:
	;
	goto L476
L495:
	;
	goto L494
L496:
	;
	goto L471
L497:
	;
	v1664 = v1644
	v1665 = int32(0)
	goto L504
L498:
	;
	v1613 = v1609
	goto L499
L499:
	;
	v1632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v1632-int32(9)))&base.B2i32(v1632 != int32(32)) != 0 {
		v1644 = v1613
		goto L497
	} else {
		goto L501
	}
L500:
	;
	v1644 = v1611
	goto L497
L501:
	;
	v1641 = v1613 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1641
	if v1641 != v1611 {
		v1613 = v1641
		goto L499
	} else {
		goto L502
	}
L502:
	;
	goto L500
L503:
	;
	v1732 = int32(1)
	v1734 = int32(0)
	v1742 = v1734
	v1743 = v1734
	v1744 = v1734
	v1745 = v1734
	v1747 = v1734
	v1753 = v1734
	v1754 = v1732
	v1759 = v1734
	goto L520
L504:
	;
	if base.Ui32(v1611) <= base.Ui32(v1664) {
		goto L506
	} else {
		goto L507
	}
L505:
	;
	if v1665 == int32(0) {
		goto L5
	} else {
		goto L519
	}
L506:
	;
	goto L505
L507:
	;
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1664))))
	if base.Ui32((v1685-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L515
	} else {
		goto L516
	}
L508:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(177)+v1665))) = uint8(v1694)
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(112)+v1665<<(uint(int32(2))%32)))) = v1705
	v1716 = int32(1)
	v1717 = v1664 + v1716
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1717
	v1719 = int32(15)
	v1721 = v1665 + v1716
	if v1721 != v1719 {
		v1664 = v1717
		v1665 = v1721
		goto L504
	} else {
		goto L518
	}
L509:
	;
	v1705 = int32(500)
	goto L508
L510:
	;
	v1705 = int32(100)
	goto L508
L511:
	;
	v1705 = int32(50)
	goto L508
L512:
	;
	v1705 = int32(10)
	goto L508
L513:
	;
	v1705 = int32(5)
	goto L508
L514:
	;
	v1705 = int32(1000)
	goto L508
L515:
	;
	v1694 = v1685 - int32(32)
	goto L517
L516:
	;
	v1694 = v1685
	goto L517
L517:
	;
	switch v1694&int32(255) - int32(67) {
	case 0:
		goto L510
	case 1:
		goto L509
	default:
		goto L506
	case 6:
		v1705 = int32(1)
		goto L508
	case 9:
		goto L511
	case 10:
		goto L514
	case 19:
		goto L513
	case 21:
		goto L512
	}
L518:
	;
	v1731 = v1719
	goto L503
L519:
	;
	v1731 = v1665
	goto L503
L520:
	;
	v1764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+int32(177)+v1742))))
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(112)+v1742<<(uint(int32(2))%32))))
	if int32(4) < v1770 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	if v1876 < int32(0) {
		goto L5
	} else {
		goto L567
	}
L522:
	;
	v1776 = v1744
	goto L524
L523:
	;
	v1776 = int32(0)
	goto L524
L524:
	;
	if int32(49) < v1770 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	v1781 = v1747
	goto L527
L526:
	;
	v1781 = int32(0)
	goto L527
L527:
	;
	if int32(499) < v1770 {
		goto L528
	} else {
		goto L529
	}
L528:
	;
	v1785 = v1753
	goto L530
L529:
	;
	v1785 = int32(0)
	goto L530
L530:
	;
	if v1759&base.B2i32(v1743 <= v1770)|v1776|(v1781|v1785) != 0 {
		goto L5
	} else {
		goto L531
	}
L531:
	;
	switch v1764 - int32(68) {
	case 0:
		goto L533
	default:
		v1796 = v1744
		v1797 = v1747
		v1798 = v1753
		goto L532
	case 8:
		goto L534
	case 18:
		goto L535
	}
L532:
	;
	if base.Ui32(v1742) < base.Ui32(v1731-v1732) {
		goto L538
	} else {
		goto L539
	}
L533:
	;
	v1796 = v1744
	v1797 = v1747
	v1798 = v1753 + int32(1)
	goto L532
L534:
	;
	v1796 = v1744
	v1797 = v1747 + int32(1)
	v1798 = v1753
	goto L532
L535:
	;
	v1796 = v1744 + int32(1)
	v1797 = v1747
	v1798 = v1753
	goto L532
L536:
	;
	v1876 = v1875 + v1745
	v1878 = v1869 + int32(1)
	if base.Ui32(v1878) < base.Ui32(v1731) {
		v1742 = v1878
		v1743 = v1866
		v1744 = v1867
		v1745 = v1876
		v1747 = v1868
		v1753 = v1870
		v1754 = v1871
		v1759 = v1874
		goto L520
	} else {
		goto L566
	}
L537:
	;
	v1866 = v1743
	v1867 = v1796
	v1868 = v1797
	v1869 = v1742
	v1870 = v1798
	v1871 = v1863
	v1874 = v1759
	v1875 = v1770
	goto L536
L538:
	;
	v1801 = v1742 + int32(1)
	v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1801+(v22+int32(177))))))
	v1811 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(112)+v1801<<(uint(int32(2))%32))))
	if v1770 < v1811 {
		goto L541
	} else {
		goto L542
	}
L539:
	;
	v1859 = v1754
	goto L540
L540:
	;
	v1863 = v1859
	goto L537
L541:
	;
	switch v1764 - int32(67) {
	case 0:
		goto L545
	default:
		goto L5
	case 6:
		goto L547
	case 21:
		goto L546
	}
L542:
	;
	goto L543
L543:
	;
	if v1805 != v1764 {
		goto L562
	} else {
		goto L563
	}
L544:
	;
	if int32(4) < v1811 {
		goto L548
	} else {
		goto L549
	}
L545:
	;
	switch v1805 - int32(68) {
	case 0, 9:
		goto L544
	default:
		goto L5
	}
L546:
	;
	switch v1805 - int32(67) {
	case 0, 9:
		goto L544
	default:
		goto L5
	}
L547:
	;
	switch v1805 - int32(86) {
	case 0, 2:
		goto L544
	default:
		goto L5
	}
L548:
	;
	v1826 = v1796
	goto L550
L549:
	;
	v1826 = int32(0)
	goto L550
L550:
	;
	if int32(49) < v1811 {
		goto L551
	} else {
		goto L552
	}
L551:
	;
	v1831 = v1797
	goto L553
L552:
	;
	v1831 = int32(0)
	goto L553
L553:
	;
	if int32(499) < v1811 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v1835 = v1798
	goto L556
L555:
	;
	v1835 = int32(0)
	goto L556
L556:
	;
	if base.B2i32(int32(1) < v1754)|v1826|(v1831|v1835) != 0 {
		goto L5
	} else {
		goto L557
	}
L557:
	;
	switch v1805 - int32(68) {
	case 0:
		goto L559
	default:
		v1846 = v1796
		v1847 = v1797
		v1848 = v1798
		goto L558
	case 8:
		goto L560
	case 18:
		goto L561
	}
L558:
	;
	v1849 = int32(1)
	v1866 = v1770
	v1867 = v1846
	v1868 = v1847
	v1869 = v1801
	v1870 = v1848
	v1871 = v1849
	v1874 = v1849
	v1875 = v1811 - v1770
	goto L536
L559:
	;
	v1846 = v1796
	v1847 = v1797
	v1848 = v1798 + int32(1)
	goto L558
L560:
	;
	v1846 = v1796
	v1847 = v1797 + int32(1)
	v1848 = v1798
	goto L558
L561:
	;
	v1846 = v1796 + int32(1)
	v1847 = v1797
	v1848 = v1798
	goto L558
L562:
	;
	v1863 = int32(1)
	goto L537
L563:
	;
	goto L564
L564:
	;
	if int32(2) < v1754 {
		goto L5
	} else {
		goto L565
	}
L565:
	;
	v1859 = v1754 + int32(1)
	goto L540
L566:
	;
	goto L521
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1876
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v1887 = F_pg_sprintf(m, v1883, int32(_a_F_NUM_processor_7), v22+int32(16))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L76
	} else {
		goto L568
	}
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v1887 + v1883
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1891))) = v1887
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1893)+4)) = int32(0)
	goto L110
L569:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v1904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1903))))
	if v1904 == int32(35) {
		goto L110
	} else {
		goto L570
	}
L570:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v1907 == int32(45) {
		goto L110
	} else {
		goto L571
	}
L571:
	;
	if v397 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1912 = F_get_th(m, v1903, int32(2))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L76
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v1993 = v1992 + l4
	if base.Ui32(v1993) <= base.Ui32(v1991) {
		goto L110
	} else {
		goto L597
	}
L575:
	;
	if (v1912^v1910)&int32(3) != 0 {
		goto L579
	} else {
		goto L580
	}
L576:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v1910 + int32(1)
	goto L141
L577:
	;
	goto L576
L578:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1968))) = uint8(v1967)
	if v1967&int32(255) == int32(0) {
		goto L577
	} else {
		goto L593
	}
L579:
	;
	v1919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1912))))
	v1966 = v1912
	v1967 = v1919
	v1968 = v1910
	goto L578
L580:
	;
	goto L581
L581:
	;
	if v1912&int32(3) != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v1923 = v1912
	v1925 = v1910
	goto L585
L583:
	;
	v1937 = v1912
	v1939 = v1910
	goto L584
L584:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1937)))
	v1944 = int32(-2139062144)
	if (int32(16843008)-v1941|v1941)&v1944 != v1944 {
		v1966 = v1937
		v1967 = v1941
		v1968 = v1939
		goto L578
	} else {
		goto L589
	}
L585:
	;
	v1926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1923))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1925))) = uint8(v1926)
	if v1926 == int32(0) {
		goto L577
	} else {
		goto L587
	}
L586:
	;
	v1937 = v1933
	v1939 = v1931
	goto L584
L587:
	;
	v1930 = int32(1)
	v1931 = v1925 + v1930
	v1933 = v1923 + v1930
	if v1933&int32(3) != 0 {
		v1923 = v1933
		v1925 = v1931
		goto L585
	} else {
		goto L588
	}
L588:
	;
	goto L586
L589:
	;
	v1949 = v1937
	v1950 = v1941
	v1951 = v1939
	goto L590
L590:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1951))) = v1950
	v1953 = int32(4)
	v1954 = v1951 + v1953
	v1956 = v1949 + v1953
	v1958 = *(*int32)(unsafe.Add(mBase, uint32(v1949)+4))
	v1961 = int32(-2139062144)
	if (int32(16843008)-v1958|v1958)&v1961 == v1961 {
		v1949 = v1956
		v1950 = v1958
		v1951 = v1954
		goto L590
	} else {
		goto L592
	}
L591:
	;
	v1966 = v1956
	v1967 = v1958
	v1968 = v1954
	goto L578
L592:
	;
	goto L591
L593:
	;
	v1975 = v1966
	v1977 = v1968
	goto L594
L594:
	;
	v1978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1975)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1977)+1)) = uint8(v1978)
	v1980 = int32(1)
	if v1978 != 0 {
		v1975 = v1975 + v1980
		v1977 = v1977 + v1980
		goto L594
	} else {
		goto L596
	}
L595:
	;
	goto L577
L596:
	;
	goto L595
L597:
	;
	v1995 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1991))))
	if base.B2i32(base.Ui64(v1995) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v1995)%64)&int64(288080842570334209) != int64(0)) != 0 {
		goto L110
	} else {
		goto L598
	}
L598:
	;
	v2005 = F_pg_mblen_range(m, v1991, v1993)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L76
	} else {
		goto L599
	}
L599:
	;
	v2007 = v2005 + v1991
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2007
	if base.Ui32(v1993) <= base.Ui32(v2007) {
		goto L110
	} else {
		goto L600
	}
L600:
	;
	v2010 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2007))))
	if base.B2i32(base.Ui64(v2010) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2010)%64)&int64(288080842570334209) != int64(0)) != 0 {
		goto L110
	} else {
		goto L601
	}
L601:
	;
	v2020 = F_pg_mblen_range(m, v2007, v1993)
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L76
	} else {
		goto L602
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2020 + v2007
	goto L110
L603:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v2032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2031))))
	if v2032 == int32(35) {
		goto L110
	} else {
		goto L604
	}
L604:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v2035 == int32(45) {
		goto L110
	} else {
		goto L605
	}
L605:
	;
	if v397 != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2040 = F_get_th(m, v2031, int32(1))
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L76
	} else {
		goto L609
	}
L607:
	;
	goto L608
L608:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v2121 = v2120 + l4
	if base.Ui32(v2121) <= base.Ui32(v2119) {
		goto L110
	} else {
		goto L631
	}
L609:
	;
	if (v2040^v2038)&int32(3) != 0 {
		goto L613
	} else {
		goto L614
	}
L610:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2038 + int32(1)
	goto L141
L611:
	;
	goto L610
L612:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2096))) = uint8(v2095)
	if v2095&int32(255) == int32(0) {
		goto L611
	} else {
		goto L627
	}
L613:
	;
	v2047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2040))))
	v2094 = v2040
	v2095 = v2047
	v2096 = v2038
	goto L612
L614:
	;
	goto L615
L615:
	;
	if v2040&int32(3) != 0 {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v2051 = v2040
	v2053 = v2038
	goto L619
L617:
	;
	v2065 = v2040
	v2067 = v2038
	goto L618
L618:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v2065)))
	v2072 = int32(-2139062144)
	if (int32(16843008)-v2069|v2069)&v2072 != v2072 {
		v2094 = v2065
		v2095 = v2069
		v2096 = v2067
		goto L612
	} else {
		goto L623
	}
L619:
	;
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2051))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2053))) = uint8(v2054)
	if v2054 == int32(0) {
		goto L611
	} else {
		goto L621
	}
L620:
	;
	v2065 = v2061
	v2067 = v2059
	goto L618
L621:
	;
	v2058 = int32(1)
	v2059 = v2053 + v2058
	v2061 = v2051 + v2058
	if v2061&int32(3) != 0 {
		v2051 = v2061
		v2053 = v2059
		goto L619
	} else {
		goto L622
	}
L622:
	;
	goto L620
L623:
	;
	v2077 = v2065
	v2078 = v2069
	v2079 = v2067
	goto L624
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2079))) = v2078
	v2081 = int32(4)
	v2082 = v2079 + v2081
	v2084 = v2077 + v2081
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2077)+4))
	v2089 = int32(-2139062144)
	if (int32(16843008)-v2086|v2086)&v2089 == v2089 {
		v2077 = v2084
		v2078 = v2086
		v2079 = v2082
		goto L624
	} else {
		goto L626
	}
L625:
	;
	v2094 = v2084
	v2095 = v2086
	v2096 = v2082
	goto L612
L626:
	;
	goto L625
L627:
	;
	v2103 = v2094
	v2105 = v2096
	goto L628
L628:
	;
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2103)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2105)+1)) = uint8(v2106)
	v2108 = int32(1)
	if v2106 != 0 {
		v2103 = v2103 + v2108
		v2105 = v2105 + v2108
		goto L628
	} else {
		goto L630
	}
L629:
	;
	goto L611
L630:
	;
	goto L629
L631:
	;
	v2123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2119))))
	if base.B2i32(base.Ui64(v2123) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2123)%64)&int64(288080842570334209) != int64(0)) != 0 {
		goto L110
	} else {
		goto L632
	}
L632:
	;
	v2133 = F_pg_mblen_range(m, v2119, v2121)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L76
	} else {
		goto L633
	}
L633:
	;
	v2135 = v2133 + v2119
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2135
	if base.Ui32(v2121) <= base.Ui32(v2135) {
		goto L110
	} else {
		goto L634
	}
L634:
	;
	v2138 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2135))))
	if base.B2i32(base.Ui64(v2138) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2138)%64)&int64(288080842570334209) != int64(0)) != 0 {
		goto L110
	} else {
		goto L635
	}
L635:
	;
	v2148 = F_pg_mblen_range(m, v2135, v2121)
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L76
	} else {
		goto L636
	}
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2148 + v2135
	goto L110
L637:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v2152 == int32(45) {
		goto L640
	} else {
		goto L641
	}
L638:
	;
	goto L639
L639:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2165))))
	if v2166 == int32(45) {
		goto L644
	} else {
		goto L645
	}
L640:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2156 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2155))) = uint8(v2156)
	goto L141
L641:
	;
	goto L642
L642:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v2159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2158)+12)))
	if v2159&int32(32) != 0 {
		goto L110
	} else {
		goto L643
	}
L643:
	;
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2163 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2162))) = uint8(v2163)
	goto L141
L644:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v2170 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2169))) = uint8(v2170)
	goto L141
L645:
	;
	goto L646
L646:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v2183 = v2182 + l4
	if base.B2i32(base.Ui32(v2166) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2166))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v2183) <= base.Ui32(v2165)) != 0 {
		goto L110
	} else {
		goto L647
	}
L647:
	;
	v2186 = F_pg_mblen_range(m, v2165, v2183)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L76
	} else {
		goto L648
	}
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2186 + v2165
	goto L110
L649:
	;
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	if v2190 == int32(43) {
		goto L652
	} else {
		goto L653
	}
L650:
	;
	goto L651
L651:
	;
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2203))))
	if v2204 == int32(43) {
		goto L656
	} else {
		goto L657
	}
L652:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2194 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v2193))) = uint8(v2194)
	goto L141
L653:
	;
	goto L654
L654:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v2197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2196)+12)))
	if v2197&int32(32) != 0 {
		goto L110
	} else {
		goto L655
	}
L655:
	;
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2201 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2200))) = uint8(v2201)
	goto L141
L656:
	;
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v2208 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v2207))) = uint8(v2208)
	goto L141
L657:
	;
	goto L658
L658:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v2221 = v2220 + l4
	if base.B2i32(base.Ui32(v2204) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2204))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v2221) <= base.Ui32(v2203)) != 0 {
		goto L110
	} else {
		goto L659
	}
L659:
	;
	v2224 = F_pg_mblen_range(m, v2203, v2221)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L76
	} else {
		goto L660
	}
L660:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2224 + v2203
	goto L110
L661:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v2228))) = uint8(v2229)
	goto L141
L662:
	;
	goto L663
L663:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2231))))
	switch v2232 - int32(43) {
	case 0:
		goto L665
	default:
		goto L664
	case 2:
		goto L666
	}
L664:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	v2252 = v2251 + l4
	if base.B2i32(base.Ui32(v2232) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2232))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v2252) <= base.Ui32(v2231)) != 0 {
		goto L110
	} else {
		goto L667
	}
L665:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v2239 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v2238))) = uint8(v2239)
	goto L141
L666:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v2236 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v2235))) = uint8(v2236)
	goto L141
L667:
	;
	v2255 = F_pg_mblen_range(m, v2231, v2252)
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L76
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v2255 + v2231
	goto L110
L669:
	;
	goto L109
L670:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
	v2330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2329))) = uint8(v2330)
	goto L6
L671:
	;
	goto L672
L672:
	;
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v2334 = v2332 - int32(1)
	v2335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2334))))
	if v2335 == int32(46) {
		goto L674
	} else {
		goto L675
	}
L673:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v2342)+4)) = v2343
	goto L6
L674:
	;
	v2338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2334))) = uint8(v2338)
	goto L673
L675:
	;
	goto L676
L676:
	;
	v2340 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2332))) = uint8(v2340)
	goto L673
L677:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L76
	} else {
		goto L678
	}
L678:
	;
	F_errmsg(m, int32(_a_F_NUM_processor_8), int32(0))
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L76
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F_NUM_processor_9), int32(_a_F_NUM_processor_10), int32(_a_F_NUM_processor_11))
	mBase = m.M
	v2401 = m.ExcPending
	if v2401 != 0 {
		goto L76
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L76
	} else {
		goto L682
	}
L682:
	;
	F_errmsg(m, int32(_a_F_NUM_processor_12), int32(0))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L76
	} else {
		goto L683
	}
L683:
	;
	F_errfinish(m, int32(_a_F_NUM_processor_9), int32(_a_F_NUM_processor_13), int32(_a_F_NUM_processor_11))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L76
	} else {
		goto L684
	}
L684:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_NormalizeSubWord(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v680 int32
	_ = v680
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v857 int32
	_ = v857
	var v870 int32
	_ = v870
	v4 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(1040)
	m.G0 = v24
	v26 = F_strlen(m, l1)
	mBase = m.M
	v30 = int32(512)
	base.MemoryFill(m, v24+int32(528), v4, v30)
	base.MemoryFill(m, v24+int32(16), v4, v30)
	if int32(256) < v26 {
		v870 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(1040)
	return v870
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = F_palloc_mul(m, int32(4), int32(1024))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	v49 = F_FindWord(m, l0, l1, int32(_a_F_NormalizeSubWord_0), l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = F_pstrdup(m, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v59 = v42
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v60 == int32(0) {
		v301 = v59
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v51
	v59 = v42 + int32(4)
	goto L8
L10:
	;
	if v39 == int32(0) {
		v839 = v301
		goto L66
	} else {
		goto L67
	}
L11:
	;
	v66 = v60
	v68 = v59
	v72 = v4
	goto L12
L12:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v84&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v301 = v288
	goto L10
L14:
	;
	v203 = v68
	v204 = int32(0)
	goto L43
L15:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if base.Ui32(int32(255)) < base.Ui32(v87) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v95 = v66
	goto L17
L17:
	;
	if v26 < v72 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v179 = v66 + int32(4)
	v185 = v72
	goto L14
L19:
	;
	goto L20
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v92 == int32(0) {
		v301 = v68
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v95 = v92
	goto L17
L22:
	;
	v97 = v72
	goto L24
L23:
	;
	v97 = v26
	goto L24
L24:
	;
	v101 = v95
	v107 = v72
	goto L25
L25:
	;
	if v107 == v97 {
		v301 = v68
		goto L10
	} else {
		goto L27
	}
L26:
	;
	v301 = v68
	goto L10
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v122 = int32(base.Ui32(v120) >> (uint(int32(1)) % 32))
	if v122 == int32(0) {
		v301 = v68
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v126 = v101 + int32(4)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v107))))
	v136 = v126 + v122*int32(12)
	v138 = v126
	goto L29
L29:
	;
	v154 = int32(12)
	v155 = base.I32_div_s(v136-v138, v154)
	v160 = v138 + int32(base.Ui32(v155)>>(uint(int32(1))%32))*v154
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v163 = v161 & int32(255)
	if v131 == v163 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v166 = v107 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v161) {
		v179 = v160
		v185 = v166
		goto L14
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v172 = base.B2i32(base.Ui32(v163) < base.Ui32(v131))
	if base.Ui32(v163) < base.Ui32(v131) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v169 != 0 {
		v101 = v169
		v107 = v166
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v301 = v68
	goto L10
L36:
	;
	v173 = v160 + int32(12)
	goto L38
L37:
	;
	v173 = v138
	goto L38
L38:
	;
	if base.Ui32(v163) < base.Ui32(v131) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v174 = v136
	goto L41
L40:
	;
	v174 = v160
	goto L41
L41:
	;
	if base.Ui32(v173) < base.Ui32(v174) {
		v136 = v174
		v138 = v173
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	v220 = v204 << (uint(int32(2)) % 32)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v221)))
	v225 = v24 + int32(528)
	v227 = F_CheckAffix(m, l1, v26, v223, l2, v225, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v295 != 0 {
		v66 = v295
		v68 = v288
		v72 = v185
		goto L12
	} else {
		goto L65
	}
L45:
	;
	v290 = v204 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if base.Ui32(v290) < base.Ui32(int32(base.Ui32(v291)>>(uint(int32(8))%32))) {
		v203 = v288
		v204 = v290
		goto L43
	} else {
		goto L64
	}
L46:
	;
	if v227 == int32(0) {
		v288 = v203
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v231+v220)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = F_FindWord(m, l0, v225, v234, l2)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v235 == int32(0) {
		v288 = v203
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v239 = int32(0)
	if int32(4088) < v203-v42 {
		v283 = v239
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v288 = v203 + v283<<(uint(int32(2))%32)
	goto L45
L51:
	;
	if v203 != v42 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v203-int32(4))))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if base.B2i32(v249 == int32(0))|base.B2i32(v249 != v252) != 0 {
		v270 = v249
		v271 = v252
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v277 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L3
	} else {
		goto L63
	}
L55:
	;
	if v270-v271 == int32(0) {
		v283 = v239
		goto L50
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v255 = v225
	v256 = v246
	goto L58
L58:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	if v260 == int32(0) {
		v270 = v260
		v271 = v259
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v270 = v260
	v271 = v259
	goto L56
L60:
	;
	v263 = int32(1)
	if v260 == v259 {
		v255 = v255 + v263
		v256 = v256 + v263
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	goto L54
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v277
	v283 = int32(1)
	goto L50
L64:
	;
	goto L44
L65:
	;
	goto L13
L66:
	;
	if v839 != v42 {
		v870 = v42
		goto L1
	} else {
		goto L182
	}
L67:
	;
	v325 = v301
	v327 = v39
	v336 = v4
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v343&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v839 = v811
	goto L66
L70:
	;
	v464 = v325
	v469 = int32(0)
	goto L99
L71:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if base.Ui32(int32(255)) < base.Ui32(v346) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v354 = v327
	goto L73
L73:
	;
	if v26 < v336 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v444 = v327 + int32(4)
	v453 = v336
	goto L70
L75:
	;
	goto L76
L76:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	if v351 == int32(0) {
		v839 = v325
		goto L66
	} else {
		goto L77
	}
L77:
	;
	v354 = v351
	goto L73
L78:
	;
	v356 = v336
	goto L80
L79:
	;
	v356 = v26
	goto L80
L80:
	;
	v364 = v354
	v373 = v336
	goto L81
L81:
	;
	if v356 == v373 {
		v839 = v325
		goto L66
	} else {
		goto L83
	}
L82:
	;
	v839 = v325
	goto L66
L83:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v381 = int32(base.Ui32(v379) >> (uint(int32(1)) % 32))
	if v381 == int32(0) {
		v839 = v325
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v385 = v364 + int32(4)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26+(v373^int32(-1))))))
	v396 = v385
	v399 = v385 + v381*int32(12)
	goto L85
L85:
	;
	v415 = int32(12)
	v416 = base.I32_div_s(v399-v396, v415)
	v421 = v396 + int32(base.Ui32(v416)>>(uint(int32(1))%32))*v415
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	v424 = v422 & int32(255)
	if v392 == v424 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L82
L87:
	;
	v427 = v373 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v422) {
		v444 = v421
		v453 = v427
		goto L70
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v433 = base.B2i32(base.Ui32(v424) < base.Ui32(v392))
	if base.Ui32(v424) < base.Ui32(v392) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v421)+8))
	if v430 != 0 {
		v364 = v430
		v373 = v427
		goto L81
	} else {
		goto L91
	}
L91:
	;
	v839 = v325
	goto L66
L92:
	;
	v434 = v421 + int32(12)
	goto L94
L93:
	;
	v434 = v396
	goto L94
L94:
	;
	if base.Ui32(v424) < base.Ui32(v392) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v435 = v399
	goto L97
L96:
	;
	v435 = v421
	goto L97
L97:
	;
	if base.Ui32(v434) < base.Ui32(v435) {
		v396 = v434
		v399 = v435
		goto L85
	} else {
		goto L98
	}
L98:
	;
	goto L86
L99:
	;
	v481 = v469 << (uint(int32(2)) % 32)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v481+v482)))
	v486 = v24 + int32(528)
	v489 = F_CheckAffix(m, l1, v26, v484, l2, v486, v24+int32(12))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L3
	} else {
		goto L102
	}
L100:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v444)+8))
	if v833 != 0 {
		v325 = v811
		v327 = v833
		v336 = v453
		goto L68
	} else {
		goto L181
	}
L101:
	;
	v828 = v469 + int32(1)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	if base.Ui32(v828) < base.Ui32(int32(base.Ui32(v829)>>(uint(int32(8))%32))) {
		v464 = v811
		v469 = v828
		goto L99
	} else {
		goto L180
	}
L102:
	;
	if v489 == int32(0) {
		v811 = v464
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v493+v481)))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v497 = F_FindWord(m, l0, v486, v496, l2)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	if v497 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v499 = int32(0)
	if int32(4088) < v464-v42 {
		v543 = v499
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v548 = v464
	goto L107
L107:
	;
	v551 = F_strlen(m, v24+int32(528))
	mBase = m.M
	v552 = int32(0)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v553 == v552 {
		v811 = v548
		goto L101
	} else {
		goto L122
	}
L108:
	;
	v548 = v464 + v543<<(uint(int32(2))%32)
	goto L107
L109:
	;
	if v464 != v42 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v464-int32(4))))
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if base.B2i32(v509 == int32(0))|base.B2i32(v509 != v512) != 0 {
		v530 = v509
		v531 = v512
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v537 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L3
	} else {
		goto L121
	}
L113:
	;
	if v530-v531 == int32(0) {
		v543 = v499
		goto L108
	} else {
		goto L120
	}
L114:
	;
	goto L113
L115:
	;
	v515 = v486
	v516 = v506
	goto L116
L116:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516)+1)))
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	if v520 == int32(0) {
		v530 = v520
		v531 = v519
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v530 = v520
	v531 = v519
	goto L114
L118:
	;
	v523 = int32(1)
	if v520 == v519 {
		v515 = v515 + v523
		v516 = v516 + v523
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L112
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v464))) = v537
	v543 = int32(1)
	goto L108
L122:
	;
	v559 = v553
	v561 = v548
	v565 = v552
	goto L123
L123:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v559))))
	if v577&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v811 = v798
	goto L101
L125:
	;
	v698 = v561
	v699 = int32(0)
	goto L154
L126:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v559)+4))
	if base.Ui32(int32(255)) < base.Ui32(v580) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v588 = v559
	goto L128
L128:
	;
	if v551 < v565 {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v674 = v559 + int32(4)
	v680 = v565
	goto L125
L130:
	;
	goto L131
L131:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v559)+12))
	if v585 == int32(0) {
		v811 = v561
		goto L101
	} else {
		goto L132
	}
L132:
	;
	v588 = v585
	goto L128
L133:
	;
	v590 = v565
	goto L135
L134:
	;
	v590 = v551
	goto L135
L135:
	;
	v594 = v588
	v600 = v565
	goto L136
L136:
	;
	if v600 == v590 {
		v811 = v561
		goto L101
	} else {
		goto L138
	}
L137:
	;
	v811 = v561
	goto L101
L138:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	v615 = int32(base.Ui32(v613) >> (uint(int32(1)) % 32))
	if v615 == int32(0) {
		v811 = v561
		goto L101
	} else {
		goto L139
	}
L139:
	;
	v619 = v594 + int32(4)
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(528)+v600))))
	v631 = v619 + v615*int32(12)
	v633 = v619
	goto L140
L140:
	;
	v649 = int32(12)
	v650 = base.I32_div_s(v631-v633, v649)
	v655 = v633 + int32(base.Ui32(v650)>>(uint(int32(1))%32))*v649
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	v658 = v656 & int32(255)
	if v626 == v658 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L137
L142:
	;
	v661 = v600 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v656) {
		v674 = v655
		v680 = v661
		goto L125
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v667 = base.B2i32(base.Ui32(v658) < base.Ui32(v626))
	if base.Ui32(v658) < base.Ui32(v626) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v655)+8))
	if v664 != 0 {
		v594 = v664
		v600 = v661
		goto L136
	} else {
		goto L146
	}
L146:
	;
	v811 = v561
	goto L101
L147:
	;
	v668 = v655 + int32(12)
	goto L149
L148:
	;
	v668 = v633
	goto L149
L149:
	;
	if base.Ui32(v658) < base.Ui32(v626) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v669 = v631
	goto L152
L151:
	;
	v669 = v655
	goto L152
L152:
	;
	if base.Ui32(v668) < base.Ui32(v669) {
		v631 = v669
		v633 = v668
		goto L140
	} else {
		goto L153
	}
L153:
	;
	goto L141
L154:
	;
	v717 = v699 << (uint(int32(2)) % 32)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v717+v718)))
	v722 = v24 + int32(16)
	v725 = F_CheckAffix(m, v24+int32(528), v551, v720, l2, v722, v24+int32(12))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L3
	} else {
		goto L157
	}
L155:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v674)+8))
	if v805 != 0 {
		v559 = v805
		v561 = v798
		v565 = v680
		goto L123
	} else {
		goto L179
	}
L156:
	;
	v800 = v699 + int32(1)
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if base.Ui32(v800) < base.Ui32(int32(base.Ui32(v801)>>(uint(int32(8))%32))) {
		v698 = v798
		v699 = v800
		goto L154
	} else {
		goto L178
	}
L157:
	;
	if v725 == int32(0) {
		v798 = v698
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v674)+4))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v729+v717)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v444)+4))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v733+v481)))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	if v732&v736&int32(128) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v742 = int32(_a_F_NormalizeSubWord_0)
	goto L161
L160:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v742 = v741
	goto L161
L161:
	;
	v743 = F_FindWord(m, l0, v722, v742, l2)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	if v743 == int32(0) {
		v798 = v698
		goto L156
	} else {
		goto L163
	}
L163:
	;
	v747 = int32(0)
	if int32(4088) < v698-v42 {
		v793 = v747
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v798 = v698 + v793<<(uint(int32(2))%32)
	goto L156
L165:
	;
	if v698 != v42 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v753 = v24 + int32(16)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v698-int32(4))))
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v756))))
	if base.B2i32(v759 == int32(0))|base.B2i32(v759 != v762) != 0 {
		v780 = v759
		v781 = v762
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	v787 = F_pstrdup(m, v24+int32(16))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L3
	} else {
		goto L177
	}
L169:
	;
	if v780-v781 == int32(0) {
		v793 = v747
		goto L164
	} else {
		goto L176
	}
L170:
	;
	goto L169
L171:
	;
	v765 = v753
	v766 = v756
	goto L172
L172:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+1)))
	if v770 == int32(0) {
		v780 = v770
		v781 = v769
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v780 = v770
	v781 = v769
	goto L170
L174:
	;
	v773 = int32(1)
	if v770 == v769 {
		v765 = v765 + v773
		v766 = v766 + v773
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	goto L168
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v698)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v698))) = v787
	v793 = int32(1)
	goto L164
L178:
	;
	goto L155
L179:
	;
	goto L124
L180:
	;
	goto L100
L181:
	;
	goto L69
L182:
	;
	F_pfree(m, v42)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	v870 = int32(0)
	goto L1
}
func F_NumRelids(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v3 = F_pull_varnos(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v8 = F_bms_del_members(m, v3, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = int64(0)
			if v8 == int32(0) {
				v55 = int32(0)
			} else {
				v16 = v8 + int32(8)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				if v17 == int32(1) {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v55 = base.I32_popcnt(v20)
				} else {
					v23 = v17 << (uint(int32(2)) % 32)
					if v23 <= int32(7) {
						if v23 == int32(0) {
							v50 = v11
						} else {
							v28 = v23
							v29 = v16
							v30 = v11
							for {
								v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
								v32 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_NumRelids[0]))))
								v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
								v34 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v33)+uint32(_c_F_NumRelids[0]))))
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
								v36 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_NumRelids[0]))))
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
								v38 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v37)+uint32(_c_F_NumRelids[0]))))
								v42 = v32 + (v34 + (v36 + (v30 + v38)))
								v43 = int32(4)
								v46 = v28 - v43
								if v46 != 0 {
									v28 = v46
									v29 = v29 + v43
									v30 = v42
									continue
								} else {
									break
								}
								break
							}
							v50 = v42
						}
					} else {
						v47 = F_pg_popcount_optimized(m, v16, v23)
						mBase = m.M
						v50 = v47
					}
					v55 = base.I32_wrap_i64(v50)
				}
			}
			F_bms_free(m, v8)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				return v55
			}
		}
	}
}
func F_nameconcatoid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v8))) = uint32(v11)
	v17 = F_pg_snprintf(m, v8+int32(16), int32(20), int32(_a_F_nameconcatoid_0), v8)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = F_strlen(m, v10)
		mBase = m.M
		if int32(64) <= v17+v21 {
			v27 = F_pg_mbcliplen(m, v10, v21, int32(63)-v17)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v29 = v27
				v31 = F_palloc0(m, int32(64))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					if v29 != 0 {
						base.MemoryCopy(m, v31, v10, v29)
					} else {
					}
					if v17 != 0 {
						base.MemoryCopy(m, v29+v31, v8+int32(16), v17)
					} else {
					}
					m.G0 = v8 + int32(48)
					return base.I64_extend_i32_u(v31)
				}
			}
		} else {
			v29 = v21
			v31 = F_palloc0(m, int32(64))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				if v29 != 0 {
					base.MemoryCopy(m, v31, v10, v29)
				} else {
				}
				if v17 != 0 {
					base.MemoryCopy(m, v29+v31, v8+int32(16), v17)
				} else {
				}
				m.G0 = v8 + int32(48)
				return base.I64_extend_i32_u(v31)
			}
		}
	}
}
func F_namegttext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(1755), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < base.I32_wrap_i64(v6)))
	}
}
func F_namehashfast(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	v3 = base.I32_wrap_i64(l0)
	v4 = F_strlen(m, v3)
	mBase = m.M
	v10 = v4 - int32(1636608432)
	if v3&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v4) {
			v119 = v3
			v120 = v4
			v121 = v10
			v122 = v10
			v123 = v10
			for {
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
				v126 = v125 + v122
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
				v130 = v129 + v123
				v132 = int32(4)
				v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
				v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
				v139 = v130 + v126
				v140 = v134 + v139
				v141 = v138 + v140
				v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
				v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
				v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
				v154 = v145 + v141
				v155 = v149 + v154
				v156 = v153 + v155
				v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
				v161 = int32(12)
				v162 = v119 + v161
				v164 = v120 - v161
				if base.Ui32(int32(11)) < base.Ui32(v164) {
					v119 = v162
					v120 = v164
					v121 = v155
					v122 = v156
					v123 = v160
					continue
				} else {
					break
				}
				break
			}
			v167 = v162
			v168 = v164
			v169 = v155
			v170 = v156
			v171 = v160
		} else {
			v167 = v3
			v168 = v4
			v169 = v10
			v170 = v10
			v171 = v10
		}
		switch v168 - int32(1) {
		case 0:
			v230 = v169
			v231 = v170
			v232 = v171
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 1:
			v223 = v169
			v224 = v170
			v225 = v171
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 2:
			v216 = v169
			v217 = v170
			v218 = v171
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 3:
			v210 = v170
			v211 = v171
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 4:
			v206 = v170
			v207 = v171
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 5:
			v200 = v170
			v201 = v171
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 6:
			v194 = v170
			v195 = v171
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 7:
			v189 = v171
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 8:
			v184 = v171
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 9:
			v179 = v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+10)))
			v179 = v175<<(uint(int32(24))%32) + v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+7)))
			v194 = v190<<(uint(int32(24))%32) + v170
			v195 = v189
			v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v200 = v196<<(uint(int32(16))%32) + v194
			v201 = v195
			v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v206 = v202<<(uint(int32(8))%32) + v200
			v207 = v201
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v210 = v206 + v208
			v211 = v207
			v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+3)))
			v216 = v212<<(uint(int32(24))%32) + v169
			v217 = v210
			v218 = v211
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v223 = v219<<(uint(int32(16))%32) + v216
			v224 = v217
			v225 = v218
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v230 = v226<<(uint(int32(8))%32) + v223
			v231 = v224
			v232 = v225
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v237 = v230 + v233
			v238 = v231
			v239 = v232
		default:
			v237 = v169
			v238 = v170
			v239 = v171
		}
	} else {
		if base.Ui32(v4) < base.Ui32(int32(12)) {
			v65 = v3
			v66 = v4
			v67 = v10
			v68 = v10
			v69 = v10
		} else {
			v17 = v3
			v18 = v4
			v19 = v10
			v20 = v10
			v21 = v10
			for {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v24 = v23 + v20
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
				v28 = v27 + v21
				v30 = int32(4)
				v32 = v25 + v19 - v28 ^ base.I32_rotl(v28, v30)
				v36 = v24 - v32 ^ base.I32_rotl(v32, int32(6))
				v37 = v28 + v24
				v38 = v32 + v37
				v39 = v36 + v38
				v43 = v37 - v36 ^ base.I32_rotl(v36, int32(8))
				v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
				v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
				v52 = v43 + v39
				v53 = v47 + v52
				v54 = v51 + v53
				v58 = v52 - v51 ^ base.I32_rotl(v51, v30)
				v59 = int32(12)
				v60 = v17 + v59
				v62 = v18 - v59
				if base.Ui32(int32(11)) < base.Ui32(v62) {
					v17 = v60
					v18 = v62
					v19 = v53
					v20 = v54
					v21 = v58
					continue
				} else {
					break
				}
				break
			}
			v65 = v60
			v66 = v62
			v67 = v53
			v68 = v54
			v69 = v58
		}
		switch v66 - int32(1) {
		case 0:
			v116 = v67
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 1:
			v111 = v67
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 2:
			v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+2)))
			v111 = v107<<(uint(int32(16))%32) + v67
			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+1)))
			v116 = v112<<(uint(int32(8))%32) + v111
			v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
			v237 = v116 + v117
			v238 = v68
			v239 = v69
		case 3:
			v104 = v68
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 4:
			v101 = v68
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 5:
			v96 = v68
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 6:
			v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+6)))
			v96 = v92<<(uint(int32(16))%32) + v68
			v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+5)))
			v101 = v97<<(uint(int32(8))%32) + v96
			v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
			v104 = v101 + v102
			v105 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v237 = v105 + v67
			v238 = v104
			v239 = v69
		case 7:
			v87 = v69
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 8:
			v82 = v69
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 9:
			v77 = v69
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		case 10:
			v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+10)))
			v77 = v73<<(uint(int32(24))%32) + v69
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+9)))
			v82 = v78<<(uint(int32(16))%32) + v77
			v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+8)))
			v87 = v83<<(uint(int32(8))%32) + v82
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
			v237 = v88 + v67
			v238 = v90 + v68
			v239 = v87
		default:
			v237 = v67
			v238 = v68
			v239 = v69
		}
	}
	v242 = int32(14)
	v244 = v238 ^ v239 - base.I32_rotl(v238, v242)
	v248 = v244 ^ v237 - base.I32_rotl(v244, int32(11))
	v252 = v248 ^ v238 - base.I32_rotl(v248, int32(25))
	v256 = v252 ^ v244 - base.I32_rotl(v252, int32(16))
	v260 = v256 ^ v248 - base.I32_rotl(v256, int32(4))
	v264 = v260 ^ v252 - base.I32_rotl(v260, v242)
	return v264 ^ v256 - base.I32_rotl(v264, int32(24))
}
func F_nameout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pstrdup(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_ndistinct_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v9 {
	case 0:
		v10 = int32(23)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = F_errsave_start(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 == int32(0) {
				v162 = v10
				m.G0 = v7 + int32(128)
				return v162
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v21
					F_errmsg(m, int32(_a_F_ndistinct_object_start_0), v7+int32(16))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v30 = F_errdetail(m, int32(_a_F_ndistinct_object_start_1), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v11, int32(_a_F_ndistinct_object_start_2), int32(76), int32(_a_F_ndistinct_object_start_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v162 = v10
								m.G0 = v7 + int32(128)
								return v162
							}
						}
					}
				}
			}
		}
	case 1:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
		v162 = int32(0)
		m.G0 = v7 + int32(128)
		return v162
	case 2:
		v37 = int32(23)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v39 = F_errsave_start(m, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			if v39 == int32(0) {
				v162 = v37
				m.G0 = v7 + int32(128)
				return v162
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v46
					F_errmsg(m, int32(_a_F_ndistinct_object_start_0), v7+int32(32))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v55 = F_errdetail(m, int32(_a_F_ndistinct_object_start_4), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v38, int32(_a_F_ndistinct_object_start_2), int32(84), int32(_a_F_ndistinct_object_start_3))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v162 = v37
								m.G0 = v7 + int32(128)
								return v162
							}
						}
					}
				}
			}
		}
	case 3:
		v62 = int32(23)
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v64 = F_errsave_start(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			if v64 == int32(0) {
				v162 = v62
				m.G0 = v7 + int32(128)
				return v162
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v71
					F_errmsg(m, int32(_a_F_ndistinct_object_start_0), v7-int32(-64))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(_a_F_ndistinct_object_start_5)
						v83 = F_errdetail(m, int32(_a_F_ndistinct_object_start_6), v7+int32(48))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v63, int32(_a_F_ndistinct_object_start_2), int32(93), int32(_a_F_ndistinct_object_start_3))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v162 = v62
								m.G0 = v7 + int32(128)
								return v162
							}
						}
					}
				}
			}
		}
	case 4:
		v90 = int32(23)
		v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v92 = F_errsave_start(m, v91)
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
			return int32(0)
		} else {
			if v92 == int32(0) {
				v162 = v90
				m.G0 = v7 + int32(128)
				return v162
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v99
					F_errmsg(m, int32(_a_F_ndistinct_object_start_0), v7+int32(80))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v108 = F_errdetail(m, int32(_a_F_ndistinct_object_start_7), int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v91, int32(_a_F_ndistinct_object_start_2), int32(101), int32(_a_F_ndistinct_object_start_3))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v162 = v90
								m.G0 = v7 + int32(128)
								return v162
							}
						}
					}
				}
			}
		}
	case 5:
		v115 = int32(23)
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v117 = F_errsave_start(m, v116)
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int32(0)
		} else {
			if v117 == int32(0) {
				v162 = v115
				m.G0 = v7 + int32(128)
				return v162
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int32(0)
				} else {
					v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+112)) = v124
					F_errmsg(m, int32(_a_F_ndistinct_object_start_0), v7+int32(112))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+96)) = int32(_a_F_ndistinct_object_start_8)
						v136 = F_errdetail(m, int32(_a_F_ndistinct_object_start_9), v7+int32(96))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v116, int32(_a_F_ndistinct_object_start_2), int32(110), int32(_a_F_ndistinct_object_start_3))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								v162 = v115
								m.G0 = v7 + int32(128)
								return v162
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v146 = m.ExcPending
		if v146 != 0 {
			return int32(0)
		} else {
			v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v147
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_ndistinct_object_start_10)
			F_errmsg_internal(m, int32(_a_F_ndistinct_object_start_11), v7)
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ndistinct_object_start_2), int32(116), int32(_a_F_ndistinct_object_start_3))
				mBase = m.M
				v158 = m.ExcPending
				if v158 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ndistinct_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ndistinct_scalar[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_ndistinct_scalar[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+144)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v17 - int32(4) {
	case 0:
		v22 = F_pg_strtoint16_safe(m, l1, v9+int32(144))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+148)))
			if v26 == int32(1) {
				v29 = int32(23)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v31 = F_errsave_start(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						v212 = v29
						m.G0 = v9 + int32(160)
						return v212
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v38
							F_errmsg(m, int32(_a_F_ndistinct_scalar_0), v9+int32(32))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_ndistinct_scalar_1)
								v50 = F_errdetail(m, int32(_a_F_ndistinct_scalar_2), v9+int32(16))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v30, int32(_a_F_ndistinct_scalar_3), int32(436), int32(_a_F_ndistinct_scalar_4))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return int32(0)
									} else {
										v212 = v29
										m.G0 = v9 + int32(160)
										return v212
									}
								}
							}
						}
					}
				}
			} else {
				if int32(-9) < v22 {
					v60 = v22
				} else {
					v60 = int32(0)
				}
				if v60 == int32(0) {
					v63 = int32(23)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v65 = F_errsave_start(m, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						if v65 == int32(0) {
							v212 = v63
							m.G0 = v9 + int32(160)
							return v212
						} else {
							F_errcode(m, int32(33685634))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v72
								F_errmsg(m, int32(_a_F_ndistinct_scalar_0), v9-int32(-64))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v22
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = int32(_a_F_ndistinct_scalar_1)
									v85 = F_errdetail(m, int32(_a_F_ndistinct_scalar_5), v9+int32(48))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, v64, int32(_a_F_ndistinct_scalar_3), int32(450), int32(_a_F_ndistinct_scalar_4))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v212 = v63
											m.G0 = v9 + int32(160)
											return v212
										}
									}
								}
							}
						}
					}
				} else {
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v92 == int32(0) {
						v146 = F_lappend_int(m, v92, v22)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146
							v212 = int32(0)
							m.G0 = v9 + int32(160)
							return v212
						}
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
						if v95 <= int32(0) {
							v146 = F_lappend_int(m, v92, v22)
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146
								v212 = int32(0)
								m.G0 = v9 + int32(160)
								return v212
							}
						} else {
							v98 = int32(_a_F_ndistinct_scalar_6)
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v100+v95<<(uint(int32(2))%32)-int32(4))))
							v110 = base.I32_extend16_s(v106)
							if int32(0) < v110 {
								v114 = base.B2i32(base.Ui32(v106&v98) < base.Ui32(v22&v98))
							} else {
								v114 = base.B2i32(v22 < v110)
							}
							if v114 != 0 {
								v146 = F_lappend_int(m, v92, v22)
								mBase = m.M
								v147 = m.ExcPending
								if v147 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146
									v212 = int32(0)
									m.G0 = v9 + int32(160)
									return v212
								}
							} else {
								v115 = int32(23)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v117 = F_errsave_start(m, v116)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									if v117 == int32(0) {
										v212 = v115
										m.G0 = v9 + int32(160)
										return v212
									} else {
										F_errcode(m, int32(33685634))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v124
											F_errmsg(m, int32(_a_F_ndistinct_scalar_0), v9+int32(96))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+88)) = v110
												*(*int32)(unsafe.Add(mBase, uint32(v9)+84)) = v22
												*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = int32(_a_F_ndistinct_scalar_1)
												v138 = F_errdetail(m, int32(_a_F_ndistinct_scalar_7), v9+int32(80))
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v116, int32(_a_F_ndistinct_scalar_3), int32(464), int32(_a_F_ndistinct_scalar_4))
													mBase = m.M
													v144 = m.ExcPending
													if v144 != 0 {
														return int32(0)
													} else {
														v212 = v115
														m.G0 = v9 + int32(160)
														return v212
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 1:
		v152 = F_pg_strtoint32_safe(m, l1, v9+int32(144))
		mBase = m.M
		v153 = m.ExcPending
		if v153 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v152
			v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+148)))
			if v155 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(2)
				v212 = int32(0)
				m.G0 = v9 + int32(160)
				return v212
			} else {
				v161 = int32(23)
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v163 = F_errsave_start(m, v162)
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return int32(0)
				} else {
					if v163 == int32(0) {
						v212 = v161
						m.G0 = v9 + int32(160)
						return v212
					} else {
						F_errcode(m, int32(33685634))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v170
							F_errmsg(m, int32(_a_F_ndistinct_scalar_0), v9+int32(128))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = int32(_a_F_ndistinct_scalar_8)
								v182 = F_errdetail(m, int32(_a_F_ndistinct_scalar_2), v9+int32(112))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, v162, int32(_a_F_ndistinct_scalar_3), int32(491), int32(_a_F_ndistinct_scalar_4))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										v212 = v161
										m.G0 = v9 + int32(160)
										return v212
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		v189 = int32(23)
		v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v191 = F_errsave_start(m, v190)
		mBase = m.M
		v192 = m.ExcPending
		if v192 != 0 {
			return int32(0)
		} else {
			if v191 == int32(0) {
				v212 = v189
				m.G0 = v9 + int32(160)
				return v212
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v197 = m.ExcPending
				if v197 != 0 {
					return int32(0)
				} else {
					v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v198
					F_errmsg(m, int32(_a_F_ndistinct_scalar_0), v9)
					mBase = m.M
					v202 = m.ExcPending
					if v202 != 0 {
						return int32(0)
					} else {
						v205 = F_errdetail(m, int32(_a_F_ndistinct_scalar_9), int32(0))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v190, int32(_a_F_ndistinct_scalar_3), int32(498), int32(_a_F_ndistinct_scalar_4))
							mBase = m.M
							v211 = m.ExcPending
							if v211 != 0 {
								return int32(0)
							} else {
								v212 = v189
								m.G0 = v9 + int32(160)
								return v212
							}
						}
					}
				}
			}
		}
	}
}
func F_newcolor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 == int32(0) {
		v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+16)))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v14 = v11 + v10*int32(24)
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v15)
			v94 = v14
			*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(0)
			v100 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v100
			v102 = int32(_a_F_newcolor_0)
			*(*uint16)(unsafe.Add(mBase, uint32(v94)+8)) = uint16(v102)
			*(*int64)(unsafe.Add(mBase, uint32(v94))) = v100
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v109 = base.I32_div_s(v94-v106, int32(24))
			return base.I32_extend16_s(v109)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if base.Ui32(v17) < base.Ui32(v18-int32(1)) {
				v23 = v17 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v94 = v25 + v23*int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(0)
				v100 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v100
				v102 = int32(_a_F_newcolor_0)
				*(*uint16)(unsafe.Add(mBase, uint32(v94)+8)) = uint16(v102)
				*(*int64)(unsafe.Add(mBase, uint32(v94))) = v100
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v109 = base.I32_div_s(v94-v106, int32(24))
				return base.I32_extend16_s(v109)
			} else {
				if v17 == int32(_a_F_newcolor_1) {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = int32(101)
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if v34 != 0 {
						v36 = v34
					} else {
						v36 = int32(20)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v36
					return int32(-1)
				} else {
					v40 = int32(_a_F_newcolor_2)
					v42 = v18 << (uint(int32(1)) % 32)
					if base.Ui32(v40) <= base.Ui32(v42) {
						v45 = v40
					} else {
						v45 = v42
					}
					v47 = v45 * int32(24)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v50 = l0 + int32(108)
					if v48 == v50 {
						v53 = F_palloc_extended(m, v47, int32(2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							if v53 == int32(0) {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
								if v72 != 0 {
									v74 = v72
								} else {
									v74 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
								return int32(-1)
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v61 = v59 * int32(24)
								if v61 == int32(0) {
									v83 = v53
								} else {
									base.MemoryCopy(m, v53, v50, v61)
									v83 = v53
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v83
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v89 = v87 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
								v94 = v83 + v89*int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(0)
								v100 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v100
								v102 = int32(_a_F_newcolor_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v94)+8)) = uint16(v102)
								*(*int64)(unsafe.Add(mBase, uint32(v94))) = v100
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v109 = base.I32_div_s(v94-v106, int32(24))
								return base.I32_extend16_s(v109)
							}
						}
					} else {
						v65 = F_repalloc_extended(m, v48, v47)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							if v65 != 0 {
								v83 = v65
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v45
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v83
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v89 = v87 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v89
								v94 = v83 + v89*int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = int32(0)
								v100 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v100
								v102 = int32(_a_F_newcolor_0)
								*(*uint16)(unsafe.Add(mBase, uint32(v94)+8)) = uint16(v102)
								*(*int64)(unsafe.Add(mBase, uint32(v94))) = v100
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v109 = base.I32_div_s(v94-v106, int32(24))
								return base.I32_extend16_s(v109)
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v68)+24)) = int32(101)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
								if v72 != 0 {
									v74 = v72
								} else {
									v74 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v71)+12)) = v74
								return int32(-1)
							}
						}
					}
				}
			}
		}
	} else {
		return int32(-1)
	}
}
func F_newdfa(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v84 int32
	_ = v84
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(base.Ui32(v9+int32(31)) >> (uint(int32(5)) % 32))
	v15 = v9 << (uint(int32(1)) % 32)
	if base.Ui32(int32(20)) < base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v149 != 0 {
		goto L54
	} else {
		goto L55
	}
L2:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+4)) = int32(0)
	if v124&int32(32) != 0 {
		goto L51
	} else {
		goto L52
	}
L3:
	;
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+69)) = uint8(v100)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+68)) = uint8(base.B2i32(l3 == v100))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+36)) = v99 + int32(3916)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+32)) = v99 + int32(1516)
	v112 = v99 + int32(1352)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+24)) = v112
	v115 = v99 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v99)+28)) = v112 + v15<<(uint(int32(2))%32)
	v121 = v115
	v122 = v99
	goto L2
L4:
	;
	v29 = v15 | int32(1)
	v30 = base.I32_div_u_s(int32(2147483647), v29)
	if base.Ui32(v13) < base.Ui32(v30) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(15) < v18 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if l3 != 0 {
		v99 = l3
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = F_palloc_extended(m, int32(_a_F_newdfa_0), int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v23 != 0 {
		v99 = v23
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	v44 = F_palloc_extended(m, int32(72), int32(2))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L19
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = base.I32_div_u_s(int32(2147483647), v15)
	if base.Ui32(v32) < base.Ui32(v34) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v36 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v38 = v36
	goto L18
L17:
	;
	v38 = int32(19)
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v38
	return int32(0)
L19:
	;
	if v44 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	goto L1
L21:
	;
	goto L22
L22:
	;
	v49 = F_palloc_mul_extended(m, int32(32), v15)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v49
	v54 = F_palloc_mul_extended(m, int32(4), v29*v13)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = v54 + v15*v13<<(uint(int32(2))%32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v65 = F_palloc_mul_extended(m, int32(4), v63*v15)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v65
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v71 = F_palloc_mul_extended(m, int32(8), v69*v15)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	v73 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+68)) = uint16(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v71
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	if v76 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v77 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	if v85 != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	F_pfree(m, v76)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L34
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v80 == int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if v71 != 0 {
		v121 = v76
		v122 = v44
		goto L2
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L29
L35:
	;
	F_pfree(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v44)+32))
	if v88 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	F_pfree(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	if v91 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	F_pfree(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+68)))
	if v94 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v44)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L1
L50:
	;
	goto L49
L51:
	;
	v130 = int32(7)
	goto L53
L52:
	;
	v130 = v15
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+56)) = v121
	*(*int64)(unsafe.Add(mBase, uint32(v122)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v122)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v122)+60)) = int64(4294967295)
	return v122
L54:
	;
	v151 = v149
	goto L56
L55:
	;
	v151 = int32(12)
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v151
	return int32(0)
}
func F_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v474 int32
	_ = v474
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v508 int32
	_ = v508
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	return int32(1)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(94)
	goto L1
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(91)
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(28)
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1026)+8)) = v1027 | int32(128)
	v1031 = int32(60)
	if v1022 == v1031 {
		goto L336
	} else {
		goto L337
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L32
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1009 != 0 {
		goto L333
	} else {
		goto L334
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1001 != 0 {
		goto L330
	} else {
		goto L331
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v995 != 0 {
		goto L327
	} else {
		goto L328
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v988 != 0 {
		goto L324
	} else {
		goto L325
	}
L14:
	;
	if base.Ui32(v962) <= base.Ui32(v959) {
		goto L317
	} else {
		goto L318
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L1
L16:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v944)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v944)+8)) = v945 | int32(16)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v949)+8)) = v950 | int32(256)
	goto L15
L17:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v933 != int32(91) {
		goto L312
	} else {
		goto L313
	}
L18:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)+8))
	v926 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v924)+8)) = v925 | v926
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v926
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v918 + int32(4)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v918)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v922
	goto L1
L20:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v905)+8)) = v906 | int32(16)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v910)+8)) = v911 | int32(256)
	goto L19
L21:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v874)+8))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+24))
	v901 = m.T0[v900].(func(*base.Module, int32, int32) int32)(m, v872, v874)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L53
	} else {
		goto L309
	}
L22:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v895)+24))
	v897 = m.T0[v896].(func(*base.Module, int32, int32) int32)(m, v312, v353)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L53
	} else {
		goto L307
	}
L23:
	;
	return v888
L24:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v867&int32(2) == int32(0) {
		goto L300
	} else {
		goto L301
	}
L25:
	;
	if base.Ui32(v167) < base.Ui32(v138) {
		goto L24
	} else {
		goto L299
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	goto L1
L27:
	;
	if v138-v167 < int32(21) {
		goto L288
	} else {
		goto L289
	}
L28:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v823 == int32(40) {
		goto L285
	} else {
		goto L286
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	goto L1
L30:
	;
	if base.Ui32(v746) <= base.Ui32(v749) {
		goto L277
	} else {
		goto L278
	}
L31:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v791)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v791)+8)) = v792 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(8589934668)
	goto L1
L32:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = int32(0)
	if base.B2i32(v18&int32(1024) == v21)|base.B2i32(v16 != int32(110)) == v21 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v784)+8)) = v785 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(12884901964)
	goto L1
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(65)
	goto L1
L35:
	;
	goto L36
L36:
	;
	if v18&int32(32) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.B2i32(base.Ui32(v137) < base.Ui32(v138))|base.B2i32(base.Ui32(int32(9)) < base.Ui32(v131)) != 0 {
		goto L67
	} else {
		goto L68
	}
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v131 = v129
	goto L37
L39:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v34))|base.B2i32(int32(1)<<(uint(v34)%32)&int32(54) == int32(0)) != 0 {
		v131 = v34
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = v45
	v50 = v44
	goto L41
L41:
	;
	if base.Ui32(v50) <= base.Ui32(v47) {
		v92 = v47
		v95 = v50
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if base.Ui32(v92) < base.Ui32(v95) {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v57 = v47
	v59 = v55
	v60 = v50
	goto L45
L45:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	if v64 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v92 = v88
	v95 = v86
	goto L43
L47:
	;
	v88 = v84 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v88
	if base.Ui32(v88) < base.Ui32(v86) {
		v57 = v88
		v59 = v85
		v60 = v86
		goto L45
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v63) {
		v92 = v57
		v95 = v60
		goto L43
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+48))
	v74 = m.T0[v73].(func(*base.Module, int32, int32) int32)(m, v63, v59)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v63)+uint32(_c_F_next[1]))))
	if v69 < int32(0) {
		v84 = v57
		v85 = v59
		v86 = v60
		goto L47
	} else {
		goto L52
	}
L52:
	;
	v92 = v57
	v95 = v60
	goto L43
L53:
	;
	return int32(0)
L54:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v74 == int32(0) {
		v92 = v79
		v95 = v78
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v84 = v79
	v85 = v83
	v86 = v78
	goto L47
L56:
	;
	goto L46
L57:
	;
	v109 = v92
	goto L63
L58:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v99 == int32(35) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v92 == v45 {
		goto L38
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v104 | int32(128)
	goto L38
L63:
	;
	v116 = v109 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	if base.Ui32(v95) <= base.Ui32(v116) {
		v47 = v116
		v50 = v95
		goto L41
	} else {
		goto L65
	}
L64:
	;
	v47 = v116
	v50 = v95
	goto L41
L65:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v119 != int32(10) {
		v109 = v116
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v167 = v137 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	switch v131 - int32(2) {
	case 0:
		goto L85
	case 1:
		goto L4
	case 2, 3:
		goto L84
	case 4:
		goto L83
	case 5:
		goto L82
	case 6:
		goto L81
	case 7:
		goto L80
	default:
		goto L79
	}
L68:
	;
	v144 = int32(1) << (uint(v131) % 32)
	if v144&int32(960) == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v144&int32(14) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	goto L13
L72:
	;
	if v144&int32(48) == int32(0) {
		goto L67
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	goto L1
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v159 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v161 = v159
	goto L78
L77:
	;
	v161 = int32(9)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v161
	goto L8
L79:
	;
	if v169 != int32(40) {
		goto L204
	} else {
		goto L205
	}
L80:
	;
	if base.B2i32(v169 != int32(58))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L201
	} else {
		goto L202
	}
L81:
	;
	if base.B2i32(v169 != int32(61))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L198
	} else {
		goto L199
	}
L82:
	;
	if base.B2i32(v169 != int32(46))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L195
	} else {
		goto L196
	}
L83:
	;
	switch v169 - int32(91) {
	case 0:
		goto L172
	case 1:
		goto L173
	case 2:
		goto L174
	default:
		goto L171
	}
L84:
	;
	v365 = v169 - int32(48)
	if base.Ui32(int32(10)) <= base.Ui32(v365) {
		goto L156
	} else {
		goto L157
	}
L85:
	;
	switch v169 - int32(36) {
	case 0:
		goto L87
	default:
		goto L4
	case 6:
		goto L91
	case 10:
		goto L89
	case 55:
		goto L90
	case 56:
		goto L86
	case 58:
		goto L88
	}
L86:
	;
	if base.Ui32(v138) <= base.Ui32(v167) {
		goto L140
	} else {
		goto L141
	}
L87:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v227&int32(32) == int32(0) {
		goto L112
	} else {
		goto L113
	}
L88:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v215 != int32(40) {
		goto L108
	} else {
		goto L109
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	goto L1
L90:
	;
	if v138-v167 < int32(21) {
		goto L96
	} else {
		goto L97
	}
L91:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v174 - int32(94) {
	case 0, 16:
		goto L93
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L92
	default:
		goto L94
	}
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	goto L1
L93:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(180388626544)
	goto L1
L94:
	;
	if v174 != int32(40) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v138) <= base.Ui32(v167) {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v186 != int32(91) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v189 != int32(58) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	switch v192 - int32(60) {
	case 0, 2:
		goto L100
	default:
		goto L96
	}
L100:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if v195 != int32(58) {
		goto L96
	} else {
		goto L101
	}
L101:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v198 != int32(93) {
		goto L96
	} else {
		goto L102
	}
L102:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	if v201 != int32(93) {
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v1022 = v192
	goto L7
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	goto L1
L105:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v208 != int32(94) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	goto L6
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(403726925936)
	goto L1
L108:
	;
	if v215 != int32(110) {
		goto L107
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v220)+8)) = v221 | int32(256)
	goto L5
L111:
	;
	goto L5
L112:
	;
	v959 = v167
	v962 = v138
	goto L14
L113:
	;
	goto L114
L114:
	;
	v233 = v167
	v236 = v138
	goto L115
L115:
	;
	if base.Ui32(v236) <= base.Ui32(v233) {
		v276 = v233
		v279 = v236
		goto L117
	} else {
		goto L118
	}
L117:
	;
	if base.Ui32(v276) < base.Ui32(v279) {
		goto L131
	} else {
		goto L132
	}
L118:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v243 = v233
	v245 = v241
	v246 = v236
	goto L119
L119:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+2)))
	if v250 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	v276 = v272
	v279 = v270
	goto L117
L121:
	;
	v272 = v268 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v272
	if base.Ui32(v272) < base.Ui32(v270) {
		v243 = v272
		v245 = v269
		v246 = v270
		goto L119
	} else {
		goto L129
	}
L122:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v249) {
		v276 = v243
		v279 = v246
		goto L117
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v245)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	v260 = m.T0[v259].(func(*base.Module, int32, int32) int32)(m, v249, v245)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L53
	} else {
		goto L127
	}
L125:
	;
	v255 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_next[1]))))
	if v255 < int32(0) {
		v268 = v243
		v269 = v245
		v270 = v246
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v276 = v243
	v279 = v246
	goto L117
L127:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v260 == int32(0) {
		v276 = v263
		v279 = v262
		goto L117
	} else {
		goto L128
	}
L128:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v268 = v263
	v269 = v267
	v270 = v262
	goto L121
L129:
	;
	goto L120
L130:
	;
	v295 = v276
	goto L136
L131:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v283 == int32(35) {
		goto L130
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if v276 == v167 {
		v959 = v276
		v962 = v279
		goto L14
	} else {
		goto L135
	}
L134:
	;
	goto L133
L135:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v287)+8)) = v288 | int32(128)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v959 = v293
	v962 = v292
	goto L14
L136:
	;
	v302 = v295 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v302
	if base.Ui32(v279) <= base.Ui32(v302) {
		v233 = v302
		v236 = v279
		goto L115
	} else {
		goto L138
	}
L137:
	;
	v233 = v302
	v236 = v279
	goto L115
L138:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if v305 != int32(10) {
		v295 = v302
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	goto L12
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	switch v312 - int32(40) {
	case 0:
		goto L148
	case 1:
		goto L147
	default:
		goto L143
	case 9, 10, 11, 12, 13, 14, 15, 16, 17:
		goto L144
	case 20:
		goto L146
	case 22:
		goto L145
	case 83:
		goto L149
	}
L143:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+2)))
	if v354 != int32(1) {
		goto L22
	} else {
		goto L150
	}
L144:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+8)) = v343 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v312 - int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(98)
	goto L1
L145:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+8)) = v336 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(62)
	goto L1
L146:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+8)) = v329 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(60)
	goto L1
L147:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	goto L1
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	goto L1
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(5)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+8)) = v318 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	goto L1
L150:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v312) {
		goto L15
	} else {
		goto L151
	}
L151:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+uint32(_c_F_next[1]))))
	if v359&int32(3) == int32(0) {
		goto L15
	} else {
		goto L152
	}
L152:
	;
	goto L16
L153:
	;
	if base.B2i32(v131 != int32(5))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L168
	} else {
		goto L169
	}
L154:
	;
	if v131 == int32(4) {
		goto L162
	} else {
		goto L163
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(44)
	goto L1
L156:
	;
	if v169 == int32(44) {
		goto L155
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v365
	goto L1
L159:
	;
	if v169 == int32(92) {
		goto L153
	} else {
		goto L160
	}
L160:
	;
	if v169 == int32(125) {
		goto L154
	} else {
		goto L161
	}
L161:
	;
	goto L11
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v383&int32(2) == int32(0))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	goto L11
L165:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	goto L1
L166:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v390 != int32(63) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v396)+8)) = v397 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(125)
	goto L1
L168:
	;
	goto L11
L169:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v409 != int32(125) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	goto L3
L171:
	;
	if v169 == int32(45) {
		goto L17
	} else {
		goto L194
	}
L172:
	;
	if base.Ui32(v138) <= base.Ui32(v167) {
		goto L187
	} else {
		goto L188
	}
L173:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v431)+8)) = v432 | int32(64)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v436&int32(2) == int32(0) {
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v418 == int32(91) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(399431958640)
	goto L1
L176:
	;
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(93)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2) - v426&int32(1)
	goto L1
L178:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(395136991344)
	goto L1
L179:
	;
	goto L180
L180:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v443)+8)) = v444 | int32(128)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v449) <= base.Ui32(v448) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	goto L12
L182:
	;
	goto L183
L183:
	;
	v452 = F_lexescape(m, l0)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L53
	} else {
		goto L184
	}
L184:
	;
	if v452 == int32(0) {
		v888 = int32(0)
		goto L23
	} else {
		goto L185
	}
L185:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v458 = v456 - int32(99)
	v465 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v458))|base.B2i32(int32(1)<<(uint(v458)%32)&int32(_a_F_next_0) == v465) == v465 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L12
L187:
	;
	goto L13
L188:
	;
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	switch v474 - int32(46) {
	case 0:
		goto L193
	default:
		goto L190
	case 12:
		goto L191
	case 15:
		goto L192
	}
L190:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(390842024048)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v167
	goto L1
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(9)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+8)) = v493 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(67)
	goto L1
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v483)+8)) = v484 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(69)
	goto L1
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(73)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(7)
	goto L1
L194:
	;
	goto L4
L195:
	;
	goto L4
L196:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v508 != int32(93) {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(197568495704)
	goto L3
L198:
	;
	goto L4
L199:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v519 != int32(93) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(261993005144)
	goto L3
L201:
	;
	goto L4
L202:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v530 != int32(93) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(249108103256)
	goto L3
L204:
	;
	switch v169 - int32(36) {
	case 0:
		goto L2
	default:
		goto L4
	case 5:
		goto L28
	case 6:
		goto L210
	case 7:
		goto L209
	case 10:
		goto L26
	case 27:
		goto L208
	case 55:
		goto L27
	case 56:
		goto L25
	case 58:
		goto L5
	case 87:
		goto L207
	case 88:
		goto L211
	}
L205:
	;
	goto L206
L206:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v727&int32(2) == int32(0))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L29
	} else {
		goto L260
	}
L207:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v609&int32(32) != 0 {
		goto L221
	} else {
		goto L222
	}
L208:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v587&int32(2) == int32(0))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L218
	} else {
		goto L219
	}
L209:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v565&int32(2) == int32(0))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L215
	} else {
		goto L216
	}
L210:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v543&int32(2) == int32(0))|base.B2i32(base.Ui32(v138) <= base.Ui32(v167)) != 0 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(124)
	goto L1
L212:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	goto L1
L213:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v550 != int32(63) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+8)) = v557 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42)
	goto L1
L215:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967339)
	goto L1
L216:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v572 != int32(63) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v578)+8)) = v579 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(43)
	goto L1
L218:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967359)
	goto L1
L219:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v594 != int32(63) {
		goto L218
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + int32(8)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v600)+8)) = v601 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(63)
	goto L1
L221:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v615 = v613
	v618 = v612
	goto L224
L222:
	;
	v693 = v167
	v695 = v138
	goto L223
L223:
	;
	if base.Ui32(v695) <= base.Ui32(v693) {
		goto L252
	} else {
		goto L253
	}
L224:
	;
	if base.Ui32(v618) <= base.Ui32(v615) {
		v658 = v615
		v661 = v618
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v693 = v690
	v695 = v689
	goto L223
L226:
	;
	goto L225
L227:
	;
	if base.Ui32(v658) < base.Ui32(v661) {
		goto L241
	} else {
		goto L242
	}
L228:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v625 = v615
	v626 = v623
	v628 = v618
	goto L229
L229:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625)))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+2)))
	if v632 == int32(1) {
		goto L232
	} else {
		goto L233
	}
L230:
	;
	v658 = v654
	v661 = v652
	goto L227
L231:
	;
	v654 = v650 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v654
	if base.Ui32(v654) < base.Ui32(v652) {
		v625 = v654
		v626 = v651
		v628 = v652
		goto L229
	} else {
		goto L239
	}
L232:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v631) {
		v658 = v625
		v661 = v628
		goto L227
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v626)+8))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)+48))
	v642 = m.T0[v641].(func(*base.Module, int32, int32) int32)(m, v631, v626)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L53
	} else {
		goto L237
	}
L235:
	;
	v637 = int32(*(*int8)(unsafe.Add(mBase, uint32(v631)+uint32(_c_F_next[1]))))
	if v637 < int32(0) {
		v650 = v625
		v651 = v626
		v652 = v628
		goto L231
	} else {
		goto L236
	}
L236:
	;
	v658 = v625
	v661 = v628
	goto L227
L237:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v642 == int32(0) {
		v658 = v645
		v661 = v644
		goto L227
	} else {
		goto L238
	}
L238:
	;
	v649 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v650 = v645
	v651 = v649
	v652 = v644
	goto L231
L239:
	;
	goto L230
L240:
	;
	v676 = v658
	goto L248
L241:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	if v665 == int32(35) {
		goto L240
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	if v658 != v613 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	goto L243
L245:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v669)+8)) = v670 | int32(128)
	goto L247
L246:
	;
	goto L247
L247:
	;
	goto L226
L248:
	;
	v683 = v676 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v683
	if base.Ui32(v661) <= base.Ui32(v683) {
		v615 = v683
		v618 = v661
		goto L224
	} else {
		goto L250
	}
L249:
	;
	v615 = v683
	v618 = v661
	goto L224
L250:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	if v686 != int32(10) {
		v676 = v683
		goto L248
	} else {
		goto L251
	}
L251:
	;
	goto L249
L252:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v715)+8)) = v716 | int32(8)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v720)+8)) = v721 | int32(256)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(528280977520)
	goto L1
L253:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v693)))
	v701 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+2)))
	if v702 == int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v699-int32(48)) {
		goto L252
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v709)+16))
	v711 = m.T0[v710].(func(*base.Module, int32, int32) int32)(m, v699, v701)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L53
	} else {
		goto L258
	}
L257:
	;
	goto L18
L258:
	;
	if v711 != 0 {
		goto L18
	} else {
		goto L259
	}
L259:
	;
	goto L252
L260:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v734 != int32(63) {
		goto L29
	} else {
		goto L261
	}
L261:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v737)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+8)) = v738 | int32(128)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v744 = v742 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v744
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v746) <= base.Ui32(v744) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	goto L10
L263:
	;
	goto L264
L264:
	;
	v749 = v742 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v742)+4))
	if v751 != int32(35) {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	goto L33
L266:
	;
	switch v751 - int32(33) {
	case 0:
		goto L31
	default:
		goto L10
	case 25:
		goto L269
	case 27:
		goto L30
	case 28:
		goto L265
	}
L267:
	;
	goto L268
L268:
	;
	if base.Ui32(v746) <= base.Ui32(v749) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(40)
	goto L1
L270:
	;
	v780 = int32(0)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v781 == v780 {
		goto L32
	} else {
		goto L276
	}
L271:
	;
	v760 = v749
	goto L272
L272:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v768 = v760 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v768
	if v766 == int32(41) {
		goto L270
	} else {
		goto L274
	}
L273:
	;
	goto L270
L274:
	;
	if base.Ui32(v768) < base.Ui32(v746) {
		v760 = v768
		goto L272
	} else {
		goto L275
	}
L275:
	;
	goto L273
L276:
	;
	v888 = v780
	goto L23
L277:
	;
	goto L10
L278:
	;
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v742 + int32(12)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v742)+8))
	if v802 != int32(33) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	goto L10
L281:
	;
	if v802 != int32(61) {
		goto L280
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v814)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v814)+8)) = v815 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(76)
	goto L1
L284:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v807)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v807)+8)) = v808 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967372)
	goto L1
L285:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v826)+8)) = v827 | int32(256)
	goto L287
L286:
	;
	goto L287
L287:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	goto L1
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v138) <= base.Ui32(v167) {
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v837 != int32(91) {
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	if v840 != int32(58) {
		goto L288
	} else {
		goto L291
	}
L291:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	switch v843 - int32(60) {
	case 0, 2:
		goto L292
	default:
		goto L288
	}
L292:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	if v846 != int32(58) {
		goto L288
	} else {
		goto L293
	}
L293:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v137)+20))
	if v849 != int32(93) {
		goto L288
	} else {
		goto L294
	}
L294:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	if v852 != int32(93) {
		goto L288
	} else {
		goto L295
	}
L295:
	;
	v1022 = v843
	goto L7
L296:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	goto L1
L297:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v859 != int32(94) {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	goto L6
L299:
	;
	goto L12
L300:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874)+2)))
	if v875 != int32(1) {
		goto L21
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v885 = F_lexescape(m, l0)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L53
	} else {
		goto L306
	}
L303:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v872) {
		goto L19
	} else {
		goto L304
	}
L304:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v872)+uint32(_c_F_next[1]))))
	if v880&int32(3) == int32(0) {
		goto L19
	} else {
		goto L305
	}
L305:
	;
	goto L20
L306:
	;
	v888 = v885
	goto L23
L307:
	;
	if v897 != 0 {
		goto L16
	} else {
		goto L308
	}
L308:
	;
	goto L15
L309:
	;
	if v901 == int32(0) {
		goto L19
	} else {
		goto L310
	}
L310:
	;
	goto L20
L311:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528402)
	goto L1
L312:
	;
	if base.Ui32(v138) <= base.Ui32(v167) {
		goto L311
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528432)
	goto L1
L315:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	if v937 != int32(93) {
		goto L311
	} else {
		goto L316
	}
L316:
	;
	goto L314
L317:
	;
	goto L2
L318:
	;
	goto L319
L319:
	;
	if v962-v959 < int32(5) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(154618822768)
	return int32(1)
L321:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	if v969 != int32(92) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v972 != int32(41) {
		goto L320
	} else {
		goto L323
	}
L323:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v975)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v975)+8)) = v976 | int32(256)
	goto L2
L324:
	;
	v990 = v988
	goto L326
L325:
	;
	v990 = int32(7)
	goto L326
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v990
	goto L8
L327:
	;
	v997 = v995
	goto L329
L328:
	;
	v997 = int32(5)
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v997
	goto L8
L330:
	;
	v1003 = v1001
	goto L332
L331:
	;
	v1003 = int32(10)
	goto L332
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1003
	goto L8
L333:
	;
	v1011 = v1009
	goto L335
L334:
	;
	v1011 = int32(13)
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1011
	goto L8
L336:
	;
	v1035 = v1031
	goto L338
L337:
	;
	v1035 = int32(62)
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1035
	goto L1
}
func F_nfanode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v9 = F_newnfa(m, l0, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v13 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	F_dupnfa(m, v9, v16, v17, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v105 = v4
	goto L5
L5:
	;
	return v105
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 != 0 {
		v32 = v25
		v33 = v4
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l2 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	F_specialcolors(m, v9)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v28 != 0 {
		v32 = v28
		v33 = v4
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v29 = F_optimize(m, v9)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = v31
	v33 = v29
	goto L7
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	if v32 != 0 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v37 = v32
	goto L15
L15:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L18
	}
L16:
	;
	F_makesearch(m, l0, v9)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = v36
	goto L15
L18:
	;
	F_compact(m, v9, l1+int32(36))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L12
L20:
	;
	v43 = v42
	goto L23
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(0)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v69 != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+136))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+136)) = v51 + v52*int32(-36) - int32(8)
	F_pfree(m, v43)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	if v49 != 0 {
		v43 = v49
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v70 = v69
	goto L30
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(0)
	F_pfree(m, v9)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+76))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+136))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77)+136)) = v78 + v79*int32(-40) - int32(8)
	F_pfree(m, v70)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	if v76 != 0 {
		v70 = v76
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v105 = v33
	goto L5
}
func F_nfatree(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	v17 = F_nfanode(m, l0, l1, int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L9
	}
L4:
	;
	v8 = F_nfatree(m, l0, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	if v12 != 0 {
		v7 = v12
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	return v17
}
func F_normalize_exec_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1355 int32
	_ = v1355
	var v1390 int32
	_ = v1390
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1427 int32
	_ = v1427
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = m.G0
	v20 = v18 - int32(_a_F_normalize_exec_path_0)
	m.G0 = v20
	if l0 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v20 + int32(_a_F_normalize_exec_path_0)
	if v1390 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L2:
	;
	v1390 = int32(0)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0])) = int32(28)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v27 = int32(_a_F_normalize_exec_path_1)
	v30 = F_memchr(m, l0, int32(0), v27)
	mBase = m.M
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v32 = v30 - l0
	goto L9
L8:
	;
	v32 = v27
	goto L9
L9:
	;
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0])) = int32(44)
	goto L2
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(int32(4095)) < base.Ui32(v32) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0])) = int32(37)
	goto L2
L14:
	;
	v40 = int32(_a_F_normalize_exec_path_2)
	v41 = v40 - v32
	v44 = v41 + (v20 + v40)
	v46 = v32 + int32(1)
	if base.Ui32(int32(512)) <= base.Ui32(v46) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v216 = v41
	v219 = v2
	v223 = v2
	v225 = v2
	v227 = v2
	goto L62
L16:
	;
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v53 = v44 + v46
	if (v44^l0)&int32(3) == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	base.MemoryCopy(m, v44, l0, v46)
	goto L21
L20:
	;
	goto L21
L21:
	;
	goto L15
L22:
	;
	if base.Ui32(v185) < base.Ui32(v53) {
		goto L56
	} else {
		goto L57
	}
L23:
	;
	if v44&int32(3) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	if base.Ui32(v53) < base.Ui32(int32(4)) {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	v89 = v53 & int32(-4)
	if base.Ui32(v53) < base.Ui32(int32(64)) {
		v139 = v83
		v140 = v84
		goto L37
	} else {
		goto L38
	}
L27:
	;
	v83 = l0
	v84 = v44
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v46 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = l0
	v84 = v44
	goto L26
L31:
	;
	goto L32
L32:
	;
	v66 = l0
	v67 = v44
	goto L33
L33:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v71)
	v73 = int32(1)
	v74 = v66 + v73
	v76 = v67 + v73
	if v76&int32(3) == int32(0) {
		v83 = v74
		v84 = v76
		goto L26
	} else {
		goto L35
	}
L34:
	;
	v83 = v74
	v84 = v76
	goto L26
L35:
	;
	if base.Ui32(v76) < base.Ui32(v53) {
		v66 = v74
		v67 = v76
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if base.Ui32(v89) <= base.Ui32(v140) {
		v184 = v139
		v185 = v140
		goto L22
	} else {
		goto L43
	}
L38:
	;
	v93 = v89 + int32(-64)
	if base.Ui32(v93) < base.Ui32(v84) {
		v139 = v83
		v140 = v84
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v96 = v83
	v97 = v84
	goto L40
L40:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v96)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+36)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+40)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v96)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v96)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+48)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+52)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v96)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+56)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v96)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+60)) = v131
	v133 = int32(-64)
	v134 = v96 - v133
	v136 = v97 - v133
	if base.Ui32(v136) <= base.Ui32(v93) {
		v96 = v134
		v97 = v136
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v139 = v134
	v140 = v136
	goto L37
L42:
	;
	goto L41
L43:
	;
	v146 = v139
	v147 = v140
	goto L44
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v151
	v153 = int32(4)
	v154 = v146 + v153
	v156 = v147 + v153
	if base.Ui32(v156) < base.Ui32(v89) {
		v146 = v154
		v147 = v156
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v184 = v154
	v185 = v156
	goto L22
L46:
	;
	goto L45
L47:
	;
	v184 = l0
	v185 = v44
	goto L22
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v46) < base.Ui32(int32(4)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v184 = l0
	v185 = v44
	goto L22
L51:
	;
	goto L52
L52:
	;
	v165 = l0
	v166 = v44
	goto L53
L53:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v170)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)) = uint8(v172)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)) = uint8(v174)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)) = uint8(v176)
	v178 = int32(4)
	v179 = v165 + v178
	v181 = v166 + v178
	if base.Ui32(v181) <= base.Ui32(v53-int32(4)) {
		v165 = v179
		v166 = v181
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v184 = v179
	v185 = v181
	goto L22
L55:
	;
	goto L54
L56:
	;
	v191 = v184
	v192 = v185
	goto L59
L57:
	;
	goto L58
L58:
	;
	goto L15
L59:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v196)
	v198 = int32(1)
	v201 = v192 + v198
	if v201 != v53 {
		v191 = v191 + v198
		v192 = v201
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	goto L60
L62:
	;
	v230 = v20 + int32(_a_F_normalize_exec_path_2)
	v231 = v230 + v216
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v232 == int32(47) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v1339 = v20 + int32(_a_F_normalize_exec_path_2) + v1326
	v1340 = v1339
	goto L364
L65:
	;
	v1326 = v1312
	v1327 = v1313
	v1331 = v1317
	v1333 = int32(0)
	goto L64
L66:
	;
	v235 = int32(1)
	v237 = v216 + v235
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v237))))
	v240 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v240)
	v242 = int32(0)
	if v239 != v240 {
		v1312 = v237
		v1313 = v235
		v1317 = v242
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L77
L69:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+2)))
	if v245 == int32(47) {
		v1312 = v237
		v1313 = v235
		v1317 = v242
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v248 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v248)
	v1312 = v237
	v1313 = int32(2)
	v1317 = v242
	goto L65
L71:
	;
	v845 = v219 + v844
	if base.Ui32(int32(4095)) < base.Ui32(v845) {
		goto L13
	} else {
		goto L223
	}
L72:
	;
	v843 = v216
	v844 = v347
	goto L71
L73:
	;
	v347 = v336 - v231
	if v347|v225 != 0 {
		goto L96
	} else {
		goto L97
	}
L74:
	;
	goto L73
L75:
	;
	v326 = v321
	goto L92
L76:
	;
	v321 = v313
	goto L75
L77:
	;
	if v231&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v259 = v231
	goto L83
L81:
	;
	v273 = v231
	goto L82
L82:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v282 = int32(-2139062144)
	if (int32(16843008)-v279|v279)&v282 != v282 {
		v313 = v273
		goto L76
	} else {
		goto L87
	}
L83:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if base.B2i32(v264 == int32(0))|base.B2i32(int32(47) == v264) != 0 {
		v336 = v259
		goto L74
	} else {
		goto L85
	}
L84:
	;
	v273 = v270
	goto L82
L85:
	;
	v270 = v259 + int32(1)
	if v270&int32(3) != 0 {
		v259 = v270
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v288 = v273
	v290 = v279
	goto L88
L88:
	;
	v294 = v290 ^ int32(791621423)
	v297 = int32(-2139062144)
	if (int32(16843008)-v294|v294)&v297 != v297 {
		v313 = v288
		goto L76
	} else {
		goto L90
	}
L89:
	;
	v321 = v303
	goto L75
L90:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	v303 = v288 + int32(4)
	v307 = int32(-2139062144)
	if (v301|(int32(16843008)-v301))&v307 == v307 {
		v288 = v303
		v290 = v301
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v328 == int32(0) {
		v336 = v326
		goto L74
	} else {
		goto L94
	}
L93:
	;
	v336 = v326
	goto L74
L94:
	;
	if v328 != int32(47) {
		v326 = v326 + int32(1)
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	if v347 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v375 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v219))) = uint8(v375)
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v379 != int32(47) {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	if v219 == int32(0) {
		goto L72
	} else {
		goto L102
	}
L100:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v351 != int32(46) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v1326 = v216 + int32(1)
	v1327 = v219
	v1331 = v223
	v1333 = v225
	goto L64
L102:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v219-int32(1)))))
	if v361 == int32(47) {
		goto L72
	} else {
		goto L103
	}
L103:
	;
	if v216 == int32(0) {
		goto L13
	} else {
		goto L104
	}
L104:
	;
	v366 = int32(1)
	v367 = v216 - v366
	v371 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v367+(v20+int32(_a_F_normalize_exec_path_2))))) = uint8(v371)
	v843 = v367
	v844 = v347 + v366
	goto L71
L105:
	;
	v383 = v20 + int32(_a_F_normalize_exec_path_2)
	v385 = F_getcwd(m, v383, int32(_a_F_normalize_exec_path_1))
	mBase = m.M
	if v385 == int32(0) {
		v1390 = v375
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v834 = F_strlen(m, v20)
	mBase = m.M
	v836 = v834 + int32(1)
	v837 = F_emscripten_builtin_malloc(m, v836)
	mBase = m.M
	if v837 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L108:
	;
	v388 = int32(0)
	v389 = F_strlen(m, v383)
	mBase = m.M
	if v223 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v390 = v389
	v392 = v388
	v397 = v223
	goto L112
L110:
	;
	v467 = v389
	v469 = v388
	goto L111
L111:
	;
	v480 = v219 - v469
	if v469 == v219 {
		v494 = v467
		goto L126
	} else {
		goto L127
	}
L112:
	;
	v406 = v392 + int32(2)
	if base.Ui32(v406) < base.Ui32(v219) {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v467 = v464
	v469 = v408
	goto L111
L114:
	;
	v408 = v392 + int32(3)
	goto L116
L115:
	;
	v408 = v406
	goto L116
L116:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v390) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v464 = v463 + v450
	v466 = v397 - int32(1)
	if v466 != 0 {
		v390 = v464
		v392 = v408
		v397 = v466
		goto L112
	} else {
		goto L125
	}
L118:
	;
	v411 = v390
	goto L121
L119:
	;
	v436 = v390
	goto L120
L120:
	;
	v450 = v436
	v463 = int32(0)
	goto L117
L121:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411+v20+int32(4095)))))
	if v428 == int32(47) {
		v450 = v411
		v463 = int32(-1)
		goto L117
	} else {
		goto L123
	}
L122:
	;
	v436 = int32(1)
	goto L120
L123:
	;
	v431 = int32(1)
	v432 = v411 - v431
	if base.Ui32(v431) < base.Ui32(v432) {
		v411 = v432
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	goto L113
L126:
	;
	if base.Ui32(v494+v480-int32(4095)) < base.Ui32(int32(-4096)) {
		goto L13
	} else {
		goto L129
	}
L127:
	;
	v484 = v20 + int32(_a_F_normalize_exec_path_2) + v467
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v484-int32(1)))))
	if v487 == int32(47) {
		v494 = v467
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v490 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v484))) = uint8(v490)
	v494 = v467 + int32(1)
	goto L126
L129:
	;
	v501 = v494 + v20
	v502 = v20 + v469
	v504 = v480 + int32(1)
	if v501 == v502 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v649 = v20 + int32(_a_F_normalize_exec_path_2)
	if base.Ui32(int32(512)) <= base.Ui32(v494) {
		goto L173
	} else {
		goto L174
	}
L131:
	;
	goto L130
L132:
	;
	v508 = v501 + v504
	if base.Ui32(v502-v508) <= base.Ui32(int32(0)-v504<<(uint(int32(1))%32)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v515 = F___memcpy(m, v501, v502, v504)
	mBase = m.M
	goto L130
L134:
	;
	goto L135
L135:
	;
	v518 = (v501 ^ v502) & int32(3)
	if base.Ui32(v501) < base.Ui32(v502) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	if v620 == int32(0) {
		goto L131
	} else {
		goto L168
	}
L137:
	;
	if base.Ui32(v598) <= base.Ui32(int32(3)) {
		v618 = v596
		v619 = v597
		v620 = v598
		goto L136
	} else {
		goto L164
	}
L138:
	;
	if v518 != 0 {
		v618 = v501
		v619 = v502
		v620 = v504
		goto L136
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if v518 != 0 {
		v580 = v504
		goto L147
	} else {
		goto L148
	}
L141:
	;
	if v501&int32(3) == int32(0) {
		v596 = v501
		v597 = v502
		v598 = v504
		goto L137
	} else {
		goto L142
	}
L142:
	;
	v524 = v501
	v525 = v502
	v526 = v504
	goto L143
L143:
	;
	if v526 == int32(0) {
		goto L131
	} else {
		goto L145
	}
L144:
	;
	v596 = v538
	v597 = v534
	v598 = v536
	goto L137
L145:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	*(*uint8)(unsafe.Add(mBase, uint32(v524))) = uint8(v531)
	v533 = int32(1)
	v534 = v525 + v533
	v536 = v526 - v533
	v538 = v524 + v533
	if v538&int32(3) != 0 {
		v524 = v538
		v525 = v534
		v526 = v536
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	if v580 == int32(0) {
		goto L131
	} else {
		goto L160
	}
L148:
	;
	if v508&int32(3) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v545 = v504
	goto L152
L150:
	;
	v560 = v504
	goto L151
L151:
	;
	if base.Ui32(v560) <= base.Ui32(int32(3)) {
		v580 = v560
		goto L147
	} else {
		goto L156
	}
L152:
	;
	if v545 == int32(0) {
		goto L131
	} else {
		goto L154
	}
L153:
	;
	v560 = v551
	goto L151
L154:
	;
	v551 = v545 - int32(1)
	v552 = v501 + v551
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v551))))
	*(*uint8)(unsafe.Add(mBase, uint32(v552))) = uint8(v554)
	if v552&int32(3) != 0 {
		v545 = v551
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v567 = v560
	goto L157
L157:
	;
	v571 = v567 - int32(4)
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v502+v571)))
	*(*int32)(unsafe.Add(mBase, uint32(v501+v571))) = v574
	if base.Ui32(int32(3)) < base.Ui32(v571) {
		v567 = v571
		goto L157
	} else {
		goto L159
	}
L158:
	;
	v580 = v571
	goto L147
L159:
	;
	goto L158
L160:
	;
	v587 = v580
	goto L161
L161:
	;
	v591 = v587 - int32(1)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502+v591))))
	*(*uint8)(unsafe.Add(mBase, uint32(v501+v591))) = uint8(v594)
	if v591 != 0 {
		v587 = v591
		goto L161
	} else {
		goto L163
	}
L162:
	;
	goto L131
L163:
	;
	goto L162
L164:
	;
	v603 = v596
	v604 = v597
	v605 = v598
	goto L165
L165:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	*(*int32)(unsafe.Add(mBase, uint32(v603))) = v608
	v610 = int32(4)
	v611 = v604 + v610
	v613 = v603 + v610
	v615 = v605 - v610
	if base.Ui32(int32(3)) < base.Ui32(v615) {
		v603 = v613
		v604 = v611
		v605 = v615
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v618 = v613
	v619 = v611
	v620 = v615
	goto L136
L167:
	;
	goto L166
L168:
	;
	v625 = v618
	v626 = v619
	v627 = v620
	goto L169
L169:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	*(*uint8)(unsafe.Add(mBase, uint32(v625))) = uint8(v630)
	v632 = int32(1)
	v637 = v627 - v632
	if v637 != 0 {
		v625 = v625 + v632
		v626 = v626 + v632
		v627 = v637
		goto L169
	} else {
		goto L171
	}
L170:
	;
	goto L131
L171:
	;
	goto L170
L172:
	;
	goto L107
L173:
	;
	if v494 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	goto L175
L175:
	;
	v656 = v20 + v494
	if (v20^v649)&int32(3) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	base.MemoryCopy(m, v20, v649, v494)
	goto L178
L177:
	;
	goto L178
L178:
	;
	goto L172
L179:
	;
	if base.Ui32(v788) < base.Ui32(v656) {
		goto L213
	} else {
		goto L214
	}
L180:
	;
	if v20&int32(3) == int32(0) {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	goto L182
L182:
	;
	if base.Ui32(v656) < base.Ui32(int32(4)) {
		goto L204
	} else {
		goto L205
	}
L183:
	;
	v692 = v656 & int32(-4)
	if base.Ui32(v656) < base.Ui32(int32(64)) {
		v742 = v686
		v743 = v687
		goto L194
	} else {
		goto L195
	}
L184:
	;
	v686 = v649
	v687 = v20
	goto L183
L185:
	;
	goto L186
L186:
	;
	if v494 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v686 = v649
	v687 = v20
	goto L183
L188:
	;
	goto L189
L189:
	;
	v669 = v649
	v670 = v20
	goto L190
L190:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v669))))
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v674)
	v676 = int32(1)
	v677 = v669 + v676
	v679 = v670 + v676
	if v679&int32(3) == int32(0) {
		v686 = v677
		v687 = v679
		goto L183
	} else {
		goto L192
	}
L191:
	;
	v686 = v677
	v687 = v679
	goto L183
L192:
	;
	if base.Ui32(v679) < base.Ui32(v656) {
		v669 = v677
		v670 = v679
		goto L190
	} else {
		goto L193
	}
L193:
	;
	goto L191
L194:
	;
	if base.Ui32(v692) <= base.Ui32(v743) {
		v787 = v742
		v788 = v743
		goto L179
	} else {
		goto L200
	}
L195:
	;
	v696 = v692 + int32(-64)
	if base.Ui32(v696) < base.Ui32(v687) {
		v742 = v686
		v743 = v687
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v699 = v686
	v700 = v687
	goto L197
L197:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v704
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+4)) = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+8)) = v708
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+12)) = v710
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v699)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+16)) = v712
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v699)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+20)) = v714
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v699)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+24)) = v716
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v699)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+28)) = v718
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v699)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+32)) = v720
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v699)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+36)) = v722
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v699)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+40)) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v699)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+44)) = v726
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v699)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+48)) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v699)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+52)) = v730
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v699)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+56)) = v732
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v699)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v700)+60)) = v734
	v736 = int32(-64)
	v737 = v699 - v736
	v739 = v700 - v736
	if base.Ui32(v739) <= base.Ui32(v696) {
		v699 = v737
		v700 = v739
		goto L197
	} else {
		goto L199
	}
L198:
	;
	v742 = v737
	v743 = v739
	goto L194
L199:
	;
	goto L198
L200:
	;
	v749 = v742
	v750 = v743
	goto L201
L201:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	*(*int32)(unsafe.Add(mBase, uint32(v750))) = v754
	v756 = int32(4)
	v757 = v749 + v756
	v759 = v750 + v756
	if base.Ui32(v759) < base.Ui32(v692) {
		v749 = v757
		v750 = v759
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v787 = v757
	v788 = v759
	goto L179
L203:
	;
	goto L202
L204:
	;
	v787 = v649
	v788 = v20
	goto L179
L205:
	;
	goto L206
L206:
	;
	if base.Ui32(v494) < base.Ui32(int32(4)) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v787 = v649
	v788 = v20
	goto L179
L208:
	;
	goto L209
L209:
	;
	v768 = v649
	v769 = v20
	goto L210
L210:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768))))
	*(*uint8)(unsafe.Add(mBase, uint32(v769))) = uint8(v773)
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v769)+1)) = uint8(v775)
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v769)+2)) = uint8(v777)
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v768)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v769)+3)) = uint8(v779)
	v781 = int32(4)
	v782 = v768 + v781
	v784 = v769 + v781
	if base.Ui32(v784) <= base.Ui32(v656-int32(4)) {
		v768 = v782
		v769 = v784
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v787 = v782
	v788 = v784
	goto L179
L212:
	;
	goto L211
L213:
	;
	v794 = v787
	v795 = v788
	goto L216
L214:
	;
	goto L215
L215:
	;
	goto L172
L216:
	;
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	*(*uint8)(unsafe.Add(mBase, uint32(v795))) = uint8(v799)
	v801 = int32(1)
	v804 = v795 + v801
	if v804 != v656 {
		v794 = v794 + v801
		v795 = v804
		goto L216
	} else {
		goto L218
	}
L217:
	;
	goto L215
L218:
	;
	goto L217
L219:
	;
	v1390 = v842
	goto L1
L220:
	;
	v842 = int32(0)
	goto L219
L221:
	;
	goto L222
L222:
	;
	v841 = F___memcpy(m, v837, v20, v836)
	mBase = m.M
	v842 = v841
	goto L219
L223:
	;
	v848 = v20 + v219
	v850 = v20 + int32(_a_F_normalize_exec_path_2)
	v851 = v850 + v843
	if base.Ui32(int32(512)) <= base.Ui32(v844) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	v1022 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20+v845))) = uint8(v1022)
	v1024 = v843 + v844
	v1025 = int32(1)
	if v347 != int32(2) {
		v1047 = v1025
		goto L273
	} else {
		goto L274
	}
L225:
	;
	if v844 != 0 {
		goto L228
	} else {
		goto L229
	}
L226:
	;
	goto L227
L227:
	;
	v858 = v848 + v844
	if (v848^v851)&int32(3) == int32(0) {
		goto L232
	} else {
		goto L233
	}
L228:
	;
	base.MemoryCopy(m, v848, v851, v844)
	goto L230
L229:
	;
	goto L230
L230:
	;
	goto L224
L231:
	;
	if base.Ui32(v990) < base.Ui32(v858) {
		goto L265
	} else {
		goto L266
	}
L232:
	;
	if v848&int32(3) == int32(0) {
		goto L236
	} else {
		goto L237
	}
L233:
	;
	goto L234
L234:
	;
	if base.Ui32(v858) < base.Ui32(int32(4)) {
		goto L256
	} else {
		goto L257
	}
L235:
	;
	v894 = v858 & int32(-4)
	if base.Ui32(v858) < base.Ui32(int32(64)) {
		v944 = v888
		v945 = v889
		goto L246
	} else {
		goto L247
	}
L236:
	;
	v888 = v851
	v889 = v848
	goto L235
L237:
	;
	goto L238
L238:
	;
	if v844 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v888 = v851
	v889 = v848
	goto L235
L240:
	;
	goto L241
L241:
	;
	v871 = v851
	v872 = v848
	goto L242
L242:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871))))
	*(*uint8)(unsafe.Add(mBase, uint32(v872))) = uint8(v876)
	v878 = int32(1)
	v879 = v871 + v878
	v881 = v872 + v878
	if v881&int32(3) == int32(0) {
		v888 = v879
		v889 = v881
		goto L235
	} else {
		goto L244
	}
L243:
	;
	v888 = v879
	v889 = v881
	goto L235
L244:
	;
	if base.Ui32(v881) < base.Ui32(v858) {
		v871 = v879
		v872 = v881
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	if base.Ui32(v894) <= base.Ui32(v945) {
		v989 = v944
		v990 = v945
		goto L231
	} else {
		goto L252
	}
L247:
	;
	v898 = v894 + int32(-64)
	if base.Ui32(v898) < base.Ui32(v889) {
		v944 = v888
		v945 = v889
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v901 = v888
	v902 = v889
	goto L249
L249:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v901)))
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = v906
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v901)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+4)) = v908
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v901)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+8)) = v910
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v901)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+12)) = v912
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v901)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+16)) = v914
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v901)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+20)) = v916
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v901)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+24)) = v918
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v901)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+28)) = v920
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v901)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+32)) = v922
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v901)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+36)) = v924
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v901)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+40)) = v926
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v901)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+44)) = v928
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v901)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+48)) = v930
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v901)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+52)) = v932
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v901)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+56)) = v934
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v901)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v902)+60)) = v936
	v938 = int32(-64)
	v939 = v901 - v938
	v941 = v902 - v938
	if base.Ui32(v941) <= base.Ui32(v898) {
		v901 = v939
		v902 = v941
		goto L249
	} else {
		goto L251
	}
L250:
	;
	v944 = v939
	v945 = v941
	goto L246
L251:
	;
	goto L250
L252:
	;
	v951 = v944
	v952 = v945
	goto L253
L253:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v951)))
	*(*int32)(unsafe.Add(mBase, uint32(v952))) = v956
	v958 = int32(4)
	v959 = v951 + v958
	v961 = v952 + v958
	if base.Ui32(v961) < base.Ui32(v894) {
		v951 = v959
		v952 = v961
		goto L253
	} else {
		goto L255
	}
L254:
	;
	v989 = v959
	v990 = v961
	goto L231
L255:
	;
	goto L254
L256:
	;
	v989 = v851
	v990 = v848
	goto L231
L257:
	;
	goto L258
L258:
	;
	if base.Ui32(v844) < base.Ui32(int32(4)) {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v989 = v851
	v990 = v848
	goto L231
L260:
	;
	goto L261
L261:
	;
	v970 = v851
	v971 = v848
	goto L262
L262:
	;
	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	*(*uint8)(unsafe.Add(mBase, uint32(v971))) = uint8(v975)
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v971)+1)) = uint8(v977)
	v979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v971)+2)) = uint8(v979)
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v971)+3)) = uint8(v981)
	v983 = int32(4)
	v984 = v970 + v983
	v986 = v971 + v983
	if base.Ui32(v986) <= base.Ui32(v858-int32(4)) {
		v970 = v984
		v971 = v986
		goto L262
	} else {
		goto L264
	}
L263:
	;
	v989 = v984
	v990 = v986
	goto L231
L264:
	;
	goto L263
L265:
	;
	v996 = v989
	v997 = v990
	goto L268
L266:
	;
	goto L267
L267:
	;
	goto L224
L268:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	*(*uint8)(unsafe.Add(mBase, uint32(v997))) = uint8(v1001)
	v1003 = int32(1)
	v1006 = v997 + v1003
	if v1006 != v858 {
		v996 = v996 + v1003
		v997 = v1006
		goto L268
	} else {
		goto L270
	}
L269:
	;
	goto L267
L270:
	;
	goto L269
L271:
	;
	v1270 = v1269
	v1273 = v219
	goto L344
L272:
	;
	v1269 = int32(1)
	goto L271
L273:
	;
	v1051 = F_readlink(m, v20, v20+int32(_a_F_normalize_exec_path_2), v1024)
	mBase = m.M
	if v1051 == v1024 {
		goto L13
	} else {
		goto L281
	}
L274:
	;
	v1028 = v1024 + v850
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028-int32(2)))))
	if v1031 != int32(46) {
		v1047 = v1025
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028-int32(1)))))
	if v1036 != int32(46) {
		v1047 = v1025
		goto L273
	} else {
		goto L276
	}
L276:
	;
	if base.Ui32(v219) <= base.Ui32(v223*int32(3)) {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1326 = v1024
	v1327 = v845
	v1331 = v223 + int32(1)
	v1333 = v225
	goto L64
L278:
	;
	goto L279
L279:
	;
	v1044 = int32(0)
	if v225 == v1044 {
		goto L272
	} else {
		goto L280
	}
L280:
	;
	v1047 = v1044
	goto L273
L281:
	;
	if v1051 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0])) = int32(44)
	goto L2
L283:
	;
	goto L284
L284:
	;
	if v1051 < int32(0) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0]))
	if v1061 != int32(28) {
		goto L2
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v1073 = v227 + int32(1)
	if v1073 == int32(40) {
		goto L293
	} else {
		goto L294
	}
L288:
	;
	v1064 = int32(0)
	if v1047 == v1064 {
		v1269 = v1064
		goto L271
	} else {
		goto L289
	}
L289:
	;
	if v347 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1067 = v845
	goto L292
L291:
	;
	v1067 = v219
	goto L292
L292:
	;
	v1071 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+int32(_a_F_normalize_exec_path_2)+v1024))))
	v1326 = v1024
	v1327 = v1067
	v1331 = v223
	v1333 = v1071
	goto L64
L293:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_normalize_exec_path[0])) = int32(32)
	goto L2
L294:
	;
	goto L295
L295:
	;
	v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v1051+int32(4095)))))
	if v1082 == int32(47) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1087 = v1024
	goto L299
L297:
	;
	v1106 = v1024
	goto L298
L298:
	;
	v1119 = v1106 - v1051
	v1121 = v20 + int32(_a_F_normalize_exec_path_2)
	v1122 = v1119 + v1121
	if v1122 == v1121 {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087+(v20+int32(_a_F_normalize_exec_path_2))))))
	if v1103 == int32(47) {
		v1087 = v1087 + int32(1)
		goto L299
	} else {
		goto L301
	}
L300:
	;
	v1106 = v1087
	goto L298
L301:
	;
	goto L300
L302:
	;
	v216 = v1119
	v227 = v1073
	goto L62
L303:
	;
	goto L302
L304:
	;
	v1126 = v1122 + v1051
	if base.Ui32(v1121-v1126) <= base.Ui32(int32(0)-v1051<<(uint(int32(1))%32)) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1133 = F___memcpy(m, v1122, v1121, v1051)
	mBase = m.M
	goto L302
L306:
	;
	goto L307
L307:
	;
	v1136 = (v1122 ^ v1121) & int32(3)
	if base.Ui32(v1122) < base.Ui32(v1121) {
		goto L310
	} else {
		goto L311
	}
L308:
	;
	if v1238 == int32(0) {
		goto L303
	} else {
		goto L340
	}
L309:
	;
	if base.Ui32(v1216) <= base.Ui32(int32(3)) {
		v1236 = v1214
		v1237 = v1215
		v1238 = v1216
		goto L308
	} else {
		goto L336
	}
L310:
	;
	if v1136 != 0 {
		v1236 = v1122
		v1237 = v1121
		v1238 = v1051
		goto L308
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	if v1136 != 0 {
		v1198 = v1051
		goto L319
	} else {
		goto L320
	}
L313:
	;
	if v1122&int32(3) == int32(0) {
		v1214 = v1122
		v1215 = v1121
		v1216 = v1051
		goto L309
	} else {
		goto L314
	}
L314:
	;
	v1142 = v1122
	v1143 = v1121
	v1144 = v1051
	goto L315
L315:
	;
	if v1144 == int32(0) {
		goto L303
	} else {
		goto L317
	}
L316:
	;
	v1214 = v1156
	v1215 = v1152
	v1216 = v1154
	goto L309
L317:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1143))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1142))) = uint8(v1149)
	v1151 = int32(1)
	v1152 = v1143 + v1151
	v1154 = v1144 - v1151
	v1156 = v1142 + v1151
	if v1156&int32(3) != 0 {
		v1142 = v1156
		v1143 = v1152
		v1144 = v1154
		goto L315
	} else {
		goto L318
	}
L318:
	;
	goto L316
L319:
	;
	if v1198 == int32(0) {
		goto L303
	} else {
		goto L332
	}
L320:
	;
	if v1126&int32(3) != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1163 = v1051
	goto L324
L322:
	;
	v1178 = v1051
	goto L323
L323:
	;
	if base.Ui32(v1178) <= base.Ui32(int32(3)) {
		v1198 = v1178
		goto L319
	} else {
		goto L328
	}
L324:
	;
	if v1163 == int32(0) {
		goto L303
	} else {
		goto L326
	}
L325:
	;
	v1178 = v1169
	goto L323
L326:
	;
	v1169 = v1163 - int32(1)
	v1170 = v1122 + v1169
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121+v1169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1170))) = uint8(v1172)
	if v1170&int32(3) != 0 {
		v1163 = v1169
		goto L324
	} else {
		goto L327
	}
L327:
	;
	goto L325
L328:
	;
	v1185 = v1178
	goto L329
L329:
	;
	v1189 = v1185 - int32(4)
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1121+v1189)))
	*(*int32)(unsafe.Add(mBase, uint32(v1122+v1189))) = v1192
	if base.Ui32(int32(3)) < base.Ui32(v1189) {
		v1185 = v1189
		goto L329
	} else {
		goto L331
	}
L330:
	;
	v1198 = v1189
	goto L319
L331:
	;
	goto L330
L332:
	;
	v1205 = v1198
	goto L333
L333:
	;
	v1209 = v1205 - int32(1)
	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121+v1209))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1122+v1209))) = uint8(v1212)
	if v1209 != 0 {
		v1205 = v1209
		goto L333
	} else {
		goto L335
	}
L334:
	;
	goto L303
L335:
	;
	goto L334
L336:
	;
	v1221 = v1214
	v1222 = v1215
	v1223 = v1216
	goto L337
L337:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1222)))
	*(*int32)(unsafe.Add(mBase, uint32(v1221))) = v1226
	v1228 = int32(4)
	v1229 = v1222 + v1228
	v1231 = v1221 + v1228
	v1233 = v1223 - v1228
	if base.Ui32(int32(3)) < base.Ui32(v1233) {
		v1221 = v1231
		v1222 = v1229
		v1223 = v1233
		goto L337
	} else {
		goto L339
	}
L338:
	;
	v1236 = v1231
	v1237 = v1229
	v1238 = v1233
	goto L308
L339:
	;
	goto L338
L340:
	;
	v1243 = v1236
	v1244 = v1237
	v1245 = v1238
	goto L341
L341:
	;
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1243))) = uint8(v1248)
	v1250 = int32(1)
	v1255 = v1245 - v1250
	if v1255 != 0 {
		v1243 = v1243 + v1250
		v1244 = v1244 + v1250
		v1245 = v1255
		goto L341
	} else {
		goto L343
	}
L342:
	;
	goto L303
L343:
	;
	goto L342
L344:
	;
	if v1270 == int32(0) {
		goto L347
	} else {
		goto L348
	}
L346:
	;
	v1270 = int32(1)
	goto L344
L347:
	;
	if v1273 != 0 {
		goto L346
	} else {
		goto L350
	}
L348:
	;
	goto L349
L349:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v1273-int32(1)))))
	if v1289 != int32(47) {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	v1312 = v1024
	v1313 = int32(0)
	v1317 = v223
	goto L65
L351:
	;
	v1270 = int32(0)
	v1273 = v1273 - int32(1)
	goto L344
L352:
	;
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1294 = int32(0)
	if v1273 == int32(1) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1326 = v1024
	v1327 = int32(1)
	v1331 = v223
	v1333 = v1294
	goto L64
L356:
	;
	goto L357
L357:
	;
	v1299 = v1273 - int32(1)
	v1301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v1301 == int32(47) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1304 = int32(2)
	goto L360
L359:
	;
	v1304 = v1299
	goto L360
L360:
	;
	if v1273 != int32(2) {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1307 = v1299
	goto L363
L362:
	;
	v1307 = v1304
	goto L363
L363:
	;
	v1326 = v1024
	v1327 = v1307
	v1331 = v223
	v1333 = v1294
	goto L64
L364:
	;
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1340))))
	if v1355 == int32(47) {
		v1340 = v1340 + int32(1)
		goto L364
	} else {
		goto L366
	}
L365:
	;
	v216 = v1326 + (v1340 - v1339)
	v219 = v1327
	v223 = v1331
	v225 = v1333
	goto L62
L366:
	;
	goto L365
L367:
	;
	m.G0 = v16 + int32(16)
	return v1548
L368:
	;
	v1408 = int32(-1)
	v1411 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L371
	} else {
		goto L372
	}
L369:
	;
	goto L370
L370:
	;
	goto L380
L371:
	;
	return int32(0)
L372:
	;
	if v1411 == int32(0) {
		v1548 = v1408
		goto L367
	} else {
		goto L373
	}
L373:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L371
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(_a_F_normalize_exec_path_3), v16)
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L371
	} else {
		goto L375
	}
L375:
	;
	F_errfinish(m, int32(_a_F_normalize_exec_path_4), int32(254), int32(_a_F_normalize_exec_path_5))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L371
	} else {
		goto L376
	}
L376:
	;
	v1548 = v1408
	goto L367
L377:
	;
	F_emscripten_builtin_free(m, v1390)
	mBase = m.M
	v1548 = v2
	goto L367
L378:
	;
	v1544 = F_strlen(m, v1533)
	mBase = m.M
	goto L377
L380:
	;
	goto L381
L381:
	;
	v1434 = int32(1023)
	if (l0^v1390)&int32(3) != 0 {
		goto L385
	} else {
		goto L386
	}
L382:
	;
	v1537 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1534))) = uint8(v1537)
	goto L378
L383:
	;
	v1518 = v1513
	v1519 = v1514
	v1520 = v1515
	goto L404
L384:
	;
	if v1508 == int32(0) {
		v1533 = v1506
		v1534 = v1507
		goto L382
	} else {
		goto L403
	}
L385:
	;
	v1506 = v1390
	v1507 = l0
	v1508 = v1434
	goto L384
L386:
	;
	goto L387
L387:
	;
	v1438 = int32(0)
	if base.B2i32(v1390&int32(3) == v1438)|int32(0) == v1438 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	if v1474 == int32(0) {
		v1533 = v1471
		v1534 = v1472
		goto L382
	} else {
		goto L397
	}
L389:
	;
	v1450 = v1390
	v1451 = l0
	v1452 = v1434
	goto L392
L390:
	;
	goto L391
L391:
	;
	v1471 = v1390
	v1472 = l0
	v1473 = v1434
	v1474 = int32(1)
	goto L388
L392:
	;
	v1454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1450))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1451))) = uint8(v1454)
	if v1454 == int32(0) {
		v1513 = v1450
		v1514 = v1451
		v1515 = v1452
		goto L383
	} else {
		goto L394
	}
L393:
	;
	v1471 = v1465
	v1472 = v1459
	v1473 = v1461
	v1474 = v1463
	goto L388
L394:
	;
	v1458 = int32(1)
	v1459 = v1451 + v1458
	v1461 = v1452 - v1458
	v1462 = int32(0)
	v1463 = base.B2i32(v1461 != v1462)
	v1465 = v1450 + v1458
	if v1465&int32(3) == v1462 {
		v1471 = v1465
		v1472 = v1459
		v1473 = v1461
		v1474 = v1463
		goto L388
	} else {
		goto L395
	}
L395:
	;
	if v1461 != 0 {
		v1450 = v1465
		v1451 = v1459
		v1452 = v1461
		goto L392
	} else {
		goto L396
	}
L396:
	;
	goto L393
L397:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471))))
	if base.B2i32(v1477 == int32(0))|base.B2i32(base.Ui32(v1473) < base.Ui32(int32(4))) != 0 {
		v1506 = v1471
		v1507 = v1472
		v1508 = v1473
		goto L384
	} else {
		goto L398
	}
L398:
	;
	v1484 = v1471
	v1485 = v1472
	v1486 = v1473
	goto L399
L399:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1484)))
	v1492 = int32(-2139062144)
	if (int32(16843008)-v1489|v1489)&v1492 != v1492 {
		v1513 = v1484
		v1514 = v1485
		v1515 = v1486
		goto L383
	} else {
		goto L401
	}
L400:
	;
	v1506 = v1500
	v1507 = v1498
	v1508 = v1502
	goto L384
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1485))) = v1489
	v1497 = int32(4)
	v1498 = v1485 + v1497
	v1500 = v1484 + v1497
	v1502 = v1486 - v1497
	if base.Ui32(int32(3)) < base.Ui32(v1502) {
		v1484 = v1500
		v1485 = v1498
		v1486 = v1502
		goto L399
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	v1513 = v1506
	v1514 = v1507
	v1515 = v1508
	goto L383
L404:
	;
	v1522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1519))) = uint8(v1522)
	if v1522 == int32(0) {
		v1533 = v1518
		v1534 = v1519
		goto L382
	} else {
		goto L406
	}
L405:
	;
	v1533 = v1529
	v1534 = v1527
	goto L382
L406:
	;
	v1526 = int32(1)
	v1527 = v1519 + v1526
	v1529 = v1518 + v1526
	v1531 = v1520 - v1526
	if v1531 != 0 {
		v1518 = v1529
		v1519 = v1527
		v1520 = v1531
		goto L404
	} else {
		goto L407
	}
L407:
	;
	goto L405
}
func F_nulltestsel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 float32
	_ = v22
	var v23 float64
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 float64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 float64
	_ = v81
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	F_examine_variable(m, l0, l2, l3, v8+int32(-32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return float64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
		if v18 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v22 = *(*float32)(unsafe.Add(mBase, uint32(v19+v20)+8))
			v23 = base.F64_promote_f32(v22)
			switch l1 {
			case 0:
				v68 = v23
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
				m.T0[v69].(func(*base.Module, int32))(m, v18)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return float64(0)
				} else {
					if base.F64_lt(v68, float64(0)) != 0 {
						v81 = float64(0)
					} else {
						if base.F64_gt(v68, float64(1)) == int32(0) {
							v81 = v68
						} else {
							v81 = float64(1)
						}
					}
					m.G0 = v10 - int32(-64)
					return v81
				}
			case 1:
				v68 = base.F64_sub(float64(1), v23)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
				m.T0[v69].(func(*base.Module, int32))(m, v18)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return float64(0)
				} else {
					if base.F64_lt(v68, float64(0)) != 0 {
						v81 = float64(0)
					} else {
						if base.F64_gt(v68, float64(1)) == int32(0) {
							v81 = v68
						} else {
							v81 = float64(1)
						}
					}
					m.G0 = v10 - int32(-64)
					return v81
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return float64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
					F_errmsg_internal(m, int32(_a_F_nulltestsel_0), v8+int32(-48))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1818), int32(_a_F_nulltestsel_2))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
			if v39 == int32(0) {
				switch l1 {
				case 0:
					v81 = float64(0.005)
					m.G0 = v10 - int32(-64)
					return v81
				case 1:
					v81 = float64(0.995)
					m.G0 = v10 - int32(-64)
					return v81
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errmsg_internal(m, int32(_a_F_nulltestsel_0), v10)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1846), int32(_a_F_nulltestsel_2))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
				if v42 != int32(6) {
					switch l1 {
					case 0:
						v81 = float64(0.005)
						m.G0 = v10 - int32(-64)
						return v81
					case 1:
						v81 = float64(0.995)
						m.G0 = v10 - int32(-64)
						return v81
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return float64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg_internal(m, int32(_a_F_nulltestsel_0), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1846), int32(_a_F_nulltestsel_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
					if int32(0) <= v45 {
						switch l1 {
						case 0:
							v81 = float64(0.005)
							m.G0 = v10 - int32(-64)
							return v81
						case 1:
							v81 = float64(0.995)
							m.G0 = v10 - int32(-64)
							return v81
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return float64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
								F_errmsg_internal(m, int32(_a_F_nulltestsel_0), v10)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1846), int32(_a_F_nulltestsel_2))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if l1 != 0 {
							v50 = float64(1)
						} else {
							v50 = float64(0)
						}
						v81 = v50
						m.G0 = v10 - int32(-64)
						return v81
					}
				}
			}
		}
	}
}
func F_numericvar_serialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_enlargeStringInfo(m, l0, int32(4))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v10+v11))) = base.I32_rotr(v6, int32(24))&v15 | base.I32_rotr(v6&v15, int32(8))
	v23 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_enlargeStringInfo(m, l0, v23)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v30+v31))) = base.I32_rotr(v26, int32(24))&v35 | base.I32_rotr(v26&v35, int32(8))
	v43 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30 + v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_enlargeStringInfo(m, l0, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v50+v51))) = base.I32_rotr(v46, int32(24))&v55 | base.I32_rotr(v46&v55, int32(8))
	v63 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50 + v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	F_enlargeStringInfo(m, l0, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v70+v71))) = base.I32_rotr(v66, int32(24))&v75 | base.I32_rotr(v66&v75, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v70 + int32(4)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v86 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v92 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95+v92<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, l0, int32(2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = int32(8)
	v110 = v99<<(uint(v106)%32) | int32(base.Ui32(v99)>>(uint(v106)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v103+v104))) = uint16(v110)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(2)
	v116 = v92 + int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v116 < v117 {
		v92 = v116
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
}
func F_numst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	v6 = l1 + int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 != 0 {
		v9 = v7
		v10 = v6
		for {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v10
			v14 = v10 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
			if v15 != 0 {
				v17 = v15
				v18 = v14
				for {
					v19 = F_numst(m, v17, v18)
					mBase = m.M
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
					if v20 != 0 {
						v17 = v20
						v18 = v19
						continue
					} else {
						break
					}
					break
				}
				v23 = v19
			} else {
				v23 = v14
			}
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			if v24 != 0 {
				v9 = v24
				v10 = v23
				continue
			} else {
				break
			}
			break
		}
		v27 = v23
	} else {
		v27 = v6
	}
	return v27
}
