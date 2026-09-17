package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_ghstore_in_0), int32(102), int32(_a_F_ghstore_in_1), int32(_a_F_ghstore_in_2), int32(_a_F_ghstore_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ghstore_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v653 int32
	_ = v653
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v758 int32
	_ = v758
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1012 int32
	_ = v1012
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1085 int32
	_ = v1085
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1191 int32
	_ = v1191
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1358 int32
	_ = v1358
	var v1382 int32
	_ = v1382
	var v1394 int32
	_ = v1394
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1552 int32
	_ = v1552
	var v1564 int32
	_ = v1564
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1634 int32
	_ = v1634
	var v1649 int32
	_ = v1649
	var v1656 int32
	_ = v1656
	var v1665 int32
	_ = v1665
	var v1676 int32
	_ = v1676
	var v1683 int32
	_ = v1683
	var v1691 int32
	_ = v1691
	v2 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v31 == v2 {
		v48 = v2
	} else {
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
		if v35 == int32(0) {
			v48 = v2
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v38 != int32(7) {
				v48 = v2
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				if v41 != int32(17) {
					v48 = v2
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
					v48 = v44 ^ int32(1)
				}
			}
		}
	}
	if v48&int32(1) != 0 {
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v52 = F_get_fn_opclass_options(m, v51)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
			v57 = v56
			v59 = (v26 + int32(_a_F_ghstore_picksplit_0)) & int32(_a_F_ghstore_picksplit_1)
			v63 = v59<<(uint(int32(1))%32) + int32(4)
			v64 = F_palloc(m, v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29))) = v64
				v67 = F_palloc(m, v63)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v67
					if base.Ui32(int32(2)) <= base.Ui32(v59) {
						v73 = v25 + int32(4)
						v77 = int32(-1)
						v78 = v2
						v85 = int32(1)
						v87 = v2
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v85<<(uint(int32(4))%32))))
							v105 = v85 + int32(1)
							v106 = v105
							v107 = v77
							v108 = v78
							v116 = v105
							v117 = v87
							for {
								v130 = int32(4)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v116<<(uint(v130)%32))))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
								v142 = v140 & v130
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
								if v143&v130 != 0 {
									if v142 != 0 {
										v311 = int32(0)
									} else {
										v148 = v57 << (uint(int32(3)) % 32)
										if v57 <= int32(0) {
											v311 = v148
										} else {
											v247 = v133 + int32(8)
											v248 = v148
											v250 = v247
											v251 = int32(0)
											v254 = int32(0)
											for {
												v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
												v260 = int32(1)
												v295 = v251 + v259&v260 + int32(base.Ui32(v259)>>(uint(int32(7))%32)) + int32(base.Ui32(v259)>>(uint(v260)%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(2))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(3))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(4))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(5))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(6))%32))&v260
												v299 = v254 + v260
												if v299 != v57 {
													v250 = v250 + v260
													v251 = v295
													v254 = v299
													continue
												} else {
													break
												}
												break
											}
											v311 = v248 - v295
										}
									}
								} else {
									if v142 != 0 {
										v154 = v57 << (uint(int32(3)) % 32)
										if v57 <= int32(0) {
											v311 = v154
										} else {
											v247 = v103 + int32(8)
											v248 = v154
											v250 = v247
											v251 = int32(0)
											v254 = int32(0)
											for {
												v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
												v260 = int32(1)
												v295 = v251 + v259&v260 + int32(base.Ui32(v259)>>(uint(int32(7))%32)) + int32(base.Ui32(v259)>>(uint(v260)%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(2))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(3))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(4))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(5))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(6))%32))&v260
												v299 = v254 + v260
												if v299 != v57 {
													v250 = v250 + v260
													v251 = v295
													v254 = v299
													continue
												} else {
													break
												}
												break
											}
											v311 = v248 - v295
										}
									} else {
										if v57 <= int32(0) {
											v311 = int32(0)
										} else {
											v162 = int32(8)
											v163 = v133 + v162
											v165 = v103 + v162
											v166 = int32(0)
											v168 = int32(1)
											v170 = v57 << (uint(int32(3)) % 32)
											if v170 <= v168 {
												v173 = v168
											} else {
												v173 = v170
											}
											if v173 != int32(1) {
												v181 = v166
												v182 = v166
												v183 = int32(0)
												for {
													v191 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
													v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v191))))
													v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v191))))
													v196 = v193 ^ v195
													v198 = v182 & int32(6)
													v199 = int32(1)
													v208 = int32(base.Ui32(v196)>>(uint(v198|v199)%32))&v199 + (int32(base.Ui32(v196)>>(uint(v198)%32))&v199 + v181)
													v209 = int32(2)
													v210 = v182 + v209
													v212 = v183 + v209
													if v212 != v173&int32(2147483640) {
														v181 = v208
														v182 = v210
														v183 = v212
														continue
													} else {
														break
													}
													break
												}
												if v173&int32(1) == int32(0) {
													v238 = v208
												} else {
													v216 = v208
													v217 = v210
													v226 = int32(base.Ui32(v217) >> (uint(int32(3)) % 32))
													v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v226))))
													v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v165))))
													v238 = v216 + int32(base.Ui32(v228^v230)>>(uint(v217&int32(7))%32))&int32(1)
												}
											} else {
												v216 = v166
												v217 = v166
												v226 = int32(base.Ui32(v217) >> (uint(int32(3)) % 32))
												v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v226))))
												v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v165))))
												v238 = v216 + int32(base.Ui32(v228^v230)>>(uint(v217&int32(7))%32))&int32(1)
											}
											v311 = v238
										}
									}
								}
								v312 = base.B2i32(v107 < v311)
								if v107 < v311 {
									v313 = v311
								} else {
									v313 = v107
								}
								if v107 < v311 {
									v314 = v106
								} else {
									v314 = v117
								}
								if v107 < v311 {
									v315 = v85
								} else {
									v315 = v108
								}
								v317 = v106 + int32(1)
								v319 = v317 & int32(_a_F_ghstore_picksplit_1)
								if base.Ui32(v319) <= base.Ui32(v59) {
									v106 = v317
									v107 = v313
									v108 = v315
									v116 = v319
									v117 = v314
									continue
								} else {
									break
								}
								break
							}
							if v59 != v105 {
								v77 = v313
								v78 = v315
								v85 = v105
								v87 = v314
								continue
							} else {
								break
							}
							break
						}
						v324 = v315
						v333 = v314
					} else {
						v324 = v2
						v333 = v2
					}
					v346 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v346
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v346
					v350 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v351 = int32(8)
					v353 = v57 + v351
					v355 = v25 + int32(4)
					v357 = int32(_a_F_ghstore_picksplit_1)
					v365 = base.B2i32(v324&v357 == v346) | base.B2i32(v333&v357 == v346)
					if v365 != 0 {
						v366 = int32(1)
					} else {
						v366 = v324
					}
					v369 = int32(4)
					v372 = *(*int32)(unsafe.Add(mBase, uint32(v355+v366&int32(_a_F_ghstore_picksplit_1)<<(uint(v369)%32))))
					v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
					v375 = v373 & v369
					if v375 != 0 {
						v376 = v351
					} else {
						v376 = v353
					}
					v377 = F_palloc(m, v376)
					mBase = m.M
					v378 = m.ExcPending
					if v378 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v375
						*(*int32)(unsafe.Add(mBase, uint32(v377))) = v376 << (uint(int32(2)) % 32)
						v383 = int32(0)
						if v375|base.B2i32(v57 == v383) == v383 {
							v388 = int32(8)
							base.MemoryCopy(m, v377+v388, v372+v388, v57)
						} else {
						}
						if v365 != 0 {
							v395 = int32(2)
						} else {
							v395 = v333
						}
						v398 = int32(4)
						v401 = *(*int32)(unsafe.Add(mBase, uint32(v355+v395&int32(_a_F_ghstore_picksplit_1)<<(uint(v398)%32))))
						v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
						v404 = v402 & v398
						if v404 != 0 {
							v405 = int32(8)
						} else {
							v405 = v353
						}
						v406 = F_palloc(m, v405)
						mBase = m.M
						v407 = m.ExcPending
						if v407 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v406)+4)) = v404
							*(*int32)(unsafe.Add(mBase, uint32(v406))) = v405 << (uint(int32(2)) % 32)
							v412 = int32(0)
							if v404|base.B2i32(v57 == v412) == v412 {
								v417 = int32(8)
								base.MemoryCopy(m, v406+v417, v401+v417, v57)
							} else {
							}
							v422 = int32(_a_F_ghstore_picksplit_1)
							v423 = v26 + v422
							v425 = v423 & v422
							v428 = F_palloc(m, v425<<(uint(int32(3))%32))
							mBase = m.M
							v429 = m.ExcPending
							if v429 != 0 {
								return int32(0)
							} else {
								if v26&int32(_a_F_ghstore_picksplit_1) == int32(1) {
									F_pg_qsort(m, v428, v425, int32(8), int32(_a_F_ghstore_picksplit_2))
									mBase = m.M
									v437 = m.ExcPending
									if v437 != 0 {
										return int32(0)
									} else {
										v1676 = v350
										v1683 = v67
										v1691 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1676))) = uint16(v1691)
										*(*uint16)(unsafe.Add(mBase, uint32(v1683))) = uint16(v1691)
										*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v406
										*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v377
										return v29
									}
								} else {
									v438 = int32(1)
									v440 = v438
									v441 = v438
									for {
										v466 = v428 + v440<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v466-int32(8)))) = uint16(v441)
										v470 = int32(4)
										v475 = *(*int32)(unsafe.Add(mBase, uint32(v355+v440<<(uint(v470)%32))))
										v482 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
										v484 = v482 & v470
										v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
										if v485&v470 != 0 {
											if v484 != 0 {
												v653 = int32(0)
											} else {
												v490 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v653 = v490
												} else {
													v589 = v475 + int32(8)
													v590 = v490
													v592 = v589
													v593 = int32(0)
													v596 = int32(0)
													for {
														v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
														v602 = int32(1)
														v637 = v593 + v601&v602 + int32(base.Ui32(v601)>>(uint(int32(7))%32)) + int32(base.Ui32(v601)>>(uint(v602)%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(2))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(3))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(4))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(5))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(6))%32))&v602
														v641 = v596 + v602
														if v641 != v57 {
															v592 = v592 + v602
															v593 = v637
															v596 = v641
															continue
														} else {
															break
														}
														break
													}
													v653 = v590 - v637
												}
											}
										} else {
											if v484 != 0 {
												v496 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v653 = v496
												} else {
													v589 = v377 + int32(8)
													v590 = v496
													v592 = v589
													v593 = int32(0)
													v596 = int32(0)
													for {
														v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
														v602 = int32(1)
														v637 = v593 + v601&v602 + int32(base.Ui32(v601)>>(uint(int32(7))%32)) + int32(base.Ui32(v601)>>(uint(v602)%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(2))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(3))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(4))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(5))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(6))%32))&v602
														v641 = v596 + v602
														if v641 != v57 {
															v592 = v592 + v602
															v593 = v637
															v596 = v641
															continue
														} else {
															break
														}
														break
													}
													v653 = v590 - v637
												}
											} else {
												if v57 <= int32(0) {
													v653 = int32(0)
												} else {
													v504 = int32(8)
													v505 = v475 + v504
													v507 = v377 + v504
													v508 = int32(0)
													v510 = int32(1)
													v512 = v57 << (uint(int32(3)) % 32)
													if v512 <= v510 {
														v515 = v510
													} else {
														v515 = v512
													}
													if v515 != int32(1) {
														v523 = v508
														v524 = v508
														v525 = int32(0)
														for {
															v533 = int32(base.Ui32(v524) >> (uint(int32(3)) % 32))
															v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v533))))
															v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507+v533))))
															v538 = v535 ^ v537
															v540 = v524 & int32(6)
															v541 = int32(1)
															v550 = int32(base.Ui32(v538)>>(uint(v540|v541)%32))&v541 + (int32(base.Ui32(v538)>>(uint(v540)%32))&v541 + v523)
															v551 = int32(2)
															v552 = v524 + v551
															v554 = v525 + v551
															if v554 != v515&int32(2147483640) {
																v523 = v550
																v524 = v552
																v525 = v554
																continue
															} else {
																break
															}
															break
														}
														if v515&int32(1) == int32(0) {
															v580 = v550
														} else {
															v558 = v550
															v559 = v552
															v568 = int32(base.Ui32(v559) >> (uint(int32(3)) % 32))
															v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v568))))
															v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v507))))
															v580 = v558 + int32(base.Ui32(v570^v572)>>(uint(v559&int32(7))%32))&int32(1)
														}
													} else {
														v558 = v508
														v559 = v508
														v568 = int32(base.Ui32(v559) >> (uint(int32(3)) % 32))
														v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v568))))
														v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v507))))
														v580 = v558 + int32(base.Ui32(v570^v572)>>(uint(v559&int32(7))%32))&int32(1)
													}
													v653 = v580
												}
											}
										}
										v660 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
										v661 = int32(4)
										v662 = v660 & v661
										v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
										if v663&v661 != 0 {
											if v662 != 0 {
												v831 = int32(0)
											} else {
												v668 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v831 = v668
												} else {
													v767 = v475 + int32(8)
													v768 = v668
													v770 = v767
													v771 = int32(0)
													v774 = int32(0)
													for {
														v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
														v780 = int32(1)
														v815 = v771 + v779&v780 + int32(base.Ui32(v779)>>(uint(int32(7))%32)) + int32(base.Ui32(v779)>>(uint(v780)%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(2))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(3))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(4))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(5))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(6))%32))&v780
														v819 = v774 + v780
														if v819 != v57 {
															v770 = v770 + v780
															v771 = v815
															v774 = v819
															continue
														} else {
															break
														}
														break
													}
													v831 = v768 - v815
												}
											}
										} else {
											if v662 != 0 {
												v674 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v831 = v674
												} else {
													v767 = v406 + int32(8)
													v768 = v674
													v770 = v767
													v771 = int32(0)
													v774 = int32(0)
													for {
														v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
														v780 = int32(1)
														v815 = v771 + v779&v780 + int32(base.Ui32(v779)>>(uint(int32(7))%32)) + int32(base.Ui32(v779)>>(uint(v780)%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(2))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(3))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(4))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(5))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(6))%32))&v780
														v819 = v774 + v780
														if v819 != v57 {
															v770 = v770 + v780
															v771 = v815
															v774 = v819
															continue
														} else {
															break
														}
														break
													}
													v831 = v768 - v815
												}
											} else {
												if v57 <= int32(0) {
													v831 = int32(0)
												} else {
													v682 = int32(8)
													v683 = v475 + v682
													v685 = v406 + v682
													v686 = int32(0)
													v688 = int32(1)
													v690 = v57 << (uint(int32(3)) % 32)
													if v690 <= v688 {
														v693 = v688
													} else {
														v693 = v690
													}
													if v693 != int32(1) {
														v701 = v686
														v702 = v686
														v703 = int32(0)
														for {
															v711 = int32(base.Ui32(v702) >> (uint(int32(3)) % 32))
															v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v711))))
															v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685+v711))))
															v716 = v713 ^ v715
															v718 = v702 & int32(6)
															v719 = int32(1)
															v728 = int32(base.Ui32(v716)>>(uint(v718|v719)%32))&v719 + (int32(base.Ui32(v716)>>(uint(v718)%32))&v719 + v701)
															v729 = int32(2)
															v730 = v702 + v729
															v732 = v703 + v729
															if v732 != v693&int32(2147483640) {
																v701 = v728
																v702 = v730
																v703 = v732
																continue
															} else {
																break
															}
															break
														}
														if v693&int32(1) == int32(0) {
															v758 = v728
														} else {
															v736 = v728
															v737 = v730
															v746 = int32(base.Ui32(v737) >> (uint(int32(3)) % 32))
															v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v746))))
															v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v685))))
															v758 = v736 + int32(base.Ui32(v748^v750)>>(uint(v737&int32(7))%32))&int32(1)
														}
													} else {
														v736 = v686
														v737 = v686
														v746 = int32(base.Ui32(v737) >> (uint(int32(3)) % 32))
														v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v746))))
														v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v685))))
														v758 = v736 + int32(base.Ui32(v748^v750)>>(uint(v737&int32(7))%32))&int32(1)
													}
													v831 = v758
												}
											}
										}
										v832 = v653 - v831
										v834 = v832 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v466-v470))) = v832 ^ v834 - v834
										v839 = v441 + int32(1)
										v840 = int32(_a_F_ghstore_picksplit_1)
										v841 = v839 & v840
										if base.Ui32(v841) <= base.Ui32(v423&v840) {
											v440 = v841
											v441 = v839
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v428, v425, int32(8), int32(_a_F_ghstore_picksplit_2))
									mBase = m.M
									v848 = m.ExcPending
									if v848 != 0 {
										return int32(0)
									} else {
										v849 = int32(1)
										if base.Ui32(v425) <= base.Ui32(v849) {
											v852 = v849
										} else {
											v852 = v425
										}
										v854 = v57 & int32(2147483644)
										v856 = v57 & int32(3)
										v857 = int32(8)
										v858 = v406 + v857
										v860 = v377 + v857
										v864 = int32(0)
										v873 = v350
										v880 = v67
										for {
											v891 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428+v864<<(uint(int32(3))%32)))))
											if v366&int32(_a_F_ghstore_picksplit_1) == v891 {
												*(*uint16)(unsafe.Add(mBase, uint32(v873))) = uint16(v366)
												v894 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v894 + int32(1)
												v1649 = v873 + int32(2)
												v1656 = v880
											} else {
												if v395&int32(_a_F_ghstore_picksplit_1) == v891 {
													*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v395)
													v1634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1634 + int32(1)
													v1649 = v873
													v1656 = v880 + int32(2)
												} else {
													v904 = int32(4)
													v907 = *(*int32)(unsafe.Add(mBase, uint32(v355+v891<<(uint(v904)%32))))
													v914 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
													v916 = v914 & v904
													v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
													if v917&v904 != 0 {
														if v916 != 0 {
															v1085 = int32(0)
														} else {
															v922 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1085 = v922
															} else {
																v1021 = v907 + int32(8)
																v1022 = v922
																v1024 = v1021
																v1025 = int32(0)
																v1028 = int32(0)
																for {
																	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
																	v1034 = int32(1)
																	v1069 = v1025 + v1033&v1034 + int32(base.Ui32(v1033)>>(uint(int32(7))%32)) + int32(base.Ui32(v1033)>>(uint(v1034)%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(2))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(3))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(4))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(5))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(6))%32))&v1034
																	v1073 = v1028 + v1034
																	if v1073 != v57 {
																		v1024 = v1024 + v1034
																		v1025 = v1069
																		v1028 = v1073
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1085 = v1022 - v1069
															}
														}
													} else {
														if v916 != 0 {
															v928 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1085 = v928
															} else {
																v1021 = v377 + int32(8)
																v1022 = v928
																v1024 = v1021
																v1025 = int32(0)
																v1028 = int32(0)
																for {
																	v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
																	v1034 = int32(1)
																	v1069 = v1025 + v1033&v1034 + int32(base.Ui32(v1033)>>(uint(int32(7))%32)) + int32(base.Ui32(v1033)>>(uint(v1034)%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(2))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(3))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(4))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(5))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(6))%32))&v1034
																	v1073 = v1028 + v1034
																	if v1073 != v57 {
																		v1024 = v1024 + v1034
																		v1025 = v1069
																		v1028 = v1073
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1085 = v1022 - v1069
															}
														} else {
															if v57 <= int32(0) {
																v1085 = int32(0)
															} else {
																v936 = int32(8)
																v937 = v907 + v936
																v939 = v377 + v936
																v940 = int32(0)
																v942 = int32(1)
																v944 = v57 << (uint(int32(3)) % 32)
																if v944 <= v942 {
																	v947 = v942
																} else {
																	v947 = v944
																}
																if v947 != int32(1) {
																	v955 = v940
																	v956 = v940
																	v957 = int32(0)
																	for {
																		v965 = int32(base.Ui32(v956) >> (uint(int32(3)) % 32))
																		v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v965))))
																		v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939+v965))))
																		v970 = v967 ^ v969
																		v972 = v956 & int32(6)
																		v973 = int32(1)
																		v982 = int32(base.Ui32(v970)>>(uint(v972|v973)%32))&v973 + (int32(base.Ui32(v970)>>(uint(v972)%32))&v973 + v955)
																		v983 = int32(2)
																		v984 = v956 + v983
																		v986 = v957 + v983
																		if v986 != v947&int32(2147483640) {
																			v955 = v982
																			v956 = v984
																			v957 = v986
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v947&int32(1) == int32(0) {
																		v1012 = v982
																	} else {
																		v990 = v982
																		v991 = v984
																		v1000 = int32(base.Ui32(v991) >> (uint(int32(3)) % 32))
																		v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v1000))))
																		v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000+v939))))
																		v1012 = v990 + int32(base.Ui32(v1002^v1004)>>(uint(v991&int32(7))%32))&int32(1)
																	}
																} else {
																	v990 = v940
																	v991 = v940
																	v1000 = int32(base.Ui32(v991) >> (uint(int32(3)) % 32))
																	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v1000))))
																	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000+v939))))
																	v1012 = v990 + int32(base.Ui32(v1002^v1004)>>(uint(v991&int32(7))%32))&int32(1)
																}
																v1085 = v1012
															}
														}
													}
													v1093 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
													v1094 = int32(4)
													v1095 = v1093 & v1094
													v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
													if v1096&v1094 != 0 {
														if v1095 != 0 {
															v1264 = int32(0)
														} else {
															v1101 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1264 = v1101
															} else {
																v1200 = v907 + int32(8)
																v1201 = v1101
																v1203 = v1200
																v1204 = int32(0)
																v1207 = int32(0)
																for {
																	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203))))
																	v1213 = int32(1)
																	v1248 = v1204 + v1212&v1213 + int32(base.Ui32(v1212)>>(uint(int32(7))%32)) + int32(base.Ui32(v1212)>>(uint(v1213)%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(2))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(3))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(4))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(5))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(6))%32))&v1213
																	v1252 = v1207 + v1213
																	if v1252 != v57 {
																		v1203 = v1203 + v1213
																		v1204 = v1248
																		v1207 = v1252
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1264 = v1201 - v1248
															}
														}
													} else {
														if v1095 != 0 {
															v1107 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1264 = v1107
															} else {
																v1200 = v406 + int32(8)
																v1201 = v1107
																v1203 = v1200
																v1204 = int32(0)
																v1207 = int32(0)
																for {
																	v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203))))
																	v1213 = int32(1)
																	v1248 = v1204 + v1212&v1213 + int32(base.Ui32(v1212)>>(uint(int32(7))%32)) + int32(base.Ui32(v1212)>>(uint(v1213)%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(2))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(3))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(4))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(5))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(6))%32))&v1213
																	v1252 = v1207 + v1213
																	if v1252 != v57 {
																		v1203 = v1203 + v1213
																		v1204 = v1248
																		v1207 = v1252
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1264 = v1201 - v1248
															}
														} else {
															if v57 <= int32(0) {
																v1264 = int32(0)
															} else {
																v1115 = int32(8)
																v1116 = v907 + v1115
																v1118 = v406 + v1115
																v1119 = int32(0)
																v1121 = int32(1)
																v1123 = v57 << (uint(int32(3)) % 32)
																if v1123 <= v1121 {
																	v1126 = v1121
																} else {
																	v1126 = v1123
																}
																if v1126 != int32(1) {
																	v1134 = v1119
																	v1135 = v1119
																	v1136 = int32(0)
																	for {
																		v1144 = int32(base.Ui32(v1135) >> (uint(int32(3)) % 32))
																		v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1144))))
																		v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118+v1144))))
																		v1149 = v1146 ^ v1148
																		v1151 = v1135 & int32(6)
																		v1152 = int32(1)
																		v1161 = int32(base.Ui32(v1149)>>(uint(v1151|v1152)%32))&v1152 + (int32(base.Ui32(v1149)>>(uint(v1151)%32))&v1152 + v1134)
																		v1162 = int32(2)
																		v1163 = v1135 + v1162
																		v1165 = v1136 + v1162
																		if v1165 != v1126&int32(2147483640) {
																			v1134 = v1161
																			v1135 = v1163
																			v1136 = v1165
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v1126&int32(1) == int32(0) {
																		v1191 = v1161
																	} else {
																		v1169 = v1161
																		v1170 = v1163
																		v1179 = int32(base.Ui32(v1170) >> (uint(int32(3)) % 32))
																		v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1179))))
																		v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179+v1118))))
																		v1191 = v1169 + int32(base.Ui32(v1181^v1183)>>(uint(v1170&int32(7))%32))&int32(1)
																	}
																} else {
																	v1169 = v1119
																	v1170 = v1119
																	v1179 = int32(base.Ui32(v1170) >> (uint(int32(3)) % 32))
																	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1179))))
																	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179+v1118))))
																	v1191 = v1169 + int32(base.Ui32(v1181^v1183)>>(uint(v1170&int32(7))%32))&int32(1)
																}
																v1264 = v1191
															}
														}
													}
													v1266 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													v1267 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													v1268 = v1266 - v1267
													if base.F64_lt(base.F64_convert_i32_s(v1085), base.F64_add(base.F64_convert_i32_s(v1264), base.F64_mul(base.F64_convert_i32_s(v1268*v1268*v1268), float64(-0.0001)))) != 0 {
														v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
														if v1276&int32(4) != 0 {
														} else {
															v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+4)))
															if v1279&int32(4) != 0 {
																if v57 == int32(0) {
																} else {
																	base.MemoryFill(m, v860, int32(255), v57)
																}
															} else {
																if v57 <= int32(0) {
																} else {
																	v1289 = v907 + int32(8)
																	v1290 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v57) {
																		v1296 = v1290
																		v1297 = v1290
																		for {
																			v1319 = v1296 + v860
																			v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319))))
																			v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1296+v1289))))
																			v1323 = v1320 | v1322
																			*(*uint8)(unsafe.Add(mBase, uint32(v1319))) = uint8(v1323)
																			v1326 = v1296 | int32(1)
																			v1327 = v860 + v1326
																			v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327))))
																			v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326+v1289))))
																			v1331 = v1328 | v1330
																			*(*uint8)(unsafe.Add(mBase, uint32(v1327))) = uint8(v1331)
																			v1334 = v1296 | int32(2)
																			v1335 = v860 + v1334
																			v1336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335))))
																			v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334+v1289))))
																			v1339 = v1336 | v1338
																			*(*uint8)(unsafe.Add(mBase, uint32(v1335))) = uint8(v1339)
																			v1342 = v1296 | int32(3)
																			v1343 = v860 + v1342
																			v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343))))
																			v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342+v1289))))
																			v1347 = v1344 | v1346
																			*(*uint8)(unsafe.Add(mBase, uint32(v1343))) = uint8(v1347)
																			v1349 = int32(4)
																			v1350 = v1296 + v1349
																			v1352 = v1297 + v1349
																			if v1352 != v854 {
																				v1296 = v1350
																				v1297 = v1352
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v856 == int32(0) {
																		} else {
																			v1358 = v1350
																			v1382 = v1358
																			v1394 = v1290
																			for {
																				v1404 = v1382 + v860
																				v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404))))
																				v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1289))))
																				v1408 = v1405 | v1407
																				*(*uint8)(unsafe.Add(mBase, uint32(v1404))) = uint8(v1408)
																				v1410 = int32(1)
																				v1413 = v1394 + v1410
																				if v1413 != v856 {
																					v1382 = v1382 + v1410
																					v1394 = v1413
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v1358 = v1290
																		v1382 = v1358
																		v1394 = v1290
																		for {
																			v1404 = v1382 + v860
																			v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404))))
																			v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1289))))
																			v1408 = v1405 | v1407
																			*(*uint8)(unsafe.Add(mBase, uint32(v1404))) = uint8(v1408)
																			v1410 = int32(1)
																			v1413 = v1394 + v1410
																			if v1413 != v856 {
																				v1382 = v1382 + v1410
																				v1394 = v1413
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v873))) = uint16(v891)
														v1440 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1440 + int32(1)
														v1649 = v873 + int32(2)
														v1656 = v880
													} else {
														v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
														if v1446&int32(4) != 0 {
														} else {
															v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+4)))
															if v1449&int32(4) != 0 {
																if v57 == int32(0) {
																} else {
																	base.MemoryFill(m, v858, int32(255), v57)
																}
															} else {
																if v57 <= int32(0) {
																} else {
																	v1459 = v907 + int32(8)
																	v1460 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v57) {
																		v1466 = v1460
																		v1467 = v1460
																		for {
																			v1489 = v1466 + v858
																			v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489))))
																			v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1466+v1459))))
																			v1493 = v1490 | v1492
																			*(*uint8)(unsafe.Add(mBase, uint32(v1489))) = uint8(v1493)
																			v1496 = v1466 | int32(1)
																			v1497 = v858 + v1496
																			v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497))))
																			v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1496+v1459))))
																			v1501 = v1498 | v1500
																			*(*uint8)(unsafe.Add(mBase, uint32(v1497))) = uint8(v1501)
																			v1504 = v1466 | int32(2)
																			v1505 = v858 + v1504
																			v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505))))
																			v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1504+v1459))))
																			v1509 = v1506 | v1508
																			*(*uint8)(unsafe.Add(mBase, uint32(v1505))) = uint8(v1509)
																			v1512 = v1466 | int32(3)
																			v1513 = v858 + v1512
																			v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513))))
																			v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512+v1459))))
																			v1517 = v1514 | v1516
																			*(*uint8)(unsafe.Add(mBase, uint32(v1513))) = uint8(v1517)
																			v1519 = int32(4)
																			v1520 = v1466 + v1519
																			v1522 = v1467 + v1519
																			if v1522 != v854 {
																				v1466 = v1520
																				v1467 = v1522
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v856 == int32(0) {
																		} else {
																			v1528 = v1520
																			v1552 = v1528
																			v1564 = v1460
																			for {
																				v1574 = v1552 + v858
																				v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574))))
																				v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552+v1459))))
																				v1578 = v1575 | v1577
																				*(*uint8)(unsafe.Add(mBase, uint32(v1574))) = uint8(v1578)
																				v1580 = int32(1)
																				v1583 = v1564 + v1580
																				if v1583 != v856 {
																					v1552 = v1552 + v1580
																					v1564 = v1583
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v1528 = v1460
																		v1552 = v1528
																		v1564 = v1460
																		for {
																			v1574 = v1552 + v858
																			v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574))))
																			v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552+v1459))))
																			v1578 = v1575 | v1577
																			*(*uint8)(unsafe.Add(mBase, uint32(v1574))) = uint8(v1578)
																			v1580 = int32(1)
																			v1583 = v1564 + v1580
																			if v1583 != v856 {
																				v1552 = v1552 + v1580
																				v1564 = v1583
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																}
															}
														}
														*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v891)
														v1634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1634 + int32(1)
														v1649 = v873
														v1656 = v880 + int32(2)
													}
												}
											}
											v1665 = v864 + int32(1)
											if v1665 != v852 {
												v864 = v1665
												v873 = v1649
												v880 = v1656
												continue
											} else {
												break
											}
											break
										}
										v1676 = v1649
										v1683 = v1656
										v1691 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1676))) = uint16(v1691)
										*(*uint16)(unsafe.Add(mBase, uint32(v1683))) = uint16(v1691)
										*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v406
										*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v377
										return v29
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v57 = int32(16)
		v59 = (v26 + int32(_a_F_ghstore_picksplit_0)) & int32(_a_F_ghstore_picksplit_1)
		v63 = v59<<(uint(int32(1))%32) + int32(4)
		v64 = F_palloc(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v64
			v67 = F_palloc(m, v63)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v67
				if base.Ui32(int32(2)) <= base.Ui32(v59) {
					v73 = v25 + int32(4)
					v77 = int32(-1)
					v78 = v2
					v85 = int32(1)
					v87 = v2
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v85<<(uint(int32(4))%32))))
						v105 = v85 + int32(1)
						v106 = v105
						v107 = v77
						v108 = v78
						v116 = v105
						v117 = v87
						for {
							v130 = int32(4)
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v116<<(uint(v130)%32))))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
							v142 = v140 & v130
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
							if v143&v130 != 0 {
								if v142 != 0 {
									v311 = int32(0)
								} else {
									v148 = v57 << (uint(int32(3)) % 32)
									if v57 <= int32(0) {
										v311 = v148
									} else {
										v247 = v133 + int32(8)
										v248 = v148
										v250 = v247
										v251 = int32(0)
										v254 = int32(0)
										for {
											v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
											v260 = int32(1)
											v295 = v251 + v259&v260 + int32(base.Ui32(v259)>>(uint(int32(7))%32)) + int32(base.Ui32(v259)>>(uint(v260)%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(2))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(3))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(4))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(5))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(6))%32))&v260
											v299 = v254 + v260
											if v299 != v57 {
												v250 = v250 + v260
												v251 = v295
												v254 = v299
												continue
											} else {
												break
											}
											break
										}
										v311 = v248 - v295
									}
								}
							} else {
								if v142 != 0 {
									v154 = v57 << (uint(int32(3)) % 32)
									if v57 <= int32(0) {
										v311 = v154
									} else {
										v247 = v103 + int32(8)
										v248 = v154
										v250 = v247
										v251 = int32(0)
										v254 = int32(0)
										for {
											v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
											v260 = int32(1)
											v295 = v251 + v259&v260 + int32(base.Ui32(v259)>>(uint(int32(7))%32)) + int32(base.Ui32(v259)>>(uint(v260)%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(2))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(3))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(4))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(5))%32))&v260 + int32(base.Ui32(v259)>>(uint(int32(6))%32))&v260
											v299 = v254 + v260
											if v299 != v57 {
												v250 = v250 + v260
												v251 = v295
												v254 = v299
												continue
											} else {
												break
											}
											break
										}
										v311 = v248 - v295
									}
								} else {
									if v57 <= int32(0) {
										v311 = int32(0)
									} else {
										v162 = int32(8)
										v163 = v133 + v162
										v165 = v103 + v162
										v166 = int32(0)
										v168 = int32(1)
										v170 = v57 << (uint(int32(3)) % 32)
										if v170 <= v168 {
											v173 = v168
										} else {
											v173 = v170
										}
										if v173 != int32(1) {
											v181 = v166
											v182 = v166
											v183 = int32(0)
											for {
												v191 = int32(base.Ui32(v182) >> (uint(int32(3)) % 32))
												v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v191))))
												v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v191))))
												v196 = v193 ^ v195
												v198 = v182 & int32(6)
												v199 = int32(1)
												v208 = int32(base.Ui32(v196)>>(uint(v198|v199)%32))&v199 + (int32(base.Ui32(v196)>>(uint(v198)%32))&v199 + v181)
												v209 = int32(2)
												v210 = v182 + v209
												v212 = v183 + v209
												if v212 != v173&int32(2147483640) {
													v181 = v208
													v182 = v210
													v183 = v212
													continue
												} else {
													break
												}
												break
											}
											if v173&int32(1) == int32(0) {
												v238 = v208
											} else {
												v216 = v208
												v217 = v210
												v226 = int32(base.Ui32(v217) >> (uint(int32(3)) % 32))
												v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v226))))
												v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v165))))
												v238 = v216 + int32(base.Ui32(v228^v230)>>(uint(v217&int32(7))%32))&int32(1)
											}
										} else {
											v216 = v166
											v217 = v166
											v226 = int32(base.Ui32(v217) >> (uint(int32(3)) % 32))
											v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v226))))
											v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226+v165))))
											v238 = v216 + int32(base.Ui32(v228^v230)>>(uint(v217&int32(7))%32))&int32(1)
										}
										v311 = v238
									}
								}
							}
							v312 = base.B2i32(v107 < v311)
							if v107 < v311 {
								v313 = v311
							} else {
								v313 = v107
							}
							if v107 < v311 {
								v314 = v106
							} else {
								v314 = v117
							}
							if v107 < v311 {
								v315 = v85
							} else {
								v315 = v108
							}
							v317 = v106 + int32(1)
							v319 = v317 & int32(_a_F_ghstore_picksplit_1)
							if base.Ui32(v319) <= base.Ui32(v59) {
								v106 = v317
								v107 = v313
								v108 = v315
								v116 = v319
								v117 = v314
								continue
							} else {
								break
							}
							break
						}
						if v59 != v105 {
							v77 = v313
							v78 = v315
							v85 = v105
							v87 = v314
							continue
						} else {
							break
						}
						break
					}
					v324 = v315
					v333 = v314
				} else {
					v324 = v2
					v333 = v2
				}
				v346 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v346
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v346
				v350 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v351 = int32(8)
				v353 = v57 + v351
				v355 = v25 + int32(4)
				v357 = int32(_a_F_ghstore_picksplit_1)
				v365 = base.B2i32(v324&v357 == v346) | base.B2i32(v333&v357 == v346)
				if v365 != 0 {
					v366 = int32(1)
				} else {
					v366 = v324
				}
				v369 = int32(4)
				v372 = *(*int32)(unsafe.Add(mBase, uint32(v355+v366&int32(_a_F_ghstore_picksplit_1)<<(uint(v369)%32))))
				v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
				v375 = v373 & v369
				if v375 != 0 {
					v376 = v351
				} else {
					v376 = v353
				}
				v377 = F_palloc(m, v376)
				mBase = m.M
				v378 = m.ExcPending
				if v378 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v377)+4)) = v375
					*(*int32)(unsafe.Add(mBase, uint32(v377))) = v376 << (uint(int32(2)) % 32)
					v383 = int32(0)
					if v375|base.B2i32(v57 == v383) == v383 {
						v388 = int32(8)
						base.MemoryCopy(m, v377+v388, v372+v388, v57)
					} else {
					}
					if v365 != 0 {
						v395 = int32(2)
					} else {
						v395 = v333
					}
					v398 = int32(4)
					v401 = *(*int32)(unsafe.Add(mBase, uint32(v355+v395&int32(_a_F_ghstore_picksplit_1)<<(uint(v398)%32))))
					v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
					v404 = v402 & v398
					if v404 != 0 {
						v405 = int32(8)
					} else {
						v405 = v353
					}
					v406 = F_palloc(m, v405)
					mBase = m.M
					v407 = m.ExcPending
					if v407 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v406)+4)) = v404
						*(*int32)(unsafe.Add(mBase, uint32(v406))) = v405 << (uint(int32(2)) % 32)
						v412 = int32(0)
						if v404|base.B2i32(v57 == v412) == v412 {
							v417 = int32(8)
							base.MemoryCopy(m, v406+v417, v401+v417, v57)
						} else {
						}
						v422 = int32(_a_F_ghstore_picksplit_1)
						v423 = v26 + v422
						v425 = v423 & v422
						v428 = F_palloc(m, v425<<(uint(int32(3))%32))
						mBase = m.M
						v429 = m.ExcPending
						if v429 != 0 {
							return int32(0)
						} else {
							if v26&int32(_a_F_ghstore_picksplit_1) == int32(1) {
								F_pg_qsort(m, v428, v425, int32(8), int32(_a_F_ghstore_picksplit_2))
								mBase = m.M
								v437 = m.ExcPending
								if v437 != 0 {
									return int32(0)
								} else {
									v1676 = v350
									v1683 = v67
									v1691 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1676))) = uint16(v1691)
									*(*uint16)(unsafe.Add(mBase, uint32(v1683))) = uint16(v1691)
									*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v406
									*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v377
									return v29
								}
							} else {
								v438 = int32(1)
								v440 = v438
								v441 = v438
								for {
									v466 = v428 + v440<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v466-int32(8)))) = uint16(v441)
									v470 = int32(4)
									v475 = *(*int32)(unsafe.Add(mBase, uint32(v355+v440<<(uint(v470)%32))))
									v482 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
									v484 = v482 & v470
									v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
									if v485&v470 != 0 {
										if v484 != 0 {
											v653 = int32(0)
										} else {
											v490 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v653 = v490
											} else {
												v589 = v475 + int32(8)
												v590 = v490
												v592 = v589
												v593 = int32(0)
												v596 = int32(0)
												for {
													v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
													v602 = int32(1)
													v637 = v593 + v601&v602 + int32(base.Ui32(v601)>>(uint(int32(7))%32)) + int32(base.Ui32(v601)>>(uint(v602)%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(2))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(3))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(4))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(5))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(6))%32))&v602
													v641 = v596 + v602
													if v641 != v57 {
														v592 = v592 + v602
														v593 = v637
														v596 = v641
														continue
													} else {
														break
													}
													break
												}
												v653 = v590 - v637
											}
										}
									} else {
										if v484 != 0 {
											v496 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v653 = v496
											} else {
												v589 = v377 + int32(8)
												v590 = v496
												v592 = v589
												v593 = int32(0)
												v596 = int32(0)
												for {
													v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592))))
													v602 = int32(1)
													v637 = v593 + v601&v602 + int32(base.Ui32(v601)>>(uint(int32(7))%32)) + int32(base.Ui32(v601)>>(uint(v602)%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(2))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(3))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(4))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(5))%32))&v602 + int32(base.Ui32(v601)>>(uint(int32(6))%32))&v602
													v641 = v596 + v602
													if v641 != v57 {
														v592 = v592 + v602
														v593 = v637
														v596 = v641
														continue
													} else {
														break
													}
													break
												}
												v653 = v590 - v637
											}
										} else {
											if v57 <= int32(0) {
												v653 = int32(0)
											} else {
												v504 = int32(8)
												v505 = v475 + v504
												v507 = v377 + v504
												v508 = int32(0)
												v510 = int32(1)
												v512 = v57 << (uint(int32(3)) % 32)
												if v512 <= v510 {
													v515 = v510
												} else {
													v515 = v512
												}
												if v515 != int32(1) {
													v523 = v508
													v524 = v508
													v525 = int32(0)
													for {
														v533 = int32(base.Ui32(v524) >> (uint(int32(3)) % 32))
														v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v533))))
														v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507+v533))))
														v538 = v535 ^ v537
														v540 = v524 & int32(6)
														v541 = int32(1)
														v550 = int32(base.Ui32(v538)>>(uint(v540|v541)%32))&v541 + (int32(base.Ui32(v538)>>(uint(v540)%32))&v541 + v523)
														v551 = int32(2)
														v552 = v524 + v551
														v554 = v525 + v551
														if v554 != v515&int32(2147483640) {
															v523 = v550
															v524 = v552
															v525 = v554
															continue
														} else {
															break
														}
														break
													}
													if v515&int32(1) == int32(0) {
														v580 = v550
													} else {
														v558 = v550
														v559 = v552
														v568 = int32(base.Ui32(v559) >> (uint(int32(3)) % 32))
														v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v568))))
														v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v507))))
														v580 = v558 + int32(base.Ui32(v570^v572)>>(uint(v559&int32(7))%32))&int32(1)
													}
												} else {
													v558 = v508
													v559 = v508
													v568 = int32(base.Ui32(v559) >> (uint(int32(3)) % 32))
													v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505+v568))))
													v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v568+v507))))
													v580 = v558 + int32(base.Ui32(v570^v572)>>(uint(v559&int32(7))%32))&int32(1)
												}
												v653 = v580
											}
										}
									}
									v660 = *(*int32)(unsafe.Add(mBase, uint32(v475)+4))
									v661 = int32(4)
									v662 = v660 & v661
									v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
									if v663&v661 != 0 {
										if v662 != 0 {
											v831 = int32(0)
										} else {
											v668 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v831 = v668
											} else {
												v767 = v475 + int32(8)
												v768 = v668
												v770 = v767
												v771 = int32(0)
												v774 = int32(0)
												for {
													v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
													v780 = int32(1)
													v815 = v771 + v779&v780 + int32(base.Ui32(v779)>>(uint(int32(7))%32)) + int32(base.Ui32(v779)>>(uint(v780)%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(2))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(3))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(4))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(5))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(6))%32))&v780
													v819 = v774 + v780
													if v819 != v57 {
														v770 = v770 + v780
														v771 = v815
														v774 = v819
														continue
													} else {
														break
													}
													break
												}
												v831 = v768 - v815
											}
										}
									} else {
										if v662 != 0 {
											v674 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v831 = v674
											} else {
												v767 = v406 + int32(8)
												v768 = v674
												v770 = v767
												v771 = int32(0)
												v774 = int32(0)
												for {
													v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
													v780 = int32(1)
													v815 = v771 + v779&v780 + int32(base.Ui32(v779)>>(uint(int32(7))%32)) + int32(base.Ui32(v779)>>(uint(v780)%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(2))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(3))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(4))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(5))%32))&v780 + int32(base.Ui32(v779)>>(uint(int32(6))%32))&v780
													v819 = v774 + v780
													if v819 != v57 {
														v770 = v770 + v780
														v771 = v815
														v774 = v819
														continue
													} else {
														break
													}
													break
												}
												v831 = v768 - v815
											}
										} else {
											if v57 <= int32(0) {
												v831 = int32(0)
											} else {
												v682 = int32(8)
												v683 = v475 + v682
												v685 = v406 + v682
												v686 = int32(0)
												v688 = int32(1)
												v690 = v57 << (uint(int32(3)) % 32)
												if v690 <= v688 {
													v693 = v688
												} else {
													v693 = v690
												}
												if v693 != int32(1) {
													v701 = v686
													v702 = v686
													v703 = int32(0)
													for {
														v711 = int32(base.Ui32(v702) >> (uint(int32(3)) % 32))
														v713 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v711))))
														v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685+v711))))
														v716 = v713 ^ v715
														v718 = v702 & int32(6)
														v719 = int32(1)
														v728 = int32(base.Ui32(v716)>>(uint(v718|v719)%32))&v719 + (int32(base.Ui32(v716)>>(uint(v718)%32))&v719 + v701)
														v729 = int32(2)
														v730 = v702 + v729
														v732 = v703 + v729
														if v732 != v693&int32(2147483640) {
															v701 = v728
															v702 = v730
															v703 = v732
															continue
														} else {
															break
														}
														break
													}
													if v693&int32(1) == int32(0) {
														v758 = v728
													} else {
														v736 = v728
														v737 = v730
														v746 = int32(base.Ui32(v737) >> (uint(int32(3)) % 32))
														v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v746))))
														v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v685))))
														v758 = v736 + int32(base.Ui32(v748^v750)>>(uint(v737&int32(7))%32))&int32(1)
													}
												} else {
													v736 = v686
													v737 = v686
													v746 = int32(base.Ui32(v737) >> (uint(int32(3)) % 32))
													v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v746))))
													v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v746+v685))))
													v758 = v736 + int32(base.Ui32(v748^v750)>>(uint(v737&int32(7))%32))&int32(1)
												}
												v831 = v758
											}
										}
									}
									v832 = v653 - v831
									v834 = v832 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v466-v470))) = v832 ^ v834 - v834
									v839 = v441 + int32(1)
									v840 = int32(_a_F_ghstore_picksplit_1)
									v841 = v839 & v840
									if base.Ui32(v841) <= base.Ui32(v423&v840) {
										v440 = v841
										v441 = v839
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v428, v425, int32(8), int32(_a_F_ghstore_picksplit_2))
								mBase = m.M
								v848 = m.ExcPending
								if v848 != 0 {
									return int32(0)
								} else {
									v849 = int32(1)
									if base.Ui32(v425) <= base.Ui32(v849) {
										v852 = v849
									} else {
										v852 = v425
									}
									v854 = v57 & int32(2147483644)
									v856 = v57 & int32(3)
									v857 = int32(8)
									v858 = v406 + v857
									v860 = v377 + v857
									v864 = int32(0)
									v873 = v350
									v880 = v67
									for {
										v891 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v428+v864<<(uint(int32(3))%32)))))
										if v366&int32(_a_F_ghstore_picksplit_1) == v891 {
											*(*uint16)(unsafe.Add(mBase, uint32(v873))) = uint16(v366)
											v894 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v894 + int32(1)
											v1649 = v873 + int32(2)
											v1656 = v880
										} else {
											if v395&int32(_a_F_ghstore_picksplit_1) == v891 {
												*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v395)
												v1634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1634 + int32(1)
												v1649 = v873
												v1656 = v880 + int32(2)
											} else {
												v904 = int32(4)
												v907 = *(*int32)(unsafe.Add(mBase, uint32(v355+v891<<(uint(v904)%32))))
												v914 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
												v916 = v914 & v904
												v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
												if v917&v904 != 0 {
													if v916 != 0 {
														v1085 = int32(0)
													} else {
														v922 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1085 = v922
														} else {
															v1021 = v907 + int32(8)
															v1022 = v922
															v1024 = v1021
															v1025 = int32(0)
															v1028 = int32(0)
															for {
																v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
																v1034 = int32(1)
																v1069 = v1025 + v1033&v1034 + int32(base.Ui32(v1033)>>(uint(int32(7))%32)) + int32(base.Ui32(v1033)>>(uint(v1034)%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(2))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(3))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(4))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(5))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(6))%32))&v1034
																v1073 = v1028 + v1034
																if v1073 != v57 {
																	v1024 = v1024 + v1034
																	v1025 = v1069
																	v1028 = v1073
																	continue
																} else {
																	break
																}
																break
															}
															v1085 = v1022 - v1069
														}
													}
												} else {
													if v916 != 0 {
														v928 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1085 = v928
														} else {
															v1021 = v377 + int32(8)
															v1022 = v928
															v1024 = v1021
															v1025 = int32(0)
															v1028 = int32(0)
															for {
																v1033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024))))
																v1034 = int32(1)
																v1069 = v1025 + v1033&v1034 + int32(base.Ui32(v1033)>>(uint(int32(7))%32)) + int32(base.Ui32(v1033)>>(uint(v1034)%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(2))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(3))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(4))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(5))%32))&v1034 + int32(base.Ui32(v1033)>>(uint(int32(6))%32))&v1034
																v1073 = v1028 + v1034
																if v1073 != v57 {
																	v1024 = v1024 + v1034
																	v1025 = v1069
																	v1028 = v1073
																	continue
																} else {
																	break
																}
																break
															}
															v1085 = v1022 - v1069
														}
													} else {
														if v57 <= int32(0) {
															v1085 = int32(0)
														} else {
															v936 = int32(8)
															v937 = v907 + v936
															v939 = v377 + v936
															v940 = int32(0)
															v942 = int32(1)
															v944 = v57 << (uint(int32(3)) % 32)
															if v944 <= v942 {
																v947 = v942
															} else {
																v947 = v944
															}
															if v947 != int32(1) {
																v955 = v940
																v956 = v940
																v957 = int32(0)
																for {
																	v965 = int32(base.Ui32(v956) >> (uint(int32(3)) % 32))
																	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v965))))
																	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939+v965))))
																	v970 = v967 ^ v969
																	v972 = v956 & int32(6)
																	v973 = int32(1)
																	v982 = int32(base.Ui32(v970)>>(uint(v972|v973)%32))&v973 + (int32(base.Ui32(v970)>>(uint(v972)%32))&v973 + v955)
																	v983 = int32(2)
																	v984 = v956 + v983
																	v986 = v957 + v983
																	if v986 != v947&int32(2147483640) {
																		v955 = v982
																		v956 = v984
																		v957 = v986
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v947&int32(1) == int32(0) {
																	v1012 = v982
																} else {
																	v990 = v982
																	v991 = v984
																	v1000 = int32(base.Ui32(v991) >> (uint(int32(3)) % 32))
																	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v1000))))
																	v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000+v939))))
																	v1012 = v990 + int32(base.Ui32(v1002^v1004)>>(uint(v991&int32(7))%32))&int32(1)
																}
															} else {
																v990 = v940
																v991 = v940
																v1000 = int32(base.Ui32(v991) >> (uint(int32(3)) % 32))
																v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937+v1000))))
																v1004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000+v939))))
																v1012 = v990 + int32(base.Ui32(v1002^v1004)>>(uint(v991&int32(7))%32))&int32(1)
															}
															v1085 = v1012
														}
													}
												}
												v1093 = *(*int32)(unsafe.Add(mBase, uint32(v907)+4))
												v1094 = int32(4)
												v1095 = v1093 & v1094
												v1096 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
												if v1096&v1094 != 0 {
													if v1095 != 0 {
														v1264 = int32(0)
													} else {
														v1101 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1264 = v1101
														} else {
															v1200 = v907 + int32(8)
															v1201 = v1101
															v1203 = v1200
															v1204 = int32(0)
															v1207 = int32(0)
															for {
																v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203))))
																v1213 = int32(1)
																v1248 = v1204 + v1212&v1213 + int32(base.Ui32(v1212)>>(uint(int32(7))%32)) + int32(base.Ui32(v1212)>>(uint(v1213)%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(2))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(3))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(4))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(5))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(6))%32))&v1213
																v1252 = v1207 + v1213
																if v1252 != v57 {
																	v1203 = v1203 + v1213
																	v1204 = v1248
																	v1207 = v1252
																	continue
																} else {
																	break
																}
																break
															}
															v1264 = v1201 - v1248
														}
													}
												} else {
													if v1095 != 0 {
														v1107 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1264 = v1107
														} else {
															v1200 = v406 + int32(8)
															v1201 = v1107
															v1203 = v1200
															v1204 = int32(0)
															v1207 = int32(0)
															for {
																v1212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203))))
																v1213 = int32(1)
																v1248 = v1204 + v1212&v1213 + int32(base.Ui32(v1212)>>(uint(int32(7))%32)) + int32(base.Ui32(v1212)>>(uint(v1213)%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(2))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(3))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(4))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(5))%32))&v1213 + int32(base.Ui32(v1212)>>(uint(int32(6))%32))&v1213
																v1252 = v1207 + v1213
																if v1252 != v57 {
																	v1203 = v1203 + v1213
																	v1204 = v1248
																	v1207 = v1252
																	continue
																} else {
																	break
																}
																break
															}
															v1264 = v1201 - v1248
														}
													} else {
														if v57 <= int32(0) {
															v1264 = int32(0)
														} else {
															v1115 = int32(8)
															v1116 = v907 + v1115
															v1118 = v406 + v1115
															v1119 = int32(0)
															v1121 = int32(1)
															v1123 = v57 << (uint(int32(3)) % 32)
															if v1123 <= v1121 {
																v1126 = v1121
															} else {
																v1126 = v1123
															}
															if v1126 != int32(1) {
																v1134 = v1119
																v1135 = v1119
																v1136 = int32(0)
																for {
																	v1144 = int32(base.Ui32(v1135) >> (uint(int32(3)) % 32))
																	v1146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1144))))
																	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1118+v1144))))
																	v1149 = v1146 ^ v1148
																	v1151 = v1135 & int32(6)
																	v1152 = int32(1)
																	v1161 = int32(base.Ui32(v1149)>>(uint(v1151|v1152)%32))&v1152 + (int32(base.Ui32(v1149)>>(uint(v1151)%32))&v1152 + v1134)
																	v1162 = int32(2)
																	v1163 = v1135 + v1162
																	v1165 = v1136 + v1162
																	if v1165 != v1126&int32(2147483640) {
																		v1134 = v1161
																		v1135 = v1163
																		v1136 = v1165
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v1126&int32(1) == int32(0) {
																	v1191 = v1161
																} else {
																	v1169 = v1161
																	v1170 = v1163
																	v1179 = int32(base.Ui32(v1170) >> (uint(int32(3)) % 32))
																	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1179))))
																	v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179+v1118))))
																	v1191 = v1169 + int32(base.Ui32(v1181^v1183)>>(uint(v1170&int32(7))%32))&int32(1)
																}
															} else {
																v1169 = v1119
																v1170 = v1119
																v1179 = int32(base.Ui32(v1170) >> (uint(int32(3)) % 32))
																v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116+v1179))))
																v1183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1179+v1118))))
																v1191 = v1169 + int32(base.Ui32(v1181^v1183)>>(uint(v1170&int32(7))%32))&int32(1)
															}
															v1264 = v1191
														}
													}
												}
												v1266 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												v1267 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
												v1268 = v1266 - v1267
												if base.F64_lt(base.F64_convert_i32_s(v1085), base.F64_add(base.F64_convert_i32_s(v1264), base.F64_mul(base.F64_convert_i32_s(v1268*v1268*v1268), float64(-0.0001)))) != 0 {
													v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v377)+4)))
													if v1276&int32(4) != 0 {
													} else {
														v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+4)))
														if v1279&int32(4) != 0 {
															if v57 == int32(0) {
															} else {
																base.MemoryFill(m, v860, int32(255), v57)
															}
														} else {
															if v57 <= int32(0) {
															} else {
																v1289 = v907 + int32(8)
																v1290 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v57) {
																	v1296 = v1290
																	v1297 = v1290
																	for {
																		v1319 = v1296 + v860
																		v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319))))
																		v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1296+v1289))))
																		v1323 = v1320 | v1322
																		*(*uint8)(unsafe.Add(mBase, uint32(v1319))) = uint8(v1323)
																		v1326 = v1296 | int32(1)
																		v1327 = v860 + v1326
																		v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1327))))
																		v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326+v1289))))
																		v1331 = v1328 | v1330
																		*(*uint8)(unsafe.Add(mBase, uint32(v1327))) = uint8(v1331)
																		v1334 = v1296 | int32(2)
																		v1335 = v860 + v1334
																		v1336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335))))
																		v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1334+v1289))))
																		v1339 = v1336 | v1338
																		*(*uint8)(unsafe.Add(mBase, uint32(v1335))) = uint8(v1339)
																		v1342 = v1296 | int32(3)
																		v1343 = v860 + v1342
																		v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343))))
																		v1346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342+v1289))))
																		v1347 = v1344 | v1346
																		*(*uint8)(unsafe.Add(mBase, uint32(v1343))) = uint8(v1347)
																		v1349 = int32(4)
																		v1350 = v1296 + v1349
																		v1352 = v1297 + v1349
																		if v1352 != v854 {
																			v1296 = v1350
																			v1297 = v1352
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v856 == int32(0) {
																	} else {
																		v1358 = v1350
																		v1382 = v1358
																		v1394 = v1290
																		for {
																			v1404 = v1382 + v860
																			v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404))))
																			v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1289))))
																			v1408 = v1405 | v1407
																			*(*uint8)(unsafe.Add(mBase, uint32(v1404))) = uint8(v1408)
																			v1410 = int32(1)
																			v1413 = v1394 + v1410
																			if v1413 != v856 {
																				v1382 = v1382 + v1410
																				v1394 = v1413
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v1358 = v1290
																	v1382 = v1358
																	v1394 = v1290
																	for {
																		v1404 = v1382 + v860
																		v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1404))))
																		v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1382+v1289))))
																		v1408 = v1405 | v1407
																		*(*uint8)(unsafe.Add(mBase, uint32(v1404))) = uint8(v1408)
																		v1410 = int32(1)
																		v1413 = v1394 + v1410
																		if v1413 != v856 {
																			v1382 = v1382 + v1410
																			v1394 = v1413
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v873))) = uint16(v891)
													v1440 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1440 + int32(1)
													v1649 = v873 + int32(2)
													v1656 = v880
												} else {
													v1446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406)+4)))
													if v1446&int32(4) != 0 {
													} else {
														v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907)+4)))
														if v1449&int32(4) != 0 {
															if v57 == int32(0) {
															} else {
																base.MemoryFill(m, v858, int32(255), v57)
															}
														} else {
															if v57 <= int32(0) {
															} else {
																v1459 = v907 + int32(8)
																v1460 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v57) {
																	v1466 = v1460
																	v1467 = v1460
																	for {
																		v1489 = v1466 + v858
																		v1490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489))))
																		v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1466+v1459))))
																		v1493 = v1490 | v1492
																		*(*uint8)(unsafe.Add(mBase, uint32(v1489))) = uint8(v1493)
																		v1496 = v1466 | int32(1)
																		v1497 = v858 + v1496
																		v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1497))))
																		v1500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1496+v1459))))
																		v1501 = v1498 | v1500
																		*(*uint8)(unsafe.Add(mBase, uint32(v1497))) = uint8(v1501)
																		v1504 = v1466 | int32(2)
																		v1505 = v858 + v1504
																		v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1505))))
																		v1508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1504+v1459))))
																		v1509 = v1506 | v1508
																		*(*uint8)(unsafe.Add(mBase, uint32(v1505))) = uint8(v1509)
																		v1512 = v1466 | int32(3)
																		v1513 = v858 + v1512
																		v1514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1513))))
																		v1516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1512+v1459))))
																		v1517 = v1514 | v1516
																		*(*uint8)(unsafe.Add(mBase, uint32(v1513))) = uint8(v1517)
																		v1519 = int32(4)
																		v1520 = v1466 + v1519
																		v1522 = v1467 + v1519
																		if v1522 != v854 {
																			v1466 = v1520
																			v1467 = v1522
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v856 == int32(0) {
																	} else {
																		v1528 = v1520
																		v1552 = v1528
																		v1564 = v1460
																		for {
																			v1574 = v1552 + v858
																			v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574))))
																			v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552+v1459))))
																			v1578 = v1575 | v1577
																			*(*uint8)(unsafe.Add(mBase, uint32(v1574))) = uint8(v1578)
																			v1580 = int32(1)
																			v1583 = v1564 + v1580
																			if v1583 != v856 {
																				v1552 = v1552 + v1580
																				v1564 = v1583
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v1528 = v1460
																	v1552 = v1528
																	v1564 = v1460
																	for {
																		v1574 = v1552 + v858
																		v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574))))
																		v1577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1552+v1459))))
																		v1578 = v1575 | v1577
																		*(*uint8)(unsafe.Add(mBase, uint32(v1574))) = uint8(v1578)
																		v1580 = int32(1)
																		v1583 = v1564 + v1580
																		if v1583 != v856 {
																			v1552 = v1552 + v1580
																			v1564 = v1583
																			continue
																		} else {
																			break
																		}
																		break
																	}
																}
															}
														}
													}
													*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v891)
													v1634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1634 + int32(1)
													v1649 = v873
													v1656 = v880 + int32(2)
												}
											}
										}
										v1665 = v864 + int32(1)
										if v1665 != v852 {
											v864 = v1665
											v873 = v1649
											v880 = v1656
											continue
										} else {
											break
										}
										break
									}
									v1676 = v1649
									v1683 = v1656
									v1691 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1676))) = uint16(v1691)
									*(*uint16)(unsafe.Add(mBase, uint32(v1683))) = uint16(v1691)
									*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v406
									*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v377
									return v29
								}
							}
						}
					}
				}
			}
		}
	}
}
