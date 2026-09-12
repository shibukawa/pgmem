package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessStartupPacket(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
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
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v586 int32
	_ = v586
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1168 int32
	_ = v1168
	var v1175 int32
	_ = v1175
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1210 int32
	_ = v1210
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1323 int32
	_ = v1323
	var v1329 int32
	_ = v1329
	var v1332 int32
	_ = v1332
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1485 int32
	_ = v1485
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	F_pq_startmsgread(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(-1)
	v24 = F_pq_getbytes(m, v12+int32(-20), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L402
	}
L4:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L398
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L1
	} else {
		goto L394
	}
L6:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L389
	}
L7:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L1
	} else {
		goto L384
	}
L8:
	;
	m.G0 = v14 - int32(-64)
	return v1377
L9:
	;
	if v24 == int32(-1) {
		v1377 = v20
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v33 = F_pq_getbytes(m, v12+int32(-20)|int32(1), int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v33 == int32(-1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if l1 != 0 {
		v1377 = v20
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v56 = int32(24)
	v58 = int32(65280)
	v60 = int32(8)
	v70 = v55<<(uint(v56)%32) | v55&v58<<(uint(v60)%32) | (int32(base.Ui32(v55)>>(uint(v60)%32))&v58 | int32(base.Ui32(v55)>>(uint(v56)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v70 - int32(4)
	if base.Ui32(v70-int32(10005)) <= base.Ui32(int32(-9998)) {
		goto L22
	} else {
		goto L23
	}
L15:
	;
	if l2 != 0 {
		v1377 = v20
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v39 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v39 == int32(0) {
		v1377 = v20
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errmsg(m, int32(101357), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(476542), int32(540), int32(101498))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v1377 = v20
	goto L8
L22:
	;
	v78 = int32(-1)
	v81 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v99 = F_palloc(m, v70-int32(3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	if v81 == int32(0) {
		v1377 = v78
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(101324), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(476542), int32(552), int32(101498))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v1377 = v78
	goto L8
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99+v101))) = uint8(v103)
	v105 = int32(-1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v107 = F_pq_getbytes(m, v99, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v107 == int32(-1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v113 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v130 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[646])) = uint8(v130)
	goto L40
L35:
	;
	if v113 == int32(0) {
		v1377 = v105
		goto L8
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(101357), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(476542), int32(568), int32(101498))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v1377 = v105
	goto L8
L40:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v133 = int32(24)
	v135 = int32(65280)
	v137 = int32(8)
	v147 = v132<<(uint(v133)%32) | v132&v135<<(uint(v137)%32) | (int32(base.Ui32(v132)>>(uint(v137)%32))&v135 | int32(base.Ui32(v132)>>(uint(v133)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v147
	if v132 != int32(806801924) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v640 = int32(0)
	v642 = int32(196610)
	if base.Ui32(v642) <= base.Ui32(v147) {
		goto L168
	} else {
		goto L169
	}
L42:
	;
	if v132 != int32(790024708) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if l2 != 0 {
		goto L41
	} else {
		goto L147
	}
L45:
	;
	if v132 != int32(773247492) {
		goto L41
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l1 != 0 {
		goto L41
	} else {
		goto L126
	}
L48:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.Ui32(v155) <= base.Ui32(int32(7)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v158 = int32(-1)
	v161 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v178 = v155 - int32(8)
	if v178 < int32(257) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	if v161 == int32(0) {
		v1377 = v158
		goto L8
	} else {
		goto L53
	}
L53:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errmsg(m, int32(101232), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(476542), int32(896), int32(101444))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v1377 = v158
	goto L8
L57:
	;
	v182 = v178
	goto L59
L58:
	;
	v182 = int32(0)
	goto L59
L59:
	;
	if v182 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v185 = int32(-1)
	v188 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v205 = int32(24)
	v207 = int32(65280)
	v209 = int32(8)
	v219 = v204<<(uint(v205)%32) | v204&v207<<(uint(v209)%32) | (int32(base.Ui32(v204)>>(uint(v209)%32))&v207 | int32(base.Ui32(v204)>>(uint(v205)%32)))
	v221 = v99 + v209
	v223 = m.G0
	v225 = v223 - int32(48)
	m.G0 = v225
	if v219 != 0 {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	if v188 == int32(0) {
		v1377 = v185
		goto L8
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(101178), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(476542), int32(904), int32(101444))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v1377 = v185
	goto L8
L68:
	;
	m.G0 = v225 + int32(48)
	v1377 = int32(-1)
	goto L8
L69:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	if v178 == v319 {
		goto L96
	} else {
		goto L97
	}
L70:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	if int32(0) < v228+int32(38) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v306 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L91
	}
L73:
	;
	v233 = v228
	v235 = int32(0)
	goto L76
L74:
	;
	goto L75
L75:
	;
	v289 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L87
	}
L76:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[606]))
	v248 = v245 + v235<<(uint(int32(7))%32)
	v250 = v248 + int32(8)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v219 == v251 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v250)+96)) = int32(1)
	v257 = v248 + int32(104)
	if v253 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v269 = v233
	goto L80
L80:
	;
	v272 = v235 + int32(1)
	if v272 < v269+int32(38) {
		v233 = v269
		v235 = v272
		goto L76
	} else {
		goto L86
	}
L81:
	;
	F_s_lock(m, v257, int32(478339), int32(754), int32(74257))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	if v263 == v219 {
		goto L69
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(0)
	v268 = *(*int32)(unsafe.Add(mBase, _consts[607]))
	v269 = v268
	goto L80
L86:
	;
	goto L77
L87:
	;
	if v289 == int32(0) {
		goto L68
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+32)) = v219
	F_errmsg(m, int32(120581), v225+int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(478339), int32(798), int32(74257))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	goto L68
L91:
	;
	if v306 == int32(0) {
		goto L68
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(534444), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(478339), int32(733), int32(74257))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	goto L68
L95:
	;
	v462 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L122
	}
L96:
	;
	v322 = v248 + int32(16)
	v323 = int32(0)
	if v178 == v323 {
		v426 = v323
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = int32(0)
	goto L95
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+96)) = int32(0)
	if v426 != 0 {
		goto L95
	} else {
		goto L114
	}
L100:
	;
	v328 = v178 & int32(3)
	if base.Ui32(v178) < base.Ui32(int32(4)) {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	if v328 != 0 {
		goto L108
	} else {
		goto L109
	}
L102:
	;
	v369 = v322
	v370 = v221
	v372 = int32(0)
	goto L101
L103:
	;
	goto L104
L104:
	;
	v335 = v322
	v336 = v221
	v338 = int32(0)
	v344 = int32(0)
	goto L105
L105:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336))))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+1)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+1)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+2)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+2)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+3)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+3)))
	v361 = v338 | (v346 ^ v347) | (v350 ^ v351) | (v354 ^ v355) | (v358 ^ v359)
	v362 = int32(4)
	v363 = v336 + v362
	v365 = v335 + v362
	v367 = v344 + v362
	if v367 != v178&int32(-4) {
		v335 = v365
		v336 = v363
		v338 = v361
		v344 = v367
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v369 = v365
	v370 = v363
	v372 = v361
	goto L101
L107:
	;
	goto L106
L108:
	;
	v380 = v369
	v381 = v370
	v383 = v372
	v388 = v323
	goto L111
L109:
	;
	v405 = v372
	goto L110
L110:
	;
	v426 = base.B2i32(v405 != int32(0))
	goto L99
L111:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
	v394 = v383 | (v391 ^ v392)
	v395 = int32(1)
	v400 = v388 + v395
	if v400 != v328 {
		v380 = v380 + v395
		v381 = v381 + v395
		v383 = v394
		v388 = v400
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v405 = v394
	goto L110
L113:
	;
	goto L112
L114:
	;
	v431 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v431 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v219
	F_errmsg_internal(m, int32(451336), v225)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v445 = F_kill(m, int32(0)-v219, int32(2))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	F_errfinish(m, int32(478339), int32(772), int32(74257))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	goto L68
L122:
	;
	if v462 == int32(0) {
		goto L68
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v219
	F_errmsg(m, int32(451293), v225+int32(16))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(478339), int32(789), int32(74257))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L68
L126:
	;
	v492 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)) = uint8(v492)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, _consts[647])))
	if v495 != int32(1) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	goto L134
L128:
	;
	v500 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v500 == int32(0) {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(430371), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(476542), int32(612), int32(101498))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L127
L133:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v556 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	goto L144
L134:
	;
	v527 = F_secure_write(m, l0, v12+int32(-16), int32(1))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L136
	}
L135:
	;
	v535 = int32(-1)
	v538 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L139
	}
L136:
	;
	if v527 == int32(1) {
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v532 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v532 == int32(27) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	if v538 == int32(0) {
		v1377 = v535
		goto L8
	} else {
		goto L140
	}
L140:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errmsg(m, int32(282217), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(476542), int32(621), int32(101498))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v1377 = v535
	goto L8
L144:
	;
	if int32(0) < v554-v556 {
		goto L7
	} else {
		goto L145
	}
L145:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v564 = F_ProcessStartupPacket(m, l0, int32(1), base.B2i32(v561 == int32(83)))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v1377 = v564
	goto L8
L147:
	;
	v566 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)) = uint8(v566)
	v569 = int32(*(*uint8)(unsafe.Add(mBase, _consts[647])))
	if v569 != int32(1) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	goto L155
L149:
	;
	v574 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	if v574 == int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(430391), int32(0))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(476542), int32(666), int32(101498))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	goto L148
L154:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _consts[372]))
	v630 = *(*int32)(unsafe.Add(mBase, _consts[370]))
	goto L165
L155:
	;
	v601 = F_secure_write(m, l0, v12+int32(-16), int32(1))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v609 = int32(-1)
	v612 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L160
	}
L157:
	;
	if v601 == int32(1) {
		goto L154
	} else {
		goto L158
	}
L158:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v606 == int32(27) {
		goto L155
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	if v612 == int32(0) {
		v1377 = v609
		goto L8
	} else {
		goto L161
	}
L161:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errmsg(m, int32(282261), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(476542), int32(675), int32(101498))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v1377 = v609
	goto L8
L165:
	;
	if int32(0) < v628-v630 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+48)))
	v638 = F_ProcessStartupPacket(m, l0, base.B2i32(v634 == int32(71)), int32(1))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v1377 = v638
	goto L8
L168:
	;
	v645 = v642
	goto L170
L169:
	;
	v645 = v147
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v645
	if base.Ui32(v147-int32(262144)) <= base.Ui32(int32(-65537)) {
		goto L5
	} else {
		goto L171
	}
L171:
	;
	v651 = int32(4449520)
	v652 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v655 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v655
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = int32(0)
	v659 = int32(4)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v660 < int32(5) {
		v1080 = v659
		v1081 = v640
		v1086 = v660
		goto L172
	} else {
		goto L173
	}
L172:
	;
	if v1080 != v1086-int32(1) {
		goto L4
	} else {
		goto L312
	}
L173:
	;
	v664 = v659
	v665 = v640
	v670 = v660
	goto L174
L174:
	;
	v674 = v664 + v99
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v675 == int32(0) {
		v1080 = v664
		v1081 = v665
		v1086 = v670
		goto L172
	} else {
		goto L176
	}
L175:
	;
	v1080 = v1076
	v1081 = v1015
	v1086 = v1077
	goto L172
L176:
	;
	if v674&int32(3) == int32(0) {
		v701 = v674
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v737 = v734 + v664 + int32(1)
	if v670 <= v737 {
		v1080 = v664
		v1081 = v665
		v1086 = v670
		goto L172
	} else {
		goto L194
	}
L178:
	;
	v734 = v726 - v674
	goto L177
L179:
	;
	v705 = v701
	goto L188
L180:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v685 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v734 = int32(0)
	goto L177
L182:
	;
	goto L183
L183:
	;
	v690 = v674
	goto L184
L184:
	;
	v694 = v690 + int32(1)
	if v694&int32(3) == int32(0) {
		v701 = v694
		goto L179
	} else {
		goto L186
	}
L185:
	;
	v726 = v694
	goto L178
L186:
	;
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694))))
	if v699 != 0 {
		v690 = v694
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v714 = int32(-2139062144)
	if (int32(16843008)-v711|v711)&v714 == v714 {
		v705 = v705 + int32(4)
		goto L188
	} else {
		goto L190
	}
L189:
	;
	v720 = v705
	goto L191
L190:
	;
	goto L189
L191:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v720))))
	if v724 != 0 {
		v720 = v720 + int32(1)
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v726 = v720
	goto L178
L193:
	;
	goto L192
L194:
	;
	v739 = v99 + v737
	v740 = int32(348134)
	v743 = int32(*(*uint8)(unsafe.Add(mBase, _consts[649])))
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v744 == int32(0) {
		v763 = v743
		v764 = v744
		goto L197
	} else {
		goto L198
	}
L195:
	;
	if v739&int32(3) == int32(0) {
		v1040 = v739
		goto L296
	} else {
		goto L297
	}
L196:
	;
	if v764-v763 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L197:
	;
	goto L196
L198:
	;
	if v743 != v744 {
		v763 = v743
		v764 = v744
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v748 = v674
	v749 = v740
	goto L200
L200:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+1)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v748)+1)))
	if v753 == int32(0) {
		v763 = v752
		v764 = v753
		goto L197
	} else {
		goto L202
	}
L201:
	;
	v763 = v752
	v764 = v753
	goto L197
L202:
	;
	v756 = int32(1)
	if v752 == v753 {
		v748 = v748 + v756
		v749 = v749 + v756
		goto L200
	} else {
		goto L203
	}
L203:
	;
	goto L201
L204:
	;
	v768 = F_pstrdup(m, v739)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	v771 = int32(208663)
	v774 = int32(*(*uint8)(unsafe.Add(mBase, _consts[650])))
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v775 == int32(0) {
		v794 = v774
		v795 = v775
		goto L209
	} else {
		goto L210
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v768
	v1015 = v665
	goto L195
L208:
	;
	if v795-v794 == int32(0) {
		goto L216
	} else {
		goto L217
	}
L209:
	;
	goto L208
L210:
	;
	if v774 != v775 {
		v794 = v774
		v795 = v775
		goto L209
	} else {
		goto L211
	}
L211:
	;
	v779 = v674
	v780 = v771
	goto L212
L212:
	;
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780)+1)))
	v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779)+1)))
	if v784 == int32(0) {
		v794 = v783
		v795 = v784
		goto L209
	} else {
		goto L214
	}
L213:
	;
	v794 = v783
	v795 = v784
	goto L209
L214:
	;
	v787 = int32(1)
	if v783 == v784 {
		v779 = v779 + v787
		v780 = v780 + v787
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v799 = F_pstrdup(m, v739)
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v802 = int32(130199)
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _consts[651])))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v806 == int32(0) {
		v825 = v805
		v826 = v806
		goto L221
	} else {
		goto L222
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+364)) = v799
	v1015 = v665
	goto L195
L220:
	;
	if v826-v825 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L221:
	;
	goto L220
L222:
	;
	if v805 != v806 {
		v825 = v805
		v826 = v806
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v810 = v674
	v811 = v802
	goto L224
L224:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+1)))
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810)+1)))
	if v815 == int32(0) {
		v825 = v814
		v826 = v815
		goto L221
	} else {
		goto L226
	}
L225:
	;
	v825 = v814
	v826 = v815
	goto L221
L226:
	;
	v818 = int32(1)
	if v814 == v815 {
		v810 = v810 + v818
		v811 = v811 + v818
		goto L224
	} else {
		goto L227
	}
L227:
	;
	goto L225
L228:
	;
	v830 = F_pstrdup(m, v739)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v833 = int32(256058)
	v836 = int32(*(*uint8)(unsafe.Add(mBase, _consts[652])))
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v837 == int32(0) {
		v856 = v836
		v857 = v837
		goto L233
	} else {
		goto L234
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v830
	v1015 = v665
	goto L195
L232:
	;
	if v857-v856 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L233:
	;
	goto L232
L234:
	;
	if v836 != v837 {
		v856 = v836
		v857 = v837
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v841 = v674
	v842 = v833
	goto L236
L236:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842)+1)))
	v846 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+1)))
	if v846 == int32(0) {
		v856 = v845
		v857 = v846
		goto L233
	} else {
		goto L238
	}
L237:
	;
	v856 = v845
	v857 = v846
	goto L233
L238:
	;
	v849 = int32(1)
	if v845 == v846 {
		v841 = v841 + v849
		v842 = v842 + v849
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v861 = int32(348134)
	v864 = int32(*(*uint8)(unsafe.Add(mBase, _consts[649])))
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	if v865 == int32(0) {
		v884 = v864
		v885 = v865
		goto L244
	} else {
		goto L245
	}
L241:
	;
	goto L242
L242:
	;
	v922 = int32(618410)
	goto L263
L243:
	;
	if v885-v884 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L244:
	;
	goto L243
L245:
	;
	if v864 != v865 {
		v884 = v864
		v885 = v865
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v869 = v739
	v870 = v861
	goto L247
L247:
	;
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+1)))
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v869)+1)))
	if v874 == int32(0) {
		v884 = v873
		v885 = v874
		goto L244
	} else {
		goto L249
	}
L248:
	;
	v884 = v873
	v885 = v874
	goto L244
L249:
	;
	v877 = int32(1)
	if v873 == v874 {
		v869 = v869 + v877
		v870 = v870 + v877
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v890 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[653])) = uint8(v890)
	*(*uint8)(unsafe.Add(mBase, _consts[654])) = uint8(v890)
	v1015 = v665
	goto L195
L252:
	;
	goto L253
L253:
	;
	v896 = F_strlen(m, v739)
	mBase = m.M
	v897 = F_parse_bool_with_len(m, v739, v896, int32(4359936))
	mBase = m.M
	goto L254
L254:
	;
	if v897 != 0 {
		v1015 = v665
		goto L195
	} else {
		goto L255
	}
L255:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(256058)
	F_errmsg(m, int32(690044), v12+int32(-32))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	F_errhint(m, int32(631506), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	F_errfinish(m, int32(476542), int32(784), int32(101498))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
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
	if v959-v960 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L263:
	;
	goto L264
L264:
	;
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v929 != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v930 = v674
	v931 = v922
	v932 = int32(5)
	v933 = v929
	goto L269
L266:
	;
	v955 = v922
	v959 = int32(0)
	goto L267
L267:
	;
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955))))
	goto L261
L268:
	;
	v955 = v950
	v959 = v952
	goto L267
L269:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if v933 != v935 {
		v950 = v931
		v952 = v933
		goto L268
	} else {
		goto L271
	}
L270:
	;
	v950 = v944
	v952 = int32(0)
	goto L268
L271:
	;
	if v935 == int32(0) {
		v950 = v931
		v952 = v933
		goto L268
	} else {
		goto L272
	}
L272:
	;
	v940 = v932 - int32(1)
	if v940 == int32(0) {
		v950 = v931
		v952 = v933
		goto L268
	} else {
		goto L273
	}
L273:
	;
	v943 = int32(1)
	v944 = v931 + v943
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+1)))
	if v945 != 0 {
		v930 = v930 + v943
		v931 = v944
		v932 = v940
		v933 = v945
		goto L269
	} else {
		goto L274
	}
L274:
	;
	goto L270
L275:
	;
	v970 = F_pstrdup(m, v674)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
	v975 = F_pstrdup(m, v674)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	v972 = F_lappend(m, v665, v970)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	v1015 = v972
	goto L195
L280:
	;
	v977 = F_lappend(m, v974, v975)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v977
	v980 = F_pstrdup(m, v739)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	v982 = F_lappend(m, v977, v980)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v982
	v985 = int32(363469)
	v988 = int32(*(*uint8)(unsafe.Add(mBase, _consts[655])))
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v674))))
	if v989 == int32(0) {
		v1008 = v988
		v1009 = v989
		goto L285
	} else {
		goto L286
	}
L284:
	;
	if v1009-v1008 != 0 {
		v1015 = v665
		goto L195
	} else {
		goto L292
	}
L285:
	;
	goto L284
L286:
	;
	if v988 != v989 {
		v1008 = v988
		v1009 = v989
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v993 = v674
	v994 = v985
	goto L288
L288:
	;
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994)+1)))
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993)+1)))
	if v998 == int32(0) {
		v1008 = v997
		v1009 = v998
		goto L285
	} else {
		goto L290
	}
L289:
	;
	v1008 = v997
	v1009 = v998
	goto L285
L290:
	;
	v1001 = int32(1)
	if v997 == v998 {
		v993 = v993 + v1001
		v994 = v994 + v1001
		goto L288
	} else {
		goto L291
	}
L291:
	;
	goto L289
L292:
	;
	v1012 = F_pg_clean_ascii(m, v739, int32(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+376)) = v1012
	v1015 = v665
	goto L195
L294:
	;
	v1076 = v1073 + v737 + int32(1)
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v1076 < v1077 {
		v664 = v1076
		v665 = v1015
		v670 = v1077
		goto L174
	} else {
		goto L311
	}
L295:
	;
	v1073 = v1065 - v739
	goto L294
L296:
	;
	v1044 = v1040
	goto L305
L297:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739))))
	if v1024 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1073 = int32(0)
	goto L294
L299:
	;
	goto L300
L300:
	;
	v1029 = v739
	goto L301
L301:
	;
	v1033 = v1029 + int32(1)
	if v1033&int32(3) == int32(0) {
		v1040 = v1033
		goto L296
	} else {
		goto L303
	}
L302:
	;
	v1065 = v1033
	goto L295
L303:
	;
	v1038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1033))))
	if v1038 != 0 {
		v1029 = v1033
		goto L301
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1044)))
	v1053 = int32(-2139062144)
	if (int32(16843008)-v1050|v1050)&v1053 == v1053 {
		v1044 = v1044 + int32(4)
		goto L305
	} else {
		goto L307
	}
L306:
	;
	v1059 = v1044
	goto L308
L307:
	;
	goto L306
L308:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1059))))
	if v1063 != 0 {
		v1059 = v1059 + int32(1)
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v1065 = v1059
	goto L295
L310:
	;
	goto L309
L311:
	;
	goto L175
L312:
	;
	v1093 = int32(0)
	if base.B2i32(v1081 == v1093)&base.B2i32(base.Ui32(v147&int32(65535)) <= base.Ui32(int32(2))) == v1093 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	F_pq_beginmessage(m, v12+int32(-16), int32(118))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1222 == int32(0) {
		goto L3
	} else {
		goto L330
	}
L316:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	F_enlargeStringInfo(m, v12+int32(-16), int32(4))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v1118 = int32(24)
	v1120 = int32(65280)
	v1122 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1115+v1116))) = v1109<<(uint(v1118)%32) | v1109&v1120<<(uint(v1122)%32) | (int32(base.Ui32(v1109)>>(uint(v1122)%32))&v1120 | int32(base.Ui32(v1109)>>(uint(v1118)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v1115 + int32(4)
	if v1081 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	v1138 = v1137
	goto L320
L319:
	;
	v1138 = int32(0)
	goto L320
L320:
	;
	F_enlargeStringInfo(m, v12+int32(-16), int32(4))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v1147 = int32(24)
	v1149 = int32(65280)
	v1151 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v1144+v1145))) = v1138<<(uint(v1147)%32) | v1138&v1149<<(uint(v1151)%32) | (int32(base.Ui32(v1138)>>(uint(v1151)%32))&v1149 | int32(base.Ui32(v1138)>>(uint(v1147)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v1144 + int32(4)
	if v1081 == int32(0) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	F_pq_endmessage(m, v12+int32(-16))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L329
	}
L323:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if v1168 <= int32(0) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1175 = int32(0)
	goto L325
L325:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+12))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1185+v1175<<(uint(int32(2))%32))))
	F_pq_sendstring(m, v12+int32(-16), v1189)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L327
	}
L326:
	;
	goto L322
L327:
	;
	v1193 = v1175 + int32(1)
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if v1193 < v1194 {
		v1175 = v1193
		goto L325
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	goto L315
L330:
	;
	v1225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1222))))
	if v1225 == int32(0) {
		goto L3
	} else {
		goto L331
	}
L331:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	if v1228 != 0 {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	if v1233&int32(3) == int32(0) {
		v1257 = v1233
		goto L340
	} else {
		goto L341
	}
L333:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228))))
	if v1229 != 0 {
		v1233 = v1228
		goto L332
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1230 = F_pstrdup(m, v1222)
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L337
	}
L336:
	;
	goto L335
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+360)) = v1230
	v1233 = v1230
	goto L332
L338:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v1290) {
		goto L355
	} else {
		goto L356
	}
L339:
	;
	v1290 = v1282 - v1233
	goto L338
L340:
	;
	v1261 = v1257
	goto L349
L341:
	;
	v1241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1233))))
	if v1241 == int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1290 = int32(0)
	goto L338
L343:
	;
	goto L344
L344:
	;
	v1246 = v1233
	goto L345
L345:
	;
	v1250 = v1246 + int32(1)
	if v1250&int32(3) == int32(0) {
		v1257 = v1250
		goto L340
	} else {
		goto L347
	}
L346:
	;
	v1282 = v1250
	goto L339
L347:
	;
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1250))))
	if v1255 != 0 {
		v1246 = v1250
		goto L345
	} else {
		goto L348
	}
L348:
	;
	goto L346
L349:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1261)))
	v1270 = int32(-2139062144)
	if (int32(16843008)-v1267|v1267)&v1270 == v1270 {
		v1261 = v1261 + int32(4)
		goto L349
	} else {
		goto L351
	}
L350:
	;
	v1276 = v1261
	goto L352
L351:
	;
	goto L350
L352:
	;
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276))))
	if v1280 != 0 {
		v1276 = v1276 + int32(1)
		goto L352
	} else {
		goto L354
	}
L353:
	;
	v1282 = v1276
	goto L339
L354:
	;
	goto L353
L355:
	;
	v1293 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1233)+63)) = uint8(v1293)
	goto L357
L356:
	;
	goto L357
L357:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v1295&int32(3) == int32(0) {
		v1319 = v1295
		goto L360
	} else {
		goto L361
	}
L358:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v1352) {
		goto L375
	} else {
		goto L376
	}
L359:
	;
	v1352 = v1344 - v1295
	goto L358
L360:
	;
	v1323 = v1319
	goto L369
L361:
	;
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295))))
	if v1303 == int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1352 = int32(0)
	goto L358
L363:
	;
	goto L364
L364:
	;
	v1308 = v1295
	goto L365
L365:
	;
	v1312 = v1308 + int32(1)
	if v1312&int32(3) == int32(0) {
		v1319 = v1312
		goto L360
	} else {
		goto L367
	}
L366:
	;
	v1344 = v1312
	goto L359
L367:
	;
	v1317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1312))))
	if v1317 != 0 {
		v1308 = v1312
		goto L365
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1323)))
	v1332 = int32(-2139062144)
	if (int32(16843008)-v1329|v1329)&v1332 == v1332 {
		v1323 = v1323 + int32(4)
		goto L369
	} else {
		goto L371
	}
L370:
	;
	v1338 = v1323
	goto L372
L371:
	;
	goto L370
L372:
	;
	v1342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338))))
	if v1342 != 0 {
		v1338 = v1338 + int32(1)
		goto L372
	} else {
		goto L374
	}
L373:
	;
	v1344 = v1338
	goto L359
L374:
	;
	goto L373
L375:
	;
	v1355 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1295)+63)) = uint8(v1355)
	goto L377
L376:
	;
	goto L377
L377:
	;
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, _consts[654])))
	if v1362 != 0 {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v1363 = int32(6)
	goto L380
L379:
	;
	v1363 = int32(1)
	goto L380
L380:
	;
	*(*int32)(unsafe.Add(mBase, _consts[172])) = v1363
	if v1362 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v652
	v1377 = int32(0)
	goto L8
L382:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, _consts[653])))
	if v1368 != 0 {
		goto L381
	} else {
		goto L383
	}
L383:
	;
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	v1370 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1369))) = uint8(v1370)
	goto L381
L384:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	F_errmsg(m, int32(74191), int32(0))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	F_errdetail(m, int32(587444), int32(0))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(476542), int32(640), int32(101498))
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L390
	}
L390:
	;
	F_errmsg(m, int32(73988), int32(0))
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L391
	}
L391:
	;
	F_errdetail(m, int32(587444), int32(0))
	mBase = m.M
	v1423 = m.ExcPending
	if v1423 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(476542), int32(694), int32(101498))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(12884901891)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v147 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(base.Ui32(v147) >> (uint(int32(16)) % 32))
	F_errmsg(m, int32(37240), v14)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(476542), int32(725), int32(101498))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L399
	}
L399:
	;
	F_errmsg(m, int32(334413), int32(0))
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L1
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(476542), int32(825), int32(101498))
	mBase = m.M
	v1469 = m.ExcPending
	if v1469 != 0 {
		goto L1
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errcode(m, int32(514))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	F_errmsg(m, int32(101272), int32(0))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L1
	} else {
		goto L404
	}
L404:
	;
	F_errfinish(m, int32(476542), int32(842), int32(101498))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_startup_progress_timeout_expired(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v31 int64
	_ = v31
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	if v11 != 0 {
		v15 = m.G0
		v16 = int32(16)
		v17 = v15 - v16
		m.G0 = v17
		F___gettimeofday(m, v17)
		mBase = m.M
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+8)))
		m.G0 = v17 + v16
		v31 = *(*int64)(unsafe.Add(mBase, _consts[476]))
		v38 = v21 + v20*int64(1000000) - int64(946684800000000) - v31
		if v38 <= int64(0) {
			v50 = int32(0)
			v51 = int32(0)
		} else {
			v42 = int64(1000000)
			v43 = base.I64_div_u_s(v38, v42)
			v50 = base.I32_wrap_i64(v43)
			v51 = base.I32_wrap_i64(v38 - v43*v42)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(12)))) = v50
		*(*int32)(unsafe.Add(mBase, uint32(v8+int32(8)))) = v51
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v54
		v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v56
		*(*int32)(unsafe.Add(mBase, _consts[475])) = int32(0)
	} else {
	}
	m.G0 = v8 + int32(16)
	return base.B2i32(v11 != int32(0))
}
