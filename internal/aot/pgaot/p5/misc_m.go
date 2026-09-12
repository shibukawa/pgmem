package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MJEvalInnerValues(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(2)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(2)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v13&v12 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	F_MemoryContextReset(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v96 = v12
	goto L6
L6:
	;
	return v96
L7:
	;
	return int32(0)
L8:
	;
	v24 = int32(4489440)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l1
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v31 <= v30 {
		v87 = v30
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v25
	v96 = v87
	goto L6
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int32)(m, v35, v18, v34+int32(17))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v39
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+17)))
	if v43 != int32(1) {
		v52 = int32(0)
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v53 < int32(2) {
		v87 = v52
		goto L9
	} else {
		goto L17
	}
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+29)))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v52 = int32(1)
	goto L12
L15:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)))
	if v47 == int32(1) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v52 = int32(2)
	goto L12
L17:
	;
	v59 = v52
	v61 = int32(1)
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v67 = v64 + v61*int32(56)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int32)(m, v68, v18, v67+int32(17))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L20
	}
L19:
	;
	v87 = v80
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+12)) = v72
	v75 = int32(1)
	if base.Ui32(v59) <= base.Ui32(v75) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v78 = v75
	goto L23
L22:
	;
	v78 = v59
	goto L23
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+17)))
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v78
	goto L26
L25:
	;
	v80 = v59
	goto L26
L26:
	;
	v82 = v61 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v82 < v83 {
		v59 = v80
		v61 = v82
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
}
func F_MemoizeHash_equal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+136))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+132))
	v19 = F_ExecStoreMinimalTuple(m, v16, v17, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+197)))
	if v24 == v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return v117
L4:
	;
	v27 = int32(4489440)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+6)))
	if v35 < v34 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v17
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+140))
	if v98 == int32(0) {
		v117 = v23
		goto L3
	} else {
		goto L27
	}
L7:
	;
	F_slot_getsomeattrs_int(m, v17, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
	if v41 < v40 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	F_slot_getsomeattrs_int(m, v14, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v29 <= int32(0) {
		v88 = v23
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v28
	v117 = v88
	goto L3
L16:
	;
	v48 = int32(0)
	goto L17
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v48))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v48))))
	if v58 != v61 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = v82
	goto L15
L19:
	;
	v88 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	if v58 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v82 = int32(1)
	v84 = v48 + v82
	if v84 != v29 {
		v48 = v84
		goto L17
	} else {
		goto L26
	}
L23:
	;
	v65 = v48 << (uint(int32(2)) % 32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65+v66)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v65)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v75 = v72 + v48<<(uint(int32(4))%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+26)))
	v77 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75)+24)))
	v78 = F_datum_image_eq(m, v68, v71, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v78 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v88 = int32(0)
	goto L15
L26:
	;
	goto L18
L27:
	;
	v101 = int32(4489440)
	v102 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	v109 = m.T0[v108].(func(*base.Module, int32, int32, int32) int32)(m, v98, v15, v11+int32(15))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v102
	v117 = base.B2i32(v109 != int32(0))
	goto L3
}
func F_MemoizeHash_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
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
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v612 int32
	_ = v612
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
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
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
	v2 = int32(0)
	v10 = int32(4489440)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+197)))
	if v19 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v11
	v1218 = int32(16)
	v1222 = (int32(base.Ui32(v1208)>>(uint(v1218)%32)) ^ v1208) * int32(-2048144789)
	v1227 = (int32(base.Ui32(v1222)>>(uint(int32(13))%32)) ^ v1222) * int32(-1028477387)
	return int32(base.Ui32(v1227)>>(uint(v1218)%32)) ^ v1227
L2:
	;
	if v13 <= int32(0) {
		v1208 = v2
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if v13 <= int32(0) {
		v1208 = v2
		goto L1
	} else {
		goto L193
	}
L5:
	;
	v25 = int32(0)
	v26 = v2
	goto L6
L6:
	;
	v35 = base.I32_rotl(v26, int32(1))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v25))))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1208 = v1160
	goto L1
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41+v25<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v49 = v46 + v25<<(uint(int32(4))%32)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+26)))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+24)))
	v52 = m.G0
	v54 = v52 - int32(16)
	m.G0 = v54
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = v45
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v1160 = v35
	goto L10
L10:
	;
	v1166 = v25 + int32(1)
	if v1166 != v13 {
		v25 = v1166
		v26 = v1160
		goto L6
	} else {
		goto L192
	}
L11:
	;
	m.G0 = v54 + int32(16)
	v1160 = v35 ^ v1154
	goto L10
L12:
	;
	v58 = v54 + int32(12)
	v65 = int32(-1636608428)
	if v58&int32(3) != 0 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L14
L14:
	;
	if int32(0) < v51 {
		goto L55
	} else {
		goto L56
	}
L15:
	;
	v1154 = v319 ^ v311 - base.I32_rotl(v319, int32(24))
	goto L11
L16:
	;
	v297 = int32(14)
	v299 = v293 ^ v294 - base.I32_rotl(v293, v297)
	v303 = v299 ^ v292 - base.I32_rotl(v299, int32(11))
	v307 = v303 ^ v293 - base.I32_rotl(v303, int32(25))
	v311 = v307 ^ v299 - base.I32_rotl(v307, int32(16))
	v315 = v311 ^ v303 - base.I32_rotl(v311, int32(4))
	v319 = v315 ^ v307 - base.I32_rotl(v315, v297)
	goto L15
L17:
	;
	switch int32(3) {
	case 0:
		v285 = v65
		v286 = v65
		v287 = v65
		goto L44
	case 1:
		v278 = v65
		v279 = v65
		v280 = v65
		goto L45
	case 2:
		v271 = v65
		v272 = v65
		v273 = v65
		goto L46
	case 3:
		v265 = v65
		v266 = v65
		goto L47
	case 4:
		v261 = v65
		v262 = v65
		goto L48
	case 5:
		v255 = v65
		v256 = v65
		goto L49
	case 6:
		v249 = v65
		v250 = v65
		goto L50
	case 7:
		v244 = v65
		goto L51
	case 8:
		v239 = v65
		goto L52
	case 9:
		v234 = v65
		goto L53
	case 10:
		goto L54
	default:
		v292 = v65
		v293 = v65
		v294 = v65
		goto L16
	}
L19:
	;
	goto L22
L20:
	;
	goto L21
L21:
	;
	goto L24
L22:
	;
	goto L17
L23:
	;
	switch int32(3) {
	case 0:
		v171 = v65
		goto L30
	case 1:
		v166 = v65
		goto L31
	case 2:
		goto L32
	case 3:
		v159 = v65
		goto L33
	case 4:
		v156 = v65
		goto L34
	case 5:
		v151 = v65
		goto L35
	case 6:
		goto L36
	case 7:
		v142 = v65
		goto L37
	case 8:
		v137 = v65
		goto L38
	case 9:
		v132 = v65
		goto L39
	case 10:
		goto L40
	default:
		v292 = v65
		v293 = v65
		v294 = v65
		goto L16
	}
L24:
	;
	goto L23
L30:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v292 = v171 + v172
	v293 = v65
	v294 = v65
	goto L16
L31:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v171 = v167<<(uint(int32(8))%32) + v166
	goto L30
L32:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	v166 = v162<<(uint(int32(16))%32) + v65
	goto L31
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v292 = v160 + v65
	v293 = v159
	v294 = v65
	goto L16
L34:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	v159 = v156 + v157
	goto L33
L35:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+5)))
	v156 = v152<<(uint(int32(8))%32) + v151
	goto L34
L36:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+6)))
	v151 = v147<<(uint(int32(16))%32) + v65
	goto L35
L37:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v292 = v143 + v65
	v293 = v145 + v65
	v294 = v142
	goto L16
L38:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+8)))
	v142 = v138<<(uint(int32(8))%32) + v137
	goto L37
L39:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+9)))
	v137 = v133<<(uint(int32(16))%32) + v132
	goto L38
L40:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+10)))
	v132 = v128<<(uint(int32(24))%32) + v65
	goto L39
L44:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v292 = v285 + v288
	v293 = v286
	v294 = v287
	goto L16
L45:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v285 = v281<<(uint(int32(8))%32) + v278
	v286 = v279
	v287 = v280
	goto L44
L46:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+2)))
	v278 = v274<<(uint(int32(16))%32) + v271
	v279 = v272
	v280 = v273
	goto L45
L47:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+3)))
	v271 = v267<<(uint(int32(24))%32) + v65
	v272 = v265
	v273 = v266
	goto L46
L48:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	v265 = v261 + v263
	v266 = v262
	goto L47
L49:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+5)))
	v261 = v257<<(uint(int32(8))%32) + v255
	v262 = v256
	goto L48
L50:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+6)))
	v255 = v251<<(uint(int32(16))%32) + v249
	v256 = v250
	goto L49
L51:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+7)))
	v249 = v245<<(uint(int32(24))%32) + v65
	v250 = v244
	goto L50
L52:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+8)))
	v244 = v240<<(uint(int32(8))%32) + v239
	goto L51
L53:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+9)))
	v239 = v235<<(uint(int32(16))%32) + v234
	goto L52
L54:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+10)))
	v234 = v230<<(uint(int32(24))%32) + v65
	goto L53
L55:
	;
	v331 = v51 - int32(1636608432)
	if v45&int32(3) != 0 {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	goto L57
L57:
	;
	switch v51 + int32(2) {
	case 0:
		goto L98
	case 1:
		goto L100
	default:
		goto L99
	}
L58:
	;
	v1154 = v585 ^ v577 - base.I32_rotl(v585, int32(24))
	goto L11
L59:
	;
	v563 = int32(14)
	v565 = v559 ^ v560 - base.I32_rotl(v559, v563)
	v569 = v565 ^ v558 - base.I32_rotl(v565, int32(11))
	v573 = v569 ^ v559 - base.I32_rotl(v569, int32(25))
	v577 = v573 ^ v565 - base.I32_rotl(v573, int32(16))
	v581 = v577 ^ v569 - base.I32_rotl(v577, int32(4))
	v585 = v581 ^ v573 - base.I32_rotl(v581, v563)
	goto L58
L60:
	;
	switch v489 - int32(1) {
	case 0:
		v551 = v490
		v552 = v491
		v553 = v492
		goto L87
	case 1:
		v544 = v490
		v545 = v491
		v546 = v492
		goto L88
	case 2:
		v537 = v490
		v538 = v491
		v539 = v492
		goto L89
	case 3:
		v531 = v491
		v532 = v492
		goto L90
	case 4:
		v527 = v491
		v528 = v492
		goto L91
	case 5:
		v521 = v491
		v522 = v492
		goto L92
	case 6:
		v515 = v491
		v516 = v492
		goto L93
	case 7:
		v510 = v492
		goto L94
	case 8:
		v505 = v492
		goto L95
	case 9:
		v500 = v492
		goto L96
	case 10:
		goto L97
	default:
		v558 = v490
		v559 = v491
		v560 = v492
		goto L59
	}
L61:
	;
	v440 = v45
	v441 = v51
	v442 = v331
	v443 = v331
	v444 = v331
	goto L84
L62:
	;
	if base.Ui32(int32(11)) < base.Ui32(v51) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(v51) < base.Ui32(int32(12)) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v488 = v45
	v489 = v51
	v490 = v331
	v491 = v331
	v492 = v331
	goto L60
L66:
	;
	switch v387 - int32(1) {
	case 0:
		v437 = v388
		goto L73
	case 1:
		v432 = v388
		goto L74
	case 2:
		goto L75
	case 3:
		v425 = v389
		goto L76
	case 4:
		v422 = v389
		goto L77
	case 5:
		v417 = v389
		goto L78
	case 6:
		goto L79
	case 7:
		v408 = v390
		goto L80
	case 8:
		v403 = v390
		goto L81
	case 9:
		v398 = v390
		goto L82
	case 10:
		goto L83
	default:
		v558 = v388
		v559 = v389
		v560 = v390
		goto L59
	}
L67:
	;
	v386 = v45
	v387 = v51
	v388 = v331
	v389 = v331
	v390 = v331
	goto L66
L68:
	;
	goto L69
L69:
	;
	v338 = v45
	v339 = v51
	v340 = v331
	v341 = v331
	v342 = v331
	goto L70
L70:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v345 = v344 + v341
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	v349 = v348 + v342
	v351 = int32(4)
	v353 = v346 + v340 - v349 ^ base.I32_rotl(v349, v351)
	v357 = v345 - v353 ^ base.I32_rotl(v353, int32(6))
	v358 = v349 + v345
	v359 = v353 + v358
	v360 = v357 + v359
	v364 = v358 - v357 ^ base.I32_rotl(v357, int32(8))
	v368 = v359 - v364 ^ base.I32_rotl(v364, int32(16))
	v372 = v360 - v368 ^ base.I32_rotl(v368, int32(19))
	v373 = v364 + v360
	v374 = v368 + v373
	v375 = v372 + v374
	v379 = v373 - v372 ^ base.I32_rotl(v372, v351)
	v380 = int32(12)
	v381 = v338 + v380
	v383 = v339 - v380
	if base.Ui32(int32(11)) < base.Ui32(v383) {
		v338 = v381
		v339 = v383
		v340 = v374
		v341 = v375
		v342 = v379
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v386 = v381
	v387 = v383
	v388 = v374
	v389 = v375
	v390 = v379
	goto L66
L72:
	;
	goto L71
L73:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	v558 = v437 + v438
	v559 = v389
	v560 = v390
	goto L59
L74:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+1)))
	v437 = v433<<(uint(int32(8))%32) + v432
	goto L73
L75:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+2)))
	v432 = v428<<(uint(int32(16))%32) + v388
	goto L74
L76:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v558 = v426 + v388
	v559 = v425
	v560 = v390
	goto L59
L77:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+4)))
	v425 = v422 + v423
	goto L76
L78:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+5)))
	v422 = v418<<(uint(int32(8))%32) + v417
	goto L77
L79:
	;
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+6)))
	v417 = v413<<(uint(int32(16))%32) + v389
	goto L78
L80:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v386)))
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	v558 = v409 + v388
	v559 = v411 + v389
	v560 = v408
	goto L59
L81:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+8)))
	v408 = v404<<(uint(int32(8))%32) + v403
	goto L80
L82:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+9)))
	v403 = v399<<(uint(int32(16))%32) + v398
	goto L81
L83:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+10)))
	v398 = v394<<(uint(int32(24))%32) + v390
	goto L82
L84:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v447 = v446 + v443
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v440)+8))
	v451 = v450 + v444
	v453 = int32(4)
	v455 = v448 + v442 - v451 ^ base.I32_rotl(v451, v453)
	v459 = v447 - v455 ^ base.I32_rotl(v455, int32(6))
	v460 = v451 + v447
	v461 = v455 + v460
	v462 = v459 + v461
	v466 = v460 - v459 ^ base.I32_rotl(v459, int32(8))
	v470 = v461 - v466 ^ base.I32_rotl(v466, int32(16))
	v474 = v462 - v470 ^ base.I32_rotl(v470, int32(19))
	v475 = v466 + v462
	v476 = v470 + v475
	v477 = v474 + v476
	v481 = v475 - v474 ^ base.I32_rotl(v474, v453)
	v482 = int32(12)
	v483 = v440 + v482
	v485 = v441 - v482
	if base.Ui32(int32(11)) < base.Ui32(v485) {
		v440 = v483
		v441 = v485
		v442 = v476
		v443 = v477
		v444 = v481
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v488 = v483
	v489 = v485
	v490 = v476
	v491 = v477
	v492 = v481
	goto L60
L86:
	;
	goto L85
L87:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488))))
	v558 = v551 + v554
	v559 = v552
	v560 = v553
	goto L59
L88:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+1)))
	v551 = v547<<(uint(int32(8))%32) + v544
	v552 = v545
	v553 = v546
	goto L87
L89:
	;
	v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+2)))
	v544 = v540<<(uint(int32(16))%32) + v537
	v545 = v538
	v546 = v539
	goto L88
L90:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+3)))
	v537 = v533<<(uint(int32(24))%32) + v490
	v538 = v531
	v539 = v532
	goto L89
L91:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+4)))
	v531 = v527 + v529
	v532 = v528
	goto L90
L92:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+5)))
	v527 = v523<<(uint(int32(8))%32) + v521
	v528 = v522
	goto L91
L93:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+6)))
	v521 = v517<<(uint(int32(16))%32) + v515
	v522 = v516
	goto L92
L94:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+7)))
	v515 = v511<<(uint(int32(24))%32) + v491
	v516 = v510
	goto L93
L95:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+8)))
	v510 = v506<<(uint(int32(8))%32) + v505
	goto L94
L96:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+9)))
	v505 = v501<<(uint(int32(16))%32) + v500
	goto L95
L97:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+10)))
	v500 = v496<<(uint(int32(24))%32) + v492
	goto L96
L98:
	;
	v887 = F_strlen(m, v45)
	mBase = m.M
	v889 = v887 + int32(1)
	v895 = v889 - int32(1636608432)
	if v45&int32(3) != 0 {
		goto L156
	} else {
		goto L157
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L101
	} else {
		goto L149
	}
L100:
	;
	v592 = F_toast_raw_datum_size(m, v45)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	return int32(0)
L102:
	;
	v596 = F_pg_detoast_datum_packed(m, v45)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v598 = int32(1)
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v600&v598 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v603 = v598
	goto L106
L105:
	;
	v603 = int32(4)
	goto L106
L106:
	;
	v604 = v596 + v603
	v606 = v592 - int32(4)
	v612 = v606 - int32(1636608432)
	if v604&int32(3) != 0 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	if v45 == v596 {
		v1154 = v870
		goto L11
	} else {
		goto L147
	}
L108:
	;
	v844 = int32(14)
	v846 = v840 ^ v841 - base.I32_rotl(v840, v844)
	v850 = v846 ^ v839 - base.I32_rotl(v846, int32(11))
	v854 = v850 ^ v840 - base.I32_rotl(v850, int32(25))
	v858 = v854 ^ v846 - base.I32_rotl(v854, int32(16))
	v862 = v858 ^ v850 - base.I32_rotl(v858, int32(4))
	v866 = v862 ^ v854 - base.I32_rotl(v862, v844)
	v870 = v866 ^ v858 - base.I32_rotl(v866, int32(24))
	goto L107
L109:
	;
	switch v770 - int32(1) {
	case 0:
		v832 = v771
		v833 = v772
		v834 = v773
		goto L136
	case 1:
		v825 = v771
		v826 = v772
		v827 = v773
		goto L137
	case 2:
		v818 = v771
		v819 = v772
		v820 = v773
		goto L138
	case 3:
		v812 = v772
		v813 = v773
		goto L139
	case 4:
		v808 = v772
		v809 = v773
		goto L140
	case 5:
		v802 = v772
		v803 = v773
		goto L141
	case 6:
		v796 = v772
		v797 = v773
		goto L142
	case 7:
		v791 = v773
		goto L143
	case 8:
		v786 = v773
		goto L144
	case 9:
		v781 = v773
		goto L145
	case 10:
		goto L146
	default:
		v839 = v771
		v840 = v772
		v841 = v773
		goto L108
	}
L110:
	;
	v721 = v604
	v722 = v606
	v723 = v612
	v724 = v612
	v725 = v612
	goto L133
L111:
	;
	if base.Ui32(int32(11)) < base.Ui32(v606) {
		goto L110
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	if base.Ui32(v606) < base.Ui32(int32(12)) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v769 = v604
	v770 = v606
	v771 = v612
	v772 = v612
	v773 = v612
	goto L109
L115:
	;
	switch v668 - int32(1) {
	case 0:
		v718 = v669
		goto L122
	case 1:
		v713 = v669
		goto L123
	case 2:
		goto L124
	case 3:
		v706 = v670
		goto L125
	case 4:
		v703 = v670
		goto L126
	case 5:
		v698 = v670
		goto L127
	case 6:
		goto L128
	case 7:
		v689 = v671
		goto L129
	case 8:
		v684 = v671
		goto L130
	case 9:
		v679 = v671
		goto L131
	case 10:
		goto L132
	default:
		v839 = v669
		v840 = v670
		v841 = v671
		goto L108
	}
L116:
	;
	v667 = v604
	v668 = v606
	v669 = v612
	v670 = v612
	v671 = v612
	goto L115
L117:
	;
	goto L118
L118:
	;
	v619 = v604
	v620 = v606
	v621 = v612
	v622 = v612
	v623 = v612
	goto L119
L119:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	v626 = v625 + v622
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v619)))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v619)+8))
	v630 = v629 + v623
	v632 = int32(4)
	v634 = v627 + v621 - v630 ^ base.I32_rotl(v630, v632)
	v638 = v626 - v634 ^ base.I32_rotl(v634, int32(6))
	v639 = v630 + v626
	v640 = v634 + v639
	v641 = v638 + v640
	v645 = v639 - v638 ^ base.I32_rotl(v638, int32(8))
	v649 = v640 - v645 ^ base.I32_rotl(v645, int32(16))
	v653 = v641 - v649 ^ base.I32_rotl(v649, int32(19))
	v654 = v645 + v641
	v655 = v649 + v654
	v656 = v653 + v655
	v660 = v654 - v653 ^ base.I32_rotl(v653, v632)
	v661 = int32(12)
	v662 = v619 + v661
	v664 = v620 - v661
	if base.Ui32(int32(11)) < base.Ui32(v664) {
		v619 = v662
		v620 = v664
		v621 = v655
		v622 = v656
		v623 = v660
		goto L119
	} else {
		goto L121
	}
L120:
	;
	v667 = v662
	v668 = v664
	v669 = v655
	v670 = v656
	v671 = v660
	goto L115
L121:
	;
	goto L120
L122:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	v839 = v718 + v719
	v840 = v670
	v841 = v671
	goto L108
L123:
	;
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+1)))
	v718 = v714<<(uint(int32(8))%32) + v713
	goto L122
L124:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+2)))
	v713 = v709<<(uint(int32(16))%32) + v669
	goto L123
L125:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v839 = v707 + v669
	v840 = v706
	v841 = v671
	goto L108
L126:
	;
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+4)))
	v706 = v703 + v704
	goto L125
L127:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+5)))
	v703 = v699<<(uint(int32(8))%32) + v698
	goto L126
L128:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+6)))
	v698 = v694<<(uint(int32(16))%32) + v670
	goto L127
L129:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v667)+4))
	v839 = v690 + v669
	v840 = v692 + v670
	v841 = v689
	goto L108
L130:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+8)))
	v689 = v685<<(uint(int32(8))%32) + v684
	goto L129
L131:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+9)))
	v684 = v680<<(uint(int32(16))%32) + v679
	goto L130
L132:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+10)))
	v679 = v675<<(uint(int32(24))%32) + v671
	goto L131
L133:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v721)+4))
	v728 = v727 + v724
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v721)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v721)+8))
	v732 = v731 + v725
	v734 = int32(4)
	v736 = v729 + v723 - v732 ^ base.I32_rotl(v732, v734)
	v740 = v728 - v736 ^ base.I32_rotl(v736, int32(6))
	v741 = v732 + v728
	v742 = v736 + v741
	v743 = v740 + v742
	v747 = v741 - v740 ^ base.I32_rotl(v740, int32(8))
	v751 = v742 - v747 ^ base.I32_rotl(v747, int32(16))
	v755 = v743 - v751 ^ base.I32_rotl(v751, int32(19))
	v756 = v747 + v743
	v757 = v751 + v756
	v758 = v755 + v757
	v762 = v756 - v755 ^ base.I32_rotl(v755, v734)
	v763 = int32(12)
	v764 = v721 + v763
	v766 = v722 - v763
	if base.Ui32(int32(11)) < base.Ui32(v766) {
		v721 = v764
		v722 = v766
		v723 = v757
		v724 = v758
		v725 = v762
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v769 = v764
	v770 = v766
	v771 = v757
	v772 = v758
	v773 = v762
	goto L109
L135:
	;
	goto L134
L136:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769))))
	v839 = v832 + v835
	v840 = v833
	v841 = v834
	goto L108
L137:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+1)))
	v832 = v828<<(uint(int32(8))%32) + v825
	v833 = v826
	v834 = v827
	goto L136
L138:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+2)))
	v825 = v821<<(uint(int32(16))%32) + v818
	v826 = v819
	v827 = v820
	goto L137
L139:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+3)))
	v818 = v814<<(uint(int32(24))%32) + v771
	v819 = v812
	v820 = v813
	goto L138
L140:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+4)))
	v812 = v808 + v810
	v813 = v809
	goto L139
L141:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+5)))
	v808 = v804<<(uint(int32(8))%32) + v802
	v809 = v803
	goto L140
L142:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+6)))
	v802 = v798<<(uint(int32(16))%32) + v796
	v803 = v797
	goto L141
L143:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+7)))
	v796 = v792<<(uint(int32(24))%32) + v772
	v797 = v791
	goto L142
L144:
	;
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+8)))
	v791 = v787<<(uint(int32(8))%32) + v786
	goto L143
L145:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+9)))
	v786 = v782<<(uint(int32(16))%32) + v781
	goto L144
L146:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769)+10)))
	v781 = v777<<(uint(int32(24))%32) + v773
	goto L145
L147:
	;
	F_pfree(m, v596)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L101
	} else {
		goto L148
	}
L148:
	;
	v1154 = v870
	goto L11
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v51
	F_errmsg_internal(m, int32(482620), v54)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L101
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(496448), int32(372), int32(322667))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L101
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v1154 = v1149 ^ v1141 - base.I32_rotl(v1149, int32(24))
	goto L11
L153:
	;
	v1127 = int32(14)
	v1129 = v1123 ^ v1124 - base.I32_rotl(v1123, v1127)
	v1133 = v1129 ^ v1122 - base.I32_rotl(v1129, int32(11))
	v1137 = v1133 ^ v1123 - base.I32_rotl(v1133, int32(25))
	v1141 = v1137 ^ v1129 - base.I32_rotl(v1137, int32(16))
	v1145 = v1141 ^ v1133 - base.I32_rotl(v1141, int32(4))
	v1149 = v1145 ^ v1137 - base.I32_rotl(v1145, v1127)
	goto L152
L154:
	;
	switch v1053 - int32(1) {
	case 0:
		v1115 = v1054
		v1116 = v1055
		v1117 = v1056
		goto L181
	case 1:
		v1108 = v1054
		v1109 = v1055
		v1110 = v1056
		goto L182
	case 2:
		v1101 = v1054
		v1102 = v1055
		v1103 = v1056
		goto L183
	case 3:
		v1095 = v1055
		v1096 = v1056
		goto L184
	case 4:
		v1091 = v1055
		v1092 = v1056
		goto L185
	case 5:
		v1085 = v1055
		v1086 = v1056
		goto L186
	case 6:
		v1079 = v1055
		v1080 = v1056
		goto L187
	case 7:
		v1074 = v1056
		goto L188
	case 8:
		v1069 = v1056
		goto L189
	case 9:
		v1064 = v1056
		goto L190
	case 10:
		goto L191
	default:
		v1122 = v1054
		v1123 = v1055
		v1124 = v1056
		goto L153
	}
L155:
	;
	v1004 = v45
	v1005 = v889
	v1006 = v895
	v1007 = v895
	v1008 = v895
	goto L178
L156:
	;
	if base.Ui32(int32(11)) < base.Ui32(v889) {
		goto L155
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	if base.Ui32(v889) < base.Ui32(int32(12)) {
		goto L161
	} else {
		goto L162
	}
L159:
	;
	v1052 = v45
	v1053 = v889
	v1054 = v895
	v1055 = v895
	v1056 = v895
	goto L154
L160:
	;
	switch v951 - int32(1) {
	case 0:
		v1001 = v952
		goto L167
	case 1:
		v996 = v952
		goto L168
	case 2:
		goto L169
	case 3:
		v989 = v953
		goto L170
	case 4:
		v986 = v953
		goto L171
	case 5:
		v981 = v953
		goto L172
	case 6:
		goto L173
	case 7:
		v972 = v954
		goto L174
	case 8:
		v967 = v954
		goto L175
	case 9:
		v962 = v954
		goto L176
	case 10:
		goto L177
	default:
		v1122 = v952
		v1123 = v953
		v1124 = v954
		goto L153
	}
L161:
	;
	v950 = v45
	v951 = v889
	v952 = v895
	v953 = v895
	v954 = v895
	goto L160
L162:
	;
	goto L163
L163:
	;
	v902 = v45
	v903 = v889
	v904 = v895
	v905 = v895
	v906 = v895
	goto L164
L164:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v902)+4))
	v909 = v908 + v905
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v902)))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v902)+8))
	v913 = v912 + v906
	v915 = int32(4)
	v917 = v910 + v904 - v913 ^ base.I32_rotl(v913, v915)
	v921 = v909 - v917 ^ base.I32_rotl(v917, int32(6))
	v922 = v913 + v909
	v923 = v917 + v922
	v924 = v921 + v923
	v928 = v922 - v921 ^ base.I32_rotl(v921, int32(8))
	v932 = v923 - v928 ^ base.I32_rotl(v928, int32(16))
	v936 = v924 - v932 ^ base.I32_rotl(v932, int32(19))
	v937 = v928 + v924
	v938 = v932 + v937
	v939 = v936 + v938
	v943 = v937 - v936 ^ base.I32_rotl(v936, v915)
	v944 = int32(12)
	v945 = v902 + v944
	v947 = v903 - v944
	if base.Ui32(int32(11)) < base.Ui32(v947) {
		v902 = v945
		v903 = v947
		v904 = v938
		v905 = v939
		v906 = v943
		goto L164
	} else {
		goto L166
	}
L165:
	;
	v950 = v945
	v951 = v947
	v952 = v938
	v953 = v939
	v954 = v943
	goto L160
L166:
	;
	goto L165
L167:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950))))
	v1122 = v1001 + v1002
	v1123 = v953
	v1124 = v954
	goto L153
L168:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+1)))
	v1001 = v997<<(uint(int32(8))%32) + v996
	goto L167
L169:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+2)))
	v996 = v992<<(uint(int32(16))%32) + v952
	goto L168
L170:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v1122 = v990 + v952
	v1123 = v989
	v1124 = v954
	goto L153
L171:
	;
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+4)))
	v989 = v986 + v987
	goto L170
L172:
	;
	v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+5)))
	v986 = v982<<(uint(int32(8))%32) + v981
	goto L171
L173:
	;
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+6)))
	v981 = v977<<(uint(int32(16))%32) + v953
	goto L172
L174:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	v1122 = v973 + v952
	v1123 = v975 + v953
	v1124 = v972
	goto L153
L175:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+8)))
	v972 = v968<<(uint(int32(8))%32) + v967
	goto L174
L176:
	;
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+9)))
	v967 = v963<<(uint(int32(16))%32) + v962
	goto L175
L177:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+10)))
	v962 = v958<<(uint(int32(24))%32) + v954
	goto L176
L178:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	v1011 = v1010 + v1007
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	v1015 = v1014 + v1008
	v1017 = int32(4)
	v1019 = v1012 + v1006 - v1015 ^ base.I32_rotl(v1015, v1017)
	v1023 = v1011 - v1019 ^ base.I32_rotl(v1019, int32(6))
	v1024 = v1015 + v1011
	v1025 = v1019 + v1024
	v1026 = v1023 + v1025
	v1030 = v1024 - v1023 ^ base.I32_rotl(v1023, int32(8))
	v1034 = v1025 - v1030 ^ base.I32_rotl(v1030, int32(16))
	v1038 = v1026 - v1034 ^ base.I32_rotl(v1034, int32(19))
	v1039 = v1030 + v1026
	v1040 = v1034 + v1039
	v1041 = v1038 + v1040
	v1045 = v1039 - v1038 ^ base.I32_rotl(v1038, v1017)
	v1046 = int32(12)
	v1047 = v1004 + v1046
	v1049 = v1005 - v1046
	if base.Ui32(int32(11)) < base.Ui32(v1049) {
		v1004 = v1047
		v1005 = v1049
		v1006 = v1040
		v1007 = v1041
		v1008 = v1045
		goto L178
	} else {
		goto L180
	}
L179:
	;
	v1052 = v1047
	v1053 = v1049
	v1054 = v1040
	v1055 = v1041
	v1056 = v1045
	goto L154
L180:
	;
	goto L179
L181:
	;
	v1118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052))))
	v1122 = v1115 + v1118
	v1123 = v1116
	v1124 = v1117
	goto L153
L182:
	;
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+1)))
	v1115 = v1111<<(uint(int32(8))%32) + v1108
	v1116 = v1109
	v1117 = v1110
	goto L181
L183:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+2)))
	v1108 = v1104<<(uint(int32(16))%32) + v1101
	v1109 = v1102
	v1110 = v1103
	goto L182
L184:
	;
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+3)))
	v1101 = v1097<<(uint(int32(24))%32) + v1054
	v1102 = v1095
	v1103 = v1096
	goto L183
L185:
	;
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+4)))
	v1095 = v1091 + v1093
	v1096 = v1092
	goto L184
L186:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+5)))
	v1091 = v1087<<(uint(int32(8))%32) + v1085
	v1092 = v1086
	goto L185
L187:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+6)))
	v1085 = v1081<<(uint(int32(16))%32) + v1079
	v1086 = v1080
	goto L186
L188:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+7)))
	v1079 = v1075<<(uint(int32(24))%32) + v1055
	v1080 = v1074
	goto L187
L189:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+8)))
	v1074 = v1070<<(uint(int32(8))%32) + v1069
	goto L188
L190:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+9)))
	v1069 = v1065<<(uint(int32(16))%32) + v1064
	goto L189
L191:
	;
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1052)+10)))
	v1064 = v1060<<(uint(int32(24))%32) + v1056
	goto L190
L192:
	;
	goto L7
L193:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
	v1173 = int32(0)
	v1174 = v2
	goto L194
L194:
	;
	v1183 = base.I32_rotl(v1174, int32(1))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184+v1173))))
	if v1186 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1208 = v1202
	goto L1
L196:
	;
	v1193 = v1173 << (uint(int32(2)) % 32)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1193)))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1196+v1193)))
	v1199 = F_FunctionCall1Coll(m, v1171+v1173*int32(28), v1195, v1198)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L101
	} else {
		goto L199
	}
L197:
	;
	v1202 = v1183
	goto L198
L198:
	;
	v1205 = v1173 + int32(1)
	if v1205 != v13 {
		v1173 = v1205
		v1174 = v1202
		goto L194
	} else {
		goto L200
	}
L199:
	;
	v1202 = v1199 ^ v1183
	goto L198
L200:
	;
	goto L195
}
func F_ModifyWaitEvent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = v6 + l1<<(uint(int32(4))%32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(16) {
		switch l2 - int32(16) {
		case 0, 16:
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(base.B2i32(base.Ui32(int32(31)) < base.Ui32(l2)))
			return
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(91046), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(492668), int32(680), int32(91237))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if l2 == v10 {
			if l2&int32(1) == int32(0) {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v36 != l3 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
					v41 = int32(1)
					if l2 == v41 {
						if l3 != 0 {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v46 = *(*int32)(unsafe.Add(mBase, _consts[450]))
							if v44 != v46 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(129237), int32(0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										F_errfinish(m, int32(492668), int32(704), int32(91237))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
								return
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
							return
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v53 = v49 + v50<<(uint(int32(3))%32)
						v54 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v53)+6)) = uint16(v54)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						switch v58 - int32(1) {
						case 0, 15:
							v85 = v41
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v85)
						default:
							v61 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v61)
							v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
							v64 = int32(1)
							v67 = int32(base.Ui32(v63)>>(uint(v64)%32)) & v64
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v67)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v69&int32(4) != 0 {
								v73 = v67 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v73)
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								v76 = v73
								v77 = v75
							} else {
								v76 = v67
								v77 = v69
							}
							if v77&int32(128) == int32(0) {
							} else {
								v85 = v76 | int32(8192)
								*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v85)
							}
						}
						return
					}
				} else {
					return
				}
			}
		} else {
			if v10&int32(1) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(91083), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						F_errfinish(m, int32(492668), int32(696), int32(91237))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
				v41 = int32(1)
				if l2 == v41 {
					if l3 != 0 {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v46 = *(*int32)(unsafe.Add(mBase, _consts[450]))
						if v44 != v46 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(129237), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									F_errfinish(m, int32(492668), int32(704), int32(91237))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l3
						return
					}
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v53 = v49 + v50<<(uint(int32(3))%32)
					v54 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v53)+6)) = uint16(v54)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v53))) = v56
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					switch v58 - int32(1) {
					case 0, 15:
						v85 = v41
						*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v85)
					default:
						v61 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v61)
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
						v64 = int32(1)
						v67 = int32(base.Ui32(v63)>>(uint(v64)%32)) & v64
						*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v67)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						if v69&int32(4) != 0 {
							v73 = v67 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v73)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v76 = v73
							v77 = v75
						} else {
							v76 = v67
							v77 = v69
						}
						if v77&int32(128) == int32(0) {
						} else {
							v85 = v76 | int32(8192)
							*(*uint16)(unsafe.Add(mBase, uint32(v53)+4)) = uint16(v85)
						}
					}
					return
				}
			}
		}
	}
}
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	if l2 != 0 {
		base.MemoryCopy(m, l0, l1, l2)
	} else {
	}
	return l0
}
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	base.MemoryFill(m, l0, base.I32_extend8_s(l1), l2)
	return l0
}
func F___mmap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	if base.Ui32(int32(2147483647)) <= base.Ui32(l0) {
		*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(48)
		return int32(-1)
	} else {
		if l1&int32(16) != 0 {
			v18 = int32(-63)
		} else {
			v18 = int32(-48)
		}
		if l1&int32(32) != 0 {
			v21 = int32(65536)
			v25 = (l0 + int32(15)) & int32(-16)
			v27 = v25 + int32(40)
			if base.Ui32(int32(-65600)) <= base.Ui32(v27) {
				*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(48)
				v191 = int32(0)
			} else {
				if base.Ui32(v27) < base.Ui32(int32(11)) {
					v78 = int32(16)
				} else {
					v78 = (v25 + int32(51)) & int32(-8)
				}
				v82 = F_emscripten_builtin_malloc(m, v78+int32(65548))
				mBase = m.M
				if v82 == int32(0) {
					v191 = int32(0)
				} else {
					v86 = v82 - int32(8)
					if int32(65535)&v82 == int32(0) {
						v146 = v86
					} else {
						v93 = v82 - int32(4)
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
						v104 = (v21+v82-int32(1))&int32(-65536) - int32(8)
						if base.Ui32(v104-v86) <= base.Ui32(int32(15)) {
							v109 = v21
						} else {
							v109 = int32(0)
						}
						v110 = v104 + v109
						v111 = v110 - v86
						v112 = v94&int32(-8) - v111
						if v94&int32(3) == int32(0) {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
							*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v110))) = v117 + v111
							v146 = v110
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
							v122 = int32(1)
							v125 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v112 | v121&v122 | v125
							v128 = v110 + v112
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v128)+4)) = v129 | v122
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
							*(*int32)(unsafe.Add(mBase, uint32(v93))) = v111 | v133&v122 | v125
							v140 = v86 + v111
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v140)+4)) = v141 | v122
							F_dispose_chunk(m, v86, v111)
							mBase = m.M
							v146 = v110
						}
					}
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
					if v152&int32(3) == int32(0) {
					} else {
						v158 = v152 & int32(-8)
						if base.Ui32(v158) <= base.Ui32(v78+int32(16)) {
						} else {
							v162 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v78 | v152&v162 | int32(2)
							v168 = v146 + v78
							v169 = v158 - v78
							*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v169 | int32(3)
							v173 = v146 + v158
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v174 | v162
							F_dispose_chunk(m, v168, v169)
							mBase = m.M
						}
					}
					v191 = v146 + int32(8)
				}
			}
			if v191 != 0 {
				v214 = F__emscripten_memset_bulkmem(m, v191, base.I32_extend8_s(int32(0)), v25)
				mBase = m.M
				v215 = v191 + v25
				*(*int32)(unsafe.Add(mBase, uint32(v215))) = v191
				*(*int64)(unsafe.Add(mBase, uint32(v215)+8)) = int64(-4294967295)
				v220 = v215
				*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = int32(3)
				*(*int64)(unsafe.Add(mBase, uint32(v220)+24)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = l0
				v228 = int32(4654216)
				v229 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
				*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = v229
				*(*int32)(unsafe.Add(mBase, _consts[1469])) = v220
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
				v237 = v233
			} else {
				v237 = int32(-48)
			}
		} else {
			v204 = F_emscripten_builtin_malloc(m, int32(40))
			mBase = m.M
			v207 = m.Env.X_mmap_js(m, l0, int32(3), l1, l2, int64(0), v204+int32(8), v204)
			mBase = m.M
			if int32(0) <= v207 {
				*(*int32)(unsafe.Add(mBase, uint32(v204)+12)) = l2
				v220 = v204
				*(*int32)(unsafe.Add(mBase, uint32(v220)+32)) = int32(3)
				*(*int64)(unsafe.Add(mBase, uint32(v220)+24)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v220)+16)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = l0
				v228 = int32(4654216)
				v229 = *(*int32)(unsafe.Add(mBase, _consts[1469]))
				*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = v229
				*(*int32)(unsafe.Add(mBase, _consts[1469])) = v220
				v233 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
				v237 = v233
			} else {
				F_emscripten_builtin_free(m, v204)
				mBase = m.M
				v237 = v207
			}
		}
		if l1&int32(32) != 0 {
			v241 = v18
		} else {
			v241 = int32(-63)
		}
		if v237 != int32(-63) {
			v244 = v237
		} else {
			v244 = v241
		}
		if base.Ui32(int32(-4095)) <= base.Ui32(v244) {
			*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0) - v244
			v252 = int32(-1)
		} else {
			v252 = v244
		}
		return v252
	}
}
func F_macaddr8tomacaddr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc0(m, int32(6))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+3)))
		if v9 == int32(255) {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+4)))
			if v12 == int32(254) {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
				*(*uint8)(unsafe.Add(mBase, uint32(v5))) = uint8(v35)
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)) = uint8(v37)
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)))
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)) = uint8(v39)
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+5)))
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)) = uint8(v41)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v43)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+7)))
				*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)) = uint8(v45)
				return v5
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(228910), int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(595159), int32(0))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499586), int32(559), int32(228739))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(228910), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(595159), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499586), int32(559), int32(228739))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
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
		}
	}
}
func F_makeA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_palloc0(m, int32(32))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(71)
		return v8
	}
}
func F_makeConfigurationDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3602)
	v17 = v13 + v14
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = F_deleteDependencyRecordsFor(m, int32(3602), v18, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v30 = F_new_object_addresses(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	F_deleteSharedDependencyRecordsFor(m, int32(3602), v18, int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(2615)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v34
	F_add_exact_object_address(m, v9+int32(-12), v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	F_recordDependencyOnOwner(m, int32(3602), v18, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	F_recordDependencyOnCurrentExtension(m, l0, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(3601)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v50
	F_add_exact_object_address(m, v9+int32(-12), v30)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_record_object_address_dependencies(m, l0, v30, int32(110))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v9+int32(-60), int32(1), int32(3), int32(184), v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v69 = int32(1)
	v74 = F_systable_beginscan(m, l3, int32(3609), v69, int32(0), v69, v9+int32(-60))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v76 = F_systable_getnext(m, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v76 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v79 = v76
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_systable_endscan(m, v74)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L27
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(3600)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86+v87)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v91
	F_add_exact_object_address(m, v9+int32(-12), v30)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v99 = F_systable_getnext(m, v74)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v99 != 0 {
		v79 = v99
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	goto L14
L28:
	;
	F_free_object_addresses(m, v30)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	m.G0 = v11 - int32(-64)
	return
}
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = l5
	v7 = l6
	v10 = F_palloc0(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(7)
		if l3 != int32(-1) {
			v20 = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(-1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v7)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v6)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
			return v10
		} else {
			if v6 != 0 {
				v20 = l4
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v7)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				return v10
			} else {
				v18 = F_pg_detoast_datum(m, l4)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v7)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v6)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
					return v10
				}
			}
		}
	}
}
func F_makeSimpleA_Expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_palloc0(m, int32(32))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(71)
		v20 = F_makeString(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v20
			v27 = F_list_make1_impl(m, int32(1), v10+int32(8))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v27
				m.G0 = v10 + int32(16)
				return v13
			}
		}
	}
}
func F_make_ands_implicit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v26 = v2
		m.G0 = v6 + int32(16)
		return v26
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v10 - int32(7) {
		case 0:
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v15 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
				v22 = F_list_make1_impl(m, int32(1), v6+int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v22
					m.G0 = v6 + int32(16)
					return v26
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 != 0 {
					v26 = v2
					m.G0 = v6 + int32(16)
					return v26
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
					v22 = F_list_make1_impl(m, int32(1), v6+int32(8))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = v22
						m.G0 = v6 + int32(16)
						return v26
					}
				}
			}
		default:
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
			v22 = F_list_make1_impl(m, int32(1), v6+int32(8))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = v22
				m.G0 = v6 + int32(16)
				return v26
			}
		case 14:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v13 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
				v22 = F_list_make1_impl(m, int32(1), v6+int32(8))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = v22
					m.G0 = v6 + int32(16)
					return v26
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v26 = v14
				m.G0 = v6 + int32(16)
				return v26
			}
		}
	}
}
func F_make_distinct_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v12 = F_make_op(m, l0, l1, l2, l3, v11, l4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		if v16 == int32(16) {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
			if v19 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(530947)
						F_errmsg(m, int32(106831), v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							F_parser_errposition(m, l0, l4)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(494448), int32(3103), int32(234826))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
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
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(18)
				m.G0 = v9 + int32(32)
				return v12
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(530947)
					F_errmsg(m, int32(283728), v9+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_parser_errposition(m, l0, l4)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494448), int32(3097), int32(234826))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
		}
	}
}
func F_make_greater_string(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v331 int32
	_ = v331
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(17) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	if int32(0) < v177 {
		goto L78
	} else {
		goto L79
	}
L2:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v169 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L3:
	;
	v162 = v159
	v163 = v159
	goto L2
L4:
	;
	v143 = v59 + int32(5)
	v144 = F_palloc(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L65
	}
L5:
	;
	v129 = F_palloc(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L57
	}
L6:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v128 = int32(base.Ui32(v122)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 == int32(19) {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	return int32(0)
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v28&int32(254) == int32(2) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v23&int32(1) == int32(0) {
		goto L6
	} else {
		goto L21
	}
L15:
	;
	v37 = v26
	goto L17
L16:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L17
L17:
	;
	if v28 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v40 = v26
	goto L20
L19:
	;
	v40 = v37
	goto L20
L20:
	;
	v128 = v40
	goto L5
L21:
	;
	v45 = int32(1)
	v128 = int32(base.Ui32(v23)>>(uint(v45)%32)) - v45
	goto L5
L22:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1078])))
	if v68 != 0 {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	v59 = F_strlen(m, v58)
	mBase = m.M
	if v59 != 0 {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v54 = F_DirectFunctionCall1Coll(m, int32(581), int32(0), v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v56 = F_text_to_cstring(m, v49)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L10
	} else {
		goto L28
	}
L27:
	;
	v58 = v54
	goto L23
L28:
	;
	v58 = v56
	goto L23
L29:
	;
	v60 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v162 = v65
	v163 = int32(0)
	goto L2
L32:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v62 != int32(1) {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v15 != int32(19) {
		goto L4
	} else {
		goto L51
	}
L35:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[1079]))
	if v70 == l2 {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v73 = int32(8009)
	v74 = int32(507609)
	v76 = int32(1)
	v79 = F_varstr_cmp(m, v74, v76, v73, v76, l2)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	if v79 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v83 = v73
	goto L42
L41:
	;
	v83 = v74
	goto L42
L42:
	;
	v84 = int32(1)
	v87 = F_varstr_cmp(m, v83, v84, int32(27054), v84, l2)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	if v87 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v91 = int32(27054)
	goto L46
L45:
	;
	v91 = v83
	goto L46
L46:
	;
	v92 = int32(1)
	v95 = F_varstr_cmp(m, v91, v92, int32(547044), v92, l2)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1079])) = l2
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v104 = int32(57)
	goto L50
L49:
	;
	v104 = v101
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[1078])) = uint8(v104)
	goto L34
L51:
	;
	v112 = F_palloc(m, v59+int32(2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	if v59 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1078])))
	v118 = v115 + v59
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)) = uint8(v119)
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v117)
	v159 = v112
	goto L3
L54:
	;
	v114 = F__emscripten_memcpy_bulkmem(m, v112, v58, v59)
	mBase = m.M
	v115 = v114
	goto L56
L55:
	;
	v115 = v112
	goto L56
L56:
	;
	goto L53
L57:
	;
	v131 = int32(1)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v133&v131 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v136 = v131
	goto L60
L59:
	;
	v136 = int32(4)
	goto L60
L60:
	;
	if v128 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v177 = v128
	v179 = v129
	v180 = v140
	v181 = int32(0)
	v182 = int32(1454)
	goto L1
L62:
	;
	v138 = F__emscripten_memcpy_bulkmem(m, v129, v19+v136, v128)
	mBase = m.M
	goto L64
L63:
	;
	goto L64
L64:
	;
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v143 << (uint(int32(2)) % 32)
	v150 = v144 + int32(4)
	if v59 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1078])))
	*(*uint8)(unsafe.Add(mBase, uint32(v152+v59))) = uint8(v155)
	v159 = v144
	goto L3
L67:
	;
	v151 = F__emscripten_memcpy_bulkmem(m, v150, v58, v59)
	mBase = m.M
	v152 = v151
	goto L69
L68:
	;
	v152 = v150
	goto L69
L69:
	;
	goto L66
L70:
	;
	v172 = int32(1654)
	goto L72
L71:
	;
	v172 = int32(1655)
	goto L72
L72:
	;
	if v169 == int32(6) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v175 = int32(1653)
	goto L75
L74:
	;
	v175 = v172
	goto L75
L75:
	;
	v177 = v59
	v179 = v58
	v180 = v162
	v181 = v163
	v182 = v175
	goto L1
L76:
	;
	F_pfree(m, v179)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L116
	}
L77:
	;
	F_pfree(m, v181)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L115
	}
L78:
	;
	v190 = v177
	goto L81
L79:
	;
	goto L80
L80:
	;
	v297 = int32(0)
	if v181 == v297 {
		v316 = v297
		goto L76
	} else {
		goto L114
	}
L81:
	;
	if base.B2i32(v15 == int32(17)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L80
L83:
	;
	v206 = F_pg_mbcliplen(m, v179, v190, v190-int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L86
	}
L84:
	;
	v209 = int32(1)
	goto L85
L85:
	;
	v211 = v190 + v179 - v209
	v212 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, v211, v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L87
	}
L86:
	;
	v209 = v190 - v206
	goto L85
L87:
	;
	if v212 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v215 = v190 + int32(4)
	goto L91
L89:
	;
	goto L90
L90:
	;
	v277 = v190 - v209
	v279 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179+v277))) = uint8(v279)
	if v279 < v277 {
		v190 = v277
		goto L81
	} else {
		goto L113
	}
L91:
	;
	if v15 == int32(17) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L90
L93:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	v254 = F_FunctionCall2Coll(m, l1, l2, v180, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L104
	}
L94:
	;
	v234 = F_palloc(m, v215)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v249 = F_string_to_const(m, v179, v15)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L103
	}
L97:
	;
	if v190 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v215 << (uint(int32(2)) % 32)
	v242 = int32(-1)
	v243 = int32(0)
	v247 = F_makeConst(m, int32(17), v242, v243, v242, v234, v243, v243)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L102
	}
L99:
	;
	v238 = F__emscripten_memcpy_bulkmem(m, v234+int32(4), v179, v190)
	mBase = m.M
	goto L101
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	v252 = v247
	goto L93
L103:
	;
	v252 = v249
	goto L93
L104:
	;
	if v254 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	if v181 != 0 {
		v300 = v252
		goto L77
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	F_pfree(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L10
	} else {
		goto L109
	}
L108:
	;
	v316 = v252
	goto L76
L109:
	;
	F_pfree(m, v252)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	v261 = m.T0[v182].(func(*base.Module, int32, int32) int32)(m, v211, v209)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	if v261 != 0 {
		goto L91
	} else {
		goto L112
	}
L112:
	;
	goto L92
L113:
	;
	goto L82
L114:
	;
	v300 = v297
	goto L77
L115:
	;
	v316 = v300
	goto L76
L116:
	;
	return v316
}
func F_make_parsestate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v4 = F_palloc0(m, int32(124))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+85)) = uint8(v8)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+72)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		if l0 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+104)) = v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+108)) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+112)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+116)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+120)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+88)) = v25
		} else {
		}
		return v4
	}
}
func F_make_pathkeys_for_sortclauses_extended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
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
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v221 int32
	_ = v221
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v24 == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v221
L2:
	;
	v221 = v8
	goto L1
L3:
	;
	goto L4
L4:
	;
	v38 = v8
	v39 = v8
	v40 = v24
	goto L5
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 <= v39 {
		v221 = v38
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v221 = v202
	goto L1
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v39<<(uint(int32(2))%32))))
	v51 = F_get_sortgroupclause_expr(m, v50, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v204 != 0 {
		v38 = v202
		v39 = v203 + int32(1)
		v40 = v204
		goto L5
	} else {
		goto L43
	}
L11:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v58)
	v202 = v38
	v203 = v39
	v204 = v40
	goto L10
L12:
	;
	goto L13
L13:
	;
	if l4 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v61 = F_bms_make_singleton(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	v67 = v51
	v68 = v55
	goto L16
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+17)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+16)))
	v78 = F_get_ordering_op_properties(m, v68, v20+int32(12), v20+int32(8), v20+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L20
	}
L17:
	;
	v64 = F_remove_nulling_relids(m, v51, v61, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v67 = v64
	v68 = v66
	goto L16
L19:
	;
	v189 = F_lappend(m, v38, v90)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L8
	} else {
		goto L42
	}
L20:
	;
	if v78 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = F_exprCollation(m, v67)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L39
	}
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v84 = int32(1)
	v90 = F_make_pathkey_from_sortinfo(m, l0, v67, v82, v83, v80, v71&v84, v70&v84, v69, int32(0), v84)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if l6 == int32(0) {
		v99 = v92
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+40)))
	if v100 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+44))
	if v95 != 0 {
		v99 = v92
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+44)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v99 = v98
	goto L26
L29:
	;
	if l3 == int32(0) {
		v202 = v38
		v203 = v39
		v204 = v40
		goto L10
	} else {
		goto L37
	}
L30:
	;
	if v38 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v103 <= int32(0) {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v115 = int32(0)
	goto L33
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v106+v115<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v99 == v129 {
		goto L29
	} else {
		goto L35
	}
L34:
	;
	goto L19
L35:
	;
	v132 = v115 + int32(1)
	if v132 != v103 {
		v115 = v132
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v154 = F_list_delete_nth_cell(m, v153, v39)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v154
	v202 = v38
	v203 = v39 - int32(1)
	v204 = v154
	goto L10
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v68
	F_errmsg_internal(m, int32(208952), v20)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(492971), int32(273), int32(234015))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v202 = v189
	v203 = v39
	v204 = v40
	goto L10
L43:
	;
	goto L6
}
func F_make_pathkeys_for_window(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v10 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L35
	} else {
		goto L49
	}
L2:
	;
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v55 = int32(1)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v19 <= int32(0) {
		v47 = int32(1)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v55 = v47
	goto L2
L7:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v19
	goto L10
L9:
	;
	v25 = v22
	goto L10
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v28 = int32(0)
	goto L11
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26+v28<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = int32(0)
	v39 = base.B2i32(v37 != v38)
	if v37 == v38 {
		v47 = v39
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v47 = v39
	goto L6
L13:
	;
	v43 = v28 + int32(1)
	if v43 != v25 {
		v28 = v43
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v56 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L35
	} else {
		goto L44
	}
L18:
	;
	if v101 == int32(0) {
		goto L1
	} else {
		goto L31
	}
L19:
	;
	v101 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v65 <= int32(0) {
		v93 = int32(1)
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v101 = v93
	goto L18
L23:
	;
	v68 = int32(0)
	if v68 < v65 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v71 = v65
	goto L26
L25:
	;
	v71 = v68
	goto L26
L26:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v74 = int32(0)
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v72+v74<<(uint(int32(2))%32))))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v84 = int32(0)
	v85 = base.B2i32(v83 != v84)
	if v83 == v84 {
		v93 = v85
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v93 = v85
	goto L22
L29:
	;
	v89 = v74 + int32(1)
	if v89 != v71 {
		v74 = v89
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v105 = l1 + int32(12)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v106 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v108 = int32(0)
	v112 = F_make_pathkeys_for_sortclauses_extended(m, l0, v105, l2, int32(1), v108, v8+int32(15), v108)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v117 = int32(0)
	goto L34
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v118 == int32(0) {
		v128 = v117
		goto L37
	} else {
		goto L38
	}
L35:
	;
	return int32(0)
L36:
	;
	v117 = v112
	goto L34
L37:
	;
	m.G0 = v8 + int32(16)
	return v128
L38:
	;
	v121 = F_make_pathkeys_for_sortclauses(m, l0, v118, l2)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v117 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v128 = v121
	goto L37
L41:
	;
	goto L42
L42:
	;
	v125 = F_append_pathkeys(m, v117, v121)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L35
	} else {
		goto L43
	}
L43:
	;
	v128 = v125
	goto L37
L44:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L35
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(509114), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_errdetail(m, int32(582686), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L35
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(494727), int32(6323), int32(31801))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L35
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L35
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(509014), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L35
	} else {
		goto L51
	}
L51:
	;
	F_errdetail(m, int32(582631), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L35
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(494727), int32(6328), int32(31801))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_restrictinfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	if l1 == int32(0) {
		v25 = F_make_plain_restrictinfo(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8, l9)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			return v25
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v13 != int32(21) {
			v25 = F_make_plain_restrictinfo(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8, l9)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v25
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v16 != int32(1) {
				v25 = F_make_plain_restrictinfo(m, l0, l1, int32(0), l2, l3, l4, l5, l6, l7, l8, l9)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					return v25
				}
			} else {
				v19 = F_make_sub_restrictinfos(m, l0, l1, l2, l3, l4, l5, l6, l7, l8, l9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			}
		}
	}
}
func F_make_scalar_array_op(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v3 = l2
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = F_exprType(m, l3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = F_exprType(m, l4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 != int32(705) {
				v26 = F_get_base_element_type(m, v22)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(413725), int32(0))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, l0, l5)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(494703), int32(804), int32(234711))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
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
					} else {
						v30 = v26
						v32 = F_oper(m, l0, l1, v18, v30, int32(0), l5)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
							v36 = v34 + v35
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
							if v37 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return int32(0)
									} else {
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
										v151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
										v152 = F_op_signature_string(m, l1, v150, v151)
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = v152
											F_errmsg(m, int32(202160), v15)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, l0, l5)
												mBase = m.M
												v159 = m.ExcPending
												if v159 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494703), int32(819), int32(234711))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
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
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l4
								*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
								v48 = F_list_make2_impl(m, v13+int32(-28), v13+int32(-32))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v30
									*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v18
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v52
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
									*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v54
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
									v63 = F_enforce_generic_type_consistency(m, v13+int32(-8), v13+int32(-16), int32(2), v61, int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										if v63 != int32(16) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v168 = m.ExcPending
											if v168 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(283674), int32(0))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v177 = m.ExcPending
														if v177 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494703), int32(845), int32(234711))
															mBase = m.M
															v182 = m.ExcPending
															if v182 != 0 {
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
										} else {
											v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
											v68 = F_get_func_retset(m, v67)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												if v68 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v189 = m.ExcPending
														if v189 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(107079), int32(0))
															mBase = m.M
															v193 = m.ExcPending
															if v193 != 0 {
																return int32(0)
															} else {
																F_parser_errposition(m, l0, l5)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(494703), int32(850), int32(234711))
																	mBase = m.M
																	v200 = m.ExcPending
																	if v200 != 0 {
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
												} else {
													v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
													if v70 <= int32(3830) {
														switch v70 - int32(2277) {
														case 0, 6:
															v93 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v103 = F_palloc0(m, int32(36))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v103
																	}
																}
															}
														case 1, 2, 3, 4, 5:
															v89 = F_get_array_type(m, v70)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																if v89 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v207 = m.ExcPending
																		if v207 != 0 {
																			return int32(0)
																		} else {
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v209 = F_format_type_be(m, v208)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																				F_errmsg(m, int32(193150), v13+int32(-48))
																				mBase = m.M
																				v216 = m.ExcPending
																				if v216 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(494703), int32(871), int32(234711))
																						mBase = m.M
																						v223 = m.ExcPending
																						if v223 != 0 {
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
																	}
																} else {
																	v93 = v89
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v103 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v103
																			}
																		}
																	}
																}
															}
														default:
															if v70 == int32(2776) {
																v93 = v22
																*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return int32(0)
																} else {
																	v103 = F_palloc0(m, int32(36))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																		v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																		v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																		*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																		*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																		F_ReleaseCatCache(m, v32)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v15 - int32(-64)
																			return v103
																		}
																	}
																}
															} else {
																if v70 != int32(3500) {
																	v89 = F_get_array_type(m, v70)
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int32(0)
																	} else {
																		if v89 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(67137668))
																				mBase = m.M
																				v207 = m.ExcPending
																				if v207 != 0 {
																					return int32(0)
																				} else {
																					v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																					v209 = F_format_type_be(m, v208)
																					mBase = m.M
																					v210 = m.ExcPending
																					if v210 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																						F_errmsg(m, int32(193150), v13+int32(-48))
																						mBase = m.M
																						v216 = m.ExcPending
																						if v216 != 0 {
																							return int32(0)
																						} else {
																							F_parser_errposition(m, l0, l5)
																							mBase = m.M
																							v218 = m.ExcPending
																							if v218 != 0 {
																								return int32(0)
																							} else {
																								F_errfinish(m, int32(494703), int32(871), int32(234711))
																								mBase = m.M
																								v223 = m.ExcPending
																								if v223 != 0 {
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
																			}
																		} else {
																			v93 = v89
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																			F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return int32(0)
																			} else {
																				v103 = F_palloc0(m, int32(36))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																					v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																					v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																					v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																					*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																					*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																					F_ReleaseCatCache(m, v32)
																					mBase = m.M
																					v120 = m.ExcPending
																					if v120 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v15 - int32(-64)
																						return v103
																					}
																				}
																			}
																		}
																	}
																} else {
																	v93 = v22
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v103 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v103
																			}
																		}
																	}
																}
															}
														}
													} else {
														if base.Ui32(v70-int32(5077)) < base.Ui32(int32(4)) {
															v93 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v103 = F_palloc0(m, int32(36))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v103
																	}
																}
															}
														} else {
															if base.Ui32(v70-int32(4537)) < base.Ui32(int32(2)) {
																v93 = v22
																*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
																	return int32(0)
																} else {
																	v103 = F_palloc0(m, int32(36))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																		v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																		v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																		v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																		*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																		*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																		F_ReleaseCatCache(m, v32)
																		mBase = m.M
																		v120 = m.ExcPending
																		if v120 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v15 - int32(-64)
																			return v103
																		}
																	}
																}
															} else {
																if v70 == int32(3831) {
																	v93 = v22
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v103 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v103
																			}
																		}
																	}
																} else {
																	v89 = F_get_array_type(m, v70)
																	mBase = m.M
																	v90 = m.ExcPending
																	if v90 != 0 {
																		return int32(0)
																	} else {
																		if v89 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(67137668))
																				mBase = m.M
																				v207 = m.ExcPending
																				if v207 != 0 {
																					return int32(0)
																				} else {
																					v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																					v209 = F_format_type_be(m, v208)
																					mBase = m.M
																					v210 = m.ExcPending
																					if v210 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																						F_errmsg(m, int32(193150), v13+int32(-48))
																						mBase = m.M
																						v216 = m.ExcPending
																						if v216 != 0 {
																							return int32(0)
																						} else {
																							F_parser_errposition(m, l0, l5)
																							mBase = m.M
																							v218 = m.ExcPending
																							if v218 != 0 {
																								return int32(0)
																							} else {
																								F_errfinish(m, int32(494703), int32(871), int32(234711))
																								mBase = m.M
																								v223 = m.ExcPending
																								if v223 != 0 {
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
																			}
																		} else {
																			v93 = v89
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																			F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return int32(0)
																			} else {
																				v103 = F_palloc0(m, int32(36))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																					v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																					v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																					v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																					*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																					*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																					F_ReleaseCatCache(m, v32)
																					mBase = m.M
																					v120 = m.ExcPending
																					if v120 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v15 - int32(-64)
																						return v103
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
									}
								}
							}
						}
					}
				}
			} else {
				v30 = int32(705)
				v32 = F_oper(m, l0, l1, v18, v30, int32(0), l5)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
					v36 = v34 + v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
					if v37 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
								v152 = F_op_signature_string(m, l1, v150, v151)
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v152
									F_errmsg(m, int32(202160), v15)
									mBase = m.M
									v157 = m.ExcPending
									if v157 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l5)
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(494703), int32(819), int32(234711))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
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
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = l4
						*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l4
						v48 = F_list_make2_impl(m, v13+int32(-28), v13+int32(-32))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v30
							*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v18
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v52
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
							*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v54
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+88))
							v63 = F_enforce_generic_type_consistency(m, v13+int32(-8), v13+int32(-16), int32(2), v61, int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								if v63 != int32(16) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v168 = m.ExcPending
									if v168 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(283674), int32(0))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, l0, l5)
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(494703), int32(845), int32(234711))
													mBase = m.M
													v182 = m.ExcPending
													if v182 != 0 {
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
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
									v68 = F_get_func_retset(m, v67)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										if v68 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v186 = m.ExcPending
											if v186 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(107079), int32(0))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(494703), int32(850), int32(234711))
															mBase = m.M
															v200 = m.ExcPending
															if v200 != 0 {
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
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
											if v70 <= int32(3830) {
												switch v70 - int32(2277) {
												case 0, 6:
													v93 = v22
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
													*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
													F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														v103 = F_palloc0(m, int32(36))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
															*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
															*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
															*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
															*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
															F_ReleaseCatCache(m, v32)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																m.G0 = v15 - int32(-64)
																return v103
															}
														}
													}
												case 1, 2, 3, 4, 5:
													v89 = F_get_array_type(m, v70)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int32(0)
													} else {
														if v89 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67137668))
																mBase = m.M
																v207 = m.ExcPending
																if v207 != 0 {
																	return int32(0)
																} else {
																	v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																	v209 = F_format_type_be(m, v208)
																	mBase = m.M
																	v210 = m.ExcPending
																	if v210 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																		F_errmsg(m, int32(193150), v13+int32(-48))
																		mBase = m.M
																		v216 = m.ExcPending
																		if v216 != 0 {
																			return int32(0)
																		} else {
																			F_parser_errposition(m, l0, l5)
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(494703), int32(871), int32(234711))
																				mBase = m.M
																				v223 = m.ExcPending
																				if v223 != 0 {
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
															}
														} else {
															v93 = v89
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v103 = F_palloc0(m, int32(36))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v103
																	}
																}
															}
														}
													}
												default:
													if v70 == int32(2776) {
														v93 = v22
														*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
														*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
														F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															v103 = F_palloc0(m, int32(36))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																F_ReleaseCatCache(m, v32)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v15 - int32(-64)
																	return v103
																}
															}
														}
													} else {
														if v70 != int32(3500) {
															v89 = F_get_array_type(m, v70)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																if v89 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v207 = m.ExcPending
																		if v207 != 0 {
																			return int32(0)
																		} else {
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v209 = F_format_type_be(m, v208)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																				F_errmsg(m, int32(193150), v13+int32(-48))
																				mBase = m.M
																				v216 = m.ExcPending
																				if v216 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(494703), int32(871), int32(234711))
																						mBase = m.M
																						v223 = m.ExcPending
																						if v223 != 0 {
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
																	}
																} else {
																	v93 = v89
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v103 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v103
																			}
																		}
																	}
																}
															}
														} else {
															v93 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v103 = F_palloc0(m, int32(36))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v103
																	}
																}
															}
														}
													}
												}
											} else {
												if base.Ui32(v70-int32(5077)) < base.Ui32(int32(4)) {
													v93 = v22
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
													*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
													F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														v103 = F_palloc0(m, int32(36))
														mBase = m.M
														v104 = m.ExcPending
														if v104 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
															v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
															v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
															*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
															*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
															*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
															*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
															F_ReleaseCatCache(m, v32)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return int32(0)
															} else {
																m.G0 = v15 - int32(-64)
																return v103
															}
														}
													}
												} else {
													if base.Ui32(v70-int32(4537)) < base.Ui32(int32(2)) {
														v93 = v22
														*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
														*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
														F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return int32(0)
														} else {
															v103 = F_palloc0(m, int32(36))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																F_ReleaseCatCache(m, v32)
																mBase = m.M
																v120 = m.ExcPending
																if v120 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v15 - int32(-64)
																	return v103
																}
															}
														}
													} else {
														if v70 == int32(3831) {
															v93 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return int32(0)
															} else {
																v103 = F_palloc0(m, int32(36))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																	v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v120 = m.ExcPending
																	if v120 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v103
																	}
																}
															}
														} else {
															v89 = F_get_array_type(m, v70)
															mBase = m.M
															v90 = m.ExcPending
															if v90 != 0 {
																return int32(0)
															} else {
																if v89 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v207 = m.ExcPending
																		if v207 != 0 {
																			return int32(0)
																		} else {
																			v208 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v209 = F_format_type_be(m, v208)
																			mBase = m.M
																			v210 = m.ExcPending
																			if v210 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v209
																				F_errmsg(m, int32(193150), v13+int32(-48))
																				mBase = m.M
																				v216 = m.ExcPending
																				if v216 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(494703), int32(871), int32(234711))
																						mBase = m.M
																						v223 = m.ExcPending
																						if v223 != 0 {
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
																	}
																} else {
																	v93 = v89
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v93
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		v103 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(20)
																			v107 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
																			v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v110
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v103)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v103)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v112
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v103
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
							}
						}
					}
				}
			}
		}
	}
}
func F_match_db_entries(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, _consts[107]))
	return base.B2i32(v3 == v5)
}
func F_matchingjoinsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(0.01))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_mda_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	if l0 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	v10 = int32(1)
	v11 = l0 - v10
	v13 = v11 << (uint(int32(2)) % 32)
	v14 = l1 + v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2+v13)))
	v20 = base.I32_rem_s(v15+v10, v19)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v20
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v48
L5:
	;
	v22 = v11
	v25 = v20
	goto L8
L6:
	;
	goto L7
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	if v25 != 0 {
		v48 = v22
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v27 = int32(1)
	v28 = v22 - v27
	v30 = v28 << (uint(int32(2)) % 32)
	v31 = l1 + v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+v30)))
	v37 = base.I32_rem_s(v32+v27, v36)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v37
	if v28 != 0 {
		v22 = v28
		v25 = v37
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v47 = int32(0)
	goto L14
L13:
	;
	v47 = int32(-1)
	goto L14
L14:
	;
	v48 = v47
	goto L4
}
func F_mdextend(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	if l2 != int32(-1) {
		v16 = F__mdfd_getseg(m, l0, l1, l2, l4, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(8192)
			v31 = F_FileWriteV(m, v18, v11+int32(120), int32(1), base.I64_extend_i32_u(l2<<(uint(int32(13))%32)&int32(1073733632)), int32(167772177))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				if v31 != int32(8192) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						if v31 < int32(0) {
							F_errcode_for_file_access(m)
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v110 = *(*int32)(unsafe.Add(mBase, _consts[551]))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v108*int32(48))+32))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v114
								F_errmsg(m, int32(299135), v11+int32(16))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									F_errhint(m, int32(628594), int32(0))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return
									} else {
										F_errfinish(m, int32(499060), int32(519), int32(425077))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_errcode(m, int32(4293))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v46 = *(*int32)(unsafe.Add(mBase, _consts[551]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44*int32(48))+32))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(8192)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v31
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v50
								F_errmsg(m, int32(47586), v11+int32(32))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errhint(m, int32(628594), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errfinish(m, int32(499060), int32(526), int32(425077))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l4 != 0 {
						m.G0 = v11 + int32(128)
						return
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v70 != int32(-1) {
							m.G0 = v11 + int32(128)
							return
						} else {
							F_register_dirty_segment(m, l0, l1, v16)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v11 + int32(128)
								return
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return
			} else {
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_GetRelationPath(m, v11+int32(48), v87, v88, v89, v90, l1)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(48)
					F_errmsg(m, int32(153690), v11)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						F_errfinish(m, int32(499060), int32(504), int32(425077))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_mdfd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = F_mdopenfork(m, l0, l1, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v12 = F__mdfd_getseg(m, l0, l1, l2, int32(0), int32(1))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = l2 << (uint(int32(13)) % 32) & int32(1073733632)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v20 = F_FileAccess(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if v20 < int32(0) {
					v31 = int32(-1)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[551]))
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v19*int32(48))))
					v31 = v30
				}
				return v31
			}
		}
	}
}
func F_mdinit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(213456), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[795])) = v8
		return
	}
}
func F_mdnblocks(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_mdopenfork(m, l0, l1, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = l0 + l1<<(uint(int32(2))%32)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v22 = v20 - int32(1)
	v28 = v22
	v29 = v19 + v22<<(uint(int32(3))%32)
	goto L6
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	m.G0 = v9 + int32(16)
	return v56
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v33 = F_FileSize(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v56 = v48 << (uint(int32(17)) % 32)
	goto L5
L8:
	;
	if v33 < int64(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v39 = base.I32_wrap_i64(int64(base.Ui64(v33) >> (uint(int64(13)) % 64)))
	if base.Ui32(int32(131073)) <= base.Ui32(v39) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v39 != int32(131072) {
		v56 = v28<<(uint(int32(17))%32) + v39
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v48 = v28 + int32(1)
	v50 = F__mdfd_openseg(m, l0, l1, v48, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v50 != 0 {
		v28 = v48
		v29 = v50
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[551]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v67*int32(48))+32))
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
	F_errmsg(m, int32(298762), v9)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(499060), int32(1882), int32(153510))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	F_errmsg_internal(m, int32(337033), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(499060), int32(1255), int32(153511))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_mdopen(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v2
	return
}
func F_mdopenfork(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v13 = l0 + l1<<(uint(int32(2))%32)
	v15 = v13 + int32(40)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if int32(0) < v16 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
		v99 = v19
		m.G0 = v9 + int32(80)
		return v99
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_GetRelationPath(m, v9+int32(8), v22, v23, v24, v25, l1)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _consts[624]))
			if v35&int32(1) != 0 {
				v38 = int32(16386)
			} else {
				v38 = int32(2)
			}
			v39 = F_PathNameOpenFile(m, v9+int32(8), v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v39 < int32(0) {
					if l2&int32(2) != 0 {
						v47 = *(*int32)(unsafe.Add(mBase, _consts[86]))
						if v47 == int32(44) {
							v99 = int32(0)
							m.G0 = v9 + int32(80)
							return v99
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(8)
									F_errmsg(m, int32(298247), v9)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499060), int32(686), int32(314654))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
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
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(8)
								F_errmsg(m, int32(298247), v9)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499060), int32(686), int32(314654))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
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
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					if v68 == int32(0) {
						v75 = *(*int32)(unsafe.Add(mBase, _consts[795]))
						v77 = F_MemoryContextAlloc(m, v75, int32(8))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56)) = v77
							v92 = v77
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
							v99 = v92
							m.G0 = v9 + int32(80)
							return v99
						}
					} else {
						v84 = l0 + l1<<(uint(int32(2))%32) + int32(56)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
						if int32(0) < v68 {
							v92 = v85
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
							v99 = v92
							m.G0 = v9 + int32(80)
							return v99
						} else {
							v89 = F_repalloc(m, v85, int32(8))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v84))) = v89
								v92 = v89
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v92))) = v39
								v99 = v92
								m.G0 = v9 + int32(80)
								return v99
							}
						}
					}
				}
			}
		}
	}
}
func F_mdregistersync(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	v10 = F_mdnblocks(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v16 = l0 + l1<<(uint(int32(2))%32) + int32(40)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v20 = v17
	goto L3
L3:
	;
	v30 = F__mdfd_openseg(m, l0, l1, v20, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	if int32(0) < v20 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	if v30 != 0 {
		v20 = v20 + int32(1)
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	v38 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v43 = v20
	goto L10
L8:
	;
	goto L9
L9:
	;
	return
L10:
	;
	v49 = v43 - int32(1)
	v51 = v49 << (uint(int32(3)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v53 = v51 + v52
	F_register_dirty_segment(m, l0, l1, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if v17 < v43 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	F_FileClose(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(int32(1)) < base.Ui32(v43) {
		v43 = v49
		goto L10
	} else {
		goto L30
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v49 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v49
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v79
	goto L17
L19:
	;
	if v60 <= int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v60 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_pfree(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v79 = int32(0)
	goto L18
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[795]))
	v73 = F_MemoryContextAlloc(m, v72, v51)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v49 <= v60 {
		goto L17
	} else {
		goto L28
	}
L27:
	;
	v79 = v73
	goto L18
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v77 = F_repalloc(m, v76, v51)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v79 = v77
	goto L18
L30:
	;
	goto L11
}
func F_memcmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	v39 = v34
	v40 = v35
	v41 = v36
	goto L12
L3:
	;
	if (l0|l1)&int32(3) != 0 {
		v34 = l0
		v35 = l1
		v36 = l2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v27 = l0
	v28 = l1
	v29 = l2
	goto L5
L5:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v11 = l0
	v12 = l1
	v13 = l2
	goto L7
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v16 != v17 {
		v34 = v11
		v35 = v12
		v36 = v13
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v27 = v22
	v28 = v20
	v29 = v24
	goto L5
L9:
	;
	v19 = int32(4)
	v20 = v12 + v19
	v22 = v11 + v19
	v24 = v13 - v19
	if base.Ui32(int32(3)) < base.Ui32(v24) {
		v11 = v22
		v12 = v20
		v13 = v24
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v34 = v27
	v35 = v28
	v36 = v29
	goto L2
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v44 == v45 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return v44 - v45
L14:
	;
	v47 = int32(1)
	v52 = v41 - v47
	if v52 != 0 {
		v39 = v39 + v47
		v40 = v40 + v47
		v41 = v52
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	goto L1
}
func F_merge_acl_with_grant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v85
L2:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if int32(0) < v17 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v85 = l0
	goto L1
L6:
	;
	goto L5
L7:
	;
	v22 = int32(1)
	goto L9
L8:
	;
	v22 = int32(2)
	goto L9
L9:
	;
	v24 = l5 & int64(4294967295)
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = int64(0)
	goto L12
L11:
	;
	v26 = v24
	goto L12
L12:
	;
	if l1 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v27 = v24
	goto L15
L14:
	;
	v27 = v26
	goto L15
L15:
	;
	v29 = l5 << (uint(int64(32)) % 64)
	if l1 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v31 = int64(0)
	goto L18
L17:
	;
	v31 = v29
	goto L18
L18:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v32 = v29
	goto L21
L20:
	;
	v32 = v31
	goto L21
L21:
	;
	v35 = l0
	v44 = int32(0)
	goto L22
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v51
	if base.B2i32(v51 == int32(0))&(l1&l2) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v85 = v76
	goto L1
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v27 | v32
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l6
	v76 = F_aclupdate(m, v35, v15, v22, l7, l3)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L27
	} else {
		goto L32
	}
L27:
	;
	return int32(0)
L28:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(164600), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(497029), int32(211), int32(96784))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_pfree(m, v35)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v81 = v44 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v81 < v82 {
		v35 = v76
		v44 = v81
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
}
func F_metaphone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v802 int32
	_ = v802
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_text_to_cstring(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L6
	} else {
		goto L360
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L6
	} else {
		goto L356
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L6
	} else {
		goto L352
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L6
	} else {
		goto L349
	}
L5:
	;
	v927 = F_cstring_to_text(m, v921)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L6
	} else {
		goto L348
	}
L6:
	;
	return int32(0)
L7:
	;
	v20 = F_strlen(m, v16)
	mBase = m.M
	if v20 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v921 = int32(741336)
	goto L5
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(256)) <= base.Ui32(v20) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if int32(256) <= v26 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v26 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v38 = F_palloc(m, v26+int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.Ui32(v40-int32(97)) < base.Ui32(int32(26)) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	switch v88 - int32(65) {
	case 0:
		goto L41
	default:
		v190 = v87
		v192 = v2
		goto L35
	case 4, 8, 14, 20:
		goto L37
	case 6, 10, 15:
		goto L40
	case 22:
		goto L39
	case 23:
		goto L38
	}
L18:
	;
	if base.Ui32(int32(-27)) < base.Ui32(v47&int32(223)-int32(91)) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v47 = v40 & int32(95)
	goto L21
L20:
	;
	v47 = v40
	goto L21
L21:
	;
	goto L18
L22:
	;
	v87 = int32(0)
	v88 = v47
	v90 = v16
	goto L17
L23:
	;
	goto L24
L24:
	;
	v56 = int32(0)
	v57 = v47
	goto L25
L25:
	;
	if v57 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v87 = v71
	v88 = v80
	v90 = v72
	goto L17
L27:
	;
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v68)
	v921 = v38
	goto L5
L28:
	;
	goto L29
L29:
	;
	v71 = v56 + int32(1)
	v72 = v16 + v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32(v73-int32(97)) < base.Ui32(int32(26)) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if base.Ui32(v80&int32(223)-int32(91)) < base.Ui32(int32(-26)) {
		v56 = v71
		v57 = v80
		goto L25
	} else {
		goto L34
	}
L31:
	;
	v80 = v73 & int32(95)
	goto L33
L32:
	;
	v80 = v73
	goto L33
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	v194 = v190 + v16
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if base.Ui32(v195-int32(97)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L71
	}
L36:
	;
	v187 = int32(1)
	v190 = v87 + v187
	v192 = v187
	goto L35
L37:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v88)
	goto L36
L38:
	;
	v184 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v184)
	goto L36
L39:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if base.Ui32(v134-int32(97)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if base.Ui32(v119-int32(97)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L41:
	;
	v99 = int32(1)
	v101 = v87 + v99
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v101))))
	if base.Ui32(v103-int32(97)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v110 == int32(69) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v110 = v103 & int32(95)
	goto L45
L44:
	;
	v110 = v103
	goto L45
L45:
	;
	goto L42
L46:
	;
	v113 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v113)
	v190 = v87 + int32(2)
	v192 = v99
	goto L35
L47:
	;
	goto L48
L48:
	;
	v117 = int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v117)
	v190 = v101
	v192 = v99
	goto L35
L49:
	;
	if v126 != int32(78) {
		v190 = v87
		v192 = v2
		goto L35
	} else {
		goto L53
	}
L50:
	;
	v126 = v119 & int32(95)
	goto L52
L51:
	;
	v126 = v119
	goto L52
L52:
	;
	goto L49
L53:
	;
	v129 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v129)
	v190 = v87 + int32(2)
	v192 = int32(1)
	goto L35
L54:
	;
	v148 = int32(0)
	v151 = base.I32_extend8_s(v141) & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v151|int32(32)-int32(97)) {
		v174 = v148
		goto L60
	} else {
		goto L61
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v141)
	v190 = v87 + int32(2)
	v192 = int32(1)
	goto L35
L56:
	;
	switch v141 - int32(72) {
	case 0, 10:
		goto L55
	default:
		goto L54
	}
L57:
	;
	v141 = v134 & int32(95)
	goto L59
L58:
	;
	v141 = v134
	goto L59
L59:
	;
	goto L56
L60:
	;
	if v174&int32(1) == int32(0) {
		v190 = v87
		v192 = v2
		goto L35
	} else {
		goto L67
	}
L61:
	;
	if base.Ui32(v151-int32(97)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v167 = base.I32_extend8_s(v164) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v167) {
		v174 = v148
		goto L60
	} else {
		goto L66
	}
L63:
	;
	v164 = v151 & int32(95)
	goto L65
L64:
	;
	v164 = v151
	goto L65
L65:
	;
	goto L62
L66:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v167)+uint32(_consts[1459]))))
	v174 = v172
	goto L60
L67:
	;
	v179 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v179)
	v190 = v87 + int32(2)
	v192 = int32(1)
	goto L35
L68:
	;
	v915 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v906+v38))) = uint8(v915)
	v921 = v38
	goto L5
L69:
	;
	if v202 == int32(0) {
		v906 = v192
		goto L68
	} else {
		goto L73
	}
L70:
	;
	v202 = v195 & int32(95)
	goto L72
L71:
	;
	v202 = v195
	goto L72
L72:
	;
	goto L69
L73:
	;
	if base.Ui32(v26) <= base.Ui32(v192) {
		v906 = v192
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v206 = v190
	v207 = v202
	v208 = v192
	v209 = v194
	goto L75
L75:
	;
	if base.Ui32(int32(25)) < base.Ui32(v207&int32(223)-int32(65)) {
		v885 = v206
		v887 = v208
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v906 = v887
	goto L68
L77:
	;
	v891 = v885 + int32(1)
	v892 = v16 + v891
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	if base.Ui32(v893-int32(97)) < base.Ui32(int32(26)) {
		goto L343
	} else {
		goto L344
	}
L78:
	;
	v223 = base.B2i32(v206 <= int32(0))
	if v206 <= int32(0) {
		goto L84
	} else {
		goto L85
	}
L79:
	;
	v885 = v206 + v883
	v887 = v882
	goto L77
L80:
	;
	v882 = v208 + int32(1)
	v883 = int32(0)
	goto L79
L81:
	;
	v875 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v875)
	goto L80
L82:
	;
	v871 = int32(1)
	v882 = v208 + v871
	v883 = v871
	goto L79
L83:
	;
	v768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v768-int32(97)) < base.Ui32(int32(26)) {
		goto L301
	} else {
		goto L302
	}
L84:
	;
	v237 = int32(0)
	switch v207 - int32(66) {
	case 0:
		goto L107
	case 1:
		goto L83
	case 2:
		goto L106
	default:
		v882 = v208
		v883 = v237
		goto L79
	case 4, 8, 10, 11, 12, 16:
		goto L93
	case 5:
		goto L105
	case 6:
		goto L104
	case 9:
		goto L103
	case 14:
		goto L102
	case 15:
		goto L101
	case 17:
		goto L100
	case 18:
		goto L99
	case 20:
		goto L98
	case 21:
		goto L97
	case 22:
		goto L96
	case 23:
		goto L95
	case 24:
		goto L94
	}
L85:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v226-int32(97)) < base.Ui32(int32(26)) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v233 != v207 {
		goto L84
	} else {
		goto L90
	}
L87:
	;
	v233 = v226 & int32(95)
	goto L89
L88:
	;
	v233 = v226
	goto L89
L89:
	;
	goto L86
L90:
	;
	if v207 != int32(67) {
		v885 = v206
		v887 = v208
		goto L77
	} else {
		goto L91
	}
L91:
	;
	goto L83
L92:
	;
	v882 = v208 + int32(1)
	v883 = v237
	goto L79
L93:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v207)
	goto L92
L94:
	;
	v760 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v760)
	goto L92
L95:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v724-int32(97)) < base.Ui32(int32(26)) {
		goto L289
	} else {
		goto L290
	}
L96:
	;
	v714 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v714)
	v717 = v208 + int32(1)
	if v26 <= v717 {
		goto L285
	} else {
		goto L286
	}
L97:
	;
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v678-int32(97)) < base.Ui32(int32(26)) {
		goto L275
	} else {
		goto L276
	}
L98:
	;
	v676 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v676)
	goto L92
L99:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v646-int32(97)) < base.Ui32(int32(26)) {
		goto L266
	} else {
		goto L267
	}
L100:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v569-int32(97)) < base.Ui32(int32(26)) {
		goto L237
	} else {
		goto L238
	}
L101:
	;
	v567 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v567)
	goto L92
L102:
	;
	v549 = v208 + v38
	v551 = v208 + int32(1)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v552-int32(97)) < base.Ui32(int32(26)) {
		goto L226
	} else {
		goto L227
	}
L103:
	;
	if v223 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L104:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v460-int32(97)) < base.Ui32(int32(26)) {
		goto L194
	} else {
		goto L195
	}
L105:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v306-int32(97)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L106:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32(v257-int32(97)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L107:
	;
	if v223 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v244-int32(97)) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	goto L110
L110:
	;
	v255 = int32(66)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v255)
	goto L92
L111:
	;
	if v251 == int32(77) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L115
	}
L112:
	;
	v251 = v244 & int32(95)
	goto L114
L113:
	;
	v251 = v244
	goto L114
L114:
	;
	goto L111
L115:
	;
	goto L110
L116:
	;
	v304 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v304)
	goto L92
L117:
	;
	if v264 != int32(71) {
		goto L116
	} else {
		goto L121
	}
L118:
	;
	v264 = v257 & int32(95)
	goto L120
L119:
	;
	v264 = v257
	goto L120
L120:
	;
	goto L117
L121:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v267-int32(97)) < base.Ui32(int32(26)) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v276 = v274 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v276|int32(32)-int32(97)) {
		goto L116
	} else {
		goto L126
	}
L123:
	;
	v274 = v267 & int32(95)
	goto L125
L124:
	;
	v274 = v267
	goto L125
L125:
	;
	goto L122
L126:
	;
	if base.Ui32(v276-int32(97)) < base.Ui32(int32(26)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v292 = base.I32_extend8_s(v289) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v292) {
		goto L116
	} else {
		goto L131
	}
L128:
	;
	v289 = v276 & int32(95)
	goto L130
L129:
	;
	v289 = v276
	goto L130
L130:
	;
	goto L127
L131:
	;
	if int32(1)<<(uint(v292)%32)&int32(50331375) != 0 {
		goto L116
	} else {
		goto L132
	}
L132:
	;
	v300 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v300)
	goto L82
L133:
	;
	v416 = v313 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v416|int32(32)-int32(97)) {
		goto L177
	} else {
		goto L178
	}
L134:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v373-int32(97)) < base.Ui32(int32(26)) {
		goto L161
	} else {
		goto L162
	}
L135:
	;
	if v206 < int32(3) {
		goto L140
	} else {
		goto L141
	}
L136:
	;
	switch v313 - int32(72) {
	case 0:
		goto L135
	default:
		goto L133
	case 6:
		goto L134
	}
L137:
	;
	v313 = v306 & int32(95)
	goto L139
L138:
	;
	v313 = v306
	goto L139
L139:
	;
	goto L136
L140:
	;
	v371 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v371)
	goto L82
L141:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(3)))))
	if base.Ui32(v320-int32(97)) < base.Ui32(int32(26)) {
		goto L144
	} else {
		goto L145
	}
L142:
	;
	if v206 == int32(3) {
		goto L140
	} else {
		goto L154
	}
L143:
	;
	v329 = v327 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v329|int32(32)-int32(97)) {
		goto L142
	} else {
		goto L147
	}
L144:
	;
	v327 = v320 & int32(95)
	goto L146
L145:
	;
	v327 = v320
	goto L146
L146:
	;
	goto L143
L147:
	;
	if base.Ui32(v329-int32(97)) < base.Ui32(int32(26)) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v345 = base.I32_extend8_s(v342) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v345) {
		goto L142
	} else {
		goto L152
	}
L149:
	;
	v342 = v329 & int32(95)
	goto L151
L150:
	;
	v342 = v329
	goto L151
L151:
	;
	goto L148
L152:
	;
	if int32(1)<<(uint(v345)%32)&int32(67108725) == int32(0) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L153
	}
L153:
	;
	goto L142
L154:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(4)))))
	if base.Ui32(v359-int32(97)) < base.Ui32(int32(26)) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	if v366 == int32(72) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L159
	}
L156:
	;
	v366 = v359 & int32(95)
	goto L158
L157:
	;
	v366 = v359
	goto L158
L158:
	;
	goto L155
L159:
	;
	goto L140
L160:
	;
	if base.Ui32(int32(25)) < base.Ui32(v380&int32(223)-int32(65)) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L164
	}
L161:
	;
	v380 = v373 & int32(95)
	goto L163
L162:
	;
	v380 = v373
	goto L163
L163:
	;
	goto L160
L164:
	;
	if v380 == int32(69) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v389 == int32(0) {
		v399 = v389
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v413 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v413)
	goto L92
L168:
	;
	v401 = v399 & int32(255)
	if base.Ui32(v401-int32(97)) < base.Ui32(int32(26)) {
		goto L173
	} else {
		goto L174
	}
L169:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v392 == int32(0) {
		v399 = v392
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if v395 == int32(0) {
		v399 = v395
		goto L168
	} else {
		goto L171
	}
L171:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
	v399 = v398
	goto L168
L172:
	;
	if v408 == int32(68) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L176
	}
L173:
	;
	v408 = v401 & int32(95)
	goto L175
L174:
	;
	v408 = v401
	goto L175
L175:
	;
	goto L172
L176:
	;
	goto L167
L177:
	;
	v458 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v458)
	goto L92
L178:
	;
	if base.Ui32(v416-int32(97)) < base.Ui32(int32(26)) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v432 = base.I32_extend8_s(v429) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v432) {
		goto L177
	} else {
		goto L183
	}
L180:
	;
	v429 = v416 & int32(95)
	goto L182
L181:
	;
	v429 = v416
	goto L182
L182:
	;
	goto L179
L183:
	;
	if int32(1)<<(uint(v432)%32)&int32(50331375) != 0 {
		goto L177
	} else {
		goto L184
	}
L184:
	;
	if v223 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v443-int32(97)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L186:
	;
	goto L187
L187:
	;
	v454 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v454)
	goto L92
L188:
	;
	if v450 == int32(71) {
		goto L177
	} else {
		goto L192
	}
L189:
	;
	v450 = v443 & int32(95)
	goto L191
L190:
	;
	v450 = v443
	goto L191
L191:
	;
	goto L188
L192:
	;
	goto L187
L193:
	;
	v469 = v467 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v469|int32(32)-int32(97)) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L197
	}
L194:
	;
	v467 = v460 & int32(95)
	goto L196
L195:
	;
	v467 = v460
	goto L196
L196:
	;
	goto L193
L197:
	;
	if base.Ui32(v469-int32(97)) < base.Ui32(int32(26)) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v485 = base.I32_extend8_s(v482) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v485) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L202
	}
L199:
	;
	v482 = v469 & int32(95)
	goto L201
L200:
	;
	v482 = v469
	goto L201
L201:
	;
	goto L198
L202:
	;
	if int32(1)<<(uint(v485)%32)&int32(66043630) != 0 {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L203
	}
L203:
	;
	if v206 <= int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v530 = int32(72)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v530)
	goto L92
L205:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v494-int32(97)) < base.Ui32(int32(26)) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v503 = v501 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v503|int32(32)-int32(97)) {
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v501 = v494 & int32(95)
	goto L209
L208:
	;
	v501 = v494
	goto L209
L209:
	;
	goto L206
L210:
	;
	if base.Ui32(v503-int32(97)) < base.Ui32(int32(26)) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v519 = base.I32_extend8_s(v516) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v519) {
		goto L204
	} else {
		goto L215
	}
L212:
	;
	v516 = v503 & int32(95)
	goto L214
L213:
	;
	v516 = v503
	goto L214
L214:
	;
	goto L211
L215:
	;
	if int32(1)<<(uint(v519)%32)&int32(66289595) == int32(0) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L216
	}
L216:
	;
	goto L204
L217:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v536-int32(97)) < base.Ui32(int32(26)) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	v547 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v547)
	goto L92
L220:
	;
	if v543 == int32(67) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L224
	}
L221:
	;
	v543 = v536 & int32(95)
	goto L223
L222:
	;
	v543 = v536
	goto L223
L223:
	;
	goto L220
L224:
	;
	goto L219
L225:
	;
	if v559 == int32(72) {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v559 = v552 & int32(95)
	goto L228
L227:
	;
	v559 = v552
	goto L228
L228:
	;
	goto L225
L229:
	;
	v562 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v549))) = uint8(v562)
	v882 = v551
	v883 = v237
	goto L79
L230:
	;
	goto L231
L231:
	;
	v564 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v549))) = uint8(v564)
	v882 = v551
	v883 = v237
	goto L79
L232:
	;
	v644 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v644)
	goto L92
L233:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v595 == int32(0) {
		v602 = v595
		goto L245
	} else {
		goto L246
	}
L234:
	;
	v593 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v593)
	goto L82
L235:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v579-int32(97)) < base.Ui32(int32(26)) {
		goto L242
	} else {
		goto L243
	}
L236:
	;
	switch v576 - int32(67) {
	case 0:
		goto L233
	default:
		goto L232
	case 5:
		goto L234
	case 6:
		goto L235
	}
L237:
	;
	v576 = v569 & int32(95)
	goto L239
L238:
	;
	v576 = v569
	goto L239
L239:
	;
	goto L236
L240:
	;
	v590 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v590)
	goto L92
L241:
	;
	switch v586 - int32(65) {
	case 0, 14:
		goto L240
	default:
		goto L232
	}
L242:
	;
	v586 = v579 & int32(95)
	goto L244
L243:
	;
	v586 = v579
	goto L244
L244:
	;
	goto L241
L245:
	;
	v604 = v602 & int32(255)
	if base.Ui32(v604-int32(97)) < base.Ui32(int32(26)) {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v598 == int32(0) {
		v602 = v598
		goto L245
	} else {
		goto L247
	}
L247:
	;
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	v602 = v601
	goto L245
L248:
	;
	if v611 != int32(72) {
		goto L232
	} else {
		goto L252
	}
L249:
	;
	v611 = v604 & int32(95)
	goto L251
L250:
	;
	v611 = v604
	goto L251
L251:
	;
	goto L248
L252:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v614 == int32(0) {
		v624 = v614
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v626 = v624 & int32(255)
	if base.Ui32(v626-int32(97)) < base.Ui32(int32(26)) {
		goto L258
	} else {
		goto L259
	}
L254:
	;
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if v617 == int32(0) {
		v624 = v617
		goto L253
	} else {
		goto L255
	}
L255:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if v620 == int32(0) {
		v624 = v620
		goto L253
	} else {
		goto L256
	}
L256:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
	v624 = v623
	goto L253
L257:
	;
	if v633 != int32(87) {
		goto L232
	} else {
		goto L261
	}
L258:
	;
	v633 = v626 & int32(95)
	goto L260
L259:
	;
	v633 = v626
	goto L260
L260:
	;
	goto L257
L261:
	;
	v637 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v637)
	v882 = v208 + int32(1)
	v883 = int32(2)
	goto L79
L262:
	;
	v673 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v673)
	goto L92
L263:
	;
	v670 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v670)
	goto L82
L264:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v656-int32(97)) < base.Ui32(int32(26)) {
		goto L271
	} else {
		goto L272
	}
L265:
	;
	switch v653 - int32(72) {
	case 0:
		goto L263
	case 1:
		goto L264
	default:
		goto L262
	}
L266:
	;
	v653 = v646 & int32(95)
	goto L268
L267:
	;
	v653 = v646
	goto L268
L268:
	;
	goto L265
L269:
	;
	v667 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v667)
	goto L92
L270:
	;
	switch v663 - int32(65) {
	case 0, 14:
		goto L269
	default:
		goto L262
	}
L271:
	;
	v663 = v656 & int32(95)
	goto L273
L272:
	;
	v663 = v656
	goto L273
L273:
	;
	goto L270
L274:
	;
	v687 = v685 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v687|int32(32)-int32(97)) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L278
	}
L275:
	;
	v685 = v678 & int32(95)
	goto L277
L276:
	;
	v685 = v678
	goto L277
L277:
	;
	goto L274
L278:
	;
	if base.Ui32(v687-int32(97)) < base.Ui32(int32(26)) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v703 = base.I32_extend8_s(v700) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v703) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L283
	}
L280:
	;
	v700 = v687 & int32(95)
	goto L282
L281:
	;
	v700 = v687
	goto L282
L282:
	;
	goto L279
L283:
	;
	if int32(1)<<(uint(v703)%32)&int32(66043630) != 0 {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L284
	}
L284:
	;
	v711 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v711)
	goto L92
L285:
	;
	v882 = v717
	v883 = v237
	goto L79
L286:
	;
	goto L287
L287:
	;
	v720 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v717+v38))) = uint8(v720)
	v882 = v208 + int32(2)
	v883 = v237
	goto L79
L288:
	;
	v733 = v731 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v733|int32(32)-int32(97)) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L292
	}
L289:
	;
	v731 = v724 & int32(95)
	goto L291
L290:
	;
	v731 = v724
	goto L291
L291:
	;
	goto L288
L292:
	;
	if base.Ui32(v733-int32(97)) < base.Ui32(int32(26)) {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v749 = base.I32_extend8_s(v746) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v749) {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L297
	}
L294:
	;
	v746 = v733 & int32(95)
	goto L296
L295:
	;
	v746 = v733
	goto L296
L296:
	;
	goto L293
L297:
	;
	if int32(1)<<(uint(v749)%32)&int32(66043630) != 0 {
		v882 = v208
		v883 = v237
		goto L79
	} else {
		goto L298
	}
L298:
	;
	v757 = int32(89)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v757)
	goto L92
L299:
	;
	if v775 != int32(72) {
		goto L81
	} else {
		goto L327
	}
L300:
	;
	v777 = v775 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v777|int32(32)-int32(97)) {
		goto L299
	} else {
		goto L304
	}
L301:
	;
	v775 = v768 & int32(95)
	goto L303
L302:
	;
	v775 = v768
	goto L303
L303:
	;
	goto L300
L304:
	;
	if base.Ui32(v777-int32(97)) < base.Ui32(int32(26)) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	v793 = base.I32_extend8_s(v790) - int32(65)
	if base.Ui32(int32(25)) < base.Ui32(v793) {
		goto L299
	} else {
		goto L309
	}
L306:
	;
	v790 = v777 & int32(95)
	goto L308
L307:
	;
	v790 = v777
	goto L308
L308:
	;
	goto L305
L309:
	;
	if int32(1)<<(uint(v793)%32)&int32(50331375) != 0 {
		goto L299
	} else {
		goto L310
	}
L310:
	;
	if v775 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	if v223 == int32(0) {
		goto L319
	} else {
		goto L320
	}
L312:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v802-int32(97)) < base.Ui32(int32(26)) {
		goto L314
	} else {
		goto L315
	}
L313:
	;
	if v775 != int32(73) {
		goto L311
	} else {
		goto L317
	}
L314:
	;
	v809 = v802 & int32(95)
	goto L316
L315:
	;
	v809 = v802
	goto L316
L316:
	;
	goto L313
L317:
	;
	if v809 != int32(65) {
		goto L311
	} else {
		goto L318
	}
L318:
	;
	v815 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v815)
	goto L80
L319:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v823-int32(97)) < base.Ui32(int32(26)) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	goto L321
L321:
	;
	v835 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v835)
	goto L80
L322:
	;
	if v830 == int32(83) {
		v882 = v208
		v883 = int32(0)
		goto L79
	} else {
		goto L326
	}
L323:
	;
	v830 = v823 & int32(95)
	goto L325
L324:
	;
	v830 = v823
	goto L325
L325:
	;
	goto L322
L326:
	;
	goto L321
L327:
	;
	v840 = int32(75)
	v841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
	if base.Ui32(v841-int32(97)) < base.Ui32(int32(26)) {
		goto L330
	} else {
		goto L331
	}
L328:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v38))) = uint8(v866)
	goto L82
L329:
	;
	if v848 == int32(82) {
		v866 = v840
		goto L328
	} else {
		goto L333
	}
L330:
	;
	v848 = v841 & int32(95)
	goto L332
L331:
	;
	v848 = v841
	goto L332
L332:
	;
	goto L329
L333:
	;
	if v223 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209-int32(1)))))
	if base.Ui32(v855-int32(97)) < base.Ui32(int32(26)) {
		goto L338
	} else {
		goto L339
	}
L335:
	;
	goto L336
L336:
	;
	v866 = int32(88)
	goto L328
L337:
	;
	if v862 == int32(83) {
		v866 = v840
		goto L328
	} else {
		goto L341
	}
L338:
	;
	v862 = v855 & int32(95)
	goto L340
L339:
	;
	v862 = v855
	goto L340
L340:
	;
	goto L337
L341:
	;
	goto L336
L342:
	;
	if v900 == int32(0) {
		v906 = v887
		goto L68
	} else {
		goto L346
	}
L343:
	;
	v900 = v893 & int32(95)
	goto L345
L344:
	;
	v900 = v893
	goto L345
L345:
	;
	goto L342
L346:
	;
	if v887 < v26 {
		v206 = v891
		v207 = v900
		v208 = v887
		v209 = v892
		goto L75
	} else {
		goto L347
	}
L347:
	;
	goto L76
L348:
	;
	m.G0 = v13 + int32(32)
	return v927
L349:
	;
	F_errmsg_internal(m, int32(554710), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L6
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(497330), int32(368), int32(372176))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L6
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L6
	} else {
		goto L353
	}
L353:
	;
	F_errmsg(m, int32(329683), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L6
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(497330), int32(287), int32(372177))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L6
	} else {
		goto L355
	}
L355:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L356:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L6
	} else {
		goto L357
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(255)
	F_errmsg(m, int32(159195), v13+int32(16))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L6
	} else {
		goto L358
	}
L358:
	;
	F_errfinish(m, int32(497330), int32(282), int32(372177))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L6
	} else {
		goto L359
	}
L359:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L360:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L6
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(255)
	F_errmsg(m, int32(159241), v13)
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L6
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(497330), int32(275), int32(372177))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L6
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_missing_hash(m *base.Module, l0 int32, l1 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
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
func F_mix_decrypt_resync(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != int32(2) {
		v64 = l1
		v66 = l3
		v67 = v10
		v72 = l2
		if v72 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
			return v72
		} else {
			v83 = v64
			v85 = v66
			v86 = v67
			for {
				v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
				*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(84))))) = uint8(v92)
				v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(52))))))
				v96 = v92 ^ v95
				*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v96)
				v98 = int32(1)
				v103 = v86 + v98
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v105 = v104 + v72
				if v103 < v105 {
					v83 = v83 + v98
					v85 = v85 + v98
					v86 = v103
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
			return v72
		}
	} else {
		v15 = int32(2) - v10
		if l2 < v15 {
			v17 = l2
		} else {
			v17 = v15
		}
		if v17 <= int32(0) {
			v51 = l1
			v53 = l3
			v54 = v10 + v17
		} else {
			v26 = l1
			v28 = l3
			v30 = v10
			for {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
				*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(84))))) = uint8(v35)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(52))))))
				v39 = v35 ^ v38
				*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v39)
				v41 = int32(1)
				v42 = v28 + v41
				v44 = v26 + v41
				v46 = v30 + v41
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v48 = v47 + v17
				if v46 < v48 {
					v26 = v44
					v28 = v42
					v30 = v46
					continue
				} else {
					break
				}
				break
			}
			v51 = v44
			v53 = v42
			v54 = v48
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54
		if v54 == int32(2) {
			v110 = l0 + int32(20)
			v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v115 = v113 - int32(2)
			if v115 != 0 {
				v116 = F__emscripten_memcpy_bulkmem(m, v110, l0+int32(86), v115)
				mBase = m.M
				v117 = v116
			} else {
				v117 = v110
			}
			v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+84)))
			*(*uint16)(unsafe.Add(mBase, uint32(v117+v115))) = uint16(v119)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			return v17
		} else {
			v64 = v51
			v66 = v53
			v67 = v54
			v72 = l2 - v17
			if v72 <= int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
				return v72
			} else {
				v83 = v64
				v85 = v66
				v86 = v67
				for {
					v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
					*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(84))))) = uint8(v92)
					v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+(l0+int32(52))))))
					v96 = v92 ^ v95
					*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v96)
					v98 = int32(1)
					v103 = v86 + v98
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v105 = v104 + v72
					if v103 < v105 {
						v83 = v83 + v98
						v85 = v85 + v98
						v86 = v103
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
				return v72
			}
		}
	}
}
func F_mkSPNode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
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
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
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
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v449 int32
	_ = v449
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	if l2 <= l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(32)
	return v449
L2:
	;
	v449 = v5
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = l2 - l1
	v25 = int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if l1+v25 == l2 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v24&v25 == int32(0) {
		v116 = v89
		goto L18
	} else {
		goto L19
	}
L6:
	;
	v88 = v5
	v89 = v5
	v91 = l1
	goto L5
L7:
	;
	goto L8
L8:
	;
	v39 = v5
	v40 = v5
	v42 = l1
	v44 = v5
	goto L9
L9:
	;
	v54 = v42 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v27+v54)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if l3 < v57 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v88 = v77
	v89 = v78
	v91 = v80
	goto L5
L11:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v56)+8)))
	v65 = v62
	v66 = v40 + base.B2i32(v39&int32(255) != v62)
	goto L13
L12:
	;
	v65 = v39
	v66 = v40
	goto L13
L13:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(4)+v54)))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if l3 < v69 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v68)+8)))
	v77 = v74
	v78 = v66 + base.B2i32(v65&int32(255) != v74)
	goto L16
L15:
	;
	v77 = v65
	v78 = v66
	goto L16
L16:
	;
	v79 = int32(2)
	v80 = v42 + v79
	v82 = v44 + v79
	if v82 != v24&int32(-2) {
		v39 = v77
		v40 = v78
		v42 = v80
		v44 = v82
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L10
L18:
	;
	if v116 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v27+v91<<(uint(int32(2))%32))))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v108 <= l3 {
		v116 = v89
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v107)+8)))
	v116 = v89 + base.B2i32(v111 != v88&int32(255))
	goto L18
L21:
	;
	v449 = int32(0)
	goto L1
L22:
	;
	goto L23
L23:
	;
	v122 = v116 << (uint(int32(3)) % 32)
	v124 = v122 | int32(4)
	if base.Ui32(int32(1025)) <= base.Ui32(v124) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v116
	v153 = l3 + int32(1)
	v158 = l1
	v162 = v150 + int32(4)
	v164 = l1
	v172 = int32(0)
	goto L35
L25:
	;
	v127 = F_palloc0(m, v124)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v134 = (v122 + int32(11)) & int32(4088)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v134) <= base.Ui32(v135) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	return int32(0)
L29:
	;
	v150 = v127
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v142 - v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v134 + v143
	v150 = v143
	goto L24
L31:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v142 = v135
	v143 = v137
	goto L30
L32:
	;
	goto L33
L33:
	;
	v138 = int32(8192)
	v140 = F_palloc0(m, v138)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v142 = v138
	v143 = v140
	goto L30
L35:
	;
	v176 = v164 << (uint(int32(2)) % 32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176+v177)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v180 <= l3 {
		v422 = v158
		v424 = v162
		v430 = v172
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v436 = F_mkSPNode(m, l0, v422, l2, v153)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L28
	} else {
		goto L92
	}
L37:
	;
	v434 = v164 + int32(1)
	if v434 != l2 {
		v158 = v422
		v162 = v424
		v164 = v434
		v172 = v430
		goto L35
	} else {
		goto L91
	}
L38:
	;
	v183 = v172 & int32(255)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v179)+8)))
	if v183 == v185 {
		v199 = v158
		v200 = v179
		v201 = v162
		v202 = v172
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v200)+8)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v208 = v204 | v205&int32(-256)
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v210+v176)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v213 != v153 {
		v422 = v199
		v424 = v201
		v430 = v202
		goto L37
	} else {
		goto L45
	}
L40:
	;
	if v183 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = v158
	v200 = v179
	v201 = v162
	v202 = v185
	goto L39
L42:
	;
	goto L43
L43:
	;
	v189 = F_mkSPNode(m, l0, v158, v164, v153)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v189
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v194+v176)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+l3)+8)))
	v199 = v164
	v200 = v196
	v201 = v162 + int32(8)
	v202 = v198
	goto L39
L45:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v205&int32(256) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v387 | int32(256)
	v401 = F_makeCompoundFlags(m, l0, int32(base.Ui32(v387)>>(uint(int32(13))%32)))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L28
	} else {
		goto L86
	}
L47:
	;
	v387 = v208&int32(8191) | v215<<(uint(int32(13))%32)
	v395 = int32(0)
	goto L46
L48:
	;
	if int32(base.Ui32(v205)>>(uint(int32(13))%32)) == v215 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v223 = F_makeCompoundFlags(m, l0, v215)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v228+v176)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v235 = int32(base.Ui32(v233) >> (uint(int32(13)) % 32))
	v237 = v235 << (uint(int32(2)) % 32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v232+v237)))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v240 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v387 = v368&int32(8191) | v369<<(uint(int32(13))%32)
	v395 = v223&int32(base.Ui32(v205)>>(uint(int32(9))%32)) ^ int32(1)
	goto L46
L52:
	;
	v368 = v233
	v369 = v231
	goto L51
L53:
	;
	goto L54
L54:
	;
	v244 = v231 << (uint(int32(2)) % 32)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v232+v244)))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246))))
	if v247 == int32(0) {
		v368 = v233
		v369 = v235
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v250 <= v251+int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v250 << (uint(int32(1)) % 32)
	v260 = F_repalloc(m, v232, v250<<(uint(int32(3))%32))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L28
	} else {
		goto L59
	}
L57:
	;
	v268 = v232
	v269 = v239
	v270 = v246
	v271 = v251
	goto L58
L58:
	;
	v272 = int32(2)
	v274 = v271<<(uint(v272)%32) + v268
	v275 = F_strlen(m, v269)
	mBase = m.M
	v276 = F_strlen(m, v270)
	mBase = m.M
	v277 = v275 + v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v278 == v272 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v260
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v244)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v260+v237)))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v268 = v260
	v269 = v266
	v270 = v264
	v271 = v267
	goto L58
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v274)+4)) = int32(0)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v362 + int32(1)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v368 = v366
	v369 = v362
	goto L51
L61:
	;
	v282 = v277 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v282) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	goto L63
L63:
	;
	v318 = v277 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v318) {
		goto L76
	} else {
		goto L77
	}
L64:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v307+v237)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v307+v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v309
	v315 = F_pg_sprintf(m, v304, int32(177051), v21)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L28
	} else {
		goto L74
	}
L65:
	;
	v285 = F_palloc0(m, v282)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L28
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v290 = (v277 + int32(9)) & int32(4088)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v290) <= base.Ui32(v291) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v304 = v285
	goto L64
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v298 - v290
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v299 + v290
	v304 = v299
	goto L64
L70:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v298 = v291
	v299 = v293
	goto L69
L71:
	;
	goto L72
L72:
	;
	v294 = int32(8192)
	v296 = F_palloc0(m, v294)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L28
	} else {
		goto L73
	}
L73:
	;
	v298 = v294
	v299 = v296
	goto L69
L74:
	;
	v355 = v304
	goto L60
L75:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v343+v237)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v345
	v353 = F_pg_sprintf(m, v340, int32(175588), v21+int32(16))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L28
	} else {
		goto L85
	}
L76:
	;
	v321 = F_palloc0(m, v318)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L28
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v326 = (v277 + int32(8)) & int32(4088)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v326) <= base.Ui32(v327) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v340 = v321
	goto L75
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v334 - v326
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v335 + v326
	v340 = v335
	goto L75
L81:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v334 = v327
	v335 = v329
	goto L80
L82:
	;
	goto L83
L83:
	;
	v330 = int32(8192)
	v332 = F_palloc0(m, v330)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L28
	} else {
		goto L84
	}
L84:
	;
	v334 = v330
	v335 = v332
	goto L80
L85:
	;
	v355 = v340
	goto L60
L86:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v401 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v413 = v403 | int32(7680)
	goto L89
L88:
	;
	v413 = v403&int32(-7681) | v401<<(uint(int32(9))%32)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v413
	if v395&int32(1) == int32(0) {
		v422 = v199
		v424 = v201
		v430 = v202
		goto L37
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v413 & int32(-513)
	v422 = v199
	v424 = v201
	v430 = v202
	goto L37
L91:
	;
	goto L36
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v436
	v449 = v150
	goto L1
}
func F_moveArrayTypeName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v7 = F_get_typisdefined(m, l0)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v30 = int32(1)
			return v30
		} else {
			v13 = int32(0)
			v14 = F_get_element_type(m, l0)
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v30 = v13
					return v30
				} else {
					v18 = F_get_array_type(m, v14)
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 != l0 {
							v30 = v13
							return v30
						} else {
							v21 = F_makeArrayTypeName(m, l1, l2)
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								F_RenameTypeInternal(m, l0, v21, l2)
								v24 = m.ExcPending
								if v24 != 0 {
									return int32(0)
								} else {
									F_CommandCounterIncrement(m)
									v26 = m.ExcPending
									if v26 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v21)
										v28 = m.ExcPending
										if v28 != 0 {
											return int32(0)
										} else {
											v30 = int32(1)
											return v30
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
func F_moveins(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int64
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int64
	_ = v301
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v8 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L5
L3:
	;
	goto L4
L4:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(4) <= v101 {
		goto L38
	} else {
		goto L39
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	F_createarc(m, l0, v21, v22, v23, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v33 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L5
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	if v67 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v38 = v36 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v38) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if int32(1)<<(uint(v38)%32)&int32(163841) == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v47 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v33*int32(24))+12)) = v56
	v60 = v56
	goto L16
L18:
	;
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v58
	v60 = v58
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = v48
	goto L22
L21:
	;
	goto L22
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
	goto L11
L23:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v66
	goto L23
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v66
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v67
	goto L29
L28:
	;
	goto L29
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v73 - int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v78 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v84 = v18 + int32(8)
	if v77 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v77
	goto L30
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v77
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v78
	goto L36
L35:
	;
	goto L36
L36:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v86 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v84))) = v92
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	goto L10
L37:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v197 != 0 {
		goto L74
	} else {
		goto L75
	}
L38:
	;
	if int32(32) < v8 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L43
L41:
	;
	if base.Ui32(int32(32)) < base.Ui32(v101) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v115 == int32(0) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	F_cparc(m, l0, v115, v118, l2)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115)+4)))
	if v128 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L43
L48:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v115)+20))
	if v162 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v133 = v131 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v133) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	if int32(1)<<(uint(v133)%32)&int32(163841) == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v142 != 0 {
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v115)+36))
	if v143 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v155 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+20))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v147+v128*int32(24))+12)) = v151
	v155 = v151
	goto L53
L55:
	;
	goto L56
L56:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v115)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+32)) = v153
	v155 = v153
	goto L53
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v143
	goto L59
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v115)+32)) = int64(0)
	goto L48
L60:
	;
	if v161 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+20)) = v161
	goto L60
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v161
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+20)) = v162
	goto L66
L65:
	;
	goto L66
L66:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v168 - int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+24))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v115)+28))
	if v173 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v179 = v115 + int32(8)
	if v172 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v172
	goto L67
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = v172
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+28)) = v173
	goto L73
L72:
	;
	goto L73
L73:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = v181 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(0)
	v187 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v179)+16)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v179)+8)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v179))) = v187
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+16)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v115
	goto L47
L74:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_sortins(m, l0, l1)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	F_sortins(m, l0, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	if v205 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v206 == int32(0) {
		v349 = v206
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v349 == int32(0) {
		goto L1
	} else {
		goto L134
	}
L82:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v209 == int32(0) {
		v349 = v206
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v215 = v206
	v216 = v209
	goto L84
L84:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v220 < v222 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v349 = v340
	goto L81
L86:
	;
	if v340 == int32(0) {
		v349 = v340
		goto L81
	} else {
		goto L132
	}
L87:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	v340 = v215
	v341 = v338
	goto L86
L88:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	if v314 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L89:
	;
	if v222 < v220 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+4)))
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216)+4)))
	if v225 < v226 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	if v226 < v225 {
		goto L87
	} else {
		goto L92
	}
L92:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v229 < v230 {
		goto L88
	} else {
		goto L93
	}
L93:
	;
	if v230 < v229 {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v216)+24))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v215)+8))
	v242 = int32(*(*int16)(unsafe.Add(mBase, uint32(v215)+4)))
	if v242 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v340 = v234
	v341 = v233
	goto L86
L96:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	if v276 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L97:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v247 = v245 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v247) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	if int32(1)<<(uint(v247)%32)&int32(163841) == int32(0) {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v256 != 0 {
		goto L96
	} else {
		goto L100
	}
L100:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v215)+36))
	if v257 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v269 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261+v242*int32(24))+12)) = v265
	v269 = v265
	goto L101
L103:
	;
	goto L104
L104:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v257)+32)) = v267
	v269 = v267
	goto L101
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269)+36)) = v257
	goto L107
L106:
	;
	goto L107
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v215)+32)) = int64(0)
	goto L96
L108:
	;
	if v275 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = v275
	goto L108
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+16)) = v275
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v275)+20)) = v276
	goto L114
L113:
	;
	goto L114
L114:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v241)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = v282 - int32(1)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v215)+28))
	if v287 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v293 = v215 + int32(8)
	if v286 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = v286
	goto L115
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287)+24)) = v286
	goto L115
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+28)) = v287
	goto L121
L120:
	;
	goto L121
L121:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v240)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = v295 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = int32(0)
	v301 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v293)+16)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v293)+8)) = v301
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v301
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v215
	goto L95
L122:
	;
	if v313 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+16)) = v313
	goto L122
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+24)) = v313
	goto L122
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+28)) = v314
	goto L128
L127:
	;
	goto L128
L128:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+8)) = v320 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = l2
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v215)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+24)) = v325
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v329 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+28)) = v215
	goto L131
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v215
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v332 + int32(1)
	v340 = v313
	v341 = v216
	goto L86
L132:
	;
	if v341 != 0 {
		v215 = v340
		v216 = v341
		goto L84
	} else {
		goto L133
	}
L133:
	;
	goto L85
L134:
	;
	v358 = v349
	goto L135
L135:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v358)+24))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v358)+28))
	if v364 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L1
L137:
	;
	if v363 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+16)) = v363
	goto L137
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v364)+24)) = v363
	goto L137
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363)+28)) = v364
	goto L143
L142:
	;
	goto L143
L143:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v362)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+8)) = v370 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+12)) = l2
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v358)+24)) = v375
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v379 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+28)) = v358
	goto L146
L145:
	;
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v358
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v382 + int32(1)
	if v363 != 0 {
		v358 = v363
		goto L135
	} else {
		goto L147
	}
L147:
	;
	goto L136
}
func F_mul_d_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(1475), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_multixactmemberssyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_SlruSyncFileTag(m, int32(4384228), l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_my_log2(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	v3 = int32(1073741823)
	if v3 <= l0 {
		v6 = v3
	} else {
		v6 = l0
	}
	if base.Ui32(int32(2)) <= base.Ui32(v6) {
		v14 = int32(32) - base.I32_clz(v6-int32(1))
	} else {
		v14 = int32(0)
	}
	return v14
}
