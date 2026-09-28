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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v104 int32
	_ = v104
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v172 int64
	_ = v172
	var v176 int64
	_ = v176
	var v181 int64
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
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
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int64
	_ = v214
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
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
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v516 int64
	_ = v516
	var v517 int64
	_ = v517
	var v522 int64
	_ = v522
	var v527 int64
	_ = v527
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v568 int32
	_ = v568
	var v585 int64
	_ = v585
	var v589 int32
	_ = v589
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int64
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v648 int64
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 float64
	_ = v676
	var v678 int32
	_ = v678
	var v680 float64
	_ = v680
	var v682 float64
	_ = v682
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v767 int32
	_ = v767
	var v778 int32
	_ = v778
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v825 int32
	_ = v825
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	v15 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(1232)
	m.G0 = v35
	v39 = F_pairingheap_allocate(m, int32(_a_F_HnswSearchLayer_0), v15)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = F_pairingheap_allocate(m, int32(_a_F_HnswSearchLayer_1), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v50 = l7 << (uint(base.B2i32(l4 == int32(0))) % 32)
	v51 = F_palloc_mul(m, int32(8), v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v55 = base.B2i32(l10 == int32(0)) | l12
	if v55 != int32(1) {
		v89 = v15
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l5 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L6:
	;
	if l5 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if l10 != 0 {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_HnswSearchLayer[0]))
	v64 = F_tidhash_create(m, v59, l3*l7<<(uint(int32(1))%32), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v68 = l3 * l7 << (uint(int32(1)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_HnswSearchLayer[0]))
	if l0 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v79 = v64
	goto L7
L12:
	;
	v72 = F_offsethash_create(m, v70, v68, int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v75 = F_pointerhash_create(m, v70, v68, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v79 = v72
	goto L7
L16:
	;
	v79 = v75
	goto L7
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v79
	goto L19
L18:
	;
	goto L19
L19:
	;
	if l11 == int32(0) {
		v89 = v79
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v85 = F_pairingheap_allocate(m, int32(_a_F_HnswSearchLayer_2), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v85
	v89 = v79
	goto L5
L22:
	;
	v94 = F_mul_size(m, int32(12), v50)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v100 = v15
	v101 = v15
	goto L24
L24:
	;
	if l2 == int32(0) {
		v276 = v15
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v96 = F_add_size(m, int32(8), v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v98 = F_palloc(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	v101 = v96
	goto L24
L28:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v287 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L29:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v104 <= int32(0) {
		v276 = v15
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v123 = v15
	v128 = v15
	goto L31
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v123<<(uint(int32(2))%32))))
	if v55 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v276 = v250
	goto L28
L33:
	;
	F_pairingheap_add(m, v39, v143)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L61
	}
L34:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	if l5 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if l13 == int32(0) {
		goto L33
	} else {
		goto L60
	}
L36:
	;
	if v146 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if l0 != 0 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v151 = l0 + v146 - int32(1)
	goto L41
L40:
	;
	v151 = int32(0)
	goto L41
L41:
	;
	if l0 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v152 = v151
	goto L44
L43:
	;
	v152 = v146
	goto L44
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+76))
	v155 = int32(base.Ui32(v153) >> (uint(int32(16)) % 32))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+80)))
	if l10 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v158 = v157
	goto L47
L46:
	;
	v158 = v89
	goto L47
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+34)) = uint16(v153)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+32)) = uint16(v155)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)) = uint16(v156)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v156)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v163
	v167 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v169 = v167 << (uint(int64(32)) % 64)
	v170 = int64(33)
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+32)))
	v176 = (int64(base.Ui64(v169)>>(uint(v170)%64)) ^ (v172 | v169)) * int64(-49064778989728563)
	v181 = (int64(base.Ui64(v176)>>(uint(v170)%64)) ^ v176) * int64(-4265267296055464877)
	v188 = F_tidhash_insert_hash_internal(m, v158, v35+int32(20), base.I32_wrap_i64(int64(base.Ui64(v181)>>(uint(v170)%64))^v181), v35+int32(28))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L35
L49:
	;
	if l10 != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if l10 != 0 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v192 = v191
	goto L54
L53:
	;
	v192 = v89
	goto L54
L54:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0+v146)+67))
	v198 = F_offsethash_insert_hash_internal(m, v192, v146-int32(1), v195, v35+int32(28))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L35
L56:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v201 = v200
	goto L58
L57:
	;
	v201 = v89
	goto L58
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v146)+68))
	v205 = F_pointerhash_insert_hash_internal(m, v201, v146, v202, v35+int32(28))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L35
L60:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v214 + int64(1)
	goto L33
L61:
	;
	F_pairingheap_add(m, v45, v143+int32(12))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if l0 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if l9 != 0 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	v240 = v231
	goto L63
L65:
	;
	goto L66
L66:
	;
	v232 = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	if v233 == v232 {
		v240 = v232
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v240 = l0 + v233 - int32(1)
	goto L63
L68:
	;
	v241 = int32(0)
	v244 = base.AtomicRmwOr32(m, v241, int32(_a_F_HnswSearchLayer_3), v241)
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+64)))
	v249 = base.B2i32(v245 != v241)
	goto L70
L69:
	;
	v249 = int32(1)
	goto L70
L70:
	;
	v250 = v249 + v128
	v252 = v123 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v252 < v253 {
		v123 = v252
		v128 = v250
		goto L31
	} else {
		goto L71
	}
L71:
	;
	goto L32
L72:
	;
	v811 = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v812 != 0 {
		goto L184
	} else {
		goto L185
	}
L73:
	;
	v293 = l4 << (uint(int32(2)) % 32)
	v315 = v276
	goto L74
L74:
	;
	v326 = F_pairingheap_remove_first(m, v39)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	goto L72
L76:
	;
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v326)+32))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v329)+20))
	if base.F64_gt(v328, v330) != 0 {
		goto L72
	} else {
		goto L77
	}
L77:
	;
	if l0 != 0 {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	if l13 != 0 {
		goto L130
	} else {
		goto L131
	}
L79:
	;
	v453 = int32(0)
	v456 = F_HnswLoadNeighborTids(m, v451, v35+int32(32), l5, l7, v50, l4)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L116
	}
L80:
	;
	v357 = v353 + int32(92)
	v359 = F_LWLockAcquire(m, v357, int32(1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L91
	}
L81:
	;
	v353 = v337
	v355 = l0 + v343 - int32(1)
	goto L80
L82:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326)+24))
	v333 = l0 + v332
	if v332 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v326)+24))
	if l5 != 0 {
		v451 = v345
		goto L79
	} else {
		goto L90
	}
L85:
	;
	v337 = v333 - int32(1)
	goto L87
L86:
	;
	v337 = int32(0)
	goto L87
L87:
	;
	if l5 != 0 {
		v451 = v337
		goto L79
	} else {
		goto L88
	}
L88:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v333)+71))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0+v338+v293-int32(1))))
	if v343 != 0 {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v353 = v337
	v355 = int32(0)
	goto L80
L90:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+72))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346+v293)))
	v353 = v345
	v355 = v348
	goto L80
L91:
	;
	if v101 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	base.MemoryCopy(m, v100, v355, v101)
	goto L94
L93:
	;
	goto L94
L94:
	;
	F_LWLockRelease(m, v357)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v364 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v366 <= v364 {
		v568 = v364
		goto L78
	} else {
		goto L96
	}
L96:
	;
	v381 = v364
	v384 = v364
	goto L97
L97:
	;
	v403 = v100 + int32(8) + v381*int32(12)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if l0 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v568 = v444
	goto L78
L99:
	;
	v447 = v381 + int32(1)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v447 < v448 {
		v381 = v447
		v384 = v444
		goto L97
	} else {
		goto L115
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v384<<(uint(int32(3))%32)))) = v439
	v444 = v384 + int32(1)
	goto L99
L101:
	;
	if l10 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	if l10 != 0 {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v411 = v410
	goto L106
L105:
	;
	v411 = v89
	goto L106
L106:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v404)+68))
	v415 = F_pointerhash_insert_hash_internal(m, v411, v404, v412, v35+int32(32))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	if v417 != 0 {
		v444 = v384
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	v439 = v418
	goto L100
L109:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v421 = v420
	goto L111
L110:
	;
	v421 = v89
	goto L111
L111:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0+v404)+67))
	v427 = F_offsethash_insert_hash_internal(m, v421, v404-int32(1), v424, v35+int32(32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	if v429 != 0 {
		v444 = v384
		goto L99
	} else {
		goto L113
	}
L113:
	;
	v430 = int32(0)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v431 == v430 {
		v439 = v430
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v439 = l0 + v431 - int32(1)
	goto L100
L115:
	;
	goto L98
L116:
	;
	if l7 <= int32(0) {
		v568 = v453
		goto L78
	} else {
		goto L117
	}
L117:
	;
	v460 = int32(0)
	if v456 == v460 {
		v568 = v453
		goto L78
	} else {
		goto L118
	}
L118:
	;
	v478 = v453
	v479 = v460
	goto L119
L119:
	;
	v499 = v35 + int32(32) + v479*int32(6)
	v500 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v499)+4)))
	if v500 == int32(0) {
		v568 = v478
		goto L78
	} else {
		goto L121
	}
L120:
	;
	v568 = v549
	goto L78
L121:
	;
	if l10 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v504 = v503
	goto L124
L123:
	;
	v504 = v89
	goto L124
L124:
	;
	v506 = v499 + int32(4)
	v507 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v506))))
	v508 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v499))))
	v509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506))))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v509)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v511
	v516 = v507 << (uint(int64(32)) % 64)
	v517 = int64(33)
	v522 = (int64(base.Ui64(v516)>>(uint(v517)%64)) ^ (v516 | v508)) * int64(-49064778989728563)
	v527 = (int64(base.Ui64(v522)>>(uint(v517)%64)) ^ v522) * int64(-4265267296055464877)
	v534 = F_tidhash_insert_hash_internal(m, v504, v35+int32(12), base.I32_wrap_i64(int64(base.Ui64(v527)>>(uint(v517)%64))^v527), v35+int32(28))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+28)))
	if v536 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v541 = v51 + v478<<(uint(int32(3))%32)
	v542 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v499)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v541)+4)) = uint16(v542)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v499)))
	*(*int32)(unsafe.Add(mBase, uint32(v541))) = v544
	v549 = v478 + int32(1)
	goto L128
L127:
	;
	v549 = v478
	goto L128
L128:
	;
	v551 = v479 + int32(1)
	if v551 != v50 {
		v478 = v549
		v479 = v551
		goto L119
	} else {
		goto L129
	}
L129:
	;
	goto L120
L130:
	;
	v585 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v585 + base.I64_extend_i32_s(v568)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v589 = int32(0)
	if v589 < v568 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v608 = v589
	v613 = v315
	goto L136
L134:
	;
	v767 = v315
	goto L135
L135:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v778 != 0 {
		v315 = v767
		goto L74
	} else {
		goto L183
	}
L136:
	;
	v626 = v51 + v608<<(uint(int32(3))%32)
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if l5 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v767 = v740
	goto L135
L138:
	;
	v744 = v608 + int32(1)
	if v744 != v568 {
		v608 = v744
		v613 = v740
		goto L136
	} else {
		goto L182
	}
L139:
	;
	v682 = *(*float64)(unsafe.Add(mBase, uint32(v627)+20))
	if base.B2i32(v613 < l3)|base.F64_lt(v680, v682) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L140:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v630
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v634 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L142
L142:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v626)+4)))
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v626)+2)))
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v626))))
	v655 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v655
	if l3 <= v613 {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	v648 = F_FunctionCall2Coll(m, v632, v633, v634, base.I64_extend_i32_u(v646))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L148
	}
L144:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v630)+88))
	v646 = v637
	goto L143
L145:
	;
	goto L146
L146:
	;
	v638 = int32(0)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v630)+88))
	if v639 == v638 {
		v646 = v638
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v646 = l0 + v639 - int32(1)
	goto L143
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v648
	v678 = v630
	v680 = base.F64_reinterpret_i64(v648)
	goto L139
L149:
	;
	v667 = v627 + int32(20)
	goto L151
L150:
	;
	v667 = v655
	goto L151
L151:
	;
	if l11 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v668 = v655
	goto L154
L153:
	;
	v668 = v667
	goto L154
L154:
	;
	F_HnswLoadElementImpl(m, v653|v654<<(uint(int32(16))%32), v652, v35+int32(32), l1, l5, l6, l8, v668, v35+int32(28))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	if v673 == int32(0) {
		v740 = v613
		goto L138
	} else {
		goto L156
	}
L156:
	;
	v676 = *(*float64)(unsafe.Add(mBase, uint32(v35)+32))
	v678 = v673
	v680 = v676
	goto L139
L157:
	;
	if l11 == int32(0) {
		v740 = v613
		goto L138
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+65)))
	if v703 < l4 {
		v740 = v613
		goto L138
	} else {
		goto L166
	}
L160:
	;
	v690 = F_palloc(m, int32(40))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v690)+32)) = v680
	if l0 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v696 = v678 - l0 + int32(1)
	goto L164
L163:
	;
	v696 = v678
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v690)+24)) = v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v698, v690+int32(12))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v740 = v613
	goto L138
L166:
	;
	v706 = F_palloc(m, int32(40))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v706)+32)) = v680
	if l0 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v712 = v678 - l0 + int32(1)
	goto L170
L169:
	;
	v712 = v678
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+24)) = v712
	F_pairingheap_add(m, v39, v706)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_pairingheap_add(m, v45, v706+int32(12))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	if l9 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v720 = int32(0)
	v723 = base.AtomicRmwOr32(m, v720, int32(_a_F_HnswSearchLayer_3), v720)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678)+64)))
	if v724 == v720 {
		v740 = v613
		goto L138
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if v613 < l3 {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	goto L175
L177:
	;
	v740 = v613 + int32(1)
	goto L138
L178:
	;
	v730 = F_pairingheap_remove_first(m, v45)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	if l11 == int32(0) {
		goto L177
	} else {
		goto L180
	}
L180:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v734, v730)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	goto L177
L182:
	;
	goto L137
L183:
	;
	goto L75
L184:
	;
	v825 = v811
	goto L187
L185:
	;
	v864 = v811
	goto L186
L186:
	;
	m.G0 = v35 + int32(1232)
	return v864
L187:
	;
	v845 = F_pairingheap_remove_first(m, v45)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	v864 = v849
	goto L186
L189:
	;
	v849 = F_lappend(m, v825, v845-int32(12))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v851 != 0 {
		v825 = v849
		goto L187
	} else {
		goto L191
	}
L191:
	;
	goto L188
}
func F_hnsw_sparsevec_support(m *base.Module, l0 int32) int64 {
	return int64(4171864)
}
