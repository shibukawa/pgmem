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
	var v39 int64
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
	var v72 int64
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+24))
	v39 = m.T0[v38].(func(*base.Module, int32, int32, int32) int64)(m, v35, v18, v34+int32(25))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v39
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+25)))
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
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+37)))
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
	v67 = v64 + v58<<(uint(int32(6))%32)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v72 = m.T0[v71].(func(*base.Module, int32, int32, int32) int64)(m, v68, v18, v67+int32(25))
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
	*(*int64)(unsafe.Add(mBase, uint32(v67)+16)) = v72
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
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+25)))
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
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
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+132))
	v20 = F_ExecStoreMinimalTuple(m, v17, v18, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+197)))
	if v25 == v24 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(16)
	return v123
L4:
	;
	v28 = int32(_a_F_MemoizeHash_equal_0)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+120))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+6)))
	if v36 < v35 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v18
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v14)+140))
	if v102 == int32(0) {
		v123 = v24
		goto L3
	} else {
		goto L27
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	m.T0[v39].(func(*base.Module, int32, int32))(m, v18, v35)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+6)))
	if v44 < v43 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
	m.T0[v47].(func(*base.Module, int32, int32))(m, v15, v43)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v30 <= int32(0) {
		v93 = v24
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v29
	v123 = v93
	goto L3
L16:
	;
	v53 = int32(0)
	goto L17
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v53))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65+v53))))
	if v64 != v67 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v93 = v24
	goto L15
L19:
	;
	v93 = int32(0)
	goto L15
L20:
	;
	goto L21
L21:
	;
	if v64 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = v53 + int32(1)
	if v87 != v30 {
		v53 = v87
		goto L17
	} else {
		goto L26
	}
L23:
	;
	v71 = v53 << (uint(int32(3)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v71+v72)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v75+v71)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v79 = v78 + v71
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+32)))
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+30)))
	v82 = F_datum_image_eq(m, v74, v77, v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v82 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v93 = int32(0)
	goto L15
L26:
	;
	goto L18
L27:
	;
	v105 = int32(_a_F_MemoizeHash_equal_0)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v108
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+24))
	v113 = m.T0[v112].(func(*base.Module, int32, int32, int32) int64)(m, v102, v16, v12+int32(15))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_equal[0])) = v106
	v123 = base.B2i32(v113 != int64(0))
	goto L3
}
func F_MemoizeHash_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v64 int64
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
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
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
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
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
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
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
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
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
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
	var v947 int32
	_ = v947
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
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
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
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
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
	var v1049 int32
	_ = v1049
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
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1221 int64
	_ = v1221
	var v1222 int64
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	v2 = int32(0)
	v12 = int32(_a_F_MemoizeHash_hash_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+120))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+136))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0])) = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+197)))
	if v21 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MemoizeHash_hash[0])) = v13
	v1243 = int32(16)
	v1247 = (int32(base.Ui32(v1233)>>(uint(v1243)%32)) ^ v1233) * int32(-2048144789)
	v1252 = (int32(base.Ui32(v1247)>>(uint(int32(13))%32)) ^ v1247) * int32(-1028477387)
	return int32(base.Ui32(v1252)>>(uint(v1243)%32)) ^ v1252
L2:
	;
	if v15 <= int32(0) {
		v1233 = v2
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	if v15 <= int32(0) {
		v1233 = v2
		goto L1
	} else {
		goto L197
	}
L5:
	;
	v27 = int32(0)
	v30 = v2
	goto L6
L6:
	;
	v39 = base.I32_rotl(v30, int32(1))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v27))))
	if v42 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v1233 = v1179
	goto L1
L8:
	;
	v46 = v27 << (uint(int32(3)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v46+v47)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v51 = v50 + v46
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+32)))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+30)))
	v54 = m.G0
	v56 = v54 - int32(16)
	m.G0 = v56
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v49
	if v52 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v1179 = v39
	goto L10
L10:
	;
	v1185 = v27 + int32(1)
	if v1185 != v15 {
		v27 = v1185
		v30 = v1179
		goto L6
	} else {
		goto L196
	}
L11:
	;
	m.G0 = v56 + int32(16)
	v1179 = v1169 ^ v39
	goto L10
L12:
	;
	switch v53 - int32(1) {
	case 0:
		v64 = int64(56)
		goto L16
	case 1:
		goto L18
	default:
		goto L15
	case 3:
		goto L17
	}
L13:
	;
	goto L14
L14:
	;
	if int32(0) < v53 {
		goto L59
	} else {
		goto L60
	}
L15:
	;
	v70 = v56 + int32(8)
	v77 = int32(-1636608424)
	if v70&int32(3) != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v49 << (uint(v64) % 64) >> (uint(v64) % 64)
	goto L15
L17:
	;
	v64 = int64(32)
	goto L16
L18:
	;
	v64 = int64(48)
	goto L16
L19:
	;
	v1169 = v331 ^ v323 - base.I32_rotl(v331, int32(24))
	goto L11
L20:
	;
	v309 = int32(14)
	v311 = v305 ^ v306 - base.I32_rotl(v305, v309)
	v315 = v311 ^ v304 - base.I32_rotl(v311, int32(11))
	v319 = v315 ^ v305 - base.I32_rotl(v315, int32(25))
	v323 = v319 ^ v311 - base.I32_rotl(v319, int32(16))
	v327 = v323 ^ v315 - base.I32_rotl(v323, int32(4))
	v331 = v327 ^ v319 - base.I32_rotl(v327, v309)
	goto L19
L21:
	;
	switch int32(7) {
	case 0:
		v297 = v77
		v298 = v77
		v299 = v77
		goto L48
	case 1:
		v290 = v77
		v291 = v77
		v292 = v77
		goto L49
	case 2:
		v283 = v77
		v284 = v77
		v285 = v77
		goto L50
	case 3:
		v277 = v77
		v278 = v77
		goto L51
	case 4:
		v273 = v77
		v274 = v77
		goto L52
	case 5:
		v267 = v77
		v268 = v77
		goto L53
	case 6:
		v261 = v77
		v262 = v77
		goto L54
	case 7:
		v256 = v77
		goto L55
	case 8:
		v251 = v77
		goto L56
	case 9:
		v246 = v77
		goto L57
	case 10:
		goto L58
	default:
		v304 = v77
		v305 = v77
		v306 = v77
		goto L20
	}
L23:
	;
	goto L26
L24:
	;
	goto L25
L25:
	;
	goto L28
L26:
	;
	goto L21
L27:
	;
	switch int32(7) {
	case 0:
		v183 = v77
		goto L34
	case 1:
		v178 = v77
		goto L35
	case 2:
		goto L36
	case 3:
		v171 = v77
		goto L37
	case 4:
		v168 = v77
		goto L38
	case 5:
		v163 = v77
		goto L39
	case 6:
		goto L40
	case 7:
		v154 = v77
		goto L41
	case 8:
		v149 = v77
		goto L42
	case 9:
		v144 = v77
		goto L43
	case 10:
		goto L44
	default:
		v304 = v77
		v305 = v77
		v306 = v77
		goto L20
	}
L28:
	;
	goto L27
L34:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v304 = v183 + v184
	v305 = v77
	v306 = v77
	goto L20
L35:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v183 = v179<<(uint(int32(8))%32) + v178
	goto L34
L36:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	v178 = v174<<(uint(int32(16))%32) + v77
	goto L35
L37:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v304 = v172 + v77
	v305 = v171
	v306 = v77
	goto L20
L38:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	v171 = v168 + v169
	goto L37
L39:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+5)))
	v168 = v164<<(uint(int32(8))%32) + v163
	goto L38
L40:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+6)))
	v163 = v159<<(uint(int32(16))%32) + v77
	goto L39
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v304 = v155 + v77
	v305 = v157 + v77
	v306 = v154
	goto L20
L42:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
	v154 = v150<<(uint(int32(8))%32) + v149
	goto L41
L43:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
	v149 = v145<<(uint(int32(16))%32) + v144
	goto L42
L44:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	v144 = v140<<(uint(int32(24))%32) + v77
	goto L43
L48:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v304 = v297 + v300
	v305 = v298
	v306 = v299
	goto L20
L49:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v297 = v293<<(uint(int32(8))%32) + v290
	v298 = v291
	v299 = v292
	goto L48
L50:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+2)))
	v290 = v286<<(uint(int32(16))%32) + v283
	v291 = v284
	v292 = v285
	goto L49
L51:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+3)))
	v283 = v279<<(uint(int32(24))%32) + v77
	v284 = v277
	v285 = v278
	goto L50
L52:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)))
	v277 = v273 + v275
	v278 = v274
	goto L51
L53:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+5)))
	v273 = v269<<(uint(int32(8))%32) + v267
	v274 = v268
	goto L52
L54:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+6)))
	v267 = v263<<(uint(int32(16))%32) + v261
	v268 = v262
	goto L53
L55:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+7)))
	v261 = v257<<(uint(int32(24))%32) + v77
	v262 = v256
	goto L54
L56:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)))
	v256 = v252<<(uint(int32(8))%32) + v251
	goto L55
L57:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+9)))
	v251 = v247<<(uint(int32(16))%32) + v246
	goto L56
L58:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+10)))
	v246 = v242<<(uint(int32(24))%32) + v77
	goto L57
L59:
	;
	v338 = base.I32_wrap_i64(v49)
	v344 = v53 - int32(1636608432)
	if v338&int32(3) != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	goto L61
L61:
	;
	switch v53 + int32(2) {
	case 0:
		goto L102
	case 1:
		goto L104
	default:
		goto L103
	}
L62:
	;
	v1169 = v598 ^ v590 - base.I32_rotl(v598, int32(24))
	goto L11
L63:
	;
	v576 = int32(14)
	v578 = v572 ^ v573 - base.I32_rotl(v572, v576)
	v582 = v578 ^ v571 - base.I32_rotl(v578, int32(11))
	v586 = v582 ^ v572 - base.I32_rotl(v582, int32(25))
	v590 = v586 ^ v578 - base.I32_rotl(v586, int32(16))
	v594 = v590 ^ v582 - base.I32_rotl(v590, int32(4))
	v598 = v594 ^ v586 - base.I32_rotl(v594, v576)
	goto L62
L64:
	;
	switch v502 - int32(1) {
	case 0:
		v564 = v503
		v565 = v504
		v566 = v505
		goto L91
	case 1:
		v557 = v503
		v558 = v504
		v559 = v505
		goto L92
	case 2:
		v550 = v503
		v551 = v504
		v552 = v505
		goto L93
	case 3:
		v544 = v504
		v545 = v505
		goto L94
	case 4:
		v540 = v504
		v541 = v505
		goto L95
	case 5:
		v534 = v504
		v535 = v505
		goto L96
	case 6:
		v528 = v504
		v529 = v505
		goto L97
	case 7:
		v523 = v505
		goto L98
	case 8:
		v518 = v505
		goto L99
	case 9:
		v513 = v505
		goto L100
	case 10:
		goto L101
	default:
		v571 = v503
		v572 = v504
		v573 = v505
		goto L63
	}
L65:
	;
	v453 = v338
	v454 = v53
	v455 = v344
	v456 = v344
	v457 = v344
	goto L88
L66:
	;
	if base.Ui32(int32(11)) < base.Ui32(v53) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(v53) < base.Ui32(int32(12)) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v501 = v338
	v502 = v53
	v503 = v344
	v504 = v344
	v505 = v344
	goto L64
L70:
	;
	switch v400 - int32(1) {
	case 0:
		v450 = v401
		goto L77
	case 1:
		v445 = v401
		goto L78
	case 2:
		goto L79
	case 3:
		v438 = v402
		goto L80
	case 4:
		v435 = v402
		goto L81
	case 5:
		v430 = v402
		goto L82
	case 6:
		goto L83
	case 7:
		v421 = v403
		goto L84
	case 8:
		v416 = v403
		goto L85
	case 9:
		v411 = v403
		goto L86
	case 10:
		goto L87
	default:
		v571 = v401
		v572 = v402
		v573 = v403
		goto L63
	}
L71:
	;
	v399 = v338
	v400 = v53
	v401 = v344
	v402 = v344
	v403 = v344
	goto L70
L72:
	;
	goto L73
L73:
	;
	v351 = v338
	v352 = v53
	v353 = v344
	v354 = v344
	v355 = v344
	goto L74
L74:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v358 = v357 + v354
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
	v362 = v361 + v355
	v364 = int32(4)
	v366 = v359 + v353 - v362 ^ base.I32_rotl(v362, v364)
	v370 = v358 - v366 ^ base.I32_rotl(v366, int32(6))
	v371 = v362 + v358
	v372 = v366 + v371
	v373 = v370 + v372
	v377 = v371 - v370 ^ base.I32_rotl(v370, int32(8))
	v381 = v372 - v377 ^ base.I32_rotl(v377, int32(16))
	v385 = v373 - v381 ^ base.I32_rotl(v381, int32(19))
	v386 = v377 + v373
	v387 = v381 + v386
	v388 = v385 + v387
	v392 = v386 - v385 ^ base.I32_rotl(v385, v364)
	v393 = int32(12)
	v394 = v351 + v393
	v396 = v352 - v393
	if base.Ui32(int32(11)) < base.Ui32(v396) {
		v351 = v394
		v352 = v396
		v353 = v387
		v354 = v388
		v355 = v392
		goto L74
	} else {
		goto L76
	}
L75:
	;
	v399 = v394
	v400 = v396
	v401 = v387
	v402 = v388
	v403 = v392
	goto L70
L76:
	;
	goto L75
L77:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v571 = v450 + v451
	v572 = v402
	v573 = v403
	goto L63
L78:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+1)))
	v450 = v446<<(uint(int32(8))%32) + v445
	goto L77
L79:
	;
	v441 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+2)))
	v445 = v441<<(uint(int32(16))%32) + v401
	goto L78
L80:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v571 = v439 + v401
	v572 = v438
	v573 = v403
	goto L63
L81:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+4)))
	v438 = v435 + v436
	goto L80
L82:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+5)))
	v435 = v431<<(uint(int32(8))%32) + v430
	goto L81
L83:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+6)))
	v430 = v426<<(uint(int32(16))%32) + v402
	goto L82
L84:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	v571 = v422 + v401
	v572 = v424 + v402
	v573 = v421
	goto L63
L85:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+8)))
	v421 = v417<<(uint(int32(8))%32) + v416
	goto L84
L86:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+9)))
	v416 = v412<<(uint(int32(16))%32) + v411
	goto L85
L87:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+10)))
	v411 = v407<<(uint(int32(24))%32) + v403
	goto L86
L88:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	v460 = v459 + v456
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v464 = v463 + v457
	v466 = int32(4)
	v468 = v461 + v455 - v464 ^ base.I32_rotl(v464, v466)
	v472 = v460 - v468 ^ base.I32_rotl(v468, int32(6))
	v473 = v464 + v460
	v474 = v468 + v473
	v475 = v472 + v474
	v479 = v473 - v472 ^ base.I32_rotl(v472, int32(8))
	v483 = v474 - v479 ^ base.I32_rotl(v479, int32(16))
	v487 = v475 - v483 ^ base.I32_rotl(v483, int32(19))
	v488 = v479 + v475
	v489 = v483 + v488
	v490 = v487 + v489
	v494 = v488 - v487 ^ base.I32_rotl(v487, v466)
	v495 = int32(12)
	v496 = v453 + v495
	v498 = v454 - v495
	if base.Ui32(int32(11)) < base.Ui32(v498) {
		v453 = v496
		v454 = v498
		v455 = v489
		v456 = v490
		v457 = v494
		goto L88
	} else {
		goto L90
	}
L89:
	;
	v501 = v496
	v502 = v498
	v503 = v489
	v504 = v490
	v505 = v494
	goto L64
L90:
	;
	goto L89
L91:
	;
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501))))
	v571 = v564 + v567
	v572 = v565
	v573 = v566
	goto L63
L92:
	;
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+1)))
	v564 = v560<<(uint(int32(8))%32) + v557
	v565 = v558
	v566 = v559
	goto L91
L93:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+2)))
	v557 = v553<<(uint(int32(16))%32) + v550
	v558 = v551
	v559 = v552
	goto L92
L94:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+3)))
	v550 = v546<<(uint(int32(24))%32) + v503
	v551 = v544
	v552 = v545
	goto L93
L95:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+4)))
	v544 = v540 + v542
	v545 = v541
	goto L94
L96:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+5)))
	v540 = v536<<(uint(int32(8))%32) + v534
	v541 = v535
	goto L95
L97:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+6)))
	v534 = v530<<(uint(int32(16))%32) + v528
	v535 = v529
	goto L96
L98:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+7)))
	v528 = v524<<(uint(int32(24))%32) + v504
	v529 = v523
	goto L97
L99:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+8)))
	v523 = v519<<(uint(int32(8))%32) + v518
	goto L98
L100:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+9)))
	v518 = v514<<(uint(int32(16))%32) + v513
	goto L99
L101:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v501)+10)))
	v513 = v509<<(uint(int32(24))%32) + v505
	goto L100
L102:
	;
	v901 = base.I32_wrap_i64(v49)
	v902 = F_strlen(m, v901)
	mBase = m.M
	v904 = v902 + int32(1)
	v910 = v904 - int32(1636608432)
	if v901&int32(3) != 0 {
		goto L160
	} else {
		goto L161
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L105
	} else {
		goto L153
	}
L104:
	;
	v605 = F_toast_raw_datum_size(m, v49)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	return int32(0)
L106:
	;
	v609 = base.I32_wrap_i64(v49)
	v610 = F_pg_detoast_datum_packed(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v612 = int32(1)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v610))))
	if v614&v612 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v617 = v612
	goto L110
L109:
	;
	v617 = int32(4)
	goto L110
L110:
	;
	v618 = v610 + v617
	v620 = v605 - int32(4)
	v626 = v620 - int32(1636608432)
	if v618&int32(3) != 0 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	if v610 == v609 {
		v1169 = v884
		goto L11
	} else {
		goto L151
	}
L112:
	;
	v858 = int32(14)
	v860 = v854 ^ v855 - base.I32_rotl(v854, v858)
	v864 = v860 ^ v853 - base.I32_rotl(v860, int32(11))
	v868 = v864 ^ v854 - base.I32_rotl(v864, int32(25))
	v872 = v868 ^ v860 - base.I32_rotl(v868, int32(16))
	v876 = v872 ^ v864 - base.I32_rotl(v872, int32(4))
	v880 = v876 ^ v868 - base.I32_rotl(v876, v858)
	v884 = v880 ^ v872 - base.I32_rotl(v880, int32(24))
	goto L111
L113:
	;
	switch v784 - int32(1) {
	case 0:
		v846 = v785
		v847 = v786
		v848 = v787
		goto L140
	case 1:
		v839 = v785
		v840 = v786
		v841 = v787
		goto L141
	case 2:
		v832 = v785
		v833 = v786
		v834 = v787
		goto L142
	case 3:
		v826 = v786
		v827 = v787
		goto L143
	case 4:
		v822 = v786
		v823 = v787
		goto L144
	case 5:
		v816 = v786
		v817 = v787
		goto L145
	case 6:
		v810 = v786
		v811 = v787
		goto L146
	case 7:
		v805 = v787
		goto L147
	case 8:
		v800 = v787
		goto L148
	case 9:
		v795 = v787
		goto L149
	case 10:
		goto L150
	default:
		v853 = v785
		v854 = v786
		v855 = v787
		goto L112
	}
L114:
	;
	v735 = v618
	v736 = v620
	v737 = v626
	v738 = v626
	v739 = v626
	goto L137
L115:
	;
	if base.Ui32(int32(11)) < base.Ui32(v620) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if base.Ui32(v620) < base.Ui32(int32(12)) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	v783 = v618
	v784 = v620
	v785 = v626
	v786 = v626
	v787 = v626
	goto L113
L119:
	;
	switch v682 - int32(1) {
	case 0:
		v732 = v683
		goto L126
	case 1:
		v727 = v683
		goto L127
	case 2:
		goto L128
	case 3:
		v720 = v684
		goto L129
	case 4:
		v717 = v684
		goto L130
	case 5:
		v712 = v684
		goto L131
	case 6:
		goto L132
	case 7:
		v703 = v685
		goto L133
	case 8:
		v698 = v685
		goto L134
	case 9:
		v693 = v685
		goto L135
	case 10:
		goto L136
	default:
		v853 = v683
		v854 = v684
		v855 = v685
		goto L112
	}
L120:
	;
	v681 = v618
	v682 = v620
	v683 = v626
	v684 = v626
	v685 = v626
	goto L119
L121:
	;
	goto L122
L122:
	;
	v633 = v618
	v634 = v620
	v635 = v626
	v636 = v626
	v637 = v626
	goto L123
L123:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	v640 = v639 + v636
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v633)+8))
	v644 = v643 + v637
	v646 = int32(4)
	v648 = v641 + v635 - v644 ^ base.I32_rotl(v644, v646)
	v652 = v640 - v648 ^ base.I32_rotl(v648, int32(6))
	v653 = v644 + v640
	v654 = v648 + v653
	v655 = v652 + v654
	v659 = v653 - v652 ^ base.I32_rotl(v652, int32(8))
	v663 = v654 - v659 ^ base.I32_rotl(v659, int32(16))
	v667 = v655 - v663 ^ base.I32_rotl(v663, int32(19))
	v668 = v659 + v655
	v669 = v663 + v668
	v670 = v667 + v669
	v674 = v668 - v667 ^ base.I32_rotl(v667, v646)
	v675 = int32(12)
	v676 = v633 + v675
	v678 = v634 - v675
	if base.Ui32(int32(11)) < base.Ui32(v678) {
		v633 = v676
		v634 = v678
		v635 = v669
		v636 = v670
		v637 = v674
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v681 = v676
	v682 = v678
	v683 = v669
	v684 = v670
	v685 = v674
	goto L119
L125:
	;
	goto L124
L126:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v853 = v732 + v733
	v854 = v684
	v855 = v685
	goto L112
L127:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+1)))
	v732 = v728<<(uint(int32(8))%32) + v727
	goto L126
L128:
	;
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+2)))
	v727 = v723<<(uint(int32(16))%32) + v683
	goto L127
L129:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v853 = v721 + v683
	v854 = v720
	v855 = v685
	goto L112
L130:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+4)))
	v720 = v717 + v718
	goto L129
L131:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+5)))
	v717 = v713<<(uint(int32(8))%32) + v712
	goto L130
L132:
	;
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+6)))
	v712 = v708<<(uint(int32(16))%32) + v684
	goto L131
L133:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v681)))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v681)+4))
	v853 = v704 + v683
	v854 = v706 + v684
	v855 = v703
	goto L112
L134:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+8)))
	v703 = v699<<(uint(int32(8))%32) + v698
	goto L133
L135:
	;
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+9)))
	v698 = v694<<(uint(int32(16))%32) + v693
	goto L134
L136:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681)+10)))
	v693 = v689<<(uint(int32(24))%32) + v685
	goto L135
L137:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	v742 = v741 + v738
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v735)))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v735)+8))
	v746 = v745 + v739
	v748 = int32(4)
	v750 = v743 + v737 - v746 ^ base.I32_rotl(v746, v748)
	v754 = v742 - v750 ^ base.I32_rotl(v750, int32(6))
	v755 = v746 + v742
	v756 = v750 + v755
	v757 = v754 + v756
	v761 = v755 - v754 ^ base.I32_rotl(v754, int32(8))
	v765 = v756 - v761 ^ base.I32_rotl(v761, int32(16))
	v769 = v757 - v765 ^ base.I32_rotl(v765, int32(19))
	v770 = v761 + v757
	v771 = v765 + v770
	v772 = v769 + v771
	v776 = v770 - v769 ^ base.I32_rotl(v769, v748)
	v777 = int32(12)
	v778 = v735 + v777
	v780 = v736 - v777
	if base.Ui32(int32(11)) < base.Ui32(v780) {
		v735 = v778
		v736 = v780
		v737 = v771
		v738 = v772
		v739 = v776
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v783 = v778
	v784 = v780
	v785 = v771
	v786 = v772
	v787 = v776
	goto L113
L139:
	;
	goto L138
L140:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783))))
	v853 = v846 + v849
	v854 = v847
	v855 = v848
	goto L112
L141:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+1)))
	v846 = v842<<(uint(int32(8))%32) + v839
	v847 = v840
	v848 = v841
	goto L140
L142:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+2)))
	v839 = v835<<(uint(int32(16))%32) + v832
	v840 = v833
	v841 = v834
	goto L141
L143:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+3)))
	v832 = v828<<(uint(int32(24))%32) + v785
	v833 = v826
	v834 = v827
	goto L142
L144:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+4)))
	v826 = v822 + v824
	v827 = v823
	goto L143
L145:
	;
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+5)))
	v822 = v818<<(uint(int32(8))%32) + v816
	v823 = v817
	goto L144
L146:
	;
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+6)))
	v816 = v812<<(uint(int32(16))%32) + v810
	v817 = v811
	goto L145
L147:
	;
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+7)))
	v810 = v806<<(uint(int32(24))%32) + v786
	v811 = v805
	goto L146
L148:
	;
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+8)))
	v805 = v801<<(uint(int32(8))%32) + v800
	goto L147
L149:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+9)))
	v800 = v796<<(uint(int32(16))%32) + v795
	goto L148
L150:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v783)+10)))
	v795 = v791<<(uint(int32(24))%32) + v787
	goto L149
L151:
	;
	F_pfree(m, v610)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L105
	} else {
		goto L152
	}
L152:
	;
	v1169 = v884
	goto L11
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56))) = v53
	F_errmsg_internal(m, int32(_a_F_MemoizeHash_hash_1), v56)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L105
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_MemoizeHash_hash_2), int32(408), int32(_a_F_MemoizeHash_hash_3))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L105
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	v1169 = v1164 ^ v1156 - base.I32_rotl(v1164, int32(24))
	goto L11
L157:
	;
	v1142 = int32(14)
	v1144 = v1138 ^ v1139 - base.I32_rotl(v1138, v1142)
	v1148 = v1144 ^ v1137 - base.I32_rotl(v1144, int32(11))
	v1152 = v1148 ^ v1138 - base.I32_rotl(v1148, int32(25))
	v1156 = v1152 ^ v1144 - base.I32_rotl(v1152, int32(16))
	v1160 = v1156 ^ v1148 - base.I32_rotl(v1156, int32(4))
	v1164 = v1160 ^ v1152 - base.I32_rotl(v1160, v1142)
	goto L156
L158:
	;
	switch v1068 - int32(1) {
	case 0:
		v1130 = v1069
		v1131 = v1070
		v1132 = v1071
		goto L185
	case 1:
		v1123 = v1069
		v1124 = v1070
		v1125 = v1071
		goto L186
	case 2:
		v1116 = v1069
		v1117 = v1070
		v1118 = v1071
		goto L187
	case 3:
		v1110 = v1070
		v1111 = v1071
		goto L188
	case 4:
		v1106 = v1070
		v1107 = v1071
		goto L189
	case 5:
		v1100 = v1070
		v1101 = v1071
		goto L190
	case 6:
		v1094 = v1070
		v1095 = v1071
		goto L191
	case 7:
		v1089 = v1071
		goto L192
	case 8:
		v1084 = v1071
		goto L193
	case 9:
		v1079 = v1071
		goto L194
	case 10:
		goto L195
	default:
		v1137 = v1069
		v1138 = v1070
		v1139 = v1071
		goto L157
	}
L159:
	;
	v1019 = v901
	v1020 = v904
	v1021 = v910
	v1022 = v910
	v1023 = v910
	goto L182
L160:
	;
	if base.Ui32(int32(11)) < base.Ui32(v904) {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	if base.Ui32(v904) < base.Ui32(int32(12)) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v1067 = v901
	v1068 = v904
	v1069 = v910
	v1070 = v910
	v1071 = v910
	goto L158
L164:
	;
	switch v966 - int32(1) {
	case 0:
		v1016 = v967
		goto L171
	case 1:
		v1011 = v967
		goto L172
	case 2:
		goto L173
	case 3:
		v1004 = v968
		goto L174
	case 4:
		v1001 = v968
		goto L175
	case 5:
		v996 = v968
		goto L176
	case 6:
		goto L177
	case 7:
		v987 = v969
		goto L178
	case 8:
		v982 = v969
		goto L179
	case 9:
		v977 = v969
		goto L180
	case 10:
		goto L181
	default:
		v1137 = v967
		v1138 = v968
		v1139 = v969
		goto L157
	}
L165:
	;
	v965 = v901
	v966 = v904
	v967 = v910
	v968 = v910
	v969 = v910
	goto L164
L166:
	;
	goto L167
L167:
	;
	v917 = v901
	v918 = v904
	v919 = v910
	v920 = v910
	v921 = v910
	goto L168
L168:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v917)+4))
	v924 = v923 + v920
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v917)))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v917)+8))
	v928 = v927 + v921
	v930 = int32(4)
	v932 = v925 + v919 - v928 ^ base.I32_rotl(v928, v930)
	v936 = v924 - v932 ^ base.I32_rotl(v932, int32(6))
	v937 = v928 + v924
	v938 = v932 + v937
	v939 = v936 + v938
	v943 = v937 - v936 ^ base.I32_rotl(v936, int32(8))
	v947 = v938 - v943 ^ base.I32_rotl(v943, int32(16))
	v951 = v939 - v947 ^ base.I32_rotl(v947, int32(19))
	v952 = v943 + v939
	v953 = v947 + v952
	v954 = v951 + v953
	v958 = v952 - v951 ^ base.I32_rotl(v951, v930)
	v959 = int32(12)
	v960 = v917 + v959
	v962 = v918 - v959
	if base.Ui32(int32(11)) < base.Ui32(v962) {
		v917 = v960
		v918 = v962
		v919 = v953
		v920 = v954
		v921 = v958
		goto L168
	} else {
		goto L170
	}
L169:
	;
	v965 = v960
	v966 = v962
	v967 = v953
	v968 = v954
	v969 = v958
	goto L164
L170:
	;
	goto L169
L171:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	v1137 = v1016 + v1017
	v1138 = v968
	v1139 = v969
	goto L157
L172:
	;
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+1)))
	v1016 = v1012<<(uint(int32(8))%32) + v1011
	goto L171
L173:
	;
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+2)))
	v1011 = v1007<<(uint(int32(16))%32) + v967
	goto L172
L174:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	v1137 = v1005 + v967
	v1138 = v1004
	v1139 = v969
	goto L157
L175:
	;
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+4)))
	v1004 = v1001 + v1002
	goto L174
L176:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+5)))
	v1001 = v997<<(uint(int32(8))%32) + v996
	goto L175
L177:
	;
	v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+6)))
	v996 = v992<<(uint(int32(16))%32) + v968
	goto L176
L178:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v965)))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v965)+4))
	v1137 = v988 + v967
	v1138 = v990 + v968
	v1139 = v987
	goto L157
L179:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+8)))
	v987 = v983<<(uint(int32(8))%32) + v982
	goto L178
L180:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+9)))
	v982 = v978<<(uint(int32(16))%32) + v977
	goto L179
L181:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965)+10)))
	v977 = v973<<(uint(int32(24))%32) + v969
	goto L180
L182:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+4))
	v1026 = v1025 + v1022
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1019)))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+8))
	v1030 = v1029 + v1023
	v1032 = int32(4)
	v1034 = v1027 + v1021 - v1030 ^ base.I32_rotl(v1030, v1032)
	v1038 = v1026 - v1034 ^ base.I32_rotl(v1034, int32(6))
	v1039 = v1030 + v1026
	v1040 = v1034 + v1039
	v1041 = v1038 + v1040
	v1045 = v1039 - v1038 ^ base.I32_rotl(v1038, int32(8))
	v1049 = v1040 - v1045 ^ base.I32_rotl(v1045, int32(16))
	v1053 = v1041 - v1049 ^ base.I32_rotl(v1049, int32(19))
	v1054 = v1045 + v1041
	v1055 = v1049 + v1054
	v1056 = v1053 + v1055
	v1060 = v1054 - v1053 ^ base.I32_rotl(v1053, v1032)
	v1061 = int32(12)
	v1062 = v1019 + v1061
	v1064 = v1020 - v1061
	if base.Ui32(int32(11)) < base.Ui32(v1064) {
		v1019 = v1062
		v1020 = v1064
		v1021 = v1055
		v1022 = v1056
		v1023 = v1060
		goto L182
	} else {
		goto L184
	}
L183:
	;
	v1067 = v1062
	v1068 = v1064
	v1069 = v1055
	v1070 = v1056
	v1071 = v1060
	goto L158
L184:
	;
	goto L183
L185:
	;
	v1133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067))))
	v1137 = v1130 + v1133
	v1138 = v1131
	v1139 = v1132
	goto L157
L186:
	;
	v1126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+1)))
	v1130 = v1126<<(uint(int32(8))%32) + v1123
	v1131 = v1124
	v1132 = v1125
	goto L185
L187:
	;
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+2)))
	v1123 = v1119<<(uint(int32(16))%32) + v1116
	v1124 = v1117
	v1125 = v1118
	goto L186
L188:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+3)))
	v1116 = v1112<<(uint(int32(24))%32) + v1069
	v1117 = v1110
	v1118 = v1111
	goto L187
L189:
	;
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+4)))
	v1110 = v1106 + v1108
	v1111 = v1107
	goto L188
L190:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+5)))
	v1106 = v1102<<(uint(int32(8))%32) + v1100
	v1107 = v1101
	goto L189
L191:
	;
	v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+6)))
	v1100 = v1096<<(uint(int32(16))%32) + v1094
	v1101 = v1095
	goto L190
L192:
	;
	v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+7)))
	v1094 = v1090<<(uint(int32(24))%32) + v1070
	v1095 = v1089
	goto L191
L193:
	;
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+8)))
	v1089 = v1085<<(uint(int32(8))%32) + v1084
	goto L192
L194:
	;
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+9)))
	v1084 = v1080<<(uint(int32(16))%32) + v1079
	goto L193
L195:
	;
	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067)+10)))
	v1079 = v1075<<(uint(int32(24))%32) + v1071
	goto L194
L196:
	;
	goto L7
L197:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v14)+152))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+148))
	v1192 = int32(0)
	v1195 = v2
	goto L198
L198:
	;
	v1204 = base.I32_rotl(v1195, int32(1))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205+v1192))))
	if v1207 == int32(0) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v1233 = v1226
	goto L1
L200:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1192<<(uint(int32(2))%32))))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v1221 = *(*int64)(unsafe.Add(mBase, uint32(v1217+v1192<<(uint(int32(3))%32))))
	v1222 = F_FunctionCall1Coll(m, v1190+v1192*int32(28), v1216, v1221)
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L105
	} else {
		goto L203
	}
L201:
	;
	v1226 = v1204
	goto L202
L202:
	;
	v1228 = v1192 + int32(1)
	if v1228 != v15 {
		v1192 = v1228
		v1195 = v1226
		goto L198
	} else {
		goto L204
	}
L203:
	;
	v1226 = v1204 ^ base.I32_wrap_i64(v1222)
	goto L202
L204:
	;
	goto L199
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
func F_macaddr8tomacaddr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc0(m, int32(6))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+3)))
		if v11 == int32(255) {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)))
			if v14 == int32(254) {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
				*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v40)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)) = uint8(v42)
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+2)) = uint8(v44)
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+5)))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+3)) = uint8(v46)
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+6)))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v48)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+7)))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+5)) = uint8(v50)
				v54 = base.I64_extend_i32_u(v7)
				return v54
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v18 = F_errsave_start(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					if v18 == int32(0) {
						v54 = int64(0)
						return v54
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_macaddr8tomacaddr_0), int32(0))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int64(0)
							} else {
								F_errhint(m, int32(_a_F_macaddr8tomacaddr_1), int32(0))
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v17, int32(_a_F_macaddr8tomacaddr_2), int32(559), int32(_a_F_macaddr8tomacaddr_3))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return int64(0)
									} else {
										return int64(0)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = F_errsave_start(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				if v18 == int32(0) {
					v54 = int64(0)
					return v54
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_macaddr8tomacaddr_0), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							F_errhint(m, int32(_a_F_macaddr8tomacaddr_1), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v17, int32(_a_F_macaddr8tomacaddr_2), int32(559), int32(_a_F_macaddr8tomacaddr_3))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									return int64(0)
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
	var v63 int64
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
	v11 = v9 - int32(80)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(2615)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v34
	v39 = v11 + int32(68)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(3601)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v50
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
	v59 = v11 + int32(8)
	v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+4)))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(3600)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v82+v83)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v87
	F_add_exact_object_address(m, v11+int32(68), v30)
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
	m.G0 = v11 + int32(80)
	return
}
func F_makeConst(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32, l6 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	v6 = l5
	v7 = l6
	v10 = F_palloc0(m, int32(40))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(7)
		if base.B2i32(l3 != int32(-1))|v6 == int32(0) {
			v22 = F_pg_detoast_datum(m, base.I32_wrap_i64(l4))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = base.I64_extend_i32_u(v22)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(-1)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+33)) = uint8(v7)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v6)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v25
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l0
				return v10
			}
		} else {
			v25 = l4
			*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(-1)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+33)) = uint8(v7)
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v6)
			*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v25
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
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v2 {
		v31 = v2
		m.G0 = v7 + int32(16)
		return v31
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = v11 - int32(7)
		if v13 != 0 {
			if v13 != int32(14) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
				v27 = F_list_make1_impl(m, int32(1), v7+int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = v27
					m.G0 = v7 + int32(16)
					return v31
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v16 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
					v27 = F_list_make1_impl(m, int32(1), v7+int32(8))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v27
						m.G0 = v7 + int32(16)
						return v31
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v31 = v17
					m.G0 = v7 + int32(16)
					return v31
				}
			}
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v18 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
				v27 = F_list_make1_impl(m, int32(1), v7+int32(8))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = v27
					m.G0 = v7 + int32(16)
					return v31
				}
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != int64(0) {
					v31 = v2
					m.G0 = v7 + int32(16)
					return v31
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l0
					v27 = F_list_make1_impl(m, int32(1), v7+int32(8))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = v27
						m.G0 = v7 + int32(16)
						return v31
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
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
								F_errfinish(m, int32(_a_F_make_distinct_op_2), int32(3104), int32(_a_F_make_distinct_op_3))
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
							F_errfinish(m, int32(_a_F_make_distinct_op_2), int32(3098), int32(_a_F_make_distinct_op_3))
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
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
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
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
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
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
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
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v352 int32
	_ = v352
	v4 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == int32(17) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if int32(0) < v196 {
		goto L89
	} else {
		goto L90
	}
L2:
	;
	v183 = F_palloc(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L80
	}
L3:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v182 = int32(base.Ui32(v176)>>(uint(int32(2))%32)) - int32(4)
	goto L2
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	v182 = v39
	goto L2
L18:
	;
	v44 = int32(1)
	v182 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
	goto L2
L19:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	if v81 != 0 {
		goto L39
	} else {
		goto L40
	}
L20:
	;
	v60 = F_strlen(m, v59)
	mBase = m.M
	if v60 != 0 {
		goto L26
	} else {
		goto L27
	}
L21:
	;
	v53 = F_DirectFunctionCall1Coll(m, int32(626), int32(0), v48)
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
	v57 = F_text_to_cstring(m, base.I32_wrap_i64(v48))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v59 = base.I32_wrap_i64(v53)
	goto L20
L25:
	;
	v59 = v57
	goto L20
L26:
	;
	v61 = F_pg_newlocale_from_collation(m, l2)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v73 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v63 != int32(1) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v196 = v60
	v198 = v4
	v199 = v59
	v200 = v66
	v201 = v79
	goto L1
L32:
	;
	v76 = int32(1853)
	goto L34
L33:
	;
	v76 = int32(1854)
	goto L34
L34:
	;
	if v73 == int32(6) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v79 = int32(1852)
	goto L37
L36:
	;
	v79 = v76
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
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2]))
	if v83 == l2 {
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v86 = int32(_a_F_make_greater_string_0)
	v87 = int32(_a_F_make_greater_string_1)
	v89 = int32(1)
	v92 = F_varstr_cmp(m, v87, v89, v86, v89, l2)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	if v92 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v96 = v86
	goto L46
L45:
	;
	v96 = v87
	goto L46
L46:
	;
	v97 = int32(1)
	v100 = F_varstr_cmp(m, v96, v97, int32(_a_F_make_greater_string_2), v97, l2)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	if v100 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v104 = int32(_a_F_make_greater_string_2)
	goto L50
L49:
	;
	v104 = v96
	goto L50
L50:
	;
	v105 = int32(1)
	v108 = F_varstr_cmp(m, v104, v105, int32(_a_F_make_greater_string_3), v105, l2)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[2])) = l2
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v117 = int32(57)
	goto L54
L53:
	;
	v117 = v114
	goto L54
L54:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])) = uint8(v117)
	goto L38
L55:
	;
	v125 = F_palloc(m, v60+int32(2))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v149 = v60 + int32(5)
	v150 = F_palloc(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L69
	}
L58:
	;
	if v60 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v125, v59, v60)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	v130 = v60 + v125
	v131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)) = uint8(v131)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v129)
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v141 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v196 = v60
	v198 = v125
	v199 = v59
	v200 = base.I64_extend_i32_u(v125)
	v201 = v147
	goto L1
L63:
	;
	v144 = int32(1853)
	goto L65
L64:
	;
	v144 = int32(1854)
	goto L65
L65:
	;
	if v141 == int32(6) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v147 = int32(1852)
	goto L68
L67:
	;
	v147 = v144
	goto L68
L68:
	;
	goto L62
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v149 << (uint(int32(2)) % 32)
	v156 = v150 + int32(4)
	if v60 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	base.MemoryCopy(m, v156, v59, v60)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_make_greater_string[0])))
	*(*uint8)(unsafe.Add(mBase, uint32(v156+v60))) = uint8(v160)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_make_greater_string[1]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v169 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v196 = v60
	v198 = v150
	v199 = v59
	v200 = base.I64_extend_i32_u(v150)
	v201 = v175
	goto L1
L74:
	;
	v172 = int32(1853)
	goto L76
L75:
	;
	v172 = int32(1854)
	goto L76
L76:
	;
	if v169 == int32(6) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v175 = int32(1852)
	goto L79
L78:
	;
	v175 = v172
	goto L79
L79:
	;
	goto L73
L80:
	;
	if v182 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v185 = int32(1)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v187&v185 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v196 = v182
	v198 = v4
	v199 = v183
	v200 = v193
	v201 = int32(1585)
	goto L1
L84:
	;
	v190 = v185
	goto L86
L85:
	;
	v190 = int32(4)
	goto L86
L86:
	;
	base.MemoryCopy(m, v183, v19+v190, v182)
	goto L83
L87:
	;
	F_pfree(m, v199)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L126
	}
L88:
	;
	F_pfree(m, v198)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L125
	}
L89:
	;
	v209 = v196
	goto L92
L90:
	;
	goto L91
L91:
	;
	v318 = int32(0)
	if v198 == v318 {
		v337 = v318
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
	v225 = F_pg_mbcliplen(m, v199, v209, v209-int32(1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L97
	}
L95:
	;
	v228 = int32(1)
	goto L96
L96:
	;
	v230 = v209 + v199 - v228
	v231 = m.T0[v201].(func(*base.Module, int32, int32) int32)(m, v230, v228)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L98
	}
L97:
	;
	v228 = v209 - v225
	goto L96
L98:
	;
	if v231 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v234 = v209 + int32(4)
	goto L102
L100:
	;
	goto L101
L101:
	;
	v298 = v209 - v228
	v300 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199+v298))) = uint8(v300)
	if v300 < v298 {
		v209 = v298
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
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v271)+24))
	v273 = F_FunctionCall2Coll(m, l1, l2, v200, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L114
	}
L105:
	;
	v253 = F_palloc(m, v234)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v268 = F_string_to_const(m, v199, v15)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L7
	} else {
		goto L113
	}
L108:
	;
	if v209 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	base.MemoryCopy(m, v253+int32(4), v199, v209)
	goto L111
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v234 << (uint(int32(2)) % 32)
	v260 = int32(-1)
	v261 = int32(0)
	v266 = F_makeConst(m, int32(17), v260, v261, v260, base.I64_extend_i32_u(v253), v261, v261)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	v271 = v266
	goto L104
L113:
	;
	v271 = v268
	goto L104
L114:
	;
	if v273 != int64(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v198 != 0 {
		v321 = v271
		goto L88
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v271)+24))
	F_pfree(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L7
	} else {
		goto L119
	}
L118:
	;
	v337 = v271
	goto L87
L119:
	;
	F_pfree(m, v271)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	v282 = m.T0[v201].(func(*base.Module, int32, int32) int32)(m, v230, v228)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L121
	}
L121:
	;
	if v282 != 0 {
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
	v321 = v318
	goto L88
L125:
	;
	v337 = v321
	goto L87
L126:
	;
	return v337
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
	v4 = F_palloc0(m, int32(120))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+81)) = uint8(v8)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+68)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = l0
		if l0 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+100)) = v15
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+104)) = v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+108)) = v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+112)) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+116)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+84)) = v25
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
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+340))
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
	var v146 int32
	_ = v146
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
	var v166 int32
	_ = v166
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
	v146 = F_errdetail(m, int32(_a_F_make_pathkeys_for_window_1), int32(0))
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
	v166 = F_errdetail(m, int32(_a_F_make_pathkeys_for_window_6), int32(0))
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
										F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(877), int32(_a_F_make_scalar_array_op_2))
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
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(892), int32(_a_F_make_scalar_array_op_2))
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
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(918), int32(_a_F_make_scalar_array_op_2))
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
																	F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(923), int32(_a_F_make_scalar_array_op_2))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
																								F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
											F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(892), int32(_a_F_make_scalar_array_op_2))
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
													F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(918), int32(_a_F_make_scalar_array_op_2))
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
															F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(923), int32(_a_F_make_scalar_array_op_2))
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
																				F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
																						F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
																				F_errfinish(m, int32(_a_F_make_scalar_array_op_1), int32(944), int32(_a_F_make_scalar_array_op_2))
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
func F_match_db_entries(m *base.Module, l0 int32, l1 int64) int32 {
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
func F_matchingjoinsel(m *base.Module, l0 int32) int64 {
	return int64(4576918229304087675)
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
			v31 = F_FileWriteV(m, v18, v11+int32(56), int32(1), base.I64_extend_i32_u(l2<<(uint(int32(13))%32)&int32(1073733632)), int32(167772179))
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
										F_errfinish(m, int32(_a_F_mdextend_3), int32(530), int32(_a_F_mdextend_4))
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
										F_errfinish(m, int32(_a_F_mdextend_3), int32(537), int32(_a_F_mdextend_4))
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
						F_errfinish(m, int32(_a_F_mdextend_3), int32(515), int32(_a_F_mdextend_4))
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
	F_errfinish(m, int32(_a_F_mdnblocks_3), int32(1893), int32(_a_F_mdnblocks_4))
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
	F_errfinish(m, int32(_a_F_mdnblocks_3), int32(1266), int32(_a_F_mdnblocks_6))
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
										F_errfinish(m, int32(_a_F_mdopenfork_2), int32(697), int32(_a_F_mdopenfork_3))
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
									F_errfinish(m, int32(_a_F_mdopenfork_2), int32(697), int32(_a_F_mdopenfork_3))
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
func F_metaphone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v814 int32
	_ = v814
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v853 int32
	_ = v853
	var v862 int32
	_ = v862
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v901 int32
	_ = v901
	var v910 int32
	_ = v910
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v962 int32
	_ = v962
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var v995 int32
	_ = v995
	var v1004 int32
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1057 int32
	_ = v1057
	var v1066 int32
	_ = v1066
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_text_to_cstring(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L6
	} else {
		goto L307
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L6
	} else {
		goto L303
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L6
	} else {
		goto L299
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L6
	} else {
		goto L296
	}
L5:
	;
	v1088 = F_cstring_to_text(m, v1081)
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L6
	} else {
		goto L295
	}
L6:
	;
	return int64(0)
L7:
	;
	v21 = F_strlen(m, v17)
	mBase = m.M
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v1081 = int32(_a_F_metaphone_0)
	goto L5
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(256)) <= base.Ui32(v21) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if int32(256) <= v27 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if v27 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v37 = F_palloc(m, v27+int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.Ui32((v39-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	switch v99&int32(255) - int32(65) {
	case 0:
		goto L38
	default:
		v208 = v97
		v211 = v2
		goto L32
	case 4, 8, 14, 20:
		goto L34
	case 6, 10, 15:
		goto L37
	case 22:
		goto L36
	case 23:
		goto L35
	}
L17:
	;
	v48 = v39 - int32(32)
	goto L19
L18:
	;
	v48 = v39
	goto L19
L19:
	;
	if base.Ui32((v48&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v97 = int32(0)
	v98 = v17
	v99 = v48
	goto L16
L21:
	;
	goto L22
L22:
	;
	v59 = int32(0)
	v61 = v48
	goto L23
L23:
	;
	if v61&int32(255) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v97 = v77
	v98 = v78
	v99 = v88
	goto L16
L25:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v74)
	v1081 = v37
	goto L5
L26:
	;
	goto L27
L27:
	;
	v77 = v59 + int32(1)
	v78 = v17 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.Ui32((v79-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v88 = v79 - int32(32)
	goto L30
L29:
	;
	v88 = v79
	goto L30
L30:
	;
	if base.Ui32(int32(25)) < base.Ui32((v88&int32(223)-int32(65))&int32(255)) {
		v59 = v77
		v61 = v88
		goto L23
	} else {
		goto L31
	}
L31:
	;
	goto L24
L32:
	;
	v212 = v208 + v17
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if base.Ui32((v213-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L33:
	;
	v205 = int32(1)
	v208 = v97 + v205
	v211 = v205
	goto L32
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v99)
	goto L33
L35:
	;
	v202 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v202)
	goto L33
L36:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if base.Ui32((v155-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if base.Ui32((v136-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v112 = int32(1)
	v114 = v97 + v112
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v114))))
	if base.Ui32((v116-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v125 = v116 - int32(32)
	goto L41
L40:
	;
	v125 = v116
	goto L41
L41:
	;
	if v125&int32(255) == int32(69) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v130)
	v208 = v97 + int32(2)
	v211 = v112
	goto L32
L43:
	;
	goto L44
L44:
	;
	v134 = int32(65)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v134)
	v208 = v114
	v211 = v112
	goto L32
L45:
	;
	v145 = v136 - int32(32)
	goto L47
L46:
	;
	v145 = v136
	goto L47
L47:
	;
	if v145&int32(255) != int32(78) {
		v208 = v97
		v211 = v2
		goto L32
	} else {
		goto L48
	}
L48:
	;
	v150 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v150)
	v208 = v97 + int32(2)
	v211 = int32(1)
	goto L32
L49:
	;
	if base.Ui32((v164-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v164)
	v208 = v97 + int32(2)
	v211 = int32(1)
	goto L32
L51:
	;
	v164 = v155 - int32(32)
	goto L53
L52:
	;
	v164 = v155
	goto L53
L53:
	;
	switch v164&int32(255) - int32(72) {
	case 0, 10:
		goto L50
	default:
		goto L49
	}
L54:
	;
	v181 = v164 - int32(32)
	goto L56
L55:
	;
	v181 = v164
	goto L56
L56:
	;
	if base.Ui32(int32(25)) < base.Ui32((v181-int32(65))&int32(255)) {
		v208 = v97
		v211 = v2
		goto L32
	} else {
		goto L57
	}
L57:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181&int32(255))+uint32(_c_F_metaphone[0]))))
	if v192&int32(1) == int32(0) {
		v208 = v97
		v211 = v2
		goto L32
	} else {
		goto L58
	}
L58:
	;
	v197 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v37))) = uint8(v197)
	v208 = v97 + int32(2)
	v211 = int32(1)
	goto L32
L59:
	;
	v1075 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1066+v37))) = uint8(v1075)
	v1081 = v37
	goto L5
L60:
	;
	v222 = v213 - int32(32)
	goto L62
L61:
	;
	v222 = v213
	goto L62
L62:
	;
	if base.B2i32(v222&int32(255) == int32(0))|base.B2i32(base.Ui32(v27) <= base.Ui32(v211)) != 0 {
		v1066 = v211
		goto L59
	} else {
		goto L63
	}
L63:
	;
	v229 = v208
	v230 = v212
	v231 = v222
	v232 = v211
	goto L64
L64:
	;
	if base.Ui32(int32(25)) < base.Ui32((v231&int32(223)-int32(65))&int32(255)) {
		v1038 = v229
		v1041 = v232
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v1066 = v1041
	goto L59
L66:
	;
	v1046 = v1038 + int32(1)
	v1047 = v17 + v1046
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
	if base.Ui32((v1048-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L290
	} else {
		goto L291
	}
L67:
	;
	v249 = base.B2i32(v229 <= int32(0))
	if v229 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L68:
	;
	v1038 = v229 + v1035
	v1041 = v1034
	goto L66
L69:
	;
	v1034 = v232 + int32(1)
	v1035 = int32(0)
	goto L68
L70:
	;
	v1023 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v1023)
	goto L69
L71:
	;
	v1019 = int32(1)
	v1034 = v232 + v1019
	v1035 = v1019
	goto L68
L72:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v892-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L257
	} else {
		goto L258
	}
L73:
	;
	v271 = int32(0)
	switch v231&int32(255) - int32(66) {
	case 0:
		goto L95
	case 1:
		goto L72
	case 2:
		goto L94
	default:
		v1034 = v232
		v1035 = v271
		goto L68
	case 4, 8, 10, 11, 12, 16:
		goto L81
	case 5:
		goto L93
	case 6:
		goto L92
	case 9:
		goto L91
	case 14:
		goto L90
	case 15:
		goto L89
	case 17:
		goto L88
	case 18:
		goto L87
	case 20:
		goto L86
	case 21:
		goto L85
	case 22:
		goto L84
	case 23:
		goto L83
	case 24:
		goto L82
	}
L74:
	;
	v250 = int32(255)
	v251 = v231 & v250
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v254-int32(97))&v250) < base.Ui32(int32(26)) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v263 = v254 - int32(32)
	goto L77
L76:
	;
	v263 = v254
	goto L77
L77:
	;
	if v251 != v263&int32(255) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	if v251 != int32(67) {
		v1038 = v229
		v1041 = v232
		goto L66
	} else {
		goto L79
	}
L79:
	;
	goto L72
L80:
	;
	v1034 = v232 + int32(1)
	v1035 = v271
	goto L68
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v231)
	goto L80
L82:
	;
	v882 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v882)
	goto L80
L83:
	;
	v844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v844-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L248
	} else {
		goto L249
	}
L84:
	;
	v834 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v834)
	v837 = v232 + int32(1)
	if v27 <= v837 {
		goto L245
	} else {
		goto L246
	}
L85:
	;
	v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v796-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L237
	} else {
		goto L238
	}
L86:
	;
	v794 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v794)
	goto L80
L87:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v748-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L230
	} else {
		goto L231
	}
L88:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v651-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L205
	} else {
		goto L206
	}
L89:
	;
	v649 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v649)
	goto L80
L90:
	;
	v627 = v232 + v37
	v629 = v232 + int32(1)
	v630 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v630-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L195
	} else {
		goto L196
	}
L91:
	;
	if v249 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L92:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v529-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L168
	} else {
		goto L169
	}
L93:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v353-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L94:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if base.Ui32((v298-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L104
	} else {
		goto L105
	}
L95:
	;
	if v249 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v280-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v296 = int32(66)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v296)
	goto L80
L99:
	;
	v289 = v280 - int32(32)
	goto L101
L100:
	;
	v289 = v280
	goto L101
L101:
	;
	if v289&int32(255) == int32(77) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L102
	}
L102:
	;
	goto L98
L103:
	;
	v351 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v351)
	goto L80
L104:
	;
	v307 = v298 - int32(32)
	goto L106
L105:
	;
	v307 = v298
	goto L106
L106:
	;
	if v307&int32(255) != int32(71) {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v312-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v321 = v312 - int32(32)
	goto L110
L109:
	;
	v321 = v312
	goto L110
L110:
	;
	if base.Ui32((v321-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v330 = v321 - int32(32)
	goto L113
L112:
	;
	v330 = v321
	goto L113
L113:
	;
	if base.Ui32(int32(25)) < base.Ui32((v330-int32(65))&int32(255)) {
		goto L103
	} else {
		goto L114
	}
L114:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330&int32(255))+uint32(_c_F_metaphone[0]))))
	if v341&int32(8) == int32(0) {
		goto L103
	} else {
		goto L115
	}
L115:
	;
	v347 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v347)
	goto L71
L116:
	;
	if base.Ui32((v362-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L156
	} else {
		goto L157
	}
L117:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v429-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L118:
	;
	if int32(3) <= v229 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v362 = v353 - int32(32)
	goto L121
L120:
	;
	v362 = v353
	goto L121
L121:
	;
	switch v362&int32(255) - int32(72) {
	case 0:
		goto L118
	default:
		goto L116
	case 6:
		goto L117
	}
L122:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(3)))))
	if base.Ui32((v372-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v382 = int32(0)
	goto L124
L124:
	;
	if base.Ui32((v382-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v381 = v372 - int32(32)
	goto L127
L126:
	;
	v381 = v372
	goto L127
L127:
	;
	v382 = v381
	goto L124
L128:
	;
	if v229 < int32(4) {
		goto L134
	} else {
		goto L135
	}
L129:
	;
	v391 = v382 - int32(32)
	goto L131
L130:
	;
	v391 = v382
	goto L131
L131:
	;
	if base.Ui32(int32(25)) < base.Ui32((v391-int32(65))&int32(255)) {
		goto L128
	} else {
		goto L132
	}
L132:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391&int32(255))+uint32(_c_F_metaphone[0]))))
	if v402&int32(16) == int32(0) {
		goto L128
	} else {
		goto L133
	}
L133:
	;
	v1034 = v232
	v1035 = v271
	goto L68
L134:
	;
	v427 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v427)
	goto L71
L135:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(4)))))
	if base.Ui32((v411-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v420 = v411 - int32(32)
	goto L138
L137:
	;
	v420 = v411
	goto L138
L138:
	;
	if v420&int32(255) != int32(72) {
		goto L134
	} else {
		goto L139
	}
L139:
	;
	v1034 = v232
	v1035 = v271
	goto L68
L140:
	;
	v438 = v429 - int32(32)
	goto L142
L141:
	;
	v438 = v429
	goto L142
L142:
	;
	if base.Ui32(int32(25)) < base.Ui32((v438&int32(223)-int32(65))&int32(255)) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L143
	}
L143:
	;
	if v438&int32(255) == int32(69) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v451 == int32(0) {
		v461 = v451
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	v477 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v477)
	goto L80
L147:
	;
	if base.Ui32((v461-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v454 == int32(0) {
		v461 = v454
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if v457 == int32(0) {
		v461 = v457
		goto L147
	} else {
		goto L150
	}
L150:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+3)))
	v461 = v460
	goto L147
L151:
	;
	v470 = v461 - int32(32)
	goto L153
L152:
	;
	v470 = v461
	goto L153
L153:
	;
	if v470&int32(255) == int32(68) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L154
	}
L154:
	;
	goto L146
L155:
	;
	v527 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v527)
	goto L80
L156:
	;
	v487 = v362 - int32(32)
	goto L158
L157:
	;
	v487 = v362
	goto L158
L158:
	;
	if base.Ui32(int32(25)) < base.Ui32((v487-int32(65))&int32(255)) {
		goto L155
	} else {
		goto L159
	}
L159:
	;
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487&int32(255))+uint32(_c_F_metaphone[0]))))
	if v498&int32(8) == int32(0) {
		goto L155
	} else {
		goto L160
	}
L160:
	;
	if v249 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v507-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	goto L163
L163:
	;
	v523 = int32(74)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v523)
	goto L80
L164:
	;
	v516 = v507 - int32(32)
	goto L166
L165:
	;
	v516 = v507
	goto L166
L166:
	;
	if v516&int32(255) == int32(71) {
		goto L155
	} else {
		goto L167
	}
L167:
	;
	goto L163
L168:
	;
	v538 = v529 - int32(32)
	goto L170
L169:
	;
	v538 = v529
	goto L170
L170:
	;
	if base.Ui32((v538-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v547 = v538 - int32(32)
	goto L173
L172:
	;
	v547 = v538
	goto L173
L173:
	;
	if base.Ui32(int32(25)) < base.Ui32((v547-int32(65))&int32(255)) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L174
	}
L174:
	;
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547&int32(255))+uint32(_c_F_metaphone[0]))))
	if v558&int32(1) == int32(0) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L175
	}
L175:
	;
	if v229 <= int32(0) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v603 = int32(72)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v603)
	goto L69
L177:
	;
	v577 = int32(0)
	goto L179
L178:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v566-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if base.Ui32((v577-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v575 = v566 - int32(32)
	goto L182
L181:
	;
	v575 = v566
	goto L182
L182:
	;
	v577 = v575
	goto L179
L183:
	;
	v586 = v577 - int32(32)
	goto L185
L184:
	;
	v586 = v577
	goto L185
L185:
	;
	if base.Ui32(int32(25)) < base.Ui32((v586-int32(65))&int32(255)) {
		goto L176
	} else {
		goto L186
	}
L186:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586&int32(255))+uint32(_c_F_metaphone[0]))))
	if v597&int32(4) == int32(0) {
		goto L176
	} else {
		goto L187
	}
L187:
	;
	v1034 = v232
	v1035 = v271
	goto L68
L188:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v609-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	goto L190
L190:
	;
	v625 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v625)
	goto L80
L191:
	;
	v618 = v609 - int32(32)
	goto L193
L192:
	;
	v618 = v609
	goto L193
L193:
	;
	if v618&int32(255) == int32(67) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L194
	}
L194:
	;
	goto L190
L195:
	;
	v639 = v630 - int32(32)
	goto L197
L196:
	;
	v639 = v630
	goto L197
L197:
	;
	if v639&int32(255) == int32(72) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v644 = int32(70)
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v644)
	v1034 = v629
	v1035 = v271
	goto L68
L199:
	;
	goto L200
L200:
	;
	v646 = int32(80)
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v646)
	v1034 = v629
	v1035 = v271
	goto L68
L201:
	;
	v746 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v746)
	goto L80
L202:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v692 == int32(0) {
		v699 = v692
		goto L212
	} else {
		goto L213
	}
L203:
	;
	v690 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v690)
	goto L71
L204:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v665-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v660 = v651 - int32(32)
	goto L207
L206:
	;
	v660 = v651
	goto L207
L207:
	;
	switch v660&int32(255) - int32(67) {
	case 0:
		goto L202
	default:
		goto L201
	case 5:
		goto L203
	case 6:
		goto L204
	}
L208:
	;
	v674 = v665 - int32(32)
	goto L210
L209:
	;
	v674 = v665
	goto L210
L210:
	;
	v678 = v674&int32(255) - int32(65)
	v679 = int32(0)
	if base.B2i32(v678 == v679)|base.B2i32(v678 == int32(14)) == v679 {
		goto L201
	} else {
		goto L211
	}
L211:
	;
	v687 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v687)
	goto L80
L212:
	;
	if base.Ui32((v699-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v695 == int32(0) {
		v699 = v695
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	v699 = v698
	goto L212
L215:
	;
	v708 = v699 - int32(32)
	goto L217
L216:
	;
	v708 = v699
	goto L217
L217:
	;
	if v708&int32(255) != int32(72) {
		goto L201
	} else {
		goto L218
	}
L218:
	;
	v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	if v713 == int32(0) {
		v723 = v713
		goto L219
	} else {
		goto L220
	}
L219:
	;
	if base.Ui32((v723-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v716 == int32(0) {
		v723 = v716
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if v719 == int32(0) {
		v723 = v719
		goto L219
	} else {
		goto L222
	}
L222:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+3)))
	v723 = v722
	goto L219
L223:
	;
	v732 = v723 - int32(32)
	goto L225
L224:
	;
	v732 = v723
	goto L225
L225:
	;
	if v732&int32(255) != int32(87) {
		goto L201
	} else {
		goto L226
	}
L226:
	;
	v738 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v738)
	v1034 = v232 + int32(1)
	v1035 = int32(2)
	goto L68
L227:
	;
	v791 = int32(84)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v791)
	goto L80
L228:
	;
	v787 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v787)
	goto L71
L229:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v762-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L233
	} else {
		goto L234
	}
L230:
	;
	v757 = v748 - int32(32)
	goto L232
L231:
	;
	v757 = v748
	goto L232
L232:
	;
	switch v757&int32(255) - int32(72) {
	case 0:
		goto L228
	case 1:
		goto L229
	default:
		goto L227
	}
L233:
	;
	v771 = v762 - int32(32)
	goto L235
L234:
	;
	v771 = v762
	goto L235
L235:
	;
	v775 = v771&int32(255) - int32(65)
	v776 = int32(0)
	if base.B2i32(v775 == v776)|base.B2i32(v775 == int32(14)) == v776 {
		goto L227
	} else {
		goto L236
	}
L236:
	;
	v784 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v784)
	goto L80
L237:
	;
	v805 = v796 - int32(32)
	goto L239
L238:
	;
	v805 = v796
	goto L239
L239:
	;
	if base.Ui32((v805-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v814 = v805 - int32(32)
	goto L242
L241:
	;
	v814 = v805
	goto L242
L242:
	;
	if base.Ui32(int32(25)) < base.Ui32((v814-int32(65))&int32(255)) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L243
	}
L243:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v814&int32(255))+uint32(_c_F_metaphone[0]))))
	if v825&int32(1) == int32(0) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L244
	}
L244:
	;
	v831 = int32(87)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v831)
	goto L80
L245:
	;
	v1034 = v837
	v1035 = v271
	goto L68
L246:
	;
	goto L247
L247:
	;
	v840 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v837+v37))) = uint8(v840)
	v1034 = v232 + int32(2)
	v1035 = v271
	goto L68
L248:
	;
	v853 = v844 - int32(32)
	goto L250
L249:
	;
	v853 = v844
	goto L250
L250:
	;
	if base.Ui32((v853-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v862 = v853 - int32(32)
	goto L253
L252:
	;
	v862 = v853
	goto L253
L253:
	;
	if base.Ui32(int32(25)) < base.Ui32((v862-int32(65))&int32(255)) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L254
	}
L254:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v862&int32(255))+uint32(_c_F_metaphone[0]))))
	if v873&int32(1) == int32(0) {
		v1034 = v232
		v1035 = v271
		goto L68
	} else {
		goto L255
	}
L255:
	;
	v879 = int32(89)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v879)
	goto L80
L256:
	;
	if v901&int32(255) != int32(72) {
		goto L70
	} else {
		goto L277
	}
L257:
	;
	v901 = v892 - int32(32)
	goto L259
L258:
	;
	v901 = v892
	goto L259
L259:
	;
	if base.Ui32((v901-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v910 = v901 - int32(32)
	goto L262
L261:
	;
	v910 = v901
	goto L262
L262:
	;
	if base.Ui32(int32(25)) < base.Ui32((v910-int32(65))&int32(255)) {
		goto L256
	} else {
		goto L263
	}
L263:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v910&int32(255))+uint32(_c_F_metaphone[0]))))
	if v921&int32(8) == int32(0) {
		goto L256
	} else {
		goto L264
	}
L264:
	;
	v927 = v901 & int32(255)
	if base.B2i32(v927 == int32(0))|base.B2i32(v927 != int32(73)) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	if v229 <= int32(0) {
		goto L271
	} else {
		goto L272
	}
L266:
	;
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v933-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	v942 = v933 - int32(32)
	goto L269
L268:
	;
	v942 = v933
	goto L269
L269:
	;
	if v942&int32(255) != int32(65) {
		goto L265
	} else {
		goto L270
	}
L270:
	;
	v948 = int32(88)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v948)
	goto L69
L271:
	;
	v970 = int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v970)
	goto L69
L272:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v953-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	v962 = v953 - int32(32)
	goto L275
L274:
	;
	v962 = v953
	goto L275
L275:
	;
	if v962&int32(255) != int32(83) {
		goto L271
	} else {
		goto L276
	}
L276:
	;
	v1034 = v232
	v1035 = int32(0)
	goto L68
L277:
	;
	v976 = int32(75)
	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if base.Ui32((v977-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v37))) = uint8(v1012)
	goto L71
L279:
	;
	v986 = v977 - int32(32)
	goto L281
L280:
	;
	v986 = v977
	goto L281
L281:
	;
	if v986&int32(255) == int32(82) {
		v1012 = v976
		goto L278
	} else {
		goto L282
	}
L282:
	;
	if v249 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230-int32(1)))))
	if base.Ui32((v995-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	v1012 = int32(88)
	goto L278
L286:
	;
	v1004 = v995 - int32(32)
	goto L288
L287:
	;
	v1004 = v995
	goto L288
L288:
	;
	if v1004&int32(255) == int32(83) {
		v1012 = v976
		goto L278
	} else {
		goto L289
	}
L289:
	;
	goto L285
L290:
	;
	v1057 = v1048 - int32(32)
	goto L292
L291:
	;
	v1057 = v1048
	goto L292
L292:
	;
	if v1057&int32(255) == int32(0) {
		v1066 = v1041
		goto L59
	} else {
		goto L293
	}
L293:
	;
	if v1041 < v27 {
		v229 = v1046
		v230 = v1047
		v231 = v1057
		v232 = v1041
		goto L64
	} else {
		goto L294
	}
L294:
	;
	goto L65
L295:
	;
	m.G0 = v14 + int32(32)
	return base.I64_extend_i32_u(v1088)
L296:
	;
	F_errmsg_internal(m, int32(_a_F_metaphone_1), int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L6
	} else {
		goto L297
	}
L297:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(377), int32(_a_F_metaphone_3))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L6
	} else {
		goto L298
	}
L298:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L299:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L6
	} else {
		goto L300
	}
L300:
	;
	F_errmsg(m, int32(_a_F_metaphone_4), int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L6
	} else {
		goto L301
	}
L301:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(294), int32(_a_F_metaphone_5))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L6
	} else {
		goto L302
	}
L302:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L303:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L6
	} else {
		goto L304
	}
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(255)
	F_errmsg(m, int32(_a_F_metaphone_6), v14+int32(16))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L6
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(289), int32(_a_F_metaphone_5))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L6
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L6
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(255)
	F_errmsg(m, int32(_a_F_metaphone_7), v14)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L6
	} else {
		goto L309
	}
L309:
	;
	F_errfinish(m, int32(_a_F_metaphone_2), int32(282), int32(_a_F_metaphone_5))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L6
	} else {
		goto L310
	}
L310:
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
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
				*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(88))))) = uint8(v106)
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(56))))))
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
				*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(88))))) = uint8(v35)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+(l0+int32(56))))))
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
			v79 = l0 + int32(24)
			v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v82 = v80 - int32(2)
			if v82 != 0 {
				base.MemoryCopy(m, v79, l0+int32(90), v82)
			} else {
			}
			v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
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
					*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(88))))) = uint8(v106)
					v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+(l0+int32(56))))))
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
	var v41 int32
	_ = v41
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
	var v91 int32
	_ = v91
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
	var v178 int32
	_ = v178
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
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v475 int32
	_ = v475
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	if l2 <= l1 {
		v475 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v21 + int32(32)
	return v475
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
	v41 = l1
	v47 = v5
	goto L7
L5:
	;
	v87 = v5
	v89 = v5
	v91 = l1
	goto L6
L6:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v24+v91<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v105 <= l3 {
		v119 = v89
		goto L3
	} else {
		goto L17
	}
L7:
	;
	v53 = v24 + v41<<(uint(int32(2))%32)
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
	v77 = v41 + v76
	v79 = v47 + v76
	if v79 != v28&int32(-2) {
		v37 = v74
		v39 = v75
		v41 = v77
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
	v91 = v77
	goto L6
L17:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v104)+8)))
	v119 = v89 + base.B2i32(v108 != v87&int32(255))
	goto L3
L18:
	;
	v475 = int32(0)
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
	v178 = l1
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
	v189 = v178 << (uint(int32(2)) % 32)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	if v193 <= l3 {
		v449 = v171
		v452 = v176
		v455 = v182
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v462 = F_mkSPNode(m, l0, v449, l2, v166)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L25
	} else {
		goto L92
	}
L34:
	;
	v460 = v178 + int32(1)
	if v460 != l2 {
		v171 = v449
		v176 = v452
		v178 = v460
		v182 = v455
		goto L32
	} else {
		goto L91
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
		v449 = v212
		v452 = v213
		v455 = v214
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
	v202 = F_mkSPNode(m, l0, v171, v178, v166)
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
	v212 = v178
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
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v405 | int32(256)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v416+int32(base.Ui32(v405)>>(uint(int32(11))%32))&int32(_a_F_mkSPNode_1))))
	v423 = F_getCompoundAffixFlagValue(m, l0, v422)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L25
	} else {
		goto L83
	}
L44:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238+v231<<(uint(int32(2))%32))))
	v243 = F_getCompoundAffixFlagValue(m, l0, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L25
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v405 = v220&int32(_a_F_mkSPNode_2) | v231<<(uint(int32(13))%32)
	v412 = int32(0)
	goto L43
L47:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v248+v189)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v255 = int32(base.Ui32(v253) >> (uint(int32(13)) % 32))
	v257 = v255 << (uint(int32(2)) % 32)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v252+v257)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v260 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v405 = v390&int32(_a_F_mkSPNode_2) | v387<<(uint(int32(13))%32)
	v412 = v243&int32(base.Ui32(v215)>>(uint(int32(9))%32)) ^ int32(1)
	goto L43
L49:
	;
	v387 = v251
	v390 = v253
	goto L48
L50:
	;
	goto L51
L51:
	;
	v264 = v251 << (uint(int32(2)) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v252+v264)))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266))))
	if v267 == int32(0) {
		v387 = v255
		v390 = v253
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v270 <= v271+int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v270 << (uint(int32(1)) % 32)
	v280 = F_repalloc(m, v252, v270<<(uint(int32(3))%32))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L25
	} else {
		goto L56
	}
L54:
	;
	v288 = v252
	v289 = v259
	v290 = v266
	v291 = v271
	goto L55
L55:
	;
	v292 = int32(2)
	v294 = v291<<(uint(v292)%32) + v288
	v295 = F_strlen(m, v289)
	mBase = m.M
	v296 = F_strlen(m, v290)
	mBase = m.M
	v297 = v295 + v296
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v298 == v292 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v280
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v264)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v280+v257)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v288 = v280
	v289 = v286
	v290 = v284
	v291 = v287
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v294)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294))) = v375
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v381 + int32(1)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v387 = v381
	v390 = v385
	goto L48
L58:
	;
	v302 = v297 + int32(2)
	if base.Ui32(int32(1025)) <= base.Ui32(v302) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v338 = v297 + int32(1)
	if base.Ui32(int32(1025)) <= base.Ui32(v338) {
		goto L73
	} else {
		goto L74
	}
L61:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v327+v257)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327+v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v329
	v335 = F_pg_sprintf(m, v324, int32(_a_F_mkSPNode_3), v21)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L25
	} else {
		goto L71
	}
L62:
	;
	v305 = F_palloc0(m, v302)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L25
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v310 = (v297 + int32(9)) & int32(4088)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v310) <= base.Ui32(v311) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v324 = v305
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v318 - v310
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v319 + v310
	v324 = v319
	goto L61
L67:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v318 = v311
	v319 = v313
	goto L66
L68:
	;
	goto L69
L69:
	;
	v314 = int32(_a_F_mkSPNode_0)
	v316 = F_palloc0(m, v314)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	v318 = v314
	v319 = v316
	goto L66
L71:
	;
	v375 = v324
	goto L57
L72:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v363+v257)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v264)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v365
	v373 = F_pg_sprintf(m, v360, int32(_a_F_mkSPNode_4), v21+int32(16))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L25
	} else {
		goto L82
	}
L73:
	;
	v341 = F_palloc0(m, v338)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L25
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v346 = (v297 + int32(8)) & int32(4088)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if base.Ui32(v346) <= base.Ui32(v347) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v360 = v341
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v354 - v346
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v355 + v346
	v360 = v355
	goto L72
L78:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v354 = v347
	v355 = v349
	goto L77
L79:
	;
	goto L80
L80:
	;
	v350 = int32(_a_F_mkSPNode_0)
	v352 = F_palloc0(m, v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v354 = v350
	v355 = v352
	goto L77
L82:
	;
	v375 = v360
	goto L57
L83:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v431 = v423 & int32(15)
	v434 = v425&int32(-7681) | v431<<(uint(int32(9))%32)
	if base.Ui32(v431) < base.Ui32(int32(2)) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v437 = v425 | int32(_a_F_mkSPNode_5)
	goto L86
L85:
	;
	v437 = v434
	goto L86
L86:
	;
	if v423&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v440 = v437
	goto L89
L88:
	;
	v440 = v434
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v440
	if v412&int32(1) == int32(0) {
		v449 = v212
		v452 = v213
		v455 = v214
		goto L34
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v440 & int32(-513)
	v449 = v212
	v452 = v213
	v455 = v214
	goto L34
L91:
	;
	goto L33
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v452)+4)) = v462
	v475 = v163
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
func F_mul_d_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(1605), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_mul_size_error(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		F_errcode(m, int32(261))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
			F_errmsg(m, int32(_a_F_mul_size_error_0), v6)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_mul_size_error_1), int32(1774), int32(_a_F_mul_size_error_2))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
