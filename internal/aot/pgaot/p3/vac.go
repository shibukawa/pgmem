package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_update_datfrozenxid(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v85 int32
	_ = v85
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v328 int64
	_ = v328
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
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v430 int32
	_ = v430
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int64
	_ = v538
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v550 int32
	_ = v550
	var __phi550 int32
	_ = __phi550
	var v551 int32
	_ = v551
	var __phi551 int32
	_ = __phi551
	var v553 int32
	_ = v553
	var __phi553 int32
	_ = __phi553
	var v558 int32
	_ = v558
	var __phi558 int32
	_ = __phi558
	var v566 int64
	_ = v566
	var __phi566 int64
	_ = __phi566
	var v567 int64
	_ = v567
	var __phi567 int64
	_ = __phi567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int64
	_ = v584
	var v585 int64
	_ = v585
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int64
	_ = v599
	var v600 int64
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
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
	var v624 int32
	_ = v624
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int64
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v658 int64
	_ = v658
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int64
	_ = v679
	var v680 int64
	_ = v680
	var v686 int32
	_ = v686
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int64
	_ = v766
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int64
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v784 int64
	_ = v784
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int64
	_ = v805
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v817 int64
	_ = v817
	var v818 int32
	_ = v818
	var v820 int64
	_ = v820
	var v822 int32
	_ = v822
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v893 int32
	_ = v893
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int64
	_ = v960
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v984 int32
	_ = v984
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int64
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1019 int64
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1058 int64
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1083 int64
	_ = v1083
	var v1106 int32
	_ = v1106
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1181 int32
	_ = v1181
	var v1185 int32
	_ = v1185
	v1 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(96)
	m.G0 = v22
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(16908288)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = int64(0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v33
	v38 = F_LockAcquire(m, v26, int32(7), v1, v1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	m.G0 = v26 + int32(16)
	v44 = F_GetOldestNonRemovableTransactionId(m, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v48 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v50 = base.I32_wrap_i64(v48)
	v51 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v55 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	m.G0 = v22 + int32(96)
	return
L8:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_pfree(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L84
	}
L9:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	F_systable_inplace_update_finish(m, v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L83
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+88)) = v142
	v258 = v142
	goto L9
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L80
	}
L12:
	;
	F_systable_endscan(m, v62)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L78
	}
L13:
	;
	v57 = int32(0)
	v62 = F_systable_beginscan(m, v55, v57, v57, v57, v57, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v64 = F_systable_getnext(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v66 = v64
	v68 = v44
	v70 = v46
	goto L19
L17:
	;
	v140 = v44
	v142 = v46
	goto L18
L18:
	;
	F_systable_endscan(m, v62)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L50
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+22)))
	v87 = v85 + v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+136))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+140))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v90 == int32(114) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v140 = v134
	v142 = v135
	goto L18
L21:
	;
	v136 = F_systable_getnext(m, v62)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L48
	}
L22:
	;
	if v88 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v93 == int32(109) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+119)))
	if v96 != int32(116) {
		v134 = v68
		v135 = v70
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v88))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v124 = v68
	goto L28
L28:
	;
	if v89 == int32(0) {
		v134 = v124
		v135 = v70
		goto L21
	} else {
		goto L41
	}
L29:
	;
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L30:
	;
	v110 = base.B2i32(base.Ui32(v50) < base.Ui32(v88))
	goto L29
L31:
	;
	goto L32
L32:
	;
	v110 = int32(base.Ui32(v50-v88) >> (uint(int32(31)) % 32))
	goto L29
L33:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v68))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v88)) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v122 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v122 = base.B2i32(base.Ui32(v88) < base.Ui32(v68))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v122 = int32(base.Ui32(v88-v68) >> (uint(int32(31)) % 32))
	goto L34
L38:
	;
	v123 = v88
	goto L40
L39:
	;
	v123 = v68
	goto L40
L40:
	;
	v124 = v123
	goto L28
L41:
	;
	goto L42
L42:
	;
	if int32(base.Ui32(v51-v89)>>(uint(int32(31))%32)) != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	goto L44
L44:
	;
	if int32(base.Ui32(v89-v70)>>(uint(int32(31))%32)) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = v89
	goto L47
L46:
	;
	v133 = v70
	goto L47
L47:
	;
	v134 = v124
	v135 = v133
	goto L21
L48:
	;
	if v136 != 0 {
		v66 = v136
		v68 = v134
		v70 = v135
		goto L19
	} else {
		goto L49
	}
L49:
	;
	goto L20
L50:
	;
	F_relation_close(m, v55, int32(1))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v164 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v167 = v22 + int32(32)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	F_ScanKeyInit(m, v167, int32(1), int32(3), int32(184), v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_systable_inplace_update_begin(m, v164, int32(2672), v167, v22+int32(92), v22+int32(28))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v22)+92))
	if v182 == int32(0) {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v185 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+22)))
	v188 = v186 + v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+84))
	if v189 == v140 {
		v219 = v140
		v220 = v185
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v188)+88))
	if v142 != v221 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v140))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v189)) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+84)) = v140
	v219 = v140
	v220 = int32(1)
	goto L56
L59:
	;
	if v202 != 0 {
		goto L58
	} else {
		goto L63
	}
L60:
	;
	v202 = base.B2i32(base.Ui32(v189) < base.Ui32(v140))
	goto L59
L61:
	;
	goto L62
L62:
	;
	v202 = int32(base.Ui32(v189-v140) >> (uint(int32(31)) % 32))
	goto L59
L63:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188)+84))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v203))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v215 != 0 {
		goto L58
	} else {
		goto L68
	}
L65:
	;
	v215 = base.B2i32(base.Ui32(v50) < base.Ui32(v203))
	goto L64
L66:
	;
	goto L67
L67:
	;
	v215 = int32(base.Ui32(v50-v203) >> (uint(int32(31)) % 32))
	goto L64
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v188)+84))
	v219 = v216
	v220 = v185
	goto L56
L69:
	;
	goto L72
L70:
	;
	v231 = v142
	goto L71
L71:
	;
	if v220 != 0 {
		v258 = v231
		goto L9
	} else {
		goto L76
	}
L72:
	;
	if int32(base.Ui32(v221-v142)>>(uint(int32(31))%32)) != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v188)+88))
	goto L74
L74:
	;
	if int32(base.Ui32(v51-v226)>>(uint(int32(31))%32)) != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v188)+88))
	v231 = v230
	goto L71
L76:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_systable_inplace_update_cancel(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v264 = v231
	v265 = int32(0)
	goto L8
L78:
	;
	F_relation_close(m, v55, int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	goto L7
L80:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v247
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_0), v22)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_1), int32(1776), int32(_a_F_vac_update_datfrozenxid_2))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v264 = v258
	v265 = int32(1)
	goto L8
L84:
	;
	F_relation_close(m, v164, int32(3))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v265 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v274 = int32(1)
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v280 = F_LWLockAcquire(m, v276+int32(384), v274)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v328 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L100
	}
L89:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[2]))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+36))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v283)+20))
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v283)+8))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+16))
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v289+int32(384))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if base.B2i32(v285 == int32(0))|base.B2i32(base.Ui32(v287) < base.Ui32(int32(3))) != 0 {
		v320 = v274
		goto L91
	} else {
		goto L92
	}
L91:
	;
	if v320 == int32(0) {
		goto L7
	} else {
		goto L99
	}
L92:
	;
	v299 = base.I32_wrap_i64(v286)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v285))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v299)) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v311 != 0 {
		v320 = v274
		goto L91
	} else {
		goto L97
	}
L94:
	;
	v311 = base.B2i32(base.Ui32(v285) <= base.Ui32(v299))
	goto L93
L95:
	;
	goto L96
L96:
	;
	v311 = base.B2i32(int32(0) <= v299-v285)
	goto L93
L97:
	;
	v313 = int32(0)
	v316 = F_SearchSysCacheExists(m, int32(21), v284, v313, v313, v313)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v320 = v316 ^ int32(1)
	goto L91
L99:
	;
	goto L88
L100:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v335 = F_LWLockAcquire(m, v331+int32(_a_F_vac_update_datfrozenxid_3), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[0]))
	v341 = F_table_open(m, int32(1262), int32(1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+188))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	m.T0[v486].(func(*base.Module, int32))(m, v345)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L151
	}
L103:
	;
	v343 = int32(0)
	v345 = F_table_beginscan_catalog(m, v341, v343, v343)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v347 = F_heap_getnext(m, v345)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v347 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v467 = v219
	v469 = v264
	v470 = v1
	v471 = v338
	v472 = v338
	v473 = v1
	goto L102
L107:
	;
	goto L108
L108:
	;
	v351 = base.I32_wrap_i64(v328)
	v352 = v347
	v354 = v219
	v356 = v264
	v357 = v1
	v358 = v338
	v359 = v338
	v360 = v1
	goto L109
L109:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v352)+16))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+22)))
	v373 = v371 + v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+84))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v373)+88))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+80))
	goto L112
L110:
	;
	v467 = v457
	v469 = v458
	v470 = v459
	v471 = v460
	v472 = v461
	v473 = v462
	goto L102
L111:
	;
	v463 = F_heap_getnext(m, v345)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L149
	}
L112:
	;
	if v376 == int32(-2) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v381 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v374))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L116:
	;
	if v381 == int32(0) {
		v457 = v354
		v458 = v356
		v459 = v357
		v460 = v358
		v461 = v359
		v462 = v360
		goto L111
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v373 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_4), v22+int32(16))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_1), int32(1905), int32(_a_F_vac_update_datfrozenxid_5))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v457 = v354
	v458 = v356
	v459 = v357
	v460 = v358
	v461 = v359
	v462 = v360
	goto L111
L120:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v374))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v351)) == int32(0) {
		goto L132
	} else {
		goto L133
	}
L121:
	;
	if v409 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v409 = base.B2i32(base.Ui32(v50) < base.Ui32(v374))
	goto L121
L123:
	;
	goto L124
L124:
	;
	v409 = int32(base.Ui32(v50-v374) >> (uint(int32(31)) % 32))
	goto L121
L125:
	;
	goto L128
L126:
	;
	goto L127
L127:
	;
	v418 = int32(1)
	goto L120
L128:
	;
	if int32(base.Ui32(v51-v375)>>(uint(int32(31))%32)) == int32(0) {
		v418 = v357
		goto L120
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	goto L145
L131:
	;
	if v430 != 0 {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v430 = base.B2i32(base.Ui32(v351) < base.Ui32(v374))
	goto L131
L133:
	;
	goto L134
L134:
	;
	v430 = int32(base.Ui32(v351-v374) >> (uint(int32(31)) % 32))
	goto L131
L135:
	;
	v447 = v354
	v448 = v359
	v449 = int32(1)
	goto L130
L136:
	;
	goto L137
L137:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v354))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v374)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v443 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v443 = base.B2i32(base.Ui32(v374) < base.Ui32(v354))
	goto L138
L140:
	;
	goto L141
L141:
	;
	v443 = int32(base.Ui32(v374-v354) >> (uint(int32(31)) % 32))
	goto L138
L142:
	;
	v447 = v354
	v448 = v359
	v449 = v360
	goto L130
L143:
	;
	goto L144
L144:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v447 = v374
	v448 = v446
	v449 = v360
	goto L130
L145:
	;
	if int32(base.Ui32(v375-v356)>>(uint(int32(31))%32)) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v454 = v375
	v455 = v453
	goto L148
L147:
	;
	v454 = v356
	v455 = v358
	goto L148
L148:
	;
	v457 = v447
	v458 = v454
	v459 = v418
	v460 = v455
	v461 = v448
	v462 = v449
	goto L111
L149:
	;
	if v463 != 0 {
		v352 = v463
		v354 = v457
		v356 = v458
		v357 = v459
		v358 = v460
		v359 = v461
		v360 = v462
		goto L109
	} else {
		goto L150
	}
L150:
	;
	goto L110
L151:
	;
	F_relation_close(m, v341, int32(1))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v473 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v1181+int32(_a_F_vac_update_datfrozenxid_3))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L1
	} else {
		goto L314
	}
L154:
	;
	v494 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	if v470 != 0 {
		goto L153
	} else {
		goto L162
	}
L157:
	;
	if v494 == int32(0) {
		goto L153
	} else {
		goto L158
	}
L158:
	;
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_6), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errdetail(m, int32(_a_F_vac_update_datfrozenxid_7), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_1), int32(1951), int32(_a_F_vac_update_datfrozenxid_5))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	goto L153
L162:
	;
	v511 = int32(0)
	v513 = m.G0
	v515 = v513 - int32(16)
	m.G0 = v515
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v522 = F_LWLockAcquire(m, v518+int32(_a_F_vac_update_datfrozenxid_8), int32(1))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v529 = F_LWLockAcquire(m, v525+int32(3456), int32(1))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[3]))
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v532)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v532)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v535
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v532)+8))
	v538 = *(*int64)(unsafe.Add(mBase, uint32(v532)))
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v540+int32(3456))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	if base.B2i32(v535 == v537)&base.B2i32(v533 == v538) != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v707+int32(_a_F_vac_update_datfrozenxid_8))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L202
	}
L167:
	;
	__phi550 = v511
	__phi551 = int32(-1)
	__phi553 = v535
	__phi558 = v511
	__phi566 = int64(-1)
	__phi567 = v533
	v550 = __phi550
	v551 = __phi551
	v553 = __phi553
	v558 = __phi558
	v566 = __phi566
	v567 = __phi567
	goto L168
L168:
	;
	if v566 != v567 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v621 < int32(0) {
		goto L166
	} else {
		goto L197
	}
L170:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	if int32(0) <= v551 {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	v620 = v550
	v621 = v551
	v622 = v558
	goto L172
L172:
	;
	v623 = v620 + v553
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v623)+8))
	if base.Ui32(v624) < base.Ui32(int32(3)) {
		v648 = v622
		goto L182
	} else {
		goto L183
	}
L173:
	;
	if v558 != 0 {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	v595 = v558
	v596 = v571
	goto L175
L175:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)+28))
	v599 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v600 = base.I64_rem_s(v567, v599)
	v606 = F_LWLockAcquire(m, v597+base.I32_wrap_i64(v600)<<(uint(int32(7))%32), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L180
	}
L176:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v571)+12))
	v576 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v574+v551))) = uint8(v576)
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v580 = v579
	goto L178
L177:
	;
	v580 = v571
	goto L178
L178:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v580)+28))
	v584 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v585 = base.I64_rem_s(v566, v584)
	F_LWLockRelease(m, v582+base.I32_wrap_i64(v585)<<(uint(int32(7))%32))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v593 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v595 = int32(0)
	v596 = v593
	goto L175
L180:
	;
	v611 = F_SimpleLruReadPage(m, int32(_a_F_vac_update_datfrozenxid_9), v567, int32(1), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615+v611<<(uint(int32(2))%32))))
	v620 = v619
	v621 = v611
	v622 = v595
	goto L172
L182:
	;
	v649 = *(*int64)(unsafe.Add(mBase, uint32(v515)+8))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v623)))
	v652 = v650 + v651
	v656 = base.B2i32(base.Ui32(v652-int32(_a_F_vac_update_datfrozenxid_10)) < base.Ui32(int32(-8193)))
	v658 = v649 + base.I64_extend_i32_u(v656)
	*(*int64)(unsafe.Add(mBase, uint32(v515)+8)) = v658
	if base.Ui32(v652-int32(_a_F_vac_update_datfrozenxid_10)) < base.Ui32(int32(-8193)) {
		goto L193
	} else {
		goto L194
	}
L183:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v467))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v624)) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v638 == int32(0) {
		v648 = v622
		goto L182
	} else {
		goto L188
	}
L185:
	;
	v638 = base.B2i32(base.Ui32(v624) < base.Ui32(v467))
	goto L184
L186:
	;
	goto L187
L187:
	;
	v638 = int32(base.Ui32(v624-v467) >> (uint(int32(31)) % 32))
	goto L184
L188:
	;
	v643 = F_TransactionIdDidCommit(m, v624)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	if v643 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v645 = int32(2)
	goto L192
L191:
	;
	v645 = int32(0)
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v623)+8)) = v645
	v648 = int32(1)
	goto L182
L193:
	;
	v661 = int32(0)
	goto L195
L194:
	;
	v661 = v652
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515))) = v661
	if base.B2i32(v661 != v537)|base.B2i32(v658 != v538) != 0 {
		__phi550 = v620
		__phi551 = v621
		__phi553 = v661
		__phi558 = v648
		__phi566 = v567
		__phi567 = v658
		v550 = __phi550
		v551 = __phi551
		v553 = __phi553
		v558 = __phi558
		v566 = __phi566
		v567 = __phi567
		goto L168
	} else {
		goto L196
	}
L196:
	;
	goto L169
L197:
	;
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	if v648 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)+12))
	v672 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v670+v621))) = uint8(v672)
	v675 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[4]))
	v676 = v675
	goto L200
L199:
	;
	v676 = v669
	goto L200
L200:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+28))
	v679 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[5])))
	v680 = base.I64_rem_s(v567, v679)
	F_LWLockRelease(m, v677+base.I32_wrap_i64(v680)<<(uint(int32(7))%32))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	goto L166
L202:
	;
	m.G0 = v515 + int32(16)
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v720 = F_LWLockAcquire(m, v716+int32(_a_F_vac_update_datfrozenxid_11), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[2]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+40))
	if v724 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v745+int32(_a_F_vac_update_datfrozenxid_11))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L211
	}
L205:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v467))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v724)) == int32(0) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v738 == int32(0) {
		goto L204
	} else {
		goto L210
	}
L207:
	;
	v738 = base.B2i32(base.Ui32(v724) < base.Ui32(v467))
	goto L206
L208:
	;
	goto L209
L209:
	;
	v738 = int32(base.Ui32(v724-v467) >> (uint(int32(31)) % 32))
	goto L206
L210:
	;
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v742)+40)) = v467
	goto L204
L211:
	;
	v750 = m.G0
	v752 = v750 - int32(32)
	m.G0 = v752
	*(*int64)(unsafe.Add(mBase, uint32(v752)+8)) = base.I64_extend_i32_u(int32(base.Ui32(v467) >> (uint(int32(15)) % 32)))
	v762 = F_SlruScanDirectory(m, int32(_a_F_vac_update_datfrozenxid_12), int32(288), v752+int32(8))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	if v762 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	F_AdvanceOldestClogXid(m, v467)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v788 = int32(32)
	m.G0 = v752 + v788
	v791 = m.G0
	v793 = v791 - v788
	m.G0 = v793
	v796 = base.I32_div_u_s(v467, int32(819))
	*(*int64)(unsafe.Add(mBase, uint32(v793)+8)) = base.I64_extend_i32_u(v796)
	v803 = F_SlruScanDirectory(m, int32(_a_F_vac_update_datfrozenxid_13), int32(288), v793+int32(8))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L1
	} else {
		goto L222
	}
L216:
	;
	v766 = *(*int64)(unsafe.Add(mBase, uint32(v752)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v752)+28)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v752)+24)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v752)+16)) = v766
	F_XLogBeginInsert(m)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v772 = int32(16)
	F_XLogRegisterData(m, v752+v772, v772)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v779 = F_XLogInsert(m, int32(3), int32(16))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_XLogFlush(m, v779)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v784 = *(*int64)(unsafe.Add(mBase, uint32(v752)+8))
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_12), v784)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	goto L215
L222:
	;
	if v803 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v805 = *(*int64)(unsafe.Add(mBase, uint32(v793)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v793)+24)) = v467
	*(*int64)(unsafe.Add(mBase, uint32(v793)+16)) = v805
	F_XLogBeginInsert(m)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	m.G0 = v793 + int32(32)
	v827 = m.G0
	v829 = v827 - int32(144)
	m.G0 = v829
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v836 = F_LWLockAcquire(m, v832+int32(_a_F_vac_update_datfrozenxid_14), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L230
	}
L226:
	;
	F_XLogRegisterData(m, v793+int32(16), int32(12))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v817 = F_XLogInsert(m, int32(18), int32(16))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v820 = *(*int64)(unsafe.Add(mBase, uint32(v793)+8))
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_13), v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	goto L225
L230:
	;
	v839 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v843 = F_LWLockAcquire(m, v839+int32(1664), int32(1))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[6]))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v846)))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v846)+12))
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v851+int32(1664))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	if v469-v849 <= int32(0) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v1148 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v1148+int32(_a_F_vac_update_datfrozenxid_14))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L1
	} else {
		goto L311
	}
L234:
	;
	goto L233
L235:
	;
	goto L236
L236:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v829)+104)) = int64(-1)
	v865 = F_SlruScanDirectory(m, int32(_a_F_vac_update_datfrozenxid_15), int32(294), v829+int32(104))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v867 = int32(1)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v829)+104))
	v870 = v868 << (uint(int32(11)) % 32)
	if base.Ui32(v870) <= base.Ui32(v867) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v873 = v867
	goto L240
L239:
	;
	v873 = v870
	goto L240
L240:
	;
	if v849-v873 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	goto L233
L242:
	;
	goto L243
L243:
	;
	if v848 == v849 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v848 != v469 {
		goto L256
	} else {
		goto L257
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829)+116)) = v847
	goto L244
L246:
	;
	goto L247
L247:
	;
	v881 = F_find_multixact_start(m, v849, v829+int32(116))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	if v881 != 0 {
		goto L244
	} else {
		goto L249
	}
L249:
	;
	v885 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	if v885 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829)+100)) = v873
	*(*int32)(unsafe.Add(mBase, uint32(v829)+96)) = v849
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_16), v829+int32(96))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	goto L233
L254:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_17), int32(3259), int32(_a_F_vac_update_datfrozenxid_18))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v902 = F_find_multixact_start(m, v469, v829+int32(120))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	v922 = v847
	goto L258
L258:
	;
	if v922 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L259:
	;
	if v902 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v908 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v829)+120))
	v922 = v921
	goto L258
L263:
	;
	if v908 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829)+80)) = v469
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_19), v829+int32(80))
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	goto L233
L267:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_17), int32(3277), int32(_a_F_vac_update_datfrozenxid_18))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	goto L266
L269:
	;
	v927 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v940 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L1
	} else {
		goto L278
	}
L272:
	;
	if v927 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = v469
	F_errmsg(m, int32(_a_F_vac_update_datfrozenxid_20), v829)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	goto L233
L276:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_17), int32(3294), int32(_a_F_vac_update_datfrozenxid_18))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v829)+116))
	if v940 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	v993 = int32(_a_F_vac_update_datfrozenxid_21)
	v995 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7]))
	v996 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7])) = v995 + v996
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[8]))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v1000)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1000)+120)) = v1001 | v996
	*(*int32)(unsafe.Add(mBase, uint32(v829)+140)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v829)+136)) = v942
	*(*int32)(unsafe.Add(mBase, uint32(v829)+132)) = v469
	*(*int32)(unsafe.Add(mBase, uint32(v829)+128)) = v849
	*(*int32)(unsafe.Add(mBase, uint32(v829)+124)) = v471
	F_XLogBeginInsert(m)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L1
	} else {
		goto L285
	}
L280:
	;
	v945 = int32(1636)
	v946 = base.I32_div_u_s(v942, v945)
	v947 = int32(5)
	v950 = base.I32_div_u_s(v922, v945)
	v952 = int32(base.Ui32(v950) >> (uint(v947) % 32))
	v990 = int32(base.Ui32(v946) >> (uint(v947) % 32))
	v991 = v952
	v992 = base.I64_extend_i32_u(v952)
	goto L279
L281:
	;
	goto L282
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829)+60)) = v922
	*(*int32)(unsafe.Add(mBase, uint32(v829)+56)) = v942
	v956 = int32(1636)
	v957 = base.I32_div_u_s(v922, v956)
	v958 = int32(5)
	v959 = int32(base.Ui32(v957) >> (uint(v958) % 32))
	v960 = base.I64_extend_i32_u(v959)
	*(*int64)(unsafe.Add(mBase, uint32(v829)+72)) = v960
	v965 = base.I32_div_u_s(v942, v956)
	v967 = int32(base.Ui32(v965) >> (uint(v958) % 32))
	*(*int64)(unsafe.Add(mBase, uint32(v829-int32(-64)))) = base.I64_extend_i32_u(v967)
	*(*int32)(unsafe.Add(mBase, uint32(v829)+36)) = v469
	v971 = int32(16)
	*(*int64)(unsafe.Add(mBase, uint32(v829)+48)) = base.I64_extend_i32_u(int32(base.Ui32(v469) >> (uint(v971) % 32)))
	*(*int32)(unsafe.Add(mBase, uint32(v829)+32)) = v849
	*(*int64)(unsafe.Add(mBase, uint32(v829)+40)) = base.I64_extend_i32_u(int32(base.Ui32(v849) >> (uint(v971) % 32)))
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_22), v829+int32(32))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_17), int32(3307), int32(_a_F_vac_update_datfrozenxid_18))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	v990 = v967
	v991 = v959
	v992 = v960
	goto L279
L285:
	;
	F_XLogRegisterData(m, v829+int32(124), int32(20))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	v1019 = F_XLogInsert(m, int32(6), int32(48))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_XLogFlush(m, v1019)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	v1028 = F_LWLockAcquire(m, v1024+int32(1664), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v1031)+16)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v1031)+12)) = v469
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[1]))
	F_LWLockRelease(m, v1035+int32(1664))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	if v990 != v991 {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1058 = base.I64_extend_i32_u(v990)
	goto L294
L292:
	;
	goto L293
L293:
	;
	v1106 = int32(1)
	if v469 == v1106 {
		goto L307
	} else {
		goto L308
	}
L294:
	;
	v1063 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L1
	} else {
		goto L296
	}
L295:
	;
	goto L293
L296:
	;
	if v1063 != 0 {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v829)+16)) = v1058
	F_errmsg_internal(m, int32(_a_F_vac_update_datfrozenxid_23), v829+int32(16))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L300
	}
L298:
	;
	goto L299
L299:
	;
	F_SlruDeleteSegment(m, v1058)
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L302
	}
L300:
	;
	F_errfinish(m, int32(_a_F_vac_update_datfrozenxid_17), int32(3132), int32(_a_F_vac_update_datfrozenxid_24))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	goto L299
L302:
	;
	if v1058 != int64(82040) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1083 = v1058 + int64(1)
	goto L305
L304:
	;
	v1083 = int64(0)
	goto L305
L305:
	;
	if v1083 != v992 {
		v1058 = v1083
		goto L294
	} else {
		goto L306
	}
L306:
	;
	goto L295
L307:
	;
	v1112 = int32(_a_F_vac_update_datfrozenxid_25)
	goto L309
L308:
	;
	v1112 = int32(base.Ui32(v469-v1106) >> (uint(int32(11)) % 32))
	goto L309
L309:
	;
	F_SimpleLruTruncate(m, int32(_a_F_vac_update_datfrozenxid_15), base.I64_extend_i32_u(v1112))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[8]))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1117)+120)) = v1118 & int32(-2)
	v1122 = int32(_a_F_vac_update_datfrozenxid_21)
	v1124 = *(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_vac_update_datfrozenxid[7])) = v1124 - int32(1)
	goto L233
L311:
	;
	m.G0 = v829 + int32(144)
	F_SetTransactionIdLimit(m, v467, v472)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_SetMultiXactIdLimit(m, v469, v471, int32(0))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	goto L153
L314:
	;
	goto L7
}
func F_vac_update_relstats(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v47 float32
	_ = v47
	var v48 float32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = v19 - int32(-64)
	F_ScanKeyInit(m, v27, int32(1), int32(3), int32(184), v21)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_systable_inplace_update_begin(m, v24, int32(2662), v27, v19+int32(60), v19+int32(56))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if v40 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+96))
	if l1 != v44 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L83
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = l1
	goto L10
L9:
	;
	goto L10
L10:
	;
	v47 = base.F32_demote_f64(l2)
	v48 = *(*float32)(unsafe.Add(mBase, uint32(v43)+100))
	if base.F32_eq(v47, v48) != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = base.B2i32(l1 != v44)
	goto L13
L12:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v43)+100)) = v47
	v53 = int32(1)
	goto L13
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+104))
	if l3 != v54 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = l3
	v58 = int32(1)
	goto L16
L15:
	;
	v58 = v53
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+108))
	if l4 != v59 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+108)) = l4
	v63 = int32(1)
	goto L19
L18:
	;
	v63 = v58
	goto L19
L19:
	;
	if l10 != 0 {
		v88 = v63
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v43)+136))
	if l8 != 0 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	if l5 != 0 {
		v72 = v63
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+124)))
	if v73 != int32(1) {
		v80 = v72
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+116)))
	if v64&int32(1) == int32(0) {
		v72 = v63
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+116)) = uint8(v69)
	v72 = int32(1)
	goto L22
L25:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+125)))
	if v81 != int32(1) {
		v88 = v80
		goto L20
	} else {
		goto L28
	}
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v76 != 0 {
		v80 = v72
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+124)) = uint8(v77)
	v80 = int32(1)
	goto L25
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v84 != 0 {
		v88 = v80
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v85 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+125)) = uint8(v85)
	v88 = int32(1)
	goto L20
L30:
	;
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v90)
	goto L32
L31:
	;
	goto L32
L32:
	;
	v92 = int32(0)
	if base.B2i32(v89 == l6)|base.B2i32(base.Ui32(l6) < base.Ui32(int32(3))) != 0 {
		v136 = v88
		v137 = v92
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v43)+140))
	if l9 != 0 {
		goto L49
	} else {
		goto L50
	}
L34:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l6))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v89)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v108 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v108 = base.B2i32(base.Ui32(v89) < base.Ui32(l6))
	goto L35
L37:
	;
	goto L38
L38:
	;
	v108 = int32(base.Ui32(v89-l6) >> (uint(int32(31)) % 32))
	goto L35
L39:
	;
	v111 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+136)) = l6
	v129 = int32(1)
	v131 = v108 ^ v129
	if l8 == int32(0) {
		v136 = v129
		v137 = v131
		goto L33
	} else {
		goto L48
	}
L42:
	;
	v113 = base.I32_wrap_i64(v111)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v89))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v113)) == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v125 == int32(0) {
		v136 = v88
		v137 = v92
		goto L33
	} else {
		goto L47
	}
L44:
	;
	v125 = base.B2i32(base.Ui32(v113) < base.Ui32(v89))
	goto L43
L45:
	;
	goto L46
L46:
	;
	v125 = int32(base.Ui32(v113-v89) >> (uint(int32(31)) % 32))
	goto L43
L47:
	;
	goto L41
L48:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8))) = uint8(v134)
	v136 = v129
	v137 = v131
	goto L33
L49:
	;
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v140)
	goto L51
L50:
	;
	goto L51
L51:
	;
	if base.B2i32(l7 == int32(0))|base.B2i32(v139 == l7) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	F_relation_close(m, v24, int32(3))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L68
	}
L53:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	F_systable_inplace_update_cancel(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L67
	}
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	F_systable_inplace_update_finish(m, v170, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L66
	}
L55:
	;
	v166 = int32(0)
	if v136 == v166 {
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v148 = int32(base.Ui32(v139-l7) >> (uint(int32(31)) % 32))
	goto L57
L57:
	;
	if v148 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v151 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+140)) = l7
	v160 = v148 ^ int32(1)
	if l9 == int32(0) {
		v169 = v160
		goto L54
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if int32(base.Ui32(v151-v139)>>(uint(int32(31))%32)) == int32(0) {
		goto L55
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l9))) = uint8(v163)
	v169 = v160
	goto L54
L65:
	;
	v169 = v166
	goto L54
L66:
	;
	v177 = v169
	goto L52
L67:
	;
	v177 = v166
	goto L52
L68:
	;
	if v137 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v177 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	v185 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v185 == int32(0) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v192 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_3), v19+int32(32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1595), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L69
L76:
	;
	m.G0 = v19 + int32(112)
	return
L77:
	;
	v213 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v213 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v220 + int32(4)
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_0), v19+int32(16))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1601), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L76
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	F_errmsg_internal(m, int32(_a_F_vac_update_relstats_4), v19)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_vac_update_relstats_1), int32(1474), int32(_a_F_vac_update_relstats_2))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
