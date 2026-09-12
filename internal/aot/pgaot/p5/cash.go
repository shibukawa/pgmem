package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cash_div_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(236186), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(490272), int32(725), int32(318620))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
		v30 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v26), base.F64_convert_i64_s(v4)))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_cash_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v5 < v3)
}
func F_cash_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v83 int32
	_ = v83
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v226 int32
	_ = v226
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int64
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v559 int32
	_ = v559
	var v567 int64
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v587 int64
	_ = v587
	var v588 int64
	_ = v588
	var v593 int32
	_ = v593
	var v601 int32
	_ = v601
	var v604 int64
	_ = v604
	var v605 int64
	_ = v605
	var v610 int64
	_ = v610
	var v613 int64
	_ = v613
	var v616 int64
	_ = v616
	var v619 int64
	_ = v619
	var v620 int64
	_ = v620
	var v624 int64
	_ = v624
	var v631 int64
	_ = v631
	var v642 int64
	_ = v642
	var v643 int64
	_ = v643
	var v649 int64
	_ = v649
	var v652 int64
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int64
	_ = v796
	var v797 int64
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v809 int64
	_ = v809
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v826 int32
	_ = v826
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v848 int64
	_ = v848
	var v849 int64
	_ = v849
	var v870 int64
	_ = v870
	var v871 int64
	_ = v871
	var v876 int32
	_ = v876
	var v879 int64
	_ = v879
	var v880 int64
	_ = v880
	var v885 int64
	_ = v885
	var v888 int64
	_ = v888
	var v891 int64
	_ = v891
	var v894 int64
	_ = v894
	var v895 int64
	_ = v895
	var v899 int64
	_ = v899
	var v906 int64
	_ = v906
	var v917 int64
	_ = v917
	var v918 int64
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v942 int64
	_ = v942
	var v960 int64
	_ = v960
	var v965 int32
	_ = v965
	var v988 int32
	_ = v988
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1014 int64
	_ = v1014
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1075 int32
	_ = v1075
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1087 int32
	_ = v1087
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1119 int32
	_ = v1119
	var v1130 int32
	_ = v1130
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1223 int32
	_ = v1223
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1282 int32
	_ = v1282
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1336 int32
	_ = v1336
	var v1344 int32
	_ = v1344
	var v1349 int32
	_ = v1349
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1361 int32
	_ = v1361
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1377 int64
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int64
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	v17 = int64(0)
	v22 = m.G0
	v24 = v22 - int32(112)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_PGLC_localeconv(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+41)))
	v35 = int32(646195)
	v36 = int32(46)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v38 == int32(0) {
		v50 = v35
		v51 = v36
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.Ui32(int32(10)) < base.Ui32(v32) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v41 != 0 {
		v50 = v35
		v51 = v36
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v38&int32(255) == int32(44) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v48 = int32(646159)
	goto L8
L7:
	;
	v48 = int32(646195)
	goto L8
L8:
	;
	v50 = v48
	v51 = base.I32_extend8_s(v38)
	goto L3
L9:
	;
	v53 = int32(2)
	goto L11
L10:
	;
	v53 = v32
	goto L11
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v28)+36))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v62 = v27
	goto L12
L12:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v83-int32(9)))&base.B2i32(v83 != int32(32)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v59&int32(255) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v62 = v62 + int32(1)
	goto L12
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	v99 = v58
	goto L19
L18:
	;
	v99 = int32(666123)
	goto L19
L19:
	;
	if v99&int32(3) == int32(0) {
		v123 = v99
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v156 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L21:
	;
	v156 = v148 - v99
	goto L20
L22:
	;
	v127 = v123
	goto L31
L23:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v107 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v156 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v112 = v99
	goto L27
L27:
	;
	v116 = v112 + int32(1)
	if v116&int32(3) == int32(0) {
		v123 = v116
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v148 = v116
	goto L21
L29:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v121 != 0 {
		v112 = v116
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v136 = int32(-2139062144)
	if (int32(16843008)-v133|v133)&v136 == v136 {
		v127 = v127 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v142 = v127
	goto L34
L33:
	;
	goto L32
L34:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 != 0 {
		v142 = v142 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v148 = v142
	goto L21
L36:
	;
	goto L35
L37:
	;
	if v200 != 0 {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v200 = int32(0)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v162 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v163 = v62
	v164 = v99
	v165 = v156
	v166 = v162
	goto L45
L42:
	;
	v188 = v99
	v192 = int32(0)
	goto L43
L43:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v200 = v192 - v193
	goto L37
L44:
	;
	v188 = v183
	v192 = v185
	goto L43
L45:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	if v166 != v168 {
		v183 = v164
		v185 = v166
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v183 = v177
	v185 = int32(0)
	goto L44
L47:
	;
	if v168 == int32(0) {
		v183 = v164
		v185 = v166
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v173 = v165 - int32(1)
	if v173 == int32(0) {
		v183 = v164
		v185 = v166
		goto L44
	} else {
		goto L49
	}
L49:
	;
	v176 = int32(1)
	v177 = v164 + v176
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v178 != 0 {
		v163 = v163 + v176
		v164 = v177
		v165 = v173
		v166 = v178
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v201 = int32(0)
	goto L53
L52:
	;
	v201 = v156
	goto L53
L53:
	;
	if v57 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v204 = v56
	goto L56
L55:
	;
	v204 = int32(646197)
	goto L56
L56:
	;
	v205 = v62 + v201
	goto L57
L57:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v226-int32(9)))&base.B2i32(v226 != int32(32)) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v238 = int64(-1)
	if v55&int32(255) != 0 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v205 = v205 + int32(1)
	goto L57
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	if v61 != 0 {
		goto L133
	} else {
		goto L134
	}
L63:
	;
	v242 = v54
	goto L65
L64:
	;
	v242 = int32(646185)
	goto L65
L65:
	;
	if v242&int32(3) == int32(0) {
		v266 = v242
		goto L68
	} else {
		goto L69
	}
L66:
	;
	if v299 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L67:
	;
	v299 = v291 - v242
	goto L66
L68:
	;
	v270 = v266
	goto L77
L69:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v250 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v299 = int32(0)
	goto L66
L71:
	;
	goto L72
L72:
	;
	v255 = v242
	goto L73
L73:
	;
	v259 = v255 + int32(1)
	if v259&int32(3) == int32(0) {
		v266 = v259
		goto L68
	} else {
		goto L75
	}
L74:
	;
	v291 = v259
	goto L67
L75:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v264 != 0 {
		v255 = v259
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v279 = int32(-2139062144)
	if (int32(16843008)-v276|v276)&v279 == v279 {
		v270 = v270 + int32(4)
		goto L77
	} else {
		goto L79
	}
L78:
	;
	v285 = v270
	goto L80
L79:
	;
	goto L78
L80:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v289 != 0 {
		v285 = v285 + int32(1)
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v291 = v285
	goto L67
L82:
	;
	goto L81
L83:
	;
	if v343 == int32(0) {
		v454 = v299
		v455 = v238
		goto L62
	} else {
		goto L97
	}
L84:
	;
	v343 = int32(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v305 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v306 = v205
	v307 = v242
	v308 = v299
	v309 = v305
	goto L91
L88:
	;
	v331 = v242
	v335 = int32(0)
	goto L89
L89:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	v343 = v335 - v336
	goto L83
L90:
	;
	v331 = v326
	v335 = v328
	goto L89
L91:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v309 != v311 {
		v326 = v307
		v328 = v309
		goto L90
	} else {
		goto L93
	}
L92:
	;
	v326 = v320
	v328 = int32(0)
	goto L90
L93:
	;
	if v311 == int32(0) {
		v326 = v307
		v328 = v309
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v316 = v308 - int32(1)
	if v316 == int32(0) {
		v326 = v307
		v328 = v309
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v319 = int32(1)
	v320 = v307 + v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v306)+1)))
	if v321 != 0 {
		v306 = v306 + v319
		v307 = v320
		v308 = v316
		v309 = v321
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	if v226 == int32(40) {
		v454 = int32(1)
		v455 = v238
		goto L62
	} else {
		goto L98
	}
L98:
	;
	v349 = int32(0)
	if v204&int32(3) == v349 {
		v373 = v204
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if v406 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L100:
	;
	v406 = v398 - v204
	goto L99
L101:
	;
	v377 = v373
	goto L110
L102:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v357 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v406 = int32(0)
	goto L99
L104:
	;
	goto L105
L105:
	;
	v362 = v204
	goto L106
L106:
	;
	v366 = v362 + int32(1)
	if v366&int32(3) == int32(0) {
		v373 = v366
		goto L101
	} else {
		goto L108
	}
L107:
	;
	v398 = v366
	goto L100
L108:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v371 != 0 {
		v362 = v366
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v386 = int32(-2139062144)
	if (int32(16843008)-v383|v383)&v386 == v386 {
		v377 = v377 + int32(4)
		goto L110
	} else {
		goto L112
	}
L111:
	;
	v392 = v377
	goto L113
L112:
	;
	goto L111
L113:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392))))
	if v396 != 0 {
		v392 = v392 + int32(1)
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v398 = v392
	goto L100
L115:
	;
	goto L114
L116:
	;
	if v450 != 0 {
		goto L130
	} else {
		goto L131
	}
L117:
	;
	v450 = int32(0)
	goto L116
L118:
	;
	goto L119
L119:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v412 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v413 = v205
	v414 = v204
	v415 = v406
	v416 = v412
	goto L124
L121:
	;
	v438 = v204
	v442 = int32(0)
	goto L122
L122:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v438))))
	v450 = v442 - v443
	goto L116
L123:
	;
	v438 = v433
	v442 = v435
	goto L122
L124:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414))))
	if v416 != v418 {
		v433 = v414
		v435 = v416
		goto L123
	} else {
		goto L126
	}
L125:
	;
	v433 = v427
	v435 = int32(0)
	goto L123
L126:
	;
	if v418 == int32(0) {
		v433 = v414
		v435 = v416
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v423 = v415 - int32(1)
	if v423 == int32(0) {
		v433 = v414
		v435 = v416
		goto L123
	} else {
		goto L128
	}
L128:
	;
	v426 = int32(1)
	v427 = v414 + v426
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413)+1)))
	if v428 != 0 {
		v413 = v413 + v426
		v414 = v427
		v415 = v423
		v416 = v428
		goto L124
	} else {
		goto L129
	}
L129:
	;
	goto L125
L130:
	;
	v451 = v349
	goto L132
L131:
	;
	v451 = v406
	goto L132
L132:
	;
	v454 = v451
	v455 = int64(1)
	goto L62
L133:
	;
	v456 = v60
	goto L135
L134:
	;
	v456 = v50
	goto L135
L135:
	;
	v458 = v205 + v454
	goto L136
L136:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v479-int32(9)))&base.B2i32(v479 != int32(32)) == int32(0) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v491 = int32(0)
	if v156 == v491 {
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v458 = v458 + int32(1)
	goto L136
L139:
	;
	goto L140
L140:
	;
	goto L137
L141:
	;
	if v535 != 0 {
		goto L155
	} else {
		goto L156
	}
L142:
	;
	v535 = int32(0)
	goto L141
L143:
	;
	goto L144
L144:
	;
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v497 != 0 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v498 = v458
	v499 = v99
	v500 = v156
	v501 = v497
	goto L149
L146:
	;
	v523 = v99
	v527 = int32(0)
	goto L147
L147:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
	v535 = v527 - v528
	goto L141
L148:
	;
	v523 = v518
	v527 = v520
	goto L147
L149:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499))))
	if v501 != v503 {
		v518 = v499
		v520 = v501
		goto L148
	} else {
		goto L151
	}
L150:
	;
	v518 = v512
	v520 = int32(0)
	goto L148
L151:
	;
	if v503 == int32(0) {
		v518 = v499
		v520 = v501
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v508 = v500 - int32(1)
	if v508 == int32(0) {
		v518 = v499
		v520 = v501
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v511 = int32(1)
	v512 = v499 + v511
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+1)))
	if v513 != 0 {
		v498 = v498 + v511
		v499 = v512
		v500 = v508
		v501 = v513
		goto L149
	} else {
		goto L154
	}
L154:
	;
	goto L150
L155:
	;
	v536 = v491
	goto L157
L156:
	;
	v536 = v156
	goto L157
L157:
	;
	v538 = v458 + v536
	goto L160
L158:
	;
	m.G0 = v24 + int32(112)
	return v1384
L159:
	;
	if v849 < v567 {
		goto L228
	} else {
		goto L229
	}
L160:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
	if base.Ui32(v559-int32(9)) < base.Ui32(int32(5)) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v571 = v538
	v572 = v559
	v574 = int32(0)
	v587 = v17
	v588 = v17
	goto L167
L162:
	;
	goto L161
L163:
	;
	v538 = v538 + int32(1)
	goto L160
L164:
	;
	if v559 == int32(32) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v567 = base.I64_extend8_s(base.I64_extend_i32_u(v53))
	if v559 != 0 {
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v832 = v538
	v848 = v17
	v849 = v17
	goto L159
L167:
	;
	v593 = v572 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v593&int32(255)) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	if base.Ui32(int32(4)) < base.Ui32((v572-int32(53))&int32(255)) {
		v832 = v571
		v848 = v587
		v849 = v588
		goto L159
	} else {
		goto L221
	}
L169:
	;
	goto L168
L170:
	;
	v800 = v793 + int32(1)
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800))))
	if v801 != 0 {
		v571 = v800
		v572 = v801
		v574 = v794
		v587 = v796
		v588 = v797
		goto L167
	} else {
		goto L220
	}
L171:
	;
	if (base.B2i32(v51 != base.I32_extend8_s(v572))|v574)&int32(1) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L172:
	;
	if v574&base.B2i32(v567 <= v588) != 0 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v601 = v24 + int32(96)
	v604 = int64(10)
	v605 = int64(0)
	v610 = int64(32)
	v613 = int64(base.Ui64(v587) >> (uint(v610) % 64))
	v616 = int64(4294967295)
	v619 = v587 & v616
	v620 = v604 * v619
	v624 = int64(base.Ui64(v620)>>(uint(v610)%64)) + v604*v613
	v631 = v619*v605 + v624&v616
	*(*int64)(unsafe.Add(mBase, uint32(v601)+8)) = v587*v605 + v587>>(uint(int64(63))%64)*v604 + v605*v613 + int64(base.Ui64(v624)>>(uint(v610)%64)) + int64(base.Ui64(v631)>>(uint(v610)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v601))) = v620&v616 | v631<<(uint(v610)%64)
	goto L174
L174:
	;
	v642 = *(*int64)(unsafe.Add(mBase, uint32(v24)+104))
	v643 = *(*int64)(unsafe.Add(mBase, uint32(v24)+96))
	if v642 == v643>>(uint(int64(63))%64) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v793 = v571
	v794 = v574
	v796 = v652
	v797 = v588 + base.I64_extend_i32_u(v574)&int64(1)
	goto L170
L176:
	;
	v649 = base.I64_extend_i32_u(v593) & int64(255)
	v652 = v643 - v649
	if base.B2i32(v649 != int64(0)) == base.B2i32(v652 < v643) {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v656 = int32(0)
	v657 = F_errsave_start(m, v26)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L180
	}
L179:
	;
	goto L178
L180:
	;
	if v657 == int32(0) {
		v1384 = v656
		goto L158
	} else {
		goto L181
	}
L181:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = int32(20814)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v27
	F_errmsg(m, int32(187276), v24+int32(80))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errsave_finish(m, v26, int32(490272), int32(293), int32(275506))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v1384 = v656
	goto L158
L185:
	;
	v793 = v571
	v794 = int32(1)
	v796 = v587
	v797 = v588
	goto L170
L186:
	;
	goto L187
L187:
	;
	if v456&int32(3) == int32(0) {
		v712 = v456
		goto L190
	} else {
		goto L191
	}
L188:
	;
	if v745 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L189:
	;
	v745 = v737 - v456
	goto L188
L190:
	;
	v716 = v712
	goto L199
L191:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456))))
	if v696 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v745 = int32(0)
	goto L188
L193:
	;
	goto L194
L194:
	;
	v701 = v456
	goto L195
L195:
	;
	v705 = v701 + int32(1)
	if v705&int32(3) == int32(0) {
		v712 = v705
		goto L190
	} else {
		goto L197
	}
L196:
	;
	v737 = v705
	goto L189
L197:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v710 != 0 {
		v701 = v705
		goto L195
	} else {
		goto L198
	}
L198:
	;
	goto L196
L199:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v716)))
	v725 = int32(-2139062144)
	if (int32(16843008)-v722|v722)&v725 == v725 {
		v716 = v716 + int32(4)
		goto L199
	} else {
		goto L201
	}
L200:
	;
	v731 = v716
	goto L202
L201:
	;
	goto L200
L202:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731))))
	if v735 != 0 {
		v731 = v731 + int32(1)
		goto L202
	} else {
		goto L204
	}
L203:
	;
	v737 = v731
	goto L189
L204:
	;
	goto L203
L205:
	;
	if v789 != 0 {
		goto L169
	} else {
		goto L219
	}
L206:
	;
	v789 = int32(0)
	goto L205
L207:
	;
	goto L208
L208:
	;
	v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v751 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v752 = v571
	v753 = v456
	v754 = v745
	v755 = v751
	goto L213
L210:
	;
	v777 = v456
	v781 = int32(0)
	goto L211
L211:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
	v789 = v781 - v782
	goto L205
L212:
	;
	v777 = v772
	v781 = v774
	goto L211
L213:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	if v755 != v757 {
		v772 = v753
		v774 = v755
		goto L212
	} else {
		goto L215
	}
L214:
	;
	v772 = v766
	v774 = int32(0)
	goto L212
L215:
	;
	if v757 == int32(0) {
		v772 = v753
		v774 = v755
		goto L212
	} else {
		goto L216
	}
L216:
	;
	v762 = v754 - int32(1)
	if v762 == int32(0) {
		v772 = v753
		v774 = v755
		goto L212
	} else {
		goto L217
	}
L217:
	;
	v765 = int32(1)
	v766 = v753 + v765
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+1)))
	if v767 != 0 {
		v752 = v752 + v765
		v753 = v766
		v754 = v762
		v755 = v767
		goto L213
	} else {
		goto L218
	}
L218:
	;
	goto L214
L219:
	;
	v793 = v571 + v745 - int32(1)
	v794 = v574
	v796 = v587
	v797 = v588
	goto L170
L220:
	;
	v832 = v800
	v848 = v796
	v849 = v797
	goto L159
L221:
	;
	v809 = v587 - int64(1)
	if v809 < v587 {
		v832 = v571
		v848 = v809
		v849 = v588
		goto L159
	} else {
		goto L222
	}
L222:
	;
	v811 = int32(0)
	v812 = F_errsave_start(m, v26)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	if v812 == int32(0) {
		v1384 = v811
		goto L158
	} else {
		goto L224
	}
L224:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = int32(20814)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v27
	F_errmsg(m, int32(187276), v24-int32(-64))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_errsave_finish(m, v26, int32(490272), int32(318), int32(275506))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v1384 = v811
	goto L158
L228:
	;
	v870 = v848
	v871 = v849
	goto L231
L229:
	;
	v960 = v848
	goto L230
L230:
	;
	v965 = v832
	goto L243
L231:
	;
	v876 = v24 + int32(48)
	v879 = int64(10)
	v880 = int64(0)
	v885 = int64(32)
	v888 = int64(base.Ui64(v870) >> (uint(v885) % 64))
	v891 = int64(4294967295)
	v894 = v870 & v891
	v895 = v879 * v894
	v899 = int64(base.Ui64(v895)>>(uint(v885)%64)) + v879*v888
	v906 = v894*v880 + v899&v891
	*(*int64)(unsafe.Add(mBase, uint32(v876)+8)) = v870*v880 + v870>>(uint(int64(63))%64)*v879 + v880*v888 + int64(base.Ui64(v899)>>(uint(v885)%64)) + int64(base.Ui64(v906)>>(uint(v885)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v876))) = v895&v891 | v906<<(uint(v885)%64)
	goto L233
L232:
	;
	v960 = v918
	goto L230
L233:
	;
	v917 = *(*int64)(unsafe.Add(mBase, uint32(v24)+56))
	v918 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
	if v917 != v918>>(uint(int64(63))%64) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v922 = int32(0)
	v923 = F_errsave_start(m, v26)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v942 = v871 + int64(1)
	if v942 != v567 {
		v870 = v918
		v871 = v942
		goto L231
	} else {
		goto L242
	}
L237:
	;
	if v923 == int32(0) {
		v1384 = v922
		goto L158
	} else {
		goto L238
	}
L238:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(20814)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27
	F_errmsg(m, int32(187276), v24)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	F_errsave_finish(m, v26, int32(490272), int32(328), int32(275506))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1384 = v922
	goto L158
L242:
	;
	goto L232
L243:
	;
	v988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v965))))
	if base.Ui32((v988-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v965 = v965 + int32(1)
		goto L243
	} else {
		goto L245
	}
L244:
	;
	v996 = v965
	v998 = v988
	v1014 = v455
	goto L246
L245:
	;
	goto L244
L246:
	;
	switch v998 & int32(255) {
	case 0:
		goto L249
	default:
		goto L250
	case 9, 10, 11, 12, 13, 32, 41:
		v1380 = int32(1)
		v1381 = v1014
		goto L248
	}
L248:
	;
	v1382 = v1380 + v996
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382))))
	v996 = v1382
	v998 = v1383
	v1014 = v1381
	goto L246
L249:
	;
	if int64(0) < v1014 {
		goto L354
	} else {
		goto L355
	}
L250:
	;
	if v242&int32(3) == int32(0) {
		v1042 = v242
		goto L253
	} else {
		goto L254
	}
L251:
	;
	if v1075 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L252:
	;
	v1075 = v1067 - v242
	goto L251
L253:
	;
	v1046 = v1042
	goto L262
L254:
	;
	v1026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242))))
	if v1026 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1075 = int32(0)
	goto L251
L256:
	;
	goto L257
L257:
	;
	v1031 = v242
	goto L258
L258:
	;
	v1035 = v1031 + int32(1)
	if v1035&int32(3) == int32(0) {
		v1042 = v1035
		goto L253
	} else {
		goto L260
	}
L259:
	;
	v1067 = v1035
	goto L252
L260:
	;
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	if v1040 != 0 {
		v1031 = v1035
		goto L258
	} else {
		goto L261
	}
L261:
	;
	goto L259
L262:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1046)))
	v1055 = int32(-2139062144)
	if (int32(16843008)-v1052|v1052)&v1055 == v1055 {
		v1046 = v1046 + int32(4)
		goto L262
	} else {
		goto L264
	}
L263:
	;
	v1061 = v1046
	goto L265
L264:
	;
	goto L263
L265:
	;
	v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1061))))
	if v1065 != 0 {
		v1061 = v1061 + int32(1)
		goto L265
	} else {
		goto L267
	}
L266:
	;
	v1067 = v1061
	goto L252
L267:
	;
	goto L266
L268:
	;
	if v1119 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L269:
	;
	v1119 = int32(0)
	goto L268
L270:
	;
	goto L271
L271:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	if v1081 != 0 {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1082 = v996
	v1083 = v242
	v1084 = v1075
	v1085 = v1081
	goto L276
L273:
	;
	v1107 = v242
	v1111 = int32(0)
	goto L274
L274:
	;
	v1112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107))))
	v1119 = v1111 - v1112
	goto L268
L275:
	;
	v1107 = v1102
	v1111 = v1104
	goto L274
L276:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083))))
	if v1085 != v1087 {
		v1102 = v1083
		v1104 = v1085
		goto L275
	} else {
		goto L278
	}
L277:
	;
	v1102 = v1096
	v1104 = int32(0)
	goto L275
L278:
	;
	if v1087 == int32(0) {
		v1102 = v1083
		v1104 = v1085
		goto L275
	} else {
		goto L279
	}
L279:
	;
	v1092 = v1084 - int32(1)
	if v1092 == int32(0) {
		v1102 = v1083
		v1104 = v1085
		goto L275
	} else {
		goto L280
	}
L280:
	;
	v1095 = int32(1)
	v1096 = v1083 + v1095
	v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+1)))
	if v1097 != 0 {
		v1082 = v1082 + v1095
		v1083 = v1096
		v1084 = v1092
		v1085 = v1097
		goto L276
	} else {
		goto L281
	}
L281:
	;
	goto L277
L282:
	;
	v1380 = v1075
	v1381 = int64(-1)
	goto L248
L283:
	;
	goto L284
L284:
	;
	if v204&int32(3) == int32(0) {
		v1146 = v204
		goto L287
	} else {
		goto L288
	}
L285:
	;
	if v1179 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L286:
	;
	v1179 = v1171 - v204
	goto L285
L287:
	;
	v1150 = v1146
	goto L296
L288:
	;
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v1130 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L289:
	;
	v1179 = int32(0)
	goto L285
L290:
	;
	goto L291
L291:
	;
	v1135 = v204
	goto L292
L292:
	;
	v1139 = v1135 + int32(1)
	if v1139&int32(3) == int32(0) {
		v1146 = v1139
		goto L287
	} else {
		goto L294
	}
L293:
	;
	v1171 = v1139
	goto L286
L294:
	;
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1139))))
	if v1144 != 0 {
		v1135 = v1139
		goto L292
	} else {
		goto L295
	}
L295:
	;
	goto L293
L296:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1150)))
	v1159 = int32(-2139062144)
	if (int32(16843008)-v1156|v1156)&v1159 == v1159 {
		v1150 = v1150 + int32(4)
		goto L296
	} else {
		goto L298
	}
L297:
	;
	v1165 = v1150
	goto L299
L298:
	;
	goto L297
L299:
	;
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165))))
	if v1169 != 0 {
		v1165 = v1165 + int32(1)
		goto L299
	} else {
		goto L301
	}
L300:
	;
	v1171 = v1165
	goto L286
L301:
	;
	goto L300
L302:
	;
	if v1223 == int32(0) {
		v1380 = v1179
		v1381 = v1014
		goto L248
	} else {
		goto L316
	}
L303:
	;
	v1223 = int32(0)
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	if v1185 != 0 {
		goto L306
	} else {
		goto L307
	}
L306:
	;
	v1186 = v996
	v1187 = v204
	v1188 = v1179
	v1189 = v1185
	goto L310
L307:
	;
	v1211 = v204
	v1215 = int32(0)
	goto L308
L308:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
	v1223 = v1215 - v1216
	goto L302
L309:
	;
	v1211 = v1206
	v1215 = v1208
	goto L308
L310:
	;
	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187))))
	if v1189 != v1191 {
		v1206 = v1187
		v1208 = v1189
		goto L309
	} else {
		goto L312
	}
L311:
	;
	v1206 = v1200
	v1208 = int32(0)
	goto L309
L312:
	;
	if v1191 == int32(0) {
		v1206 = v1187
		v1208 = v1189
		goto L309
	} else {
		goto L313
	}
L313:
	;
	v1196 = v1188 - int32(1)
	if v1196 == int32(0) {
		v1206 = v1187
		v1208 = v1189
		goto L309
	} else {
		goto L314
	}
L314:
	;
	v1199 = int32(1)
	v1200 = v1187 + v1199
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186)+1)))
	if v1201 != 0 {
		v1186 = v1186 + v1199
		v1187 = v1200
		v1188 = v1196
		v1189 = v1201
		goto L310
	} else {
		goto L315
	}
L315:
	;
	goto L311
L316:
	;
	if v99&int32(3) == int32(0) {
		v1249 = v99
		goto L319
	} else {
		goto L320
	}
L317:
	;
	if v1282 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L318:
	;
	v1282 = v1274 - v99
	goto L317
L319:
	;
	v1253 = v1249
	goto L328
L320:
	;
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v1233 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1282 = int32(0)
	goto L317
L322:
	;
	goto L323
L323:
	;
	v1238 = v99
	goto L324
L324:
	;
	v1242 = v1238 + int32(1)
	if v1242&int32(3) == int32(0) {
		v1249 = v1242
		goto L319
	} else {
		goto L326
	}
L325:
	;
	v1274 = v1242
	goto L318
L326:
	;
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1242))))
	if v1247 != 0 {
		v1238 = v1242
		goto L324
	} else {
		goto L327
	}
L327:
	;
	goto L325
L328:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1253)))
	v1262 = int32(-2139062144)
	if (int32(16843008)-v1259|v1259)&v1262 == v1262 {
		v1253 = v1253 + int32(4)
		goto L328
	} else {
		goto L330
	}
L329:
	;
	v1268 = v1253
	goto L331
L330:
	;
	goto L329
L331:
	;
	v1272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1268))))
	if v1272 != 0 {
		v1268 = v1268 + int32(1)
		goto L331
	} else {
		goto L333
	}
L332:
	;
	v1274 = v1268
	goto L318
L333:
	;
	goto L332
L334:
	;
	if v1326 == int32(0) {
		v1380 = v1282
		v1381 = v1014
		goto L248
	} else {
		goto L348
	}
L335:
	;
	v1326 = int32(0)
	goto L334
L336:
	;
	goto L337
L337:
	;
	v1288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996))))
	if v1288 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	v1289 = v996
	v1290 = v99
	v1291 = v1282
	v1292 = v1288
	goto L342
L339:
	;
	v1314 = v99
	v1318 = int32(0)
	goto L340
L340:
	;
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1314))))
	v1326 = v1318 - v1319
	goto L334
L341:
	;
	v1314 = v1309
	v1318 = v1311
	goto L340
L342:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290))))
	if v1292 != v1294 {
		v1309 = v1290
		v1311 = v1292
		goto L341
	} else {
		goto L344
	}
L343:
	;
	v1309 = v1303
	v1311 = int32(0)
	goto L341
L344:
	;
	if v1294 == int32(0) {
		v1309 = v1290
		v1311 = v1292
		goto L341
	} else {
		goto L345
	}
L345:
	;
	v1299 = v1291 - int32(1)
	if v1299 == int32(0) {
		v1309 = v1290
		v1311 = v1292
		goto L341
	} else {
		goto L346
	}
L346:
	;
	v1302 = int32(1)
	v1303 = v1290 + v1302
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1289)+1)))
	if v1304 != 0 {
		v1289 = v1289 + v1302
		v1290 = v1303
		v1291 = v1299
		v1292 = v1304
		goto L342
	} else {
		goto L347
	}
L347:
	;
	goto L343
L348:
	;
	v1329 = int32(0)
	v1330 = F_errsave_start(m, v26)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	if v1330 == int32(0) {
		v1384 = v1329
		goto L158
	} else {
		goto L350
	}
L350:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = int32(20814)
	F_errmsg(m, int32(701283), v24+int32(16))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	F_errsave_finish(m, v26, int32(490272), int32(355), int32(275506))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v1384 = v1329
	goto L158
L354:
	;
	if v960 == int64(-9223372036854775807-1) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	v1377 = v960
	goto L356
L356:
	;
	v1378 = F_Int64GetDatum(m, v1377)
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L1
	} else {
		goto L365
	}
L357:
	;
	v1354 = int32(0)
	v1355 = F_errsave_start(m, v26)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L1
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1377 = int64(0) - v960
	goto L356
L360:
	;
	if v1355 == int32(0) {
		v1384 = v1354
		goto L158
	} else {
		goto L361
	}
L361:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L362
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(20814)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v27
	F_errmsg(m, int32(187276), v24+int32(32))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	F_errsave_finish(m, v26, int32(490272), int32(368), int32(275506))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L364
	}
L364:
	;
	v1384 = v1354
	goto L158
L365:
	;
	v1384 = v1378
	goto L158
}
func F_cash_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	return base.B2i32(v3 <= v5)
}
func F_cash_mul_flt8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = F_cash_mul_float8(m, v3, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_Int64GetDatum(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_cash_words(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	F_initStringInfo(m, v9)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 < int64(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, v9, int32(709438))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v24 = v12
	goto L5
L5:
	;
	v26 = base.I64_div_u_s(v24, int64(100))
	v28 = base.I64_rem_u_s(v26, int64(1000))
	v30 = base.I64_div_u_s(v24, int64(100000000000000))
	v33 = base.I32_rem_u_s(base.I32_wrap_i64(v30), int32(1000))
	if base.Ui64(int64(100000000000000000)) <= base.Ui64(v24) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v24 = int64(0) - v12
	goto L5
L7:
	;
	v37 = base.I64_div_u_s(v24, int64(100000000000000000))
	F_append_num_word(m, v9, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_appendStringInfoString(m, v9, int32(715491))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	F_append_num_word(m, v9, base.I64_extend_i32_u(v33))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = base.I64_div_u_s(v24, int64(100000000000))
	v53 = base.I32_rem_u_s(base.I32_wrap_i64(v50), int32(1000))
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	F_appendStringInfoString(m, v9, int32(715480))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_append_num_word(m, v9, base.I64_extend_i32_u(v53))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v61 = base.I64_div_u_s(v24, int64(100000000))
	v63 = base.I64_rem_u_s(v61, int64(1000))
	if v63 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_appendStringInfoString(m, v9, int32(715515))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_append_num_word(m, v9, v63)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v72 = base.I64_div_u_s(v24, int64(100000))
	v74 = base.I64_rem_u_s(v72, int64(1000))
	if v74 != int64(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_appendStringInfoString(m, v9, int32(715505))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_append_num_word(m, v9, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v28 != int64(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_appendStringInfoString(m, v9, int32(719265))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_append_num_word(m, v9, v28)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if base.Ui64(v24) <= base.Ui64(int64(99)) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_appendStringInfoString(m, v9, int32(236920))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v93 = int64(100)
	if base.Ui64(v24-v93) < base.Ui64(v93) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v97 = int32(719356)
	goto L42
L41:
	;
	v97 = int32(719342)
	goto L42
L42:
	;
	F_appendStringInfoString(m, v9, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v102 = v24 - v26*int64(100)
	F_append_num_word(m, v9, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v102 == int64(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v109 = int32(95610)
	goto L47
L46:
	;
	v109 = int32(121476)
	goto L47
L47:
	;
	F_appendStringInfoString(m, v9, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if base.Ui32((v113-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v125))) = uint8(v124)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v129 = F_cstring_to_text_with_len(m, v127, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v122 = v113 - int32(32)
	goto L52
L51:
	;
	v122 = v113
	goto L52
L52:
	;
	v124 = v122 & int32(255)
	goto L49
L53:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	F_pfree(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v9 + int32(16)
	return v129
}
