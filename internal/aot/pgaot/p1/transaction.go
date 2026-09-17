package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v300 int32
	_ = v300
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int64
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v393 int64
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int64
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int64
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v445 int64
	_ = v445
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v602 int32
	_ = v602
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int64
	_ = v621
	var v622 int64
	_ = v622
	var v627 int32
	_ = v627
	var v629 float64
	_ = v629
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v640 int64
	_ = v640
	var v641 int64
	_ = v641
	var v649 int64
	_ = v649
	var v651 int32
	_ = v651
	var v652 int64
	_ = v652
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int64
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v690 int64
	_ = v690
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v699 int32
	_ = v699
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int64
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int64
	_ = v717
	var v720 int32
	_ = v720
	var v730 int32
	_ = v730
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v800 int64
	_ = v800
	var v801 int64
	_ = v801
	var v802 int64
	_ = v802
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int64
	_ = v813
	var v814 int64
	_ = v814
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v847 int32
	_ = v847
	var v856 int32
	_ = v856
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v927 int64
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int64
	_ = v935
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int64
	_ = v946
	var v947 int64
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int64
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1131 int64
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int64
	_ = v1155
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int64
	_ = v1167
	var v1168 int64
	_ = v1168
	var v1176 int64
	_ = v1176
	var v1178 int64
	_ = v1178
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int64
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1196 int64
	_ = v1196
	var v1198 int64
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int64
	_ = v1202
	var v1207 int64
	_ = v1207
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1219 int64
	_ = v1219
	var v1220 int64
	_ = v1220
	var v1228 int64
	_ = v1228
	var v1230 int64
	_ = v1230
	var v1233 int64
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1245 int64
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1264 int32
	_ = v1264
	var v1268 int64
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1361 int64
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int64
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1386 int64
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int64
	_ = v1399
	var v1404 int32
	_ = v1404
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1513 int32
	_ = v1513
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1564 int32
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1582 int32
	_ = v1582
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1648 int32
	_ = v1648
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1754 int32
	_ = v1754
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1816 int32
	_ = v1816
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1829 int32
	_ = v1829
	var v1835 int32
	_ = v1835
	var v1840 int32
	_ = v1840
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1926 int32
	_ = v1926
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1970 int32
	_ = v1970
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1980 int32
	_ = v1980
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2051 int32
	_ = v2051
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2061 int32
	_ = v2061
	var v2062 int32
	_ = v2062
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2077 int32
	_ = v2077
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2092 int32
	_ = v2092
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int64
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int64
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2118 int32
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2133 int32
	_ = v2133
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2169 int32
	_ = v2169
	var v2171 int32
	_ = v2171
	var v2173 int32
	_ = v2173
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2298 int32
	_ = v2298
	var v2300 int32
	_ = v2300
	var v2307 int32
	_ = v2307
	var v2310 int32
	_ = v2310
	var v2312 int32
	_ = v2312
	var v2315 int32
	_ = v2315
	var v2317 int32
	_ = v2317
	var v2330 int32
	_ = v2330
	var v2333 int32
	_ = v2333
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2371 int32
	_ = v2371
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2398 int32
	_ = v2398
	var v2404 int64
	_ = v2404
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	v24 = m.G0
	v26 = v24 - int32(48)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v30 == int32(5) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v33 + int32(1)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v37 = int32(10)
	goto L6
L4:
	;
	if v74 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[1]))
	goto L9
L7:
	;
	v57 = int32(0)
	goto L14
L9:
	;
	goto L10
L10:
	;
	if int32(0)|base.B2i32(v44 == int32(15)) != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if v44 <= v37 {
		v74 = int32(1)
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[2]))
	if v61 != int32(2) {
		v74 = v57
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[3])))
	if v65&int32(1) != 0 {
		v74 = v57
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[4]))
	v74 = int32(0) | base.B2i32(v71 <= v37)
	goto L5
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_CommitTransaction_0), v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if v81 == int32(2) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	return
L21:
	;
	goto L19
L22:
	;
	goto L31
L23:
	;
	v86 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	if v86 == int32(0) {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	if base.Ui32(v90) <= base.Ui32(int32(5)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90<<(uint(int32(2))%32))+uint32(_c_F_CommitTransaction[5])))
	v97 = v95
	goto L28
L27:
	;
	v97 = int32(_a_F_CommitTransaction_1)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v97
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_2), v26+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), int32(2247), int32(_a_F_CommitTransaction_0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	F_AfterTriggerFireDeferred(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L20
	} else {
		goto L33
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[6]))
	if v139 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v136 = F_PreCommit_Portals(m, int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L20
	} else {
		goto L34
	}
L34:
	;
	if v136 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v141 = int32(5)
	if v30 == v141 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	F_AtEOXact_Parallel(m, int32(1))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L20
	} else {
		goto L46
	}
L39:
	;
	v144 = int32(6)
	goto L41
L40:
	;
	v144 = v141
	goto L41
L41:
	;
	v145 = v139
	goto L42
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	m.T0[v170].(func(*base.Module, int32, int32))(m, v144, v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L20
	} else {
		goto L44
	}
L43:
	;
	goto L38
L44:
	;
	if v168 != 0 {
		v145 = v168
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if v30 == int32(5) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L20
	} else {
		goto L60
	}
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v224
	F_errmsg_internal(m, v223, v26)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L20
	} else {
		goto L58
	}
L49:
	;
	if v199 == int32(1) {
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v199 == int32(0) {
		goto L47
	} else {
		goto L55
	}
L52:
	;
	v206 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L20
	} else {
		goto L53
	}
L53:
	;
	if v206 == int32(0) {
		goto L47
	} else {
		goto L54
	}
L54:
	;
	v222 = int32(2295)
	v223 = int32(_a_F_CommitTransaction_4)
	goto L48
L55:
	;
	v216 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L20
	} else {
		goto L56
	}
L56:
	;
	if v216 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v222 = int32(2301)
	v223 = int32(_a_F_CommitTransaction_5)
	goto L48
L58:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), v222, int32(_a_F_CommitTransaction_0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L20
	} else {
		goto L59
	}
L59:
	;
	goto L47
L60:
	;
	F_PreCommit_on_commit_actions(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v239 = base.B2i32(v30 == int32(5))
	F_smgrDoPendingSyncs(m, int32(1), v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	v245 = m.G0
	v247 = v245 - int32(_a_F_CommitTransaction_6)
	m.G0 = v247
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v250|v252 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if v239 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L20
	} else {
		goto L208
	}
L66:
	;
	m.G0 = v247 + int32(_a_F_CommitTransaction_6)
	goto L64
L67:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[9])))
	if v257 != int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	if v276 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v262 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	if v262 == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_7), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(868), int32(_a_F_CommitTransaction_7))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L20
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v569 == int32(0) {
		goto L66
	} else {
		goto L119
	}
L75:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	if v279 == int32(0) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v282 <= int32(0) {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v300 = int32(0)
	goto L78
L78:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+v300<<(uint(int32(2))%32))))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	if v313 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L74
L80:
	;
	v542 = v300 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v542 < v543 {
		v300 = v542
		goto L78
	} else {
		goto L118
	}
L81:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])))
	if v315 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[9])))
	if v317 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v340 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[11])))
	if v340 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v322 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L20
	} else {
		goto L85
	}
L85:
	;
	if v322 == int32(0) {
		goto L83
	} else {
		goto L86
	}
L86:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+32)) = v327
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_9), v247+int32(32))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1054), int32(_a_F_CommitTransaction_10))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	goto L83
L89:
	;
	F_before_shmem_exit(m, int32(514), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L20
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v355 = F_LWLockAcquire(m, v351+int32(3456), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L20
	} else {
		goto L93
	}
L92:
	;
	v348 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[11])) = uint8(v348)
	goto L91
L93:
	;
	v357 = int32(-1)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)+28))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v361)+16))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v361)))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v361)+40))
	if v367 != v357 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v374 = v363
	v376 = v367
	v381 = v362
	v387 = v357
	v393 = v364
	goto L97
L95:
	;
	v426 = v363
	v433 = v362
	v439 = v357
	v445 = v364
	goto L96
L96:
	;
	v449 = int32(5)
	v451 = v361 + v359<<(uint(v449)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v451)+84)) = v433
	*(*int32)(unsafe.Add(mBase, uint32(v451)+80)) = v426
	*(*int64)(unsafe.Add(mBase, uint32(v451)+72)) = v445
	v456 = v361 + int32(56)
	v457 = int32(_a_F_CommitTransaction_11)
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v456+v458<<(uint(v449)%32)))) = v463
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v471 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v466<<(uint(v449)%32))+60)) = v471
	if v439 != int32(-1) {
		goto L112
	} else {
		goto L113
	}
L97:
	;
	v398 = v376 << (uint(int32(5)) % 32)
	v399 = v361 + int32(56) + v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v399)+4))
	if v400 != v373 {
		v412 = v374
		v414 = v381
		v416 = v393
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v426 = v412
	v433 = v414
	v439 = v419
	v445 = v416
	goto L96
L99:
	;
	if v376 < v359 {
		goto L107
	} else {
		goto L108
	}
L100:
	;
	v402 = v398 + v361
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v399)+16))
	if v393 < v403 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v402)+84))
	v412 = v409
	v414 = v411
	v416 = v410
	goto L99
L102:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v402)+80))
	v409 = v405
	v410 = v403
	goto L101
L103:
	;
	goto L104
L104:
	;
	if v393 != v403 {
		v412 = v374
		v414 = v381
		v416 = v393
		goto L99
	} else {
		goto L105
	}
L105:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v402)+80))
	if v407 < v374 {
		v412 = v374
		v414 = v381
		v416 = v393
		goto L99
	} else {
		goto L106
	}
L106:
	;
	v409 = v407
	v410 = v393
	goto L101
L107:
	;
	v419 = v376
	goto L109
L108:
	;
	v419 = v387
	goto L109
L109:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v398+v361-int32(-64))))
	if v423 != int32(-1) {
		v374 = v412
		v376 = v423
		v381 = v414
		v387 = v419
		v393 = v416
		goto L97
	} else {
		goto L110
	}
L110:
	;
	goto L98
L111:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v505+int32(3456))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L20
	} else {
		goto L115
	}
L112:
	;
	v475 = int32(_a_F_CommitTransaction_11)
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v477 = int32(5)
	v481 = v439 << (uint(v477) % 32)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v456+v481)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v456+v476<<(uint(v477)%32))+8)) = v483
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v481+v361-int32(-64)))) = v489
	goto L111
L113:
	;
	goto L114
L114:
	;
	v491 = int32(_a_F_CommitTransaction_11)
	v492 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v361)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v492<<(uint(int32(5))%32)-int32(-64)))) = v498
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v361)+40)) = v501
	goto L111
L115:
	;
	v511 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])) = uint8(v511)
	if base.B2i32(v426 == v365)&base.B2i32(v445 == v366) != 0 {
		goto L80
	} else {
		goto L116
	}
L116:
	;
	F_asyncQueueReadAllNotifications(m)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L20
	} else {
		goto L117
	}
L117:
	;
	goto L80
L118:
	;
	goto L79
L119:
	;
	v572 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L20
	} else {
		goto L120
	}
L120:
	;
	F_LockSharedObject(m, int32(1262), int32(0), int32(8))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L20
	} else {
		goto L121
	}
L121:
	;
	v580 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v580)+4))
	if v581 == int32(0) {
		goto L66
	} else {
		goto L122
	}
L122:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v581)+12))
	if v584 == int32(0) {
		goto L66
	} else {
		goto L123
	}
L123:
	;
	v602 = v584
	goto L124
L124:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v617 = F_LWLockAcquire(m, v613+int32(3456), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L20
	} else {
		goto L126
	}
L125:
	;
	goto L66
L126:
	;
	v620 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v621 = *(*int64)(unsafe.Add(mBase, uint32(v620)))
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v620)+16))
	if v621 == v622 {
		v778 = v620
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v800 = int64(*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17])))
	v801 = *(*int64)(unsafe.Add(mBase, uint32(v778)))
	v802 = *(*int64)(unsafe.Add(mBase, uint32(v778)+16))
	if v800 <= v801-v802 {
		goto L65
	} else {
		goto L163
	}
L128:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[17]))
	v629 = base.F64_div(base.F64_convert_i64_s(v621-v622), base.F64_convert_i32_s(v627))
	if base.F64_lt(v629, float64(0.5)) != 0 {
		v778 = v620
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v635 = m.G0
	v636 = int32(16)
	v637 = v635 - v636
	m.G0 = v637
	F_gettimeofday(m, v637)
	mBase = m.M
	v640 = *(*int64)(unsafe.Add(mBase, uint32(v637)))
	v641 = int64(*(*int32)(unsafe.Add(mBase, uint32(v637)+8)))
	m.G0 = v637 + v636
	v649 = v641 + v640*int64(1000000) - int64(946684800000000)
	goto L130
L130:
	;
	v651 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v652 = *(*int64)(unsafe.Add(mBase, uint32(v651)+48))
	goto L131
L131:
	;
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	if base.B2i32(base.I64_extend_i32_s(int32(_a_F_CommitTransaction_12))*int64(1000) <= v649-v652) == int32(0) {
		v778 = v660
		goto L127
	} else {
		goto L132
	}
L132:
	;
	v663 = int32(-1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660)+40))
	if v664 != v663 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v668 = v660 + int32(56)
	v669 = *(*int64)(unsafe.Add(mBase, uint32(v660)))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v660)+8))
	v674 = v670
	v675 = v664
	v678 = v663
	v690 = v669
	goto L136
L134:
	;
	v730 = v663
	goto L135
L135:
	;
	v748 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L20
	} else {
		goto L152
	}
L136:
	;
	v695 = v675 << (uint(int32(5)) % 32)
	v696 = v660 + v695
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v696)+72))
	if v697 <= v690 {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	v730 = v716
	goto L135
L138:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v696-int32(-64))))
	if v720 != int32(-1) {
		v674 = v715
		v675 = v720
		v678 = v716
		v690 = v717
		goto L136
	} else {
		goto L151
	}
L139:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v695+v668)))
	v715 = v709
	v716 = v713
	v717 = v711
	goto L138
L140:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v696)+80))
	if v690 != v697 {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v704 = v674
	goto L142
L142:
	;
	if v690 != v697 {
		v715 = v704
		v716 = v678
		v717 = v690
		goto L138
	} else {
		goto L149
	}
L143:
	;
	v709 = v699
	v711 = v697
	goto L139
L144:
	;
	goto L145
L145:
	;
	if v674 < v699 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v702 = v674
	goto L148
L147:
	;
	v702 = v699
	goto L148
L148:
	;
	v704 = v702
	goto L142
L149:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v695+v668)+24))
	if v704 != v707 {
		v715 = v704
		v716 = v678
		v717 = v690
		goto L138
	} else {
		goto L150
	}
L150:
	;
	v709 = v707
	v711 = v690
	goto L139
L151:
	;
	goto L137
L152:
	;
	if v748 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v247)+16)) = base.F64_mul(v629, float64(100))
	F_errmsg(m, int32(_a_F_CommitTransaction_13), v247+int32(16))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L20
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v774 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v774)+48)) = v649
	v778 = v774
	goto L127
L156:
	;
	if v730 != int32(-1) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v730
	F_errdetail(m, int32(_a_F_CommitTransaction_14), v247)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1569), int32(_a_F_CommitTransaction_15))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L20
	} else {
		goto L162
	}
L160:
	;
	F_errhint(m, int32(_a_F_CommitTransaction_16), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L20
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	goto L155
L163:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v247)+56)) = v801
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v778)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+48)) = v806
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+28))
	v813 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[19])))
	v814 = base.I64_rem_s(v801, v813)
	v818 = v811 + base.I32_wrap_i64(v814)<<(uint(int32(7))%32)
	v820 = F_LWLockAcquire(m, v818, int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	if base.B2i32(v801 != int64(0))|v806 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+12))
	v841 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v836+v839))) = uint8(v841)
	v847 = v806
	v856 = v602
	goto L172
L166:
	;
	v829 = F_SimpleLruZeroPage(m, int32(_a_F_CommitTransaction_17), int64(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L20
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v834 = F_SimpleLruReadPage(m, int32(_a_F_CommitTransaction_17), v801, int32(1), int32(0))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L20
	} else {
		goto L170
	}
L169:
	;
	v836 = v829
	goto L165
L170:
	;
	v836 = v834
	goto L165
L171:
	;
	v974 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v974))) = v935
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v247)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v974)+12)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v974)+8)) = v976
	F_LWLockRelease(m, v970)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L20
	} else {
		goto L205
	}
L172:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v856)))
	v867 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v866)+2)))
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v866))))
	v870 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+68)) = v870
	v872 = v868 + v867
	v876 = (v872 + int32(21)) & int32(_a_F_CommitTransaction_18)
	*(*int32)(unsafe.Add(mBase, uint32(v247)+64)) = v876
	v878 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L20
	} else {
		goto L174
	}
L173:
	;
	v970 = v818
	v972 = int32(0)
	goto L171
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247)+72)) = v878
	v882 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v247)+76)) = v882
	v885 = v872 + int32(2)
	if v885 != 0 {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	base.MemoryCopy(m, v247+int32(80), v866+int32(4), v885)
	goto L177
L176:
	;
	goto L177
L177:
	;
	if v876+v847 <= int32(_a_F_CommitTransaction_19) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	if v913 != 0 {
		goto L185
	} else {
		goto L186
	}
L179:
	;
	v893 = v856 + int32(4)
	v896 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v896)+4))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v897)+12))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
	if base.Ui32(v893) < base.Ui32(v898+v899<<(uint(int32(2))%32)) {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	goto L181
L181:
	;
	v905 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+80)) = uint16(v905)
	*(*int64)(unsafe.Add(mBase, uint32(v247)+68)) = int64(0)
	v910 = int32(_a_F_CommitTransaction_19) - v847
	*(*int32)(unsafe.Add(mBase, uint32(v247)+64)) = v910
	v913 = v910
	v915 = v856
	goto L178
L182:
	;
	v904 = v893
	goto L184
L183:
	;
	v904 = int32(0)
	goto L184
L184:
	;
	v913 = v876
	v915 = v904
	goto L178
L185:
	;
	v917 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v917)+4))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v918+v836<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v922+v847, v247-int32(-64), v913)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v927 = *(*int64)(unsafe.Add(mBase, uint32(v247)+56))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v247)+48))
	v929 = v928 + v913
	v931 = v929 - int32(_a_F_CommitTransaction_20)
	v933 = base.B2i32(base.Ui32(v931) < base.Ui32(int32(-8193)))
	v935 = v927 + base.I64_extend_i32_u(v933)
	*(*int64)(unsafe.Add(mBase, uint32(v247)+56)) = v935
	if base.Ui32(v931) < base.Ui32(int32(-8193)) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v938 = int32(0)
	goto L190
L189:
	;
	v938 = v929
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247)+48)) = v938
	if base.Ui32(v931) <= base.Ui32(int32(-8194)) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v943 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[18]))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+28))
	v946 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[19])))
	v947 = base.I64_rem_s(v935, v946)
	v951 = v944 + base.I32_wrap_i64(v947)<<(uint(int32(7))%32)
	if v818 == v951 {
		goto L195
	} else {
		goto L196
	}
L192:
	;
	goto L193
L193:
	;
	if v915 != 0 {
		v847 = v929
		v856 = v915
		goto L172
	} else {
		goto L204
	}
L194:
	;
	v960 = F_SimpleLruZeroPage(m, int32(_a_F_CommitTransaction_17), v935)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L20
	} else {
		goto L200
	}
L195:
	;
	v958 = v818
	goto L194
L196:
	;
	goto L197
L197:
	;
	F_LWLockRelease(m, v818)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L20
	} else {
		goto L198
	}
L198:
	;
	v956 = F_LWLockAcquire(m, v951, int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L20
	} else {
		goto L199
	}
L199:
	;
	v958 = v951
	goto L194
L200:
	;
	if v935&int64(3) == int64(0) {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v967 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])) = uint8(v967)
	goto L203
L202:
	;
	goto L203
L203:
	;
	v970 = v958
	v972 = v915
	goto L171
L204:
	;
	goto L173
L205:
	;
	v982 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v982+int32(3456))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L20
	} else {
		goto L206
	}
L206:
	;
	if v972 != 0 {
		v602 = v972
		goto L124
	} else {
		goto L207
	}
L207:
	;
	goto L125
L208:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L20
	} else {
		goto L209
	}
L209:
	;
	F_errmsg(m, int32(_a_F_CommitTransaction_21), int32(0))
	mBase = m.M
	v1023 = m.ExcPending
	if v1023 != 0 {
		goto L20
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(945), int32(_a_F_CommitTransaction_7))
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L20
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L20
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v1033 = int32(_a_F_CommitTransaction_22)
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[21]))
	v1036 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[21])) = v1035 + v1036
	F_AtEOXact_RelationMap(m, v1036, v239)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L20
	} else {
		goto L216
	}
L215:
	;
	goto L214
L216:
	;
	v1042 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+76)) = uint8(v1042)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v1042
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(3)
	v1049 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[22]))
	if v1042 < v1049 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L20
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	if v30 != int32(5) {
		goto L223
	} else {
		goto L224
	}
L220:
	;
	goto L219
L221:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L20
	} else {
		goto L490
	}
L222:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	F_ProcArrayEndTransaction(m, v1416, v1404)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L20
	} else {
		goto L330
	}
L223:
	;
	v1057 = int32(0)
	v1059 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v1057
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v1057
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+35)) = uint8(v1057)
	v1067 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	if int32(2) <= v1067 {
		goto L226
	} else {
		goto L227
	}
L224:
	;
	goto L225
L225:
	;
	v1386 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[27]))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v1388)+68)) = int32(1)
	v1393 = v1388 + int32(68)
	if v1389 != 0 {
		goto L323
	} else {
		goto L324
	}
L226:
	;
	F_LogLogicalInvalidations(m)
	mBase = m.M
	v1071 = m.ExcPending
	if v1071 != 0 {
		goto L20
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1075 = F_smgrGetPendingDeletes(m, int32(1), v26+int32(44))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L20
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+52))
	if v1079 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1078)+48))
	v1081 = v1080
	goto L233
L232:
	;
	v1081 = v1057
	goto L233
L233:
	;
	v1085 = F_pgstat_get_transactional_drops(m, int32(1), v26+int32(40))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	v1087 = int32(0)
	v1089 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[25]))
	if v1087 < v1089 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1096 = F_xactGetCommittedInvalidationMessages(m, v26+int32(36), v26+int32(35))
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L20
	} else {
		goto L238
	}
L236:
	;
	v1098 = v1087
	goto L237
L237:
	;
	v1100 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	if v1059 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	v1098 = v1096
	goto L237
L239:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	if v1376 != 0 {
		goto L317
	} else {
		goto L318
	}
L240:
	;
	v1245 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	v1247 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[28]))
	if v1240&base.B2i32(int32(0) < v1247) != 0 {
		goto L272
	} else {
		goto L273
	}
L241:
	;
	if v1085|v1075 != 0 {
		goto L221
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1140 = int32(_a_F_CommitTransaction_23)
	v1142 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	v1143 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v1142 + v1143
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[30])))
	v1149 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1149)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1149)+120)) = v1150 | v1143
	v1155 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31]))
	if v1155 == int64(0) {
		goto L253
	} else {
		goto L254
	}
L244:
	;
	if v1098 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+35)))
	v1106 = m.G0
	v1108 = v1106 - int32(16)
	m.G0 = v1108
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1108)+8)) = uint8(v1105)
	v1114 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v1108))) = v1114
	v1117 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+4)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+12)) = v1098
	F_XLogBeginInsert(m)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L20
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v1137 = int32(0)
	if v1100 != int64(0) {
		v1240 = v1137
		goto L240
	} else {
		goto L252
	}
L248:
	;
	F_XLogRegisterData(m, v1108, int32(16))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L20
	} else {
		goto L249
	}
L249:
	;
	F_XLogRegisterData(m, v1104, v1098<<(uint(int32(4))%32))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L20
	} else {
		goto L250
	}
L250:
	;
	v1131 = F_XLogInsert(m, int32(8), int32(32))
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L20
	} else {
		goto L251
	}
L251:
	;
	m.G0 = v1108 + int32(16)
	v1240 = int32(0)
	goto L240
L252:
	;
	v1372 = v1137
	goto L239
L253:
	;
	v1162 = m.G0
	v1163 = int32(16)
	v1164 = v1162 - v1163
	m.G0 = v1164
	F_gettimeofday(m, v1164)
	mBase = m.M
	v1167 = *(*int64)(unsafe.Add(mBase, uint32(v1164)))
	v1168 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1164)+8)))
	m.G0 = v1164 + v1163
	v1176 = v1168 + v1167*int64(1000000) - int64(946684800000000)
	goto L256
L254:
	;
	v1178 = v1155
	goto L255
L255:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v1184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+35)))
	v1186 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[33]))
	v1187 = int32(0)
	v1189 = F_XactLogCommitRecord(m, v1178, v1079, v1081, v1075, v1181, v1085, v1182, v1098, v1183, v1184, v1186, v1187, v1187)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L20
	} else {
		goto L257
	}
L256:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31])) = v1176
	v1178 = v1176
	goto L255
L257:
	;
	if base.Ui32(int32(2)) <= base.Ui32((v1147+int32(1))&int32(_a_F_CommitTransaction_24)) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1237 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_CommitTransaction[30])))
	F_TransactionTreeSetCommitTsData(m, v1059, v1079, v1081, v1233, v1237)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L20
	} else {
		goto L268
	}
L259:
	;
	v1196 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[34]))
	v1198 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	F_replorigin_session_advance(m, v1196, v1198)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L20
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v1207 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31]))
	if v1207 == int64(0) {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v1202 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[35]))
	if v1202 != int64(0) {
		v1233 = v1202
		goto L258
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v1214 = m.G0
	v1215 = int32(16)
	v1216 = v1214 - v1215
	m.G0 = v1216
	F_gettimeofday(m, v1216)
	mBase = m.M
	v1219 = *(*int64)(unsafe.Add(mBase, uint32(v1216)))
	v1220 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1216)+8)))
	m.G0 = v1216 + v1215
	v1228 = v1220 + v1219*int64(1000000) - int64(946684800000000)
	goto L267
L265:
	;
	v1230 = v1207
	goto L266
L266:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[35])) = v1230
	v1233 = v1230
	goto L258
L267:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[31])) = v1228
	v1230 = v1228
	goto L266
L268:
	;
	v1240 = base.B2i32(v1100 != int64(0))
	goto L240
L269:
	;
	v1287 = v1079 - int32(1)
	if v1287 < int32(0) {
		v1355 = v1059
		goto L283
	} else {
		goto L284
	}
L270:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[23]))
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v1272)+120)) = v1273 & int32(-2)
	v1277 = int32(_a_F_CommitTransaction_23)
	v1279 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v1279 - int32(1)
	goto L269
L271:
	;
	F_XLogSetAsyncXactLSN(m, v1245)
	mBase = m.M
	v1264 = m.ExcPending
	if v1264 != 0 {
		goto L20
	} else {
		goto L279
	}
L272:
	;
	F_XLogFlush(m, v1245)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L20
	} else {
		goto L276
	}
L273:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[36])))
	if v1252&int32(1) != 0 {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	if v1075 <= int32(0) {
		goto L271
	} else {
		goto L275
	}
L275:
	;
	goto L272
L276:
	;
	if v1059 == int32(0) {
		goto L269
	} else {
		goto L277
	}
L277:
	;
	F_TransactionIdCommitTree(m, v1059, v1079, v1081)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L20
	} else {
		goto L278
	}
L278:
	;
	goto L270
L279:
	;
	if v1059 == int32(0) {
		goto L269
	} else {
		goto L280
	}
L280:
	;
	v1268 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	F_TransactionIdAsyncCommitTree(m, v1059, v1079, v1081, v1268)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L20
	} else {
		goto L281
	}
L281:
	;
	goto L270
L282:
	;
	if v1240 != 0 {
		goto L313
	} else {
		goto L314
	}
L283:
	;
	goto L282
L284:
	;
	if v1079&int32(1) != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1292 = int32(2)
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1081+v1287<<(uint(v1292)%32))))
	if base.B2i32(base.Ui32(v1292) < base.Ui32(v1295))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1059)) == int32(0) {
		goto L290
	} else {
		goto L291
	}
L286:
	;
	v1310 = v1059
	v1312 = v1287
	goto L287
L287:
	;
	if v1287 == int32(0) {
		v1355 = v1310
		goto L283
	} else {
		goto L295
	}
L288:
	;
	v1310 = v1307
	v1312 = v1079 - int32(2)
	goto L287
L289:
	;
	v1307 = v1295
	goto L288
L290:
	;
	if base.Ui32(v1059) < base.Ui32(v1295) {
		goto L289
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	if int32(0) <= v1059-v1295 {
		v1307 = v1059
		goto L288
	} else {
		goto L294
	}
L293:
	;
	v1307 = v1059
	goto L288
L294:
	;
	goto L289
L295:
	;
	v1315 = v1310
	v1316 = v1312
	goto L296
L296:
	;
	v1320 = int32(3)
	v1324 = v1081 + v1316<<(uint(int32(2))%32)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1324)))
	if base.B2i32(base.Ui32(v1315) < base.Ui32(v1320))|base.B2i32(base.Ui32(v1325) < base.Ui32(v1320)) == int32(0) {
		goto L300
	} else {
		goto L301
	}
L297:
	;
	v1355 = v1350
	goto L283
L298:
	;
	v1338 = *(*int32)(unsafe.Add(mBase, uint32(v1324-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v1338))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v1335)) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L299:
	;
	v1335 = v1325
	goto L298
L300:
	;
	if v1315-v1325 < int32(0) {
		goto L299
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	if base.Ui32(v1325) <= base.Ui32(v1315) {
		v1335 = v1315
		goto L298
	} else {
		goto L304
	}
L303:
	;
	v1335 = v1315
	goto L298
L304:
	;
	goto L299
L305:
	;
	if int32(1) < v1316 {
		v1315 = v1350
		v1316 = v1316 - int32(2)
		goto L296
	} else {
		goto L312
	}
L306:
	;
	v1350 = v1338
	goto L305
L307:
	;
	if base.Ui32(v1335) < base.Ui32(v1338) {
		goto L306
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	if int32(0) <= v1335-v1338 {
		v1350 = v1335
		goto L305
	} else {
		goto L311
	}
L310:
	;
	v1350 = v1335
	goto L305
L311:
	;
	goto L306
L312:
	;
	goto L297
L313:
	;
	v1361 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	F_SyncRepWaitForLSN(m, v1361, int32(1))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L20
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v1366 = int32(_a_F_CommitTransaction_25)
	v1367 = *(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26]))
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[37])) = v1367
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[26])) = int64(0)
	v1372 = v1355
	goto L239
L316:
	;
	goto L315
L317:
	;
	F_pfree(m, v1376)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L20
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	if v1085 == int32(0) {
		v1404 = v1372
		goto L222
	} else {
		goto L321
	}
L320:
	;
	goto L319
L321:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
	F_pfree(m, v1381)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L20
	} else {
		goto L322
	}
L322:
	;
	v1404 = v1372
	goto L222
L323:
	;
	F_s_lock(m, v1393, int32(_a_F_CommitTransaction_26), int32(1598), int32(_a_F_CommitTransaction_27))
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L20
	} else {
		goto L326
	}
L324:
	;
	goto L325
L325:
	;
	v1399 = *(*int64)(unsafe.Add(mBase, uint32(v1388)+72))
	if base.Ui64(v1399) < base.Ui64(v1386) {
		goto L327
	} else {
		goto L328
	}
L326:
	;
	goto L325
L327:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1388)+72)) = v1386
	goto L329
L328:
	;
	goto L329
L329:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1393))) = int32(0)
	v1404 = int32(0)
	goto L222
L330:
	;
	v1420 = base.B2i32(v30 == int32(5))
	v1422 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[6]))
	if v1422 != 0 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1423 = v1422
	goto L334
L332:
	;
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[38])) = int32(0)
	v1478 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39]))
	v1479 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v1478, v1479, v1479, v1479)
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L20
	} else {
		goto L338
	}
L334:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+8))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1423)+4))
	m.T0[v1448].(func(*base.Module, int32, int32))(m, v1420, v1447)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L20
	} else {
		goto L336
	}
L335:
	;
	goto L333
L336:
	;
	if v1446 != 0 {
		v1423 = v1446
		goto L334
	} else {
		goto L337
	}
L337:
	;
	goto L335
L338:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L20
	} else {
		goto L339
	}
L339:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v1488 = m.ExcPending
	if v1488 != 0 {
		goto L20
	} else {
		goto L340
	}
L340:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L20
	} else {
		goto L341
	}
L341:
	;
	F_AtEOXact_Inval(m, int32(1))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L20
	} else {
		goto L342
	}
L342:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[40]))
	v1496 = int32(_a_F_CommitTransaction_11)
	v1497 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v1498 = int32(2)
	v1501 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1495+v1497<<(uint(v1498)%32)))) = v1501
	v1504 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[41]))
	v1506 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v1504+v1506<<(uint(v1498)%32)))) = v1501
	v1513 = int32(_a_F_CommitTransaction_28)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[42])) = v1513
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[43])) = v1513
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[44])) = v1501
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[45])) = v1501
	goto L343
L343:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39]))
	v1527 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v1525, int32(2), v1527, v1527)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L20
	} else {
		goto L344
	}
L344:
	;
	v1532 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39]))
	v1534 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v1532, int32(3), v1534, v1534)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L20
	} else {
		goto L345
	}
L345:
	;
	F_smgrDoPendingDeletes(m, int32(1))
	mBase = m.M
	v1540 = m.ExcPending
	if v1540 != 0 {
		goto L20
	} else {
		goto L346
	}
L346:
	;
	v1542 = m.G0
	v1544 = v1542 - int32(48)
	m.G0 = v1544
	v1547 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	v1549 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v1547|v1549 != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[9])))
	if v1552 != int32(1) {
		goto L350
	} else {
		goto L351
	}
L348:
	;
	goto L349
L349:
	;
	m.G0 = v1544 + int32(48)
	v2292 = int32(1)
	F_AtEOXact_GUC(m, v2292, v2292)
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L20
	} else {
		goto L471
	}
L350:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7]))
	if v1571 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L351:
	;
	v1557 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L20
	} else {
		goto L352
	}
L352:
	;
	if v1557 == int32(0) {
		goto L350
	} else {
		goto L353
	}
L353:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_29), int32(0))
	mBase = m.M
	v1564 = m.ExcPending
	if v1564 != 0 {
		goto L20
	} else {
		goto L354
	}
L354:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(979), int32(_a_F_CommitTransaction_29))
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L20
	} else {
		goto L355
	}
L355:
	;
	goto L350
L356:
	;
	v1902 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])))
	if v1902 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L357:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+4))
	if v1574 == int32(0) {
		goto L356
	} else {
		goto L358
	}
L358:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if v1577 <= int32(0) {
		goto L356
	} else {
		goto L359
	}
L359:
	;
	v1582 = int32(0)
	goto L360
L360:
	;
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+12))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1603+v1582<<(uint(int32(2))%32))))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1607)))
	switch v1608 {
	case 0:
		goto L365
	case 1:
		goto L364
	case 2:
		goto L363
	default:
		goto L362
	}
L361:
	;
	goto L356
L362:
	;
	v1875 = v1582 + int32(1)
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v1574)+4))
	if v1875 < v1876 {
		v1582 = v1875
		goto L360
	} else {
		goto L412
	}
L363:
	;
	v1819 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[9])))
	if v1819 != int32(1) {
		goto L405
	} else {
		goto L406
	}
L364:
	;
	v1718 = v1607 + int32(4)
	v1720 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[9])))
	if v1720 != int32(1) {
		goto L382
	} else {
		goto L383
	}
L365:
	;
	v1610 = v1607 + int32(4)
	m.Env.Pgmem_listen(m, v1610, int32(1))
	mBase = m.M
	v1614 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46]))
	if v1614 == int32(0) {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1703 = int32(_a_F_CommitTransaction_30)
	v1704 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[47]))
	v1707 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[48]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[47])) = v1707
	v1709 = F_pstrdup(m, v1610)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L20
	} else {
		goto L380
	}
L367:
	;
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+4))
	if v1617 <= int32(0) {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1614)+12))
	v1622 = int32(0)
	goto L369
L369:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1620+v1622<<(uint(int32(2))%32))))
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648))))
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610))))
	if base.B2i32(v1651 == int32(0))|base.B2i32(v1651 != v1654) != 0 {
		v1672 = v1651
		v1673 = v1654
		goto L372
	} else {
		goto L373
	}
L370:
	;
	goto L366
L371:
	;
	if v1672-v1673 == int32(0) {
		goto L362
	} else {
		goto L378
	}
L372:
	;
	goto L371
L373:
	;
	v1657 = v1648
	v1658 = v1610
	goto L374
L374:
	;
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1658)+1)))
	v1662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1657)+1)))
	if v1662 == int32(0) {
		v1672 = v1662
		v1673 = v1661
		goto L372
	} else {
		goto L376
	}
L375:
	;
	v1672 = v1662
	v1673 = v1661
	goto L372
L376:
	;
	v1665 = int32(1)
	if v1662 == v1661 {
		v1657 = v1657 + v1665
		v1658 = v1658 + v1665
		goto L374
	} else {
		goto L377
	}
L377:
	;
	goto L375
L378:
	;
	v1678 = v1622 + int32(1)
	if v1617 != v1678 {
		v1622 = v1678
		goto L369
	} else {
		goto L379
	}
L379:
	;
	goto L370
L380:
	;
	v1711 = F_lappend(m, v1614, v1709)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L20
	} else {
		goto L381
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[47])) = v1704
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46])) = v1711
	goto L362
L382:
	;
	v1743 = int32(0)
	m.Env.Pgmem_listen(m, v1718, v1743)
	mBase = m.M
	v1746 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46]))
	if v1746 == v1743 {
		goto L362
	} else {
		goto L388
	}
L383:
	;
	v1725 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L20
	} else {
		goto L384
	}
L384:
	;
	if v1725 == int32(0) {
		goto L382
	} else {
		goto L385
	}
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+16)) = v1718
	v1731 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+20)) = v1731
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_31), v1544+int32(16))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L20
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1175), int32(_a_F_CommitTransaction_32))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L20
	} else {
		goto L387
	}
L387:
	;
	goto L382
L388:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+4))
	if v1749 <= int32(0) {
		goto L362
	} else {
		goto L389
	}
L389:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1746)+12))
	v1754 = int32(0)
	goto L390
L390:
	;
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v1752+v1754<<(uint(int32(2))%32))))
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1780))))
	v1786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718))))
	if base.B2i32(v1783 == int32(0))|base.B2i32(v1783 != v1786) != 0 {
		v1804 = v1783
		v1805 = v1786
		goto L393
	} else {
		goto L394
	}
L391:
	;
	goto L362
L392:
	;
	if v1804-v1805 == int32(0) {
		goto L399
	} else {
		goto L400
	}
L393:
	;
	goto L392
L394:
	;
	v1789 = v1780
	v1790 = v1718
	goto L395
L395:
	;
	v1793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1790)+1)))
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1789)+1)))
	if v1794 == int32(0) {
		v1804 = v1794
		v1805 = v1793
		goto L393
	} else {
		goto L397
	}
L396:
	;
	v1804 = v1794
	v1805 = v1793
	goto L393
L397:
	;
	v1797 = int32(1)
	if v1794 == v1793 {
		v1789 = v1789 + v1797
		v1790 = v1790 + v1797
		goto L395
	} else {
		goto L398
	}
L398:
	;
	goto L396
L399:
	;
	v1810 = F_list_delete_nth_cell(m, v1746, v1754)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L20
	} else {
		goto L402
	}
L400:
	;
	goto L401
L401:
	;
	v1816 = v1754 + int32(1)
	if v1749 != v1816 {
		v1754 = v1816
		goto L390
	} else {
		goto L404
	}
L402:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46])) = v1810
	F_pfree(m, v1780)
	mBase = m.M
	v1814 = m.ExcPending
	if v1814 != 0 {
		goto L20
	} else {
		goto L403
	}
L403:
	;
	goto L362
L404:
	;
	goto L391
L405:
	;
	m.Env.Pgmem_listen(m, int32(_a_F_CommitTransaction_33), int32(2))
	mBase = m.M
	v1845 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46]))
	F_list_free_deep(m, v1845)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L20
	} else {
		goto L411
	}
L406:
	;
	v1824 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L20
	} else {
		goto L407
	}
L407:
	;
	if v1824 == int32(0) {
		goto L405
	} else {
		goto L408
	}
L408:
	;
	v1829 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v1544)+32)) = v1829
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_34), v1544+int32(32))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L20
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1205), int32(_a_F_CommitTransaction_35))
	mBase = m.M
	v1840 = m.ExcPending
	if v1840 != 0 {
		goto L20
	} else {
		goto L410
	}
L410:
	;
	goto L405
L411:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46])) = int32(0)
	goto L362
L412:
	;
	goto L361
L413:
	;
	v2043 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8]))
	if v2043 != 0 {
		goto L428
	} else {
		goto L429
	}
L414:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[46]))
	if v1906 != 0 {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v1912 = F_LWLockAcquire(m, v1908+int32(3456), int32(0))
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L20
	} else {
		goto L416
	}
L416:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v1917 = v1915 + int32(56)
	v1918 = int32(_a_F_CommitTransaction_11)
	v1919 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v1920 = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v1917+v1919<<(uint(v1920)%32)))) = int32(-1)
	v1926 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v1915+v1926<<(uint(v1920)%32))+60)) = int32(0)
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1915)+40))
	v1934 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	if v1932 == v1934 {
		goto L418
	} else {
		goto L419
	}
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1915+v1980<<(uint(int32(5))%32)-int32(-64)))) = int32(-1)
	v2011 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v2011+int32(3456))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L20
	} else {
		goto L427
	}
L418:
	;
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1915+v1932<<(uint(int32(5))%32)-int32(-64))))
	*(*int32)(unsafe.Add(mBase, uint32(v1915)+40)) = v1941
	v1980 = v1932
	goto L417
L419:
	;
	goto L420
L420:
	;
	v1943 = v1932
	goto L421
L421:
	;
	if v1943 == int32(-1) {
		goto L423
	} else {
		goto L424
	}
L422:
	;
	v1976 = *(*int32)(unsafe.Add(mBase, uint32(v1917+v1934<<(uint(int32(5))%32))+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1970)+8)) = v1976
	v1979 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[14]))
	v1980 = v1979
	goto L417
L423:
	;
	v1980 = v1934
	goto L417
L424:
	;
	goto L425
L425:
	;
	v1970 = v1917 + v1943<<(uint(int32(5))%32)
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1970)+8))
	if v1971 != v1934 {
		v1943 = v1971
		goto L421
	} else {
		goto L426
	}
L426:
	;
	goto L422
L427:
	;
	v2017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[10])) = uint8(v2017)
	goto L413
L428:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[49]))
	v2048 = F_palloc(m, v2045<<(uint(int32(2))%32))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L20
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	v2254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])))
	if v2254 != 0 {
		goto L467
	} else {
		goto L468
	}
L431:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[49]))
	v2054 = F_palloc(m, v2051<<(uint(int32(2))%32))
	mBase = m.M
	v2055 = m.ExcPending
	if v2055 != 0 {
		goto L20
	} else {
		goto L432
	}
L432:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	v2061 = F_LWLockAcquire(m, v2057+int32(3456), int32(0))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L20
	} else {
		goto L433
	}
L433:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[15]))
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+40))
	if v2065 == int32(-1) {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	F_pfree(m, v2048)
	mBase = m.M
	v2227 = m.ExcPending
	if v2227 != 0 {
		goto L20
	} else {
		goto L465
	}
L435:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v2069+int32(3456))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L20
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v2080 = int32(0)
	v2081 = v2065
	v2092 = v2077
	goto L439
L438:
	;
	goto L434
L439:
	;
	v2103 = v2081 << (uint(int32(5)) % 32)
	v2104 = v2064 + int32(56) + v2103
	v2105 = *(*int64)(unsafe.Add(mBase, uint32(v2104)+16))
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v2104)))
	v2107 = *(*int64)(unsafe.Add(mBase, uint32(v2064)))
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+4))
	if v2092 == v2108 {
		goto L443
	} else {
		goto L444
	}
L440:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[13]))
	F_LWLockRelease(m, v2138+int32(3456))
	mBase = m.M
	v2142 = m.ExcPending
	if v2142 != 0 {
		goto L20
	} else {
		goto L450
	}
L441:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2103+v2064-int32(-64))))
	if v2133 != int32(-1) {
		v2080 = v2128
		v2081 = v2133
		v2092 = v2129
		goto L439
	} else {
		goto L449
	}
L442:
	;
	v2118 = v2080 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v2048+v2118))) = v2106
	*(*int32)(unsafe.Add(mBase, uint32(v2118+v2054))) = v2081
	v2126 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[16]))
	v2128 = v2080 + int32(1)
	v2129 = v2126
	goto L441
L443:
	;
	if v2107 != v2105 {
		goto L442
	} else {
		goto L446
	}
L444:
	;
	goto L445
L445:
	;
	if v2107-v2105 < int64(4) {
		v2128 = v2080
		v2129 = v2092
		goto L441
	} else {
		goto L448
	}
L446:
	;
	v2111 = *(*int32)(unsafe.Add(mBase, uint32(v2104)+24))
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v2064)+8))
	if v2111 != v2112 {
		goto L442
	} else {
		goto L447
	}
L447:
	;
	v2128 = v2080
	v2129 = v2092
	goto L441
L448:
	;
	goto L442
L449:
	;
	goto L440
L450:
	;
	if v2128 <= int32(0) {
		goto L434
	} else {
		goto L451
	}
L451:
	;
	v2145 = int32(0)
	goto L452
L452:
	;
	v2169 = v2145 << (uint(int32(2)) % 32)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v2048+v2169)))
	v2173 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[12]))
	if v2171 == v2173 {
		goto L455
	} else {
		goto L456
	}
L453:
	;
	goto L434
L454:
	;
	v2201 = v2145 + int32(1)
	if v2201 != v2128 {
		v2145 = v2201
		goto L452
	} else {
		goto L464
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[50])) = int32(1)
	goto L454
L456:
	;
	goto L457
L457:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v2169+v2054)))
	v2181 = F_SendProcSignal(m, v2171, int32(1), v2180)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L20
	} else {
		goto L458
	}
L458:
	;
	if int32(0) <= v2181 {
		goto L454
	} else {
		goto L459
	}
L459:
	;
	v2187 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L20
	} else {
		goto L460
	}
L460:
	;
	if v2187 == int32(0) {
		goto L454
	} else {
		goto L461
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1544))) = v2171
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_36), v1544)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L20
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_8), int32(1665), int32(_a_F_CommitTransaction_37))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L20
	} else {
		goto L463
	}
L463:
	;
	goto L454
L464:
	;
	goto L453
L465:
	;
	F_pfree(m, v2054)
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L20
	} else {
		goto L466
	}
L466:
	;
	goto L430
L467:
	;
	v2256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[20])) = uint8(v2256)
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L20
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v2261 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[8])) = v2261
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[7])) = v2261
	goto L349
L470:
	;
	goto L469
L471:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v2298 = m.ExcPending
	if v2298 != 0 {
		goto L20
	} else {
		goto L472
	}
L472:
	;
	v2300 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[51])) = v2300
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[52])) = v2300
	goto L473
L473:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v2307 = m.ExcPending
	if v2307 != 0 {
		goto L20
	} else {
		goto L474
	}
L474:
	;
	F_AtEOXact_Namespace(m, int32(1), v1420)
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L20
	} else {
		goto L475
	}
L475:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L20
	} else {
		goto L476
	}
L476:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L20
	} else {
		goto L477
	}
L477:
	;
	v2317 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[53])) = v2317
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[54])) = v2317
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[55])) = v2317
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[56])) = v2317
	goto L478
L478:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v2330 = m.ExcPending
	if v2330 != 0 {
		goto L20
	} else {
		goto L479
	}
L479:
	;
	F_AtEOXact_PgStat(m, int32(1), v1420)
	mBase = m.M
	v2333 = m.ExcPending
	if v2333 != 0 {
		goto L20
	} else {
		goto L480
	}
L480:
	;
	F_AtEOXact_Snapshot(m, int32(1), int32(0))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L20
	} else {
		goto L481
	}
L481:
	;
	F_AtEOXact_ApplyLauncher(m, int32(1))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L20
	} else {
		goto L482
	}
L482:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(1))
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L20
	} else {
		goto L483
	}
L483:
	;
	v2347 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CommitTransaction[57])))
	if v2347 != int32(1) {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39]))
	F_ResourceOwnerDelete(m, v2377)
	mBase = m.M
	v2379 = m.ExcPending
	if v2379 != 0 {
		goto L20
	} else {
		goto L488
	}
L485:
	;
	goto L484
L486:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[58]))
	if v2351 == int32(0) {
		goto L485
	} else {
		goto L487
	}
L487:
	;
	v2354 = int32(_a_F_CommitTransaction_23)
	v2356 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	v2357 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v2356 + v2357
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v2351)))
	*(*int32)(unsafe.Add(mBase, uint32(v2351))) = v2360 + v2357
	*(*int64)(unsafe.Add(mBase, uint32(v2351)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2351))) = v2360 + int32(2)
	v2371 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[29])) = v2371 - v2357
	goto L485
L488:
	;
	v2380 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+40)) = v2380
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[39])) = v2380
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[59])) = v2380
	v2390 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[0]))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[47])) = v2391
	v2394 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[60]))
	F_MemoryContextReset(m, v2394)
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L20
	} else {
		goto L489
	}
L489:
	;
	v2398 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[61])) = v2398
	*(*int32)(unsafe.Add(mBase, uint32(v2390)+36)) = v2398
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v2398
	v2404 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v2404
	*(*int64)(unsafe.Add(mBase, uint32(v29)+28)) = v2404
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v2398
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v2404
	*(*int64)(unsafe.Add(mBase, _c_F_CommitTransaction[24])) = v2404
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[62])) = v2398
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v2398
	v2420 = int32(_a_F_CommitTransaction_22)
	v2422 = *(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_CommitTransaction[21])) = v2422 - int32(1)
	m.G0 = v26 + int32(48)
	return
L490:
	;
	F_errmsg_internal(m, int32(_a_F_CommitTransaction_38), int32(0))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L20
	} else {
		goto L491
	}
L491:
	;
	F_errfinish(m, int32(_a_F_CommitTransaction_3), int32(1365), int32(_a_F_CommitTransaction_39))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L20
	} else {
		goto L492
	}
L492:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IsTransactionBlock(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_IsTransactionBlock[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+24))
	return base.B2i32(base.Ui32(int32(1)) < base.Ui32(v3))
}
func F_PopTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v11 == int32(0) {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
		if v38 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[0])) = v38
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
			*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[1])) = v42
			*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[2])) = v42
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
			*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[3])) = v47
			*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[4])) = v47
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			if v51 != 0 {
				F_pfree(m, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_PopTransaction_0), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PopTransaction_1), int32(_a_F_PopTransaction_2), int32(_a_F_PopTransaction_3))
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
		v16 = F_errstart(m, int32(19), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 == int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				if v38 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[0])) = v38
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
					*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[1])) = v42
					*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[2])) = v42
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
					*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[3])) = v47
					*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[4])) = v47
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v51 != 0 {
						F_pfree(m, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							F_pfree(m, v10)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					} else {
						F_pfree(m, v10)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_PopTransaction_0), int32(0))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PopTransaction_1), int32(_a_F_PopTransaction_2), int32(_a_F_PopTransaction_3))
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
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
				if base.Ui32(v20) <= base.Ui32(int32(5)) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20<<(uint(int32(2))%32))+uint32(_c_F_PopTransaction[5])))
					v27 = v25
				} else {
					v27 = int32(_a_F_PopTransaction_4)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
				F_errmsg_internal(m, int32(_a_F_PopTransaction_5), v7)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PopTransaction_1), int32(_a_F_PopTransaction_6), int32(_a_F_PopTransaction_3))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
						if v38 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[0])) = v38
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
							*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[1])) = v42
							*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[2])) = v42
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
							*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[3])) = v47
							*(*int32)(unsafe.Add(mBase, _c_F_PopTransaction[4])) = v47
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							if v51 != 0 {
								F_pfree(m, v51)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									F_pfree(m, v10)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								F_pfree(m, v10)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_PopTransaction_0), int32(0))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_PopTransaction_1), int32(_a_F_PopTransaction_2), int32(_a_F_PopTransaction_3))
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
					}
				}
			}
		}
	}
}
func F_PrepareTransaction(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
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
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v556 int32
	_ = v556
	var v558 int64
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int64
	_ = v605
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1152 int32
	_ = v1152
	var v1162 int32
	_ = v1162
	var v1164 int32
	_ = v1164
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1235 int32
	_ = v1235
	var v1238 int64
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1435 int32
	_ = v1435
	var v1458 int32
	_ = v1458
	var v1461 int64
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1578 int32
	_ = v1578
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1593 int32
	_ = v1593
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1685 int32
	_ = v1685
	var v1692 int32
	_ = v1692
	var v1721 int64
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1733 int64
	_ = v1733
	var v1735 int64
	_ = v1735
	var v1744 int64
	_ = v1744
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1755 int64
	_ = v1755
	var v1757 int64
	_ = v1757
	var v1766 int64
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1779 int32
	_ = v1779
	var v1784 int32
	_ = v1784
	var v1785 int64
	_ = v1785
	var v1786 int64
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1794 int64
	_ = v1794
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1845 int64
	_ = v1845
	var v1851 int32
	_ = v1851
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1907 int32
	_ = v1907
	var v1938 int32
	_ = v1938
	var v1940 int64
	_ = v1940
	var v1942 int64
	_ = v1942
	var v1944 int32
	_ = v1944
	var v1951 int32
	_ = v1951
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	var v2058 int32
	_ = v2058
	var v2062 int32
	_ = v2062
	var v2066 int32
	_ = v2066
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2112 int32
	_ = v2112
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2120 int32
	_ = v2120
	var v2129 int32
	_ = v2129
	var v2154 int32
	_ = v2154
	var v2155 int64
	_ = v2155
	var v2157 int64
	_ = v2157
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2197 int32
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2270 int32
	_ = v2270
	var v2271 int64
	_ = v2271
	var v2273 int64
	_ = v2273
	var v2275 int64
	_ = v2275
	var v2277 int64
	_ = v2277
	var v2279 int64
	_ = v2279
	var v2281 int64
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2368 int32
	_ = v2368
	var v2375 int32
	_ = v2375
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2392 int32
	_ = v2392
	var v2395 int32
	_ = v2395
	var v2399 int32
	_ = v2399
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2410 int32
	_ = v2410
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2427 int32
	_ = v2427
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2435 int32
	_ = v2435
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2468 int32
	_ = v2468
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2479 int64
	_ = v2479
	var v2482 int64
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2571 int32
	_ = v2571
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2578 int64
	_ = v2578
	var v2579 int32
	_ = v2579
	var v2586 int64
	_ = v2586
	var v2588 int32
	_ = v2588
	var v2589 int64
	_ = v2589
	var v2590 int64
	_ = v2590
	var v2592 int32
	_ = v2592
	var v2594 int64
	_ = v2594
	var v2597 int32
	_ = v2597
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2606 int32
	_ = v2606
	var v2610 int32
	_ = v2610
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2619 int32
	_ = v2619
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2628 int32
	_ = v2628
	var v2630 int32
	_ = v2630
	var v2634 int64
	_ = v2634
	var v2637 int32
	_ = v2637
	var v2639 int32
	_ = v2639
	var v2650 int32
	_ = v2650
	var v2653 int32
	_ = v2653
	var v2657 int32
	_ = v2657
	var v2662 int32
	_ = v2662
	var v2666 int32
	_ = v2666
	var v2668 int32
	_ = v2668
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2693 int32
	_ = v2693
	var v2716 int32
	_ = v2716
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2740 int32
	_ = v2740
	var v2743 int32
	_ = v2743
	var v2744 int32
	_ = v2744
	var v2750 int32
	_ = v2750
	var v2755 int32
	_ = v2755
	var v2772 int32
	_ = v2772
	var v2774 int32
	_ = v2774
	var v2777 int32
	_ = v2777
	var v2778 int32
	_ = v2778
	var v2782 int32
	_ = v2782
	var v2788 int32
	_ = v2788
	var v2792 int32
	_ = v2792
	var v2796 int32
	_ = v2796
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2806 int32
	_ = v2806
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2817 int32
	_ = v2817
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2874 int32
	_ = v2874
	var v2877 int32
	_ = v2877
	var v2878 int32
	_ = v2878
	var v2881 int32
	_ = v2881
	var v2883 int32
	_ = v2883
	var v2890 int32
	_ = v2890
	var v2918 int64
	_ = v2918
	var v2921 int32
	_ = v2921
	var v2923 int32
	_ = v2923
	var v2957 int32
	_ = v2957
	var v2989 int32
	_ = v2989
	var v2990 int32
	_ = v2990
	var v3023 int32
	_ = v3023
	var v3025 int32
	_ = v3025
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3039 int32
	_ = v3039
	var v3057 int32
	_ = v3057
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3063 int32
	_ = v3063
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3074 int32
	_ = v3074
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3119 int32
	_ = v3119
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3128 int32
	_ = v3128
	var v3136 int32
	_ = v3136
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3141 int32
	_ = v3141
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3150 int32
	_ = v3150
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3154 int32
	_ = v3154
	var v3156 int32
	_ = v3156
	var v3157 int32
	_ = v3157
	var v3158 int32
	_ = v3158
	var v3163 int32
	_ = v3163
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3176 int32
	_ = v3176
	var v3204 int32
	_ = v3204
	var v3207 int32
	_ = v3207
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3213 int32
	_ = v3213
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3228 int32
	_ = v3228
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3238 int32
	_ = v3238
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3244 int32
	_ = v3244
	var v3270 int32
	_ = v3270
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3277 int32
	_ = v3277
	var v3279 int32
	_ = v3279
	var v3308 int32
	_ = v3308
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3316 int32
	_ = v3316
	var v3347 int32
	_ = v3347
	var v3354 int32
	_ = v3354
	var v3355 int32
	_ = v3355
	var v3359 int32
	_ = v3359
	var v3364 int32
	_ = v3364
	var v3395 int32
	_ = v3395
	var v3399 int32
	_ = v3399
	var v3400 int32
	_ = v3400
	var v3406 int32
	_ = v3406
	var v3411 int32
	_ = v3411
	var v3414 int32
	_ = v3414
	var v3420 int32
	_ = v3420
	var v3484 int32
	_ = v3484
	var v3486 int32
	_ = v3486
	var v3488 int32
	_ = v3488
	var v3496 int32
	_ = v3496
	var v3497 int32
	_ = v3497
	var v3519 int32
	_ = v3519
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	var v3534 int32
	_ = v3534
	var v3537 int32
	_ = v3537
	var v3541 int32
	_ = v3541
	var v3546 int32
	_ = v3546
	var v3550 int32
	_ = v3550
	var v3554 int32
	_ = v3554
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3567 int32
	_ = v3567
	var v3572 int32
	_ = v3572
	var v3574 int32
	_ = v3574
	var v3576 int32
	_ = v3576
	var v3580 int32
	_ = v3580
	var v3581 int32
	_ = v3581
	var v3583 int32
	_ = v3583
	var v3584 int32
	_ = v3584
	var v3585 int32
	_ = v3585
	var v3589 int32
	_ = v3589
	var v3598 int32
	_ = v3598
	var v3599 int64
	_ = v3599
	var v3603 int32
	_ = v3603
	var v3606 int32
	_ = v3606
	var v3610 int32
	_ = v3610
	var v3611 int32
	_ = v3611
	var v3612 int32
	_ = v3612
	var v3613 int32
	_ = v3613
	var v3615 int32
	_ = v3615
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3627 int32
	_ = v3627
	var v3631 int32
	_ = v3631
	var v3633 int32
	_ = v3633
	var v3635 int32
	_ = v3635
	var v3663 int32
	_ = v3663
	var v3665 int32
	_ = v3665
	var v3666 int32
	_ = v3666
	var v3668 int32
	_ = v3668
	var v3699 int32
	_ = v3699
	var v3700 int32
	_ = v3700
	var v3704 int32
	_ = v3704
	var v3706 int32
	_ = v3706
	var v3709 int32
	_ = v3709
	var v3711 int32
	_ = v3711
	var v3713 int32
	_ = v3713
	var v3714 int32
	_ = v3714
	var v3716 int32
	_ = v3716
	var v3744 int32
	_ = v3744
	var v3747 int32
	_ = v3747
	var v3810 int32
	_ = v3810
	var v3815 int32
	_ = v3815
	var v3816 int32
	_ = v3816
	var v3817 int32
	_ = v3817
	var v3820 int32
	_ = v3820
	var v3849 int32
	_ = v3849
	var v3854 int32
	_ = v3854
	var v3856 int32
	_ = v3856
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3891 int32
	_ = v3891
	var v3920 int32
	_ = v3920
	var v3925 int32
	_ = v3925
	var v3927 int32
	_ = v3927
	var v3991 int32
	_ = v3991
	var v3993 int32
	_ = v3993
	var v4022 int32
	_ = v4022
	var v4025 int32
	_ = v4025
	var v4056 int32
	_ = v4056
	var v4058 int32
	_ = v4058
	var v4062 int32
	_ = v4062
	var v4064 int32
	_ = v4064
	var v4065 int32
	_ = v4065
	var v4067 int32
	_ = v4067
	var v4071 int32
	_ = v4071
	var v4072 int32
	_ = v4072
	var v4074 int32
	_ = v4074
	var v4075 int32
	_ = v4075
	var v4080 int32
	_ = v4080
	var v4087 int32
	_ = v4087
	var v4091 int32
	_ = v4091
	var v4093 int32
	_ = v4093
	var v4095 int32
	_ = v4095
	var v4097 int32
	_ = v4097
	var v4101 int32
	_ = v4101
	var v4104 int32
	_ = v4104
	var v4116 int32
	_ = v4116
	var v4120 int32
	_ = v4120
	var v4122 int32
	_ = v4122
	var v4124 int32
	_ = v4124
	var v4133 int32
	_ = v4133
	var v4135 int32
	_ = v4135
	var v4138 int32
	_ = v4138
	var v4140 int32
	_ = v4140
	var v4142 int32
	_ = v4142
	var v4145 int32
	_ = v4145
	var v4147 int32
	_ = v4147
	var v4151 int32
	_ = v4151
	var v4152 int32
	_ = v4152
	var v4154 int32
	_ = v4154
	var v4158 int32
	_ = v4158
	var v4162 int32
	_ = v4162
	var v4166 int32
	_ = v4166
	var v4169 int32
	_ = v4169
	var v4172 int32
	_ = v4172
	var v4174 int32
	_ = v4174
	var v4181 int32
	_ = v4181
	var v4185 int32
	_ = v4185
	var v4187 int32
	_ = v4187
	var v4190 int32
	_ = v4190
	var v4192 int32
	_ = v4192
	var v4205 int32
	_ = v4205
	var v4206 int32
	_ = v4206
	var v4209 int32
	_ = v4209
	var v4212 int32
	_ = v4212
	var v4215 int32
	_ = v4215
	var v4219 int32
	_ = v4219
	var v4223 int32
	_ = v4223
	var v4226 int32
	_ = v4226
	var v4228 int32
	_ = v4228
	var v4229 int32
	_ = v4229
	var v4232 int32
	_ = v4232
	var v4243 int32
	_ = v4243
	var v4252 int32
	_ = v4252
	var v4254 int32
	_ = v4254
	var v4255 int32
	_ = v4255
	var v4265 int32
	_ = v4265
	var v4266 int32
	_ = v4266
	var v4269 int32
	_ = v4269
	var v4271 int32
	_ = v4271
	var v4273 int32
	_ = v4273
	var v4279 int64
	_ = v4279
	var v4295 int32
	_ = v4295
	var v4297 int32
	_ = v4297
	var v4307 int32
	_ = v4307
	var v4310 int32
	_ = v4310
	var v4314 int32
	_ = v4314
	var v4319 int32
	_ = v4319
	var v4323 int32
	_ = v4323
	var v4326 int32
	_ = v4326
	var v4330 int32
	_ = v4330
	var v4335 int32
	_ = v4335
	v30 = m.G0
	v32 = v30 - int32(16)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v36 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v35))))
	if v36 == int64(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_AssignTransactionId(m, v35)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v36
	goto L3
L3:
	;
	v43 = int32(10)
	goto L8
L4:
	;
	return
L5:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v42 = v41
	goto L3
L6:
	;
	if v80 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	goto L6
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[1]))
	goto L11
L9:
	;
	v63 = int32(0)
	goto L16
L11:
	;
	goto L12
L12:
	;
	if int32(0)|base.B2i32(v50 == int32(15)) != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v50 <= v43 {
		v80 = int32(1)
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[2]))
	if v67 != int32(2) {
		v80 = v63
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[3])))
	if v71&int32(1) != 0 {
		v80 = v63
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[4]))
	v80 = int32(0) | base.B2i32(v77 <= v43)
	goto L7
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	F_ShowTransactionStateRec(m, int32(_a_F_PrepareTransaction_0), v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v87 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v114 = base.I32_wrap_i64(v42)
	goto L32
L24:
	;
	v92 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v92 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if base.Ui32(v96) <= base.Ui32(int32(5)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96<<(uint(int32(2))%32))+uint32(_c_F_PrepareTransaction[5])))
	v103 = v101
	goto L29
L28:
	;
	v103 = int32(_a_F_PrepareTransaction_1)
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v103
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_2), v32)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2531), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	goto L23
L32:
	;
	F_AfterTriggerFireDeferred(m)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L34
	}
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[6]))
	if v150 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v147 = F_PreCommit_Portals(m, int32(1))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v147 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v152 = v150
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_AfterTriggerEndXact(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L44
	}
L40:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	m.T0[v183].(func(*base.Module, int32, int32))(m, int32(7), v182)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	if v180 != 0 {
		v152 = v180
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_PreCommit_on_commit_actions(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_smgrDoPendingSyncs(m, int32(1), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_AtEOXact_LargeObject(m, int32(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_PreCommit_CheckForSerializationFailure(m)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[7])))
	if v229&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4323 = m.ExcPending
	if v4323 != 0 {
		goto L4
	} else {
		goto L648
	}
L50:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[8]))
	if v235 != 0 {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v4307 = m.ExcPending
	if v4307 != 0 {
		goto L4
	} else {
		goto L644
	}
L53:
	;
	v236 = int32(_a_F_PrepareTransaction_4)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[9])) = v238 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(5)
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[10]))
	if int32(0) < v245 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_disable_timeout(m, int32(8))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v254 = m.G0
	v255 = int32(16)
	v256 = v254 - v255
	m.G0 = v256
	F_gettimeofday(m, v256)
	mBase = m.M
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v256)))
	v260 = int64(*(*int32)(unsafe.Add(mBase, uint32(v256)+8)))
	m.G0 = v256 + v255
	goto L58
L57:
	;
	goto L56
L58:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[11]))
	v272 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[12]))
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[13]))
	v276 = m.G0
	v278 = v276 - int32(48)
	m.G0 = v278
	v280 = F_strlen(m, v270)
	mBase = m.M
	if base.Ui32(v280) < base.Ui32(int32(200)) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v511 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[11])) = v511
	v513 = m.G0
	v515 = v513 - int32(96)
	m.G0 = v515
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v423)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+8)) = v511
	*(*int32)(unsafe.Add(mBase, uint32(v515)+4)) = v511
	v528 = F_palloc0(m, int32(12))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L4
	} else {
		goto L106
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L4
	} else {
		goto L101
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L4
	} else {
		goto L96
	}
L62:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	if v284 == int32(0) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L4
	} else {
		goto L92
	}
L65:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[16])))
	if v288 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_before_shmem_exit(m, int32(393), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v303 = F_LWLockAcquire(m, v299+int32(2304), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L70
	}
L69:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[16])) = uint8(v296)
	goto L68
L70:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+4))
	if v307 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v423 == int32(0) {
		goto L60
	} else {
		goto L90
	}
L72:
	;
	v313 = int32(0)
	goto L73
L73:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v306+int32(8)+v313<<(uint(int32(2))%32))))
	v346 = v344 + int32(47)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	if base.B2i32(v349 == int32(0))|base.B2i32(v349 != v352) != 0 {
		v370 = v349
		v371 = v352
		goto L76
	} else {
		goto L77
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L4
	} else {
		goto L86
	}
L75:
	;
	if v370-v371 != 0 {
		goto L82
	} else {
		goto L83
	}
L76:
	;
	goto L75
L77:
	;
	v355 = v346
	v356 = v270
	goto L78
L78:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+1)))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355)+1)))
	if v360 == int32(0) {
		v370 = v360
		v371 = v359
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v370 = v360
	v371 = v359
	goto L76
L80:
	;
	v363 = int32(1)
	if v360 == v359 {
		v355 = v355 + v363
		v356 = v356 + v363
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v374 = v313 + int32(1)
	if v307 != v374 {
		v313 = v374
		goto L73
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L74
L85:
	;
	goto L71
L86:
	;
	F_errcode(m, int32(_a_F_PrepareTransaction_5))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278)+16)) = v270
	F_errmsg(m, int32(_a_F_PrepareTransaction_6), v278+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(396), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
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
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = v426
	F_MarkAsPreparingGuts(m, v423, v114, v270, v260+v259*int64(1000000)-int64(946684800000000), v272, v274)
	mBase = m.M
	v429 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v423)+45)) = uint8(v429)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[18]))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v433 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v432+v433<<(uint(int32(2))%32))+8)) = v423
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v442+int32(2304))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v278 + int32(48)
	goto L59
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v278))) = v270
	F_errmsg(m, int32(_a_F_PrepareTransaction_9), v278)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(369), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_10), int32(0))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errhint(m, int32(_a_F_PrepareTransaction_11), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(376), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(_a_F_PrepareTransaction_12))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_13), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v278)+32)) = v498
	F_errhint(m, int32(_a_F_PrepareTransaction_14), v278+int32(32))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(406), int32(_a_F_PrepareTransaction_8))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19])) = v528
	*(*int64)(unsafe.Add(mBase, uint32(v528)+4)) = int64(0)
	v534 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v534
	v537 = F_palloc(m, v534)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19]))
	*(*int32)(unsafe.Add(mBase, uint32(v540))) = v537
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v540
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v515)+32)) = v521
	*(*int64)(unsafe.Add(mBase, uint32(v515)+24)) = int64(1475953972)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v519+v520*int32(640))+60))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+36)) = v556
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v423)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v515)+40)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v423)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v515)+48)) = v560
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+52))
	if v566 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+52)) = v571
	v576 = F_smgrGetPendingDeletes(m, int32(1), v515+int32(16))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L4
	} else {
		goto L112
	}
L109:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v565)+48))
	v569 = v567
	goto L111
L110:
	;
	v569 = int32(0)
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515+int32(20)))) = v569
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v565)+52))
	goto L108
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+56)) = v576
	v582 = F_smgrGetPendingDeletes(m, int32(0), v515+int32(12))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+60)) = v582
	v588 = F_pgstat_get_transactional_drops(m, int32(1), v515+int32(4))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+64)) = v588
	v594 = F_pgstat_get_transactional_drops(m, int32(0), v515+int32(8))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+68)) = v594
	v599 = F_xactGetCommittedInvalidationMessages(m, v515, v515+int32(76))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v515)+72)) = v599
	v603 = v423 + int32(47)
	v604 = F_strlen(m, v603)
	mBase = m.M
	v605 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v515)+80)) = v605
	*(*int64)(unsafe.Add(mBase, uint32(v515)+88)) = v605
	v610 = v604 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v515)+78)) = uint16(v610)
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(int32(72)) <= base.Ui32(v613) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	v654 = int32(72)
	base.MemoryCopy(m, v647+v650, v515+int32(24), v654)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v648)+4))
	v658 = v656 + v654
	*(*int32)(unsafe.Add(mBase, uint32(v648)+4)) = v658
	v662 = v649 - v654
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v662
	v664 = int32(_a_F_PrepareTransaction_15)
	v666 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v668 = v666 + v654
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v668
	v670 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v515)+78)))
	v674 = (v670 + int32(7)) & int32(_a_F_PrepareTransaction_16)
	if base.Ui32(v674) <= base.Ui32(v662) {
		goto L124
	} else {
		goto L125
	}
L118:
	;
	v617 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	v647 = v618
	v648 = v617
	v649 = v613
	goto L117
L119:
	;
	goto L120
L120:
	;
	v620 = F_palloc0(m, int32(12))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L4
	} else {
		goto L121
	}
L121:
	;
	v622 = int32(_a_F_PrepareTransaction_17)
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v623)+8)) = v620
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v620
	*(*int64)(unsafe.Add(mBase, uint32(v620)+4)) = int64(0)
	v630 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v630
	v632 = int32(_a_F_PrepareTransaction_18)
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v634 + int32(1)
	v639 = F_palloc(m, v630)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v642 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v642))) = v639
	v645 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v647 = v639
	v648 = v642
	v649 = v645
	goto L117
L123:
	;
	if v670 != 0 {
		goto L132
	} else {
		goto L133
	}
L124:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v648)))
	v710 = v662
	v711 = v648
	v712 = v668
	v713 = v658
	v714 = v676
	goto L123
L125:
	;
	goto L126
L126:
	;
	v678 = F_palloc0(m, int32(12))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v680 = int32(_a_F_PrepareTransaction_17)
	v681 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v681)+8)) = v678
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v678
	*(*int64)(unsafe.Add(mBase, uint32(v678)+4)) = int64(0)
	v688 = int32(512)
	if base.Ui32(v674) <= base.Ui32(v688) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v691 = v688
	goto L130
L129:
	;
	v691 = v674
	goto L130
L130:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v691
	v693 = int32(_a_F_PrepareTransaction_18)
	v695 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v695 + int32(1)
	v699 = F_palloc(m, v691)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v699
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v702)+4))
	v706 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v708 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v710 = v708
	v711 = v702
	v712 = v706
	v713 = v704
	v714 = v699
	goto L123
L132:
	;
	base.MemoryCopy(m, v713+v714, v603, v670)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v711)+4))
	v718 = v717 + v674
	*(*int32)(unsafe.Add(mBase, uint32(v711)+4)) = v718
	v721 = v674 + v712
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v721
	v724 = v710 - v674
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v724
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v515)+52))
	if v726 <= int32(0) {
		v814 = v724
		v816 = v721
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v515)+56))
	if int32(0) < v818 {
		goto L157
	} else {
		goto L158
	}
L136:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v515)+20))
	v731 = v726 << (uint(int32(2)) % 32)
	v735 = (v731 + int32(7)) & int32(-8)
	if base.Ui32(v735) <= base.Ui32(v724) {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v731 != 0 {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v711)))
	v769 = v711
	v770 = v724
	v771 = v718
	v772 = v737
	goto L137
L139:
	;
	goto L140
L140:
	;
	v739 = F_palloc0(m, int32(12))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v741 = int32(_a_F_PrepareTransaction_17)
	v742 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v742)+8)) = v739
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v739
	*(*int64)(unsafe.Add(mBase, uint32(v739)+4)) = int64(0)
	v749 = int32(512)
	if base.Ui32(v735) <= base.Ui32(v749) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v752 = v749
	goto L144
L143:
	;
	v752 = v735
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v752
	v754 = int32(_a_F_PrepareTransaction_18)
	v756 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v756 + int32(1)
	v760 = F_palloc(m, v752)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L145
	}
L145:
	;
	v763 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v763))) = v760
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v763)+4))
	v767 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v769 = v763
	v770 = v767
	v771 = v765
	v772 = v760
	goto L137
L146:
	;
	base.MemoryCopy(m, v771+v772, v729, v731)
	goto L148
L147:
	;
	goto L148
L148:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v769)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v769)+4)) = v775 + v735
	v779 = v770 - v735
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v779
	v781 = int32(_a_F_PrepareTransaction_15)
	v783 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v784 = v783 + v735
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v784
	v787 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v787)))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	v792 = v788 + v789*int32(640)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v515)+20))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v515)+52))
	if int32(65) <= v794 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v804 = v802 << (uint(int32(2)) % 32)
	if v804 != 0 {
		goto L154
	} else {
		goto L155
	}
L150:
	;
	v797 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v792)+277)) = uint8(v797)
	v802 = int32(64)
	goto L149
L151:
	;
	goto L152
L152:
	;
	if v794 <= int32(0) {
		v814 = v779
		v816 = v784
		goto L135
	} else {
		goto L153
	}
L153:
	;
	v802 = v794
	goto L149
L154:
	;
	base.MemoryCopy(m, v792+int32(280), v793, v804)
	goto L156
L155:
	;
	goto L156
L156:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v792)+276)) = uint8(v802)
	v814 = v779
	v816 = v784
	goto L135
L157:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v515)+16))
	v823 = v818 * int32(12)
	v827 = (v823 + int32(7)) & int32(-8)
	if base.Ui32(v827) <= base.Ui32(v814) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v515)+60))
	if int32(0) < v889 {
		goto L173
	} else {
		goto L174
	}
L160:
	;
	if v823 != 0 {
		goto L169
	} else {
		goto L170
	}
L161:
	;
	v830 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v830)))
	v863 = v830
	v864 = v814
	v865 = v831
	v866 = v816
	goto L160
L162:
	;
	goto L163
L163:
	;
	v833 = F_palloc0(m, int32(12))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	v835 = int32(_a_F_PrepareTransaction_17)
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v836)+8)) = v833
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v833
	*(*int64)(unsafe.Add(mBase, uint32(v833)+4)) = int64(0)
	v843 = int32(512)
	if base.Ui32(v827) <= base.Ui32(v843) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v846 = v843
	goto L167
L166:
	;
	v846 = v827
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v846
	v848 = int32(_a_F_PrepareTransaction_18)
	v850 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v850 + int32(1)
	v854 = F_palloc(m, v846)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v857 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v857))) = v854
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v863 = v857
	v864 = v862
	v865 = v854
	v866 = v860
	goto L160
L169:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	base.MemoryCopy(m, v865+v867, v821, v823)
	goto L171
L170:
	;
	goto L171
L171:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v863)+4)) = v870 + v827
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v827 + v866
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v864 - v827
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v515)+16))
	F_pfree(m, v879)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	goto L159
L173:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	v894 = v889 * int32(12)
	v898 = (v894 + int32(7)) & int32(-8)
	v900 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(v898) <= base.Ui32(v900) {
		goto L177
	} else {
		goto L178
	}
L174:
	;
	goto L175
L175:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v515)+64))
	if int32(0) < v960 {
		goto L189
	} else {
		goto L190
	}
L176:
	;
	if v894 != 0 {
		goto L185
	} else {
		goto L186
	}
L177:
	;
	v903 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v903)))
	v934 = v903
	v935 = v900
	v936 = v904
	goto L176
L178:
	;
	goto L179
L179:
	;
	v906 = F_palloc0(m, int32(12))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v908 = int32(_a_F_PrepareTransaction_17)
	v909 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v909)+8)) = v906
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v906
	*(*int64)(unsafe.Add(mBase, uint32(v906)+4)) = int64(0)
	v916 = int32(512)
	if base.Ui32(v898) <= base.Ui32(v916) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v919 = v916
	goto L183
L182:
	;
	v919 = v898
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v919
	v921 = int32(_a_F_PrepareTransaction_18)
	v923 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v923 + int32(1)
	v927 = F_palloc(m, v919)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	v930 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v927
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v934 = v930
	v935 = v933
	v936 = v927
	goto L176
L185:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	base.MemoryCopy(m, v936+v937, v892, v894)
	goto L187
L186:
	;
	goto L187
L187:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v934)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v934)+4)) = v940 + v898
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v935 - v898
	v946 = int32(_a_F_PrepareTransaction_15)
	v948 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v948 + v898
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v515)+12))
	F_pfree(m, v951)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	goto L175
L189:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	v965 = v960 << (uint(int32(4)) % 32)
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(v965) <= base.Ui32(v967) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v515)+68))
	if int32(0) < v1026 {
		goto L205
	} else {
		goto L206
	}
L192:
	;
	if v965 != 0 {
		goto L201
	} else {
		goto L202
	}
L193:
	;
	v970 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v970)))
	v1001 = v970
	v1002 = v971
	v1003 = v967
	goto L192
L194:
	;
	goto L195
L195:
	;
	v973 = F_palloc0(m, int32(12))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v975 = int32(_a_F_PrepareTransaction_17)
	v976 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v976)+8)) = v973
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v973
	*(*int64)(unsafe.Add(mBase, uint32(v973)+4)) = int64(0)
	v983 = int32(512)
	if base.Ui32(v965) <= base.Ui32(v983) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v986 = v983
	goto L199
L198:
	;
	v986 = v965
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v986
	v988 = int32(_a_F_PrepareTransaction_18)
	v990 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v990 + int32(1)
	v994 = F_palloc(m, v986)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v997 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v997))) = v994
	v1000 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v1001 = v997
	v1002 = v994
	v1003 = v1000
	goto L192
L201:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	base.MemoryCopy(m, v1002+v1004, v963, v965)
	goto L203
L202:
	;
	goto L203
L203:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1001)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1001)+4)) = v1007 + v965
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v1003 - v965
	v1013 = int32(_a_F_PrepareTransaction_15)
	v1015 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v1015 + v965
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v515)+4))
	F_pfree(m, v1018)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	goto L191
L205:
	;
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	v1031 = v1026 << (uint(int32(4)) % 32)
	v1033 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(v1031) <= base.Ui32(v1033) {
		goto L209
	} else {
		goto L210
	}
L206:
	;
	goto L207
L207:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v515)+72))
	if int32(0) < v1092 {
		goto L221
	} else {
		goto L222
	}
L208:
	;
	if v1031 != 0 {
		goto L217
	} else {
		goto L218
	}
L209:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	v1067 = v1036
	v1068 = v1037
	v1069 = v1033
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1039 = F_palloc0(m, int32(12))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	v1041 = int32(_a_F_PrepareTransaction_17)
	v1042 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1042)+8)) = v1039
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v1039)+4)) = int64(0)
	v1049 = int32(512)
	if base.Ui32(v1031) <= base.Ui32(v1049) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1052 = v1049
	goto L215
L214:
	;
	v1052 = v1031
	goto L215
L215:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v1052
	v1054 = int32(_a_F_PrepareTransaction_18)
	v1056 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1056 + int32(1)
	v1060 = F_palloc(m, v1052)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L4
	} else {
		goto L216
	}
L216:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1063))) = v1060
	v1066 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v1067 = v1063
	v1068 = v1060
	v1069 = v1066
	goto L208
L217:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+4))
	base.MemoryCopy(m, v1068+v1070, v1029, v1031)
	goto L219
L218:
	;
	goto L219
L219:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1067)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1067)+4)) = v1073 + v1031
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v1069 - v1031
	v1079 = int32(_a_F_PrepareTransaction_15)
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v1081 + v1031
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v515)+8))
	F_pfree(m, v1084)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L4
	} else {
		goto L220
	}
L220:
	;
	goto L207
L221:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v1097 = v1092 << (uint(int32(4)) % 32)
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(v1097) <= base.Ui32(v1099) {
		goto L225
	} else {
		goto L226
	}
L222:
	;
	goto L223
L223:
	;
	m.G0 = v515 + int32(96)
	v1162 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[24]))
	v1164 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[25]))
	if v1162|v1164 != 0 {
		goto L237
	} else {
		goto L238
	}
L224:
	;
	if v1097 != 0 {
		goto L233
	} else {
		goto L234
	}
L225:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1102)))
	v1133 = v1102
	v1134 = v1103
	v1135 = v1099
	goto L224
L226:
	;
	goto L227
L227:
	;
	v1105 = F_palloc0(m, int32(12))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1107 = int32(_a_F_PrepareTransaction_17)
	v1108 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+8)) = v1105
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v1105
	*(*int64)(unsafe.Add(mBase, uint32(v1105)+4)) = int64(0)
	v1115 = int32(512)
	if base.Ui32(v1097) <= base.Ui32(v1115) {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1118 = v1115
	goto L231
L230:
	;
	v1118 = v1097
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v1118
	v1120 = int32(_a_F_PrepareTransaction_18)
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v1122 + int32(1)
	v1126 = F_palloc(m, v1118)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1129))) = v1126
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v1133 = v1129
	v1134 = v1126
	v1135 = v1132
	goto L224
L233:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	base.MemoryCopy(m, v1134+v1136, v1095, v1097)
	goto L235
L234:
	;
	goto L235
L235:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1133)+4)) = v1139 + v1097
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v1135 - v1097
	v1145 = int32(_a_F_PrepareTransaction_15)
	v1147 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v1147 + v1097
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	F_pfree(m, v1150)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	goto L223
L237:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L4
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1182 = m.G0
	v1184 = v1182 - int32(80)
	m.G0 = v1184
	*(*int64)(unsafe.Add(mBase, uint32(v1184)+48)) = int64(85899345936)
	v1189 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+72)) = v1189
	v1196 = F_hash_create(m, int32(_a_F_PrepareTransaction_19), int32(256), v1184+int32(32), int32(1064))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L4
	} else {
		goto L244
	}
L240:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L4
	} else {
		goto L241
	}
L241:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_20), int32(0))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L4
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_21), int32(841), int32(_a_F_PrepareTransaction_22))
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	v1199 = v1184 + int32(8)
	v1201 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[27]))
	F_hash_seq_init(m, v1199, v1201)
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	v1204 = F_hash_seq_search(m, v1199)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L4
	} else {
		goto L248
	}
L246:
	;
	v2088 = m.G0
	v2090 = v2088 - int32(32)
	m.G0 = v2090
	v2093 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28]))
	if v2093 != 0 {
		goto L374
	} else {
		goto L375
	}
L247:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L4
	} else {
		goto L370
	}
L248:
	;
	if v1204 != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1206 = v1204
	goto L252
L250:
	;
	goto L251
L251:
	;
	F_hash_destroy(m, v1196)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L4
	} else {
		goto L286
	}
L252:
	;
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1206)+14)))
	if v1235 == int32(6) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	goto L251
L254:
	;
	v1388 = F_hash_seq_search(m, v1184+int32(8))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L4
	} else {
		goto L284
	}
L255:
	;
	v1238 = *(*int64)(unsafe.Add(mBase, uint32(v1206)+32))
	if v1238 <= int64(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+48))
	v1245 = F_hash_search(m, v1196, v1206, int32(1), v1184+int32(7))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1184)+7)))
	if v1247 == int32(0) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1250 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1245)+16)) = uint16(v1250)
	goto L260
L259:
	;
	goto L260
L260:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1206)+40))
	v1254 = v1252 - int32(1)
	if v1254 < int32(0) {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245)+16)))
	if v1351 != int32(1) {
		goto L254
	} else {
		goto L282
	}
L262:
	;
	if v1252&int32(1) != 0 {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1241+v1254<<(uint(int32(4))%32))))
	if v1262 != 0 {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1269 = v1254
	goto L265
L265:
	;
	if v1254 == int32(0) {
		goto L261
	} else {
		goto L270
	}
L266:
	;
	v1269 = v1252 - int32(2)
	goto L265
L267:
	;
	v1263 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+17)) = uint8(v1263)
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+16)) = uint8(v1265)
	goto L266
L270:
	;
	v1272 = v1269
	goto L271
L271:
	;
	v1303 = v1241 + v1272<<(uint(int32(4))%32)
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1303)))
	if v1304 == int32(0) {
		goto L274
	} else {
		goto L275
	}
L272:
	;
	goto L261
L273:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1303-int32(16))))
	if v1313 != 0 {
		goto L278
	} else {
		goto L279
	}
L274:
	;
	v1307 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+16)) = uint8(v1307)
	goto L273
L275:
	;
	goto L276
L276:
	;
	v1309 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+17)) = uint8(v1309)
	goto L273
L277:
	;
	if int32(1) < v1272 {
		v1272 = v1272 - int32(2)
		goto L271
	} else {
		goto L281
	}
L278:
	;
	v1314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+17)) = uint8(v1314)
	goto L277
L279:
	;
	goto L280
L280:
	;
	v1316 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1245)+16)) = uint8(v1316)
	goto L277
L281:
	;
	goto L272
L282:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1245)+17)))
	if v1354 == int32(1) {
		goto L247
	} else {
		goto L283
	}
L283:
	;
	goto L254
L284:
	;
	if v1388 != 0 {
		v1206 = v1388
		goto L252
	} else {
		goto L285
	}
L285:
	;
	goto L253
L286:
	;
	v1422 = v1184 + int32(32)
	v1424 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[27]))
	F_hash_seq_init(m, v1422, v1424)
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L4
	} else {
		goto L287
	}
L287:
	;
	v1427 = F_hash_seq_search(m, v1422)
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L4
	} else {
		goto L291
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2062 = m.ExcPending
	if v2062 != 0 {
		goto L4
	} else {
		goto L367
	}
L289:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
		goto L4
	} else {
		goto L364
	}
L290:
	;
	F_LWLockRelease(m, v1692)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L4
	} else {
		goto L357
	}
L291:
	;
	if v1427 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1435 = v1427
	goto L295
L293:
	;
	goto L294
L294:
	;
	m.G0 = v1184 + int32(80)
	goto L246
L295:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1435)+14)))
	if v1458 == int32(6) {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	goto L294
L297:
	;
	v1983 = F_hash_seq_search(m, v1184+int32(32))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L4
	} else {
		goto L355
	}
L298:
	;
	v1461 = *(*int64)(unsafe.Add(mBase, uint32(v1435)+32))
	if v1461 <= int64(0) {
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+40))
	v1466 = v1464 - int32(1)
	if v1466 < int32(0) {
		goto L297
	} else {
		goto L300
	}
L300:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+48))
	v1470 = int32(3)
	v1471 = v1464 & v1470
	if base.Ui32(v1466) < base.Ui32(v1470) {
		goto L303
	} else {
		goto L304
	}
L301:
	;
	if v1630&int32(1) == int32(0) {
		goto L297
	} else {
		goto L313
	}
L302:
	;
	v1578 = v1549
	v1585 = v1556
	v1586 = v1557
	v1593 = int32(0)
	goto L310
L303:
	;
	v1475 = int32(0)
	v1549 = v1466
	v1556 = v1475
	v1557 = v1475
	goto L302
L304:
	;
	goto L305
L305:
	;
	v1479 = int32(0)
	v1482 = v1466
	v1489 = v1479
	v1490 = v1479
	v1498 = v1479
	goto L306
L306:
	;
	v1511 = int32(4)
	v1513 = v1469 + v1482<<(uint(v1511)%32)
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1513-int32(48))))
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v1513-int32(32))))
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1513-int32(16))))
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1513)))
	v1527 = int32(0)
	v1529 = base.B2i32(v1516|v1519|v1523|v1525 != v1527) | v1490
	v1541 = base.B2i32(v1525 == v1527) | (base.B2i32(v1523 == v1527) | (base.B2i32(v1516 == v1527) | base.B2i32(v1519 == v1527))) | v1489
	v1543 = v1482 - v1511
	v1545 = v1498 + v1511
	if v1545 != v1464&int32(-4) {
		v1482 = v1543
		v1489 = v1541
		v1490 = v1529
		v1498 = v1545
		goto L306
	} else {
		goto L308
	}
L307:
	;
	if v1471 == int32(0) {
		v1629 = v1541
		v1630 = v1529
		goto L301
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v1549 = v1543
	v1556 = v1541
	v1557 = v1529
	goto L302
L310:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v1469+v1578<<(uint(int32(4))%32))))
	v1611 = int32(0)
	v1613 = base.B2i32(v1610 != v1611) | v1586
	v1616 = base.B2i32(v1610 == v1611) | v1585
	v1617 = int32(1)
	v1620 = v1593 + v1617
	if v1620 != v1471 {
		v1578 = v1578 - v1617
		v1585 = v1616
		v1586 = v1613
		v1593 = v1620
		goto L310
	} else {
		goto L312
	}
L311:
	;
	v1629 = v1616
	v1630 = v1613
	goto L301
L312:
	;
	goto L311
L313:
	;
	if v1629&int32(1) == int32(0) {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	v1938 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1435)+52)) = uint8(v1938)
	v1940 = *(*int64)(unsafe.Add(mBase, uint32(v1435)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v1184)+16)) = v1940
	v1942 = *(*int64)(unsafe.Add(mBase, uint32(v1435)))
	*(*int64)(unsafe.Add(mBase, uint32(v1184)+8)) = v1942
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+24)) = v1944
	F_RegisterTwoPhaseRecord(m, int32(1), v1184+int32(8), int32(20))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L4
	} else {
		goto L354
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1435)+28)) = v1898
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	*(*int32)(unsafe.Add(mBase, uint32(v1435)+24)) = v1907
	goto L314
L316:
	;
	F_LWLockRelease(m, v1681+int32(584))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L4
	} else {
		goto L347
	}
L317:
	;
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+16))
	v1794 = int64(1) << (uint(base.I64_extend_i32_u(v1788+base.I32_wrap_i64(v1785)-int32(1))) % 64)
	if v1794&v1786 == int64(0) {
		goto L316
	} else {
		goto L338
	}
L318:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+28))
	if v1659 != 0 {
		goto L314
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L4
	} else {
		goto L334
	}
L321:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+20))
	v1664 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[29]))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+4))
	v1667 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v1671 = F_LWLockAcquire(m, v1667+int32(584), int32(0))
	mBase = m.M
	v1672 = m.ExcPending
	if v1672 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	v1673 = int32(268435455)
	v1677 = (v1664 + v1673) & (v1665 * int32(_a_F_PrepareTransaction_23))
	v1679 = v1677 & v1673
	v1681 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+604))
	v1685 = v1682 + v1677<<(uint(int32(6))%32)
	v1692 = v1661 + v1662&int32(15)<<(uint(int32(7))%32) + int32(_a_F_PrepareTransaction_24)
	v1721 = int64(0)
	goto L323
L323:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1685+base.I32_wrap_i64(v1721)<<(uint(int32(2))%32))))
	if v1727 == v1665 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	goto L316
L325:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+600))
	v1733 = *(*int64)(unsafe.Add(mBase, uint32(v1729+v1679<<(uint(int32(3))%32))))
	v1735 = v1721 * int64(3)
	if int64(base.Ui64(v1733)>>(uint(v1735)%64))&int64(7) != int64(0) {
		v1785 = v1735
		v1786 = v1733
		goto L317
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1744 = v1721 | int64(1)
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1685+base.I32_wrap_i64(v1744)<<(uint(int32(2))%32))))
	if v1749 == v1665 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	goto L327
L329:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1681)+600))
	v1755 = *(*int64)(unsafe.Add(mBase, uint32(v1751+v1679<<(uint(int32(3))%32))))
	v1757 = v1744 * int64(3)
	if int64(base.Ui64(v1755)>>(uint(v1757)%64))&int64(7) != int64(0) {
		v1785 = v1757
		v1786 = v1755
		goto L317
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	v1766 = v1721 + int64(2)
	if v1766 != int64(16) {
		v1721 = v1766
		goto L323
	} else {
		goto L333
	}
L332:
	;
	goto L331
L333:
	;
	goto L324
L334:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L4
	} else {
		goto L335
	}
L335:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v1779 = m.ExcPending
	if v1779 != 0 {
		goto L4
	} else {
		goto L336
	}
L336:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3494), int32(_a_F_PrepareTransaction_27))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L4
	} else {
		goto L337
	}
L337:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L338:
	;
	v1799 = F_LWLockAcquire(m, v1692, int32(0))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L4
	} else {
		goto L339
	}
L339:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+20))
	v1805 = F_SetupLockInTable(m, int32(_a_F_PrepareTransaction_28), v1803, v1435, v1804, v1788)
	mBase = m.M
	v1806 = m.ExcPending
	if v1806 != 0 {
		goto L4
	} else {
		goto L340
	}
L340:
	;
	if v1805 == int32(0) {
		goto L290
	} else {
		goto L341
	}
L341:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(v1805)))
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+128))
	v1811 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1809)+128)) = v1810 + v1811
	v1816 = v1809 + v1788<<(uint(int32(2))%32)
	v1818 = v1816 + int32(88)
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	*(*int32)(unsafe.Add(mBase, uint32(v1818))) = v1819 + v1811
	v1824 = v1811 << (uint(v1788) % 32)
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1809)+16)) = v1824 | v1825
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(v1818)))
	v1829 = *(*int32)(unsafe.Add(mBase, uint32(v1816)+44))
	if v1828 == v1829 {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1831 = *(*int32)(unsafe.Add(mBase, uint32(v1809)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1809)+20)) = v1831 & (v1824 ^ int32(-1))
	goto L344
L343:
	;
	goto L344
L344:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1805)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1805)+12)) = v1836 | v1824
	v1840 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v1840)+600))
	v1844 = v1841 + v1679<<(uint(int32(3))%32)
	v1845 = *(*int64)(unsafe.Add(mBase, uint32(v1844)))
	*(*int64)(unsafe.Add(mBase, uint32(v1844))) = v1845 & (v1794 ^ int64(-1))
	F_LWLockRelease(m, v1692)
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L4
	} else {
		goto L345
	}
L345:
	;
	v1853 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	F_LWLockRelease(m, v1853+int32(584))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L4
	} else {
		goto L346
	}
L346:
	;
	v1898 = v1805
	goto L315
L347:
	;
	v1867 = F_LWLockAcquire(m, v1692, int32(1))
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L4
	} else {
		goto L348
	}
L348:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[31]))
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+20))
	v1872 = int32(0)
	v1874 = F_hash_search_with_hash_value(m, v1870, v1435, v1871, v1872, v1872)
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L4
	} else {
		goto L349
	}
L349:
	;
	if v1874 == int32(0) {
		goto L289
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+8)) = v1874
	v1880 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1184)+12)) = v1880
	v1883 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[32]))
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1435)+20))
	v1890 = int32(0)
	v1892 = F_hash_search_with_hash_value(m, v1883, v1184+int32(8), v1886^v1880<<(uint(int32(4))%32), v1890, v1890)
	mBase = m.M
	v1893 = m.ExcPending
	if v1893 != 0 {
		goto L4
	} else {
		goto L351
	}
L351:
	;
	if v1892 == int32(0) {
		goto L288
	} else {
		goto L352
	}
L352:
	;
	F_LWLockRelease(m, v1692)
	mBase = m.M
	v1897 = m.ExcPending
	if v1897 != 0 {
		goto L4
	} else {
		goto L353
	}
L353:
	;
	v1898 = v1892
	goto L315
L354:
	;
	goto L297
L355:
	;
	if v1983 != 0 {
		v1435 = v1983
		goto L295
	} else {
		goto L356
	}
L356:
	;
	goto L296
L357:
	;
	v2020 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	F_LWLockRelease(m, v2020+int32(584))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L4
	} else {
		goto L358
	}
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L4
	} else {
		goto L359
	}
L359:
	;
	F_errcode(m, int32(_a_F_PrepareTransaction_12))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L4
	} else {
		goto L360
	}
L360:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_29), int32(0))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L4
	} else {
		goto L361
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1184))) = int32(_a_F_PrepareTransaction_30)
	F_errhint(m, int32(_a_F_PrepareTransaction_31), v1184)
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L4
	} else {
		goto L362
	}
L362:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(2969), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2045 = m.ExcPending
	if v2045 != 0 {
		goto L4
	} else {
		goto L363
	}
L363:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L364:
	;
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_33), int32(0))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(2997), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L367:
	;
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_34), int32(0))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3010), int32(_a_F_PrepareTransaction_32))
	mBase = m.M
	v2071 = m.ExcPending
	if v2071 != 0 {
		goto L4
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L4
	} else {
		goto L371
	}
L371:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L4
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3426), int32(_a_F_PrepareTransaction_35))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L4
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2090)+8)) = int32(0)
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v2093)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v2090)+12)) = v2096
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v2093)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v2090)+16)) = v2098
	F_RegisterTwoPhaseRecord(m, int32(4), v2090+int32(8), int32(24))
	mBase = m.M
	v2105 = m.ExcPending
	if v2105 != 0 {
		goto L4
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	m.G0 = v2090 + int32(32)
	v2235 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[33]))
	if v2235 != 0 {
		goto L387
	} else {
		goto L388
	}
L377:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v2111 = F_LWLockAcquire(m, v2107+int32(3840), int32(1))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L4
	} else {
		goto L378
	}
L378:
	;
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v2093)+52))
	if v2113 == int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v2197+int32(3840))
	mBase = m.M
	v2201 = m.ExcPending
	if v2201 != 0 {
		goto L4
	} else {
		goto L386
	}
L380:
	;
	v2117 = v2093 + int32(48)
	if v2113 == v2117 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2120 = v2090 + int32(12)
	v2129 = v2113
	goto L382
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2090)+8)) = int32(1)
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v2129-int32(16))))
	v2155 = *(*int64)(unsafe.Add(mBase, uint32(v2154)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2120)+8)) = v2155
	v2157 = *(*int64)(unsafe.Add(mBase, uint32(v2154)))
	*(*int64)(unsafe.Add(mBase, uint32(v2120))) = v2157
	F_RegisterTwoPhaseRecord(m, int32(4), v2090+int32(8), int32(24))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L4
	} else {
		goto L384
	}
L383:
	;
	goto L379
L384:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2129)+4))
	if v2165 != v2117 {
		v2129 = v2165
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	goto L376
L387:
	;
	v2236 = m.G0
	v2238 = v2236 + int32(-64)
	m.G0 = v2238
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2235)+20))
	if v2240 != 0 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	goto L389
L389:
	;
	v2357 = m.G0
	v2359 = v2357 - int32(16)
	m.G0 = v2359
	v2362 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34]))
	v2364 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(v2362+v2364<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v2359)+12)) = v2368
	if v2368 != 0 {
		goto L397
	} else {
		goto L398
	}
L390:
	;
	v2241 = v2240
	goto L393
L391:
	;
	goto L392
L392:
	;
	m.G0 = v2238 - int32(-64)
	goto L389
L393:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+64))
	v2271 = *(*int64)(unsafe.Add(mBase, uint32(v2241)))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+8)) = v2271
	v2273 = *(*int64)(unsafe.Add(mBase, uint32(v2241)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+16)) = v2273
	v2275 = *(*int64)(unsafe.Add(mBase, uint32(v2241)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+24)) = v2275
	v2277 = *(*int64)(unsafe.Add(mBase, uint32(v2241)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+32)) = v2277
	v2279 = *(*int64)(unsafe.Add(mBase, uint32(v2241)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+40)) = v2279
	v2281 = *(*int64)(unsafe.Add(mBase, uint32(v2241)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v2238)+48)) = v2281
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v2270)))
	*(*int32)(unsafe.Add(mBase, uint32(v2238)+56)) = v2283
	v2285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2270)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2238)+60)) = uint8(v2285)
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2241)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2238)+61)) = uint8(v2287)
	F_RegisterTwoPhaseRecord(m, int32(2), v2236+int32(-56), int32(56))
	mBase = m.M
	v2294 = m.ExcPending
	if v2294 != 0 {
		goto L4
	} else {
		goto L395
	}
L394:
	;
	goto L392
L395:
	;
	v2295 = *(*int32)(unsafe.Add(mBase, uint32(v2241)+68))
	if v2295 != 0 {
		v2241 = v2295
		goto L393
	} else {
		goto L396
	}
L396:
	;
	goto L394
L397:
	;
	F_RegisterTwoPhaseRecord(m, int32(3), v2359+int32(12), int32(4))
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L4
	} else {
		goto L400
	}
L398:
	;
	goto L399
L399:
	;
	m.G0 = v2359 + int32(16)
	v2380 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[36]))
	if v2380 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L399
L401:
	;
	v2406 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	if base.Ui32(int32(8)) <= base.Ui32(v2406) {
		goto L412
	} else {
		goto L413
	}
L402:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2392 = m.ExcPending
	if v2392 != 0 {
		goto L4
	} else {
		goto L407
	}
L403:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[37]))
	if v2382 != 0 {
		goto L402
	} else {
		goto L404
	}
L404:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[38]))
	if v2384 != 0 {
		goto L402
	} else {
		goto L405
	}
L405:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[39]))
	if v2386 == int32(0) {
		goto L401
	} else {
		goto L406
	}
L406:
	;
	goto L402
L407:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2395 = m.ExcPending
	if v2395 != 0 {
		goto L4
	} else {
		goto L408
	}
L408:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_36), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L4
	} else {
		goto L409
	}
L409:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_37), int32(596), int32(_a_F_PrepareTransaction_38))
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L4
	} else {
		goto L410
	}
L410:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L411:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+4))
	v2443 = v2440 + v2442
	v2444 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2443)+6)) = uint16(v2444)
	*(*uint8)(unsafe.Add(mBase, uint32(v2443)+4)) = uint8(v2444)
	*(*int32)(unsafe.Add(mBase, uint32(v2443))) = v2444
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2439)+4))
	v2451 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v2439)+4)) = v2450 + v2451
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v2441 - v2451
	v2458 = int32(_a_F_PrepareTransaction_15)
	v2460 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[23])) = v2460 + v2451
	v2465 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19]))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2465)))
	v2468 = v2460 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v2466)+4)) = v2468
	v2471 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_PrepareTransaction[40])))
	v2473 = v2471 - int32(1)
	if base.Ui32(v2473&int32(_a_F_PrepareTransaction_39)) <= base.Ui32(int32(_a_F_PrepareTransaction_40)) {
		goto L417
	} else {
		goto L418
	}
L412:
	;
	v2410 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2410)))
	v2439 = v2410
	v2440 = v2411
	v2441 = v2406
	goto L411
L413:
	;
	goto L414
L414:
	;
	v2413 = F_palloc0(m, int32(12))
	mBase = m.M
	v2414 = m.ExcPending
	if v2414 != 0 {
		goto L4
	} else {
		goto L415
	}
L415:
	;
	v2415 = int32(_a_F_PrepareTransaction_17)
	v2416 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v2416)+8)) = v2413
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v2413
	*(*int64)(unsafe.Add(mBase, uint32(v2413)+4)) = int64(0)
	v2423 = int32(512)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20])) = v2423
	v2425 = int32(_a_F_PrepareTransaction_18)
	v2427 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v2427 + int32(1)
	v2432 = F_palloc(m, v2423)
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L4
	} else {
		goto L416
	}
L416:
	;
	v2435 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v2435))) = v2432
	v2438 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[20]))
	v2439 = v2435
	v2440 = v2432
	v2441 = v2438
	goto L411
L417:
	;
	v2479 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[41]))
	*(*int64)(unsafe.Add(mBase, uint32(v2466)+56)) = v2479
	v2482 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[42]))
	*(*int64)(unsafe.Add(mBase, uint32(v2466)+64)) = v2482
	goto L419
L418:
	;
	goto L419
L419:
	;
	if base.Ui32(v2468) < base.Ui32(int32(1073741824)) {
		goto L421
	} else {
		goto L422
	}
L420:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[43])) = int64(0)
	v2666 = m.G0
	v2668 = v2666 - int32(32)
	m.G0 = v2668
	v2671 = F_TwoPhaseGetDummyProc(m, v114, int32(0))
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L4
	} else {
		goto L448
	}
L421:
	;
	v2488 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21]))
	F_XLogEnsureRecordSpace(m, int32(0), v2488)
	mBase = m.M
	v2490 = m.ExcPending
	if v2490 != 0 {
		goto L4
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L4
	} else {
		goto L444
	}
L424:
	;
	v2491 = int32(_a_F_PrepareTransaction_41)
	v2493 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	v2494 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v2493 + v2494
	v2498 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v2498)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v2498)+120)) = v2499 | v2494
	F_XLogBeginInsert(m)
	mBase = m.M
	v2504 = m.ExcPending
	if v2504 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	v2506 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19]))
	if v2506 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2507 = v2506
	goto L429
L427:
	;
	goto L428
L428:
	;
	v2571 = int32(_a_F_PrepareTransaction_42)
	v2573 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])))
	v2574 = v2573 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[45])) = uint8(v2574)
	goto L433
L429:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(v2507)))
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+4))
	F_XLogRegisterData(m, v2536, v2537)
	mBase = m.M
	v2539 = m.ExcPending
	if v2539 != 0 {
		goto L4
	} else {
		goto L431
	}
L430:
	;
	goto L428
L431:
	;
	v2540 = *(*int32)(unsafe.Add(mBase, uint32(v2507)+8))
	if v2540 != 0 {
		v2507 = v2540
		goto L429
	} else {
		goto L432
	}
L432:
	;
	goto L430
L433:
	;
	v2578 = F_XLogInsert(m, int32(1), int32(16))
	mBase = m.M
	v2579 = m.ExcPending
	if v2579 != 0 {
		goto L4
	} else {
		goto L434
	}
L434:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v423)+24)) = v2578
	if base.Ui32(v2473&int32(_a_F_PrepareTransaction_39)) <= base.Ui32(int32(_a_F_PrepareTransaction_40)) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2586 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[41]))
	F_replorigin_session_advance(m, v2586, v2578)
	mBase = m.M
	v2588 = m.ExcPending
	if v2588 != 0 {
		goto L4
	} else {
		goto L438
	}
L436:
	;
	v2590 = v2578
	goto L437
L437:
	;
	F_XLogFlush(m, v2590)
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L4
	} else {
		goto L439
	}
L438:
	;
	v2589 = *(*int64)(unsafe.Add(mBase, uint32(v423)+24))
	v2590 = v2589
	goto L437
L439:
	;
	v2594 = *(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[46]))
	*(*int64)(unsafe.Add(mBase, uint32(v423)+16)) = v2594
	v2597 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v2601 = F_LWLockAcquire(m, v2597+int32(2304), int32(0))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L4
	} else {
		goto L440
	}
L440:
	;
	v2603 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v423)+44)) = uint8(v2603)
	v2606 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v2606+int32(2304))
	mBase = m.M
	v2610 = m.ExcPending
	if v2610 != 0 {
		goto L4
	} else {
		goto L441
	}
L441:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2612)))
	v2614 = *(*int32)(unsafe.Add(mBase, uint32(v423)+4))
	F_ProcArrayAdd(m, v2613+v2614*int32(640))
	mBase = m.M
	v2619 = m.ExcPending
	if v2619 != 0 {
		goto L4
	} else {
		goto L442
	}
L442:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2621)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v2621)+120)) = v2622 & int32(-2)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[47])) = v423
	v2628 = int32(_a_F_PrepareTransaction_41)
	v2630 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v2630 - int32(1)
	v2634 = *(*int64)(unsafe.Add(mBase, uint32(v423)+24))
	F_SyncRepWaitForLSN(m, v2634, int32(0))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L4
	} else {
		goto L443
	}
L443:
	;
	v2639 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[22])) = v2639
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[19])) = v2639
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[21])) = v2639
	goto L420
L444:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v2653 = m.ExcPending
	if v2653 != 0 {
		goto L4
	} else {
		goto L445
	}
L445:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_43), int32(0))
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L4
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_7), int32(1174), int32(_a_F_PrepareTransaction_44))
	mBase = m.M
	v2662 = m.ExcPending
	if v2662 != 0 {
		goto L4
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	v2673 = int32(_a_F_PrepareTransaction_41)
	v2675 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v2675 + int32(1)
	v2680 = v2668 + int32(12)
	v2682 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[27]))
	F_hash_seq_init(m, v2680, v2682)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L4
	} else {
		goto L449
	}
L449:
	;
	v2685 = F_hash_seq_search(m, v2680)
	mBase = m.M
	v2686 = m.ExcPending
	if v2686 != 0 {
		goto L4
	} else {
		goto L454
	}
L450:
	;
	v3574 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v3576 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v3580 = F_LWLockAcquire(m, v3576+int32(512), int32(0))
	mBase = m.M
	v3581 = m.ExcPending
	if v3581 != 0 {
		goto L4
	} else {
		goto L560
	}
L451:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3563 = m.ExcPending
	if v3563 != 0 {
		goto L4
	} else {
		goto L557
	}
L452:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3550 = m.ExcPending
	if v3550 != 0 {
		goto L4
	} else {
		goto L554
	}
L453:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3534 = m.ExcPending
	if v3534 != 0 {
		goto L4
	} else {
		goto L550
	}
L454:
	;
	if v2685 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v2693 = v2685
	goto L458
L456:
	;
	goto L457
L457:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v3025 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v3034 = v3025
	v3035 = v3023
	v3039 = int32(0)
	goto L484
L458:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+28))
	if v2716 == int32(0) {
		goto L461
	} else {
		goto L462
	}
L459:
	;
	goto L457
L460:
	;
	v2989 = F_hash_seq_search(m, v2668+int32(12))
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L4
	} else {
		goto L482
	}
L461:
	;
	F_RemoveLocalLock(m, v2693)
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L4
	} else {
		goto L481
	}
L462:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+24))
	if v2719 == int32(0) {
		goto L461
	} else {
		goto L463
	}
L463:
	;
	v2722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2693)+14)))
	if v2722 == int32(6) {
		goto L460
	} else {
		goto L464
	}
L464:
	;
	v2725 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+40))
	v2727 = v2725 - int32(1)
	if v2727 < int32(0) {
		goto L460
	} else {
		goto L465
	}
L465:
	;
	v2730 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+48))
	v2731 = int32(3)
	v2732 = v2725 & v2731
	if base.Ui32(v2727) < base.Ui32(v2731) {
		goto L468
	} else {
		goto L469
	}
L466:
	;
	if v2883&int32(1) == int32(0) {
		goto L460
	} else {
		goto L478
	}
L467:
	;
	v2839 = v2810
	v2840 = v2811
	v2846 = v2817
	v2847 = int32(0)
	goto L475
L468:
	;
	v2736 = int32(0)
	v2810 = v2736
	v2811 = v2727
	v2817 = v2736
	goto L467
L469:
	;
	goto L470
L470:
	;
	v2740 = int32(0)
	v2743 = v2740
	v2744 = v2727
	v2750 = v2740
	v2755 = v2740
	goto L471
L471:
	;
	v2772 = int32(4)
	v2774 = v2730 + v2744<<(uint(v2772)%32)
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2774-int32(48))))
	v2778 = int32(0)
	v2782 = *(*int32)(unsafe.Add(mBase, uint32(v2774-int32(32))))
	v2788 = *(*int32)(unsafe.Add(mBase, uint32(v2774-int32(16))))
	v2792 = *(*int32)(unsafe.Add(mBase, uint32(v2774)))
	v2796 = base.B2i32(v2777 == v2778) | base.B2i32(v2782 == v2778) | base.B2i32(v2788 == v2778) | base.B2i32(v2792 == v2778) | v2750
	v2802 = base.B2i32(v2782|v2777|v2788|v2792 != v2778) | v2743
	v2804 = v2744 - v2772
	v2806 = v2755 + v2772
	if v2806 != v2725&int32(-4) {
		v2743 = v2802
		v2744 = v2804
		v2750 = v2796
		v2755 = v2806
		goto L471
	} else {
		goto L473
	}
L472:
	;
	if v2732 == int32(0) {
		v2883 = v2802
		v2890 = v2796
		goto L466
	} else {
		goto L474
	}
L473:
	;
	goto L472
L474:
	;
	v2810 = v2802
	v2811 = v2804
	v2817 = v2796
	goto L467
L475:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2730+v2840<<(uint(int32(4))%32))))
	v2872 = int32(0)
	v2874 = base.B2i32(v2871 == v2872) | v2846
	v2877 = base.B2i32(v2871 != v2872) | v2839
	v2878 = int32(1)
	v2881 = v2847 + v2878
	if v2881 != v2732 {
		v2839 = v2877
		v2840 = v2840 - v2878
		v2846 = v2874
		v2847 = v2881
		goto L475
	} else {
		goto L477
	}
L476:
	;
	v2883 = v2877
	v2890 = v2874
	goto L466
L477:
	;
	goto L476
L478:
	;
	if v2890&int32(1) != 0 {
		goto L453
	} else {
		goto L479
	}
L479:
	;
	v2918 = *(*int64)(unsafe.Add(mBase, uint32(v2693)+32))
	if v2918 <= int64(0) {
		goto L461
	} else {
		goto L480
	}
L480:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2716)+16))
	v2923 = *(*int32)(unsafe.Add(mBase, uint32(v2693)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v2716)+16)) = v2921 | int32(1)<<(uint(v2923)%32)
	goto L461
L481:
	;
	goto L460
L482:
	;
	if v2989 != 0 {
		v2693 = v2989
		goto L458
	} else {
		goto L483
	}
L483:
	;
	goto L459
L484:
	;
	v3057 = v3039 << (uint(int32(3)) % 32)
	v3058 = v3034 + v3057
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v3058)+152))
	if v3059 == int32(0) {
		v3496 = v3034
		v3497 = v3035
		goto L486
	} else {
		goto L487
	}
L485:
	;
	v3522 = int32(_a_F_PrepareTransaction_41)
	v3524 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v3524 - int32(1)
	m.G0 = v2668 + int32(32)
	goto L450
L486:
	;
	v3519 = v3039 + int32(1)
	if v3519 != int32(16) {
		v3034 = v3496
		v3035 = v3497
		v3039 = v3519
		goto L484
	} else {
		goto L549
	}
L487:
	;
	v3063 = v3058 + int32(148)
	if v3059 == v3063 {
		v3496 = v3034
		v3497 = v3035
		goto L486
	} else {
		goto L488
	}
L488:
	;
	v3069 = v3035 + v3039<<(uint(int32(7))%32) + int32(_a_F_PrepareTransaction_24)
	v3071 = F_LWLockAcquire(m, v3069, int32(0))
	mBase = m.M
	v3072 = m.ExcPending
	if v3072 != 0 {
		goto L4
	} else {
		goto L489
	}
L489:
	;
	v3073 = *(*int32)(unsafe.Add(mBase, uint32(v3063)+4))
	v3074 = int32(0)
	if base.B2i32(v3073 == v3074)|base.B2i32(v3073 == v3063) == v3074 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v3080 = v3057 + (v2671 + int32(148))
	v3081 = v3073
	goto L493
L491:
	;
	goto L492
L492:
	;
	F_LWLockRelease(m, v3069)
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L4
	} else {
		goto L548
	}
L493:
	;
	v3110 = *(*int32)(unsafe.Add(mBase, uint32(v3081)+4))
	v3112 = v3081 - int32(28)
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v3112)))
	v3114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3113)+14)))
	if v3114 == int32(6) {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	goto L492
L495:
	;
	if v3110 != v3063 {
		v3081 = v3110
		goto L493
	} else {
		goto L547
	}
L496:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(v3081-int32(12))))
	if v3119 == int32(0) {
		goto L495
	} else {
		goto L497
	}
L497:
	;
	v3124 = *(*int32)(unsafe.Add(mBase, uint32(v3081-int32(16))))
	if v3119 != v3124 {
		goto L452
	} else {
		goto L498
	}
L498:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(v3081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3126)+4)) = v3110
	v3128 = *(*int32)(unsafe.Add(mBase, uint32(v3081)))
	*(*int32)(unsafe.Add(mBase, uint32(v3110))) = v3128
	*(*int32)(unsafe.Add(mBase, uint32(v2668)+8)) = v2671
	*(*int32)(unsafe.Add(mBase, uint32(v2668)+4)) = v3113
	*(*int32)(unsafe.Add(mBase, uint32(v3081-int32(20)))) = v2671
	v3136 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[32]))
	v3138 = v2668 + int32(4)
	v3139 = m.G0
	v3141 = v3139 - int32(32)
	m.G0 = v3141
	v3143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3136)+34)))
	if v3143 != int32(1) {
		goto L502
	} else {
		goto L503
	}
L499:
	;
	if v3347 == int32(0) {
		goto L451
	} else {
		goto L543
	}
L500:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3399 = m.ExcPending
	if v3399 != 0 {
		goto L4
	} else {
		goto L540
	}
L501:
	;
	F_hash_corrupted(m, v3136)
	mBase = m.M
	v3395 = m.ExcPending
	if v3395 != 0 {
		goto L4
	} else {
		goto L539
	}
L502:
	;
	v3146 = *(*int32)(unsafe.Add(mBase, uint32(v3136)))
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v3146)+396))
	v3149 = v3112 - int32(4)
	v3150 = *(*int32)(unsafe.Add(mBase, uint32(v3149)))
	v3151 = v3147 & v3150
	v3152 = *(*int32)(unsafe.Add(mBase, uint32(v3146)+392))
	if base.Ui32(v3152) < base.Ui32(v3151) {
		goto L505
	} else {
		goto L506
	}
L503:
	;
	goto L504
L504:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L4
	} else {
		goto L536
	}
L505:
	;
	v3154 = *(*int32)(unsafe.Add(mBase, uint32(v3146)+400))
	v3156 = v3154 & v3151
	goto L507
L506:
	;
	v3156 = v3151
	goto L507
L507:
	;
	v3157 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+4))
	v3158 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+44))
	v3163 = *(*int32)(unsafe.Add(mBase, uint32(v3157+int32(base.Ui32(v3156)>>(uint(v3158)%32))<<(uint(int32(2))%32))))
	if v3163 == int32(0) {
		goto L501
	} else {
		goto L508
	}
L508:
	;
	v3167 = v3112 - int32(8)
	v3168 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+40))
	v3176 = v3163 + (v3168-int32(1))&v3156<<(uint(int32(2))%32)
	goto L509
L509:
	;
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(v3176)))
	if v3204 != v3167 {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	if v3204 == int32(0) {
		goto L500
	} else {
		goto L515
	}
L511:
	;
	v3207 = v3204
	goto L513
L512:
	;
	v3207 = int32(0)
	goto L513
L513:
	;
	if v3207 != 0 {
		v3176 = v3204
		goto L509
	} else {
		goto L514
	}
L514:
	;
	goto L510
L515:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+36))
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+8))
	v3212 = m.T0[v3211].(func(*base.Module, int32, int32) int32)(m, v3138, v3210)
	mBase = m.M
	v3213 = m.ExcPending
	if v3213 != 0 {
		goto L4
	} else {
		goto L516
	}
L516:
	;
	v3214 = *(*int32)(unsafe.Add(mBase, uint32(v3136)))
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v3214)+396))
	v3216 = v3212 & v3215
	v3217 = *(*int32)(unsafe.Add(mBase, uint32(v3214)+392))
	if base.Ui32(v3217) < base.Ui32(v3216) {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(v3214)+400))
	v3221 = v3219 & v3216
	goto L519
L518:
	;
	v3221 = v3216
	goto L519
L519:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+4))
	v3223 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+44))
	v3228 = *(*int32)(unsafe.Add(mBase, uint32(v3222+int32(base.Ui32(v3221)>>(uint(v3223)%32))<<(uint(int32(2))%32))))
	if v3228 == int32(0) {
		goto L501
	} else {
		goto L520
	}
L520:
	;
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+36))
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+40))
	v3238 = v3228 + (v3232-int32(1))&v3221<<(uint(int32(2))%32)
	v3239 = *(*int32)(unsafe.Add(mBase, uint32(v3238)))
	if v3239 != 0 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	m.G0 = v3141 + int32(32)
	goto L499
L522:
	;
	v3240 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+12))
	v3244 = v3239
	goto L525
L523:
	;
	v3279 = v3238
	goto L524
L524:
	;
	if v3221 != v3156 {
		goto L532
	} else {
		goto L533
	}
L525:
	;
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(v3244)+4))
	if v3270 != v3212 {
		goto L527
	} else {
		goto L528
	}
L526:
	;
	v3279 = v3244
	goto L524
L527:
	;
	v3277 = *(*int32)(unsafe.Add(mBase, uint32(v3244)))
	if v3277 != 0 {
		v3244 = v3277
		goto L525
	} else {
		goto L531
	}
L528:
	;
	v3274 = m.T0[v3240].(func(*base.Module, int32, int32, int32) int32)(m, v3244+int32(8), v3138, v3231)
	mBase = m.M
	v3275 = m.ExcPending
	if v3275 != 0 {
		goto L4
	} else {
		goto L529
	}
L529:
	;
	if v3274 != 0 {
		goto L527
	} else {
		goto L530
	}
L530:
	;
	v3347 = int32(0)
	goto L521
L531:
	;
	goto L526
L532:
	;
	v3308 = *(*int32)(unsafe.Add(mBase, uint32(v3167)))
	*(*int32)(unsafe.Add(mBase, uint32(v3176))) = v3308
	*(*int32)(unsafe.Add(mBase, uint32(v3279))) = v3167
	*(*int32)(unsafe.Add(mBase, uint32(v3167))) = int32(0)
	goto L534
L533:
	;
	goto L534
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3149))) = v3212
	v3314 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+16))
	v3315 = m.T0[v3314].(func(*base.Module, int32, int32, int32) int32)(m, v3112, v3138, v3231)
	mBase = m.M
	v3316 = m.ExcPending
	if v3316 != 0 {
		goto L4
	} else {
		goto L535
	}
L535:
	;
	v3347 = int32(1)
	goto L521
L536:
	;
	v3355 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3141))) = v3355
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_45), v3141)
	mBase = m.M
	v3359 = m.ExcPending
	if v3359 != 0 {
		goto L4
	} else {
		goto L537
	}
L537:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_46), int32(1170), int32(_a_F_PrepareTransaction_47))
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L4
	} else {
		goto L538
	}
L538:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L540:
	;
	v3400 = *(*int32)(unsafe.Add(mBase, uint32(v3136)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v3141)+16)) = v3400
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_48), v3141+int32(16))
	mBase = m.M
	v3406 = m.ExcPending
	if v3406 != 0 {
		goto L4
	} else {
		goto L541
	}
L541:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_46), int32(1191), int32(_a_F_PrepareTransaction_47))
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L4
	} else {
		goto L542
	}
L542:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L543:
	;
	v3414 = *(*int32)(unsafe.Add(mBase, uint32(v3080)+4))
	if v3414 == int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3080)+4)) = v3080
	*(*int32)(unsafe.Add(mBase, uint32(v3080))) = v3080
	goto L546
L545:
	;
	goto L546
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3081)+4)) = v3080
	v3420 = *(*int32)(unsafe.Add(mBase, uint32(v3080)))
	*(*int32)(unsafe.Add(mBase, uint32(v3081))) = v3420
	*(*int32)(unsafe.Add(mBase, uint32(v3420)+4)) = v3081
	*(*int32)(unsafe.Add(mBase, uint32(v3080))) = v3081
	goto L495
L547:
	;
	goto L494
L548:
	;
	v3486 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v3488 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[30]))
	v3496 = v3488
	v3497 = v3486
	goto L486
L549:
	;
	goto L485
L550:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L4
	} else {
		goto L551
	}
L551:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_25), int32(0))
	mBase = m.M
	v3541 = m.ExcPending
	if v3541 != 0 {
		goto L4
	} else {
		goto L552
	}
L552:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3610), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3546 = m.ExcPending
	if v3546 != 0 {
		goto L4
	} else {
		goto L553
	}
L553:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L554:
	;
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_50), int32(0))
	mBase = m.M
	v3554 = m.ExcPending
	if v3554 != 0 {
		goto L4
	} else {
		goto L555
	}
L555:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3669), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3559 = m.ExcPending
	if v3559 != 0 {
		goto L4
	} else {
		goto L556
	}
L556:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L557:
	;
	F_errmsg_internal(m, int32(_a_F_PrepareTransaction_51), int32(0))
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L4
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_26), int32(3707), int32(_a_F_PrepareTransaction_49))
	mBase = m.M
	v3572 = m.ExcPending
	if v3572 != 0 {
		goto L4
	} else {
		goto L559
	}
L559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L560:
	;
	v3583 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v3584 = *(*int32)(unsafe.Add(mBase, uint32(v3583)+4))
	v3585 = *(*int32)(unsafe.Add(mBase, uint32(v3574)+48))
	v3589 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3584+v3585<<(uint(int32(2))%32)))) = v3589
	*(*int32)(unsafe.Add(mBase, uint32(v3574)+56)) = v3589
	*(*int64)(unsafe.Add(mBase, uint32(v3574)+36)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3574)+73)) = uint8(v3589)
	v3598 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[48]))
	v3599 = *(*int64)(unsafe.Add(mBase, uint32(v3598)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v3598)+56)) = v3599 + int64(1)
	v3603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3574)+276)))
	if v3603 == v3589 {
		goto L562
	} else {
		goto L563
	}
L561:
	;
	v3627 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v3627+int32(512))
	mBase = m.M
	v3631 = m.ExcPending
	if v3631 != 0 {
		goto L4
	} else {
		goto L566
	}
L562:
	;
	v3606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3574)+277)))
	if v3606 != int32(1) {
		goto L561
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	v3610 = v3585 << (uint(int32(1)) % 32)
	v3611 = int32(_a_F_PrepareTransaction_52)
	v3612 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v3613 = *(*int32)(unsafe.Add(mBase, uint32(v3612)+8))
	v3615 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v3610+v3613))) = uint8(v3615)
	v3618 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[14]))
	v3619 = *(*int32)(unsafe.Add(mBase, uint32(v3618)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v3619+v3610)+1)) = uint8(v3615)
	*(*uint16)(unsafe.Add(mBase, uint32(v3574)+276)) = uint16(v3615)
	goto L561
L565:
	;
	goto L564
L566:
	;
	v3633 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[6]))
	if v3633 != 0 {
		goto L567
	} else {
		goto L568
	}
L567:
	;
	v3635 = v3633
	goto L570
L568:
	;
	goto L569
L569:
	;
	v3699 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49]))
	v3700 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v3699, v3700, v3700, v3700)
	mBase = m.M
	v3704 = m.ExcPending
	if v3704 != 0 {
		goto L4
	} else {
		goto L574
	}
L570:
	;
	v3663 = *(*int32)(unsafe.Add(mBase, uint32(v3635)))
	v3665 = *(*int32)(unsafe.Add(mBase, uint32(v3635)+8))
	v3666 = *(*int32)(unsafe.Add(mBase, uint32(v3635)+4))
	m.T0[v3666].(func(*base.Module, int32, int32))(m, int32(4), v3665)
	mBase = m.M
	v3668 = m.ExcPending
	if v3668 != 0 {
		goto L4
	} else {
		goto L572
	}
L571:
	;
	goto L569
L572:
	;
	if v3663 != 0 {
		v3635 = v3663
		goto L570
	} else {
		goto L573
	}
L573:
	;
	goto L571
L574:
	;
	F_AtEOXact_Aio(m)
	mBase = m.M
	v3706 = m.ExcPending
	if v3706 != 0 {
		goto L4
	} else {
		goto L575
	}
L575:
	;
	F_AtEOXact_RelationCache(m, int32(1))
	mBase = m.M
	v3709 = m.ExcPending
	if v3709 != 0 {
		goto L4
	} else {
		goto L576
	}
L576:
	;
	F_AtEOXact_TypeCache(m)
	mBase = m.M
	v3711 = m.ExcPending
	if v3711 != 0 {
		goto L4
	} else {
		goto L577
	}
L577:
	;
	v3713 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[33]))
	if v3713 != 0 {
		goto L578
	} else {
		goto L579
	}
L578:
	;
	v3714 = *(*int32)(unsafe.Add(mBase, uint32(v3713)+20))
	if v3714 != 0 {
		goto L581
	} else {
		goto L582
	}
L579:
	;
	goto L580
L580:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[33])) = int32(0)
	F_pgstat_clear_snapshot(m)
	mBase = m.M
	v3810 = m.ExcPending
	if v3810 != 0 {
		goto L4
	} else {
		goto L587
	}
L581:
	;
	v3716 = v3714
	goto L584
L582:
	;
	goto L583
L583:
	;
	goto L580
L584:
	;
	v3744 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v3744)+8)) = int32(0)
	v3747 = *(*int32)(unsafe.Add(mBase, uint32(v3716)+68))
	if v3747 != 0 {
		v3716 = v3747
		goto L584
	} else {
		goto L586
	}
L585:
	;
	goto L583
L586:
	;
	goto L585
L587:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[50])) = int32(0)
	v3815 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[51]))
	if v3815 != 0 {
		goto L588
	} else {
		goto L589
	}
L588:
	;
	v3816 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+20))
	v3817 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+28))
	if v3816 < v3817 {
		goto L591
	} else {
		goto L592
	}
L589:
	;
	goto L590
L590:
	;
	v3991 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[52]))
	if v3991 != 0 {
		goto L605
	} else {
		goto L606
	}
L591:
	;
	v3820 = v3816
	goto L594
L592:
	;
	goto L593
L593:
	;
	v3887 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+24))
	v3888 = *(*int32)(unsafe.Add(mBase, uint32(v3815)+32))
	if v3887 < v3888 {
		goto L598
	} else {
		goto L599
	}
L594:
	;
	v3849 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[53]))
	F_LocalExecuteInvalidationMessage(m, v3849+v3820<<(uint(int32(4))%32))
	mBase = m.M
	v3854 = m.ExcPending
	if v3854 != 0 {
		goto L4
	} else {
		goto L596
	}
L595:
	;
	goto L593
L596:
	;
	v3856 = v3820 + int32(1)
	if v3856 != v3817 {
		v3820 = v3856
		goto L594
	} else {
		goto L597
	}
L597:
	;
	goto L595
L598:
	;
	v3891 = v3887
	goto L601
L599:
	;
	goto L600
L600:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[51])) = int32(0)
	goto L590
L601:
	;
	v3920 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[54]))
	F_LocalExecuteInvalidationMessage(m, v3920+v3891<<(uint(int32(4))%32))
	mBase = m.M
	v3925 = m.ExcPending
	if v3925 != 0 {
		goto L4
	} else {
		goto L603
	}
L602:
	;
	goto L600
L603:
	;
	v3927 = v3891 + int32(1)
	if v3927 != v3888 {
		v3891 = v3927
		goto L601
	} else {
		goto L604
	}
L604:
	;
	goto L602
L605:
	;
	v3993 = v3991
	goto L608
L606:
	;
	goto L607
L607:
	;
	v4056 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34]))
	v4058 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v4062 = *(*int32)(unsafe.Add(mBase, uint32(v4056+v4058<<(uint(int32(2))%32))))
	if v4062 != 0 {
		goto L612
	} else {
		goto L613
	}
L608:
	;
	v4022 = *(*int32)(unsafe.Add(mBase, uint32(v3993)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[52])) = v4022
	F_pfree(m, v3993)
	mBase = m.M
	v4025 = m.ExcPending
	if v4025 != 0 {
		goto L4
	} else {
		goto L610
	}
L609:
	;
	goto L607
L610:
	;
	if v4022 != 0 {
		v3993 = v4022
		goto L608
	} else {
		goto L611
	}
L611:
	;
	goto L609
L612:
	;
	v4064 = F_TwoPhaseGetDummyProcNumber(m, v114, int32(0))
	mBase = m.M
	v4065 = m.ExcPending
	if v4065 != 0 {
		goto L4
	} else {
		goto L615
	}
L613:
	;
	v4095 = v4058
	goto L614
L614:
	;
	v4097 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[55]))
	v4101 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4097+v4095<<(uint(int32(2))%32)))) = v4101
	v4104 = int32(_a_F_PrepareTransaction_53)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[56])) = v4104
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[57])) = v4104
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[58])) = v4101
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[59])) = v4101
	v4116 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28]))
	if v4116 != 0 {
		goto L618
	} else {
		goto L619
	}
L615:
	;
	v4067 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v4071 = F_LWLockAcquire(m, v4067+int32(1664), int32(0))
	mBase = m.M
	v4072 = m.ExcPending
	if v4072 != 0 {
		goto L4
	} else {
		goto L616
	}
L616:
	;
	v4074 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[34]))
	v4075 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v4074+v4064<<(uint(v4075)%32)))) = v4062
	v4080 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	*(*int32)(unsafe.Add(mBase, uint32(v4074+v4080<<(uint(v4075)%32)))) = int32(0)
	v4087 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v4087+int32(1664))
	mBase = m.M
	v4091 = m.ExcPending
	if v4091 != 0 {
		goto L4
	} else {
		goto L617
	}
L617:
	;
	v4093 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[35]))
	v4095 = v4093
	goto L614
L618:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v4116)+112)) = int64(-4294967296)
	v4120 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[60]))
	F_hash_destroy(m, v4120)
	mBase = m.M
	v4122 = m.ExcPending
	if v4122 != 0 {
		goto L4
	} else {
		goto L621
	}
L619:
	;
	goto L620
L620:
	;
	v4133 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49]))
	v4135 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v4133, int32(2), v4135, v4135)
	mBase = m.M
	v4138 = m.ExcPending
	if v4138 != 0 {
		goto L4
	} else {
		goto L622
	}
L621:
	;
	v4124 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[28])) = v4124
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[60])) = v4124
	*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[61])) = uint8(v4124)
	goto L620
L622:
	;
	v4140 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49]))
	v4142 = int32(1)
	F_ResourceOwnerReleaseInternal(m, v4140, int32(3), v4142, v4142)
	mBase = m.M
	v4145 = m.ExcPending
	if v4145 != 0 {
		goto L4
	} else {
		goto L623
	}
L623:
	;
	v4147 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	v4151 = F_LWLockAcquire(m, v4147+int32(2304), int32(0))
	mBase = m.M
	v4152 = m.ExcPending
	if v4152 != 0 {
		goto L4
	} else {
		goto L624
	}
L624:
	;
	v4154 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[47]))
	*(*int32)(unsafe.Add(mBase, uint32(v4154)+40)) = int32(-1)
	v4158 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[17]))
	F_LWLockRelease(m, v4158+int32(2304))
	mBase = m.M
	v4162 = m.ExcPending
	if v4162 != 0 {
		goto L4
	} else {
		goto L625
	}
L625:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[47])) = int32(0)
	v4166 = int32(1)
	F_AtEOXact_GUC(m, v4166, v4166)
	mBase = m.M
	v4169 = m.ExcPending
	if v4169 != 0 {
		goto L4
	} else {
		goto L626
	}
L626:
	;
	F_AtEOXact_SPI(m, int32(1))
	mBase = m.M
	v4172 = m.ExcPending
	if v4172 != 0 {
		goto L4
	} else {
		goto L627
	}
L627:
	;
	v4174 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[62])) = v4174
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[63])) = v4174
	goto L628
L628:
	;
	F_AtEOXact_on_commit_actions(m, int32(1))
	mBase = m.M
	v4181 = m.ExcPending
	if v4181 != 0 {
		goto L4
	} else {
		goto L629
	}
L629:
	;
	F_AtEOXact_Namespace(m, int32(1), int32(0))
	mBase = m.M
	v4185 = m.ExcPending
	if v4185 != 0 {
		goto L4
	} else {
		goto L630
	}
L630:
	;
	F_smgrdestroyall(m)
	mBase = m.M
	v4187 = m.ExcPending
	if v4187 != 0 {
		goto L4
	} else {
		goto L631
	}
L631:
	;
	F_AtEOXact_Files(m, int32(1))
	mBase = m.M
	v4190 = m.ExcPending
	if v4190 != 0 {
		goto L4
	} else {
		goto L632
	}
L632:
	;
	v4192 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[64])) = v4192
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[65])) = v4192
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[66])) = v4192
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[67])) = v4192
	goto L633
L633:
	;
	F_AtEOXact_HashTables(m, int32(1))
	mBase = m.M
	v4205 = m.ExcPending
	if v4205 != 0 {
		goto L4
	} else {
		goto L634
	}
L634:
	;
	v4206 = int32(1)
	F_AtEOXact_Snapshot(m, v4206, v4206)
	mBase = m.M
	v4209 = m.ExcPending
	if v4209 != 0 {
		goto L4
	} else {
		goto L635
	}
L635:
	;
	F_AtEOXact_ApplyLauncher(m, int32(0))
	mBase = m.M
	v4212 = m.ExcPending
	if v4212 != 0 {
		goto L4
	} else {
		goto L636
	}
L636:
	;
	F_AtEOXact_LogicalRepWorkers(m, int32(0))
	mBase = m.M
	v4215 = m.ExcPending
	if v4215 != 0 {
		goto L4
	} else {
		goto L637
	}
L637:
	;
	v4219 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrepareTransaction[68])))
	if v4219 != int32(1) {
		goto L639
	} else {
		goto L640
	}
L638:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[69])) = int32(0)
	v4252 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49]))
	F_ResourceOwnerDelete(m, v4252)
	mBase = m.M
	v4254 = m.ExcPending
	if v4254 != 0 {
		goto L4
	} else {
		goto L642
	}
L639:
	;
	goto L638
L640:
	;
	v4223 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[70]))
	if v4223 == int32(0) {
		goto L639
	} else {
		goto L641
	}
L641:
	;
	v4226 = int32(_a_F_PrepareTransaction_41)
	v4228 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	v4229 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v4228 + v4229
	v4232 = *(*int32)(unsafe.Add(mBase, uint32(v4223)))
	*(*int32)(unsafe.Add(mBase, uint32(v4223))) = v4232 + v4229
	*(*int64)(unsafe.Add(mBase, uint32(v4223)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4223))) = v4232 + int32(2)
	v4243 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[44])) = v4243 - v4229
	goto L639
L642:
	;
	v4255 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+40)) = v4255
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[49])) = v4255
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[71])) = v4255
	v4265 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[0]))
	v4266 = *(*int32)(unsafe.Add(mBase, uint32(v4265)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[26])) = v4266
	v4269 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[72]))
	F_MemoryContextReset(m, v4269)
	mBase = m.M
	v4271 = m.ExcPending
	if v4271 != 0 {
		goto L4
	} else {
		goto L643
	}
L643:
	;
	v4273 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[73])) = v4273
	*(*int32)(unsafe.Add(mBase, uint32(v4265)+36)) = v4273
	*(*int32)(unsafe.Add(mBase, uint32(v35)+56)) = v4273
	v4279 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v4279
	*(*int64)(unsafe.Add(mBase, uint32(v35)+28)) = v4279
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v4273
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v4279
	*(*int64)(unsafe.Add(mBase, _c_F_PrepareTransaction[74])) = v4279
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[75])) = v4273
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v4273
	v4295 = int32(_a_F_PrepareTransaction_4)
	v4297 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_PrepareTransaction[9])) = v4297 - int32(1)
	m.G0 = v32 + int32(16)
	return
L644:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4310 = m.ExcPending
	if v4310 != 0 {
		goto L4
	} else {
		goto L645
	}
L645:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_54), int32(0))
	mBase = m.M
	v4314 = m.ExcPending
	if v4314 != 0 {
		goto L4
	} else {
		goto L646
	}
L646:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2616), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v4319 = m.ExcPending
	if v4319 != 0 {
		goto L4
	} else {
		goto L647
	}
L647:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L648:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v4326 = m.ExcPending
	if v4326 != 0 {
		goto L4
	} else {
		goto L649
	}
L649:
	;
	F_errmsg(m, int32(_a_F_PrepareTransaction_55), int32(0))
	mBase = m.M
	v4330 = m.ExcPending
	if v4330 != 0 {
		goto L4
	} else {
		goto L650
	}
L650:
	;
	F_errfinish(m, int32(_a_F_PrepareTransaction_3), int32(2626), int32(_a_F_PrepareTransaction_0))
	mBase = m.M
	v4335 = m.ExcPending
	if v4335 != 0 {
		goto L4
	} else {
		goto L651
	}
L651:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RecordTransactionAbort(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v78 int64
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int32
	_ = v119
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v405 int32
	_ = v405
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v424 int64
	_ = v424
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v2
	if v22 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L7
	} else {
		goto L110
	}
L2:
	;
	m.G0 = v18 + int32(16)
	return v468
L3:
	;
	if l0 != 0 {
		v468 = v2
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v30 = F_TransactionIdDidCommit(m, v22)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[1])) = int64(0)
	v468 = v2
	goto L2
L7:
	;
	return int32(0)
L8:
	;
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[2])))
	v39 = F_smgrGetPendingDeletes(m, int32(0), v18+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+52))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	v45 = v44
	goto L13
L12:
	;
	v45 = v2
	goto L13
L13:
	;
	v51 = F_pgstat_get_transactional_drops(m, int32(0), v18+int32(8))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v53 = int32(_a_F_RecordTransactionAbort_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[3])) = v55 + int32(1)
	if l0 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[4]))
	v106 = int32(0)
	v108 = F_XactLogAbortRecord(m, v101, v43, v45, v39, v102, v51, v103, v105, v106, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L22
	}
L16:
	;
	v62 = m.G0
	v63 = int32(16)
	v64 = v62 - v63
	m.G0 = v64
	F_gettimeofday(m, v64)
	mBase = m.M
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	m.G0 = v64 + v63
	goto L19
L17:
	;
	goto L18
L18:
	;
	v78 = *(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[5]))
	if v78 != int64(0) {
		v101 = v78
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v101 = v68 + v67*int64(1000000) - int64(946684800000000)
	goto L15
L20:
	;
	v85 = m.G0
	v86 = int32(16)
	v87 = v85 - v86
	m.G0 = v87
	F_gettimeofday(m, v87)
	mBase = m.M
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	v91 = int64(*(*int32)(unsafe.Add(mBase, uint32(v87)+8)))
	m.G0 = v87 + v86
	v99 = v91 + v90*int64(1000000) - int64(946684800000000)
	goto L21
L21:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[5])) = v99
	v101 = v99
	goto L15
L22:
	;
	if base.Ui32((v35-int32(1))&int32(_a_F_RecordTransactionAbort_1)) <= base.Ui32(int32(_a_F_RecordTransactionAbort_2)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v115 = *(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[6]))
	v117 = *(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[1]))
	F_replorigin_session_advance(m, v115, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if l0 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v123 = *(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[1]))
	F_XLogSetAsyncXactLSN(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_TransactionIdAbortTree(m, v22, v43, v45)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v128 = int32(_a_F_RecordTransactionAbort_0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[3]))
	v131 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[3])) = v130 - v131
	v137 = v43 - v131
	if v137 < int32(0) {
		v205 = v22
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if l0 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	goto L32
L34:
	;
	if v43&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v142 = int32(2)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v45+v137<<(uint(v142)%32))))
	if base.B2i32(base.Ui32(v142) < base.Ui32(v145))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v22)) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v160 = v22
	v162 = v137
	goto L37
L37:
	;
	if v137 == int32(0) {
		v205 = v160
		goto L33
	} else {
		goto L45
	}
L38:
	;
	v160 = v157
	v162 = v43 - int32(2)
	goto L37
L39:
	;
	v157 = v145
	goto L38
L40:
	;
	if base.Ui32(v22) < base.Ui32(v145) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if int32(0) <= v22-v145 {
		v157 = v22
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v157 = v22
	goto L38
L44:
	;
	goto L39
L45:
	;
	v165 = v160
	v166 = v162
	goto L46
L46:
	;
	v170 = int32(3)
	v174 = v45 + v166<<(uint(int32(2))%32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if base.B2i32(base.Ui32(v165) < base.Ui32(v170))|base.B2i32(base.Ui32(v175) < base.Ui32(v170)) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v205 = v200
	goto L33
L48:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v174-int32(4))))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v188))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v185)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	v185 = v175
	goto L48
L50:
	;
	if v165-v175 < int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v175) <= base.Ui32(v165) {
		v185 = v165
		goto L48
	} else {
		goto L54
	}
L53:
	;
	v185 = v165
	goto L48
L54:
	;
	goto L49
L55:
	;
	if int32(1) < v166 {
		v165 = v200
		v166 = v166 - int32(2)
		goto L46
	} else {
		goto L62
	}
L56:
	;
	v200 = v188
	goto L55
L57:
	;
	if base.Ui32(v185) < base.Ui32(v188) {
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if int32(0) <= v185-v188 {
		v200 = v185
		goto L55
	} else {
		goto L61
	}
L60:
	;
	v200 = v185
	goto L55
L61:
	;
	goto L56
L62:
	;
	goto L47
L63:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v455 != 0 {
		goto L104
	} else {
		goto L105
	}
L64:
	;
	v210 = m.G0
	v212 = v210 - int32(32)
	m.G0 = v212
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[7]))
	v219 = F_LWLockAcquire(m, v215+int32(512), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[1])) = int64(0)
	goto L63
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[8]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+48))
	v227 = int32(1)
	v229 = v223 + v226<<(uint(v227)%32)
	v231 = v43 - v227
	if int32(0) <= v231 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v234 = v231
	goto L71
L69:
	;
	v329 = v225
	goto L70
L70:
	;
	v341 = v329 + int32(280)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+276)))
	v343 = v342
	goto L87
L71:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v252 = v250 + int32(280)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v45+v234<<(uint(int32(2))%32))))
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+276)))
	v258 = v257
	goto L75
L72:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v329 = v324
	goto L70
L73:
	;
	if int32(0) < v234 {
		v234 = v234 - int32(1)
		goto L71
	} else {
		goto L84
	}
L74:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+277)))
	if v299 != 0 {
		goto L73
	} else {
		goto L79
	}
L75:
	;
	if v258 == int32(0) {
		goto L74
	} else {
		goto L77
	}
L76:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v252+v257<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v287
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v290 = int32(1)
	v291 = v289 - v290
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v291)
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+276)))
	v297 = v295 - v290
	*(*uint8)(unsafe.Add(mBase, uint32(v294)+276)) = uint8(v297)
	goto L73
L77:
	;
	v276 = v258 - int32(1)
	v279 = v252 + v276<<(uint(int32(2))%32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v280 != v256 {
		v258 = v276
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v302 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	if v302 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+16)) = v256
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_3), v212+int32(16))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_4), int32(4046), int32(_a_F_RecordTransactionAbort_5))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	goto L73
L84:
	;
	goto L72
L85:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[10]))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v403)+48))
	v405 = base.I32_wrap_i64(v404)
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v205))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v405)) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L86:
	;
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+277)))
	if v384 != 0 {
		goto L85
	} else {
		goto L91
	}
L87:
	;
	if v343 == int32(0) {
		goto L86
	} else {
		goto L89
	}
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v341+v342<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v372
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	v375 = int32(1)
	v376 = v374 - v375
	*(*uint8)(unsafe.Add(mBase, uint32(v229))) = uint8(v376)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[9]))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+276)))
	v382 = v380 - v375
	*(*uint8)(unsafe.Add(mBase, uint32(v379)+276)) = uint8(v382)
	goto L85
L89:
	;
	v361 = v343 - int32(1)
	v364 = v341 + v361<<(uint(int32(2))%32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v365 != v22 {
		v343 = v361
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	v387 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	if v387 == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v22
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_3), v212)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_4), int32(4062), int32(_a_F_RecordTransactionAbort_5))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	goto L85
L96:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[10]))
	if v417 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v417 = base.B2i32(base.Ui32(v405) < base.Ui32(v205))
	goto L96
L98:
	;
	goto L99
L99:
	;
	v417 = int32(base.Ui32(v405-v205) >> (uint(int32(31)) % 32))
	goto L96
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v419)+48)) = v404 + base.I64_extend_i32_s(v205-v405)
	goto L102
L101:
	;
	goto L102
L102:
	;
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v419)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v419)+56)) = v424 + int64(1)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_RecordTransactionAbort[7]))
	F_LWLockRelease(m, v429+int32(512))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	m.G0 = v212 + int32(32)
	goto L63
L104:
	;
	F_pfree(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v51 == int32(0) {
		v468 = v205
		goto L2
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	F_pfree(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v468 = v205
	goto L2
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v22
	F_errmsg_internal(m, int32(_a_F_RecordTransactionAbort_6), v18)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_RecordTransactionAbort_7), int32(1794), int32(_a_F_RecordTransactionAbort_8))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SaveTransactionCharacteristics(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SaveTransactionCharacteristics[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v3
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SaveTransactionCharacteristics[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v6)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SaveTransactionCharacteristics[2])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v9)
	return
}
func F_StartTransactionCommand(m *base.Module) {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(int32(19)) < base.Ui32(v10) {
		v46 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[2])) = v46
		m.G0 = v6 + int32(16)
		return
	} else {
		if v10 != 0 {
			if int32(1)<<(uint(v10)%32)&int32(_a_F_StartTransactionCommand_0) == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[2])) = v46
				m.G0 = v6 + int32(16)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
					if base.Ui32(v23) <= base.Ui32(int32(19)) {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_c_F_StartTransactionCommand[3])))
						v30 = v28
					} else {
						v30 = int32(_a_F_StartTransactionCommand_1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v30
					F_errmsg_internal(m, int32(_a_F_StartTransactionCommand_2), v6)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_StartTransactionCommand_3), int32(3114), int32(_a_F_StartTransactionCommand_4))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
			F_StartTransaction(m)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(1)
				v46 = *(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_StartTransactionCommand[2])) = v46
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_TransactionIdFollows(m *base.Module, l0 int32, l1 int32) int32 {
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l1))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		return base.B2i32(base.Ui32(l1) < base.Ui32(l0))
	} else {
		return base.B2i32(int32(0) < l0-l1)
	}
}
func F_TransactionIdIsInProgress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
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
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[0]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[1]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v15))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[2]))
	F_LWLockRelease(m, v625+int32(512))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L67
	} else {
		goto L174
	}
L2:
	;
	return int32(0)
L3:
	;
	if v27 != 0 {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v27 = base.B2i32(base.Ui32(l0) < base.Ui32(v15))
	goto L3
L5:
	;
	goto L6
L6:
	;
	v27 = int32(base.Ui32(l0-v15) >> (uint(int32(31)) % 32))
	goto L3
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3]))
	if v29 == l0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32(l0) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v150 != 0 {
		goto L49
	} else {
		goto L50
	}
L10:
	;
	v150 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[4]))
	if v41 == l0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v150 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[5]))
	if v45 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v150 = v142
	goto L9
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[6]))
	if v49 == int32(0) {
		v142 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[7]))
	v113 = int32(0)
	v115 = v45 - int32(1)
	goto L39
L20:
	;
	v54 = v49
	goto L21
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if v59 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v142 = int32(0)
	goto L16
L23:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v54)+80))
	if v106 != 0 {
		v54 = v106
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v62 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v65 = int32(1)
	if l0 == v62 {
		v142 = v65
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+52))
	v69 = v67 - int32(1)
	if v69 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v74 = int32(0)
	v76 = v69
	goto L28
L28:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v54)+48))
	v82 = int32(2)
	v83 = base.I32_div_s(v76-v74, v82)
	v84 = v83 + v74
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v80+v84<<(uint(v82)%32))))
	if v88 == l0 {
		v142 = v65
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v92 = F_TransactionIdPrecedes(m, v88, l0)
	mBase = m.M
	if v92 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v93 = v84 + int32(1)
	goto L33
L32:
	;
	v93 = v74
	goto L33
L33:
	;
	if v92 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v96 = v76
	goto L36
L35:
	;
	v96 = v84 - int32(1)
	goto L36
L36:
	;
	if v93 <= v96 {
		v74 = v93
		v76 = v96
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v120 = int32(2)
	v121 = base.I32_div_s(v115-v113, v120)
	v122 = v121 + v113
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v111+v122<<(uint(v120)%32))))
	v127 = base.B2i32(v126 == l0)
	if v126 == l0 {
		v142 = v127
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v142 = v127
	goto L16
L41:
	;
	v130 = base.B2i32(base.Ui32(v126) < base.Ui32(l0))
	if base.Ui32(v126) < base.Ui32(l0) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = v122 + int32(1)
	goto L44
L43:
	;
	v131 = v113
	goto L44
L44:
	;
	if base.Ui32(v126) < base.Ui32(l0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v134 = v115
	goto L47
L46:
	;
	v134 = v122 - int32(1)
	goto L47
L47:
	;
	if v131 <= v134 {
		v113 = v131
		v115 = v134
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
L49:
	;
	return int32(1)
L50:
	;
	goto L51
L51:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	if v154 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[2]))
	F_LWLockRelease(m, v538+int32(512))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L67
	} else {
		goto L158
	}
L53:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[0]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v427))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L67
	} else {
		goto L123
	}
L55:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])))
	if v160 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	goto L57
L57:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[10]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[11])) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[2]))
	v198 = F_LWLockAcquire(m, v194+int32(512), int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v182 = F_emscripten_builtin_malloc(m, v179<<(uint(int32(2))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8])) = v182
	if v182 == int32(0) {
		goto L54
	} else {
		goto L66
	}
L59:
	;
	if v170 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[12]))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+316))
	v168 = base.B2i32(v166 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])) = uint8(v168)
	v170 = v168
	goto L62
L61:
	;
	v170 = int32(0)
	goto L62
L62:
	;
	goto L59
L63:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[13]))
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[14]))
	v179 = (v172 + v174) * int32(65)
	goto L58
L64:
	;
	goto L65
L65:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v179 = v178
	goto L58
L66:
	;
	goto L57
L67:
	;
	return int32(0)
L68:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[15]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+48))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v204)) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v216 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v216 = base.B2i32(base.Ui32(v204) < base.Ui32(l0))
	goto L69
L71:
	;
	goto L72
L72:
	;
	v216 = int32(base.Ui32(v204-l0) >> (uint(int32(31)) % 32))
	goto L69
L73:
	;
	goto L1
L74:
	;
	goto L75
L75:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if int32(0) < v217 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[16]))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
	v227 = int32(0)
	v230 = v2
	goto L79
L77:
	;
	v321 = v2
	goto L78
L78:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])))
	if v330 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L79:
	;
	if v227 == v222 {
		v307 = v230
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v321 = v307
	goto L78
L81:
	;
	v315 = v227 + int32(1)
	if v315 != v217 {
		v227 = v315
		v230 = v307
		goto L79
	} else {
		goto L98
	}
L82:
	;
	v239 = v227 << (uint(int32(2)) % 32)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[11]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v239+v241)))
	if v243 == int32(0) {
		v307 = v230
		goto L81
	} else {
		goto L83
	}
L83:
	;
	if l0 == v243 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	goto L1
L85:
	;
	goto L86
L86:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v243))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v258 != 0 {
		v307 = v230
		goto L81
	} else {
		goto L91
	}
L88:
	;
	v258 = base.B2i32(base.Ui32(l0) < base.Ui32(v243))
	goto L87
L89:
	;
	goto L90
L90:
	;
	v258 = int32(base.Ui32(l0-v243) >> (uint(int32(31)) % 32))
	goto L87
L91:
	;
	v261 = v192 + v227<<(uint(int32(1))%32)
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[17]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v239+(v13+int32(36)))))
	v274 = v262
	goto L93
L92:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v292 != int32(1) {
		v307 = v230
		goto L81
	} else {
		goto L97
	}
L93:
	;
	if v274 <= int32(0) {
		goto L92
	} else {
		goto L95
	}
L94:
	;
	goto L1
L95:
	;
	v286 = v274 - int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v264+v266*int32(640)+int32(280)+v286<<(uint(int32(2))%32))))
	if v290 != l0 {
		v274 = v286
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v296+v230<<(uint(int32(2))%32)))) = v243
	v307 = v230 + int32(1)
	goto L81
L98:
	;
	goto L80
L99:
	;
	if v340 == int32(0) {
		v529 = v321
		goto L52
	} else {
		goto L103
	}
L100:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[12]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+316))
	v338 = base.B2i32(v336 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[9])) = uint8(v338)
	v340 = v338
	goto L102
L101:
	;
	v340 = int32(0)
	goto L102
L102:
	;
	goto L99
L103:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[0]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v344)+20))
	v347 = v345 - int32(1)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+16))
	if v347 < v348 {
		goto L53
	} else {
		goto L104
	}
L104:
	;
	v351 = v347
	v352 = v348
	goto L105
L105:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[18]))
	v363 = v351 + v352
	v364 = int32(2)
	v365 = base.I32_div_s(v363, v364)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v362+v365<<(uint(v364)%32))))
	if v369 != l0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v363 < int32(-1) {
		goto L53
	} else {
		goto L121
	}
L107:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v369))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	if v384 != 0 {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v384 = base.B2i32(base.Ui32(l0) < base.Ui32(v369))
	goto L110
L112:
	;
	goto L113
L113:
	;
	v384 = int32(base.Ui32(l0-v369) >> (uint(int32(31)) % 32))
	goto L110
L114:
	;
	v385 = v352
	goto L116
L115:
	;
	v385 = v365 + int32(1)
	goto L116
L116:
	;
	if v384 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v388 = v365 - int32(1)
	goto L119
L118:
	;
	v388 = v351
	goto L119
L119:
	;
	if v385 <= v388 {
		v351 = v388
		v352 = v385
		goto L105
	} else {
		goto L120
	}
L120:
	;
	goto L53
L121:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[19]))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393+v365))))
	if v395 != int32(1) {
		goto L53
	} else {
		goto L122
	}
L122:
	;
	goto L1
L123:
	;
	F_errcode(m, int32(_a_F_TransactionIdIsInProgress_0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L67
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_TransactionIdIsInProgress_1), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L67
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_TransactionIdIsInProgress_2), int32(1465), int32(_a_F_TransactionIdIsInProgress_3))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L67
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
	if v439 == int32(0) {
		v529 = v321
		goto L52
	} else {
		goto L131
	}
L128:
	;
	v439 = base.B2i32(base.Ui32(l0) <= base.Ui32(v427))
	goto L127
L129:
	;
	goto L130
L130:
	;
	v439 = base.B2i32(l0-v427 <= int32(0))
	goto L127
L131:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	v444 = int32(0)
	v447 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[0]))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	if v448 <= v449 {
		v515 = v444
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v529 = v515
	goto L52
L133:
	;
	v452 = v444
	v453 = v449
	v454 = v444
	goto L134
L134:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[19]))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v453))))
	if v465 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v515 = v508
	goto L132
L136:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[18]))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v469+v453<<(uint(int32(2))%32))))
	if v452 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L137:
	;
	v508 = v452
	v509 = v454
	goto L138
L138:
	;
	v512 = v453 + int32(1)
	if v512 != v448 {
		v452 = v508
		v453 = v512
		v454 = v509
		goto L134
	} else {
		goto L157
	}
L139:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v454))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v473)) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v489 = v454
	goto L141
L141:
	;
	if l0 != 0 {
		goto L149
	} else {
		goto L150
	}
L142:
	;
	if v487 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v487 = base.B2i32(base.Ui32(v473) < base.Ui32(v454))
	goto L142
L144:
	;
	goto L145
L145:
	;
	v487 = int32(base.Ui32(v473-v454) >> (uint(int32(31)) % 32))
	goto L142
L146:
	;
	v488 = v473
	goto L148
L147:
	;
	v488 = v454
	goto L148
L148:
	;
	v489 = v488
	goto L141
L149:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v473)) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443+v452<<(uint(int32(2))%32)))) = v473
	v508 = v452 + int32(1)
	v509 = v489
	goto L138
L152:
	;
	if v501 != 0 {
		v515 = v452
		goto L132
	} else {
		goto L156
	}
L153:
	;
	v501 = base.B2i32(base.Ui32(l0) <= base.Ui32(v473))
	goto L152
L154:
	;
	goto L155
L155:
	;
	v501 = base.B2i32(int32(0) <= v473-l0)
	goto L152
L156:
	;
	goto L151
L157:
	;
	goto L135
L158:
	;
	if v529 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v543 = F_TransactionIdDidAbort(m, l0)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L67
	} else {
		goto L163
	}
L160:
	;
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[3])) = l0
	goto L2
L162:
	;
	goto L161
L163:
	;
	if v543 != 0 {
		goto L162
	} else {
		goto L164
	}
L164:
	;
	v545 = F_SubTransGetTopmostTransaction(m, l0)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L67
	} else {
		goto L165
	}
L165:
	;
	if v545 == l0 {
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdIsInProgress[8]))
	v552 = int32(0)
	goto L167
L167:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v549+v552<<(uint(int32(2))%32))))
	v565 = base.B2i32(v545 == v564)
	if v565 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	if v565 == int32(0) {
		goto L162
	} else {
		goto L173
	}
L169:
	;
	v569 = v552 + int32(1)
	if v569 != v529 {
		v552 = v569
		goto L167
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	goto L168
L172:
	;
	goto L171
L173:
	;
	return int32(1)
L174:
	;
	return int32(1)
}
func F_TransactionIdSetTreeStatus(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v37 int32
	_ = v37
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v394 int64
	_ = v394
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
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
	var v472 int32
	_ = v472
	var v479 int32
	_ = v479
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v581 int32
	_ = v581
	var v593 int64
	_ = v593
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v627 int32
	_ = v627
	var v630 int64
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v670 int32
	_ = v670
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int64
	_ = v734
	var v739 int32
	_ = v739
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v824 int32
	_ = v824
	var v843 int64
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v877 int32
	_ = v877
	var v880 int64
	_ = v880
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v891 int64
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	v6 = int32(0)
	v25 = int32(base.Ui32(l0) >> (uint(int32(15)) % 32))
	v26 = base.I64_extend_i32_u(v25)
	if l1 <= v6 {
		v70 = v6
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v559 = l1 - v70
	v560 = int32(0)
	if base.B2i32(l3 != int32(1))|base.B2i32(v559 <= v560) == v560 {
		goto L96
	} else {
		goto L97
	}
L2:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v115 = base.I32_rem_u_s(base.I32_wrap_i64(v26), v114)
	v118 = v111 + v115<<(uint(int32(7))%32)
	if int32(5) < l1 {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	if l1 != v70 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v37 = v6
	goto L5
L5:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2+v37<<(uint(int32(2))%32))))
	if int32(base.Ui32(v55)>>(uint(int32(15))%32)) != v25 {
		v70 = v37
		goto L3
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v60 = v37 + int32(1)
	if v60 != l1 {
		v37 = v60
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	goto L2
L10:
	;
	return
L11:
	;
	F_TransactionIdSetPageStatusInternal(m, l0, l1, l2, l3, l4, v26)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L38
	} else {
		goto L94
	}
L12:
	;
	v505 = F_LWLockAcquire(m, v118, int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L38
	} else {
		goto L93
	}
L13:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+36))
	if l0 != v123 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+276)))
	if l1 != v125 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if l1 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v128 = v122 + int32(280)
	v130 = l1 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v130) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L18
L18:
	;
	v194 = F_LWLockConditionalAcquire(m, v118, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	if v192 != 0 {
		goto L12
	} else {
		goto L37
	}
L20:
	;
	v192 = int32(0)
	goto L19
L21:
	;
	v166 = v161
	v167 = v162
	v168 = v163
	goto L31
L22:
	;
	if (l2|v128)&int32(3) != 0 {
		v161 = l2
		v162 = v128
		v163 = v130
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v154 = l2
	v155 = v128
	v156 = v130
	goto L24
L24:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v138 = l2
	v139 = v128
	v140 = v130
	goto L26
L26:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v143 != v144 {
		v161 = v138
		v162 = v139
		v163 = v140
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v154 = v149
	v155 = v147
	v156 = v151
	goto L24
L28:
	;
	v146 = int32(4)
	v147 = v139 + v146
	v149 = v138 + v146
	v151 = v140 - v146
	if base.Ui32(int32(3)) < base.Ui32(v151) {
		v138 = v149
		v139 = v147
		v140 = v151
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v161 = v154
	v162 = v155
	v163 = v156
	goto L21
L31:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 == v172 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v192 = v171 - v172
	goto L19
L33:
	;
	v174 = int32(1)
	v179 = v168 - v174
	if v179 != 0 {
		v166 = v166 + v174
		v167 = v167 + v174
		v168 = v179
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L20
L37:
	;
	goto L18
L38:
	;
	return
L39:
	;
	if v194 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v199)+576)) = l4
	*(*int64)(unsafe.Add(mBase, uint32(v199)+568)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v199)+564)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v199)+560)) = l0
	v204 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)) = uint8(v204)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v210 = v206
	goto L42
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = int32(-1)
	v479 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)) = uint8(v479)
	goto L12
L42:
	;
	if v210 != int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+28))
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v199)+568))
	v332 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v333 = base.I64_rem_s(v330, v332)
	v337 = v329 + base.I32_wrap_i64(v333)<<(uint(int32(7))%32)
	v339 = F_LWLockAcquire(m, v337, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L38
	} else {
		goto L65
	}
L44:
	;
	goto L43
L45:
	;
	v210 = v325
	goto L42
L46:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v234+v210*int32(640))+568))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v199)+568))
	if v238 != v239 {
		goto L41
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v316 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = v316
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[4]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v322 = base.B2i32(v320 == v316)
	if v320 == v316 {
		goto L61
	} else {
		goto L62
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v199)+556)) = v210
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[4]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v245 = base.B2i32(v244 == v210)
	if v244 == v210 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v246 = v243
	goto L52
L51:
	;
	v246 = v244
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v246
	if v245 == int32(0) {
		v325 = v244
		goto L45
	} else {
		goto L53
	}
L53:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v252))) = int32(134217784)
	v256 = int32(0)
	goto L54
L54:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+552)))
	if v281 != 0 {
		v256 = v256 + int32(1)
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[5]))
	v284 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v284
	if v256 <= v284 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v288 = v256
	goto L58
L58:
	;
	v312 = int32(1)
	if base.Ui32(v312) < base.Ui32(v288) {
		v288 = v288 - v312
		goto L58
	} else {
		goto L60
	}
L59:
	;
	goto L10
L60:
	;
	goto L59
L61:
	;
	v323 = v319
	goto L63
L62:
	;
	v323 = v320
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v323
	if v320 == v316 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	v325 = v320
	goto L45
L65:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v197)+56))
	v342 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v197)+56)) = v342
	if v341 != v342 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v346 = v341
	v349 = v337
	v350 = v330
	goto L69
L67:
	;
	v413 = v337
	goto L68
L68:
	;
	if v413 != 0 {
		goto L82
	} else {
		goto L83
	}
L69:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)))
	v374 = v371 + v346*int32(640)
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v374)+568))
	if v350 == v375 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v413 = v396
	goto L68
L71:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v374)+560))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+276)))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v374)+564))
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v374)+576))
	F_TransactionIdSetPageStatusInternal(m, v399, v400, v374+int32(280), v403, v404, v398)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L38
	} else {
		goto L80
	}
L72:
	;
	v396 = v349
	v397 = v350
	v398 = v350
	goto L71
L73:
	;
	goto L74
L74:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+28))
	v381 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v382 = base.I64_rem_s(v375, v381)
	v386 = v379 + base.I32_wrap_i64(v382)<<(uint(int32(7))%32)
	if v386 == v349 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v394 = v375
	goto L77
L76:
	;
	F_LWLockRelease(m, v349)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L38
	} else {
		goto L78
	}
L77:
	;
	v396 = v386
	v397 = v375
	v398 = v394
	goto L71
L78:
	;
	v391 = F_LWLockAcquire(m, v386, int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L38
	} else {
		goto L79
	}
L79:
	;
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v374)+568))
	v394 = v393
	goto L77
L80:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v374)+556))
	if v407 != int32(-1) {
		v346 = v407
		v349 = v396
		v350 = v397
		goto L69
	} else {
		goto L81
	}
L81:
	;
	goto L70
L82:
	;
	F_LWLockRelease(m, v413)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L38
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v341 == int32(-1) {
		goto L10
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v438 = v341
	goto L87
L87:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[3]))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	v465 = v462 + v438*int32(640)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v465)+556))
	*(*int32)(unsafe.Add(mBase, uint32(v465)+556)) = int32(-1)
	v469 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465)+552)) = uint8(v469)
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[2]))
	if v472 != v465 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L10
L89:
	;
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v466 != int32(-1) {
		v438 = v466
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L11
L94:
	;
	F_LWLockRelease(m, v118)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L38
	} else {
		goto L95
	}
L95:
	;
	goto L10
L96:
	;
	v567 = l2 + v70<<(uint(int32(2))%32)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	v581 = v6
	v593 = base.I64_extend_i32_u(int32(base.Ui32(v568) >> (uint(int32(15)) % 32)))
	goto L99
L97:
	;
	goto L98
L98:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v797)+28))
	v800 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v801 = base.I32_rem_u_s(v25, v800)
	v804 = v798 + v801<<(uint(int32(7))%32)
	v806 = F_LWLockAcquire(m, v804, int32(0))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L38
	} else {
		goto L125
	}
L99:
	;
	v596 = v581 + int32(1)
	if v596 < v559 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	v598 = v559
	goto L103
L102:
	;
	v598 = v596
	goto L103
L103:
	;
	v599 = v598 - v581
	v602 = v581
	v607 = int32(0)
	goto L105
L104:
	;
	v643 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v643)+28))
	v647 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v648 = base.I32_rem_u_s(base.I32_wrap_i64(v593), v647)
	v651 = v644 + v648<<(uint(int32(7))%32)
	v653 = F_LWLockAcquire(m, v651, int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L38
	} else {
		goto L111
	}
L105:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v567+v602<<(uint(int32(2))%32))))
	v630 = base.I64_extend_i32_u(int32(base.Ui32(v627) >> (uint(int32(15)) % 32)))
	if v630 != v593 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v639 = v599
	v640 = v598
	v641 = v593
	goto L104
L107:
	;
	v639 = v607
	v640 = v602
	v641 = v630
	goto L104
L108:
	;
	goto L109
L109:
	;
	v632 = int32(1)
	v635 = v607 + v632
	if v635 != v599 {
		v602 = v602 + v632
		v607 = v635
		goto L105
	} else {
		goto L110
	}
L110:
	;
	goto L106
L111:
	;
	v657 = base.B2i32(l4 == int64(0))
	v659 = F_SimpleLruReadPage(m, int32(_a_F_TransactionIdSetTreeStatus_0), v593, v657, int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L38
	} else {
		goto L112
	}
L112:
	;
	if int32(0) < v639 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v670 = int32(0)
	goto L116
L114:
	;
	goto L115
L115:
	;
	v765 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)+12))
	v768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v766+v659))) = uint8(v768)
	F_LWLockRelease(m, v651)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L38
	} else {
		goto L123
	}
L116:
	;
	v693 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[6])))
	v694 = int32(1)
	v696 = int32(2)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v567+v581<<(uint(int32(2))%32)+v670<<(uint(v696)%32))))
	v703 = int32(base.Ui32(v699&int32(_a_F_TransactionIdSetTreeStatus_1)) >> (uint(v696) % 32))
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v705)+4))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v706+v659<<(uint(v696)%32))))
	v711 = v703 + v710
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v711))))
	v716 = v699 << (uint(v694) % 32) & int32(6)
	if base.B2i32(v693 == v694)&base.B2i32(int32(base.Ui32(v712)>>(uint(v716)%32))&int32(3) == v694) != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L115
L118:
	;
	v739 = v670 + int32(1)
	if v739 != v639 {
		v670 = v739
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v725 = v712 | int32(3)<<(uint(v716)%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v711))) = uint8(v725)
	if l4 == int64(0) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+36))
	v733 = v729 + v659<<(uint(int32(13))%32) + v703&int32(_a_F_TransactionIdSetTreeStatus_2)
	v734 = *(*int64)(unsafe.Add(mBase, uint32(v733)))
	if base.Ui64(l4) <= base.Ui64(v734) {
		goto L118
	} else {
		goto L121
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v733))) = l4
	goto L118
L122:
	;
	goto L117
L123:
	;
	if v640 < v559 {
		v581 = v640
		v593 = v641
		goto L99
	} else {
		goto L124
	}
L124:
	;
	goto L100
L125:
	;
	F_TransactionIdSetPageStatusInternal(m, l0, v70, l2, l3, l4, v26)
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L38
	} else {
		goto L126
	}
L126:
	;
	F_LWLockRelease(m, v804)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L38
	} else {
		goto L127
	}
L127:
	;
	if int32(0) < v559 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v816 = l2 + v70<<(uint(int32(2))%32)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	v824 = int32(0)
	v843 = base.I64_extend_i32_u(int32(base.Ui32(v817) >> (uint(int32(15)) % 32)))
	goto L131
L129:
	;
	goto L130
L130:
	;
	return
L131:
	;
	v846 = v824 + int32(1)
	if v846 < v559 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L130
L133:
	;
	v848 = v559
	goto L135
L134:
	;
	v848 = v846
	goto L135
L135:
	;
	v849 = v848 - v824
	v852 = v824
	v857 = int32(0)
	goto L137
L136:
	;
	v893 = *(*int32)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[0]))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)+28))
	v897 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_TransactionIdSetTreeStatus[1])))
	v898 = base.I32_rem_u_s(base.I32_wrap_i64(v843), v897)
	v901 = v894 + v898<<(uint(int32(7))%32)
	v903 = F_LWLockAcquire(m, v901, int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L38
	} else {
		goto L143
	}
L137:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v816+v852<<(uint(int32(2))%32))))
	v880 = base.I64_extend_i32_u(int32(base.Ui32(v877) >> (uint(int32(15)) % 32)))
	if v880 != v843 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v887 = v849
	v889 = v848
	v891 = v843
	goto L136
L139:
	;
	v887 = v857
	v889 = v852
	v891 = v880
	goto L136
L140:
	;
	goto L141
L141:
	;
	v882 = int32(1)
	v885 = v857 + v882
	if v885 != v849 {
		v852 = v852 + v882
		v857 = v885
		goto L137
	} else {
		goto L142
	}
L142:
	;
	goto L138
L143:
	;
	F_TransactionIdSetPageStatusInternal(m, int32(0), v887, v816+v824<<(uint(int32(2))%32), l3, l4, v843)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L38
	} else {
		goto L144
	}
L144:
	;
	F_LWLockRelease(m, v901)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L38
	} else {
		goto L145
	}
L145:
	;
	if v889 < v559 {
		v824 = v889
		v843 = v891
		goto L131
	} else {
		goto L146
	}
L146:
	;
	goto L132
}
func F_check_transaction_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_transaction_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
