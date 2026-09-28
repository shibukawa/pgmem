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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v408 int32
	_ = v408
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v437 int32
	_ = v437
	var v439 int64
	_ = v439
	var v442 int32
	_ = v442
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int64
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
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
	var v545 int32
	_ = v545
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v705 int32
	_ = v705
	var v716 int32
	_ = v716
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
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
	var v740 int32
	_ = v740
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v772 int32
	_ = v772
	var v773 int64
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
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
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	v8 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(704)
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
	m.G0 = v25 + int32(704)
	return v732
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
	v772 = int32(m.ExcTag)
	v773 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v772 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v736
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v738
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v740
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v733
	v754 = v739 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v754)
	F_ReleaseCatCache(m, v733)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L6
	} else {
		goto L134
	}
L8:
	;
	v731 = v44
	v732 = v729
	v733 = v67
	v734 = v112
	v735 = v114
	v736 = v41
	v737 = v42
	v738 = v43
	v739 = v47
	v740 = v45
	goto L7
L9:
	;
	if v538 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L10:
	;
	v537 = v44
	v538 = v37
	v539 = v38
	v540 = v39
	v541 = v40
	v542 = v41
	v543 = v42
	v544 = v47
	v545 = v45
	goto L9
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	v59 = v47 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v38
	v67 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(v53))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v67 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v112 = v67 + int32(4)
	v114 = v67 + int32(16)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+22)))
	v117 = v115 + v116
	if l1 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v53
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_0), v25)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(515), int32(_a_F_cached_function_compile_2))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0]))
	v529 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1]))
	goto L106
L21:
	;
	v514 = int32(0)
	if l4 == v514 {
		v519 = v510
		v521 = v514
		v522 = v45
		goto L20
	} else {
		goto L105
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	v502 = int32(1)
	v504 = v47 & v502
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v504)
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[2]))
	v508 = F_MemoryContextAllocZero(m, v507, l4)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L6
	} else {
		goto L104
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v130 = v25 + int32(184)
	F_compute_function_hashkey(m, l0, v117, v130, l4, l5, l6)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	v156 = v115
	v157 = l1
	goto L25
L25:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v158 == v159 {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v134 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v146 = int32(0)
	v148 = F_hash_search(m, v134, v130, v146, v146)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if v148 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+428))
	if v152 == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v156 = v155
	v157 = v152
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v171 = v157 + int32(8)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171))))
	v174 = int32(16)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
	if v172|v173<<(uint(v174)%32) == v177|v178<<(uint(v174)%32) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v189 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L34:
	;
	if v188 != 0 {
		v729 = v157
		goto L8
	} else {
		goto L40
	}
L35:
	;
	goto L34
L36:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+4)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	if v184 == v185 {
		v188 = int32(1)
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v188 = int32(0)
	goto L35
L39:
	;
	goto L38
L40:
	;
	goto L33
L41:
	;
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
	if v267 == int64(0) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	v206 = F_hash_search(m, v203, v189, int32(2), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L44
	}
L43:
	;
	v250 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = v250
	if v192 == v250 {
		goto L41
	} else {
		goto L50
	}
L44:
	;
	if v206 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v219 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v219 == int32(0) {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_3), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(229), int32(_a_F_cached_function_compile_4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L43
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_FreeTupleDesc(m, v192)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	goto L41
L52:
	;
	if v489 != 0 {
		v510 = v488
		v511 = v157
		goto L21
	} else {
		goto L103
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_compute_function_hashkey(m, l0, v117, v25+int32(184), l4, l5, l6)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L6
	} else {
		goto L102
	}
L54:
	;
	v469 = int32(0)
	if l1 == v469 {
		v488 = v469
		v489 = v469
		goto L52
	} else {
		goto L101
	}
L55:
	;
	v468 = int32(1)
	if l1 != 0 {
		v473 = v157
		v474 = v468
		goto L53
	} else {
		goto L100
	}
L56:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	if v270 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	v289 = int32(0)
	goto L58
L58:
	;
	v290 = int32(0)
	if base.B2i32(l1 == v290)|v289 == v290 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	m.T0[v270].(func(*base.Module, int32))(m, v157)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+16)) = int32(0)
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v157)+24))
	v289 = base.B2i32(v286 == int64(0))
	goto L58
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v305 = v25 + int32(184)
	F_compute_function_hashkey(m, l0, v117, v305, l4, l5, l6)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v289 == int32(0) {
		goto L54
	} else {
		goto L99
	}
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v309 == int32(0) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v321 = int32(0)
	v323 = F_hash_search(m, v309, v305, v321, v321)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v323 == int32(0) {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+428))
	if v327 == int32(0) {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	if v330 == v332 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v344 = v327 + int32(8)
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+2)))
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344))))
	v347 = int32(16)
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
	if v345|v346<<(uint(v347)%32) == v350|v351<<(uint(v347)%32) {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	goto L71
L71:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	if v362 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	if v361 != 0 {
		v729 = v327
		goto L8
	} else {
		goto L78
	}
L73:
	;
	goto L72
L74:
	;
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344)+4)))
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)))
	if v357 == v358 {
		v361 = int32(1)
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v361 = int32(0)
	goto L73
L77:
	;
	goto L76
L78:
	;
	goto L71
L79:
	;
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v327)+24))
	if v439 != int64(0) {
		goto L22
	} else {
		goto L90
	}
L80:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v362)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	v379 = F_hash_search(m, v376, v362, int32(2), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L6
	} else {
		goto L82
	}
L81:
	;
	v423 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v423
	if v365 == v423 {
		goto L79
	} else {
		goto L88
	}
L82:
	;
	if v379 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	v392 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v392 == int32(0) {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_3), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(229), int32(_a_F_cached_function_compile_4))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L81
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	F_FreeTupleDesc(m, v365)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	goto L79
L90:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v327)+16))
	if v442 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v510 = v327
	v511 = v327
	goto L21
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v67
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v59)
	m.T0[v442].(func(*base.Module, int32))(m, v327)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v456 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = v456
	v458 = *(*int64)(unsafe.Add(mBase, uint32(v327)+24))
	v460 = base.B2i32(v458 == int64(0))
	if v460 == v456 {
		goto L22
	} else {
		goto L95
	}
L95:
	;
	if v458 == int64(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v464 = v327
	goto L98
L97:
	;
	v464 = int32(0)
	goto L98
L98:
	;
	v510 = v464
	v511 = v327
	goto L21
L99:
	;
	goto L55
L100:
	;
	v488 = v157
	v489 = v468
	goto L52
L101:
	;
	v473 = v469
	v474 = v469
	goto L53
L102:
	;
	v488 = v473
	v489 = v474
	goto L52
L103:
	;
	goto L22
L104:
	;
	v519 = v508
	v521 = v502
	v522 = v508
	goto L20
L105:
	;
	base.MemoryFill(m, v511, int32(0), l4)
	v519 = v510
	v521 = v514
	v522 = v45
	goto L20
L106:
	;
	v531 = v25 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v531)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v531))) = v25 + int32(12)
	goto L109
L107:
	;
	v537 = v519
	v538 = int32(0)
	v539 = v67
	v540 = v112
	v541 = v114
	v542 = v527
	v543 = v529
	v544 = v521
	v545 = v522
	goto L9
L109:
	;
	goto L107
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v25 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	v564 = v544 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	m.T0[l2].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v539, v25+int32(184), v537, l6)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0])) = v542
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v543
	v705 = v544 & int32(1)
	if v705 != 0 {
		goto L129
	} else {
		goto L130
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[0])) = v542
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[1])) = v543
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v541)))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	*(*int32)(unsafe.Add(mBase, uint32(v537)+4)) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	*(*int32)(unsafe.Add(mBase, uint32(v537)+8)) = v577
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v540)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v537)+12)) = uint16(v579)
	*(*int32)(unsafe.Add(mBase, uint32(v537)+16)) = l3
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3]))
	if v583 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+636)) = int32(1791)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+632)) = int32(1792)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+624)) = int64(1855425872300)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	v607 = F_hash_create(m, int32(_a_F_cached_function_compile_5), int64(128), v25+int32(616), int32(200))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	v610 = v583
	v611 = v43
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	v626 = F_hash_search(m, v610, v25+int32(184), int32(1), v25+int32(615))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L6
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[3])) = v607
	v610 = v607
	v611 = v607
	goto L116
L118:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+615)))
	if v628 != int32(1) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v25)+208))
	if v673 != 0 {
		goto L125
	} else {
		goto L126
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	v642 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L6
	} else {
		goto L121
	}
L121:
	;
	if v642 == int32(0) {
		goto L119
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	F_errmsg_internal(m, int32(_a_F_cached_function_compile_6), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	F_errfinish(m, int32(_a_F_cached_function_compile_1), int32(181), int32(_a_F_cached_function_compile_7))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L119
L125:
	;
	v674 = int32(_a_F_cached_function_compile_8)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4]))
	v678 = *(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4])) = v678
	*(*int32)(unsafe.Add(mBase, uint32(v626)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v564)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v25)+208))
	v692 = F_CreateTupleDescCopy(m, v691)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L6
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626)+428)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v537))) = v626
	v731 = v537
	v732 = v537
	v733 = v539
	v734 = v540
	v735 = v541
	v736 = v542
	v737 = v543
	v738 = v611
	v739 = v544
	v740 = v545
	goto L7
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v626)+24)) = v692
	*(*int32)(unsafe.Add(mBase, _c_F_cached_function_compile[4])) = v675
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v705)
	F_pfree(m, v537)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L6
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+672)) = v542
	*(*int32)(unsafe.Add(mBase, uint32(v25)+668)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v25)+676)) = v543
	*(*int32)(unsafe.Add(mBase, uint32(v25)+680)) = v537
	*(*int32)(unsafe.Add(mBase, uint32(v25)+688)) = v545
	*(*int32)(unsafe.Add(mBase, uint32(v25)+692)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v25)+696)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v25)+700)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)) = uint8(v705)
	F_pg_re_throw(m)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
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
	v777 = int32(v773)
	m.G0 = v25
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v777)+4))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v777)))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v780)))
	if v25+int32(12) == v783 {
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
	if v787 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v780)+4))
	v787 = v785
	goto L140
L139:
	;
	v787 = int32(0)
	goto L140
L140:
	;
	goto L137
L141:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v25)+700))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v25)+696))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v25)+692))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v25)+688))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+687)))
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v25)+680))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v25)+676))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v25)+672))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v25)+668))
	v36 = v787
	v37 = v779
	v38 = v788
	v39 = v790
	v40 = v789
	v41 = v795
	v42 = v794
	v43 = v796
	v44 = v793
	v45 = v791
	v47 = v792
	goto L2
L142:
	;
	goto L143
L143:
	;
	F___wasm_longjmp(m, v780, v779)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
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
