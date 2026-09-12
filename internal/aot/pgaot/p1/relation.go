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
	var v76 int32
	_ = v76
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
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
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
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
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
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
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v854 int32
	_ = v854
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v940 int32
	_ = v940
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1018 int32
	_ = v1018
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1040 int32
	_ = v1040
	var v1045 int32
	_ = v1045
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	v5 = l4
	v8 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(352)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
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
		v163 = v8
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l2 == int32(0) {
		v1200 = v163
		v1204 = v34
		goto L24
	} else {
		goto L25
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 <= int32(0) {
		v163 = v8
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v68 = v8
	v76 = v8
	goto L11
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v68<<(uint(int32(2))%32))))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v96 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
	v99 = v91 + v92<<(uint(int32(4))%32) + v96*int32(100)
	v101 = v99 - int32(80)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+68))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+76))
	v106 = int32(*(*int8)(unsafe.Add(mBase, uint32(v99)+10)))
	v107 = F_cookDefault(m, v36, v90, v102, v103, v99-int32(76), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v163 = v140
	goto L8
L13:
	;
	v143 = v68 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v143 < v144 {
		v68 = v143
		v76 = v140
		goto L11
	} else {
		goto L23
	}
L14:
	;
	if v107 == int32(0) {
		v140 = v76
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
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
	v117 = F_StoreAttrDefault(m, l0, v116, v107, l5)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L20
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
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+24)))
	if v115 != 0 {
		v140 = v76
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v120 = F_palloc(m, int32(28))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v122 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = int32(2)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+26)) = uint8(v122)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+22)) = uint8(v5)
	v132 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+20)) = uint16(v132)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v107
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+12)) = uint16(v127)
	v136 = F_lappend(m, v76, v120)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v140 = v136
	goto L13
L23:
	;
	goto L12
L24:
	;
	F_SetRelationNumChecks(m, l0, v1204)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L288
	}
L25:
	;
	v174 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v175 <= v174 {
		v1200 = v163
		v1204 = v34
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v197 = v8
	v199 = v163
	v201 = v174
	v203 = v34
	v204 = v8
	goto L34
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L4
	} else {
		goto L284
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L4
	} else {
		goto L280
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L4
	} else {
		goto L276
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L4
	} else {
		goto L272
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L4
	} else {
		goto L268
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L4
	} else {
		goto L264
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L4
	} else {
		goto L260
	}
L34:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208+v201<<(uint(int32(2))%32))))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	switch v213 - int32(1) {
	case 0:
		goto L38
	default:
		v987 = v197
		v989 = v199
		v993 = v203
		v994 = v204
		goto L37
	case 4:
		goto L39
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L256
	}
L36:
	;
	goto L35
L37:
	;
	v999 = v201 + int32(1)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v999 < v1000 {
		v197 = v987
		v199 = v989
		v201 = v999
		v203 = v993
		v204 = v994
		goto L34
	} else {
		goto L255
	}
L38:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)+12))
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+4))
	v700 = F_get_attnum(m, v695, v699)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L182
	}
L39:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	if v216 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	if v256 != 0 {
		goto L58
	} else {
		goto L59
	}
L41:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v219 = F_transformExpr(m, v36, v216, int32(28))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v212)+24))
	v251 = F_stringToNode(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L55
	}
L44:
	;
	v222 = F_coerce_to_boolean(m, v36, v219, int32(534383))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_assign_expr_collations(m, v36, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v226 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v227 == int32(1) {
		v255 = v222
		goto L40
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v217 + int32(4)
	F_errmsg(m, int32(90142), v29+int32(144))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(496316), int32(3433), int32(90671))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v255 = v251
	goto L40
L56:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+14)))
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+16)))
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	v671 = F_StoreRelCheck(m, l0, v651, v255, v668, v669, v5, v53, v670, l5)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L4
	} else {
		goto L179
	}
L57:
	;
	F_systable_endscan(m, v404)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L4
	} else {
		goto L177
	}
L58:
	;
	if v197 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v609 = int32(0)
	v611 = F_pull_var_clause(m, v255, v609)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L4
	} else {
		goto L170
	}
L61:
	;
	v370 = F_lappend(m, v197, v256)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L85
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v259 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v262 = int32(0)
	if v262 < v259 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v265 = v259
	goto L66
L65:
	;
	v265 = v262
	goto L66
L66:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v274 = int32(0)
	goto L67
L67:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v266+v274<<(uint(int32(2))%32))))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v301 == int32(0) {
		v320 = v300
		v321 = v301
		goto L70
	} else {
		goto L71
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L81
	}
L69:
	;
	if v321-v320 != 0 {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	goto L69
L71:
	;
	if v300 != v301 {
		v320 = v300
		v321 = v301
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v305 = v297
	v306 = v256
	goto L73
L73:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	if v310 == int32(0) {
		v320 = v309
		v321 = v310
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v320 = v309
	v321 = v310
	goto L70
L75:
	;
	v313 = int32(1)
	if v309 == v310 {
		v305 = v305 + v313
		v306 = v306 + v313
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v324 = v274 + int32(1)
	if v265 != v324 {
		v274 = v324
		goto L67
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L68
L80:
	;
	goto L61
L81:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v256
	F_errmsg(m, int32(116279), v29+int32(128))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(496316), int32(2526), int32(120306))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+16)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+14)))
	v377 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v29+int32(208), int32(9), int32(3), int32(184), v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_ScanKeyInit(m, v29+int32(256), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_ScanKeyInit(m, v29+int32(304), int32(2), int32(3), int32(62), v256)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v404 = F_systable_beginscan(m, v377, int32(2665), int32(1), int32(0), int32(3), v29+int32(208))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	v406 = F_systable_getnext(m, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	if v406 == int32(0) {
		goto L57
	} else {
		goto L92
	}
L92:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v406)+16))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+22)))
	v413 = v411 + v412
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+72)))
	if v414 == int32(99) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v377)+52))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+20)))
	if v418&int32(1) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v486 = int32(0)
	goto L95
L95:
	;
	if v5 == int32(0) {
		v496 = l3
		goto L122
	} else {
		goto L123
	}
L96:
	;
	v480 = F_text_to_cstring(m, v479)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L119
	}
L97:
	;
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v426))))
	v479 = v476
	goto L96
L98:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v417)+452))
	if int32(0) <= v423 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+26)))
	if v452&int32(8) != 0 {
		goto L112
	} else {
		goto L113
	}
L101:
	;
	v426 = v413 + v423
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417)+458)))
	if v427 != int32(1) {
		v479 = v426
		goto L96
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v450 = F_nocachegetattr(m, v406, int32(28), v417)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L111
	}
L104:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v417)+456)))
	switch v430 - int32(1) {
	case 0:
		goto L97
	case 1:
		goto L107
	default:
		goto L105
	case 3:
		goto L106
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L108
	}
L106:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v479 = v434
	goto L96
L107:
	;
	v433 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426))))
	v479 = v433
	goto L96
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = base.I32_extend16_s(v430)
	F_errmsg_internal(m, int32(483562), v29)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(326787), int32(70), int32(67821))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	v479 = v450
	goto L96
L112:
	;
	v456 = F_nocachegetattr(m, v406, int32(28), v417)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L4
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L116
	}
L115:
	;
	v479 = v456
	goto L96
L116:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v462 + int32(4)
	F_errmsg_internal(m, int32(185975), v29+int32(112))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(496316), int32(2761), int32(90778))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v482 = F_stringToNode(m, v480)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	v484 = F_equal(m, v255, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v486 = v484
	goto L95
L122:
	;
	if v496&v486 == int32(0) {
		goto L36
	} else {
		goto L125
	}
L123:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+103)))
	if v490 != 0 {
		v496 = l3
		goto L122
	} else {
		goto L124
	}
L124:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+131)))
	v496 = l3 | (v492 ^ int32(1))
	goto L122
L125:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+106)))
	if v500 == int32(1) {
		goto L33
	} else {
		goto L126
	}
L126:
	;
	v505 = int32(*(*int16)(unsafe.Add(mBase, uint32(v413)+104)))
	if v372&int32(1)&base.B2i32(int32(0) < v505) != 0 {
		goto L32
	} else {
		goto L127
	}
L127:
	;
	if v373&int32(1) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v5 != 0 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+75)))
	if v513 != int32(1) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+76)))
	if v516 == int32(0) {
		goto L31
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	if (v374|v53)&int32(1) == int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	if v374&int32(1) == int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+75)))
	if v523 != int32(1) {
		goto L30
	} else {
		goto L135
	}
L135:
	;
	goto L132
L136:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+75)))
	if v531 == int32(1) {
		goto L30
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v536 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L4
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	if v536 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v256
	F_errmsg(m, int32(250662), v29+int32(48))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L4
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v549 = F_heap_copytuple(m, v406)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L4
	} else {
		goto L146
	}
L144:
	;
	F_errfinish(m, int32(496316), int32(2826), int32(90778))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v549)+16))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+22)))
	v553 = v551 + v552
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v554)+131)))
	if v555 == int32(1) {
		goto L151
	} else {
		goto L152
	}
L147:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L165
	}
L148:
	;
	if v372&int32(1) != 0 {
		goto L156
	} else {
		goto L157
	}
L149:
	;
	v566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v553)+104)))
	v568 = v566 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+104)) = uint16(v568)
	if base.I32_extend16_s(v568) != v568 {
		goto L147
	} else {
		goto L155
	}
L150:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+103)) = uint8(v564)
	goto L148
L151:
	;
	v558 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+104)) = uint16(v558)
	v564 = int32(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	if v5 == int32(0) {
		goto L149
	} else {
		goto L154
	}
L154:
	;
	v564 = int32(1)
	goto L150
L155:
	;
	goto L148
L156:
	;
	v575 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+106)) = uint8(v575)
	goto L158
L157:
	;
	goto L158
L158:
	;
	if v374&int32(1) == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	F_CatalogTupleUpdate(m, v377, v549+int32(4), v549)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L162
	}
L160:
	;
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v553)+75)))
	if v581 != 0 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v582 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+75)) = uint16(v582)
	goto L159
L162:
	;
	F_systable_endscan(m, v404)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_sequence_close(m, v377, int32(3))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v987 = v370
	v989 = v199
	v993 = v203
	v994 = v204
	goto L37
L165:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	F_errmsg(m, int32(120714), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(496316), int32(2849), int32(90778))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v628)+68))
	v633 = F_ChooseConstraintName(m, v628+int32(4), v627, int32(318478), v632, v197)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L175
	}
L170:
	;
	v613 = F_list_union(m, v611)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	if v613 == int32(0) {
		v627 = v609
		goto L169
	} else {
		goto L172
	}
L172:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v617 != int32(1) {
		v627 = v609
		goto L169
	} else {
		goto L173
	}
L173:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v613)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v621)))
	v623 = int32(*(*int16)(unsafe.Add(mBase, uint32(v622)+8)))
	v625 = F_get_attname(m, v620, v623, int32(1))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	v627 = v625
	goto L169
L175:
	;
	v635 = F_lappend(m, v197, v633)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v651 = v633
	v657 = v635
	goto L56
L177:
	;
	F_sequence_close(m, v377, int32(3))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	v651 = v256
	v657 = v370
	goto L56
L179:
	;
	v674 = F_palloc(m, int32(28))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+16)) = v255
	v677 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v674)+12)) = uint16(v677)
	*(*int32)(unsafe.Add(mBase, uint32(v674)+8)) = v651
	*(*int32)(unsafe.Add(mBase, uint32(v674)+4)) = v671
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = int32(5)
	v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+20)) = uint8(v683)
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v674)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+21)) = uint8(v685)
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v674)+26)) = uint8(v689)
	v693 = F_lappend(m, v199, v674)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	v987 = v657
	v989 = v693
	v993 = v203 + int32(1)
	v994 = v204
	goto L37
L182:
	;
	if v700 == int32(0) {
		goto L29
	} else {
		goto L183
	}
L183:
	;
	if v700 < int32(0) {
		goto L28
	} else {
		goto L184
	}
L184:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+15)))
	v710 = m.G0
	v712 = v710 - int32(96)
	m.G0 = v712
	v714 = F_findNotNullConstraintAttnum(m, v706, v700)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L189
	}
L185:
	;
	if v714 != int32(0) {
		v987 = v197
		v989 = v199
		v993 = v203
		v994 = v204
		goto L37
	} else {
		goto L243
	}
L186:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L4
	} else {
		goto L236
	}
L187:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L4
	} else {
		goto L230
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L4
	} else {
		goto L224
	}
L189:
	;
	if v714 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v718 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L4
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	m.G0 = v712 + int32(96)
	goto L185
L193:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720)+22)))
	v722 = v720 + v721
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+106)))
	if v708 != v723 {
		goto L188
	} else {
		goto L194
	}
L194:
	;
	if v709 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+76)))
	if v727 == int32(0) {
		goto L187
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if v5 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	if v5 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L200:
	;
	if v707 == int32(0) {
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v735 = v722 + int32(4)
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735))))
	v739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	if v739 == int32(0) {
		v758 = v738
		v759 = v739
		goto L203
	} else {
		goto L204
	}
L202:
	;
	if v759-v758 != 0 {
		goto L186
	} else {
		goto L210
	}
L203:
	;
	goto L202
L204:
	;
	if v738 != v739 {
		v758 = v738
		v759 = v739
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v743 = v707
	v744 = v735
	goto L206
L206:
	;
	v747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
	if v748 == int32(0) {
		v758 = v747
		v759 = v748
		goto L203
	} else {
		goto L208
	}
L207:
	;
	v758 = v747
	v759 = v748
	goto L203
L208:
	;
	v751 = int32(1)
	if v747 == v748 {
		v743 = v743 + v751
		v744 = v744 + v751
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	goto L199
L211:
	;
	F_sequence_close(m, v718, int32(3))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L223
	}
L212:
	;
	F_CatalogTupleUpdate(m, v718, v714+int32(4), v714)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L222
	}
L213:
	;
	v764 = int32(*(*int16)(unsafe.Add(mBase, uint32(v722)+104)))
	v766 = v764 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v722)+104)) = uint16(v766)
	if base.I32_extend16_s(v766) == v766 {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v722)+103)))
	if v786 != 0 {
		goto L211
	} else {
		goto L221
	}
L216:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L218
	}
L218:
	;
	F_errmsg(m, int32(120714), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(493315), int32(803), int32(418034))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	v787 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v722)+103)) = uint8(v787)
	goto L212
L222:
	;
	goto L211
L223:
	;
	goto L192
L224:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	v814 = F_get_rel_name(m, v706)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+84)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v712)+80)) = v722 + int32(4)
	F_errmsg(m, int32(706246), v712+int32(80))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+64)) = int32(521090)
	F_errhint(m, int32(605334), v712-int32(-64))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(493315), int32(767), int32(418034))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L4
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L231
	}
L231:
	;
	v844 = F_get_rel_name(m, v706)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+52)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v712)+48)) = v722 + int32(4)
	F_errmsg(m, int32(706323), v712+int32(48))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+32)) = int32(518465)
	F_errhint(m, int32(605294), v712+int32(32))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L4
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(493315), int32(779), int32(418034))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L4
	} else {
		goto L237
	}
L237:
	;
	v875 = F_get_attname(m, v706, v700, int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L4
	} else {
		goto L238
	}
L238:
	;
	v877 = F_get_rel_name(m, v706)
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712)+24)) = v877
	*(*int32)(unsafe.Add(mBase, uint32(v712)+20)) = v875
	*(*int32)(unsafe.Add(mBase, uint32(v712)+16)) = v707
	F_errmsg(m, int32(720453), v712+int32(16))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v712))) = v735
	F_errdetail(m, int32(619169), v712)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(493315), int32(795), int32(418034))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	if v896 != 0 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v915 = F_lappend(m, v204, v914)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L251
	}
L245:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v899 = F_ConstraintNameIsUsed(m, int32(0), v898, v896)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+12))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v906)))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v902)+68))
	v911 = F_ChooseConstraintName(m, v902+int32(4), v908, int32(302258), v910, v204)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L4
	} else {
		goto L250
	}
L248:
	;
	if v899 != 0 {
		goto L27
	} else {
		goto L249
	}
L249:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	v914 = v901
	goto L244
L250:
	;
	v914 = v911
	goto L244
L251:
	;
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	v918 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+208)) = uint16(v700)
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v920)+68))
	v923 = int32(0)
	v925 = int32(1)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v940 = int32(32)
	v950 = F_CreateConstraintEntry(m, v914, v921, int32(110), v923, v923, v925, v918, v923, v927, v29+int32(208), v925, v925, v923, v923, v923, v923, v923, v923, v923, v923, v940, v940, v923, v923, v940, v923, v923, v923, v5, v53, v917, v923, v923)
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L4
	} else {
		goto L252
	}
L252:
	;
	v953 = F_palloc(m, int32(28))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	v955 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+20)) = uint8(v955)
	*(*int32)(unsafe.Add(mBase, uint32(v953)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v953)+12)) = uint16(v700)
	*(*int32)(unsafe.Add(mBase, uint32(v953)+8)) = v914
	*(*int32)(unsafe.Add(mBase, uint32(v953)+4)) = v950
	*(*int32)(unsafe.Add(mBase, uint32(v953))) = v955
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v953)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+21)) = uint8(v964)
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v953)+26)) = uint8(v968)
	v970 = F_lappend(m, v199, v953)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	v987 = v197
	v989 = v970
	v993 = v203
	v994 = v915
	goto L37
L255:
	;
	v1200 = v989
	v1204 = v993
	goto L24
L256:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v1009 + int32(4)
	F_errmsg(m, int32(116585), v29+int32(96))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(496316), int32(2781), int32(90778))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L4
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L4
	} else {
		goto L261
	}
L261:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1031 + int32(4)
	F_errmsg(m, int32(705378), v29+int32(16))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(496316), int32(2788), int32(90778))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L4
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1053 + int32(4)
	F_errmsg(m, int32(705451), v29+int32(32))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(496316), int32(2799), int32(90778))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L4
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v1075 + int32(4)
	F_errmsg(m, int32(705520), v29+int32(80))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(496316), int32(2809), int32(90778))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L4
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L4
	} else {
		goto L273
	}
L273:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1097 + int32(4)
	F_errmsg(m, int32(705589), v29-int32(-64))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(496316), int32(2821), int32(90778))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L4
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+12))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1120)))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+4))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v1123 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v1122
	F_errmsg(m, int32(71628), v29+int32(160))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	F_errfinish(m, int32(496316), int32(2622), int32(120306))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L280:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v212)+32))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1145)+12))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1146)))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v1148
	F_errmsg(m, int32(711457), v29+int32(176))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L4
	} else {
		goto L282
	}
L282:
	;
	F_errfinish(m, int32(496316), int32(2627), int32(120306))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L284:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v212)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v29)+196)) = v1167 + int32(4)
	F_errmsg(m, int32(116585), v29+int32(192))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L4
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(496316), int32(2655), int32(120306))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	m.G0 = v29 + int32(352)
	return v1200
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v14 = F_SearchSysCache1(m, int32(57), l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(46291), v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_errfinish(m, int32(496316), int32(1586), int32(385222))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				F_CatalogTupleDelete(m, v11, v14+int32(4))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_sequence_close(m, v11, int32(3))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
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
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v12 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v18
	v22 = F_smgropen(m, v10+int32(8), v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v40 = v12
	goto L3
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+118)))
	if v42 != int32(116) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v22
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = v38
	goto L3
L7:
	;
	v34 = v26
	goto L9
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+72))
	v34 = v32
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+72)) = v34 + int32(1)
	goto L6
L10:
	;
	m.G0 = v10 + int32(48)
	return
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v46 <= int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	if v210 <= int32(0) {
		goto L10
	} else {
		goto L53
	}
L14:
	;
	v53 = v2
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v60 = v57 + v53<<(uint(int32(6))%32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v61 != v62 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v205 = v53 + int32(1)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v205 < v207 {
		v53 = v205
		goto L15
	} else {
		goto L52
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v64 != v65 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v67 != v68 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(229400)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(495416)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v87 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v86 | v87
	if v86&v87 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	goto L26
L24:
	;
	v111 = v86
	goto L25
L25:
	;
	v119 = int32(4122140)
	v120 = *(*int32)(unsafe.Add(mBase, _consts[276]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))+8))
	if v122 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	F_perform_spin_delay(m, v10+int32(24))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v111 = v103
	goto L25
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v104 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v103 | v104
	if v103&v104 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v139 != v140 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[276])) = v137
	goto L31
L33:
	;
	if int32(999) < v120 {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v120 < int32(11) {
		goto L31
	} else {
		goto L40
	}
L36:
	;
	v127 = int32(900)
	if v127 <= v120 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v130 = v127
	goto L39
L38:
	;
	v130 = v120
	goto L39
L39:
	;
	v137 = v130 + int32(100)
	goto L32
L40:
	;
	v137 = v120 - int32(1)
	goto L32
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v111 & int32(-4194305)
	goto L17
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v142 != v143 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v145 = int32(25165824)
	if v111&v145 != v145 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v149 != v150 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v153 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = (v152 + v153) & int32(-4194305)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v159 = int32(4431512)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[278]))
	*(*int32)(unsafe.Add(mBase, _consts[278])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v153
	v167 = v158 + v153
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v167
	v170 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerRemember(m, v170, v167, int32(1628528))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v175 = v60 + int32(48)
	v177 = F_LWLockAcquire(m, v175, int32(1))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_FlushBuffer(m, v60, v40, int32(3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_LWLockRelease(m, v175)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	F_ResourceOwnerForget(m, v185, v186+int32(1), int32(1628528))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_UnpinBufferNoOwner(m, v60)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L17
L52:
	;
	goto L16
L53:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v217 = int32(0)
	v220 = v210
	v221 = v214
	goto L54
L54:
	;
	v225 = v221 + v217<<(uint(int32(6))%32)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v226 != v227 {
		v273 = v220
		v274 = v221
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L10
L56:
	;
	v276 = v217 + int32(1)
	if v276 < v273 {
		v217 = v276
		v220 = v273
		v221 = v274
		goto L54
	} else {
		goto L66
	}
L57:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v229 != v230 {
		v273 = v220
		v274 = v221
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v232 != v233 {
		v273 = v220
		v274 = v221
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v225)+24))
	v236 = int32(25165824)
	if v235&v236 != v236 {
		v273 = v220
		v274 = v221
		goto L56
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(1079)
	v243 = int32(4508504)
	v244 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v10 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v244
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_ResourceOwnerEnlarge(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_PinLocalBuffer(m, v225, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	F_FlushLocalBuffer(m, v225, v40)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v225)+20))
	F_UnpinLocalBuffer(m, v261+int32(1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v267
	v270 = *(*int32)(unsafe.Add(mBase, _consts[742]))
	v272 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v273 = v270
	v274 = v272
	goto L56
L66:
	;
	goto L55
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
func F_NewRelationCreateToastTable(m *base.Module, l0 int32, l1 int32) {
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
			F_sequence_close(m, v4, int32(0))
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
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2 - int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		F_ResourceOwnerForget(m, v9, l0, int32(1739200))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+120)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v201
L2:
	;
	if l1 == int32(0) {
		v201 = v15
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v57 = v14 << (uint(int32(2)) % 32)
	v58 = F_palloc0(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L18
	}
L5:
	;
	v20 = F_palloc(m, v14<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v14 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v20
L9:
	;
	goto L10
L10:
	;
	v29 = v3
	goto L11
L11:
	;
	v41 = v29 << (uint(int32(2)) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15+v41)))
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v20
L13:
	;
	v46 = F_datumCopy(m, v43, int32(0), int32(-1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	v48 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v41))) = v48
	v52 = v29 + int32(1)
	if v52 != v14 {
		v29 = v52
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v48 = v46
	goto L15
L17:
	;
	goto L12
L18:
	;
	v60 = int32(0)
	v61 = base.B2i32(v14 <= v60)
	if v61 == v60 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = v3
	goto L22
L20:
	;
	goto L21
L21:
	;
	v116 = int32(4515600)
	v117 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v119
	v121 = F_palloc(m, v57)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L32
	}
L22:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1166])))
	if v79 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v101 = v68 + int32(1)
	if v101 != v14 {
		v68 = v101
		goto L22
	} else {
		goto L31
	}
L25:
	;
	if v55 == int32(2659) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v87 = base.I32_extend16_s(v68 + int32(1))
	v88 = F_get_attoptions(m, v55, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v91 = F_index_opclass_options(m, l0, v87, v88, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58+v68<<(uint(int32(2))%32)))) = v91
	if v88 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	F_pfree(m, v88)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L24
L31:
	;
	goto L23
L32:
	;
	if v61 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_pfree(m, v58)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L53
	}
L34:
	;
	v127 = int32(0)
	goto L37
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v121
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v117
	if l1 != 0 {
		v201 = v58
		goto L1
	} else {
		goto L52
	}
L37:
	;
	v139 = v127 << (uint(int32(2)) % 32)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v58+v139)))
	if v141 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v121
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v117
	if l1 != 0 {
		v201 = v58
		goto L1
	} else {
		goto L44
	}
L39:
	;
	v144 = F_datumCopy(m, v141, int32(0), int32(-1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	v146 = int32(0)
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121+v139))) = v146
	v150 = v127 + int32(1)
	if v150 != v14 {
		v127 = v150
		goto L37
	} else {
		goto L43
	}
L42:
	;
	v146 = v144
	goto L41
L43:
	;
	goto L38
L44:
	;
	v158 = int32(0)
	goto L45
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v58+v158<<(uint(int32(2))%32))))
	if v171 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L33
L47:
	;
	F_pfree(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v175 = v158 + int32(1)
	if v175 != v14 {
		v158 = v175
		goto L45
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L46
L52:
	;
	goto L33
L53:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v201 = v194
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	v2 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v478 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v478
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v21 != 0 {
		v478 = v21
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = m.G0
	v24 = v22 - int32(80)
	m.G0 = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v28 = F_SearchSysCache1(m, int32(45), v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v478 = v476
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L9
	} else {
		goto L113
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L9
	} else {
		goto L105
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L9
	} else {
		goto L102
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L9
	} else {
		goto L99
	}
L9:
	;
	return int32(0)
L10:
	;
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(21830), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v388 = m.ExcPending
	if v388 != 0 {
		goto L9
	} else {
		goto L96
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v43 = F_MemoryContextStrdup(m, v38, v40+int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v43
	goto L16
L16:
	;
	v47 = F_MemoryContextAllocZero(m, v38, int32(56))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v51 = v49 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = base.I32_extend8_s(v52)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)) = uint16(v55)
	v58 = v52 - int32(104)
	if base.Ui32(int32(10)) < base.Ui32(v58) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	if int32(1)<<(uint(v58)%32)&int32(1041) == int32(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v71 = F_SysCacheGetAttrNotNull(m, int32(45), v28, int32(6))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v75 = F_SysCacheGetAttrNotNull(m, int32(45), v28, int32(7))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	v81 = F_SysCacheGetAttr(m, int32(45), v28, int32(8), v24+int32(79))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+79)))
	if v83 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v38
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v113 = F_palloc0(m, v110<<(uint(int32(1))%32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L9
	} else {
		goto L33
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v106 = v87
	goto L23
L25:
	;
	goto L26
L26:
	;
	v88 = F_text_to_cstring(m, v81)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v90 = F_stringToNode(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v95 = F_eval_const_expressions(m, int32(0), v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_fix_opfuncids(m, v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	v99 = int32(4515600)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v38
	v103 = F_copyObjectImpl(m, v95)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v103
	v106 = v100
	goto L23
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v113
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v119 = F_palloc0(m, v116<<(uint(int32(2))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v119
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v125 = F_palloc0(m, v122<<(uint(int32(2))%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v125
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v131 = F_palloc0(m, v128*int32(28))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v131
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v137 = F_palloc0(m, v134<<(uint(int32(2))%32))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v137
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v143 = F_palloc0(m, v140<<(uint(int32(2))%32))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v143
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v149 = F_palloc0(m, v146<<(uint(int32(2))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v149
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v155 = F_palloc0(m, v152<<(uint(int32(1))%32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v155
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v159 = F_palloc0(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v159
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v163 = F_palloc0(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v163
	v166 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v169 = F_palloc0(m, v166<<(uint(int32(2))%32))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v169
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v106
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v178 = v176 << (uint(int32(1)) % 32)
	if v178 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if v181 != 0 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v179 = F__emscripten_memcpy_bulkmem(m, v175, v51+int32(36), v178)
	mBase = m.M
	goto L47
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v183 = v182
	goto L50
L49:
	;
	v183 = v2
	goto L50
L50:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	if int32(0) < v184 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v174 == int32(104) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L78
	}
L54:
	;
	v191 = int32(2)
	goto L56
L55:
	;
	v191 = int32(1)
	goto L56
L56:
	;
	v192 = int32(24)
	v201 = int32(0)
	v205 = v183
	goto L57
L57:
	;
	v214 = v201 << (uint(int32(1)) % 32)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v214+v215))))
	v220 = v201 << (uint(int32(2)) % 32)
	v221 = v71 + v192 + v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v223 = F_SearchSysCache1(m, int32(14), v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L59
	}
L58:
	;
	goto L53
L59:
	;
	if v223 == int32(0) {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v223)+16))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+22)))
	v231 = v229 + v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v227+v220))) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v234+v220))) = v236
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v231)+80))
	v239 = F_get_opfamily_proc(m, v238, v236, v236, v191)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	if v239 == int32(0) {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	F_fmgr_info_cxt(m, v239, v243+v201*int32(28), v38)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v75+v192))))
	*(*int32)(unsafe.Add(mBase, uint32(v249+v220))) = v252
	if v217 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310+v220)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	F_get_typlenbyvalalign(m, v312, v313+v214, v315+v201, v317+v201)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L9
	} else {
		goto L75
	}
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v263 = v256 + v257<<(uint(int32(4))%32) + v217*int32(100)
	v265 = v263 - int32(80)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v254+v220))) = v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v268+v220))) = v270
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v263)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v272+v220))) = v274
	v309 = v205
	goto L64
L66:
	;
	goto L67
L67:
	;
	if v205 == int32(0) {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v279 = F_exprType(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v281+v220))) = v279
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v285 = F_exprTypmod(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v287+v220))) = v285
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v291 = F_exprCollation(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v293+v220))) = v291
	v297 = v205 + int32(4)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if base.Ui32(v297) < base.Ui32(v300+v301<<(uint(int32(2))%32)) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v306 = v297
	goto L74
L73:
	;
	v306 = int32(0)
	goto L74
L74:
	;
	v309 = v306
	goto L64
L75:
	;
	F_ReleaseCatCache(m, v223)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	v324 = v201 + int32(1)
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v324 < v325 {
		v201 = v324
		v205 = v309
		goto L57
	} else {
		goto L77
	}
L77:
	;
	goto L58
L78:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v350 != v346 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v38
	m.G0 = v24 + int32(80)
	goto L4
L80:
	;
	if v350 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	if v346 != 0 {
		goto L90
	} else {
		goto L91
	}
L84:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v355 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if v354 == int32(0) {
		goto L83
	} else {
		goto L89
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+28)) = v354
	goto L85
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+20)) = v354
	goto L85
L89:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v360
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v346
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v346)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v367
	if v367 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = int32(0)
	goto L82
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+24)) = v38
	goto L95
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+20)) = v38
	goto L79
L96:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v389
	F_errmsg_internal(m, int32(46327), v24)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(499442), int32(99), int32(22590))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
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
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v403
	F_errmsg_internal(m, int32(729738), v24+int32(16))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(499442), int32(119), int32(22590))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
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
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v419
	F_errmsg_internal(m, int32(42306), v24+int32(32))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(499442), int32(202), int32(22590))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v231)+84))
	v440 = F_format_type_be(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L9
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v231 + int32(8)
	if v438 == int32(104) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v451 = int32(323688)
	goto L110
L109:
	;
	v451 = int32(410413)
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v451
	F_errmsg(m, int32(190271), v24+int32(48))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(499442), int32(221), int32(22590))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errmsg_internal(m, int32(144658), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(499442), int32(240), int32(22590))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetPartitionQual(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+131)))
	if v4 == int32(1) {
		v7 = F_generate_partition_qual(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = v7
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	if v11 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return v82
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v9, int32(2), int32(3), int32(184), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v82 = v15
	goto L1
L7:
	;
	v27 = F_table_open(m, int32(3381), int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v30 = int32(1)
	v33 = F_systable_beginscan(m, v27, int32(3379), v30, int32(0), v30, v9)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v35 = F_systable_getnext(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = v2
	v39 = v35
	goto L14
L12:
	;
	v52 = v2
	goto L13
L13:
	;
	F_systable_endscan(m, v33)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L19
	}
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43+v44)))
	v47 = F_lappend_oid(m, v38, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	v52 = v47
	goto L13
L16:
	;
	v49 = F_systable_getnext(m, v33)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v49 != 0 {
		v38 = v47
		v39 = v49
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_sequence_close(m, v27, int32(1))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_list_sort(m, v52, int32(467))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v65 = int32(4515600)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[299]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v72 = F_list_copy(m, v52)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v72
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v66
	F_list_free(m, v71)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v82 = v52
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L34
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L34
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if int32(0) < v37 {
		goto L23
	} else {
		goto L24
	}
L5:
	;
	v10 = int32(4504396)
	goto L7
L6:
	;
	v10 = int32(4505444)
	goto L7
L7:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[298]))
	if v12 == int32(0) {
		v36 = v10
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v24 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+72))
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v27&int32(1) != 0 {
		goto L2
	} else {
		goto L15
	}
L12:
	;
	v27 = int32(1)
	goto L14
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+76)))
	v27 = v26
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
	v32 = int32(4503872)
	goto L18
L17:
	;
	v32 = int32(4504920)
	goto L18
L18:
	;
	if l3 != 0 {
		v36 = v32
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
	v35 = int32(4505968)
	goto L22
L21:
	;
	v35 = int32(4506492)
	goto L22
L22:
	;
	v36 = v35
	goto L4
L23:
	;
	v45 = int32(0)
	goto L27
L24:
	;
	goto L25
L25:
	;
	v70 = v36 + v37<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = l0
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v73 + int32(1)
	return
L26:
	;
	if int32(64) <= v37 {
		goto L1
	} else {
		goto L33
	}
L27:
	;
	v52 = v36 + int32(8) + v45<<(uint(int32(3))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != l0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = l1
	return
L29:
	;
	v56 = v45 + int32(1)
	if v37 != v56 {
		v45 = v56
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
	F_errmsg_internal(m, int32(255742), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(495564), int32(348), int32(239066))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	F_errmsg_internal(m, int32(412492), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(495564), int32(351), int32(239066))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
	F_errmsg_internal(m, int32(238567), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(495564), int32(403), int32(356199))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
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
	var v116 int32
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
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
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
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
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
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
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
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v696 int32
	_ = v696
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v739 int32
	_ = v739
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1084 int32
	_ = v1084
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1254 int32
	_ = v1254
	var v1268 int64
	_ = v1268
	var v1279 int32
	_ = v1279
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(320)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	v28 = v26 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v28
	if v28 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgrclose(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v33 = v23 + int32(76)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[1154]))
	if v35 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v42
	v44 = int32(4439080)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v33
	*(*int32)(unsafe.Add(mBase, _consts[1155])) = v33
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1155]))
	v42 = v37
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = int32(4439080)
	*(*int32)(unsafe.Add(mBase, _consts[1154])) = v39
	v42 = v39
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
	F_pfree(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v59
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+119)))
	if v64|int32(32) != int32(105) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	m.G0 = v21 + int32(320)
	return
L19:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	goto L265
L20:
	;
	v1084 = int32(0)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L12
	} else {
		goto L261
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L12
	} else {
		goto L258
	}
L23:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v176 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v69 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+117)))
	if v72 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v174 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v174)
	goto L18
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v83 = F_ScanPgRelation(m, v79, base.B2i32(v79 != int32(2662)), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L12
	} else {
		goto L31
	}
L28:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1166])))
	if v76 != 0 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if v83 == int32(0) {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+16))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+22)))
	goto L34
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v94 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v92 = F__emscripten_memcpy_bulkmem(m, v87, v88+v89, int32(144))
	mBase = m.M
	goto L36
L36:
	;
	goto L33
L37:
	;
	F_pfree(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_RelationParseRelOptions(m, l0, v83)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L12
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	F_pfree(m, v83)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v104 = int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v105) < base.Ui32(int32(12000)) {
		v114 = v104
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v114 != 0 {
		goto L26
	} else {
		goto L48
	}
L45:
	;
	goto L44
L46:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	if v109 == int32(99) {
		v114 = v104
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v112 = F_isTempToastNamespace(m, v109)
	mBase = m.M
	v114 = v112
	goto L45
L48:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v117 = F_SearchSysCache1(m, int32(34), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	if v117 == int32(0) {
		goto L22
	} else {
		goto L50
	}
L50:
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
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v167 = v165
	goto L53
L52:
	;
	v167 = int32(2)
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v167
	F_ReleaseCatCache(m, v117)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	goto L26
L55:
	;
	F_RelationInitPhysicalAddr(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v205 = F_RelationBuildDesc(m, v203, int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L66
	}
L58:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1166])))
	if v182 != int32(1) {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v185)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v190 = F_ScanPgRelation(m, v187, v185, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+16))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+22)))
	goto L62
L61:
	;
	F_pfree(m, v190)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L12
	} else {
		goto L65
	}
L62:
	;
	v197 = F__emscripten_memcpy_bulkmem(m, v192, v193+v194, int32(144))
	mBase = m.M
	goto L64
L64:
	;
	goto L61
L65:
	;
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v201)
	goto L18
L66:
	;
	if v205 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[655]))
	goto L70
L68:
	;
	goto L69
L69:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v205)+52))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	if v227 != v229 {
		v696 = v2
		goto L76
	} else {
		goto L77
	}
L70:
	;
	if v210 != int32(0) {
		goto L18
	} else {
		goto L71
	}
L71:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v203
	F_errmsg_internal(m, int32(360443), v21)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(499514), int32(2689), int32(264505))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v205)+68))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v719 != 0 {
		goto L185
	} else {
		goto L186
	}
L76:
	;
	v717 = v696
	goto L75
L77:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v231 != v232 {
		v696 = v2
		goto L76
	} else {
		goto L78
	}
L78:
	;
	if v227 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	if v397 != 0 {
		goto L115
	} else {
		goto L116
	}
L80:
	;
	v236 = int32(31)
	v241 = v227 << (uint(int32(4)) % 32)
	v243 = int32(20)
	v253 = v2
	goto L81
L81:
	;
	v266 = int32(0)
	v268 = v253 * int32(100)
	v269 = v241 + v226 + v243 + v268
	v270 = int32(4)
	v271 = v269 + v270
	v272 = v268 + (v228 + v241 + v243)
	v274 = v272 + v270
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v278 == v266 {
		v297 = v277
		v298 = v278
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v717 = int32(0)
	goto L75
L83:
	;
	if v298-v297 != 0 {
		v717 = v266
		goto L75
	} else {
		goto L91
	}
L84:
	;
	goto L83
L85:
	;
	if v277 != v278 {
		v297 = v277
		v298 = v278
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v282 = v271
	v283 = v274
	goto L87
L87:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v286
		v298 = v287
		goto L84
	} else {
		goto L89
	}
L88:
	;
	v297 = v286
	v298 = v287
	goto L84
L89:
	;
	v290 = int32(1)
	if v286 == v287 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v269)+68))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v272)+68))
	if v301 != v302 {
		v717 = int32(0)
		goto L75
	} else {
		goto L92
	}
L92:
	;
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+72)))
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272)+72)))
	if v305 != v306 {
		v717 = int32(0)
		goto L75
	} else {
		goto L93
	}
L93:
	;
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+80)))
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272)+80)))
	if v309 != v310 {
		v717 = int32(0)
		goto L75
	} else {
		goto L94
	}
L94:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v269)+76))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v272)+76))
	if v313 != v314 {
		v717 = int32(0)
		goto L75
	} else {
		goto L95
	}
L95:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+82)))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+82)))
	if v317 != v318 {
		v717 = int32(0)
		goto L75
	} else {
		goto L96
	}
L96:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+83)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+83)))
	if v321 != v322 {
		v717 = int32(0)
		goto L75
	} else {
		goto L97
	}
L97:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+84)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+84)))
	if v325 != v326 {
		v717 = int32(0)
		goto L75
	} else {
		goto L98
	}
L98:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+85)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+85)))
	if v329 != v330 {
		v717 = int32(0)
		goto L75
	} else {
		goto L99
	}
L99:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+86)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+86)))
	if v333 != v334 {
		v717 = int32(0)
		goto L75
	} else {
		goto L100
	}
L100:
	;
	if v333 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+87)))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+87)))
	if v348 != v349 {
		v717 = int32(0)
		goto L75
	} else {
		goto L104
	}
L102:
	;
	v339 = v253 << (uint(int32(4)) % 32)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v236+v339))))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339+(v228+v236)))))
	if v341 == v343 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v717 = int32(0)
	goto L75
L104:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+89)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+89)))
	if v352 != v353 {
		v717 = int32(0)
		goto L75
	} else {
		goto L105
	}
L105:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+90)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+90)))
	if v356 != v357 {
		v717 = int32(0)
		goto L75
	} else {
		goto L106
	}
L106:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+91)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+91)))
	if v360 != v361 {
		v717 = int32(0)
		goto L75
	} else {
		goto L107
	}
L107:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+92)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+92)))
	if v364 != v365 {
		v717 = int32(0)
		goto L75
	} else {
		goto L108
	}
L108:
	;
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+94)))
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272)+94)))
	if v368 != v369 {
		v717 = int32(0)
		goto L75
	} else {
		goto L109
	}
L109:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v269)+96))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v272)+96))
	if v371 == v372 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v375 = v253 + int32(1)
	if v375 == v227 {
		goto L79
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L82
L113:
	;
	v253 = v375
	goto L81
L114:
	;
	v696 = int32(1)
	goto L76
L115:
	;
	if v396 == int32(0) {
		v696 = v2
		goto L76
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v396 != 0 {
		v696 = v2
		goto L76
	} else {
		goto L182
	}
L118:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+16)))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+16)))
	if v400 != v401 {
		v696 = v2
		goto L76
	} else {
		goto L119
	}
L119:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+17)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+17)))
	if v403 != v404 {
		v696 = v2
		goto L76
	} else {
		goto L120
	}
L120:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+18)))
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+18)))
	if v406 != v407 {
		v696 = v2
		goto L76
	} else {
		goto L121
	}
L121:
	;
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+12)))
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396)+12)))
	if v409 != v410 {
		v696 = v2
		goto L76
	} else {
		goto L122
	}
L122:
	;
	if v409 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	if v495 != 0 {
		goto L141
	} else {
		goto L142
	}
L124:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v396)))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
	v421 = int32(0)
	goto L125
L125:
	;
	v437 = v421 << (uint(int32(3)) % 32)
	v438 = v415 + v437
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v438))))
	v440 = v437 + v414
	v441 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v440))))
	if v439 != v441 {
		v717 = int32(0)
		goto L75
	} else {
		goto L127
	}
L126:
	;
	v717 = int32(0)
	goto L75
L127:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v443))))
	if v448 == int32(0) {
		v467 = v447
		v468 = v448
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if v468-v467 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	goto L128
L130:
	;
	if v447 != v448 {
		v467 = v447
		v468 = v448
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v452 = v443
	v453 = v444
	goto L132
L132:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+1)))
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452)+1)))
	if v457 == int32(0) {
		v467 = v456
		v468 = v457
		goto L129
	} else {
		goto L134
	}
L133:
	;
	v467 = v456
	v468 = v457
	goto L129
L134:
	;
	v460 = int32(1)
	if v456 == v457 {
		v452 = v452 + v460
		v453 = v453 + v460
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v473 = v421 + int32(1)
	if v473 == v409 {
		goto L123
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	goto L126
L139:
	;
	v421 = v473
	goto L125
L140:
	;
	v566 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+14)))
	v567 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396)+14)))
	if v566 != v567 {
		v696 = v2
		goto L76
	} else {
		goto L156
	}
L141:
	;
	if v494 == int32(0) {
		v696 = v2
		goto L76
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	if v494 != 0 {
		v696 = v2
		goto L76
	} else {
		goto L155
	}
L144:
	;
	if v227 <= int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	v507 = int32(0)
	v510 = v227
	goto L146
L146:
	;
	v523 = v507 << (uint(int32(3)) % 32)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v397)+8))
	v525 = v523 + v524
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v396)+8))
	v528 = v527 + v523
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528))))
	if v526 != v529 {
		v717 = int32(0)
		goto L75
	} else {
		goto L148
	}
L147:
	;
	goto L140
L148:
	;
	if v526 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	v536 = v226 + int32(20) + v507<<(uint(int32(4))%32)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v536)+6)))
	v538 = int32(*(*int16)(unsafe.Add(mBase, uint32(v536)+4)))
	v539 = F_datumIsEqual(m, v532, v533, v537, v538)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L12
	} else {
		goto L152
	}
L150:
	;
	v544 = v510
	goto L151
L151:
	;
	v546 = v507 + int32(1)
	if v546 < v544 {
		v507 = v546
		v510 = v544
		goto L146
	} else {
		goto L154
	}
L152:
	;
	if v539 == int32(0) {
		v717 = int32(0)
		goto L75
	} else {
		goto L153
	}
L153:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v544 = v543
	goto L151
L154:
	;
	goto L147
L155:
	;
	goto L140
L156:
	;
	if v566 == int32(0) {
		goto L114
	} else {
		goto L157
	}
L157:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	v580 = int32(0)
	goto L158
L158:
	;
	v593 = v580 * int32(12)
	v594 = v572 + v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	v596 = v593 + v571
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v596)))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	if v601 == int32(0) {
		v620 = v600
		v621 = v601
		goto L161
	} else {
		goto L162
	}
L159:
	;
	goto L114
L160:
	;
	if v621-v620 != 0 {
		v696 = v2
		goto L76
	} else {
		goto L168
	}
L161:
	;
	goto L160
L162:
	;
	if v600 != v601 {
		v620 = v600
		v621 = v601
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v605 = v595
	v606 = v597
	goto L164
L164:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v606)+1)))
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605)+1)))
	if v610 == int32(0) {
		v620 = v609
		v621 = v610
		goto L161
	} else {
		goto L166
	}
L165:
	;
	v620 = v609
	v621 = v610
	goto L161
L166:
	;
	v613 = int32(1)
	if v609 == v610 {
		v605 = v605 + v613
		v606 = v606 + v613
		goto L164
	} else {
		goto L167
	}
L167:
	;
	goto L165
L168:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v596)+4))
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v623))))
	if v628 == int32(0) {
		v647 = v627
		v648 = v628
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v648-v647 != 0 {
		v696 = v2
		goto L76
	} else {
		goto L177
	}
L170:
	;
	goto L169
L171:
	;
	if v627 != v628 {
		v647 = v627
		v648 = v628
		goto L170
	} else {
		goto L172
	}
L172:
	;
	v632 = v623
	v633 = v624
	goto L173
L173:
	;
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+1)))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+1)))
	if v637 == int32(0) {
		v647 = v636
		v648 = v637
		goto L170
	} else {
		goto L175
	}
L174:
	;
	v647 = v636
	v648 = v637
	goto L170
L175:
	;
	v640 = int32(1)
	if v636 == v637 {
		v632 = v632 + v640
		v633 = v633 + v640
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+8)))
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+8)))
	if v650 != v651 {
		v696 = v2
		goto L76
	} else {
		goto L178
	}
L178:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+9)))
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+9)))
	if v653 != v654 {
		v696 = v2
		goto L76
	} else {
		goto L179
	}
L179:
	;
	v656 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594)+10)))
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596)+10)))
	if v656 != v657 {
		v696 = v2
		goto L76
	} else {
		goto L180
	}
L180:
	;
	v660 = v580 + int32(1)
	if v566 != v660 {
		v580 = v660
		goto L158
	} else {
		goto L181
	}
L181:
	;
	goto L159
L182:
	;
	goto L114
L183:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v205)+80))
	if v818|v819 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L184:
	;
	v817 = int32(1)
	goto L183
L185:
	;
	if v718 == int32(0) {
		v817 = v2
		goto L183
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	if v718 != 0 {
		v817 = v2
		goto L183
	} else {
		goto L202
	}
L188:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v719)))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v718)))
	if v722 != v723 {
		v817 = v2
		goto L183
	} else {
		goto L189
	}
L189:
	;
	if v722 <= int32(0) {
		goto L184
	} else {
		goto L190
	}
L190:
	;
	v739 = v2
	goto L191
L191:
	;
	v746 = v739 << (uint(int32(2)) % 32)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v746+v747)))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v751+v746)))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	if v750 != v754 {
		v817 = v2
		goto L183
	} else {
		goto L193
	}
L192:
	;
	goto L184
L193:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v753)+4))
	if v756 != v757 {
		v817 = v2
		goto L183
	} else {
		goto L194
	}
L194:
	;
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+16)))
	v760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+16)))
	if v759 != v760 {
		v817 = v2
		goto L183
	} else {
		goto L195
	}
L195:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+17)))
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+17)))
	if v762 != v763 {
		v817 = v2
		goto L183
	} else {
		goto L196
	}
L196:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v749)+8))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v753)+8))
	v767 = F_equal(m, v765, v766)
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L12
	} else {
		goto L197
	}
L197:
	;
	if v767 == int32(0) {
		v817 = v2
		goto L183
	} else {
		goto L198
	}
L198:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v749)+12))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v773 = F_equal(m, v771, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L12
	} else {
		goto L199
	}
L199:
	;
	if v773 == int32(0) {
		v817 = v2
		goto L183
	} else {
		goto L200
	}
L200:
	;
	v778 = v739 + int32(1)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v719)))
	if v778 < v779 {
		v739 = v778
		goto L191
	} else {
		goto L201
	}
L201:
	;
	goto L192
L202:
	;
	goto L184
L203:
	;
	v1084 = int32(1)
	goto L19
L204:
	;
	goto L205
L205:
	;
	v824 = int32(0)
	if base.B2i32(v818 != v824)^base.B2i32(v819 != v824) != 0 {
		v1084 = v824
		goto L19
	} else {
		goto L206
	}
L206:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	if v831 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	v833 = v832
	goto L209
L208:
	;
	v833 = int32(0)
	goto L209
L209:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v819)+4))
	if v834 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	v837 = v835
	goto L212
L211:
	;
	v837 = int32(0)
	goto L212
L212:
	;
	if v837 != v833 {
		goto L20
	} else {
		goto L213
	}
L213:
	;
	v853 = v2
	goto L214
L214:
	;
	v857 = int32(0)
	if v831 == v857 {
		v867 = v857
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v868 = int32(1)
	if v834 == int32(0) {
		v1084 = v868
		goto L19
	} else {
		goto L219
	}
L217:
	;
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	if v861 <= v853 {
		v867 = int32(0)
		goto L216
	} else {
		goto L218
	}
L218:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v831)+12))
	v867 = v863 + v853<<(uint(int32(2))%32)
	goto L216
L219:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v834)+4))
	if v871 <= v853 {
		v1084 = v868
		goto L19
	} else {
		goto L220
	}
L220:
	;
	if v867 == int32(0) {
		v1084 = v868
		goto L19
	} else {
		goto L221
	}
L221:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v834)+12))
	v878 = v875 + v853<<(uint(int32(2))%32)
	if v878 == int32(0) {
		v1084 = v868
		goto L19
	} else {
		goto L222
	}
L222:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v878)))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	if v882 != 0 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v853 = v853 + int32(1)
	goto L214
L224:
	;
	v883 = int32(0)
	if v881 == v883 {
		v1084 = v883
		goto L19
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	if v881 != 0 {
		goto L20
	} else {
		goto L257
	}
L227:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+4)))
	v887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+4)))
	if v886 != v887 {
		v1084 = v883
		goto L19
	} else {
		goto L228
	}
L228:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v882)+24)))
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881)+24)))
	if v889 != v890 {
		v1084 = v883
		goto L19
	} else {
		goto L229
	}
L229:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v881)))
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v893))))
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v892))))
	if v897 == int32(0) {
		v916 = v896
		v917 = v897
		goto L231
	} else {
		goto L232
	}
L230:
	;
	if v917-v916 != 0 {
		v1084 = v883
		goto L19
	} else {
		goto L238
	}
L231:
	;
	goto L230
L232:
	;
	if v896 != v897 {
		v916 = v896
		v917 = v897
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v901 = v892
	v902 = v893
	goto L234
L234:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v902)+1)))
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v901)+1)))
	if v906 == int32(0) {
		v916 = v905
		v917 = v906
		goto L231
	} else {
		goto L236
	}
L235:
	;
	v916 = v905
	v917 = v906
	goto L231
L236:
	;
	v909 = int32(1)
	if v905 == v906 {
		v901 = v901 + v909
		v902 = v902 + v909
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v882)+8))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v919)+16))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v881)+8))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+16))
	if v920 != v922 {
		v1084 = v883
		goto L19
	} else {
		goto L239
	}
L239:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v919)+8))
	if v924 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v919)+4))
	v934 = (v927<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L242
L241:
	;
	v934 = v924
	goto L242
L242:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v921)+8))
	if v935 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v921)+4))
	v945 = (v938<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L245
L244:
	;
	v945 = v935
	goto L245
L245:
	;
	if int32(0) < v920 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v952 = int32(0)
	goto L249
L247:
	;
	goto L248
L248:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v882)+16))
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v881)+16))
	v999 = F_equal(m, v997, v998)
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L12
	} else {
		goto L253
	}
L249:
	;
	v970 = v952 << (uint(int32(2)) % 32)
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v919+v934+v970)))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v970+(v921+v945))))
	if v972 != v974 {
		v1084 = v883
		goto L19
	} else {
		goto L251
	}
L250:
	;
	goto L248
L251:
	;
	v977 = v952 + int32(1)
	if v977 != v920 {
		v952 = v977
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	if v999 == int32(0) {
		v1084 = v883
		goto L19
	} else {
		goto L254
	}
L254:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v882)+20))
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v881)+20))
	v1005 = F_equal(m, v1003, v1004)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L12
	} else {
		goto L255
	}
L255:
	;
	if v1005 == int32(0) {
		v1084 = v883
		goto L19
	} else {
		goto L256
	}
L256:
	;
	goto L223
L257:
	;
	goto L223
L258:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v1033
	F_errmsg_internal(m, int32(40143), v21+int32(32))
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L12
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(499514), int32(2343), int32(242256))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L12
	} else {
		goto L260
	}
L260:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L261:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1049
	F_errmsg_internal(m, int32(40100), v21+int32(16))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(499514), int32(2314), int32(242256))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L12
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	goto L269
L265:
	;
	v1102 = F__emscripten_memcpy_bulkmem(m, v21+int32(44), v205, int32(276))
	mBase = m.M
	goto L267
L267:
	;
	goto L264
L268:
	;
	goto L273
L269:
	;
	v1105 = F__emscripten_memcpy_bulkmem(m, v205, l0, int32(276))
	mBase = m.M
	goto L271
L271:
	;
	goto L268
L272:
	;
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+12))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+12)) = v1113
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+12)) = v1112
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+16))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+16)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+16)) = v1116
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+32))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+32)) = v1121
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+32)) = v1120
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+36))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+36)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+36)) = v1124
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+40))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+40)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+40)) = v1128
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+44))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+44)) = v1133
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+44)) = v1132
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+48))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+48)) = v1137
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+48)) = v1136
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+48))
	goto L277
L273:
	;
	v1110 = F__emscripten_memcpy_bulkmem(m, l0, v21+int32(44), int32(276))
	mBase = m.M
	goto L275
L275:
	;
	goto L272
L276:
	;
	if v717 != 0 {
		goto L280
	} else {
		goto L281
	}
L277:
	;
	v1142 = F__emscripten_memcpy_bulkmem(m, v1136, v1140, int32(144))
	mBase = m.M
	goto L279
L279:
	;
	goto L276
L280:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+52))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+52)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+52)) = v1144
	goto L282
L281:
	;
	goto L282
L282:
	;
	if v817 != 0 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+68))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+68)) = v1150
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+68)) = v1149
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+72))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+72)) = v1154
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+72)) = v1153
	goto L285
L284:
	;
	goto L285
L285:
	;
	if v1084 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+80))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+80)) = v1159
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+80)) = v1158
	goto L288
L287:
	;
	goto L288
L288:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+264))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+264)) = v1164
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+264)) = v1163
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+272))
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+272)) = v1168
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+272)) = v1167
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105)+268)))
	v1172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110)+268)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1105)+268)) = uint8(v1172)
	*(*uint8)(unsafe.Add(mBase, uint32(v1110)+268)) = uint8(v1171)
	if v1098 != 0 {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+92))
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+92)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+92)) = v1175
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+96))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+96)) = v1180
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+96)) = v1179
	goto L291
L290:
	;
	goto L291
L291:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+104))
	if v1184 == int32(0) {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	F_RelationDestroyRelation(m, v1105, v717^int32(1))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L12
	} else {
		goto L339
	}
L293:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+112))
	if v1187 == int32(0) {
		goto L292
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	v1190 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+116)) = v1190
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+108)) = v1190
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+100)) = v1190
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+104))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+104))
	if v1197 != 0 {
		goto L298
	} else {
		goto L299
	}
L296:
	;
	goto L295
L297:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1105)+112))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+112))
	if v1233 != 0 {
		goto L319
	} else {
		goto L320
	}
L298:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+16))
	if v1201 != v1197 {
		goto L302
	} else {
		goto L303
	}
L299:
	;
	goto L300
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+104)) = v1196
	goto L297
L301:
	;
	goto L297
L302:
	;
	if v1201 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L303:
	;
	goto L304
L304:
	;
	goto L301
L305:
	;
	if v1197 != 0 {
		goto L312
	} else {
		goto L313
	}
L306:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+28))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+24))
	if v1206 != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	if v1205 == int32(0) {
		goto L305
	} else {
		goto L311
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+28)) = v1205
	goto L307
L309:
	;
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1201)+20)) = v1205
	goto L307
L311:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1205)+24)) = v1211
	goto L305
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+16)) = v1197
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1197)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+28)) = v1218
	if v1218 != 0 {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	goto L314
L314:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1196)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1196)+16)) = int32(0)
	goto L304
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+24)) = v1196
	goto L317
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+20)) = v1196
	goto L301
L318:
	;
	v1268 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+100)) = v1268
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+116)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+108)) = v1268
	goto L292
L319:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+16))
	if v1237 != v1233 {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1110)+112)) = v1232
	goto L318
L322:
	;
	goto L318
L323:
	;
	if v1237 == int32(0) {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	goto L325
L325:
	;
	goto L322
L326:
	;
	if v1233 != 0 {
		goto L333
	} else {
		goto L334
	}
L327:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+28))
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+24))
	if v1242 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1241 == int32(0) {
		goto L326
	} else {
		goto L332
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+28)) = v1241
	goto L328
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1237)+20)) = v1241
	goto L328
L332:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1241)+24)) = v1247
	goto L326
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+16)) = v1233
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+28)) = v1254
	if v1254 != 0 {
		goto L336
	} else {
		goto L337
	}
L334:
	;
	goto L335
L335:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1232)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1232)+16)) = int32(0)
	goto L325
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1254)+24)) = v1232
	goto L338
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1233)+20)) = v1232
	goto L322
L339:
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
	var v41 int32
	_ = v41
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1168]))
	v10 = v8 - int32(1)
	if v10 < v2 {
		v41 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v41
L2:
	;
	v14 = v2
	v15 = v10
	goto L3
L3:
	;
	v20 = int32(2)
	v21 = base.I32_div_s(v15-v14, v20)
	v22 = v21 + v14
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(v20)%32))+uint32(_consts[1169])))
	v28 = base.B2i32(v27 == l0)
	if v27 == l0 {
		v41 = v28
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v41 = v28
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
	v32 = v14
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
	v35 = v15
	goto L11
L10:
	;
	v35 = v22 - int32(1)
	goto L11
L11:
	;
	if v32 <= v35 {
		v14 = v32
		v15 = v35
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
	var v16 int32
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
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v18 = F_SearchSysCacheCopy(m, int32(57), v16, int32(0))
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
							F_sequence_close(m, v13, int32(3))
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
							F_sequence_close(m, v13, int32(3))
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
					F_errmsg_internal(m, int32(46291), v9)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(496316), int32(3160), int32(154641))
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v15 = F_LockRelease(m, v6, l1, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l1 == v3 {
		v12 = int32(0)
		v15 = F_SearchSysCacheExists(m, int32(57), l0, v12, v12, v12)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v68 = v3
				m.G0 = v7 + int32(16)
				return v68
			} else {
				v31 = F_RelationIdGetRelation(m, l0)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					if v31 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(56067), v7)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496628), int32(115), int32(281814))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+118)))
						if v36 == int32(116) {
							v39 = int32(4410724)
							v41 = *(*int32)(unsafe.Add(mBase, _consts[22]))
							*(*int32)(unsafe.Add(mBase, _consts[22])) = v41 | int32(1)
						} else {
						}
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
						switch v47 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v55 = int32(*(*uint8)(unsafe.Add(mBase, _consts[23])))
							if v55 == int32(0) {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
								if v58 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = int32(0)
								} else {
								}
								v61 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v61
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v61)
							} else {
								v65 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v65)
							}
						default:
							v50 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v50
							*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v50)
						}
						v68 = v31
						m.G0 = v7 + int32(16)
						return v68
					}
				}
			}
		}
	} else {
		F_LockRelationOid(m, l0, l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = int32(0)
			v27 = F_SearchSysCacheExists(m, int32(57), l0, v24, v24, v24)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				if v27 != 0 {
					v31 = F_RelationIdGetRelation(m, l0)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						if v31 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
								F_errmsg_internal(m, int32(56067), v7)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496628), int32(115), int32(281814))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+118)))
							if v36 == int32(116) {
								v39 = int32(4410724)
								v41 = *(*int32)(unsafe.Add(mBase, _consts[22]))
								*(*int32)(unsafe.Add(mBase, _consts[22])) = v41 | int32(1)
							} else {
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
							switch v47 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v55 = int32(*(*uint8)(unsafe.Add(mBase, _consts[23])))
								if v55 == int32(0) {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
									if v58 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v58)+128)) = int32(0)
									} else {
									}
									v61 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v61
									*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v61)
								} else {
									v65 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v65)
								}
							default:
								v50 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v50
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v50)
							}
							v68 = v31
							m.G0 = v7 + int32(16)
							return v68
						}
					}
				} else {
					F_UnlockRelationOid(m, l0, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v68 = v3
						m.G0 = v7 + int32(16)
						return v68
					}
				}
			}
		}
	}
}
