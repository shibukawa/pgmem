package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_subquery_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 float64, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
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
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v595 int32
	_ = v595
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int64
	_ = v740
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
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
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v879 int32
	_ = v879
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1058 int32
	_ = v1058
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1150 int32
	_ = v1150
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
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
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1367 int32
	_ = v1367
	var v1453 int32
	_ = v1453
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1464 int32
	_ = v1464
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
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
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1928 int32
	_ = v1928
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1950 int32
	_ = v1950
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1964 int32
	_ = v1964
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2028 int32
	_ = v2028
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2097 int32
	_ = v2097
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2125 int32
	_ = v2125
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2202 int32
	_ = v2202
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2253 int32
	_ = v2253
	var v2276 int32
	_ = v2276
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2311 int32
	_ = v2311
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2339 int32
	_ = v2339
	var v2342 int32
	_ = v2342
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2373 int32
	_ = v2373
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2383 int32
	_ = v2383
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2434 int32
	_ = v2434
	var v2435 int32
	_ = v2435
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2458 int32
	_ = v2458
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2527 int32
	_ = v2527
	var v2529 int32
	_ = v2529
	var v2532 int32
	_ = v2532
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2541 int32
	_ = v2541
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2564 int32
	_ = v2564
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2598 int32
	_ = v2598
	var v2599 int32
	_ = v2599
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2606 int32
	_ = v2606
	var v2609 int32
	_ = v2609
	var v2610 int32
	_ = v2610
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2618 int32
	_ = v2618
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2625 int32
	_ = v2625
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2662 int32
	_ = v2662
	var v2665 int32
	_ = v2665
	var v2672 int32
	_ = v2672
	var v2710 int32
	_ = v2710
	var v2711 int32
	_ = v2711
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2729 int32
	_ = v2729
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2735 int32
	_ = v2735
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2747 int32
	_ = v2747
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2804 int32
	_ = v2804
	var v2807 int32
	_ = v2807
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2812 int32
	_ = v2812
	var v2813 int32
	_ = v2813
	var v2814 int32
	_ = v2814
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2828 int32
	_ = v2828
	var v2831 int32
	_ = v2831
	var v2850 int32
	_ = v2850
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2881 int32
	_ = v2881
	var v2882 int32
	_ = v2882
	var v2885 int32
	_ = v2885
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2913 int32
	_ = v2913
	var v2916 int32
	_ = v2916
	var v2917 int32
	_ = v2917
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2921 int32
	_ = v2921
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2937 int32
	_ = v2937
	var v2941 int32
	_ = v2941
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2961 int32
	_ = v2961
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2967 int32
	_ = v2967
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2975 int32
	_ = v2975
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3019 int32
	_ = v3019
	var v3021 int32
	_ = v3021
	var v3022 int32
	_ = v3022
	var v3025 int32
	_ = v3025
	var v3026 int32
	_ = v3026
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3114 int32
	_ = v3114
	var v3117 int32
	_ = v3117
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3165 int32
	_ = v3165
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3217 int32
	_ = v3217
	var v3220 int32
	_ = v3220
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3230 int32
	_ = v3230
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3237 int32
	_ = v3237
	var v3238 int32
	_ = v3238
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3253 int32
	_ = v3253
	var v3267 int32
	_ = v3267
	var v3293 int32
	_ = v3293
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3302 int32
	_ = v3302
	var v3303 int32
	_ = v3303
	var v3304 int32
	_ = v3304
	var v3307 int32
	_ = v3307
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3333 int32
	_ = v3333
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3339 int32
	_ = v3339
	var v3340 int32
	_ = v3340
	var v3341 int32
	_ = v3341
	var v3342 int32
	_ = v3342
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3372 int32
	_ = v3372
	var v3401 int32
	_ = v3401
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3411 int32
	_ = v3411
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3420 int32
	_ = v3420
	var v3423 int32
	_ = v3423
	var v3424 int32
	_ = v3424
	var v3425 int32
	_ = v3425
	var v3427 int32
	_ = v3427
	var v3428 int32
	_ = v3428
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3433 int32
	_ = v3433
	var v3434 int32
	_ = v3434
	var v3436 int32
	_ = v3436
	var v3439 int32
	_ = v3439
	var v3453 int32
	_ = v3453
	var v3484 int32
	_ = v3484
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3490 int32
	_ = v3490
	var v3491 int32
	_ = v3491
	var v3492 int32
	_ = v3492
	var v3493 int32
	_ = v3493
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3497 int32
	_ = v3497
	var v3498 int32
	_ = v3498
	var v3499 int32
	_ = v3499
	var v3500 int32
	_ = v3500
	var v3503 int32
	_ = v3503
	var v3504 int32
	_ = v3504
	var v3553 int32
	_ = v3553
	var v3557 int32
	_ = v3557
	var v3562 int32
	_ = v3562
	var v3606 int32
	_ = v3606
	var v3607 int32
	_ = v3607
	var v3609 int32
	_ = v3609
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3620 int32
	_ = v3620
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3625 int32
	_ = v3625
	var v3626 int32
	_ = v3626
	var v3628 int32
	_ = v3628
	var v3629 int32
	_ = v3629
	var v3631 int32
	_ = v3631
	var v3632 int32
	_ = v3632
	var v3634 int32
	_ = v3634
	var v3638 int32
	_ = v3638
	var v3639 int32
	_ = v3639
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3694 int32
	_ = v3694
	var v3695 int32
	_ = v3695
	var v3698 int32
	_ = v3698
	var v3699 int32
	_ = v3699
	var v3703 int32
	_ = v3703
	var v3704 int32
	_ = v3704
	var v3792 int32
	_ = v3792
	var v3795 int32
	_ = v3795
	var v3797 int32
	_ = v3797
	var v3799 int32
	_ = v3799
	var v3800 int32
	_ = v3800
	var v3803 int32
	_ = v3803
	var v3805 int64
	_ = v3805
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3810 int32
	_ = v3810
	var v3814 int32
	_ = v3814
	var v3815 int64
	_ = v3815
	var v3816 int32
	_ = v3816
	var v3817 int64
	_ = v3817
	var v3820 int64
	_ = v3820
	var v3822 int64
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3826 int64
	_ = v3826
	var v3827 int32
	_ = v3827
	var v3828 int32
	_ = v3828
	var v3831 int32
	_ = v3831
	var v3835 int32
	_ = v3835
	var v3836 int32
	_ = v3836
	var v3837 int64
	_ = v3837
	var v3838 int64
	_ = v3838
	var v3841 int64
	_ = v3841
	var v3843 int64
	_ = v3843
	var v3849 float64
	_ = v3849
	var v3853 float64
	_ = v3853
	var v3861 float64
	_ = v3861
	var v3870 float64
	_ = v3870
	var v3871 int64
	_ = v3871
	var v3883 float64
	_ = v3883
	var v3893 float64
	_ = v3893
	var v3900 float64
	_ = v3900
	var v3901 float64
	_ = v3901
	var v3905 float64
	_ = v3905
	var v3909 float64
	_ = v3909
	var v3911 float64
	_ = v3911
	var v3913 float64
	_ = v3913
	var v3914 int64
	_ = v3914
	var v3915 int64
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3919 int32
	_ = v3919
	var v3921 int32
	_ = v3921
	var v3923 int32
	_ = v3923
	var v3924 int32
	_ = v3924
	var v3925 int32
	_ = v3925
	var v3928 int32
	_ = v3928
	var v3930 int32
	_ = v3930
	var v3970 int32
	_ = v3970
	var v3971 int32
	_ = v3971
	var v3974 int32
	_ = v3974
	var v3975 int32
	_ = v3975
	var v3979 int32
	_ = v3979
	var v3980 int32
	_ = v3980
	var v3981 int32
	_ = v3981
	var v3984 int32
	_ = v3984
	var v3987 int32
	_ = v3987
	var v3989 int32
	_ = v3989
	var v3990 int32
	_ = v3990
	var v3991 int32
	_ = v3991
	var v3996 int32
	_ = v3996
	var v3997 int32
	_ = v3997
	var v3998 int32
	_ = v3998
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4003 int32
	_ = v4003
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4009 int32
	_ = v4009
	var v4011 int32
	_ = v4011
	var v4012 int32
	_ = v4012
	var v4017 int32
	_ = v4017
	var v4018 int32
	_ = v4018
	var v4019 int32
	_ = v4019
	var v4020 int32
	_ = v4020
	var v4023 int32
	_ = v4023
	var v4024 int32
	_ = v4024
	var v4027 int32
	_ = v4027
	var v4028 int32
	_ = v4028
	var v4031 int32
	_ = v4031
	var v4032 int32
	_ = v4032
	var v4034 int32
	_ = v4034
	var v4039 int32
	_ = v4039
	var v4040 int32
	_ = v4040
	var v4041 int32
	_ = v4041
	var v4042 int32
	_ = v4042
	var v4045 int32
	_ = v4045
	var v4046 int32
	_ = v4046
	var v4047 int32
	_ = v4047
	var v4048 int32
	_ = v4048
	var v4049 int32
	_ = v4049
	var v4050 int32
	_ = v4050
	var v4051 int32
	_ = v4051
	var v4052 int32
	_ = v4052
	var v4053 int32
	_ = v4053
	var v4054 int32
	_ = v4054
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4060 int32
	_ = v4060
	var v4061 int32
	_ = v4061
	var v4062 int32
	_ = v4062
	var v4063 int32
	_ = v4063
	var v4065 int32
	_ = v4065
	var v4068 int32
	_ = v4068
	var v4072 int32
	_ = v4072
	var v4073 int32
	_ = v4073
	var v4113 int32
	_ = v4113
	var v4114 int32
	_ = v4114
	var v4115 int32
	_ = v4115
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4120 int32
	_ = v4120
	var v4121 int32
	_ = v4121
	var v4124 int32
	_ = v4124
	var v4130 int32
	_ = v4130
	var v4132 int32
	_ = v4132
	var v4133 int32
	_ = v4133
	var v4176 int32
	_ = v4176
	var v4183 int32
	_ = v4183
	var v4186 int32
	_ = v4186
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4198 int32
	_ = v4198
	var v4199 int32
	_ = v4199
	var v4203 int32
	_ = v4203
	var v4207 int32
	_ = v4207
	var v4213 int32
	_ = v4213
	var v4216 float64
	_ = v4216
	var v4219 float64
	_ = v4219
	var v4221 int32
	_ = v4221
	var v4223 int32
	_ = v4223
	var v4226 float64
	_ = v4226
	var v4229 int32
	_ = v4229
	var v4263 int32
	_ = v4263
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4276 int32
	_ = v4276
	var v4280 int32
	_ = v4280
	var v4283 int32
	_ = v4283
	var v4284 int32
	_ = v4284
	var v4286 int32
	_ = v4286
	var v4288 int32
	_ = v4288
	var v4298 float64
	_ = v4298
	var v4299 float64
	_ = v4299
	var v4300 float64
	_ = v4300
	var v4301 float64
	_ = v4301
	var v4302 float64
	_ = v4302
	var v4303 int32
	_ = v4303
	var v4304 int32
	_ = v4304
	var v4305 float64
	_ = v4305
	var v4309 float64
	_ = v4309
	var v4311 float64
	_ = v4311
	var v4319 int32
	_ = v4319
	var v4320 int32
	_ = v4320
	var v4321 int32
	_ = v4321
	var v4322 int32
	_ = v4322
	var v4323 int32
	_ = v4323
	var v4325 int32
	_ = v4325
	var v4328 int32
	_ = v4328
	var v4330 int32
	_ = v4330
	var v4331 int32
	_ = v4331
	var v4335 int32
	_ = v4335
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4340 int32
	_ = v4340
	var v4341 int32
	_ = v4341
	var v4346 int32
	_ = v4346
	var v4347 int32
	_ = v4347
	var v4348 int32
	_ = v4348
	var v4351 int32
	_ = v4351
	var v4359 int32
	_ = v4359
	var v4397 int32
	_ = v4397
	var v4401 int32
	_ = v4401
	var v4406 int32
	_ = v4406
	var v4410 int32
	_ = v4410
	var v4413 int32
	_ = v4413
	var v4417 int32
	_ = v4417
	var v4421 int32
	_ = v4421
	var v4426 int32
	_ = v4426
	var v4427 int32
	_ = v4427
	var v4428 int32
	_ = v4428
	var v4429 int32
	_ = v4429
	var v4430 int32
	_ = v4430
	var v4431 int32
	_ = v4431
	var v4433 int32
	_ = v4433
	var v4436 int32
	_ = v4436
	var v4440 int32
	_ = v4440
	var v4447 int32
	_ = v4447
	var v4457 int32
	_ = v4457
	var v4481 int32
	_ = v4481
	var v4485 int32
	_ = v4485
	var v4486 int32
	_ = v4486
	var v4489 int32
	_ = v4489
	var v4490 int32
	_ = v4490
	var v4493 int32
	_ = v4493
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4498 int32
	_ = v4498
	var v4504 int32
	_ = v4504
	var v4505 int32
	_ = v4505
	var v4506 int32
	_ = v4506
	var v4509 int32
	_ = v4509
	var v4511 int32
	_ = v4511
	var v4530 int32
	_ = v4530
	var v4555 int32
	_ = v4555
	var v4556 int32
	_ = v4556
	var v4557 int32
	_ = v4557
	var v4558 int32
	_ = v4558
	var v4559 int32
	_ = v4559
	var v4560 int32
	_ = v4560
	var v4564 int32
	_ = v4564
	var v4565 int32
	_ = v4565
	var v4566 int32
	_ = v4566
	var v4567 int32
	_ = v4567
	var v4568 int32
	_ = v4568
	var v4570 int32
	_ = v4570
	var v4571 int32
	_ = v4571
	var v4573 int32
	_ = v4573
	var v4574 int32
	_ = v4574
	var v4575 int32
	_ = v4575
	var v4577 int32
	_ = v4577
	var v4583 int32
	_ = v4583
	var v4584 int32
	_ = v4584
	var v4587 int32
	_ = v4587
	var v4598 int32
	_ = v4598
	var v4607 int32
	_ = v4607
	var v4631 int32
	_ = v4631
	var v4635 int32
	_ = v4635
	var v4636 int32
	_ = v4636
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4644 int32
	_ = v4644
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4649 int32
	_ = v4649
	var v4652 int32
	_ = v4652
	var v4654 int32
	_ = v4654
	var v4655 int32
	_ = v4655
	var v4668 int32
	_ = v4668
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4705 int32
	_ = v4705
	var v4706 int32
	_ = v4706
	var v4709 int32
	_ = v4709
	var v4712 int32
	_ = v4712
	var v4714 int32
	_ = v4714
	var v4721 int32
	_ = v4721
	var v4755 int32
	_ = v4755
	var v4756 int32
	_ = v4756
	var v4760 int32
	_ = v4760
	var v4761 int32
	_ = v4761
	var v4762 int32
	_ = v4762
	var v4764 int32
	_ = v4764
	var v4765 int32
	_ = v4765
	var v4769 int32
	_ = v4769
	var v4770 int32
	_ = v4770
	var v4771 int32
	_ = v4771
	var v4773 int32
	_ = v4773
	var v4774 int32
	_ = v4774
	var v4775 int32
	_ = v4775
	var v4781 int32
	_ = v4781
	var v4784 int32
	_ = v4784
	var v4788 int32
	_ = v4788
	var v4792 int32
	_ = v4792
	var v4797 int32
	_ = v4797
	var v4798 int32
	_ = v4798
	var v4799 int32
	_ = v4799
	var v4800 int32
	_ = v4800
	var v4803 int32
	_ = v4803
	var v4804 int32
	_ = v4804
	var v4807 int32
	_ = v4807
	var v4811 int32
	_ = v4811
	var v4812 int32
	_ = v4812
	var v4817 int32
	_ = v4817
	var v4821 int32
	_ = v4821
	var v4826 int32
	_ = v4826
	var v4830 int32
	_ = v4830
	var v4834 int32
	_ = v4834
	var v4839 int32
	_ = v4839
	var v4843 int32
	_ = v4843
	var v4846 int32
	_ = v4846
	var v4847 int32
	_ = v4847
	var v4848 int32
	_ = v4848
	var v4849 int32
	_ = v4849
	var v4850 int32
	_ = v4850
	var v4852 int32
	_ = v4852
	var v4857 int32
	_ = v4857
	var v4859 int32
	_ = v4859
	var v4865 int32
	_ = v4865
	var v4870 int32
	_ = v4870
	var v4885 int32
	_ = v4885
	var v4914 int32
	_ = v4914
	var v4917 int32
	_ = v4917
	var v4919 int32
	_ = v4919
	var v4930 int32
	_ = v4930
	var v4941 int32
	_ = v4941
	var v4964 int32
	_ = v4964
	var v4968 int32
	_ = v4968
	var v4982 int32
	_ = v4982
	var v5016 int32
	_ = v5016
	var v5017 int32
	_ = v5017
	var v5018 int32
	_ = v5018
	var v5020 int32
	_ = v5020
	var v5021 int32
	_ = v5021
	var v5022 int32
	_ = v5022
	var v5023 int32
	_ = v5023
	var v5024 int32
	_ = v5024
	var v5025 int32
	_ = v5025
	var v5026 int32
	_ = v5026
	var v5029 int32
	_ = v5029
	var v5030 int32
	_ = v5030
	var v5031 int32
	_ = v5031
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5039 int32
	_ = v5039
	var v5040 int32
	_ = v5040
	var v5050 int32
	_ = v5050
	var v5057 int32
	_ = v5057
	var v5078 int32
	_ = v5078
	var v5082 int32
	_ = v5082
	var v5083 int32
	_ = v5083
	var v5085 int32
	_ = v5085
	var v5095 int32
	_ = v5095
	var v5096 int32
	_ = v5096
	var v5129 int32
	_ = v5129
	var v5133 int32
	_ = v5133
	var v5134 int32
	_ = v5134
	var v5135 int32
	_ = v5135
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5148 int32
	_ = v5148
	var v5149 int32
	_ = v5149
	var v5183 int32
	_ = v5183
	var v5185 int32
	_ = v5185
	var v5186 int32
	_ = v5186
	var v5195 int32
	_ = v5195
	var v5236 int32
	_ = v5236
	var v5271 int32
	_ = v5271
	var v5273 int32
	_ = v5273
	var v5274 int32
	_ = v5274
	var v5288 int32
	_ = v5288
	var v5289 int32
	_ = v5289
	var v5291 int32
	_ = v5291
	var v5294 int32
	_ = v5294
	var v5295 int32
	_ = v5295
	var v5300 int32
	_ = v5300
	var v5308 int32
	_ = v5308
	var v5310 int32
	_ = v5310
	var v5312 int32
	_ = v5312
	var v5313 int32
	_ = v5313
	var v5316 int32
	_ = v5316
	var v5320 int32
	_ = v5320
	var v5328 int32
	_ = v5328
	var v5332 int32
	_ = v5332
	var v5333 int32
	_ = v5333
	var v5334 int32
	_ = v5334
	var v5335 int32
	_ = v5335
	var v5338 int32
	_ = v5338
	var v5339 int32
	_ = v5339
	var v5342 int32
	_ = v5342
	var v5346 int32
	_ = v5346
	var v5351 int32
	_ = v5351
	var v5363 int32
	_ = v5363
	var v5387 int32
	_ = v5387
	var v5392 int32
	_ = v5392
	var v5393 int32
	_ = v5393
	var v5397 int32
	_ = v5397
	var v5399 int32
	_ = v5399
	var v5402 int32
	_ = v5402
	var v5409 int32
	_ = v5409
	var v5446 int32
	_ = v5446
	var v5447 int32
	_ = v5447
	var v5456 int32
	_ = v5456
	var v5457 int32
	_ = v5457
	var v5459 int32
	_ = v5459
	var v5462 int32
	_ = v5462
	var v5463 int32
	_ = v5463
	var v5468 int32
	_ = v5468
	var v5475 int32
	_ = v5475
	var v5477 int32
	_ = v5477
	var v5479 int32
	_ = v5479
	var v5482 int32
	_ = v5482
	var v5484 int32
	_ = v5484
	var v5486 int32
	_ = v5486
	var v5493 int32
	_ = v5493
	var v5500 int32
	_ = v5500
	var v5501 int32
	_ = v5501
	var v5502 int32
	_ = v5502
	var v5507 int32
	_ = v5507
	var v5508 int32
	_ = v5508
	var v5519 int32
	_ = v5519
	var v5520 int32
	_ = v5520
	var v5521 int32
	_ = v5521
	var v5615 int32
	_ = v5615
	var v5616 int32
	_ = v5616
	var v5633 int32
	_ = v5633
	var v5655 int32
	_ = v5655
	var v5656 int32
	_ = v5656
	var v5660 int32
	_ = v5660
	var v5699 int32
	_ = v5699
	var v5701 int32
	_ = v5701
	var v5703 int32
	_ = v5703
	var v5704 int32
	_ = v5704
	var v5718 int32
	_ = v5718
	var v5720 int32
	_ = v5720
	var v5721 int32
	_ = v5721
	var v5722 int32
	_ = v5722
	var v5726 int32
	_ = v5726
	var v5727 int32
	_ = v5727
	var v5729 int32
	_ = v5729
	var v5730 int32
	_ = v5730
	var v5734 int32
	_ = v5734
	var v5735 int32
	_ = v5735
	var v5737 int32
	_ = v5737
	var v5738 int32
	_ = v5738
	var v5739 int32
	_ = v5739
	var v5743 int32
	_ = v5743
	var v5750 int32
	_ = v5750
	var v5756 int32
	_ = v5756
	var v5784 int32
	_ = v5784
	var v5785 int32
	_ = v5785
	var v5787 int32
	_ = v5787
	var v5790 int32
	_ = v5790
	var v5793 int32
	_ = v5793
	var v5794 int32
	_ = v5794
	var v5795 int32
	_ = v5795
	var v5801 int32
	_ = v5801
	var v5802 int32
	_ = v5802
	var v5804 int32
	_ = v5804
	var v5843 int32
	_ = v5843
	var v5844 int32
	_ = v5844
	var v5845 int32
	_ = v5845
	var v5847 int32
	_ = v5847
	var v5850 int32
	_ = v5850
	var v5852 int32
	_ = v5852
	var v5858 int32
	_ = v5858
	var v5860 int32
	_ = v5860
	var v5861 int32
	_ = v5861
	var v5862 int32
	_ = v5862
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5866 int32
	_ = v5866
	var v5868 int32
	_ = v5868
	var v5869 int32
	_ = v5869
	var v5871 int32
	_ = v5871
	var v5873 int32
	_ = v5873
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5883 int32
	_ = v5883
	var v5888 int32
	_ = v5888
	var v5890 int32
	_ = v5890
	var v5929 int32
	_ = v5929
	var v5930 int32
	_ = v5930
	var v5931 int32
	_ = v5931
	var v5933 int32
	_ = v5933
	var v5934 int32
	_ = v5934
	var v5936 int32
	_ = v5936
	var v5938 int32
	_ = v5938
	var v5947 int32
	_ = v5947
	var v5985 int32
	_ = v5985
	var v5991 int32
	_ = v5991
	var v6003 int32
	_ = v6003
	var v6029 int32
	_ = v6029
	var v6032 int32
	_ = v6032
	var v6035 int32
	_ = v6035
	var v6036 int32
	_ = v6036
	var v6037 int32
	_ = v6037
	var v6039 int32
	_ = v6039
	var v6043 int32
	_ = v6043
	var v6046 int32
	_ = v6046
	var v6050 int32
	_ = v6050
	var v6052 int32
	_ = v6052
	var v6090 int32
	_ = v6090
	var v6091 int32
	_ = v6091
	var v6094 int32
	_ = v6094
	var v6098 int32
	_ = v6098
	var v6101 int32
	_ = v6101
	var v6102 int32
	_ = v6102
	var v6105 int32
	_ = v6105
	var v6106 int32
	_ = v6106
	var v6107 int32
	_ = v6107
	var v6115 int32
	_ = v6115
	var v6116 int32
	_ = v6116
	var v6123 int32
	_ = v6123
	var v6162 int32
	_ = v6162
	var v6205 int32
	_ = v6205
	var v6209 int32
	_ = v6209
	var v6250 int32
	_ = v6250
	var v6254 int32
	_ = v6254
	var v6255 int32
	_ = v6255
	var v6256 int32
	_ = v6256
	var v6259 int32
	_ = v6259
	var v6308 int32
	_ = v6308
	var v6310 int32
	_ = v6310
	var v6311 int32
	_ = v6311
	var v6312 int32
	_ = v6312
	var v6313 int32
	_ = v6313
	var v6314 int32
	_ = v6314
	var v6362 int32
	_ = v6362
	var v6366 int32
	_ = v6366
	var v6371 int32
	_ = v6371
	var v6374 int32
	_ = v6374
	var v6375 int32
	_ = v6375
	var v6379 int32
	_ = v6379
	var v6380 int32
	_ = v6380
	var v6381 int32
	_ = v6381
	var v6384 int32
	_ = v6384
	var v6393 int32
	_ = v6393
	var v6406 int32
	_ = v6406
	var v6428 int32
	_ = v6428
	var v6429 int32
	_ = v6429
	var v6431 int32
	_ = v6431
	var v6432 int32
	_ = v6432
	var v6434 int32
	_ = v6434
	var v6435 int32
	_ = v6435
	var v6444 int32
	_ = v6444
	var v6445 int32
	_ = v6445
	var v6454 int32
	_ = v6454
	var v6456 int32
	_ = v6456
	var v6457 int32
	_ = v6457
	var v6458 int32
	_ = v6458
	var v6464 int32
	_ = v6464
	var v6470 int32
	_ = v6470
	var v6471 int32
	_ = v6471
	var v6481 int32
	_ = v6481
	var v6514 int32
	_ = v6514
	var v6515 int32
	_ = v6515
	var v6517 int32
	_ = v6517
	var v6520 int32
	_ = v6520
	var v6521 int32
	_ = v6521
	var v6523 int32
	_ = v6523
	var v6524 int32
	_ = v6524
	var v6525 int32
	_ = v6525
	var v6528 int32
	_ = v6528
	var v6537 int32
	_ = v6537
	var v6550 int32
	_ = v6550
	var v6574 int32
	_ = v6574
	var v6592 int32
	_ = v6592
	var v6593 int32
	_ = v6593
	var v6617 int32
	_ = v6617
	var v6618 int32
	_ = v6618
	var v6620 int32
	_ = v6620
	var v6665 int32
	_ = v6665
	var v6676 int32
	_ = v6676
	var v6681 int32
	_ = v6681
	var v6712 int32
	_ = v6712
	var v6713 int32
	_ = v6713
	var v6714 int32
	_ = v6714
	var v6716 int32
	_ = v6716
	var v6731 int32
	_ = v6731
	var v6759 int32
	_ = v6759
	var v6761 int32
	_ = v6761
	var v6762 int32
	_ = v6762
	var v6764 int32
	_ = v6764
	var v6765 int32
	_ = v6765
	var v6767 int32
	_ = v6767
	var v6768 int32
	_ = v6768
	var v6770 int32
	_ = v6770
	var v6772 int32
	_ = v6772
	var v6774 int32
	_ = v6774
	var v6776 int32
	_ = v6776
	var v6787 int32
	_ = v6787
	var v6824 int32
	_ = v6824
	var v6826 int32
	_ = v6826
	var v6828 int32
	_ = v6828
	var v6831 int32
	_ = v6831
	var v6833 int32
	_ = v6833
	var v6835 int32
	_ = v6835
	var v6844 int32
	_ = v6844
	var v6881 int32
	_ = v6881
	var v6883 int32
	_ = v6883
	var v6885 int32
	_ = v6885
	var v6888 int32
	_ = v6888
	var v6890 int32
	_ = v6890
	var v6892 int32
	_ = v6892
	var v6935 int32
	_ = v6935
	var v6949 int32
	_ = v6949
	var v6979 int32
	_ = v6979
	var v6989 int32
	_ = v6989
	var v6993 int32
	_ = v6993
	var v6994 int32
	_ = v6994
	var v6996 int32
	_ = v6996
	var v6999 int32
	_ = v6999
	var v7002 int32
	_ = v7002
	var v7003 int32
	_ = v7003
	var v7004 int32
	_ = v7004
	var v7008 int32
	_ = v7008
	var v7010 int32
	_ = v7010
	var v7017 float64
	_ = v7017
	var v7021 int64
	_ = v7021
	var v7022 int64
	_ = v7022
	var v7024 int32
	_ = v7024
	var v7028 int32
	_ = v7028
	var v7030 int32
	_ = v7030
	var v7031 int32
	_ = v7031
	var v7034 int32
	_ = v7034
	var v7036 int32
	_ = v7036
	var v7039 int32
	_ = v7039
	var v7040 int32
	_ = v7040
	var v7041 int32
	_ = v7041
	var v7044 int32
	_ = v7044
	var v7045 int32
	_ = v7045
	var v7048 int32
	_ = v7048
	var v7056 int32
	_ = v7056
	var v7057 int32
	_ = v7057
	var v7060 int32
	_ = v7060
	var v7089 int32
	_ = v7089
	var v7093 int32
	_ = v7093
	var v7094 int32
	_ = v7094
	var v7096 int32
	_ = v7096
	var v7099 int32
	_ = v7099
	var v7104 int32
	_ = v7104
	var v7109 int32
	_ = v7109
	var v7143 int32
	_ = v7143
	var v7147 int32
	_ = v7147
	var v7148 int32
	_ = v7148
	var v7151 int32
	_ = v7151
	var v7154 int32
	_ = v7154
	var v7197 int32
	_ = v7197
	var v7200 int32
	_ = v7200
	var v7243 int32
	_ = v7243
	var v7244 int32
	_ = v7244
	var v7247 int32
	_ = v7247
	var v7287 int32
	_ = v7287
	var v7288 int32
	_ = v7288
	var v7290 int32
	_ = v7290
	var v7293 int32
	_ = v7293
	var v7296 int32
	_ = v7296
	var v7297 int32
	_ = v7297
	var v7299 int32
	_ = v7299
	var v7302 int32
	_ = v7302
	var v7309 int32
	_ = v7309
	var v7311 int32
	_ = v7311
	var v7316 int32
	_ = v7316
	var v7317 int32
	_ = v7317
	var v7322 int32
	_ = v7322
	var v7325 int32
	_ = v7325
	var v7328 int32
	_ = v7328
	var v7332 int32
	_ = v7332
	var v7412 int32
	_ = v7412
	var v7414 int32
	_ = v7414
	var v7415 int32
	_ = v7415
	var v7425 int32
	_ = v7425
	var v7426 int32
	_ = v7426
	var v7459 int32
	_ = v7459
	var v7461 int32
	_ = v7461
	var v7469 int32
	_ = v7469
	var v7472 int32
	_ = v7472
	var v7473 int32
	_ = v7473
	var v7477 int32
	_ = v7477
	var v7478 int32
	_ = v7478
	var v7479 int32
	_ = v7479
	var v7485 int32
	_ = v7485
	var v7488 int32
	_ = v7488
	var v7491 int32
	_ = v7491
	var v7492 int32
	_ = v7492
	var v7494 int32
	_ = v7494
	var v7502 int32
	_ = v7502
	var v7503 int32
	_ = v7503
	var v7505 int32
	_ = v7505
	var v7511 int32
	_ = v7511
	var v7517 int32
	_ = v7517
	var v7519 int32
	_ = v7519
	var v7522 int32
	_ = v7522
	var v7523 int32
	_ = v7523
	var v7524 int32
	_ = v7524
	var v7525 int32
	_ = v7525
	var v7527 int32
	_ = v7527
	var v7528 int32
	_ = v7528
	var v7530 int32
	_ = v7530
	var v7531 int32
	_ = v7531
	var v7533 int32
	_ = v7533
	var v7534 int32
	_ = v7534
	var v7535 int32
	_ = v7535
	var v7536 int32
	_ = v7536
	var v7545 int32
	_ = v7545
	var v7546 int32
	_ = v7546
	var v7579 int32
	_ = v7579
	var v7580 int32
	_ = v7580
	var v7581 int32
	_ = v7581
	var v7582 int32
	_ = v7582
	var v7583 int32
	_ = v7583
	var v7586 int32
	_ = v7586
	var v7589 int32
	_ = v7589
	var v7598 int32
	_ = v7598
	var v7599 int32
	_ = v7599
	var v7632 int32
	_ = v7632
	var v7636 int32
	_ = v7636
	var v7637 int32
	_ = v7637
	var v7638 int32
	_ = v7638
	var v7639 int32
	_ = v7639
	var v7640 int32
	_ = v7640
	var v7641 int32
	_ = v7641
	var v7643 int32
	_ = v7643
	var v7644 int32
	_ = v7644
	var v7653 int32
	_ = v7653
	var v7688 int32
	_ = v7688
	var v7689 int32
	_ = v7689
	var v7690 int32
	_ = v7690
	var v7691 int32
	_ = v7691
	var v7694 int32
	_ = v7694
	var v7698 int32
	_ = v7698
	var v7699 int32
	_ = v7699
	var v7700 int32
	_ = v7700
	var v7703 int32
	_ = v7703
	var v7704 int32
	_ = v7704
	var v7715 int32
	_ = v7715
	var v7748 int32
	_ = v7748
	var v7749 int32
	_ = v7749
	var v7752 int32
	_ = v7752
	var v7753 int32
	_ = v7753
	var v7759 int32
	_ = v7759
	var v7760 int32
	_ = v7760
	var v7803 int32
	_ = v7803
	var v7805 int32
	_ = v7805
	var v7813 int32
	_ = v7813
	var v7820 int32
	_ = v7820
	var v7849 int32
	_ = v7849
	var v7850 int32
	_ = v7850
	var v7854 int32
	_ = v7854
	var v7855 int32
	_ = v7855
	var v7858 int32
	_ = v7858
	var v7859 int32
	_ = v7859
	var v7869 int32
	_ = v7869
	var v7870 int32
	_ = v7870
	var v7903 int32
	_ = v7903
	var v7904 int32
	_ = v7904
	var v7907 int32
	_ = v7907
	var v7911 int32
	_ = v7911
	var v7912 int32
	_ = v7912
	var v7913 int32
	_ = v7913
	var v7915 int32
	_ = v7915
	var v7916 int32
	_ = v7916
	var v7926 int32
	_ = v7926
	var v7959 int32
	_ = v7959
	var v7960 int32
	_ = v7960
	var v7962 int32
	_ = v7962
	var v7963 int32
	_ = v7963
	var v7970 int32
	_ = v7970
	var v8008 int32
	_ = v8008
	var v8009 int32
	_ = v8009
	var v8010 int32
	_ = v8010
	var v8013 int32
	_ = v8013
	var v8014 int32
	_ = v8014
	var v8022 int32
	_ = v8022
	var v8026 int32
	_ = v8026
	var v8027 int32
	_ = v8027
	var v8032 int32
	_ = v8032
	var v8035 int32
	_ = v8035
	var v8037 int32
	_ = v8037
	var v8041 int32
	_ = v8041
	var v8043 int32
	_ = v8043
	var v8050 float64
	_ = v8050
	var v8054 int64
	_ = v8054
	var v8055 int64
	_ = v8055
	var v8057 int32
	_ = v8057
	var v8060 int32
	_ = v8060
	var v8061 int32
	_ = v8061
	var v8064 int32
	_ = v8064
	var v8075 int32
	_ = v8075
	var v8109 int32
	_ = v8109
	var v8110 int32
	_ = v8110
	var v8113 int32
	_ = v8113
	var v8114 int32
	_ = v8114
	var v8120 int32
	_ = v8120
	var v8121 int32
	_ = v8121
	var v8164 int32
	_ = v8164
	var v8165 int32
	_ = v8165
	var v8174 int32
	_ = v8174
	var v8178 int32
	_ = v8178
	var v8210 int32
	_ = v8210
	var v8214 int32
	_ = v8214
	var v8215 int32
	_ = v8215
	var v8219 int32
	_ = v8219
	var v8221 int32
	_ = v8221
	var v8231 int32
	_ = v8231
	var v8232 int32
	_ = v8232
	var v8265 int32
	_ = v8265
	var v8266 int32
	_ = v8266
	var v8269 int32
	_ = v8269
	var v8273 int32
	_ = v8273
	var v8274 int32
	_ = v8274
	var v8275 int32
	_ = v8275
	var v8277 int32
	_ = v8277
	var v8278 int32
	_ = v8278
	var v8288 int32
	_ = v8288
	var v8321 int32
	_ = v8321
	var v8322 int32
	_ = v8322
	var v8324 int32
	_ = v8324
	var v8325 int32
	_ = v8325
	var v8332 int32
	_ = v8332
	var v8368 int32
	_ = v8368
	var v8377 int32
	_ = v8377
	var v8381 int32
	_ = v8381
	var v8382 int32
	_ = v8382
	var v8387 int32
	_ = v8387
	var v8390 int32
	_ = v8390
	var v8396 int32
	_ = v8396
	var v8398 int32
	_ = v8398
	var v8405 float64
	_ = v8405
	var v8409 int64
	_ = v8409
	var v8410 int64
	_ = v8410
	var v8412 int32
	_ = v8412
	var v8416 int32
	_ = v8416
	var v8418 int32
	_ = v8418
	var v8420 int32
	_ = v8420
	var v8421 int32
	_ = v8421
	var v8422 int32
	_ = v8422
	var v8423 int32
	_ = v8423
	var v8424 int32
	_ = v8424
	var v8430 int32
	_ = v8430
	var v8431 int32
	_ = v8431
	var v8432 int32
	_ = v8432
	var v8434 int32
	_ = v8434
	var v8435 int32
	_ = v8435
	var v8437 int32
	_ = v8437
	var v8438 int32
	_ = v8438
	var v8439 int32
	_ = v8439
	var v8442 int32
	_ = v8442
	var v8443 int32
	_ = v8443
	var v8447 int32
	_ = v8447
	var v8451 int32
	_ = v8451
	var v8452 int32
	_ = v8452
	var v8455 int32
	_ = v8455
	var v8491 int32
	_ = v8491
	var v8495 int32
	_ = v8495
	var v8496 int32
	_ = v8496
	var v8499 int32
	_ = v8499
	var v8500 int32
	_ = v8500
	var v8501 int32
	_ = v8501
	var v8502 int32
	_ = v8502
	var v8504 int32
	_ = v8504
	var v8507 int32
	_ = v8507
	var v8508 int32
	_ = v8508
	var v8515 int32
	_ = v8515
	var v8602 int32
	_ = v8602
	var v8607 int32
	_ = v8607
	var v8608 int32
	_ = v8608
	var v8613 int32
	_ = v8613
	var v8614 int32
	_ = v8614
	var v8617 int32
	_ = v8617
	var v8620 int32
	_ = v8620
	var v8635 int32
	_ = v8635
	var v8661 int32
	_ = v8661
	var v8662 int32
	_ = v8662
	var v8665 int32
	_ = v8665
	var v8666 int32
	_ = v8666
	var v8669 int32
	_ = v8669
	var v8670 int32
	_ = v8670
	var v8671 int32
	_ = v8671
	var v8673 int32
	_ = v8673
	var v8678 int32
	_ = v8678
	var v8680 int32
	_ = v8680
	var v8684 int32
	_ = v8684
	var v8685 int32
	_ = v8685
	var v8690 int32
	_ = v8690
	var v8724 int32
	_ = v8724
	var v8728 int32
	_ = v8728
	var v8729 int32
	_ = v8729
	var v8732 int32
	_ = v8732
	var v8733 int32
	_ = v8733
	var v8734 int32
	_ = v8734
	var v8735 int32
	_ = v8735
	var v8737 int32
	_ = v8737
	var v8740 int32
	_ = v8740
	var v8741 int32
	_ = v8741
	var v8750 int32
	_ = v8750
	var v8826 int32
	_ = v8826
	var v8827 int32
	_ = v8827
	var v8828 int32
	_ = v8828
	var v8829 int32
	_ = v8829
	var v8831 int32
	_ = v8831
	var v8832 int32
	_ = v8832
	var v8835 int32
	_ = v8835
	var v8836 int32
	_ = v8836
	var v8839 int32
	_ = v8839
	var v8840 int32
	_ = v8840
	var v8880 int32
	_ = v8880
	var v8884 int32
	_ = v8884
	var v8885 int32
	_ = v8885
	var v8888 int32
	_ = v8888
	var v8890 int32
	_ = v8890
	var v8891 int32
	_ = v8891
	var v8892 int32
	_ = v8892
	var v8896 int32
	_ = v8896
	var v8900 int32
	_ = v8900
	var v8901 int32
	_ = v8901
	var v8902 int32
	_ = v8902
	var v8903 int32
	_ = v8903
	var v8904 int32
	_ = v8904
	var v8906 int32
	_ = v8906
	var v8907 int32
	_ = v8907
	var v8909 int32
	_ = v8909
	var v8951 int32
	_ = v8951
	var v8953 int32
	_ = v8953
	var v8954 int32
	_ = v8954
	var v8959 int32
	_ = v8959
	var v8963 int32
	_ = v8963
	var v8968 int32
	_ = v8968
	var v8969 int32
	_ = v8969
	var v9010 int32
	_ = v9010
	var v9012 int32
	_ = v9012
	var v9013 int32
	_ = v9013
	var v9016 int32
	_ = v9016
	var v9017 int32
	_ = v9017
	var v9020 int32
	_ = v9020
	var v9021 int32
	_ = v9021
	var v9061 int32
	_ = v9061
	var v9065 int32
	_ = v9065
	var v9066 int32
	_ = v9066
	var v9069 int32
	_ = v9069
	var v9071 int32
	_ = v9071
	var v9072 int32
	_ = v9072
	var v9073 int32
	_ = v9073
	var v9077 int32
	_ = v9077
	var v9081 int32
	_ = v9081
	var v9082 int32
	_ = v9082
	var v9083 int32
	_ = v9083
	var v9084 int32
	_ = v9084
	var v9085 int32
	_ = v9085
	var v9087 int32
	_ = v9087
	var v9088 int32
	_ = v9088
	var v9090 int32
	_ = v9090
	var v9131 int32
	_ = v9131
	var v9134 int32
	_ = v9134
	var v9138 int32
	_ = v9138
	var v9140 int32
	_ = v9140
	var v9179 int32
	_ = v9179
	var v9183 int32
	_ = v9183
	var v9184 int32
	_ = v9184
	var v9185 int32
	_ = v9185
	var v9187 int32
	_ = v9187
	var v9190 int32
	_ = v9190
	var v9193 int32
	_ = v9193
	var v9195 int32
	_ = v9195
	var v9196 int32
	_ = v9196
	var v9197 int32
	_ = v9197
	var v9201 int32
	_ = v9201
	var v9205 int32
	_ = v9205
	var v9206 int32
	_ = v9206
	var v9207 int32
	_ = v9207
	var v9211 int32
	_ = v9211
	var v9215 int32
	_ = v9215
	var v9216 int32
	_ = v9216
	var v9218 int32
	_ = v9218
	var v9219 int32
	_ = v9219
	var v9220 int32
	_ = v9220
	var v9221 int32
	_ = v9221
	var v9222 int32
	_ = v9222
	var v9223 int32
	_ = v9223
	var v9225 int32
	_ = v9225
	var v9228 int32
	_ = v9228
	var v9229 int32
	_ = v9229
	var v9235 int32
	_ = v9235
	var v9236 int32
	_ = v9236
	var v9238 int32
	_ = v9238
	var v9239 int32
	_ = v9239
	var v9240 int32
	_ = v9240
	var v9248 int32
	_ = v9248
	var v9249 int32
	_ = v9249
	var v9250 int32
	_ = v9250
	var v9254 int32
	_ = v9254
	var v9258 int32
	_ = v9258
	var v9259 int32
	_ = v9259
	var v9261 int32
	_ = v9261
	var v9262 int32
	_ = v9262
	var v9263 int32
	_ = v9263
	var v9264 int32
	_ = v9264
	var v9265 int32
	_ = v9265
	var v9267 int32
	_ = v9267
	var v9270 int32
	_ = v9270
	var v9274 int32
	_ = v9274
	var v9276 int32
	_ = v9276
	var v9277 int32
	_ = v9277
	var v9278 int32
	_ = v9278
	var v9284 int32
	_ = v9284
	var v9285 int32
	_ = v9285
	var v9286 int32
	_ = v9286
	var v9290 int32
	_ = v9290
	var v9294 int32
	_ = v9294
	var v9295 int32
	_ = v9295
	var v9297 int32
	_ = v9297
	var v9298 int32
	_ = v9298
	var v9299 int32
	_ = v9299
	var v9300 int32
	_ = v9300
	var v9301 int32
	_ = v9301
	var v9305 int32
	_ = v9305
	var v9306 int32
	_ = v9306
	var v9308 int32
	_ = v9308
	var v9349 int32
	_ = v9349
	var v9352 int32
	_ = v9352
	var v9355 int32
	_ = v9355
	var v9359 int32
	_ = v9359
	var v9360 int32
	_ = v9360
	var v9363 int32
	_ = v9363
	var v9367 int32
	_ = v9367
	var v9368 int32
	_ = v9368
	var v9408 int32
	_ = v9408
	var v9412 int32
	_ = v9412
	var v9413 int32
	_ = v9413
	var v9416 int32
	_ = v9416
	var v9418 int32
	_ = v9418
	var v9419 int32
	_ = v9419
	var v9420 int32
	_ = v9420
	var v9424 int32
	_ = v9424
	var v9428 int32
	_ = v9428
	var v9429 int32
	_ = v9429
	var v9430 int32
	_ = v9430
	var v9431 int32
	_ = v9431
	var v9432 int32
	_ = v9432
	var v9434 int32
	_ = v9434
	var v9435 int32
	_ = v9435
	var v9437 int32
	_ = v9437
	var v9479 int32
	_ = v9479
	var v9480 int32
	_ = v9480
	var v9524 int32
	_ = v9524
	var v9528 int32
	_ = v9528
	var v9531 int32
	_ = v9531
	var v9533 int32
	_ = v9533
	var v9534 int32
	_ = v9534
	var v9536 int32
	_ = v9536
	var v9537 int32
	_ = v9537
	var v9539 int32
	_ = v9539
	var v9542 int32
	_ = v9542
	var v9543 int32
	_ = v9543
	var v9544 int32
	_ = v9544
	var v9546 int32
	_ = v9546
	var v9548 int32
	_ = v9548
	var v9549 int32
	_ = v9549
	var v9557 int32
	_ = v9557
	var v9558 int32
	_ = v9558
	var v9560 int32
	_ = v9560
	var v9561 int32
	_ = v9561
	var v9562 int32
	_ = v9562
	var v9565 int32
	_ = v9565
	var v9566 int32
	_ = v9566
	var v9569 int32
	_ = v9569
	var v9570 int32
	_ = v9570
	var v9573 int32
	_ = v9573
	var v9614 int32
	_ = v9614
	var v9615 int32
	_ = v9615
	var v9616 int32
	_ = v9616
	var v9619 int32
	_ = v9619
	var v9620 int32
	_ = v9620
	var v9624 int32
	_ = v9624
	var v9627 int32
	_ = v9627
	var v9631 int32
	_ = v9631
	var v9632 int32
	_ = v9632
	var v9633 int32
	_ = v9633
	var v9634 int32
	_ = v9634
	var v9635 int32
	_ = v9635
	var v9642 int32
	_ = v9642
	var v9647 int32
	_ = v9647
	var v9648 int32
	_ = v9648
	var v9651 int32
	_ = v9651
	var v9653 int32
	_ = v9653
	var v9673 int32
	_ = v9673
	var v9697 int32
	_ = v9697
	var v9701 int32
	_ = v9701
	var v9702 int32
	_ = v9702
	var v9703 int32
	_ = v9703
	var v9704 int32
	_ = v9704
	var v9711 int32
	_ = v9711
	var v9716 int32
	_ = v9716
	var v9717 int32
	_ = v9717
	var v9720 int32
	_ = v9720
	var v9723 int32
	_ = v9723
	var v9724 int32
	_ = v9724
	var v9735 int32
	_ = v9735
	var v9767 int32
	_ = v9767
	var v9770 int32
	_ = v9770
	var v9782 int32
	_ = v9782
	var v9815 int32
	_ = v9815
	var v9819 int32
	_ = v9819
	var v9821 int32
	_ = v9821
	var v9822 int32
	_ = v9822
	var v9823 int32
	_ = v9823
	var v9824 int32
	_ = v9824
	var v9827 int32
	_ = v9827
	var v9828 int32
	_ = v9828
	var v9829 int32
	_ = v9829
	var v9830 int32
	_ = v9830
	var v9833 int32
	_ = v9833
	var v9834 int32
	_ = v9834
	var v9836 int32
	_ = v9836
	var v9837 int32
	_ = v9837
	var v9838 int32
	_ = v9838
	var v9839 int32
	_ = v9839
	var v9842 int32
	_ = v9842
	var v9843 int32
	_ = v9843
	var v9844 int32
	_ = v9844
	var v9845 int32
	_ = v9845
	var v9848 int32
	_ = v9848
	var v9849 int32
	_ = v9849
	var v9853 int32
	_ = v9853
	var v9857 int32
	_ = v9857
	var v9860 int32
	_ = v9860
	var v9870 int32
	_ = v9870
	var v9903 int32
	_ = v9903
	var v9907 int32
	_ = v9907
	var v9910 int32
	_ = v9910
	var v9911 int32
	_ = v9911
	var v9913 int32
	_ = v9913
	var v9914 int32
	_ = v9914
	var v9918 int32
	_ = v9918
	var v9927 int32
	_ = v9927
	var v9928 int32
	_ = v9928
	var v9960 int32
	_ = v9960
	var v9964 int32
	_ = v9964
	var v9965 int32
	_ = v9965
	var v9966 int32
	_ = v9966
	var v9967 int32
	_ = v9967
	var v9968 int32
	_ = v9968
	var v9969 int32
	_ = v9969
	var v9973 int32
	_ = v9973
	var v9974 int32
	_ = v9974
	var v9981 int32
	_ = v9981
	var v9982 int32
	_ = v9982
	var v10026 int32
	_ = v10026
	var v10027 int32
	_ = v10027
	var v10029 int32
	_ = v10029
	var v10031 int32
	_ = v10031
	var v10041 int32
	_ = v10041
	var v10043 int32
	_ = v10043
	var v10075 int32
	_ = v10075
	var v10076 int32
	_ = v10076
	var v10077 int32
	_ = v10077
	var v10080 int32
	_ = v10080
	var v10081 int32
	_ = v10081
	var v10085 int32
	_ = v10085
	var v10086 int32
	_ = v10086
	var v10090 int32
	_ = v10090
	var v10099 int32
	_ = v10099
	var v10109 int32
	_ = v10109
	var v10133 int32
	_ = v10133
	var v10136 int32
	_ = v10136
	var v10137 int32
	_ = v10137
	var v10138 int32
	_ = v10138
	var v10139 int32
	_ = v10139
	var v10142 int32
	_ = v10142
	var v10143 int32
	_ = v10143
	var v10144 int32
	_ = v10144
	var v10145 int32
	_ = v10145
	var v10149 int32
	_ = v10149
	var v10151 int32
	_ = v10151
	var v10152 int32
	_ = v10152
	var v10171 int32
	_ = v10171
	var v10196 int32
	_ = v10196
	var v10197 int32
	_ = v10197
	var v10198 int32
	_ = v10198
	var v10203 int32
	_ = v10203
	var v10211 int32
	_ = v10211
	var v10246 int32
	_ = v10246
	var v10289 int32
	_ = v10289
	var v10290 int32
	_ = v10290
	var v10294 int32
	_ = v10294
	var v10295 int32
	_ = v10295
	var v10296 int32
	_ = v10296
	var v10299 int32
	_ = v10299
	var v10300 int32
	_ = v10300
	var v10301 int32
	_ = v10301
	var v10304 int32
	_ = v10304
	var v10313 int32
	_ = v10313
	var v10320 int32
	_ = v10320
	var v10323 int32
	_ = v10323
	var v10347 int32
	_ = v10347
	var v10348 int32
	_ = v10348
	var v10349 int32
	_ = v10349
	var v10352 int32
	_ = v10352
	var v10353 int32
	_ = v10353
	var v10357 int32
	_ = v10357
	var v10360 int32
	_ = v10360
	var v10362 int32
	_ = v10362
	var v10363 int32
	_ = v10363
	var v10364 int32
	_ = v10364
	var v10365 int32
	_ = v10365
	var v10366 int32
	_ = v10366
	var v10367 int32
	_ = v10367
	var v10371 int32
	_ = v10371
	var v10373 int32
	_ = v10373
	var v10374 int32
	_ = v10374
	var v10376 int32
	_ = v10376
	var v10381 int32
	_ = v10381
	var v10382 int32
	_ = v10382
	var v10392 int32
	_ = v10392
	var v10405 int32
	_ = v10405
	var v10429 int32
	_ = v10429
	var v10430 int32
	_ = v10430
	var v10431 int32
	_ = v10431
	var v10433 int32
	_ = v10433
	var v10436 int32
	_ = v10436
	var v10441 int32
	_ = v10441
	var v10452 int32
	_ = v10452
	var v10454 int32
	_ = v10454
	var v10459 int32
	_ = v10459
	var v10486 int32
	_ = v10486
	var v10490 int32
	_ = v10490
	var v10491 int32
	_ = v10491
	var v10501 int32
	_ = v10501
	var v10502 int32
	_ = v10502
	var v10504 int32
	_ = v10504
	var v10505 int32
	_ = v10505
	var v10508 int32
	_ = v10508
	var v10517 int32
	_ = v10517
	var v10549 int32
	_ = v10549
	var v10558 int32
	_ = v10558
	var v10595 int32
	_ = v10595
	var v10596 int32
	_ = v10596
	var v10598 int32
	_ = v10598
	var v10601 int32
	_ = v10601
	var v10604 int32
	_ = v10604
	var v10607 int32
	_ = v10607
	var v10608 int32
	_ = v10608
	var v10611 int32
	_ = v10611
	var v10612 int32
	_ = v10612
	var v10615 int32
	_ = v10615
	var v10622 int32
	_ = v10622
	var v10623 int32
	_ = v10623
	var v10628 int32
	_ = v10628
	var v10637 int32
	_ = v10637
	var v10638 int32
	_ = v10638
	var v10640 int32
	_ = v10640
	var v10641 int32
	_ = v10641
	var v10653 int32
	_ = v10653
	var v10687 int32
	_ = v10687
	var v10688 int32
	_ = v10688
	var v10690 int32
	_ = v10690
	var v10698 int32
	_ = v10698
	var v10700 int32
	_ = v10700
	var v10733 int32
	_ = v10733
	var v10735 int32
	_ = v10735
	var v10741 int32
	_ = v10741
	var v10751 int32
	_ = v10751
	var v10784 int32
	_ = v10784
	var v10805 int32
	_ = v10805
	var v10807 int32
	_ = v10807
	var v10826 int32
	_ = v10826
	var v10829 int32
	_ = v10829
	var v10831 float64
	_ = v10831
	var v10832 int32
	_ = v10832
	var v10834 int32
	_ = v10834
	var v10836 int32
	_ = v10836
	var v10837 int32
	_ = v10837
	var v10840 int32
	_ = v10840
	var v10841 int32
	_ = v10841
	var v10842 int32
	_ = v10842
	var v10845 int32
	_ = v10845
	var v10846 int32
	_ = v10846
	var v10849 int32
	_ = v10849
	var v10890 int32
	_ = v10890
	var v10891 int32
	_ = v10891
	var v10896 int32
	_ = v10896
	var v10899 int32
	_ = v10899
	var v10903 int32
	_ = v10903
	var v10906 int32
	_ = v10906
	var v10909 int32
	_ = v10909
	var v10910 int32
	_ = v10910
	var v10911 int32
	_ = v10911
	var v10912 int32
	_ = v10912
	var v10918 int32
	_ = v10918
	var v10919 int32
	_ = v10919
	var v10920 int32
	_ = v10920
	var v10921 int32
	_ = v10921
	var v10924 int32
	_ = v10924
	var v10927 int32
	_ = v10927
	var v10931 int32
	_ = v10931
	var v10933 int32
	_ = v10933
	var v10950 int32
	_ = v10950
	var v10977 int32
	_ = v10977
	var v10978 int32
	_ = v10978
	var v10982 int32
	_ = v10982
	var v10983 int32
	_ = v10983
	var v10984 int32
	_ = v10984
	var v10985 int32
	_ = v10985
	var v10986 int32
	_ = v10986
	var v10990 int32
	_ = v10990
	var v10994 int32
	_ = v10994
	var v10996 int32
	_ = v10996
	var v10997 int32
	_ = v10997
	var v10999 int32
	_ = v10999
	var v11000 int32
	_ = v11000
	var v11001 int32
	_ = v11001
	var v11004 int32
	_ = v11004
	var v11005 int32
	_ = v11005
	var v11007 int32
	_ = v11007
	var v11009 int32
	_ = v11009
	var v11012 int32
	_ = v11012
	var v11013 int32
	_ = v11013
	var v11014 int32
	_ = v11014
	var v11015 int32
	_ = v11015
	var v11016 int32
	_ = v11016
	var v11017 int32
	_ = v11017
	var v11018 int32
	_ = v11018
	var v11019 int32
	_ = v11019
	var v11020 int32
	_ = v11020
	var v11021 int32
	_ = v11021
	var v11022 int32
	_ = v11022
	var v11024 int32
	_ = v11024
	var v11025 int32
	_ = v11025
	var v11028 int32
	_ = v11028
	var v11031 int32
	_ = v11031
	var v11032 int64
	_ = v11032
	var v11039 int32
	_ = v11039
	var v11040 int32
	_ = v11040
	var v11041 int32
	_ = v11041
	var v11043 int32
	_ = v11043
	var v11045 int32
	_ = v11045
	var v11046 int32
	_ = v11046
	var v11057 int32
	_ = v11057
	var v11130 int32
	_ = v11130
	var v11133 int32
	_ = v11133
	var v11136 int32
	_ = v11136
	var v11144 int32
	_ = v11144
	var v11180 int32
	_ = v11180
	var v11184 int32
	_ = v11184
	var v11185 int32
	_ = v11185
	var v11188 int32
	_ = v11188
	var v11189 int32
	_ = v11189
	var v11192 int32
	_ = v11192
	var v11193 int32
	_ = v11193
	var v11194 int32
	_ = v11194
	var v11195 int32
	_ = v11195
	var v11198 int32
	_ = v11198
	var v11199 int32
	_ = v11199
	var v11202 int32
	_ = v11202
	var v11203 int32
	_ = v11203
	var v11208 int32
	_ = v11208
	var v11209 int32
	_ = v11209
	var v11252 int32
	_ = v11252
	var v11261 int32
	_ = v11261
	var v11297 int32
	_ = v11297
	var v11301 int32
	_ = v11301
	var v11302 int32
	_ = v11302
	var v11303 int32
	_ = v11303
	var v11304 int32
	_ = v11304
	var v11306 int32
	_ = v11306
	var v11307 int32
	_ = v11307
	var v11308 int32
	_ = v11308
	var v11309 int32
	_ = v11309
	var v11310 int32
	_ = v11310
	var v11313 int32
	_ = v11313
	var v11314 int32
	_ = v11314
	var v11359 int32
	_ = v11359
	var v11360 int32
	_ = v11360
	var v11361 int32
	_ = v11361
	var v11362 int32
	_ = v11362
	var v11363 int32
	_ = v11363
	var v11364 int32
	_ = v11364
	var v11365 int32
	_ = v11365
	var v11366 int32
	_ = v11366
	var v11367 int32
	_ = v11367
	var v11368 int32
	_ = v11368
	var v11370 int32
	_ = v11370
	var v11373 int32
	_ = v11373
	var v11374 int32
	_ = v11374
	var v11377 int32
	_ = v11377
	var v11383 int32
	_ = v11383
	var v11394 int32
	_ = v11394
	var v11395 int32
	_ = v11395
	var v11399 int32
	_ = v11399
	var v11403 float64
	_ = v11403
	var v11407 int32
	_ = v11407
	var v11408 int32
	_ = v11408
	var v11440 int32
	_ = v11440
	var v11444 int32
	_ = v11444
	var v11445 float64
	_ = v11445
	var v11446 int32
	_ = v11446
	var v11447 int32
	_ = v11447
	var v11448 int32
	_ = v11448
	var v11451 int32
	_ = v11451
	var v11454 int32
	_ = v11454
	var v11455 float64
	_ = v11455
	var v11456 int32
	_ = v11456
	var v11458 int32
	_ = v11458
	var v11459 int32
	_ = v11459
	var v11465 float64
	_ = v11465
	var v11469 int32
	_ = v11469
	var v11470 int32
	_ = v11470
	var v11503 float64
	_ = v11503
	var v11506 float64
	_ = v11506
	var v11508 float64
	_ = v11508
	var v11511 float64
	_ = v11511
	var v11515 int32
	_ = v11515
	var v11516 float64
	_ = v11516
	var v11517 float64
	_ = v11517
	var v11520 float64
	_ = v11520
	var v11521 float64
	_ = v11521
	var v11525 int32
	_ = v11525
	var v11529 int32
	_ = v11529
	var v11530 int32
	_ = v11530
	var v11531 int32
	_ = v11531
	var v11532 int32
	_ = v11532
	var v11533 int32
	_ = v11533
	var v11535 int32
	_ = v11535
	var v11542 int32
	_ = v11542
	var v11590 int32
	_ = v11590
	var v11591 int32
	_ = v11591
	var v11595 int32
	_ = v11595
	var v11600 int32
	_ = v11600
	var v11642 float64
	_ = v11642
	var v11643 int32
	_ = v11643
	var v11644 int32
	_ = v11644
	var v11645 int32
	_ = v11645
	var v11646 int32
	_ = v11646
	var v11647 int32
	_ = v11647
	var v11648 int32
	_ = v11648
	var v11650 int32
	_ = v11650
	var v11651 float64
	_ = v11651
	var v11652 float64
	_ = v11652
	var v11660 int32
	_ = v11660
	var v11661 int32
	_ = v11661
	var v11662 int32
	_ = v11662
	var v11663 int32
	_ = v11663
	var v11664 int32
	_ = v11664
	var v11665 int32
	_ = v11665
	var v11666 int32
	_ = v11666
	var v11667 int32
	_ = v11667
	var v11668 int32
	_ = v11668
	var v11669 int32
	_ = v11669
	var v11670 int32
	_ = v11670
	var v11671 int32
	_ = v11671
	var v11672 int32
	_ = v11672
	var v11673 int32
	_ = v11673
	var v11675 int32
	_ = v11675
	var v11676 int32
	_ = v11676
	var v11677 int32
	_ = v11677
	var v11678 int32
	_ = v11678
	var v11679 int32
	_ = v11679
	var v11680 int32
	_ = v11680
	var v11683 int32
	_ = v11683
	var v11686 int32
	_ = v11686
	var v11698 int32
	_ = v11698
	var v11699 int32
	_ = v11699
	var v11700 int32
	_ = v11700
	var v11704 int32
	_ = v11704
	var v11712 int32
	_ = v11712
	var v11733 int32
	_ = v11733
	var v11734 int32
	_ = v11734
	var v11736 int32
	_ = v11736
	var v11737 int32
	_ = v11737
	var v11739 int32
	_ = v11739
	var v11740 int32
	_ = v11740
	var v11743 int32
	_ = v11743
	var v11744 int32
	_ = v11744
	var v11747 int32
	_ = v11747
	var v11751 int32
	_ = v11751
	var v11752 int32
	_ = v11752
	var v11753 int32
	_ = v11753
	var v11760 int32
	_ = v11760
	var v11761 float64
	_ = v11761
	var v11763 float64
	_ = v11763
	var v11769 int32
	_ = v11769
	var v11777 int32
	_ = v11777
	var v11780 int32
	_ = v11780
	var v11781 int32
	_ = v11781
	var v11782 int32
	_ = v11782
	var v11783 int32
	_ = v11783
	var v11784 int32
	_ = v11784
	var v11785 int32
	_ = v11785
	var v11787 int32
	_ = v11787
	var v11788 int32
	_ = v11788
	var v11790 int32
	_ = v11790
	var v11792 int32
	_ = v11792
	var v11800 int32
	_ = v11800
	var v11801 float64
	_ = v11801
	var v11806 int32
	_ = v11806
	var v11807 int32
	_ = v11807
	var v11808 int32
	_ = v11808
	var v11812 int32
	_ = v11812
	var v11813 int32
	_ = v11813
	var v11817 int32
	_ = v11817
	var v11824 int32
	_ = v11824
	var v11859 int32
	_ = v11859
	var v11860 int32
	_ = v11860
	var v11862 int32
	_ = v11862
	var v11864 int32
	_ = v11864
	var v11872 int32
	_ = v11872
	var v11875 int32
	_ = v11875
	var v11876 int32
	_ = v11876
	var v11877 int32
	_ = v11877
	var v11879 int32
	_ = v11879
	var v11881 int32
	_ = v11881
	var v11883 int32
	_ = v11883
	var v11884 int32
	_ = v11884
	var v11887 int32
	_ = v11887
	var v11888 int32
	_ = v11888
	var v11890 int32
	_ = v11890
	var v11932 int32
	_ = v11932
	var v11933 int32
	_ = v11933
	var v11935 int32
	_ = v11935
	var v11937 int32
	_ = v11937
	var v11939 int32
	_ = v11939
	var v11943 float64
	_ = v11943
	var v11944 int32
	_ = v11944
	var v11945 int32
	_ = v11945
	var v11950 float64
	_ = v11950
	var v11987 int32
	_ = v11987
	var v11988 int32
	_ = v11988
	var v11989 int32
	_ = v11989
	var v11990 int32
	_ = v11990
	var v11993 int32
	_ = v11993
	var v11995 float64
	_ = v11995
	var v12009 int32
	_ = v12009
	var v12032 int32
	_ = v12032
	var v12033 int32
	_ = v12033
	var v12044 int32
	_ = v12044
	var v12051 int32
	_ = v12051
	var v12078 int32
	_ = v12078
	var v12082 int32
	_ = v12082
	var v12083 int32
	_ = v12083
	var v12086 int32
	_ = v12086
	var v12087 int32
	_ = v12087
	var v12097 int32
	_ = v12097
	var v12098 int32
	_ = v12098
	var v12131 int32
	_ = v12131
	var v12135 int32
	_ = v12135
	var v12136 int32
	_ = v12136
	var v12137 int32
	_ = v12137
	var v12138 int32
	_ = v12138
	var v12140 int32
	_ = v12140
	var v12141 int32
	_ = v12141
	var v12150 int32
	_ = v12150
	var v12184 int32
	_ = v12184
	var v12187 int32
	_ = v12187
	var v12188 int32
	_ = v12188
	var v12198 int32
	_ = v12198
	var v12199 int32
	_ = v12199
	var v12232 int32
	_ = v12232
	var v12236 int32
	_ = v12236
	var v12237 int32
	_ = v12237
	var v12238 int32
	_ = v12238
	var v12239 int32
	_ = v12239
	var v12241 int32
	_ = v12241
	var v12242 int32
	_ = v12242
	var v12251 int32
	_ = v12251
	var v12286 int32
	_ = v12286
	var v12287 int32
	_ = v12287
	var v12296 int32
	_ = v12296
	var v12330 int32
	_ = v12330
	var v12333 int32
	_ = v12333
	var v12334 int32
	_ = v12334
	var v12344 int32
	_ = v12344
	var v12345 int32
	_ = v12345
	var v12378 int32
	_ = v12378
	var v12382 int32
	_ = v12382
	var v12383 int32
	_ = v12383
	var v12384 int32
	_ = v12384
	var v12385 int32
	_ = v12385
	var v12387 int32
	_ = v12387
	var v12388 int32
	_ = v12388
	var v12397 int32
	_ = v12397
	var v12431 int32
	_ = v12431
	var v12432 int32
	_ = v12432
	var v12433 int32
	_ = v12433
	var v12437 int32
	_ = v12437
	var v12438 int32
	_ = v12438
	var v12450 int32
	_ = v12450
	var v12456 int32
	_ = v12456
	var v12484 int32
	_ = v12484
	var v12485 int32
	_ = v12485
	var v12487 int32
	_ = v12487
	var v12488 int32
	_ = v12488
	var v12492 int32
	_ = v12492
	var v12495 int32
	_ = v12495
	var v12496 int32
	_ = v12496
	var v12500 int32
	_ = v12500
	var v12502 int32
	_ = v12502
	var v12503 int32
	_ = v12503
	var v12505 int32
	_ = v12505
	var v12507 int32
	_ = v12507
	var v12508 int32
	_ = v12508
	var v12524 int32
	_ = v12524
	var v12552 int32
	_ = v12552
	var v12553 int32
	_ = v12553
	var v12555 int32
	_ = v12555
	var v12557 int32
	_ = v12557
	var v12559 int32
	_ = v12559
	var v12560 int32
	_ = v12560
	var v12561 int32
	_ = v12561
	var v12562 int32
	_ = v12562
	var v12563 int32
	_ = v12563
	var v12564 int32
	_ = v12564
	var v12566 int32
	_ = v12566
	var v12578 int32
	_ = v12578
	var v12606 int32
	_ = v12606
	var v12607 int32
	_ = v12607
	var v12608 int32
	_ = v12608
	var v12610 int32
	_ = v12610
	var v12615 int32
	_ = v12615
	var v12616 int32
	_ = v12616
	var v12617 int32
	_ = v12617
	var v12618 int32
	_ = v12618
	var v12622 int32
	_ = v12622
	var v12623 int32
	_ = v12623
	var v12627 int32
	_ = v12627
	var v12634 int32
	_ = v12634
	var v12669 int32
	_ = v12669
	var v12670 int32
	_ = v12670
	var v12672 int32
	_ = v12672
	var v12673 int32
	_ = v12673
	var v12677 int32
	_ = v12677
	var v12680 int32
	_ = v12680
	var v12686 int32
	_ = v12686
	var v12689 int32
	_ = v12689
	var v12692 int32
	_ = v12692
	var v12693 int32
	_ = v12693
	var v12696 int32
	_ = v12696
	var v12703 int32
	_ = v12703
	var v12704 int32
	_ = v12704
	var v12707 int32
	_ = v12707
	var v12718 int32
	_ = v12718
	var v12722 int32
	_ = v12722
	var v12725 int32
	_ = v12725
	var v12728 int32
	_ = v12728
	var v12729 int32
	_ = v12729
	var v12730 int32
	_ = v12730
	var v12732 int32
	_ = v12732
	var v12733 int32
	_ = v12733
	var v12734 int32
	_ = v12734
	var v12736 int32
	_ = v12736
	var v12739 int32
	_ = v12739
	var v12740 int32
	_ = v12740
	var v12741 int32
	_ = v12741
	var v12746 int32
	_ = v12746
	var v12747 int32
	_ = v12747
	var v12749 int32
	_ = v12749
	var v12790 int32
	_ = v12790
	var v12791 int32
	_ = v12791
	var v12792 int32
	_ = v12792
	var v12793 int32
	_ = v12793
	var v12795 int32
	_ = v12795
	var v12796 int32
	_ = v12796
	var v12797 int32
	_ = v12797
	var v12798 int32
	_ = v12798
	var v12801 int32
	_ = v12801
	var v12804 int32
	_ = v12804
	var v12805 int32
	_ = v12805
	var v12806 int32
	_ = v12806
	var v12808 int32
	_ = v12808
	var v12809 int32
	_ = v12809
	var v12810 int32
	_ = v12810
	var v12812 int32
	_ = v12812
	var v12814 int32
	_ = v12814
	var v12816 int32
	_ = v12816
	var v12817 int32
	_ = v12817
	var v12818 int32
	_ = v12818
	var v12819 int32
	_ = v12819
	var v12820 int32
	_ = v12820
	var v12821 int32
	_ = v12821
	var v12829 int32
	_ = v12829
	var v12836 int32
	_ = v12836
	var v12863 int32
	_ = v12863
	var v12864 int32
	_ = v12864
	var v12872 int32
	_ = v12872
	var v12873 int32
	_ = v12873
	var v12874 int32
	_ = v12874
	var v12875 int32
	_ = v12875
	var v12881 int32
	_ = v12881
	var v12882 int32
	_ = v12882
	var v12883 int32
	_ = v12883
	var v12884 int32
	_ = v12884
	var v12891 int32
	_ = v12891
	var v12892 int32
	_ = v12892
	var v12893 int32
	_ = v12893
	var v12894 int32
	_ = v12894
	var v12901 int32
	_ = v12901
	var v12902 int32
	_ = v12902
	var v12903 int32
	_ = v12903
	var v12904 int32
	_ = v12904
	var v12905 int32
	_ = v12905
	var v12923 int32
	_ = v12923
	var v12924 int32
	_ = v12924
	var v12930 int32
	_ = v12930
	var v12931 int32
	_ = v12931
	var v12932 int32
	_ = v12932
	var v12933 int32
	_ = v12933
	var v12934 int32
	_ = v12934
	var v12936 int32
	_ = v12936
	var v12939 int32
	_ = v12939
	var v12940 int32
	_ = v12940
	var v12941 int32
	_ = v12941
	var v12942 int32
	_ = v12942
	var v12943 int32
	_ = v12943
	var v12944 int32
	_ = v12944
	var v12945 int32
	_ = v12945
	var v12946 int32
	_ = v12946
	var v12947 int32
	_ = v12947
	var v12950 int32
	_ = v12950
	var v12951 int32
	_ = v12951
	var v12952 int32
	_ = v12952
	var v12954 int32
	_ = v12954
	var v12963 int32
	_ = v12963
	var v12964 int64
	_ = v12964
	var v12978 int32
	_ = v12978
	var v12979 int32
	_ = v12979
	var v12980 int32
	_ = v12980
	var v12987 int32
	_ = v12987
	var v12993 int32
	_ = v12993
	var v12994 int32
	_ = v12994
	var v12995 int32
	_ = v12995
	var v13000 int32
	_ = v13000
	var v13001 int32
	_ = v13001
	var v13002 int32
	_ = v13002
	var v13004 int32
	_ = v13004
	var v13008 int32
	_ = v13008
	var v13009 int32
	_ = v13009
	var v13012 int32
	_ = v13012
	var v13014 int32
	_ = v13014
	var v13016 int32
	_ = v13016
	var v13018 int32
	_ = v13018
	var v13020 int32
	_ = v13020
	var v13022 int32
	_ = v13022
	var v13023 int32
	_ = v13023
	var v13026 int32
	_ = v13026
	var v13029 int32
	_ = v13029
	var v13030 int32
	_ = v13030
	var v13031 int32
	_ = v13031
	var v13034 int32
	_ = v13034
	var v13045 int32
	_ = v13045
	var v13055 int32
	_ = v13055
	var v13081 int32
	_ = v13081
	var v13082 int32
	_ = v13082
	var v13083 int32
	_ = v13083
	var v13084 int32
	_ = v13084
	var v13085 int32
	_ = v13085
	var v13086 int32
	_ = v13086
	var v13089 int32
	_ = v13089
	var v13095 int32
	_ = v13095
	var v13096 int32
	_ = v13096
	var v13098 int32
	_ = v13098
	var v13100 int32
	_ = v13100
	var v13101 int32
	_ = v13101
	var v13102 int32
	_ = v13102
	var v13103 int32
	_ = v13103
	var v13105 int32
	_ = v13105
	var v13106 int32
	_ = v13106
	var v13107 int32
	_ = v13107
	var v13108 int32
	_ = v13108
	var v13117 int32
	_ = v13117
	var v13120 int32
	_ = v13120
	var v13123 int32
	_ = v13123
	var v13124 int32
	_ = v13124
	var v13126 int32
	_ = v13126
	var v13134 int32
	_ = v13134
	var v13135 int32
	_ = v13135
	var v13136 int32
	_ = v13136
	var v13137 int32
	_ = v13137
	var v13141 int32
	_ = v13141
	var v13145 int32
	_ = v13145
	var v13153 int32
	_ = v13153
	var v13158 int32
	_ = v13158
	var v13159 int32
	_ = v13159
	var v13162 int32
	_ = v13162
	var v13163 int32
	_ = v13163
	var v13164 int32
	_ = v13164
	var v13165 int32
	_ = v13165
	var v13172 int32
	_ = v13172
	var v13175 int32
	_ = v13175
	var v13178 int32
	_ = v13178
	var v13179 int32
	_ = v13179
	var v13183 int32
	_ = v13183
	var v13187 int32
	_ = v13187
	var v13188 int32
	_ = v13188
	var v13192 int32
	_ = v13192
	var v13196 int32
	_ = v13196
	var v13202 int32
	_ = v13202
	var v13207 int32
	_ = v13207
	var v13208 int32
	_ = v13208
	var v13209 int32
	_ = v13209
	var v13212 int32
	_ = v13212
	var v13215 int32
	_ = v13215
	var v13216 int32
	_ = v13216
	var v13219 int32
	_ = v13219
	var v13220 int32
	_ = v13220
	var v13221 int32
	_ = v13221
	var v13224 int32
	_ = v13224
	var v13226 int32
	_ = v13226
	var v13228 int32
	_ = v13228
	var v13230 int32
	_ = v13230
	var v13232 int32
	_ = v13232
	var v13235 int32
	_ = v13235
	var v13239 int32
	_ = v13239
	var v13248 int32
	_ = v13248
	var v13291 int32
	_ = v13291
	var v13292 int32
	_ = v13292
	var v13295 int32
	_ = v13295
	var v13296 int32
	_ = v13296
	var v13298 int32
	_ = v13298
	var v13304 int32
	_ = v13304
	var v13344 int32
	_ = v13344
	var v13345 int32
	_ = v13345
	var v13346 int32
	_ = v13346
	var v13350 int32
	_ = v13350
	var v13351 int32
	_ = v13351
	var v13354 int32
	_ = v13354
	var v13356 int32
	_ = v13356
	var v13358 int32
	_ = v13358
	var v13360 int32
	_ = v13360
	var v13362 int32
	_ = v13362
	var v13364 int32
	_ = v13364
	var v13367 int32
	_ = v13367
	var v13397 int32
	_ = v13397
	var v13411 int32
	_ = v13411
	var v13415 int32
	_ = v13415
	var v13416 int32
	_ = v13416
	var v13418 int32
	_ = v13418
	var v13419 int32
	_ = v13419
	var v13421 int32
	_ = v13421
	var v13439 int32
	_ = v13439
	var v13442 int32
	_ = v13442
	var v13443 int32
	_ = v13443
	var v13446 int32
	_ = v13446
	var v13447 int32
	_ = v13447
	var v13452 int32
	_ = v13452
	var v13459 int32
	_ = v13459
	var v13463 int32
	_ = v13463
	var v13469 int32
	_ = v13469
	var v13473 int32
	_ = v13473
	var v13477 int32
	_ = v13477
	var v13481 int32
	_ = v13481
	var v13487 int32
	_ = v13487
	var v13499 int32
	_ = v13499
	var v13500 int32
	_ = v13500
	var v13503 int32
	_ = v13503
	var v13505 int32
	_ = v13505
	var v13511 int32
	_ = v13511
	var v13515 int32
	_ = v13515
	var v13527 int32
	_ = v13527
	var v13529 int32
	_ = v13529
	var v13549 int32
	_ = v13549
	var v13552 int32
	_ = v13552
	var v13553 int32
	_ = v13553
	var v13554 int32
	_ = v13554
	var v13555 int32
	_ = v13555
	var v13556 int32
	_ = v13556
	var v13557 int32
	_ = v13557
	var v13559 int32
	_ = v13559
	var v13577 int32
	_ = v13577
	var v13580 int32
	_ = v13580
	var v13581 int32
	_ = v13581
	var v13584 int32
	_ = v13584
	var v13585 int32
	_ = v13585
	var v13590 int32
	_ = v13590
	var v13597 int32
	_ = v13597
	var v13601 int32
	_ = v13601
	var v13607 int32
	_ = v13607
	var v13611 int32
	_ = v13611
	var v13615 int32
	_ = v13615
	var v13619 int32
	_ = v13619
	var v13625 int32
	_ = v13625
	var v13637 int32
	_ = v13637
	var v13638 int32
	_ = v13638
	var v13640 int32
	_ = v13640
	var v13644 int32
	_ = v13644
	var v13645 int32
	_ = v13645
	var v13647 int32
	_ = v13647
	var v13648 int32
	_ = v13648
	var v13650 int32
	_ = v13650
	var v13653 int32
	_ = v13653
	var v13654 int32
	_ = v13654
	var v13659 int64
	_ = v13659
	var v13660 int32
	_ = v13660
	var v13661 int32
	_ = v13661
	var v13662 int32
	_ = v13662
	var v13663 int32
	_ = v13663
	var v13667 int32
	_ = v13667
	var v13670 int32
	_ = v13670
	var v13671 int32
	_ = v13671
	var v13681 int32
	_ = v13681
	var v13714 int64
	_ = v13714
	var v13715 int32
	_ = v13715
	var v13719 int32
	_ = v13719
	var v13722 int32
	_ = v13722
	var v13723 int32
	_ = v13723
	var v13725 int32
	_ = v13725
	var v13726 int32
	_ = v13726
	var v13728 int64
	_ = v13728
	var v13730 int32
	_ = v13730
	var v13731 int32
	_ = v13731
	var v13773 int64
	_ = v13773
	var v13774 int64
	_ = v13774
	var v13777 int64
	_ = v13777
	var v13780 int32
	_ = v13780
	var v13781 int32
	_ = v13781
	var v13788 int32
	_ = v13788
	var v13823 int32
	_ = v13823
	var v13824 int32
	_ = v13824
	var v13825 int32
	_ = v13825
	var v13829 int32
	_ = v13829
	var v13832 int32
	_ = v13832
	var v13834 int32
	_ = v13834
	var v13836 int32
	_ = v13836
	var v13839 int32
	_ = v13839
	var v13846 int32
	_ = v13846
	var v13857 int32
	_ = v13857
	var v13858 int32
	_ = v13858
	var v13880 int32
	_ = v13880
	var v13884 int32
	_ = v13884
	var v13885 int32
	_ = v13885
	var v13888 int32
	_ = v13888
	var v13889 int32
	_ = v13889
	var v13892 int32
	_ = v13892
	var v13900 int32
	_ = v13900
	var v13911 int32
	_ = v13911
	var v13933 int32
	_ = v13933
	var v13937 int32
	_ = v13937
	var v13938 int32
	_ = v13938
	var v13939 int32
	_ = v13939
	var v13940 int32
	_ = v13940
	var v13941 int32
	_ = v13941
	var v13942 int32
	_ = v13942
	var v13943 int32
	_ = v13943
	var v13944 int32
	_ = v13944
	var v13945 int32
	_ = v13945
	var v13946 int32
	_ = v13946
	var v13947 int32
	_ = v13947
	var v13948 int32
	_ = v13948
	var v13949 int32
	_ = v13949
	var v13950 int32
	_ = v13950
	var v13951 int32
	_ = v13951
	var v13952 int32
	_ = v13952
	var v13953 int32
	_ = v13953
	var v13954 int32
	_ = v13954
	var v13956 int32
	_ = v13956
	var v13957 int32
	_ = v13957
	var v13958 int32
	_ = v13958
	var v13960 int32
	_ = v13960
	var v13961 int32
	_ = v13961
	var v13963 int32
	_ = v13963
	var v13964 int32
	_ = v13964
	var v13971 int32
	_ = v13971
	var v13983 int32
	_ = v13983
	var v14006 int32
	_ = v14006
	var v14008 int32
	_ = v14008
	var v14009 int32
	_ = v14009
	var v14013 int32
	_ = v14013
	var v14014 int32
	_ = v14014
	var v14027 int32
	_ = v14027
	var v14033 int32
	_ = v14033
	var v14056 int32
	_ = v14056
	var v14057 int32
	_ = v14057
	var v14060 int32
	_ = v14060
	var v14061 int32
	_ = v14061
	var v14062 int32
	_ = v14062
	var v14070 int32
	_ = v14070
	var v14073 int32
	_ = v14073
	var v14075 int32
	_ = v14075
	var v14077 int32
	_ = v14077
	var v14079 int32
	_ = v14079
	var v14081 int32
	_ = v14081
	var v14088 int32
	_ = v14088
	var v14089 float64
	_ = v14089
	var v14090 float64
	_ = v14090
	var v14091 float64
	_ = v14091
	var v14092 int32
	_ = v14092
	var v14094 int32
	_ = v14094
	var v14096 int32
	_ = v14096
	var v14098 int32
	_ = v14098
	var v14099 int32
	_ = v14099
	var v14100 int32
	_ = v14100
	var v14101 int32
	_ = v14101
	var v14102 int32
	_ = v14102
	var v14103 int32
	_ = v14103
	var v14106 int32
	_ = v14106
	var v14110 int32
	_ = v14110
	var v14125 int32
	_ = v14125
	var v14141 float64
	_ = v14141
	var v14146 float64
	_ = v14146
	var v14152 int32
	_ = v14152
	var v14156 int32
	_ = v14156
	var v14157 int64
	_ = v14157
	var v14161 int32
	_ = v14161
	var v14165 int32
	_ = v14165
	var v14166 int32
	_ = v14166
	var v14167 float64
	_ = v14167
	var v14168 float64
	_ = v14168
	var v14170 int64
	_ = v14170
	var v14175 int32
	_ = v14175
	var v14176 int32
	_ = v14176
	var v14177 int32
	_ = v14177
	var v14178 int64
	_ = v14178
	var v14180 int64
	_ = v14180
	var v14182 int32
	_ = v14182
	var v14183 float64
	_ = v14183
	var v14184 float64
	_ = v14184
	var v14186 int64
	_ = v14186
	var v14190 int32
	_ = v14190
	var v14191 int32
	_ = v14191
	var v14192 int64
	_ = v14192
	var v14194 int64
	_ = v14194
	var v14198 float64
	_ = v14198
	var v14199 float64
	_ = v14199
	var v14201 float64
	_ = v14201
	var v14204 float64
	_ = v14204
	var v14206 int32
	_ = v14206
	var v14207 int32
	_ = v14207
	var v14239 float64
	_ = v14239
	var v14244 float64
	_ = v14244
	var v14251 float64
	_ = v14251
	var v14253 float64
	_ = v14253
	var v14263 float64
	_ = v14263
	var v14265 int32
	_ = v14265
	var v14266 int32
	_ = v14266
	var v14267 int32
	_ = v14267
	var v14268 int32
	_ = v14268
	var v14269 int32
	_ = v14269
	var v14270 int32
	_ = v14270
	var v14271 int32
	_ = v14271
	var v14273 float64
	_ = v14273
	var v14274 int32
	_ = v14274
	var v14276 int32
	_ = v14276
	var v14279 float64
	_ = v14279
	var v14281 int32
	_ = v14281
	var v14282 int32
	_ = v14282
	var v14283 int32
	_ = v14283
	var v14284 int32
	_ = v14284
	var v14285 int32
	_ = v14285
	var v14286 int32
	_ = v14286
	var v14288 float64
	_ = v14288
	var v14289 int32
	_ = v14289
	var v14291 int32
	_ = v14291
	var v14296 float64
	_ = v14296
	var v14309 int32
	_ = v14309
	var v14310 float64
	_ = v14310
	var v14311 float64
	_ = v14311
	var v14316 int32
	_ = v14316
	var v14317 int32
	_ = v14317
	var v14321 int32
	_ = v14321
	var v14322 int32
	_ = v14322
	var v14325 int32
	_ = v14325
	var v14327 int32
	_ = v14327
	var v14329 int32
	_ = v14329
	var v14330 int64
	_ = v14330
	var v14338 float64
	_ = v14338
	var v14351 float64
	_ = v14351
	var v14353 int32
	_ = v14353
	var v14356 int32
	_ = v14356
	var v14361 float64
	_ = v14361
	var v14364 float64
	_ = v14364
	var v14377 float64
	_ = v14377
	var v14382 float64
	_ = v14382
	var v14388 float64
	_ = v14388
	var v14395 float64
	_ = v14395
	var v14396 float64
	_ = v14396
	var v14399 float64
	_ = v14399
	var v14400 float64
	_ = v14400
	var v14401 float64
	_ = v14401
	var v14403 float64
	_ = v14403
	var v14408 int32
	_ = v14408
	var v14409 int32
	_ = v14409
	var v14432 int32
	_ = v14432
	var v14453 int32
	_ = v14453
	var v14496 int32
	_ = v14496
	var v14497 int32
	_ = v14497
	var v14499 int32
	_ = v14499
	var v14502 int32
	_ = v14502
	var v14504 float64
	_ = v14504
	var v14506 int32
	_ = v14506
	var v14507 int32
	_ = v14507
	var v14510 int32
	_ = v14510
	var v14511 int32
	_ = v14511
	var v14515 int32
	_ = v14515
	var v14516 int32
	_ = v14516
	var v14524 int32
	_ = v14524
	var v14525 int32
	_ = v14525
	var v14534 float64
	_ = v14534
	var v14538 int64
	_ = v14538
	var v14539 int64
	_ = v14539
	var v14543 int32
	_ = v14543
	var v14549 int32
	_ = v14549
	var v14552 int32
	_ = v14552
	var v14556 int32
	_ = v14556
	var v14558 int32
	_ = v14558
	var v14559 int32
	_ = v14559
	var v14562 int32
	_ = v14562
	var v14563 int32
	_ = v14563
	var v14565 int32
	_ = v14565
	var v14568 int32
	_ = v14568
	var v14570 float64
	_ = v14570
	var v14572 int32
	_ = v14572
	var v14576 int32
	_ = v14576
	var v14577 int32
	_ = v14577
	var v14581 int32
	_ = v14581
	var v14582 int32
	_ = v14582
	var v14590 int32
	_ = v14590
	var v14591 int32
	_ = v14591
	var v14600 float64
	_ = v14600
	var v14604 int64
	_ = v14604
	var v14605 int64
	_ = v14605
	var v14607 int32
	_ = v14607
	var v14612 int32
	_ = v14612
	var v14613 int32
	_ = v14613
	var v14614 int32
	_ = v14614
	var v14616 int32
	_ = v14616
	var v14618 int32
	_ = v14618
	var v14620 int32
	_ = v14620
	var v14622 int32
	_ = v14622
	var v14624 int32
	_ = v14624
	var v14625 int32
	_ = v14625
	var v14626 int32
	_ = v14626
	var v14629 int32
	_ = v14629
	var v14632 int32
	_ = v14632
	var v14633 int32
	_ = v14633
	var v14636 int32
	_ = v14636
	var v14637 int32
	_ = v14637
	var v14639 int32
	_ = v14639
	var v14641 int32
	_ = v14641
	var v14643 int32
	_ = v14643
	var v14645 int32
	_ = v14645
	var v14647 int32
	_ = v14647
	var v14649 int32
	_ = v14649
	var v14650 int32
	_ = v14650
	var v14651 int32
	_ = v14651
	var v14652 int32
	_ = v14652
	var v14653 int32
	_ = v14653
	var v14654 int32
	_ = v14654
	var v14655 int32
	_ = v14655
	var v14656 float64
	_ = v14656
	var v14657 int32
	_ = v14657
	var v14659 float64
	_ = v14659
	var v14660 int32
	_ = v14660
	var v14661 int32
	_ = v14661
	var v14670 int32
	_ = v14670
	var v14673 int32
	_ = v14673
	var v14676 int32
	_ = v14676
	var v14677 int32
	_ = v14677
	var v14679 int32
	_ = v14679
	var v14687 int32
	_ = v14687
	var v14688 int32
	_ = v14688
	var v14689 int32
	_ = v14689
	var v14690 int32
	_ = v14690
	var v14694 int32
	_ = v14694
	var v14698 int32
	_ = v14698
	var v14706 int32
	_ = v14706
	var v14709 int32
	_ = v14709
	var v14712 int32
	_ = v14712
	var v14729 int32
	_ = v14729
	var v14757 int32
	_ = v14757
	var v14758 int32
	_ = v14758
	var v14762 int32
	_ = v14762
	var v14763 int32
	_ = v14763
	var v14764 int32
	_ = v14764
	var v14765 int32
	_ = v14765
	var v14768 int32
	_ = v14768
	var v14769 int32
	_ = v14769
	var v14779 int32
	_ = v14779
	var v14813 int32
	_ = v14813
	var v14817 int32
	_ = v14817
	var v14818 int32
	_ = v14818
	var v14820 int32
	_ = v14820
	var v14838 int32
	_ = v14838
	var v14841 int32
	_ = v14841
	var v14842 int32
	_ = v14842
	var v14845 int32
	_ = v14845
	var v14846 int32
	_ = v14846
	var v14851 int32
	_ = v14851
	var v14858 int32
	_ = v14858
	var v14862 int32
	_ = v14862
	var v14868 int32
	_ = v14868
	var v14872 int32
	_ = v14872
	var v14876 int32
	_ = v14876
	var v14880 int32
	_ = v14880
	var v14886 int32
	_ = v14886
	var v14898 int32
	_ = v14898
	var v14900 int32
	_ = v14900
	var v14901 int32
	_ = v14901
	var v14905 int32
	_ = v14905
	var v14910 int32
	_ = v14910
	var v14914 int32
	_ = v14914
	var v14918 int32
	_ = v14918
	var v14919 int32
	_ = v14919
	var v14921 int32
	_ = v14921
	var v14922 int32
	_ = v14922
	var v14927 int32
	_ = v14927
	var v14928 int32
	_ = v14928
	var v14931 int32
	_ = v14931
	var v14937 int32
	_ = v14937
	var v14938 int32
	_ = v14938
	var v14939 int32
	_ = v14939
	var v14941 int32
	_ = v14941
	var v14942 int32
	_ = v14942
	var v14946 int32
	_ = v14946
	var v14947 int32
	_ = v14947
	var v14949 int32
	_ = v14949
	var v14950 int32
	_ = v14950
	var v14951 int32
	_ = v14951
	var v14952 int32
	_ = v14952
	var v14954 int32
	_ = v14954
	var v14959 int32
	_ = v14959
	var v14960 int32
	_ = v14960
	var v15004 int32
	_ = v15004
	var v15005 int32
	_ = v15005
	var v15049 int32
	_ = v15049
	var v15052 int32
	_ = v15052
	var v15053 int32
	_ = v15053
	var v15060 int32
	_ = v15060
	var v15063 int32
	_ = v15063
	var v15066 int32
	_ = v15066
	var v15067 int32
	_ = v15067
	var v15071 int32
	_ = v15071
	var v15075 int32
	_ = v15075
	var v15076 int32
	_ = v15076
	var v15080 int32
	_ = v15080
	var v15084 int32
	_ = v15084
	var v15090 int32
	_ = v15090
	var v15093 int32
	_ = v15093
	var v15095 int32
	_ = v15095
	var v15096 int32
	_ = v15096
	var v15099 int32
	_ = v15099
	var v15100 int32
	_ = v15100
	var v15102 int32
	_ = v15102
	var v15103 int32
	_ = v15103
	var v15106 int32
	_ = v15106
	var v15112 int32
	_ = v15112
	var v15115 int32
	_ = v15115
	var v15119 int32
	_ = v15119
	var v15120 int32
	_ = v15120
	var v15125 int32
	_ = v15125
	var v15127 int32
	_ = v15127
	var v15128 int32
	_ = v15128
	var v15129 int32
	_ = v15129
	var v15171 int32
	_ = v15171
	var v15174 int32
	_ = v15174
	var v15177 int32
	_ = v15177
	var v15183 int32
	_ = v15183
	var v15186 int32
	_ = v15186
	var v15190 int32
	_ = v15190
	var v15192 int32
	_ = v15192
	var v15197 float64
	_ = v15197
	var v15199 int32
	_ = v15199
	var v15204 int32
	_ = v15204
	var v15205 int32
	_ = v15205
	var v15208 int32
	_ = v15208
	var v15209 int32
	_ = v15209
	var v15217 int32
	_ = v15217
	var v15218 int32
	_ = v15218
	var v15227 float64
	_ = v15227
	var v15231 int64
	_ = v15231
	var v15232 int64
	_ = v15232
	var v15234 int32
	_ = v15234
	var v15237 int32
	_ = v15237
	var v15240 int32
	_ = v15240
	var v15241 int32
	_ = v15241
	var v15242 int32
	_ = v15242
	var v15246 int32
	_ = v15246
	var v15248 int32
	_ = v15248
	var v15250 int32
	_ = v15250
	var v15252 int32
	_ = v15252
	var v15254 int32
	_ = v15254
	var v15256 int32
	_ = v15256
	var v15259 int32
	_ = v15259
	var v15271 int32
	_ = v15271
	var v15304 int32
	_ = v15304
	var v15305 int32
	_ = v15305
	var v15309 int32
	_ = v15309
	var v15310 int32
	_ = v15310
	var v15312 int32
	_ = v15312
	var v15330 int32
	_ = v15330
	var v15333 int32
	_ = v15333
	var v15334 int32
	_ = v15334
	var v15337 int32
	_ = v15337
	var v15338 int32
	_ = v15338
	var v15343 int32
	_ = v15343
	var v15350 int32
	_ = v15350
	var v15354 int32
	_ = v15354
	var v15360 int32
	_ = v15360
	var v15364 int32
	_ = v15364
	var v15368 int32
	_ = v15368
	var v15372 int32
	_ = v15372
	var v15378 int32
	_ = v15378
	var v15390 int32
	_ = v15390
	var v15392 int32
	_ = v15392
	var v15393 int32
	_ = v15393
	var v15397 int32
	_ = v15397
	var v15402 int32
	_ = v15402
	var v15403 int32
	_ = v15403
	var v15407 int32
	_ = v15407
	var v15410 int32
	_ = v15410
	var v15411 int32
	_ = v15411
	var v15412 int32
	_ = v15412
	var v15413 int32
	_ = v15413
	var v15416 int32
	_ = v15416
	var v15418 int32
	_ = v15418
	var v15419 int32
	_ = v15419
	var v15420 int32
	_ = v15420
	var v15421 int32
	_ = v15421
	var v15422 int32
	_ = v15422
	var v15423 int32
	_ = v15423
	var v15424 int32
	_ = v15424
	var v15425 int32
	_ = v15425
	var v15427 int32
	_ = v15427
	var v15433 int32
	_ = v15433
	var v15434 int32
	_ = v15434
	var v15477 int32
	_ = v15477
	var v15480 int32
	_ = v15480
	var v15483 int32
	_ = v15483
	var v15486 int32
	_ = v15486
	var v15489 int32
	_ = v15489
	var v15490 int32
	_ = v15490
	var v15499 int32
	_ = v15499
	var v15533 int32
	_ = v15533
	var v15534 int32
	_ = v15534
	var v15538 int32
	_ = v15538
	var v15539 int32
	_ = v15539
	var v15541 int32
	_ = v15541
	var v15559 int32
	_ = v15559
	var v15562 int32
	_ = v15562
	var v15563 int32
	_ = v15563
	var v15566 int32
	_ = v15566
	var v15567 int32
	_ = v15567
	var v15572 int32
	_ = v15572
	var v15579 int32
	_ = v15579
	var v15583 int32
	_ = v15583
	var v15589 int32
	_ = v15589
	var v15593 int32
	_ = v15593
	var v15597 int32
	_ = v15597
	var v15601 int32
	_ = v15601
	var v15607 int32
	_ = v15607
	var v15619 int32
	_ = v15619
	var v15621 int32
	_ = v15621
	var v15622 int32
	_ = v15622
	var v15626 int32
	_ = v15626
	var v15631 int32
	_ = v15631
	var v15632 int32
	_ = v15632
	var v15636 int32
	_ = v15636
	var v15639 int32
	_ = v15639
	var v15640 int32
	_ = v15640
	var v15641 int32
	_ = v15641
	var v15642 int32
	_ = v15642
	var v15643 int32
	_ = v15643
	var v15647 float64
	_ = v15647
	var v15648 int32
	_ = v15648
	var v15649 float64
	_ = v15649
	var v15651 int32
	_ = v15651
	var v15657 float64
	_ = v15657
	var v15661 float64
	_ = v15661
	var v15663 float64
	_ = v15663
	var v15665 float64
	_ = v15665
	var v15666 float64
	_ = v15666
	var v15675 float64
	_ = v15675
	var v15679 float64
	_ = v15679
	var v15681 int32
	_ = v15681
	var v15682 int32
	_ = v15682
	var v15685 int32
	_ = v15685
	var v15686 int32
	_ = v15686
	var v15687 int32
	_ = v15687
	var v15688 int32
	_ = v15688
	var v15689 int32
	_ = v15689
	var v15690 int32
	_ = v15690
	var v15691 int32
	_ = v15691
	var v15692 int32
	_ = v15692
	var v15693 int32
	_ = v15693
	var v15694 int32
	_ = v15694
	var v15696 int32
	_ = v15696
	var v15701 int32
	_ = v15701
	var v15702 int32
	_ = v15702
	var v15745 int32
	_ = v15745
	var v15748 int32
	_ = v15748
	var v15754 int32
	_ = v15754
	var v15757 int32
	_ = v15757
	var v15761 int32
	_ = v15761
	var v15762 int32
	_ = v15762
	var v15765 int32
	_ = v15765
	var v15766 int32
	_ = v15766
	var v15768 int32
	_ = v15768
	var v15779 int32
	_ = v15779
	var v15812 int32
	_ = v15812
	var v15813 int32
	_ = v15813
	var v15814 int32
	_ = v15814
	var v15817 int32
	_ = v15817
	var v15818 int32
	_ = v15818
	var v15819 int32
	_ = v15819
	var v15822 int32
	_ = v15822
	var v15823 int32
	_ = v15823
	var v15824 int32
	_ = v15824
	var v15827 int32
	_ = v15827
	var v15829 int32
	_ = v15829
	var v15831 int32
	_ = v15831
	var v15833 int32
	_ = v15833
	var v15835 int32
	_ = v15835
	var v15837 int32
	_ = v15837
	var v15840 int32
	_ = v15840
	var v15844 int32
	_ = v15844
	var v15850 int32
	_ = v15850
	var v15853 int32
	_ = v15853
	var v15854 int32
	_ = v15854
	var v15855 int32
	_ = v15855
	var v15860 int32
	_ = v15860
	var v15866 int32
	_ = v15866
	var v15869 int32
	_ = v15869
	var v15878 float64
	_ = v15878
	var v15882 int64
	_ = v15882
	var v15883 int64
	_ = v15883
	var v15885 int32
	_ = v15885
	var v15889 int32
	_ = v15889
	var v15890 int32
	_ = v15890
	var v15891 int32
	_ = v15891
	var v15892 int32
	_ = v15892
	var v15893 int32
	_ = v15893
	var v15895 int32
	_ = v15895
	var v15896 int32
	_ = v15896
	var v15900 int32
	_ = v15900
	var v15901 int32
	_ = v15901
	var v15908 float64
	_ = v15908
	var v15915 int32
	_ = v15915
	var v15917 float64
	_ = v15917
	var v15920 float64
	_ = v15920
	var v15921 float64
	_ = v15921
	var v15923 float64
	_ = v15923
	var v15931 int32
	_ = v15931
	var v15932 int32
	_ = v15932
	var v15933 int32
	_ = v15933
	var v15936 int32
	_ = v15936
	var v15939 int32
	_ = v15939
	var v15942 int32
	_ = v15942
	var v15945 int32
	_ = v15945
	var v15946 int32
	_ = v15946
	var v15947 int64
	_ = v15947
	var v15951 int32
	_ = v15951
	var v15952 int32
	_ = v15952
	var v15953 int32
	_ = v15953
	var v15954 int32
	_ = v15954
	var v15956 int32
	_ = v15956
	var v15957 int32
	_ = v15957
	var v15960 int32
	_ = v15960
	var v15961 int32
	_ = v15961
	var v15969 int32
	_ = v15969
	var v15970 int32
	_ = v15970
	var v15973 int32
	_ = v15973
	var v15977 int32
	_ = v15977
	var v15979 int32
	_ = v15979
	var v15986 int32
	_ = v15986
	var v15987 int32
	_ = v15987
	var v15988 int32
	_ = v15988
	var v15993 int32
	_ = v15993
	var v15995 int32
	_ = v15995
	var v15998 int32
	_ = v15998
	var v16006 int32
	_ = v16006
	var v16007 int32
	_ = v16007
	var v16010 int32
	_ = v16010
	var v16011 int32
	_ = v16011
	var v16012 int32
	_ = v16012
	var v16013 int32
	_ = v16013
	var v16018 int32
	_ = v16018
	var v16028 int32
	_ = v16028
	var v16029 int32
	_ = v16029
	var v16032 int32
	_ = v16032
	var v16036 int32
	_ = v16036
	var v16039 int32
	_ = v16039
	var v16041 int32
	_ = v16041
	var v16044 int32
	_ = v16044
	var v16051 int32
	_ = v16051
	var v16053 int32
	_ = v16053
	var v16061 int32
	_ = v16061
	var v16062 int32
	_ = v16062
	var v16075 int32
	_ = v16075
	var v16081 int32
	_ = v16081
	var v16084 int32
	_ = v16084
	var v16087 int32
	_ = v16087
	var v16092 int32
	_ = v16092
	var v16097 int32
	_ = v16097
	var v16099 int32
	_ = v16099
	var v16100 int32
	_ = v16100
	var v16120 int32
	_ = v16120
	var v16121 int32
	_ = v16121
	var v16122 int32
	_ = v16122
	var v16124 int32
	_ = v16124
	var v16127 int32
	_ = v16127
	var v16128 int32
	_ = v16128
	var v16131 int32
	_ = v16131
	var v16132 int32
	_ = v16132
	var v16139 int32
	_ = v16139
	var v16145 int32
	_ = v16145
	var v16146 int32
	_ = v16146
	var v16147 int32
	_ = v16147
	var v16148 int32
	_ = v16148
	var v16151 int32
	_ = v16151
	var v16153 int32
	_ = v16153
	var v16154 int32
	_ = v16154
	var v16155 int32
	_ = v16155
	var v16156 int32
	_ = v16156
	var v16157 int32
	_ = v16157
	var v16158 int32
	_ = v16158
	var v16159 int32
	_ = v16159
	var v16161 int32
	_ = v16161
	var v16162 int32
	_ = v16162
	var v16164 int32
	_ = v16164
	var v16165 int32
	_ = v16165
	var v16166 int32
	_ = v16166
	var v16167 int32
	_ = v16167
	var v16168 int32
	_ = v16168
	var v16169 int32
	_ = v16169
	var v16170 int32
	_ = v16170
	var v16172 int32
	_ = v16172
	var v16173 int32
	_ = v16173
	var v16174 int32
	_ = v16174
	var v16175 int32
	_ = v16175
	var v16176 int32
	_ = v16176
	var v16177 int32
	_ = v16177
	var v16178 int32
	_ = v16178
	var v16179 int32
	_ = v16179
	var v16180 int32
	_ = v16180
	var v16196 int32
	_ = v16196
	var v16201 int32
	_ = v16201
	var v16225 int32
	_ = v16225
	var v16229 int32
	_ = v16229
	var v16230 int32
	_ = v16230
	var v16231 int32
	_ = v16231
	var v16232 int32
	_ = v16232
	var v16233 int32
	_ = v16233
	var v16234 int32
	_ = v16234
	var v16236 int32
	_ = v16236
	var v16237 int32
	_ = v16237
	var v16238 int32
	_ = v16238
	var v16240 int32
	_ = v16240
	var v16243 int32
	_ = v16243
	var v16244 int32
	_ = v16244
	var v16245 int32
	_ = v16245
	var v16246 int32
	_ = v16246
	var v16247 int32
	_ = v16247
	var v16249 int32
	_ = v16249
	var v16250 int32
	_ = v16250
	var v16252 int32
	_ = v16252
	var v16253 int32
	_ = v16253
	var v16267 int32
	_ = v16267
	var v16296 int32
	_ = v16296
	var v16297 int32
	_ = v16297
	var v16316 int32
	_ = v16316
	var v16339 int32
	_ = v16339
	var v16342 int32
	_ = v16342
	var v16344 int32
	_ = v16344
	var v16345 int32
	_ = v16345
	var v16346 int32
	_ = v16346
	var v16347 int32
	_ = v16347
	var v16348 int32
	_ = v16348
	var v16351 int32
	_ = v16351
	var v16357 int32
	_ = v16357
	var v16362 int32
	_ = v16362
	var v16367 int32
	_ = v16367
	var v16369 int32
	_ = v16369
	var v16370 int32
	_ = v16370
	var v16390 int32
	_ = v16390
	var v16397 int32
	_ = v16397
	var v16399 int32
	_ = v16399
	var v16400 int32
	_ = v16400
	var v16403 int32
	_ = v16403
	var v16407 int32
	_ = v16407
	var v16410 int32
	_ = v16410
	var v16412 int32
	_ = v16412
	var v16415 int32
	_ = v16415
	var v16422 int32
	_ = v16422
	var v16424 int32
	_ = v16424
	var v16432 int32
	_ = v16432
	var v16433 int32
	_ = v16433
	var v16446 int32
	_ = v16446
	var v16457 int32
	_ = v16457
	var v16462 int32
	_ = v16462
	var v16467 int32
	_ = v16467
	var v16469 int32
	_ = v16469
	var v16470 int32
	_ = v16470
	var v16490 int32
	_ = v16490
	var v16496 int32
	_ = v16496
	var v16497 int32
	_ = v16497
	var v16498 int32
	_ = v16498
	var v16501 int32
	_ = v16501
	var v16507 int32
	_ = v16507
	var v16508 int32
	_ = v16508
	var v16510 int32
	_ = v16510
	var v16511 int32
	_ = v16511
	var v16517 int32
	_ = v16517
	var v16518 int32
	_ = v16518
	var v16519 int32
	_ = v16519
	var v16520 int32
	_ = v16520
	var v16526 int32
	_ = v16526
	var v16527 int32
	_ = v16527
	var v16528 int32
	_ = v16528
	var v16529 int32
	_ = v16529
	var v16535 int32
	_ = v16535
	var v16536 int32
	_ = v16536
	var v16537 int32
	_ = v16537
	var v16538 int32
	_ = v16538
	var v16548 int32
	_ = v16548
	var v16549 int32
	_ = v16549
	var v16550 int32
	_ = v16550
	var v16552 int32
	_ = v16552
	var v16555 int32
	_ = v16555
	var v16561 int32
	_ = v16561
	var v16562 int32
	_ = v16562
	var v16564 int32
	_ = v16564
	var v16565 int32
	_ = v16565
	var v16571 int32
	_ = v16571
	var v16572 int32
	_ = v16572
	var v16573 int32
	_ = v16573
	var v16574 int32
	_ = v16574
	var v16576 int32
	_ = v16576
	var v16582 int32
	_ = v16582
	var v16583 int32
	_ = v16583
	var v16584 int32
	_ = v16584
	var v16585 int32
	_ = v16585
	var v16591 int32
	_ = v16591
	var v16592 int32
	_ = v16592
	var v16593 int32
	_ = v16593
	var v16594 int32
	_ = v16594
	var v16595 int32
	_ = v16595
	var v16603 int32
	_ = v16603
	var v16614 int32
	_ = v16614
	var v16616 int32
	_ = v16616
	var v16619 int32
	_ = v16619
	var v16621 int32
	_ = v16621
	var v16622 int32
	_ = v16622
	var v16642 int32
	_ = v16642
	var v16643 int32
	_ = v16643
	var v16647 int32
	_ = v16647
	var v16648 int32
	_ = v16648
	var v16651 int32
	_ = v16651
	var v16657 int32
	_ = v16657
	var v16662 int32
	_ = v16662
	var v16664 int32
	_ = v16664
	var v16667 int32
	_ = v16667
	var v16669 int32
	_ = v16669
	var v16670 int32
	_ = v16670
	var v16690 int32
	_ = v16690
	var v16691 int32
	_ = v16691
	var v16692 int32
	_ = v16692
	var v16693 int32
	_ = v16693
	var v16694 int32
	_ = v16694
	var v16696 int32
	_ = v16696
	var v16697 int32
	_ = v16697
	var v16698 int32
	_ = v16698
	var v16699 int32
	_ = v16699
	var v16700 int32
	_ = v16700
	var v16703 int32
	_ = v16703
	var v16704 int32
	_ = v16704
	var v16708 int32
	_ = v16708
	var v16709 int32
	_ = v16709
	var v16718 int32
	_ = v16718
	var v16720 float64
	_ = v16720
	var v16722 float64
	_ = v16722
	var v16724 float64
	_ = v16724
	var v16726 int32
	_ = v16726
	var v16727 int32
	_ = v16727
	var v16730 int32
	_ = v16730
	var v16747 int32
	_ = v16747
	var v16753 int32
	_ = v16753
	var v16756 int32
	_ = v16756
	var v16757 int32
	_ = v16757
	var v16758 int32
	_ = v16758
	var v16763 int32
	_ = v16763
	var v16772 int32
	_ = v16772
	var v16781 float64
	_ = v16781
	var v16785 int64
	_ = v16785
	var v16786 int64
	_ = v16786
	var v16788 int32
	_ = v16788
	var v16790 int32
	_ = v16790
	var v16792 int32
	_ = v16792
	var v16793 int32
	_ = v16793
	var v16795 int32
	_ = v16795
	var v16801 int32
	_ = v16801
	var v16805 int32
	_ = v16805
	var v16806 int32
	_ = v16806
	var v16811 int32
	_ = v16811
	var v16820 int32
	_ = v16820
	var v16829 float64
	_ = v16829
	var v16833 int64
	_ = v16833
	var v16834 int64
	_ = v16834
	var v16836 int32
	_ = v16836
	var v16839 int32
	_ = v16839
	var v16842 int32
	_ = v16842
	var v16843 int32
	_ = v16843
	var v16847 int32
	_ = v16847
	var v16850 int32
	_ = v16850
	var v16853 int32
	_ = v16853
	var v16856 int32
	_ = v16856
	var v16857 int32
	_ = v16857
	var v16858 int64
	_ = v16858
	var v16861 int32
	_ = v16861
	var v16864 int32
	_ = v16864
	var v16875 int32
	_ = v16875
	var v16909 int32
	_ = v16909
	var v16913 int32
	_ = v16913
	var v16915 int32
	_ = v16915
	var v16917 int32
	_ = v16917
	var v16918 int32
	_ = v16918
	var v16961 int32
	_ = v16961
	var v16962 int32
	_ = v16962
	var v16966 int32
	_ = v16966
	var v16967 int32
	_ = v16967
	var v16971 int32
	_ = v16971
	var v16974 int32
	_ = v16974
	var v16975 int32
	_ = v16975
	var v16978 int32
	_ = v16978
	var v16979 int32
	_ = v16979
	var v16980 int64
	_ = v16980
	var v16992 int32
	_ = v16992
	var v17030 int32
	_ = v17030
	var v17033 int32
	_ = v17033
	var v17040 int32
	_ = v17040
	var v17043 int32
	_ = v17043
	var v17048 int32
	_ = v17048
	var v17055 int32
	_ = v17055
	var v17058 int32
	_ = v17058
	var v17062 int32
	_ = v17062
	var v17066 int32
	_ = v17066
	var v17071 int32
	_ = v17071
	var v17073 int32
	_ = v17073
	var v17076 int32
	_ = v17076
	var v17077 int32
	_ = v17077
	var v17078 int32
	_ = v17078
	var v17085 float64
	_ = v17085
	var v17086 int32
	_ = v17086
	var v17089 int32
	_ = v17089
	var v17092 int32
	_ = v17092
	var v17101 int32
	_ = v17101
	var v17104 int32
	_ = v17104
	var v17109 int32
	_ = v17109
	var v17110 float64
	_ = v17110
	var v17111 int32
	_ = v17111
	var v17113 int32
	_ = v17113
	var v17114 int32
	_ = v17114
	var v17115 float64
	_ = v17115
	var v17116 float64
	_ = v17116
	var v17119 int32
	_ = v17119
	var v17120 float64
	_ = v17120
	var v17121 float64
	_ = v17121
	var v17123 float64
	_ = v17123
	var v17124 int32
	_ = v17124
	var v17125 int32
	_ = v17125
	var v17129 int32
	_ = v17129
	var v17131 int32
	_ = v17131
	var v17133 int32
	_ = v17133
	var v17137 int32
	_ = v17137
	var v17140 int32
	_ = v17140
	var v17146 float64
	_ = v17146
	var v17150 int32
	_ = v17150
	var v17151 float64
	_ = v17151
	var v17152 float64
	_ = v17152
	var v17155 int32
	_ = v17155
	var v17162 int32
	_ = v17162
	var v17168 float64
	_ = v17168
	var v17169 int32
	_ = v17169
	var v17172 int32
	_ = v17172
	var v17178 int32
	_ = v17178
	var v17188 int32
	_ = v17188
	var v17192 int32
	_ = v17192
	var v17193 float64
	_ = v17193
	var v17196 float64
	_ = v17196
	var v17199 int32
	_ = v17199
	var v17202 int32
	_ = v17202
	var v17203 int32
	_ = v17203
	var v17217 int32
	_ = v17217
	var v17221 int32
	_ = v17221
	var v17224 int32
	_ = v17224
	var v17228 int32
	_ = v17228
	var v17238 int32
	_ = v17238
	var v17242 int32
	_ = v17242
	var v17243 float64
	_ = v17243
	var v17246 float64
	_ = v17246
	var v17250 int32
	_ = v17250
	var v17251 int32
	_ = v17251
	var v17274 int32
	_ = v17274
	v4 = l3
	v7 = int32(0)
	v39 = int64(0)
	v42 = m.G0
	v44 = v42 - int32(16)
	m.G0 = v44
	v47 = F_palloc0(m, int32(384))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(267)
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v59 = v55 + int32(1)
	goto L5
L4:
	;
	v59 = int32(1)
	goto L5
L5:
	;
	v60 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v47)+20)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v59
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[0]))
	v66 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+116)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v47)+280)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v47)+72)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v47)+80)) = v60
	*(*int64)(unsafe.Add(mBase, uint32(v47)+85)) = v60
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v76 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v77 = F_bms_make_singleton(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v79 = v66
	goto L8
L8:
	;
	v80 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+321)) = uint8(v80)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+319)) = uint16(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+312)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v47)+120)) = v79
	v87 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v47)+124)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v47)+132)) = v87
	base.MemoryFill(m, v47+int32(192), v80, int32(88))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+322)) = uint8(v4)
	if v4 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v79 = v77
	goto L8
L10:
	;
	v98 = F_assign_special_exec_param(m, v47)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v100 = int32(-1)
	goto L12
L12:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+372)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+348)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v47)+344)) = v100
	v107 = F_palloc0(m, int32(8))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v100 = v98
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(272)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v107
	v116 = F_list_make1_impl(m, int32(1), v44+int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+84)) = v116
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v119 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v120 = int32(0)
	v121 = m.G0
	v123 = v121 - int32(32)
	m.G0 = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+48))
	if v126 == v120 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	goto L18
L18:
	;
	v482 = int32(0)
	v485 = m.G0
	v487 = v485 - int32(32)
	m.G0 = v487
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v489 == int32(5) {
		goto L92
	} else {
		goto L93
	}
L19:
	;
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L89
	}
L21:
	;
	m.G0 = v123 + int32(32)
	goto L19
L22:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v129 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v134 = v120
	goto L24
L24:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v134<<(uint(int32(2))%32))))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+36))
	if v180|base.B2i32(v179 != int32(1)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L21
L26:
	;
	v381 = v134 + int32(1)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	if v381 < v382 {
		v134 = v381
		goto L24
	} else {
		goto L88
	}
L27:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v188 = F_lappend_int(m, v186, int32(-1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	switch v191 {
	case 0:
		goto L33
	default:
		goto L31
	case 2:
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v188
	goto L26
L31:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v269 = F_copyObjectImpl(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L63
	}
L32:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+32)))
	if v194|base.B2i32(v179 != int32(1)) != 0 {
		goto L31
	} else {
		goto L35
	}
L33:
	;
	if v180 != int32(1) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v198 == int32(67) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v177)+36))
	if v211 < int32(2) {
		goto L45
	} else {
		goto L46
	}
L37:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v178)+140))
	if v201 != 0 {
		goto L31
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v209 = F_expression_tree_walker_impl(m, v178, int32(842), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v203 = int32(0)
	v205 = F_query_tree_walker_impl(m, v178, int32(842), v203, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v205 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	if v209 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L36
L45:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v248 = F_contain_volatile_functions(m, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L59
	}
L46:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = v215
	if v214 == v215 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v219 != int32(67) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v243 = F_expression_tree_walker_impl(m, v214, int32(843), v123+int32(20))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L57
	}
L49:
	;
	if v219 != int32(101) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = int32(1)
	v236 = F_query_tree_walker_impl(m, v214, int32(843), v123+int32(20), int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	if v224 != int32(6) {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+92)))
	if v227 == int32(0) {
		goto L45
	} else {
		goto L54
	}
L54:
	;
	goto L31
L55:
	;
	if v236 == int32(0) {
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L31
L57:
	;
	if v243 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	goto L45
L59:
	;
	if v248 != 0 {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v123)+20)) = v250
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+28)) = v254
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v259 = F_inline_cte_walker(m, v256, v123+int32(20))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v263 = F_lappend_int(m, v261, int32(-1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v263
	goto L26
L63:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+32)))
	v275 = F_subquery_planner(m, v271, v269, v47, v272, float64(0), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	if v277 != 0 {
		goto L20
	} else {
		goto L65
	}
L65:
	;
	v280 = F_fetch_upper_rel(m, v275, int32(7), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v280)+48))
	v283 = F_create_plan(m, v275, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v286 = F_palloc0(m, int32(72))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v286)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v286))) = int64(30064771095)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v283)+44))
	if v292 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v315 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+48)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v286)+40)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286)+38)) = uint8(v315)
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+36)) = uint16(v315)
	*(*int32)(unsafe.Add(mBase, uint32(v286)+32)) = v314
	v324 = F_assign_special_exec_param(m, v47)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L76
	}
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v286)+24)) = int64(-4294965018)
	v314 = int32(0)
	goto L69
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+26)))
	if v297 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v299 = F_exprType(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+24)) = v299
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v303 = F_exprTypmod(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+28)) = v303
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v307 = F_exprCollation(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v314 = v307
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v123)+16)) = v324
	v331 = F_list_make1_impl(m, int32(471), v123+int32(12))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+40)) = v331
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	v336 = F_lappend(m, v335, v283)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+8)) = v336
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	v342 = F_lappend(m, v341, v282)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v344)+12)) = v342
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)+16))
	v348 = F_lappend(m, v347, v275)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v350)+16)) = v348
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+8))
	if v353 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+4))
	v356 = v354
	goto L83
L82:
	;
	v356 = int32(0)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+16)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
	v359 = F_lappend(m, v358, v286)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v364 = F_lappend_int(m, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+76)) = v364
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v177)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v367
	v370 = F_psprintf(m, int32(_a_F_subquery_planner_0), v123)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v286)+20)) = v370
	F_cost_subplan(m, v286, v283)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L26
L88:
	;
	goto L25
L89:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_1), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_2), int32(977), int32(_a_F_subquery_planner_3))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v492 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v487)+30)) = uint16(v492)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v494 != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	m.G0 = v487 + int32(32)
	F_replace_empty_jointree(m, l1)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L174
	}
L95:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v495 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v695 = v482
	v700 = v482
	v733 = int32(0)
	goto L97
L97:
	;
	v735 = F_palloc0(m, int32(136))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L126
	}
L98:
	;
	v677 = int32(1)
	v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+31)))
	v680 = v678 & v677
	if v680 != 0 {
		goto L117
	} else {
		goto L118
	}
L99:
	;
	if v495 != int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v500 = int32(0)
	if v500 < v495 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v595 = v7
	goto L102
L102:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v621+v595<<(uint(int32(2))%32))))
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v625)+8))
	if v626 == int32(7) {
		goto L98
	} else {
		goto L116
	}
L103:
	;
	v503 = v495
	goto L105
L104:
	;
	v503 = v500
	goto L105
L105:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v517 = v482
	v524 = v7
	goto L106
L106:
	;
	v552 = v508 + v524<<(uint(int32(2))%32)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	if v554 != int32(7) {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v503&int32(1) == int32(0) {
		goto L98
	} else {
		goto L115
	}
L108:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v561 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v557+(v487+int32(29))))) = uint8(v561)
	goto L110
L109:
	;
	goto L110
L110:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v563)+8))
	if v564 != int32(7) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	v571 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v567+(v487+int32(29))))) = uint8(v571)
	goto L113
L112:
	;
	goto L113
L113:
	;
	v573 = int32(2)
	v574 = v524 + v573
	v576 = v517 + v573
	if v576 != v503&int32(2147483646) {
		v517 = v576
		v524 = v574
		goto L106
	} else {
		goto L114
	}
L114:
	;
	goto L107
L115:
	;
	v595 = v574
	goto L102
L116:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	v633 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v629+(v487+int32(29))))) = uint8(v633)
	goto L98
L117:
	;
	v681 = int32(2)
	goto L119
L118:
	;
	v681 = v677
	goto L119
L119:
	;
	if v680 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v684 = int32(3)
	goto L122
L121:
	;
	v684 = int32(0)
	goto L122
L122:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v487)+30)))
	v687 = v685 & int32(1)
	if v687 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v688 = v681
	goto L125
L124:
	;
	v688 = v684
	goto L125
L125:
	;
	v695 = v678
	v700 = v688
	v733 = int32(0) - v687
	goto L97
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v735)+44)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v735)+12)) = int32(2)
	v740 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v735)+48)) = v740
	*(*int64)(unsafe.Add(mBase, uint32(v735))) = int64(101)
	*(*int64)(unsafe.Add(mBase, uint32(v735)+56)) = v740
	v746 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v735)+64)) = v746
	v750 = F_makeAlias(m, int32(_a_F_subquery_planner_4), v746)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v735)+8)) = v750
	v753 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v735)+124)) = uint16(v753)
	v755 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v735)+20)) = uint8(v755)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v758 = F_lappend(m, v757, v735)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v758
	if v758 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	v762 = v761
	goto L131
L130:
	;
	v762 = v482
	goto L131
L131:
	;
	v764 = F_palloc0(m, int32(8))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = int32(63)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v764)+4)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v487)+16)) = v764
	*(*int32)(unsafe.Add(mBase, uint32(v487)+24)) = v764
	v775 = F_list_make1_impl(m, int32(1), v487+int32(16))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+8))
	v779 = F_makeFromExpr(m, v775, v778)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v783)+12))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v784)))
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	switch v786 - int32(63) {
	case 0:
		v804 = int32(4)
		goto L135
	case 1:
		goto L136
	default:
		goto L137
	}
L135:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v785+v804)))
	v808 = F_palloc0(m, int32(40))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L1
	} else {
		goto L141
	}
L136:
	;
	v804 = int32(36)
	goto L135
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v785)))
	*(*int32)(unsafe.Add(mBase, uint32(v487))) = v793
	F_errmsg_internal(m, int32(_a_F_subquery_planner_5), v487)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_6), int32(283), int32(_a_F_subquery_planner_7))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v808)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v808)+16)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v808)+12)) = v779
	v814 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v808)+8)) = uint8(v814)
	*(*int32)(unsafe.Add(mBase, uint32(v808)+4)) = v700
	*(*int32)(unsafe.Add(mBase, uint32(v808))) = int32(64)
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v808)+36)) = v762
	*(*int32)(unsafe.Add(mBase, uint32(v808)+32)) = v814
	*(*int32)(unsafe.Add(mBase, uint32(v808)+28)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v487)+12)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v487)+20)) = v808
	v829 = F_list_make1_impl(m, int32(1), v487+int32(12))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v831)+4)) = v829
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v834 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v833)+8)) = v834
	v836 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if (base.B2i32(v836 == v834)|(v695^int32(-1)))&int32(1) == v834 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v847 = F_bms_make_singleton(m, v846)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	if v733&int32(1) != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v849 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v851 = F_add_nulling_relids(m, v836, v847, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v851
	goto L145
L149:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v857 = F_bms_make_singleton(m, v806)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	v1058 = int32(0)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v1058
	goto L94
L152:
	;
	v859 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v861 = F_add_nulling_relids(m, v856, v857, v859)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v861
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v864 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v979 = F_bms_make_singleton(m, v806)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L167
	}
L156:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	if v867 <= int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v879 = int32(0)
	goto L158
L158:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v864)+12))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v912+v879<<(uint(int32(2))%32))))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v916)+16))
	v918 = F_bms_make_singleton(m, v806)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	goto L155
L160:
	;
	v920 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v922 = F_add_nulling_relids(m, v917, v918, v920)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+16)) = v922
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v916)+20))
	v926 = F_bms_make_singleton(m, v806)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v928 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v930 = F_add_nulling_relids(m, v925, v926, v928)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v916)+20)) = v930
	v934 = v879 + int32(1)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v864)+4))
	if v934 < v935 {
		v879 = v934
		goto L158
	} else {
		goto L166
	}
L166:
	;
	goto L159
L167:
	;
	v981 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v983 = F_add_nulling_relids(m, v978, v979, v981)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v983
	v986 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v986)+12))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v987+v806<<(uint(int32(2))%32)-int32(4))))
	v994 = int32(0)
	v996 = F_makeWholeRowVar(m, v993, v806, v994, v994)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v998 = F_bms_make_singleton(m, v762)
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v996)+24)) = v998
	v1002 = F_palloc0(m, int32(20))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+16)) = int32(-1)
	v1006 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1002)+12)) = uint8(v1006)
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1002)+4)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v1002))) = int32(52)
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1014 = F_make_and_qual(m, v1002, v1013)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v1058 = v1014
	goto L151
L174:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+39)))
	if v1106 == int32(1) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v1109 = m.G0
	v1111 = v1109 - int32(16)
	m.G0 = v1111
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1113)+60))
	v1117 = F_pull_up_sublinks_jointree_recurse(m, v47, v1114, v1111+int32(12))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v1140 = int32(0)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1141)+52))
	if v1142 == v1140 {
		goto L184
	} else {
		goto L185
	}
L178:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1117)))
	if v1119 != int32(65) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1111)+4)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(v1111)+8)) = v1117
	v1127 = F_list_make1_impl(m, int32(1), v1111+int32(4))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	v1132 = v1117
	goto L181
L181:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1133)+60)) = v1132
	m.G0 = v1111 + int32(16)
	goto L177
L182:
	;
	v1130 = F_makeFromExpr(m, v1127, int32(0))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v1132 = v1130
	goto L181
L184:
	;
	v1258 = F_expand_virtual_generated_columns(m, v47)
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L195
	}
L185:
	;
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	if v1145 <= int32(0) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v1150 = v1140
	goto L187
L187:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+12))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1150<<(uint(int32(2))%32))))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+12))
	if v1194 != int32(3) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L184
L189:
	;
	v1214 = v1150 + int32(1)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	if v1214 < v1215 {
		v1150 = v1214
		goto L187
	} else {
		goto L194
	}
L190:
	;
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1193)+68))
	v1198 = F_eval_const_expressions(m, v47, v1197)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1193)+68)) = v1198
	v1201 = F_inline_set_returning_function(m, v47, v1193)
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	if v1201 == int32(0) {
		goto L189
	} else {
		goto L193
	}
L193:
	;
	v1205 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1193)+72)) = uint8(v1205)
	*(*uint8)(unsafe.Add(mBase, uint32(v1193)+40)) = uint8(v1205)
	*(*int32)(unsafe.Add(mBase, uint32(v1193)+36)) = v1201
	*(*int32)(unsafe.Add(mBase, uint32(v1193)+12)) = int32(1)
	goto L189
L194:
	;
	goto L188
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v1258
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	v1263 = int32(0)
	v1265 = F_pull_up_subqueries_recurse(m, v47, v1262, v1263, v1263)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1267)+60)) = v1265
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+144))
	if v1269 != 0 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v1270 = m.G0
	v1272 = v1270 - int32(16)
	m.G0 = v1272
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+322)))
	if v1274 != 0 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v1453 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+324)) = v1453
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+316)) = uint16(v1453)
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	if v1458 != 0 {
		goto L218
	} else {
		goto L219
	}
L200:
	;
	m.G0 = v1272 + int32(16)
	goto L199
L201:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+144))
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1276)+20))
	v1278 = F_is_simple_union_all_recurse(m, v1276, v1275, v1277)
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	if v1278 == int32(0) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v1284 = v1276
	goto L204
L204:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	if v1323 != 0 {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+52))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1327)+12))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1323)+4))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1328+v1329<<(uint(int32(2))%32)-int32(4))))
	v1336 = F_copyObjectImpl(m, v1335)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L210
	}
L206:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1323)))
	if v1324 == int32(142) {
		v1284 = v1323
		goto L204
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	goto L205
L209:
	;
	goto L208
L210:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+52))
	v1339 = F_lappend(m, v1338, v1336)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+52)) = v1339
	if v1339 != 0 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1339)+4))
	v1344 = v1342
	goto L214
L213:
	;
	v1344 = int32(0)
	goto L214
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1323)+4)) = v1344
	v1346 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1335)+20)) = uint8(v1346)
	v1349 = F_palloc0(m, int32(8))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1349)+4)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v1349))) = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v1272)+8)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v1272)+12)) = v1349
	v1359 = F_list_make1_impl(m, int32(1), v1272+int32(8))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1275)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1361)+4)) = v1359
	v1363 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1275)+144)) = v1363
	F_pull_up_union_leaf_queries(m, v1276, v47, v1329, v1275, v1363)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	goto L200
L218:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+4))
	if int32(0) < v1459 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	v1630 = v1453
	v1640 = v7
	v1642 = v7
	goto L220
L220:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+32))
	if v1669 != 0 {
		goto L252
	} else {
		goto L253
	}
L221:
	;
	v1464 = v1453
	v1474 = v7
	v1476 = v7
	goto L224
L222:
	;
	v1598 = v7
	v1600 = v7
	goto L223
L223:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	v1630 = v1627
	v1640 = v1598
	v1642 = v1600
	goto L220
L224:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+12))
	v1506 = v1503 + v1464<<(uint(int32(2))%32)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1506)))
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+12))
	switch v1508 {
	case 0:
		goto L230
	default:
		v1567 = v1474
		v1568 = v1476
		goto L226
	case 2:
		goto L229
	case 8:
		goto L228
	case 9:
		goto L227
	}
L225:
	;
	v1598 = v1567
	v1600 = v1568
	goto L223
L226:
	;
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507)+124)))
	if v1569 == int32(1) {
		goto L240
	} else {
		goto L241
	}
L227:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(v1556)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+324)) = (v1506-v1557)>>(uint(int32(2))%32) + int32(1)
	v1567 = v1474
	v1568 = v1476
	goto L226
L228:
	;
	v1567 = int32(1)
	v1568 = v1476
	goto L226
L229:
	;
	v1545 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)) = uint8(v1545)
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+44))
	v1567 = v1474
	v1568 = base.B2i32(v1545<<(uint(v1548)%32)&int32(174) != int32(0)) | v1476
	goto L226
L230:
	;
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507)+20)))
	if v1509 != int32(1) {
		v1567 = v1474
		v1568 = v1476
		goto L226
	} else {
		goto L231
	}
L231:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+16))
	v1513 = m.G0
	v1515 = v1513 - int32(16)
	m.G0 = v1515
	v1518 = F_SearchSysCache1(m, int32(57), v1512)
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	if v1518 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1518)+16))
	v1536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535)+22)))
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535+v1536)+126)))
	F_ReleaseCatCache(m, v1518)
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L1
	} else {
		goto L239
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1515))) = v1512
	F_errmsg_internal(m, int32(_a_F_subquery_planner_8), v1515)
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_9), int32(362), int32(_a_F_subquery_planner_10))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	m.G0 = v1515 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v1507)+20)) = uint8(v1538)
	v1567 = v1474
	v1568 = v1476
	goto L226
L240:
	;
	v1572 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+317)) = uint8(v1572)
	goto L242
L241:
	;
	goto L242
L242:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1507)+128))
	if v1574 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v47)+312))
	v1576 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if base.Ui32(v1576) < base.Ui32(v1575) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	v1583 = v1464 + int32(1)
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+4))
	if v1583 < v1584 {
		v1464 = v1583
		v1474 = v1567
		v1476 = v1568
		goto L224
	} else {
		goto L249
	}
L246:
	;
	v1578 = v1575
	goto L248
L247:
	;
	v1578 = v1576
	goto L248
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+312)) = v1578
	goto L245
L249:
	;
	goto L225
L250:
	;
	v1800 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+140))
	if v1801 != 0 {
		goto L272
	} else {
		goto L273
	}
L251:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+4))
	if v1686 <= int32(0) {
		goto L250
	} else {
		goto L258
	}
L252:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1630)+12))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1670+v1669<<(uint(int32(2))%32)-int32(4))))
	v1677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1676)+20)))
	if v1677 != 0 {
		v1685 = v1630
		goto L251
	} else {
		goto L255
	}
L253:
	;
	v1682 = v1630
	goto L254
L254:
	;
	if v1682 == int32(0) {
		goto L250
	} else {
		goto L257
	}
L255:
	;
	v1678 = F_bms_make_singleton(m, v1669)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+124)) = v1678
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	v1682 = v1681
	goto L254
L257:
	;
	v1685 = v1682
	goto L251
L258:
	;
	v1691 = int32(0)
	goto L259
L259:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+12))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1731+v1691<<(uint(int32(2))%32))))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)+28))
	if v1736 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L260:
	;
	goto L250
L261:
	;
	v1756 = v1691 + int32(1)
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+4))
	if v1756 < v1757 {
		v1691 = v1756
		goto L259
	} else {
		goto L269
	}
L262:
	;
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735)+21)))
	if v1739 != int32(118) {
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+56))
	v1743 = F_getRTEPermissionInfo(m, v1742, v1735)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1745 = F_ExecCheckOneRelPerms(m, v1743)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L1
	} else {
		goto L265
	}
L265:
	;
	if v1745 != 0 {
		goto L261
	} else {
		goto L266
	}
L266:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1743)+4))
	v1750 = F_get_rel_name(m, v1749)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_aclcheck_error(m, int32(1), int32(51), v1750)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	goto L261
L269:
	;
	goto L260
L270:
	;
	v2193 = int32(0)
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+112))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+318)) = uint8(base.B2i32(v2194 != v2193))
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+76))
	if v2199 == v2193 {
		v2226 = v2193
		goto L325
	} else {
		goto L326
	}
L271:
	;
	v1818 = int32(0)
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+60))
	v1822 = F_get_relids_in_jointree(m, v1819, v1818, v1818)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L277
	}
L272:
	;
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v1801)+12))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v1802)))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1803)+8))
	F_CheckSelectLocking(m, v1800, v1804)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+4))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v1807))|base.B2i32(int32(1)<<(uint(v1807)%32)&int32(52) == int32(0)) != 0 {
		goto L270
	} else {
		goto L276
	}
L275:
	;
	goto L271
L276:
	;
	goto L271
L277:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+32))
	if v1824 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1825 = F_bms_del_member(m, v1822, v1824)
	mBase = m.M
	v1826 = m.ExcPending
	if v1826 != 0 {
		goto L1
	} else {
		goto L281
	}
L279:
	;
	v1827 = v1822
	goto L280
L280:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+140))
	if v1828 == int32(0) {
		v1980 = v1818
		v1983 = v1827
		goto L282
	} else {
		goto L283
	}
L281:
	;
	v1827 = v1825
	goto L280
L282:
	;
	v2006 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+52))
	if v2006 == int32(0) {
		v2125 = v1980
		goto L307
	} else {
		goto L308
	}
L283:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1828)+4))
	if v1831 <= int32(0) {
		v1980 = v1818
		v1983 = v1827
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1836 = v1831
	v1837 = int32(0)
	v1850 = v1818
	v1853 = v1827
	goto L285
L285:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1800)+52))
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1876)+12))
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(v1828)+12))
	v1879 = int32(2)
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1878+v1837<<(uint(v1879)%32))))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+4))
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1877+v1883<<(uint(v1879)%32)-int32(4))))
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+12))
	if v1890 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L286:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L304
	}
L287:
	;
	goto L286
L288:
	;
	v1893 = F_bms_del_member(m, v1853, v1883)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	v1944 = v1836
	v1946 = v1850
	v1947 = v1853
	goto L290
L290:
	;
	v1950 = v1837 + int32(1)
	if v1950 < v1944 {
		v1836 = v1944
		v1837 = v1950
		v1850 = v1946
		v1853 = v1947
		goto L285
	} else {
		goto L303
	}
L291:
	;
	v1896 = F_palloc0(m, int32(36))
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1896))) = int32(374)
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+4)) = v1900
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+8)) = v1900
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1903)+72))
	v1906 = v1904 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1903)+72)) = v1906
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+12)) = v1906
	v1909 = int32(5)
	v1910 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+12))
	if v1910 != 0 {
		v1928 = v1909
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+16)) = v1928
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+20)) = int32(1) << (uint(v1928) % 32)
	v1934 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+24)) = v1934
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+12))
	v1937 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1896)+32)) = uint8(v1937)
	*(*int32)(unsafe.Add(mBase, uint32(v1896)+28)) = v1936
	v1940 = F_lappend(m, v1850, v1896)
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L1
	} else {
		goto L302
	}
L294:
	;
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+8))
	v1912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1889)+21)))
	if v1912 == int32(102) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1889)+16))
	v1916 = F_GetFdwRoutineByRelId(m, v1915)
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
		goto L1
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	if base.Ui32(int32(5)) <= base.Ui32(v1911) {
		goto L287
	} else {
		goto L301
	}
L298:
	;
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1916)+104))
	if v1918 == int32(0) {
		v1928 = v1909
		goto L293
	} else {
		goto L299
	}
L299:
	;
	v1921 = m.T0[v1918].(func(*base.Module, int32, int32) int32)(m, v1889, v1911)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L1
	} else {
		goto L300
	}
L300:
	;
	v1928 = v1921
	goto L293
L301:
	;
	v1928 = int32(4) - v1911
	goto L293
L302:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1828)+4))
	v1944 = v1942
	v1946 = v1940
	v1947 = v1893
	goto L290
L303:
	;
	v1980 = v1946
	v1983 = v1947
	goto L282
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v1911
	F_errmsg_internal(m, int32(_a_F_subquery_planner_11), v44)
	mBase = m.M
	v1959 = m.ExcPending
	if v1959 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(2554), int32(_a_F_subquery_planner_13))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v2125
	goto L270
L308:
	;
	v2009 = int32(0)
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v2006)+4))
	if v2010 <= v2009 {
		v2125 = v1980
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v2014 = v2009
	v2028 = v1980
	goto L310
L310:
	;
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2006)+12))
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v2054+v2014<<(uint(int32(2))%32))))
	v2060 = v2014 + int32(1)
	v2061 = F_bms_is_member(m, v2060, v1983)
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L1
	} else {
		goto L312
	}
L311:
	;
	v2125 = v2107
	goto L307
L312:
	;
	if v2061 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v2064 = F_palloc0(m, int32(36))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L1
	} else {
		goto L316
	}
L314:
	;
	v2107 = v2028
	goto L315
L315:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2006)+4))
	if v2060 < v2108 {
		v2014 = v2060
		v2028 = v2107
		goto L310
	} else {
		goto L324
	}
L316:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+8)) = v2060
	*(*int32)(unsafe.Add(mBase, uint32(v2064))) = int32(374)
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+4)) = v2060
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2070)+72))
	v2073 = v2071 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v2070)+72)) = v2073
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+12)) = v2073
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v2058)+12))
	if v2077 != 0 {
		v2092 = int32(5)
		goto L317
	} else {
		goto L318
	}
L317:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2064)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+16)) = v2092
	v2097 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2064)+32)) = uint8(v2097)
	*(*int32)(unsafe.Add(mBase, uint32(v2064)+20)) = int32(1) << (uint(v2092) % 32)
	v2102 = F_lappend(m, v2028, v2064)
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L1
	} else {
		goto L323
	}
L318:
	;
	v2079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2058)+21)))
	if v2079 != int32(102) {
		v2092 = int32(4)
		goto L317
	} else {
		goto L319
	}
L319:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(v2058)+16))
	v2084 = F_GetFdwRoutineByRelId(m, v2083)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v2084)+104))
	if v2086 == int32(0) {
		v2092 = int32(5)
		goto L317
	} else {
		goto L321
	}
L321:
	;
	v2090 = m.T0[v2086].(func(*base.Module, int32, int32) int32)(m, v2058, int32(0))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	v2092 = v2090
	goto L317
L323:
	;
	v2107 = v2102
	goto L315
L324:
	;
	goto L311
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+76)) = v2226
	v2228 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+152))
	if v2228 == int32(0) {
		v2311 = v2193
		goto L339
	} else {
		goto L340
	}
L326:
	;
	v2202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2202 == int32(1) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2206 = F_flatten_join_alias_vars(m, v47, v2205, v2199)
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L330
	}
L328:
	;
	v2208 = v2199
	goto L329
L329:
	;
	v2209 = F_eval_const_expressions(m, v47, v2208)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L1
	} else {
		goto L331
	}
L330:
	;
	v2208 = v2206
	goto L329
L331:
	;
	F_convert_saop_to_hashed_saop(m, v2209)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2213)+39)))
	if v2214 == int32(1) {
		goto L333
	} else {
		goto L334
	}
L333:
	;
	v2218 = F_SS_process_sublinks(m, v47, v2209, int32(0))
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L1
	} else {
		goto L336
	}
L334:
	;
	v2220 = v2209
	goto L335
L335:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2221) < base.Ui32(int32(2)) {
		v2226 = v2220
		goto L325
	} else {
		goto L337
	}
L336:
	;
	v2220 = v2218
	goto L335
L337:
	;
	v2224 = F_SS_replace_correlation_vars(m, v47, v2220)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v2226 = v2224
	goto L325
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+152)) = v2311
	v2335 = int32(0)
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+96))
	if v2336 == v2335 {
		v2363 = v2335
		goto L350
	} else {
		goto L351
	}
L340:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+4))
	if v2231 <= int32(0) {
		v2311 = v2193
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v2237 = int32(0)
	v2253 = v2193
	goto L342
L342:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+12))
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2276+v2237<<(uint(int32(2))%32))))
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2280)+16))
	v2283 = F_preprocess_expression(m, v47, v2281, int32(0))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L1
	} else {
		goto L344
	}
L343:
	;
	v2311 = v2288
	goto L339
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2280)+16)) = v2283
	if v2283 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v2286 = F_lappend(m, v2253, v2280)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	v2288 = v2253
	goto L347
L347:
	;
	v2290 = v2237 + int32(1)
	v2291 = *(*int32)(unsafe.Add(mBase, uint32(v2228)+4))
	if v2290 < v2291 {
		v2237 = v2290
		v2253 = v2288
		goto L342
	} else {
		goto L349
	}
L348:
	;
	v2288 = v2286
	goto L347
L349:
	;
	goto L343
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+96)) = v2363
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	F_preprocess_qual_conditions(m, v47, v2365)
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L1
	} else {
		goto L364
	}
L351:
	;
	v2339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2339 == int32(1) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2343 = F_flatten_join_alias_vars(m, v47, v2342, v2336)
	mBase = m.M
	v2344 = m.ExcPending
	if v2344 != 0 {
		goto L1
	} else {
		goto L355
	}
L353:
	;
	v2345 = v2336
	goto L354
L354:
	;
	v2346 = F_eval_const_expressions(m, v47, v2345)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L1
	} else {
		goto L356
	}
L355:
	;
	v2345 = v2343
	goto L354
L356:
	;
	F_convert_saop_to_hashed_saop(m, v2346)
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L1
	} else {
		goto L357
	}
L357:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2350)+39)))
	if v2351 == int32(1) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v2355 = F_SS_process_sublinks(m, v47, v2346, int32(0))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L361
	}
L359:
	;
	v2357 = v2346
	goto L360
L360:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2358) < base.Ui32(int32(2)) {
		v2363 = v2357
		goto L350
	} else {
		goto L362
	}
L361:
	;
	v2357 = v2355
	goto L360
L362:
	;
	v2361 = F_SS_replace_correlation_vars(m, v47, v2357)
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L1
	} else {
		goto L363
	}
L363:
	;
	v2363 = v2361
	goto L350
L364:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+112))
	v2370 = F_preprocess_expression(m, v47, v2368, int32(0))
	mBase = m.M
	v2371 = m.ExcPending
	if v2371 != 0 {
		goto L1
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+112)) = v2370
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+116))
	if v2373 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v2527 = int32(0)
	v2529 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+128))
	if v2529 == v2527 {
		v2554 = v2527
		goto L400
	} else {
		goto L401
	}
L367:
	;
	v2376 = int32(0)
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+4))
	if v2377 <= v2376 {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v2383 = v2376
	goto L369
L369:
	;
	v2421 = int32(0)
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+12))
	v2426 = *(*int32)(unsafe.Add(mBase, uint32(v2422+v2383<<(uint(int32(2))%32))))
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2426)+24))
	if v2427 == v2421 {
		v2452 = v2421
		goto L371
	} else {
		goto L372
	}
L370:
	;
	goto L366
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2426)+24)) = v2452
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2426)+28))
	if v2454 == int32(0) {
		goto L385
	} else {
		goto L386
	}
L372:
	;
	v2430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2430 == int32(1) {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v2433 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2434 = F_flatten_join_alias_vars(m, v47, v2433, v2427)
	mBase = m.M
	v2435 = m.ExcPending
	if v2435 != 0 {
		goto L1
	} else {
		goto L376
	}
L374:
	;
	v2436 = v2427
	goto L375
L375:
	;
	v2437 = F_eval_const_expressions(m, v47, v2436)
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L377
	}
L376:
	;
	v2436 = v2434
	goto L375
L377:
	;
	v2439 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2439)+39)))
	if v2440 == int32(1) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v2444 = F_SS_process_sublinks(m, v47, v2437, int32(0))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L1
	} else {
		goto L381
	}
L379:
	;
	v2446 = v2437
	goto L380
L380:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2447) < base.Ui32(int32(2)) {
		v2452 = v2446
		goto L371
	} else {
		goto L382
	}
L381:
	;
	v2446 = v2444
	goto L380
L382:
	;
	v2450 = F_SS_replace_correlation_vars(m, v47, v2446)
	mBase = m.M
	v2451 = m.ExcPending
	if v2451 != 0 {
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v2452 = v2450
	goto L371
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2426)+28)) = v2480
	v2483 = v2383 + int32(1)
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(v2373)+4))
	if v2483 < v2484 {
		v2383 = v2483
		goto L369
	} else {
		goto L399
	}
L385:
	;
	v2480 = int32(0)
	goto L384
L386:
	;
	goto L387
L387:
	;
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2458 == int32(1) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2462 = F_flatten_join_alias_vars(m, v47, v2461, v2454)
	mBase = m.M
	v2463 = m.ExcPending
	if v2463 != 0 {
		goto L1
	} else {
		goto L391
	}
L389:
	;
	v2464 = v2454
	goto L390
L390:
	;
	v2465 = F_eval_const_expressions(m, v47, v2464)
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L1
	} else {
		goto L392
	}
L391:
	;
	v2464 = v2462
	goto L390
L392:
	;
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2467)+39)))
	if v2468 == int32(1) {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v2472 = F_SS_process_sublinks(m, v47, v2465, int32(0))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L1
	} else {
		goto L396
	}
L394:
	;
	v2474 = v2465
	goto L395
L395:
	;
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2475) < base.Ui32(int32(2)) {
		v2480 = v2474
		goto L384
	} else {
		goto L397
	}
L396:
	;
	v2474 = v2472
	goto L395
L397:
	;
	v2478 = F_SS_replace_correlation_vars(m, v47, v2474)
	mBase = m.M
	v2479 = m.ExcPending
	if v2479 != 0 {
		goto L1
	} else {
		goto L398
	}
L398:
	;
	v2480 = v2478
	goto L384
L399:
	;
	goto L370
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+128)) = v2554
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+132))
	if v2556 == int32(0) {
		v2581 = v2527
		goto L413
	} else {
		goto L414
	}
L401:
	;
	v2532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2532 == int32(1) {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2536 = F_flatten_join_alias_vars(m, v47, v2535, v2529)
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L1
	} else {
		goto L405
	}
L403:
	;
	v2538 = v2529
	goto L404
L404:
	;
	v2539 = F_eval_const_expressions(m, v47, v2538)
	mBase = m.M
	v2540 = m.ExcPending
	if v2540 != 0 {
		goto L1
	} else {
		goto L406
	}
L405:
	;
	v2538 = v2536
	goto L404
L406:
	;
	v2541 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2541)+39)))
	if v2542 == int32(1) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v2546 = F_SS_process_sublinks(m, v47, v2539, int32(0))
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L1
	} else {
		goto L410
	}
L408:
	;
	v2548 = v2539
	goto L409
L409:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2549) < base.Ui32(int32(2)) {
		v2554 = v2548
		goto L400
	} else {
		goto L411
	}
L410:
	;
	v2548 = v2546
	goto L409
L411:
	;
	v2552 = F_SS_replace_correlation_vars(m, v47, v2548)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L412
	}
L412:
	;
	v2554 = v2552
	goto L400
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+132)) = v2581
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	if v2583 != 0 {
		goto L426
	} else {
		goto L427
	}
L414:
	;
	v2559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2559 == int32(1) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2563 = F_flatten_join_alias_vars(m, v47, v2562, v2556)
	mBase = m.M
	v2564 = m.ExcPending
	if v2564 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	v2565 = v2556
	goto L417
L417:
	;
	v2566 = F_eval_const_expressions(m, v47, v2565)
	mBase = m.M
	v2567 = m.ExcPending
	if v2567 != 0 {
		goto L1
	} else {
		goto L419
	}
L418:
	;
	v2565 = v2563
	goto L417
L419:
	;
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2568)+39)))
	if v2569 == int32(1) {
		goto L420
	} else {
		goto L421
	}
L420:
	;
	v2573 = F_SS_process_sublinks(m, v47, v2566, int32(0))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L1
	} else {
		goto L423
	}
L421:
	;
	v2575 = v2566
	goto L422
L422:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2576) < base.Ui32(int32(2)) {
		v2581 = v2575
		goto L413
	} else {
		goto L424
	}
L423:
	;
	v2575 = v2573
	goto L422
L424:
	;
	v2579 = F_SS_replace_correlation_vars(m, v47, v2575)
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	v2581 = v2579
	goto L413
L426:
	;
	v2584 = int32(0)
	v2586 = *(*int32)(unsafe.Add(mBase, uint32(v2583)+8))
	if v2586 == v2584 {
		v2611 = v2584
		goto L429
	} else {
		goto L430
	}
L427:
	;
	goto L428
L428:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+64))
	if v2662 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L429:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2612)+8)) = v2611
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v2614)+12))
	v2617 = F_preprocess_expression(m, v47, v2615, int32(0))
	mBase = m.M
	v2618 = m.ExcPending
	if v2618 != 0 {
		goto L1
	} else {
		goto L442
	}
L430:
	;
	v2589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2589 == int32(1) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2593 = F_flatten_join_alias_vars(m, v47, v2592, v2586)
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L1
	} else {
		goto L434
	}
L432:
	;
	v2595 = v2586
	goto L433
L433:
	;
	v2596 = F_eval_const_expressions(m, v47, v2595)
	mBase = m.M
	v2597 = m.ExcPending
	if v2597 != 0 {
		goto L1
	} else {
		goto L435
	}
L434:
	;
	v2595 = v2593
	goto L433
L435:
	;
	v2598 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2598)+39)))
	if v2599 == int32(1) {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	v2603 = F_SS_process_sublinks(m, v47, v2596, int32(0))
	mBase = m.M
	v2604 = m.ExcPending
	if v2604 != 0 {
		goto L1
	} else {
		goto L439
	}
L437:
	;
	v2605 = v2596
	goto L438
L438:
	;
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2606) < base.Ui32(int32(2)) {
		v2611 = v2605
		goto L429
	} else {
		goto L440
	}
L439:
	;
	v2605 = v2603
	goto L438
L440:
	;
	v2609 = F_SS_replace_correlation_vars(m, v47, v2605)
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v2611 = v2609
	goto L429
L442:
	;
	v2619 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2619)+12)) = v2617
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2621)+20))
	if v2622 == int32(0) {
		v2649 = v2584
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v2650 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2650)+20)) = v2649
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	v2653 = *(*int32)(unsafe.Add(mBase, uint32(v2652)+24))
	v2655 = F_preprocess_expression(m, v47, v2653, int32(0))
	mBase = m.M
	v2656 = m.ExcPending
	if v2656 != 0 {
		goto L1
	} else {
		goto L457
	}
L444:
	;
	v2625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2625 == int32(1) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2629 = F_flatten_join_alias_vars(m, v47, v2628, v2622)
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L1
	} else {
		goto L448
	}
L446:
	;
	v2631 = v2622
	goto L447
L447:
	;
	v2632 = F_eval_const_expressions(m, v47, v2631)
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L1
	} else {
		goto L449
	}
L448:
	;
	v2631 = v2629
	goto L447
L449:
	;
	F_convert_saop_to_hashed_saop(m, v2632)
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2636)+39)))
	if v2637 == int32(1) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v2641 = F_SS_process_sublinks(m, v47, v2632, int32(0))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L1
	} else {
		goto L454
	}
L452:
	;
	v2643 = v2632
	goto L453
L453:
	;
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2644) < base.Ui32(int32(2)) {
		v2649 = v2643
		goto L443
	} else {
		goto L455
	}
L454:
	;
	v2643 = v2641
	goto L453
L455:
	;
	v2647 = F_SS_replace_correlation_vars(m, v47, v2643)
	mBase = m.M
	v2648 = m.ExcPending
	if v2648 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	v2649 = v2647
	goto L443
L457:
	;
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2657)+24)) = v2655
	goto L428
L458:
	;
	v2795 = int32(0)
	v2796 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+72))
	v2798 = F_preprocess_expression(m, v47, v2796, v2795)
	mBase = m.M
	v2799 = m.ExcPending
	if v2799 != 0 {
		goto L1
	} else {
		goto L479
	}
L459:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2662)+4))
	if v2665 <= int32(0) {
		goto L458
	} else {
		goto L460
	}
L460:
	;
	v2672 = int32(0)
	goto L461
L461:
	;
	v2710 = int32(0)
	v2711 = *(*int32)(unsafe.Add(mBase, uint32(v2662)+12))
	v2715 = *(*int32)(unsafe.Add(mBase, uint32(v2711+v2672<<(uint(int32(2))%32))))
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2715)+20))
	if v2716 == v2710 {
		v2743 = v2710
		goto L463
	} else {
		goto L464
	}
L462:
	;
	goto L458
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2715)+20)) = v2743
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v2715)+16))
	v2747 = F_preprocess_expression(m, v47, v2745, int32(0))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L1
	} else {
		goto L477
	}
L464:
	;
	v2719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2719 == int32(1) {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2723 = F_flatten_join_alias_vars(m, v47, v2722, v2716)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L1
	} else {
		goto L468
	}
L466:
	;
	v2725 = v2716
	goto L467
L467:
	;
	v2726 = F_eval_const_expressions(m, v47, v2725)
	mBase = m.M
	v2727 = m.ExcPending
	if v2727 != 0 {
		goto L1
	} else {
		goto L469
	}
L468:
	;
	v2725 = v2723
	goto L467
L469:
	;
	F_convert_saop_to_hashed_saop(m, v2726)
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2730)+39)))
	if v2731 == int32(1) {
		goto L471
	} else {
		goto L472
	}
L471:
	;
	v2735 = F_SS_process_sublinks(m, v47, v2726, int32(0))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L1
	} else {
		goto L474
	}
L472:
	;
	v2737 = v2726
	goto L473
L473:
	;
	v2738 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2738) < base.Ui32(int32(2)) {
		v2743 = v2737
		goto L463
	} else {
		goto L475
	}
L474:
	;
	v2737 = v2735
	goto L473
L475:
	;
	v2741 = F_SS_replace_correlation_vars(m, v47, v2737)
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v2743 = v2741
	goto L463
L477:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2715)+16)) = v2747
	v2751 = v2672 + int32(1)
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v2662)+4))
	if v2751 < v2752 {
		v2672 = v2751
		goto L461
	} else {
		goto L478
	}
L478:
	;
	goto L462
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+72)) = v2798
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	if v2801 == int32(0) {
		v2826 = v2795
		goto L480
	} else {
		goto L481
	}
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v2826
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	if v2828 == int32(0) {
		goto L493
	} else {
		goto L494
	}
L481:
	;
	v2804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2804 == int32(1) {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v2807 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2808 = F_flatten_join_alias_vars(m, v47, v2807, v2801)
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L1
	} else {
		goto L485
	}
L483:
	;
	v2810 = v2801
	goto L484
L484:
	;
	v2811 = F_eval_const_expressions(m, v47, v2810)
	mBase = m.M
	v2812 = m.ExcPending
	if v2812 != 0 {
		goto L1
	} else {
		goto L486
	}
L485:
	;
	v2810 = v2808
	goto L484
L486:
	;
	v2813 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2813)+39)))
	if v2814 == int32(1) {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2818 = F_SS_process_sublinks(m, v47, v2811, int32(0))
	mBase = m.M
	v2819 = m.ExcPending
	if v2819 != 0 {
		goto L1
	} else {
		goto L490
	}
L488:
	;
	v2820 = v2811
	goto L489
L489:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2821) < base.Ui32(int32(2)) {
		v2826 = v2820
		goto L480
	} else {
		goto L491
	}
L490:
	;
	v2820 = v2818
	goto L489
L491:
	;
	v2824 = F_SS_replace_correlation_vars(m, v47, v2820)
	mBase = m.M
	v2825 = m.ExcPending
	if v2825 != 0 {
		goto L1
	} else {
		goto L492
	}
L492:
	;
	v2826 = v2824
	goto L480
L493:
	;
	v3114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v3114 == int32(0) {
		goto L553
	} else {
		goto L554
	}
L494:
	;
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+4))
	if v2831 <= int32(0) {
		goto L493
	} else {
		goto L495
	}
L495:
	;
	v2850 = int32(0)
	goto L496
L496:
	;
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+12))
	v2880 = *(*int32)(unsafe.Add(mBase, uint32(v2876+v2850<<(uint(int32(2))%32))))
	v2881 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+12))
	switch v2881 {
	case 0:
		goto L504
	case 1:
		goto L503
	default:
		goto L498
	case 3:
		goto L502
	case 4:
		goto L501
	case 5:
		goto L500
	case 9:
		goto L499
	}
L497:
	;
	goto L493
L498:
	;
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+128))
	if v2967 == int32(0) {
		goto L545
	} else {
		goto L546
	}
L499:
	;
	v2937 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+120))
	if v2937 == int32(0) {
		goto L531
	} else {
		goto L532
	}
L500:
	;
	v2929 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+80))
	v2932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+124)))
	if v2932 != 0 {
		goto L526
	} else {
		goto L527
	}
L501:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+76))
	v2924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+124)))
	if v2924 != 0 {
		goto L522
	} else {
		goto L523
	}
L502:
	;
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+68))
	v2916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+124)))
	if v2916 != 0 {
		goto L518
	} else {
		goto L519
	}
L503:
	;
	v2902 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2880)+124)))
	if v2902 != int32(1) {
		goto L498
	} else {
		goto L515
	}
L504:
	;
	v2882 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+32))
	if v2882 == int32(0) {
		goto L498
	} else {
		goto L505
	}
L505:
	;
	v2885 = F_eval_const_expressions(m, v47, v2882)
	mBase = m.M
	v2886 = m.ExcPending
	if v2886 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2887)+39)))
	if v2888 == int32(1) {
		goto L507
	} else {
		goto L508
	}
L507:
	;
	v2892 = F_SS_process_sublinks(m, v47, v2885, int32(0))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L510
	}
L508:
	;
	v2894 = v2885
	goto L509
L509:
	;
	v2895 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(int32(2)) <= base.Ui32(v2895) {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	v2894 = v2892
	goto L509
L511:
	;
	v2898 = F_SS_replace_correlation_vars(m, v47, v2894)
	mBase = m.M
	v2899 = m.ExcPending
	if v2899 != 0 {
		goto L1
	} else {
		goto L514
	}
L512:
	;
	v2900 = v2894
	goto L513
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+32)) = v2900
	goto L498
L514:
	;
	v2900 = v2898
	goto L513
L515:
	;
	v2905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2905 != int32(1) {
		goto L498
	} else {
		goto L516
	}
L516:
	;
	v2908 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(v2880)+36))
	v2910 = F_flatten_join_alias_vars(m, v47, v2908, v2909)
	mBase = m.M
	v2911 = m.ExcPending
	if v2911 != 0 {
		goto L1
	} else {
		goto L517
	}
L517:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+36)) = v2910
	goto L498
L518:
	;
	v2917 = int32(3)
	goto L520
L519:
	;
	v2917 = int32(2)
	goto L520
L520:
	;
	v2918 = F_preprocess_expression(m, v47, v2913, v2917)
	mBase = m.M
	v2919 = m.ExcPending
	if v2919 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+68)) = v2918
	goto L498
L522:
	;
	v2925 = int32(12)
	goto L524
L523:
	;
	v2925 = int32(11)
	goto L524
L524:
	;
	v2926 = F_preprocess_expression(m, v47, v2921, v2925)
	mBase = m.M
	v2927 = m.ExcPending
	if v2927 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+76)) = v2926
	goto L498
L526:
	;
	v2933 = int32(5)
	goto L528
L527:
	;
	v2933 = int32(4)
	goto L528
L528:
	;
	v2934 = F_preprocess_expression(m, v47, v2929, v2933)
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+80)) = v2934
	goto L498
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2880)+120)) = v2963
	goto L498
L531:
	;
	v2963 = int32(0)
	goto L530
L532:
	;
	goto L533
L533:
	;
	v2941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+316)))
	if v2941 == int32(1) {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2945 = F_flatten_join_alias_vars(m, v47, v2944, v2937)
	mBase = m.M
	v2946 = m.ExcPending
	if v2946 != 0 {
		goto L1
	} else {
		goto L537
	}
L535:
	;
	v2947 = v2937
	goto L536
L536:
	;
	v2948 = F_eval_const_expressions(m, v47, v2947)
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L1
	} else {
		goto L538
	}
L537:
	;
	v2947 = v2945
	goto L536
L538:
	;
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2950)+39)))
	if v2951 == int32(1) {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v2955 = F_SS_process_sublinks(m, v47, v2948, int32(0))
	mBase = m.M
	v2956 = m.ExcPending
	if v2956 != 0 {
		goto L1
	} else {
		goto L542
	}
L540:
	;
	v2957 = v2948
	goto L541
L541:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(v2958) < base.Ui32(int32(2)) {
		v2963 = v2957
		goto L530
	} else {
		goto L543
	}
L542:
	;
	v2957 = v2955
	goto L541
L543:
	;
	v2961 = F_SS_replace_correlation_vars(m, v47, v2957)
	mBase = m.M
	v2962 = m.ExcPending
	if v2962 != 0 {
		goto L1
	} else {
		goto L544
	}
L544:
	;
	v2963 = v2961
	goto L530
L545:
	;
	v3070 = v2850 + int32(1)
	v3071 = *(*int32)(unsafe.Add(mBase, uint32(v2828)+4))
	if v3070 < v3071 {
		v2850 = v3070
		goto L496
	} else {
		goto L552
	}
L546:
	;
	v2970 = int32(0)
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2967)+4))
	if v2971 <= v2970 {
		goto L545
	} else {
		goto L547
	}
L547:
	;
	v2975 = v2970
	goto L548
L548:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, uint32(v2967)+12))
	v3018 = v3015 + v2975<<(uint(int32(2))%32)
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v3018)))
	v3021 = F_preprocess_expression(m, v47, v3019, int32(0))
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L1
	} else {
		goto L550
	}
L549:
	;
	goto L545
L550:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3018))) = v3021
	v3025 = v2975 + int32(1)
	v3026 = *(*int32)(unsafe.Add(mBase, uint32(v2967)+4))
	if v3025 < v3026 {
		v2975 = v3025
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	goto L497
L553:
	;
	v3217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+45)))
	if v3217 == int32(1) {
		goto L560
	} else {
		goto L561
	}
L554:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+52))
	if v3117 == int32(0) {
		goto L553
	} else {
		goto L555
	}
L555:
	;
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+4))
	if v3120 <= int32(0) {
		goto L553
	} else {
		goto L556
	}
L556:
	;
	v3125 = int32(0)
	goto L557
L557:
	;
	v3165 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+12))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3165+v3125<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3169)+52)) = int32(0)
	v3173 = v3125 + int32(1)
	v3174 = *(*int32)(unsafe.Add(mBase, uint32(v3117)+4))
	if v3173 < v3174 {
		v3125 = v3173
		goto L557
	} else {
		goto L559
	}
L558:
	;
	goto L553
L559:
	;
	goto L558
L560:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+76))
	v3222 = F_flatten_group_exprs(m, v47, v3220, v3221)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L563
	}
L561:
	;
	goto L562
L562:
	;
	v3230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+38)))
	if v3230 == int32(1) {
		goto L565
	} else {
		goto L566
	}
L563:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+76)) = v3222
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+112))
	v3227 = F_flatten_group_exprs(m, v47, v3225, v3226)
	mBase = m.M
	v3228 = m.ExcPending
	if v3228 != 0 {
		goto L1
	} else {
		goto L564
	}
L564:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+112)) = v3227
	goto L562
L565:
	;
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+76))
	v3234 = F_expression_returns_set(m, v3233)
	mBase = m.M
	v3235 = m.ExcPending
	if v3235 != 0 {
		goto L1
	} else {
		goto L568
	}
L566:
	;
	goto L567
L567:
	;
	v3237 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+108))
	if v3237 != 0 {
		goto L569
	} else {
		goto L570
	}
L568:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1258)+38)) = uint8(v3234)
	goto L567
L569:
	;
	v3238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+104)))
	v3240 = F_expand_grouping_sets(m, v3237, v3238, int32(-1))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		goto L1
	} else {
		goto L572
	}
L570:
	;
	goto L571
L571:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+112))
	if v3243 == int32(0) {
		goto L574
	} else {
		goto L575
	}
L572:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+108)) = v3240
	goto L571
L573:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1258)+112)) = v3372
	if v1642&int32(1) != 0 {
		goto L607
	} else {
		goto L608
	}
L574:
	;
	v3372 = int32(0)
	goto L573
L575:
	;
	goto L576
L576:
	;
	v3247 = int32(0)
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+4))
	if v3248 <= v3247 {
		v3372 = v3247
		goto L573
	} else {
		goto L577
	}
L577:
	;
	v3253 = int32(0)
	v3267 = v3247
	goto L578
L578:
	;
	v3293 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+12))
	v3297 = *(*int32)(unsafe.Add(mBase, uint32(v3293+v3253<<(uint(int32(2))%32))))
	v3298 = F_contain_agg_clause(m, v3297)
	mBase = m.M
	v3299 = m.ExcPending
	if v3299 != 0 {
		goto L1
	} else {
		goto L582
	}
L579:
	;
	v3372 = v3352
	goto L573
L580:
	;
	v3354 = v3253 + int32(1)
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3243)+4))
	if v3354 < v3355 {
		v3253 = v3354
		v3267 = v3352
		goto L578
	} else {
		goto L604
	}
L581:
	;
	v3348 = F_lappend(m, v3267, v3297)
	mBase = m.M
	v3349 = m.ExcPending
	if v3349 != 0 {
		goto L1
	} else {
		goto L603
	}
L582:
	;
	if v3298 != 0 {
		goto L581
	} else {
		goto L583
	}
L583:
	;
	v3300 = F_contain_volatile_functions(m, v3297)
	mBase = m.M
	v3301 = m.ExcPending
	if v3301 != 0 {
		goto L1
	} else {
		goto L584
	}
L584:
	;
	if v3300 != 0 {
		goto L581
	} else {
		goto L585
	}
L585:
	;
	v3302 = F_contain_subplans(m, v3297)
	mBase = m.M
	v3303 = m.ExcPending
	if v3303 != 0 {
		goto L1
	} else {
		goto L586
	}
L586:
	;
	if v3302 != 0 {
		goto L581
	} else {
		goto L587
	}
L587:
	;
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+100))
	if v3304 == int32(0) {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v3336 = F_copyObjectImpl(m, v3297)
	mBase = m.M
	v3337 = m.ExcPending
	if v3337 != 0 {
		goto L1
	} else {
		goto L600
	}
L589:
	;
	v3307 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+108))
	if v3307 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v3327 = F_preprocess_expression(m, v47, v3297, int32(0))
	mBase = m.M
	v3328 = m.ExcPending
	if v3328 != 0 {
		goto L1
	} else {
		goto L598
	}
L591:
	;
	v3310 = *(*int32)(unsafe.Add(mBase, uint32(v47)+324))
	v3311 = F_pull_varnos(m, v47, v3297)
	mBase = m.M
	v3312 = m.ExcPending
	if v3312 != 0 {
		goto L1
	} else {
		goto L592
	}
L592:
	;
	v3313 = F_bms_is_member(m, v3310, v3311)
	mBase = m.M
	v3314 = m.ExcPending
	if v3314 != 0 {
		goto L1
	} else {
		goto L593
	}
L593:
	;
	if v3313 != 0 {
		goto L581
	} else {
		goto L594
	}
L594:
	;
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+100))
	if v3315 == int32(0) {
		goto L588
	} else {
		goto L595
	}
L595:
	;
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+108))
	if v3318 == int32(0) {
		goto L590
	} else {
		goto L596
	}
L596:
	;
	v3321 = *(*int32)(unsafe.Add(mBase, uint32(v3318)+12))
	v3322 = *(*int32)(unsafe.Add(mBase, uint32(v3321)))
	if v3322 == int32(0) {
		goto L588
	} else {
		goto L597
	}
L597:
	;
	goto L590
L598:
	;
	v3329 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	v3330 = *(*int32)(unsafe.Add(mBase, uint32(v3329)+8))
	v3331 = F_list_concat(m, v3330, v3327)
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	v3333 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v3333)+8)) = v3331
	v3352 = v3267
	goto L580
L600:
	;
	v3339 = F_preprocess_expression(m, v47, v3336, int32(0))
	mBase = m.M
	v3340 = m.ExcPending
	if v3340 != 0 {
		goto L1
	} else {
		goto L601
	}
L601:
	;
	v3341 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	v3342 = *(*int32)(unsafe.Add(mBase, uint32(v3341)+8))
	v3343 = F_list_concat(m, v3342, v3339)
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L1
	} else {
		goto L602
	}
L602:
	;
	v3345 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v3345)+8)) = v3343
	goto L581
L603:
	;
	v3352 = v3348
	goto L580
L604:
	;
	goto L579
L605:
	;
	v3792 = int32(0)
	v3795 = m.G0
	v3797 = v3795 - int32(352)
	m.G0 = v3797
	v3799 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3800 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+132))
	if v3800 == v3792 {
		goto L651
	} else {
		goto L652
	}
L606:
	;
	v3606 = int32(0)
	v3607 = m.G0
	v3609 = v3607 - int32(16)
	m.G0 = v3609
	*(*int32)(unsafe.Add(mBase, uint32(v3609)+12)) = v3606
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3614 = *(*int32)(unsafe.Add(mBase, uint32(v3613)+60))
	v3618 = F_remove_useless_results_recurse(m, v47, v3614, v3606, v3609+int32(12))
	mBase = m.M
	v3619 = m.ExcPending
	if v3619 != 0 {
		goto L1
	} else {
		goto L633
	}
L607:
	;
	v3401 = m.G0
	v3403 = v3401 - int32(16)
	m.G0 = v3403
	v3405 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3406 = *(*int32)(unsafe.Add(mBase, uint32(v3405)+60))
	v3407 = F_reduce_outer_joins_pass1(m, v3406)
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
		goto L1
	} else {
		goto L611
	}
L608:
	;
	goto L609
L609:
	;
	if v1640 == int32(0) {
		goto L605
	} else {
		goto L632
	}
L610:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3553 = m.ExcPending
	if v3553 != 0 {
		goto L1
	} else {
		goto L629
	}
L611:
	;
	if v3407 == int32(0) {
		goto L610
	} else {
		goto L612
	}
L612:
	;
	v3411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3407)+4)))
	if v3411 == int32(0) {
		goto L610
	} else {
		goto L613
	}
L613:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3403)+8)) = int64(0)
	v3416 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(v3416)+60))
	v3420 = int32(0)
	F_reduce_outer_joins_pass2(m, v3417, v3407, v3403+int32(8), v47, v3420, v3420)
	mBase = m.M
	v3423 = m.ExcPending
	if v3423 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	v3424 = *(*int32)(unsafe.Add(mBase, uint32(v3403)+8))
	if v3424 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3425 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3427 = F_remove_nulling_relids(m, v3425, v3424, int32(0))
	mBase = m.M
	v3428 = m.ExcPending
	if v3428 != 0 {
		goto L1
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	v3436 = *(*int32)(unsafe.Add(mBase, uint32(v3403)+12))
	if v3436 == int32(0) {
		goto L620
	} else {
		goto L621
	}
L618:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3427
	v3430 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(v3403)+8))
	v3433 = F_remove_nulling_relids(m, v3430, v3431, int32(0))
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L1
	} else {
		goto L619
	}
L619:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v3433
	goto L617
L620:
	;
	m.G0 = v3403 + int32(16)
	goto L606
L621:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(v3436)+4))
	if v3439 <= int32(0) {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v3453 = int32(0)
	goto L623
L623:
	;
	v3484 = *(*int32)(unsafe.Add(mBase, uint32(v3436)+12))
	v3488 = *(*int32)(unsafe.Add(mBase, uint32(v3484+v3453<<(uint(int32(2))%32))))
	v3489 = *(*int32)(unsafe.Add(mBase, uint32(v3488)))
	v3490 = F_bms_make_singleton(m, v3489)
	mBase = m.M
	v3491 = m.ExcPending
	if v3491 != 0 {
		goto L1
	} else {
		goto L625
	}
L624:
	;
	goto L620
L625:
	;
	v3492 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3493 = *(*int32)(unsafe.Add(mBase, uint32(v3488)+4))
	v3494 = F_remove_nulling_relids(m, v3492, v3490, v3493)
	mBase = m.M
	v3495 = m.ExcPending
	if v3495 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3494
	v3497 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v3498 = *(*int32)(unsafe.Add(mBase, uint32(v3488)+4))
	v3499 = F_remove_nulling_relids(m, v3497, v3490, v3498)
	mBase = m.M
	v3500 = m.ExcPending
	if v3500 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v3499
	v3503 = v3453 + int32(1)
	v3504 = *(*int32)(unsafe.Add(mBase, uint32(v3436)+4))
	if v3503 < v3504 {
		v3453 = v3503
		goto L623
	} else {
		goto L628
	}
L628:
	;
	goto L624
L629:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_14), int32(0))
	mBase = m.M
	v3557 = m.ExcPending
	if v3557 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_6), int32(3121), int32(_a_F_subquery_planner_15))
	mBase = m.M
	v3562 = m.ExcPending
	if v3562 != 0 {
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
	goto L606
L633:
	;
	v3620 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3620)+60)) = v3618
	v3622 = *(*int32)(unsafe.Add(mBase, uint32(v3609)+12))
	if v3622 != 0 {
		goto L634
	} else {
		goto L635
	}
L634:
	;
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3625 = F_remove_nulling_relids(m, v3623, v3622, int32(0))
	mBase = m.M
	v3626 = m.ExcPending
	if v3626 != 0 {
		goto L1
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v3634 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	if v3634 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3625
	v3628 = *(*int32)(unsafe.Add(mBase, uint32(v47)+128))
	v3629 = *(*int32)(unsafe.Add(mBase, uint32(v3609)+12))
	v3631 = F_remove_nulling_relids(m, v3628, v3629, int32(0))
	mBase = m.M
	v3632 = m.ExcPending
	if v3632 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+128)) = v3631
	goto L636
L639:
	;
	m.G0 = v3609 + int32(16)
	goto L605
L640:
	;
	v3638 = v3606
	v3639 = v3634
	goto L641
L641:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, uint32(v3639)+4))
	if v3678 <= v3638 {
		goto L639
	} else {
		goto L643
	}
L642:
	;
	goto L639
L643:
	;
	v3680 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(v3680)+52))
	v3682 = *(*int32)(unsafe.Add(mBase, uint32(v3681)+12))
	v3683 = *(*int32)(unsafe.Add(mBase, uint32(v3639)+12))
	v3684 = int32(2)
	v3687 = *(*int32)(unsafe.Add(mBase, uint32(v3683+v3638<<(uint(v3684)%32))))
	v3688 = *(*int32)(unsafe.Add(mBase, uint32(v3687)+4))
	v3694 = *(*int32)(unsafe.Add(mBase, uint32(v3682+v3688<<(uint(v3684)%32)-int32(4))))
	v3695 = *(*int32)(unsafe.Add(mBase, uint32(v3694)+12))
	if v3695 == int32(8) {
		goto L644
	} else {
		goto L645
	}
L644:
	;
	v3698 = F_list_delete_nth_cell(m, v3639, v3638)
	mBase = m.M
	v3699 = m.ExcPending
	if v3699 != 0 {
		goto L1
	} else {
		goto L647
	}
L645:
	;
	v3703 = v3639
	v3704 = v3638
	goto L646
L646:
	;
	if v3703 != 0 {
		v3638 = v3704 + int32(1)
		v3639 = v3703
		goto L641
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v3698
	v3703 = v3698
	v3704 = v3638 - int32(1)
	goto L646
L648:
	;
	goto L642
L649:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v47)+296)) = v3911
	v3917 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+144))
	if v3917 != 0 {
		goto L724
	} else {
		goto L725
	}
L650:
	;
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+128))
	if v3823 == int32(0) {
		v3843 = v39
		goto L662
	} else {
		goto L663
	}
L651:
	;
	v3803 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+128))
	if v3803 != 0 {
		v3822 = v39
		goto L650
	} else {
		goto L654
	}
L652:
	;
	goto L653
L653:
	;
	v3805 = int64(-1)
	v3806 = F_estimate_expression_value(m, v47, v3800)
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L1
	} else {
		goto L655
	}
L654:
	;
	v3911 = l4
	v3913 = float64(-1)
	v3914 = v39
	v3915 = v39
	goto L649
L655:
	;
	if v3806 == int32(0) {
		v3822 = v3805
		goto L650
	} else {
		goto L656
	}
L656:
	;
	v3810 = *(*int32)(unsafe.Add(mBase, uint32(v3806)))
	if v3810 != int32(7) {
		v3822 = v3805
		goto L650
	} else {
		goto L657
	}
L657:
	;
	v3814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3806)+24)))
	if v3814 != 0 {
		v3822 = int64(0)
		goto L650
	} else {
		goto L658
	}
L658:
	;
	v3815 = int64(1)
	v3816 = *(*int32)(unsafe.Add(mBase, uint32(v3806)+20))
	v3817 = *(*int64)(unsafe.Add(mBase, uint32(v3816)))
	if v3817 <= v3815 {
		goto L659
	} else {
		goto L660
	}
L659:
	;
	v3820 = v3815
	goto L661
L660:
	;
	v3820 = v3817
	goto L661
L661:
	;
	v3822 = v3820
	goto L650
L662:
	;
	if v3822 != int64(0) {
		goto L673
	} else {
		goto L674
	}
L663:
	;
	v3826 = int64(-1)
	v3827 = F_estimate_expression_value(m, v47, v3823)
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	if v3827 == int32(0) {
		v3843 = v3826
		goto L662
	} else {
		goto L665
	}
L665:
	;
	v3831 = *(*int32)(unsafe.Add(mBase, uint32(v3827)))
	if v3831 != int32(7) {
		v3843 = v3826
		goto L662
	} else {
		goto L666
	}
L666:
	;
	v3835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3827)+24)))
	if v3835 != 0 {
		v3843 = int64(0)
		goto L662
	} else {
		goto L667
	}
L667:
	;
	v3836 = *(*int32)(unsafe.Add(mBase, uint32(v3827)+20))
	v3837 = *(*int64)(unsafe.Add(mBase, uint32(v3836)))
	v3838 = int64(0)
	if v3838 < v3837 {
		goto L668
	} else {
		goto L669
	}
L668:
	;
	v3841 = v3837
	goto L670
L669:
	;
	v3841 = v3838
	goto L670
L670:
	;
	v3843 = v3841
	goto L662
L671:
	;
	if int64(0) <= v3843 {
		goto L705
	} else {
		goto L706
	}
L672:
	;
	if base.F64_lt(l4, v3853) != 0 {
		goto L702
	} else {
		goto L703
	}
L673:
	;
	v3849 = base.F64_add(base.F64_convert_i64_u(v3822), base.F64_convert_i64_u(v3843))
	if v3843|v3822 < int64(0) {
		goto L676
	} else {
		goto L677
	}
L674:
	;
	goto L675
L675:
	;
	v3870 = float64(-1)
	v3871 = int64(0)
	if base.B2i32(base.F64_gt(l4, float64(0)) == int32(0))|base.B2i32(v3843 == v3871) != 0 {
		v3911 = l4
		v3913 = v3870
		v3914 = v3843
		v3915 = v3871
		goto L649
	} else {
		goto L690
	}
L676:
	;
	v3853 = float64(0.1)
	goto L678
L677:
	;
	v3853 = v3849
	goto L678
L678:
	;
	if base.F64_ge(l4, float64(1)) != 0 {
		goto L679
	} else {
		goto L680
	}
L679:
	;
	if base.F64_ge(v3853, float64(1)) == int32(0) {
		v3901 = l4
		goto L671
	} else {
		goto L682
	}
L680:
	;
	goto L681
L681:
	;
	if base.F64_gt(l4, float64(0)) == int32(0) {
		goto L686
	} else {
		goto L687
	}
L682:
	;
	if base.F64_lt(l4, v3853) != 0 {
		goto L683
	} else {
		goto L684
	}
L683:
	;
	v3861 = l4
	goto L685
L684:
	;
	v3861 = v3853
	goto L685
L685:
	;
	v3901 = v3861
	goto L671
L686:
	;
	v3901 = v3853
	goto L671
L687:
	;
	goto L688
L688:
	;
	if base.F64_ge(v3853, float64(1)) == int32(0) {
		goto L672
	} else {
		goto L689
	}
L689:
	;
	v3901 = v3853
	goto L671
L690:
	;
	if v3843 < int64(0) {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	v3883 = float64(0.1)
	goto L693
L692:
	;
	v3883 = base.F64_convert_i64_u(v3843)
	goto L693
L693:
	;
	if base.F64_ge(l4, float64(1)) != 0 {
		goto L694
	} else {
		goto L695
	}
L694:
	;
	if base.F64_ge(v3883, float64(1)) == int32(0) {
		goto L697
	} else {
		goto L698
	}
L695:
	;
	goto L696
L696:
	;
	if base.F64_ge(v3883, float64(1)) != 0 {
		v3911 = l4
		v3913 = v3870
		v3914 = v3843
		v3915 = v3871
		goto L649
	} else {
		goto L700
	}
L697:
	;
	v3911 = v3883
	v3913 = v3870
	v3914 = v3843
	v3915 = v3871
	goto L649
L698:
	;
	goto L699
L699:
	;
	v3911 = base.F64_add(l4, v3883)
	v3913 = v3870
	v3914 = v3843
	v3915 = v3871
	goto L649
L700:
	;
	v3893 = base.F64_add(l4, v3883)
	if base.F64_ge(v3893, float64(1)) == int32(0) {
		v3911 = v3893
		v3913 = v3870
		v3914 = v3843
		v3915 = v3871
		goto L649
	} else {
		goto L701
	}
L701:
	;
	v3911 = float64(0)
	v3913 = v3870
	v3914 = v3843
	v3915 = v3871
	goto L649
L702:
	;
	v3900 = l4
	goto L704
L703:
	;
	v3900 = v3853
	goto L704
L704:
	;
	v3901 = v3900
	goto L671
L705:
	;
	v3905 = v3849
	goto L707
L706:
	;
	v3905 = float64(-1)
	goto L707
L707:
	;
	if int64(0) < v3822 {
		goto L708
	} else {
		goto L709
	}
L708:
	;
	v3909 = v3905
	goto L710
L709:
	;
	v3909 = float64(-1)
	goto L710
L710:
	;
	v3911 = v3901
	v3913 = v3909
	v3914 = v3843
	v3915 = v3822
	goto L649
L711:
	;
	F_SS_identify_outer_params(m, v16801)
	mBase = m.M
	v17073 = m.ExcPending
	if v17073 != 0 {
		goto L1
	} else {
		goto L2737
	}
L712:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v17055 = m.ExcPending
	if v17055 != 0 {
		goto L1
	} else {
		goto L2732
	}
L713:
	;
	v15234 = *(*int32)(unsafe.Add(mBase, uint32(v15209)+124))
	if v15234 == int32(0) {
		goto L2343
	} else {
		goto L2344
	}
L714:
	;
	v8412 = int32(0)
	v8416 = m.G0
	v8418 = v8416 - int32(80)
	m.G0 = v8418
	v8420 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v8421 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+4))
	v8422 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+52))
	v8423 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+32))
	if v8423 != 0 {
		goto L1260
	} else {
		goto L1261
	}
L715:
	;
	v8057 = *(*int32)(unsafe.Add(mBase, uint32(v8035)+28))
	if v8057 == int32(0) {
		v8377 = v8022
		v8381 = v8026
		v8382 = v8027
		v8387 = v8032
		v8390 = v8035
		v8396 = v8041
		v8398 = v8043
		v8405 = v8050
		v8409 = v8054
		v8410 = v8055
		goto L714
	} else {
		goto L1234
	}
L716:
	;
	if v6949 == int32(0) {
		v8022 = v47
		v8026 = l5
		v8027 = v3797
		v8032 = v3799
		v8035 = v4573
		v8037 = v4571
		v8041 = v44
		v8043 = v7
		v8050 = v3913
		v8054 = v3914
		v8055 = v3915
		goto L715
	} else {
		goto L1120
	}
L717:
	;
	v5018 = int32(1)
	v5020 = v4919 << (uint(int32(2)) % 32)
	v5021 = F_palloc0(m, v5020)
	mBase = m.M
	v5022 = m.ExcPending
	if v5022 != 0 {
		goto L1
	} else {
		goto L900
	}
L718:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+72)) = v4982
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+192)) = v4982
	v5016 = F_list_make1_impl(m, int32(1), v3797+int32(72))
	mBase = m.M
	v5017 = m.ExcPending
	if v5017 != 0 {
		goto L1
	} else {
		goto L899
	}
L719:
	;
	v4914 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+12))
	if v4914 == int32(0) {
		v4982 = v4885
		goto L718
	} else {
		goto L894
	}
L720:
	;
	if v4800 == int32(0) {
		v8022 = v47
		v8026 = l5
		v8027 = v3797
		v8032 = v3799
		v8035 = v4573
		v8037 = v4571
		v8041 = v44
		v8043 = v7
		v8050 = v3913
		v8054 = v3914
		v8055 = v3915
		goto L715
	} else {
		goto L893
	}
L721:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4843 = m.ExcPending
	if v4843 != 0 {
		goto L1
	} else {
		goto L885
	}
L722:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4830 = m.ExcPending
	if v4830 != 0 {
		goto L1
	} else {
		goto L882
	}
L723:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4817 = m.ExcPending
	if v4817 != 0 {
		goto L1
	} else {
		goto L879
	}
L724:
	;
	v3919 = m.G0
	v3921 = v3919 - int32(32)
	m.G0 = v3921
	v3923 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3924 = *(*int32)(unsafe.Add(mBase, uint32(v3923)+144))
	v3925 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+92)) = uint8(v3925)
	F_setup_simple_rel_arrays(m, v47)
	mBase = m.M
	v3928 = m.ExcPending
	if v3928 != 0 {
		goto L1
	} else {
		goto L728
	}
L725:
	;
	goto L726
L726:
	;
	v4570 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+108))
	if v4570 != 0 {
		goto L831
	} else {
		goto L832
	}
L727:
	;
	v4427 = *(*int32)(unsafe.Add(mBase, uint32(v47)+264))
	v4428 = F_copyObjectImpl(m, v4427)
	mBase = m.M
	v4429 = m.ExcPending
	if v4429 != 0 {
		goto L1
	} else {
		goto L810
	}
L728:
	;
	v3930 = v3924
	goto L729
L729:
	;
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3930)+12))
	if v3970 != 0 {
		goto L731
	} else {
		goto L732
	}
L730:
	;
	v3974 = *(*int32)(unsafe.Add(mBase, uint32(v47)+36))
	v3975 = *(*int32)(unsafe.Add(mBase, uint32(v3970)+4))
	v3979 = *(*int32)(unsafe.Add(mBase, uint32(v3974+v3975<<(uint(int32(2))%32))))
	v3980 = *(*int32)(unsafe.Add(mBase, uint32(v3979)+36))
	v3981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+322)))
	if v3981 == int32(1) {
		goto L738
	} else {
		goto L739
	}
L731:
	;
	v3971 = *(*int32)(unsafe.Add(mBase, uint32(v3970)))
	if v3971 == int32(142) {
		v3930 = v3970
		goto L729
	} else {
		goto L734
	}
L732:
	;
	goto L733
L733:
	;
	goto L730
L734:
	;
	goto L733
L735:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4410 = m.ExcPending
	if v4410 != 0 {
		goto L1
	} else {
		goto L805
	}
L736:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4397 = m.ExcPending
	if v4397 != 0 {
		goto L1
	} else {
		goto L802
	}
L737:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+264)) = v4359
	m.G0 = v3921 + int32(32)
	goto L727
L738:
	;
	v3984 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+4))
	if v3984 != int32(1) {
		goto L736
	} else {
		goto L741
	}
L739:
	;
	goto L740
L740:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+20))
	v4340 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+28))
	v4341 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+76))
	v4346 = F_recurse_set_operations(m, v3924, v47, int32(0), v4339, v4340, v4341, v3921+int32(8), v3921+int32(28))
	mBase = m.M
	v4347 = m.ExcPending
	if v4347 != 0 {
		goto L1
	} else {
		goto L801
	}
L741:
	;
	v3987 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+12))
	v3989 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+20))
	v3990 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+28))
	v3991 = *(*int32)(unsafe.Add(mBase, uint32(v3980)+76))
	v3996 = F_recurse_set_operations(m, v3987, v47, int32(0), v3989, v3990, v3991, v3921+int32(28), v3921+int32(27))
	mBase = m.M
	v3997 = m.ExcPending
	if v3997 != 0 {
		goto L1
	} else {
		goto L742
	}
L742:
	;
	v3998 = *(*int32)(unsafe.Add(mBase, uint32(v3996)+76))
	if v3998 == int32(1) {
		goto L743
	} else {
		goto L744
	}
L743:
	;
	v4001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3921)+27)))
	v4002 = *(*int32)(unsafe.Add(mBase, uint32(v3921)+28))
	v4003 = int32(0)
	F_build_setop_child_paths(m, v47, v3996, v4001, v4002, v4003, v4003)
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L746
	}
L744:
	;
	goto L745
L745:
	;
	v4007 = *(*int32)(unsafe.Add(mBase, uint32(v3996)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+348)) = v4007
	v4009 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+16))
	v4011 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+20))
	v4012 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+28))
	v4017 = F_recurse_set_operations(m, v4009, v47, int32(0), v4011, v4012, v3991, v3921+int32(20), v3921+int32(19))
	mBase = m.M
	v4018 = m.ExcPending
	if v4018 != 0 {
		goto L1
	} else {
		goto L747
	}
L746:
	;
	goto L745
L747:
	;
	v4019 = *(*int32)(unsafe.Add(mBase, uint32(v3921)+20))
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+76))
	if v4020 == int32(1) {
		goto L748
	} else {
		goto L749
	}
L748:
	;
	v4023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3921)+19)))
	v4024 = int32(0)
	F_build_setop_child_paths(m, v47, v4017, v4023, v4019, v4024, v4024)
	mBase = m.M
	v4027 = m.ExcPending
	if v4027 != 0 {
		goto L1
	} else {
		goto L751
	}
L749:
	;
	goto L750
L750:
	;
	v4028 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+348)) = int32(0)
	v4031 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+28))
	v4032 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3921)+12)) = v4019
	v4034 = *(*int32)(unsafe.Add(mBase, uint32(v3921)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3921)+4)) = v4034
	*(*int32)(unsafe.Add(mBase, uint32(v3921))) = v4019
	v4039 = F_list_make2_impl(m, v3921+int32(4), v3921)
	mBase = m.M
	v4040 = m.ExcPending
	if v4040 != 0 {
		goto L1
	} else {
		goto L752
	}
L751:
	;
	goto L750
L752:
	;
	v4041 = F_generate_append_tlist(m, v4032, v4031, v4039, v3991)
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L1
	} else {
		goto L753
	}
L753:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3921)+8)) = v4041
	v4045 = *(*int32)(unsafe.Add(mBase, uint32(v3996)+8))
	v4046 = *(*int32)(unsafe.Add(mBase, uint32(v4017)+8))
	v4047 = F_bms_union(m, v4045, v4046)
	mBase = m.M
	v4048 = m.ExcPending
	if v4048 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	v4049 = F_fetch_upper_rel(m, v47, int32(0), v4047)
	mBase = m.M
	v4050 = m.ExcPending
	if v4050 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	v4051 = F_make_pathtarget_from_tlist(m, v4041)
	mBase = m.M
	v4052 = m.ExcPending
	if v4052 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	v4053 = F_set_pathtarget_cost_width(m, v47, v4051)
	mBase = m.M
	v4054 = m.ExcPending
	if v4054 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4049)+28)) = v4053
	v4056 = int32(0)
	v4057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3924)+8)))
	if v4057 == v4056 {
		goto L758
	} else {
		goto L759
	}
L758:
	;
	v4060 = *(*int32)(unsafe.Add(mBase, uint32(v3924)+32))
	v4061 = F_copyObjectImpl(m, v4060)
	mBase = m.M
	v4062 = m.ExcPending
	if v4062 != 0 {
		goto L1
	} else {
		goto L761
	}
L759:
	;
	v4223 = v4053
	v4226 = float64(0)
	v4229 = v4056
	goto L760
L760:
	;
	v4263 = *(*int32)(unsafe.Add(mBase, uint32(v47)+344))
	v4265 = F_palloc0(m, int32(96))
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L788
	}
L761:
	;
	if v4061 != 0 {
		goto L762
	} else {
		goto L763
	}
L762:
	;
	v4063 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+12))
	v4065 = v4063
	goto L764
L763:
	;
	v4065 = int32(0)
	goto L764
L764:
	;
	if v4041 == int32(0) {
		goto L765
	} else {
		goto L766
	}
L765:
	;
	v4176 = int32(0)
	if v4061 == v4176 {
		goto L775
	} else {
		goto L776
	}
L766:
	;
	v4068 = *(*int32)(unsafe.Add(mBase, uint32(v4041)+4))
	if v4068 <= int32(0) {
		goto L765
	} else {
		goto L767
	}
L767:
	;
	v4072 = int32(0)
	v4073 = v4065
	goto L768
L768:
	;
	v4113 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+12))
	v4114 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+4))
	v4115 = *(*int32)(unsafe.Add(mBase, uint32(v4073)))
	v4116 = *(*int32)(unsafe.Add(mBase, uint32(v4041)+12))
	v4117 = int32(2)
	v4120 = *(*int32)(unsafe.Add(mBase, uint32(v4116+v4072<<(uint(v4117)%32))))
	v4121 = *(*int32)(unsafe.Add(mBase, uint32(v4120)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4115)+4)) = v4121
	v4124 = v4073 + int32(4)
	if base.Ui32(v4124) < base.Ui32(v4113+v4114<<(uint(v4117)%32)) {
		goto L770
	} else {
		goto L771
	}
L769:
	;
	goto L765
L770:
	;
	v4130 = v4124
	goto L772
L771:
	;
	v4130 = int32(0)
	goto L772
L772:
	;
	v4132 = v4072 + int32(1)
	v4133 = *(*int32)(unsafe.Add(mBase, uint32(v4041)+4))
	if v4132 < v4133 {
		v4072 = v4132
		v4073 = v4130
		goto L768
	} else {
		goto L773
	}
L773:
	;
	goto L769
L774:
	;
	if v4213 == int32(0) {
		goto L735
	} else {
		goto L787
	}
L775:
	;
	v4213 = int32(1)
	goto L774
L776:
	;
	goto L777
L777:
	;
	v4183 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+4))
	if v4183 <= int32(0) {
		v4207 = int32(1)
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v4213 = v4207
	goto L774
L779:
	;
	v4186 = int32(0)
	if v4186 < v4183 {
		goto L780
	} else {
		goto L781
	}
L780:
	;
	v4189 = v4183
	goto L782
L781:
	;
	v4189 = v4186
	goto L782
L782:
	;
	v4190 = *(*int32)(unsafe.Add(mBase, uint32(v4061)+12))
	v4194 = v4176
	goto L783
L783:
	;
	v4198 = *(*int32)(unsafe.Add(mBase, uint32(v4190+v4194<<(uint(int32(2))%32))))
	v4199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4198)+18)))
	if v4199 != int32(1) {
		v4207 = v4199
		goto L778
	} else {
		goto L785
	}
L784:
	;
	v4207 = v4199
	goto L778
L785:
	;
	v4203 = v4194 + int32(1)
	if v4203 != v4189 {
		v4194 = v4203
		goto L783
	} else {
		goto L786
	}
L786:
	;
	goto L784
L787:
	;
	v4216 = *(*float64)(unsafe.Add(mBase, uint32(v4028)+32))
	v4219 = *(*float64)(unsafe.Add(mBase, uint32(v4007)+32))
	v4221 = *(*int32)(unsafe.Add(mBase, uint32(v4049)+28))
	v4223 = v4221
	v4226 = base.F64_add(base.F64_mul(v4216, float64(10)), v4219)
	v4229 = v4061
	goto L760
L788:
	;
	v4267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4265)+20)) = uint8(v4267)
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+16)) = v4267
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+12)) = v4223
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+8)) = v4049
	*(*int64)(unsafe.Add(mBase, uint32(v4265))) = int64(1443109011770)
	v4276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4049)+26)))
	if v4276 != int32(1) {
		v4284 = v4267
		goto L789
	} else {
		goto L790
	}
L789:
	;
	v4286 = v4284 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4265)+21)) = uint8(v4286)
	v4288 = *(*int32)(unsafe.Add(mBase, uint32(v4007)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v4265)+88)) = v4226
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+84)) = v4263
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+80)) = v4229
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+76)) = v4028
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+72)) = v4007
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+24)) = v4288
	v4298 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	v4299 = *(*float64)(unsafe.Add(mBase, uint32(v4007)+56))
	v4300 = *(*float64)(unsafe.Add(mBase, uint32(v4028)+56))
	v4301 = *(*float64)(unsafe.Add(mBase, uint32(v4007)+32))
	v4302 = *(*float64)(unsafe.Add(mBase, uint32(v4028)+32))
	v4303 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+40))
	v4304 = *(*int32)(unsafe.Add(mBase, uint32(v4007)+40))
	v4305 = *(*float64)(unsafe.Add(mBase, uint32(v4007)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4265)+48)) = v4305
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+40)) = v4303 + v4304
	v4309 = float64(10)
	v4311 = base.F64_add(v4301, base.F64_mul(v4302, v4309))
	*(*float64)(unsafe.Add(mBase, uint32(v4265)+32)) = v4311
	*(*float64)(unsafe.Add(mBase, uint32(v4265)+56)) = base.F64_add(base.F64_mul(v4298, v4311), base.F64_add(v4299, base.F64_mul(v4300, v4309)))
	v4319 = *(*int32)(unsafe.Add(mBase, uint32(v4265)+12))
	v4320 = *(*int32)(unsafe.Add(mBase, uint32(v4007)+12))
	v4321 = *(*int32)(unsafe.Add(mBase, uint32(v4320)+32))
	v4322 = *(*int32)(unsafe.Add(mBase, uint32(v4028)+12))
	v4323 = *(*int32)(unsafe.Add(mBase, uint32(v4322)+32))
	if v4323 < v4321 {
		goto L792
	} else {
		goto L793
	}
L790:
	;
	v4280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4007)+21)))
	if v4280 != int32(1) {
		v4284 = int32(0)
		goto L789
	} else {
		goto L791
	}
L791:
	;
	v4283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4028)+21)))
	v4284 = v4283
	goto L789
L792:
	;
	v4325 = v4321
	goto L794
L793:
	;
	v4325 = v4323
	goto L794
L794:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4319)+32)) = v4325
	F_add_path(m, v4049, v4265)
	mBase = m.M
	v4328 = m.ExcPending
	if v4328 != 0 {
		goto L1
	} else {
		goto L795
	}
L795:
	;
	v4330 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v4330 != 0 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v4331 = int32(0)
	m.T0[v4330].(func(*base.Module, int32, int32, int32, int32, int32))(m, v47, v4331, v4331, v4049, v4331)
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L1
	} else {
		goto L799
	}
L797:
	;
	goto L798
L798:
	;
	F_set_cheapest(m, v4049)
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L1
	} else {
		goto L800
	}
L799:
	;
	goto L798
L800:
	;
	v4351 = v4049
	v4359 = v4041
	goto L737
L801:
	;
	v4348 = *(*int32)(unsafe.Add(mBase, uint32(v3921)+8))
	v4351 = v4346
	v4359 = v4348
	goto L737
L802:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_16), int32(0))
	mBase = m.M
	v4401 = m.ExcPending
	if v4401 != 0 {
		goto L1
	} else {
		goto L803
	}
L803:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_17), int32(377), int32(_a_F_subquery_planner_18))
	mBase = m.M
	v4406 = m.ExcPending
	if v4406 != 0 {
		goto L1
	} else {
		goto L804
	}
L804:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L805:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4413 = m.ExcPending
	if v4413 != 0 {
		goto L1
	} else {
		goto L806
	}
L806:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_19), int32(0))
	mBase = m.M
	v4417 = m.ExcPending
	if v4417 != 0 {
		goto L1
	} else {
		goto L807
	}
L807:
	;
	F_errdetail(m, int32(_a_F_subquery_planner_20), int32(0))
	mBase = m.M
	v4421 = m.ExcPending
	if v4421 != 0 {
		goto L1
	} else {
		goto L808
	}
L808:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_17), int32(441), int32(_a_F_subquery_planner_18))
	mBase = m.M
	v4426 = m.ExcPending
	if v4426 != 0 {
		goto L1
	} else {
		goto L809
	}
L809:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L810:
	;
	v4430 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+76))
	if v4430 != 0 {
		goto L811
	} else {
		goto L812
	}
L811:
	;
	v4431 = *(*int32)(unsafe.Add(mBase, uint32(v4430)+12))
	v4433 = v4431
	goto L813
L812:
	;
	v4433 = int32(0)
	goto L813
L813:
	;
	if v4428 == int32(0) {
		v4530 = v4433
		goto L814
	} else {
		goto L815
	}
L814:
	;
	if v4530 != 0 {
		goto L722
	} else {
		goto L827
	}
L815:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(v4428)+4))
	if v4436 <= int32(0) {
		v4530 = v4433
		goto L814
	} else {
		goto L816
	}
L816:
	;
	v4440 = v4436
	v4447 = int32(0)
	v4457 = v4433
	goto L817
L817:
	;
	v4481 = *(*int32)(unsafe.Add(mBase, uint32(v4428)+12))
	v4485 = *(*int32)(unsafe.Add(mBase, uint32(v4481+v4447<<(uint(int32(2))%32))))
	v4486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4485)+26)))
	if v4486 == int32(0) {
		goto L819
	} else {
		goto L820
	}
L818:
	;
	v4530 = v4509
	goto L814
L819:
	;
	v4489 = *(*int32)(unsafe.Add(mBase, uint32(v4457)))
	v4490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4489)+26)))
	if v4490 == int32(1) {
		goto L723
	} else {
		goto L822
	}
L820:
	;
	v4506 = v4440
	v4509 = v4457
	goto L821
L821:
	;
	v4511 = v4447 + int32(1)
	if v4511 < v4506 {
		v4440 = v4506
		v4447 = v4511
		v4457 = v4509
		goto L817
	} else {
		goto L826
	}
L822:
	;
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v4430)+12))
	v4494 = *(*int32)(unsafe.Add(mBase, uint32(v4430)+4))
	v4495 = *(*int32)(unsafe.Add(mBase, uint32(v4489)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4485)+16)) = v4495
	v4498 = v4457 + int32(4)
	if base.Ui32(v4498) < base.Ui32(v4493+v4494<<(uint(int32(2))%32)) {
		goto L823
	} else {
		goto L824
	}
L823:
	;
	v4504 = v4498
	goto L825
L824:
	;
	v4504 = int32(0)
	goto L825
L825:
	;
	v4505 = *(*int32)(unsafe.Add(mBase, uint32(v4428)+4))
	v4506 = v4505
	v4509 = v4504
	goto L821
L826:
	;
	goto L818
L827:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+264)) = v4428
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v4351)+48))
	v4556 = *(*int32)(unsafe.Add(mBase, uint32(v4555)+12))
	v4557 = *(*int32)(unsafe.Add(mBase, uint32(v4556)+4))
	v4558 = F_is_parallel_safe(m, v47, v4557)
	mBase = m.M
	v4559 = m.ExcPending
	if v4559 != 0 {
		goto L1
	} else {
		goto L828
	}
L828:
	;
	v4560 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+188)) = v4560
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+184)) = v4560
	v4564 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+140))
	if v4564 != 0 {
		goto L721
	} else {
		goto L829
	}
L829:
	;
	v4565 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+124))
	v4566 = *(*int32)(unsafe.Add(mBase, uint32(v47)+264))
	v4567 = F_make_pathkeys_for_sortclauses(m, v47, v4565, v4566)
	mBase = m.M
	v4568 = m.ExcPending
	if v4568 != 0 {
		goto L1
	} else {
		goto L830
	}
L830:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+176)) = v4567
	v15197 = v3913
	v15199 = v47
	v15204 = v3797
	v15205 = v4351
	v15208 = v4556
	v15209 = v3799
	v15217 = v4558
	v15218 = v44
	v15227 = v3913
	v15231 = v3914
	v15232 = v3915
	goto L713
L831:
	;
	v4571 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v4573 = F_palloc0(m, int32(40))
	mBase = m.M
	v4574 = m.ExcPending
	if v4574 != 0 {
		goto L1
	} else {
		goto L834
	}
L832:
	;
	goto L833
L833:
	;
	v4807 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+100))
	if v4807 == int32(0) {
		v8377 = v47
		v8381 = l5
		v8382 = v3797
		v8387 = v3799
		v8390 = v7
		v8396 = v44
		v8398 = v7
		v8405 = v3913
		v8409 = v3914
		v8410 = v3915
		goto L714
	} else {
		goto L877
	}
L834:
	;
	v4575 = *(*int32)(unsafe.Add(mBase, uint32(v4571)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+256)) = v4575
	v4577 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4573)+16)) = uint8(v4577)
	*(*int32)(unsafe.Add(mBase, uint32(v4573)+28)) = v4577
	*(*int64)(unsafe.Add(mBase, uint32(v4573)+20)) = int64(0)
	v4583 = int32(4)
	v4584 = *(*int32)(unsafe.Add(mBase, uint32(v4571)+100))
	if v4584 == v4577 {
		v4668 = v4583
		goto L835
	} else {
		goto L836
	}
L835:
	;
	v4702 = F_palloc(m, v4668)
	mBase = m.M
	v4703 = m.ExcPending
	if v4703 != 0 {
		goto L1
	} else {
		goto L852
	}
L836:
	;
	v4587 = *(*int32)(unsafe.Add(mBase, uint32(v4584)+4))
	if v4587 <= int32(0) {
		v4668 = v4583
		goto L835
	} else {
		goto L837
	}
L837:
	;
	v4598 = v3792
	v4607 = v7
	goto L838
L838:
	;
	v4631 = *(*int32)(unsafe.Add(mBase, uint32(v4584)+12))
	v4635 = *(*int32)(unsafe.Add(mBase, uint32(v4631+v4607<<(uint(int32(2))%32))))
	v4636 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+4))
	v4637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4635)+18)))
	if v4637 == int32(0) {
		goto L840
	} else {
		goto L841
	}
L839:
	;
	v4668 = v4652<<(uint(int32(2))%32) + int32(4)
	goto L835
L840:
	;
	v4640 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+24))
	v4641 = F_bms_add_member(m, v4640, v4636)
	mBase = m.M
	v4642 = m.ExcPending
	if v4642 != 0 {
		goto L1
	} else {
		goto L843
	}
L841:
	;
	goto L842
L842:
	;
	v4644 = *(*int32)(unsafe.Add(mBase, uint32(v4635)+12))
	if v4644 == int32(0) {
		goto L844
	} else {
		goto L845
	}
L843:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4573)+24)) = v4641
	goto L842
L844:
	;
	v4647 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+20))
	v4648 = F_bms_add_member(m, v4647, v4636)
	mBase = m.M
	v4649 = m.ExcPending
	if v4649 != 0 {
		goto L1
	} else {
		goto L847
	}
L845:
	;
	goto L846
L846:
	;
	if base.Ui32(v4598) < base.Ui32(v4636) {
		goto L848
	} else {
		goto L849
	}
L847:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4573)+20)) = v4648
	goto L846
L848:
	;
	v4652 = v4636
	goto L850
L849:
	;
	v4652 = v4598
	goto L850
L850:
	;
	v4654 = v4607 + int32(1)
	v4655 = *(*int32)(unsafe.Add(mBase, uint32(v4584)+4))
	if v4654 < v4655 {
		v4598 = v4652
		v4607 = v4654
		goto L838
	} else {
		goto L851
	}
L851:
	;
	goto L839
L852:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4573)+32)) = v4702
	v4705 = *(*int32)(unsafe.Add(mBase, uint32(v4571)+108))
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+20))
	if v4706 != 0 {
		goto L853
	} else {
		goto L854
	}
L853:
	;
	if v4705 == int32(0) {
		v8022 = v47
		v8026 = l5
		v8027 = v3797
		v8032 = v3799
		v8035 = v4573
		v8037 = v4571
		v8041 = v44
		v8043 = v7
		v8050 = v3913
		v8054 = v3914
		v8055 = v3915
		goto L715
	} else {
		goto L856
	}
L854:
	;
	goto L855
L855:
	;
	if v4705 != 0 {
		v4885 = v4705
		goto L719
	} else {
		goto L876
	}
L856:
	;
	v4709 = *(*int32)(unsafe.Add(mBase, uint32(v4705)+4))
	if v4709 <= int32(0) {
		v8022 = v47
		v8026 = l5
		v8027 = v3797
		v8032 = v3799
		v8035 = v4573
		v8037 = v4571
		v8041 = v44
		v8043 = v7
		v8050 = v3913
		v8054 = v3914
		v8055 = v3915
		goto L715
	} else {
		goto L857
	}
L857:
	;
	v4712 = int32(0)
	v4714 = v4712
	v4721 = v4712
	goto L858
L858:
	;
	v4755 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+20))
	v4756 = *(*int32)(unsafe.Add(mBase, uint32(v4705)+12))
	v4760 = *(*int32)(unsafe.Add(mBase, uint32(v4756+v4721<<(uint(int32(2))%32))))
	v4761 = F_bms_overlap_list(m, v4755, v4760)
	mBase = m.M
	v4762 = m.ExcPending
	if v4762 != 0 {
		goto L1
	} else {
		goto L861
	}
L859:
	;
	goto L720
L860:
	;
	v4803 = v4721 + int32(1)
	v4804 = *(*int32)(unsafe.Add(mBase, uint32(v4705)+4))
	if v4803 < v4804 {
		v4714 = v4800
		v4721 = v4803
		goto L858
	} else {
		goto L875
	}
L861:
	;
	if v4761 != 0 {
		goto L862
	} else {
		goto L863
	}
L862:
	;
	v4764 = F_palloc0(m, int32(16))
	mBase = m.M
	v4765 = m.ExcPending
	if v4765 != 0 {
		goto L1
	} else {
		goto L865
	}
L863:
	;
	goto L864
L864:
	;
	v4798 = F_lappend(m, v4714, v4760)
	mBase = m.M
	v4799 = m.ExcPending
	if v4799 != 0 {
		goto L1
	} else {
		goto L874
	}
L865:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4764)+4)) = v4760
	*(*int32)(unsafe.Add(mBase, uint32(v4764))) = int32(308)
	v4769 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+28))
	v4770 = F_lappend(m, v4769, v4764)
	mBase = m.M
	v4771 = m.ExcPending
	if v4771 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4573)+28)) = v4770
	v4773 = *(*int32)(unsafe.Add(mBase, uint32(v4573)+24))
	v4774 = F_bms_overlap_list(m, v4773, v4760)
	mBase = m.M
	v4775 = m.ExcPending
	if v4775 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	if v4774 == int32(0) {
		v4800 = v4714
		goto L860
	} else {
		goto L868
	}
L868:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4781 = m.ExcPending
	if v4781 != 0 {
		goto L1
	} else {
		goto L869
	}
L869:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4784 = m.ExcPending
	if v4784 != 0 {
		goto L1
	} else {
		goto L870
	}
L870:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_21), int32(0))
	mBase = m.M
	v4788 = m.ExcPending
	if v4788 != 0 {
		goto L1
	} else {
		goto L871
	}
L871:
	;
	F_errdetail(m, int32(_a_F_subquery_planner_22), int32(0))
	mBase = m.M
	v4792 = m.ExcPending
	if v4792 != 0 {
		goto L1
	} else {
		goto L872
	}
L872:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(2259), int32(_a_F_subquery_planner_23))
	mBase = m.M
	v4797 = m.ExcPending
	if v4797 != 0 {
		goto L1
	} else {
		goto L873
	}
L873:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L874:
	;
	v4800 = v4798
	goto L860
L875:
	;
	goto L859
L876:
	;
	v4982 = int32(0)
	goto L718
L877:
	;
	v4811 = F_preprocess_groupclause(m, v47, int32(0))
	mBase = m.M
	v4812 = m.ExcPending
	if v4812 != 0 {
		goto L1
	} else {
		goto L878
	}
L878:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+256)) = v4811
	v8377 = v47
	v8381 = l5
	v8382 = v3797
	v8387 = v3799
	v8390 = v7
	v8396 = v44
	v8398 = v7
	v8405 = v3913
	v8409 = v3914
	v8410 = v3915
	goto L714
L879:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_24), int32(0))
	mBase = m.M
	v4821 = m.ExcPending
	if v4821 != 0 {
		goto L1
	} else {
		goto L880
	}
L880:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(_a_F_subquery_planner_25), int32(_a_F_subquery_planner_26))
	mBase = m.M
	v4826 = m.ExcPending
	if v4826 != 0 {
		goto L1
	} else {
		goto L881
	}
L881:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L882:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_24), int32(0))
	mBase = m.M
	v4834 = m.ExcPending
	if v4834 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(_a_F_subquery_planner_27), int32(_a_F_subquery_planner_26))
	mBase = m.M
	v4839 = m.ExcPending
	if v4839 != 0 {
		goto L1
	} else {
		goto L884
	}
L884:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L885:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4846 = m.ExcPending
	if v4846 != 0 {
		goto L1
	} else {
		goto L886
	}
L886:
	;
	v4847 = *(*int32)(unsafe.Add(mBase, uint32(v3799)+140))
	v4848 = *(*int32)(unsafe.Add(mBase, uint32(v4847)+12))
	v4849 = *(*int32)(unsafe.Add(mBase, uint32(v4848)))
	v4850 = *(*int32)(unsafe.Add(mBase, uint32(v4849)+8))
	v4852 = v4850 - int32(1)
	if base.Ui32(v4852) <= base.Ui32(int32(3)) {
		goto L888
	} else {
		goto L889
	}
L887:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+80)) = v4859
	F_errmsg(m, int32(_a_F_subquery_planner_28), v3797+int32(80))
	mBase = m.M
	v4865 = m.ExcPending
	if v4865 != 0 {
		goto L1
	} else {
		goto L891
	}
L888:
	;
	v4857 = *(*int32)(unsafe.Add(mBase, uint32(v4852<<(uint(int32(2))%32))+uint32(_c_F_subquery_planner[3])))
	v4859 = v4857
	goto L890
L889:
	;
	v4859 = int32(_a_F_subquery_planner_29)
	goto L890
L890:
	;
	goto L887
L891:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(1514), int32(_a_F_subquery_planner_30))
	mBase = m.M
	v4870 = m.ExcPending
	if v4870 != 0 {
		goto L1
	} else {
		goto L892
	}
L892:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L893:
	;
	v4885 = v4800
	goto L719
L894:
	;
	v4917 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+4))
	v4919 = v4917 + int32(1)
	v4930 = v4914
	v4941 = v3792
	goto L895
L895:
	;
	v4964 = *(*int32)(unsafe.Add(mBase, uint32(v4930)))
	if v4964 != 0 {
		goto L717
	} else {
		goto L897
	}
L896:
	;
	v4982 = v4885
	goto L718
L897:
	;
	v4968 = v4930 + int32(4)
	if base.Ui32(v4968) < base.Ui32(v4914+v4917<<(uint(int32(2))%32)) {
		v4930 = v4968
		v4941 = v4941 + int32(1)
		goto L895
	} else {
		goto L898
	}
L898:
	;
	goto L896
L899:
	;
	v6949 = v5016
	goto L716
L900:
	;
	v5023 = F_palloc0(m, v5020)
	mBase = m.M
	v5024 = m.ExcPending
	if v5024 != 0 {
		goto L1
	} else {
		goto L901
	}
L901:
	;
	v5025 = F_palloc0(m, v5020)
	mBase = m.M
	v5026 = m.ExcPending
	if v5026 != 0 {
		goto L1
	} else {
		goto L902
	}
L902:
	;
	v5029 = F_palloc(m, v4919<<(uint(int32(1))%32))
	mBase = m.M
	v5030 = m.ExcPending
	if v5030 != 0 {
		goto L1
	} else {
		goto L903
	}
L903:
	;
	v5031 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+12))
	v5034 = (v4930 - v5031) >> (uint(int32(2)) % 32)
	v5035 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+4))
	if v5034 < v5035 {
		goto L904
	} else {
		goto L905
	}
L904:
	;
	v5039 = v5018
	v5040 = v3792
	v5050 = v5034
	v5057 = v7
	goto L907
L905:
	;
	v5660 = v5018
	goto L906
L906:
	;
	v5699 = int32(0)
	v5701 = v5660 - int32(1)
	v5703 = F_palloc(m, int32(32))
	mBase = m.M
	v5704 = m.ExcPending
	if v5704 != 0 {
		goto L1
	} else {
		goto L984
	}
L907:
	;
	v5078 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+12))
	v5082 = *(*int32)(unsafe.Add(mBase, uint32(v5078+v5050<<(uint(int32(2))%32))))
	if v5082 != 0 {
		goto L913
	} else {
		goto L914
	}
L908:
	;
	v5660 = v5615
	goto L906
L909:
	;
	v5655 = v5050 + int32(1)
	v5656 = *(*int32)(unsafe.Add(mBase, uint32(v4885)+4))
	if v5655 < v5656 {
		v5039 = v5615
		v5040 = v5616
		v5050 = v5655
		v5057 = v5633
		goto L907
	} else {
		goto L982
	}
L910:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+76)) = v5082
	*(*int32)(unsafe.Add(mBase, uint32(v3797)+304)) = v5082
	v5387 = v5039 << (uint(int32(2)) % 32)
	v5392 = F_list_make1_impl(m, int32(1), v3797+int32(76))
	mBase = m.M
	v5393 = m.ExcPending
	if v5393 != 0 {
		goto L1
	} else {
		goto L955
	}
L911:
	;
	v5339 = int32(0)
	if v5339 < v5057 {
		goto L952
	} else {
		goto L953
	}
L912:
	;
	if v5039 <= v5040 {
		v5346 = v5040
		v5351 = v5195
		v5363 = v5057
		goto L910
	} else {
		goto L931
	}
L913:
	;
	v5083 = int32(0)
	v5085 = *(*int32)(unsafe.Add(mBase, uint32(v5082)+4))
	if v5083 < v5085 {
		goto L916
	} else {
		goto L917
	}
L914:
	;
	goto L915
L915:
	;
	v5186 = int32(0)
	if v5057 != 0 {
		goto L911
	} else {
		goto L930
	}
L916:
	;
	v5095 = v5083
	v5096 = v5083
	goto L919
L917:
	;
	v5148 = v5083
	v5149 = v5085
	goto L918
L918:
	;
	if v5149 == v5057 {
		v5195 = v5148
		goto L912
	} else {
		goto L923
	}
L919:
	;
	v5129 = *(*int32)(unsafe.Add(mBase, uint32(v5082)+12))
	v5133 = *(*int32)(unsafe.Add(mBase, uint32(v5129+v5095<<(uint(int32(2))%32))))
	v5134 = F_bms_add_member(m, v5096, v5133)
	mBase = m.M
	v5135 = m.ExcPending
	if v5135 != 0 {
		goto L1
	} else {
		goto L921
	}
L920:
	;
	v5148 = v5134
	v5149 = v5138
	goto L918
L921:
	;
	v5137 = v5095 + int32(1)
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v5082)+4))
	if v5137 < v5138 {
		v5095 = v5137
		v5096 = v5134
		goto L919
	} else {
		goto L922
	}
L922:
	;
	goto L920
L923:
	;
	if v5057 < v5149 {
		goto L924
	} else {
		goto L925
	}
L924:
	;
	v5183 = v5039
	goto L926
L925:
	;
	v5183 = v5040
	goto L926
L926:
	;
	if v5149 < v5057 {
		goto L927
	} else {
		goto L928
	}
L927:
	;
	v5185 = v5057
	goto L929
L928:
	;
	v5185 = v5149
	goto L929
L929:
	;
	v5346 = v5183
	v5351 = v5148
	v5363 = v5185
	goto L910
L930:
	;
	v5195 = v5186
	goto L912
L931:
	;
	v5236 = v5040
	goto L932
L932:
	;
	v5271 = v5236 << (uint(int32(2)) % 32)
	v5273 = *(*int32)(unsafe.Add(mBase, uint32(v5023+v5271)))
	v5274 = int32(0)
	if base.B2i32(v5273 == v5274)|base.B2i32(v5195 == v5274) != 0 {
		v5320 = base.B2i32(v5273|v5195 == v5274)
		goto L935
	} else {
		goto L936
	}
L933:
	;
	if v5236 <= int32(0) {
		v5346 = v5040
		v5351 = v5195
		v5363 = v5057
		goto L910
	} else {
		goto L949
	}
L934:
	;
	if v5320 == int32(0) {
		goto L945
	} else {
		goto L946
	}
L935:
	;
	goto L934
L936:
	;
	v5288 = *(*int32)(unsafe.Add(mBase, uint32(v5273)+4))
	v5289 = *(*int32)(unsafe.Add(mBase, uint32(v5195)+4))
	if v5288 != v5289 {
		v5320 = int32(0)
		goto L935
	} else {
		goto L937
	}
L937:
	;
	v5291 = int32(1)
	if v5288 <= v5291 {
		goto L938
	} else {
		goto L939
	}
L938:
	;
	v5294 = v5291
	goto L940
L939:
	;
	v5294 = v5288
	goto L940
L940:
	;
	v5295 = int32(8)
	v5300 = int32(0)
	goto L941
L941:
	;
	v5308 = v5300 << (uint(int32(2)) % 32)
	v5310 = *(*int32)(unsafe.Add(mBase, uint32(v5273+v5295+v5308)))
	v5312 = *(*int32)(unsafe.Add(mBase, uint32(v5195+v5295+v5308)))
	v5313 = base.B2i32(v5310 == v5312)
	if v5310 != v5312 {
		v5320 = v5313
		goto L935
	} else {
		goto L943
	}
L942:
	;
	v5320 = v5313
	goto L935
L943:
	;
	v5316 = v5300 + int32(1)
	if v5316 != v5294 {
		v5300 = v5316
		goto L941
	} else {
		goto L944
	}
L944:
	;
	goto L942
L945:
	;
	v5328 = v5236 + int32(1)
	if v5328 != v5039 {
		v5236 = v5328
		goto L932
	} else {
		goto L948
	}
L946:
	;
	goto L947
L947:
	;
	goto L933
L948:
	;
	v5346 = v5040
	v5351 = v5195
	v5363 = v5057
	goto L910
L949:
	;
	v5332 = v5271 + v5021
	v5333 = *(*int32)(unsafe.Add(mBase, uint32(v5332)))
	v5334 = F_lappend(m, v5333, v5082)
	mBase = m.M
	v5335 = m.ExcPending
	if v5335 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5332))) = v5334
	F_bms_free(m, v5195)
	mBase = m.M
	v5338 = m.ExcPending
	if v5338 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	v5615 = v5039
	v5616 = v5040
	v5633 = v5057
	goto L909
L952:
	;
	v5342 = v5057
	goto L954
L953:
	;
	v5342 = v5339
	goto L954
L954:
	;
	v5346 = v5040
	v5351 = v5186
	v5363 = v5342
	goto L910
L955:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5021+v5387))) = v5392
	*(*int32)(unsafe.Add(mBase, uint32(v5387+v5023))) = v5351
	v5397 = int32(0)
	v5399 = v5346 - int32(1)
	if v5399 <= v5397 {
		goto L957
	} else {
		goto L958
	}
L956:
	;
	v5615 = v5039 + int32(1)
	v5616 = v5346
	v5633 = v5363
	goto L909
L957:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5387+v5025))) = int32(0)
	goto L956
L958:
	;
	v5402 = v5397
	v5409 = v5399
	goto L959
L959:
	;
	v5446 = *(*int32)(unsafe.Add(mBase, uint32(v5023+v5409<<(uint(int32(2))%32))))
	v5447 = int32(0)
	if v5446 == v5447 {
		goto L962
	} else {
		goto L963
	}
L960:
	;
	if v5507 <= int32(0) {
		goto L957
	} else {
		goto L979
	}
L961:
	;
	if v5500 != 0 {
		goto L975
	} else {
		goto L976
	}
L962:
	;
	v5500 = int32(1)
	goto L961
L963:
	;
	goto L964
L964:
	;
	if v5351 == int32(0) {
		v5493 = v5447
		goto L965
	} else {
		goto L966
	}
L965:
	;
	v5500 = v5493
	goto L961
L966:
	;
	v5456 = *(*int32)(unsafe.Add(mBase, uint32(v5446)+4))
	v5457 = *(*int32)(unsafe.Add(mBase, uint32(v5351)+4))
	if v5457 < v5456 {
		v5493 = v5447
		goto L965
	} else {
		goto L967
	}
L967:
	;
	v5459 = int32(1)
	if v5456 <= v5459 {
		goto L968
	} else {
		goto L969
	}
L968:
	;
	v5462 = v5459
	goto L970
L969:
	;
	v5462 = v5456
	goto L970
L970:
	;
	v5463 = int32(8)
	v5468 = int32(0)
	goto L971
L971:
	;
	v5475 = v5468 << (uint(int32(2)) % 32)
	v5477 = *(*int32)(unsafe.Add(mBase, uint32(v5446+v5463+v5475)))
	v5479 = *(*int32)(unsafe.Add(mBase, uint32(v5351+v5463+v5475)))
	v5482 = v5477 & (v5479 ^ int32(-1))
	v5484 = base.B2i32(v5482 == int32(0))
	if v5482 != 0 {
		v5493 = v5484
		goto L965
	} else {
		goto L973
	}
L972:
	;
	v5493 = v5484
	goto L965
L973:
	;
	v5486 = v5468 + int32(1)
	if v5486 != v5462 {
		v5468 = v5486
		goto L971
	} else {
		goto L974
	}
L974:
	;
	goto L972
L975:
	;
	v5501 = int32(1)
	v5502 = v5402 + v5501
	*(*uint16)(unsafe.Add(mBase, uint32(v5029+v5502<<(uint(v5501)%32)))) = uint16(v5409)
	v5507 = v5502
	goto L977
L976:
	;
	v5507 = v5402
	goto L977
L977:
	;
	v5508 = int32(1)
	if v5508 < v5409 {
		v5402 = v5507
		v5409 = v5409 - v5508
		goto L959
	} else {
		goto L978
	}
L978:
	;
	goto L960
L979:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5029))) = uint16(v5507)
	v5519 = v5507<<(uint(int32(1))%32) + int32(2)
	v5520 = F_palloc(m, v5519)
	mBase = m.M
	v5521 = m.ExcPending
	if v5521 != 0 {
		goto L1
	} else {
		goto L980
	}
L980:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5387+v5025))) = v5520
	if v5519 == int32(0) {
		goto L956
	} else {
		goto L981
	}
L981:
	;
	base.MemoryCopy(m, v5520, v5029, v5519)
	goto L956
L982:
	;
	goto L908
L983:
	;
	v6374 = F_palloc0(m, v5660<<(uint(int32(2))%32))
	mBase = m.M
	v6375 = m.ExcPending
	if v6375 != 0 {
		goto L1
	} else {
		goto L1053
	}
L984:
	;
	if base.B2i32(base.Ui32(int32(_a_F_subquery_planner_31)) < base.Ui32(v5701))|base.B2i32(base.Ui32(int32(_a_F_subquery_planner_32)) <= base.Ui32(v5701)) == int32(0) {
		goto L985
	} else {
		goto L986
	}
L985:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+8)) = v5025
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+4)) = v5701
	*(*int32)(unsafe.Add(mBase, uint32(v5703))) = v5701
	v5718 = v5701 << (uint(int32(1)) % 32)
	v5720 = v5718 + int32(2)
	v5721 = F_palloc0(m, v5720)
	mBase = m.M
	v5722 = m.ExcPending
	if v5722 != 0 {
		goto L1
	} else {
		goto L988
	}
L986:
	;
	goto L987
L987:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6362 = m.ExcPending
	if v6362 != 0 {
		goto L1
	} else {
		goto L1050
	}
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+16)) = v5721
	v5726 = F_palloc0(m, v5718+int32(2))
	mBase = m.M
	v5727 = m.ExcPending
	if v5727 != 0 {
		goto L1
	} else {
		goto L989
	}
L989:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+20)) = v5726
	v5729 = F_palloc(m, v5720)
	mBase = m.M
	v5730 = m.ExcPending
	if v5730 != 0 {
		goto L1
	} else {
		goto L990
	}
L990:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+24)) = v5729
	v5734 = F_palloc(m, v5718+int32(4))
	mBase = m.M
	v5735 = m.ExcPending
	if v5735 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+28)) = v5734
	v5737 = *(*int32)(unsafe.Add(mBase, uint32(v5703)))
	v5738 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+24))
	v5739 = int32(_a_F_subquery_planner_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v5738))) = uint16(v5739)
	if v5737 <= int32(0) {
		goto L992
	} else {
		goto L993
	}
L992:
	;
	goto L983
L993:
	;
	v5743 = v5737
	v5750 = v5738
	v5756 = v5734
	goto L994
L994:
	;
	v5784 = int32(1)
	v5785 = int32(0)
	v5787 = v5743 + v5784
	if int32(3) <= v5787 {
		goto L997
	} else {
		goto L998
	}
L995:
	;
	goto L992
L996:
	;
	v5985 = int32(0)
	if v5985 < v5947 {
		goto L1018
	} else {
		goto L1019
	}
L997:
	;
	v5790 = int32(2)
	if v5787 <= v5790 {
		goto L1000
	} else {
		goto L1001
	}
L998:
	;
	v5888 = v5784
	v5890 = v5785
	goto L999
L999:
	;
	v5929 = v5888 << (uint(int32(1)) % 32)
	v5930 = v5750 + v5929
	v5931 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	v5933 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5931+v5929))))
	if v5933 != 0 {
		goto L1015
	} else {
		goto L1016
	}
L1000:
	;
	v5793 = v5790
	goto L1002
L1001:
	;
	v5793 = v5787
	goto L1002
L1002:
	;
	v5794 = int32(1)
	v5795 = v5793 - v5794
	v5801 = int32(0)
	v5802 = v5784
	v5804 = v5785
	goto L1003
L1003:
	;
	v5843 = v5802 << (uint(int32(1)) % 32)
	v5844 = v5750 + v5843
	v5845 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	v5847 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5845+v5843))))
	if v5847 == int32(0) {
		goto L1006
	} else {
		goto L1007
	}
L1004:
	;
	if v5795&v5794 == int32(0) {
		v5947 = v5879
		goto L996
	} else {
		goto L1014
	}
L1005:
	;
	v5861 = int32(1)
	v5862 = v5802 + v5861
	v5864 = v5862 << (uint(v5861) % 32)
	v5865 = v5750 + v5864
	v5866 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	v5868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5866+v5864))))
	if v5868 != 0 {
		goto L1010
	} else {
		goto L1011
	}
L1006:
	;
	v5850 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5844))) = uint16(v5850)
	v5852 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5756+v5804<<(uint(v5852)%32)))) = uint16(v5802)
	v5860 = v5804 + v5852
	goto L1005
L1007:
	;
	goto L1008
L1008:
	;
	v5858 = int32(_a_F_subquery_planner_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v5844))) = uint16(v5858)
	v5860 = v5804
	goto L1005
L1009:
	;
	v5880 = int32(2)
	v5881 = v5802 + v5880
	v5883 = v5801 + v5880
	if v5883 != v5795&int32(-2) {
		v5801 = v5883
		v5802 = v5881
		v5804 = v5879
		goto L1003
	} else {
		goto L1013
	}
L1010:
	;
	v5869 = int32(_a_F_subquery_planner_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v5865))) = uint16(v5869)
	v5879 = v5860
	goto L1009
L1011:
	;
	goto L1012
L1012:
	;
	v5871 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5865))) = uint16(v5871)
	v5873 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5756+v5860<<(uint(v5873)%32)))) = uint16(v5862)
	v5879 = v5860 + v5873
	goto L1009
L1013:
	;
	goto L1004
L1014:
	;
	v5888 = v5881
	v5890 = v5879
	goto L999
L1015:
	;
	v5934 = int32(_a_F_subquery_planner_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v5930))) = uint16(v5934)
	v5947 = v5890
	goto L996
L1016:
	;
	goto L1017
L1017:
	;
	v5936 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5930))) = uint16(v5936)
	v5938 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5756+v5890<<(uint(v5938)%32)))) = uint16(v5888)
	v5947 = v5890 + v5938
	goto L996
L1018:
	;
	v5991 = v5947
	v6003 = v5985
	goto L1021
L1019:
	;
	goto L1020
L1020:
	;
	v6205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5750))))
	if v6205 == int32(_a_F_subquery_planner_32) {
		goto L992
	} else {
		goto L1034
	}
L1021:
	;
	v6029 = int32(1)
	v6032 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5756+v6003<<(uint(v6029)%32)))))
	v6035 = v5750 + v6032<<(uint(v6029)%32)
	v6036 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6035))))
	v6037 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5750))))
	if v6037 <= v6036 {
		v6123 = v5991
		goto L1023
	} else {
		goto L1024
	}
L1022:
	;
	goto L1020
L1023:
	;
	v6162 = v6003 + int32(1)
	if v6162 < v6123 {
		v5991 = v6123
		v6003 = v6162
		goto L1021
	} else {
		goto L1033
	}
L1024:
	;
	v6039 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+8))
	v6043 = *(*int32)(unsafe.Add(mBase, uint32(v6039+v6032<<(uint(int32(2))%32))))
	if v6043 == int32(0) {
		v6123 = v5991
		goto L1023
	} else {
		goto L1025
	}
L1025:
	;
	v6046 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6043))))
	if v6046 <= int32(0) {
		v6123 = v5991
		goto L1023
	} else {
		goto L1026
	}
L1026:
	;
	v6050 = v6046
	v6052 = v5991
	goto L1027
L1027:
	;
	v6090 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+20))
	v6091 = int32(1)
	v6094 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6043+v6050<<(uint(v6091)%32)))))
	v6098 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6090+v6094<<(uint(v6091)%32)))))
	v6101 = v5750 + v6098<<(uint(v6091)%32)
	v6102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6101))))
	if v6102 == int32(_a_F_subquery_planner_32) {
		goto L1029
	} else {
		goto L1030
	}
L1028:
	;
	v6123 = v6115
	goto L1023
L1029:
	;
	v6105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6035))))
	v6106 = int32(1)
	v6107 = v6105 + v6106
	*(*uint16)(unsafe.Add(mBase, uint32(v6101))) = uint16(v6107)
	*(*uint16)(unsafe.Add(mBase, uint32(v5756+v6052<<(uint(v6106)%32)))) = uint16(v6098)
	v6115 = v6052 + v6106
	goto L1031
L1030:
	;
	v6115 = v6052
	goto L1031
L1031:
	;
	v6116 = int32(1)
	if v6116 < v6050 {
		v6050 = v6050 - v6116
		v6052 = v6115
		goto L1027
	} else {
		goto L1032
	}
L1032:
	;
	goto L1028
L1033:
	;
	goto L1022
L1034:
	;
	if v5701 != 0 {
		goto L1035
	} else {
		goto L1036
	}
L1035:
	;
	v6209 = int32(1)
	goto L1038
L1036:
	;
	goto L1037
L1037:
	;
	v6308 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[4]))
	if v6308 != 0 {
		goto L1045
	} else {
		goto L1046
	}
L1038:
	;
	v6250 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	v6254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6250+v6209<<(uint(int32(1))%32)))))
	if v6254 != 0 {
		goto L1040
	} else {
		goto L1041
	}
L1039:
	;
	goto L1037
L1040:
	;
	if v6209 != v5701 {
		v6209 = v6209 + int32(1)
		goto L1038
	} else {
		goto L1044
	}
L1041:
	;
	v6255 = F_hk_depth_search(m, v5703, v6209)
	mBase = m.M
	v6256 = m.ExcPending
	if v6256 != 0 {
		goto L1
	} else {
		goto L1042
	}
L1042:
	;
	if v6255 == int32(0) {
		goto L1040
	} else {
		goto L1043
	}
L1043:
	;
	v6259 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5703)+12)) = v6259 + int32(1)
	goto L1040
L1044:
	;
	goto L1039
L1045:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6310 = m.ExcPending
	if v6310 != 0 {
		goto L1
	} else {
		goto L1048
	}
L1046:
	;
	goto L1047
L1047:
	;
	v6311 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+28))
	v6312 = *(*int32)(unsafe.Add(mBase, uint32(v5703)))
	v6313 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+24))
	v6314 = int32(_a_F_subquery_planner_32)
	*(*uint16)(unsafe.Add(mBase, uint32(v6313))) = uint16(v6314)
	if int32(0) < v6312 {
		v5743 = v6312
		v5750 = v6313
		v5756 = v6311
		goto L994
	} else {
		goto L1049
	}
L1048:
	;
	goto L1047
L1049:
	;
	goto L995
L1050:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_33), int32(0))
	mBase = m.M
	v6366 = m.ExcPending
	if v6366 != 0 {
		goto L1
	} else {
		goto L1051
	}
L1051:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_34), int32(45), int32(_a_F_subquery_planner_35))
	mBase = m.M
	v6371 = m.ExcPending
	if v6371 != 0 {
		goto L1
	} else {
		goto L1052
	}
L1052:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1053:
	;
	if v5701 <= int32(0) {
		goto L1055
	} else {
		goto L1056
	}
L1054:
	;
	if int32(0) < v4941 {
		goto L1077
	} else {
		goto L1078
	}
L1055:
	;
	v6379 = F_palloc0(m, int32(4))
	mBase = m.M
	v6380 = m.ExcPending
	if v6380 != 0 {
		goto L1
	} else {
		goto L1058
	}
L1056:
	;
	goto L1057
L1057:
	;
	v6381 = int32(2)
	if v5660 <= v6381 {
		goto L1059
	} else {
		goto L1060
	}
L1058:
	;
	v6537 = v6379
	v6550 = v5699
	goto L1054
L1059:
	;
	v6384 = v6381
	goto L1061
L1060:
	;
	v6384 = v5660
	goto L1061
L1061:
	;
	v6393 = int32(1)
	v6406 = v5699
	goto L1062
L1062:
	;
	v6428 = v6393 << (uint(int32(1)) % 32)
	v6429 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	v6431 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6428+v6429))))
	v6432 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+20))
	v6434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6432+v6428))))
	v6435 = int32(0)
	if base.B2i32(v6434 <= v6435)|base.B2i32(v6393 <= v6434) == v6435 {
		goto L1065
	} else {
		goto L1066
	}
L1063:
	;
	v6470 = F_palloc0(m, v6458<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v6471 = m.ExcPending
	if v6471 != 0 {
		goto L1
	} else {
		goto L1072
	}
L1064:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6374+v6393<<(uint(int32(2))%32)))) = v6457
	v6464 = v6393 + int32(1)
	if v6464 != v6384 {
		v6393 = v6464
		v6406 = v6458
		goto L1062
	} else {
		goto L1071
	}
L1065:
	;
	v6444 = *(*int32)(unsafe.Add(mBase, uint32(v6374+v6434<<(uint(int32(2))%32))))
	v6457 = v6444
	v6458 = v6406
	goto L1064
L1066:
	;
	goto L1067
L1067:
	;
	v6445 = int32(0)
	if base.B2i32(v6431 <= v6445)|base.B2i32(v6393 <= v6431) == v6445 {
		goto L1068
	} else {
		goto L1069
	}
L1068:
	;
	v6454 = *(*int32)(unsafe.Add(mBase, uint32(v6374+v6431<<(uint(int32(2))%32))))
	v6457 = v6454
	v6458 = v6406
	goto L1064
L1069:
	;
	goto L1070
L1070:
	;
	v6456 = v6406 + int32(1)
	v6457 = v6456
	v6458 = v6456
	goto L1064
L1071:
	;
	goto L1063
L1072:
	;
	v6481 = int32(1)
	goto L1073
L1073:
	;
	v6514 = int32(2)
	v6515 = v6481 << (uint(v6514) % 32)
	v6517 = *(*int32)(unsafe.Add(mBase, uint32(v6374+v6515)))
	v6520 = v6470 + v6517<<(uint(v6514)%32)
	v6521 = *(*int32)(unsafe.Add(mBase, uint32(v6520)))
	v6523 = *(*int32)(unsafe.Add(mBase, uint32(v6515+v5021)))
	v6524 = F_list_concat(m, v6521, v6523)
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L1
	} else {
		goto L1075
	}
L1074:
	;
	v6537 = v6470
	v6550 = v6458
	goto L1054
L1075:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6520))) = v6524
	v6528 = v6481 + int32(1)
	if v6528 <= v5701 {
		v6481 = v6528
		goto L1073
	} else {
		goto L1076
	}
L1076:
	;
	goto L1074
L1077:
	;
	v6574 = *(*int32)(unsafe.Add(mBase, uint32(v6537)+4))
	v6592 = v6574
	v6593 = v4941
	goto L1080
L1078:
	;
	goto L1079
L1079:
	;
	v6665 = int32(0)
	if v6665 < v6550 {
		goto L1084
	} else {
		goto L1085
	}
L1080:
	;
	v6617 = F_lcons(m, int32(0), v6592)
	mBase = m.M
	v6618 = m.ExcPending
	if v6618 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1081:
	;
	goto L1079
L1082:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6537)+4)) = v6617
	v6620 = int32(1)
	if base.Ui32(v6620) < base.Ui32(v6593) {
		v6592 = v6617
		v6593 = v6593 - v6620
		goto L1080
	} else {
		goto L1083
	}
L1083:
	;
	goto L1081
L1084:
	;
	v6676 = int32(1)
	v6681 = v6665
	goto L1087
L1085:
	;
	v6731 = v6665
	goto L1086
L1086:
	;
	v6759 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+16))
	F_pfree(m, v6759)
	mBase = m.M
	v6761 = m.ExcPending
	if v6761 != 0 {
		goto L1
	} else {
		goto L1091
	}
L1087:
	;
	v6712 = *(*int32)(unsafe.Add(mBase, uint32(v6537+v6676<<(uint(int32(2))%32))))
	v6713 = F_lappend(m, v6681, v6712)
	mBase = m.M
	v6714 = m.ExcPending
	if v6714 != 0 {
		goto L1
	} else {
		goto L1089
	}
L1088:
	;
	v6731 = v6713
	goto L1086
L1089:
	;
	v6716 = v6676 + int32(1)
	if v6716 <= v6550 {
		v6676 = v6716
		v6681 = v6713
		goto L1087
	} else {
		goto L1090
	}
L1090:
	;
	goto L1088
L1091:
	;
	v6762 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+20))
	F_pfree(m, v6762)
	mBase = m.M
	v6764 = m.ExcPending
	if v6764 != 0 {
		goto L1
	} else {
		goto L1092
	}
L1092:
	;
	v6765 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+24))
	F_pfree(m, v6765)
	mBase = m.M
	v6767 = m.ExcPending
	if v6767 != 0 {
		goto L1
	} else {
		goto L1093
	}
L1093:
	;
	v6768 = *(*int32)(unsafe.Add(mBase, uint32(v5703)+28))
	F_pfree(m, v6768)
	mBase = m.M
	v6770 = m.ExcPending
	if v6770 != 0 {
		goto L1
	} else {
		goto L1094
	}
L1094:
	;
	F_pfree(m, v5703)
	mBase = m.M
	v6772 = m.ExcPending
	if v6772 != 0 {
		goto L1
	} else {
		goto L1095
	}
L1095:
	;
	F_pfree(m, v6537)
	mBase = m.M
	v6774 = m.ExcPending
	if v6774 != 0 {
		goto L1
	} else {
		goto L1096
	}
L1096:
	;
	F_pfree(m, v6374)
	mBase = m.M
	v6776 = m.ExcPending
	if v6776 != 0 {
		goto L1
	} else {
		goto L1097
	}
L1097:
	;
	if int32(0) < v5701 {
		goto L1099
	} else {
		goto L1100
	}
L1098:
	;
	F_pfree(m, v5023)
	mBase = m.M
	v6935 = m.ExcPending
	if v6935 != 0 {
		goto L1
	} else {
		goto L1119
	}
L1099:
	;
	v6787 = int32(1)
	goto L1102
L1100:
	;
	goto L1101
L1101:
	;
	F_pfree(m, v5025)
	mBase = m.M
	v6888 = m.ExcPending
	if v6888 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1102:
	;
	v6824 = *(*int32)(unsafe.Add(mBase, uint32(v5025+v6787<<(uint(int32(2))%32))))
	if v6824 != 0 {
		goto L1104
	} else {
		goto L1105
	}
L1103:
	;
	F_pfree(m, v5025)
	mBase = m.M
	v6831 = m.ExcPending
	if v6831 != 0 {
		goto L1
	} else {
		goto L1109
	}
L1104:
	;
	F_pfree(m, v6824)
	mBase = m.M
	v6826 = m.ExcPending
	if v6826 != 0 {
		goto L1
	} else {
		goto L1107
	}
L1105:
	;
	goto L1106
L1106:
	;
	v6828 = v6787 + int32(1)
	if v6828 <= v5701 {
		v6787 = v6828
		goto L1102
	} else {
		goto L1108
	}
L1107:
	;
	goto L1106
L1108:
	;
	goto L1103
L1109:
	;
	F_pfree(m, v5029)
	mBase = m.M
	v6833 = m.ExcPending
	if v6833 != 0 {
		goto L1
	} else {
		goto L1110
	}
L1110:
	;
	F_pfree(m, v5021)
	mBase = m.M
	v6835 = m.ExcPending
	if v6835 != 0 {
		goto L1
	} else {
		goto L1111
	}
L1111:
	;
	v6844 = int32(1)
	goto L1112
L1112:
	;
	v6881 = *(*int32)(unsafe.Add(mBase, uint32(v5023+v6844<<(uint(int32(2))%32))))
	F_bms_free(m, v6881)
	mBase = m.M
	v6883 = m.ExcPending
	if v6883 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1113:
	;
	goto L1098
L1114:
	;
	v6885 = v6844 + int32(1)
	if v6885 != v5660 {
		v6844 = v6885
		goto L1112
	} else {
		goto L1115
	}
L1115:
	;
	goto L1113
L1116:
	;
	F_pfree(m, v5029)
	mBase = m.M
	v6890 = m.ExcPending
	if v6890 != 0 {
		goto L1
	} else {
		goto L1117
	}
L1117:
	;
	F_pfree(m, v5021)
	mBase = m.M
	v6892 = m.ExcPending
	if v6892 != 0 {
		goto L1
	} else {
		goto L1118
	}
L1118:
	;
	goto L1098
L1119:
	;
	v6949 = v6731
	goto L716
L1120:
	;
	v6979 = *(*int32)(unsafe.Add(mBase, uint32(v6949)+4))
	if v6979 <= int32(0) {
		v8022 = v47
		v8026 = l5
		v8027 = v3797
		v8032 = v3799
		v8035 = v4573
		v8037 = v4571
		v8041 = v44
		v8043 = v7
		v8050 = v3913
		v8054 = v3914
		v8055 = v3915
		goto L715
	} else {
		goto L1121
	}
L1121:
	;
	v6989 = v47
	v6993 = l5
	v6994 = v3797
	v6996 = v6949
	v6999 = v3799
	v7002 = v4573
	v7003 = int32(0)
	v7004 = v4571
	v7008 = v44
	v7010 = v7
	v7017 = v3913
	v7021 = v3914
	v7022 = v3915
	goto L1122
L1122:
	;
	v7024 = *(*int32)(unsafe.Add(mBase, uint32(v6996)+12))
	v7028 = *(*int32)(unsafe.Add(mBase, uint32(v7024+v7003<<(uint(int32(2))%32))))
	v7030 = F_palloc0(m, int32(32))
	mBase = m.M
	v7031 = m.ExcPending
	if v7031 != 0 {
		goto L1
	} else {
		goto L1124
	}
L1123:
	;
	v8022 = v6989
	v8026 = v6993
	v8027 = v6994
	v8032 = v6999
	v8035 = v7002
	v8037 = v7004
	v8041 = v7008
	v8043 = v7010
	v8050 = v7017
	v8054 = v7021
	v8055 = v7022
	goto L715
L1124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7030))) = int32(309)
	v7034 = int32(0)
	v7036 = *(*int32)(unsafe.Add(mBase, uint32(v6996)+4))
	if v7036 == int32(1) {
		goto L1125
	} else {
		goto L1126
	}
L1125:
	;
	v7039 = *(*int32)(unsafe.Add(mBase, uint32(v7004)+124))
	v7040 = v7039
	goto L1127
L1126:
	;
	v7040 = v7034
	goto L1127
L1127:
	;
	v7041 = int32(0)
	if v7028 == v7041 {
		v7545 = v7041
		v7546 = v7034
		goto L1128
	} else {
		goto L1129
	}
L1128:
	;
	F_list_free(m, v7545)
	mBase = m.M
	v7579 = m.ExcPending
	if v7579 != 0 {
		goto L1
	} else {
		goto L1196
	}
L1129:
	;
	v7044 = int32(0)
	v7045 = *(*int32)(unsafe.Add(mBase, uint32(v7028)+4))
	if v7045 <= v7044 {
		v7545 = v7041
		v7546 = v7034
		goto L1128
	} else {
		goto L1130
	}
L1130:
	;
	v7048 = v7040
	v7056 = v7041
	v7057 = v7034
	v7060 = v7044
	goto L1131
L1131:
	;
	v7089 = *(*int32)(unsafe.Add(mBase, uint32(v7028)+12))
	v7093 = *(*int32)(unsafe.Add(mBase, uint32(v7089+v7060<<(uint(int32(2))%32))))
	v7094 = int32(0)
	if v7056 != 0 {
		goto L1135
	} else {
		goto L1136
	}
L1133:
	;
	v7414 = F_palloc0(m, int32(16))
	mBase = m.M
	v7415 = m.ExcPending
	if v7415 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1134:
	;
	v7412 = v7332
	goto L1133
L1135:
	;
	v7096 = int32(0)
	if v7093 == v7096 {
		v7412 = v7096
		goto L1133
	} else {
		goto L1138
	}
L1136:
	;
	goto L1137
L1137:
	;
	v7290 = int32(0)
	if v7093 == v7290 {
		v7412 = v7290
		goto L1133
	} else {
		goto L1152
	}
L1138:
	;
	v7099 = *(*int32)(unsafe.Add(mBase, uint32(v7093)+4))
	if v7099 <= int32(0) {
		v7332 = v7094
		goto L1134
	} else {
		goto L1139
	}
L1139:
	;
	v7104 = v7094
	v7109 = v7094
	goto L1140
L1140:
	;
	v7143 = *(*int32)(unsafe.Add(mBase, uint32(v7093)+12))
	v7147 = *(*int32)(unsafe.Add(mBase, uint32(v7143+v7109<<(uint(int32(2))%32))))
	v7148 = *(*int32)(unsafe.Add(mBase, uint32(v7056)+4))
	if int32(0) < v7148 {
		goto L1143
	} else {
		goto L1144
	}
L1141:
	;
	v7332 = v7247
	goto L1134
L1142:
	;
	v7287 = v7109 + int32(1)
	v7288 = *(*int32)(unsafe.Add(mBase, uint32(v7093)+4))
	if v7287 < v7288 {
		v7104 = v7247
		v7109 = v7287
		goto L1140
	} else {
		goto L1151
	}
L1143:
	;
	v7151 = *(*int32)(unsafe.Add(mBase, uint32(v7056)+12))
	v7154 = int32(0)
	goto L1146
L1144:
	;
	goto L1145
L1145:
	;
	v7243 = F_lappend_int(m, v7104, v7147)
	mBase = m.M
	v7244 = m.ExcPending
	if v7244 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1146:
	;
	v7197 = *(*int32)(unsafe.Add(mBase, uint32(v7151+v7154<<(uint(int32(2))%32))))
	if v7197 == v7147 {
		v7247 = v7104
		goto L1142
	} else {
		goto L1148
	}
L1147:
	;
	goto L1145
L1148:
	;
	v7200 = v7154 + int32(1)
	if v7148 != v7200 {
		v7154 = v7200
		goto L1146
	} else {
		goto L1149
	}
L1149:
	;
	goto L1147
L1150:
	;
	v7247 = v7243
	goto L1142
L1151:
	;
	goto L1141
L1152:
	;
	v7293 = *(*int32)(unsafe.Add(mBase, uint32(v7093)))
	v7296 = int32(8)
	v7297 = *(*int32)(unsafe.Add(mBase, uint32(v7093)+4))
	v7299 = v7297 + int32(4)
	if v7299 <= v7296 {
		goto L1153
	} else {
		goto L1154
	}
L1153:
	;
	v7302 = v7296
	goto L1155
L1154:
	;
	v7302 = v7299
	goto L1155
L1155:
	;
	if v7302&(v7302-int32(1)) != 0 {
		goto L1156
	} else {
		goto L1157
	}
L1156:
	;
	v7309 = int32(1) << (uint(int32(32)-base.I32_clz(v7302)) % 32)
	goto L1158
L1157:
	;
	v7309 = v7302
	goto L1158
L1158:
	;
	v7311 = v7309 - int32(4)
	v7316 = F_palloc(m, v7311<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v7317 = m.ExcPending
	if v7317 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7316)+8)) = v7311
	*(*int32)(unsafe.Add(mBase, uint32(v7316)+4)) = v7297
	*(*int32)(unsafe.Add(mBase, uint32(v7316))) = v7293
	v7322 = v7316 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v7316)+12)) = v7322
	v7325 = v7297 << (uint(int32(2)) % 32)
	if v7325 == int32(0) {
		v7332 = v7316
		goto L1134
	} else {
		goto L1160
	}
L1160:
	;
	v7328 = *(*int32)(unsafe.Add(mBase, uint32(v7093)+12))
	base.MemoryCopy(m, v7322, v7328, v7325)
	v7412 = v7316
	goto L1133
L1161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7414))) = int32(308)
	v7425 = v7412
	v7426 = v7056
	goto L1162
L1162:
	;
	if v7048 != 0 {
		goto L1164
	} else {
		goto L1165
	}
L1164:
	;
	v7459 = *(*int32)(unsafe.Add(mBase, uint32(v7048)+4))
	v7461 = v7459
	goto L1166
L1165:
	;
	v7461 = int32(0)
	goto L1166
L1166:
	;
	if v7426 == int32(0) {
		goto L1170
	} else {
		goto L1171
	}
L1167:
	;
	v7533 = F_lappend_int(m, v7426, v7478)
	mBase = m.M
	v7534 = m.ExcPending
	if v7534 != 0 {
		goto L1
	} else {
		goto L1194
	}
L1168:
	;
	v7522 = F_list_concat(m, v7426, v7425)
	mBase = m.M
	v7523 = m.ExcPending
	if v7523 != 0 {
		goto L1
	} else {
		goto L1190
	}
L1169:
	;
	v7473 = *(*int32)(unsafe.Add(mBase, uint32(v7048)+12))
	v7477 = *(*int32)(unsafe.Add(mBase, uint32(v7473+v7472<<(uint(int32(2))%32))))
	v7478 = *(*int32)(unsafe.Add(mBase, uint32(v7477)+4))
	v7479 = int32(0)
	if v7425 == v7479 {
		goto L1177
	} else {
		goto L1178
	}
L1170:
	;
	if v7461 <= int32(0) {
		v7519 = v7048
		goto L1168
	} else {
		goto L1173
	}
L1171:
	;
	goto L1172
L1172:
	;
	v7469 = *(*int32)(unsafe.Add(mBase, uint32(v7426)+4))
	if base.B2i32(v7425 == int32(0))|base.B2i32(v7461 <= v7469) != 0 {
		v7519 = v7048
		goto L1168
	} else {
		goto L1175
	}
L1173:
	;
	if v7425 != 0 {
		v7472 = int32(0)
		goto L1169
	} else {
		goto L1174
	}
L1174:
	;
	v7519 = v7048
	goto L1168
L1175:
	;
	v7472 = v7469
	goto L1169
L1176:
	;
	if v7517 != 0 {
		goto L1167
	} else {
		goto L1189
	}
L1177:
	;
	v7517 = int32(0)
	goto L1176
L1178:
	;
	goto L1179
L1179:
	;
	v7485 = *(*int32)(unsafe.Add(mBase, uint32(v7425)+4))
	if v7485 <= int32(0) {
		v7511 = v7479
		goto L1180
	} else {
		goto L1181
	}
L1180:
	;
	v7517 = v7511
	goto L1176
L1181:
	;
	v7488 = int32(0)
	if v7488 < v7485 {
		goto L1182
	} else {
		goto L1183
	}
L1182:
	;
	v7491 = v7485
	goto L1184
L1183:
	;
	v7491 = v7488
	goto L1184
L1184:
	;
	v7492 = *(*int32)(unsafe.Add(mBase, uint32(v7425)+12))
	v7494 = int32(0)
	goto L1185
L1185:
	;
	v7502 = *(*int32)(unsafe.Add(mBase, uint32(v7492+v7494<<(uint(int32(2))%32))))
	v7503 = base.B2i32(v7502 == v7478)
	if v7502 == v7478 {
		v7511 = v7503
		goto L1180
	} else {
		goto L1187
	}
L1186:
	;
	v7511 = v7503
	goto L1180
L1187:
	;
	v7505 = v7494 + int32(1)
	if v7505 != v7491 {
		v7494 = v7505
		goto L1185
	} else {
		goto L1188
	}
L1188:
	;
	goto L1186
L1189:
	;
	v7519 = int32(0)
	goto L1168
L1190:
	;
	v7524 = F_list_copy(m, v7522)
	mBase = m.M
	v7525 = m.ExcPending
	if v7525 != 0 {
		goto L1
	} else {
		goto L1191
	}
L1191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7414)+4)) = v7524
	v7527 = F_lcons(m, v7414, v7057)
	mBase = m.M
	v7528 = m.ExcPending
	if v7528 != 0 {
		goto L1
	} else {
		goto L1192
	}
L1192:
	;
	v7530 = v7060 + int32(1)
	v7531 = *(*int32)(unsafe.Add(mBase, uint32(v7028)+4))
	if v7530 < v7531 {
		v7048 = v7519
		v7056 = v7522
		v7057 = v7527
		v7060 = v7530
		goto L1131
	} else {
		goto L1193
	}
L1193:
	;
	v7545 = v7522
	v7546 = v7527
	goto L1128
L1194:
	;
	v7535 = F_list_delete_ptr(m, v7425, v7478)
	mBase = m.M
	v7536 = m.ExcPending
	if v7536 != 0 {
		goto L1
	} else {
		goto L1195
	}
L1195:
	;
	v7425 = v7535
	v7426 = v7533
	goto L1162
L1196:
	;
	v7580 = int32(0)
	v7581 = *(*int32)(unsafe.Add(mBase, uint32(v7546)+12))
	v7582 = *(*int32)(unsafe.Add(mBase, uint32(v7581)))
	v7583 = *(*int32)(unsafe.Add(mBase, uint32(v7582)+4))
	if v7583 == v7580 {
		v7653 = v7580
		goto L1197
	} else {
		goto L1198
	}
L1197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7030)+4)) = v7653
	v7688 = *(*int32)(unsafe.Add(mBase, uint32(v7582)+4))
	if v7688 != 0 {
		goto L1205
	} else {
		goto L1206
	}
L1198:
	;
	v7586 = *(*int32)(unsafe.Add(mBase, uint32(v7583)+4))
	if v7586 <= int32(0) {
		v7653 = v7580
		goto L1197
	} else {
		goto L1199
	}
L1199:
	;
	v7589 = *(*int32)(unsafe.Add(mBase, uint32(v6989)+4))
	v7598 = v7580
	v7599 = int32(0)
	goto L1200
L1200:
	;
	v7632 = *(*int32)(unsafe.Add(mBase, uint32(v7583)+12))
	v7636 = *(*int32)(unsafe.Add(mBase, uint32(v7632+v7599<<(uint(int32(2))%32))))
	v7637 = *(*int32)(unsafe.Add(mBase, uint32(v7589)+100))
	v7638 = F_get_sortgroupref_clause(m, v7636, v7637)
	mBase = m.M
	v7639 = m.ExcPending
	if v7639 != 0 {
		goto L1
	} else {
		goto L1202
	}
L1201:
	;
	v7653 = v7640
	goto L1197
L1202:
	;
	v7640 = F_lappend(m, v7598, v7638)
	mBase = m.M
	v7641 = m.ExcPending
	if v7641 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1203:
	;
	v7643 = v7599 + int32(1)
	v7644 = *(*int32)(unsafe.Add(mBase, uint32(v7583)+4))
	if v7643 < v7644 {
		v7598 = v7640
		v7599 = v7643
		goto L1200
	} else {
		goto L1204
	}
L1204:
	;
	goto L1201
L1205:
	;
	v7689 = *(*int32)(unsafe.Add(mBase, uint32(v7002)+24))
	v7690 = F_bms_overlap_list(m, v7689, v7688)
	mBase = m.M
	v7691 = m.ExcPending
	if v7691 != 0 {
		goto L1
	} else {
		goto L1208
	}
L1206:
	;
	v7699 = v7653
	goto L1207
L1207:
	;
	v7700 = *(*int32)(unsafe.Add(mBase, uint32(v7002)+32))
	if v7699 == int32(0) {
		goto L1212
	} else {
		goto L1213
	}
L1208:
	;
	if v7690 == int32(0) {
		goto L1209
	} else {
		goto L1210
	}
L1209:
	;
	v7694 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7030)+24)) = uint8(v7694)
	*(*uint8)(unsafe.Add(mBase, uint32(v7002)+16)) = uint8(v7694)
	goto L1211
L1210:
	;
	goto L1211
L1211:
	;
	v7698 = *(*int32)(unsafe.Add(mBase, uint32(v7030)+4))
	v7699 = v7698
	goto L1207
L1212:
	;
	v7803 = int32(0)
	v7805 = *(*int32)(unsafe.Add(mBase, uint32(v7546)+4))
	if v7803 < v7805 {
		goto L1218
	} else {
		goto L1219
	}
L1213:
	;
	v7703 = int32(0)
	v7704 = *(*int32)(unsafe.Add(mBase, uint32(v7699)+4))
	if v7704 <= v7703 {
		goto L1212
	} else {
		goto L1214
	}
L1214:
	;
	v7715 = v7703
	goto L1215
L1215:
	;
	v7748 = *(*int32)(unsafe.Add(mBase, uint32(v7699)+12))
	v7749 = int32(2)
	v7752 = *(*int32)(unsafe.Add(mBase, uint32(v7748+v7715<<(uint(v7749)%32))))
	v7753 = *(*int32)(unsafe.Add(mBase, uint32(v7752)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7700+v7753<<(uint(v7749)%32)))) = v7715
	v7759 = v7715 + int32(1)
	v7760 = *(*int32)(unsafe.Add(mBase, uint32(v7699)+4))
	if v7759 < v7760 {
		v7715 = v7759
		goto L1215
	} else {
		goto L1217
	}
L1216:
	;
	goto L1212
L1217:
	;
	goto L1216
L1218:
	;
	v7813 = v7803
	v7820 = v7803
	goto L1221
L1219:
	;
	v7970 = v7803
	goto L1220
L1220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7030)+12)) = v7546
	*(*int32)(unsafe.Add(mBase, uint32(v7030)+8)) = v7970
	v8008 = *(*int32)(unsafe.Add(mBase, uint32(v7002)))
	v8009 = F_lappend(m, v8008, v7030)
	mBase = m.M
	v8010 = m.ExcPending
	if v8010 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1221:
	;
	v7849 = int32(0)
	v7850 = *(*int32)(unsafe.Add(mBase, uint32(v7546)+12))
	v7854 = *(*int32)(unsafe.Add(mBase, uint32(v7850+v7820<<(uint(int32(2))%32))))
	v7855 = *(*int32)(unsafe.Add(mBase, uint32(v7854)+4))
	if v7855 == v7849 {
		v7926 = v7849
		goto L1223
	} else {
		goto L1224
	}
L1222:
	;
	v7970 = v7959
	goto L1220
L1223:
	;
	v7959 = F_lappend(m, v7813, v7926)
	mBase = m.M
	v7960 = m.ExcPending
	if v7960 != 0 {
		goto L1
	} else {
		goto L1230
	}
L1224:
	;
	v7858 = int32(0)
	v7859 = *(*int32)(unsafe.Add(mBase, uint32(v7855)+4))
	if v7859 <= v7858 {
		v7926 = v7849
		goto L1223
	} else {
		goto L1225
	}
L1225:
	;
	v7869 = v7858
	v7870 = v7849
	goto L1226
L1226:
	;
	v7903 = *(*int32)(unsafe.Add(mBase, uint32(v7855)+12))
	v7904 = int32(2)
	v7907 = *(*int32)(unsafe.Add(mBase, uint32(v7903+v7869<<(uint(v7904)%32))))
	v7911 = *(*int32)(unsafe.Add(mBase, uint32(v7700+v7907<<(uint(v7904)%32))))
	v7912 = F_lappend_int(m, v7870, v7911)
	mBase = m.M
	v7913 = m.ExcPending
	if v7913 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1227:
	;
	v7926 = v7912
	goto L1223
L1228:
	;
	v7915 = v7869 + int32(1)
	v7916 = *(*int32)(unsafe.Add(mBase, uint32(v7855)+4))
	if v7915 < v7916 {
		v7869 = v7915
		v7870 = v7912
		goto L1226
	} else {
		goto L1229
	}
L1229:
	;
	goto L1227
L1230:
	;
	v7962 = v7820 + int32(1)
	v7963 = *(*int32)(unsafe.Add(mBase, uint32(v7546)+4))
	if v7962 < v7963 {
		v7813 = v7959
		v7820 = v7962
		goto L1221
	} else {
		goto L1231
	}
L1231:
	;
	goto L1222
L1232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7002))) = v8009
	v8013 = v7003 + int32(1)
	v8014 = *(*int32)(unsafe.Add(mBase, uint32(v6996)+4))
	if v8013 < v8014 {
		v7003 = v8013
		goto L1122
	} else {
		goto L1233
	}
L1233:
	;
	goto L1123
L1234:
	;
	v8060 = *(*int32)(unsafe.Add(mBase, uint32(v8035)+32))
	v8061 = *(*int32)(unsafe.Add(mBase, uint32(v8037)+100))
	if v8061 == int32(0) {
		goto L1235
	} else {
		goto L1236
	}
L1235:
	;
	v8164 = int32(0)
	v8165 = *(*int32)(unsafe.Add(mBase, uint32(v8057)+4))
	if v8164 < v8165 {
		goto L1241
	} else {
		goto L1242
	}
L1236:
	;
	v8064 = *(*int32)(unsafe.Add(mBase, uint32(v8061)+4))
	if v8064 <= int32(0) {
		goto L1235
	} else {
		goto L1237
	}
L1237:
	;
	v8075 = int32(0)
	goto L1238
L1238:
	;
	v8109 = *(*int32)(unsafe.Add(mBase, uint32(v8061)+12))
	v8110 = int32(2)
	v8113 = *(*int32)(unsafe.Add(mBase, uint32(v8109+v8075<<(uint(v8110)%32))))
	v8114 = *(*int32)(unsafe.Add(mBase, uint32(v8113)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8060+v8114<<(uint(v8110)%32)))) = v8075
	v8120 = v8075 + int32(1)
	v8121 = *(*int32)(unsafe.Add(mBase, uint32(v8061)+4))
	if v8120 < v8121 {
		v8075 = v8120
		goto L1238
	} else {
		goto L1240
	}
L1239:
	;
	goto L1235
L1240:
	;
	goto L1239
L1241:
	;
	v8174 = v8164
	v8178 = int32(0)
	goto L1244
L1242:
	;
	v8332 = v8164
	goto L1243
L1243:
	;
	v8368 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8035)+16)) = uint8(v8368)
	*(*int32)(unsafe.Add(mBase, uint32(v8035)+4)) = v8332
	v8377 = v8022
	v8381 = v8026
	v8382 = v8027
	v8387 = v8032
	v8390 = v8035
	v8396 = v8041
	v8398 = v8043
	v8405 = v8050
	v8409 = v8054
	v8410 = v8055
	goto L714
L1244:
	;
	v8210 = *(*int32)(unsafe.Add(mBase, uint32(v8057)+12))
	v8214 = *(*int32)(unsafe.Add(mBase, uint32(v8210+v8178<<(uint(int32(2))%32))))
	v8215 = *(*int32)(unsafe.Add(mBase, uint32(v8214)+4))
	if v8215 == int32(0) {
		goto L1247
	} else {
		goto L1248
	}
L1245:
	;
	v8332 = v8321
	goto L1243
L1246:
	;
	v8321 = F_lappend(m, v8174, v8288)
	mBase = m.M
	v8322 = m.ExcPending
	if v8322 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1247:
	;
	v8288 = int32(0)
	goto L1246
L1248:
	;
	goto L1249
L1249:
	;
	v8219 = int32(0)
	v8221 = *(*int32)(unsafe.Add(mBase, uint32(v8215)+4))
	if v8221 <= v8219 {
		v8288 = v8219
		goto L1246
	} else {
		goto L1250
	}
L1250:
	;
	v8231 = v8219
	v8232 = v8219
	goto L1251
L1251:
	;
	v8265 = *(*int32)(unsafe.Add(mBase, uint32(v8215)+12))
	v8266 = int32(2)
	v8269 = *(*int32)(unsafe.Add(mBase, uint32(v8265+v8231<<(uint(v8266)%32))))
	v8273 = *(*int32)(unsafe.Add(mBase, uint32(v8060+v8269<<(uint(v8266)%32))))
	v8274 = F_lappend_int(m, v8232, v8273)
	mBase = m.M
	v8275 = m.ExcPending
	if v8275 != 0 {
		goto L1
	} else {
		goto L1253
	}
L1252:
	;
	v8288 = v8274
	goto L1246
L1253:
	;
	v8277 = v8231 + int32(1)
	v8278 = *(*int32)(unsafe.Add(mBase, uint32(v8215)+4))
	if v8277 < v8278 {
		v8231 = v8277
		v8232 = v8274
		goto L1251
	} else {
		goto L1254
	}
L1254:
	;
	goto L1252
L1255:
	;
	v8324 = v8178 + int32(1)
	v8325 = *(*int32)(unsafe.Add(mBase, uint32(v8057)+4))
	if v8324 < v8325 {
		v8174 = v8321
		v8178 = v8324
		goto L1244
	} else {
		goto L1256
	}
L1256:
	;
	goto L1245
L1257:
	;
	v9131 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+136))
	if v9131 == int32(0) {
		v9308 = v9090
		goto L1350
	} else {
		goto L1351
	}
L1258:
	;
	v9010 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+72))
	v9012 = F_pull_var_clause(m, v9010, int32(16))
	mBase = m.M
	v9013 = m.ExcPending
	if v9013 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1259:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8959 = m.ExcPending
	if v8959 != 0 {
		goto L1
	} else {
		goto L1329
	}
L1260:
	;
	v8424 = *(*int32)(unsafe.Add(mBase, uint32(v8422)+12))
	v8430 = *(*int32)(unsafe.Add(mBase, uint32(v8424+v8423<<(uint(int32(2))%32)-int32(4))))
	v8431 = *(*int32)(unsafe.Add(mBase, uint32(v8430)+12))
	if v8431 != 0 {
		goto L1259
	} else {
		goto L1263
	}
L1261:
	;
	v8437 = v8412
	v8438 = int32(0)
	goto L1262
L1262:
	;
	v8439 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+76))
	switch v8421 - int32(2) {
	case 0:
		goto L1266
	case 1:
		goto L1267
	default:
		goto L1265
	}
L1263:
	;
	v8432 = *(*int32)(unsafe.Add(mBase, uint32(v8430)+16))
	v8434 = F_table_open(m, v8432, int32(0))
	mBase = m.M
	v8435 = m.ExcPending
	if v8435 != 0 {
		goto L1
	} else {
		goto L1264
	}
L1264:
	;
	v8437 = v8430
	v8438 = v8434
	goto L1262
L1265:
	;
	if base.B2i32(int32(1)<<(uint(v8421)%32)&int32(52) == int32(0))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v8421)) != 0 {
		v9090 = v8439
		goto L1257
	} else {
		goto L1279
	}
L1266:
	;
	if v8439 == int32(0) {
		v8515 = v8412
		goto L1269
	} else {
		goto L1270
	}
L1267:
	;
	v8442 = F_expand_insert_targetlist(m, v8377, v8439, v8438)
	mBase = m.M
	v8443 = m.ExcPending
	if v8443 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1268:
	;
	v9090 = v8442
	goto L1257
L1269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+268)) = v8515
	goto L1265
L1270:
	;
	v8447 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+4))
	if v8447 <= int32(0) {
		v8515 = v8412
		goto L1269
	} else {
		goto L1271
	}
L1271:
	;
	v8451 = int32(1)
	v8452 = v8412
	v8455 = v8412
	goto L1272
L1272:
	;
	v8491 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+12))
	v8495 = *(*int32)(unsafe.Add(mBase, uint32(v8491+v8452<<(uint(int32(2))%32))))
	v8496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8495)+26)))
	if v8496 == int32(0) {
		goto L1274
	} else {
		goto L1275
	}
L1273:
	;
	v8515 = v8502
	goto L1269
L1274:
	;
	v8499 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8495)+8)))
	v8500 = F_lappend_int(m, v8455, v8499)
	mBase = m.M
	v8501 = m.ExcPending
	if v8501 != 0 {
		goto L1
	} else {
		goto L1277
	}
L1275:
	;
	v8502 = v8455
	goto L1276
L1276:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8495)+8)) = uint16(v8451)
	v8504 = int32(1)
	v8507 = v8452 + v8504
	v8508 = *(*int32)(unsafe.Add(mBase, uint32(v8439)+4))
	if v8507 < v8508 {
		v8451 = v8451 + v8504
		v8452 = v8507
		v8455 = v8502
		goto L1272
	} else {
		goto L1278
	}
L1277:
	;
	v8502 = v8500
	goto L1276
L1278:
	;
	goto L1273
L1279:
	;
	v8602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8437)+20)))
	if v8602 == int32(0) {
		goto L1281
	} else {
		goto L1282
	}
L1280:
	;
	v8614 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+64))
	if v8614 == int32(0) {
		v8969 = v8613
		goto L1258
	} else {
		goto L1287
	}
L1281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+264)) = v8439
	F_add_row_identity_columns(m, v8377, v8423, v8437, v8438)
	mBase = m.M
	v8607 = m.ExcPending
	if v8607 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1282:
	;
	goto L1283
L1283:
	;
	if v8421 != int32(5) {
		v9090 = v8439
		goto L1257
	} else {
		goto L1286
	}
L1284:
	;
	v8608 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	if v8421 == int32(5) {
		v8613 = v8608
		goto L1280
	} else {
		goto L1285
	}
L1285:
	;
	v9090 = v8608
	goto L1257
L1286:
	;
	v8613 = v8439
	goto L1280
L1287:
	;
	v8617 = *(*int32)(unsafe.Add(mBase, uint32(v8614)+4))
	if v8617 <= int32(0) {
		v8969 = v8613
		goto L1258
	} else {
		goto L1288
	}
L1288:
	;
	v8620 = v8613
	v8635 = v8412
	goto L1289
L1289:
	;
	v8661 = *(*int32)(unsafe.Add(mBase, uint32(v8614)+12))
	v8662 = int32(2)
	v8665 = *(*int32)(unsafe.Add(mBase, uint32(v8661+v8635<<(uint(v8662)%32))))
	v8666 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+8))
	switch v8666 - v8662 {
	case 0:
		goto L1292
	case 1:
		goto L1293
	default:
		goto L1291
	}
L1290:
	;
	v8969 = v8909
	goto L1258
L1291:
	;
	v8826 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+16))
	v8827 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+20))
	v8828 = F_list_concat_copy(m, v8826, v8827)
	mBase = m.M
	v8829 = m.ExcPending
	if v8829 != 0 {
		goto L1
	} else {
		goto L1308
	}
L1292:
	;
	v8673 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+20))
	if v8673 == int32(0) {
		goto L1296
	} else {
		goto L1297
	}
L1293:
	;
	v8669 = *(*int32)(unsafe.Add(mBase, uint32(v8665)+20))
	v8670 = F_expand_insert_targetlist(m, v8377, v8669, v8438)
	mBase = m.M
	v8671 = m.ExcPending
	if v8671 != 0 {
		goto L1
	} else {
		goto L1294
	}
L1294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8665)+20)) = v8670
	goto L1291
L1295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8665)+24)) = v8750
	goto L1291
L1296:
	;
	v8750 = int32(0)
	goto L1295
L1297:
	;
	goto L1298
L1298:
	;
	v8678 = int32(0)
	v8680 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+4))
	if v8680 <= v8678 {
		v8750 = v8678
		goto L1295
	} else {
		goto L1299
	}
L1299:
	;
	v8684 = int32(1)
	v8685 = v8678
	v8690 = v8678
	goto L1300
L1300:
	;
	v8724 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+12))
	v8728 = *(*int32)(unsafe.Add(mBase, uint32(v8724+v8685<<(uint(int32(2))%32))))
	v8729 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8728)+26)))
	if v8729 == int32(0) {
		goto L1302
	} else {
		goto L1303
	}
L1301:
	;
	v8750 = v8735
	goto L1295
L1302:
	;
	v8732 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8728)+8)))
	v8733 = F_lappend_int(m, v8690, v8732)
	mBase = m.M
	v8734 = m.ExcPending
	if v8734 != 0 {
		goto L1
	} else {
		goto L1305
	}
L1303:
	;
	v8735 = v8690
	goto L1304
L1304:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8728)+8)) = uint16(v8684)
	v8737 = int32(1)
	v8740 = v8685 + v8737
	v8741 = *(*int32)(unsafe.Add(mBase, uint32(v8673)+4))
	if v8740 < v8741 {
		v8684 = v8684 + v8737
		v8685 = v8740
		v8690 = v8735
		goto L1300
	} else {
		goto L1306
	}
L1305:
	;
	v8735 = v8733
	goto L1304
L1306:
	;
	goto L1301
L1307:
	;
	F_list_free(m, v8831)
	mBase = m.M
	v8951 = m.ExcPending
	if v8951 != 0 {
		goto L1
	} else {
		goto L1327
	}
L1308:
	;
	v8831 = F_pull_var_clause(m, v8828, int32(16))
	mBase = m.M
	v8832 = m.ExcPending
	if v8832 != 0 {
		goto L1
	} else {
		goto L1309
	}
L1309:
	;
	if v8831 == int32(0) {
		v8909 = v8620
		goto L1307
	} else {
		goto L1310
	}
L1310:
	;
	v8835 = int32(0)
	v8836 = *(*int32)(unsafe.Add(mBase, uint32(v8831)+4))
	if v8836 <= v8835 {
		v8909 = v8620
		goto L1307
	} else {
		goto L1311
	}
L1311:
	;
	v8839 = v8620
	v8840 = v8835
	goto L1312
L1312:
	;
	v8880 = *(*int32)(unsafe.Add(mBase, uint32(v8831)+12))
	v8884 = *(*int32)(unsafe.Add(mBase, uint32(v8880+v8840<<(uint(int32(2))%32))))
	v8885 = *(*int32)(unsafe.Add(mBase, uint32(v8884)))
	if v8885 == int32(6) {
		goto L1315
	} else {
		goto L1316
	}
L1313:
	;
	v8909 = v8904
	goto L1307
L1314:
	;
	v8906 = v8840 + int32(1)
	v8907 = *(*int32)(unsafe.Add(mBase, uint32(v8831)+4))
	if v8906 < v8907 {
		v8839 = v8904
		v8840 = v8906
		goto L1312
	} else {
		goto L1326
	}
L1315:
	;
	v8888 = *(*int32)(unsafe.Add(mBase, uint32(v8884)+4))
	if v8888 == v8423 {
		v8904 = v8839
		goto L1314
	} else {
		goto L1318
	}
L1316:
	;
	goto L1317
L1317:
	;
	v8890 = F_tlist_member(m, v8884, v8839)
	mBase = m.M
	v8891 = m.ExcPending
	if v8891 != 0 {
		goto L1
	} else {
		goto L1319
	}
L1318:
	;
	goto L1317
L1319:
	;
	if v8890 != 0 {
		v8904 = v8839
		goto L1314
	} else {
		goto L1320
	}
L1320:
	;
	if v8839 != 0 {
		goto L1321
	} else {
		goto L1322
	}
L1321:
	;
	v8892 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8839)+4)))
	v8896 = v8892 + int32(1)
	goto L1323
L1322:
	;
	v8896 = int32(1)
	goto L1323
L1323:
	;
	v8900 = F_makeTargetEntry(m, v8884, base.I32_extend16_s(v8896), int32(0), int32(1))
	mBase = m.M
	v8901 = m.ExcPending
	if v8901 != 0 {
		goto L1
	} else {
		goto L1324
	}
L1324:
	;
	v8902 = F_lappend(m, v8839, v8900)
	mBase = m.M
	v8903 = m.ExcPending
	if v8903 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	v8904 = v8902
	goto L1314
L1326:
	;
	goto L1313
L1327:
	;
	v8953 = v8635 + int32(1)
	v8954 = *(*int32)(unsafe.Add(mBase, uint32(v8614)+4))
	if v8953 < v8954 {
		v8620 = v8909
		v8635 = v8953
		goto L1289
	} else {
		goto L1328
	}
L1328:
	;
	goto L1290
L1329:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_36), int32(0))
	mBase = m.M
	v8963 = m.ExcPending
	if v8963 != 0 {
		goto L1
	} else {
		goto L1330
	}
L1330:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_37), int32(89), int32(_a_F_subquery_planner_38))
	mBase = m.M
	v8968 = m.ExcPending
	if v8968 != 0 {
		goto L1
	} else {
		goto L1331
	}
L1331:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1332:
	;
	if v9012 == int32(0) {
		v9090 = v8969
		goto L1257
	} else {
		goto L1333
	}
L1333:
	;
	v9016 = int32(0)
	v9017 = *(*int32)(unsafe.Add(mBase, uint32(v9012)+4))
	if v9017 <= v9016 {
		v9090 = v8969
		goto L1257
	} else {
		goto L1334
	}
L1334:
	;
	v9020 = v8969
	v9021 = v9016
	goto L1335
L1335:
	;
	v9061 = *(*int32)(unsafe.Add(mBase, uint32(v9012)+12))
	v9065 = *(*int32)(unsafe.Add(mBase, uint32(v9061+v9021<<(uint(int32(2))%32))))
	v9066 = *(*int32)(unsafe.Add(mBase, uint32(v9065)))
	if v9066 == int32(6) {
		goto L1338
	} else {
		goto L1339
	}
L1336:
	;
	v9090 = v9085
	goto L1257
L1337:
	;
	v9087 = v9021 + int32(1)
	v9088 = *(*int32)(unsafe.Add(mBase, uint32(v9012)+4))
	if v9087 < v9088 {
		v9020 = v9085
		v9021 = v9087
		goto L1335
	} else {
		goto L1349
	}
L1338:
	;
	v9069 = *(*int32)(unsafe.Add(mBase, uint32(v9065)+4))
	if v9069 == v8423 {
		v9085 = v9020
		goto L1337
	} else {
		goto L1341
	}
L1339:
	;
	goto L1340
L1340:
	;
	v9071 = F_tlist_member(m, v9065, v9020)
	mBase = m.M
	v9072 = m.ExcPending
	if v9072 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1341:
	;
	goto L1340
L1342:
	;
	if v9071 != 0 {
		v9085 = v9020
		goto L1337
	} else {
		goto L1343
	}
L1343:
	;
	if v9020 != 0 {
		goto L1344
	} else {
		goto L1345
	}
L1344:
	;
	v9073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9020)+4)))
	v9077 = v9073 + int32(1)
	goto L1346
L1345:
	;
	v9077 = int32(1)
	goto L1346
L1346:
	;
	v9081 = F_makeTargetEntry(m, v9065, base.I32_extend16_s(v9077), int32(0), int32(1))
	mBase = m.M
	v9082 = m.ExcPending
	if v9082 != 0 {
		goto L1
	} else {
		goto L1347
	}
L1347:
	;
	v9083 = F_lappend(m, v9020, v9081)
	mBase = m.M
	v9084 = m.ExcPending
	if v9084 != 0 {
		goto L1
	} else {
		goto L1348
	}
L1348:
	;
	v9085 = v9083
	goto L1337
L1349:
	;
	goto L1336
L1350:
	;
	v9349 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+96))
	if v9349 == int32(0) {
		v9480 = v9308
		goto L1389
	} else {
		goto L1390
	}
L1351:
	;
	v9134 = *(*int32)(unsafe.Add(mBase, uint32(v9131)+4))
	if v9134 <= int32(0) {
		v9308 = v9090
		goto L1350
	} else {
		goto L1352
	}
L1352:
	;
	v9138 = v9090
	v9140 = int32(0)
	goto L1353
L1353:
	;
	v9179 = *(*int32)(unsafe.Add(mBase, uint32(v9131)+12))
	v9183 = *(*int32)(unsafe.Add(mBase, uint32(v9179+v9140<<(uint(int32(2))%32))))
	v9184 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+4))
	v9185 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+8))
	if v9184 != v9185 {
		v9301 = v9138
		goto L1355
	} else {
		goto L1356
	}
L1354:
	;
	v9308 = v9301
	goto L1350
L1355:
	;
	v9305 = v9140 + int32(1)
	v9306 = *(*int32)(unsafe.Add(mBase, uint32(v9131)+4))
	if v9305 < v9306 {
		v9138 = v9301
		v9140 = v9305
		goto L1353
	} else {
		goto L1388
	}
L1356:
	;
	v9187 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+20))
	if v9187&int32(-33) != 0 {
		goto L1357
	} else {
		goto L1358
	}
L1357:
	;
	v9190 = int32(-1)
	v9193 = int32(0)
	v9195 = F_makeVar(m, v9184, v9190, int32(27), v9190, v9193, v9193)
	mBase = m.M
	v9196 = m.ExcPending
	if v9196 != 0 {
		goto L1
	} else {
		goto L1360
	}
L1358:
	;
	v9223 = v9138
	v9225 = v9187
	goto L1359
L1359:
	;
	if v9225&int32(32) != 0 {
		goto L1368
	} else {
		goto L1369
	}
L1360:
	;
	v9197 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8418)+32)) = v9197
	v9201 = int32(32)
	v9205 = F_pg_snprintf(m, v8418+int32(48), v9201, int32(_a_F_subquery_planner_39), v8418+v9201)
	mBase = m.M
	v9206 = m.ExcPending
	if v9206 != 0 {
		goto L1
	} else {
		goto L1361
	}
L1361:
	;
	if v9138 != 0 {
		goto L1362
	} else {
		goto L1363
	}
L1362:
	;
	v9207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9138)+4)))
	v9211 = v9207 + int32(1)
	goto L1364
L1363:
	;
	v9211 = int32(1)
	goto L1364
L1364:
	;
	v9215 = F_pstrdup(m, v8418+int32(48))
	mBase = m.M
	v9216 = m.ExcPending
	if v9216 != 0 {
		goto L1
	} else {
		goto L1365
	}
L1365:
	;
	v9218 = F_makeTargetEntry(m, v9195, base.I32_extend16_s(v9211), v9215, int32(1))
	mBase = m.M
	v9219 = m.ExcPending
	if v9219 != 0 {
		goto L1
	} else {
		goto L1366
	}
L1366:
	;
	v9220 = F_lappend(m, v9138, v9218)
	mBase = m.M
	v9221 = m.ExcPending
	if v9221 != 0 {
		goto L1
	} else {
		goto L1367
	}
L1367:
	;
	v9222 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+20))
	v9223 = v9220
	v9225 = v9222
	goto L1359
L1368:
	;
	v9228 = *(*int32)(unsafe.Add(mBase, uint32(v8422)+12))
	v9229 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+4))
	v9235 = *(*int32)(unsafe.Add(mBase, uint32(v9228+v9229<<(uint(int32(2))%32)-int32(4))))
	v9236 = int32(0)
	v9238 = F_makeWholeRowVar(m, v9235, v9229, v9236, v9236)
	mBase = m.M
	v9239 = m.ExcPending
	if v9239 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1369:
	;
	v9265 = v9223
	goto L1370
L1370:
	;
	v9267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9183)+32)))
	if v9267 != int32(1) {
		v9301 = v9265
		goto L1355
	} else {
		goto L1379
	}
L1371:
	;
	v9240 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8418)+16)) = v9240
	v9248 = F_pg_snprintf(m, v8418+int32(48), int32(32), int32(_a_F_subquery_planner_40), v8418+int32(16))
	mBase = m.M
	v9249 = m.ExcPending
	if v9249 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1372:
	;
	if v9223 != 0 {
		goto L1373
	} else {
		goto L1374
	}
L1373:
	;
	v9250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9223)+4)))
	v9254 = v9250 + int32(1)
	goto L1375
L1374:
	;
	v9254 = int32(1)
	goto L1375
L1375:
	;
	v9258 = F_pstrdup(m, v8418+int32(48))
	mBase = m.M
	v9259 = m.ExcPending
	if v9259 != 0 {
		goto L1
	} else {
		goto L1376
	}
L1376:
	;
	v9261 = F_makeTargetEntry(m, v9238, base.I32_extend16_s(v9254), v9258, int32(1))
	mBase = m.M
	v9262 = m.ExcPending
	if v9262 != 0 {
		goto L1
	} else {
		goto L1377
	}
L1377:
	;
	v9263 = F_lappend(m, v9223, v9261)
	mBase = m.M
	v9264 = m.ExcPending
	if v9264 != 0 {
		goto L1
	} else {
		goto L1378
	}
L1378:
	;
	v9265 = v9263
	goto L1370
L1379:
	;
	v9270 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+4))
	v9274 = int32(0)
	v9276 = F_makeVar(m, v9270, int32(-6), int32(26), int32(-1), v9274, v9274)
	mBase = m.M
	v9277 = m.ExcPending
	if v9277 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1380:
	;
	v9278 = *(*int32)(unsafe.Add(mBase, uint32(v9183)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8418))) = v9278
	v9284 = F_pg_snprintf(m, v8418+int32(48), int32(32), int32(_a_F_subquery_planner_41), v8418)
	mBase = m.M
	v9285 = m.ExcPending
	if v9285 != 0 {
		goto L1
	} else {
		goto L1381
	}
L1381:
	;
	if v9265 != 0 {
		goto L1382
	} else {
		goto L1383
	}
L1382:
	;
	v9286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9265)+4)))
	v9290 = v9286 + int32(1)
	goto L1384
L1383:
	;
	v9290 = int32(1)
	goto L1384
L1384:
	;
	v9294 = F_pstrdup(m, v8418+int32(48))
	mBase = m.M
	v9295 = m.ExcPending
	if v9295 != 0 {
		goto L1
	} else {
		goto L1385
	}
L1385:
	;
	v9297 = F_makeTargetEntry(m, v9276, base.I32_extend16_s(v9290), v9294, int32(1))
	mBase = m.M
	v9298 = m.ExcPending
	if v9298 != 0 {
		goto L1
	} else {
		goto L1386
	}
L1386:
	;
	v9299 = F_lappend(m, v9265, v9297)
	mBase = m.M
	v9300 = m.ExcPending
	if v9300 != 0 {
		goto L1
	} else {
		goto L1387
	}
L1387:
	;
	v9301 = v9299
	goto L1355
L1388:
	;
	goto L1354
L1389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+264)) = v9480
	if v8438 != 0 {
		goto L1414
	} else {
		goto L1415
	}
L1390:
	;
	v9352 = *(*int32)(unsafe.Add(mBase, uint32(v8420)+52))
	if v9352 == int32(0) {
		v9480 = v9308
		goto L1389
	} else {
		goto L1391
	}
L1391:
	;
	v9355 = *(*int32)(unsafe.Add(mBase, uint32(v9352)+4))
	if v9355 < int32(2) {
		v9480 = v9308
		goto L1389
	} else {
		goto L1392
	}
L1392:
	;
	v9359 = F_pull_var_clause(m, v9349, int32(26))
	mBase = m.M
	v9360 = m.ExcPending
	if v9360 != 0 {
		goto L1
	} else {
		goto L1394
	}
L1393:
	;
	F_list_free(m, v9359)
	mBase = m.M
	v9479 = m.ExcPending
	if v9479 != 0 {
		goto L1
	} else {
		goto L1413
	}
L1394:
	;
	if v9359 == int32(0) {
		v9437 = v9308
		goto L1393
	} else {
		goto L1395
	}
L1395:
	;
	v9363 = *(*int32)(unsafe.Add(mBase, uint32(v9359)+4))
	if v9363 <= int32(0) {
		v9437 = v9308
		goto L1393
	} else {
		goto L1396
	}
L1396:
	;
	v9367 = v9308
	v9368 = int32(0)
	goto L1397
L1397:
	;
	v9408 = *(*int32)(unsafe.Add(mBase, uint32(v9359)+12))
	v9412 = *(*int32)(unsafe.Add(mBase, uint32(v9408+v9368<<(uint(int32(2))%32))))
	v9413 = *(*int32)(unsafe.Add(mBase, uint32(v9412)))
	if v9413 == int32(6) {
		goto L1400
	} else {
		goto L1401
	}
L1398:
	;
	v9437 = v9432
	goto L1393
L1399:
	;
	v9434 = v9368 + int32(1)
	v9435 = *(*int32)(unsafe.Add(mBase, uint32(v9359)+4))
	if v9434 < v9435 {
		v9367 = v9432
		v9368 = v9434
		goto L1397
	} else {
		goto L1412
	}
L1400:
	;
	v9416 = *(*int32)(unsafe.Add(mBase, uint32(v9412)+4))
	if v9416 == v8423 {
		v9432 = v9367
		goto L1399
	} else {
		goto L1403
	}
L1401:
	;
	goto L1402
L1402:
	;
	v9418 = F_tlist_member(m, v9412, v9367)
	mBase = m.M
	v9419 = m.ExcPending
	if v9419 != 0 {
		goto L1
	} else {
		goto L1404
	}
L1403:
	;
	goto L1402
L1404:
	;
	if v9418 != 0 {
		v9432 = v9367
		goto L1399
	} else {
		goto L1405
	}
L1405:
	;
	if v9367 != 0 {
		goto L1407
	} else {
		goto L1408
	}
L1406:
	;
	v9428 = F_makeTargetEntry(m, v9412, base.I32_extend16_s(v9424), int32(0), int32(1))
	mBase = m.M
	v9429 = m.ExcPending
	if v9429 != 0 {
		goto L1
	} else {
		goto L1410
	}
L1407:
	;
	v9420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9367)+4)))
	v9424 = v9420 + int32(1)
	goto L1406
L1408:
	;
	goto L1409
L1409:
	;
	v9424 = int32(1)
	goto L1406
L1410:
	;
	v9430 = F_lappend(m, v9367, v9428)
	mBase = m.M
	v9431 = m.ExcPending
	if v9431 != 0 {
		goto L1
	} else {
		goto L1411
	}
L1411:
	;
	v9432 = v9430
	goto L1399
L1412:
	;
	goto L1398
L1413:
	;
	v9480 = v9437
	goto L1389
L1414:
	;
	F_relation_close(m, v8438, int32(0))
	mBase = m.M
	v9524 = m.ExcPending
	if v9524 != 0 {
		goto L1
	} else {
		goto L1417
	}
L1415:
	;
	goto L1416
L1416:
	;
	m.G0 = v8418 + int32(80)
	v9528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+36)))
	if v9528 == int32(1) {
		goto L1418
	} else {
		goto L1419
	}
L1417:
	;
	goto L1416
L1418:
	;
	v9531 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	F_preprocess_aggrefs(m, v8377, v9531)
	mBase = m.M
	v9533 = m.ExcPending
	if v9533 != 0 {
		goto L1
	} else {
		goto L1421
	}
L1419:
	;
	goto L1420
L1420:
	;
	v9537 = int32(0)
	v9539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+37)))
	if v9539 != int32(1) {
		v10805 = v9537
		v10807 = v9537
		goto L1423
	} else {
		goto L1424
	}
L1421:
	;
	v9534 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+112))
	F_preprocess_aggrefs(m, v8377, v9534)
	mBase = m.M
	v9536 = m.ExcPending
	if v9536 != 0 {
		goto L1
	} else {
		goto L1422
	}
L1422:
	;
	goto L1420
L1423:
	;
	v10826 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+36)))
	if v10826 == int32(1) {
		goto L1564
	} else {
		goto L1565
	}
L1424:
	;
	v9542 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	v9543 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+116))
	if v9543 != 0 {
		goto L1425
	} else {
		goto L1426
	}
L1425:
	;
	v9544 = *(*int32)(unsafe.Add(mBase, uint32(v9543)+4))
	v9546 = v9544
	goto L1427
L1426:
	;
	v9546 = int32(0)
	goto L1427
L1427:
	;
	v9548 = F_palloc(m, int32(12))
	mBase = m.M
	v9549 = m.ExcPending
	if v9549 != 0 {
		goto L1
	} else {
		goto L1428
	}
L1428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9548)+4)) = v9546
	*(*int32)(unsafe.Add(mBase, uint32(v9548))) = int32(0)
	v9557 = F_palloc0(m, v9546<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v9558 = m.ExcPending
	if v9558 != 0 {
		goto L1
	} else {
		goto L1429
	}
L1429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9548)+8)) = v9557
	v9560 = F_find_window_functions_walker(m, v9542, v9548)
	mBase = m.M
	v9561 = m.ExcPending
	if v9561 != 0 {
		goto L1
	} else {
		goto L1430
	}
L1430:
	;
	v9562 = *(*int32)(unsafe.Add(mBase, uint32(v9548)))
	if int32(0) < v9562 {
		goto L1433
	} else {
		goto L1434
	}
L1431:
	;
	v10289 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v10290 = *(*int32)(unsafe.Add(mBase, uint32(v10289)+116))
	if v10290 == int32(0) {
		goto L1508
	} else {
		goto L1509
	}
L1432:
	;
	v10031 = int32(0)
	if v10027 <= v10031 {
		goto L1431
	} else {
		goto L1485
	}
L1433:
	;
	v9565 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v9566 = *(*int32)(unsafe.Add(mBase, uint32(v9565)+116))
	if v9566 == int32(0) {
		goto L1431
	} else {
		goto L1436
	}
L1434:
	;
	goto L1435
L1435:
	;
	v10029 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8387)+37)) = uint8(v10029)
	v10805 = v9537
	v10807 = v9548
	goto L1423
L1436:
	;
	v9569 = int32(0)
	v9570 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+4))
	if v9570 <= v9569 {
		goto L1431
	} else {
		goto L1437
	}
L1437:
	;
	v9573 = v9569
	goto L1438
L1438:
	;
	v9614 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v9615 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+12))
	v9616 = int32(2)
	v9619 = *(*int32)(unsafe.Add(mBase, uint32(v9615+v9573<<(uint(v9616)%32))))
	v9620 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+48))
	v9624 = *(*int32)(unsafe.Add(mBase, uint32(v9614+v9620<<(uint(v9616)%32))))
	if v9624 == int32(0) {
		goto L1440
	} else {
		goto L1441
	}
L1439:
	;
	goto L1432
L1440:
	;
	v10026 = v9573 + int32(1)
	v10027 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+4))
	if v10026 < v10027 {
		v9573 = v10026
		goto L1438
	} else {
		goto L1484
	}
L1441:
	;
	v9627 = *(*int32)(unsafe.Add(mBase, uint32(v9624)+4))
	if v9627 <= int32(0) {
		goto L1443
	} else {
		goto L1444
	}
L1442:
	;
	v9767 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+20))
	if v9767 == v9735 {
		goto L1440
	} else {
		goto L1459
	}
L1443:
	;
	v9735 = int32(0)
	goto L1442
L1444:
	;
	goto L1445
L1445:
	;
	v9631 = *(*int32)(unsafe.Add(mBase, uint32(v9624)+12))
	v9632 = *(*int32)(unsafe.Add(mBase, uint32(v9631)))
	v9633 = *(*int32)(unsafe.Add(mBase, uint32(v9632)+4))
	v9634 = F_get_func_support(m, v9633)
	mBase = m.M
	v9635 = m.ExcPending
	if v9635 != 0 {
		goto L1
	} else {
		goto L1446
	}
L1446:
	;
	if v9634 == int32(0) {
		goto L1440
	} else {
		goto L1447
	}
L1447:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+192)) = int32(463)
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+196)) = v9632
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+200)) = v9619
	v9642 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+204)) = v9642
	v9647 = F_OidFunctionCall1Coll(m, v9634, int32(0), v8382+int32(192))
	mBase = m.M
	v9648 = m.ExcPending
	if v9648 != 0 {
		goto L1
	} else {
		goto L1448
	}
L1448:
	;
	if v9647 == int32(0) {
		goto L1440
	} else {
		goto L1449
	}
L1449:
	;
	v9651 = *(*int32)(unsafe.Add(mBase, uint32(v9647)+12))
	v9653 = *(*int32)(unsafe.Add(mBase, uint32(v9624)+4))
	if v9653 < int32(2) {
		v9735 = v9651
		goto L1442
	} else {
		goto L1450
	}
L1450:
	;
	v9673 = int32(1)
	goto L1451
L1451:
	;
	v9697 = *(*int32)(unsafe.Add(mBase, uint32(v9624)+12))
	v9701 = *(*int32)(unsafe.Add(mBase, uint32(v9697+v9673<<(uint(int32(2))%32))))
	v9702 = *(*int32)(unsafe.Add(mBase, uint32(v9701)+4))
	v9703 = F_get_func_support(m, v9702)
	mBase = m.M
	v9704 = m.ExcPending
	if v9704 != 0 {
		goto L1
	} else {
		goto L1453
	}
L1452:
	;
	v9735 = v9651
	goto L1442
L1453:
	;
	if v9703 == int32(0) {
		goto L1440
	} else {
		goto L1454
	}
L1454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+192)) = int32(463)
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+196)) = v9701
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+200)) = v9619
	v9711 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+204)) = v9711
	v9716 = F_OidFunctionCall1Coll(m, v9703, int32(0), v8382+int32(192))
	mBase = m.M
	v9717 = m.ExcPending
	if v9717 != 0 {
		goto L1
	} else {
		goto L1455
	}
L1455:
	;
	if v9716 == int32(0) {
		goto L1440
	} else {
		goto L1456
	}
L1456:
	;
	v9720 = *(*int32)(unsafe.Add(mBase, uint32(v9716)+12))
	if v9651 != v9720 {
		goto L1440
	} else {
		goto L1457
	}
L1457:
	;
	v9723 = v9673 + int32(1)
	v9724 = *(*int32)(unsafe.Add(mBase, uint32(v9624)+4))
	if v9723 < v9724 {
		v9673 = v9723
		goto L1451
	} else {
		goto L1458
	}
L1458:
	;
	goto L1452
L1459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9619)+20)) = v9735
	v9770 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+4))
	if v9770 < int32(2) {
		goto L1440
	} else {
		goto L1460
	}
L1460:
	;
	v9782 = int32(0)
	goto L1461
L1461:
	;
	v9815 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+12))
	v9819 = *(*int32)(unsafe.Add(mBase, uint32(v9815+v9782<<(uint(int32(2))%32))))
	if v9819 == v9619 {
		goto L1463
	} else {
		goto L1464
	}
L1462:
	;
	goto L1440
L1463:
	;
	v9981 = v9782 + int32(1)
	v9982 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+4))
	if v9981 < v9982 {
		v9782 = v9981
		goto L1461
	} else {
		goto L1483
	}
L1464:
	;
	v9821 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+12))
	v9822 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+12))
	v9823 = F_equal(m, v9821, v9822)
	mBase = m.M
	v9824 = m.ExcPending
	if v9824 != 0 {
		goto L1
	} else {
		goto L1465
	}
L1465:
	;
	if v9823 == int32(0) {
		goto L1463
	} else {
		goto L1466
	}
L1466:
	;
	v9827 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+16))
	v9828 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+16))
	v9829 = F_equal(m, v9827, v9828)
	mBase = m.M
	v9830 = m.ExcPending
	if v9830 != 0 {
		goto L1
	} else {
		goto L1467
	}
L1467:
	;
	if v9829 == int32(0) {
		goto L1463
	} else {
		goto L1468
	}
L1468:
	;
	v9833 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+20))
	v9834 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+20))
	if v9833 != v9834 {
		goto L1463
	} else {
		goto L1469
	}
L1469:
	;
	v9836 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+24))
	v9837 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+24))
	v9838 = F_equal(m, v9836, v9837)
	mBase = m.M
	v9839 = m.ExcPending
	if v9839 != 0 {
		goto L1
	} else {
		goto L1470
	}
L1470:
	;
	if v9838 == int32(0) {
		goto L1463
	} else {
		goto L1471
	}
L1471:
	;
	v9842 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+28))
	v9843 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+28))
	v9844 = F_equal(m, v9842, v9843)
	mBase = m.M
	v9845 = m.ExcPending
	if v9845 != 0 {
		goto L1
	} else {
		goto L1472
	}
L1472:
	;
	if v9844 == int32(0) {
		goto L1463
	} else {
		goto L1473
	}
L1473:
	;
	v9848 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v9849 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+48))
	v9853 = *(*int32)(unsafe.Add(mBase, uint32(v9848+v9849<<(uint(int32(2))%32))))
	if v9853 == int32(0) {
		goto L1475
	} else {
		goto L1476
	}
L1474:
	;
	v9960 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+48))
	v9964 = *(*int32)(unsafe.Add(mBase, uint32(v9927+v9960<<(uint(int32(2))%32))))
	v9965 = F_list_concat(m, v9964, v9928)
	mBase = m.M
	v9966 = m.ExcPending
	if v9966 != 0 {
		goto L1
	} else {
		goto L1482
	}
L1475:
	;
	v9927 = v9848
	v9928 = int32(0)
	goto L1474
L1476:
	;
	goto L1477
L1477:
	;
	v9857 = *(*int32)(unsafe.Add(mBase, uint32(v9853)+4))
	if v9857 <= int32(0) {
		v9927 = v9848
		v9928 = v9853
		goto L1474
	} else {
		goto L1478
	}
L1478:
	;
	v9860 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+48))
	v9870 = int32(0)
	goto L1479
L1479:
	;
	v9903 = *(*int32)(unsafe.Add(mBase, uint32(v9853)+12))
	v9907 = *(*int32)(unsafe.Add(mBase, uint32(v9903+v9870<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9907)+32)) = v9860
	v9910 = v9870 + int32(1)
	v9911 = *(*int32)(unsafe.Add(mBase, uint32(v9853)+4))
	if v9910 < v9911 {
		v9870 = v9910
		goto L1479
	} else {
		goto L1481
	}
L1480:
	;
	v9913 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v9914 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+48))
	v9918 = *(*int32)(unsafe.Add(mBase, uint32(v9913+v9914<<(uint(int32(2))%32))))
	v9927 = v9913
	v9928 = v9918
	goto L1474
L1481:
	;
	goto L1480
L1482:
	;
	v9967 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v9968 = *(*int32)(unsafe.Add(mBase, uint32(v9819)+48))
	v9969 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9967+v9968<<(uint(v9969)%32)))) = v9965
	v9973 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v9974 = *(*int32)(unsafe.Add(mBase, uint32(v9619)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9973+v9974<<(uint(v9969)%32)))) = int32(0)
	goto L1440
L1483:
	;
	goto L1462
L1484:
	;
	goto L1439
L1485:
	;
	v10041 = v10027
	v10043 = v10031
	goto L1486
L1486:
	;
	v10075 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v10076 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+12))
	v10077 = int32(2)
	v10080 = *(*int32)(unsafe.Add(mBase, uint32(v10076+v10043<<(uint(v10077)%32))))
	v10081 = *(*int32)(unsafe.Add(mBase, uint32(v10080)+48))
	v10085 = *(*int32)(unsafe.Add(mBase, uint32(v10075+v10081<<(uint(v10077)%32))))
	if v10085 != 0 {
		goto L1488
	} else {
		goto L1489
	}
L1487:
	;
	goto L1431
L1488:
	;
	v10086 = *(*int32)(unsafe.Add(mBase, uint32(v10085)+4))
	if v10086 <= int32(0) {
		goto L1492
	} else {
		goto L1493
	}
L1489:
	;
	v10211 = v10041
	goto L1490
L1490:
	;
	v10246 = v10043 + int32(1)
	if v10246 < v10211 {
		v10041 = v10211
		v10043 = v10246
		goto L1486
	} else {
		goto L1505
	}
L1491:
	;
	F_list_free(m, v10085)
	mBase = m.M
	v10196 = m.ExcPending
	if v10196 != 0 {
		goto L1
	} else {
		goto L1504
	}
L1492:
	;
	v10171 = int32(0)
	goto L1491
L1493:
	;
	goto L1494
L1494:
	;
	v10090 = int32(0)
	v10099 = v10090
	v10109 = v10090
	goto L1495
L1495:
	;
	v10133 = *(*int32)(unsafe.Add(mBase, uint32(v10085)+12))
	v10136 = v10133 + v10099<<(uint(int32(2))%32)
	v10137 = *(*int32)(unsafe.Add(mBase, uint32(v10136)))
	v10138 = F_list_member(m, v10109, v10137)
	mBase = m.M
	v10139 = m.ExcPending
	if v10139 != 0 {
		goto L1
	} else {
		goto L1498
	}
L1496:
	;
	v10171 = v10149
	goto L1491
L1497:
	;
	v10151 = v10099 + int32(1)
	v10152 = *(*int32)(unsafe.Add(mBase, uint32(v10085)+4))
	if v10151 < v10152 {
		v10099 = v10151
		v10109 = v10149
		goto L1495
	} else {
		goto L1503
	}
L1498:
	;
	if v10138 == int32(0) {
		goto L1499
	} else {
		goto L1500
	}
L1499:
	;
	v10142 = *(*int32)(unsafe.Add(mBase, uint32(v10136)))
	v10143 = F_lappend(m, v10109, v10142)
	mBase = m.M
	v10144 = m.ExcPending
	if v10144 != 0 {
		goto L1
	} else {
		goto L1502
	}
L1500:
	;
	goto L1501
L1501:
	;
	v10145 = *(*int32)(unsafe.Add(mBase, uint32(v9548)))
	*(*int32)(unsafe.Add(mBase, uint32(v9548))) = v10145 - int32(1)
	v10149 = v10109
	goto L1497
L1502:
	;
	v10149 = v10143
	goto L1497
L1503:
	;
	goto L1496
L1504:
	;
	v10197 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v10198 = *(*int32)(unsafe.Add(mBase, uint32(v10080)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10197+v10198<<(uint(int32(2))%32)))) = v10171
	v10203 = *(*int32)(unsafe.Add(mBase, uint32(v9566)+4))
	v10211 = v10203
	goto L1490
L1505:
	;
	goto L1487
L1506:
	;
	F_pfree(m, v10751)
	mBase = m.M
	v10784 = m.ExcPending
	if v10784 != 0 {
		goto L1
	} else {
		goto L1563
	}
L1507:
	;
	F_pg_qsort(m, v10735, int32(0), int32(8), int32(832))
	mBase = m.M
	v10741 = m.ExcPending
	if v10741 != 0 {
		goto L1
	} else {
		goto L1562
	}
L1508:
	;
	v10294 = F_palloc(m, int32(0))
	mBase = m.M
	v10295 = m.ExcPending
	if v10295 != 0 {
		goto L1
	} else {
		goto L1511
	}
L1509:
	;
	goto L1510
L1510:
	;
	v10296 = *(*int32)(unsafe.Add(mBase, uint32(v10290)+4))
	v10299 = F_palloc(m, v10296<<(uint(int32(3))%32))
	mBase = m.M
	v10300 = m.ExcPending
	if v10300 != 0 {
		goto L1
	} else {
		goto L1512
	}
L1511:
	;
	v10735 = v10294
	goto L1507
L1512:
	;
	v10301 = *(*int32)(unsafe.Add(mBase, uint32(v10290)+4))
	if v10301 <= int32(0) {
		v10735 = v10299
		goto L1507
	} else {
		goto L1513
	}
L1513:
	;
	v10304 = int32(0)
	v10313 = v10304
	v10320 = v10301
	v10323 = v10304
	goto L1514
L1514:
	;
	v10347 = *(*int32)(unsafe.Add(mBase, uint32(v9548)+8))
	v10348 = *(*int32)(unsafe.Add(mBase, uint32(v10290)+12))
	v10349 = int32(2)
	v10352 = *(*int32)(unsafe.Add(mBase, uint32(v10348+v10313<<(uint(v10349)%32))))
	v10353 = *(*int32)(unsafe.Add(mBase, uint32(v10352)+48))
	v10357 = *(*int32)(unsafe.Add(mBase, uint32(v10347+v10353<<(uint(v10349)%32))))
	if v10357 != 0 {
		goto L1516
	} else {
		goto L1517
	}
L1515:
	;
	F_pg_qsort(m, v10299, v10374, int32(8), int32(832))
	mBase = m.M
	v10381 = m.ExcPending
	if v10381 != 0 {
		goto L1
	} else {
		goto L1522
	}
L1516:
	;
	v10360 = v10299 + v10323<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v10360))) = v10352
	v10362 = *(*int32)(unsafe.Add(mBase, uint32(v10352)+12))
	v10363 = F_list_copy(m, v10362)
	mBase = m.M
	v10364 = m.ExcPending
	if v10364 != 0 {
		goto L1
	} else {
		goto L1519
	}
L1517:
	;
	v10373 = v10320
	v10374 = v10323
	goto L1518
L1518:
	;
	v10376 = v10313 + int32(1)
	if v10376 < v10373 {
		v10313 = v10376
		v10320 = v10373
		v10323 = v10374
		goto L1514
	} else {
		goto L1521
	}
L1519:
	;
	v10365 = *(*int32)(unsafe.Add(mBase, uint32(v10352)+16))
	v10366 = F_list_concat_unique(m, v10363, v10365)
	mBase = m.M
	v10367 = m.ExcPending
	if v10367 != 0 {
		goto L1
	} else {
		goto L1520
	}
L1520:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10360)+4)) = v10366
	v10371 = *(*int32)(unsafe.Add(mBase, uint32(v10290)+4))
	v10373 = v10371
	v10374 = v10323 + int32(1)
	goto L1518
L1521:
	;
	goto L1515
L1522:
	;
	v10382 = int32(0)
	if v10374 <= v10382 {
		v10751 = v10299
		goto L1506
	} else {
		goto L1523
	}
L1523:
	;
	v10392 = v10382
	v10405 = v9537
	goto L1524
L1524:
	;
	v10429 = *(*int32)(unsafe.Add(mBase, uint32(v10299+v10392<<(uint(int32(3))%32))))
	v10430 = F_lappend(m, v10405, v10429)
	mBase = m.M
	v10431 = m.ExcPending
	if v10431 != 0 {
		goto L1
	} else {
		goto L1526
	}
L1525:
	;
	F_pfree(m, v10299)
	mBase = m.M
	v10436 = m.ExcPending
	if v10436 != 0 {
		goto L1
	} else {
		goto L1528
	}
L1526:
	;
	v10433 = v10392 + int32(1)
	if v10433 != v10374 {
		v10392 = v10433
		v10405 = v10430
		goto L1524
	} else {
		goto L1527
	}
L1527:
	;
	goto L1525
L1528:
	;
	if v10430 == int32(0) {
		goto L1529
	} else {
		goto L1530
	}
L1529:
	;
	v10805 = int32(0)
	v10807 = v9548
	goto L1423
L1530:
	;
	goto L1531
L1531:
	;
	v10441 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+4))
	if v10441 <= int32(0) {
		v10805 = v10430
		v10807 = v9548
		goto L1423
	} else {
		goto L1532
	}
L1532:
	;
	v10452 = v10441
	v10454 = int32(1)
	v10459 = int32(0)
	goto L1533
L1533:
	;
	v10486 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+12))
	v10490 = *(*int32)(unsafe.Add(mBase, uint32(v10486+v10459<<(uint(int32(2))%32))))
	v10491 = *(*int32)(unsafe.Add(mBase, uint32(v10490)+4))
	if v10491 == int32(0) {
		goto L1535
	} else {
		goto L1536
	}
L1534:
	;
	v10805 = v10430
	v10807 = v9548
	goto L1423
L1535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+64)) = v10454
	v10501 = F_pg_snprintf(m, v8382+int32(192), int32(16), int32(_a_F_subquery_planner_42), v8382-int32(-64))
	mBase = m.M
	v10502 = m.ExcPending
	if v10502 != 0 {
		goto L1
	} else {
		goto L1538
	}
L1536:
	;
	v10698 = v10452
	v10700 = v10454
	goto L1537
L1537:
	;
	v10733 = v10459 + int32(1)
	if v10733 < v10698 {
		v10452 = v10698
		v10454 = v10700
		v10459 = v10733
		goto L1533
	} else {
		goto L1561
	}
L1538:
	;
	v10504 = v10454 + int32(1)
	v10505 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+4))
	if v10505 <= int32(0) {
		v10653 = v10504
		goto L1539
	} else {
		goto L1540
	}
L1539:
	;
	v10687 = F_pstrdup(m, v8382+int32(192))
	mBase = m.M
	v10688 = m.ExcPending
	if v10688 != 0 {
		goto L1
	} else {
		goto L1560
	}
L1540:
	;
	v10508 = v10505
	v10517 = v10504
	goto L1541
L1541:
	;
	v10549 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+12))
	v10558 = int32(0)
	goto L1543
L1542:
	;
	v10653 = v10640
	goto L1539
L1543:
	;
	v10595 = *(*int32)(unsafe.Add(mBase, uint32(v10549+v10558<<(uint(int32(2))%32))))
	v10596 = *(*int32)(unsafe.Add(mBase, uint32(v10595)+4))
	if v10596 != 0 {
		goto L1546
	} else {
		goto L1547
	}
L1544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+48)) = v10517
	v10637 = F_pg_snprintf(m, v8382+int32(192), int32(16), int32(_a_F_subquery_planner_42), v8382+int32(48))
	mBase = m.M
	v10638 = m.ExcPending
	if v10638 != 0 {
		goto L1
	} else {
		goto L1558
	}
L1545:
	;
	goto L1544
L1546:
	;
	v10598 = v8382 + int32(192)
	v10601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10596))))
	v10604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10598))))
	if base.B2i32(v10601 == int32(0))|base.B2i32(v10601 != v10604) != 0 {
		v10622 = v10601
		v10623 = v10604
		goto L1550
	} else {
		goto L1551
	}
L1547:
	;
	goto L1548
L1548:
	;
	v10628 = v10558 + int32(1)
	if v10628 != v10508 {
		v10558 = v10628
		goto L1543
	} else {
		goto L1557
	}
L1549:
	;
	if v10622-v10623 == int32(0) {
		goto L1545
	} else {
		goto L1556
	}
L1550:
	;
	goto L1549
L1551:
	;
	v10607 = v10596
	v10608 = v10598
	goto L1552
L1552:
	;
	v10611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10608)+1)))
	v10612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10607)+1)))
	if v10612 == int32(0) {
		v10622 = v10612
		v10623 = v10611
		goto L1550
	} else {
		goto L1554
	}
L1553:
	;
	v10622 = v10612
	v10623 = v10611
	goto L1550
L1554:
	;
	v10615 = int32(1)
	if v10612 == v10611 {
		v10607 = v10607 + v10615
		v10608 = v10608 + v10615
		goto L1552
	} else {
		goto L1555
	}
L1555:
	;
	goto L1553
L1556:
	;
	goto L1548
L1557:
	;
	v10653 = v10517
	goto L1539
L1558:
	;
	v10640 = v10517 + int32(1)
	v10641 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+4))
	if int32(0) < v10641 {
		v10508 = v10641
		v10517 = v10640
		goto L1541
	} else {
		goto L1559
	}
L1559:
	;
	goto L1542
L1560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10490)+4)) = v10687
	v10690 = *(*int32)(unsafe.Add(mBase, uint32(v10430)+4))
	v10698 = v10690
	v10700 = v10653
	goto L1537
L1561:
	;
	goto L1534
L1562:
	;
	v10751 = v10735
	goto L1506
L1563:
	;
	v10805 = v9537
	v10807 = v9548
	goto L1423
L1564:
	;
	v10829 = int32(0)
	v10831 = float64(0)
	v10832 = m.G0
	v10834 = v10832 - int32(16)
	m.G0 = v10834
	v10836 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v10837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10836)+36)))
	if v10837 != int32(1) {
		goto L1569
	} else {
		goto L1570
	}
L1565:
	;
	goto L1566
L1566:
	;
	v11642 = float64(-1)
	v11643 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+100))
	if v11643 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1671
	}
L1567:
	;
	goto L1566
L1568:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11590 = m.ExcPending
	if v11590 != 0 {
		goto L1
	} else {
		goto L1667
	}
L1569:
	;
	m.G0 = v10834 + int32(16)
	goto L1567
L1570:
	;
	v10840 = *(*int32)(unsafe.Add(mBase, uint32(v10836)+100))
	if v10840 != 0 {
		goto L1569
	} else {
		goto L1571
	}
L1571:
	;
	v10841 = *(*int32)(unsafe.Add(mBase, uint32(v10836)+108))
	if v10841 != 0 {
		goto L1572
	} else {
		goto L1573
	}
L1572:
	;
	v10842 = *(*int32)(unsafe.Add(mBase, uint32(v10841)+4))
	if int32(1) < v10842 {
		goto L1569
	} else {
		goto L1575
	}
L1573:
	;
	goto L1574
L1574:
	;
	v10845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10836)+37)))
	if v10845 != 0 {
		goto L1569
	} else {
		goto L1576
	}
L1575:
	;
	goto L1574
L1576:
	;
	v10846 = *(*int32)(unsafe.Add(mBase, uint32(v10836)+48))
	if v10846 != 0 {
		goto L1569
	} else {
		goto L1577
	}
L1577:
	;
	v10849 = v10836 + int32(60)
	goto L1581
L1578:
	;
	v10924 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10834)+12)) = v10924
	v10927 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+328))
	if v10927 == v10924 {
		v11130 = int32(1)
		goto L1592
	} else {
		goto L1593
	}
L1579:
	;
	v10921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10919)+20)))
	if v10921 != int32(1) {
		goto L1569
	} else {
		goto L1591
	}
L1580:
	;
	v10919 = *(*int32)(unsafe.Add(mBase, uint32(v10918)))
	v10920 = *(*int32)(unsafe.Add(mBase, uint32(v10919)+12))
	switch v10920 {
	case 0:
		goto L1578
	case 1:
		goto L1579
	default:
		goto L1569
	}
L1581:
	;
	v10890 = *(*int32)(unsafe.Add(mBase, uint32(v10849)))
	v10891 = *(*int32)(unsafe.Add(mBase, uint32(v10890)))
	if v10891 != int32(65) {
		goto L1584
	} else {
		goto L1585
	}
L1582:
	;
	v10910 = *(*int32)(unsafe.Add(mBase, uint32(v10836)+52))
	v10911 = *(*int32)(unsafe.Add(mBase, uint32(v10910)+12))
	v10912 = *(*int32)(unsafe.Add(mBase, uint32(v10890)+4))
	v10918 = v10911 + v10912<<(uint(int32(2))%32) - int32(4)
	goto L1580
L1583:
	;
	goto L1582
L1584:
	;
	if v10891 != int32(63) {
		goto L1569
	} else {
		goto L1587
	}
L1585:
	;
	goto L1586
L1586:
	;
	v10903 = *(*int32)(unsafe.Add(mBase, uint32(v10890)+4))
	if v10903 == int32(0) {
		goto L1569
	} else {
		goto L1589
	}
L1587:
	;
	v10896 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+36))
	if v10896 == int32(0) {
		goto L1583
	} else {
		goto L1588
	}
L1588:
	;
	v10899 = *(*int32)(unsafe.Add(mBase, uint32(v10890)+4))
	v10918 = v10896 + v10899<<(uint(int32(2))%32)
	goto L1580
L1589:
	;
	v10906 = *(*int32)(unsafe.Add(mBase, uint32(v10903)+4))
	if v10906 != int32(1) {
		goto L1569
	} else {
		goto L1590
	}
L1590:
	;
	v10909 = *(*int32)(unsafe.Add(mBase, uint32(v10903)+12))
	v10849 = v10909
	goto L1581
L1591:
	;
	goto L1578
L1592:
	;
	if v11130 == int32(0) {
		goto L1569
	} else {
		goto L1614
	}
L1593:
	;
	v10931 = v10834 + int32(12)
	v10933 = *(*int32)(unsafe.Add(mBase, uint32(v10927)+4))
	if v10933 <= int32(0) {
		v11057 = int32(1)
		goto L1594
	} else {
		goto L1595
	}
L1594:
	;
	v11130 = v11057
	goto L1592
L1595:
	;
	v10950 = v10829
	goto L1596
L1596:
	;
	v10977 = int32(0)
	v10978 = *(*int32)(unsafe.Add(mBase, uint32(v10927)+12))
	v10982 = *(*int32)(unsafe.Add(mBase, uint32(v10978+v10950<<(uint(int32(2))%32))))
	v10983 = *(*int32)(unsafe.Add(mBase, uint32(v10982)+4))
	v10984 = *(*int32)(unsafe.Add(mBase, uint32(v10983)+12))
	v10985 = *(*int32)(unsafe.Add(mBase, uint32(v10984)))
	v10986 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+32))
	if v10986 == v10977 {
		v11130 = v10977
		goto L1592
	} else {
		goto L1598
	}
L1597:
	;
	v11057 = v11043
	goto L1594
L1598:
	;
	v10990 = *(*int32)(unsafe.Add(mBase, uint32(v10986)+4))
	if v10990 != int32(1) {
		v11130 = int32(0)
		goto L1592
	} else {
		goto L1599
	}
L1599:
	;
	v10994 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+36))
	if v10994 != 0 {
		v11130 = int32(0)
		goto L1592
	} else {
		goto L1600
	}
L1600:
	;
	v10996 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+44))
	if v10996 != 0 {
		v11130 = int32(0)
		goto L1592
	} else {
		goto L1601
	}
L1601:
	;
	v10997 = int32(0)
	v10999 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+4))
	v11000 = F_SearchSysCache1(m, v10997, v10999)
	mBase = m.M
	v11001 = m.ExcPending
	if v11001 != 0 {
		goto L1
	} else {
		goto L1602
	}
L1602:
	;
	if v11000 == int32(0) {
		v11057 = v10997
		goto L1594
	} else {
		goto L1603
	}
L1603:
	;
	v11004 = *(*int32)(unsafe.Add(mBase, uint32(v11000)+16))
	v11005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11004)+22)))
	v11007 = *(*int32)(unsafe.Add(mBase, uint32(v11004+v11005)+44))
	F_ReleaseCatCache(m, v11000)
	mBase = m.M
	v11009 = m.ExcPending
	if v11009 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1604:
	;
	if v11007 == int32(0) {
		v11057 = v10997
		goto L1594
	} else {
		goto L1605
	}
L1605:
	;
	v11012 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+32))
	v11013 = *(*int32)(unsafe.Add(mBase, uint32(v11012)+12))
	v11014 = *(*int32)(unsafe.Add(mBase, uint32(v11013)))
	v11015 = *(*int32)(unsafe.Add(mBase, uint32(v11014)+4))
	v11016 = F_contain_mutable_functions(m, v11015)
	mBase = m.M
	v11017 = m.ExcPending
	if v11017 != 0 {
		goto L1
	} else {
		goto L1606
	}
L1606:
	;
	if v11016 != 0 {
		v11057 = v10997
		goto L1594
	} else {
		goto L1607
	}
L1607:
	;
	v11018 = *(*int32)(unsafe.Add(mBase, uint32(v11014)+4))
	v11019 = F_exprType(m, v11018)
	mBase = m.M
	v11020 = m.ExcPending
	if v11020 != 0 {
		goto L1
	} else {
		goto L1608
	}
L1608:
	;
	v11021 = F_type_is_rowtype(m, v11019)
	mBase = m.M
	v11022 = m.ExcPending
	if v11022 != 0 {
		goto L1
	} else {
		goto L1609
	}
L1609:
	;
	if v11021 != 0 {
		v11057 = v10997
		goto L1594
	} else {
		goto L1610
	}
L1610:
	;
	v11024 = F_palloc0(m, int32(40))
	mBase = m.M
	v11025 = m.ExcPending
	if v11025 != 0 {
		goto L1
	} else {
		goto L1611
	}
L1611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11024))) = int32(325)
	v11028 = *(*int32)(unsafe.Add(mBase, uint32(v10985)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11024)+8)) = v11007
	*(*int32)(unsafe.Add(mBase, uint32(v11024)+4)) = v11028
	v11031 = *(*int32)(unsafe.Add(mBase, uint32(v11014)+4))
	v11032 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11024)+16)) = v11032
	*(*int32)(unsafe.Add(mBase, uint32(v11024)+12)) = v11031
	*(*int64)(unsafe.Add(mBase, uint32(v11024)+24)) = v11032
	*(*int32)(unsafe.Add(mBase, uint32(v11024)+32)) = int32(0)
	v11039 = *(*int32)(unsafe.Add(mBase, uint32(v10931)))
	v11040 = F_lappend(m, v11039, v11024)
	mBase = m.M
	v11041 = m.ExcPending
	if v11041 != 0 {
		goto L1
	} else {
		goto L1612
	}
L1612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10931))) = v11040
	v11043 = int32(1)
	v11045 = v10950 + v11043
	v11046 = *(*int32)(unsafe.Add(mBase, uint32(v10927)+4))
	if v11045 < v11046 {
		v10950 = v11045
		goto L1596
	} else {
		goto L1613
	}
L1613:
	;
	goto L1597
L1614:
	;
	v11133 = *(*int32)(unsafe.Add(mBase, uint32(v10834)+12))
	if v11133 == int32(0) {
		goto L1615
	} else {
		goto L1616
	}
L1615:
	;
	v11359 = F_fetch_upper_rel(m, v8377, int32(2), int32(0))
	mBase = m.M
	v11360 = m.ExcPending
	if v11360 != 0 {
		goto L1
	} else {
		goto L1638
	}
L1616:
	;
	v11136 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if int32(0) < v11136 {
		goto L1617
	} else {
		goto L1618
	}
L1617:
	;
	v11144 = v10829
	goto L1620
L1618:
	;
	goto L1619
L1619:
	;
	v11252 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if v11252 <= int32(0) {
		goto L1615
	} else {
		goto L1631
	}
L1620:
	;
	v11180 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+12))
	v11184 = *(*int32)(unsafe.Add(mBase, uint32(v11180+v11144<<(uint(int32(2))%32))))
	v11185 = *(*int32)(unsafe.Add(mBase, uint32(v11184)+8))
	v11188 = F_get_equality_op_for_ordering_op(m, v11185, v10834+int32(11))
	mBase = m.M
	v11189 = m.ExcPending
	if v11189 != 0 {
		goto L1
	} else {
		goto L1622
	}
L1621:
	;
	goto L1619
L1622:
	;
	if v11188 == int32(0) {
		goto L1568
	} else {
		goto L1623
	}
L1623:
	;
	v11192 = *(*int32)(unsafe.Add(mBase, uint32(v11184)+8))
	v11193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10834)+11)))
	v11194 = F_build_minmax_path(m, v8377, v11184, v11188, v11192, v11193, v11193)
	mBase = m.M
	v11195 = m.ExcPending
	if v11195 != 0 {
		goto L1
	} else {
		goto L1624
	}
L1624:
	;
	if v11194 == int32(0) {
		goto L1625
	} else {
		goto L1626
	}
L1625:
	;
	v11198 = *(*int32)(unsafe.Add(mBase, uint32(v11184)+8))
	v11199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10834)+11)))
	v11202 = F_build_minmax_path(m, v8377, v11184, v11188, v11198, v11199, v11199^int32(1))
	mBase = m.M
	v11203 = m.ExcPending
	if v11203 != 0 {
		goto L1
	} else {
		goto L1628
	}
L1626:
	;
	goto L1627
L1627:
	;
	v11208 = v11144 + int32(1)
	v11209 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if v11208 < v11209 {
		v11144 = v11208
		goto L1620
	} else {
		goto L1630
	}
L1628:
	;
	if v11202 == int32(0) {
		goto L1569
	} else {
		goto L1629
	}
L1629:
	;
	goto L1627
L1630:
	;
	goto L1621
L1631:
	;
	v11261 = int32(0)
	goto L1632
L1632:
	;
	v11297 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+12))
	v11301 = *(*int32)(unsafe.Add(mBase, uint32(v11297+v11261<<(uint(int32(2))%32))))
	v11302 = *(*int32)(unsafe.Add(mBase, uint32(v11301)+12))
	v11303 = F_exprType(m, v11302)
	mBase = m.M
	v11304 = m.ExcPending
	if v11304 != 0 {
		goto L1
	} else {
		goto L1634
	}
L1633:
	;
	goto L1615
L1634:
	;
	v11306 = *(*int32)(unsafe.Add(mBase, uint32(v11301)+12))
	v11307 = F_exprCollation(m, v11306)
	mBase = m.M
	v11308 = m.ExcPending
	if v11308 != 0 {
		goto L1
	} else {
		goto L1635
	}
L1635:
	;
	v11309 = F_generate_new_exec_param(m, v8377, v11303, int32(-1), v11307)
	mBase = m.M
	v11310 = m.ExcPending
	if v11310 != 0 {
		goto L1
	} else {
		goto L1636
	}
L1636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11301)+32)) = v11309
	v11313 = v11261 + int32(1)
	v11314 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if v11313 < v11314 {
		v11261 = v11313
		goto L1632
	} else {
		goto L1637
	}
L1637:
	;
	goto L1633
L1638:
	;
	v11361 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	v11362 = F_make_pathtarget_from_tlist(m, v11361)
	mBase = m.M
	v11363 = m.ExcPending
	if v11363 != 0 {
		goto L1
	} else {
		goto L1639
	}
L1639:
	;
	v11364 = F_set_pathtarget_cost_width(m, v8377, v11362)
	mBase = m.M
	v11365 = m.ExcPending
	if v11365 != 0 {
		goto L1
	} else {
		goto L1640
	}
L1640:
	;
	v11366 = *(*int32)(unsafe.Add(mBase, uint32(v10836)+112))
	v11367 = int32(0)
	v11368 = m.G0
	v11370 = v11368 - int32(16)
	m.G0 = v11370
	v11373 = F_palloc0(m, int32(80))
	mBase = m.M
	v11374 = m.ExcPending
	if v11374 != 0 {
		goto L1
	} else {
		goto L1641
	}
L1641:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+76)) = v11366
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+72)) = v11133
	v11377 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+64)) = v11377
	*(*int64)(unsafe.Add(mBase, uint32(v11373)+32)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+24)) = v11377
	v11383 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v11373)+20)) = uint16(v11383)
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+16)) = v11377
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+12)) = v11364
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+8)) = v11359
	*(*int64)(unsafe.Add(mBase, uint32(v11373))) = int64(1421634175287)
	if v11133 == v11377 {
		goto L1643
	} else {
		goto L1644
	}
L1642:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11373)+40)) = v11469
	v11503 = *(*float64)(unsafe.Add(mBase, uint32(v11364)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v11373)+48)) = base.F64_add(v11465, v11503)
	v11506 = *(*float64)(unsafe.Add(mBase, uint32(v11364)+16))
	v11508 = *(*float64)(unsafe.Add(mBase, uint32(v11364)+24))
	v11511 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v11373)+56)) = base.F64_add(base.F64_add(base.F64_add(v11465, v11506), v11508), v11511)
	if v11366 != 0 {
		goto L1655
	} else {
		goto L1656
	}
L1643:
	;
	v11465 = v10831
	v11469 = v11367
	v11470 = int32(1)
	goto L1642
L1644:
	;
	goto L1645
L1645:
	;
	v11394 = int32(1)
	v11395 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if v11395 <= int32(0) {
		v11465 = v10831
		v11469 = v11367
		v11470 = v11394
		goto L1642
	} else {
		goto L1646
	}
L1646:
	;
	v11399 = int32(0)
	v11403 = v10831
	v11407 = v11367
	v11408 = v11394
	goto L1647
L1647:
	;
	v11440 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+12))
	v11444 = *(*int32)(unsafe.Add(mBase, uint32(v11440+v11399<<(uint(int32(2))%32))))
	v11445 = *(*float64)(unsafe.Add(mBase, uint32(v11444)+24))
	v11446 = *(*int32)(unsafe.Add(mBase, uint32(v11444)+20))
	v11447 = *(*int32)(unsafe.Add(mBase, uint32(v11446)+40))
	v11448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11446)+21)))
	if v11448 == int32(0) {
		goto L1649
	} else {
		goto L1650
	}
L1648:
	;
	v11465 = v11455
	v11469 = v11456
	v11470 = v11454
	goto L1642
L1649:
	;
	v11451 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11373)+21)) = uint8(v11451)
	v11454 = v11451
	goto L1651
L1650:
	;
	v11454 = v11408
	goto L1651
L1651:
	;
	v11455 = base.F64_add(v11403, v11445)
	v11456 = v11407 + v11447
	v11458 = v11399 + int32(1)
	v11459 = *(*int32)(unsafe.Add(mBase, uint32(v11133)+4))
	if v11458 < v11459 {
		v11399 = v11458
		v11403 = v11455
		v11407 = v11456
		v11408 = v11454
		goto L1647
	} else {
		goto L1652
	}
L1652:
	;
	goto L1648
L1653:
	;
	m.G0 = v11370 + int32(16)
	F_add_path(m, v11359, v11373)
	mBase = m.M
	v11542 = m.ExcPending
	if v11542 != 0 {
		goto L1
	} else {
		goto L1666
	}
L1654:
	;
	v11529 = *(*int32)(unsafe.Add(mBase, uint32(v11364)+4))
	v11530 = F_is_parallel_safe(m, v8377, v11529)
	mBase = m.M
	v11531 = m.ExcPending
	if v11531 != 0 {
		goto L1
	} else {
		goto L1661
	}
L1655:
	;
	F_cost_qual_eval(m, v11370, v11366, v8377)
	mBase = m.M
	v11515 = m.ExcPending
	if v11515 != 0 {
		goto L1
	} else {
		goto L1658
	}
L1656:
	;
	goto L1657
L1657:
	;
	if v11470 == int32(0) {
		goto L1653
	} else {
		goto L1660
	}
L1658:
	;
	v11516 = *(*float64)(unsafe.Add(mBase, uint32(v11370)))
	v11517 = *(*float64)(unsafe.Add(mBase, uint32(v11373)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v11373)+48)) = base.F64_add(v11516, v11517)
	v11520 = *(*float64)(unsafe.Add(mBase, uint32(v11373)+56))
	v11521 = *(*float64)(unsafe.Add(mBase, uint32(v11370)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11373)+56)) = base.F64_add(v11520, base.F64_add(v11516, v11521))
	v11525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11373)+21)))
	if v11525 != 0 {
		goto L1654
	} else {
		goto L1659
	}
L1659:
	;
	goto L1653
L1660:
	;
	goto L1654
L1661:
	;
	if v11530 != 0 {
		goto L1662
	} else {
		goto L1663
	}
L1662:
	;
	v11532 = F_is_parallel_safe(m, v8377, v11366)
	mBase = m.M
	v11533 = m.ExcPending
	if v11533 != 0 {
		goto L1
	} else {
		goto L1665
	}
L1663:
	;
	v11535 = int32(0)
	goto L1664
L1664:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11373)+21)) = uint8(v11535)
	goto L1653
L1665:
	;
	v11535 = v11532
	goto L1664
L1666:
	;
	goto L1569
L1667:
	;
	v11591 = *(*int32)(unsafe.Add(mBase, uint32(v11184)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10834))) = v11591
	F_errmsg_internal(m, int32(_a_F_subquery_planner_43), v10834)
	mBase = m.M
	v11595 = m.ExcPending
	if v11595 != 0 {
		goto L1
	} else {
		goto L1668
	}
L1668:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_44), int32(166), int32(_a_F_subquery_planner_45))
	mBase = m.M
	v11600 = m.ExcPending
	if v11600 != 0 {
		goto L1
	} else {
		goto L1669
	}
L1669:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1670:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v8377)+304)) = v11652
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+156)) = v8381
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+152)) = v8390
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+148)) = v10805
	v11660 = F_query_planner(m, v8377, int32(833), v8382+int32(148))
	mBase = m.M
	v11661 = m.ExcPending
	if v11661 != 0 {
		goto L1
	} else {
		goto L1680
	}
L1671:
	;
	v11644 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+108))
	if v11644 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1672
	}
L1672:
	;
	v11645 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+120))
	if v11645 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1673
	}
L1673:
	;
	v11646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+36)))
	if v11646 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1674
	}
L1674:
	;
	v11647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+37)))
	if v11647 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1675
	}
L1675:
	;
	v11648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+38)))
	if v11648 != 0 {
		v11652 = v11642
		goto L1670
	} else {
		goto L1676
	}
L1676:
	;
	v11650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+318)))
	if v11650 != 0 {
		goto L1677
	} else {
		goto L1678
	}
L1677:
	;
	v11651 = float64(-1)
	goto L1679
L1678:
	;
	v11651 = v8405
	goto L1679
L1679:
	;
	v11652 = v11651
	goto L1670
L1680:
	;
	v11662 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	v11663 = F_make_pathtarget_from_tlist(m, v11662)
	mBase = m.M
	v11664 = m.ExcPending
	if v11664 != 0 {
		goto L1
	} else {
		goto L1681
	}
L1681:
	;
	v11665 = F_set_pathtarget_cost_width(m, v8377, v11663)
	mBase = m.M
	v11666 = m.ExcPending
	if v11666 != 0 {
		goto L1
	} else {
		goto L1682
	}
L1682:
	;
	v11667 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	v11668 = F_is_parallel_safe(m, v8377, v11667)
	mBase = m.M
	v11669 = m.ExcPending
	if v11669 != 0 {
		goto L1
	} else {
		goto L1683
	}
L1683:
	;
	v11670 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+124))
	if v11670 != 0 {
		goto L1684
	} else {
		goto L1685
	}
L1684:
	;
	v11671 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v11672 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	if v11672 != 0 {
		goto L1687
	} else {
		goto L1688
	}
L1685:
	;
	v11993 = v11665
	v11995 = v8405
	v12009 = v11668
	goto L1686
L1686:
	;
	if v10805 != 0 {
		goto L1754
	} else {
		goto L1755
	}
L1687:
	;
	v11673 = *(*int32)(unsafe.Add(mBase, uint32(v11672)+4))
	v11675 = v11673
	goto L1689
L1688:
	;
	v11675 = int32(0)
	goto L1689
L1689:
	;
	v11676 = F_palloc0(m, v11675)
	mBase = m.M
	v11677 = m.ExcPending
	if v11677 != 0 {
		goto L1
	} else {
		goto L1690
	}
L1690:
	;
	v11678 = F_palloc0(m, v11675)
	mBase = m.M
	v11679 = m.ExcPending
	if v11679 != 0 {
		goto L1
	} else {
		goto L1691
	}
L1691:
	;
	v11680 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	if v11680 == int32(0) {
		v11950 = v8405
		v11987 = v11665
		goto L1692
	} else {
		goto L1693
	}
L1692:
	;
	v11988 = *(*int32)(unsafe.Add(mBase, uint32(v11987)+4))
	v11989 = F_is_parallel_safe(m, v8377, v11988)
	mBase = m.M
	v11990 = m.ExcPending
	if v11990 != 0 {
		goto L1
	} else {
		goto L1753
	}
L1693:
	;
	v11683 = *(*int32)(unsafe.Add(mBase, uint32(v11680)+4))
	if v11683 <= int32(0) {
		v11950 = v8405
		v11987 = v11665
		goto L1692
	} else {
		goto L1694
	}
L1694:
	;
	v11686 = int32(0)
	v11698 = v11686
	v11699 = v11686
	v11700 = v11686
	v11704 = v11686
	v11712 = v11686
	goto L1695
L1695:
	;
	v11733 = v11698 << (uint(int32(2)) % 32)
	v11734 = *(*int32)(unsafe.Add(mBase, uint32(v11680)+12))
	v11736 = *(*int32)(unsafe.Add(mBase, uint32(v11733+v11734)))
	v11737 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+8))
	if v11737 != 0 {
		goto L1699
	} else {
		goto L1700
	}
L1696:
	;
	v11790 = int32(1)
	v11792 = v11785 & (v11783 ^ v11790)
	if (v11792|v11784)&v11790 != 0 {
		goto L1719
	} else {
		goto L1720
	}
L1697:
	;
	v11787 = v11698 + int32(1)
	v11788 = *(*int32)(unsafe.Add(mBase, uint32(v11680)+4))
	if v11787 < v11788 {
		v11698 = v11787
		v11699 = v11782
		v11700 = v11783
		v11704 = v11784
		v11712 = v11785
		goto L1695
	} else {
		goto L1718
	}
L1698:
	;
	if v11700&int32(1) != 0 {
		goto L1713
	} else {
		goto L1714
	}
L1699:
	;
	v11739 = *(*int32)(unsafe.Add(mBase, uint32(v11733+v11737)))
	if v11739 != 0 {
		goto L1698
	} else {
		goto L1702
	}
L1700:
	;
	goto L1701
L1701:
	;
	v11740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11671)+38)))
	if v11740 != int32(1) {
		goto L1703
	} else {
		goto L1704
	}
L1702:
	;
	goto L1701
L1703:
	;
	v11751 = F_contain_volatile_functions(m, v11736)
	mBase = m.M
	v11752 = m.ExcPending
	if v11752 != 0 {
		goto L1
	} else {
		goto L1707
	}
L1704:
	;
	v11743 = F_expression_returns_set(m, v11736)
	mBase = m.M
	v11744 = m.ExcPending
	if v11744 != 0 {
		goto L1
	} else {
		goto L1705
	}
L1705:
	;
	if v11743 == int32(0) {
		goto L1703
	} else {
		goto L1706
	}
L1706:
	;
	v11747 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11698+v11676))) = uint8(v11747)
	v11782 = v11699
	v11783 = v11700
	v11784 = v11704
	v11785 = v11747
	goto L1697
L1707:
	;
	if v11751 != 0 {
		goto L1708
	} else {
		goto L1709
	}
L1708:
	;
	v11753 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11678+v11698))) = uint8(v11753)
	v11782 = v11699
	v11783 = v11700
	v11784 = v11753
	v11785 = v11712
	goto L1697
L1709:
	;
	goto L1710
L1710:
	;
	F_cost_qual_eval_node(m, v8382+int32(192), v11736, v8377)
	mBase = m.M
	v11760 = m.ExcPending
	if v11760 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1711:
	;
	v11761 = *(*float64)(unsafe.Add(mBase, uint32(v8382)+200))
	v11763 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[5]))
	if base.F64_gt(v11761, base.F64_mul(v11763, float64(10))) == int32(0) {
		v11782 = v11699
		v11783 = v11700
		v11784 = v11704
		v11785 = v11712
		goto L1697
	} else {
		goto L1712
	}
L1712:
	;
	v11769 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11678+v11698))) = uint8(v11769)
	v11782 = v11769
	v11783 = v11700
	v11784 = v11704
	v11785 = v11712
	goto L1697
L1713:
	;
	v11782 = v11699
	v11783 = int32(1)
	v11784 = v11704
	v11785 = v11712
	goto L1697
L1714:
	;
	goto L1715
L1715:
	;
	v11777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11671)+38)))
	if v11777 != int32(1) {
		v11782 = v11699
		v11783 = int32(0)
		v11784 = v11704
		v11785 = v11712
		goto L1697
	} else {
		goto L1716
	}
L1716:
	;
	v11780 = F_expression_returns_set(m, v11736)
	mBase = m.M
	v11781 = m.ExcPending
	if v11781 != 0 {
		goto L1
	} else {
		goto L1717
	}
L1717:
	;
	v11782 = v11699
	v11783 = v11780
	v11784 = v11704
	v11785 = v11712
	goto L1697
L1718:
	;
	goto L1696
L1719:
	;
	v11806 = F_create_empty_pathtarget(m)
	mBase = m.M
	v11807 = m.ExcPending
	if v11807 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1720:
	;
	if v11782&int32(1) == int32(0) {
		v11950 = v8405
		v11987 = v11665
		goto L1692
	} else {
		goto L1721
	}
L1721:
	;
	v11800 = *(*int32)(unsafe.Add(mBase, uint32(v11671)+132))
	if v11800 != 0 {
		goto L1719
	} else {
		goto L1722
	}
L1722:
	;
	v11801 = *(*float64)(unsafe.Add(mBase, uint32(v8377)+296))
	if base.F64_gt(v11801, float64(0)) == int32(0) {
		v11950 = v8405
		v11987 = v11665
		goto L1692
	} else {
		goto L1723
	}
L1723:
	;
	goto L1719
L1724:
	;
	v11808 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	if v11808 == int32(0) {
		goto L1726
	} else {
		goto L1727
	}
L1725:
	;
	v11932 = F_pull_var_clause(m, v11890, int32(21))
	mBase = m.M
	v11933 = m.ExcPending
	if v11933 != 0 {
		goto L1
	} else {
		goto L1745
	}
L1726:
	;
	v11890 = int32(0)
	goto L1725
L1727:
	;
	goto L1728
L1728:
	;
	v11812 = int32(0)
	v11813 = *(*int32)(unsafe.Add(mBase, uint32(v11808)+4))
	if v11813 <= v11812 {
		v11890 = v11812
		goto L1725
	} else {
		goto L1729
	}
L1729:
	;
	v11817 = v11812
	v11824 = int32(0)
	goto L1730
L1730:
	;
	v11859 = v11824 << (uint(int32(2)) % 32)
	v11860 = *(*int32)(unsafe.Add(mBase, uint32(v11808)+12))
	v11862 = *(*int32)(unsafe.Add(mBase, uint32(v11859+v11860)))
	v11864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11678+v11824))))
	if v11864 == int32(0) {
		goto L1734
	} else {
		goto L1735
	}
L1731:
	;
	v11890 = v11884
	goto L1725
L1732:
	;
	v11887 = v11824 + int32(1)
	v11888 = *(*int32)(unsafe.Add(mBase, uint32(v11808)+4))
	if v11887 < v11888 {
		v11817 = v11884
		v11824 = v11887
		goto L1730
	} else {
		goto L1744
	}
L1733:
	;
	v11877 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+8))
	if v11877 != 0 {
		goto L1740
	} else {
		goto L1741
	}
L1734:
	;
	if v11792&int32(1) == int32(0) {
		goto L1733
	} else {
		goto L1737
	}
L1735:
	;
	goto L1736
L1736:
	;
	v11875 = F_lappend(m, v11817, v11862)
	mBase = m.M
	v11876 = m.ExcPending
	if v11876 != 0 {
		goto L1
	} else {
		goto L1739
	}
L1737:
	;
	v11872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11824+v11676))))
	if v11872 != int32(1) {
		goto L1733
	} else {
		goto L1738
	}
L1738:
	;
	goto L1736
L1739:
	;
	v11884 = v11875
	goto L1732
L1740:
	;
	v11879 = *(*int32)(unsafe.Add(mBase, uint32(v11859+v11877)))
	v11881 = v11879
	goto L1742
L1741:
	;
	v11881 = int32(0)
	goto L1742
L1742:
	;
	F_add_column_to_pathtarget(m, v11806, v11862, v11881)
	mBase = m.M
	v11883 = m.ExcPending
	if v11883 != 0 {
		goto L1
	} else {
		goto L1743
	}
L1743:
	;
	v11884 = v11817
	goto L1732
L1744:
	;
	goto L1731
L1745:
	;
	F_add_new_columns_to_pathtarget(m, v11806, v11932)
	mBase = m.M
	v11935 = m.ExcPending
	if v11935 != 0 {
		goto L1
	} else {
		goto L1746
	}
L1746:
	;
	F_list_free(m, v11932)
	mBase = m.M
	v11937 = m.ExcPending
	if v11937 != 0 {
		goto L1
	} else {
		goto L1747
	}
L1747:
	;
	F_list_free(m, v11890)
	mBase = m.M
	v11939 = m.ExcPending
	if v11939 != 0 {
		goto L1
	} else {
		goto L1748
	}
L1748:
	;
	if v11792&int32(1) != 0 {
		goto L1749
	} else {
		goto L1750
	}
L1749:
	;
	v11943 = float64(-1)
	goto L1751
L1750:
	;
	v11943 = v8405
	goto L1751
L1751:
	;
	v11944 = F_set_pathtarget_cost_width(m, v8377, v11806)
	mBase = m.M
	v11945 = m.ExcPending
	if v11945 != 0 {
		goto L1
	} else {
		goto L1752
	}
L1752:
	;
	v11950 = v11943
	v11987 = v11944
	goto L1692
L1753:
	;
	v11993 = v11987
	v11995 = v11950
	v12009 = v11989
	goto L1686
L1754:
	;
	v12032 = int32(0)
	v12033 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	if v12032 < v12033 {
		goto L1757
	} else {
		goto L1758
	}
L1755:
	;
	v12566 = v11993
	v12578 = v12009
	goto L1756
L1756:
	;
	v12606 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+100))
	if v12606 != 0 {
		goto L1808
	} else {
		goto L1809
	}
L1757:
	;
	v12044 = v12032
	v12051 = int32(0)
	goto L1760
L1758:
	;
	v12296 = v12032
	goto L1759
L1759:
	;
	v12330 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+256))
	if v12330 == int32(0) {
		v12397 = v12296
		goto L1777
	} else {
		goto L1778
	}
L1760:
	;
	v12078 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+12))
	v12082 = *(*int32)(unsafe.Add(mBase, uint32(v12078+v12051<<(uint(int32(2))%32))))
	v12083 = *(*int32)(unsafe.Add(mBase, uint32(v12082)+12))
	if v12083 == int32(0) {
		v12150 = v12044
		goto L1762
	} else {
		goto L1763
	}
L1761:
	;
	v12296 = v12251
	goto L1759
L1762:
	;
	v12184 = *(*int32)(unsafe.Add(mBase, uint32(v12082)+16))
	if v12184 == int32(0) {
		v12251 = v12150
		goto L1769
	} else {
		goto L1770
	}
L1763:
	;
	v12086 = int32(0)
	v12087 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+4))
	if v12087 <= v12086 {
		v12150 = v12044
		goto L1762
	} else {
		goto L1764
	}
L1764:
	;
	v12097 = v12044
	v12098 = v12086
	goto L1765
L1765:
	;
	v12131 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+12))
	v12135 = *(*int32)(unsafe.Add(mBase, uint32(v12131+v12098<<(uint(int32(2))%32))))
	v12136 = *(*int32)(unsafe.Add(mBase, uint32(v12135)+4))
	v12137 = F_bms_add_member(m, v12097, v12136)
	mBase = m.M
	v12138 = m.ExcPending
	if v12138 != 0 {
		goto L1
	} else {
		goto L1767
	}
L1766:
	;
	v12150 = v12137
	goto L1762
L1767:
	;
	v12140 = v12098 + int32(1)
	v12141 = *(*int32)(unsafe.Add(mBase, uint32(v12083)+4))
	if v12140 < v12141 {
		v12097 = v12137
		v12098 = v12140
		goto L1765
	} else {
		goto L1768
	}
L1768:
	;
	goto L1766
L1769:
	;
	v12286 = v12051 + int32(1)
	v12287 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	if v12286 < v12287 {
		v12044 = v12251
		v12051 = v12286
		goto L1760
	} else {
		goto L1776
	}
L1770:
	;
	v12187 = int32(0)
	v12188 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+4))
	if v12188 <= v12187 {
		v12251 = v12150
		goto L1769
	} else {
		goto L1771
	}
L1771:
	;
	v12198 = v12150
	v12199 = v12187
	goto L1772
L1772:
	;
	v12232 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+12))
	v12236 = *(*int32)(unsafe.Add(mBase, uint32(v12232+v12199<<(uint(int32(2))%32))))
	v12237 = *(*int32)(unsafe.Add(mBase, uint32(v12236)+4))
	v12238 = F_bms_add_member(m, v12198, v12237)
	mBase = m.M
	v12239 = m.ExcPending
	if v12239 != 0 {
		goto L1
	} else {
		goto L1774
	}
L1773:
	;
	v12251 = v12238
	goto L1769
L1774:
	;
	v12241 = v12199 + int32(1)
	v12242 = *(*int32)(unsafe.Add(mBase, uint32(v12184)+4))
	if v12241 < v12242 {
		v12198 = v12238
		v12199 = v12241
		goto L1772
	} else {
		goto L1775
	}
L1775:
	;
	goto L1773
L1776:
	;
	goto L1761
L1777:
	;
	v12431 = F_create_empty_pathtarget(m)
	mBase = m.M
	v12432 = m.ExcPending
	if v12432 != 0 {
		goto L1
	} else {
		goto L1784
	}
L1778:
	;
	v12333 = int32(0)
	v12334 = *(*int32)(unsafe.Add(mBase, uint32(v12330)+4))
	if v12334 <= v12333 {
		v12397 = v12296
		goto L1777
	} else {
		goto L1779
	}
L1779:
	;
	v12344 = v12296
	v12345 = v12333
	goto L1780
L1780:
	;
	v12378 = *(*int32)(unsafe.Add(mBase, uint32(v12330)+12))
	v12382 = *(*int32)(unsafe.Add(mBase, uint32(v12378+v12345<<(uint(int32(2))%32))))
	v12383 = *(*int32)(unsafe.Add(mBase, uint32(v12382)+4))
	v12384 = F_bms_add_member(m, v12344, v12383)
	mBase = m.M
	v12385 = m.ExcPending
	if v12385 != 0 {
		goto L1
	} else {
		goto L1782
	}
L1781:
	;
	v12397 = v12384
	goto L1777
L1782:
	;
	v12387 = v12345 + int32(1)
	v12388 = *(*int32)(unsafe.Add(mBase, uint32(v12330)+4))
	if v12387 < v12388 {
		v12344 = v12384
		v12345 = v12387
		goto L1780
	} else {
		goto L1783
	}
L1783:
	;
	goto L1781
L1784:
	;
	v12433 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	if v12433 == int32(0) {
		goto L1786
	} else {
		goto L1787
	}
L1785:
	;
	v12552 = F_pull_var_clause(m, v12524, int32(25))
	mBase = m.M
	v12553 = m.ExcPending
	if v12553 != 0 {
		goto L1
	} else {
		goto L1801
	}
L1786:
	;
	v12524 = int32(0)
	goto L1785
L1787:
	;
	goto L1788
L1788:
	;
	v12437 = int32(0)
	v12438 = *(*int32)(unsafe.Add(mBase, uint32(v12433)+4))
	if v12438 <= v12437 {
		v12524 = v12437
		goto L1785
	} else {
		goto L1789
	}
L1789:
	;
	v12450 = int32(0)
	v12456 = v12437
	goto L1790
L1790:
	;
	v12484 = v12450 << (uint(int32(2)) % 32)
	v12485 = *(*int32)(unsafe.Add(mBase, uint32(v12433)+12))
	v12487 = *(*int32)(unsafe.Add(mBase, uint32(v12484+v12485)))
	v12488 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+8))
	if v12488 == int32(0) {
		goto L1793
	} else {
		goto L1794
	}
L1791:
	;
	v12524 = v12505
	goto L1785
L1792:
	;
	v12507 = v12450 + int32(1)
	v12508 = *(*int32)(unsafe.Add(mBase, uint32(v12433)+4))
	if v12507 < v12508 {
		v12450 = v12507
		v12456 = v12505
		goto L1790
	} else {
		goto L1800
	}
L1793:
	;
	v12502 = F_lappend(m, v12456, v12487)
	mBase = m.M
	v12503 = m.ExcPending
	if v12503 != 0 {
		goto L1
	} else {
		goto L1799
	}
L1794:
	;
	v12492 = *(*int32)(unsafe.Add(mBase, uint32(v12484+v12488)))
	if v12492 == int32(0) {
		goto L1793
	} else {
		goto L1795
	}
L1795:
	;
	v12495 = F_bms_is_member(m, v12492, v12397)
	mBase = m.M
	v12496 = m.ExcPending
	if v12496 != 0 {
		goto L1
	} else {
		goto L1796
	}
L1796:
	;
	if v12495 == int32(0) {
		goto L1793
	} else {
		goto L1797
	}
L1797:
	;
	F_add_column_to_pathtarget(m, v12431, v12487, v12492)
	mBase = m.M
	v12500 = m.ExcPending
	if v12500 != 0 {
		goto L1
	} else {
		goto L1798
	}
L1798:
	;
	v12505 = v12456
	goto L1792
L1799:
	;
	v12505 = v12502
	goto L1792
L1800:
	;
	goto L1791
L1801:
	;
	F_add_new_columns_to_pathtarget(m, v12431, v12552)
	mBase = m.M
	v12555 = m.ExcPending
	if v12555 != 0 {
		goto L1
	} else {
		goto L1802
	}
L1802:
	;
	F_list_free(m, v12552)
	mBase = m.M
	v12557 = m.ExcPending
	if v12557 != 0 {
		goto L1
	} else {
		goto L1803
	}
L1803:
	;
	F_list_free(m, v12524)
	mBase = m.M
	v12559 = m.ExcPending
	if v12559 != 0 {
		goto L1
	} else {
		goto L1804
	}
L1804:
	;
	v12560 = F_set_pathtarget_cost_width(m, v8377, v12431)
	mBase = m.M
	v12561 = m.ExcPending
	if v12561 != 0 {
		goto L1
	} else {
		goto L1805
	}
L1805:
	;
	v12562 = *(*int32)(unsafe.Add(mBase, uint32(v12560)+4))
	v12563 = F_is_parallel_safe(m, v8377, v12562)
	mBase = m.M
	v12564 = m.ExcPending
	if v12564 != 0 {
		goto L1
	} else {
		goto L1806
	}
L1806:
	;
	v12566 = v12560
	v12578 = v12563
	goto L1756
L1807:
	;
	v12864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+38)))
	if v12864 == int32(1) {
		goto L1865
	} else {
		goto L1866
	}
L1808:
	;
	v12615 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v12616 = F_create_empty_pathtarget(m)
	mBase = m.M
	v12617 = m.ExcPending
	if v12617 != 0 {
		goto L1
	} else {
		goto L1813
	}
L1809:
	;
	v12607 = *(*int32)(unsafe.Add(mBase, uint32(v8387)+108))
	if v12607 != 0 {
		goto L1808
	} else {
		goto L1810
	}
L1810:
	;
	v12608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+36)))
	if v12608 != 0 {
		goto L1808
	} else {
		goto L1811
	}
L1811:
	;
	v12610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+318)))
	if v12610 != int32(1) {
		v12829 = v12566
		v12836 = int32(0)
		v12863 = v12578
		goto L1807
	} else {
		goto L1812
	}
L1812:
	;
	goto L1808
L1813:
	;
	v12618 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+4))
	if v12618 == int32(0) {
		goto L1815
	} else {
		goto L1816
	}
L1814:
	;
	v12790 = *(*int32)(unsafe.Add(mBase, uint32(v12615)+112))
	if v12790 != 0 {
		goto L1848
	} else {
		goto L1849
	}
L1815:
	;
	v12749 = int32(0)
	goto L1814
L1816:
	;
	goto L1817
L1817:
	;
	v12622 = int32(0)
	v12623 = *(*int32)(unsafe.Add(mBase, uint32(v12618)+4))
	if v12623 <= v12622 {
		v12749 = v12622
		goto L1814
	} else {
		goto L1818
	}
L1818:
	;
	v12627 = v12622
	v12634 = int32(0)
	goto L1819
L1819:
	;
	v12669 = v12634 << (uint(int32(2)) % 32)
	v12670 = *(*int32)(unsafe.Add(mBase, uint32(v12618)+12))
	v12672 = *(*int32)(unsafe.Add(mBase, uint32(v12669+v12670)))
	v12673 = *(*int32)(unsafe.Add(mBase, uint32(v11665)+8))
	if v12673 == int32(0) {
		goto L1822
	} else {
		goto L1823
	}
L1820:
	;
	v12749 = v12741
	goto L1814
L1821:
	;
	v12746 = v12634 + int32(1)
	v12747 = *(*int32)(unsafe.Add(mBase, uint32(v12618)+4))
	if v12746 < v12747 {
		v12627 = v12741
		v12634 = v12746
		goto L1819
	} else {
		goto L1847
	}
L1822:
	;
	v12739 = F_lappend(m, v12627, v12672)
	mBase = m.M
	v12740 = m.ExcPending
	if v12740 != 0 {
		goto L1
	} else {
		goto L1846
	}
L1823:
	;
	v12677 = *(*int32)(unsafe.Add(mBase, uint32(v12673+v12669)))
	if v12677 == int32(0) {
		goto L1822
	} else {
		goto L1824
	}
L1824:
	;
	v12680 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+256))
	if v12680 == int32(0) {
		goto L1822
	} else {
		goto L1825
	}
L1825:
	;
	if v12680 != 0 {
		goto L1828
	} else {
		goto L1829
	}
L1826:
	;
	if v12718 == int32(0) {
		goto L1822
	} else {
		goto L1839
	}
L1827:
	;
	goto L1826
L1828:
	;
	v12686 = *(*int32)(unsafe.Add(mBase, uint32(v12680)+4))
	if v12686 <= int32(0) {
		v12718 = int32(0)
		goto L1827
	} else {
		goto L1831
	}
L1829:
	;
	goto L1830
L1830:
	;
	v12718 = int32(0)
	goto L1827
L1831:
	;
	v12689 = int32(0)
	if v12689 < v12686 {
		goto L1832
	} else {
		goto L1833
	}
L1832:
	;
	v12692 = v12686
	goto L1834
L1833:
	;
	v12692 = v12689
	goto L1834
L1834:
	;
	v12693 = *(*int32)(unsafe.Add(mBase, uint32(v12680)+12))
	v12696 = int32(0)
	goto L1835
L1835:
	;
	v12703 = *(*int32)(unsafe.Add(mBase, uint32(v12693+v12696<<(uint(int32(2))%32))))
	v12704 = *(*int32)(unsafe.Add(mBase, uint32(v12703)+4))
	if v12704 == v12677 {
		v12718 = v12703
		goto L1827
	} else {
		goto L1837
	}
L1836:
	;
	goto L1830
L1837:
	;
	v12707 = v12696 + int32(1)
	if v12707 != v12692 {
		v12696 = v12707
		goto L1835
	} else {
		goto L1838
	}
L1838:
	;
	goto L1836
L1839:
	;
	v12722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12615)+45)))
	if v12722 != int32(1) {
		v12734 = v12672
		goto L1840
	} else {
		goto L1841
	}
L1840:
	;
	F_add_column_to_pathtarget(m, v12616, v12734, v12677)
	mBase = m.M
	v12736 = m.ExcPending
	if v12736 != 0 {
		goto L1
	} else {
		goto L1845
	}
L1841:
	;
	v12725 = *(*int32)(unsafe.Add(mBase, uint32(v12615)+108))
	if v12725 == int32(0) {
		v12734 = v12672
		goto L1840
	} else {
		goto L1842
	}
L1842:
	;
	v12728 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+324))
	v12729 = F_bms_make_singleton(m, v12728)
	mBase = m.M
	v12730 = m.ExcPending
	if v12730 != 0 {
		goto L1
	} else {
		goto L1843
	}
L1843:
	;
	v12732 = F_remove_nulling_relids(m, v12672, v12729, int32(0))
	mBase = m.M
	v12733 = m.ExcPending
	if v12733 != 0 {
		goto L1
	} else {
		goto L1844
	}
L1844:
	;
	v12734 = v12732
	goto L1840
L1845:
	;
	v12741 = v12627
	goto L1821
L1846:
	;
	v12741 = v12739
	goto L1821
L1847:
	;
	goto L1820
L1848:
	;
	v12791 = F_lappend(m, v12749, v12790)
	mBase = m.M
	v12792 = m.ExcPending
	if v12792 != 0 {
		goto L1
	} else {
		goto L1851
	}
L1849:
	;
	v12793 = v12749
	goto L1850
L1850:
	;
	v12795 = F_pull_var_clause(m, v12793, int32(26))
	mBase = m.M
	v12796 = m.ExcPending
	if v12796 != 0 {
		goto L1
	} else {
		goto L1852
	}
L1851:
	;
	v12793 = v12791
	goto L1850
L1852:
	;
	v12797 = int32(1)
	v12798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12615)+45)))
	if v12798 != v12797 {
		v12810 = v12795
		goto L1853
	} else {
		goto L1854
	}
L1853:
	;
	F_add_new_columns_to_pathtarget(m, v12616, v12810)
	mBase = m.M
	v12812 = m.ExcPending
	if v12812 != 0 {
		goto L1
	} else {
		goto L1858
	}
L1854:
	;
	v12801 = *(*int32)(unsafe.Add(mBase, uint32(v12615)+108))
	if v12801 == int32(0) {
		v12810 = v12795
		goto L1853
	} else {
		goto L1855
	}
L1855:
	;
	v12804 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+324))
	v12805 = F_bms_make_singleton(m, v12804)
	mBase = m.M
	v12806 = m.ExcPending
	if v12806 != 0 {
		goto L1
	} else {
		goto L1856
	}
L1856:
	;
	v12808 = F_remove_nulling_relids(m, v12795, v12805, int32(0))
	mBase = m.M
	v12809 = m.ExcPending
	if v12809 != 0 {
		goto L1
	} else {
		goto L1857
	}
L1857:
	;
	v12810 = v12808
	goto L1853
L1858:
	;
	F_list_free(m, v12810)
	mBase = m.M
	v12814 = m.ExcPending
	if v12814 != 0 {
		goto L1
	} else {
		goto L1859
	}
L1859:
	;
	F_list_free(m, v12793)
	mBase = m.M
	v12816 = m.ExcPending
	if v12816 != 0 {
		goto L1
	} else {
		goto L1860
	}
L1860:
	;
	v12817 = F_set_pathtarget_cost_width(m, v8377, v12616)
	mBase = m.M
	v12818 = m.ExcPending
	if v12818 != 0 {
		goto L1
	} else {
		goto L1861
	}
L1861:
	;
	v12819 = *(*int32)(unsafe.Add(mBase, uint32(v12817)+4))
	v12820 = F_is_parallel_safe(m, v8377, v12819)
	mBase = m.M
	v12821 = m.ExcPending
	if v12821 != 0 {
		goto L1
	} else {
		goto L1862
	}
L1862:
	;
	v12829 = v12817
	v12836 = v12797
	v12863 = v12820
	goto L1807
L1863:
	;
	v12952 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+160))
	F_apply_scanjoin_target_to_paths(m, v8377, v11660, v12945, v12952, v12863, v12951)
	mBase = m.M
	v12954 = m.ExcPending
	if v12954 != 0 {
		goto L1
	} else {
		goto L1878
	}
L1864:
	;
	v12936 = *(*int32)(unsafe.Add(mBase, uint32(v12933)+4))
	if v12936 != int32(1) {
		goto L1874
	} else {
		goto L1875
	}
L1865:
	;
	F_split_pathtarget_at_srfs(m, v8377, v11665, v11993, v8382+int32(188), v8382+int32(184))
	mBase = m.M
	v12872 = m.ExcPending
	if v12872 != 0 {
		goto L1
	} else {
		goto L1868
	}
L1866:
	;
	goto L1867
L1867:
	;
	v12905 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+188)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+184)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+176)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+180)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+168)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+172)) = v12905
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+144)) = v12829
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+44)) = v12829
	v12923 = F_list_make1_impl(m, int32(1), v8382+int32(44))
	mBase = m.M
	v12924 = m.ExcPending
	if v12924 != 0 {
		goto L1
	} else {
		goto L1872
	}
L1868:
	;
	v12873 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+188))
	v12874 = *(*int32)(unsafe.Add(mBase, uint32(v12873)+12))
	v12875 = *(*int32)(unsafe.Add(mBase, uint32(v12874)))
	F_split_pathtarget_at_srfs(m, v8377, v11993, v12566, v8382+int32(180), v8382+int32(176))
	mBase = m.M
	v12881 = m.ExcPending
	if v12881 != 0 {
		goto L1
	} else {
		goto L1869
	}
L1869:
	;
	v12882 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+180))
	v12883 = *(*int32)(unsafe.Add(mBase, uint32(v12882)+12))
	v12884 = *(*int32)(unsafe.Add(mBase, uint32(v12883)))
	F_split_pathtarget_at_srfs_extended(m, v8377, v12566, v12829, v8382+int32(172), v8382+int32(168), int32(1))
	mBase = m.M
	v12891 = m.ExcPending
	if v12891 != 0 {
		goto L1
	} else {
		goto L1870
	}
L1870:
	;
	v12892 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+172))
	v12893 = *(*int32)(unsafe.Add(mBase, uint32(v12892)+12))
	v12894 = *(*int32)(unsafe.Add(mBase, uint32(v12893)))
	F_split_pathtarget_at_srfs(m, v8377, v12829, int32(0), v8382+int32(164), v8382+int32(160))
	mBase = m.M
	v12901 = m.ExcPending
	if v12901 != 0 {
		goto L1
	} else {
		goto L1871
	}
L1871:
	;
	v12902 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+164))
	v12903 = *(*int32)(unsafe.Add(mBase, uint32(v12902)+12))
	v12904 = *(*int32)(unsafe.Add(mBase, uint32(v12903)))
	v12930 = v12894
	v12931 = v12884
	v12932 = v12904
	v12933 = v12902
	v12934 = v12875
	goto L1864
L1872:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+164)) = v12923
	if v12923 != 0 {
		v12930 = v12566
		v12931 = v11993
		v12932 = v12829
		v12933 = v12923
		v12934 = v11665
		goto L1864
	} else {
		goto L1873
	}
L1873:
	;
	v12945 = v12905
	v12946 = v12566
	v12947 = v11993
	v12950 = v11665
	v12951 = int32(0)
	goto L1863
L1874:
	;
	v12945 = v12933
	v12946 = v12930
	v12947 = v12931
	v12950 = v12934
	v12951 = int32(0)
	goto L1863
L1875:
	;
	goto L1876
L1876:
	;
	v12939 = *(*int32)(unsafe.Add(mBase, uint32(v12932)+4))
	v12940 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+28))
	v12941 = *(*int32)(unsafe.Add(mBase, uint32(v12940)+4))
	v12942 = F_equal(m, v12939, v12941)
	mBase = m.M
	v12943 = m.ExcPending
	if v12943 != 0 {
		goto L1
	} else {
		goto L1877
	}
L1877:
	;
	v12944 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+164))
	v12945 = v12944
	v12946 = v12930
	v12947 = v12931
	v12950 = v12934
	v12951 = v12942
	goto L1863
L1878:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+248)) = v12950
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+252)) = v12950
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+244)) = v12947
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+240)) = v12947
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+236)) = v12947
	*(*int32)(unsafe.Add(mBase, uint32(v8377)+232)) = v12946
	if v12836 == int32(0) {
		goto L1880
	} else {
		goto L1881
	}
L1879:
	;
	if v10805 == int32(0) {
		goto L1976
	} else {
		goto L1977
	}
L1880:
	;
	v13304 = v11660
	goto L1879
L1881:
	;
	goto L1882
L1882:
	;
	v12963 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v12964 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8382)+336)) = v12964
	*(*int64)(unsafe.Add(mBase, uint32(v8382)+328)) = v12964
	*(*int64)(unsafe.Add(mBase, uint32(v8382)+320)) = v12964
	*(*int64)(unsafe.Add(mBase, uint32(v8382)+312)) = v12964
	*(*int64)(unsafe.Add(mBase, uint32(v8382)+304)) = v12964
	F_get_agg_clause_costs(m, v8377, int32(0), v8382+int32(304))
	mBase = m.M
	v12978 = m.ExcPending
	if v12978 != 0 {
		goto L1
	} else {
		goto L1883
	}
L1883:
	;
	v12979 = *(*int32)(unsafe.Add(mBase, uint32(v12963)+112))
	v12980 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+4))
	v12987 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v12980))|base.B2i32(int32(1)<<(uint(v12980)%32)&int32(44) == v12987) == v12987 {
		goto L1885
	} else {
		goto L1886
	}
L1884:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13002)+28)) = v12946
	v13004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11660)+26)))
	if v12578&v13004 != int32(1) {
		goto L1890
	} else {
		goto L1891
	}
L1885:
	;
	v12993 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+8))
	v12994 = F_fetch_upper_rel(m, v8377, int32(2), v12993)
	mBase = m.M
	v12995 = m.ExcPending
	if v12995 != 0 {
		goto L1
	} else {
		goto L1888
	}
L1886:
	;
	goto L1887
L1887:
	;
	v13000 = F_fetch_upper_rel(m, v8377, int32(2), int32(0))
	mBase = m.M
	v13001 = m.ExcPending
	if v13001 != 0 {
		goto L1
	} else {
		goto L1889
	}
L1888:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12994)+4)) = int32(5)
	v13002 = v12994
	goto L1884
L1889:
	;
	v13002 = v13000
	goto L1884
L1890:
	;
	v13014 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v13002)+156)) = v13014
	v13016 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v13002)+160)) = v13016
	v13018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11660)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13002)+164)) = uint8(v13018)
	v13020 = *(*int32)(unsafe.Add(mBase, uint32(v11660)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v13002)+168)) = v13020
	v13022 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v13023 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+318)))
	if v13023 == int32(0) {
		goto L1896
	} else {
		goto L1897
	}
L1891:
	;
	v13008 = F_is_parallel_safe(m, v8377, v12979)
	mBase = m.M
	v13009 = m.ExcPending
	if v13009 != 0 {
		goto L1
	} else {
		goto L1892
	}
L1892:
	;
	if v13008 == int32(0) {
		goto L1890
	} else {
		goto L1893
	}
L1893:
	;
	v13012 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13002)+26)) = uint8(v13012)
	goto L1890
L1894:
	;
	F_set_cheapest(m, v13002)
	mBase = m.M
	v13291 = m.ExcPending
	if v13291 != 0 {
		goto L1
	} else {
		goto L1972
	}
L1895:
	;
	if v8390 != 0 {
		goto L1916
	} else {
		goto L1917
	}
L1896:
	;
	v13026 = *(*int32)(unsafe.Add(mBase, uint32(v13022)+108))
	if v13026 == int32(0) {
		goto L1895
	} else {
		goto L1899
	}
L1897:
	;
	goto L1898
L1898:
	;
	v13029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13022)+36)))
	if v13029 != 0 {
		goto L1895
	} else {
		goto L1900
	}
L1899:
	;
	goto L1898
L1900:
	;
	v13030 = *(*int32)(unsafe.Add(mBase, uint32(v13022)+100))
	if v13030 != 0 {
		goto L1895
	} else {
		goto L1901
	}
L1901:
	;
	v13031 = *(*int32)(unsafe.Add(mBase, uint32(v13022)+108))
	if v13031 == int32(0) {
		goto L1902
	} else {
		goto L1903
	}
L1902:
	;
	v13100 = *(*int32)(unsafe.Add(mBase, uint32(v13002)+28))
	v13101 = *(*int32)(unsafe.Add(mBase, uint32(v13022)+112))
	v13102 = F_create_group_result_path(m, v8377, v13002, v13100, v13101)
	mBase = m.M
	v13103 = m.ExcPending
	if v13103 != 0 {
		goto L1
	} else {
		goto L1912
	}
L1903:
	;
	v13034 = *(*int32)(unsafe.Add(mBase, uint32(v13031)+4))
	if v13034 < int32(2) {
		goto L1902
	} else {
		goto L1904
	}
L1904:
	;
	v13045 = v13034
	v13055 = int32(0)
	goto L1905
L1905:
	;
	v13081 = *(*int32)(unsafe.Add(mBase, uint32(v13002)+28))
	v13082 = *(*int32)(unsafe.Add(mBase, uint32(v13022)+112))
	v13083 = F_create_group_result_path(m, v8377, v13002, v13081, v13082)
	mBase = m.M
	v13084 = m.ExcPending
	if v13084 != 0 {
		goto L1
	} else {
		goto L1907
	}
L1906:
	;
	v13089 = int32(0)
	v13095 = F_create_append_path(m, v8377, v13002, v13085, v13089, v13089, v13089, v13089, v13089, float64(-1))
	mBase = m.M
	v13096 = m.ExcPending
	if v13096 != 0 {
		goto L1
	} else {
		goto L1910
	}
L1907:
	;
	v13085 = F_lappend(m, v13055, v13083)
	mBase = m.M
	v13086 = m.ExcPending
	if v13086 != 0 {
		goto L1
	} else {
		goto L1908
	}
L1908:
	;
	if base.Ui32(int32(1)) < base.Ui32(v13045) {
		v13045 = v13045 - int32(1)
		v13055 = v13085
		goto L1905
	} else {
		goto L1909
	}
L1909:
	;
	goto L1906
L1910:
	;
	F_add_path(m, v13002, v13095)
	mBase = m.M
	v13098 = m.ExcPending
	if v13098 != 0 {
		goto L1
	} else {
		goto L1911
	}
L1911:
	;
	goto L1894
L1912:
	;
	F_add_path(m, v13002, v13102)
	mBase = m.M
	v13105 = m.ExcPending
	if v13105 != 0 {
		goto L1
	} else {
		goto L1913
	}
L1913:
	;
	goto L1894
L1914:
	;
	v13159 = *(*int32)(unsafe.Add(mBase, uint32(v12963)+100))
	if v13159 == int32(0) {
		v13207 = v13158
		goto L1934
	} else {
		goto L1935
	}
L1915:
	;
	v13158 = int32(1)
	goto L1914
L1916:
	;
	v13106 = *(*int32)(unsafe.Add(mBase, uint32(v8390)))
	if v13106 != 0 {
		goto L1915
	} else {
		goto L1919
	}
L1917:
	;
	goto L1918
L1918:
	;
	v13107 = int32(0)
	v13108 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+256))
	if v13108 == v13107 {
		goto L1921
	} else {
		goto L1922
	}
L1919:
	;
	goto L1918
L1920:
	;
	if v13153 == int32(0) {
		v13158 = v13107
		goto L1914
	} else {
		goto L1933
	}
L1921:
	;
	v13153 = int32(1)
	goto L1920
L1922:
	;
	goto L1923
L1923:
	;
	v13117 = *(*int32)(unsafe.Add(mBase, uint32(v13108)+4))
	if v13117 <= int32(0) {
		v13145 = int32(1)
		goto L1924
	} else {
		goto L1925
	}
L1924:
	;
	v13153 = v13145
	goto L1920
L1925:
	;
	v13120 = int32(0)
	if v13120 < v13117 {
		goto L1926
	} else {
		goto L1927
	}
L1926:
	;
	v13123 = v13117
	goto L1928
L1927:
	;
	v13123 = v13120
	goto L1928
L1928:
	;
	v13124 = *(*int32)(unsafe.Add(mBase, uint32(v13108)+12))
	v13126 = int32(0)
	goto L1929
L1929:
	;
	v13134 = *(*int32)(unsafe.Add(mBase, uint32(v13124+v13126<<(uint(int32(2))%32))))
	v13135 = *(*int32)(unsafe.Add(mBase, uint32(v13134)+12))
	v13136 = int32(0)
	v13137 = base.B2i32(v13135 != v13136)
	if v13135 == v13136 {
		v13145 = v13137
		goto L1924
	} else {
		goto L1931
	}
L1930:
	;
	v13145 = v13137
	goto L1924
L1931:
	;
	v13141 = v13126 + int32(1)
	if v13141 != v13123 {
		v13126 = v13141
		goto L1929
	} else {
		goto L1932
	}
L1932:
	;
	goto L1930
L1933:
	;
	goto L1915
L1934:
	;
	v13208 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v13209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13208)+36)))
	if v13209 == int32(0) {
		goto L1957
	} else {
		goto L1958
	}
L1935:
	;
	v13162 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+336))
	if v13162 != 0 {
		v13207 = v13158
		goto L1934
	} else {
		goto L1936
	}
L1936:
	;
	if v8390 != 0 {
		goto L1938
	} else {
		goto L1939
	}
L1937:
	;
	v13207 = v13158 | int32(2)
	goto L1934
L1938:
	;
	v13163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8390)+16)))
	if v13163 != 0 {
		goto L1937
	} else {
		goto L1941
	}
L1939:
	;
	goto L1940
L1940:
	;
	v13164 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+256))
	v13165 = int32(0)
	if v13164 == v13165 {
		goto L1943
	} else {
		goto L1944
	}
L1941:
	;
	v13207 = v13158
	goto L1934
L1942:
	;
	if v13202 == int32(0) {
		v13207 = v13158
		goto L1934
	} else {
		goto L1955
	}
L1943:
	;
	v13202 = int32(1)
	goto L1942
L1944:
	;
	goto L1945
L1945:
	;
	v13172 = *(*int32)(unsafe.Add(mBase, uint32(v13164)+4))
	if v13172 <= int32(0) {
		v13196 = int32(1)
		goto L1946
	} else {
		goto L1947
	}
L1946:
	;
	v13202 = v13196
	goto L1942
L1947:
	;
	v13175 = int32(0)
	if v13175 < v13172 {
		goto L1948
	} else {
		goto L1949
	}
L1948:
	;
	v13178 = v13172
	goto L1950
L1949:
	;
	v13178 = v13175
	goto L1950
L1950:
	;
	v13179 = *(*int32)(unsafe.Add(mBase, uint32(v13164)+12))
	v13183 = v13165
	goto L1951
L1951:
	;
	v13187 = *(*int32)(unsafe.Add(mBase, uint32(v13179+v13183<<(uint(int32(2))%32))))
	v13188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13187)+18)))
	if v13188 != int32(1) {
		v13196 = v13188
		goto L1946
	} else {
		goto L1953
	}
L1952:
	;
	v13196 = v13188
	goto L1946
L1953:
	;
	v13192 = v13183 + int32(1)
	if v13192 != v13178 {
		v13183 = v13192
		goto L1951
	} else {
		goto L1954
	}
L1954:
	;
	goto L1952
L1955:
	;
	goto L1937
L1956:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8382)+280)) = uint8(v12578)
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+192)) = v13221
	v13224 = *(*int32)(unsafe.Add(mBase, uint32(v12963)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+284)) = v13224
	v13226 = *(*int32)(unsafe.Add(mBase, uint32(v12963)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+288)) = v13226
	v13228 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8382)+196)) = uint8(v13228)
	v13230 = int32(1)
	v13232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[6])))
	if v13232 == v13230 {
		goto L1967
	} else {
		goto L1968
	}
L1957:
	;
	v13212 = *(*int32)(unsafe.Add(mBase, uint32(v13208)+100))
	if v13212 == int32(0) {
		v13221 = v13207
		goto L1956
	} else {
		goto L1960
	}
L1958:
	;
	goto L1959
L1959:
	;
	v13215 = *(*int32)(unsafe.Add(mBase, uint32(v13208)+108))
	if v13215 != 0 {
		v13221 = v13207
		goto L1956
	} else {
		goto L1961
	}
L1960:
	;
	goto L1959
L1961:
	;
	v13216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+340)))
	if v13216 != 0 {
		v13221 = v13207
		goto L1956
	} else {
		goto L1962
	}
L1962:
	;
	v13219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8377)+341)))
	if v13219 != 0 {
		goto L1963
	} else {
		goto L1964
	}
L1963:
	;
	v13220 = v13207
	goto L1965
L1964:
	;
	v13220 = v13207 | int32(4)
	goto L1965
L1965:
	;
	v13221 = v13220
	goto L1956
L1966:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8382)+292)) = v13239
	F_create_ordinary_grouping_paths(m, v8377, v11660, v13002, v8382+int32(304), v8390, v8382+int32(192), v8382+int32(348))
	mBase = m.M
	v13248 = m.ExcPending
	if v13248 != 0 {
		goto L1
	} else {
		goto L1971
	}
L1967:
	;
	v13235 = *(*int32)(unsafe.Add(mBase, uint32(v12963)+108))
	if v13235 == int32(0) {
		v13239 = v13230
		goto L1966
	} else {
		goto L1970
	}
L1968:
	;
	goto L1969
L1969:
	;
	v13239 = int32(0)
	goto L1966
L1970:
	;
	goto L1969
L1971:
	;
	goto L1894
L1972:
	;
	v13292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8387)+38)))
	if v13292 != int32(1) {
		v13304 = v13002
		goto L1879
	} else {
		goto L1973
	}
L1973:
	;
	v13295 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+172))
	v13296 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+168))
	F_adjust_paths_for_srfs(m, v8377, v13002, v13295, v13296)
	mBase = m.M
	v13298 = m.ExcPending
	if v13298 != 0 {
		goto L1
	} else {
		goto L1974
	}
L1974:
	;
	v13304 = v13002
	goto L1879
L1975:
	;
	v14607 = *(*int32)(unsafe.Add(mBase, uint32(v14582)+120))
	if v14607 == int32(0) {
		goto L2201
	} else {
		goto L2202
	}
L1976:
	;
	v14568 = v12947
	v14570 = v11995
	v14572 = v8377
	v14576 = v13304
	v14577 = v8382
	v14581 = v12950
	v14582 = v8387
	v14590 = v11668
	v14591 = v8396
	v14600 = v8405
	v14604 = v8409
	v14605 = v8410
	goto L1975
L1977:
	;
	goto L1978
L1978:
	;
	v13344 = F_fetch_upper_rel(m, v8377, int32(3), int32(0))
	mBase = m.M
	v13345 = m.ExcPending
	if v13345 != 0 {
		goto L1
	} else {
		goto L1979
	}
L1979:
	;
	v13346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13304)+26)))
	if v12009&v13346 != int32(1) {
		goto L1980
	} else {
		goto L1981
	}
L1980:
	;
	v13356 = *(*int32)(unsafe.Add(mBase, uint32(v13304)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v13344)+156)) = v13356
	v13358 = *(*int32)(unsafe.Add(mBase, uint32(v13304)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v13344)+160)) = v13358
	v13360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13304)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13344)+164)) = uint8(v13360)
	v13362 = *(*int32)(unsafe.Add(mBase, uint32(v13304)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v13344)+168)) = v13362
	v13364 = *(*int32)(unsafe.Add(mBase, uint32(v13304)+32))
	if v13364 == int32(0) {
		v14502 = v12947
		v14504 = v11995
		v14506 = v8377
		v14507 = v13362
		v14510 = v13344
		v14511 = v8382
		v14515 = v12950
		v14516 = v8387
		v14524 = v11668
		v14525 = v8396
		v14534 = v8405
		v14538 = v8409
		v14539 = v8410
		goto L1984
	} else {
		goto L1985
	}
L1981:
	;
	v13350 = F_is_parallel_safe(m, v8377, v10805)
	mBase = m.M
	v13351 = m.ExcPending
	if v13351 != 0 {
		goto L1
	} else {
		goto L1982
	}
L1982:
	;
	if v13350 == int32(0) {
		goto L1980
	} else {
		goto L1983
	}
L1983:
	;
	v13354 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13344)+26)) = uint8(v13354)
	goto L1980
L1984:
	;
	if v14507 == int32(0) {
		goto L2190
	} else {
		goto L2191
	}
L1985:
	;
	v13367 = *(*int32)(unsafe.Add(mBase, uint32(v13364)+4))
	if v13367 <= int32(0) {
		v14502 = v12947
		v14504 = v11995
		v14506 = v8377
		v14507 = v13362
		v14510 = v13344
		v14511 = v8382
		v14515 = v12950
		v14516 = v8387
		v14524 = v11668
		v14525 = v8396
		v14534 = v8405
		v14538 = v8409
		v14539 = v8410
		goto L1984
	} else {
		goto L1986
	}
L1986:
	;
	v13397 = v8398
	goto L1987
L1987:
	;
	v13411 = *(*int32)(unsafe.Add(mBase, uint32(v13364)+12))
	v13415 = *(*int32)(unsafe.Add(mBase, uint32(v13411+v13397<<(uint(int32(2))%32))))
	v13416 = *(*int32)(unsafe.Add(mBase, uint32(v13304)+48))
	if v13415 == v13416 {
		goto L1990
	} else {
		goto L1991
	}
L1988:
	;
	v14499 = *(*int32)(unsafe.Add(mBase, uint32(v13344)+168))
	v14502 = v12947
	v14504 = v11995
	v14506 = v8377
	v14507 = v14499
	v14510 = v13344
	v14511 = v8382
	v14515 = v12950
	v14516 = v8387
	v14524 = v11668
	v14525 = v8396
	v14534 = v8405
	v14538 = v8409
	v14539 = v8410
	goto L1984
L1989:
	;
	v14496 = v13397 + int32(1)
	v14497 = *(*int32)(unsafe.Add(mBase, uint32(v13364)+4))
	if v14496 < v14497 {
		v13397 = v14496
		goto L1987
	} else {
		goto L2189
	}
L1990:
	;
	v13503 = int32(0)
	v13505 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	if v13503 < v13505 {
		goto L2026
	} else {
		goto L2027
	}
L1991:
	;
	v13418 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+168))
	v13419 = *(*int32)(unsafe.Add(mBase, uint32(v13415)+64))
	v13421 = v8382 + int32(304)
	if v13418 == v13419 {
		goto L1994
	} else {
		goto L1995
	}
L1992:
	;
	if v13499 != 0 {
		goto L1990
	} else {
		goto L2024
	}
L1993:
	;
	v13487 = *(*int32)(unsafe.Add(mBase, uint32(v13418)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13487
	v13499 = int32(1)
	goto L1992
L1994:
	;
	if v13418 != 0 {
		goto L1993
	} else {
		goto L1997
	}
L1995:
	;
	goto L1996
L1996:
	;
	if v13418 == int32(0) {
		goto L1998
	} else {
		goto L1999
	}
L1997:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = int32(0)
	v13499 = int32(1)
	goto L1992
L1998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = int32(0)
	v13499 = int32(1)
	goto L1992
L1999:
	;
	goto L2000
L2000:
	;
	if v13419 == int32(0) {
		goto L2001
	} else {
		goto L2002
	}
L2001:
	;
	v13439 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13439
	v13499 = v13439
	goto L1992
L2002:
	;
	goto L2003
L2003:
	;
	v13442 = *(*int32)(unsafe.Add(mBase, uint32(v13419)+4))
	v13443 = int32(0)
	if v13443 < v13442 {
		goto L2004
	} else {
		goto L2005
	}
L2004:
	;
	v13446 = v13442
	goto L2006
L2005:
	;
	v13446 = v13443
	goto L2006
L2006:
	;
	v13447 = *(*int32)(unsafe.Add(mBase, uint32(v13418)+4))
	v13452 = int32(0)
	goto L2007
L2007:
	;
	if v13452 < v13447 {
		goto L2009
	} else {
		goto L2010
	}
L2009:
	;
	v13459 = *(*int32)(unsafe.Add(mBase, uint32(v13418)+12))
	v13463 = v13459 + v13452<<(uint(int32(2))%32)
	goto L2011
L2010:
	;
	v13463 = int32(0)
	goto L2011
L2011:
	;
	if v13452 == v13446 {
		goto L2012
	} else {
		goto L2013
	}
L2012:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13446
	v13499 = base.B2i32(v13463 == int32(0))
	goto L1992
L2013:
	;
	goto L2014
L2014:
	;
	v13469 = base.B2i32(v13463 == int32(0))
	if v13463 == int32(0) {
		goto L2015
	} else {
		goto L2016
	}
L2015:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13452
	v13499 = v13469
	goto L1992
L2016:
	;
	goto L2017
L2017:
	;
	v13473 = *(*int32)(unsafe.Add(mBase, uint32(v13419)+12))
	if v13473 == int32(0) {
		goto L2018
	} else {
		goto L2019
	}
L2018:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13452
	v13499 = v13469
	goto L1992
L2019:
	;
	goto L2020
L2020:
	;
	v13477 = *(*int32)(unsafe.Add(mBase, uint32(v13463)))
	v13481 = *(*int32)(unsafe.Add(mBase, uint32(v13473+v13452<<(uint(int32(2))%32))))
	if v13477 != v13481 {
		goto L2021
	} else {
		goto L2022
	}
L2021:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13421))) = v13452
	v13499 = int32(0)
	goto L1992
L2022:
	;
	v13452 = v13452 + int32(1)
	goto L2007
L2024:
	;
	v13500 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+304))
	if v13500 <= int32(0) {
		goto L1989
	} else {
		goto L2025
	}
L2025:
	;
	goto L1990
L2026:
	;
	v13511 = v12946
	v13515 = v13503
	v13527 = v13503
	v13529 = v13415
	goto L2029
L2027:
	;
	v14432 = v13415
	goto L2028
L2028:
	;
	F_add_path(m, v13344, v14432)
	mBase = m.M
	v14453 = m.ExcPending
	if v14453 != 0 {
		goto L1
	} else {
		goto L2188
	}
L2029:
	;
	v13549 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+12))
	v13552 = v13549 + v13515<<(uint(int32(2))%32)
	v13553 = *(*int32)(unsafe.Add(mBase, uint32(v13552)))
	v13554 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+264))
	v13555 = F_make_pathkeys_for_window(m, v8377, v13553, v13554)
	mBase = m.M
	v13556 = m.ExcPending
	if v13556 != 0 {
		goto L1
	} else {
		goto L2032
	}
L2030:
	;
	v14432 = v14060
	goto L2028
L2031:
	;
	v13653 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+12))
	v13654 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	if base.Ui32(v13552+int32(4)) < base.Ui32(v13653+v13654<<(uint(int32(2))%32)) {
		goto L2073
	} else {
		goto L2074
	}
L2032:
	;
	v13557 = *(*int32)(unsafe.Add(mBase, uint32(v13529)+64))
	v13559 = v8382 + int32(192)
	if v13555 == v13557 {
		goto L2035
	} else {
		goto L2036
	}
L2033:
	;
	if v13637 != 0 {
		v13650 = v13529
		goto L2031
	} else {
		goto L2065
	}
L2034:
	;
	v13625 = *(*int32)(unsafe.Add(mBase, uint32(v13555)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13625
	v13637 = int32(1)
	goto L2033
L2035:
	;
	if v13555 != 0 {
		goto L2034
	} else {
		goto L2038
	}
L2036:
	;
	goto L2037
L2037:
	;
	if v13555 == int32(0) {
		goto L2039
	} else {
		goto L2040
	}
L2038:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = int32(0)
	v13637 = int32(1)
	goto L2033
L2039:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = int32(0)
	v13637 = int32(1)
	goto L2033
L2040:
	;
	goto L2041
L2041:
	;
	if v13557 == int32(0) {
		goto L2042
	} else {
		goto L2043
	}
L2042:
	;
	v13577 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13577
	v13637 = v13577
	goto L2033
L2043:
	;
	goto L2044
L2044:
	;
	v13580 = *(*int32)(unsafe.Add(mBase, uint32(v13557)+4))
	v13581 = int32(0)
	if v13581 < v13580 {
		goto L2045
	} else {
		goto L2046
	}
L2045:
	;
	v13584 = v13580
	goto L2047
L2046:
	;
	v13584 = v13581
	goto L2047
L2047:
	;
	v13585 = *(*int32)(unsafe.Add(mBase, uint32(v13555)+4))
	v13590 = int32(0)
	goto L2048
L2048:
	;
	if v13590 < v13585 {
		goto L2050
	} else {
		goto L2051
	}
L2050:
	;
	v13597 = *(*int32)(unsafe.Add(mBase, uint32(v13555)+12))
	v13601 = v13597 + v13590<<(uint(int32(2))%32)
	goto L2052
L2051:
	;
	v13601 = int32(0)
	goto L2052
L2052:
	;
	if v13590 == v13584 {
		goto L2053
	} else {
		goto L2054
	}
L2053:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13584
	v13637 = base.B2i32(v13601 == int32(0))
	goto L2033
L2054:
	;
	goto L2055
L2055:
	;
	v13607 = base.B2i32(v13601 == int32(0))
	if v13601 == int32(0) {
		goto L2056
	} else {
		goto L2057
	}
L2056:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13590
	v13637 = v13607
	goto L2033
L2057:
	;
	goto L2058
L2058:
	;
	v13611 = *(*int32)(unsafe.Add(mBase, uint32(v13557)+12))
	if v13611 == int32(0) {
		goto L2059
	} else {
		goto L2060
	}
L2059:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13590
	v13637 = v13607
	goto L2033
L2060:
	;
	goto L2061
L2061:
	;
	v13615 = *(*int32)(unsafe.Add(mBase, uint32(v13601)))
	v13619 = *(*int32)(unsafe.Add(mBase, uint32(v13611+v13590<<(uint(int32(2))%32))))
	if v13615 != v13619 {
		goto L2062
	} else {
		goto L2063
	}
L2062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13559))) = v13590
	v13637 = int32(0)
	goto L2033
L2063:
	;
	v13590 = v13590 + int32(1)
	goto L2048
L2065:
	;
	v13638 = *(*int32)(unsafe.Add(mBase, uint32(v8382)+192))
	if v13638 != 0 {
		goto L2067
	} else {
		goto L2068
	}
L2066:
	;
	v13647 = F_create_incremental_sort_path(m, v8377, v13344, v13529, v13555, v13638, float64(-1))
	mBase = m.M
	v13648 = m.ExcPending
	if v13648 != 0 {
		goto L1
	} else {
		goto L2072
	}
L2067:
	;
	v13640 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	if v13640&int32(1) != 0 {
		goto L2066
	} else {
		goto L2070
	}
L2068:
	;
	goto L2069
L2069:
	;
	v13644 = F_create_sort_path(m, v13344, v13529, v13555, float64(-1))
	mBase = m.M
	v13645 = m.ExcPending
	if v13645 != 0 {
		goto L1
	} else {
		goto L2071
	}
L2070:
	;
	goto L2069
L2071:
	;
	v13650 = v13644
	goto L2031
L2072:
	;
	v13650 = v13647
	goto L2031
L2073:
	;
	v13659 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13511)+32)))
	v13660 = F_copy_pathtarget(m, v13511)
	mBase = m.M
	v13661 = m.ExcPending
	if v13661 != 0 {
		goto L1
	} else {
		goto L2076
	}
L2074:
	;
	v13781 = v12947
	v13788 = v13654
	goto L2075
L2075:
	;
	v13823 = v13788 - int32(1)
	v13824 = *(*int32)(unsafe.Add(mBase, uint32(v10807)+8))
	v13825 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+48))
	v13829 = *(*int32)(unsafe.Add(mBase, uint32(v13824+v13825<<(uint(int32(2))%32))))
	if v13829 == int32(0) {
		goto L2090
	} else {
		goto L2091
	}
L2076:
	;
	v13662 = *(*int32)(unsafe.Add(mBase, uint32(v10807)+8))
	v13663 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+48))
	v13667 = *(*int32)(unsafe.Add(mBase, uint32(v13662+v13663<<(uint(int32(2))%32))))
	if v13667 == int32(0) {
		v13773 = v13659
		goto L2077
	} else {
		goto L2078
	}
L2077:
	;
	v13774 = int64(1073741823)
	if v13774 <= v13773 {
		goto L2086
	} else {
		goto L2087
	}
L2078:
	;
	v13670 = int32(0)
	v13671 = *(*int32)(unsafe.Add(mBase, uint32(v13667)+4))
	if v13671 <= v13670 {
		v13773 = v13659
		goto L2077
	} else {
		goto L2079
	}
L2079:
	;
	v13681 = v13670
	v13714 = v13659
	goto L2080
L2080:
	;
	v13715 = *(*int32)(unsafe.Add(mBase, uint32(v13667)+12))
	v13719 = *(*int32)(unsafe.Add(mBase, uint32(v13715+v13681<<(uint(int32(2))%32))))
	F_add_column_to_pathtarget(m, v13660, v13719, int32(0))
	mBase = m.M
	v13722 = m.ExcPending
	if v13722 != 0 {
		goto L1
	} else {
		goto L2082
	}
L2081:
	;
	v13773 = v13728
	goto L2077
L2082:
	;
	v13723 = *(*int32)(unsafe.Add(mBase, uint32(v13719)+8))
	v13725 = F_get_typavgwidth(m, v13723, int32(-1))
	mBase = m.M
	v13726 = m.ExcPending
	if v13726 != 0 {
		goto L1
	} else {
		goto L2083
	}
L2083:
	;
	v13728 = v13714 + base.I64_extend_i32_s(v13725)
	v13730 = v13681 + int32(1)
	v13731 = *(*int32)(unsafe.Add(mBase, uint32(v13667)+4))
	if v13730 < v13731 {
		v13681 = v13730
		v13714 = v13728
		goto L2080
	} else {
		goto L2084
	}
L2084:
	;
	goto L2081
L2085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13660)+32)) = base.I32_wrap_i64(v13777)
	v13780 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	v13781 = v13660
	v13788 = v13780
	goto L2075
L2086:
	;
	v13777 = v13774
	goto L2088
L2087:
	;
	v13777 = v13773
	goto L2088
L2088:
	;
	goto L2085
L2089:
	;
	v14056 = base.B2i32(v13515 == v13823)
	if v13515 == v13823 {
		goto L2116
	} else {
		goto L2117
	}
L2090:
	;
	v13832 = int32(0)
	v14014 = v13832
	v14027 = v13832
	v14033 = v13527
	goto L2089
L2091:
	;
	goto L2092
L2092:
	;
	v13834 = int32(0)
	v13836 = *(*int32)(unsafe.Add(mBase, uint32(v13829)+4))
	if v13836 <= v13834 {
		v14014 = v13834
		v14027 = v13829
		v14033 = v13527
		goto L2089
	} else {
		goto L2093
	}
L2093:
	;
	v13839 = v13834
	v13846 = v13836
	v13857 = v13834
	v13858 = v13527
	goto L2094
L2094:
	;
	v13880 = *(*int32)(unsafe.Add(mBase, uint32(v13829)+12))
	v13884 = *(*int32)(unsafe.Add(mBase, uint32(v13880+v13857<<(uint(int32(2))%32))))
	v13885 = *(*int32)(unsafe.Add(mBase, uint32(v13884)+28))
	if v13885 == int32(0) {
		v13964 = v13839
		v13971 = v13846
		v13983 = v13858
		goto L2096
	} else {
		goto L2097
	}
L2095:
	;
	v14008 = *(*int32)(unsafe.Add(mBase, uint32(v10807)+8))
	v14009 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+48))
	v14013 = *(*int32)(unsafe.Add(mBase, uint32(v14008+v14009<<(uint(int32(2))%32))))
	v14014 = v13964
	v14027 = v14013
	v14033 = v13983
	goto L2089
L2096:
	;
	v14006 = v13857 + int32(1)
	if v14006 < v13971 {
		v13839 = v13964
		v13846 = v13971
		v13857 = v14006
		v13858 = v13983
		goto L2094
	} else {
		goto L2115
	}
L2097:
	;
	v13888 = int32(0)
	v13889 = *(*int32)(unsafe.Add(mBase, uint32(v13885)+4))
	if v13889 <= v13888 {
		v13964 = v13839
		v13971 = v13846
		v13983 = v13858
		goto L2096
	} else {
		goto L2098
	}
L2098:
	;
	v13892 = v13839
	v13900 = v13888
	v13911 = v13858
	goto L2099
L2099:
	;
	v13933 = *(*int32)(unsafe.Add(mBase, uint32(v13885)+12))
	v13937 = *(*int32)(unsafe.Add(mBase, uint32(v13933+v13900<<(uint(int32(2))%32))))
	v13938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13937)+12)))
	if v13938 != 0 {
		goto L2102
	} else {
		goto L2103
	}
L2100:
	;
	v13963 = *(*int32)(unsafe.Add(mBase, uint32(v13829)+4))
	v13964 = v13953
	v13971 = v13963
	v13983 = v13958
	goto L2096
L2101:
	;
	v13947 = F_copyObjectImpl(m, v13946)
	mBase = m.M
	v13948 = m.ExcPending
	if v13948 != 0 {
		goto L1
	} else {
		goto L2107
	}
L2102:
	;
	v13939 = F_copyObjectImpl(m, v13884)
	mBase = m.M
	v13940 = m.ExcPending
	if v13940 != 0 {
		goto L1
	} else {
		goto L2105
	}
L2103:
	;
	goto L2104
L2104:
	;
	v13942 = *(*int32)(unsafe.Add(mBase, uint32(v13937)+16))
	v13943 = F_copyObjectImpl(m, v13942)
	mBase = m.M
	v13944 = m.ExcPending
	if v13944 != 0 {
		goto L1
	} else {
		goto L2106
	}
L2105:
	;
	v13941 = *(*int32)(unsafe.Add(mBase, uint32(v13937)+16))
	v13945 = v13939
	v13946 = v13941
	goto L2101
L2106:
	;
	v13945 = v13943
	v13946 = v13884
	goto L2101
L2107:
	;
	v13949 = *(*int32)(unsafe.Add(mBase, uint32(v13937)+4))
	v13950 = *(*int32)(unsafe.Add(mBase, uint32(v13937)+8))
	v13951 = F_make_opclause(m, v13949, v13945, v13947, v13950)
	mBase = m.M
	v13952 = m.ExcPending
	if v13952 != 0 {
		goto L1
	} else {
		goto L2108
	}
L2108:
	;
	v13953 = F_lappend(m, v13892, v13951)
	mBase = m.M
	v13954 = m.ExcPending
	if v13954 != 0 {
		goto L1
	} else {
		goto L2109
	}
L2109:
	;
	if v13515 != v13823 {
		goto L2110
	} else {
		goto L2111
	}
L2110:
	;
	v13956 = F_lappend(m, v13911, v13951)
	mBase = m.M
	v13957 = m.ExcPending
	if v13957 != 0 {
		goto L1
	} else {
		goto L2113
	}
L2111:
	;
	v13958 = v13911
	goto L2112
L2112:
	;
	v13960 = v13900 + int32(1)
	v13961 = *(*int32)(unsafe.Add(mBase, uint32(v13885)+4))
	if v13960 < v13961 {
		v13892 = v13953
		v13900 = v13960
		v13911 = v13958
		goto L2099
	} else {
		goto L2114
	}
L2113:
	;
	v13958 = v13956
	goto L2112
L2114:
	;
	goto L2100
L2115:
	;
	goto L2095
L2116:
	;
	v14057 = v14033
	goto L2118
L2117:
	;
	v14057 = int32(0)
	goto L2118
L2118:
	;
	v14060 = F_palloc0(m, int32(96))
	mBase = m.M
	v14061 = m.ExcPending
	if v14061 != 0 {
		goto L1
	} else {
		goto L2119
	}
L2119:
	;
	v14062 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14060)+20)) = uint8(v14062)
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+16)) = v14062
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+12)) = v13781
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+8)) = v13344
	*(*int64)(unsafe.Add(mBase, uint32(v14060))) = int64(1571958030648)
	v14070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13344)+26)))
	if v14070 == int32(1) {
		goto L2120
	} else {
		goto L2121
	}
L2120:
	;
	v14073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13650)+21)))
	v14075 = v14073
	goto L2122
L2121:
	;
	v14075 = int32(0)
	goto L2122
L2122:
	;
	v14077 = v14075 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14060)+21)) = uint8(v14077)
	v14079 = *(*int32)(unsafe.Add(mBase, uint32(v13650)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+24)) = v14079
	v14081 = *(*int32)(unsafe.Add(mBase, uint32(v13650)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v14060)+88)) = uint8(v14056)
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+84)) = v14014
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+80)) = v14057
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+76)) = v13553
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+72)) = v13650
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+64)) = v14081
	v14088 = *(*int32)(unsafe.Add(mBase, uint32(v13650)+40))
	v14089 = *(*float64)(unsafe.Add(mBase, uint32(v13650)+48))
	v14090 = *(*float64)(unsafe.Add(mBase, uint32(v13650)+56))
	v14091 = *(*float64)(unsafe.Add(mBase, uint32(v13650)+32))
	v14092 = int32(0)
	v14094 = m.G0
	v14096 = v14094 - int32(48)
	m.G0 = v14096
	v14098 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+12))
	if v14098 != 0 {
		goto L2123
	} else {
		goto L2124
	}
L2123:
	;
	v14099 = *(*int32)(unsafe.Add(mBase, uint32(v14098)+4))
	v14100 = v14099
	goto L2125
L2124:
	;
	v14100 = int32(0)
	goto L2125
L2125:
	;
	v14101 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+16))
	if v14101 != 0 {
		goto L2126
	} else {
		goto L2127
	}
L2126:
	;
	v14102 = *(*int32)(unsafe.Add(mBase, uint32(v14101)+4))
	v14103 = v14102
	goto L2128
L2127:
	;
	v14103 = v14092
	goto L2128
L2128:
	;
	if v14027 == int32(0) {
		v14239 = v14090
		v14244 = v14089
		goto L2129
	} else {
		goto L2130
	}
L2129:
	;
	v14251 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[5]))
	v14253 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+48)) = v14244
	*(*int32)(unsafe.Add(mBase, uint32(v14060)+40)) = v14088
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+32)) = v14091
	v14263 = base.F64_add(base.F64_mul(v14253, v14091), base.F64_add(base.F64_mul(base.F64_mul(v14251, base.F64_convert_i32_s(v14103+v14100)), v14091), v14239))
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+56)) = v14263
	v14265 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+20))
	v14266 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+12))
	if v14266 != 0 {
		goto L2138
	} else {
		goto L2139
	}
L2130:
	;
	v14106 = *(*int32)(unsafe.Add(mBase, uint32(v14027)+4))
	if v14106 <= int32(0) {
		v14239 = v14090
		v14244 = v14089
		goto L2129
	} else {
		goto L2131
	}
L2131:
	;
	v14110 = v14096 + int32(32)
	v14125 = v14092
	v14141 = v14090
	v14146 = v14089
	goto L2132
L2132:
	;
	v14152 = *(*int32)(unsafe.Add(mBase, uint32(v14027)+12))
	v14156 = *(*int32)(unsafe.Add(mBase, uint32(v14152+v14125<<(uint(int32(2))%32))))
	v14157 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+16)) = v14157
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+8)) = v14157
	v14161 = *(*int32)(unsafe.Add(mBase, uint32(v14156)+4))
	F_add_function_cost(m, v8377, v14161, v14156, v14096+int32(8))
	mBase = m.M
	v14165 = m.ExcPending
	if v14165 != 0 {
		goto L1
	} else {
		goto L2134
	}
L2133:
	;
	v14239 = v14204
	v14244 = v14199
	goto L2129
L2134:
	;
	v14166 = *(*int32)(unsafe.Add(mBase, uint32(v14156)+20))
	v14167 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+16))
	v14168 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14096)+24)) = v8377
	v14170 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14110)+8)) = v14170
	*(*int64)(unsafe.Add(mBase, uint32(v14110))) = v14170
	v14175 = v14096 + int32(24)
	v14176 = F_cost_qual_eval_walker(m, v14166, v14175)
	mBase = m.M
	v14177 = m.ExcPending
	if v14177 != 0 {
		goto L1
	} else {
		goto L2135
	}
L2135:
	;
	v14178 = *(*int64)(unsafe.Add(mBase, uint32(v14110)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+16)) = v14178
	v14180 = *(*int64)(unsafe.Add(mBase, uint32(v14110)))
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+8)) = v14180
	v14182 = *(*int32)(unsafe.Add(mBase, uint32(v14156)+24))
	v14183 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+16))
	v14184 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14096)+24)) = v8377
	v14186 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14110)+8)) = v14186
	*(*int64)(unsafe.Add(mBase, uint32(v14110))) = v14186
	v14190 = F_cost_qual_eval_walker(m, v14182, v14175)
	mBase = m.M
	v14191 = m.ExcPending
	if v14191 != 0 {
		goto L1
	} else {
		goto L2136
	}
L2136:
	;
	v14192 = *(*int64)(unsafe.Add(mBase, uint32(v14110)))
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+8)) = v14192
	v14194 = *(*int64)(unsafe.Add(mBase, uint32(v14110)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v14096)+16)) = v14194
	v14198 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+8))
	v14199 = base.F64_add(base.F64_add(v14184, base.F64_add(v14146, v14168)), v14198)
	v14201 = *(*float64)(unsafe.Add(mBase, uint32(v14096)+16))
	v14204 = base.F64_add(base.F64_mul(base.F64_add(base.F64_add(v14167, v14183), v14201), v14091), v14141)
	v14206 = v14125 + int32(1)
	v14207 = *(*int32)(unsafe.Add(mBase, uint32(v14027)+4))
	if v14206 < v14207 {
		v14125 = v14206
		v14141 = v14204
		v14146 = v14199
		goto L2132
	} else {
		goto L2137
	}
L2137:
	;
	goto L2133
L2138:
	;
	v14267 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v14268 = *(*int32)(unsafe.Add(mBase, uint32(v14267)+76))
	v14269 = F_get_sortgrouplist_exprs(m, v14266, v14268)
	mBase = m.M
	v14270 = m.ExcPending
	if v14270 != 0 {
		goto L1
	} else {
		goto L2141
	}
L2139:
	;
	v14279 = v14091
	goto L2140
L2140:
	;
	v14281 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+16))
	if v14281 != 0 {
		goto L2144
	} else {
		goto L2145
	}
L2141:
	;
	v14271 = int32(0)
	v14273 = F_estimate_num_groups(m, v8377, v14269, v14091, v14271, v14271)
	mBase = m.M
	v14274 = m.ExcPending
	if v14274 != 0 {
		goto L1
	} else {
		goto L2142
	}
L2142:
	;
	F_list_free(m, v14269)
	mBase = m.M
	v14276 = m.ExcPending
	if v14276 != 0 {
		goto L1
	} else {
		goto L2143
	}
L2143:
	;
	v14279 = base.F64_div(v14091, v14273)
	goto L2140
L2144:
	;
	v14282 = *(*int32)(unsafe.Add(mBase, uint32(v8377)+4))
	v14283 = *(*int32)(unsafe.Add(mBase, uint32(v14282)+76))
	v14284 = F_get_sortgrouplist_exprs(m, v14281, v14283)
	mBase = m.M
	v14285 = m.ExcPending
	if v14285 != 0 {
		goto L1
	} else {
		goto L2147
	}
L2145:
	;
	v14296 = float64(1)
	goto L2146
L2146:
	;
	if v14265&int32(256) != 0 {
		v14351 = v14279
		goto L2150
	} else {
		goto L2151
	}
L2147:
	;
	v14286 = int32(0)
	v14288 = F_estimate_num_groups(m, v8377, v14284, v14279, v14286, v14286)
	mBase = m.M
	v14289 = m.ExcPending
	if v14289 != 0 {
		goto L1
	} else {
		goto L2148
	}
L2148:
	;
	F_list_free(m, v14284)
	mBase = m.M
	v14291 = m.ExcPending
	if v14291 != 0 {
		goto L1
	} else {
		goto L2149
	}
L2149:
	;
	v14296 = base.F64_div(v14279, v14288)
	goto L2146
L2150:
	;
	v14353 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+12))
	if v14353 == int32(0) {
		goto L2174
	} else {
		goto L2175
	}
L2151:
	;
	if v14265&int32(1024) != 0 {
		goto L2152
	} else {
		goto L2153
	}
L2152:
	;
	if base.B2i32(v14265&int32(10) == int32(0))|v14265&int32(4) != 0 {
		v14351 = float64(1)
		goto L2150
	} else {
		goto L2155
	}
L2153:
	;
	goto L2154
L2154:
	;
	v14311 = float64(1)
	if v14265&int32(_a_F_subquery_planner_46) != int32(_a_F_subquery_planner_47) {
		v14351 = v14311
		goto L2150
	} else {
		goto L2159
	}
L2155:
	;
	v14309 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+16))
	if v14309 != 0 {
		goto L2156
	} else {
		goto L2157
	}
L2156:
	;
	v14310 = v14296
	goto L2158
L2157:
	;
	v14310 = v14279
	goto L2158
L2158:
	;
	v14351 = v14310
	goto L2150
L2159:
	;
	v14316 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+28))
	v14317 = *(*int32)(unsafe.Add(mBase, uint32(v14316)))
	if v14317 == int32(7) {
		goto L2161
	} else {
		goto L2162
	}
L2160:
	;
	if v14265&int32(4) != 0 {
		goto L2169
	} else {
		goto L2170
	}
L2161:
	;
	v14321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14316)+24)))
	if v14321 != 0 {
		v14338 = float64(1)
		goto L2160
	} else {
		goto L2164
	}
L2162:
	;
	goto L2163
L2163:
	;
	v14338 = base.F64_mul(base.F64_div(v14279, v14296), float64(0.3333333333333333))
	goto L2160
L2164:
	;
	v14322 = *(*int32)(unsafe.Add(mBase, uint32(v14316)+4))
	switch v14322 - int32(20) {
	case 0:
		goto L2166
	case 1:
		goto L2168
	default:
		goto L2165
	case 3:
		goto L2167
	}
L2165:
	;
	v14338 = base.F64_mul(base.F64_div(v14279, v14296), float64(0.3333333333333333))
	goto L2160
L2166:
	;
	v14329 = *(*int32)(unsafe.Add(mBase, uint32(v14316)+20))
	v14330 = *(*int64)(unsafe.Add(mBase, uint32(v14329)))
	v14338 = base.F64_convert_i64_s(v14330)
	goto L2160
L2167:
	;
	v14327 = *(*int32)(unsafe.Add(mBase, uint32(v14316)+20))
	v14338 = base.F64_convert_i32_s(v14327)
	goto L2160
L2168:
	;
	v14325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14316)+20)))
	v14338 = base.F64_convert_i32_s(v14325)
	goto L2160
L2169:
	;
	v14351 = base.F64_add(v14338, float64(1))
	goto L2150
L2170:
	;
	goto L2171
L2171:
	;
	if v14265&int32(10) == int32(0) {
		v14351 = v14311
		goto L2150
	} else {
		goto L2172
	}
L2172:
	;
	v14351 = base.F64_mul(v14296, base.F64_add(v14338, float64(1)))
	goto L2150
L2173:
	;
	if base.F64_gt(v14279, v14361) != 0 {
		goto L2179
	} else {
		goto L2180
	}
L2174:
	;
	v14356 = *(*int32)(unsafe.Add(mBase, uint32(v13553)+16))
	if v14356 == int32(0) {
		v14361 = v14351
		goto L2173
	} else {
		goto L2177
	}
L2175:
	;
	goto L2176
L2176:
	;
	v14361 = base.F64_add(v14351, float64(1))
	goto L2173
L2177:
	;
	goto L2176
L2178:
	;
	m.G0 = v14096 + int32(48)
	v14395 = *(*float64)(unsafe.Add(mBase, uint32(v13781)+16))
	v14396 = *(*float64)(unsafe.Add(mBase, uint32(v14060)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+48)) = base.F64_add(v14395, v14396)
	v14399 = *(*float64)(unsafe.Add(mBase, uint32(v14060)+56))
	v14400 = *(*float64)(unsafe.Add(mBase, uint32(v13781)+24))
	v14401 = *(*float64)(unsafe.Add(mBase, uint32(v14060)+32))
	v14403 = *(*float64)(unsafe.Add(mBase, uint32(v13781)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+56)) = base.F64_add(v14399, base.F64_add(base.F64_mul(v14400, v14401), v14403))
	v14408 = v13515 + int32(1)
	v14409 = *(*int32)(unsafe.Add(mBase, uint32(v10805)+4))
	if v14408 < v14409 {
		v13511 = v13781
		v13515 = v14408
		v13527 = v14033
		v13529 = v14060
		goto L2029
	} else {
		goto L2187
	}
L2179:
	;
	v14364 = v14361
	goto L2181
L2180:
	;
	v14364 = v14279
	goto L2181
L2181:
	;
	if base.F64_gt(v14364, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14364)&int64(9223372036854775807))) == int32(0) {
		goto L2182
	} else {
		goto L2183
	}
L2182:
	;
	if base.F64_le(v14364, float64(1)) != 0 {
		goto L2178
	} else {
		goto L2185
	}
L2183:
	;
	v14382 = float64(1e+100)
	goto L2184
L2184:
	;
	v14388 = *(*float64)(unsafe.Add(mBase, uint32(v14060)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v14060)+48)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_sub(v14263, v14244), v14091), base.F64_add(v14382, float64(-1))), v14388)
	goto L2178
L2185:
	;
	v14377 = base.F64_nearest(v14364)
	if base.F64_gt(v14377, float64(1)) == int32(0) {
		goto L2178
	} else {
		goto L2186
	}
L2186:
	;
	v14382 = v14377
	goto L2184
L2187:
	;
	goto L2030
L2188:
	;
	goto L1989
L2189:
	;
	goto L1988
L2190:
	;
	v14552 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v14552 != 0 {
		goto L2194
	} else {
		goto L2195
	}
L2191:
	;
	v14543 = *(*int32)(unsafe.Add(mBase, uint32(v14507)+36))
	if v14543 == int32(0) {
		goto L2190
	} else {
		goto L2192
	}
L2192:
	;
	m.T0[v14543].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14506, int32(3), v13304, v14510, int32(0))
	mBase = m.M
	v14549 = m.ExcPending
	if v14549 != 0 {
		goto L1
	} else {
		goto L2193
	}
L2193:
	;
	goto L2190
L2194:
	;
	m.T0[v14552].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14506, int32(3), v13304, v14510, int32(0))
	mBase = m.M
	v14556 = m.ExcPending
	if v14556 != 0 {
		goto L1
	} else {
		goto L2197
	}
L2195:
	;
	goto L2196
L2196:
	;
	F_set_cheapest(m, v14510)
	mBase = m.M
	v14558 = m.ExcPending
	if v14558 != 0 {
		goto L1
	} else {
		goto L2198
	}
L2197:
	;
	goto L2196
L2198:
	;
	v14559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14516)+38)))
	if v14559 != int32(1) {
		v14568 = v14502
		v14570 = v14504
		v14572 = v14506
		v14576 = v14510
		v14577 = v14511
		v14581 = v14515
		v14582 = v14516
		v14590 = v14524
		v14591 = v14525
		v14600 = v14534
		v14604 = v14538
		v14605 = v14539
		goto L1975
	} else {
		goto L2199
	}
L2199:
	;
	v14562 = *(*int32)(unsafe.Add(mBase, uint32(v14511)+180))
	v14563 = *(*int32)(unsafe.Add(mBase, uint32(v14511)+176))
	F_adjust_paths_for_srfs(m, v14506, v14510, v14562, v14563)
	mBase = m.M
	v14565 = m.ExcPending
	if v14565 != 0 {
		goto L1
	} else {
		goto L2200
	}
L2200:
	;
	v14568 = v14502
	v14570 = v14504
	v14572 = v14506
	v14576 = v14510
	v14577 = v14511
	v14581 = v14515
	v14582 = v14516
	v14590 = v14524
	v14591 = v14525
	v14600 = v14534
	v14604 = v14538
	v14605 = v14539
	goto L1975
L2201:
	;
	v15197 = v14570
	v15199 = v14572
	v15204 = v14577
	v15205 = v14576
	v15208 = v14581
	v15209 = v14582
	v15217 = v14590
	v15218 = v14591
	v15227 = v14600
	v15231 = v14604
	v15232 = v14605
	goto L713
L2202:
	;
	goto L2203
L2203:
	;
	v14612 = F_fetch_upper_rel(m, v14572, int32(5), int32(0))
	mBase = m.M
	v14613 = m.ExcPending
	if v14613 != 0 {
		goto L1
	} else {
		goto L2204
	}
L2204:
	;
	v14614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14576)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14612)+26)) = uint8(v14614)
	v14616 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v14612)+156)) = v14616
	v14618 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v14612)+160)) = v14618
	v14620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14576)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14612)+164)) = uint8(v14620)
	v14622 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v14612)+168)) = v14622
	v14624 = F_create_final_distinct_paths(m, v14572, v14576, v14612)
	mBase = m.M
	v14625 = m.ExcPending
	if v14625 != 0 {
		goto L1
	} else {
		goto L2205
	}
L2205:
	;
	v14626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14576)+26)))
	if v14626 != int32(1) {
		goto L2206
	} else {
		goto L2207
	}
L2206:
	;
	v15171 = *(*int32)(unsafe.Add(mBase, uint32(v14624)+32))
	if v15171 == int32(0) {
		goto L712
	} else {
		goto L2332
	}
L2207:
	;
	v14629 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+40))
	if v14629 == int32(0) {
		goto L2206
	} else {
		goto L2208
	}
L2208:
	;
	v14632 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+4))
	v14633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14632)+40)))
	if v14633 != 0 {
		goto L2206
	} else {
		goto L2209
	}
L2209:
	;
	v14636 = F_fetch_upper_rel(m, v14572, int32(4), int32(0))
	mBase = m.M
	v14637 = m.ExcPending
	if v14637 != 0 {
		goto L1
	} else {
		goto L2210
	}
L2210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14636)+28)) = v14568
	v14639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14576)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14636)+26)) = uint8(v14639)
	v14641 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v14636)+156)) = v14641
	v14643 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v14636)+160)) = v14643
	v14645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14576)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14636)+164)) = uint8(v14645)
	v14647 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v14636)+168)) = v14647
	v14649 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+40))
	v14650 = *(*int32)(unsafe.Add(mBase, uint32(v14649)+12))
	v14651 = *(*int32)(unsafe.Add(mBase, uint32(v14650)))
	v14652 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+260))
	v14653 = *(*int32)(unsafe.Add(mBase, uint32(v14632)+76))
	v14654 = F_get_sortgrouplist_exprs(m, v14652, v14653)
	mBase = m.M
	v14655 = m.ExcPending
	if v14655 != 0 {
		goto L1
	} else {
		goto L2211
	}
L2211:
	;
	v14656 = *(*float64)(unsafe.Add(mBase, uint32(v14651)+32))
	v14657 = int32(0)
	v14659 = F_estimate_num_groups(m, v14572, v14654, v14656, v14657, v14657)
	mBase = m.M
	v14660 = m.ExcPending
	if v14660 != 0 {
		goto L1
	} else {
		goto L2212
	}
L2212:
	;
	v14661 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+260))
	if v14661 == int32(0) {
		goto L2215
	} else {
		goto L2216
	}
L2213:
	;
	v15049 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[8])))
	if v15049 != int32(1) {
		goto L2302
	} else {
		goto L2303
	}
L2214:
	;
	if v14706 == int32(0) {
		goto L2213
	} else {
		goto L2227
	}
L2215:
	;
	v14706 = int32(1)
	goto L2214
L2216:
	;
	goto L2217
L2217:
	;
	v14670 = *(*int32)(unsafe.Add(mBase, uint32(v14661)+4))
	if v14670 <= int32(0) {
		v14698 = int32(1)
		goto L2218
	} else {
		goto L2219
	}
L2218:
	;
	v14706 = v14698
	goto L2214
L2219:
	;
	v14673 = int32(0)
	if v14673 < v14670 {
		goto L2220
	} else {
		goto L2221
	}
L2220:
	;
	v14676 = v14670
	goto L2222
L2221:
	;
	v14676 = v14673
	goto L2222
L2222:
	;
	v14677 = *(*int32)(unsafe.Add(mBase, uint32(v14661)+12))
	v14679 = int32(0)
	goto L2223
L2223:
	;
	v14687 = *(*int32)(unsafe.Add(mBase, uint32(v14677+v14679<<(uint(int32(2))%32))))
	v14688 = *(*int32)(unsafe.Add(mBase, uint32(v14687)+12))
	v14689 = int32(0)
	v14690 = base.B2i32(v14688 != v14689)
	if v14688 == v14689 {
		v14698 = v14690
		goto L2218
	} else {
		goto L2225
	}
L2224:
	;
	v14698 = v14690
	goto L2218
L2225:
	;
	v14694 = v14679 + int32(1)
	if v14694 != v14676 {
		v14679 = v14694
		goto L2223
	} else {
		goto L2226
	}
L2226:
	;
	goto L2224
L2227:
	;
	v14709 = *(*int32)(unsafe.Add(mBase, uint32(v14576)+40))
	if v14709 == int32(0) {
		goto L2213
	} else {
		goto L2228
	}
L2228:
	;
	v14712 = *(*int32)(unsafe.Add(mBase, uint32(v14709)+4))
	if v14712 <= int32(0) {
		goto L2213
	} else {
		goto L2229
	}
L2229:
	;
	v14729 = int32(0)
	goto L2230
L2230:
	;
	v14757 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+172))
	v14758 = *(*int32)(unsafe.Add(mBase, uint32(v14709)+12))
	v14762 = *(*int32)(unsafe.Add(mBase, uint32(v14758+v14729<<(uint(int32(2))%32))))
	v14763 = *(*int32)(unsafe.Add(mBase, uint32(v14762)+64))
	v14764 = F_get_useful_pathkeys_for_distinct(m, v14572, v14757, v14763)
	mBase = m.M
	v14765 = m.ExcPending
	if v14765 != 0 {
		goto L1
	} else {
		goto L2233
	}
L2231:
	;
	goto L2213
L2232:
	;
	v15004 = v14729 + int32(1)
	v15005 = *(*int32)(unsafe.Add(mBase, uint32(v14709)+4))
	if v15004 < v15005 {
		v14729 = v15004
		goto L2230
	} else {
		goto L2301
	}
L2233:
	;
	if v14764 == int32(0) {
		goto L2232
	} else {
		goto L2234
	}
L2234:
	;
	v14768 = int32(0)
	v14769 = *(*int32)(unsafe.Add(mBase, uint32(v14764)+4))
	if v14769 <= v14768 {
		goto L2232
	} else {
		goto L2235
	}
L2235:
	;
	v14779 = v14768
	goto L2236
L2236:
	;
	v14813 = *(*int32)(unsafe.Add(mBase, uint32(v14764)+12))
	v14817 = *(*int32)(unsafe.Add(mBase, uint32(v14813+v14779<<(uint(int32(2))%32))))
	v14818 = *(*int32)(unsafe.Add(mBase, uint32(v14762)+64))
	v14820 = v14577 + int32(192)
	if v14817 == v14818 {
		goto L2242
	} else {
		goto L2243
	}
L2237:
	;
	goto L2232
L2238:
	;
	v14959 = v14779 + int32(1)
	v14960 = *(*int32)(unsafe.Add(mBase, uint32(v14764)+4))
	if v14959 < v14960 {
		v14779 = v14959
		goto L2236
	} else {
		goto L2300
	}
L2239:
	;
	v14928 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+172))
	if v14928 == int32(0) {
		goto L2291
	} else {
		goto L2292
	}
L2240:
	;
	if v14898 != 0 {
		goto L2272
	} else {
		goto L2273
	}
L2241:
	;
	v14886 = *(*int32)(unsafe.Add(mBase, uint32(v14817)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14886
	v14898 = int32(1)
	goto L2240
L2242:
	;
	if v14817 != 0 {
		goto L2241
	} else {
		goto L2245
	}
L2243:
	;
	goto L2244
L2244:
	;
	if v14817 == int32(0) {
		goto L2246
	} else {
		goto L2247
	}
L2245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = int32(0)
	v14898 = int32(1)
	goto L2240
L2246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = int32(0)
	v14898 = int32(1)
	goto L2240
L2247:
	;
	goto L2248
L2248:
	;
	if v14818 == int32(0) {
		goto L2249
	} else {
		goto L2250
	}
L2249:
	;
	v14838 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14838
	v14898 = v14838
	goto L2240
L2250:
	;
	goto L2251
L2251:
	;
	v14841 = *(*int32)(unsafe.Add(mBase, uint32(v14818)+4))
	v14842 = int32(0)
	if v14842 < v14841 {
		goto L2252
	} else {
		goto L2253
	}
L2252:
	;
	v14845 = v14841
	goto L2254
L2253:
	;
	v14845 = v14842
	goto L2254
L2254:
	;
	v14846 = *(*int32)(unsafe.Add(mBase, uint32(v14817)+4))
	v14851 = int32(0)
	goto L2255
L2255:
	;
	if v14851 < v14846 {
		goto L2257
	} else {
		goto L2258
	}
L2257:
	;
	v14858 = *(*int32)(unsafe.Add(mBase, uint32(v14817)+12))
	v14862 = v14858 + v14851<<(uint(int32(2))%32)
	goto L2259
L2258:
	;
	v14862 = int32(0)
	goto L2259
L2259:
	;
	if v14851 == v14845 {
		goto L2260
	} else {
		goto L2261
	}
L2260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14845
	v14898 = base.B2i32(v14862 == int32(0))
	goto L2240
L2261:
	;
	goto L2262
L2262:
	;
	v14868 = base.B2i32(v14862 == int32(0))
	if v14862 == int32(0) {
		goto L2263
	} else {
		goto L2264
	}
L2263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14851
	v14898 = v14868
	goto L2240
L2264:
	;
	goto L2265
L2265:
	;
	v14872 = *(*int32)(unsafe.Add(mBase, uint32(v14818)+12))
	if v14872 == int32(0) {
		goto L2266
	} else {
		goto L2267
	}
L2266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14851
	v14898 = v14868
	goto L2240
L2267:
	;
	goto L2268
L2268:
	;
	v14876 = *(*int32)(unsafe.Add(mBase, uint32(v14862)))
	v14880 = *(*int32)(unsafe.Add(mBase, uint32(v14872+v14851<<(uint(int32(2))%32))))
	if v14876 != v14880 {
		goto L2269
	} else {
		goto L2270
	}
L2269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14820))) = v14851
	v14898 = int32(0)
	goto L2240
L2270:
	;
	v14851 = v14851 + int32(1)
	goto L2255
L2272:
	;
	v14927 = v14762
	goto L2239
L2273:
	;
	goto L2274
L2274:
	;
	v14900 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v14901 = *(*int32)(unsafe.Add(mBase, uint32(v14577)+192))
	if v14762 == v14651 {
		goto L2276
	} else {
		goto L2277
	}
L2275:
	;
	if v14910&int32(1) != 0 {
		goto L2281
	} else {
		goto L2282
	}
L2276:
	;
	v14910 = v14900
	goto L2275
L2277:
	;
	goto L2278
L2278:
	;
	if v14901 == int32(0) {
		goto L2238
	} else {
		goto L2279
	}
L2279:
	;
	v14905 = int32(1)
	if v14900&v14905 == int32(0) {
		goto L2238
	} else {
		goto L2280
	}
L2280:
	;
	v14910 = v14905
	goto L2275
L2281:
	;
	v14914 = v14901
	goto L2283
L2282:
	;
	v14914 = int32(0)
	goto L2283
L2283:
	;
	if v14914 == int32(0) {
		goto L2284
	} else {
		goto L2285
	}
L2284:
	;
	v14918 = F_create_sort_path(m, v14636, v14762, v14817, float64(-1))
	mBase = m.M
	v14919 = m.ExcPending
	if v14919 != 0 {
		goto L1
	} else {
		goto L2287
	}
L2285:
	;
	goto L2286
L2286:
	;
	v14921 = F_create_incremental_sort_path(m, v14572, v14636, v14762, v14817, v14901, float64(-1))
	mBase = m.M
	v14922 = m.ExcPending
	if v14922 != 0 {
		goto L1
	} else {
		goto L2289
	}
L2287:
	;
	if v14918 != 0 {
		v14927 = v14918
		goto L2239
	} else {
		goto L2288
	}
L2288:
	;
	goto L2238
L2289:
	;
	if v14921 == int32(0) {
		goto L2238
	} else {
		goto L2290
	}
L2290:
	;
	v14927 = v14921
	goto L2239
L2291:
	;
	v14931 = int32(0)
	v14937 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v14938 = m.ExcPending
	if v14938 != 0 {
		goto L1
	} else {
		goto L2294
	}
L2292:
	;
	goto L2293
L2293:
	;
	v14950 = *(*int32)(unsafe.Add(mBase, uint32(v14928)+4))
	v14951 = F_create_upper_unique_path(m, v14636, v14927, v14950, v14659)
	mBase = m.M
	v14952 = m.ExcPending
	if v14952 != 0 {
		goto L1
	} else {
		goto L2298
	}
L2294:
	;
	v14939 = int32(0)
	v14941 = F_makeConst(m, int32(20), int32(-1), v14931, int32(8), v14937, v14939, v14939)
	mBase = m.M
	v14942 = m.ExcPending
	if v14942 != 0 {
		goto L1
	} else {
		goto L2295
	}
L2295:
	;
	v14946 = F_create_limit_path(m, v14636, v14927, v14931, v14941, int32(0), int64(0), int64(1))
	mBase = m.M
	v14947 = m.ExcPending
	if v14947 != 0 {
		goto L1
	} else {
		goto L2296
	}
L2296:
	;
	F_add_partial_path(m, v14636, v14946)
	mBase = m.M
	v14949 = m.ExcPending
	if v14949 != 0 {
		goto L1
	} else {
		goto L2297
	}
L2297:
	;
	goto L2238
L2298:
	;
	F_add_partial_path(m, v14636, v14951)
	mBase = m.M
	v14954 = m.ExcPending
	if v14954 != 0 {
		goto L1
	} else {
		goto L2299
	}
L2299:
	;
	goto L2238
L2300:
	;
	goto L2237
L2301:
	;
	goto L2231
L2302:
	;
	v15103 = *(*int32)(unsafe.Add(mBase, uint32(v14636)+168))
	if v15103 == int32(0) {
		goto L2320
	} else {
		goto L2321
	}
L2303:
	;
	v15052 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+260))
	v15053 = int32(0)
	if v15052 == v15053 {
		goto L2305
	} else {
		goto L2306
	}
L2304:
	;
	if v15090 == int32(0) {
		goto L2302
	} else {
		goto L2317
	}
L2305:
	;
	v15090 = int32(1)
	goto L2304
L2306:
	;
	goto L2307
L2307:
	;
	v15060 = *(*int32)(unsafe.Add(mBase, uint32(v15052)+4))
	if v15060 <= int32(0) {
		v15084 = int32(1)
		goto L2308
	} else {
		goto L2309
	}
L2308:
	;
	v15090 = v15084
	goto L2304
L2309:
	;
	v15063 = int32(0)
	if v15063 < v15060 {
		goto L2310
	} else {
		goto L2311
	}
L2310:
	;
	v15066 = v15060
	goto L2312
L2311:
	;
	v15066 = v15063
	goto L2312
L2312:
	;
	v15067 = *(*int32)(unsafe.Add(mBase, uint32(v15052)+12))
	v15071 = v15053
	goto L2313
L2313:
	;
	v15075 = *(*int32)(unsafe.Add(mBase, uint32(v15067+v15071<<(uint(int32(2))%32))))
	v15076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15075)+18)))
	if v15076 != int32(1) {
		v15084 = v15076
		goto L2308
	} else {
		goto L2315
	}
L2314:
	;
	v15084 = v15076
	goto L2308
L2315:
	;
	v15080 = v15071 + int32(1)
	if v15080 != v15066 {
		v15071 = v15080
		goto L2313
	} else {
		goto L2316
	}
L2316:
	;
	goto L2314
L2317:
	;
	v15093 = *(*int32)(unsafe.Add(mBase, uint32(v14651)+12))
	v15095 = int32(0)
	v15096 = *(*int32)(unsafe.Add(mBase, uint32(v14572)+260))
	v15099 = F_create_agg_path(m, v14572, v14636, v14651, v15093, int32(2), v15095, v15096, v15095, v15095, v14659)
	mBase = m.M
	v15100 = m.ExcPending
	if v15100 != 0 {
		goto L1
	} else {
		goto L2318
	}
L2318:
	;
	F_add_partial_path(m, v14636, v15099)
	mBase = m.M
	v15102 = m.ExcPending
	if v15102 != 0 {
		goto L1
	} else {
		goto L2319
	}
L2319:
	;
	goto L2302
L2320:
	;
	v15115 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v15115 != 0 {
		goto L2324
	} else {
		goto L2325
	}
L2321:
	;
	v15106 = *(*int32)(unsafe.Add(mBase, uint32(v15103)+36))
	if v15106 == int32(0) {
		goto L2320
	} else {
		goto L2322
	}
L2322:
	;
	m.T0[v15106].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14572, int32(4), v14576, v14636, int32(0))
	mBase = m.M
	v15112 = m.ExcPending
	if v15112 != 0 {
		goto L1
	} else {
		goto L2323
	}
L2323:
	;
	goto L2320
L2324:
	;
	m.T0[v15115].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14572, int32(4), v14576, v14636, int32(0))
	mBase = m.M
	v15119 = m.ExcPending
	if v15119 != 0 {
		goto L1
	} else {
		goto L2327
	}
L2325:
	;
	goto L2326
L2326:
	;
	v15120 = *(*int32)(unsafe.Add(mBase, uint32(v14636)+40))
	if v15120 == int32(0) {
		goto L2206
	} else {
		goto L2328
	}
L2327:
	;
	goto L2326
L2328:
	;
	F_generate_useful_gather_paths(m, v14572, v14636, int32(1))
	mBase = m.M
	v15125 = m.ExcPending
	if v15125 != 0 {
		goto L1
	} else {
		goto L2329
	}
L2329:
	;
	F_set_cheapest(m, v14636)
	mBase = m.M
	v15127 = m.ExcPending
	if v15127 != 0 {
		goto L1
	} else {
		goto L2330
	}
L2330:
	;
	v15128 = F_create_final_distinct_paths(m, v14572, v14636, v14624)
	mBase = m.M
	v15129 = m.ExcPending
	if v15129 != 0 {
		goto L1
	} else {
		goto L2331
	}
L2331:
	;
	goto L2206
L2332:
	;
	v15174 = *(*int32)(unsafe.Add(mBase, uint32(v14612)+168))
	if v15174 == int32(0) {
		goto L2333
	} else {
		goto L2334
	}
L2333:
	;
	v15186 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v15186 != 0 {
		goto L2337
	} else {
		goto L2338
	}
L2334:
	;
	v15177 = *(*int32)(unsafe.Add(mBase, uint32(v15174)+36))
	if v15177 == int32(0) {
		goto L2333
	} else {
		goto L2335
	}
L2335:
	;
	m.T0[v15177].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14572, int32(5), v14576, v14624, int32(0))
	mBase = m.M
	v15183 = m.ExcPending
	if v15183 != 0 {
		goto L1
	} else {
		goto L2336
	}
L2336:
	;
	goto L2333
L2337:
	;
	m.T0[v15186].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14572, int32(5), v14576, v14624, int32(0))
	mBase = m.M
	v15190 = m.ExcPending
	if v15190 != 0 {
		goto L1
	} else {
		goto L2340
	}
L2338:
	;
	goto L2339
L2339:
	;
	F_set_cheapest(m, v14624)
	mBase = m.M
	v15192 = m.ExcPending
	if v15192 != 0 {
		goto L1
	} else {
		goto L2341
	}
L2340:
	;
	goto L2339
L2341:
	;
	v15197 = v14570
	v15199 = v14572
	v15204 = v14577
	v15205 = v14624
	v15208 = v14581
	v15209 = v14582
	v15217 = v14590
	v15218 = v14591
	v15227 = v14600
	v15231 = v14604
	v15232 = v14605
	goto L713
L2342:
	;
	v15812 = F_fetch_upper_rel(m, v15199, int32(7), int32(0))
	mBase = m.M
	v15813 = m.ExcPending
	if v15813 != 0 {
		goto L1
	} else {
		goto L2495
	}
L2343:
	;
	v15779 = v15205
	goto L2342
L2344:
	;
	goto L2345
L2345:
	;
	v15237 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+48))
	v15240 = F_fetch_upper_rel(m, v15199, int32(6), int32(0))
	mBase = m.M
	v15241 = m.ExcPending
	if v15241 != 0 {
		goto L1
	} else {
		goto L2346
	}
L2346:
	;
	v15242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15205)+26)))
	if v15217&v15242 == int32(1) {
		goto L2347
	} else {
		goto L2348
	}
L2347:
	;
	v15246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15240)+26)) = uint8(v15246)
	goto L2349
L2348:
	;
	goto L2349
L2349:
	;
	v15248 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v15240)+156)) = v15248
	v15250 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v15240)+160)) = v15250
	v15252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15205)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15240)+164)) = uint8(v15252)
	v15254 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v15240)+168)) = v15254
	v15256 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+32))
	if v15256 == int32(0) {
		goto L2350
	} else {
		goto L2351
	}
L2350:
	;
	v15477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15240)+26)))
	if v15477 == int32(0) {
		goto L2411
	} else {
		goto L2412
	}
L2351:
	;
	v15259 = *(*int32)(unsafe.Add(mBase, uint32(v15256)+4))
	if v15259 <= int32(0) {
		goto L2350
	} else {
		goto L2352
	}
L2352:
	;
	v15271 = int32(0)
	goto L2353
L2353:
	;
	v15304 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	v15305 = *(*int32)(unsafe.Add(mBase, uint32(v15256)+12))
	v15309 = *(*int32)(unsafe.Add(mBase, uint32(v15305+v15271<<(uint(int32(2))%32))))
	v15310 = *(*int32)(unsafe.Add(mBase, uint32(v15309)+64))
	v15312 = v15204 + int32(192)
	if v15304 == v15310 {
		goto L2359
	} else {
		goto L2360
	}
L2354:
	;
	goto L2350
L2355:
	;
	v15433 = v15271 + int32(1)
	v15434 = *(*int32)(unsafe.Add(mBase, uint32(v15256)+4))
	if v15433 < v15434 {
		v15271 = v15433
		goto L2353
	} else {
		goto L2410
	}
L2356:
	;
	v15418 = *(*int32)(unsafe.Add(mBase, uint32(v15416)+12))
	v15419 = *(*int32)(unsafe.Add(mBase, uint32(v15418)+4))
	v15420 = *(*int32)(unsafe.Add(mBase, uint32(v15208)+4))
	v15421 = F_equal(m, v15419, v15420)
	mBase = m.M
	v15422 = m.ExcPending
	if v15422 != 0 {
		goto L1
	} else {
		goto L2404
	}
L2357:
	;
	if v15390 != 0 {
		v15416 = v15309
		goto L2356
	} else {
		goto L2389
	}
L2358:
	;
	v15378 = *(*int32)(unsafe.Add(mBase, uint32(v15304)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15378
	v15390 = int32(1)
	goto L2357
L2359:
	;
	if v15304 != 0 {
		goto L2358
	} else {
		goto L2362
	}
L2360:
	;
	goto L2361
L2361:
	;
	if v15304 == int32(0) {
		goto L2363
	} else {
		goto L2364
	}
L2362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = int32(0)
	v15390 = int32(1)
	goto L2357
L2363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = int32(0)
	v15390 = int32(1)
	goto L2357
L2364:
	;
	goto L2365
L2365:
	;
	if v15310 == int32(0) {
		goto L2366
	} else {
		goto L2367
	}
L2366:
	;
	v15330 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15330
	v15390 = v15330
	goto L2357
L2367:
	;
	goto L2368
L2368:
	;
	v15333 = *(*int32)(unsafe.Add(mBase, uint32(v15310)+4))
	v15334 = int32(0)
	if v15334 < v15333 {
		goto L2369
	} else {
		goto L2370
	}
L2369:
	;
	v15337 = v15333
	goto L2371
L2370:
	;
	v15337 = v15334
	goto L2371
L2371:
	;
	v15338 = *(*int32)(unsafe.Add(mBase, uint32(v15304)+4))
	v15343 = int32(0)
	goto L2372
L2372:
	;
	if v15343 < v15338 {
		goto L2374
	} else {
		goto L2375
	}
L2374:
	;
	v15350 = *(*int32)(unsafe.Add(mBase, uint32(v15304)+12))
	v15354 = v15350 + v15343<<(uint(int32(2))%32)
	goto L2376
L2375:
	;
	v15354 = int32(0)
	goto L2376
L2376:
	;
	if v15343 == v15337 {
		goto L2377
	} else {
		goto L2378
	}
L2377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15337
	v15390 = base.B2i32(v15354 == int32(0))
	goto L2357
L2378:
	;
	goto L2379
L2379:
	;
	v15360 = base.B2i32(v15354 == int32(0))
	if v15354 == int32(0) {
		goto L2380
	} else {
		goto L2381
	}
L2380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15343
	v15390 = v15360
	goto L2357
L2381:
	;
	goto L2382
L2382:
	;
	v15364 = *(*int32)(unsafe.Add(mBase, uint32(v15310)+12))
	if v15364 == int32(0) {
		goto L2383
	} else {
		goto L2384
	}
L2383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15343
	v15390 = v15360
	goto L2357
L2384:
	;
	goto L2385
L2385:
	;
	v15368 = *(*int32)(unsafe.Add(mBase, uint32(v15354)))
	v15372 = *(*int32)(unsafe.Add(mBase, uint32(v15364+v15343<<(uint(int32(2))%32))))
	if v15368 != v15372 {
		goto L2386
	} else {
		goto L2387
	}
L2386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15312))) = v15343
	v15390 = int32(0)
	goto L2357
L2387:
	;
	v15343 = v15343 + int32(1)
	goto L2372
L2389:
	;
	v15392 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v15393 = *(*int32)(unsafe.Add(mBase, uint32(v15204)+192))
	if v15237 == v15309 {
		goto L2391
	} else {
		goto L2392
	}
L2390:
	;
	v15403 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	if v15402&int32(1) != 0 {
		goto L2396
	} else {
		goto L2397
	}
L2391:
	;
	v15402 = v15392
	goto L2390
L2392:
	;
	goto L2393
L2393:
	;
	if v15393 == int32(0) {
		goto L2355
	} else {
		goto L2394
	}
L2394:
	;
	v15397 = int32(1)
	if v15392&v15397 == int32(0) {
		goto L2355
	} else {
		goto L2395
	}
L2395:
	;
	v15402 = v15397
	goto L2390
L2396:
	;
	v15407 = v15393
	goto L2398
L2397:
	;
	v15407 = int32(0)
	goto L2398
L2398:
	;
	if v15407 == int32(0) {
		goto L2399
	} else {
		goto L2400
	}
L2399:
	;
	v15410 = F_create_sort_path(m, v15240, v15309, v15403, v15197)
	mBase = m.M
	v15411 = m.ExcPending
	if v15411 != 0 {
		goto L1
	} else {
		goto L2402
	}
L2400:
	;
	goto L2401
L2401:
	;
	v15412 = F_create_incremental_sort_path(m, v15199, v15240, v15309, v15403, v15393, v15197)
	mBase = m.M
	v15413 = m.ExcPending
	if v15413 != 0 {
		goto L1
	} else {
		goto L2403
	}
L2402:
	;
	v15416 = v15410
	goto L2356
L2403:
	;
	v15416 = v15412
	goto L2356
L2404:
	;
	if v15421 != 0 {
		goto L2405
	} else {
		goto L2406
	}
L2405:
	;
	v15425 = v15416
	goto L2407
L2406:
	;
	v15423 = F_apply_projection_to_path(m, v15199, v15240, v15416, v15208)
	mBase = m.M
	v15424 = m.ExcPending
	if v15424 != 0 {
		goto L1
	} else {
		goto L2408
	}
L2407:
	;
	F_add_path(m, v15240, v15425)
	mBase = m.M
	v15427 = m.ExcPending
	if v15427 != 0 {
		goto L1
	} else {
		goto L2409
	}
L2408:
	;
	v15425 = v15423
	goto L2407
L2409:
	;
	goto L2355
L2410:
	;
	goto L2354
L2411:
	;
	v15745 = *(*int32)(unsafe.Add(mBase, uint32(v15240)+168))
	if v15745 == int32(0) {
		goto L2485
	} else {
		goto L2486
	}
L2412:
	;
	v15480 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	if v15480 == int32(0) {
		goto L2411
	} else {
		goto L2413
	}
L2413:
	;
	v15483 = *(*int32)(unsafe.Add(mBase, uint32(v15205)+40))
	if v15483 == int32(0) {
		goto L2411
	} else {
		goto L2414
	}
L2414:
	;
	v15486 = *(*int32)(unsafe.Add(mBase, uint32(v15483)+4))
	if v15486 <= int32(0) {
		goto L2411
	} else {
		goto L2415
	}
L2415:
	;
	v15489 = *(*int32)(unsafe.Add(mBase, uint32(v15483)+12))
	v15490 = *(*int32)(unsafe.Add(mBase, uint32(v15489)))
	v15499 = int32(0)
	goto L2416
L2416:
	;
	v15533 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	v15534 = *(*int32)(unsafe.Add(mBase, uint32(v15483)+12))
	v15538 = *(*int32)(unsafe.Add(mBase, uint32(v15534+v15499<<(uint(int32(2))%32))))
	v15539 = *(*int32)(unsafe.Add(mBase, uint32(v15538)+64))
	v15541 = v15204 + int32(304)
	if v15533 == v15539 {
		goto L2421
	} else {
		goto L2422
	}
L2417:
	;
	goto L2411
L2418:
	;
	v15701 = v15499 + int32(1)
	v15702 = *(*int32)(unsafe.Add(mBase, uint32(v15483)+4))
	if v15701 < v15702 {
		v15499 = v15701
		goto L2416
	} else {
		goto L2484
	}
L2419:
	;
	if v15619 != 0 {
		goto L2418
	} else {
		goto L2451
	}
L2420:
	;
	v15607 = *(*int32)(unsafe.Add(mBase, uint32(v15533)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15607
	v15619 = int32(1)
	goto L2419
L2421:
	;
	if v15533 != 0 {
		goto L2420
	} else {
		goto L2424
	}
L2422:
	;
	goto L2423
L2423:
	;
	if v15533 == int32(0) {
		goto L2425
	} else {
		goto L2426
	}
L2424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = int32(0)
	v15619 = int32(1)
	goto L2419
L2425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = int32(0)
	v15619 = int32(1)
	goto L2419
L2426:
	;
	goto L2427
L2427:
	;
	if v15539 == int32(0) {
		goto L2428
	} else {
		goto L2429
	}
L2428:
	;
	v15559 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15559
	v15619 = v15559
	goto L2419
L2429:
	;
	goto L2430
L2430:
	;
	v15562 = *(*int32)(unsafe.Add(mBase, uint32(v15539)+4))
	v15563 = int32(0)
	if v15563 < v15562 {
		goto L2431
	} else {
		goto L2432
	}
L2431:
	;
	v15566 = v15562
	goto L2433
L2432:
	;
	v15566 = v15563
	goto L2433
L2433:
	;
	v15567 = *(*int32)(unsafe.Add(mBase, uint32(v15533)+4))
	v15572 = int32(0)
	goto L2434
L2434:
	;
	if v15572 < v15567 {
		goto L2436
	} else {
		goto L2437
	}
L2436:
	;
	v15579 = *(*int32)(unsafe.Add(mBase, uint32(v15533)+12))
	v15583 = v15579 + v15572<<(uint(int32(2))%32)
	goto L2438
L2437:
	;
	v15583 = int32(0)
	goto L2438
L2438:
	;
	if v15572 == v15566 {
		goto L2439
	} else {
		goto L2440
	}
L2439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15566
	v15619 = base.B2i32(v15583 == int32(0))
	goto L2419
L2440:
	;
	goto L2441
L2441:
	;
	v15589 = base.B2i32(v15583 == int32(0))
	if v15583 == int32(0) {
		goto L2442
	} else {
		goto L2443
	}
L2442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15572
	v15619 = v15589
	goto L2419
L2443:
	;
	goto L2444
L2444:
	;
	v15593 = *(*int32)(unsafe.Add(mBase, uint32(v15539)+12))
	if v15593 == int32(0) {
		goto L2445
	} else {
		goto L2446
	}
L2445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15572
	v15619 = v15589
	goto L2419
L2446:
	;
	goto L2447
L2447:
	;
	v15597 = *(*int32)(unsafe.Add(mBase, uint32(v15583)))
	v15601 = *(*int32)(unsafe.Add(mBase, uint32(v15593+v15572<<(uint(int32(2))%32))))
	if v15597 != v15601 {
		goto L2448
	} else {
		goto L2449
	}
L2448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15541))) = v15572
	v15619 = int32(0)
	goto L2419
L2449:
	;
	v15572 = v15572 + int32(1)
	goto L2434
L2451:
	;
	v15621 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v15622 = *(*int32)(unsafe.Add(mBase, uint32(v15204)+304))
	if v15538 == v15490 {
		goto L2453
	} else {
		goto L2454
	}
L2452:
	;
	v15632 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	if v15631&int32(1) != 0 {
		goto L2459
	} else {
		goto L2460
	}
L2453:
	;
	v15631 = v15621
	goto L2452
L2454:
	;
	goto L2455
L2455:
	;
	if v15622 == int32(0) {
		goto L2418
	} else {
		goto L2456
	}
L2456:
	;
	v15626 = int32(1)
	if v15621&v15626 == int32(0) {
		goto L2418
	} else {
		goto L2457
	}
L2457:
	;
	v15631 = v15626
	goto L2452
L2458:
	;
	v15647 = *(*float64)(unsafe.Add(mBase, uint32(v15643)+32))
	v15648 = *(*int32)(unsafe.Add(mBase, uint32(v15643)+24))
	v15649 = base.F64_convert_i32_s(v15648)
	v15651 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[9])))
	if v15651 == int32(1) {
		goto L2468
	} else {
		goto L2469
	}
L2459:
	;
	v15636 = v15622
	goto L2461
L2460:
	;
	v15636 = int32(0)
	goto L2461
L2461:
	;
	if v15636 == int32(0) {
		goto L2462
	} else {
		goto L2463
	}
L2462:
	;
	v15639 = F_create_sort_path(m, v15240, v15538, v15632, v15197)
	mBase = m.M
	v15640 = m.ExcPending
	if v15640 != 0 {
		goto L1
	} else {
		goto L2465
	}
L2463:
	;
	goto L2464
L2464:
	;
	v15641 = F_create_incremental_sort_path(m, v15199, v15240, v15538, v15632, v15622, v15197)
	mBase = m.M
	v15642 = m.ExcPending
	if v15642 != 0 {
		goto L1
	} else {
		goto L2466
	}
L2465:
	;
	v15643 = v15639
	goto L2458
L2466:
	;
	v15643 = v15641
	goto L2458
L2467:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v15204)+192)) = v15679
	v15681 = *(*int32)(unsafe.Add(mBase, uint32(v15643)+12))
	v15682 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+176))
	v15685 = F_create_gather_merge_path(m, v15199, v15240, v15643, v15681, v15682, v15204+int32(192))
	mBase = m.M
	v15686 = m.ExcPending
	if v15686 != 0 {
		goto L1
	} else {
		goto L2477
	}
L2468:
	;
	v15657 = base.F64_add(base.F64_mul(v15649, float64(-0.3)), float64(1))
	if base.F64_gt(v15657, float64(0)) != 0 {
		goto L2471
	} else {
		goto L2472
	}
L2469:
	;
	v15663 = v15649
	goto L2470
L2470:
	;
	v15665 = float64(1e+100)
	v15666 = base.F64_mul(v15647, v15663)
	if base.F64_gt(v15666, v15665)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v15666)&int64(9223372036854775807))) != 0 {
		v15679 = v15665
		goto L2474
	} else {
		goto L2475
	}
L2471:
	;
	v15661 = v15657
	goto L2473
L2472:
	;
	v15661 = math.Float64frombits(uint64(0x8000000000000000))
	goto L2473
L2473:
	;
	v15663 = base.F64_add(v15661, v15649)
	goto L2470
L2474:
	;
	goto L2467
L2475:
	;
	v15675 = float64(1)
	if base.F64_le(v15666, v15675) != 0 {
		v15679 = v15675
		goto L2474
	} else {
		goto L2476
	}
L2476:
	;
	v15679 = base.F64_nearest(v15666)
	goto L2474
L2477:
	;
	v15687 = *(*int32)(unsafe.Add(mBase, uint32(v15685)+12))
	v15688 = *(*int32)(unsafe.Add(mBase, uint32(v15687)+4))
	v15689 = *(*int32)(unsafe.Add(mBase, uint32(v15208)+4))
	v15690 = F_equal(m, v15688, v15689)
	mBase = m.M
	v15691 = m.ExcPending
	if v15691 != 0 {
		goto L1
	} else {
		goto L2478
	}
L2478:
	;
	if v15690 != 0 {
		goto L2479
	} else {
		goto L2480
	}
L2479:
	;
	v15694 = v15685
	goto L2481
L2480:
	;
	v15692 = F_apply_projection_to_path(m, v15199, v15240, v15685, v15208)
	mBase = m.M
	v15693 = m.ExcPending
	if v15693 != 0 {
		goto L1
	} else {
		goto L2482
	}
L2481:
	;
	F_add_path(m, v15240, v15694)
	mBase = m.M
	v15696 = m.ExcPending
	if v15696 != 0 {
		goto L1
	} else {
		goto L2483
	}
L2482:
	;
	v15694 = v15692
	goto L2481
L2483:
	;
	goto L2418
L2484:
	;
	goto L2417
L2485:
	;
	v15757 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v15757 != 0 {
		goto L2489
	} else {
		goto L2490
	}
L2486:
	;
	v15748 = *(*int32)(unsafe.Add(mBase, uint32(v15745)+36))
	if v15748 == int32(0) {
		goto L2485
	} else {
		goto L2487
	}
L2487:
	;
	m.T0[v15748].(func(*base.Module, int32, int32, int32, int32, int32))(m, v15199, int32(6), v15205, v15240, int32(0))
	mBase = m.M
	v15754 = m.ExcPending
	if v15754 != 0 {
		goto L1
	} else {
		goto L2488
	}
L2488:
	;
	goto L2485
L2489:
	;
	m.T0[v15757].(func(*base.Module, int32, int32, int32, int32, int32))(m, v15199, int32(6), v15205, v15240, int32(0))
	mBase = m.M
	v15761 = m.ExcPending
	if v15761 != 0 {
		goto L1
	} else {
		goto L2492
	}
L2490:
	;
	goto L2491
L2491:
	;
	v15762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15209)+38)))
	if v15762 != int32(1) {
		v15779 = v15240
		goto L2342
	} else {
		goto L2493
	}
L2492:
	;
	goto L2491
L2493:
	;
	v15765 = *(*int32)(unsafe.Add(mBase, uint32(v15204)+188))
	v15766 = *(*int32)(unsafe.Add(mBase, uint32(v15204)+184))
	F_adjust_paths_for_srfs(m, v15199, v15240, v15765, v15766)
	mBase = m.M
	v15768 = m.ExcPending
	if v15768 != 0 {
		goto L1
	} else {
		goto L2494
	}
L2494:
	;
	v15779 = v15240
	goto L2342
L2495:
	;
	v15814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15779)+26)))
	if v15814 != int32(1) {
		goto L2496
	} else {
		goto L2497
	}
L2496:
	;
	v15829 = *(*int32)(unsafe.Add(mBase, uint32(v15779)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v15812)+156)) = v15829
	v15831 = *(*int32)(unsafe.Add(mBase, uint32(v15779)+160))
	*(*int32)(unsafe.Add(mBase, uint32(v15812)+160)) = v15831
	v15833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15779)+164)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15812)+164)) = uint8(v15833)
	v15835 = *(*int32)(unsafe.Add(mBase, uint32(v15779)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v15812)+168)) = v15835
	v15837 = *(*int32)(unsafe.Add(mBase, uint32(v15779)+32))
	if v15837 == int32(0) {
		v16795 = v15812
		v16801 = v15199
		v16805 = v15779
		v16806 = v15204
		v16811 = v15209
		v16820 = v15218
		v16829 = v15227
		v16833 = v15231
		v16834 = v15232
		goto L2502
	} else {
		goto L2503
	}
L2497:
	;
	v15817 = *(*int32)(unsafe.Add(mBase, uint32(v15209)+128))
	v15818 = F_is_parallel_safe(m, v15199, v15817)
	mBase = m.M
	v15819 = m.ExcPending
	if v15819 != 0 {
		goto L1
	} else {
		goto L2498
	}
L2498:
	;
	if v15818 == int32(0) {
		goto L2496
	} else {
		goto L2499
	}
L2499:
	;
	v15822 = *(*int32)(unsafe.Add(mBase, uint32(v15209)+132))
	v15823 = F_is_parallel_safe(m, v15199, v15822)
	mBase = m.M
	v15824 = m.ExcPending
	if v15824 != 0 {
		goto L1
	} else {
		goto L2500
	}
L2500:
	;
	if v15823 == int32(0) {
		goto L2496
	} else {
		goto L2501
	}
L2501:
	;
	v15827 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15812)+26)) = uint8(v15827)
	goto L2496
L2502:
	;
	v16836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16795)+26)))
	if v16836 == int32(0) {
		goto L2691
	} else {
		goto L2692
	}
L2503:
	;
	v15840 = *(*int32)(unsafe.Add(mBase, uint32(v15837)+4))
	if v15840 <= int32(0) {
		v16795 = v15812
		v16801 = v15199
		v16805 = v15779
		v16806 = v15204
		v16811 = v15209
		v16820 = v15218
		v16829 = v15227
		v16833 = v15231
		v16834 = v15232
		goto L2502
	} else {
		goto L2504
	}
L2504:
	;
	v15844 = v15812
	v15850 = v15199
	v15853 = v15837
	v15854 = v15779
	v15855 = v15204
	v15860 = v15209
	v15866 = int32(0)
	v15869 = v15218
	v15878 = v15227
	v15882 = v15231
	v15883 = v15232
	goto L2505
L2505:
	;
	v15885 = *(*int32)(unsafe.Add(mBase, uint32(v15853)+12))
	v15889 = *(*int32)(unsafe.Add(mBase, uint32(v15885+v15866<<(uint(int32(2))%32))))
	v15890 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+140))
	if v15890 != 0 {
		goto L2507
	} else {
		goto L2508
	}
L2506:
	;
	v16795 = v16747
	v16801 = v16753
	v16805 = v16757
	v16806 = v16758
	v16811 = v16763
	v16820 = v16772
	v16829 = v16781
	v16833 = v16785
	v16834 = v16786
	goto L2502
L2507:
	;
	v15891 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+136))
	v15892 = F_assign_special_exec_param(m, v15850)
	mBase = m.M
	v15893 = m.ExcPending
	if v15893 != 0 {
		goto L1
	} else {
		goto L2510
	}
L2508:
	;
	v15931 = v15889
	goto L2509
L2509:
	;
	v15932 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+132))
	if v15932 != 0 {
		goto L2514
	} else {
		goto L2515
	}
L2510:
	;
	v15895 = F_palloc0(m, int32(88))
	mBase = m.M
	v15896 = m.ExcPending
	if v15896 != 0 {
		goto L1
	} else {
		goto L2511
	}
L2511:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+8)) = v15844
	*(*int64)(unsafe.Add(mBase, uint32(v15895))) = int64(1597727834427)
	v15900 = *(*int32)(unsafe.Add(mBase, uint32(v15889)+12))
	v15901 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+24)) = v15901
	*(*uint16)(unsafe.Add(mBase, uint32(v15895)+20)) = uint16(v15901)
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+16)) = v15901
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+12)) = v15900
	v15908 = *(*float64)(unsafe.Add(mBase, uint32(v15889)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+80)) = v15892
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+76)) = v15891
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+72)) = v15889
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+64)) = v15901
	*(*float64)(unsafe.Add(mBase, uint32(v15895)+32)) = v15908
	v15915 = *(*int32)(unsafe.Add(mBase, uint32(v15889)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15895)+40)) = v15915
	v15917 = *(*float64)(unsafe.Add(mBase, uint32(v15889)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v15895)+48)) = v15917
	v15920 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	v15921 = *(*float64)(unsafe.Add(mBase, uint32(v15889)+32))
	v15923 = *(*float64)(unsafe.Add(mBase, uint32(v15889)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v15895)+56)) = base.F64_add(base.F64_mul(v15920, v15921), v15923)
	v15931 = v15895
	goto L2509
L2512:
	;
	v15957 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v15957 != int32(1) {
		goto L2524
	} else {
		goto L2525
	}
L2513:
	;
	v15951 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+128))
	v15952 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+136))
	v15953 = F_create_limit_path(m, v15844, v15931, v15951, v15932, v15952, v15882, v15883)
	mBase = m.M
	v15954 = m.ExcPending
	if v15954 != 0 {
		goto L1
	} else {
		goto L2523
	}
L2514:
	;
	v15933 = *(*int32)(unsafe.Add(mBase, uint32(v15932)))
	if v15933 != int32(7) {
		goto L2513
	} else {
		goto L2517
	}
L2515:
	;
	goto L2516
L2516:
	;
	v15939 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+128))
	if v15939 == int32(0) {
		v15956 = v15931
		goto L2512
	} else {
		goto L2519
	}
L2517:
	;
	v15936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15932)+24)))
	if v15936 != int32(1) {
		goto L2513
	} else {
		goto L2518
	}
L2518:
	;
	goto L2516
L2519:
	;
	v15942 = *(*int32)(unsafe.Add(mBase, uint32(v15939)))
	if v15942 != int32(7) {
		goto L2513
	} else {
		goto L2520
	}
L2520:
	;
	v15945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15939)+24)))
	if v15945 != 0 {
		v15956 = v15931
		goto L2512
	} else {
		goto L2521
	}
L2521:
	;
	v15946 = *(*int32)(unsafe.Add(mBase, uint32(v15939)+20))
	v15947 = *(*int64)(unsafe.Add(mBase, uint32(v15946)))
	if v15947 == int64(0) {
		v15956 = v15931
		goto L2512
	} else {
		goto L2522
	}
L2522:
	;
	goto L2513
L2523:
	;
	v15956 = v15953
	goto L2512
L2524:
	;
	v15960 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+120))
	v15961 = int32(0)
	if v15960 == v15961 {
		goto L2528
	} else {
		goto L2529
	}
L2525:
	;
	v16747 = v15844
	v16753 = v15850
	v16756 = v15853
	v16757 = v15854
	v16758 = v15855
	v16763 = v15860
	v16772 = v15869
	v16781 = v15878
	v16785 = v15882
	v16786 = v15883
	v16788 = v15956
	goto L2526
L2526:
	;
	F_add_path(m, v15844, v16788)
	mBase = m.M
	v16790 = m.ExcPending
	if v16790 != 0 {
		goto L1
	} else {
		goto L2688
	}
L2527:
	;
	v16007 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+32))
	if v16006 == int32(2) {
		goto L2545
	} else {
		goto L2546
	}
L2528:
	;
	v16006 = int32(0)
	goto L2527
L2529:
	;
	goto L2530
L2530:
	;
	v15969 = int32(1)
	v15970 = *(*int32)(unsafe.Add(mBase, uint32(v15960)+4))
	if v15970 <= v15969 {
		goto L2531
	} else {
		goto L2532
	}
L2531:
	;
	v15973 = v15969
	goto L2533
L2532:
	;
	v15973 = v15970
	goto L2533
L2533:
	;
	v15977 = int32(0)
	v15979 = v15961
	goto L2534
L2534:
	;
	v15986 = *(*int32)(unsafe.Add(mBase, uint32(v15960+int32(8)+v15977<<(uint(int32(2))%32))))
	if v15986 != 0 {
		goto L2537
	} else {
		goto L2538
	}
L2535:
	;
	v16006 = v15998
	goto L2527
L2536:
	;
	goto L2535
L2537:
	;
	v15987 = int32(2)
	if v15979 != 0 {
		v15998 = v15987
		goto L2536
	} else {
		goto L2540
	}
L2538:
	;
	v15993 = v15979
	goto L2539
L2539:
	;
	v15995 = v15977 + int32(1)
	if v15995 != v15973 {
		v15977 = v15995
		v15979 = v15993
		goto L2534
	} else {
		goto L2542
	}
L2540:
	;
	v15988 = int32(1)
	if base.Ui32(v15988) < base.Ui32(base.I32_popcnt(v15986)) {
		v15998 = v15987
		goto L2536
	} else {
		goto L2541
	}
L2541:
	;
	v15993 = v15988
	goto L2539
L2542:
	;
	v15998 = v15993
	goto L2536
L2543:
	;
	v16690 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	v16691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15860)+24)))
	v16692 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+32))
	v16693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15850)+372)))
	v16694 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+140))
	if v16694 != 0 {
		goto L2679
	} else {
		goto L2680
	}
L2544:
	;
	v16643 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v16642))) = v16643
	*(*int32)(unsafe.Add(mBase, uint32(v15855))) = v16643
	v16647 = F_list_make1_impl(m, int32(1), v15855)
	mBase = m.M
	v16648 = m.ExcPending
	if v16648 != 0 {
		goto L1
	} else {
		goto L2678
	}
L2545:
	;
	v16010 = F_find_base_rel(m, v15850, v16007)
	mBase = m.M
	v16011 = m.ExcPending
	if v16011 != 0 {
		goto L1
	} else {
		goto L2548
	}
L2546:
	;
	goto L2547
L2547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+40)) = v16007
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+116)) = v16007
	v16548 = F_list_make1_impl(m, int32(471), v15855+int32(40))
	mBase = m.M
	v16549 = m.ExcPending
	if v16549 != 0 {
		goto L1
	} else {
		goto L2658
	}
L2548:
	;
	v16012 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+32))
	v16013 = int32(0)
	v16018 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+124))
	if v16018 == v16013 {
		goto L2551
	} else {
		goto L2552
	}
L2549:
	;
	if int32(0) <= v16075 {
		goto L2560
	} else {
		goto L2561
	}
L2550:
	;
	v16075 = base.I32_ctz(v16061) | v16062<<(uint(int32(5))%32)
	goto L2549
L2551:
	;
	v16075 = int32(-2)
	goto L2549
L2552:
	;
	v16028 = base.I32_div_s(int32(0), int32(32))
	v16029 = *(*int32)(unsafe.Add(mBase, uint32(v16018)+4))
	if v16029 <= v16028 {
		goto L2551
	} else {
		goto L2553
	}
L2553:
	;
	v16032 = v16018 + int32(8)
	v16036 = *(*int32)(unsafe.Add(mBase, uint32(v16032+v16028<<(uint(int32(2))%32))))
	v16039 = v16036 & int32(-1)
	if v16039 != 0 {
		v16061 = v16039
		v16062 = v16028
		goto L2550
	} else {
		goto L2554
	}
L2554:
	;
	v16041 = v16028 + int32(1)
	if v16041 == v16029 {
		goto L2551
	} else {
		goto L2555
	}
L2555:
	;
	v16044 = v16041
	goto L2556
L2556:
	;
	v16051 = *(*int32)(unsafe.Add(mBase, uint32(v16032+v16044<<(uint(int32(2))%32))))
	if v16051 != 0 {
		v16061 = v16051
		v16062 = v16044
		goto L2550
	} else {
		goto L2558
	}
L2557:
	;
	goto L2551
L2558:
	;
	v16053 = v16044 + int32(1)
	if v16053 != v16029 {
		v16044 = v16053
		goto L2556
	} else {
		goto L2559
	}
L2559:
	;
	goto L2557
L2560:
	;
	v16081 = int32(0)
	v16084 = v16075
	v16087 = v16013
	v16092 = v16013
	v16097 = v16013
	v16099 = v16013
	v16100 = v16013
	goto L2563
L2561:
	;
	v16457 = v16013
	v16462 = v16013
	v16467 = v16013
	v16469 = v16013
	v16470 = v16013
	goto L2562
L2562:
	;
	v16490 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+20)) = v16490
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+140)) = v16490
	v16496 = F_list_make1_impl(m, int32(471), v15855+int32(20))
	mBase = m.M
	v16497 = m.ExcPending
	if v16497 != 0 {
		goto L1
	} else {
		goto L2640
	}
L2563:
	;
	v16120 = F_find_base_rel(m, v15850, v16084)
	mBase = m.M
	v16121 = m.ExcPending
	if v16121 != 0 {
		goto L1
	} else {
		goto L2566
	}
L2564:
	;
	if v16351 != 0 {
		v16651 = v16351
		v16657 = v16357
		v16662 = v16362
		v16664 = v16012
		v16667 = v16367
		v16669 = v16369
		v16670 = v16370
		goto L2543
	} else {
		goto L2639
	}
L2565:
	;
	v16390 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+124))
	if v16390 == int32(0) {
		goto L2629
	} else {
		goto L2630
	}
L2566:
	;
	v16122 = int32(0)
	v16124 = *(*int32)(unsafe.Add(mBase, uint32(v16120)+32))
	if v16124 == v16122 {
		v16145 = v16122
		goto L2568
	} else {
		goto L2569
	}
L2567:
	;
	if v16145 != 0 {
		v16351 = v16081
		v16357 = v16087
		v16362 = v16092
		v16367 = v16097
		v16369 = v16099
		v16370 = v16100
		goto L2565
	} else {
		goto L2577
	}
L2568:
	;
	goto L2567
L2569:
	;
	v16127 = *(*int32)(unsafe.Add(mBase, uint32(v16124)+12))
	v16128 = v16127
	goto L2570
L2570:
	;
	v16131 = *(*int32)(unsafe.Add(mBase, uint32(v16128)))
	v16132 = *(*int32)(unsafe.Add(mBase, uint32(v16131)))
	if base.Ui32(int32(2)) <= base.Ui32(v16132-int32(301)) {
		goto L2572
	} else {
		goto L2573
	}
L2571:
	;
	v16145 = int32(1)
	goto L2568
L2572:
	;
	if v16132 != int32(290) {
		v16145 = v16122
		goto L2568
	} else {
		goto L2575
	}
L2573:
	;
	v16128 = v16131 + int32(72)
	goto L2570
L2574:
	;
	goto L2571
L2575:
	;
	v16139 = *(*int32)(unsafe.Add(mBase, uint32(v16131)+72))
	if v16139 != 0 {
		v16145 = v16122
		goto L2568
	} else {
		goto L2576
	}
L2576:
	;
	goto L2574
L2577:
	;
	v16146 = F_lappend_int(m, v16081, v16084)
	mBase = m.M
	v16147 = m.ExcPending
	if v16147 != 0 {
		goto L1
	} else {
		goto L2578
	}
L2578:
	;
	v16148 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16148 == int32(2) {
		goto L2579
	} else {
		goto L2580
	}
L2579:
	;
	v16151 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+268))
	if v16010 != v16120 {
		goto L2582
	} else {
		goto L2583
	}
L2580:
	;
	v16161 = v16100
	goto L2581
L2581:
	;
	v16162 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+152))
	if v16162 != 0 {
		goto L2587
	} else {
		goto L2588
	}
L2582:
	;
	v16153 = *(*int32)(unsafe.Add(mBase, uint32(v16120)+68))
	v16154 = *(*int32)(unsafe.Add(mBase, uint32(v16010)+68))
	v16155 = F_adjust_inherited_attnums_multilevel(m, v15850, v16151, v16153, v16154)
	mBase = m.M
	v16156 = m.ExcPending
	if v16156 != 0 {
		goto L1
	} else {
		goto L2585
	}
L2583:
	;
	v16157 = v16151
	goto L2584
L2584:
	;
	v16158 = F_lappend(m, v16100, v16157)
	mBase = m.M
	v16159 = m.ExcPending
	if v16159 != 0 {
		goto L1
	} else {
		goto L2586
	}
L2585:
	;
	v16157 = v16155
	goto L2584
L2586:
	;
	v16161 = v16158
	goto L2581
L2587:
	;
	if v16010 != v16120 {
		goto L2590
	} else {
		goto L2591
	}
L2588:
	;
	v16169 = v16099
	goto L2589
L2589:
	;
	v16170 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+96))
	if v16170 != 0 {
		goto L2595
	} else {
		goto L2596
	}
L2590:
	;
	v16164 = F_adjust_appendrel_attrs_multilevel(m, v15850, v16162, v16120, v16010)
	mBase = m.M
	v16165 = m.ExcPending
	if v16165 != 0 {
		goto L1
	} else {
		goto L2593
	}
L2591:
	;
	v16166 = v16162
	goto L2592
L2592:
	;
	v16167 = F_lappend(m, v16099, v16166)
	mBase = m.M
	v16168 = m.ExcPending
	if v16168 != 0 {
		goto L1
	} else {
		goto L2594
	}
L2593:
	;
	v16166 = v16164
	goto L2592
L2594:
	;
	v16169 = v16167
	goto L2589
L2595:
	;
	if v16010 != v16120 {
		goto L2598
	} else {
		goto L2599
	}
L2596:
	;
	v16177 = v16092
	goto L2597
L2597:
	;
	v16178 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+64))
	if v16178 != 0 {
		goto L2603
	} else {
		goto L2604
	}
L2598:
	;
	v16172 = F_adjust_appendrel_attrs_multilevel(m, v15850, v16170, v16120, v16010)
	mBase = m.M
	v16173 = m.ExcPending
	if v16173 != 0 {
		goto L1
	} else {
		goto L2601
	}
L2599:
	;
	v16174 = v16170
	goto L2600
L2600:
	;
	v16175 = F_lappend(m, v16092, v16174)
	mBase = m.M
	v16176 = m.ExcPending
	if v16176 != 0 {
		goto L1
	} else {
		goto L2602
	}
L2601:
	;
	v16174 = v16172
	goto L2600
L2602:
	;
	v16177 = v16175
	goto L2597
L2603:
	;
	v16179 = int32(0)
	v16180 = *(*int32)(unsafe.Add(mBase, uint32(v16178)+4))
	if v16179 < v16180 {
		goto L2606
	} else {
		goto L2607
	}
L2604:
	;
	v16316 = v16097
	goto L2605
L2605:
	;
	v16339 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16339 != int32(5) {
		v16351 = v16146
		v16357 = v16087
		v16362 = v16177
		v16367 = v16316
		v16369 = v16169
		v16370 = v16161
		goto L2565
	} else {
		goto L2621
	}
L2606:
	;
	v16196 = v16179
	v16201 = int32(0)
	goto L2609
L2607:
	;
	v16267 = v16179
	goto L2608
L2608:
	;
	v16296 = F_lappend(m, v16097, v16267)
	mBase = m.M
	v16297 = m.ExcPending
	if v16297 != 0 {
		goto L1
	} else {
		goto L2620
	}
L2609:
	;
	v16225 = *(*int32)(unsafe.Add(mBase, uint32(v16178)+12))
	v16229 = *(*int32)(unsafe.Add(mBase, uint32(v16225+v16201<<(uint(int32(2))%32))))
	v16230 = F_copyObjectImpl(m, v16229)
	mBase = m.M
	v16231 = m.ExcPending
	if v16231 != 0 {
		goto L1
	} else {
		goto L2611
	}
L2610:
	;
	v16267 = v16249
	goto L2608
L2611:
	;
	v16232 = *(*int32)(unsafe.Add(mBase, uint32(v16229)+16))
	v16233 = F_adjust_appendrel_attrs_multilevel(m, v15850, v16232, v16120, v16010)
	mBase = m.M
	v16234 = m.ExcPending
	if v16234 != 0 {
		goto L1
	} else {
		goto L2612
	}
L2612:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16230)+16)) = v16233
	v16236 = *(*int32)(unsafe.Add(mBase, uint32(v16229)+20))
	v16237 = F_adjust_appendrel_attrs_multilevel(m, v15850, v16236, v16120, v16010)
	mBase = m.M
	v16238 = m.ExcPending
	if v16238 != 0 {
		goto L1
	} else {
		goto L2613
	}
L2613:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16230)+20)) = v16237
	v16240 = *(*int32)(unsafe.Add(mBase, uint32(v16230)+8))
	if v16240 == int32(2) {
		goto L2614
	} else {
		goto L2615
	}
L2614:
	;
	v16243 = *(*int32)(unsafe.Add(mBase, uint32(v16229)+24))
	v16244 = *(*int32)(unsafe.Add(mBase, uint32(v16120)+68))
	v16245 = *(*int32)(unsafe.Add(mBase, uint32(v16010)+68))
	v16246 = F_adjust_inherited_attnums_multilevel(m, v15850, v16243, v16244, v16245)
	mBase = m.M
	v16247 = m.ExcPending
	if v16247 != 0 {
		goto L1
	} else {
		goto L2617
	}
L2615:
	;
	goto L2616
L2616:
	;
	v16249 = F_lappend(m, v16196, v16230)
	mBase = m.M
	v16250 = m.ExcPending
	if v16250 != 0 {
		goto L1
	} else {
		goto L2618
	}
L2617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16230)+24)) = v16246
	goto L2616
L2618:
	;
	v16252 = v16201 + int32(1)
	v16253 = *(*int32)(unsafe.Add(mBase, uint32(v16178)+4))
	if v16252 < v16253 {
		v16196 = v16249
		v16201 = v16252
		goto L2609
	} else {
		goto L2619
	}
L2619:
	;
	goto L2610
L2620:
	;
	v16316 = v16296
	goto L2605
L2621:
	;
	v16342 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+72))
	if v16010 != v16120 {
		goto L2622
	} else {
		goto L2623
	}
L2622:
	;
	v16344 = F_adjust_appendrel_attrs_multilevel(m, v15850, v16342, v16120, v16010)
	mBase = m.M
	v16345 = m.ExcPending
	if v16345 != 0 {
		goto L1
	} else {
		goto L2625
	}
L2623:
	;
	v16346 = v16342
	goto L2624
L2624:
	;
	v16347 = F_lappend(m, v16087, v16346)
	mBase = m.M
	v16348 = m.ExcPending
	if v16348 != 0 {
		goto L1
	} else {
		goto L2626
	}
L2625:
	;
	v16346 = v16344
	goto L2624
L2626:
	;
	v16351 = v16146
	v16357 = v16347
	v16362 = v16177
	v16367 = v16316
	v16369 = v16169
	v16370 = v16161
	goto L2565
L2627:
	;
	if int32(0) <= v16446 {
		v16081 = v16351
		v16084 = v16446
		v16087 = v16357
		v16092 = v16362
		v16097 = v16367
		v16099 = v16369
		v16100 = v16370
		goto L2563
	} else {
		goto L2638
	}
L2628:
	;
	v16446 = base.I32_ctz(v16432) | v16433<<(uint(int32(5))%32)
	goto L2627
L2629:
	;
	v16446 = int32(-2)
	goto L2627
L2630:
	;
	v16397 = v16084 + int32(1)
	v16399 = base.I32_div_s(v16397, int32(32))
	v16400 = *(*int32)(unsafe.Add(mBase, uint32(v16390)+4))
	if v16400 <= v16399 {
		goto L2629
	} else {
		goto L2631
	}
L2631:
	;
	v16403 = v16390 + int32(8)
	v16407 = *(*int32)(unsafe.Add(mBase, uint32(v16403+v16399<<(uint(int32(2))%32))))
	v16410 = v16407 & (int32(-1) << (uint(v16397) % 32))
	if v16410 != 0 {
		v16432 = v16410
		v16433 = v16399
		goto L2628
	} else {
		goto L2632
	}
L2632:
	;
	v16412 = v16399 + int32(1)
	if v16412 == v16400 {
		goto L2629
	} else {
		goto L2633
	}
L2633:
	;
	v16415 = v16412
	goto L2634
L2634:
	;
	v16422 = *(*int32)(unsafe.Add(mBase, uint32(v16403+v16415<<(uint(int32(2))%32))))
	if v16422 != 0 {
		v16432 = v16422
		v16433 = v16415
		goto L2628
	} else {
		goto L2636
	}
L2635:
	;
	goto L2629
L2636:
	;
	v16424 = v16415 + int32(1)
	if v16424 != v16400 {
		v16415 = v16424
		goto L2634
	} else {
		goto L2637
	}
L2637:
	;
	goto L2635
L2638:
	;
	goto L2564
L2639:
	;
	v16457 = v16357
	v16462 = v16362
	v16467 = v16367
	v16469 = v16369
	v16470 = v16370
	goto L2562
L2640:
	;
	v16498 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16498 == int32(2) {
		goto L2641
	} else {
		goto L2642
	}
L2641:
	;
	v16501 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+16)) = v16501
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+136)) = v16501
	v16507 = F_list_make1_impl(m, int32(1), v15855+int32(16))
	mBase = m.M
	v16508 = m.ExcPending
	if v16508 != 0 {
		goto L1
	} else {
		goto L2644
	}
L2642:
	;
	v16510 = v16470
	goto L2643
L2643:
	;
	v16511 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+152))
	if v16511 != 0 {
		goto L2645
	} else {
		goto L2646
	}
L2644:
	;
	v16510 = v16507
	goto L2643
L2645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+12)) = v16511
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+132)) = v16511
	v16517 = F_list_make1_impl(m, int32(1), v15855+int32(12))
	mBase = m.M
	v16518 = m.ExcPending
	if v16518 != 0 {
		goto L1
	} else {
		goto L2648
	}
L2646:
	;
	v16519 = v16469
	goto L2647
L2647:
	;
	v16520 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+96))
	if v16520 != 0 {
		goto L2649
	} else {
		goto L2650
	}
L2648:
	;
	v16519 = v16517
	goto L2647
L2649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+8)) = v16520
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+128)) = v16520
	v16526 = F_list_make1_impl(m, int32(1), v15855+int32(8))
	mBase = m.M
	v16527 = m.ExcPending
	if v16527 != 0 {
		goto L1
	} else {
		goto L2652
	}
L2650:
	;
	v16528 = v16462
	goto L2651
L2651:
	;
	v16529 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+64))
	if v16529 != 0 {
		goto L2653
	} else {
		goto L2654
	}
L2652:
	;
	v16528 = v16526
	goto L2651
L2653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+4)) = v16529
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+124)) = v16529
	v16535 = F_list_make1_impl(m, int32(1), v15855+int32(4))
	mBase = m.M
	v16536 = m.ExcPending
	if v16536 != 0 {
		goto L1
	} else {
		goto L2656
	}
L2654:
	;
	v16537 = v16467
	goto L2655
L2655:
	;
	v16538 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16538 != int32(5) {
		v16651 = v16496
		v16657 = v16457
		v16662 = v16528
		v16664 = v16012
		v16667 = v16537
		v16669 = v16519
		v16670 = v16510
		goto L2543
	} else {
		goto L2657
	}
L2656:
	;
	v16537 = v16535
	goto L2655
L2657:
	;
	v16603 = v16496
	v16614 = v16528
	v16616 = v16012
	v16619 = v16537
	v16621 = v16519
	v16622 = v16510
	v16642 = v15855 + int32(120)
	goto L2544
L2658:
	;
	v16550 = int32(0)
	v16552 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16552 == int32(2) {
		goto L2659
	} else {
		goto L2660
	}
L2659:
	;
	v16555 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+36)) = v16555
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+112)) = v16555
	v16561 = F_list_make1_impl(m, int32(1), v15855+int32(36))
	mBase = m.M
	v16562 = m.ExcPending
	if v16562 != 0 {
		goto L1
	} else {
		goto L2662
	}
L2660:
	;
	v16564 = v16550
	goto L2661
L2661:
	;
	v16565 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+152))
	if v16565 != 0 {
		goto L2663
	} else {
		goto L2664
	}
L2662:
	;
	v16564 = v16561
	goto L2661
L2663:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+32)) = v16565
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+108)) = v16565
	v16571 = F_list_make1_impl(m, int32(1), v15855+int32(32))
	mBase = m.M
	v16572 = m.ExcPending
	if v16572 != 0 {
		goto L1
	} else {
		goto L2666
	}
L2664:
	;
	v16573 = v16550
	goto L2665
L2665:
	;
	v16574 = int32(0)
	v16576 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+96))
	if v16576 != 0 {
		goto L2667
	} else {
		goto L2668
	}
L2666:
	;
	v16573 = v16571
	goto L2665
L2667:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+28)) = v16576
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+104)) = v16576
	v16582 = F_list_make1_impl(m, int32(1), v15855+int32(28))
	mBase = m.M
	v16583 = m.ExcPending
	if v16583 != 0 {
		goto L1
	} else {
		goto L2670
	}
L2668:
	;
	v16584 = v16574
	goto L2669
L2669:
	;
	v16585 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+64))
	if v16585 != 0 {
		goto L2671
	} else {
		goto L2672
	}
L2670:
	;
	v16584 = v16582
	goto L2669
L2671:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+24)) = v16585
	*(*int32)(unsafe.Add(mBase, uint32(v15855)+100)) = v16585
	v16591 = F_list_make1_impl(m, int32(1), v15855+int32(24))
	mBase = m.M
	v16592 = m.ExcPending
	if v16592 != 0 {
		goto L1
	} else {
		goto L2674
	}
L2672:
	;
	v16593 = v16574
	goto L2673
L2673:
	;
	v16594 = int32(0)
	v16595 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+4))
	if v16595 != int32(5) {
		goto L2675
	} else {
		goto L2676
	}
L2674:
	;
	v16593 = v16591
	goto L2673
L2675:
	;
	v16651 = v16548
	v16657 = int32(0)
	v16662 = v16584
	v16664 = v16594
	v16667 = v16593
	v16669 = v16573
	v16670 = v16564
	goto L2543
L2676:
	;
	goto L2677
L2677:
	;
	v16603 = v16548
	v16614 = v16584
	v16616 = v16594
	v16619 = v16593
	v16621 = v16573
	v16622 = v16564
	v16642 = v15855 + int32(96)
	goto L2544
L2678:
	;
	v16651 = v16603
	v16657 = v16647
	v16662 = v16614
	v16664 = v16616
	v16667 = v16619
	v16669 = v16621
	v16670 = v16622
	goto L2543
L2679:
	;
	v16697 = int32(0)
	goto L2681
L2680:
	;
	v16696 = *(*int32)(unsafe.Add(mBase, uint32(v15850)+136))
	v16697 = v16696
	goto L2681
L2681:
	;
	v16698 = *(*int32)(unsafe.Add(mBase, uint32(v15860)+84))
	v16699 = F_assign_special_exec_param(m, v15850)
	mBase = m.M
	v16700 = m.ExcPending
	if v16700 != 0 {
		goto L1
	} else {
		goto L2682
	}
L2682:
	;
	v16703 = F_palloc0(m, int32(136))
	mBase = m.M
	v16704 = m.ExcPending
	if v16704 != 0 {
		goto L1
	} else {
		goto L2683
	}
L2683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+8)) = v15844
	*(*int64)(unsafe.Add(mBase, uint32(v16703))) = int64(1430224109884)
	v16708 = *(*int32)(unsafe.Add(mBase, uint32(v15844)+28))
	v16709 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+64)) = v16709
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+24)) = v16709
	*(*uint16)(unsafe.Add(mBase, uint32(v16703)+20)) = uint16(v16709)
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+16)) = v16709
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+12)) = v16708
	v16718 = *(*int32)(unsafe.Add(mBase, uint32(v15956)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+40)) = v16718
	v16720 = *(*float64)(unsafe.Add(mBase, uint32(v15956)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16703)+48)) = v16720
	v16722 = *(*float64)(unsafe.Add(mBase, uint32(v15956)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16703)+56)) = v16722
	if v16662 != 0 {
		goto L2685
	} else {
		goto L2686
	}
L2684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16708)+32)) = v16730
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+128)) = v16657
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+124)) = v16667
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+120)) = v16699
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+116)) = v16698
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+112)) = v16697
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+108)) = v16662
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+104)) = v16669
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+100)) = v16670
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+96)) = v16651
	*(*uint8)(unsafe.Add(mBase, uint32(v16703)+92)) = uint8(v16693)
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+88)) = v16664
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+84)) = v16692
	*(*uint8)(unsafe.Add(mBase, uint32(v16703)+80)) = uint8(v16691)
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+76)) = v16690
	*(*int32)(unsafe.Add(mBase, uint32(v16703)+72)) = v15956
	v16747 = v15844
	v16753 = v15850
	v16756 = v15853
	v16757 = v15854
	v16758 = v15855
	v16763 = v15860
	v16772 = v15869
	v16781 = v15878
	v16785 = v15882
	v16786 = v15883
	v16788 = v16703
	goto L2526
L2685:
	;
	v16724 = *(*float64)(unsafe.Add(mBase, uint32(v15956)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v16703)+32)) = v16724
	v16726 = *(*int32)(unsafe.Add(mBase, uint32(v15956)+12))
	v16727 = *(*int32)(unsafe.Add(mBase, uint32(v16726)+32))
	v16730 = v16727
	goto L2684
L2686:
	;
	goto L2687
L2687:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16703)+32)) = int64(0)
	v16730 = int32(0)
	goto L2684
L2688:
	;
	v16792 = v15866 + int32(1)
	v16793 = *(*int32)(unsafe.Add(mBase, uint32(v16756)+4))
	if v16792 < v16793 {
		v15844 = v16747
		v15850 = v16753
		v15853 = v16756
		v15854 = v16757
		v15855 = v16758
		v15860 = v16763
		v15866 = v16792
		v15869 = v16772
		v15878 = v16781
		v15882 = v16785
		v15883 = v16786
		goto L2505
	} else {
		goto L2689
	}
L2689:
	;
	goto L2506
L2690:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16806)+216)) = v16833
	*(*int64)(unsafe.Add(mBase, uint32(v16806)+208)) = v16834
	*(*float64)(unsafe.Add(mBase, uint32(v16806)+200)) = v16829
	*(*uint8)(unsafe.Add(mBase, uint32(v16806)+192)) = uint8(v16992)
	v17030 = *(*int32)(unsafe.Add(mBase, uint32(v16795)+168))
	if v17030 == int32(0) {
		goto L2724
	} else {
		goto L2725
	}
L2691:
	;
	v16961 = *(*int32)(unsafe.Add(mBase, uint32(v16811)+132))
	if v16961 != 0 {
		goto L2712
	} else {
		goto L2713
	}
L2692:
	;
	v16839 = *(*int32)(unsafe.Add(mBase, uint32(v16801)+12))
	if base.Ui32(v16839) < base.Ui32(int32(2)) {
		goto L2691
	} else {
		goto L2693
	}
L2693:
	;
	v16842 = *(*int32)(unsafe.Add(mBase, uint32(v16811)+132))
	if v16842 != 0 {
		goto L2694
	} else {
		goto L2695
	}
L2694:
	;
	v16843 = *(*int32)(unsafe.Add(mBase, uint32(v16842)))
	if v16843 != int32(7) {
		goto L2697
	} else {
		goto L2698
	}
L2695:
	;
	goto L2696
L2696:
	;
	v16850 = *(*int32)(unsafe.Add(mBase, uint32(v16811)+128))
	if v16850 == int32(0) {
		goto L2701
	} else {
		goto L2702
	}
L2697:
	;
	v16992 = int32(1)
	goto L2690
L2698:
	;
	goto L2699
L2699:
	;
	v16847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16842)+24)))
	if v16847 != int32(1) {
		goto L2691
	} else {
		goto L2700
	}
L2700:
	;
	goto L2696
L2701:
	;
	v16861 = *(*int32)(unsafe.Add(mBase, uint32(v16805)+40))
	if v16861 == int32(0) {
		goto L2691
	} else {
		goto L2706
	}
L2702:
	;
	v16853 = *(*int32)(unsafe.Add(mBase, uint32(v16850)))
	if v16853 != int32(7) {
		goto L2691
	} else {
		goto L2703
	}
L2703:
	;
	v16856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16850)+24)))
	if v16856 != 0 {
		goto L2701
	} else {
		goto L2704
	}
L2704:
	;
	v16857 = *(*int32)(unsafe.Add(mBase, uint32(v16850)+20))
	v16858 = *(*int64)(unsafe.Add(mBase, uint32(v16857)))
	if v16858 != int64(0) {
		goto L2691
	} else {
		goto L2705
	}
L2705:
	;
	goto L2701
L2706:
	;
	v16864 = *(*int32)(unsafe.Add(mBase, uint32(v16861)+4))
	if v16864 <= int32(0) {
		goto L2691
	} else {
		goto L2707
	}
L2707:
	;
	v16875 = int32(0)
	goto L2708
L2708:
	;
	v16909 = *(*int32)(unsafe.Add(mBase, uint32(v16861)+12))
	v16913 = *(*int32)(unsafe.Add(mBase, uint32(v16909+v16875<<(uint(int32(2))%32))))
	F_add_partial_path(m, v16795, v16913)
	mBase = m.M
	v16915 = m.ExcPending
	if v16915 != 0 {
		goto L1
	} else {
		goto L2710
	}
L2709:
	;
	goto L2691
L2710:
	;
	v16917 = v16875 + int32(1)
	v16918 = *(*int32)(unsafe.Add(mBase, uint32(v16861)+4))
	if v16917 < v16918 {
		v16875 = v16917
		goto L2708
	} else {
		goto L2711
	}
L2711:
	;
	goto L2709
L2712:
	;
	v16962 = *(*int32)(unsafe.Add(mBase, uint32(v16961)))
	if v16962 != int32(7) {
		goto L2715
	} else {
		goto L2716
	}
L2713:
	;
	goto L2714
L2714:
	;
	v16971 = *(*int32)(unsafe.Add(mBase, uint32(v16811)+128))
	if v16971 == int32(0) {
		goto L2719
	} else {
		goto L2720
	}
L2715:
	;
	v16992 = int32(1)
	goto L2690
L2716:
	;
	goto L2717
L2717:
	;
	v16966 = int32(1)
	v16967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16961)+24)))
	if v16967 != v16966 {
		v16992 = v16966
		goto L2690
	} else {
		goto L2718
	}
L2718:
	;
	goto L2714
L2719:
	;
	v16992 = int32(0)
	goto L2690
L2720:
	;
	v16974 = int32(1)
	v16975 = *(*int32)(unsafe.Add(mBase, uint32(v16971)))
	if v16975 != int32(7) {
		v16992 = v16974
		goto L2690
	} else {
		goto L2721
	}
L2721:
	;
	v16978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16971)+24)))
	if v16978 != 0 {
		goto L2719
	} else {
		goto L2722
	}
L2722:
	;
	v16979 = *(*int32)(unsafe.Add(mBase, uint32(v16971)+20))
	v16980 = *(*int64)(unsafe.Add(mBase, uint32(v16979)))
	if v16980 != int64(0) {
		v16992 = v16974
		goto L2690
	} else {
		goto L2723
	}
L2723:
	;
	goto L2719
L2724:
	;
	v17043 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v17043 != 0 {
		goto L2728
	} else {
		goto L2729
	}
L2725:
	;
	v17033 = *(*int32)(unsafe.Add(mBase, uint32(v17030)+36))
	if v17033 == int32(0) {
		goto L2724
	} else {
		goto L2726
	}
L2726:
	;
	m.T0[v17033].(func(*base.Module, int32, int32, int32, int32, int32))(m, v16801, int32(7), v16805, v16795, v16806+int32(192))
	mBase = m.M
	v17040 = m.ExcPending
	if v17040 != 0 {
		goto L1
	} else {
		goto L2727
	}
L2727:
	;
	goto L2724
L2728:
	;
	m.T0[v17043].(func(*base.Module, int32, int32, int32, int32, int32))(m, v16801, int32(7), v16805, v16795, v16806+int32(192))
	mBase = m.M
	v17048 = m.ExcPending
	if v17048 != 0 {
		goto L1
	} else {
		goto L2731
	}
L2729:
	;
	goto L2730
L2730:
	;
	m.G0 = v16806 + int32(352)
	goto L711
L2731:
	;
	goto L2730
L2732:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v17058 = m.ExcPending
	if v17058 != 0 {
		goto L1
	} else {
		goto L2733
	}
L2733:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_48), int32(0))
	mBase = m.M
	v17062 = m.ExcPending
	if v17062 != 0 {
		goto L1
	} else {
		goto L2734
	}
L2734:
	;
	F_errdetail(m, int32(_a_F_subquery_planner_22), int32(0))
	mBase = m.M
	v17066 = m.ExcPending
	if v17066 != 0 {
		goto L1
	} else {
		goto L2735
	}
L2735:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_12), int32(_a_F_subquery_planner_49), int32(_a_F_subquery_planner_50))
	mBase = m.M
	v17071 = m.ExcPending
	if v17071 != 0 {
		goto L1
	} else {
		goto L2736
	}
L2736:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2737:
	;
	v17076 = F_fetch_upper_rel(m, v16801, int32(7), int32(0))
	mBase = m.M
	v17077 = m.ExcPending
	if v17077 != 0 {
		goto L1
	} else {
		goto L2738
	}
L2738:
	;
	v17078 = int32(0)
	v17085 = float64(0)
	v17086 = *(*int32)(unsafe.Add(mBase, uint32(v16801)+72))
	if v17086 == v17078 {
		goto L2740
	} else {
		goto L2741
	}
L2739:
	;
	F_set_cheapest(m, v17076)
	mBase = m.M
	v17274 = m.ExcPending
	if v17274 != 0 {
		goto L1
	} else {
		goto L2769
	}
L2740:
	;
	goto L2739
L2741:
	;
	v17089 = *(*int32)(unsafe.Add(mBase, uint32(v17086)+4))
	if v17089 <= int32(0) {
		v17162 = v17078
		v17168 = v17085
		goto L2742
	} else {
		goto L2743
	}
L2742:
	;
	v17169 = *(*int32)(unsafe.Add(mBase, uint32(v17076)+32))
	if v17169 == int32(0) {
		goto L2752
	} else {
		goto L2753
	}
L2743:
	;
	v17092 = *(*int32)(unsafe.Add(mBase, uint32(v17086)+12))
	if v17089 == int32(1) {
		goto L2745
	} else {
		goto L2746
	}
L2744:
	;
	v17150 = *(*int32)(unsafe.Add(mBase, uint32(v17092+v17137<<(uint(int32(2))%32))))
	v17151 = *(*float64)(unsafe.Add(mBase, uint32(v17150)+56))
	v17152 = *(*float64)(unsafe.Add(mBase, uint32(v17150)+64))
	v17155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17150)+38)))
	v17162 = v17155 ^ int32(1) | v17140
	v17168 = base.F64_add(v17146, base.F64_add(v17151, v17152))
	goto L2742
L2745:
	;
	v17137 = int32(0)
	v17140 = v17078
	v17146 = v17085
	goto L2744
L2746:
	;
	goto L2747
L2747:
	;
	v17101 = int32(0)
	v17104 = v17078
	v17109 = v17078
	v17110 = v17085
	goto L2748
L2748:
	;
	v17111 = int32(2)
	v17113 = v17092 + v17101<<(uint(v17111)%32)
	v17114 = *(*int32)(unsafe.Add(mBase, uint32(v17113)))
	v17115 = *(*float64)(unsafe.Add(mBase, uint32(v17114)+56))
	v17116 = *(*float64)(unsafe.Add(mBase, uint32(v17114)+64))
	v17119 = *(*int32)(unsafe.Add(mBase, uint32(v17113)+4))
	v17120 = *(*float64)(unsafe.Add(mBase, uint32(v17119)+56))
	v17121 = *(*float64)(unsafe.Add(mBase, uint32(v17119)+64))
	v17123 = base.F64_add(base.F64_add(v17110, base.F64_add(v17115, v17116)), base.F64_add(v17120, v17121))
	v17124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17119)+38)))
	v17125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17114)+38)))
	v17129 = base.B2i32(v17124&v17125 == int32(0)) | v17104
	v17131 = v17101 + v17111
	v17133 = v17109 + v17111
	if v17133 != v17089&int32(2147483646) {
		v17101 = v17131
		v17104 = v17129
		v17109 = v17133
		v17110 = v17123
		goto L2748
	} else {
		goto L2750
	}
L2749:
	;
	if v17089&int32(1) == int32(0) {
		v17162 = v17129
		v17168 = v17123
		goto L2742
	} else {
		goto L2751
	}
L2750:
	;
	goto L2749
L2751:
	;
	v17137 = v17131
	v17140 = v17129
	v17146 = v17123
	goto L2744
L2752:
	;
	if v17162&int32(1) != 0 {
		goto L2761
	} else {
		goto L2762
	}
L2753:
	;
	v17172 = *(*int32)(unsafe.Add(mBase, uint32(v17169)+4))
	if v17172 <= int32(0) {
		goto L2752
	} else {
		goto L2754
	}
L2754:
	;
	v17178 = int32(0)
	goto L2755
L2755:
	;
	v17188 = *(*int32)(unsafe.Add(mBase, uint32(v17169)+12))
	v17192 = *(*int32)(unsafe.Add(mBase, uint32(v17188+v17178<<(uint(int32(2))%32))))
	v17193 = *(*float64)(unsafe.Add(mBase, uint32(v17192)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v17192)+48)) = base.F64_add(v17168, v17193)
	v17196 = *(*float64)(unsafe.Add(mBase, uint32(v17192)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v17192)+56)) = base.F64_add(v17168, v17196)
	if v17162&int32(1) != 0 {
		goto L2757
	} else {
		goto L2758
	}
L2756:
	;
	goto L2752
L2757:
	;
	v17199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17192)+21)) = uint8(v17199)
	goto L2759
L2758:
	;
	goto L2759
L2759:
	;
	v17202 = v17178 + int32(1)
	v17203 = *(*int32)(unsafe.Add(mBase, uint32(v17169)+4))
	if v17202 < v17203 {
		v17178 = v17202
		goto L2755
	} else {
		goto L2760
	}
L2760:
	;
	goto L2756
L2761:
	;
	v17217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17076)+26)) = uint8(v17217)
	*(*int32)(unsafe.Add(mBase, uint32(v17076)+40)) = v17217
	goto L2739
L2762:
	;
	goto L2763
L2763:
	;
	v17221 = *(*int32)(unsafe.Add(mBase, uint32(v17076)+40))
	if v17221 == int32(0) {
		goto L2740
	} else {
		goto L2764
	}
L2764:
	;
	v17224 = *(*int32)(unsafe.Add(mBase, uint32(v17221)+4))
	if v17224 <= int32(0) {
		goto L2740
	} else {
		goto L2765
	}
L2765:
	;
	v17228 = int32(0)
	goto L2766
L2766:
	;
	v17238 = *(*int32)(unsafe.Add(mBase, uint32(v17221)+12))
	v17242 = *(*int32)(unsafe.Add(mBase, uint32(v17238+v17228<<(uint(int32(2))%32))))
	v17243 = *(*float64)(unsafe.Add(mBase, uint32(v17242)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v17242)+48)) = base.F64_add(v17168, v17243)
	v17246 = *(*float64)(unsafe.Add(mBase, uint32(v17242)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v17242)+56)) = base.F64_add(v17168, v17246)
	v17250 = v17228 + int32(1)
	v17251 = *(*int32)(unsafe.Add(mBase, uint32(v17221)+4))
	if v17250 < v17251 {
		v17228 = v17250
		goto L2766
	} else {
		goto L2768
	}
L2767:
	;
	goto L2740
L2768:
	;
	goto L2767
L2769:
	;
	m.G0 = v16820 + int32(16)
	return v16801
}
