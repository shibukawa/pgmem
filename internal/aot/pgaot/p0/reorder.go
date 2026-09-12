package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReorderBufferCleanupSerializedTXNs(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	v5 = m.G0
	v7 = v5 - int32(2208)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(90937)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l0
	v17 = F_pg_sprintf(m, v7+int32(48), int32(187320), v7+int32(32))
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
	v25 = F___fstatat(m, int32(-100), v7+int32(48), v7+int32(2112), int32(256))
	mBase = m.M
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L4:
	;
	m.G0 = v7 + int32(2208)
	return
L5:
	;
	if v25 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+2116))
	if v28&int32(61440) != int32(16384) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v35 = F_AllocateDir(m, v7+int32(48))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v40 = F_ReadDirExtended(m, v35, v7+int32(48), int32(17))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = v40
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_FreeDir(m, v35)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+19)))
	if v46 != int32(120) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v75 = F_ReadDirExtended(m, v35, v7+int32(48), int32(17))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L18:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+20)))
	if v49 != int32(105) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+21)))
	if v52 != int32(100) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v44 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(90937)
	v67 = F_pg_snprintf(m, v7+int32(48), int32(2060), int32(187272), v7+int32(16))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = F_unlink(m, v7+int32(48))
	mBase = m.M
	if v71 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	goto L17
L23:
	;
	if v75 != 0 {
		v44 = v75
		goto L15
	} else {
		goto L24
	}
L24:
	;
	goto L16
L25:
	;
	goto L4
L26:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(90937)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v7 + int32(48)
	F_errmsg(m, int32(309640), v7)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(518586), int32(4877), int32(185099))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferCommitChild(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v14 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v25 = F_hash_search(m, v19, v11+int32(12), int32(0), v11+int32(11))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
			if v27 == int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
				m.G0 = v11 + int32(16)
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
				if v35 == int32(0) {
					m.G0 = v11 + int32(16)
					return
				} else {
					v40 = v35
					*(*int64)(unsafe.Add(mBase, uint32(v40)+32)) = l4
					*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = l3
					F_ReorderBufferAssignChild(m, l0, l1, l2, int64(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v11 + int32(16)
						return
					}
				}
			}
		}
	} else {
		if l2 != v14 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v25 = F_hash_search(m, v19, v11+int32(12), int32(0), v11+int32(11))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
				if v27 == int32(0) {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v30
					m.G0 = v11 + int32(16)
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v34
					if v35 == int32(0) {
						m.G0 = v11 + int32(16)
						return
					} else {
						v40 = v35
						*(*int64)(unsafe.Add(mBase, uint32(v40)+32)) = l4
						*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = l3
						F_ReorderBufferAssignChild(m, l0, l1, l2, int64(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v11 + int32(16)
							return
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			if v18 != 0 {
				v40 = v18
				*(*int64)(unsafe.Add(mBase, uint32(v40)+32)) = l4
				*(*int64)(unsafe.Add(mBase, uint32(v40)+24)) = l3
				F_ReorderBufferAssignChild(m, l0, l1, l2, int64(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		}
	}
}
func F_ReorderBufferRestoreChanges(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v101 int32
	_ = v101
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v138 int32
	_ = v138
	var v143 int64
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v169 int64
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
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
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int64
	_ = v231
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int32
	_ = v243
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
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v295 int32
	_ = v295
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
	var v322 int64
	_ = v322
	var v324 int64
	_ = v324
	var v326 int64
	_ = v326
	var v328 int64
	_ = v328
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int64
	_ = v334
	var v338 int64
	_ = v338
	var v341 int32
	_ = v341
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
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v395 int64
	_ = v395
	var v397 int64
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
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
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
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
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v491 int64
	_ = v491
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(1104)
	m.G0 = v22
	v25 = l1 + int32(128)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v26 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = int64(0)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	v85 = int64(*(*int32)(unsafe.Add(mBase, _consts[164])))
	v86 = base.I64_div_u_s(v83, v85)
	v101 = v5
	goto L13
L2:
	;
	if v26 == v25 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = v26
	goto L4
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = v52
	F_ReorderBufferFreeChange(m, l0, v34-int32(52), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	if v50 != v25 {
		v34 = v50
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L6
	} else {
		goto L148
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L6
	} else {
		goto L144
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L6
	} else {
		goto L140
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L6
	} else {
		goto L136
	}
L13:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui64(v108) <= base.Ui64(v86) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	m.G0 = v22 + int32(1104)
	return v594
L15:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v111 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v594 = v101
	goto L17
L17:
	;
	goto L14
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v114 != int32(-1) {
		v190 = v114
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	if base.Ui32(v581) < base.Ui32(int32(4096)) {
		v101 = v581
		goto L13
	} else {
		goto L135
	}
L23:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v193 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	if v119 == int64(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v122 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	v124 = base.I64_div_u_s(v122, base.I64_extend_i32_s(v118))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v124
	v127 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v128 = v127
	v129 = v124
	goto L27
L26:
	;
	v128 = v118
	v129 = v119
	goto L27
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v132 = v129 * base.I64_extend_i32_s(v128)
	*(*uint32)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = uint32(v132)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(90937)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v130
	v138 = *(*int32)(unsafe.Add(mBase, _consts[510]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v138 + int32(24)
	v143 = int64(base.Ui64(v132) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+60)) = uint32(v143)
	v151 = F_pg_snprintf(m, v22+int32(80), int32(1024), int32(318071), v22+int32(48))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v156 = F_PathNameOpenFile(m, v22+int32(80), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v156
	if int32(0) <= v156 {
		v190 = v156
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v164 == int32(44) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v169 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v169 + int64(1)
	v581 = v101
	goto L22
L32:
	;
	goto L33
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v22 + int32(80)
	F_errmsg(m, int32(312266), v22)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(518586), int32(4574), int32(179692))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v213
	v222 = F_FileReadV(m, v212, v22+int32(80), int32(1), v214, int32(167772203))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L46
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v207
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v212 = v211
	v213 = v207
	goto L38
L40:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v198 = F_MemoryContextAlloc(m, v196, int32(72))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if base.Ui32(int32(71)) < base.Ui32(v193) {
		v212 = v190
		v213 = v200
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v207 = v198
	goto L39
L44:
	;
	v204 = F_repalloc(m, v200, int32(72))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v207 = v204
	goto L39
L46:
	;
	if v222 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	F_FileClose(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v222 < int32(0) {
		goto L12
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v231 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v231 + int64(1)
	v581 = v101
	goto L22
L51:
	;
	if v222 != int32(72) {
		goto L11
	} else {
		goto L52
	}
L52:
	;
	v239 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v241 = v239 + int64(72)
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v241
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v246 = v244 + int32(72)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v247 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v265 = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v262 - v265
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v261 + v265
	v275 = F_FileReadV(m, v264, v22+int32(80), int32(1), v263, int32(167772203))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L61
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v256
	v259 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v261 = v256
	v262 = v260
	v263 = v259
	goto L53
L55:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v251 = F_MemoryContextAlloc(m, v250, v246)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if base.Ui32(v246) <= base.Ui32(v247) {
		v261 = v243
		v262 = v244
		v263 = v241
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v256 = v251
	goto L54
L59:
	;
	v254 = F_repalloc(m, v243, v246)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v256 = v254
	goto L54
L61:
	;
	if v275 < int32(0) {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	if v275 != v279-int32(72) {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	v283 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v283 + base.I64_extend_i32_u(v275)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v290 = F_MemoryContextAlloc(m, v288, int32(64))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v292 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v290))) = v292
	v295 = v290 + int32(56)
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v292
	v299 = v290 + int32(48)
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = v292
	v303 = v290 + int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v303))) = v292
	v307 = v290 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v292
	v311 = v290 + int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v311))) = v292
	v315 = v290 + int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v292
	v319 = v290 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v292
	v322 = *(*int64)(unsafe.Add(mBase, uint32(v287)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v290))) = v322
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v287)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = v324
	v326 = *(*int64)(unsafe.Add(mBase, uint32(v287)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v287)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v311))) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v287)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v330
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v287)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v303))) = v332
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v287)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = v334
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v287-int32(-64))))
	*(*int64)(unsafe.Add(mBase, uint32(v295))) = v338
	v341 = v287 + int32(72)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	switch v342 {
	case 0, 1, 2, 8:
		goto L70
	case 3:
		goto L69
	case 4:
		goto L68
	case 5:
		goto L67
	default:
		goto L65
	case 11:
		goto L66
	}
L65:
	;
	v480 = v290 + int32(52)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v481 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L66:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v466 = F_MemoryContextAlloc(m, v462, v463<<(uint(int32(2))%32))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L105
	}
L67:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v287)+96))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v287)+88))
	v446 = (v440+v441)<<(uint(int32(2))%32) + int32(72)
	v447 = F_MemoryContextAllocZero(m, v439, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L6
	} else {
		goto L100
	}
L68:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v433 = v431 << (uint(int32(4)) % 32)
	v434 = F_MemoryContextAlloc(m, v430, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L95
	}
L69:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v287)+72))
	v411 = F_MemoryContextAlloc(m, v409, v410)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L85
	}
L70:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	if v343 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v348 = F_MemoryContextAlloc(m, v344, v345+int32(24))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	v374 = v341
	goto L73
L73:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	if v378 == int32(0) {
		goto L65
	} else {
		goto L79
	}
L74:
	;
	v351 = v348 + int32(16)
	v352 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v348 + v352
	*(*int32)(unsafe.Add(mBase, uint32(v290)+36)) = v348
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v341)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v356
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v341)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v348)+8)) = v358
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v341)))
	*(*int64)(unsafe.Add(mBase, uint32(v348))) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v362)+16)) = v362 + v352
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	v369 = v287 + int32(92)
	if v345 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v374 = v369 + v345
	goto L73
L76:
	;
	v370 = F__emscripten_memcpy_bulkmem(m, v367, v369, v345)
	mBase = m.M
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v385 = F_MemoryContextAlloc(m, v381, v382+int32(24))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	v388 = v385 + int32(16)
	v389 = int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v385 + v389
	*(*int32)(unsafe.Add(mBase, uint32(v290)+40)) = v385
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v374)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = v393
	v395 = *(*int64)(unsafe.Add(mBase, uint32(v374)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v385)+8)) = v395
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v374)))
	*(*int64)(unsafe.Add(mBase, uint32(v385))) = v397
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v399)+16)) = v399 + v389
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	if v382 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L65
L82:
	;
	v407 = F__emscripten_memcpy_bulkmem(m, v404, v374+int32(20), v382)
	mBase = m.M
	goto L84
L83:
	;
	goto L84
L84:
	;
	goto L81
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+20)) = v411
	v415 = v287 + int32(76)
	if v410 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v418 = v415 + v410
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+24)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v422 = F_MemoryContextAlloc(m, v421, v419)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L6
	} else {
		goto L90
	}
L87:
	;
	v416 = F__emscripten_memcpy_bulkmem(m, v411, v415, v410)
	mBase = m.M
	goto L89
L88:
	;
	goto L89
L89:
	;
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v422
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
	if v427 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L65
L92:
	;
	v428 = F__emscripten_memcpy_bulkmem(m, v422, v418+int32(4), v427)
	mBase = m.M
	goto L94
L93:
	;
	goto L94
L94:
	;
	goto L91
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+24)) = v434
	if v433 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L65
L97:
	;
	v437 = F__emscripten_memcpy_bulkmem(m, v434, v341, v433)
	mBase = m.M
	goto L99
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+20)) = v447
	if v446 != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v453 = v451 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v451)+12)) = v453
	v455 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+30)) = uint8(v455)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v451)+20)) = v453 + v457<<(uint(int32(2))%32)
	goto L65
L102:
	;
	v450 = F__emscripten_memcpy_bulkmem(m, v447, v341, v446)
	mBase = m.M
	v451 = v450
	goto L104
L103:
	;
	v451 = v447
	goto L104
L104:
	;
	goto L101
L105:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v470 = v468 << (uint(int32(2)) % 32)
	if v470 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+28)) = v472
	goto L65
L107:
	;
	v471 = F__emscripten_memcpy_bulkmem(m, v466, v341, v470)
	mBase = m.M
	v472 = v471
	goto L109
L108:
	;
	v472 = v466
	goto L109
L109:
	;
	goto L106
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v25
	goto L112
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v290)+56)) = v25
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v290)+52)) = v487
	*(*int32)(unsafe.Add(mBase, uint32(v487)+4)) = v480
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v480
	v491 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+120)) = v491 + int64(1)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v290)+8))
	switch v498 {
	case 0, 1, 2, 8:
		goto L120
	case 3:
		goto L119
	case 4:
		goto L118
	case 5:
		goto L117
	default:
		v537 = int32(64)
		goto L115
	case 11:
		goto L116
	}
L113:
	;
	v581 = v101 + int32(1)
	goto L22
L114:
	;
	if v542 == int32(0) {
		goto L113
	} else {
		goto L125
	}
L115:
	;
	v542 = v537
	goto L114
L116:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v537 = v531<<(uint(int32(2))%32) - int32(-64)
	goto L115
L117:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v523)+24))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v523)+16))
	v542 = (v524+v525)<<(uint(int32(2))%32) + int32(136)
	goto L114
L118:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v542 = v518<<(uint(int32(4))%32) - int32(-64)
	goto L114
L119:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v290)+20))
	v513 = F_strlen(m, v512)
	mBase = m.M
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v290)+24))
	v542 = v513 + v514 + int32(73)
	goto L114
L120:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	if v500 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
	v505 = v501 + int32(84)
	goto L123
L122:
	;
	v505 = int32(64)
	goto L123
L123:
	;
	if v499 == int32(0) {
		v537 = v505
		goto L115
	} else {
		goto L124
	}
L124:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	v542 = v505 + v508 + int32(20)
	goto L114
L125:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v290)+8))
	if v545 == int32(7) {
		goto L113
	} else {
		goto L126
	}
L126:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v548)+216)) = v549 + v542
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548)+40))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v553 + v542
	if v552 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v556 = v552
	goto L129
L128:
	;
	v556 = v548
	goto L129
L129:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v556)+220)) = v557 + v542
	if v549 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_remove(m, v560, v548+int32(204))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_pairingheap_add(m, v565, v548+int32(204))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L6
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	goto L113
L135:
	;
	v594 = v581
	goto L17
L136:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(308371), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(518586), int32(4598), int32(179692))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v222
	F_errmsg(m, int32(168703), v22+int32(32))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(518586), int32(4604), int32(179692))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L6
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(308371), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L6
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(518586), int32(4623), int32(179692))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L6
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L6
	} else {
		goto L149
	}
L149:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v275
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v658 - int32(72)
	F_errmsg(m, int32(168703), v22+int32(16))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(518586), int32(4629), int32(179692))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReorderBufferSetRestartPoint(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+136)) = l1
	return
}
