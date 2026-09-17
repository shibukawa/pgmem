package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TablesyncWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v466 int64
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
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
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v520 int32
	_ = v520
	var v521 int64
	_ = v521
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v693 int64
	_ = v693
	var v696 int32
	_ = v696
	var v699 int64
	_ = v699
	var v702 int64
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v777 int32
	_ = v777
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v797 int32
	_ = v797
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v851 int32
	_ = v851
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v957 int32
	_ = v957
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1029 int32
	_ = v1029
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1044 int64
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1061 int32
	_ = v1061
	var v1072 int32
	_ = v1072
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1165 int32
	_ = v1165
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1209 int32
	_ = v1209
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1288 int32
	_ = v1288
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1344 int32
	_ = v1344
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1377 int32
	_ = v1377
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1397 int32
	_ = v1397
	var v1406 int32
	_ = v1406
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1473 int32
	_ = v1473
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1581 int32
	_ = v1581
	var v1592 int32
	_ = v1592
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1677 int32
	_ = v1677
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1705 int32
	_ = v1705
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1728 int32
	_ = v1728
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1751 int32
	_ = v1751
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1780 int32
	_ = v1780
	var v1807 int32
	_ = v1807
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1927 int32
	_ = v1927
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1965 int32
	_ = v1965
	var v1967 int32
	_ = v1967
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2022 int32
	_ = v2022
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2059 int32
	_ = v2059
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2129 int32
	_ = v2129
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2183 int32
	_ = v2183
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2200 int32
	_ = v2200
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2216 int32
	_ = v2216
	var v2222 int32
	_ = v2222
	var v2255 int32
	_ = v2255
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2268 int32
	_ = v2268
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2311 int32
	_ = v2311
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2352 int32
	_ = v2352
	var v2386 int32
	_ = v2386
	var v2422 int32
	_ = v2422
	var v2432 int32
	_ = v2432
	var v2455 int32
	_ = v2455
	var v2457 int32
	_ = v2457
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2484 int32
	_ = v2484
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2525 int32
	_ = v2525
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2559 int64
	_ = v2559
	var v2566 int32
	_ = v2566
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2582 int32
	_ = v2582
	var v2583 int32
	_ = v2583
	var v2589 int32
	_ = v2589
	var v2590 int32
	_ = v2590
	var v2596 int32
	_ = v2596
	var v2602 int32
	_ = v2602
	var v2608 int32
	_ = v2608
	var v2609 int32
	_ = v2609
	var v2616 int32
	_ = v2616
	var v2617 int32
	_ = v2617
	var v2623 int32
	_ = v2623
	var v2626 int32
	_ = v2626
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2632 int32
	_ = v2632
	var v2636 int32
	_ = v2636
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2661 int32
	_ = v2661
	var v2665 int32
	_ = v2665
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2676 int32
	_ = v2676
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2718 int32
	_ = v2718
	var v2723 int64
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2731 int32
	_ = v2731
	var v2737 int32
	_ = v2737
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2746 int32
	_ = v2746
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2751 int32
	_ = v2751
	var v2752 int32
	_ = v2752
	var v2762 int32
	_ = v2762
	var v2769 int32
	_ = v2769
	var v2770 int32
	_ = v2770
	var v2780 int32
	_ = v2780
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2810 int32
	_ = v2810
	var v2816 int32
	_ = v2816
	var v2826 int32
	_ = v2826
	var v2833 int32
	_ = v2833
	var v2839 int32
	_ = v2839
	var v2841 int32
	_ = v2841
	var v2842 int64
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2852 int32
	_ = v2852
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2858 int32
	_ = v2858
	var v2885 int32
	_ = v2885
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2894 int64
	_ = v2894
	var v2901 int64
	_ = v2901
	var v2910 int32
	_ = v2910
	var v2919 int32
	_ = v2919
	var v2922 int32
	_ = v2922
	var v2923 int32
	_ = v2923
	var v2931 int32
	_ = v2931
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2943 int64
	_ = v2943
	var v2975 int32
	_ = v2975
	var v2981 int32
	_ = v2981
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2992 int32
	_ = v2992
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v2999 int32
	_ = v2999
	var v3000 int32
	_ = v3000
	var v3005 int32
	_ = v3005
	var v3011 int32
	_ = v3011
	var v3015 int32
	_ = v3015
	var v3021 int32
	_ = v3021
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3033 int32
	_ = v3033
	var v3035 int32
	_ = v3035
	var v3038 int32
	_ = v3038
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3053 int32
	_ = v3053
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3069 int32
	_ = v3069
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3084 int32
	_ = v3084
	var v3092 int32
	_ = v3092
	var v3096 int32
	_ = v3096
	var v3099 int32
	_ = v3099
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3102 int32
	_ = v3102
	var v3112 int32
	_ = v3112
	var v3133 int32
	_ = v3133
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3142 int32
	_ = v3142
	var v3170 int32
	_ = v3170
	var v3171 int64
	_ = v3171
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3181 int32
	_ = v3181
	var v3183 int32
	_ = v3183
	var v3185 int32
	_ = v3185
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3191 int32
	_ = v3191
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3205 int32
	_ = v3205
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3216 int64
	_ = v3216
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3234 int32
	_ = v3234
	var v3237 int32
	_ = v3237
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3243 int32
	_ = v3243
	var v3248 int32
	_ = v3248
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3258 int32
	_ = v3258
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3264 int32
	_ = v3264
	var v3271 int32
	_ = v3271
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3287 int32
	_ = v3287
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3295 int64
	_ = v3295
	var v3297 int32
	_ = v3297
	var v3302 int32
	_ = v3302
	v2 = int32(0)
	F_SetupApplyOrSyncWorker(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = m.G0
	v32 = v30 - int32(128)
	m.G0 = v32
	*(*int64)(unsafe.Add(mBase, uint32(v32)+56)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = int32(0)
	v39 = v32 + int32(56)
	v42 = m.G0
	v44 = v42 - int32(752)
	m.G0 = v44
	v49 = v2
	v50 = v2
	v51 = v2
	v52 = v2
	v53 = v2
	v54 = int32(-1)
	v67 = v2
	v68 = v2
	goto L3
L3:
	;
	if v54 != int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v44 + int32(752)
	v3197 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(v3197)))
	v3200 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3200)+36))
	v3203 = v32 - int32(-64)
	F_ReplicationOriginNameForLogicalRep(m, v3198, v3201, v3203)
	mBase = m.M
	v3205 = m.ExcPending
	if v3205 != 0 {
		goto L1
	} else {
		goto L549
	}
L5:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2]))
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3]))
	v82 = v44 + int32(352)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v44 + int32(348)
	goto L8
L6:
	;
	v88 = v53
	v89 = v67
	v90 = v68
	goto L7
L7:
	;
	goto L9
L8:
	;
	v88 = int32(0)
	v89 = v78
	v90 = v80
	goto L7
L9:
	;
	if v88 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	goto L4
L11:
	;
	v3170 = int32(m.ExcTag)
	v3171 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v3170 == int32(0) {
		goto L539
	} else {
		goto L540
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3099
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3100
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3101
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v3102)
	v3133 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[4]))
	v3134 = F_MemoryContextStrdup(m, v3133, v3112)
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		goto L11
	} else {
		goto L537
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v2885 = m.ExcPending
	if v2885 != 0 {
		goto L11
	} else {
		goto L495
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2455 = v44 + int32(616)
	F_appendStringInfoString(m, v2455, v2432)
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L11
	} else {
		goto L420
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoChar(m, v44+int32(616), int32(41))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L11
	} else {
		goto L419
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2255 = v44 + int32(616)
	F_appendStringInfoString(m, v2255, int32(_a_F_TablesyncWorkerMain_0))
	mBase = m.M
	v2258 = m.ExcPending
	if v2258 != 0 {
		goto L11
	} else {
		goto L402
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v44 + int32(352)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[2])) = v89
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[3])) = v90
	v2190 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2190)+29)))
	if v2191 == int32(1) {
		goto L395
	} else {
		goto L396
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+36))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v113 = F_GetSubscriptionRelState(m, v106, v105, v44+int32(600))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+30)))
	if v124 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+24)))
	v130 = v127 ^ int32(1)
	goto L25
L24:
	;
	v130 = int32(0)
	goto L25
L25:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+56)) = int32(1)
	if v133 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v141+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1344), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+40)) = uint8(v113)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v44)+600))
	v153 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+56)) = v153
	*(*int64)(unsafe.Add(mBase, uint32(v150)+48)) = v152
	v157 = v113 & int32(255)
	if v157 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v163 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v157-int32(114)))
	goto L32
L31:
	;
	v163 = v153
	goto L32
L32:
	;
	if v163 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_finish_sync_worker(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v177 = F_palloc(m, int32(64))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L11
	} else {
		goto L37
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+36))
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[5]))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v190)))
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+328)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v44)+324)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v44)+320)) = v184
	v203 = F_pg_snprintf(m, v177, int32(64), int32(_a_F_TablesyncWorkerMain_3), v44+int32(320))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+36))
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v216 = int32(1)
	v222 = m.T0[v210].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v207, v216, v216, v130&v216, v177, v44+int32(608))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7])) = v222
	if v222 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+36))
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_ReplicationOriginNameForLogicalRep(m, v269, v266, v44+int32(528))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L48
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L45
	}
L45:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v44)+608))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v244
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_4), v44)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1381), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+40)))
	switch v280 - int32(100) {
	case 0:
		goto L52
	default:
		v294 = v279
		goto L51
	case 2:
		goto L50
	}
L49:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	if v469 != 0 {
		goto L79
	} else {
		goto L80
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L11
	} else {
		goto L75
	}
L51:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+56)) = int32(1)
	if v295 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	F_ReplicationSlotDropAtPubNode(m, v288, v177, int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v294 = v293
	goto L51
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v303+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1430), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L11
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v312)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+48)) = int64(0)
	v317 = int32(100)
	*(*uint8)(unsafe.Add(mBase, uint32(v312)+40)) = uint8(v317)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_StartTransactionCommand(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L11
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v326)+48))
	v328 = int32(*(*int8)(unsafe.Add(mBase, uint32(v326)+40)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v326)+36))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_UpdateSubscriptionRelState(m, v330, v329, v328, v327, int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v343 = v44 + int32(528)
	v345 = F_replorigin_by_name(m, v343, int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	if v345 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v353 = F_replorigin_create(m, v343)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L11
	} else {
		goto L64
	}
L62:
	;
	v355 = v52
	v356 = v345
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L65
	}
L64:
	;
	v355 = v353
	v356 = v353
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v368 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_StartTransactionCommand(m)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L11
	} else {
		goto L67
	}
L67:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v377)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v384 = F_table_open(m, v378, int32(3))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v396 = int32(0)
	v398 = m.T0[v388].(func(*base.Module, int32, int32, int32, int32) int32)(m, v394, int32(_a_F_TablesyncWorkerMain_5), v396, v396)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v398)))
	if v400 == int32(1) {
		goto L49
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L11
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v398)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+304)) = v418
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_6), v44+int32(304))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1481), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v451 = F_replorigin_by_name(m, v44+int32(528), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_replorigin_session_setup(m, v451, int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[8])) = uint16(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	v466 = F_replorigin_session_get_progress(m)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = v466
	v2855 = v49
	v2856 = v50
	v2857 = v51
	v2858 = v52
	goto L13
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v469)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L11
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
	if v476 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v476)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L11
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v398)+16))
	if v483 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v483)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L11
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v398)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	v497 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+32)))
	v500 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v500)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v508 = int32(0)
	v511 = m.T0[v501].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32) int32)(m, v507, v177, v508, v508, v498, int32(2), v39)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_LockRelationOid(m, int32(_a_F_TablesyncWorkerMain_7), int32(3))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v521 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v527 = int32(1)
	F_replorigin_advance(m, v356, v521, int64(0), v527, v527)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_UnlockRelationOid(m, int32(_a_F_TablesyncWorkerMain_7), int32(3))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_replorigin_session_setup(m, v356, int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[8])) = uint16(v356)
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+31)))
	if v550 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_SwitchToUntrustedUser(m, v554, v44+int32(516))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L11
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v384)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v575 = F_pg_class_aclcheck(m, v563, v569, int64(1))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L11
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	if v575 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v578 = int32(*(*int8)(unsafe.Add(mBase, uint32(v577)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	switch v578 - int32(73) {
	case 0, 32:
		v592 = int32(20)
		goto L106
	default:
		goto L107
	case 10:
		goto L111
	case 29:
		goto L108
	case 36:
		goto L109
	case 45:
		goto L110
	}
L103:
	;
	goto L104
L104:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v384)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v611 = int32(0)
	v613 = F_check_enable_rls(m, v606, v611, v611)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L11
	} else {
		goto L113
	}
L105:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_aclcheck_error(m, v575, v594, v595+int32(4))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L11
	} else {
		goto L112
	}
L106:
	;
	v594 = v592
	goto L105
L107:
	;
	v592 = int32(41)
	goto L106
L108:
	;
	v594 = int32(18)
	goto L105
L109:
	;
	v594 = int32(23)
	goto L105
L110:
	;
	v594 = int32(51)
	goto L105
L111:
	;
	v594 = int32(37)
	goto L105
L112:
	;
	goto L104
L113:
	;
	if v613 == int32(2) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v672 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L11
	} else {
		goto L122
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(1088))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v643 = F_GetUserNameFromId(m, v637, int32(1))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v645 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v643
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_8), v44+int32(16))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1542), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_PushActiveSnapshot(m, v672)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v686 = F_get_namespace_name(m, v681)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v384)+48))
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+712)) = v690
	v693 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+704)) = v693
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+688)) = v696
	v699 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[13]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+680)) = v699
	v702 = *(*int64)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[14]))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+672)) = v702
	*(*int32)(unsafe.Add(mBase, uint32(v44)+668)) = int32(25)
	v707 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v715 = m.T0[v708].(func(*base.Module, int32) int32)(m, v714)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v718 = v688 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+640)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+636)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v726 = v44 + int32(720)
	F_initStringInfo(m, v726)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v733 = F_quote_literal_cstr(m, v686)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v739 = F_quote_literal_cstr(m, v718)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+292)) = v739
	*(*int32)(unsafe.Add(mBase, uint32(v44)+288)) = v733
	F_appendStringInfo(m, v726, int32(_a_F_TablesyncWorkerMain_9), v44+int32(288))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v765 = m.T0[v754].(func(*base.Module, int32, int32, int32, int32) int32)(m, v760, v761, int32(3), v44+int32(704))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v765)))
	if v767 != int32(2) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L11
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v765)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v813 = F_MakeTupleTableSlot(m, v807, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L11
	} else {
		goto L138
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+280)) = v785
	*(*int32)(unsafe.Add(mBase, uint32(v44)+276)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+272)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_11), v44+int32(272))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(860), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L11
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v765)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v822 = F_tuplestore_gettupleslot(m, v815, int32(1), int32(0), v813)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L11
	} else {
		goto L139
	}
L139:
	;
	if v822 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L11
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v861 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813)+6)))
	if v861 <= int32(0) {
		goto L147
	} else {
		goto L148
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(67137668))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+260)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+256)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_13), v44+int32(256))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(867), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v813, int32(1))
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L11
	} else {
		goto L150
	}
L148:
	;
	v872 = v861
	goto L149
L149:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v813)+16))
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v873)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+632)) = v874
	if base.I32_extend16_s(v872) <= int32(1) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v871 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v813)+6)))
	v872 = v871
	goto L149
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v813, int32(2))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L11
	} else {
		goto L154
	}
L152:
	;
	v887 = v873
	goto L153
L153:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v887)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+656)) = uint8(v888)
	v890 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813)+6)))
	if v890 <= int32(2) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v813)+16))
	v887 = v886
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v813, int32(3))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L11
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v813)+16))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)) = uint8(v901)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_ExecDropSingleTupleTableSlot(m, v813)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L11
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v765)+8))
	if v909 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v909)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L11
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v765)+12))
	if v916 != 0 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v916)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L11
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v765)+16))
	if v923 != 0 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v923)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L11
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v765)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L11
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	v936 = int32(0)
	v938 = base.B2i32(v715 < int32(_a_F_TablesyncWorkerMain_14))
	if v715 < int32(_a_F_TablesyncWorkerMain_14) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1302 = v44 + int32(720)
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1302)))
	v1304 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1303))) = uint8(v1304)
	*(*int32)(unsafe.Add(mBase, uint32(v1302)+12)) = v1304
	*(*int32)(unsafe.Add(mBase, uint32(v1302)+4)) = v1304
	goto L231
L174:
	;
	v1274 = v51
	v1278 = v936
	v1288 = int32(0)
	goto L173
L175:
	;
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+664)) = int32(22)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v946 = F_makeStringInfo(m)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L11
	} else {
		goto L177
	}
L177:
	;
	v949 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_GetPublicationsStr(m, v950, v946, int32(1))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v963 = v44 + int32(720)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v963)))
	v965 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v964))) = uint8(v965)
	*(*int32)(unsafe.Add(mBase, uint32(v963)+12)) = v965
	*(*int32)(unsafe.Add(mBase, uint32(v963)+4)) = v965
	goto L179
L179:
	;
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v946)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+244)) = v971
	*(*int32)(unsafe.Add(mBase, uint32(v44)+240)) = v976
	F_appendStringInfo(m, v963, int32(_a_F_TablesyncWorkerMain_15), v44+int32(240))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L11
	} else {
		goto L180
	}
L180:
	;
	v985 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v985)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v992 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v997 = m.T0[v986].(func(*base.Module, int32, int32, int32, int32) int32)(m, v992, v993, int32(1), v44+int32(664))
	mBase = m.M
	v998 = m.ExcPending
	if v998 != 0 {
		goto L11
	} else {
		goto L181
	}
L181:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v997)))
	if v999 != int32(2) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L11
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1044 = *(*int64)(unsafe.Add(mBase, uint32(v1039)+40))
	if int64(2) <= v1044 {
		goto L189
	} else {
		goto L190
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L11
	} else {
		goto L186
	}
L186:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v997)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+232)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v44)+228)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+224)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_16), v44+int32(224))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L11
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(920), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L11
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L11
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v997)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1088 = F_MakeTupleTableSlot(m, v1082, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1089 = m.ExcPending
	if v1089 != 0 {
		goto L11
	} else {
		goto L196
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(1088))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+36)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+32)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_17), v44+int32(32))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L11
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(934), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1097 = F_tuplestore_gettupleslot(m, v1090, int32(1), int32(0), v1088)
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L11
	} else {
		goto L197
	}
L197:
	;
	if v1097 != 0 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v1099 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1088)+6)))
	if v1099 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L199:
	;
	v1214 = v51
	v1218 = v936
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1214
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_ExecDropSingleTupleTableSlot(m, v1088)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L11
	} else {
		goto L217
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1088, int32(1))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L11
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+20))
	v1110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1109))))
	if v1110 != 0 {
		v1179 = v51
		v1183 = v936
		goto L205
	} else {
		goto L206
	}
L204:
	;
	goto L203
L205:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+8))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1179
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	m.T0[v1203].(func(*base.Module, int32))(m, v1088)
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L11
	} else {
		goto L216
	}
L206:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+16))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1117 = F_pg_detoast_datum(m, v1112)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L11
	} else {
		goto L207
	}
L207:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+16))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+8))
	if v1120 == int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1117)+4))
	v1130 = (v1123<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L210
L209:
	;
	v1130 = v1120
	goto L210
L210:
	;
	if v1119 <= int32(0) {
		v1179 = v51
		v1183 = v936
		goto L205
	} else {
		goto L211
	}
L211:
	;
	v1139 = v51
	v1141 = int32(0)
	v1143 = v936
	goto L212
L212:
	;
	v1165 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1117+v1130+v1141<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1139
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1170 = F_bms_add_member(m, v1143, v1165)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L11
	} else {
		goto L214
	}
L213:
	;
	v1179 = v1170
	v1183 = v1170
	goto L205
L214:
	;
	v1173 = v1141 + int32(1)
	if v1173 != v1119 {
		v1139 = v1170
		v1141 = v1173
		v1143 = v1170
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v1214 = v1179
	v1218 = v1183
	goto L200
L217:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v997)+8))
	if v1243 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1214
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v1243)
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L11
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	if v1250 != 0 {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	goto L220
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1214
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v1250)
	mBase = m.M
	v1256 = m.ExcPending
	if v1256 != 0 {
		goto L11
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v997)+16))
	if v1257 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	goto L224
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1214
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v1257)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L11
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1214
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v997)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L11
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v1274 = v1214
	v1278 = v1218
	v1288 = v946
	goto L173
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v1302, int32(_a_F_TablesyncWorkerMain_18))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	v1318 = base.B2i32(v715 < int32(_a_F_TablesyncWorkerMain_19))
	if v715 < int32(_a_F_TablesyncWorkerMain_19) {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1330 = int32(4)
	goto L235
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v44+int32(720), int32(_a_F_TablesyncWorkerMain_20))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L11
	} else {
		goto L236
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+216)) = v1335
	*(*int32)(unsafe.Add(mBase, uint32(v44)+208)) = v1335
	if base.Ui32(v715-int32(_a_F_TablesyncWorkerMain_21)) < base.Ui32(int32(_a_F_TablesyncWorkerMain_22)) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v1330 = int32(5)
	goto L235
L237:
	;
	v1344 = int32(_a_F_TablesyncWorkerMain_23)
	goto L239
L238:
	;
	v1344 = int32(_a_F_TablesyncWorkerMain_24)
	goto L239
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+212)) = v1344
	F_appendStringInfo(m, v44+int32(720), int32(_a_F_TablesyncWorkerMain_25), v44+int32(208))
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L11
	} else {
		goto L240
	}
L240:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v1354)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1361 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v1365 = m.T0[v1355].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1361, v1362, v1330, v44+int32(672))
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L11
	} else {
		goto L241
	}
L241:
	;
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1365)))
	if v1367 != int32(2) {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L11
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1412 = F_palloc0(m, int32(_a_F_TablesyncWorkerMain_26))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L11
	} else {
		goto L249
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L11
	} else {
		goto L246
	}
L246:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+200)) = v1385
	*(*int32)(unsafe.Add(mBase, uint32(v44)+196)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+192)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_11), v44+int32(192))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1001), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L11
	} else {
		goto L248
	}
L248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+648)) = v1412
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1420 = F_palloc0(m, int32(_a_F_TablesyncWorkerMain_26))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L11
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+660)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+652)) = v1420
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1431 = F_MakeTupleTableSlot(m, v1425, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L11
	} else {
		goto L251
	}
L251:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1440 = F_tuplestore_gettupleslot(m, v1433, int32(1), int32(0), v1431)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L11
	} else {
		goto L252
	}
L252:
	;
	v1442 = int32(0)
	if v1440 != 0 {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1453 = v1442
	v1455 = v1442
	goto L256
L254:
	;
	v1630 = v1442
	v1632 = v1442
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_ExecDropSingleTupleTableSlot(m, v1431)
	mBase = m.M
	v1655 = m.ExcPending
	if v1655 != 0 {
		goto L11
	} else {
		goto L299
	}
L256:
	;
	v1473 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1431)+6)))
	if v1473 <= int32(0) {
		goto L258
	} else {
		goto L259
	}
L257:
	;
	v1630 = v1602
	v1632 = v1603
	goto L255
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1431, int32(1))
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L11
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if v1278 != 0 {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	goto L260
L262:
	;
	v1606 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+8))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1606)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	m.T0[v1607].(func(*base.Module, int32))(m, v1431)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L11
	} else {
		goto L296
	}
L263:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+16))
	v1484 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1483))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1489 = F_bms_is_member(m, v1484, v1278)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L11
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1431)+6)))
	if v1493 <= int32(1) {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	if v1489 == int32(0) {
		v1602 = v1453
		v1603 = v1455
		goto L262
	} else {
		goto L267
	}
L267:
	;
	goto L265
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1431, int32(2))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L11
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+16))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1509 = F_text_to_cstring(m, v1504)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L11
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	v1511 = int32(2)
	v1512 = v1453 << (uint(v1511) % 32)
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	*(*int32)(unsafe.Add(mBase, uint32(v1512+v1513))) = v1509
	v1516 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1431)+6)))
	if v1516 <= v1511 {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1431, int32(3))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L11
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v44)+652))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+16))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(v1528)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1526+v1512))) = v1529
	v1531 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1431)+6)))
	if v1531 <= int32(3) {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1431, int32(4))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L11
	} else {
		goto L280
	}
L278:
	;
	goto L279
L279:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+16))
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(v1541)+12))
	if v1542 != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	goto L279
L281:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v44)+660))
	v1548 = F_bms_add_member(m, v1547, v1453)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L11
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	if (v1455|v1318)&int32(1) != 0 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+660)) = v1548
	goto L283
L285:
	;
	v1569 = v1455 | base.B2i32(int32(_a_F_TablesyncWorkerMain_27) < v715)
	goto L287
L286:
	;
	v1555 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1431)+6)))
	if v1555 <= int32(4) {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	v1571 = v1453 + int32(1)
	if v1571 < int32(1664) {
		v1602 = v1571
		v1603 = v1569
		goto L262
	} else {
		goto L292
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1431, int32(5))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L11
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+16))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+16))
	v1569 = base.B2i32(v1566 != int32(0))
	goto L287
L291:
	;
	goto L290
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L11
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v686
	F_errmsg_internal(m, int32(_a_F_TablesyncWorkerMain_28), v44+int32(48))
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L11
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1049), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1601 = m.ExcPending
	if v1601 != 0 {
		goto L11
	} else {
		goto L295
	}
L295:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L296:
	;
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1621 = F_tuplestore_gettupleslot(m, v1614, int32(1), int32(0), v1431)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L11
	} else {
		goto L297
	}
L297:
	;
	if v1621 != 0 {
		v1453 = v1602
		v1455 = v1603
		goto L256
	} else {
		goto L298
	}
L298:
	;
	goto L257
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+644)) = v1630
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+8))
	if v1657 != 0 {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v1657)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L11
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+12))
	if v1664 != 0 {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	goto L302
L304:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v1664)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L11
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+16))
	if v1671 != 0 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	goto L306
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v1671)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L11
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v1365)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L11
	} else {
		goto L312
	}
L311:
	;
	goto L310
L312:
	;
	v1684 = int32(0)
	if v938 == v1684 {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1692 = v44 + int32(720)
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1692)))
	v1694 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1693))) = uint8(v1694)
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+12)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+4)) = v1694
	goto L316
L314:
	;
	v1934 = v1684
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	F_pfree(m, v1965)
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L11
	} else {
		goto L363
	}
L316:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1288)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+180)) = v1700
	*(*int32)(unsafe.Add(mBase, uint32(v44)+176)) = v1705
	F_appendStringInfo(m, v1692, int32(_a_F_TablesyncWorkerMain_29), v44+int32(176))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L11
	} else {
		goto L317
	}
L317:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1714)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1721 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v44)+720))
	v1726 = m.T0[v1715].(func(*base.Module, int32, int32, int32, int32) int32)(m, v1721, v1722, int32(1), v44+int32(668))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L11
	} else {
		goto L318
	}
L318:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	if v1728 != int32(2) {
		goto L319
	} else {
		goto L320
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L11
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1767 = F_MakeTupleTableSlot(m, v1761, int32(_a_F_TablesyncWorkerMain_10))
	mBase = m.M
	v1768 = m.ExcPending
	if v1768 != 0 {
		goto L11
	} else {
		goto L325
	}
L322:
	;
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+168)) = v1739
	*(*int32)(unsafe.Add(mBase, uint32(v44)+164)) = v718
	*(*int32)(unsafe.Add(mBase, uint32(v44)+160)) = v686
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_30), v44+int32(160))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L11
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1099), int32(_a_F_TablesyncWorkerMain_12))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L11
	} else {
		goto L324
	}
L324:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L325:
	;
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1776 = F_tuplestore_gettupleslot(m, v1769, int32(1), int32(0), v1767)
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L11
	} else {
		goto L327
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_ExecDropSingleTupleTableSlot(m, v1767)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L11
	} else {
		goto L348
	}
L327:
	;
	if v1776 == int32(0) {
		v1868 = v1684
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1780 = v1684
	goto L329
L329:
	;
	v1807 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1767)+6)))
	if v1807 <= int32(0) {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	v1868 = v1849
	goto L326
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_slot_getsomeattrs_int(m, v1767, int32(1))
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L11
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v1767)+20))
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1817))))
	if v1818 == int32(1) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	goto L333
L335:
	;
	if v1780 == int32(0) {
		goto L338
	} else {
		goto L339
	}
L336:
	;
	goto L337
L337:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1767)+16))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1831)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1837 = F_text_to_cstring(m, v1832)
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L11
	} else {
		goto L342
	}
L338:
	;
	v1868 = int32(0)
	goto L326
L339:
	;
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_list_free_deep(m, v1780)
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L11
	} else {
		goto L341
	}
L341:
	;
	v1868 = int32(0)
	goto L326
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1843 = F_makeString(m, v1837)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L11
	} else {
		goto L343
	}
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1849 = F_lappend(m, v1780, v1843)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L11
	} else {
		goto L344
	}
L344:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(v1767)+8))
	v1852 = *(*int32)(unsafe.Add(mBase, uint32(v1851)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	m.T0[v1852].(func(*base.Module, int32))(m, v1767)
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L11
	} else {
		goto L345
	}
L345:
	;
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1866 = F_tuplestore_gettupleslot(m, v1859, int32(1), int32(0), v1767)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L11
	} else {
		goto L346
	}
L346:
	;
	if v1866 != 0 {
		v1780 = v1849
		goto L329
	} else {
		goto L347
	}
L347:
	;
	goto L330
L348:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+8))
	if v1901 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v1901)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L11
	} else {
		goto L352
	}
L350:
	;
	goto L351
L351:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+12))
	if v1908 != 0 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	goto L351
L353:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v1908)
	mBase = m.M
	v1914 = m.ExcPending
	if v1914 != 0 {
		goto L11
	} else {
		goto L356
	}
L354:
	;
	goto L355
L355:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1726)+16))
	if v1915 != 0 {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	goto L355
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v1915)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L11
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v1726)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L11
	} else {
		goto L361
	}
L360:
	;
	goto L359
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_free_attrmap(m, v1288)
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
		goto L11
	} else {
		goto L362
	}
L362:
	;
	v1934 = v1868
	goto L315
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_logicalrep_relmap_update(m, v44+int32(632))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L11
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v44)+632))
	v1982 = F_logicalrep_rel_open(m, v1980, int32(0))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L11
	} else {
		goto L365
	}
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v1989 = v44 + int32(616)
	F_initStringInfo(m, v1989)
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L11
	} else {
		goto L366
	}
L366:
	;
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)))
	v1995 = int32(0)
	if (base.B2i32(v1992 != int32(114))|base.B2i32(v1934 != v1995)|v1632)&int32(1) == v1995 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v44)+636))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v44)+640))
	v2009 = F_quote_qualified_identifier(m, v2007, v2008)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L11
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v44+int32(616), int32(_a_F_TablesyncWorkerMain_31))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L11
	} else {
		goto L384
	}
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+144)) = v2009
	F_appendStringInfo(m, v1989, int32(_a_F_TablesyncWorkerMain_32), v44+int32(144))
	mBase = m.M
	v2020 = m.ExcPending
	if v2020 != 0 {
		goto L11
	} else {
		goto L371
	}
L371:
	;
	v2021 = int32(_a_F_TablesyncWorkerMain_33)
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2022 == int32(0) {
		v2432 = v2021
		goto L14
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v1989, int32(_a_F_TablesyncWorkerMain_34))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L11
	} else {
		goto L373
	}
L373:
	;
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2032 <= int32(0) {
		goto L15
	} else {
		goto L374
	}
L374:
	;
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v2035)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2041 = F_quote_identifier(m, v2036)
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L11
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v1989, v2041)
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L11
	} else {
		goto L376
	}
L376:
	;
	v2049 = int32(1)
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2050 <= v2049 {
		goto L15
	} else {
		goto L377
	}
L377:
	;
	v2059 = v2049
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2085 = v44 + int32(616)
	F_appendStringInfoString(m, v2085, int32(_a_F_TablesyncWorkerMain_35))
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L11
	} else {
		goto L380
	}
L379:
	;
	goto L15
L380:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v2089+v2059<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2098 = F_quote_identifier(m, v2093)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L11
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v2085, v2098)
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L11
	} else {
		goto L382
	}
L382:
	;
	v2107 = v2059 + int32(1)
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2107 < v2108 {
		v2059 = v2107
		goto L378
	} else {
		goto L383
	}
L383:
	;
	goto L379
L384:
	;
	v2119 = int32(0)
	v2120 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2120 <= v2119 {
		goto L16
	} else {
		goto L385
	}
L385:
	;
	v2129 = v2119
	goto L386
L386:
	;
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v44)+648))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2150+v2129<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2159 = F_quote_identifier(m, v2154)
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L11
	} else {
		goto L388
	}
L387:
	;
	goto L16
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2166 = v44 + int32(616)
	F_appendStringInfoString(m, v2166, v2159)
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	if v2129 < v2169-int32(1) {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v2166, int32(_a_F_TablesyncWorkerMain_35))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L11
	} else {
		goto L393
	}
L391:
	;
	v2181 = v2169
	goto L392
L392:
	;
	v2183 = v2129 + int32(1)
	if v2183 < v2181 {
		v2129 = v2183
		goto L386
	} else {
		goto L394
	}
L393:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v44)+644))
	v2181 = v2180
	goto L392
L394:
	;
	goto L387
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_DisableSubscriptionAndExit(m)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L11
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_AbortOutOfAnyTransaction(m)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L11
	} else {
		goto L399
	}
L398:
	;
	v3099 = v49
	v3100 = v50
	v3101 = v51
	v3102 = v52
	v3112 = int32(0)
	goto L12
L399:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2208)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_pgstat_report_subscription_error(m, v2209, int32(0))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v51
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v52)
	F_pg_re_throw(m)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L11
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	v2259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+657)))
	if v2259 == int32(114) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v2255, int32(_a_F_TablesyncWorkerMain_36))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L11
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v44)+636))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v44)+640))
	v2275 = F_quote_qualified_identifier(m, v2273, v2274)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L11
	} else {
		goto L407
	}
L406:
	;
	goto L405
L407:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2282 = v44 + int32(616)
	F_appendStringInfoString(m, v2282, v2275)
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L11
	} else {
		goto L408
	}
L408:
	;
	v2285 = int32(_a_F_TablesyncWorkerMain_37)
	if v1934 == int32(0) {
		v2432 = v2285
		goto L14
	} else {
		goto L409
	}
L409:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+12))
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v2288)))
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2289)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+128)) = v2290
	F_appendStringInfo(m, v2282, int32(_a_F_TablesyncWorkerMain_38), v44+int32(128))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L11
	} else {
		goto L410
	}
L410:
	;
	v2302 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+4))
	if int32(2) <= v2302 {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v2311 = int32(1)
	goto L414
L412:
	;
	goto L413
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_list_free_deep(m, v1934)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L11
	} else {
		goto L418
	}
L414:
	;
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+12))
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2332+v2311<<(uint(int32(2))%32))))
	v2337 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v2337
	F_appendStringInfo(m, v44+int32(616), int32(_a_F_TablesyncWorkerMain_39), v44+int32(112))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L11
	} else {
		goto L416
	}
L415:
	;
	goto L413
L416:
	;
	v2351 = v2311 + int32(1)
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v1934)+4))
	if v2351 < v2352 {
		v2311 = v2351
		goto L414
	} else {
		goto L417
	}
L417:
	;
	goto L415
L418:
	;
	v2432 = v2285
	goto L14
L419:
	;
	v2432 = v2021
	goto L14
L420:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2459)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2467 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2468 = m.T0[v2460].(func(*base.Module, int32) int32)(m, v2467)
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L11
	} else {
		goto L422
	}
L421:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2517 = *(*int32)(unsafe.Add(mBase, uint32(v2516)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2523 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v44)+616))
	v2525 = int32(0)
	v2527 = m.T0[v2517].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2523, v2524, v2525, v2525)
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L11
	} else {
		goto L429
	}
L422:
	;
	if v2468 < int32(_a_F_TablesyncWorkerMain_40) {
		v2513 = v50
		v2514 = int32(0)
		goto L421
	} else {
		goto L423
	}
L423:
	;
	v2474 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v2475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2474)+26)))
	if v2475 != int32(1) {
		v2513 = v50
		v2514 = int32(0)
		goto L421
	} else {
		goto L424
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_appendStringInfoString(m, v2455, int32(_a_F_TablesyncWorkerMain_41))
	mBase = m.M
	v2484 = m.ExcPending
	if v2484 != 0 {
		goto L11
	} else {
		goto L425
	}
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2490 = F_makeString(m, int32(_a_F_TablesyncWorkerMain_42))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L11
	} else {
		goto L426
	}
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2498 = F_makeDefElem(m, int32(_a_F_TablesyncWorkerMain_43), v2490, int32(-1))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L11
	} else {
		goto L427
	}
L427:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+612)) = v2498
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = v2498
	v2510 = F_list_make1_impl(m, int32(1), v44+int32(108))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L11
	} else {
		goto L428
	}
L428:
	;
	v2513 = v2510
	v2514 = v2510
	goto L421
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(v44)+616))
	F_pfree(m, v2533)
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L11
	} else {
		goto L430
	}
L430:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2527)))
	if v2536 != int32(4) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L11
	} else {
		goto L434
	}
L432:
	;
	goto L433
L433:
	;
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+8))
	if v2576 != 0 {
		goto L438
	} else {
		goto L439
	}
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2553 = m.ExcPending
	if v2553 != 0 {
		goto L11
	} else {
		goto L435
	}
L435:
	;
	v2554 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2559 = *(*int64)(unsafe.Add(mBase, uint32(v44)+636))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+104)) = v2554
	*(*int64)(unsafe.Add(mBase, uint32(v44)+96)) = v2559
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_44), v44+int32(96))
	mBase = m.M
	v2566 = m.ExcPending
	if v2566 != 0 {
		goto L11
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1266), int32(_a_F_TablesyncWorkerMain_45))
	mBase = m.M
	v2575 = m.ExcPending
	if v2575 != 0 {
		goto L11
	} else {
		goto L437
	}
L437:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v2576)
	mBase = m.M
	v2582 = m.ExcPending
	if v2582 != 0 {
		goto L11
	} else {
		goto L441
	}
L439:
	;
	goto L440
L440:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+12))
	if v2583 != 0 {
		goto L442
	} else {
		goto L443
	}
L441:
	;
	goto L440
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v2583)
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L11
	} else {
		goto L445
	}
L443:
	;
	goto L444
L444:
	;
	v2590 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+16))
	if v2590 != 0 {
		goto L446
	} else {
		goto L447
	}
L445:
	;
	goto L444
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v2590)
	mBase = m.M
	v2596 = m.ExcPending
	if v2596 != 0 {
		goto L11
	} else {
		goto L449
	}
L447:
	;
	goto L448
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v2527)
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L11
	} else {
		goto L450
	}
L449:
	;
	goto L448
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2608 = F_makeStringInfo(m)
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L11
	} else {
		goto L451
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[15])) = v2608
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2616 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v2617 = m.ExcPending
	if v2617 != 0 {
		goto L11
	} else {
		goto L452
	}
L452:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2623 = int32(0)
	v2626 = F_addRangeTableEntryForRelation(m, v2616, v384, int32(1), v2623, v2623, v2623)
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L11
	} else {
		goto L453
	}
L453:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v1982)+12))
	if v2628 <= int32(0) {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2713 = int32(0)
	v2717 = F_BeginCopyFrom(m, v2616, v384, v2713, v2713, v2713, int32(1023), v2682, v2514)
	mBase = m.M
	v2718 = m.ExcPending
	if v2718 != 0 {
		goto L11
	} else {
		goto L463
	}
L455:
	;
	v2682 = int32(0)
	v2684 = v49
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2632 = int32(0)
	v2636 = v49
	v2640 = v2632
	v2642 = v2632
	goto L458
L458:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v1982)+16))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2661+v2640<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2636
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2670 = F_makeString(m, v2665)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L11
	} else {
		goto L460
	}
L459:
	;
	v2682 = v2676
	v2684 = v2676
	goto L454
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2636
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2676 = F_lappend(m, v2642, v2670)
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L11
	} else {
		goto L461
	}
L461:
	;
	v2679 = v2640 + int32(1)
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v1982)+12))
	if v2679 < v2680 {
		v2636 = v2676
		v2640 = v2679
		v2642 = v2676
		goto L458
	} else {
		goto L462
	}
L462:
	;
	goto L459
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2723 = F_CopyFrom(m, v2717)
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L11
	} else {
		goto L464
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_logicalrep_rel_close(m, v1982, int32(0))
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L11
	} else {
		goto L465
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2737 = m.ExcPending
	if v2737 != 0 {
		goto L11
	} else {
		goto L466
	}
L466:
	;
	v2739 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(v2739)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	v2746 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v2748 = int32(0)
	v2750 = m.T0[v2740].(func(*base.Module, int32, int32, int32, int32) int32)(m, v2746, int32(_a_F_TablesyncWorkerMain_46), v2748, v2748)
	mBase = m.M
	v2751 = m.ExcPending
	if v2751 != 0 {
		goto L11
	} else {
		goto L467
	}
L467:
	;
	v2752 = *(*int32)(unsafe.Add(mBase, uint32(v2750)))
	if v2752 != int32(1) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L11
	} else {
		goto L471
	}
L469:
	;
	goto L470
L470:
	;
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(v2750)+8))
	if v2790 != 0 {
		goto L475
	} else {
		goto L476
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errcode(m, int32(100663808))
	mBase = m.M
	v2769 = m.ExcPending
	if v2769 != 0 {
		goto L11
	} else {
		goto L472
	}
L472:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2750)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+80)) = v2770
	F_errmsg(m, int32(_a_F_TablesyncWorkerMain_47), v44+int32(80))
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L11
	} else {
		goto L473
	}
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1554), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2789 = m.ExcPending
	if v2789 != 0 {
		goto L11
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v2790)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L11
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v2750)+12))
	if v2797 != 0 {
		goto L479
	} else {
		goto L480
	}
L478:
	;
	goto L477
L479:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_tuplestore_end(m, v2797)
	mBase = m.M
	v2803 = m.ExcPending
	if v2803 != 0 {
		goto L11
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2804 = *(*int32)(unsafe.Add(mBase, uint32(v2750)+16))
	if v2804 != 0 {
		goto L483
	} else {
		goto L484
	}
L482:
	;
	goto L481
L483:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_FreeTupleDesc(m, v2804)
	mBase = m.M
	v2810 = m.ExcPending
	if v2810 != 0 {
		goto L11
	} else {
		goto L486
	}
L484:
	;
	goto L485
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_pfree(m, v2750)
	mBase = m.M
	v2816 = m.ExcPending
	if v2816 != 0 {
		goto L11
	} else {
		goto L487
	}
L486:
	;
	goto L485
L487:
	;
	if v550 == int32(0) {
		goto L488
	} else {
		goto L489
	}
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_RestoreUserContext(m, v44+int32(516))
	mBase = m.M
	v2826 = m.ExcPending
	if v2826 != 0 {
		goto L11
	} else {
		goto L491
	}
L489:
	;
	goto L490
L490:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_relation_close(m, v384, int32(0))
	mBase = m.M
	v2833 = m.ExcPending
	if v2833 != 0 {
		goto L11
	} else {
		goto L492
	}
L491:
	;
	goto L490
L492:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L11
	} else {
		goto L493
	}
L493:
	;
	v2841 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2842 = *(*int64)(unsafe.Add(mBase, uint32(v2841)+48))
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2841)+36))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2841)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2684
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2513
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v1274
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v355)
	F_UpdateSubscriptionRelState(m, v2844, v2843, int32(102), v2842, int32(0))
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L11
	} else {
		goto L494
	}
L494:
	;
	v2855 = v2684
	v2856 = v2513
	v2857 = v1274
	v2858 = v355
	goto L13
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v2892 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L11
	} else {
		goto L496
	}
L496:
	;
	if v2892 != 0 {
		goto L497
	} else {
		goto L498
	}
L497:
	;
	v2894 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+72)) = uint32(v2894)
	v2901 = int64(base.Ui64(v2894) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v44)+68)) = uint32(v2901)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+64)) = v44 + int32(528)
	F_errmsg_internal(m, int32(_a_F_TablesyncWorkerMain_48), v44-int32(-64))
	mBase = m.M
	v2910 = m.ExcPending
	if v2910 != 0 {
		goto L11
	} else {
		goto L500
	}
L498:
	;
	goto L499
L499:
	;
	v2922 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v2922)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v2922)+56)) = int32(1)
	if v2923 != 0 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	F_errfinish(m, int32(_a_F_TablesyncWorkerMain_1), int32(1581), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2919 = m.ExcPending
	if v2919 != 0 {
		goto L11
	} else {
		goto L501
	}
L501:
	;
	goto L499
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v2931 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	F_s_lock(m, v2931+int32(56), int32(_a_F_TablesyncWorkerMain_1), int32(1586), int32(_a_F_TablesyncWorkerMain_2))
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L11
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2940 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2941 = int32(119)
	*(*uint8)(unsafe.Add(mBase, uint32(v2940)+40)) = uint8(v2941)
	v2943 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v2940)+56)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2940)+48)) = v2943
	goto L506
L505:
	;
	goto L504
L506:
	;
	v2975 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[16]))
	if v2975 != 0 {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v3092 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	F_LWLockRelease(m, v3092+int32(_a_F_TablesyncWorkerMain_49))
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L11
	} else {
		goto L536
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	F_ProcessInterrupts(m)
	mBase = m.M
	v2981 = m.ExcPending
	if v2981 != 0 {
		goto L11
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	v2983 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v2984 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2983)+40)))
	if v2984 == int32(99) {
		v3099 = v2855
		v3100 = v2856
		v3101 = v2857
		v3102 = v2858
		v3112 = v177
		goto L12
	} else {
		goto L512
	}
L511:
	;
	goto L510
L512:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v2992 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	v2996 = F_LWLockAcquire(m, v2992+int32(_a_F_TablesyncWorkerMain_49), int32(1))
	mBase = m.M
	v2997 = m.ExcPending
	if v2997 != 0 {
		goto L11
	} else {
		goto L513
	}
L513:
	;
	v2999 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	v3000 = *(*int32)(unsafe.Add(mBase, uint32(v2999)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v3005 = int32(0)
	v3011 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[18]))
	if v3011 <= v3005 {
		v3043 = v3005
		goto L515
	} else {
		goto L516
	}
L514:
	;
	if v3043 != 0 {
		goto L525
	} else {
		goto L526
	}
L515:
	;
	goto L514
L516:
	;
	v3015 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[19]))
	v3021 = v3005
	goto L517
L517:
	;
	v3026 = v3015 + int32(16) + v3021*int32(112)
	v3027 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3026)+16)))
	if v3027 != int32(1) {
		goto L519
	} else {
		goto L520
	}
L518:
	;
	v3043 = int32(0)
	goto L515
L519:
	;
	v3038 = v3021 + int32(1)
	if v3038 != v3011 {
		v3021 = v3038
		goto L517
	} else {
		goto L524
	}
L520:
	;
	v3030 = *(*int32)(unsafe.Add(mBase, uint32(v3026)))
	if v3030 == int32(3) {
		goto L519
	} else {
		goto L521
	}
L521:
	;
	v3033 = *(*int32)(unsafe.Add(mBase, uint32(v3026)+32))
	if v3033 != v3000 {
		goto L519
	} else {
		goto L522
	}
L522:
	;
	v3035 = *(*int32)(unsafe.Add(mBase, uint32(v3026)+36))
	if v3035 != v3005 {
		goto L519
	} else {
		goto L523
	}
L523:
	;
	v3043 = v3026
	goto L515
L524:
	;
	goto L518
L525:
	;
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v3043)+20))
	if v3047 != 0 {
		goto L528
	} else {
		goto L529
	}
L526:
	;
	goto L527
L527:
	;
	goto L507
L528:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	F_logicalrep_worker_wakeup_ptr(m, v3043)
	mBase = m.M
	v3053 = m.ExcPending
	if v3053 != 0 {
		goto L11
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v3059 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[17]))
	F_LWLockRelease(m, v3059+int32(_a_F_TablesyncWorkerMain_49))
	mBase = m.M
	v3063 = m.ExcPending
	if v3063 != 0 {
		goto L11
	} else {
		goto L532
	}
L531:
	;
	goto L530
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v3069 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[20]))
	v3073 = F_WaitLatch(m, v3069, int32(41), int32(1000), int32(134217760))
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L11
	} else {
		goto L533
	}
L533:
	;
	if v3073&int32(1) == int32(0) {
		goto L506
	} else {
		goto L534
	}
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v2855
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v2856
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v2857
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v2858)
	v3084 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v3084))) = int32(0)
	goto L535
L535:
	;
	goto L506
L536:
	;
	v3099 = v2855
	v3100 = v2856
	v3101 = v2857
	v3102 = v2858
	v3112 = v177
	goto L12
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(52)))) = v3134
	*(*int32)(unsafe.Add(mBase, uint32(v44)+740)) = v3100
	*(*int32)(unsafe.Add(mBase, uint32(v44)+736)) = v3099
	*(*int32)(unsafe.Add(mBase, uint32(v44)+744)) = v3101
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)) = uint16(v3102)
	F_pfree(m, v3112)
	mBase = m.M
	v3142 = m.ExcPending
	if v3142 != 0 {
		goto L11
	} else {
		goto L538
	}
L538:
	;
	goto L10
L539:
	;
	v3175 = int32(v3171)
	m.G0 = v44
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(v3175)+4))
	v3178 = *(*int32)(unsafe.Add(mBase, uint32(v3175)))
	v3181 = *(*int32)(unsafe.Add(mBase, uint32(v3178)))
	if v44+int32(348) == v3181 {
		goto L542
	} else {
		goto L543
	}
L540:
	;
	m.ExcPending = 1
	goto L1
L541:
	;
	if v3185 != 0 {
		goto L545
	} else {
		goto L546
	}
L542:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v3178)+4))
	v3185 = v3183
	goto L544
L543:
	;
	v3185 = int32(0)
	goto L544
L544:
	;
	goto L541
L545:
	;
	v3186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+750)))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v44)+744))
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(v44)+740))
	v3189 = *(*int32)(unsafe.Add(mBase, uint32(v44)+736))
	v49 = v3189
	v50 = v3188
	v51 = v3187
	v52 = v3186
	v53 = v3177
	v54 = v3185
	v67 = v89
	v68 = v90
	goto L3
L546:
	;
	goto L547
L547:
	;
	F___wasm_longjmp(m, v3178, v3177)
	mBase = m.M
	v3191 = m.ExcPending
	if v3191 != 0 {
		goto L1
	} else {
		goto L548
	}
L548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L549:
	;
	F_set_apply_error_context_origin(m, v3203)
	mBase = m.M
	v3207 = m.ExcPending
	if v3207 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v3209 = int32(1)
	v3211 = v32 + int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v3211))) = uint8(v3209)
	v3216 = *(*int64)(unsafe.Add(mBase, uint32(v32+int32(56))))
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+4)) = v3208
	*(*int64)(unsafe.Add(mBase, uint32(v3211)+8)) = v3216
	v3221 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v3223 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(v3223)+24))
	v3225 = m.T0[v3224].(func(*base.Module, int32) int32)(m, v3221)
	mBase = m.M
	v3226 = m.ExcPending
	if v3226 != 0 {
		goto L1
	} else {
		goto L553
	}
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+28)) = v3275
	v3278 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[1]))
	*(*uint8)(unsafe.Add(mBase, uint32(v3278)+68)) = uint8(v3274)
	v3280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3211)+32)) = uint8(v3280)
	v3282 = *(*int32)(unsafe.Add(mBase, uint32(v3273)+52))
	v3283 = F_pstrdup(m, v3282)
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L570
	}
L552:
	;
	v3264 = int32(0)
	if v3261&int32(255) != int32(102) {
		goto L567
	} else {
		goto L568
	}
L553:
	;
	if v3225 <= int32(_a_F_TablesyncWorkerMain_50) {
		goto L554
	} else {
		goto L555
	}
L554:
	;
	if int32(_a_F_TablesyncWorkerMain_51) < v3225 {
		goto L557
	} else {
		goto L558
	}
L555:
	;
	goto L556
L556:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+16)) = int32(4)
	v3252 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(v3252)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+20)) = v3253
	v3255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3252)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3211)+24)) = uint8(v3255)
	v3258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3252)+27)))
	if v3258 == int32(112) {
		v3273 = v3252
		v3274 = v3209
		v3275 = int32(_a_F_TablesyncWorkerMain_52)
		goto L551
	} else {
		goto L566
	}
L557:
	;
	v3234 = int32(2)
	goto L559
L558:
	;
	v3234 = int32(1)
	goto L559
L559:
	;
	if int32(_a_F_TablesyncWorkerMain_53) < v3225 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3237 = int32(3)
	goto L562
L561:
	;
	v3237 = v3234
	goto L562
L562:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+16)) = v3237
	v3240 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[0]))
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3240)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+20)) = v3241
	v3243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3240)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3211)+24)) = uint8(v3243)
	if v3225 < int32(_a_F_TablesyncWorkerMain_54) {
		goto L563
	} else {
		goto L564
	}
L563:
	;
	v3273 = v3240
	v3274 = int32(0)
	v3275 = int32(0)
	goto L551
L564:
	;
	goto L565
L565:
	;
	v3248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3240)+27)))
	v3261 = v3248
	v3262 = v3240
	goto L552
L566:
	;
	v3261 = v3258
	v3262 = v3252
	goto L552
L567:
	;
	v3271 = int32(_a_F_TablesyncWorkerMain_55)
	goto L569
L568:
	;
	v3271 = v3264
	goto L569
L569:
	;
	v3273 = v3262
	v3274 = v3264
	v3275 = v3271
	goto L551
L570:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3211)+36)) = v3283
	v3287 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[7]))
	v3291 = *(*int32)(unsafe.Add(mBase, _c_F_TablesyncWorkerMain[6]))
	v3292 = *(*int32)(unsafe.Add(mBase, uint32(v3291)+32))
	v3293 = m.T0[v3292].(func(*base.Module, int32, int32) int32)(m, v3287, v32+int32(8))
	mBase = m.M
	v3294 = m.ExcPending
	if v3294 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	v3295 = *(*int64)(unsafe.Add(mBase, uint32(v32)+56))
	F_start_apply(m, v3295)
	mBase = m.M
	v3297 = m.ExcPending
	if v3297 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	m.G0 = v32 + int32(128)
	F_finish_sync_worker(m)
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L1
	} else {
		goto L573
	}
L573:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
