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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	v58 = int32(1)
	v59 = v52
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v67 = v64 + v58*int32(56)
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
	v82 = v58 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v82 < v83 {
		v58 = v82
		v59 = v80
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
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
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
	return v118
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
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+140))
	if v97 == int32(0) {
		v118 = v23
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
		v89 = v23
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
	v118 = v89
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
	v89 = v23
	goto L15
L19:
	;
	v89 = int32(0)
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
	v83 = v48 + int32(1)
	if v83 != v29 {
		v48 = v83
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
	v89 = int32(0)
	goto L15
L26:
	;
	goto L18
L27:
	;
	v100 = int32(_a_F_MemoizeHash_equal_0)
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v103
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v108 = m.T0[v107].(func(*base.Module, int32, int32, int32) int32)(m, v97, v15, v11+int32(15))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v101
	v118 = base.B2i32(v108 != int32(0))
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
	var v27 int32
	_ = v27
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
	var v1161 int32
	_ = v1161
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
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
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1227 int32
	_ = v1227
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
	v1218 = int32(16)
	v1222 = (int32(base.Ui32(v1209)>>(uint(v1218)%32)) ^ v1209) * int32(-2048144789)
	v1227 = (int32(base.Ui32(v1222)>>(uint(int32(13))%32)) ^ v1222) * int32(-1028477387)
	return int32(base.Ui32(v1227)>>(uint(v1218)%32)) ^ v1227
L2:
	;
	if v13 <= int32(0) {
		v1209 = v2
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
		v1209 = v2
		goto L1
	} else {
		goto L193
	}
L5:
	;
	v25 = int32(0)
	v27 = v2
	goto L6
L6:
	;
	v35 = base.I32_rotl(v27, int32(1))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v25))))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1209 = v1161
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
	v1161 = v35
	goto L10
L10:
	;
	v1166 = v25 + int32(1)
	if v1166 != v13 {
		v25 = v1166
		v27 = v1161
		goto L6
	} else {
		goto L192
	}
L11:
	;
	m.G0 = v54 + int32(16)
	v1161 = v1154 ^ v35
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
	v1175 = v2
	goto L194
L194:
	;
	v1183 = base.I32_rotl(v1175, int32(1))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184+v1173))))
	if v1186 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v1209 = v1203
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
	v1203 = v1183
	goto L198
L198:
	;
	v1205 = v1173 + int32(1)
	if v1205 != v13 {
		v1173 = v1205
		v1175 = v1203
		goto L194
	} else {
		goto L200
	}
L199:
	;
	v1203 = v1199 ^ v1183
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v9 = v6 + l1<<(uint(int32(4))%32)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 == int32(16) {
		v13 = int32(16)
		v14 = l2 - v13
		v15 = int32(0)
		if base.B2i32(v14 == v15)|base.B2i32(v14 == v13) == v15 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_0), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(680), int32(_a_F_ModifyWaitEvent_2))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(base.B2i32(base.Ui32(int32(31)) < base.Ui32(l2)))
			return
		}
	} else {
		if l2 == v10 {
			if l2&int32(1) == int32(0) {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v43 != l3 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
					if l2 == int32(1) {
						if l3 != 0 {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_ModifyWaitEvent[0]))
							if v51 != v53 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_3), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(704), int32(_a_F_ModifyWaitEvent_2))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
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
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v60 = v56 + v57<<(uint(int32(3))%32)
						v61 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)) = uint16(v61)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = v63
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v67 = v65 - int32(1)
						if base.B2i32(v67 == v61)|base.B2i32(v67 == int32(15)) != 0 {
							v99 = int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v99)
						} else {
							v74 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v74)
							v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
							v77 = int32(1)
							v80 = int32(base.Ui32(v76)>>(uint(v77)%32)) & v77
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v80)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							if v82&int32(4) != 0 {
								v86 = v80 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v86)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
								v89 = v86
								v90 = v88
							} else {
								v89 = v80
								v90 = v82
							}
							if v90&int32(128) == int32(0) {
							} else {
								v99 = v89 | int32(_a_F_ModifyWaitEvent_4)
								*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v99)
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
				v108 = m.ExcPending
				if v108 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_5), int32(0))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(696), int32(_a_F_ModifyWaitEvent_2))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
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
				if l2 == int32(1) {
					if l3 != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_ModifyWaitEvent[0]))
						if v51 != v53 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_ModifyWaitEvent_3), int32(0))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ModifyWaitEvent_1), int32(704), int32(_a_F_ModifyWaitEvent_2))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
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
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v60 = v56 + v57<<(uint(int32(3))%32)
					v61 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v60)+6)) = uint16(v61)
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v60))) = v63
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v67 = v65 - int32(1)
					if base.B2i32(v67 == v61)|base.B2i32(v67 == int32(15)) != 0 {
						v99 = int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v99)
					} else {
						v74 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v74)
						v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
						v77 = int32(1)
						v80 = int32(base.Ui32(v76)>>(uint(v77)%32)) & v77
						*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v80)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						if v82&int32(4) != 0 {
							v86 = v80 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v86)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v89 = v86
							v90 = v88
						} else {
							v89 = v80
							v90 = v82
						}
						if v90&int32(128) == int32(0) {
						} else {
							v99 = v89 | int32(_a_F_ModifyWaitEvent_4)
							*(*uint16)(unsafe.Add(mBase, uint32(v60)+4)) = uint16(v99)
						}
					}
					return
				}
			}
		}
	}
}
func F___memcpy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	if base.Ui32(int32(512)) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v11 = l0 + l2
	if (l0^l1)&int32(3) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	base.MemoryCopy(m, l0, l1, l2)
	goto L6
L5:
	;
	goto L6
L6:
	;
	return l0
L7:
	;
	if base.Ui32(v143) < base.Ui32(v11) {
		goto L41
	} else {
		goto L42
	}
L8:
	;
	if l0&int32(3) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(v11) < base.Ui32(int32(4)) {
		goto L32
	} else {
		goto L33
	}
L11:
	;
	v47 = v11 & int32(-4)
	if base.Ui32(v11) < base.Ui32(int32(64)) {
		v97 = v41
		v98 = v42
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v41 = l1
	v42 = l0
	goto L11
L13:
	;
	goto L14
L14:
	;
	if l2 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v41 = l1
	v42 = l0
	goto L11
L16:
	;
	goto L17
L17:
	;
	v24 = l1
	v25 = l0
	goto L18
L18:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25))) = uint8(v29)
	v31 = int32(1)
	v32 = v24 + v31
	v34 = v25 + v31
	if v34&int32(3) == int32(0) {
		v41 = v32
		v42 = v34
		goto L11
	} else {
		goto L20
	}
L19:
	;
	v41 = v32
	v42 = v34
	goto L11
L20:
	;
	if base.Ui32(v34) < base.Ui32(v11) {
		v24 = v32
		v25 = v34
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	if base.Ui32(v47) <= base.Ui32(v98) {
		v142 = v97
		v143 = v98
		goto L7
	} else {
		goto L28
	}
L23:
	;
	v51 = v47 + int32(-64)
	if base.Ui32(v51) < base.Ui32(v42) {
		v97 = v41
		v98 = v42
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v54 = v41
	v55 = v42
	goto L25
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+16)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+20)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v54)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v54)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+28)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+32)) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+36)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+40)) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v54)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+44)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+48)) = v83
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+52)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+56)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v54)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+60)) = v89
	v91 = int32(-64)
	v92 = v54 - v91
	v94 = v55 - v91
	if base.Ui32(v94) <= base.Ui32(v51) {
		v54 = v92
		v55 = v94
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v97 = v92
	v98 = v94
	goto L22
L27:
	;
	goto L26
L28:
	;
	v104 = v97
	v105 = v98
	goto L29
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v109
	v111 = int32(4)
	v112 = v104 + v111
	v114 = v105 + v111
	if base.Ui32(v114) < base.Ui32(v47) {
		v104 = v112
		v105 = v114
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v142 = v112
	v143 = v114
	goto L7
L31:
	;
	goto L30
L32:
	;
	v142 = l1
	v143 = l0
	goto L7
L33:
	;
	goto L34
L34:
	;
	if base.Ui32(l2) < base.Ui32(int32(4)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v142 = l1
	v143 = l0
	goto L7
L36:
	;
	goto L37
L37:
	;
	v123 = l1
	v124 = l0
	goto L38
L38:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	*(*uint8)(unsafe.Add(mBase, uint32(v124))) = uint8(v128)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+1)) = uint8(v130)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+2)) = uint8(v132)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+3)) = uint8(v134)
	v136 = int32(4)
	v137 = v123 + v136
	v139 = v124 + v136
	if base.Ui32(v139) <= base.Ui32(v11-int32(4)) {
		v123 = v137
		v124 = v139
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v142 = v137
	v143 = v139
	goto L7
L40:
	;
	goto L39
L41:
	;
	v149 = v142
	v150 = v143
	goto L44
L42:
	;
	goto L43
L43:
	;
	return l0
L44:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v154)
	v156 = int32(1)
	v159 = v150 + v156
	if v159 != v11 {
		v149 = v149 + v156
		v150 = v159
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
}
func F___memset(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v2 = l1
	if l2 == int32(0) {
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
		v10 = l0 + l2
		*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(1)))) = uint8(v2)
		if base.Ui32(l2) < base.Ui32(int32(3)) {
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v2)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v2)
			*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(3)))) = uint8(v2)
			*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(2)))) = uint8(v2)
			if base.Ui32(l2) < base.Ui32(int32(7)) {
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v2)
				*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(4)))) = uint8(v2)
				if base.Ui32(l2) < base.Ui32(int32(9)) {
				} else {
					v35 = (int32(0) - l0) & int32(3)
					v36 = l0 + v35
					v40 = v2 & int32(255) * int32(16843009)
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v40
					v44 = (l2 - v35) & int32(-4)
					v45 = v36 + v44
					*(*int32)(unsafe.Add(mBase, uint32(v45-int32(4)))) = v40
					if base.Ui32(v44) < base.Ui32(int32(9)) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v45-int32(8)))) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v45-int32(12)))) = v40
						if base.Ui32(v44) < base.Ui32(int32(25)) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(16)))) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(20)))) = v40
							v71 = int32(24)
							*(*int32)(unsafe.Add(mBase, uint32(v45-v71))) = v40
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(28)))) = v40
							v80 = v36&int32(4) | v71
							v81 = v44 - v80
							if base.Ui32(v81) < base.Ui32(int32(32)) {
							} else {
								v86 = base.I64_extend_i32_u(v40) * int64(4294967297)
								v89 = v80 + v36
								v90 = v81
								for {
									*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89))) = v86
									v98 = int32(32)
									v101 = v90 - v98
									if base.Ui32(int32(31)) < base.Ui32(v101) {
										v89 = v89 + v98
										v90 = v101
										continue
									} else {
										break
									}
									break
								}
							}
						}
					}
				}
			}
		}
	}
	return
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
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
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
	v39 = v9 + int32(-12)
	F_add_exact_object_address(m, v39, v30)
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
	F_add_exact_object_address(m, v39, v30)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v57 = m.ExcPending
	if v57 != 0 {
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
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v59 = v9 + int32(-60)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v59, int32(1), int32(3), int32(184), v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v67 = int32(1)
	v70 = F_systable_beginscan(m, l3, int32(3609), v67, int32(0), v67, v59)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v72 = F_systable_getnext(m, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if v72 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = v72
	goto L22
L20:
	;
	goto L21
L21:
	;
	F_systable_endscan(m, v70)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L27
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(3600)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v87
	F_add_exact_object_address(m, v9+int32(-12), v30)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L21
L24:
	;
	v95 = F_systable_getnext(m, v70)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v95 != 0 {
		v75 = v95
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
	v119 = m.ExcPending
	if v119 != 0 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v6 = l5
	v7 = l6
	v10 = F_palloc0(m, int32(32))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(7)
		if base.B2i32(l3 != int32(-1))|v6 == int32(0) {
			v21 = F_pg_detoast_datum(m, l4)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = v21
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v7)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v6)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				return v10
			}
		} else {
			v23 = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(-1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v7)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)) = uint8(v6)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v23
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
			return v10
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v2 {
		v29 = v2
		m.G0 = v7 + int32(16)
		return v29
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = v11 - int32(7)
		if v13 != 0 {
			if v13 != int32(14) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
				v25 = F_list_make1_impl(m, int32(1), v7+int32(8))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v25
					m.G0 = v7 + int32(16)
					return v29
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v16 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
					v25 = F_list_make1_impl(m, int32(1), v7+int32(8))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v25
						m.G0 = v7 + int32(16)
						return v29
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v29 = v17
					m.G0 = v7 + int32(16)
					return v29
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v18 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
				v25 = F_list_make1_impl(m, int32(1), v7+int32(8))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v25
					m.G0 = v7 + int32(16)
					return v29
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 != 0 {
					v29 = v2
					m.G0 = v7 + int32(16)
					return v29
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
					v25 = F_list_make1_impl(m, int32(1), v7+int32(8))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v25
						m.G0 = v7 + int32(16)
						return v29
					}
				}
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
	var v4 int32
	_ = v4
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
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
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v344 int32
	_ = v344
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(17) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if int32(0) < v192 {
		goto L89
	} else {
		goto L90
	}
L2:
	;
	v179 = F_palloc(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L80
	}
L3:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v178 = int32(base.Ui32(v172)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v15 == int32(19) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	return int32(0)
L8:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v23 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v29 == int32(18) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if v23&int32(1) == int32(0) {
		goto L3
	} else {
		goto L18
	}
L12:
	;
	v32 = int32(16)
	goto L14
L13:
	;
	v32 = int32(0)
	goto L14
L14:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(4)
	goto L17
L16:
	;
	v39 = v32
	goto L17
L17:
	;
	v178 = v39
	goto L2
L18:
	;
	v44 = int32(1)
	v178 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
	goto L2
L19:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	if v79 != 0 {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v58 = F_strlen(m, v57)
	mBase = m.M
	if v58 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v53 = F_DirectFunctionCall1Coll(m, int32(581), int32(0), v48)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v55 = F_text_to_cstring(m, v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v57 = v53
	goto L20
L25:
	;
	v57 = v55
	goto L20
L26:
	;
	v59 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v71 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	if v61 != int32(1) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v192 = v58
	v193 = v64
	v194 = v57
	v195 = v4
	v196 = v77
	goto L1
L32:
	;
	v74 = int32(1638)
	goto L34
L33:
	;
	v74 = int32(1639)
	goto L34
L34:
	;
	if v71 == int32(6) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v77 = int32(1637)
	goto L37
L36:
	;
	v77 = v74
	goto L37
L37:
	;
	goto L31
L38:
	;
	if v15 == int32(19) {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2]))
	if v81 == l2 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v84 = int32(_a_F_make_greater_string_0)
	v85 = int32(_a_F_make_greater_string_1)
	v87 = int32(1)
	v90 = F_varstr_cmp(m, v85, v87, v84, v87, l2)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	if v90 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v94 = v84
	goto L46
L45:
	;
	v94 = v85
	goto L46
L46:
	;
	v95 = int32(1)
	v98 = F_varstr_cmp(m, v94, v95, int32(_a_F_make_greater_string_2), v95, l2)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	if v98 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v102 = int32(_a_F_make_greater_string_2)
	goto L50
L49:
	;
	v102 = v94
	goto L50
L50:
	;
	v103 = int32(1)
	v106 = F_varstr_cmp(m, v102, v103, int32(_a_F_make_greater_string_3), v103, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2])) = l2
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v115 = int32(57)
	goto L54
L53:
	;
	v115 = v112
	goto L54
L54:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])) = uint8(v115)
	goto L38
L55:
	;
	v123 = F_palloc(m, v58+int32(2))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v146 = v58 + int32(5)
	v147 = F_palloc(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L69
	}
L58:
	;
	if v58 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v123, v57, v58)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	v128 = v58 + v123
	v129 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+1)) = uint8(v129)
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v127)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v138 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v192 = v58
	v193 = v123
	v194 = v57
	v195 = v123
	v196 = v144
	goto L1
L63:
	;
	v141 = int32(1638)
	goto L65
L64:
	;
	v141 = int32(1639)
	goto L65
L65:
	;
	if v138 == int32(6) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v144 = int32(1637)
	goto L68
L67:
	;
	v144 = v141
	goto L68
L68:
	;
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v146 << (uint(int32(2)) % 32)
	v153 = v147 + int32(4)
	if v58 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	base.MemoryCopy(m, v153, v57, v58)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v153+v58))) = uint8(v157)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v165 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v192 = v58
	v193 = v147
	v194 = v57
	v195 = v147
	v196 = v171
	goto L1
L74:
	;
	v168 = int32(1638)
	goto L76
L75:
	;
	v168 = int32(1639)
	goto L76
L76:
	;
	if v165 == int32(6) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v171 = int32(1637)
	goto L79
L78:
	;
	v171 = v168
	goto L79
L79:
	;
	goto L73
L80:
	;
	if v178 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v181 = int32(1)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v183&v181 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v192 = v178
	v193 = v189
	v194 = v179
	v195 = v4
	v196 = int32(1438)
	goto L1
L84:
	;
	v186 = v181
	goto L86
L85:
	;
	v186 = int32(4)
	goto L86
L86:
	;
	base.MemoryCopy(m, v179, v19+v186, v178)
	goto L83
L87:
	;
	F_pfree(m, v194)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L7
	} else {
		goto L126
	}
L88:
	;
	F_pfree(m, v195)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L125
	}
L89:
	;
	v204 = v192
	goto L92
L90:
	;
	goto L91
L91:
	;
	v310 = int32(0)
	if v195 == v310 {
		v329 = v310
		goto L87
	} else {
		goto L124
	}
L92:
	;
	if base.B2i32(v15 == int32(17)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L91
L94:
	;
	v220 = F_pg_mbcliplen(m, v194, v204, v204-int32(1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L97
	}
L95:
	;
	v223 = int32(1)
	goto L96
L96:
	;
	v225 = v204 + v194 - v223
	v226 = m.T0[v196].(func(*base.Module, int32, int32) int32)(m, v225, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L7
	} else {
		goto L98
	}
L97:
	;
	v223 = v204 - v220
	goto L96
L98:
	;
	if v226 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v229 = v204 + int32(4)
	goto L102
L100:
	;
	goto L101
L101:
	;
	v290 = v204 - v223
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v194+v290))) = uint8(v292)
	if v292 < v290 {
		v204 = v290
		goto L92
	} else {
		goto L123
	}
L102:
	;
	if v15 == int32(17) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L101
L104:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	v267 = F_FunctionCall2Coll(m, l1, l2, v193, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L7
	} else {
		goto L114
	}
L105:
	;
	v248 = F_palloc(m, v229)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v262 = F_string_to_const(m, v194, v15)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L7
	} else {
		goto L113
	}
L108:
	;
	if v204 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	base.MemoryCopy(m, v248+int32(4), v194, v204)
	goto L111
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v229 << (uint(int32(2)) % 32)
	v255 = int32(-1)
	v256 = int32(0)
	v260 = F_makeConst(m, int32(17), v255, v256, v255, v248, v256, v256)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v265 = v260
	goto L104
L113:
	;
	v265 = v262
	goto L104
L114:
	;
	if v267 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v195 != 0 {
		v313 = v265
		goto L88
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+20))
	F_pfree(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L119
	}
L118:
	;
	v329 = v265
	goto L87
L119:
	;
	F_pfree(m, v265)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v274 = m.T0[v196].(func(*base.Module, int32, int32) int32)(m, v225, v223)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if v274 != 0 {
		goto L102
	} else {
		goto L122
	}
L122:
	;
	goto L103
L123:
	;
	goto L93
L124:
	;
	v313 = v310
	goto L88
L125:
	;
	v329 = v313
	goto L87
L126:
	;
	return v329
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
	var v222 int32
	_ = v222
	v8 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v24 == v8 {
		v222 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v222
L2:
	;
	v38 = v8
	v39 = v8
	v40 = v24
	goto L3
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 <= v38 {
		v222 = v39
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v222 = v203
	goto L1
L5:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v38<<(uint(int32(2))%32))))
	v51 = F_get_sortgroupclause_expr(m, v50, l2)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v55 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v204 != 0 {
		v38 = v202 + int32(1)
		v39 = v203
		v40 = v204
		goto L3
	} else {
		goto L41
	}
L9:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v58)
	v202 = v38
	v203 = v39
	v204 = v40
	goto L8
L10:
	;
	goto L11
L11:
	;
	if l4 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v61 = F_bms_make_singleton(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v67 = v51
	v68 = v55
	goto L14
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+17)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+16)))
	v78 = F_get_ordering_op_properties(m, v68, v20+int32(12), v20+int32(8), v20+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L18
	}
L15:
	;
	v64 = F_remove_nulling_relids(m, v51, v61, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v67 = v64
	v68 = v66
	goto L14
L17:
	;
	v189 = F_lappend(m, v39, v90)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L40
	}
L18:
	;
	if v78 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v80 = F_exprCollation(m, v67)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L6
	} else {
		goto L37
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v84 = int32(1)
	v90 = F_make_pathkey_from_sortinfo(m, l0, v67, v82, v83, v80, v71&v84, v70&v84, v69, int32(0), v84)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if l6 == int32(0) {
		v99 = v92
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+40)))
	if v100 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+44))
	if v95 != 0 {
		v99 = v92
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+44)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v99 = v98
	goto L24
L27:
	;
	if l3 == int32(0) {
		v202 = v38
		v203 = v39
		v204 = v40
		goto L8
	} else {
		goto L35
	}
L28:
	;
	if v39 == int32(0) {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v103 <= int32(0) {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v115 = int32(0)
	goto L31
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v106+v115<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v99 == v129 {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L17
L33:
	;
	v132 = v115 + int32(1)
	if v132 != v103 {
		v115 = v132
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v154 = F_list_delete_nth_cell(m, v153, v38)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v154
	v202 = v38 - int32(1)
	v203 = v39
	v204 = v154
	goto L8
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v68
	F_errmsg_internal(m, int32(_a_F_make_pathkeys_for_sortclauses_extended_0), v20)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_make_pathkeys_for_sortclauses_extended_1), int32(273), int32(_a_F_make_pathkeys_for_sortclauses_extended_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v202 = v38
	v203 = v189
	v204 = v40
	goto L8
L41:
	;
	goto L4
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
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
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v133 = m.ExcPending
							if v133 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_make_scalar_array_op_0), int32(0))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, l0, l5)
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(804), int32(_a_F_make_scalar_array_op_2))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
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
								v148 = m.ExcPending
								if v148 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return int32(0)
									} else {
										v152 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
										v154 = F_op_signature_string(m, l1, v152, v153)
										mBase = m.M
										v155 = m.ExcPending
										if v155 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15))) = v154
											F_errmsg(m, int32(_a_F_make_scalar_array_op_3), v15)
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, l0, l5)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(819), int32(_a_F_make_scalar_array_op_2))
													mBase = m.M
													v166 = m.ExcPending
													if v166 != 0 {
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
											v170 = m.ExcPending
											if v170 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_make_scalar_array_op_4), int32(0))
													mBase = m.M
													v177 = m.ExcPending
													if v177 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(845), int32(_a_F_make_scalar_array_op_2))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
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
													v188 = m.ExcPending
													if v188 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(151027844))
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_make_scalar_array_op_5), int32(0))
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
																return int32(0)
															} else {
																F_parser_errposition(m, l0, l5)
																mBase = m.M
																v197 = m.ExcPending
																if v197 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(850), int32(_a_F_make_scalar_array_op_2))
																	mBase = m.M
																	v202 = m.ExcPending
																	if v202 != 0 {
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
															v95 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = F_palloc0(m, int32(36))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v105
																	}
																}
															}
														case 1, 2, 3, 4, 5:
															v91 = F_get_array_type(m, v70)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																if v91 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v206 = m.ExcPending
																	if v206 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return int32(0)
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v211 = F_format_type_be(m, v210)
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v220 = m.ExcPending
																					if v220 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																						mBase = m.M
																						v225 = m.ExcPending
																						if v225 != 0 {
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
																	v95 = v91
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v105 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v105
																			}
																		}
																	}
																}
															}
														default:
															if v70 == int32(2776) {
																v95 = v22
																*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v105 = F_palloc0(m, int32(36))
																	mBase = m.M
																	v106 = m.ExcPending
																	if v106 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																		v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																		v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																		v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																		*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																		v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																		*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																		*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																		*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																		*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																		*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																		F_ReleaseCatCache(m, v32)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v15 - int32(-64)
																			return v105
																		}
																	}
																}
															} else {
																if v70 != int32(3500) {
																	v91 = F_get_array_type(m, v70)
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int32(0)
																	} else {
																		if v91 == int32(0) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v206 = m.ExcPending
																			if v206 != 0 {
																				return int32(0)
																			} else {
																				F_errcode(m, int32(67137668))
																				mBase = m.M
																				v209 = m.ExcPending
																				if v209 != 0 {
																					return int32(0)
																				} else {
																					v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																					v211 = F_format_type_be(m, v210)
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																						F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																						mBase = m.M
																						v218 = m.ExcPending
																						if v218 != 0 {
																							return int32(0)
																						} else {
																							F_parser_errposition(m, l0, l5)
																							mBase = m.M
																							v220 = m.ExcPending
																							if v220 != 0 {
																								return int32(0)
																							} else {
																								F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																								mBase = m.M
																								v225 = m.ExcPending
																								if v225 != 0 {
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
																			v95 = v91
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																			*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																			F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return int32(0)
																			} else {
																				v105 = F_palloc0(m, int32(36))
																				mBase = m.M
																				v106 = m.ExcPending
																				if v106 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																					v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																					v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																					v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																					*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																					v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																					*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																					*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																					*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																					*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																					*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																					F_ReleaseCatCache(m, v32)
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v15 - int32(-64)
																						return v105
																					}
																				}
																			}
																		}
																	}
																} else {
																	v95 = v22
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v105 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v105
																			}
																		}
																	}
																}
															}
														}
													} else {
														if base.B2i32(base.Ui32(v70-int32(_a_F_make_scalar_array_op_7)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v70-int32(_a_F_make_scalar_array_op_8)) < base.Ui32(int32(2)))|base.B2i32(v70 == int32(3831)) != 0 {
															v95 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = F_palloc0(m, int32(36))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v105
																	}
																}
															}
														} else {
															v91 = F_get_array_type(m, v70)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																if v91 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v206 = m.ExcPending
																	if v206 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return int32(0)
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v211 = F_format_type_be(m, v210)
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v220 = m.ExcPending
																					if v220 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																						mBase = m.M
																						v225 = m.ExcPending
																						if v225 != 0 {
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
																	v95 = v91
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v105 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v105
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
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(52461700))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return int32(0)
							} else {
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v36)+84))
								v154 = F_op_signature_string(m, l1, v152, v153)
								mBase = m.M
								v155 = m.ExcPending
								if v155 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v154
									F_errmsg(m, int32(_a_F_make_scalar_array_op_3), v15)
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										F_parser_errposition(m, l0, l5)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(819), int32(_a_F_make_scalar_array_op_2))
											mBase = m.M
											v166 = m.ExcPending
											if v166 != 0 {
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
									v170 = m.ExcPending
									if v170 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_make_scalar_array_op_4), int32(0))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int32(0)
											} else {
												F_parser_errposition(m, l0, l5)
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(845), int32(_a_F_make_scalar_array_op_2))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
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
											v188 = m.ExcPending
											if v188 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(151027844))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_make_scalar_array_op_5), int32(0))
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														F_parser_errposition(m, l0, l5)
														mBase = m.M
														v197 = m.ExcPending
														if v197 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(850), int32(_a_F_make_scalar_array_op_2))
															mBase = m.M
															v202 = m.ExcPending
															if v202 != 0 {
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
													v95 = v22
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
													F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v105 = F_palloc0(m, int32(36))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
															v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
															*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
															*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
															*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
															*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
															F_ReleaseCatCache(m, v32)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int32(0)
															} else {
																m.G0 = v15 - int32(-64)
																return v105
															}
														}
													}
												case 1, 2, 3, 4, 5:
													v91 = F_get_array_type(m, v70)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														if v91 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v206 = m.ExcPending
															if v206 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67137668))
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																	v211 = F_format_type_be(m, v210)
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																		F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			F_parser_errposition(m, l0, l5)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																				mBase = m.M
																				v225 = m.ExcPending
																				if v225 != 0 {
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
															v95 = v91
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = F_palloc0(m, int32(36))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v105
																	}
																}
															}
														}
													}
												default:
													if v70 == int32(2776) {
														v95 = v22
														*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
														*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
														F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v105 = F_palloc0(m, int32(36))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																F_ReleaseCatCache(m, v32)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v15 - int32(-64)
																	return v105
																}
															}
														}
													} else {
														if v70 != int32(3500) {
															v91 = F_get_array_type(m, v70)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																if v91 == int32(0) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v206 = m.ExcPending
																	if v206 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(67137668))
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return int32(0)
																		} else {
																			v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																			v211 = F_format_type_be(m, v210)
																			mBase = m.M
																			v212 = m.ExcPending
																			if v212 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																				F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					F_parser_errposition(m, l0, l5)
																					mBase = m.M
																					v220 = m.ExcPending
																					if v220 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																						mBase = m.M
																						v225 = m.ExcPending
																						if v225 != 0 {
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
																	v95 = v91
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
																	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
																	F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v105 = F_palloc0(m, int32(36))
																		mBase = m.M
																		v106 = m.ExcPending
																		if v106 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																			v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																			v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																			v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																			v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																			*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																			*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																			*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																			F_ReleaseCatCache(m, v32)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v15 - int32(-64)
																				return v105
																			}
																		}
																	}
																}
															}
														} else {
															v95 = v22
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = F_palloc0(m, int32(36))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v105
																	}
																}
															}
														}
													}
												}
											} else {
												if base.B2i32(base.Ui32(v70-int32(_a_F_make_scalar_array_op_7)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v70-int32(_a_F_make_scalar_array_op_8)) < base.Ui32(int32(2)))|base.B2i32(v70 == int32(3831)) != 0 {
													v95 = v22
													*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
													*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
													F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v105 = F_palloc0(m, int32(36))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
															v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
															v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
															v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
															*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
															*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
															*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
															*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
															*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
															F_ReleaseCatCache(m, v32)
															mBase = m.M
															v122 = m.ExcPending
															if v122 != 0 {
																return int32(0)
															} else {
																m.G0 = v15 - int32(-64)
																return v105
															}
														}
													}
												} else {
													v91 = F_get_array_type(m, v70)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														if v91 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v206 = m.ExcPending
															if v206 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(67137668))
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	v210 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
																	v211 = F_format_type_be(m, v210)
																	mBase = m.M
																	v212 = m.ExcPending
																	if v212 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v211
																		F_errmsg(m, int32(_a_F_make_scalar_array_op_6), v13+int32(-48))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			F_parser_errposition(m, l0, l5)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(871), int32(_a_F_make_scalar_array_op_2))
																				mBase = m.M
																				v225 = m.ExcPending
																				if v225 != 0 {
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
															v95 = v91
															*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v95
															*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v22
															F_make_fn_arguments(m, l0, v48, v13+int32(-8), v13+int32(-16))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v105 = F_palloc0(m, int32(36))
																mBase = m.M
																v106 = m.ExcPending
																if v106 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(20)
																	v109 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
																	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+22)))
																	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109+v110)))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+100))
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+32)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v48
																	*(*uint8)(unsafe.Add(mBase, uint32(v105)+20)) = uint8(v3)
																	*(*int64)(unsafe.Add(mBase, uint32(v105)+12)) = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v114
																	F_ReleaseCatCache(m, v32)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v15 - int32(-64)
																		return v105
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
	var v86 int32
	_ = v86
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
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
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
			*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(_a_F_mdextend_0)
			v31 = F_FileWriteV(m, v18, v11+int32(56), int32(1), base.I64_extend_i32_u(l2<<(uint(int32(13))%32)&int32(1073733632)), int32(167772177))
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
							v105 = m.ExcPending
							if v105 != 0 {
								return
							} else {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
								v108 = *(*int32)(unsafe.Add(mBase, _c_F_mdextend[0]))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v106*int32(48))+32))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v112
								F_errmsg(m, int32(_a_F_mdextend_1), v11+int32(16))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_mdextend_2), int32(0))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_mdextend_3), int32(519), int32(_a_F_mdextend_4))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
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
				v86 = v11 + int32(56)
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_GetRelationPath(m, v86, v87, v88, v89, v90, l1)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v86
					F_errmsg(m, int32(_a_F_mdextend_7), v11)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_mdextend_3), int32(504), int32(_a_F_mdextend_4))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
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
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v97 int32
	_ = v97
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v13 = l0 + l1<<(uint(int32(2))%32)
	v15 = v13 + int32(40)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if int32(0) < v16 {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
		v97 = v19
		m.G0 = v9 + int32(80)
		return v97
	} else {
		v21 = v9 + int32(8)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_GetRelationPath(m, v21, v22, v23, v24, v25, l1)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[0]))
			if v33&int32(1) != 0 {
				v36 = int32(_a_F_mdopenfork_0)
			} else {
				v36 = int32(2)
			}
			v37 = F_PathNameOpenFile(m, v21, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 < int32(0) {
					if l2&int32(2) != 0 {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[1]))
						if v45 == int32(44) {
							v97 = int32(0)
							m.G0 = v9 + int32(80)
							return v97
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errcode_for_file_access(m)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(8)
									F_errmsg(m, int32(_a_F_mdopenfork_1), v9)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_mdopenfork_2), int32(686), int32(_a_F_mdopenfork_3))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
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
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errcode_for_file_access(m)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(8)
								F_errmsg(m, int32(_a_F_mdopenfork_1), v9)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_mdopenfork_2), int32(686), int32(_a_F_mdopenfork_3))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
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
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					if v66 == int32(0) {
						v73 = *(*int32)(unsafe.Add(mBase, _c_F_mdopenfork[2]))
						v75 = F_MemoryContextAlloc(m, v73, int32(8))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56)) = v75
							v90 = v75
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90))) = v37
							v97 = v90
							m.G0 = v9 + int32(80)
							return v97
						}
					} else {
						v82 = l0 + l1<<(uint(int32(2))%32) + int32(56)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
						if int32(0) < v66 {
							v90 = v83
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v90))) = v37
							v97 = v90
							m.G0 = v9 + int32(80)
							return v97
						} else {
							v87 = F_repalloc(m, v83, int32(8))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v82))) = v87
								v90 = v87
								*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v90))) = v37
								v97 = v90
								m.G0 = v9 + int32(80)
								return v97
							}
						}
					}
				}
			}
		}
	}
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L27
	} else {
		goto L31
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v73
L3:
	;
	v73 = l0
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v19 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v73 = l0
	goto L2
L7:
	;
	goto L8
L8:
	;
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v24 = int32(1)
	goto L11
L10:
	;
	v24 = int32(2)
	goto L11
L11:
	;
	v26 = l5 & int64(4294967295)
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = v26
	goto L14
L13:
	;
	v28 = int64(0)
	goto L14
L14:
	;
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v29 = v28
	goto L17
L16:
	;
	v29 = v26
	goto L17
L17:
	;
	v31 = l5 << (uint(int64(32)) % 64)
	if l1 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = int64(0)
	goto L20
L19:
	;
	v33 = v31
	goto L20
L20:
	;
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v34 = v31
	goto L23
L22:
	;
	v34 = v33
	goto L23
L23:
	;
	v38 = l0
	v39 = int32(0)
	goto L24
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v39<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v54
	if l1&l2&base.B2i32(v54 == int32(0)) != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v73 = v61
	goto L2
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v29 | v34
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = l6
	v61 = F_aclupdate(m, v38, v15, v24, l7, l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	return int32(0)
L28:
	;
	F_pfree(m, v38)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v68 = v39 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v68 < v69 {
		v38 = v61
		v39 = v68
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L25
L31:
	;
	F_errcode(m, int32(16910080))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(_a_F_merge_acl_with_grant_0), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L27
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_merge_acl_with_grant_1), int32(211), int32(_a_F_merge_acl_with_grant_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L27
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v788 int32
	_ = v788
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
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
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L6
	} else {
		goto L357
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L6
	} else {
		goto L353
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L6
	} else {
		goto L349
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L6
	} else {
		goto L346
	}
L5:
	;
	v963 = F_cstring_to_text(m, v957)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L6
	} else {
		goto L345
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
	v957 = int32(_a_F_metaphone_0)
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
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v36 = F_palloc(m, v26+int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.Ui32(v38-int32(97)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	switch v87 - int32(65) {
	case 0:
		goto L40
	default:
		v188 = v85
		v191 = v2
		goto L34
	case 4, 8, 14, 20:
		goto L36
	case 6, 10, 15:
		goto L39
	case 22:
		goto L38
	case 23:
		goto L37
	}
L17:
	;
	if base.Ui32(int32(-27)) < base.Ui32(v45&int32(223)-int32(91)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v45 = v38 & int32(95)
	goto L20
L19:
	;
	v45 = v38
	goto L20
L20:
	;
	goto L17
L21:
	;
	v85 = int32(0)
	v86 = v16
	v87 = v45
	goto L16
L22:
	;
	goto L23
L23:
	;
	v54 = int32(0)
	v56 = v45
	goto L24
L24:
	;
	if v56 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v85 = v69
	v86 = v70
	v87 = v78
	goto L16
L26:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v66)
	v957 = v36
	goto L5
L27:
	;
	goto L28
L28:
	;
	v69 = v54 + int32(1)
	v70 = v16 + v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui32(v71-int32(97)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(v78&int32(223)-int32(91)) < base.Ui32(int32(-26)) {
		v54 = v69
		v56 = v78
		goto L24
	} else {
		goto L33
	}
L30:
	;
	v78 = v71 & int32(95)
	goto L32
L31:
	;
	v78 = v71
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L25
L34:
	;
	v192 = v188 + v16
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if base.Ui32(v193-int32(97)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L35:
	;
	v185 = int32(1)
	v188 = v85 + v185
	v191 = v185
	goto L34
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v87)
	goto L35
L37:
	;
	v182 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v182)
	goto L35
L38:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if base.Ui32(v132-int32(97)) < base.Ui32(int32(26)) {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if base.Ui32(v117-int32(97)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v97 = int32(1)
	v99 = v85 + v97
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v99))))
	if base.Ui32(v101-int32(97)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v108 == int32(69) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v108 = v101 & int32(95)
	goto L44
L43:
	;
	v108 = v101
	goto L44
L44:
	;
	goto L41
L45:
	;
	v111 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v111)
	v188 = v85 + int32(2)
	v191 = v97
	goto L34
L46:
	;
	goto L47
L47:
	;
	v115 = int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v115)
	v188 = v99
	v191 = v97
	goto L34
L48:
	;
	if v124 != int32(78) {
		v188 = v85
		v191 = v2
		goto L34
	} else {
		goto L52
	}
L49:
	;
	v124 = v117 & int32(95)
	goto L51
L50:
	;
	v124 = v117
	goto L51
L51:
	;
	goto L48
L52:
	;
	v127 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v127)
	v188 = v85 + int32(2)
	v191 = int32(1)
	goto L34
L53:
	;
	v146 = int32(0)
	v149 = base.I32_extend8_s(v139) & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v149|int32(32)-int32(97)) {
		v171 = v146
		goto L59
	} else {
		goto L60
	}
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v139)
	v188 = v85 + int32(2)
	v191 = int32(1)
	goto L34
L55:
	;
	switch v139 - int32(72) {
	case 0, 10:
		goto L54
	default:
		goto L53
	}
L56:
	;
	v139 = v132 & int32(95)
	goto L58
L57:
	;
	v139 = v132
	goto L58
L58:
	;
	goto L55
L59:
	;
	if v171&int32(1) == int32(0) {
		v188 = v85
		v191 = v2
		goto L34
	} else {
		goto L66
	}
L60:
	;
	if base.Ui32(v149-int32(97)) < base.Ui32(int32(26)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v163 = base.I32_extend8_s(v162)
	if base.Ui32(int32(25)) < base.Ui32(v163-int32(65)) {
		v171 = v146
		goto L59
	} else {
		goto L65
	}
L62:
	;
	v162 = v149 & int32(95)
	goto L64
L63:
	;
	v162 = v149
	goto L64
L64:
	;
	goto L61
L65:
	;
	v170 = int32(*(*int8)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_metaphone[0]))))
	v171 = v170
	goto L59
L66:
	;
	v177 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v177)
	v188 = v85 + int32(2)
	v191 = int32(1)
	goto L34
L67:
	;
	v951 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v943+v36))) = uint8(v951)
	v957 = v36
	goto L5
L68:
	;
	if base.B2i32(v200 == int32(0))|base.B2i32(base.Ui32(v26) <= base.Ui32(v191)) != 0 {
		v943 = v191
		goto L67
	} else {
		goto L72
	}
L69:
	;
	v200 = v193 & int32(95)
	goto L71
L70:
	;
	v200 = v193
	goto L71
L71:
	;
	goto L68
L72:
	;
	v205 = v188
	v206 = v192
	v207 = v200
	v208 = v191
	goto L73
L73:
	;
	if base.Ui32(int32(25)) < base.Ui32(v207&int32(223)-int32(65)) {
		v920 = v205
		v923 = v208
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v943 = v923
	goto L67
L75:
	;
	v927 = v920 + int32(1)
	v928 = v16 + v927
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928))))
	if base.Ui32(v929-int32(97)) < base.Ui32(int32(26)) {
		goto L340
	} else {
		goto L341
	}
L76:
	;
	v222 = base.B2i32(v205 <= int32(0))
	if v205 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L77:
	;
	v920 = v205 + v918
	v923 = v917
	goto L75
L78:
	;
	v917 = v208 + int32(1)
	v918 = int32(0)
	goto L77
L79:
	;
	v908 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v908)
	goto L78
L80:
	;
	v904 = int32(1)
	v917 = v208 + v904
	v918 = v904
	goto L77
L81:
	;
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v797-int32(97)) < base.Ui32(int32(26)) {
		goto L299
	} else {
		goto L300
	}
L82:
	;
	v236 = int32(0)
	switch v207 - int32(66) {
	case 0:
		goto L105
	case 1:
		goto L81
	case 2:
		goto L104
	default:
		v917 = v208
		v918 = v236
		goto L77
	case 4, 8, 10, 11, 12, 16:
		goto L91
	case 5:
		goto L103
	case 6:
		goto L102
	case 9:
		goto L101
	case 14:
		goto L100
	case 15:
		goto L99
	case 17:
		goto L98
	case 18:
		goto L97
	case 20:
		goto L96
	case 21:
		goto L95
	case 22:
		goto L94
	case 23:
		goto L93
	case 24:
		goto L92
	}
L83:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v225-int32(97)) < base.Ui32(int32(26)) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	if v232 != v207 {
		goto L82
	} else {
		goto L88
	}
L85:
	;
	v232 = v225 & int32(95)
	goto L87
L86:
	;
	v232 = v225
	goto L87
L87:
	;
	goto L84
L88:
	;
	if v207 != int32(67) {
		v920 = v205
		v923 = v208
		goto L75
	} else {
		goto L89
	}
L89:
	;
	goto L81
L90:
	;
	v917 = v208 + int32(1)
	v918 = v236
	goto L77
L91:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v207)
	goto L90
L92:
	;
	v788 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v788)
	goto L90
L93:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v749-int32(97)) < base.Ui32(int32(26)) {
		goto L287
	} else {
		goto L288
	}
L94:
	;
	v739 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v739)
	v742 = v208 + int32(1)
	if v26 <= v742 {
		goto L283
	} else {
		goto L284
	}
L95:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v700-int32(97)) < base.Ui32(int32(26)) {
		goto L273
	} else {
		goto L274
	}
L96:
	;
	v698 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v698)
	goto L90
L97:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v660-int32(97)) < base.Ui32(int32(26)) {
		goto L264
	} else {
		goto L265
	}
L98:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v575-int32(97)) < base.Ui32(int32(26)) {
		goto L235
	} else {
		goto L236
	}
L99:
	;
	v573 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v573)
	goto L90
L100:
	;
	v555 = v208 + v36
	v557 = v208 + int32(1)
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v558-int32(97)) < base.Ui32(int32(26)) {
		goto L224
	} else {
		goto L225
	}
L101:
	;
	if v222 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L102:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v464-int32(97)) < base.Ui32(int32(26)) {
		goto L192
	} else {
		goto L193
	}
L103:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v308-int32(97)) < base.Ui32(int32(26)) {
		goto L135
	} else {
		goto L136
	}
L104:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if base.Ui32(v256-int32(97)) < base.Ui32(int32(26)) {
		goto L116
	} else {
		goto L117
	}
L105:
	;
	if v222 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v243-int32(97)) < base.Ui32(int32(26)) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v254 = int32(66)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v254)
	goto L90
L109:
	;
	if v250 == int32(77) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L113
	}
L110:
	;
	v250 = v243 & int32(95)
	goto L112
L111:
	;
	v250 = v243
	goto L112
L112:
	;
	goto L109
L113:
	;
	goto L108
L114:
	;
	v306 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v306)
	goto L90
L115:
	;
	if v263 != int32(71) {
		goto L114
	} else {
		goto L119
	}
L116:
	;
	v263 = v256 & int32(95)
	goto L118
L117:
	;
	v263 = v256
	goto L118
L118:
	;
	goto L115
L119:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v266-int32(97)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v275 = v273 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v275|int32(32)-int32(97)) {
		goto L114
	} else {
		goto L124
	}
L121:
	;
	v273 = v266 & int32(95)
	goto L123
L122:
	;
	v273 = v266
	goto L123
L123:
	;
	goto L120
L124:
	;
	if base.Ui32(v275-int32(97)) < base.Ui32(int32(26)) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v289 = base.I32_extend8_s(v288)
	if base.Ui32(int32(25)) < base.Ui32(v289-int32(65)) {
		goto L114
	} else {
		goto L129
	}
L126:
	;
	v288 = v275 & int32(95)
	goto L128
L127:
	;
	v288 = v275
	goto L128
L128:
	;
	goto L125
L129:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+uint32(_c_F_metaphone[0]))))
	if v296&int32(8) == int32(0) {
		goto L114
	} else {
		goto L130
	}
L130:
	;
	v302 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v302)
	goto L80
L131:
	;
	v417 = v315 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v417|int32(32)-int32(97)) {
		goto L175
	} else {
		goto L176
	}
L132:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v374-int32(97)) < base.Ui32(int32(26)) {
		goto L159
	} else {
		goto L160
	}
L133:
	;
	if v205 < int32(3) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	switch v315 - int32(72) {
	case 0:
		goto L133
	default:
		goto L131
	case 6:
		goto L132
	}
L135:
	;
	v315 = v308 & int32(95)
	goto L137
L136:
	;
	v315 = v308
	goto L137
L137:
	;
	goto L134
L138:
	;
	v372 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v372)
	goto L80
L139:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(3)))))
	if base.Ui32(v322-int32(97)) < base.Ui32(int32(26)) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	if v205 == int32(3) {
		goto L138
	} else {
		goto L152
	}
L141:
	;
	v331 = v329 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v331|int32(32)-int32(97)) {
		goto L140
	} else {
		goto L145
	}
L142:
	;
	v329 = v322 & int32(95)
	goto L144
L143:
	;
	v329 = v322
	goto L144
L144:
	;
	goto L141
L145:
	;
	if base.Ui32(v331-int32(97)) < base.Ui32(int32(26)) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v345 = base.I32_extend8_s(v344)
	if base.Ui32(int32(25)) < base.Ui32(v345-int32(65)) {
		goto L140
	} else {
		goto L150
	}
L147:
	;
	v344 = v331 & int32(95)
	goto L149
L148:
	;
	v344 = v331
	goto L149
L149:
	;
	goto L146
L150:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+uint32(_c_F_metaphone[0]))))
	if v352&int32(16) != 0 {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L151
	}
L151:
	;
	goto L140
L152:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(4)))))
	if base.Ui32(v360-int32(97)) < base.Ui32(int32(26)) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	if v367 == int32(72) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L157
	}
L154:
	;
	v367 = v360 & int32(95)
	goto L156
L155:
	;
	v367 = v360
	goto L156
L156:
	;
	goto L153
L157:
	;
	goto L138
L158:
	;
	if base.Ui32(int32(25)) < base.Ui32(v381&int32(223)-int32(65)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L162
	}
L159:
	;
	v381 = v374 & int32(95)
	goto L161
L160:
	;
	v381 = v374
	goto L161
L161:
	;
	goto L158
L162:
	;
	if v381 == int32(69) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v390 == int32(0) {
		v400 = v390
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	v414 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v414)
	goto L90
L166:
	;
	v402 = v400 & int32(255)
	if base.Ui32(v402-int32(97)) < base.Ui32(int32(26)) {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v393 == int32(0) {
		v400 = v393
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if v396 == int32(0) {
		v400 = v396
		goto L166
	} else {
		goto L169
	}
L169:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
	v400 = v399
	goto L166
L170:
	;
	if v409 == int32(68) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L174
	}
L171:
	;
	v409 = v402 & int32(95)
	goto L173
L172:
	;
	v409 = v402
	goto L173
L173:
	;
	goto L170
L174:
	;
	goto L165
L175:
	;
	v462 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v462)
	goto L90
L176:
	;
	if base.Ui32(v417-int32(97)) < base.Ui32(int32(26)) {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	v431 = base.I32_extend8_s(v430)
	if base.Ui32(int32(25)) < base.Ui32(v431-int32(65)) {
		goto L175
	} else {
		goto L181
	}
L178:
	;
	v430 = v417 & int32(95)
	goto L180
L179:
	;
	v430 = v417
	goto L180
L180:
	;
	goto L177
L181:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+uint32(_c_F_metaphone[0]))))
	if v438&int32(8) == int32(0) {
		goto L175
	} else {
		goto L182
	}
L182:
	;
	if v222 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v447-int32(97)) < base.Ui32(int32(26)) {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	v458 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v458)
	goto L90
L186:
	;
	if v454 == int32(71) {
		goto L175
	} else {
		goto L190
	}
L187:
	;
	v454 = v447 & int32(95)
	goto L189
L188:
	;
	v454 = v447
	goto L189
L189:
	;
	goto L186
L190:
	;
	goto L185
L191:
	;
	v473 = v471 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v473|int32(32)-int32(97)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L195
	}
L192:
	;
	v471 = v464 & int32(95)
	goto L194
L193:
	;
	v471 = v464
	goto L194
L194:
	;
	goto L191
L195:
	;
	if base.Ui32(v473-int32(97)) < base.Ui32(int32(26)) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v487 = base.I32_extend8_s(v486)
	if base.Ui32(int32(25)) < base.Ui32(v487-int32(65)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L200
	}
L197:
	;
	v486 = v473 & int32(95)
	goto L199
L198:
	;
	v486 = v473
	goto L199
L199:
	;
	goto L196
L200:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+uint32(_c_F_metaphone[0]))))
	if v494&int32(1) == int32(0) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L201
	}
L201:
	;
	if v205 <= int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v536 = int32(72)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v536)
	goto L90
L203:
	;
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v501-int32(97)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v510 = v508 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v510|int32(32)-int32(97)) {
		goto L202
	} else {
		goto L208
	}
L205:
	;
	v508 = v501 & int32(95)
	goto L207
L206:
	;
	v508 = v501
	goto L207
L207:
	;
	goto L204
L208:
	;
	if base.Ui32(v510-int32(97)) < base.Ui32(int32(26)) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	v524 = base.I32_extend8_s(v523)
	if base.Ui32(int32(25)) < base.Ui32(v524-int32(65)) {
		goto L202
	} else {
		goto L213
	}
L210:
	;
	v523 = v510 & int32(95)
	goto L212
L211:
	;
	v523 = v510
	goto L212
L212:
	;
	goto L209
L213:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524)+uint32(_c_F_metaphone[0]))))
	if v531&int32(4) != 0 {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L214
	}
L214:
	;
	goto L202
L215:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v542-int32(97)) < base.Ui32(int32(26)) {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	goto L217
L217:
	;
	v553 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v553)
	goto L90
L218:
	;
	if v549 == int32(67) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L222
	}
L219:
	;
	v549 = v542 & int32(95)
	goto L221
L220:
	;
	v549 = v542
	goto L221
L221:
	;
	goto L218
L222:
	;
	goto L217
L223:
	;
	if v565 == int32(72) {
		goto L227
	} else {
		goto L228
	}
L224:
	;
	v565 = v558 & int32(95)
	goto L226
L225:
	;
	v565 = v558
	goto L226
L226:
	;
	goto L223
L227:
	;
	v568 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v555))) = uint8(v568)
	v917 = v557
	v918 = v236
	goto L77
L228:
	;
	goto L229
L229:
	;
	v570 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v555))) = uint8(v570)
	v917 = v557
	v918 = v236
	goto L77
L230:
	;
	v658 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v658)
	goto L90
L231:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v608 == int32(0) {
		v615 = v608
		goto L243
	} else {
		goto L244
	}
L232:
	;
	v606 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v606)
	goto L80
L233:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v585-int32(97)) < base.Ui32(int32(26)) {
		goto L239
	} else {
		goto L240
	}
L234:
	;
	switch v582 - int32(67) {
	case 0:
		goto L231
	default:
		goto L230
	case 5:
		goto L232
	case 6:
		goto L233
	}
L235:
	;
	v582 = v575 & int32(95)
	goto L237
L236:
	;
	v582 = v575
	goto L237
L237:
	;
	goto L234
L238:
	;
	v594 = v592 - int32(65)
	v595 = int32(0)
	if base.B2i32(v594 == v595)|base.B2i32(v594 == int32(14)) == v595 {
		goto L230
	} else {
		goto L242
	}
L239:
	;
	v592 = v585 & int32(95)
	goto L241
L240:
	;
	v592 = v585
	goto L241
L241:
	;
	goto L238
L242:
	;
	v603 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v603)
	goto L90
L243:
	;
	v617 = v615 & int32(255)
	if base.Ui32(v617-int32(97)) < base.Ui32(int32(26)) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v611 == int32(0) {
		v615 = v611
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	v615 = v614
	goto L243
L246:
	;
	if v624 != int32(72) {
		goto L230
	} else {
		goto L250
	}
L247:
	;
	v624 = v617 & int32(95)
	goto L249
L248:
	;
	v624 = v617
	goto L249
L249:
	;
	goto L246
L250:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v627 == int32(0) {
		v637 = v627
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v639 = v637 & int32(255)
	if base.Ui32(v639-int32(97)) < base.Ui32(int32(26)) {
		goto L256
	} else {
		goto L257
	}
L252:
	;
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
	if v630 == int32(0) {
		v637 = v630
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if v633 == int32(0) {
		v637 = v633
		goto L251
	} else {
		goto L254
	}
L254:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
	v637 = v636
	goto L251
L255:
	;
	if v646 != int32(87) {
		goto L230
	} else {
		goto L259
	}
L256:
	;
	v646 = v639 & int32(95)
	goto L258
L257:
	;
	v646 = v639
	goto L258
L258:
	;
	goto L255
L259:
	;
	v650 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v650)
	v917 = v208 + int32(1)
	v918 = int32(2)
	goto L77
L260:
	;
	v695 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v695)
	goto L90
L261:
	;
	v691 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v691)
	goto L80
L262:
	;
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v670-int32(97)) < base.Ui32(int32(26)) {
		goto L268
	} else {
		goto L269
	}
L263:
	;
	switch v667 - int32(72) {
	case 0:
		goto L261
	case 1:
		goto L262
	default:
		goto L260
	}
L264:
	;
	v667 = v660 & int32(95)
	goto L266
L265:
	;
	v667 = v660
	goto L266
L266:
	;
	goto L263
L267:
	;
	v679 = v677 - int32(65)
	v680 = int32(0)
	if base.B2i32(v679 == v680)|base.B2i32(v679 == int32(14)) == v680 {
		goto L260
	} else {
		goto L271
	}
L268:
	;
	v677 = v670 & int32(95)
	goto L270
L269:
	;
	v677 = v670
	goto L270
L270:
	;
	goto L267
L271:
	;
	v688 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v688)
	goto L90
L272:
	;
	v709 = v707 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v709|int32(32)-int32(97)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L276
	}
L273:
	;
	v707 = v700 & int32(95)
	goto L275
L274:
	;
	v707 = v700
	goto L275
L275:
	;
	goto L272
L276:
	;
	if base.Ui32(v709-int32(97)) < base.Ui32(int32(26)) {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v723 = base.I32_extend8_s(v722)
	if base.Ui32(int32(25)) < base.Ui32(v723-int32(65)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L281
	}
L278:
	;
	v722 = v709 & int32(95)
	goto L280
L279:
	;
	v722 = v709
	goto L280
L280:
	;
	goto L277
L281:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v723)+uint32(_c_F_metaphone[0]))))
	if v730&int32(1) == int32(0) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L282
	}
L282:
	;
	v736 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v736)
	goto L90
L283:
	;
	v917 = v742
	v918 = v236
	goto L77
L284:
	;
	goto L285
L285:
	;
	v745 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v742+v36))) = uint8(v745)
	v917 = v208 + int32(2)
	v918 = v236
	goto L77
L286:
	;
	v758 = v756 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v758|int32(32)-int32(97)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L290
	}
L287:
	;
	v756 = v749 & int32(95)
	goto L289
L288:
	;
	v756 = v749
	goto L289
L289:
	;
	goto L286
L290:
	;
	if base.Ui32(v758-int32(97)) < base.Ui32(int32(26)) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v772 = base.I32_extend8_s(v771)
	if base.Ui32(int32(25)) < base.Ui32(v772-int32(65)) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L295
	}
L292:
	;
	v771 = v758 & int32(95)
	goto L294
L293:
	;
	v771 = v758
	goto L294
L294:
	;
	goto L291
L295:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772)+uint32(_c_F_metaphone[0]))))
	if v779&int32(1) == int32(0) {
		v917 = v208
		v918 = v236
		goto L77
	} else {
		goto L296
	}
L296:
	;
	v785 = int32(89)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v785)
	goto L90
L297:
	;
	if v804 != int32(72) {
		goto L79
	} else {
		goto L324
	}
L298:
	;
	v806 = v804 & int32(255)
	if base.Ui32(int32(25)) < base.Ui32(v806|int32(32)-int32(97)) {
		goto L297
	} else {
		goto L302
	}
L299:
	;
	v804 = v797 & int32(95)
	goto L301
L300:
	;
	v804 = v797
	goto L301
L301:
	;
	goto L298
L302:
	;
	if base.Ui32(v806-int32(97)) < base.Ui32(int32(26)) {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	v820 = base.I32_extend8_s(v819)
	if base.Ui32(int32(25)) < base.Ui32(v820-int32(65)) {
		goto L297
	} else {
		goto L307
	}
L304:
	;
	v819 = v806 & int32(95)
	goto L306
L305:
	;
	v819 = v806
	goto L306
L306:
	;
	goto L303
L307:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+uint32(_c_F_metaphone[0]))))
	if v827&int32(8) == int32(0) {
		goto L297
	} else {
		goto L308
	}
L308:
	;
	if v804 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	if v222 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L310:
	;
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v836-int32(97)) < base.Ui32(int32(26)) {
		goto L312
	} else {
		goto L313
	}
L311:
	;
	if base.B2i32(v804 != int32(73))|base.B2i32(v843 != int32(65)) != 0 {
		goto L309
	} else {
		goto L315
	}
L312:
	;
	v843 = v836 & int32(95)
	goto L314
L313:
	;
	v843 = v836
	goto L314
L314:
	;
	goto L311
L315:
	;
	v848 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v848)
	goto L78
L316:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v855-int32(97)) < base.Ui32(int32(26)) {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	goto L318
L318:
	;
	v867 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v867)
	goto L78
L319:
	;
	if v862 == int32(83) {
		v917 = v208
		v918 = int32(0)
		goto L77
	} else {
		goto L323
	}
L320:
	;
	v862 = v855 & int32(95)
	goto L322
L321:
	;
	v862 = v855
	goto L322
L322:
	;
	goto L319
L323:
	;
	goto L318
L324:
	;
	v872 = int32(75)
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
	if base.Ui32(v873-int32(97)) < base.Ui32(int32(26)) {
		goto L327
	} else {
		goto L328
	}
L325:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v208+v36))) = uint8(v898)
	goto L80
L326:
	;
	if v880 == int32(82) {
		v898 = v872
		goto L325
	} else {
		goto L330
	}
L327:
	;
	v880 = v873 & int32(95)
	goto L329
L328:
	;
	v880 = v873
	goto L329
L329:
	;
	goto L326
L330:
	;
	if v222 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206-int32(1)))))
	if base.Ui32(v887-int32(97)) < base.Ui32(int32(26)) {
		goto L335
	} else {
		goto L336
	}
L332:
	;
	goto L333
L333:
	;
	v898 = int32(88)
	goto L325
L334:
	;
	if v894 == int32(83) {
		v898 = v872
		goto L325
	} else {
		goto L338
	}
L335:
	;
	v894 = v887 & int32(95)
	goto L337
L336:
	;
	v894 = v887
	goto L337
L337:
	;
	goto L334
L338:
	;
	goto L333
L339:
	;
	if v936 == int32(0) {
		v943 = v923
		goto L67
	} else {
		goto L343
	}
L340:
	;
	v936 = v929 & int32(95)
	goto L342
L341:
	;
	v936 = v929
	goto L342
L342:
	;
	goto L339
L343:
	;
	if v923 < v26 {
		v205 = v927
		v206 = v928
		v207 = v936
		v208 = v923
		goto L73
	} else {
		goto L344
	}
L344:
	;
	goto L74
L345:
	;
	m.G0 = v13 + int32(32)
	return v963
L346:
	;
	F_errmsg_internal(m, int32(_a_F_metaphone_1), int32(0))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L6
	} else {
		goto L347
	}
L347:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(368), int32(_a_F_metaphone_3))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L6
	} else {
		goto L348
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L6
	} else {
		goto L350
	}
L350:
	;
	F_errmsg(m, int32(_a_F_metaphone_4), int32(0))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
		goto L6
	} else {
		goto L351
	}
L351:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(287), int32(_a_F_metaphone_5))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L6
	} else {
		goto L352
	}
L352:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L353:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L6
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(255)
	F_errmsg(m, int32(_a_F_metaphone_6), v13+int32(16))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L6
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(282), int32(_a_F_metaphone_5))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L6
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L6
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(255)
	F_errmsg(m, int32(_a_F_metaphone_7), v13)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L6
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(275), int32(_a_F_metaphone_5))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L6
	} else {
		goto L360
	}
L360:
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 != int32(2) {
		v64 = l1
		v66 = l3
		v67 = v10
		v72 = l2
		if int32(0) < v72 {
			v97 = v64
			v99 = v66
			v100 = v67
			for {
				v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
				*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(84))))) = uint8(v106)
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(52))))))
				v110 = v106 ^ v109
				*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v110)
				v112 = int32(1)
				v117 = v100 + v112
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v119 = v118 + v72
				if v117 < v119 {
					v97 = v97 + v112
					v99 = v99 + v112
					v100 = v117
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
			return v72
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
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
			v79 = l0 + int32(20)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v82 = v80 - int32(2)
			if v82 != 0 {
				base.MemoryCopy(m, v79, l0+int32(86), v82)
			} else {
			}
			v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+84)))
			*(*uint16)(unsafe.Add(mBase, uint32(v82+v79))) = uint16(v87)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			return v17
		} else {
			v64 = v51
			v66 = v53
			v67 = v54
			v72 = l2 - v17
			if int32(0) < v72 {
				v97 = v64
				v99 = v66
				v100 = v67
				for {
					v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
					*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(84))))) = uint8(v106)
					v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(52))))))
					v110 = v106 ^ v109
					*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v110)
					v112 = int32(1)
					v117 = v100 + v112
					v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v119 = v118 + v72
					if v117 < v119 {
						v97 = v97 + v112
						v99 = v99 + v112
						v100 = v117
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v119
				return v72
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v67 + v72
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
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v460 int32
	_ = v460
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	if l2 <= l1 {
		v460 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(32)
	return v460
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if l1+int32(1) != l2 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v119 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	v28 = l2 - l1
	v37 = v5
	v39 = v5
	v40 = l1
	v47 = v5
	goto L7
L5:
	;
	v87 = v5
	v89 = v5
	v90 = l1
	goto L6
L6:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v24+v90<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v105 <= l3 {
		v119 = v89
		goto L3
	} else {
		goto L17
	}
L7:
	;
	v53 = v24 + v40<<(uint(int32(2))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if l3 < v55 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v28&int32(1) == int32(0) {
		v119 = v75
		goto L3
	} else {
		goto L16
	}
L9:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v54)+8)))
	v63 = v60
	v64 = v39 + base.B2i32(v37&int32(255) != v60)
	goto L11
L10:
	;
	v63 = v37
	v64 = v39
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if l3 < v66 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v65)+8)))
	v74 = v71
	v75 = v64 + base.B2i32(v63&int32(255) != v71)
	goto L14
L13:
	;
	v74 = v63
	v75 = v64
	goto L14
L14:
	;
	v76 = int32(2)
	v77 = v40 + v76
	v79 = v47 + v76
	if v79 != v28&int32(-2) {
		v37 = v74
		v39 = v75
		v40 = v77
		v47 = v79
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
L16:
	;
	v87 = v74
	v89 = v75
	v90 = v77
	goto L6
L17:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v104)+8)))
	v119 = v89 + base.B2i32(v108 != v87&int32(255))
	goto L3
L18:
	;
	v460 = int32(0)
	goto L1
L19:
	;
	goto L20
L20:
	;
	v135 = v119 << (uint(int32(3)) % 32)
	if base.Ui32(int32(1024)) <= base.Ui32(v135) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v119
	v166 = l3 + int32(1)
	v171 = l1
	v176 = v163 + int32(4)
	v177 = l1
	v182 = int32(0)
	goto L32
L22:
	;
	v140 = F_palloc0(m, v135|int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v147 = (v135 + int32(11)) & int32(2040)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v147) <= base.Ui32(v148) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	return int32(0)
L26:
	;
	v163 = v140
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v155 - v147
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v147 + v156
	v163 = v156
	goto L21
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v155 = v148
	v156 = v150
	goto L27
L29:
	;
	goto L30
L30:
	;
	v151 = int32(_a_F_mkSPNode_0)
	v153 = F_palloc0(m, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v155 = v151
	v156 = v153
	goto L27
L32:
	;
	v189 = v177 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v193 <= l3 {
		v434 = v171
		v437 = v176
		v440 = v182
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v447 = F_mkSPNode(m, l0, v434, l2, v166)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L25
	} else {
		goto L89
	}
L34:
	;
	v445 = v177 + int32(1)
	if v445 != l2 {
		v171 = v434
		v176 = v437
		v177 = v445
		v182 = v440
		goto L32
	} else {
		goto L88
	}
L35:
	;
	v196 = v182 & int32(255)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v192)+8)))
	if v196 == v198 {
		v212 = v171
		v213 = v176
		v214 = v182
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v220 = v215&int32(-256) | v214&int32(255)
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222+v189)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v225 != v166 {
		v434 = v212
		v437 = v213
		v440 = v214
		goto L34
	} else {
		goto L42
	}
L37:
	;
	if v196 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v212 = v171
	v213 = v176
	v214 = v198
	goto L36
L39:
	;
	goto L40
L40:
	;
	v202 = F_mkSPNode(m, l0, v171, v177, v166)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v202
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207+v189)))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209+l3)+8)))
	v212 = v177
	v213 = v176 + int32(8)
	v214 = v211
	goto L36
L42:
	;
	v229 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if base.B2i32(v215&int32(256) == v229)|base.B2i32(v231 == int32(base.Ui32(v215)>>(uint(int32(13))%32))) == v229 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v400 | int32(256)
	v413 = F_makeCompoundFlags(m, l0, int32(base.Ui32(v400)>>(uint(int32(13))%32)))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L25
	} else {
		goto L83
	}
L44:
	;
	v238 = F_makeCompoundFlags(m, l0, v231)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L25
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v400 = v220&int32(_a_F_mkSPNode_1) | v231<<(uint(int32(13))%32)
	v407 = int32(0)
	goto L43
L47:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v243+v189)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v250 = int32(base.Ui32(v248) >> (uint(int32(13)) % 32))
	v252 = v250 << (uint(int32(2)) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v247+v252)))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v255 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v400 = v385&int32(_a_F_mkSPNode_1) | v382<<(uint(int32(13))%32)
	v407 = v238&int32(base.Ui32(v215)>>(uint(int32(9))%32)) ^ int32(1)
	goto L43
L49:
	;
	v382 = v246
	v385 = v248
	goto L48
L50:
	;
	goto L51
L51:
	;
	v259 = v246 << (uint(int32(2)) % 32)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v247+v259)))
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	if v262 == int32(0) {
		v382 = v250
		v385 = v248
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v265 <= v266+int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v265 << (uint(int32(1)) % 32)
	v275 = F_repalloc(m, v247, v265<<(uint(int32(3))%32))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L25
	} else {
		goto L56
	}
L54:
	;
	v283 = v247
	v284 = v254
	v285 = v261
	v286 = v266
	goto L55
L55:
	;
	v287 = int32(2)
	v289 = v286<<(uint(v287)%32) + v283
	v290 = F_strlen(m, v284)
	mBase = m.M
	v291 = F_strlen(m, v285)
	mBase = m.M
	v292 = v290 + v291
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v293 == v287 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v275
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275+v259)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275+v252)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v283 = v275
	v284 = v281
	v285 = v279
	v286 = v282
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = v370
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v376 + int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v382 = v376
	v385 = v380
	goto L48
L58:
	;
	v297 = v292 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v297) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v333 = v292 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v333) {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v322+v252)))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322+v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v324
	v330 = F_pg_sprintf(m, v319, int32(_a_F_mkSPNode_2), v21)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L25
	} else {
		goto L71
	}
L62:
	;
	v300 = F_palloc0(m, v297)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L25
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v305 = (v292 + int32(9)) & int32(4088)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v305) <= base.Ui32(v306) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v319 = v300
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v313 - v305
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v314 + v305
	v319 = v314
	goto L61
L67:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v313 = v306
	v314 = v308
	goto L66
L68:
	;
	goto L69
L69:
	;
	v309 = int32(_a_F_mkSPNode_0)
	v311 = F_palloc0(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	v313 = v309
	v314 = v311
	goto L66
L71:
	;
	v370 = v319
	goto L57
L72:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v358+v252)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358+v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v360
	v368 = F_pg_sprintf(m, v355, int32(_a_F_mkSPNode_3), v21+int32(16))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L25
	} else {
		goto L82
	}
L73:
	;
	v336 = F_palloc0(m, v333)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L25
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v341 = (v292 + int32(8)) & int32(4088)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v341) <= base.Ui32(v342) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v355 = v336
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v349 - v341
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v350 + v341
	v355 = v350
	goto L72
L78:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v349 = v342
	v350 = v344
	goto L77
L79:
	;
	goto L80
L80:
	;
	v345 = int32(_a_F_mkSPNode_0)
	v347 = F_palloc0(m, v345)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v349 = v345
	v350 = v347
	goto L77
L82:
	;
	v370 = v355
	goto L57
L83:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	if v413 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v425 = v415 | int32(_a_F_mkSPNode_4)
	goto L86
L85:
	;
	v425 = v415&int32(-7681) | v413<<(uint(int32(9))%32)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v425
	if v407&int32(1) == int32(0) {
		v434 = v212
		v437 = v213
		v440 = v214
		goto L34
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v425 & int32(-513)
	v434 = v212
	v437 = v213
	v440 = v214
	goto L34
L88:
	;
	goto L33
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v447
	v460 = v163
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
	var v31 int32
	_ = v31
	v7 = F_get_typisdefined(m, l0)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v31 = int32(1)
			return v31
		} else {
			v13 = int32(0)
			v14 = F_get_element_type(m, l0)
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v31 = v13
					return v31
				} else {
					v18 = F_get_array_type(m, v14)
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						if v18 != l0 {
							v31 = v13
							return v31
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
											v31 = int32(1)
											return v31
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
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
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
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
	v104 = int32(32)
	if base.B2i32(int32(4) <= v101)&(base.B2i32(v104 < v8)|base.B2i32(base.Ui32(v104) < base.Ui32(v101))) == int32(0) {
		goto L36
	} else {
		goto L37
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
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v32 < int32(0) {
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
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v37 = v35 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v37))|base.B2i32(int32(1)<<(uint(v37)%32)&int32(_a_F_moveins_0) == int32(0)) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v47 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v48 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v52+v32*int32(24))+12)) = v56
	v60 = v56
	goto L15
L17:
	;
	goto L18
L18:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v58
	v60 = v58
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = v48
	goto L21
L20:
	;
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
	goto L11
L22:
	;
	if v66 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v66
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67)+16)) = v66
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+20)) = v67
	goto L28
L27:
	;
	goto L28
L28:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v73 - int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v78 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v77 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v77
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v77
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+28)) = v78
	goto L35
L34:
	;
	goto L35
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v84 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(0)
	v91 = v18 + int32(8)
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+16)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91)+8)) = v92
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = v92
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v18
	goto L10
L36:
	;
	goto L39
L37:
	;
	goto L38
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_moveins[0]))
	if v201 != 0 {
		goto L69
	} else {
		goto L70
	}
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	F_cparc(m, l0, v119, v122, l2)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
	v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v119)+4)))
	if v131 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L39
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v119)+20))
	if v166 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v136 = v134 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v136))|base.B2i32(int32(1)<<(uint(v136)%32)&int32(_a_F_moveins_0) == int32(0)) != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v146 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v119)+36))
	if v147 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v159 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v151+v131*int32(24))+12)) = v155
	v159 = v155
	goto L48
L50:
	;
	goto L51
L51:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v119)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+32)) = v157
	v159 = v157
	goto L48
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+36)) = v147
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119)+32)) = int64(0)
	goto L44
L55:
	;
	if v165 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+20)) = v165
	goto L55
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v166)+16)) = v165
	goto L55
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v166
	goto L61
L60:
	;
	goto L61
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v172 - int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if v177 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v176 != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v176
	goto L62
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177)+24)) = v176
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+28)) = v177
	goto L68
L67:
	;
	goto L68
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v183 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(0)
	v190 = v119 + int32(8)
	v191 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+16)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190)+8)) = v191
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = v191
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v119)+16)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v119
	goto L43
L69:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_sortins(m, l0, l1)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L8
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_sortins(m, l0, l2)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	if v209 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v210 == int32(0) {
		v351 = v210
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v351 == int32(0) {
		goto L1
	} else {
		goto L128
	}
L77:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v213 == int32(0) {
		v351 = v210
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v217 = v210
	v220 = v213
	goto L79
L79:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v224 < v226 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v351 = v343
	goto L76
L81:
	;
	if v343 == int32(0) {
		v351 = v343
		goto L76
	} else {
		goto L126
	}
L82:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v220)+24))
	v343 = v217
	v345 = v342
	goto L81
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	if v318 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L84:
	;
	if v226 < v224 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217)+4)))
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+4)))
	if v229 < v230 {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	if v230 < v229 {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v233 < v234 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	if v234 < v233 {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v220)+24))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v245 = int32(*(*int16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v245 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v343 = v238
	v345 = v237
	goto L81
L91:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v217)+16))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	if v280 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v250 = v248 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v250))|base.B2i32(int32(1)<<(uint(v250)%32)&int32(_a_F_moveins_0) == int32(0)) != 0 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v260 != 0 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v217)+36))
	if v261 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	if v273 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v217)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v265+v245*int32(24))+12)) = v269
	v273 = v269
	goto L95
L97:
	;
	goto L98
L98:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v217)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+32)) = v271
	v273 = v271
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v273)+36)) = v261
	goto L101
L100:
	;
	goto L101
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v217)+32)) = int64(0)
	goto L91
L102:
	;
	if v279 != 0 {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244)+20)) = v279
	goto L102
L104:
	;
	goto L105
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v280)+16)) = v279
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = v280
	goto L108
L107:
	;
	goto L108
L108:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v244)+12)) = v286 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v217)+24))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v217)+28))
	if v291 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v290 != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v243)+16)) = v290
	goto L109
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+24)) = v290
	goto L109
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v291
	goto L115
L114:
	;
	goto L115
L115:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v243)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v243)+8)) = v297 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = int32(0)
	v304 = v217 + int32(8)
	v305 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v304)+16)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304)+8)) = v305
	*(*int64)(unsafe.Add(mBase, uint32(v304))) = v305
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+16)) = v311
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v217
	goto L90
L116:
	;
	if v317 != 0 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316)+16)) = v317
	goto L116
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+24)) = v317
	goto L116
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317)+28)) = v318
	goto L122
L121:
	;
	goto L122
L122:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v316)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v316)+8)) = v324 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = l2
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v217)+24)) = v329
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v333 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v333)+28)) = v217
	goto L125
L124:
	;
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v217
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v336 + int32(1)
	v343 = v317
	v345 = v220
	goto L81
L126:
	;
	if v345 != 0 {
		v217 = v343
		v220 = v345
		goto L79
	} else {
		goto L127
	}
L127:
	;
	goto L80
L128:
	;
	v360 = v351
	goto L129
L129:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v360)+24))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v360)+28))
	if v368 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L1
L131:
	;
	if v367 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+16)) = v367
	goto L131
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+24)) = v367
	goto L131
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+28)) = v368
	goto L137
L136:
	;
	goto L137
L137:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+8)) = v374 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+12)) = l2
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v360)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = v379
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v383 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+28)) = v360
	goto L140
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v360
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v386 + int32(1)
	if v367 != 0 {
		v360 = v367
		goto L129
	} else {
		goto L141
	}
L141:
	;
	goto L130
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1459), int32(0), v4, v5)
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
