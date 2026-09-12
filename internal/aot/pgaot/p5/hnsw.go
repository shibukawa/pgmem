package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v5 = m.T0[v4].(func(*base.Module, int32, int32) int32)(m, l1, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v5
		}
	} else {
		v10 = F_palloc(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_HnswInitNeighborArray(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = F_mul_size(m, int32(12), l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_add_size(m, int32(8), v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if l1 != 0 {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, v9, v11)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					v17 = v13
					v18 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v18)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
					return v17
				}
			} else {
				v15 = F_palloc(m, v9)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					v17 = v15
					v18 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v18)
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v18
					return v17
				}
			}
		}
	}
}
func F_HnswInitNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
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
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	v11 = F_add_size(m, v9, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = F_mul_size(m, int32(4), v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v17 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, v13, v15)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v19 = F_palloc(m, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v21 = v17
	goto L4
L9:
	;
	v21 = v19
	goto L4
L10:
	;
	v26 = v21 - l0 + int32(1)
	goto L12
L11:
	;
	v26 = int32(0)
	goto L12
L12:
	;
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v27 = v26
	goto L15
L14:
	;
	v27 = v21
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v27
	v31 = int32(0)
	goto L16
L16:
	;
	v42 = F_mul_size(m, int32(12), l2<<(uint(base.B2i32(v31 == int32(0)))%32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	return
L18:
	;
	v44 = F_add_size(m, int32(8), v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if l0 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if base.B2i32(v31 == v9) == int32(0) {
		v31 = v31 + int32(1)
		goto L16
	} else {
		goto L36
	}
L21:
	;
	if l3 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)) = uint8(v55)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v21+v31<<(uint(int32(2))%32)))) = v54
	goto L20
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v50 = m.T0[v49].(func(*base.Module, int32, int32) int32)(m, v44, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v52 = F_palloc(m, v44)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v54 = v50
	goto L24
L29:
	;
	v54 = v52
	goto L24
L30:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v21+v31<<(uint(int32(2))%32)))) = v69 - l0 + int32(1)
	goto L20
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v65 = m.T0[v64].(func(*base.Module, int32, int32) int32)(m, v44, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v67 = F_palloc(m, v44)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	v69 = v65
	goto L30
L35:
	;
	v69 = v67
	goto L30
L36:
	;
	goto L17
}
func F_HnswInitSupport(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v3 = int32(1)
	v5 = F_index_getprocinfo(m, l1, v3, v3)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v9
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+216))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v13+v15*int32(0)<<(uint(int32(2))%32)+int32(8)-int32(4))))
		if v27 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
			return
		} else {
			v34 = F_index_getprocinfo(m, l1, int32(1), int32(2))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v34
				return
			}
		}
	}
}
func F_HnswSearchLayer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v192 int64
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int64
	_ = v227
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v284 int32
	_ = v284
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 float64
	_ = v341
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
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
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v419 int64
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int64
	_ = v427
	var v428 int64
	_ = v428
	var v433 int64
	_ = v433
	var v438 int64
	_ = v438
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v583 int32
	_ = v583
	var v602 int64
	_ = v602
	var v606 int32
	_ = v606
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 float64
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 float64
	_ = v692
	var v693 int32
	_ = v693
	var v696 float64
	_ = v696
	var v697 float64
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v779 int32
	_ = v779
	var v791 int32
	_ = v791
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v841 int32
	_ = v841
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v883 int32
	_ = v883
	v15 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(1232)
	m.G0 = v38
	v42 = F_pairingheap_allocate(m, int32(7677), v15)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v48 = F_pairingheap_allocate(m, int32(7678), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v53 = l7 << (uint(base.B2i32(l4 == int32(0))) % 32)
	v54 = F_mul_size(m, int32(8), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v56 = F_palloc(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v60 = base.B2i32(l10 == int32(0)) | l12
	if v60 != int32(1) {
		v95 = v15
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if l5 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	if l5 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l10 != 0 {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v69 = F_tidhash_create(m, v64, l3*l7<<(uint(int32(1))%32), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v73 = l3 * l7 << (uint(int32(1)) % 32)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v84 = v69
	goto L8
L13:
	;
	v77 = F_offsethash_create(m, v75, v73, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v80 = F_pointerhash_create(m, v75, v73, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v84 = v77
	goto L8
L17:
	;
	v84 = v80
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v84
	goto L20
L19:
	;
	goto L20
L20:
	;
	if l11 == int32(0) {
		v95 = v84
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v90 = F_pairingheap_allocate(m, int32(7679), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v90
	v95 = v84
	goto L6
L23:
	;
	v100 = F_mul_size(m, int32(12), v53)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v106 = v15
	v107 = v15
	goto L25
L25:
	;
	v109 = l0 - int32(1)
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v102 = F_add_size(m, int32(8), v100)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v104 = F_palloc(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v106 = v104
	v107 = v102
	goto L25
L29:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v296 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L30:
	;
	v132 = v110
	v140 = v15
	goto L35
L31:
	;
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v110 < v111 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v284 = v15
	goto L29
L34:
	;
	goto L33
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v152+v132<<(uint(int32(2))%32))))
	if v60 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v284 = v256
	goto L29
L37:
	;
	F_pairingheap_add(m, v42, v156)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L68
	}
L38:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	if l5 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if l13 == int32(0) {
		goto L37
	} else {
		goto L67
	}
L40:
	;
	if v159 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if l0 != 0 {
		goto L53
	} else {
		goto L54
	}
L43:
	;
	v162 = v159 + v109
	goto L45
L44:
	;
	v162 = int32(0)
	goto L45
L45:
	;
	if l0 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v163 = v162
	goto L48
L47:
	;
	v163 = v159
	goto L48
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+76))
	v166 = int32(base.Ui32(v164) >> (uint(int32(16)) % 32))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+80)))
	if l10 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v169 = v168
	goto L51
L50:
	;
	v169 = v95
	goto L51
L51:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v38+int32(24)))) = uint16(v167)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+34)) = uint16(v164)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+32)) = uint16(v166)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+36)) = uint16(v167)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v174
	v178 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v38)+36)))
	v180 = v178 << (uint(int64(32)) % 64)
	v181 = int64(33)
	v183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v38)+32)))
	v187 = (int64(base.Ui64(v180)>>(uint(v181)%64)) ^ (v183 | v180)) * int64(-49064778989728563)
	v192 = (int64(base.Ui64(v187)>>(uint(v181)%64)) ^ v187) * int64(-4265267296055464877)
	v199 = F_tidhash_insert_hash_internal(m, v169, v38+int32(20), base.I32_wrap_i64(int64(base.Ui64(v192)>>(uint(v181)%64))^v192), v38+int32(28))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L39
L53:
	;
	if v159 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if l10 != 0 {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v203 = v159 + v109
	goto L58
L57:
	;
	v203 = int32(0)
	goto L58
L58:
	;
	if l10 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v205 = v204
	goto L61
L60:
	;
	v205 = v95
	goto L61
L61:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v211 = F_offsethash_insert_hash_internal(m, v205, v159-int32(1), v208, v38+int32(28))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L39
L63:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v214 = v213
	goto L65
L64:
	;
	v214 = v95
	goto L65
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v159)+68))
	v218 = F_pointerhash_insert_hash_internal(m, v214, v159, v215, v38+int32(28))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	goto L39
L67:
	;
	v227 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v227 + int64(1)
	goto L37
L68:
	;
	F_pairingheap_add(m, v48, v156+int32(12))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if l0 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if l9 != 0 {
		goto L77
	} else {
		goto L78
	}
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	v250 = v244
	goto L70
L72:
	;
	goto L73
L73:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	if v245 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v248 = v109 + v245
	goto L76
L75:
	;
	v248 = int32(0)
	goto L76
L76:
	;
	v250 = v248
	goto L70
L77:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+64)))
	v255 = base.B2i32(v251 != int32(0))
	goto L79
L78:
	;
	v255 = int32(1)
	goto L79
L79:
	;
	v256 = v255 + v140
	v258 = v132 + int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v258 < v259 {
		v132 = v258
		v140 = v256
		goto L35
	} else {
		goto L80
	}
L80:
	;
	goto L36
L81:
	;
	v827 = int32(0)
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v828 != 0 {
		goto L203
	} else {
		goto L204
	}
L82:
	;
	v300 = l4 << (uint(int32(2)) % 32)
	v327 = v284
	goto L83
L83:
	;
	v339 = F_pairingheap_remove_first(m, v42)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	goto L81
L85:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v339)+32))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)+20))
	if base.F64_gt(v341, v343) != 0 {
		goto L81
	} else {
		goto L86
	}
L86:
	;
	if l0 != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	if l13 != 0 {
		goto L147
	} else {
		goto L148
	}
L88:
	;
	v471 = v467 + int32(92)
	v473 = F_LWLockAcquire(m, v471, int32(1))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L116
	}
L89:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v355)+72))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v464+v300)))
	v467 = v355
	v469 = v466
	goto L88
L90:
	;
	v361 = int32(0)
	v364 = F_HnswLoadNeighborTids(m, v360, v38+int32(32), l5, l7, v53, l4)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L102
	}
L91:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v339)+24))
	v346 = v109 + v345
	if v345 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v339)+24))
	if l5 == int32(0) {
		goto L89
	} else {
		goto L101
	}
L94:
	;
	v348 = v346
	goto L96
L95:
	;
	v348 = int32(0)
	goto L96
L96:
	;
	if l5 != 0 {
		v360 = v348
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+72))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v109+v300+v349)))
	if v351 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v354 = v109 + v351
	goto L100
L99:
	;
	v354 = int32(0)
	goto L100
L100:
	;
	v467 = v346
	v469 = v354
	goto L88
L101:
	;
	v360 = v355
	goto L90
L102:
	;
	if l7 <= int32(0) {
		v583 = v361
		goto L87
	} else {
		goto L103
	}
L103:
	;
	v368 = int32(0)
	if v364 == v368 {
		v583 = v361
		goto L87
	} else {
		goto L104
	}
L104:
	;
	v386 = v368
	v387 = v361
	goto L105
L105:
	;
	v410 = v38 + int32(32) + v386*int32(6)
	v411 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v410)+4)))
	if v411 == int32(0) {
		v583 = v387
		goto L87
	} else {
		goto L107
	}
L106:
	;
	v583 = v460
	goto L87
L107:
	;
	if l10 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v415 = v414
	goto L110
L109:
	;
	v415 = v95
	goto L110
L110:
	;
	v417 = v410 + int32(4)
	v418 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v417))))
	v419 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v410))))
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417))))
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)) = uint16(v420)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v422
	v427 = v418 << (uint(int64(32)) % 64)
	v428 = int64(33)
	v433 = (int64(base.Ui64(v427)>>(uint(v428)%64)) ^ (v427 | v419)) * int64(-49064778989728563)
	v438 = (int64(base.Ui64(v433)>>(uint(v428)%64)) ^ v433) * int64(-4265267296055464877)
	v445 = F_tidhash_insert_hash_internal(m, v415, v38+int32(12), base.I32_wrap_i64(int64(base.Ui64(v438)>>(uint(v428)%64))^v438), v38+int32(28))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+28)))
	if v447 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v452 = v56 + v387<<(uint(int32(3))%32)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(v452))) = v453
	v455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417))))
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+4)) = uint16(v455)
	v460 = v387 + int32(1)
	goto L114
L113:
	;
	v460 = v387
	goto L114
L114:
	;
	v462 = v386 + int32(1)
	if v462 != v53 {
		v386 = v462
		v387 = v460
		goto L105
	} else {
		goto L115
	}
L115:
	;
	goto L106
L116:
	;
	if v107 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	F_LWLockRelease(m, v471)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L121
	}
L118:
	;
	v475 = F__emscripten_memcpy_bulkmem(m, v106, v469, v107)
	mBase = m.M
	v476 = v475
	goto L120
L119:
	;
	v476 = v106
	goto L120
L120:
	;
	goto L117
L121:
	;
	v479 = int32(0)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	if v481 <= v479 {
		v583 = v479
		goto L87
	} else {
		goto L122
	}
L122:
	;
	v499 = v479
	v500 = v479
	goto L123
L123:
	;
	v521 = v106 + int32(8) + v499*int32(12)
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	if l0 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L124:
	;
	v583 = v561
	goto L87
L125:
	;
	v564 = v499 + int32(1)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v476)))
	if v564 < v565 {
		v499 = v564
		v500 = v561
		goto L123
	} else {
		goto L146
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v500<<(uint(int32(3))%32)))) = v556
	v561 = v500 + int32(1)
	goto L125
L127:
	;
	if l10 != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	if v522 != 0 {
		goto L135
	} else {
		goto L136
	}
L130:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v529 = v528
	goto L132
L131:
	;
	v529 = v95
	goto L132
L132:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v522)+68))
	v533 = F_pointerhash_insert_hash_internal(m, v529, v522, v530, v38+int32(32))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+32)))
	if v535 != 0 {
		v561 = v500
		goto L125
	} else {
		goto L134
	}
L134:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	v556 = v536
	goto L126
L135:
	;
	v539 = v522 + v109
	goto L137
L136:
	;
	v539 = int32(0)
	goto L137
L137:
	;
	if l10 != 0 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v541 = v540
	goto L140
L139:
	;
	v541 = v95
	goto L140
L140:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v539)+68))
	v547 = F_offsethash_insert_hash_internal(m, v541, v522-int32(1), v544, v38+int32(32))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+32)))
	if v549 != 0 {
		v561 = v500
		goto L125
	} else {
		goto L142
	}
L142:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v521)))
	if v550 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v553 = v109 + v550
	goto L145
L144:
	;
	v553 = int32(0)
	goto L145
L145:
	;
	v556 = v553
	goto L126
L146:
	;
	goto L124
L147:
	;
	v602 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v602 + base.I64_extend_i32_s(v583)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v606 = int32(0)
	if v606 < v583 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v624 = v606
	v632 = v327
	goto L153
L151:
	;
	v779 = v327
	goto L152
L152:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	if v791 != 0 {
		v327 = v779
		goto L83
	} else {
		goto L202
	}
L153:
	;
	v646 = v56 + v624<<(uint(int32(3))%32)
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if l5 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L154:
	;
	v779 = v751
	goto L152
L155:
	;
	v754 = v624 + int32(1)
	if v754 != v583 {
		v624 = v754
		v632 = v751
		goto L153
	} else {
		goto L201
	}
L156:
	;
	v697 = *(*float64)(unsafe.Add(mBase, uint32(v647)+20))
	if base.F64_lt(v696, v697) != 0 {
		goto L176
	} else {
		goto L177
	}
L157:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646)+4)))
	v669 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646)+2)))
	v670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v646))))
	v671 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v671
	if l3 <= v632 {
		goto L168
	} else {
		goto L169
	}
L160:
	;
	v664 = F_FunctionCall2Coll(m, v652, v653, v654, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L167
	}
L161:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v650)+88))
	v663 = v657
	goto L160
L162:
	;
	goto L163
L163:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v650)+88))
	if v658 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v661 = v109 + v658
	goto L166
L165:
	;
	v661 = int32(0)
	goto L166
L166:
	;
	v663 = v661
	goto L160
L167:
	;
	v666 = *(*float64)(unsafe.Add(mBase, uint32(v664)))
	*(*float64)(unsafe.Add(mBase, uint32(v38)+32)) = v666
	v693 = v650
	v696 = v666
	goto L156
L168:
	;
	v683 = v647 + int32(20)
	goto L170
L169:
	;
	v683 = v671
	goto L170
L170:
	;
	if l11 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v684 = v671
	goto L173
L172:
	;
	v684 = v683
	goto L173
L173:
	;
	F_HnswLoadElementImpl(m, v669|v670<<(uint(int32(16))%32), v668, v38+int32(32), l1, l5, l6, l8, v684, v38+int32(28))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	if v689 == int32(0) {
		v751 = v632
		goto L155
	} else {
		goto L175
	}
L175:
	;
	v692 = *(*float64)(unsafe.Add(mBase, uint32(v38)+32))
	v693 = v689
	v696 = v692
	goto L156
L176:
	;
	v716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693)+65)))
	if v716 < l4 {
		v751 = v632
		goto L155
	} else {
		goto L185
	}
L177:
	;
	if v632 < l3 {
		goto L176
	} else {
		goto L178
	}
L178:
	;
	if l11 == int32(0) {
		v751 = v632
		goto L155
	} else {
		goto L179
	}
L179:
	;
	v703 = F_palloc(m, int32(40))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v703)+32)) = v696
	if l0 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v709 = v693 - l0 + int32(1)
	goto L183
L182:
	;
	v709 = v693
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v703)+24)) = v709
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v711, v703+int32(12))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v751 = v632
	goto L155
L185:
	;
	v719 = F_palloc(m, int32(40))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v719)+32)) = v696
	if l0 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v725 = v693 - l0 + int32(1)
	goto L189
L188:
	;
	v725 = v693
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v719)+24)) = v725
	F_pairingheap_add(m, v42, v719)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_pairingheap_add(m, v48, v719+int32(12))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if l9 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693)+64)))
	if v733 == int32(0) {
		v751 = v632
		goto L155
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	if v632 < l3 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L194
L196:
	;
	v751 = v632 + int32(1)
	goto L155
L197:
	;
	v739 = F_pairingheap_remove_first(m, v48)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	if l11 == int32(0) {
		goto L196
	} else {
		goto L199
	}
L199:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v743, v739)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	goto L196
L201:
	;
	goto L154
L202:
	;
	goto L84
L203:
	;
	v841 = v827
	goto L206
L204:
	;
	v883 = v827
	goto L205
L205:
	;
	m.G0 = v38 + int32(1232)
	return v883
L206:
	;
	v864 = F_pairingheap_remove_first(m, v48)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L208
	}
L207:
	;
	v883 = v868
	goto L205
L208:
	;
	v868 = F_lappend(m, v841, v864-int32(12))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v870 != 0 {
		v841 = v868
		goto L206
	} else {
		goto L210
	}
L210:
	;
	goto L207
}
func F_hnsw_sparsevec_support(m *base.Module, l0 int32) int32 {
	return int32(4120640)
}
