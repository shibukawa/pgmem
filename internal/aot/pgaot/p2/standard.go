package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standard_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v8 = int32(_a_F_standard_ExecutorFinish_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0])) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_InstrStart(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+148))
	if v19 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+128)))
	if v70&int32(32) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v30 = int32(0)
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30<<(uint(int32(2))%32))))
	goto L11
L10:
	;
	goto L6
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+152))
	if v44 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v60 = v30 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v60 < v61 {
		v30 = v60
		goto L9
	} else {
		goto L26
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	F_MemoryContextReset(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	if v48 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	F_ExecReScan(m, v36)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v52 = m.T0[v51].(func(*base.Module, int32) int32)(m, v36)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v54&int32(2) == int32(0) {
		goto L11
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L12
L25:
	;
	goto L24
L26:
	;
	goto L10
L27:
	;
	F_AfterTriggerEndQuery(m, v11)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v77 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_InstrStop(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_standard_ExecutorFinish[0])) = v9
	v82 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+136)) = uint8(v82)
	return
L34:
	;
	goto L33
}
func F_standard_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 float64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 float64
	_ = v112
	var v122 float64
	_ = v122
	var v124 float64
	_ = v124
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v180 int32
	_ = v180
	var v184 int64
	_ = v184
	var v187 int32
	_ = v187
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 float64
	_ = v246
	var v253 float64
	_ = v253
	var v259 float64
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v309 float64
	_ = v309
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v320 float64
	_ = v320
	var v321 float64
	_ = v321
	var v324 float64
	_ = v324
	var v325 float64
	_ = v325
	var v326 float64
	_ = v326
	var v329 float64
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v414 float64
	_ = v414
	var v416 float64
	_ = v416
	var v420 float64
	_ = v420
	var v421 float64
	_ = v421
	var v423 float64
	_ = v423
	var v427 float64
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v446 float64
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 float64
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 float64
	_ = v481
	var v482 float64
	_ = v482
	var v485 int32
	_ = v485
	var v486 float64
	_ = v486
	var v487 float64
	_ = v487
	var v489 float64
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v514 float64
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 float64
	_ = v520
	var v521 float64
	_ = v521
	var v524 int32
	_ = v524
	var v533 int32
	_ = v533
	var v539 float64
	_ = v539
	var v542 int32
	_ = v542
	var v544 float64
	_ = v544
	var v545 float64
	_ = v545
	var v548 float64
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int64
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v785 int32
	_ = v785
	var v792 int32
	_ = v792
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
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v856 float64
	_ = v856
	var v861 float64
	_ = v861
	var v865 int32
	_ = v865
	var v869 float64
	_ = v869
	var v874 float64
	_ = v874
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 float64
	_ = v883
	var v888 float64
	_ = v888
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	v13 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v22 = F_palloc0(m, int32(128))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+112)) = v26
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(268)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v22)+86)) = v26
	if l2&int32(2048) == v28 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	v103 = int32(0)
	v105 = v100 & base.B2i32(v102 != v103)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+95)) = uint8(v105)
	if l2&int32(256) == v103 {
		v122 = v13
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+94)) = uint8(v95)
	v97 = int32(117)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+96)) = uint8(v97)
	v100 = int32(0)
	goto L3
L5:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[1])))
	if v58&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v63 != int32(1) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	if v66 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[2]))
	if v68 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[3]))
	if int32(0) <= v72 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v75 = m.G0
	v77 = v75 - int32(16)
	m.G0 = v77
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = int32(0)
	v81 = int32(_a_F_standard_planner_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+8)) = uint16(v81)
	v85 = F_max_parallel_hazard_walker(m, l0, v77+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v87 = int32(*(*int8)(unsafe.Add(mBase, uint32(v77)+8)))
	m.G0 = v77 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+96)) = uint8(v87)
	v93 = base.B2i32(v87 != int32(117))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+94)) = uint8(v93)
	v100 = v93
	goto L3
L12:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[4])))
	if v129 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v122
	v124 = v122
	goto L12
L14:
	;
	v112 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[5]))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+24)) = v112
	if base.F64_ge(v112, float64(1)) != 0 {
		v122 = v13
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.F64_le(v112, float64(0)) == int32(0) {
		v124 = v112
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v122 = float64(1e-10)
	goto L13
L17:
	;
	v130 = int64(290864)
	goto L19
L18:
	;
	v130 = int64(290848)
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v130
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[6])))
	if v133 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v137 = v130 | int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v137
	v139 = v137
	goto L22
L21:
	;
	v139 = v130
	goto L22
L22:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[7])))
	if v141 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v145 = v139 | int64(6)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v145
	v147 = v145
	goto L25
L24:
	;
	v147 = v139
	goto L25
L25:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[8])))
	if v149 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v153 = v147 | int64(65536)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v153
	v155 = v153
	goto L28
L27:
	;
	v155 = v147
	goto L28
L28:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[9])))
	if v157 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v161 = v155 | int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v161
	v163 = v161
	goto L31
L30:
	;
	v163 = v155
	goto L31
L31:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[10])))
	if v165 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[11])))
	if v180 != int32(1) {
		v202 = v178
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v178 = v163
	goto L32
L34:
	;
	goto L35
L35:
	;
	v169 = v163 | int64(64)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v169
	v172 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[12])))
	if v172 != int32(1) {
		v178 = v169
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v176 = v163 | int64(192)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v176
	v178 = v176
	goto L32
L37:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[13])))
	if v204 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v184 = v178 | int64(256)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v184
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[12])))
	if v187 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v191 = v178 | int64(768)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v191
	v193 = v191
	goto L41
L40:
	;
	v193 = v184
	goto L41
L41:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[14])))
	if v195 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v202 = v193
	goto L37
L43:
	;
	goto L44
L44:
	;
	v199 = v193 | int64(1024)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v199
	v202 = v199
	goto L37
L45:
	;
	v208 = v202 | int64(2048)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v208
	v210 = v208
	goto L47
L46:
	;
	v210 = v202
	goto L47
L47:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[15])))
	if v212 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v216 = v210 | int64(32768)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v216
	v218 = v216
	goto L50
L49:
	;
	v218 = v210
	goto L50
L50:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[16])))
	if v220 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+104)) = v218 | int64(131072)
	goto L53
L52:
	;
	goto L53
L53:
	;
	v226 = int32(0)
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[17]))
	if v231 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	m.T0[v231].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v22, l0, l1, l2, v19+int32(24), l4)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	v237 = v124
	goto L56
L56:
	;
	v239 = F_subquery_planner(m, v22, l0, v226, v226, v226, v226, v237, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v236 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v237 = v236
	goto L56
L58:
	;
	v243 = F_fetch_upper_rel(m, v239, int32(7), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v243)+60))
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	if base.F64_le(v246, float64(0)) != 0 {
		v351 = v245
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v364 = F_create_plan(m, v239, v351)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L94
	}
L61:
	;
	if base.F64_ge(v246, float64(1)) == int32(0) {
		v259 = v246
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v243)+44))
	if v261 == int32(0) {
		v351 = v245
		goto L60
	} else {
		goto L65
	}
L63:
	;
	v253 = *(*float64)(unsafe.Add(mBase, uint32(v245)+32))
	if base.F64_gt(v253, float64(0)) == int32(0) {
		v259 = v246
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v259 = base.F64_div(v246, v253)
	goto L62
L65:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v264 <= int32(0) {
		v351 = v245
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v271 = v245
	v274 = int32(0)
	v279 = v264
	goto L67
L67:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v284+v274<<(uint(int32(2))%32))))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	if v289 != 0 {
		v343 = v271
		v344 = v279
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v351 = v343
	goto L60
L69:
	;
	v346 = v274 + int32(1)
	if v346 < v344 {
		v271 = v343
		v274 = v346
		v279 = v344
		goto L67
	} else {
		goto L93
	}
L70:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v243)+60))
	if v288 == v290 {
		v343 = v271
		v344 = v279
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v271)+40))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v288)+40))
	if v295 != v296 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v338 <= int32(0) {
		goto L90
	} else {
		goto L91
	}
L73:
	;
	if v295 < v296 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	if base.F64_le(v259, float64(0))|base.F64_ge(v259, float64(1)) != 0 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v301 = int32(-1)
	goto L78
L77:
	;
	v301 = int32(1)
	goto L78
L78:
	;
	v338 = v301
	goto L72
L79:
	;
	v338 = v334
	goto L72
L80:
	;
	v307 = int32(-1)
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v271)+56))
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v288)+56))
	if base.F64_lt(v308, v309) != 0 {
		v334 = v307
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v271)+56))
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v271)+48))
	v324 = base.F64_add(base.F64_mul(v259, base.F64_sub(v320, v321)), v321)
	v325 = *(*float64)(unsafe.Add(mBase, uint32(v288)+56))
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v288)+48))
	v329 = base.F64_add(base.F64_mul(v259, base.F64_sub(v325, v326)), v326)
	if base.F64_lt(v324, v329) != 0 {
		v334 = int32(-1)
		goto L79
	} else {
		goto L89
	}
L83:
	;
	if base.F64_gt(v308, v309) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v338 = int32(1)
	goto L72
L85:
	;
	goto L86
L86:
	;
	v313 = *(*float64)(unsafe.Add(mBase, uint32(v271)+48))
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v288)+48))
	if base.F64_lt(v313, v314) != 0 {
		v334 = v307
		goto L79
	} else {
		goto L87
	}
L87:
	;
	if base.F64_gt(v313, v314) != 0 {
		v334 = int32(1)
		goto L79
	} else {
		goto L88
	}
L88:
	;
	v338 = int32(0)
	goto L72
L89:
	;
	v334 = base.F64_lt(v329, v324)
	goto L79
L90:
	;
	v341 = v271
	goto L92
L91:
	;
	v341 = v288
	goto L92
L92:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v343 = v341
	v344 = v342
	goto L69
L93:
	;
	goto L68
L94:
	;
	if l2&int32(2) == int32(0) {
		v374 = v364
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	if v376 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v370 = F_ExecSupportsBackwardScan(m, v364)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	if v370 != 0 {
		v374 = v364
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v372 = F_materialize_finished_plan(m, v364)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v374 = v372
	goto L95
L100:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	if v557 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L101:
	;
	v554 = v374
	goto L100
L102:
	;
	goto L103
L103:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+37)))
	if v379 != int32(1) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v554 = v374
	goto L100
L105:
	;
	goto L106
L106:
	;
	if v376 != int32(2) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v388 = F_palloc0(m, int32(88))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L110
	}
L108:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v374)+60))
	if v384 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v554 = v374
	goto L100
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v388))) = int32(372)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v374)+44))
	v393 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v388)+80)) = uint8(v393)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+72)) = v393
	v397 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+56)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v388)+52)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v388)+48)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v388)+44)) = v392
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v388)+81)) = uint8(base.B2i32(v404 == int32(2)))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v374)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v388)+60)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(v374)+60)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v388)+76)) = int32(-1)
	v414 = *(*float64)(unsafe.Add(mBase, uint32(v374)+8))
	v416 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[18]))
	*(*float64)(unsafe.Add(mBase, uint32(v388)+8)) = base.F64_add(v414, v416)
	v420 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[19]))
	v421 = *(*float64)(unsafe.Add(mBase, uint32(v374)+24))
	v423 = *(*float64)(unsafe.Add(mBase, uint32(v374)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v388)+16)) = base.F64_add(base.F64_mul(v420, v421), base.F64_add(v416, v423))
	v427 = *(*float64)(unsafe.Add(mBase, uint32(v374)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v388)+24)) = v427
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v374)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v388)+36)) = uint16(v397)
	*(*int32)(unsafe.Add(mBase, uint32(v388)+32)) = v429
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v388)+60))
	v446 = float64(0)
	if v433 == v397 {
		v533 = v397
		v539 = v446
		goto L112
	} else {
		goto L113
	}
L111:
	;
	v544 = *(*float64)(unsafe.Add(mBase, uint32(v374)+8))
	v545 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+8)) = base.F64_sub(v544, v545)
	v548 = *(*float64)(unsafe.Add(mBase, uint32(v374)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v374)+16)) = base.F64_sub(v548, v545)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v552 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v551)+95)) = uint8(v552)
	v554 = v388
	goto L100
L112:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v19+int32(16)))) = v539
	v542 = v533 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(15)))) = uint8(v542)
	goto L111
L113:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	if v449 <= int32(0) {
		v533 = v397
		v539 = v446
		goto L112
	} else {
		goto L114
	}
L114:
	;
	if v449 == int32(1) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v515+v506<<(uint(int32(2))%32))))
	v520 = *(*float64)(unsafe.Add(mBase, uint32(v519)+56))
	v521 = *(*float64)(unsafe.Add(mBase, uint32(v519)+64))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+39)))
	v533 = v524 ^ int32(1) | v508
	v539 = base.F64_add(v514, base.F64_add(v520, v521))
	goto L112
L116:
	;
	v506 = int32(0)
	v508 = v397
	v514 = v446
	goto L115
L117:
	;
	goto L118
L118:
	;
	v455 = int32(0)
	if v455 < v449 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v458 = v449
	goto L121
L120:
	;
	v458 = v455
	goto L121
L121:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v433)+12))
	v468 = int32(0)
	v470 = v397
	v475 = v397
	v476 = v446
	goto L122
L122:
	;
	v477 = int32(2)
	v479 = v463 + v468<<(uint(v477)%32)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	v481 = *(*float64)(unsafe.Add(mBase, uint32(v480)+56))
	v482 = *(*float64)(unsafe.Add(mBase, uint32(v480)+64))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479)+4))
	v486 = *(*float64)(unsafe.Add(mBase, uint32(v485)+56))
	v487 = *(*float64)(unsafe.Add(mBase, uint32(v485)+64))
	v489 = base.F64_add(base.F64_add(v476, base.F64_add(v481, v482)), base.F64_add(v486, v487))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+39)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480)+39)))
	v495 = base.B2i32(v490&v491 == int32(0)) | v470
	v497 = v468 + v477
	v499 = v475 + v477
	if v499 != v458&int32(2147483646) {
		v468 = v497
		v470 = v495
		v475 = v499
		v476 = v489
		goto L122
	} else {
		goto L124
	}
L123:
	;
	if v458&int32(1) == int32(0) {
		v533 = v495
		v539 = v489
		goto L112
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v506 = v497
	v508 = v495
	v514 = v489
	goto L115
L126:
	;
	v628 = F_set_plan_references(m, v239, v554)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L140
	}
L127:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v569 = int32(0)
	goto L128
L128:
	;
	v579 = int32(0)
	if v561 == v579 {
		v589 = v579
		goto L130
	} else {
		goto L131
	}
L130:
	;
	if v560 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v583 <= v569 {
		v589 = int32(0)
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v589 = v585 + v569<<(uint(int32(2))%32)
	goto L130
L133:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v597+v569<<(uint(int32(2))%32))))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	F_SS_finalize_plan(m, v604, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L139
	}
L134:
	;
	F_SS_finalize_plan(m, v239, v554)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	if base.B2i32(v589 == int32(0))|base.B2i32(v594 <= v569) != 0 {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	if v597 != 0 {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	goto L134
L138:
	;
	goto L126
L139:
	;
	v569 = v569 + int32(1)
	goto L128
L140:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v639 = int32(0)
	goto L141
L141:
	;
	v649 = int32(0)
	if v631 == v649 {
		v659 = v649
		goto L143
	} else {
		goto L144
	}
L143:
	;
	if v630 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v631)+4))
	if v653 <= v639 {
		v659 = int32(0)
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v631)+12))
	v659 = v655 + v639<<(uint(int32(2))%32)
	goto L143
L146:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v667+v639<<(uint(int32(2))%32))))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	v929 = F_set_plan_references(m, v927, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L192
	}
L147:
	;
	v670 = F_palloc0(m, int32(120))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L151
	}
L148:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if base.B2i32(v659 == int32(0))|base.B2i32(v664 <= v639) != 0 {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v630)+12))
	if v667 != 0 {
		goto L146
	} else {
		goto L150
	}
L150:
	;
	goto L147
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670))) = int32(334)
	v674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+4)) = v674
	v676 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+24)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v670)+8)) = v676
	v680 = int32(0)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+28)) = uint8(base.B2i32(v681 != v680))
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+29)) = uint8(v685)
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+30)) = uint8(v687)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+31)) = uint8(v689)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+93)))
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+32)) = uint8(v691)
	v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+95)))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+40)) = v628
	*(*uint8)(unsafe.Add(mBase, uint32(v670)+33)) = uint8(v693)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v22)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+44)) = v696
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+48)) = v698
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v22)+36))
	v702 = F_bms_difference(m, v700, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670)+52)) = v702
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+56)) = v705
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+72)) = v707
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+64)) = v709
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+68)) = v711
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+76)) = v713
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+80)) = v715
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	if v717 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if int32(0) < v718 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	v771 = v715
	goto L155
L155:
	;
	if v771 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L156:
	;
	v724 = v680
	goto L159
L157:
	;
	goto L158
L158:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(48))))
	v771 = v766
	goto L155
L159:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v670)+60))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v717)+12))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v738+v724<<(uint(int32(2))%32))))
	v743 = F_bms_add_member(m, v737, v742)
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	goto L158
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670)+60)) = v743
	v747 = v724 + int32(1)
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if v747 < v748 {
		v724 = v747
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+88)) = v835
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v22)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+92)) = v837
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+96)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+100)) = v841
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+104)) = v843
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+112)) = v845
	v847 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+116)) = v847
	v852 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[20])))
	if v852 != int32(1) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	if v785 <= int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v792 = int32(0)
	goto L166
L166:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v670)+84))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v771)+12))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v806+v792<<(uint(int32(2))%32))))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	v812 = F_bms_add_member(m, v805, v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L168
	}
L167:
	;
	goto L163
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670)+84)) = v812
	v816 = v792 + int32(1)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v771)+4))
	if v816 < v817 {
		v792 = v816
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v914 = *(*int32)(unsafe.Add(mBase, _c_F_standard_planner[21]))
	if v914 != 0 {
		goto L184
	} else {
		goto L185
	}
L171:
	;
	v856 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[22]))
	if base.F64_ge(v856, float64(0)) == int32(0) {
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v861 = *(*float64)(unsafe.Add(mBase, uint32(v628)+16))
	if base.F64_gt(v861, v856) == int32(0) {
		goto L170
	} else {
		goto L173
	}
L173:
	;
	v865 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = v865
	v869 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[23]))
	if base.F64_ge(v869, float64(0)) == int32(0) {
		v881 = v865
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v883 = *(*float64)(unsafe.Add(mBase, _c_F_standard_planner[24]))
	if base.F64_ge(v883, float64(0)) == int32(0) {
		v895 = v881
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v874 = *(*float64)(unsafe.Add(mBase, uint32(v628)+16))
	if base.F64_gt(v874, v869) == int32(0) {
		v881 = v865
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v878 = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = v878
	v881 = v878
	goto L174
L177:
	;
	v897 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[25])))
	if v897 == int32(1) {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	v888 = *(*float64)(unsafe.Add(mBase, uint32(v628)+16))
	if base.F64_gt(v888, v883) == int32(0) {
		v895 = v881
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v893 = v881 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = v893
	v895 = v893
	goto L177
L180:
	;
	v901 = v895 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = v901
	v903 = v901
	goto L182
L181:
	;
	v903 = v895
	goto L182
L182:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_standard_planner[26])))
	if v905 != int32(1) {
		goto L170
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v670)+36)) = v903 | int32(16)
	goto L170
L184:
	;
	m.T0[v914].(func(*base.Module, int32, int32, int32, int32))(m, v22, l0, l1, v670)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v22)+112))
	if v917 != 0 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	F_DestroyPartitionDirectory(m, v917)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	m.G0 = v19 + int32(32)
	return v670
L191:
	;
	goto L190
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v659))) = v929
	v639 = v639 + int32(1)
	goto L141
}
