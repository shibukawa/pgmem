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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v585 int32
	_ = v585
	var v587 int64
	_ = v587
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
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
	if base.Ui32(int32(75)) <= base.Ui32(l0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(12)
	v23 = F_palloc(m, l0*v19+v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	if int32(2147483646) <= v50 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v25)
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v27
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v29
	F_parse_format(m, v23, v13, int32(_a_F_NUM_cache_0), v27, int32(_a_F_NUM_cache_1), int32(2), l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_pfree(m, v13)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	return v23
L9:
	;
	if v48 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v160 = v50
	goto L11
L11:
	;
	if v48 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v152 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v152
	v160 = v152
	goto L11
L13:
	;
	v56 = v48 & int32(3)
	v57 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v48) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = v57
	v72 = v5
	goto L17
L15:
	;
	v103 = v57
	goto L16
L16:
	;
	v115 = v103
	v126 = v5
	goto L21
L17:
	;
	v75 = v62 << (uint(int32(2)) % 32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_NUM_cache[2])))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+976))
	v78 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+976)) = v77 >> (uint(v78) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_NUM_cache[3])))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+976)) = v82 >> (uint(v78) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_NUM_cache[4])))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+976)) = v87 >> (uint(v78) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_NUM_cache[5])))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+976))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+976)) = v92 >> (uint(v78) % 32)
	v96 = int32(4)
	v97 = v62 + v96
	v99 = v72 + v96
	if v99 != v48&int32(2147483644) {
		v62 = v97
		v72 = v99
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v56 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v103 = v97
	goto L16
L21:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v115<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+976))
	v131 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+976)) = v130 >> (uint(v131) % 32)
	v137 = v126 + v131
	if v137 != v56 {
		v115 = v115 + v131
		v126 = v137
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
	v617 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v617)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v605)+992))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v619
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v605)+988))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v621
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v605)+980))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v623
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v605)+984))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v625
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v605)+996))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v627
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v605)+1012))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v629
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v605)+1000))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v631
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v605)+1004))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v633
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v605)+1008))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v635
	F_pfree(m, v13)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L125
	}
L25:
	;
	v585 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v573)+1012)) = v585
	v587 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v573)+1004)) = v587
	*(*int64)(unsafe.Add(mBase, uint32(v573)+996)) = v587
	*(*int64)(unsafe.Add(mBase, uint32(v573)+988)) = v587
	*(*int64)(unsafe.Add(mBase, uint32(v573)+980)) = v587
	F_parse_format(m, v573, v13, int32(_a_F_NUM_cache_0), v585, int32(_a_F_NUM_cache_1), int32(2), v573+int32(980))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L124
	}
L26:
	;
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+975)) = uint8(v443)
	v446 = v431 + int32(900)
	goto L96
L27:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[6]))
	v286 = F_MemoryContextAllocZero(m, v284, int32(1016))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L61
	}
L28:
	;
	v171 = int32(0)
	goto L30
L29:
	;
	v268 = v160 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v183)+976)) = v268
	v605 = v183
	goto L24
L30:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v171<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+975)))
	if v184 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v48 < int32(20) {
		goto L27
	} else {
		goto L44
	}
L32:
	;
	v188 = v183 + int32(900)
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if base.B2i32(v191 == int32(0))|base.B2i32(v191 != v194) != 0 {
		v212 = v191
		v213 = v194
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v218 = v171 + int32(1)
	if v218 != v48 {
		v171 = v218
		goto L30
	} else {
		goto L43
	}
L35:
	;
	if v212-v213 == int32(0) {
		goto L29
	} else {
		goto L42
	}
L36:
	;
	goto L35
L37:
	;
	v197 = v188
	v198 = v13
	goto L38
L38:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if v202 == int32(0) {
		v212 = v202
		v213 = v201
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v212 = v202
	v213 = v201
	goto L36
L40:
	;
	v205 = int32(1)
	if v202 == v201 {
		v197 = v197 + v205
		v198 = v198 + v205
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
	v222 = int32(1)
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[2]))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+975)))
	if v225 != v222 {
		v431 = v224
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v228 = v224
	v230 = v222
	goto L46
L46:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v230<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+975)))
	if v243 != int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v431 = v242
	goto L26
L49:
	;
	goto L50
L50:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242)+976))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v228)+976))
	if v246 < v247 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v249 = v242
	goto L53
L52:
	;
	v249 = v228
	goto L53
L53:
	;
	v251 = v230 + int32(1)
	if v251 == int32(20) {
		v431 = v249
		goto L26
	} else {
		goto L54
	}
L54:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2])))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+975)))
	if v257 != int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v431 = v256
	goto L26
L56:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)+976))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v249)+976))
	if v260 < v261 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v263 = v256
	goto L60
L59:
	;
	v263 = v249
	goto L60
L60:
	;
	v228 = v263
	v230 = v230 + int32(2)
	goto L46
L61:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v289<<(uint(int32(2))%32))+uint32(_c_F_NUM_cache[2]))) = v286
	v295 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+975)) = uint8(v295)
	v298 = v286 + int32(900)
	goto L65
L62:
	;
	v418 = int32(_a_F_NUM_cache_2)
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	v421 = int32(1)
	v422 = v420 + v421
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v286)+976)) = v422
	v425 = int32(_a_F_NUM_cache_3)
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[0])) = v427 + v421
	v573 = v286
	goto L25
L63:
	;
	v415 = F_strlen(m, v404)
	mBase = m.M
	goto L62
L65:
	;
	goto L66
L66:
	;
	v305 = int32(74)
	if (v298^v13)&int32(3) != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v408)
	goto L63
L68:
	;
	v389 = v384
	v390 = v385
	v391 = v386
	goto L89
L69:
	;
	if v379 == int32(0) {
		v404 = v377
		v405 = v378
		goto L67
	} else {
		goto L88
	}
L70:
	;
	v377 = v13
	v378 = v298
	v379 = v305
	goto L69
L71:
	;
	goto L72
L72:
	;
	v309 = int32(0)
	if base.B2i32(v13&int32(3) == v309)|int32(0) == v309 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v345 == int32(0) {
		v404 = v342
		v405 = v343
		goto L67
	} else {
		goto L82
	}
L74:
	;
	v321 = v13
	v322 = v298
	v323 = v305
	goto L77
L75:
	;
	goto L76
L76:
	;
	v342 = v13
	v343 = v298
	v344 = v305
	v345 = int32(1)
	goto L73
L77:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v322))) = uint8(v325)
	if v325 == int32(0) {
		v384 = v321
		v385 = v322
		v386 = v323
		goto L68
	} else {
		goto L79
	}
L78:
	;
	v342 = v336
	v343 = v330
	v344 = v332
	v345 = v334
	goto L73
L79:
	;
	v329 = int32(1)
	v330 = v322 + v329
	v332 = v323 - v329
	v333 = int32(0)
	v334 = base.B2i32(v332 != v333)
	v336 = v321 + v329
	if v336&int32(3) == v333 {
		v342 = v336
		v343 = v330
		v344 = v332
		v345 = v334
		goto L73
	} else {
		goto L80
	}
L80:
	;
	if v332 != 0 {
		v321 = v336
		v322 = v330
		v323 = v332
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if base.B2i32(v348 == int32(0))|base.B2i32(base.Ui32(v344) < base.Ui32(int32(4))) != 0 {
		v377 = v342
		v378 = v343
		v379 = v344
		goto L69
	} else {
		goto L83
	}
L83:
	;
	v355 = v342
	v356 = v343
	v357 = v344
	goto L84
L84:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v363 = int32(-2139062144)
	if (int32(16843008)-v360|v360)&v363 != v363 {
		v384 = v355
		v385 = v356
		v386 = v357
		goto L68
	} else {
		goto L86
	}
L85:
	;
	v377 = v371
	v378 = v369
	v379 = v373
	goto L69
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v360
	v368 = int32(4)
	v369 = v356 + v368
	v371 = v355 + v368
	v373 = v357 - v368
	if base.Ui32(int32(3)) < base.Ui32(v373) {
		v355 = v371
		v356 = v369
		v357 = v373
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v384 = v377
	v385 = v378
	v386 = v379
	goto L68
L89:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v393)
	if v393 == int32(0) {
		v404 = v389
		v405 = v390
		goto L67
	} else {
		goto L91
	}
L90:
	;
	v404 = v400
	v405 = v398
	goto L67
L91:
	;
	v397 = int32(1)
	v398 = v390 + v397
	v400 = v389 + v397
	v402 = v391 - v397
	if v402 != 0 {
		v389 = v400
		v390 = v398
		v391 = v402
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v566 = int32(_a_F_NUM_cache_2)
	v568 = *(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1]))
	v570 = v568 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_NUM_cache[1])) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v431)+976)) = v570
	v573 = v431
	goto L25
L94:
	;
	v563 = F_strlen(m, v552)
	mBase = m.M
	goto L93
L96:
	;
	goto L97
L97:
	;
	v453 = int32(74)
	if (v446^v13)&int32(3) != 0 {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v556 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v553))) = uint8(v556)
	goto L94
L99:
	;
	v537 = v532
	v538 = v533
	v539 = v534
	goto L120
L100:
	;
	if v527 == int32(0) {
		v552 = v525
		v553 = v526
		goto L98
	} else {
		goto L119
	}
L101:
	;
	v525 = v13
	v526 = v446
	v527 = v453
	goto L100
L102:
	;
	goto L103
L103:
	;
	v457 = int32(0)
	if base.B2i32(v13&int32(3) == v457)|int32(0) == v457 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v493 == int32(0) {
		v552 = v490
		v553 = v491
		goto L98
	} else {
		goto L113
	}
L105:
	;
	v469 = v13
	v470 = v446
	v471 = v453
	goto L108
L106:
	;
	goto L107
L107:
	;
	v490 = v13
	v491 = v446
	v492 = v453
	v493 = int32(1)
	goto L104
L108:
	;
	v473 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	*(*uint8)(unsafe.Add(mBase, uint32(v470))) = uint8(v473)
	if v473 == int32(0) {
		v532 = v469
		v533 = v470
		v534 = v471
		goto L99
	} else {
		goto L110
	}
L109:
	;
	v490 = v484
	v491 = v478
	v492 = v480
	v493 = v482
	goto L104
L110:
	;
	v477 = int32(1)
	v478 = v470 + v477
	v480 = v471 - v477
	v481 = int32(0)
	v482 = base.B2i32(v480 != v481)
	v484 = v469 + v477
	if v484&int32(3) == v481 {
		v490 = v484
		v491 = v478
		v492 = v480
		v493 = v482
		goto L104
	} else {
		goto L111
	}
L111:
	;
	if v480 != 0 {
		v469 = v484
		v470 = v478
		v471 = v480
		goto L108
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490))))
	if base.B2i32(v496 == int32(0))|base.B2i32(base.Ui32(v492) < base.Ui32(int32(4))) != 0 {
		v525 = v490
		v526 = v491
		v527 = v492
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v503 = v490
	v504 = v491
	v505 = v492
	goto L115
L115:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	v511 = int32(-2139062144)
	if (int32(16843008)-v508|v508)&v511 != v511 {
		v532 = v503
		v533 = v504
		v534 = v505
		goto L99
	} else {
		goto L117
	}
L116:
	;
	v525 = v519
	v526 = v517
	v527 = v521
	goto L100
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v504))) = v508
	v516 = int32(4)
	v517 = v504 + v516
	v519 = v503 + v516
	v521 = v505 - v516
	if base.Ui32(int32(3)) < base.Ui32(v521) {
		v503 = v519
		v504 = v517
		v505 = v521
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v532 = v525
	v533 = v526
	v534 = v527
	goto L99
L120:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	*(*uint8)(unsafe.Add(mBase, uint32(v538))) = uint8(v541)
	if v541 == int32(0) {
		v552 = v537
		v553 = v538
		goto L98
	} else {
		goto L122
	}
L121:
	;
	v552 = v548
	v553 = v546
	goto L98
L122:
	;
	v545 = int32(1)
	v546 = v538 + v545
	v548 = v537 + v545
	v550 = v539 - v545
	if v550 != 0 {
		v537 = v548
		v538 = v546
		v539 = v550
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v603 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v573)+975)) = uint8(v603)
	v605 = v573
	goto L24
L125:
	;
	return v605
}
func F_NUM_processor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v240 int32
	_ = v240
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v308 int32
	_ = v308
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
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
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v925 int32
	_ = v925
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1273 int32
	_ = v1273
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1326 int32
	_ = v1326
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1420 int32
	_ = v1420
	var v1459 int64
	_ = v1459
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1494 int32
	_ = v1494
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1554 int32
	_ = v1554
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1685 int32
	_ = v1685
	var v1723 int32
	_ = v1723
	var v1734 int32
	_ = v1734
	var v1744 int32
	_ = v1744
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1832 int32
	_ = v1832
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1852 int32
	_ = v1852
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1889 int32
	_ = v1889
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1938 int32
	_ = v1938
	var v1944 int32
	_ = v1944
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1985 int32
	_ = v1985
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2084 int32
	_ = v2084
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2114 int32
	_ = v2114
	var v2115 int32
	_ = v2115
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2121 int32
	_ = v2121
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2145 int32
	_ = v2145
	var v2156 int64
	_ = v2156
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2170 int64
	_ = v2170
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2189 int32
	_ = v2189
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2243 int32
	_ = v2243
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2273 int64
	_ = v2273
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int64
	_ = v2287
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2333 int32
	_ = v2333
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2341 int32
	_ = v2341
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2431 int32
	_ = v2431
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2460 int32
	_ = v2460
	var v2467 int32
	_ = v2467
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2522 int32
	_ = v2522
	var v2525 int32
	_ = v2525
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2569 int32
	_ = v2569
	var v2572 int32
	_ = v2572
	var v2582 int32
	_ = v2582
	var v2586 int32
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2596 int32
	_ = v2596
	var v2600 int32
	_ = v2600
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2607 int32
	_ = v2607
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2624 int32
	_ = v2624
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2667 int32
	_ = v2667
	var v2670 int32
	_ = v2670
	var v2677 int32
	_ = v2677
	var v2682 int32
	_ = v2682
	var v2685 int32
	_ = v2685
	var v2688 int32
	_ = v2688
	var v2693 int32
	_ = v2693
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2704 int32
	_ = v2704
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2710 int32
	_ = v2710
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2719 int32
	_ = v2719
	var v2723 int32
	_ = v2723
	var v2733 int32
	_ = v2733
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2766 int32
	_ = v2766
	var v2771 int32
	_ = v2771
	var v2772 int32
	_ = v2772
	var v2773 int32
	_ = v2773
	var v2775 int32
	_ = v2775
	var v2776 int32
	_ = v2776
	var v2778 int32
	_ = v2778
	var v2780 int32
	_ = v2780
	var v2783 int32
	_ = v2783
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2802 int32
	_ = v2802
	var v2810 int32
	_ = v2810
	var v2820 int32
	_ = v2820
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2828 int32
	_ = v2828
	var v2833 int32
	_ = v2833
	var v2836 int32
	_ = v2836
	var v2840 int32
	_ = v2840
	var v2846 int32
	_ = v2846
	var v2868 int32
	_ = v2868
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2876 int32
	_ = v2876
	var v2881 int32
	_ = v2881
	var v2884 int32
	_ = v2884
	var v2888 int32
	_ = v2888
	var v2894 int32
	_ = v2894
	var v2908 int32
	_ = v2908
	var v2919 int32
	_ = v2919
	var v2923 int32
	_ = v2923
	var v2939 int32
	_ = v2939
	var v2959 int32
	_ = v2959
	var v2973 int32
	_ = v2973
	var v2989 int32
	_ = v2989
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3012 int32
	_ = v3012
	var v3014 int32
	_ = v3014
	var v3069 int32
	_ = v3069
	var v3072 int32
	_ = v3072
	var v3076 int32
	_ = v3076
	var v3081 int32
	_ = v3081
	v9 = int32(0)
	v47 = m.G0
	v49 = v47 - int32(112)
	m.G0 = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v51 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v51 - int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v55&int32(_a_F_NUM_processor_0) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L70
	} else {
		goto L763
	}
L5:
	;
	m.G0 = v49 + int32(112)
	return
L6:
	;
	if l7 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if l7 != 0 {
		goto L32
	} else {
		goto L33
	}
L9:
	;
	if (l3^l2)&int32(3) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L5
L11:
	;
	goto L10
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v113)
	if v113&int32(255) == int32(0) {
		goto L11
	} else {
		goto L27
	}
L13:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v112 = l3
	v113 = v65
	v114 = l2
	goto L12
L14:
	;
	goto L15
L15:
	;
	if l3&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = l3
	v71 = l2
	goto L19
L17:
	;
	v83 = l3
	v85 = l2
	goto L18
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = int32(-2139062144)
	if (int32(16843008)-v87|v87)&v90 != v90 {
		v112 = v83
		v113 = v87
		v114 = v85
		goto L12
	} else {
		goto L23
	}
L19:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v72)
	if v72 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	v83 = v79
	v85 = v77
	goto L18
L21:
	;
	v76 = int32(1)
	v77 = v71 + v76
	v79 = v69 + v76
	if v79&int32(3) != 0 {
		v69 = v79
		v71 = v77
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v95 = v83
	v96 = v87
	v97 = v85
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v96
	v99 = int32(4)
	v100 = v97 + v99
	v102 = v95 + v99
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v107 = int32(-2139062144)
	if (int32(16843008)-v104|v104)&v107 == v107 {
		v95 = v102
		v96 = v104
		v97 = v100
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v112 = v102
	v113 = v104
	v114 = v100
	goto L12
L26:
	;
	goto L25
L27:
	;
	v121 = v112
	v123 = v114
	goto L28
L28:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)) = uint8(v124)
	v126 = int32(1)
	if v124 != 0 {
		v121 = v121 + v126
		v123 = v123 + v126
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L11
L30:
	;
	goto L29
L31:
	;
	v400 = int32(_a_F_NUM_processor_1)
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v401 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L32:
	;
	v135 = v55 & int32(768)
	if v135 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v346 = int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(l3))) = uint16(v346)
	v348 = int32(0)
	v358 = v348
	v359 = v348
	v367 = v9
	v369 = v9
	v399 = v344 + v345 - int32(1)
	goto L31
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v172 = int32(34)
	if v167&v172 != v172 {
		v308 = v9
		goto L45
	} else {
		goto L46
	}
L36:
	;
	if v135 == int32(512) {
		v167 = v55
		v168 = v9
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v141 = int32(0)
	if base.B2i32(v55&int32(32) == v141)|base.B2i32(l6 == int32(45)) == v141 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v167 = v55
	v168 = int32(1)
	goto L35
L40:
	;
	v149 = v55 & int32(-17281)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v149
	v151 = v149
	goto L42
L41:
	;
	v151 = v55
	goto L42
L42:
	;
	v158 = base.B2i32(l6 == int32(43)) & base.B2i32(v151&int32(96) == int32(32))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v159 != int32(-1) {
		v167 = v151
		v168 = v158
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v162 != v163 {
		v167 = v151
		v168 = v158
		goto L35
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	v167 = v151
	v168 = v158
	goto L35
L45:
	;
	v358 = l5
	v359 = l6
	v367 = v308
	v369 = v168
	v399 = v169 + v170 - base.B2i32(l5|v168 != int32(0))
	goto L31
L46:
	;
	v176 = int32(46)
	v177 = F___strchrnul(m, l3, v176)
	mBase = m.M
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v179 == v176 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v183 == int32(0) {
		v308 = v9
		goto L45
	} else {
		goto L51
	}
L48:
	;
	v183 = v177
	goto L50
L49:
	;
	v183 = int32(0)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v194 = v183
	goto L52
L52:
	;
	v240 = v194
	goto L54
L53:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v283 <= l5 {
		v308 = v194
		goto L45
	} else {
		goto L58
	}
L54:
	;
	v279 = v240 + int32(1)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v280 == int32(48) {
		v240 = v279
		goto L54
	} else {
		goto L56
	}
L55:
	;
	if v280 != 0 {
		v194 = v279
		goto L52
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L53
L58:
	;
	v285 = F_strlen(m, l3)
	mBase = m.M
	v287 = v285 - int32(1)
	v288 = v283 - l5
	if v287 < v288 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v290 = v287
	goto L61
L60:
	;
	v290 = v288
	goto L61
L61:
	;
	v291 = l3 + v290
	if base.Ui32(v194) < base.Ui32(v291) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v293 = v291
	goto L64
L63:
	;
	v293 = v194
	goto L64
L64:
	;
	v308 = v293
	goto L45
L65:
	;
	v451 = int32(1)
	v453 = l3 + (l7 ^ v451)
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v454 == v451 {
		goto L102
	} else {
		goto L103
	}
L66:
	;
	v446 = int32(_a_F_NUM_processor_2)
	v447 = int32(_a_F_NUM_processor_3)
	v448 = int32(_a_F_NUM_processor_4)
	v449 = int32(_a_F_NUM_processor_5)
	v450 = v400
	goto L65
L67:
	;
	goto L68
L68:
	;
	v408 = F_PGLC_localeconv(m)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v408)+32))
	if v414 != 0 {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	return
L71:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v408)+36))
	if v410 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v411 != 0 {
		v413 = v410
		goto L69
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v413 = int32(_a_F_NUM_processor_4)
	goto L69
L75:
	;
	goto L74
L76:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	if v418 != 0 {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v415 != 0 {
		v417 = v414
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v417 = int32(_a_F_NUM_processor_3)
	goto L76
L80:
	;
	goto L79
L81:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v423&int32(4) != 0 {
		goto L86
	} else {
		goto L87
	}
L82:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418))))
	if v419 != 0 {
		v421 = v418
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v421 = int32(_a_F_NUM_processor_5)
	goto L81
L85:
	;
	goto L84
L86:
	;
	v426 = v421
	goto L88
L87:
	;
	v426 = int32(_a_F_NUM_processor_5)
	goto L88
L88:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v408)+4))
	if v427 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v408)+16))
	if v438 == int32(0) {
		v446 = v437
		v447 = v417
		v448 = v413
		v449 = v426
		v450 = v400
		goto L65
	} else {
		goto L98
	}
L90:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427))))
	if v428 != 0 {
		v437 = v427
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v430 != int32(44) {
		v437 = int32(_a_F_NUM_processor_2)
		goto L89
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	if v435 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v436 = int32(_a_F_NUM_processor_2)
	goto L97
L96:
	;
	v436 = int32(_a_F_NUM_processor_5)
	goto L97
L97:
	;
	v437 = v436
	goto L89
L98:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	if v441 == int32(0) {
		v446 = v437
		v447 = v417
		v448 = v413
		v449 = v426
		v450 = v400
		goto L65
	} else {
		goto L99
	}
L99:
	;
	v446 = v437
	v447 = v417
	v448 = v413
	v449 = v426
	v450 = v438
	goto L65
L100:
	;
	v3008 = v2973 - int32(1)
	v3009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3008))))
	if v3009 == int32(46) {
		goto L760
	} else {
		goto L761
	}
L101:
	;
	if l7 == int32(0) {
		v2973 = v2923
		v2989 = v2939
		goto L100
	} else {
		goto L758
	}
L102:
	;
	v2919 = l2
	v2923 = v453
	v2939 = v9
	goto L101
L103:
	;
	goto L104
L104:
	;
	v458 = base.B2i32(v359 == int32(45))
	if v359 == int32(45) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v459 = v448
	goto L107
L106:
	;
	v459 = v447
	goto L107
L107:
	;
	v463 = base.B2i32(v359 == int32(43))
	if v359 == int32(43) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v464 = int32(32)
	goto L110
L109:
	;
	v464 = int32(62)
	goto L110
L110:
	;
	if v359 == int32(43) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v467 = int32(32)
	goto L113
L112:
	;
	v467 = int32(60)
	goto L113
L113:
	;
	v468 = l2 + l4
	v474 = l0
	v482 = l2
	v484 = v454
	v486 = v453
	v487 = v9
	v490 = v369
	v495 = v9
	v498 = v9
	v502 = v9
	v508 = v9
	goto L114
L114:
	;
	if l7 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v2919 = v2868
	v2923 = v2872
	v2939 = v2888
	goto L101
L116:
	;
	v2908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+12)))
	if v2908 != int32(1) {
		v474 = v474 + int32(12)
		v482 = v2868
		v484 = v2908
		v486 = v2872
		v487 = v2873
		v490 = v2876
		v495 = v2881
		v498 = v2884
		v502 = v2888
		v508 = v2894
		goto L114
	} else {
		goto L757
	}
L117:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v474)+8))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v612)+8))
	switch v613 {
	case 0:
		goto L160
	case 1, 2, 3, 6:
		goto L161
	default:
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	case 9:
		goto L159
	case 10:
		goto L158
	case 11:
		goto L154
	case 12:
		goto L153
	case 14, 30:
		goto L157
	case 15:
		goto L152
	case 18:
		goto L155
	case 34:
		goto L156
	}
L118:
	;
	v609 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L70
	} else {
		goto L146
	}
L119:
	;
	if base.Ui32(v468) <= base.Ui32(v482) {
		v2973 = v486
		v2989 = v502
		goto L100
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	if v484&int32(255) == int32(2) {
		goto L117
	} else {
		goto L124
	}
L122:
	;
	if v484&int32(255) != int32(2) {
		goto L118
	} else {
		goto L123
	}
L123:
	;
	goto L117
L124:
	;
	v532 = v474 + int32(1)
	if (v532^v482)&int32(3) != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v607 = F_strlen(m, v482)
	mBase = m.M
	v2868 = v607 + v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L126:
	;
	goto L125
L127:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v587))) = uint8(v586)
	if v586&int32(255) == int32(0) {
		goto L126
	} else {
		goto L142
	}
L128:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	v585 = v532
	v586 = v538
	v587 = v482
	goto L127
L129:
	;
	goto L130
L130:
	;
	if v532&int32(3) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v542 = v532
	v544 = v482
	goto L134
L132:
	;
	v556 = v532
	v558 = v482
	goto L133
L133:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v563 = int32(-2139062144)
	if (int32(16843008)-v560|v560)&v563 != v563 {
		v585 = v556
		v586 = v560
		v587 = v558
		goto L127
	} else {
		goto L138
	}
L134:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v542))))
	*(*uint8)(unsafe.Add(mBase, uint32(v544))) = uint8(v545)
	if v545 == int32(0) {
		goto L126
	} else {
		goto L136
	}
L135:
	;
	v556 = v552
	v558 = v550
	goto L133
L136:
	;
	v549 = int32(1)
	v550 = v544 + v549
	v552 = v542 + v549
	if v552&int32(3) != 0 {
		v542 = v552
		v544 = v550
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v568 = v556
	v569 = v560
	v570 = v558
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570))) = v569
	v572 = int32(4)
	v573 = v570 + v572
	v575 = v568 + v572
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v568)+4))
	v580 = int32(-2139062144)
	if (int32(16843008)-v577|v577)&v580 == v580 {
		v568 = v575
		v569 = v577
		v570 = v573
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v585 = v575
	v586 = v577
	v587 = v573
	goto L127
L141:
	;
	goto L140
L142:
	;
	v594 = v585
	v596 = v587
	goto L143
L143:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)) = uint8(v597)
	v599 = int32(1)
	if v597 != 0 {
		v594 = v594 + v599
		v596 = v596 + v599
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L126
L145:
	;
	goto L144
L146:
	;
	v2868 = v609 + v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L147:
	;
	v2868 = v2820 + int32(1)
	v2872 = v2824
	v2873 = v2825
	v2876 = v2828
	v2881 = v2833
	v2884 = v2836
	v2888 = v2840
	v2894 = v2846
	goto L116
L148:
	;
	if int32(1)<<(uint(v613)%32)&int32(78) != 0 {
		goto L644
	} else {
		goto L645
	}
L149:
	;
	v2446 = int32(1)
	v2449 = v482 + v2446
	v2450 = v2444
	v2451 = v2446
	goto L148
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L70
	} else {
		goto L640
	}
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1160)
	v2820 = v1008
	v2824 = v1009
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v1010
	v2840 = v1011
	v2846 = v1012
	goto L147
L152:
	;
	if l7 != 0 {
		goto L632
	} else {
		goto L633
	}
L153:
	;
	if l7 != 0 {
		goto L620
	} else {
		goto L621
	}
L154:
	;
	if l7 != 0 {
		goto L608
	} else {
		goto L609
	}
L155:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2183&int32(1024)|v2183&int32(2) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L575
	}
L156:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2066&int32(1024)|v2066&int32(2) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L542
	}
L157:
	;
	if l7 != 0 {
		goto L429
	} else {
		goto L430
	}
L158:
	;
	if l7 != 0 {
		goto L397
	} else {
		goto L398
	}
L159:
	;
	v1182 = F_strlen(m, v446)
	mBase = m.M
	if l7 != 0 {
		goto L345
	} else {
		goto L346
	}
L160:
	;
	if l7 != 0 {
		goto L332
	} else {
		goto L333
	}
L161:
	;
	if l7 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v614&int32(1024) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	if base.Ui32(v468) <= base.Ui32(v482) {
		v2820 = v482
		v2824 = v486
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v498
		v2840 = v502
		v2846 = v508
		goto L147
	} else {
		goto L220
	}
L165:
	;
	v617 = int32(0)
	if v490 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v2449 = v482
	v2450 = v617
	v2451 = int32(1)
	goto L148
L167:
	;
	goto L168
L168:
	;
	v620 = v614 & int32(8)
	if v495 < v358 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v614&int32(64) != 0 {
		goto L185
	} else {
		goto L186
	}
L170:
	;
	v622 = int32(0)
	if v620 == v622 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	goto L172
L172:
	;
	if v620|base.B2i32(l3 != v486) != 0 {
		goto L169
	} else {
		goto L177
	}
L173:
	;
	v2449 = v482
	v2450 = int32(1)
	v2451 = v622
	goto L148
L174:
	;
	goto L175
L175:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v626 == v495 {
		goto L169
	} else {
		goto L176
	}
L176:
	;
	v2449 = v482
	v2450 = int32(1)
	v2451 = v622
	goto L148
L177:
	;
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v631 != int32(48) {
		goto L169
	} else {
		goto L178
	}
L178:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v634 == int32(0) {
		goto L169
	} else {
		goto L179
	}
L179:
	;
	v637 = int32(0)
	if v367 == v637 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v2449 = v482
	v2450 = int32(1)
	v2451 = v637
	goto L148
L181:
	;
	goto L182
L182:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v641 == int32(46) {
		goto L169
	} else {
		goto L183
	}
L183:
	;
	v2449 = v482
	v2450 = int32(1)
	v2451 = v637
	goto L148
L184:
	;
	v2449 = v746
	v2450 = int32(0)
	v2451 = v749
	goto L148
L185:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v648 != int32(-1) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L187
L187:
	;
	if v614&int32(128) != 0 {
		goto L212
	} else {
		goto L213
	}
L188:
	;
	v2449 = v482
	v2450 = int32(1)
	v2451 = int32(0)
	goto L148
L189:
	;
	goto L190
L190:
	;
	if (v459^v482)&int32(3) != 0 {
		goto L194
	} else {
		goto L195
	}
L191:
	;
	v727 = F_strlen(m, v482)
	mBase = m.M
	v746 = v727 + v482
	v749 = int32(1)
	goto L184
L192:
	;
	goto L191
L193:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v707))) = uint8(v706)
	if v706&int32(255) == int32(0) {
		goto L192
	} else {
		goto L208
	}
L194:
	;
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	v705 = v459
	v706 = v658
	v707 = v482
	goto L193
L195:
	;
	goto L196
L196:
	;
	if v459&int32(3) != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v662 = v459
	v664 = v482
	goto L200
L198:
	;
	v676 = v459
	v678 = v482
	goto L199
L199:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676)))
	v683 = int32(-2139062144)
	if (int32(16843008)-v680|v680)&v683 != v683 {
		v705 = v676
		v706 = v680
		v707 = v678
		goto L193
	} else {
		goto L204
	}
L200:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v662))))
	*(*uint8)(unsafe.Add(mBase, uint32(v664))) = uint8(v665)
	if v665 == int32(0) {
		goto L192
	} else {
		goto L202
	}
L201:
	;
	v676 = v672
	v678 = v670
	goto L199
L202:
	;
	v669 = int32(1)
	v670 = v664 + v669
	v672 = v662 + v669
	if v672&int32(3) != 0 {
		v662 = v672
		v664 = v670
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v688 = v676
	v689 = v680
	v690 = v678
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v690))) = v689
	v692 = int32(4)
	v693 = v690 + v692
	v695 = v688 + v692
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v688)+4))
	v700 = int32(-2139062144)
	if (int32(16843008)-v697|v697)&v700 == v700 {
		v688 = v695
		v689 = v697
		v690 = v693
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v705 = v695
	v706 = v697
	v707 = v693
	goto L193
L207:
	;
	goto L206
L208:
	;
	v714 = v705
	v716 = v707
	goto L209
L209:
	;
	v717 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v716)+1)) = uint8(v717)
	v719 = int32(1)
	if v717 != 0 {
		v714 = v714 + v719
		v716 = v716 + v719
		goto L209
	} else {
		goto L211
	}
L210:
	;
	goto L192
L211:
	;
	goto L210
L212:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v467)
	v2444 = v617
	goto L149
L213:
	;
	goto L214
L214:
	;
	switch v359 - int32(43) {
	case 0:
		goto L216
	default:
		v2449 = v482
		v2450 = int32(1)
		v2451 = int32(0)
		goto L148
	case 2:
		goto L215
	}
L215:
	;
	v741 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v741)
	v743 = int32(1)
	v746 = v482 + v743
	v749 = v743
	goto L184
L216:
	;
	v735 = int32(0)
	if v614&int32(32) != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v2449 = v482
	v2450 = v735
	v2451 = int32(1)
	goto L148
L218:
	;
	goto L219
L219:
	;
	v739 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v739)
	v2444 = v735
	goto L149
L220:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	v755 = v482 + base.B2i32(v752 == int32(32))
	if base.Ui32(v468) <= base.Ui32(v755) {
		v2820 = v755
		v2824 = v486
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v498
		v2840 = v502
		v2846 = v508
		goto L147
	} else {
		goto L221
	}
L221:
	;
	if v613&int32(-2) != int32(2) {
		v908 = v755
		goto L222
	} else {
		goto L223
	}
L222:
	;
	if base.Ui32(v468) <= base.Ui32(v908) {
		v2820 = v908
		v2824 = v486
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v498
		v2840 = v502
		v2846 = v508
		goto L147
	} else {
		goto L264
	}
L223:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v761 != int32(32))|base.B2i32(v508 != int32(0)-v502) != 0 {
		v908 = v755
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v768&int32(64) == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	v891 = int32(0)
	if base.B2i32(v886 != int32(45))&(base.B2i32(v768&int32(128) == v891)|base.B2i32(v886 != int32(60))) == v891 {
		goto L260
	} else {
		goto L261
	}
L226:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v773 != int32(-1) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v776 = F_strlen(m, v448)
	mBase = m.M
	if base.B2i32(v776 == int32(0))|base.B2i32(base.Ui32(l2+(l4-v776)) < base.Ui32(v755)) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v831 = F_strlen(m, v447)
	mBase = m.M
	if base.B2i32(v831 == int32(0))|base.B2i32(base.Ui32(l2+(l4-v831)) < base.Ui32(v755)) != 0 {
		v908 = v755
		goto L222
	} else {
		goto L244
	}
L229:
	;
	if v776 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v827 != 0 {
		goto L228
	} else {
		goto L243
	}
L231:
	;
	v827 = int32(0)
	goto L230
L232:
	;
	goto L233
L233:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	if v788 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v789 = v755
	v790 = v448
	v791 = v776
	v792 = v788
	goto L238
L235:
	;
	v815 = v448
	v819 = int32(0)
	goto L236
L236:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815))))
	v827 = v819 - v820
	goto L230
L237:
	;
	v815 = v810
	v819 = v812
	goto L236
L238:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790))))
	if base.B2i32(v792 != v794)|base.B2i32(v794 == int32(0)) != 0 {
		v810 = v790
		v812 = v792
		goto L237
	} else {
		goto L240
	}
L239:
	;
	v810 = v804
	v812 = int32(0)
	goto L237
L240:
	;
	v800 = v791 - int32(1)
	if v800 == int32(0) {
		v810 = v790
		v812 = v792
		goto L237
	} else {
		goto L241
	}
L241:
	;
	v803 = int32(1)
	v804 = v790 + v803
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v789)+1)))
	if v805 != 0 {
		v789 = v789 + v803
		v790 = v804
		v791 = v800
		v792 = v805
		goto L238
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	v828 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v828)
	v908 = v776 + v755
	goto L222
L244:
	;
	if v831 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	if v882 != 0 {
		v908 = v755
		goto L222
	} else {
		goto L258
	}
L246:
	;
	v882 = int32(0)
	goto L245
L247:
	;
	goto L248
L248:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	if v843 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v844 = v755
	v845 = v447
	v846 = v831
	v847 = v843
	goto L253
L250:
	;
	v870 = v447
	v874 = int32(0)
	goto L251
L251:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870))))
	v882 = v874 - v875
	goto L245
L252:
	;
	v870 = v865
	v874 = v867
	goto L251
L253:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845))))
	if base.B2i32(v847 != v849)|base.B2i32(v849 == int32(0)) != 0 {
		v865 = v845
		v867 = v847
		goto L252
	} else {
		goto L255
	}
L254:
	;
	v865 = v859
	v867 = int32(0)
	goto L252
L255:
	;
	v855 = v846 - int32(1)
	if v855 == int32(0) {
		v865 = v845
		v867 = v847
		goto L252
	} else {
		goto L256
	}
L256:
	;
	v858 = int32(1)
	v859 = v845 + v858
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v844)+1)))
	if v860 != 0 {
		v844 = v844 + v858
		v845 = v859
		v846 = v855
		v847 = v860
		goto L253
	} else {
		goto L257
	}
L257:
	;
	goto L254
L258:
	;
	v883 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v883)
	v908 = v831 + v755
	goto L222
L259:
	;
	v908 = v755 + int32(1)
	goto L222
L260:
	;
	v899 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v899)
	goto L259
L261:
	;
	goto L262
L262:
	;
	if v886 != int32(43) {
		v908 = v755
		goto L222
	} else {
		goto L263
	}
L263:
	;
	v903 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v903)
	goto L259
L264:
	;
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if base.Ui32((v911-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	if base.Ui32(v468) <= base.Ui32(v1008) {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L293
	}
L266:
	;
	v1004 = int32(1)
	v1008 = v1001
	v1009 = v486 + v1004
	v1010 = v1002
	v1011 = v1003
	v1012 = v508
	v1013 = v1004
	goto L265
L267:
	;
	if v498 != 0 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	goto L269
L269:
	;
	v931 = int32(0)
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if base.B2i32(v932&int32(2) == v931)|v498 != 0 {
		v1008 = v908
		v1009 = v486
		v1010 = v498
		v1011 = v502
		v1012 = v508
		v1013 = v931
		goto L265
	} else {
		goto L274
	}
L270:
	;
	v918 = int32(1)
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v502 == v919 {
		v2820 = v908
		v2824 = v486
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v918
		v2840 = v502
		v2846 = v508
		goto L147
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v486))) = uint8(v911)
	v925 = int32(1)
	v1008 = v908
	v1009 = v486 + v925
	v1010 = int32(0)
	v1011 = v502
	v1012 = v508 + v925
	v1013 = v925
	goto L265
L273:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v486))) = uint8(v911)
	v1001 = v908
	v1002 = v918
	v1003 = v502 + int32(1)
	goto L266
L274:
	;
	v938 = F_strlen(m, v449)
	mBase = m.M
	if v938 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v941 = int32(0)
	v1008 = v908
	v1009 = v486
	v1010 = v941
	v1011 = v502
	v1012 = v508
	v1013 = v941
	goto L265
L276:
	;
	goto L277
L277:
	;
	v943 = int32(0)
	if base.Ui32(l2+(l4-v938)) < base.Ui32(v908) {
		v1008 = v908
		v1009 = v486
		v1010 = v943
		v1011 = v502
		v1012 = v508
		v1013 = v943
		goto L265
	} else {
		goto L278
	}
L278:
	;
	v948 = int32(0)
	if v938 == v948 {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	if v993 != 0 {
		v1008 = v908
		v1009 = v486
		v1010 = v943
		v1011 = v502
		v1012 = v508
		v1013 = v948
		goto L265
	} else {
		goto L292
	}
L280:
	;
	v993 = int32(0)
	goto L279
L281:
	;
	goto L282
L282:
	;
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	if v954 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v955 = v908
	v956 = v449
	v957 = v938
	v958 = v954
	goto L287
L284:
	;
	v981 = v449
	v985 = int32(0)
	goto L285
L285:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v981))))
	v993 = v985 - v986
	goto L279
L286:
	;
	v981 = v976
	v985 = v978
	goto L285
L287:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	if base.B2i32(v958 != v960)|base.B2i32(v960 == int32(0)) != 0 {
		v976 = v956
		v978 = v958
		goto L286
	} else {
		goto L289
	}
L288:
	;
	v976 = v970
	v978 = int32(0)
	goto L286
L289:
	;
	v966 = v957 - int32(1)
	if v966 == int32(0) {
		v976 = v956
		v978 = v958
		goto L286
	} else {
		goto L290
	}
L290:
	;
	v969 = int32(1)
	v970 = v956 + v969
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955)+1)))
	if v971 != 0 {
		v955 = v955 + v969
		v956 = v970
		v957 = v966
		v958 = v971
		goto L287
	} else {
		goto L291
	}
L291:
	;
	goto L288
L292:
	;
	v994 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v486))) = uint8(v994)
	v996 = int32(1)
	v1001 = v938 + v908 - v996
	v1002 = v996
	v1003 = v502
	goto L266
L293:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v1015 != int32(32))|base.B2i32(v1011+v1012 <= int32(0)) != 0 {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L294
	}
L294:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1024 = v1022 & int32(64)
	v1025 = int32(0)
	if base.B2i32(v1024 == v1025)|(v1013^int32(1)) == v1025 {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1033 = v1008 + int32(1)
	if base.Ui32(v468) <= base.Ui32(v1033) {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	if v1013|(v1024|base.B2i32(v1022&int32(768) == int32(0))) != 0 {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L331
	}
L298:
	;
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	if base.Ui32((v1035-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L299
	}
L299:
	;
	v1042 = F_strlen(m, v448)
	mBase = m.M
	if base.B2i32(v1042 == int32(0))|base.B2i32(base.Ui32(l2+(l4-v1042)) < base.Ui32(v1033)) != 0 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1097 = F_strlen(m, v447)
	mBase = m.M
	if base.B2i32(v1097 == int32(0))|base.B2i32(base.Ui32(l2+(l4-v1097)) < base.Ui32(v1033)) != 0 {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L316
	}
L301:
	;
	if v1042 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	if v1093 != 0 {
		goto L300
	} else {
		goto L315
	}
L303:
	;
	v1093 = int32(0)
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	if v1054 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1055 = v1033
	v1056 = v448
	v1057 = v1042
	v1058 = v1054
	goto L310
L307:
	;
	v1081 = v448
	v1085 = int32(0)
	goto L308
L308:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081))))
	v1093 = v1085 - v1086
	goto L302
L309:
	;
	v1081 = v1076
	v1085 = v1078
	goto L308
L310:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1056))))
	if base.B2i32(v1058 != v1060)|base.B2i32(v1060 == int32(0)) != 0 {
		v1076 = v1056
		v1078 = v1058
		goto L309
	} else {
		goto L312
	}
L311:
	;
	v1076 = v1070
	v1078 = int32(0)
	goto L309
L312:
	;
	v1066 = v1057 - int32(1)
	if v1066 == int32(0) {
		v1076 = v1056
		v1078 = v1058
		goto L309
	} else {
		goto L313
	}
L313:
	;
	v1069 = int32(1)
	v1070 = v1056 + v1069
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1055)+1)))
	if v1071 != 0 {
		v1055 = v1055 + v1069
		v1056 = v1070
		v1057 = v1066
		v1058 = v1071
		goto L310
	} else {
		goto L314
	}
L314:
	;
	goto L311
L315:
	;
	v1095 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1095)
	v2820 = v1008 + v1042
	v2824 = v1009
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v1010
	v2840 = v1011
	v2846 = v1012
	goto L147
L316:
	;
	if v1097 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L317:
	;
	if v1148 != 0 {
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	} else {
		goto L330
	}
L318:
	;
	v1148 = int32(0)
	goto L317
L319:
	;
	goto L320
L320:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	if v1109 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1110 = v1033
	v1111 = v447
	v1112 = v1097
	v1113 = v1109
	goto L325
L322:
	;
	v1136 = v447
	v1140 = int32(0)
	goto L323
L323:
	;
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1136))))
	v1148 = v1140 - v1141
	goto L317
L324:
	;
	v1136 = v1131
	v1140 = v1133
	goto L323
L325:
	;
	v1115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111))))
	if base.B2i32(v1113 != v1115)|base.B2i32(v1115 == int32(0)) != 0 {
		v1131 = v1111
		v1133 = v1113
		goto L324
	} else {
		goto L327
	}
L326:
	;
	v1131 = v1125
	v1133 = int32(0)
	goto L324
L327:
	;
	v1121 = v1112 - int32(1)
	if v1121 == int32(0) {
		v1131 = v1111
		v1133 = v1113
		goto L324
	} else {
		goto L328
	}
L328:
	;
	v1124 = int32(1)
	v1125 = v1111 + v1124
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+1)))
	if v1126 != 0 {
		v1110 = v1110 + v1124
		v1111 = v1125
		v1112 = v1121
		v1113 = v1126
		goto L325
	} else {
		goto L329
	}
L329:
	;
	goto L326
L330:
	;
	v1152 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v1152)
	v2820 = v1033 + v1097 - int32(1)
	v2824 = v1009
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v1010
	v2840 = v1011
	v2846 = v1012
	goto L147
L331:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008))))
	switch v1160 - int32(43) {
	case 0, 2:
		goto L151
	default:
		v2820 = v1008
		v2824 = v1009
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v1010
		v2840 = v1011
		v2846 = v1012
		goto L147
	}
L332:
	;
	if v487 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L333:
	;
	goto L334
L334:
	;
	if v487 != 0 {
		goto L341
	} else {
		goto L342
	}
L335:
	;
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1165&int32(32) != 0 {
		goto L338
	} else {
		goto L339
	}
L336:
	;
	goto L337
L337:
	;
	v1171 = int32(44)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v1171)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L338:
	;
	v2868 = v482
	v2872 = v486
	v2873 = int32(0)
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L339:
	;
	goto L340
L340:
	;
	v1169 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v1169)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L341:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v1179 == int32(44) {
		v2820 = v482
		v2824 = v486
		v2825 = v487
		v2828 = v490
		v2833 = v495
		v2836 = v498
		v2840 = v502
		v2846 = v508
		goto L147
	} else {
		goto L344
	}
L342:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1173&int32(32) == int32(0) {
		goto L341
	} else {
		goto L343
	}
L343:
	;
	v2868 = v482
	v2872 = v486
	v2873 = int32(0)
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L344:
	;
	v2868 = v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L345:
	;
	if v487 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	goto L347
L347:
	;
	if v487 != 0 {
		goto L379
	} else {
		goto L380
	}
L348:
	;
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1185&int32(32) != 0 {
		goto L351
	} else {
		goto L352
	}
L349:
	;
	goto L350
L350:
	;
	if (v446^v482)&int32(3) != 0 {
		goto L361
	} else {
		goto L362
	}
L351:
	;
	v2868 = v482
	v2872 = v486
	v2873 = int32(0)
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L352:
	;
	goto L353
L353:
	;
	v1189 = F_pg_mbstrlen(m, v446)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L70
	} else {
		goto L354
	}
L354:
	;
	if v1189 != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	base.MemoryFill(m, v482, int32(32), v1189)
	goto L357
L356:
	;
	goto L357
L357:
	;
	v2820 = v1189 + v482 - int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L358:
	;
	v2820 = v482 + v1182 - int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L359:
	;
	goto L358
L360:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1250))) = uint8(v1249)
	if v1249&int32(255) == int32(0) {
		goto L359
	} else {
		goto L375
	}
L361:
	;
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	v1248 = v446
	v1249 = v1201
	v1250 = v482
	goto L360
L362:
	;
	goto L363
L363:
	;
	if v446&int32(3) != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1205 = v446
	v1207 = v482
	goto L367
L365:
	;
	v1219 = v446
	v1221 = v482
	goto L366
L366:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	v1226 = int32(-2139062144)
	if (int32(16843008)-v1223|v1223)&v1226 != v1226 {
		v1248 = v1219
		v1249 = v1223
		v1250 = v1221
		goto L360
	} else {
		goto L371
	}
L367:
	;
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1207))) = uint8(v1208)
	if v1208 == int32(0) {
		goto L359
	} else {
		goto L369
	}
L368:
	;
	v1219 = v1215
	v1221 = v1213
	goto L366
L369:
	;
	v1212 = int32(1)
	v1213 = v1207 + v1212
	v1215 = v1205 + v1212
	if v1215&int32(3) != 0 {
		v1205 = v1215
		v1207 = v1213
		goto L367
	} else {
		goto L370
	}
L370:
	;
	goto L368
L371:
	;
	v1231 = v1219
	v1232 = v1223
	v1233 = v1221
	goto L372
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1233))) = v1232
	v1235 = int32(4)
	v1236 = v1233 + v1235
	v1238 = v1231 + v1235
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+4))
	v1243 = int32(-2139062144)
	if (int32(16843008)-v1240|v1240)&v1243 == v1243 {
		v1231 = v1238
		v1232 = v1240
		v1233 = v1236
		goto L372
	} else {
		goto L374
	}
L373:
	;
	v1248 = v1238
	v1249 = v1240
	v1250 = v1236
	goto L360
L374:
	;
	goto L373
L375:
	;
	v1257 = v1248
	v1259 = v1250
	goto L376
L376:
	;
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1257)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1259)+1)) = uint8(v1260)
	v1262 = int32(1)
	if v1260 != 0 {
		v1257 = v1257 + v1262
		v1259 = v1259 + v1262
		goto L376
	} else {
		goto L378
	}
L377:
	;
	goto L359
L378:
	;
	goto L377
L379:
	;
	if base.Ui32(l2+(l4-v1182)) < base.Ui32(v482) {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L382
	}
L380:
	;
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1273&int32(32) == int32(0) {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2868 = v482
	v2872 = v486
	v2873 = int32(0)
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L382:
	;
	if v1182 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v1326 != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L396
	}
L384:
	;
	v1326 = int32(0)
	goto L383
L385:
	;
	goto L386
L386:
	;
	v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v1287 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1288 = v482
	v1289 = v446
	v1290 = v1182
	v1291 = v1287
	goto L391
L388:
	;
	v1314 = v446
	v1318 = int32(0)
	goto L389
L389:
	;
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314))))
	v1326 = v1318 - v1319
	goto L383
L390:
	;
	v1314 = v1309
	v1318 = v1311
	goto L389
L391:
	;
	v1293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289))))
	if base.B2i32(v1291 != v1293)|base.B2i32(v1293 == int32(0)) != 0 {
		v1309 = v1289
		v1311 = v1291
		goto L390
	} else {
		goto L393
	}
L392:
	;
	v1309 = v1303
	v1311 = int32(0)
	goto L390
L393:
	;
	v1299 = v1290 - int32(1)
	if v1299 == int32(0) {
		v1309 = v1289
		v1311 = v1291
		goto L390
	} else {
		goto L394
	}
L394:
	;
	v1302 = int32(1)
	v1303 = v1289 + v1302
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288)+1)))
	if v1304 != 0 {
		v1288 = v1288 + v1302
		v1289 = v1303
		v1290 = v1299
		v1291 = v1304
		goto L391
	} else {
		goto L395
	}
L395:
	;
	goto L392
L396:
	;
	v2820 = v1182 + v482 - int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L397:
	;
	if (v450^v482)&int32(3) != 0 {
		goto L403
	} else {
		goto L404
	}
L398:
	;
	goto L399
L399:
	;
	v1408 = F_pg_mbstrlen(m, v450)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L70
	} else {
		goto L421
	}
L400:
	;
	v1404 = F_strlen(m, v450)
	mBase = m.M
	v2820 = v482 + v1404 - int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L401:
	;
	goto L400
L402:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1384))) = uint8(v1383)
	if v1383&int32(255) == int32(0) {
		goto L401
	} else {
		goto L417
	}
L403:
	;
	v1335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450))))
	v1382 = v450
	v1383 = v1335
	v1384 = v482
	goto L402
L404:
	;
	goto L405
L405:
	;
	if v450&int32(3) != 0 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1339 = v450
	v1341 = v482
	goto L409
L407:
	;
	v1353 = v450
	v1355 = v482
	goto L408
L408:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v1353)))
	v1360 = int32(-2139062144)
	if (int32(16843008)-v1357|v1357)&v1360 != v1360 {
		v1382 = v1353
		v1383 = v1357
		v1384 = v1355
		goto L402
	} else {
		goto L413
	}
L409:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1341))) = uint8(v1342)
	if v1342 == int32(0) {
		goto L401
	} else {
		goto L411
	}
L410:
	;
	v1353 = v1349
	v1355 = v1347
	goto L408
L411:
	;
	v1346 = int32(1)
	v1347 = v1341 + v1346
	v1349 = v1339 + v1346
	if v1349&int32(3) != 0 {
		v1339 = v1349
		v1341 = v1347
		goto L409
	} else {
		goto L412
	}
L412:
	;
	goto L410
L413:
	;
	v1365 = v1353
	v1366 = v1357
	v1367 = v1355
	goto L414
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1367))) = v1366
	v1369 = int32(4)
	v1370 = v1367 + v1369
	v1372 = v1365 + v1369
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+4))
	v1377 = int32(-2139062144)
	if (int32(16843008)-v1374|v1374)&v1377 == v1377 {
		v1365 = v1372
		v1366 = v1374
		v1367 = v1370
		goto L414
	} else {
		goto L416
	}
L415:
	;
	v1382 = v1372
	v1383 = v1374
	v1384 = v1370
	goto L402
L416:
	;
	goto L415
L417:
	;
	v1391 = v1382
	v1393 = v1384
	goto L418
L418:
	;
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1393)+1)) = uint8(v1394)
	v1396 = int32(1)
	if v1394 != 0 {
		v1391 = v1391 + v1396
		v1393 = v1393 + v1396
		goto L418
	} else {
		goto L420
	}
L419:
	;
	goto L401
L420:
	;
	goto L419
L421:
	;
	if v1408 <= int32(0) {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L422
	}
L422:
	;
	v1412 = v1408
	v1420 = v482
	goto L423
L423:
	;
	if base.Ui32(v468) <= base.Ui32(v1420) {
		v2868 = v1420
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L425
	}
L424:
	;
	v2868 = v1471
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L425:
	;
	v1459 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1420))))
	if base.B2i32(base.Ui64(v1459) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v1459)%64)&int64(288080842570334209) != int64(0)) != 0 {
		v2868 = v1420
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L426
	}
L426:
	;
	v1469 = F_pg_mblen_range(m, v1420, v468)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L70
	} else {
		goto L427
	}
L427:
	;
	v1471 = v1469 + v1420
	v1472 = int32(1)
	if base.Ui32(v1472) < base.Ui32(v1412) {
		v1412 = v1412 - v1472
		v1420 = v1471
		goto L423
	} else {
		goto L428
	}
L428:
	;
	goto L424
L429:
	;
	if v613 != int32(30) {
		v1554 = v486
		goto L432
	} else {
		goto L433
	}
L430:
	;
	goto L431
L431:
	;
	if base.Ui32(v468) <= base.Ui32(v482) {
		v1744 = v482
		v1782 = v482
		goto L469
	} else {
		goto L470
	}
L432:
	;
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v1591&int32(32) != 0 {
		goto L444
	} else {
		goto L445
	}
L433:
	;
	v1478 = F_strlen(m, v486)
	mBase = m.M
	v1479 = F_pnstrdup(m, v486, v1478)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L70
	} else {
		goto L434
	}
L434:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1479))))
	if v1481 == int32(0) {
		v1554 = v1479
		goto L432
	} else {
		goto L435
	}
L435:
	;
	v1484 = v1479
	v1494 = v1481
	goto L436
L436:
	;
	v1530 = int32(255)
	v1531 = v1494 & v1530
	if base.Ui32((v1531-int32(65))&v1530) < base.Ui32(int32(26)) {
		goto L439
	} else {
		goto L440
	}
L437:
	;
	v1554 = v1479
	goto L432
L438:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1484))) = uint8(v1540)
	v1542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484)+1)))
	if v1542 != 0 {
		v1484 = v1484 + int32(1)
		v1494 = v1542
		goto L436
	} else {
		goto L442
	}
L439:
	;
	v1540 = v1531 | int32(32)
	goto L441
L440:
	;
	v1540 = v1531
	goto L441
L441:
	;
	goto L438
L442:
	;
	goto L437
L443:
	;
	v1672 = F_strlen(m, v482)
	mBase = m.M
	v2820 = v1672 + v482 - int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L444:
	;
	if (v1554^v482)&int32(3) != 0 {
		goto L450
	} else {
		goto L451
	}
L445:
	;
	goto L446
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v1554
	v1670 = F_pg_sprintf(m, v482, int32(_a_F_NUM_processor_6), v49)
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L70
	} else {
		goto L468
	}
L447:
	;
	goto L443
L448:
	;
	goto L447
L449:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1648))) = uint8(v1647)
	if v1647&int32(255) == int32(0) {
		goto L448
	} else {
		goto L464
	}
L450:
	;
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1554))))
	v1646 = v1554
	v1647 = v1599
	v1648 = v482
	goto L449
L451:
	;
	goto L452
L452:
	;
	if v1554&int32(3) != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1603 = v1554
	v1605 = v482
	goto L456
L454:
	;
	v1617 = v1554
	v1619 = v482
	goto L455
L455:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1617)))
	v1624 = int32(-2139062144)
	if (int32(16843008)-v1621|v1621)&v1624 != v1624 {
		v1646 = v1617
		v1647 = v1621
		v1648 = v1619
		goto L449
	} else {
		goto L460
	}
L456:
	;
	v1606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1603))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1605))) = uint8(v1606)
	if v1606 == int32(0) {
		goto L448
	} else {
		goto L458
	}
L457:
	;
	v1617 = v1613
	v1619 = v1611
	goto L455
L458:
	;
	v1610 = int32(1)
	v1611 = v1605 + v1610
	v1613 = v1603 + v1610
	if v1613&int32(3) != 0 {
		v1603 = v1613
		v1605 = v1611
		goto L456
	} else {
		goto L459
	}
L459:
	;
	goto L457
L460:
	;
	v1629 = v1617
	v1630 = v1621
	v1631 = v1619
	goto L461
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1631))) = v1630
	v1633 = int32(4)
	v1634 = v1631 + v1633
	v1636 = v1629 + v1633
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1629)+4))
	v1641 = int32(-2139062144)
	if (int32(16843008)-v1638|v1638)&v1641 == v1641 {
		v1629 = v1636
		v1630 = v1638
		v1631 = v1634
		goto L461
	} else {
		goto L463
	}
L462:
	;
	v1646 = v1636
	v1647 = v1638
	v1648 = v1634
	goto L449
L463:
	;
	goto L462
L464:
	;
	v1655 = v1646
	v1657 = v1648
	goto L465
L465:
	;
	v1658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1655)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1657)+1)) = uint8(v1658)
	v1660 = int32(1)
	if v1658 != 0 {
		v1655 = v1655 + v1660
		v1657 = v1657 + v1660
		goto L465
	} else {
		goto L467
	}
L466:
	;
	goto L448
L467:
	;
	goto L466
L468:
	;
	goto L443
L469:
	;
	v1784 = int32(0)
	v1792 = v1744
	v1794 = v1782
	goto L476
L470:
	;
	v1685 = v482
	goto L471
L471:
	;
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1685))))
	if base.B2i32(base.Ui32(v1723-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v1723 == int32(32)) == int32(0) {
		v1744 = v1685
		v1782 = v1685
		goto L469
	} else {
		goto L473
	}
L472:
	;
	v1744 = v1734
	v1782 = v468
	goto L469
L473:
	;
	v1734 = v1685 + int32(1)
	if v1734 != v468 {
		v1685 = v1734
		goto L471
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	v1879 = int32(1)
	v1881 = int32(0)
	v1889 = v1881
	v1912 = v1879
	v1914 = v1881
	v1915 = v1881
	v1916 = v1881
	v1921 = v1881
	v1927 = v1881
	v1928 = v1881
	goto L493
L476:
	;
	if base.Ui32(v468) <= base.Ui32(v1794) {
		goto L478
	} else {
		goto L479
	}
L477:
	;
	if v1784 == int32(0) {
		goto L150
	} else {
		goto L492
	}
L478:
	;
	goto L477
L479:
	;
	v1832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1794))))
	if base.Ui32((v1832-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L488
	} else {
		goto L489
	}
L480:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v1784))) = uint8(v1843)
	*(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v1784<<(uint(int32(2))%32)))) = v1852
	v1863 = int32(15)
	v1864 = int32(1)
	v1865 = v1792 + v1864
	v1867 = v1784 + v1864
	if v1867 != v1863 {
		v1784 = v1867
		v1792 = v1865
		v1794 = v1865
		goto L476
	} else {
		goto L491
	}
L481:
	;
	v1852 = int32(500)
	goto L480
L482:
	;
	v1852 = int32(100)
	goto L480
L483:
	;
	v1852 = int32(50)
	goto L480
L484:
	;
	v1852 = int32(10)
	goto L480
L485:
	;
	v1852 = int32(5)
	goto L480
L486:
	;
	v1852 = int32(1000)
	goto L480
L487:
	;
	switch v1843 - int32(67) {
	case 0:
		goto L482
	case 1:
		goto L481
	default:
		goto L478
	case 6:
		v1852 = int32(1)
		goto L480
	case 9:
		goto L483
	case 10:
		goto L486
	case 19:
		goto L485
	case 21:
		goto L484
	}
L488:
	;
	v1841 = v1832 - int32(32)
	goto L490
L489:
	;
	v1841 = v1832
	goto L490
L490:
	;
	v1843 = v1841 & int32(255)
	goto L487
L491:
	;
	v1875 = v1865
	v1878 = v1863
	goto L475
L492:
	;
	v1875 = v1792
	v1878 = v1784
	goto L475
L493:
	;
	v1938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(97)+v1889))))
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v1889<<(uint(int32(2))%32))))
	if int32(4) < v1944 {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	if v2050 < int32(0) {
		goto L150
	} else {
		goto L540
	}
L495:
	;
	v1950 = v1914
	goto L497
L496:
	;
	v1950 = int32(0)
	goto L497
L497:
	;
	if int32(49) < v1944 {
		goto L498
	} else {
		goto L499
	}
L498:
	;
	v1955 = v1915
	goto L500
L499:
	;
	v1955 = int32(0)
	goto L500
L500:
	;
	if int32(499) < v1944 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1959 = v1916
	goto L503
L502:
	;
	v1959 = int32(0)
	goto L503
L503:
	;
	if base.B2i32(v1927 <= v1944)&v1928|v1950|(v1955|v1959) != 0 {
		goto L150
	} else {
		goto L504
	}
L504:
	;
	switch v1938 - int32(68) {
	case 0:
		goto L506
	default:
		v1970 = v1914
		v1971 = v1915
		v1972 = v1916
		goto L505
	case 8:
		goto L507
	case 18:
		goto L508
	}
L505:
	;
	if v1889 < v1878-v1879 {
		goto L511
	} else {
		goto L512
	}
L506:
	;
	v1970 = v1914
	v1971 = v1915
	v1972 = v1916 + int32(1)
	goto L505
L507:
	;
	v1970 = v1914
	v1971 = v1915 + int32(1)
	v1972 = v1916
	goto L505
L508:
	;
	v1970 = v1914 + int32(1)
	v1971 = v1915
	v1972 = v1916
	goto L505
L509:
	;
	v2050 = v2049 + v1921
	v2052 = v2040 + int32(1)
	if v2052 < v1878 {
		v1889 = v2052
		v1912 = v2041
		v1914 = v2042
		v1915 = v2043
		v1916 = v2044
		v1921 = v2050
		v1927 = v2047
		v1928 = v2048
		goto L493
	} else {
		goto L539
	}
L510:
	;
	v2040 = v1889
	v2041 = v2037
	v2042 = v1970
	v2043 = v1971
	v2044 = v1972
	v2047 = v1927
	v2048 = v1928
	v2049 = v1944
	goto L509
L511:
	;
	v1975 = v1889 + int32(1)
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1975+(v49+int32(97))))))
	v1985 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(32)+v1975<<(uint(int32(2))%32))))
	if v1944 < v1985 {
		goto L514
	} else {
		goto L515
	}
L512:
	;
	v2033 = v1912
	goto L513
L513:
	;
	v2037 = v2033
	goto L510
L514:
	;
	switch v1938 - int32(67) {
	case 0:
		goto L518
	default:
		goto L150
	case 6:
		goto L520
	case 21:
		goto L519
	}
L515:
	;
	goto L516
L516:
	;
	if v1979 != v1938 {
		goto L535
	} else {
		goto L536
	}
L517:
	;
	if int32(4) < v1985 {
		goto L521
	} else {
		goto L522
	}
L518:
	;
	switch v1979 - int32(68) {
	case 0, 9:
		goto L517
	default:
		goto L150
	}
L519:
	;
	switch v1979 - int32(67) {
	case 0, 9:
		goto L517
	default:
		goto L150
	}
L520:
	;
	switch v1979 - int32(86) {
	case 0, 2:
		goto L517
	default:
		goto L150
	}
L521:
	;
	v2000 = v1970
	goto L523
L522:
	;
	v2000 = int32(0)
	goto L523
L523:
	;
	if int32(49) < v1985 {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2005 = v1971
	goto L526
L525:
	;
	v2005 = int32(0)
	goto L526
L526:
	;
	if int32(499) < v1985 {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v2009 = v1972
	goto L529
L528:
	;
	v2009 = int32(0)
	goto L529
L529:
	;
	if base.B2i32(int32(1) < v1912)|v2000|(v2005|v2009) != 0 {
		goto L150
	} else {
		goto L530
	}
L530:
	;
	switch v1979 - int32(68) {
	case 0:
		goto L532
	default:
		v2020 = v1970
		v2021 = v1971
		v2022 = v1972
		goto L531
	case 8:
		goto L533
	case 18:
		goto L534
	}
L531:
	;
	v2023 = int32(1)
	v2040 = v1975
	v2041 = v2023
	v2042 = v2020
	v2043 = v2021
	v2044 = v2022
	v2047 = v1944
	v2048 = v2023
	v2049 = v1985 - v1944
	goto L509
L532:
	;
	v2020 = v1970
	v2021 = v1971
	v2022 = v1972 + int32(1)
	goto L531
L533:
	;
	v2020 = v1970
	v2021 = v1971 + int32(1)
	v2022 = v1972
	goto L531
L534:
	;
	v2020 = v1970 + int32(1)
	v2021 = v1971
	v2022 = v1972
	goto L531
L535:
	;
	v2037 = int32(1)
	goto L510
L536:
	;
	goto L537
L537:
	;
	if int32(2) < v1912 {
		goto L150
	} else {
		goto L538
	}
L538:
	;
	v2033 = v1912 + int32(1)
	goto L513
L539:
	;
	goto L494
L540:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v2050
	v2060 = F_pg_sprintf(m, v486, int32(_a_F_NUM_processor_7), v49+int32(16))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L70
	} else {
		goto L541
	}
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v2060
	v2868 = v1875
	v2872 = v2060 + v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L542:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v2072 == int32(35))|v458 != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L543
	}
L543:
	;
	if l7 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2077 = F_get_th(m, l3, int32(2))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L70
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	if base.Ui32(v468) <= base.Ui32(v482) {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L569
	}
L547:
	;
	if (v2077^v482)&int32(3) != 0 {
		goto L551
	} else {
		goto L552
	}
L548:
	;
	v2820 = v482 + int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L549:
	;
	goto L548
L550:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2133))) = uint8(v2132)
	if v2132&int32(255) == int32(0) {
		goto L549
	} else {
		goto L565
	}
L551:
	;
	v2084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077))))
	v2131 = v2077
	v2132 = v2084
	v2133 = v482
	goto L550
L552:
	;
	goto L553
L553:
	;
	if v2077&int32(3) != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v2088 = v2077
	v2090 = v482
	goto L557
L555:
	;
	v2102 = v2077
	v2104 = v482
	goto L556
L556:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2102)))
	v2109 = int32(-2139062144)
	if (int32(16843008)-v2106|v2106)&v2109 != v2109 {
		v2131 = v2102
		v2132 = v2106
		v2133 = v2104
		goto L550
	} else {
		goto L561
	}
L557:
	;
	v2091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2088))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2090))) = uint8(v2091)
	if v2091 == int32(0) {
		goto L549
	} else {
		goto L559
	}
L558:
	;
	v2102 = v2098
	v2104 = v2096
	goto L556
L559:
	;
	v2095 = int32(1)
	v2096 = v2090 + v2095
	v2098 = v2088 + v2095
	if v2098&int32(3) != 0 {
		v2088 = v2098
		v2090 = v2096
		goto L557
	} else {
		goto L560
	}
L560:
	;
	goto L558
L561:
	;
	v2114 = v2102
	v2115 = v2106
	v2116 = v2104
	goto L562
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2115
	v2118 = int32(4)
	v2119 = v2116 + v2118
	v2121 = v2114 + v2118
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v2114)+4))
	v2126 = int32(-2139062144)
	if (int32(16843008)-v2123|v2123)&v2126 == v2126 {
		v2114 = v2121
		v2115 = v2123
		v2116 = v2119
		goto L562
	} else {
		goto L564
	}
L563:
	;
	v2131 = v2121
	v2132 = v2123
	v2133 = v2119
	goto L550
L564:
	;
	goto L563
L565:
	;
	v2140 = v2131
	v2142 = v2133
	goto L566
L566:
	;
	v2143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2142)+1)) = uint8(v2143)
	v2145 = int32(1)
	if v2143 != 0 {
		v2140 = v2140 + v2145
		v2142 = v2142 + v2145
		goto L566
	} else {
		goto L568
	}
L567:
	;
	goto L549
L568:
	;
	goto L567
L569:
	;
	v2156 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if base.B2i32(base.Ui64(v2156) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2156)%64)&int64(288080842570334209) != int64(0)) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L570
	}
L570:
	;
	v2166 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L70
	} else {
		goto L571
	}
L571:
	;
	v2168 = v2166 + v482
	if base.Ui32(v468) <= base.Ui32(v2168) {
		v2868 = v2168
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L572
	}
L572:
	;
	v2170 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2168))))
	if base.B2i32(base.Ui64(v2170) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2170)%64)&int64(288080842570334209) != int64(0)) != 0 {
		v2868 = v2168
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L573
	}
L573:
	;
	v2180 = F_pg_mblen_range(m, v2168, v468)
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L70
	} else {
		goto L574
	}
L574:
	;
	v2868 = v2180 + v2168
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L575:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if base.B2i32(v2189 == int32(35))|v458 != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L576
	}
L576:
	;
	if l7 != 0 {
		goto L577
	} else {
		goto L578
	}
L577:
	;
	v2194 = F_get_th(m, l3, int32(1))
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L70
	} else {
		goto L580
	}
L578:
	;
	goto L579
L579:
	;
	if base.Ui32(v468) <= base.Ui32(v482) {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L602
	}
L580:
	;
	if (v2194^v482)&int32(3) != 0 {
		goto L584
	} else {
		goto L585
	}
L581:
	;
	v2820 = v482 + int32(1)
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L582:
	;
	goto L581
L583:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2250))) = uint8(v2249)
	if v2249&int32(255) == int32(0) {
		goto L582
	} else {
		goto L598
	}
L584:
	;
	v2201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2194))))
	v2248 = v2194
	v2249 = v2201
	v2250 = v482
	goto L583
L585:
	;
	goto L586
L586:
	;
	if v2194&int32(3) != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v2205 = v2194
	v2207 = v482
	goto L590
L588:
	;
	v2219 = v2194
	v2221 = v482
	goto L589
L589:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(v2219)))
	v2226 = int32(-2139062144)
	if (int32(16843008)-v2223|v2223)&v2226 != v2226 {
		v2248 = v2219
		v2249 = v2223
		v2250 = v2221
		goto L583
	} else {
		goto L594
	}
L590:
	;
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2205))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2207))) = uint8(v2208)
	if v2208 == int32(0) {
		goto L582
	} else {
		goto L592
	}
L591:
	;
	v2219 = v2215
	v2221 = v2213
	goto L589
L592:
	;
	v2212 = int32(1)
	v2213 = v2207 + v2212
	v2215 = v2205 + v2212
	if v2215&int32(3) != 0 {
		v2205 = v2215
		v2207 = v2213
		goto L590
	} else {
		goto L593
	}
L593:
	;
	goto L591
L594:
	;
	v2231 = v2219
	v2232 = v2223
	v2233 = v2221
	goto L595
L595:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2233))) = v2232
	v2235 = int32(4)
	v2236 = v2233 + v2235
	v2238 = v2231 + v2235
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+4))
	v2243 = int32(-2139062144)
	if (int32(16843008)-v2240|v2240)&v2243 == v2243 {
		v2231 = v2238
		v2232 = v2240
		v2233 = v2236
		goto L595
	} else {
		goto L597
	}
L596:
	;
	v2248 = v2238
	v2249 = v2240
	v2250 = v2236
	goto L583
L597:
	;
	goto L596
L598:
	;
	v2257 = v2248
	v2259 = v2250
	goto L599
L599:
	;
	v2260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2257)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2259)+1)) = uint8(v2260)
	v2262 = int32(1)
	if v2260 != 0 {
		v2257 = v2257 + v2262
		v2259 = v2259 + v2262
		goto L599
	} else {
		goto L601
	}
L600:
	;
	goto L582
L601:
	;
	goto L600
L602:
	;
	v2273 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if base.B2i32(base.Ui64(v2273) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2273)%64)&int64(288080842570334209) != int64(0)) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L603
	}
L603:
	;
	v2283 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L70
	} else {
		goto L604
	}
L604:
	;
	v2285 = v2283 + v482
	if base.Ui32(v468) <= base.Ui32(v2285) {
		v2868 = v2285
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L605
	}
L605:
	;
	v2287 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v2285))))
	if base.B2i32(base.Ui64(v2287) <= base.Ui64(int64(63)))&base.B2i32(int64(1)<<(uint(v2287)%64)&int64(288080842570334209) != int64(0)) != 0 {
		v2868 = v2285
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L606
	}
L606:
	;
	v2297 = F_pg_mblen_range(m, v2285, v468)
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L70
	} else {
		goto L607
	}
L607:
	;
	v2868 = v2297 + v2285
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L608:
	;
	if v359 == int32(45) {
		goto L611
	} else {
		goto L612
	}
L609:
	;
	goto L610
L610:
	;
	v2309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v2309 == int32(45) {
		goto L615
	} else {
		goto L616
	}
L611:
	;
	v2302 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v2302)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L612:
	;
	goto L613
L613:
	;
	v2304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2304&int32(32) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L614
	}
L614:
	;
	v2307 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v2307)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L615:
	;
	v2312 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2312)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L616:
	;
	goto L617
L617:
	;
	if base.B2i32(base.Ui32(v2309) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2309))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v468) <= base.Ui32(v482)) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L618
	}
L618:
	;
	v2326 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L70
	} else {
		goto L619
	}
L619:
	;
	v2868 = v2326 + v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L620:
	;
	if v359 == int32(43) {
		goto L623
	} else {
		goto L624
	}
L621:
	;
	goto L622
L622:
	;
	v2338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	if v2338 == int32(43) {
		goto L627
	} else {
		goto L628
	}
L623:
	;
	v2331 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v2331)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L624:
	;
	goto L625
L625:
	;
	v2333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+12)))
	if v2333&int32(32) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L626
	}
L626:
	;
	v2336 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v2336)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L627:
	;
	v2341 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2341)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L628:
	;
	goto L629
L629:
	;
	if base.B2i32(base.Ui32(v2338) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2338))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v468) <= base.Ui32(v482)) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L630
	}
L630:
	;
	v2355 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L70
	} else {
		goto L631
	}
L631:
	;
	v2868 = v2355 + v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L632:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v359)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L633:
	;
	goto L634
L634:
	;
	v2359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482))))
	switch v2359 - int32(43) {
	case 0:
		goto L636
	default:
		goto L635
	case 2:
		goto L637
	}
L635:
	;
	if base.B2i32(base.Ui32(v2359) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v2359))%64)&int64(288080842570334209) != int64(0))|base.B2i32(base.Ui32(v468) <= base.Ui32(v482)) != 0 {
		v2868 = v482
		v2872 = v486
		v2873 = v487
		v2876 = v490
		v2881 = v495
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L638
	}
L636:
	;
	v2364 = int32(43)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2364)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L637:
	;
	v2362 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v2362)
	v2820 = v482
	v2824 = v486
	v2825 = v487
	v2828 = v490
	v2833 = v495
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L638:
	;
	v2378 = F_pg_mblen_range(m, v482, v468)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L70
	} else {
		goto L639
	}
L639:
	;
	v2868 = v2378 + v482
	v2872 = v486
	v2873 = v487
	v2876 = v490
	v2881 = v495
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L640:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v2434 = m.ExcPending
	if v2434 != 0 {
		goto L70
	} else {
		goto L641
	}
L641:
	;
	F_errmsg(m, int32(_a_F_NUM_processor_8), int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L70
	} else {
		goto L642
	}
L642:
	;
	F_errfinish(m, int32(_a_F_NUM_processor_9), int32(_a_F_NUM_processor_10), int32(_a_F_NUM_processor_11))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L70
	} else {
		goto L643
	}
L643:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L644:
	;
	v2460 = base.B2i32(base.Ui32(v613) <= base.Ui32(int32(6)))
	goto L646
L645:
	;
	v2460 = int32(0)
	goto L646
L646:
	;
	if v2460 == int32(0) {
		goto L647
	} else {
		goto L648
	}
L647:
	;
	v2868 = v2449
	v2872 = v486
	v2873 = int32(0)
	v2876 = v2451
	v2881 = v495 + int32(1)
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L648:
	;
	goto L649
L649:
	;
	if v495 < v358 {
		goto L651
	} else {
		goto L652
	}
L650:
	;
	v2709 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2710 = int32(1)
	v2714 = v399 + base.B2i32(v358 != int32(0)) + int32(base.Ui32(v2709)>>(uint(v2710)%32))&v2710
	if v2706 == v367 {
		goto L724
	} else {
		goto L725
	}
L651:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2470 = int32(0)
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.B2i32(v2467&int32(8) == v2470)|base.B2i32(v495 < v2472) == v2470 {
		goto L655
	} else {
		goto L656
	}
L652:
	;
	goto L653
L653:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	if v2489 == int32(46) {
		goto L660
	} else {
		goto L661
	}
L654:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2449))) = uint8(v2483)
	v2704 = v2449 + int32(1)
	v2706 = v486
	v2707 = v2484
	goto L650
L655:
	;
	v2483 = int32(48)
	v2484 = int32(1)
	goto L654
L656:
	;
	goto L657
L657:
	;
	v2479 = int32(32)
	v2480 = int32(0)
	if v2467&v2479 != 0 {
		v2704 = v2449
		v2706 = v486
		v2707 = v2480
		goto L650
	} else {
		goto L658
	}
L658:
	;
	v2483 = v2479
	v2484 = v2480
	goto L654
L659:
	;
	v2699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	v2704 = v2697
	v2706 = v486 + base.B2i32(v2699 != int32(0))
	v2707 = v2698
	goto L650
L660:
	;
	if v367 != 0 {
		goto L664
	} else {
		goto L665
	}
L661:
	;
	goto L662
L662:
	;
	v2653 = int32(0)
	if base.B2i32(v367 == v2653)|base.B2i32(v613 == int32(2)) == v2653 {
		goto L711
	} else {
		goto L712
	}
L663:
	;
	v2572 = int32(0)
	if v2488&int32(32) == v2572 {
		v2697 = v2449
		v2698 = v2572
		goto L659
	} else {
		goto L689
	}
L664:
	;
	v2492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v2492 == int32(46) {
		goto L663
	} else {
		goto L667
	}
L665:
	;
	goto L666
L666:
	;
	if (v449^v2449)&int32(3) != 0 {
		goto L671
	} else {
		goto L672
	}
L667:
	;
	goto L666
L668:
	;
	v2569 = F_strlen(m, v2449)
	mBase = m.M
	v2697 = v2569 + v2449
	v2698 = int32(0)
	goto L659
L669:
	;
	goto L668
L670:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2549))) = uint8(v2548)
	if v2548&int32(255) == int32(0) {
		goto L669
	} else {
		goto L685
	}
L671:
	;
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v2547 = v449
	v2548 = v2500
	v2549 = v2449
	goto L670
L672:
	;
	goto L673
L673:
	;
	if v449&int32(3) != 0 {
		goto L674
	} else {
		goto L675
	}
L674:
	;
	v2504 = v449
	v2506 = v2449
	goto L677
L675:
	;
	v2518 = v449
	v2520 = v2449
	goto L676
L676:
	;
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2518)))
	v2525 = int32(-2139062144)
	if (int32(16843008)-v2522|v2522)&v2525 != v2525 {
		v2547 = v2518
		v2548 = v2522
		v2549 = v2520
		goto L670
	} else {
		goto L681
	}
L677:
	;
	v2507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2504))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2506))) = uint8(v2507)
	if v2507 == int32(0) {
		goto L669
	} else {
		goto L679
	}
L678:
	;
	v2518 = v2514
	v2520 = v2512
	goto L676
L679:
	;
	v2511 = int32(1)
	v2512 = v2506 + v2511
	v2514 = v2504 + v2511
	if v2514&int32(3) != 0 {
		v2504 = v2514
		v2506 = v2512
		goto L677
	} else {
		goto L680
	}
L680:
	;
	goto L678
L681:
	;
	v2530 = v2518
	v2531 = v2522
	v2532 = v2520
	goto L682
L682:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2532))) = v2531
	v2534 = int32(4)
	v2535 = v2532 + v2534
	v2537 = v2530 + v2534
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(v2530)+4))
	v2542 = int32(-2139062144)
	if (int32(16843008)-v2539|v2539)&v2542 == v2542 {
		v2530 = v2537
		v2531 = v2539
		v2532 = v2535
		goto L682
	} else {
		goto L684
	}
L683:
	;
	v2547 = v2537
	v2548 = v2539
	v2549 = v2535
	goto L670
L684:
	;
	goto L683
L685:
	;
	v2556 = v2547
	v2558 = v2549
	goto L686
L686:
	;
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2556)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2558)+1)) = uint8(v2559)
	v2561 = int32(1)
	if v2559 != 0 {
		v2556 = v2556 + v2561
		v2558 = v2558 + v2561
		goto L686
	} else {
		goto L688
	}
L687:
	;
	goto L669
L688:
	;
	goto L687
L689:
	;
	if (v449^v2449)&int32(3) != 0 {
		goto L693
	} else {
		goto L694
	}
L690:
	;
	v2651 = F_strlen(m, v2449)
	mBase = m.M
	v2697 = v2651 + v2449
	v2698 = v2572
	goto L659
L691:
	;
	goto L690
L692:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2631))) = uint8(v2630)
	if v2630&int32(255) == int32(0) {
		goto L691
	} else {
		goto L707
	}
L693:
	;
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	v2629 = v449
	v2630 = v2582
	v2631 = v2449
	goto L692
L694:
	;
	goto L695
L695:
	;
	if v449&int32(3) != 0 {
		goto L696
	} else {
		goto L697
	}
L696:
	;
	v2586 = v449
	v2588 = v2449
	goto L699
L697:
	;
	v2600 = v449
	v2602 = v2449
	goto L698
L698:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v2600)))
	v2607 = int32(-2139062144)
	if (int32(16843008)-v2604|v2604)&v2607 != v2607 {
		v2629 = v2600
		v2630 = v2604
		v2631 = v2602
		goto L692
	} else {
		goto L703
	}
L699:
	;
	v2589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2586))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2588))) = uint8(v2589)
	if v2589 == int32(0) {
		goto L691
	} else {
		goto L701
	}
L700:
	;
	v2600 = v2596
	v2602 = v2594
	goto L698
L701:
	;
	v2593 = int32(1)
	v2594 = v2588 + v2593
	v2596 = v2586 + v2593
	if v2596&int32(3) != 0 {
		v2586 = v2596
		v2588 = v2594
		goto L699
	} else {
		goto L702
	}
L702:
	;
	goto L700
L703:
	;
	v2612 = v2600
	v2613 = v2604
	v2614 = v2602
	goto L704
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2614))) = v2613
	v2616 = int32(4)
	v2617 = v2614 + v2616
	v2619 = v2612 + v2616
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2612)+4))
	v2624 = int32(-2139062144)
	if (int32(16843008)-v2621|v2621)&v2624 == v2624 {
		v2612 = v2619
		v2613 = v2621
		v2614 = v2617
		goto L704
	} else {
		goto L706
	}
L705:
	;
	v2629 = v2619
	v2630 = v2621
	v2631 = v2617
	goto L692
L706:
	;
	goto L705
L707:
	;
	v2638 = v2629
	v2640 = v2631
	goto L708
L708:
	;
	v2641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2638)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2640)+1)) = uint8(v2641)
	v2643 = int32(1)
	if v2641 != 0 {
		v2638 = v2638 + v2643
		v2640 = v2640 + v2643
		goto L708
	} else {
		goto L710
	}
L709:
	;
	goto L691
L710:
	;
	goto L709
L711:
	;
	if base.Ui32(v367) < base.Ui32(v486) {
		v2697 = v2449
		v2698 = int32(0)
		goto L659
	} else {
		goto L714
	}
L712:
	;
	goto L713
L713:
	;
	if v2488&int32(8)|base.B2i32(l3 != v486) != 0 {
		goto L715
	} else {
		goto L716
	}
L714:
	;
	goto L713
L715:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2449))) = uint8(v2489)
	v2693 = int32(1)
	v2697 = v2449 + v2693
	v2698 = v2693
	goto L659
L716:
	;
	v2667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v2667 != int32(48) {
		goto L715
	} else {
		goto L717
	}
L717:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2670 == int32(0) {
		goto L715
	} else {
		goto L718
	}
L718:
	;
	if v2488&int32(32) == int32(0) {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	v2677 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v2449))) = uint8(v2677)
	v2697 = v2449 + int32(1)
	v2698 = int32(0)
	goto L659
L720:
	;
	goto L721
L721:
	;
	v2682 = int32(0)
	if v367 == v2682 {
		v2697 = v2449
		v2698 = v2682
		goto L659
	} else {
		goto L722
	}
L722:
	;
	v2685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v2685 != int32(46) {
		v2697 = v2449
		v2698 = v2682
		goto L659
	} else {
		goto L723
	}
L723:
	;
	v2688 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v2449))) = uint8(v2688)
	v2697 = v2449 + int32(1)
	v2698 = v2682
	goto L659
L724:
	;
	v2716 = v495
	goto L726
L725:
	;
	v2716 = v2714
	goto L726
L726:
	;
	v2718 = v495 + int32(1)
	if v367 != 0 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	v2719 = v2716
	goto L729
L728:
	;
	v2719 = v2714
	goto L729
L729:
	;
	if v2718 != v2719 {
		v2868 = v2704
		v2872 = v2706
		v2873 = v2707
		v2876 = v2451
		v2881 = v2718
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L730
	}
L730:
	;
	v2723 = int32(0)
	if v2450|base.B2i32(v2709&int32(128) == v2723) == v2723 {
		goto L731
	} else {
		goto L732
	}
L731:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2704))) = uint8(v464)
	v2820 = v2704
	v2824 = v2706
	v2825 = v2707
	v2828 = v2451
	v2833 = v2718
	v2836 = v498
	v2840 = v502
	v2846 = v508
	goto L147
L732:
	;
	goto L733
L733:
	;
	if v2709&int32(64) == int32(0) {
		v2868 = v2704
		v2872 = v2706
		v2873 = v2707
		v2876 = v2451
		v2881 = v2718
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L734
	}
L734:
	;
	v2733 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v2733 != int32(1) {
		v2868 = v2704
		v2872 = v2706
		v2873 = v2707
		v2876 = v2451
		v2881 = v2718
		v2884 = v498
		v2888 = v502
		v2894 = v508
		goto L116
	} else {
		goto L735
	}
L735:
	;
	if (v459^v2704)&int32(3) != 0 {
		goto L739
	} else {
		goto L740
	}
L736:
	;
	v2810 = F_strlen(m, v2704)
	mBase = m.M
	v2868 = v2810 + v2704
	v2872 = v2706
	v2873 = v2707
	v2876 = v2451
	v2881 = v2718
	v2884 = v498
	v2888 = v502
	v2894 = v508
	goto L116
L737:
	;
	goto L736
L738:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2790))) = uint8(v2789)
	if v2789&int32(255) == int32(0) {
		goto L737
	} else {
		goto L753
	}
L739:
	;
	v2741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	v2788 = v459
	v2789 = v2741
	v2790 = v2704
	goto L738
L740:
	;
	goto L741
L741:
	;
	if v459&int32(3) != 0 {
		goto L742
	} else {
		goto L743
	}
L742:
	;
	v2745 = v459
	v2747 = v2704
	goto L745
L743:
	;
	v2759 = v459
	v2761 = v2704
	goto L744
L744:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v2759)))
	v2766 = int32(-2139062144)
	if (int32(16843008)-v2763|v2763)&v2766 != v2766 {
		v2788 = v2759
		v2789 = v2763
		v2790 = v2761
		goto L738
	} else {
		goto L749
	}
L745:
	;
	v2748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2745))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2747))) = uint8(v2748)
	if v2748 == int32(0) {
		goto L737
	} else {
		goto L747
	}
L746:
	;
	v2759 = v2755
	v2761 = v2753
	goto L744
L747:
	;
	v2752 = int32(1)
	v2753 = v2747 + v2752
	v2755 = v2745 + v2752
	if v2755&int32(3) != 0 {
		v2745 = v2755
		v2747 = v2753
		goto L745
	} else {
		goto L748
	}
L748:
	;
	goto L746
L749:
	;
	v2771 = v2759
	v2772 = v2763
	v2773 = v2761
	goto L750
L750:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2773))) = v2772
	v2775 = int32(4)
	v2776 = v2773 + v2775
	v2778 = v2771 + v2775
	v2780 = *(*int32)(unsafe.Add(mBase, uint32(v2771)+4))
	v2783 = int32(-2139062144)
	if (int32(16843008)-v2780|v2780)&v2783 == v2783 {
		v2771 = v2778
		v2772 = v2780
		v2773 = v2776
		goto L750
	} else {
		goto L752
	}
L751:
	;
	v2788 = v2778
	v2789 = v2780
	v2790 = v2776
	goto L738
L752:
	;
	goto L751
L753:
	;
	v2797 = v2788
	v2799 = v2790
	goto L754
L754:
	;
	v2800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2797)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2799)+1)) = uint8(v2800)
	v2802 = int32(1)
	if v2800 != 0 {
		v2797 = v2797 + v2802
		v2799 = v2799 + v2802
		goto L754
	} else {
		goto L756
	}
L755:
	;
	goto L737
L756:
	;
	goto L755
L757:
	;
	goto L115
L758:
	;
	v2959 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2919))) = uint8(v2959)
	goto L5
L759:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v2989
	goto L5
L760:
	;
	v3012 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3008))) = uint8(v3012)
	goto L759
L761:
	;
	goto L762
L762:
	;
	v3014 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2973))) = uint8(v3014)
	goto L759
L763:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L70
	} else {
		goto L764
	}
L764:
	;
	F_errmsg(m, int32(_a_F_NUM_processor_12), int32(0))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L70
	} else {
		goto L765
	}
L765:
	;
	F_errfinish(m, int32(_a_F_NUM_processor_9), int32(_a_F_NUM_processor_13), int32(_a_F_NUM_processor_11))
	mBase = m.M
	v3081 = m.ExcPending
	if v3081 != 0 {
		goto L70
	} else {
		goto L766
	}
L766:
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
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
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
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v856 int32
	_ = v856
	var v869 int32
	_ = v869
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
		v869 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v24 + int32(1040)
	return v869
L2:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = F_palloc(m, int32(_a_F_NormalizeSubWord_0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(0)
	v48 = F_FindWord(m, l0, l1, int32(_a_F_NormalizeSubWord_1), l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v48 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v50 = F_pstrdup(m, l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v58 = v41
	goto L8
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v59 == int32(0) {
		v300 = v58
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v50
	v58 = v41 + int32(4)
	goto L8
L10:
	;
	if v39 == int32(0) {
		v838 = v300
		goto L66
	} else {
		goto L67
	}
L11:
	;
	v65 = v59
	v67 = v58
	v71 = v4
	goto L12
L12:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v83&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v300 = v287
	goto L10
L14:
	;
	v202 = v67
	v203 = int32(0)
	goto L43
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if base.Ui32(int32(255)) < base.Ui32(v86) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v94 = v65
	goto L17
L17:
	;
	if v26 < v71 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v178 = v65 + int32(4)
	v184 = v71
	goto L14
L19:
	;
	goto L20
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	if v91 == int32(0) {
		v300 = v67
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v94 = v91
	goto L17
L22:
	;
	v96 = v71
	goto L24
L23:
	;
	v96 = v26
	goto L24
L24:
	;
	v100 = v94
	v106 = v71
	goto L25
L25:
	;
	if v106 == v96 {
		v300 = v67
		goto L10
	} else {
		goto L27
	}
L26:
	;
	v300 = v67
	goto L10
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v121 = int32(base.Ui32(v119) >> (uint(int32(1)) % 32))
	if v121 == int32(0) {
		v300 = v67
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v125 = v100 + int32(4)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v106))))
	v135 = v125 + v121*int32(12)
	v137 = v125
	goto L29
L29:
	;
	v153 = int32(12)
	v154 = base.I32_div_s(v135-v137, v153)
	v159 = v137 + int32(base.Ui32(v154)>>(uint(int32(1))%32))*v153
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v162 = v160 & int32(255)
	if v130 == v162 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v165 = v106 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v160) {
		v178 = v159
		v184 = v165
		goto L14
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v171 = base.B2i32(base.Ui32(v162) < base.Ui32(v130))
	if base.Ui32(v162) < base.Ui32(v130) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
	if v168 != 0 {
		v100 = v168
		v106 = v165
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v300 = v67
	goto L10
L36:
	;
	v172 = v159 + int32(12)
	goto L38
L37:
	;
	v172 = v137
	goto L38
L38:
	;
	if base.Ui32(v162) < base.Ui32(v130) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v173 = v135
	goto L41
L40:
	;
	v173 = v159
	goto L41
L41:
	;
	if base.Ui32(v172) < base.Ui32(v173) {
		v135 = v173
		v137 = v172
		goto L29
	} else {
		goto L42
	}
L42:
	;
	goto L30
L43:
	;
	v219 = v203 << (uint(int32(2)) % 32)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219+v220)))
	v224 = v24 + int32(528)
	v226 = F_CheckAffix(m, l1, v26, v222, l2, v224, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	if v294 != 0 {
		v65 = v294
		v67 = v287
		v71 = v184
		goto L12
	} else {
		goto L65
	}
L45:
	;
	v289 = v203 + int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if base.Ui32(v289) < base.Ui32(int32(base.Ui32(v290)>>(uint(int32(8))%32))) {
		v202 = v287
		v203 = v289
		goto L43
	} else {
		goto L64
	}
L46:
	;
	if v226 == int32(0) {
		v287 = v202
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v230+v219)))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v234 = F_FindWord(m, l0, v224, v233, l2)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	if v234 == int32(0) {
		v287 = v202
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v238 = int32(0)
	if int32(4088) < v202-v41 {
		v282 = v238
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v287 = v202 + v282<<(uint(int32(2))%32)
	goto L45
L51:
	;
	if v202 != v41 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v202-int32(4))))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if base.B2i32(v248 == int32(0))|base.B2i32(v248 != v251) != 0 {
		v269 = v248
		v270 = v251
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v276 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L3
	} else {
		goto L63
	}
L55:
	;
	if v269-v270 == int32(0) {
		v282 = v238
		goto L50
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v254 = v224
	v255 = v245
	goto L58
L58:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+1)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	if v259 == int32(0) {
		v269 = v259
		v270 = v258
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v269 = v259
	v270 = v258
	goto L56
L60:
	;
	v262 = int32(1)
	if v259 == v258 {
		v254 = v254 + v262
		v255 = v255 + v262
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
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v276
	v282 = int32(1)
	goto L50
L64:
	;
	goto L44
L65:
	;
	goto L13
L66:
	;
	if v838 != v41 {
		v869 = v41
		goto L1
	} else {
		goto L182
	}
L67:
	;
	v324 = v300
	v326 = v39
	v335 = v4
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v342&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v838 = v810
	goto L66
L70:
	;
	v463 = v324
	v468 = int32(0)
	goto L99
L71:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if base.Ui32(int32(255)) < base.Ui32(v345) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v353 = v326
	goto L73
L73:
	;
	if v26 < v335 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v443 = v326 + int32(4)
	v452 = v335
	goto L70
L75:
	;
	goto L76
L76:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	if v350 == int32(0) {
		v838 = v324
		goto L66
	} else {
		goto L77
	}
L77:
	;
	v353 = v350
	goto L73
L78:
	;
	v355 = v335
	goto L80
L79:
	;
	v355 = v26
	goto L80
L80:
	;
	v363 = v353
	v372 = v335
	goto L81
L81:
	;
	if v355 == v372 {
		v838 = v324
		goto L66
	} else {
		goto L83
	}
L82:
	;
	v838 = v324
	goto L66
L83:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v380 = int32(base.Ui32(v378) >> (uint(int32(1)) % 32))
	if v380 == int32(0) {
		v838 = v324
		goto L66
	} else {
		goto L84
	}
L84:
	;
	v384 = v363 + int32(4)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v26+(v372^int32(-1))))))
	v395 = v384
	v398 = v384 + v380*int32(12)
	goto L85
L85:
	;
	v414 = int32(12)
	v415 = base.I32_div_s(v398-v395, v414)
	v420 = v395 + int32(base.Ui32(v415)>>(uint(int32(1))%32))*v414
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v423 = v421 & int32(255)
	if v391 == v423 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L82
L87:
	;
	v426 = v372 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v421) {
		v443 = v420
		v452 = v426
		goto L70
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v432 = base.B2i32(base.Ui32(v423) < base.Ui32(v391))
	if base.Ui32(v423) < base.Ui32(v391) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v420)+8))
	if v429 != 0 {
		v363 = v429
		v372 = v426
		goto L81
	} else {
		goto L91
	}
L91:
	;
	v838 = v324
	goto L66
L92:
	;
	v433 = v420 + int32(12)
	goto L94
L93:
	;
	v433 = v395
	goto L94
L94:
	;
	if base.Ui32(v423) < base.Ui32(v391) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v434 = v398
	goto L97
L96:
	;
	v434 = v420
	goto L97
L97:
	;
	if base.Ui32(v433) < base.Ui32(v434) {
		v395 = v433
		v398 = v434
		goto L85
	} else {
		goto L98
	}
L98:
	;
	goto L86
L99:
	;
	v480 = v468 << (uint(int32(2)) % 32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v480+v481)))
	v485 = v24 + int32(528)
	v488 = F_CheckAffix(m, l1, v26, v483, l2, v485, v24+int32(12))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L102
	}
L100:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	if v832 != 0 {
		v324 = v810
		v326 = v832
		v335 = v452
		goto L68
	} else {
		goto L181
	}
L101:
	;
	v827 = v468 + int32(1)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	if base.Ui32(v827) < base.Ui32(int32(base.Ui32(v828)>>(uint(int32(8))%32))) {
		v463 = v810
		v468 = v827
		goto L99
	} else {
		goto L180
	}
L102:
	;
	if v488 == int32(0) {
		v810 = v463
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v492+v480)))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v496 = F_FindWord(m, l0, v485, v495, l2)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	if v496 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v498 = int32(0)
	if int32(4088) < v463-v41 {
		v542 = v498
		goto L108
	} else {
		goto L109
	}
L106:
	;
	v547 = v463
	goto L107
L107:
	;
	v550 = F_strlen(m, v24+int32(528))
	mBase = m.M
	v551 = int32(0)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v552 == v551 {
		v810 = v547
		goto L101
	} else {
		goto L122
	}
L108:
	;
	v547 = v463 + v542<<(uint(int32(2))%32)
	goto L107
L109:
	;
	if v463 != v41 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v463-int32(4))))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	if base.B2i32(v508 == int32(0))|base.B2i32(v508 != v511) != 0 {
		v529 = v508
		v530 = v511
		goto L114
	} else {
		goto L115
	}
L111:
	;
	goto L112
L112:
	;
	v536 = F_pstrdup(m, v24+int32(528))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L121
	}
L113:
	;
	if v529-v530 == int32(0) {
		v542 = v498
		goto L108
	} else {
		goto L120
	}
L114:
	;
	goto L113
L115:
	;
	v514 = v485
	v515 = v505
	goto L116
L116:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514)+1)))
	if v519 == int32(0) {
		v529 = v519
		v530 = v518
		goto L114
	} else {
		goto L118
	}
L117:
	;
	v529 = v519
	v530 = v518
	goto L114
L118:
	;
	v522 = int32(1)
	if v519 == v518 {
		v514 = v514 + v522
		v515 = v515 + v522
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
	*(*int32)(unsafe.Add(mBase, uint32(v463)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v536
	v542 = int32(1)
	goto L108
L122:
	;
	v558 = v552
	v560 = v547
	v564 = v551
	goto L123
L123:
	;
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if v576&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v810 = v797
	goto L101
L125:
	;
	v697 = v560
	v698 = int32(0)
	goto L154
L126:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v558)+4))
	if base.Ui32(int32(255)) < base.Ui32(v579) {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	v587 = v558
	goto L128
L128:
	;
	if v550 < v564 {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v673 = v558 + int32(4)
	v679 = v564
	goto L125
L130:
	;
	goto L131
L131:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v558)+12))
	if v584 == int32(0) {
		v810 = v560
		goto L101
	} else {
		goto L132
	}
L132:
	;
	v587 = v584
	goto L128
L133:
	;
	v589 = v564
	goto L135
L134:
	;
	v589 = v550
	goto L135
L135:
	;
	v593 = v587
	v599 = v564
	goto L136
L136:
	;
	if v599 == v589 {
		v810 = v560
		goto L101
	} else {
		goto L138
	}
L137:
	;
	v810 = v560
	goto L101
L138:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
	v614 = int32(base.Ui32(v612) >> (uint(int32(1)) % 32))
	if v614 == int32(0) {
		v810 = v560
		goto L101
	} else {
		goto L139
	}
L139:
	;
	v618 = v593 + int32(4)
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+int32(528)+v599))))
	v630 = v618 + v614*int32(12)
	v632 = v618
	goto L140
L140:
	;
	v648 = int32(12)
	v649 = base.I32_div_s(v630-v632, v648)
	v654 = v632 + int32(base.Ui32(v649)>>(uint(int32(1))%32))*v648
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)))
	v657 = v655 & int32(255)
	if v625 == v657 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	goto L137
L142:
	;
	v660 = v599 + int32(1)
	if base.Ui32(int32(256)) <= base.Ui32(v655) {
		v673 = v654
		v679 = v660
		goto L125
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v666 = base.B2i32(base.Ui32(v657) < base.Ui32(v625))
	if base.Ui32(v657) < base.Ui32(v625) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v654)+8))
	if v663 != 0 {
		v593 = v663
		v599 = v660
		goto L136
	} else {
		goto L146
	}
L146:
	;
	v810 = v560
	goto L101
L147:
	;
	v667 = v654 + int32(12)
	goto L149
L148:
	;
	v667 = v632
	goto L149
L149:
	;
	if base.Ui32(v657) < base.Ui32(v625) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v668 = v630
	goto L152
L151:
	;
	v668 = v654
	goto L152
L152:
	;
	if base.Ui32(v667) < base.Ui32(v668) {
		v630 = v668
		v632 = v667
		goto L140
	} else {
		goto L153
	}
L153:
	;
	goto L141
L154:
	;
	v716 = v698 << (uint(int32(2)) % 32)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v716+v717)))
	v721 = v24 + int32(16)
	v724 = F_CheckAffix(m, v24+int32(528), v550, v719, l2, v721, v24+int32(12))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L157
	}
L155:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v673)+8))
	if v804 != 0 {
		v558 = v804
		v560 = v797
		v564 = v679
		goto L123
	} else {
		goto L179
	}
L156:
	;
	v799 = v698 + int32(1)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	if base.Ui32(v799) < base.Ui32(int32(base.Ui32(v800)>>(uint(int32(8))%32))) {
		v697 = v797
		v698 = v799
		goto L154
	} else {
		goto L178
	}
L157:
	;
	if v724 == int32(0) {
		v797 = v697
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v673)+4))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v728+v716)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+4))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v732+v480)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v734)+4))
	if v731&v735&int32(128) != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v741 = int32(_a_F_NormalizeSubWord_1)
	goto L161
L160:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v730)))
	v741 = v740
	goto L161
L161:
	;
	v742 = F_FindWord(m, l0, v721, v741, l2)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	if v742 == int32(0) {
		v797 = v697
		goto L156
	} else {
		goto L163
	}
L163:
	;
	v746 = int32(0)
	if int32(4088) < v697-v41 {
		v792 = v746
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v797 = v697 + v792<<(uint(int32(2))%32)
	goto L156
L165:
	;
	if v697 != v41 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v752 = v24 + int32(16)
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v697-int32(4))))
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752))))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	if base.B2i32(v758 == int32(0))|base.B2i32(v758 != v761) != 0 {
		v779 = v758
		v780 = v761
		goto L170
	} else {
		goto L171
	}
L167:
	;
	goto L168
L168:
	;
	v786 = F_pstrdup(m, v24+int32(16))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L3
	} else {
		goto L177
	}
L169:
	;
	if v779-v780 == int32(0) {
		v792 = v746
		goto L164
	} else {
		goto L176
	}
L170:
	;
	goto L169
L171:
	;
	v764 = v752
	v765 = v755
	goto L172
L172:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+1)))
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v764)+1)))
	if v769 == int32(0) {
		v779 = v769
		v780 = v768
		goto L170
	} else {
		goto L174
	}
L173:
	;
	v779 = v769
	v780 = v768
	goto L170
L174:
	;
	v772 = int32(1)
	if v769 == v768 {
		v764 = v764 + v772
		v765 = v765 + v772
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
	*(*int32)(unsafe.Add(mBase, uint32(v697)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v697))) = v786
	v792 = int32(1)
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
	F_pfree(m, v41)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L3
	} else {
		goto L183
	}
L183:
	;
	v869 = int32(0)
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
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v3 = F_pull_varnos(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v8 = F_bms_del_members(m, v3, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = int32(0)
			if v8 == v10 {
				v45 = int32(0)
			} else {
				v17 = int32(1)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				if v18 <= v17 {
					v21 = v17
				} else {
					v21 = v18
				}
				v25 = int32(0)
				v27 = v10
				for {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)+v25<<(uint(int32(2))%32))))
					if v33 != 0 {
						v36 = v27 + base.I32_popcnt(v33)
					} else {
						v36 = v27
					}
					v38 = v25 + int32(1)
					if v38 != v21 {
						v25 = v38
						v27 = v36
						continue
					} else {
						break
					}
					break
				}
				v45 = v36
			}
			F_bms_free(m, v8)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				return v45
			}
		}
	}
}
func F_nameconcatoid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	v17 = F_pg_snprintf(m, v8+int32(16), int32(20), int32(_a_F_nameconcatoid_0), v8)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = F_strlen(m, v10)
		mBase = m.M
		if int32(64) <= v17+v21 {
			v27 = F_pg_mbcliplen(m, v10, v21, int32(63)-v17)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = v27
				v31 = F_palloc0(m, int32(64))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
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
					return v31
				}
			}
		} else {
			v29 = v21
			v31 = F_palloc0(m, int32(64))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
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
				return v31
			}
		}
	}
}
func F_namegttext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1542), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v6)
	}
}
func F_namehashfast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	v2 = F_strlen(m, l0)
	mBase = m.M
	v8 = v2 - int32(1636608432)
	if l0&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v2) {
			v117 = l0
			v118 = v2
			v119 = v8
			v120 = v8
			v121 = v8
			for {
				v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				v124 = v123 + v120
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v128 = v127 + v121
				v130 = int32(4)
				v132 = v125 + v119 - v128 ^ base.I32_rotl(v128, v130)
				v136 = v124 - v132 ^ base.I32_rotl(v132, int32(6))
				v137 = v128 + v124
				v138 = v132 + v137
				v139 = v136 + v138
				v143 = v137 - v136 ^ base.I32_rotl(v136, int32(8))
				v147 = v138 - v143 ^ base.I32_rotl(v143, int32(16))
				v151 = v139 - v147 ^ base.I32_rotl(v147, int32(19))
				v152 = v143 + v139
				v153 = v147 + v152
				v154 = v151 + v153
				v158 = v152 - v151 ^ base.I32_rotl(v151, v130)
				v159 = int32(12)
				v160 = v117 + v159
				v162 = v118 - v159
				if base.Ui32(int32(11)) < base.Ui32(v162) {
					v117 = v160
					v118 = v162
					v119 = v153
					v120 = v154
					v121 = v158
					continue
				} else {
					break
				}
				break
			}
			v165 = v160
			v166 = v162
			v167 = v153
			v168 = v154
			v169 = v158
		} else {
			v165 = l0
			v166 = v2
			v167 = v8
			v168 = v8
			v169 = v8
		}
		switch v166 - int32(1) {
		case 0:
			v228 = v167
			v229 = v168
			v230 = v169
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 1:
			v221 = v167
			v222 = v168
			v223 = v169
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 2:
			v214 = v167
			v215 = v168
			v216 = v169
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 3:
			v208 = v168
			v209 = v169
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 4:
			v204 = v168
			v205 = v169
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 5:
			v198 = v168
			v199 = v169
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 6:
			v192 = v168
			v193 = v169
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 7:
			v187 = v169
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 8:
			v182 = v169
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 9:
			v177 = v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		case 10:
			v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+10)))
			v177 = v173<<(uint(int32(24))%32) + v169
			v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+9)))
			v182 = v178<<(uint(int32(16))%32) + v177
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+8)))
			v187 = v183<<(uint(int32(8))%32) + v182
			v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+7)))
			v192 = v188<<(uint(int32(24))%32) + v168
			v193 = v187
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
			v198 = v194<<(uint(int32(16))%32) + v192
			v199 = v193
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+5)))
			v204 = v200<<(uint(int32(8))%32) + v198
			v205 = v199
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+4)))
			v208 = v204 + v206
			v209 = v205
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+3)))
			v214 = v210<<(uint(int32(24))%32) + v167
			v215 = v208
			v216 = v209
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
			v221 = v217<<(uint(int32(16))%32) + v214
			v222 = v215
			v223 = v216
			v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
			v228 = v224<<(uint(int32(8))%32) + v221
			v229 = v222
			v230 = v223
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
			v235 = v228 + v231
			v236 = v229
			v237 = v230
		default:
			v235 = v167
			v236 = v168
			v237 = v169
		}
	} else {
		if base.Ui32(v2) < base.Ui32(int32(12)) {
			v63 = l0
			v64 = v2
			v65 = v8
			v66 = v8
			v67 = v8
		} else {
			v15 = l0
			v16 = v2
			v17 = v8
			v18 = v8
			v19 = v8
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v22 = v21 + v18
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v26 = v25 + v19
				v28 = int32(4)
				v30 = v23 + v17 - v26 ^ base.I32_rotl(v26, v28)
				v34 = v22 - v30 ^ base.I32_rotl(v30, int32(6))
				v35 = v26 + v22
				v36 = v30 + v35
				v37 = v34 + v36
				v41 = v35 - v34 ^ base.I32_rotl(v34, int32(8))
				v45 = v36 - v41 ^ base.I32_rotl(v41, int32(16))
				v49 = v37 - v45 ^ base.I32_rotl(v45, int32(19))
				v50 = v41 + v37
				v51 = v45 + v50
				v52 = v49 + v51
				v56 = v50 - v49 ^ base.I32_rotl(v49, v28)
				v57 = int32(12)
				v58 = v15 + v57
				v60 = v16 - v57
				if base.Ui32(int32(11)) < base.Ui32(v60) {
					v15 = v58
					v16 = v60
					v17 = v51
					v18 = v52
					v19 = v56
					continue
				} else {
					break
				}
				break
			}
			v63 = v58
			v64 = v60
			v65 = v51
			v66 = v52
			v67 = v56
		}
		switch v64 - int32(1) {
		case 0:
			v114 = v65
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 1:
			v109 = v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 2:
			v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)))
			v109 = v105<<(uint(int32(16))%32) + v65
			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
			v114 = v110<<(uint(int32(8))%32) + v109
			v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
			v235 = v114 + v115
			v236 = v66
			v237 = v67
		case 3:
			v102 = v66
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 4:
			v99 = v66
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 5:
			v94 = v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 6:
			v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
			v94 = v90<<(uint(int32(16))%32) + v66
			v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+5)))
			v99 = v95<<(uint(int32(8))%32) + v94
			v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
			v102 = v99 + v100
			v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v235 = v103 + v65
			v236 = v102
			v237 = v67
		case 7:
			v85 = v67
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 8:
			v80 = v67
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 9:
			v75 = v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		case 10:
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+10)))
			v75 = v71<<(uint(int32(24))%32) + v67
			v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+9)))
			v80 = v76<<(uint(int32(16))%32) + v75
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+8)))
			v85 = v81<<(uint(int32(8))%32) + v80
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
			v88 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
			v235 = v86 + v65
			v236 = v88 + v66
			v237 = v85
		default:
			v235 = v65
			v236 = v66
			v237 = v67
		}
	}
	v240 = int32(14)
	v242 = v236 ^ v237 - base.I32_rotl(v236, v240)
	v246 = v242 ^ v235 - base.I32_rotl(v242, int32(11))
	v250 = v246 ^ v236 - base.I32_rotl(v246, int32(25))
	v254 = v250 ^ v242 - base.I32_rotl(v250, int32(16))
	v258 = v254 ^ v246 - base.I32_rotl(v254, int32(4))
	v262 = v258 ^ v250 - base.I32_rotl(v258, v240)
	return v262 ^ v254 - base.I32_rotl(v262, int32(24))
}
func F_nameout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pstrdup(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
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
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v144 != 0 {
		goto L46
	} else {
		goto L47
	}
L2:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = int32(0)
	if v119&int32(32) != 0 {
		goto L43
	} else {
		goto L44
	}
L3:
	;
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+69)) = uint8(v95)
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+68)) = uint8(base.B2i32(l3 == v95))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+36)) = v94 + int32(3916)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v94 + int32(1516)
	v107 = v94 + int32(1352)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = v107
	v110 = v94 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v94)+28)) = v107 + v15<<(uint(int32(2))%32)
	v116 = v110
	v117 = v94
	goto L2
L4:
	;
	v29 = F_palloc_extended(m, int32(72), int32(2))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L11
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
		v94 = l3
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
		v94 = v23
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L1
L11:
	;
	if v29 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	goto L1
L13:
	;
	goto L14
L14:
	;
	v36 = F_palloc_extended(m, v9<<(uint(int32(6))%32), int32(2))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v36
	v42 = int32(2)
	v45 = F_palloc_extended(m, v13*(v15|int32(1))<<(uint(v42)%32), v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v45
	v49 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v45 + v15*v13<<(uint(v49)%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v58 = F_palloc_extended(m, v9*v53<<(uint(int32(3))%32), v49)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v66 = F_palloc_extended(m, v9*v61<<(uint(int32(4))%32), int32(2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v68 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+68)) = uint16(v68)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v66
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if v71 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v72 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	F_pfree(m, v71)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L26
	}
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v75 == int32(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v66 != 0 {
		v116 = v71
		v117 = v29
		goto L2
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	goto L21
L27:
	;
	F_pfree(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v29)+32))
	if v83 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_pfree(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29)+36))
	if v86 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+68)))
	if v89 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	F_pfree(m, v29)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L1
L42:
	;
	goto L41
L43:
	;
	v125 = int32(7)
	goto L45
L44:
	;
	v125 = v15
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+56)) = v116
	*(*int64)(unsafe.Add(mBase, uint32(v117)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v117)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v117)+12)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v117)+60)) = int64(4294967295)
	return v117
L46:
	;
	v146 = v144
	goto L48
L47:
	;
	v146 = int32(12)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v146
	return int32(0)
}
func F_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v453 int32
	_ = v453
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v550 int32
	_ = v550
	var v561 int32
	_ = v561
	var v572 int32
	_ = v572
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v928 int32
	_ = v928
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		v996 = v2
		goto L14
	} else {
		goto L15
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1087 != 0 {
		goto L351
	} else {
		goto L352
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1077 != 0 {
		goto L348
	} else {
		goto L349
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1069 != 0 {
		goto L345
	} else {
		goto L346
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1059 != 0 {
		goto L342
	} else {
		goto L343
	}
L5:
	;
	return int32(1)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(36)
	goto L5
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(94)
	goto L5
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(91)
	goto L6
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(28)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1017)+8)) = v1018 | int32(128)
	v1022 = int32(60)
	if v1013 == v1022 {
		goto L339
	} else {
		goto L340
	}
L12:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v1000 != int32(91) {
		goto L334
	} else {
		goto L335
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	goto L5
L14:
	;
	return v996
L15:
	;
	goto L24
L16:
	;
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v928&int32(2) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L17:
	;
	if base.Ui32(v154) < base.Ui32(v127) {
		goto L16
	} else {
		goto L317
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	goto L5
L19:
	;
	if v127-v154 < int32(21) {
		goto L306
	} else {
		goto L307
	}
L20:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v884 == int32(40) {
		goto L303
	} else {
		goto L304
	}
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	goto L5
L22:
	;
	if base.Ui32(v808) <= base.Ui32(v811) {
		goto L295
	} else {
		goto L296
	}
L23:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v852)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v852)+8)) = v853 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(8589934668)
	goto L5
L24:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = int32(0)
	if base.B2i32(v18&int32(1024) == v21)|base.B2i32(v16 != int32(110)) == v21 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v845)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v845)+8)) = v846 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(12884901964)
	goto L5
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(65)
	goto L5
L27:
	;
	goto L28
L28:
	;
	v32 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.B2i32(v18&int32(32) == v32)|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v34))|base.B2i32(int32(1)<<(uint(v34)%32)&int32(54) == v32) == v32 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = v52
	v56 = v51
	goto L33
L30:
	;
	v125 = v34
	goto L31
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.B2i32(base.Ui32(v126) < base.Ui32(v127))|base.B2i32(base.Ui32(int32(9)) < base.Ui32(v125)) != 0 {
		goto L60
	} else {
		goto L61
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v125 = v124
	goto L31
L33:
	;
	if base.Ui32(v56) <= base.Ui32(v54) {
		v97 = v54
		v99 = v56
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if base.Ui32(v97) < base.Ui32(v99) {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v60 = v54
	v62 = v56
	goto L37
L37:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v66 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	default:
		goto L43
	}
L38:
	;
	v97 = v93
	v99 = v88
	goto L35
L39:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v87 == int32(0) {
		v97 = v89
		v99 = v88
		goto L35
	} else {
		goto L46
	}
L40:
	;
	if base.Ui32(int32(255)) < base.Ui32(v64) {
		v97 = v60
		v99 = v62
		goto L35
	} else {
		goto L45
	}
L41:
	;
	v78 = F_iswspace(m, v64)
	mBase = m.M
	v87 = v78
	goto L39
L42:
	;
	v74 = F_pg_u_isspace(m, v64)
	mBase = m.M
	v87 = v74
	goto L39
L43:
	;
	if base.Ui32(int32(127)) < base.Ui32(v64) {
		v97 = v60
		v99 = v62
		goto L35
	} else {
		goto L44
	}
L44:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_next[1]))))
	v87 = int32(base.Ui32(v71) >> (uint(int32(7)) % 32))
	goto L39
L45:
	;
	v84 = F___isspace(m, v64)
	mBase = m.M
	v87 = base.B2i32(v84 != int32(0))
	goto L39
L46:
	;
	v93 = v89 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
	if base.Ui32(v93) < base.Ui32(v88) {
		v60 = v93
		v62 = v88
		goto L37
	} else {
		goto L47
	}
L47:
	;
	goto L38
L48:
	;
	v113 = v97
	goto L56
L49:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v102 == int32(35) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v97 != v52 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v107 | int32(128)
	goto L55
L54:
	;
	goto L55
L55:
	;
	goto L32
L56:
	;
	v118 = v113 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118
	if base.Ui32(v99) <= base.Ui32(v118) {
		v54 = v118
		v56 = v99
		goto L33
	} else {
		goto L58
	}
L57:
	;
	v54 = v118
	v56 = v99
	goto L33
L58:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v121 != int32(10) {
		v113 = v118
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v154 = v126 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	switch v125 - int32(2) {
	case 0:
		goto L76
	case 1:
		goto L7
	case 2, 3:
		goto L75
	case 4:
		goto L74
	case 5:
		goto L73
	case 6:
		goto L72
	case 7:
		goto L71
	default:
		goto L70
	}
L61:
	;
	v133 = int32(1) << (uint(v125) % 32)
	if v133&int32(960) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v133&int32(14) != 0 {
		goto L13
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L4
L65:
	;
	if v133&int32(48) == int32(0) {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v146 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v148 = v146
	goto L69
L68:
	;
	v148 = int32(9)
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v148
	return int32(0)
L70:
	;
	if v156 != int32(40) {
		goto L217
	} else {
		goto L218
	}
L71:
	;
	if base.B2i32(v156 != int32(58))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L214
	} else {
		goto L215
	}
L72:
	;
	if base.B2i32(v156 != int32(61))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L211
	} else {
		goto L212
	}
L73:
	;
	if base.B2i32(v156 != int32(46))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L208
	} else {
		goto L209
	}
L74:
	;
	switch v156 - int32(91) {
	case 0:
		goto L181
	case 1:
		goto L182
	case 2:
		goto L183
	default:
		goto L180
	}
L75:
	;
	v409 = v156 - int32(48)
	if base.Ui32(int32(10)) <= base.Ui32(v409) {
		goto L165
	} else {
		goto L166
	}
L76:
	;
	switch v156 - int32(36) {
	case 0:
		goto L78
	default:
		goto L7
	case 6:
		goto L82
	case 10:
		goto L80
	case 55:
		goto L81
	case 56:
		goto L77
	case 58:
		goto L79
	}
L77:
	;
	if base.Ui32(v127) <= base.Ui32(v154) {
		goto L141
	} else {
		goto L142
	}
L78:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v214&int32(32) != 0 {
		goto L103
	} else {
		goto L104
	}
L79:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v202 != int32(40) {
		goto L99
	} else {
		goto L100
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(46)
	goto L5
L81:
	;
	if v127-v154 < int32(21) {
		goto L87
	} else {
		goto L88
	}
L82:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v161 - int32(94) {
	case 0, 16:
		goto L84
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L83
	default:
		goto L85
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	goto L5
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(180388626544)
	goto L5
L85:
	;
	if v161 != int32(40) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v127) <= base.Ui32(v154) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v173 != int32(91) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v176 != int32(58) {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	switch v179 - int32(60) {
	case 0, 2:
		goto L91
	default:
		goto L87
	}
L91:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	if v182 != int32(58) {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	if v185 != int32(93) {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	if v188 != int32(93) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v1013 = v179
	goto L11
L95:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	goto L5
L96:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v195 != int32(94) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	goto L10
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(403726925936)
	goto L5
L99:
	;
	if v202 != int32(110) {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v208 | int32(256)
	goto L9
L102:
	;
	goto L9
L103:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v224 = v222
	v226 = v221
	goto L107
L104:
	;
	v296 = v154
	v297 = v127
	goto L105
L105:
	;
	if base.Ui32(v297) <= base.Ui32(v296) {
		goto L134
	} else {
		goto L135
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v296 = v295
	v297 = v294
	goto L105
L107:
	;
	if base.Ui32(v226) <= base.Ui32(v224) {
		v267 = v224
		v269 = v226
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if base.Ui32(v267) < base.Ui32(v269) {
		goto L123
	} else {
		goto L124
	}
L110:
	;
	v230 = v224
	v232 = v226
	goto L111
L111:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v236 - int32(1) {
	case 0:
		goto L116
	case 1:
		goto L115
	case 2:
		goto L114
	default:
		goto L117
	}
L112:
	;
	v267 = v263
	v269 = v258
	goto L109
L113:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v257 == int32(0) {
		v267 = v259
		v269 = v258
		goto L109
	} else {
		goto L120
	}
L114:
	;
	if base.Ui32(int32(255)) < base.Ui32(v234) {
		v267 = v230
		v269 = v232
		goto L109
	} else {
		goto L119
	}
L115:
	;
	v248 = F_iswspace(m, v234)
	mBase = m.M
	v257 = v248
	goto L113
L116:
	;
	v244 = F_pg_u_isspace(m, v234)
	mBase = m.M
	v257 = v244
	goto L113
L117:
	;
	if base.Ui32(int32(127)) < base.Ui32(v234) {
		v267 = v230
		v269 = v232
		goto L109
	} else {
		goto L118
	}
L118:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_c_F_next[1]))))
	v257 = int32(base.Ui32(v241) >> (uint(int32(7)) % 32))
	goto L113
L119:
	;
	v254 = F___isspace(m, v234)
	mBase = m.M
	v257 = base.B2i32(v254 != int32(0))
	goto L113
L120:
	;
	v263 = v259 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v263
	if base.Ui32(v263) < base.Ui32(v258) {
		v230 = v263
		v232 = v258
		goto L111
	} else {
		goto L121
	}
L121:
	;
	goto L112
L122:
	;
	v283 = v267
	goto L130
L123:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v272 == int32(35) {
		goto L122
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if v267 != v222 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+8)) = v277 | int32(128)
	goto L129
L128:
	;
	goto L129
L129:
	;
	goto L106
L130:
	;
	v288 = v283 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v288
	if base.Ui32(v269) <= base.Ui32(v288) {
		v224 = v288
		v226 = v269
		goto L107
	} else {
		goto L132
	}
L131:
	;
	v224 = v288
	v226 = v269
	goto L107
L132:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	if v291 != int32(10) {
		v283 = v288
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	goto L8
L135:
	;
	goto L136
L136:
	;
	if v297-v296 < int32(5) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(154618822768)
	goto L5
L138:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	if v302 != int32(92) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	if v305 != int32(41) {
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v308)+8)) = v309 | int32(256)
	goto L8
L141:
	;
	goto L3
L142:
	;
	goto L143
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	switch v319 - int32(40) {
	case 0:
		goto L149
	case 1:
		goto L148
	default:
		goto L144
	case 9, 10, 11, 12, 13, 14, 15, 16, 17:
		goto L145
	case 20:
		goto L147
	case 22:
		goto L146
	case 83:
		goto L150
	}
L144:
	;
	v359 = int32(0)
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v361 - int32(1) {
	case 0:
		goto L155
	case 1:
		goto L154
	case 2:
		goto L153
	default:
		goto L156
	}
L145:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v349)+8)) = v350 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v319 - int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(98)
	goto L5
L146:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v342)+8)) = v343 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(62)
	goto L5
L147:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v335)+8)) = v336 | int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(60)
	goto L5
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	goto L5
L149:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967336)
	goto L5
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(5)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+8)) = v325 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	goto L5
L151:
	;
	if v393 != 0 {
		goto L159
	} else {
		goto L160
	}
L152:
	;
	v393 = v391
	goto L151
L153:
	;
	if base.Ui32(int32(255)) < base.Ui32(v319) {
		v391 = v359
		goto L152
	} else {
		goto L158
	}
L154:
	;
	v382 = F_iswalnum(m, v319)
	mBase = m.M
	v393 = v382
	goto L151
L155:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_next[2]))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+16)))
	v378 = F_pg_u_isalnum(m, v319, (v373^int32(-1))&int32(1))
	mBase = m.M
	v393 = v378
	goto L151
L156:
	;
	if base.Ui32(int32(127)) < base.Ui32(v319) {
		v391 = v359
		goto L152
	} else {
		goto L157
	}
L157:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_next[1]))))
	v393 = base.B2i32(v366&int32(3) != int32(0))
	goto L151
L158:
	;
	v388 = F_isalnum(m, v319)
	mBase = m.M
	v391 = base.B2i32(v388 != int32(0))
	goto L152
L159:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v394)+8)) = v395 | int32(16)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v399)+8)) = v400 | int32(256)
	goto L161
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	goto L5
L162:
	;
	if base.B2i32(v125 != int32(5))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L177
	} else {
		goto L178
	}
L163:
	;
	if v125 == int32(4) {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(44)
	goto L5
L165:
	;
	if v156 == int32(44) {
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v409
	goto L5
L168:
	;
	if v156 == int32(92) {
		goto L162
	} else {
		goto L169
	}
L169:
	;
	if v156 == int32(125) {
		goto L163
	} else {
		goto L170
	}
L170:
	;
	goto L2
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v427&int32(2) == int32(0))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	goto L2
L174:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	goto L5
L175:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v434 != int32(63) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+8)) = v441 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(125)
	goto L5
L177:
	;
	goto L2
L178:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v453 != int32(125) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967421)
	goto L6
L180:
	;
	if v156 == int32(45) {
		goto L12
	} else {
		goto L207
	}
L181:
	;
	if base.Ui32(v127) <= base.Ui32(v154) {
		goto L200
	} else {
		goto L201
	}
L182:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v475)+8)) = v476 | int32(64)
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v480&int32(2) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v462 == int32(91) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(399431958640)
	goto L5
L185:
	;
	goto L186
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(93)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(2) - v470&int32(1)
	goto L5
L187:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(395136991344)
	goto L5
L188:
	;
	goto L189
L189:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v487)+8)) = v488 | int32(128)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v493) <= base.Ui32(v492) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	goto L3
L191:
	;
	goto L192
L192:
	;
	v495 = F_lexescape(m, l0)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	return int32(0)
L194:
	;
	if v495 == int32(0) {
		v996 = v2
		goto L14
	} else {
		goto L195
	}
L195:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v503 = v501 - int32(99)
	if base.Ui32(v503) <= base.Ui32(int32(16)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v506 = int32(1)
	if v506<<(uint(v503)%32)&int32(_a_F_next_0) != 0 {
		v996 = v506
		goto L14
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	goto L3
L199:
	;
	goto L198
L200:
	;
	goto L4
L201:
	;
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	switch v516 - int32(46) {
	case 0:
		goto L206
	default:
		goto L203
	case 12:
		goto L204
	case 15:
		goto L205
	}
L203:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(390842024048)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	goto L5
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(9)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v535 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(67)
	goto L5
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(8)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v525)+8)) = v526 | int32(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(69)
	goto L5
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(73)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(7)
	goto L5
L207:
	;
	goto L7
L208:
	;
	goto L7
L209:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v550 != int32(93) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(197568495704)
	goto L6
L211:
	;
	goto L7
L212:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v561 != int32(93) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(261993005144)
	goto L6
L214:
	;
	goto L7
L215:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v572 != int32(93) {
		goto L214
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(249108103256)
	goto L6
L217:
	;
	switch v156 - int32(36) {
	case 0:
		goto L8
	default:
		goto L7
	case 5:
		goto L20
	case 6:
		goto L223
	case 7:
		goto L222
	case 10:
		goto L18
	case 27:
		goto L221
	case 55:
		goto L19
	case 56:
		goto L17
	case 58:
		goto L9
	case 87:
		goto L220
	case 88:
		goto L224
	}
L218:
	;
	goto L219
L219:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v789&int32(2) == int32(0))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L21
	} else {
		goto L278
	}
L220:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v651&int32(32) != 0 {
		goto L235
	} else {
		goto L236
	}
L221:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v629&int32(2) == int32(0))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L231
	} else {
		goto L232
	}
L222:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v607&int32(2) == int32(0))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L228
	} else {
		goto L229
	}
L223:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(v585&int32(2) == int32(0))|base.B2i32(base.Ui32(v127) <= base.Ui32(v154)) != 0 {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(124)
	goto L5
L225:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967338)
	goto L5
L226:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v592 != int32(63) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v598)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v598)+8)) = v599 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(42)
	goto L5
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967339)
	goto L5
L229:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v614 != int32(63) {
		goto L228
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v620)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v620)+8)) = v621 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(43)
	goto L5
L231:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967359)
	goto L5
L232:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v636 != int32(63) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v126 + int32(8)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v642)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v642)+8)) = v643 | int32(128)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(63)
	goto L5
L234:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+8))
	v782 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v780)+8)) = v781 | v782
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(123)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v782
	goto L5
L235:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v661 = v659
	v663 = v658
	goto L239
L236:
	;
	v733 = v154
	v734 = v127
	goto L237
L237:
	;
	if base.Ui32(v733) < base.Ui32(v734) {
		goto L266
	} else {
		goto L267
	}
L238:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v733 = v731
	v734 = v732
	goto L237
L239:
	;
	if base.Ui32(v663) <= base.Ui32(v661) {
		v704 = v661
		v706 = v663
		goto L241
	} else {
		goto L242
	}
L241:
	;
	if base.Ui32(v704) < base.Ui32(v706) {
		goto L255
	} else {
		goto L256
	}
L242:
	;
	v667 = v661
	v669 = v663
	goto L243
L243:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v673 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v673 - int32(1) {
	case 0:
		goto L248
	case 1:
		goto L247
	case 2:
		goto L246
	default:
		goto L249
	}
L244:
	;
	v704 = v700
	v706 = v695
	goto L241
L245:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v694 == int32(0) {
		v704 = v696
		v706 = v695
		goto L241
	} else {
		goto L252
	}
L246:
	;
	if base.Ui32(int32(255)) < base.Ui32(v671) {
		v704 = v667
		v706 = v669
		goto L241
	} else {
		goto L251
	}
L247:
	;
	v685 = F_iswspace(m, v671)
	mBase = m.M
	v694 = v685
	goto L245
L248:
	;
	v681 = F_pg_u_isspace(m, v671)
	mBase = m.M
	v694 = v681
	goto L245
L249:
	;
	if base.Ui32(int32(127)) < base.Ui32(v671) {
		v704 = v667
		v706 = v669
		goto L241
	} else {
		goto L250
	}
L250:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+uint32(_c_F_next[1]))))
	v694 = int32(base.Ui32(v678) >> (uint(int32(7)) % 32))
	goto L245
L251:
	;
	v691 = F___isspace(m, v671)
	mBase = m.M
	v694 = base.B2i32(v691 != int32(0))
	goto L245
L252:
	;
	v700 = v696 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v700
	if base.Ui32(v700) < base.Ui32(v695) {
		v667 = v700
		v669 = v695
		goto L243
	} else {
		goto L253
	}
L253:
	;
	goto L244
L254:
	;
	v720 = v704
	goto L262
L255:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v704)))
	if v709 == int32(35) {
		goto L254
	} else {
		goto L258
	}
L256:
	;
	goto L257
L257:
	;
	if v704 != v659 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L257
L259:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v713)+8)) = v714 | int32(128)
	goto L261
L260:
	;
	goto L261
L261:
	;
	goto L238
L262:
	;
	v725 = v720 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v725
	if base.Ui32(v706) <= base.Ui32(v725) {
		v661 = v725
		v663 = v706
		goto L239
	} else {
		goto L264
	}
L263:
	;
	v661 = v725
	v663 = v706
	goto L239
L264:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	if v728 != int32(10) {
		v720 = v725
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v733)))
	v738 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v738 - int32(1) {
	case 0:
		goto L272
	case 1:
		goto L271
	case 2:
		goto L270
	default:
		goto L273
	}
L267:
	;
	goto L268
L268:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v768)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v768)+8)) = v769 | int32(8)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v773)+8)) = v774 | int32(256)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(528280977520)
	goto L5
L269:
	;
	if v767 != 0 {
		goto L234
	} else {
		goto L277
	}
L270:
	;
	if base.Ui32(v736) <= base.Ui32(int32(255)) {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	v756 = F_isdigit(m, v736)
	mBase = m.M
	v767 = v756
	goto L269
L272:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_next[2]))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746)+16)))
	v752 = F_pg_u_isdigit(m, v736, (v747^int32(-1))&int32(1))
	mBase = m.M
	v767 = v752
	goto L269
L273:
	;
	v767 = base.B2i32(base.Ui32(v736-int32(48)) < base.Ui32(int32(10)))
	goto L269
L274:
	;
	v762 = F_isdigit(m, v736)
	mBase = m.M
	v766 = base.B2i32(v762 != int32(0))
	goto L276
L275:
	;
	v766 = int32(0)
	goto L276
L276:
	;
	v767 = v766
	goto L269
L277:
	;
	goto L268
L278:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v796 != int32(63) {
		goto L21
	} else {
		goto L279
	}
L279:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v799)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v799)+8)) = v800 | int32(128)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v806 = v804 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v808) <= base.Ui32(v806) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	goto L1
L281:
	;
	goto L282
L282:
	;
	v811 = v804 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v811
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v813 != int32(35) {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	goto L25
L284:
	;
	switch v813 - int32(33) {
	case 0:
		goto L23
	default:
		goto L1
	case 25:
		goto L287
	case 27:
		goto L22
	case 28:
		goto L283
	}
L285:
	;
	goto L286
L286:
	;
	if base.Ui32(v808) <= base.Ui32(v811) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(40)
	goto L5
L288:
	;
	v842 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v842 == int32(0) {
		goto L24
	} else {
		goto L294
	}
L289:
	;
	v822 = v811
	goto L290
L290:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v822)))
	v830 = v822 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v830
	if v828 == int32(41) {
		goto L288
	} else {
		goto L292
	}
L291:
	;
	goto L288
L292:
	;
	if base.Ui32(v830) < base.Ui32(v808) {
		v822 = v830
		goto L290
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v996 = v2
	goto L14
L295:
	;
	goto L1
L296:
	;
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v804 + int32(12)
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v804)+8))
	if v863 != int32(33) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	goto L1
L299:
	;
	if v863 != int32(61) {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v875)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v875)+8)) = v876 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(76)
	goto L5
L302:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v868)+8)) = v869 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967372)
	goto L5
L303:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v887)+8)) = v888 | int32(256)
	goto L305
L304:
	;
	goto L305
L305:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(176093659177)
	goto L5
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(6)
	if base.Ui32(v127) <= base.Ui32(v154) {
		goto L314
	} else {
		goto L315
	}
L307:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v898 != int32(91) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	if v901 != int32(58) {
		goto L306
	} else {
		goto L309
	}
L309:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	switch v904 - int32(60) {
	case 0, 2:
		goto L310
	default:
		goto L306
	}
L310:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	if v907 != int32(58) {
		goto L306
	} else {
		goto L311
	}
L311:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	if v910 != int32(93) {
		goto L306
	} else {
		goto L312
	}
L312:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v126)+24))
	if v913 != int32(93) {
		goto L306
	} else {
		goto L313
	}
L313:
	;
	v1013 = v904
	goto L11
L314:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(4294967387)
	goto L5
L315:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v920 != int32(94) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	goto L10
L317:
	;
	goto L3
L318:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v934 = int32(0)
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_next[0]))
	switch v936 - int32(1) {
	case 0:
		goto L325
	case 1:
		goto L324
	case 2:
		goto L323
	default:
		goto L326
	}
L319:
	;
	goto L320
L320:
	;
	v988 = F_lexescape(m, l0)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L193
	} else {
		goto L332
	}
L321:
	;
	if v968 != 0 {
		goto L329
	} else {
		goto L330
	}
L322:
	;
	v968 = v966
	goto L321
L323:
	;
	if base.Ui32(int32(255)) < base.Ui32(v933) {
		v966 = v934
		goto L322
	} else {
		goto L328
	}
L324:
	;
	v957 = F_iswalnum(m, v933)
	mBase = m.M
	v968 = v957
	goto L321
L325:
	;
	v947 = *(*int32)(unsafe.Add(mBase, _c_F_next[2]))
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+16)))
	v953 = F_pg_u_isalnum(m, v933, (v948^int32(-1))&int32(1))
	mBase = m.M
	v968 = v953
	goto L321
L326:
	;
	if base.Ui32(int32(127)) < base.Ui32(v933) {
		v966 = v934
		goto L322
	} else {
		goto L327
	}
L327:
	;
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v933)+uint32(_c_F_next[1]))))
	v968 = base.B2i32(v941&int32(3) != int32(0))
	goto L321
L328:
	;
	v963 = F_isalnum(m, v933)
	mBase = m.M
	v966 = base.B2i32(v963 != int32(0))
	goto L322
L329:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v969)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v969)+8)) = v970 | int32(16)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v974)+8)) = v975 | int32(256)
	goto L331
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(112)
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v982 + int32(4)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v982)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v986
	goto L5
L332:
	;
	v996 = v988
	goto L14
L333:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528402)
	return int32(1)
L334:
	;
	if base.Ui32(v127) <= base.Ui32(v154) {
		goto L333
	} else {
		goto L337
	}
L335:
	;
	goto L336
L336:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(193273528432)
	goto L5
L337:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v1004 != int32(93) {
		goto L333
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	v1026 = v1022
	goto L341
L340:
	;
	v1026 = int32(62)
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v1026
	goto L5
L342:
	;
	v1061 = v1059
	goto L344
L343:
	;
	v1061 = int32(7)
	goto L344
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1061
	return int32(0)
L345:
	;
	v1071 = v1069
	goto L347
L346:
	;
	v1071 = int32(5)
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1071
	return int32(0)
L348:
	;
	v1079 = v1077
	goto L350
L349:
	;
	v1079 = int32(10)
	goto L350
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1079
	return int32(0)
L351:
	;
	v1089 = v1087
	goto L353
L352:
	;
	v1089 = int32(13)
	goto L353
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1089
	return int32(0)
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
	F_errfinish(m, int32(_a_F_normalize_exec_path_4), int32(253), int32(_a_F_normalize_exec_path_5))
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
						F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1741), int32(_a_F_nulltestsel_2))
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
							F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1769), int32(_a_F_nulltestsel_2))
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
								F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1769), int32(_a_F_nulltestsel_2))
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
									F_errfinish(m, int32(_a_F_nulltestsel_1), int32(1769), int32(_a_F_nulltestsel_2))
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
func F_numericvar_to_int128(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v197 int64
	_ = v197
	var v199 int64
	_ = v199
	var v218 int64
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v249 int64
	_ = v249
	var v252 int64
	_ = v252
	var v255 int64
	_ = v255
	var v258 int64
	_ = v258
	var v259 int64
	_ = v259
	var v263 int64
	_ = v263
	var v270 int64
	_ = v270
	var v281 int64
	_ = v281
	var v282 int64
	_ = v282
	var v287 int64
	_ = v287
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v308 int64
	_ = v308
	var v309 int64
	_ = v309
	var v312 int64
	_ = v312
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v331 int64
	_ = v331
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v348 int64
	_ = v348
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int64
	_ = v390
	var v392 int64
	_ = v392
	var v398 int64
	_ = v398
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_palloc(m, v20<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 <= v27 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = v41 << (uint(int32(2)) % 32)
	if v43+int32(4) < int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v33 = v29 << (uint(int32(1)) % 32)
	if v33 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v25+int32(2), v38, v33)
	goto L3
L6:
	;
	m.G0 = v18 + int32(32)
	return
L7:
	;
	F_pfree(m, v25)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L54
	}
L8:
	;
	v226 = int32(1)
	v238 = v199
	v239 = v197
	goto L42
L9:
	;
	v218 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v218
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v218
	F_pfree(m, v25)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L41
	}
L10:
	;
	v49 = v25 + int32(2)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v54 = base.I32_div_s(v43+int32(7), int32(4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v55 <= v54 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if int32(0) < v114 {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v114 = v55
	v115 = v49
	v119 = v41
	goto L11
L13:
	;
	goto L14
L14:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49+v54<<(uint(int32(1))%32)))))
	if int32(_a_F_numericvar_to_int128_0) <= v60 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v63 = v54
	goto L18
L16:
	;
	v91 = v54
	goto L17
L17:
	;
	if int32(0) <= v91 {
		v114 = v54
		v115 = v49
		v119 = v41
		goto L11
	} else {
		goto L24
	}
L18:
	;
	v78 = int32(1)
	v80 = v25 + v63<<(uint(v78)%32)
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
	v85 = base.B2i32(int32(_a_F_numericvar_to_int128_1) < v83)
	if int32(_a_F_numericvar_to_int128_1) < v83 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v91 = v90
	goto L17
L20:
	;
	v86 = int32(-9999)
	goto L22
L21:
	;
	v86 = v78
	goto L22
L22:
	;
	v87 = v86 + v83
	*(*uint16)(unsafe.Add(mBase, uint32(v80))) = uint16(v87)
	v90 = v63 - int32(1)
	if int32(_a_F_numericvar_to_int128_1) < v83 {
		v63 = v90
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v108 = int32(1)
	v114 = v54 + v108
	v115 = v25
	v119 = v41 + v108
	goto L11
L25:
	;
	v197 = base.I64_extend16_s(base.I64_extend_i32_u(v181))
	v199 = v197 >> (uint(int64(63)) % 64)
	if int32(0) < v188 {
		goto L8
	} else {
		goto L40
	}
L26:
	;
	v131 = v114
	v132 = v115
	v136 = v119
	goto L29
L27:
	;
	goto L28
L28:
	;
	if v114 == int32(0) {
		goto L9
	} else {
		goto L39
	}
L29:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132))))
	if v144 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v147 = v131
	goto L34
L32:
	;
	goto L33
L33:
	;
	v170 = int32(1)
	if v170 < v131 {
		v131 = v131 - v170
		v132 = v132 + int32(2)
		v136 = v136 - v170
		goto L29
	} else {
		goto L38
	}
L34:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132+v147<<(uint(int32(1))%32)-int32(2)))))
	if v165 != 0 {
		v181 = v144
		v183 = v147
		v184 = v132
		v188 = v136
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v166 = int32(1)
	if v166 < v147 {
		v147 = v147 - v166
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L9
L38:
	;
	goto L9
L39:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v181 = v180
	v183 = v114
	v184 = v115
	v188 = v119
	goto L25
L40:
	;
	v379 = v197
	v380 = v199
	goto L7
L41:
	;
	goto L6
L42:
	;
	v242 = v18 + int32(16)
	v243 = int64(10000)
	v244 = int64(0)
	v249 = int64(32)
	v252 = int64(base.Ui64(v239) >> (uint(v249) % 64))
	v255 = int64(4294967295)
	v258 = v239 & v255
	v259 = v243 * v258
	v263 = int64(base.Ui64(v259)>>(uint(v249)%64)) + v243*v252
	v270 = v258*v244 + v263&v255
	*(*int64)(unsafe.Add(mBase, uint32(v242)+8)) = v239*v244 + v238*v243 + v244*v252 + int64(base.Ui64(v263)>>(uint(v249)%64)) + int64(base.Ui64(v270)>>(uint(v249)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v242))) = v259&v255 | v270<<(uint(v249)%64)
	goto L44
L43:
	;
	v379 = v295
	v380 = v296
	goto L7
L44:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	if v226 < v183 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v287 = int64(*(*int16)(unsafe.Add(mBase, uint32(v184+v226<<(uint(int32(1))%32)))))
	v291 = v282 + v287
	v295 = v291
	v296 = v281 + v287>>(uint(int64(63))%64) + base.I64_extend_i32_u(base.B2i32(base.Ui64(v291) < base.Ui64(v282)))
	goto L47
L46:
	;
	v295 = v282
	v296 = v281
	goto L47
L47:
	;
	v299 = int64(0)
	v303 = m.G0
	v304 = int32(16)
	v305 = v303 - v304
	m.G0 = v305
	v307 = int64(63)
	v308 = v296 >> (uint(v307) % 64)
	v309 = v295 ^ v308
	v312 = v309 + int64(base.Ui64(v296)>>(uint(v307)%64))
	F___udivmodti4(m, v305, v312, base.I64_extend_i32_u(base.B2i32(base.Ui64(v312) < base.Ui64(v309)))+(v308^v296), int64(10000), base.I64_extend_i32_u(int32(0))+v299)
	mBase = m.M
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v305)+8))
	v329 = v308 ^ v299
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v305)))
	v331 = v329 ^ v330
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v331 - v329
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v329 ^ v328 - v329 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v331) < base.Ui64(v329)))
	m.G0 = v305 + v304
	goto L48
L48:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v348 = int64(0)
	v356 = int32(0)
	if base.B2i32(v343^v239|(v345^v238) == v348)|base.B2i32(base.B2i32(v50 != int32(_a_F_numericvar_to_int128_2))|base.B2i32(v295|(v296^int64(-9223372036854775807-1)) != v348) == v356)&base.B2i32(v348 <= v238) == v356 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_pfree(m, v25)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v367 = v226 + int32(1)
	if v367 <= v188 {
		v226 = v367
		v238 = v296
		v239 = v295
		goto L42
	} else {
		goto L53
	}
L52:
	;
	goto L6
L53:
	;
	goto L43
L54:
	;
	v389 = base.B2i32(v50 == int32(_a_F_numericvar_to_int128_2))
	if v50 == int32(_a_F_numericvar_to_int128_2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v390 = int64(0) - v379
	goto L57
L56:
	;
	v390 = v379
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v390
	v392 = int64(0)
	if v50 == int32(_a_F_numericvar_to_int128_2) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v398 = v392 - (v380 + base.I64_extend_i32_u(base.B2i32(v379 != v392)))
	goto L60
L59:
	;
	v398 = v380
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+8)) = v398
	goto L6
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
