package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_ghstore_in_0), int32(102), int32(_a_F_ghstore_in_1), int32(_a_F_ghstore_in_2), int32(_a_F_ghstore_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ghstore_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v765 int32
	_ = v765
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v880 int32
	_ = v880
	var v887 int32
	_ = v887
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v915 int32
	_ = v915
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1093 int32
	_ = v1093
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1199 int32
	_ = v1199
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1256 int32
	_ = v1256
	var v1260 int32
	_ = v1260
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1287 int32
	_ = v1287
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1392 int32
	_ = v1392
	var v1404 int32
	_ = v1404
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1452 int32
	_ = v1452
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1541 int32
	_ = v1541
	var v1566 int32
	_ = v1566
	var v1578 int32
	_ = v1578
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1651 int32
	_ = v1651
	var v1666 int32
	_ = v1666
	var v1673 int32
	_ = v1673
	var v1683 int32
	_ = v1683
	var v1694 int32
	_ = v1694
	var v1701 int32
	_ = v1701
	var v1710 int32
	_ = v1710
	v2 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v31 = base.I32_wrap_i64(v30)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v33 == v2 {
		v50 = v2
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
		if v37 == int32(0) {
			v50 = v2
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			if v40 != int32(7) {
				v50 = v2
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
				if v43 != int32(17) {
					v50 = v2
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+32)))
					v50 = v46 ^ int32(1)
				}
			}
		}
	}
	if v50&int32(1) != 0 {
		v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v54 = F_get_fn_opclass_options(m, v53)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int64(0)
		} else {
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
			v59 = v58
			v61 = (v27 + int32(_a_F_ghstore_picksplit_0)) & int32(_a_F_ghstore_picksplit_1)
			v65 = v61<<(uint(int32(1))%32) + int32(4)
			v66 = F_palloc(m, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31))) = v66
				v69 = F_palloc(m, v65)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v69
					if base.Ui32(int32(2)) <= base.Ui32(v61) {
						v75 = v26 + int32(8)
						v79 = int32(-1)
						v80 = v2
						v87 = int32(1)
						v89 = v2
						for {
							v106 = *(*int32)(unsafe.Add(mBase, uint32(v75+v87*int32(24))))
							v108 = v87 + int32(1)
							v109 = v108
							v110 = v79
							v111 = v80
							v119 = v108
							v120 = v89
							for {
								v137 = *(*int32)(unsafe.Add(mBase, uint32(v75+v119*int32(24))))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
								v145 = int32(4)
								v146 = v144 & v145
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
								if v147&v145 != 0 {
									if v146 != 0 {
										v315 = int32(0)
									} else {
										v152 = v59 << (uint(int32(3)) % 32)
										if v59 <= int32(0) {
											v315 = v152
										} else {
											v251 = v137 + int32(8)
											v252 = v152
											v254 = v251
											v255 = int32(0)
											v258 = int32(0)
											for {
												v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
												v264 = int32(1)
												v299 = v255 + v263&v264 + int32(base.Ui32(v263)>>(uint(int32(7))%32)) + int32(base.Ui32(v263)>>(uint(v264)%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(2))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(3))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(4))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(5))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(6))%32))&v264
												v303 = v258 + v264
												if v303 != v59 {
													v254 = v254 + v264
													v255 = v299
													v258 = v303
													continue
												} else {
													break
												}
												break
											}
											v315 = v252 - v299
										}
									}
								} else {
									if v146 != 0 {
										v158 = v59 << (uint(int32(3)) % 32)
										if v59 <= int32(0) {
											v315 = v158
										} else {
											v251 = v106 + int32(8)
											v252 = v158
											v254 = v251
											v255 = int32(0)
											v258 = int32(0)
											for {
												v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
												v264 = int32(1)
												v299 = v255 + v263&v264 + int32(base.Ui32(v263)>>(uint(int32(7))%32)) + int32(base.Ui32(v263)>>(uint(v264)%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(2))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(3))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(4))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(5))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(6))%32))&v264
												v303 = v258 + v264
												if v303 != v59 {
													v254 = v254 + v264
													v255 = v299
													v258 = v303
													continue
												} else {
													break
												}
												break
											}
											v315 = v252 - v299
										}
									} else {
										if v59 <= int32(0) {
											v315 = int32(0)
										} else {
											v166 = int32(8)
											v167 = v137 + v166
											v169 = v106 + v166
											v170 = int32(0)
											v172 = int32(1)
											v174 = v59 << (uint(int32(3)) % 32)
											if v174 <= v172 {
												v177 = v172
											} else {
												v177 = v174
											}
											if v177 != int32(1) {
												v185 = v170
												v186 = v170
												v187 = int32(0)
												for {
													v195 = int32(base.Ui32(v186) >> (uint(int32(3)) % 32))
													v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v195))))
													v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v195))))
													v200 = v197 ^ v199
													v202 = v186 & int32(6)
													v203 = int32(1)
													v212 = int32(base.Ui32(v200)>>(uint(v202|v203)%32))&v203 + (int32(base.Ui32(v200)>>(uint(v202)%32))&v203 + v185)
													v213 = int32(2)
													v214 = v186 + v213
													v216 = v187 + v213
													if v216 != v177&int32(2147483640) {
														v185 = v212
														v186 = v214
														v187 = v216
														continue
													} else {
														break
													}
													break
												}
												if v177&int32(1) == int32(0) {
													v242 = v212
												} else {
													v220 = v212
													v221 = v214
													v230 = int32(base.Ui32(v221) >> (uint(int32(3)) % 32))
													v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v230))))
													v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v169))))
													v242 = v220 + int32(base.Ui32(v232^v234)>>(uint(v221&int32(7))%32))&int32(1)
												}
											} else {
												v220 = v170
												v221 = v170
												v230 = int32(base.Ui32(v221) >> (uint(int32(3)) % 32))
												v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v230))))
												v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v169))))
												v242 = v220 + int32(base.Ui32(v232^v234)>>(uint(v221&int32(7))%32))&int32(1)
											}
											v315 = v242
										}
									}
								}
								v316 = base.B2i32(v110 < v315)
								if v110 < v315 {
									v317 = v315
								} else {
									v317 = v110
								}
								if v110 < v315 {
									v318 = v109
								} else {
									v318 = v120
								}
								if v110 < v315 {
									v319 = v87
								} else {
									v319 = v111
								}
								v321 = v109 + int32(1)
								v323 = v321 & int32(_a_F_ghstore_picksplit_1)
								if base.Ui32(v323) <= base.Ui32(v61) {
									v109 = v321
									v110 = v317
									v111 = v319
									v119 = v323
									v120 = v318
									continue
								} else {
									break
								}
								break
							}
							if v61 != v108 {
								v79 = v317
								v80 = v319
								v87 = v108
								v89 = v318
								continue
							} else {
								break
							}
							break
						}
						v328 = v319
						v337 = v318
					} else {
						v328 = v2
						v337 = v2
					}
					v351 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v351
					*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v351
					v355 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v356 = int32(8)
					v358 = v59 + v356
					v360 = v26 + v356
					v362 = int32(_a_F_ghstore_picksplit_1)
					v370 = base.B2i32(v328&v362 == v351) | base.B2i32(v337&v362 == v351)
					if v370 != 0 {
						v371 = int32(1)
					} else {
						v371 = v328
					}
					v377 = *(*int32)(unsafe.Add(mBase, uint32(v360+v371&int32(_a_F_ghstore_picksplit_1)*int32(24))))
					v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
					v380 = v378 & int32(4)
					if v380 != 0 {
						v381 = v356
					} else {
						v381 = v358
					}
					v382 = F_palloc(m, v381)
					mBase = m.M
					v383 = m.ExcPending
					if v383 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v382)+4)) = v380
						*(*int32)(unsafe.Add(mBase, uint32(v382))) = v381 << (uint(int32(2)) % 32)
						v388 = int32(0)
						if v380|base.B2i32(v59 == v388) == v388 {
							v393 = int32(8)
							base.MemoryCopy(m, v382+v393, v377+v393, v59)
						} else {
						}
						if v370 != 0 {
							v400 = int32(2)
						} else {
							v400 = v337
						}
						v406 = *(*int32)(unsafe.Add(mBase, uint32(v360+v400&int32(_a_F_ghstore_picksplit_1)*int32(24))))
						v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
						v409 = v407 & int32(4)
						if v409 != 0 {
							v410 = int32(8)
						} else {
							v410 = v358
						}
						v411 = F_palloc(m, v410)
						mBase = m.M
						v412 = m.ExcPending
						if v412 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v411)+4)) = v409
							*(*int32)(unsafe.Add(mBase, uint32(v411))) = v410 << (uint(int32(2)) % 32)
							v417 = int32(0)
							if v409|base.B2i32(v59 == v417) == v417 {
								v422 = int32(8)
								base.MemoryCopy(m, v411+v422, v406+v422, v59)
							} else {
							}
							v430 = int32(_a_F_ghstore_picksplit_1)
							v431 = v27 + v430
							v433 = v431 & v430
							v434 = F_palloc_mul(m, int32(8), v433)
							mBase = m.M
							v435 = m.ExcPending
							if v435 != 0 {
								return int64(0)
							} else {
								if v27&int32(_a_F_ghstore_picksplit_1) == int32(1) {
									F_pg_qsort(m, v434, v433, int32(8), int32(_a_F_ghstore_picksplit_2))
									mBase = m.M
									v443 = m.ExcPending
									if v443 != 0 {
										return int64(0)
									} else {
										v1694 = v355
										v1701 = v69
										v1710 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1694))) = uint16(v1710)
										*(*uint16)(unsafe.Add(mBase, uint32(v1701))) = uint16(v1710)
										*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = base.I64_extend_i32_u(v411)
										*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = base.I64_extend_i32_u(v382)
										return v30 & int64(4294967295)
									}
								} else {
									v444 = int32(1)
									v446 = v444
									v447 = v444
									for {
										v473 = v434 + v446<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v473-int32(8)))) = uint16(v447)
										v477 = int32(4)
										v482 = *(*int32)(unsafe.Add(mBase, uint32(v360+v446*int32(24))))
										v489 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
										v491 = v489 & v477
										v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
										if v492&v477 != 0 {
											if v491 != 0 {
												v660 = int32(0)
											} else {
												v497 = v59 << (uint(int32(3)) % 32)
												if v59 <= int32(0) {
													v660 = v497
												} else {
													v596 = v482 + int32(8)
													v597 = v497
													v599 = v596
													v600 = int32(0)
													v603 = int32(0)
													for {
														v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
														v609 = int32(1)
														v644 = v600 + v608&v609 + int32(base.Ui32(v608)>>(uint(int32(7))%32)) + int32(base.Ui32(v608)>>(uint(v609)%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(2))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(3))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(4))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(5))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(6))%32))&v609
														v648 = v603 + v609
														if v648 != v59 {
															v599 = v599 + v609
															v600 = v644
															v603 = v648
															continue
														} else {
															break
														}
														break
													}
													v660 = v597 - v644
												}
											}
										} else {
											if v491 != 0 {
												v503 = v59 << (uint(int32(3)) % 32)
												if v59 <= int32(0) {
													v660 = v503
												} else {
													v596 = v382 + int32(8)
													v597 = v503
													v599 = v596
													v600 = int32(0)
													v603 = int32(0)
													for {
														v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
														v609 = int32(1)
														v644 = v600 + v608&v609 + int32(base.Ui32(v608)>>(uint(int32(7))%32)) + int32(base.Ui32(v608)>>(uint(v609)%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(2))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(3))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(4))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(5))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(6))%32))&v609
														v648 = v603 + v609
														if v648 != v59 {
															v599 = v599 + v609
															v600 = v644
															v603 = v648
															continue
														} else {
															break
														}
														break
													}
													v660 = v597 - v644
												}
											} else {
												if v59 <= int32(0) {
													v660 = int32(0)
												} else {
													v511 = int32(8)
													v512 = v482 + v511
													v514 = v382 + v511
													v515 = int32(0)
													v517 = int32(1)
													v519 = v59 << (uint(int32(3)) % 32)
													if v519 <= v517 {
														v522 = v517
													} else {
														v522 = v519
													}
													if v522 != int32(1) {
														v530 = v515
														v531 = v515
														v532 = int32(0)
														for {
															v540 = int32(base.Ui32(v531) >> (uint(int32(3)) % 32))
															v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v540))))
															v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514+v540))))
															v545 = v542 ^ v544
															v547 = v531 & int32(6)
															v548 = int32(1)
															v557 = int32(base.Ui32(v545)>>(uint(v547|v548)%32))&v548 + (int32(base.Ui32(v545)>>(uint(v547)%32))&v548 + v530)
															v558 = int32(2)
															v559 = v531 + v558
															v561 = v532 + v558
															if v561 != v522&int32(2147483640) {
																v530 = v557
																v531 = v559
																v532 = v561
																continue
															} else {
																break
															}
															break
														}
														if v522&int32(1) == int32(0) {
															v587 = v557
														} else {
															v565 = v557
															v566 = v559
															v575 = int32(base.Ui32(v566) >> (uint(int32(3)) % 32))
															v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v575))))
															v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v514))))
															v587 = v565 + int32(base.Ui32(v577^v579)>>(uint(v566&int32(7))%32))&int32(1)
														}
													} else {
														v565 = v515
														v566 = v515
														v575 = int32(base.Ui32(v566) >> (uint(int32(3)) % 32))
														v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v575))))
														v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v514))))
														v587 = v565 + int32(base.Ui32(v577^v579)>>(uint(v566&int32(7))%32))&int32(1)
													}
													v660 = v587
												}
											}
										}
										v667 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
										v668 = int32(4)
										v669 = v667 & v668
										v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
										if v670&v668 != 0 {
											if v669 != 0 {
												v838 = int32(0)
											} else {
												v675 = v59 << (uint(int32(3)) % 32)
												if v59 <= int32(0) {
													v838 = v675
												} else {
													v774 = v482 + int32(8)
													v775 = v675
													v777 = v774
													v778 = int32(0)
													v781 = int32(0)
													for {
														v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
														v787 = int32(1)
														v822 = v778 + v786&v787 + int32(base.Ui32(v786)>>(uint(int32(7))%32)) + int32(base.Ui32(v786)>>(uint(v787)%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(2))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(3))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(4))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(5))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(6))%32))&v787
														v826 = v781 + v787
														if v826 != v59 {
															v777 = v777 + v787
															v778 = v822
															v781 = v826
															continue
														} else {
															break
														}
														break
													}
													v838 = v775 - v822
												}
											}
										} else {
											if v669 != 0 {
												v681 = v59 << (uint(int32(3)) % 32)
												if v59 <= int32(0) {
													v838 = v681
												} else {
													v774 = v411 + int32(8)
													v775 = v681
													v777 = v774
													v778 = int32(0)
													v781 = int32(0)
													for {
														v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
														v787 = int32(1)
														v822 = v778 + v786&v787 + int32(base.Ui32(v786)>>(uint(int32(7))%32)) + int32(base.Ui32(v786)>>(uint(v787)%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(2))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(3))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(4))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(5))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(6))%32))&v787
														v826 = v781 + v787
														if v826 != v59 {
															v777 = v777 + v787
															v778 = v822
															v781 = v826
															continue
														} else {
															break
														}
														break
													}
													v838 = v775 - v822
												}
											} else {
												if v59 <= int32(0) {
													v838 = int32(0)
												} else {
													v689 = int32(8)
													v690 = v482 + v689
													v692 = v411 + v689
													v693 = int32(0)
													v695 = int32(1)
													v697 = v59 << (uint(int32(3)) % 32)
													if v697 <= v695 {
														v700 = v695
													} else {
														v700 = v697
													}
													if v700 != int32(1) {
														v708 = v693
														v709 = v693
														v710 = int32(0)
														for {
															v718 = int32(base.Ui32(v709) >> (uint(int32(3)) % 32))
															v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v718))))
															v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v718))))
															v723 = v720 ^ v722
															v725 = v709 & int32(6)
															v726 = int32(1)
															v735 = int32(base.Ui32(v723)>>(uint(v725|v726)%32))&v726 + (int32(base.Ui32(v723)>>(uint(v725)%32))&v726 + v708)
															v736 = int32(2)
															v737 = v709 + v736
															v739 = v710 + v736
															if v739 != v700&int32(2147483640) {
																v708 = v735
																v709 = v737
																v710 = v739
																continue
															} else {
																break
															}
															break
														}
														if v700&int32(1) == int32(0) {
															v765 = v735
														} else {
															v743 = v735
															v744 = v737
															v753 = int32(base.Ui32(v744) >> (uint(int32(3)) % 32))
															v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v753))))
															v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v692))))
															v765 = v743 + int32(base.Ui32(v755^v757)>>(uint(v744&int32(7))%32))&int32(1)
														}
													} else {
														v743 = v693
														v744 = v693
														v753 = int32(base.Ui32(v744) >> (uint(int32(3)) % 32))
														v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v753))))
														v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v692))))
														v765 = v743 + int32(base.Ui32(v755^v757)>>(uint(v744&int32(7))%32))&int32(1)
													}
													v838 = v765
												}
											}
										}
										v839 = v660 - v838
										v841 = v839 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v473-v477))) = v839 ^ v841 - v841
										v846 = v447 + int32(1)
										v847 = int32(_a_F_ghstore_picksplit_1)
										v848 = v846 & v847
										if base.Ui32(v848) <= base.Ui32(v431&v847) {
											v446 = v848
											v447 = v846
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v434, v433, int32(8), int32(_a_F_ghstore_picksplit_2))
									mBase = m.M
									v855 = m.ExcPending
									if v855 != 0 {
										return int64(0)
									} else {
										v856 = int32(1)
										if base.Ui32(v433) <= base.Ui32(v856) {
											v859 = v856
										} else {
											v859 = v433
										}
										v861 = v59 & int32(2147483644)
										v863 = v59 & int32(3)
										v864 = int32(8)
										v865 = v411 + v864
										v867 = v382 + v864
										v871 = int32(0)
										v880 = v355
										v887 = v69
										for {
											v899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434+v871<<(uint(int32(3))%32)))))
											if v371&int32(_a_F_ghstore_picksplit_1) == v899 {
												*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v371)
												v902 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v902 + int32(1)
												v1666 = v880 + int32(2)
												v1673 = v887
											} else {
												if v400&int32(_a_F_ghstore_picksplit_1) == v899 {
													*(*uint16)(unsafe.Add(mBase, uint32(v887))) = uint16(v400)
													v1651 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1651 + int32(1)
													v1666 = v880
													v1673 = v887 + int32(2)
												} else {
													v915 = *(*int32)(unsafe.Add(mBase, uint32(v360+v899*int32(24))))
													v922 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
													v923 = int32(4)
													v924 = v922 & v923
													v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
													if v925&v923 != 0 {
														if v924 != 0 {
															v1093 = int32(0)
														} else {
															v930 = v59 << (uint(int32(3)) % 32)
															if v59 <= int32(0) {
																v1093 = v930
															} else {
																v1029 = v915 + int32(8)
																v1030 = v930
																v1032 = v1029
																v1033 = int32(0)
																v1036 = int32(0)
																for {
																	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
																	v1042 = int32(1)
																	v1077 = v1033 + v1041&v1042 + int32(base.Ui32(v1041)>>(uint(int32(7))%32)) + int32(base.Ui32(v1041)>>(uint(v1042)%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(2))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(3))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(4))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(5))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(6))%32))&v1042
																	v1081 = v1036 + v1042
																	if v1081 != v59 {
																		v1032 = v1032 + v1042
																		v1033 = v1077
																		v1036 = v1081
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1093 = v1030 - v1077
															}
														}
													} else {
														if v924 != 0 {
															v936 = v59 << (uint(int32(3)) % 32)
															if v59 <= int32(0) {
																v1093 = v936
															} else {
																v1029 = v382 + int32(8)
																v1030 = v936
																v1032 = v1029
																v1033 = int32(0)
																v1036 = int32(0)
																for {
																	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
																	v1042 = int32(1)
																	v1077 = v1033 + v1041&v1042 + int32(base.Ui32(v1041)>>(uint(int32(7))%32)) + int32(base.Ui32(v1041)>>(uint(v1042)%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(2))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(3))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(4))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(5))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(6))%32))&v1042
																	v1081 = v1036 + v1042
																	if v1081 != v59 {
																		v1032 = v1032 + v1042
																		v1033 = v1077
																		v1036 = v1081
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1093 = v1030 - v1077
															}
														} else {
															if v59 <= int32(0) {
																v1093 = int32(0)
															} else {
																v944 = int32(8)
																v945 = v915 + v944
																v947 = v382 + v944
																v948 = int32(0)
																v950 = int32(1)
																v952 = v59 << (uint(int32(3)) % 32)
																if v952 <= v950 {
																	v955 = v950
																} else {
																	v955 = v952
																}
																if v955 != int32(1) {
																	v963 = v948
																	v964 = v948
																	v965 = int32(0)
																	for {
																		v973 = int32(base.Ui32(v964) >> (uint(int32(3)) % 32))
																		v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v973))))
																		v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947+v973))))
																		v978 = v975 ^ v977
																		v980 = v964 & int32(6)
																		v981 = int32(1)
																		v990 = int32(base.Ui32(v978)>>(uint(v980|v981)%32))&v981 + (int32(base.Ui32(v978)>>(uint(v980)%32))&v981 + v963)
																		v991 = int32(2)
																		v992 = v964 + v991
																		v994 = v965 + v991
																		if v994 != v955&int32(2147483640) {
																			v963 = v990
																			v964 = v992
																			v965 = v994
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v955&int32(1) == int32(0) {
																		v1020 = v990
																	} else {
																		v998 = v990
																		v999 = v992
																		v1008 = int32(base.Ui32(v999) >> (uint(int32(3)) % 32))
																		v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v1008))))
																		v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008+v947))))
																		v1020 = v998 + int32(base.Ui32(v1010^v1012)>>(uint(v999&int32(7))%32))&int32(1)
																	}
																} else {
																	v998 = v948
																	v999 = v948
																	v1008 = int32(base.Ui32(v999) >> (uint(int32(3)) % 32))
																	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v1008))))
																	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008+v947))))
																	v1020 = v998 + int32(base.Ui32(v1010^v1012)>>(uint(v999&int32(7))%32))&int32(1)
																}
																v1093 = v1020
															}
														}
													}
													v1101 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
													v1102 = int32(4)
													v1103 = v1101 & v1102
													v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
													if v1104&v1102 != 0 {
														if v1103 != 0 {
															v1272 = int32(0)
														} else {
															v1109 = v59 << (uint(int32(3)) % 32)
															if v59 <= int32(0) {
																v1272 = v1109
															} else {
																v1208 = v915 + int32(8)
																v1209 = v1109
																v1211 = v1208
																v1212 = int32(0)
																v1215 = int32(0)
																for {
																	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
																	v1221 = int32(1)
																	v1256 = v1212 + v1220&v1221 + int32(base.Ui32(v1220)>>(uint(int32(7))%32)) + int32(base.Ui32(v1220)>>(uint(v1221)%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(2))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(3))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(4))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(5))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(6))%32))&v1221
																	v1260 = v1215 + v1221
																	if v1260 != v59 {
																		v1211 = v1211 + v1221
																		v1212 = v1256
																		v1215 = v1260
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1272 = v1209 - v1256
															}
														}
													} else {
														if v1103 != 0 {
															v1115 = v59 << (uint(int32(3)) % 32)
															if v59 <= int32(0) {
																v1272 = v1115
															} else {
																v1208 = v411 + int32(8)
																v1209 = v1115
																v1211 = v1208
																v1212 = int32(0)
																v1215 = int32(0)
																for {
																	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
																	v1221 = int32(1)
																	v1256 = v1212 + v1220&v1221 + int32(base.Ui32(v1220)>>(uint(int32(7))%32)) + int32(base.Ui32(v1220)>>(uint(v1221)%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(2))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(3))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(4))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(5))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(6))%32))&v1221
																	v1260 = v1215 + v1221
																	if v1260 != v59 {
																		v1211 = v1211 + v1221
																		v1212 = v1256
																		v1215 = v1260
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1272 = v1209 - v1256
															}
														} else {
															if v59 <= int32(0) {
																v1272 = int32(0)
															} else {
																v1123 = int32(8)
																v1124 = v915 + v1123
																v1126 = v411 + v1123
																v1127 = int32(0)
																v1129 = int32(1)
																v1131 = v59 << (uint(int32(3)) % 32)
																if v1131 <= v1129 {
																	v1134 = v1129
																} else {
																	v1134 = v1131
																}
																if v1134 != int32(1) {
																	v1142 = v1127
																	v1143 = v1127
																	v1144 = int32(0)
																	for {
																		v1152 = int32(base.Ui32(v1143) >> (uint(int32(3)) % 32))
																		v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1152))))
																		v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126+v1152))))
																		v1157 = v1154 ^ v1156
																		v1159 = v1143 & int32(6)
																		v1160 = int32(1)
																		v1169 = int32(base.Ui32(v1157)>>(uint(v1159|v1160)%32))&v1160 + (int32(base.Ui32(v1157)>>(uint(v1159)%32))&v1160 + v1142)
																		v1170 = int32(2)
																		v1171 = v1143 + v1170
																		v1173 = v1144 + v1170
																		if v1173 != v1134&int32(2147483640) {
																			v1142 = v1169
																			v1143 = v1171
																			v1144 = v1173
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v1134&int32(1) == int32(0) {
																		v1199 = v1169
																	} else {
																		v1177 = v1169
																		v1178 = v1171
																		v1187 = int32(base.Ui32(v1178) >> (uint(int32(3)) % 32))
																		v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1187))))
																		v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v1126))))
																		v1199 = v1177 + int32(base.Ui32(v1189^v1191)>>(uint(v1178&int32(7))%32))&int32(1)
																	}
																} else {
																	v1177 = v1127
																	v1178 = v1127
																	v1187 = int32(base.Ui32(v1178) >> (uint(int32(3)) % 32))
																	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1187))))
																	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v1126))))
																	v1199 = v1177 + int32(base.Ui32(v1189^v1191)>>(uint(v1178&int32(7))%32))&int32(1)
																}
																v1272 = v1199
															}
														}
													}
													v1274 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
													v1275 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
													v1276 = v1274 - v1275
													if base.F64_lt(base.F64_convert_i32_s(v1093), base.F64_add(base.F64_convert_i32_s(v1272), base.F64_mul(base.F64_convert_i32_s(v1276*v1276*v1276), float64(-0.0001)))) != 0 {
														v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
														if v1284&int32(4) != 0 {
														} else {
															v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+4)))
															if v1287&int32(4) != 0 {
																if v59 == int32(0) {
																} else {
																	base.MemoryFill(m, v867, int32(255), v59)
																}
															} else {
																if v59 <= int32(0) {
																} else {
																	v1297 = v915 + int32(8)
																	v1298 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v59) {
																		v1304 = v1298
																		v1305 = v1298
																		for {
																			v1328 = v1304 + v867
																			v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328))))
																			v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304+v1297))))
																			v1332 = v1329 | v1331
																			*(*uint8)(unsafe.Add(mBase, uint32(v1328))) = uint8(v1332)
																			v1335 = v1304 | int32(1)
																			v1336 = v867 + v1335
																			v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336))))
																			v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335+v1297))))
																			v1340 = v1337 | v1339
																			*(*uint8)(unsafe.Add(mBase, uint32(v1336))) = uint8(v1340)
																			v1343 = v1304 | int32(2)
																			v1344 = v867 + v1343
																			v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344))))
																			v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343+v1297))))
																			v1348 = v1345 | v1347
																			*(*uint8)(unsafe.Add(mBase, uint32(v1344))) = uint8(v1348)
																			v1351 = v1304 | int32(3)
																			v1352 = v867 + v1351
																			v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352))))
																			v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1297))))
																			v1356 = v1353 | v1355
																			*(*uint8)(unsafe.Add(mBase, uint32(v1352))) = uint8(v1356)
																			v1358 = int32(4)
																			v1359 = v1304 + v1358
																			v1361 = v1305 + v1358
																			if v1361 != v861 {
																				v1304 = v1359
																				v1305 = v1361
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v863 == int32(0) {
																		} else {
																			v1367 = v1359
																			v1392 = v1367
																			v1404 = v1298
																			for {
																				v1415 = v1392 + v867
																				v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415))))
																				v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1297))))
																				v1419 = v1416 | v1418
																				*(*uint8)(unsafe.Add(mBase, uint32(v1415))) = uint8(v1419)
																				v1421 = int32(1)
																				v1424 = v1404 + v1421
																				if v1424 != v863 {
																					v1392 = v1392 + v1421
																					v1404 = v1424
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v1367 = v1298
																		v1392 = v1367
																		v1404 = v1298
																		for {
																			v1415 = v1392 + v867
																			v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415))))
																			v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1297))))
																			v1419 = v1416 | v1418
																			*(*uint8)(unsafe.Add(mBase, uint32(v1415))) = uint8(v1419)
																			v1421 = int32(1)
																			v1424 = v1404 + v1421
																			if v1424 != v863 {
																				v1392 = v1392 + v1421
																				v1404 = v1424
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
														*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v899)
														v1452 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1452 + int32(1)
														v1666 = v880 + int32(2)
														v1673 = v887
													} else {
														v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
														if v1458&int32(4) != 0 {
														} else {
															v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+4)))
															if v1461&int32(4) != 0 {
																if v59 == int32(0) {
																} else {
																	base.MemoryFill(m, v865, int32(255), v59)
																}
															} else {
																if v59 <= int32(0) {
																} else {
																	v1471 = v915 + int32(8)
																	v1472 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v59) {
																		v1478 = v1472
																		v1479 = v1472
																		for {
																			v1502 = v1478 + v865
																			v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502))))
																			v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478+v1471))))
																			v1506 = v1503 | v1505
																			*(*uint8)(unsafe.Add(mBase, uint32(v1502))) = uint8(v1506)
																			v1509 = v1478 | int32(1)
																			v1510 = v865 + v1509
																			v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510))))
																			v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509+v1471))))
																			v1514 = v1511 | v1513
																			*(*uint8)(unsafe.Add(mBase, uint32(v1510))) = uint8(v1514)
																			v1517 = v1478 | int32(2)
																			v1518 = v865 + v1517
																			v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518))))
																			v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1517+v1471))))
																			v1522 = v1519 | v1521
																			*(*uint8)(unsafe.Add(mBase, uint32(v1518))) = uint8(v1522)
																			v1525 = v1478 | int32(3)
																			v1526 = v865 + v1525
																			v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526))))
																			v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525+v1471))))
																			v1530 = v1527 | v1529
																			*(*uint8)(unsafe.Add(mBase, uint32(v1526))) = uint8(v1530)
																			v1532 = int32(4)
																			v1533 = v1478 + v1532
																			v1535 = v1479 + v1532
																			if v1535 != v861 {
																				v1478 = v1533
																				v1479 = v1535
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v863 == int32(0) {
																		} else {
																			v1541 = v1533
																			v1566 = v1541
																			v1578 = v1472
																			for {
																				v1589 = v1566 + v865
																				v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
																				v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566+v1471))))
																				v1593 = v1590 | v1592
																				*(*uint8)(unsafe.Add(mBase, uint32(v1589))) = uint8(v1593)
																				v1595 = int32(1)
																				v1598 = v1578 + v1595
																				if v1598 != v863 {
																					v1566 = v1566 + v1595
																					v1578 = v1598
																					continue
																				} else {
																					break
																				}
																				break
																			}
																		}
																	} else {
																		v1541 = v1472
																		v1566 = v1541
																		v1578 = v1472
																		for {
																			v1589 = v1566 + v865
																			v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
																			v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566+v1471))))
																			v1593 = v1590 | v1592
																			*(*uint8)(unsafe.Add(mBase, uint32(v1589))) = uint8(v1593)
																			v1595 = int32(1)
																			v1598 = v1578 + v1595
																			if v1598 != v863 {
																				v1566 = v1566 + v1595
																				v1578 = v1598
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
														*(*uint16)(unsafe.Add(mBase, uint32(v887))) = uint16(v899)
														v1651 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1651 + int32(1)
														v1666 = v880
														v1673 = v887 + int32(2)
													}
												}
											}
											v1683 = v871 + int32(1)
											if v1683 != v859 {
												v871 = v1683
												v880 = v1666
												v887 = v1673
												continue
											} else {
												break
											}
											break
										}
										v1694 = v1666
										v1701 = v1673
										v1710 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1694))) = uint16(v1710)
										*(*uint16)(unsafe.Add(mBase, uint32(v1701))) = uint16(v1710)
										*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = base.I64_extend_i32_u(v411)
										*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = base.I64_extend_i32_u(v382)
										return v30 & int64(4294967295)
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v59 = int32(16)
		v61 = (v27 + int32(_a_F_ghstore_picksplit_0)) & int32(_a_F_ghstore_picksplit_1)
		v65 = v61<<(uint(int32(1))%32) + int32(4)
		v66 = F_palloc(m, v65)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v66
			v69 = F_palloc(m, v65)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v69
				if base.Ui32(int32(2)) <= base.Ui32(v61) {
					v75 = v26 + int32(8)
					v79 = int32(-1)
					v80 = v2
					v87 = int32(1)
					v89 = v2
					for {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v75+v87*int32(24))))
						v108 = v87 + int32(1)
						v109 = v108
						v110 = v79
						v111 = v80
						v119 = v108
						v120 = v89
						for {
							v137 = *(*int32)(unsafe.Add(mBase, uint32(v75+v119*int32(24))))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
							v145 = int32(4)
							v146 = v144 & v145
							v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
							if v147&v145 != 0 {
								if v146 != 0 {
									v315 = int32(0)
								} else {
									v152 = v59 << (uint(int32(3)) % 32)
									if v59 <= int32(0) {
										v315 = v152
									} else {
										v251 = v137 + int32(8)
										v252 = v152
										v254 = v251
										v255 = int32(0)
										v258 = int32(0)
										for {
											v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
											v264 = int32(1)
											v299 = v255 + v263&v264 + int32(base.Ui32(v263)>>(uint(int32(7))%32)) + int32(base.Ui32(v263)>>(uint(v264)%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(2))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(3))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(4))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(5))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(6))%32))&v264
											v303 = v258 + v264
											if v303 != v59 {
												v254 = v254 + v264
												v255 = v299
												v258 = v303
												continue
											} else {
												break
											}
											break
										}
										v315 = v252 - v299
									}
								}
							} else {
								if v146 != 0 {
									v158 = v59 << (uint(int32(3)) % 32)
									if v59 <= int32(0) {
										v315 = v158
									} else {
										v251 = v106 + int32(8)
										v252 = v158
										v254 = v251
										v255 = int32(0)
										v258 = int32(0)
										for {
											v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
											v264 = int32(1)
											v299 = v255 + v263&v264 + int32(base.Ui32(v263)>>(uint(int32(7))%32)) + int32(base.Ui32(v263)>>(uint(v264)%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(2))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(3))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(4))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(5))%32))&v264 + int32(base.Ui32(v263)>>(uint(int32(6))%32))&v264
											v303 = v258 + v264
											if v303 != v59 {
												v254 = v254 + v264
												v255 = v299
												v258 = v303
												continue
											} else {
												break
											}
											break
										}
										v315 = v252 - v299
									}
								} else {
									if v59 <= int32(0) {
										v315 = int32(0)
									} else {
										v166 = int32(8)
										v167 = v137 + v166
										v169 = v106 + v166
										v170 = int32(0)
										v172 = int32(1)
										v174 = v59 << (uint(int32(3)) % 32)
										if v174 <= v172 {
											v177 = v172
										} else {
											v177 = v174
										}
										if v177 != int32(1) {
											v185 = v170
											v186 = v170
											v187 = int32(0)
											for {
												v195 = int32(base.Ui32(v186) >> (uint(int32(3)) % 32))
												v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v195))))
												v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v195))))
												v200 = v197 ^ v199
												v202 = v186 & int32(6)
												v203 = int32(1)
												v212 = int32(base.Ui32(v200)>>(uint(v202|v203)%32))&v203 + (int32(base.Ui32(v200)>>(uint(v202)%32))&v203 + v185)
												v213 = int32(2)
												v214 = v186 + v213
												v216 = v187 + v213
												if v216 != v177&int32(2147483640) {
													v185 = v212
													v186 = v214
													v187 = v216
													continue
												} else {
													break
												}
												break
											}
											if v177&int32(1) == int32(0) {
												v242 = v212
											} else {
												v220 = v212
												v221 = v214
												v230 = int32(base.Ui32(v221) >> (uint(int32(3)) % 32))
												v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v230))))
												v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v169))))
												v242 = v220 + int32(base.Ui32(v232^v234)>>(uint(v221&int32(7))%32))&int32(1)
											}
										} else {
											v220 = v170
											v221 = v170
											v230 = int32(base.Ui32(v221) >> (uint(int32(3)) % 32))
											v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167+v230))))
											v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230+v169))))
											v242 = v220 + int32(base.Ui32(v232^v234)>>(uint(v221&int32(7))%32))&int32(1)
										}
										v315 = v242
									}
								}
							}
							v316 = base.B2i32(v110 < v315)
							if v110 < v315 {
								v317 = v315
							} else {
								v317 = v110
							}
							if v110 < v315 {
								v318 = v109
							} else {
								v318 = v120
							}
							if v110 < v315 {
								v319 = v87
							} else {
								v319 = v111
							}
							v321 = v109 + int32(1)
							v323 = v321 & int32(_a_F_ghstore_picksplit_1)
							if base.Ui32(v323) <= base.Ui32(v61) {
								v109 = v321
								v110 = v317
								v111 = v319
								v119 = v323
								v120 = v318
								continue
							} else {
								break
							}
							break
						}
						if v61 != v108 {
							v79 = v317
							v80 = v319
							v87 = v108
							v89 = v318
							continue
						} else {
							break
						}
						break
					}
					v328 = v319
					v337 = v318
				} else {
					v328 = v2
					v337 = v2
				}
				v351 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v351
				*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v351
				v355 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v356 = int32(8)
				v358 = v59 + v356
				v360 = v26 + v356
				v362 = int32(_a_F_ghstore_picksplit_1)
				v370 = base.B2i32(v328&v362 == v351) | base.B2i32(v337&v362 == v351)
				if v370 != 0 {
					v371 = int32(1)
				} else {
					v371 = v328
				}
				v377 = *(*int32)(unsafe.Add(mBase, uint32(v360+v371&int32(_a_F_ghstore_picksplit_1)*int32(24))))
				v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+4))
				v380 = v378 & int32(4)
				if v380 != 0 {
					v381 = v356
				} else {
					v381 = v358
				}
				v382 = F_palloc(m, v381)
				mBase = m.M
				v383 = m.ExcPending
				if v383 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v382)+4)) = v380
					*(*int32)(unsafe.Add(mBase, uint32(v382))) = v381 << (uint(int32(2)) % 32)
					v388 = int32(0)
					if v380|base.B2i32(v59 == v388) == v388 {
						v393 = int32(8)
						base.MemoryCopy(m, v382+v393, v377+v393, v59)
					} else {
					}
					if v370 != 0 {
						v400 = int32(2)
					} else {
						v400 = v337
					}
					v406 = *(*int32)(unsafe.Add(mBase, uint32(v360+v400&int32(_a_F_ghstore_picksplit_1)*int32(24))))
					v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
					v409 = v407 & int32(4)
					if v409 != 0 {
						v410 = int32(8)
					} else {
						v410 = v358
					}
					v411 = F_palloc(m, v410)
					mBase = m.M
					v412 = m.ExcPending
					if v412 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v411)+4)) = v409
						*(*int32)(unsafe.Add(mBase, uint32(v411))) = v410 << (uint(int32(2)) % 32)
						v417 = int32(0)
						if v409|base.B2i32(v59 == v417) == v417 {
							v422 = int32(8)
							base.MemoryCopy(m, v411+v422, v406+v422, v59)
						} else {
						}
						v430 = int32(_a_F_ghstore_picksplit_1)
						v431 = v27 + v430
						v433 = v431 & v430
						v434 = F_palloc_mul(m, int32(8), v433)
						mBase = m.M
						v435 = m.ExcPending
						if v435 != 0 {
							return int64(0)
						} else {
							if v27&int32(_a_F_ghstore_picksplit_1) == int32(1) {
								F_pg_qsort(m, v434, v433, int32(8), int32(_a_F_ghstore_picksplit_2))
								mBase = m.M
								v443 = m.ExcPending
								if v443 != 0 {
									return int64(0)
								} else {
									v1694 = v355
									v1701 = v69
									v1710 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1694))) = uint16(v1710)
									*(*uint16)(unsafe.Add(mBase, uint32(v1701))) = uint16(v1710)
									*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = base.I64_extend_i32_u(v411)
									*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = base.I64_extend_i32_u(v382)
									return v30 & int64(4294967295)
								}
							} else {
								v444 = int32(1)
								v446 = v444
								v447 = v444
								for {
									v473 = v434 + v446<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v473-int32(8)))) = uint16(v447)
									v477 = int32(4)
									v482 = *(*int32)(unsafe.Add(mBase, uint32(v360+v446*int32(24))))
									v489 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
									v491 = v489 & v477
									v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
									if v492&v477 != 0 {
										if v491 != 0 {
											v660 = int32(0)
										} else {
											v497 = v59 << (uint(int32(3)) % 32)
											if v59 <= int32(0) {
												v660 = v497
											} else {
												v596 = v482 + int32(8)
												v597 = v497
												v599 = v596
												v600 = int32(0)
												v603 = int32(0)
												for {
													v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
													v609 = int32(1)
													v644 = v600 + v608&v609 + int32(base.Ui32(v608)>>(uint(int32(7))%32)) + int32(base.Ui32(v608)>>(uint(v609)%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(2))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(3))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(4))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(5))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(6))%32))&v609
													v648 = v603 + v609
													if v648 != v59 {
														v599 = v599 + v609
														v600 = v644
														v603 = v648
														continue
													} else {
														break
													}
													break
												}
												v660 = v597 - v644
											}
										}
									} else {
										if v491 != 0 {
											v503 = v59 << (uint(int32(3)) % 32)
											if v59 <= int32(0) {
												v660 = v503
											} else {
												v596 = v382 + int32(8)
												v597 = v503
												v599 = v596
												v600 = int32(0)
												v603 = int32(0)
												for {
													v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
													v609 = int32(1)
													v644 = v600 + v608&v609 + int32(base.Ui32(v608)>>(uint(int32(7))%32)) + int32(base.Ui32(v608)>>(uint(v609)%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(2))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(3))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(4))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(5))%32))&v609 + int32(base.Ui32(v608)>>(uint(int32(6))%32))&v609
													v648 = v603 + v609
													if v648 != v59 {
														v599 = v599 + v609
														v600 = v644
														v603 = v648
														continue
													} else {
														break
													}
													break
												}
												v660 = v597 - v644
											}
										} else {
											if v59 <= int32(0) {
												v660 = int32(0)
											} else {
												v511 = int32(8)
												v512 = v482 + v511
												v514 = v382 + v511
												v515 = int32(0)
												v517 = int32(1)
												v519 = v59 << (uint(int32(3)) % 32)
												if v519 <= v517 {
													v522 = v517
												} else {
													v522 = v519
												}
												if v522 != int32(1) {
													v530 = v515
													v531 = v515
													v532 = int32(0)
													for {
														v540 = int32(base.Ui32(v531) >> (uint(int32(3)) % 32))
														v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v540))))
														v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514+v540))))
														v545 = v542 ^ v544
														v547 = v531 & int32(6)
														v548 = int32(1)
														v557 = int32(base.Ui32(v545)>>(uint(v547|v548)%32))&v548 + (int32(base.Ui32(v545)>>(uint(v547)%32))&v548 + v530)
														v558 = int32(2)
														v559 = v531 + v558
														v561 = v532 + v558
														if v561 != v522&int32(2147483640) {
															v530 = v557
															v531 = v559
															v532 = v561
															continue
														} else {
															break
														}
														break
													}
													if v522&int32(1) == int32(0) {
														v587 = v557
													} else {
														v565 = v557
														v566 = v559
														v575 = int32(base.Ui32(v566) >> (uint(int32(3)) % 32))
														v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v575))))
														v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v514))))
														v587 = v565 + int32(base.Ui32(v577^v579)>>(uint(v566&int32(7))%32))&int32(1)
													}
												} else {
													v565 = v515
													v566 = v515
													v575 = int32(base.Ui32(v566) >> (uint(int32(3)) % 32))
													v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512+v575))))
													v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575+v514))))
													v587 = v565 + int32(base.Ui32(v577^v579)>>(uint(v566&int32(7))%32))&int32(1)
												}
												v660 = v587
											}
										}
									}
									v667 = *(*int32)(unsafe.Add(mBase, uint32(v482)+4))
									v668 = int32(4)
									v669 = v667 & v668
									v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
									if v670&v668 != 0 {
										if v669 != 0 {
											v838 = int32(0)
										} else {
											v675 = v59 << (uint(int32(3)) % 32)
											if v59 <= int32(0) {
												v838 = v675
											} else {
												v774 = v482 + int32(8)
												v775 = v675
												v777 = v774
												v778 = int32(0)
												v781 = int32(0)
												for {
													v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
													v787 = int32(1)
													v822 = v778 + v786&v787 + int32(base.Ui32(v786)>>(uint(int32(7))%32)) + int32(base.Ui32(v786)>>(uint(v787)%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(2))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(3))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(4))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(5))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(6))%32))&v787
													v826 = v781 + v787
													if v826 != v59 {
														v777 = v777 + v787
														v778 = v822
														v781 = v826
														continue
													} else {
														break
													}
													break
												}
												v838 = v775 - v822
											}
										}
									} else {
										if v669 != 0 {
											v681 = v59 << (uint(int32(3)) % 32)
											if v59 <= int32(0) {
												v838 = v681
											} else {
												v774 = v411 + int32(8)
												v775 = v681
												v777 = v774
												v778 = int32(0)
												v781 = int32(0)
												for {
													v786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v777))))
													v787 = int32(1)
													v822 = v778 + v786&v787 + int32(base.Ui32(v786)>>(uint(int32(7))%32)) + int32(base.Ui32(v786)>>(uint(v787)%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(2))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(3))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(4))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(5))%32))&v787 + int32(base.Ui32(v786)>>(uint(int32(6))%32))&v787
													v826 = v781 + v787
													if v826 != v59 {
														v777 = v777 + v787
														v778 = v822
														v781 = v826
														continue
													} else {
														break
													}
													break
												}
												v838 = v775 - v822
											}
										} else {
											if v59 <= int32(0) {
												v838 = int32(0)
											} else {
												v689 = int32(8)
												v690 = v482 + v689
												v692 = v411 + v689
												v693 = int32(0)
												v695 = int32(1)
												v697 = v59 << (uint(int32(3)) % 32)
												if v697 <= v695 {
													v700 = v695
												} else {
													v700 = v697
												}
												if v700 != int32(1) {
													v708 = v693
													v709 = v693
													v710 = int32(0)
													for {
														v718 = int32(base.Ui32(v709) >> (uint(int32(3)) % 32))
														v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v718))))
														v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v718))))
														v723 = v720 ^ v722
														v725 = v709 & int32(6)
														v726 = int32(1)
														v735 = int32(base.Ui32(v723)>>(uint(v725|v726)%32))&v726 + (int32(base.Ui32(v723)>>(uint(v725)%32))&v726 + v708)
														v736 = int32(2)
														v737 = v709 + v736
														v739 = v710 + v736
														if v739 != v700&int32(2147483640) {
															v708 = v735
															v709 = v737
															v710 = v739
															continue
														} else {
															break
														}
														break
													}
													if v700&int32(1) == int32(0) {
														v765 = v735
													} else {
														v743 = v735
														v744 = v737
														v753 = int32(base.Ui32(v744) >> (uint(int32(3)) % 32))
														v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v753))))
														v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v692))))
														v765 = v743 + int32(base.Ui32(v755^v757)>>(uint(v744&int32(7))%32))&int32(1)
													}
												} else {
													v743 = v693
													v744 = v693
													v753 = int32(base.Ui32(v744) >> (uint(int32(3)) % 32))
													v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v690+v753))))
													v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v692))))
													v765 = v743 + int32(base.Ui32(v755^v757)>>(uint(v744&int32(7))%32))&int32(1)
												}
												v838 = v765
											}
										}
									}
									v839 = v660 - v838
									v841 = v839 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v473-v477))) = v839 ^ v841 - v841
									v846 = v447 + int32(1)
									v847 = int32(_a_F_ghstore_picksplit_1)
									v848 = v846 & v847
									if base.Ui32(v848) <= base.Ui32(v431&v847) {
										v446 = v848
										v447 = v846
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v434, v433, int32(8), int32(_a_F_ghstore_picksplit_2))
								mBase = m.M
								v855 = m.ExcPending
								if v855 != 0 {
									return int64(0)
								} else {
									v856 = int32(1)
									if base.Ui32(v433) <= base.Ui32(v856) {
										v859 = v856
									} else {
										v859 = v433
									}
									v861 = v59 & int32(2147483644)
									v863 = v59 & int32(3)
									v864 = int32(8)
									v865 = v411 + v864
									v867 = v382 + v864
									v871 = int32(0)
									v880 = v355
									v887 = v69
									for {
										v899 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434+v871<<(uint(int32(3))%32)))))
										if v371&int32(_a_F_ghstore_picksplit_1) == v899 {
											*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v371)
											v902 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v902 + int32(1)
											v1666 = v880 + int32(2)
											v1673 = v887
										} else {
											if v400&int32(_a_F_ghstore_picksplit_1) == v899 {
												*(*uint16)(unsafe.Add(mBase, uint32(v887))) = uint16(v400)
												v1651 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1651 + int32(1)
												v1666 = v880
												v1673 = v887 + int32(2)
											} else {
												v915 = *(*int32)(unsafe.Add(mBase, uint32(v360+v899*int32(24))))
												v922 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
												v923 = int32(4)
												v924 = v922 & v923
												v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
												if v925&v923 != 0 {
													if v924 != 0 {
														v1093 = int32(0)
													} else {
														v930 = v59 << (uint(int32(3)) % 32)
														if v59 <= int32(0) {
															v1093 = v930
														} else {
															v1029 = v915 + int32(8)
															v1030 = v930
															v1032 = v1029
															v1033 = int32(0)
															v1036 = int32(0)
															for {
																v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
																v1042 = int32(1)
																v1077 = v1033 + v1041&v1042 + int32(base.Ui32(v1041)>>(uint(int32(7))%32)) + int32(base.Ui32(v1041)>>(uint(v1042)%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(2))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(3))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(4))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(5))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(6))%32))&v1042
																v1081 = v1036 + v1042
																if v1081 != v59 {
																	v1032 = v1032 + v1042
																	v1033 = v1077
																	v1036 = v1081
																	continue
																} else {
																	break
																}
																break
															}
															v1093 = v1030 - v1077
														}
													}
												} else {
													if v924 != 0 {
														v936 = v59 << (uint(int32(3)) % 32)
														if v59 <= int32(0) {
															v1093 = v936
														} else {
															v1029 = v382 + int32(8)
															v1030 = v936
															v1032 = v1029
															v1033 = int32(0)
															v1036 = int32(0)
															for {
																v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1032))))
																v1042 = int32(1)
																v1077 = v1033 + v1041&v1042 + int32(base.Ui32(v1041)>>(uint(int32(7))%32)) + int32(base.Ui32(v1041)>>(uint(v1042)%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(2))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(3))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(4))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(5))%32))&v1042 + int32(base.Ui32(v1041)>>(uint(int32(6))%32))&v1042
																v1081 = v1036 + v1042
																if v1081 != v59 {
																	v1032 = v1032 + v1042
																	v1033 = v1077
																	v1036 = v1081
																	continue
																} else {
																	break
																}
																break
															}
															v1093 = v1030 - v1077
														}
													} else {
														if v59 <= int32(0) {
															v1093 = int32(0)
														} else {
															v944 = int32(8)
															v945 = v915 + v944
															v947 = v382 + v944
															v948 = int32(0)
															v950 = int32(1)
															v952 = v59 << (uint(int32(3)) % 32)
															if v952 <= v950 {
																v955 = v950
															} else {
																v955 = v952
															}
															if v955 != int32(1) {
																v963 = v948
																v964 = v948
																v965 = int32(0)
																for {
																	v973 = int32(base.Ui32(v964) >> (uint(int32(3)) % 32))
																	v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v973))))
																	v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947+v973))))
																	v978 = v975 ^ v977
																	v980 = v964 & int32(6)
																	v981 = int32(1)
																	v990 = int32(base.Ui32(v978)>>(uint(v980|v981)%32))&v981 + (int32(base.Ui32(v978)>>(uint(v980)%32))&v981 + v963)
																	v991 = int32(2)
																	v992 = v964 + v991
																	v994 = v965 + v991
																	if v994 != v955&int32(2147483640) {
																		v963 = v990
																		v964 = v992
																		v965 = v994
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v955&int32(1) == int32(0) {
																	v1020 = v990
																} else {
																	v998 = v990
																	v999 = v992
																	v1008 = int32(base.Ui32(v999) >> (uint(int32(3)) % 32))
																	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v1008))))
																	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008+v947))))
																	v1020 = v998 + int32(base.Ui32(v1010^v1012)>>(uint(v999&int32(7))%32))&int32(1)
																}
															} else {
																v998 = v948
																v999 = v948
																v1008 = int32(base.Ui32(v999) >> (uint(int32(3)) % 32))
																v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945+v1008))))
																v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1008+v947))))
																v1020 = v998 + int32(base.Ui32(v1010^v1012)>>(uint(v999&int32(7))%32))&int32(1)
															}
															v1093 = v1020
														}
													}
												}
												v1101 = *(*int32)(unsafe.Add(mBase, uint32(v915)+4))
												v1102 = int32(4)
												v1103 = v1101 & v1102
												v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
												if v1104&v1102 != 0 {
													if v1103 != 0 {
														v1272 = int32(0)
													} else {
														v1109 = v59 << (uint(int32(3)) % 32)
														if v59 <= int32(0) {
															v1272 = v1109
														} else {
															v1208 = v915 + int32(8)
															v1209 = v1109
															v1211 = v1208
															v1212 = int32(0)
															v1215 = int32(0)
															for {
																v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
																v1221 = int32(1)
																v1256 = v1212 + v1220&v1221 + int32(base.Ui32(v1220)>>(uint(int32(7))%32)) + int32(base.Ui32(v1220)>>(uint(v1221)%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(2))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(3))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(4))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(5))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(6))%32))&v1221
																v1260 = v1215 + v1221
																if v1260 != v59 {
																	v1211 = v1211 + v1221
																	v1212 = v1256
																	v1215 = v1260
																	continue
																} else {
																	break
																}
																break
															}
															v1272 = v1209 - v1256
														}
													}
												} else {
													if v1103 != 0 {
														v1115 = v59 << (uint(int32(3)) % 32)
														if v59 <= int32(0) {
															v1272 = v1115
														} else {
															v1208 = v411 + int32(8)
															v1209 = v1115
															v1211 = v1208
															v1212 = int32(0)
															v1215 = int32(0)
															for {
																v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211))))
																v1221 = int32(1)
																v1256 = v1212 + v1220&v1221 + int32(base.Ui32(v1220)>>(uint(int32(7))%32)) + int32(base.Ui32(v1220)>>(uint(v1221)%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(2))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(3))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(4))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(5))%32))&v1221 + int32(base.Ui32(v1220)>>(uint(int32(6))%32))&v1221
																v1260 = v1215 + v1221
																if v1260 != v59 {
																	v1211 = v1211 + v1221
																	v1212 = v1256
																	v1215 = v1260
																	continue
																} else {
																	break
																}
																break
															}
															v1272 = v1209 - v1256
														}
													} else {
														if v59 <= int32(0) {
															v1272 = int32(0)
														} else {
															v1123 = int32(8)
															v1124 = v915 + v1123
															v1126 = v411 + v1123
															v1127 = int32(0)
															v1129 = int32(1)
															v1131 = v59 << (uint(int32(3)) % 32)
															if v1131 <= v1129 {
																v1134 = v1129
															} else {
																v1134 = v1131
															}
															if v1134 != int32(1) {
																v1142 = v1127
																v1143 = v1127
																v1144 = int32(0)
																for {
																	v1152 = int32(base.Ui32(v1143) >> (uint(int32(3)) % 32))
																	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1152))))
																	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1126+v1152))))
																	v1157 = v1154 ^ v1156
																	v1159 = v1143 & int32(6)
																	v1160 = int32(1)
																	v1169 = int32(base.Ui32(v1157)>>(uint(v1159|v1160)%32))&v1160 + (int32(base.Ui32(v1157)>>(uint(v1159)%32))&v1160 + v1142)
																	v1170 = int32(2)
																	v1171 = v1143 + v1170
																	v1173 = v1144 + v1170
																	if v1173 != v1134&int32(2147483640) {
																		v1142 = v1169
																		v1143 = v1171
																		v1144 = v1173
																		continue
																	} else {
																		break
																	}
																	break
																}
																if v1134&int32(1) == int32(0) {
																	v1199 = v1169
																} else {
																	v1177 = v1169
																	v1178 = v1171
																	v1187 = int32(base.Ui32(v1178) >> (uint(int32(3)) % 32))
																	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1187))))
																	v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v1126))))
																	v1199 = v1177 + int32(base.Ui32(v1189^v1191)>>(uint(v1178&int32(7))%32))&int32(1)
																}
															} else {
																v1177 = v1127
																v1178 = v1127
																v1187 = int32(base.Ui32(v1178) >> (uint(int32(3)) % 32))
																v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1124+v1187))))
																v1191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v1126))))
																v1199 = v1177 + int32(base.Ui32(v1189^v1191)>>(uint(v1178&int32(7))%32))&int32(1)
															}
															v1272 = v1199
														}
													}
												}
												v1274 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
												v1275 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
												v1276 = v1274 - v1275
												if base.F64_lt(base.F64_convert_i32_s(v1093), base.F64_add(base.F64_convert_i32_s(v1272), base.F64_mul(base.F64_convert_i32_s(v1276*v1276*v1276), float64(-0.0001)))) != 0 {
													v1284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
													if v1284&int32(4) != 0 {
													} else {
														v1287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+4)))
														if v1287&int32(4) != 0 {
															if v59 == int32(0) {
															} else {
																base.MemoryFill(m, v867, int32(255), v59)
															}
														} else {
															if v59 <= int32(0) {
															} else {
																v1297 = v915 + int32(8)
																v1298 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v59) {
																	v1304 = v1298
																	v1305 = v1298
																	for {
																		v1328 = v1304 + v867
																		v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1328))))
																		v1331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1304+v1297))))
																		v1332 = v1329 | v1331
																		*(*uint8)(unsafe.Add(mBase, uint32(v1328))) = uint8(v1332)
																		v1335 = v1304 | int32(1)
																		v1336 = v867 + v1335
																		v1337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1336))))
																		v1339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1335+v1297))))
																		v1340 = v1337 | v1339
																		*(*uint8)(unsafe.Add(mBase, uint32(v1336))) = uint8(v1340)
																		v1343 = v1304 | int32(2)
																		v1344 = v867 + v1343
																		v1345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1344))))
																		v1347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1343+v1297))))
																		v1348 = v1345 | v1347
																		*(*uint8)(unsafe.Add(mBase, uint32(v1344))) = uint8(v1348)
																		v1351 = v1304 | int32(3)
																		v1352 = v867 + v1351
																		v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352))))
																		v1355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1351+v1297))))
																		v1356 = v1353 | v1355
																		*(*uint8)(unsafe.Add(mBase, uint32(v1352))) = uint8(v1356)
																		v1358 = int32(4)
																		v1359 = v1304 + v1358
																		v1361 = v1305 + v1358
																		if v1361 != v861 {
																			v1304 = v1359
																			v1305 = v1361
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v863 == int32(0) {
																	} else {
																		v1367 = v1359
																		v1392 = v1367
																		v1404 = v1298
																		for {
																			v1415 = v1392 + v867
																			v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415))))
																			v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1297))))
																			v1419 = v1416 | v1418
																			*(*uint8)(unsafe.Add(mBase, uint32(v1415))) = uint8(v1419)
																			v1421 = int32(1)
																			v1424 = v1404 + v1421
																			if v1424 != v863 {
																				v1392 = v1392 + v1421
																				v1404 = v1424
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v1367 = v1298
																	v1392 = v1367
																	v1404 = v1298
																	for {
																		v1415 = v1392 + v867
																		v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415))))
																		v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1392+v1297))))
																		v1419 = v1416 | v1418
																		*(*uint8)(unsafe.Add(mBase, uint32(v1415))) = uint8(v1419)
																		v1421 = int32(1)
																		v1424 = v1404 + v1421
																		if v1424 != v863 {
																			v1392 = v1392 + v1421
																			v1404 = v1424
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
													*(*uint16)(unsafe.Add(mBase, uint32(v880))) = uint16(v899)
													v1452 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v1452 + int32(1)
													v1666 = v880 + int32(2)
													v1673 = v887
												} else {
													v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+4)))
													if v1458&int32(4) != 0 {
													} else {
														v1461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+4)))
														if v1461&int32(4) != 0 {
															if v59 == int32(0) {
															} else {
																base.MemoryFill(m, v865, int32(255), v59)
															}
														} else {
															if v59 <= int32(0) {
															} else {
																v1471 = v915 + int32(8)
																v1472 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v59) {
																	v1478 = v1472
																	v1479 = v1472
																	for {
																		v1502 = v1478 + v865
																		v1503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1502))))
																		v1505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478+v1471))))
																		v1506 = v1503 | v1505
																		*(*uint8)(unsafe.Add(mBase, uint32(v1502))) = uint8(v1506)
																		v1509 = v1478 | int32(1)
																		v1510 = v865 + v1509
																		v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510))))
																		v1513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1509+v1471))))
																		v1514 = v1511 | v1513
																		*(*uint8)(unsafe.Add(mBase, uint32(v1510))) = uint8(v1514)
																		v1517 = v1478 | int32(2)
																		v1518 = v865 + v1517
																		v1519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518))))
																		v1521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1517+v1471))))
																		v1522 = v1519 | v1521
																		*(*uint8)(unsafe.Add(mBase, uint32(v1518))) = uint8(v1522)
																		v1525 = v1478 | int32(3)
																		v1526 = v865 + v1525
																		v1527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1526))))
																		v1529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525+v1471))))
																		v1530 = v1527 | v1529
																		*(*uint8)(unsafe.Add(mBase, uint32(v1526))) = uint8(v1530)
																		v1532 = int32(4)
																		v1533 = v1478 + v1532
																		v1535 = v1479 + v1532
																		if v1535 != v861 {
																			v1478 = v1533
																			v1479 = v1535
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	if v863 == int32(0) {
																	} else {
																		v1541 = v1533
																		v1566 = v1541
																		v1578 = v1472
																		for {
																			v1589 = v1566 + v865
																			v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
																			v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566+v1471))))
																			v1593 = v1590 | v1592
																			*(*uint8)(unsafe.Add(mBase, uint32(v1589))) = uint8(v1593)
																			v1595 = int32(1)
																			v1598 = v1578 + v1595
																			if v1598 != v863 {
																				v1566 = v1566 + v1595
																				v1578 = v1598
																				continue
																			} else {
																				break
																			}
																			break
																		}
																	}
																} else {
																	v1541 = v1472
																	v1566 = v1541
																	v1578 = v1472
																	for {
																		v1589 = v1566 + v865
																		v1590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589))))
																		v1592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566+v1471))))
																		v1593 = v1590 | v1592
																		*(*uint8)(unsafe.Add(mBase, uint32(v1589))) = uint8(v1593)
																		v1595 = int32(1)
																		v1598 = v1578 + v1595
																		if v1598 != v863 {
																			v1566 = v1566 + v1595
																			v1578 = v1598
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
													*(*uint16)(unsafe.Add(mBase, uint32(v887))) = uint16(v899)
													v1651 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
													*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v1651 + int32(1)
													v1666 = v880
													v1673 = v887 + int32(2)
												}
											}
										}
										v1683 = v871 + int32(1)
										if v1683 != v859 {
											v871 = v1683
											v880 = v1666
											v887 = v1673
											continue
										} else {
											break
										}
										break
									}
									v1694 = v1666
									v1701 = v1673
									v1710 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1694))) = uint16(v1710)
									*(*uint16)(unsafe.Add(mBase, uint32(v1701))) = uint16(v1710)
									*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = base.I64_extend_i32_u(v411)
									*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = base.I64_extend_i32_u(v382)
									return v30 & int64(4294967295)
								}
							}
						}
					}
				}
			}
		}
	}
}
