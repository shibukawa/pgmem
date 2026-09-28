package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddRelationNewConstraints(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int64
	_ = v777
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v825 int64
	_ = v825
	var v826 int64
	_ = v826
	var v827 int64
	_ = v827
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v843 int64
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int64
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int64
	_ = v869
	var v872 int64
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1124 int32
	_ = v1124
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1220 int32
	_ = v1220
	var v1225 int32
	_ = v1225
	v5 = l4
	v8 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(384)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
	v34 = v33
	goto L3
L2:
	;
	v34 = v8
	goto L3
L3:
	;
	v36 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = l6
	v41 = int32(1)
	v42 = int32(0)
	v45 = F_addRangeTableEntryForRelation(m, v36, l0, v41, v42, v42, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v47 = int32(1)
	F_addNSItemToQuery(m, v36, v45, v47, v47, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v53 = v5 ^ int32(1)
	if l1 == int32(0) {
		v169 = v8
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l2 == int32(0) {
		v1188 = v169
		v1191 = v34
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 <= int32(0) {
		v169 = v8
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v68 = v8
	v75 = v8
	goto L11
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v68<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
	v99 = v91 + v92<<(uint(int32(3))%32) + v96*int32(100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v99)+18)))
	v107 = F_cookDefault(m, v36, v90, v102, v103, v99-int32(68), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v169 = v148
	goto L8
L13:
	;
	v150 = v68 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v150 < v151 {
		v68 = v150
		v75 = v148
		goto L11
	} else {
		goto L27
	}
L14:
	;
	if v107 == int32(0) {
		v148 = v75
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+8)))
	if v111 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if l5 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	if v112 != int32(7) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+32)))
	if v115 != 0 {
		v148 = v75
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_AddRelationNewConstraints[0]))
	F_CheckUsageOnTypesInSingleRelExpr(m, v107, v118, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
	v124 = F_StoreAttrDefault(m, l0, v123, v107, l5)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v127 = F_palloc(m, int32(28))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v129 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+8)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = int32(2)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+26)) = uint8(v129)
	*(*uint16)(unsafe.Add(mBase, uint32(v127)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+22)) = uint8(v5)
	v139 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v127)+20)) = uint16(v139)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v107
	*(*uint16)(unsafe.Add(mBase, uint32(v127)+12)) = uint16(v134)
	v143 = F_lappend(m, v75, v127)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v148 = v143
	goto L13
L27:
	;
	goto L12
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L4
	} else {
		goto L290
	}
L29:
	;
	F_SetRelationNumChecks(m, l0, v1191)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L4
	} else {
		goto L289
	}
L30:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v181 <= int32(0) {
		v1188 = v169
		v1191 = v34
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v202 = v8
	v204 = v169
	v207 = v34
	v209 = v8
	v210 = v8
	goto L32
L32:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214+v209<<(uint(int32(2))%32))))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	switch v219 - int32(1) {
	case 0:
		goto L47
	default:
		v1156 = v202
		v1158 = v204
		v1161 = v207
		v1164 = v210
		goto L34
	case 4:
		goto L48
	}
L33:
	;
	v1188 = v1158
	v1191 = v1161
	goto L29
L34:
	;
	v1169 = v209 + int32(1)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1169 < v1170 {
		v202 = v1156
		v204 = v1158
		v207 = v1161
		v209 = v1169
		v210 = v1164
		goto L32
	} else {
		goto L288
	}
L35:
	;
	if v765&int32(1) != 0 {
		goto L279
	} else {
		goto L280
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v942)+103)) = uint8(v1119)
	goto L35
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L4
	} else {
		goto L275
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L4
	} else {
		goto L271
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L4
	} else {
		goto L267
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L4
	} else {
		goto L263
	}
L41:
	;
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+14)))
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+16)))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	v1007 = F_StoreRelCheck(m, l0, v987, v246, v1004, v1005, v5, v53, v1006, l5)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L4
	} else {
		goto L260
	}
L42:
	;
	v763 = F_lappend(m, v202, v249)
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L4
	} else {
		goto L182
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L178
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L4
	} else {
		goto L174
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L4
	} else {
		goto L170
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L166
	}
L47:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	v371 = F_get_attnum(m, v366, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L94
	}
L48:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	if v222 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	if v249 != 0 {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v225 = F_transformExpr(m, v36, v222, int32(28))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v218)+24))
	v244 = F_stringToNode(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L60
	}
L53:
	;
	v228 = F_coerce_to_boolean(m, v36, v225, int32(_a_F_AddRelationNewConstraints_0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_assign_expr_collations(m, v36, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v232 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v235 != int32(1) {
		goto L46
	} else {
		goto L57
	}
L57:
	;
	if l5 != 0 {
		v246 = v228
		goto L49
	} else {
		goto L58
	}
L58:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_AddRelationNewConstraints[0]))
	F_CheckUsageOnTypesInSingleRelExpr(m, v228, v238, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v246 = v228
	goto L49
L60:
	;
	v246 = v244
	goto L49
L61:
	;
	if v202 == int32(0) {
		goto L42
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v338 = int32(0)
	v340 = F_pull_var_clause(m, v246, v338)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L87
	}
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
	if v252 <= int32(0) {
		goto L42
	} else {
		goto L65
	}
L65:
	;
	v255 = int32(0)
	if v255 < v252 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v258 = v252
	goto L68
L67:
	;
	v258 = v255
	goto L68
L68:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v267 = int32(0)
	goto L69
L69:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v259+v267<<(uint(int32(2))%32))))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if base.B2i32(v293 == int32(0))|base.B2i32(v293 != v296) != 0 {
		v314 = v293
		v315 = v296
		goto L72
	} else {
		goto L73
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L82
	}
L71:
	;
	if v314-v315 != 0 {
		goto L78
	} else {
		goto L79
	}
L72:
	;
	goto L71
L73:
	;
	v299 = v290
	v300 = v249
	goto L74
L74:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+1)))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+1)))
	if v304 == int32(0) {
		v314 = v304
		v315 = v303
		goto L72
	} else {
		goto L76
	}
L75:
	;
	v314 = v304
	v315 = v303
	goto L72
L76:
	;
	v307 = int32(1)
	if v304 == v303 {
		v299 = v299 + v307
		v300 = v300 + v307
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v318 = v267 + int32(1)
	if v258 != v318 {
		v267 = v318
		goto L69
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	goto L70
L81:
	;
	goto L42
L82:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_1))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v249
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_2), v29+int32(128))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2562), int32(_a_F_AddRelationNewConstraints_4))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357)+68))
	v362 = F_ChooseConstraintName(m, v357+int32(4), v356, int32(_a_F_AddRelationNewConstraints_5), v361, v202)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L4
	} else {
		goto L92
	}
L87:
	;
	v342 = F_list_union(m, v340)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v342 == int32(0) {
		v356 = v338
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v346 != int32(1) {
		v356 = v338
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v342)+12))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)))
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v351)+8)))
	v354 = F_get_attname(m, v349, v352, int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v356 = v354
	goto L86
L92:
	;
	v364 = F_lappend(m, v202, v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v987 = v362
	v992 = v364
	goto L41
L94:
	;
	if v371 == int32(0) {
		goto L45
	} else {
		goto L95
	}
L95:
	;
	if v371 < int32(0) {
		goto L44
	} else {
		goto L96
	}
L96:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+15)))
	v381 = m.G0
	v383 = v381 - int32(96)
	m.G0 = v383
	v385 = F_findNotNullConstraintAttnum(m, v377, v371)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	if v385 != int32(0) {
		v1156 = v202
		v1158 = v204
		v1161 = v207
		v1164 = v210
		goto L34
	} else {
		goto L154
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L147
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L141
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L135
	}
L101:
	;
	if v385 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v389 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v383 + int32(96)
	goto L97
L105:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v385)+16))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+22)))
	v393 = v391 + v392
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+106)))
	if v379 != v394 {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	if v380 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+76)))
	if v398 == int32(0) {
		goto L99
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v401 = int32(0)
	if base.B2i32(v5 == v401)|base.B2i32(v378 == v401) == v401 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v409 = v393 + int32(4)
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if base.B2i32(v412 == int32(0))|base.B2i32(v412 != v415) != 0 {
		v433 = v412
		v434 = v415
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	if v5 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L114:
	;
	if v433-v434 != 0 {
		goto L98
	} else {
		goto L121
	}
L115:
	;
	goto L114
L116:
	;
	v418 = v378
	v419 = v409
	goto L117
L117:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v419)+1)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v418)+1)))
	if v423 == int32(0) {
		v433 = v423
		v434 = v422
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v433 = v423
	v434 = v422
	goto L115
L119:
	;
	v426 = int32(1)
	if v423 == v422 {
		v418 = v418 + v426
		v419 = v419 + v426
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L113
L122:
	;
	F_relation_close(m, v389, int32(3))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L134
	}
L123:
	;
	F_CatalogTupleUpdate(m, v389, v385+int32(4), v385)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L133
	}
L124:
	;
	v439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393)+104)))
	v441 = v439 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v393)+104)) = uint16(v441)
	if base.I32_extend16_s(v441) == v441 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+103)))
	if v461 != 0 {
		goto L122
	} else {
		goto L132
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_6), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_7), int32(806), int32(_a_F_AddRelationNewConstraints_8))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v462 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+103)) = uint8(v462)
	goto L123
L133:
	;
	goto L122
L134:
	;
	goto L104
L135:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v488 = F_get_rel_name(m, v377)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+84)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v383)+80)) = v393 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_9), v383+int32(80))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+64)) = int32(_a_F_AddRelationNewConstraints_10)
	F_errhint(m, int32(_a_F_AddRelationNewConstraints_11), v383-int32(-64))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_7), int32(770), int32(_a_F_AddRelationNewConstraints_8))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v518 = F_get_rel_name(m, v377)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+52)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v383)+48)) = v393 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_12), v383+int32(48))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+32)) = int32(_a_F_AddRelationNewConstraints_13)
	F_errhint(m, int32(_a_F_AddRelationNewConstraints_14), v383+int32(32))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_7), int32(782), int32(_a_F_AddRelationNewConstraints_8))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v549 = F_get_attname(m, v377, v371, int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v551 = F_get_rel_name(m, v377)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383)+24)) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v383)+16)) = v378
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_15), v383+int32(16))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v383))) = v409
	v563 = F_errdetail(m, int32(_a_F_AddRelationNewConstraints_16), v383)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_7), int32(798), int32(_a_F_AddRelationNewConstraints_8))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	if v570 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v589 = F_lappend(m, v210, v588)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L162
	}
L156:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v573 = F_ConstraintNameIsUsed(m, int32(0), v572, v570)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+12))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v581)+4))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v576)+68))
	v585 = F_ChooseConstraintName(m, v576+int32(4), v582, int32(_a_F_AddRelationNewConstraints_17), v584, v210)
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L4
	} else {
		goto L161
	}
L159:
	;
	if v573 != 0 {
		goto L43
	} else {
		goto L160
	}
L160:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	v588 = v575
	goto L155
L161:
	;
	v588 = v585
	goto L155
L162:
	;
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+208)) = uint16(v371)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+68))
	v597 = int32(0)
	v599 = int32(1)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v614 = int32(32)
	v624 = F_CreateConstraintEntry(m, v588, v595, int32(110), v597, v597, v599, v592, v597, v601, v29+int32(208), v599, v599, v597, v597, v597, v597, v597, v597, v597, v597, v614, v614, v597, v597, v614, v597, v597, v597, v5, v53, v591, v597, v597)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v627 = F_palloc(m, int32(28))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v629 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v627)+20)) = uint8(v629)
	*(*int32)(unsafe.Add(mBase, uint32(v627)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v627)+12)) = uint16(v371)
	*(*int32)(unsafe.Add(mBase, uint32(v627)+8)) = v588
	*(*int32)(unsafe.Add(mBase, uint32(v627)+4)) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v627))) = v629
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v627)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v627)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v627)+21)) = uint8(v638)
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v627)+26)) = uint8(v642)
	v644 = F_lappend(m, v204, v627)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v1156 = v202
	v1158 = v644
	v1161 = v207
	v1164 = v589
	goto L34
L166:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_18))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v223 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_19), v29+int32(144))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(3469), int32(_a_F_AddRelationNewConstraints_20))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L4
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v673)+12))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)+4))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v677 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v676
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_21), v29+int32(160))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2658), int32(_a_F_AddRelationNewConstraints_4))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L174:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)+12))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v702
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_22), v29+int32(176))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2663), int32(_a_F_AddRelationNewConstraints_4))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L178:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_1))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v29)+196)) = v721 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_23), v29+int32(192))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2691), int32(_a_F_AddRelationNewConstraints_4))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L182:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+16)))
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+14)))
	v770 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v773 = v29 + int32(208)
	v777 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v773, int32(9), int32(3), int32(184), v777)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	F_ScanKeyInit(m, v29+int32(264), int32(10), int32(3), int32(184), int64(0))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	F_ScanKeyInit(m, v29+int32(320), int32(2), int32(3), int32(62), base.I64_extend_i32_u(v249))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v796 = F_systable_beginscan(m, v770, int32(2665), int32(1), int32(0), int32(3), v773)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v798 = F_systable_getnext(m, v796)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	if v798 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v798)+16))
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+22)))
	v803 = v801 + v802
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+72)))
	if v804 == int32(99) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	goto L191
L191:
	;
	F_systable_endscan(m, v796)
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L258
	}
L192:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v770)+52))
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+20)))
	if v808&int32(1) == int32(0) {
		goto L197
	} else {
		goto L198
	}
L193:
	;
	v881 = int32(0)
	goto L194
L194:
	;
	if v5 == int32(0) {
		v890 = l3
		goto L225
	} else {
		goto L226
	}
L195:
	;
	v874 = F_text_to_cstring(m, base.I32_wrap_i64(v872))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L4
	} else {
		goto L222
	}
L196:
	;
	v869 = int64(*(*int8)(unsafe.Add(mBase, uint32(v816))))
	v872 = v869
	goto L195
L197:
	;
	v813 = int32(*(*int16)(unsafe.Add(mBase, uint32(v807)+244)))
	if int32(0) <= v813 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+26)))
	if v845&int32(8) != 0 {
		goto L215
	} else {
		goto L216
	}
L200:
	;
	v816 = v813 + v803
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807)+248)))
	if v817 == int32(1) {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	goto L202
L202:
	;
	v843 = F_nocachegetattr(m, v798, int32(28), v807)
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L4
	} else {
		goto L214
	}
L203:
	;
	v820 = int32(*(*int16)(unsafe.Add(mBase, uint32(v807)+246)))
	if base.I32_popcnt(v820) != int32(1) {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L205
L205:
	;
	v872 = base.I64_extend_i32_u(v816)
	goto L195
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L211
	}
L207:
	;
	switch base.I32_ctz(v820) {
	case 0:
		goto L196
	case 1:
		goto L210
	case 2:
		goto L209
	case 3:
		goto L208
	default:
		goto L206
	}
L208:
	;
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v816)))
	v872 = v827
	goto L195
L209:
	;
	v826 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v816))))
	v872 = v826
	goto L195
L210:
	;
	v825 = int64(*(*int16)(unsafe.Add(mBase, uint32(v816))))
	v872 = v825
	goto L195
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v820
	F_errmsg_internal(m, int32(_a_F_AddRelationNewConstraints_24), v29)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_25), int32(123), int32(_a_F_AddRelationNewConstraints_26))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	v872 = v843
	goto L195
L215:
	;
	v849 = F_nocachegetattr(m, v798, int32(28), v807)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L219
	}
L218:
	;
	v872 = v849
	goto L195
L219:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v855 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AddRelationNewConstraints_27), v29+int32(112))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2797), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	v876 = F_stringToNode(m, v874)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	v878 = F_equal(m, v246, v876)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	v881 = v878
	goto L194
L225:
	;
	if v890&v881 == int32(0) {
		goto L40
	} else {
		goto L228
	}
L226:
	;
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+103)))
	if v884 != 0 {
		v890 = l3
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+131)))
	v890 = l3 | (v886 ^ int32(1))
	goto L225
L228:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+106)))
	if v894 == int32(1) {
		goto L39
	} else {
		goto L229
	}
L229:
	;
	v899 = int32(*(*int16)(unsafe.Add(mBase, uint32(v803)+104)))
	if v765&int32(1)&base.B2i32(int32(0) < v899) != 0 {
		goto L38
	} else {
		goto L230
	}
L230:
	;
	if v766&int32(1) == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	if v53&v767 == int32(1) {
		goto L236
	} else {
		goto L237
	}
L232:
	;
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+75)))
	if v907 != int32(1) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+76)))
	if v910 == int32(0) {
		goto L37
	} else {
		goto L234
	}
L234:
	;
	goto L231
L235:
	;
	v925 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L242
	}
L236:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+75)))
	if v916 != 0 {
		goto L235
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	if (v53|v767)&int32(1) != 0 {
		goto L235
	} else {
		goto L240
	}
L239:
	;
	goto L28
L240:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+75)))
	if v920 == int32(1) {
		goto L28
	} else {
		goto L241
	}
L241:
	;
	goto L235
L242:
	;
	if v925 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v249
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_29), v29+int32(48))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	v938 = F_heap_copytuple(m, v798)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2862), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v940)+22)))
	v942 = v940 + v941
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+131)))
	if v944 == int32(1) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v947 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v942)+104)) = uint16(v947)
	v1119 = int32(0)
	goto L36
L250:
	;
	goto L251
L251:
	;
	if v5 != 0 {
		v1119 = int32(1)
		goto L36
	} else {
		goto L252
	}
L252:
	;
	v951 = int32(*(*int16)(unsafe.Add(mBase, uint32(v942)+104)))
	v953 = v951 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v942)+104)) = uint16(v953)
	if base.I32_extend16_s(v953) == v953 {
		goto L35
	} else {
		goto L253
	}
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_6), int32(0))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2885), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L258:
	;
	F_relation_close(m, v770, int32(3))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L4
	} else {
		goto L259
	}
L259:
	;
	v987 = v249
	v992 = v763
	goto L41
L260:
	;
	v1010 = F_palloc(m, int32(28))
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L4
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+16)) = v246
	v1013 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1010)+12)) = uint16(v1013)
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+8)) = v987
	*(*int32)(unsafe.Add(mBase, uint32(v1010)+4)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = int32(5)
	v1019 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1010)+20)) = uint8(v1019)
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1010)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v1010)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v1010)+21)) = uint8(v1021)
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1010)+26)) = uint8(v1025)
	v1029 = F_lappend(m, v204, v1010)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	v1156 = v992
	v1158 = v1029
	v1161 = v207 + int32(1)
	v1164 = v210
	goto L34
L263:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_1))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v1038 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_23), v29+int32(96))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2817), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1060 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_30), v29+int32(16))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2824), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L4
	} else {
		goto L272
	}
L272:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1082 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_31), v29+int32(32))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2835), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v1104 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_32), v29+int32(80))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2845), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v1124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v942)+106)) = uint8(v1124)
	goto L281
L280:
	;
	goto L281
L281:
	;
	if v767&int32(1) == int32(0) {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	F_CatalogTupleUpdate(m, v770, v938+int32(4), v938)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L4
	} else {
		goto L285
	}
L283:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v942)+75)))
	if v1130 != 0 {
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1131 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v942)+75)) = uint16(v1131)
	goto L282
L285:
	;
	F_systable_endscan(m, v796)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	F_relation_close(m, v770, int32(3))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1156 = v763
	v1158 = v204
	v1161 = v207
	v1164 = v210
	goto L34
L288:
	;
	goto L33
L289:
	;
	m.G0 = v29 + int32(384)
	return v1188
L290:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L291
	}
L291:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1211 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_33), v29-int32(-64))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L4
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2857), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1225 = m.ExcPending
	if v1225 != 0 {
		goto L4
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DeleteRelationTuple(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v15 = F_SearchSysCache1(m, int32(57), base.I64_extend_i32_u(l0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_DeleteRelationTuple_0), v7)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_DeleteRelationTuple_1), int32(1604), int32(_a_F_DeleteRelationTuple_2))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				F_simple_heap_delete(m, v11, v15+int32(4))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_relation_close(m, v11, int32(3))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_FlushRelationBuffers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v95 int64
	_ = v95
	var v104 int64
	_ = v104
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v202 int64
	_ = v202
	var v212 int64
	_ = v212
	var v218 int64
	_ = v218
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v14 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v20
	v24 = F_smgropen(m, v12+int32(8), v17)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v14
	goto L3
L3:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+118)))
	if v44 != int32(116) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v24
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = v40
	goto L3
L7:
	;
	v36 = v28
	goto L9
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	v36 = v34
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+72)) = v36 + int32(1)
	goto L6
L10:
	;
	m.G0 = v12 + int32(48)
	return
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[0]))
	if v48 <= int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[1]))
	if v303 <= int32(0) {
		goto L10
	} else {
		goto L64
	}
L14:
	;
	v54 = v2
	goto L15
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[2]))
	v64 = v61 + v54*int32(56)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v65 != v66 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v298 = v54 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[0]))
	if v298 < v300 {
		v54 = v298
		goto L15
	} else {
		goto L63
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 != v69 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v71 != v72 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	F_ResourceOwnerEnlarge(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v80 = int64(4194304)
	v82 = base.AtomicRmwOr64(m, v64, int32(24), v80)
	if v82&v80 != int64(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v95 = v82
	goto L26
L24:
	;
	v178 = v82
	goto L25
L25:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v179 != v180 {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(_a_F_FlushRelationBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(_a_F_FlushRelationBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(_a_F_FlushRelationBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
	v104 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v104
	if v95&int64(4194304) != v104 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v178 = v165
	goto L25
L28:
	;
	goto L31
L29:
	;
	goto L30
L30:
	;
	v143 = int32(_a_F_FlushRelationBuffers_3)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[4]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(24))+8))
	if v146 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v12+int32(24))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	goto L30
L33:
	;
	v123 = int64(0)
	v126 = base.AtomicRmwCmpxchg64(m, v64, int32(24), v123, v123)
	if v126&int64(4194304) != v123 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v163 = int64(4194304)
	v165 = base.AtomicRmwOr64(m, v64, int32(24), v163)
	if v165&v163 != int64(0) {
		v95 = v165
		goto L26
	} else {
		goto L46
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[4])) = v161
	goto L36
L38:
	;
	if int32(999) < v144 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v144 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v151 = int32(900)
	if v151 <= v144 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v154 = v151
	goto L44
L43:
	;
	v154 = v144
	goto L44
L44:
	;
	v161 = v154 + int32(100)
	goto L37
L45:
	;
	v161 = v144 - int32(1)
	goto L37
L46:
	;
	goto L27
L47:
	;
	v287 = base.AtomicRmwSub64(m, v64, int32(24), int64(4194304))
	goto L17
L48:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = int64(25165824)
	if base.B2i32(v182 != v183)|base.B2i32(v178&v185 != v185) != 0 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v190 != v191 {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v193 = int64(0)
	v195 = int32(24)
	v196 = base.AtomicRmwCmpxchg64(m, v64, v195, v193, v193)
	v202 = base.AtomicRmwCmpxchg64(m, v64, v195, v196, v196&int64(-4194305)+int64(1))
	if v196 != v202 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v212 = v202
	goto L54
L52:
	;
	goto L53
L53:
	;
	v229 = int32(_a_F_FlushRelationBuffers_4)
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[5]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v236 = int32(1)
	v237 = v235 + v236
	*(*int32)(unsafe.Add(mBase, uint32(v230<<(uint(int32(2))%32))+uint32(_c_F_FlushRelationBuffers[6]))) = v237
	v240 = v230 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+uint32(_c_F_FlushRelationBuffers[7]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+uint32(_c_F_FlushRelationBuffers[8]))) = v237
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[5])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[9])) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v240)+uint32(_c_F_FlushRelationBuffers[10]))) = v236
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	F_ResourceOwnerRemember(m, v258, base.I64_extend_i32_s(v237), int32(_a_F_FlushRelationBuffers_5))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L4
	} else {
		goto L57
	}
L54:
	;
	v218 = base.AtomicRmwCmpxchg64(m, v64, int32(24), v212, v212&int64(-4194305)+int64(1))
	if v212 != v218 {
		v212 = v218
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L53
L56:
	;
	goto L55
L57:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	v265 = v263 + int32(1)
	F_BufferLockAcquire(m, v265, v64, int32(2))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_FlushBuffer(m, v64, v42, int32(3))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_BufferLockUnlock(m, v265, v64)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v64)+20))
	F_ResourceOwnerForget(m, v275, base.I64_extend_i32_s(v276+int32(1)), int32(_a_F_FlushRelationBuffers_5))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_UnpinBufferNoOwner(m, v64)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	goto L17
L63:
	;
	goto L16
L64:
	;
	v309 = v2
	goto L65
L65:
	;
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[11]))
	v319 = v316 + v309*int32(56)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v320 != v321 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L10
L67:
	;
	v368 = v309 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[1]))
	if v368 < v370 {
		v309 = v368
		goto L65
	} else {
		goto L77
	}
L68:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v323 != v324 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v326 != v327 {
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v329 = int64(0)
	v332 = base.AtomicRmwCmpxchg64(m, v319, int32(24), v329, v329)
	v333 = int64(25165824)
	if v332&v333 != v333 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = int32(1166)
	v340 = int32(_a_F_FlushRelationBuffers_6)
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[12])) = v12 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v341
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	F_ResourceOwnerEnlarge(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	F_PinLocalBuffer(m, v319, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_FlushLocalBuffer(m, v319, v42)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v319)+20))
	F_UnpinLocalBuffer(m, v358+int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[12])) = v364
	goto L67
L77:
	;
	goto L66
}
func F_GetRelationExcludedPublications(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	v10 = int64(0)
	v12 = F_SearchSysCacheList(m, int32(53), int32(1), base.I64_extend_i32_u(l0), v10, v10)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v16 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = int32(0)
	v29 = v16
	v30 = int32(0)
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)+v26<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v41 = F_lappend_oid(m, v30, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v44 = v29
	v45 = v30
	goto L11
L11:
	;
	v47 = v26 + int32(1)
	if v47 < v44 {
		v26 = v47
		v29 = v44
		v30 = v45
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v44 = v43
	v45 = v41
	goto L11
L13:
	;
	goto L8
L14:
	;
	return v45
}
func F_GetRelationIncludedPublications(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v10 = int64(0)
	v12 = F_SearchSysCacheList(m, int32(53), int32(1), base.I64_extend_i32_u(l0), v10, v10)
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v16 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = int32(0)
	v29 = v16
	v30 = int32(0)
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)+v26<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = v36 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)))
	if v39 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_ReleaseCatCacheList(m, v12)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v43 = F_lappend_oid(m, v30, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v46 = v29
	v47 = v30
	goto L11
L11:
	;
	v49 = v26 + int32(1)
	if v49 < v46 {
		v26 = v49
		v29 = v46
		v30 = v47
		goto L7
	} else {
		goto L13
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v46 = v45
	v47 = v43
	goto L11
L13:
	;
	goto L8
L14:
	;
	return v47
}
func F_LockRelationForExtension(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v16 = F_LockAcquire(m, v6, l1, v3, v3)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_LockRelationIdForSession(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v16 = F_LockAcquire(m, v6, l1, int32(1), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_NewRelationCreateToastTable(m *base.Module, l0 int32, l1 int64) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v4 = F_table_open(m, l0, int32(8))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = int32(0)
		v11 = F_create_toast_table(m, v4, v6, v6, l1, int32(8), v6, v6)
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_relation_close(m, v4, int32(0))
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_RelationDecrementReferenceCount(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2 - int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDecrementReferenceCount[0]))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDecrementReferenceCount[1]))
		F_ResourceOwnerForget(m, v9, base.I64_extend_i32_u(l0), int32(_a_F_RelationDecrementReferenceCount_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_RelationGetIndexAttOptions(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+120)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v198
L2:
	;
	if l1 == int32(0) {
		v198 = v14
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v56 = F_palloc0_mul(m, int32(4), v13)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L18
	}
L5:
	;
	v18 = F_palloc_mul(m, int32(4), v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v13 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v18
L9:
	;
	goto L10
L10:
	;
	v27 = v3
	goto L11
L11:
	;
	v38 = v27 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14+v38)))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v18
L13:
	;
	v44 = F_datumCopy(m, base.I64_extend_i32_u(v40), int32(0), int32(-1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	v47 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+v38))) = v47
	v51 = v27 + int32(1)
	if v51 != v13 {
		v27 = v51
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v47 = base.I32_wrap_i64(v44)
	goto L15
L17:
	;
	goto L12
L18:
	;
	v58 = int32(0)
	v59 = base.B2i32(v13 <= v58)
	if v59 == v58 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = v3
	goto L22
L20:
	;
	goto L21
L21:
	;
	v114 = int32(_a_F_RelationGetIndexAttOptions_0)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v117
	v120 = F_palloc_mul(m, int32(4), v13)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L31
	}
L22:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[1])))
	if base.B2i32(v76 != int32(1))|base.B2i32(v54 == int32(2659)) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v100 = v66 + int32(1)
	if v100 != v13 {
		v66 = v100
		goto L22
	} else {
		goto L30
	}
L25:
	;
	v85 = base.I32_extend16_s(v66 + int32(1))
	v86 = F_get_attoptions(m, v54, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v89 = F_index_opclass_options(m, l0, v85, v86, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56+v66<<(uint(int32(2))%32)))) = v89
	if v86 == int64(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, base.I32_wrap_i64(v86))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	goto L24
L30:
	;
	goto L23
L31:
	;
	if v59 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_pfree(m, v56)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L52
	}
L33:
	;
	v126 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v115
	if l1 != 0 {
		v198 = v56
		goto L1
	} else {
		goto L51
	}
L36:
	;
	v137 = v126 << (uint(int32(2)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v56+v137)))
	if v139 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v120
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v115
	if l1 != 0 {
		v198 = v56
		goto L1
	} else {
		goto L43
	}
L38:
	;
	v143 = F_datumCopy(m, base.I64_extend_i32_u(v139), int32(0), int32(-1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	v146 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120+v137))) = v146
	v150 = v126 + int32(1)
	if v150 != v13 {
		v126 = v150
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v146 = base.I32_wrap_i64(v143)
	goto L40
L42:
	;
	goto L37
L43:
	;
	v158 = int32(0)
	goto L44
L44:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56+v158<<(uint(int32(2))%32))))
	if v170 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L32
L46:
	;
	F_pfree(m, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v174 = v158 + int32(1)
	if v174 != v13 {
		v158 = v174
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L45
L51:
	;
	goto L32
L52:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v198 = v192
	goto L1
}
func F_RelationGetNumberOfBlocksInFork(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)))
	switch v10 - int32(83) {
	case 0, 22:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v24 != 0 {
			v48 = v24
			v49 = F_smgrnblocks(m, v48, l1)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v51 = v49
				m.G0 = v7 + int32(16)
				return v51
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v26
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v28
			v30 = F_smgropen(m, v7, v25)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v30
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
				if v34 != 0 {
					v42 = v34
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v36
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
					v42 = v40
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v42 + int32(1)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v48 = v46
				v49 = F_smgrnblocks(m, v48, l1)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = v49
					m.G0 = v7 + int32(16)
					return v51
				}
			}
		}
	default:
		v51 = int32(0)
		m.G0 = v7 + int32(16)
		return v51
	case 26, 31, 33:
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+148))
		v15 = m.T0[v14].(func(*base.Module, int32, int32) int64)(m, l0, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v51 = base.I32_wrap_i64(int64(base.Ui64(v15+int64(8191)) >> (uint(int64(13)) % 64)))
			m.G0 = v7 + int32(16)
			return v51
		}
	}
}
func F_RelationGetPartitionKey(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
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
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	v2 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	if v21 != int32(112) {
		v481 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v481
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v24 != 0 {
		v481 = v24
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = m.G0
	v27 = v25 - int32(80)
	m.G0 = v27
	v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v31 = F_SearchSysCache1(m, int32(45), v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v481 = v479
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L9
	} else {
		goto L102
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L9
	} else {
		goto L99
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L9
	} else {
		goto L96
	}
L9:
	;
	return int32(0)
L10:
	;
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[0]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_RelationGetPartitionKey_0), int32(0), int32(1024), int32(_a_F_RelationGetPartitionKey_1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L9
	} else {
		goto L93
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v46 = F_MemoryContextStrdup(m, v41, v43+int32(4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+36)) = v46
	v50 = F_MemoryContextAllocZero(m, v41, int32(56))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v50))) = base.I32_extend8_s(v55)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v58)
	v61 = v55 - int32(104)
	if base.B2i32(base.Ui32(int32(10)) < base.Ui32(v61))|base.B2i32(int32(1)<<(uint(v61)%32)&int32(1041) == int32(0)) != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v73 = F_SysCacheGetAttrNotNull(m, int32(45), v31, int32(6))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v77 = F_SysCacheGetAttrNotNull(m, int32(45), v31, int32(7))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v83 = F_SysCacheGetAttr(m, int32(45), v31, int32(8), v27+int32(79))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+79)))
	if v85 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v41
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v115 = F_palloc0_mul(m, int32(2), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L31
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1]))
	v110 = v89
	goto L21
L23:
	;
	goto L24
L24:
	;
	v91 = F_text_to_cstring(m, base.I32_wrap_i64(v83))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v93 = F_stringToNode(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v91)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v98 = F_eval_const_expressions(m, int32(0), v93)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_fix_opfuncids(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v102 = int32(_a_F_RelationGetPartitionKey_2)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v41
	v106 = F_copyObjectImpl(m, v98)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+12)) = v106
	v110 = v103
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = v115
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v120 = F_palloc0_mul(m, int32(4), v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+16)) = v120
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v125 = F_palloc0_mul(m, int32(4), v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+20)) = v125
	v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v130 = F_palloc0_mul(m, int32(28), v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v130
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v135 = F_palloc0_mul(m, int32(4), v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+28)) = v135
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v140 = F_palloc0_mul(m, int32(4), v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+32)) = v140
	v144 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v145 = F_palloc0_mul(m, int32(4), v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+36)) = v145
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v150 = F_palloc0_mul(m, int32(2), v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+40)) = v150
	v154 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v155 = F_palloc0_mul(m, int32(1), v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+44)) = v155
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v160 = F_palloc0_mul(m, int32(1), v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+48)) = v160
	v164 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v165 = F_palloc0_mul(m, int32(4), v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v50)+52)) = v165
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v110
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	v173 = v171 << (uint(int32(1)) % 32)
	if v173 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	base.MemoryCopy(m, v174, v54+int32(36), v173)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	if v178 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v180 = v179
	goto L47
L46:
	;
	v180 = v2
	goto L47
L47:
	;
	v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	if int32(0) < v181 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v170 == int32(104) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_ReleaseCatCache(m, v31)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L9
	} else {
		goto L75
	}
L51:
	;
	v188 = int32(2)
	goto L53
L52:
	;
	v188 = int32(1)
	goto L53
L53:
	;
	v190 = int32(24)
	v200 = int32(0)
	v202 = v180
	goto L54
L54:
	;
	v216 = v200 << (uint(int32(1)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(v216+v217))))
	v222 = v200 << (uint(int32(2)) % 32)
	v223 = base.I32_wrap_i64(v73) + v190 + v222
	v224 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v223))))
	v225 = F_SearchSysCache1(m, int32(14), v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L56
	}
L55:
	;
	goto L50
L56:
	;
	if v225 == int32(0) {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+22)))
	v233 = v231 + v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v229+v222))) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v233)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v236+v222))) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v233)+80))
	v241 = F_get_opfamily_proc(m, v240, v238, v238, v188)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	if v241 == int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	F_fmgr_info_cxt(m, v241, v245+v200*int32(28), v41)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v222+(base.I32_wrap_i64(v77)+v190))))
	*(*int32)(unsafe.Add(mBase, uint32(v251+v222))) = v254
	if v219 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v311+v222)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	F_get_typlenbyvalalign(m, v313, v314+v216, v316+v200, v318+v200)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L9
	} else {
		goto L72
	}
L62:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v265 = v258 + v259<<(uint(int32(3))%32) + v219*int32(100)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v256+v222))) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v270+v222))) = v272
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v265)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v274+v222))) = v276
	v310 = v202
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v202 == int32(0) {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v281 = F_exprType(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v50)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v283+v222))) = v281
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v287 = F_exprTypmod(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v289+v222))) = v287
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v293 = F_exprCollation(m, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v295+v222))) = v293
	v299 = v202 + int32(4)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if base.Ui32(v299) < base.Ui32(v302+v303<<(uint(int32(2))%32)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v308 = v299
	goto L71
L70:
	;
	v308 = int32(0)
	goto L71
L71:
	;
	v310 = v308
	goto L61
L72:
	;
	F_ReleaseCatCache(m, v225)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v325 = v200 + int32(1)
	v326 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+4)))
	if v325 < v326 {
		v200 = v325
		v202 = v310
		goto L54
	} else {
		goto L74
	}
L74:
	;
	goto L55
L75:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[2]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v354 != v350 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v41
	m.G0 = v27 + int32(80)
	goto L4
L77:
	;
	if v354 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	if v350 != 0 {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v359 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v358 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+28)) = v358
	goto L82
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+20)) = v358
	goto L82
L86:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v358)+24)) = v364
	goto L80
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v350
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v350)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v371
	if v371 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v41)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = int32(0)
	goto L79
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+24)) = v41
	goto L92
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+20)) = v41
	goto L76
L93:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v392
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_3), v27)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(99), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v406
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_6), v27+int32(16))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(119), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v422
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_7), v27+int32(32))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(202), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v233)+84))
	v443 = F_format_type_be(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v443
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v233 + int32(8)
	if v441 == int32(104) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v454 = int32(_a_F_RelationGetPartitionKey_8)
	goto L107
L106:
	;
	v454 = int32(_a_F_RelationGetPartitionKey_9)
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v454
	F_errmsg(m, int32(_a_F_RelationGetPartitionKey_10), v27+int32(48))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(221), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_11), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(240), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetPartitionQual(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+131)))
	if v3 == int32(1) {
		v6 = F_generate_partition_qual(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = v6
			return v11
		}
	} else {
		v11 = int32(0)
		return v11
	}
}
func F_RelationGetStatExtList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v11 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 - int32(-64)
	return v84
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v15 = F_list_copy(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v20 = v7 + int32(-56)
	v24 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v20, int32(2), int32(3), int32(184), v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v84 = v15
	goto L1
L7:
	;
	v29 = F_table_open(m, int32(3381), int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(1)
	v35 = F_systable_beginscan(m, v29, int32(3379), v32, int32(0), v32, v20)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v2
	v42 = v37
	goto L14
L12:
	;
	v54 = v2
	goto L13
L13:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L19
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46)))
	v49 = F_lappend_oid(m, v40, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v54 = v49
	goto L13
L16:
	;
	v51 = F_systable_getnext(m, v35)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v51 != 0 {
		v40 = v49
		v42 = v51
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_relation_close(m, v29, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_list_sort(m, v54, int32(502))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v67 = int32(_a_F_RelationGetStatExtList_0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0]))
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0])) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v74 = F_list_copy(m, v54)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v74
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0])) = v68
	F_list_free(m, v73)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v84 = v54
	goto L1
}
func F_RelationMapUpdateMap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L34
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L34
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(0) < v36 {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v10 = int32(_a_F_RelationMapUpdateMap_0)
	goto L7
L6:
	;
	v10 = int32(_a_F_RelationMapUpdateMap_1)
	goto L7
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapUpdateMap[0]))
	if v12 == int32(0) {
		v35 = v10
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapUpdateMap[1]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	goto L9
L9:
	;
	if int32(2) <= v17 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapUpdateMap[1]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	if v23 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v26&int32(1) != 0 {
		goto L2
	} else {
		goto L15
	}
L12:
	;
	v26 = int32(1)
	goto L14
L13:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+76)))
	v26 = v25
	goto L14
L14:
	;
	goto L11
L15:
	;
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v31 = int32(_a_F_RelationMapUpdateMap_2)
	goto L18
L17:
	;
	v31 = int32(_a_F_RelationMapUpdateMap_3)
	goto L18
L18:
	;
	if l3 != 0 {
		v35 = v31
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if l2 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v34 = int32(_a_F_RelationMapUpdateMap_4)
	goto L22
L21:
	;
	v34 = int32(_a_F_RelationMapUpdateMap_5)
	goto L22
L22:
	;
	v35 = v34
	goto L4
L23:
	;
	v44 = int32(0)
	goto L27
L24:
	;
	goto L25
L25:
	;
	v69 = v35 + v36<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v69)+8)) = l0
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v72 + int32(1)
	return
L26:
	;
	if int32(64) <= v36 {
		goto L1
	} else {
		goto L33
	}
L27:
	;
	v51 = v35 + int32(8) + v44<<(uint(int32(3))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 != l0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = l1
	return
L29:
	;
	v55 = v44 + int32(1)
	if v36 != v55 {
		v44 = v55
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L26
L33:
	;
	goto L25
L34:
	;
	return
L35:
	;
	F_errmsg_internal(m, int32(_a_F_RelationMapUpdateMap_6), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(349), int32(_a_F_RelationMapUpdateMap_8))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L34
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
	F_errmsg_internal(m, int32(_a_F_RelationMapUpdateMap_9), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(352), int32(_a_F_RelationMapUpdateMap_8))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L34
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errmsg_internal(m, int32(_a_F_RelationMapUpdateMap_10), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(404), int32(_a_F_RelationMapUpdateMap_11))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L34
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationRebuildRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
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
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
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
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int64
	_ = v523
	var v524 int64
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v686 int32
	_ = v686
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v725 int32
	_ = v725
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v950 int32
	_ = v950
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1067 int32
	_ = v1067
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1196 int32
	_ = v1196
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1224 int32
	_ = v1224
	var v1231 int32
	_ = v1231
	var v1246 int64
	_ = v1246
	var v1255 int32
	_ = v1255
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(320)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v27 = v25 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v27
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgrclose(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v32 = v22 + int32(76)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[0]))
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+76)) = v41
	v43 = int32(_a_F_RelationRebuildRelation_0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v32
	*(*int32)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[1])) = v32
	goto L7
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[1]))
	v41 = v36
	goto L8
L10:
	;
	goto L11
L11:
	;
	v38 = int32(_a_F_RelationRebuildRelation_0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[0])) = v38
	v41 = v38
	goto L8
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L14:
	;
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v58)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+119)))
	if v63|int32(32) != int32(105) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	m.G0 = v20 + int32(320)
	return
L19:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1084 = v20 + int32(44)
	v1085 = int32(276)
	base.MemoryCopy(m, v1084, v203, v1085)
	base.MemoryCopy(m, v203, l0, v1085)
	base.MemoryCopy(m, l0, v1084, v1085)
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v1092
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1091
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+16)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1095
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v203)+32))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+32)) = v1100
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1099
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+36)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1103
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+40)) = v1108
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1107
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v203)+44))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+44)) = v1112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1111
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+48)) = v1116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1115
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	base.MemoryCopy(m, v1115, v1119, int32(144))
	if v705 != 0 {
		goto L251
	} else {
		goto L252
	}
L20:
	;
	v1067 = int32(0)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L12
	} else {
		goto L248
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L12
	} else {
		goto L245
	}
L23:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v175 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v68 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+117)))
	if v71 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v173 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v173)
	goto L18
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v84 = F_ScanPgRelation(m, v80, base.B2i32(v80 != int32(2662)), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L31
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[2])))
	if v75&int32(1) != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if v84 == int32(0) {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	base.MemoryCopy(m, v88, v89+v90, int32(144))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v94 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_pfree(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_RelationParseRelOptions(m, l0, v84)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	F_pfree(m, v84)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	v104 = int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v105) < base.Ui32(int32(_a_F_RelationRebuildRelation_1)) {
		v114 = v104
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v114 != 0 {
		goto L26
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	if v109 == int32(99) {
		v114 = v104
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v112 = F_isTempToastNamespace(m, v109)
	mBase = m.M
	v114 = v112
	goto L41
L44:
	;
	v116 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v117 = F_SearchSysCache1(m, int32(34), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	if v117 == int32(0) {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)) = uint8(v125)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+13)) = uint8(v128)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v130)+14)) = uint8(v131)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+15)) = uint8(v134)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+16)) = uint8(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+17)) = uint8(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v142)+18)) = uint8(v143)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+19)) = uint8(v146)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v148)+20)) = uint8(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+21)) = uint8(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v154)+22)) = uint8(v155)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v159)+20)))
	v161 = int32(768)
	if v160&v161 != v161 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v167 = v165
	goto L49
L48:
	;
	v167 = int32(2)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v167
	F_ReleaseCatCache(m, v117)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	goto L26
L51:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L12
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v203 = F_RelationBuildDesc(m, v201, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L12
	} else {
		goto L58
	}
L54:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[2])))
	if v181 != int32(1) {
		goto L18
	} else {
		goto L55
	}
L55:
	;
	v184 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v189 = F_ScanPgRelation(m, v186, v184, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+22)))
	base.MemoryCopy(m, v191, v192+v193, int32(144))
	F_pfree(m, v189)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v199 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v199)
	goto L18
L58:
	;
	if v203 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_RelationRebuildRelation[3]))
	goto L62
L60:
	;
	goto L61
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v203)+52))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v225 != v227 {
		v686 = v2
		goto L68
	} else {
		goto L69
	}
L62:
	;
	if v208 != int32(0) {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v201
	F_errmsg_internal(m, int32(_a_F_RelationRebuildRelation_2), v20)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2683), int32(_a_F_RelationRebuildRelation_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v707 != 0 {
		goto L173
	} else {
		goto L174
	}
L68:
	;
	v705 = v686
	goto L67
L69:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v229 != v230 {
		v686 = v2
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if v225 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	if v390 != 0 {
		goto L106
	} else {
		goto L107
	}
L72:
	;
	v235 = v225 << (uint(int32(3)) % 32)
	v237 = int32(28)
	v246 = v2
	goto L73
L73:
	;
	v259 = int32(0)
	v261 = v246 * int32(100)
	v262 = v235 + v224 + v237 + v261
	v263 = int32(4)
	v264 = v262 + v263
	v265 = v261 + (v226 + v235 + v237)
	v267 = v265 + v263
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v267))))
	if base.B2i32(v270 == v259)|base.B2i32(v270 != v273) != 0 {
		v291 = v270
		v292 = v273
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v705 = int32(0)
	goto L67
L75:
	;
	if v291-v292 != 0 {
		v705 = v259
		goto L67
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v276 = v264
	v277 = v267
	goto L78
L78:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+1)))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+1)))
	if v281 == int32(0) {
		v291 = v281
		v292 = v280
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v291 = v281
	v292 = v280
	goto L76
L80:
	;
	v284 = int32(1)
	if v281 == v280 {
		v276 = v276 + v284
		v277 = v277 + v284
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v262)+68))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v265)+68))
	if v295 != v296 {
		v705 = int32(0)
		goto L67
	} else {
		goto L83
	}
L83:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+72)))
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265)+72)))
	if v299 != v300 {
		v705 = int32(0)
		goto L67
	} else {
		goto L84
	}
L84:
	;
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+80)))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265)+80)))
	if v303 != v304 {
		v705 = int32(0)
		goto L67
	} else {
		goto L85
	}
L85:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v262)+76))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v265)+76))
	if v307 != v308 {
		v705 = int32(0)
		goto L67
	} else {
		goto L86
	}
L86:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+82)))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+82)))
	if v311 != v312 {
		v705 = int32(0)
		goto L67
	} else {
		goto L87
	}
L87:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+83)))
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+83)))
	if v315 != v316 {
		v705 = int32(0)
		goto L67
	} else {
		goto L88
	}
L88:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+84)))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+84)))
	if v319 != v320 {
		v705 = int32(0)
		goto L67
	} else {
		goto L89
	}
L89:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+85)))
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+85)))
	if v323 != v324 {
		v705 = int32(0)
		goto L67
	} else {
		goto L90
	}
L90:
	;
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+86)))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+86)))
	if v327 != v328 {
		v705 = int32(0)
		goto L67
	} else {
		goto L91
	}
L91:
	;
	if v327 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+87)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+87)))
	if v342 != v343 {
		v705 = int32(0)
		goto L67
	} else {
		goto L95
	}
L93:
	;
	v333 = v246 << (uint(int32(3)) % 32)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v333)+35)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v333)+35)))
	if v335 == v337 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v705 = int32(0)
	goto L67
L95:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+89)))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+89)))
	if v346 != v347 {
		v705 = int32(0)
		goto L67
	} else {
		goto L96
	}
L96:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+90)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+90)))
	if v350 != v351 {
		v705 = int32(0)
		goto L67
	} else {
		goto L97
	}
L97:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+91)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+91)))
	if v354 != v355 {
		v705 = int32(0)
		goto L67
	} else {
		goto L98
	}
L98:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+92)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265)+92)))
	if v358 != v359 {
		v705 = int32(0)
		goto L67
	} else {
		goto L99
	}
L99:
	;
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+94)))
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265)+94)))
	if v362 != v363 {
		v705 = int32(0)
		goto L67
	} else {
		goto L100
	}
L100:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v262)+96))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v265)+96))
	if v365 == v366 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v369 = v246 + int32(1)
	if v369 == v225 {
		goto L71
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	goto L74
L104:
	;
	v246 = v369
	goto L73
L105:
	;
	v686 = int32(1)
	goto L68
L106:
	;
	if v389 == int32(0) {
		v686 = v2
		goto L68
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if v389 != 0 {
		v686 = v2
		goto L68
	} else {
		goto L170
	}
L109:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+16)))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+16)))
	if v393 != v394 {
		v686 = v2
		goto L68
	} else {
		goto L110
	}
L110:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+17)))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+17)))
	if v396 != v397 {
		v686 = v2
		goto L68
	} else {
		goto L111
	}
L111:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v390)+18)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+18)))
	if v399 != v400 {
		v686 = v2
		goto L68
	} else {
		goto L112
	}
L112:
	;
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+12)))
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+12)))
	if v402 != v403 {
		v686 = v2
		goto L68
	} else {
		goto L113
	}
L113:
	;
	if v402 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	if v487 != 0 {
		goto L131
	} else {
		goto L132
	}
L115:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v412 = int32(0)
	goto L116
L116:
	;
	v429 = v412 << (uint(int32(3)) % 32)
	v430 = v408 + v429
	v431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430))))
	v432 = v407 + v429
	v433 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432))))
	if v431 != v433 {
		v705 = int32(0)
		goto L67
	} else {
		goto L118
	}
L117:
	;
	v705 = int32(0)
	goto L67
L118:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v430)+4))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if base.B2i32(v439 == int32(0))|base.B2i32(v439 != v442) != 0 {
		v460 = v439
		v461 = v442
		goto L120
	} else {
		goto L121
	}
L119:
	;
	if v460-v461 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	goto L119
L121:
	;
	v445 = v435
	v446 = v436
	goto L122
L122:
	;
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+1)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+1)))
	if v450 == int32(0) {
		v460 = v450
		v461 = v449
		goto L120
	} else {
		goto L124
	}
L123:
	;
	v460 = v450
	v461 = v449
	goto L120
L124:
	;
	v453 = int32(1)
	if v450 == v449 {
		v445 = v445 + v453
		v446 = v446 + v453
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v466 = v412 + int32(1)
	if v466 == v402 {
		goto L114
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	goto L117
L129:
	;
	v412 = v466
	goto L116
L130:
	;
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v390)+14)))
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v389)+14)))
	if v556 != v557 {
		v686 = v2
		goto L68
	} else {
		goto L146
	}
L131:
	;
	if v486 == int32(0) {
		v686 = v2
		goto L68
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	if v486 != 0 {
		v686 = v2
		goto L68
	} else {
		goto L145
	}
L134:
	;
	if v225 <= int32(0) {
		goto L130
	} else {
		goto L135
	}
L135:
	;
	v497 = int32(0)
	v501 = v225
	goto L136
L136:
	;
	v514 = v497 << (uint(int32(4)) % 32)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v390)+8))
	v516 = v514 + v515
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v516))))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	v519 = v518 + v514
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519))))
	if v517 != v520 {
		v705 = int32(0)
		goto L67
	} else {
		goto L138
	}
L137:
	;
	goto L130
L138:
	;
	if v517 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v523 = *(*int64)(unsafe.Add(mBase, uint32(v516)+8))
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v519)+8))
	v527 = v224 + int32(28) + v497<<(uint(int32(3))%32)
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+4)))
	v529 = int32(*(*int16)(unsafe.Add(mBase, uint32(v527)+2)))
	v530 = F_datumIsEqual(m, v523, v524, v528, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L142
	}
L140:
	;
	v535 = v501
	goto L141
L141:
	;
	v537 = v497 + int32(1)
	if v537 < v535 {
		v497 = v537
		v501 = v535
		goto L136
	} else {
		goto L144
	}
L142:
	;
	if v530 == int32(0) {
		v705 = int32(0)
		goto L67
	} else {
		goto L143
	}
L143:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v535 = v534
	goto L141
L144:
	;
	goto L137
L145:
	;
	goto L130
L146:
	;
	if v556 == int32(0) {
		goto L105
	} else {
		goto L147
	}
L147:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v568 = int32(0)
	goto L148
L148:
	;
	v582 = v568 * int32(12)
	v583 = v562 + v582
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	v585 = v582 + v561
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)))
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586))))
	if base.B2i32(v589 == int32(0))|base.B2i32(v589 != v592) != 0 {
		v610 = v589
		v611 = v592
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L105
L150:
	;
	if v610-v611 != 0 {
		v686 = v2
		goto L68
	} else {
		goto L157
	}
L151:
	;
	goto L150
L152:
	;
	v595 = v584
	v596 = v586
	goto L153
L153:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+1)))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595)+1)))
	if v600 == int32(0) {
		v610 = v600
		v611 = v599
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v610 = v600
	v611 = v599
	goto L151
L155:
	;
	v603 = int32(1)
	if v600 == v599 {
		v595 = v595 + v603
		v596 = v596 + v603
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v583)+4))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	if base.B2i32(v617 == int32(0))|base.B2i32(v617 != v620) != 0 {
		v638 = v617
		v639 = v620
		goto L159
	} else {
		goto L160
	}
L158:
	;
	if v638-v639 != 0 {
		v686 = v2
		goto L68
	} else {
		goto L165
	}
L159:
	;
	goto L158
L160:
	;
	v623 = v613
	v624 = v614
	goto L161
L161:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624)+1)))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623)+1)))
	if v628 == int32(0) {
		v638 = v628
		v639 = v627
		goto L159
	} else {
		goto L163
	}
L162:
	;
	v638 = v628
	v639 = v627
	goto L159
L163:
	;
	v631 = int32(1)
	if v628 == v627 {
		v623 = v623 + v631
		v624 = v624 + v631
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+8)))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+8)))
	if v641 != v642 {
		v686 = v2
		goto L68
	} else {
		goto L166
	}
L166:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+9)))
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+9)))
	if v644 != v645 {
		v686 = v2
		goto L68
	} else {
		goto L167
	}
L167:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+10)))
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585)+10)))
	if v647 != v648 {
		v686 = v2
		goto L68
	} else {
		goto L168
	}
L168:
	;
	v651 = v568 + int32(1)
	if v556 != v651 {
		v568 = v651
		goto L148
	} else {
		goto L169
	}
L169:
	;
	goto L149
L170:
	;
	goto L105
L171:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v803|v804 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L172:
	;
	v802 = int32(1)
	goto L171
L173:
	;
	if v706 == int32(0) {
		v802 = v2
		goto L171
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	if v706 != 0 {
		v802 = v2
		goto L171
	} else {
		goto L190
	}
L176:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v706)))
	if v710 != v711 {
		v802 = v2
		goto L171
	} else {
		goto L177
	}
L177:
	;
	if v710 <= int32(0) {
		goto L172
	} else {
		goto L178
	}
L178:
	;
	v725 = v2
	goto L179
L179:
	;
	v733 = v725 << (uint(int32(2)) % 32)
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v707)+4))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v733+v734)))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v736)))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v706)+4))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v738+v733)))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	if v737 != v741 {
		v802 = v2
		goto L171
	} else {
		goto L181
	}
L180:
	;
	goto L172
L181:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v736)+4))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	if v743 != v744 {
		v802 = v2
		goto L171
	} else {
		goto L182
	}
L182:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+16)))
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+16)))
	if v746 != v747 {
		v802 = v2
		goto L171
	} else {
		goto L183
	}
L183:
	;
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736)+17)))
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740)+17)))
	if v749 != v750 {
		v802 = v2
		goto L171
	} else {
		goto L184
	}
L184:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v736)+8))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	v754 = F_equal(m, v752, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L12
	} else {
		goto L185
	}
L185:
	;
	if v754 == int32(0) {
		v802 = v2
		goto L171
	} else {
		goto L186
	}
L186:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v736)+12))
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v740)+12))
	v760 = F_equal(m, v758, v759)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L12
	} else {
		goto L187
	}
L187:
	;
	if v760 == int32(0) {
		v802 = v2
		goto L171
	} else {
		goto L188
	}
L188:
	;
	v765 = v725 + int32(1)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	if v765 < v766 {
		v725 = v765
		goto L179
	} else {
		goto L189
	}
L189:
	;
	goto L180
L190:
	;
	goto L172
L191:
	;
	v1067 = int32(1)
	goto L19
L192:
	;
	goto L193
L193:
	;
	v809 = int32(0)
	if base.B2i32(v803 != v809)^base.B2i32(v804 != v809) != 0 {
		v1067 = v809
		goto L19
	} else {
		goto L194
	}
L194:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	if v816 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	v818 = v817
	goto L197
L196:
	;
	v818 = int32(0)
	goto L197
L197:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v804)+4))
	if v819 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	v822 = v820
	goto L200
L199:
	;
	v822 = int32(0)
	goto L200
L200:
	;
	if v822 != v818 {
		goto L20
	} else {
		goto L201
	}
L201:
	;
	v838 = v2
	goto L202
L202:
	;
	v841 = int32(0)
	if v816 == v841 {
		v851 = v841
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v852 = int32(1)
	if v819 == int32(0) {
		v1067 = v852
		goto L19
	} else {
		goto L207
	}
L205:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	if v845 <= v838 {
		v851 = int32(0)
		goto L204
	} else {
		goto L206
	}
L206:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v816)+12))
	v851 = v847 + v838<<(uint(int32(2))%32)
	goto L204
L207:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	if base.B2i32(v851 == int32(0))|base.B2i32(v857 <= v838) != 0 {
		v1067 = v852
		goto L19
	} else {
		goto L208
	}
L208:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v819)+12))
	if v860 == int32(0) {
		v1067 = v852
		goto L19
	} else {
		goto L209
	}
L209:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v860+v838<<(uint(int32(2))%32))))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v851)))
	if v867 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v838 = v838 + int32(1)
	goto L202
L211:
	;
	v868 = int32(0)
	if v866 == v868 {
		v1067 = v868
		goto L19
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	if v866 != 0 {
		goto L20
	} else {
		goto L244
	}
L214:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+4)))
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+4)))
	if v871 != v872 {
		v1067 = v868
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+12)))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+12)))
	if v874 != v875 {
		v1067 = v868
		goto L19
	} else {
		goto L216
	}
L216:
	;
	v877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+24)))
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+24)))
	if v877 != v878 {
		v1067 = v868
		goto L19
	} else {
		goto L217
	}
L217:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880))))
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881))))
	if base.B2i32(v884 == int32(0))|base.B2i32(v884 != v887) != 0 {
		v905 = v884
		v906 = v887
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if v905-v906 != 0 {
		v1067 = v868
		goto L19
	} else {
		goto L225
	}
L219:
	;
	goto L218
L220:
	;
	v890 = v880
	v891 = v881
	goto L221
L221:
	;
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+1)))
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890)+1)))
	if v895 == int32(0) {
		v905 = v895
		v906 = v894
		goto L219
	} else {
		goto L223
	}
L222:
	;
	v905 = v895
	v906 = v894
	goto L219
L223:
	;
	v898 = int32(1)
	if v895 == v894 {
		v890 = v890 + v898
		v891 = v891 + v898
		goto L221
	} else {
		goto L224
	}
L224:
	;
	goto L222
L225:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v867)+8))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v908)+16))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v910)+16))
	if v909 != v911 {
		v1067 = v868
		goto L19
	} else {
		goto L226
	}
L226:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v908)+8))
	if v913 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v908)+4))
	v923 = (v916<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L229
L228:
	;
	v923 = v913
	goto L229
L229:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v910)+8))
	if v924 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v910)+4))
	v934 = (v927<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L232
L231:
	;
	v934 = v924
	goto L232
L232:
	;
	if int32(0) < v909 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v950 = int32(0)
	goto L236
L234:
	;
	goto L235
L235:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v867)+16))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v866)+16))
	v986 = F_equal(m, v984, v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L12
	} else {
		goto L240
	}
L236:
	;
	v958 = v950 << (uint(int32(2)) % 32)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v908+v923+v958)))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v910+v934+v958)))
	if v960 != v962 {
		v1067 = v868
		goto L19
	} else {
		goto L238
	}
L237:
	;
	goto L235
L238:
	;
	v965 = v950 + int32(1)
	if v965 != v909 {
		v950 = v965
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	if v986 == int32(0) {
		v1067 = v868
		goto L19
	} else {
		goto L241
	}
L241:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v867)+20))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v866)+20))
	v992 = F_equal(m, v990, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L12
	} else {
		goto L242
	}
L242:
	;
	if v992 == int32(0) {
		v1067 = v868
		goto L19
	} else {
		goto L243
	}
L243:
	;
	goto L210
L244:
	;
	goto L210
L245:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v1019
	F_errmsg_internal(m, int32(_a_F_RelationRebuildRelation_5), v20+int32(32))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2337), int32(_a_F_RelationRebuildRelation_6))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v1035
	F_errmsg_internal(m, int32(_a_F_RelationRebuildRelation_7), v20+int32(16))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2308), int32(_a_F_RelationRebuildRelation_6))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L12
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v203)+52))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+52)) = v1123
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1122
	goto L253
L252:
	;
	goto L253
L253:
	;
	if v802 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v1128
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v1127
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v203)+72))
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+72)) = v1132
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v1131
	goto L256
L255:
	;
	goto L256
L256:
	;
	if v1067 != 0 {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+80)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v1136
	goto L259
L258:
	;
	goto L259
L259:
	;
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v203)+264))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+264)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v1141
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v203)+272))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+272)) = v1146
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1145
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+268)))
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+268)) = uint8(v1150)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v1149)
	if v1082 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v203)+92))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+92)) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1153
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v203)+96))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+96)) = v1158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1157
	goto L262
L261:
	;
	goto L262
L262:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v203)+104))
	if v1162 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	F_RelationDestroyRelation(m, v203, v705^int32(1))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L12
	} else {
		goto L310
	}
L264:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v203)+112))
	if v1165 == int32(0) {
		goto L263
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1168 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1168
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v203)+104))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1175 != 0 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	goto L266
L268:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v203)+112))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1210 != 0 {
		goto L290
	} else {
		goto L291
	}
L269:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+16))
	if v1179 != v1175 {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	goto L271
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1174
	goto L268
L272:
	;
	goto L268
L273:
	;
	if v1179 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	goto L272
L276:
	;
	if v1175 != 0 {
		goto L283
	} else {
		goto L284
	}
L277:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+28))
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+24))
	if v1184 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	if v1183 == int32(0) {
		goto L276
	} else {
		goto L282
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+28)) = v1183
	goto L278
L280:
	;
	goto L281
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1179)+20)) = v1183
	goto L278
L282:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1174)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1183)+24)) = v1189
	goto L276
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+16)) = v1175
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+28)) = v1196
	if v1196 != 0 {
		goto L286
	} else {
		goto L287
	}
L284:
	;
	goto L285
L285:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1174)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+16)) = int32(0)
	goto L275
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+24)) = v1174
	goto L288
L287:
	;
	goto L288
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1175)+20)) = v1174
	goto L272
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+116)) = int32(0)
	v1246 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v203)+108)) = v1246
	*(*int64)(unsafe.Add(mBase, uint32(v203)+100)) = v1246
	goto L263
L290:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+16))
	if v1214 != v1210 {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	goto L292
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v1209
	goto L289
L293:
	;
	goto L289
L294:
	;
	if v1214 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L295:
	;
	goto L296
L296:
	;
	goto L293
L297:
	;
	if v1210 != 0 {
		goto L304
	} else {
		goto L305
	}
L298:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+28))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+24))
	if v1219 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	if v1218 == int32(0) {
		goto L297
	} else {
		goto L303
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1219)+28)) = v1218
	goto L299
L301:
	;
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1214)+20)) = v1218
	goto L299
L303:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+24)) = v1224
	goto L297
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+16)) = v1210
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+28)) = v1231
	if v1231 != 0 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	goto L306
L306:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1209)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1209)+16)) = int32(0)
	goto L296
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1231)+24)) = v1209
	goto L309
L308:
	;
	goto L309
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1210)+20)) = v1209
	goto L293
L310:
	;
	goto L18
}
func F_RelationSupportsSysCache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_RelationSupportsSysCache[0]))
	v10 = v8 - int32(1)
	if v10 < v2 {
		v42 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v42
L2:
	;
	v14 = v10
	v15 = v2
	goto L3
L3:
	;
	v20 = int32(2)
	v21 = base.I32_div_s(v14-v15, v20)
	v22 = v21 + v15
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(v20)%32))+uint32(_c_F_RelationSupportsSysCache[1])))
	v28 = base.B2i32(v27 == l0)
	if v27 == l0 {
		v42 = v28
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v42 = v28
	goto L1
L5:
	;
	v31 = base.B2i32(base.Ui32(v27) < base.Ui32(l0))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = v22 + int32(1)
	goto L8
L7:
	;
	v32 = v15
	goto L8
L8:
	;
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v35 = v14
	goto L11
L10:
	;
	v35 = v22 - int32(1)
	goto L11
L11:
	;
	if v32 <= v35 {
		v14 = v35
		v15 = v32
		goto L3
	} else {
		goto L12
	}
L12:
	;
	goto L4
}
func F_SetRelationNumChecks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v16 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
		v18 = F_SearchSysCacheCopy(m, int32(57), v16, int64(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				v22 = v20 + v21
				v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+122)))
				if v23 != v2 {
					*(*uint16)(unsafe.Add(mBase, uint32(v22)+122)) = uint16(v2)
					F_CatalogTupleUpdate(m, v13, v18+int32(4), v18)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_pfree(m, v18)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_relation_close(m, v13, int32(3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				} else {
					F_CacheInvalidateRelcache(m, l0)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_pfree(m, v18)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							F_relation_close(m, v13, int32(3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
					F_errmsg_internal(m, int32(_a_F_SetRelationNumChecks_0), v9)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetRelationNumChecks_1), int32(3196), int32(_a_F_SetRelationNumChecks_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
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
func F_UnlockRelationId(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	Fn14224(m, l0, l1, int32(0))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_check_relation_block_range(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	if base.Ui64(l1) < base.Ui64(int64(4294967295)) {
		v11 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if base.Ui32(v11) <= base.Ui32(base.I32_wrap_i64(l1)) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = l1
						F_errmsg(m, int32(_a_F_check_relation_block_range_0), v6+int32(16))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_relation_block_range_1), int32(215), int32(_a_F_check_relation_block_range_2))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
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
				m.G0 = v6 + int32(32)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v6))) = l1
				F_errmsg(m, int32(_a_F_check_relation_block_range_3), v6)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_relation_block_range_1), int32(210), int32(_a_F_check_relation_block_range_2))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
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
func F_relation_close(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	F_RelationClose(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		if l1 != 0 {
			F_UnlockRelationId(m, v6+int32(8), l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		} else {
			m.G0 = v6 + int32(16)
			return
		}
	}
}
func F_try_relation_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 == int32(0) {
		v13 = int64(0)
		v16 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(l0), v13, v13, v13)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				v69 = int32(0)
				m.G0 = v6 + int32(16)
				return v69
			} else {
				v35 = F_RelationIdGetRelation(m, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					if v35 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
							F_errmsg_internal(m, int32(_a_F_try_relation_open_0), v6)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_try_relation_open_1), int32(116), int32(_a_F_try_relation_open_2))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+118)))
						if v40 == int32(116) {
							v43 = int32(_a_F_try_relation_open_3)
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0])) = v45 | int32(1)
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
						switch v51 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_relation_open[1])))
							if v55 == int32(0) {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
								if v58 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = int32(0)
								} else {
								}
								v64 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v64
								*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v64)
							} else {
								v61 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v61)
							}
						default:
							v64 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v64)
						}
						v69 = v35
						m.G0 = v6 + int32(16)
						return v69
					}
				}
			}
		}
	} else {
		F_LockRelationOid(m, l0, l1)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = int64(0)
			v29 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(l0), v26, v26, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				if v29 != 0 {
					v35 = F_RelationIdGetRelation(m, l0)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
								F_errmsg_internal(m, int32(_a_F_try_relation_open_0), v6)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_try_relation_open_1), int32(116), int32(_a_F_try_relation_open_2))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+118)))
							if v40 == int32(116) {
								v43 = int32(_a_F_try_relation_open_3)
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0])) = v45 | int32(1)
							} else {
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
							switch v51 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_relation_open[1])))
								if v55 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)+272))
									if v58 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = int32(0)
									} else {
									}
									v64 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v64
									*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v64)
								} else {
									v61 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v61)
								}
							default:
								v64 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v35)+272)) = v64
								*(*uint8)(unsafe.Add(mBase, uint32(v35)+268)) = uint8(v64)
							}
							v69 = v35
							m.G0 = v6 + int32(16)
							return v69
						}
					}
				} else {
					F_UnlockRelationOid(m, l0, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v69 = int32(0)
						m.G0 = v6 + int32(16)
						return v69
					}
				}
			}
		}
	}
}
