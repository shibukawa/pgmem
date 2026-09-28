package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gtrgm_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v291 int32
	_ = v291
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
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
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v703 int64
	_ = v703
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v736 int64
	_ = v736
	var v738 int32
	_ = v738
	var v741 int64
	_ = v741
	var v742 int32
	_ = v742
	var v745 int64
	_ = v745
	var v746 int32
	_ = v746
	var v749 int64
	_ = v749
	var v750 int32
	_ = v750
	var v753 int64
	_ = v753
	var v757 int64
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v795 int64
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v828 int64
	_ = v828
	var v830 int32
	_ = v830
	var v833 int64
	_ = v833
	var v834 int64
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v945 int64
	_ = v945
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v960 int64
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int64
	_ = v967
	var v968 int32
	_ = v968
	var v969 int64
	_ = v969
	var v970 int32
	_ = v970
	var v971 int64
	_ = v971
	var v972 int32
	_ = v972
	var v973 int64
	_ = v973
	var v977 int64
	_ = v977
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v985 int64
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int64
	_ = v992
	var v996 int32
	_ = v996
	var v997 int64
	_ = v997
	var v998 int64
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1006 int64
	_ = v1006
	var v1016 int64
	_ = v1016
	var v1047 int64
	_ = v1047
	var v1058 int32
	_ = v1058
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int64
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1138 int64
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1143 int64
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int64
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int64
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int64
	_ = v1155
	var v1159 int64
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1197 int64
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1230 int64
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1235 int64
	_ = v1235
	var v1236 int64
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1346 int64
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1361 int64
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int64
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int64
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int64
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int64
	_ = v1374
	var v1378 int64
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1384 int32
	_ = v1384
	var v1386 int64
	_ = v1386
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int64
	_ = v1393
	var v1397 int32
	_ = v1397
	var v1398 int64
	_ = v1398
	var v1399 int64
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1407 int64
	_ = v1407
	var v1417 int64
	_ = v1417
	var v1448 int64
	_ = v1448
	var v1455 int32
	_ = v1455
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1587 int32
	_ = v1587
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1659 int32
	_ = v1659
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1703 int32
	_ = v1703
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1719 int32
	_ = v1719
	var v1723 int32
	_ = v1723
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1928 int32
	_ = v1928
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1967 int32
	_ = v1967
	v2 = int32(0)
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v37 = v35 + int32(_a_F_gtrgm_picksplit_0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v39 == v2 {
		v56 = v2
	} else {
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
		if v43 == int32(0) {
			v56 = v2
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
			if v46 != int32(7) {
				v56 = v2
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				if v49 != int32(17) {
					v56 = v2
				} else {
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+32)))
					v56 = v52 ^ int32(1)
				}
			}
		}
	}
	if v56&int32(1) != 0 {
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v60 = F_get_fn_opclass_options(m, v59)
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return int64(0)
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
			v65 = v64
			v68 = v37 & int32(_a_F_gtrgm_picksplit_0)
			v70 = v68 + int32(1)
			v71 = F_palloc_mul(m, int32(8), v70)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int64(0)
			} else {
				v74 = F_palloc(m, v70*v65)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					v77 = v35 & int32(_a_F_gtrgm_picksplit_0)
					if v77 != int32(1) {
						v80 = int32(3)
						v87 = int32(8)
						v91 = int32(1)
						v92 = v65<<(uint(v80)%32) - v91
						v94 = base.I32_div_s(v92, v87)
						v97 = v91
						for {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(v34+v87+v97*int32(24))))
							v134 = v71 + v97<<(uint(int32(3))%32)
							v136 = v74 + v97*v65
							*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v136
							v138 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v138)
							v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+4)))
							if v140&int32(1) != 0 {
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
								v147 = int32(base.Ui32(v143)>>(uint(int32(2))%32)) - int32(5)
								v148 = int32(3)
								v149 = base.I32_div_u_s(v147, v148)
								v152 = int32(0)
								if base.B2i32(v65&v80 != int32(0))|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v65))|base.B2i32(v136&v148 != v152) == v152 {
									if v65 == int32(0) {
									} else {
										v161 = v65 + v136
										v163 = v136 + int32(4)
										if base.Ui32(v163) < base.Ui32(v161) {
											v165 = v161
										} else {
											v165 = v163
										}
										v171 = (v136^int32(-1)+v165)&int32(-4) + int32(4)
										if v171 == int32(0) {
										} else {
											base.MemoryFill(m, v136, int32(0), v171)
										}
									}
								} else {
									v171 = v65
									if v171 == int32(0) {
									} else {
										base.MemoryFill(m, v136, int32(0), v171)
									}
								}
								v179 = v136 + v94
								v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
								v182 = v180 | int32(128)
								*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v182)
								if base.Ui32(v147) < base.Ui32(int32(3)) {
								} else {
									v188 = int32(1)
									if base.Ui32(v149) <= base.Ui32(v188) {
										v191 = v188
									} else {
										v191 = v149
									}
									v193 = int32(0)
									for {
										v225 = int32(3)
										v227 = v131 + int32(5) + v193*v225
										v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227))))
										v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+2)))
										v233 = base.I32_rem_u_s(v228|v229<<(uint(int32(16))%32), v92)
										v236 = v136 + int32(base.Ui32(v233)>>(uint(v225)%32))
										v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
										v238 = int32(1)
										v242 = v237 | v238<<(uint(v233&int32(7))%32)
										*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v242)
										v245 = v193 + v238
										if v245 != v191 {
											v193 = v245
											continue
										} else {
											break
										}
										break
									}
								}
							} else {
								if v140&int32(4) != 0 {
									v249 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v249)
								} else {
									if v65 == int32(0) {
									} else {
										base.MemoryCopy(m, v136, v131+int32(5), v65)
									}
								}
							}
							v291 = (v97 + int32(1)) & int32(_a_F_gtrgm_picksplit_0)
							if base.Ui32(v291) <= base.Ui32(v68) {
								v97 = v291
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v325 = int32(0)
					if base.Ui32(int32(2)) <= base.Ui32(v68) {
						v332 = int32(1)
						v333 = v325
						v335 = int32(-1)
						v342 = v325
						for {
							v367 = v332 + int32(1)
							v368 = v367
							v370 = v333
							v372 = v335
							v376 = v367
							v379 = v342
							for {
								v403 = F_hemdistcache_2(m, v71+v376<<(uint(int32(3))%32), v71+v332<<(uint(int32(3))%32), v65)
								mBase = m.M
								v404 = base.B2i32(v372 < v403)
								if v372 < v403 {
									v405 = v403
								} else {
									v405 = v372
								}
								if v372 < v403 {
									v406 = v368
								} else {
									v406 = v379
								}
								if v372 < v403 {
									v407 = v332
								} else {
									v407 = v370
								}
								v409 = v368 + int32(1)
								v410 = int32(_a_F_gtrgm_picksplit_0)
								v411 = v409 & v410
								if base.Ui32(v411) <= base.Ui32(v37&v410) {
									v368 = v409
									v370 = v407
									v372 = v405
									v376 = v411
									v379 = v406
									continue
								} else {
									break
								}
								break
							}
							if v367 != v68 {
								v332 = v367
								v333 = v407
								v335 = v405
								v342 = v406
								continue
							} else {
								break
							}
							break
						}
						v418 = v407
						v427 = v406
					} else {
						v418 = v325
						v427 = v325
					}
					v448 = base.I32_wrap_i64(v33)
					v450 = v68 << (uint(int32(1)) % 32)
					v451 = F_palloc(m, v450)
					mBase = m.M
					v452 = m.ExcPending
					if v452 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v448))) = v451
						v454 = F_palloc(m, v450)
						mBase = m.M
						v455 = m.ExcPending
						if v455 != 0 {
							return int64(0)
						} else {
							v456 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v456
							*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v456
							*(*int32)(unsafe.Add(mBase, uint32(v448)+20)) = v454
							v462 = int32(_a_F_gtrgm_picksplit_0)
							v470 = base.B2i32(v418&v462 == v456) | base.B2i32(v427&v462 == v456)
							if v470 != 0 {
								v471 = int32(1)
							} else {
								v471 = v418
							}
							v476 = v71 + v471&int32(_a_F_gtrgm_picksplit_0)<<(uint(int32(3))%32)
							v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
							v478 = int32(5)
							v480 = v65 + v478
							v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
							if v481 != 0 {
								v482 = v478
							} else {
								v482 = v480
							}
							v483 = F_palloc(m, v482)
							mBase = m.M
							v484 = m.ExcPending
							if v484 != 0 {
								return int64(0)
							} else {
								if v481 != 0 {
									v487 = int32(6)
								} else {
									v487 = int32(2)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)) = uint8(v487)
								v489 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v483))) = v482 << (uint(v489) % 32)
								if v470 != 0 {
									v493 = v489
								} else {
									v493 = v427
								}
								if v481 != 0 {
								} else {
									v495 = v483 + int32(5)
									if v477 != 0 {
										if v65 == int32(0) {
										} else {
											base.MemoryCopy(m, v495, v477, v65)
										}
									} else {
										if v65 == int32(0) {
										} else {
											base.MemoryFill(m, v495, int32(0), v65)
										}
									}
								}
								v508 = v71 + v493&int32(_a_F_gtrgm_picksplit_0)<<(uint(int32(3))%32)
								v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
								v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
								if v511 != 0 {
									v512 = int32(5)
								} else {
									v512 = v480
								}
								v513 = F_palloc(m, v512)
								mBase = m.M
								v514 = m.ExcPending
								if v514 != 0 {
									return int64(0)
								} else {
									if v511 != 0 {
										v517 = int32(6)
									} else {
										v517 = int32(2)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)) = uint8(v517)
									*(*int32)(unsafe.Add(mBase, uint32(v513))) = v512 << (uint(int32(2)) % 32)
									if v511 != 0 {
									} else {
										v523 = v513 + int32(5)
										if v509 != 0 {
											if v65 == int32(0) {
											} else {
												base.MemoryCopy(m, v523, v509, v65)
											}
										} else {
											if v65 == int32(0) {
											} else {
												base.MemoryFill(m, v523, int32(0), v65)
											}
										}
									}
									v535 = F_palloc_mul(m, int32(8), v68)
									mBase = m.M
									v536 = m.ExcPending
									if v536 != 0 {
										return int64(0)
									} else {
										if v77 == int32(1) {
											F_pg_qsort(m, v535, v68, int32(8), int32(_a_F_gtrgm_picksplit_1))
											mBase = m.M
											v542 = m.ExcPending
											if v542 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v448)+32)) = base.I64_extend_i32_u(v513)
												*(*int64)(unsafe.Add(mBase, uint32(v448)+8)) = base.I64_extend_i32_u(v483)
												return v33 & int64(4294967295)
											}
										} else {
											v543 = int32(5)
											v544 = v513 + v543
											v546 = v483 + v543
											v547 = int32(1)
											v549 = v547
											v551 = v547
											for {
												v582 = v551 << (uint(int32(3)) % 32)
												v583 = v535 + v582
												*(*uint16)(unsafe.Add(mBase, uint32(v583-int32(8)))) = uint16(v549)
												v589 = v582 + v71
												v590 = F_hemdistcache_2(m, v476, v589, v65)
												mBase = m.M
												v591 = F_hemdistcache_2(m, v508, v589, v65)
												mBase = m.M
												v592 = v590 - v591
												v594 = v592 >> (uint(int32(31)) % 32)
												*(*int32)(unsafe.Add(mBase, uint32(v583-int32(4)))) = v592 ^ v594 - v594
												v599 = v549 + int32(1)
												v600 = int32(_a_F_gtrgm_picksplit_0)
												v601 = v599 & v600
												if base.Ui32(v601) <= base.Ui32(v37&v600) {
													v549 = v599
													v551 = v601
													continue
												} else {
													break
												}
												break
											}
											F_pg_qsort(m, v535, v68, int32(8), int32(_a_F_gtrgm_picksplit_1))
											mBase = m.M
											v608 = m.ExcPending
											if v608 != 0 {
												return int64(0)
											} else {
												v609 = int32(1)
												if base.Ui32(v68) <= base.Ui32(v609) {
													v612 = v609
												} else {
													v612 = v68
												}
												v614 = v65 & int32(2147483644)
												v615 = int32(3)
												v616 = v65 & v615
												v618 = v65 & int32(-4)
												v620 = v65 & int32(2147483646)
												v621 = int32(1)
												v622 = v65 & v621
												v624 = v65 - v621
												v626 = v65 << (uint(v615) % 32)
												v631 = base.B2i32(int32(7) < v65)
												v643 = int32(0)
												v645 = v454
												v646 = v451
												for {
													v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v535+v643<<(uint(int32(3))%32)))))
													if v471&int32(_a_F_gtrgm_picksplit_0) == v667 {
														*(*uint16)(unsafe.Add(mBase, uint32(v646))) = uint16(v471)
														v670 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v670 + int32(1)
														v1947 = v645
														v1948 = v646 + int32(2)
													} else {
														if v493&int32(_a_F_gtrgm_picksplit_0) == v667 {
															*(*uint16)(unsafe.Add(mBase, uint32(v645))) = uint16(v493)
															v1928 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
															*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v1928 + int32(1)
															v1947 = v645 + int32(2)
															v1948 = v646
														} else {
															v682 = v71 + v667<<(uint(int32(3))%32)
															v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
															v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
															if v684&int32(4) == int32(0) {
																if v683&int32(1) != 0 {
																	v698 = v546
																	if int32(7) < v65 {
																		v945 = int64(0)
																		v946 = int32(0)
																		if v65 == v946 {
																			v1016 = int64(0)
																		} else {
																			v953 = v65 & int32(3)
																			if base.Ui32(int32(4)) <= base.Ui32(v65) {
																				v958 = v698
																				v960 = v945
																				v963 = v946
																				for {
																					v964 = int32(4)
																					v965 = v958 + v964
																					v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+3)))
																					v967 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v966)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+2)))
																					v969 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v968)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+1)))
																					v971 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v970)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
																					v973 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v972)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v977 = v967 + (v969 + (v971 + (v960 + v973)))
																					v979 = v963 + v964
																					if v979 != v65&int32(-4) {
																						v958 = v965
																						v960 = v977
																						v963 = v979
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v953 == int32(0) {
																					v1006 = v977
																				} else {
																					v983 = v965
																					v985 = v977
																					v990 = v983
																					v991 = int32(0)
																					v992 = v985
																					for {
																						v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																						v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v998 = v992 + v997
																						v999 = int32(1)
																						v1002 = v991 + v999
																						if v1002 != v953 {
																							v990 = v990 + v999
																							v991 = v1002
																							v992 = v998
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1006 = v998
																				}
																			} else {
																				v983 = v698
																				v985 = v945
																				v990 = v983
																				v991 = int32(0)
																				v992 = v985
																				for {
																					v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																					v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v998 = v992 + v997
																					v999 = int32(1)
																					v1002 = v991 + v999
																					if v1002 != v953 {
																						v990 = v990 + v999
																						v991 = v1002
																						v992 = v998
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1006 = v998
																			}
																			v1016 = v1006
																		}
																		v1047 = v1016
																	} else {
																		if v65 == int32(0) {
																			v1047 = int64(0)
																		} else {
																			v703 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v624) {
																				v706 = v698
																				v711 = int32(0)
																				v736 = v703
																				for {
																					v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+3)))
																					v741 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v738)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+2)))
																					v745 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v742)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
																					v749 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v746)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
																					v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v750)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v757 = v741 + (v745 + (v749 + (v736 + v753)))
																					v758 = int32(4)
																					v759 = v706 + v758
																					v761 = v711 + v758
																					if v761 != v618 {
																						v706 = v759
																						v711 = v761
																						v736 = v757
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																					v1047 = v757
																				} else {
																					v765 = v759
																					v795 = v757
																					v798 = v765
																					v799 = int32(0)
																					v828 = v795
																					for {
																						v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																						v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v834 = v828 + v833
																						v835 = int32(1)
																						v838 = v799 + v835
																						if v838 != v616 {
																							v798 = v798 + v835
																							v799 = v838
																							v828 = v834
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1047 = v834
																				}
																			} else {
																				v765 = v698
																				v795 = v703
																				v798 = v765
																				v799 = int32(0)
																				v828 = v795
																				for {
																					v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																					v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v834 = v828 + v833
																					v835 = int32(1)
																					v838 = v799 + v835
																					if v838 != v616 {
																						v798 = v798 + v835
																						v799 = v838
																						v828 = v834
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1047 = v834
																			}
																		}
																	}
																	v1058 = v626 + (base.I32_wrap_i64(v1047) ^ int32(-1))
																} else {
																	if int32(0) < v65 {
																		v840 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																		v841 = int32(0)
																		if v624 != 0 {
																			v845 = v841
																			v847 = v841
																			v850 = v841
																			for {
																				v878 = v845 | int32(1)
																				v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v878))))
																				v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v878))))
																				v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880^v882)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845+v546))))
																				v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845+v840))))
																				v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888^v890)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v896 = v886 + (v850 + v894)
																				v897 = int32(2)
																				v898 = v845 + v897
																				v900 = v847 + v897
																				if v900 != v620 {
																					v845 = v898
																					v847 = v900
																					v850 = v896
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v622 == int32(0) {
																				v1058 = v896
																			} else {
																				v908 = v898
																				v909 = v896
																				v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v908))))
																				v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908+v546))))
																				v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937^v939)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1058 = v909 + v943
																			}
																		} else {
																			v908 = v841
																			v909 = v841
																			v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v908))))
																			v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908+v546))))
																			v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937^v939)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1058 = v909 + v943
																		}
																	} else {
																		v1058 = int32(0)
																	}
																}
															} else {
																if v683&int32(1) != 0 {
																	v1058 = int32(0)
																} else {
																	v697 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																	v698 = v697
																	if int32(7) < v65 {
																		v945 = int64(0)
																		v946 = int32(0)
																		if v65 == v946 {
																			v1016 = int64(0)
																		} else {
																			v953 = v65 & int32(3)
																			if base.Ui32(int32(4)) <= base.Ui32(v65) {
																				v958 = v698
																				v960 = v945
																				v963 = v946
																				for {
																					v964 = int32(4)
																					v965 = v958 + v964
																					v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+3)))
																					v967 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v966)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+2)))
																					v969 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v968)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+1)))
																					v971 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v970)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
																					v973 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v972)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v977 = v967 + (v969 + (v971 + (v960 + v973)))
																					v979 = v963 + v964
																					if v979 != v65&int32(-4) {
																						v958 = v965
																						v960 = v977
																						v963 = v979
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v953 == int32(0) {
																					v1006 = v977
																				} else {
																					v983 = v965
																					v985 = v977
																					v990 = v983
																					v991 = int32(0)
																					v992 = v985
																					for {
																						v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																						v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v998 = v992 + v997
																						v999 = int32(1)
																						v1002 = v991 + v999
																						if v1002 != v953 {
																							v990 = v990 + v999
																							v991 = v1002
																							v992 = v998
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1006 = v998
																				}
																			} else {
																				v983 = v698
																				v985 = v945
																				v990 = v983
																				v991 = int32(0)
																				v992 = v985
																				for {
																					v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																					v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v998 = v992 + v997
																					v999 = int32(1)
																					v1002 = v991 + v999
																					if v1002 != v953 {
																						v990 = v990 + v999
																						v991 = v1002
																						v992 = v998
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1006 = v998
																			}
																			v1016 = v1006
																		}
																		v1047 = v1016
																	} else {
																		if v65 == int32(0) {
																			v1047 = int64(0)
																		} else {
																			v703 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v624) {
																				v706 = v698
																				v711 = int32(0)
																				v736 = v703
																				for {
																					v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+3)))
																					v741 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v738)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+2)))
																					v745 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v742)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
																					v749 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v746)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
																					v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v750)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v757 = v741 + (v745 + (v749 + (v736 + v753)))
																					v758 = int32(4)
																					v759 = v706 + v758
																					v761 = v711 + v758
																					if v761 != v618 {
																						v706 = v759
																						v711 = v761
																						v736 = v757
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																					v1047 = v757
																				} else {
																					v765 = v759
																					v795 = v757
																					v798 = v765
																					v799 = int32(0)
																					v828 = v795
																					for {
																						v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																						v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v834 = v828 + v833
																						v835 = int32(1)
																						v838 = v799 + v835
																						if v838 != v616 {
																							v798 = v798 + v835
																							v799 = v838
																							v828 = v834
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1047 = v834
																				}
																			} else {
																				v765 = v698
																				v795 = v703
																				v798 = v765
																				v799 = int32(0)
																				v828 = v795
																				for {
																					v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																					v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v834 = v828 + v833
																					v835 = int32(1)
																					v838 = v799 + v835
																					if v838 != v616 {
																						v798 = v798 + v835
																						v799 = v838
																						v828 = v834
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1047 = v834
																			}
																		}
																	}
																	v1058 = v626 + (base.I32_wrap_i64(v1047) ^ int32(-1))
																}
															}
															v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
															v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)))
															if v1086&int32(4) == int32(0) {
																if v1085&int32(1) != 0 {
																	v1100 = v544
																	if int32(7) < v65 {
																		v1346 = int64(0)
																		v1347 = int32(0)
																		if v65 == v1347 {
																			v1417 = int64(0)
																		} else {
																			v1354 = v65 & int32(3)
																			if base.Ui32(int32(4)) <= base.Ui32(v65) {
																				v1359 = v1100
																				v1361 = v1346
																				v1364 = v1347
																				for {
																					v1365 = int32(4)
																					v1366 = v1359 + v1365
																					v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+3)))
																					v1368 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+2)))
																					v1370 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1369)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+1)))
																					v1372 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359))))
																					v1374 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1378 = v1368 + (v1370 + (v1372 + (v1361 + v1374)))
																					v1380 = v1364 + v1365
																					if v1380 != v65&int32(-4) {
																						v1359 = v1366
																						v1361 = v1378
																						v1364 = v1380
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v1354 == int32(0) {
																					v1407 = v1378
																				} else {
																					v1384 = v1366
																					v1386 = v1378
																					v1391 = v1384
																					v1392 = int32(0)
																					v1393 = v1386
																					for {
																						v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																						v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v1399 = v1393 + v1398
																						v1400 = int32(1)
																						v1403 = v1392 + v1400
																						if v1403 != v1354 {
																							v1391 = v1391 + v1400
																							v1392 = v1403
																							v1393 = v1399
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1407 = v1399
																				}
																			} else {
																				v1384 = v1100
																				v1386 = v1346
																				v1391 = v1384
																				v1392 = int32(0)
																				v1393 = v1386
																				for {
																					v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																					v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1399 = v1393 + v1398
																					v1400 = int32(1)
																					v1403 = v1392 + v1400
																					if v1403 != v1354 {
																						v1391 = v1391 + v1400
																						v1392 = v1403
																						v1393 = v1399
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1407 = v1399
																			}
																			v1417 = v1407
																		}
																		v1448 = v1417
																	} else {
																		if v65 == int32(0) {
																			v1448 = int64(0)
																		} else {
																			v1105 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v624) {
																				v1108 = v1100
																				v1110 = int32(0)
																				v1138 = v1105
																				for {
																					v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+3)))
																					v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+2)))
																					v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
																					v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																					v1155 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1159 = v1143 + (v1147 + (v1151 + (v1138 + v1155)))
																					v1160 = int32(4)
																					v1161 = v1108 + v1160
																					v1163 = v1110 + v1160
																					if v1163 != v618 {
																						v1108 = v1161
																						v1110 = v1163
																						v1138 = v1159
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																					v1448 = v1159
																				} else {
																					v1167 = v1161
																					v1197 = v1159
																					v1200 = v1167
																					v1201 = int32(0)
																					v1230 = v1197
																					for {
																						v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																						v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v1236 = v1230 + v1235
																						v1237 = int32(1)
																						v1240 = v1201 + v1237
																						if v1240 != v616 {
																							v1200 = v1200 + v1237
																							v1201 = v1240
																							v1230 = v1236
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1448 = v1236
																				}
																			} else {
																				v1167 = v1100
																				v1197 = v1105
																				v1200 = v1167
																				v1201 = int32(0)
																				v1230 = v1197
																				for {
																					v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																					v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1236 = v1230 + v1235
																					v1237 = int32(1)
																					v1240 = v1201 + v1237
																					if v1240 != v616 {
																						v1200 = v1200 + v1237
																						v1201 = v1240
																						v1230 = v1236
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1448 = v1236
																			}
																		}
																	}
																	v1455 = v626 + (base.I32_wrap_i64(v1448) ^ int32(-1))
																} else {
																	if int32(0) < v65 {
																		v1242 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																		v1243 = int32(0)
																		if v624 != 0 {
																			v1246 = v1243
																			v1247 = v1243
																			v1250 = v1243
																			for {
																				v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1242))))
																				v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v544))))
																				v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279^v1281)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1288 = v1246 | int32(1)
																				v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+v1288))))
																				v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288+v1242))))
																				v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290^v1292)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1297 = v1247 + v1285 + v1296
																				v1298 = int32(2)
																				v1299 = v1246 + v1298
																				v1301 = v1250 + v1298
																				if v1301 != v620 {
																					v1246 = v1299
																					v1247 = v1297
																					v1250 = v1301
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v622 == int32(0) {
																				v1455 = v1297
																			} else {
																				v1305 = v1299
																				v1306 = v1297
																				v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v1242))))
																				v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v544))))
																				v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338^v1340)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1455 = v1306 + v1344
																			}
																		} else {
																			v1305 = v1243
																			v1306 = v1243
																			v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v1242))))
																			v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v544))))
																			v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338^v1340)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1455 = v1306 + v1344
																		}
																	} else {
																		v1455 = int32(0)
																	}
																}
															} else {
																if v1085&int32(1) != 0 {
																	v1455 = int32(0)
																} else {
																	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																	v1100 = v1099
																	if int32(7) < v65 {
																		v1346 = int64(0)
																		v1347 = int32(0)
																		if v65 == v1347 {
																			v1417 = int64(0)
																		} else {
																			v1354 = v65 & int32(3)
																			if base.Ui32(int32(4)) <= base.Ui32(v65) {
																				v1359 = v1100
																				v1361 = v1346
																				v1364 = v1347
																				for {
																					v1365 = int32(4)
																					v1366 = v1359 + v1365
																					v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+3)))
																					v1368 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+2)))
																					v1370 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1369)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+1)))
																					v1372 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359))))
																					v1374 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1378 = v1368 + (v1370 + (v1372 + (v1361 + v1374)))
																					v1380 = v1364 + v1365
																					if v1380 != v65&int32(-4) {
																						v1359 = v1366
																						v1361 = v1378
																						v1364 = v1380
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v1354 == int32(0) {
																					v1407 = v1378
																				} else {
																					v1384 = v1366
																					v1386 = v1378
																					v1391 = v1384
																					v1392 = int32(0)
																					v1393 = v1386
																					for {
																						v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																						v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v1399 = v1393 + v1398
																						v1400 = int32(1)
																						v1403 = v1392 + v1400
																						if v1403 != v1354 {
																							v1391 = v1391 + v1400
																							v1392 = v1403
																							v1393 = v1399
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1407 = v1399
																				}
																			} else {
																				v1384 = v1100
																				v1386 = v1346
																				v1391 = v1384
																				v1392 = int32(0)
																				v1393 = v1386
																				for {
																					v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																					v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1399 = v1393 + v1398
																					v1400 = int32(1)
																					v1403 = v1392 + v1400
																					if v1403 != v1354 {
																						v1391 = v1391 + v1400
																						v1392 = v1403
																						v1393 = v1399
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1407 = v1399
																			}
																			v1417 = v1407
																		}
																		v1448 = v1417
																	} else {
																		if v65 == int32(0) {
																			v1448 = int64(0)
																		} else {
																			v1105 = int64(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v624) {
																				v1108 = v1100
																				v1110 = int32(0)
																				v1138 = v1105
																				for {
																					v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+3)))
																					v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+2)))
																					v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
																					v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																					v1155 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1159 = v1143 + (v1147 + (v1151 + (v1138 + v1155)))
																					v1160 = int32(4)
																					v1161 = v1108 + v1160
																					v1163 = v1110 + v1160
																					if v1163 != v618 {
																						v1108 = v1161
																						v1110 = v1163
																						v1138 = v1159
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																					v1448 = v1159
																				} else {
																					v1167 = v1161
																					v1197 = v1159
																					v1200 = v1167
																					v1201 = int32(0)
																					v1230 = v1197
																					for {
																						v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																						v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																						v1236 = v1230 + v1235
																						v1237 = int32(1)
																						v1240 = v1201 + v1237
																						if v1240 != v616 {
																							v1200 = v1200 + v1237
																							v1201 = v1240
																							v1230 = v1236
																							continue
																						} else {
																							break
																						}
																						break
																					}
																					v1448 = v1236
																				}
																			} else {
																				v1167 = v1100
																				v1197 = v1105
																				v1200 = v1167
																				v1201 = int32(0)
																				v1230 = v1197
																				for {
																					v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																					v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1236 = v1230 + v1235
																					v1237 = int32(1)
																					v1240 = v1201 + v1237
																					if v1240 != v616 {
																						v1200 = v1200 + v1237
																						v1201 = v1240
																						v1230 = v1236
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1448 = v1236
																			}
																		}
																	}
																	v1455 = v626 + (base.I32_wrap_i64(v1448) ^ int32(-1))
																}
															}
															v1488 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
															v1489 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
															v1490 = v1488 - v1489
															if base.F64_lt(base.F64_convert_i32_s(v1058), base.F64_add(base.F64_convert_i32_s(v1455), base.F64_mul(base.F64_convert_i32_s(v1490*v1490*v1490), float64(-0.1)))) != 0 {
																v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
																if v1498&int32(4) != 0 {
																} else {
																	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
																	if v1501 == int32(1) {
																		if v65 == int32(0) {
																		} else {
																			base.MemoryFill(m, v546, int32(255), v65)
																		}
																	} else {
																		if v65 <= int32(0) {
																		} else {
																			v1510 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																			v1511 = int32(0)
																			if base.Ui32(int32(3)) <= base.Ui32(v624) {
																				v1517 = v1511
																				v1521 = v1511
																				for {
																					v1549 = v1517 + v546
																					v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549))))
																					v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1517+v1510))))
																					v1553 = v1550 | v1552
																					*(*uint8)(unsafe.Add(mBase, uint32(v1549))) = uint8(v1553)
																					v1556 = v1517 | int32(1)
																					v1557 = v546 + v1556
																					v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557))))
																					v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556+v1510))))
																					v1561 = v1558 | v1560
																					*(*uint8)(unsafe.Add(mBase, uint32(v1557))) = uint8(v1561)
																					v1564 = v1517 | int32(2)
																					v1565 = v546 + v1564
																					v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
																					v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564+v1510))))
																					v1569 = v1566 | v1568
																					*(*uint8)(unsafe.Add(mBase, uint32(v1565))) = uint8(v1569)
																					v1572 = v1517 | int32(3)
																					v1573 = v546 + v1572
																					v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
																					v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572+v1510))))
																					v1577 = v1574 | v1576
																					*(*uint8)(unsafe.Add(mBase, uint32(v1573))) = uint8(v1577)
																					v1579 = int32(4)
																					v1580 = v1517 + v1579
																					v1582 = v1521 + v1579
																					if v1582 != v614 {
																						v1517 = v1580
																						v1521 = v1582
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																				} else {
																					v1587 = v1580
																					v1619 = v1587
																					v1620 = v1511
																					for {
																						v1650 = v1619 + v546
																						v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
																						v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619+v1510))))
																						v1654 = v1651 | v1653
																						*(*uint8)(unsafe.Add(mBase, uint32(v1650))) = uint8(v1654)
																						v1656 = int32(1)
																						v1659 = v1620 + v1656
																						if v1659 != v616 {
																							v1619 = v1619 + v1656
																							v1620 = v1659
																							continue
																						} else {
																							break
																						}
																						break
																					}
																				}
																			} else {
																				v1587 = v1511
																				v1619 = v1587
																				v1620 = v1511
																				for {
																					v1650 = v1619 + v546
																					v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
																					v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619+v1510))))
																					v1654 = v1651 | v1653
																					*(*uint8)(unsafe.Add(mBase, uint32(v1650))) = uint8(v1654)
																					v1656 = int32(1)
																					v1659 = v1620 + v1656
																					if v1659 != v616 {
																						v1619 = v1619 + v1656
																						v1620 = v1659
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
																*(*uint16)(unsafe.Add(mBase, uint32(v646))) = uint16(v667)
																v1694 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
																*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v1694 + int32(1)
																v1947 = v645
																v1948 = v646 + int32(2)
															} else {
																v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)))
																if v1700&int32(4) != 0 {
																} else {
																	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
																	if v1703 == int32(1) {
																		if v65 == int32(0) {
																		} else {
																			base.MemoryFill(m, v544, int32(255), v65)
																		}
																	} else {
																		if v65 <= int32(0) {
																		} else {
																			v1712 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																			v1713 = int32(0)
																			if base.Ui32(int32(4)) <= base.Ui32(v65) {
																				v1719 = v1713
																				v1723 = v1713
																				for {
																					v1751 = v1719 + v544
																					v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751))))
																					v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1719+v1712))))
																					v1755 = v1752 | v1754
																					*(*uint8)(unsafe.Add(mBase, uint32(v1751))) = uint8(v1755)
																					v1758 = v1719 | int32(1)
																					v1759 = v544 + v1758
																					v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759))))
																					v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758+v1712))))
																					v1763 = v1760 | v1762
																					*(*uint8)(unsafe.Add(mBase, uint32(v1759))) = uint8(v1763)
																					v1766 = v1719 | int32(2)
																					v1767 = v544 + v1766
																					v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767))))
																					v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766+v1712))))
																					v1771 = v1768 | v1770
																					*(*uint8)(unsafe.Add(mBase, uint32(v1767))) = uint8(v1771)
																					v1774 = v1719 | int32(3)
																					v1775 = v544 + v1774
																					v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1775))))
																					v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1712))))
																					v1779 = v1776 | v1778
																					*(*uint8)(unsafe.Add(mBase, uint32(v1775))) = uint8(v1779)
																					v1781 = int32(4)
																					v1782 = v1719 + v1781
																					v1784 = v1723 + v1781
																					if v1784 != v614 {
																						v1719 = v1782
																						v1723 = v1784
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v616 == int32(0) {
																				} else {
																					v1789 = v1782
																					v1821 = v1789
																					v1822 = v1713
																					for {
																						v1852 = v1821 + v544
																						v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852))))
																						v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1821+v1712))))
																						v1856 = v1853 | v1855
																						*(*uint8)(unsafe.Add(mBase, uint32(v1852))) = uint8(v1856)
																						v1858 = int32(1)
																						v1861 = v1822 + v1858
																						if v1861 != v616 {
																							v1821 = v1821 + v1858
																							v1822 = v1861
																							continue
																						} else {
																							break
																						}
																						break
																					}
																				}
																			} else {
																				v1789 = v1713
																				v1821 = v1789
																				v1822 = v1713
																				for {
																					v1852 = v1821 + v544
																					v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852))))
																					v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1821+v1712))))
																					v1856 = v1853 | v1855
																					*(*uint8)(unsafe.Add(mBase, uint32(v1852))) = uint8(v1856)
																					v1858 = int32(1)
																					v1861 = v1822 + v1858
																					if v1861 != v616 {
																						v1821 = v1821 + v1858
																						v1822 = v1861
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
																*(*uint16)(unsafe.Add(mBase, uint32(v645))) = uint16(v667)
																v1928 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v1928 + int32(1)
																v1947 = v645 + int32(2)
																v1948 = v646
															}
														}
													}
													v1967 = v643 + int32(1)
													if v1967 != v612 {
														v643 = v1967
														v645 = v1947
														v646 = v1948
														continue
													} else {
														break
													}
													break
												}
												*(*int64)(unsafe.Add(mBase, uint32(v448)+32)) = base.I64_extend_i32_u(v513)
												*(*int64)(unsafe.Add(mBase, uint32(v448)+8)) = base.I64_extend_i32_u(v483)
												return v33 & int64(4294967295)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v65 = int32(12)
		v68 = v37 & int32(_a_F_gtrgm_picksplit_0)
		v70 = v68 + int32(1)
		v71 = F_palloc_mul(m, int32(8), v70)
		mBase = m.M
		v72 = m.ExcPending
		if v72 != 0 {
			return int64(0)
		} else {
			v74 = F_palloc(m, v70*v65)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int64(0)
			} else {
				v77 = v35 & int32(_a_F_gtrgm_picksplit_0)
				if v77 != int32(1) {
					v80 = int32(3)
					v87 = int32(8)
					v91 = int32(1)
					v92 = v65<<(uint(v80)%32) - v91
					v94 = base.I32_div_s(v92, v87)
					v97 = v91
					for {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v34+v87+v97*int32(24))))
						v134 = v71 + v97<<(uint(int32(3))%32)
						v136 = v74 + v97*v65
						*(*int32)(unsafe.Add(mBase, uint32(v134)+4)) = v136
						v138 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v138)
						v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+4)))
						if v140&int32(1) != 0 {
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
							v147 = int32(base.Ui32(v143)>>(uint(int32(2))%32)) - int32(5)
							v148 = int32(3)
							v149 = base.I32_div_u_s(v147, v148)
							v152 = int32(0)
							if base.B2i32(v65&v80 != int32(0))|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v65))|base.B2i32(v136&v148 != v152) == v152 {
								if v65 == int32(0) {
								} else {
									v161 = v65 + v136
									v163 = v136 + int32(4)
									if base.Ui32(v163) < base.Ui32(v161) {
										v165 = v161
									} else {
										v165 = v163
									}
									v171 = (v136^int32(-1)+v165)&int32(-4) + int32(4)
									if v171 == int32(0) {
									} else {
										base.MemoryFill(m, v136, int32(0), v171)
									}
								}
							} else {
								v171 = v65
								if v171 == int32(0) {
								} else {
									base.MemoryFill(m, v136, int32(0), v171)
								}
							}
							v179 = v136 + v94
							v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
							v182 = v180 | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v182)
							if base.Ui32(v147) < base.Ui32(int32(3)) {
							} else {
								v188 = int32(1)
								if base.Ui32(v149) <= base.Ui32(v188) {
									v191 = v188
								} else {
									v191 = v149
								}
								v193 = int32(0)
								for {
									v225 = int32(3)
									v227 = v131 + int32(5) + v193*v225
									v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227))))
									v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+2)))
									v233 = base.I32_rem_u_s(v228|v229<<(uint(int32(16))%32), v92)
									v236 = v136 + int32(base.Ui32(v233)>>(uint(v225)%32))
									v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
									v238 = int32(1)
									v242 = v237 | v238<<(uint(v233&int32(7))%32)
									*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v242)
									v245 = v193 + v238
									if v245 != v191 {
										v193 = v245
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							if v140&int32(4) != 0 {
								v249 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v249)
							} else {
								if v65 == int32(0) {
								} else {
									base.MemoryCopy(m, v136, v131+int32(5), v65)
								}
							}
						}
						v291 = (v97 + int32(1)) & int32(_a_F_gtrgm_picksplit_0)
						if base.Ui32(v291) <= base.Ui32(v68) {
							v97 = v291
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v325 = int32(0)
				if base.Ui32(int32(2)) <= base.Ui32(v68) {
					v332 = int32(1)
					v333 = v325
					v335 = int32(-1)
					v342 = v325
					for {
						v367 = v332 + int32(1)
						v368 = v367
						v370 = v333
						v372 = v335
						v376 = v367
						v379 = v342
						for {
							v403 = F_hemdistcache_2(m, v71+v376<<(uint(int32(3))%32), v71+v332<<(uint(int32(3))%32), v65)
							mBase = m.M
							v404 = base.B2i32(v372 < v403)
							if v372 < v403 {
								v405 = v403
							} else {
								v405 = v372
							}
							if v372 < v403 {
								v406 = v368
							} else {
								v406 = v379
							}
							if v372 < v403 {
								v407 = v332
							} else {
								v407 = v370
							}
							v409 = v368 + int32(1)
							v410 = int32(_a_F_gtrgm_picksplit_0)
							v411 = v409 & v410
							if base.Ui32(v411) <= base.Ui32(v37&v410) {
								v368 = v409
								v370 = v407
								v372 = v405
								v376 = v411
								v379 = v406
								continue
							} else {
								break
							}
							break
						}
						if v367 != v68 {
							v332 = v367
							v333 = v407
							v335 = v405
							v342 = v406
							continue
						} else {
							break
						}
						break
					}
					v418 = v407
					v427 = v406
				} else {
					v418 = v325
					v427 = v325
				}
				v448 = base.I32_wrap_i64(v33)
				v450 = v68 << (uint(int32(1)) % 32)
				v451 = F_palloc(m, v450)
				mBase = m.M
				v452 = m.ExcPending
				if v452 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v448))) = v451
					v454 = F_palloc(m, v450)
					mBase = m.M
					v455 = m.ExcPending
					if v455 != 0 {
						return int64(0)
					} else {
						v456 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v456
						*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v456
						*(*int32)(unsafe.Add(mBase, uint32(v448)+20)) = v454
						v462 = int32(_a_F_gtrgm_picksplit_0)
						v470 = base.B2i32(v418&v462 == v456) | base.B2i32(v427&v462 == v456)
						if v470 != 0 {
							v471 = int32(1)
						} else {
							v471 = v418
						}
						v476 = v71 + v471&int32(_a_F_gtrgm_picksplit_0)<<(uint(int32(3))%32)
						v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
						v478 = int32(5)
						v480 = v65 + v478
						v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
						if v481 != 0 {
							v482 = v478
						} else {
							v482 = v480
						}
						v483 = F_palloc(m, v482)
						mBase = m.M
						v484 = m.ExcPending
						if v484 != 0 {
							return int64(0)
						} else {
							if v481 != 0 {
								v487 = int32(6)
							} else {
								v487 = int32(2)
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)) = uint8(v487)
							v489 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v483))) = v482 << (uint(v489) % 32)
							if v470 != 0 {
								v493 = v489
							} else {
								v493 = v427
							}
							if v481 != 0 {
							} else {
								v495 = v483 + int32(5)
								if v477 != 0 {
									if v65 == int32(0) {
									} else {
										base.MemoryCopy(m, v495, v477, v65)
									}
								} else {
									if v65 == int32(0) {
									} else {
										base.MemoryFill(m, v495, int32(0), v65)
									}
								}
							}
							v508 = v71 + v493&int32(_a_F_gtrgm_picksplit_0)<<(uint(int32(3))%32)
							v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
							v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
							if v511 != 0 {
								v512 = int32(5)
							} else {
								v512 = v480
							}
							v513 = F_palloc(m, v512)
							mBase = m.M
							v514 = m.ExcPending
							if v514 != 0 {
								return int64(0)
							} else {
								if v511 != 0 {
									v517 = int32(6)
								} else {
									v517 = int32(2)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)) = uint8(v517)
								*(*int32)(unsafe.Add(mBase, uint32(v513))) = v512 << (uint(int32(2)) % 32)
								if v511 != 0 {
								} else {
									v523 = v513 + int32(5)
									if v509 != 0 {
										if v65 == int32(0) {
										} else {
											base.MemoryCopy(m, v523, v509, v65)
										}
									} else {
										if v65 == int32(0) {
										} else {
											base.MemoryFill(m, v523, int32(0), v65)
										}
									}
								}
								v535 = F_palloc_mul(m, int32(8), v68)
								mBase = m.M
								v536 = m.ExcPending
								if v536 != 0 {
									return int64(0)
								} else {
									if v77 == int32(1) {
										F_pg_qsort(m, v535, v68, int32(8), int32(_a_F_gtrgm_picksplit_1))
										mBase = m.M
										v542 = m.ExcPending
										if v542 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v448)+32)) = base.I64_extend_i32_u(v513)
											*(*int64)(unsafe.Add(mBase, uint32(v448)+8)) = base.I64_extend_i32_u(v483)
											return v33 & int64(4294967295)
										}
									} else {
										v543 = int32(5)
										v544 = v513 + v543
										v546 = v483 + v543
										v547 = int32(1)
										v549 = v547
										v551 = v547
										for {
											v582 = v551 << (uint(int32(3)) % 32)
											v583 = v535 + v582
											*(*uint16)(unsafe.Add(mBase, uint32(v583-int32(8)))) = uint16(v549)
											v589 = v582 + v71
											v590 = F_hemdistcache_2(m, v476, v589, v65)
											mBase = m.M
											v591 = F_hemdistcache_2(m, v508, v589, v65)
											mBase = m.M
											v592 = v590 - v591
											v594 = v592 >> (uint(int32(31)) % 32)
											*(*int32)(unsafe.Add(mBase, uint32(v583-int32(4)))) = v592 ^ v594 - v594
											v599 = v549 + int32(1)
											v600 = int32(_a_F_gtrgm_picksplit_0)
											v601 = v599 & v600
											if base.Ui32(v601) <= base.Ui32(v37&v600) {
												v549 = v599
												v551 = v601
												continue
											} else {
												break
											}
											break
										}
										F_pg_qsort(m, v535, v68, int32(8), int32(_a_F_gtrgm_picksplit_1))
										mBase = m.M
										v608 = m.ExcPending
										if v608 != 0 {
											return int64(0)
										} else {
											v609 = int32(1)
											if base.Ui32(v68) <= base.Ui32(v609) {
												v612 = v609
											} else {
												v612 = v68
											}
											v614 = v65 & int32(2147483644)
											v615 = int32(3)
											v616 = v65 & v615
											v618 = v65 & int32(-4)
											v620 = v65 & int32(2147483646)
											v621 = int32(1)
											v622 = v65 & v621
											v624 = v65 - v621
											v626 = v65 << (uint(v615) % 32)
											v631 = base.B2i32(int32(7) < v65)
											v643 = int32(0)
											v645 = v454
											v646 = v451
											for {
												v667 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v535+v643<<(uint(int32(3))%32)))))
												if v471&int32(_a_F_gtrgm_picksplit_0) == v667 {
													*(*uint16)(unsafe.Add(mBase, uint32(v646))) = uint16(v471)
													v670 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v670 + int32(1)
													v1947 = v645
													v1948 = v646 + int32(2)
												} else {
													if v493&int32(_a_F_gtrgm_picksplit_0) == v667 {
														*(*uint16)(unsafe.Add(mBase, uint32(v645))) = uint16(v493)
														v1928 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
														*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v1928 + int32(1)
														v1947 = v645 + int32(2)
														v1948 = v646
													} else {
														v682 = v71 + v667<<(uint(int32(3))%32)
														v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
														v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
														if v684&int32(4) == int32(0) {
															if v683&int32(1) != 0 {
																v698 = v546
																if int32(7) < v65 {
																	v945 = int64(0)
																	v946 = int32(0)
																	if v65 == v946 {
																		v1016 = int64(0)
																	} else {
																		v953 = v65 & int32(3)
																		if base.Ui32(int32(4)) <= base.Ui32(v65) {
																			v958 = v698
																			v960 = v945
																			v963 = v946
																			for {
																				v964 = int32(4)
																				v965 = v958 + v964
																				v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+3)))
																				v967 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v966)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+2)))
																				v969 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v968)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+1)))
																				v971 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v970)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
																				v973 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v972)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v977 = v967 + (v969 + (v971 + (v960 + v973)))
																				v979 = v963 + v964
																				if v979 != v65&int32(-4) {
																					v958 = v965
																					v960 = v977
																					v963 = v979
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v953 == int32(0) {
																				v1006 = v977
																			} else {
																				v983 = v965
																				v985 = v977
																				v990 = v983
																				v991 = int32(0)
																				v992 = v985
																				for {
																					v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																					v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v998 = v992 + v997
																					v999 = int32(1)
																					v1002 = v991 + v999
																					if v1002 != v953 {
																						v990 = v990 + v999
																						v991 = v1002
																						v992 = v998
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1006 = v998
																			}
																		} else {
																			v983 = v698
																			v985 = v945
																			v990 = v983
																			v991 = int32(0)
																			v992 = v985
																			for {
																				v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																				v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v998 = v992 + v997
																				v999 = int32(1)
																				v1002 = v991 + v999
																				if v1002 != v953 {
																					v990 = v990 + v999
																					v991 = v1002
																					v992 = v998
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1006 = v998
																		}
																		v1016 = v1006
																	}
																	v1047 = v1016
																} else {
																	if v65 == int32(0) {
																		v1047 = int64(0)
																	} else {
																		v703 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v624) {
																			v706 = v698
																			v711 = int32(0)
																			v736 = v703
																			for {
																				v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+3)))
																				v741 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v738)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+2)))
																				v745 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v742)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
																				v749 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v746)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
																				v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v750)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v757 = v741 + (v745 + (v749 + (v736 + v753)))
																				v758 = int32(4)
																				v759 = v706 + v758
																				v761 = v711 + v758
																				if v761 != v618 {
																					v706 = v759
																					v711 = v761
																					v736 = v757
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																				v1047 = v757
																			} else {
																				v765 = v759
																				v795 = v757
																				v798 = v765
																				v799 = int32(0)
																				v828 = v795
																				for {
																					v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																					v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v834 = v828 + v833
																					v835 = int32(1)
																					v838 = v799 + v835
																					if v838 != v616 {
																						v798 = v798 + v835
																						v799 = v838
																						v828 = v834
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1047 = v834
																			}
																		} else {
																			v765 = v698
																			v795 = v703
																			v798 = v765
																			v799 = int32(0)
																			v828 = v795
																			for {
																				v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																				v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v834 = v828 + v833
																				v835 = int32(1)
																				v838 = v799 + v835
																				if v838 != v616 {
																					v798 = v798 + v835
																					v799 = v838
																					v828 = v834
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1047 = v834
																		}
																	}
																}
																v1058 = v626 + (base.I32_wrap_i64(v1047) ^ int32(-1))
															} else {
																if int32(0) < v65 {
																	v840 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																	v841 = int32(0)
																	if v624 != 0 {
																		v845 = v841
																		v847 = v841
																		v850 = v841
																		for {
																			v878 = v845 | int32(1)
																			v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546+v878))))
																			v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v878))))
																			v886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v880^v882)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845+v546))))
																			v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v845+v840))))
																			v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v888^v890)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v896 = v886 + (v850 + v894)
																			v897 = int32(2)
																			v898 = v845 + v897
																			v900 = v847 + v897
																			if v900 != v620 {
																				v845 = v898
																				v847 = v900
																				v850 = v896
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v622 == int32(0) {
																			v1058 = v896
																		} else {
																			v908 = v898
																			v909 = v896
																			v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v908))))
																			v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908+v546))))
																			v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937^v939)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1058 = v909 + v943
																		}
																	} else {
																		v908 = v841
																		v909 = v841
																		v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v840+v908))))
																		v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908+v546))))
																		v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937^v939)+uint32(_c_F_gtrgm_picksplit[0]))))
																		v1058 = v909 + v943
																	}
																} else {
																	v1058 = int32(0)
																}
															}
														} else {
															if v683&int32(1) != 0 {
																v1058 = int32(0)
															} else {
																v697 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																v698 = v697
																if int32(7) < v65 {
																	v945 = int64(0)
																	v946 = int32(0)
																	if v65 == v946 {
																		v1016 = int64(0)
																	} else {
																		v953 = v65 & int32(3)
																		if base.Ui32(int32(4)) <= base.Ui32(v65) {
																			v958 = v698
																			v960 = v945
																			v963 = v946
																			for {
																				v964 = int32(4)
																				v965 = v958 + v964
																				v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+3)))
																				v967 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v966)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+2)))
																				v969 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v968)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958)+1)))
																				v971 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v970)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
																				v973 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v972)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v977 = v967 + (v969 + (v971 + (v960 + v973)))
																				v979 = v963 + v964
																				if v979 != v65&int32(-4) {
																					v958 = v965
																					v960 = v977
																					v963 = v979
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v953 == int32(0) {
																				v1006 = v977
																			} else {
																				v983 = v965
																				v985 = v977
																				v990 = v983
																				v991 = int32(0)
																				v992 = v985
																				for {
																					v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																					v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v998 = v992 + v997
																					v999 = int32(1)
																					v1002 = v991 + v999
																					if v1002 != v953 {
																						v990 = v990 + v999
																						v991 = v1002
																						v992 = v998
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1006 = v998
																			}
																		} else {
																			v983 = v698
																			v985 = v945
																			v990 = v983
																			v991 = int32(0)
																			v992 = v985
																			for {
																				v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v990))))
																				v997 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v996)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v998 = v992 + v997
																				v999 = int32(1)
																				v1002 = v991 + v999
																				if v1002 != v953 {
																					v990 = v990 + v999
																					v991 = v1002
																					v992 = v998
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1006 = v998
																		}
																		v1016 = v1006
																	}
																	v1047 = v1016
																} else {
																	if v65 == int32(0) {
																		v1047 = int64(0)
																	} else {
																		v703 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v624) {
																			v706 = v698
																			v711 = int32(0)
																			v736 = v703
																			for {
																				v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+3)))
																				v741 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v738)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+2)))
																				v745 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v742)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+1)))
																				v749 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v746)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
																				v753 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v750)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v757 = v741 + (v745 + (v749 + (v736 + v753)))
																				v758 = int32(4)
																				v759 = v706 + v758
																				v761 = v711 + v758
																				if v761 != v618 {
																					v706 = v759
																					v711 = v761
																					v736 = v757
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																				v1047 = v757
																			} else {
																				v765 = v759
																				v795 = v757
																				v798 = v765
																				v799 = int32(0)
																				v828 = v795
																				for {
																					v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																					v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v834 = v828 + v833
																					v835 = int32(1)
																					v838 = v799 + v835
																					if v838 != v616 {
																						v798 = v798 + v835
																						v799 = v838
																						v828 = v834
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1047 = v834
																			}
																		} else {
																			v765 = v698
																			v795 = v703
																			v798 = v765
																			v799 = int32(0)
																			v828 = v795
																			for {
																				v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
																				v833 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v830)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v834 = v828 + v833
																				v835 = int32(1)
																				v838 = v799 + v835
																				if v838 != v616 {
																					v798 = v798 + v835
																					v799 = v838
																					v828 = v834
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1047 = v834
																		}
																	}
																}
																v1058 = v626 + (base.I32_wrap_i64(v1047) ^ int32(-1))
															}
														}
														v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
														v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)))
														if v1086&int32(4) == int32(0) {
															if v1085&int32(1) != 0 {
																v1100 = v544
																if int32(7) < v65 {
																	v1346 = int64(0)
																	v1347 = int32(0)
																	if v65 == v1347 {
																		v1417 = int64(0)
																	} else {
																		v1354 = v65 & int32(3)
																		if base.Ui32(int32(4)) <= base.Ui32(v65) {
																			v1359 = v1100
																			v1361 = v1346
																			v1364 = v1347
																			for {
																				v1365 = int32(4)
																				v1366 = v1359 + v1365
																				v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+3)))
																				v1368 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+2)))
																				v1370 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1369)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+1)))
																				v1372 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359))))
																				v1374 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1378 = v1368 + (v1370 + (v1372 + (v1361 + v1374)))
																				v1380 = v1364 + v1365
																				if v1380 != v65&int32(-4) {
																					v1359 = v1366
																					v1361 = v1378
																					v1364 = v1380
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v1354 == int32(0) {
																				v1407 = v1378
																			} else {
																				v1384 = v1366
																				v1386 = v1378
																				v1391 = v1384
																				v1392 = int32(0)
																				v1393 = v1386
																				for {
																					v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																					v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1399 = v1393 + v1398
																					v1400 = int32(1)
																					v1403 = v1392 + v1400
																					if v1403 != v1354 {
																						v1391 = v1391 + v1400
																						v1392 = v1403
																						v1393 = v1399
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1407 = v1399
																			}
																		} else {
																			v1384 = v1100
																			v1386 = v1346
																			v1391 = v1384
																			v1392 = int32(0)
																			v1393 = v1386
																			for {
																				v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																				v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1399 = v1393 + v1398
																				v1400 = int32(1)
																				v1403 = v1392 + v1400
																				if v1403 != v1354 {
																					v1391 = v1391 + v1400
																					v1392 = v1403
																					v1393 = v1399
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1407 = v1399
																		}
																		v1417 = v1407
																	}
																	v1448 = v1417
																} else {
																	if v65 == int32(0) {
																		v1448 = int64(0)
																	} else {
																		v1105 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v624) {
																			v1108 = v1100
																			v1110 = int32(0)
																			v1138 = v1105
																			for {
																				v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+3)))
																				v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+2)))
																				v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
																				v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																				v1155 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1159 = v1143 + (v1147 + (v1151 + (v1138 + v1155)))
																				v1160 = int32(4)
																				v1161 = v1108 + v1160
																				v1163 = v1110 + v1160
																				if v1163 != v618 {
																					v1108 = v1161
																					v1110 = v1163
																					v1138 = v1159
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																				v1448 = v1159
																			} else {
																				v1167 = v1161
																				v1197 = v1159
																				v1200 = v1167
																				v1201 = int32(0)
																				v1230 = v1197
																				for {
																					v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																					v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1236 = v1230 + v1235
																					v1237 = int32(1)
																					v1240 = v1201 + v1237
																					if v1240 != v616 {
																						v1200 = v1200 + v1237
																						v1201 = v1240
																						v1230 = v1236
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1448 = v1236
																			}
																		} else {
																			v1167 = v1100
																			v1197 = v1105
																			v1200 = v1167
																			v1201 = int32(0)
																			v1230 = v1197
																			for {
																				v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																				v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1236 = v1230 + v1235
																				v1237 = int32(1)
																				v1240 = v1201 + v1237
																				if v1240 != v616 {
																					v1200 = v1200 + v1237
																					v1201 = v1240
																					v1230 = v1236
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1448 = v1236
																		}
																	}
																}
																v1455 = v626 + (base.I32_wrap_i64(v1448) ^ int32(-1))
															} else {
																if int32(0) < v65 {
																	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																	v1243 = int32(0)
																	if v624 != 0 {
																		v1246 = v1243
																		v1247 = v1243
																		v1250 = v1243
																		for {
																			v1279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v1242))))
																			v1281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1246+v544))))
																			v1285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1279^v1281)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1288 = v1246 | int32(1)
																			v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544+v1288))))
																			v1292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1288+v1242))))
																			v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1290^v1292)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1297 = v1247 + v1285 + v1296
																			v1298 = int32(2)
																			v1299 = v1246 + v1298
																			v1301 = v1250 + v1298
																			if v1301 != v620 {
																				v1246 = v1299
																				v1247 = v1297
																				v1250 = v1301
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		if v622 == int32(0) {
																			v1455 = v1297
																		} else {
																			v1305 = v1299
																			v1306 = v1297
																			v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v1242))))
																			v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v544))))
																			v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338^v1340)+uint32(_c_F_gtrgm_picksplit[0]))))
																			v1455 = v1306 + v1344
																		}
																	} else {
																		v1305 = v1243
																		v1306 = v1243
																		v1338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v1242))))
																		v1340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305+v544))))
																		v1344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338^v1340)+uint32(_c_F_gtrgm_picksplit[0]))))
																		v1455 = v1306 + v1344
																	}
																} else {
																	v1455 = int32(0)
																}
															}
														} else {
															if v1085&int32(1) != 0 {
																v1455 = int32(0)
															} else {
																v1099 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																v1100 = v1099
																if int32(7) < v65 {
																	v1346 = int64(0)
																	v1347 = int32(0)
																	if v65 == v1347 {
																		v1417 = int64(0)
																	} else {
																		v1354 = v65 & int32(3)
																		if base.Ui32(int32(4)) <= base.Ui32(v65) {
																			v1359 = v1100
																			v1361 = v1346
																			v1364 = v1347
																			for {
																				v1365 = int32(4)
																				v1366 = v1359 + v1365
																				v1367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+3)))
																				v1368 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1367)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+2)))
																				v1370 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1369)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359)+1)))
																				v1372 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1371)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1359))))
																				v1374 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1373)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1378 = v1368 + (v1370 + (v1372 + (v1361 + v1374)))
																				v1380 = v1364 + v1365
																				if v1380 != v65&int32(-4) {
																					v1359 = v1366
																					v1361 = v1378
																					v1364 = v1380
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v1354 == int32(0) {
																				v1407 = v1378
																			} else {
																				v1384 = v1366
																				v1386 = v1378
																				v1391 = v1384
																				v1392 = int32(0)
																				v1393 = v1386
																				for {
																					v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																					v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1399 = v1393 + v1398
																					v1400 = int32(1)
																					v1403 = v1392 + v1400
																					if v1403 != v1354 {
																						v1391 = v1391 + v1400
																						v1392 = v1403
																						v1393 = v1399
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1407 = v1399
																			}
																		} else {
																			v1384 = v1100
																			v1386 = v1346
																			v1391 = v1384
																			v1392 = int32(0)
																			v1393 = v1386
																			for {
																				v1397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1391))))
																				v1398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1397)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1399 = v1393 + v1398
																				v1400 = int32(1)
																				v1403 = v1392 + v1400
																				if v1403 != v1354 {
																					v1391 = v1391 + v1400
																					v1392 = v1403
																					v1393 = v1399
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1407 = v1399
																		}
																		v1417 = v1407
																	}
																	v1448 = v1417
																} else {
																	if v65 == int32(0) {
																		v1448 = int64(0)
																	} else {
																		v1105 = int64(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v624) {
																			v1108 = v1100
																			v1110 = int32(0)
																			v1138 = v1105
																			for {
																				v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+3)))
																				v1143 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1140)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+2)))
																				v1147 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1144)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+1)))
																				v1151 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1148)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108))))
																				v1155 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1152)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1159 = v1143 + (v1147 + (v1151 + (v1138 + v1155)))
																				v1160 = int32(4)
																				v1161 = v1108 + v1160
																				v1163 = v1110 + v1160
																				if v1163 != v618 {
																					v1108 = v1161
																					v1110 = v1163
																					v1138 = v1159
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																				v1448 = v1159
																			} else {
																				v1167 = v1161
																				v1197 = v1159
																				v1200 = v1167
																				v1201 = int32(0)
																				v1230 = v1197
																				for {
																					v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																					v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																					v1236 = v1230 + v1235
																					v1237 = int32(1)
																					v1240 = v1201 + v1237
																					if v1240 != v616 {
																						v1200 = v1200 + v1237
																						v1201 = v1240
																						v1230 = v1236
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v1448 = v1236
																			}
																		} else {
																			v1167 = v1100
																			v1197 = v1105
																			v1200 = v1167
																			v1201 = int32(0)
																			v1230 = v1197
																			for {
																				v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200))))
																				v1235 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1232)+uint32(_c_F_gtrgm_picksplit[0]))))
																				v1236 = v1230 + v1235
																				v1237 = int32(1)
																				v1240 = v1201 + v1237
																				if v1240 != v616 {
																					v1200 = v1200 + v1237
																					v1201 = v1240
																					v1230 = v1236
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v1448 = v1236
																		}
																	}
																}
																v1455 = v626 + (base.I32_wrap_i64(v1448) ^ int32(-1))
															}
														}
														v1488 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
														v1489 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
														v1490 = v1488 - v1489
														if base.F64_lt(base.F64_convert_i32_s(v1058), base.F64_add(base.F64_convert_i32_s(v1455), base.F64_mul(base.F64_convert_i32_s(v1490*v1490*v1490), float64(-0.1)))) != 0 {
															v1498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)))
															if v1498&int32(4) != 0 {
															} else {
																v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
																if v1501 == int32(1) {
																	if v65 == int32(0) {
																	} else {
																		base.MemoryFill(m, v546, int32(255), v65)
																	}
																} else {
																	if v65 <= int32(0) {
																	} else {
																		v1510 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																		v1511 = int32(0)
																		if base.Ui32(int32(3)) <= base.Ui32(v624) {
																			v1517 = v1511
																			v1521 = v1511
																			for {
																				v1549 = v1517 + v546
																				v1550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549))))
																				v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1517+v1510))))
																				v1553 = v1550 | v1552
																				*(*uint8)(unsafe.Add(mBase, uint32(v1549))) = uint8(v1553)
																				v1556 = v1517 | int32(1)
																				v1557 = v546 + v1556
																				v1558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1557))))
																				v1560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556+v1510))))
																				v1561 = v1558 | v1560
																				*(*uint8)(unsafe.Add(mBase, uint32(v1557))) = uint8(v1561)
																				v1564 = v1517 | int32(2)
																				v1565 = v546 + v1564
																				v1566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1565))))
																				v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564+v1510))))
																				v1569 = v1566 | v1568
																				*(*uint8)(unsafe.Add(mBase, uint32(v1565))) = uint8(v1569)
																				v1572 = v1517 | int32(3)
																				v1573 = v546 + v1572
																				v1574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1573))))
																				v1576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572+v1510))))
																				v1577 = v1574 | v1576
																				*(*uint8)(unsafe.Add(mBase, uint32(v1573))) = uint8(v1577)
																				v1579 = int32(4)
																				v1580 = v1517 + v1579
																				v1582 = v1521 + v1579
																				if v1582 != v614 {
																					v1517 = v1580
																					v1521 = v1582
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																			} else {
																				v1587 = v1580
																				v1619 = v1587
																				v1620 = v1511
																				for {
																					v1650 = v1619 + v546
																					v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
																					v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619+v1510))))
																					v1654 = v1651 | v1653
																					*(*uint8)(unsafe.Add(mBase, uint32(v1650))) = uint8(v1654)
																					v1656 = int32(1)
																					v1659 = v1620 + v1656
																					if v1659 != v616 {
																						v1619 = v1619 + v1656
																						v1620 = v1659
																						continue
																					} else {
																						break
																					}
																					break
																				}
																			}
																		} else {
																			v1587 = v1511
																			v1619 = v1587
																			v1620 = v1511
																			for {
																				v1650 = v1619 + v546
																				v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
																				v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1619+v1510))))
																				v1654 = v1651 | v1653
																				*(*uint8)(unsafe.Add(mBase, uint32(v1650))) = uint8(v1654)
																				v1656 = int32(1)
																				v1659 = v1620 + v1656
																				if v1659 != v616 {
																					v1619 = v1619 + v1656
																					v1620 = v1659
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
															*(*uint16)(unsafe.Add(mBase, uint32(v646))) = uint16(v667)
															v1694 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
															*(*int32)(unsafe.Add(mBase, uint32(v448)+4)) = v1694 + int32(1)
															v1947 = v645
															v1948 = v646 + int32(2)
														} else {
															v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+4)))
															if v1700&int32(4) != 0 {
															} else {
																v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v682))))
																if v1703 == int32(1) {
																	if v65 == int32(0) {
																	} else {
																		base.MemoryFill(m, v544, int32(255), v65)
																	}
																} else {
																	if v65 <= int32(0) {
																	} else {
																		v1712 = *(*int32)(unsafe.Add(mBase, uint32(v682)+4))
																		v1713 = int32(0)
																		if base.Ui32(int32(4)) <= base.Ui32(v65) {
																			v1719 = v1713
																			v1723 = v1713
																			for {
																				v1751 = v1719 + v544
																				v1752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1751))))
																				v1754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1719+v1712))))
																				v1755 = v1752 | v1754
																				*(*uint8)(unsafe.Add(mBase, uint32(v1751))) = uint8(v1755)
																				v1758 = v1719 | int32(1)
																				v1759 = v544 + v1758
																				v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759))))
																				v1762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1758+v1712))))
																				v1763 = v1760 | v1762
																				*(*uint8)(unsafe.Add(mBase, uint32(v1759))) = uint8(v1763)
																				v1766 = v1719 | int32(2)
																				v1767 = v544 + v1766
																				v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1767))))
																				v1770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1766+v1712))))
																				v1771 = v1768 | v1770
																				*(*uint8)(unsafe.Add(mBase, uint32(v1767))) = uint8(v1771)
																				v1774 = v1719 | int32(3)
																				v1775 = v544 + v1774
																				v1776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1775))))
																				v1778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774+v1712))))
																				v1779 = v1776 | v1778
																				*(*uint8)(unsafe.Add(mBase, uint32(v1775))) = uint8(v1779)
																				v1781 = int32(4)
																				v1782 = v1719 + v1781
																				v1784 = v1723 + v1781
																				if v1784 != v614 {
																					v1719 = v1782
																					v1723 = v1784
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v616 == int32(0) {
																			} else {
																				v1789 = v1782
																				v1821 = v1789
																				v1822 = v1713
																				for {
																					v1852 = v1821 + v544
																					v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852))))
																					v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1821+v1712))))
																					v1856 = v1853 | v1855
																					*(*uint8)(unsafe.Add(mBase, uint32(v1852))) = uint8(v1856)
																					v1858 = int32(1)
																					v1861 = v1822 + v1858
																					if v1861 != v616 {
																						v1821 = v1821 + v1858
																						v1822 = v1861
																						continue
																					} else {
																						break
																					}
																					break
																				}
																			}
																		} else {
																			v1789 = v1713
																			v1821 = v1789
																			v1822 = v1713
																			for {
																				v1852 = v1821 + v544
																				v1853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1852))))
																				v1855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1821+v1712))))
																				v1856 = v1853 | v1855
																				*(*uint8)(unsafe.Add(mBase, uint32(v1852))) = uint8(v1856)
																				v1858 = int32(1)
																				v1861 = v1822 + v1858
																				if v1861 != v616 {
																					v1821 = v1821 + v1858
																					v1822 = v1861
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
															*(*uint16)(unsafe.Add(mBase, uint32(v645))) = uint16(v667)
															v1928 = *(*int32)(unsafe.Add(mBase, uint32(v448)+24))
															*(*int32)(unsafe.Add(mBase, uint32(v448)+24)) = v1928 + int32(1)
															v1947 = v645 + int32(2)
															v1948 = v646
														}
													}
												}
												v1967 = v643 + int32(1)
												if v1967 != v612 {
													v643 = v1967
													v645 = v1947
													v646 = v1948
													continue
												} else {
													break
												}
												break
											}
											*(*int64)(unsafe.Add(mBase, uint32(v448)+32)) = base.I64_extend_i32_u(v513)
											*(*int64)(unsafe.Add(mBase, uint32(v448)+8)) = base.I64_extend_i32_u(v483)
											return v33 & int64(4294967295)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
