package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_subquery_planner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
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
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
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
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v698 int32
	_ = v698
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int64
	_ = v737
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
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
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
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
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1055 int32
	_ = v1055
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
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
	var v1149 int32
	_ = v1149
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
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
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
	var v1279 int32
	_ = v1279
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1450 int32
	_ = v1450
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1464 int32
	_ = v1464
	var v1475 int32
	_ = v1475
	var v1482 int32
	_ = v1482
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1513 int32
	_ = v1513
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1562 int32
	_ = v1562
	var v1569 int32
	_ = v1569
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1679 int32
	_ = v1679
	var v1686 int32
	_ = v1686
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1968 int32
	_ = v1968
	var v1969 int32
	_ = v1969
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2004 int32
	_ = v2004
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2021 int32
	_ = v2021
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2158 int32
	_ = v2158
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2240 int32
	_ = v2240
	var v2278 int32
	_ = v2278
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2353 int32
	_ = v2353
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2369 int32
	_ = v2369
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2387 int32
	_ = v2387
	var v2425 int32
	_ = v2425
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2506 int32
	_ = v2506
	var v2543 int32
	_ = v2543
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
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2565 int32
	_ = v2565
	var v2566 int32
	_ = v2566
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2575 int32
	_ = v2575
	var v2578 int32
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2582 int32
	_ = v2582
	var v2585 int32
	_ = v2585
	var v2588 int32
	_ = v2588
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2606 int32
	_ = v2606
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2616 int32
	_ = v2616
	var v2619 int32
	_ = v2619
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2647 int32
	_ = v2647
	var v2687 int32
	_ = v2687
	var v2690 int32
	_ = v2690
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2697 int32
	_ = v2697
	var v2698 int32
	_ = v2698
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2786 int32
	_ = v2786
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2797 int32
	_ = v2797
	var v2837 int32
	_ = v2837
	var v2841 int32
	_ = v2841
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2903 int32
	_ = v2903
	var v2904 int32
	_ = v2904
	var v2944 int32
	_ = v2944
	var v2948 int32
	_ = v2948
	var v2952 int32
	_ = v2952
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2961 int32
	_ = v2961
	var v2964 int32
	_ = v2964
	var v3005 int32
	_ = v3005
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3015 int32
	_ = v3015
	var v3056 int32
	_ = v3056
	var v3059 int32
	_ = v3059
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3064 int32
	_ = v3064
	var v3066 int32
	_ = v3066
	var v3067 int32
	_ = v3067
	var v3069 int32
	_ = v3069
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3079 int32
	_ = v3079
	var v3082 int32
	_ = v3082
	var v3119 int32
	_ = v3119
	var v3123 int32
	_ = v3123
	var v3124 int32
	_ = v3124
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3131 int32
	_ = v3131
	var v3132 int32
	_ = v3132
	var v3135 int32
	_ = v3135
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3164 int32
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3169 int32
	_ = v3169
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3172 int32
	_ = v3172
	var v3173 int32
	_ = v3173
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3179 int32
	_ = v3179
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3189 int32
	_ = v3189
	var v3229 int32
	_ = v3229
	var v3231 int32
	_ = v3231
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3242 int32
	_ = v3242
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3257 int32
	_ = v3257
	var v3258 int32
	_ = v3258
	var v3260 int32
	_ = v3260
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3266 int32
	_ = v3266
	var v3269 int32
	_ = v3269
	var v3273 int32
	_ = v3273
	var v3314 int32
	_ = v3314
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3322 int32
	_ = v3322
	var v3323 int32
	_ = v3323
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3397 int32
	_ = v3397
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3444 int32
	_ = v3444
	var v3448 int32
	_ = v3448
	var v3449 int32
	_ = v3449
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3454 int32
	_ = v3454
	var v3455 int32
	_ = v3455
	var v3456 int32
	_ = v3456
	var v3457 int32
	_ = v3457
	var v3458 int32
	_ = v3458
	var v3462 int32
	_ = v3462
	var v3463 int32
	_ = v3463
	var v3464 int32
	_ = v3464
	var v3466 int32
	_ = v3466
	var v3467 int32
	_ = v3467
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3472 int32
	_ = v3472
	var v3473 int32
	_ = v3473
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3478 int32
	_ = v3478
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3523 int32
	_ = v3523
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3532 int32
	_ = v3532
	var v3533 int32
	_ = v3533
	var v3539 int32
	_ = v3539
	var v3540 int32
	_ = v3540
	var v3543 int32
	_ = v3543
	var v3544 int32
	_ = v3544
	var v3548 int32
	_ = v3548
	var v3549 int32
	_ = v3549
	var v3637 int32
	_ = v3637
	var v3640 int32
	_ = v3640
	var v3642 int32
	_ = v3642
	var v3644 int32
	_ = v3644
	var v3645 int32
	_ = v3645
	var v3648 int32
	_ = v3648
	var v3650 int64
	_ = v3650
	var v3651 int32
	_ = v3651
	var v3652 int32
	_ = v3652
	var v3655 int32
	_ = v3655
	var v3659 int32
	_ = v3659
	var v3660 int64
	_ = v3660
	var v3661 int64
	_ = v3661
	var v3664 int64
	_ = v3664
	var v3666 int64
	_ = v3666
	var v3667 int32
	_ = v3667
	var v3670 int64
	_ = v3670
	var v3671 int32
	_ = v3671
	var v3672 int32
	_ = v3672
	var v3675 int32
	_ = v3675
	var v3679 int32
	_ = v3679
	var v3680 int64
	_ = v3680
	var v3681 int64
	_ = v3681
	var v3684 int64
	_ = v3684
	var v3686 int64
	_ = v3686
	var v3692 float64
	_ = v3692
	var v3696 float64
	_ = v3696
	var v3704 float64
	_ = v3704
	var v3713 float64
	_ = v3713
	var v3714 int64
	_ = v3714
	var v3726 float64
	_ = v3726
	var v3736 float64
	_ = v3736
	var v3743 float64
	_ = v3743
	var v3744 float64
	_ = v3744
	var v3748 float64
	_ = v3748
	var v3752 float64
	_ = v3752
	var v3754 float64
	_ = v3754
	var v3756 float64
	_ = v3756
	var v3757 int64
	_ = v3757
	var v3758 int64
	_ = v3758
	var v3760 int32
	_ = v3760
	var v3762 int32
	_ = v3762
	var v3764 int32
	_ = v3764
	var v3766 int32
	_ = v3766
	var v3767 int32
	_ = v3767
	var v3768 int32
	_ = v3768
	var v3771 int32
	_ = v3771
	var v3773 int32
	_ = v3773
	var v3813 int32
	_ = v3813
	var v3814 int32
	_ = v3814
	var v3817 int32
	_ = v3817
	var v3818 int32
	_ = v3818
	var v3822 int32
	_ = v3822
	var v3823 int32
	_ = v3823
	var v3824 int32
	_ = v3824
	var v3827 int32
	_ = v3827
	var v3830 int32
	_ = v3830
	var v3832 int32
	_ = v3832
	var v3833 int32
	_ = v3833
	var v3834 int32
	_ = v3834
	var v3839 int32
	_ = v3839
	var v3840 int32
	_ = v3840
	var v3841 int32
	_ = v3841
	var v3844 int32
	_ = v3844
	var v3845 int32
	_ = v3845
	var v3846 int32
	_ = v3846
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3852 int32
	_ = v3852
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3860 int32
	_ = v3860
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3866 int32
	_ = v3866
	var v3867 int32
	_ = v3867
	var v3870 int32
	_ = v3870
	var v3871 int32
	_ = v3871
	var v3874 int32
	_ = v3874
	var v3875 int32
	_ = v3875
	var v3877 int32
	_ = v3877
	var v3882 int32
	_ = v3882
	var v3883 int32
	_ = v3883
	var v3884 int32
	_ = v3884
	var v3885 int32
	_ = v3885
	var v3888 int32
	_ = v3888
	var v3889 int32
	_ = v3889
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3893 int32
	_ = v3893
	var v3894 int32
	_ = v3894
	var v3895 int32
	_ = v3895
	var v3896 int32
	_ = v3896
	var v3897 int32
	_ = v3897
	var v3899 int32
	_ = v3899
	var v3900 int32
	_ = v3900
	var v3903 int32
	_ = v3903
	var v3904 int32
	_ = v3904
	var v3905 int32
	_ = v3905
	var v3906 int32
	_ = v3906
	var v3908 int32
	_ = v3908
	var v3911 int32
	_ = v3911
	var v3915 int32
	_ = v3915
	var v3916 int32
	_ = v3916
	var v3956 int32
	_ = v3956
	var v3957 int32
	_ = v3957
	var v3958 int32
	_ = v3958
	var v3959 int32
	_ = v3959
	var v3960 int32
	_ = v3960
	var v3963 int32
	_ = v3963
	var v3964 int32
	_ = v3964
	var v3967 int32
	_ = v3967
	var v3973 int32
	_ = v3973
	var v3975 int32
	_ = v3975
	var v3976 int32
	_ = v3976
	var v4019 int32
	_ = v4019
	var v4026 int32
	_ = v4026
	var v4029 int32
	_ = v4029
	var v4032 int32
	_ = v4032
	var v4033 int32
	_ = v4033
	var v4037 int32
	_ = v4037
	var v4041 int32
	_ = v4041
	var v4042 int32
	_ = v4042
	var v4046 int32
	_ = v4046
	var v4050 int32
	_ = v4050
	var v4056 int32
	_ = v4056
	var v4059 float64
	_ = v4059
	var v4062 float64
	_ = v4062
	var v4064 int32
	_ = v4064
	var v4066 int32
	_ = v4066
	var v4071 float64
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4106 int32
	_ = v4106
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4119 int32
	_ = v4119
	var v4123 int32
	_ = v4123
	var v4126 int32
	_ = v4126
	var v4127 int32
	_ = v4127
	var v4129 int32
	_ = v4129
	var v4131 int32
	_ = v4131
	var v4141 float64
	_ = v4141
	var v4142 float64
	_ = v4142
	var v4143 float64
	_ = v4143
	var v4144 int32
	_ = v4144
	var v4145 int64
	_ = v4145
	var v4146 float64
	_ = v4146
	var v4147 float64
	_ = v4147
	var v4148 float64
	_ = v4148
	var v4152 float64
	_ = v4152
	var v4157 int64
	_ = v4157
	var v4168 int32
	_ = v4168
	var v4169 int32
	_ = v4169
	var v4170 int32
	_ = v4170
	var v4171 int32
	_ = v4171
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4177 int32
	_ = v4177
	var v4179 int32
	_ = v4179
	var v4180 int32
	_ = v4180
	var v4184 int32
	_ = v4184
	var v4186 int32
	_ = v4186
	var v4188 int32
	_ = v4188
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4195 int32
	_ = v4195
	var v4196 int32
	_ = v4196
	var v4197 int32
	_ = v4197
	var v4200 int32
	_ = v4200
	var v4203 int32
	_ = v4203
	var v4246 int32
	_ = v4246
	var v4250 int32
	_ = v4250
	var v4255 int32
	_ = v4255
	var v4259 int32
	_ = v4259
	var v4262 int32
	_ = v4262
	var v4266 int32
	_ = v4266
	var v4269 int32
	_ = v4269
	var v4270 int32
	_ = v4270
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4277 int32
	_ = v4277
	var v4278 int32
	_ = v4278
	var v4279 int32
	_ = v4279
	var v4280 int32
	_ = v4280
	var v4282 int32
	_ = v4282
	var v4285 int32
	_ = v4285
	var v4289 int32
	_ = v4289
	var v4290 int32
	_ = v4290
	var v4305 int32
	_ = v4305
	var v4330 int32
	_ = v4330
	var v4334 int32
	_ = v4334
	var v4335 int32
	_ = v4335
	var v4338 int32
	_ = v4338
	var v4339 int32
	_ = v4339
	var v4342 int32
	_ = v4342
	var v4343 int32
	_ = v4343
	var v4344 int32
	_ = v4344
	var v4347 int32
	_ = v4347
	var v4353 int32
	_ = v4353
	var v4354 int32
	_ = v4354
	var v4355 int32
	_ = v4355
	var v4358 int32
	_ = v4358
	var v4360 int32
	_ = v4360
	var v4378 int32
	_ = v4378
	var v4404 int32
	_ = v4404
	var v4405 int32
	_ = v4405
	var v4406 int32
	_ = v4406
	var v4407 int32
	_ = v4407
	var v4408 int32
	_ = v4408
	var v4409 int32
	_ = v4409
	var v4413 int32
	_ = v4413
	var v4414 int32
	_ = v4414
	var v4415 int32
	_ = v4415
	var v4416 int32
	_ = v4416
	var v4417 int32
	_ = v4417
	var v4419 int32
	_ = v4419
	var v4420 int32
	_ = v4420
	var v4422 int32
	_ = v4422
	var v4423 int32
	_ = v4423
	var v4424 int32
	_ = v4424
	var v4426 int32
	_ = v4426
	var v4432 int32
	_ = v4432
	var v4433 int32
	_ = v4433
	var v4436 int32
	_ = v4436
	var v4448 int32
	_ = v4448
	var v4455 int32
	_ = v4455
	var v4480 int32
	_ = v4480
	var v4484 int32
	_ = v4484
	var v4485 int32
	_ = v4485
	var v4486 int32
	_ = v4486
	var v4489 int32
	_ = v4489
	var v4490 int32
	_ = v4490
	var v4491 int32
	_ = v4491
	var v4493 int32
	_ = v4493
	var v4496 int32
	_ = v4496
	var v4497 int32
	_ = v4497
	var v4498 int32
	_ = v4498
	var v4501 int32
	_ = v4501
	var v4503 int32
	_ = v4503
	var v4504 int32
	_ = v4504
	var v4511 int32
	_ = v4511
	var v4551 int32
	_ = v4551
	var v4552 int32
	_ = v4552
	var v4554 int32
	_ = v4554
	var v4555 int32
	_ = v4555
	var v4558 int32
	_ = v4558
	var v4561 int32
	_ = v4561
	var v4563 int32
	_ = v4563
	var v4564 int32
	_ = v4564
	var v4604 int32
	_ = v4604
	var v4605 int32
	_ = v4605
	var v4609 int32
	_ = v4609
	var v4610 int32
	_ = v4610
	var v4611 int32
	_ = v4611
	var v4613 int32
	_ = v4613
	var v4614 int32
	_ = v4614
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4620 int32
	_ = v4620
	var v4622 int32
	_ = v4622
	var v4623 int32
	_ = v4623
	var v4624 int32
	_ = v4624
	var v4630 int32
	_ = v4630
	var v4633 int32
	_ = v4633
	var v4637 int32
	_ = v4637
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4646 int32
	_ = v4646
	var v4647 int32
	_ = v4647
	var v4648 int32
	_ = v4648
	var v4649 int32
	_ = v4649
	var v4652 int32
	_ = v4652
	var v4653 int32
	_ = v4653
	var v4656 int32
	_ = v4656
	var v4660 int32
	_ = v4660
	var v4661 int32
	_ = v4661
	var v4666 int32
	_ = v4666
	var v4670 int32
	_ = v4670
	var v4675 int32
	_ = v4675
	var v4679 int32
	_ = v4679
	var v4683 int32
	_ = v4683
	var v4688 int32
	_ = v4688
	var v4692 int32
	_ = v4692
	var v4695 int32
	_ = v4695
	var v4696 int32
	_ = v4696
	var v4697 int32
	_ = v4697
	var v4698 int32
	_ = v4698
	var v4699 int32
	_ = v4699
	var v4701 int32
	_ = v4701
	var v4706 int32
	_ = v4706
	var v4708 int32
	_ = v4708
	var v4714 int32
	_ = v4714
	var v4719 int32
	_ = v4719
	var v4725 int32
	_ = v4725
	var v4763 int32
	_ = v4763
	var v4766 int32
	_ = v4766
	var v4768 int32
	_ = v4768
	var v4773 int32
	_ = v4773
	var v4791 int32
	_ = v4791
	var v4813 int32
	_ = v4813
	var v4817 int32
	_ = v4817
	var v4822 int32
	_ = v4822
	var v4865 int32
	_ = v4865
	var v4866 int32
	_ = v4866
	var v4867 int32
	_ = v4867
	var v4869 int32
	_ = v4869
	var v4870 int32
	_ = v4870
	var v4871 int32
	_ = v4871
	var v4872 int32
	_ = v4872
	var v4873 int32
	_ = v4873
	var v4874 int32
	_ = v4874
	var v4875 int32
	_ = v4875
	var v4878 int32
	_ = v4878
	var v4879 int32
	_ = v4879
	var v4880 int32
	_ = v4880
	var v4883 int32
	_ = v4883
	var v4884 int32
	_ = v4884
	var v4893 int32
	_ = v4893
	var v4900 int32
	_ = v4900
	var v4904 int32
	_ = v4904
	var v4907 int32
	_ = v4907
	var v4927 int32
	_ = v4927
	var v4931 int32
	_ = v4931
	var v4932 int32
	_ = v4932
	var v4934 int32
	_ = v4934
	var v4938 int32
	_ = v4938
	var v4946 int32
	_ = v4946
	var v4978 int32
	_ = v4978
	var v4982 int32
	_ = v4982
	var v4983 int32
	_ = v4983
	var v4984 int32
	_ = v4984
	var v4985 int32
	_ = v4985
	var v4987 int32
	_ = v4987
	var v4998 int32
	_ = v4998
	var v4999 int32
	_ = v4999
	var v5032 int32
	_ = v5032
	var v5034 int32
	_ = v5034
	var v5035 int32
	_ = v5035
	var v5045 int32
	_ = v5045
	var v5079 int32
	_ = v5079
	var v5120 int32
	_ = v5120
	var v5122 int32
	_ = v5122
	var v5123 int32
	_ = v5123
	var v5137 int32
	_ = v5137
	var v5138 int32
	_ = v5138
	var v5140 int32
	_ = v5140
	var v5143 int32
	_ = v5143
	var v5144 int32
	_ = v5144
	var v5149 int32
	_ = v5149
	var v5157 int32
	_ = v5157
	var v5159 int32
	_ = v5159
	var v5161 int32
	_ = v5161
	var v5162 int32
	_ = v5162
	var v5165 int32
	_ = v5165
	var v5169 int32
	_ = v5169
	var v5177 int32
	_ = v5177
	var v5181 int32
	_ = v5181
	var v5182 int32
	_ = v5182
	var v5183 int32
	_ = v5183
	var v5184 int32
	_ = v5184
	var v5187 int32
	_ = v5187
	var v5188 int32
	_ = v5188
	var v5191 int32
	_ = v5191
	var v5199 int32
	_ = v5199
	var v5201 int32
	_ = v5201
	var v5210 int32
	_ = v5210
	var v5236 int32
	_ = v5236
	var v5241 int32
	_ = v5241
	var v5242 int32
	_ = v5242
	var v5246 int32
	_ = v5246
	var v5248 int32
	_ = v5248
	var v5251 int32
	_ = v5251
	var v5252 int32
	_ = v5252
	var v5295 int32
	_ = v5295
	var v5296 int32
	_ = v5296
	var v5305 int32
	_ = v5305
	var v5306 int32
	_ = v5306
	var v5308 int32
	_ = v5308
	var v5311 int32
	_ = v5311
	var v5312 int32
	_ = v5312
	var v5317 int32
	_ = v5317
	var v5324 int32
	_ = v5324
	var v5326 int32
	_ = v5326
	var v5328 int32
	_ = v5328
	var v5331 int32
	_ = v5331
	var v5333 int32
	_ = v5333
	var v5335 int32
	_ = v5335
	var v5342 int32
	_ = v5342
	var v5349 int32
	_ = v5349
	var v5350 int32
	_ = v5350
	var v5351 int32
	_ = v5351
	var v5356 int32
	_ = v5356
	var v5357 int32
	_ = v5357
	var v5368 int32
	_ = v5368
	var v5369 int32
	_ = v5369
	var v5370 int32
	_ = v5370
	var v5469 int32
	_ = v5469
	var v5480 int32
	_ = v5480
	var v5483 int32
	_ = v5483
	var v5504 int32
	_ = v5504
	var v5505 int32
	_ = v5505
	var v5528 int32
	_ = v5528
	var v5548 int32
	_ = v5548
	var v5550 int32
	_ = v5550
	var v5552 int32
	_ = v5552
	var v5553 int32
	_ = v5553
	var v5567 int32
	_ = v5567
	var v5569 int32
	_ = v5569
	var v5570 int32
	_ = v5570
	var v5571 int32
	_ = v5571
	var v5575 int32
	_ = v5575
	var v5576 int32
	_ = v5576
	var v5578 int32
	_ = v5578
	var v5579 int32
	_ = v5579
	var v5583 int32
	_ = v5583
	var v5584 int32
	_ = v5584
	var v5586 int32
	_ = v5586
	var v5587 int32
	_ = v5587
	var v5588 int32
	_ = v5588
	var v5592 int32
	_ = v5592
	var v5601 int32
	_ = v5601
	var v5604 int32
	_ = v5604
	var v5633 int32
	_ = v5633
	var v5634 int32
	_ = v5634
	var v5636 int32
	_ = v5636
	var v5639 int32
	_ = v5639
	var v5642 int32
	_ = v5642
	var v5643 int32
	_ = v5643
	var v5644 int32
	_ = v5644
	var v5650 int32
	_ = v5650
	var v5651 int32
	_ = v5651
	var v5657 int32
	_ = v5657
	var v5692 int32
	_ = v5692
	var v5693 int32
	_ = v5693
	var v5694 int32
	_ = v5694
	var v5696 int32
	_ = v5696
	var v5699 int32
	_ = v5699
	var v5701 int32
	_ = v5701
	var v5707 int32
	_ = v5707
	var v5709 int32
	_ = v5709
	var v5710 int32
	_ = v5710
	var v5711 int32
	_ = v5711
	var v5713 int32
	_ = v5713
	var v5714 int32
	_ = v5714
	var v5715 int32
	_ = v5715
	var v5717 int32
	_ = v5717
	var v5718 int32
	_ = v5718
	var v5720 int32
	_ = v5720
	var v5722 int32
	_ = v5722
	var v5728 int32
	_ = v5728
	var v5729 int32
	_ = v5729
	var v5730 int32
	_ = v5730
	var v5732 int32
	_ = v5732
	var v5737 int32
	_ = v5737
	var v5743 int32
	_ = v5743
	var v5778 int32
	_ = v5778
	var v5779 int32
	_ = v5779
	var v5780 int32
	_ = v5780
	var v5782 int32
	_ = v5782
	var v5783 int32
	_ = v5783
	var v5785 int32
	_ = v5785
	var v5787 int32
	_ = v5787
	var v5800 int32
	_ = v5800
	var v5834 int32
	_ = v5834
	var v5840 int32
	_ = v5840
	var v5844 int32
	_ = v5844
	var v5878 int32
	_ = v5878
	var v5881 int32
	_ = v5881
	var v5884 int32
	_ = v5884
	var v5885 int32
	_ = v5885
	var v5886 int32
	_ = v5886
	var v5888 int32
	_ = v5888
	var v5892 int32
	_ = v5892
	var v5895 int32
	_ = v5895
	var v5899 int32
	_ = v5899
	var v5905 int32
	_ = v5905
	var v5939 int32
	_ = v5939
	var v5940 int32
	_ = v5940
	var v5943 int32
	_ = v5943
	var v5947 int32
	_ = v5947
	var v5950 int32
	_ = v5950
	var v5951 int32
	_ = v5951
	var v5954 int32
	_ = v5954
	var v5955 int32
	_ = v5955
	var v5956 int32
	_ = v5956
	var v5964 int32
	_ = v5964
	var v5965 int32
	_ = v5965
	var v5976 int32
	_ = v5976
	var v6011 int32
	_ = v6011
	var v6054 int32
	_ = v6054
	var v6058 int32
	_ = v6058
	var v6099 int32
	_ = v6099
	var v6103 int32
	_ = v6103
	var v6104 int32
	_ = v6104
	var v6105 int32
	_ = v6105
	var v6108 int32
	_ = v6108
	var v6157 int32
	_ = v6157
	var v6159 int32
	_ = v6159
	var v6160 int32
	_ = v6160
	var v6161 int32
	_ = v6161
	var v6162 int32
	_ = v6162
	var v6163 int32
	_ = v6163
	var v6211 int32
	_ = v6211
	var v6215 int32
	_ = v6215
	var v6220 int32
	_ = v6220
	var v6223 int32
	_ = v6223
	var v6224 int32
	_ = v6224
	var v6228 int32
	_ = v6228
	var v6229 int32
	_ = v6229
	var v6230 int32
	_ = v6230
	var v6233 int32
	_ = v6233
	var v6236 int32
	_ = v6236
	var v6253 int32
	_ = v6253
	var v6277 int32
	_ = v6277
	var v6278 int32
	_ = v6278
	var v6280 int32
	_ = v6280
	var v6281 int32
	_ = v6281
	var v6283 int32
	_ = v6283
	var v6284 int32
	_ = v6284
	var v6293 int32
	_ = v6293
	var v6294 int32
	_ = v6294
	var v6303 int32
	_ = v6303
	var v6305 int32
	_ = v6305
	var v6306 int32
	_ = v6306
	var v6307 int32
	_ = v6307
	var v6313 int32
	_ = v6313
	var v6319 int32
	_ = v6319
	var v6320 int32
	_ = v6320
	var v6331 int32
	_ = v6331
	var v6363 int32
	_ = v6363
	var v6364 int32
	_ = v6364
	var v6366 int32
	_ = v6366
	var v6369 int32
	_ = v6369
	var v6370 int32
	_ = v6370
	var v6372 int32
	_ = v6372
	var v6373 int32
	_ = v6373
	var v6374 int32
	_ = v6374
	var v6377 int32
	_ = v6377
	var v6380 int32
	_ = v6380
	var v6397 int32
	_ = v6397
	var v6423 int32
	_ = v6423
	var v6440 int32
	_ = v6440
	var v6443 int32
	_ = v6443
	var v6466 int32
	_ = v6466
	var v6467 int32
	_ = v6467
	var v6469 int32
	_ = v6469
	var v6514 int32
	_ = v6514
	var v6526 int32
	_ = v6526
	var v6531 int32
	_ = v6531
	var v6561 int32
	_ = v6561
	var v6562 int32
	_ = v6562
	var v6563 int32
	_ = v6563
	var v6565 int32
	_ = v6565
	var v6581 int32
	_ = v6581
	var v6608 int32
	_ = v6608
	var v6610 int32
	_ = v6610
	var v6611 int32
	_ = v6611
	var v6613 int32
	_ = v6613
	var v6614 int32
	_ = v6614
	var v6616 int32
	_ = v6616
	var v6617 int32
	_ = v6617
	var v6619 int32
	_ = v6619
	var v6621 int32
	_ = v6621
	var v6623 int32
	_ = v6623
	var v6625 int32
	_ = v6625
	var v6630 int32
	_ = v6630
	var v6673 int32
	_ = v6673
	var v6675 int32
	_ = v6675
	var v6677 int32
	_ = v6677
	var v6680 int32
	_ = v6680
	var v6682 int32
	_ = v6682
	var v6684 int32
	_ = v6684
	var v6687 int32
	_ = v6687
	var v6730 int32
	_ = v6730
	var v6732 int32
	_ = v6732
	var v6734 int32
	_ = v6734
	var v6737 int32
	_ = v6737
	var v6739 int32
	_ = v6739
	var v6741 int32
	_ = v6741
	var v6784 int32
	_ = v6784
	var v6799 int32
	_ = v6799
	var v6828 int32
	_ = v6828
	var v6840 int32
	_ = v6840
	var v6843 int32
	_ = v6843
	var v6845 int32
	_ = v6845
	var v6846 int32
	_ = v6846
	var v6847 int32
	_ = v6847
	var v6849 int32
	_ = v6849
	var v6850 int32
	_ = v6850
	var v6852 int32
	_ = v6852
	var v6855 int32
	_ = v6855
	var v6861 int32
	_ = v6861
	var v6866 float64
	_ = v6866
	var v6870 int64
	_ = v6870
	var v6871 int64
	_ = v6871
	var v6873 int32
	_ = v6873
	var v6877 int32
	_ = v6877
	var v6879 int32
	_ = v6879
	var v6880 int32
	_ = v6880
	var v6883 int32
	_ = v6883
	var v6885 int32
	_ = v6885
	var v6888 int32
	_ = v6888
	var v6889 int32
	_ = v6889
	var v6890 int32
	_ = v6890
	var v6893 int32
	_ = v6893
	var v6894 int32
	_ = v6894
	var v6897 int32
	_ = v6897
	var v6900 int32
	_ = v6900
	var v6906 int32
	_ = v6906
	var v6907 int32
	_ = v6907
	var v6938 int32
	_ = v6938
	var v6942 int32
	_ = v6942
	var v6943 int32
	_ = v6943
	var v6945 int32
	_ = v6945
	var v6948 int32
	_ = v6948
	var v6953 int32
	_ = v6953
	var v6956 int32
	_ = v6956
	var v6992 int32
	_ = v6992
	var v6996 int32
	_ = v6996
	var v6997 int32
	_ = v6997
	var v7000 int32
	_ = v7000
	var v7003 int32
	_ = v7003
	var v7046 int32
	_ = v7046
	var v7049 int32
	_ = v7049
	var v7092 int32
	_ = v7092
	var v7093 int32
	_ = v7093
	var v7096 int32
	_ = v7096
	var v7136 int32
	_ = v7136
	var v7137 int32
	_ = v7137
	var v7139 int32
	_ = v7139
	var v7142 int32
	_ = v7142
	var v7145 int32
	_ = v7145
	var v7146 int32
	_ = v7146
	var v7148 int32
	_ = v7148
	var v7151 int32
	_ = v7151
	var v7158 int32
	_ = v7158
	var v7160 int32
	_ = v7160
	var v7165 int32
	_ = v7165
	var v7166 int32
	_ = v7166
	var v7171 int32
	_ = v7171
	var v7174 int32
	_ = v7174
	var v7177 int32
	_ = v7177
	var v7181 int32
	_ = v7181
	var v7261 int32
	_ = v7261
	var v7263 int32
	_ = v7263
	var v7264 int32
	_ = v7264
	var v7268 int32
	_ = v7268
	var v7276 int32
	_ = v7276
	var v7308 int32
	_ = v7308
	var v7310 int32
	_ = v7310
	var v7318 int32
	_ = v7318
	var v7321 int32
	_ = v7321
	var v7322 int32
	_ = v7322
	var v7326 int32
	_ = v7326
	var v7327 int32
	_ = v7327
	var v7328 int32
	_ = v7328
	var v7334 int32
	_ = v7334
	var v7337 int32
	_ = v7337
	var v7340 int32
	_ = v7340
	var v7341 int32
	_ = v7341
	var v7343 int32
	_ = v7343
	var v7351 int32
	_ = v7351
	var v7352 int32
	_ = v7352
	var v7354 int32
	_ = v7354
	var v7360 int32
	_ = v7360
	var v7366 int32
	_ = v7366
	var v7368 int32
	_ = v7368
	var v7370 int32
	_ = v7370
	var v7371 int32
	_ = v7371
	var v7372 int32
	_ = v7372
	var v7373 int32
	_ = v7373
	var v7375 int32
	_ = v7375
	var v7376 int32
	_ = v7376
	var v7378 int32
	_ = v7378
	var v7379 int32
	_ = v7379
	var v7381 int32
	_ = v7381
	var v7382 int32
	_ = v7382
	var v7383 int32
	_ = v7383
	var v7384 int32
	_ = v7384
	var v7394 int32
	_ = v7394
	var v7395 int32
	_ = v7395
	var v7427 int32
	_ = v7427
	var v7428 int32
	_ = v7428
	var v7429 int32
	_ = v7429
	var v7430 int32
	_ = v7430
	var v7431 int32
	_ = v7431
	var v7434 int32
	_ = v7434
	var v7437 int32
	_ = v7437
	var v7440 int32
	_ = v7440
	var v7448 int32
	_ = v7448
	var v7480 int32
	_ = v7480
	var v7484 int32
	_ = v7484
	var v7485 int32
	_ = v7485
	var v7486 int32
	_ = v7486
	var v7487 int32
	_ = v7487
	var v7488 int32
	_ = v7488
	var v7489 int32
	_ = v7489
	var v7491 int32
	_ = v7491
	var v7492 int32
	_ = v7492
	var v7495 int32
	_ = v7495
	var v7536 int32
	_ = v7536
	var v7537 int32
	_ = v7537
	var v7538 int32
	_ = v7538
	var v7539 int32
	_ = v7539
	var v7542 int32
	_ = v7542
	var v7546 int32
	_ = v7546
	var v7547 int32
	_ = v7547
	var v7548 int32
	_ = v7548
	var v7551 int32
	_ = v7551
	var v7552 int32
	_ = v7552
	var v7564 int32
	_ = v7564
	var v7596 int32
	_ = v7596
	var v7597 int32
	_ = v7597
	var v7600 int32
	_ = v7600
	var v7601 int32
	_ = v7601
	var v7607 int32
	_ = v7607
	var v7608 int32
	_ = v7608
	var v7651 int32
	_ = v7651
	var v7653 int32
	_ = v7653
	var v7659 int32
	_ = v7659
	var v7661 int32
	_ = v7661
	var v7697 int32
	_ = v7697
	var v7698 int32
	_ = v7698
	var v7702 int32
	_ = v7702
	var v7703 int32
	_ = v7703
	var v7706 int32
	_ = v7706
	var v7707 int32
	_ = v7707
	var v7711 int32
	_ = v7711
	var v7719 int32
	_ = v7719
	var v7751 int32
	_ = v7751
	var v7752 int32
	_ = v7752
	var v7755 int32
	_ = v7755
	var v7759 int32
	_ = v7759
	var v7760 int32
	_ = v7760
	var v7761 int32
	_ = v7761
	var v7763 int32
	_ = v7763
	var v7764 int32
	_ = v7764
	var v7775 int32
	_ = v7775
	var v7807 int32
	_ = v7807
	var v7808 int32
	_ = v7808
	var v7810 int32
	_ = v7810
	var v7811 int32
	_ = v7811
	var v7818 int32
	_ = v7818
	var v7856 int32
	_ = v7856
	var v7857 int32
	_ = v7857
	var v7858 int32
	_ = v7858
	var v7861 int32
	_ = v7861
	var v7862 int32
	_ = v7862
	var v7872 int32
	_ = v7872
	var v7875 int32
	_ = v7875
	var v7877 int32
	_ = v7877
	var v7879 int32
	_ = v7879
	var v7881 int32
	_ = v7881
	var v7884 int32
	_ = v7884
	var v7887 int32
	_ = v7887
	var v7893 int32
	_ = v7893
	var v7898 float64
	_ = v7898
	var v7902 int64
	_ = v7902
	var v7903 int64
	_ = v7903
	var v7905 int32
	_ = v7905
	var v7908 int32
	_ = v7908
	var v7909 int32
	_ = v7909
	var v7912 int32
	_ = v7912
	var v7917 int32
	_ = v7917
	var v7957 int32
	_ = v7957
	var v7958 int32
	_ = v7958
	var v7961 int32
	_ = v7961
	var v7962 int32
	_ = v7962
	var v7968 int32
	_ = v7968
	var v7969 int32
	_ = v7969
	var v8012 int32
	_ = v8012
	var v8013 int32
	_ = v8013
	var v8022 int32
	_ = v8022
	var v8027 int32
	_ = v8027
	var v8058 int32
	_ = v8058
	var v8062 int32
	_ = v8062
	var v8063 int32
	_ = v8063
	var v8067 int32
	_ = v8067
	var v8069 int32
	_ = v8069
	var v8073 int32
	_ = v8073
	var v8081 int32
	_ = v8081
	var v8113 int32
	_ = v8113
	var v8114 int32
	_ = v8114
	var v8117 int32
	_ = v8117
	var v8121 int32
	_ = v8121
	var v8122 int32
	_ = v8122
	var v8123 int32
	_ = v8123
	var v8125 int32
	_ = v8125
	var v8126 int32
	_ = v8126
	var v8137 int32
	_ = v8137
	var v8169 int32
	_ = v8169
	var v8170 int32
	_ = v8170
	var v8172 int32
	_ = v8172
	var v8173 int32
	_ = v8173
	var v8180 int32
	_ = v8180
	var v8216 int32
	_ = v8216
	var v8227 int32
	_ = v8227
	var v8230 int32
	_ = v8230
	var v8232 int32
	_ = v8232
	var v8236 int32
	_ = v8236
	var v8239 int32
	_ = v8239
	var v8242 int32
	_ = v8242
	var v8248 int32
	_ = v8248
	var v8253 float64
	_ = v8253
	var v8257 int64
	_ = v8257
	var v8258 int64
	_ = v8258
	var v8260 int32
	_ = v8260
	var v8264 int32
	_ = v8264
	var v8266 int32
	_ = v8266
	var v8268 int32
	_ = v8268
	var v8269 int32
	_ = v8269
	var v8270 int32
	_ = v8270
	var v8271 int32
	_ = v8271
	var v8272 int32
	_ = v8272
	var v8278 int32
	_ = v8278
	var v8279 int32
	_ = v8279
	var v8280 int32
	_ = v8280
	var v8282 int32
	_ = v8282
	var v8283 int32
	_ = v8283
	var v8285 int32
	_ = v8285
	var v8286 int32
	_ = v8286
	var v8287 int32
	_ = v8287
	var v8290 int32
	_ = v8290
	var v8291 int32
	_ = v8291
	var v8295 int32
	_ = v8295
	var v8299 int32
	_ = v8299
	var v8300 int32
	_ = v8300
	var v8302 int32
	_ = v8302
	var v8339 int32
	_ = v8339
	var v8343 int32
	_ = v8343
	var v8344 int32
	_ = v8344
	var v8347 int32
	_ = v8347
	var v8348 int32
	_ = v8348
	var v8349 int32
	_ = v8349
	var v8350 int32
	_ = v8350
	var v8352 int32
	_ = v8352
	var v8355 int32
	_ = v8355
	var v8356 int32
	_ = v8356
	var v8362 int32
	_ = v8362
	var v8450 int32
	_ = v8450
	var v8455 int32
	_ = v8455
	var v8456 int32
	_ = v8456
	var v8461 int32
	_ = v8461
	var v8462 int32
	_ = v8462
	var v8465 int32
	_ = v8465
	var v8468 int32
	_ = v8468
	var v8489 int32
	_ = v8489
	var v8509 int32
	_ = v8509
	var v8510 int32
	_ = v8510
	var v8513 int32
	_ = v8513
	var v8514 int32
	_ = v8514
	var v8517 int32
	_ = v8517
	var v8518 int32
	_ = v8518
	var v8519 int32
	_ = v8519
	var v8521 int32
	_ = v8521
	var v8526 int32
	_ = v8526
	var v8528 int32
	_ = v8528
	var v8532 int32
	_ = v8532
	var v8533 int32
	_ = v8533
	var v8538 int32
	_ = v8538
	var v8572 int32
	_ = v8572
	var v8576 int32
	_ = v8576
	var v8577 int32
	_ = v8577
	var v8580 int32
	_ = v8580
	var v8581 int32
	_ = v8581
	var v8582 int32
	_ = v8582
	var v8583 int32
	_ = v8583
	var v8585 int32
	_ = v8585
	var v8588 int32
	_ = v8588
	var v8589 int32
	_ = v8589
	var v8598 int32
	_ = v8598
	var v8674 int32
	_ = v8674
	var v8675 int32
	_ = v8675
	var v8676 int32
	_ = v8676
	var v8677 int32
	_ = v8677
	var v8679 int32
	_ = v8679
	var v8680 int32
	_ = v8680
	var v8683 int32
	_ = v8683
	var v8684 int32
	_ = v8684
	var v8687 int32
	_ = v8687
	var v8688 int32
	_ = v8688
	var v8728 int32
	_ = v8728
	var v8732 int32
	_ = v8732
	var v8733 int32
	_ = v8733
	var v8736 int32
	_ = v8736
	var v8738 int32
	_ = v8738
	var v8739 int32
	_ = v8739
	var v8740 int32
	_ = v8740
	var v8744 int32
	_ = v8744
	var v8748 int32
	_ = v8748
	var v8749 int32
	_ = v8749
	var v8750 int32
	_ = v8750
	var v8751 int32
	_ = v8751
	var v8752 int32
	_ = v8752
	var v8754 int32
	_ = v8754
	var v8755 int32
	_ = v8755
	var v8757 int32
	_ = v8757
	var v8799 int32
	_ = v8799
	var v8801 int32
	_ = v8801
	var v8802 int32
	_ = v8802
	var v8807 int32
	_ = v8807
	var v8811 int32
	_ = v8811
	var v8816 int32
	_ = v8816
	var v8817 int32
	_ = v8817
	var v8858 int32
	_ = v8858
	var v8860 int32
	_ = v8860
	var v8861 int32
	_ = v8861
	var v8864 int32
	_ = v8864
	var v8865 int32
	_ = v8865
	var v8868 int32
	_ = v8868
	var v8869 int32
	_ = v8869
	var v8909 int32
	_ = v8909
	var v8913 int32
	_ = v8913
	var v8914 int32
	_ = v8914
	var v8917 int32
	_ = v8917
	var v8919 int32
	_ = v8919
	var v8920 int32
	_ = v8920
	var v8921 int32
	_ = v8921
	var v8925 int32
	_ = v8925
	var v8929 int32
	_ = v8929
	var v8930 int32
	_ = v8930
	var v8931 int32
	_ = v8931
	var v8932 int32
	_ = v8932
	var v8933 int32
	_ = v8933
	var v8935 int32
	_ = v8935
	var v8936 int32
	_ = v8936
	var v8938 int32
	_ = v8938
	var v8979 int32
	_ = v8979
	var v8982 int32
	_ = v8982
	var v8986 int32
	_ = v8986
	var v8988 int32
	_ = v8988
	var v9027 int32
	_ = v9027
	var v9031 int32
	_ = v9031
	var v9032 int32
	_ = v9032
	var v9033 int32
	_ = v9033
	var v9035 int32
	_ = v9035
	var v9038 int32
	_ = v9038
	var v9041 int32
	_ = v9041
	var v9043 int32
	_ = v9043
	var v9044 int32
	_ = v9044
	var v9045 int32
	_ = v9045
	var v9049 int32
	_ = v9049
	var v9053 int32
	_ = v9053
	var v9054 int32
	_ = v9054
	var v9055 int32
	_ = v9055
	var v9059 int32
	_ = v9059
	var v9063 int32
	_ = v9063
	var v9064 int32
	_ = v9064
	var v9066 int32
	_ = v9066
	var v9067 int32
	_ = v9067
	var v9068 int32
	_ = v9068
	var v9069 int32
	_ = v9069
	var v9070 int32
	_ = v9070
	var v9071 int32
	_ = v9071
	var v9073 int32
	_ = v9073
	var v9076 int32
	_ = v9076
	var v9077 int32
	_ = v9077
	var v9083 int32
	_ = v9083
	var v9084 int32
	_ = v9084
	var v9086 int32
	_ = v9086
	var v9087 int32
	_ = v9087
	var v9088 int32
	_ = v9088
	var v9096 int32
	_ = v9096
	var v9097 int32
	_ = v9097
	var v9098 int32
	_ = v9098
	var v9102 int32
	_ = v9102
	var v9106 int32
	_ = v9106
	var v9107 int32
	_ = v9107
	var v9109 int32
	_ = v9109
	var v9110 int32
	_ = v9110
	var v9111 int32
	_ = v9111
	var v9112 int32
	_ = v9112
	var v9113 int32
	_ = v9113
	var v9115 int32
	_ = v9115
	var v9118 int32
	_ = v9118
	var v9122 int32
	_ = v9122
	var v9124 int32
	_ = v9124
	var v9125 int32
	_ = v9125
	var v9126 int32
	_ = v9126
	var v9132 int32
	_ = v9132
	var v9133 int32
	_ = v9133
	var v9134 int32
	_ = v9134
	var v9138 int32
	_ = v9138
	var v9142 int32
	_ = v9142
	var v9143 int32
	_ = v9143
	var v9145 int32
	_ = v9145
	var v9146 int32
	_ = v9146
	var v9147 int32
	_ = v9147
	var v9148 int32
	_ = v9148
	var v9149 int32
	_ = v9149
	var v9153 int32
	_ = v9153
	var v9154 int32
	_ = v9154
	var v9156 int32
	_ = v9156
	var v9197 int32
	_ = v9197
	var v9200 int32
	_ = v9200
	var v9203 int32
	_ = v9203
	var v9207 int32
	_ = v9207
	var v9208 int32
	_ = v9208
	var v9211 int32
	_ = v9211
	var v9215 int32
	_ = v9215
	var v9216 int32
	_ = v9216
	var v9256 int32
	_ = v9256
	var v9260 int32
	_ = v9260
	var v9261 int32
	_ = v9261
	var v9264 int32
	_ = v9264
	var v9266 int32
	_ = v9266
	var v9267 int32
	_ = v9267
	var v9268 int32
	_ = v9268
	var v9272 int32
	_ = v9272
	var v9276 int32
	_ = v9276
	var v9277 int32
	_ = v9277
	var v9278 int32
	_ = v9278
	var v9279 int32
	_ = v9279
	var v9280 int32
	_ = v9280
	var v9282 int32
	_ = v9282
	var v9283 int32
	_ = v9283
	var v9285 int32
	_ = v9285
	var v9327 int32
	_ = v9327
	var v9328 int32
	_ = v9328
	var v9372 int32
	_ = v9372
	var v9376 int32
	_ = v9376
	var v9379 int32
	_ = v9379
	var v9381 int32
	_ = v9381
	var v9382 int32
	_ = v9382
	var v9384 int32
	_ = v9384
	var v9385 int32
	_ = v9385
	var v9387 int32
	_ = v9387
	var v9390 int32
	_ = v9390
	var v9391 int32
	_ = v9391
	var v9392 int32
	_ = v9392
	var v9394 int32
	_ = v9394
	var v9396 int32
	_ = v9396
	var v9397 int32
	_ = v9397
	var v9405 int32
	_ = v9405
	var v9406 int32
	_ = v9406
	var v9408 int32
	_ = v9408
	var v9409 int32
	_ = v9409
	var v9410 int32
	_ = v9410
	var v9413 int32
	_ = v9413
	var v9414 int32
	_ = v9414
	var v9415 int32
	_ = v9415
	var v9421 int64
	_ = v9421
	var v9424 int32
	_ = v9424
	var v9463 int32
	_ = v9463
	var v9464 int32
	_ = v9464
	var v9465 int32
	_ = v9465
	var v9468 int32
	_ = v9468
	var v9469 int32
	_ = v9469
	var v9473 int32
	_ = v9473
	var v9476 int32
	_ = v9476
	var v9480 int32
	_ = v9480
	var v9481 int32
	_ = v9481
	var v9482 int32
	_ = v9482
	var v9483 int32
	_ = v9483
	var v9484 int32
	_ = v9484
	var v9491 int32
	_ = v9491
	var v9494 int64
	_ = v9494
	var v9495 int32
	_ = v9495
	var v9496 int32
	_ = v9496
	var v9499 int32
	_ = v9499
	var v9501 int32
	_ = v9501
	var v9504 int32
	_ = v9504
	var v9545 int32
	_ = v9545
	var v9549 int32
	_ = v9549
	var v9550 int32
	_ = v9550
	var v9551 int32
	_ = v9551
	var v9552 int32
	_ = v9552
	var v9559 int32
	_ = v9559
	var v9562 int64
	_ = v9562
	var v9563 int32
	_ = v9563
	var v9564 int32
	_ = v9564
	var v9567 int32
	_ = v9567
	var v9570 int32
	_ = v9570
	var v9571 int32
	_ = v9571
	var v9583 int32
	_ = v9583
	var v9614 int32
	_ = v9614
	var v9617 int32
	_ = v9617
	var v9637 int32
	_ = v9637
	var v9662 int32
	_ = v9662
	var v9666 int32
	_ = v9666
	var v9668 int32
	_ = v9668
	var v9669 int32
	_ = v9669
	var v9670 int32
	_ = v9670
	var v9671 int32
	_ = v9671
	var v9674 int32
	_ = v9674
	var v9675 int32
	_ = v9675
	var v9676 int32
	_ = v9676
	var v9677 int32
	_ = v9677
	var v9680 int32
	_ = v9680
	var v9681 int32
	_ = v9681
	var v9683 int32
	_ = v9683
	var v9684 int32
	_ = v9684
	var v9685 int32
	_ = v9685
	var v9686 int32
	_ = v9686
	var v9689 int32
	_ = v9689
	var v9690 int32
	_ = v9690
	var v9691 int32
	_ = v9691
	var v9692 int32
	_ = v9692
	var v9695 int32
	_ = v9695
	var v9696 int32
	_ = v9696
	var v9700 int32
	_ = v9700
	var v9704 int32
	_ = v9704
	var v9707 int32
	_ = v9707
	var v9725 int32
	_ = v9725
	var v9750 int32
	_ = v9750
	var v9754 int32
	_ = v9754
	var v9757 int32
	_ = v9757
	var v9758 int32
	_ = v9758
	var v9760 int32
	_ = v9760
	var v9761 int32
	_ = v9761
	var v9765 int32
	_ = v9765
	var v9776 int32
	_ = v9776
	var v9782 int32
	_ = v9782
	var v9807 int32
	_ = v9807
	var v9811 int32
	_ = v9811
	var v9812 int32
	_ = v9812
	var v9813 int32
	_ = v9813
	var v9814 int32
	_ = v9814
	var v9815 int32
	_ = v9815
	var v9816 int32
	_ = v9816
	var v9820 int32
	_ = v9820
	var v9821 int32
	_ = v9821
	var v9828 int32
	_ = v9828
	var v9829 int32
	_ = v9829
	var v9873 int32
	_ = v9873
	var v9874 int32
	_ = v9874
	var v9876 int32
	_ = v9876
	var v9877 int32
	_ = v9877
	var v9921 int32
	_ = v9921
	var v9922 int32
	_ = v9922
	var v9932 int32
	_ = v9932
	var v9965 int32
	_ = v9965
	var v9966 int32
	_ = v9966
	var v9967 int32
	_ = v9967
	var v9968 int32
	_ = v9968
	var v9971 int32
	_ = v9971
	var v9973 int32
	_ = v9973
	var v9974 int32
	_ = v9974
	var v9975 int32
	_ = v9975
	var v10014 int32
	_ = v10014
	var v10015 int32
	_ = v10015
	var v10016 int32
	_ = v10016
	var v10019 int32
	_ = v10019
	var v10020 int32
	_ = v10020
	var v10024 int32
	_ = v10024
	var v10027 int32
	_ = v10027
	var v10029 int32
	_ = v10029
	var v10030 int32
	_ = v10030
	var v10031 int32
	_ = v10031
	var v10032 int32
	_ = v10032
	var v10033 int32
	_ = v10033
	var v10034 int32
	_ = v10034
	var v10036 int32
	_ = v10036
	var v10039 int32
	_ = v10039
	var v10040 int32
	_ = v10040
	var v10042 int32
	_ = v10042
	var v10047 int32
	_ = v10047
	var v10048 int32
	_ = v10048
	var v10052 int32
	_ = v10052
	var v10069 int32
	_ = v10069
	var v10095 int32
	_ = v10095
	var v10096 int32
	_ = v10096
	var v10097 int32
	_ = v10097
	var v10099 int32
	_ = v10099
	var v10102 int32
	_ = v10102
	var v10107 int32
	_ = v10107
	var v10112 int32
	_ = v10112
	var v10113 int32
	_ = v10113
	var v10121 int32
	_ = v10121
	var v10152 int32
	_ = v10152
	var v10156 int32
	_ = v10156
	var v10157 int32
	_ = v10157
	var v10167 int32
	_ = v10167
	var v10168 int32
	_ = v10168
	var v10170 int32
	_ = v10170
	var v10171 int32
	_ = v10171
	var v10174 int32
	_ = v10174
	var v10184 int32
	_ = v10184
	var v10215 int32
	_ = v10215
	var v10218 int32
	_ = v10218
	var v10261 int32
	_ = v10261
	var v10262 int32
	_ = v10262
	var v10264 int32
	_ = v10264
	var v10267 int32
	_ = v10267
	var v10270 int32
	_ = v10270
	var v10273 int32
	_ = v10273
	var v10274 int32
	_ = v10274
	var v10277 int32
	_ = v10277
	var v10278 int32
	_ = v10278
	var v10281 int32
	_ = v10281
	var v10288 int32
	_ = v10288
	var v10289 int32
	_ = v10289
	var v10294 int32
	_ = v10294
	var v10303 int32
	_ = v10303
	var v10304 int32
	_ = v10304
	var v10306 int32
	_ = v10306
	var v10307 int32
	_ = v10307
	var v10320 int32
	_ = v10320
	var v10353 int32
	_ = v10353
	var v10354 int32
	_ = v10354
	var v10356 int32
	_ = v10356
	var v10358 int32
	_ = v10358
	var v10367 int32
	_ = v10367
	var v10399 int32
	_ = v10399
	var v10401 int32
	_ = v10401
	var v10413 int32
	_ = v10413
	var v10448 int32
	_ = v10448
	var v10459 int32
	_ = v10459
	var v10491 int32
	_ = v10491
	var v10510 int32
	_ = v10510
	var v10514 int32
	_ = v10514
	var v10533 int32
	_ = v10533
	var v10536 int32
	_ = v10536
	var v10538 float64
	_ = v10538
	var v10539 int32
	_ = v10539
	var v10541 int32
	_ = v10541
	var v10543 int32
	_ = v10543
	var v10544 int32
	_ = v10544
	var v10547 int32
	_ = v10547
	var v10548 int32
	_ = v10548
	var v10549 int32
	_ = v10549
	var v10552 int32
	_ = v10552
	var v10553 int32
	_ = v10553
	var v10556 int32
	_ = v10556
	var v10597 int32
	_ = v10597
	var v10598 int32
	_ = v10598
	var v10603 int32
	_ = v10603
	var v10606 int32
	_ = v10606
	var v10610 int32
	_ = v10610
	var v10613 int32
	_ = v10613
	var v10616 int32
	_ = v10616
	var v10617 int32
	_ = v10617
	var v10618 int32
	_ = v10618
	var v10619 int32
	_ = v10619
	var v10625 int32
	_ = v10625
	var v10626 int32
	_ = v10626
	var v10627 int32
	_ = v10627
	var v10628 int32
	_ = v10628
	var v10631 int32
	_ = v10631
	var v10634 int32
	_ = v10634
	var v10638 int32
	_ = v10638
	var v10640 int32
	_ = v10640
	var v10655 int32
	_ = v10655
	var v10684 int32
	_ = v10684
	var v10685 int32
	_ = v10685
	var v10689 int32
	_ = v10689
	var v10690 int32
	_ = v10690
	var v10691 int32
	_ = v10691
	var v10692 int32
	_ = v10692
	var v10693 int32
	_ = v10693
	var v10697 int32
	_ = v10697
	var v10701 int32
	_ = v10701
	var v10703 int32
	_ = v10703
	var v10704 int32
	_ = v10704
	var v10706 int64
	_ = v10706
	var v10707 int32
	_ = v10707
	var v10708 int32
	_ = v10708
	var v10711 int32
	_ = v10711
	var v10712 int32
	_ = v10712
	var v10714 int32
	_ = v10714
	var v10716 int32
	_ = v10716
	var v10719 int32
	_ = v10719
	var v10720 int32
	_ = v10720
	var v10721 int32
	_ = v10721
	var v10722 int32
	_ = v10722
	var v10723 int32
	_ = v10723
	var v10724 int32
	_ = v10724
	var v10725 int32
	_ = v10725
	var v10726 int32
	_ = v10726
	var v10727 int32
	_ = v10727
	var v10728 int32
	_ = v10728
	var v10729 int32
	_ = v10729
	var v10731 int32
	_ = v10731
	var v10732 int32
	_ = v10732
	var v10735 int32
	_ = v10735
	var v10738 int32
	_ = v10738
	var v10739 int64
	_ = v10739
	var v10746 int32
	_ = v10746
	var v10747 int32
	_ = v10747
	var v10748 int32
	_ = v10748
	var v10750 int32
	_ = v10750
	var v10752 int32
	_ = v10752
	var v10753 int32
	_ = v10753
	var v10760 int32
	_ = v10760
	var v10837 int32
	_ = v10837
	var v10840 int32
	_ = v10840
	var v10843 int32
	_ = v10843
	var v10850 int32
	_ = v10850
	var v10887 int32
	_ = v10887
	var v10891 int32
	_ = v10891
	var v10892 int32
	_ = v10892
	var v10895 int32
	_ = v10895
	var v10896 int32
	_ = v10896
	var v10899 int32
	_ = v10899
	var v10900 int32
	_ = v10900
	var v10901 int32
	_ = v10901
	var v10902 int32
	_ = v10902
	var v10905 int32
	_ = v10905
	var v10906 int32
	_ = v10906
	var v10909 int32
	_ = v10909
	var v10910 int32
	_ = v10910
	var v10915 int32
	_ = v10915
	var v10916 int32
	_ = v10916
	var v10959 int32
	_ = v10959
	var v10967 int32
	_ = v10967
	var v11004 int32
	_ = v11004
	var v11008 int32
	_ = v11008
	var v11009 int32
	_ = v11009
	var v11010 int32
	_ = v11010
	var v11011 int32
	_ = v11011
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
	var v11020 int32
	_ = v11020
	var v11021 int32
	_ = v11021
	var v11066 int32
	_ = v11066
	var v11067 int32
	_ = v11067
	var v11068 int32
	_ = v11068
	var v11069 int32
	_ = v11069
	var v11070 int32
	_ = v11070
	var v11071 int32
	_ = v11071
	var v11072 int32
	_ = v11072
	var v11073 int32
	_ = v11073
	var v11074 int32
	_ = v11074
	var v11075 int32
	_ = v11075
	var v11077 int32
	_ = v11077
	var v11080 int32
	_ = v11080
	var v11081 int32
	_ = v11081
	var v11084 int32
	_ = v11084
	var v11090 int32
	_ = v11090
	var v11101 int32
	_ = v11101
	var v11102 int32
	_ = v11102
	var v11106 int32
	_ = v11106
	var v11111 int32
	_ = v11111
	var v11112 float64
	_ = v11112
	var v11116 int32
	_ = v11116
	var v11147 int32
	_ = v11147
	var v11151 int32
	_ = v11151
	var v11152 float64
	_ = v11152
	var v11153 int32
	_ = v11153
	var v11154 int32
	_ = v11154
	var v11155 int32
	_ = v11155
	var v11158 int32
	_ = v11158
	var v11161 int32
	_ = v11161
	var v11162 float64
	_ = v11162
	var v11163 int32
	_ = v11163
	var v11165 int32
	_ = v11165
	var v11166 int32
	_ = v11166
	var v11173 int32
	_ = v11173
	var v11174 float64
	_ = v11174
	var v11178 int32
	_ = v11178
	var v11210 float64
	_ = v11210
	var v11213 float64
	_ = v11213
	var v11215 float64
	_ = v11215
	var v11218 float64
	_ = v11218
	var v11222 int32
	_ = v11222
	var v11223 float64
	_ = v11223
	var v11224 float64
	_ = v11224
	var v11227 float64
	_ = v11227
	var v11228 float64
	_ = v11228
	var v11232 int32
	_ = v11232
	var v11236 int32
	_ = v11236
	var v11237 int32
	_ = v11237
	var v11238 int32
	_ = v11238
	var v11239 int32
	_ = v11239
	var v11240 int32
	_ = v11240
	var v11242 int32
	_ = v11242
	var v11249 int32
	_ = v11249
	var v11297 int32
	_ = v11297
	var v11298 int32
	_ = v11298
	var v11302 int32
	_ = v11302
	var v11307 int32
	_ = v11307
	var v11349 float64
	_ = v11349
	var v11350 int32
	_ = v11350
	var v11351 int32
	_ = v11351
	var v11352 int32
	_ = v11352
	var v11353 int32
	_ = v11353
	var v11354 int32
	_ = v11354
	var v11355 int32
	_ = v11355
	var v11357 int32
	_ = v11357
	var v11358 float64
	_ = v11358
	var v11359 float64
	_ = v11359
	var v11367 int32
	_ = v11367
	var v11368 int32
	_ = v11368
	var v11369 int32
	_ = v11369
	var v11370 int32
	_ = v11370
	var v11371 int32
	_ = v11371
	var v11372 int32
	_ = v11372
	var v11373 int32
	_ = v11373
	var v11374 int32
	_ = v11374
	var v11375 int32
	_ = v11375
	var v11376 int32
	_ = v11376
	var v11377 int32
	_ = v11377
	var v11378 int32
	_ = v11378
	var v11379 int32
	_ = v11379
	var v11380 int32
	_ = v11380
	var v11382 int32
	_ = v11382
	var v11383 int32
	_ = v11383
	var v11384 int32
	_ = v11384
	var v11385 int32
	_ = v11385
	var v11386 int32
	_ = v11386
	var v11387 int32
	_ = v11387
	var v11390 int32
	_ = v11390
	var v11393 int32
	_ = v11393
	var v11399 int32
	_ = v11399
	var v11405 int32
	_ = v11405
	var v11408 int32
	_ = v11408
	var v11412 int32
	_ = v11412
	var v11413 int32
	_ = v11413
	var v11440 int32
	_ = v11440
	var v11441 int32
	_ = v11441
	var v11443 int32
	_ = v11443
	var v11444 int32
	_ = v11444
	var v11446 int32
	_ = v11446
	var v11447 int32
	_ = v11447
	var v11450 int32
	_ = v11450
	var v11451 int32
	_ = v11451
	var v11454 int32
	_ = v11454
	var v11458 int32
	_ = v11458
	var v11459 int32
	_ = v11459
	var v11460 int32
	_ = v11460
	var v11467 int32
	_ = v11467
	var v11468 float64
	_ = v11468
	var v11470 float64
	_ = v11470
	var v11476 int32
	_ = v11476
	var v11484 int32
	_ = v11484
	var v11487 int32
	_ = v11487
	var v11488 int32
	_ = v11488
	var v11489 int32
	_ = v11489
	var v11490 int32
	_ = v11490
	var v11491 int32
	_ = v11491
	var v11492 int32
	_ = v11492
	var v11494 int32
	_ = v11494
	var v11495 int32
	_ = v11495
	var v11497 int32
	_ = v11497
	var v11499 int32
	_ = v11499
	var v11507 int32
	_ = v11507
	var v11508 float64
	_ = v11508
	var v11513 int32
	_ = v11513
	var v11514 int32
	_ = v11514
	var v11515 int32
	_ = v11515
	var v11519 int32
	_ = v11519
	var v11520 int32
	_ = v11520
	var v11524 int32
	_ = v11524
	var v11525 int32
	_ = v11525
	var v11566 int32
	_ = v11566
	var v11567 int32
	_ = v11567
	var v11569 int32
	_ = v11569
	var v11571 int32
	_ = v11571
	var v11579 int32
	_ = v11579
	var v11582 int32
	_ = v11582
	var v11583 int32
	_ = v11583
	var v11584 int32
	_ = v11584
	var v11586 int32
	_ = v11586
	var v11588 int32
	_ = v11588
	var v11590 int32
	_ = v11590
	var v11591 int32
	_ = v11591
	var v11594 int32
	_ = v11594
	var v11595 int32
	_ = v11595
	var v11597 int32
	_ = v11597
	var v11639 int32
	_ = v11639
	var v11640 int32
	_ = v11640
	var v11642 int32
	_ = v11642
	var v11644 int32
	_ = v11644
	var v11646 int32
	_ = v11646
	var v11650 float64
	_ = v11650
	var v11651 int32
	_ = v11651
	var v11652 int32
	_ = v11652
	var v11659 float64
	_ = v11659
	var v11694 int32
	_ = v11694
	var v11695 int32
	_ = v11695
	var v11696 int32
	_ = v11696
	var v11697 int32
	_ = v11697
	var v11703 int32
	_ = v11703
	var v11704 float64
	_ = v11704
	var v11717 int32
	_ = v11717
	var v11739 int32
	_ = v11739
	var v11740 int32
	_ = v11740
	var v11745 int32
	_ = v11745
	var v11746 int32
	_ = v11746
	var v11785 int32
	_ = v11785
	var v11789 int32
	_ = v11789
	var v11790 int32
	_ = v11790
	var v11793 int32
	_ = v11793
	var v11794 int32
	_ = v11794
	var v11798 int32
	_ = v11798
	var v11806 int32
	_ = v11806
	var v11838 int32
	_ = v11838
	var v11842 int32
	_ = v11842
	var v11843 int32
	_ = v11843
	var v11844 int32
	_ = v11844
	var v11845 int32
	_ = v11845
	var v11847 int32
	_ = v11847
	var v11848 int32
	_ = v11848
	var v11851 int32
	_ = v11851
	var v11891 int32
	_ = v11891
	var v11894 int32
	_ = v11894
	var v11895 int32
	_ = v11895
	var v11899 int32
	_ = v11899
	var v11907 int32
	_ = v11907
	var v11939 int32
	_ = v11939
	var v11943 int32
	_ = v11943
	var v11944 int32
	_ = v11944
	var v11945 int32
	_ = v11945
	var v11946 int32
	_ = v11946
	var v11948 int32
	_ = v11948
	var v11949 int32
	_ = v11949
	var v11952 int32
	_ = v11952
	var v11993 int32
	_ = v11993
	var v11994 int32
	_ = v11994
	var v11997 int32
	_ = v11997
	var v12037 int32
	_ = v12037
	var v12040 int32
	_ = v12040
	var v12041 int32
	_ = v12041
	var v12045 int32
	_ = v12045
	var v12053 int32
	_ = v12053
	var v12085 int32
	_ = v12085
	var v12089 int32
	_ = v12089
	var v12090 int32
	_ = v12090
	var v12091 int32
	_ = v12091
	var v12092 int32
	_ = v12092
	var v12094 int32
	_ = v12094
	var v12095 int32
	_ = v12095
	var v12098 int32
	_ = v12098
	var v12138 int32
	_ = v12138
	var v12139 int32
	_ = v12139
	var v12140 int32
	_ = v12140
	var v12144 int32
	_ = v12144
	var v12145 int32
	_ = v12145
	var v12151 int32
	_ = v12151
	var v12158 int32
	_ = v12158
	var v12191 int32
	_ = v12191
	var v12192 int32
	_ = v12192
	var v12194 int32
	_ = v12194
	var v12195 int32
	_ = v12195
	var v12199 int32
	_ = v12199
	var v12202 int32
	_ = v12202
	var v12203 int32
	_ = v12203
	var v12207 int32
	_ = v12207
	var v12209 int32
	_ = v12209
	var v12210 int32
	_ = v12210
	var v12211 int32
	_ = v12211
	var v12214 int32
	_ = v12214
	var v12215 int32
	_ = v12215
	var v12219 int32
	_ = v12219
	var v12259 int32
	_ = v12259
	var v12260 int32
	_ = v12260
	var v12262 int32
	_ = v12262
	var v12264 int32
	_ = v12264
	var v12266 int32
	_ = v12266
	var v12267 int32
	_ = v12267
	var v12268 int32
	_ = v12268
	var v12269 int32
	_ = v12269
	var v12270 int32
	_ = v12270
	var v12271 int32
	_ = v12271
	var v12284 int32
	_ = v12284
	var v12286 int32
	_ = v12286
	var v12313 int32
	_ = v12313
	var v12314 int32
	_ = v12314
	var v12315 int32
	_ = v12315
	var v12317 int32
	_ = v12317
	var v12322 int32
	_ = v12322
	var v12323 int32
	_ = v12323
	var v12324 int32
	_ = v12324
	var v12325 int32
	_ = v12325
	var v12329 int32
	_ = v12329
	var v12330 int32
	_ = v12330
	var v12334 int32
	_ = v12334
	var v12335 int32
	_ = v12335
	var v12376 int32
	_ = v12376
	var v12377 int32
	_ = v12377
	var v12379 int32
	_ = v12379
	var v12380 int32
	_ = v12380
	var v12384 int32
	_ = v12384
	var v12387 int32
	_ = v12387
	var v12393 int32
	_ = v12393
	var v12396 int32
	_ = v12396
	var v12399 int32
	_ = v12399
	var v12400 int32
	_ = v12400
	var v12403 int32
	_ = v12403
	var v12410 int32
	_ = v12410
	var v12411 int32
	_ = v12411
	var v12414 int32
	_ = v12414
	var v12425 int32
	_ = v12425
	var v12429 int32
	_ = v12429
	var v12432 int32
	_ = v12432
	var v12435 int32
	_ = v12435
	var v12436 int32
	_ = v12436
	var v12437 int32
	_ = v12437
	var v12439 int32
	_ = v12439
	var v12440 int32
	_ = v12440
	var v12441 int32
	_ = v12441
	var v12443 int32
	_ = v12443
	var v12446 int32
	_ = v12446
	var v12447 int32
	_ = v12447
	var v12448 int32
	_ = v12448
	var v12453 int32
	_ = v12453
	var v12454 int32
	_ = v12454
	var v12456 int32
	_ = v12456
	var v12497 int32
	_ = v12497
	var v12498 int32
	_ = v12498
	var v12499 int32
	_ = v12499
	var v12500 int32
	_ = v12500
	var v12502 int32
	_ = v12502
	var v12503 int32
	_ = v12503
	var v12504 int32
	_ = v12504
	var v12505 int32
	_ = v12505
	var v12508 int32
	_ = v12508
	var v12511 int32
	_ = v12511
	var v12512 int32
	_ = v12512
	var v12513 int32
	_ = v12513
	var v12515 int32
	_ = v12515
	var v12516 int32
	_ = v12516
	var v12517 int32
	_ = v12517
	var v12519 int32
	_ = v12519
	var v12521 int32
	_ = v12521
	var v12523 int32
	_ = v12523
	var v12524 int32
	_ = v12524
	var v12525 int32
	_ = v12525
	var v12526 int32
	_ = v12526
	var v12527 int32
	_ = v12527
	var v12528 int32
	_ = v12528
	var v12530 int32
	_ = v12530
	var v12531 int32
	_ = v12531
	var v12570 int32
	_ = v12570
	var v12571 int32
	_ = v12571
	var v12579 int32
	_ = v12579
	var v12580 int32
	_ = v12580
	var v12581 int32
	_ = v12581
	var v12582 int32
	_ = v12582
	var v12588 int32
	_ = v12588
	var v12589 int32
	_ = v12589
	var v12590 int32
	_ = v12590
	var v12591 int32
	_ = v12591
	var v12598 int32
	_ = v12598
	var v12599 int32
	_ = v12599
	var v12600 int32
	_ = v12600
	var v12601 int32
	_ = v12601
	var v12608 int32
	_ = v12608
	var v12609 int32
	_ = v12609
	var v12610 int32
	_ = v12610
	var v12611 int32
	_ = v12611
	var v12612 int32
	_ = v12612
	var v12630 int32
	_ = v12630
	var v12631 int32
	_ = v12631
	var v12637 int32
	_ = v12637
	var v12638 int32
	_ = v12638
	var v12639 int32
	_ = v12639
	var v12640 int32
	_ = v12640
	var v12641 int32
	_ = v12641
	var v12643 int32
	_ = v12643
	var v12646 int32
	_ = v12646
	var v12647 int32
	_ = v12647
	var v12648 int32
	_ = v12648
	var v12649 int32
	_ = v12649
	var v12650 int32
	_ = v12650
	var v12651 int32
	_ = v12651
	var v12652 int32
	_ = v12652
	var v12654 int32
	_ = v12654
	var v12656 int32
	_ = v12656
	var v12657 int32
	_ = v12657
	var v12658 int32
	_ = v12658
	var v12659 int32
	_ = v12659
	var v12661 int32
	_ = v12661
	var v12670 int32
	_ = v12670
	var v12671 int64
	_ = v12671
	var v12685 int32
	_ = v12685
	var v12686 int32
	_ = v12686
	var v12687 int32
	_ = v12687
	var v12694 int32
	_ = v12694
	var v12700 int32
	_ = v12700
	var v12701 int32
	_ = v12701
	var v12702 int32
	_ = v12702
	var v12707 int32
	_ = v12707
	var v12708 int32
	_ = v12708
	var v12709 int32
	_ = v12709
	var v12711 int32
	_ = v12711
	var v12715 int32
	_ = v12715
	var v12716 int32
	_ = v12716
	var v12719 int32
	_ = v12719
	var v12721 int64
	_ = v12721
	var v12723 int32
	_ = v12723
	var v12725 int32
	_ = v12725
	var v12727 int32
	_ = v12727
	var v12729 int32
	_ = v12729
	var v12731 int32
	_ = v12731
	var v12732 int32
	_ = v12732
	var v12735 int32
	_ = v12735
	var v12738 int32
	_ = v12738
	var v12739 int32
	_ = v12739
	var v12740 int32
	_ = v12740
	var v12743 int32
	_ = v12743
	var v12750 int32
	_ = v12750
	var v12765 int32
	_ = v12765
	var v12792 int32
	_ = v12792
	var v12793 int32
	_ = v12793
	var v12794 int32
	_ = v12794
	var v12795 int32
	_ = v12795
	var v12796 int32
	_ = v12796
	var v12797 int32
	_ = v12797
	var v12801 int32
	_ = v12801
	var v12803 int64
	_ = v12803
	var v12807 int32
	_ = v12807
	var v12812 int32
	_ = v12812
	var v12813 int32
	_ = v12813
	var v12815 int32
	_ = v12815
	var v12817 int32
	_ = v12817
	var v12818 int32
	_ = v12818
	var v12819 int32
	_ = v12819
	var v12820 int32
	_ = v12820
	var v12822 int32
	_ = v12822
	var v12823 int32
	_ = v12823
	var v12824 int32
	_ = v12824
	var v12825 int32
	_ = v12825
	var v12834 int32
	_ = v12834
	var v12837 int32
	_ = v12837
	var v12840 int32
	_ = v12840
	var v12841 int32
	_ = v12841
	var v12843 int32
	_ = v12843
	var v12851 int32
	_ = v12851
	var v12852 int32
	_ = v12852
	var v12853 int32
	_ = v12853
	var v12854 int32
	_ = v12854
	var v12858 int32
	_ = v12858
	var v12862 int32
	_ = v12862
	var v12870 int32
	_ = v12870
	var v12875 int32
	_ = v12875
	var v12876 int32
	_ = v12876
	var v12879 int32
	_ = v12879
	var v12880 int32
	_ = v12880
	var v12881 int32
	_ = v12881
	var v12882 int32
	_ = v12882
	var v12889 int32
	_ = v12889
	var v12892 int32
	_ = v12892
	var v12895 int32
	_ = v12895
	var v12896 int32
	_ = v12896
	var v12900 int32
	_ = v12900
	var v12904 int32
	_ = v12904
	var v12905 int32
	_ = v12905
	var v12909 int32
	_ = v12909
	var v12913 int32
	_ = v12913
	var v12919 int32
	_ = v12919
	var v12924 int32
	_ = v12924
	var v12925 int32
	_ = v12925
	var v12926 int32
	_ = v12926
	var v12929 int32
	_ = v12929
	var v12932 int32
	_ = v12932
	var v12933 int32
	_ = v12933
	var v12936 int32
	_ = v12936
	var v12937 int32
	_ = v12937
	var v12938 int32
	_ = v12938
	var v12941 int32
	_ = v12941
	var v12943 int32
	_ = v12943
	var v12945 int32
	_ = v12945
	var v12947 int32
	_ = v12947
	var v12949 int32
	_ = v12949
	var v12952 int32
	_ = v12952
	var v12956 int32
	_ = v12956
	var v12965 int32
	_ = v12965
	var v13008 int32
	_ = v13008
	var v13009 int32
	_ = v13009
	var v13012 int32
	_ = v13012
	var v13013 int32
	_ = v13013
	var v13015 int32
	_ = v13015
	var v13019 int32
	_ = v13019
	var v13061 int32
	_ = v13061
	var v13062 int32
	_ = v13062
	var v13063 int32
	_ = v13063
	var v13067 int32
	_ = v13067
	var v13068 int32
	_ = v13068
	var v13071 int32
	_ = v13071
	var v13073 int32
	_ = v13073
	var v13075 int32
	_ = v13075
	var v13077 int32
	_ = v13077
	var v13079 int32
	_ = v13079
	var v13081 int32
	_ = v13081
	var v13084 int32
	_ = v13084
	var v13116 int32
	_ = v13116
	var v13128 int32
	_ = v13128
	var v13132 int32
	_ = v13132
	var v13133 int32
	_ = v13133
	var v13135 int32
	_ = v13135
	var v13136 int32
	_ = v13136
	var v13138 int32
	_ = v13138
	var v13156 int32
	_ = v13156
	var v13159 int32
	_ = v13159
	var v13160 int32
	_ = v13160
	var v13163 int32
	_ = v13163
	var v13164 int32
	_ = v13164
	var v13169 int32
	_ = v13169
	var v13176 int32
	_ = v13176
	var v13180 int32
	_ = v13180
	var v13186 int32
	_ = v13186
	var v13190 int32
	_ = v13190
	var v13194 int32
	_ = v13194
	var v13198 int32
	_ = v13198
	var v13204 int32
	_ = v13204
	var v13216 int32
	_ = v13216
	var v13217 int32
	_ = v13217
	var v13220 int32
	_ = v13220
	var v13222 int32
	_ = v13222
	var v13226 int32
	_ = v13226
	var v13232 int32
	_ = v13232
	var v13240 int32
	_ = v13240
	var v13242 int32
	_ = v13242
	var v13266 int32
	_ = v13266
	var v13269 int32
	_ = v13269
	var v13270 int32
	_ = v13270
	var v13271 int32
	_ = v13271
	var v13272 int32
	_ = v13272
	var v13273 int32
	_ = v13273
	var v13274 int32
	_ = v13274
	var v13276 int32
	_ = v13276
	var v13294 int32
	_ = v13294
	var v13297 int32
	_ = v13297
	var v13298 int32
	_ = v13298
	var v13301 int32
	_ = v13301
	var v13302 int32
	_ = v13302
	var v13307 int32
	_ = v13307
	var v13314 int32
	_ = v13314
	var v13318 int32
	_ = v13318
	var v13324 int32
	_ = v13324
	var v13328 int32
	_ = v13328
	var v13332 int32
	_ = v13332
	var v13336 int32
	_ = v13336
	var v13342 int32
	_ = v13342
	var v13354 int32
	_ = v13354
	var v13355 int32
	_ = v13355
	var v13357 int32
	_ = v13357
	var v13361 int32
	_ = v13361
	var v13362 int32
	_ = v13362
	var v13364 int32
	_ = v13364
	var v13365 int32
	_ = v13365
	var v13367 int32
	_ = v13367
	var v13370 int32
	_ = v13370
	var v13371 int32
	_ = v13371
	var v13376 int64
	_ = v13376
	var v13377 int32
	_ = v13377
	var v13378 int32
	_ = v13378
	var v13379 int32
	_ = v13379
	var v13380 int32
	_ = v13380
	var v13384 int32
	_ = v13384
	var v13387 int32
	_ = v13387
	var v13388 int32
	_ = v13388
	var v13392 int32
	_ = v13392
	var v13431 int64
	_ = v13431
	var v13432 int32
	_ = v13432
	var v13436 int32
	_ = v13436
	var v13439 int32
	_ = v13439
	var v13440 int32
	_ = v13440
	var v13442 int32
	_ = v13442
	var v13443 int32
	_ = v13443
	var v13445 int64
	_ = v13445
	var v13447 int32
	_ = v13447
	var v13448 int32
	_ = v13448
	var v13490 int64
	_ = v13490
	var v13491 int64
	_ = v13491
	var v13494 int64
	_ = v13494
	var v13497 int32
	_ = v13497
	var v13498 int32
	_ = v13498
	var v13499 int32
	_ = v13499
	var v13540 int32
	_ = v13540
	var v13541 int32
	_ = v13541
	var v13542 int32
	_ = v13542
	var v13546 int32
	_ = v13546
	var v13549 int32
	_ = v13549
	var v13551 int32
	_ = v13551
	var v13553 int32
	_ = v13553
	var v13556 int32
	_ = v13556
	var v13557 int32
	_ = v13557
	var v13573 int32
	_ = v13573
	var v13575 int32
	_ = v13575
	var v13597 int32
	_ = v13597
	var v13601 int32
	_ = v13601
	var v13602 int32
	_ = v13602
	var v13605 int32
	_ = v13605
	var v13606 int32
	_ = v13606
	var v13609 int32
	_ = v13609
	var v13618 int32
	_ = v13618
	var v13626 int32
	_ = v13626
	var v13650 int32
	_ = v13650
	var v13654 int32
	_ = v13654
	var v13655 int32
	_ = v13655
	var v13656 int32
	_ = v13656
	var v13657 int32
	_ = v13657
	var v13658 int32
	_ = v13658
	var v13659 int32
	_ = v13659
	var v13660 int32
	_ = v13660
	var v13661 int32
	_ = v13661
	var v13662 int32
	_ = v13662
	var v13663 int32
	_ = v13663
	var v13664 int32
	_ = v13664
	var v13665 int32
	_ = v13665
	var v13666 int32
	_ = v13666
	var v13667 int32
	_ = v13667
	var v13668 int32
	_ = v13668
	var v13669 int32
	_ = v13669
	var v13670 int32
	_ = v13670
	var v13671 int32
	_ = v13671
	var v13673 int32
	_ = v13673
	var v13674 int32
	_ = v13674
	var v13675 int32
	_ = v13675
	var v13677 int32
	_ = v13677
	var v13678 int32
	_ = v13678
	var v13680 int32
	_ = v13680
	var v13681 int32
	_ = v13681
	var v13682 int32
	_ = v13682
	var v13698 int32
	_ = v13698
	var v13723 int32
	_ = v13723
	var v13725 int32
	_ = v13725
	var v13726 int32
	_ = v13726
	var v13730 int32
	_ = v13730
	var v13731 int32
	_ = v13731
	var v13745 int32
	_ = v13745
	var v13748 int32
	_ = v13748
	var v13773 int32
	_ = v13773
	var v13774 int32
	_ = v13774
	var v13777 int32
	_ = v13777
	var v13778 int32
	_ = v13778
	var v13779 int32
	_ = v13779
	var v13787 int32
	_ = v13787
	var v13790 int32
	_ = v13790
	var v13792 int32
	_ = v13792
	var v13794 int32
	_ = v13794
	var v13796 int32
	_ = v13796
	var v13798 int32
	_ = v13798
	var v13805 int32
	_ = v13805
	var v13806 float64
	_ = v13806
	var v13807 float64
	_ = v13807
	var v13808 float64
	_ = v13808
	var v13809 int32
	_ = v13809
	var v13811 int32
	_ = v13811
	var v13813 int32
	_ = v13813
	var v13815 int32
	_ = v13815
	var v13816 int32
	_ = v13816
	var v13817 int32
	_ = v13817
	var v13818 int32
	_ = v13818
	var v13819 int32
	_ = v13819
	var v13820 int32
	_ = v13820
	var v13823 int32
	_ = v13823
	var v13827 int32
	_ = v13827
	var v13837 int32
	_ = v13837
	var v13858 float64
	_ = v13858
	var v13863 float64
	_ = v13863
	var v13869 int32
	_ = v13869
	var v13873 int32
	_ = v13873
	var v13874 int64
	_ = v13874
	var v13878 int32
	_ = v13878
	var v13882 int32
	_ = v13882
	var v13883 int32
	_ = v13883
	var v13884 float64
	_ = v13884
	var v13885 float64
	_ = v13885
	var v13887 int64
	_ = v13887
	var v13892 int32
	_ = v13892
	var v13893 int32
	_ = v13893
	var v13894 int32
	_ = v13894
	var v13895 int64
	_ = v13895
	var v13897 int64
	_ = v13897
	var v13899 int32
	_ = v13899
	var v13900 float64
	_ = v13900
	var v13901 float64
	_ = v13901
	var v13903 int64
	_ = v13903
	var v13907 int32
	_ = v13907
	var v13908 int32
	_ = v13908
	var v13909 int64
	_ = v13909
	var v13911 int64
	_ = v13911
	var v13915 float64
	_ = v13915
	var v13916 float64
	_ = v13916
	var v13918 float64
	_ = v13918
	var v13921 float64
	_ = v13921
	var v13923 int32
	_ = v13923
	var v13924 int32
	_ = v13924
	var v13956 float64
	_ = v13956
	var v13961 float64
	_ = v13961
	var v13968 float64
	_ = v13968
	var v13970 float64
	_ = v13970
	var v13980 float64
	_ = v13980
	var v13982 int32
	_ = v13982
	var v13983 int32
	_ = v13983
	var v13984 int32
	_ = v13984
	var v13985 int32
	_ = v13985
	var v13986 int32
	_ = v13986
	var v13987 int32
	_ = v13987
	var v13988 int32
	_ = v13988
	var v13990 float64
	_ = v13990
	var v13991 int32
	_ = v13991
	var v13993 int32
	_ = v13993
	var v13996 float64
	_ = v13996
	var v13998 int32
	_ = v13998
	var v13999 int32
	_ = v13999
	var v14000 int32
	_ = v14000
	var v14001 int32
	_ = v14001
	var v14002 int32
	_ = v14002
	var v14003 int32
	_ = v14003
	var v14005 float64
	_ = v14005
	var v14006 int32
	_ = v14006
	var v14008 int32
	_ = v14008
	var v14013 float64
	_ = v14013
	var v14026 int32
	_ = v14026
	var v14027 float64
	_ = v14027
	var v14028 float64
	_ = v14028
	var v14033 int32
	_ = v14033
	var v14034 int32
	_ = v14034
	var v14038 int32
	_ = v14038
	var v14039 int32
	_ = v14039
	var v14042 int32
	_ = v14042
	var v14044 int32
	_ = v14044
	var v14046 int64
	_ = v14046
	var v14054 float64
	_ = v14054
	var v14067 float64
	_ = v14067
	var v14069 int32
	_ = v14069
	var v14072 int32
	_ = v14072
	var v14077 float64
	_ = v14077
	var v14080 float64
	_ = v14080
	var v14093 float64
	_ = v14093
	var v14098 float64
	_ = v14098
	var v14104 float64
	_ = v14104
	var v14111 float64
	_ = v14111
	var v14112 float64
	_ = v14112
	var v14115 float64
	_ = v14115
	var v14116 float64
	_ = v14116
	var v14117 float64
	_ = v14117
	var v14119 float64
	_ = v14119
	var v14124 int32
	_ = v14124
	var v14125 int32
	_ = v14125
	var v14142 int32
	_ = v14142
	var v14169 int32
	_ = v14169
	var v14212 int32
	_ = v14212
	var v14213 int32
	_ = v14213
	var v14215 int32
	_ = v14215
	var v14217 int32
	_ = v14217
	var v14220 int32
	_ = v14220
	var v14221 int32
	_ = v14221
	var v14222 float64
	_ = v14222
	var v14224 int32
	_ = v14224
	var v14227 int32
	_ = v14227
	var v14229 int32
	_ = v14229
	var v14236 int32
	_ = v14236
	var v14239 int32
	_ = v14239
	var v14241 int32
	_ = v14241
	var v14250 float64
	_ = v14250
	var v14254 int64
	_ = v14254
	var v14255 int64
	_ = v14255
	var v14259 int32
	_ = v14259
	var v14265 int32
	_ = v14265
	var v14268 int32
	_ = v14268
	var v14272 int32
	_ = v14272
	var v14274 int32
	_ = v14274
	var v14275 int32
	_ = v14275
	var v14278 int32
	_ = v14278
	var v14279 int32
	_ = v14279
	var v14281 int32
	_ = v14281
	var v14286 int32
	_ = v14286
	var v14287 int32
	_ = v14287
	var v14288 float64
	_ = v14288
	var v14290 int32
	_ = v14290
	var v14293 int32
	_ = v14293
	var v14295 int32
	_ = v14295
	var v14302 int32
	_ = v14302
	var v14305 int32
	_ = v14305
	var v14307 int32
	_ = v14307
	var v14316 float64
	_ = v14316
	var v14320 int64
	_ = v14320
	var v14321 int64
	_ = v14321
	var v14323 int32
	_ = v14323
	var v14328 int32
	_ = v14328
	var v14329 int32
	_ = v14329
	var v14330 int32
	_ = v14330
	var v14332 int32
	_ = v14332
	var v14334 int32
	_ = v14334
	var v14336 int32
	_ = v14336
	var v14338 int32
	_ = v14338
	var v14340 int32
	_ = v14340
	var v14341 int32
	_ = v14341
	var v14342 int32
	_ = v14342
	var v14345 int32
	_ = v14345
	var v14348 int32
	_ = v14348
	var v14349 int32
	_ = v14349
	var v14352 int32
	_ = v14352
	var v14353 int32
	_ = v14353
	var v14355 int32
	_ = v14355
	var v14357 int32
	_ = v14357
	var v14359 int32
	_ = v14359
	var v14361 int32
	_ = v14361
	var v14363 int32
	_ = v14363
	var v14365 int32
	_ = v14365
	var v14366 int32
	_ = v14366
	var v14367 int32
	_ = v14367
	var v14368 int32
	_ = v14368
	var v14369 int32
	_ = v14369
	var v14370 int32
	_ = v14370
	var v14371 int32
	_ = v14371
	var v14372 float64
	_ = v14372
	var v14373 int32
	_ = v14373
	var v14375 float64
	_ = v14375
	var v14376 int32
	_ = v14376
	var v14377 int32
	_ = v14377
	var v14386 int32
	_ = v14386
	var v14389 int32
	_ = v14389
	var v14392 int32
	_ = v14392
	var v14393 int32
	_ = v14393
	var v14395 int32
	_ = v14395
	var v14403 int32
	_ = v14403
	var v14404 int32
	_ = v14404
	var v14405 int32
	_ = v14405
	var v14406 int32
	_ = v14406
	var v14410 int32
	_ = v14410
	var v14414 int32
	_ = v14414
	var v14422 int32
	_ = v14422
	var v14425 int32
	_ = v14425
	var v14428 int32
	_ = v14428
	var v14446 int32
	_ = v14446
	var v14473 int32
	_ = v14473
	var v14474 int32
	_ = v14474
	var v14478 int32
	_ = v14478
	var v14479 int32
	_ = v14479
	var v14480 int32
	_ = v14480
	var v14481 int32
	_ = v14481
	var v14484 int32
	_ = v14484
	var v14485 int32
	_ = v14485
	var v14489 int32
	_ = v14489
	var v14529 int32
	_ = v14529
	var v14533 int32
	_ = v14533
	var v14534 int32
	_ = v14534
	var v14536 int32
	_ = v14536
	var v14554 int32
	_ = v14554
	var v14557 int32
	_ = v14557
	var v14558 int32
	_ = v14558
	var v14561 int32
	_ = v14561
	var v14562 int32
	_ = v14562
	var v14567 int32
	_ = v14567
	var v14574 int32
	_ = v14574
	var v14578 int32
	_ = v14578
	var v14584 int32
	_ = v14584
	var v14588 int32
	_ = v14588
	var v14592 int32
	_ = v14592
	var v14596 int32
	_ = v14596
	var v14602 int32
	_ = v14602
	var v14614 int32
	_ = v14614
	var v14616 int32
	_ = v14616
	var v14617 int32
	_ = v14617
	var v14621 int32
	_ = v14621
	var v14626 int32
	_ = v14626
	var v14630 int32
	_ = v14630
	var v14634 int32
	_ = v14634
	var v14635 int32
	_ = v14635
	var v14637 int32
	_ = v14637
	var v14638 int32
	_ = v14638
	var v14641 int32
	_ = v14641
	var v14644 int32
	_ = v14644
	var v14647 int32
	_ = v14647
	var v14655 int32
	_ = v14655
	var v14656 int32
	_ = v14656
	var v14660 int32
	_ = v14660
	var v14661 int32
	_ = v14661
	var v14663 int32
	_ = v14663
	var v14664 int32
	_ = v14664
	var v14665 int32
	_ = v14665
	var v14666 int32
	_ = v14666
	var v14668 int32
	_ = v14668
	var v14673 int32
	_ = v14673
	var v14674 int32
	_ = v14674
	var v14718 int32
	_ = v14718
	var v14719 int32
	_ = v14719
	var v14763 int32
	_ = v14763
	var v14766 int32
	_ = v14766
	var v14767 int32
	_ = v14767
	var v14774 int32
	_ = v14774
	var v14777 int32
	_ = v14777
	var v14780 int32
	_ = v14780
	var v14781 int32
	_ = v14781
	var v14785 int32
	_ = v14785
	var v14789 int32
	_ = v14789
	var v14790 int32
	_ = v14790
	var v14794 int32
	_ = v14794
	var v14798 int32
	_ = v14798
	var v14804 int32
	_ = v14804
	var v14807 int32
	_ = v14807
	var v14809 int32
	_ = v14809
	var v14810 int32
	_ = v14810
	var v14813 int32
	_ = v14813
	var v14814 int32
	_ = v14814
	var v14816 int32
	_ = v14816
	var v14817 int32
	_ = v14817
	var v14820 int32
	_ = v14820
	var v14826 int32
	_ = v14826
	var v14829 int32
	_ = v14829
	var v14833 int32
	_ = v14833
	var v14834 int32
	_ = v14834
	var v14839 int32
	_ = v14839
	var v14841 int32
	_ = v14841
	var v14842 int32
	_ = v14842
	var v14843 int32
	_ = v14843
	var v14885 int32
	_ = v14885
	var v14888 int32
	_ = v14888
	var v14891 int32
	_ = v14891
	var v14897 int32
	_ = v14897
	var v14900 int32
	_ = v14900
	var v14904 int32
	_ = v14904
	var v14906 int32
	_ = v14906
	var v14910 int32
	_ = v14910
	var v14913 float64
	_ = v14913
	var v14915 int32
	_ = v14915
	var v14918 int32
	_ = v14918
	var v14920 int32
	_ = v14920
	var v14927 int32
	_ = v14927
	var v14930 int32
	_ = v14930
	var v14932 int32
	_ = v14932
	var v14941 float64
	_ = v14941
	var v14945 int64
	_ = v14945
	var v14946 int64
	_ = v14946
	var v14948 int32
	_ = v14948
	var v14951 int32
	_ = v14951
	var v14954 int32
	_ = v14954
	var v14955 int32
	_ = v14955
	var v14956 int32
	_ = v14956
	var v14960 int32
	_ = v14960
	var v14962 int64
	_ = v14962
	var v14964 int32
	_ = v14964
	var v14966 int32
	_ = v14966
	var v14968 int32
	_ = v14968
	var v14970 int32
	_ = v14970
	var v14972 int32
	_ = v14972
	var v14975 int32
	_ = v14975
	var v14988 int32
	_ = v14988
	var v15020 int32
	_ = v15020
	var v15021 int32
	_ = v15021
	var v15025 int32
	_ = v15025
	var v15026 int32
	_ = v15026
	var v15028 int32
	_ = v15028
	var v15046 int32
	_ = v15046
	var v15049 int32
	_ = v15049
	var v15050 int32
	_ = v15050
	var v15053 int32
	_ = v15053
	var v15054 int32
	_ = v15054
	var v15059 int32
	_ = v15059
	var v15066 int32
	_ = v15066
	var v15070 int32
	_ = v15070
	var v15076 int32
	_ = v15076
	var v15080 int32
	_ = v15080
	var v15084 int32
	_ = v15084
	var v15088 int32
	_ = v15088
	var v15094 int32
	_ = v15094
	var v15106 int32
	_ = v15106
	var v15108 int32
	_ = v15108
	var v15109 int32
	_ = v15109
	var v15113 int32
	_ = v15113
	var v15118 int32
	_ = v15118
	var v15119 int32
	_ = v15119
	var v15123 int32
	_ = v15123
	var v15126 int32
	_ = v15126
	var v15127 int32
	_ = v15127
	var v15128 int32
	_ = v15128
	var v15129 int32
	_ = v15129
	var v15131 int32
	_ = v15131
	var v15134 int32
	_ = v15134
	var v15135 int32
	_ = v15135
	var v15136 int32
	_ = v15136
	var v15137 int32
	_ = v15137
	var v15138 int32
	_ = v15138
	var v15139 int32
	_ = v15139
	var v15140 int32
	_ = v15140
	var v15141 int32
	_ = v15141
	var v15143 int32
	_ = v15143
	var v15149 int32
	_ = v15149
	var v15150 int32
	_ = v15150
	var v15193 int32
	_ = v15193
	var v15196 int32
	_ = v15196
	var v15199 int32
	_ = v15199
	var v15202 int32
	_ = v15202
	var v15205 int32
	_ = v15205
	var v15206 int32
	_ = v15206
	var v15209 int32
	_ = v15209
	var v15249 int32
	_ = v15249
	var v15250 int32
	_ = v15250
	var v15254 int32
	_ = v15254
	var v15255 int32
	_ = v15255
	var v15257 int32
	_ = v15257
	var v15275 int32
	_ = v15275
	var v15278 int32
	_ = v15278
	var v15279 int32
	_ = v15279
	var v15282 int32
	_ = v15282
	var v15283 int32
	_ = v15283
	var v15288 int32
	_ = v15288
	var v15295 int32
	_ = v15295
	var v15299 int32
	_ = v15299
	var v15305 int32
	_ = v15305
	var v15309 int32
	_ = v15309
	var v15313 int32
	_ = v15313
	var v15317 int32
	_ = v15317
	var v15323 int32
	_ = v15323
	var v15335 int32
	_ = v15335
	var v15337 int32
	_ = v15337
	var v15338 int32
	_ = v15338
	var v15342 int32
	_ = v15342
	var v15347 int32
	_ = v15347
	var v15348 int32
	_ = v15348
	var v15352 int32
	_ = v15352
	var v15355 int32
	_ = v15355
	var v15356 int32
	_ = v15356
	var v15357 int32
	_ = v15357
	var v15358 int32
	_ = v15358
	var v15359 int32
	_ = v15359
	var v15363 float64
	_ = v15363
	var v15364 int32
	_ = v15364
	var v15365 float64
	_ = v15365
	var v15367 int32
	_ = v15367
	var v15373 float64
	_ = v15373
	var v15377 float64
	_ = v15377
	var v15379 float64
	_ = v15379
	var v15381 float64
	_ = v15381
	var v15382 float64
	_ = v15382
	var v15391 float64
	_ = v15391
	var v15395 float64
	_ = v15395
	var v15397 int32
	_ = v15397
	var v15398 int32
	_ = v15398
	var v15401 int32
	_ = v15401
	var v15402 int32
	_ = v15402
	var v15403 int32
	_ = v15403
	var v15404 int32
	_ = v15404
	var v15405 int32
	_ = v15405
	var v15406 int32
	_ = v15406
	var v15407 int32
	_ = v15407
	var v15408 int32
	_ = v15408
	var v15409 int32
	_ = v15409
	var v15410 int32
	_ = v15410
	var v15412 int32
	_ = v15412
	var v15417 int32
	_ = v15417
	var v15418 int32
	_ = v15418
	var v15461 int32
	_ = v15461
	var v15464 int32
	_ = v15464
	var v15470 int32
	_ = v15470
	var v15473 int32
	_ = v15473
	var v15477 int32
	_ = v15477
	var v15478 int32
	_ = v15478
	var v15481 int32
	_ = v15481
	var v15482 int32
	_ = v15482
	var v15484 int32
	_ = v15484
	var v15489 int32
	_ = v15489
	var v15528 int32
	_ = v15528
	var v15529 int32
	_ = v15529
	var v15530 int32
	_ = v15530
	var v15533 int32
	_ = v15533
	var v15534 int32
	_ = v15534
	var v15535 int32
	_ = v15535
	var v15538 int32
	_ = v15538
	var v15539 int32
	_ = v15539
	var v15540 int32
	_ = v15540
	var v15543 int32
	_ = v15543
	var v15545 int32
	_ = v15545
	var v15547 int32
	_ = v15547
	var v15549 int32
	_ = v15549
	var v15551 int32
	_ = v15551
	var v15553 int32
	_ = v15553
	var v15556 int32
	_ = v15556
	var v15560 int32
	_ = v15560
	var v15564 int32
	_ = v15564
	var v15568 int32
	_ = v15568
	var v15570 int32
	_ = v15570
	var v15571 int32
	_ = v15571
	var v15573 int32
	_ = v15573
	var v15582 int32
	_ = v15582
	var v15583 int32
	_ = v15583
	var v15594 float64
	_ = v15594
	var v15598 int64
	_ = v15598
	var v15599 int64
	_ = v15599
	var v15601 int32
	_ = v15601
	var v15605 int32
	_ = v15605
	var v15606 int32
	_ = v15606
	var v15607 int32
	_ = v15607
	var v15608 int32
	_ = v15608
	var v15609 int32
	_ = v15609
	var v15611 int32
	_ = v15611
	var v15612 int32
	_ = v15612
	var v15616 int32
	_ = v15616
	var v15617 int32
	_ = v15617
	var v15624 float64
	_ = v15624
	var v15631 int32
	_ = v15631
	var v15633 float64
	_ = v15633
	var v15636 float64
	_ = v15636
	var v15637 float64
	_ = v15637
	var v15639 float64
	_ = v15639
	var v15647 int32
	_ = v15647
	var v15648 int32
	_ = v15648
	var v15649 int32
	_ = v15649
	var v15652 int32
	_ = v15652
	var v15655 int32
	_ = v15655
	var v15658 int32
	_ = v15658
	var v15661 int32
	_ = v15661
	var v15662 int64
	_ = v15662
	var v15666 int32
	_ = v15666
	var v15667 int32
	_ = v15667
	var v15668 int32
	_ = v15668
	var v15669 int32
	_ = v15669
	var v15671 int32
	_ = v15671
	var v15672 int32
	_ = v15672
	var v15675 int32
	_ = v15675
	var v15676 int32
	_ = v15676
	var v15684 int32
	_ = v15684
	var v15685 int32
	_ = v15685
	var v15688 int32
	_ = v15688
	var v15692 int32
	_ = v15692
	var v15694 int32
	_ = v15694
	var v15701 int32
	_ = v15701
	var v15702 int32
	_ = v15702
	var v15703 int32
	_ = v15703
	var v15708 int32
	_ = v15708
	var v15710 int32
	_ = v15710
	var v15713 int32
	_ = v15713
	var v15721 int32
	_ = v15721
	var v15722 int32
	_ = v15722
	var v15725 int32
	_ = v15725
	var v15726 int32
	_ = v15726
	var v15727 int32
	_ = v15727
	var v15728 int32
	_ = v15728
	var v15733 int32
	_ = v15733
	var v15741 int32
	_ = v15741
	var v15744 int32
	_ = v15744
	var v15747 int32
	_ = v15747
	var v15751 int32
	_ = v15751
	var v15754 int32
	_ = v15754
	var v15755 int32
	_ = v15755
	var v15759 int32
	_ = v15759
	var v15766 int32
	_ = v15766
	var v15768 int32
	_ = v15768
	var v15776 int32
	_ = v15776
	var v15777 int32
	_ = v15777
	var v15790 int32
	_ = v15790
	var v15799 int32
	_ = v15799
	var v15801 int32
	_ = v15801
	var v15808 int32
	_ = v15808
	var v15809 int32
	_ = v15809
	var v15812 int32
	_ = v15812
	var v15813 int32
	_ = v15813
	var v15815 int32
	_ = v15815
	var v15835 int32
	_ = v15835
	var v15836 int32
	_ = v15836
	var v15837 int32
	_ = v15837
	var v15839 int32
	_ = v15839
	var v15842 int32
	_ = v15842
	var v15843 int32
	_ = v15843
	var v15846 int32
	_ = v15846
	var v15847 int32
	_ = v15847
	var v15854 int32
	_ = v15854
	var v15860 int32
	_ = v15860
	var v15861 int32
	_ = v15861
	var v15862 int32
	_ = v15862
	var v15863 int32
	_ = v15863
	var v15866 int32
	_ = v15866
	var v15868 int32
	_ = v15868
	var v15869 int32
	_ = v15869
	var v15870 int32
	_ = v15870
	var v15871 int32
	_ = v15871
	var v15872 int32
	_ = v15872
	var v15873 int32
	_ = v15873
	var v15874 int32
	_ = v15874
	var v15876 int32
	_ = v15876
	var v15877 int32
	_ = v15877
	var v15879 int32
	_ = v15879
	var v15880 int32
	_ = v15880
	var v15881 int32
	_ = v15881
	var v15882 int32
	_ = v15882
	var v15883 int32
	_ = v15883
	var v15884 int32
	_ = v15884
	var v15885 int32
	_ = v15885
	var v15887 int32
	_ = v15887
	var v15888 int32
	_ = v15888
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
	var v15894 int32
	_ = v15894
	var v15895 int32
	_ = v15895
	var v15902 int32
	_ = v15902
	var v15915 int32
	_ = v15915
	var v15940 int32
	_ = v15940
	var v15944 int32
	_ = v15944
	var v15945 int32
	_ = v15945
	var v15946 int32
	_ = v15946
	var v15947 int32
	_ = v15947
	var v15948 int32
	_ = v15948
	var v15949 int32
	_ = v15949
	var v15951 int32
	_ = v15951
	var v15952 int32
	_ = v15952
	var v15953 int32
	_ = v15953
	var v15955 int32
	_ = v15955
	var v15958 int32
	_ = v15958
	var v15959 int32
	_ = v15959
	var v15960 int32
	_ = v15960
	var v15961 int32
	_ = v15961
	var v15962 int32
	_ = v15962
	var v15964 int32
	_ = v15964
	var v15965 int32
	_ = v15965
	var v15967 int32
	_ = v15967
	var v15968 int32
	_ = v15968
	var v15973 int32
	_ = v15973
	var v16011 int32
	_ = v16011
	var v16012 int32
	_ = v16012
	var v16032 int32
	_ = v16032
	var v16054 int32
	_ = v16054
	var v16057 int32
	_ = v16057
	var v16059 int32
	_ = v16059
	var v16060 int32
	_ = v16060
	var v16061 int32
	_ = v16061
	var v16062 int32
	_ = v16062
	var v16063 int32
	_ = v16063
	var v16071 int32
	_ = v16071
	var v16078 int32
	_ = v16078
	var v16079 int32
	_ = v16079
	var v16082 int32
	_ = v16082
	var v16083 int32
	_ = v16083
	var v16085 int32
	_ = v16085
	var v16105 int32
	_ = v16105
	var v16112 int32
	_ = v16112
	var v16114 int32
	_ = v16114
	var v16115 int32
	_ = v16115
	var v16118 int32
	_ = v16118
	var v16122 int32
	_ = v16122
	var v16125 int32
	_ = v16125
	var v16127 int32
	_ = v16127
	var v16130 int32
	_ = v16130
	var v16137 int32
	_ = v16137
	var v16139 int32
	_ = v16139
	var v16147 int32
	_ = v16147
	var v16148 int32
	_ = v16148
	var v16161 int32
	_ = v16161
	var v16171 int32
	_ = v16171
	var v16178 int32
	_ = v16178
	var v16179 int32
	_ = v16179
	var v16182 int32
	_ = v16182
	var v16183 int32
	_ = v16183
	var v16205 int32
	_ = v16205
	var v16211 int32
	_ = v16211
	var v16212 int32
	_ = v16212
	var v16213 int32
	_ = v16213
	var v16216 int32
	_ = v16216
	var v16222 int32
	_ = v16222
	var v16223 int32
	_ = v16223
	var v16225 int32
	_ = v16225
	var v16226 int32
	_ = v16226
	var v16232 int32
	_ = v16232
	var v16233 int32
	_ = v16233
	var v16234 int32
	_ = v16234
	var v16235 int32
	_ = v16235
	var v16241 int32
	_ = v16241
	var v16242 int32
	_ = v16242
	var v16243 int32
	_ = v16243
	var v16244 int32
	_ = v16244
	var v16250 int32
	_ = v16250
	var v16251 int32
	_ = v16251
	var v16252 int32
	_ = v16252
	var v16253 int32
	_ = v16253
	var v16263 int32
	_ = v16263
	var v16264 int32
	_ = v16264
	var v16265 int32
	_ = v16265
	var v16267 int32
	_ = v16267
	var v16270 int32
	_ = v16270
	var v16276 int32
	_ = v16276
	var v16277 int32
	_ = v16277
	var v16279 int32
	_ = v16279
	var v16280 int32
	_ = v16280
	var v16286 int32
	_ = v16286
	var v16287 int32
	_ = v16287
	var v16288 int32
	_ = v16288
	var v16289 int32
	_ = v16289
	var v16291 int32
	_ = v16291
	var v16297 int32
	_ = v16297
	var v16298 int32
	_ = v16298
	var v16299 int32
	_ = v16299
	var v16300 int32
	_ = v16300
	var v16306 int32
	_ = v16306
	var v16307 int32
	_ = v16307
	var v16308 int32
	_ = v16308
	var v16309 int32
	_ = v16309
	var v16310 int32
	_ = v16310
	var v16330 int32
	_ = v16330
	var v16331 int32
	_ = v16331
	var v16334 int32
	_ = v16334
	var v16335 int32
	_ = v16335
	var v16336 int32
	_ = v16336
	var v16337 int32
	_ = v16337
	var v16357 int32
	_ = v16357
	var v16358 int32
	_ = v16358
	var v16364 int32
	_ = v16364
	var v16365 int32
	_ = v16365
	var v16373 int32
	_ = v16373
	var v16380 int32
	_ = v16380
	var v16381 int32
	_ = v16381
	var v16384 int32
	_ = v16384
	var v16385 int32
	_ = v16385
	var v16386 int32
	_ = v16386
	var v16387 int32
	_ = v16387
	var v16407 int32
	_ = v16407
	var v16408 int32
	_ = v16408
	var v16409 int32
	_ = v16409
	var v16410 int32
	_ = v16410
	var v16412 int32
	_ = v16412
	var v16413 int32
	_ = v16413
	var v16414 int32
	_ = v16414
	var v16415 int32
	_ = v16415
	var v16416 int32
	_ = v16416
	var v16419 int32
	_ = v16419
	var v16420 int32
	_ = v16420
	var v16424 int32
	_ = v16424
	var v16425 int32
	_ = v16425
	var v16434 int32
	_ = v16434
	var v16436 float64
	_ = v16436
	var v16438 float64
	_ = v16438
	var v16440 float64
	_ = v16440
	var v16442 int32
	_ = v16442
	var v16443 int32
	_ = v16443
	var v16446 int32
	_ = v16446
	var v16462 int32
	_ = v16462
	var v16466 int32
	_ = v16466
	var v16470 int32
	_ = v16470
	var v16472 int32
	_ = v16472
	var v16473 int32
	_ = v16473
	var v16475 int32
	_ = v16475
	var v16485 int32
	_ = v16485
	var v16496 float64
	_ = v16496
	var v16500 int64
	_ = v16500
	var v16501 int64
	_ = v16501
	var v16503 int32
	_ = v16503
	var v16505 int32
	_ = v16505
	var v16507 int32
	_ = v16507
	var v16508 int32
	_ = v16508
	var v16510 int32
	_ = v16510
	var v16514 int32
	_ = v16514
	var v16518 int32
	_ = v16518
	var v16521 int32
	_ = v16521
	var v16523 int32
	_ = v16523
	var v16533 int32
	_ = v16533
	var v16544 float64
	_ = v16544
	var v16548 int64
	_ = v16548
	var v16549 int64
	_ = v16549
	var v16551 int32
	_ = v16551
	var v16554 int32
	_ = v16554
	var v16557 int32
	_ = v16557
	var v16558 int32
	_ = v16558
	var v16562 int32
	_ = v16562
	var v16565 int32
	_ = v16565
	var v16568 int32
	_ = v16568
	var v16571 int32
	_ = v16571
	var v16572 int64
	_ = v16572
	var v16575 int32
	_ = v16575
	var v16578 int32
	_ = v16578
	var v16583 int32
	_ = v16583
	var v16623 int32
	_ = v16623
	var v16627 int32
	_ = v16627
	var v16629 int32
	_ = v16629
	var v16631 int32
	_ = v16631
	var v16632 int32
	_ = v16632
	var v16675 int32
	_ = v16675
	var v16676 int32
	_ = v16676
	var v16680 int32
	_ = v16680
	var v16681 int32
	_ = v16681
	var v16685 int32
	_ = v16685
	var v16688 int32
	_ = v16688
	var v16689 int32
	_ = v16689
	var v16692 int32
	_ = v16692
	var v16693 int64
	_ = v16693
	var v16699 int32
	_ = v16699
	var v16743 int32
	_ = v16743
	var v16746 int32
	_ = v16746
	var v16753 int32
	_ = v16753
	var v16756 int32
	_ = v16756
	var v16761 int32
	_ = v16761
	var v16768 int32
	_ = v16768
	var v16771 int32
	_ = v16771
	var v16775 int32
	_ = v16775
	var v16778 int32
	_ = v16778
	var v16779 int32
	_ = v16779
	var v16784 int32
	_ = v16784
	var v16786 int32
	_ = v16786
	var v16789 int32
	_ = v16789
	var v16790 int32
	_ = v16790
	var v16791 int32
	_ = v16791
	var v16798 float64
	_ = v16798
	var v16799 int32
	_ = v16799
	var v16802 int32
	_ = v16802
	var v16805 int32
	_ = v16805
	var v16814 int32
	_ = v16814
	var v16817 int32
	_ = v16817
	var v16822 int32
	_ = v16822
	var v16823 float64
	_ = v16823
	var v16824 int32
	_ = v16824
	var v16826 int32
	_ = v16826
	var v16827 int32
	_ = v16827
	var v16828 float64
	_ = v16828
	var v16829 float64
	_ = v16829
	var v16832 int32
	_ = v16832
	var v16833 float64
	_ = v16833
	var v16834 float64
	_ = v16834
	var v16836 float64
	_ = v16836
	var v16837 int32
	_ = v16837
	var v16838 int32
	_ = v16838
	var v16842 int32
	_ = v16842
	var v16844 int32
	_ = v16844
	var v16846 int32
	_ = v16846
	var v16850 int32
	_ = v16850
	var v16853 int32
	_ = v16853
	var v16859 float64
	_ = v16859
	var v16863 int32
	_ = v16863
	var v16864 float64
	_ = v16864
	var v16865 float64
	_ = v16865
	var v16868 int32
	_ = v16868
	var v16875 int32
	_ = v16875
	var v16881 float64
	_ = v16881
	var v16882 int32
	_ = v16882
	var v16885 int32
	_ = v16885
	var v16891 int32
	_ = v16891
	var v16901 int32
	_ = v16901
	var v16905 int32
	_ = v16905
	var v16906 float64
	_ = v16906
	var v16909 float64
	_ = v16909
	var v16912 int32
	_ = v16912
	var v16915 int32
	_ = v16915
	var v16916 int32
	_ = v16916
	var v16930 int32
	_ = v16930
	var v16934 int32
	_ = v16934
	var v16937 int32
	_ = v16937
	var v16941 int32
	_ = v16941
	var v16951 int32
	_ = v16951
	var v16955 int32
	_ = v16955
	var v16956 float64
	_ = v16956
	var v16959 float64
	_ = v16959
	var v16963 int32
	_ = v16963
	var v16964 int32
	_ = v16964
	var v16987 int32
	_ = v16987
	v6 = l5
	v9 = int32(0)
	v39 = int64(0)
	v42 = m.G0
	v44 = v42 - int32(32)
	m.G0 = v44
	v47 = F_palloc0(m, int32(400))
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
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(269)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v59 = v55 + int32(1)
	goto L5
L4:
	;
	v59 = int32(1)
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v59
	if l4 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v63 = v62
	goto L8
L7:
	;
	v63 = l2
	goto L8
L8:
	;
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v47)+28)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v63
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[0]))
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+337)) = uint8(v70)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+335)) = uint16(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+328)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v47)+300)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v47)+80)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v47)+88)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v47)+93)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v47)+124)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v47)+132)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v47)+140)) = v64
	base.MemoryFill(m, v47+int32(212), v70, int32(88))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+339)) = uint8(v70)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+338)) = uint8(v6)
	if v6 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v98 = F_assign_special_exec_param(m, v47)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v100 = int32(-1)
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+364)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+360)) = v100
	v105 = F_palloc0(m, int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v100 = v98
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = int32(275)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v105
	v114 = F_list_make1_impl(m, int32(1), v44+int32(16))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+92)) = v114
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v117 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v118 = m.G0
	v120 = v118 - int32(32)
	m.G0 = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+48))
	if v123 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	goto L17
L17:
	;
	v476 = int32(0)
	v480 = m.G0
	v482 = v480 - int32(32)
	m.G0 = v482
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v484 == int32(5) {
		goto L91
	} else {
		goto L92
	}
L18:
	;
	goto L17
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L88
	}
L20:
	;
	m.G0 = v120 + int32(32)
	goto L18
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v126 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v141 = v9
	goto L23
L23:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170+v141<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+36))
	if v177|base.B2i32(v176 != int32(1)) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L20
L25:
	;
	v375 = v141 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v375 < v376 {
		v141 = v375
		goto L23
	} else {
		goto L87
	}
L26:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v47)+84))
	v185 = F_lappend_int(m, v183, int32(-1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	switch v188 {
	case 0:
		goto L32
	default:
		goto L30
	case 2:
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+84)) = v185
	goto L25
L30:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v266 = F_copyObjectImpl(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L62
	}
L31:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+32)))
	if v191|base.B2i32(v176 != int32(1)) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	if v177 != int32(1) {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	if v195 == int32(67) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v174)+36))
	if v208 < int32(2) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v175)+140))
	if v198 != 0 {
		goto L30
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v206 = F_expression_tree_walker_impl(m, v175, int32(890), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L42
	}
L39:
	;
	v200 = int32(0)
	v202 = F_query_tree_walker_impl(m, v175, int32(890), v200, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v202 != 0 {
		goto L30
	} else {
		goto L41
	}
L41:
	;
	goto L35
L42:
	;
	if v206 != 0 {
		goto L30
	} else {
		goto L43
	}
L43:
	;
	goto L35
L44:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v245 = F_contain_volatile_functions(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L58
	}
L45:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v212 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v212
	if v211 == v212 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v216 != int32(67) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v240 = F_expression_tree_walker_impl(m, v211, int32(891), v120+int32(20))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L56
	}
L48:
	;
	if v216 != int32(101) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = int32(1)
	v233 = F_query_tree_walker_impl(m, v211, int32(891), v120+int32(20), int32(16))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L54
	}
L51:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	if v221 != int32(6) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+92)))
	if v224 == int32(0) {
		goto L44
	} else {
		goto L53
	}
L53:
	;
	goto L30
L54:
	;
	if v233 == int32(0) {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L30
L56:
	;
	if v240 != 0 {
		goto L30
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	if v245 != 0 {
		goto L30
	} else {
		goto L59
	}
L59:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+20)) = v247
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = v251
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v256 = F_inline_cte_walker(m, v253, v120+int32(20))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v47)+84))
	v260 = F_lappend_int(m, v258, int32(-1))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+84)) = v260
	goto L25
L62:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v271 = F_choose_plan_name(m, v268, v269, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v273 = int32(0)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+32)))
	v277 = F_subquery_planner(m, v268, v266, v271, v47, v273, v274, float64(0), v273)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v279 != 0 {
		goto L19
	} else {
		goto L65
	}
L65:
	;
	v282 = F_fetch_upper_rel(m, v277, int32(7), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v282)+60))
	v285 = F_create_plan(m, v277, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v288 = F_palloc0(m, int32(72))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v288))) = int64(30064771095)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v277)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v288)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v288)+20)) = v292
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v285)+44))
	if v296 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v319 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v288)+37)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v288)+32)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v288)+44)) = v319
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
	*(*int64)(unsafe.Add(mBase, uint32(v288)+24)) = int64(-4294965018)
	v318 = int32(0)
	goto L69
L71:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v296)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+26)))
	if v301 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v303 = F_exprType(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+24)) = v303
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v307 = F_exprTypmod(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+28)) = v307
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	v311 = F_exprCollation(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v318 = v311
	goto L69
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+12)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v324
	v331 = F_list_make1_impl(m, int32(479), v120+int32(12))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+40)) = v331
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	v336 = F_lappend(m, v335, v285)
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
	v342 = F_lappend(m, v341, v284)
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
	v348 = F_lappend(m, v347, v277)
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
	*(*int32)(unsafe.Add(mBase, uint32(v288)+16)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v47)+80))
	v359 = F_lappend(m, v358, v288)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+80)) = v359
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v47)+84))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v47)+84)) = v364
	F_cost_subplan(m, v288, v285)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L25
L87:
	;
	goto L24
L88:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_0), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_1), int32(989), int32(_a_F_subquery_planner_2))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v487 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v482)+30)) = uint16(v487)
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v489 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v482 + int32(32)
	v1101 = F_preprocess_relation_rtes(m, v47)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L173
	}
L94:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v490 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v691 = v476
	v698 = v476
	v730 = int32(0)
	goto L96
L96:
	;
	v732 = F_palloc0(m, int32(136))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L1
	} else {
		goto L125
	}
L97:
	;
	v674 = int32(1)
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+31)))
	v677 = v675 & v674
	if v677 != 0 {
		goto L116
	} else {
		goto L117
	}
L98:
	;
	if v490 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v495 = int32(0)
	if v495 < v490 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v577 = v476
	goto L101
L101:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v618+v577<<(uint(int32(2))%32))))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	if v623 == int32(7) {
		goto L97
	} else {
		goto L115
	}
L102:
	;
	v498 = v490
	goto L104
L103:
	;
	v498 = v495
	goto L104
L104:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v504 = int32(0)
	v506 = v504
	v508 = v504
	goto L105
L105:
	;
	v549 = v503 + v506<<(uint(int32(2))%32)
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+8))
	if v551 != int32(7) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v498&int32(1) == int32(0) {
		goto L97
	} else {
		goto L114
	}
L107:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550)+4))
	v558 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v554+(v482+int32(29))))) = uint8(v558)
	goto L109
L108:
	;
	goto L109
L109:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+8))
	if v561 != int32(7) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v568 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v564+(v482+int32(29))))) = uint8(v568)
	goto L112
L111:
	;
	goto L112
L112:
	;
	v570 = int32(2)
	v571 = v506 + v570
	v573 = v508 + v570
	if v573 != v498&int32(2147483646) {
		v506 = v571
		v508 = v573
		goto L105
	} else {
		goto L113
	}
L113:
	;
	goto L106
L114:
	;
	v577 = v571
	goto L101
L115:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v622)+4))
	v630 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v626+(v482+int32(29))))) = uint8(v630)
	goto L97
L116:
	;
	v678 = int32(2)
	goto L118
L117:
	;
	v678 = v674
	goto L118
L118:
	;
	if v677 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v681 = int32(3)
	goto L121
L120:
	;
	v681 = int32(0)
	goto L121
L121:
	;
	v682 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v482)+30)))
	v684 = v682 & int32(1)
	if v684 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v685 = v678
	goto L124
L123:
	;
	v685 = v681
	goto L124
L124:
	;
	v691 = v685
	v698 = v675
	v730 = int32(0) - v684
	goto L96
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+44)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v732)+12)) = int32(2)
	v737 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v732)+48)) = v737
	*(*int64)(unsafe.Add(mBase, uint32(v732))) = int64(101)
	*(*int64)(unsafe.Add(mBase, uint32(v732)+56)) = v737
	v743 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v732)+64)) = v743
	v747 = F_makeAlias(m, int32(_a_F_subquery_planner_3), v743)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v732)+8)) = v747
	v750 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v732)+124)) = uint16(v750)
	v752 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v732)+20)) = uint8(v752)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v755 = F_lappend(m, v754, v732)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v755
	if v755 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v755)+4))
	v759 = v758
	goto L130
L129:
	;
	v759 = v476
	goto L130
L130:
	;
	v761 = F_palloc0(m, int32(8))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v761))) = int32(63)
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v761)+4)) = v765
	*(*int32)(unsafe.Add(mBase, uint32(v482)+16)) = v761
	*(*int32)(unsafe.Add(mBase, uint32(v482)+24)) = v761
	v772 = F_list_make1_impl(m, int32(1), v482+int32(16))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v774)+8))
	v776 = F_makeFromExpr(m, v772, v775)
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+4))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	switch v783 - int32(63) {
	case 0:
		v801 = int32(4)
		goto L134
	case 1:
		goto L135
	default:
		goto L136
	}
L134:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v801+v782)))
	v805 = F_palloc0(m, int32(40))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L140
	}
L135:
	;
	v801 = int32(36)
	goto L134
L136:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	*(*int32)(unsafe.Add(mBase, uint32(v482))) = v790
	F_errmsg_internal(m, int32(_a_F_subquery_planner_4), v482)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_5), int32(297), int32(_a_F_subquery_planner_6))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v805)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v805)+16)) = v782
	*(*int32)(unsafe.Add(mBase, uint32(v805)+12)) = v776
	v811 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v805)+8)) = uint8(v811)
	*(*int32)(unsafe.Add(mBase, uint32(v805)+4)) = v691
	*(*int32)(unsafe.Add(mBase, uint32(v805))) = int32(64)
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v805)+36)) = v759
	*(*int32)(unsafe.Add(mBase, uint32(v805)+32)) = v811
	*(*int32)(unsafe.Add(mBase, uint32(v805)+28)) = v816
	*(*int32)(unsafe.Add(mBase, uint32(v482)+12)) = v805
	*(*int32)(unsafe.Add(mBase, uint32(v482)+20)) = v805
	v826 = F_list_make1_impl(m, int32(1), v482+int32(12))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+4)) = v826
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v831 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v830)+8)) = v831
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if (base.B2i32(v833 == v831)|(v698^int32(-1)))&int32(1) == v831 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v844 = F_bms_make_singleton(m, v843)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if v730&int32(1) != 0 {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	v846 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v848 = F_add_nulling_relids(m, v833, v844, v846)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v848
	goto L144
L148:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v854 = F_bms_make_singleton(m, v803)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	v1055 = int32(0)
	goto L150
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v1055
	goto L93
L151:
	;
	v856 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v858 = F_add_nulling_relids(m, v853, v854, v856)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v858
	v861 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v861 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v976 = F_bms_make_singleton(m, v803)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L166
	}
L155:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v864 <= int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v870 = int32(0)
	goto L157
L157:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v861)+12))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v909+v870<<(uint(int32(2))%32))))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v913)+16))
	v915 = F_bms_make_singleton(m, v803)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	goto L154
L159:
	;
	v917 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v919 = F_add_nulling_relids(m, v914, v915, v917)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913)+16)) = v919
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v913)+20))
	v923 = F_bms_make_singleton(m, v803)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v925 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v927 = F_add_nulling_relids(m, v922, v923, v925)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v913)+20)) = v927
	v931 = v870 + int32(1)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v861)+4))
	if v931 < v932 {
		v870 = v931
		goto L157
	} else {
		goto L165
	}
L165:
	;
	goto L158
L166:
	;
	v978 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v980 = F_add_nulling_relids(m, v975, v976, v978)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v980
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v983)+12))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v984+v803<<(uint(int32(2))%32)-int32(4))))
	v991 = int32(0)
	v993 = F_makeWholeRowVar(m, v990, v803, v991, v991)
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v995 = F_bms_make_singleton(m, v759)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+24)) = v995
	v999 = F_palloc0(m, int32(20))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v999)+16)) = int32(-1)
	v1003 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v999)+12)) = uint8(v1003)
	*(*int32)(unsafe.Add(mBase, uint32(v999)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = v993
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = int32(52)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v1011 = F_make_and_qual(m, v999, v1010)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v1055 = v1011
	goto L150
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v1101
	F_replace_empty_jointree(m, v1101)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+39)))
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
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+60))
	v1260 = int32(0)
	v1262 = F_pull_up_subqueries_recurse(m, v47, v1259, v1260, v1260)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
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
	v1149 = v1140
	goto L187
L187:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+12))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1189+v1149<<(uint(int32(2))%32))))
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
	v1214 = v1149 + int32(1)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+4))
	if v1214 < v1215 {
		v1149 = v1214
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
	v1201 = F_inline_function_in_from(m, v47, v1193)
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
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1264)+60)) = v1262
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+144))
	if v1266 != 0 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1267 = m.G0
	v1269 = v1267 - int32(16)
	m.G0 = v1269
	v1271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+338)))
	if v1271 != 0 {
		goto L199
	} else {
		goto L200
	}
L197:
	;
	goto L198
L198:
	;
	v1450 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+340)) = v1450
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+332)) = uint16(v1450)
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	if v1454 == v1450 {
		v1679 = v9
		v1686 = v9
		goto L217
	} else {
		goto L218
	}
L199:
	;
	m.G0 = v1269 + int32(16)
	goto L198
L200:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+144))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+20))
	v1275 = F_is_simple_union_all_recurse(m, v1273, v1272, v1274)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	if v1275 == int32(0) {
		goto L199
	} else {
		goto L202
	}
L202:
	;
	v1279 = v1273
	goto L203
L203:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+12))
	if v1320 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+52))
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)+12))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+4))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1325+v1326<<(uint(int32(2))%32)-int32(4))))
	v1333 = F_copyObjectImpl(m, v1332)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L1
	} else {
		goto L209
	}
L205:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	if v1321 == int32(142) {
		v1279 = v1320
		goto L203
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	goto L204
L208:
	;
	goto L207
L209:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+52))
	v1336 = F_lappend(m, v1335, v1333)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1272)+52)) = v1336
	if v1336 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1336)+4))
	v1341 = v1339
	goto L213
L212:
	;
	v1341 = int32(0)
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1320)+4)) = v1341
	v1343 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1332)+20)) = uint8(v1343)
	v1346 = F_palloc0(m, int32(8))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1346)+4)) = v1326
	*(*int32)(unsafe.Add(mBase, uint32(v1346))) = int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+8)) = v1346
	*(*int32)(unsafe.Add(mBase, uint32(v1269)+12)) = v1346
	v1356 = F_list_make1_impl(m, int32(1), v1269+int32(8))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v1358)+4)) = v1356
	v1360 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1272)+144)) = v1360
	F_pull_up_union_leaf_queries(m, v1273, v47, v1326, v1272, v1360)
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	goto L199
L217:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+140))
	if v1707 != 0 {
		goto L253
	} else {
		goto L254
	}
L218:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+4))
	if int32(0) < v1457 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1464 = int32(0)
	v1475 = v9
	v1482 = v9
	goto L222
L220:
	;
	v1562 = v9
	v1569 = v9
	goto L221
L221:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	if v1589 == int32(0) {
		v1679 = v1562
		v1686 = v1569
		goto L217
	} else {
		goto L238
	}
L222:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+12))
	v1503 = int32(2)
	v1505 = v1502 + v1464<<(uint(v1503)%32)
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1505)))
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+12))
	switch v1507 - v1503 {
	case 0:
		goto L227
	default:
		v1529 = v1475
		v1530 = v1482
		goto L224
	case 6:
		goto L226
	case 7:
		goto L225
	}
L223:
	;
	v1562 = v1529
	v1569 = v1530
	goto L221
L224:
	;
	v1531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506)+124)))
	if v1531 == int32(1) {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v1521 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1521)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+340)) = (v1505-v1522)>>(uint(int32(2))%32) + int32(1)
	v1529 = v1475
	v1530 = v1482
	goto L224
L226:
	;
	v1529 = v1475
	v1530 = int32(1)
	goto L224
L227:
	;
	v1510 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+332)) = uint8(v1510)
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+44))
	v1529 = base.B2i32(v1510<<(uint(v1513)%32)&int32(174) != int32(0)) | v1475
	v1530 = v1482
	goto L224
L228:
	;
	v1534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+333)) = uint8(v1534)
	goto L230
L229:
	;
	goto L230
L230:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+128))
	if v1536 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v47)+328))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1536)+4))
	if base.Ui32(v1538) < base.Ui32(v1537) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	goto L233
L233:
	;
	v1545 = v1464 + int32(1)
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1454)+4))
	if v1545 < v1546 {
		v1464 = v1545
		v1475 = v1529
		v1482 = v1530
		goto L222
	} else {
		goto L237
	}
L234:
	;
	v1540 = v1537
	goto L236
L235:
	;
	v1540 = v1538
	goto L236
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+328)) = v1540
	goto L233
L237:
	;
	goto L223
L238:
	;
	v1592 = int32(0)
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1589)+4))
	if v1593 <= v1592 {
		v1679 = v1562
		v1686 = v1569
		goto L217
	} else {
		goto L239
	}
L239:
	;
	v1597 = v1592
	goto L240
L240:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1589)+12))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1637+v1597<<(uint(int32(2))%32))))
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1641)+28))
	if v1642 == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1679 = v1562
	v1686 = v1569
	goto L217
L242:
	;
	v1662 = v1597 + int32(1)
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1589)+4))
	if v1662 < v1663 {
		v1597 = v1662
		goto L240
	} else {
		goto L250
	}
L243:
	;
	v1645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1641)+21)))
	if v1645 != int32(118) {
		goto L242
	} else {
		goto L244
	}
L244:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+56))
	v1649 = F_getRTEPermissionInfo(m, v1648, v1641)
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1651 = F_ExecCheckOneRelPerms(m, v1649)
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	if v1651 != 0 {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1649)+4))
	v1656 = F_get_rel_name(m, v1655)
	mBase = m.M
	v1657 = m.ExcPending
	if v1657 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_aclcheck_error(m, int32(1), int32(52), v1656)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L242
L250:
	;
	goto L241
L251:
	;
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+112))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+334)) = uint8(base.B2i32(v2100 != int32(0)))
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+76))
	v2106 = F_preprocess_expression(m, v47, v2104, int32(1))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L1
	} else {
		goto L306
	}
L252:
	;
	v1724 = int32(0)
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+60))
	v1728 = F_get_relids_in_jointree(m, v1725, v1724, v1724)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L258
	}
L253:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+12))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1708)))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1709)+8))
	F_CheckSelectLocking(m, v1706, v1710)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L1
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+4))
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v1713))|base.B2i32(int32(1)<<(uint(v1713)%32)&int32(52) == int32(0)) != 0 {
		goto L251
	} else {
		goto L257
	}
L256:
	;
	goto L252
L257:
	;
	goto L252
L258:
	;
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+32))
	if v1730 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v1731 = F_bms_del_member(m, v1728, v1730)
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	v1733 = v1728
	goto L261
L261:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+140))
	if v1734 == int32(0) {
		v1872 = v1733
		v1876 = v1724
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1733 = v1731
	goto L261
L263:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+52))
	if v1913 == int32(0) {
		v2021 = v1876
		goto L288
	} else {
		goto L289
	}
L264:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+4))
	if v1737 <= int32(0) {
		v1872 = v1733
		v1876 = v1724
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1741 = v1733
	v1742 = v1737
	v1744 = int32(0)
	v1745 = v1724
	goto L266
L266:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+52))
	v1783 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+12))
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+12))
	v1785 = int32(2)
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1784+v1744<<(uint(v1785)%32))))
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1783+v1789<<(uint(v1785)%32)-int32(4))))
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1795)+12))
	if v1796 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L1
	} else {
		goto L285
	}
L268:
	;
	goto L267
L269:
	;
	v1799 = F_bms_del_member(m, v1741, v1789)
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L1
	} else {
		goto L272
	}
L270:
	;
	v1849 = v1741
	v1850 = v1742
	v1851 = v1745
	goto L271
L271:
	;
	v1857 = v1744 + int32(1)
	if v1857 < v1850 {
		v1741 = v1849
		v1742 = v1850
		v1744 = v1857
		v1745 = v1851
		goto L266
	} else {
		goto L284
	}
L272:
	;
	v1802 = F_palloc0(m, int32(36))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1802))) = int32(378)
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+4)) = v1806
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+8)) = v1806
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+84))
	v1812 = v1810 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1809)+84)) = v1812
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+12)) = v1812
	v1815 = int32(5)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1795)+12))
	if v1816 != 0 {
		v1833 = v1815
		goto L274
	} else {
		goto L275
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+16)) = v1833
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+20)) = int32(1) << (uint(v1833) % 32)
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+24)) = v1840
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+12))
	v1843 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1802)+32)) = uint8(v1843)
	*(*int32)(unsafe.Add(mBase, uint32(v1802)+28)) = v1842
	v1846 = F_lappend(m, v1745, v1802)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L283
	}
L275:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1788)+8))
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1795)+21)))
	if v1818 == int32(102) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1795)+16))
	v1822 = F_GetFdwRoutineByRelId(m, v1821)
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L279
	}
L277:
	;
	goto L278
L278:
	;
	if base.Ui32(int32(5)) <= base.Ui32(v1817) {
		goto L268
	} else {
		goto L282
	}
L279:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+104))
	if v1824 == int32(0) {
		v1833 = v1815
		goto L274
	} else {
		goto L280
	}
L280:
	;
	v1827 = m.T0[v1824].(func(*base.Module, int32, int32) int32)(m, v1795, v1817)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v1833 = v1827
	goto L274
L282:
	;
	v1833 = int32(4) - v1817
	goto L274
L283:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(v1734)+4))
	v1849 = v1799
	v1850 = v1848
	v1851 = v1846
	goto L271
L284:
	;
	v1872 = v1849
	v1876 = v1851
	goto L263
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v1817
	F_errmsg_internal(m, int32(_a_F_subquery_planner_7), v44)
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(2855), int32(_a_F_subquery_planner_9))
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+144)) = v2021
	goto L251
L289:
	;
	v1916 = int32(0)
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+4))
	if v1917 <= v1916 {
		v2021 = v1876
		goto L288
	} else {
		goto L290
	}
L290:
	;
	v1921 = v1916
	v1924 = v1876
	goto L291
L291:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+12))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1961+v1921<<(uint(int32(2))%32))))
	v1967 = v1921 + int32(1)
	v1968 = F_bms_is_member(m, v1967, v1872)
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L1
	} else {
		goto L293
	}
L292:
	;
	v2021 = v2012
	goto L288
L293:
	;
	if v1968 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1971 = F_palloc0(m, int32(36))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L297
	}
L295:
	;
	v2012 = v1924
	goto L296
L296:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v1913)+4))
	if v1967 < v2015 {
		v1921 = v1967
		v1924 = v2012
		goto L291
	} else {
		goto L305
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+8)) = v1967
	*(*int32)(unsafe.Add(mBase, uint32(v1971))) = int32(378)
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+4)) = v1967
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1977)+84))
	v1980 = v1978 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1977)+84)) = v1980
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+12)) = v1980
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+12))
	if v1984 != 0 {
		v2000 = int32(5)
		goto L298
	} else {
		goto L299
	}
L298:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1971)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+16)) = v2000
	v2004 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1971)+32)) = uint8(v2004)
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+20)) = int32(1) << (uint(v2000) % 32)
	v2009 = F_lappend(m, v1924, v1971)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L1
	} else {
		goto L304
	}
L299:
	;
	v1986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1965)+21)))
	if v1986 != int32(102) {
		v2000 = int32(4)
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v1965)+16))
	v1991 = F_GetFdwRoutineByRelId(m, v1990)
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L1
	} else {
		goto L301
	}
L301:
	;
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1991)+104))
	if v1993 == int32(0) {
		v2000 = int32(5)
		goto L298
	} else {
		goto L302
	}
L302:
	;
	v1997 = m.T0[v1993].(func(*base.Module, int32, int32) int32)(m, v1965, int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L303
	}
L303:
	;
	v2000 = v1997
	goto L298
L304:
	;
	v2012 = v2009
	goto L296
L305:
	;
	goto L292
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+76)) = v2106
	v2109 = int32(0)
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+152))
	if v2110 == v2109 {
		v2175 = v2109
		goto L307
	} else {
		goto L308
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+152)) = v2175
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+96))
	v2219 = F_preprocess_expression(m, v47, v2217, int32(1))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L1
	} else {
		goto L318
	}
L308:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+4))
	if v2113 <= int32(0) {
		v2175 = v2109
		goto L307
	} else {
		goto L309
	}
L309:
	;
	v2117 = v2109
	v2120 = int32(0)
	goto L310
L310:
	;
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+12))
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v2158+v2120<<(uint(int32(2))%32))))
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v2162)+16))
	v2165 = F_preprocess_expression(m, v47, v2163, int32(0))
	mBase = m.M
	v2166 = m.ExcPending
	if v2166 != 0 {
		goto L1
	} else {
		goto L312
	}
L311:
	;
	v2175 = v2170
	goto L307
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2162)+16)) = v2165
	if v2165 != 0 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	v2168 = F_lappend(m, v2117, v2162)
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L1
	} else {
		goto L316
	}
L314:
	;
	v2170 = v2117
	goto L315
L315:
	;
	v2172 = v2120 + int32(1)
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v2110)+4))
	if v2172 < v2173 {
		v2117 = v2170
		v2120 = v2172
		goto L310
	} else {
		goto L317
	}
L316:
	;
	v2170 = v2168
	goto L315
L317:
	;
	goto L311
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+96)) = v2219
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+60))
	F_preprocess_qual_conditions(m, v47, v2222)
	mBase = m.M
	v2224 = m.ExcPending
	if v2224 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+112))
	v2227 = F_preprocess_expression(m, v47, v2225, int32(0))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L1
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+112)) = v2227
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+116))
	if v2230 == int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+128))
	v2340 = F_preprocess_expression(m, v47, v2338, int32(6))
	mBase = m.M
	v2341 = m.ExcPending
	if v2341 != 0 {
		goto L1
	} else {
		goto L329
	}
L322:
	;
	v2233 = int32(0)
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+4))
	if v2234 <= v2233 {
		goto L321
	} else {
		goto L323
	}
L323:
	;
	v2240 = v2233
	goto L324
L324:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+12))
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2278+v2240<<(uint(int32(2))%32))))
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+24))
	v2285 = F_preprocess_expression(m, v47, v2283, int32(6))
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L1
	} else {
		goto L326
	}
L325:
	;
	goto L321
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2282)+24)) = v2285
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v2282)+28))
	v2290 = F_preprocess_expression(m, v47, v2288, int32(6))
	mBase = m.M
	v2291 = m.ExcPending
	if v2291 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2282)+28)) = v2290
	v2294 = v2240 + int32(1)
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2230)+4))
	if v2294 < v2295 {
		v2240 = v2294
		goto L324
	} else {
		goto L328
	}
L328:
	;
	goto L325
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+128)) = v2340
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+132))
	v2345 = F_preprocess_expression(m, v47, v2343, int32(6))
	mBase = m.M
	v2346 = m.ExcPending
	if v2346 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+132)) = v2345
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	if v2348 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2348)+8))
	v2351 = F_preprocess_expression(m, v47, v2349, int32(10))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L1
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+64))
	if v2377 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L334:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2353)+8)) = v2351
	v2355 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2355)+12))
	v2358 = F_preprocess_expression(m, v47, v2356, int32(0))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2360)+12)) = v2358
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	v2363 = *(*int32)(unsafe.Add(mBase, uint32(v2362)+24))
	v2365 = F_preprocess_expression(m, v47, v2363, int32(1))
	mBase = m.M
	v2366 = m.ExcPending
	if v2366 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	v2367 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2367)+24)) = v2365
	v2369 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2369)+28))
	v2372 = F_preprocess_expression(m, v47, v2370, int32(0))
	mBase = m.M
	v2373 = m.ExcPending
	if v2373 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v2374)+28)) = v2372
	goto L333
L338:
	;
	v2485 = int32(0)
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+72))
	v2488 = F_preprocess_expression(m, v47, v2486, v2485)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L1
	} else {
		goto L346
	}
L339:
	;
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2377)+4))
	if v2380 <= int32(0) {
		goto L338
	} else {
		goto L340
	}
L340:
	;
	v2387 = int32(0)
	goto L341
L341:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v2377)+12))
	v2429 = *(*int32)(unsafe.Add(mBase, uint32(v2425+v2387<<(uint(int32(2))%32))))
	v2430 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+20))
	v2432 = F_preprocess_expression(m, v47, v2430, int32(1))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L1
	} else {
		goto L343
	}
L342:
	;
	goto L338
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2429)+20)) = v2432
	v2435 = *(*int32)(unsafe.Add(mBase, uint32(v2429)+16))
	v2437 = F_preprocess_expression(m, v47, v2435, int32(0))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2429)+16)) = v2437
	v2441 = v2387 + int32(1)
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2377)+4))
	if v2441 < v2442 {
		v2387 = v2441
		goto L341
	} else {
		goto L345
	}
L345:
	;
	goto L342
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+72)) = v2488
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v2493 = F_preprocess_expression(m, v47, v2491, int32(7))
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v2493
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	if v2496 == int32(0) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+332)))
	if v2786 == int32(0) {
		goto L403
	} else {
		goto L404
	}
L349:
	;
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2496)+4))
	if v2499 <= int32(0) {
		goto L348
	} else {
		goto L350
	}
L350:
	;
	v2506 = v2485
	goto L351
L351:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2496)+12))
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v2543+v2506<<(uint(int32(2))%32))))
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+12))
	switch v2548 {
	case 0:
		goto L359
	case 1:
		goto L358
	default:
		goto L353
	case 3:
		goto L357
	case 4:
		goto L356
	case 5:
		goto L355
	case 9:
		goto L354
	}
L352:
	;
	goto L348
L353:
	;
	v2639 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+128))
	if v2639 == int32(0) {
		goto L395
	} else {
		goto L396
	}
L354:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+120))
	v2634 = F_preprocess_expression(m, v47, v2632, int32(13))
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L1
	} else {
		goto L394
	}
L355:
	;
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+80))
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+124)))
	if v2627 != 0 {
		goto L390
	} else {
		goto L391
	}
L356:
	;
	v2616 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+76))
	v2619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+124)))
	if v2619 != 0 {
		goto L386
	} else {
		goto L387
	}
L357:
	;
	v2608 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+68))
	v2611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+124)))
	if v2611 != 0 {
		goto L382
	} else {
		goto L383
	}
L358:
	;
	v2582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+124)))
	if v2582 != int32(1) {
		goto L353
	} else {
		goto L374
	}
L359:
	;
	v2549 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+32))
	if v2549 == int32(0) {
		goto L353
	} else {
		goto L360
	}
L360:
	;
	v2552 = F_eval_const_expressions(m, v47, v2549)
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L1
	} else {
		goto L361
	}
L361:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2554)+39)))
	if v2555 != int32(1) {
		v2574 = v2552
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if base.Ui32(int32(2)) <= base.Ui32(v2575) {
		goto L370
	} else {
		goto L371
	}
L363:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v2558)+80))
	if v2559 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v47
	v2565 = F_preprocess_subquery_phvs_walker(m, v2552, v44+int32(24))
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L1
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v2572 = F_SS_process_sublinks(m, v47, v2552, int32(0))
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L1
	} else {
		goto L369
	}
L367:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2567)+39)))
	if v2568 != int32(1) {
		v2574 = v2552
		goto L362
	} else {
		goto L368
	}
L368:
	;
	goto L366
L369:
	;
	v2574 = v2572
	goto L362
L370:
	;
	v2578 = F_SS_replace_correlation_vars(m, v47, v2574)
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L1
	} else {
		goto L373
	}
L371:
	;
	v2580 = v2574
	goto L372
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+32)) = v2580
	goto L353
L373:
	;
	v2580 = v2578
	goto L372
L374:
	;
	v2585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+332)))
	if v2585 == int32(1) {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v2588 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+36))
	v2590 = F_flatten_join_alias_vars(m, v47, v2588, v2589)
	mBase = m.M
	v2591 = m.ExcPending
	if v2591 != 0 {
		goto L1
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v2597 = *(*int32)(unsafe.Add(mBase, uint32(v2596)+80))
	if v2597 == int32(0) {
		goto L353
	} else {
		goto L380
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+36)) = v2590
	v2593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547)+124)))
	if v2593 != int32(1) {
		goto L353
	} else {
		goto L379
	}
L379:
	;
	goto L377
L380:
	;
	v2600 = *(*int32)(unsafe.Add(mBase, uint32(v2547)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v47
	v2606 = F_preprocess_subquery_phvs_walker(m, v2600, v44+int32(24))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	goto L353
L382:
	;
	v2612 = int32(3)
	goto L384
L383:
	;
	v2612 = int32(2)
	goto L384
L384:
	;
	v2613 = F_preprocess_expression(m, v47, v2608, v2612)
	mBase = m.M
	v2614 = m.ExcPending
	if v2614 != 0 {
		goto L1
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+68)) = v2613
	goto L353
L386:
	;
	v2620 = int32(12)
	goto L388
L387:
	;
	v2620 = int32(11)
	goto L388
L388:
	;
	v2621 = F_preprocess_expression(m, v47, v2616, v2620)
	mBase = m.M
	v2622 = m.ExcPending
	if v2622 != 0 {
		goto L1
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+76)) = v2621
	goto L353
L390:
	;
	v2628 = int32(5)
	goto L392
L391:
	;
	v2628 = int32(4)
	goto L392
L392:
	;
	v2629 = F_preprocess_expression(m, v47, v2624, v2628)
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L1
	} else {
		goto L393
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+80)) = v2629
	goto L353
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2547)+120)) = v2634
	goto L353
L395:
	;
	v2742 = v2506 + int32(1)
	v2743 = *(*int32)(unsafe.Add(mBase, uint32(v2496)+4))
	if v2742 < v2743 {
		v2506 = v2742
		goto L351
	} else {
		goto L402
	}
L396:
	;
	v2642 = int32(0)
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+4))
	if v2643 <= v2642 {
		goto L395
	} else {
		goto L397
	}
L397:
	;
	v2647 = v2642
	goto L398
L398:
	;
	v2687 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+12))
	v2690 = v2687 + v2647<<(uint(int32(2))%32)
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2690)))
	v2693 = F_preprocess_expression(m, v47, v2691, int32(0))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L1
	} else {
		goto L400
	}
L399:
	;
	goto L395
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2690))) = v2693
	v2697 = v2647 + int32(1)
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+4))
	if v2697 < v2698 {
		v2647 = v2697
		goto L398
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	goto L352
L403:
	;
	v2889 = int32(0)
	v2890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+45)))
	if v2890 != int32(1) {
		v3015 = v2889
		goto L410
	} else {
		goto L411
	}
L404:
	;
	v2789 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+52))
	if v2789 == int32(0) {
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2789)+4))
	if v2792 <= int32(0) {
		goto L403
	} else {
		goto L406
	}
L406:
	;
	v2797 = int32(0)
	goto L407
L407:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2789)+12))
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2837+v2797<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2841)+52)) = int32(0)
	v2845 = v2797 + int32(1)
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2789)+4))
	if v2845 < v2846 {
		v2797 = v2845
		goto L407
	} else {
		goto L409
	}
L408:
	;
	goto L403
L409:
	;
	goto L408
L410:
	;
	v3056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+38)))
	if v3056 == int32(1) {
		goto L426
	} else {
		goto L427
	}
L411:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+112))
	if v2893 == int32(0) {
		v2964 = v2889
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v3005 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3006 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+76))
	v3007 = F_flatten_group_exprs(m, v47, v3005, v3006)
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L424
	}
L413:
	;
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v47)+340))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+28)) = v2896
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v1101
	v2899 = *(*int32)(unsafe.Add(mBase, uint32(v2893)+4))
	if v2899 <= int32(0) {
		v2964 = v2889
		goto L412
	} else {
		goto L414
	}
L414:
	;
	v2903 = v2889
	v2904 = int32(0)
	goto L415
L415:
	;
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(v2893)+12))
	v2948 = *(*int32)(unsafe.Add(mBase, uint32(v2944+v2904<<(uint(int32(2))%32))))
	v2952 = F_expression_has_grouping_conflict(m, v2948, int32(878), v44+int32(24))
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L1
	} else {
		goto L417
	}
L416:
	;
	v2961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+45)))
	if v2961 != int32(1) {
		v3015 = v2956
		goto L410
	} else {
		goto L423
	}
L417:
	;
	if v2952 != 0 {
		goto L418
	} else {
		goto L419
	}
L418:
	;
	v2954 = F_bms_add_member(m, v2903, v2904)
	mBase = m.M
	v2955 = m.ExcPending
	if v2955 != 0 {
		goto L1
	} else {
		goto L421
	}
L419:
	;
	v2956 = v2903
	goto L420
L420:
	;
	v2958 = v2904 + int32(1)
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2893)+4))
	if v2958 < v2959 {
		v2903 = v2956
		v2904 = v2958
		goto L415
	} else {
		goto L422
	}
L421:
	;
	v2956 = v2954
	goto L420
L422:
	;
	goto L416
L423:
	;
	v2964 = v2956
	goto L412
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+76)) = v3007
	v3010 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+112))
	v3012 = F_flatten_group_exprs(m, v47, v3010, v3011)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+112)) = v3012
	v3015 = v2964
	goto L410
L426:
	;
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+76))
	v3060 = F_expression_returns_set(m, v3059)
	mBase = m.M
	v3061 = m.ExcPending
	if v3061 != 0 {
		goto L1
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+108))
	if v3063 != 0 {
		goto L430
	} else {
		goto L431
	}
L429:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1101)+38)) = uint8(v3060)
	goto L428
L430:
	;
	v3064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101)+104)))
	v3066 = F_expand_grouping_sets(m, v3063, v3064, int32(-1))
	mBase = m.M
	v3067 = m.ExcPending
	if v3067 != 0 {
		goto L1
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+112))
	if v3069 == int32(0) {
		goto L435
	} else {
		goto L436
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+108)) = v3066
	goto L432
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1101)+112)) = v3189
	if v1679&int32(1) != 0 {
		goto L470
	} else {
		goto L471
	}
L435:
	;
	v3189 = int32(0)
	goto L434
L436:
	;
	goto L437
L437:
	;
	v3073 = int32(0)
	v3074 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+4))
	if v3074 <= v3073 {
		v3189 = v3073
		goto L434
	} else {
		goto L438
	}
L438:
	;
	v3079 = int32(0)
	v3082 = v3073
	goto L439
L439:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+12))
	v3123 = *(*int32)(unsafe.Add(mBase, uint32(v3119+v3079<<(uint(int32(2))%32))))
	v3124 = F_contain_agg_clause(m, v3123)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L1
	} else {
		goto L443
	}
L440:
	;
	v3189 = v3179
	goto L434
L441:
	;
	v3182 = v3079 + int32(1)
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3069)+4))
	if v3182 < v3183 {
		v3079 = v3182
		v3082 = v3179
		goto L439
	} else {
		goto L467
	}
L442:
	;
	v3176 = F_lappend(m, v3082, v3123)
	mBase = m.M
	v3177 = m.ExcPending
	if v3177 != 0 {
		goto L1
	} else {
		goto L466
	}
L443:
	;
	if v3124 != 0 {
		goto L442
	} else {
		goto L444
	}
L444:
	;
	v3126 = F_contain_volatile_functions(m, v3123)
	mBase = m.M
	v3127 = m.ExcPending
	if v3127 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	if v3126 != 0 {
		goto L442
	} else {
		goto L446
	}
L446:
	;
	v3128 = F_contain_subplans(m, v3123)
	mBase = m.M
	v3129 = m.ExcPending
	if v3129 != 0 {
		goto L1
	} else {
		goto L447
	}
L447:
	;
	if v3128 != 0 {
		goto L442
	} else {
		goto L448
	}
L448:
	;
	v3130 = F_bms_is_member(m, v3079, v3015)
	mBase = m.M
	v3131 = m.ExcPending
	if v3131 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	if v3130 != 0 {
		goto L442
	} else {
		goto L450
	}
L450:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+100))
	if v3132 == int32(0) {
		goto L451
	} else {
		goto L452
	}
L451:
	;
	v3164 = F_copyObjectImpl(m, v3123)
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L1
	} else {
		goto L463
	}
L452:
	;
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+108))
	if v3135 == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v3155 = F_preprocess_expression(m, v47, v3123, int32(0))
	mBase = m.M
	v3156 = m.ExcPending
	if v3156 != 0 {
		goto L1
	} else {
		goto L461
	}
L454:
	;
	v3138 = *(*int32)(unsafe.Add(mBase, uint32(v47)+340))
	v3139 = F_pull_varnos(m, v47, v3123)
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	v3141 = F_bms_is_member(m, v3138, v3139)
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L1
	} else {
		goto L456
	}
L456:
	;
	if v3141 != 0 {
		goto L442
	} else {
		goto L457
	}
L457:
	;
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+100))
	if v3143 == int32(0) {
		goto L451
	} else {
		goto L458
	}
L458:
	;
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+108))
	if v3146 == int32(0) {
		goto L453
	} else {
		goto L459
	}
L459:
	;
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(v3146)+12))
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3149)))
	if v3150 == int32(0) {
		goto L451
	} else {
		goto L460
	}
L460:
	;
	goto L453
L461:
	;
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+60))
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v3157)+8))
	v3159 = F_list_concat(m, v3158, v3155)
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		goto L1
	} else {
		goto L462
	}
L462:
	;
	v3161 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v3161)+8)) = v3159
	v3179 = v3082
	goto L441
L463:
	;
	v3167 = F_preprocess_expression(m, v47, v3164, int32(0))
	mBase = m.M
	v3168 = m.ExcPending
	if v3168 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+60))
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3169)+8))
	v3171 = F_list_concat(m, v3170, v3167)
	mBase = m.M
	v3172 = m.ExcPending
	if v3172 != 0 {
		goto L1
	} else {
		goto L465
	}
L465:
	;
	v3173 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v3173)+8)) = v3171
	goto L442
L466:
	;
	v3179 = v3176
	goto L441
L467:
	;
	goto L440
L468:
	;
	v3637 = int32(0)
	v3640 = m.G0
	v3642 = v3640 - int32(368)
	m.G0 = v3642
	v3644 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3645 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+132))
	if v3645 == v3637 {
		goto L522
	} else {
		goto L523
	}
L469:
	;
	v3441 = int32(0)
	v3442 = m.G0
	v3444 = v3442 - int32(16)
	m.G0 = v3444
	*(*int32)(unsafe.Add(mBase, uint32(v3444)+12)) = v3441
	v3448 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v3449 = *(*int32)(unsafe.Add(mBase, uint32(v3448)+80))
	if v3449 != 0 {
		goto L500
	} else {
		goto L501
	}
L470:
	;
	v3229 = m.G0
	v3231 = v3229 - int32(16)
	m.G0 = v3231
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3233)+60))
	v3235 = F_reduce_outer_joins_pass1(m, v3234)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L1
	} else {
		goto L474
	}
L471:
	;
	goto L472
L472:
	;
	if v1686 == int32(0) {
		goto L468
	} else {
		goto L499
	}
L473:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		goto L1
	} else {
		goto L496
	}
L474:
	;
	if v3235 == int32(0) {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v3239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3235)+4)))
	if v3239 == int32(0) {
		goto L473
	} else {
		goto L476
	}
L476:
	;
	v3242 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+12)) = v3242
	*(*int64)(unsafe.Add(mBase, uint32(v3231)+4)) = int64(0)
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3246)+60))
	F_reduce_outer_joins_pass2(m, v3247, v3235, v3231+int32(4), v47, v3242, v3242)
	mBase = m.M
	v3253 = m.ExcPending
	if v3253 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v3231)+4))
	if v3254 != 0 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3257 = F_remove_nulling_relids(m, v3255, v3254, int32(0))
	mBase = m.M
	v3258 = m.ExcPending
	if v3258 != 0 {
		goto L1
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v3231)+8))
	if v3266 == int32(0) {
		goto L483
	} else {
		goto L484
	}
L481:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3257
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3231)+4))
	v3263 = F_remove_nulling_relids(m, v3260, v3261, int32(0))
	mBase = m.M
	v3264 = m.ExcPending
	if v3264 != 0 {
		goto L1
	} else {
		goto L482
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v3263
	goto L480
L483:
	;
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3231)+12))
	if v3377 != 0 {
		goto L492
	} else {
		goto L493
	}
L484:
	;
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v3266)+4))
	if v3269 <= int32(0) {
		goto L483
	} else {
		goto L485
	}
L485:
	;
	v3273 = int32(0)
	goto L486
L486:
	;
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(v3266)+12))
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v3314+v3273<<(uint(int32(2))%32))))
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3318)))
	v3320 = F_bms_make_singleton(m, v3319)
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L1
	} else {
		goto L488
	}
L487:
	;
	goto L483
L488:
	;
	v3322 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3323 = *(*int32)(unsafe.Add(mBase, uint32(v3318)+4))
	v3324 = F_remove_nulling_relids(m, v3322, v3320, v3323)
	mBase = m.M
	v3325 = m.ExcPending
	if v3325 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3324
	v3327 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v3328 = *(*int32)(unsafe.Add(mBase, uint32(v3318)+4))
	v3329 = F_remove_nulling_relids(m, v3327, v3320, v3328)
	mBase = m.M
	v3330 = m.ExcPending
	if v3330 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v3329
	v3333 = v3273 + int32(1)
	v3334 = *(*int32)(unsafe.Add(mBase, uint32(v3266)+4))
	if v3333 < v3334 {
		v3273 = v3333
		goto L486
	} else {
		goto L491
	}
L491:
	;
	goto L487
L492:
	;
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3378)+60))
	F_remove_redundant_nullability_quals(m, v3379, v3377)
	mBase = m.M
	v3381 = m.ExcPending
	if v3381 != 0 {
		goto L1
	} else {
		goto L495
	}
L493:
	;
	goto L494
L494:
	;
	m.G0 = v3231 + int32(16)
	goto L469
L495:
	;
	goto L494
L496:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_10), int32(0))
	mBase = m.M
	v3392 = m.ExcPending
	if v3392 != 0 {
		goto L1
	} else {
		goto L497
	}
L497:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_5), int32(3273), int32(_a_F_subquery_planner_11))
	mBase = m.M
	v3397 = m.ExcPending
	if v3397 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L499:
	;
	goto L469
L500:
	;
	v3450 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(v3450)+60))
	v3452 = int32(0)
	v3454 = F_get_relids_in_jointree(m, v3451, v3452, v3452)
	mBase = m.M
	v3455 = m.ExcPending
	if v3455 != 0 {
		goto L1
	} else {
		goto L503
	}
L501:
	;
	v3456 = v3441
	goto L502
L502:
	;
	v3457 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3458 = *(*int32)(unsafe.Add(mBase, uint32(v3457)+60))
	v3462 = F_remove_useless_results_recurse(m, v47, v3458, v3456, int32(0), v3444+int32(12))
	mBase = m.M
	v3463 = m.ExcPending
	if v3463 != 0 {
		goto L1
	} else {
		goto L504
	}
L503:
	;
	v3456 = v3454
	goto L502
L504:
	;
	v3464 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3464)+60)) = v3462
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(v3444)+12))
	if v3466 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v3467 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3469 = F_remove_nulling_relids(m, v3467, v3466, int32(0))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L1
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v3478 = *(*int32)(unsafe.Add(mBase, uint32(v47)+144))
	if v3478 == int32(0) {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v3469
	v3472 = *(*int32)(unsafe.Add(mBase, uint32(v47)+136))
	v3473 = *(*int32)(unsafe.Add(mBase, uint32(v3444)+12))
	v3475 = F_remove_nulling_relids(m, v3472, v3473, int32(0))
	mBase = m.M
	v3476 = m.ExcPending
	if v3476 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+136)) = v3475
	goto L507
L510:
	;
	m.G0 = v3444 + int32(16)
	goto L468
L511:
	;
	v3482 = int32(0)
	v3483 = v3478
	goto L512
L512:
	;
	v3523 = *(*int32)(unsafe.Add(mBase, uint32(v3483)+4))
	if v3523 <= v3482 {
		goto L510
	} else {
		goto L514
	}
L513:
	;
	goto L510
L514:
	;
	v3525 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(v3525)+52))
	v3527 = *(*int32)(unsafe.Add(mBase, uint32(v3526)+12))
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(v3483)+12))
	v3529 = int32(2)
	v3532 = *(*int32)(unsafe.Add(mBase, uint32(v3528+v3482<<(uint(v3529)%32))))
	v3533 = *(*int32)(unsafe.Add(mBase, uint32(v3532)+4))
	v3539 = *(*int32)(unsafe.Add(mBase, uint32(v3527+v3533<<(uint(v3529)%32)-int32(4))))
	v3540 = *(*int32)(unsafe.Add(mBase, uint32(v3539)+12))
	if v3540 == int32(8) {
		goto L515
	} else {
		goto L516
	}
L515:
	;
	v3543 = F_list_delete_nth_cell(m, v3483, v3482)
	mBase = m.M
	v3544 = m.ExcPending
	if v3544 != 0 {
		goto L1
	} else {
		goto L518
	}
L516:
	;
	v3548 = v3483
	v3549 = v3482
	goto L517
L517:
	;
	if v3548 != 0 {
		v3482 = v3549 + int32(1)
		v3483 = v3548
		goto L512
	} else {
		goto L519
	}
L518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+144)) = v3543
	v3548 = v3543
	v3549 = v3482 - int32(1)
	goto L517
L519:
	;
	goto L513
L520:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v47)+312)) = v3754
	v3760 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+144))
	if v3760 != 0 {
		goto L595
	} else {
		goto L596
	}
L521:
	;
	v3667 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+128))
	if v3667 == int32(0) {
		v3686 = v39
		goto L533
	} else {
		goto L534
	}
L522:
	;
	v3648 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+128))
	if v3648 != 0 {
		v3666 = v39
		goto L521
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	v3650 = int64(-1)
	v3651 = F_estimate_expression_value(m, v47, v3645)
	mBase = m.M
	v3652 = m.ExcPending
	if v3652 != 0 {
		goto L1
	} else {
		goto L526
	}
L525:
	;
	v3754 = l6
	v3756 = float64(-1)
	v3757 = v39
	v3758 = v39
	goto L520
L526:
	;
	if v3651 == int32(0) {
		v3666 = v3650
		goto L521
	} else {
		goto L527
	}
L527:
	;
	v3655 = *(*int32)(unsafe.Add(mBase, uint32(v3651)))
	if v3655 != int32(7) {
		v3666 = v3650
		goto L521
	} else {
		goto L528
	}
L528:
	;
	v3659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3651)+32)))
	if v3659 != 0 {
		v3666 = int64(0)
		goto L521
	} else {
		goto L529
	}
L529:
	;
	v3660 = int64(1)
	v3661 = *(*int64)(unsafe.Add(mBase, uint32(v3651)+24))
	if v3661 <= v3660 {
		goto L530
	} else {
		goto L531
	}
L530:
	;
	v3664 = v3660
	goto L532
L531:
	;
	v3664 = v3661
	goto L532
L532:
	;
	v3666 = v3664
	goto L521
L533:
	;
	if v3666 != int64(0) {
		goto L544
	} else {
		goto L545
	}
L534:
	;
	v3670 = int64(-1)
	v3671 = F_estimate_expression_value(m, v47, v3667)
	mBase = m.M
	v3672 = m.ExcPending
	if v3672 != 0 {
		goto L1
	} else {
		goto L535
	}
L535:
	;
	if v3671 == int32(0) {
		v3686 = v3670
		goto L533
	} else {
		goto L536
	}
L536:
	;
	v3675 = *(*int32)(unsafe.Add(mBase, uint32(v3671)))
	if v3675 != int32(7) {
		v3686 = v3670
		goto L533
	} else {
		goto L537
	}
L537:
	;
	v3679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3671)+32)))
	if v3679 != 0 {
		v3686 = int64(0)
		goto L533
	} else {
		goto L538
	}
L538:
	;
	v3680 = *(*int64)(unsafe.Add(mBase, uint32(v3671)+24))
	v3681 = int64(0)
	if v3681 < v3680 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v3684 = v3680
	goto L541
L540:
	;
	v3684 = v3681
	goto L541
L541:
	;
	v3686 = v3684
	goto L533
L542:
	;
	if int64(0) <= v3686 {
		goto L576
	} else {
		goto L577
	}
L543:
	;
	if base.F64_lt(l6, v3696) != 0 {
		goto L573
	} else {
		goto L574
	}
L544:
	;
	v3692 = base.F64_add(base.F64_convert_i64_u(v3666), base.F64_convert_i64_u(v3686))
	if v3686|v3666 < int64(0) {
		goto L547
	} else {
		goto L548
	}
L545:
	;
	goto L546
L546:
	;
	v3713 = float64(-1)
	v3714 = int64(0)
	if base.B2i32(base.F64_gt(l6, float64(0)) == int32(0))|base.B2i32(v3686 == v3714) != 0 {
		v3754 = l6
		v3756 = v3713
		v3757 = v3686
		v3758 = v3714
		goto L520
	} else {
		goto L561
	}
L547:
	;
	v3696 = float64(0.1)
	goto L549
L548:
	;
	v3696 = v3692
	goto L549
L549:
	;
	if base.F64_ge(l6, float64(1)) != 0 {
		goto L550
	} else {
		goto L551
	}
L550:
	;
	if base.F64_ge(v3696, float64(1)) == int32(0) {
		v3744 = l6
		goto L542
	} else {
		goto L553
	}
L551:
	;
	goto L552
L552:
	;
	if base.F64_gt(l6, float64(0)) == int32(0) {
		goto L557
	} else {
		goto L558
	}
L553:
	;
	if base.F64_lt(l6, v3696) != 0 {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	v3704 = l6
	goto L556
L555:
	;
	v3704 = v3696
	goto L556
L556:
	;
	v3744 = v3704
	goto L542
L557:
	;
	v3744 = v3696
	goto L542
L558:
	;
	goto L559
L559:
	;
	if base.F64_ge(v3696, float64(1)) == int32(0) {
		goto L543
	} else {
		goto L560
	}
L560:
	;
	v3744 = v3696
	goto L542
L561:
	;
	if v3686 < int64(0) {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v3726 = float64(0.1)
	goto L564
L563:
	;
	v3726 = base.F64_convert_i64_u(v3686)
	goto L564
L564:
	;
	if base.F64_ge(l6, float64(1)) != 0 {
		goto L565
	} else {
		goto L566
	}
L565:
	;
	if base.F64_ge(v3726, float64(1)) == int32(0) {
		goto L568
	} else {
		goto L569
	}
L566:
	;
	goto L567
L567:
	;
	if base.F64_ge(v3726, float64(1)) != 0 {
		v3754 = l6
		v3756 = v3713
		v3757 = v3686
		v3758 = v3714
		goto L520
	} else {
		goto L571
	}
L568:
	;
	v3754 = v3726
	v3756 = v3713
	v3757 = v3686
	v3758 = v3714
	goto L520
L569:
	;
	goto L570
L570:
	;
	v3754 = base.F64_add(l6, v3726)
	v3756 = v3713
	v3757 = v3686
	v3758 = v3714
	goto L520
L571:
	;
	v3736 = base.F64_add(l6, v3726)
	if base.F64_ge(v3736, float64(1)) == int32(0) {
		v3754 = v3736
		v3756 = v3713
		v3757 = v3686
		v3758 = v3714
		goto L520
	} else {
		goto L572
	}
L572:
	;
	v3754 = float64(0)
	v3756 = v3713
	v3757 = v3686
	v3758 = v3714
	goto L520
L573:
	;
	v3743 = l6
	goto L575
L574:
	;
	v3743 = v3696
	goto L575
L575:
	;
	v3744 = v3743
	goto L542
L576:
	;
	v3748 = v3692
	goto L578
L577:
	;
	v3748 = float64(-1)
	goto L578
L578:
	;
	if int64(0) < v3666 {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v3752 = v3748
	goto L581
L580:
	;
	v3752 = float64(-1)
	goto L581
L581:
	;
	v3754 = v3744
	v3756 = v3752
	v3757 = v3686
	v3758 = v3666
	goto L520
L582:
	;
	F_SS_identify_outer_params(m, v16518)
	mBase = m.M
	v16786 = m.ExcPending
	if v16786 != 0 {
		goto L1
	} else {
		goto L2588
	}
L583:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v16768 = m.ExcPending
	if v16768 != 0 {
		goto L1
	} else {
		goto L2583
	}
L584:
	;
	v14948 = *(*int32)(unsafe.Add(mBase, uint32(v14920)+124))
	if v14948 == int32(0) {
		goto L2194
	} else {
		goto L2195
	}
L585:
	;
	v8260 = int32(0)
	v8264 = m.G0
	v8266 = v8264 - int32(80)
	m.G0 = v8266
	v8268 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v8269 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+4))
	v8270 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+52))
	v8271 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+32))
	if v8271 != 0 {
		goto L1134
	} else {
		goto L1135
	}
L586:
	;
	v7905 = *(*int32)(unsafe.Add(mBase, uint32(v7881)+28))
	if v7905 == int32(0) {
		v8227 = v7872
		v8230 = v7875
		v8232 = v7877
		v8236 = v7881
		v8239 = v7884
		v8242 = v7887
		v8248 = v7893
		v8253 = v7898
		v8257 = v7902
		v8258 = v7903
		goto L585
	} else {
		goto L1108
	}
L587:
	;
	if v6799 == int32(0) {
		v7872 = v47
		v7875 = v3642
		v7877 = v3644
		v7879 = v4420
		v7881 = v4422
		v7884 = l7
		v7887 = v44
		v7893 = v9
		v7898 = v3756
		v7902 = v3757
		v7903 = v3758
		goto L586
	} else {
		goto L994
	}
L588:
	;
	v4867 = int32(1)
	v4869 = v4768 << (uint(int32(2)) % 32)
	v4870 = F_palloc0(m, v4869)
	mBase = m.M
	v4871 = m.ExcPending
	if v4871 != 0 {
		goto L1
	} else {
		goto L774
	}
L589:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+88)) = v4822
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+208)) = v4822
	v4865 = F_list_make1_impl(m, int32(1), v3642+int32(88))
	mBase = m.M
	v4866 = m.ExcPending
	if v4866 != 0 {
		goto L1
	} else {
		goto L773
	}
L590:
	;
	v4763 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+12))
	if v4763 == int32(0) {
		v4822 = v4725
		goto L589
	} else {
		goto L768
	}
L591:
	;
	if v4649 == int32(0) {
		v7872 = v47
		v7875 = v3642
		v7877 = v3644
		v7879 = v4420
		v7881 = v4422
		v7884 = l7
		v7887 = v44
		v7893 = v9
		v7898 = v3756
		v7902 = v3757
		v7903 = v3758
		goto L586
	} else {
		goto L767
	}
L592:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4692 = m.ExcPending
	if v4692 != 0 {
		goto L1
	} else {
		goto L759
	}
L593:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L1
	} else {
		goto L756
	}
L594:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4666 = m.ExcPending
	if v4666 != 0 {
		goto L1
	} else {
		goto L753
	}
L595:
	;
	v3762 = m.G0
	v3764 = v3762 - int32(32)
	m.G0 = v3764
	v3766 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v3767 = *(*int32)(unsafe.Add(mBase, uint32(v3766)+144))
	v3768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+100)) = uint8(v3768)
	F_setup_simple_rel_arrays(m, v47)
	mBase = m.M
	v3771 = m.ExcPending
	if v3771 != 0 {
		goto L1
	} else {
		goto L599
	}
L596:
	;
	goto L597
L597:
	;
	v4419 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+108))
	if v4419 != 0 {
		goto L705
	} else {
		goto L706
	}
L598:
	;
	v4276 = *(*int32)(unsafe.Add(mBase, uint32(v47)+284))
	v4277 = F_copyObjectImpl(m, v4276)
	mBase = m.M
	v4278 = m.ExcPending
	if v4278 != 0 {
		goto L1
	} else {
		goto L684
	}
L599:
	;
	v3773 = v3767
	goto L600
L600:
	;
	v3813 = *(*int32)(unsafe.Add(mBase, uint32(v3773)+12))
	if v3813 != 0 {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(v47)+44))
	v3818 = *(*int32)(unsafe.Add(mBase, uint32(v3813)+4))
	v3822 = *(*int32)(unsafe.Add(mBase, uint32(v3817+v3818<<(uint(int32(2))%32))))
	v3823 = *(*int32)(unsafe.Add(mBase, uint32(v3822)+36))
	v3824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+338)))
	if v3824 == int32(1) {
		goto L609
	} else {
		goto L610
	}
L602:
	;
	v3814 = *(*int32)(unsafe.Add(mBase, uint32(v3813)))
	if v3814 == int32(142) {
		v3773 = v3813
		goto L600
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	goto L601
L605:
	;
	goto L604
L606:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4259 = m.ExcPending
	if v4259 != 0 {
		goto L1
	} else {
		goto L679
	}
L607:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4246 = m.ExcPending
	if v4246 != 0 {
		goto L1
	} else {
		goto L676
	}
L608:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+284)) = v4203
	m.G0 = v3764 + int32(32)
	goto L598
L609:
	;
	v3827 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+4))
	if v3827 != int32(1) {
		goto L607
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	v4188 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+20))
	v4189 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+28))
	v4190 = *(*int32)(unsafe.Add(mBase, uint32(v3823)+76))
	v4195 = F_recurse_set_operations(m, v3767, v47, int32(0), v4188, v4189, v4190, v3764+int32(8), v3764+int32(28))
	mBase = m.M
	v4196 = m.ExcPending
	if v4196 != 0 {
		goto L1
	} else {
		goto L675
	}
L612:
	;
	v3830 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+12))
	v3832 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+20))
	v3833 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+28))
	v3834 = *(*int32)(unsafe.Add(mBase, uint32(v3823)+76))
	v3839 = F_recurse_set_operations(m, v3830, v47, int32(0), v3832, v3833, v3834, v3764+int32(28), v3764+int32(27))
	mBase = m.M
	v3840 = m.ExcPending
	if v3840 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	v3841 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+84))
	if v3841 == int32(1) {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v3844 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3764)+27)))
	v3845 = *(*int32)(unsafe.Add(mBase, uint32(v3764)+28))
	v3846 = int32(0)
	F_build_setop_child_paths(m, v47, v3839, v3844, v3845, v3846, v3846)
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L617
	}
L615:
	;
	goto L616
L616:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+364)) = v3850
	v3852 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+16))
	v3854 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+20))
	v3855 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+28))
	v3860 = F_recurse_set_operations(m, v3852, v47, int32(0), v3854, v3855, v3834, v3764+int32(20), v3764+int32(19))
	mBase = m.M
	v3861 = m.ExcPending
	if v3861 != 0 {
		goto L1
	} else {
		goto L618
	}
L617:
	;
	goto L616
L618:
	;
	v3862 = *(*int32)(unsafe.Add(mBase, uint32(v3764)+20))
	v3863 = *(*int32)(unsafe.Add(mBase, uint32(v3860)+84))
	if v3863 == int32(1) {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v3866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3764)+19)))
	v3867 = int32(0)
	F_build_setop_child_paths(m, v47, v3860, v3866, v3862, v3867, v3867)
	mBase = m.M
	v3870 = m.ExcPending
	if v3870 != 0 {
		goto L1
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	v3871 = *(*int32)(unsafe.Add(mBase, uint32(v3860)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+364)) = int32(0)
	v3874 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+28))
	v3875 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3764)+12)) = v3862
	v3877 = *(*int32)(unsafe.Add(mBase, uint32(v3764)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3764)+4)) = v3877
	*(*int32)(unsafe.Add(mBase, uint32(v3764))) = v3862
	v3882 = F_list_make2_impl(m, v3764+int32(4), v3764)
	mBase = m.M
	v3883 = m.ExcPending
	if v3883 != 0 {
		goto L1
	} else {
		goto L623
	}
L622:
	;
	goto L621
L623:
	;
	v3884 = F_generate_append_tlist(m, v3875, v3874, v3882, v3834)
	mBase = m.M
	v3885 = m.ExcPending
	if v3885 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3764)+8)) = v3884
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3839)+8))
	v3889 = *(*int32)(unsafe.Add(mBase, uint32(v3860)+8))
	v3890 = F_bms_union(m, v3888, v3889)
	mBase = m.M
	v3891 = m.ExcPending
	if v3891 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	v3892 = F_fetch_upper_rel(m, v47, int32(0), v3890)
	mBase = m.M
	v3893 = m.ExcPending
	if v3893 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	v3894 = F_make_pathtarget_from_tlist(m, v3884)
	mBase = m.M
	v3895 = m.ExcPending
	if v3895 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	v3896 = F_set_pathtarget_cost_width(m, v47, v3894)
	mBase = m.M
	v3897 = m.ExcPending
	if v3897 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3892)+40)) = v3896
	v3899 = int32(0)
	v3900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3767)+8)))
	if v3900 == v3899 {
		goto L629
	} else {
		goto L630
	}
L629:
	;
	v3903 = *(*int32)(unsafe.Add(mBase, uint32(v3767)+32))
	v3904 = F_copyObjectImpl(m, v3903)
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
		goto L1
	} else {
		goto L632
	}
L630:
	;
	v4066 = v3896
	v4071 = float64(0)
	v4072 = v3899
	goto L631
L631:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, uint32(v47)+360))
	v4108 = F_palloc0(m, int32(96))
	mBase = m.M
	v4109 = m.ExcPending
	if v4109 != 0 {
		goto L1
	} else {
		goto L659
	}
L632:
	;
	if v3904 != 0 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3906 = *(*int32)(unsafe.Add(mBase, uint32(v3904)+12))
	v3908 = v3906
	goto L635
L634:
	;
	v3908 = int32(0)
	goto L635
L635:
	;
	if v3884 == int32(0) {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v4019 = int32(0)
	if v3904 == v4019 {
		goto L646
	} else {
		goto L647
	}
L637:
	;
	v3911 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+4))
	if v3911 <= int32(0) {
		goto L636
	} else {
		goto L638
	}
L638:
	;
	v3915 = int32(0)
	v3916 = v3908
	goto L639
L639:
	;
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v3904)+12))
	v3957 = *(*int32)(unsafe.Add(mBase, uint32(v3904)+4))
	v3958 = *(*int32)(unsafe.Add(mBase, uint32(v3916)))
	v3959 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+12))
	v3960 = int32(2)
	v3963 = *(*int32)(unsafe.Add(mBase, uint32(v3959+v3915<<(uint(v3960)%32))))
	v3964 = *(*int32)(unsafe.Add(mBase, uint32(v3963)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v3958)+4)) = v3964
	v3967 = v3916 + int32(4)
	if base.Ui32(v3967) < base.Ui32(v3956+v3957<<(uint(v3960)%32)) {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	goto L636
L641:
	;
	v3973 = v3967
	goto L643
L642:
	;
	v3973 = int32(0)
	goto L643
L643:
	;
	v3975 = v3915 + int32(1)
	v3976 = *(*int32)(unsafe.Add(mBase, uint32(v3884)+4))
	if v3975 < v3976 {
		v3915 = v3975
		v3916 = v3973
		goto L639
	} else {
		goto L644
	}
L644:
	;
	goto L640
L645:
	;
	if v4056 == int32(0) {
		goto L606
	} else {
		goto L658
	}
L646:
	;
	v4056 = int32(1)
	goto L645
L647:
	;
	goto L648
L648:
	;
	v4026 = *(*int32)(unsafe.Add(mBase, uint32(v3904)+4))
	if v4026 <= int32(0) {
		v4050 = int32(1)
		goto L649
	} else {
		goto L650
	}
L649:
	;
	v4056 = v4050
	goto L645
L650:
	;
	v4029 = int32(0)
	if v4029 < v4026 {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	v4032 = v4026
	goto L653
L652:
	;
	v4032 = v4029
	goto L653
L653:
	;
	v4033 = *(*int32)(unsafe.Add(mBase, uint32(v3904)+12))
	v4037 = v4019
	goto L654
L654:
	;
	v4041 = *(*int32)(unsafe.Add(mBase, uint32(v4033+v4037<<(uint(int32(2))%32))))
	v4042 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4041)+18)))
	if v4042 != int32(1) {
		v4050 = v4042
		goto L649
	} else {
		goto L656
	}
L655:
	;
	v4050 = v4042
	goto L649
L656:
	;
	v4046 = v4037 + int32(1)
	if v4046 != v4032 {
		v4037 = v4046
		goto L654
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	v4059 = *(*float64)(unsafe.Add(mBase, uint32(v3871)+32))
	v4062 = *(*float64)(unsafe.Add(mBase, uint32(v3850)+32))
	v4064 = *(*int32)(unsafe.Add(mBase, uint32(v3892)+40))
	v4066 = v4064
	v4071 = base.F64_add(base.F64_mul(v4059, float64(10)), v4062)
	v4072 = v3904
	goto L631
L659:
	;
	v4110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4108)+20)) = uint8(v4110)
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+16)) = v4110
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+12)) = v4066
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+8)) = v3892
	*(*int64)(unsafe.Add(mBase, uint32(v4108))) = int64(1460288880956)
	v4119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3892)+26)))
	if v4119 != int32(1) {
		v4127 = v4110
		goto L660
	} else {
		goto L661
	}
L660:
	;
	v4129 = v4127 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4108)+21)) = uint8(v4129)
	v4131 = *(*int32)(unsafe.Add(mBase, uint32(v3850)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v4108)+88)) = v4071
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+84)) = v4106
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+80)) = v4072
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+76)) = v3871
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+72)) = v3850
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+64)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+24)) = v4131
	v4141 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	v4142 = *(*float64)(unsafe.Add(mBase, uint32(v3850)+56))
	v4143 = *(*float64)(unsafe.Add(mBase, uint32(v3871)+56))
	v4144 = *(*int32)(unsafe.Add(mBase, uint32(v4108)+8))
	v4145 = *(*int64)(unsafe.Add(mBase, uint32(v4144)+32))
	v4146 = *(*float64)(unsafe.Add(mBase, uint32(v3850)+32))
	v4147 = *(*float64)(unsafe.Add(mBase, uint32(v3871)+32))
	v4148 = *(*float64)(unsafe.Add(mBase, uint32(v3850)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v4108)+48)) = v4148
	v4152 = base.F64_add(v4146, base.F64_mul(v4147, float64(10)))
	*(*float64)(unsafe.Add(mBase, uint32(v4108)+32)) = v4152
	if v4131 != 0 {
		goto L663
	} else {
		goto L664
	}
L661:
	;
	v4123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3850)+21)))
	if v4123 != int32(1) {
		v4127 = int32(0)
		goto L660
	} else {
		goto L662
	}
L662:
	;
	v4126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3871)+21)))
	v4127 = v4126
	goto L660
L663:
	;
	v4157 = int64(-1)
	goto L665
L664:
	;
	v4157 = int64(-262145)
	goto L665
L665:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4108)+40)) = base.B2i32(v4145|v4157 != int64(-1))
	*(*float64)(unsafe.Add(mBase, uint32(v4108)+56)) = base.F64_add(base.F64_mul(v4141, v4152), base.F64_add(v4142, base.F64_mul(v4143, float64(10))))
	v4168 = *(*int32)(unsafe.Add(mBase, uint32(v4108)+12))
	v4169 = *(*int32)(unsafe.Add(mBase, uint32(v3850)+12))
	v4170 = *(*int32)(unsafe.Add(mBase, uint32(v4169)+32))
	v4171 = *(*int32)(unsafe.Add(mBase, uint32(v3871)+12))
	v4172 = *(*int32)(unsafe.Add(mBase, uint32(v4171)+32))
	if v4172 < v4170 {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v4174 = v4170
	goto L668
L667:
	;
	v4174 = v4172
	goto L668
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4168)+32)) = v4174
	F_add_path(m, v3892, v4108)
	mBase = m.M
	v4177 = m.ExcPending
	if v4177 != 0 {
		goto L1
	} else {
		goto L669
	}
L669:
	;
	v4179 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v4179 != 0 {
		goto L670
	} else {
		goto L671
	}
L670:
	;
	v4180 = int32(0)
	m.T0[v4179].(func(*base.Module, int32, int32, int32, int32, int32))(m, v47, v4180, v4180, v3892, v4180)
	mBase = m.M
	v4184 = m.ExcPending
	if v4184 != 0 {
		goto L1
	} else {
		goto L673
	}
L671:
	;
	goto L672
L672:
	;
	F_set_cheapest(m, v3892)
	mBase = m.M
	v4186 = m.ExcPending
	if v4186 != 0 {
		goto L1
	} else {
		goto L674
	}
L673:
	;
	goto L672
L674:
	;
	v4200 = v3892
	v4203 = v3884
	goto L608
L675:
	;
	v4197 = *(*int32)(unsafe.Add(mBase, uint32(v3764)+8))
	v4200 = v4195
	v4203 = v4197
	goto L608
L676:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_12), int32(0))
	mBase = m.M
	v4250 = m.ExcPending
	if v4250 != 0 {
		goto L1
	} else {
		goto L677
	}
L677:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_13), int32(385), int32(_a_F_subquery_planner_14))
	mBase = m.M
	v4255 = m.ExcPending
	if v4255 != 0 {
		goto L1
	} else {
		goto L678
	}
L678:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L679:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4262 = m.ExcPending
	if v4262 != 0 {
		goto L1
	} else {
		goto L680
	}
L680:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_15), int32(0))
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	v4269 = F_errdetail(m, int32(_a_F_subquery_planner_16), int32(0))
	mBase = m.M
	v4270 = m.ExcPending
	if v4270 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_13), int32(449), int32(_a_F_subquery_planner_14))
	mBase = m.M
	v4275 = m.ExcPending
	if v4275 != 0 {
		goto L1
	} else {
		goto L683
	}
L683:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L684:
	;
	v4279 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+76))
	if v4279 != 0 {
		goto L685
	} else {
		goto L686
	}
L685:
	;
	v4280 = *(*int32)(unsafe.Add(mBase, uint32(v4279)+12))
	v4282 = v4280
	goto L687
L686:
	;
	v4282 = int32(0)
	goto L687
L687:
	;
	if v4277 == int32(0) {
		v4378 = v4282
		goto L688
	} else {
		goto L689
	}
L688:
	;
	if v4378 != 0 {
		goto L593
	} else {
		goto L701
	}
L689:
	;
	v4285 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+4))
	if v4285 <= int32(0) {
		v4378 = v4282
		goto L688
	} else {
		goto L690
	}
L690:
	;
	v4289 = v4285
	v4290 = int32(0)
	v4305 = v4282
	goto L691
L691:
	;
	v4330 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+12))
	v4334 = *(*int32)(unsafe.Add(mBase, uint32(v4330+v4290<<(uint(int32(2))%32))))
	v4335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4334)+26)))
	if v4335 == int32(0) {
		goto L693
	} else {
		goto L694
	}
L692:
	;
	v4378 = v4358
	goto L688
L693:
	;
	v4338 = *(*int32)(unsafe.Add(mBase, uint32(v4305)))
	v4339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4338)+26)))
	if v4339 == int32(1) {
		goto L594
	} else {
		goto L696
	}
L694:
	;
	v4355 = v4289
	v4358 = v4305
	goto L695
L695:
	;
	v4360 = v4290 + int32(1)
	if v4360 < v4355 {
		v4289 = v4355
		v4290 = v4360
		v4305 = v4358
		goto L691
	} else {
		goto L700
	}
L696:
	;
	v4342 = *(*int32)(unsafe.Add(mBase, uint32(v4279)+12))
	v4343 = *(*int32)(unsafe.Add(mBase, uint32(v4279)+4))
	v4344 = *(*int32)(unsafe.Add(mBase, uint32(v4338)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v4334)+16)) = v4344
	v4347 = v4305 + int32(4)
	if base.Ui32(v4347) < base.Ui32(v4342+v4343<<(uint(int32(2))%32)) {
		goto L697
	} else {
		goto L698
	}
L697:
	;
	v4353 = v4347
	goto L699
L698:
	;
	v4353 = int32(0)
	goto L699
L699:
	;
	v4354 = *(*int32)(unsafe.Add(mBase, uint32(v4277)+4))
	v4355 = v4354
	v4358 = v4353
	goto L695
L700:
	;
	goto L692
L701:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+284)) = v4277
	v4404 = *(*int32)(unsafe.Add(mBase, uint32(v4200)+60))
	v4405 = *(*int32)(unsafe.Add(mBase, uint32(v4404)+12))
	v4406 = *(*int32)(unsafe.Add(mBase, uint32(v4405)+4))
	v4407 = F_is_parallel_safe(m, v47, v4406)
	mBase = m.M
	v4408 = m.ExcPending
	if v4408 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	v4409 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+204)) = v4409
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+200)) = v4409
	v4413 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+140))
	if v4413 != 0 {
		goto L592
	} else {
		goto L703
	}
L703:
	;
	v4414 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+124))
	v4415 = *(*int32)(unsafe.Add(mBase, uint32(v47)+284))
	v4416 = F_make_pathkeys_for_sortclauses(m, v47, v4414, v4415)
	mBase = m.M
	v4417 = m.ExcPending
	if v4417 != 0 {
		goto L1
	} else {
		goto L704
	}
L704:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+196)) = v4416
	v14910 = v4200
	v14913 = v3756
	v14915 = v47
	v14918 = v3642
	v14920 = v3644
	v14927 = v4405
	v14930 = v44
	v14932 = v4407
	v14941 = v3756
	v14945 = v3757
	v14946 = v3758
	goto L584
L705:
	;
	v4420 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v4422 = F_palloc0(m, int32(40))
	mBase = m.M
	v4423 = m.ExcPending
	if v4423 != 0 {
		goto L1
	} else {
		goto L708
	}
L706:
	;
	goto L707
L707:
	;
	v4656 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+100))
	if v4656 == int32(0) {
		v8227 = v47
		v8230 = v3642
		v8232 = v3644
		v8236 = v9
		v8239 = l7
		v8242 = v44
		v8248 = v9
		v8253 = v3756
		v8257 = v3757
		v8258 = v3758
		goto L585
	} else {
		goto L751
	}
L708:
	;
	v4424 = *(*int32)(unsafe.Add(mBase, uint32(v4420)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+276)) = v4424
	v4426 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4422)+16)) = uint8(v4426)
	*(*int32)(unsafe.Add(mBase, uint32(v4422)+28)) = v4426
	*(*int64)(unsafe.Add(mBase, uint32(v4422)+20)) = int64(0)
	v4432 = int32(4)
	v4433 = *(*int32)(unsafe.Add(mBase, uint32(v4420)+100))
	if v4433 == v4426 {
		v4511 = v4432
		goto L709
	} else {
		goto L710
	}
L709:
	;
	v4551 = F_palloc(m, v4511)
	mBase = m.M
	v4552 = m.ExcPending
	if v4552 != 0 {
		goto L1
	} else {
		goto L726
	}
L710:
	;
	v4436 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+4))
	if v4436 <= int32(0) {
		v4511 = v4432
		goto L709
	} else {
		goto L711
	}
L711:
	;
	v4448 = v3637
	v4455 = v9
	goto L712
L712:
	;
	v4480 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+12))
	v4484 = *(*int32)(unsafe.Add(mBase, uint32(v4480+v4455<<(uint(int32(2))%32))))
	v4485 = *(*int32)(unsafe.Add(mBase, uint32(v4484)+4))
	v4486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4484)+18)))
	if v4486 == int32(0) {
		goto L714
	} else {
		goto L715
	}
L713:
	;
	v4511 = v4501<<(uint(int32(2))%32) + int32(4)
	goto L709
L714:
	;
	v4489 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+24))
	v4490 = F_bms_add_member(m, v4489, v4485)
	mBase = m.M
	v4491 = m.ExcPending
	if v4491 != 0 {
		goto L1
	} else {
		goto L717
	}
L715:
	;
	goto L716
L716:
	;
	v4493 = *(*int32)(unsafe.Add(mBase, uint32(v4484)+12))
	if v4493 == int32(0) {
		goto L718
	} else {
		goto L719
	}
L717:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4422)+24)) = v4490
	goto L716
L718:
	;
	v4496 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+20))
	v4497 = F_bms_add_member(m, v4496, v4485)
	mBase = m.M
	v4498 = m.ExcPending
	if v4498 != 0 {
		goto L1
	} else {
		goto L721
	}
L719:
	;
	goto L720
L720:
	;
	if base.Ui32(v4448) < base.Ui32(v4485) {
		goto L722
	} else {
		goto L723
	}
L721:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4422)+20)) = v4497
	goto L720
L722:
	;
	v4501 = v4485
	goto L724
L723:
	;
	v4501 = v4448
	goto L724
L724:
	;
	v4503 = v4455 + int32(1)
	v4504 = *(*int32)(unsafe.Add(mBase, uint32(v4433)+4))
	if v4503 < v4504 {
		v4448 = v4501
		v4455 = v4503
		goto L712
	} else {
		goto L725
	}
L725:
	;
	goto L713
L726:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4422)+32)) = v4551
	v4554 = *(*int32)(unsafe.Add(mBase, uint32(v4420)+108))
	v4555 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+20))
	if v4555 != 0 {
		goto L727
	} else {
		goto L728
	}
L727:
	;
	if v4554 == int32(0) {
		v7872 = v47
		v7875 = v3642
		v7877 = v3644
		v7879 = v4420
		v7881 = v4422
		v7884 = l7
		v7887 = v44
		v7893 = v9
		v7898 = v3756
		v7902 = v3757
		v7903 = v3758
		goto L586
	} else {
		goto L730
	}
L728:
	;
	goto L729
L729:
	;
	if v4554 != 0 {
		v4725 = v4554
		goto L590
	} else {
		goto L750
	}
L730:
	;
	v4558 = *(*int32)(unsafe.Add(mBase, uint32(v4554)+4))
	if v4558 <= int32(0) {
		v7872 = v47
		v7875 = v3642
		v7877 = v3644
		v7879 = v4420
		v7881 = v4422
		v7884 = l7
		v7887 = v44
		v7893 = v9
		v7898 = v3756
		v7902 = v3757
		v7903 = v3758
		goto L586
	} else {
		goto L731
	}
L731:
	;
	v4561 = int32(0)
	v4563 = v4561
	v4564 = v4561
	goto L732
L732:
	;
	v4604 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+20))
	v4605 = *(*int32)(unsafe.Add(mBase, uint32(v4554)+12))
	v4609 = *(*int32)(unsafe.Add(mBase, uint32(v4605+v4564<<(uint(int32(2))%32))))
	v4610 = F_bms_overlap_list(m, v4604, v4609)
	mBase = m.M
	v4611 = m.ExcPending
	if v4611 != 0 {
		goto L1
	} else {
		goto L735
	}
L733:
	;
	goto L591
L734:
	;
	v4652 = v4564 + int32(1)
	v4653 = *(*int32)(unsafe.Add(mBase, uint32(v4554)+4))
	if v4652 < v4653 {
		v4563 = v4649
		v4564 = v4652
		goto L732
	} else {
		goto L749
	}
L735:
	;
	if v4610 != 0 {
		goto L736
	} else {
		goto L737
	}
L736:
	;
	v4613 = F_palloc0(m, int32(16))
	mBase = m.M
	v4614 = m.ExcPending
	if v4614 != 0 {
		goto L1
	} else {
		goto L739
	}
L737:
	;
	goto L738
L738:
	;
	v4647 = F_lappend(m, v4563, v4609)
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L1
	} else {
		goto L748
	}
L739:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4613)+4)) = v4609
	*(*int32)(unsafe.Add(mBase, uint32(v4613))) = int32(310)
	v4618 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+28))
	v4619 = F_lappend(m, v4618, v4613)
	mBase = m.M
	v4620 = m.ExcPending
	if v4620 != 0 {
		goto L1
	} else {
		goto L740
	}
L740:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4422)+28)) = v4619
	v4622 = *(*int32)(unsafe.Add(mBase, uint32(v4422)+24))
	v4623 = F_bms_overlap_list(m, v4622, v4609)
	mBase = m.M
	v4624 = m.ExcPending
	if v4624 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	if v4623 == int32(0) {
		v4649 = v4563
		goto L734
	} else {
		goto L742
	}
L742:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4630 = m.ExcPending
	if v4630 != 0 {
		goto L1
	} else {
		goto L743
	}
L743:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4633 = m.ExcPending
	if v4633 != 0 {
		goto L1
	} else {
		goto L744
	}
L744:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_17), int32(0))
	mBase = m.M
	v4637 = m.ExcPending
	if v4637 != 0 {
		goto L1
	} else {
		goto L745
	}
L745:
	;
	v4640 = F_errdetail(m, int32(_a_F_subquery_planner_18), int32(0))
	mBase = m.M
	v4641 = m.ExcPending
	if v4641 != 0 {
		goto L1
	} else {
		goto L746
	}
L746:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(2560), int32(_a_F_subquery_planner_19))
	mBase = m.M
	v4646 = m.ExcPending
	if v4646 != 0 {
		goto L1
	} else {
		goto L747
	}
L747:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L748:
	;
	v4649 = v4647
	goto L734
L749:
	;
	goto L733
L750:
	;
	v4822 = int32(0)
	goto L589
L751:
	;
	v4660 = F_preprocess_groupclause(m, v47, int32(0))
	mBase = m.M
	v4661 = m.ExcPending
	if v4661 != 0 {
		goto L1
	} else {
		goto L752
	}
L752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+276)) = v4660
	v8227 = v47
	v8230 = v3642
	v8232 = v3644
	v8236 = v9
	v8239 = l7
	v8242 = v44
	v8248 = v9
	v8253 = v3756
	v8257 = v3757
	v8258 = v3758
	goto L585
L753:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_20), int32(0))
	mBase = m.M
	v4670 = m.ExcPending
	if v4670 != 0 {
		goto L1
	} else {
		goto L754
	}
L754:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(_a_F_subquery_planner_21), int32(_a_F_subquery_planner_22))
	mBase = m.M
	v4675 = m.ExcPending
	if v4675 != 0 {
		goto L1
	} else {
		goto L755
	}
L755:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L756:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_20), int32(0))
	mBase = m.M
	v4683 = m.ExcPending
	if v4683 != 0 {
		goto L1
	} else {
		goto L757
	}
L757:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(_a_F_subquery_planner_23), int32(_a_F_subquery_planner_22))
	mBase = m.M
	v4688 = m.ExcPending
	if v4688 != 0 {
		goto L1
	} else {
		goto L758
	}
L758:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L759:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4695 = m.ExcPending
	if v4695 != 0 {
		goto L1
	} else {
		goto L760
	}
L760:
	;
	v4696 = *(*int32)(unsafe.Add(mBase, uint32(v3644)+140))
	v4697 = *(*int32)(unsafe.Add(mBase, uint32(v4696)+12))
	v4698 = *(*int32)(unsafe.Add(mBase, uint32(v4697)))
	v4699 = *(*int32)(unsafe.Add(mBase, uint32(v4698)+8))
	v4701 = v4699 - int32(1)
	if base.Ui32(v4701) <= base.Ui32(int32(3)) {
		goto L762
	} else {
		goto L763
	}
L761:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+96)) = v4708
	F_errmsg(m, int32(_a_F_subquery_planner_24), v3642+int32(96))
	mBase = m.M
	v4714 = m.ExcPending
	if v4714 != 0 {
		goto L1
	} else {
		goto L765
	}
L762:
	;
	v4706 = *(*int32)(unsafe.Add(mBase, uint32(v4701<<(uint(int32(2))%32))+uint32(_c_F_subquery_planner[3])))
	v4708 = v4706
	goto L764
L763:
	;
	v4708 = int32(_a_F_subquery_planner_25)
	goto L764
L764:
	;
	goto L761
L765:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(1816), int32(_a_F_subquery_planner_26))
	mBase = m.M
	v4719 = m.ExcPending
	if v4719 != 0 {
		goto L1
	} else {
		goto L766
	}
L766:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L767:
	;
	v4725 = v4649
	goto L590
L768:
	;
	v4766 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+4))
	v4768 = v4766 + int32(1)
	v4773 = v4763
	v4791 = v3637
	goto L769
L769:
	;
	v4813 = *(*int32)(unsafe.Add(mBase, uint32(v4773)))
	if v4813 != 0 {
		goto L588
	} else {
		goto L771
	}
L770:
	;
	v4822 = v4725
	goto L589
L771:
	;
	v4817 = v4773 + int32(4)
	if base.Ui32(v4817) < base.Ui32(v4763+v4766<<(uint(int32(2))%32)) {
		v4773 = v4817
		v4791 = v4791 + int32(1)
		goto L769
	} else {
		goto L772
	}
L772:
	;
	goto L770
L773:
	;
	v6799 = v4865
	goto L587
L774:
	;
	v4872 = F_palloc0(m, v4869)
	mBase = m.M
	v4873 = m.ExcPending
	if v4873 != 0 {
		goto L1
	} else {
		goto L775
	}
L775:
	;
	v4874 = F_palloc0(m, v4869)
	mBase = m.M
	v4875 = m.ExcPending
	if v4875 != 0 {
		goto L1
	} else {
		goto L776
	}
L776:
	;
	v4878 = F_palloc(m, v4768<<(uint(int32(1))%32))
	mBase = m.M
	v4879 = m.ExcPending
	if v4879 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v4880 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+12))
	v4883 = (v4773 - v4880) >> (uint(int32(2)) % 32)
	v4884 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+4))
	if v4883 < v4884 {
		goto L778
	} else {
		goto L779
	}
L778:
	;
	v4893 = v3637
	v4900 = v4883
	v4904 = v9
	v4907 = v4867
	goto L781
L779:
	;
	v5528 = v4867
	goto L780
L780:
	;
	v5548 = int32(0)
	v5550 = v5528 - int32(1)
	v5552 = F_palloc(m, int32(32))
	mBase = m.M
	v5553 = m.ExcPending
	if v5553 != 0 {
		goto L1
	} else {
		goto L858
	}
L781:
	;
	v4927 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+12))
	v4931 = *(*int32)(unsafe.Add(mBase, uint32(v4927+v4900<<(uint(int32(2))%32))))
	if v4931 != 0 {
		goto L787
	} else {
		goto L788
	}
L782:
	;
	v5528 = v5483
	goto L780
L783:
	;
	v5504 = v4900 + int32(1)
	v5505 = *(*int32)(unsafe.Add(mBase, uint32(v4725)+4))
	if v5504 < v5505 {
		v4893 = v5469
		v4900 = v5504
		v4904 = v5480
		v4907 = v5483
		goto L781
	} else {
		goto L856
	}
L784:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+92)) = v4931
	*(*int32)(unsafe.Add(mBase, uint32(v3642)+320)) = v4931
	v5236 = v4907 << (uint(int32(2)) % 32)
	v5241 = F_list_make1_impl(m, int32(1), v3642+int32(92))
	mBase = m.M
	v5242 = m.ExcPending
	if v5242 != 0 {
		goto L1
	} else {
		goto L829
	}
L785:
	;
	v5188 = int32(0)
	if v5188 < v4904 {
		goto L826
	} else {
		goto L827
	}
L786:
	;
	if v4907 <= v4893 {
		v5199 = v4893
		v5201 = v5045
		v5210 = v4904
		goto L784
	} else {
		goto L805
	}
L787:
	;
	v4932 = int32(0)
	v4934 = *(*int32)(unsafe.Add(mBase, uint32(v4931)+4))
	if v4932 < v4934 {
		goto L790
	} else {
		goto L791
	}
L788:
	;
	goto L789
L789:
	;
	v5035 = int32(0)
	if v4904 != 0 {
		goto L785
	} else {
		goto L804
	}
L790:
	;
	v4938 = v4932
	v4946 = v4932
	goto L793
L791:
	;
	v4998 = v4932
	v4999 = v4934
	goto L792
L792:
	;
	if v4999 == v4904 {
		v5045 = v4998
		goto L786
	} else {
		goto L797
	}
L793:
	;
	v4978 = *(*int32)(unsafe.Add(mBase, uint32(v4931)+12))
	v4982 = *(*int32)(unsafe.Add(mBase, uint32(v4978+v4938<<(uint(int32(2))%32))))
	v4983 = F_bms_add_member(m, v4946, v4982)
	mBase = m.M
	v4984 = m.ExcPending
	if v4984 != 0 {
		goto L1
	} else {
		goto L795
	}
L794:
	;
	v4998 = v4983
	v4999 = v4985
	goto L792
L795:
	;
	v4985 = *(*int32)(unsafe.Add(mBase, uint32(v4931)+4))
	v4987 = v4938 + int32(1)
	if v4987 < v4985 {
		v4938 = v4987
		v4946 = v4983
		goto L793
	} else {
		goto L796
	}
L796:
	;
	goto L794
L797:
	;
	if v4904 < v4999 {
		goto L798
	} else {
		goto L799
	}
L798:
	;
	v5032 = v4907
	goto L800
L799:
	;
	v5032 = v4893
	goto L800
L800:
	;
	if v4999 < v4904 {
		goto L801
	} else {
		goto L802
	}
L801:
	;
	v5034 = v4904
	goto L803
L802:
	;
	v5034 = v4999
	goto L803
L803:
	;
	v5199 = v5032
	v5201 = v4998
	v5210 = v5034
	goto L784
L804:
	;
	v5045 = v5035
	goto L786
L805:
	;
	v5079 = v4893
	goto L806
L806:
	;
	v5120 = v5079 << (uint(int32(2)) % 32)
	v5122 = *(*int32)(unsafe.Add(mBase, uint32(v4872+v5120)))
	v5123 = int32(0)
	if base.B2i32(v5122 == v5123)|base.B2i32(v5045 == v5123) != 0 {
		v5169 = base.B2i32(v5122|v5045 == v5123)
		goto L809
	} else {
		goto L810
	}
L807:
	;
	if v5079 <= int32(0) {
		v5199 = v4893
		v5201 = v5045
		v5210 = v4904
		goto L784
	} else {
		goto L823
	}
L808:
	;
	if v5169 == int32(0) {
		goto L819
	} else {
		goto L820
	}
L809:
	;
	goto L808
L810:
	;
	v5137 = *(*int32)(unsafe.Add(mBase, uint32(v5122)+4))
	v5138 = *(*int32)(unsafe.Add(mBase, uint32(v5045)+4))
	if v5137 != v5138 {
		v5169 = int32(0)
		goto L809
	} else {
		goto L811
	}
L811:
	;
	v5140 = int32(1)
	if v5137 <= v5140 {
		goto L812
	} else {
		goto L813
	}
L812:
	;
	v5143 = v5140
	goto L814
L813:
	;
	v5143 = v5137
	goto L814
L814:
	;
	v5144 = int32(8)
	v5149 = int32(0)
	goto L815
L815:
	;
	v5157 = v5149 << (uint(int32(2)) % 32)
	v5159 = *(*int32)(unsafe.Add(mBase, uint32(v5122+v5144+v5157)))
	v5161 = *(*int32)(unsafe.Add(mBase, uint32(v5045+v5144+v5157)))
	v5162 = base.B2i32(v5159 == v5161)
	if v5159 != v5161 {
		v5169 = v5162
		goto L809
	} else {
		goto L817
	}
L816:
	;
	v5169 = v5162
	goto L809
L817:
	;
	v5165 = v5149 + int32(1)
	if v5165 != v5143 {
		v5149 = v5165
		goto L815
	} else {
		goto L818
	}
L818:
	;
	goto L816
L819:
	;
	v5177 = v5079 + int32(1)
	if v5177 != v4907 {
		v5079 = v5177
		goto L806
	} else {
		goto L822
	}
L820:
	;
	goto L821
L821:
	;
	goto L807
L822:
	;
	v5199 = v4893
	v5201 = v5045
	v5210 = v4904
	goto L784
L823:
	;
	v5181 = v5120 + v4870
	v5182 = *(*int32)(unsafe.Add(mBase, uint32(v5181)))
	v5183 = F_lappend(m, v5182, v4931)
	mBase = m.M
	v5184 = m.ExcPending
	if v5184 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5181))) = v5183
	F_bms_free(m, v5045)
	mBase = m.M
	v5187 = m.ExcPending
	if v5187 != 0 {
		goto L1
	} else {
		goto L825
	}
L825:
	;
	v5469 = v4893
	v5480 = v4904
	v5483 = v4907
	goto L783
L826:
	;
	v5191 = v4904
	goto L828
L827:
	;
	v5191 = v5188
	goto L828
L828:
	;
	v5199 = v4893
	v5201 = v5035
	v5210 = v5191
	goto L784
L829:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4870+v5236))) = v5241
	*(*int32)(unsafe.Add(mBase, uint32(v5236+v4872))) = v5201
	v5246 = int32(0)
	v5248 = v5199 - int32(1)
	if v5248 <= v5246 {
		goto L831
	} else {
		goto L832
	}
L830:
	;
	v5469 = v5199
	v5480 = v5210
	v5483 = v4907 + int32(1)
	goto L783
L831:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5236+v4874))) = int32(0)
	goto L830
L832:
	;
	v5251 = v5246
	v5252 = v5248
	goto L833
L833:
	;
	v5295 = *(*int32)(unsafe.Add(mBase, uint32(v4872+v5252<<(uint(int32(2))%32))))
	v5296 = int32(0)
	if v5295 == v5296 {
		goto L836
	} else {
		goto L837
	}
L834:
	;
	if v5356 <= int32(0) {
		goto L831
	} else {
		goto L853
	}
L835:
	;
	if v5349 != 0 {
		goto L849
	} else {
		goto L850
	}
L836:
	;
	v5349 = int32(1)
	goto L835
L837:
	;
	goto L838
L838:
	;
	if v5201 == int32(0) {
		v5342 = v5296
		goto L839
	} else {
		goto L840
	}
L839:
	;
	v5349 = v5342
	goto L835
L840:
	;
	v5305 = *(*int32)(unsafe.Add(mBase, uint32(v5295)+4))
	v5306 = *(*int32)(unsafe.Add(mBase, uint32(v5201)+4))
	if v5306 < v5305 {
		v5342 = v5296
		goto L839
	} else {
		goto L841
	}
L841:
	;
	v5308 = int32(1)
	if v5305 <= v5308 {
		goto L842
	} else {
		goto L843
	}
L842:
	;
	v5311 = v5308
	goto L844
L843:
	;
	v5311 = v5305
	goto L844
L844:
	;
	v5312 = int32(8)
	v5317 = int32(0)
	goto L845
L845:
	;
	v5324 = v5317 << (uint(int32(2)) % 32)
	v5326 = *(*int32)(unsafe.Add(mBase, uint32(v5295+v5312+v5324)))
	v5328 = *(*int32)(unsafe.Add(mBase, uint32(v5201+v5312+v5324)))
	v5331 = v5326 & (v5328 ^ int32(-1))
	v5333 = base.B2i32(v5331 == int32(0))
	if v5331 != 0 {
		v5342 = v5333
		goto L839
	} else {
		goto L847
	}
L846:
	;
	v5342 = v5333
	goto L839
L847:
	;
	v5335 = v5317 + int32(1)
	if v5335 != v5311 {
		v5317 = v5335
		goto L845
	} else {
		goto L848
	}
L848:
	;
	goto L846
L849:
	;
	v5350 = int32(1)
	v5351 = v5251 + v5350
	*(*uint16)(unsafe.Add(mBase, uint32(v4878+v5351<<(uint(v5350)%32)))) = uint16(v5252)
	v5356 = v5351
	goto L851
L850:
	;
	v5356 = v5251
	goto L851
L851:
	;
	v5357 = int32(1)
	if v5357 < v5252 {
		v5251 = v5356
		v5252 = v5252 - v5357
		goto L833
	} else {
		goto L852
	}
L852:
	;
	goto L834
L853:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v4878))) = uint16(v5356)
	v5368 = v5356<<(uint(int32(1))%32) + int32(2)
	v5369 = F_palloc(m, v5368)
	mBase = m.M
	v5370 = m.ExcPending
	if v5370 != 0 {
		goto L1
	} else {
		goto L854
	}
L854:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5236+v4874))) = v5369
	if v5368 == int32(0) {
		goto L830
	} else {
		goto L855
	}
L855:
	;
	base.MemoryCopy(m, v5369, v4878, v5368)
	goto L830
L856:
	;
	goto L782
L857:
	;
	v6223 = F_palloc0(m, v5528<<(uint(int32(2))%32))
	mBase = m.M
	v6224 = m.ExcPending
	if v6224 != 0 {
		goto L1
	} else {
		goto L927
	}
L858:
	;
	if base.B2i32(base.Ui32(int32(_a_F_subquery_planner_27)) < base.Ui32(v5550))|base.B2i32(base.Ui32(int32(_a_F_subquery_planner_28)) <= base.Ui32(v5550)) == int32(0) {
		goto L859
	} else {
		goto L860
	}
L859:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+8)) = v4874
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+4)) = v5550
	*(*int32)(unsafe.Add(mBase, uint32(v5552))) = v5550
	v5567 = v5550 << (uint(int32(1)) % 32)
	v5569 = v5567 + int32(2)
	v5570 = F_palloc0(m, v5569)
	mBase = m.M
	v5571 = m.ExcPending
	if v5571 != 0 {
		goto L1
	} else {
		goto L862
	}
L860:
	;
	goto L861
L861:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6211 = m.ExcPending
	if v6211 != 0 {
		goto L1
	} else {
		goto L924
	}
L862:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+16)) = v5570
	v5575 = F_palloc0(m, v5567+int32(2))
	mBase = m.M
	v5576 = m.ExcPending
	if v5576 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+20)) = v5575
	v5578 = F_palloc(m, v5569)
	mBase = m.M
	v5579 = m.ExcPending
	if v5579 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+24)) = v5578
	v5583 = F_palloc(m, v5567+int32(4))
	mBase = m.M
	v5584 = m.ExcPending
	if v5584 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+28)) = v5583
	v5586 = *(*int32)(unsafe.Add(mBase, uint32(v5552)))
	v5587 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+24))
	v5588 = int32(_a_F_subquery_planner_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v5587))) = uint16(v5588)
	if v5586 <= int32(0) {
		goto L866
	} else {
		goto L867
	}
L866:
	;
	goto L857
L867:
	;
	v5592 = v5586
	v5601 = v5587
	v5604 = v5583
	goto L868
L868:
	;
	v5633 = int32(1)
	v5634 = int32(0)
	v5636 = v5592 + v5633
	if int32(3) <= v5636 {
		goto L871
	} else {
		goto L872
	}
L869:
	;
	goto L866
L870:
	;
	v5834 = int32(0)
	if v5834 < v5800 {
		goto L892
	} else {
		goto L893
	}
L871:
	;
	v5639 = int32(2)
	if v5636 <= v5639 {
		goto L874
	} else {
		goto L875
	}
L872:
	;
	v5737 = v5633
	v5743 = v5634
	goto L873
L873:
	;
	v5778 = v5737 << (uint(int32(1)) % 32)
	v5779 = v5601 + v5778
	v5780 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	v5782 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5780+v5778))))
	if v5782 != 0 {
		goto L889
	} else {
		goto L890
	}
L874:
	;
	v5642 = v5639
	goto L876
L875:
	;
	v5642 = v5636
	goto L876
L876:
	;
	v5643 = int32(1)
	v5644 = v5642 - v5643
	v5650 = int32(0)
	v5651 = v5633
	v5657 = v5634
	goto L877
L877:
	;
	v5692 = v5651 << (uint(int32(1)) % 32)
	v5693 = v5601 + v5692
	v5694 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	v5696 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5694+v5692))))
	if v5696 == int32(0) {
		goto L880
	} else {
		goto L881
	}
L878:
	;
	if v5644&v5643 == int32(0) {
		v5800 = v5728
		goto L870
	} else {
		goto L888
	}
L879:
	;
	v5710 = int32(1)
	v5711 = v5651 + v5710
	v5713 = v5711 << (uint(v5710) % 32)
	v5714 = v5601 + v5713
	v5715 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	v5717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5715+v5713))))
	if v5717 != 0 {
		goto L884
	} else {
		goto L885
	}
L880:
	;
	v5699 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5693))) = uint16(v5699)
	v5701 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5604+v5657<<(uint(v5701)%32)))) = uint16(v5651)
	v5709 = v5657 + v5701
	goto L879
L881:
	;
	goto L882
L882:
	;
	v5707 = int32(_a_F_subquery_planner_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v5693))) = uint16(v5707)
	v5709 = v5657
	goto L879
L883:
	;
	v5729 = int32(2)
	v5730 = v5651 + v5729
	v5732 = v5650 + v5729
	if v5732 != v5644&int32(-2) {
		v5650 = v5732
		v5651 = v5730
		v5657 = v5728
		goto L877
	} else {
		goto L887
	}
L884:
	;
	v5718 = int32(_a_F_subquery_planner_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v5714))) = uint16(v5718)
	v5728 = v5709
	goto L883
L885:
	;
	goto L886
L886:
	;
	v5720 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5714))) = uint16(v5720)
	v5722 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5604+v5709<<(uint(v5722)%32)))) = uint16(v5711)
	v5728 = v5709 + v5722
	goto L883
L887:
	;
	goto L878
L888:
	;
	v5737 = v5730
	v5743 = v5728
	goto L873
L889:
	;
	v5783 = int32(_a_F_subquery_planner_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v5779))) = uint16(v5783)
	v5800 = v5743
	goto L870
L890:
	;
	goto L891
L891:
	;
	v5785 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v5779))) = uint16(v5785)
	v5787 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5604+v5743<<(uint(v5787)%32)))) = uint16(v5737)
	v5800 = v5743 + v5787
	goto L870
L892:
	;
	v5840 = v5834
	v5844 = v5800
	goto L895
L893:
	;
	goto L894
L894:
	;
	v6054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5601))))
	if v6054 == int32(_a_F_subquery_planner_28) {
		goto L866
	} else {
		goto L908
	}
L895:
	;
	v5878 = int32(1)
	v5881 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5604+v5840<<(uint(v5878)%32)))))
	v5884 = v5601 + v5881<<(uint(v5878)%32)
	v5885 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5884))))
	v5886 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5601))))
	if v5886 <= v5885 {
		v5976 = v5844
		goto L897
	} else {
		goto L898
	}
L896:
	;
	goto L894
L897:
	;
	v6011 = v5840 + int32(1)
	if v6011 < v5976 {
		v5840 = v6011
		v5844 = v5976
		goto L895
	} else {
		goto L907
	}
L898:
	;
	v5888 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+8))
	v5892 = *(*int32)(unsafe.Add(mBase, uint32(v5888+v5881<<(uint(int32(2))%32))))
	if v5892 == int32(0) {
		v5976 = v5844
		goto L897
	} else {
		goto L899
	}
L899:
	;
	v5895 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5892))))
	if v5895 <= int32(0) {
		v5976 = v5844
		goto L897
	} else {
		goto L900
	}
L900:
	;
	v5899 = v5895
	v5905 = v5844
	goto L901
L901:
	;
	v5939 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+20))
	v5940 = int32(1)
	v5943 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5892+v5899<<(uint(v5940)%32)))))
	v5947 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5939+v5943<<(uint(v5940)%32)))))
	v5950 = v5601 + v5947<<(uint(v5940)%32)
	v5951 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5950))))
	if v5951 == int32(_a_F_subquery_planner_28) {
		goto L903
	} else {
		goto L904
	}
L902:
	;
	v5976 = v5964
	goto L897
L903:
	;
	v5954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5884))))
	v5955 = int32(1)
	v5956 = v5954 + v5955
	*(*uint16)(unsafe.Add(mBase, uint32(v5950))) = uint16(v5956)
	*(*uint16)(unsafe.Add(mBase, uint32(v5604+v5905<<(uint(v5955)%32)))) = uint16(v5947)
	v5964 = v5905 + v5955
	goto L905
L904:
	;
	v5964 = v5905
	goto L905
L905:
	;
	v5965 = int32(1)
	if v5965 < v5899 {
		v5899 = v5899 - v5965
		v5905 = v5964
		goto L901
	} else {
		goto L906
	}
L906:
	;
	goto L902
L907:
	;
	goto L896
L908:
	;
	if v5550 != 0 {
		goto L909
	} else {
		goto L910
	}
L909:
	;
	v6058 = int32(1)
	goto L912
L910:
	;
	goto L911
L911:
	;
	v6157 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[4]))
	if v6157 != 0 {
		goto L919
	} else {
		goto L920
	}
L912:
	;
	v6099 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	v6103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6099+v6058<<(uint(int32(1))%32)))))
	if v6103 != 0 {
		goto L914
	} else {
		goto L915
	}
L913:
	;
	goto L911
L914:
	;
	if v6058 != v5550 {
		v6058 = v6058 + int32(1)
		goto L912
	} else {
		goto L918
	}
L915:
	;
	v6104 = F_hk_depth_search(m, v5552, v6058)
	mBase = m.M
	v6105 = m.ExcPending
	if v6105 != 0 {
		goto L1
	} else {
		goto L916
	}
L916:
	;
	if v6104 == int32(0) {
		goto L914
	} else {
		goto L917
	}
L917:
	;
	v6108 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+12)) = v6108 + int32(1)
	goto L914
L918:
	;
	goto L913
L919:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v6159 = m.ExcPending
	if v6159 != 0 {
		goto L1
	} else {
		goto L922
	}
L920:
	;
	goto L921
L921:
	;
	v6160 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+28))
	v6161 = *(*int32)(unsafe.Add(mBase, uint32(v5552)))
	v6162 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+24))
	v6163 = int32(_a_F_subquery_planner_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v6162))) = uint16(v6163)
	if int32(0) < v6161 {
		v5592 = v6161
		v5601 = v6162
		v5604 = v6160
		goto L868
	} else {
		goto L923
	}
L922:
	;
	goto L921
L923:
	;
	goto L869
L924:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_29), int32(0))
	mBase = m.M
	v6215 = m.ExcPending
	if v6215 != 0 {
		goto L1
	} else {
		goto L925
	}
L925:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_30), int32(45), int32(_a_F_subquery_planner_31))
	mBase = m.M
	v6220 = m.ExcPending
	if v6220 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L927:
	;
	if v5550 <= int32(0) {
		goto L929
	} else {
		goto L930
	}
L928:
	;
	if int32(0) < v4791 {
		goto L951
	} else {
		goto L952
	}
L929:
	;
	v6228 = F_palloc0(m, int32(4))
	mBase = m.M
	v6229 = m.ExcPending
	if v6229 != 0 {
		goto L1
	} else {
		goto L932
	}
L930:
	;
	goto L931
L931:
	;
	v6230 = int32(2)
	if v5528 <= v6230 {
		goto L933
	} else {
		goto L934
	}
L932:
	;
	v6380 = v6228
	v6397 = v5548
	goto L928
L933:
	;
	v6233 = v6230
	goto L935
L934:
	;
	v6233 = v5528
	goto L935
L935:
	;
	v6236 = int32(1)
	v6253 = v5548
	goto L936
L936:
	;
	v6277 = v6236 << (uint(int32(1)) % 32)
	v6278 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	v6280 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6277+v6278))))
	v6281 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+20))
	v6283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6281+v6277))))
	v6284 = int32(0)
	if base.B2i32(v6283 <= v6284)|base.B2i32(v6236 <= v6283) == v6284 {
		goto L939
	} else {
		goto L940
	}
L937:
	;
	v6319 = F_palloc0(m, v6307<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v6320 = m.ExcPending
	if v6320 != 0 {
		goto L1
	} else {
		goto L946
	}
L938:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6223+v6236<<(uint(int32(2))%32)))) = v6306
	v6313 = v6236 + int32(1)
	if v6313 != v6233 {
		v6236 = v6313
		v6253 = v6307
		goto L936
	} else {
		goto L945
	}
L939:
	;
	v6293 = *(*int32)(unsafe.Add(mBase, uint32(v6223+v6283<<(uint(int32(2))%32))))
	v6306 = v6293
	v6307 = v6253
	goto L938
L940:
	;
	goto L941
L941:
	;
	v6294 = int32(0)
	if base.B2i32(v6280 <= v6294)|base.B2i32(v6236 <= v6280) == v6294 {
		goto L942
	} else {
		goto L943
	}
L942:
	;
	v6303 = *(*int32)(unsafe.Add(mBase, uint32(v6223+v6280<<(uint(int32(2))%32))))
	v6306 = v6303
	v6307 = v6253
	goto L938
L943:
	;
	goto L944
L944:
	;
	v6305 = v6253 + int32(1)
	v6306 = v6305
	v6307 = v6305
	goto L938
L945:
	;
	goto L937
L946:
	;
	v6331 = int32(1)
	goto L947
L947:
	;
	v6363 = int32(2)
	v6364 = v6331 << (uint(v6363) % 32)
	v6366 = *(*int32)(unsafe.Add(mBase, uint32(v6223+v6364)))
	v6369 = v6319 + v6366<<(uint(v6363)%32)
	v6370 = *(*int32)(unsafe.Add(mBase, uint32(v6369)))
	v6372 = *(*int32)(unsafe.Add(mBase, uint32(v6364+v4870)))
	v6373 = F_list_concat(m, v6370, v6372)
	mBase = m.M
	v6374 = m.ExcPending
	if v6374 != 0 {
		goto L1
	} else {
		goto L949
	}
L948:
	;
	v6380 = v6319
	v6397 = v6307
	goto L928
L949:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6369))) = v6373
	v6377 = v6331 + int32(1)
	if v6377 <= v5550 {
		v6331 = v6377
		goto L947
	} else {
		goto L950
	}
L950:
	;
	goto L948
L951:
	;
	v6423 = *(*int32)(unsafe.Add(mBase, uint32(v6380)+4))
	v6440 = v6423
	v6443 = v4791
	goto L954
L952:
	;
	goto L953
L953:
	;
	v6514 = int32(0)
	if v6514 < v6397 {
		goto L958
	} else {
		goto L959
	}
L954:
	;
	v6466 = F_lcons(m, int32(0), v6440)
	mBase = m.M
	v6467 = m.ExcPending
	if v6467 != 0 {
		goto L1
	} else {
		goto L956
	}
L955:
	;
	goto L953
L956:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6380)+4)) = v6466
	v6469 = int32(1)
	if base.Ui32(v6469) < base.Ui32(v6443) {
		v6440 = v6466
		v6443 = v6443 - v6469
		goto L954
	} else {
		goto L957
	}
L957:
	;
	goto L955
L958:
	;
	v6526 = int32(1)
	v6531 = v6514
	goto L961
L959:
	;
	v6581 = v6514
	goto L960
L960:
	;
	v6608 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+16))
	F_pfree(m, v6608)
	mBase = m.M
	v6610 = m.ExcPending
	if v6610 != 0 {
		goto L1
	} else {
		goto L965
	}
L961:
	;
	v6561 = *(*int32)(unsafe.Add(mBase, uint32(v6380+v6526<<(uint(int32(2))%32))))
	v6562 = F_lappend(m, v6531, v6561)
	mBase = m.M
	v6563 = m.ExcPending
	if v6563 != 0 {
		goto L1
	} else {
		goto L963
	}
L962:
	;
	v6581 = v6562
	goto L960
L963:
	;
	v6565 = v6526 + int32(1)
	if v6565 <= v6397 {
		v6526 = v6565
		v6531 = v6562
		goto L961
	} else {
		goto L964
	}
L964:
	;
	goto L962
L965:
	;
	v6611 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+20))
	F_pfree(m, v6611)
	mBase = m.M
	v6613 = m.ExcPending
	if v6613 != 0 {
		goto L1
	} else {
		goto L966
	}
L966:
	;
	v6614 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+24))
	F_pfree(m, v6614)
	mBase = m.M
	v6616 = m.ExcPending
	if v6616 != 0 {
		goto L1
	} else {
		goto L967
	}
L967:
	;
	v6617 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+28))
	F_pfree(m, v6617)
	mBase = m.M
	v6619 = m.ExcPending
	if v6619 != 0 {
		goto L1
	} else {
		goto L968
	}
L968:
	;
	F_pfree(m, v5552)
	mBase = m.M
	v6621 = m.ExcPending
	if v6621 != 0 {
		goto L1
	} else {
		goto L969
	}
L969:
	;
	F_pfree(m, v6380)
	mBase = m.M
	v6623 = m.ExcPending
	if v6623 != 0 {
		goto L1
	} else {
		goto L970
	}
L970:
	;
	F_pfree(m, v6223)
	mBase = m.M
	v6625 = m.ExcPending
	if v6625 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	if int32(0) < v5550 {
		goto L973
	} else {
		goto L974
	}
L972:
	;
	F_pfree(m, v4872)
	mBase = m.M
	v6784 = m.ExcPending
	if v6784 != 0 {
		goto L1
	} else {
		goto L993
	}
L973:
	;
	v6630 = int32(1)
	goto L976
L974:
	;
	goto L975
L975:
	;
	F_pfree(m, v4874)
	mBase = m.M
	v6737 = m.ExcPending
	if v6737 != 0 {
		goto L1
	} else {
		goto L990
	}
L976:
	;
	v6673 = *(*int32)(unsafe.Add(mBase, uint32(v4874+v6630<<(uint(int32(2))%32))))
	if v6673 != 0 {
		goto L978
	} else {
		goto L979
	}
L977:
	;
	F_pfree(m, v4874)
	mBase = m.M
	v6680 = m.ExcPending
	if v6680 != 0 {
		goto L1
	} else {
		goto L983
	}
L978:
	;
	F_pfree(m, v6673)
	mBase = m.M
	v6675 = m.ExcPending
	if v6675 != 0 {
		goto L1
	} else {
		goto L981
	}
L979:
	;
	goto L980
L980:
	;
	v6677 = v6630 + int32(1)
	if v6677 <= v5550 {
		v6630 = v6677
		goto L976
	} else {
		goto L982
	}
L981:
	;
	goto L980
L982:
	;
	goto L977
L983:
	;
	F_pfree(m, v4878)
	mBase = m.M
	v6682 = m.ExcPending
	if v6682 != 0 {
		goto L1
	} else {
		goto L984
	}
L984:
	;
	F_pfree(m, v4870)
	mBase = m.M
	v6684 = m.ExcPending
	if v6684 != 0 {
		goto L1
	} else {
		goto L985
	}
L985:
	;
	v6687 = int32(1)
	goto L986
L986:
	;
	v6730 = *(*int32)(unsafe.Add(mBase, uint32(v4872+v6687<<(uint(int32(2))%32))))
	F_bms_free(m, v6730)
	mBase = m.M
	v6732 = m.ExcPending
	if v6732 != 0 {
		goto L1
	} else {
		goto L988
	}
L987:
	;
	goto L972
L988:
	;
	v6734 = v6687 + int32(1)
	if v6734 != v5528 {
		v6687 = v6734
		goto L986
	} else {
		goto L989
	}
L989:
	;
	goto L987
L990:
	;
	F_pfree(m, v4878)
	mBase = m.M
	v6739 = m.ExcPending
	if v6739 != 0 {
		goto L1
	} else {
		goto L991
	}
L991:
	;
	F_pfree(m, v4870)
	mBase = m.M
	v6741 = m.ExcPending
	if v6741 != 0 {
		goto L1
	} else {
		goto L992
	}
L992:
	;
	goto L972
L993:
	;
	v6799 = v6581
	goto L587
L994:
	;
	v6828 = *(*int32)(unsafe.Add(mBase, uint32(v6799)+4))
	if v6828 <= int32(0) {
		v7872 = v47
		v7875 = v3642
		v7877 = v3644
		v7879 = v4420
		v7881 = v4422
		v7884 = l7
		v7887 = v44
		v7893 = v9
		v7898 = v3756
		v7902 = v3757
		v7903 = v3758
		goto L586
	} else {
		goto L995
	}
L995:
	;
	v6840 = v47
	v6843 = v3642
	v6845 = v3644
	v6846 = v6799
	v6847 = v4420
	v6849 = v4422
	v6850 = int32(0)
	v6852 = l7
	v6855 = v44
	v6861 = v9
	v6866 = v3756
	v6870 = v3757
	v6871 = v3758
	goto L996
L996:
	;
	v6873 = *(*int32)(unsafe.Add(mBase, uint32(v6846)+12))
	v6877 = *(*int32)(unsafe.Add(mBase, uint32(v6873+v6850<<(uint(int32(2))%32))))
	v6879 = F_palloc0(m, int32(32))
	mBase = m.M
	v6880 = m.ExcPending
	if v6880 != 0 {
		goto L1
	} else {
		goto L998
	}
L997:
	;
	v7872 = v6840
	v7875 = v6843
	v7877 = v6845
	v7879 = v6847
	v7881 = v6849
	v7884 = v6852
	v7887 = v6855
	v7893 = v6861
	v7898 = v6866
	v7902 = v6870
	v7903 = v6871
	goto L586
L998:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6879))) = int32(311)
	v6883 = int32(0)
	v6885 = *(*int32)(unsafe.Add(mBase, uint32(v6846)+4))
	if v6885 == int32(1) {
		goto L999
	} else {
		goto L1000
	}
L999:
	;
	v6888 = *(*int32)(unsafe.Add(mBase, uint32(v6847)+124))
	v6889 = v6888
	goto L1001
L1000:
	;
	v6889 = v6883
	goto L1001
L1001:
	;
	v6890 = int32(0)
	if v6877 == v6890 {
		v7394 = v6890
		v7395 = v6883
		goto L1002
	} else {
		goto L1003
	}
L1002:
	;
	F_list_free(m, v7394)
	mBase = m.M
	v7427 = m.ExcPending
	if v7427 != 0 {
		goto L1
	} else {
		goto L1070
	}
L1003:
	;
	v6893 = int32(0)
	v6894 = *(*int32)(unsafe.Add(mBase, uint32(v6877)+4))
	if v6894 <= v6893 {
		v7394 = v6890
		v7395 = v6883
		goto L1002
	} else {
		goto L1004
	}
L1004:
	;
	v6897 = v6889
	v6900 = v6893
	v6906 = v6890
	v6907 = v6883
	goto L1005
L1005:
	;
	v6938 = *(*int32)(unsafe.Add(mBase, uint32(v6877)+12))
	v6942 = *(*int32)(unsafe.Add(mBase, uint32(v6938+v6900<<(uint(int32(2))%32))))
	v6943 = int32(0)
	if v6906 != 0 {
		goto L1009
	} else {
		goto L1010
	}
L1007:
	;
	v7263 = F_palloc0(m, int32(16))
	mBase = m.M
	v7264 = m.ExcPending
	if v7264 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1008:
	;
	v7261 = v7181
	goto L1007
L1009:
	;
	v6945 = int32(0)
	if v6942 == v6945 {
		v7261 = v6945
		goto L1007
	} else {
		goto L1012
	}
L1010:
	;
	goto L1011
L1011:
	;
	v7139 = int32(0)
	if v6942 == v7139 {
		v7261 = v7139
		goto L1007
	} else {
		goto L1026
	}
L1012:
	;
	v6948 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+4))
	if v6948 <= int32(0) {
		v7181 = v6943
		goto L1008
	} else {
		goto L1013
	}
L1013:
	;
	v6953 = v6943
	v6956 = v6943
	goto L1014
L1014:
	;
	v6992 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+12))
	v6996 = *(*int32)(unsafe.Add(mBase, uint32(v6992+v6956<<(uint(int32(2))%32))))
	v6997 = *(*int32)(unsafe.Add(mBase, uint32(v6906)+4))
	if int32(0) < v6997 {
		goto L1017
	} else {
		goto L1018
	}
L1015:
	;
	v7181 = v7096
	goto L1008
L1016:
	;
	v7136 = v6956 + int32(1)
	v7137 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+4))
	if v7136 < v7137 {
		v6953 = v7096
		v6956 = v7136
		goto L1014
	} else {
		goto L1025
	}
L1017:
	;
	v7000 = *(*int32)(unsafe.Add(mBase, uint32(v6906)+12))
	v7003 = int32(0)
	goto L1020
L1018:
	;
	goto L1019
L1019:
	;
	v7092 = F_lappend_int(m, v6953, v6996)
	mBase = m.M
	v7093 = m.ExcPending
	if v7093 != 0 {
		goto L1
	} else {
		goto L1024
	}
L1020:
	;
	v7046 = *(*int32)(unsafe.Add(mBase, uint32(v7000+v7003<<(uint(int32(2))%32))))
	if v7046 == v6996 {
		v7096 = v6953
		goto L1016
	} else {
		goto L1022
	}
L1021:
	;
	goto L1019
L1022:
	;
	v7049 = v7003 + int32(1)
	if v6997 != v7049 {
		v7003 = v7049
		goto L1020
	} else {
		goto L1023
	}
L1023:
	;
	goto L1021
L1024:
	;
	v7096 = v7092
	goto L1016
L1025:
	;
	goto L1015
L1026:
	;
	v7142 = *(*int32)(unsafe.Add(mBase, uint32(v6942)))
	v7145 = int32(8)
	v7146 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+4))
	v7148 = v7146 + int32(4)
	if v7148 <= v7145 {
		goto L1027
	} else {
		goto L1028
	}
L1027:
	;
	v7151 = v7145
	goto L1029
L1028:
	;
	v7151 = v7148
	goto L1029
L1029:
	;
	if v7151&(v7151-int32(1)) != 0 {
		goto L1030
	} else {
		goto L1031
	}
L1030:
	;
	v7158 = int32(1) << (uint(int32(32)-base.I32_clz(v7151)) % 32)
	goto L1032
L1031:
	;
	v7158 = v7151
	goto L1032
L1032:
	;
	v7160 = v7158 - int32(4)
	v7165 = F_palloc(m, v7160<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v7166 = m.ExcPending
	if v7166 != 0 {
		goto L1
	} else {
		goto L1033
	}
L1033:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7165)+8)) = v7160
	*(*int32)(unsafe.Add(mBase, uint32(v7165)+4)) = v7146
	*(*int32)(unsafe.Add(mBase, uint32(v7165))) = v7142
	v7171 = v7165 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v7165)+12)) = v7171
	v7174 = v7146 << (uint(int32(2)) % 32)
	if v7174 == int32(0) {
		v7181 = v7165
		goto L1008
	} else {
		goto L1034
	}
L1034:
	;
	v7177 = *(*int32)(unsafe.Add(mBase, uint32(v6942)+12))
	base.MemoryCopy(m, v7171, v7177, v7174)
	v7261 = v7165
	goto L1007
L1035:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7263))) = int32(310)
	v7268 = v7261
	v7276 = v6906
	goto L1036
L1036:
	;
	if v6897 != 0 {
		goto L1038
	} else {
		goto L1039
	}
L1038:
	;
	v7308 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+4))
	v7310 = v7308
	goto L1040
L1039:
	;
	v7310 = int32(0)
	goto L1040
L1040:
	;
	if v7276 == int32(0) {
		goto L1044
	} else {
		goto L1045
	}
L1041:
	;
	v7381 = F_lappend_int(m, v7276, v7327)
	mBase = m.M
	v7382 = m.ExcPending
	if v7382 != 0 {
		goto L1
	} else {
		goto L1068
	}
L1042:
	;
	v7370 = F_list_concat(m, v7276, v7268)
	mBase = m.M
	v7371 = m.ExcPending
	if v7371 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1043:
	;
	v7322 = *(*int32)(unsafe.Add(mBase, uint32(v6897)+12))
	v7326 = *(*int32)(unsafe.Add(mBase, uint32(v7322+v7321<<(uint(int32(2))%32))))
	v7327 = *(*int32)(unsafe.Add(mBase, uint32(v7326)+4))
	v7328 = int32(0)
	if v7268 == v7328 {
		goto L1051
	} else {
		goto L1052
	}
L1044:
	;
	if v7310 <= int32(0) {
		v7368 = v6897
		goto L1042
	} else {
		goto L1047
	}
L1045:
	;
	goto L1046
L1046:
	;
	v7318 = *(*int32)(unsafe.Add(mBase, uint32(v7276)+4))
	if base.B2i32(v7268 == int32(0))|base.B2i32(v7310 <= v7318) != 0 {
		v7368 = v6897
		goto L1042
	} else {
		goto L1049
	}
L1047:
	;
	if v7268 != 0 {
		v7321 = int32(0)
		goto L1043
	} else {
		goto L1048
	}
L1048:
	;
	v7368 = v6897
	goto L1042
L1049:
	;
	v7321 = v7318
	goto L1043
L1050:
	;
	if v7366 != 0 {
		goto L1041
	} else {
		goto L1063
	}
L1051:
	;
	v7366 = int32(0)
	goto L1050
L1052:
	;
	goto L1053
L1053:
	;
	v7334 = *(*int32)(unsafe.Add(mBase, uint32(v7268)+4))
	if v7334 <= int32(0) {
		v7360 = v7328
		goto L1054
	} else {
		goto L1055
	}
L1054:
	;
	v7366 = v7360
	goto L1050
L1055:
	;
	v7337 = int32(0)
	if v7337 < v7334 {
		goto L1056
	} else {
		goto L1057
	}
L1056:
	;
	v7340 = v7334
	goto L1058
L1057:
	;
	v7340 = v7337
	goto L1058
L1058:
	;
	v7341 = *(*int32)(unsafe.Add(mBase, uint32(v7268)+12))
	v7343 = int32(0)
	goto L1059
L1059:
	;
	v7351 = *(*int32)(unsafe.Add(mBase, uint32(v7341+v7343<<(uint(int32(2))%32))))
	v7352 = base.B2i32(v7351 == v7327)
	if v7351 == v7327 {
		v7360 = v7352
		goto L1054
	} else {
		goto L1061
	}
L1060:
	;
	v7360 = v7352
	goto L1054
L1061:
	;
	v7354 = v7343 + int32(1)
	if v7354 != v7340 {
		v7343 = v7354
		goto L1059
	} else {
		goto L1062
	}
L1062:
	;
	goto L1060
L1063:
	;
	v7368 = int32(0)
	goto L1042
L1064:
	;
	v7372 = F_list_copy(m, v7370)
	mBase = m.M
	v7373 = m.ExcPending
	if v7373 != 0 {
		goto L1
	} else {
		goto L1065
	}
L1065:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7263)+4)) = v7372
	v7375 = F_lcons(m, v7263, v6907)
	mBase = m.M
	v7376 = m.ExcPending
	if v7376 != 0 {
		goto L1
	} else {
		goto L1066
	}
L1066:
	;
	v7378 = v6900 + int32(1)
	v7379 = *(*int32)(unsafe.Add(mBase, uint32(v6877)+4))
	if v7378 < v7379 {
		v6897 = v7368
		v6900 = v7378
		v6906 = v7370
		v6907 = v7375
		goto L1005
	} else {
		goto L1067
	}
L1067:
	;
	v7394 = v7370
	v7395 = v7375
	goto L1002
L1068:
	;
	v7383 = F_list_delete_ptr(m, v7268, v7327)
	mBase = m.M
	v7384 = m.ExcPending
	if v7384 != 0 {
		goto L1
	} else {
		goto L1069
	}
L1069:
	;
	v7268 = v7383
	v7276 = v7381
	goto L1036
L1070:
	;
	v7428 = int32(0)
	v7429 = *(*int32)(unsafe.Add(mBase, uint32(v7395)+12))
	v7430 = *(*int32)(unsafe.Add(mBase, uint32(v7429)))
	v7431 = *(*int32)(unsafe.Add(mBase, uint32(v7430)+4))
	if v7431 == v7428 {
		v7495 = v7428
		goto L1071
	} else {
		goto L1072
	}
L1071:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6879)+4)) = v7495
	v7536 = *(*int32)(unsafe.Add(mBase, uint32(v7430)+4))
	if v7536 != 0 {
		goto L1079
	} else {
		goto L1080
	}
L1072:
	;
	v7434 = *(*int32)(unsafe.Add(mBase, uint32(v7431)+4))
	if v7434 <= int32(0) {
		v7495 = v7428
		goto L1071
	} else {
		goto L1073
	}
L1073:
	;
	v7437 = *(*int32)(unsafe.Add(mBase, uint32(v6840)+4))
	v7440 = v7428
	v7448 = int32(0)
	goto L1074
L1074:
	;
	v7480 = *(*int32)(unsafe.Add(mBase, uint32(v7431)+12))
	v7484 = *(*int32)(unsafe.Add(mBase, uint32(v7480+v7448<<(uint(int32(2))%32))))
	v7485 = *(*int32)(unsafe.Add(mBase, uint32(v7437)+100))
	v7486 = F_get_sortgroupref_clause(m, v7484, v7485)
	mBase = m.M
	v7487 = m.ExcPending
	if v7487 != 0 {
		goto L1
	} else {
		goto L1076
	}
L1075:
	;
	v7495 = v7488
	goto L1071
L1076:
	;
	v7488 = F_lappend(m, v7440, v7486)
	mBase = m.M
	v7489 = m.ExcPending
	if v7489 != 0 {
		goto L1
	} else {
		goto L1077
	}
L1077:
	;
	v7491 = v7448 + int32(1)
	v7492 = *(*int32)(unsafe.Add(mBase, uint32(v7431)+4))
	if v7491 < v7492 {
		v7440 = v7488
		v7448 = v7491
		goto L1074
	} else {
		goto L1078
	}
L1078:
	;
	goto L1075
L1079:
	;
	v7537 = *(*int32)(unsafe.Add(mBase, uint32(v6849)+24))
	v7538 = F_bms_overlap_list(m, v7537, v7536)
	mBase = m.M
	v7539 = m.ExcPending
	if v7539 != 0 {
		goto L1
	} else {
		goto L1082
	}
L1080:
	;
	v7547 = v7495
	goto L1081
L1081:
	;
	v7548 = *(*int32)(unsafe.Add(mBase, uint32(v6849)+32))
	if v7547 == int32(0) {
		goto L1086
	} else {
		goto L1087
	}
L1082:
	;
	if v7538 == int32(0) {
		goto L1083
	} else {
		goto L1084
	}
L1083:
	;
	v7542 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6879)+24)) = uint8(v7542)
	*(*uint8)(unsafe.Add(mBase, uint32(v6849)+16)) = uint8(v7542)
	goto L1085
L1084:
	;
	goto L1085
L1085:
	;
	v7546 = *(*int32)(unsafe.Add(mBase, uint32(v6879)+4))
	v7547 = v7546
	goto L1081
L1086:
	;
	v7651 = int32(0)
	v7653 = *(*int32)(unsafe.Add(mBase, uint32(v7395)+4))
	if v7651 < v7653 {
		goto L1092
	} else {
		goto L1093
	}
L1087:
	;
	v7551 = int32(0)
	v7552 = *(*int32)(unsafe.Add(mBase, uint32(v7547)+4))
	if v7552 <= v7551 {
		goto L1086
	} else {
		goto L1088
	}
L1088:
	;
	v7564 = v7551
	goto L1089
L1089:
	;
	v7596 = *(*int32)(unsafe.Add(mBase, uint32(v7547)+12))
	v7597 = int32(2)
	v7600 = *(*int32)(unsafe.Add(mBase, uint32(v7596+v7564<<(uint(v7597)%32))))
	v7601 = *(*int32)(unsafe.Add(mBase, uint32(v7600)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7548+v7601<<(uint(v7597)%32)))) = v7564
	v7607 = v7564 + int32(1)
	v7608 = *(*int32)(unsafe.Add(mBase, uint32(v7547)+4))
	if v7607 < v7608 {
		v7564 = v7607
		goto L1089
	} else {
		goto L1091
	}
L1090:
	;
	goto L1086
L1091:
	;
	goto L1090
L1092:
	;
	v7659 = v7651
	v7661 = v7651
	goto L1095
L1093:
	;
	v7818 = v7651
	goto L1094
L1094:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6879)+12)) = v7395
	*(*int32)(unsafe.Add(mBase, uint32(v6879)+8)) = v7818
	v7856 = *(*int32)(unsafe.Add(mBase, uint32(v6849)))
	v7857 = F_lappend(m, v7856, v6879)
	mBase = m.M
	v7858 = m.ExcPending
	if v7858 != 0 {
		goto L1
	} else {
		goto L1106
	}
L1095:
	;
	v7697 = int32(0)
	v7698 = *(*int32)(unsafe.Add(mBase, uint32(v7395)+12))
	v7702 = *(*int32)(unsafe.Add(mBase, uint32(v7698+v7659<<(uint(int32(2))%32))))
	v7703 = *(*int32)(unsafe.Add(mBase, uint32(v7702)+4))
	if v7703 == v7697 {
		v7775 = v7697
		goto L1097
	} else {
		goto L1098
	}
L1096:
	;
	v7818 = v7807
	goto L1094
L1097:
	;
	v7807 = F_lappend(m, v7661, v7775)
	mBase = m.M
	v7808 = m.ExcPending
	if v7808 != 0 {
		goto L1
	} else {
		goto L1104
	}
L1098:
	;
	v7706 = int32(0)
	v7707 = *(*int32)(unsafe.Add(mBase, uint32(v7703)+4))
	if v7707 <= v7706 {
		v7775 = v7697
		goto L1097
	} else {
		goto L1099
	}
L1099:
	;
	v7711 = v7706
	v7719 = v7697
	goto L1100
L1100:
	;
	v7751 = *(*int32)(unsafe.Add(mBase, uint32(v7703)+12))
	v7752 = int32(2)
	v7755 = *(*int32)(unsafe.Add(mBase, uint32(v7751+v7711<<(uint(v7752)%32))))
	v7759 = *(*int32)(unsafe.Add(mBase, uint32(v7548+v7755<<(uint(v7752)%32))))
	v7760 = F_lappend_int(m, v7719, v7759)
	mBase = m.M
	v7761 = m.ExcPending
	if v7761 != 0 {
		goto L1
	} else {
		goto L1102
	}
L1101:
	;
	v7775 = v7760
	goto L1097
L1102:
	;
	v7763 = v7711 + int32(1)
	v7764 = *(*int32)(unsafe.Add(mBase, uint32(v7703)+4))
	if v7763 < v7764 {
		v7711 = v7763
		v7719 = v7760
		goto L1100
	} else {
		goto L1103
	}
L1103:
	;
	goto L1101
L1104:
	;
	v7810 = v7659 + int32(1)
	v7811 = *(*int32)(unsafe.Add(mBase, uint32(v7395)+4))
	if v7810 < v7811 {
		v7659 = v7810
		v7661 = v7807
		goto L1095
	} else {
		goto L1105
	}
L1105:
	;
	goto L1096
L1106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6849))) = v7857
	v7861 = v6850 + int32(1)
	v7862 = *(*int32)(unsafe.Add(mBase, uint32(v6846)+4))
	if v7861 < v7862 {
		v6850 = v7861
		goto L996
	} else {
		goto L1107
	}
L1107:
	;
	goto L997
L1108:
	;
	v7908 = *(*int32)(unsafe.Add(mBase, uint32(v7881)+32))
	v7909 = *(*int32)(unsafe.Add(mBase, uint32(v7879)+100))
	if v7909 == int32(0) {
		goto L1109
	} else {
		goto L1110
	}
L1109:
	;
	v8012 = int32(0)
	v8013 = *(*int32)(unsafe.Add(mBase, uint32(v7905)+4))
	if v8012 < v8013 {
		goto L1115
	} else {
		goto L1116
	}
L1110:
	;
	v7912 = *(*int32)(unsafe.Add(mBase, uint32(v7909)+4))
	if v7912 <= int32(0) {
		goto L1109
	} else {
		goto L1111
	}
L1111:
	;
	v7917 = int32(0)
	goto L1112
L1112:
	;
	v7957 = *(*int32)(unsafe.Add(mBase, uint32(v7909)+12))
	v7958 = int32(2)
	v7961 = *(*int32)(unsafe.Add(mBase, uint32(v7957+v7917<<(uint(v7958)%32))))
	v7962 = *(*int32)(unsafe.Add(mBase, uint32(v7961)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7908+v7962<<(uint(v7958)%32)))) = v7917
	v7968 = v7917 + int32(1)
	v7969 = *(*int32)(unsafe.Add(mBase, uint32(v7909)+4))
	if v7968 < v7969 {
		v7917 = v7968
		goto L1112
	} else {
		goto L1114
	}
L1113:
	;
	goto L1109
L1114:
	;
	goto L1113
L1115:
	;
	v8022 = v8012
	v8027 = int32(0)
	goto L1118
L1116:
	;
	v8180 = v8012
	goto L1117
L1117:
	;
	v8216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v7881)+16)) = uint8(v8216)
	*(*int32)(unsafe.Add(mBase, uint32(v7881)+4)) = v8180
	v8227 = v7872
	v8230 = v7875
	v8232 = v7877
	v8236 = v7881
	v8239 = v7884
	v8242 = v7887
	v8248 = v7893
	v8253 = v7898
	v8257 = v7902
	v8258 = v7903
	goto L585
L1118:
	;
	v8058 = *(*int32)(unsafe.Add(mBase, uint32(v7905)+12))
	v8062 = *(*int32)(unsafe.Add(mBase, uint32(v8058+v8027<<(uint(int32(2))%32))))
	v8063 = *(*int32)(unsafe.Add(mBase, uint32(v8062)+4))
	if v8063 == int32(0) {
		goto L1121
	} else {
		goto L1122
	}
L1119:
	;
	v8180 = v8169
	goto L1117
L1120:
	;
	v8169 = F_lappend(m, v8022, v8137)
	mBase = m.M
	v8170 = m.ExcPending
	if v8170 != 0 {
		goto L1
	} else {
		goto L1129
	}
L1121:
	;
	v8137 = int32(0)
	goto L1120
L1122:
	;
	goto L1123
L1123:
	;
	v8067 = int32(0)
	v8069 = *(*int32)(unsafe.Add(mBase, uint32(v8063)+4))
	if v8069 <= v8067 {
		v8137 = v8067
		goto L1120
	} else {
		goto L1124
	}
L1124:
	;
	v8073 = v8067
	v8081 = v8067
	goto L1125
L1125:
	;
	v8113 = *(*int32)(unsafe.Add(mBase, uint32(v8063)+12))
	v8114 = int32(2)
	v8117 = *(*int32)(unsafe.Add(mBase, uint32(v8113+v8073<<(uint(v8114)%32))))
	v8121 = *(*int32)(unsafe.Add(mBase, uint32(v7908+v8117<<(uint(v8114)%32))))
	v8122 = F_lappend_int(m, v8081, v8121)
	mBase = m.M
	v8123 = m.ExcPending
	if v8123 != 0 {
		goto L1
	} else {
		goto L1127
	}
L1126:
	;
	v8137 = v8122
	goto L1120
L1127:
	;
	v8125 = v8073 + int32(1)
	v8126 = *(*int32)(unsafe.Add(mBase, uint32(v8063)+4))
	if v8125 < v8126 {
		v8073 = v8125
		v8081 = v8122
		goto L1125
	} else {
		goto L1128
	}
L1128:
	;
	goto L1126
L1129:
	;
	v8172 = v8027 + int32(1)
	v8173 = *(*int32)(unsafe.Add(mBase, uint32(v7905)+4))
	if v8172 < v8173 {
		v8022 = v8169
		v8027 = v8172
		goto L1118
	} else {
		goto L1130
	}
L1130:
	;
	goto L1119
L1131:
	;
	v8979 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+144))
	if v8979 == int32(0) {
		v9156 = v8938
		goto L1224
	} else {
		goto L1225
	}
L1132:
	;
	v8858 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+72))
	v8860 = F_pull_var_clause(m, v8858, int32(16))
	mBase = m.M
	v8861 = m.ExcPending
	if v8861 != 0 {
		goto L1
	} else {
		goto L1206
	}
L1133:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v8807 = m.ExcPending
	if v8807 != 0 {
		goto L1
	} else {
		goto L1203
	}
L1134:
	;
	v8272 = *(*int32)(unsafe.Add(mBase, uint32(v8270)+12))
	v8278 = *(*int32)(unsafe.Add(mBase, uint32(v8272+v8271<<(uint(int32(2))%32)-int32(4))))
	v8279 = *(*int32)(unsafe.Add(mBase, uint32(v8278)+12))
	if v8279 != 0 {
		goto L1133
	} else {
		goto L1137
	}
L1135:
	;
	v8285 = v8260
	v8286 = int32(0)
	goto L1136
L1136:
	;
	v8287 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+76))
	switch v8269 - int32(2) {
	case 0:
		goto L1140
	case 1:
		goto L1141
	default:
		goto L1139
	}
L1137:
	;
	v8280 = *(*int32)(unsafe.Add(mBase, uint32(v8278)+16))
	v8282 = F_table_open(m, v8280, int32(0))
	mBase = m.M
	v8283 = m.ExcPending
	if v8283 != 0 {
		goto L1
	} else {
		goto L1138
	}
L1138:
	;
	v8285 = v8278
	v8286 = v8282
	goto L1136
L1139:
	;
	if base.B2i32(int32(1)<<(uint(v8269)%32)&int32(52) == int32(0))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v8269)) != 0 {
		v8938 = v8287
		goto L1131
	} else {
		goto L1153
	}
L1140:
	;
	if v8287 == int32(0) {
		v8362 = v8260
		goto L1143
	} else {
		goto L1144
	}
L1141:
	;
	v8290 = F_expand_insert_targetlist(m, v8227, v8287, v8286)
	mBase = m.M
	v8291 = m.ExcPending
	if v8291 != 0 {
		goto L1
	} else {
		goto L1142
	}
L1142:
	;
	v8938 = v8290
	goto L1131
L1143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+288)) = v8362
	goto L1139
L1144:
	;
	v8295 = *(*int32)(unsafe.Add(mBase, uint32(v8287)+4))
	if v8295 <= int32(0) {
		v8362 = v8260
		goto L1143
	} else {
		goto L1145
	}
L1145:
	;
	v8299 = int32(1)
	v8300 = v8260
	v8302 = v8260
	goto L1146
L1146:
	;
	v8339 = *(*int32)(unsafe.Add(mBase, uint32(v8287)+12))
	v8343 = *(*int32)(unsafe.Add(mBase, uint32(v8339+v8300<<(uint(int32(2))%32))))
	v8344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8343)+26)))
	if v8344 == int32(0) {
		goto L1148
	} else {
		goto L1149
	}
L1147:
	;
	v8362 = v8350
	goto L1143
L1148:
	;
	v8347 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8343)+8)))
	v8348 = F_lappend_int(m, v8302, v8347)
	mBase = m.M
	v8349 = m.ExcPending
	if v8349 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1149:
	;
	v8350 = v8302
	goto L1150
L1150:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8343)+8)) = uint16(v8299)
	v8352 = int32(1)
	v8355 = v8300 + v8352
	v8356 = *(*int32)(unsafe.Add(mBase, uint32(v8287)+4))
	if v8355 < v8356 {
		v8299 = v8299 + v8352
		v8300 = v8355
		v8302 = v8350
		goto L1146
	} else {
		goto L1152
	}
L1151:
	;
	v8350 = v8348
	goto L1150
L1152:
	;
	goto L1147
L1153:
	;
	v8450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8285)+20)))
	if v8450 == int32(0) {
		goto L1155
	} else {
		goto L1156
	}
L1154:
	;
	v8462 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+64))
	if v8462 == int32(0) {
		v8817 = v8461
		goto L1132
	} else {
		goto L1161
	}
L1155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+284)) = v8287
	F_add_row_identity_columns(m, v8227, v8271, v8285, v8286)
	mBase = m.M
	v8455 = m.ExcPending
	if v8455 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1156:
	;
	goto L1157
L1157:
	;
	if v8269 != int32(5) {
		v8938 = v8287
		goto L1131
	} else {
		goto L1160
	}
L1158:
	;
	v8456 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	if v8269 == int32(5) {
		v8461 = v8456
		goto L1154
	} else {
		goto L1159
	}
L1159:
	;
	v8938 = v8456
	goto L1131
L1160:
	;
	v8461 = v8287
	goto L1154
L1161:
	;
	v8465 = *(*int32)(unsafe.Add(mBase, uint32(v8462)+4))
	if v8465 <= int32(0) {
		v8817 = v8461
		goto L1132
	} else {
		goto L1162
	}
L1162:
	;
	v8468 = v8461
	v8489 = v8260
	goto L1163
L1163:
	;
	v8509 = *(*int32)(unsafe.Add(mBase, uint32(v8462)+12))
	v8510 = int32(2)
	v8513 = *(*int32)(unsafe.Add(mBase, uint32(v8509+v8489<<(uint(v8510)%32))))
	v8514 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+8))
	switch v8514 - v8510 {
	case 0:
		goto L1166
	case 1:
		goto L1167
	default:
		goto L1165
	}
L1164:
	;
	v8817 = v8757
	goto L1132
L1165:
	;
	v8674 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+16))
	v8675 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+20))
	v8676 = F_list_concat_copy(m, v8674, v8675)
	mBase = m.M
	v8677 = m.ExcPending
	if v8677 != 0 {
		goto L1
	} else {
		goto L1182
	}
L1166:
	;
	v8521 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+20))
	if v8521 == int32(0) {
		goto L1170
	} else {
		goto L1171
	}
L1167:
	;
	v8517 = *(*int32)(unsafe.Add(mBase, uint32(v8513)+20))
	v8518 = F_expand_insert_targetlist(m, v8227, v8517, v8286)
	mBase = m.M
	v8519 = m.ExcPending
	if v8519 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8513)+20)) = v8518
	goto L1165
L1169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8513)+24)) = v8598
	goto L1165
L1170:
	;
	v8598 = int32(0)
	goto L1169
L1171:
	;
	goto L1172
L1172:
	;
	v8526 = int32(0)
	v8528 = *(*int32)(unsafe.Add(mBase, uint32(v8521)+4))
	if v8528 <= v8526 {
		v8598 = v8526
		goto L1169
	} else {
		goto L1173
	}
L1173:
	;
	v8532 = int32(1)
	v8533 = v8526
	v8538 = v8526
	goto L1174
L1174:
	;
	v8572 = *(*int32)(unsafe.Add(mBase, uint32(v8521)+12))
	v8576 = *(*int32)(unsafe.Add(mBase, uint32(v8572+v8533<<(uint(int32(2))%32))))
	v8577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8576)+26)))
	if v8577 == int32(0) {
		goto L1176
	} else {
		goto L1177
	}
L1175:
	;
	v8598 = v8583
	goto L1169
L1176:
	;
	v8580 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8576)+8)))
	v8581 = F_lappend_int(m, v8538, v8580)
	mBase = m.M
	v8582 = m.ExcPending
	if v8582 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1177:
	;
	v8583 = v8538
	goto L1178
L1178:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v8576)+8)) = uint16(v8532)
	v8585 = int32(1)
	v8588 = v8533 + v8585
	v8589 = *(*int32)(unsafe.Add(mBase, uint32(v8521)+4))
	if v8588 < v8589 {
		v8532 = v8532 + v8585
		v8533 = v8588
		v8538 = v8583
		goto L1174
	} else {
		goto L1180
	}
L1179:
	;
	v8583 = v8581
	goto L1178
L1180:
	;
	goto L1175
L1181:
	;
	F_list_free(m, v8679)
	mBase = m.M
	v8799 = m.ExcPending
	if v8799 != 0 {
		goto L1
	} else {
		goto L1201
	}
L1182:
	;
	v8679 = F_pull_var_clause(m, v8676, int32(16))
	mBase = m.M
	v8680 = m.ExcPending
	if v8680 != 0 {
		goto L1
	} else {
		goto L1183
	}
L1183:
	;
	if v8679 == int32(0) {
		v8757 = v8468
		goto L1181
	} else {
		goto L1184
	}
L1184:
	;
	v8683 = int32(0)
	v8684 = *(*int32)(unsafe.Add(mBase, uint32(v8679)+4))
	if v8684 <= v8683 {
		v8757 = v8468
		goto L1181
	} else {
		goto L1185
	}
L1185:
	;
	v8687 = v8468
	v8688 = v8683
	goto L1186
L1186:
	;
	v8728 = *(*int32)(unsafe.Add(mBase, uint32(v8679)+12))
	v8732 = *(*int32)(unsafe.Add(mBase, uint32(v8728+v8688<<(uint(int32(2))%32))))
	v8733 = *(*int32)(unsafe.Add(mBase, uint32(v8732)))
	if v8733 == int32(6) {
		goto L1189
	} else {
		goto L1190
	}
L1187:
	;
	v8757 = v8752
	goto L1181
L1188:
	;
	v8754 = v8688 + int32(1)
	v8755 = *(*int32)(unsafe.Add(mBase, uint32(v8679)+4))
	if v8754 < v8755 {
		v8687 = v8752
		v8688 = v8754
		goto L1186
	} else {
		goto L1200
	}
L1189:
	;
	v8736 = *(*int32)(unsafe.Add(mBase, uint32(v8732)+4))
	if v8736 == v8271 {
		v8752 = v8687
		goto L1188
	} else {
		goto L1192
	}
L1190:
	;
	goto L1191
L1191:
	;
	v8738 = F_tlist_member(m, v8732, v8687)
	mBase = m.M
	v8739 = m.ExcPending
	if v8739 != 0 {
		goto L1
	} else {
		goto L1193
	}
L1192:
	;
	goto L1191
L1193:
	;
	if v8738 != 0 {
		v8752 = v8687
		goto L1188
	} else {
		goto L1194
	}
L1194:
	;
	if v8687 != 0 {
		goto L1195
	} else {
		goto L1196
	}
L1195:
	;
	v8740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8687)+4)))
	v8744 = v8740 + int32(1)
	goto L1197
L1196:
	;
	v8744 = int32(1)
	goto L1197
L1197:
	;
	v8748 = F_makeTargetEntry(m, v8732, base.I32_extend16_s(v8744), int32(0), int32(1))
	mBase = m.M
	v8749 = m.ExcPending
	if v8749 != 0 {
		goto L1
	} else {
		goto L1198
	}
L1198:
	;
	v8750 = F_lappend(m, v8687, v8748)
	mBase = m.M
	v8751 = m.ExcPending
	if v8751 != 0 {
		goto L1
	} else {
		goto L1199
	}
L1199:
	;
	v8752 = v8750
	goto L1188
L1200:
	;
	goto L1187
L1201:
	;
	v8801 = v8489 + int32(1)
	v8802 = *(*int32)(unsafe.Add(mBase, uint32(v8462)+4))
	if v8801 < v8802 {
		v8468 = v8757
		v8489 = v8801
		goto L1163
	} else {
		goto L1202
	}
L1202:
	;
	goto L1164
L1203:
	;
	F_errmsg_internal(m, int32(_a_F_subquery_planner_32), int32(0))
	mBase = m.M
	v8811 = m.ExcPending
	if v8811 != 0 {
		goto L1
	} else {
		goto L1204
	}
L1204:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_33), int32(91), int32(_a_F_subquery_planner_34))
	mBase = m.M
	v8816 = m.ExcPending
	if v8816 != 0 {
		goto L1
	} else {
		goto L1205
	}
L1205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1206:
	;
	if v8860 == int32(0) {
		v8938 = v8817
		goto L1131
	} else {
		goto L1207
	}
L1207:
	;
	v8864 = int32(0)
	v8865 = *(*int32)(unsafe.Add(mBase, uint32(v8860)+4))
	if v8865 <= v8864 {
		v8938 = v8817
		goto L1131
	} else {
		goto L1208
	}
L1208:
	;
	v8868 = v8817
	v8869 = v8864
	goto L1209
L1209:
	;
	v8909 = *(*int32)(unsafe.Add(mBase, uint32(v8860)+12))
	v8913 = *(*int32)(unsafe.Add(mBase, uint32(v8909+v8869<<(uint(int32(2))%32))))
	v8914 = *(*int32)(unsafe.Add(mBase, uint32(v8913)))
	if v8914 == int32(6) {
		goto L1212
	} else {
		goto L1213
	}
L1210:
	;
	v8938 = v8933
	goto L1131
L1211:
	;
	v8935 = v8869 + int32(1)
	v8936 = *(*int32)(unsafe.Add(mBase, uint32(v8860)+4))
	if v8935 < v8936 {
		v8868 = v8933
		v8869 = v8935
		goto L1209
	} else {
		goto L1223
	}
L1212:
	;
	v8917 = *(*int32)(unsafe.Add(mBase, uint32(v8913)+4))
	if v8917 == v8271 {
		v8933 = v8868
		goto L1211
	} else {
		goto L1215
	}
L1213:
	;
	goto L1214
L1214:
	;
	v8919 = F_tlist_member(m, v8913, v8868)
	mBase = m.M
	v8920 = m.ExcPending
	if v8920 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1215:
	;
	goto L1214
L1216:
	;
	if v8919 != 0 {
		v8933 = v8868
		goto L1211
	} else {
		goto L1217
	}
L1217:
	;
	if v8868 != 0 {
		goto L1218
	} else {
		goto L1219
	}
L1218:
	;
	v8921 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8868)+4)))
	v8925 = v8921 + int32(1)
	goto L1220
L1219:
	;
	v8925 = int32(1)
	goto L1220
L1220:
	;
	v8929 = F_makeTargetEntry(m, v8913, base.I32_extend16_s(v8925), int32(0), int32(1))
	mBase = m.M
	v8930 = m.ExcPending
	if v8930 != 0 {
		goto L1
	} else {
		goto L1221
	}
L1221:
	;
	v8931 = F_lappend(m, v8868, v8929)
	mBase = m.M
	v8932 = m.ExcPending
	if v8932 != 0 {
		goto L1
	} else {
		goto L1222
	}
L1222:
	;
	v8933 = v8931
	goto L1211
L1223:
	;
	goto L1210
L1224:
	;
	v9197 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+96))
	if v9197 == int32(0) {
		v9328 = v9156
		goto L1263
	} else {
		goto L1264
	}
L1225:
	;
	v8982 = *(*int32)(unsafe.Add(mBase, uint32(v8979)+4))
	if v8982 <= int32(0) {
		v9156 = v8938
		goto L1224
	} else {
		goto L1226
	}
L1226:
	;
	v8986 = v8938
	v8988 = int32(0)
	goto L1227
L1227:
	;
	v9027 = *(*int32)(unsafe.Add(mBase, uint32(v8979)+12))
	v9031 = *(*int32)(unsafe.Add(mBase, uint32(v9027+v8988<<(uint(int32(2))%32))))
	v9032 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+4))
	v9033 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+8))
	if v9032 != v9033 {
		v9149 = v8986
		goto L1229
	} else {
		goto L1230
	}
L1228:
	;
	v9156 = v9149
	goto L1224
L1229:
	;
	v9153 = v8988 + int32(1)
	v9154 = *(*int32)(unsafe.Add(mBase, uint32(v8979)+4))
	if v9153 < v9154 {
		v8986 = v9149
		v8988 = v9153
		goto L1227
	} else {
		goto L1262
	}
L1230:
	;
	v9035 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+20))
	if v9035&int32(-33) != 0 {
		goto L1231
	} else {
		goto L1232
	}
L1231:
	;
	v9038 = int32(-1)
	v9041 = int32(0)
	v9043 = F_makeVar(m, v9032, v9038, int32(27), v9038, v9041, v9041)
	mBase = m.M
	v9044 = m.ExcPending
	if v9044 != 0 {
		goto L1
	} else {
		goto L1234
	}
L1232:
	;
	v9071 = v8986
	v9073 = v9035
	goto L1233
L1233:
	;
	if v9073&int32(32) != 0 {
		goto L1242
	} else {
		goto L1243
	}
L1234:
	;
	v9045 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8266)+32)) = v9045
	v9049 = int32(32)
	v9053 = F_pg_snprintf(m, v8266+int32(48), v9049, int32(_a_F_subquery_planner_35), v8266+v9049)
	mBase = m.M
	v9054 = m.ExcPending
	if v9054 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	if v8986 != 0 {
		goto L1236
	} else {
		goto L1237
	}
L1236:
	;
	v9055 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8986)+4)))
	v9059 = v9055 + int32(1)
	goto L1238
L1237:
	;
	v9059 = int32(1)
	goto L1238
L1238:
	;
	v9063 = F_pstrdup(m, v8266+int32(48))
	mBase = m.M
	v9064 = m.ExcPending
	if v9064 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	v9066 = F_makeTargetEntry(m, v9043, base.I32_extend16_s(v9059), v9063, int32(1))
	mBase = m.M
	v9067 = m.ExcPending
	if v9067 != 0 {
		goto L1
	} else {
		goto L1240
	}
L1240:
	;
	v9068 = F_lappend(m, v8986, v9066)
	mBase = m.M
	v9069 = m.ExcPending
	if v9069 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	v9070 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+20))
	v9071 = v9068
	v9073 = v9070
	goto L1233
L1242:
	;
	v9076 = *(*int32)(unsafe.Add(mBase, uint32(v8270)+12))
	v9077 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+4))
	v9083 = *(*int32)(unsafe.Add(mBase, uint32(v9076+v9077<<(uint(int32(2))%32)-int32(4))))
	v9084 = int32(0)
	v9086 = F_makeWholeRowVar(m, v9083, v9077, v9084, v9084)
	mBase = m.M
	v9087 = m.ExcPending
	if v9087 != 0 {
		goto L1
	} else {
		goto L1245
	}
L1243:
	;
	v9113 = v9071
	goto L1244
L1244:
	;
	v9115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9031)+32)))
	if v9115 != int32(1) {
		v9149 = v9113
		goto L1229
	} else {
		goto L1253
	}
L1245:
	;
	v9088 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8266)+16)) = v9088
	v9096 = F_pg_snprintf(m, v8266+int32(48), int32(32), int32(_a_F_subquery_planner_36), v8266+int32(16))
	mBase = m.M
	v9097 = m.ExcPending
	if v9097 != 0 {
		goto L1
	} else {
		goto L1246
	}
L1246:
	;
	if v9071 != 0 {
		goto L1247
	} else {
		goto L1248
	}
L1247:
	;
	v9098 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9071)+4)))
	v9102 = v9098 + int32(1)
	goto L1249
L1248:
	;
	v9102 = int32(1)
	goto L1249
L1249:
	;
	v9106 = F_pstrdup(m, v8266+int32(48))
	mBase = m.M
	v9107 = m.ExcPending
	if v9107 != 0 {
		goto L1
	} else {
		goto L1250
	}
L1250:
	;
	v9109 = F_makeTargetEntry(m, v9086, base.I32_extend16_s(v9102), v9106, int32(1))
	mBase = m.M
	v9110 = m.ExcPending
	if v9110 != 0 {
		goto L1
	} else {
		goto L1251
	}
L1251:
	;
	v9111 = F_lappend(m, v9071, v9109)
	mBase = m.M
	v9112 = m.ExcPending
	if v9112 != 0 {
		goto L1
	} else {
		goto L1252
	}
L1252:
	;
	v9113 = v9111
	goto L1244
L1253:
	;
	v9118 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+4))
	v9122 = int32(0)
	v9124 = F_makeVar(m, v9118, int32(-6), int32(26), int32(-1), v9122, v9122)
	mBase = m.M
	v9125 = m.ExcPending
	if v9125 != 0 {
		goto L1
	} else {
		goto L1254
	}
L1254:
	;
	v9126 = *(*int32)(unsafe.Add(mBase, uint32(v9031)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8266))) = v9126
	v9132 = F_pg_snprintf(m, v8266+int32(48), int32(32), int32(_a_F_subquery_planner_37), v8266)
	mBase = m.M
	v9133 = m.ExcPending
	if v9133 != 0 {
		goto L1
	} else {
		goto L1255
	}
L1255:
	;
	if v9113 != 0 {
		goto L1256
	} else {
		goto L1257
	}
L1256:
	;
	v9134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9113)+4)))
	v9138 = v9134 + int32(1)
	goto L1258
L1257:
	;
	v9138 = int32(1)
	goto L1258
L1258:
	;
	v9142 = F_pstrdup(m, v8266+int32(48))
	mBase = m.M
	v9143 = m.ExcPending
	if v9143 != 0 {
		goto L1
	} else {
		goto L1259
	}
L1259:
	;
	v9145 = F_makeTargetEntry(m, v9124, base.I32_extend16_s(v9138), v9142, int32(1))
	mBase = m.M
	v9146 = m.ExcPending
	if v9146 != 0 {
		goto L1
	} else {
		goto L1260
	}
L1260:
	;
	v9147 = F_lappend(m, v9113, v9145)
	mBase = m.M
	v9148 = m.ExcPending
	if v9148 != 0 {
		goto L1
	} else {
		goto L1261
	}
L1261:
	;
	v9149 = v9147
	goto L1229
L1262:
	;
	goto L1228
L1263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+284)) = v9328
	if v8286 != 0 {
		goto L1288
	} else {
		goto L1289
	}
L1264:
	;
	v9200 = *(*int32)(unsafe.Add(mBase, uint32(v8268)+52))
	if v9200 == int32(0) {
		v9328 = v9156
		goto L1263
	} else {
		goto L1265
	}
L1265:
	;
	v9203 = *(*int32)(unsafe.Add(mBase, uint32(v9200)+4))
	if v9203 < int32(2) {
		v9328 = v9156
		goto L1263
	} else {
		goto L1266
	}
L1266:
	;
	v9207 = F_pull_var_clause(m, v9197, int32(26))
	mBase = m.M
	v9208 = m.ExcPending
	if v9208 != 0 {
		goto L1
	} else {
		goto L1268
	}
L1267:
	;
	F_list_free(m, v9207)
	mBase = m.M
	v9327 = m.ExcPending
	if v9327 != 0 {
		goto L1
	} else {
		goto L1287
	}
L1268:
	;
	if v9207 == int32(0) {
		v9285 = v9156
		goto L1267
	} else {
		goto L1269
	}
L1269:
	;
	v9211 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+4))
	if v9211 <= int32(0) {
		v9285 = v9156
		goto L1267
	} else {
		goto L1270
	}
L1270:
	;
	v9215 = v9156
	v9216 = int32(0)
	goto L1271
L1271:
	;
	v9256 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+12))
	v9260 = *(*int32)(unsafe.Add(mBase, uint32(v9256+v9216<<(uint(int32(2))%32))))
	v9261 = *(*int32)(unsafe.Add(mBase, uint32(v9260)))
	if v9261 == int32(6) {
		goto L1274
	} else {
		goto L1275
	}
L1272:
	;
	v9285 = v9280
	goto L1267
L1273:
	;
	v9282 = v9216 + int32(1)
	v9283 = *(*int32)(unsafe.Add(mBase, uint32(v9207)+4))
	if v9282 < v9283 {
		v9215 = v9280
		v9216 = v9282
		goto L1271
	} else {
		goto L1286
	}
L1274:
	;
	v9264 = *(*int32)(unsafe.Add(mBase, uint32(v9260)+4))
	if v9264 == v8271 {
		v9280 = v9215
		goto L1273
	} else {
		goto L1277
	}
L1275:
	;
	goto L1276
L1276:
	;
	v9266 = F_tlist_member(m, v9260, v9215)
	mBase = m.M
	v9267 = m.ExcPending
	if v9267 != 0 {
		goto L1
	} else {
		goto L1278
	}
L1277:
	;
	goto L1276
L1278:
	;
	if v9266 != 0 {
		v9280 = v9215
		goto L1273
	} else {
		goto L1279
	}
L1279:
	;
	if v9215 != 0 {
		goto L1281
	} else {
		goto L1282
	}
L1280:
	;
	v9276 = F_makeTargetEntry(m, v9260, base.I32_extend16_s(v9272), int32(0), int32(1))
	mBase = m.M
	v9277 = m.ExcPending
	if v9277 != 0 {
		goto L1
	} else {
		goto L1284
	}
L1281:
	;
	v9268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9215)+4)))
	v9272 = v9268 + int32(1)
	goto L1280
L1282:
	;
	goto L1283
L1283:
	;
	v9272 = int32(1)
	goto L1280
L1284:
	;
	v9278 = F_lappend(m, v9215, v9276)
	mBase = m.M
	v9279 = m.ExcPending
	if v9279 != 0 {
		goto L1
	} else {
		goto L1285
	}
L1285:
	;
	v9280 = v9278
	goto L1273
L1286:
	;
	goto L1272
L1287:
	;
	v9328 = v9285
	goto L1263
L1288:
	;
	F_relation_close(m, v8286, int32(0))
	mBase = m.M
	v9372 = m.ExcPending
	if v9372 != 0 {
		goto L1
	} else {
		goto L1291
	}
L1289:
	;
	goto L1290
L1290:
	;
	m.G0 = v8266 + int32(80)
	v9376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+36)))
	if v9376 == int32(1) {
		goto L1292
	} else {
		goto L1293
	}
L1291:
	;
	goto L1290
L1292:
	;
	v9379 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	F_preprocess_aggrefs(m, v8227, v9379)
	mBase = m.M
	v9381 = m.ExcPending
	if v9381 != 0 {
		goto L1
	} else {
		goto L1295
	}
L1293:
	;
	goto L1294
L1294:
	;
	v9385 = int32(0)
	v9387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+37)))
	if v9387 != int32(1) {
		v10510 = v9385
		v10514 = v9385
		goto L1297
	} else {
		goto L1298
	}
L1295:
	;
	v9382 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+112))
	F_preprocess_aggrefs(m, v8227, v9382)
	mBase = m.M
	v9384 = m.ExcPending
	if v9384 != 0 {
		goto L1
	} else {
		goto L1296
	}
L1296:
	;
	goto L1294
L1297:
	;
	v10533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+36)))
	if v10533 == int32(1) {
		goto L1416
	} else {
		goto L1417
	}
L1298:
	;
	v9390 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	v9391 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+116))
	if v9391 != 0 {
		goto L1299
	} else {
		goto L1300
	}
L1299:
	;
	v9392 = *(*int32)(unsafe.Add(mBase, uint32(v9391)+4))
	v9394 = v9392
	goto L1301
L1300:
	;
	v9394 = int32(0)
	goto L1301
L1301:
	;
	v9396 = F_palloc(m, int32(12))
	mBase = m.M
	v9397 = m.ExcPending
	if v9397 != 0 {
		goto L1
	} else {
		goto L1302
	}
L1302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9396)+4)) = v9394
	*(*int32)(unsafe.Add(mBase, uint32(v9396))) = int32(0)
	v9405 = F_palloc0(m, v9394<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v9406 = m.ExcPending
	if v9406 != 0 {
		goto L1
	} else {
		goto L1303
	}
L1303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9396)+8)) = v9405
	v9408 = F_find_window_functions_walker(m, v9390, v9396)
	mBase = m.M
	v9409 = m.ExcPending
	if v9409 != 0 {
		goto L1
	} else {
		goto L1304
	}
L1304:
	;
	v9410 = *(*int32)(unsafe.Add(mBase, uint32(v9396)))
	if int32(0) < v9410 {
		goto L1307
	} else {
		goto L1308
	}
L1305:
	;
	F_pfree(m, v10459)
	mBase = m.M
	v10491 = m.ExcPending
	if v10491 != 0 {
		goto L1
	} else {
		goto L1415
	}
L1306:
	;
	F_pg_qsort(m, v10413, int32(0), int32(8), int32(880))
	mBase = m.M
	v10448 = m.ExcPending
	if v10448 != 0 {
		goto L1
	} else {
		goto L1414
	}
L1307:
	;
	v9413 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v9414 = *(*int32)(unsafe.Add(mBase, uint32(v9413)+116))
	if v9414 != 0 {
		goto L1311
	} else {
		goto L1312
	}
L1308:
	;
	goto L1309
L1309:
	;
	v10401 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8232)+37)) = uint8(v10401)
	v10510 = v9385
	v10514 = v9396
	goto L1297
L1310:
	;
	v9965 = *(*int32)(unsafe.Add(mBase, uint32(v9932)+4))
	v9966 = F_palloc_mul(m, int32(8), v9965)
	mBase = m.M
	v9967 = m.ExcPending
	if v9967 != 0 {
		goto L1
	} else {
		goto L1364
	}
L1311:
	;
	v9415 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+4))
	if v9415 <= int32(0) {
		v9932 = v9414
		goto L1310
	} else {
		goto L1314
	}
L1312:
	;
	goto L1313
L1313:
	;
	v9921 = F_palloc_mul(m, int32(8), int32(0))
	mBase = m.M
	v9922 = m.ExcPending
	if v9922 != 0 {
		goto L1
	} else {
		goto L1363
	}
L1314:
	;
	v9421 = base.I64_extend_i32_u(v8230 + int32(208))
	v9424 = int32(0)
	goto L1315
L1315:
	;
	v9463 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v9464 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+12))
	v9465 = int32(2)
	v9468 = *(*int32)(unsafe.Add(mBase, uint32(v9464+v9424<<(uint(v9465)%32))))
	v9469 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+48))
	v9473 = *(*int32)(unsafe.Add(mBase, uint32(v9463+v9469<<(uint(v9465)%32))))
	if v9473 == int32(0) {
		goto L1317
	} else {
		goto L1318
	}
L1316:
	;
	v9876 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v9877 = *(*int32)(unsafe.Add(mBase, uint32(v9876)+116))
	if v9877 != 0 {
		v9932 = v9877
		goto L1310
	} else {
		goto L1362
	}
L1317:
	;
	v9873 = v9424 + int32(1)
	v9874 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+4))
	if v9873 < v9874 {
		v9424 = v9873
		goto L1315
	} else {
		goto L1361
	}
L1318:
	;
	v9476 = *(*int32)(unsafe.Add(mBase, uint32(v9473)+4))
	if v9476 <= int32(0) {
		goto L1320
	} else {
		goto L1321
	}
L1319:
	;
	v9614 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+20))
	if v9614 == v9583 {
		goto L1317
	} else {
		goto L1336
	}
L1320:
	;
	v9583 = int32(0)
	goto L1319
L1321:
	;
	goto L1322
L1322:
	;
	v9480 = *(*int32)(unsafe.Add(mBase, uint32(v9473)+12))
	v9481 = *(*int32)(unsafe.Add(mBase, uint32(v9480)))
	v9482 = *(*int32)(unsafe.Add(mBase, uint32(v9481)+4))
	v9483 = F_get_func_support(m, v9482)
	mBase = m.M
	v9484 = m.ExcPending
	if v9484 != 0 {
		goto L1
	} else {
		goto L1323
	}
L1323:
	;
	if v9483 == int32(0) {
		goto L1317
	} else {
		goto L1324
	}
L1324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+208)) = int32(471)
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+212)) = v9481
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+216)) = v9468
	v9491 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+220)) = v9491
	v9494 = F_OidFunctionCall1Coll(m, v9483, int32(0), v9421)
	mBase = m.M
	v9495 = m.ExcPending
	if v9495 != 0 {
		goto L1
	} else {
		goto L1325
	}
L1325:
	;
	v9496 = base.I32_wrap_i64(v9494)
	if v9496 == int32(0) {
		goto L1317
	} else {
		goto L1326
	}
L1326:
	;
	v9499 = *(*int32)(unsafe.Add(mBase, uint32(v9496)+12))
	v9501 = *(*int32)(unsafe.Add(mBase, uint32(v9473)+4))
	if v9501 < int32(2) {
		v9583 = v9499
		goto L1319
	} else {
		goto L1327
	}
L1327:
	;
	v9504 = int32(1)
	goto L1328
L1328:
	;
	v9545 = *(*int32)(unsafe.Add(mBase, uint32(v9473)+12))
	v9549 = *(*int32)(unsafe.Add(mBase, uint32(v9545+v9504<<(uint(int32(2))%32))))
	v9550 = *(*int32)(unsafe.Add(mBase, uint32(v9549)+4))
	v9551 = F_get_func_support(m, v9550)
	mBase = m.M
	v9552 = m.ExcPending
	if v9552 != 0 {
		goto L1
	} else {
		goto L1330
	}
L1329:
	;
	v9583 = v9499
	goto L1319
L1330:
	;
	if v9551 == int32(0) {
		goto L1317
	} else {
		goto L1331
	}
L1331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+208)) = int32(471)
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+212)) = v9549
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+216)) = v9468
	v9559 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+220)) = v9559
	v9562 = F_OidFunctionCall1Coll(m, v9551, int32(0), v9421)
	mBase = m.M
	v9563 = m.ExcPending
	if v9563 != 0 {
		goto L1
	} else {
		goto L1332
	}
L1332:
	;
	v9564 = base.I32_wrap_i64(v9562)
	if v9564 == int32(0) {
		goto L1317
	} else {
		goto L1333
	}
L1333:
	;
	v9567 = *(*int32)(unsafe.Add(mBase, uint32(v9564)+12))
	if v9499 != v9567 {
		goto L1317
	} else {
		goto L1334
	}
L1334:
	;
	v9570 = v9504 + int32(1)
	v9571 = *(*int32)(unsafe.Add(mBase, uint32(v9473)+4))
	if v9570 < v9571 {
		v9504 = v9570
		goto L1328
	} else {
		goto L1335
	}
L1335:
	;
	goto L1329
L1336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9468)+20)) = v9583
	v9617 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+4))
	if v9617 < int32(2) {
		goto L1317
	} else {
		goto L1337
	}
L1337:
	;
	v9637 = int32(0)
	goto L1338
L1338:
	;
	v9662 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+12))
	v9666 = *(*int32)(unsafe.Add(mBase, uint32(v9662+v9637<<(uint(int32(2))%32))))
	if v9666 == v9468 {
		goto L1340
	} else {
		goto L1341
	}
L1339:
	;
	goto L1317
L1340:
	;
	v9828 = v9637 + int32(1)
	v9829 = *(*int32)(unsafe.Add(mBase, uint32(v9414)+4))
	if v9828 < v9829 {
		v9637 = v9828
		goto L1338
	} else {
		goto L1360
	}
L1341:
	;
	v9668 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+12))
	v9669 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+12))
	v9670 = F_equal(m, v9668, v9669)
	mBase = m.M
	v9671 = m.ExcPending
	if v9671 != 0 {
		goto L1
	} else {
		goto L1342
	}
L1342:
	;
	if v9670 == int32(0) {
		goto L1340
	} else {
		goto L1343
	}
L1343:
	;
	v9674 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+16))
	v9675 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+16))
	v9676 = F_equal(m, v9674, v9675)
	mBase = m.M
	v9677 = m.ExcPending
	if v9677 != 0 {
		goto L1
	} else {
		goto L1344
	}
L1344:
	;
	if v9676 == int32(0) {
		goto L1340
	} else {
		goto L1345
	}
L1345:
	;
	v9680 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+20))
	v9681 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+20))
	if v9680 != v9681 {
		goto L1340
	} else {
		goto L1346
	}
L1346:
	;
	v9683 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+24))
	v9684 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+24))
	v9685 = F_equal(m, v9683, v9684)
	mBase = m.M
	v9686 = m.ExcPending
	if v9686 != 0 {
		goto L1
	} else {
		goto L1347
	}
L1347:
	;
	if v9685 == int32(0) {
		goto L1340
	} else {
		goto L1348
	}
L1348:
	;
	v9689 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+28))
	v9690 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+28))
	v9691 = F_equal(m, v9689, v9690)
	mBase = m.M
	v9692 = m.ExcPending
	if v9692 != 0 {
		goto L1
	} else {
		goto L1349
	}
L1349:
	;
	if v9691 == int32(0) {
		goto L1340
	} else {
		goto L1350
	}
L1350:
	;
	v9695 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v9696 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+48))
	v9700 = *(*int32)(unsafe.Add(mBase, uint32(v9695+v9696<<(uint(int32(2))%32))))
	if v9700 == int32(0) {
		goto L1352
	} else {
		goto L1353
	}
L1351:
	;
	v9807 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+48))
	v9811 = *(*int32)(unsafe.Add(mBase, uint32(v9782+v9807<<(uint(int32(2))%32))))
	v9812 = F_list_concat(m, v9811, v9776)
	mBase = m.M
	v9813 = m.ExcPending
	if v9813 != 0 {
		goto L1
	} else {
		goto L1359
	}
L1352:
	;
	v9776 = int32(0)
	v9782 = v9695
	goto L1351
L1353:
	;
	goto L1354
L1354:
	;
	v9704 = *(*int32)(unsafe.Add(mBase, uint32(v9700)+4))
	if v9704 <= int32(0) {
		v9776 = v9700
		v9782 = v9695
		goto L1351
	} else {
		goto L1355
	}
L1355:
	;
	v9707 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+48))
	v9725 = int32(0)
	goto L1356
L1356:
	;
	v9750 = *(*int32)(unsafe.Add(mBase, uint32(v9700)+12))
	v9754 = *(*int32)(unsafe.Add(mBase, uint32(v9750+v9725<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9754)+32)) = v9707
	v9757 = v9725 + int32(1)
	v9758 = *(*int32)(unsafe.Add(mBase, uint32(v9700)+4))
	if v9757 < v9758 {
		v9725 = v9757
		goto L1356
	} else {
		goto L1358
	}
L1357:
	;
	v9760 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v9761 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+48))
	v9765 = *(*int32)(unsafe.Add(mBase, uint32(v9760+v9761<<(uint(int32(2))%32))))
	v9776 = v9765
	v9782 = v9760
	goto L1351
L1358:
	;
	goto L1357
L1359:
	;
	v9814 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v9815 = *(*int32)(unsafe.Add(mBase, uint32(v9666)+48))
	v9816 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v9814+v9815<<(uint(v9816)%32)))) = v9812
	v9820 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v9821 = *(*int32)(unsafe.Add(mBase, uint32(v9468)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9820+v9821<<(uint(v9816)%32)))) = int32(0)
	goto L1317
L1360:
	;
	goto L1339
L1361:
	;
	goto L1316
L1362:
	;
	goto L1313
L1363:
	;
	v10413 = v9921
	goto L1306
L1364:
	;
	v9968 = *(*int32)(unsafe.Add(mBase, uint32(v9932)+4))
	if v9968 <= int32(0) {
		v10413 = v9966
		goto L1306
	} else {
		goto L1365
	}
L1365:
	;
	v9971 = int32(0)
	v9973 = v9971
	v9974 = v9971
	v9975 = v9968
	goto L1366
L1366:
	;
	v10014 = *(*int32)(unsafe.Add(mBase, uint32(v9396)+8))
	v10015 = *(*int32)(unsafe.Add(mBase, uint32(v9932)+12))
	v10016 = int32(2)
	v10019 = *(*int32)(unsafe.Add(mBase, uint32(v10015+v9974<<(uint(v10016)%32))))
	v10020 = *(*int32)(unsafe.Add(mBase, uint32(v10019)+48))
	v10024 = *(*int32)(unsafe.Add(mBase, uint32(v10014+v10020<<(uint(v10016)%32))))
	if v10024 != 0 {
		goto L1368
	} else {
		goto L1369
	}
L1367:
	;
	F_pg_qsort(m, v9966, v10039, int32(8), int32(880))
	mBase = m.M
	v10047 = m.ExcPending
	if v10047 != 0 {
		goto L1
	} else {
		goto L1374
	}
L1368:
	;
	v10027 = v9966 + v9973<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v10027))) = v10019
	v10029 = *(*int32)(unsafe.Add(mBase, uint32(v10019)+12))
	v10030 = F_list_copy(m, v10029)
	mBase = m.M
	v10031 = m.ExcPending
	if v10031 != 0 {
		goto L1
	} else {
		goto L1371
	}
L1369:
	;
	v10039 = v9973
	v10040 = v9975
	goto L1370
L1370:
	;
	v10042 = v9974 + int32(1)
	if v10042 < v10040 {
		v9973 = v10039
		v9974 = v10042
		v9975 = v10040
		goto L1366
	} else {
		goto L1373
	}
L1371:
	;
	v10032 = *(*int32)(unsafe.Add(mBase, uint32(v10019)+16))
	v10033 = F_list_concat_unique(m, v10030, v10032)
	mBase = m.M
	v10034 = m.ExcPending
	if v10034 != 0 {
		goto L1
	} else {
		goto L1372
	}
L1372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10027)+4)) = v10033
	v10036 = *(*int32)(unsafe.Add(mBase, uint32(v9932)+4))
	v10039 = v9973 + int32(1)
	v10040 = v10036
	goto L1370
L1373:
	;
	goto L1367
L1374:
	;
	v10048 = int32(0)
	if v10039 <= v10048 {
		v10459 = v9966
		goto L1305
	} else {
		goto L1375
	}
L1375:
	;
	v10052 = v10048
	v10069 = v9385
	goto L1376
L1376:
	;
	v10095 = *(*int32)(unsafe.Add(mBase, uint32(v9966+v10052<<(uint(int32(3))%32))))
	v10096 = F_lappend(m, v10069, v10095)
	mBase = m.M
	v10097 = m.ExcPending
	if v10097 != 0 {
		goto L1
	} else {
		goto L1378
	}
L1377:
	;
	F_pfree(m, v9966)
	mBase = m.M
	v10102 = m.ExcPending
	if v10102 != 0 {
		goto L1
	} else {
		goto L1380
	}
L1378:
	;
	v10099 = v10052 + int32(1)
	if v10099 != v10039 {
		v10052 = v10099
		v10069 = v10096
		goto L1376
	} else {
		goto L1379
	}
L1379:
	;
	goto L1377
L1380:
	;
	if v10096 == int32(0) {
		goto L1381
	} else {
		goto L1382
	}
L1381:
	;
	v10510 = int32(0)
	v10514 = v9396
	goto L1297
L1382:
	;
	goto L1383
L1383:
	;
	v10107 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+4))
	if v10107 <= int32(0) {
		v10510 = v10096
		v10514 = v9396
		goto L1297
	} else {
		goto L1384
	}
L1384:
	;
	v10112 = v10107
	v10113 = int32(0)
	v10121 = int32(1)
	goto L1385
L1385:
	;
	v10152 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+12))
	v10156 = *(*int32)(unsafe.Add(mBase, uint32(v10152+v10113<<(uint(int32(2))%32))))
	v10157 = *(*int32)(unsafe.Add(mBase, uint32(v10156)+4))
	if v10157 == int32(0) {
		goto L1387
	} else {
		goto L1388
	}
L1386:
	;
	v10510 = v10096
	v10514 = v9396
	goto L1297
L1387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+80)) = v10121
	v10167 = F_pg_snprintf(m, v8230+int32(208), int32(16), int32(_a_F_subquery_planner_38), v8230+int32(80))
	mBase = m.M
	v10168 = m.ExcPending
	if v10168 != 0 {
		goto L1
	} else {
		goto L1390
	}
L1388:
	;
	v10358 = v10112
	v10367 = v10121
	goto L1389
L1389:
	;
	v10399 = v10113 + int32(1)
	if v10399 < v10358 {
		v10112 = v10358
		v10113 = v10399
		v10121 = v10367
		goto L1385
	} else {
		goto L1413
	}
L1390:
	;
	v10170 = v10121 + int32(1)
	v10171 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+4))
	if v10171 <= int32(0) {
		v10320 = v10170
		goto L1391
	} else {
		goto L1392
	}
L1391:
	;
	v10353 = F_pstrdup(m, v8230+int32(208))
	mBase = m.M
	v10354 = m.ExcPending
	if v10354 != 0 {
		goto L1
	} else {
		goto L1412
	}
L1392:
	;
	v10174 = v10171
	v10184 = v10170
	goto L1393
L1393:
	;
	v10215 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+12))
	v10218 = int32(0)
	goto L1395
L1394:
	;
	v10320 = v10306
	goto L1391
L1395:
	;
	v10261 = *(*int32)(unsafe.Add(mBase, uint32(v10215+v10218<<(uint(int32(2))%32))))
	v10262 = *(*int32)(unsafe.Add(mBase, uint32(v10261)+4))
	if v10262 != 0 {
		goto L1398
	} else {
		goto L1399
	}
L1396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+64)) = v10184
	v10303 = F_pg_snprintf(m, v8230+int32(208), int32(16), int32(_a_F_subquery_planner_38), v8230-int32(-64))
	mBase = m.M
	v10304 = m.ExcPending
	if v10304 != 0 {
		goto L1
	} else {
		goto L1410
	}
L1397:
	;
	goto L1396
L1398:
	;
	v10264 = v8230 + int32(208)
	v10267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10262))))
	v10270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10264))))
	if base.B2i32(v10267 == int32(0))|base.B2i32(v10267 != v10270) != 0 {
		v10288 = v10267
		v10289 = v10270
		goto L1402
	} else {
		goto L1403
	}
L1399:
	;
	goto L1400
L1400:
	;
	v10294 = v10218 + int32(1)
	if v10294 != v10174 {
		v10218 = v10294
		goto L1395
	} else {
		goto L1409
	}
L1401:
	;
	if v10288-v10289 == int32(0) {
		goto L1397
	} else {
		goto L1408
	}
L1402:
	;
	goto L1401
L1403:
	;
	v10273 = v10262
	v10274 = v10264
	goto L1404
L1404:
	;
	v10277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10274)+1)))
	v10278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10273)+1)))
	if v10278 == int32(0) {
		v10288 = v10278
		v10289 = v10277
		goto L1402
	} else {
		goto L1406
	}
L1405:
	;
	v10288 = v10278
	v10289 = v10277
	goto L1402
L1406:
	;
	v10281 = int32(1)
	if v10278 == v10277 {
		v10273 = v10273 + v10281
		v10274 = v10274 + v10281
		goto L1404
	} else {
		goto L1407
	}
L1407:
	;
	goto L1405
L1408:
	;
	goto L1400
L1409:
	;
	v10320 = v10184
	goto L1391
L1410:
	;
	v10306 = v10184 + int32(1)
	v10307 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+4))
	if int32(0) < v10307 {
		v10174 = v10307
		v10184 = v10306
		goto L1393
	} else {
		goto L1411
	}
L1411:
	;
	goto L1394
L1412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10156)+4)) = v10353
	v10356 = *(*int32)(unsafe.Add(mBase, uint32(v10096)+4))
	v10358 = v10356
	v10367 = v10320
	goto L1389
L1413:
	;
	goto L1386
L1414:
	;
	v10459 = v10413
	goto L1305
L1415:
	;
	v10510 = v9385
	v10514 = v9396
	goto L1297
L1416:
	;
	v10536 = int32(0)
	v10538 = float64(0)
	v10539 = m.G0
	v10541 = v10539 - int32(16)
	m.G0 = v10541
	v10543 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v10544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10543)+36)))
	if v10544 != int32(1) {
		goto L1421
	} else {
		goto L1422
	}
L1417:
	;
	goto L1418
L1418:
	;
	v11349 = float64(-1)
	v11350 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+100))
	if v11350 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1523
	}
L1419:
	;
	goto L1418
L1420:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11297 = m.ExcPending
	if v11297 != 0 {
		goto L1
	} else {
		goto L1519
	}
L1421:
	;
	m.G0 = v10541 + int32(16)
	goto L1419
L1422:
	;
	v10547 = *(*int32)(unsafe.Add(mBase, uint32(v10543)+100))
	if v10547 != 0 {
		goto L1421
	} else {
		goto L1423
	}
L1423:
	;
	v10548 = *(*int32)(unsafe.Add(mBase, uint32(v10543)+108))
	if v10548 != 0 {
		goto L1424
	} else {
		goto L1425
	}
L1424:
	;
	v10549 = *(*int32)(unsafe.Add(mBase, uint32(v10548)+4))
	if int32(1) < v10549 {
		goto L1421
	} else {
		goto L1427
	}
L1425:
	;
	goto L1426
L1426:
	;
	v10552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10543)+37)))
	if v10552 != 0 {
		goto L1421
	} else {
		goto L1428
	}
L1427:
	;
	goto L1426
L1428:
	;
	v10553 = *(*int32)(unsafe.Add(mBase, uint32(v10543)+48))
	if v10553 != 0 {
		goto L1421
	} else {
		goto L1429
	}
L1429:
	;
	v10556 = v10543 + int32(60)
	goto L1433
L1430:
	;
	v10631 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10541)+12)) = v10631
	v10634 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+344))
	if v10634 == v10631 {
		v10837 = int32(1)
		goto L1444
	} else {
		goto L1445
	}
L1431:
	;
	v10628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10626)+20)))
	if v10628 != int32(1) {
		goto L1421
	} else {
		goto L1443
	}
L1432:
	;
	v10626 = *(*int32)(unsafe.Add(mBase, uint32(v10625)))
	v10627 = *(*int32)(unsafe.Add(mBase, uint32(v10626)+12))
	switch v10627 {
	case 0:
		goto L1430
	case 1:
		goto L1431
	default:
		goto L1421
	}
L1433:
	;
	v10597 = *(*int32)(unsafe.Add(mBase, uint32(v10556)))
	v10598 = *(*int32)(unsafe.Add(mBase, uint32(v10597)))
	if v10598 != int32(65) {
		goto L1436
	} else {
		goto L1437
	}
L1434:
	;
	v10617 = *(*int32)(unsafe.Add(mBase, uint32(v10543)+52))
	v10618 = *(*int32)(unsafe.Add(mBase, uint32(v10617)+12))
	v10619 = *(*int32)(unsafe.Add(mBase, uint32(v10597)+4))
	v10625 = v10618 + v10619<<(uint(int32(2))%32) - int32(4)
	goto L1432
L1435:
	;
	goto L1434
L1436:
	;
	if v10598 != int32(63) {
		goto L1421
	} else {
		goto L1439
	}
L1437:
	;
	goto L1438
L1438:
	;
	v10610 = *(*int32)(unsafe.Add(mBase, uint32(v10597)+4))
	if v10610 == int32(0) {
		goto L1421
	} else {
		goto L1441
	}
L1439:
	;
	v10603 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+44))
	if v10603 == int32(0) {
		goto L1435
	} else {
		goto L1440
	}
L1440:
	;
	v10606 = *(*int32)(unsafe.Add(mBase, uint32(v10597)+4))
	v10625 = v10603 + v10606<<(uint(int32(2))%32)
	goto L1432
L1441:
	;
	v10613 = *(*int32)(unsafe.Add(mBase, uint32(v10610)+4))
	if v10613 != int32(1) {
		goto L1421
	} else {
		goto L1442
	}
L1442:
	;
	v10616 = *(*int32)(unsafe.Add(mBase, uint32(v10610)+12))
	v10556 = v10616
	goto L1433
L1443:
	;
	goto L1430
L1444:
	;
	if v10837 == int32(0) {
		goto L1421
	} else {
		goto L1466
	}
L1445:
	;
	v10638 = v10541 + int32(12)
	v10640 = *(*int32)(unsafe.Add(mBase, uint32(v10634)+4))
	if v10640 <= int32(0) {
		v10760 = int32(1)
		goto L1446
	} else {
		goto L1447
	}
L1446:
	;
	v10837 = v10760
	goto L1444
L1447:
	;
	v10655 = v10536
	goto L1448
L1448:
	;
	v10684 = int32(0)
	v10685 = *(*int32)(unsafe.Add(mBase, uint32(v10634)+12))
	v10689 = *(*int32)(unsafe.Add(mBase, uint32(v10685+v10655<<(uint(int32(2))%32))))
	v10690 = *(*int32)(unsafe.Add(mBase, uint32(v10689)+4))
	v10691 = *(*int32)(unsafe.Add(mBase, uint32(v10690)+12))
	v10692 = *(*int32)(unsafe.Add(mBase, uint32(v10691)))
	v10693 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+32))
	if v10693 == v10684 {
		v10837 = v10684
		goto L1444
	} else {
		goto L1450
	}
L1449:
	;
	v10760 = v10750
	goto L1446
L1450:
	;
	v10697 = *(*int32)(unsafe.Add(mBase, uint32(v10693)+4))
	if v10697 != int32(1) {
		v10837 = int32(0)
		goto L1444
	} else {
		goto L1451
	}
L1451:
	;
	v10701 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+36))
	if v10701 != 0 {
		v10837 = int32(0)
		goto L1444
	} else {
		goto L1452
	}
L1452:
	;
	v10703 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+44))
	if v10703 != 0 {
		v10837 = int32(0)
		goto L1444
	} else {
		goto L1453
	}
L1453:
	;
	v10704 = int32(0)
	v10706 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10692)+4)))
	v10707 = F_SearchSysCache1(m, v10704, v10706)
	mBase = m.M
	v10708 = m.ExcPending
	if v10708 != 0 {
		goto L1
	} else {
		goto L1454
	}
L1454:
	;
	if v10707 == int32(0) {
		v10760 = v10704
		goto L1446
	} else {
		goto L1455
	}
L1455:
	;
	v10711 = *(*int32)(unsafe.Add(mBase, uint32(v10707)+16))
	v10712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10711)+22)))
	v10714 = *(*int32)(unsafe.Add(mBase, uint32(v10711+v10712)+44))
	F_ReleaseCatCache(m, v10707)
	mBase = m.M
	v10716 = m.ExcPending
	if v10716 != 0 {
		goto L1
	} else {
		goto L1456
	}
L1456:
	;
	if v10714 == int32(0) {
		v10760 = v10704
		goto L1446
	} else {
		goto L1457
	}
L1457:
	;
	v10719 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+32))
	v10720 = *(*int32)(unsafe.Add(mBase, uint32(v10719)+12))
	v10721 = *(*int32)(unsafe.Add(mBase, uint32(v10720)))
	v10722 = *(*int32)(unsafe.Add(mBase, uint32(v10721)+4))
	v10723 = F_contain_mutable_functions(m, v10722)
	mBase = m.M
	v10724 = m.ExcPending
	if v10724 != 0 {
		goto L1
	} else {
		goto L1458
	}
L1458:
	;
	if v10723 != 0 {
		v10760 = v10704
		goto L1446
	} else {
		goto L1459
	}
L1459:
	;
	v10725 = *(*int32)(unsafe.Add(mBase, uint32(v10721)+4))
	v10726 = F_exprType(m, v10725)
	mBase = m.M
	v10727 = m.ExcPending
	if v10727 != 0 {
		goto L1
	} else {
		goto L1460
	}
L1460:
	;
	v10728 = F_type_is_rowtype(m, v10726)
	mBase = m.M
	v10729 = m.ExcPending
	if v10729 != 0 {
		goto L1
	} else {
		goto L1461
	}
L1461:
	;
	if v10728 != 0 {
		v10760 = v10704
		goto L1446
	} else {
		goto L1462
	}
L1462:
	;
	v10731 = F_palloc0(m, int32(40))
	mBase = m.M
	v10732 = m.ExcPending
	if v10732 != 0 {
		goto L1
	} else {
		goto L1463
	}
L1463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10731))) = int32(327)
	v10735 = *(*int32)(unsafe.Add(mBase, uint32(v10692)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10731)+8)) = v10714
	*(*int32)(unsafe.Add(mBase, uint32(v10731)+4)) = v10735
	v10738 = *(*int32)(unsafe.Add(mBase, uint32(v10721)+4))
	v10739 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10731)+16)) = v10739
	*(*int32)(unsafe.Add(mBase, uint32(v10731)+12)) = v10738
	*(*int64)(unsafe.Add(mBase, uint32(v10731)+24)) = v10739
	*(*int32)(unsafe.Add(mBase, uint32(v10731)+32)) = int32(0)
	v10746 = *(*int32)(unsafe.Add(mBase, uint32(v10638)))
	v10747 = F_lappend(m, v10746, v10731)
	mBase = m.M
	v10748 = m.ExcPending
	if v10748 != 0 {
		goto L1
	} else {
		goto L1464
	}
L1464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10638))) = v10747
	v10750 = int32(1)
	v10752 = v10655 + v10750
	v10753 = *(*int32)(unsafe.Add(mBase, uint32(v10634)+4))
	if v10752 < v10753 {
		v10655 = v10752
		goto L1448
	} else {
		goto L1465
	}
L1465:
	;
	goto L1449
L1466:
	;
	v10840 = *(*int32)(unsafe.Add(mBase, uint32(v10541)+12))
	if v10840 == int32(0) {
		goto L1467
	} else {
		goto L1468
	}
L1467:
	;
	v11066 = F_fetch_upper_rel(m, v8227, int32(2), int32(0))
	mBase = m.M
	v11067 = m.ExcPending
	if v11067 != 0 {
		goto L1
	} else {
		goto L1490
	}
L1468:
	;
	v10843 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if int32(0) < v10843 {
		goto L1469
	} else {
		goto L1470
	}
L1469:
	;
	v10850 = v10536
	goto L1472
L1470:
	;
	goto L1471
L1471:
	;
	v10959 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if v10959 <= int32(0) {
		goto L1467
	} else {
		goto L1483
	}
L1472:
	;
	v10887 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+12))
	v10891 = *(*int32)(unsafe.Add(mBase, uint32(v10887+v10850<<(uint(int32(2))%32))))
	v10892 = *(*int32)(unsafe.Add(mBase, uint32(v10891)+8))
	v10895 = F_get_equality_op_for_ordering_op(m, v10892, v10541+int32(11))
	mBase = m.M
	v10896 = m.ExcPending
	if v10896 != 0 {
		goto L1
	} else {
		goto L1474
	}
L1473:
	;
	goto L1471
L1474:
	;
	if v10895 == int32(0) {
		goto L1420
	} else {
		goto L1475
	}
L1475:
	;
	v10899 = *(*int32)(unsafe.Add(mBase, uint32(v10891)+8))
	v10900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10541)+11)))
	v10901 = F_build_minmax_path(m, v8227, v10891, v10895, v10899, v10900, v10900)
	mBase = m.M
	v10902 = m.ExcPending
	if v10902 != 0 {
		goto L1
	} else {
		goto L1476
	}
L1476:
	;
	if v10901 == int32(0) {
		goto L1477
	} else {
		goto L1478
	}
L1477:
	;
	v10905 = *(*int32)(unsafe.Add(mBase, uint32(v10891)+8))
	v10906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10541)+11)))
	v10909 = F_build_minmax_path(m, v8227, v10891, v10895, v10905, v10906, v10906^int32(1))
	mBase = m.M
	v10910 = m.ExcPending
	if v10910 != 0 {
		goto L1
	} else {
		goto L1480
	}
L1478:
	;
	goto L1479
L1479:
	;
	v10915 = v10850 + int32(1)
	v10916 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if v10915 < v10916 {
		v10850 = v10915
		goto L1472
	} else {
		goto L1482
	}
L1480:
	;
	if v10909 == int32(0) {
		goto L1421
	} else {
		goto L1481
	}
L1481:
	;
	goto L1479
L1482:
	;
	goto L1473
L1483:
	;
	v10967 = int32(0)
	goto L1484
L1484:
	;
	v11004 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+12))
	v11008 = *(*int32)(unsafe.Add(mBase, uint32(v11004+v10967<<(uint(int32(2))%32))))
	v11009 = *(*int32)(unsafe.Add(mBase, uint32(v11008)+12))
	v11010 = F_exprType(m, v11009)
	mBase = m.M
	v11011 = m.ExcPending
	if v11011 != 0 {
		goto L1
	} else {
		goto L1486
	}
L1485:
	;
	goto L1467
L1486:
	;
	v11013 = *(*int32)(unsafe.Add(mBase, uint32(v11008)+12))
	v11014 = F_exprCollation(m, v11013)
	mBase = m.M
	v11015 = m.ExcPending
	if v11015 != 0 {
		goto L1
	} else {
		goto L1487
	}
L1487:
	;
	v11016 = F_generate_new_exec_param(m, v8227, v11010, int32(-1), v11014)
	mBase = m.M
	v11017 = m.ExcPending
	if v11017 != 0 {
		goto L1
	} else {
		goto L1488
	}
L1488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11008)+32)) = v11016
	v11020 = v10967 + int32(1)
	v11021 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if v11020 < v11021 {
		v10967 = v11020
		goto L1484
	} else {
		goto L1489
	}
L1489:
	;
	goto L1485
L1490:
	;
	v11068 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	v11069 = F_make_pathtarget_from_tlist(m, v11068)
	mBase = m.M
	v11070 = m.ExcPending
	if v11070 != 0 {
		goto L1
	} else {
		goto L1491
	}
L1491:
	;
	v11071 = F_set_pathtarget_cost_width(m, v8227, v11069)
	mBase = m.M
	v11072 = m.ExcPending
	if v11072 != 0 {
		goto L1
	} else {
		goto L1492
	}
L1492:
	;
	v11073 = *(*int32)(unsafe.Add(mBase, uint32(v10543)+112))
	v11074 = int32(0)
	v11075 = m.G0
	v11077 = v11075 - int32(16)
	m.G0 = v11077
	v11080 = F_palloc0(m, int32(80))
	mBase = m.M
	v11081 = m.ExcPending
	if v11081 != 0 {
		goto L1
	} else {
		goto L1493
	}
L1493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+76)) = v11073
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+72)) = v10840
	v11084 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+64)) = v11084
	*(*int64)(unsafe.Add(mBase, uint32(v11080)+32)) = int64(4607182418800017408)
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+24)) = v11084
	v11090 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v11080)+20)) = uint16(v11090)
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+16)) = v11084
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+12)) = v11071
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+8)) = v11066
	*(*int64)(unsafe.Add(mBase, uint32(v11080))) = int64(1438814044473)
	if v10840 == v11084 {
		goto L1495
	} else {
		goto L1496
	}
L1494:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11080)+40)) = v11178
	v11210 = *(*float64)(unsafe.Add(mBase, uint32(v11071)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v11080)+48)) = base.F64_add(v11174, v11210)
	v11213 = *(*float64)(unsafe.Add(mBase, uint32(v11071)+16))
	v11215 = *(*float64)(unsafe.Add(mBase, uint32(v11071)+24))
	v11218 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v11080)+56)) = base.F64_add(base.F64_add(base.F64_add(v11174, v11213), v11215), v11218)
	if v11073 != 0 {
		goto L1507
	} else {
		goto L1508
	}
L1495:
	;
	v11173 = int32(1)
	v11174 = v10538
	v11178 = v11074
	goto L1494
L1496:
	;
	goto L1497
L1497:
	;
	v11101 = int32(1)
	v11102 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if v11102 <= int32(0) {
		v11173 = v11101
		v11174 = v10538
		v11178 = v11074
		goto L1494
	} else {
		goto L1498
	}
L1498:
	;
	v11106 = int32(0)
	v11111 = v11101
	v11112 = v10538
	v11116 = v11074
	goto L1499
L1499:
	;
	v11147 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+12))
	v11151 = *(*int32)(unsafe.Add(mBase, uint32(v11147+v11106<<(uint(int32(2))%32))))
	v11152 = *(*float64)(unsafe.Add(mBase, uint32(v11151)+24))
	v11153 = *(*int32)(unsafe.Add(mBase, uint32(v11151)+20))
	v11154 = *(*int32)(unsafe.Add(mBase, uint32(v11153)+40))
	v11155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11153)+21)))
	if v11155 == int32(0) {
		goto L1501
	} else {
		goto L1502
	}
L1500:
	;
	v11173 = v11161
	v11174 = v11162
	v11178 = v11163
	goto L1494
L1501:
	;
	v11158 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11080)+21)) = uint8(v11158)
	v11161 = v11158
	goto L1503
L1502:
	;
	v11161 = v11111
	goto L1503
L1503:
	;
	v11162 = base.F64_add(v11112, v11152)
	v11163 = v11116 + v11154
	v11165 = v11106 + int32(1)
	v11166 = *(*int32)(unsafe.Add(mBase, uint32(v10840)+4))
	if v11165 < v11166 {
		v11106 = v11165
		v11111 = v11161
		v11112 = v11162
		v11116 = v11163
		goto L1499
	} else {
		goto L1504
	}
L1504:
	;
	goto L1500
L1505:
	;
	m.G0 = v11077 + int32(16)
	F_add_path(m, v11066, v11080)
	mBase = m.M
	v11249 = m.ExcPending
	if v11249 != 0 {
		goto L1
	} else {
		goto L1518
	}
L1506:
	;
	v11236 = *(*int32)(unsafe.Add(mBase, uint32(v11071)+4))
	v11237 = F_is_parallel_safe(m, v8227, v11236)
	mBase = m.M
	v11238 = m.ExcPending
	if v11238 != 0 {
		goto L1
	} else {
		goto L1513
	}
L1507:
	;
	F_cost_qual_eval(m, v11077, v11073, v8227)
	mBase = m.M
	v11222 = m.ExcPending
	if v11222 != 0 {
		goto L1
	} else {
		goto L1510
	}
L1508:
	;
	goto L1509
L1509:
	;
	if v11173 == int32(0) {
		goto L1505
	} else {
		goto L1512
	}
L1510:
	;
	v11223 = *(*float64)(unsafe.Add(mBase, uint32(v11077)))
	v11224 = *(*float64)(unsafe.Add(mBase, uint32(v11080)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v11080)+48)) = base.F64_add(v11223, v11224)
	v11227 = *(*float64)(unsafe.Add(mBase, uint32(v11080)+56))
	v11228 = *(*float64)(unsafe.Add(mBase, uint32(v11077)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11080)+56)) = base.F64_add(v11227, base.F64_add(v11223, v11228))
	v11232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11080)+21)))
	if v11232 != 0 {
		goto L1506
	} else {
		goto L1511
	}
L1511:
	;
	goto L1505
L1512:
	;
	goto L1506
L1513:
	;
	if v11237 != 0 {
		goto L1514
	} else {
		goto L1515
	}
L1514:
	;
	v11239 = F_is_parallel_safe(m, v8227, v11073)
	mBase = m.M
	v11240 = m.ExcPending
	if v11240 != 0 {
		goto L1
	} else {
		goto L1517
	}
L1515:
	;
	v11242 = int32(0)
	goto L1516
L1516:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11080)+21)) = uint8(v11242)
	goto L1505
L1517:
	;
	v11242 = v11239
	goto L1516
L1518:
	;
	goto L1421
L1519:
	;
	v11298 = *(*int32)(unsafe.Add(mBase, uint32(v10891)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10541))) = v11298
	F_errmsg_internal(m, int32(_a_F_subquery_planner_39), v10541)
	mBase = m.M
	v11302 = m.ExcPending
	if v11302 != 0 {
		goto L1
	} else {
		goto L1520
	}
L1520:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_40), int32(167), int32(_a_F_subquery_planner_41))
	mBase = m.M
	v11307 = m.ExcPending
	if v11307 != 0 {
		goto L1
	} else {
		goto L1521
	}
L1521:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1522:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v8227)+320)) = v11359
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+172)) = v8239
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+168)) = v8236
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+164)) = v10510
	v11367 = F_query_planner(m, v8227, int32(881), v8230+int32(164))
	mBase = m.M
	v11368 = m.ExcPending
	if v11368 != 0 {
		goto L1
	} else {
		goto L1532
	}
L1523:
	;
	v11351 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+108))
	if v11351 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1524
	}
L1524:
	;
	v11352 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+120))
	if v11352 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1525
	}
L1525:
	;
	v11353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+36)))
	if v11353 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1526
	}
L1526:
	;
	v11354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+37)))
	if v11354 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1527
	}
L1527:
	;
	v11355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+38)))
	if v11355 != 0 {
		v11359 = v11349
		goto L1522
	} else {
		goto L1528
	}
L1528:
	;
	v11357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8227)+334)))
	if v11357 != 0 {
		goto L1529
	} else {
		goto L1530
	}
L1529:
	;
	v11358 = float64(-1)
	goto L1531
L1530:
	;
	v11358 = v8253
	goto L1531
L1531:
	;
	v11359 = v11358
	goto L1522
L1532:
	;
	v11369 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	v11370 = F_make_pathtarget_from_tlist(m, v11369)
	mBase = m.M
	v11371 = m.ExcPending
	if v11371 != 0 {
		goto L1
	} else {
		goto L1533
	}
L1533:
	;
	v11372 = F_set_pathtarget_cost_width(m, v8227, v11370)
	mBase = m.M
	v11373 = m.ExcPending
	if v11373 != 0 {
		goto L1
	} else {
		goto L1534
	}
L1534:
	;
	v11374 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	v11375 = F_is_parallel_safe(m, v8227, v11374)
	mBase = m.M
	v11376 = m.ExcPending
	if v11376 != 0 {
		goto L1
	} else {
		goto L1535
	}
L1535:
	;
	v11377 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+124))
	if v11377 != 0 {
		goto L1536
	} else {
		goto L1537
	}
L1536:
	;
	v11378 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v11379 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	if v11379 != 0 {
		goto L1539
	} else {
		goto L1540
	}
L1537:
	;
	v11703 = v11372
	v11704 = v8253
	v11717 = v11375
	goto L1538
L1538:
	;
	if v10510 != 0 {
		goto L1606
	} else {
		goto L1607
	}
L1539:
	;
	v11380 = *(*int32)(unsafe.Add(mBase, uint32(v11379)+4))
	v11382 = v11380
	goto L1541
L1540:
	;
	v11382 = int32(0)
	goto L1541
L1541:
	;
	v11383 = F_palloc0(m, v11382)
	mBase = m.M
	v11384 = m.ExcPending
	if v11384 != 0 {
		goto L1
	} else {
		goto L1542
	}
L1542:
	;
	v11385 = F_palloc0(m, v11382)
	mBase = m.M
	v11386 = m.ExcPending
	if v11386 != 0 {
		goto L1
	} else {
		goto L1543
	}
L1543:
	;
	v11387 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	if v11387 == int32(0) {
		v11659 = v8253
		v11694 = v11372
		goto L1544
	} else {
		goto L1545
	}
L1544:
	;
	v11695 = *(*int32)(unsafe.Add(mBase, uint32(v11694)+4))
	v11696 = F_is_parallel_safe(m, v8227, v11695)
	mBase = m.M
	v11697 = m.ExcPending
	if v11697 != 0 {
		goto L1
	} else {
		goto L1605
	}
L1545:
	;
	v11390 = *(*int32)(unsafe.Add(mBase, uint32(v11387)+4))
	if v11390 <= int32(0) {
		v11659 = v8253
		v11694 = v11372
		goto L1544
	} else {
		goto L1546
	}
L1546:
	;
	v11393 = int32(0)
	v11399 = v11393
	v11405 = v11393
	v11408 = v11393
	v11412 = v11393
	v11413 = v11393
	goto L1547
L1547:
	;
	v11440 = v11399 << (uint(int32(2)) % 32)
	v11441 = *(*int32)(unsafe.Add(mBase, uint32(v11387)+12))
	v11443 = *(*int32)(unsafe.Add(mBase, uint32(v11440+v11441)))
	v11444 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+8))
	if v11444 != 0 {
		goto L1551
	} else {
		goto L1552
	}
L1548:
	;
	v11497 = int32(1)
	v11499 = v11492 & (v11490 ^ v11497)
	if (v11499|v11491)&v11497 != 0 {
		goto L1571
	} else {
		goto L1572
	}
L1549:
	;
	v11494 = v11399 + int32(1)
	v11495 = *(*int32)(unsafe.Add(mBase, uint32(v11387)+4))
	if v11494 < v11495 {
		v11399 = v11494
		v11405 = v11489
		v11408 = v11490
		v11412 = v11491
		v11413 = v11492
		goto L1547
	} else {
		goto L1570
	}
L1550:
	;
	if v11408&int32(1) != 0 {
		goto L1565
	} else {
		goto L1566
	}
L1551:
	;
	v11446 = *(*int32)(unsafe.Add(mBase, uint32(v11440+v11444)))
	if v11446 != 0 {
		goto L1550
	} else {
		goto L1554
	}
L1552:
	;
	goto L1553
L1553:
	;
	v11447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11378)+38)))
	if v11447 != int32(1) {
		goto L1555
	} else {
		goto L1556
	}
L1554:
	;
	goto L1553
L1555:
	;
	v11458 = F_contain_volatile_functions(m, v11443)
	mBase = m.M
	v11459 = m.ExcPending
	if v11459 != 0 {
		goto L1
	} else {
		goto L1559
	}
L1556:
	;
	v11450 = F_expression_returns_set(m, v11443)
	mBase = m.M
	v11451 = m.ExcPending
	if v11451 != 0 {
		goto L1
	} else {
		goto L1557
	}
L1557:
	;
	if v11450 == int32(0) {
		goto L1555
	} else {
		goto L1558
	}
L1558:
	;
	v11454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11399+v11383))) = uint8(v11454)
	v11489 = v11405
	v11490 = v11408
	v11491 = v11412
	v11492 = v11454
	goto L1549
L1559:
	;
	if v11458 != 0 {
		goto L1560
	} else {
		goto L1561
	}
L1560:
	;
	v11460 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11399+v11385))) = uint8(v11460)
	v11489 = v11405
	v11490 = v11408
	v11491 = v11460
	v11492 = v11413
	goto L1549
L1561:
	;
	goto L1562
L1562:
	;
	F_cost_qual_eval_node(m, v8230+int32(208), v11443, v8227)
	mBase = m.M
	v11467 = m.ExcPending
	if v11467 != 0 {
		goto L1
	} else {
		goto L1563
	}
L1563:
	;
	v11468 = *(*float64)(unsafe.Add(mBase, uint32(v8230)+216))
	v11470 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[5]))
	if base.F64_gt(v11468, base.F64_mul(v11470, float64(10))) == int32(0) {
		v11489 = v11405
		v11490 = v11408
		v11491 = v11412
		v11492 = v11413
		goto L1549
	} else {
		goto L1564
	}
L1564:
	;
	v11476 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11399+v11385))) = uint8(v11476)
	v11489 = v11476
	v11490 = v11408
	v11491 = v11412
	v11492 = v11413
	goto L1549
L1565:
	;
	v11489 = v11405
	v11490 = int32(1)
	v11491 = v11412
	v11492 = v11413
	goto L1549
L1566:
	;
	goto L1567
L1567:
	;
	v11484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11378)+38)))
	if v11484 != int32(1) {
		v11489 = v11405
		v11490 = int32(0)
		v11491 = v11412
		v11492 = v11413
		goto L1549
	} else {
		goto L1568
	}
L1568:
	;
	v11487 = F_expression_returns_set(m, v11443)
	mBase = m.M
	v11488 = m.ExcPending
	if v11488 != 0 {
		goto L1
	} else {
		goto L1569
	}
L1569:
	;
	v11489 = v11405
	v11490 = v11487
	v11491 = v11412
	v11492 = v11413
	goto L1549
L1570:
	;
	goto L1548
L1571:
	;
	v11513 = F_create_empty_pathtarget(m)
	mBase = m.M
	v11514 = m.ExcPending
	if v11514 != 0 {
		goto L1
	} else {
		goto L1576
	}
L1572:
	;
	if v11489&int32(1) == int32(0) {
		v11659 = v8253
		v11694 = v11372
		goto L1544
	} else {
		goto L1573
	}
L1573:
	;
	v11507 = *(*int32)(unsafe.Add(mBase, uint32(v11378)+132))
	if v11507 != 0 {
		goto L1571
	} else {
		goto L1574
	}
L1574:
	;
	v11508 = *(*float64)(unsafe.Add(mBase, uint32(v8227)+312))
	if base.F64_gt(v11508, float64(0)) == int32(0) {
		v11659 = v8253
		v11694 = v11372
		goto L1544
	} else {
		goto L1575
	}
L1575:
	;
	goto L1571
L1576:
	;
	v11515 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	if v11515 == int32(0) {
		goto L1578
	} else {
		goto L1579
	}
L1577:
	;
	v11639 = F_pull_var_clause(m, v11597, int32(21))
	mBase = m.M
	v11640 = m.ExcPending
	if v11640 != 0 {
		goto L1
	} else {
		goto L1597
	}
L1578:
	;
	v11597 = int32(0)
	goto L1577
L1579:
	;
	goto L1580
L1580:
	;
	v11519 = int32(0)
	v11520 = *(*int32)(unsafe.Add(mBase, uint32(v11515)+4))
	if v11520 <= v11519 {
		v11597 = v11519
		goto L1577
	} else {
		goto L1581
	}
L1581:
	;
	v11524 = v11519
	v11525 = int32(0)
	goto L1582
L1582:
	;
	v11566 = v11525 << (uint(int32(2)) % 32)
	v11567 = *(*int32)(unsafe.Add(mBase, uint32(v11515)+12))
	v11569 = *(*int32)(unsafe.Add(mBase, uint32(v11566+v11567)))
	v11571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11525+v11385))))
	if v11571 == int32(0) {
		goto L1586
	} else {
		goto L1587
	}
L1583:
	;
	v11597 = v11591
	goto L1577
L1584:
	;
	v11594 = v11525 + int32(1)
	v11595 = *(*int32)(unsafe.Add(mBase, uint32(v11515)+4))
	if v11594 < v11595 {
		v11524 = v11591
		v11525 = v11594
		goto L1582
	} else {
		goto L1596
	}
L1585:
	;
	v11584 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+8))
	if v11584 != 0 {
		goto L1592
	} else {
		goto L1593
	}
L1586:
	;
	if v11499&int32(1) == int32(0) {
		goto L1585
	} else {
		goto L1589
	}
L1587:
	;
	goto L1588
L1588:
	;
	v11582 = F_lappend(m, v11524, v11569)
	mBase = m.M
	v11583 = m.ExcPending
	if v11583 != 0 {
		goto L1
	} else {
		goto L1591
	}
L1589:
	;
	v11579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11525+v11383))))
	if v11579 != int32(1) {
		goto L1585
	} else {
		goto L1590
	}
L1590:
	;
	goto L1588
L1591:
	;
	v11591 = v11582
	goto L1584
L1592:
	;
	v11586 = *(*int32)(unsafe.Add(mBase, uint32(v11584+v11566)))
	v11588 = v11586
	goto L1594
L1593:
	;
	v11588 = int32(0)
	goto L1594
L1594:
	;
	F_add_column_to_pathtarget(m, v11513, v11569, v11588)
	mBase = m.M
	v11590 = m.ExcPending
	if v11590 != 0 {
		goto L1
	} else {
		goto L1595
	}
L1595:
	;
	v11591 = v11524
	goto L1584
L1596:
	;
	goto L1583
L1597:
	;
	F_add_new_columns_to_pathtarget(m, v11513, v11639)
	mBase = m.M
	v11642 = m.ExcPending
	if v11642 != 0 {
		goto L1
	} else {
		goto L1598
	}
L1598:
	;
	F_list_free(m, v11639)
	mBase = m.M
	v11644 = m.ExcPending
	if v11644 != 0 {
		goto L1
	} else {
		goto L1599
	}
L1599:
	;
	F_list_free(m, v11597)
	mBase = m.M
	v11646 = m.ExcPending
	if v11646 != 0 {
		goto L1
	} else {
		goto L1600
	}
L1600:
	;
	if v11499&int32(1) != 0 {
		goto L1601
	} else {
		goto L1602
	}
L1601:
	;
	v11650 = float64(-1)
	goto L1603
L1602:
	;
	v11650 = v8253
	goto L1603
L1603:
	;
	v11651 = F_set_pathtarget_cost_width(m, v8227, v11513)
	mBase = m.M
	v11652 = m.ExcPending
	if v11652 != 0 {
		goto L1
	} else {
		goto L1604
	}
L1604:
	;
	v11659 = v11650
	v11694 = v11651
	goto L1544
L1605:
	;
	v11703 = v11694
	v11704 = v11659
	v11717 = v11696
	goto L1538
L1606:
	;
	v11739 = int32(0)
	v11740 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	if v11739 < v11740 {
		goto L1609
	} else {
		goto L1610
	}
L1607:
	;
	v12284 = v11703
	v12286 = v11717
	goto L1608
L1608:
	;
	v12313 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+100))
	if v12313 != 0 {
		goto L1660
	} else {
		goto L1661
	}
L1609:
	;
	v11745 = v11739
	v11746 = int32(0)
	goto L1612
L1610:
	;
	v11997 = v11739
	goto L1611
L1611:
	;
	v12037 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+276))
	if v12037 == int32(0) {
		v12098 = v11997
		goto L1629
	} else {
		goto L1630
	}
L1612:
	;
	v11785 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+12))
	v11789 = *(*int32)(unsafe.Add(mBase, uint32(v11785+v11746<<(uint(int32(2))%32))))
	v11790 = *(*int32)(unsafe.Add(mBase, uint32(v11789)+12))
	if v11790 == int32(0) {
		v11851 = v11745
		goto L1614
	} else {
		goto L1615
	}
L1613:
	;
	v11997 = v11952
	goto L1611
L1614:
	;
	v11891 = *(*int32)(unsafe.Add(mBase, uint32(v11789)+16))
	if v11891 == int32(0) {
		v11952 = v11851
		goto L1621
	} else {
		goto L1622
	}
L1615:
	;
	v11793 = int32(0)
	v11794 = *(*int32)(unsafe.Add(mBase, uint32(v11790)+4))
	if v11794 <= v11793 {
		v11851 = v11745
		goto L1614
	} else {
		goto L1616
	}
L1616:
	;
	v11798 = v11745
	v11806 = v11793
	goto L1617
L1617:
	;
	v11838 = *(*int32)(unsafe.Add(mBase, uint32(v11790)+12))
	v11842 = *(*int32)(unsafe.Add(mBase, uint32(v11838+v11806<<(uint(int32(2))%32))))
	v11843 = *(*int32)(unsafe.Add(mBase, uint32(v11842)+4))
	v11844 = F_bms_add_member(m, v11798, v11843)
	mBase = m.M
	v11845 = m.ExcPending
	if v11845 != 0 {
		goto L1
	} else {
		goto L1619
	}
L1618:
	;
	v11851 = v11844
	goto L1614
L1619:
	;
	v11847 = v11806 + int32(1)
	v11848 = *(*int32)(unsafe.Add(mBase, uint32(v11790)+4))
	if v11847 < v11848 {
		v11798 = v11844
		v11806 = v11847
		goto L1617
	} else {
		goto L1620
	}
L1620:
	;
	goto L1618
L1621:
	;
	v11993 = v11746 + int32(1)
	v11994 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	if v11993 < v11994 {
		v11745 = v11952
		v11746 = v11993
		goto L1612
	} else {
		goto L1628
	}
L1622:
	;
	v11894 = int32(0)
	v11895 = *(*int32)(unsafe.Add(mBase, uint32(v11891)+4))
	if v11895 <= v11894 {
		v11952 = v11851
		goto L1621
	} else {
		goto L1623
	}
L1623:
	;
	v11899 = v11851
	v11907 = v11894
	goto L1624
L1624:
	;
	v11939 = *(*int32)(unsafe.Add(mBase, uint32(v11891)+12))
	v11943 = *(*int32)(unsafe.Add(mBase, uint32(v11939+v11907<<(uint(int32(2))%32))))
	v11944 = *(*int32)(unsafe.Add(mBase, uint32(v11943)+4))
	v11945 = F_bms_add_member(m, v11899, v11944)
	mBase = m.M
	v11946 = m.ExcPending
	if v11946 != 0 {
		goto L1
	} else {
		goto L1626
	}
L1625:
	;
	v11952 = v11945
	goto L1621
L1626:
	;
	v11948 = v11907 + int32(1)
	v11949 = *(*int32)(unsafe.Add(mBase, uint32(v11891)+4))
	if v11948 < v11949 {
		v11899 = v11945
		v11907 = v11948
		goto L1624
	} else {
		goto L1627
	}
L1627:
	;
	goto L1625
L1628:
	;
	goto L1613
L1629:
	;
	v12138 = F_create_empty_pathtarget(m)
	mBase = m.M
	v12139 = m.ExcPending
	if v12139 != 0 {
		goto L1
	} else {
		goto L1636
	}
L1630:
	;
	v12040 = int32(0)
	v12041 = *(*int32)(unsafe.Add(mBase, uint32(v12037)+4))
	if v12041 <= v12040 {
		v12098 = v11997
		goto L1629
	} else {
		goto L1631
	}
L1631:
	;
	v12045 = v11997
	v12053 = v12040
	goto L1632
L1632:
	;
	v12085 = *(*int32)(unsafe.Add(mBase, uint32(v12037)+12))
	v12089 = *(*int32)(unsafe.Add(mBase, uint32(v12085+v12053<<(uint(int32(2))%32))))
	v12090 = *(*int32)(unsafe.Add(mBase, uint32(v12089)+4))
	v12091 = F_bms_add_member(m, v12045, v12090)
	mBase = m.M
	v12092 = m.ExcPending
	if v12092 != 0 {
		goto L1
	} else {
		goto L1634
	}
L1633:
	;
	v12098 = v12091
	goto L1629
L1634:
	;
	v12094 = v12053 + int32(1)
	v12095 = *(*int32)(unsafe.Add(mBase, uint32(v12037)+4))
	if v12094 < v12095 {
		v12045 = v12091
		v12053 = v12094
		goto L1632
	} else {
		goto L1635
	}
L1635:
	;
	goto L1633
L1636:
	;
	v12140 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	if v12140 == int32(0) {
		goto L1638
	} else {
		goto L1639
	}
L1637:
	;
	v12259 = F_pull_var_clause(m, v12219, int32(25))
	mBase = m.M
	v12260 = m.ExcPending
	if v12260 != 0 {
		goto L1
	} else {
		goto L1653
	}
L1638:
	;
	v12219 = int32(0)
	goto L1637
L1639:
	;
	goto L1640
L1640:
	;
	v12144 = int32(0)
	v12145 = *(*int32)(unsafe.Add(mBase, uint32(v12140)+4))
	if v12145 <= v12144 {
		v12219 = v12144
		goto L1637
	} else {
		goto L1641
	}
L1641:
	;
	v12151 = v12144
	v12158 = int32(0)
	goto L1642
L1642:
	;
	v12191 = v12158 << (uint(int32(2)) % 32)
	v12192 = *(*int32)(unsafe.Add(mBase, uint32(v12140)+12))
	v12194 = *(*int32)(unsafe.Add(mBase, uint32(v12191+v12192)))
	v12195 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+8))
	if v12195 == int32(0) {
		goto L1645
	} else {
		goto L1646
	}
L1643:
	;
	v12219 = v12211
	goto L1637
L1644:
	;
	v12214 = v12158 + int32(1)
	v12215 = *(*int32)(unsafe.Add(mBase, uint32(v12140)+4))
	if v12214 < v12215 {
		v12151 = v12211
		v12158 = v12214
		goto L1642
	} else {
		goto L1652
	}
L1645:
	;
	v12209 = F_lappend(m, v12151, v12194)
	mBase = m.M
	v12210 = m.ExcPending
	if v12210 != 0 {
		goto L1
	} else {
		goto L1651
	}
L1646:
	;
	v12199 = *(*int32)(unsafe.Add(mBase, uint32(v12195+v12191)))
	if v12199 == int32(0) {
		goto L1645
	} else {
		goto L1647
	}
L1647:
	;
	v12202 = F_bms_is_member(m, v12199, v12098)
	mBase = m.M
	v12203 = m.ExcPending
	if v12203 != 0 {
		goto L1
	} else {
		goto L1648
	}
L1648:
	;
	if v12202 == int32(0) {
		goto L1645
	} else {
		goto L1649
	}
L1649:
	;
	F_add_column_to_pathtarget(m, v12138, v12194, v12199)
	mBase = m.M
	v12207 = m.ExcPending
	if v12207 != 0 {
		goto L1
	} else {
		goto L1650
	}
L1650:
	;
	v12211 = v12151
	goto L1644
L1651:
	;
	v12211 = v12209
	goto L1644
L1652:
	;
	goto L1643
L1653:
	;
	F_add_new_columns_to_pathtarget(m, v12138, v12259)
	mBase = m.M
	v12262 = m.ExcPending
	if v12262 != 0 {
		goto L1
	} else {
		goto L1654
	}
L1654:
	;
	F_list_free(m, v12259)
	mBase = m.M
	v12264 = m.ExcPending
	if v12264 != 0 {
		goto L1
	} else {
		goto L1655
	}
L1655:
	;
	F_list_free(m, v12219)
	mBase = m.M
	v12266 = m.ExcPending
	if v12266 != 0 {
		goto L1
	} else {
		goto L1656
	}
L1656:
	;
	v12267 = F_set_pathtarget_cost_width(m, v8227, v12138)
	mBase = m.M
	v12268 = m.ExcPending
	if v12268 != 0 {
		goto L1
	} else {
		goto L1657
	}
L1657:
	;
	v12269 = *(*int32)(unsafe.Add(mBase, uint32(v12267)+4))
	v12270 = F_is_parallel_safe(m, v8227, v12269)
	mBase = m.M
	v12271 = m.ExcPending
	if v12271 != 0 {
		goto L1
	} else {
		goto L1658
	}
L1658:
	;
	v12284 = v12267
	v12286 = v12270
	goto L1608
L1659:
	;
	v12571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+38)))
	if v12571 == int32(1) {
		goto L1717
	} else {
		goto L1718
	}
L1660:
	;
	v12322 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v12323 = F_create_empty_pathtarget(m)
	mBase = m.M
	v12324 = m.ExcPending
	if v12324 != 0 {
		goto L1
	} else {
		goto L1665
	}
L1661:
	;
	v12314 = *(*int32)(unsafe.Add(mBase, uint32(v8232)+108))
	if v12314 != 0 {
		goto L1660
	} else {
		goto L1662
	}
L1662:
	;
	v12315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+36)))
	if v12315 != 0 {
		goto L1660
	} else {
		goto L1663
	}
L1663:
	;
	v12317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8227)+334)))
	if v12317 != int32(1) {
		v12530 = v12284
		v12531 = int32(0)
		v12570 = v12286
		goto L1659
	} else {
		goto L1664
	}
L1664:
	;
	goto L1660
L1665:
	;
	v12325 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+4))
	if v12325 == int32(0) {
		goto L1667
	} else {
		goto L1668
	}
L1666:
	;
	v12497 = *(*int32)(unsafe.Add(mBase, uint32(v12322)+112))
	if v12497 != 0 {
		goto L1700
	} else {
		goto L1701
	}
L1667:
	;
	v12456 = int32(0)
	goto L1666
L1668:
	;
	goto L1669
L1669:
	;
	v12329 = int32(0)
	v12330 = *(*int32)(unsafe.Add(mBase, uint32(v12325)+4))
	if v12330 <= v12329 {
		v12456 = v12329
		goto L1666
	} else {
		goto L1670
	}
L1670:
	;
	v12334 = v12329
	v12335 = int32(0)
	goto L1671
L1671:
	;
	v12376 = v12335 << (uint(int32(2)) % 32)
	v12377 = *(*int32)(unsafe.Add(mBase, uint32(v12325)+12))
	v12379 = *(*int32)(unsafe.Add(mBase, uint32(v12376+v12377)))
	v12380 = *(*int32)(unsafe.Add(mBase, uint32(v11372)+8))
	if v12380 == int32(0) {
		goto L1674
	} else {
		goto L1675
	}
L1672:
	;
	v12456 = v12448
	goto L1666
L1673:
	;
	v12453 = v12335 + int32(1)
	v12454 = *(*int32)(unsafe.Add(mBase, uint32(v12325)+4))
	if v12453 < v12454 {
		v12334 = v12448
		v12335 = v12453
		goto L1671
	} else {
		goto L1699
	}
L1674:
	;
	v12446 = F_lappend(m, v12334, v12379)
	mBase = m.M
	v12447 = m.ExcPending
	if v12447 != 0 {
		goto L1
	} else {
		goto L1698
	}
L1675:
	;
	v12384 = *(*int32)(unsafe.Add(mBase, uint32(v12380+v12376)))
	if v12384 == int32(0) {
		goto L1674
	} else {
		goto L1676
	}
L1676:
	;
	v12387 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+276))
	if v12387 == int32(0) {
		goto L1674
	} else {
		goto L1677
	}
L1677:
	;
	if v12387 != 0 {
		goto L1680
	} else {
		goto L1681
	}
L1678:
	;
	if v12425 == int32(0) {
		goto L1674
	} else {
		goto L1691
	}
L1679:
	;
	goto L1678
L1680:
	;
	v12393 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+4))
	if v12393 <= int32(0) {
		v12425 = int32(0)
		goto L1679
	} else {
		goto L1683
	}
L1681:
	;
	goto L1682
L1682:
	;
	v12425 = int32(0)
	goto L1679
L1683:
	;
	v12396 = int32(0)
	if v12396 < v12393 {
		goto L1684
	} else {
		goto L1685
	}
L1684:
	;
	v12399 = v12393
	goto L1686
L1685:
	;
	v12399 = v12396
	goto L1686
L1686:
	;
	v12400 = *(*int32)(unsafe.Add(mBase, uint32(v12387)+12))
	v12403 = int32(0)
	goto L1687
L1687:
	;
	v12410 = *(*int32)(unsafe.Add(mBase, uint32(v12400+v12403<<(uint(int32(2))%32))))
	v12411 = *(*int32)(unsafe.Add(mBase, uint32(v12410)+4))
	if v12411 == v12384 {
		v12425 = v12410
		goto L1679
	} else {
		goto L1689
	}
L1688:
	;
	goto L1682
L1689:
	;
	v12414 = v12403 + int32(1)
	if v12414 != v12399 {
		v12403 = v12414
		goto L1687
	} else {
		goto L1690
	}
L1690:
	;
	goto L1688
L1691:
	;
	v12429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12322)+45)))
	if v12429 != int32(1) {
		v12441 = v12379
		goto L1692
	} else {
		goto L1693
	}
L1692:
	;
	F_add_column_to_pathtarget(m, v12323, v12441, v12384)
	mBase = m.M
	v12443 = m.ExcPending
	if v12443 != 0 {
		goto L1
	} else {
		goto L1697
	}
L1693:
	;
	v12432 = *(*int32)(unsafe.Add(mBase, uint32(v12322)+108))
	if v12432 == int32(0) {
		v12441 = v12379
		goto L1692
	} else {
		goto L1694
	}
L1694:
	;
	v12435 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+340))
	v12436 = F_bms_make_singleton(m, v12435)
	mBase = m.M
	v12437 = m.ExcPending
	if v12437 != 0 {
		goto L1
	} else {
		goto L1695
	}
L1695:
	;
	v12439 = F_remove_nulling_relids(m, v12379, v12436, int32(0))
	mBase = m.M
	v12440 = m.ExcPending
	if v12440 != 0 {
		goto L1
	} else {
		goto L1696
	}
L1696:
	;
	v12441 = v12439
	goto L1692
L1697:
	;
	v12448 = v12334
	goto L1673
L1698:
	;
	v12448 = v12446
	goto L1673
L1699:
	;
	goto L1672
L1700:
	;
	v12498 = F_lappend(m, v12456, v12497)
	mBase = m.M
	v12499 = m.ExcPending
	if v12499 != 0 {
		goto L1
	} else {
		goto L1703
	}
L1701:
	;
	v12500 = v12456
	goto L1702
L1702:
	;
	v12502 = F_pull_var_clause(m, v12500, int32(26))
	mBase = m.M
	v12503 = m.ExcPending
	if v12503 != 0 {
		goto L1
	} else {
		goto L1704
	}
L1703:
	;
	v12500 = v12498
	goto L1702
L1704:
	;
	v12504 = int32(1)
	v12505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12322)+45)))
	if v12505 != v12504 {
		v12517 = v12502
		goto L1705
	} else {
		goto L1706
	}
L1705:
	;
	F_add_new_columns_to_pathtarget(m, v12323, v12517)
	mBase = m.M
	v12519 = m.ExcPending
	if v12519 != 0 {
		goto L1
	} else {
		goto L1710
	}
L1706:
	;
	v12508 = *(*int32)(unsafe.Add(mBase, uint32(v12322)+108))
	if v12508 == int32(0) {
		v12517 = v12502
		goto L1705
	} else {
		goto L1707
	}
L1707:
	;
	v12511 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+340))
	v12512 = F_bms_make_singleton(m, v12511)
	mBase = m.M
	v12513 = m.ExcPending
	if v12513 != 0 {
		goto L1
	} else {
		goto L1708
	}
L1708:
	;
	v12515 = F_remove_nulling_relids(m, v12502, v12512, int32(0))
	mBase = m.M
	v12516 = m.ExcPending
	if v12516 != 0 {
		goto L1
	} else {
		goto L1709
	}
L1709:
	;
	v12517 = v12515
	goto L1705
L1710:
	;
	F_list_free(m, v12517)
	mBase = m.M
	v12521 = m.ExcPending
	if v12521 != 0 {
		goto L1
	} else {
		goto L1711
	}
L1711:
	;
	F_list_free(m, v12500)
	mBase = m.M
	v12523 = m.ExcPending
	if v12523 != 0 {
		goto L1
	} else {
		goto L1712
	}
L1712:
	;
	v12524 = F_set_pathtarget_cost_width(m, v8227, v12323)
	mBase = m.M
	v12525 = m.ExcPending
	if v12525 != 0 {
		goto L1
	} else {
		goto L1713
	}
L1713:
	;
	v12526 = *(*int32)(unsafe.Add(mBase, uint32(v12524)+4))
	v12527 = F_is_parallel_safe(m, v8227, v12526)
	mBase = m.M
	v12528 = m.ExcPending
	if v12528 != 0 {
		goto L1
	} else {
		goto L1714
	}
L1714:
	;
	v12530 = v12524
	v12531 = v12504
	v12570 = v12527
	goto L1659
L1715:
	;
	v12659 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+176))
	F_apply_scanjoin_target_to_paths(m, v8227, v11367, v12652, v12659, v12570, v12657)
	mBase = m.M
	v12661 = m.ExcPending
	if v12661 != 0 {
		goto L1
	} else {
		goto L1730
	}
L1716:
	;
	v12643 = *(*int32)(unsafe.Add(mBase, uint32(v12639)+4))
	if v12643 != int32(1) {
		goto L1726
	} else {
		goto L1727
	}
L1717:
	;
	F_split_pathtarget_at_srfs(m, v8227, v11372, v11703, v8230+int32(204), v8230+int32(200))
	mBase = m.M
	v12579 = m.ExcPending
	if v12579 != 0 {
		goto L1
	} else {
		goto L1720
	}
L1718:
	;
	goto L1719
L1719:
	;
	v12612 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+204)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+200)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+192)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+196)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+184)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+188)) = v12612
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+160)) = v12530
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+60)) = v12530
	v12630 = F_list_make1_impl(m, int32(1), v8230+int32(60))
	mBase = m.M
	v12631 = m.ExcPending
	if v12631 != 0 {
		goto L1
	} else {
		goto L1724
	}
L1720:
	;
	v12580 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+204))
	v12581 = *(*int32)(unsafe.Add(mBase, uint32(v12580)+12))
	v12582 = *(*int32)(unsafe.Add(mBase, uint32(v12581)))
	F_split_pathtarget_at_srfs(m, v8227, v11703, v12284, v8230+int32(196), v8230+int32(192))
	mBase = m.M
	v12588 = m.ExcPending
	if v12588 != 0 {
		goto L1
	} else {
		goto L1721
	}
L1721:
	;
	v12589 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+196))
	v12590 = *(*int32)(unsafe.Add(mBase, uint32(v12589)+12))
	v12591 = *(*int32)(unsafe.Add(mBase, uint32(v12590)))
	F_split_pathtarget_at_srfs_extended(m, v8227, v12284, v12530, v8230+int32(188), v8230+int32(184), int32(1))
	mBase = m.M
	v12598 = m.ExcPending
	if v12598 != 0 {
		goto L1
	} else {
		goto L1722
	}
L1722:
	;
	v12599 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+188))
	v12600 = *(*int32)(unsafe.Add(mBase, uint32(v12599)+12))
	v12601 = *(*int32)(unsafe.Add(mBase, uint32(v12600)))
	F_split_pathtarget_at_srfs(m, v8227, v12530, int32(0), v8230+int32(180), v8230+int32(176))
	mBase = m.M
	v12608 = m.ExcPending
	if v12608 != 0 {
		goto L1
	} else {
		goto L1723
	}
L1723:
	;
	v12609 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+180))
	v12610 = *(*int32)(unsafe.Add(mBase, uint32(v12609)+12))
	v12611 = *(*int32)(unsafe.Add(mBase, uint32(v12610)))
	v12637 = v12611
	v12638 = v12591
	v12639 = v12609
	v12640 = v12601
	v12641 = v12582
	goto L1716
L1724:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+180)) = v12630
	if v12630 != 0 {
		v12637 = v12530
		v12638 = v11703
		v12639 = v12630
		v12640 = v12284
		v12641 = v11372
		goto L1716
	} else {
		goto L1725
	}
L1725:
	;
	v12652 = v12612
	v12654 = v11703
	v12656 = v12284
	v12657 = int32(0)
	v12658 = v11372
	goto L1715
L1726:
	;
	v12652 = v12639
	v12654 = v12638
	v12656 = v12640
	v12657 = int32(0)
	v12658 = v12641
	goto L1715
L1727:
	;
	goto L1728
L1728:
	;
	v12646 = *(*int32)(unsafe.Add(mBase, uint32(v12637)+4))
	v12647 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+40))
	v12648 = *(*int32)(unsafe.Add(mBase, uint32(v12647)+4))
	v12649 = F_equal(m, v12646, v12648)
	mBase = m.M
	v12650 = m.ExcPending
	if v12650 != 0 {
		goto L1
	} else {
		goto L1729
	}
L1729:
	;
	v12651 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+180))
	v12652 = v12651
	v12654 = v12638
	v12656 = v12640
	v12657 = v12649
	v12658 = v12641
	goto L1715
L1730:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+268)) = v12658
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+272)) = v12658
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+264)) = v12654
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+260)) = v12654
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+256)) = v12654
	*(*int32)(unsafe.Add(mBase, uint32(v8227)+252)) = v12656
	if v12531 == int32(0) {
		goto L1732
	} else {
		goto L1733
	}
L1731:
	;
	if v10510 == int32(0) {
		goto L1828
	} else {
		goto L1829
	}
L1732:
	;
	v13019 = v11367
	goto L1731
L1733:
	;
	goto L1734
L1734:
	;
	v12670 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v12671 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+352)) = v12671
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+344)) = v12671
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+336)) = v12671
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+328)) = v12671
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+320)) = v12671
	F_get_agg_clause_costs(m, v8227, int32(0), v8230+int32(320))
	mBase = m.M
	v12685 = m.ExcPending
	if v12685 != 0 {
		goto L1
	} else {
		goto L1735
	}
L1735:
	;
	v12686 = *(*int32)(unsafe.Add(mBase, uint32(v12670)+112))
	v12687 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+4))
	v12694 = int32(0)
	if base.B2i32(base.Ui32(int32(5)) < base.Ui32(v12687))|base.B2i32(int32(1)<<(uint(v12687)%32)&int32(44) == v12694) == v12694 {
		goto L1737
	} else {
		goto L1738
	}
L1736:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12709)+40)) = v12656
	v12711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11367)+26)))
	if v12286&v12711 != int32(1) {
		goto L1742
	} else {
		goto L1743
	}
L1737:
	;
	v12700 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+8))
	v12701 = F_fetch_upper_rel(m, v8227, int32(2), v12700)
	mBase = m.M
	v12702 = m.ExcPending
	if v12702 != 0 {
		goto L1
	} else {
		goto L1740
	}
L1738:
	;
	goto L1739
L1739:
	;
	v12707 = F_fetch_upper_rel(m, v8227, int32(2), int32(0))
	mBase = m.M
	v12708 = m.ExcPending
	if v12708 != 0 {
		goto L1
	} else {
		goto L1741
	}
L1740:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12701)+4)) = int32(5)
	v12709 = v12701
	goto L1736
L1741:
	;
	v12709 = v12707
	goto L1736
L1742:
	;
	v12721 = *(*int64)(unsafe.Add(mBase, uint32(v11367)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v12709)+32)) = v12721
	v12723 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v12709)+164)) = v12723
	v12725 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v12709)+168)) = v12725
	v12727 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11367)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12709)+172)) = uint8(v12727)
	v12729 = *(*int32)(unsafe.Add(mBase, uint32(v11367)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v12709)+176)) = v12729
	v12731 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v12732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8227)+334)))
	if v12732 == int32(0) {
		goto L1748
	} else {
		goto L1749
	}
L1743:
	;
	v12715 = F_is_parallel_safe(m, v8227, v12686)
	mBase = m.M
	v12716 = m.ExcPending
	if v12716 != 0 {
		goto L1
	} else {
		goto L1744
	}
L1744:
	;
	if v12715 == int32(0) {
		goto L1742
	} else {
		goto L1745
	}
L1745:
	;
	v12719 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12709)+26)) = uint8(v12719)
	goto L1742
L1746:
	;
	F_set_cheapest(m, v12709)
	mBase = m.M
	v13008 = m.ExcPending
	if v13008 != 0 {
		goto L1
	} else {
		goto L1824
	}
L1747:
	;
	if v8236 != 0 {
		goto L1768
	} else {
		goto L1769
	}
L1748:
	;
	v12735 = *(*int32)(unsafe.Add(mBase, uint32(v12731)+108))
	if v12735 == int32(0) {
		goto L1747
	} else {
		goto L1751
	}
L1749:
	;
	goto L1750
L1750:
	;
	v12738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12731)+36)))
	if v12738 != 0 {
		goto L1747
	} else {
		goto L1752
	}
L1751:
	;
	goto L1750
L1752:
	;
	v12739 = *(*int32)(unsafe.Add(mBase, uint32(v12731)+100))
	if v12739 != 0 {
		goto L1747
	} else {
		goto L1753
	}
L1753:
	;
	v12740 = *(*int32)(unsafe.Add(mBase, uint32(v12731)+108))
	if v12740 == int32(0) {
		goto L1754
	} else {
		goto L1755
	}
L1754:
	;
	v12817 = *(*int32)(unsafe.Add(mBase, uint32(v12709)+40))
	v12818 = *(*int32)(unsafe.Add(mBase, uint32(v12731)+112))
	v12819 = F_create_group_result_path(m, v8227, v12709, v12817, v12818)
	mBase = m.M
	v12820 = m.ExcPending
	if v12820 != 0 {
		goto L1
	} else {
		goto L1764
	}
L1755:
	;
	v12743 = *(*int32)(unsafe.Add(mBase, uint32(v12740)+4))
	if v12743 < int32(2) {
		goto L1754
	} else {
		goto L1756
	}
L1756:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+212)) = int64(0)
	v12750 = v12743
	v12765 = int32(0)
	goto L1757
L1757:
	;
	v12792 = *(*int32)(unsafe.Add(mBase, uint32(v12709)+40))
	v12793 = *(*int32)(unsafe.Add(mBase, uint32(v12731)+112))
	v12794 = F_create_group_result_path(m, v8227, v12709, v12792, v12793)
	mBase = m.M
	v12795 = m.ExcPending
	if v12795 != 0 {
		goto L1
	} else {
		goto L1759
	}
L1758:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+208)) = v12796
	v12801 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+216))
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+56)) = v12801
	v12803 = *(*int64)(unsafe.Add(mBase, uint32(v8230)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v8230)+48)) = v12803
	v12807 = int32(0)
	v12812 = F_create_append_path(m, v8227, v12709, v8230+int32(48), v12807, v12807, v12807, v12807, float64(-1))
	mBase = m.M
	v12813 = m.ExcPending
	if v12813 != 0 {
		goto L1
	} else {
		goto L1762
	}
L1759:
	;
	v12796 = F_lappend(m, v12765, v12794)
	mBase = m.M
	v12797 = m.ExcPending
	if v12797 != 0 {
		goto L1
	} else {
		goto L1760
	}
L1760:
	;
	if base.Ui32(int32(1)) < base.Ui32(v12750) {
		v12750 = v12750 - int32(1)
		v12765 = v12796
		goto L1757
	} else {
		goto L1761
	}
L1761:
	;
	goto L1758
L1762:
	;
	F_add_path(m, v12709, v12812)
	mBase = m.M
	v12815 = m.ExcPending
	if v12815 != 0 {
		goto L1
	} else {
		goto L1763
	}
L1763:
	;
	goto L1746
L1764:
	;
	F_add_path(m, v12709, v12819)
	mBase = m.M
	v12822 = m.ExcPending
	if v12822 != 0 {
		goto L1
	} else {
		goto L1765
	}
L1765:
	;
	goto L1746
L1766:
	;
	v12876 = *(*int32)(unsafe.Add(mBase, uint32(v12670)+100))
	if v12876 == int32(0) {
		v12924 = v12875
		goto L1786
	} else {
		goto L1787
	}
L1767:
	;
	v12875 = int32(1)
	goto L1766
L1768:
	;
	v12823 = *(*int32)(unsafe.Add(mBase, uint32(v8236)))
	if v12823 != 0 {
		goto L1767
	} else {
		goto L1771
	}
L1769:
	;
	goto L1770
L1770:
	;
	v12824 = int32(0)
	v12825 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+276))
	if v12825 == v12824 {
		goto L1773
	} else {
		goto L1774
	}
L1771:
	;
	goto L1770
L1772:
	;
	if v12870 == int32(0) {
		v12875 = v12824
		goto L1766
	} else {
		goto L1785
	}
L1773:
	;
	v12870 = int32(1)
	goto L1772
L1774:
	;
	goto L1775
L1775:
	;
	v12834 = *(*int32)(unsafe.Add(mBase, uint32(v12825)+4))
	if v12834 <= int32(0) {
		v12862 = int32(1)
		goto L1776
	} else {
		goto L1777
	}
L1776:
	;
	v12870 = v12862
	goto L1772
L1777:
	;
	v12837 = int32(0)
	if v12837 < v12834 {
		goto L1778
	} else {
		goto L1779
	}
L1778:
	;
	v12840 = v12834
	goto L1780
L1779:
	;
	v12840 = v12837
	goto L1780
L1780:
	;
	v12841 = *(*int32)(unsafe.Add(mBase, uint32(v12825)+12))
	v12843 = int32(0)
	goto L1781
L1781:
	;
	v12851 = *(*int32)(unsafe.Add(mBase, uint32(v12841+v12843<<(uint(int32(2))%32))))
	v12852 = *(*int32)(unsafe.Add(mBase, uint32(v12851)+12))
	v12853 = int32(0)
	v12854 = base.B2i32(v12852 != v12853)
	if v12852 == v12853 {
		v12862 = v12854
		goto L1776
	} else {
		goto L1783
	}
L1782:
	;
	v12862 = v12854
	goto L1776
L1783:
	;
	v12858 = v12843 + int32(1)
	if v12858 != v12840 {
		v12843 = v12858
		goto L1781
	} else {
		goto L1784
	}
L1784:
	;
	goto L1782
L1785:
	;
	goto L1767
L1786:
	;
	v12925 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v12926 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12925)+36)))
	if v12926 == int32(0) {
		goto L1809
	} else {
		goto L1810
	}
L1787:
	;
	v12879 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+352))
	if v12879 != 0 {
		v12924 = v12875
		goto L1786
	} else {
		goto L1788
	}
L1788:
	;
	if v8236 != 0 {
		goto L1790
	} else {
		goto L1791
	}
L1789:
	;
	v12924 = v12875 | int32(2)
	goto L1786
L1790:
	;
	v12880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8236)+16)))
	if v12880 != 0 {
		goto L1789
	} else {
		goto L1793
	}
L1791:
	;
	goto L1792
L1792:
	;
	v12881 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+276))
	v12882 = int32(0)
	if v12881 == v12882 {
		goto L1795
	} else {
		goto L1796
	}
L1793:
	;
	v12924 = v12875
	goto L1786
L1794:
	;
	if v12919 == int32(0) {
		v12924 = v12875
		goto L1786
	} else {
		goto L1807
	}
L1795:
	;
	v12919 = int32(1)
	goto L1794
L1796:
	;
	goto L1797
L1797:
	;
	v12889 = *(*int32)(unsafe.Add(mBase, uint32(v12881)+4))
	if v12889 <= int32(0) {
		v12913 = int32(1)
		goto L1798
	} else {
		goto L1799
	}
L1798:
	;
	v12919 = v12913
	goto L1794
L1799:
	;
	v12892 = int32(0)
	if v12892 < v12889 {
		goto L1800
	} else {
		goto L1801
	}
L1800:
	;
	v12895 = v12889
	goto L1802
L1801:
	;
	v12895 = v12892
	goto L1802
L1802:
	;
	v12896 = *(*int32)(unsafe.Add(mBase, uint32(v12881)+12))
	v12900 = v12882
	goto L1803
L1803:
	;
	v12904 = *(*int32)(unsafe.Add(mBase, uint32(v12896+v12900<<(uint(int32(2))%32))))
	v12905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12904)+18)))
	if v12905 != int32(1) {
		v12913 = v12905
		goto L1798
	} else {
		goto L1805
	}
L1804:
	;
	v12913 = v12905
	goto L1798
L1805:
	;
	v12909 = v12900 + int32(1)
	if v12909 != v12895 {
		v12900 = v12909
		goto L1803
	} else {
		goto L1806
	}
L1806:
	;
	goto L1804
L1807:
	;
	goto L1789
L1808:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v8230)+296)) = uint8(v12286)
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+208)) = v12938
	v12941 = *(*int32)(unsafe.Add(mBase, uint32(v12670)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+300)) = v12941
	v12943 = *(*int32)(unsafe.Add(mBase, uint32(v12670)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+304)) = v12943
	v12945 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8230)+212)) = uint8(v12945)
	v12947 = int32(1)
	v12949 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[6])))
	if v12949 == v12947 {
		goto L1819
	} else {
		goto L1820
	}
L1809:
	;
	v12929 = *(*int32)(unsafe.Add(mBase, uint32(v12925)+100))
	if v12929 == int32(0) {
		v12938 = v12924
		goto L1808
	} else {
		goto L1812
	}
L1810:
	;
	goto L1811
L1811:
	;
	v12932 = *(*int32)(unsafe.Add(mBase, uint32(v12925)+108))
	if v12932 != 0 {
		v12938 = v12924
		goto L1808
	} else {
		goto L1813
	}
L1812:
	;
	goto L1811
L1813:
	;
	v12933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8227)+356)))
	if v12933 != 0 {
		v12938 = v12924
		goto L1808
	} else {
		goto L1814
	}
L1814:
	;
	v12936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8227)+357)))
	if v12936 != 0 {
		goto L1815
	} else {
		goto L1816
	}
L1815:
	;
	v12937 = v12924
	goto L1817
L1816:
	;
	v12937 = v12924 | int32(4)
	goto L1817
L1817:
	;
	v12938 = v12937
	goto L1808
L1818:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8230)+308)) = v12956
	F_create_ordinary_grouping_paths(m, v8227, v11367, v12709, v8230+int32(320), v8236, v8230+int32(208), v8230+int32(364))
	mBase = m.M
	v12965 = m.ExcPending
	if v12965 != 0 {
		goto L1
	} else {
		goto L1823
	}
L1819:
	;
	v12952 = *(*int32)(unsafe.Add(mBase, uint32(v12670)+108))
	if v12952 == int32(0) {
		v12956 = v12947
		goto L1818
	} else {
		goto L1822
	}
L1820:
	;
	goto L1821
L1821:
	;
	v12956 = int32(0)
	goto L1818
L1822:
	;
	goto L1821
L1823:
	;
	goto L1746
L1824:
	;
	v13009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8232)+38)))
	if v13009 != int32(1) {
		v13019 = v12709
		goto L1731
	} else {
		goto L1825
	}
L1825:
	;
	v13012 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+188))
	v13013 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+184))
	F_adjust_paths_for_srfs(m, v8227, v12709, v13012, v13013)
	mBase = m.M
	v13015 = m.ExcPending
	if v13015 != 0 {
		goto L1
	} else {
		goto L1826
	}
L1826:
	;
	v13019 = v12709
	goto L1731
L1827:
	;
	v14323 = *(*int32)(unsafe.Add(mBase, uint32(v14295)+120))
	if v14323 == int32(0) {
		goto L2053
	} else {
		goto L2054
	}
L1828:
	;
	v14286 = v13019
	v14287 = v12654
	v14288 = v11704
	v14290 = v8227
	v14293 = v8230
	v14295 = v8232
	v14302 = v12658
	v14305 = v8242
	v14307 = v11375
	v14316 = v8253
	v14320 = v8257
	v14321 = v8258
	goto L1827
L1829:
	;
	goto L1830
L1830:
	;
	v13061 = F_fetch_upper_rel(m, v8227, int32(3), int32(0))
	mBase = m.M
	v13062 = m.ExcPending
	if v13062 != 0 {
		goto L1
	} else {
		goto L1831
	}
L1831:
	;
	v13063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13019)+26)))
	if v11717&v13063 != int32(1) {
		goto L1832
	} else {
		goto L1833
	}
L1832:
	;
	v13073 = *(*int32)(unsafe.Add(mBase, uint32(v13019)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v13061)+164)) = v13073
	v13075 = *(*int32)(unsafe.Add(mBase, uint32(v13019)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v13061)+168)) = v13075
	v13077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13019)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13061)+172)) = uint8(v13077)
	v13079 = *(*int32)(unsafe.Add(mBase, uint32(v13019)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v13061)+176)) = v13079
	v13081 = *(*int32)(unsafe.Add(mBase, uint32(v13019)+44))
	if v13081 == int32(0) {
		v14217 = v13079
		v14220 = v13061
		v14221 = v12654
		v14222 = v11704
		v14224 = v8227
		v14227 = v8230
		v14229 = v8232
		v14236 = v12658
		v14239 = v8242
		v14241 = v11375
		v14250 = v8253
		v14254 = v8257
		v14255 = v8258
		goto L1836
	} else {
		goto L1837
	}
L1833:
	;
	v13067 = F_is_parallel_safe(m, v8227, v10510)
	mBase = m.M
	v13068 = m.ExcPending
	if v13068 != 0 {
		goto L1
	} else {
		goto L1834
	}
L1834:
	;
	if v13067 == int32(0) {
		goto L1832
	} else {
		goto L1835
	}
L1835:
	;
	v13071 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13061)+26)) = uint8(v13071)
	goto L1832
L1836:
	;
	if v14217 == int32(0) {
		goto L2042
	} else {
		goto L2043
	}
L1837:
	;
	v13084 = *(*int32)(unsafe.Add(mBase, uint32(v13081)+4))
	if v13084 <= int32(0) {
		v14217 = v13079
		v14220 = v13061
		v14221 = v12654
		v14222 = v11704
		v14224 = v8227
		v14227 = v8230
		v14229 = v8232
		v14236 = v12658
		v14239 = v8242
		v14241 = v11375
		v14250 = v8253
		v14254 = v8257
		v14255 = v8258
		goto L1836
	} else {
		goto L1838
	}
L1838:
	;
	v13116 = v8248
	goto L1839
L1839:
	;
	v13128 = *(*int32)(unsafe.Add(mBase, uint32(v13081)+12))
	v13132 = *(*int32)(unsafe.Add(mBase, uint32(v13128+v13116<<(uint(int32(2))%32))))
	v13133 = *(*int32)(unsafe.Add(mBase, uint32(v13019)+60))
	if v13132 == v13133 {
		goto L1842
	} else {
		goto L1843
	}
L1840:
	;
	v14215 = *(*int32)(unsafe.Add(mBase, uint32(v13061)+176))
	v14217 = v14215
	v14220 = v13061
	v14221 = v12654
	v14222 = v11704
	v14224 = v8227
	v14227 = v8230
	v14229 = v8232
	v14236 = v12658
	v14239 = v8242
	v14241 = v11375
	v14250 = v8253
	v14254 = v8257
	v14255 = v8258
	goto L1836
L1841:
	;
	v14212 = v13116 + int32(1)
	v14213 = *(*int32)(unsafe.Add(mBase, uint32(v13081)+4))
	if v14212 < v14213 {
		v13116 = v14212
		goto L1839
	} else {
		goto L2041
	}
L1842:
	;
	v13220 = int32(0)
	v13222 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	if v13220 < v13222 {
		goto L1878
	} else {
		goto L1879
	}
L1843:
	;
	v13135 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+188))
	v13136 = *(*int32)(unsafe.Add(mBase, uint32(v13132)+64))
	v13138 = v8230 + int32(320)
	if v13135 == v13136 {
		goto L1846
	} else {
		goto L1847
	}
L1844:
	;
	if v13216 != 0 {
		goto L1842
	} else {
		goto L1876
	}
L1845:
	;
	v13204 = *(*int32)(unsafe.Add(mBase, uint32(v13135)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13204
	v13216 = int32(1)
	goto L1844
L1846:
	;
	if v13135 != 0 {
		goto L1845
	} else {
		goto L1849
	}
L1847:
	;
	goto L1848
L1848:
	;
	if v13135 == int32(0) {
		goto L1850
	} else {
		goto L1851
	}
L1849:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = int32(0)
	v13216 = int32(1)
	goto L1844
L1850:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = int32(0)
	v13216 = int32(1)
	goto L1844
L1851:
	;
	goto L1852
L1852:
	;
	if v13136 == int32(0) {
		goto L1853
	} else {
		goto L1854
	}
L1853:
	;
	v13156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13156
	v13216 = v13156
	goto L1844
L1854:
	;
	goto L1855
L1855:
	;
	v13159 = *(*int32)(unsafe.Add(mBase, uint32(v13136)+4))
	v13160 = int32(0)
	if v13160 < v13159 {
		goto L1856
	} else {
		goto L1857
	}
L1856:
	;
	v13163 = v13159
	goto L1858
L1857:
	;
	v13163 = v13160
	goto L1858
L1858:
	;
	v13164 = *(*int32)(unsafe.Add(mBase, uint32(v13135)+4))
	v13169 = int32(0)
	goto L1859
L1859:
	;
	if v13169 < v13164 {
		goto L1861
	} else {
		goto L1862
	}
L1861:
	;
	v13176 = *(*int32)(unsafe.Add(mBase, uint32(v13135)+12))
	v13180 = v13176 + v13169<<(uint(int32(2))%32)
	goto L1863
L1862:
	;
	v13180 = int32(0)
	goto L1863
L1863:
	;
	if v13169 == v13163 {
		goto L1864
	} else {
		goto L1865
	}
L1864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13163
	v13216 = base.B2i32(v13180 == int32(0))
	goto L1844
L1865:
	;
	goto L1866
L1866:
	;
	v13186 = base.B2i32(v13180 == int32(0))
	if v13180 == int32(0) {
		goto L1867
	} else {
		goto L1868
	}
L1867:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13169
	v13216 = v13186
	goto L1844
L1868:
	;
	goto L1869
L1869:
	;
	v13190 = *(*int32)(unsafe.Add(mBase, uint32(v13136)+12))
	if v13190 == int32(0) {
		goto L1870
	} else {
		goto L1871
	}
L1870:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13169
	v13216 = v13186
	goto L1844
L1871:
	;
	goto L1872
L1872:
	;
	v13194 = *(*int32)(unsafe.Add(mBase, uint32(v13180)))
	v13198 = *(*int32)(unsafe.Add(mBase, uint32(v13190+v13169<<(uint(int32(2))%32))))
	if v13194 != v13198 {
		goto L1873
	} else {
		goto L1874
	}
L1873:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13138))) = v13169
	v13216 = int32(0)
	goto L1844
L1874:
	;
	v13169 = v13169 + int32(1)
	goto L1859
L1876:
	;
	v13217 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+320))
	if v13217 <= int32(0) {
		goto L1841
	} else {
		goto L1877
	}
L1877:
	;
	goto L1842
L1878:
	;
	v13226 = v13220
	v13232 = v12656
	v13240 = v13132
	v13242 = v13220
	goto L1881
L1879:
	;
	v14142 = v13132
	goto L1880
L1880:
	;
	F_add_path(m, v13061, v14142)
	mBase = m.M
	v14169 = m.ExcPending
	if v14169 != 0 {
		goto L1
	} else {
		goto L2040
	}
L1881:
	;
	v13266 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+12))
	v13269 = v13266 + v13226<<(uint(int32(2))%32)
	v13270 = *(*int32)(unsafe.Add(mBase, uint32(v13269)))
	v13271 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+284))
	v13272 = F_make_pathkeys_for_window(m, v8227, v13270, v13271)
	mBase = m.M
	v13273 = m.ExcPending
	if v13273 != 0 {
		goto L1
	} else {
		goto L1884
	}
L1882:
	;
	v14142 = v13777
	goto L1880
L1883:
	;
	v13370 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+12))
	v13371 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	if base.Ui32(v13269+int32(4)) < base.Ui32(v13370+v13371<<(uint(int32(2))%32)) {
		goto L1925
	} else {
		goto L1926
	}
L1884:
	;
	v13274 = *(*int32)(unsafe.Add(mBase, uint32(v13240)+64))
	v13276 = v8230 + int32(208)
	if v13272 == v13274 {
		goto L1887
	} else {
		goto L1888
	}
L1885:
	;
	if v13354 != 0 {
		v13367 = v13240
		goto L1883
	} else {
		goto L1917
	}
L1886:
	;
	v13342 = *(*int32)(unsafe.Add(mBase, uint32(v13272)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13342
	v13354 = int32(1)
	goto L1885
L1887:
	;
	if v13272 != 0 {
		goto L1886
	} else {
		goto L1890
	}
L1888:
	;
	goto L1889
L1889:
	;
	if v13272 == int32(0) {
		goto L1891
	} else {
		goto L1892
	}
L1890:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = int32(0)
	v13354 = int32(1)
	goto L1885
L1891:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = int32(0)
	v13354 = int32(1)
	goto L1885
L1892:
	;
	goto L1893
L1893:
	;
	if v13274 == int32(0) {
		goto L1894
	} else {
		goto L1895
	}
L1894:
	;
	v13294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13294
	v13354 = v13294
	goto L1885
L1895:
	;
	goto L1896
L1896:
	;
	v13297 = *(*int32)(unsafe.Add(mBase, uint32(v13274)+4))
	v13298 = int32(0)
	if v13298 < v13297 {
		goto L1897
	} else {
		goto L1898
	}
L1897:
	;
	v13301 = v13297
	goto L1899
L1898:
	;
	v13301 = v13298
	goto L1899
L1899:
	;
	v13302 = *(*int32)(unsafe.Add(mBase, uint32(v13272)+4))
	v13307 = int32(0)
	goto L1900
L1900:
	;
	if v13307 < v13302 {
		goto L1902
	} else {
		goto L1903
	}
L1902:
	;
	v13314 = *(*int32)(unsafe.Add(mBase, uint32(v13272)+12))
	v13318 = v13314 + v13307<<(uint(int32(2))%32)
	goto L1904
L1903:
	;
	v13318 = int32(0)
	goto L1904
L1904:
	;
	if v13307 == v13301 {
		goto L1905
	} else {
		goto L1906
	}
L1905:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13301
	v13354 = base.B2i32(v13318 == int32(0))
	goto L1885
L1906:
	;
	goto L1907
L1907:
	;
	v13324 = base.B2i32(v13318 == int32(0))
	if v13318 == int32(0) {
		goto L1908
	} else {
		goto L1909
	}
L1908:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13307
	v13354 = v13324
	goto L1885
L1909:
	;
	goto L1910
L1910:
	;
	v13328 = *(*int32)(unsafe.Add(mBase, uint32(v13274)+12))
	if v13328 == int32(0) {
		goto L1911
	} else {
		goto L1912
	}
L1911:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13307
	v13354 = v13324
	goto L1885
L1912:
	;
	goto L1913
L1913:
	;
	v13332 = *(*int32)(unsafe.Add(mBase, uint32(v13318)))
	v13336 = *(*int32)(unsafe.Add(mBase, uint32(v13328+v13307<<(uint(int32(2))%32))))
	if v13332 != v13336 {
		goto L1914
	} else {
		goto L1915
	}
L1914:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13276))) = v13307
	v13354 = int32(0)
	goto L1885
L1915:
	;
	v13307 = v13307 + int32(1)
	goto L1900
L1917:
	;
	v13355 = *(*int32)(unsafe.Add(mBase, uint32(v8230)+208))
	if v13355 != 0 {
		goto L1919
	} else {
		goto L1920
	}
L1918:
	;
	v13364 = F_create_incremental_sort_path(m, v8227, v13061, v13240, v13272, v13355, float64(-1))
	mBase = m.M
	v13365 = m.ExcPending
	if v13365 != 0 {
		goto L1
	} else {
		goto L1924
	}
L1919:
	;
	v13357 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	if v13357&int32(1) != 0 {
		goto L1918
	} else {
		goto L1922
	}
L1920:
	;
	goto L1921
L1921:
	;
	v13361 = F_create_sort_path(m, v13061, v13240, v13272, float64(-1))
	mBase = m.M
	v13362 = m.ExcPending
	if v13362 != 0 {
		goto L1
	} else {
		goto L1923
	}
L1922:
	;
	goto L1921
L1923:
	;
	v13367 = v13361
	goto L1883
L1924:
	;
	v13367 = v13364
	goto L1883
L1925:
	;
	v13376 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13232)+32)))
	v13377 = F_copy_pathtarget(m, v13232)
	mBase = m.M
	v13378 = m.ExcPending
	if v13378 != 0 {
		goto L1
	} else {
		goto L1928
	}
L1926:
	;
	v13498 = v12654
	v13499 = v13371
	goto L1927
L1927:
	;
	v13540 = v13499 - int32(1)
	v13541 = *(*int32)(unsafe.Add(mBase, uint32(v10514)+8))
	v13542 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+48))
	v13546 = *(*int32)(unsafe.Add(mBase, uint32(v13541+v13542<<(uint(int32(2))%32))))
	if v13546 == int32(0) {
		goto L1942
	} else {
		goto L1943
	}
L1928:
	;
	v13379 = *(*int32)(unsafe.Add(mBase, uint32(v10514)+8))
	v13380 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+48))
	v13384 = *(*int32)(unsafe.Add(mBase, uint32(v13379+v13380<<(uint(int32(2))%32))))
	if v13384 == int32(0) {
		v13490 = v13376
		goto L1929
	} else {
		goto L1930
	}
L1929:
	;
	v13491 = int64(1073741823)
	if v13491 <= v13490 {
		goto L1938
	} else {
		goto L1939
	}
L1930:
	;
	v13387 = int32(0)
	v13388 = *(*int32)(unsafe.Add(mBase, uint32(v13384)+4))
	if v13388 <= v13387 {
		v13490 = v13376
		goto L1929
	} else {
		goto L1931
	}
L1931:
	;
	v13392 = v13387
	v13431 = v13376
	goto L1932
L1932:
	;
	v13432 = *(*int32)(unsafe.Add(mBase, uint32(v13384)+12))
	v13436 = *(*int32)(unsafe.Add(mBase, uint32(v13432+v13392<<(uint(int32(2))%32))))
	F_add_column_to_pathtarget(m, v13377, v13436, int32(0))
	mBase = m.M
	v13439 = m.ExcPending
	if v13439 != 0 {
		goto L1
	} else {
		goto L1934
	}
L1933:
	;
	v13490 = v13445
	goto L1929
L1934:
	;
	v13440 = *(*int32)(unsafe.Add(mBase, uint32(v13436)+8))
	v13442 = F_get_typavgwidth(m, v13440, int32(-1))
	mBase = m.M
	v13443 = m.ExcPending
	if v13443 != 0 {
		goto L1
	} else {
		goto L1935
	}
L1935:
	;
	v13445 = v13431 + base.I64_extend_i32_s(v13442)
	v13447 = v13392 + int32(1)
	v13448 = *(*int32)(unsafe.Add(mBase, uint32(v13384)+4))
	if v13447 < v13448 {
		v13392 = v13447
		v13431 = v13445
		goto L1932
	} else {
		goto L1936
	}
L1936:
	;
	goto L1933
L1937:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13377)+32)) = base.I32_wrap_i64(v13494)
	v13497 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	v13498 = v13377
	v13499 = v13497
	goto L1927
L1938:
	;
	v13494 = v13491
	goto L1940
L1939:
	;
	v13494 = v13490
	goto L1940
L1940:
	;
	goto L1937
L1941:
	;
	v13773 = base.B2i32(v13226 == v13540)
	if v13226 == v13540 {
		goto L1968
	} else {
		goto L1969
	}
L1942:
	;
	v13549 = int32(0)
	v13731 = v13549
	v13745 = v13549
	v13748 = v13242
	goto L1941
L1943:
	;
	goto L1944
L1944:
	;
	v13551 = int32(0)
	v13553 = *(*int32)(unsafe.Add(mBase, uint32(v13546)+4))
	if v13553 <= v13551 {
		v13731 = v13551
		v13745 = v13546
		v13748 = v13242
		goto L1941
	} else {
		goto L1945
	}
L1945:
	;
	v13556 = v13551
	v13557 = v13553
	v13573 = v13242
	v13575 = v13551
	goto L1946
L1946:
	;
	v13597 = *(*int32)(unsafe.Add(mBase, uint32(v13546)+12))
	v13601 = *(*int32)(unsafe.Add(mBase, uint32(v13597+v13575<<(uint(int32(2))%32))))
	v13602 = *(*int32)(unsafe.Add(mBase, uint32(v13601)+28))
	if v13602 == int32(0) {
		v13681 = v13556
		v13682 = v13557
		v13698 = v13573
		goto L1948
	} else {
		goto L1949
	}
L1947:
	;
	v13725 = *(*int32)(unsafe.Add(mBase, uint32(v10514)+8))
	v13726 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+48))
	v13730 = *(*int32)(unsafe.Add(mBase, uint32(v13725+v13726<<(uint(int32(2))%32))))
	v13731 = v13681
	v13745 = v13730
	v13748 = v13698
	goto L1941
L1948:
	;
	v13723 = v13575 + int32(1)
	if v13723 < v13682 {
		v13556 = v13681
		v13557 = v13682
		v13573 = v13698
		v13575 = v13723
		goto L1946
	} else {
		goto L1967
	}
L1949:
	;
	v13605 = int32(0)
	v13606 = *(*int32)(unsafe.Add(mBase, uint32(v13602)+4))
	if v13606 <= v13605 {
		v13681 = v13556
		v13682 = v13557
		v13698 = v13573
		goto L1948
	} else {
		goto L1950
	}
L1950:
	;
	v13609 = v13556
	v13618 = v13605
	v13626 = v13573
	goto L1951
L1951:
	;
	v13650 = *(*int32)(unsafe.Add(mBase, uint32(v13602)+12))
	v13654 = *(*int32)(unsafe.Add(mBase, uint32(v13650+v13618<<(uint(int32(2))%32))))
	v13655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13654)+12)))
	if v13655 != 0 {
		goto L1954
	} else {
		goto L1955
	}
L1952:
	;
	v13680 = *(*int32)(unsafe.Add(mBase, uint32(v13546)+4))
	v13681 = v13670
	v13682 = v13680
	v13698 = v13675
	goto L1948
L1953:
	;
	v13664 = F_copyObjectImpl(m, v13663)
	mBase = m.M
	v13665 = m.ExcPending
	if v13665 != 0 {
		goto L1
	} else {
		goto L1959
	}
L1954:
	;
	v13656 = F_copyObjectImpl(m, v13601)
	mBase = m.M
	v13657 = m.ExcPending
	if v13657 != 0 {
		goto L1
	} else {
		goto L1957
	}
L1955:
	;
	goto L1956
L1956:
	;
	v13659 = *(*int32)(unsafe.Add(mBase, uint32(v13654)+16))
	v13660 = F_copyObjectImpl(m, v13659)
	mBase = m.M
	v13661 = m.ExcPending
	if v13661 != 0 {
		goto L1
	} else {
		goto L1958
	}
L1957:
	;
	v13658 = *(*int32)(unsafe.Add(mBase, uint32(v13654)+16))
	v13662 = v13656
	v13663 = v13658
	goto L1953
L1958:
	;
	v13662 = v13660
	v13663 = v13601
	goto L1953
L1959:
	;
	v13666 = *(*int32)(unsafe.Add(mBase, uint32(v13654)+4))
	v13667 = *(*int32)(unsafe.Add(mBase, uint32(v13654)+8))
	v13668 = F_make_opclause(m, v13666, v13662, v13664, v13667)
	mBase = m.M
	v13669 = m.ExcPending
	if v13669 != 0 {
		goto L1
	} else {
		goto L1960
	}
L1960:
	;
	v13670 = F_lappend(m, v13609, v13668)
	mBase = m.M
	v13671 = m.ExcPending
	if v13671 != 0 {
		goto L1
	} else {
		goto L1961
	}
L1961:
	;
	if v13226 != v13540 {
		goto L1962
	} else {
		goto L1963
	}
L1962:
	;
	v13673 = F_lappend(m, v13626, v13668)
	mBase = m.M
	v13674 = m.ExcPending
	if v13674 != 0 {
		goto L1
	} else {
		goto L1965
	}
L1963:
	;
	v13675 = v13626
	goto L1964
L1964:
	;
	v13677 = v13618 + int32(1)
	v13678 = *(*int32)(unsafe.Add(mBase, uint32(v13602)+4))
	if v13677 < v13678 {
		v13609 = v13670
		v13618 = v13677
		v13626 = v13675
		goto L1951
	} else {
		goto L1966
	}
L1965:
	;
	v13675 = v13673
	goto L1964
L1966:
	;
	goto L1952
L1967:
	;
	goto L1947
L1968:
	;
	v13774 = v13748
	goto L1970
L1969:
	;
	v13774 = int32(0)
	goto L1970
L1970:
	;
	v13777 = F_palloc0(m, int32(96))
	mBase = m.M
	v13778 = m.ExcPending
	if v13778 != 0 {
		goto L1
	} else {
		goto L1971
	}
L1971:
	;
	v13779 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13777)+20)) = uint8(v13779)
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+16)) = v13779
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+12)) = v13498
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+8)) = v13061
	*(*int64)(unsafe.Add(mBase, uint32(v13777))) = int64(1589137899834)
	v13787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13061)+26)))
	if v13787 == int32(1) {
		goto L1972
	} else {
		goto L1973
	}
L1972:
	;
	v13790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13367)+21)))
	v13792 = v13790
	goto L1974
L1973:
	;
	v13792 = int32(0)
	goto L1974
L1974:
	;
	v13794 = v13792 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13777)+21)) = uint8(v13794)
	v13796 = *(*int32)(unsafe.Add(mBase, uint32(v13367)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+24)) = v13796
	v13798 = *(*int32)(unsafe.Add(mBase, uint32(v13367)+64))
	*(*uint8)(unsafe.Add(mBase, uint32(v13777)+88)) = uint8(v13773)
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+84)) = v13731
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+80)) = v13774
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+76)) = v13270
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+72)) = v13367
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+64)) = v13798
	v13805 = *(*int32)(unsafe.Add(mBase, uint32(v13367)+40))
	v13806 = *(*float64)(unsafe.Add(mBase, uint32(v13367)+48))
	v13807 = *(*float64)(unsafe.Add(mBase, uint32(v13367)+56))
	v13808 = *(*float64)(unsafe.Add(mBase, uint32(v13367)+32))
	v13809 = int32(0)
	v13811 = m.G0
	v13813 = v13811 - int32(48)
	m.G0 = v13813
	v13815 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+12))
	if v13815 != 0 {
		goto L1975
	} else {
		goto L1976
	}
L1975:
	;
	v13816 = *(*int32)(unsafe.Add(mBase, uint32(v13815)+4))
	v13817 = v13816
	goto L1977
L1976:
	;
	v13817 = int32(0)
	goto L1977
L1977:
	;
	v13818 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+16))
	if v13818 != 0 {
		goto L1978
	} else {
		goto L1979
	}
L1978:
	;
	v13819 = *(*int32)(unsafe.Add(mBase, uint32(v13818)+4))
	v13820 = v13819
	goto L1980
L1979:
	;
	v13820 = v13809
	goto L1980
L1980:
	;
	if v13745 == int32(0) {
		v13956 = v13807
		v13961 = v13806
		goto L1981
	} else {
		goto L1982
	}
L1981:
	;
	v13968 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[5]))
	v13970 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+48)) = v13961
	*(*int32)(unsafe.Add(mBase, uint32(v13777)+40)) = v13805
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+32)) = v13808
	v13980 = base.F64_add(base.F64_mul(v13970, v13808), base.F64_add(base.F64_mul(base.F64_mul(v13968, base.F64_convert_i32_s(v13820+v13817)), v13808), v13956))
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+56)) = v13980
	v13982 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+20))
	v13983 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+12))
	if v13983 != 0 {
		goto L1990
	} else {
		goto L1991
	}
L1982:
	;
	v13823 = *(*int32)(unsafe.Add(mBase, uint32(v13745)+4))
	if v13823 <= int32(0) {
		v13956 = v13807
		v13961 = v13806
		goto L1981
	} else {
		goto L1983
	}
L1983:
	;
	v13827 = v13813 + int32(32)
	v13837 = v13809
	v13858 = v13807
	v13863 = v13806
	goto L1984
L1984:
	;
	v13869 = *(*int32)(unsafe.Add(mBase, uint32(v13745)+12))
	v13873 = *(*int32)(unsafe.Add(mBase, uint32(v13869+v13837<<(uint(int32(2))%32))))
	v13874 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+16)) = v13874
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+8)) = v13874
	v13878 = *(*int32)(unsafe.Add(mBase, uint32(v13873)+4))
	F_add_function_cost(m, v8227, v13878, v13873, v13813+int32(8))
	mBase = m.M
	v13882 = m.ExcPending
	if v13882 != 0 {
		goto L1
	} else {
		goto L1986
	}
L1985:
	;
	v13956 = v13921
	v13961 = v13916
	goto L1981
L1986:
	;
	v13883 = *(*int32)(unsafe.Add(mBase, uint32(v13873)+20))
	v13884 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+16))
	v13885 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13813)+24)) = v8227
	v13887 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13827)+8)) = v13887
	*(*int64)(unsafe.Add(mBase, uint32(v13827))) = v13887
	v13892 = v13813 + int32(24)
	v13893 = F_cost_qual_eval_walker(m, v13883, v13892)
	mBase = m.M
	v13894 = m.ExcPending
	if v13894 != 0 {
		goto L1
	} else {
		goto L1987
	}
L1987:
	;
	v13895 = *(*int64)(unsafe.Add(mBase, uint32(v13827)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+16)) = v13895
	v13897 = *(*int64)(unsafe.Add(mBase, uint32(v13827)))
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+8)) = v13897
	v13899 = *(*int32)(unsafe.Add(mBase, uint32(v13873)+24))
	v13900 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+16))
	v13901 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13813)+24)) = v8227
	v13903 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13827)+8)) = v13903
	*(*int64)(unsafe.Add(mBase, uint32(v13827))) = v13903
	v13907 = F_cost_qual_eval_walker(m, v13899, v13892)
	mBase = m.M
	v13908 = m.ExcPending
	if v13908 != 0 {
		goto L1
	} else {
		goto L1988
	}
L1988:
	;
	v13909 = *(*int64)(unsafe.Add(mBase, uint32(v13827)))
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+8)) = v13909
	v13911 = *(*int64)(unsafe.Add(mBase, uint32(v13827)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13813)+16)) = v13911
	v13915 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+8))
	v13916 = base.F64_add(base.F64_add(v13901, base.F64_add(v13863, v13885)), v13915)
	v13918 = *(*float64)(unsafe.Add(mBase, uint32(v13813)+16))
	v13921 = base.F64_add(base.F64_mul(base.F64_add(base.F64_add(v13884, v13900), v13918), v13808), v13858)
	v13923 = v13837 + int32(1)
	v13924 = *(*int32)(unsafe.Add(mBase, uint32(v13745)+4))
	if v13923 < v13924 {
		v13837 = v13923
		v13858 = v13921
		v13863 = v13916
		goto L1984
	} else {
		goto L1989
	}
L1989:
	;
	goto L1985
L1990:
	;
	v13984 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v13985 = *(*int32)(unsafe.Add(mBase, uint32(v13984)+76))
	v13986 = F_get_sortgrouplist_exprs(m, v13983, v13985)
	mBase = m.M
	v13987 = m.ExcPending
	if v13987 != 0 {
		goto L1
	} else {
		goto L1993
	}
L1991:
	;
	v13996 = v13808
	goto L1992
L1992:
	;
	v13998 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+16))
	if v13998 != 0 {
		goto L1996
	} else {
		goto L1997
	}
L1993:
	;
	v13988 = int32(0)
	v13990 = F_estimate_num_groups(m, v8227, v13986, v13808, v13988, v13988)
	mBase = m.M
	v13991 = m.ExcPending
	if v13991 != 0 {
		goto L1
	} else {
		goto L1994
	}
L1994:
	;
	F_list_free(m, v13986)
	mBase = m.M
	v13993 = m.ExcPending
	if v13993 != 0 {
		goto L1
	} else {
		goto L1995
	}
L1995:
	;
	v13996 = base.F64_div(v13808, v13990)
	goto L1992
L1996:
	;
	v13999 = *(*int32)(unsafe.Add(mBase, uint32(v8227)+4))
	v14000 = *(*int32)(unsafe.Add(mBase, uint32(v13999)+76))
	v14001 = F_get_sortgrouplist_exprs(m, v13998, v14000)
	mBase = m.M
	v14002 = m.ExcPending
	if v14002 != 0 {
		goto L1
	} else {
		goto L1999
	}
L1997:
	;
	v14013 = float64(1)
	goto L1998
L1998:
	;
	if v13982&int32(256) != 0 {
		v14067 = v13996
		goto L2002
	} else {
		goto L2003
	}
L1999:
	;
	v14003 = int32(0)
	v14005 = F_estimate_num_groups(m, v8227, v14001, v13996, v14003, v14003)
	mBase = m.M
	v14006 = m.ExcPending
	if v14006 != 0 {
		goto L1
	} else {
		goto L2000
	}
L2000:
	;
	F_list_free(m, v14001)
	mBase = m.M
	v14008 = m.ExcPending
	if v14008 != 0 {
		goto L1
	} else {
		goto L2001
	}
L2001:
	;
	v14013 = base.F64_div(v13996, v14005)
	goto L1998
L2002:
	;
	v14069 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+12))
	if v14069 == int32(0) {
		goto L2026
	} else {
		goto L2027
	}
L2003:
	;
	if v13982&int32(1024) != 0 {
		goto L2004
	} else {
		goto L2005
	}
L2004:
	;
	if base.B2i32(v13982&int32(10) == int32(0))|v13982&int32(4) != 0 {
		v14067 = float64(1)
		goto L2002
	} else {
		goto L2007
	}
L2005:
	;
	goto L2006
L2006:
	;
	v14028 = float64(1)
	if v13982&int32(_a_F_subquery_planner_42) != int32(_a_F_subquery_planner_43) {
		v14067 = v14028
		goto L2002
	} else {
		goto L2011
	}
L2007:
	;
	v14026 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+16))
	if v14026 != 0 {
		goto L2008
	} else {
		goto L2009
	}
L2008:
	;
	v14027 = v14013
	goto L2010
L2009:
	;
	v14027 = v13996
	goto L2010
L2010:
	;
	v14067 = v14027
	goto L2002
L2011:
	;
	v14033 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+28))
	v14034 = *(*int32)(unsafe.Add(mBase, uint32(v14033)))
	if v14034 == int32(7) {
		goto L2013
	} else {
		goto L2014
	}
L2012:
	;
	if v13982&int32(4) != 0 {
		goto L2021
	} else {
		goto L2022
	}
L2013:
	;
	v14038 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14033)+32)))
	if v14038 != 0 {
		v14054 = float64(1)
		goto L2012
	} else {
		goto L2016
	}
L2014:
	;
	goto L2015
L2015:
	;
	v14054 = base.F64_mul(base.F64_div(v13996, v14013), float64(0.3333333333333333))
	goto L2012
L2016:
	;
	v14039 = *(*int32)(unsafe.Add(mBase, uint32(v14033)+4))
	switch v14039 - int32(20) {
	case 0:
		goto L2018
	case 1:
		goto L2020
	default:
		goto L2017
	case 3:
		goto L2019
	}
L2017:
	;
	v14054 = base.F64_mul(base.F64_div(v13996, v14013), float64(0.3333333333333333))
	goto L2012
L2018:
	;
	v14046 = *(*int64)(unsafe.Add(mBase, uint32(v14033)+24))
	v14054 = base.F64_convert_i64_s(v14046)
	goto L2012
L2019:
	;
	v14044 = *(*int32)(unsafe.Add(mBase, uint32(v14033)+24))
	v14054 = base.F64_convert_i32_s(v14044)
	goto L2012
L2020:
	;
	v14042 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14033)+24)))
	v14054 = base.F64_convert_i32_s(v14042)
	goto L2012
L2021:
	;
	v14067 = base.F64_add(v14054, float64(1))
	goto L2002
L2022:
	;
	goto L2023
L2023:
	;
	if v13982&int32(10) == int32(0) {
		v14067 = v14028
		goto L2002
	} else {
		goto L2024
	}
L2024:
	;
	v14067 = base.F64_mul(v14013, base.F64_add(v14054, float64(1)))
	goto L2002
L2025:
	;
	if base.F64_gt(v13996, v14077) != 0 {
		goto L2031
	} else {
		goto L2032
	}
L2026:
	;
	v14072 = *(*int32)(unsafe.Add(mBase, uint32(v13270)+16))
	if v14072 == int32(0) {
		v14077 = v14067
		goto L2025
	} else {
		goto L2029
	}
L2027:
	;
	goto L2028
L2028:
	;
	v14077 = base.F64_add(v14067, float64(1))
	goto L2025
L2029:
	;
	goto L2028
L2030:
	;
	m.G0 = v13813 + int32(48)
	v14111 = *(*float64)(unsafe.Add(mBase, uint32(v13498)+16))
	v14112 = *(*float64)(unsafe.Add(mBase, uint32(v13777)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+48)) = base.F64_add(v14111, v14112)
	v14115 = *(*float64)(unsafe.Add(mBase, uint32(v13777)+56))
	v14116 = *(*float64)(unsafe.Add(mBase, uint32(v13498)+24))
	v14117 = *(*float64)(unsafe.Add(mBase, uint32(v13777)+32))
	v14119 = *(*float64)(unsafe.Add(mBase, uint32(v13498)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+56)) = base.F64_add(v14115, base.F64_add(base.F64_mul(v14116, v14117), v14119))
	v14124 = v13226 + int32(1)
	v14125 = *(*int32)(unsafe.Add(mBase, uint32(v10510)+4))
	if v14124 < v14125 {
		v13226 = v14124
		v13232 = v13498
		v13240 = v13777
		v13242 = v13748
		goto L1881
	} else {
		goto L2039
	}
L2031:
	;
	v14080 = v14077
	goto L2033
L2032:
	;
	v14080 = v13996
	goto L2033
L2033:
	;
	if base.F64_gt(v14080, float64(1e+100))|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v14080)&int64(9223372036854775807))) == int32(0) {
		goto L2034
	} else {
		goto L2035
	}
L2034:
	;
	if base.F64_le(v14080, float64(1)) != 0 {
		goto L2030
	} else {
		goto L2037
	}
L2035:
	;
	v14098 = float64(1e+100)
	goto L2036
L2036:
	;
	v14104 = *(*float64)(unsafe.Add(mBase, uint32(v13777)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v13777)+48)) = base.F64_add(base.F64_mul(base.F64_div(base.F64_sub(v13980, v13961), v13808), base.F64_add(v14098, float64(-1))), v14104)
	goto L2030
L2037:
	;
	v14093 = base.F64_nearest(v14080)
	if base.F64_gt(v14093, float64(1)) == int32(0) {
		goto L2030
	} else {
		goto L2038
	}
L2038:
	;
	v14098 = v14093
	goto L2036
L2039:
	;
	goto L1882
L2040:
	;
	goto L1841
L2041:
	;
	goto L1840
L2042:
	;
	v14268 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v14268 != 0 {
		goto L2046
	} else {
		goto L2047
	}
L2043:
	;
	v14259 = *(*int32)(unsafe.Add(mBase, uint32(v14217)+36))
	if v14259 == int32(0) {
		goto L2042
	} else {
		goto L2044
	}
L2044:
	;
	m.T0[v14259].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14224, int32(3), v13019, v14220, int32(0))
	mBase = m.M
	v14265 = m.ExcPending
	if v14265 != 0 {
		goto L1
	} else {
		goto L2045
	}
L2045:
	;
	goto L2042
L2046:
	;
	m.T0[v14268].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14224, int32(3), v13019, v14220, int32(0))
	mBase = m.M
	v14272 = m.ExcPending
	if v14272 != 0 {
		goto L1
	} else {
		goto L2049
	}
L2047:
	;
	goto L2048
L2048:
	;
	F_set_cheapest(m, v14220)
	mBase = m.M
	v14274 = m.ExcPending
	if v14274 != 0 {
		goto L1
	} else {
		goto L2050
	}
L2049:
	;
	goto L2048
L2050:
	;
	v14275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14229)+38)))
	if v14275 != int32(1) {
		v14286 = v14220
		v14287 = v14221
		v14288 = v14222
		v14290 = v14224
		v14293 = v14227
		v14295 = v14229
		v14302 = v14236
		v14305 = v14239
		v14307 = v14241
		v14316 = v14250
		v14320 = v14254
		v14321 = v14255
		goto L1827
	} else {
		goto L2051
	}
L2051:
	;
	v14278 = *(*int32)(unsafe.Add(mBase, uint32(v14227)+196))
	v14279 = *(*int32)(unsafe.Add(mBase, uint32(v14227)+192))
	F_adjust_paths_for_srfs(m, v14224, v14220, v14278, v14279)
	mBase = m.M
	v14281 = m.ExcPending
	if v14281 != 0 {
		goto L1
	} else {
		goto L2052
	}
L2052:
	;
	v14286 = v14220
	v14287 = v14221
	v14288 = v14222
	v14290 = v14224
	v14293 = v14227
	v14295 = v14229
	v14302 = v14236
	v14305 = v14239
	v14307 = v14241
	v14316 = v14250
	v14320 = v14254
	v14321 = v14255
	goto L1827
L2053:
	;
	v14910 = v14286
	v14913 = v14288
	v14915 = v14290
	v14918 = v14293
	v14920 = v14295
	v14927 = v14302
	v14930 = v14305
	v14932 = v14307
	v14941 = v14316
	v14945 = v14320
	v14946 = v14321
	goto L584
L2054:
	;
	goto L2055
L2055:
	;
	v14328 = F_fetch_upper_rel(m, v14290, int32(5), int32(0))
	mBase = m.M
	v14329 = m.ExcPending
	if v14329 != 0 {
		goto L1
	} else {
		goto L2056
	}
L2056:
	;
	v14330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14286)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14328)+26)) = uint8(v14330)
	v14332 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v14328)+164)) = v14332
	v14334 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v14328)+168)) = v14334
	v14336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14286)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14328)+172)) = uint8(v14336)
	v14338 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v14328)+176)) = v14338
	v14340 = F_create_final_distinct_paths(m, v14290, v14286, v14328)
	mBase = m.M
	v14341 = m.ExcPending
	if v14341 != 0 {
		goto L1
	} else {
		goto L2057
	}
L2057:
	;
	v14342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14286)+26)))
	if v14342 != int32(1) {
		goto L2058
	} else {
		goto L2059
	}
L2058:
	;
	v14885 = *(*int32)(unsafe.Add(mBase, uint32(v14340)+44))
	if v14885 == int32(0) {
		goto L583
	} else {
		goto L2183
	}
L2059:
	;
	v14345 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+52))
	if v14345 == int32(0) {
		goto L2058
	} else {
		goto L2060
	}
L2060:
	;
	v14348 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+4))
	v14349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14348)+40)))
	if v14349 != 0 {
		goto L2058
	} else {
		goto L2061
	}
L2061:
	;
	v14352 = F_fetch_upper_rel(m, v14290, int32(4), int32(0))
	mBase = m.M
	v14353 = m.ExcPending
	if v14353 != 0 {
		goto L1
	} else {
		goto L2062
	}
L2062:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14352)+40)) = v14287
	v14355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14286)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14352)+26)) = uint8(v14355)
	v14357 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v14352)+164)) = v14357
	v14359 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v14352)+168)) = v14359
	v14361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14286)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14352)+172)) = uint8(v14361)
	v14363 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v14352)+176)) = v14363
	v14365 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+52))
	v14366 = *(*int32)(unsafe.Add(mBase, uint32(v14365)+12))
	v14367 = *(*int32)(unsafe.Add(mBase, uint32(v14366)))
	v14368 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+280))
	v14369 = *(*int32)(unsafe.Add(mBase, uint32(v14348)+76))
	v14370 = F_get_sortgrouplist_exprs(m, v14368, v14369)
	mBase = m.M
	v14371 = m.ExcPending
	if v14371 != 0 {
		goto L1
	} else {
		goto L2063
	}
L2063:
	;
	v14372 = *(*float64)(unsafe.Add(mBase, uint32(v14367)+32))
	v14373 = int32(0)
	v14375 = F_estimate_num_groups(m, v14290, v14370, v14372, v14373, v14373)
	mBase = m.M
	v14376 = m.ExcPending
	if v14376 != 0 {
		goto L1
	} else {
		goto L2064
	}
L2064:
	;
	v14377 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+280))
	if v14377 == int32(0) {
		goto L2067
	} else {
		goto L2068
	}
L2065:
	;
	v14763 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[8])))
	if v14763 != int32(1) {
		goto L2153
	} else {
		goto L2154
	}
L2066:
	;
	if v14422 == int32(0) {
		goto L2065
	} else {
		goto L2079
	}
L2067:
	;
	v14422 = int32(1)
	goto L2066
L2068:
	;
	goto L2069
L2069:
	;
	v14386 = *(*int32)(unsafe.Add(mBase, uint32(v14377)+4))
	if v14386 <= int32(0) {
		v14414 = int32(1)
		goto L2070
	} else {
		goto L2071
	}
L2070:
	;
	v14422 = v14414
	goto L2066
L2071:
	;
	v14389 = int32(0)
	if v14389 < v14386 {
		goto L2072
	} else {
		goto L2073
	}
L2072:
	;
	v14392 = v14386
	goto L2074
L2073:
	;
	v14392 = v14389
	goto L2074
L2074:
	;
	v14393 = *(*int32)(unsafe.Add(mBase, uint32(v14377)+12))
	v14395 = int32(0)
	goto L2075
L2075:
	;
	v14403 = *(*int32)(unsafe.Add(mBase, uint32(v14393+v14395<<(uint(int32(2))%32))))
	v14404 = *(*int32)(unsafe.Add(mBase, uint32(v14403)+12))
	v14405 = int32(0)
	v14406 = base.B2i32(v14404 != v14405)
	if v14404 == v14405 {
		v14414 = v14406
		goto L2070
	} else {
		goto L2077
	}
L2076:
	;
	v14414 = v14406
	goto L2070
L2077:
	;
	v14410 = v14395 + int32(1)
	if v14410 != v14392 {
		v14395 = v14410
		goto L2075
	} else {
		goto L2078
	}
L2078:
	;
	goto L2076
L2079:
	;
	v14425 = *(*int32)(unsafe.Add(mBase, uint32(v14286)+52))
	if v14425 == int32(0) {
		goto L2065
	} else {
		goto L2080
	}
L2080:
	;
	v14428 = *(*int32)(unsafe.Add(mBase, uint32(v14425)+4))
	if v14428 <= int32(0) {
		goto L2065
	} else {
		goto L2081
	}
L2081:
	;
	v14446 = int32(0)
	goto L2082
L2082:
	;
	v14473 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+192))
	v14474 = *(*int32)(unsafe.Add(mBase, uint32(v14425)+12))
	v14478 = *(*int32)(unsafe.Add(mBase, uint32(v14474+v14446<<(uint(int32(2))%32))))
	v14479 = *(*int32)(unsafe.Add(mBase, uint32(v14478)+64))
	v14480 = F_get_useful_pathkeys_for_distinct(m, v14290, v14473, v14479)
	mBase = m.M
	v14481 = m.ExcPending
	if v14481 != 0 {
		goto L1
	} else {
		goto L2085
	}
L2083:
	;
	goto L2065
L2084:
	;
	v14718 = v14446 + int32(1)
	v14719 = *(*int32)(unsafe.Add(mBase, uint32(v14425)+4))
	if v14718 < v14719 {
		v14446 = v14718
		goto L2082
	} else {
		goto L2152
	}
L2085:
	;
	if v14480 == int32(0) {
		goto L2084
	} else {
		goto L2086
	}
L2086:
	;
	v14484 = int32(0)
	v14485 = *(*int32)(unsafe.Add(mBase, uint32(v14480)+4))
	if v14485 <= v14484 {
		goto L2084
	} else {
		goto L2087
	}
L2087:
	;
	v14489 = v14484
	goto L2088
L2088:
	;
	v14529 = *(*int32)(unsafe.Add(mBase, uint32(v14480)+12))
	v14533 = *(*int32)(unsafe.Add(mBase, uint32(v14529+v14489<<(uint(int32(2))%32))))
	v14534 = *(*int32)(unsafe.Add(mBase, uint32(v14478)+64))
	v14536 = v14293 + int32(208)
	if v14533 == v14534 {
		goto L2094
	} else {
		goto L2095
	}
L2089:
	;
	goto L2084
L2090:
	;
	v14673 = v14489 + int32(1)
	v14674 = *(*int32)(unsafe.Add(mBase, uint32(v14480)+4))
	if v14673 < v14674 {
		v14489 = v14673
		goto L2088
	} else {
		goto L2151
	}
L2091:
	;
	v14644 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+192))
	if v14644 == int32(0) {
		goto L2143
	} else {
		goto L2144
	}
L2092:
	;
	if v14614 != 0 {
		goto L2124
	} else {
		goto L2125
	}
L2093:
	;
	v14602 = *(*int32)(unsafe.Add(mBase, uint32(v14533)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14602
	v14614 = int32(1)
	goto L2092
L2094:
	;
	if v14533 != 0 {
		goto L2093
	} else {
		goto L2097
	}
L2095:
	;
	goto L2096
L2096:
	;
	if v14533 == int32(0) {
		goto L2098
	} else {
		goto L2099
	}
L2097:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = int32(0)
	v14614 = int32(1)
	goto L2092
L2098:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = int32(0)
	v14614 = int32(1)
	goto L2092
L2099:
	;
	goto L2100
L2100:
	;
	if v14534 == int32(0) {
		goto L2101
	} else {
		goto L2102
	}
L2101:
	;
	v14554 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14554
	v14614 = v14554
	goto L2092
L2102:
	;
	goto L2103
L2103:
	;
	v14557 = *(*int32)(unsafe.Add(mBase, uint32(v14534)+4))
	v14558 = int32(0)
	if v14558 < v14557 {
		goto L2104
	} else {
		goto L2105
	}
L2104:
	;
	v14561 = v14557
	goto L2106
L2105:
	;
	v14561 = v14558
	goto L2106
L2106:
	;
	v14562 = *(*int32)(unsafe.Add(mBase, uint32(v14533)+4))
	v14567 = int32(0)
	goto L2107
L2107:
	;
	if v14567 < v14562 {
		goto L2109
	} else {
		goto L2110
	}
L2109:
	;
	v14574 = *(*int32)(unsafe.Add(mBase, uint32(v14533)+12))
	v14578 = v14574 + v14567<<(uint(int32(2))%32)
	goto L2111
L2110:
	;
	v14578 = int32(0)
	goto L2111
L2111:
	;
	if v14567 == v14561 {
		goto L2112
	} else {
		goto L2113
	}
L2112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14561
	v14614 = base.B2i32(v14578 == int32(0))
	goto L2092
L2113:
	;
	goto L2114
L2114:
	;
	v14584 = base.B2i32(v14578 == int32(0))
	if v14578 == int32(0) {
		goto L2115
	} else {
		goto L2116
	}
L2115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14567
	v14614 = v14584
	goto L2092
L2116:
	;
	goto L2117
L2117:
	;
	v14588 = *(*int32)(unsafe.Add(mBase, uint32(v14534)+12))
	if v14588 == int32(0) {
		goto L2118
	} else {
		goto L2119
	}
L2118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14567
	v14614 = v14584
	goto L2092
L2119:
	;
	goto L2120
L2120:
	;
	v14592 = *(*int32)(unsafe.Add(mBase, uint32(v14578)))
	v14596 = *(*int32)(unsafe.Add(mBase, uint32(v14588+v14567<<(uint(int32(2))%32))))
	if v14592 != v14596 {
		goto L2121
	} else {
		goto L2122
	}
L2121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14536))) = v14567
	v14614 = int32(0)
	goto L2092
L2122:
	;
	v14567 = v14567 + int32(1)
	goto L2107
L2124:
	;
	v14641 = v14478
	goto L2091
L2125:
	;
	goto L2126
L2126:
	;
	v14616 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v14617 = *(*int32)(unsafe.Add(mBase, uint32(v14293)+208))
	if v14478 == v14367 {
		goto L2128
	} else {
		goto L2129
	}
L2127:
	;
	if v14626&int32(1) != 0 {
		goto L2133
	} else {
		goto L2134
	}
L2128:
	;
	v14626 = v14616
	goto L2127
L2129:
	;
	goto L2130
L2130:
	;
	if v14617 == int32(0) {
		goto L2090
	} else {
		goto L2131
	}
L2131:
	;
	v14621 = int32(1)
	if v14616&v14621 == int32(0) {
		goto L2090
	} else {
		goto L2132
	}
L2132:
	;
	v14626 = v14621
	goto L2127
L2133:
	;
	v14630 = v14617
	goto L2135
L2134:
	;
	v14630 = int32(0)
	goto L2135
L2135:
	;
	if v14630 == int32(0) {
		goto L2136
	} else {
		goto L2137
	}
L2136:
	;
	v14634 = F_create_sort_path(m, v14352, v14478, v14533, float64(-1))
	mBase = m.M
	v14635 = m.ExcPending
	if v14635 != 0 {
		goto L1
	} else {
		goto L2139
	}
L2137:
	;
	goto L2138
L2138:
	;
	v14637 = F_create_incremental_sort_path(m, v14290, v14352, v14478, v14533, v14617, float64(-1))
	mBase = m.M
	v14638 = m.ExcPending
	if v14638 != 0 {
		goto L1
	} else {
		goto L2141
	}
L2139:
	;
	if v14634 != 0 {
		v14641 = v14634
		goto L2091
	} else {
		goto L2140
	}
L2140:
	;
	goto L2090
L2141:
	;
	if v14637 == int32(0) {
		goto L2090
	} else {
		goto L2142
	}
L2142:
	;
	v14641 = v14637
	goto L2091
L2143:
	;
	v14647 = int32(0)
	v14655 = F_makeConst(m, int32(20), int32(-1), v14647, int32(8), int64(1), v14647, int32(1))
	mBase = m.M
	v14656 = m.ExcPending
	if v14656 != 0 {
		goto L1
	} else {
		goto L2146
	}
L2144:
	;
	goto L2145
L2145:
	;
	v14664 = *(*int32)(unsafe.Add(mBase, uint32(v14644)+4))
	v14665 = F_create_unique_path(m, v14352, v14641, v14664, v14375)
	mBase = m.M
	v14666 = m.ExcPending
	if v14666 != 0 {
		goto L1
	} else {
		goto L2149
	}
L2146:
	;
	v14660 = F_create_limit_path(m, v14352, v14641, v14647, v14655, int32(0), int64(0), int64(1))
	mBase = m.M
	v14661 = m.ExcPending
	if v14661 != 0 {
		goto L1
	} else {
		goto L2147
	}
L2147:
	;
	F_add_partial_path(m, v14352, v14660)
	mBase = m.M
	v14663 = m.ExcPending
	if v14663 != 0 {
		goto L1
	} else {
		goto L2148
	}
L2148:
	;
	goto L2090
L2149:
	;
	F_add_partial_path(m, v14352, v14665)
	mBase = m.M
	v14668 = m.ExcPending
	if v14668 != 0 {
		goto L1
	} else {
		goto L2150
	}
L2150:
	;
	goto L2090
L2151:
	;
	goto L2089
L2152:
	;
	goto L2083
L2153:
	;
	v14817 = *(*int32)(unsafe.Add(mBase, uint32(v14352)+176))
	if v14817 == int32(0) {
		goto L2171
	} else {
		goto L2172
	}
L2154:
	;
	v14766 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+280))
	v14767 = int32(0)
	if v14766 == v14767 {
		goto L2156
	} else {
		goto L2157
	}
L2155:
	;
	if v14804 == int32(0) {
		goto L2153
	} else {
		goto L2168
	}
L2156:
	;
	v14804 = int32(1)
	goto L2155
L2157:
	;
	goto L2158
L2158:
	;
	v14774 = *(*int32)(unsafe.Add(mBase, uint32(v14766)+4))
	if v14774 <= int32(0) {
		v14798 = int32(1)
		goto L2159
	} else {
		goto L2160
	}
L2159:
	;
	v14804 = v14798
	goto L2155
L2160:
	;
	v14777 = int32(0)
	if v14777 < v14774 {
		goto L2161
	} else {
		goto L2162
	}
L2161:
	;
	v14780 = v14774
	goto L2163
L2162:
	;
	v14780 = v14777
	goto L2163
L2163:
	;
	v14781 = *(*int32)(unsafe.Add(mBase, uint32(v14766)+12))
	v14785 = v14767
	goto L2164
L2164:
	;
	v14789 = *(*int32)(unsafe.Add(mBase, uint32(v14781+v14785<<(uint(int32(2))%32))))
	v14790 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14789)+18)))
	if v14790 != int32(1) {
		v14798 = v14790
		goto L2159
	} else {
		goto L2166
	}
L2165:
	;
	v14798 = v14790
	goto L2159
L2166:
	;
	v14794 = v14785 + int32(1)
	if v14794 != v14780 {
		v14785 = v14794
		goto L2164
	} else {
		goto L2167
	}
L2167:
	;
	goto L2165
L2168:
	;
	v14807 = *(*int32)(unsafe.Add(mBase, uint32(v14367)+12))
	v14809 = int32(0)
	v14810 = *(*int32)(unsafe.Add(mBase, uint32(v14290)+280))
	v14813 = F_create_agg_path(m, v14290, v14352, v14367, v14807, int32(2), v14809, v14810, v14809, v14809, v14375)
	mBase = m.M
	v14814 = m.ExcPending
	if v14814 != 0 {
		goto L1
	} else {
		goto L2169
	}
L2169:
	;
	F_add_partial_path(m, v14352, v14813)
	mBase = m.M
	v14816 = m.ExcPending
	if v14816 != 0 {
		goto L1
	} else {
		goto L2170
	}
L2170:
	;
	goto L2153
L2171:
	;
	v14829 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v14829 != 0 {
		goto L2175
	} else {
		goto L2176
	}
L2172:
	;
	v14820 = *(*int32)(unsafe.Add(mBase, uint32(v14817)+36))
	if v14820 == int32(0) {
		goto L2171
	} else {
		goto L2173
	}
L2173:
	;
	m.T0[v14820].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14290, int32(4), v14286, v14352, int32(0))
	mBase = m.M
	v14826 = m.ExcPending
	if v14826 != 0 {
		goto L1
	} else {
		goto L2174
	}
L2174:
	;
	goto L2171
L2175:
	;
	m.T0[v14829].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14290, int32(4), v14286, v14352, int32(0))
	mBase = m.M
	v14833 = m.ExcPending
	if v14833 != 0 {
		goto L1
	} else {
		goto L2178
	}
L2176:
	;
	goto L2177
L2177:
	;
	v14834 = *(*int32)(unsafe.Add(mBase, uint32(v14352)+52))
	if v14834 == int32(0) {
		goto L2058
	} else {
		goto L2179
	}
L2178:
	;
	goto L2177
L2179:
	;
	F_generate_useful_gather_paths(m, v14290, v14352, int32(1))
	mBase = m.M
	v14839 = m.ExcPending
	if v14839 != 0 {
		goto L1
	} else {
		goto L2180
	}
L2180:
	;
	F_set_cheapest(m, v14352)
	mBase = m.M
	v14841 = m.ExcPending
	if v14841 != 0 {
		goto L1
	} else {
		goto L2181
	}
L2181:
	;
	v14842 = F_create_final_distinct_paths(m, v14290, v14352, v14340)
	mBase = m.M
	v14843 = m.ExcPending
	if v14843 != 0 {
		goto L1
	} else {
		goto L2182
	}
L2182:
	;
	goto L2058
L2183:
	;
	v14888 = *(*int32)(unsafe.Add(mBase, uint32(v14328)+176))
	if v14888 == int32(0) {
		goto L2184
	} else {
		goto L2185
	}
L2184:
	;
	v14900 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v14900 != 0 {
		goto L2188
	} else {
		goto L2189
	}
L2185:
	;
	v14891 = *(*int32)(unsafe.Add(mBase, uint32(v14888)+36))
	if v14891 == int32(0) {
		goto L2184
	} else {
		goto L2186
	}
L2186:
	;
	m.T0[v14891].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14290, int32(5), v14286, v14340, int32(0))
	mBase = m.M
	v14897 = m.ExcPending
	if v14897 != 0 {
		goto L1
	} else {
		goto L2187
	}
L2187:
	;
	goto L2184
L2188:
	;
	m.T0[v14900].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14290, int32(5), v14286, v14340, int32(0))
	mBase = m.M
	v14904 = m.ExcPending
	if v14904 != 0 {
		goto L1
	} else {
		goto L2191
	}
L2189:
	;
	goto L2190
L2190:
	;
	F_set_cheapest(m, v14340)
	mBase = m.M
	v14906 = m.ExcPending
	if v14906 != 0 {
		goto L1
	} else {
		goto L2192
	}
L2191:
	;
	goto L2190
L2192:
	;
	v14910 = v14340
	v14913 = v14288
	v14915 = v14290
	v14918 = v14293
	v14920 = v14295
	v14927 = v14302
	v14930 = v14305
	v14932 = v14307
	v14941 = v14316
	v14945 = v14320
	v14946 = v14321
	goto L584
L2193:
	;
	v15528 = F_fetch_upper_rel(m, v14915, int32(7), int32(0))
	mBase = m.M
	v15529 = m.ExcPending
	if v15529 != 0 {
		goto L1
	} else {
		goto L2346
	}
L2194:
	;
	v15489 = v14910
	goto L2193
L2195:
	;
	goto L2196
L2196:
	;
	v14951 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+60))
	v14954 = F_fetch_upper_rel(m, v14915, int32(6), int32(0))
	mBase = m.M
	v14955 = m.ExcPending
	if v14955 != 0 {
		goto L1
	} else {
		goto L2197
	}
L2197:
	;
	v14956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14910)+26)))
	if v14932&v14956 == int32(1) {
		goto L2198
	} else {
		goto L2199
	}
L2198:
	;
	v14960 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14954)+26)) = uint8(v14960)
	goto L2200
L2199:
	;
	goto L2200
L2200:
	;
	v14962 = *(*int64)(unsafe.Add(mBase, uint32(v14910)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v14954)+32)) = v14962
	v14964 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v14954)+164)) = v14964
	v14966 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v14954)+168)) = v14966
	v14968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14910)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14954)+172)) = uint8(v14968)
	v14970 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v14954)+176)) = v14970
	v14972 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+44))
	if v14972 == int32(0) {
		goto L2201
	} else {
		goto L2202
	}
L2201:
	;
	v15193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14954)+26)))
	if v15193 == int32(0) {
		goto L2262
	} else {
		goto L2263
	}
L2202:
	;
	v14975 = *(*int32)(unsafe.Add(mBase, uint32(v14972)+4))
	if v14975 <= int32(0) {
		goto L2201
	} else {
		goto L2203
	}
L2203:
	;
	v14988 = int32(0)
	goto L2204
L2204:
	;
	v15020 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	v15021 = *(*int32)(unsafe.Add(mBase, uint32(v14972)+12))
	v15025 = *(*int32)(unsafe.Add(mBase, uint32(v15021+v14988<<(uint(int32(2))%32))))
	v15026 = *(*int32)(unsafe.Add(mBase, uint32(v15025)+64))
	v15028 = v14918 + int32(208)
	if v15020 == v15026 {
		goto L2210
	} else {
		goto L2211
	}
L2205:
	;
	goto L2201
L2206:
	;
	v15149 = v14988 + int32(1)
	v15150 = *(*int32)(unsafe.Add(mBase, uint32(v14972)+4))
	if v15149 < v15150 {
		v14988 = v15149
		goto L2204
	} else {
		goto L2261
	}
L2207:
	;
	v15134 = *(*int32)(unsafe.Add(mBase, uint32(v15131)+12))
	v15135 = *(*int32)(unsafe.Add(mBase, uint32(v15134)+4))
	v15136 = *(*int32)(unsafe.Add(mBase, uint32(v14927)+4))
	v15137 = F_equal(m, v15135, v15136)
	mBase = m.M
	v15138 = m.ExcPending
	if v15138 != 0 {
		goto L1
	} else {
		goto L2255
	}
L2208:
	;
	if v15106 != 0 {
		v15131 = v15025
		goto L2207
	} else {
		goto L2240
	}
L2209:
	;
	v15094 = *(*int32)(unsafe.Add(mBase, uint32(v15020)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15094
	v15106 = int32(1)
	goto L2208
L2210:
	;
	if v15020 != 0 {
		goto L2209
	} else {
		goto L2213
	}
L2211:
	;
	goto L2212
L2212:
	;
	if v15020 == int32(0) {
		goto L2214
	} else {
		goto L2215
	}
L2213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = int32(0)
	v15106 = int32(1)
	goto L2208
L2214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = int32(0)
	v15106 = int32(1)
	goto L2208
L2215:
	;
	goto L2216
L2216:
	;
	if v15026 == int32(0) {
		goto L2217
	} else {
		goto L2218
	}
L2217:
	;
	v15046 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15046
	v15106 = v15046
	goto L2208
L2218:
	;
	goto L2219
L2219:
	;
	v15049 = *(*int32)(unsafe.Add(mBase, uint32(v15026)+4))
	v15050 = int32(0)
	if v15050 < v15049 {
		goto L2220
	} else {
		goto L2221
	}
L2220:
	;
	v15053 = v15049
	goto L2222
L2221:
	;
	v15053 = v15050
	goto L2222
L2222:
	;
	v15054 = *(*int32)(unsafe.Add(mBase, uint32(v15020)+4))
	v15059 = int32(0)
	goto L2223
L2223:
	;
	if v15059 < v15054 {
		goto L2225
	} else {
		goto L2226
	}
L2225:
	;
	v15066 = *(*int32)(unsafe.Add(mBase, uint32(v15020)+12))
	v15070 = v15066 + v15059<<(uint(int32(2))%32)
	goto L2227
L2226:
	;
	v15070 = int32(0)
	goto L2227
L2227:
	;
	if v15059 == v15053 {
		goto L2228
	} else {
		goto L2229
	}
L2228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15053
	v15106 = base.B2i32(v15070 == int32(0))
	goto L2208
L2229:
	;
	goto L2230
L2230:
	;
	v15076 = base.B2i32(v15070 == int32(0))
	if v15070 == int32(0) {
		goto L2231
	} else {
		goto L2232
	}
L2231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15059
	v15106 = v15076
	goto L2208
L2232:
	;
	goto L2233
L2233:
	;
	v15080 = *(*int32)(unsafe.Add(mBase, uint32(v15026)+12))
	if v15080 == int32(0) {
		goto L2234
	} else {
		goto L2235
	}
L2234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15059
	v15106 = v15076
	goto L2208
L2235:
	;
	goto L2236
L2236:
	;
	v15084 = *(*int32)(unsafe.Add(mBase, uint32(v15070)))
	v15088 = *(*int32)(unsafe.Add(mBase, uint32(v15080+v15059<<(uint(int32(2))%32))))
	if v15084 != v15088 {
		goto L2237
	} else {
		goto L2238
	}
L2237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15028))) = v15059
	v15106 = int32(0)
	goto L2208
L2238:
	;
	v15059 = v15059 + int32(1)
	goto L2223
L2240:
	;
	v15108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v15109 = *(*int32)(unsafe.Add(mBase, uint32(v14918)+208))
	if v15025 == v14951 {
		goto L2242
	} else {
		goto L2243
	}
L2241:
	;
	v15119 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	if v15118&int32(1) != 0 {
		goto L2247
	} else {
		goto L2248
	}
L2242:
	;
	v15118 = v15108
	goto L2241
L2243:
	;
	goto L2244
L2244:
	;
	if v15109 == int32(0) {
		goto L2206
	} else {
		goto L2245
	}
L2245:
	;
	v15113 = int32(1)
	if v15108&v15113 == int32(0) {
		goto L2206
	} else {
		goto L2246
	}
L2246:
	;
	v15118 = v15113
	goto L2241
L2247:
	;
	v15123 = v15109
	goto L2249
L2248:
	;
	v15123 = int32(0)
	goto L2249
L2249:
	;
	if v15123 == int32(0) {
		goto L2250
	} else {
		goto L2251
	}
L2250:
	;
	v15126 = F_create_sort_path(m, v14954, v15025, v15119, v14913)
	mBase = m.M
	v15127 = m.ExcPending
	if v15127 != 0 {
		goto L1
	} else {
		goto L2253
	}
L2251:
	;
	goto L2252
L2252:
	;
	v15128 = F_create_incremental_sort_path(m, v14915, v14954, v15025, v15119, v15109, v14913)
	mBase = m.M
	v15129 = m.ExcPending
	if v15129 != 0 {
		goto L1
	} else {
		goto L2254
	}
L2253:
	;
	v15131 = v15126
	goto L2207
L2254:
	;
	v15131 = v15128
	goto L2207
L2255:
	;
	if v15137 != 0 {
		goto L2256
	} else {
		goto L2257
	}
L2256:
	;
	v15141 = v15131
	goto L2258
L2257:
	;
	v15139 = F_apply_projection_to_path(m, v14915, v14954, v15131, v14927)
	mBase = m.M
	v15140 = m.ExcPending
	if v15140 != 0 {
		goto L1
	} else {
		goto L2259
	}
L2258:
	;
	F_add_path(m, v14954, v15141)
	mBase = m.M
	v15143 = m.ExcPending
	if v15143 != 0 {
		goto L1
	} else {
		goto L2260
	}
L2259:
	;
	v15141 = v15139
	goto L2258
L2260:
	;
	goto L2206
L2261:
	;
	goto L2205
L2262:
	;
	v15461 = *(*int32)(unsafe.Add(mBase, uint32(v14954)+176))
	if v15461 == int32(0) {
		goto L2336
	} else {
		goto L2337
	}
L2263:
	;
	v15196 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	if v15196 == int32(0) {
		goto L2262
	} else {
		goto L2264
	}
L2264:
	;
	v15199 = *(*int32)(unsafe.Add(mBase, uint32(v14910)+52))
	if v15199 == int32(0) {
		goto L2262
	} else {
		goto L2265
	}
L2265:
	;
	v15202 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+4))
	if v15202 <= int32(0) {
		goto L2262
	} else {
		goto L2266
	}
L2266:
	;
	v15205 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+12))
	v15206 = *(*int32)(unsafe.Add(mBase, uint32(v15205)))
	v15209 = int32(0)
	goto L2267
L2267:
	;
	v15249 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	v15250 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+12))
	v15254 = *(*int32)(unsafe.Add(mBase, uint32(v15250+v15209<<(uint(int32(2))%32))))
	v15255 = *(*int32)(unsafe.Add(mBase, uint32(v15254)+64))
	v15257 = v14918 + int32(320)
	if v15249 == v15255 {
		goto L2272
	} else {
		goto L2273
	}
L2268:
	;
	goto L2262
L2269:
	;
	v15417 = v15209 + int32(1)
	v15418 = *(*int32)(unsafe.Add(mBase, uint32(v15199)+4))
	if v15417 < v15418 {
		v15209 = v15417
		goto L2267
	} else {
		goto L2335
	}
L2270:
	;
	if v15335 != 0 {
		goto L2269
	} else {
		goto L2302
	}
L2271:
	;
	v15323 = *(*int32)(unsafe.Add(mBase, uint32(v15249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15323
	v15335 = int32(1)
	goto L2270
L2272:
	;
	if v15249 != 0 {
		goto L2271
	} else {
		goto L2275
	}
L2273:
	;
	goto L2274
L2274:
	;
	if v15249 == int32(0) {
		goto L2276
	} else {
		goto L2277
	}
L2275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = int32(0)
	v15335 = int32(1)
	goto L2270
L2276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = int32(0)
	v15335 = int32(1)
	goto L2270
L2277:
	;
	goto L2278
L2278:
	;
	if v15255 == int32(0) {
		goto L2279
	} else {
		goto L2280
	}
L2279:
	;
	v15275 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15275
	v15335 = v15275
	goto L2270
L2280:
	;
	goto L2281
L2281:
	;
	v15278 = *(*int32)(unsafe.Add(mBase, uint32(v15255)+4))
	v15279 = int32(0)
	if v15279 < v15278 {
		goto L2282
	} else {
		goto L2283
	}
L2282:
	;
	v15282 = v15278
	goto L2284
L2283:
	;
	v15282 = v15279
	goto L2284
L2284:
	;
	v15283 = *(*int32)(unsafe.Add(mBase, uint32(v15249)+4))
	v15288 = int32(0)
	goto L2285
L2285:
	;
	if v15288 < v15283 {
		goto L2287
	} else {
		goto L2288
	}
L2287:
	;
	v15295 = *(*int32)(unsafe.Add(mBase, uint32(v15249)+12))
	v15299 = v15295 + v15288<<(uint(int32(2))%32)
	goto L2289
L2288:
	;
	v15299 = int32(0)
	goto L2289
L2289:
	;
	if v15288 == v15282 {
		goto L2290
	} else {
		goto L2291
	}
L2290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15282
	v15335 = base.B2i32(v15299 == int32(0))
	goto L2270
L2291:
	;
	goto L2292
L2292:
	;
	v15305 = base.B2i32(v15299 == int32(0))
	if v15299 == int32(0) {
		goto L2293
	} else {
		goto L2294
	}
L2293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15288
	v15335 = v15305
	goto L2270
L2294:
	;
	goto L2295
L2295:
	;
	v15309 = *(*int32)(unsafe.Add(mBase, uint32(v15255)+12))
	if v15309 == int32(0) {
		goto L2296
	} else {
		goto L2297
	}
L2296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15288
	v15335 = v15305
	goto L2270
L2297:
	;
	goto L2298
L2298:
	;
	v15313 = *(*int32)(unsafe.Add(mBase, uint32(v15299)))
	v15317 = *(*int32)(unsafe.Add(mBase, uint32(v15309+v15288<<(uint(int32(2))%32))))
	if v15313 != v15317 {
		goto L2299
	} else {
		goto L2300
	}
L2299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15257))) = v15288
	v15335 = int32(0)
	goto L2270
L2300:
	;
	v15288 = v15288 + int32(1)
	goto L2285
L2302:
	;
	v15337 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[7])))
	v15338 = *(*int32)(unsafe.Add(mBase, uint32(v14918)+320))
	if v15254 == v15206 {
		goto L2304
	} else {
		goto L2305
	}
L2303:
	;
	v15348 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	if v15347&int32(1) != 0 {
		goto L2310
	} else {
		goto L2311
	}
L2304:
	;
	v15347 = v15337
	goto L2303
L2305:
	;
	goto L2306
L2306:
	;
	if v15338 == int32(0) {
		goto L2269
	} else {
		goto L2307
	}
L2307:
	;
	v15342 = int32(1)
	if v15337&v15342 == int32(0) {
		goto L2269
	} else {
		goto L2308
	}
L2308:
	;
	v15347 = v15342
	goto L2303
L2309:
	;
	v15363 = *(*float64)(unsafe.Add(mBase, uint32(v15359)+32))
	v15364 = *(*int32)(unsafe.Add(mBase, uint32(v15359)+24))
	v15365 = base.F64_convert_i32_s(v15364)
	v15367 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_subquery_planner[9])))
	if v15367 == int32(1) {
		goto L2319
	} else {
		goto L2320
	}
L2310:
	;
	v15352 = v15338
	goto L2312
L2311:
	;
	v15352 = int32(0)
	goto L2312
L2312:
	;
	if v15352 == int32(0) {
		goto L2313
	} else {
		goto L2314
	}
L2313:
	;
	v15355 = F_create_sort_path(m, v14954, v15254, v15348, v14913)
	mBase = m.M
	v15356 = m.ExcPending
	if v15356 != 0 {
		goto L1
	} else {
		goto L2316
	}
L2314:
	;
	goto L2315
L2315:
	;
	v15357 = F_create_incremental_sort_path(m, v14915, v14954, v15254, v15348, v15338, v14913)
	mBase = m.M
	v15358 = m.ExcPending
	if v15358 != 0 {
		goto L1
	} else {
		goto L2317
	}
L2316:
	;
	v15359 = v15355
	goto L2309
L2317:
	;
	v15359 = v15357
	goto L2309
L2318:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14918)+208)) = v15395
	v15397 = *(*int32)(unsafe.Add(mBase, uint32(v15359)+12))
	v15398 = *(*int32)(unsafe.Add(mBase, uint32(v14915)+196))
	v15401 = F_create_gather_merge_path(m, v14915, v14954, v15359, v15397, v15398, v14918+int32(208))
	mBase = m.M
	v15402 = m.ExcPending
	if v15402 != 0 {
		goto L1
	} else {
		goto L2328
	}
L2319:
	;
	v15373 = base.F64_add(base.F64_mul(v15365, float64(-0.3)), float64(1))
	if base.F64_gt(v15373, float64(0)) != 0 {
		goto L2322
	} else {
		goto L2323
	}
L2320:
	;
	v15379 = v15365
	goto L2321
L2321:
	;
	v15381 = float64(1e+100)
	v15382 = base.F64_mul(v15363, v15379)
	if base.F64_gt(v15382, v15381)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v15382)&int64(9223372036854775807))) != 0 {
		v15395 = v15381
		goto L2325
	} else {
		goto L2326
	}
L2322:
	;
	v15377 = v15373
	goto L2324
L2323:
	;
	v15377 = math.Float64frombits(uint64(0x8000000000000000))
	goto L2324
L2324:
	;
	v15379 = base.F64_add(v15377, v15365)
	goto L2321
L2325:
	;
	goto L2318
L2326:
	;
	v15391 = float64(1)
	if base.F64_le(v15382, v15391) != 0 {
		v15395 = v15391
		goto L2325
	} else {
		goto L2327
	}
L2327:
	;
	v15395 = base.F64_nearest(v15382)
	goto L2325
L2328:
	;
	v15403 = *(*int32)(unsafe.Add(mBase, uint32(v15401)+12))
	v15404 = *(*int32)(unsafe.Add(mBase, uint32(v15403)+4))
	v15405 = *(*int32)(unsafe.Add(mBase, uint32(v14927)+4))
	v15406 = F_equal(m, v15404, v15405)
	mBase = m.M
	v15407 = m.ExcPending
	if v15407 != 0 {
		goto L1
	} else {
		goto L2329
	}
L2329:
	;
	if v15406 != 0 {
		goto L2330
	} else {
		goto L2331
	}
L2330:
	;
	v15410 = v15401
	goto L2332
L2331:
	;
	v15408 = F_apply_projection_to_path(m, v14915, v14954, v15401, v14927)
	mBase = m.M
	v15409 = m.ExcPending
	if v15409 != 0 {
		goto L1
	} else {
		goto L2333
	}
L2332:
	;
	F_add_path(m, v14954, v15410)
	mBase = m.M
	v15412 = m.ExcPending
	if v15412 != 0 {
		goto L1
	} else {
		goto L2334
	}
L2333:
	;
	v15410 = v15408
	goto L2332
L2334:
	;
	goto L2269
L2335:
	;
	goto L2268
L2336:
	;
	v15473 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v15473 != 0 {
		goto L2340
	} else {
		goto L2341
	}
L2337:
	;
	v15464 = *(*int32)(unsafe.Add(mBase, uint32(v15461)+36))
	if v15464 == int32(0) {
		goto L2336
	} else {
		goto L2338
	}
L2338:
	;
	m.T0[v15464].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14915, int32(6), v14910, v14954, int32(0))
	mBase = m.M
	v15470 = m.ExcPending
	if v15470 != 0 {
		goto L1
	} else {
		goto L2339
	}
L2339:
	;
	goto L2336
L2340:
	;
	m.T0[v15473].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14915, int32(6), v14910, v14954, int32(0))
	mBase = m.M
	v15477 = m.ExcPending
	if v15477 != 0 {
		goto L1
	} else {
		goto L2343
	}
L2341:
	;
	goto L2342
L2342:
	;
	v15478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14920)+38)))
	if v15478 != int32(1) {
		v15489 = v14954
		goto L2193
	} else {
		goto L2344
	}
L2343:
	;
	goto L2342
L2344:
	;
	v15481 = *(*int32)(unsafe.Add(mBase, uint32(v14918)+204))
	v15482 = *(*int32)(unsafe.Add(mBase, uint32(v14918)+200))
	F_adjust_paths_for_srfs(m, v14915, v14954, v15481, v15482)
	mBase = m.M
	v15484 = m.ExcPending
	if v15484 != 0 {
		goto L1
	} else {
		goto L2345
	}
L2345:
	;
	v15489 = v14954
	goto L2193
L2346:
	;
	v15530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15489)+26)))
	if v15530 != int32(1) {
		goto L2347
	} else {
		goto L2348
	}
L2347:
	;
	v15545 = *(*int32)(unsafe.Add(mBase, uint32(v15489)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v15528)+164)) = v15545
	v15547 = *(*int32)(unsafe.Add(mBase, uint32(v15489)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v15528)+168)) = v15547
	v15549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15489)+172)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15528)+172)) = uint8(v15549)
	v15551 = *(*int32)(unsafe.Add(mBase, uint32(v15489)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v15528)+176)) = v15551
	v15553 = *(*int32)(unsafe.Add(mBase, uint32(v15489)+44))
	if v15553 == int32(0) {
		v16510 = v15528
		v16514 = v15489
		v16518 = v14915
		v16521 = v14918
		v16523 = v14920
		v16533 = v14930
		v16544 = v14941
		v16548 = v14945
		v16549 = v14946
		goto L2353
	} else {
		goto L2354
	}
L2348:
	;
	v15533 = *(*int32)(unsafe.Add(mBase, uint32(v14920)+128))
	v15534 = F_is_parallel_safe(m, v14915, v15533)
	mBase = m.M
	v15535 = m.ExcPending
	if v15535 != 0 {
		goto L1
	} else {
		goto L2349
	}
L2349:
	;
	if v15534 == int32(0) {
		goto L2347
	} else {
		goto L2350
	}
L2350:
	;
	v15538 = *(*int32)(unsafe.Add(mBase, uint32(v14920)+132))
	v15539 = F_is_parallel_safe(m, v14915, v15538)
	mBase = m.M
	v15540 = m.ExcPending
	if v15540 != 0 {
		goto L1
	} else {
		goto L2351
	}
L2351:
	;
	if v15539 == int32(0) {
		goto L2347
	} else {
		goto L2352
	}
L2352:
	;
	v15543 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15528)+26)) = uint8(v15543)
	goto L2347
L2353:
	;
	v16551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16510)+26)))
	if v16551 == int32(0) {
		goto L2542
	} else {
		goto L2543
	}
L2354:
	;
	v15556 = *(*int32)(unsafe.Add(mBase, uint32(v15553)+4))
	if v15556 <= int32(0) {
		v16510 = v15528
		v16514 = v15489
		v16518 = v14915
		v16521 = v14918
		v16523 = v14920
		v16533 = v14930
		v16544 = v14941
		v16548 = v14945
		v16549 = v14946
		goto L2353
	} else {
		goto L2355
	}
L2355:
	;
	v15560 = v15528
	v15564 = v15489
	v15568 = v14915
	v15570 = v15553
	v15571 = v14918
	v15573 = v14920
	v15582 = int32(0)
	v15583 = v14930
	v15594 = v14941
	v15598 = v14945
	v15599 = v14946
	goto L2356
L2356:
	;
	v15601 = *(*int32)(unsafe.Add(mBase, uint32(v15570)+12))
	v15605 = *(*int32)(unsafe.Add(mBase, uint32(v15601+v15582<<(uint(int32(2))%32))))
	v15606 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+140))
	if v15606 != 0 {
		goto L2358
	} else {
		goto L2359
	}
L2357:
	;
	v16510 = v16462
	v16514 = v16466
	v16518 = v16470
	v16521 = v16473
	v16523 = v16475
	v16533 = v16485
	v16544 = v16496
	v16548 = v16500
	v16549 = v16501
	goto L2353
L2358:
	;
	v15607 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+144))
	v15608 = F_assign_special_exec_param(m, v15568)
	mBase = m.M
	v15609 = m.ExcPending
	if v15609 != 0 {
		goto L1
	} else {
		goto L2361
	}
L2359:
	;
	v15647 = v15605
	goto L2360
L2360:
	;
	v15648 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+132))
	if v15648 != 0 {
		goto L2365
	} else {
		goto L2366
	}
L2361:
	;
	v15611 = F_palloc0(m, int32(88))
	mBase = m.M
	v15612 = m.ExcPending
	if v15612 != 0 {
		goto L1
	} else {
		goto L2362
	}
L2362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+8)) = v15560
	*(*int64)(unsafe.Add(mBase, uint32(v15611))) = int64(1614907703613)
	v15616 = *(*int32)(unsafe.Add(mBase, uint32(v15605)+12))
	v15617 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+24)) = v15617
	*(*uint16)(unsafe.Add(mBase, uint32(v15611)+20)) = uint16(v15617)
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+16)) = v15617
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+12)) = v15616
	v15624 = *(*float64)(unsafe.Add(mBase, uint32(v15605)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+80)) = v15608
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+76)) = v15607
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+72)) = v15605
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+64)) = v15617
	*(*float64)(unsafe.Add(mBase, uint32(v15611)+32)) = v15624
	v15631 = *(*int32)(unsafe.Add(mBase, uint32(v15605)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v15611)+40)) = v15631
	v15633 = *(*float64)(unsafe.Add(mBase, uint32(v15605)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v15611)+48)) = v15633
	v15636 = *(*float64)(unsafe.Add(mBase, _c_F_subquery_planner[1]))
	v15637 = *(*float64)(unsafe.Add(mBase, uint32(v15605)+32))
	v15639 = *(*float64)(unsafe.Add(mBase, uint32(v15605)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v15611)+56)) = base.F64_add(base.F64_mul(v15636, v15637), v15639)
	v15647 = v15611
	goto L2360
L2363:
	;
	v15672 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v15672 != int32(1) {
		goto L2375
	} else {
		goto L2376
	}
L2364:
	;
	v15666 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+128))
	v15667 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+136))
	v15668 = F_create_limit_path(m, v15560, v15647, v15666, v15648, v15667, v15598, v15599)
	mBase = m.M
	v15669 = m.ExcPending
	if v15669 != 0 {
		goto L1
	} else {
		goto L2374
	}
L2365:
	;
	v15649 = *(*int32)(unsafe.Add(mBase, uint32(v15648)))
	if v15649 != int32(7) {
		goto L2364
	} else {
		goto L2368
	}
L2366:
	;
	goto L2367
L2367:
	;
	v15655 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+128))
	if v15655 == int32(0) {
		v15671 = v15647
		goto L2363
	} else {
		goto L2370
	}
L2368:
	;
	v15652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15648)+32)))
	if v15652 != int32(1) {
		goto L2364
	} else {
		goto L2369
	}
L2369:
	;
	goto L2367
L2370:
	;
	v15658 = *(*int32)(unsafe.Add(mBase, uint32(v15655)))
	if v15658 != int32(7) {
		goto L2364
	} else {
		goto L2371
	}
L2371:
	;
	v15661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15655)+32)))
	if v15661 != 0 {
		v15671 = v15647
		goto L2363
	} else {
		goto L2372
	}
L2372:
	;
	v15662 = *(*int64)(unsafe.Add(mBase, uint32(v15655)+24))
	if v15662 == int64(0) {
		v15671 = v15647
		goto L2363
	} else {
		goto L2373
	}
L2373:
	;
	goto L2364
L2374:
	;
	v15671 = v15668
	goto L2363
L2375:
	;
	v15675 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+128))
	v15676 = int32(0)
	if v15675 == v15676 {
		goto L2379
	} else {
		goto L2380
	}
L2376:
	;
	v16462 = v15560
	v16466 = v15564
	v16470 = v15568
	v16472 = v15570
	v16473 = v15571
	v16475 = v15573
	v16485 = v15583
	v16496 = v15594
	v16500 = v15598
	v16501 = v15599
	v16503 = v15671
	goto L2377
L2377:
	;
	F_add_path(m, v15560, v16503)
	mBase = m.M
	v16505 = m.ExcPending
	if v16505 != 0 {
		goto L1
	} else {
		goto L2539
	}
L2378:
	;
	v15722 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+32))
	if v15721 == int32(2) {
		goto L2396
	} else {
		goto L2397
	}
L2379:
	;
	v15721 = int32(0)
	goto L2378
L2380:
	;
	goto L2381
L2381:
	;
	v15684 = int32(1)
	v15685 = *(*int32)(unsafe.Add(mBase, uint32(v15675)+4))
	if v15685 <= v15684 {
		goto L2382
	} else {
		goto L2383
	}
L2382:
	;
	v15688 = v15684
	goto L2384
L2383:
	;
	v15688 = v15685
	goto L2384
L2384:
	;
	v15692 = int32(0)
	v15694 = v15676
	goto L2385
L2385:
	;
	v15701 = *(*int32)(unsafe.Add(mBase, uint32(v15675+int32(8)+v15692<<(uint(int32(2))%32))))
	if v15701 != 0 {
		goto L2388
	} else {
		goto L2389
	}
L2386:
	;
	v15721 = v15713
	goto L2378
L2387:
	;
	goto L2386
L2388:
	;
	v15702 = int32(2)
	if v15694 != 0 {
		v15713 = v15702
		goto L2387
	} else {
		goto L2391
	}
L2389:
	;
	v15708 = v15694
	goto L2390
L2390:
	;
	v15710 = v15692 + int32(1)
	if v15710 != v15688 {
		v15692 = v15710
		v15694 = v15708
		goto L2385
	} else {
		goto L2393
	}
L2391:
	;
	v15703 = int32(1)
	if base.Ui32(v15703) < base.Ui32(base.I32_popcnt(v15701)) {
		v15713 = v15702
		goto L2387
	} else {
		goto L2392
	}
L2392:
	;
	v15708 = v15703
	goto L2390
L2393:
	;
	v15713 = v15708
	goto L2387
L2394:
	;
	v16407 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	v16408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15573)+24)))
	v16409 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+32))
	v16410 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+140))
	if v16410 != 0 {
		goto L2530
	} else {
		goto L2531
	}
L2395:
	;
	v16358 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v16357))) = v16358
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+4)) = v16358
	v16364 = F_list_make1_impl(m, int32(1), v15571+int32(4))
	mBase = m.M
	v16365 = m.ExcPending
	if v16365 != 0 {
		goto L1
	} else {
		goto L2529
	}
L2396:
	;
	v15725 = F_find_base_rel(m, v15568, v15722)
	mBase = m.M
	v15726 = m.ExcPending
	if v15726 != 0 {
		goto L1
	} else {
		goto L2399
	}
L2397:
	;
	goto L2398
L2398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+44)) = v15722
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+132)) = v15722
	v16263 = F_list_make1_impl(m, int32(479), v15571+int32(44))
	mBase = m.M
	v16264 = m.ExcPending
	if v16264 != 0 {
		goto L1
	} else {
		goto L2509
	}
L2399:
	;
	v15727 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+32))
	v15728 = int32(0)
	v15733 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+132))
	if v15733 == v15728 {
		goto L2402
	} else {
		goto L2403
	}
L2400:
	;
	if int32(0) <= v15790 {
		goto L2411
	} else {
		goto L2412
	}
L2401:
	;
	v15790 = base.I32_ctz(v15776) | v15777<<(uint(int32(5))%32)
	goto L2400
L2402:
	;
	v15790 = int32(-2)
	goto L2400
L2403:
	;
	v15741 = int32(0)
	v15744 = *(*int32)(unsafe.Add(mBase, uint32(v15733)+4))
	if v15744 <= v15741 {
		goto L2402
	} else {
		goto L2404
	}
L2404:
	;
	v15747 = v15733 + int32(8)
	v15751 = *(*int32)(unsafe.Add(mBase, uint32(v15747)))
	v15754 = v15751 & int32(-1)
	if v15754 != 0 {
		v15776 = v15754
		v15777 = v15741
		goto L2401
	} else {
		goto L2405
	}
L2405:
	;
	v15755 = int32(1)
	if v15755 == v15744 {
		goto L2402
	} else {
		goto L2406
	}
L2406:
	;
	v15759 = v15755
	goto L2407
L2407:
	;
	v15766 = *(*int32)(unsafe.Add(mBase, uint32(v15747+v15759<<(uint(int32(2))%32))))
	if v15766 != 0 {
		v15776 = v15766
		v15777 = v15759
		goto L2401
	} else {
		goto L2409
	}
L2408:
	;
	goto L2402
L2409:
	;
	v15768 = v15759 + int32(1)
	if v15768 != v15744 {
		v15759 = v15768
		goto L2407
	} else {
		goto L2410
	}
L2410:
	;
	goto L2408
L2411:
	;
	v15799 = v15790
	v15801 = v15728
	v15808 = v15728
	v15809 = v15728
	v15812 = v15728
	v15813 = v15728
	v15815 = int32(0)
	goto L2414
L2412:
	;
	v16171 = v15728
	v16178 = v15728
	v16179 = v15728
	v16182 = v15728
	v16183 = v15728
	goto L2413
L2413:
	;
	v16205 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+24)) = v16205
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+156)) = v16205
	v16211 = F_list_make1_impl(m, int32(479), v15571+int32(24))
	mBase = m.M
	v16212 = m.ExcPending
	if v16212 != 0 {
		goto L1
	} else {
		goto L2491
	}
L2414:
	;
	v15835 = F_find_base_rel(m, v15568, v15799)
	mBase = m.M
	v15836 = m.ExcPending
	if v15836 != 0 {
		goto L1
	} else {
		goto L2417
	}
L2415:
	;
	if v16085 != 0 {
		v16373 = v16071
		v16380 = v16078
		v16381 = v16079
		v16384 = v16082
		v16385 = v16083
		v16386 = v15727
		v16387 = v16085
		goto L2394
	} else {
		goto L2490
	}
L2416:
	;
	v16105 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+132))
	if v16105 == int32(0) {
		goto L2480
	} else {
		goto L2481
	}
L2417:
	;
	v15837 = int32(0)
	v15839 = *(*int32)(unsafe.Add(mBase, uint32(v15835)+44))
	if v15839 == v15837 {
		v15860 = v15837
		goto L2419
	} else {
		goto L2420
	}
L2418:
	;
	if v15860 != 0 {
		v16071 = v15801
		v16078 = v15808
		v16079 = v15809
		v16082 = v15812
		v16083 = v15813
		v16085 = v15815
		goto L2416
	} else {
		goto L2428
	}
L2419:
	;
	goto L2418
L2420:
	;
	v15842 = *(*int32)(unsafe.Add(mBase, uint32(v15839)+12))
	v15843 = v15842
	goto L2421
L2421:
	;
	v15846 = *(*int32)(unsafe.Add(mBase, uint32(v15843)))
	v15847 = *(*int32)(unsafe.Add(mBase, uint32(v15846)))
	if base.Ui32(int32(2)) <= base.Ui32(v15847-int32(303)) {
		goto L2423
	} else {
		goto L2424
	}
L2422:
	;
	v15860 = int32(1)
	goto L2419
L2423:
	;
	if v15847 != int32(293) {
		v15860 = v15837
		goto L2419
	} else {
		goto L2426
	}
L2424:
	;
	v15843 = v15846 + int32(72)
	goto L2421
L2425:
	;
	goto L2422
L2426:
	;
	v15854 = *(*int32)(unsafe.Add(mBase, uint32(v15846)+72))
	if v15854 != 0 {
		v15860 = v15837
		goto L2419
	} else {
		goto L2427
	}
L2427:
	;
	goto L2425
L2428:
	;
	v15861 = F_lappend_int(m, v15815, v15799)
	mBase = m.M
	v15862 = m.ExcPending
	if v15862 != 0 {
		goto L1
	} else {
		goto L2429
	}
L2429:
	;
	v15863 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v15863 == int32(2) {
		goto L2430
	} else {
		goto L2431
	}
L2430:
	;
	v15866 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+288))
	if v15725 != v15835 {
		goto L2433
	} else {
		goto L2434
	}
L2431:
	;
	v15876 = v15809
	goto L2432
L2432:
	;
	v15877 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+152))
	if v15877 != 0 {
		goto L2438
	} else {
		goto L2439
	}
L2433:
	;
	v15868 = *(*int32)(unsafe.Add(mBase, uint32(v15835)+76))
	v15869 = *(*int32)(unsafe.Add(mBase, uint32(v15725)+76))
	v15870 = F_adjust_inherited_attnums_multilevel(m, v15568, v15866, v15868, v15869)
	mBase = m.M
	v15871 = m.ExcPending
	if v15871 != 0 {
		goto L1
	} else {
		goto L2436
	}
L2434:
	;
	v15872 = v15866
	goto L2435
L2435:
	;
	v15873 = F_lappend(m, v15809, v15872)
	mBase = m.M
	v15874 = m.ExcPending
	if v15874 != 0 {
		goto L1
	} else {
		goto L2437
	}
L2436:
	;
	v15872 = v15870
	goto L2435
L2437:
	;
	v15876 = v15873
	goto L2432
L2438:
	;
	if v15725 != v15835 {
		goto L2441
	} else {
		goto L2442
	}
L2439:
	;
	v15884 = v15812
	goto L2440
L2440:
	;
	v15885 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+96))
	if v15885 != 0 {
		goto L2446
	} else {
		goto L2447
	}
L2441:
	;
	v15879 = F_adjust_appendrel_attrs_multilevel(m, v15568, v15877, v15835, v15725)
	mBase = m.M
	v15880 = m.ExcPending
	if v15880 != 0 {
		goto L1
	} else {
		goto L2444
	}
L2442:
	;
	v15881 = v15877
	goto L2443
L2443:
	;
	v15882 = F_lappend(m, v15812, v15881)
	mBase = m.M
	v15883 = m.ExcPending
	if v15883 != 0 {
		goto L1
	} else {
		goto L2445
	}
L2444:
	;
	v15881 = v15879
	goto L2443
L2445:
	;
	v15884 = v15882
	goto L2440
L2446:
	;
	if v15725 != v15835 {
		goto L2449
	} else {
		goto L2450
	}
L2447:
	;
	v15892 = v15808
	goto L2448
L2448:
	;
	v15893 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+64))
	if v15893 != 0 {
		goto L2454
	} else {
		goto L2455
	}
L2449:
	;
	v15887 = F_adjust_appendrel_attrs_multilevel(m, v15568, v15885, v15835, v15725)
	mBase = m.M
	v15888 = m.ExcPending
	if v15888 != 0 {
		goto L1
	} else {
		goto L2452
	}
L2450:
	;
	v15889 = v15885
	goto L2451
L2451:
	;
	v15890 = F_lappend(m, v15808, v15889)
	mBase = m.M
	v15891 = m.ExcPending
	if v15891 != 0 {
		goto L1
	} else {
		goto L2453
	}
L2452:
	;
	v15889 = v15887
	goto L2451
L2453:
	;
	v15892 = v15890
	goto L2448
L2454:
	;
	v15894 = int32(0)
	v15895 = *(*int32)(unsafe.Add(mBase, uint32(v15893)+4))
	if v15894 < v15895 {
		goto L2457
	} else {
		goto L2458
	}
L2455:
	;
	v16032 = v15813
	goto L2456
L2456:
	;
	v16054 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v16054 != int32(5) {
		v16071 = v15801
		v16078 = v15892
		v16079 = v15876
		v16082 = v15884
		v16083 = v16032
		v16085 = v15861
		goto L2416
	} else {
		goto L2472
	}
L2457:
	;
	v15902 = v15894
	v15915 = int32(0)
	goto L2460
L2458:
	;
	v15973 = v15894
	goto L2459
L2459:
	;
	v16011 = F_lappend(m, v15813, v15973)
	mBase = m.M
	v16012 = m.ExcPending
	if v16012 != 0 {
		goto L1
	} else {
		goto L2471
	}
L2460:
	;
	v15940 = *(*int32)(unsafe.Add(mBase, uint32(v15893)+12))
	v15944 = *(*int32)(unsafe.Add(mBase, uint32(v15940+v15915<<(uint(int32(2))%32))))
	v15945 = F_copyObjectImpl(m, v15944)
	mBase = m.M
	v15946 = m.ExcPending
	if v15946 != 0 {
		goto L1
	} else {
		goto L2462
	}
L2461:
	;
	v15973 = v15964
	goto L2459
L2462:
	;
	v15947 = *(*int32)(unsafe.Add(mBase, uint32(v15944)+16))
	v15948 = F_adjust_appendrel_attrs_multilevel(m, v15568, v15947, v15835, v15725)
	mBase = m.M
	v15949 = m.ExcPending
	if v15949 != 0 {
		goto L1
	} else {
		goto L2463
	}
L2463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15945)+16)) = v15948
	v15951 = *(*int32)(unsafe.Add(mBase, uint32(v15944)+20))
	v15952 = F_adjust_appendrel_attrs_multilevel(m, v15568, v15951, v15835, v15725)
	mBase = m.M
	v15953 = m.ExcPending
	if v15953 != 0 {
		goto L1
	} else {
		goto L2464
	}
L2464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15945)+20)) = v15952
	v15955 = *(*int32)(unsafe.Add(mBase, uint32(v15945)+8))
	if v15955 == int32(2) {
		goto L2465
	} else {
		goto L2466
	}
L2465:
	;
	v15958 = *(*int32)(unsafe.Add(mBase, uint32(v15944)+24))
	v15959 = *(*int32)(unsafe.Add(mBase, uint32(v15835)+76))
	v15960 = *(*int32)(unsafe.Add(mBase, uint32(v15725)+76))
	v15961 = F_adjust_inherited_attnums_multilevel(m, v15568, v15958, v15959, v15960)
	mBase = m.M
	v15962 = m.ExcPending
	if v15962 != 0 {
		goto L1
	} else {
		goto L2468
	}
L2466:
	;
	goto L2467
L2467:
	;
	v15964 = F_lappend(m, v15902, v15945)
	mBase = m.M
	v15965 = m.ExcPending
	if v15965 != 0 {
		goto L1
	} else {
		goto L2469
	}
L2468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15945)+24)) = v15961
	goto L2467
L2469:
	;
	v15967 = v15915 + int32(1)
	v15968 = *(*int32)(unsafe.Add(mBase, uint32(v15893)+4))
	if v15967 < v15968 {
		v15902 = v15964
		v15915 = v15967
		goto L2460
	} else {
		goto L2470
	}
L2470:
	;
	goto L2461
L2471:
	;
	v16032 = v16011
	goto L2456
L2472:
	;
	v16057 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+72))
	if v15725 != v15835 {
		goto L2473
	} else {
		goto L2474
	}
L2473:
	;
	v16059 = F_adjust_appendrel_attrs_multilevel(m, v15568, v16057, v15835, v15725)
	mBase = m.M
	v16060 = m.ExcPending
	if v16060 != 0 {
		goto L1
	} else {
		goto L2476
	}
L2474:
	;
	v16061 = v16057
	goto L2475
L2475:
	;
	v16062 = F_lappend(m, v15801, v16061)
	mBase = m.M
	v16063 = m.ExcPending
	if v16063 != 0 {
		goto L1
	} else {
		goto L2477
	}
L2476:
	;
	v16061 = v16059
	goto L2475
L2477:
	;
	v16071 = v16062
	v16078 = v15892
	v16079 = v15876
	v16082 = v15884
	v16083 = v16032
	v16085 = v15861
	goto L2416
L2478:
	;
	if int32(0) <= v16161 {
		v15799 = v16161
		v15801 = v16071
		v15808 = v16078
		v15809 = v16079
		v15812 = v16082
		v15813 = v16083
		v15815 = v16085
		goto L2414
	} else {
		goto L2489
	}
L2479:
	;
	v16161 = base.I32_ctz(v16147) | v16148<<(uint(int32(5))%32)
	goto L2478
L2480:
	;
	v16161 = int32(-2)
	goto L2478
L2481:
	;
	v16112 = v15799 + int32(1)
	v16114 = int32(base.Ui32(v16112) >> (uint(int32(5)) % 32))
	v16115 = *(*int32)(unsafe.Add(mBase, uint32(v16105)+4))
	if v16115 <= v16114 {
		goto L2480
	} else {
		goto L2482
	}
L2482:
	;
	v16118 = v16105 + int32(8)
	v16122 = *(*int32)(unsafe.Add(mBase, uint32(v16118+v16114<<(uint(int32(2))%32))))
	v16125 = v16122 & (int32(-1) << (uint(v16112) % 32))
	if v16125 != 0 {
		v16147 = v16125
		v16148 = v16114
		goto L2479
	} else {
		goto L2483
	}
L2483:
	;
	v16127 = v16114 + int32(1)
	if v16127 == v16115 {
		goto L2480
	} else {
		goto L2484
	}
L2484:
	;
	v16130 = v16127
	goto L2485
L2485:
	;
	v16137 = *(*int32)(unsafe.Add(mBase, uint32(v16118+v16130<<(uint(int32(2))%32))))
	if v16137 != 0 {
		v16147 = v16137
		v16148 = v16130
		goto L2479
	} else {
		goto L2487
	}
L2486:
	;
	goto L2480
L2487:
	;
	v16139 = v16130 + int32(1)
	if v16139 != v16115 {
		v16130 = v16139
		goto L2485
	} else {
		goto L2488
	}
L2488:
	;
	goto L2486
L2489:
	;
	goto L2415
L2490:
	;
	v16171 = v16071
	v16178 = v16078
	v16179 = v16079
	v16182 = v16082
	v16183 = v16083
	goto L2413
L2491:
	;
	v16213 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v16213 == int32(2) {
		goto L2492
	} else {
		goto L2493
	}
L2492:
	;
	v16216 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+20)) = v16216
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+152)) = v16216
	v16222 = F_list_make1_impl(m, int32(1), v15571+int32(20))
	mBase = m.M
	v16223 = m.ExcPending
	if v16223 != 0 {
		goto L1
	} else {
		goto L2495
	}
L2493:
	;
	v16225 = v16179
	goto L2494
L2494:
	;
	v16226 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+152))
	if v16226 != 0 {
		goto L2496
	} else {
		goto L2497
	}
L2495:
	;
	v16225 = v16222
	goto L2494
L2496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+16)) = v16226
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+148)) = v16226
	v16232 = F_list_make1_impl(m, int32(1), v15571+int32(16))
	mBase = m.M
	v16233 = m.ExcPending
	if v16233 != 0 {
		goto L1
	} else {
		goto L2499
	}
L2497:
	;
	v16234 = v16182
	goto L2498
L2498:
	;
	v16235 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+96))
	if v16235 != 0 {
		goto L2500
	} else {
		goto L2501
	}
L2499:
	;
	v16234 = v16232
	goto L2498
L2500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+12)) = v16235
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+144)) = v16235
	v16241 = F_list_make1_impl(m, int32(1), v15571+int32(12))
	mBase = m.M
	v16242 = m.ExcPending
	if v16242 != 0 {
		goto L1
	} else {
		goto L2503
	}
L2501:
	;
	v16243 = v16178
	goto L2502
L2502:
	;
	v16244 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+64))
	if v16244 != 0 {
		goto L2504
	} else {
		goto L2505
	}
L2503:
	;
	v16243 = v16241
	goto L2502
L2504:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+8)) = v16244
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+140)) = v16244
	v16250 = F_list_make1_impl(m, int32(1), v15571+int32(8))
	mBase = m.M
	v16251 = m.ExcPending
	if v16251 != 0 {
		goto L1
	} else {
		goto L2507
	}
L2505:
	;
	v16252 = v16183
	goto L2506
L2506:
	;
	v16253 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v16253 != int32(5) {
		v16373 = v16171
		v16380 = v16243
		v16381 = v16225
		v16384 = v16234
		v16385 = v16252
		v16386 = v15727
		v16387 = v16211
		goto L2394
	} else {
		goto L2508
	}
L2507:
	;
	v16252 = v16250
	goto L2506
L2508:
	;
	v16330 = v16243
	v16331 = v16225
	v16334 = v16234
	v16335 = v16252
	v16336 = v15727
	v16337 = v16211
	v16357 = v15571 + int32(136)
	goto L2395
L2509:
	;
	v16265 = int32(0)
	v16267 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v16267 == int32(2) {
		goto L2510
	} else {
		goto L2511
	}
L2510:
	;
	v16270 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+288))
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+40)) = v16270
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+128)) = v16270
	v16276 = F_list_make1_impl(m, int32(1), v15571+int32(40))
	mBase = m.M
	v16277 = m.ExcPending
	if v16277 != 0 {
		goto L1
	} else {
		goto L2513
	}
L2511:
	;
	v16279 = v16265
	goto L2512
L2512:
	;
	v16280 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+152))
	if v16280 != 0 {
		goto L2514
	} else {
		goto L2515
	}
L2513:
	;
	v16279 = v16276
	goto L2512
L2514:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+36)) = v16280
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+124)) = v16280
	v16286 = F_list_make1_impl(m, int32(1), v15571+int32(36))
	mBase = m.M
	v16287 = m.ExcPending
	if v16287 != 0 {
		goto L1
	} else {
		goto L2517
	}
L2515:
	;
	v16288 = v16265
	goto L2516
L2516:
	;
	v16289 = int32(0)
	v16291 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+96))
	if v16291 != 0 {
		goto L2518
	} else {
		goto L2519
	}
L2517:
	;
	v16288 = v16286
	goto L2516
L2518:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+32)) = v16291
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+120)) = v16291
	v16297 = F_list_make1_impl(m, int32(1), v15571+int32(32))
	mBase = m.M
	v16298 = m.ExcPending
	if v16298 != 0 {
		goto L1
	} else {
		goto L2521
	}
L2519:
	;
	v16299 = v16289
	goto L2520
L2520:
	;
	v16300 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+64))
	if v16300 != 0 {
		goto L2522
	} else {
		goto L2523
	}
L2521:
	;
	v16299 = v16297
	goto L2520
L2522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+28)) = v16300
	*(*int32)(unsafe.Add(mBase, uint32(v15571)+116)) = v16300
	v16306 = F_list_make1_impl(m, int32(1), v15571+int32(28))
	mBase = m.M
	v16307 = m.ExcPending
	if v16307 != 0 {
		goto L1
	} else {
		goto L2525
	}
L2523:
	;
	v16308 = v16289
	goto L2524
L2524:
	;
	v16309 = int32(0)
	v16310 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+4))
	if v16310 != int32(5) {
		goto L2526
	} else {
		goto L2527
	}
L2525:
	;
	v16308 = v16306
	goto L2524
L2526:
	;
	v16373 = int32(0)
	v16380 = v16299
	v16381 = v16279
	v16384 = v16288
	v16385 = v16308
	v16386 = v16309
	v16387 = v16263
	goto L2394
L2527:
	;
	goto L2528
L2528:
	;
	v16330 = v16299
	v16331 = v16279
	v16334 = v16288
	v16335 = v16308
	v16336 = v16309
	v16337 = v16263
	v16357 = v15571 + int32(112)
	goto L2395
L2529:
	;
	v16373 = v16364
	v16380 = v16330
	v16381 = v16331
	v16384 = v16334
	v16385 = v16335
	v16386 = v16336
	v16387 = v16337
	goto L2394
L2530:
	;
	v16413 = int32(0)
	goto L2532
L2531:
	;
	v16412 = *(*int32)(unsafe.Add(mBase, uint32(v15568)+144))
	v16413 = v16412
	goto L2532
L2532:
	;
	v16414 = *(*int32)(unsafe.Add(mBase, uint32(v15573)+84))
	v16415 = F_assign_special_exec_param(m, v15568)
	mBase = m.M
	v16416 = m.ExcPending
	if v16416 != 0 {
		goto L1
	} else {
		goto L2533
	}
L2533:
	;
	v16419 = F_palloc0(m, int32(128))
	mBase = m.M
	v16420 = m.ExcPending
	if v16420 != 0 {
		goto L1
	} else {
		goto L2534
	}
L2534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+8)) = v15560
	*(*int64)(unsafe.Add(mBase, uint32(v16419))) = int64(1447403979070)
	v16424 = *(*int32)(unsafe.Add(mBase, uint32(v15560)+40))
	v16425 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+64)) = v16425
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+24)) = v16425
	*(*uint16)(unsafe.Add(mBase, uint32(v16419)+20)) = uint16(v16425)
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+16)) = v16425
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+12)) = v16424
	v16434 = *(*int32)(unsafe.Add(mBase, uint32(v15671)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+40)) = v16434
	v16436 = *(*float64)(unsafe.Add(mBase, uint32(v15671)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16419)+48)) = v16436
	v16438 = *(*float64)(unsafe.Add(mBase, uint32(v15671)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16419)+56)) = v16438
	if v16380 != 0 {
		goto L2536
	} else {
		goto L2537
	}
L2535:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16424)+32)) = v16446
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+124)) = v16373
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+120)) = v16385
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+116)) = v16415
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+112)) = v16414
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+108)) = v16413
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+104)) = v16380
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+100)) = v16384
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+96)) = v16381
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+92)) = v16387
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+88)) = v16386
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+84)) = v16409
	*(*uint8)(unsafe.Add(mBase, uint32(v16419)+80)) = uint8(v16408)
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+76)) = v16407
	*(*int32)(unsafe.Add(mBase, uint32(v16419)+72)) = v15671
	v16462 = v15560
	v16466 = v15564
	v16470 = v15568
	v16472 = v15570
	v16473 = v15571
	v16475 = v15573
	v16485 = v15583
	v16496 = v15594
	v16500 = v15598
	v16501 = v15599
	v16503 = v16419
	goto L2377
L2536:
	;
	v16440 = *(*float64)(unsafe.Add(mBase, uint32(v15671)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v16419)+32)) = v16440
	v16442 = *(*int32)(unsafe.Add(mBase, uint32(v15671)+12))
	v16443 = *(*int32)(unsafe.Add(mBase, uint32(v16442)+32))
	v16446 = v16443
	goto L2535
L2537:
	;
	goto L2538
L2538:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16419)+32)) = int64(0)
	v16446 = int32(0)
	goto L2535
L2539:
	;
	v16507 = v15582 + int32(1)
	v16508 = *(*int32)(unsafe.Add(mBase, uint32(v16472)+4))
	if v16507 < v16508 {
		v15560 = v16462
		v15564 = v16466
		v15568 = v16470
		v15570 = v16472
		v15571 = v16473
		v15573 = v16475
		v15582 = v16507
		v15583 = v16485
		v15594 = v16496
		v15598 = v16500
		v15599 = v16501
		goto L2356
	} else {
		goto L2540
	}
L2540:
	;
	goto L2357
L2541:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16521)+232)) = v16548
	*(*int64)(unsafe.Add(mBase, uint32(v16521)+224)) = v16549
	*(*float64)(unsafe.Add(mBase, uint32(v16521)+216)) = v16544
	*(*uint8)(unsafe.Add(mBase, uint32(v16521)+208)) = uint8(v16699)
	v16743 = *(*int32)(unsafe.Add(mBase, uint32(v16510)+176))
	if v16743 == int32(0) {
		goto L2575
	} else {
		goto L2576
	}
L2542:
	;
	v16675 = *(*int32)(unsafe.Add(mBase, uint32(v16523)+132))
	if v16675 != 0 {
		goto L2563
	} else {
		goto L2564
	}
L2543:
	;
	v16554 = *(*int32)(unsafe.Add(mBase, uint32(v16518)+12))
	if base.Ui32(v16554) < base.Ui32(int32(2)) {
		goto L2542
	} else {
		goto L2544
	}
L2544:
	;
	v16557 = *(*int32)(unsafe.Add(mBase, uint32(v16523)+132))
	if v16557 != 0 {
		goto L2545
	} else {
		goto L2546
	}
L2545:
	;
	v16558 = *(*int32)(unsafe.Add(mBase, uint32(v16557)))
	if v16558 != int32(7) {
		goto L2548
	} else {
		goto L2549
	}
L2546:
	;
	goto L2547
L2547:
	;
	v16565 = *(*int32)(unsafe.Add(mBase, uint32(v16523)+128))
	if v16565 == int32(0) {
		goto L2552
	} else {
		goto L2553
	}
L2548:
	;
	v16699 = int32(1)
	goto L2541
L2549:
	;
	goto L2550
L2550:
	;
	v16562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16557)+32)))
	if v16562 != int32(1) {
		goto L2542
	} else {
		goto L2551
	}
L2551:
	;
	goto L2547
L2552:
	;
	v16575 = *(*int32)(unsafe.Add(mBase, uint32(v16514)+52))
	if v16575 == int32(0) {
		goto L2542
	} else {
		goto L2557
	}
L2553:
	;
	v16568 = *(*int32)(unsafe.Add(mBase, uint32(v16565)))
	if v16568 != int32(7) {
		goto L2542
	} else {
		goto L2554
	}
L2554:
	;
	v16571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16565)+32)))
	if v16571 != 0 {
		goto L2552
	} else {
		goto L2555
	}
L2555:
	;
	v16572 = *(*int64)(unsafe.Add(mBase, uint32(v16565)+24))
	if v16572 != int64(0) {
		goto L2542
	} else {
		goto L2556
	}
L2556:
	;
	goto L2552
L2557:
	;
	v16578 = *(*int32)(unsafe.Add(mBase, uint32(v16575)+4))
	if v16578 <= int32(0) {
		goto L2542
	} else {
		goto L2558
	}
L2558:
	;
	v16583 = int32(0)
	goto L2559
L2559:
	;
	v16623 = *(*int32)(unsafe.Add(mBase, uint32(v16575)+12))
	v16627 = *(*int32)(unsafe.Add(mBase, uint32(v16623+v16583<<(uint(int32(2))%32))))
	F_add_partial_path(m, v16510, v16627)
	mBase = m.M
	v16629 = m.ExcPending
	if v16629 != 0 {
		goto L1
	} else {
		goto L2561
	}
L2560:
	;
	goto L2542
L2561:
	;
	v16631 = v16583 + int32(1)
	v16632 = *(*int32)(unsafe.Add(mBase, uint32(v16575)+4))
	if v16631 < v16632 {
		v16583 = v16631
		goto L2559
	} else {
		goto L2562
	}
L2562:
	;
	goto L2560
L2563:
	;
	v16676 = *(*int32)(unsafe.Add(mBase, uint32(v16675)))
	if v16676 != int32(7) {
		goto L2566
	} else {
		goto L2567
	}
L2564:
	;
	goto L2565
L2565:
	;
	v16685 = *(*int32)(unsafe.Add(mBase, uint32(v16523)+128))
	if v16685 == int32(0) {
		goto L2570
	} else {
		goto L2571
	}
L2566:
	;
	v16699 = int32(1)
	goto L2541
L2567:
	;
	goto L2568
L2568:
	;
	v16680 = int32(1)
	v16681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16675)+32)))
	if v16681 != v16680 {
		v16699 = v16680
		goto L2541
	} else {
		goto L2569
	}
L2569:
	;
	goto L2565
L2570:
	;
	v16699 = int32(0)
	goto L2541
L2571:
	;
	v16688 = int32(1)
	v16689 = *(*int32)(unsafe.Add(mBase, uint32(v16685)))
	if v16689 != int32(7) {
		v16699 = v16688
		goto L2541
	} else {
		goto L2572
	}
L2572:
	;
	v16692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16685)+32)))
	if v16692 != 0 {
		goto L2570
	} else {
		goto L2573
	}
L2573:
	;
	v16693 = *(*int64)(unsafe.Add(mBase, uint32(v16685)+24))
	if v16693 != int64(0) {
		v16699 = v16688
		goto L2541
	} else {
		goto L2574
	}
L2574:
	;
	goto L2570
L2575:
	;
	v16756 = *(*int32)(unsafe.Add(mBase, _c_F_subquery_planner[2]))
	if v16756 != 0 {
		goto L2579
	} else {
		goto L2580
	}
L2576:
	;
	v16746 = *(*int32)(unsafe.Add(mBase, uint32(v16743)+36))
	if v16746 == int32(0) {
		goto L2575
	} else {
		goto L2577
	}
L2577:
	;
	m.T0[v16746].(func(*base.Module, int32, int32, int32, int32, int32))(m, v16518, int32(7), v16514, v16510, v16521+int32(208))
	mBase = m.M
	v16753 = m.ExcPending
	if v16753 != 0 {
		goto L1
	} else {
		goto L2578
	}
L2578:
	;
	goto L2575
L2579:
	;
	m.T0[v16756].(func(*base.Module, int32, int32, int32, int32, int32))(m, v16518, int32(7), v16514, v16510, v16521+int32(208))
	mBase = m.M
	v16761 = m.ExcPending
	if v16761 != 0 {
		goto L1
	} else {
		goto L2582
	}
L2580:
	;
	goto L2581
L2581:
	;
	m.G0 = v16521 + int32(368)
	goto L582
L2582:
	;
	goto L2581
L2583:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v16771 = m.ExcPending
	if v16771 != 0 {
		goto L1
	} else {
		goto L2584
	}
L2584:
	;
	F_errmsg(m, int32(_a_F_subquery_planner_44), int32(0))
	mBase = m.M
	v16775 = m.ExcPending
	if v16775 != 0 {
		goto L1
	} else {
		goto L2585
	}
L2585:
	;
	v16778 = F_errdetail(m, int32(_a_F_subquery_planner_18), int32(0))
	mBase = m.M
	v16779 = m.ExcPending
	if v16779 != 0 {
		goto L1
	} else {
		goto L2586
	}
L2586:
	;
	F_errfinish(m, int32(_a_F_subquery_planner_8), int32(_a_F_subquery_planner_45), int32(_a_F_subquery_planner_46))
	mBase = m.M
	v16784 = m.ExcPending
	if v16784 != 0 {
		goto L1
	} else {
		goto L2587
	}
L2587:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2588:
	;
	v16789 = F_fetch_upper_rel(m, v16518, int32(7), int32(0))
	mBase = m.M
	v16790 = m.ExcPending
	if v16790 != 0 {
		goto L1
	} else {
		goto L2589
	}
L2589:
	;
	v16791 = int32(0)
	v16798 = float64(0)
	v16799 = *(*int32)(unsafe.Add(mBase, uint32(v16518)+80))
	if v16799 == v16791 {
		goto L2591
	} else {
		goto L2592
	}
L2590:
	;
	F_set_cheapest(m, v16789)
	mBase = m.M
	v16987 = m.ExcPending
	if v16987 != 0 {
		goto L1
	} else {
		goto L2620
	}
L2591:
	;
	goto L2590
L2592:
	;
	v16802 = *(*int32)(unsafe.Add(mBase, uint32(v16799)+4))
	if v16802 <= int32(0) {
		v16875 = v16791
		v16881 = v16798
		goto L2593
	} else {
		goto L2594
	}
L2593:
	;
	v16882 = *(*int32)(unsafe.Add(mBase, uint32(v16789)+44))
	if v16882 == int32(0) {
		goto L2603
	} else {
		goto L2604
	}
L2594:
	;
	v16805 = *(*int32)(unsafe.Add(mBase, uint32(v16799)+12))
	if v16802 == int32(1) {
		goto L2596
	} else {
		goto L2597
	}
L2595:
	;
	v16863 = *(*int32)(unsafe.Add(mBase, uint32(v16805+v16850<<(uint(int32(2))%32))))
	v16864 = *(*float64)(unsafe.Add(mBase, uint32(v16863)+56))
	v16865 = *(*float64)(unsafe.Add(mBase, uint32(v16863)+64))
	v16868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16863)+39)))
	v16875 = v16868 ^ int32(1) | v16853
	v16881 = base.F64_add(v16859, base.F64_add(v16864, v16865))
	goto L2593
L2596:
	;
	v16850 = int32(0)
	v16853 = v16791
	v16859 = v16798
	goto L2595
L2597:
	;
	goto L2598
L2598:
	;
	v16814 = int32(0)
	v16817 = v16791
	v16822 = v16791
	v16823 = v16798
	goto L2599
L2599:
	;
	v16824 = int32(2)
	v16826 = v16805 + v16814<<(uint(v16824)%32)
	v16827 = *(*int32)(unsafe.Add(mBase, uint32(v16826)))
	v16828 = *(*float64)(unsafe.Add(mBase, uint32(v16827)+56))
	v16829 = *(*float64)(unsafe.Add(mBase, uint32(v16827)+64))
	v16832 = *(*int32)(unsafe.Add(mBase, uint32(v16826)+4))
	v16833 = *(*float64)(unsafe.Add(mBase, uint32(v16832)+56))
	v16834 = *(*float64)(unsafe.Add(mBase, uint32(v16832)+64))
	v16836 = base.F64_add(base.F64_add(v16823, base.F64_add(v16828, v16829)), base.F64_add(v16833, v16834))
	v16837 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16832)+39)))
	v16838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16827)+39)))
	v16842 = base.B2i32(v16837&v16838 == int32(0)) | v16817
	v16844 = v16814 + v16824
	v16846 = v16822 + v16824
	if v16846 != v16802&int32(2147483646) {
		v16814 = v16844
		v16817 = v16842
		v16822 = v16846
		v16823 = v16836
		goto L2599
	} else {
		goto L2601
	}
L2600:
	;
	if v16802&int32(1) == int32(0) {
		v16875 = v16842
		v16881 = v16836
		goto L2593
	} else {
		goto L2602
	}
L2601:
	;
	goto L2600
L2602:
	;
	v16850 = v16844
	v16853 = v16842
	v16859 = v16836
	goto L2595
L2603:
	;
	if v16875&int32(1) != 0 {
		goto L2612
	} else {
		goto L2613
	}
L2604:
	;
	v16885 = *(*int32)(unsafe.Add(mBase, uint32(v16882)+4))
	if v16885 <= int32(0) {
		goto L2603
	} else {
		goto L2605
	}
L2605:
	;
	v16891 = int32(0)
	goto L2606
L2606:
	;
	v16901 = *(*int32)(unsafe.Add(mBase, uint32(v16882)+12))
	v16905 = *(*int32)(unsafe.Add(mBase, uint32(v16901+v16891<<(uint(int32(2))%32))))
	v16906 = *(*float64)(unsafe.Add(mBase, uint32(v16905)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16905)+48)) = base.F64_add(v16881, v16906)
	v16909 = *(*float64)(unsafe.Add(mBase, uint32(v16905)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16905)+56)) = base.F64_add(v16881, v16909)
	if v16875&int32(1) != 0 {
		goto L2608
	} else {
		goto L2609
	}
L2607:
	;
	goto L2603
L2608:
	;
	v16912 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16905)+21)) = uint8(v16912)
	goto L2610
L2609:
	;
	goto L2610
L2610:
	;
	v16915 = v16891 + int32(1)
	v16916 = *(*int32)(unsafe.Add(mBase, uint32(v16882)+4))
	if v16915 < v16916 {
		v16891 = v16915
		goto L2606
	} else {
		goto L2611
	}
L2611:
	;
	goto L2607
L2612:
	;
	v16930 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16789)+26)) = uint8(v16930)
	*(*int32)(unsafe.Add(mBase, uint32(v16789)+52)) = v16930
	goto L2590
L2613:
	;
	goto L2614
L2614:
	;
	v16934 = *(*int32)(unsafe.Add(mBase, uint32(v16789)+52))
	if v16934 == int32(0) {
		goto L2591
	} else {
		goto L2615
	}
L2615:
	;
	v16937 = *(*int32)(unsafe.Add(mBase, uint32(v16934)+4))
	if v16937 <= int32(0) {
		goto L2591
	} else {
		goto L2616
	}
L2616:
	;
	v16941 = int32(0)
	goto L2617
L2617:
	;
	v16951 = *(*int32)(unsafe.Add(mBase, uint32(v16934)+12))
	v16955 = *(*int32)(unsafe.Add(mBase, uint32(v16951+v16941<<(uint(int32(2))%32))))
	v16956 = *(*float64)(unsafe.Add(mBase, uint32(v16955)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v16955)+48)) = base.F64_add(v16881, v16956)
	v16959 = *(*float64)(unsafe.Add(mBase, uint32(v16955)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v16955)+56)) = base.F64_add(v16881, v16959)
	v16963 = v16941 + int32(1)
	v16964 = *(*int32)(unsafe.Add(mBase, uint32(v16934)+4))
	if v16963 < v16964 {
		v16941 = v16963
		goto L2617
	} else {
		goto L2619
	}
L2618:
	;
	goto L2591
L2619:
	;
	goto L2618
L2620:
	;
	m.G0 = v16533 + int32(32)
	return v16518
}
