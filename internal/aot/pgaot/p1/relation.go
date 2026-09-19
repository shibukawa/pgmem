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
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
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
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
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
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
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
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v966 int32
	_ = v966
	var v971 int32
	_ = v971
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1070 int32
	_ = v1070
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1092 int32
	_ = v1092
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1178 int32
	_ = v1178
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
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
		v162 = v8
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l2 == int32(0) {
		v1167 = v162
		v1170 = v34
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v56 <= int32(0) {
		v162 = v8
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
	v162 = v141
	goto L8
L13:
	;
	v143 = v68 + int32(1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v143 < v144 {
		v68 = v143
		v75 = v141
		goto L11
	} else {
		goto L23
	}
L14:
	;
	if v107 == int32(0) {
		v141 = v75
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
		v141 = v75
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
	v136 = F_lappend(m, v75, v120)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v141 = v136
	goto L13
L23:
	;
	goto L12
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L4
	} else {
		goto L282
	}
L25:
	;
	F_SetRelationNumChecks(m, l0, v1170)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L4
	} else {
		goto L281
	}
L26:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v174 <= int32(0) {
		v1167 = v162
		v1170 = v34
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v195 = v8
	v197 = v162
	v200 = v34
	v202 = v8
	v203 = v8
	goto L28
L28:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v202<<(uint(int32(2))%32))))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	switch v212 - int32(1) {
	case 0:
		goto L39
	default:
		v1135 = v195
		v1137 = v197
		v1140 = v200
		v1143 = v203
		goto L30
	case 4:
		goto L40
	}
L29:
	;
	v1167 = v1137
	v1170 = v1140
	goto L25
L30:
	;
	v1148 = v202 + int32(1)
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1148 < v1149 {
		v195 = v1135
		v197 = v1137
		v200 = v1140
		v202 = v1148
		v203 = v1143
		goto L28
	} else {
		goto L280
	}
L31:
	;
	if v749&int32(1) != 0 {
		goto L271
	} else {
		goto L272
	}
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v921)+103)) = uint8(v1098)
	goto L31
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L4
	} else {
		goto L267
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L4
	} else {
		goto L263
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L4
	} else {
		goto L259
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L4
	} else {
		goto L255
	}
L37:
	;
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+14)))
	v984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	v986 = F_StoreRelCheck(m, l0, v966, v252, v983, v984, v5, v53, v985, l5)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L4
	} else {
		goto L252
	}
L38:
	;
	v747 = F_lappend(m, v195, v255)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L178
	}
L39:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	v377 = F_get_attnum(m, v372, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L92
	}
L40:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	if v215 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if v255 != 0 {
		goto L57
	} else {
		goto L58
	}
L42:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v218 = F_transformExpr(m, v36, v215, int32(28))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	v250 = F_stringToNode(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L56
	}
L45:
	;
	v221 = F_coerce_to_boolean(m, v36, v218, int32(_a_F_AddRelationNewConstraints_0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_assign_expr_collations(m, v36, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v225 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v226 == int32(1) {
		v252 = v221
		goto L41
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+144)) = v216 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_2), v29+int32(144))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(3433), int32(_a_F_AddRelationNewConstraints_4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v252 = v250
	goto L41
L57:
	;
	if v195 == int32(0) {
		goto L38
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v344 = int32(0)
	v346 = F_pull_var_clause(m, v252, v344)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L83
	}
L60:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v258 <= int32(0) {
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v261 = int32(0)
	if v261 < v258 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v264 = v258
	goto L64
L63:
	;
	v264 = v261
	goto L64
L64:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v273 = int32(0)
	goto L65
L65:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v265+v273<<(uint(int32(2))%32))))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if base.B2i32(v299 == int32(0))|base.B2i32(v299 != v302) != 0 {
		v320 = v299
		v321 = v302
		goto L68
	} else {
		goto L69
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L78
	}
L67:
	;
	if v320-v321 != 0 {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	goto L67
L69:
	;
	v305 = v296
	v306 = v255
	goto L70
L70:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+1)))
	if v310 == int32(0) {
		v320 = v310
		v321 = v309
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v320 = v310
	v321 = v309
	goto L68
L72:
	;
	v313 = int32(1)
	if v310 == v309 {
		v305 = v305 + v313
		v306 = v306 + v313
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v324 = v273 + int32(1)
	if v264 != v324 {
		v273 = v324
		goto L65
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L66
L77:
	;
	goto L38
L78:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_5))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v255
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_6), v29+int32(128))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2526), int32(_a_F_AddRelationNewConstraints_7))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363)+68))
	v368 = F_ChooseConstraintName(m, v363+int32(4), v362, int32(_a_F_AddRelationNewConstraints_8), v367, v195)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L4
	} else {
		goto L88
	}
L83:
	;
	v348 = F_list_union(m, v346)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	if v348 == int32(0) {
		v362 = v344
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v352 != int32(1) {
		v362 = v344
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v348)+12))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	v358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v357)+8)))
	v360 = F_get_attname(m, v355, v358, int32(1))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v362 = v360
	goto L82
L88:
	;
	v370 = F_lappend(m, v195, v368)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	v966 = v368
	v971 = v370
	goto L37
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L174
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L170
	}
L92:
	;
	if v377 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v377 < int32(0) {
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L166
	}
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+15)))
	v385 = m.G0
	v387 = v385 - int32(96)
	m.G0 = v387
	v389 = F_findNotNullConstraintAttnum(m, v381, v377)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L4
	} else {
		goto L101
	}
L97:
	;
	if v389 != int32(0) {
		v1135 = v195
		v1137 = v197
		v1140 = v200
		v1143 = v203
		goto L30
	} else {
		goto L154
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L147
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L4
	} else {
		goto L141
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L135
	}
L101:
	;
	if v389 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v393 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v387 + int32(96)
	goto L97
L105:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389)+16))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+22)))
	v397 = v395 + v396
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+106)))
	if v383 != v398 {
		goto L100
	} else {
		goto L106
	}
L106:
	;
	if v384 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+76)))
	if v402 == int32(0) {
		goto L99
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v405 = int32(0)
	if base.B2i32(v5 == v405)|base.B2i32(v382 == v405) == v405 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v413 = v397 + int32(4)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382))))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	if base.B2i32(v416 == int32(0))|base.B2i32(v416 != v419) != 0 {
		v437 = v416
		v438 = v419
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
	if v437-v438 != 0 {
		goto L98
	} else {
		goto L121
	}
L115:
	;
	goto L114
L116:
	;
	v422 = v382
	v423 = v413
	goto L117
L117:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+1)))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+1)))
	if v427 == int32(0) {
		v437 = v427
		v438 = v426
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v437 = v427
	v438 = v426
	goto L115
L119:
	;
	v430 = int32(1)
	if v427 == v426 {
		v422 = v422 + v430
		v423 = v423 + v430
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
	F_relation_close(m, v393, int32(3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L134
	}
L123:
	;
	F_CatalogTupleUpdate(m, v393, v389+int32(4), v389)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L133
	}
L124:
	;
	v443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v397)+104)))
	v445 = v443 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+104)) = uint16(v445)
	if base.I32_extend16_s(v445) == v445 {
		goto L123
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397)+103)))
	if v465 != 0 {
		goto L122
	} else {
		goto L132
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L128
	}
L128:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_9), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_10), int32(803), int32(_a_F_AddRelationNewConstraints_11))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
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
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+103)) = uint8(v466)
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
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v492 = F_get_rel_name(m, v381)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L4
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+84)) = v492
	*(*int32)(unsafe.Add(mBase, uint32(v387)+80)) = v397 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_12), v387+int32(80))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+64)) = int32(_a_F_AddRelationNewConstraints_13)
	F_errhint(m, int32(_a_F_AddRelationNewConstraints_14), v387-int32(-64))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_10), int32(767), int32(_a_F_AddRelationNewConstraints_11))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
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
	v521 = m.ExcPending
	if v521 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v522 = F_get_rel_name(m, v381)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+52)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v387)+48)) = v397 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_15), v387+int32(48))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+32)) = int32(_a_F_AddRelationNewConstraints_16)
	F_errhint(m, int32(_a_F_AddRelationNewConstraints_17), v387+int32(32))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_10), int32(779), int32(_a_F_AddRelationNewConstraints_11))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
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
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	v553 = F_get_attname(m, v381, v377, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v555 = F_get_rel_name(m, v381)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+24)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v387)+20)) = v553
	*(*int32)(unsafe.Add(mBase, uint32(v387)+16)) = v382
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_18), v387+int32(16))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v413
	F_errdetail(m, int32(_a_F_AddRelationNewConstraints_19), v387)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_10), int32(795), int32(_a_F_AddRelationNewConstraints_11))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
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
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if v574 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v593 = F_lappend(m, v203, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L162
	}
L156:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v577 = F_ConstraintNameIsUsed(m, int32(0), v576, v574)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v583)+12))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v584)))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v585)+4))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v580)+68))
	v589 = F_ChooseConstraintName(m, v580+int32(4), v586, int32(_a_F_AddRelationNewConstraints_20), v588, v203)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L4
	} else {
		goto L161
	}
L159:
	;
	if v577 != 0 {
		goto L90
	} else {
		goto L160
	}
L160:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	v592 = v579
	goto L155
L161:
	;
	v592 = v589
	goto L155
L162:
	;
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+208)) = uint16(v377)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v598)+68))
	v601 = int32(0)
	v603 = int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v618 = int32(32)
	v628 = F_CreateConstraintEntry(m, v592, v599, int32(110), v601, v601, v603, v596, v601, v605, v29+int32(208), v603, v603, v601, v601, v601, v601, v601, v601, v601, v601, v618, v618, v601, v601, v618, v601, v601, v601, v5, v53, v595, v601, v601)
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v631 = F_palloc(m, int32(28))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v633 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v631)+20)) = uint8(v633)
	*(*int32)(unsafe.Add(mBase, uint32(v631)+16)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+12)) = uint16(v377)
	*(*int32)(unsafe.Add(mBase, uint32(v631)+8)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v631)+4)) = v628
	*(*int32)(unsafe.Add(mBase, uint32(v631))) = v633
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v631)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v631)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v631)+21)) = uint8(v642)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v631)+26)) = uint8(v646)
	v648 = F_lappend(m, v197, v631)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v1135 = v195
	v1137 = v648
	v1140 = v200
	v1143 = v593
	goto L30
L166:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+12))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+164)) = v661 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v660
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_21), v29+int32(160))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2622), int32(_a_F_AddRelationNewConstraints_7))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v211)+32))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)+12))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v684)))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v686
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_22), v29+int32(176))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2627), int32(_a_F_AddRelationNewConstraints_7))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
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
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_5))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L4
	} else {
		goto L175
	}
L175:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v706
	*(*int32)(unsafe.Add(mBase, uint32(v29)+196)) = v705 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_23), v29+int32(192))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2655), int32(_a_F_AddRelationNewConstraints_7))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
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
	v749 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+14)))
	v754 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	v757 = v29 + int32(208)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v757, int32(9), int32(3), int32(184), v761)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	F_ScanKeyInit(m, v29+int32(256), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_ScanKeyInit(m, v29+int32(304), int32(2), int32(3), int32(62), v255)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	v779 = F_systable_beginscan(m, v754, int32(2665), int32(1), int32(0), int32(3), v757)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v781 = F_systable_getnext(m, v779)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	if v781 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v781)+16))
	v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+22)))
	v786 = v784 + v785
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+72)))
	if v787 == int32(99) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	goto L187
L187:
	;
	F_systable_endscan(m, v779)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L4
	} else {
		goto L250
	}
L188:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v754)+52))
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+20)))
	if v791&int32(1) == int32(0) {
		goto L193
	} else {
		goto L194
	}
L189:
	;
	v860 = int32(0)
	goto L190
L190:
	;
	if v5 == int32(0) {
		v869 = l3
		goto L217
	} else {
		goto L218
	}
L191:
	;
	v853 = F_text_to_cstring(m, v852)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L4
	} else {
		goto L214
	}
L192:
	;
	v849 = int32(*(*int8)(unsafe.Add(mBase, uint32(v799))))
	v852 = v849
	goto L191
L193:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v790)+452))
	if int32(0) <= v796 {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	goto L195
L195:
	;
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v784)+26)))
	if v825&int32(8) != 0 {
		goto L207
	} else {
		goto L208
	}
L196:
	;
	v799 = v796 + v786
	v800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790)+458)))
	if v800 != int32(1) {
		v852 = v799
		goto L191
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v823 = F_nocachegetattr(m, v781, int32(28), v790)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L206
	}
L199:
	;
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v790)+456)))
	switch v803 - int32(1) {
	case 0:
		goto L192
	case 1:
		goto L202
	default:
		goto L200
	case 3:
		goto L201
	}
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L4
	} else {
		goto L203
	}
L201:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v799)))
	v852 = v807
	goto L191
L202:
	;
	v806 = int32(*(*int16)(unsafe.Add(mBase, uint32(v799))))
	v852 = v806
	goto L191
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = base.I32_extend16_s(v803)
	F_errmsg_internal(m, int32(_a_F_AddRelationNewConstraints_24), v29)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_25), int32(70), int32(_a_F_AddRelationNewConstraints_26))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	v852 = v823
	goto L191
L207:
	;
	v829 = F_nocachegetattr(m, v781, int32(28), v790)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L4
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L211
	}
L210:
	;
	v852 = v829
	goto L191
L211:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v835 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AddRelationNewConstraints_27), v29+int32(112))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2761), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
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
	v855 = F_stringToNode(m, v853)
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L215
	}
L215:
	;
	v857 = F_equal(m, v252, v855)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	v860 = v857
	goto L190
L217:
	;
	if v869&v860 == int32(0) {
		goto L36
	} else {
		goto L220
	}
L218:
	;
	v863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+103)))
	if v863 != 0 {
		v869 = l3
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+131)))
	v869 = l3 | (v865 ^ int32(1))
	goto L217
L220:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+106)))
	if v873 == int32(1) {
		goto L35
	} else {
		goto L221
	}
L221:
	;
	v878 = int32(*(*int16)(unsafe.Add(mBase, uint32(v786)+104)))
	if v749&int32(1)&base.B2i32(int32(0) < v878) != 0 {
		goto L34
	} else {
		goto L222
	}
L222:
	;
	if v750&int32(1) == int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	if v53&v751 == int32(1) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+75)))
	if v886 != int32(1) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+76)))
	if v889 == int32(0) {
		goto L33
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	v904 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L4
	} else {
		goto L234
	}
L228:
	;
	v895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+75)))
	if v895 != 0 {
		goto L227
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	if (v53|v751)&int32(1) != 0 {
		goto L227
	} else {
		goto L232
	}
L231:
	;
	goto L24
L232:
	;
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786)+75)))
	if v899 == int32(1) {
		goto L24
	} else {
		goto L233
	}
L233:
	;
	goto L227
L234:
	;
	if v904 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v255
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_29), v29+int32(48))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L4
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	v917 = F_heap_copytuple(m, v781)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L4
	} else {
		goto L240
	}
L238:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2826), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v917)+16))
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+22)))
	v921 = v919 + v920
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922)+131)))
	if v923 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v926 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v921)+104)) = uint16(v926)
	v1098 = int32(0)
	goto L32
L242:
	;
	goto L243
L243:
	;
	if v5 != 0 {
		v1098 = int32(1)
		goto L32
	} else {
		goto L244
	}
L244:
	;
	v930 = int32(*(*int16)(unsafe.Add(mBase, uint32(v921)+104)))
	v932 = v930 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v921)+104)) = uint16(v932)
	if base.I32_extend16_s(v932) == v932 {
		goto L31
	} else {
		goto L245
	}
L245:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L246
	}
L246:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L4
	} else {
		goto L247
	}
L247:
	;
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_9), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L4
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2849), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L4
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	F_relation_close(m, v754, int32(3))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L4
	} else {
		goto L251
	}
L251:
	;
	v966 = v255
	v971 = v747
	goto L37
L252:
	;
	v989 = F_palloc(m, int32(28))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L4
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v989)+16)) = v252
	v992 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v989)+12)) = uint16(v992)
	*(*int32)(unsafe.Add(mBase, uint32(v989)+8)) = v966
	*(*int32)(unsafe.Add(mBase, uint32(v989)+4)) = v986
	*(*int32)(unsafe.Add(mBase, uint32(v989))) = int32(5)
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+20)) = uint8(v998)
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v989)+24)) = uint16(v53)
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+22)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+21)) = uint8(v1000)
	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v989)+26)) = uint8(v1004)
	v1008 = F_lappend(m, v197, v989)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	v1135 = v971
	v1137 = v1008
	v1140 = v200 + int32(1)
	v1143 = v203
	goto L30
L255:
	;
	F_errcode(m, int32(_a_F_AddRelationNewConstraints_5))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v1017 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_23), v29+int32(96))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2781), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1039 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_30), v29+int32(16))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L4
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2788), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L4
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L4
	} else {
		goto L264
	}
L264:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v1061 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_31), v29+int32(32))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2799), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
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
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v1083 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_32), v29+int32(80))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2809), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
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
	v1103 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v921)+106)) = uint8(v1103)
	goto L273
L272:
	;
	goto L273
L273:
	;
	if v751&int32(1) == int32(0) {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	F_CatalogTupleUpdate(m, v754, v917+int32(4), v917)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L4
	} else {
		goto L277
	}
L275:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+75)))
	if v1109 != 0 {
		goto L274
	} else {
		goto L276
	}
L276:
	;
	v1110 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v921)+75)) = uint16(v1110)
	goto L274
L277:
	;
	F_systable_endscan(m, v779)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	F_relation_close(m, v754, int32(3))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	v1135 = v747
	v1137 = v197
	v1140 = v200
	v1143 = v203
	goto L30
L280:
	;
	goto L29
L281:
	;
	m.G0 = v29 + int32(352)
	return v1167
L282:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L4
	} else {
		goto L283
	}
L283:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v255
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1190 + int32(4)
	F_errmsg(m, int32(_a_F_AddRelationNewConstraints_33), v29-int32(-64))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(_a_F_AddRelationNewConstraints_3), int32(2821), int32(_a_F_AddRelationNewConstraints_28))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
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
					F_errmsg_internal(m, int32(_a_F_DeleteRelationTuple_0), v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_DeleteRelationTuple_1), int32(1586), int32(_a_F_DeleteRelationTuple_2))
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
				F_simple_heap_delete(m, v11, v14+int32(4))
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
						F_relation_close(m, v11, int32(3))
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
	var v88 int32
	_ = v88
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
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
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[0]))
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
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[1]))
	if v209 <= int32(0) {
		goto L10
	} else {
		goto L52
	}
L14:
	;
	v53 = v2
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[2]))
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
	v204 = v53 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[0]))
	if v204 < v206 {
		v53 = v204
		goto L15
	} else {
		goto L51
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
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(_a_F_FlushRelationBuffers_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(_a_F_FlushRelationBuffers_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = int32(_a_F_FlushRelationBuffers_2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
	v86 = int32(_a_F_FlushRelationBuffers_3)
	v88 = base.AtomicRmwOr32(m, v60, int32(24), v86)
	if v88&v86 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	goto L26
L24:
	;
	v110 = v88
	goto L25
L25:
	;
	v117 = int32(_a_F_FlushRelationBuffers_4)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[4]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(24))+8))
	if v120 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	F_perform_spin_delay(m, v10+int32(24))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	v110 = v104
	goto L25
L28:
	;
	v102 = int32(_a_F_FlushRelationBuffers_3)
	v104 = base.AtomicRmwOr32(m, v60, int32(24), v102)
	if v104&v102 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v137 != v138 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[4])) = v135
	goto L31
L33:
	;
	if int32(999) < v118 {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v118 < int32(11) {
		goto L31
	} else {
		goto L40
	}
L36:
	;
	v125 = int32(900)
	if v125 <= v118 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v128 = v125
	goto L39
L38:
	;
	v128 = v118
	goto L39
L39:
	;
	v135 = v128 + int32(100)
	goto L32
L40:
	;
	v135 = v118 - int32(1)
	goto L32
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v110 & int32(-4194305)
	goto L17
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v143 = int32(25165824)
	if base.B2i32(v140 != v141)|base.B2i32(v110&v143 != v143) != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v148 != v149 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v152 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = (v151 + v152) & int32(-4194305)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v158 = int32(_a_F_FlushRelationBuffers_5)
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[5])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+4)) = v152
	v166 = v157 + v152
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v166
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	F_ResourceOwnerRemember(m, v169, v166, int32(_a_F_FlushRelationBuffers_6))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v174 = v60 + int32(48)
	v176 = F_LWLockAcquire(m, v174, int32(1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_FlushBuffer(m, v60, v40, int32(3))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_LWLockRelease(m, v174)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	F_ResourceOwnerForget(m, v184, v185+int32(1), int32(_a_F_FlushRelationBuffers_6))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_UnpinBufferNoOwner(m, v60)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L17
L51:
	;
	goto L16
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[6]))
	v217 = v209
	v218 = v2
	v220 = v213
	goto L53
L53:
	;
	v223 = v220 + v218<<(uint(int32(6))%32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v224 != v225 {
		v271 = v217
		v272 = v220
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L10
L55:
	;
	v274 = v218 + int32(1)
	if v274 < v271 {
		v217 = v271
		v218 = v274
		v220 = v272
		goto L53
	} else {
		goto L65
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v227 != v228 {
		v271 = v217
		v272 = v220
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v230 != v231 {
		v271 = v217
		v272 = v220
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	v234 = int32(25165824)
	if v233&v234 != v234 {
		v271 = v217
		v272 = v220
		goto L55
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(1079)
	v241 = int32(_a_F_FlushRelationBuffers_7)
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[7])) = v10 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v242
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[3]))
	F_ResourceOwnerEnlarge(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_PinLocalBuffer(m, v223, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_FlushLocalBuffer(m, v223, v40)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	F_UnpinLocalBuffer(m, v259+int32(1))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[7])) = v265
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[1]))
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_FlushRelationBuffers[6]))
	v271 = v268
	v272 = v270
	goto L55
L65:
	;
	goto L54
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
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2 - int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDecrementReferenceCount[0]))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDecrementReferenceCount[1]))
		F_ResourceOwnerForget(m, v9, l0, int32(_a_F_RelationDecrementReferenceCount_0))
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+120)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v187
L2:
	;
	if l1 == int32(0) {
		v187 = v13
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v53 = v12 << (uint(int32(2)) % 32)
	v54 = F_palloc0(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L18
	}
L5:
	;
	v18 = F_palloc(m, v12<<(uint(int32(2))%32))
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
	if v12 <= int32(0) {
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
	v37 = v27 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13+v37)))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v18
L13:
	;
	v42 = F_datumCopy(m, v39, int32(0), int32(-1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	v44 = int32(0)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18+v37))) = v44
	v48 = v27 + int32(1)
	if v48 != v12 {
		v27 = v48
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v44 = v42
	goto L15
L17:
	;
	goto L12
L18:
	;
	v56 = int32(0)
	v57 = base.B2i32(v12 <= v56)
	if v57 == v56 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = v3
	goto L22
L20:
	;
	goto L21
L21:
	;
	v108 = int32(_a_F_RelationGetIndexAttOptions_0)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v111
	v113 = F_palloc(m, v53)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L31
	}
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[1])))
	if base.B2i32(v73 != int32(1))|base.B2i32(v51 == int32(2659)) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	v95 = v64 + int32(1)
	if v95 != v12 {
		v64 = v95
		goto L22
	} else {
		goto L30
	}
L25:
	;
	v82 = base.I32_extend16_s(v64 + int32(1))
	v83 = F_get_attoptions(m, v51, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v86 = F_index_opclass_options(m, l0, v82, v83, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v64<<(uint(int32(2))%32)))) = v86
	if v83 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	F_pfree(m, v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	if v57 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	F_pfree(m, v54)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L52
	}
L33:
	;
	v119 = int32(0)
	goto L36
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v113
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v109
	if l1 != 0 {
		v187 = v54
		goto L1
	} else {
		goto L51
	}
L36:
	;
	v129 = v119 << (uint(int32(2)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v54+v129)))
	if v131 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+252)) = v113
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttOptions[0])) = v109
	if l1 != 0 {
		v187 = v54
		goto L1
	} else {
		goto L43
	}
L38:
	;
	v134 = F_datumCopy(m, v131, int32(0), int32(-1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	v136 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113+v129))) = v136
	v140 = v119 + int32(1)
	if v140 != v12 {
		v119 = v140
		goto L36
	} else {
		goto L42
	}
L41:
	;
	v136 = v134
	goto L40
L42:
	;
	goto L37
L43:
	;
	v148 = int32(0)
	goto L44
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v54+v148<<(uint(int32(2))%32))))
	if v159 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	goto L32
L46:
	;
	F_pfree(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v163 = v148 + int32(1)
	if v163 != v12 {
		v148 = v163
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
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+252))
	v187 = v180
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
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
	var v308 int32
	_ = v308
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
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	v2 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v477 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v477
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v21 != 0 {
		v477 = v21
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
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v477 = v475
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L9
	} else {
		goto L110
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L9
	} else {
		goto L102
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L9
	} else {
		goto L99
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L9
	} else {
		goto L96
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[0]))
	v38 = F_AllocSetContextCreateInternal(m, v33, int32(_a_F_RelationGetPartitionKey_0), int32(0), int32(1024), int32(_a_F_RelationGetPartitionKey_1))
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
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L93
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
	v47 = F_MemoryContextAllocZero(m, v38, int32(56))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v51 = v49 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = base.I32_extend8_s(v52)
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)) = uint16(v55)
	v58 = v52 - int32(104)
	if base.B2i32(base.Ui32(int32(10)) < base.Ui32(v58))|base.B2i32(int32(1)<<(uint(v58)%32)&int32(1041) == int32(0)) != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v70 = F_SysCacheGetAttrNotNull(m, int32(45), v28, int32(6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v74 = F_SysCacheGetAttrNotNull(m, int32(45), v28, int32(7))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v80 = F_SysCacheGetAttr(m, int32(45), v28, int32(8), v24+int32(79))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+79)))
	if v82 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v38
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v112 = F_palloc0(m, v109<<(uint(int32(1))%32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L31
	}
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1]))
	v105 = v86
	goto L21
L23:
	;
	goto L24
L24:
	;
	v87 = F_text_to_cstring(m, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	v89 = F_stringToNode(m, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	F_pfree(m, v87)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v94 = F_eval_const_expressions(m, int32(0), v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	F_fix_opfuncids(m, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v98 = int32(_a_F_RelationGetPartitionKey_2)
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v38
	v102 = F_copyObjectImpl(m, v94)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v102
	v105 = v99
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v112
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v118 = F_palloc0(m, v115<<(uint(int32(2))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v118
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v124 = F_palloc0(m, v121<<(uint(int32(2))%32))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v124
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v130 = F_palloc0(m, v127*int32(28))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v130
	v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v136 = F_palloc0(m, v133<<(uint(int32(2))%32))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = v136
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v142 = F_palloc0(m, v139<<(uint(int32(2))%32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v142
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v148 = F_palloc0(m, v145<<(uint(int32(2))%32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v148
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v154 = F_palloc0(m, v151<<(uint(int32(1))%32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+40)) = v154
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v158 = F_palloc0(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+44)) = v158
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v162 = F_palloc0(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+48)) = v162
	v165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v168 = F_palloc0(m, v165<<(uint(int32(2))%32))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v168
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[1])) = v105
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	v176 = v174 << (uint(int32(1)) % 32)
	if v176 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	base.MemoryCopy(m, v177, v51+int32(36), v176)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if v181 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v183 = v182
	goto L47
L46:
	;
	v183 = v2
	goto L47
L47:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	if int32(0) < v184 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v173 == int32(104) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L75
	}
L51:
	;
	v191 = int32(2)
	goto L53
L52:
	;
	v191 = int32(1)
	goto L53
L53:
	;
	v192 = int32(24)
	v200 = int32(0)
	v203 = v183
	goto L54
L54:
	;
	v214 = v200 << (uint(int32(1)) % 32)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v214+v215))))
	v220 = v200 << (uint(int32(2)) % 32)
	v221 = v70 + v192 + v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v223 = F_SearchSysCache1(m, int32(14), v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L56
	}
L55:
	;
	goto L50
L56:
	;
	if v223 == int32(0) {
		goto L7
	} else {
		goto L57
	}
L57:
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
		goto L58
	}
L58:
	;
	if v239 == int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	F_fmgr_info_cxt(m, v239, v243+v200*int32(28), v38)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v220+(v74+v192))))
	*(*int32)(unsafe.Add(mBase, uint32(v249+v220))) = v252
	if v217 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v310+v220)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	F_get_typlenbyvalalign(m, v312, v313+v214, v315+v200, v317+v200)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L9
	} else {
		goto L72
	}
L62:
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
	v308 = v203
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v203 == int32(0) {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v279 = F_exprType(m, v278)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v47)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v281+v220))) = v279
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v285 = F_exprTypmod(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v287+v220))) = v285
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v291 = F_exprCollation(m, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v293+v220))) = v291
	v297 = v203 + int32(4)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	if base.Ui32(v297) < base.Ui32(v300+v301<<(uint(int32(2))%32)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v306 = v297
	goto L71
L70:
	;
	v306 = int32(0)
	goto L71
L71:
	;
	v308 = v306
	goto L61
L72:
	;
	F_ReleaseCatCache(m, v223)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v324 = v200 + int32(1)
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+4)))
	if v324 < v325 {
		v200 = v324
		v203 = v308
		goto L54
	} else {
		goto L74
	}
L74:
	;
	goto L55
L75:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionKey[2]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v350 != v346 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v38
	m.G0 = v24 + int32(80)
	goto L4
L77:
	;
	if v350 == int32(0) {
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
	if v346 != 0 {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v355 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v354 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355)+28)) = v354
	goto L82
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350)+20)) = v354
	goto L82
L86:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+24)) = v360
	goto L80
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v346
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v346)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v367
	if v367 != 0 {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	goto L89
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = int32(0)
	goto L79
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+24)) = v38
	goto L92
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+20)) = v38
	goto L76
L93:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v388
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_3), v24)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(99), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v402
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_6), v24+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(119), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
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
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v418
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionKey_7), v24+int32(32))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(202), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
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
	v436 = m.ExcPending
	if v436 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v231)+84))
	v439 = F_format_type_be(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v24)+56)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v24)+48)) = v231 + int32(8)
	if v437 == int32(104) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v450 = int32(_a_F_RelationGetPartitionKey_8)
	goto L107
L106:
	;
	v450 = int32(_a_F_RelationGetPartitionKey_9)
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+52)) = v450
	F_errmsg(m, int32(_a_F_RelationGetPartitionKey_10), v24+int32(48))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(221), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
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
	v469 = m.ExcPending
	if v469 != 0 {
		goto L9
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionKey_4), int32(240), int32(_a_F_RelationGetPartitionKey_5))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
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
	var v40 int32
	_ = v40
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
	v40 = v35
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
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
		v40 = v49
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_relation_close(m, v27, int32(1))
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
	v65 = int32(_a_F_RelationGetStatExtList_0)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0]))
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0])) = v69
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
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetStatExtList[0])) = v66
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
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(348), int32(_a_F_RelationMapUpdateMap_8))
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
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(351), int32(_a_F_RelationMapUpdateMap_8))
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
	F_errfinish(m, int32(_a_F_RelationMapUpdateMap_7), int32(403), int32(_a_F_RelationMapUpdateMap_11))
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
	var v523 int32
	_ = v523
	var v524 int32
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
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
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
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v931 int32
	_ = v931
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1043 int32
	_ = v1043
	var v1064 int32
	_ = v1064
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1193 int32
	_ = v1193
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1228 int32
	_ = v1228
	var v1243 int64
	_ = v1243
	var v1252 int32
	_ = v1252
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
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v1081 = v20 + int32(44)
	v1082 = int32(276)
	base.MemoryCopy(m, v1081, v203, v1082)
	base.MemoryCopy(m, v203, l0, v1082)
	base.MemoryCopy(m, l0, v1081, v1082)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+12)) = v1089
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1088
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+16)) = v1093
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1092
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v203)+32))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+32)) = v1097
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1096
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v203)+36))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+36)) = v1101
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1100
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v203)+40))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+40)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1104
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v203)+44))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+44)) = v1109
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1108
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+48)) = v1113
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1112
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	base.MemoryCopy(m, v1112, v1116, int32(144))
	if v705 != 0 {
		goto L250
	} else {
		goto L251
	}
L20:
	;
	v1064 = int32(0)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L12
	} else {
		goto L247
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L12
	} else {
		goto L244
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
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2689), int32(_a_F_RelationRebuildRelation_4))
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
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v226)+16))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	if v390 != 0 {
		goto L106
	} else {
		goto L107
	}
L72:
	;
	v235 = v225 << (uint(int32(4)) % 32)
	v237 = int32(20)
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
	v333 = v246 << (uint(int32(4)) % 32)
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v333)+31)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v333)+31)))
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
	v514 = v497 << (uint(int32(3)) % 32)
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
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v519)+4))
	v527 = v224 + int32(20) + v497<<(uint(int32(4))%32)
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v527)+6)))
	v529 = int32(*(*int16)(unsafe.Add(mBase, uint32(v527)+4)))
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
	v1064 = int32(1)
	goto L19
L192:
	;
	goto L193
L193:
	;
	v809 = int32(0)
	if base.B2i32(v803 != v809)^base.B2i32(v804 != v809) != 0 {
		v1064 = v809
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
		v1064 = v852
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
		v1064 = v852
		goto L19
	} else {
		goto L208
	}
L208:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v819)+12))
	if v860 == int32(0) {
		v1064 = v852
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
		v1064 = v868
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
		goto L243
	}
L214:
	;
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+4)))
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+4)))
	if v871 != v872 {
		v1064 = v868
		goto L19
	} else {
		goto L215
	}
L215:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+24)))
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v866)+24)))
	if v874 != v875 {
		v1064 = v868
		goto L19
	} else {
		goto L216
	}
L216:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877))))
	v884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v878))))
	if base.B2i32(v881 == int32(0))|base.B2i32(v881 != v884) != 0 {
		v902 = v881
		v903 = v884
		goto L218
	} else {
		goto L219
	}
L217:
	;
	if v902-v903 != 0 {
		v1064 = v868
		goto L19
	} else {
		goto L224
	}
L218:
	;
	goto L217
L219:
	;
	v887 = v877
	v888 = v878
	goto L220
L220:
	;
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888)+1)))
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)))
	if v892 == int32(0) {
		v902 = v892
		v903 = v891
		goto L218
	} else {
		goto L222
	}
L221:
	;
	v902 = v892
	v903 = v891
	goto L218
L222:
	;
	v895 = int32(1)
	if v892 == v891 {
		v887 = v887 + v895
		v888 = v888 + v895
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v867)+8))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+16))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v866)+8))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v907)+16))
	if v906 != v908 {
		v1064 = v868
		goto L19
	} else {
		goto L225
	}
L225:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v905)+8))
	if v910 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v905)+4))
	v920 = (v913<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L228
L227:
	;
	v920 = v910
	goto L228
L228:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v907)+8))
	if v921 == int32(0) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
	v931 = (v924<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L231
L230:
	;
	v931 = v921
	goto L231
L231:
	;
	if int32(0) < v906 {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v947 = int32(0)
	goto L235
L233:
	;
	goto L234
L234:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v867)+16))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v866)+16))
	v983 = F_equal(m, v981, v982)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L12
	} else {
		goto L239
	}
L235:
	;
	v955 = v947 << (uint(int32(2)) % 32)
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v905+v920+v955)))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v907+v931+v955)))
	if v957 != v959 {
		v1064 = v868
		goto L19
	} else {
		goto L237
	}
L236:
	;
	goto L234
L237:
	;
	v962 = v947 + int32(1)
	if v962 != v906 {
		v947 = v962
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	if v983 == int32(0) {
		v1064 = v868
		goto L19
	} else {
		goto L240
	}
L240:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v867)+20))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v866)+20))
	v989 = F_equal(m, v987, v988)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L12
	} else {
		goto L241
	}
L241:
	;
	if v989 == int32(0) {
		v1064 = v868
		goto L19
	} else {
		goto L242
	}
L242:
	;
	goto L210
L243:
	;
	goto L210
L244:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v1016
	F_errmsg_internal(m, int32(_a_F_RelationRebuildRelation_5), v20+int32(32))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2343), int32(_a_F_RelationRebuildRelation_6))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v1032
	F_errmsg_internal(m, int32(_a_F_RelationRebuildRelation_7), v20+int32(16))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L12
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_RelationRebuildRelation_3), int32(2314), int32(_a_F_RelationRebuildRelation_6))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L12
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L250:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v203)+52))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+52)) = v1120
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1119
	goto L252
L251:
	;
	goto L252
L252:
	;
	if v802 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v203)+68))
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+68)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v1124
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v203)+72))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+72)) = v1129
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v1128
	goto L255
L254:
	;
	goto L255
L255:
	;
	if v1064 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+80)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v1133
	goto L258
L257:
	;
	goto L258
L258:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v203)+264))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+264)) = v1139
	*(*int32)(unsafe.Add(mBase, uint32(l0)+264)) = v1138
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v203)+272))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+272)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v1142
	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+268)))
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+268)) = uint8(v1147)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)) = uint8(v1146)
	if v1079 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v203)+92))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+92)) = v1151
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v1150
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v203)+96))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+96)) = v1155
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1154
	goto L261
L260:
	;
	goto L261
L261:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v203)+104))
	if v1159 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	F_RelationDestroyRelation(m, v203, v705^int32(1))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L12
	} else {
		goto L309
	}
L263:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v203)+112))
	if v1162 == int32(0) {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1165 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v1165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1165
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v203)+104))
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v1172 != 0 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v203)+112))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v1207 != 0 {
		goto L289
	} else {
		goto L290
	}
L268:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+16))
	if v1176 != v1172 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v1171
	goto L267
L271:
	;
	goto L267
L272:
	;
	if v1176 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	goto L271
L275:
	;
	if v1172 != 0 {
		goto L282
	} else {
		goto L283
	}
L276:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+28))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+24))
	if v1181 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	if v1180 == int32(0) {
		goto L275
	} else {
		goto L281
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1181)+28)) = v1180
	goto L277
L279:
	;
	goto L280
L280:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1176)+20)) = v1180
	goto L277
L281:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1171)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1180)+24)) = v1186
	goto L275
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+16)) = v1172
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1172)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+28)) = v1193
	if v1193 != 0 {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	goto L284
L284:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1171)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+16)) = int32(0)
	goto L274
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1193)+24)) = v1171
	goto L287
L286:
	;
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1172)+20)) = v1171
	goto L271
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+116)) = int32(0)
	v1243 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v203)+108)) = v1243
	*(*int64)(unsafe.Add(mBase, uint32(v203)+100)) = v1243
	goto L262
L289:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+16))
	if v1211 != v1207 {
		goto L293
	} else {
		goto L294
	}
L290:
	;
	goto L291
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v1206
	goto L288
L292:
	;
	goto L288
L293:
	;
	if v1211 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	goto L295
L295:
	;
	goto L292
L296:
	;
	if v1207 != 0 {
		goto L303
	} else {
		goto L304
	}
L297:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+28))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+24))
	if v1216 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if v1215 == int32(0) {
		goto L296
	} else {
		goto L302
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1216)+28)) = v1215
	goto L298
L300:
	;
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1211)+20)) = v1215
	goto L298
L302:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1215)+24)) = v1221
	goto L296
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+16)) = v1207
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+28)) = v1228
	if v1228 != 0 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1206)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+16)) = int32(0)
	goto L295
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1228)+24)) = v1206
	goto L308
L307:
	;
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1207)+20)) = v1206
	goto L292
L309:
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
						F_errfinish(m, int32(_a_F_SetRelationNumChecks_1), int32(3160), int32(_a_F_SetRelationNumChecks_2))
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
	Fn13861(m, l0, l1, int32(0))
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
							F_errfinish(m, int32(_a_F_check_relation_block_range_1), int32(214), int32(_a_F_check_relation_block_range_2))
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
					F_errfinish(m, int32(_a_F_check_relation_block_range_1), int32(209), int32(_a_F_check_relation_block_range_2))
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
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
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
				v65 = v3
				m.G0 = v7 + int32(16)
				return v65
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
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(_a_F_try_relation_open_0), v7)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_try_relation_open_1), int32(115), int32(_a_F_try_relation_open_2))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
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
							v39 = int32(_a_F_try_relation_open_3)
							v41 = *(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0]))
							*(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0])) = v41 | int32(1)
						} else {
						}
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
						switch v47 - int32(83) {
						case 0, 22, 26, 29, 31, 33:
							v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_relation_open[1])))
							if v51 == int32(0) {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
								if v54 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v54)+128)) = int32(0)
								} else {
								}
								v60 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v60
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v60)
							} else {
								v57 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v57)
							}
						default:
							v60 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v60
							*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v60)
						}
						v65 = v31
						m.G0 = v7 + int32(16)
						return v65
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
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
								F_errmsg_internal(m, int32(_a_F_try_relation_open_0), v7)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_try_relation_open_1), int32(115), int32(_a_F_try_relation_open_2))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
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
								v39 = int32(_a_F_try_relation_open_3)
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_try_relation_open[0])) = v41 | int32(1)
							} else {
							}
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
							switch v47 - int32(83) {
							case 0, 22, 26, 29, 31, 33:
								v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_relation_open[1])))
								if v51 == int32(0) {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+272))
									if v54 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v54)+128)) = int32(0)
									} else {
									}
									v60 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v60
									*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v60)
								} else {
									v57 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v57)
								}
							default:
								v60 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v31)+272)) = v60
								*(*uint8)(unsafe.Add(mBase, uint32(v31)+268)) = uint8(v60)
							}
							v65 = v31
							m.G0 = v7 + int32(16)
							return v65
						}
					}
				} else {
					F_UnlockRelationOid(m, l0, l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v65 = v3
						m.G0 = v7 + int32(16)
						return v65
					}
				}
			}
		}
	}
}
