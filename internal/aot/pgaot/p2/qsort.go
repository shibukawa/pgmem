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
	var v173 int32
	_ = v173
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
	var v319 int32
	_ = v319
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
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
	var v521 int32
	_ = v521
	var v532 int32
	_ = v532
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
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
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1179 int32
	_ = v1179
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1319 int32
	_ = v1319
	var v1329 int32
	_ = v1329
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1353 int32
	_ = v1353
	var v1366 int32
	_ = v1366
	var v1372 int32
	_ = v1372
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1424 int32
	_ = v1424
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1523 int32
	_ = v1523
	v25 = int32(0) - l2
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v1262 = l0
		v1263 = l1
		v1264 = l2
		v1265 = l3
		v1266 = l4
		v1283 = v25
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(v1263) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L136
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
	v1262 = v34
	v1263 = v1259
	v1264 = v36
	v1265 = v37
	v1266 = v38
	v1283 = v55
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
	v319 = v57
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
	v173 = v157
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
	v219 = v173 + v216
	if v219 != v50 {
		v170 = v217
		v173 = v219
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
	if base.Ui32(v314) < base.Ui32(v317) {
		v544 = v317
		v546 = v319
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v1256 = base.I32_div_u_s(v940, v36)
	F_qsort_arg(m, v83-v940, v1256, v36, v37, v38)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L11
	} else {
		goto L134
	}
L42:
	;
	if base.Ui32(v544) <= base.Ui32(v314) {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	v341 = v317
	v343 = v319
	goto L44
L44:
	;
	v354 = m.T0[v37].(func(*base.Module, int32, int32, int32) int32)(m, v341, v34, v38)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v544 = v532
	v546 = v521
	goto L42
L46:
	;
	if int32(0) < v354 {
		v544 = v341
		v546 = v343
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
	v521 = v343
	goto L50
L50:
	;
	v532 = v36 + v341
	if base.Ui32(v532) <= base.Ui32(v314) {
		v341 = v532
		v343 = v521
		goto L44
	} else {
		goto L63
	}
L51:
	;
	v521 = v36 + v343
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
	v391 = v373 + v343
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	v393 = v373 + v341
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393))))
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v394)
	*(*uint8)(unsafe.Add(mBase, uint32(v393))) = uint8(v392)
	v398 = v373 | int32(1)
	v399 = v343 + v398
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v401 = v398 + v341
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	*(*uint8)(unsafe.Add(mBase, uint32(v399))) = uint8(v402)
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v400)
	v406 = v373 | int32(2)
	v407 = v343 + v406
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	v409 = v406 + v341
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	*(*uint8)(unsafe.Add(mBase, uint32(v407))) = uint8(v410)
	*(*uint8)(unsafe.Add(mBase, uint32(v409))) = uint8(v408)
	v414 = v373 | int32(3)
	v415 = v343 + v414
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v417 = v414 + v341
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
	v474 = v457 + v343
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	v476 = v457 + v341
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
		goto L122
	} else {
		goto L123
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
	v784 = v546 - v34
	v785 = v544 - v546
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
	if base.Ui32(v544) <= base.Ui32(v759) {
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
	v790 = v544 - v787
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
	v980 = v964 + v544
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980))))
	v982 = v947 + v964
	v983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	*(*uint8)(unsafe.Add(mBase, uint32(v980))) = uint8(v983)
	*(*uint8)(unsafe.Add(mBase, uint32(v982))) = uint8(v981)
	v987 = v964 | int32(1)
	v988 = v544 + v987
	v989 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988))))
	v990 = v947 + v987
	v991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
	*(*uint8)(unsafe.Add(mBase, uint32(v988))) = uint8(v991)
	*(*uint8)(unsafe.Add(mBase, uint32(v990))) = uint8(v989)
	v995 = v964 | int32(2)
	v996 = v544 + v995
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	v998 = v947 + v995
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	*(*uint8)(unsafe.Add(mBase, uint32(v996))) = uint8(v999)
	*(*uint8)(unsafe.Add(mBase, uint32(v998))) = uint8(v997)
	v1003 = v964 | int32(3)
	v1004 = v544 + v1003
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
	v1063 = v1047 + v544
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
	v1098 = base.I32_div_u_s(v785, v36)
	F_qsort_arg(m, v34, v1098, v36, v37, v38)
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	v1101 = v83 - v940
	v1102 = base.I32_div_u_s(v940, v36)
	if base.Ui32(int32(7)) <= base.Ui32(v1102) {
		v34 = v1101
		v35 = v1102
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v1262 = v1101
	v1263 = v1102
	v1264 = v36
	v1265 = v37
	v1266 = v38
	v1283 = v55
	goto L2
L122:
	;
	v314 = v565 + v55
	v316 = v567
	v317 = v36 + v544
	v319 = v546
	goto L40
L123:
	;
	v1107 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v52) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v1114 = v1107
	v1118 = v1107
	goto L127
L125:
	;
	v1179 = v1107
	goto L126
L126:
	;
	v1202 = v1179
	v1209 = v1107
	goto L131
L127:
	;
	v1136 = v1118 + v544
	v1137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1136))))
	v1138 = v1118 + v565
	v1139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1136))) = uint8(v1139)
	*(*uint8)(unsafe.Add(mBase, uint32(v1138))) = uint8(v1137)
	v1143 = v1118 | int32(1)
	v1144 = v544 + v1143
	v1145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1144))))
	v1146 = v1143 + v565
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1144))) = uint8(v1147)
	*(*uint8)(unsafe.Add(mBase, uint32(v1146))) = uint8(v1145)
	v1151 = v1118 | int32(2)
	v1152 = v544 + v1151
	v1153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1152))))
	v1154 = v1151 + v565
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1152))) = uint8(v1155)
	*(*uint8)(unsafe.Add(mBase, uint32(v1154))) = uint8(v1153)
	v1159 = v1118 | int32(3)
	v1160 = v544 + v1159
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160))))
	v1162 = v1159 + v565
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1160))) = uint8(v1163)
	*(*uint8)(unsafe.Add(mBase, uint32(v1162))) = uint8(v1161)
	v1166 = int32(4)
	v1167 = v1118 + v1166
	v1169 = v1114 + v1166
	if v1169 != v50 {
		v1114 = v1169
		v1118 = v1167
		goto L127
	} else {
		goto L129
	}
L128:
	;
	if v51 == int32(0) {
		goto L122
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	v1179 = v1167
	goto L126
L131:
	;
	v1219 = v1202 + v544
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219))))
	v1221 = v1202 + v565
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1219))) = uint8(v1222)
	*(*uint8)(unsafe.Add(mBase, uint32(v1221))) = uint8(v1220)
	v1225 = int32(1)
	v1228 = v1209 + v1225
	if v1228 != v51 {
		v1202 = v1202 + v1225
		v1209 = v1228
		goto L131
	} else {
		goto L133
	}
L132:
	;
	goto L122
L133:
	;
	goto L132
L134:
	;
	v1259 = base.I32_div_u_s(v785, v36)
	if base.Ui32(int32(7)) <= base.Ui32(v1259) {
		v59 = v1259
		goto L6
	} else {
		goto L135
	}
L135:
	;
	goto L7
L136:
	;
	v1287 = v1263 * v1264
	if base.Ui32(v1287) <= base.Ui32(v1264) {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v1292 = int32(3)
	v1293 = v1264 & v1292
	v1319 = v1262 + v1264
	goto L138
L138:
	;
	if base.Ui32(v1319) <= base.Ui32(v1262) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L1
L140:
	;
	v1523 = v1264 + v1319
	if base.Ui32(v1523) < base.Ui32(v1262+v1287) {
		v1319 = v1523
		goto L138
	} else {
		goto L159
	}
L141:
	;
	v1329 = v1319
	goto L142
L142:
	;
	v1346 = v1329 + v1283
	v1347 = m.T0[v1265].(func(*base.Module, int32, int32, int32) int32)(m, v1346, v1329, v1266)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L11
	} else {
		goto L144
	}
L143:
	;
	goto L140
L144:
	;
	if v1347 <= int32(0) {
		goto L140
	} else {
		goto L145
	}
L145:
	;
	if v1264 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	if base.Ui32(v1262) < base.Ui32(v1346) {
		v1329 = v1346
		goto L142
	} else {
		goto L158
	}
L147:
	;
	v1353 = int32(0)
	if base.B2i32(base.Ui32(v1264-int32(1)) < base.Ui32(v1292)) == v1353 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1366 = v1353
	v1372 = v1353
	goto L151
L149:
	;
	v1424 = v1353
	goto L150
L150:
	;
	v1447 = v1424
	v1452 = v1353
	goto L155
L151:
	;
	v1382 = v1329 + v1366
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382))))
	v1384 = v1346 + v1366
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1384))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1382))) = uint8(v1385)
	*(*uint8)(unsafe.Add(mBase, uint32(v1384))) = uint8(v1383)
	v1389 = v1366 | int32(1)
	v1390 = v1329 + v1389
	v1391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1390))))
	v1392 = v1389 + v1346
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1390))) = uint8(v1393)
	*(*uint8)(unsafe.Add(mBase, uint32(v1392))) = uint8(v1391)
	v1397 = v1366 | int32(2)
	v1398 = v1329 + v1397
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1398))))
	v1400 = v1397 + v1346
	v1401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1398))) = uint8(v1401)
	*(*uint8)(unsafe.Add(mBase, uint32(v1400))) = uint8(v1399)
	v1405 = v1366 | int32(3)
	v1406 = v1329 + v1405
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406))))
	v1408 = v1405 + v1346
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1408))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1406))) = uint8(v1409)
	*(*uint8)(unsafe.Add(mBase, uint32(v1408))) = uint8(v1407)
	v1412 = int32(4)
	v1413 = v1366 + v1412
	v1415 = v1372 + v1412
	if v1415 != v1264&int32(-4) {
		v1366 = v1413
		v1372 = v1415
		goto L151
	} else {
		goto L153
	}
L152:
	;
	if v1293 == int32(0) {
		goto L146
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	v1424 = v1413
	goto L150
L155:
	;
	v1465 = v1329 + v1447
	v1466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465))))
	v1467 = v1447 + v1346
	v1468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1467))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1465))) = uint8(v1468)
	*(*uint8)(unsafe.Add(mBase, uint32(v1467))) = uint8(v1466)
	v1471 = int32(1)
	v1474 = v1452 + v1471
	if v1474 != v1293 {
		v1447 = v1447 + v1471
		v1452 = v1474
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L146
L157:
	;
	goto L156
L158:
	;
	goto L143
L159:
	;
	goto L139
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
	var v77 int32
	_ = v77
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
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v419 int64
	_ = v419
	var v421 int64
	_ = v421
	var v423 int64
	_ = v423
	var v425 int64
	_ = v425
	var v427 int64
	_ = v427
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v467 int32
	_ = v467
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int64
	_ = v479
	var v481 int64
	_ = v481
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v486 int64
	_ = v486
	var v488 int64
	_ = v488
	var v490 int64
	_ = v490
	var v492 int64
	_ = v492
	var v494 int64
	_ = v494
	var v496 int64
	_ = v496
	var v499 int32
	_ = v499
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v527 int64
	_ = v527
	var v529 int64
	_ = v529
	var v531 int64
	_ = v531
	var v533 int64
	_ = v533
	var v535 int64
	_ = v535
	var v537 int64
	_ = v537
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v543 int64
	_ = v543
	var v545 int32
	_ = v545
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = l0
	v21 = l1
	goto L1
L1:
	;
	v36 = v20 + int32(24)
	v38 = v21
	goto L3
L2:
	;
	m.G0 = v18 + int32(32)
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
	v139 = v20 + v38*int32(24)
	v144 = v36
	goto L25
L14:
	;
	v77 = v36
	goto L15
L15:
	;
	if base.Ui32(v77) <= base.Ui32(v20) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v135 = v77 + int32(24)
	if base.Ui32(v135) < base.Ui32(v20+v38*int32(24)) {
		v77 = v135
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v83 = v77
	goto L19
L19:
	;
	v95 = v83 - int32(24)
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
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v95)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	*(*int64)(unsafe.Add(mBase, uint32(v83))) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v95)+16)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v114
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v95))) = v116
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
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v156 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v172 = v20 + int32(base.Ui32(v38)>>(uint(int32(1))%32))*int32(24)
	if v38 != int32(7) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v161 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v144-int32(24), v144, l3)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v161 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v166 = v144 + int32(24)
	if base.Ui32(v139) <= base.Ui32(v166) {
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
	v144 = v166
	goto L25
L36:
	;
	v176 = v139 - int32(24)
	if base.Ui32(v38) < base.Ui32(int32(41)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v210 = v172
	goto L38
L38:
	;
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v217
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v219
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v210)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v210)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v223
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v225
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v210)+16)) = v227
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v210)+8)) = v229
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v210))) = v231
	v234 = v139 - int32(24)
	v239 = v234
	v241 = v36
	v242 = v36
	v244 = v234
	goto L47
L39:
	;
	v207 = F_qsort_interruptible_med3(m, v201, v202, v204, l2, l3)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L46
	}
L40:
	;
	v201 = v20
	v202 = v172
	v204 = v176
	goto L39
L41:
	;
	goto L42
L42:
	;
	v180 = int32(base.Ui32(v38) >> (uint(int32(3)) % 32))
	v182 = v180 * int32(24)
	v187 = F_qsort_interruptible_med3(m, v20, v20+v182, v20+v180*int32(48), l2, l3)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v190 = v180 * int32(-24)
	v193 = F_qsort_interruptible_med3(m, v172+v190, v172, v172+v182, l2, l3)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v199 = F_qsort_interruptible_med3(m, v176+v180*int32(-48), v190+v176, v176, l2, l3)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v201 = v187
	v202 = v193
	v204 = v199
	goto L39
L46:
	;
	v210 = v207
	goto L38
L47:
	;
	if base.Ui32(v239) < base.Ui32(v242) {
		v306 = v241
		v307 = v242
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if base.Ui32(v307) <= base.Ui32(v239) {
		goto L64
	} else {
		goto L65
	}
L50:
	;
	v257 = v241
	v258 = v242
	goto L51
L51:
	;
	v266 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v258, v20, l3)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L8
	} else {
		goto L53
	}
L52:
	;
	v306 = v292
	v307 = v298
	goto L49
L53:
	;
	if int32(0) < v266 {
		v306 = v257
		v307 = v258
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if v266 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v257)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v272
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v257)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v257)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v276
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v258)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+16)) = v278
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v258)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v257)+8)) = v280
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v258)))
	*(*int64)(unsafe.Add(mBase, uint32(v257))) = v282
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+16)) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v286
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v258))) = v288
	v292 = v257 + int32(24)
	goto L57
L56:
	;
	v292 = v257
	goto L57
L57:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v294 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L8
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v298 = v258 + int32(24)
	if base.Ui32(v298) <= base.Ui32(v239) {
		v257 = v292
		v258 = v298
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
	v527 = *(*int64)(unsafe.Add(mBase, uint32(v307)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v527
	v529 = *(*int64)(unsafe.Add(mBase, uint32(v307)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v529
	v531 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v531
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+16)) = v533
	v535 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+8)) = v535
	v537 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v307))) = v537
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v539
	v541 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v541
	v543 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v543
	v545 = int32(24)
	v239 = v320 - v545
	v241 = v306
	v242 = v307 + v545
	v244 = v325
	goto L47
L64:
	;
	v320 = v239
	v325 = v244
	goto L67
L65:
	;
	v369 = v239
	v374 = v244
	goto L66
L66:
	;
	v381 = int32(24)
	v382 = base.I32_div_s(v306-v20, v381)
	v385 = base.I32_div_s(v307-v306, v381)
	if v382 < v385 {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	v331 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v320, v20, l3)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L8
	} else {
		goto L69
	}
L68:
	;
	v369 = v363
	v374 = v357
	goto L66
L69:
	;
	if v331 < int32(0) {
		goto L63
	} else {
		goto L70
	}
L70:
	;
	if v331 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v337
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v339
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v341
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v325)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+16)) = v343
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v325)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v345
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
	*(*int64)(unsafe.Add(mBase, uint32(v320))) = v347
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+16)) = v349
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v325)+8)) = v351
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v325))) = v353
	v357 = v325 - int32(24)
	goto L73
L72:
	;
	v357 = v325
	goto L73
L73:
	;
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v359 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v363 = v320 - int32(24)
	if base.Ui32(v307) <= base.Ui32(v363) {
		v320 = v363
		v325 = v357
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
	v387 = v382
	goto L81
L80:
	;
	v387 = v385
	goto L81
L81:
	;
	if v387 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v405 = int32(0)
	goto L85
L83:
	;
	goto L84
L84:
	;
	v448 = int32(24)
	v449 = base.I32_div_s(v374-v369, v448)
	v452 = base.I32_div_s(v139-v374, v448)
	v454 = v452 - int32(1)
	if v449 < v454 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v408 = v405 * int32(24)
	v409 = v20 + v408
	v410 = *(*int64)(unsafe.Add(mBase, uint32(v409)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v410
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v409)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v412
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v414
	v416 = v408 + (v307 + v387*int32(-24))
	v417 = *(*int64)(unsafe.Add(mBase, uint32(v416)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v409)+16)) = v417
	v419 = *(*int64)(unsafe.Add(mBase, uint32(v416)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v409)+8)) = v419
	v421 = *(*int64)(unsafe.Add(mBase, uint32(v416)))
	*(*int64)(unsafe.Add(mBase, uint32(v409))) = v421
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v416)+16)) = v423
	v425 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v416)+8)) = v425
	v427 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v416))) = v427
	v430 = v405 + int32(1)
	if v430 != v387 {
		v405 = v430
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
	v456 = v449
	goto L90
L89:
	;
	v456 = v454
	goto L90
L90:
	;
	if v456 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v467 = int32(0)
	goto L94
L92:
	;
	goto L93
L93:
	;
	if base.Ui32(v385) <= base.Ui32(v449) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v477 = v467 * int32(24)
	v478 = v307 + v477
	v479 = *(*int64)(unsafe.Add(mBase, uint32(v478)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v479
	v481 = *(*int64)(unsafe.Add(mBase, uint32(v478)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v481
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v478)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v483
	v485 = v477 + (v139 + v456*int32(-24))
	v486 = *(*int64)(unsafe.Add(mBase, uint32(v485)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v478)+16)) = v486
	v488 = *(*int64)(unsafe.Add(mBase, uint32(v485)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v478)+8)) = v488
	v490 = *(*int64)(unsafe.Add(mBase, uint32(v485)))
	*(*int64)(unsafe.Add(mBase, uint32(v478))) = v490
	v492 = *(*int64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v485)+16)) = v492
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v485)+8)) = v494
	v496 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v485))) = v496
	v499 = v467 + int32(1)
	if v499 != v456 {
		v467 = v499
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
	F_qsort_tuple(m, v20, v385, l2, l3)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L8
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_qsort_tuple(m, v139+v449*int32(-24), v449, l2, l3)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L8
	} else {
		goto L101
	}
L100:
	;
	v20 = v139 + v449*int32(-24)
	v21 = v449
	goto L1
L101:
	;
	v38 = v385
	goto L3
}
