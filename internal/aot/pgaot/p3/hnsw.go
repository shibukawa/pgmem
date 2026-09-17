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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v6 = F_HnswNewBuffer(m, l0, l3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if v6 < int32(0) {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[0]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v11+(v6^int32(-1))<<(uint(int32(6))%32))+16))
			v26 = v17
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[1]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19+v6<<(uint(int32(6))%32)+int32(-64))+16))
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
				F_LockBuffer(m, v6, int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[2]))
					if v41 != 0 {
						F_ProcessInterrupts(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_LockBuffer(m, v6, int32(2))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
								if v6 < int32(0) {
									v51 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
									v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v6^int32(-1))<<(uint(int32(2))%32))))
									v65 = v57
								} else {
									v59 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
									v65 = v59 + v6<<(uint(int32(13))%32) + int32(-8192)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
								F_PageInit(m, v65, int32(_a_F_HnswBuildAppendPage_0), int32(8))
								mBase = m.M
								v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
								v71 = v65 + v70
								v72 = int32(_a_F_HnswBuildAppendPage_1)
								*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)) = uint16(v72)
								*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(-1)
								return
							}
						}
					} else {
						F_LockBuffer(m, v6, int32(2))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
							if v6 < int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[3]))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v6^int32(-1))<<(uint(int32(2))%32))))
								v65 = v57
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, _c_F_HnswBuildAppendPage[4]))
								v65 = v59 + v6<<(uint(int32(13))%32) + int32(-8192)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
							F_PageInit(m, v65, int32(_a_F_HnswBuildAppendPage_0), int32(8))
							mBase = m.M
							v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+16)))
							v71 = v65 + v70
							v72 = int32(_a_F_HnswBuildAppendPage_1)
							*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)) = uint16(v72)
							*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(-1)
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
func F_HnswInsertTupleOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
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
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
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
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v630 int32
	_ = v630
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v698 int32
	_ = v698
	var v709 int32
	_ = v709
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v810 int32
	_ = v810
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v944 int32
	_ = v944
	var v961 int32
	_ = v961
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1142 int32
	_ = v1142
	var v1160 int32
	_ = v1160
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1240 int32
	_ = v1240
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1314 int32
	_ = v1314
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1398 int32
	_ = v1398
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1416 int32
	_ = v1416
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	v32 = m.G0
	v34 = v32 - int32(80)
	m.G0 = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v37 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_LockPage(m, l0, int32(0), int32(5))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v42 = int32(64)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v42 = v41
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	F_HnswGetMetaPageInfo(m, l0, v34+int32(56), v34+int32(60))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v59 = F_log(m, base.F64_convert_i32_s(v56))
	mBase = m.M
	v61 = int32(63)
	v63 = base.I32_div_u_s(int32(1358), v56)
	v65 = v63 - int32(2)
	if base.Ui32(v61) <= base.Ui32(v65) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v68 = v61
	goto L10
L9:
	;
	v68 = v65
	goto L10
L10:
	;
	v70 = F_HnswInitElement(m, int32(0), l3, v56, base.F64_div(float64(1), v59), v68, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+88)) = l2
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	F_HnswFindElementNeighbors(m, v91, v70, v89, l0, l1, v92, v42, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L20
	}
L13:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+65)))
	if base.Ui32(v74) <= base.Ui32(v75) {
		v89 = v73
		v90 = int32(5)
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
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v81 = int32(7)
	F_LockPage(m, l0, int32(0), v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v86 = F_HnswGetEntryPoint(m, l0)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v86
	v89 = v86
	v90 = v81
	goto L12
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L5
	} else {
		goto L370
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L5
	} else {
		goto L367
	}
L23:
	;
	F_UnlockPage(m, l0, int32(0), v90)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L5
	} else {
		goto L366
	}
L24:
	;
	v287 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L84
	}
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	v105 = v70 + int32(4)
	v112 = int32(0)
	goto L26
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v99+int32(8)+v112*int32(12))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+88))
	v147 = F_datumIsEqual(m, v103, v144, int32(0), int32(-1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L5
	} else {
		goto L28
	}
L27:
	;
	goto L24
L28:
	;
	if v147 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v143)+76))
	v152 = F_ReadBuffer(m, l0, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_LockBuffer(m, v152, int32(2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+80)))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182+v183<<(uint(int32(2))%32))+20))
	v190 = v182 + v187&int32(_a_F_HnswInsertTupleOnDisk_0)
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+8)))
	if v191 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	if v152 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v176 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L39
	}
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161+(v152^int32(-1))<<(uint(int32(2))%32))))
	v181 = int32(0)
	v182 = v167
	goto L32
L37:
	;
	goto L38
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v181 = int32(0)
	v182 = v170 + v152<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L39:
	;
	v179 = F_GenericXLogRegisterBuffer(m, v176, v152, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v181 = v176
	v182 = v179
	goto L32
L41:
	;
	F_UnlockReleaseBuffer(m, v152)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L78
	}
L42:
	;
	v235 = v190 + v232*int32(6)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v235)+8)) = uint16(v236)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v238
	if l4 != 0 {
		goto L73
	} else {
		goto L74
	}
L43:
	;
	v229 = base.B2i32(v191 == int32(0))
	if l4 != 0 {
		v245 = v229
		goto L41
	} else {
		goto L70
	}
L44:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+14)))
	if v194 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v232 = int32(1)
	goto L42
L46:
	;
	goto L47
L47:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+20)))
	if v198 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v232 = int32(2)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+26)))
	if v202 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v232 = int32(3)
	goto L42
L52:
	;
	goto L53
L53:
	;
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+32)))
	if v206 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v232 = int32(4)
	goto L42
L55:
	;
	goto L56
L56:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+38)))
	if v210 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v232 = int32(5)
	goto L42
L58:
	;
	goto L59
L59:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+44)))
	if v214 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v232 = int32(6)
	goto L42
L61:
	;
	goto L62
L62:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+50)))
	if v218 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = int32(7)
	goto L42
L64:
	;
	goto L65
L65:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+56)))
	if v222 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v232 = int32(8)
	goto L42
L67:
	;
	goto L68
L68:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+62)))
	if v226 != 0 {
		goto L43
	} else {
		goto L69
	}
L69:
	;
	v232 = int32(9)
	goto L42
L70:
	;
	F_pfree(m, v181)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	v245 = v229
	goto L41
L72:
	;
	v245 = int32(1)
	goto L41
L73:
	;
	F_MarkBufferDirty(m, v152)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L5
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_GenericXLogFinish(m, v181)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	if v245 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v250 = v191
	goto L81
L80:
	;
	v250 = int32(0)
	goto L81
L81:
	;
	if v250 != 0 {
		goto L23
	} else {
		goto L82
	}
L82:
	;
	v252 = v112 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v252 < v253 {
		v112 = v252
		goto L26
	} else {
		goto L83
	}
L83:
	;
	goto L27
L84:
	;
	F_LockBuffer(m, v287, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	if v287 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+48))
	F_UnlockReleaseBuffer(m, v287)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L90
	}
L87:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295+(v287^int32(-1))<<(uint(int32(2))%32))))
	v309 = v301
	goto L86
L88:
	;
	goto L89
L89:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v309 = v303 + v287<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L90:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v315 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v340 = F_add_size(m, int32(72), v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L100
	}
L92:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if base.Ui32((v319-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v339 = int32(6)
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v331 = int32(1)
	if v315&v331 != 0 {
		v339 = int32(base.Ui32(v315) >> (uint(v331) % 32))
		goto L91
	} else {
		goto L99
	}
L95:
	;
	v326 = int32(18)
	if v319 == v326 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v330 = v326
	goto L98
L97:
	;
	v330 = int32(2)
	goto L98
L98:
	;
	v339 = v330
	goto L91
L99:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v339 = int32(base.Ui32(v335) >> (uint(int32(2)) % 32))
	goto L91
L100:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v346 = F_add_size(m, v344, int32(2))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	v348 = F_mul_size(m, v346, v97)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	v350 = F_mul_size(m, int32(6), v348)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	v352 = F_add_size(m, int32(4), v350)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	v358 = F_add_size(m, int32(0), int32(2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	v360 = F_mul_size(m, v358, v97)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v362 = F_mul_size(m, int32(6), v360)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	v364 = F_add_size(m, int32(4), v362)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v370 = (v340 + int32(7)) & int32(-8)
	v371 = F_palloc0(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	goto L112
L110:
	;
	v549 = (v352 + int32(7)) & int32(-8)
	v550 = F_palloc0(m, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L161
	}
L111:
	;
	v386 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v371))) = uint8(v386)
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v389 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+2)) = uint8(v389)
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+1)) = uint8(v388)
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+3)) = uint8(v392)
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if v394 == v389 {
		goto L135
	} else {
		goto L136
	}
L112:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v70)+88))
	goto L111
L116:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v516 == int32(1) {
		goto L148
	} else {
		goto L149
	}
L117:
	;
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+62)) = uint16(v511)
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v70)+58))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+58)) = v513
	goto L116
L118:
	;
	v507 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+62)) = uint16(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+58)) = int32(-1)
	goto L116
L119:
	;
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+56)) = uint16(v499)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+52)) = v501
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(9)) < base.Ui32(v503) {
		goto L117
	} else {
		goto L146
	}
L120:
	;
	v495 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+56)) = uint16(v495)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+52)) = int32(-1)
	goto L118
L121:
	;
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+50)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+50)) = uint16(v487)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v70)+46))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+46)) = v489
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(8)) < base.Ui32(v491) {
		goto L119
	} else {
		goto L145
	}
L122:
	;
	v483 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+50)) = uint16(v483)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+46)) = int32(-1)
	goto L120
L123:
	;
	v475 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+44)) = uint16(v475)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v70)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+40)) = v477
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(7)) < base.Ui32(v479) {
		goto L121
	} else {
		goto L144
	}
L124:
	;
	v471 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+44)) = uint16(v471)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+40)) = int32(-1)
	goto L122
L125:
	;
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+38)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+38)) = uint16(v463)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v70)+34))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+34)) = v465
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(6)) < base.Ui32(v467) {
		goto L123
	} else {
		goto L143
	}
L126:
	;
	v459 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+38)) = uint16(v459)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+34)) = int32(-1)
	goto L124
L127:
	;
	v451 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+32)) = uint16(v451)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v70)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+28)) = v453
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(5)) < base.Ui32(v455) {
		goto L125
	} else {
		goto L142
	}
L128:
	;
	v447 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+32)) = uint16(v447)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+28)) = int32(-1)
	goto L126
L129:
	;
	v439 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+26)) = uint16(v439)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v70)+22))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+22)) = v441
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(4)) < base.Ui32(v443) {
		goto L127
	} else {
		goto L141
	}
L130:
	;
	v435 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+26)) = uint16(v435)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+22)) = int32(-1)
	goto L128
L131:
	;
	v427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+20)) = uint16(v427)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = v429
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(3)) < base.Ui32(v431) {
		goto L129
	} else {
		goto L140
	}
L132:
	;
	v423 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+20)) = uint16(v423)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = int32(-1)
	goto L130
L133:
	;
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+14)) = uint16(v415)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v70)+10))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+10)) = v417
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(2)) < base.Ui32(v419) {
		goto L131
	} else {
		goto L139
	}
L134:
	;
	v411 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+14)) = uint16(v411)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+10)) = int32(-1)
	goto L132
L135:
	;
	v397 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+8)) = uint16(v397)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = int32(-1)
	goto L134
L136:
	;
	goto L137
L137:
	;
	v402 = v371 + int32(4)
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v402)+4)) = uint16(v403)
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v402))) = v405
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+64)))
	if base.Ui32(int32(1)) < base.Ui32(v407) {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	goto L134
L139:
	;
	goto L132
L140:
	;
	goto L130
L141:
	;
	goto L128
L142:
	;
	goto L126
L143:
	;
	goto L124
L144:
	;
	goto L122
L145:
	;
	goto L120
L146:
	;
	goto L118
L147:
	;
	if v541 != 0 {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	v520 = int32(18)
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+1)))
	if v522 == v520 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v533 = int32(1)
	if v516&v533 != 0 {
		v541 = int32(base.Ui32(v516) >> (uint(v533) % 32))
		goto L147
	} else {
		goto L157
	}
L151:
	;
	v525 = v520
	goto L153
L152:
	;
	v525 = int32(2)
	goto L153
L153:
	;
	if base.Ui32((v522-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v532 = int32(6)
	goto L156
L155:
	;
	v532 = v525
	goto L156
L156:
	;
	v541 = v532
	goto L147
L157:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v376)))
	v541 = int32(base.Ui32(v537) >> (uint(int32(2)) % 32))
	goto L147
L158:
	;
	base.MemoryCopy(m, v371+int32(72), v376, v541)
	goto L160
L159:
	;
	goto L160
L160:
	;
	goto L110
L161:
	;
	v552 = int32(0)
	v563 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v550))) = uint8(v563)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v575 = v567
	v579 = v552
	goto L163
L162:
	;
	v683 = v370 + v549
	v684 = int32(4)
	v685 = v683 | v684
	v698 = v310
	v709 = int32(-1)
	goto L189
L163:
	;
	goto L168
L164:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v550)+2)) = uint16(v670)
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)) = uint8(v681)
	goto L162
L165:
	;
	if base.B2i32(v97 <= v552) == int32(0) {
		goto L171
	} else {
		goto L172
	}
L168:
	;
	goto L169
L169:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v70)+72))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v594+v575<<(uint(int32(2))%32))))
	goto L165
L171:
	;
	v608 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v598)))
	v617 = v608
	v622 = v579
	goto L174
L172:
	;
	v670 = v579
	goto L173
L173:
	;
	if int32(0) < v575 {
		v575 = v575 - int32(1)
		v579 = v670
		goto L163
	} else {
		goto L185
	}
L174:
	;
	v630 = v550 + int32(4) + v622*int32(6)
	if v617 < v611 {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	v670 = v655
	goto L173
L176:
	;
	v654 = int32(1)
	v655 = v622 + v654
	*(*uint16)(unsafe.Add(mBase, uint32(v630)+4)) = uint16(v652)
	*(*uint16)(unsafe.Add(mBase, uint32(v630)+2)) = uint16(v653)
	v659 = v617 + v654
	if v659 != v97<<(uint(base.B2i32(v575 == v608))%32) {
		v617 = v659
		v622 = v655
		goto L174
	} else {
		goto L184
	}
L177:
	;
	goto L181
L178:
	;
	goto L179
L179:
	;
	v648 = int32(_a_F_HnswInsertTupleOnDisk_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v630))) = uint16(v648)
	v652 = int32(0)
	v653 = v648
	goto L176
L180:
	;
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v637)+80)))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v637)+76))
	v646 = int32(base.Ui32(v644) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v630))) = uint16(v646)
	v652 = v643
	v653 = v644
	goto L176
L181:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v598+int32(8)+v617*int32(12))))
	goto L180
L184:
	;
	goto L175
L185:
	;
	goto L164
L186:
	;
	if v1168 < int32(0) {
		goto L308
	} else {
		goto L309
	}
L187:
	;
	v1160 = int32(0)
	v1164 = v1160
	v1165 = v1132
	v1168 = v1135
	v1172 = v1139
	v1175 = v1142
	v1180 = v1160
	goto L186
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1099
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v1102
	v1132 = v1099
	v1135 = v1102
	v1139 = v1106
	v1142 = v1109
	goto L187
L189:
	;
	v727 = F_ReadBuffer(m, l0, v698)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L5
	} else {
		goto L191
	}
L190:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(68), v34-int32(-64), v756, v757, l4)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L5
	} else {
		goto L287
	}
L191:
	;
	F_LockBuffer(m, v727, int32(2))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L5
	} else {
		goto L192
	}
L192:
	;
	if l4 != 0 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v709 == int32(-1) {
		goto L202
	} else {
		goto L203
	}
L194:
	;
	if v727 < int32(0) {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v751 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L200
	}
L197:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v736+(v727^int32(-1))<<(uint(int32(2))%32))))
	v756 = int32(0)
	v757 = v742
	goto L193
L198:
	;
	goto L199
L199:
	;
	v745 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v756 = int32(0)
	v757 = v745 + v727<<(uint(int32(13))%32) + int32(-8192)
	goto L193
L200:
	;
	v754 = F_GenericXLogRegisterBuffer(m, v751, v727, int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L201
	}
L201:
	;
	v756 = v751
	v757 = v754
	goto L193
L202:
	;
	v761 = int32(4)
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+14)))
	v763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+12)))
	v764 = v762 - v763
	if v764 <= v761 {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v772 = v709
	goto L204
L204:
	;
	v773 = int32(4)
	v774 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+14)))
	v775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+12)))
	v776 = v774 - v775
	if v776 <= v773 {
		goto L213
	} else {
		goto L214
	}
L205:
	;
	if base.Ui32(v767-int32(4)) < base.Ui32((v364+int32(7))&int32(-8)+v370|v684) {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	v767 = v761
	goto L208
L207:
	;
	v767 = v764
	goto L208
L208:
	;
	goto L205
L209:
	;
	v771 = int32(-1)
	goto L211
L210:
	;
	v771 = v698
	goto L211
L211:
	;
	v772 = v771
	goto L204
L212:
	;
	if base.Ui32(v685) <= base.Ui32(v779-int32(4)) {
		v1099 = v757
		v1102 = v727
		v1106 = v756
		v1109 = v772
		goto L188
	} else {
		goto L216
	}
L213:
	;
	v779 = v773
	goto L215
L214:
	;
	v779 = v776
	goto L215
L215:
	;
	goto L212
L216:
	;
	v783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+12)))
	if base.Ui32(v783) < base.Ui32(int32(25)) {
		v961 = v772
		goto L219
	} else {
		goto L220
	}
L217:
	;
	v1026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+16)))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v757+v1026)))
	if v1028 != int32(-1) {
		goto L279
	} else {
		goto L280
	}
L218:
	;
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+3)))
	if v727 != v893 {
		goto L268
	} else {
		goto L269
	}
L219:
	;
	if base.Ui32(v683) < base.Ui32(int32(_a_F_HnswInsertTupleOnDisk_2)) {
		goto L217
	} else {
		goto L260
	}
L220:
	;
	v791 = int32(base.Ui32(v783+int32(_a_F_HnswInsertTupleOnDisk_3))>>(uint(int32(2))%32)) & int32(_a_F_HnswInsertTupleOnDisk_1)
	if v791 == int32(0) {
		v961 = v772
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v799 = int32(1)
	v810 = v772
	goto L222
L222:
	;
	v832 = v757 + int32(20) + v799&int32(_a_F_HnswInsertTupleOnDisk_1)<<(uint(int32(2))%32)
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v836 = v757 + v833&int32(_a_F_HnswInsertTupleOnDisk_0)
	v837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836))))
	if v837 != int32(1) {
		v936 = v810
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v961 = v936
	goto L219
L224:
	;
	v944 = v799 + int32(1)
	if base.Ui32(v944&int32(_a_F_HnswInsertTupleOnDisk_1)) <= base.Ui32(v791) {
		v799 = v944
		v810 = v936
		goto L222
	} else {
		goto L259
	}
L225:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+2)))
	if v840 == int32(0) {
		v936 = v810
		goto L224
	} else {
		goto L226
	}
L226:
	;
	if v727 < int32(0) {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+68)))
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+66)))
	v864 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v836)+64)))
	v867 = v863 | v864<<(uint(int32(16))%32)
	if v867 == v861 {
		goto L232
	} else {
		goto L233
	}
L228:
	;
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v846+(v727^int32(-1))<<(uint(int32(6))%32))+16))
	v861 = v852
	goto L227
L229:
	;
	goto L230
L230:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v854+v727<<(uint(int32(6))%32)+int32(-64))+16))
	v861 = v860
	goto L227
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v894
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v832)))
	v899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+14)))
	v900 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+12)))
	v901 = v899 - v900
	v902 = int32(0)
	if v902 < v901 {
		goto L241
	} else {
		goto L242
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v727
	v893 = v727
	v894 = v757
	goto L231
L233:
	;
	goto L234
L234:
	;
	v870 = F_ReadBuffer(m, l0, v867)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+76)) = v870
	F_LockBuffer(m, v870, int32(2))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	if v870 < int32(0) {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v879+(v870^int32(-1))<<(uint(int32(2))%32))))
	v893 = v870
	v894 = v885
	goto L231
L238:
	;
	goto L239
L239:
	;
	v887 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v893 = v870
	v894 = v887 + v870<<(uint(int32(13))%32) + int32(-8192)
	goto L231
L240:
	;
	v906 = int32(base.Ui32(v896)>>(uint(int32(17))%32)) + v905
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v894+v862<<(uint(int32(2))%32))+20))
	v912 = int32(base.Ui32(v910) >> (uint(int32(17)) % 32))
	if v861 != v867 {
		goto L245
	} else {
		goto L246
	}
L241:
	;
	v905 = v901
	goto L243
L242:
	;
	v905 = v902
	goto L243
L243:
	;
	goto L240
L244:
	;
	if v810 == int32(-1) {
		goto L253
	} else {
		goto L254
	}
L245:
	;
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v894)+14)))
	v915 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v894)+12)))
	v916 = v914 - v915
	v917 = int32(0)
	if v917 < v916 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	goto L247
L247:
	;
	if base.Ui32(v906) < base.Ui32(v370) {
		v925 = v912
		goto L244
	} else {
		goto L252
	}
L248:
	;
	v925 = v920 + v912
	goto L244
L249:
	;
	v920 = v916
	goto L251
L250:
	;
	v920 = v917
	goto L251
L251:
	;
	goto L248
L252:
	;
	v925 = v906 - v370 + v912
	goto L244
L253:
	;
	v928 = v861
	goto L255
L254:
	;
	v928 = v810
	goto L255
L255:
	;
	if base.B2i32(base.Ui32(v370) <= base.Ui32(v906))&base.B2i32(base.Ui32(v549) <= base.Ui32(v925)) != 0 {
		goto L218
	} else {
		goto L256
	}
L256:
	;
	if v727 == v893 {
		v936 = v928
		goto L224
	} else {
		goto L257
	}
L257:
	;
	F_UnlockReleaseBuffer(m, v893)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L5
	} else {
		goto L258
	}
L258:
	;
	v936 = v928
	goto L224
L259:
	;
	goto L223
L260:
	;
	v979 = int32(4)
	v980 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+14)))
	v981 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+12)))
	v982 = v980 - v981
	if v982 <= v979 {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	if base.Ui32(v985-int32(4)) < base.Ui32(v370) {
		goto L217
	} else {
		goto L265
	}
L262:
	;
	v985 = v979
	goto L264
L263:
	;
	v985 = v982
	goto L264
L264:
	;
	goto L261
L265:
	;
	v989 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v757)+16)))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v757+v989)))
	if v991 != int32(-1) {
		goto L217
	} else {
		goto L266
	}
L266:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(76), v34+int32(72), v756, v757, l4)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	v1132 = v757
	v1135 = v727
	v1139 = v756
	v1142 = v961
	goto L187
L268:
	;
	if l4 != 0 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	goto L270
L270:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v371)+3)) = uint8(v1000)
	*(*uint8)(unsafe.Add(mBase, uint32(v550)+1)) = uint8(v1000)
	v1164 = v799
	v1165 = v757
	v1168 = v727
	v1172 = v756
	v1175 = v928
	v1180 = v862
	goto L186
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+72)) = v1022
	goto L270
L272:
	;
	if v893 < int32(0) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	goto L274
L274:
	;
	v1020 = F_GenericXLogRegisterBuffer(m, v756, v893, int32(0))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L5
	} else {
		goto L278
	}
L275:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v1005+(v893^int32(-1))<<(uint(int32(2))%32))))
	v1022 = v1011
	goto L271
L276:
	;
	goto L277
L277:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v1022 = v1013 + v893<<(uint(int32(13))%32) + int32(-8192)
	goto L271
L278:
	;
	v1022 = v1020
	goto L271
L279:
	;
	if l4 == int32(0) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	goto L190
L282:
	;
	F_pfree(m, v756)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L5
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	F_UnlockReleaseBuffer(m, v727)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L5
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	v698 = v1028
	v709 = v961
	goto L189
L287:
	;
	if l4 != 0 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	v1080 = int32(4)
	v1081 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1079)+14)))
	v1082 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1079)+12)))
	v1083 = v1081 - v1082
	if v1083 <= v1080 {
		goto L302
	} else {
		goto L303
	}
L289:
	;
	F_MarkBufferDirty(m, v727)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L5
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	F_GenericXLogFinish(m, v756)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L5
	} else {
		goto L297
	}
L292:
	;
	F_UnlockReleaseBuffer(m, v727)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L5
	} else {
		goto L293
	}
L293:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	if v1047 < int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[0]))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1052+(v1047^int32(-1))<<(uint(int32(2))%32))))
	v1077 = v1047
	v1078 = int32(0)
	v1079 = v1058
	goto L288
L295:
	;
	goto L296
L296:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[1]))
	v1077 = v1047
	v1078 = int32(0)
	v1079 = v1061 + v1047<<(uint(int32(13))%32) + int32(-8192)
	goto L288
L297:
	;
	F_UnlockReleaseBuffer(m, v727)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L5
	} else {
		goto L298
	}
L298:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v1072 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L5
	} else {
		goto L299
	}
L299:
	;
	v1075 = F_GenericXLogRegisterBuffer(m, v1072, v1071, int32(0))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L5
	} else {
		goto L300
	}
L300:
	;
	v1077 = v1071
	v1078 = v1072
	v1079 = v1075
	goto L288
L301:
	;
	if base.Ui32(v685) <= base.Ui32(v1086-int32(4)) {
		v1099 = v1079
		v1102 = v1077
		v1106 = v1078
		v1109 = v961
		goto L188
	} else {
		goto L305
	}
L302:
	;
	v1086 = v1080
	goto L304
L303:
	;
	v1086 = v1083
	goto L304
L304:
	;
	goto L301
L305:
	;
	F_HnswInsertAppendPage(m, l0, v34+int32(76), v34+int32(72), v1078, v1079, l4)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L5
	} else {
		goto L306
	}
L306:
	;
	v1132 = v1079
	v1135 = v1077
	v1139 = v1078
	v1142 = v961
	goto L187
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+76)) = v1211
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v1213 < int32(0) {
		goto L312
	} else {
		goto L313
	}
L308:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1196+(v1168^int32(-1))<<(uint(int32(6))%32))+16))
	v1211 = v1202
	goto L307
L309:
	;
	goto L310
L310:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1204+v1168<<(uint(int32(6))%32)+int32(-64))+16))
	v1211 = v1210
	goto L307
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+84)) = v1232
	if base.Ui32(int32(2048)) <= base.Ui32((v1164-int32(1))&int32(_a_F_HnswInsertTupleOnDisk_1)) {
		goto L317
	} else {
		goto L318
	}
L312:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[2]))
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1217+(v1213^int32(-1))<<(uint(int32(6))%32))+16))
	v1232 = v1223
	goto L311
L313:
	;
	goto L314
L314:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertTupleOnDisk[3]))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1225+v1213<<(uint(int32(6))%32)+int32(-64))+16))
	v1232 = v1231
	goto L311
L315:
	;
	if l4 != 0 {
		goto L341
	} else {
		goto L342
	}
L316:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1321 = int32(0)
	v1323 = F_PageAddItemExtended(m, v1320, v550, v549, v1321, v1321)
	mBase = m.M
	v1324 = m.ExcPending
	if v1324 != 0 {
		goto L5
	} else {
		goto L338
	}
L317:
	;
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1240) {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	goto L319
L319:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)) = uint16(v1180)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v1164)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+68)) = uint16(v1180)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+66)) = uint16(v1232)
	v1292 = int32(base.Ui32(v1232) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+64)) = uint16(v1292)
	v1296 = F_PageIndexTupleOverwrite(m, v1165, v1164&int32(_a_F_HnswInsertTupleOnDisk_1), v371, v370)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L5
	} else {
		goto L331
	}
L320:
	;
	v1248 = int32(base.Ui32(v1240+int32(_a_F_HnswInsertTupleOnDisk_3)) >> (uint(int32(2)) % 32))
	goto L322
L321:
	;
	v1248 = int32(0)
	goto L322
L322:
	;
	v1249 = int32(1)
	v1250 = v1248 + v1249
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)) = uint16(v1250)
	if v1168 != v1213 {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1256 = v1249
	goto L325
L324:
	;
	v1256 = v1248 + int32(2)
	goto L325
L325:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)) = uint16(v1256)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+68)) = uint16(v1256)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+66)) = uint16(v1232)
	v1261 = int32(base.Ui32(v1232) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+64)) = uint16(v1261)
	v1263 = int32(0)
	v1265 = F_PageAddItemExtended(m, v1165, v371, v370, v1263, v1263)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L5
	} else {
		goto L326
	}
L326:
	;
	v1267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+80)))
	if v1265 == v1267 {
		goto L316
	} else {
		goto L327
	}
L327:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v1273 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(48))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(325), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L331:
	;
	if v1296 == int32(0) {
		goto L22
	} else {
		goto L332
	}
L332:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v1301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)))
	v1302 = F_PageIndexTupleOverwrite(m, v1300, v1301, v550, v549)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	if v1302 != 0 {
		goto L315
	} else {
		goto L334
	}
L334:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L5
	} else {
		goto L335
	}
L335:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1308 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34)
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L5
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(320), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1319 = m.ExcPending
	if v1319 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	v1325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+82)))
	if v1323 != v1325 {
		goto L21
	} else {
		goto L339
	}
L339:
	;
	goto L315
L340:
	;
	if v1175 == int32(-1) {
		goto L348
	} else {
		goto L349
	}
L341:
	;
	F_MarkBufferDirty(m, v1168)
	mBase = m.M
	v1331 = m.ExcPending
	if v1331 != 0 {
		goto L5
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	F_GenericXLogFinish(m, v1172)
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L5
	} else {
		goto L347
	}
L344:
	;
	if v1168 == v1213 {
		goto L340
	} else {
		goto L345
	}
L345:
	;
	F_MarkBufferDirty(m, v1213)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L5
	} else {
		goto L346
	}
L346:
	;
	goto L340
L347:
	;
	goto L340
L348:
	;
	v1337 = v1232
	goto L350
L349:
	;
	v1337 = v1175
	goto L350
L350:
	;
	F_UnlockReleaseBuffer(m, v1168)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L5
	} else {
		goto L351
	}
L351:
	;
	if v1168 != v1213 {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	F_UnlockReleaseBuffer(m, v1213)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L5
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	if base.B2i32(v1337 == int32(-1))|base.B2i32(v1337 == v310) == int32(0) {
		goto L356
	} else {
		goto L357
	}
L355:
	;
	goto L354
L356:
	;
	v1349 = int32(0)
	F_HnswUpdateMetaPage(m, l0, v1349, v1349, v1337, v1349, l4)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L5
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	F_HnswUpdateNeighborsOnDisk(m, l0, l1, v70, v97, int32(0), l4)
	mBase = m.M
	v1356 = m.ExcPending
	if v1356 != 0 {
		goto L5
	} else {
		goto L360
	}
L359:
	;
	goto L358
L360:
	;
	if v96 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+65)))
	v1358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+65)))
	if base.Ui32(v1357) <= base.Ui32(v1358) {
		goto L23
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	F_HnswUpdateMetaPage(m, l0, int32(1), v70, int32(-1), int32(0), l4)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L5
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	goto L23
L366:
	;
	m.G0 = v34 + int32(80)
	return int32(1)
L367:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1408 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(16))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L5
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(317), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L5
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1426 + int32(4)
	F_errmsg_internal(m, int32(_a_F_HnswInsertTupleOnDisk_4), v34+int32(32))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L5
	} else {
		goto L371
	}
L371:
	;
	F_errfinish(m, int32(_a_F_HnswInsertTupleOnDisk_5), int32(328), int32(_a_F_HnswInsertTupleOnDisk_6))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L5
	} else {
		goto L372
	}
L372:
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
func F_HnswNormValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_DirectFunctionCall1Coll(m, v4, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
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
		F_LockBuffer(m, v12, int32(2))
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
