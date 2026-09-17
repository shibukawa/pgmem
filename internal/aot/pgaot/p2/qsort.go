package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v228 int32
	_ = v228
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
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
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v519 int32
	_ = v519
	var v532 int32
	_ = v532
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v661 int32
	_ = v661
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v745 int32
	_ = v745
	var v759 int32
	_ = v759
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v865 int32
	_ = v865
	var v888 int32
	_ = v888
	var v896 int32
	_ = v896
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1024 int32
	_ = v1024
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1181 int32
	_ = v1181
	var v1204 int32
	_ = v1204
	var v1211 int32
	_ = v1211
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1321 int32
	_ = v1321
	var v1331 int32
	_ = v1331
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1368 int32
	_ = v1368
	var v1374 int32
	_ = v1374
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1426 int32
	_ = v1426
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1525 int32
	_ = v1525
	v25 = int32(0) - l2
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v1266 = l0
		v1267 = l1
		v1268 = l2
		v1269 = l3
		v1270 = l4
		v1287 = v25
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v1289 = v1267 * v1268
	if base.Ui32(v1289) <= base.Ui32(v1268) {
		goto L1
	} else {
		goto L144
	}
L3:
	;
	v34 = l0
	v35 = l1
	v36 = l2
	v37 = l3
	v38 = l4
	v50 = l2 & int32(-4)
	v51 = l2 & int32(3)
	v52 = l2 - int32(1)
	v55 = v25
	goto L4
L4:
	;
	v57 = v34 + v36
	v59 = v35
	goto L6
L5:
	;
	v1266 = v34
	v1267 = v1263
	v1268 = v36
	v1269 = v37
	v1270 = v38
	v1287 = v55
	goto L2
L6:
	;
	v81 = v59 * v36
	if base.Ui32(v81) <= base.Ui32(v36) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v83 = v34 + v81
	v91 = v57
	goto L9
L9:
	;
	v108 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v91+v55, v91, v38)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v117 = v34 + int32(base.Ui32(v59)>>(uint(int32(1))%32))*v36
	if v59 != int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	return
L12:
	;
	if v108 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v112 = v36 + v91
	if base.Ui32(v112) < base.Ui32(v83) {
		v91 = v112
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L10
L16:
	;
	goto L1
L17:
	;
	v123 = v34 + (v59-int32(1))*v36
	if base.Ui32(v59) < base.Ui32(int32(41)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v151 = v117
	goto L19
L19:
	;
	if v36 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v148 = F_qsort_interruptible_med3(m, v146, v144, v145, v37, v38)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L11
	} else {
		goto L27
	}
L21:
	;
	v144 = v117
	v145 = v123
	v146 = v34
	goto L20
L22:
	;
	goto L23
L23:
	;
	v128 = int32(base.Ui32(v59)>>(uint(int32(3))%32)) * v36
	v131 = v128 << (uint(int32(1)) % 32)
	v133 = F_qsort_interruptible_med3(m, v34, v34+v128, v34+v131, v37, v38)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v137 = F_qsort_interruptible_med3(m, v117-v128, v117, v128+v117, v37, v38)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v141 = F_qsort_interruptible_med3(m, v123-v131, v123-v128, v123, v37, v38)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v144 = v137
	v145 = v141
	v146 = v133
	goto L20
L27:
	;
	v151 = v148
	goto L19
L28:
	;
	v306 = v34 + (v59-int32(1))*v36
	v314 = v306
	v316 = v306
	v317 = v57
	v318 = v57
	goto L40
L29:
	;
	v157 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v52) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v170 = v157
	v174 = v157
	goto L33
L31:
	;
	v228 = v157
	goto L32
L32:
	;
	v251 = v228
	v255 = v157
	goto L37
L33:
	;
	v186 = v34 + v170
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v188 = v151 + v170
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v189)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v187)
	v193 = v170 | int32(1)
	v194 = v34 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v196 = v193 + v151
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v197)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v195)
	v201 = v170 | int32(2)
	v202 = v34 + v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v204 = v201 + v151
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v205)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v203)
	v209 = v170 | int32(3)
	v210 = v34 + v209
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	v212 = v209 + v151
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	*(*uint8)(unsafe.Add(mBase, uint32(v210))) = uint8(v213)
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v211)
	v216 = int32(4)
	v217 = v170 + v216
	v219 = v174 + v216
	if v219 != v50 {
		v170 = v217
		v174 = v219
		goto L33
	} else {
		goto L35
	}
L34:
	;
	if v51 == int32(0) {
		goto L28
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v228 = v217
	goto L32
L37:
	;
	v269 = v34 + v251
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	v271 = v251 + v151
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v272)
	*(*uint8)(unsafe.Add(mBase, uint32(v271))) = uint8(v270)
	v275 = int32(1)
	v278 = v255 + v275
	if v278 != v51 {
		v251 = v251 + v275
		v255 = v278
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L28
L39:
	;
	goto L38
L40:
	;
	if base.Ui32(v314) < base.Ui32(v318) {
		v544 = v317
		v545 = v318
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v36) < base.Ui32(v940) {
		goto L138
	} else {
		goto L139
	}
L42:
	;
	if base.Ui32(v545) <= base.Ui32(v314) {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	v341 = v317
	v342 = v318
	goto L44
L44:
	;
	v354 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v342, v34, v38)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v544 = v519
	v545 = v532
	goto L42
L46:
	;
	if int32(0) < v354 {
		v544 = v341
		v545 = v342
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v354 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v36 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v519 = v341
	goto L50
L50:
	;
	v532 = v36 + v342
	if base.Ui32(v532) <= base.Ui32(v314) {
		v341 = v519
		v342 = v532
		goto L44
	} else {
		goto L63
	}
L51:
	;
	v519 = v36 + v341
	goto L50
L52:
	;
	v362 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v52) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v369 = v362
	v373 = v362
	goto L56
L54:
	;
	v434 = v362
	goto L55
L55:
	;
	v457 = v434
	v464 = v362
	goto L60
L56:
	;
	v391 = v373 + v341
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	v393 = v373 + v342
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v394)
	*(*uint8)(unsafe.Add(mBase, uint32(v393))) = uint8(v392)
	v398 = v373 | int32(1)
	v399 = v341 + v398
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v401 = v398 + v342
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v402)
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v400)
	v406 = v373 | int32(2)
	v407 = v341 + v406
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	v409 = v406 + v342
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v407))) = uint8(v410)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v408)
	v414 = v373 | int32(3)
	v415 = v341 + v414
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v417 = v414 + v342
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417))))
	*(*uint8)(unsafe.Add(mBase, uint32(v415))) = uint8(v418)
	*(*uint8)(unsafe.Add(mBase, uint32(v417))) = uint8(v416)
	v421 = int32(4)
	v422 = v373 + v421
	v424 = v369 + v421
	if v424 != v50 {
		v369 = v424
		v373 = v422
		goto L56
	} else {
		goto L58
	}
L57:
	;
	if v51 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	v434 = v422
	goto L55
L60:
	;
	v474 = v457 + v341
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v476 = v457 + v342
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	*(*uint8)(unsafe.Add(mBase, uint32(v474))) = uint8(v477)
	*(*uint8)(unsafe.Add(mBase, uint32(v476))) = uint8(v475)
	v480 = int32(1)
	v483 = v464 + v480
	if v483 != v51 {
		v457 = v457 + v480
		v464 = v483
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	goto L45
L64:
	;
	goto L41
L65:
	;
	if v36 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L66:
	;
	v565 = v314
	v567 = v316
	goto L69
L67:
	;
	v768 = v314
	v770 = v316
	goto L68
L68:
	;
	v784 = v544 - v34
	v785 = v545 - v544
	if v784 < v785 {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v581 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v565, v34, v38)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	v768 = v759
	v770 = v745
	goto L68
L71:
	;
	if v581 < int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	if v581 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v36 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v745 = v567
	goto L75
L75:
	;
	v759 = v565 + v55
	if base.Ui32(v545) <= base.Ui32(v759) {
		v565 = v759
		v567 = v745
		goto L69
	} else {
		goto L88
	}
L76:
	;
	v745 = v567 + v55
	goto L75
L77:
	;
	v589 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v52) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v596 = v589
	v600 = v589
	goto L81
L79:
	;
	v661 = v589
	goto L80
L80:
	;
	v684 = v661
	v691 = v589
	goto L85
L81:
	;
	v618 = v600 + v565
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618))))
	v620 = v600 + v567
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	*(*uint8)(unsafe.Add(mBase, uint32(v618))) = uint8(v621)
	*(*uint8)(unsafe.Add(mBase, uint32(v620))) = uint8(v619)
	v625 = v600 | int32(1)
	v626 = v565 + v625
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	v628 = v625 + v567
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628))))
	*(*uint8)(unsafe.Add(mBase, uint32(v626))) = uint8(v629)
	*(*uint8)(unsafe.Add(mBase, uint32(v628))) = uint8(v627)
	v633 = v600 | int32(2)
	v634 = v565 + v633
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v634))))
	v636 = v633 + v567
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	*(*uint8)(unsafe.Add(mBase, uint32(v634))) = uint8(v637)
	*(*uint8)(unsafe.Add(mBase, uint32(v636))) = uint8(v635)
	v641 = v600 | int32(3)
	v642 = v565 + v641
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
	v644 = v641 + v567
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v644))))
	*(*uint8)(unsafe.Add(mBase, uint32(v642))) = uint8(v645)
	*(*uint8)(unsafe.Add(mBase, uint32(v644))) = uint8(v643)
	v648 = int32(4)
	v649 = v600 + v648
	v651 = v596 + v648
	if v651 != v50 {
		v596 = v651
		v600 = v649
		goto L81
	} else {
		goto L83
	}
L82:
	;
	if v51 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v661 = v649
	goto L80
L85:
	;
	v701 = v684 + v565
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701))))
	v703 = v684 + v567
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703))))
	*(*uint8)(unsafe.Add(mBase, uint32(v701))) = uint8(v704)
	*(*uint8)(unsafe.Add(mBase, uint32(v703))) = uint8(v702)
	v707 = int32(1)
	v710 = v691 + v707
	if v710 != v51 {
		v684 = v684 + v707
		v691 = v710
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L76
L87:
	;
	goto L86
L88:
	;
	goto L70
L89:
	;
	v940 = v770 - v768
	v942 = v83 - (v36 + v770)
	if base.Ui32(v940) < base.Ui32(v942) {
		goto L105
	} else {
		goto L106
	}
L90:
	;
	v787 = v784
	goto L92
L91:
	;
	v787 = v785
	goto L92
L92:
	;
	if v787 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v790 = v545 - v787
	v792 = v787 & int32(3)
	v793 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v787) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v801 = int32(0)
	v805 = v793
	goto L97
L95:
	;
	v865 = v793
	goto L96
L96:
	;
	v888 = v865
	v896 = v793
	goto L101
L97:
	;
	v823 = v34 + v805
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v823))))
	v825 = v805 + v790
	v826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v825))))
	*(*uint8)(unsafe.Add(mBase, uint32(v823))) = uint8(v826)
	*(*uint8)(unsafe.Add(mBase, uint32(v825))) = uint8(v824)
	v830 = v805 | int32(1)
	v831 = v34 + v830
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831))))
	v833 = v790 + v830
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v833))))
	*(*uint8)(unsafe.Add(mBase, uint32(v831))) = uint8(v834)
	*(*uint8)(unsafe.Add(mBase, uint32(v833))) = uint8(v832)
	v838 = v805 | int32(2)
	v839 = v34 + v838
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	v841 = v790 + v838
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841))))
	*(*uint8)(unsafe.Add(mBase, uint32(v839))) = uint8(v842)
	*(*uint8)(unsafe.Add(mBase, uint32(v841))) = uint8(v840)
	v846 = v805 | int32(3)
	v847 = v34 + v846
	v848 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v847))))
	v849 = v790 + v846
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	*(*uint8)(unsafe.Add(mBase, uint32(v847))) = uint8(v850)
	*(*uint8)(unsafe.Add(mBase, uint32(v849))) = uint8(v848)
	v853 = int32(4)
	v854 = v805 + v853
	v856 = v801 + v853
	if v856 != v787&int32(-4) {
		v801 = v856
		v805 = v854
		goto L97
	} else {
		goto L99
	}
L98:
	;
	if v792 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L99:
	;
	goto L98
L100:
	;
	v865 = v854
	goto L96
L101:
	;
	v906 = v34 + v888
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906))))
	v908 = v888 + v790
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908))))
	*(*uint8)(unsafe.Add(mBase, uint32(v906))) = uint8(v909)
	*(*uint8)(unsafe.Add(mBase, uint32(v908))) = uint8(v907)
	v912 = int32(1)
	v915 = v896 + v912
	if v915 != v792 {
		v888 = v888 + v912
		v896 = v915
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L89
L103:
	;
	goto L102
L104:
	;
	if base.Ui32(v940) < base.Ui32(v785) {
		goto L64
	} else {
		goto L119
	}
L105:
	;
	v944 = v940
	goto L107
L106:
	;
	v944 = v942
	goto L107
L107:
	;
	if v944 == int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v947 = v83 - v944
	v949 = v944 & int32(3)
	v950 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v944) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v964 = v950
	v966 = int32(0)
	goto L112
L110:
	;
	v1024 = v950
	goto L111
L111:
	;
	v1046 = v950
	v1047 = v1024
	goto L116
L112:
	;
	v980 = v964 + v545
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980))))
	v982 = v947 + v964
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	*(*uint8)(unsafe.Add(mBase, uint32(v980))) = uint8(v983)
	*(*uint8)(unsafe.Add(mBase, uint32(v982))) = uint8(v981)
	v987 = v964 | int32(1)
	v988 = v545 + v987
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988))))
	v990 = v947 + v987
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	*(*uint8)(unsafe.Add(mBase, uint32(v988))) = uint8(v991)
	*(*uint8)(unsafe.Add(mBase, uint32(v990))) = uint8(v989)
	v995 = v964 | int32(2)
	v996 = v545 + v995
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	v998 = v947 + v995
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	*(*uint8)(unsafe.Add(mBase, uint32(v996))) = uint8(v999)
	*(*uint8)(unsafe.Add(mBase, uint32(v998))) = uint8(v997)
	v1003 = v964 | int32(3)
	v1004 = v545 + v1003
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004))))
	v1006 = v947 + v1003
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1004))) = uint8(v1007)
	*(*uint8)(unsafe.Add(mBase, uint32(v1006))) = uint8(v1005)
	v1010 = int32(4)
	v1011 = v964 + v1010
	v1013 = v966 + v1010
	if v1013 != v944&int32(-4) {
		v964 = v1011
		v966 = v1013
		goto L112
	} else {
		goto L114
	}
L113:
	;
	if v949 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L114:
	;
	goto L113
L115:
	;
	v1024 = v1011
	goto L111
L116:
	;
	v1063 = v1047 + v545
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1063))))
	v1065 = v947 + v1047
	v1066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1065))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1063))) = uint8(v1066)
	*(*uint8)(unsafe.Add(mBase, uint32(v1065))) = uint8(v1064)
	v1069 = int32(1)
	v1072 = v1046 + v1069
	if v1072 != v949 {
		v1046 = v1072
		v1047 = v1047 + v1069
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L104
L118:
	;
	goto L117
L119:
	;
	if base.Ui32(v36) < base.Ui32(v785) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1099 = base.I32_div_u_s(v785, v36)
	F_qsort_arg(m, v34, v1099, v36, v37, v38)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L11
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if base.Ui32(v940) <= base.Ui32(v36) {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v1103 = v83 - v940
	v1104 = base.I32_div_u_s(v940, v36)
	if base.Ui32(int32(7)) <= base.Ui32(v1104) {
		v34 = v1103
		v35 = v1104
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v1266 = v1103
	v1267 = v1104
	v1268 = v36
	v1269 = v37
	v1270 = v38
	v1287 = v55
	goto L2
L126:
	;
	v314 = v565 + v55
	v316 = v567
	v317 = v544
	v318 = v36 + v545
	goto L40
L127:
	;
	v1109 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v52) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1116 = v1109
	v1120 = v1109
	goto L131
L129:
	;
	v1181 = v1109
	goto L130
L130:
	;
	v1204 = v1181
	v1211 = v1109
	goto L135
L131:
	;
	v1138 = v1120 + v545
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138))))
	v1140 = v1120 + v565
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1138))) = uint8(v1141)
	*(*uint8)(unsafe.Add(mBase, uint32(v1140))) = uint8(v1139)
	v1145 = v1120 | int32(1)
	v1146 = v545 + v1145
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	v1148 = v1145 + v565
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1148))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1146))) = uint8(v1149)
	*(*uint8)(unsafe.Add(mBase, uint32(v1148))) = uint8(v1147)
	v1153 = v1120 | int32(2)
	v1154 = v545 + v1153
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154))))
	v1156 = v1153 + v565
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1156))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1154))) = uint8(v1157)
	*(*uint8)(unsafe.Add(mBase, uint32(v1156))) = uint8(v1155)
	v1161 = v1120 | int32(3)
	v1162 = v545 + v1161
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162))))
	v1164 = v1161 + v565
	v1165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1162))) = uint8(v1165)
	*(*uint8)(unsafe.Add(mBase, uint32(v1164))) = uint8(v1163)
	v1168 = int32(4)
	v1169 = v1120 + v1168
	v1171 = v1116 + v1168
	if v1171 != v50 {
		v1116 = v1171
		v1120 = v1169
		goto L131
	} else {
		goto L133
	}
L132:
	;
	if v51 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v1181 = v1169
	goto L130
L135:
	;
	v1221 = v1204 + v545
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	v1223 = v1204 + v565
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1221))) = uint8(v1224)
	*(*uint8)(unsafe.Add(mBase, uint32(v1223))) = uint8(v1222)
	v1227 = int32(1)
	v1230 = v1211 + v1227
	if v1230 != v51 {
		v1204 = v1204 + v1227
		v1211 = v1230
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L126
L137:
	;
	goto L136
L138:
	;
	v1259 = base.I32_div_u_s(v940, v36)
	F_qsort_arg(m, v83-v940, v1259, v36, v37, v38)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L11
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui32(v785) <= base.Ui32(v36) {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v1263 = base.I32_div_u_s(v785, v36)
	if base.Ui32(int32(7)) <= base.Ui32(v1263) {
		v59 = v1263
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L7
L144:
	;
	v1294 = int32(3)
	v1295 = v1268 & v1294
	v1321 = v1266 + v1268
	goto L145
L145:
	;
	if base.Ui32(v1321) <= base.Ui32(v1266) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L1
L147:
	;
	v1525 = v1268 + v1321
	if base.Ui32(v1525) < base.Ui32(v1266+v1289) {
		v1321 = v1525
		goto L145
	} else {
		goto L166
	}
L148:
	;
	v1331 = v1321
	goto L149
L149:
	;
	v1348 = v1331 + v1287
	v1349 = m.T0[v1269].(func(*base.Module, int32, int32, int32) int32)(m, v1348, v1331, v1270)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	goto L147
L151:
	;
	if v1349 <= int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	if v1268 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v1266) < base.Ui32(v1348) {
		v1331 = v1348
		goto L149
	} else {
		goto L165
	}
L154:
	;
	v1355 = int32(0)
	if base.B2i32(base.Ui32(v1268-int32(1)) < base.Ui32(v1294)) == v1355 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1368 = v1355
	v1374 = v1355
	goto L158
L156:
	;
	v1426 = v1355
	goto L157
L157:
	;
	v1449 = v1426
	v1455 = v1355
	goto L162
L158:
	;
	v1384 = v1331 + v1368
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1384))))
	v1386 = v1348 + v1368
	v1387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1386))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1384))) = uint8(v1387)
	*(*uint8)(unsafe.Add(mBase, uint32(v1386))) = uint8(v1385)
	v1391 = v1368 | int32(1)
	v1392 = v1331 + v1391
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392))))
	v1394 = v1391 + v1348
	v1395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1392))) = uint8(v1395)
	*(*uint8)(unsafe.Add(mBase, uint32(v1394))) = uint8(v1393)
	v1399 = v1368 | int32(2)
	v1400 = v1331 + v1399
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400))))
	v1402 = v1399 + v1348
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1402))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1400))) = uint8(v1403)
	*(*uint8)(unsafe.Add(mBase, uint32(v1402))) = uint8(v1401)
	v1407 = v1368 | int32(3)
	v1408 = v1331 + v1407
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1408))))
	v1410 = v1407 + v1348
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1410))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1408))) = uint8(v1411)
	*(*uint8)(unsafe.Add(mBase, uint32(v1410))) = uint8(v1409)
	v1414 = int32(4)
	v1415 = v1368 + v1414
	v1417 = v1374 + v1414
	if v1417 != v1268&int32(-4) {
		v1368 = v1415
		v1374 = v1417
		goto L158
	} else {
		goto L160
	}
L159:
	;
	if v1295 == int32(0) {
		goto L153
	} else {
		goto L161
	}
L160:
	;
	goto L159
L161:
	;
	v1426 = v1415
	goto L157
L162:
	;
	v1467 = v1331 + v1449
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1467))))
	v1469 = v1449 + v1348
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1469))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1467))) = uint8(v1470)
	*(*uint8)(unsafe.Add(mBase, uint32(v1469))) = uint8(v1468)
	v1473 = int32(1)
	v1476 = v1455 + v1473
	if v1476 != v1295 {
		v1449 = v1449 + v1473
		v1455 = v1476
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L153
L164:
	;
	goto L163
L165:
	;
	goto L150
L166:
	;
	goto L146
}
func F_qsort_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int64
	_ = v251
	var v253 int64
	_ = v253
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int64
	_ = v310
	var v312 int64
	_ = v312
	var v314 int64
	_ = v314
	var v316 int64
	_ = v316
	var v318 int64
	_ = v318
	var v320 int64
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v384 int64
	_ = v384
	var v386 int64
	_ = v386
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v442 int64
	_ = v442
	var v444 int32
	_ = v444
	var v445 int64
	_ = v445
	var v447 int64
	_ = v447
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v454 int32
	_ = v454
	var v475 int32
	_ = v475
	var v487 int32
	_ = v487
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v494 int64
	_ = v494
	var v496 int64
	_ = v496
	var v498 int64
	_ = v498
	var v500 int64
	_ = v500
	var v502 int32
	_ = v502
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = l0
	v21 = l1
	goto L1
L1:
	;
	v36 = v20 + int32(16)
	v38 = v21
	goto L3
L2:
	;
	m.G0 = v18 + int32(16)
	return
L3:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v53 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v58 = v20 + v38<<(uint(int32(4))%32)
	if base.Ui32(v38) <= base.Ui32(int32(6)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	goto L4
L11:
	;
	if base.Ui32(v38) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v135 = v36
	goto L25
L14:
	;
	v76 = v36
	goto L15
L15:
	;
	if base.Ui32(v76) <= base.Ui32(v20) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v129 = v76 + int32(16)
	if base.Ui32(v129) < base.Ui32(v58) {
		v76 = v129
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v83 = v76
	goto L19
L19:
	;
	v95 = v83 - int32(16)
	v96 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v95, v83, l3)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	if v96 <= int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v95))) = v110
	if base.Ui32(v20) < base.Ui32(v95) {
		v83 = v95
		goto L19
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	goto L16
L25:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v147 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v163 = v20 + v38<<(uint(int32(3))%32)&int32(-16)
	if v38 != int32(7) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v152 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v135-int32(16), v135, l3)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v152 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v157 = v135 + int32(16)
	if base.Ui32(v58) <= base.Ui32(v157) {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L26
L35:
	;
	v135 = v157
	goto L25
L36:
	;
	v167 = v58 - int32(16)
	if base.Ui32(v38) < base.Ui32(int32(41)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v196 = v163
	goto L38
L38:
	;
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v200
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v204
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v206
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v196)+8)) = v208
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v196))) = v210
	v213 = v58 - int32(16)
	v218 = v213
	v220 = v36
	v221 = v36
	v223 = v213
	goto L47
L39:
	;
	v193 = F_qsort_interruptible_med3(m, v188, v189, v190, l2, l3)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L8
	} else {
		goto L46
	}
L40:
	;
	v188 = v20
	v189 = v163
	v190 = v167
	goto L39
L41:
	;
	goto L42
L42:
	;
	v171 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	v173 = v171 << (uint(int32(4)) % 32)
	v176 = v171 << (uint(int32(5)) % 32)
	v178 = F_qsort_interruptible_med3(m, v20, v20+v173, v20+v176, l2, l3)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v182 = F_qsort_interruptible_med3(m, v163-v173, v163, v163+v173, l2, l3)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v186 = F_qsort_interruptible_med3(m, v167-v176, v167-v173, v167, l2, l3)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v188 = v178
	v189 = v182
	v190 = v186
	goto L39
L46:
	;
	v196 = v193
	goto L38
L47:
	;
	if base.Ui32(v218) < base.Ui32(v220) {
		v279 = v220
		v280 = v221
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui32(v279) <= base.Ui32(v218) {
		goto L64
	} else {
		goto L65
	}
L50:
	;
	v236 = v220
	v237 = v221
	goto L51
L51:
	;
	v245 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v236, v20, l3)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L53
	}
L52:
	;
	v279 = v271
	v280 = v265
	goto L49
L53:
	;
	if int32(0) < v245 {
		v279 = v236
		v280 = v237
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if v245 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v237)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v251
	v253 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v236)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v237)+8)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	*(*int64)(unsafe.Add(mBase, uint32(v237))) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v236)+8)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v261
	v265 = v237 + int32(16)
	goto L57
L56:
	;
	v265 = v237
	goto L57
L57:
	;
	v267 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v267 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v271 = v236 + int32(16)
	if base.Ui32(v271) <= base.Ui32(v218) {
		v236 = v271
		v237 = v265
		goto L51
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	goto L52
L63:
	;
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v279)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v279)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v492
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v293)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v279)+8)) = v494
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	*(*int64)(unsafe.Add(mBase, uint32(v279))) = v496
	v498 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v293)+8)) = v498
	v500 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v500
	v502 = int32(16)
	v218 = v293 - v502
	v220 = v279 + v502
	v221 = v280
	v223 = v298
	goto L47
L64:
	;
	v293 = v218
	v298 = v223
	goto L67
L65:
	;
	v336 = v218
	v341 = v223
	goto L66
L66:
	;
	v348 = int32(4)
	v349 = (v280 - v20) >> (uint(v348) % 32)
	v352 = (v279 - v280) >> (uint(v348) % 32)
	if v349 < v352 {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	v304 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v293, v20, l3)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L8
	} else {
		goto L69
	}
L68:
	;
	v336 = v330
	v341 = v324
	goto L66
L69:
	;
	if v304 < int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	if v304 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v293)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v310
	v312 = *(*int64)(unsafe.Add(mBase, uint32(v293)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v312
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v298)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v293)+8)) = v314
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	*(*int64)(unsafe.Add(mBase, uint32(v293))) = v316
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+8)) = v318
	v320 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v298))) = v320
	v324 = v298 - int32(16)
	goto L73
L72:
	;
	v324 = v298
	goto L73
L73:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v326 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v330 = v293 - int32(16)
	if base.Ui32(v279) <= base.Ui32(v330) {
		v293 = v330
		v298 = v324
		goto L67
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	goto L68
L79:
	;
	v354 = v349
	goto L81
L80:
	;
	v354 = v352
	goto L81
L81:
	;
	if v354 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v370 = int32(0)
	goto L85
L83:
	;
	goto L84
L84:
	;
	v409 = int32(4)
	v410 = (v341 - v336) >> (uint(v409) % 32)
	v415 = (v58-v341)>>(uint(v409)%32) - int32(1)
	if v410 < v415 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v375 = v370 << (uint(int32(4)) % 32)
	v376 = v20 + v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v376)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v379
	v381 = v375 + (v279 - v354<<(uint(int32(4))%32))
	v382 = *(*int64)(unsafe.Add(mBase, uint32(v381)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v376)+8)) = v382
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v381)))
	*(*int64)(unsafe.Add(mBase, uint32(v376))) = v384
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v381)+8)) = v386
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v381))) = v388
	v391 = v370 + int32(1)
	if v391 != v354 {
		v370 = v391
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L84
L87:
	;
	goto L86
L88:
	;
	v417 = v410
	goto L90
L89:
	;
	v417 = v415
	goto L90
L90:
	;
	if v417 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v429 = int32(0)
	goto L94
L92:
	;
	goto L93
L93:
	;
	if base.Ui32(v352) <= base.Ui32(v410) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v438 = v429 << (uint(int32(4)) % 32)
	v439 = v279 + v438
	v440 = *(*int64)(unsafe.Add(mBase, uint32(v439)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v440
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v439)))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v442
	v444 = v438 + (v58 - v417<<(uint(int32(4))%32))
	v445 = *(*int64)(unsafe.Add(mBase, uint32(v444)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v439)+8)) = v445
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v444)))
	*(*int64)(unsafe.Add(mBase, uint32(v439))) = v447
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v444)+8)) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	*(*int64)(unsafe.Add(mBase, uint32(v444))) = v451
	v454 = v429 + int32(1)
	if v454 != v417 {
		v429 = v454
		goto L94
	} else {
		goto L96
	}
L95:
	;
	goto L93
L96:
	;
	goto L95
L97:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v352) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v410) {
		goto L105
	} else {
		goto L106
	}
L100:
	;
	F_qsort_tuple(m, v20, v352, l2, l3)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	if base.Ui32(v410) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	v20 = v58 - v410<<(uint(int32(4))%32)
	v21 = v410
	goto L1
L105:
	;
	F_qsort_tuple(m, v58-v410<<(uint(int32(4))%32), v410, l2, l3)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L8
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	if base.Ui32(int32(1)) < base.Ui32(v352) {
		v38 = v352
		goto L3
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	goto L10
}
