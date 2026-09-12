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
	v24 = int32(_a_F_MJEvalInnerValues_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_MJEvalInnerValues[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MJEvalInnerValues[0])) = v27
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
	*(*int32)(unsafe.Add(mBase, _c_F_MJEvalInnerValues[0])) = v25
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
	v27 = int32(_a_F_MemoizeHash_equal_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v31
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
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v28
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
	v101 = int32(_a_F_MemoizeHash_equal_0)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v104
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
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v102
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
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
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
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	v2 = int32(0)
	v10 = int32(_a_F_MemoizeHash_hash_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+120))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0])) = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+197)))
	if v19 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0])) = v11
	v1274 = int32(16)
	v1278 = (int32(base.Ui32(v1264)>>(uint(v1274)%32)) ^ v1264) * int32(-2048144789)
	v1283 = (int32(base.Ui32(v1278)>>(uint(int32(13))%32)) ^ v1278) * int32(-1028477387)
	return int32(base.Ui32(v1283)>>(uint(v1274)%32)) ^ v1283
L2:
	;
	if v13 <= int32(0) {
		v1264 = v2
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
		v1264 = v2
		goto L1
	} else {
		goto L210
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
	v1264 = v1216
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
	v1216 = v35
	goto L10
L10:
	;
	v1222 = v25 + int32(1)
	if v1222 != v13 {
		v25 = v1222
		v26 = v1216
		goto L6
	} else {
		goto L209
	}
L11:
	;
	m.G0 = v54 + int32(16)
	v1216 = v35 ^ v1210
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
	v1210 = v319 ^ v311 - base.I32_rotl(v319, int32(24))
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
	v1210 = v585 ^ v577 - base.I32_rotl(v585, int32(24))
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
	if v45&int32(3) == int32(0) {
		v910 = v45
		goto L154
	} else {
		goto L155
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
		v1210 = v870
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
	v1210 = v870
	goto L11
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v51
	F_errmsg_internal(m, int32(_a_F_MemoizeHash_hash_1), v54)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L101
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_MemoizeHash_hash_2), int32(372), int32(_a_F_MemoizeHash_hash_3))
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
	v945 = v943 + int32(1)
	v951 = v945 - int32(1636608432)
	if v45&int32(3) != 0 {
		goto L173
	} else {
		goto L174
	}
L153:
	;
	v943 = v935 - v45
	goto L152
L154:
	;
	v914 = v910
	goto L163
L155:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v894 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v943 = int32(0)
	goto L152
L157:
	;
	goto L158
L158:
	;
	v899 = v45
	goto L159
L159:
	;
	v903 = v899 + int32(1)
	if v903&int32(3) == int32(0) {
		v910 = v903
		goto L154
	} else {
		goto L161
	}
L160:
	;
	v935 = v903
	goto L153
L161:
	;
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v903))))
	if v908 != 0 {
		v899 = v903
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v914)))
	v923 = int32(-2139062144)
	if (int32(16843008)-v920|v920)&v923 == v923 {
		v914 = v914 + int32(4)
		goto L163
	} else {
		goto L165
	}
L164:
	;
	v929 = v914
	goto L166
L165:
	;
	goto L164
L166:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929))))
	if v933 != 0 {
		v929 = v929 + int32(1)
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v935 = v929
	goto L153
L168:
	;
	goto L167
L169:
	;
	v1210 = v1205 ^ v1197 - base.I32_rotl(v1205, int32(24))
	goto L11
L170:
	;
	v1183 = int32(14)
	v1185 = v1179 ^ v1180 - base.I32_rotl(v1179, v1183)
	v1189 = v1185 ^ v1178 - base.I32_rotl(v1185, int32(11))
	v1193 = v1189 ^ v1179 - base.I32_rotl(v1189, int32(25))
	v1197 = v1193 ^ v1185 - base.I32_rotl(v1193, int32(16))
	v1201 = v1197 ^ v1189 - base.I32_rotl(v1197, int32(4))
	v1205 = v1201 ^ v1193 - base.I32_rotl(v1201, v1183)
	goto L169
L171:
	;
	switch v1109 - int32(1) {
	case 0:
		v1171 = v1110
		v1172 = v1111
		v1173 = v1112
		goto L198
	case 1:
		v1164 = v1110
		v1165 = v1111
		v1166 = v1112
		goto L199
	case 2:
		v1157 = v1110
		v1158 = v1111
		v1159 = v1112
		goto L200
	case 3:
		v1151 = v1111
		v1152 = v1112
		goto L201
	case 4:
		v1147 = v1111
		v1148 = v1112
		goto L202
	case 5:
		v1141 = v1111
		v1142 = v1112
		goto L203
	case 6:
		v1135 = v1111
		v1136 = v1112
		goto L204
	case 7:
		v1130 = v1112
		goto L205
	case 8:
		v1125 = v1112
		goto L206
	case 9:
		v1120 = v1112
		goto L207
	case 10:
		goto L208
	default:
		v1178 = v1110
		v1179 = v1111
		v1180 = v1112
		goto L170
	}
L172:
	;
	v1060 = v45
	v1061 = v945
	v1062 = v951
	v1063 = v951
	v1064 = v951
	goto L195
L173:
	;
	if base.Ui32(int32(11)) < base.Ui32(v945) {
		goto L172
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if base.Ui32(v945) < base.Ui32(int32(12)) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v1108 = v45
	v1109 = v945
	v1110 = v951
	v1111 = v951
	v1112 = v951
	goto L171
L177:
	;
	switch v1007 - int32(1) {
	case 0:
		v1057 = v1008
		goto L184
	case 1:
		v1052 = v1008
		goto L185
	case 2:
		goto L186
	case 3:
		v1045 = v1009
		goto L187
	case 4:
		v1042 = v1009
		goto L188
	case 5:
		v1037 = v1009
		goto L189
	case 6:
		goto L190
	case 7:
		v1028 = v1010
		goto L191
	case 8:
		v1023 = v1010
		goto L192
	case 9:
		v1018 = v1010
		goto L193
	case 10:
		goto L194
	default:
		v1178 = v1008
		v1179 = v1009
		v1180 = v1010
		goto L170
	}
L178:
	;
	v1006 = v45
	v1007 = v945
	v1008 = v951
	v1009 = v951
	v1010 = v951
	goto L177
L179:
	;
	goto L180
L180:
	;
	v958 = v45
	v959 = v945
	v960 = v951
	v961 = v951
	v962 = v951
	goto L181
L181:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	v965 = v964 + v961
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v958)+8))
	v969 = v968 + v962
	v971 = int32(4)
	v973 = v966 + v960 - v969 ^ base.I32_rotl(v969, v971)
	v977 = v965 - v973 ^ base.I32_rotl(v973, int32(6))
	v978 = v969 + v965
	v979 = v973 + v978
	v980 = v977 + v979
	v984 = v978 - v977 ^ base.I32_rotl(v977, int32(8))
	v988 = v979 - v984 ^ base.I32_rotl(v984, int32(16))
	v992 = v980 - v988 ^ base.I32_rotl(v988, int32(19))
	v993 = v984 + v980
	v994 = v988 + v993
	v995 = v992 + v994
	v999 = v993 - v992 ^ base.I32_rotl(v992, v971)
	v1000 = int32(12)
	v1001 = v958 + v1000
	v1003 = v959 - v1000
	if base.Ui32(int32(11)) < base.Ui32(v1003) {
		v958 = v1001
		v959 = v1003
		v960 = v994
		v961 = v995
		v962 = v999
		goto L181
	} else {
		goto L183
	}
L182:
	;
	v1006 = v1001
	v1007 = v1003
	v1008 = v994
	v1009 = v995
	v1010 = v999
	goto L177
L183:
	;
	goto L182
L184:
	;
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	v1178 = v1057 + v1058
	v1179 = v1009
	v1180 = v1010
	goto L170
L185:
	;
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+1)))
	v1057 = v1053<<(uint(int32(8))%32) + v1052
	goto L184
L186:
	;
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+2)))
	v1052 = v1048<<(uint(int32(16))%32) + v1008
	goto L185
L187:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	v1178 = v1046 + v1008
	v1179 = v1045
	v1180 = v1010
	goto L170
L188:
	;
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+4)))
	v1045 = v1042 + v1043
	goto L187
L189:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+5)))
	v1042 = v1038<<(uint(int32(8))%32) + v1037
	goto L188
L190:
	;
	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+6)))
	v1037 = v1033<<(uint(int32(16))%32) + v1009
	goto L189
L191:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1006)))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1006)+4))
	v1178 = v1029 + v1008
	v1179 = v1031 + v1009
	v1180 = v1028
	goto L170
L192:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+8)))
	v1028 = v1024<<(uint(int32(8))%32) + v1023
	goto L191
L193:
	;
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+9)))
	v1023 = v1019<<(uint(int32(16))%32) + v1018
	goto L192
L194:
	;
	v1014 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+10)))
	v1018 = v1014<<(uint(int32(24))%32) + v1010
	goto L193
L195:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+4))
	v1067 = v1066 + v1063
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1060)))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1060)+8))
	v1071 = v1070 + v1064
	v1073 = int32(4)
	v1075 = v1068 + v1062 - v1071 ^ base.I32_rotl(v1071, v1073)
	v1079 = v1067 - v1075 ^ base.I32_rotl(v1075, int32(6))
	v1080 = v1071 + v1067
	v1081 = v1075 + v1080
	v1082 = v1079 + v1081
	v1086 = v1080 - v1079 ^ base.I32_rotl(v1079, int32(8))
	v1090 = v1081 - v1086 ^ base.I32_rotl(v1086, int32(16))
	v1094 = v1082 - v1090 ^ base.I32_rotl(v1090, int32(19))
	v1095 = v1086 + v1082
	v1096 = v1090 + v1095
	v1097 = v1094 + v1096
	v1101 = v1095 - v1094 ^ base.I32_rotl(v1094, v1073)
	v1102 = int32(12)
	v1103 = v1060 + v1102
	v1105 = v1061 - v1102
	if base.Ui32(int32(11)) < base.Ui32(v1105) {
		v1060 = v1103
		v1061 = v1105
		v1062 = v1096
		v1063 = v1097
		v1064 = v1101
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v1108 = v1103
	v1109 = v1105
	v1110 = v1096
	v1111 = v1097
	v1112 = v1101
	goto L171
L197:
	;
	goto L196
L198:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	v1178 = v1171 + v1174
	v1179 = v1172
	v1180 = v1173
	goto L170
L199:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
	v1171 = v1167<<(uint(int32(8))%32) + v1164
	v1172 = v1165
	v1173 = v1166
	goto L198
L200:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+2)))
	v1164 = v1160<<(uint(int32(16))%32) + v1157
	v1165 = v1158
	v1166 = v1159
	goto L199
L201:
	;
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+3)))
	v1157 = v1153<<(uint(int32(24))%32) + v1110
	v1158 = v1151
	v1159 = v1152
	goto L200
L202:
	;
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+4)))
	v1151 = v1147 + v1149
	v1152 = v1148
	goto L201
L203:
	;
	v1143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+5)))
	v1147 = v1143<<(uint(int32(8))%32) + v1141
	v1148 = v1142
	goto L202
L204:
	;
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+6)))
	v1141 = v1137<<(uint(int32(16))%32) + v1135
	v1142 = v1136
	goto L203
L205:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+7)))
	v1135 = v1131<<(uint(int32(24))%32) + v1111
	v1136 = v1130
	goto L204
L206:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+8)))
	v1130 = v1126<<(uint(int32(8))%32) + v1125
	goto L205
L207:
	;
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+9)))
	v1125 = v1121<<(uint(int32(16))%32) + v1120
	goto L206
L208:
	;
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+10)))
	v1120 = v1116<<(uint(int32(24))%32) + v1112
	goto L207
L209:
	;
	goto L7
L210:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v12)+152))
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
	v1229 = int32(0)
	v1230 = v2
	goto L211
L211:
	;
	v1239 = base.I32_rotl(v1230, int32(1))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v1242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1240+v1229))))
	if v1242 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1264 = v1258
	goto L1
L213:
	;
	v1249 = v1229 << (uint(int32(2)) % 32)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1226+v1249)))
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1252+v1249)))
	v1255 = F_FunctionCall1Coll(m, v1227+v1229*int32(28), v1251, v1254)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L101
	} else {
		goto L216
	}
L214:
	;
	v1258 = v1239
	goto L215
L215:
	;
	v1261 = v1229 + int32(1)
	if v1261 != v13 {
		v1229 = v1261
		v1230 = v1258
		goto L211
	} else {
		goto L217
	}
L216:
	;
	v1258 = v1255 ^ v1239
	goto L215
L217:
	;
	goto L212
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
				F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(680), int32(_a_F_ModifyWaitEvent_2))
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
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_ModifyWaitEvent[0]))
							if v44 != v46 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_3), int32(0))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(704), int32(_a_F_ModifyWaitEvent_2))
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
								v85 = v76 | int32(_a_F_ModifyWaitEvent_4)
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
					F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_5), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(696), int32(_a_F_ModifyWaitEvent_2))
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
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_ModifyWaitEvent[0]))
						if v44 != v46 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_3), int32(0))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(704), int32(_a_F_ModifyWaitEvent_2))
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
							v85 = v76 | int32(_a_F_ModifyWaitEvent_4)
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
		*(*int32)(unsafe.Add(mBase, _c_F___mmap[0])) = int32(48)
		return int32(-1)
	} else {
		if l1&int32(16) != 0 {
			v18 = int32(-63)
		} else {
			v18 = int32(-48)
		}
		if l1&int32(32) != 0 {
			v21 = int32(_a_F___mmap_0)
			v25 = (l0 + int32(15)) & int32(-16)
			v27 = v25 + int32(40)
			if base.Ui32(int32(-65600)) <= base.Ui32(v27) {
				*(*int32)(unsafe.Add(mBase, _c_F___mmap[0])) = int32(48)
				v191 = int32(0)
			} else {
				if base.Ui32(v27) < base.Ui32(int32(11)) {
					v78 = int32(16)
				} else {
					v78 = (v25 + int32(51)) & int32(-8)
				}
				v82 = F_emscripten_builtin_malloc(m, v78+int32(_a_F___mmap_1))
				mBase = m.M
				if v82 == int32(0) {
					v191 = int32(0)
				} else {
					v86 = v82 - int32(8)
					if int32(_a_F___mmap_2)&v82 == int32(0) {
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
				v228 = int32(_a_F___mmap_3)
				v229 = *(*int32)(unsafe.Add(mBase, _c_F___mmap[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = v229
				*(*int32)(unsafe.Add(mBase, _c_F___mmap[1])) = v220
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
				v228 = int32(_a_F___mmap_3)
				v229 = *(*int32)(unsafe.Add(mBase, _c_F___mmap[1]))
				*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = v229
				*(*int32)(unsafe.Add(mBase, _c_F___mmap[1])) = v220
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
			*(*int32)(unsafe.Add(mBase, _c_F___mmap[0])) = int32(0) - v244
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
						F_errmsg(m, int32(_a_F_macaddr8tomacaddr_0), int32(0))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_macaddr8tomacaddr_1), int32(0))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_macaddr8tomacaddr_2), int32(559), int32(_a_F_macaddr8tomacaddr_3))
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
					F_errmsg(m, int32(_a_F_macaddr8tomacaddr_0), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_macaddr8tomacaddr_1), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_macaddr8tomacaddr_2), int32(559), int32(_a_F_macaddr8tomacaddr_3))
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
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_make_distinct_op_0)
						F_errmsg(m, int32(_a_F_make_distinct_op_1), v9)
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
								F_errfinish(m, int32(_a_F_make_distinct_op_2), int32(3103), int32(_a_F_make_distinct_op_3))
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
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_make_distinct_op_0)
					F_errmsg(m, int32(_a_F_make_distinct_op_4), v9+int32(16))
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
							F_errfinish(m, int32(_a_F_make_distinct_op_2), int32(3097), int32(_a_F_make_distinct_op_3))
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
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v387 int32
	_ = v387
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(17) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	if int32(0) < v233 {
		goto L95
	} else {
		goto L96
	}
L2:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[0]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v225 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L3:
	;
	v218 = v215
	v219 = v215
	goto L2
L4:
	;
	v199 = v115 + int32(5)
	v200 = F_palloc(m, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L10
	} else {
		goto L82
	}
L5:
	;
	v185 = F_palloc(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L74
	}
L6:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v184 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
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
	v184 = v40
	goto L5
L21:
	;
	v45 = int32(1)
	v184 = int32(base.Ui32(v23)>>(uint(v45)%32)) - v45
	goto L5
L22:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[1])))
	if v124 != 0 {
		goto L52
	} else {
		goto L53
	}
L23:
	;
	if v58&int32(3) == int32(0) {
		v82 = v58
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v54 = F_DirectFunctionCall1Coll(m, int32(580), int32(0), v49)
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
	if v115 != 0 {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v115 = v107 - v58
	goto L29
L31:
	;
	v86 = v82
	goto L40
L32:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v66 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v115 = int32(0)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v71 = v58
	goto L36
L36:
	;
	v75 = v71 + int32(1)
	if v75&int32(3) == int32(0) {
		v82 = v75
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v107 = v75
	goto L30
L38:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v80 != 0 {
		v71 = v75
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v95 = int32(-2139062144)
	if (int32(16843008)-v92|v92)&v95 == v95 {
		v86 = v86 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v101 = v86
	goto L43
L42:
	;
	goto L41
L43:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v105 != 0 {
		v101 = v101 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v107 = v101
	goto L30
L45:
	;
	goto L44
L46:
	;
	v116 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v218 = v121
	v219 = int32(0)
	goto L2
L49:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if v118 != int32(1) {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	if v15 != int32(19) {
		goto L4
	} else {
		goto L68
	}
L52:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2]))
	if v126 == l2 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v129 = int32(_a_F_make_greater_string_0)
	v130 = int32(_a_F_make_greater_string_1)
	v132 = int32(1)
	v135 = F_varstr_cmp(m, v130, v132, v129, v132, l2)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	if v135 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v139 = v129
	goto L59
L58:
	;
	v139 = v130
	goto L59
L59:
	;
	v140 = int32(1)
	v143 = F_varstr_cmp(m, v139, v140, int32(_a_F_make_greater_string_2), v140, l2)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	if v143 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v147 = int32(_a_F_make_greater_string_2)
	goto L63
L62:
	;
	v147 = v139
	goto L63
L63:
	;
	v148 = int32(1)
	v151 = F_varstr_cmp(m, v147, v148, int32(_a_F_make_greater_string_3), v148, l2)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2])) = l2
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v151 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v160 = int32(57)
	goto L67
L66:
	;
	v160 = v157
	goto L67
L67:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[1])) = uint8(v160)
	goto L51
L68:
	;
	v168 = F_palloc(m, v115+int32(2))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L10
	} else {
		goto L69
	}
L69:
	;
	if v115 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[1])))
	v174 = v171 + v115
	v175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)) = uint8(v175)
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v173)
	v215 = v168
	goto L3
L71:
	;
	v170 = F__emscripten_memcpy_bulkmem(m, v168, v58, v115)
	mBase = m.M
	v171 = v170
	goto L73
L72:
	;
	v171 = v168
	goto L73
L73:
	;
	goto L70
L74:
	;
	v187 = int32(1)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v189&v187 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v192 = v187
	goto L77
L76:
	;
	v192 = int32(4)
	goto L77
L77:
	;
	if v184 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v233 = v184
	v235 = v185
	v236 = v196
	v237 = int32(0)
	v238 = int32(1453)
	goto L1
L79:
	;
	v194 = F__emscripten_memcpy_bulkmem(m, v185, v19+v192, v184)
	mBase = m.M
	goto L81
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v199 << (uint(int32(2)) % 32)
	v206 = v200 + int32(4)
	if v115 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v115))) = uint8(v211)
	v215 = v200
	goto L3
L84:
	;
	v207 = F__emscripten_memcpy_bulkmem(m, v206, v58, v115)
	mBase = m.M
	v208 = v207
	goto L86
L85:
	;
	v208 = v206
	goto L86
L86:
	;
	goto L83
L87:
	;
	v228 = int32(1653)
	goto L89
L88:
	;
	v228 = int32(1654)
	goto L89
L89:
	;
	if v225 == int32(6) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v231 = int32(1652)
	goto L92
L91:
	;
	v231 = v228
	goto L92
L92:
	;
	v233 = v115
	v235 = v58
	v236 = v218
	v237 = v219
	v238 = v231
	goto L1
L93:
	;
	F_pfree(m, v235)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L133
	}
L94:
	;
	F_pfree(m, v237)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L10
	} else {
		goto L132
	}
L95:
	;
	v246 = v233
	goto L98
L96:
	;
	goto L97
L97:
	;
	v353 = int32(0)
	if v237 == v353 {
		v372 = v353
		goto L93
	} else {
		goto L131
	}
L98:
	;
	if base.B2i32(v15 == int32(17)) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L97
L100:
	;
	v262 = F_pg_mbcliplen(m, v235, v246, v246-int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L10
	} else {
		goto L103
	}
L101:
	;
	v265 = int32(1)
	goto L102
L102:
	;
	v267 = v246 + v235 - v265
	v268 = m.T0[v238].(func(*base.Module, int32, int32) int32)(m, v267, v265)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L10
	} else {
		goto L104
	}
L103:
	;
	v265 = v246 - v262
	goto L102
L104:
	;
	if v268 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v271 = v246 + int32(4)
	goto L108
L106:
	;
	goto L107
L107:
	;
	v333 = v246 - v265
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235+v333))) = uint8(v335)
	if v335 < v333 {
		v246 = v333
		goto L98
	} else {
		goto L130
	}
L108:
	;
	if v15 == int32(17) {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	goto L107
L110:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+20))
	v310 = F_FunctionCall2Coll(m, l1, l2, v236, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L121
	}
L111:
	;
	v290 = F_palloc(m, v271)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L10
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v305 = F_string_to_const(m, v235, v15)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L10
	} else {
		goto L120
	}
L114:
	;
	if v246 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = v271 << (uint(int32(2)) % 32)
	v298 = int32(-1)
	v299 = int32(0)
	v303 = F_makeConst(m, int32(17), v298, v299, v298, v290, v299, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L119
	}
L116:
	;
	v294 = F__emscripten_memcpy_bulkmem(m, v290+int32(4), v235, v246)
	mBase = m.M
	goto L118
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	v308 = v303
	goto L110
L120:
	;
	v308 = v305
	goto L110
L121:
	;
	if v310 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	if v237 != 0 {
		v356 = v308
		goto L94
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308)+20))
	F_pfree(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L10
	} else {
		goto L126
	}
L125:
	;
	v372 = v308
	goto L93
L126:
	;
	F_pfree(m, v308)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	v317 = m.T0[v238].(func(*base.Module, int32, int32) int32)(m, v267, v265)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L128
	}
L128:
	;
	if v317 != 0 {
		goto L108
	} else {
		goto L129
	}
L129:
	;
	goto L109
L130:
	;
	goto L99
L131:
	;
	v356 = v353
	goto L94
L132:
	;
	v372 = v356
	goto L93
L133:
	;
	return v372
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
	F_errmsg_internal(m, int32(_a_F_make_pathkeys_for_sortclauses_extended_0), v20)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_make_pathkeys_for_sortclauses_extended_1), int32(273), int32(_a_F_make_pathkeys_for_sortclauses_extended_2))
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
	F_errmsg(m, int32(_a_F_make_pathkeys_for_window_0), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	F_errdetail(m, int32(_a_F_make_pathkeys_for_window_1), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L35
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_make_pathkeys_for_window_2), int32(_a_F_make_pathkeys_for_window_3), int32(_a_F_make_pathkeys_for_window_4))
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
	F_errmsg(m, int32(_a_F_make_pathkeys_for_window_5), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L35
	} else {
		goto L51
	}
L51:
	;
	F_errdetail(m, int32(_a_F_make_pathkeys_for_window_6), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L35
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_make_pathkeys_for_window_2), int32(_a_F_make_pathkeys_for_window_7), int32(_a_F_make_pathkeys_for_window_4))
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
								F_errmsg(m, int32(_a_F_make_scalar_array_op_0), int32(0))
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
										F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(804), int32(_a_F_make_scalar_array_op_2))
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
											F_errmsg(m, int32(_a_F_make_scalar_array_op_3), v15)
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
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(819), int32(_a_F_make_scalar_array_op_2))
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
													F_errmsg(m, int32(_a_F_make_scalar_array_op_4), int32(0))
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
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(845), int32(_a_F_make_scalar_array_op_2))
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
															F_errmsg(m, int32(_a_F_make_scalar_array_op_5), int32(0))
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
																	F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(850), int32(_a_F_make_scalar_array_op_2))
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
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
																						F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																								F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
														if base.Ui32(v70-int32(_a_F_make_scalar_array_op_7)) < base.Ui32(int32(4)) {
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
															if base.Ui32(v70-int32(_a_F_make_scalar_array_op_8)) < base.Ui32(int32(2)) {
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
																						F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																								F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
									F_errmsg(m, int32(_a_F_make_scalar_array_op_3), v15)
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
											F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(819), int32(_a_F_make_scalar_array_op_2))
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
											F_errmsg(m, int32(_a_F_make_scalar_array_op_4), int32(0))
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
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(845), int32(_a_F_make_scalar_array_op_2))
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
													F_errmsg(m, int32(_a_F_make_scalar_array_op_5), int32(0))
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
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(850), int32(_a_F_make_scalar_array_op_2))
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
																		F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																				F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
												if base.Ui32(v70-int32(_a_F_make_scalar_array_op_7)) < base.Ui32(int32(4)) {
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
													if base.Ui32(v70-int32(_a_F_make_scalar_array_op_8)) < base.Ui32(int32(2)) {
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
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_match_db_entries[0]))
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
			*(*int32)(unsafe.Add(mBase, uint32(v11)+124)) = int32(_a_F_mdextend_0)
			v31 = F_FileWriteV(m, v18, v11+int32(120), int32(1), base.I64_extend_i32_u(l2<<(uint(int32(13))%32)&int32(1073733632)), int32(167772177))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				if v31 != int32(_a_F_mdextend_0) {
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
								v110 = *(*int32)(unsafe.Add(mBase, _c_F_mdextend[0]))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v108*int32(48))+32))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v114
								F_errmsg(m, int32(_a_F_mdextend_1), v11+int32(16))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_mdextend_2), int32(0))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_mdextend_3), int32(519), int32(_a_F_mdextend_4))
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
							F_errcode(m, int32(_a_F_mdextend_5))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_mdextend[0]))
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v44*int32(48))+32))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(_a_F_mdextend_0)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v31
								*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v50
								F_errmsg(m, int32(_a_F_mdextend_6), v11+int32(32))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_mdextend_2), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_mdextend_3), int32(526), int32(_a_F_mdextend_4))
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
					F_errmsg(m, int32(_a_F_mdextend_7), v11)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_mdextend_3), int32(504), int32(_a_F_mdextend_4))
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
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_mdfd[0]))
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_mdinit[0]))
	v8 = F_AllocSetContextCreateInternal(m, v3, int32(_a_F_mdinit_0), int32(0), int32(_a_F_mdinit_1), int32(_a_F_mdinit_2))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_mdinit[1])) = v8
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
	if base.Ui32(int32(_a_F_mdnblocks_0)) <= base.Ui32(v39) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v39 != int32(_a_F_mdnblocks_1) {
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
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_mdnblocks[0]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69+v67*int32(48))+32))
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v73
	F_errmsg(m, int32(_a_F_mdnblocks_2), v9)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_mdnblocks_3), int32(1882), int32(_a_F_mdnblocks_4))
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
	F_errmsg_internal(m, int32(_a_F_mdnblocks_5), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_mdnblocks_3), int32(1255), int32(_a_F_mdnblocks_6))
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
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[0]))
			if v35&int32(1) != 0 {
				v38 = int32(_a_F_mdopenfork_0)
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
						v47 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[1]))
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
									F_errmsg(m, int32(_a_F_mdopenfork_1), v9)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_mdopenfork_2), int32(686), int32(_a_F_mdopenfork_3))
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
								F_errmsg(m, int32(_a_F_mdopenfork_1), v9)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_mdopenfork_2), int32(686), int32(_a_F_mdopenfork_3))
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
						v75 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[2]))
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
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_mdregistersync[0]))
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
	F_errmsg(m, int32(_a_F_merge_acl_with_grant_0), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_merge_acl_with_grant_1), int32(211), int32(_a_F_merge_acl_with_grant_2))
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
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v561 int32
	_ = v561
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
	return v561
L2:
	;
	v561 = v5
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
	v561 = int32(0)
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
	v138 = int32(_a_F_mkSPNode_0)
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
		v534 = v158
		v536 = v162
		v542 = v172
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v548 = F_mkSPNode(m, l0, v534, l2, v153)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L28
	} else {
		goto L126
	}
L37:
	;
	v546 = v164 + int32(1)
	if v546 != l2 {
		v158 = v534
		v162 = v536
		v164 = v546
		v172 = v542
		goto L35
	} else {
		goto L125
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
		v534 = v199
		v536 = v201
		v542 = v202
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
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v499 | int32(256)
	v513 = F_makeCompoundFlags(m, l0, int32(base.Ui32(v499)>>(uint(int32(13))%32)))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L28
	} else {
		goto L120
	}
L47:
	;
	v499 = v208&int32(_a_F_mkSPNode_1) | v215<<(uint(int32(13))%32)
	v507 = int32(0)
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
	v499 = v480&int32(_a_F_mkSPNode_1) | v481<<(uint(int32(13))%32)
	v507 = v223&int32(base.Ui32(v205)>>(uint(int32(9))%32)) ^ int32(1)
	goto L46
L52:
	;
	v480 = v233
	v481 = v231
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
		v480 = v233
		v481 = v235
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
	v274 = v271<<(uint(int32(2))%32) + v268
	if v269&int32(3) == int32(0) {
		v298 = v269
		goto L62
	} else {
		goto L63
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
	if v270&int32(3) == int32(0) {
		v355 = v270
		goto L79
	} else {
		goto L80
	}
L61:
	;
	v331 = v323 - v269
	goto L60
L62:
	;
	v302 = v298
	goto L71
L63:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v282 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v331 = int32(0)
	goto L60
L65:
	;
	goto L66
L66:
	;
	v287 = v269
	goto L67
L67:
	;
	v291 = v287 + int32(1)
	if v291&int32(3) == int32(0) {
		v298 = v291
		goto L62
	} else {
		goto L69
	}
L68:
	;
	v323 = v291
	goto L61
L69:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291))))
	if v296 != 0 {
		v287 = v291
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v311 = int32(-2139062144)
	if (int32(16843008)-v308|v308)&v311 == v311 {
		v302 = v302 + int32(4)
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v317 = v302
	goto L74
L73:
	;
	goto L72
L74:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
	if v321 != 0 {
		v317 = v317 + int32(1)
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v323 = v317
	goto L61
L76:
	;
	goto L75
L77:
	;
	v389 = v331 + v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v390 == int32(2) {
		goto L95
	} else {
		goto L96
	}
L78:
	;
	v388 = v380 - v270
	goto L77
L79:
	;
	v359 = v355
	goto L88
L80:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if v339 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v388 = int32(0)
	goto L77
L82:
	;
	goto L83
L83:
	;
	v344 = v270
	goto L84
L84:
	;
	v348 = v344 + int32(1)
	if v348&int32(3) == int32(0) {
		v355 = v348
		goto L79
	} else {
		goto L86
	}
L85:
	;
	v380 = v348
	goto L78
L86:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348))))
	if v353 != 0 {
		v344 = v348
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	v368 = int32(-2139062144)
	if (int32(16843008)-v365|v365)&v368 == v368 {
		v359 = v359 + int32(4)
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v374 = v359
	goto L91
L90:
	;
	goto L89
L91:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374))))
	if v378 != 0 {
		v374 = v374 + int32(1)
		goto L91
	} else {
		goto L93
	}
L92:
	;
	v380 = v374
	goto L78
L93:
	;
	goto L92
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v467
	*(*int32)(unsafe.Add(mBase, uint32(v274)+4)) = int32(0)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v474 + int32(1)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	v480 = v478
	v481 = v474
	goto L51
L95:
	;
	v394 = v389 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v394) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	goto L97
L97:
	;
	v430 = v389 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v430) {
		goto L110
	} else {
		goto L111
	}
L98:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v419+v237)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v419+v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v421
	v427 = F_pg_sprintf(m, v416, int32(_a_F_mkSPNode_2), v21)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L28
	} else {
		goto L108
	}
L99:
	;
	v397 = F_palloc0(m, v394)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L28
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v402 = (v389 + int32(9)) & int32(4088)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v402) <= base.Ui32(v403) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v416 = v397
	goto L98
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v410 - v402
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v411 + v402
	v416 = v411
	goto L98
L104:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v410 = v403
	v411 = v405
	goto L103
L105:
	;
	goto L106
L106:
	;
	v406 = int32(_a_F_mkSPNode_0)
	v408 = F_palloc0(m, v406)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L28
	} else {
		goto L107
	}
L107:
	;
	v410 = v406
	v411 = v408
	goto L103
L108:
	;
	v467 = v416
	goto L94
L109:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v455+v237)))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455+v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v457
	v465 = F_pg_sprintf(m, v452, int32(_a_F_mkSPNode_3), v21+int32(16))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L28
	} else {
		goto L119
	}
L110:
	;
	v433 = F_palloc0(m, v430)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L28
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v438 = (v389 + int32(8)) & int32(4088)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v438) <= base.Ui32(v439) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	v452 = v433
	goto L109
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v446 - v438
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v447 + v438
	v452 = v447
	goto L109
L115:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v446 = v439
	v447 = v441
	goto L114
L116:
	;
	goto L117
L117:
	;
	v442 = int32(_a_F_mkSPNode_0)
	v444 = F_palloc0(m, v442)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L28
	} else {
		goto L118
	}
L118:
	;
	v446 = v442
	v447 = v444
	goto L114
L119:
	;
	v467 = v452
	goto L94
L120:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v513 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v525 = v515 | int32(_a_F_mkSPNode_4)
	goto L123
L122:
	;
	v525 = v515&int32(-7681) | v513<<(uint(int32(9))%32)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v525
	if v507&int32(1) == int32(0) {
		v534 = v199
		v536 = v201
		v542 = v202
		goto L37
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v525 & int32(-513)
	v534 = v199
	v536 = v201
	v542 = v202
	goto L37
L125:
	;
	goto L36
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v536)+4)) = v548
	v561 = v150
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
	if int32(1)<<(uint(v38)%32)&int32(_a_F_moveins_0) == int32(0) {
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
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_moveins[0]))
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
	if int32(1)<<(uint(v133)%32)&int32(_a_F_moveins_0) == int32(0) {
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
	if int32(1)<<(uint(v247)%32)&int32(_a_F_moveins_0) == int32(0) {
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1474), int32(0), v4, v5)
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
	v4 = F_SlruSyncFileTag(m, int32(_a_F_multixactmemberssyncfiletag_0), l0, l1)
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
