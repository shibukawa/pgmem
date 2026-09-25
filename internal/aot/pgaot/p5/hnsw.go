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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v178 int64
	_ = v178
	var v183 int64
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int64
	_ = v216
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 float64
	_ = v330
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int64
	_ = v518
	var v519 int64
	_ = v519
	var v524 int64
	_ = v524
	var v529 int64
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v570 int32
	_ = v570
	var v587 int64
	_ = v587
	var v591 int32
	_ = v591
	var v610 int32
	_ = v610
	var v615 int32
	_ = v615
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 float64
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 float64
	_ = v677
	var v679 int32
	_ = v679
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
	var v743 int32
	_ = v743
	var v766 int32
	_ = v766
	var v777 int32
	_ = v777
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v824 int32
	_ = v824
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v863 int32
	_ = v863
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
	v51 = F_mul_size(m, int32(8), v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v53 = F_palloc(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v57 = base.B2i32(l10 == int32(0)) | l12
	if v57 != int32(1) {
		v91 = v15
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
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_HnswSearchLayer[0]))
	v66 = F_tidhash_create(m, v61, l3*l7<<(uint(int32(1))%32), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v70 = l3 * l7 << (uint(int32(1)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_HnswSearchLayer[0]))
	if l0 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v81 = v66
	goto L8
L13:
	;
	v74 = F_offsethash_create(m, v72, v70, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v77 = F_pointerhash_create(m, v72, v70, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v81 = v74
	goto L8
L17:
	;
	v81 = v77
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v81
	goto L20
L19:
	;
	goto L20
L20:
	;
	if l11 == int32(0) {
		v91 = v81
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v87 = F_pairingheap_allocate(m, int32(_a_F_HnswSearchLayer_2), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = v87
	v91 = v81
	goto L6
L23:
	;
	v96 = F_mul_size(m, int32(12), v50)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v102 = v15
	v103 = v15
	goto L25
L25:
	;
	if l2 == int32(0) {
		v278 = v15
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v98 = F_add_size(m, int32(8), v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v100 = F_palloc(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v102 = v100
	v103 = v98
	goto L25
L29:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v289 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v106 <= int32(0) {
		v278 = v15
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v125 = v15
	v130 = v15
	goto L32
L32:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v125<<(uint(int32(2))%32))))
	if v57 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v278 = v252
	goto L29
L34:
	;
	F_pairingheap_add(m, v39, v145)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L62
	}
L35:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if l5 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if l13 == int32(0) {
		goto L34
	} else {
		goto L61
	}
L37:
	;
	if v148 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if l0 != 0 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v153 = l0 + v148 - int32(1)
	goto L42
L41:
	;
	v153 = int32(0)
	goto L42
L42:
	;
	if l0 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v154 = v153
	goto L45
L44:
	;
	v154 = v148
	goto L45
L45:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+76))
	v157 = int32(base.Ui32(v155) >> (uint(int32(16)) % 32))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+80)))
	if l10 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v160 = v159
	goto L48
L47:
	;
	v160 = v91
	goto L48
L48:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+34)) = uint16(v155)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+32)) = uint16(v157)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)) = uint16(v158)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+24)) = uint16(v158)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v165
	v169 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v171 = v169 << (uint(int64(32)) % 64)
	v172 = int64(33)
	v174 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35)+32)))
	v178 = (int64(base.Ui64(v171)>>(uint(v172)%64)) ^ (v174 | v171)) * int64(-49064778989728563)
	v183 = (int64(base.Ui64(v178)>>(uint(v172)%64)) ^ v178) * int64(-4265267296055464877)
	v190 = F_tidhash_insert_hash_internal(m, v160, v35+int32(20), base.I32_wrap_i64(int64(base.Ui64(v183)>>(uint(v172)%64))^v183), v35+int32(28))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L36
L50:
	;
	if l10 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	if l10 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v194 = v193
	goto L55
L54:
	;
	v194 = v91
	goto L55
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0+v148)+67))
	v200 = F_offsethash_insert_hash_internal(m, v194, v148-int32(1), v197, v35+int32(28))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L36
L57:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v203 = v202
	goto L59
L58:
	;
	v203 = v91
	goto L59
L59:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v148)+68))
	v207 = F_pointerhash_insert_hash_internal(m, v203, v148, v204, v35+int32(28))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L36
L61:
	;
	v216 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v216 + int64(1)
	goto L34
L62:
	;
	F_pairingheap_add(m, v45, v145+int32(12))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if l0 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if l9 != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	v242 = v233
	goto L64
L66:
	;
	goto L67
L67:
	;
	v234 = int32(0)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	if v235 == v234 {
		v242 = v234
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v242 = l0 + v235 - int32(1)
	goto L64
L69:
	;
	v243 = int32(0)
	v246 = base.AtomicRmwOr32(m, v243, int32(_a_F_HnswSearchLayer_3), v243)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+64)))
	v251 = base.B2i32(v247 != v243)
	goto L71
L70:
	;
	v251 = int32(1)
	goto L71
L71:
	;
	v252 = v251 + v130
	v254 = v125 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v254 < v255 {
		v125 = v254
		v130 = v252
		goto L32
	} else {
		goto L72
	}
L72:
	;
	goto L33
L73:
	;
	v810 = int32(0)
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v811 != 0 {
		goto L185
	} else {
		goto L186
	}
L74:
	;
	v295 = l4 << (uint(int32(2)) % 32)
	v317 = v278
	goto L75
L75:
	;
	v328 = F_pairingheap_remove_first(m, v39)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	goto L73
L77:
	;
	v330 = *(*float64)(unsafe.Add(mBase, uint32(v328)+32))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v331)+20))
	if base.F64_gt(v330, v332) != 0 {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	if l0 != 0 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	if l13 != 0 {
		goto L131
	} else {
		goto L132
	}
L80:
	;
	v455 = int32(0)
	v458 = F_HnswLoadNeighborTids(m, v453, v35+int32(32), l5, l7, v50, l4)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L117
	}
L81:
	;
	v359 = v355 + int32(92)
	v361 = F_LWLockAcquire(m, v359, int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L92
	}
L82:
	;
	v355 = v339
	v357 = l0 + v345 - int32(1)
	goto L81
L83:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v335 = l0 + v334
	if v334 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	if l5 != 0 {
		v453 = v347
		goto L80
	} else {
		goto L91
	}
L86:
	;
	v339 = v335 - int32(1)
	goto L88
L87:
	;
	v339 = int32(0)
	goto L88
L88:
	;
	if l5 != 0 {
		v453 = v339
		goto L80
	} else {
		goto L89
	}
L89:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v335)+71))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0+v340+v295-int32(1))))
	if v345 != 0 {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v355 = v339
	v357 = int32(0)
	goto L81
L91:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+72))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v348+v295)))
	v355 = v347
	v357 = v350
	goto L81
L92:
	;
	if v103 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	base.MemoryCopy(m, v102, v357, v103)
	goto L95
L94:
	;
	goto L95
L95:
	;
	F_LWLockRelease(m, v359)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v366 = int32(0)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v368 <= v366 {
		v570 = v366
		goto L79
	} else {
		goto L97
	}
L97:
	;
	v383 = v366
	v386 = v366
	goto L98
L98:
	;
	v405 = v102 + int32(8) + v383*int32(12)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if l0 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	v570 = v446
	goto L79
L100:
	;
	v449 = v383 + int32(1)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v449 < v450 {
		v383 = v449
		v386 = v446
		goto L98
	} else {
		goto L116
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53+v386<<(uint(int32(3))%32)))) = v441
	v446 = v386 + int32(1)
	goto L100
L102:
	;
	if l10 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if l10 != 0 {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v413 = v412
	goto L107
L106:
	;
	v413 = v91
	goto L107
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v406)+68))
	v417 = F_pointerhash_insert_hash_internal(m, v413, v406, v414, v35+int32(32))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	if v419 != 0 {
		v446 = v386
		goto L100
	} else {
		goto L109
	}
L109:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	v441 = v420
	goto L101
L110:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v423 = v422
	goto L112
L111:
	;
	v423 = v91
	goto L112
L112:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(l0+v406)+67))
	v429 = F_offsethash_insert_hash_internal(m, v423, v406-int32(1), v426, v35+int32(32))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+32)))
	if v431 != 0 {
		v446 = v386
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v432 = int32(0)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	if v433 == v432 {
		v441 = v432
		goto L101
	} else {
		goto L115
	}
L115:
	;
	v441 = l0 + v433 - int32(1)
	goto L101
L116:
	;
	goto L99
L117:
	;
	if l7 <= int32(0) {
		v570 = v455
		goto L79
	} else {
		goto L118
	}
L118:
	;
	v462 = int32(0)
	if v458 == v462 {
		v570 = v455
		goto L79
	} else {
		goto L119
	}
L119:
	;
	v480 = v455
	v481 = v462
	goto L120
L120:
	;
	v501 = v35 + int32(32) + v481*int32(6)
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v501)+4)))
	if v502 == int32(0) {
		v570 = v480
		goto L79
	} else {
		goto L122
	}
L121:
	;
	v570 = v551
	goto L79
L122:
	;
	if l10 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v506 = v505
	goto L125
L124:
	;
	v506 = v91
	goto L125
L125:
	;
	v508 = v501 + int32(4)
	v509 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v508))))
	v510 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v501))))
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508))))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v511)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v513
	v518 = v509 << (uint(int64(32)) % 64)
	v519 = int64(33)
	v524 = (int64(base.Ui64(v518)>>(uint(v519)%64)) ^ (v518 | v510)) * int64(-49064778989728563)
	v529 = (int64(base.Ui64(v524)>>(uint(v519)%64)) ^ v524) * int64(-4265267296055464877)
	v536 = F_tidhash_insert_hash_internal(m, v506, v35+int32(12), base.I32_wrap_i64(int64(base.Ui64(v529)>>(uint(v519)%64))^v529), v35+int32(28))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+28)))
	if v538 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v543 = v53 + v480<<(uint(int32(3))%32)
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v501)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v543)+4)) = uint16(v544)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v501)))
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v546
	v551 = v480 + int32(1)
	goto L129
L128:
	;
	v551 = v480
	goto L129
L129:
	;
	v553 = v481 + int32(1)
	if v553 != v50 {
		v480 = v551
		v481 = v553
		goto L120
	} else {
		goto L130
	}
L130:
	;
	goto L121
L131:
	;
	v587 = *(*int64)(unsafe.Add(mBase, uint32(l13)))
	*(*int64)(unsafe.Add(mBase, uint32(l13))) = v587 + base.I64_extend_i32_s(v570)
	goto L133
L132:
	;
	goto L133
L133:
	;
	v591 = int32(0)
	if v591 < v570 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v610 = v591
	v615 = v317
	goto L137
L135:
	;
	v766 = v317
	goto L136
L136:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v777 != 0 {
		v317 = v766
		goto L75
	} else {
		goto L184
	}
L137:
	;
	v628 = v53 + v610<<(uint(int32(3))%32)
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if l5 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v766 = v740
	goto L136
L139:
	;
	v743 = v610 + int32(1)
	if v743 != v570 {
		v610 = v743
		v615 = v740
		goto L137
	} else {
		goto L183
	}
L140:
	;
	v682 = *(*float64)(unsafe.Add(mBase, uint32(v629)+20))
	if base.B2i32(v615 < l3)|base.F64_lt(v680, v682) == int32(0) {
		goto L158
	} else {
		goto L159
	}
L141:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	goto L143
L143:
	;
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+4)))
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+2)))
	v655 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628))))
	v656 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+28)) = v656
	if l3 <= v615 {
		goto L150
	} else {
		goto L151
	}
L144:
	;
	v649 = F_FunctionCall2Coll(m, v634, v635, v636, v648)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L1
	} else {
		goto L149
	}
L145:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v632)+88))
	v648 = v639
	goto L144
L146:
	;
	goto L147
L147:
	;
	v640 = int32(0)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v632)+88))
	if v641 == v640 {
		v648 = v640
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v648 = l0 + v641 - int32(1)
	goto L144
L149:
	;
	v651 = *(*float64)(unsafe.Add(mBase, uint32(v649)))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+32)) = v651
	v679 = v632
	v680 = v651
	goto L140
L150:
	;
	v668 = v629 + int32(20)
	goto L152
L151:
	;
	v668 = v656
	goto L152
L152:
	;
	if l11 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v669 = v656
	goto L155
L154:
	;
	v669 = v668
	goto L155
L155:
	;
	F_HnswLoadElementImpl(m, v654|v655<<(uint(int32(16))%32), v653, v35+int32(32), l1, l5, l6, l8, v669, v35+int32(28))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
	if v674 == int32(0) {
		v740 = v615
		goto L139
	} else {
		goto L157
	}
L157:
	;
	v677 = *(*float64)(unsafe.Add(mBase, uint32(v35)+32))
	v679 = v674
	v680 = v677
	goto L140
L158:
	;
	if l11 == int32(0) {
		v740 = v615
		goto L139
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+65)))
	if v703 < l4 {
		v740 = v615
		goto L139
	} else {
		goto L167
	}
L161:
	;
	v690 = F_palloc(m, int32(40))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v690)+32)) = v680
	if l0 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v696 = v679 - l0 + int32(1)
	goto L165
L164:
	;
	v696 = v679
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v690)+24)) = v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v698, v690+int32(12))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v740 = v615
	goto L139
L167:
	;
	v706 = F_palloc(m, int32(40))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v706)+32)) = v680
	if l0 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v712 = v679 - l0 + int32(1)
	goto L171
L170:
	;
	v712 = v679
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v706)+24)) = v712
	F_pairingheap_add(m, v39, v706)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_pairingheap_add(m, v45, v706+int32(12))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if l9 != 0 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v720 = int32(0)
	v723 = base.AtomicRmwOr32(m, v720, int32(_a_F_HnswSearchLayer_3), v720)
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v679)+64)))
	if v724 == v720 {
		v740 = v615
		goto L139
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	if v615 < l3 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	v740 = v615 + int32(1)
	goto L139
L179:
	;
	v730 = F_pairingheap_remove_first(m, v45)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	if l11 == int32(0) {
		goto L178
	} else {
		goto L181
	}
L181:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	F_pairingheap_add(m, v734, v730)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	goto L178
L183:
	;
	goto L138
L184:
	;
	goto L76
L185:
	;
	v824 = v810
	goto L188
L186:
	;
	v863 = v810
	goto L187
L187:
	;
	m.G0 = v35 + int32(1232)
	return v863
L188:
	;
	v844 = F_pairingheap_remove_first(m, v45)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L190
	}
L189:
	;
	v863 = v848
	goto L187
L190:
	;
	v848 = F_lappend(m, v824, v844-int32(12))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v850 != 0 {
		v824 = v848
		goto L188
	} else {
		goto L192
	}
L192:
	;
	goto L189
}
func F_hnsw_sparsevec_support(m *base.Module, l0 int32) int32 {
	return int32(_a_F_hnsw_sparsevec_support_0)
}
