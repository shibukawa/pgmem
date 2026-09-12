package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BeginCopyFrom(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v255 int32
	_ = v255
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
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
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
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v535 int32
	_ = v535
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v921 int32
	_ = v921
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1094 int32
	_ = v1094
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1152 int64
	_ = v1152
	var v1178 int32
	_ = v1178
	var v1185 int32
	_ = v1185
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1316 int64
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1397 int32
	_ = v1397
	v5 = l4
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(272)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[355]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+264)) = v25
	v28 = *(*int64)(unsafe.Add(mBase, _consts[356]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+256)) = v28
	v31 = *(*int64)(unsafe.Add(mBase, _consts[357]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v31
	v34 = *(*int64)(unsafe.Add(mBase, _consts[358]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = v34
	v37 = *(*int64)(unsafe.Add(mBase, _consts[359]))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+224)) = v37
	v40 = F_palloc0(m, int32(352))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v50 = F_AllocSetContextCreateInternal(m, v45, int32(510304), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+204)) = v50
	v53 = int32(4520272)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v50
	v58 = v40 + int32(56)
	F_ProcessCopyOptions(m, l0, v58, int32(1), l7)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+62)))
	if v63 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+60)))
	if v68 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v70 = int32(1617912)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v74 = F_CopyGetAttnums(m, v73, l1, l6)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	v69 = int32(1617928)
	goto L10
L9:
	;
	v69 = int32(1617944)
	goto L10
L10:
	;
	v70 = v69
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v78 = base.I32_extend16_s(v77)
	v79 = F_palloc0(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+120)) = v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+116)))
	if v82 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v852 = F_palloc0(m, v851)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L206
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L193
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L189
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L185
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L181
	}
L18:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
	if v196 != 0 {
		goto L45
	} else {
		goto L46
	}
L19:
	;
	v87 = F__emscripten_memset_bulkmem(m, v79, base.I32_extend8_s(int32(1)), v78)
	mBase = m.M
	goto L22
L20:
	;
	goto L21
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v40)+112))
	if v88 == int32(0) {
		goto L18
	} else {
		goto L23
	}
L22:
	;
	goto L18
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v92 = F_CopyGetAttnums(m, v73, v91, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v92 == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v96 <= int32(0) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v101 = int32(0)
	goto L27
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v101<<(uint(int32(2))%32))))
	v125 = v123 - int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v128 = int32(0)
	if v127 == v128 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L18
L29:
	;
	if v166 == int32(0) {
		goto L17
	} else {
		goto L42
	}
L30:
	;
	v166 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v134 <= int32(0) {
		v159 = v128
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v166 = v159
	goto L29
L34:
	;
	v137 = int32(0)
	if v137 < v134 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v140 = v134
	goto L37
L36:
	;
	v140 = v137
	goto L37
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v143 = int32(0)
	goto L38
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141+v143<<(uint(int32(2))%32))))
	v152 = base.B2i32(v151 == v123)
	if v151 == v123 {
		v159 = v152
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v159 = v152
	goto L33
L40:
	;
	v154 = v143 + int32(1)
	if v154 != v140 {
		v143 = v154
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v40)+120))
	v171 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v169+v125))) = uint8(v171)
	v174 = v101 + v171
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v174 < v175 {
		v101 = v174
		goto L27
	} else {
		goto L43
	}
L43:
	;
	goto L28
L44:
	;
	v215 = F_palloc0(m, v78)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	v198 = F_palloc0(m, int32(12))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+220)) = int32(0)
	goto L44
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+220)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = int32(447)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v40)+220))
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+4)) = uint8(v204)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
	if v206 != int32(1) {
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v40)+220))
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)) = uint8(v210)
	goto L44
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+132)) = v215
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+128)))
	if v218 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+136)))
	if v332 != int32(1) {
		goto L77
	} else {
		goto L78
	}
L52:
	;
	v223 = F__emscripten_memset_bulkmem(m, v215, base.I32_extend8_s(int32(1)), v78)
	mBase = m.M
	goto L55
L53:
	;
	goto L54
L54:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v40)+124))
	if v224 == int32(0) {
		goto L51
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v228 = F_CopyGetAttnums(m, v73, v227, v224)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v228 == int32(0) {
		goto L51
	} else {
		goto L58
	}
L58:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v232 <= int32(0) {
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v237 = int32(0)
	goto L60
L60:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255+v237<<(uint(int32(2))%32))))
	v261 = v259 - int32(1)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v264 = int32(0)
	if v263 == v264 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L51
L62:
	;
	if v302 == int32(0) {
		goto L16
	} else {
		goto L75
	}
L63:
	;
	v302 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if v270 <= int32(0) {
		v295 = v264
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v302 = v295
	goto L62
L67:
	;
	v273 = int32(0)
	if v273 < v270 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v276 = v270
	goto L70
L69:
	;
	v276 = v273
	goto L70
L70:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v279 = int32(0)
	goto L71
L71:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v277+v279<<(uint(int32(2))%32))))
	v288 = base.B2i32(v287 == v259)
	if v287 == v259 {
		v295 = v288
		goto L66
	} else {
		goto L73
	}
L72:
	;
	v295 = v288
	goto L66
L73:
	;
	v290 = v279 + int32(1)
	if v290 != v276 {
		v279 = v290
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v40)+132))
	v307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v305+v261))) = uint8(v307)
	v310 = v237 + v307
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v310 < v311 {
		v237 = v310
		goto L60
	} else {
		goto L76
	}
L76:
	;
	goto L61
L77:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v444 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L78:
	;
	v335 = F_palloc0(m, v78)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+168)) = v335
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v40)+160))
	v340 = F_CopyGetAttnums(m, v73, v338, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v340 == int32(0) {
		goto L77
	} else {
		goto L81
	}
L81:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	if v344 <= int32(0) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v349 = int32(0)
	goto L83
L83:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v367+v349<<(uint(int32(2))%32))))
	v373 = v371 - int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v376 = int32(0)
	if v375 == v376 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L77
L85:
	;
	if v414 == int32(0) {
		goto L15
	} else {
		goto L98
	}
L86:
	;
	v414 = int32(0)
	goto L85
L87:
	;
	goto L88
L88:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v382 <= int32(0) {
		v407 = v376
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v414 = v407
	goto L85
L90:
	;
	v385 = int32(0)
	if v385 < v382 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v388 = v382
	goto L93
L92:
	;
	v388 = v385
	goto L93
L93:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v391 = int32(0)
	goto L94
L94:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v389+v391<<(uint(int32(2))%32))))
	v400 = base.B2i32(v399 == v371)
	if v399 == v371 {
		v407 = v400
		goto L89
	} else {
		goto L96
	}
L95:
	;
	v407 = v400
	goto L89
L96:
	;
	v402 = v391 + int32(1)
	if v402 != v388 {
		v391 = v402
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v40)+168))
	v419 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v417+v373))) = uint8(v419)
	v422 = v349 + v419
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	if v422 < v423 {
		v349 = v422
		goto L83
	} else {
		goto L99
	}
L99:
	;
	goto L84
L100:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _consts[360]))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	goto L103
L101:
	;
	v450 = v444
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	goto L107
L103:
	;
	v450 = v449
	goto L102
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+172)) = l2
	v479 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v479
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v479
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+48))
	v486 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v40)+184)) = v486
	*(*int64)(unsafe.Add(mBase, uint32(v40)+192)) = v486
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+200)) = uint8(v479)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+176)) = v485 + int32(4)
	v496 = F_palloc(m, int32(65537))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L115
	}
L105:
	;
	v464 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)) = uint8(v464)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v468 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	goto L112
L106:
	;
	v462 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)) = uint8(v462)
	goto L104
L107:
	;
	if v454 == v450 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if v456 == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v460 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	goto L110
L110:
	;
	if v461 != 0 {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	goto L106
L112:
	;
	v470 = F_FindDefaultConversionProc(m, v466, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = v470
	if v470 == int32(0) {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	goto L104
L115:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40)+328)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+324)) = v496
	v501 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+336)) = uint8(v501)
	F_initStringInfo(m, v40+int32(264))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if l0 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+248)) = v507
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+252)) = v509
	goto L119
L118:
	;
	goto L119
L119:
	;
	v513 = F_palloc(m, v78*int32(28))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v516 = v77 << (uint(int32(16)) % 32) >> (uint(int32(14)) % 32)
	v517 = F_palloc(m, v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v519 = F_palloc(m, v516)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v521 = F_palloc(m, v516)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	if v78 <= int32(0) {
		v846 = v479
		v848 = v9
		goto L13
	} else {
		goto L124
	}
L124:
	;
	v535 = int32(1)
	v542 = v479
	v544 = v9
	goto L125
L125:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v552 = v535 - int32(1)
	v555 = v73 + int32(20) + v547<<(uint(int32(4))%32) + v552*int32(100)
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+91)))
	if v556 != 0 {
		v702 = v542
		v703 = v544
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v846 = v702
	v848 = v703
	goto L13
L127:
	;
	if v535 != v78 {
		v535 = v535 + int32(1)
		v542 = v702
		v544 = v703
		goto L125
	} else {
		goto L180
	}
L128:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v555)+68))
	v562 = v552 << (uint(int32(2)) % 32)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	m.T0[v565].(func(*base.Module, int32, int32, int32, int32))(m, v40, v557, v513+v552*int32(28), v517+v562)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v568 = v562 + v521
	v569 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v40)+80))
	if v571 == v569 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v575 = int32(0)
	if v574 == v575 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	goto L132
L132:
	;
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555)+90)))
	if v614 != 0 {
		v702 = v542
		v703 = v544
		goto L127
	} else {
		goto L147
	}
L133:
	;
	if v613 != 0 {
		v702 = v542
		v703 = v544
		goto L127
	} else {
		goto L146
	}
L134:
	;
	v613 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	if v581 <= int32(0) {
		v606 = v575
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v613 = v606
	goto L133
L138:
	;
	v584 = int32(0)
	if v584 < v581 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v587 = v581
	goto L141
L140:
	;
	v587 = v584
	goto L141
L141:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v574)+12))
	v590 = int32(0)
	goto L142
L142:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v588+v590<<(uint(int32(2))%32))))
	v599 = base.B2i32(v598 == v535)
	if v598 == v535 {
		v606 = v599
		goto L137
	} else {
		goto L144
	}
L143:
	;
	v606 = v599
	goto L137
L144:
	;
	v601 = v590 + int32(1)
	if v601 != v587 {
		v590 = v601
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	goto L132
L147:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	v616 = F_build_column_default(m, v615, v535)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v616 == int32(0) {
		v702 = v542
		v703 = v544
		goto L127
	} else {
		goto L149
	}
L149:
	;
	v620 = F_expression_planner(m, v616)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v623 = F_ExecInitExpr(m, v620, int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v623
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	v627 = int32(0)
	if v626 == v627 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	if v665 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L153:
	;
	v665 = int32(0)
	goto L152
L154:
	;
	goto L155
L155:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	if v633 <= int32(0) {
		v658 = v627
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v665 = v658
	goto L152
L157:
	;
	v636 = int32(0)
	if v636 < v633 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v639 = v633
	goto L160
L159:
	;
	v639 = v636
	goto L160
L160:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v626)+12))
	v642 = int32(0)
	goto L161
L161:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v640+v642<<(uint(int32(2))%32))))
	v651 = base.B2i32(v650 == v535)
	if v650 == v535 {
		v658 = v651
		goto L156
	} else {
		goto L163
	}
L162:
	;
	v658 = v651
	goto L156
L163:
	;
	v653 = v642 + int32(1)
	if v653 != v639 {
		v642 = v653
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519+base.I32_extend16_s(v544)<<(uint(int32(2))%32)))) = v552
	v675 = v544 + int32(1)
	goto L167
L166:
	;
	v675 = v544
	goto L167
L167:
	;
	if v542&int32(1) != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v702 = int32(1)
	v703 = v675
	goto L127
L169:
	;
	goto L170
L170:
	;
	v679 = int32(0)
	if v620 == v679 {
		v699 = v679
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v702 = v699
	v703 = v675
	goto L127
L172:
	;
	v685 = F_check_functions_in_node(m, v620, int32(862), int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	if v685 != 0 {
		v699 = int32(1)
		goto L171
	} else {
		goto L174
	}
L174:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	if v687 == int32(67) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v691 = int32(0)
	v693 = F_query_tree_walker_impl(m, v620, int32(863), v691, v691)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v697 = F_expression_tree_walker_impl(m, v620, int32(863), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	v699 = v693
	goto L171
L179:
	;
	v699 = v697
	goto L171
L180:
	;
	goto L126
L181:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = int32(533507)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = v73 + v126<<(uint(int32(4))%32) + v125*int32(100) + int32(24)
	F_errmsg(m, int32(510018), v22+int32(112))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errfinish(m, int32(498284), int32(1612), int32(289110))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L185:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(533522)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v73 + v262<<(uint(int32(4))%32) + v261*int32(100) + int32(24)
	F_errmsg(m, int32(510018), v22+int32(96))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(498284), int32(1655), int32(289110))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v73 + v374<<(uint(int32(4))%32) + v373*int32(100) + int32(24)
	F_errmsg_internal(m, int32(510056), v22+int32(80))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(498284), int32(1679), int32(289110))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	if base.Ui32(v796) <= base.Ui32(int32(41)) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v808 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)+4))
	goto L199
L196:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v796<<(uint(int32(3))%32))+uint32(_consts[362])))
	v806 = v805
	goto L198
L197:
	;
	v806 = int32(758841)
	goto L198
L198:
	;
	goto L195
L199:
	;
	if base.Ui32(v809) <= base.Ui32(int32(41)) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v806
	F_errmsg(m, int32(71486), v22-int32(-64))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L204
	}
L201:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v809<<(uint(int32(3))%32))+uint32(_consts[362])))
	v819 = v818
	goto L203
L202:
	;
	v819 = int32(758841)
	goto L203
L203:
	;
	goto L200
L204:
	;
	F_errfinish(m, int32(498284), int32(1709), int32(289110))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v40)+240)) = v852
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v856 != 0 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)+56))
	v859 = v857
	goto L209
L208:
	;
	v859 = int32(0)
	goto L209
L209:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v862 == int32(0) {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v921 = v846 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+244)) = uint8(v921)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+236)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v40)+232)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v40)+216)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v40)+212)) = v513
	*(*int64)(unsafe.Add(mBase, uint32(v40)+344)) = int64(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v40)+208)) = uint16(v848)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+44)) = uint8(v5)
	if l5 != 0 {
		goto L225
	} else {
		goto L226
	}
L211:
	;
	goto L210
L212:
	;
	v866 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v866 != int32(1) {
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v869 = int32(4514932)
	v871 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v872 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v871 + v872
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v875 + v872
	*(*int32)(unsafe.Add(mBase, uint32(v862)+220)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v862)+224)) = v859
	v882 = v862 + int32(232)
	if v882&int32(3) == int32(0) {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v862)))
	v909 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = v908 + v909
	v912 = int32(4514932)
	v914 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v914 - v909
	goto L211
L215:
	;
	v888 = v862 + int32(392)
	if base.Ui32(v888) <= base.Ui32(v882) {
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v905 = F___memset(m, v882, int32(0), int32(160))
	mBase = m.M
	goto L214
L218:
	;
	v892 = v862 + int32(236)
	if base.Ui32(v892) < base.Ui32(v888) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v894 = v888
	goto L221
L220:
	;
	v894 = v892
	goto L221
L221:
	;
	v902 = F___memset(m, v882, int32(0), (v894-v862-int32(233))&int32(-4)+int32(4))
	mBase = m.M
	goto L214
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L308
	}
L223:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L304
	}
L224:
	;
	v1178 = int32(0)
	v1185 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v1185 == v1178 {
		goto L287
	} else {
		goto L288
	}
L225:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = int64(4)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(2)
	goto L224
L226:
	;
	goto L227
L227:
	;
	if l3 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = int64(3)
	v941 = *(*int32)(unsafe.Add(mBase, _consts[363]))
	if v941 == int32(2) {
		goto L231
	} else {
		goto L232
	}
L229:
	;
	goto L230
L230:
	;
	v1057 = F_pstrdup(m, l3)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L250
	}
L231:
	;
	v944 = m.G0
	v946 = v944 - int32(16)
	m.G0 = v946
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v40)+36))
	if v948 != 0 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1055 = *(*int32)(unsafe.Add(mBase, _consts[364]))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v1055
	goto L224
L234:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
	v951 = v949
	goto L236
L235:
	;
	v951 = int32(0)
	goto L236
L236:
	;
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+60)))
	F_pq_beginmessage(m, v946, int32(71))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_enlargeStringInfo(m, v946, int32(1))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	*(*uint8)(unsafe.Add(mBase, uint32(v959+v960))) = uint8(v952)
	*(*int32)(unsafe.Add(mBase, uint32(v946)+4)) = v959 + int32(1)
	F_enlargeStringInfo(m, v946, int32(2))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	v972 = int32(8)
	v978 = v951<<(uint(v972)%32) | int32(base.Ui32(v951&int32(65280))>>(uint(v972)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v969+v970))) = uint16(v978)
	*(*int32)(unsafe.Add(mBase, uint32(v946)+4)) = v969 + int32(2)
	if int32(0) < v951 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v986 = v952 << (uint(int32(8)) % 32)
	v989 = int32(0)
	goto L243
L241:
	;
	goto L242
L242:
	;
	F_pq_endmessage(m, v946)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L247
	}
L243:
	;
	F_enlargeStringInfo(m, v946, int32(2))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L245
	}
L244:
	;
	goto L242
L245:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v946)+4))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1010+v1011))) = uint16(v986)
	*(*int32)(unsafe.Add(mBase, uint32(v946)+4)) = v1010 + int32(2)
	v1018 = v989 + int32(1)
	if v1018 != v951 {
		v989 = v1018
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = int32(1)
	v1043 = F_makeStringInfo(m)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v1043
	v1047 = *(*int32)(unsafe.Add(mBase, _consts[268]))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1047)+4))
	v1049 = m.T0[v1048].(func(*base.Module) int32)(m)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	m.G0 = v946 + int32(16)
	goto L224
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v1057
	v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+44)))
	if v1060 == int32(1) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = int64(2)
	v1066 = F_OpenPipeStream(m, v1057, int32(231353))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+232)) = int64(1)
	v1088 = F_AllocateFile(m, v1057, int32(231353))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L1
	} else {
		goto L260
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v1066
	if v1066 != 0 {
		goto L224
	} else {
		goto L255
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1075
	F_errmsg(m, int32(300511), v22)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(498284), int32(1864), int32(289110))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v1088
	if v1088 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+76))
	if v1124 < int32(0) {
		goto L274
	} else {
		goto L275
	}
L264:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v1101
	F_errmsg(m, int32(294941), v22+int32(16))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	if base.B2i32(v1094 != int32(44))&base.B2i32(v1094 != int32(2)) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_errhint(m, int32(575792), int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	F_errfinish(m, int32(498284), int32(1883), int32(289110))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	if v1136 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L273:
	;
	if v1129 < int32(0) {
		goto L277
	} else {
		goto L278
	}
L274:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+60))
	v1129 = v1127
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+60))
	v1129 = v1128
	goto L273
L277:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = int32(8)
	v1136 = int32(-1)
	goto L279
L278:
	;
	v1136 = v1129
	goto L279
L279:
	;
	goto L272
L280:
	;
	if v1146 != 0 {
		goto L223
	} else {
		goto L284
	}
L281:
	;
	v1142 = F___syscall_ret(m, int32(-8))
	mBase = m.M
	v1146 = v1142
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1145 = F___fstatat(m, v1136, int32(758841), v22+int32(128), int32(4096))
	mBase = m.M
	v1146 = v1145
	goto L280
L284:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v22)+132))
	if v1147&int32(61440) == int32(16384) {
		goto L222
	} else {
		goto L285
	}
L285:
	;
	v1152 = *(*int64)(unsafe.Add(mBase, uint32(v22)+152))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+240)) = v1152
	goto L224
L286:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1351)+4))
	m.T0[v1352].(func(*base.Module, int32, int32))(m, v40, v73)
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L303
	}
L287:
	;
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v1191&int32(1) == int32(0) {
		goto L287
	} else {
		goto L290
	}
L290:
	;
	v1196 = int32(4514932)
	v1198 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v1199 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1198 + v1199
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1185)))
	*(*int32)(unsafe.Add(mBase, uint32(v1185))) = v1202 + v1199
	goto L292
L291:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1185)))
	v1333 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1185))) = v1332 + v1333
	v1336 = int32(4514932)
	v1338 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1338 - v1333
	goto L287
L292:
	;
	goto L294
L294:
	;
	goto L295
L295:
	;
	goto L299
L299:
	;
	v1297 = int32(0)
	v1300 = v1178
	goto L300
L300:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(256)+v1300<<(uint(int32(2))%32))))
	v1310 = int32(3)
	v1316 = *(*int64)(unsafe.Add(mBase, uint32(v22+int32(224)+v1300<<(uint(v1310)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1185+int32(232)+v1309<<(uint(v1310)%32)))) = v1316
	v1318 = int32(1)
	v1321 = v1297 + v1318
	if v1321 != int32(3) {
		v1297 = v1321
		v1300 = v1300 + v1318
		goto L300
	} else {
		goto L302
	}
L301:
	;
	goto L291
L302:
	;
	goto L301
L303:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v54
	m.G0 = v22 + int32(272)
	return v40
L304:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v1367
	F_errmsg(m, int32(298441), v22+int32(48))
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_errfinish(m, int32(498284), int32(1890), int32(289110))
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L308:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v1386
	F_errmsg(m, int32(13583), v22+int32(32))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L310
	}
L310:
	;
	F_errfinish(m, int32(498284), int32(1895), int32(289110))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromBinaryOneRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
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
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+4)))
	v31 = v29
	goto L3
L2:
	;
	v31 = int32(0)
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+52))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+184)) = v35 + int64(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v39-v40 <= int32(1) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L21
	} else {
		goto L137
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L21
	} else {
		goto L133
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L21
	} else {
		goto L129
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L21
	} else {
		goto L125
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L21
	} else {
		goto L121
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v111 + int32(1)
	goto L4
L10:
	;
	m.G0 = v25 + int32(16)
	return v528
L11:
	;
	v132 = int32(65535)
	if v115&v132 == v132 {
		goto L33
	} else {
		goto L34
	}
L12:
	;
	v47 = v40
	v50 = v39
	v52 = int32(0)
	v54 = v25 + int32(10)
	goto L16
L13:
	;
	goto L14
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104+v40))))
	v108 = v40 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v108
	v111 = v108
	v114 = v39
	v115 = v106
	v117 = v104
	goto L11
L15:
	;
	v528 = int32(0)
	goto L10
L16:
	;
	if v47 == v50 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v90 != int32(2) {
		goto L15
	} else {
		goto L32
	}
L18:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v76 = v47
	v77 = v50
	goto L20
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v81 = int32(2) - v52
	v82 = v77 - v76
	if v81 < v82 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	return int32(0)
L22:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v73 != 0 {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v76 = v75
	v77 = v74
	goto L20
L24:
	;
	v84 = v81
	goto L26
L25:
	;
	v84 = v82
	goto L26
L26:
	;
	if v84 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v87 = v76 + v84
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v87
	v90 = v84 + v52
	if v90 < int32(2) {
		v47 = v87
		v50 = v77
		v52 = v90
		v54 = v84 + v86
		goto L16
	} else {
		goto L31
	}
L28:
	;
	v85 = F__emscripten_memcpy_bulkmem(m, v54, v78+v76, v84)
	mBase = m.M
	v86 = v85
	goto L30
L29:
	;
	v86 = v54
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L17
L32:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+10)))
	v111 = v87
	v114 = v77
	v115 = v95
	v117 = v78
	goto L11
L33:
	;
	v136 = int32(0)
	if v136 < v114-v111 {
		goto L9
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v189 = int32(8)
	v196 = base.I32_extend16_s(v115<<(uint(v189)%32) | int32(base.Ui32(v115&int32(65280))>>(uint(v189)%32)))
	if v31 != v196 {
		goto L8
	} else {
		goto L53
	}
L36:
	;
	v143 = v111
	v146 = v114
	v148 = v136
	v149 = v117
	v150 = v25 + int32(9)
	goto L37
L37:
	;
	if v143 == v146 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v528 = int32(0)
	goto L10
L39:
	;
	goto L38
L40:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L21
	} else {
		goto L43
	}
L41:
	;
	v171 = v143
	v172 = v146
	v173 = v149
	goto L42
L42:
	;
	v176 = int32(1) - v148
	v177 = v172 - v171
	if v176 < v177 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v167 != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v171 = v170
	v172 = v169
	v173 = v168
	goto L42
L45:
	;
	v179 = v176
	goto L47
L46:
	;
	v179 = v177
	goto L47
L47:
	;
	if v179 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v182 = v171 + v179
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v182
	v185 = v179 + v148
	if v185 <= int32(0) {
		v143 = v182
		v146 = v172
		v148 = v185
		v149 = v173
		v150 = v179 + v181
		goto L37
	} else {
		goto L52
	}
L49:
	;
	v180 = F__emscripten_memcpy_bulkmem(m, v150, v171+v173, v179)
	mBase = m.M
	v181 = v180
	goto L51
L50:
	;
	v181 = v150
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L4
L53:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v198 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v528 = int32(1)
	goto L10
L55:
	;
	goto L56
L56:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v203 <= int32(0) {
		v528 = int32(1)
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v209 = l0 + int32(264)
	v222 = int32(0)
	goto L58
L58:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v234 = int32(4)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	v238 = int32(2)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v237+v222<<(uint(v238)%32))))
	v243 = v241 - int32(1)
	v246 = v34 + int32(20) + v233<<(uint(v234)%32) + v243*int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v246 + v234
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v246)+76))
	v252 = v243 << (uint(v238) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v32+v252)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v257-v258 < v234 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v528 = v522
	goto L10
L60:
	;
	v361 = l3 + v243
	v364 = v33 + v243*int32(28)
	if v343 == int32(-1) {
		goto L86
	} else {
		goto L87
	}
L61:
	;
	v264 = v258
	v268 = int32(0)
	v269 = v257
	v270 = v25 + int32(12)
	goto L65
L62:
	;
	goto L63
L63:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v333+v258)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v258 + int32(4)
	v343 = v335
	goto L60
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L21
	} else {
		goto L81
	}
L65:
	;
	if v264 == v269 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v305 != int32(4) {
		goto L64
	} else {
		goto L80
	}
L67:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L21
	} else {
		goto L70
	}
L68:
	;
	v291 = v264
	v292 = v269
	goto L69
L69:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v296 = int32(4) - v268
	v297 = v292 - v291
	if v296 < v297 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v288 != 0 {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v291 = v290
	v292 = v289
	goto L69
L72:
	;
	v299 = v296
	goto L74
L73:
	;
	v299 = v297
	goto L74
L74:
	;
	if v299 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v302 = v291 + v299
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v302
	v305 = v299 + v268
	if v305 < int32(4) {
		v264 = v302
		v268 = v305
		v269 = v292
		v270 = v299 + v301
		goto L65
	} else {
		goto L79
	}
L76:
	;
	v300 = F__emscripten_memcpy_bulkmem(m, v270, v293+v291, v299)
	mBase = m.M
	v301 = v300
	goto L78
L77:
	;
	v301 = v270
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L66
L80:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v343 = v310
	goto L60
L81:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L21
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(506335), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L21
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(499867), int32(2023), int32(349915))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L21
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
	*(*int32)(unsafe.Add(mBase, uint32(l2+v252))) = v497
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = int32(0)
	v522 = int32(1)
	v524 = v222 + v522
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v524 < v525 {
		v222 = v524
		goto L58
	} else {
		goto L120
	}
L86:
	;
	v367 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v367)
	v370 = F_ReceiveFunctionCall(m, v364, int32(0), v254, v250)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L21
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v372 = int32(24)
	v374 = int32(65280)
	v376 = int32(8)
	v386 = v343<<(uint(v372)%32) | v343&v374<<(uint(v376)%32) | (int32(base.Ui32(v343)>>(uint(v376)%32))&v374 | int32(base.Ui32(v343)>>(uint(v372)%32)))
	if v386 < int32(0) {
		goto L7
	} else {
		goto L90
	}
L89:
	;
	v497 = v370
	goto L85
L90:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v390)
	*(*int32)(unsafe.Add(mBase, uint32(v209)+12)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v209)+4)) = v390
	goto L91
L91:
	;
	F_enlargeStringInfo(m, v209, v386)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L21
	} else {
		goto L92
	}
L92:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	if v386 <= v400-v401 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+268)) = v386
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+264))
	v487 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v485+v386))) = uint8(v487)
	v489 = F_ReceiveFunctionCall(m, v364, v209, v254, v250)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L21
	} else {
		goto L118
	}
L94:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	if v386 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	goto L96
L96:
	;
	v412 = v401
	v416 = int32(0)
	v418 = v398
	goto L101
L97:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v408 + v386
	goto L93
L98:
	;
	v406 = F__emscripten_memcpy_bulkmem(m, v398, v404+v401, v386)
	mBase = m.M
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L97
L101:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	if v412 == v433 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	if v458 != v386 {
		goto L6
	} else {
		goto L117
	}
L103:
	;
	goto L102
L104:
	;
	F_CopyLoadRawBuf(m, l0)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L21
	} else {
		goto L107
	}
L105:
	;
	v440 = v412
	v441 = v433
	goto L106
L106:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v444 = v386 - v416
	v445 = v441 - v440
	if v444 < v445 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)))
	if v437 != 0 {
		v458 = v416
		goto L103
	} else {
		goto L108
	}
L108:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v440 = v439
	v441 = v438
	goto L106
L109:
	;
	v447 = v444
	goto L111
L110:
	;
	v447 = v445
	goto L111
L111:
	;
	if v447 != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v451 = v450 + v447
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = v451
	v454 = v447 + v416
	if v454 < v386 {
		v412 = v451
		v416 = v454
		v418 = v447 + v449
		goto L101
	} else {
		goto L116
	}
L113:
	;
	v448 = F__emscripten_memcpy_bulkmem(m, v418, v442+v440, v447)
	mBase = m.M
	v449 = v448
	goto L115
L114:
	;
	v449 = v418
	goto L115
L115:
	;
	goto L112
L116:
	;
	v458 = v454
	goto L103
L117:
	;
	goto L93
L118:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+276))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+268))
	if v491 != v492 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	v494 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v361))) = uint8(v494)
	v497 = v489
	goto L85
L120:
	;
	goto L59
L121:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v196
	F_errmsg(m, int32(480214), v25)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L21
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(499867), int32(1130), int32(32243))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(342273), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L21
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(499867), int32(2032), int32(349915))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L21
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(506335), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(499867), int32(2042), int32(349915))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(112602), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(499867), int32(2055), int32(349915))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	F_errcode(m, int32(67240066))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L21
	} else {
		goto L138
	}
L138:
	;
	F_errmsg(m, int32(221778), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L21
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(499867), int32(1122), int32(32243))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L21
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CopyFromErrorCallback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v147 int32
	_ = v147
	v7 = m.G0
	v9 = v7 - int32(176)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+200)))
	if v11 == int32(1) {
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
			F_errcontext_msg(m, int32(198255), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v9 + int32(176)
				return
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
		if v23 == int32(1) {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
				if v22 != 0 {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v31
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v30
					F_errcontext_msg(m, int32(185041), v9+int32(32))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v30
					F_errcontext_msg(m, int32(38527), v9+int32(16))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v9 + int32(176)
						return
					}
				}
			}
		} else {
			if v22 != 0 {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
				if v47 != 0 {
					v48 = F_strlen(m, v47)
					mBase = m.M
					if v48 <= int32(100) {
						v51 = F_pstrdup(m, v47)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v65 = v51
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
								v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v65
								*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v70
								F_errcontext_msg(m, int32(726487), v9+int32(144))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									F_pfree(m, v65)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v54 = F_pg_mbcliplen(m, v47, v48, int32(100))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v58 = F_palloc(m, v54+int32(4))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								if v54 != 0 {
									v60 = F__emscripten_memcpy_bulkmem(m, v58, v47, v54)
									mBase = m.M
									v61 = v60
								} else {
									v61 = v58
								}
								*(*int32)(unsafe.Add(mBase, uint32(v61+v54))) = int32(3026478)
								v65 = v58
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
									v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v65
									*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(v9)+152)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v70
									F_errcontext_msg(m, int32(726487), v9+int32(144))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_pfree(m, v65)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
						v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v89
						*(*int64)(unsafe.Add(mBase, uint32(v9)+120)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v87
						F_errcontext_msg(m, int32(65095), v9+int32(112))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			} else {
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+304)))
				if v98 == int32(1) {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
					v102 = F_strlen(m, v101)
					mBase = m.M
					if v102 <= int32(100) {
						v105 = F_pstrdup(m, v101)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return
						} else {
							v119 = v105
							F_set_errcontext_domain(m, int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
								v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v119
								*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v125
								*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v124
								F_errcontext_msg(m, int32(726215), v9-int32(-64))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return
								} else {
									F_pfree(m, v119)
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										m.G0 = v9 + int32(176)
										return
									}
								}
							}
						}
					} else {
						v108 = F_pg_mbcliplen(m, v101, v102, int32(100))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return
						} else {
							v112 = F_palloc(m, v108+int32(4))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return
							} else {
								if v108 != 0 {
									v114 = F__emscripten_memcpy_bulkmem(m, v112, v101, v108)
									mBase = m.M
									v115 = v114
								} else {
									v115 = v112
								}
								*(*int32)(unsafe.Add(mBase, uint32(v115+v108))) = int32(3026478)
								v119 = v112
								F_set_errcontext_domain(m, int32(0))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
									v125 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v119
									*(*int64)(unsafe.Add(mBase, uint32(v9)+72)) = v125
									*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v124
									F_errcontext_msg(m, int32(726215), v9-int32(-64))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										F_pfree(m, v119)
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return
										} else {
											m.G0 = v9 + int32(176)
											return
										}
									}
								}
							}
						}
					}
				} else {
					F_set_errcontext_domain(m, int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
						*(*int64)(unsafe.Add(mBase, uint32(v9)+104)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v139
						F_errcontext_msg(m, int32(38527), v9+int32(96))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return
						} else {
							m.G0 = v9 + int32(176)
							return
						}
					}
				}
			}
		}
	}
}
func F_CopyFromTextLikeInFunc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_getTypeInputInfo(m, l1, v7+int32(12), l3)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		F_fmgr_info(m, v13, l2)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			m.G0 = v7 + int32(16)
			return
		}
	}
}
func F_RemoveFromWaitQueue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v20 - v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = v24 - v21
	v30 = v10 + v12<<(uint(int32(2))%32)
	v32 = v30 + int32(44)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = v33 - v21
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+88))
	if v35 == v37 {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v39 & base.I32_rotl(int32(-2), v12)
	} else {
	}
	v44 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(l0)+92)) = int64(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(v44)%32))+uint32(_consts[100])))
	F_CleanUpLock(m, v10, v9, v52, l1, int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		return
	} else {
		return
	}
}
