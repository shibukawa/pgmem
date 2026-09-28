package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswBuildAppendPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v6 = F_HnswNewBuffer(m, l0, l3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if v6 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(v6^int32(-1))*int32(56))+16))
			v26 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[1]))
			v20 = int32(56)
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v6*v20-v20)+16))
			v26 = v25
		}
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
		*(*int32)(unsafe.Add(mBase, uint32(v27+v28))) = v26
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_MarkBufferDirty(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			F_UnlockReleaseBuffer(m, v34)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_UnlockBuffer(m, v6)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[2]))
					if v40 != 0 {
						F_ProcessInterrupts(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							F_LockBufferInternal(m, v6, int32(3))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
								if v6 < int32(0) {
									v50 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v6^int32(-1))<<(uint(int32(2))%32))))
									v64 = v56
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
									v64 = v58 + v6<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v64
								F_PageInit(m, v64, int32(_a_F_HnswBuildAppendPage_0), int32(8))
								mBase = m.M
								v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+16)))
								v70 = v64 + v69
								v71 = int32(_a_F_HnswBuildAppendPage_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v71)
								*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(-1)
								return
							}
						}
					} else {
						F_LockBufferInternal(m, v6, int32(3))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
							if v6 < int32(0) {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v6^int32(-1))<<(uint(int32(2))%32))))
								v64 = v56
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
								v64 = v58 + v6<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v64
							F_PageInit(m, v64, int32(_a_F_HnswBuildAppendPage_0), int32(8))
							mBase = m.M
							v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+16)))
							v70 = v64 + v69
							v71 = int32(_a_F_HnswBuildAppendPage_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v71)
							*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(-1)
							return
						}
					}
				}
			}
		}
	}
}
func F_HnswInitElementFromBlock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = l1
	v5 = F_palloc(m, int32(108))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+80)) = uint16(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+76)) = l0
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+88)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v5)+72)) = v11
		return v5
	}
}
func F_HnswInsertTupleOnDisk(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v702 int32
	_ = v702
	var v713 int32
	_ = v713
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
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
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v804 int32
	_ = v804
	var v815 int32
	_ = v815
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
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
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v967 int32
	_ = v967
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
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
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1168 int32
	_ = v1168
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1205 int32
	_ = v1205
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1249 int32
	_ = v1249
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1408 int32
	_ = v1408
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	v3 = l2
	v6 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(80)
	m.G0 = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v38 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_LockPage(m, l0, int32(0), int32(5))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v43 = int32(64)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v43 = v42
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	F_HnswGetMetaPageInfo(m, l0, v35+int32(56), v35+int32(60))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v60 = F_log(m, base.F64_convert_i32_s(v57))
	mBase = m.M
	v62 = int32(63)
	v64 = base.I32_div_u_s(int32(1358), v57)
	v66 = v64 - int32(2)
	if base.Ui32(v62) <= base.Ui32(v66) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v69 = v62
	goto L10
L9:
	;
	v69 = v66
	goto L10
L10:
	;
	v71 = F_HnswInitElement(m, int32(0), l3, v57, base.F64_div(float64(1), v60), v69, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v71)+88)) = uint32(v3)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	if v74 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	F_HnswFindElementNeighbors(m, v92, v71, v90, l0, l1, v93, v43, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L20
	}
L13:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+65)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+65)))
	if base.Ui32(v75) <= base.Ui32(v76) {
		v90 = v74
		v91 = int32(5)
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_UnlockPage(m, l0, int32(0), int32(5))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v82 = int32(7)
	F_LockPage(m, l0, int32(0), v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v87 = F_HnswGetEntryPoint(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+60)) = v87
	v90 = v87
	v91 = v82
	goto L12
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v35)+60))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L5
	} else {
		goto L372
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L5
	} else {
		goto L369
	}
L23:
	;
	F_UnlockPage(m, l0, int32(0), v91)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L5
	} else {
		goto L368
	}
L24:
	;
	v289 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L84
	}
L25:
	;
	v104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v71)+88)))
	v106 = v71 + int32(4)
	v117 = v6
	goto L26
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v100+int32(8)+v117*int32(12))))
	v145 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v144)+88)))
	v148 = F_datumIsEqual(m, v104, v145, int32(0), int32(-1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	if v148 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)+76))
	v153 = F_ReadBuffer(m, l0, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_LockBufferInternal(m, v153, int32(3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	if l4 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+80)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v184<<(uint(int32(2))%32))+20))
	v191 = v183 + v188&int32(_a_F_HnswInsertTupleOnDisk_0)
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+8)))
	if v192 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	if v153 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v177 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L39
	}
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v162+(v153^int32(-1))<<(uint(int32(2))%32))))
	v182 = int32(0)
	v183 = v168
	goto L32
L37:
	;
	goto L38
L38:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v182 = int32(0)
	v183 = v171 + v153<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L39:
	;
	v180 = F_GenericXLogRegisterBuffer(m, v177, v153, int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v182 = v177
	v183 = v180
	goto L32
L41:
	;
	F_UnlockReleaseBuffer(m, v153)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L78
	}
L42:
	;
	v236 = v191 + v233*int32(6)
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+8)) = uint16(v237)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v236)+4)) = v239
	if l4 != 0 {
		goto L73
	} else {
		goto L74
	}
L43:
	;
	v230 = base.B2i32(v192 == int32(0))
	if l4 != 0 {
		v246 = v230
		goto L41
	} else {
		goto L70
	}
L44:
	;
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+14)))
	if v195 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v233 = int32(1)
	goto L42
L46:
	;
	goto L47
L47:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+20)))
	if v199 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v233 = int32(2)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+26)))
	if v203 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v233 = int32(3)
	goto L42
L52:
	;
	goto L53
L53:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+32)))
	if v207 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v233 = int32(4)
	goto L42
L55:
	;
	goto L56
L56:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+38)))
	if v211 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v233 = int32(5)
	goto L42
L58:
	;
	goto L59
L59:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+44)))
	if v215 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v233 = int32(6)
	goto L42
L61:
	;
	goto L62
L62:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+50)))
	if v219 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v233 = int32(7)
	goto L42
L64:
	;
	goto L65
L65:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+56)))
	if v223 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v233 = int32(8)
	goto L42
L67:
	;
	goto L68
L68:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+62)))
	if v227 != 0 {
		goto L43
	} else {
		goto L69
	}
L69:
	;
	v233 = int32(9)
	goto L42
L70:
	;
	F_pfree(m, v182)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v246 = v230
	goto L41
L72:
	;
	v246 = int32(1)
	goto L41
L73:
	;
	F_MarkBufferDirty(m, v153)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_GenericXLogFinish(m, v182)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	goto L72
L78:
	;
	if v246 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v251 = v192
	goto L81
L80:
	;
	v251 = int32(0)
	goto L81
L81:
	;
	if v251 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	v253 = v117 + int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v253 < v254 {
		v117 = v253
		goto L26
	} else {
		goto L83
	}
L83:
	;
	goto L27
L84:
	;
	F_LockBufferInternal(m, v289, int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	if v289 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+48))
	F_UnlockReleaseBuffer(m, v289)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L90
	}
L87:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v297+(v289^int32(-1))<<(uint(int32(2))%32))))
	v311 = v303
	goto L86
L88:
	;
	goto L89
L89:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v311 = v305 + v289<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L90:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	if v317 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v343 = F_add_size(m, int32(72), v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L102
	}
L92:
	;
	v321 = int32(18)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	if v323 == v321 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v334 = int32(1)
	if v317&v334 != 0 {
		v342 = int32(base.Ui32(v317) >> (uint(v334) % 32))
		goto L91
	} else {
		goto L101
	}
L95:
	;
	v326 = v321
	goto L97
L96:
	;
	v326 = int32(2)
	goto L97
L97:
	;
	if base.Ui32((v323-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v333 = int32(6)
	goto L100
L99:
	;
	v333 = v326
	goto L100
L100:
	;
	v342 = v333
	goto L91
L101:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v342 = int32(base.Ui32(v338) >> (uint(int32(2)) % 32))
	goto L91
L102:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+65)))
	v349 = F_add_size(m, v347, int32(2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v351 = F_mul_size(m, v349, v98)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v353 = F_mul_size(m, int32(6), v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v355 = F_add_size(m, int32(4), v353)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v361 = F_add_size(m, int32(0), int32(2))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v363 = F_mul_size(m, v361, v98)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v365 = F_mul_size(m, int32(6), v363)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	v367 = F_add_size(m, int32(4), v365)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	v373 = (v343 + int32(7)) & int32(-8)
	v374 = F_palloc0(m, v373)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	goto L114
L112:
	;
	v552 = (v355 + int32(7)) & int32(-8)
	v553 = F_palloc0(m, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L163
	}
L113:
	;
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v374))) = uint8(v389)
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+65)))
	v392 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+2)) = uint8(v392)
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)) = uint8(v391)
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+3)) = uint8(v395)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if v397 == v392 {
		goto L137
	} else {
		goto L138
	}
L114:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v71)+88))
	goto L113
L118:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379))))
	if v519 == int32(1) {
		goto L150
	} else {
		goto L151
	}
L119:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+62)) = uint16(v514)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v71)+58))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+58)) = v516
	goto L118
L120:
	;
	v510 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+62)) = uint16(v510)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+58)) = int32(-1)
	goto L118
L121:
	;
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+56)) = uint16(v502)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+52)) = v504
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(9)) < base.Ui32(v506) {
		goto L119
	} else {
		goto L148
	}
L122:
	;
	v498 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+56)) = uint16(v498)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+52)) = int32(-1)
	goto L120
L123:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+50)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+50)) = uint16(v490)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v71)+46))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+46)) = v492
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(8)) < base.Ui32(v494) {
		goto L121
	} else {
		goto L147
	}
L124:
	;
	v486 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+50)) = uint16(v486)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+46)) = int32(-1)
	goto L122
L125:
	;
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+44)) = uint16(v478)
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v71)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+40)) = v480
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(7)) < base.Ui32(v482) {
		goto L123
	} else {
		goto L146
	}
L126:
	;
	v474 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+44)) = uint16(v474)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+40)) = int32(-1)
	goto L124
L127:
	;
	v466 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+38)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+38)) = uint16(v466)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v71)+34))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+34)) = v468
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(6)) < base.Ui32(v470) {
		goto L125
	} else {
		goto L145
	}
L128:
	;
	v462 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+38)) = uint16(v462)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+34)) = int32(-1)
	goto L126
L129:
	;
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+32)) = uint16(v454)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+28)) = v456
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(5)) < base.Ui32(v458) {
		goto L127
	} else {
		goto L144
	}
L130:
	;
	v450 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+32)) = uint16(v450)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+28)) = int32(-1)
	goto L128
L131:
	;
	v442 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+26)) = uint16(v442)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v71)+22))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+22)) = v444
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(4)) < base.Ui32(v446) {
		goto L129
	} else {
		goto L143
	}
L132:
	;
	v438 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+26)) = uint16(v438)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+22)) = int32(-1)
	goto L130
L133:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+20)) = uint16(v430)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+16)) = v432
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(3)) < base.Ui32(v434) {
		goto L131
	} else {
		goto L142
	}
L134:
	;
	v426 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+20)) = uint16(v426)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+16)) = int32(-1)
	goto L132
L135:
	;
	v418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+14)) = uint16(v418)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v71)+10))
	*(*int32)(unsafe.Add(mBase, uint32(v374)+10)) = v420
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(2)) < base.Ui32(v422) {
		goto L133
	} else {
		goto L141
	}
L136:
	;
	v414 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+14)) = uint16(v414)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+10)) = int32(-1)
	goto L134
L137:
	;
	v400 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+8)) = uint16(v400)
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = int32(-1)
	goto L136
L138:
	;
	goto L139
L139:
	;
	v405 = v374 + int32(4)
	v406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v405)+4)) = uint16(v406)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v405))) = v408
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+64)))
	if base.Ui32(int32(1)) < base.Ui32(v410) {
		goto L135
	} else {
		goto L140
	}
L140:
	;
	goto L136
L141:
	;
	goto L134
L142:
	;
	goto L132
L143:
	;
	goto L130
L144:
	;
	goto L128
L145:
	;
	goto L126
L146:
	;
	goto L124
L147:
	;
	goto L122
L148:
	;
	goto L120
L149:
	;
	if v544 != 0 {
		goto L160
	} else {
		goto L161
	}
L150:
	;
	v523 = int32(18)
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v525 == v523 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v536 = int32(1)
	if v519&v536 != 0 {
		v544 = int32(base.Ui32(v519) >> (uint(v536) % 32))
		goto L149
	} else {
		goto L159
	}
L153:
	;
	v528 = v523
	goto L155
L154:
	;
	v528 = int32(2)
	goto L155
L155:
	;
	if base.Ui32((v525-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v535 = int32(6)
	goto L158
L157:
	;
	v535 = v528
	goto L158
L158:
	;
	v544 = v535
	goto L149
L159:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v544 = int32(base.Ui32(v540) >> (uint(int32(2)) % 32))
	goto L149
L160:
	;
	base.MemoryCopy(m, v374+int32(72), v379, v544)
	goto L162
L161:
	;
	goto L162
L162:
	;
	goto L112
L163:
	;
	v555 = int32(0)
	v566 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v553))) = uint8(v566)
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+65)))
	v578 = v570
	v582 = v555
	goto L165
L164:
	;
	v686 = v373 + v552
	v687 = int32(4)
	v688 = v686 | v687
	v702 = v312
	v713 = int32(-1)
	goto L191
L165:
	;
	goto L170
L166:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v553)+2)) = uint16(v673)
	v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)) = uint8(v684)
	goto L164
L167:
	;
	if base.B2i32(v98 <= v555) == int32(0) {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L171
L171:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v71)+72))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v597+v578<<(uint(int32(2))%32))))
	goto L167
L173:
	;
	v611 = int32(0)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v601)))
	v620 = v611
	v625 = v582
	goto L176
L174:
	;
	v673 = v582
	goto L175
L175:
	;
	if int32(0) < v578 {
		v578 = v578 - int32(1)
		v582 = v673
		goto L165
	} else {
		goto L187
	}
L176:
	;
	v633 = v553 + int32(4) + v625*int32(6)
	if v620 < v614 {
		goto L179
	} else {
		goto L180
	}
L177:
	;
	v673 = v658
	goto L175
L178:
	;
	v657 = int32(1)
	v658 = v625 + v657
	*(*uint16)(unsafe.Add(mBase, uint32(v633)+4)) = uint16(v655)
	*(*uint16)(unsafe.Add(mBase, uint32(v633)+2)) = uint16(v656)
	v662 = v620 + v657
	if v662 != v98<<(uint(base.B2i32(v578 == v611))%32) {
		v620 = v662
		v625 = v658
		goto L176
	} else {
		goto L186
	}
L179:
	;
	goto L183
L180:
	;
	goto L181
L181:
	;
	v651 = int32(_a_F_HnswInsertTupleOnDisk_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v633))) = uint16(v651)
	v655 = int32(0)
	v656 = v651
	goto L178
L182:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v640)+80)))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v640)+76))
	v649 = int32(base.Ui32(v647) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v633))) = uint16(v649)
	v655 = v646
	v656 = v647
	goto L178
L183:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v601+int32(8)+v620*int32(12))))
	goto L182
L186:
	;
	goto L177
L187:
	;
	goto L166
L188:
	;
	if v1175 < int32(0) {
		goto L310
	} else {
		goto L311
	}
L189:
	;
	v1168 = int32(0)
	v1173 = v1168
	v1175 = v1141
	v1178 = v1144
	v1180 = v1146
	v1184 = v1150
	v1189 = v1168
	goto L188
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+72)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(v35)+76)) = v1107
	v1141 = v1107
	v1144 = v1110
	v1146 = v1112
	v1150 = v1116
	goto L189
L191:
	;
	v731 = F_ReadBuffer(m, l0, v702)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L5
	} else {
		goto L193
	}
L192:
	;
	F_HnswInsertAppendPage(m, l0, v35+int32(68), v35-int32(-64), v760, v761, l4)
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L5
	} else {
		goto L289
	}
L193:
	;
	F_LockBufferInternal(m, v731, int32(3))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	if l4 != 0 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v713 == int32(-1) {
		goto L204
	} else {
		goto L205
	}
L196:
	;
	if v731 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v755 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L5
	} else {
		goto L202
	}
L199:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v740+(v731^int32(-1))<<(uint(int32(2))%32))))
	v760 = int32(0)
	v761 = v746
	goto L195
L200:
	;
	goto L201
L201:
	;
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v760 = int32(0)
	v761 = v749 + v731<<(uint(int32(13))%32) + int32(-8192)
	goto L195
L202:
	;
	v758 = F_GenericXLogRegisterBuffer(m, v755, v731, int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	v760 = v755
	v761 = v758
	goto L195
L204:
	;
	v765 = int32(4)
	v766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+14)))
	v767 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+12)))
	v768 = v766 - v767
	if v768 <= v765 {
		goto L208
	} else {
		goto L209
	}
L205:
	;
	v776 = v713
	goto L206
L206:
	;
	v777 = int32(4)
	v778 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+14)))
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+12)))
	v780 = v778 - v779
	if v780 <= v777 {
		goto L215
	} else {
		goto L216
	}
L207:
	;
	if base.Ui32(v771-int32(4)) < base.Ui32((v367+int32(7))&int32(-8)+v373|v687) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v771 = v765
	goto L210
L209:
	;
	v771 = v768
	goto L210
L210:
	;
	goto L207
L211:
	;
	v775 = int32(-1)
	goto L213
L212:
	;
	v775 = v702
	goto L213
L213:
	;
	v776 = v775
	goto L206
L214:
	;
	if base.Ui32(v688) <= base.Ui32(v783-int32(4)) {
		v1107 = v731
		v1110 = v761
		v1112 = v760
		v1116 = v776
		goto L190
	} else {
		goto L218
	}
L215:
	;
	v783 = v777
	goto L217
L216:
	;
	v783 = v780
	goto L217
L217:
	;
	goto L214
L218:
	;
	v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+12)))
	if base.Ui32(v787) < base.Ui32(int32(25)) {
		v967 = v776
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v1032 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+16)))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v761+v1032)))
	if v1034 != int32(-1) {
		goto L281
	} else {
		goto L282
	}
L220:
	;
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+3)))
	if v731 != v898 {
		goto L270
	} else {
		goto L271
	}
L221:
	;
	if base.Ui32(v686) < base.Ui32(int32(_a_F_HnswInsertTupleOnDisk_2)) {
		goto L219
	} else {
		goto L262
	}
L222:
	;
	v795 = int32(base.Ui32(v787+int32(_a_F_HnswInsertTupleOnDisk_3))>>(uint(int32(2))%32)) & int32(_a_F_HnswInsertTupleOnDisk_1)
	if v795 == int32(0) {
		v967 = v776
		goto L221
	} else {
		goto L223
	}
L223:
	;
	v804 = int32(1)
	v815 = v776
	goto L224
L224:
	;
	v837 = v761 + int32(20) + v804&int32(_a_F_HnswInsertTupleOnDisk_1)<<(uint(int32(2))%32)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v841 = v761 + v838&int32(_a_F_HnswInsertTupleOnDisk_0)
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841))))
	if v842 != int32(1) {
		v941 = v815
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v967 = v941
	goto L221
L226:
	;
	v949 = v804 + int32(1)
	if base.Ui32(v949&int32(_a_F_HnswInsertTupleOnDisk_1)) <= base.Ui32(v795) {
		v804 = v949
		v815 = v941
		goto L224
	} else {
		goto L261
	}
L227:
	;
	v845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+2)))
	if v845 == int32(0) {
		v941 = v815
		goto L226
	} else {
		goto L228
	}
L228:
	;
	if v731 < int32(0) {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	v867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v841)+68)))
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v841)+66)))
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v841)+64)))
	v872 = v868 | v869<<(uint(int32(16))%32)
	if v872 == v866 {
		goto L234
	} else {
		goto L235
	}
L230:
	;
	v851 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v851+(v731^int32(-1))*int32(56))+16))
	v866 = v857
	goto L229
L231:
	;
	goto L232
L232:
	;
	v859 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v860 = int32(56)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v859+v731*v860-v860)+16))
	v866 = v865
	goto L229
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+72)) = v899
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v837)))
	v904 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+14)))
	v905 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+12)))
	v906 = v904 - v905
	v907 = int32(0)
	if v907 < v906 {
		goto L243
	} else {
		goto L244
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+76)) = v731
	v898 = v731
	v899 = v761
	goto L233
L235:
	;
	goto L236
L236:
	;
	v875 = F_ReadBuffer(m, l0, v872)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L5
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+76)) = v875
	F_LockBufferInternal(m, v875, int32(3))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L5
	} else {
		goto L238
	}
L238:
	;
	if v875 < int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v884 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v884+(v875^int32(-1))<<(uint(int32(2))%32))))
	v898 = v875
	v899 = v890
	goto L233
L240:
	;
	goto L241
L241:
	;
	v892 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v898 = v875
	v899 = v892 + v875<<(uint(int32(13))%32) + int32(-8192)
	goto L233
L242:
	;
	v911 = int32(base.Ui32(v901)>>(uint(int32(17))%32)) + v910
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v899+v867<<(uint(int32(2))%32))+20))
	v917 = int32(base.Ui32(v915) >> (uint(int32(17)) % 32))
	if v866 != v872 {
		goto L247
	} else {
		goto L248
	}
L243:
	;
	v910 = v906
	goto L245
L244:
	;
	v910 = v907
	goto L245
L245:
	;
	goto L242
L246:
	;
	if v815 == int32(-1) {
		goto L255
	} else {
		goto L256
	}
L247:
	;
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v899)+14)))
	v920 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v899)+12)))
	v921 = v919 - v920
	v922 = int32(0)
	if v922 < v921 {
		goto L251
	} else {
		goto L252
	}
L248:
	;
	goto L249
L249:
	;
	if base.Ui32(v911) < base.Ui32(v373) {
		v930 = v917
		goto L246
	} else {
		goto L254
	}
L250:
	;
	v930 = v925 + v917
	goto L246
L251:
	;
	v925 = v921
	goto L253
L252:
	;
	v925 = v922
	goto L253
L253:
	;
	goto L250
L254:
	;
	v930 = v911 - v373 + v917
	goto L246
L255:
	;
	v933 = v866
	goto L257
L256:
	;
	v933 = v815
	goto L257
L257:
	;
	if base.B2i32(base.Ui32(v373) <= base.Ui32(v911))&base.B2i32(base.Ui32(v552) <= base.Ui32(v930)) != 0 {
		goto L220
	} else {
		goto L258
	}
L258:
	;
	if v731 == v898 {
		v941 = v933
		goto L226
	} else {
		goto L259
	}
L259:
	;
	F_UnlockReleaseBuffer(m, v898)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	v941 = v933
	goto L226
L261:
	;
	goto L225
L262:
	;
	v985 = int32(4)
	v986 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+14)))
	v987 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+12)))
	v988 = v986 - v987
	if v988 <= v985 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	if base.Ui32(v991-int32(4)) < base.Ui32(v373) {
		goto L219
	} else {
		goto L267
	}
L264:
	;
	v991 = v985
	goto L266
L265:
	;
	v991 = v988
	goto L266
L266:
	;
	goto L263
L267:
	;
	v995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v761)+16)))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v761+v995)))
	if v997 != int32(-1) {
		goto L219
	} else {
		goto L268
	}
L268:
	;
	F_HnswInsertAppendPage(m, l0, v35+int32(76), v35+int32(72), v760, v761, l4)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L5
	} else {
		goto L269
	}
L269:
	;
	v1141 = v731
	v1144 = v761
	v1146 = v760
	v1150 = v967
	goto L189
L270:
	;
	if l4 != 0 {
		goto L274
	} else {
		goto L275
	}
L271:
	;
	goto L272
L272:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v374)+3)) = uint8(v1006)
	*(*uint8)(unsafe.Add(mBase, uint32(v553)+1)) = uint8(v1006)
	v1173 = v804
	v1175 = v731
	v1178 = v761
	v1180 = v760
	v1184 = v933
	v1189 = v867
	goto L188
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+72)) = v1028
	goto L272
L274:
	;
	if v898 < int32(0) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	goto L276
L276:
	;
	v1026 = F_GenericXLogRegisterBuffer(m, v760, v898, int32(0))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L5
	} else {
		goto L280
	}
L277:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1011+(v898^int32(-1))<<(uint(int32(2))%32))))
	v1028 = v1017
	goto L273
L278:
	;
	goto L279
L279:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v1028 = v1019 + v898<<(uint(int32(13))%32) + int32(-8192)
	goto L273
L280:
	;
	v1028 = v1026
	goto L273
L281:
	;
	if l4 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	goto L192
L284:
	;
	F_pfree(m, v760)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L5
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	F_UnlockReleaseBuffer(m, v731)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L5
	} else {
		goto L288
	}
L287:
	;
	goto L286
L288:
	;
	v702 = v1034
	v713 = v967
	goto L191
L289:
	;
	if l4 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1086 = int32(4)
	v1087 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1085)+14)))
	v1088 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1085)+12)))
	v1089 = v1087 - v1088
	if v1089 <= v1086 {
		goto L304
	} else {
		goto L305
	}
L291:
	;
	F_MarkBufferDirty(m, v731)
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L5
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	F_GenericXLogFinish(m, v760)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L5
	} else {
		goto L299
	}
L294:
	;
	F_UnlockReleaseBuffer(m, v731)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L5
	} else {
		goto L295
	}
L295:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	if v1053 < int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1058+(v1053^int32(-1))<<(uint(int32(2))%32))))
	v1083 = v1053
	v1084 = int32(0)
	v1085 = v1064
	goto L290
L297:
	;
	goto L298
L298:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v1083 = v1053
	v1084 = int32(0)
	v1085 = v1067 + v1053<<(uint(int32(13))%32) + int32(-8192)
	goto L290
L299:
	;
	F_UnlockReleaseBuffer(m, v731)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v35)+68))
	v1078 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L5
	} else {
		goto L301
	}
L301:
	;
	v1081 = F_GenericXLogRegisterBuffer(m, v1078, v1077, int32(0))
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L5
	} else {
		goto L302
	}
L302:
	;
	v1083 = v1077
	v1084 = v1078
	v1085 = v1081
	goto L290
L303:
	;
	if base.Ui32(v688) <= base.Ui32(v1092-int32(4)) {
		v1107 = v1083
		v1110 = v1085
		v1112 = v1084
		v1116 = v967
		goto L190
	} else {
		goto L307
	}
L304:
	;
	v1092 = v1086
	goto L306
L305:
	;
	v1092 = v1089
	goto L306
L306:
	;
	goto L303
L307:
	;
	F_HnswInsertAppendPage(m, l0, v35+int32(76), v35+int32(72), v1084, v1085, l4)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	v1141 = v1083
	v1144 = v1085
	v1146 = v1084
	v1150 = v967
	goto L189
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v1220
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v35)+76))
	if v1222 < int32(0) {
		goto L314
	} else {
		goto L315
	}
L310:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1205+(v1175^int32(-1))*int32(56))+16))
	v1220 = v1211
	goto L309
L311:
	;
	goto L312
L312:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1214 = int32(56)
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1213+v1175*v1214-v1214)+16))
	v1220 = v1219
	goto L309
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+84)) = v1241
	if base.Ui32(int32(2048)) <= base.Ui32((v1173-int32(1))&int32(_a_F_HnswInsertTupleOnDisk_1)) {
		goto L319
	} else {
		goto L320
	}
L314:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1226+(v1222^int32(-1))*int32(56))+16))
	v1241 = v1232
	goto L313
L315:
	;
	goto L316
L316:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1235 = int32(56)
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1234+v1222*v1235-v1235)+16))
	v1241 = v1240
	goto L313
L317:
	;
	if l4 != 0 {
		goto L343
	} else {
		goto L344
	}
L318:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	v1330 = int32(0)
	v1332 = F_PageAddItemExtended(m, v1329, v553, v552, v1330, v1330)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L5
	} else {
		goto L340
	}
L319:
	;
	v1249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1178)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1249) {
		goto L322
	} else {
		goto L323
	}
L320:
	;
	goto L321
L321:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+82)) = uint16(v1189)
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v1173)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+68)) = uint16(v1189)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+66)) = uint16(v1241)
	v1301 = int32(base.Ui32(v1241) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+64)) = uint16(v1301)
	v1305 = F_PageIndexTupleOverwrite(m, v1178, v1173&int32(_a_F_HnswInsertTupleOnDisk_1), v374, v373)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L5
	} else {
		goto L333
	}
L322:
	;
	v1257 = int32(base.Ui32(v1249+int32(_a_F_HnswInsertTupleOnDisk_3)) >> (uint(int32(2)) % 32))
	goto L324
L323:
	;
	v1257 = int32(0)
	goto L324
L324:
	;
	v1258 = int32(1)
	v1259 = v1257 + v1258
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v1259)
	if v1175 != v1222 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1265 = v1258
	goto L327
L326:
	;
	v1265 = v1257 + int32(2)
	goto L327
L327:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+82)) = uint16(v1265)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+68)) = uint16(v1265)
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+66)) = uint16(v1241)
	v1270 = int32(base.Ui32(v1241) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v374)+64)) = uint16(v1270)
	v1272 = int32(0)
	v1274 = F_PageAddItemExtended(m, v1178, v374, v373, v1272, v1272)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	v1276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)))
	if v1274 == v1276 {
		goto L318
	} else {
		goto L329
	}
L329:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+48)) = v1282 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v35+int32(48))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L5
	} else {
		goto L331
	}
L331:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(325), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L333:
	;
	if v1305 == int32(0) {
		goto L22
	} else {
		goto L334
	}
L334:
	;
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v35)+72))
	v1310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+82)))
	v1311 = F_PageIndexTupleOverwrite(m, v1309, v1310, v553, v552)
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	if v1311 != 0 {
		goto L317
	} else {
		goto L336
	}
L336:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v1317 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v35)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L5
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(320), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L5
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	v1334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+82)))
	if v1332 != v1334 {
		goto L21
	} else {
		goto L341
	}
L341:
	;
	goto L317
L342:
	;
	if v1184 == int32(-1) {
		goto L350
	} else {
		goto L351
	}
L343:
	;
	F_MarkBufferDirty(m, v1175)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L5
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	F_GenericXLogFinish(m, v1180)
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L5
	} else {
		goto L349
	}
L346:
	;
	if v1175 == v1222 {
		goto L342
	} else {
		goto L347
	}
L347:
	;
	F_MarkBufferDirty(m, v1222)
	mBase = m.M
	v1343 = m.ExcPending
	if v1343 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	goto L342
L349:
	;
	goto L342
L350:
	;
	v1346 = v1241
	goto L352
L351:
	;
	v1346 = v1184
	goto L352
L352:
	;
	F_UnlockReleaseBuffer(m, v1175)
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	if v1175 != v1222 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	F_UnlockReleaseBuffer(m, v1222)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L5
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	if base.B2i32(v1346 == int32(-1))|base.B2i32(v1346 == v312) == int32(0) {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L356
L358:
	;
	v1358 = int32(0)
	F_HnswUpdateMetaPage(m, l0, v1358, v1358, v1346, v1358, l4)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L5
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	F_HnswUpdateNeighborsOnDisk(m, l0, l1, v71, v98, int32(0), l4)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L5
	} else {
		goto L362
	}
L361:
	;
	goto L360
L362:
	;
	if v97 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+65)))
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+65)))
	if base.Ui32(v1366) <= base.Ui32(v1367) {
		goto L23
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	F_HnswUpdateMetaPage(m, l0, int32(1), v71, int32(-1), int32(0), l4)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L5
	} else {
		goto L367
	}
L366:
	;
	goto L365
L367:
	;
	goto L23
L368:
	;
	m.G0 = v35 + int32(80)
	return int32(1)
L369:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v1418 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v35+int32(16))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L5
	} else {
		goto L370
	}
L370:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(317), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L5
	} else {
		goto L371
	}
L371:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L372:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v1436 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v35+int32(32))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L5
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(328), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L5
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_HnswMemoryContextAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	v4 = F_MemoryContextAlloc(m, v3, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+132)) = v12
		return v4
	}
}
func F_HnswNormValue(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_DirectFunctionCall1Coll(m, v4, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_HnswSharedMemoryAlloc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v8 = (l0 + int32(7)) & int32(-8)
	if base.Ui32(v8) < base.Ui32(int32(_a_F_HnswSharedMemoryAlloc_0)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+68))
		v13 = F_add_size(m, v12, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
			if base.Ui32(v18) < base.Ui32(v13) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswSharedMemoryAlloc_1), int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_HnswSharedMemoryAlloc_2), int32(673), int32(_a_F_HnswSharedMemoryAlloc_3))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+204))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v13
				return v21 + v20
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_HnswSharedMemoryAlloc_4), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_HnswSharedMemoryAlloc_2), int32(669), int32(_a_F_HnswSharedMemoryAlloc_3))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_HnswUpdateMetaPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = int32(0)
	v12 = F_ReadBufferExtended(m, l0, l4, v7, v7, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBufferInternal(m, v12, int32(3))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if l5 != 0 {
				if v12 < int32(0) {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateMetaPage[0]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v12^int32(-1))<<(uint(int32(2))%32))))
					v39 = v7
					v40 = v26
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateMetaPage[1]))
					v39 = v7
					v40 = v28 + v12<<(uint(int32(13))%32) + int32(-8192)
				}
				if l1 == int32(0) {
				} else {
					if l2 == int32(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = int64(-281470681743361)
					} else {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
						if l1 != int32(2) {
							v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+46)))
							if v47 <= v50 {
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
								v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
								*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
							}
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
							v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
							*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
						}
					}
				}
				if l3 != int32(-1) {
					*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l3
				} else {
				}
				if l5 != 0 {
					F_MarkBufferDirty(m, v12)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_UnlockReleaseBuffer(m, v12)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_GenericXLogFinish(m, v39)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_UnlockReleaseBuffer(m, v12)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v34 = F_GenericXLogStart(m, l0)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v37 = F_GenericXLogRegisterBuffer(m, v34, v12, int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v39 = v34
						v40 = v37
						if l1 == int32(0) {
						} else {
							if l2 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v40)+40)) = int64(-281470681743361)
							} else {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
								if l1 != int32(2) {
									v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(v40)+46)))
									if v47 <= v50 {
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
										*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
										*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
										*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
									}
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v40)+40)) = v52
									v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+80)))
									*(*uint16)(unsafe.Add(mBase, uint32(v40)+46)) = uint16(v47)
									*(*uint16)(unsafe.Add(mBase, uint32(v40)+44)) = uint16(v54)
								}
							}
						}
						if l3 != int32(-1) {
							*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = l3
						} else {
						}
						if l5 != 0 {
							F_MarkBufferDirty(m, v12)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							F_GenericXLogFinish(m, v39)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
