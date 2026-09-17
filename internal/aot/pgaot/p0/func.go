package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ParseFuncOrColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v32 int32
	_ = v32
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v434 int32
	_ = v434
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
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
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v789 int32
	_ = v789
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
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
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v887 int32
	_ = v887
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1347 int32
	_ = v1347
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1368 int32
	_ = v1368
	var v1370 int32
	_ = v1370
	var v1375 int32
	_ = v1375
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1392 int32
	_ = v1392
	var v1416 int32
	_ = v1416
	var v1424 int32
	_ = v1424
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1541 int32
	_ = v1541
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1561 int32
	_ = v1561
	var v1585 int32
	_ = v1585
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1678 int32
	_ = v1678
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1762 int32
	_ = v1762
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1796 int32
	_ = v1796
	var v1798 int32
	_ = v1798
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1824 int32
	_ = v1824
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1836 int32
	_ = v1836
	var v1846 int32
	_ = v1846
	var v1876 int32
	_ = v1876
	var v1879 int32
	_ = v1879
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1891 int32
	_ = v1891
	var v1896 int32
	_ = v1896
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1923 int32
	_ = v1923
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1976 int32
	_ = v1976
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1989 int32
	_ = v1989
	var v1990 int32
	_ = v1990
	var v1996 int32
	_ = v1996
	var v1998 int32
	_ = v1998
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2034 int32
	_ = v2034
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2046 int32
	_ = v2046
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	v8 = int32(0)
	v32 = m.G0
	v34 = v32 - int32(1008)
	m.G0 = v34
	if l4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+27)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+26)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+25)))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+24)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l4)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if v43 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v54 = v8
	v55 = v8
	v56 = v8
	v57 = v8
	v58 = v8
	v59 = v8
	v60 = v8
	v61 = int32(0)
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v46 = F_transformWhereClause(m, l0, v43, int32(8), int32(_a_F_ParseFuncOrColumn_0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v51 = int32(0)
	goto L6
L6:
	;
	v54 = v42
	v55 = v40
	v56 = v39
	v57 = v41
	v58 = v38
	v59 = v37
	v60 = v36
	v61 = v51
	goto L3
L7:
	;
	return int32(0)
L8:
	;
	v51 = v46
	goto L6
L9:
	;
	v490 = int32(0)
	v492 = int32(1)
	if v467|(base.B2i32(l1 == v490)|(v59|(v58|(v56|(l5|base.B2i32(v469 != v492)|base.B2i32(v54 != v490)|base.B2i32(v61 != v490)))|base.B2i32(v57 != v490)))&v492) != 0 {
		v532 = v8
		goto L87
	} else {
		goto L88
	}
L10:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v192 < v193 {
		goto L35
	} else {
		goto L36
	}
L11:
	;
	v189 = int32(0)
	v460 = v189
	v467 = v189
	v469 = v168
	v480 = int32(1)
	v489 = v189
	goto L9
L12:
	;
	v168 = v8
	goto L11
L13:
	;
	goto L14
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v64 <= int32(100) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v152 != 0 {
		goto L10
	} else {
		goto L34
	}
L16:
	;
	v67 = int32(0)
	v73 = l2
	v79 = v67
	v82 = v8
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L29
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v102 <= v79 {
		v152 = v73
		v154 = v82
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v152 = v73
	v154 = v128
	goto L15
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v79<<(uint(int32(2))%32))))
	v109 = F_exprType(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(576)+v82<<(uint(int32(2))%32)))) = v109
	v125 = int32(1)
	v128 = v82 + v125
	if v73 != 0 {
		v79 = v79 + v125
		v82 = v128
		goto L19
	} else {
		goto L28
	}
L23:
	;
	if v109 != int32(2278) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if base.B2i32(v113 != int32(8))|(base.B2i32(l4 == v67)|v55) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v117 = F_list_delete_nth_cell(m, v73, v79)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	if v117 != 0 {
		v73 = v117
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v152 = v117
	v154 = v82
	goto L15
L28:
	;
	goto L20
L29:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v136 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+512)) = v136
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_1), int32(_a_F_ParseFuncOrColumn_2), v136, v34+int32(512))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(142), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v168 = v154
	goto L11
L35:
	;
	v208 = v8
	v213 = v8
	goto L38
L36:
	;
	v434 = v192
	goto L37
L37:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)))
	v460 = v152
	v467 = v434
	v469 = v154
	v480 = v8
	v489 = v457
	goto L9
L38:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v213<<(uint(int32(2))%32))))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	if v232 == int32(16) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v434 = v420
	goto L37
L40:
	;
	v422 = v213 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v422 < v423 {
		v208 = v420
		v213 = v422
		goto L38
	} else {
		goto L75
	}
L41:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v387 = F_lappend(m, v208, v386)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L74
	}
L42:
	;
	if v208 == int32(0) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v333 = int32(0)
	if v208 == v333 {
		v420 = v333
		goto L40
	} else {
		goto L68
	}
L45:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v237 <= int32(0) {
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v240 = int32(0)
	if v240 < v237 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v243 = v237
	goto L49
L48:
	;
	v243 = v240
	goto L49
L49:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v255 = int32(0)
	goto L50
L50:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v245+v255<<(uint(int32(2))%32))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if base.B2i32(v284 == int32(0))|base.B2i32(v284 != v287) != 0 {
		v305 = v284
		v306 = v287
		goto L53
	} else {
		goto L54
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L7
	} else {
		goto L63
	}
L52:
	;
	if v305-v306 != 0 {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v290 = v244
	v291 = v281
	goto L55
L55:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	if v295 == int32(0) {
		v305 = v295
		v306 = v294
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v305 = v295
	v306 = v294
	goto L53
L57:
	;
	v298 = int32(1)
	if v295 == v294 {
		v290 = v290 + v298
		v291 = v291 + v298
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v309 = v255 + int32(1)
	if v243 != v309 {
		v255 = v309
		goto L50
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L51
L62:
	;
	goto L41
L63:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+528)) = v318
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_5), v34+int32(528))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	F_parser_errposition(m, l0, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(196), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_6), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	v347 = F_exprLocation(m, v231)
	mBase = m.M
	F_parser_errposition(m, l0, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(206), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L7
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v420 = v387
	goto L40
L75:
	;
	goto L39
L76:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L7
	} else {
		goto L507
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2043 = m.ExcPending
	if v2043 != 0 {
		goto L7
	} else {
		goto L501
	}
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L7
	} else {
		goto L496
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L7
	} else {
		goto L491
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L7
	} else {
		goto L485
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L7
	} else {
		goto L480
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L7
	} else {
		goto L475
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L7
	} else {
		goto L469
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L7
	} else {
		goto L464
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1876 = m.ExcPending
	if v1876 != 0 {
		goto L7
	} else {
		goto L458
	}
L86:
	;
	m.G0 = v34 + int32(1008)
	return v1846
L87:
	;
	v534 = v34 + int32(536)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v534)+4)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v534)+16)) = v534
	v540 = int32(_a_F_ParseFuncOrColumn_7)
	v541 = *(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v534)+8)) = v541
	*(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0])) = v34 + int32(544)
	goto L100
L88:
	;
	v511 = int32(1)
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v512 != v511 {
		v532 = v8
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v34)+576))
	if v515 != int32(2249) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v518 = F_typeOrDomainTypeRelid(m, v515)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L7
	} else {
		goto L93
	}
L91:
	;
	v522 = v511
	goto L92
L92:
	;
	if l4 != 0 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v522 = base.B2i32(v518 != int32(0))
	goto L92
L94:
	;
	v532 = v522
	goto L87
L95:
	;
	goto L96
L96:
	;
	if v522 == int32(0) {
		v532 = v522
		goto L87
	} else {
		goto L97
	}
L97:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	v528 = F_ParseComplexProjection(m, l0, v527, v489, l6)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	if v528 != 0 {
		v1846 = v528
		goto L86
	} else {
		goto L99
	}
L99:
	;
	v532 = int32(1)
	goto L87
L100:
	;
	v549 = int32(1)
	v566 = F_func_get_detail(m, l1, v460, v467, v469, v34+int32(576), v59^v549, v549, l5, v34+int32(984), v34+int32(988), v34+int32(567), v34+int32(560), v34+int32(556), v34+int32(572), v34+int32(568))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v534)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ParseFuncOrColumn[0])) = v569
	goto L102
L102:
	;
	if l5 != 0 {
		goto L122
	} else {
		goto L123
	}
L103:
	;
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v34)+568))
	if v1424 == int32(0) {
		v1514 = v469
		goto L369
	} else {
		goto L370
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+160)) = v1008
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_8), v34+int32(160))
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L7
	} else {
		goto L365
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L7
	} else {
		goto L358
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+128)) = v965
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_9), v34+int32(128))
	mBase = m.M
	v1336 = m.ExcPending
	if v1336 != 0 {
		goto L7
	} else {
		goto L354
	}
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L7
	} else {
		goto L348
	}
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L7
	} else {
		goto L345
	}
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L7
	} else {
		goto L342
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L7
	} else {
		goto L334
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L7
	} else {
		goto L328
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1208 = m.ExcPending
	if v1208 != 0 {
		goto L7
	} else {
		goto L322
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L7
	} else {
		goto L319
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L7
	} else {
		goto L313
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L7
	} else {
		goto L307
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L7
	} else {
		goto L301
	}
L117:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L7
	} else {
		goto L295
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L7
	} else {
		goto L289
	}
L119:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L7
	} else {
		goto L282
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L7
	} else {
		goto L275
	}
L121:
	;
	if base.B2i32(v566 != int32(6))&base.B2i32(v610 != int32(2)) == int32(0) {
		goto L137
	} else {
		goto L138
	}
L122:
	;
	if v566&int32(3) != int32(2) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v566 == int32(3) {
		goto L120
	} else {
		goto L136
	}
L125:
	;
	v576 = v566 & int32(6)
	if v576 != int32(4) {
		v610 = v576
		goto L121
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L7
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	v589 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v589
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_10), v34)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_11), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(292), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L7
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
	v610 = v566 & int32(6)
	goto L121
L137:
	;
	if v56 != 0 {
		goto L119
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if v610 == int32(2) {
		v1416 = v8
		goto L103
	} else {
		goto L146
	}
L140:
	;
	if v58 != 0 {
		goto L118
	} else {
		goto L141
	}
L141:
	;
	if v55 != 0 {
		goto L117
	} else {
		goto L142
	}
L142:
	;
	if v54 != 0 {
		goto L116
	} else {
		goto L143
	}
L143:
	;
	if v61 != 0 {
		goto L115
	} else {
		goto L144
	}
L144:
	;
	if v57 != 0 {
		goto L114
	} else {
		goto L145
	}
L145:
	;
	goto L139
L146:
	;
	switch v566 - int32(4) {
	case 0:
		goto L149
	case 1:
		goto L148
	default:
		goto L147
	}
L147:
	;
	if v566 == int32(6) {
		goto L237
	} else {
		goto L238
	}
L148:
	;
	if v57 == int32(0) {
		goto L107
	} else {
		goto L229
	}
L149:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v34)+984))
	v624 = F_SearchSysCache1(m, int32(0), v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L7
	} else {
		goto L150
	}
L150:
	;
	if v624 == int32(0) {
		goto L113
	} else {
		goto L151
	}
L151:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v624)+16))
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v628)+22)))
	v630 = v628 + v629
	v631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v630)+6)))
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630)+4)))
	F_ReleaseCatCache(m, v624)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L7
	} else {
		goto L152
	}
L152:
	;
	if v632 != int32(110) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if v55 == int32(0) {
		goto L112
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	if v55 == int32(0) {
		v1416 = int32(110)
		goto L103
	} else {
		goto L222
	}
L156:
	;
	if v57 != 0 {
		goto L111
	} else {
		goto L157
	}
L157:
	;
	if v54 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v641 = v639
	goto L160
L159:
	;
	v641 = int32(0)
	goto L160
L160:
	;
	v642 = v469 - v641
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v34)+556))
	if v643 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if v480 != 0 {
		goto L204
	} else {
		goto L205
	}
L162:
	;
	if v632 != int32(104) {
		v1416 = v632
		goto L103
	} else {
		goto L203
	}
L163:
	;
	if v631 == v642 {
		goto L162
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v682 = int32(1)
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v34)+560))
	if v683 <= v682 {
		goto L175
	} else {
		goto L176
	}
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L7
	} else {
		goto L167
	}
L167:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L7
	} else {
		goto L168
	}
L168:
	;
	v656 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L7
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+272)) = v656
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(272))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L7
	} else {
		goto L170
	}
L170:
	;
	v664 = F_NameListToString(m, l1)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L7
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+264)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v34)+260)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v34)+256)) = v664
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_13), int32(_a_F_ParseFuncOrColumn_14), v631, v34+int32(256))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L7
	} else {
		goto L172
	}
L172:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L7
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(426), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L7
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	v686 = v682
	goto L177
L176:
	;
	v686 = v683
	goto L177
L177:
	;
	if v631 < v469-v686+int32(1) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	if v631 == v642 {
		goto L162
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	if v632 == int32(104) {
		goto L190
	} else {
		goto L191
	}
L181:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L7
	} else {
		goto L182
	}
L182:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L7
	} else {
		goto L183
	}
L183:
	;
	v701 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L7
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+304)) = v701
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(304))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L7
	} else {
		goto L185
	}
L185:
	;
	v709 = F_NameListToString(m, l1)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L7
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+296)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v34)+292)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v34)+288)) = v709
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_13), int32(_a_F_ParseFuncOrColumn_14), v631, v34+int32(288))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L7
	} else {
		goto L187
	}
L187:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L7
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(457), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L7
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	if v683 == v641<<(uint(int32(1))%32) {
		goto L161
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	if v683 <= v641 {
		goto L110
	} else {
		goto L202
	}
L193:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L7
	} else {
		goto L195
	}
L195:
	;
	v741 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L7
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+336)) = v741
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(336))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	v749 = F_NameListToString(m, l1)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L7
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+320)) = v749
	*(*int32)(unsafe.Add(mBase, uint32(v34)+328)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v34)+324)) = v683 - v641
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_15), v34+int32(320))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L7
	} else {
		goto L200
	}
L200:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(482), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L7
	} else {
		goto L201
	}
L201:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L202:
	;
	goto L162
L203:
	;
	goto L161
L204:
	;
	v774 = int32(0)
	goto L206
L205:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v774 = v773
	goto L206
L206:
	;
	v775 = v774 - v641
	v776 = v775 - v641
	if v776 < int32(0) {
		goto L109
	} else {
		goto L207
	}
L207:
	;
	if v775 <= v776 {
		v1416 = v632
		goto L103
	} else {
		goto L208
	}
L208:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v34)+572))
	v789 = v776
	goto L209
L209:
	;
	v812 = int32(2)
	v813 = v789 << (uint(v812) % 32)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v780+v813)))
	v819 = (v789 - v776 + v775) << (uint(v812) % 32)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v780+v819)))
	if v815 != v821 {
		goto L108
	} else {
		goto L211
	}
L210:
	;
	v1416 = v632
	goto L103
L211:
	;
	if v815 == int32(2276) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v826 = v825 + v819
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+1004)) = v827
	v829 = v825 + v813
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+1000)) = v830
	*(*int32)(unsafe.Add(mBase, uint32(v34)+252)) = v827
	*(*int32)(unsafe.Add(mBase, uint32(v34)+248)) = v830
	v838 = F_list_make2_impl(m, v34+int32(252), v34+int32(248))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L7
	} else {
		goto L215
	}
L213:
	;
	goto L214
L214:
	;
	v887 = v789 + int32(1)
	if v887 != v775 {
		v789 = v887
		goto L209
	} else {
		goto L221
	}
L215:
	;
	v842 = F_select_common_type(m, l0, v838, int32(_a_F_ParseFuncOrColumn_16), int32(0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L7
	} else {
		goto L216
	}
L216:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+996)) = v844
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+992)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v34)+244)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v34)+240)) = v846
	v854 = F_list_make2_impl(m, v34+int32(244), v34+int32(240))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L7
	} else {
		goto L217
	}
L217:
	;
	v856 = F_select_common_typmod(m, v854, v842)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L7
	} else {
		goto L218
	}
L218:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v829)))
	v860 = v34 + int32(576)
	v861 = v813 + v860
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	v866 = F_coerce_type(m, l0, v858, v862, v842, v856, int32(0), int32(2), int32(-1))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L7
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = v866
	*(*int32)(unsafe.Add(mBase, uint32(v861))) = v842
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v826)))
	v871 = v860 + v819
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	v876 = F_coerce_type(m, l0, v870, v872, v842, v856, int32(0), int32(2), int32(-1))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L7
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v826))) = v876
	*(*int32)(unsafe.Add(mBase, uint32(v871))) = v842
	goto L214
L221:
	;
	goto L210
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L7
	} else {
		goto L224
	}
L224:
	;
	v899 = F_NameListToString(m, l1)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L7
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+224)) = v899
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_17), v34+int32(224))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L7
	} else {
		goto L226
	}
L226:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L7
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(516), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L7
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	if v55 == int32(0) {
		v1416 = v8
		goto L103
	} else {
		goto L230
	}
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L7
	} else {
		goto L231
	}
L231:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L7
	} else {
		goto L232
	}
L232:
	;
	v925 = F_NameListToString(m, l1)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L7
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+432)) = v925
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_18), v34+int32(432))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L7
	} else {
		goto L234
	}
L234:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L7
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(536), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L7
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
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v942)))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v34)+576))
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v34)+988))
	v949 = F_coerce_type(m, l0, v943, v944, v945, int32(-1), int32(3), int32(0), l6)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L7
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	if v566 == int32(1) {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1846 = v949
	goto L86
L241:
	;
	if l4 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	if l4 == int32(0) {
		goto L255
	} else {
		goto L256
	}
L244:
	;
	v1846 = int32(0)
	goto L86
L245:
	;
	goto L246
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L7
	} else {
		goto L247
	}
L247:
	;
	F_errcode(m, int32(84439172))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L7
	} else {
		goto L248
	}
L248:
	;
	v965 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L7
	} else {
		goto L249
	}
L249:
	;
	if l5 != 0 {
		goto L106
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+144)) = v965
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_19), v34+int32(144))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L7
	} else {
		goto L251
	}
L251:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_20), int32(0))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L7
	} else {
		goto L252
	}
L252:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L7
	} else {
		goto L253
	}
L253:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(577), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L7
	} else {
		goto L254
	}
L254:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L255:
	;
	v1846 = int32(0)
	goto L86
L256:
	;
	goto L257
L257:
	;
	if v532 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v987)))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v988)+4))
	v990 = F_ParseComplexProjection(m, l0, v989, v489, l6)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L7
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	if v54 != 0 {
		goto L263
	} else {
		goto L264
	}
L261:
	;
	if v990 != 0 {
		v1846 = v990
		goto L86
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v55|base.B2i32(v993 < int32(2)) == int32(0) {
		goto L105
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L7
	} else {
		goto L267
	}
L266:
	;
	goto L265
L267:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L7
	} else {
		goto L268
	}
L268:
	;
	v1008 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L7
	} else {
		goto L269
	}
L269:
	;
	if l5 != 0 {
		goto L104
	} else {
		goto L270
	}
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+176)) = v1008
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(176))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L7
	} else {
		goto L271
	}
L271:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_21), int32(0))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L7
	} else {
		goto L272
	}
L272:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L7
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(636), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L7
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L7
	} else {
		goto L276
	}
L276:
	;
	v1036 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L7
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+496)) = v1036
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_22), v34+int32(496))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L7
	} else {
		goto L278
	}
L278:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_23), int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L7
	} else {
		goto L279
	}
L279:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1049 = m.ExcPending
	if v1049 != 0 {
		goto L7
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(302), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L7
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L7
	} else {
		goto L283
	}
L283:
	;
	v1062 = F_NameListToString(m, l1)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L7
	} else {
		goto L284
	}
L284:
	;
	v1064 = F_NameListToString(m, l1)
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L7
	} else {
		goto L285
	}
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v1064
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1062
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_24), v34+int32(16))
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L7
	} else {
		goto L286
	}
L286:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L7
	} else {
		goto L287
	}
L287:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(318), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L7
	} else {
		goto L288
	}
L288:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L289:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L7
	} else {
		goto L290
	}
L290:
	;
	v1087 = F_NameListToString(m, l1)
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L7
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1087
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_25), v34+int32(32))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L7
	} else {
		goto L292
	}
L292:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L7
	} else {
		goto L293
	}
L293:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(324), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L7
	} else {
		goto L294
	}
L294:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L295:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L7
	} else {
		goto L296
	}
L296:
	;
	v1109 = F_NameListToString(m, l1)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L7
	} else {
		goto L297
	}
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+48)) = v1109
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_26), v34+int32(48))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L7
	} else {
		goto L298
	}
L298:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L7
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(330), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L7
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L7
	} else {
		goto L302
	}
L302:
	;
	v1131 = F_NameListToString(m, l1)
	mBase = m.M
	v1132 = m.ExcPending
	if v1132 != 0 {
		goto L7
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+480)) = v1131
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_27), v34+int32(480))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L7
	} else {
		goto L304
	}
L304:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L7
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(336), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L7
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L7
	} else {
		goto L308
	}
L308:
	;
	v1153 = F_NameListToString(m, l1)
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L7
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+464)) = v1153
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_28), v34+int32(464))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L7
	} else {
		goto L310
	}
L310:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L7
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(342), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L7
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L7
	} else {
		goto L314
	}
L314:
	;
	v1175 = F_NameListToString(m, l1)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L7
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+448)) = v1175
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_29), v34+int32(448))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L7
	} else {
		goto L316
	}
L316:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L7
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(348), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L7
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+208)) = v623
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_30), v34+int32(208))
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L7
	} else {
		goto L320
	}
L320:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(369), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L7
	} else {
		goto L321
	}
L321:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L322:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L7
	} else {
		goto L323
	}
L323:
	;
	v1212 = F_NameListToString(m, l1)
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L7
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+400)) = v1212
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_31), v34+int32(400))
	mBase = m.M
	v1219 = m.ExcPending
	if v1219 != 0 {
		goto L7
	} else {
		goto L325
	}
L325:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L7
	} else {
		goto L326
	}
L326:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(386), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L7
	} else {
		goto L327
	}
L327:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L328:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L7
	} else {
		goto L329
	}
L329:
	;
	v1234 = F_NameListToString(m, l1)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L7
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+384)) = v1234
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_32), v34+int32(384))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L7
	} else {
		goto L331
	}
L331:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L7
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(392), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L7
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
	F_errcode(m, int32(52461700))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L7
	} else {
		goto L335
	}
L335:
	;
	v1258 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L7
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+368)) = v1258
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(368))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L7
	} else {
		goto L337
	}
L337:
	;
	v1266 = F_NameListToString(m, l1)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L7
	} else {
		goto L338
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+356)) = v631
	*(*int32)(unsafe.Add(mBase, uint32(v34)+352)) = v1266
	F_errhint_plural(m, int32(_a_F_ParseFuncOrColumn_33), int32(_a_F_ParseFuncOrColumn_34), v631, v34+int32(352))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L7
	} else {
		goto L339
	}
L339:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L7
	} else {
		goto L340
	}
L340:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(498), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L7
	} else {
		goto L341
	}
L341:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L342:
	;
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_35), int32(0))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L7
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(1755), int32(_a_F_ParseFuncOrColumn_36))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L7
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	F_errmsg_internal(m, int32(_a_F_ParseFuncOrColumn_37), int32(0))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L7
	} else {
		goto L346
	}
L346:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(1768), int32(_a_F_ParseFuncOrColumn_36))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L7
	} else {
		goto L347
	}
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1315 = m.ExcPending
	if v1315 != 0 {
		goto L7
	} else {
		goto L349
	}
L349:
	;
	v1316 = F_NameListToString(m, l1)
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L7
	} else {
		goto L350
	}
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+416)) = v1316
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_38), v34+int32(416))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L7
	} else {
		goto L351
	}
L351:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L7
	} else {
		goto L352
	}
L352:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(529), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L7
	} else {
		goto L353
	}
L353:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L354:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_39), int32(0))
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L7
	} else {
		goto L355
	}
L355:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L7
	} else {
		goto L356
	}
L356:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(568), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L7
	} else {
		goto L357
	}
L357:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L358:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L7
	} else {
		goto L359
	}
L359:
	;
	v1357 = F_func_signature_string(m, l1, v469, v467, v34+int32(576))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L7
	} else {
		goto L360
	}
L360:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+192)) = v1357
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_12), v34+int32(192))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L7
	} else {
		goto L361
	}
L361:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_40), int32(0))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L7
	} else {
		goto L362
	}
L362:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L7
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(617), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L7
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_41), int32(0))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L7
	} else {
		goto L366
	}
L366:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1387 = m.ExcPending
	if v1387 != 0 {
		goto L7
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(627), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L7
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
	v1541 = int32(0)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v34)+572))
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v34)+988))
	v1547 = F_enforce_generic_type_consistency(m, v34+int32(576), v1544, v1514, v1545, v1541)
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L7
	} else {
		goto L387
	}
L370:
	;
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	if v1427 <= int32(0) {
		v1514 = v469
		goto L369
	} else {
		goto L371
	}
L371:
	;
	v1430 = int32(100)
	if v469 <= v1430 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1433 = v1430
	goto L374
L373:
	;
	v1433 = v469
	goto L374
L374:
	;
	v1440 = v469
	v1444 = int32(0)
	goto L375
L375:
	;
	if v1444 != v1433-v469 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1490 = m.ExcPending
	if v1490 != 0 {
		goto L7
	} else {
		goto L382
	}
L377:
	;
	v1470 = int32(2)
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+12))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1444<<(uint(v1470)%32))))
	v1478 = F_exprType(m, v1477)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L7
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	goto L376
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+int32(576)+v1440<<(uint(v1470)%32)))) = v1478
	v1481 = int32(1)
	v1482 = v1440 + v1481
	v1484 = v1444 + v1481
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	if v1484 < v1485 {
		v1440 = v1482
		v1444 = v1484
		goto L375
	} else {
		goto L381
	}
L381:
	;
	v1514 = v1482
	goto L369
L382:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v1493 = m.ExcPending
	if v1493 != 0 {
		goto L7
	} else {
		goto L383
	}
L383:
	;
	v1494 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+112)) = v1494
	F_errmsg_plural(m, int32(_a_F_ParseFuncOrColumn_1), int32(_a_F_ParseFuncOrColumn_2), v1494, v34+int32(112))
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L7
	} else {
		goto L384
	}
L384:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1504 = m.ExcPending
	if v1504 != 0 {
		goto L7
	} else {
		goto L385
	}
L385:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(659), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1509 = m.ExcPending
	if v1509 != 0 {
		goto L7
	} else {
		goto L386
	}
L386:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+988)) = v1547
	if v480 != 0 {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v34)+556))
	v1652 = int32(0)
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v34)+560))
	if base.B2i32(v1651 == int32(2276))|base.B2i32(v1657 <= v1652) == v1652 {
		goto L401
	} else {
		goto L402
	}
L389:
	;
	v1550 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v1550 <= int32(0) {
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v1561 = v1541
	goto L391
L391:
	;
	v1585 = v1561 << (uint(int32(2)) % 32)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1585+(v34+int32(576)))))
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1585+v1544)))
	if v1589 == v1591 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	goto L388
L393:
	;
	v1617 = v1561 + int32(1)
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v1617 < v1618 {
		v1561 = v1617
		goto L391
	} else {
		goto L400
	}
L394:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v1594 = v1593 + v1585
	v1595 = *(*int32)(unsafe.Add(mBase, uint32(v1594)))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1595)))
	if v1596 == int32(16) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v1595)+4))
	v1600 = int32(-1)
	v1604 = F_coerce_type(m, l0, v1599, v1589, v1591, v1600, int32(0), int32(2), v1600)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L7
	} else {
		goto L398
	}
L396:
	;
	goto L397
L397:
	;
	v1607 = int32(-1)
	v1611 = F_coerce_type(m, l0, v1595, v1589, v1591, v1607, int32(0), int32(2), v1607)
	mBase = m.M
	v1612 = m.ExcPending
	if v1612 != 0 {
		goto L7
	} else {
		goto L399
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1595)+4)) = v1604
	goto L393
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1594))) = v1611
	goto L393
L400:
	;
	goto L392
L401:
	;
	v1664 = F_palloc0(m, int32(36))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L7
	} else {
		goto L404
	}
L402:
	;
	v1701 = v460
	v1702 = v59 & base.B2i32(v1651 != v1652)
	v1704 = v1651
	goto L403
L403:
	;
	v1705 = int32(0)
	if base.B2i32(v1702 == v1705)|base.B2i32(v469 <= v1705)|base.B2i32(v1704 != int32(2276)) == v1705 {
		goto L417
	} else {
		goto L418
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664))) = int32(35)
	v1668 = v469 - v1657
	v1669 = F_list_copy_tail(m, v460, v1668)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L7
	} else {
		goto L405
	}
L405:
	;
	v1671 = int32(0)
	if base.B2i32(v460 == v1671)|base.B2i32(v1668 <= v1671) != 0 {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664)+16)) = v1669
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1669)+12))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1683)))
	v1685 = F_exprType(m, v1684)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L7
	} else {
		goto L413
	}
L407:
	;
	v1681 = int32(0)
	goto L409
L408:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	if v1668 < v1678 {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L406
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v460)+4)) = v1668
	goto L412
L411:
	;
	goto L412
L412:
	;
	v1681 = v460
	goto L409
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664)+12)) = v1685
	v1688 = F_get_array_type(m, v1685)
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L7
	} else {
		goto L414
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1664)+4)) = v1688
	if v1688 == int32(0) {
		goto L85
	} else {
		goto L415
	}
L415:
	;
	v1693 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1664)+20)) = uint8(v1693)
	v1695 = F_exprLocation(m, v1669)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v1664)+32)) = v1695
	v1698 = F_lappend(m, v1681, v1664)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L7
	} else {
		goto L416
	}
L416:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v34)+556))
	v1701 = v1698
	v1702 = int32(1)
	v1704 = v1700
	goto L403
L417:
	;
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v469<<(uint(int32(2))%32)+v34)+572))
	v1719 = F_get_base_element_type(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L7
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	v1723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+567)))
	if v1723 == int32(1) {
		goto L422
	} else {
		goto L423
	}
L420:
	;
	if v1719 == int32(0) {
		goto L84
	} else {
		goto L421
	}
L421:
	;
	goto L419
L422:
	;
	F_check_srf_call_placement(m, l0, l3, l6)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L7
	} else {
		goto L425
	}
L423:
	;
	goto L424
L424:
	;
	if v610 == int32(2) {
		goto L427
	} else {
		goto L428
	}
L425:
	;
	goto L424
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v1836
	v1846 = v1836
	goto L86
L427:
	;
	v1731 = F_palloc0(m, int32(36))
	mBase = m.M
	v1732 = m.ExcPending
	if v1732 != 0 {
		goto L7
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	if base.B2i32(v566 != int32(4))|v57 == int32(0) {
		goto L433
	} else {
		goto L434
	}
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1731))) = int32(15)
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v34)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+4)) = v1735
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v34)+988))
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+32)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+28)) = v1701
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+16)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v1731)+13)) = uint8(v1702)
	*(*uint8)(unsafe.Add(mBase, uint32(v1731)+12)) = uint8(v1723)
	*(*int32)(unsafe.Add(mBase, uint32(v1731)+8)) = v1737
	if v1723 != 0 {
		v1836 = v1731
		goto L426
	} else {
		goto L431
	}
L431:
	;
	v1846 = v1731
	goto L86
L432:
	;
	F_transformAggregateCall(m, l0, v1750, v1701, v54, v58)
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L7
	} else {
		goto L456
	}
L433:
	;
	v1750 = F_palloc0(m, int32(72))
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L7
	} else {
		goto L436
	}
L434:
	;
	goto L435
L435:
	;
	v1798 = F_palloc0(m, int32(44))
	mBase = m.M
	v1799 = m.ExcPending
	if v1799 != 0 {
		goto L7
	} else {
		goto L445
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1750))) = int32(9)
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v34)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+4)) = v1754
	v1756 = *(*int32)(unsafe.Add(mBase, uint32(v34)+988))
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+68)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+64)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v1750)+56)) = int64(-4294967296)
	v1762 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+51)) = uint8(v1762)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+50)) = uint8(v1416)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+49)) = uint8(v1702)
	*(*uint8)(unsafe.Add(mBase, uint32(v1750)+48)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+44)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+20)) = v1762
	*(*int32)(unsafe.Add(mBase, uint32(v1750)+8)) = v1756
	if v55|(v56|base.B2i32(v1701 != v1762)) == v1762 {
		goto L83
	} else {
		goto L437
	}
L437:
	;
	if v1723 != 0 {
		goto L82
	} else {
		goto L438
	}
L438:
	;
	if v467 == int32(0) {
		goto L432
	} else {
		goto L439
	}
L439:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1782 = m.ExcPending
	if v1782 != 0 {
		goto L7
	} else {
		goto L440
	}
L440:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L7
	} else {
		goto L441
	}
L441:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_42), int32(0))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L7
	} else {
		goto L442
	}
L442:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1791 = m.ExcPending
	if v1791 != 0 {
		goto L7
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(814), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L7
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1798))) = int32(11)
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(v34)+984))
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+4)) = v1802
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v34)+988))
	*(*uint8)(unsafe.Add(mBase, uint32(v1798)+37)) = uint8(base.B2i32(v566 == int32(4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1798)+36)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+20)) = v1701
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+8)) = v1804
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+40)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1798)+24)) = v61
	if v58 != 0 {
		goto L81
	} else {
		goto L446
	}
L446:
	;
	v1816 = base.B2i32(v566 != int32(4))
	v1817 = int32(0)
	if v56|(v1816|base.B2i32(v1701 != v1817)) == v1817 {
		goto L80
	} else {
		goto L447
	}
L447:
	;
	if v54 != 0 {
		goto L79
	} else {
		goto L448
	}
L448:
	;
	if v61 != 0 {
		goto L449
	} else {
		goto L450
	}
L449:
	;
	v1824 = v1816
	goto L451
L450:
	;
	v1824 = int32(0)
	goto L451
L451:
	;
	if v1824 != 0 {
		goto L78
	} else {
		goto L452
	}
L452:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v1825 != l3 {
		goto L77
	} else {
		goto L453
	}
L453:
	;
	if v1723 != 0 {
		goto L76
	} else {
		goto L454
	}
L454:
	;
	F_transformWindowFuncCall(m, l0, v1798, v57)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L7
	} else {
		goto L455
	}
L455:
	;
	v1846 = v1798
	goto L86
L456:
	;
	v1831 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+567)))
	if v1831 != int32(1) {
		v1846 = v1750
		goto L86
	} else {
		goto L457
	}
L457:
	;
	v1836 = v1750
	goto L426
L458:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L7
	} else {
		goto L459
	}
L459:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1664)+12))
	v1881 = F_format_type_be(m, v1880)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L7
	} else {
		goto L460
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+64)) = v1881
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_43), v34-int32(-64))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L7
	} else {
		goto L461
	}
L461:
	;
	v1889 = F_exprLocation(m, v1669)
	mBase = m.M
	F_parser_errposition(m, l0, v1889)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L7
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(712), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L7
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
	F_errcode(m, int32(67141764))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L7
	} else {
		goto L465
	}
L465:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_44), int32(0))
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L7
	} else {
		goto L466
	}
L466:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+12))
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1701)+4))
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1908+v1909<<(uint(int32(2))%32)-int32(4))))
	v1916 = F_exprLocation(m, v1915)
	mBase = m.M
	F_parser_errposition(m, l0, v1916)
	mBase = m.M
	v1918 = m.ExcPending
	if v1918 != 0 {
		goto L7
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(738), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L7
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1930 = m.ExcPending
	if v1930 != 0 {
		goto L7
	} else {
		goto L470
	}
L470:
	;
	v1931 = F_NameListToString(m, l1)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L7
	} else {
		goto L471
	}
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+96)) = v1931
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_45), v34+int32(96))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L7
	} else {
		goto L472
	}
L472:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L7
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(793), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L7
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
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L7
	} else {
		goto L476
	}
L476:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_46), int32(0))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L7
	} else {
		goto L477
	}
L477:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L7
	} else {
		goto L478
	}
L478:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(799), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L7
	} else {
		goto L479
	}
L479:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L480:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L7
	} else {
		goto L481
	}
L481:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_47), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L7
	} else {
		goto L482
	}
L482:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L7
	} else {
		goto L483
	}
L483:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(847), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L7
	} else {
		goto L484
	}
L484:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L485:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L7
	} else {
		goto L486
	}
L486:
	;
	v1989 = F_NameListToString(m, l1)
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L7
	} else {
		goto L487
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v1989
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_45), v34+int32(80))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L7
	} else {
		goto L488
	}
L488:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L7
	} else {
		goto L489
	}
L489:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(858), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L7
	} else {
		goto L490
	}
L490:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L491:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L7
	} else {
		goto L492
	}
L492:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_48), int32(0))
	mBase = m.M
	v2014 = m.ExcPending
	if v2014 != 0 {
		goto L7
	} else {
		goto L493
	}
L493:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L7
	} else {
		goto L494
	}
L494:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(867), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L7
	} else {
		goto L495
	}
L495:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L496:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L7
	} else {
		goto L497
	}
L497:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_49), int32(0))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L7
	} else {
		goto L498
	}
L498:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L7
	} else {
		goto L499
	}
L499:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(876), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L7
	} else {
		goto L500
	}
L500:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L501:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2046 = m.ExcPending
	if v2046 != 0 {
		goto L7
	} else {
		goto L502
	}
L502:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_50), int32(0))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L7
	} else {
		goto L503
	}
L503:
	;
	F_errhint(m, int32(_a_F_ParseFuncOrColumn_51), int32(0))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L7
	} else {
		goto L504
	}
L504:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2056 = F_exprLocation(m, v2055)
	mBase = m.M
	F_parser_errposition(m, l0, v2056)
	mBase = m.M
	v2058 = m.ExcPending
	if v2058 != 0 {
		goto L7
	} else {
		goto L505
	}
L505:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(887), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L7
	} else {
		goto L506
	}
L506:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L507:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L7
	} else {
		goto L508
	}
L508:
	;
	F_errmsg(m, int32(_a_F_ParseFuncOrColumn_52), int32(0))
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L7
	} else {
		goto L509
	}
L509:
	;
	F_parser_errposition(m, l0, l6)
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L7
	} else {
		goto L510
	}
L510:
	;
	F_errfinish(m, int32(_a_F_ParseFuncOrColumn_3), int32(893), int32(_a_F_ParseFuncOrColumn_4))
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L7
	} else {
		goto L511
	}
L511:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_func_match_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v5 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = l2
	v13 = v5
	goto L4
L2:
	;
	v33 = v5
	goto L3
L3:
	;
	return v33
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v19 = F_can_coerce_type(m, l0, l1, v11+int32(32), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v33 = v28
	goto L3
L6:
	;
	return int32(0)
L7:
	;
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v11
	v28 = v13 + int32(1)
	goto L10
L9:
	;
	v28 = v13
	goto L10
L10:
	;
	if v15 != 0 {
		v11 = v15
		v13 = v28
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
}
func F_func_parallel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_func_parallel_0), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_func_parallel_1), int32(1946), int32(_a_F_func_parallel_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+102)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
func F_get_func_input_arg_names(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 == v4 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L37
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v117
	m.G0 = v11 + int32(16)
	return v112
L4:
	;
	v112 = int32(0)
	v117 = v4
	goto L3
L5:
	;
	goto L6
L6:
	;
	v16 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v20 != int32(1) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v23 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v24 != int32(25) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_deconstruct_array_builtin(m, v16, int32(25), v11+int32(8), int32(0), v11+int32(12))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if l1 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v55 = int32(0)
	if v53 <= v55 {
		v112 = v55
		v117 = v4
		goto L3
	} else {
		goto L22
	}
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v53 = v37
	v54 = v4
	goto L13
L15:
	;
	goto L16
L16:
	;
	v38 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v40 != int32(1) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v43 != v44 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v46 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v47 != int32(18) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v53 = v43
	v54 = v38 + int32(24)
	goto L13
L22:
	;
	v60 = F_palloc(m, v53<<(uint(int32(2))%32))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v62 <= int32(0) {
		v112 = v55
		v117 = v60
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v66 = int32(0)
	v67 = v55
	v70 = v62
	goto L25
L25:
	;
	if v54 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v112 = v105
	v117 = v60
	goto L3
L27:
	;
	v109 = v66 + int32(1)
	if v109 < v106 {
		v66 = v109
		v67 = v105
		v70 = v106
		goto L25
	} else {
		goto L36
	}
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v54))))
	v77 = v75 - int32(98)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v77))|base.B2i32(int32(1)<<(uint(v77)%32)&int32(_a_F_get_func_input_arg_names_0) == int32(0)) != 0 {
		v105 = v67
		v106 = v70
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v88 = int32(2)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v66<<(uint(v88)%32))))
	v96 = F_text_to_cstring(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v100 = v96
	goto L35
L34:
	;
	v100 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60+v67<<(uint(v88)%32)))) = v100
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v105 = v67 + int32(1)
	v106 = v104
	goto L27
L36:
	;
	goto L26
L37:
	;
	F_errmsg_internal(m, int32(_a_F_get_func_input_arg_names_1), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_get_func_input_arg_names_2), int32(1549), int32(_a_F_get_func_input_arg_names_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v142
	F_errmsg_internal(m, int32(_a_F_get_func_input_arg_names_4), v11)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_get_func_input_arg_names_2), int32(1559), int32(_a_F_get_func_input_arg_names_3))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makeFuncCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	v7 = F_palloc0(m, int32(36))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+12)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(76)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+20)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l2
		return v7
	}
}
