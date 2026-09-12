package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(363560)
			F_errmsg(m, int32(192260), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(492127), int32(102), int32(279360))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
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
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
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
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v731 int32
	_ = v731
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v955 int32
	_ = v955
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1259 int32
	_ = v1259
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1344 int32
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1484 int32
	_ = v1484
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1538 int32
	_ = v1538
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1609 int32
	_ = v1609
	var v1626 int32
	_ = v1626
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1669 int32
	_ = v1669
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1693 int32
	_ = v1693
	var v1707 int32
	_ = v1707
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1752 int32
	_ = v1752
	var v1778 int32
	_ = v1778
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1838 int32
	_ = v1838
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1893 int32
	_ = v1893
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1919 int32
	_ = v1919
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
			v59 = (v26 + int32(65534)) & int32(65535)
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
						v81 = int32(-1)
						v84 = v2
						v85 = v2
						v86 = int32(1)
						for {
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
							v105 = v86 + int32(1)
							v106 = v105
							v110 = v105
							v111 = v81
							v114 = v84
							v115 = v85
							for {
								v130 = int32(4)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v110<<(uint(v130)%32))))
								v134 = int32(0)
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
								v142 = v140 & v130
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
								if v143&v130 != 0 {
									if v142 != 0 {
										v357 = int32(0)
									} else {
										v148 = v57 << (uint(int32(3)) % 32)
										if v57 <= int32(0) {
											v357 = v148
										} else {
											v154 = v133 + int32(8)
											v157 = int32(0)
											v159 = v134
											for {
												v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
												v164 = int32(1)
												v199 = v163&v164 + v157 + int32(base.Ui32(v163)>>(uint(int32(7))%32)) + int32(base.Ui32(v163)>>(uint(v164)%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(2))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(3))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(4))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(5))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(6))%32))&v164
												v203 = v159 + v164
												if v203 != v57 {
													v154 = v154 + v164
													v157 = v199
													v159 = v203
													continue
												} else {
													break
												}
												break
											}
											v357 = v148 - v199
										}
									}
								} else {
									if v142 != 0 {
										v207 = v57 << (uint(int32(3)) % 32)
										if v57 <= int32(0) {
											v357 = v207
										} else {
											v213 = v103 + int32(8)
											v216 = int32(0)
											v218 = v134
											for {
												v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
												v223 = int32(1)
												v258 = v222&v223 + v216 + int32(base.Ui32(v222)>>(uint(int32(7))%32)) + int32(base.Ui32(v222)>>(uint(v223)%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(2))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(3))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(4))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(5))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(6))%32))&v223
												v262 = v218 + v223
												if v262 != v57 {
													v213 = v213 + v223
													v216 = v258
													v218 = v262
													continue
												} else {
													break
												}
												break
											}
											v357 = v207 - v258
										}
									} else {
										if v57 <= int32(0) {
											v357 = int32(0)
										} else {
											v268 = int32(8)
											v269 = v133 + v268
											v271 = v103 + v268
											v272 = int32(1)
											v274 = v57 << (uint(int32(3)) % 32)
											if v274 <= v272 {
												v277 = v272
											} else {
												v277 = v274
											}
											v278 = int32(1)
											if v277 == v278 {
												v282 = int32(0)
												v323 = v282
												v324 = v282
											} else {
												v286 = int32(0)
												v289 = v286
												v290 = v286
												v291 = v286
												for {
													v299 = int32(base.Ui32(v290) >> (uint(int32(3)) % 32))
													v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v299))))
													v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v271))))
													v305 = base.I32_extend8_s(v301 ^ v303)
													v307 = v290 & int32(6)
													v308 = int32(1)
													v317 = int32(base.Ui32(v305)>>(uint(v307|v308)%32))&v308 + (int32(base.Ui32(v305)>>(uint(v307)%32))&v308 + v289)
													v318 = int32(2)
													v319 = v290 + v318
													v321 = v291 + v318
													if v321 != v277&int32(2147483640) {
														v289 = v317
														v290 = v319
														v291 = v321
														continue
													} else {
														break
													}
													break
												}
												v323 = v317
												v324 = v319
											}
											if v277&v278 != 0 {
												v333 = int32(base.Ui32(v324) >> (uint(int32(3)) % 32))
												v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v333))))
												v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333+v271))))
												v347 = int32(base.Ui32(base.I32_extend8_s(v335^v337))>>(uint(v324&int32(7))%32))&int32(1) + v323
											} else {
												v347 = v323
											}
											v357 = v347
										}
									}
								}
								v358 = base.B2i32(v111 < v357)
								if v111 < v357 {
									v359 = v357
								} else {
									v359 = v111
								}
								if v111 < v357 {
									v360 = v106
								} else {
									v360 = v115
								}
								if v111 < v357 {
									v361 = v86
								} else {
									v361 = v114
								}
								v363 = v106 + int32(1)
								v365 = v363 & int32(65535)
								if base.Ui32(v365) <= base.Ui32(v59) {
									v106 = v363
									v110 = v365
									v111 = v359
									v114 = v361
									v115 = v360
									continue
								} else {
									break
								}
								break
							}
							if v59 != v105 {
								v81 = v359
								v84 = v361
								v85 = v360
								v86 = v105
								continue
							} else {
								break
							}
							break
						}
						v376 = v361
						v377 = v360
					} else {
						v376 = v2
						v377 = v2
					}
					v392 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v392
					*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v392
					v396 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
					v397 = int32(8)
					v399 = v57 + v397
					v401 = v25 + int32(4)
					v403 = int32(65535)
					v411 = base.B2i32(v376&v403 == v392) | base.B2i32(v377&v403 == v392)
					if v411 != 0 {
						v412 = int32(1)
					} else {
						v412 = v376
					}
					v415 = int32(4)
					v418 = *(*int32)(unsafe.Add(mBase, uint32(v401+v412&int32(65535)<<(uint(v415)%32))))
					v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
					v421 = v419 & v415
					if v421 != 0 {
						v422 = v397
					} else {
						v422 = v399
					}
					v423 = F_palloc(m, v422)
					mBase = m.M
					v424 = m.ExcPending
					if v424 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v421
						*(*int32)(unsafe.Add(mBase, uint32(v423))) = v422 << (uint(int32(2)) % 32)
						if v421 == int32(0) {
							v431 = int32(8)
							if v57 != 0 {
								v435 = F__emscripten_memcpy_bulkmem(m, v423+v431, v418+v431, v57)
								mBase = m.M
							} else {
							}
						} else {
						}
						if v411 != 0 {
							v439 = int32(2)
						} else {
							v439 = v377
						}
						v442 = int32(4)
						v445 = *(*int32)(unsafe.Add(mBase, uint32(v401+v439&int32(65535)<<(uint(v442)%32))))
						v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
						v448 = v446 & v442
						if v448 != 0 {
							v449 = int32(8)
						} else {
							v449 = v399
						}
						v450 = F_palloc(m, v449)
						mBase = m.M
						v451 = m.ExcPending
						if v451 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v450)+4)) = v448
							*(*int32)(unsafe.Add(mBase, uint32(v450))) = v449 << (uint(int32(2)) % 32)
							if v448 == int32(0) {
								v458 = int32(8)
								if v57 != 0 {
									v462 = F__emscripten_memcpy_bulkmem(m, v450+v458, v445+v458, v57)
									mBase = m.M
								} else {
								}
							} else {
							}
							v464 = int32(65535)
							v465 = v26 + v464
							v467 = v465 & v464
							v470 = F_palloc(m, v467<<(uint(int32(3))%32))
							mBase = m.M
							v471 = m.ExcPending
							if v471 != 0 {
								return int32(0)
							} else {
								if v26&int32(65535) == int32(1) {
									F_pg_qsort(m, v470, v467, int32(8), int32(6879))
									mBase = m.M
									v479 = m.ExcPending
									if v479 != 0 {
										return int32(0)
									} else {
										v1907 = v396
										v1911 = v67
										v1919 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1907))) = uint16(v1919)
										*(*uint16)(unsafe.Add(mBase, uint32(v1911))) = uint16(v1919)
										*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v450
										*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v423
										return v29
									}
								} else {
									v480 = int32(1)
									v482 = v480
									v487 = v480
									for {
										v508 = v470 + v482<<(uint(int32(3))%32)
										*(*uint16)(unsafe.Add(mBase, uint32(v508-int32(8)))) = uint16(v487)
										v512 = int32(4)
										v517 = *(*int32)(unsafe.Add(mBase, uint32(v401+v482<<(uint(v512)%32))))
										v518 = int32(0)
										v524 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
										v526 = v524 & v512
										v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
										if v527&v512 != 0 {
											if v526 != 0 {
												v741 = int32(0)
											} else {
												v532 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v741 = v532
												} else {
													v538 = v517 + int32(8)
													v541 = int32(0)
													v543 = v518
													for {
														v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
														v548 = int32(1)
														v583 = v547&v548 + v541 + int32(base.Ui32(v547)>>(uint(int32(7))%32)) + int32(base.Ui32(v547)>>(uint(v548)%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(2))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(3))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(4))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(5))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(6))%32))&v548
														v587 = v543 + v548
														if v587 != v57 {
															v538 = v538 + v548
															v541 = v583
															v543 = v587
															continue
														} else {
															break
														}
														break
													}
													v741 = v532 - v583
												}
											}
										} else {
											if v526 != 0 {
												v591 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v741 = v591
												} else {
													v597 = v423 + int32(8)
													v600 = int32(0)
													v602 = v518
													for {
														v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
														v607 = int32(1)
														v642 = v606&v607 + v600 + int32(base.Ui32(v606)>>(uint(int32(7))%32)) + int32(base.Ui32(v606)>>(uint(v607)%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(2))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(3))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(4))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(5))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(6))%32))&v607
														v646 = v602 + v607
														if v646 != v57 {
															v597 = v597 + v607
															v600 = v642
															v602 = v646
															continue
														} else {
															break
														}
														break
													}
													v741 = v591 - v642
												}
											} else {
												if v57 <= int32(0) {
													v741 = int32(0)
												} else {
													v652 = int32(8)
													v653 = v517 + v652
													v655 = v423 + v652
													v656 = int32(1)
													v658 = v57 << (uint(int32(3)) % 32)
													if v658 <= v656 {
														v661 = v656
													} else {
														v661 = v658
													}
													v662 = int32(1)
													if v661 == v662 {
														v666 = int32(0)
														v707 = v666
														v708 = v666
													} else {
														v670 = int32(0)
														v673 = v670
														v674 = v670
														v675 = v670
														for {
															v683 = int32(base.Ui32(v674) >> (uint(int32(3)) % 32))
															v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v683))))
															v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v655))))
															v689 = base.I32_extend8_s(v685 ^ v687)
															v691 = v674 & int32(6)
															v692 = int32(1)
															v701 = int32(base.Ui32(v689)>>(uint(v691|v692)%32))&v692 + (int32(base.Ui32(v689)>>(uint(v691)%32))&v692 + v673)
															v702 = int32(2)
															v703 = v674 + v702
															v705 = v675 + v702
															if v705 != v661&int32(2147483640) {
																v673 = v701
																v674 = v703
																v675 = v705
																continue
															} else {
																break
															}
															break
														}
														v707 = v701
														v708 = v703
													}
													if v661&v662 != 0 {
														v717 = int32(base.Ui32(v708) >> (uint(int32(3)) % 32))
														v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v717))))
														v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717+v655))))
														v731 = int32(base.Ui32(base.I32_extend8_s(v719^v721))>>(uint(v708&int32(7))%32))&int32(1) + v707
													} else {
														v731 = v707
													}
													v741 = v731
												}
											}
										}
										v742 = int32(0)
										v748 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
										v749 = int32(4)
										v750 = v748 & v749
										v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
										if v751&v749 != 0 {
											if v750 != 0 {
												v965 = int32(0)
											} else {
												v756 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v965 = v756
												} else {
													v762 = v517 + int32(8)
													v765 = int32(0)
													v767 = v742
													for {
														v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
														v772 = int32(1)
														v807 = v771&v772 + v765 + int32(base.Ui32(v771)>>(uint(int32(7))%32)) + int32(base.Ui32(v771)>>(uint(v772)%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(2))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(3))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(4))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(5))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(6))%32))&v772
														v811 = v767 + v772
														if v811 != v57 {
															v762 = v762 + v772
															v765 = v807
															v767 = v811
															continue
														} else {
															break
														}
														break
													}
													v965 = v756 - v807
												}
											}
										} else {
											if v750 != 0 {
												v815 = v57 << (uint(int32(3)) % 32)
												if v57 <= int32(0) {
													v965 = v815
												} else {
													v821 = v450 + int32(8)
													v824 = int32(0)
													v826 = v742
													for {
														v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
														v831 = int32(1)
														v866 = v830&v831 + v824 + int32(base.Ui32(v830)>>(uint(int32(7))%32)) + int32(base.Ui32(v830)>>(uint(v831)%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(2))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(3))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(4))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(5))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(6))%32))&v831
														v870 = v826 + v831
														if v870 != v57 {
															v821 = v821 + v831
															v824 = v866
															v826 = v870
															continue
														} else {
															break
														}
														break
													}
													v965 = v815 - v866
												}
											} else {
												if v57 <= int32(0) {
													v965 = int32(0)
												} else {
													v876 = int32(8)
													v877 = v517 + v876
													v879 = v450 + v876
													v880 = int32(1)
													v882 = v57 << (uint(int32(3)) % 32)
													if v882 <= v880 {
														v885 = v880
													} else {
														v885 = v882
													}
													v886 = int32(1)
													if v885 == v886 {
														v890 = int32(0)
														v931 = v890
														v932 = v890
													} else {
														v894 = int32(0)
														v897 = v894
														v898 = v894
														v899 = v894
														for {
															v907 = int32(base.Ui32(v898) >> (uint(int32(3)) % 32))
															v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877+v907))))
															v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907+v879))))
															v913 = base.I32_extend8_s(v909 ^ v911)
															v915 = v898 & int32(6)
															v916 = int32(1)
															v925 = int32(base.Ui32(v913)>>(uint(v915|v916)%32))&v916 + (int32(base.Ui32(v913)>>(uint(v915)%32))&v916 + v897)
															v926 = int32(2)
															v927 = v898 + v926
															v929 = v899 + v926
															if v929 != v885&int32(2147483640) {
																v897 = v925
																v898 = v927
																v899 = v929
																continue
															} else {
																break
															}
															break
														}
														v931 = v925
														v932 = v927
													}
													if v885&v886 != 0 {
														v941 = int32(base.Ui32(v932) >> (uint(int32(3)) % 32))
														v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877+v941))))
														v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941+v879))))
														v955 = int32(base.Ui32(base.I32_extend8_s(v943^v945))>>(uint(v932&int32(7))%32))&int32(1) + v931
													} else {
														v955 = v931
													}
													v965 = v955
												}
											}
										}
										v966 = v741 - v965
										v968 = v966 >> (uint(int32(31)) % 32)
										*(*int32)(unsafe.Add(mBase, uint32(v508-v512))) = v966 ^ v968 - v968
										v973 = v487 + int32(1)
										v974 = int32(65535)
										v975 = v973 & v974
										if base.Ui32(v975) <= base.Ui32(v465&v974) {
											v482 = v975
											v487 = v973
											continue
										} else {
											break
										}
										break
									}
									F_pg_qsort(m, v470, v467, int32(8), int32(6879))
									mBase = m.M
									v982 = m.ExcPending
									if v982 != 0 {
										return int32(0)
									} else {
										v983 = int32(1)
										if base.Ui32(v467) <= base.Ui32(v983) {
											v986 = v983
										} else {
											v986 = v467
										}
										v988 = v57 & int32(2147483644)
										v990 = v57 & int32(3)
										v991 = int32(8)
										v992 = v450 + v991
										v994 = v423 + v991
										v1003 = int32(0)
										v1010 = v396
										v1014 = v67
										for {
											v1025 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470+v1003<<(uint(int32(3))%32)))))
											if v412&int32(65535) == v1025 {
												*(*uint16)(unsafe.Add(mBase, uint32(v1010))) = uint16(v412)
												v1028 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1028 + int32(1)
												v1880 = v1010 + int32(2)
												v1884 = v1014
											} else {
												if v439&int32(65535) == v1025 {
													*(*uint16)(unsafe.Add(mBase, uint32(v1014))) = uint16(v439)
													v1038 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1038 + int32(1)
													v1880 = v1010
													v1884 = v1014 + int32(2)
												} else {
													v1042 = int32(4)
													v1045 = *(*int32)(unsafe.Add(mBase, uint32(v401+v1025<<(uint(v1042)%32))))
													v1046 = int32(0)
													v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
													v1054 = v1052 & v1042
													v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
													if v1055&v1042 != 0 {
														if v1054 != 0 {
															v1269 = int32(0)
														} else {
															v1060 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1269 = v1060
															} else {
																v1066 = v1045 + int32(8)
																v1069 = int32(0)
																v1071 = v1046
																for {
																	v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066))))
																	v1076 = int32(1)
																	v1111 = v1075&v1076 + v1069 + int32(base.Ui32(v1075)>>(uint(int32(7))%32)) + int32(base.Ui32(v1075)>>(uint(v1076)%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(2))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(3))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(4))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(5))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(6))%32))&v1076
																	v1115 = v1071 + v1076
																	if v1115 != v57 {
																		v1066 = v1066 + v1076
																		v1069 = v1111
																		v1071 = v1115
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1269 = v1060 - v1111
															}
														}
													} else {
														if v1054 != 0 {
															v1119 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1269 = v1119
															} else {
																v1125 = v423 + int32(8)
																v1128 = int32(0)
																v1130 = v1046
																for {
																	v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125))))
																	v1135 = int32(1)
																	v1170 = v1134&v1135 + v1128 + int32(base.Ui32(v1134)>>(uint(int32(7))%32)) + int32(base.Ui32(v1134)>>(uint(v1135)%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(2))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(3))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(4))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(5))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(6))%32))&v1135
																	v1174 = v1130 + v1135
																	if v1174 != v57 {
																		v1125 = v1125 + v1135
																		v1128 = v1170
																		v1130 = v1174
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1269 = v1119 - v1170
															}
														} else {
															if v57 <= int32(0) {
																v1269 = int32(0)
															} else {
																v1180 = int32(8)
																v1181 = v1045 + v1180
																v1183 = v423 + v1180
																v1184 = int32(1)
																v1186 = v57 << (uint(int32(3)) % 32)
																if v1186 <= v1184 {
																	v1189 = v1184
																} else {
																	v1189 = v1186
																}
																v1190 = int32(1)
																if v1189 == v1190 {
																	v1194 = int32(0)
																	v1235 = v1194
																	v1236 = v1194
																} else {
																	v1198 = int32(0)
																	v1201 = v1198
																	v1202 = v1198
																	v1203 = v1198
																	for {
																		v1211 = int32(base.Ui32(v1202) >> (uint(int32(3)) % 32))
																		v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1211))))
																		v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211+v1183))))
																		v1217 = base.I32_extend8_s(v1213 ^ v1215)
																		v1219 = v1202 & int32(6)
																		v1220 = int32(1)
																		v1229 = int32(base.Ui32(v1217)>>(uint(v1219|v1220)%32))&v1220 + (int32(base.Ui32(v1217)>>(uint(v1219)%32))&v1220 + v1201)
																		v1230 = int32(2)
																		v1231 = v1202 + v1230
																		v1233 = v1203 + v1230
																		if v1233 != v1189&int32(2147483640) {
																			v1201 = v1229
																			v1202 = v1231
																			v1203 = v1233
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v1235 = v1229
																	v1236 = v1231
																}
																if v1189&v1190 != 0 {
																	v1245 = int32(base.Ui32(v1236) >> (uint(int32(3)) % 32))
																	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1245))))
																	v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245+v1183))))
																	v1259 = int32(base.Ui32(base.I32_extend8_s(v1247^v1249))>>(uint(v1236&int32(7))%32))&int32(1) + v1235
																} else {
																	v1259 = v1235
																}
																v1269 = v1259
															}
														}
													}
													v1271 = int32(0)
													v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
													v1278 = int32(4)
													v1279 = v1277 & v1278
													v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
													if v1280&v1278 != 0 {
														if v1279 != 0 {
															v1494 = int32(0)
														} else {
															v1285 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1494 = v1285
															} else {
																v1291 = v1045 + int32(8)
																v1294 = int32(0)
																v1296 = v1271
																for {
																	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291))))
																	v1301 = int32(1)
																	v1336 = v1300&v1301 + v1294 + int32(base.Ui32(v1300)>>(uint(int32(7))%32)) + int32(base.Ui32(v1300)>>(uint(v1301)%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(2))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(3))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(4))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(5))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(6))%32))&v1301
																	v1340 = v1296 + v1301
																	if v1340 != v57 {
																		v1291 = v1291 + v1301
																		v1294 = v1336
																		v1296 = v1340
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1494 = v1285 - v1336
															}
														}
													} else {
														if v1279 != 0 {
															v1344 = v57 << (uint(int32(3)) % 32)
															if v57 <= int32(0) {
																v1494 = v1344
															} else {
																v1350 = v450 + int32(8)
																v1353 = int32(0)
																v1355 = v1271
																for {
																	v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350))))
																	v1360 = int32(1)
																	v1395 = v1359&v1360 + v1353 + int32(base.Ui32(v1359)>>(uint(int32(7))%32)) + int32(base.Ui32(v1359)>>(uint(v1360)%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(2))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(3))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(4))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(5))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(6))%32))&v1360
																	v1399 = v1355 + v1360
																	if v1399 != v57 {
																		v1350 = v1350 + v1360
																		v1353 = v1395
																		v1355 = v1399
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1494 = v1344 - v1395
															}
														} else {
															if v57 <= int32(0) {
																v1494 = int32(0)
															} else {
																v1405 = int32(8)
																v1406 = v1045 + v1405
																v1408 = v450 + v1405
																v1409 = int32(1)
																v1411 = v57 << (uint(int32(3)) % 32)
																if v1411 <= v1409 {
																	v1414 = v1409
																} else {
																	v1414 = v1411
																}
																v1415 = int32(1)
																if v1414 == v1415 {
																	v1419 = int32(0)
																	v1460 = v1419
																	v1461 = v1419
																} else {
																	v1423 = int32(0)
																	v1426 = v1423
																	v1427 = v1423
																	v1428 = v1423
																	for {
																		v1436 = int32(base.Ui32(v1427) >> (uint(int32(3)) % 32))
																		v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1436))))
																		v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1408))))
																		v1442 = base.I32_extend8_s(v1438 ^ v1440)
																		v1444 = v1427 & int32(6)
																		v1445 = int32(1)
																		v1454 = int32(base.Ui32(v1442)>>(uint(v1444|v1445)%32))&v1445 + (int32(base.Ui32(v1442)>>(uint(v1444)%32))&v1445 + v1426)
																		v1455 = int32(2)
																		v1456 = v1427 + v1455
																		v1458 = v1428 + v1455
																		if v1458 != v1414&int32(2147483640) {
																			v1426 = v1454
																			v1427 = v1456
																			v1428 = v1458
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v1460 = v1454
																	v1461 = v1456
																}
																if v1414&v1415 != 0 {
																	v1470 = int32(base.Ui32(v1461) >> (uint(int32(3)) % 32))
																	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1470))))
																	v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470+v1408))))
																	v1484 = int32(base.Ui32(base.I32_extend8_s(v1472^v1474))>>(uint(v1461&int32(7))%32))&int32(1) + v1460
																} else {
																	v1484 = v1460
																}
																v1494 = v1484
															}
														}
													}
													v1496 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													v1497 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													v1498 = v1496 - v1497
													if base.F64_lt(base.F64_convert_i32_s(v1269), base.F64_add(base.F64_convert_i32_s(v1494), base.F64_mul(base.F64_convert_i32_s(v1498*v1498*v1498), float64(-0.0001)))) != 0 {
														v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
														if v1506&int32(4) != 0 {
														} else {
															v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
															if v1509&int32(4) != 0 {
																v1514 = F__emscripten_memset_bulkmem(m, v994, base.I32_extend8_s(int32(255)), v57)
																mBase = m.M
															} else {
																if v57 <= int32(0) {
																} else {
																	v1518 = v1045 + int32(8)
																	v1519 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v57) {
																		v1524 = v1519
																		v1538 = v1519
																		for {
																			v1548 = v1524 + v994
																			v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548))))
																			v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1518))))
																			v1552 = v1549 | v1551
																			*(*uint8)(unsafe.Add(mBase, uint32(v1548))) = uint8(v1552)
																			v1555 = v1524 | int32(1)
																			v1556 = v994 + v1555
																			v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556))))
																			v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1555))))
																			v1560 = v1557 | v1559
																			*(*uint8)(unsafe.Add(mBase, uint32(v1556))) = uint8(v1560)
																			v1563 = v1524 | int32(2)
																			v1564 = v994 + v1563
																			v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564))))
																			v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1563))))
																			v1568 = v1565 | v1567
																			*(*uint8)(unsafe.Add(mBase, uint32(v1564))) = uint8(v1568)
																			v1571 = v1524 | int32(3)
																			v1572 = v994 + v1571
																			v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572))))
																			v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1571))))
																			v1576 = v1573 | v1575
																			*(*uint8)(unsafe.Add(mBase, uint32(v1572))) = uint8(v1576)
																			v1578 = int32(4)
																			v1579 = v1524 + v1578
																			v1581 = v1538 + v1578
																			if v1581 != v988 {
																				v1524 = v1579
																				v1538 = v1581
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v1583 = v1579
																	} else {
																		v1583 = v1519
																	}
																	if v990 == int32(0) {
																	} else {
																		v1609 = v1583
																		v1626 = v1519
																		for {
																			v1633 = v1609 + v994
																			v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633))))
																			v1636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1518))))
																			v1637 = v1634 | v1636
																			*(*uint8)(unsafe.Add(mBase, uint32(v1633))) = uint8(v1637)
																			v1639 = int32(1)
																			v1642 = v1626 + v1639
																			if v1642 != v990 {
																				v1609 = v1609 + v1639
																				v1626 = v1642
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
														*(*uint16)(unsafe.Add(mBase, uint32(v1010))) = uint16(v1025)
														v1669 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1669 + int32(1)
														v1880 = v1010 + int32(2)
														v1884 = v1014
													} else {
														v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
														if v1675&int32(4) != 0 {
														} else {
															v1678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
															if v1678&int32(4) != 0 {
																v1683 = F__emscripten_memset_bulkmem(m, v992, base.I32_extend8_s(int32(255)), v57)
																mBase = m.M
															} else {
																if v57 <= int32(0) {
																} else {
																	v1687 = v1045 + int32(8)
																	v1688 = int32(0)
																	if base.Ui32(int32(4)) <= base.Ui32(v57) {
																		v1693 = v1688
																		v1707 = v1688
																		for {
																			v1717 = v1693 + v992
																			v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1717))))
																			v1720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693+v1687))))
																			v1721 = v1718 | v1720
																			*(*uint8)(unsafe.Add(mBase, uint32(v1717))) = uint8(v1721)
																			v1724 = v1693 | int32(1)
																			v1725 = v992 + v1724
																			v1726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1725))))
																			v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1724))))
																			v1729 = v1726 | v1728
																			*(*uint8)(unsafe.Add(mBase, uint32(v1725))) = uint8(v1729)
																			v1732 = v1693 | int32(2)
																			v1733 = v992 + v1732
																			v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1733))))
																			v1736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1732))))
																			v1737 = v1734 | v1736
																			*(*uint8)(unsafe.Add(mBase, uint32(v1733))) = uint8(v1737)
																			v1740 = v1693 | int32(3)
																			v1741 = v992 + v1740
																			v1742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1741))))
																			v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1740))))
																			v1745 = v1742 | v1744
																			*(*uint8)(unsafe.Add(mBase, uint32(v1741))) = uint8(v1745)
																			v1747 = int32(4)
																			v1748 = v1693 + v1747
																			v1750 = v1707 + v1747
																			if v1750 != v988 {
																				v1693 = v1748
																				v1707 = v1750
																				continue
																			} else {
																				break
																			}
																			break
																		}
																		v1752 = v1748
																	} else {
																		v1752 = v1688
																	}
																	if v990 == int32(0) {
																	} else {
																		v1778 = v1752
																		v1795 = v1688
																		for {
																			v1802 = v1778 + v992
																			v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1802))))
																			v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1778+v1687))))
																			v1806 = v1803 | v1805
																			*(*uint8)(unsafe.Add(mBase, uint32(v1802))) = uint8(v1806)
																			v1808 = int32(1)
																			v1811 = v1795 + v1808
																			if v1811 != v990 {
																				v1778 = v1778 + v1808
																				v1795 = v1811
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
														*(*uint16)(unsafe.Add(mBase, uint32(v1014))) = uint16(v1025)
														v1838 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
														*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1838 + int32(1)
														v1880 = v1010
														v1884 = v1014 + int32(2)
													}
												}
											}
											v1893 = v1003 + int32(1)
											if v1893 != v986 {
												v1003 = v1893
												v1010 = v1880
												v1014 = v1884
												continue
											} else {
												break
											}
											break
										}
										v1907 = v1880
										v1911 = v1884
										v1919 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v1907))) = uint16(v1919)
										*(*uint16)(unsafe.Add(mBase, uint32(v1911))) = uint16(v1919)
										*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v450
										*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v423
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
		v59 = (v26 + int32(65534)) & int32(65535)
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
					v81 = int32(-1)
					v84 = v2
					v85 = v2
					v86 = int32(1)
					for {
						v103 = *(*int32)(unsafe.Add(mBase, uint32(v73+v86<<(uint(int32(4))%32))))
						v105 = v86 + int32(1)
						v106 = v105
						v110 = v105
						v111 = v81
						v114 = v84
						v115 = v85
						for {
							v130 = int32(4)
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v73+v110<<(uint(v130)%32))))
							v134 = int32(0)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
							v142 = v140 & v130
							v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
							if v143&v130 != 0 {
								if v142 != 0 {
									v357 = int32(0)
								} else {
									v148 = v57 << (uint(int32(3)) % 32)
									if v57 <= int32(0) {
										v357 = v148
									} else {
										v154 = v133 + int32(8)
										v157 = int32(0)
										v159 = v134
										for {
											v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
											v164 = int32(1)
											v199 = v163&v164 + v157 + int32(base.Ui32(v163)>>(uint(int32(7))%32)) + int32(base.Ui32(v163)>>(uint(v164)%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(2))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(3))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(4))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(5))%32))&v164 + int32(base.Ui32(v163)>>(uint(int32(6))%32))&v164
											v203 = v159 + v164
											if v203 != v57 {
												v154 = v154 + v164
												v157 = v199
												v159 = v203
												continue
											} else {
												break
											}
											break
										}
										v357 = v148 - v199
									}
								}
							} else {
								if v142 != 0 {
									v207 = v57 << (uint(int32(3)) % 32)
									if v57 <= int32(0) {
										v357 = v207
									} else {
										v213 = v103 + int32(8)
										v216 = int32(0)
										v218 = v134
										for {
											v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
											v223 = int32(1)
											v258 = v222&v223 + v216 + int32(base.Ui32(v222)>>(uint(int32(7))%32)) + int32(base.Ui32(v222)>>(uint(v223)%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(2))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(3))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(4))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(5))%32))&v223 + int32(base.Ui32(v222)>>(uint(int32(6))%32))&v223
											v262 = v218 + v223
											if v262 != v57 {
												v213 = v213 + v223
												v216 = v258
												v218 = v262
												continue
											} else {
												break
											}
											break
										}
										v357 = v207 - v258
									}
								} else {
									if v57 <= int32(0) {
										v357 = int32(0)
									} else {
										v268 = int32(8)
										v269 = v133 + v268
										v271 = v103 + v268
										v272 = int32(1)
										v274 = v57 << (uint(int32(3)) % 32)
										if v274 <= v272 {
											v277 = v272
										} else {
											v277 = v274
										}
										v278 = int32(1)
										if v277 == v278 {
											v282 = int32(0)
											v323 = v282
											v324 = v282
										} else {
											v286 = int32(0)
											v289 = v286
											v290 = v286
											v291 = v286
											for {
												v299 = int32(base.Ui32(v290) >> (uint(int32(3)) % 32))
												v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v299))))
												v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299+v271))))
												v305 = base.I32_extend8_s(v301 ^ v303)
												v307 = v290 & int32(6)
												v308 = int32(1)
												v317 = int32(base.Ui32(v305)>>(uint(v307|v308)%32))&v308 + (int32(base.Ui32(v305)>>(uint(v307)%32))&v308 + v289)
												v318 = int32(2)
												v319 = v290 + v318
												v321 = v291 + v318
												if v321 != v277&int32(2147483640) {
													v289 = v317
													v290 = v319
													v291 = v321
													continue
												} else {
													break
												}
												break
											}
											v323 = v317
											v324 = v319
										}
										if v277&v278 != 0 {
											v333 = int32(base.Ui32(v324) >> (uint(int32(3)) % 32))
											v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v333))))
											v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333+v271))))
											v347 = int32(base.Ui32(base.I32_extend8_s(v335^v337))>>(uint(v324&int32(7))%32))&int32(1) + v323
										} else {
											v347 = v323
										}
										v357 = v347
									}
								}
							}
							v358 = base.B2i32(v111 < v357)
							if v111 < v357 {
								v359 = v357
							} else {
								v359 = v111
							}
							if v111 < v357 {
								v360 = v106
							} else {
								v360 = v115
							}
							if v111 < v357 {
								v361 = v86
							} else {
								v361 = v114
							}
							v363 = v106 + int32(1)
							v365 = v363 & int32(65535)
							if base.Ui32(v365) <= base.Ui32(v59) {
								v106 = v363
								v110 = v365
								v111 = v359
								v114 = v361
								v115 = v360
								continue
							} else {
								break
							}
							break
						}
						if v59 != v105 {
							v81 = v359
							v84 = v361
							v85 = v360
							v86 = v105
							continue
						} else {
							break
						}
						break
					}
					v376 = v361
					v377 = v360
				} else {
					v376 = v2
					v377 = v2
				}
				v392 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v392
				*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v392
				v396 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v397 = int32(8)
				v399 = v57 + v397
				v401 = v25 + int32(4)
				v403 = int32(65535)
				v411 = base.B2i32(v376&v403 == v392) | base.B2i32(v377&v403 == v392)
				if v411 != 0 {
					v412 = int32(1)
				} else {
					v412 = v376
				}
				v415 = int32(4)
				v418 = *(*int32)(unsafe.Add(mBase, uint32(v401+v412&int32(65535)<<(uint(v415)%32))))
				v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
				v421 = v419 & v415
				if v421 != 0 {
					v422 = v397
				} else {
					v422 = v399
				}
				v423 = F_palloc(m, v422)
				mBase = m.M
				v424 = m.ExcPending
				if v424 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v423)+4)) = v421
					*(*int32)(unsafe.Add(mBase, uint32(v423))) = v422 << (uint(int32(2)) % 32)
					if v421 == int32(0) {
						v431 = int32(8)
						if v57 != 0 {
							v435 = F__emscripten_memcpy_bulkmem(m, v423+v431, v418+v431, v57)
							mBase = m.M
						} else {
						}
					} else {
					}
					if v411 != 0 {
						v439 = int32(2)
					} else {
						v439 = v377
					}
					v442 = int32(4)
					v445 = *(*int32)(unsafe.Add(mBase, uint32(v401+v439&int32(65535)<<(uint(v442)%32))))
					v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
					v448 = v446 & v442
					if v448 != 0 {
						v449 = int32(8)
					} else {
						v449 = v399
					}
					v450 = F_palloc(m, v449)
					mBase = m.M
					v451 = m.ExcPending
					if v451 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v450)+4)) = v448
						*(*int32)(unsafe.Add(mBase, uint32(v450))) = v449 << (uint(int32(2)) % 32)
						if v448 == int32(0) {
							v458 = int32(8)
							if v57 != 0 {
								v462 = F__emscripten_memcpy_bulkmem(m, v450+v458, v445+v458, v57)
								mBase = m.M
							} else {
							}
						} else {
						}
						v464 = int32(65535)
						v465 = v26 + v464
						v467 = v465 & v464
						v470 = F_palloc(m, v467<<(uint(int32(3))%32))
						mBase = m.M
						v471 = m.ExcPending
						if v471 != 0 {
							return int32(0)
						} else {
							if v26&int32(65535) == int32(1) {
								F_pg_qsort(m, v470, v467, int32(8), int32(6879))
								mBase = m.M
								v479 = m.ExcPending
								if v479 != 0 {
									return int32(0)
								} else {
									v1907 = v396
									v1911 = v67
									v1919 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1907))) = uint16(v1919)
									*(*uint16)(unsafe.Add(mBase, uint32(v1911))) = uint16(v1919)
									*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v450
									*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v423
									return v29
								}
							} else {
								v480 = int32(1)
								v482 = v480
								v487 = v480
								for {
									v508 = v470 + v482<<(uint(int32(3))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v508-int32(8)))) = uint16(v487)
									v512 = int32(4)
									v517 = *(*int32)(unsafe.Add(mBase, uint32(v401+v482<<(uint(v512)%32))))
									v518 = int32(0)
									v524 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
									v526 = v524 & v512
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
									if v527&v512 != 0 {
										if v526 != 0 {
											v741 = int32(0)
										} else {
											v532 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v741 = v532
											} else {
												v538 = v517 + int32(8)
												v541 = int32(0)
												v543 = v518
												for {
													v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v538))))
													v548 = int32(1)
													v583 = v547&v548 + v541 + int32(base.Ui32(v547)>>(uint(int32(7))%32)) + int32(base.Ui32(v547)>>(uint(v548)%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(2))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(3))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(4))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(5))%32))&v548 + int32(base.Ui32(v547)>>(uint(int32(6))%32))&v548
													v587 = v543 + v548
													if v587 != v57 {
														v538 = v538 + v548
														v541 = v583
														v543 = v587
														continue
													} else {
														break
													}
													break
												}
												v741 = v532 - v583
											}
										}
									} else {
										if v526 != 0 {
											v591 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v741 = v591
											} else {
												v597 = v423 + int32(8)
												v600 = int32(0)
												v602 = v518
												for {
													v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v597))))
													v607 = int32(1)
													v642 = v606&v607 + v600 + int32(base.Ui32(v606)>>(uint(int32(7))%32)) + int32(base.Ui32(v606)>>(uint(v607)%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(2))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(3))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(4))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(5))%32))&v607 + int32(base.Ui32(v606)>>(uint(int32(6))%32))&v607
													v646 = v602 + v607
													if v646 != v57 {
														v597 = v597 + v607
														v600 = v642
														v602 = v646
														continue
													} else {
														break
													}
													break
												}
												v741 = v591 - v642
											}
										} else {
											if v57 <= int32(0) {
												v741 = int32(0)
											} else {
												v652 = int32(8)
												v653 = v517 + v652
												v655 = v423 + v652
												v656 = int32(1)
												v658 = v57 << (uint(int32(3)) % 32)
												if v658 <= v656 {
													v661 = v656
												} else {
													v661 = v658
												}
												v662 = int32(1)
												if v661 == v662 {
													v666 = int32(0)
													v707 = v666
													v708 = v666
												} else {
													v670 = int32(0)
													v673 = v670
													v674 = v670
													v675 = v670
													for {
														v683 = int32(base.Ui32(v674) >> (uint(int32(3)) % 32))
														v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v683))))
														v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683+v655))))
														v689 = base.I32_extend8_s(v685 ^ v687)
														v691 = v674 & int32(6)
														v692 = int32(1)
														v701 = int32(base.Ui32(v689)>>(uint(v691|v692)%32))&v692 + (int32(base.Ui32(v689)>>(uint(v691)%32))&v692 + v673)
														v702 = int32(2)
														v703 = v674 + v702
														v705 = v675 + v702
														if v705 != v661&int32(2147483640) {
															v673 = v701
															v674 = v703
															v675 = v705
															continue
														} else {
															break
														}
														break
													}
													v707 = v701
													v708 = v703
												}
												if v661&v662 != 0 {
													v717 = int32(base.Ui32(v708) >> (uint(int32(3)) % 32))
													v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653+v717))))
													v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v717+v655))))
													v731 = int32(base.Ui32(base.I32_extend8_s(v719^v721))>>(uint(v708&int32(7))%32))&int32(1) + v707
												} else {
													v731 = v707
												}
												v741 = v731
											}
										}
									}
									v742 = int32(0)
									v748 = *(*int32)(unsafe.Add(mBase, uint32(v517)+4))
									v749 = int32(4)
									v750 = v748 & v749
									v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
									if v751&v749 != 0 {
										if v750 != 0 {
											v965 = int32(0)
										} else {
											v756 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v965 = v756
											} else {
												v762 = v517 + int32(8)
												v765 = int32(0)
												v767 = v742
												for {
													v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
													v772 = int32(1)
													v807 = v771&v772 + v765 + int32(base.Ui32(v771)>>(uint(int32(7))%32)) + int32(base.Ui32(v771)>>(uint(v772)%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(2))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(3))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(4))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(5))%32))&v772 + int32(base.Ui32(v771)>>(uint(int32(6))%32))&v772
													v811 = v767 + v772
													if v811 != v57 {
														v762 = v762 + v772
														v765 = v807
														v767 = v811
														continue
													} else {
														break
													}
													break
												}
												v965 = v756 - v807
											}
										}
									} else {
										if v750 != 0 {
											v815 = v57 << (uint(int32(3)) % 32)
											if v57 <= int32(0) {
												v965 = v815
											} else {
												v821 = v450 + int32(8)
												v824 = int32(0)
												v826 = v742
												for {
													v830 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v821))))
													v831 = int32(1)
													v866 = v830&v831 + v824 + int32(base.Ui32(v830)>>(uint(int32(7))%32)) + int32(base.Ui32(v830)>>(uint(v831)%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(2))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(3))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(4))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(5))%32))&v831 + int32(base.Ui32(v830)>>(uint(int32(6))%32))&v831
													v870 = v826 + v831
													if v870 != v57 {
														v821 = v821 + v831
														v824 = v866
														v826 = v870
														continue
													} else {
														break
													}
													break
												}
												v965 = v815 - v866
											}
										} else {
											if v57 <= int32(0) {
												v965 = int32(0)
											} else {
												v876 = int32(8)
												v877 = v517 + v876
												v879 = v450 + v876
												v880 = int32(1)
												v882 = v57 << (uint(int32(3)) % 32)
												if v882 <= v880 {
													v885 = v880
												} else {
													v885 = v882
												}
												v886 = int32(1)
												if v885 == v886 {
													v890 = int32(0)
													v931 = v890
													v932 = v890
												} else {
													v894 = int32(0)
													v897 = v894
													v898 = v894
													v899 = v894
													for {
														v907 = int32(base.Ui32(v898) >> (uint(int32(3)) % 32))
														v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877+v907))))
														v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907+v879))))
														v913 = base.I32_extend8_s(v909 ^ v911)
														v915 = v898 & int32(6)
														v916 = int32(1)
														v925 = int32(base.Ui32(v913)>>(uint(v915|v916)%32))&v916 + (int32(base.Ui32(v913)>>(uint(v915)%32))&v916 + v897)
														v926 = int32(2)
														v927 = v898 + v926
														v929 = v899 + v926
														if v929 != v885&int32(2147483640) {
															v897 = v925
															v898 = v927
															v899 = v929
															continue
														} else {
															break
														}
														break
													}
													v931 = v925
													v932 = v927
												}
												if v885&v886 != 0 {
													v941 = int32(base.Ui32(v932) >> (uint(int32(3)) % 32))
													v943 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v877+v941))))
													v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v941+v879))))
													v955 = int32(base.Ui32(base.I32_extend8_s(v943^v945))>>(uint(v932&int32(7))%32))&int32(1) + v931
												} else {
													v955 = v931
												}
												v965 = v955
											}
										}
									}
									v966 = v741 - v965
									v968 = v966 >> (uint(int32(31)) % 32)
									*(*int32)(unsafe.Add(mBase, uint32(v508-v512))) = v966 ^ v968 - v968
									v973 = v487 + int32(1)
									v974 = int32(65535)
									v975 = v973 & v974
									if base.Ui32(v975) <= base.Ui32(v465&v974) {
										v482 = v975
										v487 = v973
										continue
									} else {
										break
									}
									break
								}
								F_pg_qsort(m, v470, v467, int32(8), int32(6879))
								mBase = m.M
								v982 = m.ExcPending
								if v982 != 0 {
									return int32(0)
								} else {
									v983 = int32(1)
									if base.Ui32(v467) <= base.Ui32(v983) {
										v986 = v983
									} else {
										v986 = v467
									}
									v988 = v57 & int32(2147483644)
									v990 = v57 & int32(3)
									v991 = int32(8)
									v992 = v450 + v991
									v994 = v423 + v991
									v1003 = int32(0)
									v1010 = v396
									v1014 = v67
									for {
										v1025 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v470+v1003<<(uint(int32(3))%32)))))
										if v412&int32(65535) == v1025 {
											*(*uint16)(unsafe.Add(mBase, uint32(v1010))) = uint16(v412)
											v1028 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1028 + int32(1)
											v1880 = v1010 + int32(2)
											v1884 = v1014
										} else {
											if v439&int32(65535) == v1025 {
												*(*uint16)(unsafe.Add(mBase, uint32(v1014))) = uint16(v439)
												v1038 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1038 + int32(1)
												v1880 = v1010
												v1884 = v1014 + int32(2)
											} else {
												v1042 = int32(4)
												v1045 = *(*int32)(unsafe.Add(mBase, uint32(v401+v1025<<(uint(v1042)%32))))
												v1046 = int32(0)
												v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
												v1054 = v1052 & v1042
												v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
												if v1055&v1042 != 0 {
													if v1054 != 0 {
														v1269 = int32(0)
													} else {
														v1060 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1269 = v1060
														} else {
															v1066 = v1045 + int32(8)
															v1069 = int32(0)
															v1071 = v1046
															for {
																v1075 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1066))))
																v1076 = int32(1)
																v1111 = v1075&v1076 + v1069 + int32(base.Ui32(v1075)>>(uint(int32(7))%32)) + int32(base.Ui32(v1075)>>(uint(v1076)%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(2))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(3))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(4))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(5))%32))&v1076 + int32(base.Ui32(v1075)>>(uint(int32(6))%32))&v1076
																v1115 = v1071 + v1076
																if v1115 != v57 {
																	v1066 = v1066 + v1076
																	v1069 = v1111
																	v1071 = v1115
																	continue
																} else {
																	break
																}
																break
															}
															v1269 = v1060 - v1111
														}
													}
												} else {
													if v1054 != 0 {
														v1119 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1269 = v1119
														} else {
															v1125 = v423 + int32(8)
															v1128 = int32(0)
															v1130 = v1046
															for {
																v1134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1125))))
																v1135 = int32(1)
																v1170 = v1134&v1135 + v1128 + int32(base.Ui32(v1134)>>(uint(int32(7))%32)) + int32(base.Ui32(v1134)>>(uint(v1135)%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(2))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(3))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(4))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(5))%32))&v1135 + int32(base.Ui32(v1134)>>(uint(int32(6))%32))&v1135
																v1174 = v1130 + v1135
																if v1174 != v57 {
																	v1125 = v1125 + v1135
																	v1128 = v1170
																	v1130 = v1174
																	continue
																} else {
																	break
																}
																break
															}
															v1269 = v1119 - v1170
														}
													} else {
														if v57 <= int32(0) {
															v1269 = int32(0)
														} else {
															v1180 = int32(8)
															v1181 = v1045 + v1180
															v1183 = v423 + v1180
															v1184 = int32(1)
															v1186 = v57 << (uint(int32(3)) % 32)
															if v1186 <= v1184 {
																v1189 = v1184
															} else {
																v1189 = v1186
															}
															v1190 = int32(1)
															if v1189 == v1190 {
																v1194 = int32(0)
																v1235 = v1194
																v1236 = v1194
															} else {
																v1198 = int32(0)
																v1201 = v1198
																v1202 = v1198
																v1203 = v1198
																for {
																	v1211 = int32(base.Ui32(v1202) >> (uint(int32(3)) % 32))
																	v1213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1211))))
																	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1211+v1183))))
																	v1217 = base.I32_extend8_s(v1213 ^ v1215)
																	v1219 = v1202 & int32(6)
																	v1220 = int32(1)
																	v1229 = int32(base.Ui32(v1217)>>(uint(v1219|v1220)%32))&v1220 + (int32(base.Ui32(v1217)>>(uint(v1219)%32))&v1220 + v1201)
																	v1230 = int32(2)
																	v1231 = v1202 + v1230
																	v1233 = v1203 + v1230
																	if v1233 != v1189&int32(2147483640) {
																		v1201 = v1229
																		v1202 = v1231
																		v1203 = v1233
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1235 = v1229
																v1236 = v1231
															}
															if v1189&v1190 != 0 {
																v1245 = int32(base.Ui32(v1236) >> (uint(int32(3)) % 32))
																v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1181+v1245))))
																v1249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245+v1183))))
																v1259 = int32(base.Ui32(base.I32_extend8_s(v1247^v1249))>>(uint(v1236&int32(7))%32))&int32(1) + v1235
															} else {
																v1259 = v1235
															}
															v1269 = v1259
														}
													}
												}
												v1271 = int32(0)
												v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1045)+4))
												v1278 = int32(4)
												v1279 = v1277 & v1278
												v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
												if v1280&v1278 != 0 {
													if v1279 != 0 {
														v1494 = int32(0)
													} else {
														v1285 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1494 = v1285
														} else {
															v1291 = v1045 + int32(8)
															v1294 = int32(0)
															v1296 = v1271
															for {
																v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1291))))
																v1301 = int32(1)
																v1336 = v1300&v1301 + v1294 + int32(base.Ui32(v1300)>>(uint(int32(7))%32)) + int32(base.Ui32(v1300)>>(uint(v1301)%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(2))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(3))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(4))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(5))%32))&v1301 + int32(base.Ui32(v1300)>>(uint(int32(6))%32))&v1301
																v1340 = v1296 + v1301
																if v1340 != v57 {
																	v1291 = v1291 + v1301
																	v1294 = v1336
																	v1296 = v1340
																	continue
																} else {
																	break
																}
																break
															}
															v1494 = v1285 - v1336
														}
													}
												} else {
													if v1279 != 0 {
														v1344 = v57 << (uint(int32(3)) % 32)
														if v57 <= int32(0) {
															v1494 = v1344
														} else {
															v1350 = v450 + int32(8)
															v1353 = int32(0)
															v1355 = v1271
															for {
																v1359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350))))
																v1360 = int32(1)
																v1395 = v1359&v1360 + v1353 + int32(base.Ui32(v1359)>>(uint(int32(7))%32)) + int32(base.Ui32(v1359)>>(uint(v1360)%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(2))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(3))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(4))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(5))%32))&v1360 + int32(base.Ui32(v1359)>>(uint(int32(6))%32))&v1360
																v1399 = v1355 + v1360
																if v1399 != v57 {
																	v1350 = v1350 + v1360
																	v1353 = v1395
																	v1355 = v1399
																	continue
																} else {
																	break
																}
																break
															}
															v1494 = v1344 - v1395
														}
													} else {
														if v57 <= int32(0) {
															v1494 = int32(0)
														} else {
															v1405 = int32(8)
															v1406 = v1045 + v1405
															v1408 = v450 + v1405
															v1409 = int32(1)
															v1411 = v57 << (uint(int32(3)) % 32)
															if v1411 <= v1409 {
																v1414 = v1409
															} else {
																v1414 = v1411
															}
															v1415 = int32(1)
															if v1414 == v1415 {
																v1419 = int32(0)
																v1460 = v1419
																v1461 = v1419
															} else {
																v1423 = int32(0)
																v1426 = v1423
																v1427 = v1423
																v1428 = v1423
																for {
																	v1436 = int32(base.Ui32(v1427) >> (uint(int32(3)) % 32))
																	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1436))))
																	v1440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1436+v1408))))
																	v1442 = base.I32_extend8_s(v1438 ^ v1440)
																	v1444 = v1427 & int32(6)
																	v1445 = int32(1)
																	v1454 = int32(base.Ui32(v1442)>>(uint(v1444|v1445)%32))&v1445 + (int32(base.Ui32(v1442)>>(uint(v1444)%32))&v1445 + v1426)
																	v1455 = int32(2)
																	v1456 = v1427 + v1455
																	v1458 = v1428 + v1455
																	if v1458 != v1414&int32(2147483640) {
																		v1426 = v1454
																		v1427 = v1456
																		v1428 = v1458
																		continue
																	} else {
																		break
																	}
																	break
																}
																v1460 = v1454
																v1461 = v1456
															}
															if v1414&v1415 != 0 {
																v1470 = int32(base.Ui32(v1461) >> (uint(int32(3)) % 32))
																v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406+v1470))))
																v1474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1470+v1408))))
																v1484 = int32(base.Ui32(base.I32_extend8_s(v1472^v1474))>>(uint(v1461&int32(7))%32))&int32(1) + v1460
															} else {
																v1484 = v1460
															}
															v1494 = v1484
														}
													}
												}
												v1496 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
												v1497 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
												v1498 = v1496 - v1497
												if base.F64_lt(base.F64_convert_i32_s(v1269), base.F64_add(base.F64_convert_i32_s(v1494), base.F64_mul(base.F64_convert_i32_s(v1498*v1498*v1498), float64(-0.0001)))) != 0 {
													v1506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423)+4)))
													if v1506&int32(4) != 0 {
													} else {
														v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
														if v1509&int32(4) != 0 {
															v1514 = F__emscripten_memset_bulkmem(m, v994, base.I32_extend8_s(int32(255)), v57)
															mBase = m.M
														} else {
															if v57 <= int32(0) {
															} else {
																v1518 = v1045 + int32(8)
																v1519 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v57) {
																	v1524 = v1519
																	v1538 = v1519
																	for {
																		v1548 = v1524 + v994
																		v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1548))))
																		v1551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1524+v1518))))
																		v1552 = v1549 | v1551
																		*(*uint8)(unsafe.Add(mBase, uint32(v1548))) = uint8(v1552)
																		v1555 = v1524 | int32(1)
																		v1556 = v994 + v1555
																		v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556))))
																		v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1555))))
																		v1560 = v1557 | v1559
																		*(*uint8)(unsafe.Add(mBase, uint32(v1556))) = uint8(v1560)
																		v1563 = v1524 | int32(2)
																		v1564 = v994 + v1563
																		v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1564))))
																		v1567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1563))))
																		v1568 = v1565 | v1567
																		*(*uint8)(unsafe.Add(mBase, uint32(v1564))) = uint8(v1568)
																		v1571 = v1524 | int32(3)
																		v1572 = v994 + v1571
																		v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1572))))
																		v1575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1518+v1571))))
																		v1576 = v1573 | v1575
																		*(*uint8)(unsafe.Add(mBase, uint32(v1572))) = uint8(v1576)
																		v1578 = int32(4)
																		v1579 = v1524 + v1578
																		v1581 = v1538 + v1578
																		if v1581 != v988 {
																			v1524 = v1579
																			v1538 = v1581
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v1583 = v1579
																} else {
																	v1583 = v1519
																}
																if v990 == int32(0) {
																} else {
																	v1609 = v1583
																	v1626 = v1519
																	for {
																		v1633 = v1609 + v994
																		v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633))))
																		v1636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1518))))
																		v1637 = v1634 | v1636
																		*(*uint8)(unsafe.Add(mBase, uint32(v1633))) = uint8(v1637)
																		v1639 = int32(1)
																		v1642 = v1626 + v1639
																		if v1642 != v990 {
																			v1609 = v1609 + v1639
																			v1626 = v1642
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
													*(*uint16)(unsafe.Add(mBase, uint32(v1010))) = uint16(v1025)
													v1669 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v1669 + int32(1)
													v1880 = v1010 + int32(2)
													v1884 = v1014
												} else {
													v1675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+4)))
													if v1675&int32(4) != 0 {
													} else {
														v1678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
														if v1678&int32(4) != 0 {
															v1683 = F__emscripten_memset_bulkmem(m, v992, base.I32_extend8_s(int32(255)), v57)
															mBase = m.M
														} else {
															if v57 <= int32(0) {
															} else {
																v1687 = v1045 + int32(8)
																v1688 = int32(0)
																if base.Ui32(int32(4)) <= base.Ui32(v57) {
																	v1693 = v1688
																	v1707 = v1688
																	for {
																		v1717 = v1693 + v992
																		v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1717))))
																		v1720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693+v1687))))
																		v1721 = v1718 | v1720
																		*(*uint8)(unsafe.Add(mBase, uint32(v1717))) = uint8(v1721)
																		v1724 = v1693 | int32(1)
																		v1725 = v992 + v1724
																		v1726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1725))))
																		v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1724))))
																		v1729 = v1726 | v1728
																		*(*uint8)(unsafe.Add(mBase, uint32(v1725))) = uint8(v1729)
																		v1732 = v1693 | int32(2)
																		v1733 = v992 + v1732
																		v1734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1733))))
																		v1736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1732))))
																		v1737 = v1734 | v1736
																		*(*uint8)(unsafe.Add(mBase, uint32(v1733))) = uint8(v1737)
																		v1740 = v1693 | int32(3)
																		v1741 = v992 + v1740
																		v1742 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1741))))
																		v1744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687+v1740))))
																		v1745 = v1742 | v1744
																		*(*uint8)(unsafe.Add(mBase, uint32(v1741))) = uint8(v1745)
																		v1747 = int32(4)
																		v1748 = v1693 + v1747
																		v1750 = v1707 + v1747
																		if v1750 != v988 {
																			v1693 = v1748
																			v1707 = v1750
																			continue
																		} else {
																			break
																		}
																		break
																	}
																	v1752 = v1748
																} else {
																	v1752 = v1688
																}
																if v990 == int32(0) {
																} else {
																	v1778 = v1752
																	v1795 = v1688
																	for {
																		v1802 = v1778 + v992
																		v1803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1802))))
																		v1805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1778+v1687))))
																		v1806 = v1803 | v1805
																		*(*uint8)(unsafe.Add(mBase, uint32(v1802))) = uint8(v1806)
																		v1808 = int32(1)
																		v1811 = v1795 + v1808
																		if v1811 != v990 {
																			v1778 = v1778 + v1808
																			v1795 = v1811
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
													*(*uint16)(unsafe.Add(mBase, uint32(v1014))) = uint16(v1025)
													v1838 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
													*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v1838 + int32(1)
													v1880 = v1010
													v1884 = v1014 + int32(2)
												}
											}
										}
										v1893 = v1003 + int32(1)
										if v1893 != v986 {
											v1003 = v1893
											v1010 = v1880
											v1014 = v1884
											continue
										} else {
											break
										}
										break
									}
									v1907 = v1880
									v1911 = v1884
									v1919 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v1907))) = uint16(v1919)
									*(*uint16)(unsafe.Add(mBase, uint32(v1911))) = uint16(v1919)
									*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v450
									*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v423
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
