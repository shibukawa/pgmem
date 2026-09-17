package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateCachedPlan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int64
	_ = v38
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCachedPlan[0]))
	v13 = F_AllocSetContextCreateInternal(m, v8, int32(_a_F_CreateCachedPlan_0), int32(0), int32(1024), int32(_a_F_CreateCachedPlan_1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(_a_F_CreateCachedPlan_2)
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_CreateCachedPlan[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_CreateCachedPlan[0])) = v13
		v22 = F_palloc0(m, int32(144))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(195726186)
			v26 = F_copyObjectImpl(m, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v26
				v31 = F_pstrdup(m, l1)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v31
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l2
					v38 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v22)+20)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+28)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+36)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+41)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+60)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v13
					*(*int64)(unsafe.Add(mBase, uint32(v22)+68)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+76)) = v38
					*(*uint16)(unsafe.Add(mBase, uint32(v22)+84)) = uint16(v35)
					*(*int64)(unsafe.Add(mBase, uint32(v22)+88)) = v38
					*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v22)+120)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = int64(-4616189618054758400)
					*(*int64)(unsafe.Add(mBase, uint32(v22)+128)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(v22)+136)) = v38
					*(*int32)(unsafe.Add(mBase, _c_F_CreateCachedPlan[0])) = v18
					return v22
				}
			}
		}
	}
}
func F_cached_function_compile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v266 int64
	_ = v266
	var v269 int32
	_ = v269
	var v282 int32
	_ = v282
	var v285 int64
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v436 int32
	_ = v436
	var v438 int64
	_ = v438
	var v441 int32
	_ = v441
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int64
	_ = v457
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v657 int32
	_ = v657
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v704 int32
	_ = v704
	var v715 int32
	_ = v715
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
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
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v771 int32
	_ = v771
	var v772 int64
	_ = v772
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
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
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	v8 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(688)
	m.G0 = v25
	v36 = int32(-1)
	v37 = v8
	v38 = v8
	v39 = v8
	v40 = v8
	v41 = v8
	v42 = v8
	v43 = v8
	v44 = v8
	v45 = v8
	v47 = v8
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	m.G0 = v25 + int32(688)
	return v731
L4:
	;
	if v36 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	goto L3
L6:
	;
	v771 = int32(m.ExcTag)
	v772 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v771 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v732
	v753 = v738 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v753)
	F_ReleaseCatCache(m, v732)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L6
	} else {
		goto L134
	}
L8:
	;
	v730 = v44
	v731 = v728
	v732 = v66
	v733 = v111
	v734 = v113
	v735 = v41
	v736 = v42
	v737 = v43
	v738 = v47
	v739 = v45
	goto L7
L9:
	;
	if v537 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L10:
	;
	v536 = v44
	v537 = v37
	v538 = v38
	v539 = v39
	v540 = v40
	v541 = v41
	v542 = v42
	v543 = v47
	v544 = v45
	goto L9
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	v59 = v47 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v38
	v66 = F_SearchSysCache1(m, int32(47), v53)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v66 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v111 = v66 + int32(4)
	v113 = v66 + int32(16)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+22)))
	v116 = v114 + v115
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v53
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_0), v25)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(501), int32(_a_F_cached_function_compile_2))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0]))
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1]))
	goto L106
L21:
	;
	v513 = int32(0)
	if l4 == v513 {
		v518 = v509
		v520 = v513
		v521 = v45
		goto L20
	} else {
		goto L105
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	v501 = int32(1)
	v503 = v47 & v501
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v503)
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[2]))
	v507 = F_MemoryContextAllocZero(m, v506, l4)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L104
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v129 = v25 + int32(172)
	F_compute_function_hashkey(m, l0, v116, v129, l4, l5, l6)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	v155 = v114
	v156 = l1
	goto L25
L25:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v157 == v158 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v133 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v145 = int32(0)
	v147 = F_hash_search(m, v133, v129, v145, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v147 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+428))
	if v151 == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v155 = v154
	v156 = v151
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v170 = v156 + int32(8)
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+2)))
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v173 = int32(16)
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+2)))
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111))))
	if v171|v172<<(uint(v173)%32) == v176|v177<<(uint(v173)%32) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v188 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	if v187 != 0 {
		v728 = v156
		goto L8
	} else {
		goto L40
	}
L35:
	;
	goto L34
L36:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+4)))
	if v183 == v184 {
		v187 = int32(1)
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v187 = int32(0)
	goto L35
L39:
	;
	goto L38
L40:
	;
	goto L33
L41:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v156)+24))
	if v266 == int64(0) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	v205 = F_hash_search(m, v202, v188, int32(2), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L44
	}
L43:
	;
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v249
	if v191 == v249 {
		goto L41
	} else {
		goto L50
	}
L44:
	;
	if v205 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v218 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v218 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_3), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(229), int32(_a_F_cached_function_compile_4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L43
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_FreeTupleDesc(m, v191)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	goto L41
L52:
	;
	if v488 != 0 {
		v509 = v487
		v510 = v156
		goto L21
	} else {
		goto L103
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_compute_function_hashkey(m, l0, v116, v25+int32(172), l4, l5, l6)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L6
	} else {
		goto L102
	}
L54:
	;
	v468 = int32(0)
	if l1 == v468 {
		v487 = v468
		v488 = v468
		goto L52
	} else {
		goto L101
	}
L55:
	;
	v467 = int32(1)
	if l1 != 0 {
		v472 = v156
		v473 = v467
		goto L53
	} else {
		goto L100
	}
L56:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	if v269 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v288 = int32(0)
	goto L58
L58:
	;
	v289 = int32(0)
	if base.B2i32(l1 == v289)|v288 == v289 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	m.T0[v269].(func(*base.Module, int32))(m, v156)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+16)) = int32(0)
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v156)+24))
	v288 = base.B2i32(v285 == int64(0))
	goto L58
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v304 = v25 + int32(172)
	F_compute_function_hashkey(m, l0, v116, v304, l4, l5, l6)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v288 == int32(0) {
		goto L54
	} else {
		goto L99
	}
L64:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v308 == int32(0) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v320 = int32(0)
	v322 = F_hash_search(m, v308, v304, v320, v320)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v322 == int32(0) {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v322)+428))
	if v326 == int32(0) {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v329 == v331 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v343 = v326 + int32(8)
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+2)))
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343))))
	v346 = int32(16)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+2)))
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111))))
	if v344|v345<<(uint(v346)%32) == v349|v350<<(uint(v346)%32) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	goto L71
L71:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	if v361 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v360 != 0 {
		v728 = v326
		goto L8
	} else {
		goto L78
	}
L73:
	;
	goto L72
L74:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v343)+4)))
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+4)))
	if v356 == v357 {
		v360 = int32(1)
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v360 = int32(0)
	goto L73
L77:
	;
	goto L76
L78:
	;
	goto L71
L79:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v326)+24))
	if v438 != int64(0) {
		goto L22
	} else {
		goto L90
	}
L80:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	v378 = F_hash_search(m, v375, v361, int32(2), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L82
	}
L81:
	;
	v422 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = v422
	if v364 == v422 {
		goto L79
	} else {
		goto L88
	}
L82:
	;
	if v378 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	v391 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v391 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_3), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(229), int32(_a_F_cached_function_compile_4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	F_FreeTupleDesc(m, v364)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	goto L79
L90:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v326)+16))
	if v441 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v509 = v326
	v510 = v326
	goto L21
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v66
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v59)
	m.T0[v441].(func(*base.Module, int32))(m, v326)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v455 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v326)+16)) = v455
	v457 = *(*int64)(unsafe.Add(mBase, uint32(v326)+24))
	v459 = base.B2i32(v457 == int64(0))
	if v459 == v455 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	if v457 == int64(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v463 = v326
	goto L98
L97:
	;
	v463 = int32(0)
	goto L98
L98:
	;
	v509 = v463
	v510 = v326
	goto L21
L99:
	;
	goto L55
L100:
	;
	v487 = v156
	v488 = v467
	goto L52
L101:
	;
	v472 = v468
	v473 = v468
	goto L53
L102:
	;
	v487 = v472
	v488 = v473
	goto L52
L103:
	;
	goto L22
L104:
	;
	v518 = v507
	v520 = v501
	v521 = v507
	goto L20
L105:
	;
	base.MemoryFill(m, v510, int32(0), l4)
	v518 = v509
	v520 = v513
	v521 = v45
	goto L20
L106:
	;
	v530 = v25 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v530)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v530))) = v25 + int32(12)
	goto L109
L107:
	;
	v536 = v518
	v537 = int32(0)
	v538 = v66
	v539 = v111
	v540 = v113
	v541 = v526
	v542 = v528
	v543 = v520
	v544 = v521
	goto L9
L109:
	;
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v25 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	v563 = v543 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	m.T0[l2].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v538, v25+int32(172), v536, l6)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0])) = v541
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v542
	v704 = v543 & int32(1)
	if v704 != 0 {
		goto L129
	} else {
		goto L130
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0])) = v541
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v542
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v573)))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+4)) = v574
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v539)))
	*(*int32)(unsafe.Add(mBase, uint32(v536)+8)) = v576
	v578 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v539)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v536)+12)) = uint16(v578)
	*(*int32)(unsafe.Add(mBase, uint32(v536)+16)) = l3
	v582 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v582 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+632)) = int32(1578)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+628)) = int32(1579)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+620)) = int64(1855425872300)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	v606 = F_hash_create(m, int32(_a_F_cached_function_compile_5), int32(128), v25+int32(604), int32(200))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	v609 = v582
	v610 = v43
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	v625 = F_hash_search(m, v609, v25+int32(172), int32(1), v25+int32(603))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3])) = v606
	v609 = v606
	v610 = v606
	goto L116
L118:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+603)))
	if v627 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v25)+196))
	if v672 != 0 {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	v641 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	if v641 == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_6), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(181), int32(_a_F_cached_function_compile_7))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L119
L125:
	;
	v673 = int32(_a_F_cached_function_compile_8)
	v674 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4]))
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4])) = v677
	*(*int32)(unsafe.Add(mBase, uint32(v625)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v610
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v563)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v25)+196))
	v691 = F_CreateTupleDescCopy(m, v690)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L6
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v625)+428)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v536))) = v625
	v730 = v536
	v731 = v536
	v732 = v538
	v733 = v539
	v734 = v540
	v735 = v541
	v736 = v542
	v737 = v610
	v738 = v543
	v739 = v544
	goto L7
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v625)+24)) = v691
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4])) = v674
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v704)
	F_pfree(m, v536)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+656)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+652)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+660)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+664)) = v536
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v544
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+684)) = v538
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)) = uint8(v704)
	F_pg_re_throw(m)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L6
	} else {
		goto L133
	}
L132:
	;
	goto L131
L133:
	;
	goto L1
L134:
	;
	goto L5
L135:
	;
	v776 = int32(v772)
	m.G0 = v25
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v776)+4))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	if v25+int32(12) == v782 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	m.ExcPending = 1
	goto L144
L137:
	;
	if v786 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v786 = v784
	goto L140
L139:
	;
	v786 = int32(0)
	goto L140
L140:
	;
	goto L137
L141:
	;
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v25)+684))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v25)+680))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v25)+676))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v25)+672))
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+671)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v25)+664))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v25)+660))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v25)+656))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v25)+652))
	v36 = v786
	v37 = v778
	v38 = v787
	v39 = v789
	v40 = v788
	v41 = v794
	v42 = v793
	v43 = v795
	v44 = v792
	v45 = v790
	v47 = v791
	goto L2
L142:
	;
	goto L143
L143:
	;
	F___wasm_longjmp(m, v779, v778)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	return int32(0)
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
