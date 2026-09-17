package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F___emscripten_stdout_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	return int64(0)
}
func F_emscripten_builtin_malloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
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
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var __phi185 int32
	_ = __phi185
	var v189 int32
	_ = v189
	var __phi189 int32
	_ = __phi189
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var __phi384 int32
	_ = __phi384
	var v385 int32
	_ = v385
	var __phi385 int32
	_ = __phi385
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v835 int32
	_ = v835
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v874 int32
	_ = v874
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v890 int64
	_ = v890
	var v892 int32
	_ = v892
	var v893 int64
	_ = v893
	var v908 int32
	_ = v908
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1018 int32
	_ = v1018
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1046 int32
	_ = v1046
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1065 int32
	_ = v1065
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var __phi1202 int32
	_ = __phi1202
	var v1207 int32
	_ = v1207
	var __phi1207 int32
	_ = __phi1207
	var v1216 int32
	_ = v1216
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1385 int32
	_ = v1385
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1413 int32
	_ = v1413
	var v1434 int32
	_ = v1434
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1493 int32
	_ = v1493
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1579 int32
	_ = v1579
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1607 int32
	_ = v1607
	var v1628 int32
	_ = v1628
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1655 int32
	_ = v1655
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1686 int32
	_ = v1686
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1727 int32
	_ = v1727
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if base.Ui32(l0) <= base.Ui32(int32(244)) {
		goto L11
	} else {
		goto L12
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v1727
L2:
	;
	if v168 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L3:
	;
	if v367 == int32(0) {
		v1478 = v212
		goto L314
	} else {
		goto L315
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v813))) = v664
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v813)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v813)+4)) = v1118 + v668
	v1121 = int32(-8)
	v1123 = int32(7)
	v1125 = v664 + (v1121-v664)&v1123
	*(*int32)(unsafe.Add(mBase, uint32(v1125)+4)) = v406 | int32(3)
	v1133 = v825 + (v1121-v825)&v1123
	v1134 = v406 + v1125
	v1135 = v1133 - v1134
	v1137 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0]))
	if v1137 == v1133 {
		goto L248
	} else {
		goto L249
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[1])) = int32(48)
	v1727 = int32(0)
	goto L1
L6:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2]))
	if base.Ui32(v1083) <= base.Ui32(v406) {
		goto L5
	} else {
		goto L246
	}
L7:
	;
	v807 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[3]))
	if base.Ui32(v664) < base.Ui32(v807) {
		goto L198
	} else {
		goto L199
	}
L8:
	;
	v1434 = int32(0)
	goto L3
L9:
	;
	v1628 = int32(0)
	goto L2
L10:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4]))
	if base.Ui32(v406) <= base.Ui32(v416) {
		goto L120
	} else {
		goto L121
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5]))
	v22 = int32(11)
	if base.Ui32(l0) < base.Ui32(v22) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(int32(-65)) < base.Ui32(l0) {
		v406 = int32(-1)
		goto L10
	} else {
		goto L64
	}
L14:
	;
	v28 = int32(16)
	goto L16
L15:
	;
	v28 = (l0 + v22) & int32(504)
	goto L16
L16:
	;
	v29 = int32(3)
	v30 = int32(base.Ui32(v28) >> (uint(v29) % 32))
	v31 = int32(base.Ui32(v20) >> (uint(v30) % 32))
	if v31&v29 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = (v31^int32(-1))&int32(1) + v30
	v40 = v38 << (uint(int32(3)) % 32)
	v42 = v40 + int32(_a_F_emscripten_builtin_malloc_0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v42 == v44 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4]))
	if base.Ui32(v28) <= base.Ui32(v64) {
		v406 = v28
		goto L10
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v40 | int32(3)
	v58 = v40 + v43
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59 | int32(1)
	v1727 = v43 + int32(8)
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v20 & base.I32_rotl(int32(-2), v38)
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v44
	goto L20
L24:
	;
	if v31 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v67 = int32(2) << (uint(v30) % 32)
	v73 = base.I32_ctz((v67 | (int32(0) - v67)) & (v31 << (uint(v30) % 32)))
	v75 = v73 << (uint(int32(3)) % 32)
	v77 = v75 + int32(_a_F_emscripten_builtin_malloc_0)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v77 == v79 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	goto L27
L27:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7]))
	if v131 == int32(0) {
		v406 = v28
		goto L10
	} else {
		goto L39
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v28 | int32(3)
	v92 = v78 + v28
	v93 = v75 - v28
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v93 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v78+v75))) = v93
	if v64 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v84 = v20 & base.I32_rotl(int32(-2), v73)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v84
	v88 = v84
	goto L28
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v75)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v79
	v88 = v20
	goto L28
L32:
	;
	v100 = v64 & int32(-8)
	v102 = v100 + int32(_a_F_emscripten_builtin_malloc_0)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8]))
	v108 = int32(1) << (uint(int32(base.Ui32(v64)>>(uint(int32(3))%32))) % 32)
	if v88&v108 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4])) = v93
	v1727 = v78 + int32(8)
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v104)+12)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = v116
	goto L34
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v108 | v88
	v116 = v102
	goto L35
L37:
	;
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v100)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v116 = v115
	goto L35
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v131)<<(uint(int32(2))%32))+uint32(_c_F_emscripten_builtin_malloc[9])))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	v143 = v137
	v144 = v137
	v146 = v138&int32(-8) - v28
	goto L40
L40:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	if v154 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144)+24))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	if v144 != v169 {
		goto L53
	} else {
		goto L54
	}
L42:
	;
	goto L41
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v143)+20))
	if v157 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	v160 = v154
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v164 = v161&int32(-8) - v28
	v165 = base.B2i32(base.Ui32(v164) < base.Ui32(v146))
	if base.Ui32(v164) < base.Ui32(v146) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v160 = v157
	goto L45
L47:
	;
	v166 = v164
	goto L49
L48:
	;
	v166 = v146
	goto L49
L49:
	;
	if base.Ui32(v164) < base.Ui32(v146) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v167 = v160
	goto L52
L51:
	;
	v167 = v144
	goto L52
L52:
	;
	v143 = v160
	v144 = v167
	v146 = v166
	goto L40
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+12)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v169)+8)) = v171
	v1628 = v169
	goto L2
L54:
	;
	goto L55
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v174 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v182 = v174
	v183 = v144 + int32(20)
	goto L58
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	if v177 == int32(0) {
		goto L9
	} else {
		goto L59
	}
L58:
	;
	__phi185 = v182
	__phi189 = v183
	v185 = __phi185
	v189 = __phi189
	goto L60
L59:
	;
	v182 = v177
	v183 = v144 + int32(16)
	goto L58
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v185)+20))
	if v198 != 0 {
		__phi185 = v198
		__phi189 = v185 + int32(20)
		v185 = __phi185
		v189 = __phi189
		goto L60
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = int32(0)
	v1628 = v185
	goto L2
L62:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	if v201 != 0 {
		__phi185 = v201
		__phi189 = v185 + int32(16)
		v185 = __phi185
		v189 = __phi189
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v208 = l0 + int32(11)
	v210 = v208 & int32(-8)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7]))
	if v212 == int32(0) {
		v406 = v210
		goto L10
	} else {
		goto L65
	}
L65:
	;
	v217 = int32(0) - v210
	if base.Ui32(l0) <= base.Ui32(int32(16777204)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v223 = base.I32_clz(int32(base.Ui32(v208) >> (uint(int32(8)) % 32)))
	v226 = int32(1)
	v234 = int32(base.Ui32(v210)>>(uint(int32(38)-v223)%32))&v226 - v223<<(uint(v226)%32) + int32(62)
	goto L68
L67:
	;
	v234 = int32(31)
	goto L68
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234<<(uint(int32(2))%32))+uint32(_c_F_emscripten_builtin_malloc[9])))
	if v237 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	if v354 == int32(0) {
		v406 = v210
		goto L10
	} else {
		goto L107
	}
L70:
	;
	v327 = v315
	v331 = v319
	v332 = v320
	goto L95
L71:
	;
	if v283|v288 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L72:
	;
	v283 = int32(0)
	v287 = v217
	v288 = v2
	goto L71
L73:
	;
	goto L74
L74:
	;
	v241 = int32(0)
	if v234 != int32(31) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v249 = int32(25) - int32(base.Ui32(v234)>>(uint(int32(1))%32))
	goto L77
L76:
	;
	v249 = v241
	goto L77
L77:
	;
	v251 = v241
	v252 = v237
	v253 = v210 << (uint(v249) % 32)
	v255 = v217
	v256 = v2
	goto L78
L78:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v266 = v263&int32(-8) - v210
	if base.Ui32(v255) <= base.Ui32(v266) {
		v269 = v255
		v270 = v256
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v283 = v280
	v287 = v269
	v288 = v270
	goto L71
L80:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v252)+20))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v252+int32(base.Ui32(v253)>>(uint(int32(29))%32))&int32(4))+16))
	if v271 == v277 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	if v266 != 0 {
		v269 = v266
		v270 = v252
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v315 = v252
	v319 = int32(0)
	v320 = v252
	goto L70
L83:
	;
	v279 = v251
	goto L85
L84:
	;
	v279 = v271
	goto L85
L85:
	;
	if v271 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v280 = v279
	goto L88
L87:
	;
	v280 = v251
	goto L88
L88:
	;
	if v277 != 0 {
		v251 = v280
		v252 = v277
		v253 = v253 << (uint(int32(1)) % 32)
		v255 = v269
		v256 = v270
		goto L78
	} else {
		goto L89
	}
L89:
	;
	goto L79
L90:
	;
	v298 = int32(0)
	v300 = int32(2) << (uint(v234) % 32)
	v304 = (v300 | (v298 - v300)) & v212
	if v304 == v298 {
		v406 = v210
		goto L10
	} else {
		goto L93
	}
L91:
	;
	v311 = v283
	v312 = v288
	goto L92
L92:
	;
	if v311 == int32(0) {
		v353 = v287
		v354 = v312
		goto L69
	} else {
		goto L94
	}
L93:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_ctz(v304)<<(uint(int32(2))%32))+uint32(_c_F_emscripten_builtin_malloc[9])))
	v311 = v310
	v312 = v298
	goto L92
L94:
	;
	v315 = v311
	v319 = v287
	v320 = v312
	goto L70
L95:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v342 = v339&int32(-8) - v210
	v343 = base.B2i32(base.Ui32(v342) < base.Ui32(v331))
	if base.Ui32(v342) < base.Ui32(v331) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v353 = v344
	v354 = v345
	goto L69
L97:
	;
	v344 = v342
	goto L99
L98:
	;
	v344 = v331
	goto L99
L99:
	;
	if base.Ui32(v342) < base.Ui32(v331) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v345 = v327
	goto L102
L101:
	;
	v345 = v332
	goto L102
L102:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v327)+16))
	if v346 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v348 = v346
	goto L105
L104:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v327)+20))
	v348 = v347
	goto L105
L105:
	;
	if v348 != 0 {
		v327 = v348
		v331 = v344
		v332 = v345
		goto L95
	} else {
		goto L106
	}
L106:
	;
	goto L96
L107:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4]))
	if base.Ui32(v364-v210) <= base.Ui32(v353) {
		v406 = v210
		goto L10
	} else {
		goto L108
	}
L108:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v354)+24))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	if v354 != v368 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v354)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v370)+12)) = v368
	*(*int32)(unsafe.Add(mBase, uint32(v368)+8)) = v370
	v1434 = v368
	goto L3
L110:
	;
	goto L111
L111:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	if v373 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v381 = v373
	v382 = v354 + int32(20)
	goto L114
L113:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	if v376 == int32(0) {
		goto L8
	} else {
		goto L115
	}
L114:
	;
	__phi384 = v381
	__phi385 = v382
	v384 = __phi384
	v385 = __phi385
	goto L116
L115:
	;
	v381 = v376
	v382 = v354 + int32(16)
	goto L114
L116:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v384)+20))
	if v397 != 0 {
		__phi384 = v397
		__phi385 = v384 + int32(20)
		v384 = __phi384
		v385 = __phi385
		goto L116
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = int32(0)
	v1434 = v384
	goto L3
L118:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v384)+16))
	if v400 != 0 {
		__phi384 = v400
		__phi385 = v384 + int32(16)
		v384 = __phi384
		v385 = __phi385
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8]))
	v420 = v416 - v406
	if base.Ui32(int32(16)) <= base.Ui32(v420) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2]))
	if base.Ui32(v406) < base.Ui32(v451) {
		goto L127
	} else {
		goto L128
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4])) = v442
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8])) = v443
	v1727 = v419 + int32(8)
	goto L1
L124:
	;
	v423 = v419 + v406
	*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v420 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v419+v416))) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = v406 | int32(3)
	v442 = v420
	v443 = v423
	goto L123
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v419)+4)) = v416 | int32(3)
	v435 = v419 + v416
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v435)+4)) = v436 | int32(1)
	v440 = int32(0)
	v442 = v440
	v443 = v440
	goto L123
L127:
	;
	v454 = v451 - v406
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v454
	v456 = int32(_a_F_emscripten_builtin_malloc_1)
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0]))
	v459 = v458 + v406
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v459
	*(*int32)(unsafe.Add(mBase, uint32(v459)+4)) = v454 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v458)+4)) = v406 | int32(3)
	v1727 = v458 + int32(8)
	goto L1
L128:
	;
	goto L129
L129:
	;
	v469 = int32(0)
	v471 = v406 + int32(47)
	v473 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[10]))
	if v473 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v498 = v471 + v497
	v500 = int32(0) - v497
	v501 = v498 & v500
	if base.Ui32(v501) <= base.Ui32(v406) {
		v1727 = v469
		goto L1
	} else {
		goto L134
	}
L131:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[11]))
	v497 = v475
	goto L130
L132:
	;
	goto L133
L133:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[12])) = int64(-1)
	*(*int64)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[13])) = int64(17592186048512)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[10])) = (v15+int32(12))&int32(-16) ^ int32(1431655768)
	v491 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[14])) = v491
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[15])) = v491
	v497 = int32(_a_F_emscripten_builtin_malloc_2)
	goto L130
L134:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[16]))
	if v504 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[17]))
	v507 = v506 + v501
	if base.B2i32(base.Ui32(v507) <= base.Ui32(v506))|base.B2i32(base.Ui32(v504) < base.Ui32(v507)) != 0 {
		v1727 = v469
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v514 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[15])))
	if v514&int32(4) == int32(0) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	goto L137
L139:
	;
	v674 = int32(_a_F_emscripten_builtin_malloc_3)
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[17]))
	v677 = v676 + v668
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[17])) = v677
	v680 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[18]))
	if base.Ui32(v680) < base.Ui32(v677) {
		goto L176
	} else {
		goto L177
	}
L140:
	;
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0]))
	if v520 != 0 {
		goto L147
	} else {
		goto L148
	}
L141:
	;
	goto L142
L142:
	;
	v648 = F_sbrk(m, v501)
	mBase = m.M
	v649 = int32(-1)
	v652 = F_sbrk(m, int32(0))
	mBase = m.M
	if base.B2i32(v648 == v649)|base.B2i32(v652 == v649)|base.B2i32(base.Ui32(v652) <= base.Ui32(v648)) != 0 {
		goto L5
	} else {
		goto L174
	}
L143:
	;
	v630 = int32(_a_F_emscripten_builtin_malloc_4)
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[15]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[15])) = v632 | int32(4)
	goto L142
L144:
	;
	if v583 != int32(-1) {
		v664 = v583
		v668 = v582
		goto L139
	} else {
		goto L173
	}
L145:
	;
	if v588 == int32(-1) {
		goto L143
	} else {
		goto L168
	}
L146:
	;
	v582 = (v498 - v451) & v500
	v583 = F_sbrk(m, v582)
	mBase = m.M
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if v583 == v584+v585 {
		goto L144
	} else {
		goto L167
	}
L147:
	;
	v522 = int32(_a_F_emscripten_builtin_malloc_5)
	goto L150
L148:
	;
	goto L149
L149:
	;
	v553 = F_sbrk(m, int32(0))
	mBase = m.M
	if v553 == int32(-1) {
		goto L143
	} else {
		goto L157
	}
L150:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v522)))
	if base.Ui32(v534) <= base.Ui32(v520) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L149
L152:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	if base.Ui32(v520) < base.Ui32(v534+v536) {
		goto L146
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v522)+8))
	if v539 != 0 {
		v522 = v539
		goto L150
	} else {
		goto L156
	}
L155:
	;
	goto L154
L156:
	;
	goto L151
L157:
	;
	v557 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[13]))
	v559 = v557 - int32(1)
	if v559&v553 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v567 = v501 - v553 + (v553+v559)&(int32(0)-v557)
	goto L160
L159:
	;
	v567 = v501
	goto L160
L160:
	;
	if base.Ui32(v567) <= base.Ui32(v406) {
		goto L143
	} else {
		goto L161
	}
L161:
	;
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[16]))
	if v570 != 0 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[17]))
	v573 = v572 + v567
	if base.B2i32(base.Ui32(v573) <= base.Ui32(v572))|base.B2i32(base.Ui32(v570) < base.Ui32(v573)) != 0 {
		goto L143
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v579 = F_sbrk(m, v567)
	mBase = m.M
	if v579 != v553 {
		v588 = v579
		v594 = v567
		goto L145
	} else {
		goto L166
	}
L165:
	;
	goto L164
L166:
	;
	v664 = v553
	v668 = v567
	goto L139
L167:
	;
	v588 = v583
	v594 = v582
	goto L145
L168:
	;
	if base.Ui32(v406+int32(48)) <= base.Ui32(v594) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v664 = v588
	v668 = v594
	goto L139
L170:
	;
	goto L171
L171:
	;
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[11]))
	v611 = (v606 + (v471 - v594)) & (int32(0) - v606)
	v612 = F_sbrk(m, v611)
	mBase = m.M
	if v612 == int32(-1) {
		goto L143
	} else {
		goto L172
	}
L172:
	;
	v664 = v588
	v668 = v611 + v594
	goto L139
L173:
	;
	goto L143
L174:
	;
	v658 = v652 - v648
	if base.Ui32(v658) <= base.Ui32(v406+int32(40)) {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	v664 = v648
	v668 = v658
	goto L139
L176:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[18])) = v677
	goto L178
L177:
	;
	goto L178
L178:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0]))
	if v685 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if base.B2i32(base.Ui32(v664) <= base.Ui32(v685))|base.B2i32(base.Ui32(v685) < base.Ui32(v699)) != 0 {
		goto L7
	} else {
		goto L196
	}
L180:
	;
	v687 = int32(_a_F_emscripten_builtin_malloc_5)
	goto L183
L181:
	;
	goto L182
L182:
	;
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[3]))
	if base.Ui32(v705) <= base.Ui32(v664) {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v687)+4))
	if v664 == v699+v700 {
		goto L179
	} else {
		goto L185
	}
L184:
	;
	goto L7
L185:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v687)+8))
	if v703 != 0 {
		v687 = v703
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v708 = v705
	goto L189
L188:
	;
	v708 = int32(0)
	goto L189
L189:
	;
	if v708 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[3])) = v664
	goto L192
L191:
	;
	goto L192
L192:
	;
	v713 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[19])) = v668
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[20])) = v664
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[21])) = int32(-1)
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[22])) = v723
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[23])) = v713
	v728 = v713
	goto L193
L193:
	;
	v741 = v728 << (uint(int32(3)) % 32)
	v743 = v741 + int32(_a_F_emscripten_builtin_malloc_0)
	*(*int32)(unsafe.Add(mBase, uint32(v741)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v743
	*(*int32)(unsafe.Add(mBase, uint32(v741)+uint32(_c_F_emscripten_builtin_malloc[24]))) = v743
	v747 = v728 + int32(1)
	if v747 != int32(32) {
		v728 = v747
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v751 = int32(40)
	v752 = v668 - v751
	v756 = (int32(-8) - v664) & int32(7)
	v757 = v752 - v756
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v757
	v760 = v756 + v664
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v760
	*(*int32)(unsafe.Add(mBase, uint32(v760)+4)) = v757 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v752+v664)+4)) = v751
	v770 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[26])) = v770
	goto L6
L195:
	;
	goto L194
L196:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v687)+12))
	if v775&int32(8) != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v687)+4)) = v700 + v668
	v784 = (int32(-8) - v685) & int32(7)
	v785 = v685 + v784
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v785
	v787 = int32(_a_F_emscripten_builtin_malloc_6)
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2]))
	v790 = v789 + v668
	v791 = v790 - v784
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v791
	*(*int32)(unsafe.Add(mBase, uint32(v785)+4)) = v791 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v790+v685)+4)) = int32(40)
	v801 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[26])) = v801
	goto L6
L198:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[3])) = v664
	goto L200
L199:
	;
	goto L200
L200:
	;
	v813 = int32(_a_F_emscripten_builtin_malloc_5)
	goto L202
L201:
	;
	v835 = int32(_a_F_emscripten_builtin_malloc_5)
	goto L209
L202:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v813)))
	if v664+v668 != v825 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v813)+12)))
	if v828&int32(8) == int32(0) {
		goto L4
	} else {
		goto L208
	}
L204:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v813)+8))
	if v827 != 0 {
		v813 = v827
		goto L202
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	goto L203
L207:
	;
	goto L201
L208:
	;
	goto L201
L209:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v835)))
	if base.Ui32(v847) <= base.Ui32(v685) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	v855 = int32(40)
	v856 = v668 - v855
	v859 = int32(7)
	v860 = (int32(-8) - v664) & v859
	v861 = v856 - v860
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v861
	v864 = v860 + v664
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v864)+4)) = v861 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v856+v664)+4)) = v855
	v874 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[25]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[26])) = v874
	v882 = v850 + (int32(39)-v850)&v859 - int32(47)
	if base.Ui32(v882) < base.Ui32(v685+int32(16)) {
		goto L216
	} else {
		goto L217
	}
L211:
	;
	goto L210
L212:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	v850 = v847 + v849
	if base.Ui32(v685) < base.Ui32(v850) {
		goto L211
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v835)+8))
	v835 = v853
	goto L209
L215:
	;
	goto L214
L216:
	;
	v886 = v685
	goto L218
L217:
	;
	v886 = v882
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = int32(27)
	v889 = int32(_a_F_emscripten_builtin_malloc_7)
	v890 = *(*int64)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[27]))
	*(*int64)(unsafe.Add(mBase, uint32(v886)+16)) = v890
	v892 = int32(_a_F_emscripten_builtin_malloc_5)
	v893 = *(*int64)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[20]))
	*(*int64)(unsafe.Add(mBase, uint32(v886)+8)) = v893
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[27])) = v886 + int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[19])) = v668
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[20])) = v664
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[23])) = int32(0)
	v908 = v886 + int32(24)
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v908)+4)) = int32(7)
	if base.Ui32(v908+int32(8)) < base.Ui32(v850) {
		v908 = v908 + int32(4)
		goto L219
	} else {
		goto L221
	}
L220:
	;
	if v886 == v685 {
		goto L6
	} else {
		goto L222
	}
L221:
	;
	goto L220
L222:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v886)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v886)+4)) = v928 & int32(-2)
	v932 = v886 - v685
	*(*int32)(unsafe.Add(mBase, uint32(v685)+4)) = v932 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v886))) = v932
	if base.Ui32(v932) <= base.Ui32(int32(255)) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1065+v685))) = v1054
	*(*int32)(unsafe.Add(mBase, uint32(v1055+v685))) = v1053
	goto L6
L224:
	;
	v940 = v932 & int32(248)
	v942 = v940 + int32(_a_F_emscripten_builtin_malloc_0)
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5]))
	v948 = int32(1) << (uint(int32(base.Ui32(v932)>>(uint(int32(3))%32))) % 32)
	if v944&v948 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	goto L226
L226:
	;
	if base.Ui32(v932) <= base.Ui32(int32(16777215)) {
		goto L231
	} else {
		goto L232
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v956)+12)) = v685
	v1053 = v942
	v1054 = v956
	v1055 = int32(12)
	v1065 = int32(8)
	goto L223
L228:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v944 | v948
	v956 = v942
	goto L227
L229:
	;
	goto L230
L230:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v940)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v956 = v955
	goto L227
L231:
	;
	v967 = base.I32_clz(int32(base.Ui32(v932) >> (uint(int32(8)) % 32)))
	v970 = int32(1)
	v977 = int32(base.Ui32(v932)>>(uint(int32(38)-v967)%32))&v970 | v967<<(uint(v970)%32) ^ int32(62)
	goto L233
L232:
	;
	v977 = int32(31)
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+28)) = v977
	*(*int64)(unsafe.Add(mBase, uint32(v685)+16)) = int64(0)
	v982 = v977 << (uint(int32(2)) % 32)
	v986 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7]))
	v988 = int32(1) << (uint(v977) % 32)
	if v986&v988 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1046)+12)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v1011)+8)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v685)+8)) = v1046
	v1053 = int32(0)
	v1054 = v1011
	v1055 = int32(24)
	v1065 = int32(12)
	goto L223
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+24)) = v1032
	v1053 = v685
	v1054 = v685
	v1055 = int32(8)
	v1065 = int32(12)
	goto L223
L236:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v986 | v988
	*(*int32)(unsafe.Add(mBase, uint32(v982)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v685
	v1032 = v982 + int32(_a_F_emscripten_builtin_malloc_8)
	goto L235
L237:
	;
	goto L238
L238:
	;
	if v977 != int32(31) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1003 = int32(25) - int32(base.Ui32(v977)>>(uint(int32(1))%32))
	goto L241
L240:
	;
	v1003 = int32(0)
	goto L241
L241:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v982)+uint32(_c_F_emscripten_builtin_malloc[9])))
	v1006 = v932 << (uint(v1003) % 32)
	v1011 = v1005
	goto L242
L242:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+4))
	if v1018&int32(-8) == v932 {
		goto L234
	} else {
		goto L244
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+16)) = v685
	v1032 = v1011
	goto L235
L244:
	;
	v1028 = v1011 + int32(base.Ui32(v1006)>>(uint(int32(29))%32))&int32(4)
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	if v1029 != 0 {
		v1006 = v1006 << (uint(int32(1)) % 32)
		v1011 = v1029
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v1086 = v1083 - v406
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v1086
	v1088 = int32(_a_F_emscripten_builtin_malloc_1)
	v1090 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0]))
	v1091 = v1090 + v406
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v1091
	*(*int32)(unsafe.Add(mBase, uint32(v1091)+4)) = v1086 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1090)+4)) = v406 | int32(3)
	v1727 = v1090 + int32(8)
	goto L1
L247:
	;
	v1727 = v1125 + int32(8)
	goto L1
L248:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[0])) = v1134
	v1141 = int32(_a_F_emscripten_builtin_malloc_6)
	v1143 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2]))
	v1144 = v1143 + v1135
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[2])) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+4)) = v1144 | int32(1)
	goto L247
L249:
	;
	goto L250
L250:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8]))
	if v1150 == v1133 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8])) = v1134
	v1154 = int32(_a_F_emscripten_builtin_malloc_9)
	v1156 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4]))
	v1157 = v1156 + v1135
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4])) = v1157
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+4)) = v1157 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1157+v1134))) = v1157
	goto L247
L252:
	;
	goto L253
L253:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	if v1164&int32(3) == int32(1) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1170 = v1164 & int32(-8)
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+12))
	if base.Ui32(v1164) <= base.Ui32(int32(255)) {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v1284 = v1135
	v1287 = v1164
	v1288 = v1133
	goto L256
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1288)+4)) = v1287 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+4)) = v1284 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1284+v1134))) = v1284
	if base.Ui32(v1284) <= base.Ui32(int32(255)) {
		goto L292
	} else {
		goto L293
	}
L257:
	;
	v1279 = v1133 + v1170
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+4))
	v1284 = v1135 + v1170
	v1287 = v1280
	v1288 = v1279
	goto L256
L258:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+8))
	if v1174 == v1171 {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	goto L260
L260:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+24))
	if v1171 != v1133 {
		goto L265
	} else {
		goto L266
	}
L261:
	;
	v1176 = int32(_a_F_emscripten_builtin_malloc_10)
	v1178 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v1178 & base.I32_rotl(int32(-2), int32(base.Ui32(v1164)>>(uint(int32(3))%32)))
	goto L257
L262:
	;
	goto L263
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+12)) = v1171
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+8)) = v1174
	goto L257
L264:
	;
	if v1187 == int32(0) {
		goto L257
	} else {
		goto L277
	}
L265:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1189)+12)) = v1171
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+8)) = v1189
	v1226 = v1171
	goto L264
L266:
	;
	goto L267
L267:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+20))
	if v1192 != 0 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v1226 = int32(0)
	goto L264
L269:
	;
	v1200 = v1192
	v1201 = v1133 + int32(20)
	goto L271
L270:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+16))
	if v1195 == int32(0) {
		goto L268
	} else {
		goto L272
	}
L271:
	;
	__phi1202 = v1201
	__phi1207 = v1200
	v1202 = __phi1202
	v1207 = __phi1207
	goto L273
L272:
	;
	v1200 = v1195
	v1201 = v1133 + int32(16)
	goto L271
L273:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+20))
	if v1216 != 0 {
		__phi1202 = v1207 + int32(20)
		__phi1207 = v1216
		v1202 = __phi1202
		v1207 = __phi1207
		goto L273
	} else {
		goto L275
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1202))) = int32(0)
	v1226 = v1207
	goto L264
L275:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1207)+16))
	if v1219 != 0 {
		__phi1202 = v1207 + int32(16)
		__phi1207 = v1219
		v1202 = __phi1202
		v1207 = __phi1207
		goto L273
	} else {
		goto L276
	}
L276:
	;
	goto L274
L277:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+28))
	v1239 = v1237 << (uint(int32(2)) % 32)
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1239)+uint32(_c_F_emscripten_builtin_malloc[9])))
	if v1240 == v1133 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+24)) = v1187
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+16))
	if v1259 != 0 {
		goto L288
	} else {
		goto L289
	}
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1239)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v1226
	if v1226 != 0 {
		goto L278
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+16))
	if v1133 == v1252 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	v1245 = int32(_a_F_emscripten_builtin_malloc_11)
	v1247 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v1247 & base.I32_rotl(int32(-2), v1237)
	goto L257
L283:
	;
	if v1226 == int32(0) {
		goto L257
	} else {
		goto L287
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+16)) = v1226
	goto L283
L285:
	;
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1187)+20)) = v1226
	goto L283
L287:
	;
	goto L278
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+16)) = v1259
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+24)) = v1226
	goto L290
L289:
	;
	goto L290
L290:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+20))
	if v1262 == int32(0) {
		goto L257
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1226)+20)) = v1262
	*(*int32)(unsafe.Add(mBase, uint32(v1262)+24)) = v1226
	goto L257
L292:
	;
	v1305 = v1284 & int32(248)
	v1307 = v1305 + int32(_a_F_emscripten_builtin_malloc_0)
	v1309 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5]))
	v1313 = int32(1) << (uint(int32(base.Ui32(v1284)>>(uint(int32(3))%32))) % 32)
	if v1309&v1313 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L293:
	;
	goto L294
L294:
	;
	if base.Ui32(v1284) <= base.Ui32(int32(16777215)) {
		goto L299
	} else {
		goto L300
	}
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1305)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1321)+12)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+12)) = v1307
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+8)) = v1321
	goto L247
L296:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v1309 | v1313
	v1321 = v1307
	goto L295
L297:
	;
	goto L298
L298:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1305)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v1321 = v1320
	goto L295
L299:
	;
	v1332 = base.I32_clz(int32(base.Ui32(v1284) >> (uint(int32(8)) % 32)))
	v1335 = int32(1)
	v1343 = int32(base.Ui32(v1284)>>(uint(int32(38)-v1332)%32))&v1335 | v1332<<(uint(v1335)%32) ^ int32(62)
	goto L301
L300:
	;
	v1343 = int32(31)
	goto L301
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+28)) = v1343
	*(*int64)(unsafe.Add(mBase, uint32(v1134)+16)) = int64(0)
	v1348 = v1343 << (uint(int32(2)) % 32)
	v1352 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7]))
	v1354 = int32(1) << (uint(v1343) % 32)
	if v1352&v1354 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1413)+12)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1373)+8)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+12)) = v1373
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+8)) = v1413
	goto L247
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+12)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+8)) = v1134
	goto L247
L304:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v1352 | v1354
	*(*int32)(unsafe.Add(mBase, uint32(v1348)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+24)) = v1348 + int32(_a_F_emscripten_builtin_malloc_8)
	goto L303
L305:
	;
	goto L306
L306:
	;
	if v1343 != int32(31) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1370 = int32(25) - int32(base.Ui32(v1343)>>(uint(int32(1))%32))
	goto L309
L308:
	;
	v1370 = int32(0)
	goto L309
L309:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+uint32(_c_F_emscripten_builtin_malloc[9])))
	v1373 = v1372
	v1376 = v1284 << (uint(v1370) % 32)
	goto L310
L310:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	if v1385&int32(-8) == v1284 {
		goto L302
	} else {
		goto L312
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1395)+16)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(v1134)+24)) = v1373
	goto L303
L312:
	;
	v1395 = v1373 + int32(base.Ui32(v1376)>>(uint(int32(29))%32))&int32(4)
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+16))
	if v1396 != 0 {
		v1373 = v1396
		v1376 = v1376 << (uint(int32(1)) % 32)
		goto L310
	} else {
		goto L313
	}
L313:
	;
	goto L311
L314:
	;
	if base.Ui32(v353) <= base.Ui32(int32(15)) {
		goto L331
	} else {
		goto L332
	}
L315:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v354)+28))
	v1450 = v1448 << (uint(int32(2)) % 32)
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+uint32(_c_F_emscripten_builtin_malloc[9])))
	if v1451 == v354 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1434)+24)) = v367
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v354)+16))
	if v1468 != 0 {
		goto L326
	} else {
		goto L327
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v1434
	if v1434 != 0 {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v367)+16))
	if v354 == v1461 {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	v1459 = v212 & base.I32_rotl(int32(-2), v1448)
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v1459
	v1478 = v1459
	goto L314
L321:
	;
	if v1434 == int32(0) {
		v1478 = v212
		goto L314
	} else {
		goto L325
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+16)) = v1434
	goto L321
L323:
	;
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+20)) = v1434
	goto L321
L325:
	;
	goto L316
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1434)+16)) = v1468
	*(*int32)(unsafe.Add(mBase, uint32(v1468)+24)) = v1434
	goto L328
L327:
	;
	goto L328
L328:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v354)+20))
	if v1471 == int32(0) {
		v1478 = v212
		goto L314
	} else {
		goto L329
	}
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1434)+20)) = v1471
	*(*int32)(unsafe.Add(mBase, uint32(v1471)+24)) = v1434
	v1478 = v212
	goto L314
L330:
	;
	v1727 = v354 + int32(8)
	goto L1
L331:
	;
	v1481 = v210 + v353
	*(*int32)(unsafe.Add(mBase, uint32(v354)+4)) = v1481 | int32(3)
	v1485 = v1481 + v354
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1485)+4)) = v1486 | int32(1)
	goto L330
L332:
	;
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354)+4)) = v210 | int32(3)
	v1493 = v210 + v354
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+4)) = v353 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1493+v353))) = v353
	if base.Ui32(v353) <= base.Ui32(int32(255)) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1502 = v353 & int32(248)
	v1504 = v1502 + int32(_a_F_emscripten_builtin_malloc_0)
	v1506 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5]))
	v1510 = int32(1) << (uint(int32(base.Ui32(v353)>>(uint(int32(3))%32))) % 32)
	if v1506&v1510 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L335:
	;
	goto L336
L336:
	;
	if base.Ui32(v353) <= base.Ui32(int32(16777215)) {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1502)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1518)+12)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+12)) = v1504
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+8)) = v1518
	goto L330
L338:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v1506 | v1510
	v1518 = v1504
	goto L337
L339:
	;
	goto L340
L340:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v1518 = v1517
	goto L337
L341:
	;
	v1529 = base.I32_clz(int32(base.Ui32(v353) >> (uint(int32(8)) % 32)))
	v1532 = int32(1)
	v1539 = int32(base.Ui32(v353)>>(uint(int32(38)-v1529)%32))&v1532 | v1529<<(uint(v1532)%32) ^ int32(62)
	goto L343
L342:
	;
	v1539 = int32(31)
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+28)) = v1539
	*(*int64)(unsafe.Add(mBase, uint32(v1493)+16)) = int64(0)
	v1544 = v1539 << (uint(int32(2)) % 32)
	v1548 = int32(1) << (uint(v1539) % 32)
	if v1478&v1548 == int32(0) {
		goto L346
	} else {
		goto L347
	}
L344:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1607)+12)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1568)+8)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+12)) = v1568
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+8)) = v1607
	goto L330
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+12)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+8)) = v1493
	goto L330
L346:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v1548 | v1478
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+24)) = v1544 + int32(_a_F_emscripten_builtin_malloc_8)
	goto L345
L347:
	;
	goto L348
L348:
	;
	if v1539 != int32(31) {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1564 = int32(25) - int32(base.Ui32(v1539)>>(uint(int32(1))%32))
	goto L351
L350:
	;
	v1564 = int32(0)
	goto L351
L351:
	;
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1544)+uint32(_c_F_emscripten_builtin_malloc[9])))
	v1567 = v353 << (uint(v1564) % 32)
	v1568 = v1566
	goto L352
L352:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+4))
	if v1579&int32(-8) == v353 {
		goto L344
	} else {
		goto L354
	}
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1589)+16)) = v1493
	*(*int32)(unsafe.Add(mBase, uint32(v1493)+24)) = v1568
	goto L345
L354:
	;
	v1589 = v1568 + int32(base.Ui32(v1567)>>(uint(int32(29))%32))&int32(4)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1589)+16))
	if v1590 != 0 {
		v1567 = v1567 << (uint(int32(1)) % 32)
		v1568 = v1590
		goto L352
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	if base.Ui32(v146) <= base.Ui32(int32(15)) {
		goto L373
	} else {
		goto L374
	}
L357:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v144)+28))
	v1644 = v1642 << (uint(int32(2)) % 32)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1644)+uint32(_c_F_emscripten_builtin_malloc[9])))
	if v1645 == v144 {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+24)) = v168
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	if v1662 != 0 {
		goto L368
	} else {
		goto L369
	}
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1644)+uint32(_c_F_emscripten_builtin_malloc[9]))) = v1628
	if v1628 != 0 {
		goto L358
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	if v144 == v1655 {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[7])) = v131 & base.I32_rotl(int32(-2), v1642)
	goto L356
L363:
	;
	if v1628 == int32(0) {
		goto L356
	} else {
		goto L367
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = v1628
	goto L363
L365:
	;
	goto L366
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+20)) = v1628
	goto L363
L367:
	;
	goto L358
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+16)) = v1662
	*(*int32)(unsafe.Add(mBase, uint32(v1662)+24)) = v1628
	goto L370
L369:
	;
	goto L370
L370:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	if v1665 == int32(0) {
		goto L356
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+20)) = v1665
	*(*int32)(unsafe.Add(mBase, uint32(v1665)+24)) = v1628
	goto L356
L372:
	;
	v1727 = v144 + int32(8)
	goto L1
L373:
	;
	v1674 = v28 + v146
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v1674 | int32(3)
	v1678 = v1674 + v144
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1678)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1678)+4)) = v1679 | int32(1)
	goto L372
L374:
	;
	goto L375
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v28 | int32(3)
	v1686 = v144 + v28
	*(*int32)(unsafe.Add(mBase, uint32(v1686)+4)) = v146 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v146+v1686))) = v146
	if v64 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1693 = v64 & int32(-8)
	v1695 = v1693 + int32(_a_F_emscripten_builtin_malloc_0)
	v1697 = *(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8]))
	v1701 = int32(1) << (uint(int32(base.Ui32(v64)>>(uint(int32(3))%32))) % 32)
	if v1701&v20 == int32(0) {
		goto L380
	} else {
		goto L381
	}
L377:
	;
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[8])) = v1686
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[4])) = v146
	goto L372
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+uint32(_c_F_emscripten_builtin_malloc[6]))) = v1697
	*(*int32)(unsafe.Add(mBase, uint32(v1709)+12)) = v1697
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+12)) = v1695
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+8)) = v1709
	goto L378
L380:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_emscripten_builtin_malloc[5])) = v1701 | v20
	v1709 = v1695
	goto L379
L381:
	;
	goto L382
L382:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+uint32(_c_F_emscripten_builtin_malloc[6])))
	v1709 = v1708
	goto L379
}
