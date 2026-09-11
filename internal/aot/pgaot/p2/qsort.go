package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_arg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
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
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
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
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v499 int32
	_ = v499
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
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
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v987 int32
	_ = v987
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1273 int32
	_ = v1273
	var v1289 int32
	_ = v1289
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1376 int32
	_ = v1376
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1473 int32
	_ = v1473
	v24 = int32(0) - l2
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v1225 = l0
		v1226 = l1
		v1227 = l2
		v1228 = l3
		v1229 = l4
		v1244 = v24
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v1247 = v1225 + v1227
	v1249 = v1225 + v1226*v1227
	if base.Ui32(v1249) <= base.Ui32(v1247) {
		goto L1
	} else {
		goto L144
	}
L3:
	;
	v31 = l0
	v32 = l1
	v33 = l2
	v34 = l3
	v35 = l4
	v46 = l2 & int32(3)
	v47 = l2 & int32(-4)
	v50 = v24
	goto L4
L4:
	;
	v53 = v31 + v33
	v55 = v32
	goto L6
L5:
	;
	v1225 = v31
	v1226 = v1222
	v1227 = v33
	v1228 = v34
	v1229 = v35
	v1244 = v50
	goto L2
L6:
	;
	v77 = v31 + v55*v33
	if base.Ui32(v77) <= base.Ui32(v53) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v84 = v53
	goto L9
L9:
	;
	v102 = m.T0[v34].(func(*base.Module, int32, int32, int32) int32)(m, v84+v50, v84, v35)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v111 = v31 + int32(base.Ui32(v55)>>(uint(int32(1))%32))*v33
	if v55 != int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	return
L12:
	;
	if v102 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v106 = v33 + v84
	if base.Ui32(v106) < base.Ui32(v77) {
		v84 = v106
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
	v117 = v31 + (v55-int32(1))*v33
	if base.Ui32(v55) < base.Ui32(int32(41)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v146 = v111
	goto L19
L19:
	;
	if v33 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v142 = F_qsort_interruptible_med3(m, v140, v139, v137, v34, v35)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L11
	} else {
		goto L27
	}
L21:
	;
	v137 = v117
	v139 = v111
	v140 = v31
	goto L20
L22:
	;
	goto L23
L23:
	;
	v122 = int32(base.Ui32(v55)>>(uint(int32(3))%32)) * v33
	v125 = v122 << (uint(int32(1)) % 32)
	v127 = F_qsort_interruptible_med3(m, v31, v31+v122, v31+v125, v34, v35)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v131 = F_qsort_interruptible_med3(m, v111-v122, v111, v111+v122, v34, v35)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v135 = F_qsort_interruptible_med3(m, v117-v125, v117-v122, v117, v34, v35)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v137 = v135
	v139 = v131
	v140 = v127
	goto L20
L27:
	;
	v146 = v142
	goto L19
L28:
	;
	v295 = v31 + (v55-int32(1))*v33
	v304 = v53
	v305 = v295
	v307 = v53
	v309 = v295
	goto L40
L29:
	;
	v151 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v161 = v151
	v167 = v151
	goto L33
L31:
	;
	v218 = v151
	goto L32
L32:
	;
	if v46 == int32(0) {
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v178 = v31 + v161
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	v180 = v161 + v146
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v181)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v179)
	v185 = v161 | int32(1)
	v186 = v31 + v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	v188 = v185 + v146
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v189)
	*(*uint8)(unsafe.Add(mBase, uint32(v188))) = uint8(v187)
	v193 = v161 | int32(2)
	v194 = v31 + v193
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v196 = v193 + v146
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v197)
	*(*uint8)(unsafe.Add(mBase, uint32(v196))) = uint8(v195)
	v201 = v161 | int32(3)
	v202 = v31 + v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v204 = v201 + v146
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v205)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v203)
	v208 = int32(4)
	v209 = v161 + v208
	v211 = v167 + v208
	if v211 != v47 {
		v161 = v209
		v167 = v211
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v218 = v209
	goto L32
L35:
	;
	goto L34
L36:
	;
	v242 = v218
	v250 = v151
	goto L37
L37:
	;
	v259 = v31 + v242
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	v261 = v242 + v146
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	*(*uint8)(unsafe.Add(mBase, uint32(v259))) = uint8(v262)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v260)
	v265 = int32(1)
	v268 = v250 + v265
	if v268 != v46 {
		v242 = v242 + v265
		v250 = v268
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
	if base.Ui32(v305) < base.Ui32(v307) {
		v523 = v304
		v526 = v307
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v33) < base.Ui32(v908) {
		goto L138
	} else {
		goto L139
	}
L42:
	;
	if base.Ui32(v526) <= base.Ui32(v305) {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	v327 = v304
	v330 = v307
	goto L44
L44:
	;
	v341 = m.T0[v34].(func(*base.Module, int32, int32, int32) int32)(m, v330, v31, v35)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v523 = v499
	v526 = v513
	goto L42
L46:
	;
	if int32(0) < v341 {
		v523 = v327
		v526 = v330
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v341 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v33 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v499 = v327
	goto L50
L50:
	;
	v513 = v33 + v330
	if base.Ui32(v513) <= base.Ui32(v305) {
		v327 = v499
		v330 = v513
		goto L44
	} else {
		goto L63
	}
L51:
	;
	v499 = v33 + v327
	goto L50
L52:
	;
	v349 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v359 = v349
	v361 = v349
	goto L56
L54:
	;
	v416 = v349
	goto L55
L55:
	;
	if v46 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v376 = v359 + v327
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	v378 = v359 + v330
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v378))))
	*(*uint8)(unsafe.Add(mBase, uint32(v376))) = uint8(v379)
	*(*uint8)(unsafe.Add(mBase, uint32(v378))) = uint8(v377)
	v383 = v359 | int32(1)
	v384 = v327 + v383
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384))))
	v386 = v383 + v330
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v387)
	*(*uint8)(unsafe.Add(mBase, uint32(v386))) = uint8(v385)
	v391 = v359 | int32(2)
	v392 = v327 + v391
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	v394 = v391 + v330
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	*(*uint8)(unsafe.Add(mBase, uint32(v392))) = uint8(v395)
	*(*uint8)(unsafe.Add(mBase, uint32(v394))) = uint8(v393)
	v399 = v359 | int32(3)
	v400 = v327 + v399
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	v402 = v399 + v330
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402))))
	*(*uint8)(unsafe.Add(mBase, uint32(v400))) = uint8(v403)
	*(*uint8)(unsafe.Add(mBase, uint32(v402))) = uint8(v401)
	v406 = int32(4)
	v407 = v359 + v406
	v409 = v361 + v406
	if v409 != v47 {
		v359 = v407
		v361 = v409
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v416 = v407
	goto L55
L58:
	;
	goto L57
L59:
	;
	v440 = v416
	v441 = v349
	goto L60
L60:
	;
	v457 = v440 + v327
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	v459 = v440 + v330
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459))))
	*(*uint8)(unsafe.Add(mBase, uint32(v457))) = uint8(v460)
	*(*uint8)(unsafe.Add(mBase, uint32(v459))) = uint8(v458)
	v463 = int32(1)
	v466 = v441 + v463
	if v466 != v46 {
		v440 = v440 + v463
		v441 = v466
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
	if v33 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L66:
	;
	v547 = v305
	v551 = v309
	goto L69
L67:
	;
	v743 = v305
	v747 = v309
	goto L68
L68:
	;
	v756 = v523 - v31
	v757 = v526 - v523
	if v756 < v757 {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v560 = m.T0[v34].(func(*base.Module, int32, int32, int32) int32)(m, v547, v31, v35)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	v743 = v732
	v747 = v723
	goto L68
L71:
	;
	if v560 < int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	if v560 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v33 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v723 = v551
	goto L75
L75:
	;
	v732 = v547 + v50
	if base.Ui32(v526) <= base.Ui32(v732) {
		v547 = v732
		v551 = v723
		goto L69
	} else {
		goto L88
	}
L76:
	;
	v723 = v551 + v50
	goto L75
L77:
	;
	v568 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v578 = v568
	v580 = v568
	goto L81
L79:
	;
	v635 = v568
	goto L80
L80:
	;
	if v46 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L81:
	;
	v595 = v578 + v547
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	v597 = v578 + v551
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
	*(*uint8)(unsafe.Add(mBase, uint32(v595))) = uint8(v598)
	*(*uint8)(unsafe.Add(mBase, uint32(v597))) = uint8(v596)
	v602 = v578 | int32(1)
	v603 = v547 + v602
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	v605 = v602 + v551
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605))))
	*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v606)
	*(*uint8)(unsafe.Add(mBase, uint32(v605))) = uint8(v604)
	v610 = v578 | int32(2)
	v611 = v547 + v610
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611))))
	v613 = v610 + v551
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	*(*uint8)(unsafe.Add(mBase, uint32(v611))) = uint8(v614)
	*(*uint8)(unsafe.Add(mBase, uint32(v613))) = uint8(v612)
	v618 = v578 | int32(3)
	v619 = v547 + v618
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v621 = v618 + v551
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621))))
	*(*uint8)(unsafe.Add(mBase, uint32(v619))) = uint8(v622)
	*(*uint8)(unsafe.Add(mBase, uint32(v621))) = uint8(v620)
	v625 = int32(4)
	v626 = v578 + v625
	v628 = v580 + v625
	if v628 != v47 {
		v578 = v626
		v580 = v628
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v635 = v626
	goto L80
L83:
	;
	goto L82
L84:
	;
	v659 = v635
	v660 = v568
	goto L85
L85:
	;
	v676 = v659 + v547
	v677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676))))
	v678 = v659 + v551
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v678))))
	*(*uint8)(unsafe.Add(mBase, uint32(v676))) = uint8(v679)
	*(*uint8)(unsafe.Add(mBase, uint32(v678))) = uint8(v677)
	v682 = int32(1)
	v685 = v660 + v682
	if v685 != v46 {
		v659 = v659 + v682
		v660 = v685
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
	v908 = v747 - v743
	v910 = v77 - (v33 + v747)
	if base.Ui32(v908) < base.Ui32(v910) {
		goto L105
	} else {
		goto L106
	}
L90:
	;
	v759 = v756
	goto L92
L91:
	;
	v759 = v757
	goto L92
L92:
	;
	if v759 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v762 = v526 - v759
	v764 = v759 & int32(3)
	v765 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v759) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v777 = v765
	v779 = int32(0)
	goto L97
L95:
	;
	v834 = v765
	goto L96
L96:
	;
	if v764 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L97:
	;
	v794 = v31 + v777
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	v796 = v777 + v762
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	*(*uint8)(unsafe.Add(mBase, uint32(v794))) = uint8(v797)
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v795)
	v801 = v777 | int32(1)
	v802 = v31 + v801
	v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802))))
	v804 = v801 + v762
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v804))))
	*(*uint8)(unsafe.Add(mBase, uint32(v802))) = uint8(v805)
	*(*uint8)(unsafe.Add(mBase, uint32(v804))) = uint8(v803)
	v809 = v777 | int32(2)
	v810 = v31 + v809
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810))))
	v812 = v809 + v762
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v812))))
	*(*uint8)(unsafe.Add(mBase, uint32(v810))) = uint8(v813)
	*(*uint8)(unsafe.Add(mBase, uint32(v812))) = uint8(v811)
	v817 = v777 | int32(3)
	v818 = v31 + v817
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818))))
	v820 = v817 + v762
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	*(*uint8)(unsafe.Add(mBase, uint32(v818))) = uint8(v821)
	*(*uint8)(unsafe.Add(mBase, uint32(v820))) = uint8(v819)
	v824 = int32(4)
	v825 = v777 + v824
	v827 = v779 + v824
	if v827 != v759&int32(-4) {
		v777 = v825
		v779 = v827
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v834 = v825
	goto L96
L99:
	;
	goto L98
L100:
	;
	v858 = v834
	v859 = v765
	goto L101
L101:
	;
	v875 = v31 + v858
	v876 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v875))))
	v877 = v858 + v762
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877))))
	*(*uint8)(unsafe.Add(mBase, uint32(v875))) = uint8(v878)
	*(*uint8)(unsafe.Add(mBase, uint32(v877))) = uint8(v876)
	v881 = int32(1)
	v884 = v859 + v881
	if v884 != v764 {
		v858 = v858 + v881
		v859 = v884
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
	if base.Ui32(v908) < base.Ui32(v757) {
		goto L64
	} else {
		goto L119
	}
L105:
	;
	v912 = v908
	goto L107
L106:
	;
	v912 = v910
	goto L107
L107:
	;
	if v912 == int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v915 = v77 - v912
	v917 = v912 & int32(3)
	v918 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v912) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v930 = v918
	v933 = int32(0)
	goto L112
L110:
	;
	v987 = v918
	goto L111
L111:
	;
	if v917 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L112:
	;
	v947 = v930 + v526
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947))))
	v949 = v930 + v915
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v949))))
	*(*uint8)(unsafe.Add(mBase, uint32(v947))) = uint8(v950)
	*(*uint8)(unsafe.Add(mBase, uint32(v949))) = uint8(v948)
	v954 = v930 | int32(1)
	v955 = v526 + v954
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v955))))
	v957 = v954 + v915
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v957))))
	*(*uint8)(unsafe.Add(mBase, uint32(v955))) = uint8(v958)
	*(*uint8)(unsafe.Add(mBase, uint32(v957))) = uint8(v956)
	v962 = v930 | int32(2)
	v963 = v526 + v962
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963))))
	v965 = v962 + v915
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	*(*uint8)(unsafe.Add(mBase, uint32(v963))) = uint8(v966)
	*(*uint8)(unsafe.Add(mBase, uint32(v965))) = uint8(v964)
	v970 = v930 | int32(3)
	v971 = v526 + v970
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v973 = v970 + v915
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973))))
	*(*uint8)(unsafe.Add(mBase, uint32(v971))) = uint8(v974)
	*(*uint8)(unsafe.Add(mBase, uint32(v973))) = uint8(v972)
	v977 = int32(4)
	v978 = v930 + v977
	v980 = v933 + v977
	if v980 != v912&int32(-4) {
		v930 = v978
		v933 = v980
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v987 = v978
	goto L111
L114:
	;
	goto L113
L115:
	;
	v1011 = v987
	v1019 = v918
	goto L116
L116:
	;
	v1028 = v1011 + v526
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1028))))
	v1030 = v1011 + v915
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1028))) = uint8(v1031)
	*(*uint8)(unsafe.Add(mBase, uint32(v1030))) = uint8(v1029)
	v1034 = int32(1)
	v1037 = v1019 + v1034
	if v1037 != v917 {
		v1011 = v1011 + v1034
		v1019 = v1037
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
	if base.Ui32(v33) < base.Ui32(v757) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1063 = base.I32_div_u_s(v757, v33)
	F_qsort_arg(m, v31, v1063, v33, v34, v35)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L11
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if base.Ui32(v908) <= base.Ui32(v33) {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v1067 = v77 - v908
	v1068 = base.I32_div_u_s(v908, v33)
	if base.Ui32(int32(7)) <= base.Ui32(v1068) {
		v31 = v1067
		v32 = v1068
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v1225 = v1067
	v1226 = v1068
	v1227 = v33
	v1228 = v34
	v1229 = v35
	v1244 = v50
	goto L2
L126:
	;
	v304 = v523
	v305 = v547 + v50
	v307 = v33 + v526
	v309 = v551
	goto L40
L127:
	;
	v1073 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v33) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1083 = v1073
	v1085 = v1073
	goto L131
L129:
	;
	v1140 = v1073
	goto L130
L130:
	;
	if v46 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L131:
	;
	v1100 = v1083 + v526
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100))))
	v1102 = v1083 + v547
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1100))) = uint8(v1103)
	*(*uint8)(unsafe.Add(mBase, uint32(v1102))) = uint8(v1101)
	v1107 = v1083 | int32(1)
	v1108 = v526 + v1107
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
	v1110 = v1107 + v547
	v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1108))) = uint8(v1111)
	*(*uint8)(unsafe.Add(mBase, uint32(v1110))) = uint8(v1109)
	v1115 = v1083 | int32(2)
	v1116 = v526 + v1115
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116))))
	v1118 = v1115 + v547
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1116))) = uint8(v1119)
	*(*uint8)(unsafe.Add(mBase, uint32(v1118))) = uint8(v1117)
	v1123 = v1083 | int32(3)
	v1124 = v526 + v1123
	v1125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124))))
	v1126 = v1123 + v547
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1124))) = uint8(v1127)
	*(*uint8)(unsafe.Add(mBase, uint32(v1126))) = uint8(v1125)
	v1130 = int32(4)
	v1131 = v1083 + v1130
	v1133 = v1085 + v1130
	if v1133 != v47 {
		v1083 = v1131
		v1085 = v1133
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v1140 = v1131
	goto L130
L133:
	;
	goto L132
L134:
	;
	v1164 = v1140
	v1165 = v1073
	goto L135
L135:
	;
	v1181 = v1164 + v526
	v1182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181))))
	v1183 = v1164 + v547
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1183))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1181))) = uint8(v1184)
	*(*uint8)(unsafe.Add(mBase, uint32(v1183))) = uint8(v1182)
	v1187 = int32(1)
	v1190 = v1165 + v1187
	if v1190 != v46 {
		v1164 = v1164 + v1187
		v1165 = v1190
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
	v1218 = base.I32_div_u_s(v908, v33)
	F_qsort_arg(m, v77-v908, v1218, v33, v34, v35)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L11
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui32(v757) <= base.Ui32(v33) {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v1222 = base.I32_div_u_s(v757, v33)
	if base.Ui32(int32(7)) <= base.Ui32(v1222) {
		v55 = v1222
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L7
L144:
	;
	v1254 = v1227 & int32(3)
	v1273 = v1247
	goto L145
L145:
	;
	if base.Ui32(v1273) <= base.Ui32(v1225) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L1
L147:
	;
	v1473 = v1227 + v1273
	if base.Ui32(v1473) < base.Ui32(v1249) {
		v1273 = v1473
		goto L145
	} else {
		goto L166
	}
L148:
	;
	v1289 = v1273
	goto L149
L149:
	;
	v1302 = v1289 + v1244
	v1303 = m.T0[v1228].(func(*base.Module, int32, int32, int32) int32)(m, v1302, v1289, v1229)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	goto L147
L151:
	;
	if v1303 <= int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	if v1227 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v1225) < base.Ui32(v1302) {
		v1289 = v1302
		goto L149
	} else {
		goto L165
	}
L154:
	;
	v1309 = int32(0)
	if base.B2i32(base.Ui32(v1227) < base.Ui32(int32(4))) == v1309 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1319 = v1309
	v1322 = v1309
	goto L158
L156:
	;
	v1376 = v1309
	goto L157
L157:
	;
	if v1254 == int32(0) {
		goto L153
	} else {
		goto L161
	}
L158:
	;
	v1336 = v1319 + v1289
	v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336))))
	v1338 = v1319 + v1302
	v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1336))) = uint8(v1339)
	*(*uint8)(unsafe.Add(mBase, uint32(v1338))) = uint8(v1337)
	v1343 = v1319 | int32(1)
	v1344 = v1289 + v1343
	v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344))))
	v1346 = v1343 + v1302
	v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1344))) = uint8(v1347)
	*(*uint8)(unsafe.Add(mBase, uint32(v1346))) = uint8(v1345)
	v1351 = v1319 | int32(2)
	v1352 = v1289 + v1351
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352))))
	v1354 = v1351 + v1302
	v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1354))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1352))) = uint8(v1355)
	*(*uint8)(unsafe.Add(mBase, uint32(v1354))) = uint8(v1353)
	v1359 = v1319 | int32(3)
	v1360 = v1289 + v1359
	v1361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1360))))
	v1362 = v1359 + v1302
	v1363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1362))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1360))) = uint8(v1363)
	*(*uint8)(unsafe.Add(mBase, uint32(v1362))) = uint8(v1361)
	v1366 = int32(4)
	v1367 = v1319 + v1366
	v1369 = v1322 + v1366
	if v1369 != v1227&int32(-4) {
		v1319 = v1367
		v1322 = v1369
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v1376 = v1367
	goto L157
L160:
	;
	goto L159
L161:
	;
	v1400 = v1376
	v1406 = v1309
	goto L162
L162:
	;
	v1417 = v1400 + v1289
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417))))
	v1419 = v1400 + v1302
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1417))) = uint8(v1420)
	*(*uint8)(unsafe.Add(mBase, uint32(v1419))) = uint8(v1418)
	v1423 = int32(1)
	v1426 = v1406 + v1423
	if v1426 != v1254 {
		v1400 = v1400 + v1423
		v1406 = v1426
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
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
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v368 int32
	_ = v368
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v444 int64
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v451 int64
	_ = v451
	var v453 int64
	_ = v453
	var v455 int64
	_ = v455
	var v458 int32
	_ = v458
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v496 int32
	_ = v496
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v517 int64
	_ = v517
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int64
	_ = v522
	var v524 int64
	_ = v524
	var v526 int64
	_ = v526
	var v528 int64
	_ = v528
	var v531 int32
	_ = v531
	var v555 int32
	_ = v555
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int64
	_ = v572
	var v574 int64
	_ = v574
	var v577 int32
	_ = v577
	var v578 int64
	_ = v578
	var v580 int64
	_ = v580
	var v582 int64
	_ = v582
	var v584 int64
	_ = v584
	var v586 int32
	_ = v586
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = l0
	v24 = l1
	goto L1
L1:
	;
	v42 = v23 + int32(16)
	v44 = v24
	goto L3
L2:
	;
	m.G0 = v21 + int32(16)
	return
L3:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v62 != 0 {
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v67 = v23 + v44<<(uint(int32(4))%32)
	if base.Ui32(v44) <= base.Ui32(int32(6)) {
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
	if base.Ui32(v67) <= base.Ui32(v42) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v67) <= base.Ui32(v42) {
		goto L10
	} else {
		goto L25
	}
L14:
	;
	v85 = v42
	goto L15
L15:
	;
	if base.Ui32(v85) <= base.Ui32(v23) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v152 = v85 + int32(16)
	if base.Ui32(v152) < base.Ui32(v67) {
		v85 = v152
		goto L15
	} else {
		goto L24
	}
L18:
	;
	v94 = v85
	goto L19
L19:
	;
	v109 = v94 - int32(16)
	v110 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v109, v94, l3)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	if v110 <= int32(0) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v114 = int32(8)
	v115 = v21 + v114
	v117 = v94 + v114
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
	*(*int64)(unsafe.Add(mBase, uint32(v115))) = v118
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v120
	v123 = v109 + v114
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v126
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v115)))
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v128
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = v130
	if base.Ui32(v23) < base.Ui32(v109) {
		v94 = v109
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
	v159 = v42
	goto L26
L26:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v174 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v190 = v23 + v44<<(uint(int32(3))%32)&int32(-16)
	if v44 != int32(7) {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v179 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v159-int32(16), v159, l3)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v179 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v184 = v159 + int32(16)
	if base.Ui32(v67) <= base.Ui32(v184) {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L27
L36:
	;
	v159 = v184
	goto L26
L37:
	;
	v194 = v67 - int32(16)
	if base.Ui32(v44) < base.Ui32(int32(41)) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v222 = v190
	goto L39
L39:
	;
	v227 = int32(8)
	v228 = v21 + v227
	v230 = v23 + v227
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v230)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v233
	v236 = v222 + v227
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v236)))
	*(*int64)(unsafe.Add(mBase, uint32(v230))) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v222)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v239
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v236))) = v241
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v222))) = v243
	v246 = v67 - int32(16)
	v251 = v246
	v252 = v42
	v254 = v42
	v260 = v246
	goto L48
L40:
	;
	v220 = F_qsort_interruptible_med3(m, v219, v215, v216, l2, l3)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L8
	} else {
		goto L47
	}
L41:
	;
	v215 = v190
	v216 = v194
	v219 = v23
	goto L40
L42:
	;
	goto L43
L43:
	;
	v198 = int32(base.Ui32(v44) >> (uint(int32(3)) % 32))
	v200 = v198 << (uint(int32(4)) % 32)
	v203 = v198 << (uint(int32(5)) % 32)
	v205 = F_qsort_interruptible_med3(m, v23, v23+v200, v23+v203, l2, l3)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v209 = F_qsort_interruptible_med3(m, v190-v200, v190, v190+v200, l2, l3)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v213 = F_qsort_interruptible_med3(m, v194-v203, v194-v200, v194, l2, l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v215 = v209
	v216 = v213
	v219 = v205
	goto L40
L47:
	;
	v222 = v220
	goto L39
L48:
	;
	if base.Ui32(v251) < base.Ui32(v252) {
		v323 = v252
		v325 = v254
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if base.Ui32(v323) <= base.Ui32(v251) {
		goto L65
	} else {
		goto L66
	}
L51:
	;
	v271 = v252
	v273 = v254
	goto L52
L52:
	;
	v284 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v271, v23, l3)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L8
	} else {
		goto L54
	}
L53:
	;
	v323 = v316
	v325 = v309
	goto L50
L54:
	;
	if int32(0) < v284 {
		v323 = v271
		v325 = v273
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if v284 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v290 = int32(8)
	v291 = v273 + v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v291)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v292
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v273)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v294
	v297 = v271 + v290
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v297)))
	*(*int64)(unsafe.Add(mBase, uint32(v291))) = v298
	v300 = *(*int64)(unsafe.Add(mBase, uint32(v271)))
	*(*int64)(unsafe.Add(mBase, uint32(v273))) = v300
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v297))) = v302
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v271))) = v304
	v309 = v273 + int32(16)
	goto L58
L57:
	;
	v309 = v273
	goto L58
L58:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v312 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L8
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v316 = v271 + int32(16)
	if base.Ui32(v316) <= base.Ui32(v251) {
		v271 = v316
		v273 = v309
		goto L52
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	goto L53
L64:
	;
	v570 = int32(8)
	v571 = v323 + v570
	v572 = *(*int64)(unsafe.Add(mBase, uint32(v571)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v572
	v574 = *(*int64)(unsafe.Add(mBase, uint32(v323)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v574
	v577 = v341 + v570
	v578 = *(*int64)(unsafe.Add(mBase, uint32(v577)))
	*(*int64)(unsafe.Add(mBase, uint32(v571))) = v578
	v580 = *(*int64)(unsafe.Add(mBase, uint32(v341)))
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = v580
	v582 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v577))) = v582
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v584
	v586 = int32(16)
	v251 = v341 - v586
	v252 = v323 + v586
	v254 = v325
	v260 = v350
	goto L48
L65:
	;
	v341 = v251
	v350 = v260
	goto L68
L66:
	;
	v393 = v251
	v402 = v260
	goto L67
L67:
	;
	v408 = int32(4)
	v409 = (v325 - v23) >> (uint(v408) % 32)
	v412 = (v323 - v325) >> (uint(v408) % 32)
	if v409 < v412 {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	v355 = m.T0[l2].(func(*base.Module, int32, int32, int32) int32)(m, v341, v23, l3)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L70
	}
L69:
	;
	v393 = v387
	v402 = v381
	goto L67
L70:
	;
	if v355 < int32(0) {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	if v355 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v361 = int32(8)
	v362 = v341 + v361
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v341)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v365
	v368 = v350 + v361
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v368)))
	*(*int64)(unsafe.Add(mBase, uint32(v362))) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v350)))
	*(*int64)(unsafe.Add(mBase, uint32(v341))) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v368))) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v350))) = v375
	v381 = v350 - int32(16)
	goto L74
L73:
	;
	v381 = v350
	goto L74
L74:
	;
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_tuple[0]))
	if v383 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L8
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v387 = v341 - int32(16)
	if base.Ui32(v323) <= base.Ui32(v387) {
		v341 = v387
		v350 = v381
		goto L68
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	goto L69
L80:
	;
	v414 = v409
	goto L82
L81:
	;
	v414 = v412
	goto L82
L82:
	;
	if v414 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v425 = int32(0)
	goto L86
L84:
	;
	goto L85
L85:
	;
	v479 = int32(4)
	v480 = (v402 - v393) >> (uint(v479) % 32)
	v485 = (v67-v402)>>(uint(v479)%32) - int32(1)
	if v480 < v485 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v438 = v425 << (uint(int32(4)) % 32)
	v439 = v23 + v438
	v440 = int32(8)
	v441 = v439 + v440
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v441)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v442
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v439)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v444
	v446 = v438 + (v323 - v414<<(uint(int32(4))%32))
	v448 = v446 + v440
	v449 = *(*int64)(unsafe.Add(mBase, uint32(v448)))
	*(*int64)(unsafe.Add(mBase, uint32(v441))) = v449
	v451 = *(*int64)(unsafe.Add(mBase, uint32(v446)))
	*(*int64)(unsafe.Add(mBase, uint32(v439))) = v451
	v453 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v448))) = v453
	v455 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v446))) = v455
	v458 = v425 + int32(1)
	if v458 != v414 {
		v425 = v458
		goto L86
	} else {
		goto L88
	}
L87:
	;
	goto L85
L88:
	;
	goto L87
L89:
	;
	v487 = v480
	goto L91
L90:
	;
	v487 = v485
	goto L91
L91:
	;
	if v487 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v496 = int32(0)
	goto L95
L93:
	;
	goto L94
L94:
	;
	if base.Ui32(v412) <= base.Ui32(v480) {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v511 = v496 << (uint(int32(4)) % 32)
	v512 = v323 + v511
	v513 = int32(8)
	v514 = v512 + v513
	v515 = *(*int64)(unsafe.Add(mBase, uint32(v514)))
	*(*int64)(unsafe.Add(mBase, uint32(v228))) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(v512)))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v517
	v519 = v511 + (v67 - v487<<(uint(int32(4))%32))
	v521 = v519 + v513
	v522 = *(*int64)(unsafe.Add(mBase, uint32(v521)))
	*(*int64)(unsafe.Add(mBase, uint32(v514))) = v522
	v524 = *(*int64)(unsafe.Add(mBase, uint32(v519)))
	*(*int64)(unsafe.Add(mBase, uint32(v512))) = v524
	v526 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	*(*int64)(unsafe.Add(mBase, uint32(v521))) = v526
	v528 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v519))) = v528
	v531 = v496 + int32(1)
	if v531 != v487 {
		v496 = v531
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	goto L96
L98:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v412) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v480) {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	F_qsort_tuple(m, v23, v412, l2, l3)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L8
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if base.Ui32(v480) < base.Ui32(int32(2)) {
		goto L10
	} else {
		goto L105
	}
L104:
	;
	goto L103
L105:
	;
	v23 = v67 - v480<<(uint(int32(4))%32)
	v24 = v480
	goto L1
L106:
	;
	F_qsort_tuple(m, v67-v480<<(uint(int32(4))%32), v480, l2, l3)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L8
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if base.Ui32(int32(1)) < base.Ui32(v412) {
		v44 = v412
		goto L3
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	goto L10
}
