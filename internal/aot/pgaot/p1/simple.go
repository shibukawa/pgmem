package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SimpleLruAutotuneBuffers(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	v3 = int32(16)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SimpleLruAutotuneBuffers[0]))
	v7 = base.I32_div_s(v5, int32(512))
	v9 = base.I32_rem_s(v7, v3)
	v10 = v7 - v9
	if v10 <= v3 {
		v13 = v3
	} else {
		v13 = v10
	}
	if int32(1024) < v13 {
		v16 = int32(1024)
	} else {
		v16 = v13
	}
	return v16
}
func F_build_simple_rel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v484 int32
	_ = v484
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v564 int32
	_ = v564
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v771 int32
	_ = v771
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v917 int64
	_ = v917
	var v919 int32
	_ = v919
	var v963 int32
	_ = v963
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 float64
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 float64
	_ = v1182
	var v1183 float64
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1233 int32
	_ = v1233
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1264 int32
	_ = v1264
	var v1270 int32
	_ = v1270
	var v1293 int32
	_ = v1293
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1314 int32
	_ = v1314
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1399 int32
	_ = v1399
	var v1427 int64
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1633 int64
	_ = v1633
	var v1635 int64
	_ = v1635
	var v1637 int64
	_ = v1637
	var v1639 int64
	_ = v1639
	var v1641 int64
	_ = v1641
	var v1643 int64
	_ = v1643
	var v1645 int64
	_ = v1645
	var v1647 int64
	_ = v1647
	var v1649 int64
	_ = v1649
	var v1651 int64
	_ = v1651
	var v1653 int64
	_ = v1653
	var v1655 int64
	_ = v1655
	var v1657 int64
	_ = v1657
	var v1659 int64
	_ = v1659
	var v1661 int64
	_ = v1661
	var v1663 int64
	_ = v1663
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1689 int32
	_ = v1689
	var v1711 int32
	_ = v1711
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1798 int32
	_ = v1798
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1966 int32
	_ = v1966
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1975 int32
	_ = v1975
	var v1977 int32
	_ = v1977
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2028 int32
	_ = v2028
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2056 int32
	_ = v2056
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2111 int32
	_ = v2111
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int64
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2144 int64
	_ = v2144
	var v2146 int64
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2188 int32
	_ = v2188
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2258 int32
	_ = v2258
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2357 int32
	_ = v2357
	var v2388 int32
	_ = v2388
	var v2395 int32
	_ = v2395
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2427 int32
	_ = v2427
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2448 int32
	_ = v2448
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2477 int32
	_ = v2477
	var v2506 int32
	_ = v2506
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2554 int32
	_ = v2554
	var v2558 int32
	_ = v2558
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2578 int64
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2582 int32
	_ = v2582
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2653 int32
	_ = v2653
	var v2656 int32
	_ = v2656
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2707 int32
	_ = v2707
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2753 int32
	_ = v2753
	var v2755 int32
	_ = v2755
	var v2761 int32
	_ = v2761
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2770 int32
	_ = v2770
	var v2775 int32
	_ = v2775
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2819 int32
	_ = v2819
	var v2822 int32
	_ = v2822
	var v2842 int32
	_ = v2842
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2850 int32
	_ = v2850
	var v2871 int32
	_ = v2871
	var v2877 int32
	_ = v2877
	var v2880 int32
	_ = v2880
	var v2905 int32
	_ = v2905
	var v2934 int32
	_ = v2934
	var v2962 int32
	_ = v2962
	var v2974 int32
	_ = v2974
	var v2980 int32
	_ = v2980
	var v2985 int32
	_ = v2985
	v4 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v33 = l1 << (uint(int32(2)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33+v34)))
	if v36 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39+v33)))
	v43 = F_palloc0(m, int32(304))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2974 = m.ExcPending
	if v2974 != 0 {
		goto L4
	} else {
		goto L531
	}
L4:
	;
	return int32(0)
L5:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v49 = int32(2)
	goto L8
L7:
	;
	v49 = int32(0)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(270)
	v53 = F_bms_make_singleton(m, l1)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v53
	v58 = *(*float64)(unsafe.Add(mBase, uint32(l0)+312))
	v59 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+25)) = uint16(v59)
	v62 = base.F64_gt(v58, float64(0))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)) = uint8(v62)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+32)) = v65
	v67 = F_create_empty_pathtarget(m)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+44)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v43)+52)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v43)+60)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v43)+76)) = l1
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v43)+116)) = v69
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+108)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v43)+84)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v43)+124)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v43)+132)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v43)+140)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v43)+148)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v43)+164)) = v80
	*(*int64)(unsafe.Add(mBase, uint32(v43)+156)) = int64(4294967295)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	if v98 != 0 {
		v109 = v80
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v110 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+176)) = v110
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+172)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+168)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v43)+184)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+192)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+200)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+208)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+216)) = v110
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+268)) = uint8(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+264)) = v112
	*(*int64)(unsafe.Add(mBase, uint32(v43)+256)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+236)) = v110
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+232)) = uint16(v112)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+224)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+272)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+280)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v43)+288)) = v110
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	switch v99 {
	case 0:
		goto L14
	default:
		goto L13
	case 2:
		goto L15
	}
L13:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+168))
	v109 = v108
	goto L11
L14:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+56))
	v105 = F_getRTEPermissionInfo(m, v104, v41)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)+84))
	if v100 != int32(1) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v109 = v107
	goto L11
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+112)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	switch v168 {
	case 0:
		goto L26
	case 1, 3, 4, 5, 6, 7:
		goto L29
	default:
		goto L27
	case 8:
		goto L28
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+244)) = l2
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+248))
	if v144 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+252)) = v156
	v158 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v43)+244)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v156
	*(*int64)(unsafe.Add(mBase, uint32(v43)+68)) = v158
	v166 = v156
	goto L18
L22:
	;
	v145 = v144
	goto L24
L23:
	;
	v145 = l2
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+248)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+252)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+104)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+72)) = v153
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	v166 = v155
	goto L18
L25:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[0]))
	if v2506 != 0 {
		goto L468
	} else {
		goto L469
	}
L26:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+20)))
	v213 = m.G0
	v215 = v213 - int32(48)
	m.G0 = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v219 = F_table_open(m, v211, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L38
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L35
	}
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43)+92)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+88)) = int32(-65536)
	goto L25
L29:
	;
	v169 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)) = uint16(v169)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+8))
	if v173 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v175 = v174
	goto L32
L31:
	;
	v175 = v169
	goto L32
L32:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+90)) = uint16(v175)
	v181 = F_palloc0_mul(m, int32(4), base.I32_extend16_s(v175)+int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+92)) = v181
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+90)))
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+88)))
	v190 = F_palloc0_mul(m, int32(4), v185-v186+int32(1))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = v190
	goto L25
L35:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v201
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_0), v30)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_1), int32(396), int32(_a_F_build_simple_rel_2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+188))
	if v222 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+118)))
	if v249 != int32(112) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+119)))
	switch v223 - int32(102) {
	case 0, 10:
		goto L39
	default:
		goto L41
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v233 + int32(4)
	F_errmsg(m, int32(_a_F_build_simple_rel_3), v215)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v241 = int32(*(*int8)(unsafe.Add(mBase, uint32(v240)+119)))
	F_errdetail_relkind_not_supported(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(152), int32(_a_F_build_simple_rel_5))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	goto L25
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2465 = m.ExcPending
	if v2465 != 0 {
		goto L4
	} else {
		goto L464
	}
L49:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_simple_rel[1])))
	if v254 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	v265 = int32(_a_F_build_simple_rel_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+88)) = uint16(v265)
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v267)+120)))
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+90)) = uint16(v268)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+80)) = v271
	v277 = F_palloc0(m, v268<<(uint(int32(2))%32)+int32(28))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	if v264 != 0 {
		goto L48
	} else {
		goto L56
	}
L53:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[2]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+308))
	v262 = base.B2i32(v260 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_build_simple_rel[1])) = uint8(v262)
	v264 = v262
	goto L55
L54:
	;
	v264 = int32(0)
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L51
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+92)) = v277
	v280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+90)))
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+88)))
	v287 = F_palloc0(m, (v280-v281)<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+96)) = v287
	if v212 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v219)+180))
	if v329 != 0 {
		goto L72
	} else {
		goto L73
	}
L60:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+119)))
	if v291 != int32(112) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v211
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+116))
	if v296 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L62
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+100)) = v312
	if v212 != 0 {
		goto L59
	} else {
		goto L70
	}
L65:
	;
	v312 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v300 = int32(0)
	v306 = F_hash_search(m, v296, v215+int32(32), v300, v215+int32(44))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+44)))
	if v308 != int32(1) {
		v312 = v300
		goto L64
	} else {
		goto L69
	}
L69:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	v312 = v311
	goto L64
L70:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v43)+96))
	v315 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+88)))
	F_estimate_rel_size(m, v219, v314-v315<<(uint(int32(2))%32), v43+int32(124), v43+int32(128), v43+int32(136))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	goto L59
L72:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+116))
	v332 = v330
	goto L74
L73:
	;
	v332 = int32(-1)
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+156)) = v332
	if v212 != 0 {
		goto L81
	} else {
		goto L82
	}
L75:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L4
	} else {
		goto L461
	}
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L4
	} else {
		goto L457
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L4
	} else {
		goto L454
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L4
	} else {
		goto L451
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L4
	} else {
		goto L448
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+116)) = v1293
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v1300 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v1300
	v1303 = F_RelationGetStatExtList(m, v219)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L4
	} else {
		goto L239
	}
L81:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+119)))
	if v335 != int32(112) {
		v1293 = v4
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_simple_rel[3])))
	if v339 == int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	v343 = int32(1)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v219)+56))
	if base.Ui32(v344) < base.Ui32(int32(_a_F_build_simple_rel_7)) {
		v353 = v343
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+116)))
	if v355 != int32(1) {
		v1293 = v4
		goto L80
	} else {
		goto L93
	}
L88:
	;
	if v353 != 0 {
		v1293 = v4
		goto L80
	} else {
		goto L92
	}
L89:
	;
	goto L88
L90:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+68))
	if v348 == int32(99) {
		v353 = v343
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v351 = F_isTempToastNamespace(m, v348)
	mBase = m.M
	v353 = v351
	goto L89
L92:
	;
	goto L87
L93:
	;
	v358 = F_RelationGetIndexList(m, v219)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L95
	}
L94:
	;
	F_list_free(m, v358)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L4
	} else {
		goto L237
	}
L95:
	;
	if v358 == int32(0) {
		v1264 = v4
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v362 <= int32(0) {
		v1264 = v4
		goto L94
	} else {
		goto L97
	}
L97:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v217<<(uint(int32(2))%32))))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+24))
	v392 = v4
	v393 = v4
	goto L98
L98:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v358)+12))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v398+v392<<(uint(int32(2))%32))))
	v403 = F_index_open(m, v402, v370)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	v1264 = v1233
	goto L94
L100:
	;
	v1239 = v392 + int32(1)
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	if v1239 < v1240 {
		v392 = v1239
		v393 = v1233
		goto L98
	} else {
		goto L236
	}
L101:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v403)+192))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+18)))
	if v406 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_relation_close(m, v403, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+19)))
	if v412 != int32(1) {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	v1233 = v393
	goto L100
L106:
	;
	v984 = F_RelationGetIndexExpressions(m, v403)
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L4
	} else {
		goto L173
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v442)+60)) = int64(0)
	v963 = v641
	goto L106
L108:
	;
	v917 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v442)+106)) = v917
	v919 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+68)) = v919
	*(*int64)(unsafe.Add(mBase, uint32(v442)+60)) = v917
	*(*int32)(unsafe.Add(mBase, uint32(v442)+116)) = v919
	v963 = v919
	goto L106
L109:
	;
	v820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+10)))
	if v820 != int32(1) {
		goto L107
	} else {
		goto L159
	}
L110:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v815 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v814)+92)) = uint8(v815)
	F_relation_close(m, v403, int32(0))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L4
	} else {
		goto L158
	}
L111:
	;
	v442 = F_palloc0(m, int32(120))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L120
	}
L112:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v403)+196))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	v417 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v416)+20)))
	v418 = int32(768)
	if v417&v418 == v418 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	if base.Ui32(v437) <= base.Ui32(v436) {
		goto L110
	} else {
		goto L119
	}
L114:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[4]))
	v436 = int32(2)
	v437 = v423
	goto L113
L115:
	;
	goto L116
L116:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	v426 = int32(3)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[4]))
	if base.B2i32(base.Ui32(v425) < base.Ui32(v426))|base.B2i32(base.Ui32(v429) < base.Ui32(v426)) != 0 {
		v436 = v425
		v437 = v429
		goto L113
	} else {
		goto L117
	}
L117:
	;
	if v425-v429 < int32(0) {
		goto L111
	} else {
		goto L118
	}
L118:
	;
	goto L110
L119:
	;
	goto L111
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442))) = int32(272)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+4)) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v403)+48))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+12)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v442)+8)) = v449
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v405)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+36)) = v452
	v454 = int32(*(*int16)(unsafe.Add(mBase, uint32(v405)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+40)) = v454
	v457 = F_palloc_mul(m, int32(4), v452)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+44)) = v457
	v461 = F_palloc_mul(m, int32(4), v454)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+48)) = v461
	v465 = F_palloc_mul(m, int32(4), v454)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+52)) = v465
	v469 = F_palloc_mul(m, int32(4), v454)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L4
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+56)) = v469
	v473 = F_palloc_mul(m, int32(1), v452)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+76)) = v473
	if int32(0) < v452 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v484 = int32(0)
	goto L129
L127:
	;
	goto L128
L128:
	;
	v552 = int32(0)
	v553 = base.B2i32(v454 <= v552)
	if v553 == v552 {
		goto L133
	} else {
		goto L134
	}
L129:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v442)+44))
	v512 = int32(1)
	v515 = int32(*(*int16)(unsafe.Add(mBase, uint32(v405+int32(48)+v484<<(uint(v512)%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v508+v484<<(uint(int32(2))%32)))) = v515
	v518 = v484 + v512
	v519 = F_index_can_return(m, v403, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L131
	}
L130:
	;
	goto L128
L131:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v442)+76))
	*(*uint8)(unsafe.Add(mBase, uint32(v521+v484))) = uint8(v519)
	if v518 != v452 {
		v484 = v518
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	v564 = int32(0)
	goto L136
L134:
	;
	goto L135
L135:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v403)+48))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v634)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+80)) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v403)+48))
	v638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v637)+119)))
	if v638 == int32(73) {
		goto L108
	} else {
		goto L139
	}
L136:
	;
	v585 = v564 << (uint(int32(2)) % 32)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v442)+52))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v403)+208))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v588+v585)))
	*(*int32)(unsafe.Add(mBase, uint32(v585+v586))) = v590
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v442)+56))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v403)+212))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v594+v585)))
	*(*int32)(unsafe.Add(mBase, uint32(v592+v585))) = v596
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v442)+48))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v403)+248))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v600+v585)))
	*(*int32)(unsafe.Add(mBase, uint32(v598+v585))) = v602
	v605 = v564 + int32(1)
	if v605 != v454 {
		v564 = v605
		goto L136
	} else {
		goto L138
	}
L137:
	;
	goto L135
L138:
	;
	goto L137
L139:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v403)+204))
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+106)) = uint8(v642)
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+107)) = uint8(v644)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+19)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+108)) = uint8(v646)
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+109)) = uint8(v648)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+112)) = uint8(v650)
	v652 = int32(0)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v641)+100))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+110)) = uint8(base.B2i32(v653 != v652))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v641)+104))
	if v657 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v219)+188))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v658)+168))
	v663 = base.B2i32(v659 != int32(0))
	goto L142
L141:
	;
	v663 = int32(0)
	goto L142
L142:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+111)) = uint8(v663)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v641)+112))
	if v665 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v641)+116))
	v670 = base.B2i32(v666 != int32(0))
	goto L145
L144:
	;
	v670 = int32(0)
	goto L145
L145:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+113)) = uint8(v670)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v641)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+116)) = v672
	v675 = F_RelationGetIndexAttOptions(m, v403, int32(1))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+72)) = v675
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v442)+80))
	if v678 != int32(403) {
		goto L109
	} else {
		goto L147
	}
L147:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v442)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v442)+60)) = v681
	v684 = F_palloc_mul(m, int32(1), v454)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+64)) = v684
	v688 = F_palloc_mul(m, int32(1), v454)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+68)) = v688
	if v454 <= v552 {
		v963 = v641
		goto L106
	} else {
		goto L150
	}
L150:
	;
	if v454 != int32(1) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v699 = v652
	v706 = int32(0)
	goto L154
L152:
	;
	v771 = v652
	goto L153
L153:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v442)+64))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v403)+224))
	v798 = int32(1)
	v801 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797+v771<<(uint(v798)%32)))))
	v803 = v801 & v798
	*(*uint8)(unsafe.Add(mBase, uint32(v795+v771))) = uint8(v803)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v442)+68))
	v810 = int32(base.Ui32(v801)>>(uint(v798)%32)) & v798
	*(*uint8)(unsafe.Add(mBase, uint32(v805+v771))) = uint8(v810)
	v963 = v641
	goto L106
L154:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v442)+64))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v403)+224))
	v726 = int32(1)
	v729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v725+v699<<(uint(v726)%32)))))
	v731 = v729 & v726
	*(*uint8)(unsafe.Add(mBase, uint32(v723+v699))) = uint8(v731)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v442)+68))
	v738 = int32(base.Ui32(v729)>>(uint(v726)%32)) & v726
	*(*uint8)(unsafe.Add(mBase, uint32(v733+v699))) = uint8(v738)
	v741 = v699 | v726
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v442)+64))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v403)+224))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744+v741<<(uint(v726)%32)))))
	v750 = v748 & v726
	*(*uint8)(unsafe.Add(mBase, uint32(v741+v742))) = uint8(v750)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v442)+68))
	v757 = int32(base.Ui32(v748)>>(uint(v726)%32)) & v726
	*(*uint8)(unsafe.Add(mBase, uint32(v752+v741))) = uint8(v757)
	v759 = int32(2)
	v760 = v699 + v759
	v762 = v706 + v759
	if v762 != v454&int32(_a_F_build_simple_rel_8) {
		v699 = v760
		v706 = v762
		goto L154
	} else {
		goto L156
	}
L155:
	;
	if v454&int32(1) == int32(0) {
		v963 = v641
		goto L106
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v771 = v760
	goto L153
L158:
	;
	v1233 = v393
	goto L100
L159:
	;
	v824 = F_palloc_mul(m, int32(4), v454)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+60)) = v824
	v828 = F_palloc_mul(m, int32(1), v454)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+64)) = v828
	v832 = F_palloc_mul(m, int32(1), v454)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+68)) = v832
	if v454 <= v552 {
		v963 = v641
		goto L106
	} else {
		goto L163
	}
L163:
	;
	v838 = v652
	goto L164
L164:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v442)+64))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v403)+224))
	v865 = int32(1)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+v838<<(uint(v865)%32)))))
	v870 = v868 & v865
	*(*uint8)(unsafe.Add(mBase, uint32(v862+v838))) = uint8(v870)
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v442)+68))
	v877 = int32(base.Ui32(v868)>>(uint(v865)%32)) & v865
	*(*uint8)(unsafe.Add(mBase, uint32(v872+v838))) = uint8(v877)
	v880 = v838 << (uint(int32(2)) % 32)
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v442)+52))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v880+v881)))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v442)+56))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v884+v880)))
	v888 = F_get_opfamily_member_for_cmptype(m, v883, v886, v886, v865)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L4
	} else {
		goto L166
	}
L165:
	;
	v963 = v641
	goto L106
L166:
	;
	if v888 == int32(0) {
		goto L107
	} else {
		goto L167
	}
L167:
	;
	v898 = F_get_ordering_op_properties(m, v888, v215+int32(32), v215+int32(44), v215+int32(40))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	if v898 == int32(0) {
		goto L107
	} else {
		goto L169
	}
L169:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v215)+44))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v442)+56))
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v903+v880)))
	if v902 != v905 {
		goto L107
	} else {
		goto L170
	}
L170:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v215)+40))
	if v907 != int32(1) {
		goto L107
	} else {
		goto L171
	}
L171:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v442)+60))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v910+v880))) = v912
	v915 = v838 + int32(1)
	if v915 != v454 {
		v838 = v915
		goto L164
	} else {
		goto L172
	}
L172:
	;
	goto L165
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+84)) = v984
	v987 = F_RelationGetIndexPredicate(m, v403)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+88)) = v987
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v442)+84))
	if v990 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	if v1005 != 0 {
		goto L184
	} else {
		goto L185
	}
L176:
	;
	v1005 = v987
	v1006 = int32(0)
	goto L175
L177:
	;
	goto L178
L178:
	;
	if v217 != int32(1) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	F_ChangeVarNodes(m, v990, int32(1), v217)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L4
	} else {
		goto L182
	}
L180:
	;
	v1000 = v990
	goto L181
L181:
	;
	v1001 = F_eval_const_expressions(m, l0, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L4
	} else {
		goto L183
	}
L182:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v442)+84))
	v1000 = v999
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+84)) = v1001
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v442)+88))
	v1005 = v1004
	v1006 = v1001
	goto L175
L184:
	;
	if v217 != int32(1) {
		goto L187
	} else {
		goto L188
	}
L185:
	;
	v1024 = v1006
	goto L186
L186:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v442)+12))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+76))
	v1027 = int32(0)
	if v1024 != 0 {
		goto L194
	} else {
		goto L195
	}
L187:
	;
	F_ChangeVarNodes(m, v1005, int32(1), v217)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L4
	} else {
		goto L190
	}
L188:
	;
	v1013 = v1005
	goto L189
L189:
	;
	v1014 = F_make_ands_explicit(m, v1013)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L4
	} else {
		goto L191
	}
L190:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v442)+88))
	v1013 = v1012
	goto L189
L191:
	;
	v1016 = F_eval_const_expressions(m, l0, v1014)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+88)) = v1016
	v1019 = F_make_ands_implicit(m, v1016)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+88)) = v1019
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v442)+84))
	v1024 = v1022
	goto L186
L194:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+12))
	v1030 = v1029
	goto L196
L195:
	;
	v1030 = v1027
	goto L196
L196:
	;
	v1031 = int32(0)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v442)+36))
	if v1031 < v1032 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1038 = v1031
	v1044 = v1030
	v1045 = v1027
	goto L200
L198:
	;
	v1129 = v1030
	v1130 = v1027
	goto L199
L199:
	;
	if v1129 != 0 {
		goto L78
	} else {
		goto L219
	}
L200:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v442)+44))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1062+v1038<<(uint(int32(2))%32))))
	if v1066 != 0 {
		goto L203
	} else {
		goto L204
	}
L201:
	;
	v1129 = v1107
	v1130 = v1116
	goto L199
L202:
	;
	v1110 = v1038 + int32(1)
	v1112 = int32(0)
	v1114 = F_makeTargetEntry(m, v1106, base.I32_extend16_s(v1110), v1112, v1112)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L4
	} else {
		goto L216
	}
L203:
	;
	if v1066 < int32(0) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	goto L205
L205:
	;
	if v1044 == int32(0) {
		goto L79
	} else {
		goto L212
	}
L206:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+68))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+76))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+96))
	v1090 = F_makeVar(m, v1026, base.I32_extend16_s(v1082), v1086, v1087, v1088, int32(0))
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L4
	} else {
		goto L211
	}
L207:
	;
	v1069 = base.I32_extend16_s(v1066)
	v1070 = F_SystemAttributeDefinition(m, v1069)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L4
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v219)+52))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1072)))
	v1082 = v1066
	v1084 = v1072 + v1073<<(uint(int32(3))%32) + v1066*int32(100) - int32(72)
	goto L206
L210:
	;
	v1082 = v1069
	v1084 = v1070
	goto L206
L211:
	;
	v1106 = v1090
	v1107 = v1044
	goto L202
L212:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1044)))
	v1096 = v1044 + int32(4)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v442)+84))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+12))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+4))
	if base.Ui32(v1096) < base.Ui32(v1099+v1100<<(uint(int32(2))%32)) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1105 = v1096
	goto L215
L214:
	;
	v1105 = int32(0)
	goto L215
L215:
	;
	v1106 = v1094
	v1107 = v1105
	goto L202
L216:
	;
	v1116 = F_lappend(m, v1045, v1114)
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L4
	} else {
		goto L217
	}
L217:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v442)+36))
	if v1110 < v1118 {
		v1038 = v1110
		v1044 = v1107
		v1045 = v1116
		goto L200
	} else {
		goto L218
	}
L218:
	;
	goto L201
L219:
	;
	v1147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+100)) = uint8(v1147)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+96)) = v1147
	*(*int32)(unsafe.Add(mBase, uint32(v442)+92)) = v1130
	v1152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+101)) = uint8(v1152)
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+102)) = uint8(v1154)
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+105)) = uint8(v1147)
	*(*uint8)(unsafe.Add(mBase, uint32(v442)+103)) = uint8(v1156)
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v403)+48))
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160)+119)))
	if v1161 != int32(73) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+32)) = v1204
	F_relation_close(m, v403, int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L4
	} else {
		goto L234
	}
L221:
	;
	v1204 = int32(-1)
	goto L220
L222:
	;
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v442)+88))
	if v1164 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	goto L224
L224:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v442)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v442)+16)) = int32(0)
	goto L221
L225:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v963)+68))
	if v1190 == int32(0) {
		goto L221
	} else {
		goto L232
	}
L226:
	;
	v1168 = F_RelationGetNumberOfBlocksInFork(m, v403, int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L4
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1177 = v442 + int32(24)
	F_estimate_rel_size(m, v403, int32(0), v442+int32(16), v1177, v215+int32(32))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L230
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v442)+16)) = v1168
	v1171 = *(*float64)(unsafe.Add(mBase, uint32(v43)+128))
	*(*float64)(unsafe.Add(mBase, uint32(v442)+24)) = v1171
	goto L225
L230:
	;
	v1182 = *(*float64)(unsafe.Add(mBase, uint32(v43)+128))
	v1183 = *(*float64)(unsafe.Add(mBase, uint32(v442)+24))
	if base.F64_lt(v1182, v1183) == int32(0) {
		goto L225
	} else {
		goto L231
	}
L231:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1177))) = v1182
	goto L225
L232:
	;
	v1193 = m.T0[v1190].(func(*base.Module, int32) int32)(m, v403)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	v1204 = v1193
	goto L220
L234:
	;
	v1209 = F_lcons(m, v442, v393)
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L4
	} else {
		goto L235
	}
L235:
	;
	v1233 = v1209
	goto L100
L236:
	;
	goto L99
L237:
	;
	v1293 = v1264
	goto L80
L238:
	;
	F_list_free(m, v1303)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L4
	} else {
		goto L272
	}
L239:
	;
	if v1303 == int32(0) {
		v1474 = v1300
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+4))
	if v1307 <= int32(0) {
		v1474 = v1300
		goto L238
	} else {
		goto L241
	}
L241:
	;
	v1314 = v1300
	goto L242
L242:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+12))
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1338+v1314<<(uint(int32(2))%32))))
	v1344 = F_SearchSysCache1(m, int32(64), base.I64_extend_i32_u(v1342))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L4
	} else {
		goto L244
	}
L243:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v215)+32))
	v1474 = v1469
	goto L238
L244:
	;
	if v1344 == int32(0) {
		goto L77
	} else {
		goto L245
	}
L245:
	;
	v1348 = int32(0)
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1344)+16))
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1350)+22)))
	v1352 = v1350 + v1351
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+96))
	if v1348 < v1353 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1362 = v1348
	v1365 = int32(0)
	goto L249
L247:
	;
	v1399 = v1348
	goto L248
L248:
	;
	v1427 = F_SysCacheGetAttr(m, int32(64), v1344, int32(9), v215+int32(44))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L4
	} else {
		goto L253
	}
L249:
	;
	v1389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1352+int32(104)+v1365<<(uint(int32(1))%32)))))
	v1390 = F_bms_add_member(m, v1362, v1389)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L4
	} else {
		goto L251
	}
L250:
	;
	v1399 = v1390
	goto L248
L251:
	;
	v1393 = v1365 + int32(1)
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1352)+96))
	if v1393 < v1394 {
		v1362 = v1390
		v1365 = v1393
		goto L249
	} else {
		goto L252
	}
L252:
	;
	goto L250
L253:
	;
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+44)))
	if v1429 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1433 = F_text_to_cstring(m, base.I32_wrap_i64(v1427))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L4
	} else {
		goto L257
	}
L255:
	;
	v1452 = v1348
	goto L256
L256:
	;
	v1454 = v215 + int32(32)
	F_get_relation_statistics_worker(m, v1454, v43, v1342, int32(1), v1399, v1452)
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L4
	} else {
		goto L267
	}
L257:
	;
	v1435 = F_stringToNode(m, v1433)
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	F_pfree(m, v1433)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L4
	} else {
		goto L259
	}
L259:
	;
	v1440 = F_expand_generated_columns_in_expr(m, v1435, v219, int32(1))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L260
	}
L260:
	;
	if v1299 != int32(1) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	F_ChangeVarNodes(m, v1440, int32(1), v1299)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L4
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1447 = F_eval_const_expressions(m, l0, v1440)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L4
	} else {
		goto L265
	}
L264:
	;
	goto L263
L265:
	;
	F_fix_opfuncids(m, v1447)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	v1452 = v1447
	goto L256
L267:
	;
	F_get_relation_statistics_worker(m, v1454, v43, v1342, int32(0), v1399, v1452)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L4
	} else {
		goto L268
	}
L268:
	;
	F_ReleaseCatCache(m, v1344)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L4
	} else {
		goto L269
	}
L269:
	;
	F_bms_free(m, v1399)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L4
	} else {
		goto L270
	}
L270:
	;
	v1466 = v1314 + int32(1)
	v1467 = *(*int32)(unsafe.Add(mBase, uint32(v1303)+4))
	if v1466 < v1467 {
		v1314 = v1466
		goto L242
	} else {
		goto L271
	}
L271:
	;
	goto L243
L272:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+120)) = v1474
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v1501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1500)+119)))
	if v1501 == int32(102) {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+176)) = v1518
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v1520 != 0 {
		goto L280
	} else {
		goto L281
	}
L274:
	;
	v1505 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_build_simple_rel[5])))
	if v1505&int32(2) != 0 {
		goto L76
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v1515 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+164)) = v1515
	v1518 = v1515
	goto L273
L277:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v219)+56))
	v1509 = F_GetForeignServerIdByRelId(m, v1508)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+164)) = v1509
	v1513 = F_GetFdwRoutineForRelation(m, v219, int32(1))
	mBase = m.M
	v1514 = m.ExcPending
	if v1514 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	v1518 = v1513
	goto L273
L280:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v219)+188))
	if v1740 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L281:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+52))
	if base.B2i32(v1522 == int32(0))|v212 != 0 {
		goto L280
	} else {
		goto L282
	}
L282:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	if v1526 < int32(2) {
		goto L280
	} else {
		goto L283
	}
L283:
	;
	v1529 = F_RelationGetFKeyList(m, v219)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	if v1529 == int32(0) {
		goto L280
	} else {
		goto L285
	}
L285:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+4))
	if v1533 <= int32(0) {
		goto L280
	} else {
		goto L286
	}
L286:
	;
	v1541 = int32(0)
	v1543 = v1533
	goto L287
L287:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+12))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1564+v1541<<(uint(int32(2))%32))))
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1568)+20)))
	if v1569 != int32(1) {
		v1689 = v1543
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L280
L289:
	;
	v1711 = v1541 + int32(1)
	if v1711 < v1689 {
		v1541 = v1711
		v1543 = v1689
		goto L287
	} else {
		goto L302
	}
L290:
	;
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	if v1572 <= int32(0) {
		v1689 = v1543
		goto L289
	} else {
		goto L291
	}
L291:
	;
	v1578 = v1568 + int32(86)
	v1580 = v1568 + int32(22)
	v1588 = int32(0)
	v1589 = v1572
	goto L292
L292:
	;
	v1610 = v1588 + int32(1)
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+12))
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1611+v1588<<(uint(int32(2))%32))))
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1615)+12))
	if v1616 != 0 {
		v1680 = v1589
		goto L294
	} else {
		goto L295
	}
L293:
	;
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+4))
	v1689 = v1682
	goto L289
L294:
	;
	if v1610 < v1680 {
		v1588 = v1610
		v1589 = v1680
		goto L292
	} else {
		goto L301
	}
L295:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1615)+16))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+12))
	if v1617 != v1618 {
		v1680 = v1589
		goto L294
	} else {
		goto L296
	}
L296:
	;
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1615)+20)))
	if v1620 != 0 {
		v1680 = v1589
		goto L294
	} else {
		goto L297
	}
L297:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	if v1610 == v1621 {
		v1680 = v1589
		goto L294
	} else {
		goto L298
	}
L298:
	;
	v1624 = F_palloc0(m, int32(672))
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L4
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1624))) = int32(273)
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1624)+8)) = v1610
	*(*int32)(unsafe.Add(mBase, uint32(v1624)+4)) = v1628
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1624)+12)) = v1631
	v1633 = *(*int64)(unsafe.Add(mBase, uint32(v1580)))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+16)) = v1633
	v1635 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+24)) = v1635
	v1637 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+32)) = v1637
	v1639 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+40)) = v1639
	v1641 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+48)) = v1641
	v1643 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+56)) = v1643
	v1645 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+64)) = v1645
	v1647 = *(*int64)(unsafe.Add(mBase, uint32(v1580)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+72)) = v1647
	v1649 = *(*int64)(unsafe.Add(mBase, uint32(v1578)))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+80)) = v1649
	v1651 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+88)) = v1651
	v1653 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+96)) = v1653
	v1655 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+104)) = v1655
	v1657 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+112)) = v1657
	v1659 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+120)) = v1659
	v1661 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+128)) = v1661
	v1663 = *(*int64)(unsafe.Add(mBase, uint32(v1578)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1624)+136)) = v1663
	base.MemoryCopy(m, v1624+int32(144), v1568+int32(152), int32(128))
	base.MemoryFill(m, v1624+int32(272), int32(0), int32(400))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v1675 = F_lappend(m, v1674, v1624)
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v1675
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+4))
	v1680 = v1678
	goto L294
L301:
	;
	goto L293
L302:
	;
	goto L288
L303:
	;
	if v212 == int32(0) {
		goto L307
	} else {
		goto L308
	}
L304:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1740)+24))
	if v1743 == int32(0) {
		goto L303
	} else {
		goto L305
	}
L305:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1740)+28))
	if v1746 == int32(0) {
		goto L303
	} else {
		goto L306
	}
L306:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v43)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+160)) = v1749 | int32(1)
	goto L303
L307:
	;
	F_relation_close(m, v219, int32(0))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L4
	} else {
		goto L447
	}
L308:
	;
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v219)+48))
	v1756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1755)+119)))
	if v1756 != int32(112) {
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+112))
	if v1760 != 0 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1771 = v1760
	goto L312
L311:
	;
	v1762 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[6]))
	v1764 = F_CreatePartitionDirectory(m, v1762, int32(1))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L4
	} else {
		goto L313
	}
L312:
	;
	v1772 = F_PartitionDirectoryLookup(m, v1771, v219)
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L4
	} else {
		goto L314
	}
L313:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1766)+112)) = v1764
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1768)+112))
	v1771 = v1769
	goto L312
L314:
	;
	v1774 = F_RelationGetPartitionKey(m, v219)
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L315
	}
L315:
	;
	v1776 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1774)+4)))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if v1777 == int32(0) {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+256)) = v2188
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v1772)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+264)) = v2213
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v1772)))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+260)) = v2215
	v2217 = F_RelationGetPartitionKey(m, v219)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L4
	} else {
		goto L415
	}
L317:
	;
	v2052 = F_palloc0(m, int32(28))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L4
	} else {
		goto L385
	}
L318:
	;
	v2028 = v1776 << (uint(int32(2)) % 32)
	goto L317
L319:
	;
	goto L320
L320:
	;
	v1783 = v1776 << (uint(int32(2)) % 32)
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+4))
	if v1784 <= int32(0) {
		v2028 = v1783
		goto L317
	} else {
		goto L321
	}
L321:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1774)))
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1777)+12))
	v1798 = int32(0)
	goto L322
L322:
	;
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1788+v1798<<(uint(int32(2))%32))))
	v1823 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1822))))
	if v1787 != v1823 {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	v2028 = v1783
	goto L317
L324:
	;
	v2022 = v1798 + int32(1)
	if v1784 != v2022 {
		v1798 = v2022
		goto L322
	} else {
		goto L384
	}
L325:
	;
	v1825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1822)+2)))
	if v1776&int32(_a_F_build_simple_rel_9) != v1825 {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+16))
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1783) {
		goto L330
	} else {
		goto L331
	}
L327:
	;
	if v1890 != 0 {
		goto L324
	} else {
		goto L345
	}
L328:
	;
	v1890 = int32(0)
	goto L327
L329:
	;
	v1864 = v1859
	v1865 = v1860
	v1866 = v1861
	goto L339
L330:
	;
	if (v1827|v1828)&int32(3) != 0 {
		v1859 = v1827
		v1860 = v1828
		v1861 = v1783
		goto L329
	} else {
		goto L333
	}
L331:
	;
	v1852 = v1827
	v1853 = v1828
	v1854 = v1783
	goto L332
L332:
	;
	if v1854 == int32(0) {
		goto L328
	} else {
		goto L338
	}
L333:
	;
	v1836 = v1827
	v1837 = v1828
	v1838 = v1783
	goto L334
L334:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1836)))
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1837)))
	if v1841 != v1842 {
		v1859 = v1836
		v1860 = v1837
		v1861 = v1838
		goto L329
	} else {
		goto L336
	}
L335:
	;
	v1852 = v1847
	v1853 = v1845
	v1854 = v1849
	goto L332
L336:
	;
	v1844 = int32(4)
	v1845 = v1837 + v1844
	v1847 = v1836 + v1844
	v1849 = v1838 - v1844
	if base.Ui32(int32(3)) < base.Ui32(v1849) {
		v1836 = v1847
		v1837 = v1845
		v1838 = v1849
		goto L334
	} else {
		goto L337
	}
L337:
	;
	goto L335
L338:
	;
	v1859 = v1852
	v1860 = v1853
	v1861 = v1854
	goto L329
L339:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1864))))
	v1870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865))))
	if v1869 == v1870 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	v1890 = v1869 - v1870
	goto L327
L341:
	;
	v1872 = int32(1)
	v1877 = v1866 - v1872
	if v1877 != 0 {
		v1864 = v1864 + v1872
		v1865 = v1865 + v1872
		v1866 = v1877
		goto L339
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	goto L340
L344:
	;
	goto L328
L345:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+20))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v1783) {
		goto L349
	} else {
		goto L350
	}
L346:
	;
	if v1954 != 0 {
		goto L324
	} else {
		goto L364
	}
L347:
	;
	v1954 = int32(0)
	goto L346
L348:
	;
	v1928 = v1923
	v1929 = v1924
	v1930 = v1925
	goto L358
L349:
	;
	if (v1891|v1892)&int32(3) != 0 {
		v1923 = v1891
		v1924 = v1892
		v1925 = v1783
		goto L348
	} else {
		goto L352
	}
L350:
	;
	v1916 = v1891
	v1917 = v1892
	v1918 = v1783
	goto L351
L351:
	;
	if v1918 == int32(0) {
		goto L347
	} else {
		goto L357
	}
L352:
	;
	v1900 = v1891
	v1901 = v1892
	v1902 = v1783
	goto L353
L353:
	;
	v1905 = *(*int32)(unsafe.Add(mBase, uint32(v1900)))
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1901)))
	if v1905 != v1906 {
		v1923 = v1900
		v1924 = v1901
		v1925 = v1902
		goto L348
	} else {
		goto L355
	}
L354:
	;
	v1916 = v1911
	v1917 = v1909
	v1918 = v1913
	goto L351
L355:
	;
	v1908 = int32(4)
	v1909 = v1901 + v1908
	v1911 = v1900 + v1908
	v1913 = v1902 - v1908
	if base.Ui32(int32(3)) < base.Ui32(v1913) {
		v1900 = v1911
		v1901 = v1909
		v1902 = v1913
		goto L353
	} else {
		goto L356
	}
L356:
	;
	goto L354
L357:
	;
	v1923 = v1916
	v1924 = v1917
	v1925 = v1918
	goto L348
L358:
	;
	v1933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1928))))
	v1934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1929))))
	if v1933 == v1934 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v1954 = v1933 - v1934
	goto L346
L360:
	;
	v1936 = int32(1)
	v1941 = v1930 - v1936
	if v1941 != 0 {
		v1928 = v1928 + v1936
		v1929 = v1929 + v1936
		v1930 = v1941
		goto L358
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	goto L359
L363:
	;
	goto L347
L364:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+28))
	v1956 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v1783) {
		goto L368
	} else {
		goto L369
	}
L365:
	;
	if v2018 == int32(0) {
		v2188 = v1822
		goto L316
	} else {
		goto L383
	}
L366:
	;
	v2018 = int32(0)
	goto L365
L367:
	;
	v1992 = v1987
	v1993 = v1988
	v1994 = v1989
	goto L377
L368:
	;
	if (v1955|v1956)&int32(3) != 0 {
		v1987 = v1955
		v1988 = v1956
		v1989 = v1783
		goto L367
	} else {
		goto L371
	}
L369:
	;
	v1980 = v1955
	v1981 = v1956
	v1982 = v1783
	goto L370
L370:
	;
	if v1982 == int32(0) {
		goto L366
	} else {
		goto L376
	}
L371:
	;
	v1964 = v1955
	v1965 = v1956
	v1966 = v1783
	goto L372
L372:
	;
	v1969 = *(*int32)(unsafe.Add(mBase, uint32(v1964)))
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1965)))
	if v1969 != v1970 {
		v1987 = v1964
		v1988 = v1965
		v1989 = v1966
		goto L367
	} else {
		goto L374
	}
L373:
	;
	v1980 = v1975
	v1981 = v1973
	v1982 = v1977
	goto L370
L374:
	;
	v1972 = int32(4)
	v1973 = v1965 + v1972
	v1975 = v1964 + v1972
	v1977 = v1966 - v1972
	if base.Ui32(int32(3)) < base.Ui32(v1977) {
		v1964 = v1975
		v1965 = v1973
		v1966 = v1977
		goto L372
	} else {
		goto L375
	}
L375:
	;
	goto L373
L376:
	;
	v1987 = v1980
	v1988 = v1981
	v1989 = v1982
	goto L367
L377:
	;
	v1997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1992))))
	v1998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993))))
	if v1997 == v1998 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v2018 = v1997 - v1998
	goto L365
L379:
	;
	v2000 = int32(1)
	v2005 = v1994 - v2000
	if v2005 != 0 {
		v1992 = v1992 + v2000
		v1993 = v1993 + v2000
		v1994 = v2005
		goto L377
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	goto L378
L382:
	;
	goto L366
L383:
	;
	goto L324
L384:
	;
	goto L323
L385:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v1774)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2052))) = uint8(v2054)
	v2056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1774)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2052)+2)) = uint16(v2056)
	v2059 = F_palloc_mul(m, int32(4), v1776)
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+4)) = v2059
	v2062 = int32(0)
	v2063 = base.B2i32(v2028 == v2062)
	if v2063 == v2062 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+16))
	base.MemoryCopy(m, v2059, v2066, v2028)
	goto L389
L388:
	;
	goto L389
L389:
	;
	v2069 = F_palloc_mul(m, int32(4), v1776)
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L4
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+8)) = v2069
	if v2063 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	v2074 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+20))
	base.MemoryCopy(m, v2069, v2074, v2028)
	goto L393
L392:
	;
	goto L393
L393:
	;
	v2077 = F_palloc_mul(m, int32(4), v1776)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L4
	} else {
		goto L394
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+12)) = v2077
	if v2063 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+28))
	base.MemoryCopy(m, v2077, v2082, v2028)
	goto L397
L396:
	;
	goto L397
L397:
	;
	v2085 = F_palloc_mul(m, int32(2), v1776)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L4
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+16)) = v2085
	v2089 = v1776 << (uint(int32(1)) % 32)
	if v2089 != 0 {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+40))
	base.MemoryCopy(m, v2085, v2090, v2089)
	goto L401
L400:
	;
	goto L401
L401:
	;
	v2093 = F_palloc_mul(m, int32(1), v1776)
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L4
	} else {
		goto L402
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+20)) = v2093
	if v1776 != 0 {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+44))
	base.MemoryCopy(m, v2093, v2096, v1776)
	goto L405
L404:
	;
	goto L405
L405:
	;
	v2099 = F_palloc_mul(m, int32(28), v1776)
	mBase = m.M
	v2100 = m.ExcPending
	if v2100 != 0 {
		goto L4
	} else {
		goto L406
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2052)+24)) = v2099
	if int32(0) < v1776 {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v2111 = int32(0)
	goto L410
L408:
	;
	goto L409
L409:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v2182 = F_lappend(m, v2181, v2052)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L4
	} else {
		goto L414
	}
L410:
	;
	v2133 = v2111 * int32(28)
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v2052)+24))
	v2135 = v2133 + v2134
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v1774)+24))
	v2137 = v2136 + v2133
	v2139 = *(*int32)(unsafe.Add(mBase, _c_F_build_simple_rel[6]))
	v2140 = *(*int64)(unsafe.Add(mBase, uint32(v2137)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2135)+16)) = v2140
	v2142 = *(*int32)(unsafe.Add(mBase, uint32(v2137)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2135)+24)) = v2142
	v2144 = *(*int64)(unsafe.Add(mBase, uint32(v2137)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2135)+8)) = v2144
	v2146 = *(*int64)(unsafe.Add(mBase, uint32(v2137)))
	*(*int64)(unsafe.Add(mBase, uint32(v2135))) = v2146
	*(*int32)(unsafe.Add(mBase, uint32(v2135)+20)) = v2139
	*(*int32)(unsafe.Add(mBase, uint32(v2135)+16)) = int32(0)
	goto L412
L411:
	;
	goto L409
L412:
	;
	v2152 = v2111 + int32(1)
	if v2152 != v1776 {
		v2111 = v2152
		goto L410
	} else {
		goto L413
	}
L413:
	;
	goto L411
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v2182
	v2188 = v2052
	goto L316
L415:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	v2221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2217)+4)))
	v2222 = F_palloc_mul(m, int32(4), v2221)
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+12))
	if v2224 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2224)+12))
	v2227 = v2225
	goto L419
L418:
	;
	v2227 = int32(0)
	goto L419
L419:
	;
	if int32(0) < v2221 {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2237 = int32(0)
	v2241 = v2227
	goto L423
L421:
	;
	goto L422
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+288)) = v2222
	v2342 = F_palloc0_mul(m, int32(4), v2221)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L4
	} else {
		goto L438
	}
L423:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+8))
	v2262 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2258+v2237<<(uint(int32(1))%32)))))
	if v2262 != 0 {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	goto L422
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v2296
	*(*int32)(unsafe.Add(mBase, uint32(v215)+32)) = v2296
	v2307 = F_list_make1_impl(m, int32(1), v215+int32(12))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L4
	} else {
		goto L436
	}
L426:
	;
	v2264 = v2237 << (uint(int32(2)) % 32)
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+32))
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v2264+v2265)))
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+36))
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2268+v2264)))
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+52))
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2271+v2264)))
	v2275 = F_makeVar(m, v2219, v2262, v2267, v2270, v2273, int32(0))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L4
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	if v2241 == int32(0) {
		goto L75
	} else {
		goto L430
	}
L429:
	;
	v2296 = v2275
	v2297 = v2241
	goto L425
L430:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2241)))
	v2280 = F_copyObjectImpl(m, v2279)
	mBase = m.M
	v2281 = m.ExcPending
	if v2281 != 0 {
		goto L4
	} else {
		goto L431
	}
L431:
	;
	F_ChangeVarNodes(m, v2280, int32(1), v2219)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L4
	} else {
		goto L432
	}
L432:
	;
	v2286 = v2241 + int32(4)
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2217)+12))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v2288)+12))
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2288)+4))
	if base.Ui32(v2286) < base.Ui32(v2289+v2290<<(uint(int32(2))%32)) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2295 = v2286
	goto L435
L434:
	;
	v2295 = int32(0)
	goto L435
L435:
	;
	v2296 = v2280
	v2297 = v2295
	goto L425
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2222+v2237<<(uint(int32(2))%32)))) = v2307
	v2311 = v2237 + int32(1)
	if v2311 != v2221 {
		v2237 = v2311
		v2241 = v2297
		goto L423
	} else {
		goto L437
	}
L437:
	;
	goto L424
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+292)) = v2342
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v43)+272))
	if v2345 != 0 {
		goto L307
	} else {
		goto L439
	}
L439:
	;
	v2346 = F_RelationGetPartitionQual(m, v219)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	if v2346 == int32(0) {
		goto L307
	} else {
		goto L441
	}
L441:
	;
	v2350 = F_expression_planner(m, v2346)
	mBase = m.M
	v2351 = m.ExcPending
	if v2351 != 0 {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v43)+76))
	if v2352 != int32(1) {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	F_ChangeVarNodes(m, v2350, int32(1), v2352)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L4
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+272)) = v2350
	goto L307
L446:
	;
	goto L445
L447:
	;
	m.G0 = v215 + int32(48)
	goto L47
L448:
	;
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_10), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(2175), int32(_a_F_build_simple_rel_11))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L4
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_10), int32(0))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L4
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(2187), int32(_a_F_build_simple_rel_11))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L4
	} else {
		goto L453
	}
L453:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+16)) = v1342
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_12), v215+int32(16))
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L4
	} else {
		goto L455
	}
L455:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(1736), int32(_a_F_build_simple_rel_13))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L4
	} else {
		goto L456
	}
L456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L457:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2439 = m.ExcPending
	if v2439 != 0 {
		goto L4
	} else {
		goto L458
	}
L458:
	;
	F_errmsg(m, int32(_a_F_build_simple_rel_14), int32(0))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L4
	} else {
		goto L459
	}
L459:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(543), int32(_a_F_build_simple_rel_5))
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L4
	} else {
		goto L460
	}
L460:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L461:
	;
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_15), int32(0))
	mBase = m.M
	v2456 = m.ExcPending
	if v2456 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(2838), int32(_a_F_build_simple_rel_16))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2468 = m.ExcPending
	if v2468 != 0 {
		goto L4
	} else {
		goto L465
	}
L465:
	;
	F_errmsg(m, int32(_a_F_build_simple_rel_17), int32(0))
	mBase = m.M
	v2472 = m.ExcPending
	if v2472 != 0 {
		goto L4
	} else {
		goto L466
	}
L466:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_4), int32(159), int32(_a_F_build_simple_rel_5))
	mBase = m.M
	v2477 = m.ExcPending
	if v2477 != 0 {
		goto L4
	} else {
		goto L467
	}
L467:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L468:
	;
	m.T0[v2506].(func(*base.Module, int32, int32, int32))(m, l0, v43, v41)
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L4
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	if l2 == int32(0) {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	goto L470
L472:
	;
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v2962+l1<<(uint(int32(2))%32)))) = v43
	m.G0 = v30 + int32(32)
	return v43
L473:
	;
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2515 = *(*int32)(unsafe.Add(mBase, uint32(v2511+l1<<(uint(int32(2))%32))))
	v2516 = m.G0
	v2518 = v2516 - int32(16)
	m.G0 = v2518
	*(*int32)(unsafe.Add(mBase, uint32(v2518)+12)) = v2515
	v2521 = *(*int32)(unsafe.Add(mBase, uint32(l2)+204))
	if v2521 == int32(0) {
		goto L476
	} else {
		goto L477
	}
L474:
	;
	m.G0 = v2518 + int32(16)
	if v2905 != 0 {
		goto L472
	} else {
		goto L529
	}
L475:
	;
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v41)+128))
	if v2707 == int32(0) {
		v2877 = v2684
		v2880 = v2687
		goto L507
	} else {
		goto L508
	}
L476:
	;
	v2684 = int32(-1)
	v2687 = int32(0)
	goto L475
L477:
	;
	goto L478
L478:
	;
	v2526 = int32(0)
	v2527 = int32(-1)
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2521)+4))
	if v2528 <= v2526 {
		v2684 = v2527
		v2687 = v2526
		goto L475
	} else {
		goto L479
	}
L479:
	;
	v2535 = v2527
	v2538 = v2526
	v2554 = v4
	goto L480
L480:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2521)+12))
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v2558+v2554<<(uint(int32(2))%32))))
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2562)+4))
	v2567 = F_adjust_appendrel_attrs(m, l0, v2563, int32(1), v2518+int32(12))
	mBase = m.M
	v2568 = m.ExcPending
	if v2568 != 0 {
		goto L4
	} else {
		goto L484
	}
L481:
	;
	v2684 = v2653
	v2687 = v2656
	goto L475
L482:
	;
	v2677 = v2554 + int32(1)
	v2678 = *(*int32)(unsafe.Add(mBase, uint32(v2521)+4))
	if v2677 < v2678 {
		v2535 = v2653
		v2538 = v2656
		v2554 = v2677
		goto L480
	} else {
		goto L506
	}
L483:
	;
	v2581 = F_make_ands_implicit(m, v2569)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L4
	} else {
		goto L490
	}
L484:
	;
	v2569 = F_eval_const_expressions(m, l0, v2567)
	mBase = m.M
	v2570 = m.ExcPending
	if v2570 != 0 {
		goto L4
	} else {
		goto L485
	}
L485:
	;
	if v2569 == int32(0) {
		goto L483
	} else {
		goto L486
	}
L486:
	;
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2569)))
	if v2573 != int32(7) {
		goto L483
	} else {
		goto L487
	}
L487:
	;
	v2576 = int32(0)
	v2577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2569)+32)))
	if v2577 != 0 {
		v2905 = v2576
		goto L474
	} else {
		goto L488
	}
L488:
	;
	v2578 = *(*int64)(unsafe.Add(mBase, uint32(v2569)+24))
	if v2578 != int64(0) {
		v2653 = v2535
		v2656 = v2538
		goto L482
	} else {
		goto L489
	}
L489:
	;
	v2905 = v2576
	goto L474
L490:
	;
	if v2581 == int32(0) {
		v2653 = v2535
		v2656 = v2538
		goto L482
	} else {
		goto L491
	}
L491:
	;
	v2585 = int32(0)
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+4))
	if v2586 <= v2585 {
		v2653 = v2535
		v2656 = v2538
		goto L482
	} else {
		goto L492
	}
L492:
	;
	v2592 = v2585
	v2593 = v2535
	v2596 = v2538
	goto L493
L493:
	;
	v2616 = int32(0)
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+12))
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2617+v2592<<(uint(int32(2))%32))))
	v2623 = F_contain_vars_of_level(m, v2621, v2616)
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L4
	} else {
		goto L496
	}
L494:
	;
	v2653 = v2644
	v2656 = v2640
	goto L482
L495:
	;
	v2631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2562)+8)))
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2562)+11)))
	v2633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2562)+12)))
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v2562)+20))
	v2635 = int32(0)
	v2638 = F_make_restrictinfo(m, l0, v2621, v2631, v2632, v2633, v2630, v2634, v2635, v2635, v2635)
	mBase = m.M
	v2639 = m.ExcPending
	if v2639 != 0 {
		goto L4
	} else {
		goto L500
	}
L496:
	;
	if v2623 != 0 {
		v2630 = v2616
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v2625 = F_contain_volatile_functions(m, v2621)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L4
	} else {
		goto L498
	}
L498:
	;
	if v2625 != 0 {
		v2630 = v2616
		goto L495
	} else {
		goto L499
	}
L499:
	;
	v2627 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+335)) = uint8(v2627)
	v2630 = v2627
	goto L495
L500:
	;
	v2640 = F_lappend(m, v2596, v2638)
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L4
	} else {
		goto L501
	}
L501:
	;
	v2642 = *(*int32)(unsafe.Add(mBase, uint32(v2638)+20))
	if base.Ui32(v2593) < base.Ui32(v2642) {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2644 = v2593
	goto L504
L503:
	;
	v2644 = v2642
	goto L504
L504:
	;
	v2646 = v2592 + int32(1)
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2581)+4))
	if v2646 < v2647 {
		v2592 = v2646
		v2593 = v2644
		v2596 = v2640
		goto L493
	} else {
		goto L505
	}
L505:
	;
	goto L494
L506:
	;
	goto L481
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+224)) = v2877
	*(*int32)(unsafe.Add(mBase, uint32(v43)+204)) = v2880
	v2905 = int32(1)
	goto L474
L508:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2707)+4))
	if v2710 <= int32(0) {
		v2877 = v2684
		v2880 = v2687
		goto L507
	} else {
		goto L509
	}
L509:
	;
	v2716 = v2710
	v2718 = v2684
	v2720 = int32(0)
	v2721 = v2687
	goto L510
L510:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(v2707)+12))
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v2741+v2720<<(uint(int32(2))%32))))
	if v2745 != 0 {
		goto L512
	} else {
		goto L513
	}
L511:
	;
	v2877 = v2847
	v2880 = v2850
	goto L507
L512:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+4))
	if v2746 <= int32(0) {
		v2819 = v2718
		v2822 = v2721
		goto L515
	} else {
		goto L516
	}
L513:
	;
	v2845 = v2716
	v2847 = v2718
	v2850 = v2721
	goto L514
L514:
	;
	v2871 = v2720 + int32(1)
	if v2871 < v2845 {
		v2716 = v2845
		v2718 = v2847
		v2720 = v2871
		v2721 = v2850
		goto L510
	} else {
		goto L528
	}
L515:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v2707)+4))
	v2845 = v2842
	v2847 = v2819
	v2850 = v2822
	goto L514
L516:
	;
	if base.Ui32(v2718) < base.Ui32(v2720) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2750 = v2718
	goto L519
L518:
	;
	v2750 = v2720
	goto L519
L519:
	;
	v2751 = int32(1)
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+12))
	v2753 = *(*int32)(unsafe.Add(mBase, uint32(v2752)))
	v2755 = int32(0)
	v2761 = F_make_restrictinfo(m, l0, v2753, v2751, v2755, v2755, v2755, v2720, v2755, v2755, v2755)
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L4
	} else {
		goto L520
	}
L520:
	;
	v2763 = F_lappend(m, v2721, v2761)
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L4
	} else {
		goto L521
	}
L521:
	;
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+4))
	if v2765 < int32(2) {
		v2819 = v2750
		v2822 = v2763
		goto L515
	} else {
		goto L522
	}
L522:
	;
	v2770 = v2751
	v2775 = v2763
	goto L523
L523:
	;
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+12))
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2795+v2770<<(uint(int32(2))%32))))
	v2801 = int32(0)
	v2807 = F_make_restrictinfo(m, l0, v2799, int32(1), v2801, v2801, v2801, v2720, v2801, v2801, v2801)
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L4
	} else {
		goto L525
	}
L524:
	;
	v2819 = v2750
	v2822 = v2809
	goto L515
L525:
	;
	v2809 = F_lappend(m, v2775, v2807)
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L4
	} else {
		goto L526
	}
L526:
	;
	v2812 = v2770 + int32(1)
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v2745)+4))
	if v2812 < v2813 {
		v2770 = v2812
		v2775 = v2809
		goto L523
	} else {
		goto L527
	}
L527:
	;
	goto L524
L528:
	;
	goto L511
L529:
	;
	F_mark_dummy_rel(m, v43)
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L4
	} else {
		goto L530
	}
L530:
	;
	goto L472
L531:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_build_simple_rel_18), v30+int32(16))
	mBase = m.M
	v2980 = m.ExcPending
	if v2980 != 0 {
		goto L4
	} else {
		goto L532
	}
L532:
	;
	F_errfinish(m, int32(_a_F_build_simple_rel_1), int32(220), int32(_a_F_build_simple_rel_2))
	mBase = m.M
	v2985 = m.ExcPending
	if v2985 != 0 {
		goto L4
	} else {
		goto L533
	}
L533:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_simple_union_all_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0
	goto L2
L1:
	;
	m.G0 = v8 + int32(16)
	return v65
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	return int32(0)
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v19 != int32(142) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v19 != int32(63) {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v39 != int32(1) {
		v65 = v38
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v36 = F_tlist_same_datatypes(m, v34, l2, int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v65 = v36
	goto L1
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
	if v42 != int32(1) {
		v65 = v38
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v46 = F_is_simple_union_all_recurse(m, v45, l1, l2)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v46 == int32(0) {
		v65 = v38
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v10 = v50
	goto L2
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v55
	F_errmsg_internal(m, int32(_a_F_is_simple_union_all_recurse_0), v8)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_is_simple_union_all_recurse_1), int32(2424), int32(_a_F_is_simple_union_all_recurse_2))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_simple_heap_delete(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int32(0)
		v16 = F_heap_delete(m, l0, l1, v9, v11, v11, int32(1), v6+int32(12))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				switch v16 - int32(2) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_delete_0), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_delete_1), int32(3234), int32(_a_F_simple_heap_delete_2))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 1:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_delete_3), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_delete_1), int32(3242), int32(_a_F_simple_heap_delete_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 2:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_delete_4), int32(0))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_delete_1), int32(3246), int32(_a_F_simple_heap_delete_2))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v16
						F_errmsg_internal(m, int32(_a_F_simple_heap_delete_5), v6)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_delete_1), int32(3250), int32(_a_F_simple_heap_delete_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				m.G0 = v6 + int32(32)
				return
			}
		}
	}
}
func F_simple_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = int32(0)
		v20 = F_heap_update(m, l0, l1, l2, v11, v13, v13, int32(1), v8+int32(12), v8+int32(8), l3)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 != 0 {
				switch v20 - int32(2) {
				case 0:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_update_0), int32(0))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_update_1), int32(_a_F_simple_heap_update_2), int32(_a_F_simple_heap_update_3))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 1:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_update_4), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_update_1), int32(_a_F_simple_heap_update_5), int32(_a_F_simple_heap_update_3))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 2:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_simple_heap_update_6), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_update_1), int32(_a_F_simple_heap_update_7), int32(_a_F_simple_heap_update_3))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
						F_errmsg_internal(m, int32(_a_F_simple_heap_update_8), v8)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_simple_heap_update_1), int32(_a_F_simple_heap_update_9), int32(_a_F_simple_heap_update_3))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	}
}
