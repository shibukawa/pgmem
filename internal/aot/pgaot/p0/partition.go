package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecFindPartition(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
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
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int64
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v617 int32
	_ = v617
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v708 int32
	_ = v708
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v904 int32
	_ = v904
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1091 int32
	_ = v1091
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1142 int32
	_ = v1142
	var v1153 int32
	_ = v1153
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
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
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1353 int32
	_ = v1353
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1435 int32
	_ = v1435
	var v1444 int32
	_ = v1444
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1491 int32
	_ = v1491
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1538 int32
	_ = v1538
	var v1546 int32
	_ = v1546
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1583 int32
	_ = v1583
	var v1594 int32
	_ = v1594
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1679 int32
	_ = v1679
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1719 int32
	_ = v1719
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1760 int32
	_ = v1760
	var v1777 int32
	_ = v1777
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1812 int32
	_ = v1812
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1874 int32
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1899 int32
	_ = v1899
	var v1918 int32
	_ = v1918
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1985 int32
	_ = v1985
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2019 int32
	_ = v2019
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2149 int32
	_ = v2149
	var v2155 int32
	_ = v2155
	var v2160 int32
	_ = v2160
	var v2165 int32
	_ = v2165
	var v2172 int32
	_ = v2172
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2196 int32
	_ = v2196
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
	var v2215 int32
	_ = v2215
	v6 = int32(0)
	v35 = m.G0
	v37 = v35 - int32(256)
	m.G0 = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v40 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = int32(_a_F_ExecFindPartition_0)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v60
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+131)))
	if v64 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v54 = v40
	v55 = v40 + int32(4)
	v56 = v43
	goto L1
L3:
	;
	goto L4
L4:
	;
	v44 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v49 = v44 + int32(4)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v51 != 0 {
		v54 = v51
		v55 = v49
		v56 = v50
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v52 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v54 = v52
	v55 = v49
	v56 = v50
	goto L1
L9:
	;
	v68 = F_ExecPartitionCheck(m, l1, l3, l4, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v70 == int32(0) {
		v2215 = v6
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v56
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v58
	m.G0 = v37 + int32(256)
	return v2215
L14:
	;
	v83 = v70
	v88 = l3
	v95 = v6
	goto L15
L15:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[1]))
	if v108 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v2172 == int32(0) {
		v2215 = v2155
		goto L13
	} else {
		goto L394
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v88
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v116 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v126)+4)))
	if int32(0) < v127 {
		goto L41
	} else {
		goto L42
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v125 = v123
	goto L21
L23:
	;
	if v114 != 0 {
		v122 = v114
		goto L22
	} else {
		goto L26
	}
L24:
	;
	v120 = v114
	goto L25
L25:
	;
	if v120 != 0 {
		v122 = v120
		goto L22
	} else {
		goto L28
	}
L26:
	;
	v117 = F_ExecPrepareExprList(m, v116, l4)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v117
	v120 = v117
	goto L25
L28:
	;
	v125 = int32(0)
	goto L21
L29:
	;
	v1182 = v1153 << (uint(int32(2)) % 32)
	v1185 = v83 + v1182 + int32(24)
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1185)))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1187+v1153))))
	if v1189 == int32(1) {
		goto L208
	} else {
		goto L209
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v540
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+20)) = v539
	v1153 = v540
	goto L29
L31:
	;
	v681 = F_RelationGetPartitionKey(m, v112)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L5
	} else {
		goto L113
	}
L32:
	;
	if int32(0) <= v617 {
		v1153 = v617
		goto L29
	} else {
		goto L112
	}
L33:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	if v539 != v605 {
		goto L30
	} else {
		goto L111
	}
L34:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v275)+32))
	v617 = v604
	goto L32
L35:
	;
	if int32(0) <= v540 {
		goto L33
	} else {
		goto L110
	}
L36:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if v468 < int32(16) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	v395 = v326
	goto L93
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L90
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L87
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L84
	}
L41:
	;
	v136 = v126
	v137 = int32(0)
	v140 = v125
	goto L44
L42:
	;
	v242 = v126
	v246 = v125
	v249 = v127
	goto L43
L43:
	;
	if v246 != 0 {
		goto L39
	} else {
		goto L64
	}
L44:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165+v137<<(uint(int32(1))%32)))))
	if v169 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	v242 = v234
	v246 = v219
	v249 = v235
	goto L43
L46:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v137))) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v37+int32(96)+v137<<(uint(int32(2))%32)))) = v220
	v233 = v137 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v234)+4)))
	if v233 < v235 {
		v136 = v234
		v137 = v233
		v140 = v219
		goto L44
	} else {
		goto L63
	}
L47:
	;
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+6)))
	if v170 < v169 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	if v140 == int32(0) {
		goto L40
	} else {
		goto L54
	}
L50:
	;
	F_slot_getsomeattrs_int(m, v88, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v175 = v169 - int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v88)+20))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+v175<<(uint(int32(2))%32))))
	v217 = v175 + v176
	v219 = v140
	v220 = v182
	goto L46
L53:
	;
	goto L52
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l4)+152))
	if v186 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v189 = F_MakePerTupleExprContext(m, l4)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	v191 = v186
	goto L57
L57:
	;
	v192 = int32(_a_F_ExecFindPartition_0)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v195
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v185)+20))
	v200 = m.T0[v199].(func(*base.Module, int32, int32, int32) int32)(m, v185, v191, v37+int32(240))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L59
	}
L58:
	;
	v191 = v189
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v193
	v205 = v140 + int32(4)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if base.Ui32(v205) < base.Ui32(v208+v209<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v214 = v205
	goto L62
L61:
	;
	v214 = int32(0)
	goto L62
L62:
	;
	v217 = v37 + int32(240)
	v219 = v214
	v220 = v200
	goto L46
L63:
	;
	goto L45
L64:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v271 == int32(0) {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	switch v276 - int32(104) {
	case 0:
		goto L66
	default:
		goto L38
	case 4:
		goto L68
	case 10:
		goto L67
	}
L66:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v337 = F_compute_partition_hash_value(m, v249, v331, v332, v37+int32(96), v37-int32(-64))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L5
	} else {
		goto L83
	}
L67:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+240)) = uint8(v326)
	if v326 < v249 {
		goto L37
	} else {
		goto L82
	}
L68:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+64)))
	if v279 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v275)+28))
	if v282 == int32(-1) {
		goto L34
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	if int32(16) <= v286 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v617 = v282
	goto L32
L73:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323+v294)))
	v617 = v325
	goto L32
L74:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v294 = v292 << (uint(int32(2)) % 32)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v294+v295)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v299 = F_FunctionCall2Coll(m, v289, v291, v298, v285)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L5
	} else {
		goto L77
	}
L75:
	;
	v305 = v285
	goto L76
L76:
	;
	v306 = int32(-1)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v311 = F_partition_list_bsearch(m, v307, v308, v275, v305, v37+int32(240))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L79
	}
L77:
	;
	if v299 == int32(0) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v305 = v303
	goto L76
L79:
	;
	if v311 < int32(0) {
		v539 = v311
		v540 = v306
		goto L35
	} else {
		goto L80
	}
L80:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+240)))
	if v315 != int32(1) {
		v539 = v311
		v540 = v306
		goto L35
	} else {
		goto L81
	}
L81:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+v311<<(uint(int32(2))%32))))
	v539 = v311
	v540 = v322
	goto L35
L82:
	;
	goto L36
L83:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v340 = int64(*(*int32)(unsafe.Add(mBase, uint32(v275)+20)))
	v341 = base.I64_rem_u_s(v337, v340)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v339+base.I32_wrap_i64(v341)<<(uint(int32(2))%32))))
	v617 = v346
	goto L32
L84:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1337), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_1), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1348), int32(_a_F_ExecFindPartition_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v377
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_4), v37+int32(48))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(1571), int32(_a_F_ExecFindPartition_5))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v395))))
	if v426 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v432 = int32(-1)
	v539 = v432
	v540 = v432
	goto L35
L95:
	;
	v430 = v395 + int32(1)
	if v249 != v430 {
		v395 = v430
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L36
L99:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v522 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v527 = F_partition_range_datum_bsearch(m, v520, v521, v275, v522, v37+int32(96), v37+int32(240))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L5
	} else {
		goto L109
	}
L100:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v274)+20))
	v475 = v473 << (uint(int32(2)) % 32)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v475+v476)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v479+v475)))
	v484 = F_partition_rbound_datum_cmp(m, v471, v472, v478, v481, v37+int32(96), v249)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v484 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488+v475)+4))
	v617 = v490
	goto L32
L103:
	;
	goto L104
L104:
	;
	if int32(0) <= v484 {
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v494 = v473 + int32(1)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v495 <= v494 {
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v242)+24))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v242)+28))
	v500 = v494 << (uint(int32(2)) % 32)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v500+v501)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v504+v500)))
	v509 = int32(*(*int16)(unsafe.Add(mBase, uint32(v242)+4)))
	v510 = F_partition_rbound_datum_cmp(m, v497, v498, v503, v506, v37+int32(96), v509)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	if v510 <= int32(0) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v514+v500)))
	v617 = v516
	goto L32
L109:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v275)+24))
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v529+v527<<(uint(int32(2))%32))+4))
	v539 = v527
	v540 = v533
	goto L35
L110:
	;
	goto L34
L111:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v274)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = v607 + int32(1)
	v1153 = v540
	goto L29
L112:
	;
	goto L31
L113:
	;
	v683 = int32(*(*int16)(unsafe.Add(mBase, uint32(v681)+4)))
	v684 = int32(0)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v112)+56))
	v688 = F_check_enable_rls(m, v685, v684, int32(1))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L5
	} else {
		goto L115
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L5
	} else {
		goto L198
	}
L115:
	;
	if v688 == int32(2) {
		v1091 = v684
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v693 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v695 = F_pg_class_aclcheck(m, v685, v693, int64(2))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L5
	} else {
		goto L118
	}
L117:
	;
	F_initStringInfo(m, v37+int32(240))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L5
	} else {
		goto L127
	}
L118:
	;
	if v695 == int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	if v683 <= int32(0) {
		goto L117
	} else {
		goto L120
	}
L120:
	;
	v708 = int32(0)
	goto L121
L121:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v681)+8))
	v740 = int32(*(*int16)(unsafe.Add(mBase, uint32(v736+v708<<(uint(int32(1))%32)))))
	if v740 == int32(0) {
		v1091 = v684
		goto L114
	} else {
		goto L123
	}
L122:
	;
	goto L117
L123:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[2]))
	v746 = F_pg_attribute_aclcheck(m, v685, v740, v744, int64(2))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	if v746 != 0 {
		v1091 = v684
		goto L114
	} else {
		goto L125
	}
L125:
	;
	v749 = v708 + int32(1)
	if v749 != v683 {
		v708 = v749
		goto L121
	} else {
		goto L126
	}
L126:
	;
	goto L122
L127:
	;
	v792 = F_pg_get_partkeydef_worker(m, v685, int32(7), int32(1), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L5
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v792
	F_appendStringInfo(m, v37+int32(240), int32(_a_F_ExecFindPartition_6), v37+int32(32))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	if v683 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	F_appendStringInfoChar(m, v37+int32(240), int32(41))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L5
	} else {
		goto L197
	}
L131:
	;
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+64)))
	if v805 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v681)+32))
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v808)))
	F_getTypeOutputInfo(m, v809, v37+int32(236), v37+int32(235))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	v820 = int32(_a_F_ExecFindPartition_7)
	goto L134
L134:
	;
	if v820&int32(3) == int32(0) {
		v844 = v820
		goto L140
	} else {
		goto L141
	}
L135:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v37)+96))
	v818 = F_OidOutputFunctionCall(m, v816, v817)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L5
	} else {
		goto L136
	}
L136:
	;
	v820 = v818
	goto L134
L137:
	;
	v896 = int32(1)
	if v683 == v896 {
		goto L130
	} else {
		goto L162
	}
L138:
	;
	if int32(65) <= v877 {
		goto L155
	} else {
		goto L156
	}
L139:
	;
	v877 = v869 - v820
	goto L138
L140:
	;
	v848 = v844
	goto L149
L141:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820))))
	if v828 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v877 = int32(0)
	goto L138
L143:
	;
	goto L144
L144:
	;
	v833 = v820
	goto L145
L145:
	;
	v837 = v833 + int32(1)
	if v837&int32(3) == int32(0) {
		v844 = v837
		goto L140
	} else {
		goto L147
	}
L146:
	;
	v869 = v837
	goto L139
L147:
	;
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
	if v842 != 0 {
		v833 = v837
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	v857 = int32(-2139062144)
	if (int32(16843008)-v854|v854)&v857 == v857 {
		v848 = v848 + int32(4)
		goto L149
	} else {
		goto L151
	}
L150:
	;
	v863 = v848
	goto L152
L151:
	;
	goto L150
L152:
	;
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v863))))
	if v867 != 0 {
		v863 = v863 + int32(1)
		goto L152
	} else {
		goto L154
	}
L153:
	;
	v869 = v863
	goto L139
L154:
	;
	goto L153
L155:
	;
	v883 = F_pg_mbcliplen(m, v820, v877, int32(64))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L5
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v820, v877)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L5
	} else {
		goto L161
	}
L158:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v820, v883)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L5
	} else {
		goto L160
	}
L160:
	;
	goto L137
L161:
	;
	goto L137
L162:
	;
	v904 = v896
	goto L163
L163:
	;
	v937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37-int32(-64)+v904))))
	if v937 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	goto L130
L165:
	;
	v941 = v904 << (uint(int32(2)) % 32)
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v681)+32))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v941+v942)))
	F_getTypeOutputInfo(m, v944, v37+int32(236), v37+int32(235))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L5
	} else {
		goto L168
	}
L166:
	;
	v958 = int32(_a_F_ExecFindPartition_7)
	goto L167
L167:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(_a_F_ExecFindPartition_9))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L5
	} else {
		goto L170
	}
L168:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v37)+236))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(96)+v941)))
	v956 = F_OidOutputFunctionCall(m, v951, v955)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L5
	} else {
		goto L169
	}
L169:
	;
	v958 = v956
	goto L167
L170:
	;
	if v958&int32(3) == int32(0) {
		v987 = v958
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v1040 = v904 + int32(1)
	if v1040 != v683 {
		v904 = v1040
		goto L163
	} else {
		goto L196
	}
L172:
	;
	if v1020 <= int32(64) {
		goto L189
	} else {
		goto L190
	}
L173:
	;
	v1020 = v1012 - v958
	goto L172
L174:
	;
	v991 = v987
	goto L183
L175:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v958))))
	if v971 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v1020 = int32(0)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v976 = v958
	goto L179
L179:
	;
	v980 = v976 + int32(1)
	if v980&int32(3) == int32(0) {
		v987 = v980
		goto L174
	} else {
		goto L181
	}
L180:
	;
	v1012 = v980
	goto L173
L181:
	;
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v980))))
	if v985 != 0 {
		v976 = v980
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v1000 = int32(-2139062144)
	if (int32(16843008)-v997|v997)&v1000 == v1000 {
		v991 = v991 + int32(4)
		goto L183
	} else {
		goto L185
	}
L184:
	;
	v1006 = v991
	goto L186
L185:
	;
	goto L184
L186:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006))))
	if v1010 != 0 {
		v1006 = v1006 + int32(1)
		goto L186
	} else {
		goto L188
	}
L187:
	;
	v1012 = v1006
	goto L173
L188:
	;
	goto L187
L189:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v958, v1020)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L5
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1030 = F_pg_mbcliplen(m, v958, v1020, int32(64))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L5
	} else {
		goto L193
	}
L192:
	;
	goto L171
L193:
	;
	F_appendBinaryStringInfo(m, v37+int32(240), v958, v1030)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L5
	} else {
		goto L194
	}
L194:
	;
	F_appendStringInfoString(m, v37+int32(240), int32(_a_F_ExecFindPartition_8))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L5
	} else {
		goto L195
	}
L195:
	;
	goto L171
L196:
	;
	goto L164
L197:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v37)+240))
	v1091 = v1081
	goto L114
L198:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L5
	} else {
		goto L199
	}
L199:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v112)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v1123 + int32(4)
	F_errmsg(m, int32(_a_F_ExecFindPartition_10), v37+int32(16))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L5
	} else {
		goto L200
	}
L200:
	;
	if v1091 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v1091
	F_errdetail(m, int32(_a_F_ExecFindPartition_11), v37)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L5
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	F_errtable(m, v112)
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L5
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(335), int32(_a_F_ExecFindPartition_12))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L5
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v2184)+32))
	if v2185 == v1153 {
		goto L384
	} else {
		goto L385
	}
L208:
	;
	if int32(0) <= v1186 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	if int32(0) <= v1186 {
		goto L374
	} else {
		goto L375
	}
L211:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1194+v1186<<(uint(int32(2))%32))))
	v2155 = v1198
	v2160 = int32(0)
	v2165 = v88
	v2172 = v95
	goto L207
L212:
	;
	goto L213
L213:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1200+v1182)))
	v1205 = F_ExecLookupResultRelByOid(m, l0, v1202, int32(1), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L5
	} else {
		goto L214
	}
L214:
	;
	if v1205 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1208 != 0 {
		goto L218
	} else {
		goto L219
	}
L216:
	;
	goto L217
L217:
	;
	v1219 = int32(0)
	v1221 = m.G0
	v1223 = v1221 - int32(16)
	m.G0 = v1223
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1225)+8))
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1226+v1153<<(uint(int32(2))%32))))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+8))
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1232)+4))
	v1235 = int32(_a_F_ExecFindPartition_0)
	v1236 = *(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0]))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1238
	v1241 = F_table_open(m, v1230, int32(3))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L5
	} else {
		goto L223
	}
L218:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1208)+132))
	v1211 = v1209
	goto L220
L219:
	;
	v1211 = int32(0)
	goto L220
L220:
	;
	F_CheckValidResultRel(m, v1205, int32(3), v1211, int32(0))
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1205, v1153, int32(1))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L5
	} else {
		goto L222
	}
L222:
	;
	v2155 = v1205
	v2160 = int32(0)
	v2165 = v88
	v2172 = v95
	goto L207
L223:
	;
	v1244 = F_palloc0(m, int32(216))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244))) = int32(388)
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(l4)+132))
	F_InitResultRelInfo(m, v1244, v1241, int32(0), l1, v1249)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L5
	} else {
		goto L225
	}
L225:
	;
	if v1231 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+132))
	v1255 = v1253
	goto L228
L227:
	;
	v1255 = int32(0)
	goto L228
L228:
	;
	F_CheckValidResultRel(m, v1244, int32(3), v1255, int32(0))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L5
	} else {
		goto L229
	}
L229:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259)+116)))
	if v1260 != int32(1) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	if v1231 != 0 {
		goto L238
	} else {
		goto L239
	}
L231:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+16))
	if v1263 != 0 {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	if v1231 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+132))
	v1268 = base.B2i32(v1264 != int32(0))
	goto L235
L234:
	;
	v1268 = int32(0)
	goto L235
L235:
	;
	F_ExecOpenIndices(m, v1244, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L5
	} else {
		goto L236
	}
L236:
	;
	goto L230
L237:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(l4)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1918
	v1920 = *(*int32)(unsafe.Add(mBase, uint32(l4)+80))
	v1921 = F_lappend(m, v1920, v1244)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L5
	} else {
		goto L335
	}
L238:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+100))
	if v1271 != 0 {
		goto L244
	} else {
		goto L245
	}
L239:
	;
	goto L240
L240:
	;
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1244, v1153, int32(0))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L5
	} else {
		goto L334
	}
L241:
	;
	v1509 = int32(0)
	F_ExecInitRoutingInfo(m, l0, l4, l2, v83, v1244, v1153, v1509)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L5
	} else {
		goto L263
	}
L242:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1462)+72))
	v1466 = F_map_variable_attnos(m, v1435, v1234, v1444, v1463, v1223+int32(15))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L5
	} else {
		goto L261
	}
L243:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+52))
	v1426 = F_build_attrmap_by_name(m, v1423, v1424, int32(0))
	mBase = m.M
	v1427 = m.ExcPending
	if v1427 != 0 {
		goto L5
	} else {
		goto L260
	}
L244:
	;
	v1272 = int32(0)
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+12))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)))
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+52))
	v1278 = F_build_attrmap_by_name(m, v1275, v1276, v1272)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L5
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+112))
	if v1383 == int32(0) {
		v1491 = v1219
		goto L241
	} else {
		goto L259
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+120)) = v1353
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+116)) = v1284
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+112))
	if v1376 == int32(0) {
		v1491 = v1278
		goto L241
	} else {
		goto L257
	}
L248:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1280)+72))
	v1284 = F_map_variable_attnos(m, v1274, v1234, v1278, v1281, v1223+int32(15))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L5
	} else {
		goto L249
	}
L249:
	;
	if v1284 == int32(0) {
		v1353 = v1272
		goto L247
	} else {
		goto L250
	}
L250:
	;
	v1288 = int32(0)
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+4))
	if v1289 <= v1288 {
		v1353 = v1272
		goto L247
	} else {
		goto L251
	}
L251:
	;
	v1299 = v1288
	v1305 = v1272
	goto L252
L252:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1326+v1299<<(uint(int32(2))%32))))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1330)+16))
	v1332 = F_ExecInitQual(m, v1331, l0)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L5
	} else {
		goto L254
	}
L253:
	;
	v1353 = v1334
	goto L247
L254:
	;
	v1334 = F_lappend(m, v1305, v1332)
	mBase = m.M
	v1335 = m.ExcPending
	if v1335 != 0 {
		goto L5
	} else {
		goto L255
	}
L255:
	;
	v1337 = v1299 + int32(1)
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+4))
	if v1337 < v1338 {
		v1299 = v1337
		v1305 = v1334
		goto L252
	} else {
		goto L256
	}
L256:
	;
	goto L253
L257:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1379)))
	if v1278 == int32(0) {
		v1422 = v1380
		goto L243
	} else {
		goto L258
	}
L258:
	;
	v1435 = v1380
	v1444 = v1278
	goto L242
L259:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+12))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1386)))
	v1422 = v1387
	goto L243
L260:
	;
	v1435 = v1422
	v1444 = v1426
	goto L242
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+148)) = v1466
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1472 = F_ExecBuildProjectionInfo(m, v1466, v1469, v1470, l0, v1471)
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L5
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+152)) = v1472
	v1491 = v1444
	goto L241
L263:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+132))
	if v1513 == int32(0) {
		v1899 = v1491
		goto L237
	} else {
		goto L264
	}
L264:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1518 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L5
	} else {
		goto L331
	}
L266:
	;
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+8))
	v1520 = F_RelationGetIndexList(m, v1519)
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L5
	} else {
		goto L270
	}
L267:
	;
	v1760 = v1509
	goto L268
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+156)) = v1760
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+132))
	if v1777 != int32(2) {
		v1899 = v1491
		goto L237
	} else {
		goto L308
	}
L269:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1735 != 0 {
		goto L301
	} else {
		goto L302
	}
L270:
	;
	if v1520 == int32(0) {
		v1719 = v1509
		goto L269
	} else {
		goto L271
	}
L271:
	;
	v1524 = int32(0)
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+4))
	if v1525 <= v1524 {
		v1719 = v1509
		goto L269
	} else {
		goto L272
	}
L272:
	;
	v1538 = v1524
	v1546 = v1509
	goto L273
L273:
	;
	v1562 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+12))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1562+v1538<<(uint(int32(2))%32))))
	v1567 = F_get_partition_ancestors(m, v1566)
	mBase = m.M
	v1568 = m.ExcPending
	if v1568 != 0 {
		goto L5
	} else {
		goto L275
	}
L274:
	;
	v1719 = v1679
	goto L269
L275:
	;
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(l1)+156))
	if v1569 == int32(0) {
		v1679 = v1546
		goto L276
	} else {
		goto L277
	}
L276:
	;
	F_list_free(m, v1567)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L5
	} else {
		goto L299
	}
L277:
	;
	v1572 = int32(0)
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v1569)+4))
	if v1573 <= v1572 {
		v1679 = v1546
		goto L276
	} else {
		goto L278
	}
L278:
	;
	v1583 = v1572
	v1594 = v1546
	goto L279
L279:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1569)+12))
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1610+v1583<<(uint(int32(2))%32))))
	v1615 = int32(0)
	if v1567 == v1615 {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	v1679 = v1656
	goto L276
L281:
	;
	if v1653 != 0 {
		goto L294
	} else {
		goto L295
	}
L282:
	;
	v1653 = int32(0)
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+4))
	if v1621 <= int32(0) {
		v1646 = v1615
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1653 = v1646
	goto L281
L286:
	;
	v1624 = int32(0)
	if v1624 < v1621 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1627 = v1621
	goto L289
L288:
	;
	v1627 = v1624
	goto L289
L289:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+12))
	v1630 = int32(0)
	goto L290
L290:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1628+v1630<<(uint(int32(2))%32))))
	v1639 = base.B2i32(v1638 == v1614)
	if v1638 == v1614 {
		v1646 = v1639
		goto L285
	} else {
		goto L292
	}
L291:
	;
	v1646 = v1639
	goto L285
L292:
	;
	v1641 = v1630 + int32(1)
	if v1641 != v1627 {
		v1630 = v1641
		goto L290
	} else {
		goto L293
	}
L293:
	;
	goto L291
L294:
	;
	v1654 = F_lappend_oid(m, v1594, v1566)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L5
	} else {
		goto L297
	}
L295:
	;
	v1656 = v1594
	goto L296
L296:
	;
	v1658 = v1583 + int32(1)
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1569)+4))
	if v1658 < v1659 {
		v1583 = v1658
		v1594 = v1656
		goto L279
	} else {
		goto L298
	}
L297:
	;
	v1656 = v1654
	goto L296
L298:
	;
	goto L280
L299:
	;
	v1698 = v1538 + int32(1)
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v1520)+4))
	if v1698 < v1699 {
		v1538 = v1698
		v1546 = v1679
		goto L273
	} else {
		goto L300
	}
L300:
	;
	goto L274
L301:
	;
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+4))
	v1737 = v1736
	goto L303
L302:
	;
	v1737 = v1219
	goto L303
L303:
	;
	if v1719 != 0 {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+4))
	v1740 = v1738
	goto L306
L305:
	;
	v1740 = int32(0)
	goto L306
L306:
	;
	if v1740 != v1737 {
		goto L265
	} else {
		goto L307
	}
L307:
	;
	v1760 = v1719
	goto L268
L308:
	;
	v1781 = F_palloc0(m, int32(20))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L5
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781))) = int32(386)
	v1785 = F_ExecGetRootToChildMap(m, v1244, l4)
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L5
	} else {
		goto L310
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+160)) = v1781
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+8))
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1792 = F_table_slot_create(m, v1788, v1789+int32(104))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L5
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+4)) = v1792
	if v1785 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+8)) = v1798
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+12)) = v1801
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l1)+160))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+16)) = v1804
	v1899 = v1491
	goto L237
L313:
	;
	goto L314
L314:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+140))
	v1807 = F_copyObjectImpl(m, v1806)
	mBase = m.M
	v1808 = m.ExcPending
	if v1808 != 0 {
		goto L5
	} else {
		goto L315
	}
L315:
	;
	if v1491 == int32(0) {
		goto L316
	} else {
		goto L317
	}
L316:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+52))
	v1815 = F_build_attrmap_by_name(m, v1812, v1813, int32(0))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L5
	} else {
		goto L319
	}
L317:
	;
	v1817 = v1491
	goto L318
L318:
	;
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+72))
	v1822 = F_map_variable_attnos(m, v1807, int32(-1), v1817, v1819, v1223+int32(15))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L5
	} else {
		goto L320
	}
L319:
	;
	v1817 = v1815
	goto L318
L320:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+72))
	v1828 = F_map_variable_attnos(m, v1822, v1234, v1817, v1825, v1223+int32(15))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L5
	} else {
		goto L321
	}
L321:
	;
	v1830 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+144))
	v1831 = F_ExecGetChildToRootMap(m, v1244)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L5
	} else {
		goto L322
	}
L322:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1831)+8))
	v1834 = F_adjust_partition_colnos_using_map(m, v1830, v1833)
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L5
	} else {
		goto L323
	}
L323:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1839 = F_table_slot_create(m, v1241, v1836+int32(104))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L5
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+8)) = v1839
	v1843 = F_ExecBuildUpdateProjection(m, v1828, int32(1), v1834, v1517, v1516, v1839, l0)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L5
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+12)) = v1843
	v1846 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+148))
	if v1846 == int32(0) {
		v1899 = v1817
		goto L237
	} else {
		goto L326
	}
L326:
	;
	v1849 = F_copyObjectImpl(m, v1846)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L5
	} else {
		goto L327
	}
L327:
	;
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+72))
	v1856 = F_map_variable_attnos(m, v1849, int32(-1), v1817, v1853, v1223+int32(15))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L5
	} else {
		goto L328
	}
L328:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1858)+72))
	v1862 = F_map_variable_attnos(m, v1856, v1234, v1817, v1859, v1223+int32(15))
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L5
	} else {
		goto L329
	}
L329:
	;
	v1864 = F_ExecInitQual(m, v1862, l0)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L5
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1781)+16)) = v1864
	v1899 = v1817
	goto L237
L331:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_13), int32(0))
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(730), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L5
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	v1899 = v1219
	goto L237
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+80)) = v1921
	if v1231 == int32(0) {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecFindPartition[0])) = v1236
	m.G0 = v1223 + int32(16)
	v2155 = v1244
	v2160 = int32(0)
	v2165 = v88
	v2172 = v95
	goto L207
L337:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+72))
	if v1926 != int32(5) {
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+160))
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1930)+12))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1931)))
	if v1899 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1233)+52))
	v1938 = F_build_attrmap_by_name(m, v1935, v1936, int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L5
	} else {
		goto L342
	}
L340:
	;
	v1940 = v1899
	goto L341
L341:
	;
	v1941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244)+48)))
	if v1941 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L342:
	;
	v1940 = v1938
	goto L341
L343:
	;
	F_ExecInitMergeTupleSlots(m, l0, v1244)
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L5
	} else {
		goto L346
	}
L344:
	;
	goto L345
L345:
	;
	v1947 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+164))
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1947)+12))
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1948)))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1950)+72))
	v1954 = F_map_variable_attnos(m, v1949, v1234, v1940, v1951, v1223+int32(15))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L5
	} else {
		goto L347
	}
L346:
	;
	goto L345
L347:
	;
	v1956 = F_ExecInitQual(m, v1954, l0)
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L5
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+176)) = v1956
	if v1932 == int32(0) {
		goto L336
	} else {
		goto L349
	}
L349:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+4))
	if v1961 <= int32(0) {
		goto L336
	} else {
		goto L350
	}
L350:
	;
	v1965 = v1244 + int32(164)
	v1985 = int32(0)
	goto L351
L351:
	;
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+12))
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v2000+v1985<<(uint(int32(2))%32))))
	v2005 = F_copyObjectImpl(m, v2004)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L5
	} else {
		goto L353
	}
L352:
	;
	goto L336
L353:
	;
	v2008 = F_palloc0(m, int32(16))
	mBase = m.M
	v2009 = m.ExcPending
	if v2009 != 0 {
		goto L5
	} else {
		goto L354
	}
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+4)) = v2005
	*(*int32)(unsafe.Add(mBase, uint32(v2008))) = int32(387)
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+4))
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1965+v2013<<(uint(int32(2))%32))))
	v2018 = F_lappend(m, v2017, v2008)
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L5
	} else {
		goto L355
	}
L355:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+4))
	v2021 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1965+v2020<<(uint(v2021)%32)))) = v2018
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+8))
	switch v2025 - v2021 {
	case 0:
		goto L360
	case 1:
		goto L358
	case 2, 5:
		goto L356
	default:
		goto L359
	}
L356:
	;
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+16))
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+48))
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+72))
	v2068 = F_map_variable_attnos(m, v2063, v1234, v1940, v2065, v1223+int32(15))
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L5
	} else {
		goto L370
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+8)) = v2060
	goto L356
L358:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+20))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+40))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1241)+52))
	v2057 = F_ExecBuildProjectionInfo(m, v2054, v1929, v2055, l0, v2056)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L5
	} else {
		goto L369
	}
L359:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L5
	} else {
		goto L366
	}
L360:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+24))
	if v1940 != 0 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v2029 = F_adjust_partition_colnos_using_map(m, v2028, v1940)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L5
	} else {
		goto L364
	}
L362:
	;
	v2032 = v2028
	goto L363
L363:
	;
	v2033 = *(*int32)(unsafe.Add(mBase, uint32(v2005)+20))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+8))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2035)+52))
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v1244)+40))
	v2039 = F_ExecBuildUpdateProjection(m, v2033, int32(1), v2032, v2036, v1929, v2037, int32(0))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L5
	} else {
		goto L365
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+24)) = v2029
	v2032 = v2029
	goto L363
L365:
	;
	v2060 = v2039
	goto L357
L366:
	;
	F_errmsg_internal(m, int32(_a_F_ExecFindPartition_15), int32(0))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L5
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_ExecFindPartition_2), int32(968), int32(_a_F_ExecFindPartition_14))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L5
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	v2060 = v2057
	goto L357
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2005)+16)) = v2068
	v2071 = F_ExecInitQual(m, v2068, l0)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L5
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+12)) = v2071
	v2075 = v1985 + int32(1)
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v1932)+4))
	if v2075 < v2076 {
		v1985 = v2075
		goto L351
	} else {
		goto L372
	}
L372:
	;
	goto L352
L373:
	;
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2137)))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2138)+16))
	if v2140 == int32(0) {
		v2155 = v2139
		v2160 = v2138
		v2165 = v88
		v2172 = v95
		goto L207
	} else {
		goto L378
	}
L374:
	;
	v2121 = v1186 << (uint(int32(2)) % 32)
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v2121+v39)))
	v2137 = v2121 + v2122
	v2138 = v2125
	goto L373
L375:
	;
	goto L376
L376:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v2126+v1182)))
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v2130 = F_ExecInitPartitionDispatchInfo(m, l4, l2, v2128, v83, v1153, v2129)
	mBase = m.M
	v2131 = m.ExcPending
	if v2131 != 0 {
		goto L5
	} else {
		goto L377
	}
L377:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v1185)))
	v2137 = v2132 + v2133<<(uint(int32(2))%32)
	v2138 = v2130
	goto L373
L378:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v2138)+20))
	v2144 = F_execute_attr_map_slot(m, v2143, v88, v2140)
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L5
	} else {
		goto L379
	}
L379:
	;
	if v95 != 0 {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v2147 = *(*int32)(unsafe.Add(mBase, uint32(v2146)+12))
	m.T0[v2147].(func(*base.Module, int32))(m, v95)
	mBase = m.M
	v2149 = m.ExcPending
	if v2149 != 0 {
		goto L5
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	v2155 = v2139
	v2160 = v2138
	v2165 = v2144
	v2172 = v2140
	goto L207
L383:
	;
	goto L382
L384:
	;
	if v1189 == int32(0) {
		v2198 = v2165
		goto L387
	} else {
		goto L388
	}
L385:
	;
	v2203 = v2165
	goto L386
L386:
	;
	if v2160 != 0 {
		v83 = v2160
		v88 = v2203
		v95 = v2172
		goto L15
	} else {
		goto L393
	}
L387:
	;
	v2200 = F_ExecPartitionCheck(m, v2155, v2198, l4, int32(1))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L5
	} else {
		goto L392
	}
L388:
	;
	v2189 = F_ExecGetRootToChildMap(m, v2155, l4)
	mBase = m.M
	v2190 = m.ExcPending
	if v2190 != 0 {
		goto L5
	} else {
		goto L389
	}
L389:
	;
	if v2189 == int32(0) {
		v2198 = l3
		goto L387
	} else {
		goto L390
	}
L390:
	;
	v2193 = *(*int32)(unsafe.Add(mBase, uint32(v2189)+8))
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v2155)+204))
	v2195 = F_execute_attr_map_slot(m, v2193, l3, v2194)
	mBase = m.M
	v2196 = m.ExcPending
	if v2196 != 0 {
		goto L5
	} else {
		goto L391
	}
L391:
	;
	v2198 = v2195
	goto L387
L392:
	;
	v2203 = v2198
	goto L386
L393:
	;
	goto L16
L394:
	;
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(v2172)+8))
	v2207 = *(*int32)(unsafe.Add(mBase, uint32(v2206)+12))
	m.T0[v2207].(func(*base.Module, int32))(m, v2172)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L5
	} else {
		goto L395
	}
L395:
	;
	v2215 = v2155
	goto L13
}
func F_ExecPartitionCheckEmitError(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
		v20 = F_build_attrmap_by_name_if_req(m, v17, v18, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			if v20 == int32(0) {
				v32 = l1
				v33 = v13
				v34 = v18
				v35 = v15
				v37 = F_ExecGetInsertedCols(m, v33, l2)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = F_ExecGetUpdatedCols(m, v33, l2)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = F_bms_union(m, v37, v39)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									F_errcode(m, int32(67391682))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
										F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											if v43 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
												F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													F_errtable(m, v66)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_errtable(m, v66)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
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
			} else {
				v25 = F_MakeTupleTableSlot(m, v18, int32(_a_F_ExecPartitionCheckEmitError_4))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = F_execute_attr_map_slot(m, v20, l1, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v32 = v27
						v33 = v13
						v34 = v18
						v35 = v15
						v37 = F_ExecGetInsertedCols(m, v33, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = F_ExecGetUpdatedCols(m, v33, l2)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v41 = F_bms_union(m, v37, v39)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											F_errcode(m, int32(67391682))
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return
											} else {
												v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
												F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													if v43 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
														F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return
														} else {
															v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															F_errtable(m, v66)
															mBase = m.M
															v68 = m.ExcPending
															if v68 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														F_errtable(m, v66)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
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
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+56))
		v32 = l1
		v33 = l0
		v34 = v30
		v35 = v31
		v37 = F_ExecGetInsertedCols(m, v33, l2)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = F_ExecGetUpdatedCols(m, v33, l2)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = F_bms_union(m, v37, v39)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = F_ExecBuildSlotValueDescription(m, v35, v32, v34, v41)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errcode(m, int32(67391682))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v53 + int32(4)
								F_errmsg(m, int32(_a_F_ExecPartitionCheckEmitError_0), v11+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v43 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v43
										F_errdetail(m, int32(_a_F_ExecPartitionCheckEmitError_1), v11)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											F_errtable(m, v66)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
												mBase = m.M
												v73 = m.ExcPending
												if v73 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_errtable(m, v66)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ExecPartitionCheckEmitError_2), int32(1969), int32(_a_F_ExecPartitionCheckEmitError_3))
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
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
	}
}
func F_get_partition_parent(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = F_get_partition_parent_worker(m, v12, l0, v8+int32(31))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				if l1 == int32(0) {
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+31)))
					if v22&int32(1) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
							F_errmsg_internal(m, int32(_a_F_get_partition_parent_0), v8+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(69), int32(_a_F_get_partition_parent_2))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_sequence_close(m, v12, int32(1))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(32)
							return v18
						}
					}
				} else {
					F_sequence_close(m, v12, int32(1))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(32)
						return v18
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_get_partition_parent_3), v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_get_partition_parent_1), int32(65), int32(_a_F_get_partition_parent_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
	}
}
func F_has_partition_attrs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v4 {
		v134 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v134
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+119)))
	if v18 != int32(112) {
		v134 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v28 = v27
	goto L8
L7:
	;
	v28 = v4
	goto L8
L8:
	;
	if v25 <= int32(0) {
		v134 = v4
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v34 = v4
	v35 = v28
	goto L10
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41+v34<<(uint(int32(1))%32)))))
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v134 = int32(0)
	goto L1
L12:
	;
	v128 = v34 + int32(1)
	if v128 != v25 {
		v34 = v128
		v35 = v125
		goto L10
	} else {
		goto L43
	}
L13:
	;
	v48 = F_bms_is_member(m, v45+int32(7), l1)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	F_pull_varattnos(m, v57, int32(1), v13+int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L4
	} else {
		goto L19
	}
L16:
	;
	if v48 == int32(0) {
		v125 = v35
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(1)
	if l2 == int32(0) {
		v134 = v52
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v55)
	v134 = v52
	goto L1
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v68 = int32(0)
	if l1 == v68 {
		v109 = v68
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v109 != 0 {
		goto L34
	} else {
		goto L35
	}
L21:
	;
	goto L20
L22:
	;
	if v67 == int32(0) {
		v109 = v68
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v77 < v78 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v77
	goto L26
L25:
	;
	v80 = v78
	goto L26
L26:
	;
	if v80 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = int32(1)
	goto L29
L28:
	;
	v83 = v80
	goto L29
L29:
	;
	v84 = int32(8)
	v89 = int32(0)
	goto L30
L30:
	;
	v96 = v89 << (uint(int32(2)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v67+v84+v96)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+(l1+v84))))
	v101 = v98 & v100
	v103 = base.B2i32(v101 != int32(0))
	if v101 != 0 {
		v109 = v103
		goto L21
	} else {
		goto L32
	}
L31:
	;
	v109 = v103
	goto L21
L32:
	;
	v105 = v89 + int32(1)
	if v105 != v83 {
		v89 = v105
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v117 = v35 + int32(4)
	if base.Ui32(v117) < base.Ui32(v66+v65<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v113)
	goto L39
L38:
	;
	goto L39
L39:
	;
	v134 = int32(1)
	goto L1
L40:
	;
	v123 = v117
	goto L42
L41:
	;
	v123 = int32(0)
	goto L42
L42:
	;
	v125 = v123
	goto L12
L43:
	;
	goto L11
}
func F_partition_list_bsearch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	v10 = int32(-1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v13 = v11 - int32(1)
	if v13 < int32(0) {
		v56 = v10
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v56
L2:
	;
	v21 = v10
	v22 = v13
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v30 = int32(2)
	v31 = base.I32_div_s(v21+v22+int32(1), v30)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26+v31<<(uint(v30)%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = F_FunctionCall2Coll(m, l0, v25, v36, l3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v56 = v48
	goto L1
L5:
	;
	if v48 < v49 {
		v21 = v48
		v22 = v49
		goto L3
	} else {
		goto L12
	}
L6:
	;
	return int32(0)
L7:
	;
	if v37 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(base.B2i32(v37 == int32(0)))
	if v37 != 0 {
		v48 = v31
		v49 = v22
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v48 = v21
	v49 = v31 - int32(1)
	goto L5
L11:
	;
	v56 = v31
	goto L1
L12:
	;
	goto L4
}
