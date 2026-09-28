package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_qsort_arg_entries_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = base.I32_wrap_i64(v10)
		if v14 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v17)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v51 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			v53 = F_FunctionCall2Coll(m, v49, v50, v51, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v55 = base.I32_wrap_i64(v53)
				if v55 == int32(0) {
					v58 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v58)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v65 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					v66 = F_FunctionCall2Coll(m, v62, v63, v64, v65)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = base.I32_wrap_i64(v66)
						if v68 == int32(0) {
							v71 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v71)
						} else {
						}
						if v68 < int32(0) {
							v75 = l0
						} else {
							v75 = l2
						}
						v76 = v75
						return v76
					}
				} else {
					if int32(0) < v55 {
						v76 = l1
						return v76
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						v65 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						v66 = F_FunctionCall2Coll(m, v62, v63, v64, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = base.I32_wrap_i64(v66)
							if v68 == int32(0) {
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v71)
							} else {
							}
							if v68 < int32(0) {
								v75 = l0
							} else {
								v75 = l2
							}
							v76 = v75
							return v76
						}
					}
				}
			}
		} else {
			if int32(0) <= v14 {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v51 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v52 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				v53 = F_FunctionCall2Coll(m, v49, v50, v51, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = base.I32_wrap_i64(v53)
					if v55 == int32(0) {
						v58 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v58)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						v65 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						v66 = F_FunctionCall2Coll(m, v62, v63, v64, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = base.I32_wrap_i64(v66)
							if v68 == int32(0) {
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v71)
							} else {
							}
							if v68 < int32(0) {
								v75 = l0
							} else {
								v75 = l2
							}
							v76 = v75
							return v76
						}
					} else {
						if int32(0) < v55 {
							v76 = l1
							return v76
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
							v64 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							v65 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							v66 = F_FunctionCall2Coll(m, v62, v63, v64, v65)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = base.I32_wrap_i64(v66)
								if v68 == int32(0) {
									v71 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v71)
								} else {
								}
								if v68 < int32(0) {
									v75 = l0
								} else {
									v75 = l2
								}
								v76 = v75
								return v76
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				v24 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				v25 = F_FunctionCall2Coll(m, v21, v22, v23, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = base.I32_wrap_i64(v25)
					if v27 == int32(0) {
						v30 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v30)
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						v37 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						v38 = F_FunctionCall2Coll(m, v34, v35, v36, v37)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = base.I32_wrap_i64(v38)
							if v40 == int32(0) {
								v43 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v43)
							} else {
							}
							if v40 < int32(0) {
								v47 = l2
							} else {
								v47 = l0
							}
							return v47
						}
					} else {
						if v27 < int32(0) {
							v76 = l1
							return v76
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
							v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
							v38 = F_FunctionCall2Coll(m, v34, v35, v36, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = base.I32_wrap_i64(v38)
								if v40 == int32(0) {
									v43 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v43)
								} else {
								}
								if v40 < int32(0) {
									v47 = l2
								} else {
									v47 = l0
								}
								return v47
							}
						}
					}
				}
			}
		}
	}
}
func F_qsort_interruptible(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v57 int32
	_ = v57
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v494 int32
	_ = v494
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
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
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v700 int32
	_ = v700
	var v723 int32
	_ = v723
	var v730 int32
	_ = v730
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v787 int32
	_ = v787
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v931 int32
	_ = v931
	var v954 int32
	_ = v954
	var v961 int32
	_ = v961
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
	var v980 int32
	_ = v980
	var v1015 int32
	_ = v1015
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1139 int32
	_ = v1139
	var v1162 int32
	_ = v1162
	var v1170 int32
	_ = v1170
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1298 int32
	_ = v1298
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1451 int32
	_ = v1451
	var v1474 int32
	_ = v1474
	var v1481 int32
	_ = v1481
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	v32 = l0
	v33 = l1
	v34 = l2
	v35 = l3
	v36 = l4
	v49 = l2 & int32(-4)
	v50 = l2 & int32(3)
	v51 = l2 - int32(1)
	v52 = int32(0) - l2
	goto L1
L1:
	;
	v55 = v32 + v34
	v57 = v33
	goto L3
L3:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_interruptible[0]))
	if v80 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(v57) <= base.Ui32(int32(6)) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v383 = v32 + int32(base.Ui32(v57)>>(uint(int32(1))%32))*v34
	if v57 != int32(7) {
		goto L49
	} else {
		goto L50
	}
L11:
	;
	return
L12:
	;
	if base.Ui32(v57) < base.Ui32(int32(2)) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v320 = v57 * v34
	if base.Ui32(v320) <= base.Ui32(v34) {
		goto L11
	} else {
		goto L39
	}
L15:
	;
	v87 = v57 * v34
	if base.Ui32(v87) <= base.Ui32(v34) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v93 = v34 & int32(3)
	v115 = v55
	goto L17
L17:
	;
	if base.Ui32(v115) <= base.Ui32(v32) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L11
L19:
	;
	v318 = v34 + v115
	if base.Ui32(v318) < base.Ui32(v32+v87) {
		v115 = v318
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v124 = v115
	goto L21
L21:
	;
	v141 = v124 + v52
	v142 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, v141, v124, v36)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L8
	} else {
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	if v142 <= int32(0) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	if v34 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if base.Ui32(v32) < base.Ui32(v141) {
		v124 = v141
		goto L21
	} else {
		goto L37
	}
L26:
	;
	v148 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v161 = v148
	v167 = v148
	goto L30
L28:
	;
	v219 = v148
	goto L29
L29:
	;
	v242 = v219
	v247 = v148
	goto L34
L30:
	;
	v177 = v124 + v161
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v179 = v141 + v161
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v180)
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v178)
	v184 = v161 | int32(1)
	v185 = v124 + v184
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	v187 = v184 + v141
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v188)
	*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v186)
	v192 = v161 | int32(2)
	v193 = v124 + v192
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v195 = v192 + v141
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v196)
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v194)
	v200 = v161 | int32(3)
	v201 = v124 + v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	v203 = v200 + v141
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v204)
	*(*uint8)(unsafe.Add(mBase, uint32(v203))) = uint8(v202)
	v207 = int32(4)
	v208 = v161 + v207
	v210 = v167 + v207
	if v210 != v34&int32(-4) {
		v161 = v208
		v167 = v210
		goto L30
	} else {
		goto L32
	}
L31:
	;
	if v93 == int32(0) {
		goto L25
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v219 = v208
	goto L29
L34:
	;
	v260 = v124 + v242
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v262 = v242 + v141
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	*(*uint8)(unsafe.Add(mBase, uint32(v260))) = uint8(v263)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v261)
	v266 = int32(1)
	v269 = v247 + v266
	if v269 != v93 {
		v242 = v242 + v266
		v247 = v269
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L25
L36:
	;
	goto L35
L37:
	;
	goto L22
L38:
	;
	goto L18
L39:
	;
	v322 = v32 + v320
	v330 = v55
	goto L40
L40:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_interruptible[0]))
	if v347 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L11
L42:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v351 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, v330+v52, v330, v36)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	if int32(0) < v351 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	v355 = v34 + v330
	if base.Ui32(v355) < base.Ui32(v322) {
		v330 = v355
		goto L40
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	v389 = v32 + (v57-int32(1))*v34
	if base.Ui32(v57) < base.Ui32(int32(41)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v417 = v383
	goto L51
L51:
	;
	if v34 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L52:
	;
	v414 = F_qsort_interruptible_med3(m, v412, v410, v411, v35, v36)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L8
	} else {
		goto L59
	}
L53:
	;
	v410 = v383
	v411 = v389
	v412 = v32
	goto L52
L54:
	;
	goto L55
L55:
	;
	v394 = int32(base.Ui32(v57)>>(uint(int32(3))%32)) * v34
	v397 = v394 << (uint(int32(1)) % 32)
	v399 = F_qsort_interruptible_med3(m, v32, v32+v394, v32+v397, v35, v36)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v403 = F_qsort_interruptible_med3(m, v383-v394, v383, v394+v383, v35, v36)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v407 = F_qsort_interruptible_med3(m, v389-v397, v389-v394, v389, v35, v36)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v410 = v403
	v411 = v407
	v412 = v399
	goto L52
L59:
	;
	v417 = v414
	goto L51
L60:
	;
	v572 = v32 + (v57-int32(1))*v34
	v580 = v572
	v582 = v572
	v583 = v55
	v585 = v55
	goto L72
L61:
	;
	v423 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v436 = v423
	v439 = v423
	goto L65
L63:
	;
	v494 = v423
	goto L64
L64:
	;
	v517 = v494
	v521 = v423
	goto L69
L65:
	;
	v452 = v32 + v436
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	v454 = v417 + v436
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v454))))
	*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v455)
	*(*uint8)(unsafe.Add(mBase, uint32(v454))) = uint8(v453)
	v459 = v436 | int32(1)
	v460 = v32 + v459
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v462 = v459 + v417
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v463)
	*(*uint8)(unsafe.Add(mBase, uint32(v462))) = uint8(v461)
	v467 = v436 | int32(2)
	v468 = v32 + v467
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v470 = v467 + v417
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v471)
	*(*uint8)(unsafe.Add(mBase, uint32(v470))) = uint8(v469)
	v475 = v436 | int32(3)
	v476 = v32 + v475
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	v478 = v475 + v417
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v478))))
	*(*uint8)(unsafe.Add(mBase, uint32(v476))) = uint8(v479)
	*(*uint8)(unsafe.Add(mBase, uint32(v478))) = uint8(v477)
	v482 = int32(4)
	v483 = v436 + v482
	v485 = v439 + v482
	if v485 != v49 {
		v436 = v483
		v439 = v485
		goto L65
	} else {
		goto L67
	}
L66:
	;
	if v50 == int32(0) {
		goto L60
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v494 = v483
	goto L64
L69:
	;
	v535 = v32 + v517
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v535))))
	v537 = v517 + v417
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	*(*uint8)(unsafe.Add(mBase, uint32(v535))) = uint8(v538)
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v536)
	v541 = int32(1)
	v544 = v521 + v541
	if v544 != v50 {
		v517 = v517 + v541
		v521 = v544
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L60
L71:
	;
	goto L70
L72:
	;
	if base.Ui32(v580) < base.Ui32(v583) {
		v814 = v583
		v816 = v585
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v1528 = base.I32_div_u_s(v1214, v34)
	F_qsort_interruptible(m, v322-v1214, v1528, v34, v35, v36)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L8
	} else {
		goto L173
	}
L74:
	;
	if base.Ui32(v814) <= base.Ui32(v580) {
		goto L102
	} else {
		goto L103
	}
L75:
	;
	v607 = v583
	v609 = v585
	goto L76
L76:
	;
	v620 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, v607, v32, v36)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L8
	} else {
		goto L78
	}
L77:
	;
	v814 = v802
	v816 = v787
	goto L74
L78:
	;
	if int32(0) < v620 {
		v814 = v607
		v816 = v609
		goto L74
	} else {
		goto L79
	}
L79:
	;
	if v620 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v34 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v787 = v609
	goto L82
L82:
	;
	v799 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_interruptible[0]))
	if v799 != 0 {
		goto L95
	} else {
		goto L96
	}
L83:
	;
	v787 = v34 + v609
	goto L82
L84:
	;
	v628 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v635 = v628
	v639 = v628
	goto L88
L86:
	;
	v700 = v628
	goto L87
L87:
	;
	v723 = v700
	v730 = v628
	goto L92
L88:
	;
	v657 = v639 + v609
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657))))
	v659 = v639 + v607
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v659))))
	*(*uint8)(unsafe.Add(mBase, uint32(v657))) = uint8(v660)
	*(*uint8)(unsafe.Add(mBase, uint32(v659))) = uint8(v658)
	v664 = v639 | int32(1)
	v665 = v609 + v664
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	v667 = v664 + v607
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667))))
	*(*uint8)(unsafe.Add(mBase, uint32(v665))) = uint8(v668)
	*(*uint8)(unsafe.Add(mBase, uint32(v667))) = uint8(v666)
	v672 = v639 | int32(2)
	v673 = v609 + v672
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	v675 = v672 + v607
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v675))))
	*(*uint8)(unsafe.Add(mBase, uint32(v673))) = uint8(v676)
	*(*uint8)(unsafe.Add(mBase, uint32(v675))) = uint8(v674)
	v680 = v639 | int32(3)
	v681 = v609 + v680
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v681))))
	v683 = v680 + v607
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683))))
	*(*uint8)(unsafe.Add(mBase, uint32(v681))) = uint8(v684)
	*(*uint8)(unsafe.Add(mBase, uint32(v683))) = uint8(v682)
	v687 = int32(4)
	v688 = v639 + v687
	v690 = v635 + v687
	if v690 != v49 {
		v635 = v690
		v639 = v688
		goto L88
	} else {
		goto L90
	}
L89:
	;
	if v50 == int32(0) {
		goto L83
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v700 = v688
	goto L87
L92:
	;
	v740 = v723 + v609
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v740))))
	v742 = v723 + v607
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742))))
	*(*uint8)(unsafe.Add(mBase, uint32(v740))) = uint8(v743)
	*(*uint8)(unsafe.Add(mBase, uint32(v742))) = uint8(v741)
	v746 = int32(1)
	v749 = v730 + v746
	if v749 != v50 {
		v723 = v723 + v746
		v730 = v749
		goto L92
	} else {
		goto L94
	}
L93:
	;
	goto L83
L94:
	;
	goto L93
L95:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L8
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v802 = v34 + v607
	if base.Ui32(v802) <= base.Ui32(v580) {
		v607 = v802
		v609 = v787
		goto L76
	} else {
		goto L99
	}
L98:
	;
	goto L97
L99:
	;
	goto L77
L100:
	;
	goto L73
L101:
	;
	if v34 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L102:
	;
	v835 = v580
	v837 = v582
	goto L105
L103:
	;
	v1042 = v580
	v1044 = v582
	goto L104
L104:
	;
	v1058 = v816 - v32
	v1059 = v814 - v816
	if v1058 < v1059 {
		goto L130
	} else {
		goto L131
	}
L105:
	;
	v851 = m.T0[v35].(func(*base.Module, int32, int32, int32) int32)(m, v835, v32, v36)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L8
	} else {
		goto L107
	}
L106:
	;
	v1042 = v1033
	v1044 = v1015
	goto L104
L107:
	;
	if v851 < int32(0) {
		goto L101
	} else {
		goto L108
	}
L108:
	;
	if v851 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v34 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v1015 = v837
	goto L111
L111:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, _c_F_qsort_interruptible[0]))
	if v1030 != 0 {
		goto L124
	} else {
		goto L125
	}
L112:
	;
	v1015 = v837 + v52
	goto L111
L113:
	;
	v859 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v866 = v859
	v870 = v859
	goto L117
L115:
	;
	v931 = v859
	goto L116
L116:
	;
	v954 = v931
	v961 = v859
	goto L121
L117:
	;
	v888 = v870 + v835
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888))))
	v890 = v870 + v837
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890))))
	*(*uint8)(unsafe.Add(mBase, uint32(v888))) = uint8(v891)
	*(*uint8)(unsafe.Add(mBase, uint32(v890))) = uint8(v889)
	v895 = v870 | int32(1)
	v896 = v835 + v895
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v896))))
	v898 = v895 + v837
	v899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v898))))
	*(*uint8)(unsafe.Add(mBase, uint32(v896))) = uint8(v899)
	*(*uint8)(unsafe.Add(mBase, uint32(v898))) = uint8(v897)
	v903 = v870 | int32(2)
	v904 = v835 + v903
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904))))
	v906 = v903 + v837
	v907 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906))))
	*(*uint8)(unsafe.Add(mBase, uint32(v904))) = uint8(v907)
	*(*uint8)(unsafe.Add(mBase, uint32(v906))) = uint8(v905)
	v911 = v870 | int32(3)
	v912 = v835 + v911
	v913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v912))))
	v914 = v911 + v837
	v915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v914))))
	*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v915)
	*(*uint8)(unsafe.Add(mBase, uint32(v914))) = uint8(v913)
	v918 = int32(4)
	v919 = v870 + v918
	v921 = v866 + v918
	if v921 != v49 {
		v866 = v921
		v870 = v919
		goto L117
	} else {
		goto L119
	}
L118:
	;
	if v50 == int32(0) {
		goto L112
	} else {
		goto L120
	}
L119:
	;
	goto L118
L120:
	;
	v931 = v919
	goto L116
L121:
	;
	v971 = v954 + v835
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971))))
	v973 = v954 + v837
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v973))))
	*(*uint8)(unsafe.Add(mBase, uint32(v971))) = uint8(v974)
	*(*uint8)(unsafe.Add(mBase, uint32(v973))) = uint8(v972)
	v977 = int32(1)
	v980 = v961 + v977
	if v980 != v50 {
		v954 = v954 + v977
		v961 = v980
		goto L121
	} else {
		goto L123
	}
L122:
	;
	goto L112
L123:
	;
	goto L122
L124:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L8
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v1033 = v835 + v52
	if base.Ui32(v814) <= base.Ui32(v1033) {
		v835 = v1033
		v837 = v1015
		goto L105
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	goto L106
L129:
	;
	v1214 = v1044 - v1042
	v1216 = v322 - (v34 + v1044)
	if base.Ui32(v1214) < base.Ui32(v1216) {
		goto L145
	} else {
		goto L146
	}
L130:
	;
	v1061 = v1058
	goto L132
L131:
	;
	v1061 = v1059
	goto L132
L132:
	;
	if v1061 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v1064 = v814 - v1061
	v1066 = v1061 & int32(3)
	v1067 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1061) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v1075 = int32(0)
	v1079 = v1067
	goto L137
L135:
	;
	v1139 = v1067
	goto L136
L136:
	;
	v1162 = v1139
	v1170 = v1067
	goto L141
L137:
	;
	v1097 = v32 + v1079
	v1098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1097))))
	v1099 = v1079 + v1064
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1097))) = uint8(v1100)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099))) = uint8(v1098)
	v1104 = v1079 | int32(1)
	v1105 = v32 + v1104
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105))))
	v1107 = v1064 + v1104
	v1108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1107))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1105))) = uint8(v1108)
	*(*uint8)(unsafe.Add(mBase, uint32(v1107))) = uint8(v1106)
	v1112 = v1079 | int32(2)
	v1113 = v32 + v1112
	v1114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1113))))
	v1115 = v1064 + v1112
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1115))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1113))) = uint8(v1116)
	*(*uint8)(unsafe.Add(mBase, uint32(v1115))) = uint8(v1114)
	v1120 = v1079 | int32(3)
	v1121 = v32 + v1120
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1121))))
	v1123 = v1064 + v1120
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1121))) = uint8(v1124)
	*(*uint8)(unsafe.Add(mBase, uint32(v1123))) = uint8(v1122)
	v1127 = int32(4)
	v1128 = v1079 + v1127
	v1130 = v1075 + v1127
	if v1130 != v1061&int32(-4) {
		v1075 = v1130
		v1079 = v1128
		goto L137
	} else {
		goto L139
	}
L138:
	;
	if v1066 == int32(0) {
		goto L129
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v1139 = v1128
	goto L136
L141:
	;
	v1180 = v32 + v1162
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1180))))
	v1182 = v1162 + v1064
	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1182))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1180))) = uint8(v1183)
	*(*uint8)(unsafe.Add(mBase, uint32(v1182))) = uint8(v1181)
	v1186 = int32(1)
	v1189 = v1170 + v1186
	if v1189 != v1066 {
		v1162 = v1162 + v1186
		v1170 = v1189
		goto L141
	} else {
		goto L143
	}
L142:
	;
	goto L129
L143:
	;
	goto L142
L144:
	;
	if base.Ui32(v1214) < base.Ui32(v1059) {
		goto L100
	} else {
		goto L159
	}
L145:
	;
	v1218 = v1214
	goto L147
L146:
	;
	v1218 = v1216
	goto L147
L147:
	;
	if v1218 == int32(0) {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v1221 = v322 - v1218
	v1223 = v1218 & int32(3)
	v1224 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1218) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v1238 = v1224
	v1240 = int32(0)
	goto L152
L150:
	;
	v1298 = v1224
	goto L151
L151:
	;
	v1320 = v1224
	v1321 = v1298
	goto L156
L152:
	;
	v1254 = v1238 + v814
	v1255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1254))))
	v1256 = v1221 + v1238
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1256))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1254))) = uint8(v1257)
	*(*uint8)(unsafe.Add(mBase, uint32(v1256))) = uint8(v1255)
	v1261 = v1238 | int32(1)
	v1262 = v814 + v1261
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1262))))
	v1264 = v1221 + v1261
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1264))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1262))) = uint8(v1265)
	*(*uint8)(unsafe.Add(mBase, uint32(v1264))) = uint8(v1263)
	v1269 = v1238 | int32(2)
	v1270 = v814 + v1269
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270))))
	v1272 = v1221 + v1269
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1270))) = uint8(v1273)
	*(*uint8)(unsafe.Add(mBase, uint32(v1272))) = uint8(v1271)
	v1277 = v1238 | int32(3)
	v1278 = v814 + v1277
	v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1278))))
	v1280 = v1221 + v1277
	v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1280))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1278))) = uint8(v1281)
	*(*uint8)(unsafe.Add(mBase, uint32(v1280))) = uint8(v1279)
	v1284 = int32(4)
	v1285 = v1238 + v1284
	v1287 = v1240 + v1284
	if v1287 != v1218&int32(-4) {
		v1238 = v1285
		v1240 = v1287
		goto L152
	} else {
		goto L154
	}
L153:
	;
	if v1223 == int32(0) {
		goto L144
	} else {
		goto L155
	}
L154:
	;
	goto L153
L155:
	;
	v1298 = v1285
	goto L151
L156:
	;
	v1337 = v1321 + v814
	v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1337))))
	v1339 = v1221 + v1321
	v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1337))) = uint8(v1340)
	*(*uint8)(unsafe.Add(mBase, uint32(v1339))) = uint8(v1338)
	v1343 = int32(1)
	v1346 = v1320 + v1343
	if v1346 != v1223 {
		v1320 = v1346
		v1321 = v1321 + v1343
		goto L156
	} else {
		goto L158
	}
L157:
	;
	goto L144
L158:
	;
	goto L157
L159:
	;
	v1372 = base.I32_div_u_s(v1059, v34)
	F_qsort_interruptible(m, v32, v1372, v34, v35, v36)
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L8
	} else {
		goto L160
	}
L160:
	;
	v1375 = base.I32_div_u_s(v1214, v34)
	v32 = v322 - v1214
	v33 = v1375
	goto L1
L161:
	;
	v580 = v835 + v52
	v582 = v837
	v583 = v34 + v814
	v585 = v816
	goto L72
L162:
	;
	v1379 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v51) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1386 = v1379
	v1390 = v1379
	goto L166
L164:
	;
	v1451 = v1379
	goto L165
L165:
	;
	v1474 = v1451
	v1481 = v1379
	goto L170
L166:
	;
	v1408 = v1390 + v814
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1408))))
	v1410 = v1390 + v835
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1410))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1408))) = uint8(v1411)
	*(*uint8)(unsafe.Add(mBase, uint32(v1410))) = uint8(v1409)
	v1415 = v1390 | int32(1)
	v1416 = v814 + v1415
	v1417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416))))
	v1418 = v1415 + v835
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1416))) = uint8(v1419)
	*(*uint8)(unsafe.Add(mBase, uint32(v1418))) = uint8(v1417)
	v1423 = v1390 | int32(2)
	v1424 = v814 + v1423
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1424))))
	v1426 = v1423 + v835
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1424))) = uint8(v1427)
	*(*uint8)(unsafe.Add(mBase, uint32(v1426))) = uint8(v1425)
	v1431 = v1390 | int32(3)
	v1432 = v814 + v1431
	v1433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1432))))
	v1434 = v1431 + v835
	v1435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1434))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1432))) = uint8(v1435)
	*(*uint8)(unsafe.Add(mBase, uint32(v1434))) = uint8(v1433)
	v1438 = int32(4)
	v1439 = v1390 + v1438
	v1441 = v1386 + v1438
	if v1441 != v49 {
		v1386 = v1441
		v1390 = v1439
		goto L166
	} else {
		goto L168
	}
L167:
	;
	if v50 == int32(0) {
		goto L161
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v1451 = v1439
	goto L165
L170:
	;
	v1491 = v1474 + v814
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491))))
	v1493 = v1474 + v835
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1493))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1491))) = uint8(v1494)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493))) = uint8(v1492)
	v1497 = int32(1)
	v1500 = v1481 + v1497
	if v1500 != v50 {
		v1474 = v1474 + v1497
		v1481 = v1500
		goto L170
	} else {
		goto L172
	}
L171:
	;
	goto L161
L172:
	;
	goto L171
L173:
	;
	v1531 = base.I32_div_u_s(v1059, v34)
	v57 = v1531
	goto L3
}
func F_qsort_ssup_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v123 int32
	_ = v123
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int64
	_ = v186
	var v187 int64
	_ = v187
	var v196 int64
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v225 int64
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v13 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L1:
	;
	return v280
L2:
	;
	v280 = l0
	goto L1
L3:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v260 = m.T0[v259].(func(*base.Module, int64, int64, int32) int32)(m, v258, v256, l3)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L32
	} else {
		goto L86
	}
L4:
	;
	if v227&int32(1) == int32(0) {
		v256 = v225
		goto L3
	} else {
		goto L84
	}
L5:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v243 != 0 {
		goto L2
	} else {
		goto L83
	}
L6:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v228 == int32(0) {
		goto L4
	} else {
		goto L81
	}
L7:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v207 = m.T0[v206].(func(*base.Module, int64, int64, int32) int32)(m, v205, v204, l3)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L32
	} else {
		goto L75
	}
L8:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v198 != 0 {
		v280 = l1
		goto L1
	} else {
		goto L74
	}
L9:
	;
	if v183&int32(1) == int32(0) {
		v200 = v182
		v202 = v184
		v204 = v186
		v205 = v187
		goto L7
	} else {
		goto L73
	}
L10:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v154 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L11:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v144 == int32(0) {
		v280 = l1
		goto L1
	} else {
		goto L56
	}
L12:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v132 == int32(0) {
		v280 = l1
		goto L1
	} else {
		goto L54
	}
L13:
	;
	v123 = int32(1)
	if v118&v123 != 0 {
		v225 = v121
		v227 = v123
		goto L6
	} else {
		goto L53
	}
L14:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	if v53&int32(1) != 0 {
		v118 = v111
		v121 = v112
		goto L13
	} else {
		goto L52
	}
L15:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	if v107 == int32(0) {
		v130 = v106
		goto L12
	} else {
		goto L51
	}
L16:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v94 = m.T0[v93].(func(*base.Module, int64, int64, int32) int32)(m, v92, v91, l3)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L32
	} else {
		goto L45
	}
L17:
	;
	if v22&int32(1) != 0 {
		v138 = v19
		v140 = v21
		goto L11
	} else {
		goto L44
	}
L18:
	;
	if v11&int32(1) != 0 {
		goto L15
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v11&int32(1) != 0 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v19 = l2 + int32(16)
	v21 = l2 + int32(8)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v24 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	if v22&int32(1) != 0 {
		v196 = v23
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v200 = v19
	v202 = v21
	v204 = v23
	v205 = v12
	goto L7
L24:
	;
	if v59&int32(1) != 0 {
		v138 = v58
		v140 = v52
		goto L11
	} else {
		goto L43
	}
L25:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v76 != 0 {
		v118 = v74
		v121 = v75
		goto L13
	} else {
		goto L42
	}
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v71 == int32(0) {
		v147 = v66
		v149 = v68
		goto L10
	} else {
		goto L41
	}
L27:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v29 != 0 {
		goto L15
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v39 = m.T0[v38].(func(*base.Module, int64, int64, int32) int32)(m, v37, v12, l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v31 = l2 + int32(16)
	v33 = l2 + int32(8)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	if v34 == int32(0) {
		v66 = v31
		v68 = v33
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v147 = v31
	v149 = v33
	goto L10
L32:
	;
	return int32(0)
L33:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v43 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v39 < int32(0) {
		goto L25
	} else {
		goto L37
	}
L35:
	;
	v50 = v39
	goto L36
L36:
	;
	v52 = l2 + int32(8)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(0) <= v50 {
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v50 = int32(0) - v39
	goto L36
L38:
	;
	v58 = l2 + int32(16)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	if v53&int32(1) == int32(0) {
		goto L24
	} else {
		goto L39
	}
L39:
	;
	if v59&int32(1) != 0 {
		v147 = v58
		v149 = v52
		goto L10
	} else {
		goto L40
	}
L40:
	;
	v66 = v58
	v68 = v52
	goto L26
L41:
	;
	v280 = l1
	goto L1
L42:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v182 = l2 + int32(16)
	v183 = v74
	v184 = l2 + int32(8)
	v186 = v75
	v187 = v81
	goto L9
L43:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	v87 = v58
	v89 = v52
	v91 = v84
	v92 = v54
	goto L16
L44:
	;
	v87 = v19
	v89 = v21
	v91 = v23
	v92 = v12
	goto L16
L45:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v96 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	if v94 < int32(0) {
		v147 = v87
		v149 = v89
		goto L10
	} else {
		goto L49
	}
L47:
	;
	v103 = v94
	goto L48
L48:
	;
	if v103 < int32(0) {
		v280 = l1
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v103 = int32(0) - v94
	goto L48
L50:
	;
	v147 = v87
	v149 = v89
	goto L10
L51:
	;
	v225 = v106
	v227 = int32(1)
	goto L6
L52:
	;
	v182 = l2 + int32(16)
	v183 = v111
	v184 = v52
	v186 = v112
	v187 = v54
	goto L9
L53:
	;
	v130 = v121
	goto L12
L54:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v135 == int32(0) {
		v256 = v130
		goto L3
	} else {
		goto L55
	}
L55:
	;
	goto L5
L56:
	;
	v147 = v138
	v149 = v140
	goto L10
L57:
	;
	return l2
L58:
	;
	if v153&int32(1) != 0 {
		goto L2
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v153&int32(1) != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v159 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v280 = l0
	goto L1
L63:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v162 == int32(0) {
		goto L57
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v165 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v149)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v168 = m.T0[v167].(func(*base.Module, int64, int64, int32) int32)(m, v165, v166, l3)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L32
	} else {
		goto L67
	}
L66:
	;
	v280 = l0
	goto L1
L67:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v170 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v168 < int32(0) {
		goto L2
	} else {
		goto L71
	}
L69:
	;
	v177 = v168
	goto L70
L70:
	;
	if int32(0) <= v177 {
		v280 = l0
		goto L1
	} else {
		goto L72
	}
L71:
	;
	v177 = int32(0) - v168
	goto L70
L72:
	;
	goto L57
L73:
	;
	v196 = v186
	goto L8
L74:
	;
	v225 = v196
	v227 = int32(1)
	goto L6
L75:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v209 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v207 < int32(0) {
		v280 = l1
		goto L1
	} else {
		goto L79
	}
L77:
	;
	v216 = v207
	goto L78
L78:
	;
	if int32(0) < v216 {
		v280 = l1
		goto L1
	} else {
		goto L80
	}
L79:
	;
	v216 = int32(0) - v207
	goto L78
L80:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v202)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	v225 = v219
	v227 = v220
	goto L6
L81:
	;
	if v227&int32(1) == int32(0) {
		goto L5
	} else {
		goto L82
	}
L82:
	;
	return l2
L83:
	;
	v280 = l2
	goto L1
L84:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)))
	if v248 == int32(0) {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v280 = l2
	goto L1
L86:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)))
	if v262 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	if v260 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L88:
	;
	v270 = v260
	goto L89
L89:
	;
	if int32(0) <= v270 {
		v280 = l2
		goto L1
	} else {
		goto L93
	}
L90:
	;
	return l2
L91:
	;
	goto L92
L92:
	;
	v270 = int32(0) - v260
	goto L89
L93:
	;
	goto L2
}
