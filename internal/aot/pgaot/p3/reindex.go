package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexRelationConcurrently(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
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
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v571 int32
	_ = v571
	var v592 int32
	_ = v592
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v624 int32
	_ = v624
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v806 int64
	_ = v806
	var v808 int64
	_ = v808
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v877 int64
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v891 int64
	_ = v891
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v905 int64
	_ = v905
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v919 int64
	_ = v919
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int64
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int64
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1076 int32
	_ = v1076
	var v1081 int64
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int64
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1151 int32
	_ = v1151
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int64
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1210 int32
	_ = v1210
	var v1219 int32
	_ = v1219
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1370 int32
	_ = v1370
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
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
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1485 int32
	_ = v1485
	var v1493 int64
	_ = v1493
	var v1495 int64
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1511 int32
	_ = v1511
	var v1515 int32
	_ = v1515
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1564 int64
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1578 int64
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1592 int64
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1599 int32
	_ = v1599
	var v1606 int64
	_ = v1606
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1739 int32
	_ = v1739
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1774 int32
	_ = v1774
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1784 int32
	_ = v1784
	var v1810 int32
	_ = v1810
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1844 int32
	_ = v1844
	var v1848 int32
	_ = v1848
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1901 int32
	_ = v1901
	var v1905 int64
	_ = v1905
	var v1909 int64
	_ = v1909
	var v1911 int64
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	var v1949 int32
	_ = v1949
	var v1955 int32
	_ = v1955
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1970 int32
	_ = v1970
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1980 int64
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1994 int64
	_ = v1994
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2008 int64
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2022 int64
	_ = v2022
	var v2024 int32
	_ = v2024
	var v2027 int32
	_ = v2027
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2086 int32
	_ = v2086
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2116 int32
	_ = v2116
	var v2120 int32
	_ = v2120
	var v2125 int32
	_ = v2125
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2153 int32
	_ = v2153
	var v2159 int32
	_ = v2159
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2167 int32
	_ = v2167
	var v2172 int32
	_ = v2172
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2202 int32
	_ = v2202
	var v2203 int32
	_ = v2203
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2270 int32
	_ = v2270
	var v2274 int32
	_ = v2274
	var v2277 int32
	_ = v2277
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2325 int32
	_ = v2325
	var v2330 int32
	_ = v2330
	var v2334 int32
	_ = v2334
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2360 int32
	_ = v2360
	var v2361 int32
	_ = v2361
	var v2367 int32
	_ = v2367
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2380 int32
	_ = v2380
	var v2408 int32
	_ = v2408
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2478 int32
	_ = v2478
	var v2480 int32
	_ = v2480
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2541 int32
	_ = v2541
	var v2569 int32
	_ = v2569
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2583 int32
	_ = v2583
	var v2585 int32
	_ = v2585
	var v2586 int32
	_ = v2586
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2632 int32
	_ = v2632
	var v2660 int32
	_ = v2660
	var v2664 int32
	_ = v2664
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2702 int32
	_ = v2702
	var v2703 int32
	_ = v2703
	var v2704 int32
	_ = v2704
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2718 int32
	_ = v2718
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2754 int32
	_ = v2754
	var v2755 int32
	_ = v2755
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2760 int32
	_ = v2760
	var v2761 int32
	_ = v2761
	var v2768 int32
	_ = v2768
	var v2773 int32
	_ = v2773
	var v2776 int32
	_ = v2776
	var v2777 int32
	_ = v2777
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2854 int32
	_ = v2854
	var v2855 int32
	_ = v2855
	var v2857 int64
	_ = v2857
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2862 int64
	_ = v2862
	var v2864 int32
	_ = v2864
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2873 int32
	_ = v2873
	var v2874 int32
	_ = v2874
	var v2875 int32
	_ = v2875
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2880 int32
	_ = v2880
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2893 int32
	_ = v2893
	var v2897 int32
	_ = v2897
	var v2899 int32
	_ = v2899
	var v2901 int32
	_ = v2901
	var v2904 int32
	_ = v2904
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2934 int32
	_ = v2934
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2949 int32
	_ = v2949
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2972 int32
	_ = v2972
	var v2980 int32
	_ = v2980
	var v2988 int32
	_ = v2988
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2995 int32
	_ = v2995
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3029 int32
	_ = v3029
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3037 int32
	_ = v3037
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3045 int32
	_ = v3045
	var v3072 int32
	_ = v3072
	var v3075 int32
	_ = v3075
	var v3079 int32
	_ = v3079
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3094 int32
	_ = v3094
	var v3098 int32
	_ = v3098
	var v3128 int32
	_ = v3128
	var v3132 int32
	_ = v3132
	var v3133 int64
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3157 int32
	_ = v3157
	var v3159 int32
	_ = v3159
	var v3162 int32
	_ = v3162
	var v3163 int32
	_ = v3163
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3198 int32
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3247 int64
	_ = v3247
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3266 int32
	_ = v3266
	var v3274 int32
	_ = v3274
	var v3277 int32
	_ = v3277
	var v3282 int32
	_ = v3282
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3307 int32
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3328 int32
	_ = v3328
	var v3332 int32
	_ = v3332
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3346 int64
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3357 int64
	_ = v3357
	var v3359 int32
	_ = v3359
	var v3360 int32
	_ = v3360
	var v3361 int32
	_ = v3361
	var v3367 int32
	_ = v3367
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3383 int32
	_ = v3383
	var v3385 int32
	_ = v3385
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3395 int32
	_ = v3395
	var v3399 int32
	_ = v3399
	var v3409 int32
	_ = v3409
	var v3425 int32
	_ = v3425
	var v3426 int32
	_ = v3426
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3435 int32
	_ = v3435
	var v3437 int32
	_ = v3437
	var v3439 int32
	_ = v3439
	var v3440 int32
	_ = v3440
	var v3441 int32
	_ = v3441
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3479 int32
	_ = v3479
	var v3485 int32
	_ = v3485
	var v3488 int32
	_ = v3488
	var v3491 int32
	_ = v3491
	var v3494 int32
	_ = v3494
	var v3497 int32
	_ = v3497
	var v3500 int32
	_ = v3500
	var v3507 int32
	_ = v3507
	var v3511 int32
	_ = v3511
	var v3516 int32
	_ = v3516
	var v3520 int32
	_ = v3520
	var v3526 int32
	_ = v3526
	var v3531 int32
	_ = v3531
	var v3535 int32
	_ = v3535
	var v3541 int32
	_ = v3541
	var v3546 int32
	_ = v3546
	var v3550 int32
	_ = v3550
	var v3556 int32
	_ = v3556
	var v3561 int32
	_ = v3561
	var v3565 int32
	_ = v3565
	var v3571 int32
	_ = v3571
	var v3576 int32
	_ = v3576
	var v3578 int32
	_ = v3578
	var v3579 int32
	_ = v3579
	var v3581 int32
	_ = v3581
	var v3585 int32
	_ = v3585
	var v3588 int32
	_ = v3588
	var v3589 int32
	_ = v3589
	var v3604 int32
	_ = v3604
	var v3623 int32
	_ = v3623
	var v3629 int32
	_ = v3629
	var v3632 int32
	_ = v3632
	var v3633 int32
	_ = v3633
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3643 int32
	_ = v3643
	var v3674 int32
	_ = v3674
	var v3677 int32
	_ = v3677
	var v3681 int32
	_ = v3681
	var v3686 int32
	_ = v3686
	var v3689 int32
	_ = v3689
	var v3691 int32
	_ = v3691
	var v3692 int32
	_ = v3692
	var v3695 int32
	_ = v3695
	var v3699 int32
	_ = v3699
	var v3701 int32
	_ = v3701
	var v3702 int32
	_ = v3702
	var v3710 int32
	_ = v3710
	var v3711 int32
	_ = v3711
	var v3717 int32
	_ = v3717
	var v3724 int32
	_ = v3724
	var v3757 int32
	_ = v3757
	var v3761 int32
	_ = v3761
	var v3766 int32
	_ = v3766
	var v3770 int32
	_ = v3770
	var v3773 int32
	_ = v3773
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3776 int32
	_ = v3776
	var v3780 int32
	_ = v3780
	var v3785 int32
	_ = v3785
	v4 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(416)
	m.G0 = v32
	v35 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+232)) = v35
	v38 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+224)) = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[2]))
	v46 = F_AllocSetContextCreateInternal(m, v41, int32(_a_F_ReindexRelationConcurrently_0), v4, int32(1024), int32(_a_F_ReindexRelationConcurrently_1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v50&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v53 = int32(_a_F_ReindexRelationConcurrently_2)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v57 = F_get_rel_name(m, l1)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v72 = v4
	v73 = v4
	goto L5
L5:
	;
	v74 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L6:
	;
	v59 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v61 = F_get_namespace_name(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_getrusage(m, v32+int32(264))
	mBase = m.M
	F_gettimeofday(m, v32+int32(248))
	mBase = m.M
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v54
	v72 = v57
	v73 = v61
	goto L5
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3770 = m.ExcPending
	if v3770 != 0 {
		goto L1
	} else {
		goto L653
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3757 = m.ExcPending
	if v3757 != 0 {
		goto L1
	} else {
		goto L650
	}
L12:
	;
	m.G0 = v32 + int32(416)
	return v3724
L13:
	;
	if v635 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L14:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+112))
	if v449 != 0 {
		goto L136
	} else {
		goto L137
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L132
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L127
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L123
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L119
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L115
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L111
	}
L21:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v251 = F_IndexGetRelation(m, l1, int32(base.Ui32(v246&int32(4))>>(uint(int32(2))%32)))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L78
	}
L22:
	;
	v80 = int32(_a_F_ReindexRelationConcurrently_2)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v85 = F_lappend_oid(m, int32(0), l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	switch v74&int32(255) - int32(105) {
	case 0:
		goto L21
	default:
		goto L15
	case 4, 9, 11:
		goto L22
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v81
	goto L25
L25:
	;
	if base.Ui32(l1) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v91&int32(4) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v101 != 0 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v95 = F_try_table_open(m, l1, int32(4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v98 = F_table_open(m, l1, int32(4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	if v95 != 0 {
		v100 = v95
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v3724 = v4
	goto L12
L33:
	;
	v100 = v98
	goto L27
L34:
	;
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+56))
	if base.Ui32(v104) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		v113 = v103
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v114 = F_RelationGetIndexList(m, v100)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	if v113 != 0 {
		goto L19
	} else {
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+68))
	if v108 == int32(99) {
		v113 = v103
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v111 = F_isTempToastNamespace(m, v108)
	mBase = m.M
	v113 = v111
	goto L38
L41:
	;
	goto L36
L42:
	;
	if v114 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v429 = v4
	goto L14
L44:
	;
	goto L45
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v118 <= int32(0) {
		v429 = v4
		goto L14
	} else {
		goto L46
	}
L46:
	;
	v123 = int32(0)
	v132 = v4
	goto L47
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v123<<(uint(int32(2))%32))))
	v157 = F_index_open(m, v155, int32(4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v429 = v238
	goto L14
L49:
	;
	F_relation_close(m, v157, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L76
	}
L50:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v157)+192))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+18)))
	if v160 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v165 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+15)))
	if v194 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L54:
	;
	if v165 == int32(0) {
		v238 = v132
		goto L49
	} else {
		goto L55
	}
L55:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v172 = F_get_rel_namespace(m, v155)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v174 = F_get_namespace_name(m, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v176 = F_get_rel_name(m, v155)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+132)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v174
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_4), v32+int32(128))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errhint(m, int32(_a_F_ReindexRelationConcurrently_5), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3729), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v238 = v132
	goto L49
L63:
	;
	v199 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v224 = int32(_a_F_ReindexRelationConcurrently_2)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v229 = F_palloc(m, int32(16))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L74
	}
L66:
	;
	if v199 == int32(0) {
		v238 = v132
		goto L49
	} else {
		goto L67
	}
L67:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v206 = F_get_rel_namespace(m, v155)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v208 = F_get_namespace_name(m, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v210 = F_get_rel_name(m, v155)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v208
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_8), v32+int32(112))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3735), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v238 = v132
	goto L49
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v155
	v232 = F_lappend(m, v132, v229)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v225
	v238 = v232
	goto L49
L76:
	;
	v243 = v123 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v243 < v244 {
		v123 = v243
		v132 = v238
		goto L47
	} else {
		goto L77
	}
L77:
	;
	goto L48
L78:
	;
	if v251 == int32(0) {
		v3724 = v4
		goto L12
	} else {
		goto L79
	}
L79:
	;
	goto L80
L80:
	;
	if base.Ui32(v251) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		goto L18
	} else {
		goto L81
	}
L81:
	;
	v257 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v257 != int32(99) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v263 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v261 = F_isTempToastNamespace(m, v257)
	mBase = m.M
	v263 = v261
	goto L86
L85:
	;
	v263 = int32(1)
	goto L86
L86:
	;
	goto L83
L87:
	;
	v264 = F_get_index_isvalid(m, l1)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v268&int32(4) != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	if v264 == int32(0) {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v278 != 0 {
		goto L99
	} else {
		goto L100
	}
L93:
	;
	v272 = F_try_table_open(m, v251, int32(4))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v275 = F_table_open(m, v251, int32(4))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	if v272 != 0 {
		v277 = v272
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v3724 = v4
	goto L12
L98:
	;
	v277 = v275
	goto L92
L99:
	;
	v280 = int32(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v277)+56))
	if base.Ui32(v281) < base.Ui32(int32(_a_F_ReindexRelationConcurrently_3)) {
		v290 = v280
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	F_relation_close(m, v277, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L107
	}
L102:
	;
	if v290 != 0 {
		goto L16
	} else {
		goto L106
	}
L103:
	;
	goto L102
L104:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v277)+48))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+68))
	if v285 == int32(99) {
		v290 = v280
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v288 = F_isTempToastNamespace(m, v285)
	mBase = m.M
	v290 = v288
	goto L103
L106:
	;
	goto L101
L107:
	;
	v294 = int32(_a_F_ReindexRelationConcurrently_2)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+188)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v32)+156)) = v251
	v303 = F_list_make1_impl(m, int32(480), v32+int32(156))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v306 = F_palloc(m, int32(16))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = l1
	v310 = F_lappend(m, int32(0), v306)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v295
	v635 = v310
	v636 = v303
	goto L13
L111:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_9), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3694), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v337 + int32(4)
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_10), v32+int32(144))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3714), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_9), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3824), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_11), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3835), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v390 = F_get_rel_name(m, l1)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = v390
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_10), v32+int32(160))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3860), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_12), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3889), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	v452 = F_table_open(m, v449, int32(4))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	v603 = v429
	v604 = v85
	goto L138
L138:
	;
	F_relation_close(m, v100, int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L166
	}
L139:
	;
	v454 = int32(_a_F_ReindexRelationConcurrently_2)
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v458 = F_lappend_oid(m, v85, v449)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v455
	v462 = F_RelationGetIndexList(m, v452)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	F_relation_close(m, v452, int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L165
	}
L142:
	;
	if v462 == int32(0) {
		v571 = v429
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	if v466 <= int32(0) {
		v571 = v429
		goto L141
	} else {
		goto L144
	}
L144:
	;
	v470 = int32(0)
	v479 = v429
	goto L145
L145:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v462)+12))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v498+v470<<(uint(int32(2))%32))))
	v504 = F_index_open(m, v502, int32(4))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v571 = v552
	goto L141
L147:
	;
	F_relation_close(m, v504, int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L163
	}
L148:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504)+192))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506)+18)))
	if v507 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v508 = int32(_a_F_ReindexRelationConcurrently_2)
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v513 = F_palloc(m, int32(16))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v522 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v502
	v516 = F_lappend(m, v479, v513)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v509
	v552 = v516
	goto L147
L154:
	;
	if v522 == int32(0) {
		v552 = v479
		goto L147
	} else {
		goto L155
	}
L155:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v529 = F_get_rel_namespace(m, v502)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v531 = F_get_namespace_name(m, v529)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v533 = F_get_rel_name(m, v502)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v531
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_4), v32+int32(96))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errhint(m, int32(_a_F_ReindexRelationConcurrently_5), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3782), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v552 = v479
	goto L147
L163:
	;
	v558 = v470 + int32(1)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v462)+4))
	if v558 < v559 {
		v470 = v558
		v479 = v552
		goto L145
	} else {
		goto L164
	}
L164:
	;
	goto L146
L165:
	;
	v603 = v571
	v604 = v458
	goto L138
L166:
	;
	v635 = v603
	v636 = v604
	goto L13
L167:
	;
	v3724 = int32(0)
	goto L12
L168:
	;
	goto L169
L169:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v657 == int32(1664) {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v660 = int32(0)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v660 < v662 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v671 = int32(0)
	v674 = v660
	v675 = v660
	goto L174
L172:
	;
	v1106 = v660
	v1107 = v660
	goto L173
L173:
	;
	if v636 == int32(0) {
		v1210 = v1107
		v1219 = v4
		goto L235
	} else {
		goto L236
	}
L174:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v695+v671<<(uint(int32(2))%32))))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v702 = F_index_open(m, v700, int32(4))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L176
	}
L175:
	;
	v1106 = v1037
	v1107 = v1051
	goto L173
L176:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v702)+192))
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	v707 = F_table_open(m, v705, int32(4))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(184)))) = v714
	v717 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(180)))) = v717
	goto L178
L178:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v707)+48))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+80))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v32)+180))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5])) = v721 | int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4])) = v720
	goto L179
L179:
	;
	v729 = int32(_a_F_ReindexRelationConcurrently_13)
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[6]))
	v733 = v731 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[6])) = v733
	goto L180
L180:
	;
	F_RestrictSearchPath(m)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v737 = F_RelationGetIndexExpressions(m, v702)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	if v737 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v742 = int32(1)
	goto L185
L184:
	;
	v740 = F_RelationGetIndexPredicate(m, v702)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L1
	} else {
		goto L186
	}
L185:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v699)+12)) = uint8(base.B2i32(v742 == int32(0)))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v707)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v699)+4)) = v746
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v748)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v699)+8)) = v749
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+118)))
	if v752 == int32(116) {
		goto L11
	} else {
		goto L187
	}
L186:
	;
	v742 = v740
	goto L185
L187:
	;
	v758 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v758 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = int64(4)
	v806 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v699))))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v806
	v808 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v699)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v808
	v812 = v32 + int32(224)
	v814 = v32 + int32(192)
	goto L194
L189:
	;
	goto L188
L190:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v762&int32(1) == int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v767 = int32(_a_F_ReindexRelationConcurrently_14)
	v769 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v770 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v769 + v770
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = v773 + v770
	v777 = int32(0)
	v779 = int32(_a_F_ReindexRelationConcurrently_15)
	v780 = base.AtomicRmwOr32(m, v777, v779, v777)
	*(*int32)(unsafe.Add(mBase, uint32(v758)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v758)+224)) = v746
	base.MemoryFill(m, v758+int32(232), v777, int32(160))
	v791 = base.AtomicRmwOr32(m, v777, v779, v777)
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = v792 + v770
	v798 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v798 - v770
	goto L189
L192:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v997 = F_get_rel_name(m, v996)
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L1
	} else {
		goto L209
	}
L193:
	;
	goto L192
L194:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v824 == int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v828&int32(1) == int32(0) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v833 = int32(_a_F_ReindexRelationConcurrently_14)
	v835 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v836 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v835 + v836
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v839 + v836
	v843 = int32(0)
	v846 = base.AtomicRmwOr32(m, v843, int32(_a_F_ReindexRelationConcurrently_15), v843)
	goto L198
L197:
	;
	v973 = int32(0)
	v976 = base.AtomicRmwOr32(m, v973, int32(_a_F_ReindexRelationConcurrently_15), v973)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	v978 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v977 + v978
	v981 = int32(_a_F_ReindexRelationConcurrently_14)
	v983 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v983 - v978
	goto L193
L198:
	;
	v852 = v824 + int32(232)
	goto L199
L199:
	;
	v858 = int32(0)
	v861 = int32(0)
	goto L202
L202:
	;
	v867 = int32(2)
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v812+v861<<(uint(v867)%32))))
	v871 = int32(3)
	v877 = *(*int64)(unsafe.Add(mBase, uint32(v814+v861<<(uint(v871)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v852+v870<<(uint(v871)%32)))) = v877
	v880 = v861 | int32(1)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v812+v880<<(uint(v867)%32))))
	v891 = *(*int64)(unsafe.Add(mBase, uint32(v814+v880<<(uint(v871)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v852+v884<<(uint(v871)%32)))) = v891
	v894 = v861 | v867
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v812+v894<<(uint(v867)%32))))
	v905 = *(*int64)(unsafe.Add(mBase, uint32(v814+v894<<(uint(v871)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v852+v898<<(uint(v871)%32)))) = v905
	v908 = v861 | v871
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v812+v908<<(uint(v867)%32))))
	v919 = *(*int64)(unsafe.Add(mBase, uint32(v814+v908<<(uint(v871)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v852+v912<<(uint(v871)%32)))) = v919
	v921 = int32(4)
	v924 = v858 + v921
	if v924 != int32(4) {
		v858 = v924
		v861 = v861 + v921
		goto L202
	} else {
		goto L204
	}
L203:
	;
	goto L197
L204:
	;
	goto L203
L209:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v702)+192))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	v1003 = F_get_rel_namespace(m, v1002)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1006 = F_ChooseRelationName(m, v997, int32(0), int32(_a_F_ReindexRelationConcurrently_16), v1003, int32(0))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v1008 != 0 {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v1018 = F_index_create_copy(m, v707, int32(140), v1017, v1015, v1006)
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L1
	} else {
		goto L217
	}
L213:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v707)+48))
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1009)+119)))
	if v1010 != int32(116) {
		v1015 = v1008
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v702)+48))
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1013)+92))
	v1015 = v1014
	goto L212
L216:
	;
	goto L215
L217:
	;
	v1021 = F_index_open(m, v1018, int32(4))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v1023 = int32(_a_F_ReindexRelationConcurrently_2)
	v1024 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v1028 = F_palloc(m, int32(16))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1028))) = v1018
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1028)+12)) = uint8(v1031)
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v699)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+4)) = v1033
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v699)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+8)) = v1035
	v1037 = F_lappend(m, v674, v1028)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v1040 = F_palloc(m, int32(8))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v1042 = *(*int64)(unsafe.Add(mBase, uint32(v702)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1040))) = v1042
	v1044 = F_lappend(m, v675, v1040)
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	v1047 = F_palloc(m, int32(8))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v1049 = *(*int64)(unsafe.Add(mBase, uint32(v1021)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1047))) = v1049
	v1051 = F_lappend(m, v1044, v1047)
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1024
	F_relation_close(m, v702, int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	F_relation_close(m, v1021, int32(0))
	mBase = m.M
	v1060 = m.ExcPending
	if v1060 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_AtEOXact_GUC(m, int32(0), v733)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v32)+184))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v32)+180))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[5])) = v1065
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[4])) = v1064
	goto L228
L228:
	;
	F_relation_close(m, v707, int32(0))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	if l0 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+172)) = v1018
	*(*int32)(unsafe.Add(mBase, uint32(v32)+168)) = int32(1259)
	v1076 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = v1076
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1076
	v1081 = *(*int64)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+64)) = v1081
	v1084 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+72)) = v1084
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v32)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v1086
	F_EventTriggerCollectSimpleCommand(m, v32+int32(80), v32-int32(-64), l0)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v1095 = v671 + int32(1)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v1095 < v1096 {
		v671 = v1095
		v674 = v1037
		v675 = v1051
		goto L174
	} else {
		goto L234
	}
L233:
	;
	goto L232
L234:
	;
	goto L175
L235:
	;
	if v1210 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L236:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v1129 <= int32(0) {
		v1210 = v1107
		v1219 = v4
		goto L235
	} else {
		goto L237
	}
L237:
	;
	v1136 = int32(0)
	v1142 = v1107
	v1151 = v4
	goto L238
L238:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v636)+12))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1162+v1136<<(uint(int32(2))%32))))
	v1168 = F_table_open(m, v1166, int32(4))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	v1210 = v1179
	v1219 = v1190
	goto L235
L240:
	;
	v1170 = int32(_a_F_ReindexRelationConcurrently_2)
	v1171 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v46
	v1175 = F_palloc(m, int32(8))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(v1168)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v1175))) = v1177
	v1179 = F_lappend(m, v1142, v1175)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	v1182 = F_palloc(m, int32(16))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1175)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1182))) = v1184
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1175)))
	*(*int64)(unsafe.Add(mBase, uint32(v1182)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v1182)+4)) = v1186
	v1190 = F_lappend(m, v1151, v1182)
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[3])) = v1171
	F_relation_close(m, v1168, int32(0))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1198 = v1136 + int32(1)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v636)+4))
	if v1198 < v1199 {
		v1136 = v1198
		v1142 = v1179
		v1151 = v1190
		goto L238
	} else {
		goto L246
	}
L246:
	;
	goto L239
L247:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L254
	}
L248:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	if v1232 <= int32(0) {
		goto L247
	} else {
		goto L249
	}
L249:
	;
	v1237 = int32(0)
	goto L250
L250:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+12))
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1237<<(uint(int32(2))%32))))
	F_LockRelationIdForSession(m, v1269, int32(4))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L252
	}
L251:
	;
	goto L247
L252:
	;
	v1274 = v1237 + int32(1)
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	if v1274 < v1275 {
		v1237 = v1274
		goto L250
	} else {
		goto L253
	}
L253:
	;
	goto L251
L254:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1316 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	F_WaitForLockersMultiple(m, v1219, int32(5), int32(1))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L261
	}
L258:
	;
	goto L257
L259:
	;
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1320&int32(1) == int32(0) {
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1325 = int32(_a_F_ReindexRelationConcurrently_14)
	v1327 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1328 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1327 + v1328
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1331 + v1328
	v1335 = int32(0)
	v1337 = int32(_a_F_ReindexRelationConcurrently_15)
	v1338 = base.AtomicRmwOr32(m, v1335, v1337, v1335)
	*(*int64)(unsafe.Add(mBase, uint32(v1316+int32(72))+232)) = int64(1)
	v1346 = base.AtomicRmwOr32(m, v1335, v1337, v1335)
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1316)))
	*(*int32)(unsafe.Add(mBase, uint32(v1316))) = v1347 + v1328
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1353 - v1328
	goto L258
L261:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	if v1106 != 0 {
		goto L264
	} else {
		goto L265
	}
L263:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2249 = m.ExcPending
	if v2249 != 0 {
		goto L1
	} else {
		goto L372
	}
L264:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if int32(0) < v1363 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2167 = m.ExcPending
	if v2167 != 0 {
		goto L1
	} else {
		goto L365
	}
L267:
	;
	v1370 = int32(0)
	goto L270
L268:
	;
	goto L269
L269:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L1
	} else {
		goto L309
	}
L270:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+12))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1396+v1370<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L1
	} else {
		goto L272
	}
L271:
	;
	goto L269
L272:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v1404 != 0 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400)+12)))
	if v1407 == int32(1) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L275
L277:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v1415 = F_LWLockAcquire(m, v1411+int32(512), int32(0))
	mBase = m.M
	v1416 = m.ExcPending
	if v1416 != 0 {
		goto L1
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1437 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L282
	}
L280:
	;
	v1418 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1418)+36)))
	v1421 = v1419 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1418)+36)) = uint8(v1421)
	v1424 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+12))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1418)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v1425+v1426))) = uint8(v1421)
	v1430 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v1430+int32(512))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	goto L279
L282:
	;
	F_PushActiveSnapshot(m, v1437)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+4))
	v1445 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1445 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = int64(2)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = int64(4)
	v1493 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1400))))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v1493
	v1495 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1400)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v1495
	v1499 = v32 + int32(224)
	v1501 = v32 + int32(192)
	goto L290
L285:
	;
	goto L284
L286:
	;
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1449&int32(1) == int32(0) {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1454 = int32(_a_F_ReindexRelationConcurrently_14)
	v1456 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1457 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1456 + v1457
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	*(*int32)(unsafe.Add(mBase, uint32(v1445))) = v1460 + v1457
	v1464 = int32(0)
	v1466 = int32(_a_F_ReindexRelationConcurrently_15)
	v1467 = base.AtomicRmwOr32(m, v1464, v1466, v1464)
	*(*int32)(unsafe.Add(mBase, uint32(v1445)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v1445)+224)) = v1442
	base.MemoryFill(m, v1445+int32(232), v1464, int32(160))
	v1478 = base.AtomicRmwOr32(m, v1464, v1466, v1464)
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	*(*int32)(unsafe.Add(mBase, uint32(v1445))) = v1479 + v1457
	v1485 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1485 - v1457
	goto L285
L288:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1400)+4))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1400)))
	F_index_concurrently_build(m, v1683, v1684)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L1
	} else {
		goto L305
	}
L289:
	;
	goto L288
L290:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1511 == int32(0) {
		goto L289
	} else {
		goto L291
	}
L291:
	;
	v1515 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1515&int32(1) == int32(0) {
		goto L289
	} else {
		goto L292
	}
L292:
	;
	v1520 = int32(_a_F_ReindexRelationConcurrently_14)
	v1522 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1523 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1522 + v1523
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	*(*int32)(unsafe.Add(mBase, uint32(v1511))) = v1526 + v1523
	v1530 = int32(0)
	v1533 = base.AtomicRmwOr32(m, v1530, int32(_a_F_ReindexRelationConcurrently_15), v1530)
	goto L294
L293:
	;
	v1660 = int32(0)
	v1663 = base.AtomicRmwOr32(m, v1660, int32(_a_F_ReindexRelationConcurrently_15), v1660)
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1665 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1511))) = v1664 + v1665
	v1668 = int32(_a_F_ReindexRelationConcurrently_14)
	v1670 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1670 - v1665
	goto L289
L294:
	;
	v1539 = v1511 + int32(232)
	goto L295
L295:
	;
	v1545 = int32(0)
	v1548 = int32(0)
	goto L298
L298:
	;
	v1554 = int32(2)
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1548<<(uint(v1554)%32))))
	v1558 = int32(3)
	v1564 = *(*int64)(unsafe.Add(mBase, uint32(v1501+v1548<<(uint(v1558)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1539+v1557<<(uint(v1558)%32)))) = v1564
	v1567 = v1548 | int32(1)
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1567<<(uint(v1554)%32))))
	v1578 = *(*int64)(unsafe.Add(mBase, uint32(v1501+v1567<<(uint(v1558)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1539+v1571<<(uint(v1558)%32)))) = v1578
	v1581 = v1548 | v1554
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1581<<(uint(v1554)%32))))
	v1592 = *(*int64)(unsafe.Add(mBase, uint32(v1501+v1581<<(uint(v1558)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1539+v1585<<(uint(v1558)%32)))) = v1592
	v1595 = v1548 | v1558
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1499+v1595<<(uint(v1554)%32))))
	v1606 = *(*int64)(unsafe.Add(mBase, uint32(v1501+v1595<<(uint(v1558)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1539+v1599<<(uint(v1558)%32)))) = v1606
	v1608 = int32(4)
	v1611 = v1545 + v1608
	if v1611 != int32(4) {
		v1545 = v1611
		v1548 = v1548 + v1608
		goto L298
	} else {
		goto L300
	}
L299:
	;
	goto L293
L300:
	;
	goto L299
L305:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L1
	} else {
		goto L307
	}
L307:
	;
	v1692 = v1370 + int32(1)
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1692 < v1693 {
		v1370 = v1692
		goto L270
	} else {
		goto L308
	}
L308:
	;
	goto L271
L309:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1730 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L310:
	;
	F_WaitForLockersMultiple(m, v1219, int32(5), int32(1))
	mBase = m.M
	v1774 = m.ExcPending
	if v1774 != 0 {
		goto L1
	} else {
		goto L314
	}
L311:
	;
	goto L310
L312:
	;
	v1734 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1734&int32(1) == int32(0) {
		goto L311
	} else {
		goto L313
	}
L313:
	;
	v1739 = int32(_a_F_ReindexRelationConcurrently_14)
	v1741 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1742 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1741 + v1742
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1730)))
	*(*int32)(unsafe.Add(mBase, uint32(v1730))) = v1745 + v1742
	v1749 = int32(0)
	v1751 = int32(_a_F_ReindexRelationConcurrently_15)
	v1752 = base.AtomicRmwOr32(m, v1749, v1751, v1749)
	*(*int64)(unsafe.Add(mBase, uint32(v1730+int32(72))+232)) = int64(3)
	v1760 = base.AtomicRmwOr32(m, v1749, v1751, v1749)
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1730)))
	*(*int32)(unsafe.Add(mBase, uint32(v1730))) = v1761 + v1742
	v1767 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1767 - v1742
	goto L311
L314:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1777 = int32(0)
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v1778 <= v1777 {
		goto L263
	} else {
		goto L316
	}
L316:
	;
	v1784 = v1777
	goto L317
L317:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+12))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1810+v1784<<(uint(int32(2))%32))))
	F_StartTransactionCommand(m)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L1
	} else {
		goto L319
	}
L318:
	;
	goto L263
L319:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v1818 != 0 {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L1
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v1821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814)+12)))
	if v1821 == int32(1) {
		goto L324
	} else {
		goto L325
	}
L323:
	;
	goto L322
L324:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v1829 = F_LWLockAcquire(m, v1825+int32(512), int32(0))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L1
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1851 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L1
	} else {
		goto L329
	}
L327:
	;
	v1832 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v1833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1832)+36)))
	v1835 = v1833 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v1832)+36)) = uint8(v1835)
	v1838 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(v1838)+12))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1832)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v1839+v1840))) = uint8(v1835)
	v1844 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v1844+int32(512))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	goto L326
L329:
	;
	v1853 = F_RegisterSnapshot(m, v1851)
	mBase = m.M
	v1854 = m.ExcPending
	if v1854 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	F_PushActiveSnapshot(m, v1853)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+4))
	v1861 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1861 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v1905 = int64(4)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+200)) = v1905
	*(*int64)(unsafe.Add(mBase, uint32(v32)+192)) = v1905
	v1909 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1814))))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+208)) = v1909
	v1911 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1814)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+216)) = v1911
	v1915 = v32 + int32(224)
	v1917 = v32 + int32(192)
	goto L338
L333:
	;
	goto L332
L334:
	;
	v1865 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1865&int32(1) == int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v1870 = int32(_a_F_ReindexRelationConcurrently_14)
	v1872 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1873 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1872 + v1873
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1861)))
	*(*int32)(unsafe.Add(mBase, uint32(v1861))) = v1876 + v1873
	v1880 = int32(0)
	v1882 = int32(_a_F_ReindexRelationConcurrently_15)
	v1883 = base.AtomicRmwOr32(m, v1880, v1882, v1880)
	*(*int32)(unsafe.Add(mBase, uint32(v1861)+220)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v1861)+224)) = v1858
	base.MemoryFill(m, v1861+int32(232), v1880, int32(160))
	v1894 = base.AtomicRmwOr32(m, v1880, v1882, v1880)
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1861)))
	*(*int32)(unsafe.Add(mBase, uint32(v1861))) = v1895 + v1873
	v1901 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1901 - v1873
	goto L333
L336:
	;
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v1814)+4))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v1814)))
	F_validate_index(m, v2099, v2100, v1853)
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L353
	}
L337:
	;
	goto L336
L338:
	;
	v1927 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v1927 == int32(0) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v1931&int32(1) == int32(0) {
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v1936 = int32(_a_F_ReindexRelationConcurrently_14)
	v1938 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v1939 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v1938 + v1939
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1927)))
	*(*int32)(unsafe.Add(mBase, uint32(v1927))) = v1942 + v1939
	v1946 = int32(0)
	v1949 = base.AtomicRmwOr32(m, v1946, int32(_a_F_ReindexRelationConcurrently_15), v1946)
	goto L342
L341:
	;
	v2076 = int32(0)
	v2079 = base.AtomicRmwOr32(m, v2076, int32(_a_F_ReindexRelationConcurrently_15), v2076)
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v1927)))
	v2081 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1927))) = v2080 + v2081
	v2084 = int32(_a_F_ReindexRelationConcurrently_14)
	v2086 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2086 - v2081
	goto L337
L342:
	;
	v1955 = v1927 + int32(232)
	goto L343
L343:
	;
	v1961 = int32(0)
	v1964 = int32(0)
	goto L346
L346:
	;
	v1970 = int32(2)
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v1964<<(uint(v1970)%32))))
	v1974 = int32(3)
	v1980 = *(*int64)(unsafe.Add(mBase, uint32(v1917+v1964<<(uint(v1974)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1955+v1973<<(uint(v1974)%32)))) = v1980
	v1983 = v1964 | int32(1)
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v1983<<(uint(v1970)%32))))
	v1994 = *(*int64)(unsafe.Add(mBase, uint32(v1917+v1983<<(uint(v1974)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1955+v1987<<(uint(v1974)%32)))) = v1994
	v1997 = v1964 | v1970
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v1997<<(uint(v1970)%32))))
	v2008 = *(*int64)(unsafe.Add(mBase, uint32(v1917+v1997<<(uint(v1974)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1955+v2001<<(uint(v1974)%32)))) = v2008
	v2011 = v1964 | v1974
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v2011<<(uint(v1970)%32))))
	v2022 = *(*int64)(unsafe.Add(mBase, uint32(v1917+v2011<<(uint(v1974)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1955+v2015<<(uint(v1974)%32)))) = v2022
	v2024 = int32(4)
	v2027 = v1961 + v2024
	if v2027 != int32(4) {
		v1961 = v2027
		v1964 = v1964 + v2024
		goto L346
	} else {
		goto L348
	}
L347:
	;
	goto L341
L348:
	;
	goto L347
L353:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v1853)+4))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	F_UnregisterSnapshot(m, v1853)
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2116 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	F_WaitForOlderSnapshots(m, v2103, int32(1))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L362
	}
L359:
	;
	goto L358
L360:
	;
	v2120 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2120&int32(1) == int32(0) {
		goto L359
	} else {
		goto L361
	}
L361:
	;
	v2125 = int32(_a_F_ReindexRelationConcurrently_14)
	v2127 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2128 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2127 + v2128
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2116)))
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2131 + v2128
	v2135 = int32(0)
	v2137 = int32(_a_F_ReindexRelationConcurrently_15)
	v2138 = base.AtomicRmwOr32(m, v2135, v2137, v2135)
	*(*int64)(unsafe.Add(mBase, uint32(v2116+int32(72))+232)) = int64(7)
	v2146 = base.AtomicRmwOr32(m, v2135, v2137, v2135)
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2116)))
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2147 + v2128
	v2153 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2153 - v2128
	goto L359
L362:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2161 = m.ExcPending
	if v2161 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v2163 = v1784 + int32(1)
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v2163 < v2164 {
		v1784 = v2163
		goto L317
	} else {
		goto L364
	}
L364:
	;
	goto L318
L365:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2172 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	F_WaitForLockersMultiple(m, v1219, int32(5), int32(1))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L370
	}
L367:
	;
	goto L366
L368:
	;
	v2176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2176&int32(1) == int32(0) {
		goto L367
	} else {
		goto L369
	}
L369:
	;
	v2181 = int32(_a_F_ReindexRelationConcurrently_14)
	v2183 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2184 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2183 + v2184
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2172)))
	*(*int32)(unsafe.Add(mBase, uint32(v2172))) = v2187 + v2184
	v2191 = int32(0)
	v2193 = int32(_a_F_ReindexRelationConcurrently_15)
	v2194 = base.AtomicRmwOr32(m, v2191, v2193, v2191)
	*(*int64)(unsafe.Add(mBase, uint32(v2172+int32(72))+232)) = int64(3)
	v2202 = base.AtomicRmwOr32(m, v2191, v2193, v2191)
	v2203 = *(*int32)(unsafe.Add(mBase, uint32(v2172)))
	*(*int32)(unsafe.Add(mBase, uint32(v2172))) = v2203 + v2184
	v2209 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2209 - v2184
	goto L367
L370:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L1
	} else {
		goto L371
	}
L371:
	;
	goto L263
L372:
	;
	v2251 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	v2255 = F_LWLockAcquire(m, v2251+int32(512), int32(0))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[14]))
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2258)+36)))
	v2261 = v2259 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v2258)+36)) = uint8(v2261)
	v2264 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[15]))
	v2265 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+12))
	v2266 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v2265+v2266))) = uint8(v2261)
	v2270 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[13]))
	F_LWLockRelease(m, v2270+int32(512))
	mBase = m.M
	v2274 = m.ExcPending
	if v2274 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	v2277 = int32(0)
	goto L377
L375:
	;
	F_MemoryContextDelete(m, v46)
	mBase = m.M
	v3674 = m.ExcPending
	if v3674 != 0 {
		goto L1
	} else {
		goto L644
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v73
	F_errmsg(m, v3604, v32+int32(32))
	mBase = m.M
	v3629 = m.ExcPending
	if v3629 != 0 {
		goto L1
	} else {
		goto L640
	}
L377:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v2277 < v2306 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v3588 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v3589 = m.ExcPending
	if v3589 != 0 {
		goto L1
	} else {
		goto L638
	}
L379:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v2312 = v2308 + v2277<<(uint(int32(2))%32)
	goto L381
L380:
	;
	v2312 = int32(0)
	goto L381
L381:
	;
	if v1106 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L382:
	;
	goto L378
L383:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v2320+v2277<<(uint(int32(2))%32))))
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2312)))
	v2822 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2822 != 0 {
		goto L462
	} else {
		goto L463
	}
L384:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L1
	} else {
		goto L388
	}
L385:
	;
	v2317 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if base.B2i32(v2312 == int32(0))|base.B2i32(v2317 <= v2277) != 0 {
		goto L384
	} else {
		goto L386
	}
L386:
	;
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+12))
	if v2320 != 0 {
		goto L383
	} else {
		goto L387
	}
L387:
	;
	goto L384
L388:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2330 == int32(0) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	F_WaitForLockersMultiple(m, v1219, int32(8), int32(1))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L1
	} else {
		goto L394
	}
L391:
	;
	goto L390
L392:
	;
	v2334 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2334&int32(1) == int32(0) {
		goto L391
	} else {
		goto L393
	}
L393:
	;
	v2339 = int32(_a_F_ReindexRelationConcurrently_14)
	v2341 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2342 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2341 + v2342
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2330)))
	*(*int32)(unsafe.Add(mBase, uint32(v2330))) = v2345 + v2342
	v2349 = int32(0)
	v2351 = int32(_a_F_ReindexRelationConcurrently_15)
	v2352 = base.AtomicRmwOr32(m, v2349, v2351, v2349)
	*(*int64)(unsafe.Add(mBase, uint32(v2330+int32(72))+232)) = int64(8)
	v2360 = base.AtomicRmwOr32(m, v2349, v2351, v2349)
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(v2330)))
	*(*int32)(unsafe.Add(mBase, uint32(v2330))) = v2361 + v2342
	v2367 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2367 - v2342
	goto L391
L394:
	;
	v2375 = int32(0)
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v2375 < v2376 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2380 = v2375
	goto L398
L396:
	;
	goto L397
L397:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L415
	}
L398:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v2412 = *(*int32)(unsafe.Add(mBase, uint32(v2408+v2380<<(uint(int32(2))%32))))
	v2414 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[12]))
	if v2414 != 0 {
		goto L400
	} else {
		goto L401
	}
L399:
	;
	goto L397
L400:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L1
	} else {
		goto L403
	}
L401:
	;
	goto L402
L402:
	;
	v2417 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2418 = m.ExcPending
	if v2418 != 0 {
		goto L1
	} else {
		goto L404
	}
L403:
	;
	goto L402
L404:
	;
	F_PushActiveSnapshot(m, v2417)
	mBase = m.M
	v2420 = m.ExcPending
	if v2420 != 0 {
		goto L1
	} else {
		goto L405
	}
L405:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2412)))
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2412)+4))
	v2424 = F_table_open(m, v2422, int32(4))
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L1
	} else {
		goto L406
	}
L406:
	;
	v2427 = F_index_open(m, v2421, int32(4))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L407
	}
L407:
	;
	F_TransferPredicateLocksToHeapRelation(m, v2427)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L408
	}
L408:
	;
	F_index_set_state_flags(m, v2421, int32(3))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L1
	} else {
		goto L409
	}
L409:
	;
	F_CacheInvalidateRelcache(m, v2424)
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L1
	} else {
		goto L410
	}
L410:
	;
	F_relation_close(m, v2424, int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L411
	}
L411:
	;
	F_relation_close(m, v2427, int32(0))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L413
	}
L413:
	;
	v2445 = v2380 + int32(1)
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v2445 < v2446 {
		v2380 = v2445
		goto L398
	} else {
		goto L414
	}
L414:
	;
	goto L399
L415:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	v2485 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v2485 == int32(0) {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	F_WaitForLockersMultiple(m, v1219, int32(8), int32(1))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L421
	}
L418:
	;
	goto L417
L419:
	;
	v2489 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v2489&int32(1) == int32(0) {
		goto L418
	} else {
		goto L420
	}
L420:
	;
	v2494 = int32(_a_F_ReindexRelationConcurrently_14)
	v2496 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v2497 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2496 + v2497
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(v2485)))
	*(*int32)(unsafe.Add(mBase, uint32(v2485))) = v2500 + v2497
	v2504 = int32(0)
	v2506 = int32(_a_F_ReindexRelationConcurrently_15)
	v2507 = base.AtomicRmwOr32(m, v2504, v2506, v2504)
	*(*int64)(unsafe.Add(mBase, uint32(v2485+int32(72))+232)) = int64(9)
	v2515 = base.AtomicRmwOr32(m, v2504, v2506, v2504)
	v2516 = *(*int32)(unsafe.Add(mBase, uint32(v2485)))
	*(*int32)(unsafe.Add(mBase, uint32(v2485))) = v2516 + v2497
	v2522 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v2522 - v2497
	goto L418
L421:
	;
	v2530 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L1
	} else {
		goto L422
	}
L422:
	;
	F_PushActiveSnapshot(m, v2530)
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L1
	} else {
		goto L423
	}
L423:
	;
	v2534 = F_new_object_addresses(m)
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if int32(0) < v2536 {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v2541 = int32(0)
	goto L428
L426:
	;
	goto L427
L427:
	;
	v2617 = int32(0)
	F_performMultipleDeletions(m, v2534, v2617, int32(33))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L1
	} else {
		goto L432
	}
L428:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v2573 = *(*int32)(unsafe.Add(mBase, uint32(v2569+v2541<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+168)) = int32(1259)
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v2573)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+172)) = v2576
	F_add_exact_object_address(m, v32+int32(168), v2534)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L1
	} else {
		goto L430
	}
L429:
	;
	goto L427
L430:
	;
	v2585 = v2541 + int32(1)
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	if v2585 < v2586 {
		v2541 = v2585
		goto L428
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2623 = m.ExcPending
	if v2623 != 0 {
		goto L1
	} else {
		goto L433
	}
L433:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2625 = m.ExcPending
	if v2625 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	if v1210 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v2702 = m.ExcPending
	if v2702 != 0 {
		goto L1
	} else {
		goto L442
	}
L436:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	if v2628 <= int32(0) {
		goto L435
	} else {
		goto L437
	}
L437:
	;
	v2632 = v2617
	goto L438
L438:
	;
	v2660 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+12))
	v2664 = *(*int32)(unsafe.Add(mBase, uint32(v2660+v2632<<(uint(int32(2))%32))))
	F_UnlockRelationIdForSession(m, v2664, int32(4))
	mBase = m.M
	v2667 = m.ExcPending
	if v2667 != 0 {
		goto L1
	} else {
		goto L440
	}
L439:
	;
	goto L435
L440:
	;
	v2669 = v2632 + int32(1)
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v1210)+4))
	if v2669 < v2670 {
		v2632 = v2669
		goto L438
	} else {
		goto L441
	}
L441:
	;
	goto L439
L442:
	;
	v2703 = int32(1)
	v2704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v2704&v2703 == int32(0) {
		goto L375
	} else {
		goto L443
	}
L443:
	;
	if v74 == int32(105) {
		goto L382
	} else {
		goto L444
	}
L444:
	;
	if v1106 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v2810 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L1
	} else {
		goto L460
	}
L446:
	;
	v2713 = int32(0)
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v2714 <= v2713 {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v2718 = v2713
	goto L448
L448:
	;
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+12))
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2746+v2718<<(uint(int32(2))%32))))
	v2751 = *(*int32)(unsafe.Add(mBase, uint32(v2750)))
	v2754 = F_errstart(m, int32(17), int32(0))
	mBase = m.M
	v2755 = m.ExcPending
	if v2755 != 0 {
		goto L1
	} else {
		goto L450
	}
L449:
	;
	goto L445
L450:
	;
	if v2754 != 0 {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v2756 = F_get_rel_namespace(m, v2751)
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	goto L453
L453:
	;
	v2776 = v2718 + int32(1)
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v1106)+4))
	if v2776 < v2777 {
		v2718 = v2776
		goto L448
	} else {
		goto L459
	}
L454:
	;
	v2758 = F_get_namespace_name(m, v2756)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v2760 = F_get_rel_name(m, v2751)
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v2760
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v2758
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_17), v32+int32(48))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(_a_F_ReindexRelationConcurrently_18), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v2773 = m.ExcPending
	if v2773 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	goto L453
L459:
	;
	goto L449
L460:
	;
	if v2810 == int32(0) {
		goto L375
	} else {
		goto L461
	}
L461:
	;
	v3604 = int32(_a_F_ReindexRelationConcurrently_19)
	v3623 = int32(_a_F_ReindexRelationConcurrently_20)
	goto L376
L462:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v2824 = m.ExcPending
	if v2824 != 0 {
		goto L1
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2820)))
	v2826 = F_get_rel_name(m, v2825)
	mBase = m.M
	v2827 = m.ExcPending
	if v2827 != 0 {
		goto L1
	} else {
		goto L466
	}
L465:
	;
	goto L464
L466:
	;
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+4))
	v2831 = F_get_rel_namespace(m, v2830)
	mBase = m.M
	v2832 = m.ExcPending
	if v2832 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	v2834 = F_ChooseRelationName(m, v2826, int32(0), int32(_a_F_ReindexRelationConcurrently_21), v2831, int32(0))
	mBase = m.M
	v2835 = m.ExcPending
	if v2835 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v2836 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	F_PushActiveSnapshot(m, v2836)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2819)))
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2820)))
	v2842 = m.G0
	v2844 = v2842 - int32(288)
	m.G0 = v2844
	v2847 = F_relation_open(m, v2841, int32(4))
	mBase = m.M
	v2848 = m.ExcPending
	if v2848 != 0 {
		goto L1
	} else {
		goto L471
	}
L471:
	;
	v2850 = F_relation_open(m, v2840, int32(4))
	mBase = m.M
	v2851 = m.ExcPending
	if v2851 != 0 {
		goto L1
	} else {
		goto L472
	}
L472:
	;
	v2854 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v2855 = m.ExcPending
	if v2855 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	v2857 = base.I64_extend_i32_u(v2841)
	v2859 = F_SearchSysCacheCopy(m, int32(57), v2857, int64(0))
	mBase = m.M
	v2860 = m.ExcPending
	if v2860 != 0 {
		goto L1
	} else {
		goto L479
	}
L474:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v3578 = m.ExcPending
	if v3578 != 0 {
		goto L1
	} else {
		goto L635
	}
L475:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3565 = m.ExcPending
	if v3565 != 0 {
		goto L1
	} else {
		goto L632
	}
L476:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L1
	} else {
		goto L629
	}
L477:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3535 = m.ExcPending
	if v3535 != 0 {
		goto L1
	} else {
		goto L626
	}
L478:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L1
	} else {
		goto L623
	}
L479:
	;
	if v2859 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	v2862 = base.I64_extend_i32_u(v2840)
	v2864 = F_SearchSysCacheCopy(m, int32(57), v2862, int64(0))
	mBase = m.M
	v2865 = m.ExcPending
	if v2865 != 0 {
		goto L1
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L1
	} else {
		goto L620
	}
L483:
	;
	if v2864 == int32(0) {
		goto L478
	} else {
		goto L484
	}
L484:
	;
	v2868 = *(*int32)(unsafe.Add(mBase, uint32(v2864)+16))
	v2869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2868)+22)))
	v2870 = v2868 + v2869
	v2871 = int32(4)
	v2873 = *(*int32)(unsafe.Add(mBase, uint32(v2859)+16))
	v2874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2873)+22)))
	v2875 = v2873 + v2874
	v2877 = v2875 + v2871
	v2879 = F_strncpy(m, v2870+v2871, v2877, int32(64))
	mBase = m.M
	v2880 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2879)+63)) = uint8(v2880)
	goto L485
L485:
	;
	v2883 = F_strncpy(m, v2877, v2834, int32(64))
	mBase = m.M
	v2884 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2883)+63)) = uint8(v2884)
	goto L486
L486:
	;
	v2886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2870)+131)))
	v2887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2875)+131)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2870)+131)) = uint8(v2887)
	*(*uint8)(unsafe.Add(mBase, uint32(v2875)+131)) = uint8(v2886)
	F_CatalogTupleUpdate(m, v2854, v2859+int32(4), v2859)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L487
	}
L487:
	;
	F_CatalogTupleUpdate(m, v2854, v2864+int32(4), v2864)
	mBase = m.M
	v2897 = m.ExcPending
	if v2897 != 0 {
		goto L1
	} else {
		goto L488
	}
L488:
	;
	F_pfree(m, v2859)
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	F_pfree(m, v2864)
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	v2904 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L491
	}
L491:
	;
	v2908 = F_SearchSysCacheCopy(m, int32(34), v2857, int64(0))
	mBase = m.M
	v2909 = m.ExcPending
	if v2909 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	if v2908 == int32(0) {
		goto L477
	} else {
		goto L493
	}
L493:
	;
	v2914 = F_SearchSysCacheCopy(m, int32(34), v2862, int64(0))
	mBase = m.M
	v2915 = m.ExcPending
	if v2915 != 0 {
		goto L1
	} else {
		goto L494
	}
L494:
	;
	if v2914 == int32(0) {
		goto L476
	} else {
		goto L495
	}
L495:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, uint32(v2914)+16))
	v2919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2918)+22)))
	v2920 = v2918 + v2919
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2908)+16))
	v2922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2921)+22)))
	v2923 = v2921 + v2922
	v2924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2923)+14)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+14)) = uint8(v2924)
	v2926 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2923)+14)) = uint8(v2926)
	v2928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2923)+15)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+15)) = uint8(v2928)
	*(*uint8)(unsafe.Add(mBase, uint32(v2923)+15)) = uint8(v2926)
	v2932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2923)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+16)) = uint8(v2932)
	v2934 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2923)+16)) = uint8(v2934)
	v2936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2923)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+22)) = uint8(v2936)
	v2938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2923)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+18)) = uint8(v2934)
	*(*uint8)(unsafe.Add(mBase, uint32(v2920)+17)) = uint8(v2938)
	*(*uint8)(unsafe.Add(mBase, uint32(v2923)+22)) = uint8(v2926)
	*(*uint16)(unsafe.Add(mBase, uint32(v2923)+17)) = uint16(v2926)
	F_CatalogTupleUpdate(m, v2904, v2908+int32(4), v2908)
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L1
	} else {
		goto L496
	}
L496:
	;
	F_CatalogTupleUpdate(m, v2904, v2914+int32(4), v2914)
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	F_pfree(m, v2908)
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	F_pfree(m, v2914)
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	v2958 = int32(0)
	v2959 = m.G0
	v2961 = v2959 - int32(176)
	m.G0 = v2961
	v2965 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v2966 = m.ExcPending
	if v2966 != 0 {
		goto L1
	} else {
		goto L500
	}
L500:
	;
	F_ScanKeyInit(m, v2961, int32(4), int32(3), int32(184), int64(1259))
	mBase = m.M
	v2972 = m.ExcPending
	if v2972 != 0 {
		goto L1
	} else {
		goto L501
	}
L501:
	;
	F_ScanKeyInit(m, v2961+int32(56), int32(5), int32(3), int32(184), base.I64_extend_i32_u(v2841))
	mBase = m.M
	v2980 = m.ExcPending
	if v2980 != 0 {
		goto L1
	} else {
		goto L502
	}
L502:
	;
	F_ScanKeyInit(m, v2961+int32(112), int32(6), int32(3), int32(65), int64(0))
	mBase = m.M
	v2988 = m.ExcPending
	if v2988 != 0 {
		goto L1
	} else {
		goto L503
	}
L503:
	;
	v2993 = F_systable_beginscan(m, v2965, int32(2674), int32(1), int32(0), int32(3), v2961)
	mBase = m.M
	v2994 = m.ExcPending
	if v2994 != 0 {
		goto L1
	} else {
		goto L504
	}
L504:
	;
	v2995 = F_systable_getnext(m, v2993)
	mBase = m.M
	v2996 = m.ExcPending
	if v2996 != 0 {
		goto L1
	} else {
		goto L505
	}
L505:
	;
	if v2995 != 0 {
		goto L506
	} else {
		goto L507
	}
L506:
	;
	v2997 = v2995
	v3000 = v2958
	goto L509
L507:
	;
	v3045 = v2958
	goto L508
L508:
	;
	F_systable_endscan(m, v2993)
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L1
	} else {
		goto L518
	}
L509:
	;
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v2997)+16))
	v3027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3026)+22)))
	v3028 = v3026 + v3027
	v3029 = *(*int32)(unsafe.Add(mBase, uint32(v3028)))
	if v3029 != int32(2606) {
		v3039 = v3000
		goto L511
	} else {
		goto L512
	}
L510:
	;
	v3045 = v3039
	goto L508
L511:
	;
	v3040 = F_systable_getnext(m, v2993)
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L1
	} else {
		goto L516
	}
L512:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+8))
	if v3032 != 0 {
		v3039 = v3000
		goto L511
	} else {
		goto L513
	}
L513:
	;
	v3033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3028)+24)))
	if v3033 != int32(110) {
		v3039 = v3000
		goto L511
	} else {
		goto L514
	}
L514:
	;
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v3028)+4))
	v3037 = F_lappend_oid(m, v3000, v3036)
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v3039 = v3037
	goto L511
L516:
	;
	if v3040 != 0 {
		v2997 = v3040
		v3000 = v3039
		goto L509
	} else {
		goto L517
	}
L517:
	;
	goto L510
L518:
	;
	F_relation_close(m, v2965, int32(1))
	mBase = m.M
	v3075 = m.ExcPending
	if v3075 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	m.G0 = v2961 + int32(176)
	v3079 = F_get_index_constraint(m, v2841)
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	if v3079 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	v3081 = F_lappend_oid(m, v3045, v3079)
	mBase = m.M
	v3082 = m.ExcPending
	if v3082 != 0 {
		goto L1
	} else {
		goto L524
	}
L522:
	;
	v3083 = v3045
	goto L523
L523:
	;
	v3086 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v3087 = m.ExcPending
	if v3087 != 0 {
		goto L1
	} else {
		goto L525
	}
L524:
	;
	v3083 = v3081
	goto L523
L525:
	;
	v3090 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	if v3083 == int32(0) {
		goto L527
	} else {
		goto L528
	}
L527:
	;
	v3247 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2844)+104)) = v3247
	*(*int64)(unsafe.Add(mBase, uint32(v2844)+96)) = v3247
	*(*int64)(unsafe.Add(mBase, uint32(v2844)+88)) = v3247
	*(*int64)(unsafe.Add(mBase, uint32(v2844)+80)) = v3247
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+76)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2844)+80)) = v2862
	v3258 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+72)) = v3258
	v3261 = v2844 + int32(112)
	F_ScanKeyInit(m, v3261, v3258, int32(3), int32(184), v2857)
	mBase = m.M
	v3266 = m.ExcPending
	if v3266 != 0 {
		goto L1
	} else {
		goto L553
	}
L528:
	;
	v3094 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+4))
	if v3094 <= int32(0) {
		goto L527
	} else {
		goto L529
	}
L529:
	;
	v3098 = int32(0)
	goto L530
L530:
	;
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+12))
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3128+v3098<<(uint(int32(2))%32))))
	v3133 = base.I64_extend_i32_u(v3132)
	v3135 = F_SearchSysCacheCopy(m, int32(19), v3133, int64(0))
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L532
	}
L531:
	;
	goto L527
L532:
	;
	if v3135 == int32(0) {
		goto L475
	} else {
		goto L533
	}
L533:
	;
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(v3135)+16))
	v3140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3139)+22)))
	v3141 = v3139 + v3140
	v3142 = *(*int32)(unsafe.Add(mBase, uint32(v3141)+88))
	if v2841 == v3142 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3141)+88)) = v2840
	F_CatalogTupleUpdate(m, v3086, v3135+int32(4), v3135)
	mBase = m.M
	v3148 = m.ExcPending
	if v3148 != 0 {
		goto L1
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	F_pfree(m, v3135)
	mBase = m.M
	v3150 = m.ExcPending
	if v3150 != 0 {
		goto L1
	} else {
		goto L538
	}
L537:
	;
	goto L536
L538:
	;
	v3152 = v2844 + int32(112)
	F_ScanKeyInit(m, v3152, int32(11), int32(3), int32(184), v3133)
	mBase = m.M
	v3157 = m.ExcPending
	if v3157 != 0 {
		goto L1
	} else {
		goto L539
	}
L539:
	;
	v3159 = int32(1)
	v3162 = F_systable_beginscan(m, v3090, int32(2699), v3159, int32(0), v3159, v3152)
	mBase = m.M
	v3163 = m.ExcPending
	if v3163 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	goto L541
L541:
	;
	v3193 = F_systable_getnext(m, v3162)
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L1
	} else {
		goto L543
	}
L542:
	;
	F_systable_endscan(m, v3162)
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L1
	} else {
		goto L551
	}
L543:
	;
	if v3193 != 0 {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(v3193)+16))
	v3196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3195)+22)))
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3195+v3196)+88))
	if v3198 != v2841 {
		goto L541
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	goto L542
L547:
	;
	v3200 = F_heap_copytuple(m, v3193)
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(v3200)+16))
	v3203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3202)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3202+v3203)+88)) = v2840
	F_CatalogTupleUpdate(m, v3090, v3200+int32(4), v3200)
	mBase = m.M
	v3209 = m.ExcPending
	if v3209 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	F_pfree(m, v3200)
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	goto L541
L551:
	;
	v3215 = v3098 + int32(1)
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(v3083)+4))
	if v3215 < v3216 {
		v3098 = v3215
		goto L530
	} else {
		goto L552
	}
L552:
	;
	goto L531
L553:
	;
	F_ScanKeyInit(m, v2844+int32(168), int32(2), int32(3), int32(184), int64(1259))
	mBase = m.M
	v3274 = m.ExcPending
	if v3274 != 0 {
		goto L1
	} else {
		goto L554
	}
L554:
	;
	v3277 = int32(3)
	F_ScanKeyInit(m, v2844+int32(224), v3277, v3277, int32(65), int64(0))
	mBase = m.M
	v3282 = m.ExcPending
	if v3282 != 0 {
		goto L1
	} else {
		goto L555
	}
L555:
	;
	v3285 = F_table_open(m, int32(2609), int32(3))
	mBase = m.M
	v3286 = m.ExcPending
	if v3286 != 0 {
		goto L1
	} else {
		goto L556
	}
L556:
	;
	v3291 = F_systable_beginscan(m, v3285, int32(2675), int32(1), int32(0), int32(3), v3261)
	mBase = m.M
	v3292 = m.ExcPending
	if v3292 != 0 {
		goto L1
	} else {
		goto L557
	}
L557:
	;
	v3293 = F_systable_getnext(m, v3291)
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L1
	} else {
		goto L558
	}
L558:
	;
	if v3293 != 0 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v3285)+52))
	v3302 = F_heap_modify_tuple(m, v3293, v3295, v2844+int32(80), v2844+int32(76), v2844+int32(72))
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L1
	} else {
		goto L562
	}
L560:
	;
	goto L561
L561:
	;
	F_systable_endscan(m, v3291)
	mBase = m.M
	v3310 = m.ExcPending
	if v3310 != 0 {
		goto L1
	} else {
		goto L564
	}
L562:
	;
	F_CatalogTupleUpdate(m, v3285, v3302+int32(4), v3302)
	mBase = m.M
	v3307 = m.ExcPending
	if v3307 != 0 {
		goto L1
	} else {
		goto L563
	}
L563:
	;
	goto L561
L564:
	;
	F_relation_close(m, v3285, int32(0))
	mBase = m.M
	v3313 = m.ExcPending
	if v3313 != 0 {
		goto L1
	} else {
		goto L565
	}
L565:
	;
	v3314 = F_get_rel_relispartition(m, v2841)
	mBase = m.M
	v3315 = m.ExcPending
	if v3315 != 0 {
		goto L1
	} else {
		goto L566
	}
L566:
	;
	if v3314 != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3316 = F_get_partition_ancestors(m, v2841)
	mBase = m.M
	v3317 = m.ExcPending
	if v3317 != 0 {
		goto L1
	} else {
		goto L570
	}
L568:
	;
	goto L569
L569:
	;
	F_changeDependenciesOf(m, v2840, v2841)
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L1
	} else {
		goto L574
	}
L570:
	;
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v3316)+12))
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3318)))
	v3320 = int32(0)
	v3322 = F_DeleteInheritsTuple(m, v2841, v3319, v3320, v3320)
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	F_StoreSingleInheritance(m, v2840, v3319, int32(1))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	F_list_free(m, v3316)
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	goto L569
L574:
	;
	F_changeDependenciesOn(m, v2840, v2841)
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L1
	} else {
		goto L575
	}
L575:
	;
	F_changeDependenciesOf(m, v2841, v2840)
	mBase = m.M
	v3336 = m.ExcPending
	if v3336 != 0 {
		goto L1
	} else {
		goto L576
	}
L576:
	;
	F_changeDependenciesOn(m, v2841, v2840)
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L1
	} else {
		goto L577
	}
L577:
	;
	v3342 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3343 = *(*int32)(unsafe.Add(mBase, uint32(v2847)+48))
	v3344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3343)+117)))
	if v3344 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v3345 = int32(0)
	goto L580
L579:
	;
	v3345 = v3342
	goto L580
L580:
	;
	v3346 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2847)+56)))
	v3348 = F_pgstat_fetch_entry(m, int32(2), v3345, v3346, int32(0))
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	if v3348 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v3353 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[16]))
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v2850)+48))
	v3355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3354)+117)))
	if v3355 != 0 {
		goto L585
	} else {
		goto L586
	}
L583:
	;
	goto L584
L584:
	;
	v3370 = m.G0
	v3372 = v3370 + int32(-64)
	m.G0 = v3372
	v3376 = F_table_open(m, int32(2619), int32(3))
	mBase = m.M
	v3377 = m.ExcPending
	if v3377 != 0 {
		goto L1
	} else {
		goto L590
	}
L585:
	;
	v3356 = int32(0)
	goto L587
L586:
	;
	v3356 = v3353
	goto L587
L587:
	;
	v3357 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2850)+56)))
	v3359 = F_pgstat_get_entry_ref_locked(m, int32(2), v3356, v3357, int32(0))
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L1
	} else {
		goto L588
	}
L588:
	;
	v3361 = *(*int32)(unsafe.Add(mBase, uint32(v3359)+4))
	base.MemoryCopy(m, v3361+int32(24), v3348, int32(224))
	F_pgstat_unlock_entry(m, v3359)
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L1
	} else {
		goto L589
	}
L589:
	;
	goto L584
L590:
	;
	F_ScanKeyInit(m, v3372, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v2841))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L1
	} else {
		goto L591
	}
L591:
	;
	v3385 = int32(1)
	v3388 = F_systable_beginscan(m, v3376, int32(2696), v3385, int32(0), v3385, v3372)
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
		goto L1
	} else {
		goto L593
	}
L592:
	;
	F_relation_close(m, v3376, int32(3))
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L1
	} else {
		goto L613
	}
L593:
	;
	v3390 = F_systable_getnext(m, v3388)
	mBase = m.M
	v3391 = m.ExcPending
	if v3391 != 0 {
		goto L1
	} else {
		goto L594
	}
L594:
	;
	if v3390 == int32(0) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	F_systable_endscan(m, v3388)
	mBase = m.M
	v3395 = m.ExcPending
	if v3395 != 0 {
		goto L1
	} else {
		goto L598
	}
L596:
	;
	goto L597
L597:
	;
	v3399 = int32(0)
	v3409 = v3390
	goto L599
L598:
	;
	goto L592
L599:
	;
	v3425 = F_heap_copytuple(m, v3409)
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L601
	}
L600:
	;
	F_systable_endscan(m, v3388)
	mBase = m.M
	v3443 = m.ExcPending
	if v3443 != 0 {
		goto L1
	} else {
		goto L610
	}
L601:
	;
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(v3425)+16))
	v3428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3427)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(v3427+v3428))) = v2840
	if v3399 == int32(0) {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v3433 = F_CatalogOpenIndexes(m, v3376)
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L605
	}
L603:
	;
	v3435 = v3399
	goto L604
L604:
	;
	F_CatalogTupleInsertWithInfo(m, v3376, v3425, v3435)
	mBase = m.M
	v3437 = m.ExcPending
	if v3437 != 0 {
		goto L1
	} else {
		goto L606
	}
L605:
	;
	v3435 = v3433
	goto L604
L606:
	;
	F_pfree(m, v3425)
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L607
	}
L607:
	;
	v3440 = F_systable_getnext(m, v3388)
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L1
	} else {
		goto L608
	}
L608:
	;
	if v3440 != 0 {
		v3399 = v3435
		v3409 = v3440
		goto L599
	} else {
		goto L609
	}
L609:
	;
	goto L600
L610:
	;
	if v3435 == int32(0) {
		goto L592
	} else {
		goto L611
	}
L611:
	;
	F_CatalogCloseIndexes(m, v3435)
	mBase = m.M
	v3447 = m.ExcPending
	if v3447 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	goto L592
L613:
	;
	m.G0 = v3372 - int32(-64)
	F_relation_close(m, v2854, int32(3))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	F_relation_close(m, v2904, int32(3))
	mBase = m.M
	v3488 = m.ExcPending
	if v3488 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	F_relation_close(m, v3086, int32(3))
	mBase = m.M
	v3491 = m.ExcPending
	if v3491 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	F_relation_close(m, v3090, int32(3))
	mBase = m.M
	v3494 = m.ExcPending
	if v3494 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	F_relation_close(m, v2847, int32(0))
	mBase = m.M
	v3497 = m.ExcPending
	if v3497 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	F_relation_close(m, v2850, int32(0))
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	m.G0 = v2844 + int32(288)
	goto L474
L620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844))) = v2841
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_22), v2844)
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_23), int32(1618), int32(_a_F_ReindexRelationConcurrently_24))
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L623:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+16)) = v2840
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_22), v2844+int32(16))
	mBase = m.M
	v3526 = m.ExcPending
	if v3526 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_23), int32(1622), int32(_a_F_ReindexRelationConcurrently_24))
	mBase = m.M
	v3531 = m.ExcPending
	if v3531 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+32)) = v2841
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_22), v2844+int32(32))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_23), int32(1648), int32(_a_F_ReindexRelationConcurrently_24))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+48)) = v2840
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_22), v2844+int32(48))
	mBase = m.M
	v3556 = m.ExcPending
	if v3556 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_23), int32(1652), int32(_a_F_ReindexRelationConcurrently_24))
	mBase = m.M
	v3561 = m.ExcPending
	if v3561 != 0 {
		goto L1
	} else {
		goto L631
	}
L631:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+64)) = v3132
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_25), v2844-int32(-64))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_23), int32(1716), int32(_a_F_ReindexRelationConcurrently_24))
	mBase = m.M
	v3576 = m.ExcPending
	if v3576 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	v3579 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+4))
	F_CacheInvalidateRelcacheByRelid(m, v3579)
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v3585 = m.ExcPending
	if v3585 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	v2277 = v2277 + int32(1)
	goto L377
L638:
	;
	if v3588 == int32(0) {
		goto L375
	} else {
		goto L639
	}
L639:
	;
	v3604 = int32(_a_F_ReindexRelationConcurrently_17)
	v3623 = int32(_a_F_ReindexRelationConcurrently_26)
	goto L376
L640:
	;
	v3632 = F_pg_rusage_show(m, v32+int32(248))
	mBase = m.M
	v3633 = m.ExcPending
	if v3633 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v3632
	v3638 = F_errdetail(m, int32(_a_F_ReindexRelationConcurrently_27), v32+int32(16))
	mBase = m.M
	v3639 = m.ExcPending
	if v3639 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), v3623, int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v3643 = m.ExcPending
	if v3643 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	goto L375
L644:
	;
	v3677 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[7]))
	if v3677 == int32(0) {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v3724 = v2703
	goto L12
L646:
	;
	goto L645
L647:
	;
	v3681 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[8])))
	if v3681&int32(1) == int32(0) {
		goto L646
	} else {
		goto L648
	}
L648:
	;
	v3686 = *(*int32)(unsafe.Add(mBase, uint32(v3677)+220))
	if v3686 == int32(0) {
		goto L646
	} else {
		goto L649
	}
L649:
	;
	v3689 = int32(_a_F_ReindexRelationConcurrently_14)
	v3691 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	v3692 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v3691 + v3692
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v3677)))
	*(*int32)(unsafe.Add(mBase, uint32(v3677))) = v3695 + v3692
	v3699 = int32(0)
	v3701 = int32(_a_F_ReindexRelationConcurrently_15)
	v3702 = base.AtomicRmwOr32(m, v3699, v3701, v3699)
	*(*int32)(unsafe.Add(mBase, uint32(v3677)+220)) = v3699
	*(*int32)(unsafe.Add(mBase, uint32(v3677)+224)) = v3699
	v3710 = base.AtomicRmwOr32(m, v3699, v3701, v3699)
	v3711 = *(*int32)(unsafe.Add(mBase, uint32(v3677)))
	*(*int32)(unsafe.Add(mBase, uint32(v3677))) = v3711 + v3692
	v3717 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReindexRelationConcurrently[9])) = v3717 - v3692
	goto L646
L650:
	;
	F_errmsg_internal(m, int32(_a_F_ReindexRelationConcurrently_28), int32(0))
	mBase = m.M
	v3761 = m.ExcPending
	if v3761 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3983), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v3766 = m.ExcPending
	if v3766 != 0 {
		goto L1
	} else {
		goto L652
	}
L652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L653:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3773 = m.ExcPending
	if v3773 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	v3774 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v3775 = F_get_tablespace_name(m, v3774)
	mBase = m.M
	v3776 = m.ExcPending
	if v3776 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v3775
	F_errmsg(m, int32(_a_F_ReindexRelationConcurrently_29), v32)
	mBase = m.M
	v3780 = m.ExcPending
	if v3780 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	F_errfinish(m, int32(_a_F_ReindexRelationConcurrently_6), int32(3908), int32(_a_F_ReindexRelationConcurrently_7))
	mBase = m.M
	v3785 = m.ExcPending
	if v3785 != 0 {
		goto L1
	} else {
		goto L657
	}
L657:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ResetReindexState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[0]))
	if l0 <= v3 {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[1])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[2])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[3])) = v6
		*(*int32)(unsafe.Add(mBase, _c_F_ResetReindexState[0])) = v6
	} else {
	}
	return
}
