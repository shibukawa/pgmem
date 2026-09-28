package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_InvalidateConstraintCacheCallBack(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[0]))
	if base.B2i32(v8 == v4)|base.B2i32(v8 == int32(_a_F_InvalidateConstraintCacheCallBack_0)) == v4 {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
		if base.Ui32(v18) <= base.Ui32(int32(1000)) {
			v21 = l2
		} else {
			v21 = int32(0)
		}
		if l1 != int32(3) {
			v25 = v21
		} else {
			v25 = int32(0)
		}
		v29 = v8
		for {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			if v25 == int32(0) {
				v45 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29-int32(692)))) = uint8(v45)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
				if v47 != 0 {
					v48 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v47)+2244)) = v49
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2])) = v52
					*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(0)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
					v57 = v56
				} else {
					v57 = v32
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v57
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				*(*int32)(unsafe.Add(mBase, uint32(v57))) = v60
				v62 = int32(_a_F_InvalidateConstraintCacheCallBack_2)
				v64 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v64 - int32(1)
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v29-int32(684))))
				if v37 == v25 {
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v29-int32(692)))) = uint8(v45)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
					if v47 != 0 {
						v48 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
						v49 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2]))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+2244)) = v49
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2])) = v52
						*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(0)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
						v57 = v56
					} else {
						v57 = v32
					}
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v57
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					*(*int32)(unsafe.Add(mBase, uint32(v57))) = v60
					v62 = int32(_a_F_InvalidateConstraintCacheCallBack_2)
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
					*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v64 - int32(1)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v29-int32(680))))
					if v41 != v25 {
					} else {
						v45 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v29-int32(692)))) = uint8(v45)
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
						if v47 != 0 {
							v48 = int32(_a_F_InvalidateConstraintCacheCallBack_1)
							v49 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+2244)) = v49
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[2])) = v52
							*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(0)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
							v57 = v56
						} else {
							v57 = v32
						}
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v57
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v60
						v62 = int32(_a_F_InvalidateConstraintCacheCallBack_2)
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1]))
						*(*int32)(unsafe.Add(mBase, _c_F_InvalidateConstraintCacheCallBack[1])) = v64 - int32(1)
					}
				}
			}
			if v32 != int32(_a_F_InvalidateConstraintCacheCallBack_0) {
				v29 = v32
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_InvalidateObsoleteReplicationSlots(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v164 int64
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v185 int64
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v259 int64
	_ = v259
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v279 int64
	_ = v279
	var v281 int64
	_ = v281
	var v286 int32
	_ = v286
	var v288 int64
	_ = v288
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v333 int64
	_ = v333
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int64
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
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
	var v382 int64
	_ = v382
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v409 int64
	_ = v409
	var v411 int64
	_ = v411
	var v413 int64
	_ = v413
	var v415 int64
	_ = v415
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v423 int64
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int64
	_ = v538
	var v540 int64
	_ = v540
	var v542 int64
	_ = v542
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v548 int64
	_ = v548
	var v550 int64
	_ = v550
	var v552 int64
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v732 int32
	_ = v732
	var v747 int32
	_ = v747
	v5 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(1232)
	m.G0 = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[1]))
	if v31|v33 == v5 {
		v747 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v28 + int32(1232)
	return v747
L2:
	;
	v38 = int64(*(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[2])))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v44 = F_LWLockAcquire(m, v40+int32(_a_F_InvalidateObsoleteReplicationSlots_0), int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[1]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if v49+v51 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v56+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v61 = l1 * v38
	v69 = l0 & int32(8)
	v71 = base.B2i32(base.Ui32(l3) < base.Ui32(int32(3)))
	v83 = v5
	v84 = v5
	v85 = v5
	v86 = v5
	goto L9
L8:
	;
	v747 = v5
	goto L1
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[4]))
	v101 = v98 + v85*int32(296)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v102 != int32(1) {
		v627 = v83
		v628 = v84
		v630 = v86
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v716+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L3
	} else {
		goto L141
	}
L11:
	;
	goto L10
L12:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v680 = F_LWLockAcquire(m, v676+int32(_a_F_InvalidateObsoleteReplicationSlots_0), int32(1))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L3
	} else {
		goto L139
	}
L13:
	;
	v642 = v85 + int32(1)
	v644 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[1]))
	v646 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if v642 < v644+v646 {
		v83 = v627
		v84 = v628
		v85 = v642
		v86 = v630
		goto L9
	} else {
		goto L138
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	if v105 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = v127
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v129 == v127 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[5])))
	if v109&int32(1) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v116 = base.AtomicRmwXchg32(m, v101, int32(0), int32(1))
	if v116 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v101, int32(_a_F_InvalidateObsoleteReplicationSlots_1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L3
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v101)+112))
	v121 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v121))
	v627 = v83 | base.B2i32(v120 == v121)
	v628 = v84
	v630 = v86
	goto L13
L21:
	;
	goto L20
L22:
	;
	v600 = base.AtomicRmwXchg32(m, v101, int32(0), int32(1))
	if v600 != 0 {
		goto L130
	} else {
		goto L131
	}
L23:
	;
	v573 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v134 = v101 + int32(224)
	v136 = v101 + int32(24)
	v138 = int32(0)
	v141 = v138
	v147 = v138
	v150 = v138
	v164 = int64(0)
	goto L28
L26:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	v661 = base.B2i32(v568 != int32(0)) | v84
	v663 = int32(1)
	goto L12
L27:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v497+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L3
	} else {
		goto L120
	}
L28:
	;
	if v69 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v487+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L118
	}
L30:
	;
	v170 = m.G0
	v171 = int32(16)
	v172 = v170 - v171
	m.G0 = v172
	F_gettimeofday(m, v172)
	mBase = m.M
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v172)))
	v176 = int64(*(*int32)(unsafe.Add(mBase, uint32(v172)+8)))
	m.G0 = v172 + v171
	goto L33
L31:
	;
	v185 = int64(0)
	goto L32
L32:
	;
	v188 = base.AtomicRmwXchg32(m, v101, int32(0), int32(1))
	if v188 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v185 = v176 + v175*int64(1000000) - int64(946684800000000)
	goto L32
L34:
	;
	F_s_lock(m, v101, int32(_a_F_InvalidateObsoleteReplicationSlots_1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v101)+112))
	if v192 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	goto L36
L38:
	;
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v136)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+200)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v136)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+192)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v136)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+184)) = v331
	v333 = *(*int64)(unsafe.Add(mBase, uint32(v136)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+176)) = v333
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v136)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+168)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v136)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+160)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v136)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+152)) = v339
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+144)) = v341
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v101)+8))
	if v343 == int32(-1) {
		goto L87
	} else {
		goto L88
	}
L39:
	;
	v320 = int32(0)
	v322 = v320
	v323 = v319
	v325 = v164
	v326 = v320
	goto L38
L40:
	;
	v319 = int32(2)
	goto L39
L41:
	;
	v296 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v296))
	if v141&int32(1) != 0 {
		goto L80
	} else {
		goto L81
	}
L42:
	;
	v193 = int32(0)
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v101)+104))
	if base.B2i32(l0&int32(1) == v193)|base.B2i32(v195 == int64(0))|base.B2i32(base.Ui64(v61) <= base.Ui64(v195)) == v193 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v203 = int32(1)
	v322 = v203
	v323 = v203
	v325 = v164
	v326 = int32(0)
	goto L38
L44:
	;
	goto L45
L45:
	;
	if l0&int32(2) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if l0&int32(4) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L47:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	v209 = int32(0)
	if base.B2i32(v208 == v209)|base.B2i32(base.B2i32(l2 == v209)|base.B2i32(l2 == v208) == v209) != 0 {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	if v219 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if v218 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L50:
	;
	if v71|base.B2i32(base.Ui32(v219) < base.Ui32(int32(3))) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if int32(0) < v219-l3 {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(v219) <= base.Ui32(l3) {
		goto L40
	} else {
		goto L55
	}
L54:
	;
	goto L40
L55:
	;
	goto L49
L56:
	;
	if v71|base.B2i32(base.Ui32(v218) < base.Ui32(int32(3))) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v218-l3 <= int32(0) {
		goto L40
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(v218) <= base.Ui32(l3) {
		goto L40
	} else {
		goto L61
	}
L60:
	;
	goto L46
L61:
	;
	goto L46
L62:
	;
	if v69 == int32(0) {
		goto L41
	} else {
		goto L65
	}
L63:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	if v246 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v319 = int32(4)
	goto L39
L65:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[6]))
	if base.B2i32(v253 == int32(0))|base.B2i32(v195 == int64(0)) != 0 {
		goto L41
	} else {
		goto L66
	}
L66:
	;
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	if v259 <= int64(0) {
		goto L41
	} else {
		goto L67
	}
L67:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[7])))
	if v264 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v274 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[8]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+308))
	v272 = base.B2i32(v270 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[7])) = uint8(v272)
	v274 = v272
	goto L71
L70:
	;
	v274 = int32(0)
	goto L71
L71:
	;
	goto L68
L72:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+201)))
	if v275 != 0 {
		goto L41
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[6]))
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	v279 = v185 - v278
	v281 = base.I64_div_u_s(v279, int64(1000000))
	if int64(0) < v279 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v286 = base.I32_wrap_i64(v281)
	goto L78
L77:
	;
	v286 = int32(0)
	goto L78
L78:
	;
	if v286 < v277 {
		goto L41
	} else {
		goto L79
	}
L79:
	;
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v101)+272))
	v322 = int32(0)
	v323 = int32(8)
	v325 = v288
	v326 = int32(1)
	goto L38
L80:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v302+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v308 = int32(0)
	if v150 == v308 {
		v573 = v308
		goto L22
	} else {
		goto L85
	}
L83:
	;
	if v150 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	v573 = int32(1)
	goto L22
L85:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	v627 = v83
	v628 = base.B2i32(v311 != int32(0)) | v84
	v630 = int32(1)
	goto L13
L86:
	;
	v373 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v373))
	if v326 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9])) = v101
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+112)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v349
	v352 = int32(1)
	v353 = int32(0)
	if v322 == v353 {
		v371 = v352
		v372 = v353
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[11]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v343*int32(768))+12))
	v371 = v150
	v372 = v370
	goto L86
L90:
	;
	v356 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+280)) = v356
	*(*int64)(unsafe.Add(mBase, uint32(v101)+104)) = v356
	v360 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v360))
	v494 = v360
	v495 = v352
	goto L27
L91:
	;
	v382 = v185 - v325
	if v382 <= int64(0) {
		goto L95
	} else {
		goto L96
	}
L92:
	;
	goto L93
L93:
	;
	if v343 == int32(-1) {
		v494 = v372
		v495 = v371
		goto L27
	} else {
		goto L98
	}
L94:
	;
	goto L93
L95:
	;
	v394 = int32(0)
	v395 = int32(0)
	goto L97
L96:
	;
	v386 = int64(1000000)
	v387 = base.I64_div_u_s(v382, v386)
	v394 = base.I32_wrap_i64(v387)
	v395 = base.I32_wrap_i64(v382 - v387*v386)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(140)))) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v28+int32(208)))) = v395
	goto L94
L98:
	;
	F_ConditionVariablePrepareToSleep(m, v134)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v403+int32(_a_F_InvalidateObsoleteReplicationSlots_0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	if v372 != v147 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v28)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+72)) = v409
	v411 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+80)) = v411
	v413 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+88)) = v413
	v415 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+96)) = v415
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+104)) = v417
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+112)) = v419
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v28)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+120)) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+128)) = v423
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	F_ReportSlotInvalidation(m, v323, int32(1), v372, v28+int32(72), v195, v61, l3, v428)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L104
	}
L102:
	;
	v471 = v147
	goto L103
L103:
	;
	F_ConditionVariableSleep(m, v134, int32(134217778))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L3
	} else {
		goto L115
	}
L104:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[12]))
	if v432 == int32(13) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v471 = v372
	goto L103
L106:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[11]))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
	v440 = v437 + v343*int32(768)
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v446 = F_LWLockAcquire(m, v442+int32(512), int32(1))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L3
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v469 = F_pgmem_kill(m, v372, int32(15))
	mBase = m.M
	goto L105
L109:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	if v372 == v448 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v452 = base.AtomicRmwOr32(m, v440, int32(340), int32(16))
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[11]))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v459 = base.I32_div_s(v440-v456, int32(768))
	v460 = F_SendProcSignal(m, v372, int32(9), v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L3
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	F_LWLockRelease(m, v463+int32(512))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L3
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	goto L105
L115:
	;
	v475 = int32(1)
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[3]))
	v481 = F_LWLockAcquire(m, v477+int32(_a_F_InvalidateObsoleteReplicationSlots_0), v475)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+140)) = int32(0)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+4)))
	if v485 != 0 {
		v141 = v475
		v147 = v471
		v150 = v371
		v164 = v325
		goto L28
	} else {
		goto L117
	}
L117:
	;
	goto L29
L118:
	;
	if v371 == int32(0) {
		v573 = v475
		goto L22
	} else {
		goto L119
	}
L119:
	;
	goto L26
L120:
	;
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	v506 = base.AtomicRmwXchg32(m, v503, int32(0), int32(1))
	if v506 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_s_lock(m, v503, int32(_a_F_InvalidateObsoleteReplicationSlots_1))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L3
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v510 = int32(_a_F_InvalidateObsoleteReplicationSlots_2)
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	v512 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v511)+12)) = uint16(v512)
	v514 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v503))), uint32(v514))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = int32(_a_F_InvalidateObsoleteReplicationSlots_3)
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v520 + int32(24)
	v525 = v28 + int32(208)
	v529 = F_pg_sprintf(m, v525, int32(_a_F_InvalidateObsoleteReplicationSlots_4), v28-int32(-64))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L3
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[9]))
	F_SaveSlotToPath(m, v532, v525, int32(21))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v28)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v538
	v540 = *(*int64)(unsafe.Add(mBase, uint32(v28)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v540
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v28)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v542
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v28)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+24)) = v544
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v28)+176))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+32)) = v546
	v548 = *(*int64)(unsafe.Add(mBase, uint32(v28)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = v548
	v550 = *(*int64)(unsafe.Add(mBase, uint32(v28)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+48)) = v550
	v552 = *(*int64)(unsafe.Add(mBase, uint32(v28)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+56)) = v552
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v28)+140))
	F_ReportSlotInvalidation(m, v323, int32(0), v494, v28, v195, v61, l3, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	if v495 == int32(0) {
		v573 = int32(1)
		goto L22
	} else {
		goto L129
	}
L129:
	;
	goto L26
L130:
	;
	F_s_lock(m, v101, int32(_a_F_InvalidateObsoleteReplicationSlots_1))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L3
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v101)+88))
	if v604 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	goto L132
L134:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v101)+112))
	v607 = v605
	goto L136
L135:
	;
	v607 = int32(1)
	goto L136
L136:
	;
	v608 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v101))), uint32(v608))
	if v573 != 0 {
		v661 = v84
		v663 = v86
		goto L12
	} else {
		goto L137
	}
L137:
	;
	v627 = base.B2i32(v607 == v608) | v83&int32(1)
	v628 = v84
	v630 = v86
	goto L13
L138:
	;
	v701 = v627
	v702 = v628
	v704 = v630
	goto L11
L139:
	;
	v682 = int32(0)
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[1]))
	v686 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateObsoleteReplicationSlots[0]))
	if v682 < v684+v686 {
		v83 = v682
		v84 = v661
		v85 = int32(0)
		v86 = v663
		goto L9
	} else {
		goto L140
	}
L140:
	;
	v701 = v682
	v702 = v661
	v704 = v663
	goto L11
L141:
	;
	if v704 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L3
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if (v702^int32(-1)|v701)&int32(1) != 0 {
		v747 = v704
		goto L1
	} else {
		goto L147
	}
L145:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L3
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	F_RequestDisableLogicalDecoding(m)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v747 = v704
	goto L1
}
func F_InvalidateOprCacheCallBack(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprCacheCallBack[0]))
	F_hash_seq_init(m, v6+int32(12), v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L4
L3:
	;
	m.G0 = v6 + int32(32)
	return
L4:
	;
	v19 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprCacheCallBack[0]))
	v27 = F_hash_search(m, v24, v19, int32(2), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v27 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateOprCacheCallBack_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_InvalidateOprCacheCallBack_1), int32(1125), int32(_a_F_InvalidateOprCacheCallBack_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateOprProofCacheCallBack(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(12)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateOprProofCacheCallBack[0]))
	F_hash_seq_init(m, v9, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_hash_seq_search(m, v9)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	m.G0 = v6 + int32(32)
	return
L7:
	;
	v19 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)) = uint16(v19)
	v23 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v23 != 0 {
		v18 = v23
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
}
func F_InvalidatePublicationRels(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= int32(4095) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_CacheInvalidateRelcacheAll(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L12
	}
L6:
	;
	v12 = v2
	goto L7
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+v12<<(uint(int32(2))%32))))
	F_CacheInvalidateRelcacheByRelid(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L1
L9:
	;
	return
L10:
	;
	v21 = v12 + int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 < v22 {
		v12 = v21
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L1
}
func F_InvalidateSystemCaches(m *base.Module) {
	var v2 int32
	_ = v2
	F_InvalidateSystemCachesExtended(m)
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		return
	}
}
