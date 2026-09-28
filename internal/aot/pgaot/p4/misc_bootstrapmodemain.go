package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BootstrapModeMain(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v574 int64
	_ = v574
	var v575 int64
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int64
	_ = v620
	var v632 int64
	_ = v632
	var v634 int32
	_ = v634
	var v639 int64
	_ = v639
	var v643 int64
	_ = v643
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v658 int64
	_ = v658
	var v660 int32
	_ = v660
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v795 int32
	_ = v795
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int64
	_ = v827
	var v829 int64
	_ = v829
	var v831 int64
	_ = v831
	var v833 int64
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v870 int64
	_ = v870
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1093 int32
	_ = v1093
	var v1098 int32
	_ = v1098
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1124 int32
	_ = v1124
	var v1130 int64
	_ = v1130
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1230 int32
	_ = v1230
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
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
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1416 int32
	_ = v1416
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	var v1478 int32
	_ = v1478
	var v1495 int32
	_ = v1495
	var v1501 int32
	_ = v1501
	var v1505 int32
	_ = v1505
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1570 int32
	_ = v1570
	var v1576 int32
	_ = v1576
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1680 int32
	_ = v1680
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1718 int32
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1764 int32
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1773 int32
	_ = v1773
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1782 int32
	_ = v1782
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1904 int32
	_ = v1904
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1950 int32
	_ = v1950
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1963 int32
	_ = v1963
	var v1966 int32
	_ = v1966
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2005 int32
	_ = v2005
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2026 int32
	_ = v2026
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2047 int32
	_ = v2047
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2060 int32
	_ = v2060
	var v2064 int32
	_ = v2064
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2089 int32
	_ = v2089
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2131 int32
	_ = v2131
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2186 int32
	_ = v2186
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2218 int32
	_ = v2218
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2233 int32
	_ = v2233
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2257 int32
	_ = v2257
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2306 int32
	_ = v2306
	var v2307 int32
	_ = v2307
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2365 int32
	_ = v2365
	var v2369 int32
	_ = v2369
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2397 int32
	_ = v2397
	var v2400 int32
	_ = v2400
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2414 int32
	_ = v2414
	var v2427 int32
	_ = v2427
	var v2444 int32
	_ = v2444
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2457 int32
	_ = v2457
	var v2482 int32
	_ = v2482
	var v2485 int32
	_ = v2485
	var v2489 int32
	_ = v2489
	var v2491 int32
	_ = v2491
	var v2496 int32
	_ = v2496
	var v2503 int32
	_ = v2503
	var v2524 int32
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2541 int32
	_ = v2541
	var v2545 int32
	_ = v2545
	var v2575 int32
	_ = v2575
	var v2579 int32
	_ = v2579
	var v2584 int32
	_ = v2584
	var v2585 int32
	_ = v2585
	var v2593 int32
	_ = v2593
	var v2596 int32
	_ = v2596
	var v2597 int32
	_ = v2597
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2612 int32
	_ = v2612
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2642 int32
	_ = v2642
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2648 int32
	_ = v2648
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2653 int32
	_ = v2653
	var v2659 int32
	_ = v2659
	var v2662 int32
	_ = v2662
	var v2686 int32
	_ = v2686
	var v2689 int32
	_ = v2689
	var v2690 int32
	_ = v2690
	var v2710 int32
	_ = v2710
	var v2712 int32
	_ = v2712
	var v2717 int32
	_ = v2717
	var v2745 int32
	_ = v2745
	var v2746 int32
	_ = v2746
	var v2750 int32
	_ = v2750
	var v2753 int32
	_ = v2753
	var v2757 int32
	_ = v2757
	var v2760 int32
	_ = v2760
	var v2777 int32
	_ = v2777
	var v2780 int32
	_ = v2780
	var v2786 int32
	_ = v2786
	var v2787 int32
	_ = v2787
	var v2788 int32
	_ = v2788
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2795 int32
	_ = v2795
	var v2802 int32
	_ = v2802
	var v2819 int32
	_ = v2819
	var v2824 int32
	_ = v2824
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2832 int32
	_ = v2832
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2859 int32
	_ = v2859
	var v2862 int32
	_ = v2862
	var v2883 int32
	_ = v2883
	var v2886 int32
	_ = v2886
	var v2887 int32
	_ = v2887
	var v2892 int32
	_ = v2892
	var v2914 int32
	_ = v2914
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2919 int32
	_ = v2919
	var v2920 int32
	_ = v2920
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2930 int32
	_ = v2930
	var v2932 int32
	_ = v2932
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2943 int32
	_ = v2943
	var v2944 int32
	_ = v2944
	var v2945 int32
	_ = v2945
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2962 int32
	_ = v2962
	var v2963 int32
	_ = v2963
	var v2967 int32
	_ = v2967
	var v2968 int32
	_ = v2968
	var v2971 int32
	_ = v2971
	var v2972 int32
	_ = v2972
	var v2973 int32
	_ = v2973
	var v2978 int32
	_ = v2978
	var v3001 int32
	_ = v3001
	var v3002 int32
	_ = v3002
	var v3011 int32
	_ = v3011
	var v3017 int32
	_ = v3017
	var v3018 int32
	_ = v3018
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3031 int32
	_ = v3031
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3036 int32
	_ = v3036
	var v3039 int32
	_ = v3039
	var v3042 int32
	_ = v3042
	var v3047 int32
	_ = v3047
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3078 int32
	_ = v3078
	var v3100 int32
	_ = v3100
	var v3101 int32
	_ = v3101
	var v3104 int32
	_ = v3104
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3109 int32
	_ = v3109
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3120 int32
	_ = v3120
	var v3125 int32
	_ = v3125
	var v3127 int32
	_ = v3127
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3137 int32
	_ = v3137
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3158 int32
	_ = v3158
	var v3161 int32
	_ = v3161
	var v3162 int32
	_ = v3162
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3170 int32
	_ = v3170
	var v3171 int32
	_ = v3171
	var v3173 int32
	_ = v3173
	var v3175 int32
	_ = v3175
	var v3177 int32
	_ = v3177
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3187 int32
	_ = v3187
	var v3195 int32
	_ = v3195
	var v3196 int32
	_ = v3196
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3210 int32
	_ = v3210
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3229 int32
	_ = v3229
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3235 int32
	_ = v3235
	var v3236 int32
	_ = v3236
	var v3240 int32
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3242 int32
	_ = v3242
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3252 int32
	_ = v3252
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3275 int32
	_ = v3275
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3290 int32
	_ = v3290
	var v3294 int32
	_ = v3294
	var v3295 int32
	_ = v3295
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3300 int32
	_ = v3300
	var v3308 int32
	_ = v3308
	var v3309 int32
	_ = v3309
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3322 int32
	_ = v3322
	var v3325 int32
	_ = v3325
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3346 int32
	_ = v3346
	var v3348 int32
	_ = v3348
	var v3351 int32
	_ = v3351
	var v3355 int32
	_ = v3355
	var v3358 int32
	_ = v3358
	var v3359 int32
	_ = v3359
	var v3364 int32
	_ = v3364
	var v3368 int32
	_ = v3368
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3396 int32
	_ = v3396
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3407 int32
	_ = v3407
	var v3408 int32
	_ = v3408
	var v3413 int32
	_ = v3413
	var v3426 int32
	_ = v3426
	var v3443 int32
	_ = v3443
	var v3447 int32
	_ = v3447
	var v3449 int32
	_ = v3449
	var v3451 int32
	_ = v3451
	var v3452 int32
	_ = v3452
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3459 int32
	_ = v3459
	var v3460 int32
	_ = v3460
	var v3463 int32
	_ = v3463
	var v3465 int32
	_ = v3465
	var v3480 int32
	_ = v3480
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3494 int32
	_ = v3494
	var v3495 int32
	_ = v3495
	var v3517 int32
	_ = v3517
	var v3518 int32
	_ = v3518
	var v3519 int32
	_ = v3519
	var v3524 int32
	_ = v3524
	var v3528 int32
	_ = v3528
	var v3529 int32
	_ = v3529
	var v3533 int32
	_ = v3533
	var v3534 int32
	_ = v3534
	var v3539 int32
	_ = v3539
	var v3543 int32
	_ = v3543
	var v3546 int32
	_ = v3546
	var v3547 int32
	_ = v3547
	var v3571 int32
	_ = v3571
	var v3574 int32
	_ = v3574
	var v3575 int32
	_ = v3575
	var v3577 int32
	_ = v3577
	var v3578 int32
	_ = v3578
	var v3582 int32
	_ = v3582
	var v3583 int32
	_ = v3583
	var v3588 int32
	_ = v3588
	var v3601 int32
	_ = v3601
	var v3618 int32
	_ = v3618
	var v3622 int32
	_ = v3622
	var v3624 int32
	_ = v3624
	var v3628 int32
	_ = v3628
	var v3634 int32
	_ = v3634
	var v3635 int32
	_ = v3635
	var v3664 int32
	_ = v3664
	var v3668 int32
	_ = v3668
	var v3674 int32
	_ = v3674
	var v3678 int32
	_ = v3678
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3685 int32
	_ = v3685
	var v3687 int32
	_ = v3687
	var v3688 int32
	_ = v3688
	var v3690 int32
	_ = v3690
	var v3693 int32
	_ = v3693
	var v3699 int32
	_ = v3699
	var v3706 int32
	_ = v3706
	var v3707 int32
	_ = v3707
	var v3708 int32
	_ = v3708
	var v3709 int32
	_ = v3709
	var v3712 int32
	_ = v3712
	var v3714 int32
	_ = v3714
	var v3717 int32
	_ = v3717
	var v3730 int32
	_ = v3730
	var v3734 int32
	_ = v3734
	var v3737 int32
	_ = v3737
	var v3738 int32
	_ = v3738
	var v3739 int32
	_ = v3739
	var v3740 int32
	_ = v3740
	var v3741 int32
	_ = v3741
	var v3743 int32
	_ = v3743
	var v3744 int32
	_ = v3744
	var v3746 int32
	_ = v3746
	var v3749 int32
	_ = v3749
	var v3755 int32
	_ = v3755
	var v3761 int32
	_ = v3761
	var v3765 int32
	_ = v3765
	var v3768 int32
	_ = v3768
	var v3769 int32
	_ = v3769
	var v3770 int32
	_ = v3770
	var v3771 int32
	_ = v3771
	var v3772 int32
	_ = v3772
	var v3774 int32
	_ = v3774
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3780 int32
	_ = v3780
	var v3787 int32
	_ = v3787
	var v3792 int32
	_ = v3792
	var v3796 int32
	_ = v3796
	var v3801 int32
	_ = v3801
	var v3806 int32
	_ = v3806
	var v3807 int32
	_ = v3807
	var v3809 int32
	_ = v3809
	var v3812 int32
	_ = v3812
	var v3814 int32
	_ = v3814
	var v3816 int32
	_ = v3816
	var v3818 int32
	_ = v3818
	var v3821 int32
	_ = v3821
	var v3824 int32
	_ = v3824
	var v3828 int32
	_ = v3828
	var v3830 int32
	_ = v3830
	var v3833 int32
	_ = v3833
	var v3836 int32
	_ = v3836
	var v3837 int32
	_ = v3837
	var v3845 int32
	_ = v3845
	var v3850 int32
	_ = v3850
	var v3854 int32
	_ = v3854
	var v3855 int32
	_ = v3855
	var v3857 int32
	_ = v3857
	var v3858 int32
	_ = v3858
	var v3861 int32
	_ = v3861
	var v3862 int32
	_ = v3862
	var v3874 int32
	_ = v3874
	var v3895 int32
	_ = v3895
	var v3898 int32
	_ = v3898
	var v3902 int32
	_ = v3902
	var v3904 int32
	_ = v3904
	var v3905 int32
	_ = v3905
	var v3907 int32
	_ = v3907
	var v3909 int32
	_ = v3909
	var v3910 int32
	_ = v3910
	var v3911 int32
	_ = v3911
	var v3915 int32
	_ = v3915
	var v3922 int32
	_ = v3922
	var v3925 int32
	_ = v3925
	var v3926 int32
	_ = v3926
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3929 int32
	_ = v3929
	var v3939 int32
	_ = v3939
	var v3944 int32
	_ = v3944
	var v3948 int32
	_ = v3948
	var v3950 int32
	_ = v3950
	var v3983 int32
	_ = v3983
	var v3986 int32
	_ = v3986
	var v3988 int32
	_ = v3988
	var v3990 int32
	_ = v3990
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3994 int32
	_ = v3994
	var v4000 int32
	_ = v4000
	var v4001 int32
	_ = v4001
	var v4002 int32
	_ = v4002
	var v4004 int32
	_ = v4004
	var v4009 int32
	_ = v4009
	var v4014 int32
	_ = v4014
	var v4015 int32
	_ = v4015
	var v4017 int32
	_ = v4017
	var v4020 int32
	_ = v4020
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4028 int32
	_ = v4028
	var v4030 int32
	_ = v4030
	var v4032 int32
	_ = v4032
	var v4034 int32
	_ = v4034
	var v4035 int32
	_ = v4035
	var v4036 int32
	_ = v4036
	var v4042 int32
	_ = v4042
	var v4043 int32
	_ = v4043
	var v4044 int32
	_ = v4044
	var v4046 int32
	_ = v4046
	var v4051 int32
	_ = v4051
	var v4056 int32
	_ = v4056
	var v4057 int32
	_ = v4057
	var v4059 int32
	_ = v4059
	var v4063 int32
	_ = v4063
	var v4067 int32
	_ = v4067
	var v4068 int32
	_ = v4068
	var v4073 int32
	_ = v4073
	var v4076 int32
	_ = v4076
	var v4079 int64
	_ = v4079
	var v4083 int32
	_ = v4083
	var v4087 int32
	_ = v4087
	var v4091 int32
	_ = v4091
	var v4096 int32
	_ = v4096
	var v4099 int32
	_ = v4099
	var v4102 int32
	_ = v4102
	var v4104 int32
	_ = v4104
	var v4106 int32
	_ = v4106
	var v4108 int32
	_ = v4108
	var v4109 int32
	_ = v4109
	var v4110 int32
	_ = v4110
	var v4116 int32
	_ = v4116
	var v4117 int32
	_ = v4117
	var v4118 int32
	_ = v4118
	var v4120 int32
	_ = v4120
	var v4125 int32
	_ = v4125
	var v4130 int32
	_ = v4130
	var v4131 int32
	_ = v4131
	var v4133 int32
	_ = v4133
	var v4137 int32
	_ = v4137
	var v4139 int32
	_ = v4139
	var v4140 int32
	_ = v4140
	var v4143 int32
	_ = v4143
	var v4146 int32
	_ = v4146
	var v4148 int32
	_ = v4148
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4156 int32
	_ = v4156
	var v4161 int32
	_ = v4161
	var v4164 int32
	_ = v4164
	var v4168 int32
	_ = v4168
	var v4172 int32
	_ = v4172
	var v4175 int32
	_ = v4175
	var v4176 int32
	_ = v4176
	var v4179 int32
	_ = v4179
	var v4182 int32
	_ = v4182
	var v4189 int32
	_ = v4189
	var v4190 int32
	_ = v4190
	var v4194 int32
	_ = v4194
	var v4195 int32
	_ = v4195
	var v4201 int32
	_ = v4201
	var v4206 int32
	_ = v4206
	var v4209 int32
	_ = v4209
	var v4213 int32
	_ = v4213
	var v4216 int32
	_ = v4216
	var v4219 int32
	_ = v4219
	var v4220 int32
	_ = v4220
	var v4227 int32
	_ = v4227
	var v4235 int32
	_ = v4235
	var v4236 int32
	_ = v4236
	var v4239 int32
	_ = v4239
	var v4240 int32
	_ = v4240
	var v4248 int32
	_ = v4248
	var v4253 int32
	_ = v4253
	var v4257 int32
	_ = v4257
	var v4260 int32
	_ = v4260
	var v4262 int32
	_ = v4262
	var v4264 int32
	_ = v4264
	var v4266 int32
	_ = v4266
	var v4267 int32
	_ = v4267
	var v4268 int32
	_ = v4268
	var v4274 int32
	_ = v4274
	var v4275 int32
	_ = v4275
	var v4276 int32
	_ = v4276
	var v4278 int32
	_ = v4278
	var v4283 int32
	_ = v4283
	var v4288 int32
	_ = v4288
	var v4289 int32
	_ = v4289
	var v4291 int32
	_ = v4291
	var v4296 int32
	_ = v4296
	var v4297 int32
	_ = v4297
	var v4301 int32
	_ = v4301
	var v4306 int32
	_ = v4306
	var v4311 int32
	_ = v4311
	var v4313 int32
	_ = v4313
	var v4316 int32
	_ = v4316
	var v4319 int32
	_ = v4319
	var v4321 int32
	_ = v4321
	var v4325 int32
	_ = v4325
	var v4326 int32
	_ = v4326
	var v4328 int32
	_ = v4328
	var v4332 int32
	_ = v4332
	var v4337 int32
	_ = v4337
	var v4339 int32
	_ = v4339
	var v4341 int32
	_ = v4341
	var v4342 int32
	_ = v4342
	var v4345 int32
	_ = v4345
	var v4346 int32
	_ = v4346
	var v4348 int32
	_ = v4348
	var v4350 int32
	_ = v4350
	var v4352 int32
	_ = v4352
	var v4354 int32
	_ = v4354
	var v4357 int32
	_ = v4357
	var v4358 int32
	_ = v4358
	var v4362 int32
	_ = v4362
	var v4367 int32
	_ = v4367
	var v4369 int32
	_ = v4369
	var v4370 int32
	_ = v4370
	var v4385 int32
	_ = v4385
	var v4388 int32
	_ = v4388
	var v4390 int32
	_ = v4390
	var v4392 int32
	_ = v4392
	var v4394 int32
	_ = v4394
	var v4395 int32
	_ = v4395
	var v4396 int32
	_ = v4396
	var v4402 int32
	_ = v4402
	var v4403 int32
	_ = v4403
	var v4404 int32
	_ = v4404
	var v4406 int32
	_ = v4406
	var v4407 int32
	_ = v4407
	var v4412 int32
	_ = v4412
	var v4413 int32
	_ = v4413
	var v4416 int32
	_ = v4416
	var v4422 int32
	_ = v4422
	var v4427 int32
	_ = v4427
	var v4429 int32
	_ = v4429
	var v4434 int32
	_ = v4434
	var v4439 int32
	_ = v4439
	var v4440 int32
	_ = v4440
	var v4442 int32
	_ = v4442
	var v4447 int32
	_ = v4447
	var v4452 int32
	_ = v4452
	var v4454 int32
	_ = v4454
	var v4455 int32
	_ = v4455
	var v4459 int32
	_ = v4459
	var v4460 int32
	_ = v4460
	var v4465 int32
	_ = v4465
	var v4469 int64
	_ = v4469
	var v4490 int32
	_ = v4490
	var v4491 int32
	_ = v4491
	var v4494 int32
	_ = v4494
	var v4495 int32
	_ = v4495
	var v4504 int32
	_ = v4504
	var v4507 int32
	_ = v4507
	var v4510 int32
	_ = v4510
	var v4512 int32
	_ = v4512
	var v4514 int32
	_ = v4514
	var v4516 int32
	_ = v4516
	var v4517 int32
	_ = v4517
	var v4518 int32
	_ = v4518
	var v4524 int32
	_ = v4524
	var v4525 int32
	_ = v4525
	var v4526 int32
	_ = v4526
	var v4528 int32
	_ = v4528
	var v4529 int32
	_ = v4529
	var v4534 int32
	_ = v4534
	var v4535 int32
	_ = v4535
	var v4538 int32
	_ = v4538
	var v4544 int32
	_ = v4544
	var v4549 int32
	_ = v4549
	var v4551 int32
	_ = v4551
	var v4556 int32
	_ = v4556
	var v4561 int32
	_ = v4561
	var v4562 int32
	_ = v4562
	var v4564 int32
	_ = v4564
	var v4569 int32
	_ = v4569
	var v4574 int32
	_ = v4574
	var v4576 int32
	_ = v4576
	var v4577 int32
	_ = v4577
	var v4581 int32
	_ = v4581
	var v4582 int32
	_ = v4582
	var v4587 int32
	_ = v4587
	var v4588 int64
	_ = v4588
	var v4603 int32
	_ = v4603
	var v4614 int32
	_ = v4614
	var v4615 int32
	_ = v4615
	var v4618 int32
	_ = v4618
	var v4619 int32
	_ = v4619
	var v4628 int32
	_ = v4628
	var v4631 int32
	_ = v4631
	var v4634 int32
	_ = v4634
	var v4636 int32
	_ = v4636
	var v4638 int32
	_ = v4638
	var v4640 int32
	_ = v4640
	var v4641 int32
	_ = v4641
	var v4642 int32
	_ = v4642
	var v4648 int32
	_ = v4648
	var v4649 int32
	_ = v4649
	var v4650 int32
	_ = v4650
	var v4653 int32
	_ = v4653
	var v4654 int32
	_ = v4654
	var v4655 int32
	_ = v4655
	var v4661 int32
	_ = v4661
	var v4666 int32
	_ = v4666
	var v4668 int32
	_ = v4668
	var v4673 int32
	_ = v4673
	var v4678 int32
	_ = v4678
	var v4679 int32
	_ = v4679
	var v4681 int32
	_ = v4681
	var v4684 int32
	_ = v4684
	var v4687 int32
	_ = v4687
	var v4690 int32
	_ = v4690
	var v4691 int32
	_ = v4691
	var v4693 int32
	_ = v4693
	var v4697 int32
	_ = v4697
	var v4698 int32
	_ = v4698
	var v4700 int32
	_ = v4700
	var v4701 int32
	_ = v4701
	var v4702 int32
	_ = v4702
	var v4703 int32
	_ = v4703
	var v4709 int32
	_ = v4709
	var v4713 int32
	_ = v4713
	var v4718 int32
	_ = v4718
	var v4721 int32
	_ = v4721
	var v4723 int32
	_ = v4723
	var v4724 int32
	_ = v4724
	var v4730 int32
	_ = v4730
	var v4736 int32
	_ = v4736
	var v4741 int32
	_ = v4741
	var v4744 int32
	_ = v4744
	var v4750 int32
	_ = v4750
	var v4753 int32
	_ = v4753
	var v4755 int32
	_ = v4755
	var v4757 int32
	_ = v4757
	var v4759 int32
	_ = v4759
	var v4760 int32
	_ = v4760
	var v4761 int32
	_ = v4761
	var v4767 int32
	_ = v4767
	var v4768 int32
	_ = v4768
	var v4769 int32
	_ = v4769
	var v4771 int32
	_ = v4771
	var v4776 int32
	_ = v4776
	var v4781 int32
	_ = v4781
	var v4782 int32
	_ = v4782
	var v4784 int32
	_ = v4784
	var v4788 int32
	_ = v4788
	var v4793 int32
	_ = v4793
	var v4815 int32
	_ = v4815
	var v4817 int32
	_ = v4817
	var v4818 int32
	_ = v4818
	var v4820 int32
	_ = v4820
	var v4821 int32
	_ = v4821
	var v4823 int32
	_ = v4823
	var v4824 int32
	_ = v4824
	var v4826 int32
	_ = v4826
	var v4827 int32
	_ = v4827
	var v4828 int32
	_ = v4828
	var v4832 int32
	_ = v4832
	var v4835 int32
	_ = v4835
	var v4838 int32
	_ = v4838
	var v4839 int32
	_ = v4839
	var v4841 int32
	_ = v4841
	var v4842 int32
	_ = v4842
	var v4872 int32
	_ = v4872
	var v4875 int32
	_ = v4875
	var v4877 int32
	_ = v4877
	var v4879 int32
	_ = v4879
	var v4881 int32
	_ = v4881
	var v4882 int32
	_ = v4882
	var v4883 int32
	_ = v4883
	var v4889 int32
	_ = v4889
	var v4890 int32
	_ = v4890
	var v4891 int32
	_ = v4891
	var v4894 int32
	_ = v4894
	var v4895 int32
	_ = v4895
	var v4896 int32
	_ = v4896
	var v4897 int32
	_ = v4897
	var v4898 int32
	_ = v4898
	var v4904 int32
	_ = v4904
	var v4905 int32
	_ = v4905
	var v4907 int32
	_ = v4907
	var v4908 int32
	_ = v4908
	var v4913 int32
	_ = v4913
	var v4919 int32
	_ = v4919
	var v4920 int32
	_ = v4920
	var v4921 int32
	_ = v4921
	var v4927 int32
	_ = v4927
	var v4928 int32
	_ = v4928
	var v4936 int32
	_ = v4936
	var v4937 int32
	_ = v4937
	var v4939 int32
	_ = v4939
	var v4941 int32
	_ = v4941
	var v4947 int32
	_ = v4947
	var v4950 int32
	_ = v4950
	var v4951 int32
	_ = v4951
	var v4952 int32
	_ = v4952
	var v4953 int32
	_ = v4953
	var v4955 int32
	_ = v4955
	var v4958 int32
	_ = v4958
	var v4961 int32
	_ = v4961
	var v4962 int32
	_ = v4962
	var v4966 int32
	_ = v4966
	var v4971 int32
	_ = v4971
	var v4974 int32
	_ = v4974
	var v4976 int32
	_ = v4976
	var v4979 int32
	_ = v4979
	var v4983 int32
	_ = v4983
	var v4985 int32
	_ = v4985
	var v4986 int32
	_ = v4986
	var v4988 int32
	_ = v4988
	var v4989 int32
	_ = v4989
	var v4992 int32
	_ = v4992
	var v4996 int32
	_ = v4996
	var v5001 int32
	_ = v5001
	var v5002 int32
	_ = v5002
	var v5003 int32
	_ = v5003
	var v5012 int32
	_ = v5012
	var v5017 int32
	_ = v5017
	var v5019 int32
	_ = v5019
	var v5021 int32
	_ = v5021
	var v5024 int32
	_ = v5024
	var v5036 int32
	_ = v5036
	var v5056 int32
	_ = v5056
	var v5063 int32
	_ = v5063
	var v5064 int32
	_ = v5064
	var v5065 int32
	_ = v5065
	var v5066 int32
	_ = v5066
	var v5067 int32
	_ = v5067
	var v5069 int32
	_ = v5069
	var v5075 int32
	_ = v5075
	var v5078 int32
	_ = v5078
	var v5079 int32
	_ = v5079
	var v5080 int32
	_ = v5080
	var v5085 int32
	_ = v5085
	var v5087 int32
	_ = v5087
	var v5090 int32
	_ = v5090
	var v5094 int32
	_ = v5094
	var v5095 int32
	_ = v5095
	var v5106 int32
	_ = v5106
	var v5111 int32
	_ = v5111
	var v5112 int32
	_ = v5112
	var v5118 int32
	_ = v5118
	var v5123 int32
	_ = v5123
	var v5125 int32
	_ = v5125
	var v5126 int32
	_ = v5126
	var v5128 int32
	_ = v5128
	var v5139 int32
	_ = v5139
	var v5140 int32
	_ = v5140
	var v5157 int32
	_ = v5157
	var v5160 int32
	_ = v5160
	var v5170 int32
	_ = v5170
	var v5190 int32
	_ = v5190
	var v5192 int32
	_ = v5192
	var v5199 int32
	_ = v5199
	var v5200 int32
	_ = v5200
	var v5201 int32
	_ = v5201
	var v5202 int32
	_ = v5202
	var v5203 int32
	_ = v5203
	var v5205 int32
	_ = v5205
	var v5211 int32
	_ = v5211
	var v5214 int32
	_ = v5214
	var v5215 int32
	_ = v5215
	var v5216 int32
	_ = v5216
	var v5221 int32
	_ = v5221
	var v5223 int32
	_ = v5223
	var v5226 int32
	_ = v5226
	var v5230 int32
	_ = v5230
	var v5231 int32
	_ = v5231
	var v5242 int32
	_ = v5242
	var v5271 int32
	_ = v5271
	var v5276 int32
	_ = v5276
	var v5278 int32
	_ = v5278
	var v5281 int32
	_ = v5281
	var v5284 int32
	_ = v5284
	var v5295 int32
	_ = v5295
	var v5315 int32
	_ = v5315
	var v5317 int32
	_ = v5317
	var v5324 int32
	_ = v5324
	var v5325 int32
	_ = v5325
	var v5326 int32
	_ = v5326
	var v5327 int32
	_ = v5327
	var v5328 int32
	_ = v5328
	var v5330 int32
	_ = v5330
	var v5336 int32
	_ = v5336
	var v5339 int32
	_ = v5339
	var v5340 int32
	_ = v5340
	var v5341 int32
	_ = v5341
	var v5346 int32
	_ = v5346
	var v5348 int32
	_ = v5348
	var v5351 int32
	_ = v5351
	var v5355 int32
	_ = v5355
	var v5356 int32
	_ = v5356
	var v5367 int32
	_ = v5367
	var v5398 int32
	_ = v5398
	var v5402 int32
	_ = v5402
	var v5407 int32
	_ = v5407
	var v5408 int32
	_ = v5408
	var v5410 int32
	_ = v5410
	var v5411 int32
	_ = v5411
	var v5413 int32
	_ = v5413
	var v5414 int32
	_ = v5414
	var v5416 int32
	_ = v5416
	var v5417 int32
	_ = v5417
	var v5419 int32
	_ = v5419
	var v5420 int32
	_ = v5420
	var v5422 int32
	_ = v5422
	var v5423 int32
	_ = v5423
	var v5425 int32
	_ = v5425
	var v5426 int32
	_ = v5426
	var v5429 int32
	_ = v5429
	var v5430 int32
	_ = v5430
	var v5432 int32
	_ = v5432
	var v5437 int32
	_ = v5437
	var v5444 int32
	_ = v5444
	var v5466 int32
	_ = v5466
	var v5469 int32
	_ = v5469
	var v5470 int32
	_ = v5470
	var v5472 int32
	_ = v5472
	var v5473 int32
	_ = v5473
	var v5475 int32
	_ = v5475
	var v5476 int32
	_ = v5476
	var v5478 int32
	_ = v5478
	var v5479 int32
	_ = v5479
	var v5481 int32
	_ = v5481
	var v5482 int32
	_ = v5482
	var v5484 int32
	_ = v5484
	var v5487 int32
	_ = v5487
	var v5488 int32
	_ = v5488
	var v5490 int32
	_ = v5490
	var v5493 int32
	_ = v5493
	var v5496 int32
	_ = v5496
	var v5497 int32
	_ = v5497
	var v5502 int32
	_ = v5502
	var v5529 int32
	_ = v5529
	var v5531 int32
	_ = v5531
	var v5552 int32
	_ = v5552
	var v5553 int32
	_ = v5553
	var v5556 int32
	_ = v5556
	var v5557 int32
	_ = v5557
	var v5560 int32
	_ = v5560
	var v5561 int32
	_ = v5561
	var v5564 int32
	_ = v5564
	var v5568 int32
	_ = v5568
	var v5571 int32
	_ = v5571
	var v5583 int32
	_ = v5583
	var v5602 int32
	_ = v5602
	var v5603 int32
	_ = v5603
	var v5606 int32
	_ = v5606
	var v5610 int32
	_ = v5610
	var v5621 int32
	_ = v5621
	var v5645 int32
	_ = v5645
	var v5697 int32
	_ = v5697
	var v5701 int64
	_ = v5701
	var v5703 int32
	_ = v5703
	var v5704 int32
	_ = v5704
	var v5706 int32
	_ = v5706
	var v5710 int32
	_ = v5710
	var v5712 int32
	_ = v5712
	var v5716 int32
	_ = v5716
	var v5717 int32
	_ = v5717
	var v5724 int32
	_ = v5724
	var v5729 int32
	_ = v5729
	var v5731 int32
	_ = v5731
	var v5732 int32
	_ = v5732
	var v5733 int32
	_ = v5733
	var v5739 int32
	_ = v5739
	var v5740 int32
	_ = v5740
	var v5758 int32
	_ = v5758
	var v5764 int32
	_ = v5764
	var v5765 int32
	_ = v5765
	var v5769 int32
	_ = v5769
	var v5773 int32
	_ = v5773
	var v5776 int32
	_ = v5776
	var v5778 int32
	_ = v5778
	var v5780 int32
	_ = v5780
	var v5785 int64
	_ = v5785
	var v5786 int32
	_ = v5786
	var v5788 int32
	_ = v5788
	var v5789 int32
	_ = v5789
	var v5798 int32
	_ = v5798
	var v5799 int32
	_ = v5799
	var v5807 int32
	_ = v5807
	var v5812 int32
	_ = v5812
	var v5814 int32
	_ = v5814
	var v5833 int32
	_ = v5833
	var v5839 int32
	_ = v5839
	var v5857 int32
	_ = v5857
	var v5859 int32
	_ = v5859
	var v5861 int32
	_ = v5861
	var v5864 int32
	_ = v5864
	var v5865 int32
	_ = v5865
	var v5869 int32
	_ = v5869
	var v5870 int32
	_ = v5870
	var v5872 int64
	_ = v5872
	var v5873 int32
	_ = v5873
	var v5874 int64
	_ = v5874
	var v5876 int32
	_ = v5876
	var v5877 int32
	_ = v5877
	var v5878 int32
	_ = v5878
	var v5879 int32
	_ = v5879
	var v5880 int32
	_ = v5880
	var v5881 int32
	_ = v5881
	var v5882 int32
	_ = v5882
	var v5884 int32
	_ = v5884
	var v5885 int32
	_ = v5885
	var v5893 int32
	_ = v5893
	var v5914 int32
	_ = v5914
	var v5915 int32
	_ = v5915
	var v5916 int32
	_ = v5916
	var v5917 int32
	_ = v5917
	var v5921 int64
	_ = v5921
	var v5924 int32
	_ = v5924
	var v5932 int32
	_ = v5932
	var v5934 int32
	_ = v5934
	var v5935 int32
	_ = v5935
	var v5946 int32
	_ = v5946
	var v5951 int32
	_ = v5951
	var v5956 int32
	_ = v5956
	var v5957 int32
	_ = v5957
	var v5959 int64
	_ = v5959
	var v5960 int32
	_ = v5960
	var v5990 int32
	_ = v5990
	var v5991 int32
	_ = v5991
	var v5992 int32
	_ = v5992
	var v5997 int64
	_ = v5997
	var v5998 int32
	_ = v5998
	var v5999 int32
	_ = v5999
	var v6003 int32
	_ = v6003
	var v6008 int32
	_ = v6008
	var v6015 int32
	_ = v6015
	var v6019 int32
	_ = v6019
	var v6024 int32
	_ = v6024
	var v6028 int32
	_ = v6028
	var v6032 int32
	_ = v6032
	var v6037 int32
	_ = v6037
	var v6041 int32
	_ = v6041
	var v6045 int32
	_ = v6045
	var v6050 int32
	_ = v6050
	var v6051 int32
	_ = v6051
	var v6053 int32
	_ = v6053
	var v6057 int32
	_ = v6057
	var v6059 int32
	_ = v6059
	var v6063 int32
	_ = v6063
	var v6064 int32
	_ = v6064
	var v6070 int32
	_ = v6070
	var v6075 int32
	_ = v6075
	var v6077 int32
	_ = v6077
	var v6079 int32
	_ = v6079
	var v6080 int32
	_ = v6080
	var v6081 int32
	_ = v6081
	var v6086 int32
	_ = v6086
	var v6092 int32
	_ = v6092
	var v6094 int32
	_ = v6094
	var v6095 int32
	_ = v6095
	var v6096 int32
	_ = v6096
	var v6097 int32
	_ = v6097
	var v6110 int32
	_ = v6110
	var v6115 int32
	_ = v6115
	var v6124 int32
	_ = v6124
	var v6129 int32
	_ = v6129
	var v6130 int32
	_ = v6130
	var v6131 int32
	_ = v6131
	var v6132 int32
	_ = v6132
	var v6133 int32
	_ = v6133
	var v6134 int32
	_ = v6134
	var v6135 int32
	_ = v6135
	var v6136 int32
	_ = v6136
	var v6137 int32
	_ = v6137
	var v6138 int32
	_ = v6138
	var v6139 int32
	_ = v6139
	var v6140 int32
	_ = v6140
	var v6141 int32
	_ = v6141
	var v6142 int32
	_ = v6142
	var v6143 int32
	_ = v6143
	var v6144 int32
	_ = v6144
	var v6145 int32
	_ = v6145
	var v6146 int32
	_ = v6146
	var v6147 int32
	_ = v6147
	var v6148 int32
	_ = v6148
	var v6149 int32
	_ = v6149
	var v6150 int32
	_ = v6150
	var v6151 int32
	_ = v6151
	var v6152 int32
	_ = v6152
	var v6153 int32
	_ = v6153
	var v6154 int32
	_ = v6154
	var v6155 int32
	_ = v6155
	var v6156 int32
	_ = v6156
	var v6157 int32
	_ = v6157
	var v6158 int32
	_ = v6158
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
	var v6164 int32
	_ = v6164
	var v6165 int32
	_ = v6165
	var v6166 int32
	_ = v6166
	var v6167 int32
	_ = v6167
	var v6168 int32
	_ = v6168
	var v6169 int32
	_ = v6169
	var v6170 int32
	_ = v6170
	var v6171 int32
	_ = v6171
	var v6172 int32
	_ = v6172
	var v6173 int32
	_ = v6173
	var v6174 int32
	_ = v6174
	var v6175 int32
	_ = v6175
	var v6176 int32
	_ = v6176
	var v6177 int32
	_ = v6177
	var v6178 int32
	_ = v6178
	var v6179 int32
	_ = v6179
	var v6180 int32
	_ = v6180
	var v6181 int32
	_ = v6181
	var v6182 int32
	_ = v6182
	var v6183 int32
	_ = v6183
	var v6184 int32
	_ = v6184
	var v6185 int32
	_ = v6185
	var v6186 int32
	_ = v6186
	var v6189 int32
	_ = v6189
	var v6215 int32
	_ = v6215
	var v6218 int32
	_ = v6218
	var v6219 int32
	_ = v6219
	var v6220 int32
	_ = v6220
	var v6223 int32
	_ = v6223
	var v6226 int32
	_ = v6226
	var v6227 int32
	_ = v6227
	var v6230 int32
	_ = v6230
	var v6236 int32
	_ = v6236
	var v6239 int32
	_ = v6239
	var v6243 int32
	_ = v6243
	var v6247 int32
	_ = v6247
	var v6250 int32
	_ = v6250
	var v6251 int32
	_ = v6251
	var v6252 int32
	_ = v6252
	var v6253 int32
	_ = v6253
	var v6254 int32
	_ = v6254
	var v6256 int32
	_ = v6256
	var v6257 int32
	_ = v6257
	var v6259 int32
	_ = v6259
	var v6262 int32
	_ = v6262
	var v6270 int32
	_ = v6270
	var v6283 int32
	_ = v6283
	var v6285 int32
	_ = v6285
	var v6287 int32
	_ = v6287
	var v6313 int32
	_ = v6313
	var v6332 int32
	_ = v6332
	var v6336 int32
	_ = v6336
	var v6338 int32
	_ = v6338
	var v6341 int32
	_ = v6341
	var v6347 int32
	_ = v6347
	var v6352 int32
	_ = v6352
	var v6356 int32
	_ = v6356
	var v6360 int32
	_ = v6360
	var v6365 int32
	_ = v6365
	var v6369 int32
	_ = v6369
	var v6373 int32
	_ = v6373
	var v6378 int32
	_ = v6378
	var v6389 int32
	_ = v6389
	var v6391 int32
	_ = v6391
	var v6405 int32
	_ = v6405
	var v6415 int32
	_ = v6415
	var v6435 int32
	_ = v6435
	var v6440 int32
	_ = v6440
	var v6442 int32
	_ = v6442
	var v6446 int32
	_ = v6446
	var v6447 int32
	_ = v6447
	var v6449 int32
	_ = v6449
	var v6456 int32
	_ = v6456
	var v6458 int32
	_ = v6458
	var v6462 int32
	_ = v6462
	var v6464 int32
	_ = v6464
	var v6466 int32
	_ = v6466
	var v6468 int32
	_ = v6468
	var v6470 int32
	_ = v6470
	var v6474 int32
	_ = v6474
	var v6476 int32
	_ = v6476
	var v6479 int32
	_ = v6479
	var v6482 int32
	_ = v6482
	var v6486 int32
	_ = v6486
	var v6489 int32
	_ = v6489
	var v6490 int32
	_ = v6490
	var v6496 int32
	_ = v6496
	var v6501 int32
	_ = v6501
	var v6507 int32
	_ = v6507
	var v6512 int32
	_ = v6512
	var v6516 int32
	_ = v6516
	var v6519 int32
	_ = v6519
	var v6525 int32
	_ = v6525
	var v6528 int32
	_ = v6528
	var v6531 int32
	_ = v6531
	var v6537 int32
	_ = v6537
	var v6541 int32
	_ = v6541
	var v6545 int32
	_ = v6545
	var v6550 int32
	_ = v6550
	v4 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(128)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_InitStandaloneProcess(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_InitializeGUCOptions(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v29 + int32(96)
	v39 = l0 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = int32(_a_F_BootstrapModeMain_0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = l1 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v37)+28)) = int32(_a_F_BootstrapModeMain_1)
	v48 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v37)+20)) = v48
	*(*int64)(unsafe.Add(mBase, uint32(v37)+12)) = v48
	goto L4
L4:
	;
	v54 = int32(0)
	v70 = v4
	goto L12
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6541 = m.ExcPending
	if v6541 != 0 {
		goto L1
	} else {
		goto L1246
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(2)
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6537 = m.ExcPending
	if v6537 != 0 {
		goto L1
	} else {
		goto L1245
	}
L7:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6531 = m.ExcPending
	if v6531 != 0 {
		goto L1
	} else {
		goto L1244
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_2), v29+int32(16))
	mBase = m.M
	v6525 = m.ExcPending
	if v6525 != 0 {
		goto L1
	} else {
		goto L1242
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_write_stderr(m, int32(_a_F_BootstrapModeMain_3), v29)
	mBase = m.M
	v6516 = m.ExcPending
	if v6516 != 0 {
		goto L1
	} else {
		goto L1240
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v139
	F_errmsg(m, int32(_a_F_BootstrapModeMain_4), v29+int32(32))
	mBase = m.M
	v6507 = m.ExcPending
	if v6507 != 0 {
		goto L1
	} else {
		goto L1238
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6486 = m.ExcPending
	if v6486 != 0 {
		goto L1
	} else {
		goto L1234
	}
L12:
	;
	v82 = F_pg_getopt_next(m, v29+int32(96))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L23
	}
L13:
	;
	if v82 != int32(-1) {
		goto L9
	} else {
		goto L92
	}
L14:
	;
	goto L13
L15:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_5), v314, int32(0), int32(1))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L91
	}
L16:
	;
	v192 = int32(_a_F_BootstrapModeMain_6)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	goto L63
L17:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_7), int32(_a_F_BootstrapModeMain_8), int32(1), int32(4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L59
	}
L18:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v167
	v173 = F_psprintf(m, int32(_a_F_BootstrapModeMain_9), v29+int32(80))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L55
	}
L19:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	v165 = F_pstrdup(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L20:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	F_ParseLongOption(m, v122, v29+int32(92), v29+int32(88))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L42
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	v94 = F_strcmp(m, int32(_a_F_BootstrapModeMain_10), v92)
	mBase = m.M
	if v94 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_11), v87, int32(1), int32(4))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	switch v82 - int32(45) {
	case 0:
		goto L21
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 22, 24, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 56, 57, 58, 59, 60, 61, 63, 64, 65, 66, 67, 68:
		goto L9
	case 21:
		goto L22
	case 23:
		goto L19
	case 25:
		goto L17
	case 43:
		goto L15
	case 54:
		goto L20
	case 55:
		goto L18
	case 62:
		v54 = int32(1)
		goto L12
	case 69:
		goto L16
	default:
		goto L14
	}
L24:
	;
	goto L12
L25:
	;
	if v119 != int32(5) {
		goto L11
	} else {
		goto L41
	}
L26:
	;
	v119 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v99 = F_strcmp(m, int32(_a_F_BootstrapModeMain_12), v92)
	mBase = m.M
	if v99 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = int32(1)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v105 = F_strncmp(m, int32(_a_F_BootstrapModeMain_13), v92, int32(9))
	mBase = m.M
	if v105 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v119 = int32(2)
	goto L25
L33:
	;
	goto L34
L34:
	;
	v110 = F_strcmp(m, int32(_a_F_BootstrapModeMain_14), v92)
	mBase = m.M
	if v110 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v119 = int32(3)
	goto L25
L36:
	;
	goto L37
L37:
	;
	v117 = F_strcmp(m, int32(_a_F_BootstrapModeMain_15), v92)
	mBase = m.M
	if v117 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v118 = int32(5)
	goto L40
L39:
	;
	v118 = int32(4)
	goto L40
L40:
	;
	v119 = v118
	goto L25
L41:
	;
	goto L20
L42:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	if v129 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_SetConfigOption(m, v153, v129, int32(1), int32(4))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L46:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v82 == int32(45) {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v139
	F_errmsg(m, int32(_a_F_BootstrapModeMain_16), v29+int32(48))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(300), int32(_a_F_BootstrapModeMain_18))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	F_pfree(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	F_pfree(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L12
L54:
	;
	v70 = v165
	goto L12
L55:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_19), v173, int32(1), int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_SetConfigOption(m, int32(_a_F_BootstrapModeMain_20), v173, int32(1), int32(4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_pfree(m, v173)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	goto L12
L60:
	;
	goto L12
L61:
	;
	v310 = F_strlen(m, v299)
	mBase = m.M
	goto L60
L63:
	;
	goto L64
L64:
	;
	v200 = int32(1023)
	if (v192^v193)&int32(3) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v303 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v303)
	goto L61
L66:
	;
	v284 = v279
	v285 = v280
	v286 = v281
	goto L87
L67:
	;
	if v274 == int32(0) {
		v299 = v272
		v300 = v273
		goto L65
	} else {
		goto L86
	}
L68:
	;
	v272 = v193
	v273 = v192
	v274 = v200
	goto L67
L69:
	;
	goto L70
L70:
	;
	v204 = int32(0)
	if base.B2i32(v193&int32(3) == v204)|int32(0) == v204 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v240 == int32(0) {
		v299 = v237
		v300 = v238
		goto L65
	} else {
		goto L80
	}
L72:
	;
	v216 = v193
	v217 = v192
	v218 = v200
	goto L75
L73:
	;
	goto L74
L74:
	;
	v237 = v193
	v238 = v192
	v239 = v200
	v240 = int32(1)
	goto L71
L75:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v217))) = uint8(v220)
	if v220 == int32(0) {
		v279 = v216
		v280 = v217
		v281 = v218
		goto L66
	} else {
		goto L77
	}
L76:
	;
	v237 = v231
	v238 = v225
	v239 = v227
	v240 = v229
	goto L71
L77:
	;
	v224 = int32(1)
	v225 = v217 + v224
	v227 = v218 - v224
	v228 = int32(0)
	v229 = base.B2i32(v227 != v228)
	v231 = v216 + v224
	if v231&int32(3) == v228 {
		v237 = v231
		v238 = v225
		v239 = v227
		v240 = v229
		goto L71
	} else {
		goto L78
	}
L78:
	;
	if v227 != 0 {
		v216 = v231
		v217 = v225
		v218 = v227
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
L80:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if base.B2i32(v243 == int32(0))|base.B2i32(base.Ui32(v239) < base.Ui32(int32(4))) != 0 {
		v272 = v237
		v273 = v238
		v274 = v239
		goto L67
	} else {
		goto L81
	}
L81:
	;
	v250 = v237
	v251 = v238
	v252 = v239
	goto L82
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)))
	v258 = int32(-2139062144)
	if (int32(16843008)-v255|v255)&v258 != v258 {
		v279 = v250
		v280 = v251
		v281 = v252
		goto L66
	} else {
		goto L84
	}
L83:
	;
	v272 = v266
	v273 = v264
	v274 = v268
	goto L67
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v255
	v263 = int32(4)
	v264 = v251 + v263
	v266 = v250 + v263
	v268 = v252 - v263
	if base.Ui32(int32(3)) < base.Ui32(v268) {
		v250 = v266
		v251 = v264
		v252 = v268
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v279 = v272
	v280 = v273
	v281 = v274
	goto L66
L87:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v288)
	if v288 == int32(0) {
		v299 = v284
		v300 = v285
		goto L65
	} else {
		goto L89
	}
L88:
	;
	v299 = v295
	v300 = v293
	goto L65
L89:
	;
	v292 = int32(1)
	v293 = v285 + v292
	v295 = v284 + v292
	v297 = v286 - v292
	if v297 != 0 {
		v284 = v295
		v285 = v293
		v286 = v297
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	goto L12
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	if v39 != v321 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	v323 = F_SelectConfigFiles(m, v70, v31)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	if v323 == int32(0) {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_checkDataDir(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_ChangeToDataDir(m)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_CreateDataDirLockFile(m, int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v335 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[1])) = uint8(v335)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[0])) = int32(0)
	F_RegisterBuiltinShmemCallbacks(m)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_InitializeMaxBackends(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_InitPostmasterChildSlots(m)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v350 = int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[2]))
	if v353&(v353-v350) != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	F_ShmemCallRequestCallbacks(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L112
	}
L103:
	;
	v360 = v350 << (uint(int32(32)-base.I32_clz(v353)) % 32)
	goto L105
L104:
	;
	v360 = v353
	goto L105
L105:
	;
	if base.Ui32(v360) <= base.Ui32(int32(31)) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v363 = int32(31)
	goto L108
L107:
	;
	v363 = v360
	goto L108
L108:
	;
	if base.Ui32(int32(_a_F_BootstrapModeMain_21)) <= base.Ui32(v360) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v368 = int32(1024)
	goto L111
L110:
	;
	v368 = int32(base.Ui32(v363) >> (uint(int32(4)) % 32))
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[3])) = v368
	goto L102
L112:
	;
	F_CreateSharedMemoryAndSemaphores(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_set_max_safe_fds(m)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	F_InitProcess(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_BaseInit(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v384 = m.G0
	v386 = v384 - int32(32)
	m.G0 = v386
	v388 = int32(2)
	switch v388 {
	case 0, 2:
		goto L119
	default:
		goto L120
	}
L118:
	;
	v422 = int32(2)
	v426 = m.G0
	v428 = v426 - int32(32)
	m.G0 = v428
	switch v422 {
	case 0, 2:
		goto L129
	default:
		goto L130
	}
L119:
	;
	F_sigemptyset(m, v386+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = int32(268435456)
	switch v388 {
	case 0:
		goto L124
	default:
		goto L122
	case 2:
		goto L123
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[4])) = int32(0)
	goto L119
L121:
	;
	goto L126
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = int32(_a_F_BootstrapModeMain_22)
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = int32(0)
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v386)+12)) = int32(-2)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v418 = F___sigaction(m, int32(1), v386+int32(12), int32(0))
	mBase = m.M
	m.G0 = v386 + int32(32)
	goto L118
L128:
	;
	v468 = m.G0
	v470 = v468 - int32(32)
	m.G0 = v470
	v472 = int32(2)
	switch v472 {
	case 0, 2:
		goto L139
	default:
		goto L140
	}
L129:
	;
	F_sigemptyset(m, v428+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v428)+24)) = int32(268435456)
	switch v422 {
	case 0:
		goto L134
	default:
		goto L132
	case 2:
		goto L133
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[5])) = int32(0)
	goto L129
L131:
	;
	goto L136
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v428)+12)) = int32(_a_F_BootstrapModeMain_22)
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428)+12)) = int32(0)
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428)+12)) = int32(-2)
	goto L131
L136:
	;
	goto L137
L137:
	;
	v460 = F___sigaction(m, v422, v428+int32(12), int32(0))
	mBase = m.M
	m.G0 = v428 + int32(32)
	goto L128
L138:
	;
	v510 = m.G0
	v512 = v510 - int32(32)
	m.G0 = v512
	v514 = int32(2)
	switch v514 {
	case 0, 2:
		goto L149
	default:
		goto L150
	}
L139:
	;
	F_sigemptyset(m, v470+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v470)+24)) = int32(268435456)
	switch v472 {
	case 0:
		goto L144
	default:
		goto L142
	case 2:
		goto L143
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[6])) = int32(0)
	goto L139
L141:
	;
	goto L146
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v470)+12)) = int32(_a_F_BootstrapModeMain_22)
	goto L141
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470)+12)) = int32(0)
	goto L141
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v470)+12)) = int32(-2)
	goto L141
L146:
	;
	goto L147
L147:
	;
	v502 = F___sigaction(m, int32(15), v470+int32(12), int32(0))
	mBase = m.M
	m.G0 = v470 + int32(32)
	goto L138
L148:
	;
	v548 = m.G0
	v552 = (v548 - int32(_a_F_BootstrapModeMain_23)) & int32(-4096)
	m.G0 = v552
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7]))
	v559 = F_LWLockAcquire(m, v555+int32(1152), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L158
	}
L149:
	;
	F_sigemptyset(m, v512+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v512)+24)) = int32(268435456)
	switch v514 {
	case 0:
		goto L154
	default:
		goto L152
	case 2:
		goto L153
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[8])) = int32(0)
	goto L149
L151:
	;
	goto L156
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v512)+12)) = int32(_a_F_BootstrapModeMain_22)
	goto L151
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512)+12)) = int32(0)
	goto L151
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v512)+12)) = int32(-2)
	goto L151
L156:
	;
	goto L157
L157:
	;
	v544 = F___sigaction(m, int32(3), v512+int32(12), int32(0))
	mBase = m.M
	m.G0 = v512 + int32(32)
	goto L148
L158:
	;
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[9]))
	v563 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v562)+312)) = uint8(v563)
	v566 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7]))
	F_LWLockRelease(m, v566+int32(1152))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_gettimeofday(m, v552+int32(4080))
	mBase = m.M
	v574 = int64(*(*int32)(unsafe.Add(mBase, uint32(v552)+4088)))
	v575 = *(*int64)(unsafe.Add(mBase, uint32(v552)+4080))
	v576 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	v578 = v552 + int32(_a_F_BootstrapModeMain_24)
	v579 = int32(0)
	base.MemoryFill(m, v578, v579, int32(_a_F_BootstrapModeMain_25))
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10]))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[11])))
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	v588 = F_time(m)
	mBase = m.M
	v589 = int32(_a_F_BootstrapModeMain_26)
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = int32(_a_F_BootstrapModeMain_27)
	*(*int64)(unsafe.Add(mBase, uint32(v590)+8)) = int64(3)
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+4)) = v579
	F_MultiXactSetNextMXact(m, int32(1), int64(1))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_AdvanceOldestClogXid(m, int32(3))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	F_SetTransactionIdLimit(m, int32(3), int32(1))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v610 = int32(1)
	F_SetMultiXactIdLimit(m, v610, v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v614 = int32(0)
	F_SetCommitTsLimit(m, v614, v614)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[14]))) = int64(4295151906)
	v620 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[15]))) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[16]))) = int32(_a_F_BootstrapModeMain_25)
	v632 = base.I64_extend_i32_u(v576&int32(4095)) | (v574<<(uint(int64(12))%64) | v575<<(uint(int64(32))%64))
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[17]))) = v632
	v634 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[18]))) = v634
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[19]))) = v620
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[20]))) = v588
	v639 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[21]))) = v639
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[22]))) = int64(4294967299)
	v643 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[23]))) = v643
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[24]))) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[25]))) = int64(3)
	v650 = base.B2i32(v587 == int32(2))
	*(*uint8)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[26]))) = uint8(v650)
	*(*int32)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[27]))) = v587
	*(*uint8)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[28]))) = uint8(v585)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[29]))) = v639
	v656 = int32(40)
	v658 = base.I64_extend_i32_u(v583 + v656)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[30]))) = v658
	v660 = int32(_a_F_BootstrapModeMain_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[31]))) = uint16(v660)
	*(*uint16)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[32]))) = uint16(v634)
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[33]))) = int64(122)
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[34]))) = v667
	*(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[35]))) = base.I64_extend_i32_s(v667)
	v671 = int32(-1)
	v675 = m.Env.Pgmem_crc32c(m, v671, v578|int32(64), int32(98))
	mBase = m.M
	v679 = m.Env.Pgmem_crc32c(m, v675, v578|v656, int32(20))
	mBase = m.M
	v681 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[36])) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[37]))) = v679 ^ v671
	v689 = F_XLogFileInit(m, v643, v681)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[38])) = v689
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v696))) = int32(167772231)
	v699 = int32(_a_F_BootstrapModeMain_25)
	v700 = F_write(m, v689, v578, v699)
	mBase = m.M
	if v700 != v699 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v704 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	goto L168
L168:
	;
	v725 = int32(_a_F_BootstrapModeMain_29)
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	v727 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = v727
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v730))) = int32(167772230)
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[38]))
	v737 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[41])))
	if v737 != int32(1) {
		v751 = v727
		goto L183
	} else {
		goto L184
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(51)
	goto L171
L170:
	;
	goto L171
L171:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_30), int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_32), int32(_a_F_BootstrapModeMain_33))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	v1117 = int32(0)
	F_InitPostgres(m, v1117, v1117, v1117, v1117, v1117, v1117)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L1
	} else {
		goto L257
	}
L177:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L253
	}
L178:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L1
	} else {
		goto L249
	}
L179:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L1
	} else {
		goto L245
	}
L180:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L1
	} else {
		goto L241
	}
L181:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L237
	}
L182:
	;
	if v751 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L183:
	;
	goto L182
L184:
	;
	goto L185
L185:
	;
	v742 = F_fsync(m, v734)
	mBase = m.M
	if v742 != int32(-1) {
		v751 = v742
		goto L183
	} else {
		goto L187
	}
L186:
	;
	v751 = int32(-1)
	goto L183
L187:
	;
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v746 == int32(27) {
		goto L185
	} else {
		goto L188
	}
L188:
	;
	goto L186
L189:
	;
	v755 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = int32(0)
	v759 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[38]))
	v760 = F_close(m, v759)
	mBase = m.M
	if v760 != 0 {
		goto L181
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L233
	}
L192:
	;
	v762 = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[38])) = v762
	v765 = v552 + int32(_a_F_BootstrapModeMain_34)
	v767 = int32(0)
	v771 = m.G0
	v773 = v771 - int32(16)
	m.G0 = v773
	*(*int32)(unsafe.Add(mBase, uint32(v773))) = v767
	v779 = F_open(m, int32(_a_F_BootstrapModeMain_35), v767, v773)
	mBase = m.M
	if v779 != v762 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v812 == int32(0) {
		goto L180
	} else {
		goto L206
	}
L194:
	;
	goto L198
L195:
	;
	v812 = v767
	goto L196
L196:
	;
	m.G0 = v773 + int32(16)
	goto L193
L197:
	;
	v807 = F_close(m, v779)
	mBase = m.M
	v812 = v805
	goto L196
L198:
	;
	v785 = v765
	v786 = int32(32)
	goto L199
L199:
	;
	v791 = F_read(m, v779, v785, v786)
	mBase = m.M
	if v791 <= int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v805 = int32(1)
	goto L197
L201:
	;
	v795 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v795 == int32(27) {
		goto L199
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v800 = v786 - v791
	if v800 != 0 {
		v785 = v785 + v791
		v786 = v800
		goto L199
	} else {
		goto L205
	}
L204:
	;
	v805 = int32(0)
	goto L197
L205:
	;
	goto L200
L206:
	;
	v819 = int32(_a_F_BootstrapModeMain_36)
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[42]))
	v821 = int32(8)
	v823 = int32(0)
	base.MemoryFill(m, v820+v821, v823, int32(304))
	*(*int64)(unsafe.Add(mBase, uint32(v820))) = v632
	v827 = *(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[43])))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+273)) = v827
	v829 = *(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[44])))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+281)) = v829
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[45])))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+289)) = v831
	v833 = *(*int64)(unsafe.Add(mBase, uint32(v552)+uint32(_c_F_BootstrapModeMain[46])))
	*(*int64)(unsafe.Add(mBase, uint32(v820)+297)) = v833
	*(*int64)(unsafe.Add(mBase, uint32(v820)+136)) = int64(1000)
	v837 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+16)) = v837
	v840 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[47]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+188)) = v840
	v843 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[48]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+192)) = v843
	v846 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[49]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+196)) = v846
	v849 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[50]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+200)) = v849
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+204)) = v852
	v855 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+180)) = v855
	v858 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[51])))
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+184)) = uint8(v858)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[52])))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+268)) = v54
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+208)) = uint8(v861)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+264)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v820)+128)) = v823
	*(*int64)(unsafe.Add(mBase, uint32(v820)+120)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+112)) = v588
	v870 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+104)) = v870
	*(*int64)(unsafe.Add(mBase, uint32(v820)+96)) = int64(4294967299)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+88)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+80)) = int64(4294977296)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+72)) = int64(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+64)) = uint8(v650)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+60)) = v587
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+56)) = uint8(v585)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+48)) = v870
	*(*int64)(unsafe.Add(mBase, uint32(v820)+40)) = v658
	*(*int64)(unsafe.Add(mBase, uint32(v820)+32)) = v658
	*(*int64)(unsafe.Add(mBase, uint32(v820)+24)) = v588
	*(*int64)(unsafe.Add(mBase, uint32(v820)+232)) = int64(35184372088864)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+224)) = int64(562949953429504)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+216)) = int64(4698053236609777664)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+212)) = v821
	*(*int64)(unsafe.Add(mBase, uint32(v820)+8)) = int64(870199737544869745)
	v899 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[10]))
	v900 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+308)) = v900
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+272)) = uint8(v837)
	*(*uint8)(unsafe.Add(mBase, uint32(v820)+260)) = uint8(v837)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+252)) = int64(8796093024204)
	*(*int64)(unsafe.Add(mBase, uint32(v820)+244)) = int64(137438953536)
	*(*int32)(unsafe.Add(mBase, uint32(v820)+240)) = v899
	v913 = m.Env.Pgmem_crc32c(m, v900, v820, int32(308))
	mBase = m.M
	v915 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[42]))
	*(*int32)(unsafe.Add(mBase, uint32(v915)+308)) = v913 ^ v900
	base.MemoryFill(m, v552+int32(_a_F_BootstrapModeMain_37), v823, int32(_a_F_BootstrapModeMain_38))
	base.MemoryCopy(m, v765, v915, int32(312))
	v928 = F_BasicOpenFile(m, int32(_a_F_BootstrapModeMain_39), int32(194))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	if v928 < int32(0) {
		goto L179
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(0)
	v936 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v936))) = int32(167772172)
	v939 = int32(_a_F_BootstrapModeMain_25)
	v940 = F_write(m, v928, v765, v939)
	mBase = m.M
	if v940 != v939 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v944 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	v968 = int32(_a_F_BootstrapModeMain_29)
	v969 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	v970 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v969))) = v970
	v973 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v973))) = int32(167772170)
	v978 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[41])))
	if v978 != int32(1) {
		v992 = v970
		goto L220
	} else {
		goto L221
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(51)
	goto L214
L213:
	;
	goto L214
L214:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+4064)) = int32(_a_F_BootstrapModeMain_39)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_40), v552+int32(4064))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_41), int32(_a_F_BootstrapModeMain_42))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	if v992 != 0 {
		goto L178
	} else {
		goto L226
	}
L220:
	;
	goto L219
L221:
	;
	goto L222
L222:
	;
	v983 = F_fsync(m, v928)
	mBase = m.M
	if v983 != int32(-1) {
		v992 = v983
		goto L220
	} else {
		goto L224
	}
L223:
	;
	v992 = int32(-1)
	goto L220
L224:
	;
	v987 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v987 == int32(27) {
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v994 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v994))) = int32(0)
	v997 = F_close(m, v928)
	mBase = m.M
	if v997 != 0 {
		goto L177
	} else {
		goto L227
	}
L227:
	;
	F_SimpleLruZeroAndWritePage(m, int32(_a_F_BootstrapModeMain_43), int64(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_SimpleLruZeroAndWritePage(m, int32(_a_F_BootstrapModeMain_44), int64(0))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_SimpleLruZeroAndWritePage(m, int32(_a_F_BootstrapModeMain_45), int64(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	F_SimpleLruZeroAndWritePage(m, int32(_a_F_BootstrapModeMain_46), int64(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	F_ReadControlFile(m)
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	m.G0 = v548
	goto L176
L233:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_47), int32(0))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_48), int32(_a_F_BootstrapModeMain_33))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_49), int32(0))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_50), int32(_a_F_BootstrapModeMain_33))
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errmsg(m, int32(_a_F_BootstrapModeMain_51), int32(0))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_52), int32(_a_F_BootstrapModeMain_53))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1068 = m.ExcPending
	if v1068 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+4016)) = int32(_a_F_BootstrapModeMain_39)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_54), v552+int32(4016))
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_55), int32(_a_F_BootstrapModeMain_42))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L1
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+4048)) = int32(_a_F_BootstrapModeMain_39)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_56), v552+int32(4048))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L251
	}
L251:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_57), int32(_a_F_BootstrapModeMain_42))
	mBase = m.M
	v1098 = m.ExcPending
	if v1098 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L253:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552)+4032)) = int32(_a_F_BootstrapModeMain_39)
	F_errmsg(m, int32(_a_F_BootstrapModeMain_58), v552+int32(4032))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_31), int32(_a_F_BootstrapModeMain_59), int32(_a_F_BootstrapModeMain_42))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_60), int32(0), int32(160))
	v1130 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[53])) = v1130
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[54])) = v1130
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])) = v1130
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[56])) = v1130
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[57])) = v1130
	v1146 = F_boot_yylex_init(m, v29+int32(92))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	if v1146 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v29)+92))
	v1151 = m.G0
	v1153 = v1151 - int32(1136)
	m.G0 = v1153
	*(*int32)(unsafe.Add(mBase, uint32(v1153)+1132)) = int32(0)
	v1158 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58]))
	v1161 = v1153 + int32(128)
	v1163 = v1153 + int32(928)
	v1168 = v1150
	v1172 = v1161
	v1175 = v1153
	v1176 = int32(-2)
	v1177 = v1163
	v1178 = v4
	v1179 = v1163
	v1181 = v1158
	v1182 = v1161
	v1184 = int32(200)
	v1187 = v4
	goto L268
L261:
	;
	if v6415+int32(928) != v6405 {
		goto L1220
	} else {
		goto L1221
	}
L262:
	;
	v6405 = v6391
	v6415 = v6389
	goto L261
L263:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6369 = m.ExcPending
	if v6369 != 0 {
		goto L1
	} else {
		goto L1217
	}
L264:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v6356 = m.ExcPending
	if v6356 != 0 {
		goto L1
	} else {
		goto L1214
	}
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6336 = m.ExcPending
	if v6336 != 0 {
		goto L1
	} else {
		goto L1211
	}
L266:
	;
	F_boot_yyerror(m, v1168, int32(_a_F_BootstrapModeMain_61))
	mBase = m.M
	v6332 = m.ExcPending
	if v6332 != 0 {
		goto L1
	} else {
		goto L1210
	}
L267:
	;
	v6313 = v6287
	goto L1207
L268:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1179))) = uint8(v1178)
	if base.Ui32(v1177+v1184-int32(1)) <= base.Ui32(v1179) {
		goto L270
	} else {
		goto L271
	}
L269:
	;
	switch v3749 {
	case 0:
		goto L1204
	default:
		v6283 = v3737
		v6285 = v3739
		v6287 = v3741
		goto L267
	case 3:
		goto L1203
	}
L270:
	;
	if int32(_a_F_BootstrapModeMain_62) < v1184 {
		goto L266
	} else {
		goto L273
	}
L271:
	;
	v1242 = v1172
	v1243 = v1177
	v1244 = v1179
	v1245 = v1182
	v1246 = v1184
	goto L272
L272:
	;
	if v1178 == int32(46) {
		goto L290
	} else {
		goto L291
	}
L273:
	;
	v1198 = int32(_a_F_BootstrapModeMain_27)
	v1200 = v1184 << (uint(int32(1)) % 32)
	if v1198 <= v1200 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	v1203 = v1198
	goto L276
L275:
	;
	v1203 = v1200
	goto L276
L276:
	;
	v1208 = F_palloc(m, v1203*int32(5)+int32(3))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	if v1208 == int32(0) {
		goto L266
	} else {
		goto L278
	}
L278:
	;
	v1212 = v1179 - v1177
	v1214 = v1212 + int32(1)
	if v1214 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	base.MemoryCopy(m, v1208, v1177, v1214)
	goto L281
L280:
	;
	goto L281
L281:
	;
	v1219 = base.I32_div_s(v1203+int32(3), int32(4))
	v1220 = int32(2)
	v1222 = v1208 + v1219<<(uint(v1220)%32)
	v1224 = v1214 << (uint(v1220) % 32)
	if v1224 != 0 {
		goto L282
	} else {
		goto L283
	}
L282:
	;
	base.MemoryCopy(m, v1222, v1182, v1224)
	goto L284
L283:
	;
	goto L284
L284:
	;
	if v1175+int32(928) != v1177 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	F_pfree(m, v1177)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L1
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	if v1203-int32(1) <= v1212 {
		v6405 = v1208
		v6415 = v1175
		goto L261
	} else {
		goto L289
	}
L288:
	;
	goto L287
L289:
	;
	v1242 = v1222 + v1224 - int32(4)
	v1243 = v1208
	v1244 = v1208 + v1212
	v1245 = v1222
	v1246 = v1203
	goto L272
L290:
	;
	v6405 = v1243
	v6415 = v1175
	goto L261
L291:
	;
	goto L292
L292:
	;
	v1253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1178<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[59]))))
	if v1253 == int32(-53) {
		v3730 = v1168
		v3734 = v1242
		v3737 = v1175
		v3738 = v1176
		v3739 = v1243
		v3740 = v1178
		v3741 = v1244
		v3743 = v1181
		v3744 = v1245
		v3746 = v1246
		v3749 = v1187
		goto L296
	} else {
		goto L297
	}
L293:
	;
	goto L269
L294:
	;
	v1168 = v6243
	v1172 = v6247
	v1175 = v6250
	v1176 = v6251
	v1177 = v6252
	v1178 = v6253
	v1179 = v6254 + int32(1)
	v1181 = v6256
	v1182 = v6257
	v1184 = v6259
	v1187 = v6262
	goto L268
L295:
	;
	v3787 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3771)+uint32(_c_F_BootstrapModeMain[60]))))
	v3792 = *(*int32)(unsafe.Add(mBase, uint32(v3765+(int32(1)-v3787)<<(uint(int32(2))%32))))
	switch v3771 - int32(14) {
	case 0:
		goto L718
	case 1:
		goto L717
	case 2:
		goto L716
	case 3:
		goto L715
	case 4:
		goto L714
	case 5:
		goto L713
	case 6:
		goto L712
	case 7:
		goto L711
	case 8:
		goto L710
	case 9:
		goto L709
	case 10:
		goto L708
	case 11:
		goto L707
	case 12:
		goto L706
	case 13:
		goto L705
	case 14, 16, 25:
		goto L704
	case 15, 17, 19:
		goto L703
	case 18:
		goto L702
	default:
		v6189 = v3792
		goto L675
	case 22:
		goto L701
	case 23:
		goto L700
	case 24:
		goto L699
	case 26:
		goto L698
	case 30:
		goto L697
	case 31:
		goto L696
	case 32:
		goto L695
	case 33:
		goto L694
	case 34:
		goto L693
	case 35:
		goto L692
	case 36:
		goto L691
	case 37:
		goto L690
	case 38:
		goto L689
	case 39:
		goto L688
	case 40:
		goto L687
	case 41:
		goto L686
	case 42:
		goto L685
	case 43:
		goto L684
	case 44:
		goto L683
	case 45:
		goto L682
	case 46:
		goto L681
	case 47:
		goto L680
	case 48:
		goto L679
	case 49:
		goto L678
	case 50:
		goto L677
	case 51:
		goto L676
	}
L296:
	;
	v3755 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3740)+uint32(_c_F_BootstrapModeMain[61]))))
	if v3755 == int32(0) {
		goto L293
	} else {
		goto L674
	}
L297:
	;
	if v1176 == int32(-2) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v3709 = v1253 + v3708
	if base.Ui32(int32(169)) < base.Ui32(v3709) {
		v3730 = v3674
		v3734 = v3678
		v3737 = v3681
		v3738 = v3707
		v3739 = v3683
		v3740 = v3684
		v3741 = v3685
		v3743 = v3687
		v3744 = v3688
		v3746 = v3690
		v3749 = v3693
		goto L296
	} else {
		goto L669
	}
L299:
	;
	v1258 = m.G0
	v1260 = v1258 - int32(16)
	m.G0 = v1260
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+92)) = v1175 + int32(1132)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+40))
	if v1265 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v3674 = v1168
	v3678 = v1242
	v3681 = v1175
	v3682 = v1176
	v3683 = v1243
	v3684 = v1178
	v3685 = v1244
	v3687 = v1181
	v3688 = v1245
	v3690 = v1246
	v3693 = v1187
	goto L301
L301:
	;
	if v3682 <= int32(0) {
		goto L664
	} else {
		goto L665
	}
L302:
	;
	v3674 = v1376
	v3678 = v1380
	v3681 = v1383
	v3682 = v2201
	v3683 = v1385
	v3684 = v1386
	v3685 = v1387
	v3687 = v1389
	v3688 = v1390
	v3690 = v1392
	v3693 = v1395
	goto L301
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+40)) = int32(1)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+44))
	if v1270 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	goto L305
L305:
	;
	v1338 = v1168
	v1342 = v1242
	v1345 = v1175
	v1347 = v1243
	v1348 = v1178
	v1349 = v1244
	v1351 = v1181
	v1352 = v1245
	v1353 = v1260
	v1354 = v1246
	v1357 = v1187
	goto L322
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+44)) = int32(1)
	goto L308
L307:
	;
	goto L308
L308:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+4))
	if v1275 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L309:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[62]))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1279
	goto L311
L310:
	;
	goto L311
L311:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+8))
	if v1281 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[58]))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+8)) = v1285
	goto L314
L313:
	;
	goto L314
L314:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+20))
	if v1287 != 0 {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+28)) = v1316
	v1320 = v1314 + v1313<<(uint(int32(2))%32)
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+80)) = v1322
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+36)) = v1322
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1320)))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1325)))
	*(*int32)(unsafe.Add(mBase, uint32(v1168)+4)) = v1326
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1168)+24)) = uint8(v1328)
	goto L305
L316:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1287+v1288<<(uint(int32(2))%32))))
	if v1292 != 0 {
		v1313 = v1288
		v1314 = v1287
		v1315 = v1292
		goto L315
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	F_boot_yyensure_buffer_stack(m, v1168)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L1
	} else {
		goto L320
	}
L319:
	;
	goto L318
L320:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+4))
	v1298 = F_boot_yy_create_buffer(m, v1297, v1168)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+20))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	v1302 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1300+v1301<<(uint(v1302)%32)))) = v1298
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+20))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1306+v1307<<(uint(v1302)%32))))
	v1313 = v1307
	v1314 = v1306
	v1315 = v1311
	goto L315
L322:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+36))
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1361))) = uint8(v1362)
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+20))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+12))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1364+v1365<<(uint(int32(2))%32))))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+28))
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1338)+44))
	v1373 = v1361
	v1376 = v1338
	v1377 = v1361
	v1378 = v1370 + v1371
	v1380 = v1342
	v1383 = v1345
	v1385 = v1347
	v1386 = v1348
	v1387 = v1349
	v1389 = v1351
	v1390 = v1352
	v1391 = v1353
	v1392 = v1354
	v1395 = v1357
	goto L324
L324:
	;
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1377))))
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1399)+uint32(_c_F_BootstrapModeMain[63]))))
	v1402 = v1378 << (uint(int32(1)) % 32)
	v1405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1402)+uint32(_c_F_BootstrapModeMain[64]))))
	if v1405 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+68)) = v1377
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+64)) = v1378
	goto L328
L327:
	;
	goto L328
L328:
	;
	v1410 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1402)+uint32(_c_F_BootstrapModeMain[65]))))
	v1411 = v1410 + v1400
	v1416 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1411<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v1416 != v1378 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1420 = v1400
	v1423 = v1378
	v1424 = v1400
	goto L332
L330:
	;
	v1478 = v1411
	goto L331
L331:
	;
	v1495 = int32(1)
	v1501 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1478<<(uint(v1495)%32))+uint32(_c_F_BootstrapModeMain[67]))))
	if v1501 != int32(127) {
		v1377 = v1377 + v1495
		v1378 = v1501
		goto L324
	} else {
		goto L338
	}
L332:
	;
	v1448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1423<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[68]))))
	if int32(128) <= v1448 {
		goto L334
	} else {
		goto L335
	}
L333:
	;
	v1478 = v1460
	goto L331
L334:
	;
	v1451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1424)+uint32(_c_F_BootstrapModeMain[69]))))
	v1452 = v1451
	goto L336
L335:
	;
	v1452 = v1420
	goto L336
L336:
	;
	v1454 = v1452 & int32(255)
	v1455 = int32(1)
	v1459 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1448<<(uint(v1455)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v1460 = v1454 + v1459
	v1465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1460<<(uint(v1455)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v1465 != v1448&int32(_a_F_BootstrapModeMain_63) {
		v1420 = v1452
		v1423 = v1448
		v1424 = v1454
		goto L332
	} else {
		goto L337
	}
L337:
	;
	goto L333
L338:
	;
	v1505 = v1373
	goto L339
L339:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+64))
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+68))
	v1533 = v1505
	v1537 = v1530
	v1540 = v1531
	goto L341
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+80)) = v1533
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+32)) = v1540 - v1533
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1540))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1376)+24)) = uint8(v1561)
	v1563 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1540))) = uint8(v1563)
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v1540
	v1570 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1537<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[64]))))
	v1576 = v1570
	goto L343
L343:
	;
	switch v1576 {
	case 0:
		goto L389
	case 1:
		goto L388
	case 2:
		goto L387
	case 3:
		goto L386
	case 4:
		goto L385
	case 5:
		goto L384
	case 6:
		goto L383
	case 7:
		goto L382
	case 8:
		goto L381
	case 9:
		goto L380
	case 10:
		goto L379
	case 11:
		goto L378
	case 12:
		goto L377
	case 13:
		goto L376
	case 14:
		goto L375
	case 15:
		goto L374
	case 16:
		goto L373
	case 17:
		goto L372
	case 18:
		goto L371
	case 19:
		goto L370
	case 20:
		goto L369
	case 21:
		goto L368
	case 22:
		goto L367
	case 23:
		goto L366
	case 24:
		goto L365
	case 25:
		goto L364
	case 26:
		goto L363
	case 27:
		goto L362
	case 28:
		goto L361
	case 29:
		goto L360
	case 30:
		goto L357
	case 31:
		goto L356
	case 32:
		goto L355
	case 33:
		v2201 = int32(0)
		goto L358
	default:
		goto L354
	}
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v3635
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+48)) = int32(0)
	v3664 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+44))
	v3668 = base.I32_div_s(v3664-int32(1), int32(2))
	v1576 = v3668 + int32(33)
	goto L343
L346:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_64))
	mBase = m.M
	v3634 = m.ExcPending
	if v3634 != 0 {
		goto L1
	} else {
		goto L663
	}
L347:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_65))
	mBase = m.M
	v3628 = m.ExcPending
	if v3628 != 0 {
		goto L1
	} else {
		goto L662
	}
L348:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L349:
	;
	v3480 = v3459 + v3465
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v3480
	v3485 = *(*int32)(unsafe.Add(mBase, uint32(v3460+v3463<<(uint(int32(2))%32))))
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(v3485)+28))
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+44))
	v3488 = v3486 + v3487
	if base.Ui32(v3480) <= base.Ui32(v3454) {
		v1533 = v3454
		v1537 = v3488
		v1540 = v3480
		goto L341
	} else {
		goto L643
	}
L350:
	;
	v3101 = *(*int32)(unsafe.Add(mBase, uint32(v3100)))
	*(*int32)(unsafe.Add(mBase, uint32(v3101)+16)) = v3078
	v3104 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	if v3104 != 0 {
		v3227 = int32(0)
		goto L587
	} else {
		goto L588
	}
L351:
	;
	v3069 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3070 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3078 = v3047
	v3100 = v3069 + v3070<<(uint(int32(2))%32)
	goto L350
L352:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_66))
	mBase = m.M
	v3042 = m.ExcPending
	if v3042 != 0 {
		goto L1
	} else {
		goto L586
	}
L353:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_67))
	mBase = m.M
	v3039 = m.ExcPending
	if v3039 != 0 {
		goto L1
	} else {
		goto L585
	}
L354:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_68))
	mBase = m.M
	v3036 = m.ExcPending
	if v3036 != 0 {
		goto L1
	} else {
		goto L584
	}
L355:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1540))) = uint8(v2265)
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2271 = v2267 + v2268<<(uint(int32(2))%32)
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(v2271)))
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2272)+44))
	if v2273 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L356:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2244 {
		goto L471
	} else {
		goto L472
	}
L357:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2205 {
		goto L465
	} else {
		goto L466
	}
L358:
	;
	m.G0 = v1391 + int32(16)
	goto L302
L359:
	;
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2196))) = v2194
	v2201 = int32(258)
	goto L358
L360:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2174 {
		goto L461
	} else {
		goto L462
	}
L361:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2153 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2153 {
		goto L457
	} else {
		goto L458
	}
L362:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2131 {
		goto L454
	} else {
		goto L455
	}
L363:
	;
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2110 {
		goto L451
	} else {
		goto L452
	}
L364:
	;
	v2089 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2089 {
		goto L448
	} else {
		goto L449
	}
L365:
	;
	v2068 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2068 {
		goto L445
	} else {
		goto L446
	}
L366:
	;
	v2047 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2047 {
		goto L442
	} else {
		goto L443
	}
L367:
	;
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2026 {
		goto L439
	} else {
		goto L440
	}
L368:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v2005 {
		goto L436
	} else {
		goto L437
	}
L369:
	;
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1984 {
		goto L433
	} else {
		goto L434
	}
L370:
	;
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1963 {
		goto L430
	} else {
		goto L431
	}
L371:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1942 {
		goto L427
	} else {
		goto L428
	}
L372:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1921 {
		goto L424
	} else {
		goto L425
	}
L373:
	;
	v1904 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1904 <= int32(0) {
		v1338 = v1376
		v1342 = v1380
		v1345 = v1383
		v1347 = v1385
		v1348 = v1386
		v1349 = v1387
		v1351 = v1389
		v1352 = v1390
		v1353 = v1391
		v1354 = v1392
		v1357 = v1395
		goto L322
	} else {
		goto L423
	}
L374:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1887 <= int32(0) {
		v1338 = v1376
		v1342 = v1380
		v1345 = v1383
		v1347 = v1385
		v1348 = v1386
		v1349 = v1387
		v1351 = v1389
		v1352 = v1390
		v1353 = v1391
		v1354 = v1392
		v1357 = v1395
		goto L322
	} else {
		goto L422
	}
L375:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1859 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1860 {
		goto L419
	} else {
		goto L420
	}
L376:
	;
	v1840 = int32(262)
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1841 <= int32(0) {
		v2201 = v1840
		goto L358
	} else {
		goto L418
	}
L377:
	;
	v1822 = int32(261)
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1823 <= int32(0) {
		v2201 = v1822
		goto L358
	} else {
		goto L417
	}
L378:
	;
	v1804 = int32(260)
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1805 <= int32(0) {
		v2201 = v1804
		goto L358
	} else {
		goto L416
	}
L379:
	;
	v1786 = int32(259)
	v1787 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1787 <= int32(0) {
		v2201 = v1786
		goto L358
	} else {
		goto L415
	}
L380:
	;
	v1768 = int32(263)
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if v1769 <= int32(0) {
		v2201 = v1768
		goto L358
	} else {
		goto L414
	}
L381:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1747 {
		goto L411
	} else {
		goto L412
	}
L382:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1726 {
		goto L408
	} else {
		goto L409
	}
L383:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1705 {
		goto L405
	} else {
		goto L406
	}
L384:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1684 {
		goto L402
	} else {
		goto L403
	}
L385:
	;
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1663 {
		goto L399
	} else {
		goto L400
	}
L386:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1642 {
		goto L396
	} else {
		goto L397
	}
L387:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1621 {
		goto L393
	} else {
		goto L394
	}
L388:
	;
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+32))
	if int32(0) < v1600 {
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1540))) = uint8(v1598)
	v1505 = v1533
	goto L339
L390:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1603+v1604<<(uint(int32(2))%32))))
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1609+v1600-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1608)+28)) = base.B2i32(v1613 == int32(10))
	goto L392
L391:
	;
	goto L392
L392:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1617))) = int32(_a_F_BootstrapModeMain_69)
	v2201 = int32(264)
	goto L358
L393:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1624+v1625<<(uint(int32(2))%32))))
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1630+v1621-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1629)+28)) = base.B2i32(v1634 == int32(10))
	goto L395
L394:
	;
	goto L395
L395:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1638))) = int32(_a_F_BootstrapModeMain_70)
	v2201 = int32(265)
	goto L358
L396:
	;
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1645+v1646<<(uint(int32(2))%32))))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651+v1642-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1650)+28)) = base.B2i32(v1655 == int32(10))
	goto L398
L397:
	;
	goto L398
L398:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1659))) = int32(_a_F_BootstrapModeMain_71)
	v2201 = int32(266)
	goto L358
L399:
	;
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1666+v1667<<(uint(int32(2))%32))))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672+v1663-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1671)+28)) = base.B2i32(v1676 == int32(10))
	goto L401
L400:
	;
	goto L401
L401:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1680))) = int32(_a_F_BootstrapModeMain_72)
	v2201 = int32(276)
	goto L358
L402:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1687+v1688<<(uint(int32(2))%32))))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1693+v1684-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1692)+28)) = base.B2i32(v1697 == int32(10))
	goto L404
L403:
	;
	goto L404
L404:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1701))) = int32(_a_F_BootstrapModeMain_73)
	v2201 = int32(277)
	goto L358
L405:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1708+v1709<<(uint(int32(2))%32))))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1714+v1705-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1713)+28)) = base.B2i32(v1718 == int32(10))
	goto L407
L406:
	;
	goto L407
L407:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1722))) = int32(_a_F_BootstrapModeMain_74)
	v2201 = int32(278)
	goto L358
L408:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1730 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1729+v1730<<(uint(int32(2))%32))))
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1735+v1726-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1734)+28)) = base.B2i32(v1739 == int32(10))
	goto L410
L409:
	;
	goto L410
L410:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1743))) = int32(_a_F_BootstrapModeMain_75)
	v2201 = int32(279)
	goto L358
L411:
	;
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v1750+v1751<<(uint(int32(2))%32))))
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1756+v1747-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1755)+28)) = base.B2i32(v1760 == int32(10))
	goto L413
L412:
	;
	goto L413
L413:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1764))) = int32(_a_F_BootstrapModeMain_76)
	v2201 = int32(267)
	goto L358
L414:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1773 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1777 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1773<<(uint(int32(2))%32))))
	v1778 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1778+v1769-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1777)+28)) = base.B2i32(v1782 == int32(10))
	v2201 = v1768
	goto L358
L415:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(v1790+v1791<<(uint(int32(2))%32))))
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1800 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1796+v1787-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1795)+28)) = base.B2i32(v1800 == int32(10))
	v2201 = v1786
	goto L358
L416:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1808+v1809<<(uint(int32(2))%32))))
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1814+v1805-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1813)+28)) = base.B2i32(v1818 == int32(10))
	v2201 = v1804
	goto L358
L417:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1826+v1827<<(uint(int32(2))%32))))
	v1832 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1832+v1823-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1831)+28)) = base.B2i32(v1836 == int32(10))
	v2201 = v1822
	goto L358
L418:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1844+v1845<<(uint(int32(2))%32))))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850+v1841-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1849)+28)) = base.B2i32(v1854 == int32(10))
	v2201 = v1840
	goto L358
L419:
	;
	v1866 = *(*int32)(unsafe.Add(mBase, uint32(v1859+v1858<<(uint(int32(2))%32))))
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1867+v1860-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1866)+28)) = base.B2i32(v1871 == int32(10))
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1877 = v1875
	v1878 = v1876
	goto L421
L420:
	;
	v1877 = v1859
	v1878 = v1858
	goto L421
L421:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v1878<<(uint(int32(2))%32)+v1877)))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1882)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1882)+32)) = v1883 + int32(1)
	v1338 = v1376
	v1342 = v1380
	v1345 = v1383
	v1347 = v1385
	v1348 = v1386
	v1349 = v1387
	v1351 = v1389
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1357 = v1395
	goto L322
L422:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1890+v1891<<(uint(int32(2))%32))))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1896+v1887-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1895)+28)) = base.B2i32(v1900 == int32(10))
	v1338 = v1376
	v1342 = v1380
	v1345 = v1383
	v1347 = v1385
	v1348 = v1386
	v1349 = v1387
	v1351 = v1389
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1357 = v1395
	goto L322
L423:
	;
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1907+v1908<<(uint(int32(2))%32))))
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1913+v1904-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1912)+28)) = base.B2i32(v1917 == int32(10))
	v1338 = v1376
	v1342 = v1380
	v1345 = v1383
	v1347 = v1385
	v1348 = v1386
	v1349 = v1387
	v1351 = v1389
	v1352 = v1390
	v1353 = v1391
	v1354 = v1392
	v1357 = v1395
	goto L322
L424:
	;
	v1924 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1924+v1925<<(uint(int32(2))%32))))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1930+v1921-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1929)+28)) = base.B2i32(v1934 == int32(10))
	goto L426
L425:
	;
	goto L426
L426:
	;
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1938))) = int32(_a_F_BootstrapModeMain_77)
	v2201 = int32(268)
	goto L358
L427:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1946 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v1945+v1946<<(uint(int32(2))%32))))
	v1951 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1951+v1942-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1950)+28)) = base.B2i32(v1955 == int32(10))
	goto L429
L428:
	;
	goto L429
L429:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1959))) = int32(_a_F_BootstrapModeMain_78)
	v2201 = int32(272)
	goto L358
L430:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1966+v1967<<(uint(int32(2))%32))))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1972+v1963-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1971)+28)) = base.B2i32(v1976 == int32(10))
	goto L432
L431:
	;
	goto L432
L432:
	;
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v1980))) = int32(_a_F_BootstrapModeMain_79)
	v2201 = int32(273)
	goto L358
L433:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1987+v1988<<(uint(int32(2))%32))))
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v1997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1993+v1984-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v1992)+28)) = base.B2i32(v1997 == int32(10))
	goto L435
L434:
	;
	goto L435
L435:
	;
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2001))) = int32(_a_F_BootstrapModeMain_80)
	v2201 = int32(274)
	goto L358
L436:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2013 = *(*int32)(unsafe.Add(mBase, uint32(v2008+v2009<<(uint(int32(2))%32))))
	v2014 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2018 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2014+v2005-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2013)+28)) = base.B2i32(v2018 == int32(10))
	goto L438
L437:
	;
	goto L438
L438:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2022))) = int32(_a_F_BootstrapModeMain_81)
	v2201 = int32(269)
	goto L358
L439:
	;
	v2029 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2030 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2034 = *(*int32)(unsafe.Add(mBase, uint32(v2029+v2030<<(uint(int32(2))%32))))
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2035+v2026-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2034)+28)) = base.B2i32(v2039 == int32(10))
	goto L441
L440:
	;
	goto L441
L441:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2043))) = int32(_a_F_BootstrapModeMain_82)
	v2201 = int32(270)
	goto L358
L442:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2050+v2051<<(uint(int32(2))%32))))
	v2056 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2056+v2047-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2055)+28)) = base.B2i32(v2060 == int32(10))
	goto L444
L443:
	;
	goto L444
L444:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2064))) = int32(_a_F_BootstrapModeMain_83)
	v2201 = int32(271)
	goto L358
L445:
	;
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(v2071+v2072<<(uint(int32(2))%32))))
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2077+v2068-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2076)+28)) = base.B2i32(v2081 == int32(10))
	goto L447
L446:
	;
	goto L447
L447:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2085))) = int32(_a_F_BootstrapModeMain_84)
	v2201 = int32(275)
	goto L358
L448:
	;
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v2092+v2093<<(uint(int32(2))%32))))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2098+v2089-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2097)+28)) = base.B2i32(v2102 == int32(10))
	goto L450
L449:
	;
	goto L450
L450:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2106))) = int32(_a_F_BootstrapModeMain_85)
	v2201 = int32(280)
	goto L358
L451:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2113+v2114<<(uint(int32(2))%32))))
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2119+v2110-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2118)+28)) = base.B2i32(v2123 == int32(10))
	goto L453
L452:
	;
	goto L453
L453:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2127))) = int32(_a_F_BootstrapModeMain_86)
	v2201 = int32(281)
	goto L358
L454:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2134+v2135<<(uint(int32(2))%32))))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2140+v2131-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2139)+28)) = base.B2i32(v2144 == int32(10))
	goto L456
L455:
	;
	goto L456
L456:
	;
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v2148))) = int32(_a_F_BootstrapModeMain_87)
	v2201 = int32(282)
	goto L358
L457:
	;
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v2156+v2157<<(uint(int32(2))%32))))
	v2165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2152+v2153-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2161)+28)) = base.B2i32(v2165 == int32(10))
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2170 = v2169
	goto L459
L458:
	;
	v2170 = v2152
	goto L459
L459:
	;
	v2171 = F_pstrdup(m, v2170)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	v2194 = v2171
	goto L359
L461:
	;
	v2177 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v2177+v2178<<(uint(int32(2))%32))))
	v2186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2173+v2174-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2182)+28)) = base.B2i32(v2186 == int32(10))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2191 = v2190
	goto L463
L462:
	;
	v2191 = v2173
	goto L463
L463:
	;
	v2192 = F_DeescapeQuotedString(m, v2191)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v2194 = v2192
	goto L359
L465:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2208+v2209<<(uint(int32(2))%32))))
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2214+v2205-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2213)+28)) = base.B2i32(v2218 == int32(10))
	goto L467
L466:
	;
	goto L467
L467:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2227 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2231 = *(*int32)(unsafe.Add(mBase, uint32(v2226+v2227<<(uint(int32(2))%32))))
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2231)+32))
	v2233 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1391)+4)) = v2233
	*(*int32)(unsafe.Add(mBase, uint32(v1391))) = v2232
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_88), v1391)
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_89), int32(124), int32(_a_F_BootstrapModeMain_90))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L1
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2252 = *(*int32)(unsafe.Add(mBase, uint32(v2247+v2248<<(uint(int32(2))%32))))
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2253+v2244-int32(1)))))
	*(*int32)(unsafe.Add(mBase, uint32(v2252)+28)) = base.B2i32(v2257 == int32(10))
	goto L473
L472:
	;
	goto L473
L473:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_91))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L1
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
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2272)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v2276
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2271)))
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2278))) = v2279
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2283 = int32(2)
	v2286 = *(*int32)(unsafe.Add(mBase, uint32(v2281+v2282<<(uint(v2283)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2286)+44)) = int32(1)
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v2289+v2290<<(uint(v2283)%32))))
	v2295 = v2294
	v2296 = v2289
	v2297 = v2290
	goto L477
L476:
	;
	v2295 = v2272
	v2296 = v2267
	v2297 = v2268
	goto L477
L477:
	;
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+36))
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+4))
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	v2301 = v2299 + v2300
	if base.Ui32(v2298) <= base.Ui32(v2301) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2303 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2306 = v2264 ^ int32(-1) + v1540
	v2307 = v2303 + v2306
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v2307
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2296+v2297<<(uint(int32(2))%32))))
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+28))
	v2314 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+44))
	v2315 = v2313 + v2314
	if int32(0) < v2306 {
		goto L481
	} else {
		goto L482
	}
L479:
	;
	goto L480
L480:
	;
	if base.Ui32(v2301+int32(1)) < base.Ui32(v2298) {
		goto L353
	} else {
		goto L513
	}
L481:
	;
	v2322 = v2303
	v2323 = v2315
	goto L484
L482:
	;
	v2457 = v2315
	goto L483
L483:
	;
	v2482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2457<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[64]))))
	if v2482 != 0 {
		goto L502
	} else {
		goto L503
	}
L484:
	;
	v2345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2322))))
	if v2345 != 0 {
		goto L486
	} else {
		goto L487
	}
L485:
	;
	v2457 = v2448
	goto L483
L486:
	;
	v2346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2345)+uint32(_c_F_BootstrapModeMain[63]))))
	v2347 = v2346
	goto L488
L487:
	;
	v2347 = int32(1)
	goto L488
L488:
	;
	v2349 = v2323 << (uint(int32(1)) % 32)
	v2352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2349)+uint32(_c_F_BootstrapModeMain[64]))))
	if v2352 != 0 {
		goto L489
	} else {
		goto L490
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+68)) = v2322
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+64)) = v2323
	goto L491
L490:
	;
	goto L491
L491:
	;
	v2356 = v2347 & int32(255)
	v2359 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2349)+uint32(_c_F_BootstrapModeMain[65]))))
	v2360 = v2356 + v2359
	v2365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2360<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v2365 != v2323 {
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v2369 = v2347
	v2372 = v2323
	v2373 = v2356
	goto L495
L493:
	;
	v2427 = v2360
	goto L494
L494:
	;
	v2444 = int32(1)
	v2448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2427<<(uint(v2444)%32))+uint32(_c_F_BootstrapModeMain[67]))))
	v2450 = v2322 + v2444
	if v2450 != v2307 {
		v2322 = v2450
		v2323 = v2448
		goto L484
	} else {
		goto L501
	}
L495:
	;
	v2397 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2372<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[68]))))
	if int32(128) <= v2397 {
		goto L497
	} else {
		goto L498
	}
L496:
	;
	v2427 = v2409
	goto L494
L497:
	;
	v2400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2373)+uint32(_c_F_BootstrapModeMain[69]))))
	v2401 = v2400
	goto L499
L498:
	;
	v2401 = v2369
	goto L499
L499:
	;
	v2403 = v2401 & int32(255)
	v2404 = int32(1)
	v2408 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2397<<(uint(v2404)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v2409 = v2403 + v2408
	v2414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2409<<(uint(v2404)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v2414 != v2397&int32(_a_F_BootstrapModeMain_63) {
		v2369 = v2401
		v2372 = v2397
		v2373 = v2403
		goto L495
	} else {
		goto L500
	}
L500:
	;
	goto L496
L501:
	;
	goto L485
L502:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+68)) = v2307
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+64)) = v2457
	goto L504
L503:
	;
	goto L504
L504:
	;
	v2485 = int32(1)
	v2489 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2457<<(uint(v2485)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v2491 = v2489 + v2485
	v2496 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2491<<(uint(v2485)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v2496 != v2457 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2503 = v2457
	goto L508
L506:
	;
	v2545 = v2491
	goto L507
L507:
	;
	if v2545 == int32(0) {
		v1505 = v2303
		goto L339
	} else {
		goto L511
	}
L508:
	;
	v2524 = int32(1)
	v2528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2503<<(uint(v2524)%32))+uint32(_c_F_BootstrapModeMain[68]))))
	v2529 = base.I32_extend16_s(v2528)
	v2534 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2529<<(uint(v2524)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v2536 = v2534 + v2524
	v2541 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2536<<(uint(v2524)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v2528 != v2541 {
		v2503 = v2529
		goto L508
	} else {
		goto L510
	}
L509:
	;
	v2545 = v2536
	goto L507
L510:
	;
	goto L509
L511:
	;
	v2575 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2545<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[67]))))
	if v2575 == int32(127) {
		v1505 = v2303
		goto L339
	} else {
		goto L512
	}
L512:
	;
	v2579 = v2307 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v2579
	v1373 = v2303
	v1377 = v2579
	v1378 = v2575
	goto L324
L513:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+80))
	v2585 = *(*int32)(unsafe.Add(mBase, uint32(v2295)+40))
	if v2585 == int32(0) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	if v2298-v2584 != int32(1) {
		v3454 = v2584
		v3459 = v2299
		v3460 = v2296
		v3463 = v2297
		v3465 = v2300
		goto L349
	} else {
		goto L517
	}
L515:
	;
	goto L516
L516:
	;
	v2593 = v2584 ^ int32(-1) + v2298
	if int32(0) < v2593 {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	v3635 = v2584
	goto L345
L518:
	;
	v2596 = int32(7)
	v2597 = v2593 & v2596
	if base.Ui32(v2298-v2584-int32(2)) < base.Ui32(v2596) {
		goto L523
	} else {
		goto L524
	}
L519:
	;
	v2753 = v2295
	v2757 = v2296
	v2760 = v2297
	goto L520
L520:
	;
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2753)+44))
	if v2777 == int32(2) {
		goto L533
	} else {
		goto L534
	}
L521:
	;
	v2745 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2746 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2750 = *(*int32)(unsafe.Add(mBase, uint32(v2745+v2746<<(uint(int32(2))%32))))
	v2753 = v2750
	v2757 = v2745
	v2760 = v2746
	goto L520
L522:
	;
	v2686 = v2659
	v2689 = v2662
	v2690 = int32(0)
	goto L530
L523:
	;
	v2659 = v2584
	v2662 = v2299
	goto L522
L524:
	;
	goto L525
L525:
	;
	v2608 = v2584
	v2611 = v2299
	v2612 = int32(0)
	goto L526
L526:
	;
	v2632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611))) = uint8(v2632)
	v2634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+1)) = uint8(v2634)
	v2636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+2)) = uint8(v2636)
	v2638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+3)) = uint8(v2638)
	v2640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+4)) = uint8(v2640)
	v2642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+5)) = uint8(v2642)
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+6)) = uint8(v2644)
	v2646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2608)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2611)+7)) = uint8(v2646)
	v2648 = int32(8)
	v2649 = v2611 + v2648
	v2651 = v2608 + v2648
	v2653 = v2612 + v2648
	if v2653 != v2593&int32(2147483640) {
		v2608 = v2651
		v2611 = v2649
		v2612 = v2653
		goto L526
	} else {
		goto L528
	}
L527:
	;
	if v2597 == int32(0) {
		goto L521
	} else {
		goto L529
	}
L528:
	;
	goto L527
L529:
	;
	v2659 = v2651
	v2662 = v2649
	goto L522
L530:
	;
	v2710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2686))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2689))) = uint8(v2710)
	v2712 = int32(1)
	v2717 = v2690 + v2712
	if v2717 != v2597 {
		v2686 = v2686 + v2712
		v2689 = v2689 + v2712
		v2690 = v2717
		goto L530
	} else {
		goto L532
	}
L531:
	;
	goto L521
L532:
	;
	goto L531
L533:
	;
	v2780 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v2780
	v3078 = v2780
	v3100 = v2757 + v2760<<(uint(int32(2))%32)
	goto L350
L534:
	;
	goto L535
L535:
	;
	v2786 = int32(0)
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(v2753)+12))
	v2788 = v2584 - v2298
	v2789 = v2787 + v2788
	if v2789 <= v2786 {
		goto L536
	} else {
		goto L537
	}
L536:
	;
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+36))
	v2793 = v2792
	v2795 = v2753
	v2802 = v2787
	goto L539
L537:
	;
	v2859 = v2753
	v2862 = v2789
	goto L538
L538:
	;
	v2883 = int32(_a_F_BootstrapModeMain_25)
	if base.Ui32(v2883) <= base.Ui32(v2862) {
		goto L555
	} else {
		goto L556
	}
L539:
	;
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v2795)+20))
	if v2819 == int32(0) {
		goto L541
	} else {
		goto L542
	}
L540:
	;
	v2859 = v2852
	v2862 = v2854
	goto L538
L541:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+4)) = int32(0)
	goto L346
L542:
	;
	goto L543
L543:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2795)+4))
	v2826 = v2802 << (uint(int32(1)) % 32)
	if v2826 <= int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2830 = base.I32_div_s(v2802, int32(8))
	v2832 = v2830 + v2802
	goto L546
L545:
	;
	v2832 = v2826
	goto L546
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+12)) = v2832
	v2835 = v2832 + int32(2)
	if v2824 != 0 {
		goto L548
	} else {
		goto L549
	}
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+4)) = v2840
	if v2840 == int32(0) {
		goto L346
	} else {
		goto L553
	}
L548:
	;
	v2836 = F_repalloc(m, v2824, v2835)
	mBase = m.M
	v2837 = m.ExcPending
	if v2837 != 0 {
		goto L1
	} else {
		goto L551
	}
L549:
	;
	goto L550
L550:
	;
	v2838 = F_palloc(m, v2835)
	mBase = m.M
	v2839 = m.ExcPending
	if v2839 != 0 {
		goto L1
	} else {
		goto L552
	}
L551:
	;
	v2840 = v2836
	goto L547
L552:
	;
	v2840 = v2838
	goto L547
L553:
	;
	v2845 = v2840 + (v2793 - v2824)
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v2845
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v2847+v2848<<(uint(int32(2))%32))))
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2852)+12))
	v2854 = v2853 + v2788
	if v2854 <= int32(0) {
		v2793 = v2845
		v2795 = v2852
		v2802 = v2853
		goto L539
	} else {
		goto L554
	}
L554:
	;
	goto L540
L555:
	;
	v2886 = v2883
	goto L557
L556:
	;
	v2886 = v2862
	goto L557
L557:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v2859)+24))
	if v2887 != 0 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v2892 = v2786
	goto L562
L559:
	;
	goto L560
L560:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(0)
	v2962 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2963 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2967 = *(*int32)(unsafe.Add(mBase, uint32(v2962+v2963<<(uint(int32(2))%32))))
	v2968 = *(*int32)(unsafe.Add(mBase, uint32(v2967)+4))
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v2972 = F_fread(m, v2968+v2593, int32(1), v2886, v2971)
	mBase = m.M
	v2973 = m.ExcPending
	if v2973 != 0 {
		goto L1
	} else {
		goto L573
	}
L561:
	;
	switch v2918 {
	case 0:
		goto L569
	default:
		v2957 = v2932
		goto L567
	case 11:
		goto L568
	}
L562:
	;
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v2915 = F_do_getc(m, v2914)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L1
	} else {
		goto L565
	}
L563:
	;
	v2932 = v2886
	goto L561
L564:
	;
	v2919 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(v2919+v2920<<(uint(int32(2))%32))))
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2924)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v2925+v2593+v2892))) = uint8(v2915)
	v2930 = v2892 + int32(1)
	if v2930 != v2886 {
		v2892 = v2930
		goto L562
	} else {
		goto L566
	}
L565:
	;
	v2918 = v2915 + int32(1)
	switch v2918 {
	case 0, 11:
		v2932 = v2892
		goto L561
	default:
		goto L564
	}
L566:
	;
	goto L563
L567:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v2957
	v3047 = v2957
	goto L351
L568:
	;
	v2944 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v2945 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v2944+v2945<<(uint(int32(2))%32))))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v2949)+4))
	v2953 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v2950+v2593+v2932))) = uint8(v2953)
	v2957 = v2932 + int32(1)
	goto L567
L569:
	;
	v2933 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v2934 = *(*int32)(unsafe.Add(mBase, uint32(v2933)))
	goto L570
L570:
	;
	if int32(base.Ui32(v2934)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v2957 = v2932
		goto L567
	} else {
		goto L571
	}
L571:
	;
	F_yy_fatal_error_1(m, int32(_a_F_BootstrapModeMain_66))
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L1
	} else {
		goto L572
	}
L572:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L573:
	;
	v2978 = v2972
	goto L574
L574:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v2978
	if v2978 != 0 {
		v3047 = v2978
		goto L351
	} else {
		goto L576
	}
L576:
	;
	v3001 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v3002 = *(*int32)(unsafe.Add(mBase, uint32(v3001)))
	goto L577
L577:
	;
	if int32(base.Ui32(v3002)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v3047 = int32(0)
	goto L351
L579:
	;
	goto L580
L580:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	if v3011 != int32(27) {
		goto L352
	} else {
		goto L581
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = int32(0)
	v3017 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v3018 = *(*int32)(unsafe.Add(mBase, uint32(v3017)))
	*(*int32)(unsafe.Add(mBase, uint32(v3017))) = v3018 & int32(-49)
	goto L582
L582:
	;
	v3022 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3027 = *(*int32)(unsafe.Add(mBase, uint32(v3022+v3023<<(uint(int32(2))%32))))
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v3027)+4))
	v3031 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v3032 = F_fread(m, v3028+v2593, int32(1), v2886, v3031)
	mBase = m.M
	v3033 = m.ExcPending
	if v3033 != 0 {
		goto L1
	} else {
		goto L583
	}
L583:
	;
	v2978 = v3032
	goto L574
L584:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L585:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L586:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L587:
	;
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	v3229 = v3228 + v2593
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3235 = *(*int32)(unsafe.Add(mBase, uint32(v3230+v3231<<(uint(int32(2))%32))))
	v3236 = *(*int32)(unsafe.Add(mBase, uint32(v3235)+12))
	if v3236 < v3229 {
		goto L611
	} else {
		goto L612
	}
L588:
	;
	if v2593 == int32(0) {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	if v3108 != 0 {
		goto L594
	} else {
		goto L595
	}
L590:
	;
	goto L591
L591:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3215 = int32(2)
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(v3213+v3214<<(uint(v3215)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3218)+44)) = v3215
	v3227 = v3215
	goto L587
L592:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3177)+40)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v3177))) = v3107
	v3182 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	if v3182 != 0 {
		goto L607
	} else {
		goto L608
	}
L593:
	;
	v3132 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	v3133 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(v3130+v3133<<(uint(int32(2))%32))))
	if v3137 == int32(0) {
		goto L601
	} else {
		goto L602
	}
L594:
	;
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v3108+v3109<<(uint(int32(2))%32))))
	if v3113 != 0 {
		v3130 = v3108
		goto L593
	} else {
		goto L597
	}
L595:
	;
	goto L596
L596:
	;
	F_boot_yyensure_buffer_stack(m, v1376)
	mBase = m.M
	v3115 = m.ExcPending
	if v3115 != 0 {
		goto L1
	} else {
		goto L598
	}
L597:
	;
	goto L596
L598:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+4))
	v3117 = F_boot_yy_create_buffer(m, v3116, v1376)
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L1
	} else {
		goto L599
	}
L599:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v3119+v3120<<(uint(int32(2))%32)))) = v3117
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	if v3125 != 0 {
		v3130 = v3125
		goto L593
	} else {
		goto L600
	}
L600:
	;
	v3127 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39]))
	v3175 = v3127
	v3177 = int32(0)
	goto L592
L601:
	;
	v3175 = v3132
	v3177 = int32(0)
	goto L592
L602:
	;
	goto L603
L603:
	;
	v3141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+16)) = v3141
	v3143 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3143))) = uint8(v3141)
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v3146)+1)) = uint8(v3141)
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+44)) = v3141
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+28)) = int32(1)
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3137)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3137)+8)) = v3153
	v3155 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	if v3155 == v3141 {
		v3175 = v3132
		v3177 = v3137
		goto L592
	} else {
		goto L604
	}
L604:
	;
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3161 = v3155 + v3158<<(uint(int32(2))%32)
	v3162 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	if v3137 != v3162 {
		v3175 = v3132
		v3177 = v3137
		goto L592
	} else {
		goto L605
	}
L605:
	;
	v3164 = *(*int32)(unsafe.Add(mBase, uint32(v3162)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v3164
	v3166 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v3166)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+80)) = v3167
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v3167
	v3170 = *(*int32)(unsafe.Add(mBase, uint32(v3161)))
	v3171 = *(*int32)(unsafe.Add(mBase, uint32(v3170)))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+4)) = v3171
	v3173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1376)+24)) = uint8(v3173)
	v3175 = v3132
	v3177 = v3137
	goto L592
L606:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3177)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[39])) = v3175
	v3195 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3196 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3199 = v3195 + v3196<<(uint(int32(2))%32)
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(v3200)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v3201
	v3203 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(v3203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v3204
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+80)) = v3204
	v3207 = *(*int32)(unsafe.Add(mBase, uint32(v3199)))
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v3207)))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+4)) = v3208
	v3210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3204))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1376)+24)) = uint8(v3210)
	v3227 = int32(1)
	goto L587
L607:
	;
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3182+v3183<<(uint(int32(2))%32))))
	if v3177 == v3187 {
		goto L606
	} else {
		goto L610
	}
L608:
	;
	goto L609
L609:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3177)+32)) = int64(1)
	goto L606
L610:
	;
	goto L609
L611:
	;
	v3240 = v3229 + v3228>>(uint(int32(1))%32)
	v3241 = *(*int32)(unsafe.Add(mBase, uint32(v3235)+4))
	if v3241 != 0 {
		goto L615
	} else {
		goto L616
	}
L612:
	;
	v3271 = v3230
	v3272 = v3229
	v3273 = v3231
	goto L613
L613:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+28)) = v3272
	v3275 = int32(2)
	v3278 = *(*int32)(unsafe.Add(mBase, uint32(v3271+v3273<<(uint(v3275)%32))))
	v3279 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+4))
	v3281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3279+v3272))) = uint8(v3281)
	v3283 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3283+v3284<<(uint(v3275)%32))))
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(v3288)+4))
	v3290 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v3289+v3290)+1)) = uint8(v3281)
	v3294 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3295 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3298 = v3294 + v3295<<(uint(v3275)%32)
	v3299 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3300 = *(*int32)(unsafe.Add(mBase, uint32(v3299)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+80)) = v3300
	if v3227 == int32(1) {
		v3635 = v3300
		goto L345
	} else {
		goto L621
	}
L614:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3249 = int32(2)
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(v3247+v3248<<(uint(v3249)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3252)+4)) = v3246
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3254+v3255<<(uint(v3249)%32))))
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(v3259)+4))
	if v3260 == int32(0) {
		goto L347
	} else {
		goto L620
	}
L615:
	;
	v3242 = F_repalloc(m, v3241, v3240)
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L1
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	v3244 = F_palloc(m, v3240)
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L1
	} else {
		goto L619
	}
L618:
	;
	v3246 = v3242
	goto L614
L619:
	;
	v3246 = v3244
	goto L614
L620:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3259)+12)) = v3240 - int32(2)
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+12))
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	v3269 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+20))
	v3271 = v3269
	v3272 = v3267 + v2593
	v3273 = v3266
	goto L613
L621:
	;
	switch v3227 - int32(1) {
	case 0:
		goto L348
	case 1:
		goto L622
	default:
		goto L623
	}
L622:
	;
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+28))
	v3452 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3453 = *(*int32)(unsafe.Add(mBase, uint32(v3452)+4))
	v3454 = v3300
	v3459 = v3453
	v3460 = v3294
	v3463 = v3295
	v3465 = v3451
	goto L349
L623:
	;
	v3308 = v2264 ^ int32(-1) + v1540
	v3309 = v3300 + v3308
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+36)) = v3309
	v3311 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(v3311)+28))
	v3313 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+44))
	v3314 = v3312 + v3313
	if v3308 <= int32(0) {
		v1373 = v3300
		v1377 = v3309
		v1378 = v3314
		goto L324
	} else {
		goto L624
	}
L624:
	;
	v3322 = v3314
	v3325 = v3300
	goto L625
L625:
	;
	v3344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3325))))
	if v3344 != 0 {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	v1373 = v3300
	v1377 = v3309
	v1378 = v3447
	goto L324
L627:
	;
	v3345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3344)+uint32(_c_F_BootstrapModeMain[63]))))
	v3346 = v3345
	goto L629
L628:
	;
	v3346 = int32(1)
	goto L629
L629:
	;
	v3348 = v3322 << (uint(int32(1)) % 32)
	v3351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3348)+uint32(_c_F_BootstrapModeMain[64]))))
	if v3351 != 0 {
		goto L630
	} else {
		goto L631
	}
L630:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+68)) = v3325
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+64)) = v3322
	goto L632
L631:
	;
	goto L632
L632:
	;
	v3355 = v3346 & int32(255)
	v3358 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3348)+uint32(_c_F_BootstrapModeMain[65]))))
	v3359 = v3355 + v3358
	v3364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3359<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v3364 != v3322 {
		goto L633
	} else {
		goto L634
	}
L633:
	;
	v3368 = v3346
	v3371 = v3322
	v3372 = v3355
	goto L636
L634:
	;
	v3426 = v3359
	goto L635
L635:
	;
	v3443 = int32(1)
	v3447 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3426<<(uint(v3443)%32))+uint32(_c_F_BootstrapModeMain[67]))))
	v3449 = v3325 + v3443
	if v3309 != v3449 {
		v3322 = v3447
		v3325 = v3449
		goto L625
	} else {
		goto L642
	}
L636:
	;
	v3396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3371<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[68]))))
	if int32(128) <= v3396 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v3426 = v3408
	goto L635
L638:
	;
	v3399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3372)+uint32(_c_F_BootstrapModeMain[69]))))
	v3400 = v3399
	goto L640
L639:
	;
	v3400 = v3368
	goto L640
L640:
	;
	v3402 = v3400 & int32(255)
	v3403 = int32(1)
	v3407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3396<<(uint(v3403)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v3408 = v3402 + v3407
	v3413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3408<<(uint(v3403)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v3413 != v3396&int32(_a_F_BootstrapModeMain_63) {
		v3368 = v3400
		v3371 = v3396
		v3372 = v3402
		goto L636
	} else {
		goto L641
	}
L641:
	;
	goto L637
L642:
	;
	goto L626
L643:
	;
	v3494 = v3454
	v3495 = v3488
	goto L644
L644:
	;
	v3517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3494))))
	if v3517 != 0 {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v1533 = v3454
	v1537 = v3622
	v1540 = v3480
	goto L341
L646:
	;
	v3518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3517)+uint32(_c_F_BootstrapModeMain[63]))))
	v3519 = v3518
	goto L648
L647:
	;
	v3519 = int32(1)
	goto L648
L648:
	;
	v3524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3495<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[64]))))
	if v3524 != 0 {
		goto L649
	} else {
		goto L650
	}
L649:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+68)) = v3494
	*(*int32)(unsafe.Add(mBase, uint32(v1376)+64)) = v3495
	goto L651
L650:
	;
	goto L651
L651:
	;
	v3528 = v3519 & int32(255)
	v3529 = int32(1)
	v3533 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3495<<(uint(v3529)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v3534 = v3528 + v3533
	v3539 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3534<<(uint(v3529)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v3539 != v3495 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3543 = v3519
	v3546 = v3495
	v3547 = v3528
	goto L655
L653:
	;
	v3601 = v3534
	goto L654
L654:
	;
	v3618 = int32(1)
	v3622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3601<<(uint(v3618)%32))+uint32(_c_F_BootstrapModeMain[67]))))
	v3624 = v3494 + v3618
	if v3624 != v3480 {
		v3494 = v3624
		v3495 = v3622
		goto L644
	} else {
		goto L661
	}
L655:
	;
	v3571 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3546<<(uint(int32(1))%32))+uint32(_c_F_BootstrapModeMain[68]))))
	if int32(128) <= v3571 {
		goto L657
	} else {
		goto L658
	}
L656:
	;
	v3601 = v3583
	goto L654
L657:
	;
	v3574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3547)+uint32(_c_F_BootstrapModeMain[69]))))
	v3575 = v3574
	goto L659
L658:
	;
	v3575 = v3543
	goto L659
L659:
	;
	v3577 = v3575 & int32(255)
	v3578 = int32(1)
	v3582 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3571<<(uint(v3578)%32))+uint32(_c_F_BootstrapModeMain[65]))))
	v3583 = v3577 + v3582
	v3588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3583<<(uint(v3578)%32))+uint32(_c_F_BootstrapModeMain[66]))))
	if v3588 != v3571&int32(_a_F_BootstrapModeMain_63) {
		v3543 = v3575
		v3546 = v3571
		v3547 = v3577
		goto L655
	} else {
		goto L660
	}
L660:
	;
	goto L656
L661:
	;
	goto L645
L662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L663:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L664:
	;
	v3699 = int32(0)
	v3707 = v3699
	v3708 = v3699
	goto L298
L665:
	;
	goto L666
L666:
	;
	if v3682 == int32(256) {
		v6283 = v3681
		v6285 = v3683
		v6287 = v3685
		goto L267
	} else {
		goto L667
	}
L667:
	;
	if base.Ui32(int32(282)) < base.Ui32(v3682) {
		v3707 = v3682
		v3708 = int32(2)
		goto L298
	} else {
		goto L668
	}
L668:
	;
	v3706 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3682)+uint32(_c_F_BootstrapModeMain[70]))))
	v3707 = v3682
	v3708 = v3706
	goto L298
L669:
	;
	v3712 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3709)+uint32(_c_F_BootstrapModeMain[71]))))
	if v3708 != v3712 {
		v3730 = v3674
		v3734 = v3678
		v3737 = v3681
		v3738 = v3707
		v3739 = v3683
		v3740 = v3684
		v3741 = v3685
		v3743 = v3687
		v3744 = v3688
		v3746 = v3690
		v3749 = v3693
		goto L296
	} else {
		goto L670
	}
L670:
	;
	v3714 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3709)+uint32(_c_F_BootstrapModeMain[72]))))
	if int32(0) < v3714 {
		goto L671
	} else {
		goto L672
	}
L671:
	;
	v3717 = *(*int32)(unsafe.Add(mBase, uint32(v3681)+1132))
	*(*int32)(unsafe.Add(mBase, uint32(v3678)+4)) = v3717
	v6243 = v3674
	v6247 = v3678 + int32(4)
	v6250 = v3681
	v6251 = int32(-2)
	v6252 = v3683
	v6253 = v3714
	v6254 = v3685
	v6256 = v3687
	v6257 = v3688
	v6259 = v3690
	v6262 = v3693 - base.B2i32(v3693 != int32(0))
	goto L294
L672:
	;
	goto L673
L673:
	;
	v3761 = v3674
	v3765 = v3678
	v3768 = v3681
	v3769 = v3707
	v3770 = v3683
	v3771 = int32(0) - v3714
	v3772 = v3685
	v3774 = v3687
	v3775 = v3688
	v3777 = v3690
	v3780 = v3693
	goto L295
L674:
	;
	v3761 = v3730
	v3765 = v3734
	v3768 = v3737
	v3769 = v3738
	v3770 = v3739
	v3771 = v3755
	v3772 = v3741
	v3774 = v3743
	v3775 = v3744
	v3777 = v3746
	v3780 = v3749
	goto L295
L675:
	;
	v6215 = v3765 - v3787<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v6215)+4)) = v6189
	v6218 = v6215 + int32(4)
	v6219 = v3772 - v3787
	v6220 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6219))))
	v6223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3771)+uint32(_c_F_BootstrapModeMain[73]))))
	v6226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6223)+uint32(_c_F_BootstrapModeMain[74]))))
	v6227 = v6220 + v6226
	if base.Ui32(int32(169)) < base.Ui32(v6227) {
		goto L1200
	} else {
		goto L1201
	}
L676:
	;
	v6184 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6185 = F_pstrdup(m, v6184)
	mBase = m.M
	v6186 = m.ExcPending
	if v6186 != 0 {
		goto L1
	} else {
		goto L1199
	}
L677:
	;
	v6181 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6182 = F_pstrdup(m, v6181)
	mBase = m.M
	v6183 = m.ExcPending
	if v6183 != 0 {
		goto L1
	} else {
		goto L1198
	}
L678:
	;
	v6178 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6179 = F_pstrdup(m, v6178)
	mBase = m.M
	v6180 = m.ExcPending
	if v6180 != 0 {
		goto L1
	} else {
		goto L1197
	}
L679:
	;
	v6175 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6176 = F_pstrdup(m, v6175)
	mBase = m.M
	v6177 = m.ExcPending
	if v6177 != 0 {
		goto L1
	} else {
		goto L1196
	}
L680:
	;
	v6172 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6173 = F_pstrdup(m, v6172)
	mBase = m.M
	v6174 = m.ExcPending
	if v6174 != 0 {
		goto L1
	} else {
		goto L1195
	}
L681:
	;
	v6169 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6170 = F_pstrdup(m, v6169)
	mBase = m.M
	v6171 = m.ExcPending
	if v6171 != 0 {
		goto L1
	} else {
		goto L1194
	}
L682:
	;
	v6166 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6167 = F_pstrdup(m, v6166)
	mBase = m.M
	v6168 = m.ExcPending
	if v6168 != 0 {
		goto L1
	} else {
		goto L1193
	}
L683:
	;
	v6163 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6164 = F_pstrdup(m, v6163)
	mBase = m.M
	v6165 = m.ExcPending
	if v6165 != 0 {
		goto L1
	} else {
		goto L1192
	}
L684:
	;
	v6160 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6161 = F_pstrdup(m, v6160)
	mBase = m.M
	v6162 = m.ExcPending
	if v6162 != 0 {
		goto L1
	} else {
		goto L1191
	}
L685:
	;
	v6157 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6158 = F_pstrdup(m, v6157)
	mBase = m.M
	v6159 = m.ExcPending
	if v6159 != 0 {
		goto L1
	} else {
		goto L1190
	}
L686:
	;
	v6154 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6155 = F_pstrdup(m, v6154)
	mBase = m.M
	v6156 = m.ExcPending
	if v6156 != 0 {
		goto L1
	} else {
		goto L1189
	}
L687:
	;
	v6151 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6152 = F_pstrdup(m, v6151)
	mBase = m.M
	v6153 = m.ExcPending
	if v6153 != 0 {
		goto L1
	} else {
		goto L1188
	}
L688:
	;
	v6148 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6149 = F_pstrdup(m, v6148)
	mBase = m.M
	v6150 = m.ExcPending
	if v6150 != 0 {
		goto L1
	} else {
		goto L1187
	}
L689:
	;
	v6145 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6146 = F_pstrdup(m, v6145)
	mBase = m.M
	v6147 = m.ExcPending
	if v6147 != 0 {
		goto L1
	} else {
		goto L1186
	}
L690:
	;
	v6142 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6143 = F_pstrdup(m, v6142)
	mBase = m.M
	v6144 = m.ExcPending
	if v6144 != 0 {
		goto L1
	} else {
		goto L1185
	}
L691:
	;
	v6139 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6140 = F_pstrdup(m, v6139)
	mBase = m.M
	v6141 = m.ExcPending
	if v6141 != 0 {
		goto L1
	} else {
		goto L1184
	}
L692:
	;
	v6136 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6137 = F_pstrdup(m, v6136)
	mBase = m.M
	v6138 = m.ExcPending
	if v6138 != 0 {
		goto L1
	} else {
		goto L1183
	}
L693:
	;
	v6133 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6134 = F_pstrdup(m, v6133)
	mBase = m.M
	v6135 = m.ExcPending
	if v6135 != 0 {
		goto L1
	} else {
		goto L1182
	}
L694:
	;
	v6130 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6131 = F_pstrdup(m, v6130)
	mBase = m.M
	v6132 = m.ExcPending
	if v6132 != 0 {
		goto L1
	} else {
		goto L1181
	}
L695:
	;
	v6129 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6189 = v6129
	goto L675
L696:
	;
	v6051 = int32(_a_F_BootstrapModeMain_92)
	v6053 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75])) = v6053 + int32(1)
	v6057 = m.G0
	v6059 = v6057 - int32(32)
	m.G0 = v6059
	v6063 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v6064 = m.ExcPending
	if v6064 != 0 {
		goto L1
	} else {
		goto L1169
	}
L697:
	;
	v5703 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5704 = int32(_a_F_BootstrapModeMain_92)
	v5706 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75])) = v5706 + int32(1)
	v5710 = m.G0
	v5712 = v5710 - int32(112)
	m.G0 = v5712
	v5716 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5717 = m.ExcPending
	if v5717 != 0 {
		goto L1
	} else {
		goto L1110
	}
L698:
	;
	v5697 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v5701 = F_strtox_2(m, v5697, int32(0), int32(10), int64(4294967295))
	mBase = m.M
	goto L1109
L699:
	;
	v6189 = int32(2)
	goto L675
L700:
	;
	v6189 = int32(3)
	goto L675
L701:
	;
	v4937 = int32(_a_F_BootstrapModeMain_93)
	v4939 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4941 = v4939 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76])) = v4941
	if int32(41) <= v4941 {
		goto L263
	} else {
		goto L989
	}
L702:
	;
	v4936 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v6189 = v4936
	goto L675
L703:
	;
	v6189 = int32(0)
	goto L675
L704:
	;
	v6189 = int32(1)
	goto L675
L705:
	;
	v4907 = F_palloc0(m, int32(40))
	mBase = m.M
	v4908 = m.ExcPending
	if v4908 != 0 {
		goto L1
	} else {
		goto L986
	}
L706:
	;
	v4898 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+96)) = v4898
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+108)) = v4898
	v4904 = F_list_make1_impl(m, int32(1), v3768+int32(96))
	mBase = m.M
	v4905 = m.ExcPending
	if v4905 != 0 {
		goto L1
	} else {
		goto L985
	}
L707:
	;
	v4894 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4895 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4896 = F_lappend(m, v4894, v4895)
	mBase = m.M
	v4897 = m.ExcPending
	if v4897 != 0 {
		goto L1
	} else {
		goto L984
	}
L708:
	;
	v4771 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4771 == int32(0) {
		goto L961
	} else {
		goto L962
	}
L709:
	;
	v4653 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4654 = m.ExcPending
	if v4654 != 0 {
		goto L1
	} else {
		goto L928
	}
L710:
	;
	v4528 = F_palloc0(m, int32(72))
	mBase = m.M
	v4529 = m.ExcPending
	if v4529 != 0 {
		goto L1
	} else {
		goto L906
	}
L711:
	;
	v4406 = F_palloc0(m, int32(72))
	mBase = m.M
	v4407 = m.ExcPending
	if v4407 != 0 {
		goto L1
	} else {
		goto L884
	}
L712:
	;
	v4311 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	v4313 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	if v4311 != v4313 {
		goto L265
	} else {
		goto L854
	}
L713:
	;
	v4278 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4278 == int32(0) {
		goto L844
	} else {
		goto L845
	}
L714:
	;
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4120 == int32(0) {
		goto L801
	} else {
		goto L802
	}
L715:
	;
	v4099 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4099
	v4102 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4102)
	mBase = m.M
	v4104 = m.ExcPending
	if v4104 != 0 {
		goto L1
	} else {
		goto L793
	}
L716:
	;
	v4046 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4046 == int32(0) {
		goto L779
	} else {
		goto L780
	}
L717:
	;
	v4004 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4004 == int32(0) {
		goto L766
	} else {
		goto L767
	}
L718:
	;
	v3796 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v3796 == int32(0) {
		goto L719
	} else {
		goto L720
	}
L719:
	;
	v3801 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v3806 = F_AllocSetContextCreateInternal(m, v3801, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v3807 = m.ExcPending
	if v3807 != 0 {
		goto L1
	} else {
		goto L722
	}
L720:
	;
	v3809 = v3796
	goto L721
L721:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v3809
	v3812 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v3814 = m.G0
	v3816 = v3814 - int32(48)
	m.G0 = v3816
	v3818 = F_strlen(m, v3812)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v3818) {
		goto L723
	} else {
		goto L724
	}
L722:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v3806
	v3809 = v3806
	goto L721
L723:
	;
	v3821 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3812)+63)) = uint8(v3821)
	goto L725
L724:
	;
	goto L725
L725:
	;
	v3824 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[80]))
	if v3824 == int32(0) {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v3828 = m.ExcPending
	if v3828 != 0 {
		goto L1
	} else {
		goto L729
	}
L727:
	;
	goto L728
L728:
	;
	v3830 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	if v3830 != 0 {
		goto L730
	} else {
		goto L731
	}
L729:
	;
	goto L728
L730:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v3833 = m.ExcPending
	if v3833 != 0 {
		goto L1
	} else {
		goto L733
	}
L731:
	;
	goto L732
L732:
	;
	v3836 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3837 = m.ExcPending
	if v3837 != 0 {
		goto L1
	} else {
		goto L734
	}
L733:
	;
	goto L732
L734:
	;
	if v3836 != 0 {
		goto L735
	} else {
		goto L736
	}
L735:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+36)) = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+32)) = v3812
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_96), v3816+int32(32))
	mBase = m.M
	v3845 = m.ExcPending
	if v3845 != 0 {
		goto L1
	} else {
		goto L738
	}
L736:
	;
	goto L737
L737:
	;
	v3854 = F_makeRangeVar(m, int32(0), v3812, int32(-1))
	mBase = m.M
	v3855 = m.ExcPending
	if v3855 != 0 {
		goto L1
	} else {
		goto L740
	}
L738:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(502), int32(_a_F_BootstrapModeMain_97))
	mBase = m.M
	v3850 = m.ExcPending
	if v3850 != 0 {
		goto L1
	} else {
		goto L739
	}
L739:
	;
	goto L737
L740:
	;
	v3857 = F_table_openrv(m, v3854, int32(0))
	mBase = m.M
	v3858 = m.ExcPending
	if v3858 != 0 {
		goto L1
	} else {
		goto L741
	}
L741:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81])) = v3857
	v3861 = *(*int32)(unsafe.Add(mBase, uint32(v3857)+48))
	v3862 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3861)+120)))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76])) = v3862
	if int32(0) < v3862 {
		goto L742
	} else {
		goto L743
	}
L742:
	;
	v3874 = int32(0)
	goto L745
L743:
	;
	goto L744
L744:
	;
	m.G0 = v3816 + int32(48)
	v3983 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v3983
	v3986 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v3986)
	mBase = m.M
	v3988 = m.ExcPending
	if v3988 != 0 {
		goto L1
	} else {
		goto L758
	}
L745:
	;
	v3895 = v3874 << (uint(int32(2)) % 32)
	v3898 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[82])))
	if v3898 == int32(0) {
		goto L747
	} else {
		goto L748
	}
L746:
	;
	goto L744
L747:
	;
	v3902 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[83]))
	v3904 = F_MemoryContextAllocZero(m, v3902, int32(100))
	mBase = m.M
	v3905 = m.ExcPending
	if v3905 != 0 {
		goto L1
	} else {
		goto L750
	}
L748:
	;
	v3907 = v3898
	goto L749
L749:
	;
	v3909 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v3910 = *(*int32)(unsafe.Add(mBase, uint32(v3909)+52))
	v3911 = *(*int32)(unsafe.Add(mBase, uint32(v3910)))
	v3915 = int32(100)
	base.MemoryCopy(m, v3907, v3910+v3911<<(uint(int32(3))%32)+v3874*v3915+int32(28), v3915)
	v3922 = *(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[82])))
	v3925 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v3926 = m.ExcPending
	if v3926 != 0 {
		goto L1
	} else {
		goto L751
	}
L750:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3895)+uint32(_c_F_BootstrapModeMain[82]))) = v3904
	v3907 = v3904
	goto L749
L751:
	;
	if v3925 != 0 {
		goto L752
	} else {
		goto L753
	}
L752:
	;
	v3927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3922)+72)))
	v3928 = int32(*(*int16)(unsafe.Add(mBase, uint32(v3922)+74)))
	v3929 = *(*int32)(unsafe.Add(mBase, uint32(v3922)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v3816+int32(16)))) = v3929
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+12)) = v3928
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+8)) = v3927
	*(*int32)(unsafe.Add(mBase, uint32(v3816)+4)) = v3922 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v3816))) = v3874
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_98), v3816)
	mBase = m.M
	v3939 = m.ExcPending
	if v3939 != 0 {
		goto L1
	} else {
		goto L755
	}
L753:
	;
	goto L754
L754:
	;
	v3948 = v3874 + int32(1)
	v3950 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	if v3948 < v3950 {
		v3874 = v3948
		goto L745
	} else {
		goto L757
	}
L755:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(519), int32(_a_F_BootstrapModeMain_97))
	mBase = m.M
	v3944 = m.ExcPending
	if v3944 != 0 {
		goto L1
	} else {
		goto L756
	}
L756:
	;
	goto L754
L757:
	;
	goto L746
L758:
	;
	v3990 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v3990 != 0 {
		goto L759
	} else {
		goto L760
	}
L759:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v3992 = m.ExcPending
	if v3992 != 0 {
		goto L1
	} else {
		goto L762
	}
L760:
	;
	goto L761
L761:
	;
	v3993 = int32(0)
	v3994 = F_isatty(m, v3993)
	mBase = m.M
	if v3994 == v3993 {
		v6189 = v3792
		goto L675
	} else {
		goto L763
	}
L762:
	;
	goto L761
L763:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4000 = m.ExcPending
	if v4000 != 0 {
		goto L1
	} else {
		goto L764
	}
L764:
	;
	v4001 = F_fflush(m, v3774)
	mBase = m.M
	v4002 = m.ExcPending
	if v4002 != 0 {
		goto L1
	} else {
		goto L765
	}
L765:
	;
	v6189 = v3792
	goto L675
L766:
	;
	v4009 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4014 = F_AllocSetContextCreateInternal(m, v4009, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4015 = m.ExcPending
	if v4015 != 0 {
		goto L1
	} else {
		goto L769
	}
L767:
	;
	v4017 = v4004
	goto L768
L768:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4017
	v4020 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	F_closerel(m, v4020)
	mBase = m.M
	v4022 = m.ExcPending
	if v4022 != 0 {
		goto L1
	} else {
		goto L770
	}
L769:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4014
	v4017 = v4014
	goto L768
L770:
	;
	v4025 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4025
	v4028 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4028)
	mBase = m.M
	v4030 = m.ExcPending
	if v4030 != 0 {
		goto L1
	} else {
		goto L771
	}
L771:
	;
	v4032 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4032 != 0 {
		goto L772
	} else {
		goto L773
	}
L772:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4034 = m.ExcPending
	if v4034 != 0 {
		goto L1
	} else {
		goto L775
	}
L773:
	;
	goto L774
L774:
	;
	v4035 = int32(0)
	v4036 = F_isatty(m, v4035)
	mBase = m.M
	if v4036 == v4035 {
		v6189 = v3792
		goto L675
	} else {
		goto L776
	}
L775:
	;
	goto L774
L776:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4042 = m.ExcPending
	if v4042 != 0 {
		goto L1
	} else {
		goto L777
	}
L777:
	;
	v4043 = F_fflush(m, v3774)
	mBase = m.M
	v4044 = m.ExcPending
	if v4044 != 0 {
		goto L1
	} else {
		goto L778
	}
L778:
	;
	v6189 = v3792
	goto L675
L779:
	;
	v4051 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4056 = F_AllocSetContextCreateInternal(m, v4051, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4057 = m.ExcPending
	if v4057 != 0 {
		goto L1
	} else {
		goto L782
	}
L780:
	;
	v4059 = v4046
	goto L781
L781:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4059
	v4063 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76])) = v4063
	v4067 = F_errstart(m, int32(11), v4063)
	mBase = m.M
	v4068 = m.ExcPending
	if v4068 != 0 {
		goto L1
	} else {
		goto L783
	}
L782:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4056
	v4059 = v4056
	goto L781
L783:
	;
	if v4067 == int32(0) {
		v6189 = v3792
		goto L675
	} else {
		goto L784
	}
L784:
	;
	v4073 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4076 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4079 = *(*int64)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	*(*int64)(unsafe.Add(mBase, uint32(v3768)+8)) = v4079
	if v4076 != 0 {
		goto L785
	} else {
		goto L786
	}
L785:
	;
	v4083 = int32(_a_F_BootstrapModeMain_100)
	goto L787
L786:
	;
	v4083 = int32(_a_F_BootstrapModeMain_1)
	goto L787
L787:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+4)) = v4083
	if v4073 != 0 {
		goto L788
	} else {
		goto L789
	}
L788:
	;
	v4087 = int32(_a_F_BootstrapModeMain_101)
	goto L790
L789:
	;
	v4087 = int32(_a_F_BootstrapModeMain_1)
	goto L790
L790:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768))) = v4087
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_102), v3768)
	mBase = m.M
	v4091 = m.ExcPending
	if v4091 != 0 {
		goto L1
	} else {
		goto L791
	}
L791:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(166), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4096 = m.ExcPending
	if v4096 != 0 {
		goto L1
	} else {
		goto L792
	}
L792:
	;
	v6189 = v3792
	goto L675
L793:
	;
	v4106 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4106 != 0 {
		goto L794
	} else {
		goto L795
	}
L794:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4108 = m.ExcPending
	if v4108 != 0 {
		goto L1
	} else {
		goto L797
	}
L795:
	;
	goto L796
L796:
	;
	v4109 = int32(0)
	v4110 = F_isatty(m, v4109)
	mBase = m.M
	if v4110 == v4109 {
		v6189 = v3792
		goto L675
	} else {
		goto L798
	}
L797:
	;
	goto L796
L798:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4116 = m.ExcPending
	if v4116 != 0 {
		goto L1
	} else {
		goto L799
	}
L799:
	;
	v4117 = F_fflush(m, v3774)
	mBase = m.M
	v4118 = m.ExcPending
	if v4118 != 0 {
		goto L1
	} else {
		goto L800
	}
L800:
	;
	v6189 = v3792
	goto L675
L801:
	;
	v4125 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4130 = F_AllocSetContextCreateInternal(m, v4125, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4131 = m.ExcPending
	if v4131 != 0 {
		goto L1
	} else {
		goto L804
	}
L802:
	;
	v4133 = v4120
	goto L803
L803:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4133
	v4137 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4139 = F_CreateTupleDesc(m, v4137, int32(_a_F_BootstrapModeMain_60))
	mBase = m.M
	v4140 = m.ExcPending
	if v4140 != 0 {
		goto L1
	} else {
		goto L805
	}
L804:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4130
	v4133 = v4130
	goto L803
L805:
	;
	v4143 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(24))))
	v4146 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	if v4146 != 0 {
		goto L807
	} else {
		goto L808
	}
L806:
	;
	v4257 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4257
	v4260 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4260)
	mBase = m.M
	v4262 = m.ExcPending
	if v4262 != 0 {
		goto L1
	} else {
		goto L836
	}
L807:
	;
	v4148 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	if v4148 != 0 {
		goto L810
	} else {
		goto L811
	}
L808:
	;
	goto L809
L809:
	;
	v4209 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(36))))
	if v4143 != 0 {
		goto L828
	} else {
		goto L829
	}
L810:
	;
	v4151 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4152 = m.ExcPending
	if v4152 != 0 {
		goto L1
	} else {
		goto L813
	}
L811:
	;
	goto L812
L812:
	;
	v4168 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(36))))
	if v4143 != 0 {
		goto L820
	} else {
		goto L821
	}
L813:
	;
	if v4151 != 0 {
		goto L814
	} else {
		goto L815
	}
L814:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_105), int32(0))
	mBase = m.M
	v4156 = m.ExcPending
	if v4156 != 0 {
		goto L1
	} else {
		goto L817
	}
L815:
	;
	goto L816
L816:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4164 = m.ExcPending
	if v4164 != 0 {
		goto L1
	} else {
		goto L819
	}
L817:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(202), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4161 = m.ExcPending
	if v4161 != 0 {
		goto L1
	} else {
		goto L818
	}
L818:
	;
	goto L816
L819:
	;
	goto L812
L820:
	;
	v4172 = int32(1664)
	goto L822
L821:
	;
	v4172 = int32(0)
	goto L822
L822:
	;
	v4175 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	v4176 = int32(0)
	v4179 = int32(112)
	v4182 = int32(1)
	v4189 = F_heap_create(m, v4168, int32(11), v4172, v4175, v4176, int32(2), v4139, int32(114), v4179, base.B2i32(v4143 != v4176), v4182, v4182, v3768+v4179, v3768+int32(124), v4182)
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L1
	} else {
		goto L823
	}
L823:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81])) = v4189
	v4194 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4195 = m.ExcPending
	if v4195 != 0 {
		goto L1
	} else {
		goto L824
	}
L824:
	;
	if v4194 == int32(0) {
		goto L806
	} else {
		goto L825
	}
L825:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_106), int32(0))
	mBase = m.M
	v4201 = m.ExcPending
	if v4201 != 0 {
		goto L1
	} else {
		goto L826
	}
L826:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(221), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4206 = m.ExcPending
	if v4206 != 0 {
		goto L1
	} else {
		goto L827
	}
L827:
	;
	goto L806
L828:
	;
	v4213 = int32(1664)
	goto L830
L829:
	;
	v4213 = int32(0)
	goto L830
L830:
	;
	v4216 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	v4219 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4220 = int32(0)
	v4227 = base.B2i32(v4143 != v4220)
	v4235 = F_heap_create_with_catalog(m, v4209, int32(11), v4213, v4216, v4219, v4220, int32(10), int32(2), v4139, v4220, int32(114), int32(112), v4227, v4227, v4220, int64(0), v4220, int32(1), v4220, v4220, v4220)
	mBase = m.M
	v4236 = m.ExcPending
	if v4236 != 0 {
		goto L1
	} else {
		goto L831
	}
L831:
	;
	v4239 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4240 = m.ExcPending
	if v4240 != 0 {
		goto L1
	} else {
		goto L832
	}
L832:
	;
	if v4239 == int32(0) {
		goto L806
	} else {
		goto L833
	}
L833:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+16)) = v4235
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_107), v3768+int32(16))
	mBase = m.M
	v4248 = m.ExcPending
	if v4248 != 0 {
		goto L1
	} else {
		goto L834
	}
L834:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(248), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4253 = m.ExcPending
	if v4253 != 0 {
		goto L1
	} else {
		goto L835
	}
L835:
	;
	goto L806
L836:
	;
	v4264 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4264 != 0 {
		goto L837
	} else {
		goto L838
	}
L837:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4266 = m.ExcPending
	if v4266 != 0 {
		goto L1
	} else {
		goto L840
	}
L838:
	;
	goto L839
L839:
	;
	v4267 = int32(0)
	v4268 = F_isatty(m, v4267)
	mBase = m.M
	if v4268 == v4267 {
		v6189 = v3792
		goto L675
	} else {
		goto L841
	}
L840:
	;
	goto L839
L841:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4274 = m.ExcPending
	if v4274 != 0 {
		goto L1
	} else {
		goto L842
	}
L842:
	;
	v4275 = F_fflush(m, v3774)
	mBase = m.M
	v4276 = m.ExcPending
	if v4276 != 0 {
		goto L1
	} else {
		goto L843
	}
L843:
	;
	v6189 = v3792
	goto L675
L844:
	;
	v4283 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4288 = F_AllocSetContextCreateInternal(m, v4283, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4289 = m.ExcPending
	if v4289 != 0 {
		goto L1
	} else {
		goto L847
	}
L845:
	;
	v4291 = v4278
	goto L846
L846:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4291
	v4296 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4297 = m.ExcPending
	if v4297 != 0 {
		goto L1
	} else {
		goto L848
	}
L847:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4288
	v4291 = v4288
	goto L846
L848:
	;
	if v4296 != 0 {
		goto L849
	} else {
		goto L850
	}
L849:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_108), int32(0))
	mBase = m.M
	v4301 = m.ExcPending
	if v4301 != 0 {
		goto L1
	} else {
		goto L852
	}
L850:
	;
	goto L851
L851:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75])) = int32(0)
	v6189 = v3792
	goto L675
L852:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(258), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4306 = m.ExcPending
	if v4306 != 0 {
		goto L1
	} else {
		goto L853
	}
L853:
	;
	goto L851
L854:
	;
	v4316 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	if v4316 == int32(0) {
		goto L264
	} else {
		goto L855
	}
L855:
	;
	v4319 = m.G0
	v4321 = v4319 - int32(16)
	m.G0 = v4321
	v4325 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4326 = m.ExcPending
	if v4326 != 0 {
		goto L1
	} else {
		goto L856
	}
L856:
	;
	if v4325 != 0 {
		goto L857
	} else {
		goto L858
	}
L857:
	;
	v4328 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v4321))) = v4328
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_109), v4321)
	mBase = m.M
	v4332 = m.ExcPending
	if v4332 != 0 {
		goto L1
	} else {
		goto L860
	}
L858:
	;
	goto L859
L859:
	;
	v4339 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4341 = F_CreateTupleDesc(m, v4339, int32(_a_F_BootstrapModeMain_60))
	mBase = m.M
	v4342 = m.ExcPending
	if v4342 != 0 {
		goto L1
	} else {
		goto L862
	}
L860:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(679), int32(_a_F_BootstrapModeMain_110))
	mBase = m.M
	v4337 = m.ExcPending
	if v4337 != 0 {
		goto L1
	} else {
		goto L861
	}
L861:
	;
	goto L859
L862:
	;
	v4345 = F_heap_form_tuple(m, v4341, int32(_a_F_BootstrapModeMain_111), int32(_a_F_BootstrapModeMain_112))
	mBase = m.M
	v4346 = m.ExcPending
	if v4346 != 0 {
		goto L1
	} else {
		goto L863
	}
L863:
	;
	F_pfree(m, v4341)
	mBase = m.M
	v4348 = m.ExcPending
	if v4348 != 0 {
		goto L1
	} else {
		goto L864
	}
L864:
	;
	v4350 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	F_simple_heap_insert(m, v4350, v4345)
	mBase = m.M
	v4352 = m.ExcPending
	if v4352 != 0 {
		goto L1
	} else {
		goto L865
	}
L865:
	;
	F_pfree(m, v4345)
	mBase = m.M
	v4354 = m.ExcPending
	if v4354 != 0 {
		goto L1
	} else {
		goto L866
	}
L866:
	;
	v4357 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4358 = m.ExcPending
	if v4358 != 0 {
		goto L1
	} else {
		goto L867
	}
L867:
	;
	if v4357 != 0 {
		goto L868
	} else {
		goto L869
	}
L868:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_113), int32(0))
	mBase = m.M
	v4362 = m.ExcPending
	if v4362 != 0 {
		goto L1
	} else {
		goto L871
	}
L869:
	;
	goto L870
L870:
	;
	v4369 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	v4370 = int32(0)
	if base.B2i32(v4369 <= v4370)|base.B2i32(v4369 == v4370) == v4370 {
		goto L873
	} else {
		goto L874
	}
L871:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(687), int32(_a_F_BootstrapModeMain_110))
	mBase = m.M
	v4367 = m.ExcPending
	if v4367 != 0 {
		goto L1
	} else {
		goto L872
	}
L872:
	;
	goto L870
L873:
	;
	base.MemoryFill(m, int32(_a_F_BootstrapModeMain_112), int32(0), v4369)
	goto L875
L874:
	;
	goto L875
L875:
	;
	m.G0 = v4321 + int32(16)
	v4385 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4385
	v4388 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4388)
	mBase = m.M
	v4390 = m.ExcPending
	if v4390 != 0 {
		goto L1
	} else {
		goto L876
	}
L876:
	;
	v4392 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4392 != 0 {
		goto L877
	} else {
		goto L878
	}
L877:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4394 = m.ExcPending
	if v4394 != 0 {
		goto L1
	} else {
		goto L880
	}
L878:
	;
	goto L879
L879:
	;
	v4395 = int32(0)
	v4396 = F_isatty(m, v4395)
	mBase = m.M
	if v4396 == v4395 {
		v6189 = v3792
		goto L675
	} else {
		goto L881
	}
L880:
	;
	goto L879
L881:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4402 = m.ExcPending
	if v4402 != 0 {
		goto L1
	} else {
		goto L882
	}
L882:
	;
	v4403 = F_fflush(m, v3774)
	mBase = m.M
	v4404 = m.ExcPending
	if v4404 != 0 {
		goto L1
	} else {
		goto L883
	}
L883:
	;
	v6189 = v3792
	goto L675
L884:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4406))) = int32(204)
	v4412 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4413 = m.ExcPending
	if v4413 != 0 {
		goto L1
	} else {
		goto L885
	}
L885:
	;
	if v4412 != 0 {
		goto L886
	} else {
		goto L887
	}
L886:
	;
	v4416 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+48)) = v4416
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_114), v3768+int32(48))
	mBase = m.M
	v4422 = m.ExcPending
	if v4422 != 0 {
		goto L1
	} else {
		goto L889
	}
L887:
	;
	goto L888
L888:
	;
	v4429 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4429 == int32(0) {
		goto L891
	} else {
		goto L892
	}
L889:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(279), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4427 = m.ExcPending
	if v4427 != 0 {
		goto L1
	} else {
		goto L890
	}
L890:
	;
	goto L888
L891:
	;
	v4434 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4439 = F_AllocSetContextCreateInternal(m, v4434, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4440 = m.ExcPending
	if v4440 != 0 {
		goto L1
	} else {
		goto L894
	}
L892:
	;
	v4442 = v4429
	goto L893
L893:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4442
	v4447 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+4)) = v4447
	v4452 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4454 = F_makeRangeVar(m, int32(0), v4452, int32(-1))
	mBase = m.M
	v4455 = m.ExcPending
	if v4455 != 0 {
		goto L1
	} else {
		goto L895
	}
L894:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4439
	v4442 = v4439
	goto L893
L895:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+8)) = v4454
	v4459 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4460 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+16)) = v4460
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+12)) = v4459
	v4465 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	*(*uint16)(unsafe.Add(mBase, uint32(v4406)+62)) = uint16(v4460)
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+20)) = v4465
	v4469 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+24)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+32)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+40)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+48)) = v4469
	*(*int64)(unsafe.Add(mBase, uint32(v4406)+53)) = v4469
	*(*int32)(unsafe.Add(mBase, uint32(v4406)+65)) = v4460
	*(*uint16)(unsafe.Add(mBase, uint32(v4406)+69)) = uint16(v4460)
	v4490 = F_RangeVarGetRelidExtended(m, v4454, v4460, v4460, v4460, v4460)
	mBase = m.M
	v4491 = m.ExcPending
	if v4491 != 0 {
		goto L1
	} else {
		goto L896
	}
L896:
	;
	v4494 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	v4495 = int32(0)
	F_DefineIndex(m, v3768+int32(112), v4460, v4490, v4406, v4494, v4495, v4495, int32(-1), v4495, v4495, v4495, int32(1), v4495)
	mBase = m.M
	v4504 = m.ExcPending
	if v4504 != 0 {
		goto L1
	} else {
		goto L897
	}
L897:
	;
	v4507 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4507
	v4510 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4510)
	mBase = m.M
	v4512 = m.ExcPending
	if v4512 != 0 {
		goto L1
	} else {
		goto L898
	}
L898:
	;
	v4514 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4514 != 0 {
		goto L899
	} else {
		goto L900
	}
L899:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4516 = m.ExcPending
	if v4516 != 0 {
		goto L1
	} else {
		goto L902
	}
L900:
	;
	goto L901
L901:
	;
	v4517 = int32(0)
	v4518 = F_isatty(m, v4517)
	mBase = m.M
	if v4518 == v4517 {
		v6189 = v3792
		goto L675
	} else {
		goto L903
	}
L902:
	;
	goto L901
L903:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4524 = m.ExcPending
	if v4524 != 0 {
		goto L1
	} else {
		goto L904
	}
L904:
	;
	v4525 = F_fflush(m, v3774)
	mBase = m.M
	v4526 = m.ExcPending
	if v4526 != 0 {
		goto L1
	} else {
		goto L905
	}
L905:
	;
	v6189 = v3792
	goto L675
L906:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4528))) = int32(204)
	v4534 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v4535 = m.ExcPending
	if v4535 != 0 {
		goto L1
	} else {
		goto L907
	}
L907:
	;
	if v4534 != 0 {
		goto L908
	} else {
		goto L909
	}
L908:
	;
	v4538 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+64)) = v4538
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_115), v3768-int32(-64))
	mBase = m.M
	v4544 = m.ExcPending
	if v4544 != 0 {
		goto L1
	} else {
		goto L911
	}
L909:
	;
	goto L910
L910:
	;
	v4551 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4551 == int32(0) {
		goto L913
	} else {
		goto L914
	}
L911:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(333), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4549 = m.ExcPending
	if v4549 != 0 {
		goto L1
	} else {
		goto L912
	}
L912:
	;
	goto L910
L913:
	;
	v4556 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4561 = F_AllocSetContextCreateInternal(m, v4556, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4562 = m.ExcPending
	if v4562 != 0 {
		goto L1
	} else {
		goto L916
	}
L914:
	;
	v4564 = v4551
	goto L915
L915:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4564
	v4569 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+4)) = v4569
	v4574 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(20))))
	v4576 = F_makeRangeVar(m, int32(0), v4574, int32(-1))
	mBase = m.M
	v4577 = m.ExcPending
	if v4577 != 0 {
		goto L1
	} else {
		goto L917
	}
L916:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4561
	v4564 = v4561
	goto L915
L917:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+8)) = v4576
	v4581 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4582 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+16)) = v4582
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+12)) = v4581
	v4587 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	v4588 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4528)+24)) = v4588
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+20)) = v4587
	*(*int64)(unsafe.Add(mBase, uint32(v4528)+32)) = v4588
	*(*int64)(unsafe.Add(mBase, uint32(v4528)+40)) = v4588
	*(*int64)(unsafe.Add(mBase, uint32(v4528)+48)) = v4588
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+56)) = v4582
	*(*int32)(unsafe.Add(mBase, uint32(v4528)+65)) = v4582
	*(*uint16)(unsafe.Add(mBase, uint32(v4528)+62)) = uint16(v4582)
	v4603 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v4528)+60)) = uint8(v4603)
	*(*uint16)(unsafe.Add(mBase, uint32(v4528)+69)) = uint16(v4582)
	v4614 = F_RangeVarGetRelidExtended(m, v4576, v4582, v4582, v4582, v4582)
	mBase = m.M
	v4615 = m.ExcPending
	if v4615 != 0 {
		goto L1
	} else {
		goto L918
	}
L918:
	;
	v4618 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(28))))
	v4619 = int32(0)
	F_DefineIndex(m, v3768+int32(112), v4582, v4614, v4528, v4618, v4619, v4619, int32(-1), v4619, v4619, v4619, int32(1), v4619)
	mBase = m.M
	v4628 = m.ExcPending
	if v4628 != 0 {
		goto L1
	} else {
		goto L919
	}
L919:
	;
	v4631 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4631
	v4634 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4634)
	mBase = m.M
	v4636 = m.ExcPending
	if v4636 != 0 {
		goto L1
	} else {
		goto L920
	}
L920:
	;
	v4638 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4638 != 0 {
		goto L921
	} else {
		goto L922
	}
L921:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4640 = m.ExcPending
	if v4640 != 0 {
		goto L1
	} else {
		goto L924
	}
L922:
	;
	goto L923
L923:
	;
	v4641 = int32(0)
	v4642 = F_isatty(m, v4641)
	mBase = m.M
	if v4642 == v4641 {
		v6189 = v3792
		goto L675
	} else {
		goto L925
	}
L924:
	;
	goto L923
L925:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4648 = m.ExcPending
	if v4648 != 0 {
		goto L1
	} else {
		goto L926
	}
L926:
	;
	v4649 = F_fflush(m, v3774)
	mBase = m.M
	v4650 = m.ExcPending
	if v4650 != 0 {
		goto L1
	} else {
		goto L927
	}
L927:
	;
	v6189 = v3792
	goto L675
L928:
	;
	if v4653 != 0 {
		goto L929
	} else {
		goto L930
	}
L929:
	;
	v4655 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+80)) = v4655
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_116), v3768+int32(80))
	mBase = m.M
	v4661 = m.ExcPending
	if v4661 != 0 {
		goto L1
	} else {
		goto L932
	}
L930:
	;
	goto L931
L931:
	;
	v4668 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	if v4668 == int32(0) {
		goto L934
	} else {
		goto L935
	}
L932:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(384), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v4666 = m.ExcPending
	if v4666 != 0 {
		goto L1
	} else {
		goto L933
	}
L933:
	;
	goto L931
L934:
	;
	v4673 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4678 = F_AllocSetContextCreateInternal(m, v4673, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4679 = m.ExcPending
	if v4679 != 0 {
		goto L1
	} else {
		goto L937
	}
L935:
	;
	v4681 = v4668
	goto L936
L936:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4681
	v4684 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4687 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4690 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(8))))
	v4691 = m.G0
	v4693 = v4691 - int32(32)
	m.G0 = v4693
	v4697 = F_makeRangeVar(m, int32(0), v4684, int32(-1))
	mBase = m.M
	v4698 = m.ExcPending
	if v4698 != 0 {
		goto L1
	} else {
		goto L940
	}
L937:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4678
	v4681 = v4678
	goto L936
L938:
	;
	v4721 = int32(0)
	v4723 = F_create_toast_table(m, v4700, v4687, v4690, int64(0), int32(8), v4721, v4721)
	mBase = m.M
	v4724 = m.ExcPending
	if v4724 != 0 {
		goto L1
	} else {
		goto L945
	}
L939:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4709 = m.ExcPending
	if v4709 != 0 {
		goto L1
	} else {
		goto L942
	}
L940:
	;
	v4700 = F_table_openrv(m, v4697, int32(8))
	mBase = m.M
	v4701 = m.ExcPending
	if v4701 != 0 {
		goto L1
	} else {
		goto L941
	}
L941:
	;
	v4702 = *(*int32)(unsafe.Add(mBase, uint32(v4700)+48))
	v4703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4702)+119)))
	switch v4703 - int32(109) {
	case 0, 5:
		goto L938
	default:
		goto L939
	}
L942:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4693))) = v4684
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_117), v4693)
	mBase = m.M
	v4713 = m.ExcPending
	if v4713 != 0 {
		goto L1
	} else {
		goto L943
	}
L943:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_118), int32(108), int32(_a_F_BootstrapModeMain_119))
	mBase = m.M
	v4718 = m.ExcPending
	if v4718 != 0 {
		goto L1
	} else {
		goto L944
	}
L944:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L945:
	;
	if v4723 == int32(0) {
		goto L946
	} else {
		goto L947
	}
L946:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4730 = m.ExcPending
	if v4730 != 0 {
		goto L1
	} else {
		goto L949
	}
L947:
	;
	goto L948
L948:
	;
	F_relation_close(m, v4700, int32(0))
	mBase = m.M
	v4744 = m.ExcPending
	if v4744 != 0 {
		goto L1
	} else {
		goto L952
	}
L949:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4693)+16)) = v4684
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_120), v4693+int32(16))
	mBase = m.M
	v4736 = m.ExcPending
	if v4736 != 0 {
		goto L1
	} else {
		goto L950
	}
L950:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_118), int32(114), int32(_a_F_BootstrapModeMain_119))
	mBase = m.M
	v4741 = m.ExcPending
	if v4741 != 0 {
		goto L1
	} else {
		goto L951
	}
L951:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L952:
	;
	m.G0 = v4693 + int32(32)
	v4750 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4750
	v4753 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4753)
	mBase = m.M
	v4755 = m.ExcPending
	if v4755 != 0 {
		goto L1
	} else {
		goto L953
	}
L953:
	;
	v4757 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4757 != 0 {
		goto L954
	} else {
		goto L955
	}
L954:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4759 = m.ExcPending
	if v4759 != 0 {
		goto L1
	} else {
		goto L957
	}
L955:
	;
	goto L956
L956:
	;
	v4760 = int32(0)
	v4761 = F_isatty(m, v4760)
	mBase = m.M
	if v4761 == v4760 {
		v6189 = v3792
		goto L675
	} else {
		goto L958
	}
L957:
	;
	goto L956
L958:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4767 = m.ExcPending
	if v4767 != 0 {
		goto L1
	} else {
		goto L959
	}
L959:
	;
	v4768 = F_fflush(m, v3774)
	mBase = m.M
	v4769 = m.ExcPending
	if v4769 != 0 {
		goto L1
	} else {
		goto L960
	}
L960:
	;
	v6189 = v3792
	goto L675
L961:
	;
	v4776 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	v4781 = F_AllocSetContextCreateInternal(m, v4776, int32(_a_F_BootstrapModeMain_94), int32(0), int32(_a_F_BootstrapModeMain_25), int32(_a_F_BootstrapModeMain_95))
	mBase = m.M
	v4782 = m.ExcPending
	if v4782 != 0 {
		goto L1
	} else {
		goto L964
	}
L962:
	;
	v4784 = v4771
	goto L963
L963:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4784
	v4788 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[85]))
	if v4788 != 0 {
		goto L965
	} else {
		goto L966
	}
L964:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77])) = v4781
	v4784 = v4781
	goto L963
L965:
	;
	v4793 = v4788
	goto L968
L966:
	;
	goto L967
L967:
	;
	v4872 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[78]))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[79])) = v4872
	v4875 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[77]))
	F_MemoryContextReset(m, v4875)
	mBase = m.M
	v4877 = m.ExcPending
	if v4877 != 0 {
		goto L1
	} else {
		goto L976
	}
L968:
	;
	v4815 = *(*int32)(unsafe.Add(mBase, uint32(v4793)))
	v4817 = F_table_open(m, v4815, int32(0))
	mBase = m.M
	v4818 = m.ExcPending
	if v4818 != 0 {
		goto L1
	} else {
		goto L970
	}
L969:
	;
	goto L967
L970:
	;
	v4820 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[85]))
	v4821 = *(*int32)(unsafe.Add(mBase, uint32(v4820)+4))
	v4823 = F_index_open(m, v4821, int32(0))
	mBase = m.M
	v4824 = m.ExcPending
	if v4824 != 0 {
		goto L1
	} else {
		goto L971
	}
L971:
	;
	v4826 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[85]))
	v4827 = *(*int32)(unsafe.Add(mBase, uint32(v4826)+8))
	v4828 = int32(0)
	F_index_build(m, v4817, v4823, v4827, v4828, v4828, v4828)
	mBase = m.M
	v4832 = m.ExcPending
	if v4832 != 0 {
		goto L1
	} else {
		goto L972
	}
L972:
	;
	F_relation_close(m, v4823, int32(0))
	mBase = m.M
	v4835 = m.ExcPending
	if v4835 != 0 {
		goto L1
	} else {
		goto L973
	}
L973:
	;
	F_relation_close(m, v4817, int32(0))
	mBase = m.M
	v4838 = m.ExcPending
	if v4838 != 0 {
		goto L1
	} else {
		goto L974
	}
L974:
	;
	v4839 = int32(_a_F_BootstrapModeMain_121)
	v4841 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[85]))
	v4842 = *(*int32)(unsafe.Add(mBase, uint32(v4841)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[85])) = v4842
	if v4842 != 0 {
		v4793 = v4842
		goto L968
	} else {
		goto L975
	}
L975:
	;
	goto L969
L976:
	;
	v4879 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[84]))
	if v4879 != 0 {
		goto L977
	} else {
		goto L978
	}
L977:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v4881 = m.ExcPending
	if v4881 != 0 {
		goto L1
	} else {
		goto L980
	}
L978:
	;
	goto L979
L979:
	;
	v4882 = int32(0)
	v4883 = F_isatty(m, v4882)
	mBase = m.M
	if v4883 == v4882 {
		v6189 = v3792
		goto L675
	} else {
		goto L981
	}
L980:
	;
	goto L979
L981:
	;
	F_pg_printf(m, int32(_a_F_BootstrapModeMain_99), int32(0))
	mBase = m.M
	v4889 = m.ExcPending
	if v4889 != 0 {
		goto L1
	} else {
		goto L982
	}
L982:
	;
	v4890 = F_fflush(m, v3774)
	mBase = m.M
	v4891 = m.ExcPending
	if v4891 != 0 {
		goto L1
	} else {
		goto L983
	}
L983:
	;
	v6189 = v3792
	goto L675
L984:
	;
	v6189 = v4896
	goto L675
L985:
	;
	v6189 = v4904
	goto L675
L986:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4907))) = int32(92)
	v4913 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v4907)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v4907)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4907)+4)) = v4913
	v4919 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4920 = F_makeString(m, v4919)
	mBase = m.M
	v4921 = m.ExcPending
	if v4921 != 0 {
		goto L1
	} else {
		goto L987
	}
L987:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+100)) = v4920
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+104)) = v4920
	v4927 = F_list_make1_impl(m, int32(1), v3768+int32(100))
	mBase = m.M
	v4928 = m.ExcPending
	if v4928 != 0 {
		goto L1
	} else {
		goto L988
	}
L988:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4907)+36)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v4907)+28)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4907)+20)) = v4927
	v6189 = v4907
	goto L675
L989:
	;
	v4947 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(12))))
	v4950 = *(*int32)(unsafe.Add(mBase, uint32(v3765-int32(4))))
	v4951 = *(*int32)(unsafe.Add(mBase, uint32(v3765)))
	v4952 = int32(0)
	v4953 = m.G0
	v4955 = v4953 - int32(48)
	m.G0 = v4955
	v4958 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	if v4958 != 0 {
		goto L990
	} else {
		goto L991
	}
L990:
	;
	v4961 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v4962 = m.ExcPending
	if v4962 != 0 {
		goto L1
	} else {
		goto L993
	}
L991:
	;
	goto L992
L992:
	;
	v4976 = v4939 << (uint(int32(2)) % 32)
	v4979 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	if v4979 == int32(0) {
		goto L1000
	} else {
		goto L1001
	}
L993:
	;
	if v4961 != 0 {
		goto L994
	} else {
		goto L995
	}
L994:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_122), int32(0))
	mBase = m.M
	v4966 = m.ExcPending
	if v4966 != 0 {
		goto L1
	} else {
		goto L997
	}
L995:
	;
	goto L996
L996:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v4974 = m.ExcPending
	if v4974 != 0 {
		goto L1
	} else {
		goto L999
	}
L997:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(572), int32(_a_F_BootstrapModeMain_123))
	mBase = m.M
	v4971 = m.ExcPending
	if v4971 != 0 {
		goto L1
	} else {
		goto L998
	}
L998:
	;
	goto L996
L999:
	;
	goto L992
L1000:
	;
	v4983 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[83]))
	v4985 = F_MemoryContextAllocZero(m, v4983, int32(100))
	mBase = m.M
	v4986 = m.ExcPending
	if v4986 != 0 {
		goto L1
	} else {
		goto L1003
	}
L1001:
	;
	v4988 = v4979
	goto L1002
L1002:
	;
	v4989 = int32(0)
	base.MemoryFill(m, v4988, v4989, int32(100))
	v4992 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v4996 = F_strncpy(m, v4992+int32(4), v4947, int32(64))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v4996)+63)) = uint8(v4989)
	goto L1004
L1003:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82]))) = v4985
	v4988 = v4985
	goto L1002
L1004:
	;
	v5001 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5002 = m.ExcPending
	if v5002 != 0 {
		goto L1
	} else {
		goto L1005
	}
L1005:
	;
	if v5001 != 0 {
		goto L1006
	} else {
		goto L1007
	}
L1006:
	;
	v5003 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	*(*int32)(unsafe.Add(mBase, uint32(v4955)+36)) = v4950
	*(*int32)(unsafe.Add(mBase, uint32(v4955)+32)) = v5003 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_124), v4955+int32(32))
	mBase = m.M
	v5012 = m.ExcPending
	if v5012 != 0 {
		goto L1
	} else {
		goto L1009
	}
L1007:
	;
	goto L1008
L1008:
	;
	v5019 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5021 = v4939 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v5019)+74)) = uint16(v5021)
	v5024 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[80]))
	if v5024 == int32(0) {
		goto L1015
	} else {
		goto L1016
	}
L1009:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(581), int32(_a_F_BootstrapModeMain_123))
	mBase = m.M
	v5017 = m.ExcPending
	if v5017 != 0 {
		goto L1
	} else {
		goto L1010
	}
L1010:
	;
	goto L1008
L1011:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v5529)+80)) = uint16(v5531)
	v5552 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5553 = *(*int32)(unsafe.Add(mBase, uint32(v5552)+96))
	if v5553 != 0 {
		goto L1093
	} else {
		goto L1094
	}
L1012:
	;
	v5529 = v5502
	v5531 = int32(1)
	goto L1011
L1013:
	;
	v5466 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[86])) = v5444
	v5469 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5470 = *(*int32)(unsafe.Add(mBase, uint32(v5444)))
	*(*int32)(unsafe.Add(mBase, uint32(v5469)+68)) = v5470
	v5472 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5473 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5444)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v5472)+72)) = uint16(v5473)
	v5475 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5444)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5475)+82)) = uint8(v5476)
	v5478 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5444)+132)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5478)+83)) = uint8(v5479)
	v5481 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5444)+133)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5481)+84)) = uint8(v5482)
	v5484 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5484)+85)) = uint8(v5466)
	v5487 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5488 = *(*int32)(unsafe.Add(mBase, uint32(v5444)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v5487)+96)) = v5488
	v5490 = *(*int32)(unsafe.Add(mBase, uint32(v5444)+96))
	if v5490 == v5466 {
		goto L1090
	} else {
		goto L1091
	}
L1014:
	;
	v5408 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5410 = v5036 * int32(92)
	v5411 = *(*int32)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[87])))
	*(*int32)(unsafe.Add(mBase, uint32(v5408)+68)) = v5411
	v5413 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[88]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v5413)+72)) = uint16(v5414)
	v5416 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[89]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5416)+82)) = uint8(v5417)
	v5419 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[90]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5419)+83)) = uint8(v5420)
	v5422 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[91]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v5422)+84)) = uint8(v5423)
	v5425 = int32(0)
	v5426 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5426)+85)) = uint8(v5425)
	v5429 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5430 = *(*int32)(unsafe.Add(mBase, uint32(v5410)+uint32(_c_F_BootstrapModeMain[92])))
	*(*int32)(unsafe.Add(mBase, uint32(v5429)+96)) = v5430
	v5432 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	if int32(1)<<(uint(v5036)%32)&int32(_a_F_BootstrapModeMain_125) != 0 {
		v5529 = v5432
		v5531 = v5425
		goto L1011
	} else {
		goto L1088
	}
L1015:
	;
	v5036 = v4952
	goto L1018
L1016:
	;
	v5139 = v5024
	v5140 = v4952
	goto L1017
L1017:
	;
	v5157 = *(*int32)(unsafe.Add(mBase, uint32(v5139)+4))
	if int32(0) < v5157 {
		goto L1043
	} else {
		goto L1044
	}
L1018:
	;
	v5056 = v5036*int32(92) + int32(_a_F_BootstrapModeMain_126)
	goto L1022
L1019:
	;
	v5139 = v5128
	v5140 = v5126
	goto L1017
L1020:
	;
	if v5094-v5095 == int32(0) {
		goto L1014
	} else {
		goto L1033
	}
L1022:
	;
	goto L1023
L1023:
	;
	v5063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4950))))
	if v5063 != 0 {
		goto L1024
	} else {
		goto L1025
	}
L1024:
	;
	v5064 = v4950
	v5065 = v5056
	v5066 = int32(64)
	v5067 = v5063
	goto L1028
L1025:
	;
	v5090 = v5056
	v5094 = int32(0)
	goto L1026
L1026:
	;
	v5095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5090))))
	goto L1020
L1027:
	;
	v5090 = v5085
	v5094 = v5087
	goto L1026
L1028:
	;
	v5069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5065))))
	if base.B2i32(v5067 != v5069)|base.B2i32(v5069 == int32(0)) != 0 {
		v5085 = v5065
		v5087 = v5067
		goto L1027
	} else {
		goto L1030
	}
L1029:
	;
	v5085 = v5079
	v5087 = int32(0)
	goto L1027
L1030:
	;
	v5075 = v5066 - int32(1)
	if v5075 == int32(0) {
		v5085 = v5065
		v5087 = v5067
		goto L1027
	} else {
		goto L1031
	}
L1031:
	;
	v5078 = int32(1)
	v5079 = v5065 + v5078
	v5080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5064)+1)))
	if v5080 != 0 {
		v5064 = v5064 + v5078
		v5065 = v5079
		v5066 = v5075
		v5067 = v5080
		goto L1028
	} else {
		goto L1032
	}
L1032:
	;
	goto L1029
L1033:
	;
	v5106 = v5036 + int32(1)
	if v5106 != int32(23) {
		v5036 = v5106
		goto L1018
	} else {
		goto L1034
	}
L1034:
	;
	v5111 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5112 = m.ExcPending
	if v5112 != 0 {
		goto L1
	} else {
		goto L1035
	}
L1035:
	;
	if v5111 != 0 {
		goto L1036
	} else {
		goto L1037
	}
L1036:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4955)+16)) = v4950
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_127), v4955+int32(16))
	mBase = m.M
	v5118 = m.ExcPending
	if v5118 != 0 {
		goto L1
	} else {
		goto L1039
	}
L1037:
	;
	goto L1038
L1038:
	;
	F_populate_typ_list(m)
	mBase = m.M
	v5125 = m.ExcPending
	if v5125 != 0 {
		goto L1
	} else {
		goto L1041
	}
L1039:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(990), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5123 = m.ExcPending
	if v5123 != 0 {
		goto L1
	} else {
		goto L1040
	}
L1040:
	;
	goto L1038
L1041:
	;
	v5126 = int32(0)
	v5128 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[80]))
	if v5128 == v5126 {
		v5036 = v5126
		goto L1018
	} else {
		goto L1042
	}
L1042:
	;
	goto L1019
L1043:
	;
	v5160 = *(*int32)(unsafe.Add(mBase, uint32(v5139)+12))
	v5170 = v5140
	goto L1046
L1044:
	;
	goto L1045
L1045:
	;
	F_list_free_deep(m, v5139)
	mBase = m.M
	v5271 = m.ExcPending
	if v5271 != 0 {
		goto L1
	} else {
		goto L1063
	}
L1046:
	;
	v5190 = *(*int32)(unsafe.Add(mBase, uint32(v5160+v5170<<(uint(int32(2))%32))))
	v5192 = v5190 + int32(8)
	goto L1050
L1047:
	;
	goto L1045
L1048:
	;
	if v5230-v5231 == int32(0) {
		v5444 = v5190
		goto L1013
	} else {
		goto L1061
	}
L1050:
	;
	goto L1051
L1051:
	;
	v5199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5192))))
	if v5199 != 0 {
		goto L1052
	} else {
		goto L1053
	}
L1052:
	;
	v5200 = v5192
	v5201 = v4950
	v5202 = int32(64)
	v5203 = v5199
	goto L1056
L1053:
	;
	v5226 = v4950
	v5230 = int32(0)
	goto L1054
L1054:
	;
	v5231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5226))))
	goto L1048
L1055:
	;
	v5226 = v5221
	v5230 = v5223
	goto L1054
L1056:
	;
	v5205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5201))))
	if base.B2i32(v5203 != v5205)|base.B2i32(v5205 == int32(0)) != 0 {
		v5221 = v5201
		v5223 = v5203
		goto L1055
	} else {
		goto L1058
	}
L1057:
	;
	v5221 = v5215
	v5223 = int32(0)
	goto L1055
L1058:
	;
	v5211 = v5202 - int32(1)
	if v5211 == int32(0) {
		v5221 = v5201
		v5223 = v5203
		goto L1055
	} else {
		goto L1059
	}
L1059:
	;
	v5214 = int32(1)
	v5215 = v5201 + v5214
	v5216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5200)+1)))
	if v5216 != 0 {
		v5200 = v5200 + v5214
		v5201 = v5215
		v5202 = v5211
		v5203 = v5216
		goto L1056
	} else {
		goto L1060
	}
L1060:
	;
	goto L1057
L1061:
	;
	v5242 = v5170 + int32(1)
	if v5242 != v5157 {
		v5170 = v5242
		goto L1046
	} else {
		goto L1062
	}
L1062:
	;
	goto L1047
L1063:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[80])) = int32(0)
	F_populate_typ_list(m)
	mBase = m.M
	v5276 = m.ExcPending
	if v5276 != 0 {
		goto L1
	} else {
		goto L1064
	}
L1064:
	;
	v5278 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[80]))
	if v5278 == int32(0) {
		goto L1065
	} else {
		goto L1066
	}
L1065:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5398 = m.ExcPending
	if v5398 != 0 {
		goto L1
	} else {
		goto L1085
	}
L1066:
	;
	v5281 = *(*int32)(unsafe.Add(mBase, uint32(v5278)+4))
	if v5281 <= int32(0) {
		goto L1065
	} else {
		goto L1067
	}
L1067:
	;
	v5284 = *(*int32)(unsafe.Add(mBase, uint32(v5278)+12))
	v5295 = int32(0)
	goto L1068
L1068:
	;
	v5315 = *(*int32)(unsafe.Add(mBase, uint32(v5284+v5295<<(uint(int32(2))%32))))
	v5317 = v5315 + int32(8)
	goto L1072
L1069:
	;
	goto L1065
L1070:
	;
	if v5355-v5356 == int32(0) {
		v5444 = v5315
		goto L1013
	} else {
		goto L1083
	}
L1072:
	;
	goto L1073
L1073:
	;
	v5324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5317))))
	if v5324 != 0 {
		goto L1074
	} else {
		goto L1075
	}
L1074:
	;
	v5325 = v5317
	v5326 = v4950
	v5327 = int32(64)
	v5328 = v5324
	goto L1078
L1075:
	;
	v5351 = v4950
	v5355 = int32(0)
	goto L1076
L1076:
	;
	v5356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5351))))
	goto L1070
L1077:
	;
	v5351 = v5346
	v5355 = v5348
	goto L1076
L1078:
	;
	v5330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5326))))
	if base.B2i32(v5328 != v5330)|base.B2i32(v5330 == int32(0)) != 0 {
		v5346 = v5326
		v5348 = v5328
		goto L1077
	} else {
		goto L1080
	}
L1079:
	;
	v5346 = v5340
	v5348 = int32(0)
	goto L1077
L1080:
	;
	v5336 = v5327 - int32(1)
	if v5336 == int32(0) {
		v5346 = v5326
		v5348 = v5328
		goto L1077
	} else {
		goto L1081
	}
L1081:
	;
	v5339 = int32(1)
	v5340 = v5326 + v5339
	v5341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5325)+1)))
	if v5341 != 0 {
		v5325 = v5325 + v5339
		v5326 = v5340
		v5327 = v5336
		v5328 = v5341
		goto L1078
	} else {
		goto L1082
	}
L1082:
	;
	goto L1079
L1083:
	;
	v5367 = v5295 + int32(1)
	if v5281 != v5367 {
		v5295 = v5367
		goto L1068
	} else {
		goto L1084
	}
L1084:
	;
	goto L1069
L1085:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v4955))) = v4950
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_129), v4955)
	mBase = m.M
	v5402 = m.ExcPending
	if v5402 != 0 {
		goto L1
	} else {
		goto L1086
	}
L1086:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(994), int32(_a_F_BootstrapModeMain_128))
	mBase = m.M
	v5407 = m.ExcPending
	if v5407 != 0 {
		goto L1
	} else {
		goto L1087
	}
L1087:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1088:
	;
	v5437 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5432)+72)))
	if int32(0) <= v5437 {
		v5529 = v5432
		v5531 = v5425
		goto L1011
	} else {
		goto L1089
	}
L1089:
	;
	v5502 = v5432
	goto L1012
L1090:
	;
	v5497 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5529 = v5497
	v5531 = v5466
	goto L1011
L1091:
	;
	v5493 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5444)+80)))
	if int32(0) <= v5493 {
		goto L1090
	} else {
		goto L1092
	}
L1092:
	;
	v5496 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5502 = v5496
	goto L1012
L1093:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5552)+96)) = int32(950)
	v5556 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	v5557 = v5556
	goto L1095
L1094:
	;
	v5557 = v5552
	goto L1095
L1095:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5557)+76)) = int32(-1)
	v5560 = int32(1)
	v5561 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	*(*uint8)(unsafe.Add(mBase, uint32(v5561)+92)) = uint8(v5560)
	v5564 = *(*int32)(unsafe.Add(mBase, uint32(v4976)+uint32(_c_F_BootstrapModeMain[82])))
	switch v4951 - int32(2) {
	case 0:
		goto L1099
	case 1:
		v5645 = v5560
		goto L1097
	default:
		goto L1098
	}
L1096:
	;
	m.G0 = v4955 + int32(48)
	v6189 = v3792
	goto L675
L1097:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v5564)+86)) = uint8(v5645)
	goto L1096
L1098:
	;
	v5568 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5564)+72)))
	if v5568 <= int32(0) {
		goto L1096
	} else {
		goto L1100
	}
L1099:
	;
	v5645 = int32(0)
	goto L1097
L1100:
	;
	v5571 = int32(0)
	if v4939 <= v5571 {
		v5621 = v5571
		goto L1101
	} else {
		goto L1102
	}
L1101:
	;
	if v4939 != v5621 {
		goto L1096
	} else {
		goto L1108
	}
L1102:
	;
	v5583 = v5571
	goto L1103
L1103:
	;
	v5602 = *(*int32)(unsafe.Add(mBase, uint32(v5583<<(uint(int32(2))%32))+uint32(_c_F_BootstrapModeMain[82])))
	v5603 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5602)+72)))
	if v5603 <= int32(0) {
		v5621 = v5583
		goto L1101
	} else {
		goto L1105
	}
L1104:
	;
	v5645 = v5560
	goto L1097
L1105:
	;
	v5606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5602)+86)))
	if v5606 != int32(1) {
		v5621 = v5583
		goto L1101
	} else {
		goto L1106
	}
L1106:
	;
	v5610 = v5583 + int32(1)
	if v5610 != v4939 {
		v5583 = v5610
		goto L1103
	} else {
		goto L1107
	}
L1107:
	;
	goto L1104
L1108:
	;
	v5645 = v5560
	goto L1097
L1109:
	;
	v6189 = base.I32_wrap_i64(v5701)
	goto L675
L1110:
	;
	if v5716 != 0 {
		goto L1111
	} else {
		goto L1112
	}
L1111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5712)+36)) = v5703
	*(*int32)(unsafe.Add(mBase, uint32(v5712)+32)) = v5706
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_130), v5712+int32(32))
	mBase = m.M
	v5724 = m.ExcPending
	if v5724 != 0 {
		goto L1
	} else {
		goto L1114
	}
L1112:
	;
	goto L1113
L1113:
	;
	v5731 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v5732 = *(*int32)(unsafe.Add(mBase, uint32(v5731)+52))
	v5733 = *(*int32)(unsafe.Add(mBase, uint32(v5732)))
	v5739 = v5732 + v5733<<(uint(int32(3))%32) + v5706*int32(100)
	v5740 = *(*int32)(unsafe.Add(mBase, uint32(v5739)+96))
	F_boot_get_type_io_data(m, v5740, v5712+int32(74), v5712+int32(73), v5712+int32(72), v5712+int32(71), v5712-int32(-64), v5712+int32(60), v5712+int32(56), v5712+int32(52))
	mBase = m.M
	v5758 = m.ExcPending
	if v5758 != 0 {
		goto L1
	} else {
		goto L1116
	}
L1114:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(716), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v5729 = m.ExcPending
	if v5729 != 0 {
		goto L1
	} else {
		goto L1115
	}
L1115:
	;
	goto L1113
L1116:
	;
	if v5740 == int32(194) {
		goto L1122
	} else {
		goto L1123
	}
L1117:
	;
	v6189 = v3792
	goto L675
L1118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6041 = m.ExcPending
	if v6041 != 0 {
		goto L1
	} else {
		goto L1166
	}
L1119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6028 = m.ExcPending
	if v6028 != 0 {
		goto L1
	} else {
		goto L1163
	}
L1120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6015 = m.ExcPending
	if v6015 != 0 {
		goto L1
	} else {
		goto L1160
	}
L1121:
	;
	v5990 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v5991 = m.ExcPending
	if v5991 != 0 {
		goto L1
	} else {
		goto L1153
	}
L1122:
	;
	if v5706 != int32(23) {
		goto L1125
	} else {
		goto L1126
	}
L1123:
	;
	goto L1124
L1124:
	;
	v5956 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+60))
	v5957 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+64))
	v5959 = F_OidInputFunctionCall(m, v5956, v5703, v5957, int32(-1))
	mBase = m.M
	v5960 = m.ExcPending
	if v5960 != 0 {
		goto L1
	} else {
		goto L1152
	}
L1125:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v5932 = m.ExcPending
	if v5932 != 0 {
		goto L1
	} else {
		goto L1149
	}
L1126:
	;
	v5764 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v5765 = *(*int32)(unsafe.Add(mBase, uint32(v5764)+56))
	if v5765 != int32(1255) {
		goto L1125
	} else {
		goto L1127
	}
L1127:
	;
	v5769 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[55])))
	if v5769 == int32(1) {
		goto L1120
	} else {
		goto L1128
	}
L1128:
	;
	v5773 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[93])))
	if v5773 == int32(1) {
		goto L1119
	} else {
		goto L1129
	}
L1129:
	;
	v5776 = int32(0)
	v5778 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[94]))
	v5780 = int32(*(*int16)(unsafe.Add(mBase, _c_F_BootstrapModeMain[95])))
	v5785 = F_OidFunctionCall3Coll(m, int32(750), base.I64_extend_i32_u(v5703), int64(2275), int64(-1))
	mBase = m.M
	v5786 = m.ExcPending
	if v5786 != 0 {
		goto L1
	} else {
		goto L1130
	}
L1130:
	;
	v5788 = F_pg_detoast_datum(m, base.I32_wrap_i64(v5785))
	mBase = m.M
	v5789 = m.ExcPending
	if v5789 != 0 {
		goto L1
	} else {
		goto L1131
	}
L1131:
	;
	F_deconstruct_array_builtin(m, v5788, int32(2275), v5712+int32(108), v5712+int32(104), v5712+int32(100))
	mBase = m.M
	v5798 = m.ExcPending
	if v5798 != 0 {
		goto L1
	} else {
		goto L1132
	}
L1132:
	;
	v5799 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+100))
	if v5780 < v5799 {
		goto L1118
	} else {
		goto L1133
	}
L1133:
	;
	if int32(0) < v5799 {
		goto L1134
	} else {
		goto L1135
	}
L1134:
	;
	v5807 = int32(0)
	v5812 = v5776
	v5814 = v5799
	goto L1137
L1135:
	;
	v5893 = v5776
	goto L1136
L1136:
	;
	v5914 = F_nodeToString(m, v5893)
	mBase = m.M
	v5915 = m.ExcPending
	if v5915 != 0 {
		goto L1
	} else {
		goto L1147
	}
L1137:
	;
	v5833 = int32(2)
	v5839 = *(*int32)(unsafe.Add(mBase, uint32(v5778+int32(24)+(v5780-v5814)<<(uint(v5833)%32)+v5807<<(uint(v5833)%32))))
	F_boot_get_type_io_data(m, v5839, v5712+int32(98), v5712+int32(97), v5712+int32(96), v5712+int32(95), v5712+int32(88), v5712+int32(84), v5712+int32(80), v5712+int32(76))
	mBase = m.M
	v5857 = m.ExcPending
	if v5857 != 0 {
		goto L1
	} else {
		goto L1139
	}
L1138:
	;
	v5893 = v5881
	goto L1136
L1139:
	;
	v5859 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+104))
	v5861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5859+v5807))))
	if v5861 == int32(0) {
		goto L1140
	} else {
		goto L1141
	}
L1140:
	;
	v5864 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+84))
	v5865 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+108))
	v5869 = *(*int32)(unsafe.Add(mBase, uint32(v5865+v5807<<(uint(int32(3))%32))))
	v5870 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+88))
	v5872 = F_OidInputFunctionCall(m, v5864, v5869, v5870, int32(-1))
	mBase = m.M
	v5873 = m.ExcPending
	if v5873 != 0 {
		goto L1
	} else {
		goto L1143
	}
L1141:
	;
	v5874 = int64(0)
	goto L1142
L1142:
	;
	v5876 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+76))
	v5877 = int32(*(*int16)(unsafe.Add(mBase, uint32(v5712)+98)))
	v5878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5712)+97)))
	v5879 = F_makeConst(m, v5839, int32(-1), v5876, v5877, v5874, v5861, v5878)
	mBase = m.M
	v5880 = m.ExcPending
	if v5880 != 0 {
		goto L1
	} else {
		goto L1144
	}
L1143:
	;
	v5874 = v5872
	goto L1142
L1144:
	;
	v5881 = F_lappend(m, v5812, v5879)
	mBase = m.M
	v5882 = m.ExcPending
	if v5882 != 0 {
		goto L1
	} else {
		goto L1145
	}
L1145:
	;
	v5884 = v5807 + int32(1)
	v5885 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+100))
	if v5884 < v5885 {
		v5807 = v5884
		v5812 = v5881
		v5814 = v5885
		goto L1137
	} else {
		goto L1146
	}
L1146:
	;
	goto L1138
L1147:
	;
	v5916 = F_cstring_to_text(m, v5914)
	mBase = m.M
	v5917 = m.ExcPending
	if v5917 != 0 {
		goto L1
	} else {
		goto L1148
	}
L1148:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[96])) = base.I64_extend_i32_u(v5916)
	v5921 = int64(*(*int16)(unsafe.Add(mBase, uint32(v5712)+100)))
	*(*int64)(unsafe.Add(mBase, _c_F_BootstrapModeMain[97])) = v5921
	v5924 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[98])) = uint8(v5924)
	*(*uint8)(unsafe.Add(mBase, _c_F_BootstrapModeMain[99])) = uint8(v5924)
	goto L1121
L1149:
	;
	v5934 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v5935 = *(*int32)(unsafe.Add(mBase, uint32(v5934)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v5712)+20)) = v5739 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v5712)+16)) = v5935 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_132), v5712+int32(16))
	mBase = m.M
	v5946 = m.ExcPending
	if v5946 != 0 {
		goto L1
	} else {
		goto L1150
	}
L1150:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(741), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v5951 = m.ExcPending
	if v5951 != 0 {
		goto L1
	} else {
		goto L1151
	}
L1151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1152:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5706<<(uint(int32(3))%32))+uint32(_c_F_BootstrapModeMain[100]))) = v5959
	goto L1121
L1153:
	;
	if v5990 != 0 {
		goto L1154
	} else {
		goto L1155
	}
L1154:
	;
	v5992 = *(*int32)(unsafe.Add(mBase, uint32(v5712)+56))
	v5997 = *(*int64)(unsafe.Add(mBase, uint32(v5706<<(uint(int32(3))%32))+uint32(_c_F_BootstrapModeMain[100])))
	v5998 = F_OidOutputFunctionCall(m, v5992, v5997)
	mBase = m.M
	v5999 = m.ExcPending
	if v5999 != 0 {
		goto L1
	} else {
		goto L1157
	}
L1155:
	;
	goto L1156
L1156:
	;
	m.G0 = v5712 + int32(112)
	goto L1117
L1157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5712))) = v5998
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_133), v5712)
	mBase = m.M
	v6003 = m.ExcPending
	if v6003 != 0 {
		goto L1
	} else {
		goto L1158
	}
L1158:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(755), int32(_a_F_BootstrapModeMain_131))
	mBase = m.M
	v6008 = m.ExcPending
	if v6008 != 0 {
		goto L1
	} else {
		goto L1159
	}
L1159:
	;
	goto L1156
L1160:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_134), int32(0))
	mBase = m.M
	v6019 = m.ExcPending
	if v6019 != 0 {
		goto L1
	} else {
		goto L1161
	}
L1161:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(787), int32(_a_F_BootstrapModeMain_135))
	mBase = m.M
	v6024 = m.ExcPending
	if v6024 != 0 {
		goto L1
	} else {
		goto L1162
	}
L1162:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1163:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_136), int32(0))
	mBase = m.M
	v6032 = m.ExcPending
	if v6032 != 0 {
		goto L1
	} else {
		goto L1164
	}
L1164:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(789), int32(_a_F_BootstrapModeMain_135))
	mBase = m.M
	v6037 = m.ExcPending
	if v6037 != 0 {
		goto L1
	} else {
		goto L1165
	}
L1165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1166:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_137), int32(0))
	mBase = m.M
	v6045 = m.ExcPending
	if v6045 != 0 {
		goto L1
	} else {
		goto L1167
	}
L1167:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(804), int32(_a_F_BootstrapModeMain_135))
	mBase = m.M
	v6050 = m.ExcPending
	if v6050 != 0 {
		goto L1
	} else {
		goto L1168
	}
L1168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1169:
	;
	if v6063 != 0 {
		goto L1170
	} else {
		goto L1171
	}
L1170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6059)+16)) = v6053
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_138), v6059+int32(16))
	mBase = m.M
	v6070 = m.ExcPending
	if v6070 != 0 {
		goto L1
	} else {
		goto L1173
	}
L1171:
	;
	goto L1172
L1172:
	;
	v6077 = v6053 * int32(100)
	v6079 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v6080 = *(*int32)(unsafe.Add(mBase, uint32(v6079)+52))
	v6081 = *(*int32)(unsafe.Add(mBase, uint32(v6080)))
	v6086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6077+(v6080+v6081<<(uint(int32(3))%32)))+114)))
	if v6086 == int32(1) {
		goto L1175
	} else {
		goto L1176
	}
L1173:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(870), int32(_a_F_BootstrapModeMain_139))
	mBase = m.M
	v6075 = m.ExcPending
	if v6075 != 0 {
		goto L1
	} else {
		goto L1174
	}
L1174:
	;
	goto L1172
L1175:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v6092 = m.ExcPending
	if v6092 != 0 {
		goto L1
	} else {
		goto L1178
	}
L1176:
	;
	goto L1177
L1177:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6053<<(uint(int32(3))%32))+uint32(_c_F_BootstrapModeMain[100]))) = int64(0)
	v6124 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6053)+uint32(_c_F_BootstrapModeMain[57]))) = uint8(v6124)
	m.G0 = v6059 + int32(32)
	v6189 = v3792
	goto L675
L1178:
	;
	v6094 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	v6095 = *(*int32)(unsafe.Add(mBase, uint32(v6094)+52))
	v6096 = *(*int32)(unsafe.Add(mBase, uint32(v6095)))
	v6097 = *(*int32)(unsafe.Add(mBase, uint32(v6094)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v6059)+4)) = v6097 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v6059))) = v6095 + v6096<<(uint(int32(3))%32) + v6077 + int32(32)
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_140), v6059)
	mBase = m.M
	v6110 = m.ExcPending
	if v6110 != 0 {
		goto L1
	} else {
		goto L1179
	}
L1179:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(876), int32(_a_F_BootstrapModeMain_139))
	mBase = m.M
	v6115 = m.ExcPending
	if v6115 != 0 {
		goto L1
	} else {
		goto L1180
	}
L1180:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1181:
	;
	v6189 = v6131
	goto L675
L1182:
	;
	v6189 = v6134
	goto L675
L1183:
	;
	v6189 = v6137
	goto L675
L1184:
	;
	v6189 = v6140
	goto L675
L1185:
	;
	v6189 = v6143
	goto L675
L1186:
	;
	v6189 = v6146
	goto L675
L1187:
	;
	v6189 = v6149
	goto L675
L1188:
	;
	v6189 = v6152
	goto L675
L1189:
	;
	v6189 = v6155
	goto L675
L1190:
	;
	v6189 = v6158
	goto L675
L1191:
	;
	v6189 = v6161
	goto L675
L1192:
	;
	v6189 = v6164
	goto L675
L1193:
	;
	v6189 = v6167
	goto L675
L1194:
	;
	v6189 = v6170
	goto L675
L1195:
	;
	v6189 = v6173
	goto L675
L1196:
	;
	v6189 = v6176
	goto L675
L1197:
	;
	v6189 = v6179
	goto L675
L1198:
	;
	v6189 = v6182
	goto L675
L1199:
	;
	v6189 = v6185
	goto L675
L1200:
	;
	v6239 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6223)+uint32(_c_F_BootstrapModeMain[101]))))
	v6243 = v3761
	v6247 = v6218
	v6250 = v3768
	v6251 = v3769
	v6252 = v3770
	v6253 = v6239
	v6254 = v6219
	v6256 = v3774
	v6257 = v3775
	v6259 = v3777
	v6262 = v3780
	goto L294
L1201:
	;
	v6230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6227)+uint32(_c_F_BootstrapModeMain[71]))))
	if v6230 != v6220&int32(255) {
		goto L1200
	} else {
		goto L1202
	}
L1202:
	;
	v6236 = int32(*(*int8)(unsafe.Add(mBase, uint32(v6227)+uint32(_c_F_BootstrapModeMain[72]))))
	v6243 = v3761
	v6247 = v6218
	v6250 = v3768
	v6251 = v3769
	v6252 = v3770
	v6253 = v6236
	v6254 = v6219
	v6256 = v3774
	v6257 = v3775
	v6259 = v3777
	v6262 = v3780
	goto L294
L1203:
	;
	if v3738 == int32(0) {
		v6389 = v3737
		v6391 = v3739
		goto L262
	} else {
		goto L1206
	}
L1204:
	;
	F_boot_yyerror(m, v3730, int32(_a_F_BootstrapModeMain_141))
	mBase = m.M
	v6270 = m.ExcPending
	if v6270 != 0 {
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
	v6283 = v3737
	v6285 = v3739
	v6287 = v3741
	goto L267
L1207:
	;
	if v6285 == v6313 {
		v6389 = v6283
		v6391 = v6285
		goto L262
	} else {
		goto L1209
	}
L1209:
	;
	v6313 = v6313 - int32(1)
	goto L1207
L1210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1211:
	;
	v6338 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[76]))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+32)) = v6338
	v6341 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[75]))
	*(*int32)(unsafe.Add(mBase, uint32(v3768)+36)) = v6341
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_142), v3768+int32(32))
	mBase = m.M
	v6347 = m.ExcPending
	if v6347 != 0 {
		goto L1
	} else {
		goto L1212
	}
L1212:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(265), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v6352 = m.ExcPending
	if v6352 != 0 {
		goto L1
	} else {
		goto L1213
	}
L1213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1214:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_143), int32(0))
	mBase = m.M
	v6360 = m.ExcPending
	if v6360 != 0 {
		goto L1
	} else {
		goto L1215
	}
L1215:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(267), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v6365 = m.ExcPending
	if v6365 != 0 {
		goto L1
	} else {
		goto L1216
	}
L1216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1217:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_144), int32(0))
	mBase = m.M
	v6373 = m.ExcPending
	if v6373 != 0 {
		goto L1
	} else {
		goto L1218
	}
L1218:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_103), int32(449), int32(_a_F_BootstrapModeMain_104))
	mBase = m.M
	v6378 = m.ExcPending
	if v6378 != 0 {
		goto L1
	} else {
		goto L1219
	}
L1219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1220:
	;
	F_pfree(m, v6405)
	mBase = m.M
	v6435 = m.ExcPending
	if v6435 != 0 {
		goto L1
	} else {
		goto L1223
	}
L1221:
	;
	goto L1222
L1222:
	;
	m.G0 = v6415 + int32(1136)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v6440 = m.ExcPending
	if v6440 != 0 {
		goto L1
	} else {
		goto L1224
	}
L1223:
	;
	goto L1222
L1224:
	;
	v6442 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7]))
	v6446 = F_LWLockAcquire(m, v6442+int32(3200), int32(0))
	mBase = m.M
	v6447 = m.ExcPending
	if v6447 != 0 {
		goto L1
	} else {
		goto L1225
	}
L1225:
	;
	v6449 = int32(0)
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_145), v6449, v6449, v6449, v6449, int32(1664), int32(_a_F_BootstrapModeMain_146))
	mBase = m.M
	v6456 = m.ExcPending
	if v6456 != 0 {
		goto L1
	} else {
		goto L1226
	}
L1226:
	;
	v6458 = int32(0)
	v6462 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[102]))
	v6464 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[103]))
	v6466 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[104]))
	F_write_relmap_file(m, int32(_a_F_BootstrapModeMain_147), v6458, v6458, v6458, v6462, v6464, v6466)
	mBase = m.M
	v6468 = m.ExcPending
	if v6468 != 0 {
		goto L1
	} else {
		goto L1227
	}
L1227:
	;
	v6470 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[7]))
	F_LWLockRelease(m, v6470+int32(3200))
	mBase = m.M
	v6474 = m.ExcPending
	if v6474 != 0 {
		goto L1
	} else {
		goto L1228
	}
L1228:
	;
	v6476 = *(*int32)(unsafe.Add(mBase, _c_F_BootstrapModeMain[81]))
	if v6476 != 0 {
		goto L1229
	} else {
		goto L1230
	}
L1229:
	;
	F_closerel(m, int32(0))
	mBase = m.M
	v6479 = m.ExcPending
	if v6479 != 0 {
		goto L1
	} else {
		goto L1232
	}
L1230:
	;
	goto L1231
L1231:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v6482 = m.ExcPending
	if v6482 != 0 {
		goto L1
	} else {
		goto L1233
	}
L1232:
	;
	goto L1231
L1233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1234:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v6489 = m.ExcPending
	if v6489 != 0 {
		goto L1
	} else {
		goto L1235
	}
L1235:
	;
	v6490 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v6490
	F_errmsg(m, int32(_a_F_BootstrapModeMain_148), v29-int32(-64))
	mBase = m.M
	v6496 = m.ExcPending
	if v6496 != 0 {
		goto L1
	} else {
		goto L1236
	}
L1236:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(280), int32(_a_F_BootstrapModeMain_18))
	mBase = m.M
	v6501 = m.ExcPending
	if v6501 != 0 {
		goto L1
	} else {
		goto L1237
	}
L1237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1238:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(295), int32(_a_F_BootstrapModeMain_18))
	mBase = m.M
	v6512 = m.ExcPending
	if v6512 != 0 {
		goto L1
	} else {
		goto L1239
	}
L1239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1240:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6519 = m.ExcPending
	if v6519 != 0 {
		goto L1
	} else {
		goto L1241
	}
L1241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1242:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v6528 = m.ExcPending
	if v6528 != 0 {
		goto L1
	} else {
		goto L1243
	}
L1243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L1246:
	;
	F_errmsg_internal(m, int32(_a_F_BootstrapModeMain_149), int32(0))
	mBase = m.M
	v6545 = m.ExcPending
	if v6545 != 0 {
		goto L1
	} else {
		goto L1247
	}
L1247:
	;
	F_errfinish(m, int32(_a_F_BootstrapModeMain_17), int32(427), int32(_a_F_BootstrapModeMain_18))
	mBase = m.M
	v6550 = m.ExcPending
	if v6550 != 0 {
		goto L1
	} else {
		goto L1248
	}
L1248:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
